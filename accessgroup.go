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

// Creates a new Access group.
func (r *AccessGroupService) New(ctx context.Context, accountOrZone string, accountOrZoneID string, body AccessGroupNewParams, opts ...option.RequestOption) (res *AccessGroupNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessGroupNewResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/groups", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all Access groups.
func (r *AccessGroupService) List(ctx context.Context, accountOrZone string, accountOrZoneID string, opts ...option.RequestOption) (res *[]AccessGroupListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessGroupListResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/groups", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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
func (r *AccessGroupService) Replace(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, body AccessGroupReplaceParams, opts ...option.RequestOption) (res *AccessGroupReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessGroupReplaceResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/groups/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessGroupNewResponse struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessGroupNewResponseExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessGroupNewResponseInclude `json:"include"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	IsDefault []AccessGroupNewResponseIsDefault `json:"is_default"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []AccessGroupNewResponseRequire `json:"require"`
	UpdatedAt time.Time                       `json:"updated_at" format:"date-time"`
	JSON      accessGroupNewResponseJSON      `json:"-"`
}

// accessGroupNewResponseJSON contains the JSON metadata for the struct
// [AccessGroupNewResponse]
type accessGroupNewResponseJSON struct {
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

func (r *AccessGroupNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupNewResponseExcludeAccessEmailRule],
// [AccessGroupNewResponseExcludeAccessEmailListRule],
// [AccessGroupNewResponseExcludeAccessDomainRule],
// [AccessGroupNewResponseExcludeAccessEveryoneRule],
// [AccessGroupNewResponseExcludeAccessIPRule],
// [AccessGroupNewResponseExcludeAccessIPListRule],
// [AccessGroupNewResponseExcludeAccessCertificateRule],
// [AccessGroupNewResponseExcludeAccessAccessGroupRule],
// [AccessGroupNewResponseExcludeAccessAzureGroupRule],
// [AccessGroupNewResponseExcludeAccessGitHubOrganizationRule],
// [AccessGroupNewResponseExcludeAccessGsuiteGroupRule],
// [AccessGroupNewResponseExcludeAccessOktaGroupRule],
// [AccessGroupNewResponseExcludeAccessSamlGroupRule],
// [AccessGroupNewResponseExcludeAccessServiceTokenRule],
// [AccessGroupNewResponseExcludeAccessAnyValidServiceTokenRule],
// [AccessGroupNewResponseExcludeAccessExternalEvaluationRule],
// [AccessGroupNewResponseExcludeAccessCountryRule],
// [AccessGroupNewResponseExcludeAccessAuthenticationMethodRule] or
// [AccessGroupNewResponseExcludeAccessDevicePostureRule].
type AccessGroupNewResponseExclude interface {
	implementsAccessGroupNewResponseExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupNewResponseExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupNewResponseExcludeAccessEmailRule struct {
	Email AccessGroupNewResponseExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupNewResponseExcludeAccessEmailRuleJSON  `json:"-"`
}

// accessGroupNewResponseExcludeAccessEmailRuleJSON contains the JSON metadata for
// the struct [AccessGroupNewResponseExcludeAccessEmailRule]
type accessGroupNewResponseExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseExcludeAccessEmailRule) implementsAccessGroupNewResponseExclude() {}

type AccessGroupNewResponseExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                `json:"email,required" format:"email"`
	JSON  accessGroupNewResponseExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupNewResponseExcludeAccessEmailRuleEmailJSON contains the JSON metadata
// for the struct [AccessGroupNewResponseExcludeAccessEmailRuleEmail]
type accessGroupNewResponseExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupNewResponseExcludeAccessEmailListRule struct {
	EmailList AccessGroupNewResponseExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupNewResponseExcludeAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupNewResponseExcludeAccessEmailListRuleJSON contains the JSON metadata
// for the struct [AccessGroupNewResponseExcludeAccessEmailListRule]
type accessGroupNewResponseExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseExcludeAccessEmailListRule) implementsAccessGroupNewResponseExclude() {}

type AccessGroupNewResponseExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                        `json:"id,required"`
	JSON accessGroupNewResponseExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupNewResponseExcludeAccessEmailListRuleEmailListJSON contains the JSON
// metadata for the struct
// [AccessGroupNewResponseExcludeAccessEmailListRuleEmailList]
type accessGroupNewResponseExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupNewResponseExcludeAccessDomainRule struct {
	EmailDomain AccessGroupNewResponseExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupNewResponseExcludeAccessDomainRuleJSON        `json:"-"`
}

// accessGroupNewResponseExcludeAccessDomainRuleJSON contains the JSON metadata for
// the struct [AccessGroupNewResponseExcludeAccessDomainRule]
type accessGroupNewResponseExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseExcludeAccessDomainRule) implementsAccessGroupNewResponseExclude() {}

type AccessGroupNewResponseExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                       `json:"domain,required"`
	JSON   accessGroupNewResponseExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupNewResponseExcludeAccessDomainRuleEmailDomainJSON contains the JSON
// metadata for the struct
// [AccessGroupNewResponseExcludeAccessDomainRuleEmailDomain]
type accessGroupNewResponseExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupNewResponseExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                         `json:"everyone,required"`
	JSON     accessGroupNewResponseExcludeAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupNewResponseExcludeAccessEveryoneRuleJSON contains the JSON metadata
// for the struct [AccessGroupNewResponseExcludeAccessEveryoneRule]
type accessGroupNewResponseExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseExcludeAccessEveryoneRule) implementsAccessGroupNewResponseExclude() {}

// Matches an IP address block.
type AccessGroupNewResponseExcludeAccessIPRule struct {
	IP   AccessGroupNewResponseExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupNewResponseExcludeAccessIPRuleJSON `json:"-"`
}

// accessGroupNewResponseExcludeAccessIPRuleJSON contains the JSON metadata for the
// struct [AccessGroupNewResponseExcludeAccessIPRule]
type accessGroupNewResponseExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseExcludeAccessIPRule) implementsAccessGroupNewResponseExclude() {}

type AccessGroupNewResponseExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                          `json:"ip,required"`
	JSON accessGroupNewResponseExcludeAccessIPRuleIPJSON `json:"-"`
}

// accessGroupNewResponseExcludeAccessIPRuleIPJSON contains the JSON metadata for
// the struct [AccessGroupNewResponseExcludeAccessIPRuleIP]
type accessGroupNewResponseExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupNewResponseExcludeAccessIPListRule struct {
	IPList AccessGroupNewResponseExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupNewResponseExcludeAccessIPListRuleJSON   `json:"-"`
}

// accessGroupNewResponseExcludeAccessIPListRuleJSON contains the JSON metadata for
// the struct [AccessGroupNewResponseExcludeAccessIPListRule]
type accessGroupNewResponseExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseExcludeAccessIPListRule) implementsAccessGroupNewResponseExclude() {}

type AccessGroupNewResponseExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                  `json:"id,required"`
	JSON accessGroupNewResponseExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupNewResponseExcludeAccessIPListRuleIPListJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseExcludeAccessIPListRuleIPList]
type accessGroupNewResponseExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupNewResponseExcludeAccessCertificateRule struct {
	Certificate interface{}                                            `json:"certificate,required"`
	JSON        accessGroupNewResponseExcludeAccessCertificateRuleJSON `json:"-"`
}

// accessGroupNewResponseExcludeAccessCertificateRuleJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseExcludeAccessCertificateRule]
type accessGroupNewResponseExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseExcludeAccessCertificateRule) implementsAccessGroupNewResponseExclude() {
}

// Matches an Access group.
type AccessGroupNewResponseExcludeAccessAccessGroupRule struct {
	Group AccessGroupNewResponseExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupNewResponseExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupNewResponseExcludeAccessAccessGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseExcludeAccessAccessGroupRule]
type accessGroupNewResponseExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseExcludeAccessAccessGroupRule) implementsAccessGroupNewResponseExclude() {
}

type AccessGroupNewResponseExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                      `json:"id,required"`
	JSON accessGroupNewResponseExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupNewResponseExcludeAccessAccessGroupRuleGroupJSON contains the JSON
// metadata for the struct
// [AccessGroupNewResponseExcludeAccessAccessGroupRuleGroup]
type accessGroupNewResponseExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupNewResponseExcludeAccessAzureGroupRule struct {
	AzureAd AccessGroupNewResponseExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupNewResponseExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupNewResponseExcludeAccessAzureGroupRuleJSON contains the JSON metadata
// for the struct [AccessGroupNewResponseExcludeAccessAzureGroupRule]
type accessGroupNewResponseExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseExcludeAccessAzureGroupRule) implementsAccessGroupNewResponseExclude() {
}

type AccessGroupNewResponseExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                       `json:"connection_id,required"`
	JSON         accessGroupNewResponseExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupNewResponseExcludeAccessAzureGroupRuleAzureAdJSON contains the JSON
// metadata for the struct
// [AccessGroupNewResponseExcludeAccessAzureGroupRuleAzureAd]
type accessGroupNewResponseExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupNewResponseExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupNewResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupNewResponseExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupNewResponseExcludeAccessGitHubOrganizationRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupNewResponseExcludeAccessGitHubOrganizationRule]
type accessGroupNewResponseExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseExcludeAccessGitHubOrganizationRule) implementsAccessGroupNewResponseExclude() {
}

type AccessGroupNewResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                          `json:"name,required"`
	JSON accessGroupNewResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupNewResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupNewResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupNewResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupNewResponseExcludeAccessGsuiteGroupRule struct {
	Gsuite AccessGroupNewResponseExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupNewResponseExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupNewResponseExcludeAccessGsuiteGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseExcludeAccessGsuiteGroupRule]
type accessGroupNewResponseExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseExcludeAccessGsuiteGroupRule) implementsAccessGroupNewResponseExclude() {
}

type AccessGroupNewResponseExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                       `json:"email,required"`
	JSON  accessGroupNewResponseExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupNewResponseExcludeAccessGsuiteGroupRuleGsuiteJSON contains the JSON
// metadata for the struct
// [AccessGroupNewResponseExcludeAccessGsuiteGroupRuleGsuite]
type accessGroupNewResponseExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupNewResponseExcludeAccessOktaGroupRule struct {
	Okta AccessGroupNewResponseExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupNewResponseExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupNewResponseExcludeAccessOktaGroupRuleJSON contains the JSON metadata
// for the struct [AccessGroupNewResponseExcludeAccessOktaGroupRule]
type accessGroupNewResponseExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseExcludeAccessOktaGroupRule) implementsAccessGroupNewResponseExclude() {}

type AccessGroupNewResponseExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                   `json:"email,required"`
	JSON  accessGroupNewResponseExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupNewResponseExcludeAccessOktaGroupRuleOktaJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseExcludeAccessOktaGroupRuleOkta]
type accessGroupNewResponseExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupNewResponseExcludeAccessSamlGroupRule struct {
	Saml AccessGroupNewResponseExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupNewResponseExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupNewResponseExcludeAccessSamlGroupRuleJSON contains the JSON metadata
// for the struct [AccessGroupNewResponseExcludeAccessSamlGroupRule]
type accessGroupNewResponseExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseExcludeAccessSamlGroupRule) implementsAccessGroupNewResponseExclude() {}

type AccessGroupNewResponseExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                   `json:"attribute_value,required"`
	JSON           accessGroupNewResponseExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupNewResponseExcludeAccessSamlGroupRuleSamlJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseExcludeAccessSamlGroupRuleSaml]
type accessGroupNewResponseExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupNewResponseExcludeAccessServiceTokenRule struct {
	ServiceToken AccessGroupNewResponseExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupNewResponseExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupNewResponseExcludeAccessServiceTokenRuleJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseExcludeAccessServiceTokenRule]
type accessGroupNewResponseExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseExcludeAccessServiceTokenRule) implementsAccessGroupNewResponseExclude() {
}

type AccessGroupNewResponseExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                              `json:"token_id,required"`
	JSON    accessGroupNewResponseExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupNewResponseExcludeAccessServiceTokenRuleServiceTokenJSON contains the
// JSON metadata for the struct
// [AccessGroupNewResponseExcludeAccessServiceTokenRuleServiceToken]
type accessGroupNewResponseExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupNewResponseExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                     `json:"any_valid_service_token,required"`
	JSON                 accessGroupNewResponseExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupNewResponseExcludeAccessAnyValidServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupNewResponseExcludeAccessAnyValidServiceTokenRule]
type accessGroupNewResponseExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseExcludeAccessAnyValidServiceTokenRule) implementsAccessGroupNewResponseExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupNewResponseExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupNewResponseExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupNewResponseExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupNewResponseExcludeAccessExternalEvaluationRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupNewResponseExcludeAccessExternalEvaluationRule]
type accessGroupNewResponseExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseExcludeAccessExternalEvaluationRule) implementsAccessGroupNewResponseExclude() {
}

type AccessGroupNewResponseExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                          `json:"keys_url,required"`
	JSON    accessGroupNewResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupNewResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupNewResponseExcludeAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupNewResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupNewResponseExcludeAccessCountryRule struct {
	Geo  AccessGroupNewResponseExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupNewResponseExcludeAccessCountryRuleJSON `json:"-"`
}

// accessGroupNewResponseExcludeAccessCountryRuleJSON contains the JSON metadata
// for the struct [AccessGroupNewResponseExcludeAccessCountryRule]
type accessGroupNewResponseExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseExcludeAccessCountryRule) implementsAccessGroupNewResponseExclude() {}

type AccessGroupNewResponseExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                `json:"country_code,required"`
	JSON        accessGroupNewResponseExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupNewResponseExcludeAccessCountryRuleGeoJSON contains the JSON metadata
// for the struct [AccessGroupNewResponseExcludeAccessCountryRuleGeo]
type accessGroupNewResponseExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupNewResponseExcludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupNewResponseExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupNewResponseExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupNewResponseExcludeAccessAuthenticationMethodRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupNewResponseExcludeAccessAuthenticationMethodRule]
type accessGroupNewResponseExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseExcludeAccessAuthenticationMethodRule) implementsAccessGroupNewResponseExclude() {
}

type AccessGroupNewResponseExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                    `json:"auth_method,required"`
	JSON       accessGroupNewResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupNewResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupNewResponseExcludeAccessAuthenticationMethodRuleAuthMethod]
type accessGroupNewResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupNewResponseExcludeAccessDevicePostureRule struct {
	DevicePosture AccessGroupNewResponseExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupNewResponseExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupNewResponseExcludeAccessDevicePostureRuleJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseExcludeAccessDevicePostureRule]
type accessGroupNewResponseExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseExcludeAccessDevicePostureRule) implementsAccessGroupNewResponseExclude() {
}

type AccessGroupNewResponseExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                `json:"integration_uid,required"`
	JSON           accessGroupNewResponseExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupNewResponseExcludeAccessDevicePostureRuleDevicePostureJSON contains
// the JSON metadata for the struct
// [AccessGroupNewResponseExcludeAccessDevicePostureRuleDevicePosture]
type accessGroupNewResponseExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupNewResponseExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupNewResponseIncludeAccessEmailRule],
// [AccessGroupNewResponseIncludeAccessEmailListRule],
// [AccessGroupNewResponseIncludeAccessDomainRule],
// [AccessGroupNewResponseIncludeAccessEveryoneRule],
// [AccessGroupNewResponseIncludeAccessIPRule],
// [AccessGroupNewResponseIncludeAccessIPListRule],
// [AccessGroupNewResponseIncludeAccessCertificateRule],
// [AccessGroupNewResponseIncludeAccessAccessGroupRule],
// [AccessGroupNewResponseIncludeAccessAzureGroupRule],
// [AccessGroupNewResponseIncludeAccessGitHubOrganizationRule],
// [AccessGroupNewResponseIncludeAccessGsuiteGroupRule],
// [AccessGroupNewResponseIncludeAccessOktaGroupRule],
// [AccessGroupNewResponseIncludeAccessSamlGroupRule],
// [AccessGroupNewResponseIncludeAccessServiceTokenRule],
// [AccessGroupNewResponseIncludeAccessAnyValidServiceTokenRule],
// [AccessGroupNewResponseIncludeAccessExternalEvaluationRule],
// [AccessGroupNewResponseIncludeAccessCountryRule],
// [AccessGroupNewResponseIncludeAccessAuthenticationMethodRule] or
// [AccessGroupNewResponseIncludeAccessDevicePostureRule].
type AccessGroupNewResponseInclude interface {
	implementsAccessGroupNewResponseInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupNewResponseInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupNewResponseIncludeAccessEmailRule struct {
	Email AccessGroupNewResponseIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupNewResponseIncludeAccessEmailRuleJSON  `json:"-"`
}

// accessGroupNewResponseIncludeAccessEmailRuleJSON contains the JSON metadata for
// the struct [AccessGroupNewResponseIncludeAccessEmailRule]
type accessGroupNewResponseIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIncludeAccessEmailRule) implementsAccessGroupNewResponseInclude() {}

type AccessGroupNewResponseIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                `json:"email,required" format:"email"`
	JSON  accessGroupNewResponseIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupNewResponseIncludeAccessEmailRuleEmailJSON contains the JSON metadata
// for the struct [AccessGroupNewResponseIncludeAccessEmailRuleEmail]
type accessGroupNewResponseIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupNewResponseIncludeAccessEmailListRule struct {
	EmailList AccessGroupNewResponseIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupNewResponseIncludeAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupNewResponseIncludeAccessEmailListRuleJSON contains the JSON metadata
// for the struct [AccessGroupNewResponseIncludeAccessEmailListRule]
type accessGroupNewResponseIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIncludeAccessEmailListRule) implementsAccessGroupNewResponseInclude() {}

type AccessGroupNewResponseIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                        `json:"id,required"`
	JSON accessGroupNewResponseIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupNewResponseIncludeAccessEmailListRuleEmailListJSON contains the JSON
// metadata for the struct
// [AccessGroupNewResponseIncludeAccessEmailListRuleEmailList]
type accessGroupNewResponseIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupNewResponseIncludeAccessDomainRule struct {
	EmailDomain AccessGroupNewResponseIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupNewResponseIncludeAccessDomainRuleJSON        `json:"-"`
}

// accessGroupNewResponseIncludeAccessDomainRuleJSON contains the JSON metadata for
// the struct [AccessGroupNewResponseIncludeAccessDomainRule]
type accessGroupNewResponseIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIncludeAccessDomainRule) implementsAccessGroupNewResponseInclude() {}

type AccessGroupNewResponseIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                       `json:"domain,required"`
	JSON   accessGroupNewResponseIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupNewResponseIncludeAccessDomainRuleEmailDomainJSON contains the JSON
// metadata for the struct
// [AccessGroupNewResponseIncludeAccessDomainRuleEmailDomain]
type accessGroupNewResponseIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupNewResponseIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                         `json:"everyone,required"`
	JSON     accessGroupNewResponseIncludeAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupNewResponseIncludeAccessEveryoneRuleJSON contains the JSON metadata
// for the struct [AccessGroupNewResponseIncludeAccessEveryoneRule]
type accessGroupNewResponseIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIncludeAccessEveryoneRule) implementsAccessGroupNewResponseInclude() {}

// Matches an IP address block.
type AccessGroupNewResponseIncludeAccessIPRule struct {
	IP   AccessGroupNewResponseIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupNewResponseIncludeAccessIPRuleJSON `json:"-"`
}

// accessGroupNewResponseIncludeAccessIPRuleJSON contains the JSON metadata for the
// struct [AccessGroupNewResponseIncludeAccessIPRule]
type accessGroupNewResponseIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIncludeAccessIPRule) implementsAccessGroupNewResponseInclude() {}

type AccessGroupNewResponseIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                          `json:"ip,required"`
	JSON accessGroupNewResponseIncludeAccessIPRuleIPJSON `json:"-"`
}

// accessGroupNewResponseIncludeAccessIPRuleIPJSON contains the JSON metadata for
// the struct [AccessGroupNewResponseIncludeAccessIPRuleIP]
type accessGroupNewResponseIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupNewResponseIncludeAccessIPListRule struct {
	IPList AccessGroupNewResponseIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupNewResponseIncludeAccessIPListRuleJSON   `json:"-"`
}

// accessGroupNewResponseIncludeAccessIPListRuleJSON contains the JSON metadata for
// the struct [AccessGroupNewResponseIncludeAccessIPListRule]
type accessGroupNewResponseIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIncludeAccessIPListRule) implementsAccessGroupNewResponseInclude() {}

type AccessGroupNewResponseIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                  `json:"id,required"`
	JSON accessGroupNewResponseIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupNewResponseIncludeAccessIPListRuleIPListJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseIncludeAccessIPListRuleIPList]
type accessGroupNewResponseIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupNewResponseIncludeAccessCertificateRule struct {
	Certificate interface{}                                            `json:"certificate,required"`
	JSON        accessGroupNewResponseIncludeAccessCertificateRuleJSON `json:"-"`
}

// accessGroupNewResponseIncludeAccessCertificateRuleJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseIncludeAccessCertificateRule]
type accessGroupNewResponseIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIncludeAccessCertificateRule) implementsAccessGroupNewResponseInclude() {
}

// Matches an Access group.
type AccessGroupNewResponseIncludeAccessAccessGroupRule struct {
	Group AccessGroupNewResponseIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupNewResponseIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupNewResponseIncludeAccessAccessGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseIncludeAccessAccessGroupRule]
type accessGroupNewResponseIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIncludeAccessAccessGroupRule) implementsAccessGroupNewResponseInclude() {
}

type AccessGroupNewResponseIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                      `json:"id,required"`
	JSON accessGroupNewResponseIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupNewResponseIncludeAccessAccessGroupRuleGroupJSON contains the JSON
// metadata for the struct
// [AccessGroupNewResponseIncludeAccessAccessGroupRuleGroup]
type accessGroupNewResponseIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupNewResponseIncludeAccessAzureGroupRule struct {
	AzureAd AccessGroupNewResponseIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupNewResponseIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupNewResponseIncludeAccessAzureGroupRuleJSON contains the JSON metadata
// for the struct [AccessGroupNewResponseIncludeAccessAzureGroupRule]
type accessGroupNewResponseIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIncludeAccessAzureGroupRule) implementsAccessGroupNewResponseInclude() {
}

type AccessGroupNewResponseIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                       `json:"connection_id,required"`
	JSON         accessGroupNewResponseIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupNewResponseIncludeAccessAzureGroupRuleAzureAdJSON contains the JSON
// metadata for the struct
// [AccessGroupNewResponseIncludeAccessAzureGroupRuleAzureAd]
type accessGroupNewResponseIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupNewResponseIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupNewResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupNewResponseIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupNewResponseIncludeAccessGitHubOrganizationRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupNewResponseIncludeAccessGitHubOrganizationRule]
type accessGroupNewResponseIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIncludeAccessGitHubOrganizationRule) implementsAccessGroupNewResponseInclude() {
}

type AccessGroupNewResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                          `json:"name,required"`
	JSON accessGroupNewResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupNewResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupNewResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupNewResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupNewResponseIncludeAccessGsuiteGroupRule struct {
	Gsuite AccessGroupNewResponseIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupNewResponseIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupNewResponseIncludeAccessGsuiteGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseIncludeAccessGsuiteGroupRule]
type accessGroupNewResponseIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIncludeAccessGsuiteGroupRule) implementsAccessGroupNewResponseInclude() {
}

type AccessGroupNewResponseIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                       `json:"email,required"`
	JSON  accessGroupNewResponseIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupNewResponseIncludeAccessGsuiteGroupRuleGsuiteJSON contains the JSON
// metadata for the struct
// [AccessGroupNewResponseIncludeAccessGsuiteGroupRuleGsuite]
type accessGroupNewResponseIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupNewResponseIncludeAccessOktaGroupRule struct {
	Okta AccessGroupNewResponseIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupNewResponseIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupNewResponseIncludeAccessOktaGroupRuleJSON contains the JSON metadata
// for the struct [AccessGroupNewResponseIncludeAccessOktaGroupRule]
type accessGroupNewResponseIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIncludeAccessOktaGroupRule) implementsAccessGroupNewResponseInclude() {}

type AccessGroupNewResponseIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                   `json:"email,required"`
	JSON  accessGroupNewResponseIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupNewResponseIncludeAccessOktaGroupRuleOktaJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseIncludeAccessOktaGroupRuleOkta]
type accessGroupNewResponseIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupNewResponseIncludeAccessSamlGroupRule struct {
	Saml AccessGroupNewResponseIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupNewResponseIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupNewResponseIncludeAccessSamlGroupRuleJSON contains the JSON metadata
// for the struct [AccessGroupNewResponseIncludeAccessSamlGroupRule]
type accessGroupNewResponseIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIncludeAccessSamlGroupRule) implementsAccessGroupNewResponseInclude() {}

type AccessGroupNewResponseIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                   `json:"attribute_value,required"`
	JSON           accessGroupNewResponseIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupNewResponseIncludeAccessSamlGroupRuleSamlJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseIncludeAccessSamlGroupRuleSaml]
type accessGroupNewResponseIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupNewResponseIncludeAccessServiceTokenRule struct {
	ServiceToken AccessGroupNewResponseIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupNewResponseIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupNewResponseIncludeAccessServiceTokenRuleJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseIncludeAccessServiceTokenRule]
type accessGroupNewResponseIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIncludeAccessServiceTokenRule) implementsAccessGroupNewResponseInclude() {
}

type AccessGroupNewResponseIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                              `json:"token_id,required"`
	JSON    accessGroupNewResponseIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupNewResponseIncludeAccessServiceTokenRuleServiceTokenJSON contains the
// JSON metadata for the struct
// [AccessGroupNewResponseIncludeAccessServiceTokenRuleServiceToken]
type accessGroupNewResponseIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupNewResponseIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                     `json:"any_valid_service_token,required"`
	JSON                 accessGroupNewResponseIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupNewResponseIncludeAccessAnyValidServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupNewResponseIncludeAccessAnyValidServiceTokenRule]
type accessGroupNewResponseIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIncludeAccessAnyValidServiceTokenRule) implementsAccessGroupNewResponseInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupNewResponseIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupNewResponseIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupNewResponseIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupNewResponseIncludeAccessExternalEvaluationRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupNewResponseIncludeAccessExternalEvaluationRule]
type accessGroupNewResponseIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIncludeAccessExternalEvaluationRule) implementsAccessGroupNewResponseInclude() {
}

type AccessGroupNewResponseIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                          `json:"keys_url,required"`
	JSON    accessGroupNewResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupNewResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupNewResponseIncludeAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupNewResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupNewResponseIncludeAccessCountryRule struct {
	Geo  AccessGroupNewResponseIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupNewResponseIncludeAccessCountryRuleJSON `json:"-"`
}

// accessGroupNewResponseIncludeAccessCountryRuleJSON contains the JSON metadata
// for the struct [AccessGroupNewResponseIncludeAccessCountryRule]
type accessGroupNewResponseIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIncludeAccessCountryRule) implementsAccessGroupNewResponseInclude() {}

type AccessGroupNewResponseIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                `json:"country_code,required"`
	JSON        accessGroupNewResponseIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupNewResponseIncludeAccessCountryRuleGeoJSON contains the JSON metadata
// for the struct [AccessGroupNewResponseIncludeAccessCountryRuleGeo]
type accessGroupNewResponseIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupNewResponseIncludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupNewResponseIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupNewResponseIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupNewResponseIncludeAccessAuthenticationMethodRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupNewResponseIncludeAccessAuthenticationMethodRule]
type accessGroupNewResponseIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIncludeAccessAuthenticationMethodRule) implementsAccessGroupNewResponseInclude() {
}

type AccessGroupNewResponseIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                    `json:"auth_method,required"`
	JSON       accessGroupNewResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupNewResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupNewResponseIncludeAccessAuthenticationMethodRuleAuthMethod]
type accessGroupNewResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupNewResponseIncludeAccessDevicePostureRule struct {
	DevicePosture AccessGroupNewResponseIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupNewResponseIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupNewResponseIncludeAccessDevicePostureRuleJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseIncludeAccessDevicePostureRule]
type accessGroupNewResponseIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIncludeAccessDevicePostureRule) implementsAccessGroupNewResponseInclude() {
}

type AccessGroupNewResponseIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                `json:"integration_uid,required"`
	JSON           accessGroupNewResponseIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupNewResponseIncludeAccessDevicePostureRuleDevicePostureJSON contains
// the JSON metadata for the struct
// [AccessGroupNewResponseIncludeAccessDevicePostureRuleDevicePosture]
type accessGroupNewResponseIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupNewResponseIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupNewResponseIsDefaultAccessEmailRule],
// [AccessGroupNewResponseIsDefaultAccessEmailListRule],
// [AccessGroupNewResponseIsDefaultAccessDomainRule],
// [AccessGroupNewResponseIsDefaultAccessEveryoneRule],
// [AccessGroupNewResponseIsDefaultAccessIPRule],
// [AccessGroupNewResponseIsDefaultAccessIPListRule],
// [AccessGroupNewResponseIsDefaultAccessCertificateRule],
// [AccessGroupNewResponseIsDefaultAccessAccessGroupRule],
// [AccessGroupNewResponseIsDefaultAccessAzureGroupRule],
// [AccessGroupNewResponseIsDefaultAccessGitHubOrganizationRule],
// [AccessGroupNewResponseIsDefaultAccessGsuiteGroupRule],
// [AccessGroupNewResponseIsDefaultAccessOktaGroupRule],
// [AccessGroupNewResponseIsDefaultAccessSamlGroupRule],
// [AccessGroupNewResponseIsDefaultAccessServiceTokenRule],
// [AccessGroupNewResponseIsDefaultAccessAnyValidServiceTokenRule],
// [AccessGroupNewResponseIsDefaultAccessExternalEvaluationRule],
// [AccessGroupNewResponseIsDefaultAccessCountryRule],
// [AccessGroupNewResponseIsDefaultAccessAuthenticationMethodRule] or
// [AccessGroupNewResponseIsDefaultAccessDevicePostureRule].
type AccessGroupNewResponseIsDefault interface {
	implementsAccessGroupNewResponseIsDefault()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupNewResponseIsDefault)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupNewResponseIsDefaultAccessEmailRule struct {
	Email AccessGroupNewResponseIsDefaultAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupNewResponseIsDefaultAccessEmailRuleJSON  `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessEmailRuleJSON contains the JSON metadata
// for the struct [AccessGroupNewResponseIsDefaultAccessEmailRule]
type accessGroupNewResponseIsDefaultAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIsDefaultAccessEmailRule) implementsAccessGroupNewResponseIsDefault() {}

type AccessGroupNewResponseIsDefaultAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                  `json:"email,required" format:"email"`
	JSON  accessGroupNewResponseIsDefaultAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseIsDefaultAccessEmailRuleEmail]
type accessGroupNewResponseIsDefaultAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupNewResponseIsDefaultAccessEmailListRule struct {
	EmailList AccessGroupNewResponseIsDefaultAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupNewResponseIsDefaultAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessEmailListRuleJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseIsDefaultAccessEmailListRule]
type accessGroupNewResponseIsDefaultAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIsDefaultAccessEmailListRule) implementsAccessGroupNewResponseIsDefault() {
}

type AccessGroupNewResponseIsDefaultAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                          `json:"id,required"`
	JSON accessGroupNewResponseIsDefaultAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessEmailListRuleEmailListJSON contains the
// JSON metadata for the struct
// [AccessGroupNewResponseIsDefaultAccessEmailListRuleEmailList]
type accessGroupNewResponseIsDefaultAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupNewResponseIsDefaultAccessDomainRule struct {
	EmailDomain AccessGroupNewResponseIsDefaultAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupNewResponseIsDefaultAccessDomainRuleJSON        `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessDomainRuleJSON contains the JSON metadata
// for the struct [AccessGroupNewResponseIsDefaultAccessDomainRule]
type accessGroupNewResponseIsDefaultAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIsDefaultAccessDomainRule) implementsAccessGroupNewResponseIsDefault() {
}

type AccessGroupNewResponseIsDefaultAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                         `json:"domain,required"`
	JSON   accessGroupNewResponseIsDefaultAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessDomainRuleEmailDomainJSON contains the JSON
// metadata for the struct
// [AccessGroupNewResponseIsDefaultAccessDomainRuleEmailDomain]
type accessGroupNewResponseIsDefaultAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupNewResponseIsDefaultAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                           `json:"everyone,required"`
	JSON     accessGroupNewResponseIsDefaultAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessEveryoneRuleJSON contains the JSON metadata
// for the struct [AccessGroupNewResponseIsDefaultAccessEveryoneRule]
type accessGroupNewResponseIsDefaultAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIsDefaultAccessEveryoneRule) implementsAccessGroupNewResponseIsDefault() {
}

// Matches an IP address block.
type AccessGroupNewResponseIsDefaultAccessIPRule struct {
	IP   AccessGroupNewResponseIsDefaultAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupNewResponseIsDefaultAccessIPRuleJSON `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessIPRuleJSON contains the JSON metadata for
// the struct [AccessGroupNewResponseIsDefaultAccessIPRule]
type accessGroupNewResponseIsDefaultAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIsDefaultAccessIPRule) implementsAccessGroupNewResponseIsDefault() {}

type AccessGroupNewResponseIsDefaultAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                            `json:"ip,required"`
	JSON accessGroupNewResponseIsDefaultAccessIPRuleIPJSON `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessIPRuleIPJSON contains the JSON metadata for
// the struct [AccessGroupNewResponseIsDefaultAccessIPRuleIP]
type accessGroupNewResponseIsDefaultAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupNewResponseIsDefaultAccessIPListRule struct {
	IPList AccessGroupNewResponseIsDefaultAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupNewResponseIsDefaultAccessIPListRuleJSON   `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessIPListRuleJSON contains the JSON metadata
// for the struct [AccessGroupNewResponseIsDefaultAccessIPListRule]
type accessGroupNewResponseIsDefaultAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIsDefaultAccessIPListRule) implementsAccessGroupNewResponseIsDefault() {
}

type AccessGroupNewResponseIsDefaultAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                    `json:"id,required"`
	JSON accessGroupNewResponseIsDefaultAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessIPListRuleIPListJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseIsDefaultAccessIPListRuleIPList]
type accessGroupNewResponseIsDefaultAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupNewResponseIsDefaultAccessCertificateRule struct {
	Certificate interface{}                                              `json:"certificate,required"`
	JSON        accessGroupNewResponseIsDefaultAccessCertificateRuleJSON `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessCertificateRuleJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseIsDefaultAccessCertificateRule]
type accessGroupNewResponseIsDefaultAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIsDefaultAccessCertificateRule) implementsAccessGroupNewResponseIsDefault() {
}

// Matches an Access group.
type AccessGroupNewResponseIsDefaultAccessAccessGroupRule struct {
	Group AccessGroupNewResponseIsDefaultAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupNewResponseIsDefaultAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessAccessGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseIsDefaultAccessAccessGroupRule]
type accessGroupNewResponseIsDefaultAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIsDefaultAccessAccessGroupRule) implementsAccessGroupNewResponseIsDefault() {
}

type AccessGroupNewResponseIsDefaultAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                        `json:"id,required"`
	JSON accessGroupNewResponseIsDefaultAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessAccessGroupRuleGroupJSON contains the JSON
// metadata for the struct
// [AccessGroupNewResponseIsDefaultAccessAccessGroupRuleGroup]
type accessGroupNewResponseIsDefaultAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupNewResponseIsDefaultAccessAzureGroupRule struct {
	AzureAd AccessGroupNewResponseIsDefaultAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupNewResponseIsDefaultAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseIsDefaultAccessAzureGroupRule]
type accessGroupNewResponseIsDefaultAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIsDefaultAccessAzureGroupRule) implementsAccessGroupNewResponseIsDefault() {
}

type AccessGroupNewResponseIsDefaultAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                         `json:"connection_id,required"`
	JSON         accessGroupNewResponseIsDefaultAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessAzureGroupRuleAzureAdJSON contains the JSON
// metadata for the struct
// [AccessGroupNewResponseIsDefaultAccessAzureGroupRuleAzureAd]
type accessGroupNewResponseIsDefaultAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupNewResponseIsDefaultAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupNewResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupNewResponseIsDefaultAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessGitHubOrganizationRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupNewResponseIsDefaultAccessGitHubOrganizationRule]
type accessGroupNewResponseIsDefaultAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIsDefaultAccessGitHubOrganizationRule) implementsAccessGroupNewResponseIsDefault() {
}

type AccessGroupNewResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                            `json:"name,required"`
	JSON accessGroupNewResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupNewResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupNewResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupNewResponseIsDefaultAccessGsuiteGroupRule struct {
	Gsuite AccessGroupNewResponseIsDefaultAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupNewResponseIsDefaultAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessGsuiteGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseIsDefaultAccessGsuiteGroupRule]
type accessGroupNewResponseIsDefaultAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIsDefaultAccessGsuiteGroupRule) implementsAccessGroupNewResponseIsDefault() {
}

type AccessGroupNewResponseIsDefaultAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                         `json:"email,required"`
	JSON  accessGroupNewResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON contains the JSON
// metadata for the struct
// [AccessGroupNewResponseIsDefaultAccessGsuiteGroupRuleGsuite]
type accessGroupNewResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupNewResponseIsDefaultAccessOktaGroupRule struct {
	Okta AccessGroupNewResponseIsDefaultAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupNewResponseIsDefaultAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseIsDefaultAccessOktaGroupRule]
type accessGroupNewResponseIsDefaultAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIsDefaultAccessOktaGroupRule) implementsAccessGroupNewResponseIsDefault() {
}

type AccessGroupNewResponseIsDefaultAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                     `json:"email,required"`
	JSON  accessGroupNewResponseIsDefaultAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessOktaGroupRuleOktaJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseIsDefaultAccessOktaGroupRuleOkta]
type accessGroupNewResponseIsDefaultAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupNewResponseIsDefaultAccessSamlGroupRule struct {
	Saml AccessGroupNewResponseIsDefaultAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupNewResponseIsDefaultAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseIsDefaultAccessSamlGroupRule]
type accessGroupNewResponseIsDefaultAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIsDefaultAccessSamlGroupRule) implementsAccessGroupNewResponseIsDefault() {
}

type AccessGroupNewResponseIsDefaultAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                     `json:"attribute_value,required"`
	JSON           accessGroupNewResponseIsDefaultAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessSamlGroupRuleSamlJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseIsDefaultAccessSamlGroupRuleSaml]
type accessGroupNewResponseIsDefaultAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupNewResponseIsDefaultAccessServiceTokenRule struct {
	ServiceToken AccessGroupNewResponseIsDefaultAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupNewResponseIsDefaultAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessServiceTokenRuleJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseIsDefaultAccessServiceTokenRule]
type accessGroupNewResponseIsDefaultAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIsDefaultAccessServiceTokenRule) implementsAccessGroupNewResponseIsDefault() {
}

type AccessGroupNewResponseIsDefaultAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                `json:"token_id,required"`
	JSON    accessGroupNewResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON contains
// the JSON metadata for the struct
// [AccessGroupNewResponseIsDefaultAccessServiceTokenRuleServiceToken]
type accessGroupNewResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupNewResponseIsDefaultAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                       `json:"any_valid_service_token,required"`
	JSON                 accessGroupNewResponseIsDefaultAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessAnyValidServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupNewResponseIsDefaultAccessAnyValidServiceTokenRule]
type accessGroupNewResponseIsDefaultAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIsDefaultAccessAnyValidServiceTokenRule) implementsAccessGroupNewResponseIsDefault() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupNewResponseIsDefaultAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupNewResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupNewResponseIsDefaultAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessExternalEvaluationRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupNewResponseIsDefaultAccessExternalEvaluationRule]
type accessGroupNewResponseIsDefaultAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIsDefaultAccessExternalEvaluationRule) implementsAccessGroupNewResponseIsDefault() {
}

type AccessGroupNewResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                            `json:"keys_url,required"`
	JSON    accessGroupNewResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupNewResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupNewResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupNewResponseIsDefaultAccessCountryRule struct {
	Geo  AccessGroupNewResponseIsDefaultAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupNewResponseIsDefaultAccessCountryRuleJSON `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessCountryRuleJSON contains the JSON metadata
// for the struct [AccessGroupNewResponseIsDefaultAccessCountryRule]
type accessGroupNewResponseIsDefaultAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIsDefaultAccessCountryRule) implementsAccessGroupNewResponseIsDefault() {
}

type AccessGroupNewResponseIsDefaultAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                  `json:"country_code,required"`
	JSON        accessGroupNewResponseIsDefaultAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseIsDefaultAccessCountryRuleGeo]
type accessGroupNewResponseIsDefaultAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupNewResponseIsDefaultAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupNewResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupNewResponseIsDefaultAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessAuthenticationMethodRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupNewResponseIsDefaultAccessAuthenticationMethodRule]
type accessGroupNewResponseIsDefaultAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIsDefaultAccessAuthenticationMethodRule) implementsAccessGroupNewResponseIsDefault() {
}

type AccessGroupNewResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                      `json:"auth_method,required"`
	JSON       accessGroupNewResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupNewResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod]
type accessGroupNewResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupNewResponseIsDefaultAccessDevicePostureRule struct {
	DevicePosture AccessGroupNewResponseIsDefaultAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupNewResponseIsDefaultAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessDevicePostureRuleJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseIsDefaultAccessDevicePostureRule]
type accessGroupNewResponseIsDefaultAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseIsDefaultAccessDevicePostureRule) implementsAccessGroupNewResponseIsDefault() {
}

type AccessGroupNewResponseIsDefaultAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                  `json:"integration_uid,required"`
	JSON           accessGroupNewResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupNewResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON contains
// the JSON metadata for the struct
// [AccessGroupNewResponseIsDefaultAccessDevicePostureRuleDevicePosture]
type accessGroupNewResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupNewResponseIsDefaultAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupNewResponseRequireAccessEmailRule],
// [AccessGroupNewResponseRequireAccessEmailListRule],
// [AccessGroupNewResponseRequireAccessDomainRule],
// [AccessGroupNewResponseRequireAccessEveryoneRule],
// [AccessGroupNewResponseRequireAccessIPRule],
// [AccessGroupNewResponseRequireAccessIPListRule],
// [AccessGroupNewResponseRequireAccessCertificateRule],
// [AccessGroupNewResponseRequireAccessAccessGroupRule],
// [AccessGroupNewResponseRequireAccessAzureGroupRule],
// [AccessGroupNewResponseRequireAccessGitHubOrganizationRule],
// [AccessGroupNewResponseRequireAccessGsuiteGroupRule],
// [AccessGroupNewResponseRequireAccessOktaGroupRule],
// [AccessGroupNewResponseRequireAccessSamlGroupRule],
// [AccessGroupNewResponseRequireAccessServiceTokenRule],
// [AccessGroupNewResponseRequireAccessAnyValidServiceTokenRule],
// [AccessGroupNewResponseRequireAccessExternalEvaluationRule],
// [AccessGroupNewResponseRequireAccessCountryRule],
// [AccessGroupNewResponseRequireAccessAuthenticationMethodRule] or
// [AccessGroupNewResponseRequireAccessDevicePostureRule].
type AccessGroupNewResponseRequire interface {
	implementsAccessGroupNewResponseRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupNewResponseRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupNewResponseRequireAccessEmailRule struct {
	Email AccessGroupNewResponseRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupNewResponseRequireAccessEmailRuleJSON  `json:"-"`
}

// accessGroupNewResponseRequireAccessEmailRuleJSON contains the JSON metadata for
// the struct [AccessGroupNewResponseRequireAccessEmailRule]
type accessGroupNewResponseRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseRequireAccessEmailRule) implementsAccessGroupNewResponseRequire() {}

type AccessGroupNewResponseRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                `json:"email,required" format:"email"`
	JSON  accessGroupNewResponseRequireAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupNewResponseRequireAccessEmailRuleEmailJSON contains the JSON metadata
// for the struct [AccessGroupNewResponseRequireAccessEmailRuleEmail]
type accessGroupNewResponseRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupNewResponseRequireAccessEmailListRule struct {
	EmailList AccessGroupNewResponseRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupNewResponseRequireAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupNewResponseRequireAccessEmailListRuleJSON contains the JSON metadata
// for the struct [AccessGroupNewResponseRequireAccessEmailListRule]
type accessGroupNewResponseRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseRequireAccessEmailListRule) implementsAccessGroupNewResponseRequire() {}

type AccessGroupNewResponseRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                        `json:"id,required"`
	JSON accessGroupNewResponseRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupNewResponseRequireAccessEmailListRuleEmailListJSON contains the JSON
// metadata for the struct
// [AccessGroupNewResponseRequireAccessEmailListRuleEmailList]
type accessGroupNewResponseRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupNewResponseRequireAccessDomainRule struct {
	EmailDomain AccessGroupNewResponseRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupNewResponseRequireAccessDomainRuleJSON        `json:"-"`
}

// accessGroupNewResponseRequireAccessDomainRuleJSON contains the JSON metadata for
// the struct [AccessGroupNewResponseRequireAccessDomainRule]
type accessGroupNewResponseRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseRequireAccessDomainRule) implementsAccessGroupNewResponseRequire() {}

type AccessGroupNewResponseRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                       `json:"domain,required"`
	JSON   accessGroupNewResponseRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupNewResponseRequireAccessDomainRuleEmailDomainJSON contains the JSON
// metadata for the struct
// [AccessGroupNewResponseRequireAccessDomainRuleEmailDomain]
type accessGroupNewResponseRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupNewResponseRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                         `json:"everyone,required"`
	JSON     accessGroupNewResponseRequireAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupNewResponseRequireAccessEveryoneRuleJSON contains the JSON metadata
// for the struct [AccessGroupNewResponseRequireAccessEveryoneRule]
type accessGroupNewResponseRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseRequireAccessEveryoneRule) implementsAccessGroupNewResponseRequire() {}

// Matches an IP address block.
type AccessGroupNewResponseRequireAccessIPRule struct {
	IP   AccessGroupNewResponseRequireAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupNewResponseRequireAccessIPRuleJSON `json:"-"`
}

// accessGroupNewResponseRequireAccessIPRuleJSON contains the JSON metadata for the
// struct [AccessGroupNewResponseRequireAccessIPRule]
type accessGroupNewResponseRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseRequireAccessIPRule) implementsAccessGroupNewResponseRequire() {}

type AccessGroupNewResponseRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                          `json:"ip,required"`
	JSON accessGroupNewResponseRequireAccessIPRuleIPJSON `json:"-"`
}

// accessGroupNewResponseRequireAccessIPRuleIPJSON contains the JSON metadata for
// the struct [AccessGroupNewResponseRequireAccessIPRuleIP]
type accessGroupNewResponseRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupNewResponseRequireAccessIPListRule struct {
	IPList AccessGroupNewResponseRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupNewResponseRequireAccessIPListRuleJSON   `json:"-"`
}

// accessGroupNewResponseRequireAccessIPListRuleJSON contains the JSON metadata for
// the struct [AccessGroupNewResponseRequireAccessIPListRule]
type accessGroupNewResponseRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseRequireAccessIPListRule) implementsAccessGroupNewResponseRequire() {}

type AccessGroupNewResponseRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                  `json:"id,required"`
	JSON accessGroupNewResponseRequireAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupNewResponseRequireAccessIPListRuleIPListJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseRequireAccessIPListRuleIPList]
type accessGroupNewResponseRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupNewResponseRequireAccessCertificateRule struct {
	Certificate interface{}                                            `json:"certificate,required"`
	JSON        accessGroupNewResponseRequireAccessCertificateRuleJSON `json:"-"`
}

// accessGroupNewResponseRequireAccessCertificateRuleJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseRequireAccessCertificateRule]
type accessGroupNewResponseRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseRequireAccessCertificateRule) implementsAccessGroupNewResponseRequire() {
}

// Matches an Access group.
type AccessGroupNewResponseRequireAccessAccessGroupRule struct {
	Group AccessGroupNewResponseRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupNewResponseRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupNewResponseRequireAccessAccessGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseRequireAccessAccessGroupRule]
type accessGroupNewResponseRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseRequireAccessAccessGroupRule) implementsAccessGroupNewResponseRequire() {
}

type AccessGroupNewResponseRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                      `json:"id,required"`
	JSON accessGroupNewResponseRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupNewResponseRequireAccessAccessGroupRuleGroupJSON contains the JSON
// metadata for the struct
// [AccessGroupNewResponseRequireAccessAccessGroupRuleGroup]
type accessGroupNewResponseRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupNewResponseRequireAccessAzureGroupRule struct {
	AzureAd AccessGroupNewResponseRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupNewResponseRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupNewResponseRequireAccessAzureGroupRuleJSON contains the JSON metadata
// for the struct [AccessGroupNewResponseRequireAccessAzureGroupRule]
type accessGroupNewResponseRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseRequireAccessAzureGroupRule) implementsAccessGroupNewResponseRequire() {
}

type AccessGroupNewResponseRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                       `json:"connection_id,required"`
	JSON         accessGroupNewResponseRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupNewResponseRequireAccessAzureGroupRuleAzureAdJSON contains the JSON
// metadata for the struct
// [AccessGroupNewResponseRequireAccessAzureGroupRuleAzureAd]
type accessGroupNewResponseRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupNewResponseRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupNewResponseRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupNewResponseRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupNewResponseRequireAccessGitHubOrganizationRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupNewResponseRequireAccessGitHubOrganizationRule]
type accessGroupNewResponseRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseRequireAccessGitHubOrganizationRule) implementsAccessGroupNewResponseRequire() {
}

type AccessGroupNewResponseRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                          `json:"name,required"`
	JSON accessGroupNewResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupNewResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupNewResponseRequireAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupNewResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupNewResponseRequireAccessGsuiteGroupRule struct {
	Gsuite AccessGroupNewResponseRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupNewResponseRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupNewResponseRequireAccessGsuiteGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseRequireAccessGsuiteGroupRule]
type accessGroupNewResponseRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseRequireAccessGsuiteGroupRule) implementsAccessGroupNewResponseRequire() {
}

type AccessGroupNewResponseRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                       `json:"email,required"`
	JSON  accessGroupNewResponseRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupNewResponseRequireAccessGsuiteGroupRuleGsuiteJSON contains the JSON
// metadata for the struct
// [AccessGroupNewResponseRequireAccessGsuiteGroupRuleGsuite]
type accessGroupNewResponseRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupNewResponseRequireAccessOktaGroupRule struct {
	Okta AccessGroupNewResponseRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupNewResponseRequireAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupNewResponseRequireAccessOktaGroupRuleJSON contains the JSON metadata
// for the struct [AccessGroupNewResponseRequireAccessOktaGroupRule]
type accessGroupNewResponseRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseRequireAccessOktaGroupRule) implementsAccessGroupNewResponseRequire() {}

type AccessGroupNewResponseRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                   `json:"email,required"`
	JSON  accessGroupNewResponseRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupNewResponseRequireAccessOktaGroupRuleOktaJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseRequireAccessOktaGroupRuleOkta]
type accessGroupNewResponseRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupNewResponseRequireAccessSamlGroupRule struct {
	Saml AccessGroupNewResponseRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupNewResponseRequireAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupNewResponseRequireAccessSamlGroupRuleJSON contains the JSON metadata
// for the struct [AccessGroupNewResponseRequireAccessSamlGroupRule]
type accessGroupNewResponseRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseRequireAccessSamlGroupRule) implementsAccessGroupNewResponseRequire() {}

type AccessGroupNewResponseRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                   `json:"attribute_value,required"`
	JSON           accessGroupNewResponseRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupNewResponseRequireAccessSamlGroupRuleSamlJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseRequireAccessSamlGroupRuleSaml]
type accessGroupNewResponseRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupNewResponseRequireAccessServiceTokenRule struct {
	ServiceToken AccessGroupNewResponseRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupNewResponseRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupNewResponseRequireAccessServiceTokenRuleJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseRequireAccessServiceTokenRule]
type accessGroupNewResponseRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseRequireAccessServiceTokenRule) implementsAccessGroupNewResponseRequire() {
}

type AccessGroupNewResponseRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                              `json:"token_id,required"`
	JSON    accessGroupNewResponseRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupNewResponseRequireAccessServiceTokenRuleServiceTokenJSON contains the
// JSON metadata for the struct
// [AccessGroupNewResponseRequireAccessServiceTokenRuleServiceToken]
type accessGroupNewResponseRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupNewResponseRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                     `json:"any_valid_service_token,required"`
	JSON                 accessGroupNewResponseRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupNewResponseRequireAccessAnyValidServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupNewResponseRequireAccessAnyValidServiceTokenRule]
type accessGroupNewResponseRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseRequireAccessAnyValidServiceTokenRule) implementsAccessGroupNewResponseRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupNewResponseRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupNewResponseRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupNewResponseRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupNewResponseRequireAccessExternalEvaluationRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupNewResponseRequireAccessExternalEvaluationRule]
type accessGroupNewResponseRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseRequireAccessExternalEvaluationRule) implementsAccessGroupNewResponseRequire() {
}

type AccessGroupNewResponseRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                          `json:"keys_url,required"`
	JSON    accessGroupNewResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupNewResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupNewResponseRequireAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupNewResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupNewResponseRequireAccessCountryRule struct {
	Geo  AccessGroupNewResponseRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupNewResponseRequireAccessCountryRuleJSON `json:"-"`
}

// accessGroupNewResponseRequireAccessCountryRuleJSON contains the JSON metadata
// for the struct [AccessGroupNewResponseRequireAccessCountryRule]
type accessGroupNewResponseRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseRequireAccessCountryRule) implementsAccessGroupNewResponseRequire() {}

type AccessGroupNewResponseRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                `json:"country_code,required"`
	JSON        accessGroupNewResponseRequireAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupNewResponseRequireAccessCountryRuleGeoJSON contains the JSON metadata
// for the struct [AccessGroupNewResponseRequireAccessCountryRuleGeo]
type accessGroupNewResponseRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupNewResponseRequireAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupNewResponseRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupNewResponseRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupNewResponseRequireAccessAuthenticationMethodRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupNewResponseRequireAccessAuthenticationMethodRule]
type accessGroupNewResponseRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseRequireAccessAuthenticationMethodRule) implementsAccessGroupNewResponseRequire() {
}

type AccessGroupNewResponseRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                    `json:"auth_method,required"`
	JSON       accessGroupNewResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupNewResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupNewResponseRequireAccessAuthenticationMethodRuleAuthMethod]
type accessGroupNewResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupNewResponseRequireAccessDevicePostureRule struct {
	DevicePosture AccessGroupNewResponseRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupNewResponseRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupNewResponseRequireAccessDevicePostureRuleJSON contains the JSON
// metadata for the struct [AccessGroupNewResponseRequireAccessDevicePostureRule]
type accessGroupNewResponseRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupNewResponseRequireAccessDevicePostureRule) implementsAccessGroupNewResponseRequire() {
}

type AccessGroupNewResponseRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                `json:"integration_uid,required"`
	JSON           accessGroupNewResponseRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupNewResponseRequireAccessDevicePostureRuleDevicePostureJSON contains
// the JSON metadata for the struct
// [AccessGroupNewResponseRequireAccessDevicePostureRuleDevicePosture]
type accessGroupNewResponseRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupNewResponseRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupListResponse struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessGroupListResponseExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessGroupListResponseInclude `json:"include"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	IsDefault []AccessGroupListResponseIsDefault `json:"is_default"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []AccessGroupListResponseRequire `json:"require"`
	UpdatedAt time.Time                        `json:"updated_at" format:"date-time"`
	JSON      accessGroupListResponseJSON      `json:"-"`
}

// accessGroupListResponseJSON contains the JSON metadata for the struct
// [AccessGroupListResponse]
type accessGroupListResponseJSON struct {
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

func (r *AccessGroupListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupListResponseExcludeAccessEmailRule],
// [AccessGroupListResponseExcludeAccessEmailListRule],
// [AccessGroupListResponseExcludeAccessDomainRule],
// [AccessGroupListResponseExcludeAccessEveryoneRule],
// [AccessGroupListResponseExcludeAccessIPRule],
// [AccessGroupListResponseExcludeAccessIPListRule],
// [AccessGroupListResponseExcludeAccessCertificateRule],
// [AccessGroupListResponseExcludeAccessAccessGroupRule],
// [AccessGroupListResponseExcludeAccessAzureGroupRule],
// [AccessGroupListResponseExcludeAccessGitHubOrganizationRule],
// [AccessGroupListResponseExcludeAccessGsuiteGroupRule],
// [AccessGroupListResponseExcludeAccessOktaGroupRule],
// [AccessGroupListResponseExcludeAccessSamlGroupRule],
// [AccessGroupListResponseExcludeAccessServiceTokenRule],
// [AccessGroupListResponseExcludeAccessAnyValidServiceTokenRule],
// [AccessGroupListResponseExcludeAccessExternalEvaluationRule],
// [AccessGroupListResponseExcludeAccessCountryRule],
// [AccessGroupListResponseExcludeAccessAuthenticationMethodRule] or
// [AccessGroupListResponseExcludeAccessDevicePostureRule].
type AccessGroupListResponseExclude interface {
	implementsAccessGroupListResponseExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupListResponseExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupListResponseExcludeAccessEmailRule struct {
	Email AccessGroupListResponseExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupListResponseExcludeAccessEmailRuleJSON  `json:"-"`
}

// accessGroupListResponseExcludeAccessEmailRuleJSON contains the JSON metadata for
// the struct [AccessGroupListResponseExcludeAccessEmailRule]
type accessGroupListResponseExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseExcludeAccessEmailRule) implementsAccessGroupListResponseExclude() {}

type AccessGroupListResponseExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                 `json:"email,required" format:"email"`
	JSON  accessGroupListResponseExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupListResponseExcludeAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct [AccessGroupListResponseExcludeAccessEmailRuleEmail]
type accessGroupListResponseExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupListResponseExcludeAccessEmailListRule struct {
	EmailList AccessGroupListResponseExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupListResponseExcludeAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupListResponseExcludeAccessEmailListRuleJSON contains the JSON metadata
// for the struct [AccessGroupListResponseExcludeAccessEmailListRule]
type accessGroupListResponseExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseExcludeAccessEmailListRule) implementsAccessGroupListResponseExclude() {
}

type AccessGroupListResponseExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                         `json:"id,required"`
	JSON accessGroupListResponseExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupListResponseExcludeAccessEmailListRuleEmailListJSON contains the JSON
// metadata for the struct
// [AccessGroupListResponseExcludeAccessEmailListRuleEmailList]
type accessGroupListResponseExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupListResponseExcludeAccessDomainRule struct {
	EmailDomain AccessGroupListResponseExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupListResponseExcludeAccessDomainRuleJSON        `json:"-"`
}

// accessGroupListResponseExcludeAccessDomainRuleJSON contains the JSON metadata
// for the struct [AccessGroupListResponseExcludeAccessDomainRule]
type accessGroupListResponseExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseExcludeAccessDomainRule) implementsAccessGroupListResponseExclude() {}

type AccessGroupListResponseExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                        `json:"domain,required"`
	JSON   accessGroupListResponseExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupListResponseExcludeAccessDomainRuleEmailDomainJSON contains the JSON
// metadata for the struct
// [AccessGroupListResponseExcludeAccessDomainRuleEmailDomain]
type accessGroupListResponseExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupListResponseExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                          `json:"everyone,required"`
	JSON     accessGroupListResponseExcludeAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupListResponseExcludeAccessEveryoneRuleJSON contains the JSON metadata
// for the struct [AccessGroupListResponseExcludeAccessEveryoneRule]
type accessGroupListResponseExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseExcludeAccessEveryoneRule) implementsAccessGroupListResponseExclude() {
}

// Matches an IP address block.
type AccessGroupListResponseExcludeAccessIPRule struct {
	IP   AccessGroupListResponseExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupListResponseExcludeAccessIPRuleJSON `json:"-"`
}

// accessGroupListResponseExcludeAccessIPRuleJSON contains the JSON metadata for
// the struct [AccessGroupListResponseExcludeAccessIPRule]
type accessGroupListResponseExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseExcludeAccessIPRule) implementsAccessGroupListResponseExclude() {}

type AccessGroupListResponseExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                           `json:"ip,required"`
	JSON accessGroupListResponseExcludeAccessIPRuleIPJSON `json:"-"`
}

// accessGroupListResponseExcludeAccessIPRuleIPJSON contains the JSON metadata for
// the struct [AccessGroupListResponseExcludeAccessIPRuleIP]
type accessGroupListResponseExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupListResponseExcludeAccessIPListRule struct {
	IPList AccessGroupListResponseExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupListResponseExcludeAccessIPListRuleJSON   `json:"-"`
}

// accessGroupListResponseExcludeAccessIPListRuleJSON contains the JSON metadata
// for the struct [AccessGroupListResponseExcludeAccessIPListRule]
type accessGroupListResponseExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseExcludeAccessIPListRule) implementsAccessGroupListResponseExclude() {}

type AccessGroupListResponseExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                   `json:"id,required"`
	JSON accessGroupListResponseExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupListResponseExcludeAccessIPListRuleIPListJSON contains the JSON
// metadata for the struct [AccessGroupListResponseExcludeAccessIPListRuleIPList]
type accessGroupListResponseExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupListResponseExcludeAccessCertificateRule struct {
	Certificate interface{}                                             `json:"certificate,required"`
	JSON        accessGroupListResponseExcludeAccessCertificateRuleJSON `json:"-"`
}

// accessGroupListResponseExcludeAccessCertificateRuleJSON contains the JSON
// metadata for the struct [AccessGroupListResponseExcludeAccessCertificateRule]
type accessGroupListResponseExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseExcludeAccessCertificateRule) implementsAccessGroupListResponseExclude() {
}

// Matches an Access group.
type AccessGroupListResponseExcludeAccessAccessGroupRule struct {
	Group AccessGroupListResponseExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupListResponseExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupListResponseExcludeAccessAccessGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupListResponseExcludeAccessAccessGroupRule]
type accessGroupListResponseExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseExcludeAccessAccessGroupRule) implementsAccessGroupListResponseExclude() {
}

type AccessGroupListResponseExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                       `json:"id,required"`
	JSON accessGroupListResponseExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupListResponseExcludeAccessAccessGroupRuleGroupJSON contains the JSON
// metadata for the struct
// [AccessGroupListResponseExcludeAccessAccessGroupRuleGroup]
type accessGroupListResponseExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupListResponseExcludeAccessAzureGroupRule struct {
	AzureAd AccessGroupListResponseExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupListResponseExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupListResponseExcludeAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupListResponseExcludeAccessAzureGroupRule]
type accessGroupListResponseExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseExcludeAccessAzureGroupRule) implementsAccessGroupListResponseExclude() {
}

type AccessGroupListResponseExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                        `json:"connection_id,required"`
	JSON         accessGroupListResponseExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupListResponseExcludeAccessAzureGroupRuleAzureAdJSON contains the JSON
// metadata for the struct
// [AccessGroupListResponseExcludeAccessAzureGroupRuleAzureAd]
type accessGroupListResponseExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupListResponseExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupListResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupListResponseExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupListResponseExcludeAccessGitHubOrganizationRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupListResponseExcludeAccessGitHubOrganizationRule]
type accessGroupListResponseExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseExcludeAccessGitHubOrganizationRule) implementsAccessGroupListResponseExclude() {
}

type AccessGroupListResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                           `json:"name,required"`
	JSON accessGroupListResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupListResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupListResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupListResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupListResponseExcludeAccessGsuiteGroupRule struct {
	Gsuite AccessGroupListResponseExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupListResponseExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupListResponseExcludeAccessGsuiteGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupListResponseExcludeAccessGsuiteGroupRule]
type accessGroupListResponseExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseExcludeAccessGsuiteGroupRule) implementsAccessGroupListResponseExclude() {
}

type AccessGroupListResponseExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                        `json:"email,required"`
	JSON  accessGroupListResponseExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupListResponseExcludeAccessGsuiteGroupRuleGsuiteJSON contains the JSON
// metadata for the struct
// [AccessGroupListResponseExcludeAccessGsuiteGroupRuleGsuite]
type accessGroupListResponseExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupListResponseExcludeAccessOktaGroupRule struct {
	Okta AccessGroupListResponseExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupListResponseExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupListResponseExcludeAccessOktaGroupRuleJSON contains the JSON metadata
// for the struct [AccessGroupListResponseExcludeAccessOktaGroupRule]
type accessGroupListResponseExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseExcludeAccessOktaGroupRule) implementsAccessGroupListResponseExclude() {
}

type AccessGroupListResponseExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                    `json:"email,required"`
	JSON  accessGroupListResponseExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupListResponseExcludeAccessOktaGroupRuleOktaJSON contains the JSON
// metadata for the struct [AccessGroupListResponseExcludeAccessOktaGroupRuleOkta]
type accessGroupListResponseExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupListResponseExcludeAccessSamlGroupRule struct {
	Saml AccessGroupListResponseExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupListResponseExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupListResponseExcludeAccessSamlGroupRuleJSON contains the JSON metadata
// for the struct [AccessGroupListResponseExcludeAccessSamlGroupRule]
type accessGroupListResponseExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseExcludeAccessSamlGroupRule) implementsAccessGroupListResponseExclude() {
}

type AccessGroupListResponseExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                    `json:"attribute_value,required"`
	JSON           accessGroupListResponseExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupListResponseExcludeAccessSamlGroupRuleSamlJSON contains the JSON
// metadata for the struct [AccessGroupListResponseExcludeAccessSamlGroupRuleSaml]
type accessGroupListResponseExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupListResponseExcludeAccessServiceTokenRule struct {
	ServiceToken AccessGroupListResponseExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupListResponseExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupListResponseExcludeAccessServiceTokenRuleJSON contains the JSON
// metadata for the struct [AccessGroupListResponseExcludeAccessServiceTokenRule]
type accessGroupListResponseExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseExcludeAccessServiceTokenRule) implementsAccessGroupListResponseExclude() {
}

type AccessGroupListResponseExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                               `json:"token_id,required"`
	JSON    accessGroupListResponseExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupListResponseExcludeAccessServiceTokenRuleServiceTokenJSON contains
// the JSON metadata for the struct
// [AccessGroupListResponseExcludeAccessServiceTokenRuleServiceToken]
type accessGroupListResponseExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupListResponseExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                      `json:"any_valid_service_token,required"`
	JSON                 accessGroupListResponseExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupListResponseExcludeAccessAnyValidServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupListResponseExcludeAccessAnyValidServiceTokenRule]
type accessGroupListResponseExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseExcludeAccessAnyValidServiceTokenRule) implementsAccessGroupListResponseExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupListResponseExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupListResponseExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupListResponseExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupListResponseExcludeAccessExternalEvaluationRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupListResponseExcludeAccessExternalEvaluationRule]
type accessGroupListResponseExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseExcludeAccessExternalEvaluationRule) implementsAccessGroupListResponseExclude() {
}

type AccessGroupListResponseExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                           `json:"keys_url,required"`
	JSON    accessGroupListResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupListResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupListResponseExcludeAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupListResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupListResponseExcludeAccessCountryRule struct {
	Geo  AccessGroupListResponseExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupListResponseExcludeAccessCountryRuleJSON `json:"-"`
}

// accessGroupListResponseExcludeAccessCountryRuleJSON contains the JSON metadata
// for the struct [AccessGroupListResponseExcludeAccessCountryRule]
type accessGroupListResponseExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseExcludeAccessCountryRule) implementsAccessGroupListResponseExclude() {}

type AccessGroupListResponseExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                 `json:"country_code,required"`
	JSON        accessGroupListResponseExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupListResponseExcludeAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct [AccessGroupListResponseExcludeAccessCountryRuleGeo]
type accessGroupListResponseExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupListResponseExcludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupListResponseExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupListResponseExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupListResponseExcludeAccessAuthenticationMethodRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupListResponseExcludeAccessAuthenticationMethodRule]
type accessGroupListResponseExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseExcludeAccessAuthenticationMethodRule) implementsAccessGroupListResponseExclude() {
}

type AccessGroupListResponseExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                     `json:"auth_method,required"`
	JSON       accessGroupListResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupListResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupListResponseExcludeAccessAuthenticationMethodRuleAuthMethod]
type accessGroupListResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupListResponseExcludeAccessDevicePostureRule struct {
	DevicePosture AccessGroupListResponseExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupListResponseExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupListResponseExcludeAccessDevicePostureRuleJSON contains the JSON
// metadata for the struct [AccessGroupListResponseExcludeAccessDevicePostureRule]
type accessGroupListResponseExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseExcludeAccessDevicePostureRule) implementsAccessGroupListResponseExclude() {
}

type AccessGroupListResponseExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                 `json:"integration_uid,required"`
	JSON           accessGroupListResponseExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupListResponseExcludeAccessDevicePostureRuleDevicePostureJSON contains
// the JSON metadata for the struct
// [AccessGroupListResponseExcludeAccessDevicePostureRuleDevicePosture]
type accessGroupListResponseExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupListResponseExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupListResponseIncludeAccessEmailRule],
// [AccessGroupListResponseIncludeAccessEmailListRule],
// [AccessGroupListResponseIncludeAccessDomainRule],
// [AccessGroupListResponseIncludeAccessEveryoneRule],
// [AccessGroupListResponseIncludeAccessIPRule],
// [AccessGroupListResponseIncludeAccessIPListRule],
// [AccessGroupListResponseIncludeAccessCertificateRule],
// [AccessGroupListResponseIncludeAccessAccessGroupRule],
// [AccessGroupListResponseIncludeAccessAzureGroupRule],
// [AccessGroupListResponseIncludeAccessGitHubOrganizationRule],
// [AccessGroupListResponseIncludeAccessGsuiteGroupRule],
// [AccessGroupListResponseIncludeAccessOktaGroupRule],
// [AccessGroupListResponseIncludeAccessSamlGroupRule],
// [AccessGroupListResponseIncludeAccessServiceTokenRule],
// [AccessGroupListResponseIncludeAccessAnyValidServiceTokenRule],
// [AccessGroupListResponseIncludeAccessExternalEvaluationRule],
// [AccessGroupListResponseIncludeAccessCountryRule],
// [AccessGroupListResponseIncludeAccessAuthenticationMethodRule] or
// [AccessGroupListResponseIncludeAccessDevicePostureRule].
type AccessGroupListResponseInclude interface {
	implementsAccessGroupListResponseInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupListResponseInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupListResponseIncludeAccessEmailRule struct {
	Email AccessGroupListResponseIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupListResponseIncludeAccessEmailRuleJSON  `json:"-"`
}

// accessGroupListResponseIncludeAccessEmailRuleJSON contains the JSON metadata for
// the struct [AccessGroupListResponseIncludeAccessEmailRule]
type accessGroupListResponseIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIncludeAccessEmailRule) implementsAccessGroupListResponseInclude() {}

type AccessGroupListResponseIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                 `json:"email,required" format:"email"`
	JSON  accessGroupListResponseIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupListResponseIncludeAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct [AccessGroupListResponseIncludeAccessEmailRuleEmail]
type accessGroupListResponseIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupListResponseIncludeAccessEmailListRule struct {
	EmailList AccessGroupListResponseIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupListResponseIncludeAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupListResponseIncludeAccessEmailListRuleJSON contains the JSON metadata
// for the struct [AccessGroupListResponseIncludeAccessEmailListRule]
type accessGroupListResponseIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIncludeAccessEmailListRule) implementsAccessGroupListResponseInclude() {
}

type AccessGroupListResponseIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                         `json:"id,required"`
	JSON accessGroupListResponseIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupListResponseIncludeAccessEmailListRuleEmailListJSON contains the JSON
// metadata for the struct
// [AccessGroupListResponseIncludeAccessEmailListRuleEmailList]
type accessGroupListResponseIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupListResponseIncludeAccessDomainRule struct {
	EmailDomain AccessGroupListResponseIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupListResponseIncludeAccessDomainRuleJSON        `json:"-"`
}

// accessGroupListResponseIncludeAccessDomainRuleJSON contains the JSON metadata
// for the struct [AccessGroupListResponseIncludeAccessDomainRule]
type accessGroupListResponseIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIncludeAccessDomainRule) implementsAccessGroupListResponseInclude() {}

type AccessGroupListResponseIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                        `json:"domain,required"`
	JSON   accessGroupListResponseIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupListResponseIncludeAccessDomainRuleEmailDomainJSON contains the JSON
// metadata for the struct
// [AccessGroupListResponseIncludeAccessDomainRuleEmailDomain]
type accessGroupListResponseIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupListResponseIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                          `json:"everyone,required"`
	JSON     accessGroupListResponseIncludeAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupListResponseIncludeAccessEveryoneRuleJSON contains the JSON metadata
// for the struct [AccessGroupListResponseIncludeAccessEveryoneRule]
type accessGroupListResponseIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIncludeAccessEveryoneRule) implementsAccessGroupListResponseInclude() {
}

// Matches an IP address block.
type AccessGroupListResponseIncludeAccessIPRule struct {
	IP   AccessGroupListResponseIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupListResponseIncludeAccessIPRuleJSON `json:"-"`
}

// accessGroupListResponseIncludeAccessIPRuleJSON contains the JSON metadata for
// the struct [AccessGroupListResponseIncludeAccessIPRule]
type accessGroupListResponseIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIncludeAccessIPRule) implementsAccessGroupListResponseInclude() {}

type AccessGroupListResponseIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                           `json:"ip,required"`
	JSON accessGroupListResponseIncludeAccessIPRuleIPJSON `json:"-"`
}

// accessGroupListResponseIncludeAccessIPRuleIPJSON contains the JSON metadata for
// the struct [AccessGroupListResponseIncludeAccessIPRuleIP]
type accessGroupListResponseIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupListResponseIncludeAccessIPListRule struct {
	IPList AccessGroupListResponseIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupListResponseIncludeAccessIPListRuleJSON   `json:"-"`
}

// accessGroupListResponseIncludeAccessIPListRuleJSON contains the JSON metadata
// for the struct [AccessGroupListResponseIncludeAccessIPListRule]
type accessGroupListResponseIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIncludeAccessIPListRule) implementsAccessGroupListResponseInclude() {}

type AccessGroupListResponseIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                   `json:"id,required"`
	JSON accessGroupListResponseIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupListResponseIncludeAccessIPListRuleIPListJSON contains the JSON
// metadata for the struct [AccessGroupListResponseIncludeAccessIPListRuleIPList]
type accessGroupListResponseIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupListResponseIncludeAccessCertificateRule struct {
	Certificate interface{}                                             `json:"certificate,required"`
	JSON        accessGroupListResponseIncludeAccessCertificateRuleJSON `json:"-"`
}

// accessGroupListResponseIncludeAccessCertificateRuleJSON contains the JSON
// metadata for the struct [AccessGroupListResponseIncludeAccessCertificateRule]
type accessGroupListResponseIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIncludeAccessCertificateRule) implementsAccessGroupListResponseInclude() {
}

// Matches an Access group.
type AccessGroupListResponseIncludeAccessAccessGroupRule struct {
	Group AccessGroupListResponseIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupListResponseIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupListResponseIncludeAccessAccessGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupListResponseIncludeAccessAccessGroupRule]
type accessGroupListResponseIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIncludeAccessAccessGroupRule) implementsAccessGroupListResponseInclude() {
}

type AccessGroupListResponseIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                       `json:"id,required"`
	JSON accessGroupListResponseIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupListResponseIncludeAccessAccessGroupRuleGroupJSON contains the JSON
// metadata for the struct
// [AccessGroupListResponseIncludeAccessAccessGroupRuleGroup]
type accessGroupListResponseIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupListResponseIncludeAccessAzureGroupRule struct {
	AzureAd AccessGroupListResponseIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupListResponseIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupListResponseIncludeAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupListResponseIncludeAccessAzureGroupRule]
type accessGroupListResponseIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIncludeAccessAzureGroupRule) implementsAccessGroupListResponseInclude() {
}

type AccessGroupListResponseIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                        `json:"connection_id,required"`
	JSON         accessGroupListResponseIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupListResponseIncludeAccessAzureGroupRuleAzureAdJSON contains the JSON
// metadata for the struct
// [AccessGroupListResponseIncludeAccessAzureGroupRuleAzureAd]
type accessGroupListResponseIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupListResponseIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupListResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupListResponseIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupListResponseIncludeAccessGitHubOrganizationRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupListResponseIncludeAccessGitHubOrganizationRule]
type accessGroupListResponseIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIncludeAccessGitHubOrganizationRule) implementsAccessGroupListResponseInclude() {
}

type AccessGroupListResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                           `json:"name,required"`
	JSON accessGroupListResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupListResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupListResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupListResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupListResponseIncludeAccessGsuiteGroupRule struct {
	Gsuite AccessGroupListResponseIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupListResponseIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupListResponseIncludeAccessGsuiteGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupListResponseIncludeAccessGsuiteGroupRule]
type accessGroupListResponseIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIncludeAccessGsuiteGroupRule) implementsAccessGroupListResponseInclude() {
}

type AccessGroupListResponseIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                        `json:"email,required"`
	JSON  accessGroupListResponseIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupListResponseIncludeAccessGsuiteGroupRuleGsuiteJSON contains the JSON
// metadata for the struct
// [AccessGroupListResponseIncludeAccessGsuiteGroupRuleGsuite]
type accessGroupListResponseIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupListResponseIncludeAccessOktaGroupRule struct {
	Okta AccessGroupListResponseIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupListResponseIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupListResponseIncludeAccessOktaGroupRuleJSON contains the JSON metadata
// for the struct [AccessGroupListResponseIncludeAccessOktaGroupRule]
type accessGroupListResponseIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIncludeAccessOktaGroupRule) implementsAccessGroupListResponseInclude() {
}

type AccessGroupListResponseIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                    `json:"email,required"`
	JSON  accessGroupListResponseIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupListResponseIncludeAccessOktaGroupRuleOktaJSON contains the JSON
// metadata for the struct [AccessGroupListResponseIncludeAccessOktaGroupRuleOkta]
type accessGroupListResponseIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupListResponseIncludeAccessSamlGroupRule struct {
	Saml AccessGroupListResponseIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupListResponseIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupListResponseIncludeAccessSamlGroupRuleJSON contains the JSON metadata
// for the struct [AccessGroupListResponseIncludeAccessSamlGroupRule]
type accessGroupListResponseIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIncludeAccessSamlGroupRule) implementsAccessGroupListResponseInclude() {
}

type AccessGroupListResponseIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                    `json:"attribute_value,required"`
	JSON           accessGroupListResponseIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupListResponseIncludeAccessSamlGroupRuleSamlJSON contains the JSON
// metadata for the struct [AccessGroupListResponseIncludeAccessSamlGroupRuleSaml]
type accessGroupListResponseIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupListResponseIncludeAccessServiceTokenRule struct {
	ServiceToken AccessGroupListResponseIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupListResponseIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupListResponseIncludeAccessServiceTokenRuleJSON contains the JSON
// metadata for the struct [AccessGroupListResponseIncludeAccessServiceTokenRule]
type accessGroupListResponseIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIncludeAccessServiceTokenRule) implementsAccessGroupListResponseInclude() {
}

type AccessGroupListResponseIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                               `json:"token_id,required"`
	JSON    accessGroupListResponseIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupListResponseIncludeAccessServiceTokenRuleServiceTokenJSON contains
// the JSON metadata for the struct
// [AccessGroupListResponseIncludeAccessServiceTokenRuleServiceToken]
type accessGroupListResponseIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupListResponseIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                      `json:"any_valid_service_token,required"`
	JSON                 accessGroupListResponseIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupListResponseIncludeAccessAnyValidServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupListResponseIncludeAccessAnyValidServiceTokenRule]
type accessGroupListResponseIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIncludeAccessAnyValidServiceTokenRule) implementsAccessGroupListResponseInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupListResponseIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupListResponseIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupListResponseIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupListResponseIncludeAccessExternalEvaluationRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupListResponseIncludeAccessExternalEvaluationRule]
type accessGroupListResponseIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIncludeAccessExternalEvaluationRule) implementsAccessGroupListResponseInclude() {
}

type AccessGroupListResponseIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                           `json:"keys_url,required"`
	JSON    accessGroupListResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupListResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupListResponseIncludeAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupListResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupListResponseIncludeAccessCountryRule struct {
	Geo  AccessGroupListResponseIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupListResponseIncludeAccessCountryRuleJSON `json:"-"`
}

// accessGroupListResponseIncludeAccessCountryRuleJSON contains the JSON metadata
// for the struct [AccessGroupListResponseIncludeAccessCountryRule]
type accessGroupListResponseIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIncludeAccessCountryRule) implementsAccessGroupListResponseInclude() {}

type AccessGroupListResponseIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                 `json:"country_code,required"`
	JSON        accessGroupListResponseIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupListResponseIncludeAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct [AccessGroupListResponseIncludeAccessCountryRuleGeo]
type accessGroupListResponseIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupListResponseIncludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupListResponseIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupListResponseIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupListResponseIncludeAccessAuthenticationMethodRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupListResponseIncludeAccessAuthenticationMethodRule]
type accessGroupListResponseIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIncludeAccessAuthenticationMethodRule) implementsAccessGroupListResponseInclude() {
}

type AccessGroupListResponseIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                     `json:"auth_method,required"`
	JSON       accessGroupListResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupListResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupListResponseIncludeAccessAuthenticationMethodRuleAuthMethod]
type accessGroupListResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupListResponseIncludeAccessDevicePostureRule struct {
	DevicePosture AccessGroupListResponseIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupListResponseIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupListResponseIncludeAccessDevicePostureRuleJSON contains the JSON
// metadata for the struct [AccessGroupListResponseIncludeAccessDevicePostureRule]
type accessGroupListResponseIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIncludeAccessDevicePostureRule) implementsAccessGroupListResponseInclude() {
}

type AccessGroupListResponseIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                 `json:"integration_uid,required"`
	JSON           accessGroupListResponseIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupListResponseIncludeAccessDevicePostureRuleDevicePostureJSON contains
// the JSON metadata for the struct
// [AccessGroupListResponseIncludeAccessDevicePostureRuleDevicePosture]
type accessGroupListResponseIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupListResponseIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupListResponseIsDefaultAccessEmailRule],
// [AccessGroupListResponseIsDefaultAccessEmailListRule],
// [AccessGroupListResponseIsDefaultAccessDomainRule],
// [AccessGroupListResponseIsDefaultAccessEveryoneRule],
// [AccessGroupListResponseIsDefaultAccessIPRule],
// [AccessGroupListResponseIsDefaultAccessIPListRule],
// [AccessGroupListResponseIsDefaultAccessCertificateRule],
// [AccessGroupListResponseIsDefaultAccessAccessGroupRule],
// [AccessGroupListResponseIsDefaultAccessAzureGroupRule],
// [AccessGroupListResponseIsDefaultAccessGitHubOrganizationRule],
// [AccessGroupListResponseIsDefaultAccessGsuiteGroupRule],
// [AccessGroupListResponseIsDefaultAccessOktaGroupRule],
// [AccessGroupListResponseIsDefaultAccessSamlGroupRule],
// [AccessGroupListResponseIsDefaultAccessServiceTokenRule],
// [AccessGroupListResponseIsDefaultAccessAnyValidServiceTokenRule],
// [AccessGroupListResponseIsDefaultAccessExternalEvaluationRule],
// [AccessGroupListResponseIsDefaultAccessCountryRule],
// [AccessGroupListResponseIsDefaultAccessAuthenticationMethodRule] or
// [AccessGroupListResponseIsDefaultAccessDevicePostureRule].
type AccessGroupListResponseIsDefault interface {
	implementsAccessGroupListResponseIsDefault()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupListResponseIsDefault)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupListResponseIsDefaultAccessEmailRule struct {
	Email AccessGroupListResponseIsDefaultAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupListResponseIsDefaultAccessEmailRuleJSON  `json:"-"`
}

// accessGroupListResponseIsDefaultAccessEmailRuleJSON contains the JSON metadata
// for the struct [AccessGroupListResponseIsDefaultAccessEmailRule]
type accessGroupListResponseIsDefaultAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIsDefaultAccessEmailRule) implementsAccessGroupListResponseIsDefault() {
}

type AccessGroupListResponseIsDefaultAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                   `json:"email,required" format:"email"`
	JSON  accessGroupListResponseIsDefaultAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupListResponseIsDefaultAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct [AccessGroupListResponseIsDefaultAccessEmailRuleEmail]
type accessGroupListResponseIsDefaultAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupListResponseIsDefaultAccessEmailListRule struct {
	EmailList AccessGroupListResponseIsDefaultAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupListResponseIsDefaultAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupListResponseIsDefaultAccessEmailListRuleJSON contains the JSON
// metadata for the struct [AccessGroupListResponseIsDefaultAccessEmailListRule]
type accessGroupListResponseIsDefaultAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIsDefaultAccessEmailListRule) implementsAccessGroupListResponseIsDefault() {
}

type AccessGroupListResponseIsDefaultAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                           `json:"id,required"`
	JSON accessGroupListResponseIsDefaultAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupListResponseIsDefaultAccessEmailListRuleEmailListJSON contains the
// JSON metadata for the struct
// [AccessGroupListResponseIsDefaultAccessEmailListRuleEmailList]
type accessGroupListResponseIsDefaultAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupListResponseIsDefaultAccessDomainRule struct {
	EmailDomain AccessGroupListResponseIsDefaultAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupListResponseIsDefaultAccessDomainRuleJSON        `json:"-"`
}

// accessGroupListResponseIsDefaultAccessDomainRuleJSON contains the JSON metadata
// for the struct [AccessGroupListResponseIsDefaultAccessDomainRule]
type accessGroupListResponseIsDefaultAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIsDefaultAccessDomainRule) implementsAccessGroupListResponseIsDefault() {
}

type AccessGroupListResponseIsDefaultAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                          `json:"domain,required"`
	JSON   accessGroupListResponseIsDefaultAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupListResponseIsDefaultAccessDomainRuleEmailDomainJSON contains the
// JSON metadata for the struct
// [AccessGroupListResponseIsDefaultAccessDomainRuleEmailDomain]
type accessGroupListResponseIsDefaultAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupListResponseIsDefaultAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                            `json:"everyone,required"`
	JSON     accessGroupListResponseIsDefaultAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupListResponseIsDefaultAccessEveryoneRuleJSON contains the JSON
// metadata for the struct [AccessGroupListResponseIsDefaultAccessEveryoneRule]
type accessGroupListResponseIsDefaultAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIsDefaultAccessEveryoneRule) implementsAccessGroupListResponseIsDefault() {
}

// Matches an IP address block.
type AccessGroupListResponseIsDefaultAccessIPRule struct {
	IP   AccessGroupListResponseIsDefaultAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupListResponseIsDefaultAccessIPRuleJSON `json:"-"`
}

// accessGroupListResponseIsDefaultAccessIPRuleJSON contains the JSON metadata for
// the struct [AccessGroupListResponseIsDefaultAccessIPRule]
type accessGroupListResponseIsDefaultAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIsDefaultAccessIPRule) implementsAccessGroupListResponseIsDefault() {}

type AccessGroupListResponseIsDefaultAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                             `json:"ip,required"`
	JSON accessGroupListResponseIsDefaultAccessIPRuleIPJSON `json:"-"`
}

// accessGroupListResponseIsDefaultAccessIPRuleIPJSON contains the JSON metadata
// for the struct [AccessGroupListResponseIsDefaultAccessIPRuleIP]
type accessGroupListResponseIsDefaultAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupListResponseIsDefaultAccessIPListRule struct {
	IPList AccessGroupListResponseIsDefaultAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupListResponseIsDefaultAccessIPListRuleJSON   `json:"-"`
}

// accessGroupListResponseIsDefaultAccessIPListRuleJSON contains the JSON metadata
// for the struct [AccessGroupListResponseIsDefaultAccessIPListRule]
type accessGroupListResponseIsDefaultAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIsDefaultAccessIPListRule) implementsAccessGroupListResponseIsDefault() {
}

type AccessGroupListResponseIsDefaultAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                     `json:"id,required"`
	JSON accessGroupListResponseIsDefaultAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupListResponseIsDefaultAccessIPListRuleIPListJSON contains the JSON
// metadata for the struct [AccessGroupListResponseIsDefaultAccessIPListRuleIPList]
type accessGroupListResponseIsDefaultAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupListResponseIsDefaultAccessCertificateRule struct {
	Certificate interface{}                                               `json:"certificate,required"`
	JSON        accessGroupListResponseIsDefaultAccessCertificateRuleJSON `json:"-"`
}

// accessGroupListResponseIsDefaultAccessCertificateRuleJSON contains the JSON
// metadata for the struct [AccessGroupListResponseIsDefaultAccessCertificateRule]
type accessGroupListResponseIsDefaultAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIsDefaultAccessCertificateRule) implementsAccessGroupListResponseIsDefault() {
}

// Matches an Access group.
type AccessGroupListResponseIsDefaultAccessAccessGroupRule struct {
	Group AccessGroupListResponseIsDefaultAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupListResponseIsDefaultAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupListResponseIsDefaultAccessAccessGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupListResponseIsDefaultAccessAccessGroupRule]
type accessGroupListResponseIsDefaultAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIsDefaultAccessAccessGroupRule) implementsAccessGroupListResponseIsDefault() {
}

type AccessGroupListResponseIsDefaultAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                         `json:"id,required"`
	JSON accessGroupListResponseIsDefaultAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupListResponseIsDefaultAccessAccessGroupRuleGroupJSON contains the JSON
// metadata for the struct
// [AccessGroupListResponseIsDefaultAccessAccessGroupRuleGroup]
type accessGroupListResponseIsDefaultAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupListResponseIsDefaultAccessAzureGroupRule struct {
	AzureAd AccessGroupListResponseIsDefaultAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupListResponseIsDefaultAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupListResponseIsDefaultAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupListResponseIsDefaultAccessAzureGroupRule]
type accessGroupListResponseIsDefaultAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIsDefaultAccessAzureGroupRule) implementsAccessGroupListResponseIsDefault() {
}

type AccessGroupListResponseIsDefaultAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                          `json:"connection_id,required"`
	JSON         accessGroupListResponseIsDefaultAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupListResponseIsDefaultAccessAzureGroupRuleAzureAdJSON contains the
// JSON metadata for the struct
// [AccessGroupListResponseIsDefaultAccessAzureGroupRuleAzureAd]
type accessGroupListResponseIsDefaultAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupListResponseIsDefaultAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupListResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupListResponseIsDefaultAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupListResponseIsDefaultAccessGitHubOrganizationRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupListResponseIsDefaultAccessGitHubOrganizationRule]
type accessGroupListResponseIsDefaultAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIsDefaultAccessGitHubOrganizationRule) implementsAccessGroupListResponseIsDefault() {
}

type AccessGroupListResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                             `json:"name,required"`
	JSON accessGroupListResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupListResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupListResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupListResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupListResponseIsDefaultAccessGsuiteGroupRule struct {
	Gsuite AccessGroupListResponseIsDefaultAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupListResponseIsDefaultAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupListResponseIsDefaultAccessGsuiteGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupListResponseIsDefaultAccessGsuiteGroupRule]
type accessGroupListResponseIsDefaultAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIsDefaultAccessGsuiteGroupRule) implementsAccessGroupListResponseIsDefault() {
}

type AccessGroupListResponseIsDefaultAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                          `json:"email,required"`
	JSON  accessGroupListResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupListResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON contains the
// JSON metadata for the struct
// [AccessGroupListResponseIsDefaultAccessGsuiteGroupRuleGsuite]
type accessGroupListResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupListResponseIsDefaultAccessOktaGroupRule struct {
	Okta AccessGroupListResponseIsDefaultAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupListResponseIsDefaultAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupListResponseIsDefaultAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupListResponseIsDefaultAccessOktaGroupRule]
type accessGroupListResponseIsDefaultAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIsDefaultAccessOktaGroupRule) implementsAccessGroupListResponseIsDefault() {
}

type AccessGroupListResponseIsDefaultAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                      `json:"email,required"`
	JSON  accessGroupListResponseIsDefaultAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupListResponseIsDefaultAccessOktaGroupRuleOktaJSON contains the JSON
// metadata for the struct
// [AccessGroupListResponseIsDefaultAccessOktaGroupRuleOkta]
type accessGroupListResponseIsDefaultAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupListResponseIsDefaultAccessSamlGroupRule struct {
	Saml AccessGroupListResponseIsDefaultAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupListResponseIsDefaultAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupListResponseIsDefaultAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupListResponseIsDefaultAccessSamlGroupRule]
type accessGroupListResponseIsDefaultAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIsDefaultAccessSamlGroupRule) implementsAccessGroupListResponseIsDefault() {
}

type AccessGroupListResponseIsDefaultAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                      `json:"attribute_value,required"`
	JSON           accessGroupListResponseIsDefaultAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupListResponseIsDefaultAccessSamlGroupRuleSamlJSON contains the JSON
// metadata for the struct
// [AccessGroupListResponseIsDefaultAccessSamlGroupRuleSaml]
type accessGroupListResponseIsDefaultAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupListResponseIsDefaultAccessServiceTokenRule struct {
	ServiceToken AccessGroupListResponseIsDefaultAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupListResponseIsDefaultAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupListResponseIsDefaultAccessServiceTokenRuleJSON contains the JSON
// metadata for the struct [AccessGroupListResponseIsDefaultAccessServiceTokenRule]
type accessGroupListResponseIsDefaultAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIsDefaultAccessServiceTokenRule) implementsAccessGroupListResponseIsDefault() {
}

type AccessGroupListResponseIsDefaultAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                 `json:"token_id,required"`
	JSON    accessGroupListResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupListResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON contains
// the JSON metadata for the struct
// [AccessGroupListResponseIsDefaultAccessServiceTokenRuleServiceToken]
type accessGroupListResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupListResponseIsDefaultAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                        `json:"any_valid_service_token,required"`
	JSON                 accessGroupListResponseIsDefaultAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupListResponseIsDefaultAccessAnyValidServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupListResponseIsDefaultAccessAnyValidServiceTokenRule]
type accessGroupListResponseIsDefaultAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIsDefaultAccessAnyValidServiceTokenRule) implementsAccessGroupListResponseIsDefault() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupListResponseIsDefaultAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupListResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupListResponseIsDefaultAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupListResponseIsDefaultAccessExternalEvaluationRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupListResponseIsDefaultAccessExternalEvaluationRule]
type accessGroupListResponseIsDefaultAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIsDefaultAccessExternalEvaluationRule) implementsAccessGroupListResponseIsDefault() {
}

type AccessGroupListResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                             `json:"keys_url,required"`
	JSON    accessGroupListResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupListResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupListResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupListResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupListResponseIsDefaultAccessCountryRule struct {
	Geo  AccessGroupListResponseIsDefaultAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupListResponseIsDefaultAccessCountryRuleJSON `json:"-"`
}

// accessGroupListResponseIsDefaultAccessCountryRuleJSON contains the JSON metadata
// for the struct [AccessGroupListResponseIsDefaultAccessCountryRule]
type accessGroupListResponseIsDefaultAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIsDefaultAccessCountryRule) implementsAccessGroupListResponseIsDefault() {
}

type AccessGroupListResponseIsDefaultAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                   `json:"country_code,required"`
	JSON        accessGroupListResponseIsDefaultAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupListResponseIsDefaultAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct [AccessGroupListResponseIsDefaultAccessCountryRuleGeo]
type accessGroupListResponseIsDefaultAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupListResponseIsDefaultAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupListResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupListResponseIsDefaultAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupListResponseIsDefaultAccessAuthenticationMethodRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupListResponseIsDefaultAccessAuthenticationMethodRule]
type accessGroupListResponseIsDefaultAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIsDefaultAccessAuthenticationMethodRule) implementsAccessGroupListResponseIsDefault() {
}

type AccessGroupListResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                       `json:"auth_method,required"`
	JSON       accessGroupListResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupListResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupListResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod]
type accessGroupListResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupListResponseIsDefaultAccessDevicePostureRule struct {
	DevicePosture AccessGroupListResponseIsDefaultAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupListResponseIsDefaultAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupListResponseIsDefaultAccessDevicePostureRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupListResponseIsDefaultAccessDevicePostureRule]
type accessGroupListResponseIsDefaultAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseIsDefaultAccessDevicePostureRule) implementsAccessGroupListResponseIsDefault() {
}

type AccessGroupListResponseIsDefaultAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                   `json:"integration_uid,required"`
	JSON           accessGroupListResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupListResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupListResponseIsDefaultAccessDevicePostureRuleDevicePosture]
type accessGroupListResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupListResponseIsDefaultAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupListResponseRequireAccessEmailRule],
// [AccessGroupListResponseRequireAccessEmailListRule],
// [AccessGroupListResponseRequireAccessDomainRule],
// [AccessGroupListResponseRequireAccessEveryoneRule],
// [AccessGroupListResponseRequireAccessIPRule],
// [AccessGroupListResponseRequireAccessIPListRule],
// [AccessGroupListResponseRequireAccessCertificateRule],
// [AccessGroupListResponseRequireAccessAccessGroupRule],
// [AccessGroupListResponseRequireAccessAzureGroupRule],
// [AccessGroupListResponseRequireAccessGitHubOrganizationRule],
// [AccessGroupListResponseRequireAccessGsuiteGroupRule],
// [AccessGroupListResponseRequireAccessOktaGroupRule],
// [AccessGroupListResponseRequireAccessSamlGroupRule],
// [AccessGroupListResponseRequireAccessServiceTokenRule],
// [AccessGroupListResponseRequireAccessAnyValidServiceTokenRule],
// [AccessGroupListResponseRequireAccessExternalEvaluationRule],
// [AccessGroupListResponseRequireAccessCountryRule],
// [AccessGroupListResponseRequireAccessAuthenticationMethodRule] or
// [AccessGroupListResponseRequireAccessDevicePostureRule].
type AccessGroupListResponseRequire interface {
	implementsAccessGroupListResponseRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupListResponseRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupListResponseRequireAccessEmailRule struct {
	Email AccessGroupListResponseRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupListResponseRequireAccessEmailRuleJSON  `json:"-"`
}

// accessGroupListResponseRequireAccessEmailRuleJSON contains the JSON metadata for
// the struct [AccessGroupListResponseRequireAccessEmailRule]
type accessGroupListResponseRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseRequireAccessEmailRule) implementsAccessGroupListResponseRequire() {}

type AccessGroupListResponseRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                 `json:"email,required" format:"email"`
	JSON  accessGroupListResponseRequireAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupListResponseRequireAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct [AccessGroupListResponseRequireAccessEmailRuleEmail]
type accessGroupListResponseRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupListResponseRequireAccessEmailListRule struct {
	EmailList AccessGroupListResponseRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupListResponseRequireAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupListResponseRequireAccessEmailListRuleJSON contains the JSON metadata
// for the struct [AccessGroupListResponseRequireAccessEmailListRule]
type accessGroupListResponseRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseRequireAccessEmailListRule) implementsAccessGroupListResponseRequire() {
}

type AccessGroupListResponseRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                         `json:"id,required"`
	JSON accessGroupListResponseRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupListResponseRequireAccessEmailListRuleEmailListJSON contains the JSON
// metadata for the struct
// [AccessGroupListResponseRequireAccessEmailListRuleEmailList]
type accessGroupListResponseRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupListResponseRequireAccessDomainRule struct {
	EmailDomain AccessGroupListResponseRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupListResponseRequireAccessDomainRuleJSON        `json:"-"`
}

// accessGroupListResponseRequireAccessDomainRuleJSON contains the JSON metadata
// for the struct [AccessGroupListResponseRequireAccessDomainRule]
type accessGroupListResponseRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseRequireAccessDomainRule) implementsAccessGroupListResponseRequire() {}

type AccessGroupListResponseRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                        `json:"domain,required"`
	JSON   accessGroupListResponseRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupListResponseRequireAccessDomainRuleEmailDomainJSON contains the JSON
// metadata for the struct
// [AccessGroupListResponseRequireAccessDomainRuleEmailDomain]
type accessGroupListResponseRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupListResponseRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                          `json:"everyone,required"`
	JSON     accessGroupListResponseRequireAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupListResponseRequireAccessEveryoneRuleJSON contains the JSON metadata
// for the struct [AccessGroupListResponseRequireAccessEveryoneRule]
type accessGroupListResponseRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseRequireAccessEveryoneRule) implementsAccessGroupListResponseRequire() {
}

// Matches an IP address block.
type AccessGroupListResponseRequireAccessIPRule struct {
	IP   AccessGroupListResponseRequireAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupListResponseRequireAccessIPRuleJSON `json:"-"`
}

// accessGroupListResponseRequireAccessIPRuleJSON contains the JSON metadata for
// the struct [AccessGroupListResponseRequireAccessIPRule]
type accessGroupListResponseRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseRequireAccessIPRule) implementsAccessGroupListResponseRequire() {}

type AccessGroupListResponseRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                           `json:"ip,required"`
	JSON accessGroupListResponseRequireAccessIPRuleIPJSON `json:"-"`
}

// accessGroupListResponseRequireAccessIPRuleIPJSON contains the JSON metadata for
// the struct [AccessGroupListResponseRequireAccessIPRuleIP]
type accessGroupListResponseRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupListResponseRequireAccessIPListRule struct {
	IPList AccessGroupListResponseRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupListResponseRequireAccessIPListRuleJSON   `json:"-"`
}

// accessGroupListResponseRequireAccessIPListRuleJSON contains the JSON metadata
// for the struct [AccessGroupListResponseRequireAccessIPListRule]
type accessGroupListResponseRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseRequireAccessIPListRule) implementsAccessGroupListResponseRequire() {}

type AccessGroupListResponseRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                   `json:"id,required"`
	JSON accessGroupListResponseRequireAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupListResponseRequireAccessIPListRuleIPListJSON contains the JSON
// metadata for the struct [AccessGroupListResponseRequireAccessIPListRuleIPList]
type accessGroupListResponseRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupListResponseRequireAccessCertificateRule struct {
	Certificate interface{}                                             `json:"certificate,required"`
	JSON        accessGroupListResponseRequireAccessCertificateRuleJSON `json:"-"`
}

// accessGroupListResponseRequireAccessCertificateRuleJSON contains the JSON
// metadata for the struct [AccessGroupListResponseRequireAccessCertificateRule]
type accessGroupListResponseRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseRequireAccessCertificateRule) implementsAccessGroupListResponseRequire() {
}

// Matches an Access group.
type AccessGroupListResponseRequireAccessAccessGroupRule struct {
	Group AccessGroupListResponseRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupListResponseRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupListResponseRequireAccessAccessGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupListResponseRequireAccessAccessGroupRule]
type accessGroupListResponseRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseRequireAccessAccessGroupRule) implementsAccessGroupListResponseRequire() {
}

type AccessGroupListResponseRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                       `json:"id,required"`
	JSON accessGroupListResponseRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupListResponseRequireAccessAccessGroupRuleGroupJSON contains the JSON
// metadata for the struct
// [AccessGroupListResponseRequireAccessAccessGroupRuleGroup]
type accessGroupListResponseRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupListResponseRequireAccessAzureGroupRule struct {
	AzureAd AccessGroupListResponseRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupListResponseRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupListResponseRequireAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupListResponseRequireAccessAzureGroupRule]
type accessGroupListResponseRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseRequireAccessAzureGroupRule) implementsAccessGroupListResponseRequire() {
}

type AccessGroupListResponseRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                        `json:"connection_id,required"`
	JSON         accessGroupListResponseRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupListResponseRequireAccessAzureGroupRuleAzureAdJSON contains the JSON
// metadata for the struct
// [AccessGroupListResponseRequireAccessAzureGroupRuleAzureAd]
type accessGroupListResponseRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupListResponseRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupListResponseRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupListResponseRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupListResponseRequireAccessGitHubOrganizationRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupListResponseRequireAccessGitHubOrganizationRule]
type accessGroupListResponseRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseRequireAccessGitHubOrganizationRule) implementsAccessGroupListResponseRequire() {
}

type AccessGroupListResponseRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                           `json:"name,required"`
	JSON accessGroupListResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupListResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupListResponseRequireAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupListResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupListResponseRequireAccessGsuiteGroupRule struct {
	Gsuite AccessGroupListResponseRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupListResponseRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupListResponseRequireAccessGsuiteGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupListResponseRequireAccessGsuiteGroupRule]
type accessGroupListResponseRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseRequireAccessGsuiteGroupRule) implementsAccessGroupListResponseRequire() {
}

type AccessGroupListResponseRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                        `json:"email,required"`
	JSON  accessGroupListResponseRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupListResponseRequireAccessGsuiteGroupRuleGsuiteJSON contains the JSON
// metadata for the struct
// [AccessGroupListResponseRequireAccessGsuiteGroupRuleGsuite]
type accessGroupListResponseRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupListResponseRequireAccessOktaGroupRule struct {
	Okta AccessGroupListResponseRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupListResponseRequireAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupListResponseRequireAccessOktaGroupRuleJSON contains the JSON metadata
// for the struct [AccessGroupListResponseRequireAccessOktaGroupRule]
type accessGroupListResponseRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseRequireAccessOktaGroupRule) implementsAccessGroupListResponseRequire() {
}

type AccessGroupListResponseRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                    `json:"email,required"`
	JSON  accessGroupListResponseRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupListResponseRequireAccessOktaGroupRuleOktaJSON contains the JSON
// metadata for the struct [AccessGroupListResponseRequireAccessOktaGroupRuleOkta]
type accessGroupListResponseRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupListResponseRequireAccessSamlGroupRule struct {
	Saml AccessGroupListResponseRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupListResponseRequireAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupListResponseRequireAccessSamlGroupRuleJSON contains the JSON metadata
// for the struct [AccessGroupListResponseRequireAccessSamlGroupRule]
type accessGroupListResponseRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseRequireAccessSamlGroupRule) implementsAccessGroupListResponseRequire() {
}

type AccessGroupListResponseRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                    `json:"attribute_value,required"`
	JSON           accessGroupListResponseRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupListResponseRequireAccessSamlGroupRuleSamlJSON contains the JSON
// metadata for the struct [AccessGroupListResponseRequireAccessSamlGroupRuleSaml]
type accessGroupListResponseRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupListResponseRequireAccessServiceTokenRule struct {
	ServiceToken AccessGroupListResponseRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupListResponseRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupListResponseRequireAccessServiceTokenRuleJSON contains the JSON
// metadata for the struct [AccessGroupListResponseRequireAccessServiceTokenRule]
type accessGroupListResponseRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseRequireAccessServiceTokenRule) implementsAccessGroupListResponseRequire() {
}

type AccessGroupListResponseRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                               `json:"token_id,required"`
	JSON    accessGroupListResponseRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupListResponseRequireAccessServiceTokenRuleServiceTokenJSON contains
// the JSON metadata for the struct
// [AccessGroupListResponseRequireAccessServiceTokenRuleServiceToken]
type accessGroupListResponseRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupListResponseRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                      `json:"any_valid_service_token,required"`
	JSON                 accessGroupListResponseRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupListResponseRequireAccessAnyValidServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupListResponseRequireAccessAnyValidServiceTokenRule]
type accessGroupListResponseRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseRequireAccessAnyValidServiceTokenRule) implementsAccessGroupListResponseRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupListResponseRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupListResponseRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupListResponseRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupListResponseRequireAccessExternalEvaluationRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupListResponseRequireAccessExternalEvaluationRule]
type accessGroupListResponseRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseRequireAccessExternalEvaluationRule) implementsAccessGroupListResponseRequire() {
}

type AccessGroupListResponseRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                           `json:"keys_url,required"`
	JSON    accessGroupListResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupListResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupListResponseRequireAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupListResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupListResponseRequireAccessCountryRule struct {
	Geo  AccessGroupListResponseRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupListResponseRequireAccessCountryRuleJSON `json:"-"`
}

// accessGroupListResponseRequireAccessCountryRuleJSON contains the JSON metadata
// for the struct [AccessGroupListResponseRequireAccessCountryRule]
type accessGroupListResponseRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseRequireAccessCountryRule) implementsAccessGroupListResponseRequire() {}

type AccessGroupListResponseRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                 `json:"country_code,required"`
	JSON        accessGroupListResponseRequireAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupListResponseRequireAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct [AccessGroupListResponseRequireAccessCountryRuleGeo]
type accessGroupListResponseRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupListResponseRequireAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupListResponseRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupListResponseRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupListResponseRequireAccessAuthenticationMethodRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupListResponseRequireAccessAuthenticationMethodRule]
type accessGroupListResponseRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseRequireAccessAuthenticationMethodRule) implementsAccessGroupListResponseRequire() {
}

type AccessGroupListResponseRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                     `json:"auth_method,required"`
	JSON       accessGroupListResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupListResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupListResponseRequireAccessAuthenticationMethodRuleAuthMethod]
type accessGroupListResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupListResponseRequireAccessDevicePostureRule struct {
	DevicePosture AccessGroupListResponseRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupListResponseRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupListResponseRequireAccessDevicePostureRuleJSON contains the JSON
// metadata for the struct [AccessGroupListResponseRequireAccessDevicePostureRule]
type accessGroupListResponseRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupListResponseRequireAccessDevicePostureRule) implementsAccessGroupListResponseRequire() {
}

type AccessGroupListResponseRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                 `json:"integration_uid,required"`
	JSON           accessGroupListResponseRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupListResponseRequireAccessDevicePostureRuleDevicePostureJSON contains
// the JSON metadata for the struct
// [AccessGroupListResponseRequireAccessDevicePostureRuleDevicePosture]
type accessGroupListResponseRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupListResponseRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
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

type AccessGroupReplaceResponse struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessGroupReplaceResponseExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessGroupReplaceResponseInclude `json:"include"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	IsDefault []AccessGroupReplaceResponseIsDefault `json:"is_default"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []AccessGroupReplaceResponseRequire `json:"require"`
	UpdatedAt time.Time                           `json:"updated_at" format:"date-time"`
	JSON      accessGroupReplaceResponseJSON      `json:"-"`
}

// accessGroupReplaceResponseJSON contains the JSON metadata for the struct
// [AccessGroupReplaceResponse]
type accessGroupReplaceResponseJSON struct {
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

func (r *AccessGroupReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupReplaceResponseExcludeAccessEmailRule],
// [AccessGroupReplaceResponseExcludeAccessEmailListRule],
// [AccessGroupReplaceResponseExcludeAccessDomainRule],
// [AccessGroupReplaceResponseExcludeAccessEveryoneRule],
// [AccessGroupReplaceResponseExcludeAccessIPRule],
// [AccessGroupReplaceResponseExcludeAccessIPListRule],
// [AccessGroupReplaceResponseExcludeAccessCertificateRule],
// [AccessGroupReplaceResponseExcludeAccessAccessGroupRule],
// [AccessGroupReplaceResponseExcludeAccessAzureGroupRule],
// [AccessGroupReplaceResponseExcludeAccessGitHubOrganizationRule],
// [AccessGroupReplaceResponseExcludeAccessGsuiteGroupRule],
// [AccessGroupReplaceResponseExcludeAccessOktaGroupRule],
// [AccessGroupReplaceResponseExcludeAccessSamlGroupRule],
// [AccessGroupReplaceResponseExcludeAccessServiceTokenRule],
// [AccessGroupReplaceResponseExcludeAccessAnyValidServiceTokenRule],
// [AccessGroupReplaceResponseExcludeAccessExternalEvaluationRule],
// [AccessGroupReplaceResponseExcludeAccessCountryRule],
// [AccessGroupReplaceResponseExcludeAccessAuthenticationMethodRule] or
// [AccessGroupReplaceResponseExcludeAccessDevicePostureRule].
type AccessGroupReplaceResponseExclude interface {
	implementsAccessGroupReplaceResponseExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupReplaceResponseExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupReplaceResponseExcludeAccessEmailRule struct {
	Email AccessGroupReplaceResponseExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupReplaceResponseExcludeAccessEmailRuleJSON  `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessEmailRuleJSON contains the JSON metadata
// for the struct [AccessGroupReplaceResponseExcludeAccessEmailRule]
type accessGroupReplaceResponseExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseExcludeAccessEmailRule) implementsAccessGroupReplaceResponseExclude() {
}

type AccessGroupReplaceResponseExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                    `json:"email,required" format:"email"`
	JSON  accessGroupReplaceResponseExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseExcludeAccessEmailRuleEmail]
type accessGroupReplaceResponseExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupReplaceResponseExcludeAccessEmailListRule struct {
	EmailList AccessGroupReplaceResponseExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupReplaceResponseExcludeAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessEmailListRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseExcludeAccessEmailListRule]
type accessGroupReplaceResponseExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseExcludeAccessEmailListRule) implementsAccessGroupReplaceResponseExclude() {
}

type AccessGroupReplaceResponseExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                            `json:"id,required"`
	JSON accessGroupReplaceResponseExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessEmailListRuleEmailListJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseExcludeAccessEmailListRuleEmailList]
type accessGroupReplaceResponseExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupReplaceResponseExcludeAccessDomainRule struct {
	EmailDomain AccessGroupReplaceResponseExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupReplaceResponseExcludeAccessDomainRuleJSON        `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessDomainRuleJSON contains the JSON metadata
// for the struct [AccessGroupReplaceResponseExcludeAccessDomainRule]
type accessGroupReplaceResponseExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseExcludeAccessDomainRule) implementsAccessGroupReplaceResponseExclude() {
}

type AccessGroupReplaceResponseExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                           `json:"domain,required"`
	JSON   accessGroupReplaceResponseExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessDomainRuleEmailDomainJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseExcludeAccessDomainRuleEmailDomain]
type accessGroupReplaceResponseExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupReplaceResponseExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                             `json:"everyone,required"`
	JSON     accessGroupReplaceResponseExcludeAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessEveryoneRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseExcludeAccessEveryoneRule]
type accessGroupReplaceResponseExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseExcludeAccessEveryoneRule) implementsAccessGroupReplaceResponseExclude() {
}

// Matches an IP address block.
type AccessGroupReplaceResponseExcludeAccessIPRule struct {
	IP   AccessGroupReplaceResponseExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupReplaceResponseExcludeAccessIPRuleJSON `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessIPRuleJSON contains the JSON metadata for
// the struct [AccessGroupReplaceResponseExcludeAccessIPRule]
type accessGroupReplaceResponseExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseExcludeAccessIPRule) implementsAccessGroupReplaceResponseExclude() {
}

type AccessGroupReplaceResponseExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                              `json:"ip,required"`
	JSON accessGroupReplaceResponseExcludeAccessIPRuleIPJSON `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessIPRuleIPJSON contains the JSON metadata
// for the struct [AccessGroupReplaceResponseExcludeAccessIPRuleIP]
type accessGroupReplaceResponseExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupReplaceResponseExcludeAccessIPListRule struct {
	IPList AccessGroupReplaceResponseExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupReplaceResponseExcludeAccessIPListRuleJSON   `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessIPListRuleJSON contains the JSON metadata
// for the struct [AccessGroupReplaceResponseExcludeAccessIPListRule]
type accessGroupReplaceResponseExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseExcludeAccessIPListRule) implementsAccessGroupReplaceResponseExclude() {
}

type AccessGroupReplaceResponseExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                      `json:"id,required"`
	JSON accessGroupReplaceResponseExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessIPListRuleIPListJSON contains the JSON
// metadata for the struct
// [AccessGroupReplaceResponseExcludeAccessIPListRuleIPList]
type accessGroupReplaceResponseExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupReplaceResponseExcludeAccessCertificateRule struct {
	Certificate interface{}                                                `json:"certificate,required"`
	JSON        accessGroupReplaceResponseExcludeAccessCertificateRuleJSON `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessCertificateRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseExcludeAccessCertificateRule]
type accessGroupReplaceResponseExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseExcludeAccessCertificateRule) implementsAccessGroupReplaceResponseExclude() {
}

// Matches an Access group.
type AccessGroupReplaceResponseExcludeAccessAccessGroupRule struct {
	Group AccessGroupReplaceResponseExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupReplaceResponseExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessAccessGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseExcludeAccessAccessGroupRule]
type accessGroupReplaceResponseExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseExcludeAccessAccessGroupRule) implementsAccessGroupReplaceResponseExclude() {
}

type AccessGroupReplaceResponseExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                          `json:"id,required"`
	JSON accessGroupReplaceResponseExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessAccessGroupRuleGroupJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseExcludeAccessAccessGroupRuleGroup]
type accessGroupReplaceResponseExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupReplaceResponseExcludeAccessAzureGroupRule struct {
	AzureAd AccessGroupReplaceResponseExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupReplaceResponseExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseExcludeAccessAzureGroupRule]
type accessGroupReplaceResponseExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseExcludeAccessAzureGroupRule) implementsAccessGroupReplaceResponseExclude() {
}

type AccessGroupReplaceResponseExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                           `json:"connection_id,required"`
	JSON         accessGroupReplaceResponseExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessAzureGroupRuleAzureAdJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseExcludeAccessAzureGroupRuleAzureAd]
type accessGroupReplaceResponseExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupReplaceResponseExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupReplaceResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupReplaceResponseExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessGitHubOrganizationRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseExcludeAccessGitHubOrganizationRule]
type accessGroupReplaceResponseExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseExcludeAccessGitHubOrganizationRule) implementsAccessGroupReplaceResponseExclude() {
}

type AccessGroupReplaceResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                              `json:"name,required"`
	JSON accessGroupReplaceResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupReplaceResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupReplaceResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupReplaceResponseExcludeAccessGsuiteGroupRule struct {
	Gsuite AccessGroupReplaceResponseExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupReplaceResponseExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessGsuiteGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseExcludeAccessGsuiteGroupRule]
type accessGroupReplaceResponseExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseExcludeAccessGsuiteGroupRule) implementsAccessGroupReplaceResponseExclude() {
}

type AccessGroupReplaceResponseExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                           `json:"email,required"`
	JSON  accessGroupReplaceResponseExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessGsuiteGroupRuleGsuiteJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseExcludeAccessGsuiteGroupRuleGsuite]
type accessGroupReplaceResponseExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupReplaceResponseExcludeAccessOktaGroupRule struct {
	Okta AccessGroupReplaceResponseExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupReplaceResponseExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseExcludeAccessOktaGroupRule]
type accessGroupReplaceResponseExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseExcludeAccessOktaGroupRule) implementsAccessGroupReplaceResponseExclude() {
}

type AccessGroupReplaceResponseExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                       `json:"email,required"`
	JSON  accessGroupReplaceResponseExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessOktaGroupRuleOktaJSON contains the JSON
// metadata for the struct
// [AccessGroupReplaceResponseExcludeAccessOktaGroupRuleOkta]
type accessGroupReplaceResponseExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupReplaceResponseExcludeAccessSamlGroupRule struct {
	Saml AccessGroupReplaceResponseExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupReplaceResponseExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseExcludeAccessSamlGroupRule]
type accessGroupReplaceResponseExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseExcludeAccessSamlGroupRule) implementsAccessGroupReplaceResponseExclude() {
}

type AccessGroupReplaceResponseExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                       `json:"attribute_value,required"`
	JSON           accessGroupReplaceResponseExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessSamlGroupRuleSamlJSON contains the JSON
// metadata for the struct
// [AccessGroupReplaceResponseExcludeAccessSamlGroupRuleSaml]
type accessGroupReplaceResponseExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupReplaceResponseExcludeAccessServiceTokenRule struct {
	ServiceToken AccessGroupReplaceResponseExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupReplaceResponseExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessServiceTokenRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupReplaceResponseExcludeAccessServiceTokenRule]
type accessGroupReplaceResponseExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseExcludeAccessServiceTokenRule) implementsAccessGroupReplaceResponseExclude() {
}

type AccessGroupReplaceResponseExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                  `json:"token_id,required"`
	JSON    accessGroupReplaceResponseExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessServiceTokenRuleServiceTokenJSON contains
// the JSON metadata for the struct
// [AccessGroupReplaceResponseExcludeAccessServiceTokenRuleServiceToken]
type accessGroupReplaceResponseExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupReplaceResponseExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                         `json:"any_valid_service_token,required"`
	JSON                 accessGroupReplaceResponseExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessAnyValidServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseExcludeAccessAnyValidServiceTokenRule]
type accessGroupReplaceResponseExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseExcludeAccessAnyValidServiceTokenRule) implementsAccessGroupReplaceResponseExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupReplaceResponseExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupReplaceResponseExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupReplaceResponseExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessExternalEvaluationRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseExcludeAccessExternalEvaluationRule]
type accessGroupReplaceResponseExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseExcludeAccessExternalEvaluationRule) implementsAccessGroupReplaceResponseExclude() {
}

type AccessGroupReplaceResponseExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                              `json:"keys_url,required"`
	JSON    accessGroupReplaceResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupReplaceResponseExcludeAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupReplaceResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupReplaceResponseExcludeAccessCountryRule struct {
	Geo  AccessGroupReplaceResponseExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupReplaceResponseExcludeAccessCountryRuleJSON `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessCountryRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseExcludeAccessCountryRule]
type accessGroupReplaceResponseExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseExcludeAccessCountryRule) implementsAccessGroupReplaceResponseExclude() {
}

type AccessGroupReplaceResponseExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                    `json:"country_code,required"`
	JSON        accessGroupReplaceResponseExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseExcludeAccessCountryRuleGeo]
type accessGroupReplaceResponseExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupReplaceResponseExcludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupReplaceResponseExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupReplaceResponseExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessAuthenticationMethodRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseExcludeAccessAuthenticationMethodRule]
type accessGroupReplaceResponseExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseExcludeAccessAuthenticationMethodRule) implementsAccessGroupReplaceResponseExclude() {
}

type AccessGroupReplaceResponseExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                        `json:"auth_method,required"`
	JSON       accessGroupReplaceResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupReplaceResponseExcludeAccessAuthenticationMethodRuleAuthMethod]
type accessGroupReplaceResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupReplaceResponseExcludeAccessDevicePostureRule struct {
	DevicePosture AccessGroupReplaceResponseExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupReplaceResponseExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessDevicePostureRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupReplaceResponseExcludeAccessDevicePostureRule]
type accessGroupReplaceResponseExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseExcludeAccessDevicePostureRule) implementsAccessGroupReplaceResponseExclude() {
}

type AccessGroupReplaceResponseExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                    `json:"integration_uid,required"`
	JSON           accessGroupReplaceResponseExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupReplaceResponseExcludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupReplaceResponseExcludeAccessDevicePostureRuleDevicePosture]
type accessGroupReplaceResponseExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupReplaceResponseIncludeAccessEmailRule],
// [AccessGroupReplaceResponseIncludeAccessEmailListRule],
// [AccessGroupReplaceResponseIncludeAccessDomainRule],
// [AccessGroupReplaceResponseIncludeAccessEveryoneRule],
// [AccessGroupReplaceResponseIncludeAccessIPRule],
// [AccessGroupReplaceResponseIncludeAccessIPListRule],
// [AccessGroupReplaceResponseIncludeAccessCertificateRule],
// [AccessGroupReplaceResponseIncludeAccessAccessGroupRule],
// [AccessGroupReplaceResponseIncludeAccessAzureGroupRule],
// [AccessGroupReplaceResponseIncludeAccessGitHubOrganizationRule],
// [AccessGroupReplaceResponseIncludeAccessGsuiteGroupRule],
// [AccessGroupReplaceResponseIncludeAccessOktaGroupRule],
// [AccessGroupReplaceResponseIncludeAccessSamlGroupRule],
// [AccessGroupReplaceResponseIncludeAccessServiceTokenRule],
// [AccessGroupReplaceResponseIncludeAccessAnyValidServiceTokenRule],
// [AccessGroupReplaceResponseIncludeAccessExternalEvaluationRule],
// [AccessGroupReplaceResponseIncludeAccessCountryRule],
// [AccessGroupReplaceResponseIncludeAccessAuthenticationMethodRule] or
// [AccessGroupReplaceResponseIncludeAccessDevicePostureRule].
type AccessGroupReplaceResponseInclude interface {
	implementsAccessGroupReplaceResponseInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupReplaceResponseInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupReplaceResponseIncludeAccessEmailRule struct {
	Email AccessGroupReplaceResponseIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupReplaceResponseIncludeAccessEmailRuleJSON  `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessEmailRuleJSON contains the JSON metadata
// for the struct [AccessGroupReplaceResponseIncludeAccessEmailRule]
type accessGroupReplaceResponseIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIncludeAccessEmailRule) implementsAccessGroupReplaceResponseInclude() {
}

type AccessGroupReplaceResponseIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                    `json:"email,required" format:"email"`
	JSON  accessGroupReplaceResponseIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseIncludeAccessEmailRuleEmail]
type accessGroupReplaceResponseIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupReplaceResponseIncludeAccessEmailListRule struct {
	EmailList AccessGroupReplaceResponseIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupReplaceResponseIncludeAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessEmailListRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseIncludeAccessEmailListRule]
type accessGroupReplaceResponseIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIncludeAccessEmailListRule) implementsAccessGroupReplaceResponseInclude() {
}

type AccessGroupReplaceResponseIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                            `json:"id,required"`
	JSON accessGroupReplaceResponseIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessEmailListRuleEmailListJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseIncludeAccessEmailListRuleEmailList]
type accessGroupReplaceResponseIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupReplaceResponseIncludeAccessDomainRule struct {
	EmailDomain AccessGroupReplaceResponseIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupReplaceResponseIncludeAccessDomainRuleJSON        `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessDomainRuleJSON contains the JSON metadata
// for the struct [AccessGroupReplaceResponseIncludeAccessDomainRule]
type accessGroupReplaceResponseIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIncludeAccessDomainRule) implementsAccessGroupReplaceResponseInclude() {
}

type AccessGroupReplaceResponseIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                           `json:"domain,required"`
	JSON   accessGroupReplaceResponseIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessDomainRuleEmailDomainJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseIncludeAccessDomainRuleEmailDomain]
type accessGroupReplaceResponseIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupReplaceResponseIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                             `json:"everyone,required"`
	JSON     accessGroupReplaceResponseIncludeAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessEveryoneRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseIncludeAccessEveryoneRule]
type accessGroupReplaceResponseIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIncludeAccessEveryoneRule) implementsAccessGroupReplaceResponseInclude() {
}

// Matches an IP address block.
type AccessGroupReplaceResponseIncludeAccessIPRule struct {
	IP   AccessGroupReplaceResponseIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupReplaceResponseIncludeAccessIPRuleJSON `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessIPRuleJSON contains the JSON metadata for
// the struct [AccessGroupReplaceResponseIncludeAccessIPRule]
type accessGroupReplaceResponseIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIncludeAccessIPRule) implementsAccessGroupReplaceResponseInclude() {
}

type AccessGroupReplaceResponseIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                              `json:"ip,required"`
	JSON accessGroupReplaceResponseIncludeAccessIPRuleIPJSON `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessIPRuleIPJSON contains the JSON metadata
// for the struct [AccessGroupReplaceResponseIncludeAccessIPRuleIP]
type accessGroupReplaceResponseIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupReplaceResponseIncludeAccessIPListRule struct {
	IPList AccessGroupReplaceResponseIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupReplaceResponseIncludeAccessIPListRuleJSON   `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessIPListRuleJSON contains the JSON metadata
// for the struct [AccessGroupReplaceResponseIncludeAccessIPListRule]
type accessGroupReplaceResponseIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIncludeAccessIPListRule) implementsAccessGroupReplaceResponseInclude() {
}

type AccessGroupReplaceResponseIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                      `json:"id,required"`
	JSON accessGroupReplaceResponseIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessIPListRuleIPListJSON contains the JSON
// metadata for the struct
// [AccessGroupReplaceResponseIncludeAccessIPListRuleIPList]
type accessGroupReplaceResponseIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupReplaceResponseIncludeAccessCertificateRule struct {
	Certificate interface{}                                                `json:"certificate,required"`
	JSON        accessGroupReplaceResponseIncludeAccessCertificateRuleJSON `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessCertificateRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseIncludeAccessCertificateRule]
type accessGroupReplaceResponseIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIncludeAccessCertificateRule) implementsAccessGroupReplaceResponseInclude() {
}

// Matches an Access group.
type AccessGroupReplaceResponseIncludeAccessAccessGroupRule struct {
	Group AccessGroupReplaceResponseIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupReplaceResponseIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessAccessGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseIncludeAccessAccessGroupRule]
type accessGroupReplaceResponseIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIncludeAccessAccessGroupRule) implementsAccessGroupReplaceResponseInclude() {
}

type AccessGroupReplaceResponseIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                          `json:"id,required"`
	JSON accessGroupReplaceResponseIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessAccessGroupRuleGroupJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseIncludeAccessAccessGroupRuleGroup]
type accessGroupReplaceResponseIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupReplaceResponseIncludeAccessAzureGroupRule struct {
	AzureAd AccessGroupReplaceResponseIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupReplaceResponseIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseIncludeAccessAzureGroupRule]
type accessGroupReplaceResponseIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIncludeAccessAzureGroupRule) implementsAccessGroupReplaceResponseInclude() {
}

type AccessGroupReplaceResponseIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                           `json:"connection_id,required"`
	JSON         accessGroupReplaceResponseIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessAzureGroupRuleAzureAdJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseIncludeAccessAzureGroupRuleAzureAd]
type accessGroupReplaceResponseIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupReplaceResponseIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupReplaceResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupReplaceResponseIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessGitHubOrganizationRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseIncludeAccessGitHubOrganizationRule]
type accessGroupReplaceResponseIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIncludeAccessGitHubOrganizationRule) implementsAccessGroupReplaceResponseInclude() {
}

type AccessGroupReplaceResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                              `json:"name,required"`
	JSON accessGroupReplaceResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupReplaceResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupReplaceResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupReplaceResponseIncludeAccessGsuiteGroupRule struct {
	Gsuite AccessGroupReplaceResponseIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupReplaceResponseIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessGsuiteGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseIncludeAccessGsuiteGroupRule]
type accessGroupReplaceResponseIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIncludeAccessGsuiteGroupRule) implementsAccessGroupReplaceResponseInclude() {
}

type AccessGroupReplaceResponseIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                           `json:"email,required"`
	JSON  accessGroupReplaceResponseIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessGsuiteGroupRuleGsuiteJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseIncludeAccessGsuiteGroupRuleGsuite]
type accessGroupReplaceResponseIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupReplaceResponseIncludeAccessOktaGroupRule struct {
	Okta AccessGroupReplaceResponseIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupReplaceResponseIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseIncludeAccessOktaGroupRule]
type accessGroupReplaceResponseIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIncludeAccessOktaGroupRule) implementsAccessGroupReplaceResponseInclude() {
}

type AccessGroupReplaceResponseIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                       `json:"email,required"`
	JSON  accessGroupReplaceResponseIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessOktaGroupRuleOktaJSON contains the JSON
// metadata for the struct
// [AccessGroupReplaceResponseIncludeAccessOktaGroupRuleOkta]
type accessGroupReplaceResponseIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupReplaceResponseIncludeAccessSamlGroupRule struct {
	Saml AccessGroupReplaceResponseIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupReplaceResponseIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseIncludeAccessSamlGroupRule]
type accessGroupReplaceResponseIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIncludeAccessSamlGroupRule) implementsAccessGroupReplaceResponseInclude() {
}

type AccessGroupReplaceResponseIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                       `json:"attribute_value,required"`
	JSON           accessGroupReplaceResponseIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessSamlGroupRuleSamlJSON contains the JSON
// metadata for the struct
// [AccessGroupReplaceResponseIncludeAccessSamlGroupRuleSaml]
type accessGroupReplaceResponseIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupReplaceResponseIncludeAccessServiceTokenRule struct {
	ServiceToken AccessGroupReplaceResponseIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupReplaceResponseIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessServiceTokenRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupReplaceResponseIncludeAccessServiceTokenRule]
type accessGroupReplaceResponseIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIncludeAccessServiceTokenRule) implementsAccessGroupReplaceResponseInclude() {
}

type AccessGroupReplaceResponseIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                  `json:"token_id,required"`
	JSON    accessGroupReplaceResponseIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessServiceTokenRuleServiceTokenJSON contains
// the JSON metadata for the struct
// [AccessGroupReplaceResponseIncludeAccessServiceTokenRuleServiceToken]
type accessGroupReplaceResponseIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupReplaceResponseIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                         `json:"any_valid_service_token,required"`
	JSON                 accessGroupReplaceResponseIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessAnyValidServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseIncludeAccessAnyValidServiceTokenRule]
type accessGroupReplaceResponseIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIncludeAccessAnyValidServiceTokenRule) implementsAccessGroupReplaceResponseInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupReplaceResponseIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupReplaceResponseIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupReplaceResponseIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessExternalEvaluationRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseIncludeAccessExternalEvaluationRule]
type accessGroupReplaceResponseIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIncludeAccessExternalEvaluationRule) implementsAccessGroupReplaceResponseInclude() {
}

type AccessGroupReplaceResponseIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                              `json:"keys_url,required"`
	JSON    accessGroupReplaceResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupReplaceResponseIncludeAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupReplaceResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupReplaceResponseIncludeAccessCountryRule struct {
	Geo  AccessGroupReplaceResponseIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupReplaceResponseIncludeAccessCountryRuleJSON `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessCountryRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseIncludeAccessCountryRule]
type accessGroupReplaceResponseIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIncludeAccessCountryRule) implementsAccessGroupReplaceResponseInclude() {
}

type AccessGroupReplaceResponseIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                    `json:"country_code,required"`
	JSON        accessGroupReplaceResponseIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseIncludeAccessCountryRuleGeo]
type accessGroupReplaceResponseIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupReplaceResponseIncludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupReplaceResponseIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupReplaceResponseIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessAuthenticationMethodRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseIncludeAccessAuthenticationMethodRule]
type accessGroupReplaceResponseIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIncludeAccessAuthenticationMethodRule) implementsAccessGroupReplaceResponseInclude() {
}

type AccessGroupReplaceResponseIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                        `json:"auth_method,required"`
	JSON       accessGroupReplaceResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupReplaceResponseIncludeAccessAuthenticationMethodRuleAuthMethod]
type accessGroupReplaceResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupReplaceResponseIncludeAccessDevicePostureRule struct {
	DevicePosture AccessGroupReplaceResponseIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupReplaceResponseIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessDevicePostureRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupReplaceResponseIncludeAccessDevicePostureRule]
type accessGroupReplaceResponseIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIncludeAccessDevicePostureRule) implementsAccessGroupReplaceResponseInclude() {
}

type AccessGroupReplaceResponseIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                    `json:"integration_uid,required"`
	JSON           accessGroupReplaceResponseIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupReplaceResponseIncludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupReplaceResponseIncludeAccessDevicePostureRuleDevicePosture]
type accessGroupReplaceResponseIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupReplaceResponseIsDefaultAccessEmailRule],
// [AccessGroupReplaceResponseIsDefaultAccessEmailListRule],
// [AccessGroupReplaceResponseIsDefaultAccessDomainRule],
// [AccessGroupReplaceResponseIsDefaultAccessEveryoneRule],
// [AccessGroupReplaceResponseIsDefaultAccessIPRule],
// [AccessGroupReplaceResponseIsDefaultAccessIPListRule],
// [AccessGroupReplaceResponseIsDefaultAccessCertificateRule],
// [AccessGroupReplaceResponseIsDefaultAccessAccessGroupRule],
// [AccessGroupReplaceResponseIsDefaultAccessAzureGroupRule],
// [AccessGroupReplaceResponseIsDefaultAccessGitHubOrganizationRule],
// [AccessGroupReplaceResponseIsDefaultAccessGsuiteGroupRule],
// [AccessGroupReplaceResponseIsDefaultAccessOktaGroupRule],
// [AccessGroupReplaceResponseIsDefaultAccessSamlGroupRule],
// [AccessGroupReplaceResponseIsDefaultAccessServiceTokenRule],
// [AccessGroupReplaceResponseIsDefaultAccessAnyValidServiceTokenRule],
// [AccessGroupReplaceResponseIsDefaultAccessExternalEvaluationRule],
// [AccessGroupReplaceResponseIsDefaultAccessCountryRule],
// [AccessGroupReplaceResponseIsDefaultAccessAuthenticationMethodRule] or
// [AccessGroupReplaceResponseIsDefaultAccessDevicePostureRule].
type AccessGroupReplaceResponseIsDefault interface {
	implementsAccessGroupReplaceResponseIsDefault()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupReplaceResponseIsDefault)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupReplaceResponseIsDefaultAccessEmailRule struct {
	Email AccessGroupReplaceResponseIsDefaultAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupReplaceResponseIsDefaultAccessEmailRuleJSON  `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessEmailRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseIsDefaultAccessEmailRule]
type accessGroupReplaceResponseIsDefaultAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIsDefaultAccessEmailRule) implementsAccessGroupReplaceResponseIsDefault() {
}

type AccessGroupReplaceResponseIsDefaultAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                      `json:"email,required" format:"email"`
	JSON  accessGroupReplaceResponseIsDefaultAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct
// [AccessGroupReplaceResponseIsDefaultAccessEmailRuleEmail]
type accessGroupReplaceResponseIsDefaultAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupReplaceResponseIsDefaultAccessEmailListRule struct {
	EmailList AccessGroupReplaceResponseIsDefaultAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupReplaceResponseIsDefaultAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessEmailListRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseIsDefaultAccessEmailListRule]
type accessGroupReplaceResponseIsDefaultAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIsDefaultAccessEmailListRule) implementsAccessGroupReplaceResponseIsDefault() {
}

type AccessGroupReplaceResponseIsDefaultAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                              `json:"id,required"`
	JSON accessGroupReplaceResponseIsDefaultAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessEmailListRuleEmailListJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseIsDefaultAccessEmailListRuleEmailList]
type accessGroupReplaceResponseIsDefaultAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupReplaceResponseIsDefaultAccessDomainRule struct {
	EmailDomain AccessGroupReplaceResponseIsDefaultAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupReplaceResponseIsDefaultAccessDomainRuleJSON        `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessDomainRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseIsDefaultAccessDomainRule]
type accessGroupReplaceResponseIsDefaultAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIsDefaultAccessDomainRule) implementsAccessGroupReplaceResponseIsDefault() {
}

type AccessGroupReplaceResponseIsDefaultAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                             `json:"domain,required"`
	JSON   accessGroupReplaceResponseIsDefaultAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessDomainRuleEmailDomainJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseIsDefaultAccessDomainRuleEmailDomain]
type accessGroupReplaceResponseIsDefaultAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupReplaceResponseIsDefaultAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                               `json:"everyone,required"`
	JSON     accessGroupReplaceResponseIsDefaultAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessEveryoneRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseIsDefaultAccessEveryoneRule]
type accessGroupReplaceResponseIsDefaultAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIsDefaultAccessEveryoneRule) implementsAccessGroupReplaceResponseIsDefault() {
}

// Matches an IP address block.
type AccessGroupReplaceResponseIsDefaultAccessIPRule struct {
	IP   AccessGroupReplaceResponseIsDefaultAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupReplaceResponseIsDefaultAccessIPRuleJSON `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessIPRuleJSON contains the JSON metadata
// for the struct [AccessGroupReplaceResponseIsDefaultAccessIPRule]
type accessGroupReplaceResponseIsDefaultAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIsDefaultAccessIPRule) implementsAccessGroupReplaceResponseIsDefault() {
}

type AccessGroupReplaceResponseIsDefaultAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                `json:"ip,required"`
	JSON accessGroupReplaceResponseIsDefaultAccessIPRuleIPJSON `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessIPRuleIPJSON contains the JSON metadata
// for the struct [AccessGroupReplaceResponseIsDefaultAccessIPRuleIP]
type accessGroupReplaceResponseIsDefaultAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupReplaceResponseIsDefaultAccessIPListRule struct {
	IPList AccessGroupReplaceResponseIsDefaultAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupReplaceResponseIsDefaultAccessIPListRuleJSON   `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessIPListRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseIsDefaultAccessIPListRule]
type accessGroupReplaceResponseIsDefaultAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIsDefaultAccessIPListRule) implementsAccessGroupReplaceResponseIsDefault() {
}

type AccessGroupReplaceResponseIsDefaultAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                        `json:"id,required"`
	JSON accessGroupReplaceResponseIsDefaultAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessIPListRuleIPListJSON contains the JSON
// metadata for the struct
// [AccessGroupReplaceResponseIsDefaultAccessIPListRuleIPList]
type accessGroupReplaceResponseIsDefaultAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupReplaceResponseIsDefaultAccessCertificateRule struct {
	Certificate interface{}                                                  `json:"certificate,required"`
	JSON        accessGroupReplaceResponseIsDefaultAccessCertificateRuleJSON `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessCertificateRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupReplaceResponseIsDefaultAccessCertificateRule]
type accessGroupReplaceResponseIsDefaultAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIsDefaultAccessCertificateRule) implementsAccessGroupReplaceResponseIsDefault() {
}

// Matches an Access group.
type AccessGroupReplaceResponseIsDefaultAccessAccessGroupRule struct {
	Group AccessGroupReplaceResponseIsDefaultAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupReplaceResponseIsDefaultAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessAccessGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupReplaceResponseIsDefaultAccessAccessGroupRule]
type accessGroupReplaceResponseIsDefaultAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIsDefaultAccessAccessGroupRule) implementsAccessGroupReplaceResponseIsDefault() {
}

type AccessGroupReplaceResponseIsDefaultAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                            `json:"id,required"`
	JSON accessGroupReplaceResponseIsDefaultAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessAccessGroupRuleGroupJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseIsDefaultAccessAccessGroupRuleGroup]
type accessGroupReplaceResponseIsDefaultAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupReplaceResponseIsDefaultAccessAzureGroupRule struct {
	AzureAd AccessGroupReplaceResponseIsDefaultAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupReplaceResponseIsDefaultAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupReplaceResponseIsDefaultAccessAzureGroupRule]
type accessGroupReplaceResponseIsDefaultAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIsDefaultAccessAzureGroupRule) implementsAccessGroupReplaceResponseIsDefault() {
}

type AccessGroupReplaceResponseIsDefaultAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                             `json:"connection_id,required"`
	JSON         accessGroupReplaceResponseIsDefaultAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessAzureGroupRuleAzureAdJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseIsDefaultAccessAzureGroupRuleAzureAd]
type accessGroupReplaceResponseIsDefaultAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupReplaceResponseIsDefaultAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupReplaceResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupReplaceResponseIsDefaultAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessGitHubOrganizationRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseIsDefaultAccessGitHubOrganizationRule]
type accessGroupReplaceResponseIsDefaultAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIsDefaultAccessGitHubOrganizationRule) implementsAccessGroupReplaceResponseIsDefault() {
}

type AccessGroupReplaceResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                `json:"name,required"`
	JSON accessGroupReplaceResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupReplaceResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupReplaceResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupReplaceResponseIsDefaultAccessGsuiteGroupRule struct {
	Gsuite AccessGroupReplaceResponseIsDefaultAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupReplaceResponseIsDefaultAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessGsuiteGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupReplaceResponseIsDefaultAccessGsuiteGroupRule]
type accessGroupReplaceResponseIsDefaultAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIsDefaultAccessGsuiteGroupRule) implementsAccessGroupReplaceResponseIsDefault() {
}

type AccessGroupReplaceResponseIsDefaultAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                             `json:"email,required"`
	JSON  accessGroupReplaceResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseIsDefaultAccessGsuiteGroupRuleGsuite]
type accessGroupReplaceResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupReplaceResponseIsDefaultAccessOktaGroupRule struct {
	Okta AccessGroupReplaceResponseIsDefaultAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupReplaceResponseIsDefaultAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseIsDefaultAccessOktaGroupRule]
type accessGroupReplaceResponseIsDefaultAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIsDefaultAccessOktaGroupRule) implementsAccessGroupReplaceResponseIsDefault() {
}

type AccessGroupReplaceResponseIsDefaultAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                         `json:"email,required"`
	JSON  accessGroupReplaceResponseIsDefaultAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessOktaGroupRuleOktaJSON contains the JSON
// metadata for the struct
// [AccessGroupReplaceResponseIsDefaultAccessOktaGroupRuleOkta]
type accessGroupReplaceResponseIsDefaultAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupReplaceResponseIsDefaultAccessSamlGroupRule struct {
	Saml AccessGroupReplaceResponseIsDefaultAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupReplaceResponseIsDefaultAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseIsDefaultAccessSamlGroupRule]
type accessGroupReplaceResponseIsDefaultAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIsDefaultAccessSamlGroupRule) implementsAccessGroupReplaceResponseIsDefault() {
}

type AccessGroupReplaceResponseIsDefaultAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                         `json:"attribute_value,required"`
	JSON           accessGroupReplaceResponseIsDefaultAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessSamlGroupRuleSamlJSON contains the JSON
// metadata for the struct
// [AccessGroupReplaceResponseIsDefaultAccessSamlGroupRuleSaml]
type accessGroupReplaceResponseIsDefaultAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupReplaceResponseIsDefaultAccessServiceTokenRule struct {
	ServiceToken AccessGroupReplaceResponseIsDefaultAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupReplaceResponseIsDefaultAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessServiceTokenRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupReplaceResponseIsDefaultAccessServiceTokenRule]
type accessGroupReplaceResponseIsDefaultAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIsDefaultAccessServiceTokenRule) implementsAccessGroupReplaceResponseIsDefault() {
}

type AccessGroupReplaceResponseIsDefaultAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                    `json:"token_id,required"`
	JSON    accessGroupReplaceResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessGroupReplaceResponseIsDefaultAccessServiceTokenRuleServiceToken]
type accessGroupReplaceResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupReplaceResponseIsDefaultAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                           `json:"any_valid_service_token,required"`
	JSON                 accessGroupReplaceResponseIsDefaultAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessAnyValidServiceTokenRuleJSON contains
// the JSON metadata for the struct
// [AccessGroupReplaceResponseIsDefaultAccessAnyValidServiceTokenRule]
type accessGroupReplaceResponseIsDefaultAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIsDefaultAccessAnyValidServiceTokenRule) implementsAccessGroupReplaceResponseIsDefault() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupReplaceResponseIsDefaultAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupReplaceResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupReplaceResponseIsDefaultAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessExternalEvaluationRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseIsDefaultAccessExternalEvaluationRule]
type accessGroupReplaceResponseIsDefaultAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIsDefaultAccessExternalEvaluationRule) implementsAccessGroupReplaceResponseIsDefault() {
}

type AccessGroupReplaceResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                `json:"keys_url,required"`
	JSON    accessGroupReplaceResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupReplaceResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupReplaceResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupReplaceResponseIsDefaultAccessCountryRule struct {
	Geo  AccessGroupReplaceResponseIsDefaultAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupReplaceResponseIsDefaultAccessCountryRuleJSON `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessCountryRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseIsDefaultAccessCountryRule]
type accessGroupReplaceResponseIsDefaultAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIsDefaultAccessCountryRule) implementsAccessGroupReplaceResponseIsDefault() {
}

type AccessGroupReplaceResponseIsDefaultAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                      `json:"country_code,required"`
	JSON        accessGroupReplaceResponseIsDefaultAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct
// [AccessGroupReplaceResponseIsDefaultAccessCountryRuleGeo]
type accessGroupReplaceResponseIsDefaultAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupReplaceResponseIsDefaultAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupReplaceResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupReplaceResponseIsDefaultAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessAuthenticationMethodRuleJSON contains
// the JSON metadata for the struct
// [AccessGroupReplaceResponseIsDefaultAccessAuthenticationMethodRule]
type accessGroupReplaceResponseIsDefaultAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIsDefaultAccessAuthenticationMethodRule) implementsAccessGroupReplaceResponseIsDefault() {
}

type AccessGroupReplaceResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                          `json:"auth_method,required"`
	JSON       accessGroupReplaceResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupReplaceResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod]
type accessGroupReplaceResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupReplaceResponseIsDefaultAccessDevicePostureRule struct {
	DevicePosture AccessGroupReplaceResponseIsDefaultAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupReplaceResponseIsDefaultAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessDevicePostureRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupReplaceResponseIsDefaultAccessDevicePostureRule]
type accessGroupReplaceResponseIsDefaultAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseIsDefaultAccessDevicePostureRule) implementsAccessGroupReplaceResponseIsDefault() {
}

type AccessGroupReplaceResponseIsDefaultAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                      `json:"integration_uid,required"`
	JSON           accessGroupReplaceResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupReplaceResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupReplaceResponseIsDefaultAccessDevicePostureRuleDevicePosture]
type accessGroupReplaceResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseIsDefaultAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupReplaceResponseRequireAccessEmailRule],
// [AccessGroupReplaceResponseRequireAccessEmailListRule],
// [AccessGroupReplaceResponseRequireAccessDomainRule],
// [AccessGroupReplaceResponseRequireAccessEveryoneRule],
// [AccessGroupReplaceResponseRequireAccessIPRule],
// [AccessGroupReplaceResponseRequireAccessIPListRule],
// [AccessGroupReplaceResponseRequireAccessCertificateRule],
// [AccessGroupReplaceResponseRequireAccessAccessGroupRule],
// [AccessGroupReplaceResponseRequireAccessAzureGroupRule],
// [AccessGroupReplaceResponseRequireAccessGitHubOrganizationRule],
// [AccessGroupReplaceResponseRequireAccessGsuiteGroupRule],
// [AccessGroupReplaceResponseRequireAccessOktaGroupRule],
// [AccessGroupReplaceResponseRequireAccessSamlGroupRule],
// [AccessGroupReplaceResponseRequireAccessServiceTokenRule],
// [AccessGroupReplaceResponseRequireAccessAnyValidServiceTokenRule],
// [AccessGroupReplaceResponseRequireAccessExternalEvaluationRule],
// [AccessGroupReplaceResponseRequireAccessCountryRule],
// [AccessGroupReplaceResponseRequireAccessAuthenticationMethodRule] or
// [AccessGroupReplaceResponseRequireAccessDevicePostureRule].
type AccessGroupReplaceResponseRequire interface {
	implementsAccessGroupReplaceResponseRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupReplaceResponseRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupReplaceResponseRequireAccessEmailRule struct {
	Email AccessGroupReplaceResponseRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupReplaceResponseRequireAccessEmailRuleJSON  `json:"-"`
}

// accessGroupReplaceResponseRequireAccessEmailRuleJSON contains the JSON metadata
// for the struct [AccessGroupReplaceResponseRequireAccessEmailRule]
type accessGroupReplaceResponseRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseRequireAccessEmailRule) implementsAccessGroupReplaceResponseRequire() {
}

type AccessGroupReplaceResponseRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                    `json:"email,required" format:"email"`
	JSON  accessGroupReplaceResponseRequireAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupReplaceResponseRequireAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseRequireAccessEmailRuleEmail]
type accessGroupReplaceResponseRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupReplaceResponseRequireAccessEmailListRule struct {
	EmailList AccessGroupReplaceResponseRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupReplaceResponseRequireAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupReplaceResponseRequireAccessEmailListRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseRequireAccessEmailListRule]
type accessGroupReplaceResponseRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseRequireAccessEmailListRule) implementsAccessGroupReplaceResponseRequire() {
}

type AccessGroupReplaceResponseRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                            `json:"id,required"`
	JSON accessGroupReplaceResponseRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupReplaceResponseRequireAccessEmailListRuleEmailListJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseRequireAccessEmailListRuleEmailList]
type accessGroupReplaceResponseRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupReplaceResponseRequireAccessDomainRule struct {
	EmailDomain AccessGroupReplaceResponseRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupReplaceResponseRequireAccessDomainRuleJSON        `json:"-"`
}

// accessGroupReplaceResponseRequireAccessDomainRuleJSON contains the JSON metadata
// for the struct [AccessGroupReplaceResponseRequireAccessDomainRule]
type accessGroupReplaceResponseRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseRequireAccessDomainRule) implementsAccessGroupReplaceResponseRequire() {
}

type AccessGroupReplaceResponseRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                           `json:"domain,required"`
	JSON   accessGroupReplaceResponseRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupReplaceResponseRequireAccessDomainRuleEmailDomainJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseRequireAccessDomainRuleEmailDomain]
type accessGroupReplaceResponseRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupReplaceResponseRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                             `json:"everyone,required"`
	JSON     accessGroupReplaceResponseRequireAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupReplaceResponseRequireAccessEveryoneRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseRequireAccessEveryoneRule]
type accessGroupReplaceResponseRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseRequireAccessEveryoneRule) implementsAccessGroupReplaceResponseRequire() {
}

// Matches an IP address block.
type AccessGroupReplaceResponseRequireAccessIPRule struct {
	IP   AccessGroupReplaceResponseRequireAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupReplaceResponseRequireAccessIPRuleJSON `json:"-"`
}

// accessGroupReplaceResponseRequireAccessIPRuleJSON contains the JSON metadata for
// the struct [AccessGroupReplaceResponseRequireAccessIPRule]
type accessGroupReplaceResponseRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseRequireAccessIPRule) implementsAccessGroupReplaceResponseRequire() {
}

type AccessGroupReplaceResponseRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                              `json:"ip,required"`
	JSON accessGroupReplaceResponseRequireAccessIPRuleIPJSON `json:"-"`
}

// accessGroupReplaceResponseRequireAccessIPRuleIPJSON contains the JSON metadata
// for the struct [AccessGroupReplaceResponseRequireAccessIPRuleIP]
type accessGroupReplaceResponseRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupReplaceResponseRequireAccessIPListRule struct {
	IPList AccessGroupReplaceResponseRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupReplaceResponseRequireAccessIPListRuleJSON   `json:"-"`
}

// accessGroupReplaceResponseRequireAccessIPListRuleJSON contains the JSON metadata
// for the struct [AccessGroupReplaceResponseRequireAccessIPListRule]
type accessGroupReplaceResponseRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseRequireAccessIPListRule) implementsAccessGroupReplaceResponseRequire() {
}

type AccessGroupReplaceResponseRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                      `json:"id,required"`
	JSON accessGroupReplaceResponseRequireAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupReplaceResponseRequireAccessIPListRuleIPListJSON contains the JSON
// metadata for the struct
// [AccessGroupReplaceResponseRequireAccessIPListRuleIPList]
type accessGroupReplaceResponseRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupReplaceResponseRequireAccessCertificateRule struct {
	Certificate interface{}                                                `json:"certificate,required"`
	JSON        accessGroupReplaceResponseRequireAccessCertificateRuleJSON `json:"-"`
}

// accessGroupReplaceResponseRequireAccessCertificateRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseRequireAccessCertificateRule]
type accessGroupReplaceResponseRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseRequireAccessCertificateRule) implementsAccessGroupReplaceResponseRequire() {
}

// Matches an Access group.
type AccessGroupReplaceResponseRequireAccessAccessGroupRule struct {
	Group AccessGroupReplaceResponseRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupReplaceResponseRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupReplaceResponseRequireAccessAccessGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseRequireAccessAccessGroupRule]
type accessGroupReplaceResponseRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseRequireAccessAccessGroupRule) implementsAccessGroupReplaceResponseRequire() {
}

type AccessGroupReplaceResponseRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                          `json:"id,required"`
	JSON accessGroupReplaceResponseRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupReplaceResponseRequireAccessAccessGroupRuleGroupJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseRequireAccessAccessGroupRuleGroup]
type accessGroupReplaceResponseRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupReplaceResponseRequireAccessAzureGroupRule struct {
	AzureAd AccessGroupReplaceResponseRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupReplaceResponseRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupReplaceResponseRequireAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseRequireAccessAzureGroupRule]
type accessGroupReplaceResponseRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseRequireAccessAzureGroupRule) implementsAccessGroupReplaceResponseRequire() {
}

type AccessGroupReplaceResponseRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                           `json:"connection_id,required"`
	JSON         accessGroupReplaceResponseRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupReplaceResponseRequireAccessAzureGroupRuleAzureAdJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseRequireAccessAzureGroupRuleAzureAd]
type accessGroupReplaceResponseRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupReplaceResponseRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupReplaceResponseRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupReplaceResponseRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupReplaceResponseRequireAccessGitHubOrganizationRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseRequireAccessGitHubOrganizationRule]
type accessGroupReplaceResponseRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseRequireAccessGitHubOrganizationRule) implementsAccessGroupReplaceResponseRequire() {
}

type AccessGroupReplaceResponseRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                              `json:"name,required"`
	JSON accessGroupReplaceResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupReplaceResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupReplaceResponseRequireAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupReplaceResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupReplaceResponseRequireAccessGsuiteGroupRule struct {
	Gsuite AccessGroupReplaceResponseRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupReplaceResponseRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupReplaceResponseRequireAccessGsuiteGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseRequireAccessGsuiteGroupRule]
type accessGroupReplaceResponseRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseRequireAccessGsuiteGroupRule) implementsAccessGroupReplaceResponseRequire() {
}

type AccessGroupReplaceResponseRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                           `json:"email,required"`
	JSON  accessGroupReplaceResponseRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupReplaceResponseRequireAccessGsuiteGroupRuleGsuiteJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseRequireAccessGsuiteGroupRuleGsuite]
type accessGroupReplaceResponseRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupReplaceResponseRequireAccessOktaGroupRule struct {
	Okta AccessGroupReplaceResponseRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupReplaceResponseRequireAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupReplaceResponseRequireAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseRequireAccessOktaGroupRule]
type accessGroupReplaceResponseRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseRequireAccessOktaGroupRule) implementsAccessGroupReplaceResponseRequire() {
}

type AccessGroupReplaceResponseRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                       `json:"email,required"`
	JSON  accessGroupReplaceResponseRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupReplaceResponseRequireAccessOktaGroupRuleOktaJSON contains the JSON
// metadata for the struct
// [AccessGroupReplaceResponseRequireAccessOktaGroupRuleOkta]
type accessGroupReplaceResponseRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupReplaceResponseRequireAccessSamlGroupRule struct {
	Saml AccessGroupReplaceResponseRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupReplaceResponseRequireAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupReplaceResponseRequireAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseRequireAccessSamlGroupRule]
type accessGroupReplaceResponseRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseRequireAccessSamlGroupRule) implementsAccessGroupReplaceResponseRequire() {
}

type AccessGroupReplaceResponseRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                       `json:"attribute_value,required"`
	JSON           accessGroupReplaceResponseRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupReplaceResponseRequireAccessSamlGroupRuleSamlJSON contains the JSON
// metadata for the struct
// [AccessGroupReplaceResponseRequireAccessSamlGroupRuleSaml]
type accessGroupReplaceResponseRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupReplaceResponseRequireAccessServiceTokenRule struct {
	ServiceToken AccessGroupReplaceResponseRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupReplaceResponseRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupReplaceResponseRequireAccessServiceTokenRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupReplaceResponseRequireAccessServiceTokenRule]
type accessGroupReplaceResponseRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseRequireAccessServiceTokenRule) implementsAccessGroupReplaceResponseRequire() {
}

type AccessGroupReplaceResponseRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                  `json:"token_id,required"`
	JSON    accessGroupReplaceResponseRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupReplaceResponseRequireAccessServiceTokenRuleServiceTokenJSON contains
// the JSON metadata for the struct
// [AccessGroupReplaceResponseRequireAccessServiceTokenRuleServiceToken]
type accessGroupReplaceResponseRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupReplaceResponseRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                         `json:"any_valid_service_token,required"`
	JSON                 accessGroupReplaceResponseRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupReplaceResponseRequireAccessAnyValidServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseRequireAccessAnyValidServiceTokenRule]
type accessGroupReplaceResponseRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseRequireAccessAnyValidServiceTokenRule) implementsAccessGroupReplaceResponseRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupReplaceResponseRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupReplaceResponseRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupReplaceResponseRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupReplaceResponseRequireAccessExternalEvaluationRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseRequireAccessExternalEvaluationRule]
type accessGroupReplaceResponseRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseRequireAccessExternalEvaluationRule) implementsAccessGroupReplaceResponseRequire() {
}

type AccessGroupReplaceResponseRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                              `json:"keys_url,required"`
	JSON    accessGroupReplaceResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupReplaceResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupReplaceResponseRequireAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupReplaceResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupReplaceResponseRequireAccessCountryRule struct {
	Geo  AccessGroupReplaceResponseRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupReplaceResponseRequireAccessCountryRuleJSON `json:"-"`
}

// accessGroupReplaceResponseRequireAccessCountryRuleJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseRequireAccessCountryRule]
type accessGroupReplaceResponseRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseRequireAccessCountryRule) implementsAccessGroupReplaceResponseRequire() {
}

type AccessGroupReplaceResponseRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                    `json:"country_code,required"`
	JSON        accessGroupReplaceResponseRequireAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupReplaceResponseRequireAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct [AccessGroupReplaceResponseRequireAccessCountryRuleGeo]
type accessGroupReplaceResponseRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupReplaceResponseRequireAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupReplaceResponseRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupReplaceResponseRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupReplaceResponseRequireAccessAuthenticationMethodRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupReplaceResponseRequireAccessAuthenticationMethodRule]
type accessGroupReplaceResponseRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseRequireAccessAuthenticationMethodRule) implementsAccessGroupReplaceResponseRequire() {
}

type AccessGroupReplaceResponseRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                        `json:"auth_method,required"`
	JSON       accessGroupReplaceResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupReplaceResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupReplaceResponseRequireAccessAuthenticationMethodRuleAuthMethod]
type accessGroupReplaceResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupReplaceResponseRequireAccessDevicePostureRule struct {
	DevicePosture AccessGroupReplaceResponseRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupReplaceResponseRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupReplaceResponseRequireAccessDevicePostureRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupReplaceResponseRequireAccessDevicePostureRule]
type accessGroupReplaceResponseRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupReplaceResponseRequireAccessDevicePostureRule) implementsAccessGroupReplaceResponseRequire() {
}

type AccessGroupReplaceResponseRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                    `json:"integration_uid,required"`
	JSON           accessGroupReplaceResponseRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupReplaceResponseRequireAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupReplaceResponseRequireAccessDevicePostureRuleDevicePosture]
type accessGroupReplaceResponseRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupNewParams struct {
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessGroupNewParamsInclude] `json:"include,required"`
	// The name of the Access group.
	Name param.Field[string] `json:"name,required"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessGroupNewParamsExclude] `json:"exclude"`
	// Whether this is the default group
	IsDefault param.Field[bool] `json:"is_default"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessGroupNewParamsRequire] `json:"require"`
}

