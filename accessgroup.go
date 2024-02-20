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
	Errors   []AccessGroupUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessGroupUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessGroupUpdateResponse                   `json:"result,required"`
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
