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
	path := fmt.Sprintf("%s/%s/access/groups/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured Access group.
func (r *AccessGroupService) Update(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, body AccessGroupUpdateParams, opts ...option.RequestOption) (res *AccessGroupUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/groups/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes an Access group.
func (r *AccessGroupService) Delete(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, opts ...option.RequestOption) (res *AccessGroupDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/groups/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a new Access group.
func (r *AccessGroupService) AccessGroupsNewAnAccessGroup(ctx context.Context, accountOrZone string, accountOrZoneID string, body AccessGroupAccessGroupsNewAnAccessGroupParams, opts ...option.RequestOption) (res *AccessGroupAccessGroupsNewAnAccessGroupResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/groups", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all Access groups.
func (r *AccessGroupService) AccessGroupsListAccessGroups(ctx context.Context, accountOrZone string, accountOrZoneID string, opts ...option.RequestOption) (res *AccessGroupAccessGroupsListAccessGroupsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/groups", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccessGroupGetResponse struct {
	Errors   []AccessGroupGetResponseError   `json:"errors"`
	Messages []AccessGroupGetResponseMessage `json:"messages"`
	Result   AccessGroupGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessGroupGetResponseSuccess `json:"success"`
	JSON    accessGroupGetResponseJSON    `json:"-"`
}

// accessGroupGetResponseJSON contains the JSON metadata for the struct
// [AccessGroupGetResponse]
type accessGroupGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupGetResponseError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    accessGroupGetResponseErrorJSON `json:"-"`
}

// accessGroupGetResponseErrorJSON contains the JSON metadata for the struct
// [AccessGroupGetResponseError]
type accessGroupGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupGetResponseMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    accessGroupGetResponseMessageJSON `json:"-"`
}

// accessGroupGetResponseMessageJSON contains the JSON metadata for the struct
// [AccessGroupGetResponseMessage]
type accessGroupGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupGetResponseResult struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessGroupGetResponseResultExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessGroupGetResponseResultInclude `json:"include"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	IsDefault []AccessGroupGetResponseResultIsDefault `json:"is_default"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []AccessGroupGetResponseResultRequire `json:"require"`
	UpdatedAt time.Time                             `json:"updated_at" format:"date-time"`
	JSON      accessGroupGetResponseResultJSON      `json:"-"`
}

// accessGroupGetResponseResultJSON contains the JSON metadata for the struct
// [AccessGroupGetResponseResult]
type accessGroupGetResponseResultJSON struct {
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

func (r *AccessGroupGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupGetResponseResultExcludeAccessEmailRule],
// [AccessGroupGetResponseResultExcludeAccessEmailListRule],
// [AccessGroupGetResponseResultExcludeAccessDomainRule],
// [AccessGroupGetResponseResultExcludeAccessEveryoneRule],
// [AccessGroupGetResponseResultExcludeAccessIPRule],
// [AccessGroupGetResponseResultExcludeAccessIPListRule],
// [AccessGroupGetResponseResultExcludeAccessCertificateRule],
// [AccessGroupGetResponseResultExcludeAccessAccessGroupRule],
// [AccessGroupGetResponseResultExcludeAccessAzureGroupRule],
// [AccessGroupGetResponseResultExcludeAccessGitHubOrganizationRule],
// [AccessGroupGetResponseResultExcludeAccessGsuiteGroupRule],
// [AccessGroupGetResponseResultExcludeAccessOktaGroupRule],
// [AccessGroupGetResponseResultExcludeAccessSamlGroupRule],
// [AccessGroupGetResponseResultExcludeAccessServiceTokenRule],
// [AccessGroupGetResponseResultExcludeAccessAnyValidServiceTokenRule],
// [AccessGroupGetResponseResultExcludeAccessExternalEvaluationRule],
// [AccessGroupGetResponseResultExcludeAccessCountryRule],
// [AccessGroupGetResponseResultExcludeAccessAuthenticationMethodRule] or
// [AccessGroupGetResponseResultExcludeAccessDevicePostureRule].
type AccessGroupGetResponseResultExclude interface {
	implementsAccessGroupGetResponseResultExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupGetResponseResultExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupGetResponseResultExcludeAccessEmailRule struct {
	Email AccessGroupGetResponseResultExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupGetResponseResultExcludeAccessEmailRuleJSON  `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessEmailRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseResultExcludeAccessEmailRule]
type accessGroupGetResponseResultExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultExcludeAccessEmailRule) implementsAccessGroupGetResponseResultExclude() {
}

type AccessGroupGetResponseResultExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                      `json:"email,required" format:"email"`
	JSON  accessGroupGetResponseResultExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultExcludeAccessEmailRuleEmail]
type accessGroupGetResponseResultExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupGetResponseResultExcludeAccessEmailListRule struct {
	EmailList AccessGroupGetResponseResultExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupGetResponseResultExcludeAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessEmailListRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseResultExcludeAccessEmailListRule]
type accessGroupGetResponseResultExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultExcludeAccessEmailListRule) implementsAccessGroupGetResponseResultExclude() {
}

type AccessGroupGetResponseResultExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                              `json:"id,required"`
	JSON accessGroupGetResponseResultExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessEmailListRuleEmailListJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseResultExcludeAccessEmailListRuleEmailList]
type accessGroupGetResponseResultExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupGetResponseResultExcludeAccessDomainRule struct {
	EmailDomain AccessGroupGetResponseResultExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupGetResponseResultExcludeAccessDomainRuleJSON        `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessDomainRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseResultExcludeAccessDomainRule]
type accessGroupGetResponseResultExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultExcludeAccessDomainRule) implementsAccessGroupGetResponseResultExclude() {
}

type AccessGroupGetResponseResultExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                             `json:"domain,required"`
	JSON   accessGroupGetResponseResultExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessDomainRuleEmailDomainJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseResultExcludeAccessDomainRuleEmailDomain]
type accessGroupGetResponseResultExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupGetResponseResultExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                               `json:"everyone,required"`
	JSON     accessGroupGetResponseResultExcludeAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessEveryoneRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseResultExcludeAccessEveryoneRule]
type accessGroupGetResponseResultExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultExcludeAccessEveryoneRule) implementsAccessGroupGetResponseResultExclude() {
}

// Matches an IP address block.
type AccessGroupGetResponseResultExcludeAccessIPRule struct {
	IP   AccessGroupGetResponseResultExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupGetResponseResultExcludeAccessIPRuleJSON `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessIPRuleJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseResultExcludeAccessIPRule]
type accessGroupGetResponseResultExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultExcludeAccessIPRule) implementsAccessGroupGetResponseResultExclude() {
}

type AccessGroupGetResponseResultExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                `json:"ip,required"`
	JSON accessGroupGetResponseResultExcludeAccessIPRuleIPJSON `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessIPRuleIPJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseResultExcludeAccessIPRuleIP]
type accessGroupGetResponseResultExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupGetResponseResultExcludeAccessIPListRule struct {
	IPList AccessGroupGetResponseResultExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupGetResponseResultExcludeAccessIPListRuleJSON   `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessIPListRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseResultExcludeAccessIPListRule]
type accessGroupGetResponseResultExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultExcludeAccessIPListRule) implementsAccessGroupGetResponseResultExclude() {
}

type AccessGroupGetResponseResultExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                        `json:"id,required"`
	JSON accessGroupGetResponseResultExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessIPListRuleIPListJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultExcludeAccessIPListRuleIPList]
type accessGroupGetResponseResultExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupGetResponseResultExcludeAccessCertificateRule struct {
	Certificate interface{}                                                  `json:"certificate,required"`
	JSON        accessGroupGetResponseResultExcludeAccessCertificateRuleJSON `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessCertificateRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultExcludeAccessCertificateRule]
type accessGroupGetResponseResultExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultExcludeAccessCertificateRule) implementsAccessGroupGetResponseResultExclude() {
}

// Matches an Access group.
type AccessGroupGetResponseResultExcludeAccessAccessGroupRule struct {
	Group AccessGroupGetResponseResultExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupGetResponseResultExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessAccessGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultExcludeAccessAccessGroupRule]
type accessGroupGetResponseResultExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultExcludeAccessAccessGroupRule) implementsAccessGroupGetResponseResultExclude() {
}

type AccessGroupGetResponseResultExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                            `json:"id,required"`
	JSON accessGroupGetResponseResultExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessAccessGroupRuleGroupJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseResultExcludeAccessAccessGroupRuleGroup]
type accessGroupGetResponseResultExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupGetResponseResultExcludeAccessAzureGroupRule struct {
	AzureAd AccessGroupGetResponseResultExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupGetResponseResultExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultExcludeAccessAzureGroupRule]
type accessGroupGetResponseResultExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultExcludeAccessAzureGroupRule) implementsAccessGroupGetResponseResultExclude() {
}

type AccessGroupGetResponseResultExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                             `json:"connection_id,required"`
	JSON         accessGroupGetResponseResultExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessAzureGroupRuleAzureAdJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseResultExcludeAccessAzureGroupRuleAzureAd]
type accessGroupGetResponseResultExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupGetResponseResultExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupGetResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupGetResponseResultExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessGitHubOrganizationRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseResultExcludeAccessGitHubOrganizationRule]
type accessGroupGetResponseResultExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultExcludeAccessGitHubOrganizationRule) implementsAccessGroupGetResponseResultExclude() {
}

type AccessGroupGetResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                `json:"name,required"`
	JSON accessGroupGetResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupGetResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupGetResponseResultExcludeAccessGsuiteGroupRule struct {
	Gsuite AccessGroupGetResponseResultExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupGetResponseResultExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessGsuiteGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultExcludeAccessGsuiteGroupRule]
type accessGroupGetResponseResultExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultExcludeAccessGsuiteGroupRule) implementsAccessGroupGetResponseResultExclude() {
}

type AccessGroupGetResponseResultExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                             `json:"email,required"`
	JSON  accessGroupGetResponseResultExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessGsuiteGroupRuleGsuiteJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseResultExcludeAccessGsuiteGroupRuleGsuite]
type accessGroupGetResponseResultExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupGetResponseResultExcludeAccessOktaGroupRule struct {
	Okta AccessGroupGetResponseResultExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupGetResponseResultExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseResultExcludeAccessOktaGroupRule]
type accessGroupGetResponseResultExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultExcludeAccessOktaGroupRule) implementsAccessGroupGetResponseResultExclude() {
}

type AccessGroupGetResponseResultExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                         `json:"email,required"`
	JSON  accessGroupGetResponseResultExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessOktaGroupRuleOktaJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultExcludeAccessOktaGroupRuleOkta]
type accessGroupGetResponseResultExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupGetResponseResultExcludeAccessSamlGroupRule struct {
	Saml AccessGroupGetResponseResultExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupGetResponseResultExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseResultExcludeAccessSamlGroupRule]
type accessGroupGetResponseResultExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultExcludeAccessSamlGroupRule) implementsAccessGroupGetResponseResultExclude() {
}

type AccessGroupGetResponseResultExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                         `json:"attribute_value,required"`
	JSON           accessGroupGetResponseResultExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessSamlGroupRuleSamlJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultExcludeAccessSamlGroupRuleSaml]
type accessGroupGetResponseResultExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupGetResponseResultExcludeAccessServiceTokenRule struct {
	ServiceToken AccessGroupGetResponseResultExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupGetResponseResultExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessServiceTokenRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultExcludeAccessServiceTokenRule]
type accessGroupGetResponseResultExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultExcludeAccessServiceTokenRule) implementsAccessGroupGetResponseResultExclude() {
}

type AccessGroupGetResponseResultExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                    `json:"token_id,required"`
	JSON    accessGroupGetResponseResultExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseResultExcludeAccessServiceTokenRuleServiceToken]
type accessGroupGetResponseResultExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupGetResponseResultExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                           `json:"any_valid_service_token,required"`
	JSON                 accessGroupGetResponseResultExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessAnyValidServiceTokenRuleJSON contains
// the JSON metadata for the struct
// [AccessGroupGetResponseResultExcludeAccessAnyValidServiceTokenRule]
type accessGroupGetResponseResultExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultExcludeAccessAnyValidServiceTokenRule) implementsAccessGroupGetResponseResultExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupGetResponseResultExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupGetResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupGetResponseResultExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessExternalEvaluationRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseResultExcludeAccessExternalEvaluationRule]
type accessGroupGetResponseResultExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultExcludeAccessExternalEvaluationRule) implementsAccessGroupGetResponseResultExclude() {
}

type AccessGroupGetResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                `json:"keys_url,required"`
	JSON    accessGroupGetResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupGetResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupGetResponseResultExcludeAccessCountryRule struct {
	Geo  AccessGroupGetResponseResultExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupGetResponseResultExcludeAccessCountryRuleJSON `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessCountryRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseResultExcludeAccessCountryRule]
type accessGroupGetResponseResultExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultExcludeAccessCountryRule) implementsAccessGroupGetResponseResultExclude() {
}

type AccessGroupGetResponseResultExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                      `json:"country_code,required"`
	JSON        accessGroupGetResponseResultExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultExcludeAccessCountryRuleGeo]
type accessGroupGetResponseResultExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupGetResponseResultExcludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupGetResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupGetResponseResultExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessAuthenticationMethodRuleJSON contains
// the JSON metadata for the struct
// [AccessGroupGetResponseResultExcludeAccessAuthenticationMethodRule]
type accessGroupGetResponseResultExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultExcludeAccessAuthenticationMethodRule) implementsAccessGroupGetResponseResultExclude() {
}

type AccessGroupGetResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                          `json:"auth_method,required"`
	JSON       accessGroupGetResponseResultExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod]
type accessGroupGetResponseResultExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupGetResponseResultExcludeAccessDevicePostureRule struct {
	DevicePosture AccessGroupGetResponseResultExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupGetResponseResultExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessDevicePostureRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultExcludeAccessDevicePostureRule]
type accessGroupGetResponseResultExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultExcludeAccessDevicePostureRule) implementsAccessGroupGetResponseResultExclude() {
}

type AccessGroupGetResponseResultExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                      `json:"integration_uid,required"`
	JSON           accessGroupGetResponseResultExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupGetResponseResultExcludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseResultExcludeAccessDevicePostureRuleDevicePosture]
type accessGroupGetResponseResultExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupGetResponseResultIncludeAccessEmailRule],
// [AccessGroupGetResponseResultIncludeAccessEmailListRule],
// [AccessGroupGetResponseResultIncludeAccessDomainRule],
// [AccessGroupGetResponseResultIncludeAccessEveryoneRule],
// [AccessGroupGetResponseResultIncludeAccessIPRule],
// [AccessGroupGetResponseResultIncludeAccessIPListRule],
// [AccessGroupGetResponseResultIncludeAccessCertificateRule],
// [AccessGroupGetResponseResultIncludeAccessAccessGroupRule],
// [AccessGroupGetResponseResultIncludeAccessAzureGroupRule],
// [AccessGroupGetResponseResultIncludeAccessGitHubOrganizationRule],
// [AccessGroupGetResponseResultIncludeAccessGsuiteGroupRule],
// [AccessGroupGetResponseResultIncludeAccessOktaGroupRule],
// [AccessGroupGetResponseResultIncludeAccessSamlGroupRule],
// [AccessGroupGetResponseResultIncludeAccessServiceTokenRule],
// [AccessGroupGetResponseResultIncludeAccessAnyValidServiceTokenRule],
// [AccessGroupGetResponseResultIncludeAccessExternalEvaluationRule],
// [AccessGroupGetResponseResultIncludeAccessCountryRule],
// [AccessGroupGetResponseResultIncludeAccessAuthenticationMethodRule] or
// [AccessGroupGetResponseResultIncludeAccessDevicePostureRule].
type AccessGroupGetResponseResultInclude interface {
	implementsAccessGroupGetResponseResultInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupGetResponseResultInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupGetResponseResultIncludeAccessEmailRule struct {
	Email AccessGroupGetResponseResultIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupGetResponseResultIncludeAccessEmailRuleJSON  `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessEmailRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseResultIncludeAccessEmailRule]
type accessGroupGetResponseResultIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIncludeAccessEmailRule) implementsAccessGroupGetResponseResultInclude() {
}

type AccessGroupGetResponseResultIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                      `json:"email,required" format:"email"`
	JSON  accessGroupGetResponseResultIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultIncludeAccessEmailRuleEmail]
type accessGroupGetResponseResultIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupGetResponseResultIncludeAccessEmailListRule struct {
	EmailList AccessGroupGetResponseResultIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupGetResponseResultIncludeAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessEmailListRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseResultIncludeAccessEmailListRule]
type accessGroupGetResponseResultIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIncludeAccessEmailListRule) implementsAccessGroupGetResponseResultInclude() {
}

type AccessGroupGetResponseResultIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                              `json:"id,required"`
	JSON accessGroupGetResponseResultIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessEmailListRuleEmailListJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseResultIncludeAccessEmailListRuleEmailList]
type accessGroupGetResponseResultIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupGetResponseResultIncludeAccessDomainRule struct {
	EmailDomain AccessGroupGetResponseResultIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupGetResponseResultIncludeAccessDomainRuleJSON        `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessDomainRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseResultIncludeAccessDomainRule]
type accessGroupGetResponseResultIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIncludeAccessDomainRule) implementsAccessGroupGetResponseResultInclude() {
}

type AccessGroupGetResponseResultIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                             `json:"domain,required"`
	JSON   accessGroupGetResponseResultIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessDomainRuleEmailDomainJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseResultIncludeAccessDomainRuleEmailDomain]
type accessGroupGetResponseResultIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupGetResponseResultIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                               `json:"everyone,required"`
	JSON     accessGroupGetResponseResultIncludeAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessEveryoneRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseResultIncludeAccessEveryoneRule]
type accessGroupGetResponseResultIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIncludeAccessEveryoneRule) implementsAccessGroupGetResponseResultInclude() {
}

// Matches an IP address block.
type AccessGroupGetResponseResultIncludeAccessIPRule struct {
	IP   AccessGroupGetResponseResultIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupGetResponseResultIncludeAccessIPRuleJSON `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessIPRuleJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseResultIncludeAccessIPRule]
type accessGroupGetResponseResultIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIncludeAccessIPRule) implementsAccessGroupGetResponseResultInclude() {
}

type AccessGroupGetResponseResultIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                `json:"ip,required"`
	JSON accessGroupGetResponseResultIncludeAccessIPRuleIPJSON `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessIPRuleIPJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseResultIncludeAccessIPRuleIP]
type accessGroupGetResponseResultIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupGetResponseResultIncludeAccessIPListRule struct {
	IPList AccessGroupGetResponseResultIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupGetResponseResultIncludeAccessIPListRuleJSON   `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessIPListRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseResultIncludeAccessIPListRule]
type accessGroupGetResponseResultIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIncludeAccessIPListRule) implementsAccessGroupGetResponseResultInclude() {
}

type AccessGroupGetResponseResultIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                        `json:"id,required"`
	JSON accessGroupGetResponseResultIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessIPListRuleIPListJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultIncludeAccessIPListRuleIPList]
type accessGroupGetResponseResultIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupGetResponseResultIncludeAccessCertificateRule struct {
	Certificate interface{}                                                  `json:"certificate,required"`
	JSON        accessGroupGetResponseResultIncludeAccessCertificateRuleJSON `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessCertificateRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultIncludeAccessCertificateRule]
type accessGroupGetResponseResultIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIncludeAccessCertificateRule) implementsAccessGroupGetResponseResultInclude() {
}

// Matches an Access group.
type AccessGroupGetResponseResultIncludeAccessAccessGroupRule struct {
	Group AccessGroupGetResponseResultIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupGetResponseResultIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessAccessGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultIncludeAccessAccessGroupRule]
type accessGroupGetResponseResultIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIncludeAccessAccessGroupRule) implementsAccessGroupGetResponseResultInclude() {
}

type AccessGroupGetResponseResultIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                            `json:"id,required"`
	JSON accessGroupGetResponseResultIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessAccessGroupRuleGroupJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseResultIncludeAccessAccessGroupRuleGroup]
type accessGroupGetResponseResultIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupGetResponseResultIncludeAccessAzureGroupRule struct {
	AzureAd AccessGroupGetResponseResultIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupGetResponseResultIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultIncludeAccessAzureGroupRule]
type accessGroupGetResponseResultIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIncludeAccessAzureGroupRule) implementsAccessGroupGetResponseResultInclude() {
}

type AccessGroupGetResponseResultIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                             `json:"connection_id,required"`
	JSON         accessGroupGetResponseResultIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessAzureGroupRuleAzureAdJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseResultIncludeAccessAzureGroupRuleAzureAd]
type accessGroupGetResponseResultIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupGetResponseResultIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupGetResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupGetResponseResultIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessGitHubOrganizationRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseResultIncludeAccessGitHubOrganizationRule]
type accessGroupGetResponseResultIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIncludeAccessGitHubOrganizationRule) implementsAccessGroupGetResponseResultInclude() {
}

type AccessGroupGetResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                `json:"name,required"`
	JSON accessGroupGetResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupGetResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupGetResponseResultIncludeAccessGsuiteGroupRule struct {
	Gsuite AccessGroupGetResponseResultIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupGetResponseResultIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessGsuiteGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultIncludeAccessGsuiteGroupRule]
type accessGroupGetResponseResultIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIncludeAccessGsuiteGroupRule) implementsAccessGroupGetResponseResultInclude() {
}

type AccessGroupGetResponseResultIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                             `json:"email,required"`
	JSON  accessGroupGetResponseResultIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessGsuiteGroupRuleGsuiteJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseResultIncludeAccessGsuiteGroupRuleGsuite]
type accessGroupGetResponseResultIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupGetResponseResultIncludeAccessOktaGroupRule struct {
	Okta AccessGroupGetResponseResultIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupGetResponseResultIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseResultIncludeAccessOktaGroupRule]
type accessGroupGetResponseResultIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIncludeAccessOktaGroupRule) implementsAccessGroupGetResponseResultInclude() {
}

type AccessGroupGetResponseResultIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                         `json:"email,required"`
	JSON  accessGroupGetResponseResultIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessOktaGroupRuleOktaJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultIncludeAccessOktaGroupRuleOkta]
type accessGroupGetResponseResultIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupGetResponseResultIncludeAccessSamlGroupRule struct {
	Saml AccessGroupGetResponseResultIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupGetResponseResultIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseResultIncludeAccessSamlGroupRule]
type accessGroupGetResponseResultIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIncludeAccessSamlGroupRule) implementsAccessGroupGetResponseResultInclude() {
}

type AccessGroupGetResponseResultIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                         `json:"attribute_value,required"`
	JSON           accessGroupGetResponseResultIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessSamlGroupRuleSamlJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultIncludeAccessSamlGroupRuleSaml]
type accessGroupGetResponseResultIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupGetResponseResultIncludeAccessServiceTokenRule struct {
	ServiceToken AccessGroupGetResponseResultIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupGetResponseResultIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessServiceTokenRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultIncludeAccessServiceTokenRule]
type accessGroupGetResponseResultIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIncludeAccessServiceTokenRule) implementsAccessGroupGetResponseResultInclude() {
}

type AccessGroupGetResponseResultIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                    `json:"token_id,required"`
	JSON    accessGroupGetResponseResultIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseResultIncludeAccessServiceTokenRuleServiceToken]
type accessGroupGetResponseResultIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupGetResponseResultIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                           `json:"any_valid_service_token,required"`
	JSON                 accessGroupGetResponseResultIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessAnyValidServiceTokenRuleJSON contains
// the JSON metadata for the struct
// [AccessGroupGetResponseResultIncludeAccessAnyValidServiceTokenRule]
type accessGroupGetResponseResultIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIncludeAccessAnyValidServiceTokenRule) implementsAccessGroupGetResponseResultInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupGetResponseResultIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupGetResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupGetResponseResultIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessExternalEvaluationRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseResultIncludeAccessExternalEvaluationRule]
type accessGroupGetResponseResultIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIncludeAccessExternalEvaluationRule) implementsAccessGroupGetResponseResultInclude() {
}

type AccessGroupGetResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                `json:"keys_url,required"`
	JSON    accessGroupGetResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupGetResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupGetResponseResultIncludeAccessCountryRule struct {
	Geo  AccessGroupGetResponseResultIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupGetResponseResultIncludeAccessCountryRuleJSON `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessCountryRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseResultIncludeAccessCountryRule]
type accessGroupGetResponseResultIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIncludeAccessCountryRule) implementsAccessGroupGetResponseResultInclude() {
}

type AccessGroupGetResponseResultIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                      `json:"country_code,required"`
	JSON        accessGroupGetResponseResultIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultIncludeAccessCountryRuleGeo]
type accessGroupGetResponseResultIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupGetResponseResultIncludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupGetResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupGetResponseResultIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessAuthenticationMethodRuleJSON contains
// the JSON metadata for the struct
// [AccessGroupGetResponseResultIncludeAccessAuthenticationMethodRule]
type accessGroupGetResponseResultIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIncludeAccessAuthenticationMethodRule) implementsAccessGroupGetResponseResultInclude() {
}

type AccessGroupGetResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                          `json:"auth_method,required"`
	JSON       accessGroupGetResponseResultIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod]
type accessGroupGetResponseResultIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupGetResponseResultIncludeAccessDevicePostureRule struct {
	DevicePosture AccessGroupGetResponseResultIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupGetResponseResultIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessDevicePostureRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultIncludeAccessDevicePostureRule]
type accessGroupGetResponseResultIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIncludeAccessDevicePostureRule) implementsAccessGroupGetResponseResultInclude() {
}

type AccessGroupGetResponseResultIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                      `json:"integration_uid,required"`
	JSON           accessGroupGetResponseResultIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupGetResponseResultIncludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseResultIncludeAccessDevicePostureRuleDevicePosture]
type accessGroupGetResponseResultIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupGetResponseResultIsDefaultAccessEmailRule],
// [AccessGroupGetResponseResultIsDefaultAccessEmailListRule],
// [AccessGroupGetResponseResultIsDefaultAccessDomainRule],
// [AccessGroupGetResponseResultIsDefaultAccessEveryoneRule],
// [AccessGroupGetResponseResultIsDefaultAccessIPRule],
// [AccessGroupGetResponseResultIsDefaultAccessIPListRule],
// [AccessGroupGetResponseResultIsDefaultAccessCertificateRule],
// [AccessGroupGetResponseResultIsDefaultAccessAccessGroupRule],
// [AccessGroupGetResponseResultIsDefaultAccessAzureGroupRule],
// [AccessGroupGetResponseResultIsDefaultAccessGitHubOrganizationRule],
// [AccessGroupGetResponseResultIsDefaultAccessGsuiteGroupRule],
// [AccessGroupGetResponseResultIsDefaultAccessOktaGroupRule],
// [AccessGroupGetResponseResultIsDefaultAccessSamlGroupRule],
// [AccessGroupGetResponseResultIsDefaultAccessServiceTokenRule],
// [AccessGroupGetResponseResultIsDefaultAccessAnyValidServiceTokenRule],
// [AccessGroupGetResponseResultIsDefaultAccessExternalEvaluationRule],
// [AccessGroupGetResponseResultIsDefaultAccessCountryRule],
// [AccessGroupGetResponseResultIsDefaultAccessAuthenticationMethodRule] or
// [AccessGroupGetResponseResultIsDefaultAccessDevicePostureRule].
type AccessGroupGetResponseResultIsDefault interface {
	implementsAccessGroupGetResponseResultIsDefault()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupGetResponseResultIsDefault)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupGetResponseResultIsDefaultAccessEmailRule struct {
	Email AccessGroupGetResponseResultIsDefaultAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupGetResponseResultIsDefaultAccessEmailRuleJSON  `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessEmailRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseResultIsDefaultAccessEmailRule]
type accessGroupGetResponseResultIsDefaultAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIsDefaultAccessEmailRule) implementsAccessGroupGetResponseResultIsDefault() {
}

type AccessGroupGetResponseResultIsDefaultAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                        `json:"email,required" format:"email"`
	JSON  accessGroupGetResponseResultIsDefaultAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultIsDefaultAccessEmailRuleEmail]
type accessGroupGetResponseResultIsDefaultAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupGetResponseResultIsDefaultAccessEmailListRule struct {
	EmailList AccessGroupGetResponseResultIsDefaultAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupGetResponseResultIsDefaultAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessEmailListRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultIsDefaultAccessEmailListRule]
type accessGroupGetResponseResultIsDefaultAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIsDefaultAccessEmailListRule) implementsAccessGroupGetResponseResultIsDefault() {
}

type AccessGroupGetResponseResultIsDefaultAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                `json:"id,required"`
	JSON accessGroupGetResponseResultIsDefaultAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessEmailListRuleEmailListJSON contains
// the JSON metadata for the struct
// [AccessGroupGetResponseResultIsDefaultAccessEmailListRuleEmailList]
type accessGroupGetResponseResultIsDefaultAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupGetResponseResultIsDefaultAccessDomainRule struct {
	EmailDomain AccessGroupGetResponseResultIsDefaultAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupGetResponseResultIsDefaultAccessDomainRuleJSON        `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessDomainRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseResultIsDefaultAccessDomainRule]
type accessGroupGetResponseResultIsDefaultAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIsDefaultAccessDomainRule) implementsAccessGroupGetResponseResultIsDefault() {
}

type AccessGroupGetResponseResultIsDefaultAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                               `json:"domain,required"`
	JSON   accessGroupGetResponseResultIsDefaultAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessDomainRuleEmailDomainJSON contains
// the JSON metadata for the struct
// [AccessGroupGetResponseResultIsDefaultAccessDomainRuleEmailDomain]
type accessGroupGetResponseResultIsDefaultAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupGetResponseResultIsDefaultAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                 `json:"everyone,required"`
	JSON     accessGroupGetResponseResultIsDefaultAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessEveryoneRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultIsDefaultAccessEveryoneRule]
type accessGroupGetResponseResultIsDefaultAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIsDefaultAccessEveryoneRule) implementsAccessGroupGetResponseResultIsDefault() {
}

// Matches an IP address block.
type AccessGroupGetResponseResultIsDefaultAccessIPRule struct {
	IP   AccessGroupGetResponseResultIsDefaultAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupGetResponseResultIsDefaultAccessIPRuleJSON `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessIPRuleJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseResultIsDefaultAccessIPRule]
type accessGroupGetResponseResultIsDefaultAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIsDefaultAccessIPRule) implementsAccessGroupGetResponseResultIsDefault() {
}

type AccessGroupGetResponseResultIsDefaultAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                  `json:"ip,required"`
	JSON accessGroupGetResponseResultIsDefaultAccessIPRuleIPJSON `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessIPRuleIPJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseResultIsDefaultAccessIPRuleIP]
type accessGroupGetResponseResultIsDefaultAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupGetResponseResultIsDefaultAccessIPListRule struct {
	IPList AccessGroupGetResponseResultIsDefaultAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupGetResponseResultIsDefaultAccessIPListRuleJSON   `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessIPListRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseResultIsDefaultAccessIPListRule]
type accessGroupGetResponseResultIsDefaultAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIsDefaultAccessIPListRule) implementsAccessGroupGetResponseResultIsDefault() {
}

type AccessGroupGetResponseResultIsDefaultAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                          `json:"id,required"`
	JSON accessGroupGetResponseResultIsDefaultAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseResultIsDefaultAccessIPListRuleIPList]
type accessGroupGetResponseResultIsDefaultAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupGetResponseResultIsDefaultAccessCertificateRule struct {
	Certificate interface{}                                                    `json:"certificate,required"`
	JSON        accessGroupGetResponseResultIsDefaultAccessCertificateRuleJSON `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessCertificateRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultIsDefaultAccessCertificateRule]
type accessGroupGetResponseResultIsDefaultAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIsDefaultAccessCertificateRule) implementsAccessGroupGetResponseResultIsDefault() {
}

// Matches an Access group.
type AccessGroupGetResponseResultIsDefaultAccessAccessGroupRule struct {
	Group AccessGroupGetResponseResultIsDefaultAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupGetResponseResultIsDefaultAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessAccessGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultIsDefaultAccessAccessGroupRule]
type accessGroupGetResponseResultIsDefaultAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIsDefaultAccessAccessGroupRule) implementsAccessGroupGetResponseResultIsDefault() {
}

type AccessGroupGetResponseResultIsDefaultAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                              `json:"id,required"`
	JSON accessGroupGetResponseResultIsDefaultAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessAccessGroupRuleGroupJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseResultIsDefaultAccessAccessGroupRuleGroup]
type accessGroupGetResponseResultIsDefaultAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupGetResponseResultIsDefaultAccessAzureGroupRule struct {
	AzureAd AccessGroupGetResponseResultIsDefaultAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupGetResponseResultIsDefaultAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultIsDefaultAccessAzureGroupRule]
type accessGroupGetResponseResultIsDefaultAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIsDefaultAccessAzureGroupRule) implementsAccessGroupGetResponseResultIsDefault() {
}

type AccessGroupGetResponseResultIsDefaultAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                               `json:"connection_id,required"`
	JSON         accessGroupGetResponseResultIsDefaultAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessAzureGroupRuleAzureAdJSON contains
// the JSON metadata for the struct
// [AccessGroupGetResponseResultIsDefaultAccessAzureGroupRuleAzureAd]
type accessGroupGetResponseResultIsDefaultAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupGetResponseResultIsDefaultAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupGetResponseResultIsDefaultAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupGetResponseResultIsDefaultAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessGitHubOrganizationRuleJSON contains
// the JSON metadata for the struct
// [AccessGroupGetResponseResultIsDefaultAccessGitHubOrganizationRule]
type accessGroupGetResponseResultIsDefaultAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIsDefaultAccessGitHubOrganizationRule) implementsAccessGroupGetResponseResultIsDefault() {
}

type AccessGroupGetResponseResultIsDefaultAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                  `json:"name,required"`
	JSON accessGroupGetResponseResultIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseResultIsDefaultAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupGetResponseResultIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupGetResponseResultIsDefaultAccessGsuiteGroupRule struct {
	Gsuite AccessGroupGetResponseResultIsDefaultAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupGetResponseResultIsDefaultAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessGsuiteGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultIsDefaultAccessGsuiteGroupRule]
type accessGroupGetResponseResultIsDefaultAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIsDefaultAccessGsuiteGroupRule) implementsAccessGroupGetResponseResultIsDefault() {
}

type AccessGroupGetResponseResultIsDefaultAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                               `json:"email,required"`
	JSON  accessGroupGetResponseResultIsDefaultAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessGsuiteGroupRuleGsuiteJSON contains