func (r AccessGroupNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [AccessGroupNewParamsIncludeAccessEmailRule],
// [AccessGroupNewParamsIncludeAccessEmailListRule],
// [AccessGroupNewParamsIncludeAccessDomainRule],
// [AccessGroupNewParamsIncludeAccessEveryoneRule],
// [AccessGroupNewParamsIncludeAccessIPRule],
// [AccessGroupNewParamsIncludeAccessIPListRule],
// [AccessGroupNewParamsIncludeAccessCertificateRule],
// [AccessGroupNewParamsIncludeAccessAccessGroupRule],
// [AccessGroupNewParamsIncludeAccessAzureGroupRule],
// [AccessGroupNewParamsIncludeAccessGitHubOrganizationRule],
// [AccessGroupNewParamsIncludeAccessGsuiteGroupRule],
// [AccessGroupNewParamsIncludeAccessOktaGroupRule],
// [AccessGroupNewParamsIncludeAccessSamlGroupRule],
// [AccessGroupNewParamsIncludeAccessServiceTokenRule],
// [AccessGroupNewParamsIncludeAccessAnyValidServiceTokenRule],
// [AccessGroupNewParamsIncludeAccessExternalEvaluationRule],
// [AccessGroupNewParamsIncludeAccessCountryRule],
// [AccessGroupNewParamsIncludeAccessAuthenticationMethodRule],
// [AccessGroupNewParamsIncludeAccessDevicePostureRule].
type AccessGroupNewParamsInclude interface {
	implementsAccessGroupNewParamsInclude()
}

// Matches a specific email.
type AccessGroupNewParamsIncludeAccessEmailRule struct {
	Email param.Field[AccessGroupNewParamsIncludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessGroupNewParamsIncludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsIncludeAccessEmailRule) implementsAccessGroupNewParamsInclude() {}

type AccessGroupNewParamsIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccessGroupNewParamsIncludeAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccessGroupNewParamsIncludeAccessEmailListRule struct {
	EmailList param.Field[AccessGroupNewParamsIncludeAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccessGroupNewParamsIncludeAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsIncludeAccessEmailListRule) implementsAccessGroupNewParamsInclude() {}

type AccessGroupNewParamsIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupNewParamsIncludeAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccessGroupNewParamsIncludeAccessDomainRule struct {
	EmailDomain param.Field[AccessGroupNewParamsIncludeAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccessGroupNewParamsIncludeAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsIncludeAccessDomainRule) implementsAccessGroupNewParamsInclude() {}

type AccessGroupNewParamsIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccessGroupNewParamsIncludeAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccessGroupNewParamsIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccessGroupNewParamsIncludeAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsIncludeAccessEveryoneRule) implementsAccessGroupNewParamsInclude() {}

// Matches an IP address block.
type AccessGroupNewParamsIncludeAccessIPRule struct {
	IP param.Field[AccessGroupNewParamsIncludeAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessGroupNewParamsIncludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsIncludeAccessIPRule) implementsAccessGroupNewParamsInclude() {}

type AccessGroupNewParamsIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccessGroupNewParamsIncludeAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccessGroupNewParamsIncludeAccessIPListRule struct {
	IPList param.Field[AccessGroupNewParamsIncludeAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccessGroupNewParamsIncludeAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsIncludeAccessIPListRule) implementsAccessGroupNewParamsInclude() {}

type AccessGroupNewParamsIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupNewParamsIncludeAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccessGroupNewParamsIncludeAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccessGroupNewParamsIncludeAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsIncludeAccessCertificateRule) implementsAccessGroupNewParamsInclude() {}

// Matches an Access group.
type AccessGroupNewParamsIncludeAccessAccessGroupRule struct {
	Group param.Field[AccessGroupNewParamsIncludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessGroupNewParamsIncludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsIncludeAccessAccessGroupRule) implementsAccessGroupNewParamsInclude() {}

type AccessGroupNewParamsIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupNewParamsIncludeAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupNewParamsIncludeAccessAzureGroupRule struct {
	AzureAd param.Field[AccessGroupNewParamsIncludeAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccessGroupNewParamsIncludeAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsIncludeAccessAzureGroupRule) implementsAccessGroupNewParamsInclude() {}

type AccessGroupNewParamsIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccessGroupNewParamsIncludeAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupNewParamsIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccessGroupNewParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccessGroupNewParamsIncludeAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsIncludeAccessGitHubOrganizationRule) implementsAccessGroupNewParamsInclude() {
}

type AccessGroupNewParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccessGroupNewParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupNewParamsIncludeAccessGsuiteGroupRule struct {
	Gsuite param.Field[AccessGroupNewParamsIncludeAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccessGroupNewParamsIncludeAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsIncludeAccessGsuiteGroupRule) implementsAccessGroupNewParamsInclude() {}

type AccessGroupNewParamsIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessGroupNewParamsIncludeAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupNewParamsIncludeAccessOktaGroupRule struct {
	Okta param.Field[AccessGroupNewParamsIncludeAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccessGroupNewParamsIncludeAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsIncludeAccessOktaGroupRule) implementsAccessGroupNewParamsInclude() {}

type AccessGroupNewParamsIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessGroupNewParamsIncludeAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupNewParamsIncludeAccessSamlGroupRule struct {
	Saml param.Field[AccessGroupNewParamsIncludeAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccessGroupNewParamsIncludeAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsIncludeAccessSamlGroupRule) implementsAccessGroupNewParamsInclude() {}

type AccessGroupNewParamsIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccessGroupNewParamsIncludeAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccessGroupNewParamsIncludeAccessServiceTokenRule struct {
	ServiceToken param.Field[AccessGroupNewParamsIncludeAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccessGroupNewParamsIncludeAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsIncludeAccessServiceTokenRule) implementsAccessGroupNewParamsInclude() {}

type AccessGroupNewParamsIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccessGroupNewParamsIncludeAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccessGroupNewParamsIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccessGroupNewParamsIncludeAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsIncludeAccessAnyValidServiceTokenRule) implementsAccessGroupNewParamsInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupNewParamsIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessGroupNewParamsIncludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessGroupNewParamsIncludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsIncludeAccessExternalEvaluationRule) implementsAccessGroupNewParamsInclude() {
}

type AccessGroupNewParamsIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccessGroupNewParamsIncludeAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccessGroupNewParamsIncludeAccessCountryRule struct {
	Geo param.Field[AccessGroupNewParamsIncludeAccessCountryRuleGeo] `json:"geo,required"`
}

func (r AccessGroupNewParamsIncludeAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsIncludeAccessCountryRule) implementsAccessGroupNewParamsInclude() {}

type AccessGroupNewParamsIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccessGroupNewParamsIncludeAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccessGroupNewParamsIncludeAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[AccessGroupNewParamsIncludeAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccessGroupNewParamsIncludeAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsIncludeAccessAuthenticationMethodRule) implementsAccessGroupNewParamsInclude() {
}

type AccessGroupNewParamsIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccessGroupNewParamsIncludeAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccessGroupNewParamsIncludeAccessDevicePostureRule struct {
	DevicePosture param.Field[AccessGroupNewParamsIncludeAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccessGroupNewParamsIncludeAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsIncludeAccessDevicePostureRule) implementsAccessGroupNewParamsInclude() {}

type AccessGroupNewParamsIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccessGroupNewParamsIncludeAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [AccessGroupNewParamsExcludeAccessEmailRule],
// [AccessGroupNewParamsExcludeAccessEmailListRule],
// [AccessGroupNewParamsExcludeAccessDomainRule],
// [AccessGroupNewParamsExcludeAccessEveryoneRule],
// [AccessGroupNewParamsExcludeAccessIPRule],
// [AccessGroupNewParamsExcludeAccessIPListRule],
// [AccessGroupNewParamsExcludeAccessCertificateRule],
// [AccessGroupNewParamsExcludeAccessAccessGroupRule],
// [AccessGroupNewParamsExcludeAccessAzureGroupRule],
// [AccessGroupNewParamsExcludeAccessGitHubOrganizationRule],
// [AccessGroupNewParamsExcludeAccessGsuiteGroupRule],
// [AccessGroupNewParamsExcludeAccessOktaGroupRule],
// [AccessGroupNewParamsExcludeAccessSamlGroupRule],
// [AccessGroupNewParamsExcludeAccessServiceTokenRule],
// [AccessGroupNewParamsExcludeAccessAnyValidServiceTokenRule],
// [AccessGroupNewParamsExcludeAccessExternalEvaluationRule],
// [AccessGroupNewParamsExcludeAccessCountryRule],
// [AccessGroupNewParamsExcludeAccessAuthenticationMethodRule],
// [AccessGroupNewParamsExcludeAccessDevicePostureRule].
type AccessGroupNewParamsExclude interface {
	implementsAccessGroupNewParamsExclude()
}

// Matches a specific email.
type AccessGroupNewParamsExcludeAccessEmailRule struct {
	Email param.Field[AccessGroupNewParamsExcludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessGroupNewParamsExcludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsExcludeAccessEmailRule) implementsAccessGroupNewParamsExclude() {}

type AccessGroupNewParamsExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccessGroupNewParamsExcludeAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccessGroupNewParamsExcludeAccessEmailListRule struct {
	EmailList param.Field[AccessGroupNewParamsExcludeAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccessGroupNewParamsExcludeAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsExcludeAccessEmailListRule) implementsAccessGroupNewParamsExclude() {}

type AccessGroupNewParamsExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupNewParamsExcludeAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccessGroupNewParamsExcludeAccessDomainRule struct {
	EmailDomain param.Field[AccessGroupNewParamsExcludeAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccessGroupNewParamsExcludeAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsExcludeAccessDomainRule) implementsAccessGroupNewParamsExclude() {}

type AccessGroupNewParamsExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccessGroupNewParamsExcludeAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccessGroupNewParamsExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccessGroupNewParamsExcludeAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsExcludeAccessEveryoneRule) implementsAccessGroupNewParamsExclude() {}

// Matches an IP address block.
type AccessGroupNewParamsExcludeAccessIPRule struct {
	IP param.Field[AccessGroupNewParamsExcludeAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessGroupNewParamsExcludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsExcludeAccessIPRule) implementsAccessGroupNewParamsExclude() {}

type AccessGroupNewParamsExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccessGroupNewParamsExcludeAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccessGroupNewParamsExcludeAccessIPListRule struct {
	IPList param.Field[AccessGroupNewParamsExcludeAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccessGroupNewParamsExcludeAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsExcludeAccessIPListRule) implementsAccessGroupNewParamsExclude() {}

type AccessGroupNewParamsExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupNewParamsExcludeAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccessGroupNewParamsExcludeAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccessGroupNewParamsExcludeAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsExcludeAccessCertificateRule) implementsAccessGroupNewParamsExclude() {}

// Matches an Access group.
type AccessGroupNewParamsExcludeAccessAccessGroupRule struct {
	Group param.Field[AccessGroupNewParamsExcludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessGroupNewParamsExcludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsExcludeAccessAccessGroupRule) implementsAccessGroupNewParamsExclude() {}

type AccessGroupNewParamsExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupNewParamsExcludeAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupNewParamsExcludeAccessAzureGroupRule struct {
	AzureAd param.Field[AccessGroupNewParamsExcludeAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccessGroupNewParamsExcludeAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsExcludeAccessAzureGroupRule) implementsAccessGroupNewParamsExclude() {}

type AccessGroupNewParamsExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccessGroupNewParamsExcludeAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupNewParamsExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccessGroupNewParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccessGroupNewParamsExcludeAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsExcludeAccessGitHubOrganizationRule) implementsAccessGroupNewParamsExclude() {
}

type AccessGroupNewParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccessGroupNewParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupNewParamsExcludeAccessGsuiteGroupRule struct {
	Gsuite param.Field[AccessGroupNewParamsExcludeAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccessGroupNewParamsExcludeAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsExcludeAccessGsuiteGroupRule) implementsAccessGroupNewParamsExclude() {}

type AccessGroupNewParamsExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessGroupNewParamsExcludeAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupNewParamsExcludeAccessOktaGroupRule struct {
	Okta param.Field[AccessGroupNewParamsExcludeAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccessGroupNewParamsExcludeAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsExcludeAccessOktaGroupRule) implementsAccessGroupNewParamsExclude() {}

type AccessGroupNewParamsExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessGroupNewParamsExcludeAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupNewParamsExcludeAccessSamlGroupRule struct {
	Saml param.Field[AccessGroupNewParamsExcludeAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccessGroupNewParamsExcludeAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsExcludeAccessSamlGroupRule) implementsAccessGroupNewParamsExclude() {}

type AccessGroupNewParamsExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccessGroupNewParamsExcludeAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccessGroupNewParamsExcludeAccessServiceTokenRule struct {
	ServiceToken param.Field[AccessGroupNewParamsExcludeAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccessGroupNewParamsExcludeAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsExcludeAccessServiceTokenRule) implementsAccessGroupNewParamsExclude() {}

type AccessGroupNewParamsExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccessGroupNewParamsExcludeAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccessGroupNewParamsExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccessGroupNewParamsExcludeAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsExcludeAccessAnyValidServiceTokenRule) implementsAccessGroupNewParamsExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupNewParamsExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessGroupNewParamsExcludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessGroupNewParamsExcludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsExcludeAccessExternalEvaluationRule) implementsAccessGroupNewParamsExclude() {
}

type AccessGroupNewParamsExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccessGroupNewParamsExcludeAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccessGroupNewParamsExcludeAccessCountryRule struct {
	Geo param.Field[AccessGroupNewParamsExcludeAccessCountryRuleGeo] `json:"geo,required"`
}

func (r AccessGroupNewParamsExcludeAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsExcludeAccessCountryRule) implementsAccessGroupNewParamsExclude() {}

type AccessGroupNewParamsExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccessGroupNewParamsExcludeAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccessGroupNewParamsExcludeAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[AccessGroupNewParamsExcludeAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccessGroupNewParamsExcludeAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsExcludeAccessAuthenticationMethodRule) implementsAccessGroupNewParamsExclude() {
}

type AccessGroupNewParamsExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccessGroupNewParamsExcludeAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccessGroupNewParamsExcludeAccessDevicePostureRule struct {
	DevicePosture param.Field[AccessGroupNewParamsExcludeAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccessGroupNewParamsExcludeAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsExcludeAccessDevicePostureRule) implementsAccessGroupNewParamsExclude() {}

type AccessGroupNewParamsExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccessGroupNewParamsExcludeAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [AccessGroupNewParamsRequireAccessEmailRule],
// [AccessGroupNewParamsRequireAccessEmailListRule],
// [AccessGroupNewParamsRequireAccessDomainRule],
// [AccessGroupNewParamsRequireAccessEveryoneRule],
// [AccessGroupNewParamsRequireAccessIPRule],
// [AccessGroupNewParamsRequireAccessIPListRule],
// [AccessGroupNewParamsRequireAccessCertificateRule],
// [AccessGroupNewParamsRequireAccessAccessGroupRule],
// [AccessGroupNewParamsRequireAccessAzureGroupRule],
// [AccessGroupNewParamsRequireAccessGitHubOrganizationRule],
// [AccessGroupNewParamsRequireAccessGsuiteGroupRule],
// [AccessGroupNewParamsRequireAccessOktaGroupRule],
// [AccessGroupNewParamsRequireAccessSamlGroupRule],
// [AccessGroupNewParamsRequireAccessServiceTokenRule],
// [AccessGroupNewParamsRequireAccessAnyValidServiceTokenRule],
// [AccessGroupNewParamsRequireAccessExternalEvaluationRule],
// [AccessGroupNewParamsRequireAccessCountryRule],
// [AccessGroupNewParamsRequireAccessAuthenticationMethodRule],
// [AccessGroupNewParamsRequireAccessDevicePostureRule].
type AccessGroupNewParamsRequire interface {
	implementsAccessGroupNewParamsRequire()
}

// Matches a specific email.
type AccessGroupNewParamsRequireAccessEmailRule struct {
	Email param.Field[AccessGroupNewParamsRequireAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessGroupNewParamsRequireAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsRequireAccessEmailRule) implementsAccessGroupNewParamsRequire() {}

type AccessGroupNewParamsRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccessGroupNewParamsRequireAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccessGroupNewParamsRequireAccessEmailListRule struct {
	EmailList param.Field[AccessGroupNewParamsRequireAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccessGroupNewParamsRequireAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsRequireAccessEmailListRule) implementsAccessGroupNewParamsRequire() {}

type AccessGroupNewParamsRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupNewParamsRequireAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccessGroupNewParamsRequireAccessDomainRule struct {
	EmailDomain param.Field[AccessGroupNewParamsRequireAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccessGroupNewParamsRequireAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsRequireAccessDomainRule) implementsAccessGroupNewParamsRequire() {}

type AccessGroupNewParamsRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccessGroupNewParamsRequireAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccessGroupNewParamsRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccessGroupNewParamsRequireAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsRequireAccessEveryoneRule) implementsAccessGroupNewParamsRequire() {}

// Matches an IP address block.
type AccessGroupNewParamsRequireAccessIPRule struct {
	IP param.Field[AccessGroupNewParamsRequireAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessGroupNewParamsRequireAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsRequireAccessIPRule) implementsAccessGroupNewParamsRequire() {}

type AccessGroupNewParamsRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccessGroupNewParamsRequireAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccessGroupNewParamsRequireAccessIPListRule struct {
	IPList param.Field[AccessGroupNewParamsRequireAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccessGroupNewParamsRequireAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsRequireAccessIPListRule) implementsAccessGroupNewParamsRequire() {}

type AccessGroupNewParamsRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupNewParamsRequireAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccessGroupNewParamsRequireAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccessGroupNewParamsRequireAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsRequireAccessCertificateRule) implementsAccessGroupNewParamsRequire() {}

// Matches an Access group.
type AccessGroupNewParamsRequireAccessAccessGroupRule struct {
	Group param.Field[AccessGroupNewParamsRequireAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessGroupNewParamsRequireAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsRequireAccessAccessGroupRule) implementsAccessGroupNewParamsRequire() {}

type AccessGroupNewParamsRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupNewParamsRequireAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupNewParamsRequireAccessAzureGroupRule struct {
	AzureAd param.Field[AccessGroupNewParamsRequireAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccessGroupNewParamsRequireAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsRequireAccessAzureGroupRule) implementsAccessGroupNewParamsRequire() {}

type AccessGroupNewParamsRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccessGroupNewParamsRequireAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupNewParamsRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccessGroupNewParamsRequireAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccessGroupNewParamsRequireAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsRequireAccessGitHubOrganizationRule) implementsAccessGroupNewParamsRequire() {
}

type AccessGroupNewParamsRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccessGroupNewParamsRequireAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupNewParamsRequireAccessGsuiteGroupRule struct {
	Gsuite param.Field[AccessGroupNewParamsRequireAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccessGroupNewParamsRequireAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsRequireAccessGsuiteGroupRule) implementsAccessGroupNewParamsRequire() {}

type AccessGroupNewParamsRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessGroupNewParamsRequireAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupNewParamsRequireAccessOktaGroupRule struct {
	Okta param.Field[AccessGroupNewParamsRequireAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccessGroupNewParamsRequireAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsRequireAccessOktaGroupRule) implementsAccessGroupNewParamsRequire() {}

type AccessGroupNewParamsRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessGroupNewParamsRequireAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupNewParamsRequireAccessSamlGroupRule struct {
	Saml param.Field[AccessGroupNewParamsRequireAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccessGroupNewParamsRequireAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsRequireAccessSamlGroupRule) implementsAccessGroupNewParamsRequire() {}

type AccessGroupNewParamsRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccessGroupNewParamsRequireAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccessGroupNewParamsRequireAccessServiceTokenRule struct {
	ServiceToken param.Field[AccessGroupNewParamsRequireAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccessGroupNewParamsRequireAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsRequireAccessServiceTokenRule) implementsAccessGroupNewParamsRequire() {}

type AccessGroupNewParamsRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccessGroupNewParamsRequireAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccessGroupNewParamsRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccessGroupNewParamsRequireAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsRequireAccessAnyValidServiceTokenRule) implementsAccessGroupNewParamsRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupNewParamsRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessGroupNewParamsRequireAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessGroupNewParamsRequireAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsRequireAccessExternalEvaluationRule) implementsAccessGroupNewParamsRequire() {
}

type AccessGroupNewParamsRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccessGroupNewParamsRequireAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccessGroupNewParamsRequireAccessCountryRule struct {
	Geo param.Field[AccessGroupNewParamsRequireAccessCountryRuleGeo] `json:"geo,required"`
}

func (r AccessGroupNewParamsRequireAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsRequireAccessCountryRule) implementsAccessGroupNewParamsRequire() {}

type AccessGroupNewParamsRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccessGroupNewParamsRequireAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccessGroupNewParamsRequireAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[AccessGroupNewParamsRequireAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccessGroupNewParamsRequireAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsRequireAccessAuthenticationMethodRule) implementsAccessGroupNewParamsRequire() {
}

type AccessGroupNewParamsRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccessGroupNewParamsRequireAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccessGroupNewParamsRequireAccessDevicePostureRule struct {
	DevicePosture param.Field[AccessGroupNewParamsRequireAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccessGroupNewParamsRequireAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsRequireAccessDevicePostureRule) implementsAccessGroupNewParamsRequire() {}

type AccessGroupNewParamsRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccessGroupNewParamsRequireAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessGroupNewResponseEnvelope struct {
	Errors   []AccessGroupNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessGroupNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessGroupNewResponse                   `json:"result,required"`
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

type AccessGroupNewResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accessGroupNewResponseEnvelopeErrorsJSON `json:"-"`
}

// accessGroupNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessGroupNewResponseEnvelopeErrors]
type accessGroupNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupNewResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accessGroupNewResponseEnvelopeMessagesJSON `json:"-"`
}

// accessGroupNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessGroupNewResponseEnvelopeMessages]
type accessGroupNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessGroupNewResponseEnvelopeSuccess bool

const (
	AccessGroupNewResponseEnvelopeSuccessTrue AccessGroupNewResponseEnvelopeSuccess = true
)

type AccessGroupListResponseEnvelope struct {
	Errors   []AccessGroupListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessGroupListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessGroupListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AccessGroupListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AccessGroupListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       accessGroupListResponseEnvelopeJSON       `json:"-"`
}

// accessGroupListResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessGroupListResponseEnvelope]
type accessGroupListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupListResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accessGroupListResponseEnvelopeErrorsJSON `json:"-"`
}

// accessGroupListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessGroupListResponseEnvelopeErrors]
type accessGroupListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupListResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accessGroupListResponseEnvelopeMessagesJSON `json:"-"`
}

// accessGroupListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessGroupListResponseEnvelopeMessages]
type accessGroupListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessGroupListResponseEnvelopeSuccess bool

const (
	AccessGroupListResponseEnvelopeSuccessTrue AccessGroupListResponseEnvelopeSuccess = true
)

type AccessGroupListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                       `json:"total_count"`
	JSON       accessGroupListResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessGroupListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [AccessGroupListResponseEnvelopeResultInfo]
type accessGroupListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupDeleteResponseEnvelope struct {
	Errors   []AccessGroupDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessGroupDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessGroupDeleteResponse                   `json:"result,required"`
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

type AccessGroupGetResponseEnvelope struct {
	Errors   []AccessGroupGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessGroupGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessGroupGetResponse                   `json:"result,required"`
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

type AccessGroupReplaceParams struct {
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessGroupReplaceParamsInclude] `json:"include,required"`
	// The name of the Access group.
	Name param.Field[string] `json:"name,required"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessGroupReplaceParamsExclude] `json:"exclude"`
	// Whether this is the default group
	IsDefault param.Field[bool] `json:"is_default"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessGroupReplaceParamsRequire] `json:"require"`
}

func (r AccessGroupReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [AccessGroupReplaceParamsIncludeAccessEmailRule],
// [AccessGroupReplaceParamsIncludeAccessEmailListRule],
// [AccessGroupReplaceParamsIncludeAccessDomainRule],
// [AccessGroupReplaceParamsIncludeAccessEveryoneRule],
// [AccessGroupReplaceParamsIncludeAccessIPRule],
// [AccessGroupReplaceParamsIncludeAccessIPListRule],
// [AccessGroupReplaceParamsIncludeAccessCertificateRule],
// [AccessGroupReplaceParamsIncludeAccessAccessGroupRule],
// [AccessGroupReplaceParamsIncludeAccessAzureGroupRule],
// [AccessGroupReplaceParamsIncludeAccessGitHubOrganizationRule],
// [AccessGroupReplaceParamsIncludeAccessGsuiteGroupRule],
// [AccessGroupReplaceParamsIncludeAccessOktaGroupRule],
// [AccessGroupReplaceParamsIncludeAccessSamlGroupRule],
// [AccessGroupReplaceParamsIncludeAccessServiceTokenRule],
// [AccessGroupReplaceParamsIncludeAccessAnyValidServiceTokenRule],
// [AccessGroupReplaceParamsIncludeAccessExternalEvaluationRule],
// [AccessGroupReplaceParamsIncludeAccessCountryRule],
// [AccessGroupReplaceParamsIncludeAccessAuthenticationMethodRule],
// [AccessGroupReplaceParamsIncludeAccessDevicePostureRule].
type AccessGroupReplaceParamsInclude interface {
	implementsAccessGroupReplaceParamsInclude()
}

// Matches a specific email.
type AccessGroupReplaceParamsIncludeAccessEmailRule struct {
	Email param.Field[AccessGroupReplaceParamsIncludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsIncludeAccessEmailRule) implementsAccessGroupReplaceParamsInclude() {}

type AccessGroupReplaceParamsIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccessGroupReplaceParamsIncludeAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccessGroupReplaceParamsIncludeAccessEmailListRule struct {
	EmailList param.Field[AccessGroupReplaceParamsIncludeAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsIncludeAccessEmailListRule) implementsAccessGroupReplaceParamsInclude() {
}

type AccessGroupReplaceParamsIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccessGroupReplaceParamsIncludeAccessDomainRule struct {
	EmailDomain param.Field[AccessGroupReplaceParamsIncludeAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsIncludeAccessDomainRule) implementsAccessGroupReplaceParamsInclude() {
}

type AccessGroupReplaceParamsIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccessGroupReplaceParamsIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsIncludeAccessEveryoneRule) implementsAccessGroupReplaceParamsInclude() {
}

// Matches an IP address block.
type AccessGroupReplaceParamsIncludeAccessIPRule struct {
	IP param.Field[AccessGroupReplaceParamsIncludeAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsIncludeAccessIPRule) implementsAccessGroupReplaceParamsInclude() {}

type AccessGroupReplaceParamsIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccessGroupReplaceParamsIncludeAccessIPListRule struct {
	IPList param.Field[AccessGroupReplaceParamsIncludeAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsIncludeAccessIPListRule) implementsAccessGroupReplaceParamsInclude() {
}

type AccessGroupReplaceParamsIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccessGroupReplaceParamsIncludeAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsIncludeAccessCertificateRule) implementsAccessGroupReplaceParamsInclude() {
}

// Matches an Access group.
type AccessGroupReplaceParamsIncludeAccessAccessGroupRule struct {
	Group param.Field[AccessGroupReplaceParamsIncludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsIncludeAccessAccessGroupRule) implementsAccessGroupReplaceParamsInclude() {
}

type AccessGroupReplaceParamsIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupReplaceParamsIncludeAccessAzureGroupRule struct {
	AzureAd param.Field[AccessGroupReplaceParamsIncludeAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsIncludeAccessAzureGroupRule) implementsAccessGroupReplaceParamsInclude() {
}

type AccessGroupReplaceParamsIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupReplaceParamsIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccessGroupReplaceParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsIncludeAccessGitHubOrganizationRule) implementsAccessGroupReplaceParamsInclude() {
}

type AccessGroupReplaceParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupReplaceParamsIncludeAccessGsuiteGroupRule struct {
	Gsuite param.Field[AccessGroupReplaceParamsIncludeAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsIncludeAccessGsuiteGroupRule) implementsAccessGroupReplaceParamsInclude() {
}

type AccessGroupReplaceParamsIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupReplaceParamsIncludeAccessOktaGroupRule struct {
	Okta param.Field[AccessGroupReplaceParamsIncludeAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsIncludeAccessOktaGroupRule) implementsAccessGroupReplaceParamsInclude() {
}

type AccessGroupReplaceParamsIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupReplaceParamsIncludeAccessSamlGroupRule struct {
	Saml param.Field[AccessGroupReplaceParamsIncludeAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsIncludeAccessSamlGroupRule) implementsAccessGroupReplaceParamsInclude() {
}

type AccessGroupReplaceParamsIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccessGroupReplaceParamsIncludeAccessServiceTokenRule struct {
	ServiceToken param.Field[AccessGroupReplaceParamsIncludeAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsIncludeAccessServiceTokenRule) implementsAccessGroupReplaceParamsInclude() {
}

type AccessGroupReplaceParamsIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccessGroupReplaceParamsIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsIncludeAccessAnyValidServiceTokenRule) implementsAccessGroupReplaceParamsInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupReplaceParamsIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessGroupReplaceParamsIncludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsIncludeAccessExternalEvaluationRule) implementsAccessGroupReplaceParamsInclude() {
}

type AccessGroupReplaceParamsIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccessGroupReplaceParamsIncludeAccessCountryRule struct {
	Geo param.Field[AccessGroupReplaceParamsIncludeAccessCountryRuleGeo] `json:"geo,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsIncludeAccessCountryRule) implementsAccessGroupReplaceParamsInclude() {
}

type AccessGroupReplaceParamsIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccessGroupReplaceParamsIncludeAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[AccessGroupReplaceParamsIncludeAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsIncludeAccessAuthenticationMethodRule) implementsAccessGroupReplaceParamsInclude() {
}

type AccessGroupReplaceParamsIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccessGroupReplaceParamsIncludeAccessDevicePostureRule struct {
	DevicePosture param.Field[AccessGroupReplaceParamsIncludeAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsIncludeAccessDevicePostureRule) implementsAccessGroupReplaceParamsInclude() {
}

type AccessGroupReplaceParamsIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccessGroupReplaceParamsIncludeAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [AccessGroupReplaceParamsExcludeAccessEmailRule],
// [AccessGroupReplaceParamsExcludeAccessEmailListRule],
// [AccessGroupReplaceParamsExcludeAccessDomainRule],
// [AccessGroupReplaceParamsExcludeAccessEveryoneRule],
// [AccessGroupReplaceParamsExcludeAccessIPRule],
// [AccessGroupReplaceParamsExcludeAccessIPListRule],
// [AccessGroupReplaceParamsExcludeAccessCertificateRule],
// [AccessGroupReplaceParamsExcludeAccessAccessGroupRule],
// [AccessGroupReplaceParamsExcludeAccessAzureGroupRule],
// [AccessGroupReplaceParamsExcludeAccessGitHubOrganizationRule],
// [AccessGroupReplaceParamsExcludeAccessGsuiteGroupRule],
// [AccessGroupReplaceParamsExcludeAccessOktaGroupRule],
// [AccessGroupReplaceParamsExcludeAccessSamlGroupRule],
// [AccessGroupReplaceParamsExcludeAccessServiceTokenRule],
// [AccessGroupReplaceParamsExcludeAccessAnyValidServiceTokenRule],
// [AccessGroupReplaceParamsExcludeAccessExternalEvaluationRule],
// [AccessGroupReplaceParamsExcludeAccessCountryRule],
// [AccessGroupReplaceParamsExcludeAccessAuthenticationMethodRule],
// [AccessGroupReplaceParamsExcludeAccessDevicePostureRule].
type AccessGroupReplaceParamsExclude interface {
	implementsAccessGroupReplaceParamsExclude()
}

// Matches a specific email.
type AccessGroupReplaceParamsExcludeAccessEmailRule struct {
	Email param.Field[AccessGroupReplaceParamsExcludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsExcludeAccessEmailRule) implementsAccessGroupReplaceParamsExclude() {}

type AccessGroupReplaceParamsExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccessGroupReplaceParamsExcludeAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccessGroupReplaceParamsExcludeAccessEmailListRule struct {
	EmailList param.Field[AccessGroupReplaceParamsExcludeAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsExcludeAccessEmailListRule) implementsAccessGroupReplaceParamsExclude() {
}

type AccessGroupReplaceParamsExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccessGroupReplaceParamsExcludeAccessDomainRule struct {
	EmailDomain param.Field[AccessGroupReplaceParamsExcludeAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsExcludeAccessDomainRule) implementsAccessGroupReplaceParamsExclude() {
}

type AccessGroupReplaceParamsExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccessGroupReplaceParamsExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsExcludeAccessEveryoneRule) implementsAccessGroupReplaceParamsExclude() {
}

// Matches an IP address block.
type AccessGroupReplaceParamsExcludeAccessIPRule struct {
	IP param.Field[AccessGroupReplaceParamsExcludeAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsExcludeAccessIPRule) implementsAccessGroupReplaceParamsExclude() {}

type AccessGroupReplaceParamsExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccessGroupReplaceParamsExcludeAccessIPListRule struct {
	IPList param.Field[AccessGroupReplaceParamsExcludeAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsExcludeAccessIPListRule) implementsAccessGroupReplaceParamsExclude() {
}

type AccessGroupReplaceParamsExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccessGroupReplaceParamsExcludeAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsExcludeAccessCertificateRule) implementsAccessGroupReplaceParamsExclude() {
}

// Matches an Access group.
type AccessGroupReplaceParamsExcludeAccessAccessGroupRule struct {
	Group param.Field[AccessGroupReplaceParamsExcludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsExcludeAccessAccessGroupRule) implementsAccessGroupReplaceParamsExclude() {
}

type AccessGroupReplaceParamsExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupReplaceParamsExcludeAccessAzureGroupRule struct {
	AzureAd param.Field[AccessGroupReplaceParamsExcludeAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsExcludeAccessAzureGroupRule) implementsAccessGroupReplaceParamsExclude() {
}

type AccessGroupReplaceParamsExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupReplaceParamsExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccessGroupReplaceParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsExcludeAccessGitHubOrganizationRule) implementsAccessGroupReplaceParamsExclude() {
}

type AccessGroupReplaceParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupReplaceParamsExcludeAccessGsuiteGroupRule struct {
	Gsuite param.Field[AccessGroupReplaceParamsExcludeAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsExcludeAccessGsuiteGroupRule) implementsAccessGroupReplaceParamsExclude() {
}

type AccessGroupReplaceParamsExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupReplaceParamsExcludeAccessOktaGroupRule struct {
	Okta param.Field[AccessGroupReplaceParamsExcludeAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsExcludeAccessOktaGroupRule) implementsAccessGroupReplaceParamsExclude() {
}

type AccessGroupReplaceParamsExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupReplaceParamsExcludeAccessSamlGroupRule struct {
	Saml param.Field[AccessGroupReplaceParamsExcludeAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsExcludeAccessSamlGroupRule) implementsAccessGroupReplaceParamsExclude() {
}

type AccessGroupReplaceParamsExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccessGroupReplaceParamsExcludeAccessServiceTokenRule struct {
	ServiceToken param.Field[AccessGroupReplaceParamsExcludeAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsExcludeAccessServiceTokenRule) implementsAccessGroupReplaceParamsExclude() {
}

type AccessGroupReplaceParamsExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccessGroupReplaceParamsExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsExcludeAccessAnyValidServiceTokenRule) implementsAccessGroupReplaceParamsExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupReplaceParamsExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessGroupReplaceParamsExcludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsExcludeAccessExternalEvaluationRule) implementsAccessGroupReplaceParamsExclude() {
}

type AccessGroupReplaceParamsExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccessGroupReplaceParamsExcludeAccessCountryRule struct {
	Geo param.Field[AccessGroupReplaceParamsExcludeAccessCountryRuleGeo] `json:"geo,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsExcludeAccessCountryRule) implementsAccessGroupReplaceParamsExclude() {
}

type AccessGroupReplaceParamsExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccessGroupReplaceParamsExcludeAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[AccessGroupReplaceParamsExcludeAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsExcludeAccessAuthenticationMethodRule) implementsAccessGroupReplaceParamsExclude() {
}

type AccessGroupReplaceParamsExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccessGroupReplaceParamsExcludeAccessDevicePostureRule struct {
	DevicePosture param.Field[AccessGroupReplaceParamsExcludeAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsExcludeAccessDevicePostureRule) implementsAccessGroupReplaceParamsExclude() {
}

type AccessGroupReplaceParamsExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccessGroupReplaceParamsExcludeAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [AccessGroupReplaceParamsRequireAccessEmailRule],
// [AccessGroupReplaceParamsRequireAccessEmailListRule],
// [AccessGroupReplaceParamsRequireAccessDomainRule],
// [AccessGroupReplaceParamsRequireAccessEveryoneRule],
// [AccessGroupReplaceParamsRequireAccessIPRule],
// [AccessGroupReplaceParamsRequireAccessIPListRule],
// [AccessGroupReplaceParamsRequireAccessCertificateRule],
// [AccessGroupReplaceParamsRequireAccessAccessGroupRule],
// [AccessGroupReplaceParamsRequireAccessAzureGroupRule],
// [AccessGroupReplaceParamsRequireAccessGitHubOrganizationRule],
// [AccessGroupReplaceParamsRequireAccessGsuiteGroupRule],
// [AccessGroupReplaceParamsRequireAccessOktaGroupRule],
// [AccessGroupReplaceParamsRequireAccessSamlGroupRule],
// [AccessGroupReplaceParamsRequireAccessServiceTokenRule],
// [AccessGroupReplaceParamsRequireAccessAnyValidServiceTokenRule],
// [AccessGroupReplaceParamsRequireAccessExternalEvaluationRule],
// [AccessGroupReplaceParamsRequireAccessCountryRule],
// [AccessGroupReplaceParamsRequireAccessAuthenticationMethodRule],
// [AccessGroupReplaceParamsRequireAccessDevicePostureRule].
type AccessGroupReplaceParamsRequire interface {
	implementsAccessGroupReplaceParamsRequire()
}

// Matches a specific email.
type AccessGroupReplaceParamsRequireAccessEmailRule struct {
	Email param.Field[AccessGroupReplaceParamsRequireAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessGroupReplaceParamsRequireAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsRequireAccessEmailRule) implementsAccessGroupReplaceParamsRequire() {}

type AccessGroupReplaceParamsRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccessGroupReplaceParamsRequireAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccessGroupReplaceParamsRequireAccessEmailListRule struct {
	EmailList param.Field[AccessGroupReplaceParamsRequireAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccessGroupReplaceParamsRequireAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsRequireAccessEmailListRule) implementsAccessGroupReplaceParamsRequire() {
}

type AccessGroupReplaceParamsRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupReplaceParamsRequireAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccessGroupReplaceParamsRequireAccessDomainRule struct {
	EmailDomain param.Field[AccessGroupReplaceParamsRequireAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccessGroupReplaceParamsRequireAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsRequireAccessDomainRule) implementsAccessGroupReplaceParamsRequire() {
}

type AccessGroupReplaceParamsRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccessGroupReplaceParamsRequireAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccessGroupReplaceParamsRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccessGroupReplaceParamsRequireAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsRequireAccessEveryoneRule) implementsAccessGroupReplaceParamsRequire() {
}

// Matches an IP address block.
type AccessGroupReplaceParamsRequireAccessIPRule struct {
	IP param.Field[AccessGroupReplaceParamsRequireAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessGroupReplaceParamsRequireAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsRequireAccessIPRule) implementsAccessGroupReplaceParamsRequire() {}

type AccessGroupReplaceParamsRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccessGroupReplaceParamsRequireAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccessGroupReplaceParamsRequireAccessIPListRule struct {
	IPList param.Field[AccessGroupReplaceParamsRequireAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccessGroupReplaceParamsRequireAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsRequireAccessIPListRule) implementsAccessGroupReplaceParamsRequire() {
}

type AccessGroupReplaceParamsRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupReplaceParamsRequireAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccessGroupReplaceParamsRequireAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccessGroupReplaceParamsRequireAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsRequireAccessCertificateRule) implementsAccessGroupReplaceParamsRequire() {
}

// Matches an Access group.
type AccessGroupReplaceParamsRequireAccessAccessGroupRule struct {
	Group param.Field[AccessGroupReplaceParamsRequireAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessGroupReplaceParamsRequireAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsRequireAccessAccessGroupRule) implementsAccessGroupReplaceParamsRequire() {
}

type AccessGroupReplaceParamsRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccessGroupReplaceParamsRequireAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupReplaceParamsRequireAccessAzureGroupRule struct {
	AzureAd param.Field[AccessGroupReplaceParamsRequireAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccessGroupReplaceParamsRequireAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsRequireAccessAzureGroupRule) implementsAccessGroupReplaceParamsRequire() {
}

type AccessGroupReplaceParamsRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccessGroupReplaceParamsRequireAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupReplaceParamsRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccessGroupReplaceParamsRequireAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccessGroupReplaceParamsRequireAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsRequireAccessGitHubOrganizationRule) implementsAccessGroupReplaceParamsRequire() {
}

type AccessGroupReplaceParamsRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccessGroupReplaceParamsRequireAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupReplaceParamsRequireAccessGsuiteGroupRule struct {
	Gsuite param.Field[AccessGroupReplaceParamsRequireAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccessGroupReplaceParamsRequireAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsRequireAccessGsuiteGroupRule) implementsAccessGroupReplaceParamsRequire() {
}

type AccessGroupReplaceParamsRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessGroupReplaceParamsRequireAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupReplaceParamsRequireAccessOktaGroupRule struct {
	Okta param.Field[AccessGroupReplaceParamsRequireAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccessGroupReplaceParamsRequireAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsRequireAccessOktaGroupRule) implementsAccessGroupReplaceParamsRequire() {
}

type AccessGroupReplaceParamsRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccessGroupReplaceParamsRequireAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupReplaceParamsRequireAccessSamlGroupRule struct {
	Saml param.Field[AccessGroupReplaceParamsRequireAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccessGroupReplaceParamsRequireAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsRequireAccessSamlGroupRule) implementsAccessGroupReplaceParamsRequire() {
}

type AccessGroupReplaceParamsRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccessGroupReplaceParamsRequireAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccessGroupReplaceParamsRequireAccessServiceTokenRule struct {
	ServiceToken param.Field[AccessGroupReplaceParamsRequireAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccessGroupReplaceParamsRequireAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsRequireAccessServiceTokenRule) implementsAccessGroupReplaceParamsRequire() {
}

type AccessGroupReplaceParamsRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccessGroupReplaceParamsRequireAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccessGroupReplaceParamsRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccessGroupReplaceParamsRequireAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsRequireAccessAnyValidServiceTokenRule) implementsAccessGroupReplaceParamsRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupReplaceParamsRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessGroupReplaceParamsRequireAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessGroupReplaceParamsRequireAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsRequireAccessExternalEvaluationRule) implementsAccessGroupReplaceParamsRequire() {
}

type AccessGroupReplaceParamsRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccessGroupReplaceParamsRequireAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccessGroupReplaceParamsRequireAccessCountryRule struct {
	Geo param.Field[AccessGroupReplaceParamsRequireAccessCountryRuleGeo] `json:"geo,required"`
}

func (r AccessGroupReplaceParamsRequireAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsRequireAccessCountryRule) implementsAccessGroupReplaceParamsRequire() {
}

type AccessGroupReplaceParamsRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccessGroupReplaceParamsRequireAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccessGroupReplaceParamsRequireAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[AccessGroupReplaceParamsRequireAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccessGroupReplaceParamsRequireAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsRequireAccessAuthenticationMethodRule) implementsAccessGroupReplaceParamsRequire() {
}

type AccessGroupReplaceParamsRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccessGroupReplaceParamsRequireAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccessGroupReplaceParamsRequireAccessDevicePostureRule struct {
	DevicePosture param.Field[AccessGroupReplaceParamsRequireAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccessGroupReplaceParamsRequireAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupReplaceParamsRequireAccessDevicePostureRule) implementsAccessGroupReplaceParamsRequire() {
}

type AccessGroupReplaceParamsRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccessGroupReplaceParamsRequireAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessGroupReplaceResponseEnvelope struct {
	Errors   []AccessGroupReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessGroupReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessGroupReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessGroupReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessGroupReplaceResponseEnvelopeJSON    `json:"-"`
}

// accessGroupReplaceResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessGroupReplaceResponseEnvelope]
type accessGroupReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupReplaceResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accessGroupReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// accessGroupReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessGroupReplaceResponseEnvelopeErrors]
type accessGroupReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupReplaceResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accessGroupReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// accessGroupReplaceResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AccessGroupReplaceResponseEnvelopeMessages]
type accessGroupReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessGroupReplaceResponseEnvelopeSuccess bool

const (
	AccessGroupReplaceResponseEnvelopeSuccessTrue AccessGroupReplaceResponseEnvelopeSuccess = true
)