// the JSON metadata for the struct
// [AccessGroupGetResponseResultIsDefaultAccessGsuiteGroupRuleGsuite]
type accessGroupGetResponseResultIsDefaultAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupGetResponseResultIsDefaultAccessOktaGroupRule struct {
	Okta AccessGroupGetResponseResultIsDefaultAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupGetResponseResultIsDefaultAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultIsDefaultAccessOktaGroupRule]
type accessGroupGetResponseResultIsDefaultAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIsDefaultAccessOktaGroupRule) implementsAccessGroupGetResponseResultIsDefault() {
}

type AccessGroupGetResponseResultIsDefaultAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                           `json:"email,required"`
	JSON  accessGroupGetResponseResultIsDefaultAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessOktaGroupRuleOktaJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseResultIsDefaultAccessOktaGroupRuleOkta]
type accessGroupGetResponseResultIsDefaultAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupGetResponseResultIsDefaultAccessSamlGroupRule struct {
	Saml AccessGroupGetResponseResultIsDefaultAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupGetResponseResultIsDefaultAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultIsDefaultAccessSamlGroupRule]
type accessGroupGetResponseResultIsDefaultAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIsDefaultAccessSamlGroupRule) implementsAccessGroupGetResponseResultIsDefault() {
}

type AccessGroupGetResponseResultIsDefaultAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                           `json:"attribute_value,required"`
	JSON           accessGroupGetResponseResultIsDefaultAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessSamlGroupRuleSamlJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseResultIsDefaultAccessSamlGroupRuleSaml]
type accessGroupGetResponseResultIsDefaultAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupGetResponseResultIsDefaultAccessServiceTokenRule struct {
	ServiceToken AccessGroupGetResponseResultIsDefaultAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupGetResponseResultIsDefaultAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseResultIsDefaultAccessServiceTokenRule]
type accessGroupGetResponseResultIsDefaultAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIsDefaultAccessServiceTokenRule) implementsAccessGroupGetResponseResultIsDefault() {
}

type AccessGroupGetResponseResultIsDefaultAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                      `json:"token_id,required"`
	JSON    accessGroupGetResponseResultIsDefaultAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseResultIsDefaultAccessServiceTokenRuleServiceToken]
type accessGroupGetResponseResultIsDefaultAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupGetResponseResultIsDefaultAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                             `json:"any_valid_service_token,required"`
	JSON                 accessGroupGetResponseResultIsDefaultAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessAnyValidServiceTokenRuleJSON contains
// the JSON metadata for the struct
// [AccessGroupGetResponseResultIsDefaultAccessAnyValidServiceTokenRule]
type accessGroupGetResponseResultIsDefaultAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIsDefaultAccessAnyValidServiceTokenRule) implementsAccessGroupGetResponseResultIsDefault() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupGetResponseResultIsDefaultAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupGetResponseResultIsDefaultAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupGetResponseResultIsDefaultAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessExternalEvaluationRuleJSON contains
// the JSON metadata for the struct
// [AccessGroupGetResponseResultIsDefaultAccessExternalEvaluationRule]
type accessGroupGetResponseResultIsDefaultAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIsDefaultAccessExternalEvaluationRule) implementsAccessGroupGetResponseResultIsDefault() {
}

type AccessGroupGetResponseResultIsDefaultAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                  `json:"keys_url,required"`
	JSON    accessGroupGetResponseResultIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseResultIsDefaultAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupGetResponseResultIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupGetResponseResultIsDefaultAccessCountryRule struct {
	Geo  AccessGroupGetResponseResultIsDefaultAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupGetResponseResultIsDefaultAccessCountryRuleJSON `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessCountryRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseResultIsDefaultAccessCountryRule]
type accessGroupGetResponseResultIsDefaultAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIsDefaultAccessCountryRule) implementsAccessGroupGetResponseResultIsDefault() {
}

type AccessGroupGetResponseResultIsDefaultAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                        `json:"country_code,required"`
	JSON        accessGroupGetResponseResultIsDefaultAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultIsDefaultAccessCountryRuleGeo]
type accessGroupGetResponseResultIsDefaultAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupGetResponseResultIsDefaultAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupGetResponseResultIsDefaultAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupGetResponseResultIsDefaultAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessAuthenticationMethodRuleJSON contains
// the JSON metadata for the struct
// [AccessGroupGetResponseResultIsDefaultAccessAuthenticationMethodRule]
type accessGroupGetResponseResultIsDefaultAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIsDefaultAccessAuthenticationMethodRule) implementsAccessGroupGetResponseResultIsDefault() {
}

type AccessGroupGetResponseResultIsDefaultAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                            `json:"auth_method,required"`
	JSON       accessGroupGetResponseResultIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseResultIsDefaultAccessAuthenticationMethodRuleAuthMethod]
type accessGroupGetResponseResultIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupGetResponseResultIsDefaultAccessDevicePostureRule struct {
	DevicePosture AccessGroupGetResponseResultIsDefaultAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupGetResponseResultIsDefaultAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessDevicePostureRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseResultIsDefaultAccessDevicePostureRule]
type accessGroupGetResponseResultIsDefaultAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultIsDefaultAccessDevicePostureRule) implementsAccessGroupGetResponseResultIsDefault() {
}

type AccessGroupGetResponseResultIsDefaultAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                        `json:"integration_uid,required"`
	JSON           accessGroupGetResponseResultIsDefaultAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupGetResponseResultIsDefaultAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseResultIsDefaultAccessDevicePostureRuleDevicePosture]
type accessGroupGetResponseResultIsDefaultAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultIsDefaultAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupGetResponseResultRequireAccessEmailRule],
// [AccessGroupGetResponseResultRequireAccessEmailListRule],
// [AccessGroupGetResponseResultRequireAccessDomainRule],
// [AccessGroupGetResponseResultRequireAccessEveryoneRule],
// [AccessGroupGetResponseResultRequireAccessIPRule],
// [AccessGroupGetResponseResultRequireAccessIPListRule],
// [AccessGroupGetResponseResultRequireAccessCertificateRule],
// [AccessGroupGetResponseResultRequireAccessAccessGroupRule],
// [AccessGroupGetResponseResultRequireAccessAzureGroupRule],
// [AccessGroupGetResponseResultRequireAccessGitHubOrganizationRule],
// [AccessGroupGetResponseResultRequireAccessGsuiteGroupRule],
// [AccessGroupGetResponseResultRequireAccessOktaGroupRule],
// [AccessGroupGetResponseResultRequireAccessSamlGroupRule],
// [AccessGroupGetResponseResultRequireAccessServiceTokenRule],
// [AccessGroupGetResponseResultRequireAccessAnyValidServiceTokenRule],
// [AccessGroupGetResponseResultRequireAccessExternalEvaluationRule],
// [AccessGroupGetResponseResultRequireAccessCountryRule],
// [AccessGroupGetResponseResultRequireAccessAuthenticationMethodRule] or
// [AccessGroupGetResponseResultRequireAccessDevicePostureRule].
type AccessGroupGetResponseResultRequire interface {
	implementsAccessGroupGetResponseResultRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupGetResponseResultRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupGetResponseResultRequireAccessEmailRule struct {
	Email AccessGroupGetResponseResultRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupGetResponseResultRequireAccessEmailRuleJSON  `json:"-"`
}

// accessGroupGetResponseResultRequireAccessEmailRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseResultRequireAccessEmailRule]
type accessGroupGetResponseResultRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultRequireAccessEmailRule) implementsAccessGroupGetResponseResultRequire() {
}

type AccessGroupGetResponseResultRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                      `json:"email,required" format:"email"`
	JSON  accessGroupGetResponseResultRequireAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupGetResponseResultRequireAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultRequireAccessEmailRuleEmail]
type accessGroupGetResponseResultRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupGetResponseResultRequireAccessEmailListRule struct {
	EmailList AccessGroupGetResponseResultRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupGetResponseResultRequireAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupGetResponseResultRequireAccessEmailListRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseResultRequireAccessEmailListRule]
type accessGroupGetResponseResultRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultRequireAccessEmailListRule) implementsAccessGroupGetResponseResultRequire() {
}

type AccessGroupGetResponseResultRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                              `json:"id,required"`
	JSON accessGroupGetResponseResultRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupGetResponseResultRequireAccessEmailListRuleEmailListJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseResultRequireAccessEmailListRuleEmailList]
type accessGroupGetResponseResultRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupGetResponseResultRequireAccessDomainRule struct {
	EmailDomain AccessGroupGetResponseResultRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupGetResponseResultRequireAccessDomainRuleJSON        `json:"-"`
}

// accessGroupGetResponseResultRequireAccessDomainRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseResultRequireAccessDomainRule]
type accessGroupGetResponseResultRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultRequireAccessDomainRule) implementsAccessGroupGetResponseResultRequire() {
}

type AccessGroupGetResponseResultRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                             `json:"domain,required"`
	JSON   accessGroupGetResponseResultRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupGetResponseResultRequireAccessDomainRuleEmailDomainJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseResultRequireAccessDomainRuleEmailDomain]
type accessGroupGetResponseResultRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupGetResponseResultRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                               `json:"everyone,required"`
	JSON     accessGroupGetResponseResultRequireAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupGetResponseResultRequireAccessEveryoneRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseResultRequireAccessEveryoneRule]
type accessGroupGetResponseResultRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultRequireAccessEveryoneRule) implementsAccessGroupGetResponseResultRequire() {
}

// Matches an IP address block.
type AccessGroupGetResponseResultRequireAccessIPRule struct {
	IP   AccessGroupGetResponseResultRequireAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupGetResponseResultRequireAccessIPRuleJSON `json:"-"`
}

// accessGroupGetResponseResultRequireAccessIPRuleJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseResultRequireAccessIPRule]
type accessGroupGetResponseResultRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultRequireAccessIPRule) implementsAccessGroupGetResponseResultRequire() {
}

type AccessGroupGetResponseResultRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                `json:"ip,required"`
	JSON accessGroupGetResponseResultRequireAccessIPRuleIPJSON `json:"-"`
}

// accessGroupGetResponseResultRequireAccessIPRuleIPJSON contains the JSON metadata
// for the struct [AccessGroupGetResponseResultRequireAccessIPRuleIP]
type accessGroupGetResponseResultRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupGetResponseResultRequireAccessIPListRule struct {
	IPList AccessGroupGetResponseResultRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupGetResponseResultRequireAccessIPListRuleJSON   `json:"-"`
}

// accessGroupGetResponseResultRequireAccessIPListRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseResultRequireAccessIPListRule]
type accessGroupGetResponseResultRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultRequireAccessIPListRule) implementsAccessGroupGetResponseResultRequire() {
}

type AccessGroupGetResponseResultRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                        `json:"id,required"`
	JSON accessGroupGetResponseResultRequireAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupGetResponseResultRequireAccessIPListRuleIPListJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultRequireAccessIPListRuleIPList]
type accessGroupGetResponseResultRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupGetResponseResultRequireAccessCertificateRule struct {
	Certificate interface{}                                                  `json:"certificate,required"`
	JSON        accessGroupGetResponseResultRequireAccessCertificateRuleJSON `json:"-"`
}

// accessGroupGetResponseResultRequireAccessCertificateRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultRequireAccessCertificateRule]
type accessGroupGetResponseResultRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultRequireAccessCertificateRule) implementsAccessGroupGetResponseResultRequire() {
}

// Matches an Access group.
type AccessGroupGetResponseResultRequireAccessAccessGroupRule struct {
	Group AccessGroupGetResponseResultRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupGetResponseResultRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupGetResponseResultRequireAccessAccessGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultRequireAccessAccessGroupRule]
type accessGroupGetResponseResultRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultRequireAccessAccessGroupRule) implementsAccessGroupGetResponseResultRequire() {
}

type AccessGroupGetResponseResultRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                            `json:"id,required"`
	JSON accessGroupGetResponseResultRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupGetResponseResultRequireAccessAccessGroupRuleGroupJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseResultRequireAccessAccessGroupRuleGroup]
type accessGroupGetResponseResultRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupGetResponseResultRequireAccessAzureGroupRule struct {
	AzureAd AccessGroupGetResponseResultRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupGetResponseResultRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupGetResponseResultRequireAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultRequireAccessAzureGroupRule]
type accessGroupGetResponseResultRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultRequireAccessAzureGroupRule) implementsAccessGroupGetResponseResultRequire() {
}

type AccessGroupGetResponseResultRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                             `json:"connection_id,required"`
	JSON         accessGroupGetResponseResultRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupGetResponseResultRequireAccessAzureGroupRuleAzureAdJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseResultRequireAccessAzureGroupRuleAzureAd]
type accessGroupGetResponseResultRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupGetResponseResultRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupGetResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupGetResponseResultRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupGetResponseResultRequireAccessGitHubOrganizationRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseResultRequireAccessGitHubOrganizationRule]
type accessGroupGetResponseResultRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultRequireAccessGitHubOrganizationRule) implementsAccessGroupGetResponseResultRequire() {
}

type AccessGroupGetResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                `json:"name,required"`
	JSON accessGroupGetResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupGetResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupGetResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupGetResponseResultRequireAccessGsuiteGroupRule struct {
	Gsuite AccessGroupGetResponseResultRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupGetResponseResultRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupGetResponseResultRequireAccessGsuiteGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultRequireAccessGsuiteGroupRule]
type accessGroupGetResponseResultRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultRequireAccessGsuiteGroupRule) implementsAccessGroupGetResponseResultRequire() {
}

type AccessGroupGetResponseResultRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                             `json:"email,required"`
	JSON  accessGroupGetResponseResultRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupGetResponseResultRequireAccessGsuiteGroupRuleGsuiteJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseResultRequireAccessGsuiteGroupRuleGsuite]
type accessGroupGetResponseResultRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupGetResponseResultRequireAccessOktaGroupRule struct {
	Okta AccessGroupGetResponseResultRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupGetResponseResultRequireAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupGetResponseResultRequireAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseResultRequireAccessOktaGroupRule]
type accessGroupGetResponseResultRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultRequireAccessOktaGroupRule) implementsAccessGroupGetResponseResultRequire() {
}

type AccessGroupGetResponseResultRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                         `json:"email,required"`
	JSON  accessGroupGetResponseResultRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupGetResponseResultRequireAccessOktaGroupRuleOktaJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultRequireAccessOktaGroupRuleOkta]
type accessGroupGetResponseResultRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupGetResponseResultRequireAccessSamlGroupRule struct {
	Saml AccessGroupGetResponseResultRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupGetResponseResultRequireAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupGetResponseResultRequireAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseResultRequireAccessSamlGroupRule]
type accessGroupGetResponseResultRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultRequireAccessSamlGroupRule) implementsAccessGroupGetResponseResultRequire() {
}

type AccessGroupGetResponseResultRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                         `json:"attribute_value,required"`
	JSON           accessGroupGetResponseResultRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupGetResponseResultRequireAccessSamlGroupRuleSamlJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultRequireAccessSamlGroupRuleSaml]
type accessGroupGetResponseResultRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupGetResponseResultRequireAccessServiceTokenRule struct {
	ServiceToken AccessGroupGetResponseResultRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupGetResponseResultRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupGetResponseResultRequireAccessServiceTokenRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultRequireAccessServiceTokenRule]
type accessGroupGetResponseResultRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultRequireAccessServiceTokenRule) implementsAccessGroupGetResponseResultRequire() {
}

type AccessGroupGetResponseResultRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                    `json:"token_id,required"`
	JSON    accessGroupGetResponseResultRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupGetResponseResultRequireAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseResultRequireAccessServiceTokenRuleServiceToken]
type accessGroupGetResponseResultRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupGetResponseResultRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                           `json:"any_valid_service_token,required"`
	JSON                 accessGroupGetResponseResultRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupGetResponseResultRequireAccessAnyValidServiceTokenRuleJSON contains
// the JSON metadata for the struct
// [AccessGroupGetResponseResultRequireAccessAnyValidServiceTokenRule]
type accessGroupGetResponseResultRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultRequireAccessAnyValidServiceTokenRule) implementsAccessGroupGetResponseResultRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupGetResponseResultRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupGetResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupGetResponseResultRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupGetResponseResultRequireAccessExternalEvaluationRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupGetResponseResultRequireAccessExternalEvaluationRule]
type accessGroupGetResponseResultRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultRequireAccessExternalEvaluationRule) implementsAccessGroupGetResponseResultRequire() {
}

type AccessGroupGetResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                `json:"keys_url,required"`
	JSON    accessGroupGetResponseResultRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupGetResponseResultRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupGetResponseResultRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupGetResponseResultRequireAccessCountryRule struct {
	Geo  AccessGroupGetResponseResultRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupGetResponseResultRequireAccessCountryRuleJSON `json:"-"`
}

// accessGroupGetResponseResultRequireAccessCountryRuleJSON contains the JSON
// metadata for the struct [AccessGroupGetResponseResultRequireAccessCountryRule]
type accessGroupGetResponseResultRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultRequireAccessCountryRule) implementsAccessGroupGetResponseResultRequire() {
}

type AccessGroupGetResponseResultRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                      `json:"country_code,required"`
	JSON        accessGroupGetResponseResultRequireAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupGetResponseResultRequireAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultRequireAccessCountryRuleGeo]
type accessGroupGetResponseResultRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupGetResponseResultRequireAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupGetResponseResultRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupGetResponseResultRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupGetResponseResultRequireAccessAuthenticationMethodRuleJSON contains
// the JSON metadata for the struct
// [AccessGroupGetResponseResultRequireAccessAuthenticationMethodRule]
type accessGroupGetResponseResultRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultRequireAccessAuthenticationMethodRule) implementsAccessGroupGetResponseResultRequire() {
}

type AccessGroupGetResponseResultRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                          `json:"auth_method,required"`
	JSON       accessGroupGetResponseResultRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupGetResponseResultRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseResultRequireAccessAuthenticationMethodRuleAuthMethod]
type accessGroupGetResponseResultRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupGetResponseResultRequireAccessDevicePostureRule struct {
	DevicePosture AccessGroupGetResponseResultRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupGetResponseResultRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupGetResponseResultRequireAccessDevicePostureRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupGetResponseResultRequireAccessDevicePostureRule]
type accessGroupGetResponseResultRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupGetResponseResultRequireAccessDevicePostureRule) implementsAccessGroupGetResponseResultRequire() {
}

type AccessGroupGetResponseResultRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                      `json:"integration_uid,required"`
	JSON           accessGroupGetResponseResultRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupGetResponseResultRequireAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupGetResponseResultRequireAccessDevicePostureRuleDevicePosture]
type accessGroupGetResponseResultRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupGetResponseResultRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessGroupGetResponseSuccess bool

const (
	AccessGroupGetResponseSuccessTrue AccessGroupGetResponseSuccess = true
)

type AccessGroupUpdateResponse struct {
	Errors   []AccessGroupUpdateResponseError   `json:"errors"`
	Messages []AccessGroupUpdateResponseMessage `json:"messages"`
	Result   AccessGroupUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessGroupUpdateResponseSuccess `json:"success"`
	JSON    accessGroupUpdateResponseJSON    `json:"-"`
}

// accessGroupUpdateResponseJSON contains the JSON metadata for the struct
// [AccessGroupUpdateResponse]
type accessGroupUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupUpdateResponseError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    accessGroupUpdateResponseErrorJSON `json:"-"`
}

// accessGroupUpdateResponseErrorJSON contains the JSON metadata for the struct
// [AccessGroupUpdateResponseError]
type accessGroupUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupUpdateResponseMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    accessGroupUpdateResponseMessageJSON `json:"-"`
}

// accessGroupUpdateResponseMessageJSON contains the JSON metadata for the struct
// [AccessGroupUpdateResponseMessage]
type accessGroupUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupUpdateResponseResult struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessGroupUpdateResponseResultExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessGroupUpdateResponseResultInclude `json:"include"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	IsDefault []AccessGroupUpdateResponseResultIsDefault `json:"is_default"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []AccessGroupUpdateResponseResultRequire `json:"require"`
	UpdatedAt time.Time                                `json:"updated_at" format:"date-time"`
	JSON      accessGroupUpdateResponseResultJSON      `json:"-"`
}

// accessGroupUpdateResponseResultJSON contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResult]
type accessGroupUpdateResponseResultJSON struct {
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

func (r *AccessGroupUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupUpdateResponseResultExcludeAccessEmailRule],
// [AccessGroupUpdateResponseResultExcludeAccessEmailListRule],
// [AccessGroupUpdateResponseResultExcludeAccessDomainRule],
// [AccessGroupUpdateResponseResultExcludeAccessEveryoneRule],
// [AccessGroupUpdateResponseResultExcludeAccessIPRule],
// [AccessGroupUpdateResponseResultExcludeAccessIPListRule],
// [AccessGroupUpdateResponseResultExcludeAccessCertificateRule],
// [AccessGroupUpdateResponseResultExcludeAccessAccessGroupRule],
// [AccessGroupUpdateResponseResultExcludeAccessAzureGroupRule],
// [AccessGroupUpdateResponseResultExcludeAccessGitHubOrganizationRule],
// [AccessGroupUpdateResponseResultExcludeAccessGsuiteGroupRule],
// [AccessGroupUpdateResponseResultExcludeAccessOktaGroupRule],
// [AccessGroupUpdateResponseResultExcludeAccessSamlGroupRule],
// [AccessGroupUpdateResponseResultExcludeAccessServiceTokenRule],
// [AccessGroupUpdateResponseResultExcludeAccessAnyValidServiceTokenRule],
// [AccessGroupUpdateResponseResultExcludeAccessExternalEvaluationRule],
// [AccessGroupUpdateResponseResultExcludeAccessCountryRule],
// [AccessGroupUpdateResponseResultExcludeAccessAuthenticationMethodRule] or
// [AccessGroupUpdateResponseResultExcludeAccessDevicePostureRule].
type AccessGroupUpdateResponseResultExclude interface {
	implementsAccessGroupUpdateResponseResultExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupUpdateResponseResultExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupUpdateResponseResultExcludeAccessEmailRule struct {
	Email AccessGroupUpdateResponseResultExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupUpdateResponseResultExcludeAccessEmailRuleJSON  `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessEmailRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseResultExcludeAccessEmailRule]
type accessGroupUpdateResponseResultExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultExcludeAccessEmailRule) implementsAccessGroupUpdateResponseResultExclude() {
}

type AccessGroupUpdateResponseResultExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                         `json:"email,required" format:"email"`
	JSON  accessGroupUpdateResponseResultExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseResultExcludeAccessEmailRuleEmail]
type accessGroupUpdateResponseResultExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupUpdateResponseResultExcludeAccessEmailListRule struct {
	EmailList AccessGroupUpdateResponseResultExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupUpdateResponseResultExcludeAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessEmailListRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseResultExcludeAccessEmailListRule]
type accessGroupUpdateResponseResultExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultExcludeAccessEmailListRule) implementsAccessGroupUpdateResponseResultExclude() {
}

type AccessGroupUpdateResponseResultExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                 `json:"id,required"`
	JSON accessGroupUpdateResponseResultExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessEmailListRuleEmailListJSON contains
// the JSON metadata for the struct
// [AccessGroupUpdateResponseResultExcludeAccessEmailListRuleEmailList]
type accessGroupUpdateResponseResultExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupUpdateResponseResultExcludeAccessDomainRule struct {
	EmailDomain AccessGroupUpdateResponseResultExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupUpdateResponseResultExcludeAccessDomainRuleJSON        `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessDomainRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseResultExcludeAccessDomainRule]
type accessGroupUpdateResponseResultExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultExcludeAccessDomainRule) implementsAccessGroupUpdateResponseResultExclude() {
}

type AccessGroupUpdateResponseResultExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                `json:"domain,required"`
	JSON   accessGroupUpdateResponseResultExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessDomainRuleEmailDomainJSON contains
// the JSON metadata for the struct
// [AccessGroupUpdateResponseResultExcludeAccessDomainRuleEmailDomain]
type accessGroupUpdateResponseResultExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupUpdateResponseResultExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                  `json:"everyone,required"`
	JSON     accessGroupUpdateResponseResultExcludeAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessEveryoneRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseResultExcludeAccessEveryoneRule]
type accessGroupUpdateResponseResultExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultExcludeAccessEveryoneRule) implementsAccessGroupUpdateResponseResultExclude() {
}

// Matches an IP address block.
type AccessGroupUpdateResponseResultExcludeAccessIPRule struct {
	IP   AccessGroupUpdateResponseResultExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupUpdateResponseResultExcludeAccessIPRuleJSON `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessIPRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseResultExcludeAccessIPRule]
type accessGroupUpdateResponseResultExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultExcludeAccessIPRule) implementsAccessGroupUpdateResponseResultExclude() {
}

type AccessGroupUpdateResponseResultExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                   `json:"ip,required"`
	JSON accessGroupUpdateResponseResultExcludeAccessIPRuleIPJSON `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessIPRuleIPJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseResultExcludeAccessIPRuleIP]
type accessGroupUpdateResponseResultExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupUpdateResponseResultExcludeAccessIPListRule struct {
	IPList AccessGroupUpdateResponseResultExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupUpdateResponseResultExcludeAccessIPListRuleJSON   `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessIPListRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseResultExcludeAccessIPListRule]
type accessGroupUpdateResponseResultExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultExcludeAccessIPListRule) implementsAccessGroupUpdateResponseResultExclude() {
}

type AccessGroupUpdateResponseResultExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                           `json:"id,required"`
	JSON accessGroupUpdateResponseResultExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultExcludeAccessIPListRuleIPList]
type accessGroupUpdateResponseResultExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupUpdateResponseResultExcludeAccessCertificateRule struct {
	Certificate interface{}                                                     `json:"certificate,required"`
	JSON        accessGroupUpdateResponseResultExcludeAccessCertificateRuleJSON `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultExcludeAccessCertificateRule]
type accessGroupUpdateResponseResultExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultExcludeAccessCertificateRule) implementsAccessGroupUpdateResponseResultExclude() {
}

// Matches an Access group.
type AccessGroupUpdateResponseResultExcludeAccessAccessGroupRule struct {
	Group AccessGroupUpdateResponseResultExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupUpdateResponseResultExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultExcludeAccessAccessGroupRule]
type accessGroupUpdateResponseResultExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultExcludeAccessAccessGroupRule) implementsAccessGroupUpdateResponseResultExclude() {
}

type AccessGroupUpdateResponseResultExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                               `json:"id,required"`
	JSON accessGroupUpdateResponseResultExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [AccessGroupUpdateResponseResultExcludeAccessAccessGroupRuleGroup]
type accessGroupUpdateResponseResultExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupUpdateResponseResultExcludeAccessAzureGroupRule struct {
	AzureAd AccessGroupUpdateResponseResultExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupUpdateResponseResultExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseResultExcludeAccessAzureGroupRule]
type accessGroupUpdateResponseResultExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultExcludeAccessAzureGroupRule) implementsAccessGroupUpdateResponseResultExclude() {
}

type AccessGroupUpdateResponseResultExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                `json:"connection_id,required"`
	JSON         accessGroupUpdateResponseResultExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessAzureGroupRuleAzureAdJSON contains
// the JSON metadata for the struct
// [AccessGroupUpdateResponseResultExcludeAccessAzureGroupRuleAzureAd]
type accessGroupUpdateResponseResultExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupUpdateResponseResultExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupUpdateResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupUpdateResponseResultExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessGitHubOrganizationRuleJSON contains
// the JSON metadata for the struct
// [AccessGroupUpdateResponseResultExcludeAccessGitHubOrganizationRule]
type accessGroupUpdateResponseResultExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultExcludeAccessGitHubOrganizationRule) implementsAccessGroupUpdateResponseResultExclude() {
}

type AccessGroupUpdateResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                   `json:"name,required"`
	JSON accessGroupUpdateResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupUpdateResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupUpdateResponseResultExcludeAccessGsuiteGroupRule struct {
	Gsuite AccessGroupUpdateResponseResultExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupUpdateResponseResultExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultExcludeAccessGsuiteGroupRule]
type accessGroupUpdateResponseResultExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultExcludeAccessGsuiteGroupRule) implementsAccessGroupUpdateResponseResultExclude() {
}

type AccessGroupUpdateResponseResultExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                `json:"email,required"`
	JSON  accessGroupUpdateResponseResultExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessGsuiteGroupRuleGsuiteJSON contains
// the JSON metadata for the struct
// [AccessGroupUpdateResponseResultExcludeAccessGsuiteGroupRuleGsuite]
type accessGroupUpdateResponseResultExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupUpdateResponseResultExcludeAccessOktaGroupRule struct {
	Okta AccessGroupUpdateResponseResultExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupUpdateResponseResultExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseResultExcludeAccessOktaGroupRule]
type accessGroupUpdateResponseResultExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultExcludeAccessOktaGroupRule) implementsAccessGroupUpdateResponseResultExclude() {
}

type AccessGroupUpdateResponseResultExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                            `json:"email,required"`
	JSON  accessGroupUpdateResponseResultExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessOktaGroupRuleOktaJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultExcludeAccessOktaGroupRuleOkta]
type accessGroupUpdateResponseResultExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupUpdateResponseResultExcludeAccessSamlGroupRule struct {
	Saml AccessGroupUpdateResponseResultExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupUpdateResponseResultExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseResultExcludeAccessSamlGroupRule]
type accessGroupUpdateResponseResultExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultExcludeAccessSamlGroupRule) implementsAccessGroupUpdateResponseResultExclude() {
}

type AccessGroupUpdateResponseResultExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                            `json:"attribute_value,required"`
	JSON           accessGroupUpdateResponseResultExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessSamlGroupRuleSamlJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultExcludeAccessSamlGroupRuleSaml]
type accessGroupUpdateResponseResultExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupUpdateResponseResultExcludeAccessServiceTokenRule struct {
	ServiceToken AccessGroupUpdateResponseResultExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupUpdateResponseResultExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultExcludeAccessServiceTokenRule]
type accessGroupUpdateResponseResultExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultExcludeAccessServiceTokenRule) implementsAccessGroupUpdateResponseResultExclude() {
}

type AccessGroupUpdateResponseResultExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                       `json:"token_id,required"`
	JSON    accessGroupUpdateResponseResultExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResultExcludeAccessServiceTokenRuleServiceToken]
type accessGroupUpdateResponseResultExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupUpdateResponseResultExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                              `json:"any_valid_service_token,required"`
	JSON                 accessGroupUpdateResponseResultExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResultExcludeAccessAnyValidServiceTokenRule]
type accessGroupUpdateResponseResultExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultExcludeAccessAnyValidServiceTokenRule) implementsAccessGroupUpdateResponseResultExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupUpdateResponseResultExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupUpdateResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupUpdateResponseResultExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessExternalEvaluationRuleJSON contains
// the JSON metadata for the struct
// [AccessGroupUpdateResponseResultExcludeAccessExternalEvaluationRule]
type accessGroupUpdateResponseResultExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultExcludeAccessExternalEvaluationRule) implementsAccessGroupUpdateResponseResultExclude() {
}

type AccessGroupUpdateResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                   `json:"keys_url,required"`
	JSON    accessGroupUpdateResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupUpdateResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupUpdateResponseResultExcludeAccessCountryRule struct {
	Geo  AccessGroupUpdateResponseResultExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupUpdateResponseResultExcludeAccessCountryRuleJSON `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessCountryRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseResultExcludeAccessCountryRule]
type accessGroupUpdateResponseResultExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultExcludeAccessCountryRule) implementsAccessGroupUpdateResponseResultExclude() {
}

type AccessGroupUpdateResponseResultExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                         `json:"country_code,required"`
	JSON        accessGroupUpdateResponseResultExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseResultExcludeAccessCountryRuleGeo]
type accessGroupUpdateResponseResultExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupUpdateResponseResultExcludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupUpdateResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupUpdateResponseResultExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResultExcludeAccessAuthenticationMethodRule]
type accessGroupUpdateResponseResultExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultExcludeAccessAuthenticationMethodRule) implementsAccessGroupUpdateResponseResultExclude() {
}

type AccessGroupUpdateResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                             `json:"auth_method,required"`
	JSON       accessGroupUpdateResponseResultExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod]
type accessGroupUpdateResponseResultExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupUpdateResponseResultExcludeAccessDevicePostureRule struct {
	DevicePosture AccessGroupUpdateResponseResultExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupUpdateResponseResultExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessDevicePostureRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultExcludeAccessDevicePostureRule]
type accessGroupUpdateResponseResultExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultExcludeAccessDevicePostureRule) implementsAccessGroupUpdateResponseResultExclude() {
}

type AccessGroupUpdateResponseResultExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                         `json:"integration_uid,required"`
	JSON           accessGroupUpdateResponseResultExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupUpdateResponseResultExcludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResultExcludeAccessDevicePostureRuleDevicePosture]
type accessGroupUpdateResponseResultExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupUpdateResponseResultIncludeAccessEmailRule],
// [AccessGroupUpdateResponseResultIncludeAccessEmailListRule],
// [AccessGroupUpdateResponseResultIncludeAccessDomainRule],
// [AccessGroupUpdateResponseResultIncludeAccessEveryoneRule],
// [AccessGroupUpdateResponseResultIncludeAccessIPRule],
// [AccessGroupUpdateResponseResultIncludeAccessIPListRule],
// [AccessGroupUpdateResponseResultIncludeAccessCertificateRule],
// [AccessGroupUpdateResponseResultIncludeAccessAccessGroupRule],
// [AccessGroupUpdateResponseResultIncludeAccessAzureGroupRule],
// [AccessGroupUpdateResponseResultIncludeAccessGitHubOrganizationRule],
// [AccessGroupUpdateResponseResultIncludeAccessGsuiteGroupRule],
// [AccessGroupUpdateResponseResultIncludeAccessOktaGroupRule],
// [AccessGroupUpdateResponseResultIncludeAccessSamlGroupRule],
// [AccessGroupUpdateResponseResultIncludeAccessServiceTokenRule],
// [AccessGroupUpdateResponseResultIncludeAccessAnyValidServiceTokenRule],
// [AccessGroupUpdateResponseResultIncludeAccessExternalEvaluationRule],
// [AccessGroupUpdateResponseResultIncludeAccessCountryRule],
// [AccessGroupUpdateResponseResultIncludeAccessAuthenticationMethodRule] or
// [AccessGroupUpdateResponseResultIncludeAccessDevicePostureRule].
type AccessGroupUpdateResponseResultInclude interface {
	implementsAccessGroupUpdateResponseResultInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupUpdateResponseResultInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupUpdateResponseResultIncludeAccessEmailRule struct {
	Email AccessGroupUpdateResponseResultIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupUpdateResponseResultIncludeAccessEmailRuleJSON  `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessEmailRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseResultIncludeAccessEmailRule]
type accessGroupUpdateResponseResultIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIncludeAccessEmailRule) implementsAccessGroupUpdateResponseResultInclude() {
}

type AccessGroupUpdateResponseResultIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                         `json:"email,required" format:"email"`
	JSON  accessGroupUpdateResponseResultIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseResultIncludeAccessEmailRuleEmail]
type accessGroupUpdateResponseResultIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupUpdateResponseResultIncludeAccessEmailListRule struct {
	EmailList AccessGroupUpdateResponseResultIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupUpdateResponseResultIncludeAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessEmailListRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseResultIncludeAccessEmailListRule]
type accessGroupUpdateResponseResultIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIncludeAccessEmailListRule) implementsAccessGroupUpdateResponseResultInclude() {
}

type AccessGroupUpdateResponseResultIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                 `json:"id,required"`
	JSON accessGroupUpdateResponseResultIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessEmailListRuleEmailListJSON contains
// the JSON metadata for the struct
// [AccessGroupUpdateResponseResultIncludeAccessEmailListRuleEmailList]
type accessGroupUpdateResponseResultIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupUpdateResponseResultIncludeAccessDomainRule struct {
	EmailDomain AccessGroupUpdateResponseResultIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupUpdateResponseResultIncludeAccessDomainRuleJSON        `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessDomainRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseResultIncludeAccessDomainRule]
type accessGroupUpdateResponseResultIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIncludeAccessDomainRule) implementsAccessGroupUpdateResponseResultInclude() {
}

type AccessGroupUpdateResponseResultIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                `json:"domain,required"`
	JSON   accessGroupUpdateResponseResultIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessDomainRuleEmailDomainJSON contains
// the JSON metadata for the struct
// [AccessGroupUpdateResponseResultIncludeAccessDomainRuleEmailDomain]
type accessGroupUpdateResponseResultIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupUpdateResponseResultIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                  `json:"everyone,required"`
	JSON     accessGroupUpdateResponseResultIncludeAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessEveryoneRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseResultIncludeAccessEveryoneRule]
type accessGroupUpdateResponseResultIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIncludeAccessEveryoneRule) implementsAccessGroupUpdateResponseResultInclude() {
}

// Matches an IP address block.
type AccessGroupUpdateResponseResultIncludeAccessIPRule struct {
	IP   AccessGroupUpdateResponseResultIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupUpdateResponseResultIncludeAccessIPRuleJSON `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessIPRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseResultIncludeAccessIPRule]
type accessGroupUpdateResponseResultIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIncludeAccessIPRule) implementsAccessGroupUpdateResponseResultInclude() {
}

type AccessGroupUpdateResponseResultIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                   `json:"ip,required"`
	JSON accessGroupUpdateResponseResultIncludeAccessIPRuleIPJSON `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessIPRuleIPJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseResultIncludeAccessIPRuleIP]
type accessGroupUpdateResponseResultIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupUpdateResponseResultIncludeAccessIPListRule struct {
	IPList AccessGroupUpdateResponseResultIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupUpdateResponseResultIncludeAccessIPListRuleJSON   `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessIPListRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseResultIncludeAccessIPListRule]
type accessGroupUpdateResponseResultIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIncludeAccessIPListRule) implementsAccessGroupUpdateResponseResultInclude() {
}

type AccessGroupUpdateResponseResultIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                           `json:"id,required"`
	JSON accessGroupUpdateResponseResultIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultIncludeAccessIPListRuleIPList]
type accessGroupUpdateResponseResultIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupUpdateResponseResultIncludeAccessCertificateRule struct {
	Certificate interface{}                                                     `json:"certificate,required"`
	JSON        accessGroupUpdateResponseResultIncludeAccessCertificateRuleJSON `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultIncludeAccessCertificateRule]
type accessGroupUpdateResponseResultIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIncludeAccessCertificateRule) implementsAccessGroupUpdateResponseResultInclude() {
}

// Matches an Access group.
type AccessGroupUpdateResponseResultIncludeAccessAccessGroupRule struct {
	Group AccessGroupUpdateResponseResultIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupUpdateResponseResultIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultIncludeAccessAccessGroupRule]
type accessGroupUpdateResponseResultIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIncludeAccessAccessGroupRule) implementsAccessGroupUpdateResponseResultInclude() {
}

type AccessGroupUpdateResponseResultIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                               `json:"id,required"`
	JSON accessGroupUpdateResponseResultIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [AccessGroupUpdateResponseResultIncludeAccessAccessGroupRuleGroup]
type accessGroupUpdateResponseResultIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupUpdateResponseResultIncludeAccessAzureGroupRule struct {
	AzureAd AccessGroupUpdateResponseResultIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupUpdateResponseResultIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseResultIncludeAccessAzureGroupRule]
type accessGroupUpdateResponseResultIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIncludeAccessAzureGroupRule) implementsAccessGroupUpdateResponseResultInclude() {
}

type AccessGroupUpdateResponseResultIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                `json:"connection_id,required"`
	JSON         accessGroupUpdateResponseResultIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessAzureGroupRuleAzureAdJSON contains
// the JSON metadata for the struct
// [AccessGroupUpdateResponseResultIncludeAccessAzureGroupRuleAzureAd]
type accessGroupUpdateResponseResultIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupUpdateResponseResultIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupUpdateResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupUpdateResponseResultIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessGitHubOrganizationRuleJSON contains
// the JSON metadata for the struct
// [AccessGroupUpdateResponseResultIncludeAccessGitHubOrganizationRule]
type accessGroupUpdateResponseResultIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIncludeAccessGitHubOrganizationRule) implementsAccessGroupUpdateResponseResultInclude() {
}

type AccessGroupUpdateResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                   `json:"name,required"`
	JSON accessGroupUpdateResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupUpdateResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupUpdateResponseResultIncludeAccessGsuiteGroupRule struct {
	Gsuite AccessGroupUpdateResponseResultIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupUpdateResponseResultIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultIncludeAccessGsuiteGroupRule]
type accessGroupUpdateResponseResultIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIncludeAccessGsuiteGroupRule) implementsAccessGroupUpdateResponseResultInclude() {
}

type AccessGroupUpdateResponseResultIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                `json:"email,required"`
	JSON  accessGroupUpdateResponseResultIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessGsuiteGroupRuleGsuiteJSON contains
// the JSON metadata for the struct
// [AccessGroupUpdateResponseResultIncludeAccessGsuiteGroupRuleGsuite]
type accessGroupUpdateResponseResultIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupUpdateResponseResultIncludeAccessOktaGroupRule struct {
	Okta AccessGroupUpdateResponseResultIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupUpdateResponseResultIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseResultIncludeAccessOktaGroupRule]
type accessGroupUpdateResponseResultIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIncludeAccessOktaGroupRule) implementsAccessGroupUpdateResponseResultInclude() {
}

type AccessGroupUpdateResponseResultIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                            `json:"email,required"`
	JSON  accessGroupUpdateResponseResultIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessOktaGroupRuleOktaJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultIncludeAccessOktaGroupRuleOkta]
type accessGroupUpdateResponseResultIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupUpdateResponseResultIncludeAccessSamlGroupRule struct {
	Saml AccessGroupUpdateResponseResultIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupUpdateResponseResultIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseResultIncludeAccessSamlGroupRule]
type accessGroupUpdateResponseResultIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIncludeAccessSamlGroupRule) implementsAccessGroupUpdateResponseResultInclude() {
}

type AccessGroupUpdateResponseResultIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                            `json:"attribute_value,required"`
	JSON           accessGroupUpdateResponseResultIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessSamlGroupRuleSamlJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultIncludeAccessSamlGroupRuleSaml]
type accessGroupUpdateResponseResultIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupUpdateResponseResultIncludeAccessServiceTokenRule struct {
	ServiceToken AccessGroupUpdateResponseResultIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupUpdateResponseResultIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultIncludeAccessServiceTokenRule]
type accessGroupUpdateResponseResultIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIncludeAccessServiceTokenRule) implementsAccessGroupUpdateResponseResultInclude() {
}

type AccessGroupUpdateResponseResultIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                       `json:"token_id,required"`
	JSON    accessGroupUpdateResponseResultIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResultIncludeAccessServiceTokenRuleServiceToken]
type accessGroupUpdateResponseResultIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupUpdateResponseResultIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                              `json:"any_valid_service_token,required"`
	JSON                 accessGroupUpdateResponseResultIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResultIncludeAccessAnyValidServiceTokenRule]
type accessGroupUpdateResponseResultIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIncludeAccessAnyValidServiceTokenRule) implementsAccessGroupUpdateResponseResultInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupUpdateResponseResultIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupUpdateResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupUpdateResponseResultIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessExternalEvaluationRuleJSON contains
// the JSON metadata for the struct
// [AccessGroupUpdateResponseResultIncludeAccessExternalEvaluationRule]
type accessGroupUpdateResponseResultIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIncludeAccessExternalEvaluationRule) implementsAccessGroupUpdateResponseResultInclude() {
}

type AccessGroupUpdateResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                   `json:"keys_url,required"`
	JSON    accessGroupUpdateResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupUpdateResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupUpdateResponseResultIncludeAccessCountryRule struct {
	Geo  AccessGroupUpdateResponseResultIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupUpdateResponseResultIncludeAccessCountryRuleJSON `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessCountryRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseResultIncludeAccessCountryRule]
type accessGroupUpdateResponseResultIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIncludeAccessCountryRule) implementsAccessGroupUpdateResponseResultInclude() {
}

type AccessGroupUpdateResponseResultIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                         `json:"country_code,required"`
	JSON        accessGroupUpdateResponseResultIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseResultIncludeAccessCountryRuleGeo]
type accessGroupUpdateResponseResultIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupUpdateResponseResultIncludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupUpdateResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupUpdateResponseResultIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResultIncludeAccessAuthenticationMethodRule]
type accessGroupUpdateResponseResultIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIncludeAccessAuthenticationMethodRule) implementsAccessGroupUpdateResponseResultInclude() {
}

type AccessGroupUpdateResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                             `json:"auth_method,required"`
	JSON       accessGroupUpdateResponseResultIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod]
type accessGroupUpdateResponseResultIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupUpdateResponseResultIncludeAccessDevicePostureRule struct {
	DevicePosture AccessGroupUpdateResponseResultIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupUpdateResponseResultIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessDevicePostureRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultIncludeAccessDevicePostureRule]
type accessGroupUpdateResponseResultIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIncludeAccessDevicePostureRule) implementsAccessGroupUpdateResponseResultInclude() {
}

type AccessGroupUpdateResponseResultIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                         `json:"integration_uid,required"`
	JSON           accessGroupUpdateResponseResultIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupUpdateResponseResultIncludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResultIncludeAccessDevicePostureRuleDevicePosture]
type accessGroupUpdateResponseResultIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupUpdateResponseResultIsDefaultAccessEmailRule],
// [AccessGroupUpdateResponseResultIsDefaultAccessEmailListRule],
// [AccessGroupUpdateResponseResultIsDefaultAccessDomainRule],
// [AccessGroupUpdateResponseResultIsDefaultAccessEveryoneRule],
// [AccessGroupUpdateResponseResultIsDefaultAccessIPRule],
// [AccessGroupUpdateResponseResultIsDefaultAccessIPListRule],
// [AccessGroupUpdateResponseResultIsDefaultAccessCertificateRule],
// [AccessGroupUpdateResponseResultIsDefaultAccessAccessGroupRule],
// [AccessGroupUpdateResponseResultIsDefaultAccessAzureGroupRule],
// [AccessGroupUpdateResponseResultIsDefaultAccessGitHubOrganizationRule],
// [AccessGroupUpdateResponseResultIsDefaultAccessGsuiteGroupRule],
// [AccessGroupUpdateResponseResultIsDefaultAccessOktaGroupRule],
// [AccessGroupUpdateResponseResultIsDefaultAccessSamlGroupRule],
// [AccessGroupUpdateResponseResultIsDefaultAccessServiceTokenRule],
// [AccessGroupUpdateResponseResultIsDefaultAccessAnyValidServiceTokenRule],
// [AccessGroupUpdateResponseResultIsDefaultAccessExternalEvaluationRule],
// [AccessGroupUpdateResponseResultIsDefaultAccessCountryRule],
// [AccessGroupUpdateResponseResultIsDefaultAccessAuthenticationMethodRule] or
// [AccessGroupUpdateResponseResultIsDefaultAccessDevicePostureRule].
type AccessGroupUpdateResponseResultIsDefault interface {
	implementsAccessGroupUpdateResponseResultIsDefault()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupUpdateResponseResultIsDefault)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupUpdateResponseResultIsDefaultAccessEmailRule struct {
	Email AccessGroupUpdateResponseResultIsDefaultAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupUpdateResponseResultIsDefaultAccessEmailRuleJSON  `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessEmailRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessEmailRule]
type accessGroupUpdateResponseResultIsDefaultAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIsDefaultAccessEmailRule) implementsAccessGroupUpdateResponseResultIsDefault() {
}

type AccessGroupUpdateResponseResultIsDefaultAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                           `json:"email,required" format:"email"`
	JSON  accessGroupUpdateResponseResultIsDefaultAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessEmailRuleEmail]
type accessGroupUpdateResponseResultIsDefaultAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupUpdateResponseResultIsDefaultAccessEmailListRule struct {
	EmailList AccessGroupUpdateResponseResultIsDefaultAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupUpdateResponseResultIsDefaultAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessEmailListRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessEmailListRule]
type accessGroupUpdateResponseResultIsDefaultAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIsDefaultAccessEmailListRule) implementsAccessGroupUpdateResponseResultIsDefault() {
}

type AccessGroupUpdateResponseResultIsDefaultAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                   `json:"id,required"`
	JSON accessGroupUpdateResponseResultIsDefaultAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessEmailListRuleEmailList]
type accessGroupUpdateResponseResultIsDefaultAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupUpdateResponseResultIsDefaultAccessDomainRule struct {
	EmailDomain AccessGroupUpdateResponseResultIsDefaultAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupUpdateResponseResultIsDefaultAccessDomainRuleJSON        `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessDomainRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessDomainRule]
type accessGroupUpdateResponseResultIsDefaultAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIsDefaultAccessDomainRule) implementsAccessGroupUpdateResponseResultIsDefault() {
}

type AccessGroupUpdateResponseResultIsDefaultAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                  `json:"domain,required"`
	JSON   accessGroupUpdateResponseResultIsDefaultAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessDomainRuleEmailDomainJSON contains
// the JSON metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessDomainRuleEmailDomain]
type accessGroupUpdateResponseResultIsDefaultAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupUpdateResponseResultIsDefaultAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                    `json:"everyone,required"`
	JSON     accessGroupUpdateResponseResultIsDefaultAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessEveryoneRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessEveryoneRule]
type accessGroupUpdateResponseResultIsDefaultAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIsDefaultAccessEveryoneRule) implementsAccessGroupUpdateResponseResultIsDefault() {
}

// Matches an IP address block.
type AccessGroupUpdateResponseResultIsDefaultAccessIPRule struct {
	IP   AccessGroupUpdateResponseResultIsDefaultAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupUpdateResponseResultIsDefaultAccessIPRuleJSON `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessIPRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseResultIsDefaultAccessIPRule]
type accessGroupUpdateResponseResultIsDefaultAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIsDefaultAccessIPRule) implementsAccessGroupUpdateResponseResultIsDefault() {
}

type AccessGroupUpdateResponseResultIsDefaultAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                     `json:"ip,required"`
	JSON accessGroupUpdateResponseResultIsDefaultAccessIPRuleIPJSON `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessIPRuleIPJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseResultIsDefaultAccessIPRuleIP]
type accessGroupUpdateResponseResultIsDefaultAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupUpdateResponseResultIsDefaultAccessIPListRule struct {
	IPList AccessGroupUpdateResponseResultIsDefaultAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupUpdateResponseResultIsDefaultAccessIPListRuleJSON   `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessIPListRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessIPListRule]
type accessGroupUpdateResponseResultIsDefaultAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIsDefaultAccessIPListRule) implementsAccessGroupUpdateResponseResultIsDefault() {
}

type AccessGroupUpdateResponseResultIsDefaultAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                             `json:"id,required"`
	JSON accessGroupUpdateResponseResultIsDefaultAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessIPListRuleIPList]
type accessGroupUpdateResponseResultIsDefaultAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupUpdateResponseResultIsDefaultAccessCertificateRule struct {
	Certificate interface{}                                                       `json:"certificate,required"`
	JSON        accessGroupUpdateResponseResultIsDefaultAccessCertificateRuleJSON `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessCertificateRule]
type accessGroupUpdateResponseResultIsDefaultAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIsDefaultAccessCertificateRule) implementsAccessGroupUpdateResponseResultIsDefault() {
}

// Matches an Access group.
type AccessGroupUpdateResponseResultIsDefaultAccessAccessGroupRule struct {
	Group AccessGroupUpdateResponseResultIsDefaultAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupUpdateResponseResultIsDefaultAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessAccessGroupRule]
type accessGroupUpdateResponseResultIsDefaultAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIsDefaultAccessAccessGroupRule) implementsAccessGroupUpdateResponseResultIsDefault() {
}

type AccessGroupUpdateResponseResultIsDefaultAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                 `json:"id,required"`
	JSON accessGroupUpdateResponseResultIsDefaultAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessAccessGroupRuleGroup]
type accessGroupUpdateResponseResultIsDefaultAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupUpdateResponseResultIsDefaultAccessAzureGroupRule struct {
	AzureAd AccessGroupUpdateResponseResultIsDefaultAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupUpdateResponseResultIsDefaultAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessAzureGroupRule]
type accessGroupUpdateResponseResultIsDefaultAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIsDefaultAccessAzureGroupRule) implementsAccessGroupUpdateResponseResultIsDefault() {
}

type AccessGroupUpdateResponseResultIsDefaultAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                  `json:"connection_id,required"`
	JSON         accessGroupUpdateResponseResultIsDefaultAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessAzureGroupRuleAzureAdJSON contains
// the JSON metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessAzureGroupRuleAzureAd]
type accessGroupUpdateResponseResultIsDefaultAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupUpdateResponseResultIsDefaultAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupUpdateResponseResultIsDefaultAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupUpdateResponseResultIsDefaultAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessGitHubOrganizationRule]
type accessGroupUpdateResponseResultIsDefaultAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIsDefaultAccessGitHubOrganizationRule) implementsAccessGroupUpdateResponseResultIsDefault() {
}

type AccessGroupUpdateResponseResultIsDefaultAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                     `json:"name,required"`
	JSON accessGroupUpdateResponseResultIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupUpdateResponseResultIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupUpdateResponseResultIsDefaultAccessGsuiteGroupRule struct {
	Gsuite AccessGroupUpdateResponseResultIsDefaultAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupUpdateResponseResultIsDefaultAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessGsuiteGroupRule]
type accessGroupUpdateResponseResultIsDefaultAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIsDefaultAccessGsuiteGroupRule) implementsAccessGroupUpdateResponseResultIsDefault() {
}

type AccessGroupUpdateResponseResultIsDefaultAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                  `json:"email,required"`
	JSON  accessGroupUpdateResponseResultIsDefaultAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessGsuiteGroupRuleGsuiteJSON contains
// the JSON metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessGsuiteGroupRuleGsuite]
type accessGroupUpdateResponseResultIsDefaultAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupUpdateResponseResultIsDefaultAccessOktaGroupRule struct {
	Okta AccessGroupUpdateResponseResultIsDefaultAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupUpdateResponseResultIsDefaultAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessOktaGroupRule]
type accessGroupUpdateResponseResultIsDefaultAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIsDefaultAccessOktaGroupRule) implementsAccessGroupUpdateResponseResultIsDefault() {
}

type AccessGroupUpdateResponseResultIsDefaultAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                              `json:"email,required"`
	JSON  accessGroupUpdateResponseResultIsDefaultAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessOktaGroupRuleOktaJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessOktaGroupRuleOkta]
type accessGroupUpdateResponseResultIsDefaultAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupUpdateResponseResultIsDefaultAccessSamlGroupRule struct {
	Saml AccessGroupUpdateResponseResultIsDefaultAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupUpdateResponseResultIsDefaultAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessSamlGroupRule]
type accessGroupUpdateResponseResultIsDefaultAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIsDefaultAccessSamlGroupRule) implementsAccessGroupUpdateResponseResultIsDefault() {
}

type AccessGroupUpdateResponseResultIsDefaultAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                              `json:"attribute_value,required"`
	JSON           accessGroupUpdateResponseResultIsDefaultAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessSamlGroupRuleSamlJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessSamlGroupRuleSaml]
type accessGroupUpdateResponseResultIsDefaultAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupUpdateResponseResultIsDefaultAccessServiceTokenRule struct {
	ServiceToken AccessGroupUpdateResponseResultIsDefaultAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupUpdateResponseResultIsDefaultAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessServiceTokenRule]
type accessGroupUpdateResponseResultIsDefaultAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIsDefaultAccessServiceTokenRule) implementsAccessGroupUpdateResponseResultIsDefault() {
}

type AccessGroupUpdateResponseResultIsDefaultAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                         `json:"token_id,required"`
	JSON    accessGroupUpdateResponseResultIsDefaultAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessServiceTokenRuleServiceToken]
type accessGroupUpdateResponseResultIsDefaultAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupUpdateResponseResultIsDefaultAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                `json:"any_valid_service_token,required"`
	JSON                 accessGroupUpdateResponseResultIsDefaultAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessAnyValidServiceTokenRule]
type accessGroupUpdateResponseResultIsDefaultAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIsDefaultAccessAnyValidServiceTokenRule) implementsAccessGroupUpdateResponseResultIsDefault() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupUpdateResponseResultIsDefaultAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupUpdateResponseResultIsDefaultAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupUpdateResponseResultIsDefaultAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessExternalEvaluationRule]
type accessGroupUpdateResponseResultIsDefaultAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIsDefaultAccessExternalEvaluationRule) implementsAccessGroupUpdateResponseResultIsDefault() {
}

type AccessGroupUpdateResponseResultIsDefaultAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                     `json:"keys_url,required"`
	JSON    accessGroupUpdateResponseResultIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupUpdateResponseResultIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupUpdateResponseResultIsDefaultAccessCountryRule struct {
	Geo  AccessGroupUpdateResponseResultIsDefaultAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupUpdateResponseResultIsDefaultAccessCountryRuleJSON `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessCountryRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessCountryRule]
type accessGroupUpdateResponseResultIsDefaultAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIsDefaultAccessCountryRule) implementsAccessGroupUpdateResponseResultIsDefault() {
}

type AccessGroupUpdateResponseResultIsDefaultAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                           `json:"country_code,required"`
	JSON        accessGroupUpdateResponseResultIsDefaultAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessCountryRuleGeo]
type accessGroupUpdateResponseResultIsDefaultAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupUpdateResponseResultIsDefaultAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupUpdateResponseResultIsDefaultAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupUpdateResponseResultIsDefaultAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessAuthenticationMethodRule]
type accessGroupUpdateResponseResultIsDefaultAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIsDefaultAccessAuthenticationMethodRule) implementsAccessGroupUpdateResponseResultIsDefault() {
}

type AccessGroupUpdateResponseResultIsDefaultAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                               `json:"auth_method,required"`
	JSON       accessGroupUpdateResponseResultIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessAuthenticationMethodRuleAuthMethod]
type accessGroupUpdateResponseResultIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupUpdateResponseResultIsDefaultAccessDevicePostureRule struct {
	DevicePosture AccessGroupUpdateResponseResultIsDefaultAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupUpdateResponseResultIsDefaultAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessDevicePostureRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessDevicePostureRule]
type accessGroupUpdateResponseResultIsDefaultAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultIsDefaultAccessDevicePostureRule) implementsAccessGroupUpdateResponseResultIsDefault() {
}

type AccessGroupUpdateResponseResultIsDefaultAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                           `json:"integration_uid,required"`
	JSON           accessGroupUpdateResponseResultIsDefaultAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupUpdateResponseResultIsDefaultAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResultIsDefaultAccessDevicePostureRuleDevicePosture]
type accessGroupUpdateResponseResultIsDefaultAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultIsDefaultAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupUpdateResponseResultRequireAccessEmailRule],
// [AccessGroupUpdateResponseResultRequireAccessEmailListRule],
// [AccessGroupUpdateResponseResultRequireAccessDomainRule],
// [AccessGroupUpdateResponseResultRequireAccessEveryoneRule],
// [AccessGroupUpdateResponseResultRequireAccessIPRule],
// [AccessGroupUpdateResponseResultRequireAccessIPListRule],
// [AccessGroupUpdateResponseResultRequireAccessCertificateRule],
// [AccessGroupUpdateResponseResultRequireAccessAccessGroupRule],
// [AccessGroupUpdateResponseResultRequireAccessAzureGroupRule],
// [AccessGroupUpdateResponseResultRequireAccessGitHubOrganizationRule],
// [AccessGroupUpdateResponseResultRequireAccessGsuiteGroupRule],
// [AccessGroupUpdateResponseResultRequireAccessOktaGroupRule],
// [AccessGroupUpdateResponseResultRequireAccessSamlGroupRule],
// [AccessGroupUpdateResponseResultRequireAccessServiceTokenRule],
// [AccessGroupUpdateResponseResultRequireAccessAnyValidServiceTokenRule],
// [AccessGroupUpdateResponseResultRequireAccessExternalEvaluationRule],
// [AccessGroupUpdateResponseResultRequireAccessCountryRule],
// [AccessGroupUpdateResponseResultRequireAccessAuthenticationMethodRule] or
// [AccessGroupUpdateResponseResultRequireAccessDevicePostureRule].
type AccessGroupUpdateResponseResultRequire interface {
	implementsAccessGroupUpdateResponseResultRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupUpdateResponseResultRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupUpdateResponseResultRequireAccessEmailRule struct {
	Email AccessGroupUpdateResponseResultRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupUpdateResponseResultRequireAccessEmailRuleJSON  `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessEmailRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseResultRequireAccessEmailRule]
type accessGroupUpdateResponseResultRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultRequireAccessEmailRule) implementsAccessGroupUpdateResponseResultRequire() {
}

type AccessGroupUpdateResponseResultRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                         `json:"email,required" format:"email"`
	JSON  accessGroupUpdateResponseResultRequireAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseResultRequireAccessEmailRuleEmail]
type accessGroupUpdateResponseResultRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupUpdateResponseResultRequireAccessEmailListRule struct {
	EmailList AccessGroupUpdateResponseResultRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupUpdateResponseResultRequireAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessEmailListRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseResultRequireAccessEmailListRule]
type accessGroupUpdateResponseResultRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultRequireAccessEmailListRule) implementsAccessGroupUpdateResponseResultRequire() {
}

type AccessGroupUpdateResponseResultRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                 `json:"id,required"`
	JSON accessGroupUpdateResponseResultRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessEmailListRuleEmailListJSON contains
// the JSON metadata for the struct
// [AccessGroupUpdateResponseResultRequireAccessEmailListRuleEmailList]
type accessGroupUpdateResponseResultRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupUpdateResponseResultRequireAccessDomainRule struct {
	EmailDomain AccessGroupUpdateResponseResultRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupUpdateResponseResultRequireAccessDomainRuleJSON        `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessDomainRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseResultRequireAccessDomainRule]
type accessGroupUpdateResponseResultRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultRequireAccessDomainRule) implementsAccessGroupUpdateResponseResultRequire() {
}

type AccessGroupUpdateResponseResultRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                `json:"domain,required"`
	JSON   accessGroupUpdateResponseResultRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessDomainRuleEmailDomainJSON contains
// the JSON metadata for the struct
// [AccessGroupUpdateResponseResultRequireAccessDomainRuleEmailDomain]
type accessGroupUpdateResponseResultRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupUpdateResponseResultRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                  `json:"everyone,required"`
	JSON     accessGroupUpdateResponseResultRequireAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessEveryoneRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseResultRequireAccessEveryoneRule]
type accessGroupUpdateResponseResultRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultRequireAccessEveryoneRule) implementsAccessGroupUpdateResponseResultRequire() {
}

// Matches an IP address block.
type AccessGroupUpdateResponseResultRequireAccessIPRule struct {
	IP   AccessGroupUpdateResponseResultRequireAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupUpdateResponseResultRequireAccessIPRuleJSON `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessIPRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseResultRequireAccessIPRule]
type accessGroupUpdateResponseResultRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultRequireAccessIPRule) implementsAccessGroupUpdateResponseResultRequire() {
}

type AccessGroupUpdateResponseResultRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                   `json:"ip,required"`
	JSON accessGroupUpdateResponseResultRequireAccessIPRuleIPJSON `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessIPRuleIPJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseResultRequireAccessIPRuleIP]
type accessGroupUpdateResponseResultRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupUpdateResponseResultRequireAccessIPListRule struct {
	IPList AccessGroupUpdateResponseResultRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupUpdateResponseResultRequireAccessIPListRuleJSON   `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessIPListRuleJSON contains the JSON
// metadata for the struct [AccessGroupUpdateResponseResultRequireAccessIPListRule]
type accessGroupUpdateResponseResultRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultRequireAccessIPListRule) implementsAccessGroupUpdateResponseResultRequire() {
}

type AccessGroupUpdateResponseResultRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                           `json:"id,required"`
	JSON accessGroupUpdateResponseResultRequireAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultRequireAccessIPListRuleIPList]
type accessGroupUpdateResponseResultRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupUpdateResponseResultRequireAccessCertificateRule struct {
	Certificate interface{}                                                     `json:"certificate,required"`
	JSON        accessGroupUpdateResponseResultRequireAccessCertificateRuleJSON `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultRequireAccessCertificateRule]
type accessGroupUpdateResponseResultRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultRequireAccessCertificateRule) implementsAccessGroupUpdateResponseResultRequire() {
}

// Matches an Access group.
type AccessGroupUpdateResponseResultRequireAccessAccessGroupRule struct {
	Group AccessGroupUpdateResponseResultRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupUpdateResponseResultRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultRequireAccessAccessGroupRule]
type accessGroupUpdateResponseResultRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultRequireAccessAccessGroupRule) implementsAccessGroupUpdateResponseResultRequire() {
}

type AccessGroupUpdateResponseResultRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                               `json:"id,required"`
	JSON accessGroupUpdateResponseResultRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [AccessGroupUpdateResponseResultRequireAccessAccessGroupRuleGroup]
type accessGroupUpdateResponseResultRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupUpdateResponseResultRequireAccessAzureGroupRule struct {
	AzureAd AccessGroupUpdateResponseResultRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupUpdateResponseResultRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseResultRequireAccessAzureGroupRule]
type accessGroupUpdateResponseResultRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultRequireAccessAzureGroupRule) implementsAccessGroupUpdateResponseResultRequire() {
}

type AccessGroupUpdateResponseResultRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                `json:"connection_id,required"`
	JSON         accessGroupUpdateResponseResultRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessAzureGroupRuleAzureAdJSON contains
// the JSON metadata for the struct
// [AccessGroupUpdateResponseResultRequireAccessAzureGroupRuleAzureAd]
type accessGroupUpdateResponseResultRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupUpdateResponseResultRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupUpdateResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupUpdateResponseResultRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessGitHubOrganizationRuleJSON contains
// the JSON metadata for the struct
// [AccessGroupUpdateResponseResultRequireAccessGitHubOrganizationRule]
type accessGroupUpdateResponseResultRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultRequireAccessGitHubOrganizationRule) implementsAccessGroupUpdateResponseResultRequire() {
}

type AccessGroupUpdateResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                   `json:"name,required"`
	JSON accessGroupUpdateResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupUpdateResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupUpdateResponseResultRequireAccessGsuiteGroupRule struct {
	Gsuite AccessGroupUpdateResponseResultRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupUpdateResponseResultRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultRequireAccessGsuiteGroupRule]
type accessGroupUpdateResponseResultRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultRequireAccessGsuiteGroupRule) implementsAccessGroupUpdateResponseResultRequire() {
}

type AccessGroupUpdateResponseResultRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                `json:"email,required"`
	JSON  accessGroupUpdateResponseResultRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessGsuiteGroupRuleGsuiteJSON contains
// the JSON metadata for the struct
// [AccessGroupUpdateResponseResultRequireAccessGsuiteGroupRuleGsuite]
type accessGroupUpdateResponseResultRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupUpdateResponseResultRequireAccessOktaGroupRule struct {
	Okta AccessGroupUpdateResponseResultRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupUpdateResponseResultRequireAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseResultRequireAccessOktaGroupRule]
type accessGroupUpdateResponseResultRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultRequireAccessOktaGroupRule) implementsAccessGroupUpdateResponseResultRequire() {
}

type AccessGroupUpdateResponseResultRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                            `json:"email,required"`
	JSON  accessGroupUpdateResponseResultRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessOktaGroupRuleOktaJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultRequireAccessOktaGroupRuleOkta]
type accessGroupUpdateResponseResultRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupUpdateResponseResultRequireAccessSamlGroupRule struct {
	Saml AccessGroupUpdateResponseResultRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupUpdateResponseResultRequireAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseResultRequireAccessSamlGroupRule]
type accessGroupUpdateResponseResultRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultRequireAccessSamlGroupRule) implementsAccessGroupUpdateResponseResultRequire() {
}

type AccessGroupUpdateResponseResultRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                            `json:"attribute_value,required"`
	JSON           accessGroupUpdateResponseResultRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessSamlGroupRuleSamlJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultRequireAccessSamlGroupRuleSaml]
type accessGroupUpdateResponseResultRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupUpdateResponseResultRequireAccessServiceTokenRule struct {
	ServiceToken AccessGroupUpdateResponseResultRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupUpdateResponseResultRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultRequireAccessServiceTokenRule]
type accessGroupUpdateResponseResultRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultRequireAccessServiceTokenRule) implementsAccessGroupUpdateResponseResultRequire() {
}

type AccessGroupUpdateResponseResultRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                       `json:"token_id,required"`
	JSON    accessGroupUpdateResponseResultRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResultRequireAccessServiceTokenRuleServiceToken]
type accessGroupUpdateResponseResultRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupUpdateResponseResultRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                              `json:"any_valid_service_token,required"`
	JSON                 accessGroupUpdateResponseResultRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResultRequireAccessAnyValidServiceTokenRule]
type accessGroupUpdateResponseResultRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultRequireAccessAnyValidServiceTokenRule) implementsAccessGroupUpdateResponseResultRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupUpdateResponseResultRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupUpdateResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupUpdateResponseResultRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessExternalEvaluationRuleJSON contains
// the JSON metadata for the struct
// [AccessGroupUpdateResponseResultRequireAccessExternalEvaluationRule]
type accessGroupUpdateResponseResultRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultRequireAccessExternalEvaluationRule) implementsAccessGroupUpdateResponseResultRequire() {
}

type AccessGroupUpdateResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                   `json:"keys_url,required"`
	JSON    accessGroupUpdateResponseResultRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupUpdateResponseResultRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupUpdateResponseResultRequireAccessCountryRule struct {
	Geo  AccessGroupUpdateResponseResultRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupUpdateResponseResultRequireAccessCountryRuleJSON `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessCountryRuleJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseResultRequireAccessCountryRule]
type accessGroupUpdateResponseResultRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultRequireAccessCountryRule) implementsAccessGroupUpdateResponseResultRequire() {
}

type AccessGroupUpdateResponseResultRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                         `json:"country_code,required"`
	JSON        accessGroupUpdateResponseResultRequireAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct
// [AccessGroupUpdateResponseResultRequireAccessCountryRuleGeo]
type accessGroupUpdateResponseResultRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupUpdateResponseResultRequireAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupUpdateResponseResultRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupUpdateResponseResultRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResultRequireAccessAuthenticationMethodRule]
type accessGroupUpdateResponseResultRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultRequireAccessAuthenticationMethodRule) implementsAccessGroupUpdateResponseResultRequire() {
}

type AccessGroupUpdateResponseResultRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                             `json:"auth_method,required"`
	JSON       accessGroupUpdateResponseResultRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResultRequireAccessAuthenticationMethodRuleAuthMethod]
type accessGroupUpdateResponseResultRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupUpdateResponseResultRequireAccessDevicePostureRule struct {
	DevicePosture AccessGroupUpdateResponseResultRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupUpdateResponseResultRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessDevicePostureRuleJSON contains the
// JSON metadata for the struct
// [AccessGroupUpdateResponseResultRequireAccessDevicePostureRule]
type accessGroupUpdateResponseResultRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupUpdateResponseResultRequireAccessDevicePostureRule) implementsAccessGroupUpdateResponseResultRequire() {
}

type AccessGroupUpdateResponseResultRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                         `json:"integration_uid,required"`
	JSON           accessGroupUpdateResponseResultRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupUpdateResponseResultRequireAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupUpdateResponseResultRequireAccessDevicePostureRuleDevicePosture]
type accessGroupUpdateResponseResultRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupUpdateResponseResultRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessGroupUpdateResponseSuccess bool

const (
	AccessGroupUpdateResponseSuccessTrue AccessGroupUpdateResponseSuccess = true
)

type AccessGroupDeleteResponse struct {
	Errors   []AccessGroupDeleteResponseError   `json:"errors"`
	Messages []AccessGroupDeleteResponseMessage `json:"messages"`
	Result   AccessGroupDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessGroupDeleteResponseSuccess `json:"success"`
	JSON    accessGroupDeleteResponseJSON    `json:"-"`
}

// accessGroupDeleteResponseJSON contains the JSON metadata for the struct
// [AccessGroupDeleteResponse]
type accessGroupDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupDeleteResponseError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    accessGroupDeleteResponseErrorJSON `json:"-"`
}

// accessGroupDeleteResponseErrorJSON contains the JSON metadata for the struct
// [AccessGroupDeleteResponseError]
type accessGroupDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupDeleteResponseMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    accessGroupDeleteResponseMessageJSON `json:"-"`
}

// accessGroupDeleteResponseMessageJSON contains the JSON metadata for the struct
// [AccessGroupDeleteResponseMessage]
type accessGroupDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupDeleteResponseResult struct {
	// UUID
	ID   string                              `json:"id"`
	JSON accessGroupDeleteResponseResultJSON `json:"-"`
}

// accessGroupDeleteResponseResultJSON contains the JSON metadata for the struct
// [AccessGroupDeleteResponseResult]
type accessGroupDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessGroupDeleteResponseSuccess bool

const (
	AccessGroupDeleteResponseSuccessTrue AccessGroupDeleteResponseSuccess = true
)

type AccessGroupAccessGroupsNewAnAccessGroupResponse struct {
	Errors   []AccessGroupAccessGroupsNewAnAccessGroupResponseError   `json:"errors"`
	Messages []AccessGroupAccessGroupsNewAnAccessGroupResponseMessage `json:"messages"`
	Result   AccessGroupAccessGroupsNewAnAccessGroupResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessGroupAccessGroupsNewAnAccessGroupResponseSuccess `json:"success"`
	JSON    accessGroupAccessGroupsNewAnAccessGroupResponseJSON    `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseJSON contains the JSON metadata
// for the struct [AccessGroupAccessGroupsNewAnAccessGroupResponse]
type accessGroupAccessGroupsNewAnAccessGroupResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseError struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    accessGroupAccessGroupsNewAnAccessGroupResponseErrorJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseErrorJSON contains the JSON
// metadata for the struct [AccessGroupAccessGroupsNewAnAccessGroupResponseError]
type accessGroupAccessGroupsNewAnAccessGroupResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseMessage struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    accessGroupAccessGroupsNewAnAccessGroupResponseMessageJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseMessageJSON contains the JSON
// metadata for the struct [AccessGroupAccessGroupsNewAnAccessGroupResponseMessage]
type accessGroupAccessGroupsNewAnAccessGroupResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResult struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude `json:"include"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	IsDefault []AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault `json:"is_default"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire `json:"require"`
	UpdatedAt time.Time                                                      `json:"updated_at" format:"date-time"`
	JSON      accessGroupAccessGroupsNewAnAccessGroupResponseResultJSON      `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultJSON contains the JSON
// metadata for the struct [AccessGroupAccessGroupsNewAnAccessGroupResponseResult]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultJSON struct {
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

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEmailRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEmailListRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessDomainRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEveryoneRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessIPRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessIPListRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessCertificateRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAccessGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAzureGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessGitHubOrganizationRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessGsuiteGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessOktaGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessSamlGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessServiceTokenRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAnyValidServiceTokenRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessExternalEvaluationRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessCountryRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAuthenticationMethodRule]
// or
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessDevicePostureRule].
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude interface {
	implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEmailRule struct {
	Email AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEmailRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEmailRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEmailRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEmailRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                               `json:"email,required" format:"email"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEmailRuleEmail]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEmailListRule struct {
	EmailList AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEmailListRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEmailListRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                       `json:"id,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEmailListRuleEmailList]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessDomainRule struct {
	EmailDomain AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessDomainRuleJSON        `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessDomainRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessDomainRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessDomainRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                      `json:"domain,required"`
	JSON   accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessDomainRuleEmailDomain]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                        `json:"everyone,required"`
	JSON     accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEveryoneRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessEveryoneRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

// Matches an IP address block.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessIPRule struct {
	IP   AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessIPRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessIPRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessIPRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessIPRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                         `json:"ip,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessIPRuleIPJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessIPRuleIP]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessIPListRule struct {
	IPList AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessIPListRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessIPListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessIPListRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessIPListRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                                 `json:"id,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessIPListRuleIPList]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessCertificateRule struct {
	Certificate interface{}                                                                           `json:"certificate,required"`
	JSON        accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessCertificateRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessCertificateRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessCertificateRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

// Matches an Access group.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAccessGroupRule struct {
	Group AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAccessGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAccessGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                                     `json:"id,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAccessGroupRuleGroup]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAzureGroupRule struct {
	AzureAd AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAzureGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAzureGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                      `json:"connection_id,required"`
	JSON         accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAzureGroupRuleAzureAd]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessGitHubOrganizationRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessGitHubOrganizationRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                         `json:"name,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessGsuiteGroupRule struct {
	Gsuite AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessGsuiteGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessGsuiteGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                      `json:"email,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessGsuiteGroupRuleGsuite]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessOktaGroupRule struct {
	Okta AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessOktaGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessOktaGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                                  `json:"email,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessOktaGroupRuleOkta]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessSamlGroupRule struct {
	Saml AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessSamlGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessSamlGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                                  `json:"attribute_value,required"`
	JSON           accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessSamlGroupRuleSaml]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessServiceTokenRule struct {
	ServiceToken AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessServiceTokenRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessServiceTokenRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                             `json:"token_id,required"`
	JSON    accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessServiceTokenRuleServiceToken]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                                    `json:"any_valid_service_token,required"`
	JSON                 accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAnyValidServiceTokenRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAnyValidServiceTokenRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessExternalEvaluationRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessExternalEvaluationRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                         `json:"keys_url,required"`
	JSON    accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessCountryRule struct {
	Geo  AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessCountryRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessCountryRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessCountryRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                               `json:"country_code,required"`
	JSON        accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessCountryRuleGeo]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAuthenticationMethodRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAuthenticationMethodRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                                   `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessDevicePostureRule struct {
	DevicePosture AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessDevicePostureRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessDevicePostureRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                               `json:"integration_uid,required"`
	JSON           accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessDevicePostureRuleDevicePosture]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEmailRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEmailListRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessDomainRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEveryoneRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessIPRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessIPListRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessCertificateRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAccessGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAzureGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessGitHubOrganizationRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessGsuiteGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessOktaGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessSamlGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessServiceTokenRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAnyValidServiceTokenRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessExternalEvaluationRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessCountryRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAuthenticationMethodRule]
// or
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessDevicePostureRule].
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude interface {
	implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEmailRule struct {
	Email AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEmailRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEmailRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEmailRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEmailRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                               `json:"email,required" format:"email"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEmailRuleEmail]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEmailListRule struct {
	EmailList AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEmailListRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEmailListRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                       `json:"id,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEmailListRuleEmailList]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessDomainRule struct {
	EmailDomain AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessDomainRuleJSON        `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessDomainRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessDomainRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessDomainRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                      `json:"domain,required"`
	JSON   accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessDomainRuleEmailDomain]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                        `json:"everyone,required"`
	JSON     accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEveryoneRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessEveryoneRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

// Matches an IP address block.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessIPRule struct {
	IP   AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessIPRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessIPRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessIPRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessIPRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                         `json:"ip,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessIPRuleIPJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessIPRuleIP]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessIPListRule struct {
	IPList AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessIPListRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessIPListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessIPListRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessIPListRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                                 `json:"id,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessIPListRuleIPList]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessCertificateRule struct {
	Certificate interface{}                                                                           `json:"certificate,required"`
	JSON        accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessCertificateRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessCertificateRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessCertificateRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

// Matches an Access group.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAccessGroupRule struct {
	Group AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAccessGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAccessGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                                     `json:"id,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAccessGroupRuleGroup]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAzureGroupRule struct {
	AzureAd AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAzureGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAzureGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                      `json:"connection_id,required"`
	JSON         accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAzureGroupRuleAzureAd]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessGitHubOrganizationRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessGitHubOrganizationRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                         `json:"name,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessGsuiteGroupRule struct {
	Gsuite AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessGsuiteGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessGsuiteGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                      `json:"email,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessGsuiteGroupRuleGsuite]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessOktaGroupRule struct {
	Okta AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessOktaGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessOktaGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                                  `json:"email,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessOktaGroupRuleOkta]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessSamlGroupRule struct {
	Saml AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessSamlGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessSamlGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                                  `json:"attribute_value,required"`
	JSON           accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessSamlGroupRuleSaml]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessServiceTokenRule struct {
	ServiceToken AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessServiceTokenRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessServiceTokenRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                             `json:"token_id,required"`
	JSON    accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessServiceTokenRuleServiceToken]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                                    `json:"any_valid_service_token,required"`
	JSON                 accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAnyValidServiceTokenRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAnyValidServiceTokenRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessExternalEvaluationRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessExternalEvaluationRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                         `json:"keys_url,required"`
	JSON    accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessCountryRule struct {
	Geo  AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessCountryRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessCountryRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessCountryRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                               `json:"country_code,required"`
	JSON        accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessCountryRuleGeo]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAuthenticationMethodRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAuthenticationMethodRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                                   `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessDevicePostureRule struct {
	DevicePosture AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessDevicePostureRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessDevicePostureRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                               `json:"integration_uid,required"`
	JSON           accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessDevicePostureRuleDevicePosture]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEmailRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEmailListRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessDomainRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEveryoneRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessIPRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessIPListRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessCertificateRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAccessGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAzureGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessGitHubOrganizationRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessGsuiteGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessOktaGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessSamlGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessServiceTokenRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAnyValidServiceTokenRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessExternalEvaluationRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessCountryRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAuthenticationMethodRule]
// or
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessDevicePostureRule].
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault interface {
	implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEmailRule struct {
	Email AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEmailRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEmailRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEmailRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEmailRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                                 `json:"email,required" format:"email"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEmailRuleEmail]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEmailListRule struct {
	EmailList AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEmailListRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEmailListRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                         `json:"id,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEmailListRuleEmailList]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessDomainRule struct {
	EmailDomain AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessDomainRuleJSON        `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessDomainRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessDomainRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessDomainRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                        `json:"domain,required"`
	JSON   accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessDomainRuleEmailDomain]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                          `json:"everyone,required"`
	JSON     accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEveryoneRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessEveryoneRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

// Matches an IP address block.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessIPRule struct {
	IP   AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessIPRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessIPRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessIPRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessIPRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                           `json:"ip,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessIPRuleIPJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessIPRuleIP]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessIPListRule struct {
	IPList AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessIPListRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessIPListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessIPListRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessIPListRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                                   `json:"id,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessIPListRuleIPList]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessCertificateRule struct {
	Certificate interface{}                                                                             `json:"certificate,required"`
	JSON        accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessCertificateRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessCertificateRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessCertificateRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

// Matches an Access group.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAccessGroupRule struct {
	Group AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAccessGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAccessGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                                       `json:"id,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAccessGroupRuleGroup]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAzureGroupRule struct {
	AzureAd AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAzureGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAzureGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                        `json:"connection_id,required"`
	JSON         accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAzureGroupRuleAzureAd]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessGitHubOrganizationRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessGitHubOrganizationRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                           `json:"name,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessGsuiteGroupRule struct {
	Gsuite AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessGsuiteGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessGsuiteGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                        `json:"email,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessGsuiteGroupRuleGsuite]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessOktaGroupRule struct {
	Okta AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessOktaGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessOktaGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                                    `json:"email,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessOktaGroupRuleOkta]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessSamlGroupRule struct {
	Saml AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessSamlGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessSamlGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                                    `json:"attribute_value,required"`
	JSON           accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessSamlGroupRuleSaml]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessServiceTokenRule struct {
	ServiceToken AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessServiceTokenRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessServiceTokenRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                               `json:"token_id,required"`
	JSON    accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessServiceTokenRuleServiceToken]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                                      `json:"any_valid_service_token,required"`
	JSON                 accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAnyValidServiceTokenRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAnyValidServiceTokenRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessExternalEvaluationRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessExternalEvaluationRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                           `json:"keys_url,required"`
	JSON    accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessCountryRule struct {
	Geo  AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessCountryRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessCountryRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessCountryRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                                 `json:"country_code,required"`
	JSON        accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessCountryRuleGeo]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAuthenticationMethodRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAuthenticationMethodRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                                     `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAuthenticationMethodRuleAuthMethod]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessDevicePostureRule struct {
	DevicePosture AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessDevicePostureRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessDevicePostureRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                                 `json:"integration_uid,required"`
	JSON           accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessDevicePostureRuleDevicePosture]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEmailRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEmailListRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessDomainRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEveryoneRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessIPRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessIPListRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessCertificateRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAccessGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAzureGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessGitHubOrganizationRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessGsuiteGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessOktaGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessSamlGroupRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessServiceTokenRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAnyValidServiceTokenRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessExternalEvaluationRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessCountryRule],
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAuthenticationMethodRule]
// or
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessDevicePostureRule].
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire interface {
	implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEmailRule struct {
	Email AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEmailRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEmailRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEmailRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEmailRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                               `json:"email,required" format:"email"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEmailRuleEmail]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEmailListRule struct {
	EmailList AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEmailListRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEmailListRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                       `json:"id,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEmailListRuleEmailList]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessDomainRule struct {
	EmailDomain AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessDomainRuleJSON        `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessDomainRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessDomainRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessDomainRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                      `json:"domain,required"`
	JSON   accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessDomainRuleEmailDomain]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                        `json:"everyone,required"`
	JSON     accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEveryoneRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessEveryoneRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

// Matches an IP address block.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessIPRule struct {
	IP   AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessIPRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessIPRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessIPRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessIPRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                         `json:"ip,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessIPRuleIPJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessIPRuleIP]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessIPListRule struct {
	IPList AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessIPListRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessIPListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessIPListRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessIPListRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                                 `json:"id,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessIPListRuleIPList]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessCertificateRule struct {
	Certificate interface{}                                                                           `json:"certificate,required"`
	JSON        accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessCertificateRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessCertificateRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessCertificateRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

// Matches an Access group.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAccessGroupRule struct {
	Group AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAccessGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAccessGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                                     `json:"id,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAccessGroupRuleGroup]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAzureGroupRule struct {
	AzureAd AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAzureGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAzureGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                      `json:"connection_id,required"`
	JSON         accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAzureGroupRuleAzureAd]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessGitHubOrganizationRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessGitHubOrganizationRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                         `json:"name,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessGsuiteGroupRule struct {
	Gsuite AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessGsuiteGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessGsuiteGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                      `json:"email,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessGsuiteGroupRuleGsuite]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessOktaGroupRule struct {
	Okta AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessOktaGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessOktaGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                                  `json:"email,required"`
	JSON  accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessOktaGroupRuleOkta]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessSamlGroupRule struct {
	Saml AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessSamlGroupRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessSamlGroupRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                                  `json:"attribute_value,required"`
	JSON           accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessSamlGroupRuleSaml]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessServiceTokenRule struct {
	ServiceToken AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessServiceTokenRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessServiceTokenRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                             `json:"token_id,required"`
	JSON    accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessServiceTokenRuleServiceToken]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                                    `json:"any_valid_service_token,required"`
	JSON                 accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAnyValidServiceTokenRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAnyValidServiceTokenRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessExternalEvaluationRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessExternalEvaluationRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                         `json:"keys_url,required"`
	JSON    accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessCountryRule struct {
	Geo  AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessCountryRuleJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessCountryRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessCountryRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                               `json:"country_code,required"`
	JSON        accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessCountryRuleGeo]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAuthenticationMethodRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAuthenticationMethodRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                                   `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAuthenticationMethodRuleAuthMethod]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessDevicePostureRule struct {
	DevicePosture AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessDevicePostureRule]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessDevicePostureRule) implementsAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                               `json:"integration_uid,required"`
	JSON           accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessDevicePostureRuleDevicePosture]
type accessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsNewAnAccessGroupResponseResultRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessGroupAccessGroupsNewAnAccessGroupResponseSuccess bool

const (
	AccessGroupAccessGroupsNewAnAccessGroupResponseSuccessTrue AccessGroupAccessGroupsNewAnAccessGroupResponseSuccess = true
)

type AccessGroupAccessGroupsListAccessGroupsResponse struct {
	Errors     []AccessGroupAccessGroupsListAccessGroupsResponseError    `json:"errors"`
	Messages   []AccessGroupAccessGroupsListAccessGroupsResponseMessage  `json:"messages"`
	Result     []AccessGroupAccessGroupsListAccessGroupsResponseResult   `json:"result"`
	ResultInfo AccessGroupAccessGroupsListAccessGroupsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccessGroupAccessGroupsListAccessGroupsResponseSuccess `json:"success"`
	JSON    accessGroupAccessGroupsListAccessGroupsResponseJSON    `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseJSON contains the JSON metadata
// for the struct [AccessGroupAccessGroupsListAccessGroupsResponse]
type accessGroupAccessGroupsListAccessGroupsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupAccessGroupsListAccessGroupsResponseError struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    accessGroupAccessGroupsListAccessGroupsResponseErrorJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseErrorJSON contains the JSON
// metadata for the struct [AccessGroupAccessGroupsListAccessGroupsResponseError]
type accessGroupAccessGroupsListAccessGroupsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupAccessGroupsListAccessGroupsResponseMessage struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    accessGroupAccessGroupsListAccessGroupsResponseMessageJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseMessageJSON contains the JSON
// metadata for the struct [AccessGroupAccessGroupsListAccessGroupsResponseMessage]
type accessGroupAccessGroupsListAccessGroupsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupAccessGroupsListAccessGroupsResponseResult struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessGroupAccessGroupsListAccessGroupsResponseResultExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessGroupAccessGroupsListAccessGroupsResponseResultInclude `json:"include"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	IsDefault []AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault `json:"is_default"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []AccessGroupAccessGroupsListAccessGroupsResponseResultRequire `json:"require"`
	UpdatedAt time.Time                                                      `json:"updated_at" format:"date-time"`
	JSON      accessGroupAccessGroupsListAccessGroupsResponseResultJSON      `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultJSON contains the JSON
// metadata for the struct [AccessGroupAccessGroupsListAccessGroupsResponseResult]
type accessGroupAccessGroupsListAccessGroupsResponseResultJSON struct {
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

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEmailRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEmailListRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessDomainRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEveryoneRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessIPRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessIPListRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessCertificateRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAccessGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAzureGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessGitHubOrganizationRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessGsuiteGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessOktaGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessSamlGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessServiceTokenRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAnyValidServiceTokenRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessExternalEvaluationRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessCountryRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAuthenticationMethodRule]
// or
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessDevicePostureRule].
type AccessGroupAccessGroupsListAccessGroupsResponseResultExclude interface {
	implementsAccessGroupAccessGroupsListAccessGroupsResponseResultExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupAccessGroupsListAccessGroupsResponseResultExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEmailRule struct {
	Email AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEmailRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEmailRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEmailRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEmailRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                               `json:"email,required" format:"email"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEmailRuleEmail]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEmailListRule struct {
	EmailList AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEmailListRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEmailListRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                       `json:"id,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEmailListRuleEmailList]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessDomainRule struct {
	EmailDomain AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessDomainRuleJSON        `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessDomainRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessDomainRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessDomainRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                      `json:"domain,required"`
	JSON   accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessDomainRuleEmailDomain]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                        `json:"everyone,required"`
	JSON     accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEveryoneRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessEveryoneRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

// Matches an IP address block.
type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessIPRule struct {
	IP   AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessIPRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessIPRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessIPRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessIPRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                         `json:"ip,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessIPRuleIPJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessIPRuleIP]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessIPListRule struct {
	IPList AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessIPListRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessIPListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessIPListRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessIPListRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                                 `json:"id,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessIPListRuleIPList]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessCertificateRule struct {
	Certificate interface{}                                                                           `json:"certificate,required"`
	JSON        accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessCertificateRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessCertificateRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessCertificateRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

// Matches an Access group.
type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAccessGroupRule struct {
	Group AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAccessGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAccessGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                                     `json:"id,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAccessGroupRuleGroup]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAzureGroupRule struct {
	AzureAd AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAzureGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAzureGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                      `json:"connection_id,required"`
	JSON         accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAzureGroupRuleAzureAd]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessGitHubOrganizationRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessGitHubOrganizationRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                         `json:"name,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessGsuiteGroupRule struct {
	Gsuite AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessGsuiteGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessGsuiteGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                      `json:"email,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessGsuiteGroupRuleGsuite]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessOktaGroupRule struct {
	Okta AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessOktaGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessOktaGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                                  `json:"email,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessOktaGroupRuleOkta]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessSamlGroupRule struct {
	Saml AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessSamlGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessSamlGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                                  `json:"attribute_value,required"`
	JSON           accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessSamlGroupRuleSaml]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessServiceTokenRule struct {
	ServiceToken AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessServiceTokenRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessServiceTokenRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                             `json:"token_id,required"`
	JSON    accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessServiceTokenRuleServiceToken]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                                    `json:"any_valid_service_token,required"`
	JSON                 accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAnyValidServiceTokenRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAnyValidServiceTokenRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessExternalEvaluationRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessExternalEvaluationRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                         `json:"keys_url,required"`
	JSON    accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessCountryRule struct {
	Geo  AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessCountryRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessCountryRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessCountryRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                               `json:"country_code,required"`
	JSON        accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessCountryRuleGeo]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAuthenticationMethodRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAuthenticationMethodRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                                   `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessDevicePostureRule struct {
	DevicePosture AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessDevicePostureRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessDevicePostureRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                               `json:"integration_uid,required"`
	JSON           accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessDevicePostureRuleDevicePosture]
type accessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEmailRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEmailListRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessDomainRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEveryoneRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessIPRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessIPListRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessCertificateRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAccessGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAzureGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessGitHubOrganizationRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessGsuiteGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessOktaGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessSamlGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessServiceTokenRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAnyValidServiceTokenRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessExternalEvaluationRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessCountryRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAuthenticationMethodRule]
// or
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessDevicePostureRule].
type AccessGroupAccessGroupsListAccessGroupsResponseResultInclude interface {
	implementsAccessGroupAccessGroupsListAccessGroupsResponseResultInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupAccessGroupsListAccessGroupsResponseResultInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEmailRule struct {
	Email AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEmailRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEmailRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEmailRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEmailRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                               `json:"email,required" format:"email"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEmailRuleEmail]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEmailListRule struct {
	EmailList AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEmailListRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEmailListRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                       `json:"id,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEmailListRuleEmailList]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessDomainRule struct {
	EmailDomain AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessDomainRuleJSON        `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessDomainRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessDomainRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessDomainRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                      `json:"domain,required"`
	JSON   accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessDomainRuleEmailDomain]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                        `json:"everyone,required"`
	JSON     accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEveryoneRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessEveryoneRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

// Matches an IP address block.
type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessIPRule struct {
	IP   AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessIPRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessIPRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessIPRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessIPRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                         `json:"ip,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessIPRuleIPJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessIPRuleIP]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessIPListRule struct {
	IPList AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessIPListRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessIPListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessIPListRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessIPListRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                                 `json:"id,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessIPListRuleIPList]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessCertificateRule struct {
	Certificate interface{}                                                                           `json:"certificate,required"`
	JSON        accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessCertificateRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessCertificateRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessCertificateRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

// Matches an Access group.
type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAccessGroupRule struct {
	Group AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAccessGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAccessGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                                     `json:"id,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAccessGroupRuleGroup]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAzureGroupRule struct {
	AzureAd AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAzureGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAzureGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                      `json:"connection_id,required"`
	JSON         accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAzureGroupRuleAzureAd]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessGitHubOrganizationRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessGitHubOrganizationRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                         `json:"name,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessGsuiteGroupRule struct {
	Gsuite AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessGsuiteGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessGsuiteGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                      `json:"email,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessGsuiteGroupRuleGsuite]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessOktaGroupRule struct {
	Okta AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessOktaGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessOktaGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                                  `json:"email,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessOktaGroupRuleOkta]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessSamlGroupRule struct {
	Saml AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessSamlGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessSamlGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                                  `json:"attribute_value,required"`
	JSON           accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessSamlGroupRuleSaml]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessServiceTokenRule struct {
	ServiceToken AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessServiceTokenRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessServiceTokenRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                             `json:"token_id,required"`
	JSON    accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessServiceTokenRuleServiceToken]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                                    `json:"any_valid_service_token,required"`
	JSON                 accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAnyValidServiceTokenRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAnyValidServiceTokenRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessExternalEvaluationRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessExternalEvaluationRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                         `json:"keys_url,required"`
	JSON    accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessCountryRule struct {
	Geo  AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessCountryRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessCountryRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessCountryRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                               `json:"country_code,required"`
	JSON        accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessCountryRuleGeo]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAuthenticationMethodRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAuthenticationMethodRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                                   `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessDevicePostureRule struct {
	DevicePosture AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessDevicePostureRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessDevicePostureRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                               `json:"integration_uid,required"`
	JSON           accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessDevicePostureRuleDevicePosture]
type accessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEmailRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEmailListRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessDomainRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEveryoneRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessIPRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessIPListRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessCertificateRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAccessGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAzureGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessGitHubOrganizationRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessGsuiteGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessOktaGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessSamlGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessServiceTokenRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAnyValidServiceTokenRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessExternalEvaluationRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessCountryRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAuthenticationMethodRule]
// or
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessDevicePostureRule].
type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault interface {
	implementsAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEmailRule struct {
	Email AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEmailRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEmailRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEmailRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEmailRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                                 `json:"email,required" format:"email"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEmailRuleEmail]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEmailListRule struct {
	EmailList AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEmailListRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEmailListRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                         `json:"id,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEmailListRuleEmailList]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessDomainRule struct {
	EmailDomain AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessDomainRuleJSON        `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessDomainRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessDomainRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessDomainRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                        `json:"domain,required"`
	JSON   accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessDomainRuleEmailDomain]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                          `json:"everyone,required"`
	JSON     accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEveryoneRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessEveryoneRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

// Matches an IP address block.
type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessIPRule struct {
	IP   AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessIPRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessIPRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessIPRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessIPRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                           `json:"ip,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessIPRuleIPJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessIPRuleIP]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessIPListRule struct {
	IPList AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessIPListRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessIPListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessIPListRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessIPListRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                                   `json:"id,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessIPListRuleIPList]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessCertificateRule struct {
	Certificate interface{}                                                                             `json:"certificate,required"`
	JSON        accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessCertificateRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessCertificateRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessCertificateRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

// Matches an Access group.
type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAccessGroupRule struct {
	Group AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAccessGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAccessGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                                       `json:"id,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAccessGroupRuleGroup]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAzureGroupRule struct {
	AzureAd AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAzureGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAzureGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                        `json:"connection_id,required"`
	JSON         accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAzureGroupRuleAzureAd]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessGitHubOrganizationRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessGitHubOrganizationRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                           `json:"name,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessGsuiteGroupRule struct {
	Gsuite AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessGsuiteGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessGsuiteGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                        `json:"email,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessGsuiteGroupRuleGsuite]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessOktaGroupRule struct {
	Okta AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessOktaGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessOktaGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                                    `json:"email,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessOktaGroupRuleOkta]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessSamlGroupRule struct {
	Saml AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessSamlGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessSamlGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                                    `json:"attribute_value,required"`
	JSON           accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessSamlGroupRuleSaml]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessServiceTokenRule struct {
	ServiceToken AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessServiceTokenRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessServiceTokenRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                               `json:"token_id,required"`
	JSON    accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessServiceTokenRuleServiceToken]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                                      `json:"any_valid_service_token,required"`
	JSON                 accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAnyValidServiceTokenRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAnyValidServiceTokenRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessExternalEvaluationRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessExternalEvaluationRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                           `json:"keys_url,required"`
	JSON    accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessCountryRule struct {
	Geo  AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessCountryRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessCountryRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessCountryRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                                 `json:"country_code,required"`
	JSON        accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessCountryRuleGeo]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAuthenticationMethodRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAuthenticationMethodRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                                     `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAuthenticationMethodRuleAuthMethod]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessDevicePostureRule struct {
	DevicePosture AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessDevicePostureRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessDevicePostureRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                                 `json:"integration_uid,required"`
	JSON           accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessDevicePostureRuleDevicePosture]
type accessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEmailRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEmailListRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessDomainRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEveryoneRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessIPRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessIPListRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessCertificateRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAccessGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAzureGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessGitHubOrganizationRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessGsuiteGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessOktaGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessSamlGroupRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessServiceTokenRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAnyValidServiceTokenRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessExternalEvaluationRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessCountryRule],
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAuthenticationMethodRule]
// or
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessDevicePostureRule].
type AccessGroupAccessGroupsListAccessGroupsResponseResultRequire interface {
	implementsAccessGroupAccessGroupsListAccessGroupsResponseResultRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupAccessGroupsListAccessGroupsResponseResultRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEmailRule struct {
	Email AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEmailRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEmailRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEmailRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEmailRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                               `json:"email,required" format:"email"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEmailRuleEmail]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEmailListRule struct {
	EmailList AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEmailListRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEmailListRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                       `json:"id,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEmailListRuleEmailList]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessDomainRule struct {
	EmailDomain AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessDomainRuleJSON        `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessDomainRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessDomainRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessDomainRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                      `json:"domain,required"`
	JSON   accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessDomainRuleEmailDomain]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                        `json:"everyone,required"`
	JSON     accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEveryoneRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessEveryoneRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

// Matches an IP address block.
type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessIPRule struct {
	IP   AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessIPRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessIPRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessIPRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessIPRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                         `json:"ip,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessIPRuleIPJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessIPRuleIP]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessIPListRule struct {
	IPList AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessIPListRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessIPListRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessIPListRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessIPListRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                                 `json:"id,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessIPListRuleIPList]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessCertificateRule struct {
	Certificate interface{}                                                                           `json:"certificate,required"`
	JSON        accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessCertificateRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessCertificateRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessCertificateRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

// Matches an Access group.
type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAccessGroupRule struct {
	Group AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAccessGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAccessGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                                     `json:"id,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAccessGroupRuleGroup]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAzureGroupRule struct {
	AzureAd AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAzureGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAzureGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                      `json:"connection_id,required"`
	JSON         accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAzureGroupRuleAzureAd]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessGitHubOrganizationRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessGitHubOrganizationRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                         `json:"name,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessGsuiteGroupRule struct {
	Gsuite AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessGsuiteGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessGsuiteGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                      `json:"email,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessGsuiteGroupRuleGsuite]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessOktaGroupRule struct {
	Okta AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessOktaGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessOktaGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                                  `json:"email,required"`
	JSON  accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessOktaGroupRuleOkta]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessSamlGroupRule struct {
	Saml AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessSamlGroupRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessSamlGroupRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                                  `json:"attribute_value,required"`
	JSON           accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessSamlGroupRuleSaml]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessServiceTokenRule struct {
	ServiceToken AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessServiceTokenRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessServiceTokenRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                             `json:"token_id,required"`
	JSON    accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessServiceTokenRuleServiceToken]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                                    `json:"any_valid_service_token,required"`
	JSON                 accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAnyValidServiceTokenRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAnyValidServiceTokenRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessExternalEvaluationRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessExternalEvaluationRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                         `json:"keys_url,required"`
	JSON    accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessCountryRule struct {
	Geo  AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessCountryRuleJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessCountryRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessCountryRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessCountryRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                               `json:"country_code,required"`
	JSON        accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessCountryRuleGeo]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAuthenticationMethodRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAuthenticationMethodRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                                   `json:"auth_method,required"`
	JSON       accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAuthenticationMethodRuleAuthMethod]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessDevicePostureRule struct {
	DevicePosture AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessDevicePostureRule]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessDevicePostureRule) implementsAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                               `json:"integration_uid,required"`
	JSON           accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessDevicePostureRuleDevicePosture]
type accessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessGroupAccessGroupsListAccessGroupsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                       `json:"total_count"`
	JSON       accessGroupAccessGroupsListAccessGroupsResponseResultInfoJSON `json:"-"`
}

// accessGroupAccessGroupsListAccessGroupsResponseResultInfoJSON contains the JSON
// metadata for the struct
// [AccessGroupAccessGroupsListAccessGroupsResponseResultInfo]
type accessGroupAccessGroupsListAccessGroupsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupAccessGroupsListAccessGroupsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessGroupAccessGroupsListAccessGroupsResponseSuccess bool

const (
	AccessGroupAccessGroupsListAccessGroupsResponseSuccessTrue AccessGroupAccessGroupsListAccessGroupsResponseSuccess = true
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
