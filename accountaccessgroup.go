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

// AccountAccessGroupService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountAccessGroupService] method
// instead.
type AccountAccessGroupService struct {
	Options []option.RequestOption
}

// NewAccountAccessGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAccessGroupService(opts ...option.RequestOption) (r *AccountAccessGroupService) {
	r = &AccountAccessGroupService{}
	r.Options = opts
	return
}

// Fetches a single Access group.
func (r *AccountAccessGroupService) Get(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *AccountAccessGroupGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/groups/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured Access group.
func (r *AccountAccessGroupService) Update(ctx context.Context, identifier string, uuid string, body AccountAccessGroupUpdateParams, opts ...option.RequestOption) (res *AccountAccessGroupUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/groups/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes an Access group.
func (r *AccountAccessGroupService) Delete(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *AccountAccessGroupDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/groups/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a new Access group.
func (r *AccountAccessGroupService) AccessGroupsNewAnAccessGroup(ctx context.Context, identifier string, body AccountAccessGroupAccessGroupsNewAnAccessGroupParams, opts ...option.RequestOption) (res *AccountAccessGroupAccessGroupsNewAnAccessGroupResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/groups", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all Access groups.
func (r *AccountAccessGroupService) AccessGroupsListAccessGroups(ctx context.Context, identifier string, opts ...option.RequestOption) (res *AccountAccessGroupAccessGroupsListAccessGroupsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/groups", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountAccessGroupGetResponse struct {
	Errors   []AccountAccessGroupGetResponseError   `json:"errors"`
	Messages []AccountAccessGroupGetResponseMessage `json:"messages"`
	Result   AccountAccessGroupGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessGroupGetResponseSuccess `json:"success"`
	JSON    accountAccessGroupGetResponseJSON    `json:"-"`
}

// accountAccessGroupGetResponseJSON contains the JSON metadata for the struct
// [AccountAccessGroupGetResponse]
type accountAccessGroupGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessGroupGetResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    accountAccessGroupGetResponseErrorJSON `json:"-"`
}

// accountAccessGroupGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseError]
type accountAccessGroupGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessGroupGetResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accountAccessGroupGetResponseMessageJSON `json:"-"`
}

// accountAccessGroupGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountAccessGroupGetResponseMessage]
type accountAccessGroupGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessGroupGetResponseResult struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccountAccessGroupGetResponseResultExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccountAccessGroupGetResponseResultInclude `json:"include"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	IsDefault []AccountAccessGroupGetResponseResultIsDefault `json:"is_default"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []AccountAccessGroupGetResponseResultRequire `json:"require"`
	UpdatedAt time.Time                                    `json:"updated_at" format:"date-time"`
	JSON      accountAccessGroupGetResponseResultJSON      `json:"-"`
}

// accountAccessGroupGetResponseResultJSON contains the JSON metadata for the
// struct [AccountAccessGroupGetResponseResult]
type accountAccessGroupGetResponseResultJSON struct {
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

func (r *AccountAccessGroupGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccountAccessGroupGetResponseResultExcludePajwohLqEmailRule],
// [AccountAccessGroupGetResponseResultExcludePajwohLqEmailListRule],
// [AccountAccessGroupGetResponseResultExcludePajwohLqDomainRule],
// [AccountAccessGroupGetResponseResultExcludePajwohLqEveryoneRule],
// [AccountAccessGroupGetResponseResultExcludePajwohLqIPRule],
// [AccountAccessGroupGetResponseResultExcludePajwohLqIPListRule],
// [AccountAccessGroupGetResponseResultExcludePajwohLqCertificateRule],
// [AccountAccessGroupGetResponseResultExcludePajwohLqAccessGroupRule],
// [AccountAccessGroupGetResponseResultExcludePajwohLqAzureGroupRule],
// [AccountAccessGroupGetResponseResultExcludePajwohLqGitHubOrganizationRule],
// [AccountAccessGroupGetResponseResultExcludePajwohLqGsuiteGroupRule],
// [AccountAccessGroupGetResponseResultExcludePajwohLqOktaGroupRule],
// [AccountAccessGroupGetResponseResultExcludePajwohLqSamlGroupRule],
// [AccountAccessGroupGetResponseResultExcludePajwohLqServiceTokenRule],
// [AccountAccessGroupGetResponseResultExcludePajwohLqAnyValidServiceTokenRule],
// [AccountAccessGroupGetResponseResultExcludePajwohLqExternalEvaluationRule],
// [AccountAccessGroupGetResponseResultExcludePajwohLqCountryRule],
// [AccountAccessGroupGetResponseResultExcludePajwohLqAuthenticationMethodRule] or
// [AccountAccessGroupGetResponseResultExcludePajwohLqDevicePostureRule].
type AccountAccessGroupGetResponseResultExclude interface {
	implementsAccountAccessGroupGetResponseResultExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountAccessGroupGetResponseResultExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccountAccessGroupGetResponseResultExcludePajwohLqEmailRule struct {
	Email AccountAccessGroupGetResponseResultExcludePajwohLqEmailRuleEmail `json:"email,required"`
	JSON  accountAccessGroupGetResponseResultExcludePajwohLqEmailRuleJSON  `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqEmailRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqEmailRule]
type accountAccessGroupGetResponseResultExcludePajwohLqEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultExcludePajwohLqEmailRule) implementsAccountAccessGroupGetResponseResultExclude() {
}

type AccountAccessGroupGetResponseResultExcludePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email string                                                               `json:"email,required" format:"email"`
	JSON  accountAccessGroupGetResponseResultExcludePajwohLqEmailRuleEmailJSON `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqEmailRuleEmailJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqEmailRuleEmail]
type accountAccessGroupGetResponseResultExcludePajwohLqEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccountAccessGroupGetResponseResultExcludePajwohLqEmailListRule struct {
	EmailList AccountAccessGroupGetResponseResultExcludePajwohLqEmailListRuleEmailList `json:"email_list,required"`
	JSON      accountAccessGroupGetResponseResultExcludePajwohLqEmailListRuleJSON      `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqEmailListRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqEmailListRule]
type accountAccessGroupGetResponseResultExcludePajwohLqEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultExcludePajwohLqEmailListRule) implementsAccountAccessGroupGetResponseResultExclude() {
}

type AccountAccessGroupGetResponseResultExcludePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                       `json:"id,required"`
	JSON accountAccessGroupGetResponseResultExcludePajwohLqEmailListRuleEmailListJSON `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqEmailListRuleEmailList]
type accountAccessGroupGetResponseResultExcludePajwohLqEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccountAccessGroupGetResponseResultExcludePajwohLqDomainRule struct {
	EmailDomain AccountAccessGroupGetResponseResultExcludePajwohLqDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accountAccessGroupGetResponseResultExcludePajwohLqDomainRuleJSON        `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqDomainRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqDomainRule]
type accountAccessGroupGetResponseResultExcludePajwohLqDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultExcludePajwohLqDomainRule) implementsAccountAccessGroupGetResponseResultExclude() {
}

type AccountAccessGroupGetResponseResultExcludePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                      `json:"domain,required"`
	JSON   accountAccessGroupGetResponseResultExcludePajwohLqDomainRuleEmailDomainJSON `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqDomainRuleEmailDomain]
type accountAccessGroupGetResponseResultExcludePajwohLqDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccountAccessGroupGetResponseResultExcludePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                        `json:"everyone,required"`
	JSON     accountAccessGroupGetResponseResultExcludePajwohLqEveryoneRuleJSON `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqEveryoneRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqEveryoneRule]
type accountAccessGroupGetResponseResultExcludePajwohLqEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultExcludePajwohLqEveryoneRule) implementsAccountAccessGroupGetResponseResultExclude() {
}

// Matches an IP address block.
type AccountAccessGroupGetResponseResultExcludePajwohLqIPRule struct {
	IP   AccountAccessGroupGetResponseResultExcludePajwohLqIPRuleIP   `json:"ip,required"`
	JSON accountAccessGroupGetResponseResultExcludePajwohLqIPRuleJSON `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqIPRuleJSON contains the JSON
// metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqIPRule]
type accountAccessGroupGetResponseResultExcludePajwohLqIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultExcludePajwohLqIPRule) implementsAccountAccessGroupGetResponseResultExclude() {
}

type AccountAccessGroupGetResponseResultExcludePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                         `json:"ip,required"`
	JSON accountAccessGroupGetResponseResultExcludePajwohLqIPRuleIPJSON `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqIPRuleIPJSON contains the JSON
// metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqIPRuleIP]
type accountAccessGroupGetResponseResultExcludePajwohLqIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccountAccessGroupGetResponseResultExcludePajwohLqIPListRule struct {
	IPList AccountAccessGroupGetResponseResultExcludePajwohLqIPListRuleIPList `json:"ip_list,required"`
	JSON   accountAccessGroupGetResponseResultExcludePajwohLqIPListRuleJSON   `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqIPListRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqIPListRule]
type accountAccessGroupGetResponseResultExcludePajwohLqIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultExcludePajwohLqIPListRule) implementsAccountAccessGroupGetResponseResultExclude() {
}

type AccountAccessGroupGetResponseResultExcludePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                 `json:"id,required"`
	JSON accountAccessGroupGetResponseResultExcludePajwohLqIPListRuleIPListJSON `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqIPListRuleIPListJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqIPListRuleIPList]
type accountAccessGroupGetResponseResultExcludePajwohLqIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccountAccessGroupGetResponseResultExcludePajwohLqCertificateRule struct {
	Certificate interface{}                                                           `json:"certificate,required"`
	JSON        accountAccessGroupGetResponseResultExcludePajwohLqCertificateRuleJSON `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqCertificateRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqCertificateRule]
type accountAccessGroupGetResponseResultExcludePajwohLqCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultExcludePajwohLqCertificateRule) implementsAccountAccessGroupGetResponseResultExclude() {
}

// Matches an Access group.
type AccountAccessGroupGetResponseResultExcludePajwohLqAccessGroupRule struct {
	Group AccountAccessGroupGetResponseResultExcludePajwohLqAccessGroupRuleGroup `json:"group,required"`
	JSON  accountAccessGroupGetResponseResultExcludePajwohLqAccessGroupRuleJSON  `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqAccessGroupRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqAccessGroupRule]
type accountAccessGroupGetResponseResultExcludePajwohLqAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultExcludePajwohLqAccessGroupRule) implementsAccountAccessGroupGetResponseResultExclude() {
}

type AccountAccessGroupGetResponseResultExcludePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                     `json:"id,required"`
	JSON accountAccessGroupGetResponseResultExcludePajwohLqAccessGroupRuleGroupJSON `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqAccessGroupRuleGroup]
type accountAccessGroupGetResponseResultExcludePajwohLqAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccountAccessGroupGetResponseResultExcludePajwohLqAzureGroupRule struct {
	AzureAd AccountAccessGroupGetResponseResultExcludePajwohLqAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accountAccessGroupGetResponseResultExcludePajwohLqAzureGroupRuleJSON    `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqAzureGroupRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqAzureGroupRule]
type accountAccessGroupGetResponseResultExcludePajwohLqAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultExcludePajwohLqAzureGroupRule) implementsAccountAccessGroupGetResponseResultExclude() {
}

type AccountAccessGroupGetResponseResultExcludePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                      `json:"connection_id,required"`
	JSON         accountAccessGroupGetResponseResultExcludePajwohLqAzureGroupRuleAzureAdJSON `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqAzureGroupRuleAzureAd]
type accountAccessGroupGetResponseResultExcludePajwohLqAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccountAccessGroupGetResponseResultExcludePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization AccountAccessGroupGetResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accountAccessGroupGetResponseResultExcludePajwohLqGitHubOrganizationRuleJSON               `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqGitHubOrganizationRule]
type accountAccessGroupGetResponseResultExcludePajwohLqGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultExcludePajwohLqGitHubOrganizationRule) implementsAccountAccessGroupGetResponseResultExclude() {
}

type AccountAccessGroupGetResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                         `json:"name,required"`
	JSON accountAccessGroupGetResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization]
type accountAccessGroupGetResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccountAccessGroupGetResponseResultExcludePajwohLqGsuiteGroupRule struct {
	Gsuite AccountAccessGroupGetResponseResultExcludePajwohLqGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accountAccessGroupGetResponseResultExcludePajwohLqGsuiteGroupRuleJSON   `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqGsuiteGroupRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqGsuiteGroupRule]
type accountAccessGroupGetResponseResultExcludePajwohLqGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultExcludePajwohLqGsuiteGroupRule) implementsAccountAccessGroupGetResponseResultExclude() {
}

type AccountAccessGroupGetResponseResultExcludePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                      `json:"email,required"`
	JSON  accountAccessGroupGetResponseResultExcludePajwohLqGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqGsuiteGroupRuleGsuite]
type accountAccessGroupGetResponseResultExcludePajwohLqGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccountAccessGroupGetResponseResultExcludePajwohLqOktaGroupRule struct {
	Okta AccountAccessGroupGetResponseResultExcludePajwohLqOktaGroupRuleOkta `json:"okta,required"`
	JSON accountAccessGroupGetResponseResultExcludePajwohLqOktaGroupRuleJSON `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqOktaGroupRule]
type accountAccessGroupGetResponseResultExcludePajwohLqOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultExcludePajwohLqOktaGroupRule) implementsAccountAccessGroupGetResponseResultExclude() {
}

type AccountAccessGroupGetResponseResultExcludePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                  `json:"email,required"`
	JSON  accountAccessGroupGetResponseResultExcludePajwohLqOktaGroupRuleOktaJSON `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqOktaGroupRuleOkta]
type accountAccessGroupGetResponseResultExcludePajwohLqOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccountAccessGroupGetResponseResultExcludePajwohLqSamlGroupRule struct {
	Saml AccountAccessGroupGetResponseResultExcludePajwohLqSamlGroupRuleSaml `json:"saml,required"`
	JSON accountAccessGroupGetResponseResultExcludePajwohLqSamlGroupRuleJSON `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqSamlGroupRule]
type accountAccessGroupGetResponseResultExcludePajwohLqSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultExcludePajwohLqSamlGroupRule) implementsAccountAccessGroupGetResponseResultExclude() {
}

type AccountAccessGroupGetResponseResultExcludePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                  `json:"attribute_value,required"`
	JSON           accountAccessGroupGetResponseResultExcludePajwohLqSamlGroupRuleSamlJSON `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqSamlGroupRuleSaml]
type accountAccessGroupGetResponseResultExcludePajwohLqSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccountAccessGroupGetResponseResultExcludePajwohLqServiceTokenRule struct {
	ServiceToken AccountAccessGroupGetResponseResultExcludePajwohLqServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accountAccessGroupGetResponseResultExcludePajwohLqServiceTokenRuleJSON         `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqServiceTokenRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqServiceTokenRule]
type accountAccessGroupGetResponseResultExcludePajwohLqServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultExcludePajwohLqServiceTokenRule) implementsAccountAccessGroupGetResponseResultExclude() {
}

type AccountAccessGroupGetResponseResultExcludePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                             `json:"token_id,required"`
	JSON    accountAccessGroupGetResponseResultExcludePajwohLqServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqServiceTokenRuleServiceToken]
type accountAccessGroupGetResponseResultExcludePajwohLqServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccountAccessGroupGetResponseResultExcludePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                    `json:"any_valid_service_token,required"`
	JSON                 accountAccessGroupGetResponseResultExcludePajwohLqAnyValidServiceTokenRuleJSON `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqAnyValidServiceTokenRule]
type accountAccessGroupGetResponseResultExcludePajwohLqAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultExcludePajwohLqAnyValidServiceTokenRule) implementsAccountAccessGroupGetResponseResultExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccountAccessGroupGetResponseResultExcludePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation AccountAccessGroupGetResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accountAccessGroupGetResponseResultExcludePajwohLqExternalEvaluationRuleJSON               `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqExternalEvaluationRule]
type accountAccessGroupGetResponseResultExcludePajwohLqExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultExcludePajwohLqExternalEvaluationRule) implementsAccountAccessGroupGetResponseResultExclude() {
}

type AccountAccessGroupGetResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                         `json:"keys_url,required"`
	JSON    accountAccessGroupGetResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation]
type accountAccessGroupGetResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccountAccessGroupGetResponseResultExcludePajwohLqCountryRule struct {
	Geo  AccountAccessGroupGetResponseResultExcludePajwohLqCountryRuleGeo  `json:"geo,required"`
	JSON accountAccessGroupGetResponseResultExcludePajwohLqCountryRuleJSON `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqCountryRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqCountryRule]
type accountAccessGroupGetResponseResultExcludePajwohLqCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultExcludePajwohLqCountryRule) implementsAccountAccessGroupGetResponseResultExclude() {
}

type AccountAccessGroupGetResponseResultExcludePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                               `json:"country_code,required"`
	JSON        accountAccessGroupGetResponseResultExcludePajwohLqCountryRuleGeoJSON `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqCountryRuleGeoJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqCountryRuleGeo]
type accountAccessGroupGetResponseResultExcludePajwohLqCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccountAccessGroupGetResponseResultExcludePajwohLqAuthenticationMethodRule struct {
	AuthMethod AccountAccessGroupGetResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accountAccessGroupGetResponseResultExcludePajwohLqAuthenticationMethodRuleJSON       `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqAuthenticationMethodRule]
type accountAccessGroupGetResponseResultExcludePajwohLqAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultExcludePajwohLqAuthenticationMethodRule) implementsAccountAccessGroupGetResponseResultExclude() {
}

type AccountAccessGroupGetResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                   `json:"auth_method,required"`
	JSON       accountAccessGroupGetResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod]
type accountAccessGroupGetResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccountAccessGroupGetResponseResultExcludePajwohLqDevicePostureRule struct {
	DevicePosture AccountAccessGroupGetResponseResultExcludePajwohLqDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accountAccessGroupGetResponseResultExcludePajwohLqDevicePostureRuleJSON          `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqDevicePostureRule]
type accountAccessGroupGetResponseResultExcludePajwohLqDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultExcludePajwohLqDevicePostureRule) implementsAccountAccessGroupGetResponseResultExclude() {
}

type AccountAccessGroupGetResponseResultExcludePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                               `json:"integration_uid,required"`
	JSON           accountAccessGroupGetResponseResultExcludePajwohLqDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accountAccessGroupGetResponseResultExcludePajwohLqDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultExcludePajwohLqDevicePostureRuleDevicePosture]
type accountAccessGroupGetResponseResultExcludePajwohLqDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultExcludePajwohLqDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccountAccessGroupGetResponseResultIncludePajwohLqEmailRule],
// [AccountAccessGroupGetResponseResultIncludePajwohLqEmailListRule],
// [AccountAccessGroupGetResponseResultIncludePajwohLqDomainRule],
// [AccountAccessGroupGetResponseResultIncludePajwohLqEveryoneRule],
// [AccountAccessGroupGetResponseResultIncludePajwohLqIPRule],
// [AccountAccessGroupGetResponseResultIncludePajwohLqIPListRule],
// [AccountAccessGroupGetResponseResultIncludePajwohLqCertificateRule],
// [AccountAccessGroupGetResponseResultIncludePajwohLqAccessGroupRule],
// [AccountAccessGroupGetResponseResultIncludePajwohLqAzureGroupRule],
// [AccountAccessGroupGetResponseResultIncludePajwohLqGitHubOrganizationRule],
// [AccountAccessGroupGetResponseResultIncludePajwohLqGsuiteGroupRule],
// [AccountAccessGroupGetResponseResultIncludePajwohLqOktaGroupRule],
// [AccountAccessGroupGetResponseResultIncludePajwohLqSamlGroupRule],
// [AccountAccessGroupGetResponseResultIncludePajwohLqServiceTokenRule],
// [AccountAccessGroupGetResponseResultIncludePajwohLqAnyValidServiceTokenRule],
// [AccountAccessGroupGetResponseResultIncludePajwohLqExternalEvaluationRule],
// [AccountAccessGroupGetResponseResultIncludePajwohLqCountryRule],
// [AccountAccessGroupGetResponseResultIncludePajwohLqAuthenticationMethodRule] or
// [AccountAccessGroupGetResponseResultIncludePajwohLqDevicePostureRule].
type AccountAccessGroupGetResponseResultInclude interface {
	implementsAccountAccessGroupGetResponseResultInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountAccessGroupGetResponseResultInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccountAccessGroupGetResponseResultIncludePajwohLqEmailRule struct {
	Email AccountAccessGroupGetResponseResultIncludePajwohLqEmailRuleEmail `json:"email,required"`
	JSON  accountAccessGroupGetResponseResultIncludePajwohLqEmailRuleJSON  `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqEmailRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqEmailRule]
type accountAccessGroupGetResponseResultIncludePajwohLqEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIncludePajwohLqEmailRule) implementsAccountAccessGroupGetResponseResultInclude() {
}

type AccountAccessGroupGetResponseResultIncludePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email string                                                               `json:"email,required" format:"email"`
	JSON  accountAccessGroupGetResponseResultIncludePajwohLqEmailRuleEmailJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqEmailRuleEmailJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqEmailRuleEmail]
type accountAccessGroupGetResponseResultIncludePajwohLqEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccountAccessGroupGetResponseResultIncludePajwohLqEmailListRule struct {
	EmailList AccountAccessGroupGetResponseResultIncludePajwohLqEmailListRuleEmailList `json:"email_list,required"`
	JSON      accountAccessGroupGetResponseResultIncludePajwohLqEmailListRuleJSON      `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqEmailListRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqEmailListRule]
type accountAccessGroupGetResponseResultIncludePajwohLqEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIncludePajwohLqEmailListRule) implementsAccountAccessGroupGetResponseResultInclude() {
}

type AccountAccessGroupGetResponseResultIncludePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                       `json:"id,required"`
	JSON accountAccessGroupGetResponseResultIncludePajwohLqEmailListRuleEmailListJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqEmailListRuleEmailList]
type accountAccessGroupGetResponseResultIncludePajwohLqEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccountAccessGroupGetResponseResultIncludePajwohLqDomainRule struct {
	EmailDomain AccountAccessGroupGetResponseResultIncludePajwohLqDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accountAccessGroupGetResponseResultIncludePajwohLqDomainRuleJSON        `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqDomainRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqDomainRule]
type accountAccessGroupGetResponseResultIncludePajwohLqDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIncludePajwohLqDomainRule) implementsAccountAccessGroupGetResponseResultInclude() {
}

type AccountAccessGroupGetResponseResultIncludePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                      `json:"domain,required"`
	JSON   accountAccessGroupGetResponseResultIncludePajwohLqDomainRuleEmailDomainJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqDomainRuleEmailDomain]
type accountAccessGroupGetResponseResultIncludePajwohLqDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccountAccessGroupGetResponseResultIncludePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                        `json:"everyone,required"`
	JSON     accountAccessGroupGetResponseResultIncludePajwohLqEveryoneRuleJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqEveryoneRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqEveryoneRule]
type accountAccessGroupGetResponseResultIncludePajwohLqEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIncludePajwohLqEveryoneRule) implementsAccountAccessGroupGetResponseResultInclude() {
}

// Matches an IP address block.
type AccountAccessGroupGetResponseResultIncludePajwohLqIPRule struct {
	IP   AccountAccessGroupGetResponseResultIncludePajwohLqIPRuleIP   `json:"ip,required"`
	JSON accountAccessGroupGetResponseResultIncludePajwohLqIPRuleJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqIPRuleJSON contains the JSON
// metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqIPRule]
type accountAccessGroupGetResponseResultIncludePajwohLqIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIncludePajwohLqIPRule) implementsAccountAccessGroupGetResponseResultInclude() {
}

type AccountAccessGroupGetResponseResultIncludePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                         `json:"ip,required"`
	JSON accountAccessGroupGetResponseResultIncludePajwohLqIPRuleIPJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqIPRuleIPJSON contains the JSON
// metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqIPRuleIP]
type accountAccessGroupGetResponseResultIncludePajwohLqIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccountAccessGroupGetResponseResultIncludePajwohLqIPListRule struct {
	IPList AccountAccessGroupGetResponseResultIncludePajwohLqIPListRuleIPList `json:"ip_list,required"`
	JSON   accountAccessGroupGetResponseResultIncludePajwohLqIPListRuleJSON   `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqIPListRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqIPListRule]
type accountAccessGroupGetResponseResultIncludePajwohLqIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIncludePajwohLqIPListRule) implementsAccountAccessGroupGetResponseResultInclude() {
}

type AccountAccessGroupGetResponseResultIncludePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                 `json:"id,required"`
	JSON accountAccessGroupGetResponseResultIncludePajwohLqIPListRuleIPListJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqIPListRuleIPListJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqIPListRuleIPList]
type accountAccessGroupGetResponseResultIncludePajwohLqIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccountAccessGroupGetResponseResultIncludePajwohLqCertificateRule struct {
	Certificate interface{}                                                           `json:"certificate,required"`
	JSON        accountAccessGroupGetResponseResultIncludePajwohLqCertificateRuleJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqCertificateRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqCertificateRule]
type accountAccessGroupGetResponseResultIncludePajwohLqCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIncludePajwohLqCertificateRule) implementsAccountAccessGroupGetResponseResultInclude() {
}

// Matches an Access group.
type AccountAccessGroupGetResponseResultIncludePajwohLqAccessGroupRule struct {
	Group AccountAccessGroupGetResponseResultIncludePajwohLqAccessGroupRuleGroup `json:"group,required"`
	JSON  accountAccessGroupGetResponseResultIncludePajwohLqAccessGroupRuleJSON  `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqAccessGroupRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqAccessGroupRule]
type accountAccessGroupGetResponseResultIncludePajwohLqAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIncludePajwohLqAccessGroupRule) implementsAccountAccessGroupGetResponseResultInclude() {
}

type AccountAccessGroupGetResponseResultIncludePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                     `json:"id,required"`
	JSON accountAccessGroupGetResponseResultIncludePajwohLqAccessGroupRuleGroupJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqAccessGroupRuleGroup]
type accountAccessGroupGetResponseResultIncludePajwohLqAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccountAccessGroupGetResponseResultIncludePajwohLqAzureGroupRule struct {
	AzureAd AccountAccessGroupGetResponseResultIncludePajwohLqAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accountAccessGroupGetResponseResultIncludePajwohLqAzureGroupRuleJSON    `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqAzureGroupRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqAzureGroupRule]
type accountAccessGroupGetResponseResultIncludePajwohLqAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIncludePajwohLqAzureGroupRule) implementsAccountAccessGroupGetResponseResultInclude() {
}

type AccountAccessGroupGetResponseResultIncludePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                      `json:"connection_id,required"`
	JSON         accountAccessGroupGetResponseResultIncludePajwohLqAzureGroupRuleAzureAdJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqAzureGroupRuleAzureAd]
type accountAccessGroupGetResponseResultIncludePajwohLqAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccountAccessGroupGetResponseResultIncludePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization AccountAccessGroupGetResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accountAccessGroupGetResponseResultIncludePajwohLqGitHubOrganizationRuleJSON               `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqGitHubOrganizationRule]
type accountAccessGroupGetResponseResultIncludePajwohLqGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIncludePajwohLqGitHubOrganizationRule) implementsAccountAccessGroupGetResponseResultInclude() {
}

type AccountAccessGroupGetResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                         `json:"name,required"`
	JSON accountAccessGroupGetResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization]
type accountAccessGroupGetResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccountAccessGroupGetResponseResultIncludePajwohLqGsuiteGroupRule struct {
	Gsuite AccountAccessGroupGetResponseResultIncludePajwohLqGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accountAccessGroupGetResponseResultIncludePajwohLqGsuiteGroupRuleJSON   `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqGsuiteGroupRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqGsuiteGroupRule]
type accountAccessGroupGetResponseResultIncludePajwohLqGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIncludePajwohLqGsuiteGroupRule) implementsAccountAccessGroupGetResponseResultInclude() {
}

type AccountAccessGroupGetResponseResultIncludePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                      `json:"email,required"`
	JSON  accountAccessGroupGetResponseResultIncludePajwohLqGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqGsuiteGroupRuleGsuite]
type accountAccessGroupGetResponseResultIncludePajwohLqGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccountAccessGroupGetResponseResultIncludePajwohLqOktaGroupRule struct {
	Okta AccountAccessGroupGetResponseResultIncludePajwohLqOktaGroupRuleOkta `json:"okta,required"`
	JSON accountAccessGroupGetResponseResultIncludePajwohLqOktaGroupRuleJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqOktaGroupRule]
type accountAccessGroupGetResponseResultIncludePajwohLqOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIncludePajwohLqOktaGroupRule) implementsAccountAccessGroupGetResponseResultInclude() {
}

type AccountAccessGroupGetResponseResultIncludePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                  `json:"email,required"`
	JSON  accountAccessGroupGetResponseResultIncludePajwohLqOktaGroupRuleOktaJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqOktaGroupRuleOkta]
type accountAccessGroupGetResponseResultIncludePajwohLqOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccountAccessGroupGetResponseResultIncludePajwohLqSamlGroupRule struct {
	Saml AccountAccessGroupGetResponseResultIncludePajwohLqSamlGroupRuleSaml `json:"saml,required"`
	JSON accountAccessGroupGetResponseResultIncludePajwohLqSamlGroupRuleJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqSamlGroupRule]
type accountAccessGroupGetResponseResultIncludePajwohLqSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIncludePajwohLqSamlGroupRule) implementsAccountAccessGroupGetResponseResultInclude() {
}

type AccountAccessGroupGetResponseResultIncludePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                  `json:"attribute_value,required"`
	JSON           accountAccessGroupGetResponseResultIncludePajwohLqSamlGroupRuleSamlJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqSamlGroupRuleSaml]
type accountAccessGroupGetResponseResultIncludePajwohLqSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccountAccessGroupGetResponseResultIncludePajwohLqServiceTokenRule struct {
	ServiceToken AccountAccessGroupGetResponseResultIncludePajwohLqServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accountAccessGroupGetResponseResultIncludePajwohLqServiceTokenRuleJSON         `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqServiceTokenRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqServiceTokenRule]
type accountAccessGroupGetResponseResultIncludePajwohLqServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIncludePajwohLqServiceTokenRule) implementsAccountAccessGroupGetResponseResultInclude() {
}

type AccountAccessGroupGetResponseResultIncludePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                             `json:"token_id,required"`
	JSON    accountAccessGroupGetResponseResultIncludePajwohLqServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqServiceTokenRuleServiceToken]
type accountAccessGroupGetResponseResultIncludePajwohLqServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccountAccessGroupGetResponseResultIncludePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                    `json:"any_valid_service_token,required"`
	JSON                 accountAccessGroupGetResponseResultIncludePajwohLqAnyValidServiceTokenRuleJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqAnyValidServiceTokenRule]
type accountAccessGroupGetResponseResultIncludePajwohLqAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIncludePajwohLqAnyValidServiceTokenRule) implementsAccountAccessGroupGetResponseResultInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccountAccessGroupGetResponseResultIncludePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation AccountAccessGroupGetResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accountAccessGroupGetResponseResultIncludePajwohLqExternalEvaluationRuleJSON               `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqExternalEvaluationRule]
type accountAccessGroupGetResponseResultIncludePajwohLqExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIncludePajwohLqExternalEvaluationRule) implementsAccountAccessGroupGetResponseResultInclude() {
}

type AccountAccessGroupGetResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                         `json:"keys_url,required"`
	JSON    accountAccessGroupGetResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation]
type accountAccessGroupGetResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccountAccessGroupGetResponseResultIncludePajwohLqCountryRule struct {
	Geo  AccountAccessGroupGetResponseResultIncludePajwohLqCountryRuleGeo  `json:"geo,required"`
	JSON accountAccessGroupGetResponseResultIncludePajwohLqCountryRuleJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqCountryRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqCountryRule]
type accountAccessGroupGetResponseResultIncludePajwohLqCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIncludePajwohLqCountryRule) implementsAccountAccessGroupGetResponseResultInclude() {
}

type AccountAccessGroupGetResponseResultIncludePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                               `json:"country_code,required"`
	JSON        accountAccessGroupGetResponseResultIncludePajwohLqCountryRuleGeoJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqCountryRuleGeoJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqCountryRuleGeo]
type accountAccessGroupGetResponseResultIncludePajwohLqCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccountAccessGroupGetResponseResultIncludePajwohLqAuthenticationMethodRule struct {
	AuthMethod AccountAccessGroupGetResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accountAccessGroupGetResponseResultIncludePajwohLqAuthenticationMethodRuleJSON       `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqAuthenticationMethodRule]
type accountAccessGroupGetResponseResultIncludePajwohLqAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIncludePajwohLqAuthenticationMethodRule) implementsAccountAccessGroupGetResponseResultInclude() {
}

type AccountAccessGroupGetResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                   `json:"auth_method,required"`
	JSON       accountAccessGroupGetResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod]
type accountAccessGroupGetResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccountAccessGroupGetResponseResultIncludePajwohLqDevicePostureRule struct {
	DevicePosture AccountAccessGroupGetResponseResultIncludePajwohLqDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accountAccessGroupGetResponseResultIncludePajwohLqDevicePostureRuleJSON          `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqDevicePostureRule]
type accountAccessGroupGetResponseResultIncludePajwohLqDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIncludePajwohLqDevicePostureRule) implementsAccountAccessGroupGetResponseResultInclude() {
}

type AccountAccessGroupGetResponseResultIncludePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                               `json:"integration_uid,required"`
	JSON           accountAccessGroupGetResponseResultIncludePajwohLqDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIncludePajwohLqDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIncludePajwohLqDevicePostureRuleDevicePosture]
type accountAccessGroupGetResponseResultIncludePajwohLqDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIncludePajwohLqDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqEmailRule],
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqEmailListRule],
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqDomainRule],
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqEveryoneRule],
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqIPRule],
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqIPListRule],
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqCertificateRule],
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqAccessGroupRule],
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqAzureGroupRule],
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqGitHubOrganizationRule],
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqGsuiteGroupRule],
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqOktaGroupRule],
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqSamlGroupRule],
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqServiceTokenRule],
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqAnyValidServiceTokenRule],
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqExternalEvaluationRule],
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqCountryRule],
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqAuthenticationMethodRule]
// or [AccountAccessGroupGetResponseResultIsDefaultPajwohLqDevicePostureRule].
type AccountAccessGroupGetResponseResultIsDefault interface {
	implementsAccountAccessGroupGetResponseResultIsDefault()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountAccessGroupGetResponseResultIsDefault)(nil)).Elem(), "")
}

// Matches a specific email.
type AccountAccessGroupGetResponseResultIsDefaultPajwohLqEmailRule struct {
	Email AccountAccessGroupGetResponseResultIsDefaultPajwohLqEmailRuleEmail `json:"email,required"`
	JSON  accountAccessGroupGetResponseResultIsDefaultPajwohLqEmailRuleJSON  `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqEmailRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqEmailRule]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIsDefaultPajwohLqEmailRule) implementsAccountAccessGroupGetResponseResultIsDefault() {
}

type AccountAccessGroupGetResponseResultIsDefaultPajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                 `json:"email,required" format:"email"`
	JSON  accountAccessGroupGetResponseResultIsDefaultPajwohLqEmailRuleEmailJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqEmailRuleEmailJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqEmailRuleEmail]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccountAccessGroupGetResponseResultIsDefaultPajwohLqEmailListRule struct {
	EmailList AccountAccessGroupGetResponseResultIsDefaultPajwohLqEmailListRuleEmailList `json:"email_list,required"`
	JSON      accountAccessGroupGetResponseResultIsDefaultPajwohLqEmailListRuleJSON      `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqEmailListRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqEmailListRule]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIsDefaultPajwohLqEmailListRule) implementsAccountAccessGroupGetResponseResultIsDefault() {
}

type AccountAccessGroupGetResponseResultIsDefaultPajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                         `json:"id,required"`
	JSON accountAccessGroupGetResponseResultIsDefaultPajwohLqEmailListRuleEmailListJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqEmailListRuleEmailList]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccountAccessGroupGetResponseResultIsDefaultPajwohLqDomainRule struct {
	EmailDomain AccountAccessGroupGetResponseResultIsDefaultPajwohLqDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accountAccessGroupGetResponseResultIsDefaultPajwohLqDomainRuleJSON        `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqDomainRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqDomainRule]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIsDefaultPajwohLqDomainRule) implementsAccountAccessGroupGetResponseResultIsDefault() {
}

type AccountAccessGroupGetResponseResultIsDefaultPajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                        `json:"domain,required"`
	JSON   accountAccessGroupGetResponseResultIsDefaultPajwohLqDomainRuleEmailDomainJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqDomainRuleEmailDomain]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccountAccessGroupGetResponseResultIsDefaultPajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                          `json:"everyone,required"`
	JSON     accountAccessGroupGetResponseResultIsDefaultPajwohLqEveryoneRuleJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqEveryoneRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqEveryoneRule]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIsDefaultPajwohLqEveryoneRule) implementsAccountAccessGroupGetResponseResultIsDefault() {
}

// Matches an IP address block.
type AccountAccessGroupGetResponseResultIsDefaultPajwohLqIPRule struct {
	IP   AccountAccessGroupGetResponseResultIsDefaultPajwohLqIPRuleIP   `json:"ip,required"`
	JSON accountAccessGroupGetResponseResultIsDefaultPajwohLqIPRuleJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqIPRuleJSON contains the JSON
// metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqIPRule]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIsDefaultPajwohLqIPRule) implementsAccountAccessGroupGetResponseResultIsDefault() {
}

type AccountAccessGroupGetResponseResultIsDefaultPajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                           `json:"ip,required"`
	JSON accountAccessGroupGetResponseResultIsDefaultPajwohLqIPRuleIPJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqIPRuleIPJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqIPRuleIP]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccountAccessGroupGetResponseResultIsDefaultPajwohLqIPListRule struct {
	IPList AccountAccessGroupGetResponseResultIsDefaultPajwohLqIPListRuleIPList `json:"ip_list,required"`
	JSON   accountAccessGroupGetResponseResultIsDefaultPajwohLqIPListRuleJSON   `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqIPListRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqIPListRule]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIsDefaultPajwohLqIPListRule) implementsAccountAccessGroupGetResponseResultIsDefault() {
}

type AccountAccessGroupGetResponseResultIsDefaultPajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                   `json:"id,required"`
	JSON accountAccessGroupGetResponseResultIsDefaultPajwohLqIPListRuleIPListJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqIPListRuleIPList]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccountAccessGroupGetResponseResultIsDefaultPajwohLqCertificateRule struct {
	Certificate interface{}                                                             `json:"certificate,required"`
	JSON        accountAccessGroupGetResponseResultIsDefaultPajwohLqCertificateRuleJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqCertificateRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqCertificateRule]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIsDefaultPajwohLqCertificateRule) implementsAccountAccessGroupGetResponseResultIsDefault() {
}

// Matches an Access group.
type AccountAccessGroupGetResponseResultIsDefaultPajwohLqAccessGroupRule struct {
	Group AccountAccessGroupGetResponseResultIsDefaultPajwohLqAccessGroupRuleGroup `json:"group,required"`
	JSON  accountAccessGroupGetResponseResultIsDefaultPajwohLqAccessGroupRuleJSON  `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqAccessGroupRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqAccessGroupRule]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIsDefaultPajwohLqAccessGroupRule) implementsAccountAccessGroupGetResponseResultIsDefault() {
}

type AccountAccessGroupGetResponseResultIsDefaultPajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                       `json:"id,required"`
	JSON accountAccessGroupGetResponseResultIsDefaultPajwohLqAccessGroupRuleGroupJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqAccessGroupRuleGroup]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccountAccessGroupGetResponseResultIsDefaultPajwohLqAzureGroupRule struct {
	AzureAd AccountAccessGroupGetResponseResultIsDefaultPajwohLqAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accountAccessGroupGetResponseResultIsDefaultPajwohLqAzureGroupRuleJSON    `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqAzureGroupRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqAzureGroupRule]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIsDefaultPajwohLqAzureGroupRule) implementsAccountAccessGroupGetResponseResultIsDefault() {
}

type AccountAccessGroupGetResponseResultIsDefaultPajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                        `json:"connection_id,required"`
	JSON         accountAccessGroupGetResponseResultIsDefaultPajwohLqAzureGroupRuleAzureAdJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqAzureGroupRuleAzureAd]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccountAccessGroupGetResponseResultIsDefaultPajwohLqGitHubOrganizationRule struct {
	GitHubOrganization AccountAccessGroupGetResponseResultIsDefaultPajwohLqGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accountAccessGroupGetResponseResultIsDefaultPajwohLqGitHubOrganizationRuleJSON               `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqGitHubOrganizationRule]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIsDefaultPajwohLqGitHubOrganizationRule) implementsAccountAccessGroupGetResponseResultIsDefault() {
}

type AccountAccessGroupGetResponseResultIsDefaultPajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                           `json:"name,required"`
	JSON accountAccessGroupGetResponseResultIsDefaultPajwohLqGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqGitHubOrganizationRuleGitHubOrganization]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccountAccessGroupGetResponseResultIsDefaultPajwohLqGsuiteGroupRule struct {
	Gsuite AccountAccessGroupGetResponseResultIsDefaultPajwohLqGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accountAccessGroupGetResponseResultIsDefaultPajwohLqGsuiteGroupRuleJSON   `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqGsuiteGroupRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqGsuiteGroupRule]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIsDefaultPajwohLqGsuiteGroupRule) implementsAccountAccessGroupGetResponseResultIsDefault() {
}

type AccountAccessGroupGetResponseResultIsDefaultPajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                        `json:"email,required"`
	JSON  accountAccessGroupGetResponseResultIsDefaultPajwohLqGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqGsuiteGroupRuleGsuite]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccountAccessGroupGetResponseResultIsDefaultPajwohLqOktaGroupRule struct {
	Okta AccountAccessGroupGetResponseResultIsDefaultPajwohLqOktaGroupRuleOkta `json:"okta,required"`
	JSON accountAccessGroupGetResponseResultIsDefaultPajwohLqOktaGroupRuleJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqOktaGroupRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqOktaGroupRule]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIsDefaultPajwohLqOktaGroupRule) implementsAccountAccessGroupGetResponseResultIsDefault() {
}

type AccountAccessGroupGetResponseResultIsDefaultPajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                    `json:"email,required"`
	JSON  accountAccessGroupGetResponseResultIsDefaultPajwohLqOktaGroupRuleOktaJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqOktaGroupRuleOkta]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccountAccessGroupGetResponseResultIsDefaultPajwohLqSamlGroupRule struct {
	Saml AccountAccessGroupGetResponseResultIsDefaultPajwohLqSamlGroupRuleSaml `json:"saml,required"`
	JSON accountAccessGroupGetResponseResultIsDefaultPajwohLqSamlGroupRuleJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqSamlGroupRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqSamlGroupRule]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIsDefaultPajwohLqSamlGroupRule) implementsAccountAccessGroupGetResponseResultIsDefault() {
}

type AccountAccessGroupGetResponseResultIsDefaultPajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                    `json:"attribute_value,required"`
	JSON           accountAccessGroupGetResponseResultIsDefaultPajwohLqSamlGroupRuleSamlJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqSamlGroupRuleSaml]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccountAccessGroupGetResponseResultIsDefaultPajwohLqServiceTokenRule struct {
	ServiceToken AccountAccessGroupGetResponseResultIsDefaultPajwohLqServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accountAccessGroupGetResponseResultIsDefaultPajwohLqServiceTokenRuleJSON         `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqServiceTokenRule]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIsDefaultPajwohLqServiceTokenRule) implementsAccountAccessGroupGetResponseResultIsDefault() {
}

type AccountAccessGroupGetResponseResultIsDefaultPajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                               `json:"token_id,required"`
	JSON    accountAccessGroupGetResponseResultIsDefaultPajwohLqServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqServiceTokenRuleServiceToken]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccountAccessGroupGetResponseResultIsDefaultPajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                      `json:"any_valid_service_token,required"`
	JSON                 accountAccessGroupGetResponseResultIsDefaultPajwohLqAnyValidServiceTokenRuleJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqAnyValidServiceTokenRule]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIsDefaultPajwohLqAnyValidServiceTokenRule) implementsAccountAccessGroupGetResponseResultIsDefault() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccountAccessGroupGetResponseResultIsDefaultPajwohLqExternalEvaluationRule struct {
	ExternalEvaluation AccountAccessGroupGetResponseResultIsDefaultPajwohLqExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accountAccessGroupGetResponseResultIsDefaultPajwohLqExternalEvaluationRuleJSON               `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqExternalEvaluationRule]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIsDefaultPajwohLqExternalEvaluationRule) implementsAccountAccessGroupGetResponseResultIsDefault() {
}

type AccountAccessGroupGetResponseResultIsDefaultPajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                           `json:"keys_url,required"`
	JSON    accountAccessGroupGetResponseResultIsDefaultPajwohLqExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqExternalEvaluationRuleExternalEvaluation]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccountAccessGroupGetResponseResultIsDefaultPajwohLqCountryRule struct {
	Geo  AccountAccessGroupGetResponseResultIsDefaultPajwohLqCountryRuleGeo  `json:"geo,required"`
	JSON accountAccessGroupGetResponseResultIsDefaultPajwohLqCountryRuleJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqCountryRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqCountryRule]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIsDefaultPajwohLqCountryRule) implementsAccountAccessGroupGetResponseResultIsDefault() {
}

type AccountAccessGroupGetResponseResultIsDefaultPajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                 `json:"country_code,required"`
	JSON        accountAccessGroupGetResponseResultIsDefaultPajwohLqCountryRuleGeoJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqCountryRuleGeoJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqCountryRuleGeo]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccountAccessGroupGetResponseResultIsDefaultPajwohLqAuthenticationMethodRule struct {
	AuthMethod AccountAccessGroupGetResponseResultIsDefaultPajwohLqAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accountAccessGroupGetResponseResultIsDefaultPajwohLqAuthenticationMethodRuleJSON       `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqAuthenticationMethodRule]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIsDefaultPajwohLqAuthenticationMethodRule) implementsAccountAccessGroupGetResponseResultIsDefault() {
}

type AccountAccessGroupGetResponseResultIsDefaultPajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                     `json:"auth_method,required"`
	JSON       accountAccessGroupGetResponseResultIsDefaultPajwohLqAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqAuthenticationMethodRuleAuthMethod]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccountAccessGroupGetResponseResultIsDefaultPajwohLqDevicePostureRule struct {
	DevicePosture AccountAccessGroupGetResponseResultIsDefaultPajwohLqDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accountAccessGroupGetResponseResultIsDefaultPajwohLqDevicePostureRuleJSON          `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqDevicePostureRule]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultIsDefaultPajwohLqDevicePostureRule) implementsAccountAccessGroupGetResponseResultIsDefault() {
}

type AccountAccessGroupGetResponseResultIsDefaultPajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                 `json:"integration_uid,required"`
	JSON           accountAccessGroupGetResponseResultIsDefaultPajwohLqDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accountAccessGroupGetResponseResultIsDefaultPajwohLqDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultIsDefaultPajwohLqDevicePostureRuleDevicePosture]
type accountAccessGroupGetResponseResultIsDefaultPajwohLqDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultIsDefaultPajwohLqDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccountAccessGroupGetResponseResultRequirePajwohLqEmailRule],
// [AccountAccessGroupGetResponseResultRequirePajwohLqEmailListRule],
// [AccountAccessGroupGetResponseResultRequirePajwohLqDomainRule],
// [AccountAccessGroupGetResponseResultRequirePajwohLqEveryoneRule],
// [AccountAccessGroupGetResponseResultRequirePajwohLqIPRule],
// [AccountAccessGroupGetResponseResultRequirePajwohLqIPListRule],
// [AccountAccessGroupGetResponseResultRequirePajwohLqCertificateRule],
// [AccountAccessGroupGetResponseResultRequirePajwohLqAccessGroupRule],
// [AccountAccessGroupGetResponseResultRequirePajwohLqAzureGroupRule],
// [AccountAccessGroupGetResponseResultRequirePajwohLqGitHubOrganizationRule],
// [AccountAccessGroupGetResponseResultRequirePajwohLqGsuiteGroupRule],
// [AccountAccessGroupGetResponseResultRequirePajwohLqOktaGroupRule],
// [AccountAccessGroupGetResponseResultRequirePajwohLqSamlGroupRule],
// [AccountAccessGroupGetResponseResultRequirePajwohLqServiceTokenRule],
// [AccountAccessGroupGetResponseResultRequirePajwohLqAnyValidServiceTokenRule],
// [AccountAccessGroupGetResponseResultRequirePajwohLqExternalEvaluationRule],
// [AccountAccessGroupGetResponseResultRequirePajwohLqCountryRule],
// [AccountAccessGroupGetResponseResultRequirePajwohLqAuthenticationMethodRule] or
// [AccountAccessGroupGetResponseResultRequirePajwohLqDevicePostureRule].
type AccountAccessGroupGetResponseResultRequire interface {
	implementsAccountAccessGroupGetResponseResultRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountAccessGroupGetResponseResultRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type AccountAccessGroupGetResponseResultRequirePajwohLqEmailRule struct {
	Email AccountAccessGroupGetResponseResultRequirePajwohLqEmailRuleEmail `json:"email,required"`
	JSON  accountAccessGroupGetResponseResultRequirePajwohLqEmailRuleJSON  `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqEmailRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqEmailRule]
type accountAccessGroupGetResponseResultRequirePajwohLqEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultRequirePajwohLqEmailRule) implementsAccountAccessGroupGetResponseResultRequire() {
}

type AccountAccessGroupGetResponseResultRequirePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email string                                                               `json:"email,required" format:"email"`
	JSON  accountAccessGroupGetResponseResultRequirePajwohLqEmailRuleEmailJSON `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqEmailRuleEmailJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqEmailRuleEmail]
type accountAccessGroupGetResponseResultRequirePajwohLqEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccountAccessGroupGetResponseResultRequirePajwohLqEmailListRule struct {
	EmailList AccountAccessGroupGetResponseResultRequirePajwohLqEmailListRuleEmailList `json:"email_list,required"`
	JSON      accountAccessGroupGetResponseResultRequirePajwohLqEmailListRuleJSON      `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqEmailListRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqEmailListRule]
type accountAccessGroupGetResponseResultRequirePajwohLqEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultRequirePajwohLqEmailListRule) implementsAccountAccessGroupGetResponseResultRequire() {
}

type AccountAccessGroupGetResponseResultRequirePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                       `json:"id,required"`
	JSON accountAccessGroupGetResponseResultRequirePajwohLqEmailListRuleEmailListJSON `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqEmailListRuleEmailList]
type accountAccessGroupGetResponseResultRequirePajwohLqEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccountAccessGroupGetResponseResultRequirePajwohLqDomainRule struct {
	EmailDomain AccountAccessGroupGetResponseResultRequirePajwohLqDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accountAccessGroupGetResponseResultRequirePajwohLqDomainRuleJSON        `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqDomainRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqDomainRule]
type accountAccessGroupGetResponseResultRequirePajwohLqDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultRequirePajwohLqDomainRule) implementsAccountAccessGroupGetResponseResultRequire() {
}

type AccountAccessGroupGetResponseResultRequirePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                      `json:"domain,required"`
	JSON   accountAccessGroupGetResponseResultRequirePajwohLqDomainRuleEmailDomainJSON `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqDomainRuleEmailDomain]
type accountAccessGroupGetResponseResultRequirePajwohLqDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccountAccessGroupGetResponseResultRequirePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                        `json:"everyone,required"`
	JSON     accountAccessGroupGetResponseResultRequirePajwohLqEveryoneRuleJSON `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqEveryoneRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqEveryoneRule]
type accountAccessGroupGetResponseResultRequirePajwohLqEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultRequirePajwohLqEveryoneRule) implementsAccountAccessGroupGetResponseResultRequire() {
}

// Matches an IP address block.
type AccountAccessGroupGetResponseResultRequirePajwohLqIPRule struct {
	IP   AccountAccessGroupGetResponseResultRequirePajwohLqIPRuleIP   `json:"ip,required"`
	JSON accountAccessGroupGetResponseResultRequirePajwohLqIPRuleJSON `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqIPRuleJSON contains the JSON
// metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqIPRule]
type accountAccessGroupGetResponseResultRequirePajwohLqIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultRequirePajwohLqIPRule) implementsAccountAccessGroupGetResponseResultRequire() {
}

type AccountAccessGroupGetResponseResultRequirePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                         `json:"ip,required"`
	JSON accountAccessGroupGetResponseResultRequirePajwohLqIPRuleIPJSON `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqIPRuleIPJSON contains the JSON
// metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqIPRuleIP]
type accountAccessGroupGetResponseResultRequirePajwohLqIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccountAccessGroupGetResponseResultRequirePajwohLqIPListRule struct {
	IPList AccountAccessGroupGetResponseResultRequirePajwohLqIPListRuleIPList `json:"ip_list,required"`
	JSON   accountAccessGroupGetResponseResultRequirePajwohLqIPListRuleJSON   `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqIPListRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqIPListRule]
type accountAccessGroupGetResponseResultRequirePajwohLqIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultRequirePajwohLqIPListRule) implementsAccountAccessGroupGetResponseResultRequire() {
}

type AccountAccessGroupGetResponseResultRequirePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                 `json:"id,required"`
	JSON accountAccessGroupGetResponseResultRequirePajwohLqIPListRuleIPListJSON `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqIPListRuleIPListJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqIPListRuleIPList]
type accountAccessGroupGetResponseResultRequirePajwohLqIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccountAccessGroupGetResponseResultRequirePajwohLqCertificateRule struct {
	Certificate interface{}                                                           `json:"certificate,required"`
	JSON        accountAccessGroupGetResponseResultRequirePajwohLqCertificateRuleJSON `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqCertificateRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqCertificateRule]
type accountAccessGroupGetResponseResultRequirePajwohLqCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultRequirePajwohLqCertificateRule) implementsAccountAccessGroupGetResponseResultRequire() {
}

// Matches an Access group.
type AccountAccessGroupGetResponseResultRequirePajwohLqAccessGroupRule struct {
	Group AccountAccessGroupGetResponseResultRequirePajwohLqAccessGroupRuleGroup `json:"group,required"`
	JSON  accountAccessGroupGetResponseResultRequirePajwohLqAccessGroupRuleJSON  `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqAccessGroupRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqAccessGroupRule]
type accountAccessGroupGetResponseResultRequirePajwohLqAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultRequirePajwohLqAccessGroupRule) implementsAccountAccessGroupGetResponseResultRequire() {
}

type AccountAccessGroupGetResponseResultRequirePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                     `json:"id,required"`
	JSON accountAccessGroupGetResponseResultRequirePajwohLqAccessGroupRuleGroupJSON `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqAccessGroupRuleGroup]
type accountAccessGroupGetResponseResultRequirePajwohLqAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccountAccessGroupGetResponseResultRequirePajwohLqAzureGroupRule struct {
	AzureAd AccountAccessGroupGetResponseResultRequirePajwohLqAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accountAccessGroupGetResponseResultRequirePajwohLqAzureGroupRuleJSON    `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqAzureGroupRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqAzureGroupRule]
type accountAccessGroupGetResponseResultRequirePajwohLqAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultRequirePajwohLqAzureGroupRule) implementsAccountAccessGroupGetResponseResultRequire() {
}

type AccountAccessGroupGetResponseResultRequirePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                      `json:"connection_id,required"`
	JSON         accountAccessGroupGetResponseResultRequirePajwohLqAzureGroupRuleAzureAdJSON `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqAzureGroupRuleAzureAd]
type accountAccessGroupGetResponseResultRequirePajwohLqAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccountAccessGroupGetResponseResultRequirePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization AccountAccessGroupGetResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accountAccessGroupGetResponseResultRequirePajwohLqGitHubOrganizationRuleJSON               `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqGitHubOrganizationRule]
type accountAccessGroupGetResponseResultRequirePajwohLqGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultRequirePajwohLqGitHubOrganizationRule) implementsAccountAccessGroupGetResponseResultRequire() {
}

type AccountAccessGroupGetResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                         `json:"name,required"`
	JSON accountAccessGroupGetResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization]
type accountAccessGroupGetResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccountAccessGroupGetResponseResultRequirePajwohLqGsuiteGroupRule struct {
	Gsuite AccountAccessGroupGetResponseResultRequirePajwohLqGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accountAccessGroupGetResponseResultRequirePajwohLqGsuiteGroupRuleJSON   `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqGsuiteGroupRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqGsuiteGroupRule]
type accountAccessGroupGetResponseResultRequirePajwohLqGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultRequirePajwohLqGsuiteGroupRule) implementsAccountAccessGroupGetResponseResultRequire() {
}

type AccountAccessGroupGetResponseResultRequirePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                      `json:"email,required"`
	JSON  accountAccessGroupGetResponseResultRequirePajwohLqGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqGsuiteGroupRuleGsuite]
type accountAccessGroupGetResponseResultRequirePajwohLqGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccountAccessGroupGetResponseResultRequirePajwohLqOktaGroupRule struct {
	Okta AccountAccessGroupGetResponseResultRequirePajwohLqOktaGroupRuleOkta `json:"okta,required"`
	JSON accountAccessGroupGetResponseResultRequirePajwohLqOktaGroupRuleJSON `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqOktaGroupRule]
type accountAccessGroupGetResponseResultRequirePajwohLqOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultRequirePajwohLqOktaGroupRule) implementsAccountAccessGroupGetResponseResultRequire() {
}

type AccountAccessGroupGetResponseResultRequirePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                  `json:"email,required"`
	JSON  accountAccessGroupGetResponseResultRequirePajwohLqOktaGroupRuleOktaJSON `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqOktaGroupRuleOkta]
type accountAccessGroupGetResponseResultRequirePajwohLqOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccountAccessGroupGetResponseResultRequirePajwohLqSamlGroupRule struct {
	Saml AccountAccessGroupGetResponseResultRequirePajwohLqSamlGroupRuleSaml `json:"saml,required"`
	JSON accountAccessGroupGetResponseResultRequirePajwohLqSamlGroupRuleJSON `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqSamlGroupRule]
type accountAccessGroupGetResponseResultRequirePajwohLqSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultRequirePajwohLqSamlGroupRule) implementsAccountAccessGroupGetResponseResultRequire() {
}

type AccountAccessGroupGetResponseResultRequirePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                  `json:"attribute_value,required"`
	JSON           accountAccessGroupGetResponseResultRequirePajwohLqSamlGroupRuleSamlJSON `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqSamlGroupRuleSaml]
type accountAccessGroupGetResponseResultRequirePajwohLqSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccountAccessGroupGetResponseResultRequirePajwohLqServiceTokenRule struct {
	ServiceToken AccountAccessGroupGetResponseResultRequirePajwohLqServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accountAccessGroupGetResponseResultRequirePajwohLqServiceTokenRuleJSON         `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqServiceTokenRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqServiceTokenRule]
type accountAccessGroupGetResponseResultRequirePajwohLqServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultRequirePajwohLqServiceTokenRule) implementsAccountAccessGroupGetResponseResultRequire() {
}

type AccountAccessGroupGetResponseResultRequirePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                             `json:"token_id,required"`
	JSON    accountAccessGroupGetResponseResultRequirePajwohLqServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqServiceTokenRuleServiceToken]
type accountAccessGroupGetResponseResultRequirePajwohLqServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccountAccessGroupGetResponseResultRequirePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                    `json:"any_valid_service_token,required"`
	JSON                 accountAccessGroupGetResponseResultRequirePajwohLqAnyValidServiceTokenRuleJSON `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqAnyValidServiceTokenRule]
type accountAccessGroupGetResponseResultRequirePajwohLqAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultRequirePajwohLqAnyValidServiceTokenRule) implementsAccountAccessGroupGetResponseResultRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccountAccessGroupGetResponseResultRequirePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation AccountAccessGroupGetResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accountAccessGroupGetResponseResultRequirePajwohLqExternalEvaluationRuleJSON               `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqExternalEvaluationRule]
type accountAccessGroupGetResponseResultRequirePajwohLqExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultRequirePajwohLqExternalEvaluationRule) implementsAccountAccessGroupGetResponseResultRequire() {
}

type AccountAccessGroupGetResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                         `json:"keys_url,required"`
	JSON    accountAccessGroupGetResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation]
type accountAccessGroupGetResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccountAccessGroupGetResponseResultRequirePajwohLqCountryRule struct {
	Geo  AccountAccessGroupGetResponseResultRequirePajwohLqCountryRuleGeo  `json:"geo,required"`
	JSON accountAccessGroupGetResponseResultRequirePajwohLqCountryRuleJSON `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqCountryRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqCountryRule]
type accountAccessGroupGetResponseResultRequirePajwohLqCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultRequirePajwohLqCountryRule) implementsAccountAccessGroupGetResponseResultRequire() {
}

type AccountAccessGroupGetResponseResultRequirePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                               `json:"country_code,required"`
	JSON        accountAccessGroupGetResponseResultRequirePajwohLqCountryRuleGeoJSON `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqCountryRuleGeoJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqCountryRuleGeo]
type accountAccessGroupGetResponseResultRequirePajwohLqCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccountAccessGroupGetResponseResultRequirePajwohLqAuthenticationMethodRule struct {
	AuthMethod AccountAccessGroupGetResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accountAccessGroupGetResponseResultRequirePajwohLqAuthenticationMethodRuleJSON       `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqAuthenticationMethodRule]
type accountAccessGroupGetResponseResultRequirePajwohLqAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultRequirePajwohLqAuthenticationMethodRule) implementsAccountAccessGroupGetResponseResultRequire() {
}

type AccountAccessGroupGetResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                   `json:"auth_method,required"`
	JSON       accountAccessGroupGetResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod]
type accountAccessGroupGetResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccountAccessGroupGetResponseResultRequirePajwohLqDevicePostureRule struct {
	DevicePosture AccountAccessGroupGetResponseResultRequirePajwohLqDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accountAccessGroupGetResponseResultRequirePajwohLqDevicePostureRuleJSON          `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqDevicePostureRule]
type accountAccessGroupGetResponseResultRequirePajwohLqDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupGetResponseResultRequirePajwohLqDevicePostureRule) implementsAccountAccessGroupGetResponseResultRequire() {
}

type AccountAccessGroupGetResponseResultRequirePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                               `json:"integration_uid,required"`
	JSON           accountAccessGroupGetResponseResultRequirePajwohLqDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accountAccessGroupGetResponseResultRequirePajwohLqDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupGetResponseResultRequirePajwohLqDevicePostureRuleDevicePosture]
type accountAccessGroupGetResponseResultRequirePajwohLqDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupGetResponseResultRequirePajwohLqDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessGroupGetResponseSuccess bool

const (
	AccountAccessGroupGetResponseSuccessTrue AccountAccessGroupGetResponseSuccess = true
)

type AccountAccessGroupUpdateResponse struct {
	Errors   []AccountAccessGroupUpdateResponseError   `json:"errors"`
	Messages []AccountAccessGroupUpdateResponseMessage `json:"messages"`
	Result   AccountAccessGroupUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessGroupUpdateResponseSuccess `json:"success"`
	JSON    accountAccessGroupUpdateResponseJSON    `json:"-"`
}

// accountAccessGroupUpdateResponseJSON contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponse]
type accountAccessGroupUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessGroupUpdateResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountAccessGroupUpdateResponseErrorJSON `json:"-"`
}

// accountAccessGroupUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccountAccessGroupUpdateResponseError]
type accountAccessGroupUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessGroupUpdateResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountAccessGroupUpdateResponseMessageJSON `json:"-"`
}

// accountAccessGroupUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccountAccessGroupUpdateResponseMessage]
type accountAccessGroupUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessGroupUpdateResponseResult struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccountAccessGroupUpdateResponseResultExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccountAccessGroupUpdateResponseResultInclude `json:"include"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	IsDefault []AccountAccessGroupUpdateResponseResultIsDefault `json:"is_default"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []AccountAccessGroupUpdateResponseResultRequire `json:"require"`
	UpdatedAt time.Time                                       `json:"updated_at" format:"date-time"`
	JSON      accountAccessGroupUpdateResponseResultJSON      `json:"-"`
}

// accountAccessGroupUpdateResponseResultJSON contains the JSON metadata for the
// struct [AccountAccessGroupUpdateResponseResult]
type accountAccessGroupUpdateResponseResultJSON struct {
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

func (r *AccountAccessGroupUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqEmailRule],
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqEmailListRule],
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqDomainRule],
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqEveryoneRule],
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqIPRule],
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqIPListRule],
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqCertificateRule],
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqAccessGroupRule],
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqAzureGroupRule],
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqGitHubOrganizationRule],
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqGsuiteGroupRule],
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqOktaGroupRule],
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqSamlGroupRule],
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqServiceTokenRule],
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqAnyValidServiceTokenRule],
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqExternalEvaluationRule],
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqCountryRule],
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqAuthenticationMethodRule]
// or [AccountAccessGroupUpdateResponseResultExcludePajwohLqDevicePostureRule].
type AccountAccessGroupUpdateResponseResultExclude interface {
	implementsAccountAccessGroupUpdateResponseResultExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountAccessGroupUpdateResponseResultExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccountAccessGroupUpdateResponseResultExcludePajwohLqEmailRule struct {
	Email AccountAccessGroupUpdateResponseResultExcludePajwohLqEmailRuleEmail `json:"email,required"`
	JSON  accountAccessGroupUpdateResponseResultExcludePajwohLqEmailRuleJSON  `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqEmailRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqEmailRule]
type accountAccessGroupUpdateResponseResultExcludePajwohLqEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultExcludePajwohLqEmailRule) implementsAccountAccessGroupUpdateResponseResultExclude() {
}

type AccountAccessGroupUpdateResponseResultExcludePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                  `json:"email,required" format:"email"`
	JSON  accountAccessGroupUpdateResponseResultExcludePajwohLqEmailRuleEmailJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqEmailRuleEmailJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqEmailRuleEmail]
type accountAccessGroupUpdateResponseResultExcludePajwohLqEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccountAccessGroupUpdateResponseResultExcludePajwohLqEmailListRule struct {
	EmailList AccountAccessGroupUpdateResponseResultExcludePajwohLqEmailListRuleEmailList `json:"email_list,required"`
	JSON      accountAccessGroupUpdateResponseResultExcludePajwohLqEmailListRuleJSON      `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqEmailListRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqEmailListRule]
type accountAccessGroupUpdateResponseResultExcludePajwohLqEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultExcludePajwohLqEmailListRule) implementsAccountAccessGroupUpdateResponseResultExclude() {
}

type AccountAccessGroupUpdateResponseResultExcludePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                          `json:"id,required"`
	JSON accountAccessGroupUpdateResponseResultExcludePajwohLqEmailListRuleEmailListJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqEmailListRuleEmailList]
type accountAccessGroupUpdateResponseResultExcludePajwohLqEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccountAccessGroupUpdateResponseResultExcludePajwohLqDomainRule struct {
	EmailDomain AccountAccessGroupUpdateResponseResultExcludePajwohLqDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accountAccessGroupUpdateResponseResultExcludePajwohLqDomainRuleJSON        `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqDomainRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqDomainRule]
type accountAccessGroupUpdateResponseResultExcludePajwohLqDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultExcludePajwohLqDomainRule) implementsAccountAccessGroupUpdateResponseResultExclude() {
}

type AccountAccessGroupUpdateResponseResultExcludePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                         `json:"domain,required"`
	JSON   accountAccessGroupUpdateResponseResultExcludePajwohLqDomainRuleEmailDomainJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqDomainRuleEmailDomain]
type accountAccessGroupUpdateResponseResultExcludePajwohLqDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccountAccessGroupUpdateResponseResultExcludePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                           `json:"everyone,required"`
	JSON     accountAccessGroupUpdateResponseResultExcludePajwohLqEveryoneRuleJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqEveryoneRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqEveryoneRule]
type accountAccessGroupUpdateResponseResultExcludePajwohLqEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultExcludePajwohLqEveryoneRule) implementsAccountAccessGroupUpdateResponseResultExclude() {
}

// Matches an IP address block.
type AccountAccessGroupUpdateResponseResultExcludePajwohLqIPRule struct {
	IP   AccountAccessGroupUpdateResponseResultExcludePajwohLqIPRuleIP   `json:"ip,required"`
	JSON accountAccessGroupUpdateResponseResultExcludePajwohLqIPRuleJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqIPRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqIPRule]
type accountAccessGroupUpdateResponseResultExcludePajwohLqIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultExcludePajwohLqIPRule) implementsAccountAccessGroupUpdateResponseResultExclude() {
}

type AccountAccessGroupUpdateResponseResultExcludePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                            `json:"ip,required"`
	JSON accountAccessGroupUpdateResponseResultExcludePajwohLqIPRuleIPJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqIPRuleIPJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqIPRuleIP]
type accountAccessGroupUpdateResponseResultExcludePajwohLqIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccountAccessGroupUpdateResponseResultExcludePajwohLqIPListRule struct {
	IPList AccountAccessGroupUpdateResponseResultExcludePajwohLqIPListRuleIPList `json:"ip_list,required"`
	JSON   accountAccessGroupUpdateResponseResultExcludePajwohLqIPListRuleJSON   `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqIPListRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqIPListRule]
type accountAccessGroupUpdateResponseResultExcludePajwohLqIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultExcludePajwohLqIPListRule) implementsAccountAccessGroupUpdateResponseResultExclude() {
}

type AccountAccessGroupUpdateResponseResultExcludePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                    `json:"id,required"`
	JSON accountAccessGroupUpdateResponseResultExcludePajwohLqIPListRuleIPListJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqIPListRuleIPList]
type accountAccessGroupUpdateResponseResultExcludePajwohLqIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccountAccessGroupUpdateResponseResultExcludePajwohLqCertificateRule struct {
	Certificate interface{}                                                              `json:"certificate,required"`
	JSON        accountAccessGroupUpdateResponseResultExcludePajwohLqCertificateRuleJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqCertificateRule]
type accountAccessGroupUpdateResponseResultExcludePajwohLqCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultExcludePajwohLqCertificateRule) implementsAccountAccessGroupUpdateResponseResultExclude() {
}

// Matches an Access group.
type AccountAccessGroupUpdateResponseResultExcludePajwohLqAccessGroupRule struct {
	Group AccountAccessGroupUpdateResponseResultExcludePajwohLqAccessGroupRuleGroup `json:"group,required"`
	JSON  accountAccessGroupUpdateResponseResultExcludePajwohLqAccessGroupRuleJSON  `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqAccessGroupRule]
type accountAccessGroupUpdateResponseResultExcludePajwohLqAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultExcludePajwohLqAccessGroupRule) implementsAccountAccessGroupUpdateResponseResultExclude() {
}

type AccountAccessGroupUpdateResponseResultExcludePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                        `json:"id,required"`
	JSON accountAccessGroupUpdateResponseResultExcludePajwohLqAccessGroupRuleGroupJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqAccessGroupRuleGroup]
type accountAccessGroupUpdateResponseResultExcludePajwohLqAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccountAccessGroupUpdateResponseResultExcludePajwohLqAzureGroupRule struct {
	AzureAd AccountAccessGroupUpdateResponseResultExcludePajwohLqAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accountAccessGroupUpdateResponseResultExcludePajwohLqAzureGroupRuleJSON    `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqAzureGroupRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqAzureGroupRule]
type accountAccessGroupUpdateResponseResultExcludePajwohLqAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultExcludePajwohLqAzureGroupRule) implementsAccountAccessGroupUpdateResponseResultExclude() {
}

type AccountAccessGroupUpdateResponseResultExcludePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                         `json:"connection_id,required"`
	JSON         accountAccessGroupUpdateResponseResultExcludePajwohLqAzureGroupRuleAzureAdJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqAzureGroupRuleAzureAd]
type accountAccessGroupUpdateResponseResultExcludePajwohLqAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccountAccessGroupUpdateResponseResultExcludePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization AccountAccessGroupUpdateResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accountAccessGroupUpdateResponseResultExcludePajwohLqGitHubOrganizationRuleJSON               `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqGitHubOrganizationRule]
type accountAccessGroupUpdateResponseResultExcludePajwohLqGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultExcludePajwohLqGitHubOrganizationRule) implementsAccountAccessGroupUpdateResponseResultExclude() {
}

type AccountAccessGroupUpdateResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                            `json:"name,required"`
	JSON accountAccessGroupUpdateResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization]
type accountAccessGroupUpdateResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccountAccessGroupUpdateResponseResultExcludePajwohLqGsuiteGroupRule struct {
	Gsuite AccountAccessGroupUpdateResponseResultExcludePajwohLqGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accountAccessGroupUpdateResponseResultExcludePajwohLqGsuiteGroupRuleJSON   `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqGsuiteGroupRule]
type accountAccessGroupUpdateResponseResultExcludePajwohLqGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultExcludePajwohLqGsuiteGroupRule) implementsAccountAccessGroupUpdateResponseResultExclude() {
}

type AccountAccessGroupUpdateResponseResultExcludePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                         `json:"email,required"`
	JSON  accountAccessGroupUpdateResponseResultExcludePajwohLqGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqGsuiteGroupRuleGsuite]
type accountAccessGroupUpdateResponseResultExcludePajwohLqGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccountAccessGroupUpdateResponseResultExcludePajwohLqOktaGroupRule struct {
	Okta AccountAccessGroupUpdateResponseResultExcludePajwohLqOktaGroupRuleOkta `json:"okta,required"`
	JSON accountAccessGroupUpdateResponseResultExcludePajwohLqOktaGroupRuleJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqOktaGroupRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqOktaGroupRule]
type accountAccessGroupUpdateResponseResultExcludePajwohLqOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultExcludePajwohLqOktaGroupRule) implementsAccountAccessGroupUpdateResponseResultExclude() {
}

type AccountAccessGroupUpdateResponseResultExcludePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                     `json:"email,required"`
	JSON  accountAccessGroupUpdateResponseResultExcludePajwohLqOktaGroupRuleOktaJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqOktaGroupRuleOkta]
type accountAccessGroupUpdateResponseResultExcludePajwohLqOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccountAccessGroupUpdateResponseResultExcludePajwohLqSamlGroupRule struct {
	Saml AccountAccessGroupUpdateResponseResultExcludePajwohLqSamlGroupRuleSaml `json:"saml,required"`
	JSON accountAccessGroupUpdateResponseResultExcludePajwohLqSamlGroupRuleJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqSamlGroupRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqSamlGroupRule]
type accountAccessGroupUpdateResponseResultExcludePajwohLqSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultExcludePajwohLqSamlGroupRule) implementsAccountAccessGroupUpdateResponseResultExclude() {
}

type AccountAccessGroupUpdateResponseResultExcludePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                     `json:"attribute_value,required"`
	JSON           accountAccessGroupUpdateResponseResultExcludePajwohLqSamlGroupRuleSamlJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqSamlGroupRuleSaml]
type accountAccessGroupUpdateResponseResultExcludePajwohLqSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccountAccessGroupUpdateResponseResultExcludePajwohLqServiceTokenRule struct {
	ServiceToken AccountAccessGroupUpdateResponseResultExcludePajwohLqServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accountAccessGroupUpdateResponseResultExcludePajwohLqServiceTokenRuleJSON         `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqServiceTokenRule]
type accountAccessGroupUpdateResponseResultExcludePajwohLqServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultExcludePajwohLqServiceTokenRule) implementsAccountAccessGroupUpdateResponseResultExclude() {
}

type AccountAccessGroupUpdateResponseResultExcludePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                `json:"token_id,required"`
	JSON    accountAccessGroupUpdateResponseResultExcludePajwohLqServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqServiceTokenRuleServiceToken]
type accountAccessGroupUpdateResponseResultExcludePajwohLqServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccountAccessGroupUpdateResponseResultExcludePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                       `json:"any_valid_service_token,required"`
	JSON                 accountAccessGroupUpdateResponseResultExcludePajwohLqAnyValidServiceTokenRuleJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqAnyValidServiceTokenRule]
type accountAccessGroupUpdateResponseResultExcludePajwohLqAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultExcludePajwohLqAnyValidServiceTokenRule) implementsAccountAccessGroupUpdateResponseResultExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccountAccessGroupUpdateResponseResultExcludePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation AccountAccessGroupUpdateResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accountAccessGroupUpdateResponseResultExcludePajwohLqExternalEvaluationRuleJSON               `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqExternalEvaluationRule]
type accountAccessGroupUpdateResponseResultExcludePajwohLqExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultExcludePajwohLqExternalEvaluationRule) implementsAccountAccessGroupUpdateResponseResultExclude() {
}

type AccountAccessGroupUpdateResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                            `json:"keys_url,required"`
	JSON    accountAccessGroupUpdateResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation]
type accountAccessGroupUpdateResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccountAccessGroupUpdateResponseResultExcludePajwohLqCountryRule struct {
	Geo  AccountAccessGroupUpdateResponseResultExcludePajwohLqCountryRuleGeo  `json:"geo,required"`
	JSON accountAccessGroupUpdateResponseResultExcludePajwohLqCountryRuleJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqCountryRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqCountryRule]
type accountAccessGroupUpdateResponseResultExcludePajwohLqCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultExcludePajwohLqCountryRule) implementsAccountAccessGroupUpdateResponseResultExclude() {
}

type AccountAccessGroupUpdateResponseResultExcludePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                  `json:"country_code,required"`
	JSON        accountAccessGroupUpdateResponseResultExcludePajwohLqCountryRuleGeoJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqCountryRuleGeoJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqCountryRuleGeo]
type accountAccessGroupUpdateResponseResultExcludePajwohLqCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccountAccessGroupUpdateResponseResultExcludePajwohLqAuthenticationMethodRule struct {
	AuthMethod AccountAccessGroupUpdateResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accountAccessGroupUpdateResponseResultExcludePajwohLqAuthenticationMethodRuleJSON       `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqAuthenticationMethodRule]
type accountAccessGroupUpdateResponseResultExcludePajwohLqAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultExcludePajwohLqAuthenticationMethodRule) implementsAccountAccessGroupUpdateResponseResultExclude() {
}

type AccountAccessGroupUpdateResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                      `json:"auth_method,required"`
	JSON       accountAccessGroupUpdateResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod]
type accountAccessGroupUpdateResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccountAccessGroupUpdateResponseResultExcludePajwohLqDevicePostureRule struct {
	DevicePosture AccountAccessGroupUpdateResponseResultExcludePajwohLqDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accountAccessGroupUpdateResponseResultExcludePajwohLqDevicePostureRuleJSON          `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqDevicePostureRule]
type accountAccessGroupUpdateResponseResultExcludePajwohLqDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultExcludePajwohLqDevicePostureRule) implementsAccountAccessGroupUpdateResponseResultExclude() {
}

type AccountAccessGroupUpdateResponseResultExcludePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                  `json:"integration_uid,required"`
	JSON           accountAccessGroupUpdateResponseResultExcludePajwohLqDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultExcludePajwohLqDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultExcludePajwohLqDevicePostureRuleDevicePosture]
type accountAccessGroupUpdateResponseResultExcludePajwohLqDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultExcludePajwohLqDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqEmailRule],
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqEmailListRule],
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqDomainRule],
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqEveryoneRule],
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqIPRule],
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqIPListRule],
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqCertificateRule],
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqAccessGroupRule],
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqAzureGroupRule],
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqGitHubOrganizationRule],
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqGsuiteGroupRule],
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqOktaGroupRule],
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqSamlGroupRule],
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqServiceTokenRule],
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqAnyValidServiceTokenRule],
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqExternalEvaluationRule],
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqCountryRule],
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqAuthenticationMethodRule]
// or [AccountAccessGroupUpdateResponseResultIncludePajwohLqDevicePostureRule].
type AccountAccessGroupUpdateResponseResultInclude interface {
	implementsAccountAccessGroupUpdateResponseResultInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountAccessGroupUpdateResponseResultInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccountAccessGroupUpdateResponseResultIncludePajwohLqEmailRule struct {
	Email AccountAccessGroupUpdateResponseResultIncludePajwohLqEmailRuleEmail `json:"email,required"`
	JSON  accountAccessGroupUpdateResponseResultIncludePajwohLqEmailRuleJSON  `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqEmailRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqEmailRule]
type accountAccessGroupUpdateResponseResultIncludePajwohLqEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIncludePajwohLqEmailRule) implementsAccountAccessGroupUpdateResponseResultInclude() {
}

type AccountAccessGroupUpdateResponseResultIncludePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                  `json:"email,required" format:"email"`
	JSON  accountAccessGroupUpdateResponseResultIncludePajwohLqEmailRuleEmailJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqEmailRuleEmailJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqEmailRuleEmail]
type accountAccessGroupUpdateResponseResultIncludePajwohLqEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccountAccessGroupUpdateResponseResultIncludePajwohLqEmailListRule struct {
	EmailList AccountAccessGroupUpdateResponseResultIncludePajwohLqEmailListRuleEmailList `json:"email_list,required"`
	JSON      accountAccessGroupUpdateResponseResultIncludePajwohLqEmailListRuleJSON      `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqEmailListRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqEmailListRule]
type accountAccessGroupUpdateResponseResultIncludePajwohLqEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIncludePajwohLqEmailListRule) implementsAccountAccessGroupUpdateResponseResultInclude() {
}

type AccountAccessGroupUpdateResponseResultIncludePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                          `json:"id,required"`
	JSON accountAccessGroupUpdateResponseResultIncludePajwohLqEmailListRuleEmailListJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqEmailListRuleEmailList]
type accountAccessGroupUpdateResponseResultIncludePajwohLqEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccountAccessGroupUpdateResponseResultIncludePajwohLqDomainRule struct {
	EmailDomain AccountAccessGroupUpdateResponseResultIncludePajwohLqDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accountAccessGroupUpdateResponseResultIncludePajwohLqDomainRuleJSON        `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqDomainRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqDomainRule]
type accountAccessGroupUpdateResponseResultIncludePajwohLqDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIncludePajwohLqDomainRule) implementsAccountAccessGroupUpdateResponseResultInclude() {
}

type AccountAccessGroupUpdateResponseResultIncludePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                         `json:"domain,required"`
	JSON   accountAccessGroupUpdateResponseResultIncludePajwohLqDomainRuleEmailDomainJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqDomainRuleEmailDomain]
type accountAccessGroupUpdateResponseResultIncludePajwohLqDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccountAccessGroupUpdateResponseResultIncludePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                           `json:"everyone,required"`
	JSON     accountAccessGroupUpdateResponseResultIncludePajwohLqEveryoneRuleJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqEveryoneRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqEveryoneRule]
type accountAccessGroupUpdateResponseResultIncludePajwohLqEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIncludePajwohLqEveryoneRule) implementsAccountAccessGroupUpdateResponseResultInclude() {
}

// Matches an IP address block.
type AccountAccessGroupUpdateResponseResultIncludePajwohLqIPRule struct {
	IP   AccountAccessGroupUpdateResponseResultIncludePajwohLqIPRuleIP   `json:"ip,required"`
	JSON accountAccessGroupUpdateResponseResultIncludePajwohLqIPRuleJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqIPRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqIPRule]
type accountAccessGroupUpdateResponseResultIncludePajwohLqIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIncludePajwohLqIPRule) implementsAccountAccessGroupUpdateResponseResultInclude() {
}

type AccountAccessGroupUpdateResponseResultIncludePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                            `json:"ip,required"`
	JSON accountAccessGroupUpdateResponseResultIncludePajwohLqIPRuleIPJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqIPRuleIPJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqIPRuleIP]
type accountAccessGroupUpdateResponseResultIncludePajwohLqIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccountAccessGroupUpdateResponseResultIncludePajwohLqIPListRule struct {
	IPList AccountAccessGroupUpdateResponseResultIncludePajwohLqIPListRuleIPList `json:"ip_list,required"`
	JSON   accountAccessGroupUpdateResponseResultIncludePajwohLqIPListRuleJSON   `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqIPListRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqIPListRule]
type accountAccessGroupUpdateResponseResultIncludePajwohLqIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIncludePajwohLqIPListRule) implementsAccountAccessGroupUpdateResponseResultInclude() {
}

type AccountAccessGroupUpdateResponseResultIncludePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                    `json:"id,required"`
	JSON accountAccessGroupUpdateResponseResultIncludePajwohLqIPListRuleIPListJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqIPListRuleIPList]
type accountAccessGroupUpdateResponseResultIncludePajwohLqIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccountAccessGroupUpdateResponseResultIncludePajwohLqCertificateRule struct {
	Certificate interface{}                                                              `json:"certificate,required"`
	JSON        accountAccessGroupUpdateResponseResultIncludePajwohLqCertificateRuleJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqCertificateRule]
type accountAccessGroupUpdateResponseResultIncludePajwohLqCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIncludePajwohLqCertificateRule) implementsAccountAccessGroupUpdateResponseResultInclude() {
}

// Matches an Access group.
type AccountAccessGroupUpdateResponseResultIncludePajwohLqAccessGroupRule struct {
	Group AccountAccessGroupUpdateResponseResultIncludePajwohLqAccessGroupRuleGroup `json:"group,required"`
	JSON  accountAccessGroupUpdateResponseResultIncludePajwohLqAccessGroupRuleJSON  `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqAccessGroupRule]
type accountAccessGroupUpdateResponseResultIncludePajwohLqAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIncludePajwohLqAccessGroupRule) implementsAccountAccessGroupUpdateResponseResultInclude() {
}

type AccountAccessGroupUpdateResponseResultIncludePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                        `json:"id,required"`
	JSON accountAccessGroupUpdateResponseResultIncludePajwohLqAccessGroupRuleGroupJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqAccessGroupRuleGroup]
type accountAccessGroupUpdateResponseResultIncludePajwohLqAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccountAccessGroupUpdateResponseResultIncludePajwohLqAzureGroupRule struct {
	AzureAd AccountAccessGroupUpdateResponseResultIncludePajwohLqAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accountAccessGroupUpdateResponseResultIncludePajwohLqAzureGroupRuleJSON    `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqAzureGroupRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqAzureGroupRule]
type accountAccessGroupUpdateResponseResultIncludePajwohLqAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIncludePajwohLqAzureGroupRule) implementsAccountAccessGroupUpdateResponseResultInclude() {
}

type AccountAccessGroupUpdateResponseResultIncludePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                         `json:"connection_id,required"`
	JSON         accountAccessGroupUpdateResponseResultIncludePajwohLqAzureGroupRuleAzureAdJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqAzureGroupRuleAzureAd]
type accountAccessGroupUpdateResponseResultIncludePajwohLqAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccountAccessGroupUpdateResponseResultIncludePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization AccountAccessGroupUpdateResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accountAccessGroupUpdateResponseResultIncludePajwohLqGitHubOrganizationRuleJSON               `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqGitHubOrganizationRule]
type accountAccessGroupUpdateResponseResultIncludePajwohLqGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIncludePajwohLqGitHubOrganizationRule) implementsAccountAccessGroupUpdateResponseResultInclude() {
}

type AccountAccessGroupUpdateResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                            `json:"name,required"`
	JSON accountAccessGroupUpdateResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization]
type accountAccessGroupUpdateResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccountAccessGroupUpdateResponseResultIncludePajwohLqGsuiteGroupRule struct {
	Gsuite AccountAccessGroupUpdateResponseResultIncludePajwohLqGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accountAccessGroupUpdateResponseResultIncludePajwohLqGsuiteGroupRuleJSON   `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqGsuiteGroupRule]
type accountAccessGroupUpdateResponseResultIncludePajwohLqGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIncludePajwohLqGsuiteGroupRule) implementsAccountAccessGroupUpdateResponseResultInclude() {
}

type AccountAccessGroupUpdateResponseResultIncludePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                         `json:"email,required"`
	JSON  accountAccessGroupUpdateResponseResultIncludePajwohLqGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqGsuiteGroupRuleGsuite]
type accountAccessGroupUpdateResponseResultIncludePajwohLqGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccountAccessGroupUpdateResponseResultIncludePajwohLqOktaGroupRule struct {
	Okta AccountAccessGroupUpdateResponseResultIncludePajwohLqOktaGroupRuleOkta `json:"okta,required"`
	JSON accountAccessGroupUpdateResponseResultIncludePajwohLqOktaGroupRuleJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqOktaGroupRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqOktaGroupRule]
type accountAccessGroupUpdateResponseResultIncludePajwohLqOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIncludePajwohLqOktaGroupRule) implementsAccountAccessGroupUpdateResponseResultInclude() {
}

type AccountAccessGroupUpdateResponseResultIncludePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                     `json:"email,required"`
	JSON  accountAccessGroupUpdateResponseResultIncludePajwohLqOktaGroupRuleOktaJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqOktaGroupRuleOkta]
type accountAccessGroupUpdateResponseResultIncludePajwohLqOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccountAccessGroupUpdateResponseResultIncludePajwohLqSamlGroupRule struct {
	Saml AccountAccessGroupUpdateResponseResultIncludePajwohLqSamlGroupRuleSaml `json:"saml,required"`
	JSON accountAccessGroupUpdateResponseResultIncludePajwohLqSamlGroupRuleJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqSamlGroupRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqSamlGroupRule]
type accountAccessGroupUpdateResponseResultIncludePajwohLqSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIncludePajwohLqSamlGroupRule) implementsAccountAccessGroupUpdateResponseResultInclude() {
}

type AccountAccessGroupUpdateResponseResultIncludePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                     `json:"attribute_value,required"`
	JSON           accountAccessGroupUpdateResponseResultIncludePajwohLqSamlGroupRuleSamlJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqSamlGroupRuleSaml]
type accountAccessGroupUpdateResponseResultIncludePajwohLqSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccountAccessGroupUpdateResponseResultIncludePajwohLqServiceTokenRule struct {
	ServiceToken AccountAccessGroupUpdateResponseResultIncludePajwohLqServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accountAccessGroupUpdateResponseResultIncludePajwohLqServiceTokenRuleJSON         `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqServiceTokenRule]
type accountAccessGroupUpdateResponseResultIncludePajwohLqServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIncludePajwohLqServiceTokenRule) implementsAccountAccessGroupUpdateResponseResultInclude() {
}

type AccountAccessGroupUpdateResponseResultIncludePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                `json:"token_id,required"`
	JSON    accountAccessGroupUpdateResponseResultIncludePajwohLqServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqServiceTokenRuleServiceToken]
type accountAccessGroupUpdateResponseResultIncludePajwohLqServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccountAccessGroupUpdateResponseResultIncludePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                       `json:"any_valid_service_token,required"`
	JSON                 accountAccessGroupUpdateResponseResultIncludePajwohLqAnyValidServiceTokenRuleJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqAnyValidServiceTokenRule]
type accountAccessGroupUpdateResponseResultIncludePajwohLqAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIncludePajwohLqAnyValidServiceTokenRule) implementsAccountAccessGroupUpdateResponseResultInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccountAccessGroupUpdateResponseResultIncludePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation AccountAccessGroupUpdateResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accountAccessGroupUpdateResponseResultIncludePajwohLqExternalEvaluationRuleJSON               `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqExternalEvaluationRule]
type accountAccessGroupUpdateResponseResultIncludePajwohLqExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIncludePajwohLqExternalEvaluationRule) implementsAccountAccessGroupUpdateResponseResultInclude() {
}

type AccountAccessGroupUpdateResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                            `json:"keys_url,required"`
	JSON    accountAccessGroupUpdateResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation]
type accountAccessGroupUpdateResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccountAccessGroupUpdateResponseResultIncludePajwohLqCountryRule struct {
	Geo  AccountAccessGroupUpdateResponseResultIncludePajwohLqCountryRuleGeo  `json:"geo,required"`
	JSON accountAccessGroupUpdateResponseResultIncludePajwohLqCountryRuleJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqCountryRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqCountryRule]
type accountAccessGroupUpdateResponseResultIncludePajwohLqCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIncludePajwohLqCountryRule) implementsAccountAccessGroupUpdateResponseResultInclude() {
}

type AccountAccessGroupUpdateResponseResultIncludePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                  `json:"country_code,required"`
	JSON        accountAccessGroupUpdateResponseResultIncludePajwohLqCountryRuleGeoJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqCountryRuleGeoJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqCountryRuleGeo]
type accountAccessGroupUpdateResponseResultIncludePajwohLqCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccountAccessGroupUpdateResponseResultIncludePajwohLqAuthenticationMethodRule struct {
	AuthMethod AccountAccessGroupUpdateResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accountAccessGroupUpdateResponseResultIncludePajwohLqAuthenticationMethodRuleJSON       `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqAuthenticationMethodRule]
type accountAccessGroupUpdateResponseResultIncludePajwohLqAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIncludePajwohLqAuthenticationMethodRule) implementsAccountAccessGroupUpdateResponseResultInclude() {
}

type AccountAccessGroupUpdateResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                      `json:"auth_method,required"`
	JSON       accountAccessGroupUpdateResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod]
type accountAccessGroupUpdateResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccountAccessGroupUpdateResponseResultIncludePajwohLqDevicePostureRule struct {
	DevicePosture AccountAccessGroupUpdateResponseResultIncludePajwohLqDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accountAccessGroupUpdateResponseResultIncludePajwohLqDevicePostureRuleJSON          `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqDevicePostureRule]
type accountAccessGroupUpdateResponseResultIncludePajwohLqDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIncludePajwohLqDevicePostureRule) implementsAccountAccessGroupUpdateResponseResultInclude() {
}

type AccountAccessGroupUpdateResponseResultIncludePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                  `json:"integration_uid,required"`
	JSON           accountAccessGroupUpdateResponseResultIncludePajwohLqDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIncludePajwohLqDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIncludePajwohLqDevicePostureRuleDevicePosture]
type accountAccessGroupUpdateResponseResultIncludePajwohLqDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIncludePajwohLqDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqEmailRule],
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqEmailListRule],
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqDomainRule],
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqEveryoneRule],
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqIPRule],
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqIPListRule],
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqCertificateRule],
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAccessGroupRule],
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAzureGroupRule],
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqGitHubOrganizationRule],
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqGsuiteGroupRule],
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqOktaGroupRule],
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqSamlGroupRule],
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqServiceTokenRule],
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAnyValidServiceTokenRule],
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqExternalEvaluationRule],
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqCountryRule],
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAuthenticationMethodRule]
// or [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqDevicePostureRule].
type AccountAccessGroupUpdateResponseResultIsDefault interface {
	implementsAccountAccessGroupUpdateResponseResultIsDefault()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountAccessGroupUpdateResponseResultIsDefault)(nil)).Elem(), "")
}

// Matches a specific email.
type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqEmailRule struct {
	Email AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqEmailRuleEmail `json:"email,required"`
	JSON  accountAccessGroupUpdateResponseResultIsDefaultPajwohLqEmailRuleJSON  `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqEmailRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqEmailRule]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqEmailRule) implementsAccountAccessGroupUpdateResponseResultIsDefault() {
}

type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                    `json:"email,required" format:"email"`
	JSON  accountAccessGroupUpdateResponseResultIsDefaultPajwohLqEmailRuleEmailJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqEmailRuleEmail]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqEmailListRule struct {
	EmailList AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqEmailListRuleEmailList `json:"email_list,required"`
	JSON      accountAccessGroupUpdateResponseResultIsDefaultPajwohLqEmailListRuleJSON      `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqEmailListRule]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqEmailListRule) implementsAccountAccessGroupUpdateResponseResultIsDefault() {
}

type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                            `json:"id,required"`
	JSON accountAccessGroupUpdateResponseResultIsDefaultPajwohLqEmailListRuleEmailListJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqEmailListRuleEmailList]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqDomainRule struct {
	EmailDomain AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accountAccessGroupUpdateResponseResultIsDefaultPajwohLqDomainRuleJSON        `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqDomainRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqDomainRule]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqDomainRule) implementsAccountAccessGroupUpdateResponseResultIsDefault() {
}

type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                           `json:"domain,required"`
	JSON   accountAccessGroupUpdateResponseResultIsDefaultPajwohLqDomainRuleEmailDomainJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqDomainRuleEmailDomain]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                             `json:"everyone,required"`
	JSON     accountAccessGroupUpdateResponseResultIsDefaultPajwohLqEveryoneRuleJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqEveryoneRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqEveryoneRule]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqEveryoneRule) implementsAccountAccessGroupUpdateResponseResultIsDefault() {
}

// Matches an IP address block.
type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqIPRule struct {
	IP   AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqIPRuleIP   `json:"ip,required"`
	JSON accountAccessGroupUpdateResponseResultIsDefaultPajwohLqIPRuleJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqIPRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqIPRule]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqIPRule) implementsAccountAccessGroupUpdateResponseResultIsDefault() {
}

type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                              `json:"ip,required"`
	JSON accountAccessGroupUpdateResponseResultIsDefaultPajwohLqIPRuleIPJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqIPRuleIPJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqIPRuleIP]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqIPListRule struct {
	IPList AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqIPListRuleIPList `json:"ip_list,required"`
	JSON   accountAccessGroupUpdateResponseResultIsDefaultPajwohLqIPListRuleJSON   `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqIPListRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqIPListRule]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqIPListRule) implementsAccountAccessGroupUpdateResponseResultIsDefault() {
}

type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                      `json:"id,required"`
	JSON accountAccessGroupUpdateResponseResultIsDefaultPajwohLqIPListRuleIPListJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqIPListRuleIPList]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqCertificateRule struct {
	Certificate interface{}                                                                `json:"certificate,required"`
	JSON        accountAccessGroupUpdateResponseResultIsDefaultPajwohLqCertificateRuleJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqCertificateRule]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqCertificateRule) implementsAccountAccessGroupUpdateResponseResultIsDefault() {
}

// Matches an Access group.
type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAccessGroupRule struct {
	Group AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAccessGroupRuleGroup `json:"group,required"`
	JSON  accountAccessGroupUpdateResponseResultIsDefaultPajwohLqAccessGroupRuleJSON  `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAccessGroupRule]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAccessGroupRule) implementsAccountAccessGroupUpdateResponseResultIsDefault() {
}

type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                          `json:"id,required"`
	JSON accountAccessGroupUpdateResponseResultIsDefaultPajwohLqAccessGroupRuleGroupJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAccessGroupRuleGroup]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAzureGroupRule struct {
	AzureAd AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accountAccessGroupUpdateResponseResultIsDefaultPajwohLqAzureGroupRuleJSON    `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAzureGroupRule]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAzureGroupRule) implementsAccountAccessGroupUpdateResponseResultIsDefault() {
}

type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                           `json:"connection_id,required"`
	JSON         accountAccessGroupUpdateResponseResultIsDefaultPajwohLqAzureGroupRuleAzureAdJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAzureGroupRuleAzureAd]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqGitHubOrganizationRule struct {
	GitHubOrganization AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accountAccessGroupUpdateResponseResultIsDefaultPajwohLqGitHubOrganizationRuleJSON               `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqGitHubOrganizationRule]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqGitHubOrganizationRule) implementsAccountAccessGroupUpdateResponseResultIsDefault() {
}

type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                              `json:"name,required"`
	JSON accountAccessGroupUpdateResponseResultIsDefaultPajwohLqGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqGitHubOrganizationRuleGitHubOrganization]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqGsuiteGroupRule struct {
	Gsuite AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accountAccessGroupUpdateResponseResultIsDefaultPajwohLqGsuiteGroupRuleJSON   `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqGsuiteGroupRule]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqGsuiteGroupRule) implementsAccountAccessGroupUpdateResponseResultIsDefault() {
}

type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                           `json:"email,required"`
	JSON  accountAccessGroupUpdateResponseResultIsDefaultPajwohLqGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqGsuiteGroupRuleGsuite]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqOktaGroupRule struct {
	Okta AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqOktaGroupRuleOkta `json:"okta,required"`
	JSON accountAccessGroupUpdateResponseResultIsDefaultPajwohLqOktaGroupRuleJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqOktaGroupRule]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqOktaGroupRule) implementsAccountAccessGroupUpdateResponseResultIsDefault() {
}

type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                       `json:"email,required"`
	JSON  accountAccessGroupUpdateResponseResultIsDefaultPajwohLqOktaGroupRuleOktaJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqOktaGroupRuleOkta]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqSamlGroupRule struct {
	Saml AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqSamlGroupRuleSaml `json:"saml,required"`
	JSON accountAccessGroupUpdateResponseResultIsDefaultPajwohLqSamlGroupRuleJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqSamlGroupRule]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqSamlGroupRule) implementsAccountAccessGroupUpdateResponseResultIsDefault() {
}

type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                       `json:"attribute_value,required"`
	JSON           accountAccessGroupUpdateResponseResultIsDefaultPajwohLqSamlGroupRuleSamlJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqSamlGroupRuleSaml]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqServiceTokenRule struct {
	ServiceToken AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accountAccessGroupUpdateResponseResultIsDefaultPajwohLqServiceTokenRuleJSON         `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqServiceTokenRule]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqServiceTokenRule) implementsAccountAccessGroupUpdateResponseResultIsDefault() {
}

type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                  `json:"token_id,required"`
	JSON    accountAccessGroupUpdateResponseResultIsDefaultPajwohLqServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqServiceTokenRuleServiceToken]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                         `json:"any_valid_service_token,required"`
	JSON                 accountAccessGroupUpdateResponseResultIsDefaultPajwohLqAnyValidServiceTokenRuleJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAnyValidServiceTokenRule]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAnyValidServiceTokenRule) implementsAccountAccessGroupUpdateResponseResultIsDefault() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqExternalEvaluationRule struct {
	ExternalEvaluation AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accountAccessGroupUpdateResponseResultIsDefaultPajwohLqExternalEvaluationRuleJSON               `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqExternalEvaluationRule]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqExternalEvaluationRule) implementsAccountAccessGroupUpdateResponseResultIsDefault() {
}

type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                              `json:"keys_url,required"`
	JSON    accountAccessGroupUpdateResponseResultIsDefaultPajwohLqExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqExternalEvaluationRuleExternalEvaluation]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqCountryRule struct {
	Geo  AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqCountryRuleGeo  `json:"geo,required"`
	JSON accountAccessGroupUpdateResponseResultIsDefaultPajwohLqCountryRuleJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqCountryRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqCountryRule]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqCountryRule) implementsAccountAccessGroupUpdateResponseResultIsDefault() {
}

type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                    `json:"country_code,required"`
	JSON        accountAccessGroupUpdateResponseResultIsDefaultPajwohLqCountryRuleGeoJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqCountryRuleGeo]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAuthenticationMethodRule struct {
	AuthMethod AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accountAccessGroupUpdateResponseResultIsDefaultPajwohLqAuthenticationMethodRuleJSON       `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAuthenticationMethodRule]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAuthenticationMethodRule) implementsAccountAccessGroupUpdateResponseResultIsDefault() {
}

type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                        `json:"auth_method,required"`
	JSON       accountAccessGroupUpdateResponseResultIsDefaultPajwohLqAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAuthenticationMethodRuleAuthMethod]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqDevicePostureRule struct {
	DevicePosture AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accountAccessGroupUpdateResponseResultIsDefaultPajwohLqDevicePostureRuleJSON          `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqDevicePostureRule]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqDevicePostureRule) implementsAccountAccessGroupUpdateResponseResultIsDefault() {
}

type AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                    `json:"integration_uid,required"`
	JSON           accountAccessGroupUpdateResponseResultIsDefaultPajwohLqDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultIsDefaultPajwohLqDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqDevicePostureRuleDevicePosture]
type accountAccessGroupUpdateResponseResultIsDefaultPajwohLqDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultIsDefaultPajwohLqDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqEmailRule],
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqEmailListRule],
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqDomainRule],
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqEveryoneRule],
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqIPRule],
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqIPListRule],
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqCertificateRule],
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqAccessGroupRule],
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqAzureGroupRule],
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqGitHubOrganizationRule],
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqGsuiteGroupRule],
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqOktaGroupRule],
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqSamlGroupRule],
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqServiceTokenRule],
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqAnyValidServiceTokenRule],
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqExternalEvaluationRule],
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqCountryRule],
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqAuthenticationMethodRule]
// or [AccountAccessGroupUpdateResponseResultRequirePajwohLqDevicePostureRule].
type AccountAccessGroupUpdateResponseResultRequire interface {
	implementsAccountAccessGroupUpdateResponseResultRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountAccessGroupUpdateResponseResultRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type AccountAccessGroupUpdateResponseResultRequirePajwohLqEmailRule struct {
	Email AccountAccessGroupUpdateResponseResultRequirePajwohLqEmailRuleEmail `json:"email,required"`
	JSON  accountAccessGroupUpdateResponseResultRequirePajwohLqEmailRuleJSON  `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqEmailRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqEmailRule]
type accountAccessGroupUpdateResponseResultRequirePajwohLqEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultRequirePajwohLqEmailRule) implementsAccountAccessGroupUpdateResponseResultRequire() {
}

type AccountAccessGroupUpdateResponseResultRequirePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                  `json:"email,required" format:"email"`
	JSON  accountAccessGroupUpdateResponseResultRequirePajwohLqEmailRuleEmailJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqEmailRuleEmailJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqEmailRuleEmail]
type accountAccessGroupUpdateResponseResultRequirePajwohLqEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccountAccessGroupUpdateResponseResultRequirePajwohLqEmailListRule struct {
	EmailList AccountAccessGroupUpdateResponseResultRequirePajwohLqEmailListRuleEmailList `json:"email_list,required"`
	JSON      accountAccessGroupUpdateResponseResultRequirePajwohLqEmailListRuleJSON      `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqEmailListRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqEmailListRule]
type accountAccessGroupUpdateResponseResultRequirePajwohLqEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultRequirePajwohLqEmailListRule) implementsAccountAccessGroupUpdateResponseResultRequire() {
}

type AccountAccessGroupUpdateResponseResultRequirePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                          `json:"id,required"`
	JSON accountAccessGroupUpdateResponseResultRequirePajwohLqEmailListRuleEmailListJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqEmailListRuleEmailList]
type accountAccessGroupUpdateResponseResultRequirePajwohLqEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccountAccessGroupUpdateResponseResultRequirePajwohLqDomainRule struct {
	EmailDomain AccountAccessGroupUpdateResponseResultRequirePajwohLqDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accountAccessGroupUpdateResponseResultRequirePajwohLqDomainRuleJSON        `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqDomainRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqDomainRule]
type accountAccessGroupUpdateResponseResultRequirePajwohLqDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultRequirePajwohLqDomainRule) implementsAccountAccessGroupUpdateResponseResultRequire() {
}

type AccountAccessGroupUpdateResponseResultRequirePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                         `json:"domain,required"`
	JSON   accountAccessGroupUpdateResponseResultRequirePajwohLqDomainRuleEmailDomainJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqDomainRuleEmailDomain]
type accountAccessGroupUpdateResponseResultRequirePajwohLqDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccountAccessGroupUpdateResponseResultRequirePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                           `json:"everyone,required"`
	JSON     accountAccessGroupUpdateResponseResultRequirePajwohLqEveryoneRuleJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqEveryoneRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqEveryoneRule]
type accountAccessGroupUpdateResponseResultRequirePajwohLqEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultRequirePajwohLqEveryoneRule) implementsAccountAccessGroupUpdateResponseResultRequire() {
}

// Matches an IP address block.
type AccountAccessGroupUpdateResponseResultRequirePajwohLqIPRule struct {
	IP   AccountAccessGroupUpdateResponseResultRequirePajwohLqIPRuleIP   `json:"ip,required"`
	JSON accountAccessGroupUpdateResponseResultRequirePajwohLqIPRuleJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqIPRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqIPRule]
type accountAccessGroupUpdateResponseResultRequirePajwohLqIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultRequirePajwohLqIPRule) implementsAccountAccessGroupUpdateResponseResultRequire() {
}

type AccountAccessGroupUpdateResponseResultRequirePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                            `json:"ip,required"`
	JSON accountAccessGroupUpdateResponseResultRequirePajwohLqIPRuleIPJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqIPRuleIPJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqIPRuleIP]
type accountAccessGroupUpdateResponseResultRequirePajwohLqIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccountAccessGroupUpdateResponseResultRequirePajwohLqIPListRule struct {
	IPList AccountAccessGroupUpdateResponseResultRequirePajwohLqIPListRuleIPList `json:"ip_list,required"`
	JSON   accountAccessGroupUpdateResponseResultRequirePajwohLqIPListRuleJSON   `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqIPListRuleJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqIPListRule]
type accountAccessGroupUpdateResponseResultRequirePajwohLqIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultRequirePajwohLqIPListRule) implementsAccountAccessGroupUpdateResponseResultRequire() {
}

type AccountAccessGroupUpdateResponseResultRequirePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                    `json:"id,required"`
	JSON accountAccessGroupUpdateResponseResultRequirePajwohLqIPListRuleIPListJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqIPListRuleIPList]
type accountAccessGroupUpdateResponseResultRequirePajwohLqIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccountAccessGroupUpdateResponseResultRequirePajwohLqCertificateRule struct {
	Certificate interface{}                                                              `json:"certificate,required"`
	JSON        accountAccessGroupUpdateResponseResultRequirePajwohLqCertificateRuleJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqCertificateRule]
type accountAccessGroupUpdateResponseResultRequirePajwohLqCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultRequirePajwohLqCertificateRule) implementsAccountAccessGroupUpdateResponseResultRequire() {
}

// Matches an Access group.
type AccountAccessGroupUpdateResponseResultRequirePajwohLqAccessGroupRule struct {
	Group AccountAccessGroupUpdateResponseResultRequirePajwohLqAccessGroupRuleGroup `json:"group,required"`
	JSON  accountAccessGroupUpdateResponseResultRequirePajwohLqAccessGroupRuleJSON  `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqAccessGroupRule]
type accountAccessGroupUpdateResponseResultRequirePajwohLqAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultRequirePajwohLqAccessGroupRule) implementsAccountAccessGroupUpdateResponseResultRequire() {
}

type AccountAccessGroupUpdateResponseResultRequirePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                        `json:"id,required"`
	JSON accountAccessGroupUpdateResponseResultRequirePajwohLqAccessGroupRuleGroupJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqAccessGroupRuleGroup]
type accountAccessGroupUpdateResponseResultRequirePajwohLqAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccountAccessGroupUpdateResponseResultRequirePajwohLqAzureGroupRule struct {
	AzureAd AccountAccessGroupUpdateResponseResultRequirePajwohLqAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accountAccessGroupUpdateResponseResultRequirePajwohLqAzureGroupRuleJSON    `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqAzureGroupRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqAzureGroupRule]
type accountAccessGroupUpdateResponseResultRequirePajwohLqAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultRequirePajwohLqAzureGroupRule) implementsAccountAccessGroupUpdateResponseResultRequire() {
}

type AccountAccessGroupUpdateResponseResultRequirePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                         `json:"connection_id,required"`
	JSON         accountAccessGroupUpdateResponseResultRequirePajwohLqAzureGroupRuleAzureAdJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqAzureGroupRuleAzureAd]
type accountAccessGroupUpdateResponseResultRequirePajwohLqAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccountAccessGroupUpdateResponseResultRequirePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization AccountAccessGroupUpdateResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accountAccessGroupUpdateResponseResultRequirePajwohLqGitHubOrganizationRuleJSON               `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqGitHubOrganizationRule]
type accountAccessGroupUpdateResponseResultRequirePajwohLqGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultRequirePajwohLqGitHubOrganizationRule) implementsAccountAccessGroupUpdateResponseResultRequire() {
}

type AccountAccessGroupUpdateResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                            `json:"name,required"`
	JSON accountAccessGroupUpdateResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization]
type accountAccessGroupUpdateResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccountAccessGroupUpdateResponseResultRequirePajwohLqGsuiteGroupRule struct {
	Gsuite AccountAccessGroupUpdateResponseResultRequirePajwohLqGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accountAccessGroupUpdateResponseResultRequirePajwohLqGsuiteGroupRuleJSON   `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqGsuiteGroupRule]
type accountAccessGroupUpdateResponseResultRequirePajwohLqGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultRequirePajwohLqGsuiteGroupRule) implementsAccountAccessGroupUpdateResponseResultRequire() {
}

type AccountAccessGroupUpdateResponseResultRequirePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                         `json:"email,required"`
	JSON  accountAccessGroupUpdateResponseResultRequirePajwohLqGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqGsuiteGroupRuleGsuite]
type accountAccessGroupUpdateResponseResultRequirePajwohLqGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccountAccessGroupUpdateResponseResultRequirePajwohLqOktaGroupRule struct {
	Okta AccountAccessGroupUpdateResponseResultRequirePajwohLqOktaGroupRuleOkta `json:"okta,required"`
	JSON accountAccessGroupUpdateResponseResultRequirePajwohLqOktaGroupRuleJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqOktaGroupRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqOktaGroupRule]
type accountAccessGroupUpdateResponseResultRequirePajwohLqOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultRequirePajwohLqOktaGroupRule) implementsAccountAccessGroupUpdateResponseResultRequire() {
}

type AccountAccessGroupUpdateResponseResultRequirePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                     `json:"email,required"`
	JSON  accountAccessGroupUpdateResponseResultRequirePajwohLqOktaGroupRuleOktaJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqOktaGroupRuleOkta]
type accountAccessGroupUpdateResponseResultRequirePajwohLqOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccountAccessGroupUpdateResponseResultRequirePajwohLqSamlGroupRule struct {
	Saml AccountAccessGroupUpdateResponseResultRequirePajwohLqSamlGroupRuleSaml `json:"saml,required"`
	JSON accountAccessGroupUpdateResponseResultRequirePajwohLqSamlGroupRuleJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqSamlGroupRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqSamlGroupRule]
type accountAccessGroupUpdateResponseResultRequirePajwohLqSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultRequirePajwohLqSamlGroupRule) implementsAccountAccessGroupUpdateResponseResultRequire() {
}

type AccountAccessGroupUpdateResponseResultRequirePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                     `json:"attribute_value,required"`
	JSON           accountAccessGroupUpdateResponseResultRequirePajwohLqSamlGroupRuleSamlJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqSamlGroupRuleSaml]
type accountAccessGroupUpdateResponseResultRequirePajwohLqSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccountAccessGroupUpdateResponseResultRequirePajwohLqServiceTokenRule struct {
	ServiceToken AccountAccessGroupUpdateResponseResultRequirePajwohLqServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accountAccessGroupUpdateResponseResultRequirePajwohLqServiceTokenRuleJSON         `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqServiceTokenRule]
type accountAccessGroupUpdateResponseResultRequirePajwohLqServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultRequirePajwohLqServiceTokenRule) implementsAccountAccessGroupUpdateResponseResultRequire() {
}

type AccountAccessGroupUpdateResponseResultRequirePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                `json:"token_id,required"`
	JSON    accountAccessGroupUpdateResponseResultRequirePajwohLqServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqServiceTokenRuleServiceToken]
type accountAccessGroupUpdateResponseResultRequirePajwohLqServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccountAccessGroupUpdateResponseResultRequirePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                       `json:"any_valid_service_token,required"`
	JSON                 accountAccessGroupUpdateResponseResultRequirePajwohLqAnyValidServiceTokenRuleJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqAnyValidServiceTokenRule]
type accountAccessGroupUpdateResponseResultRequirePajwohLqAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultRequirePajwohLqAnyValidServiceTokenRule) implementsAccountAccessGroupUpdateResponseResultRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccountAccessGroupUpdateResponseResultRequirePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation AccountAccessGroupUpdateResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accountAccessGroupUpdateResponseResultRequirePajwohLqExternalEvaluationRuleJSON               `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqExternalEvaluationRule]
type accountAccessGroupUpdateResponseResultRequirePajwohLqExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultRequirePajwohLqExternalEvaluationRule) implementsAccountAccessGroupUpdateResponseResultRequire() {
}

type AccountAccessGroupUpdateResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                            `json:"keys_url,required"`
	JSON    accountAccessGroupUpdateResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation]
type accountAccessGroupUpdateResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccountAccessGroupUpdateResponseResultRequirePajwohLqCountryRule struct {
	Geo  AccountAccessGroupUpdateResponseResultRequirePajwohLqCountryRuleGeo  `json:"geo,required"`
	JSON accountAccessGroupUpdateResponseResultRequirePajwohLqCountryRuleJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqCountryRuleJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqCountryRule]
type accountAccessGroupUpdateResponseResultRequirePajwohLqCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultRequirePajwohLqCountryRule) implementsAccountAccessGroupUpdateResponseResultRequire() {
}

type AccountAccessGroupUpdateResponseResultRequirePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                  `json:"country_code,required"`
	JSON        accountAccessGroupUpdateResponseResultRequirePajwohLqCountryRuleGeoJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqCountryRuleGeoJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqCountryRuleGeo]
type accountAccessGroupUpdateResponseResultRequirePajwohLqCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccountAccessGroupUpdateResponseResultRequirePajwohLqAuthenticationMethodRule struct {
	AuthMethod AccountAccessGroupUpdateResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accountAccessGroupUpdateResponseResultRequirePajwohLqAuthenticationMethodRuleJSON       `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqAuthenticationMethodRule]
type accountAccessGroupUpdateResponseResultRequirePajwohLqAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultRequirePajwohLqAuthenticationMethodRule) implementsAccountAccessGroupUpdateResponseResultRequire() {
}

type AccountAccessGroupUpdateResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                      `json:"auth_method,required"`
	JSON       accountAccessGroupUpdateResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod]
type accountAccessGroupUpdateResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccountAccessGroupUpdateResponseResultRequirePajwohLqDevicePostureRule struct {
	DevicePosture AccountAccessGroupUpdateResponseResultRequirePajwohLqDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accountAccessGroupUpdateResponseResultRequirePajwohLqDevicePostureRuleJSON          `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqDevicePostureRule]
type accountAccessGroupUpdateResponseResultRequirePajwohLqDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupUpdateResponseResultRequirePajwohLqDevicePostureRule) implementsAccountAccessGroupUpdateResponseResultRequire() {
}

type AccountAccessGroupUpdateResponseResultRequirePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                  `json:"integration_uid,required"`
	JSON           accountAccessGroupUpdateResponseResultRequirePajwohLqDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accountAccessGroupUpdateResponseResultRequirePajwohLqDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupUpdateResponseResultRequirePajwohLqDevicePostureRuleDevicePosture]
type accountAccessGroupUpdateResponseResultRequirePajwohLqDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupUpdateResponseResultRequirePajwohLqDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessGroupUpdateResponseSuccess bool

const (
	AccountAccessGroupUpdateResponseSuccessTrue AccountAccessGroupUpdateResponseSuccess = true
)

type AccountAccessGroupDeleteResponse struct {
	Errors   []AccountAccessGroupDeleteResponseError   `json:"errors"`
	Messages []AccountAccessGroupDeleteResponseMessage `json:"messages"`
	Result   AccountAccessGroupDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessGroupDeleteResponseSuccess `json:"success"`
	JSON    accountAccessGroupDeleteResponseJSON    `json:"-"`
}

// accountAccessGroupDeleteResponseJSON contains the JSON metadata for the struct
// [AccountAccessGroupDeleteResponse]
type accountAccessGroupDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessGroupDeleteResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountAccessGroupDeleteResponseErrorJSON `json:"-"`
}

// accountAccessGroupDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountAccessGroupDeleteResponseError]
type accountAccessGroupDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessGroupDeleteResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountAccessGroupDeleteResponseMessageJSON `json:"-"`
}

// accountAccessGroupDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountAccessGroupDeleteResponseMessage]
type accountAccessGroupDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessGroupDeleteResponseResult struct {
	// UUID
	ID   string                                     `json:"id"`
	JSON accountAccessGroupDeleteResponseResultJSON `json:"-"`
}

// accountAccessGroupDeleteResponseResultJSON contains the JSON metadata for the
// struct [AccountAccessGroupDeleteResponseResult]
type accountAccessGroupDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessGroupDeleteResponseSuccess bool

const (
	AccountAccessGroupDeleteResponseSuccessTrue AccountAccessGroupDeleteResponseSuccess = true
)

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponse struct {
	Errors   []AccountAccessGroupAccessGroupsNewAnAccessGroupResponseError   `json:"errors"`
	Messages []AccountAccessGroupAccessGroupsNewAnAccessGroupResponseMessage `json:"messages"`
	Result   AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessGroupAccessGroupsNewAnAccessGroupResponseSuccess `json:"success"`
	JSON    accountAccessGroupAccessGroupsNewAnAccessGroupResponseJSON    `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseJSON contains the JSON
// metadata for the struct [AccountAccessGroupAccessGroupsNewAnAccessGroupResponse]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseError struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    accountAccessGroupAccessGroupsNewAnAccessGroupResponseErrorJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseError]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseMessage struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    accountAccessGroupAccessGroupsNewAnAccessGroupResponseMessageJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseMessageJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseMessage]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResult struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude `json:"include"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	IsDefault []AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault `json:"is_default"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire `json:"require"`
	UpdatedAt time.Time                                                             `json:"updated_at" format:"date-time"`
	JSON      accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultJSON      `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResult]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultJSON struct {
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

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEmailRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEmailListRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqDomainRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEveryoneRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqIPRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqIPListRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqCertificateRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAccessGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAzureGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqGitHubOrganizationRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqGsuiteGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqOktaGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqSamlGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqServiceTokenRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAnyValidServiceTokenRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqExternalEvaluationRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqCountryRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAuthenticationMethodRule]
// or
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqDevicePostureRule].
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude interface {
	implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEmailRule struct {
	Email AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEmailRuleEmail `json:"email,required"`
	JSON  accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEmailRuleJSON  `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEmailRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEmailRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEmailRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                                        `json:"email,required" format:"email"`
	JSON  accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEmailRuleEmailJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEmailRuleEmail]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEmailListRule struct {
	EmailList AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEmailListRuleEmailList `json:"email_list,required"`
	JSON      accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEmailListRuleJSON      `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEmailListRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEmailListRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                                `json:"id,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEmailListRuleEmailListJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEmailListRuleEmailList]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqDomainRule struct {
	EmailDomain AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqDomainRuleJSON        `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqDomainRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqDomainRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqDomainRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                               `json:"domain,required"`
	JSON   accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqDomainRuleEmailDomainJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqDomainRuleEmailDomain]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                                 `json:"everyone,required"`
	JSON     accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEveryoneRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEveryoneRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqEveryoneRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

// Matches an IP address block.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqIPRule struct {
	IP   AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqIPRuleIP   `json:"ip,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqIPRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqIPRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqIPRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqIPRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                                  `json:"ip,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqIPRuleIPJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqIPRuleIP]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqIPListRule struct {
	IPList AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqIPListRuleIPList `json:"ip_list,required"`
	JSON   accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqIPListRuleJSON   `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqIPListRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqIPListRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqIPListRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                                          `json:"id,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqIPListRuleIPListJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqIPListRuleIPList]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqCertificateRule struct {
	Certificate interface{}                                                                                    `json:"certificate,required"`
	JSON        accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqCertificateRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqCertificateRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqCertificateRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

// Matches an Access group.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAccessGroupRule struct {
	Group AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAccessGroupRuleGroup `json:"group,required"`
	JSON  accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAccessGroupRuleJSON  `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAccessGroupRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAccessGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                                              `json:"id,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAccessGroupRuleGroupJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAccessGroupRuleGroup]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAzureGroupRule struct {
	AzureAd AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAzureGroupRuleJSON    `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAzureGroupRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAzureGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                               `json:"connection_id,required"`
	JSON         accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAzureGroupRuleAzureAdJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAzureGroupRuleAzureAd]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqGitHubOrganizationRuleJSON               `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqGitHubOrganizationRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqGitHubOrganizationRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                                  `json:"name,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqGsuiteGroupRule struct {
	Gsuite AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqGsuiteGroupRuleJSON   `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqGsuiteGroupRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqGsuiteGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                               `json:"email,required"`
	JSON  accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqGsuiteGroupRuleGsuite]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqOktaGroupRule struct {
	Okta AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqOktaGroupRuleOkta `json:"okta,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqOktaGroupRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqOktaGroupRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqOktaGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                                           `json:"email,required"`
	JSON  accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqOktaGroupRuleOktaJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqOktaGroupRuleOkta]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqSamlGroupRule struct {
	Saml AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqSamlGroupRuleSaml `json:"saml,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqSamlGroupRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqSamlGroupRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqSamlGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                                           `json:"attribute_value,required"`
	JSON           accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqSamlGroupRuleSamlJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqSamlGroupRuleSaml]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqServiceTokenRule struct {
	ServiceToken AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqServiceTokenRuleJSON         `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqServiceTokenRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqServiceTokenRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                                      `json:"token_id,required"`
	JSON    accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqServiceTokenRuleServiceToken]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                                             `json:"any_valid_service_token,required"`
	JSON                 accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAnyValidServiceTokenRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAnyValidServiceTokenRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAnyValidServiceTokenRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqExternalEvaluationRuleJSON               `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqExternalEvaluationRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqExternalEvaluationRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                                  `json:"keys_url,required"`
	JSON    accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqCountryRule struct {
	Geo  AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqCountryRuleGeo  `json:"geo,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqCountryRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqCountryRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqCountryRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqCountryRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                                        `json:"country_code,required"`
	JSON        accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqCountryRuleGeoJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqCountryRuleGeo]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAuthenticationMethodRule struct {
	AuthMethod AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAuthenticationMethodRuleJSON       `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAuthenticationMethodRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAuthenticationMethodRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                                            `json:"auth_method,required"`
	JSON       accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqDevicePostureRule struct {
	DevicePosture AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqDevicePostureRuleJSON          `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqDevicePostureRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqDevicePostureRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                                        `json:"integration_uid,required"`
	JSON           accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqDevicePostureRuleDevicePosture]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultExcludePajwohLqDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEmailRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEmailListRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqDomainRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEveryoneRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqIPRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqIPListRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqCertificateRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAccessGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAzureGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqGitHubOrganizationRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqGsuiteGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqOktaGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqSamlGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqServiceTokenRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAnyValidServiceTokenRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqExternalEvaluationRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqCountryRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAuthenticationMethodRule]
// or
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqDevicePostureRule].
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude interface {
	implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEmailRule struct {
	Email AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEmailRuleEmail `json:"email,required"`
	JSON  accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEmailRuleJSON  `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEmailRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEmailRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEmailRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                                        `json:"email,required" format:"email"`
	JSON  accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEmailRuleEmailJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEmailRuleEmail]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEmailListRule struct {
	EmailList AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEmailListRuleEmailList `json:"email_list,required"`
	JSON      accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEmailListRuleJSON      `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEmailListRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEmailListRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                                `json:"id,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEmailListRuleEmailListJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEmailListRuleEmailList]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqDomainRule struct {
	EmailDomain AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqDomainRuleJSON        `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqDomainRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqDomainRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqDomainRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                               `json:"domain,required"`
	JSON   accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqDomainRuleEmailDomainJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqDomainRuleEmailDomain]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                                 `json:"everyone,required"`
	JSON     accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEveryoneRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEveryoneRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqEveryoneRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

// Matches an IP address block.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqIPRule struct {
	IP   AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqIPRuleIP   `json:"ip,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqIPRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqIPRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqIPRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqIPRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                                  `json:"ip,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqIPRuleIPJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqIPRuleIP]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqIPListRule struct {
	IPList AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqIPListRuleIPList `json:"ip_list,required"`
	JSON   accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqIPListRuleJSON   `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqIPListRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqIPListRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqIPListRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                                          `json:"id,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqIPListRuleIPListJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqIPListRuleIPList]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqCertificateRule struct {
	Certificate interface{}                                                                                    `json:"certificate,required"`
	JSON        accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqCertificateRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqCertificateRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqCertificateRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

// Matches an Access group.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAccessGroupRule struct {
	Group AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAccessGroupRuleGroup `json:"group,required"`
	JSON  accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAccessGroupRuleJSON  `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAccessGroupRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAccessGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                                              `json:"id,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAccessGroupRuleGroupJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAccessGroupRuleGroup]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAzureGroupRule struct {
	AzureAd AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAzureGroupRuleJSON    `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAzureGroupRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAzureGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                               `json:"connection_id,required"`
	JSON         accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAzureGroupRuleAzureAdJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAzureGroupRuleAzureAd]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqGitHubOrganizationRuleJSON               `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqGitHubOrganizationRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqGitHubOrganizationRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                                  `json:"name,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqGsuiteGroupRule struct {
	Gsuite AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqGsuiteGroupRuleJSON   `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqGsuiteGroupRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqGsuiteGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                               `json:"email,required"`
	JSON  accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqGsuiteGroupRuleGsuite]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqOktaGroupRule struct {
	Okta AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqOktaGroupRuleOkta `json:"okta,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqOktaGroupRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqOktaGroupRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqOktaGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                                           `json:"email,required"`
	JSON  accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqOktaGroupRuleOktaJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqOktaGroupRuleOkta]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqSamlGroupRule struct {
	Saml AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqSamlGroupRuleSaml `json:"saml,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqSamlGroupRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqSamlGroupRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqSamlGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                                           `json:"attribute_value,required"`
	JSON           accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqSamlGroupRuleSamlJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqSamlGroupRuleSaml]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqServiceTokenRule struct {
	ServiceToken AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqServiceTokenRuleJSON         `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqServiceTokenRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqServiceTokenRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                                      `json:"token_id,required"`
	JSON    accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqServiceTokenRuleServiceToken]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                                             `json:"any_valid_service_token,required"`
	JSON                 accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAnyValidServiceTokenRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAnyValidServiceTokenRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAnyValidServiceTokenRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqExternalEvaluationRuleJSON               `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqExternalEvaluationRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqExternalEvaluationRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                                  `json:"keys_url,required"`
	JSON    accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqCountryRule struct {
	Geo  AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqCountryRuleGeo  `json:"geo,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqCountryRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqCountryRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqCountryRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqCountryRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                                        `json:"country_code,required"`
	JSON        accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqCountryRuleGeoJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqCountryRuleGeo]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAuthenticationMethodRule struct {
	AuthMethod AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAuthenticationMethodRuleJSON       `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAuthenticationMethodRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAuthenticationMethodRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                                            `json:"auth_method,required"`
	JSON       accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqDevicePostureRule struct {
	DevicePosture AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqDevicePostureRuleJSON          `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqDevicePostureRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqDevicePostureRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                                        `json:"integration_uid,required"`
	JSON           accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqDevicePostureRuleDevicePosture]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIncludePajwohLqDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEmailRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEmailListRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqDomainRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEveryoneRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqIPRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqIPListRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqCertificateRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAccessGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAzureGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqGitHubOrganizationRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqGsuiteGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqOktaGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqSamlGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqServiceTokenRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAnyValidServiceTokenRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqExternalEvaluationRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqCountryRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAuthenticationMethodRule]
// or
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqDevicePostureRule].
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault interface {
	implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault)(nil)).Elem(), "")
}

// Matches a specific email.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEmailRule struct {
	Email AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEmailRuleEmail `json:"email,required"`
	JSON  accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEmailRuleJSON  `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEmailRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEmailRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEmailRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                                          `json:"email,required" format:"email"`
	JSON  accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEmailRuleEmailJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEmailRuleEmail]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEmailListRule struct {
	EmailList AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEmailListRuleEmailList `json:"email_list,required"`
	JSON      accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEmailListRuleJSON      `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEmailListRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEmailListRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                                  `json:"id,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEmailListRuleEmailListJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEmailListRuleEmailList]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqDomainRule struct {
	EmailDomain AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqDomainRuleJSON        `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqDomainRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqDomainRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqDomainRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                                 `json:"domain,required"`
	JSON   accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqDomainRuleEmailDomainJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqDomainRuleEmailDomain]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                                   `json:"everyone,required"`
	JSON     accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEveryoneRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEveryoneRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqEveryoneRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

// Matches an IP address block.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqIPRule struct {
	IP   AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqIPRuleIP   `json:"ip,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqIPRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqIPRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqIPRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqIPRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                                    `json:"ip,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqIPRuleIPJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqIPRuleIP]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqIPListRule struct {
	IPList AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqIPListRuleIPList `json:"ip_list,required"`
	JSON   accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqIPListRuleJSON   `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqIPListRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqIPListRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqIPListRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                                            `json:"id,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqIPListRuleIPListJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqIPListRuleIPList]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqCertificateRule struct {
	Certificate interface{}                                                                                      `json:"certificate,required"`
	JSON        accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqCertificateRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqCertificateRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqCertificateRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

// Matches an Access group.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAccessGroupRule struct {
	Group AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAccessGroupRuleGroup `json:"group,required"`
	JSON  accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAccessGroupRuleJSON  `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAccessGroupRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAccessGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                                                `json:"id,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAccessGroupRuleGroupJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAccessGroupRuleGroup]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAzureGroupRule struct {
	AzureAd AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAzureGroupRuleJSON    `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAzureGroupRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAzureGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                                 `json:"connection_id,required"`
	JSON         accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAzureGroupRuleAzureAdJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAzureGroupRuleAzureAd]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqGitHubOrganizationRule struct {
	GitHubOrganization AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqGitHubOrganizationRuleJSON               `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqGitHubOrganizationRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqGitHubOrganizationRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                                    `json:"name,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqGitHubOrganizationRuleGitHubOrganization]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqGsuiteGroupRule struct {
	Gsuite AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqGsuiteGroupRuleJSON   `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqGsuiteGroupRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqGsuiteGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                                 `json:"email,required"`
	JSON  accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqGsuiteGroupRuleGsuite]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqOktaGroupRule struct {
	Okta AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqOktaGroupRuleOkta `json:"okta,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqOktaGroupRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqOktaGroupRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqOktaGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                                             `json:"email,required"`
	JSON  accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqOktaGroupRuleOktaJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqOktaGroupRuleOkta]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqSamlGroupRule struct {
	Saml AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqSamlGroupRuleSaml `json:"saml,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqSamlGroupRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqSamlGroupRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqSamlGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                                             `json:"attribute_value,required"`
	JSON           accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqSamlGroupRuleSamlJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqSamlGroupRuleSaml]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqServiceTokenRule struct {
	ServiceToken AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqServiceTokenRuleJSON         `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqServiceTokenRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqServiceTokenRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                                        `json:"token_id,required"`
	JSON    accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqServiceTokenRuleServiceToken]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                                               `json:"any_valid_service_token,required"`
	JSON                 accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAnyValidServiceTokenRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAnyValidServiceTokenRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAnyValidServiceTokenRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqExternalEvaluationRule struct {
	ExternalEvaluation AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqExternalEvaluationRuleJSON               `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqExternalEvaluationRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqExternalEvaluationRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                                    `json:"keys_url,required"`
	JSON    accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqExternalEvaluationRuleExternalEvaluation]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqCountryRule struct {
	Geo  AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqCountryRuleGeo  `json:"geo,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqCountryRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqCountryRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqCountryRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqCountryRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                                          `json:"country_code,required"`
	JSON        accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqCountryRuleGeoJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqCountryRuleGeo]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAuthenticationMethodRule struct {
	AuthMethod AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAuthenticationMethodRuleJSON       `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAuthenticationMethodRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAuthenticationMethodRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                                              `json:"auth_method,required"`
	JSON       accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAuthenticationMethodRuleAuthMethod]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqDevicePostureRule struct {
	DevicePosture AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqDevicePostureRuleJSON          `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqDevicePostureRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqDevicePostureRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                                          `json:"integration_uid,required"`
	JSON           accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqDevicePostureRuleDevicePosture]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultIsDefaultPajwohLqDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEmailRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEmailListRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqDomainRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEveryoneRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqIPRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqIPListRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqCertificateRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAccessGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAzureGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqGitHubOrganizationRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqGsuiteGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqOktaGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqSamlGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqServiceTokenRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAnyValidServiceTokenRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqExternalEvaluationRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqCountryRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAuthenticationMethodRule]
// or
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqDevicePostureRule].
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire interface {
	implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEmailRule struct {
	Email AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEmailRuleEmail `json:"email,required"`
	JSON  accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEmailRuleJSON  `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEmailRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEmailRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEmailRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                                        `json:"email,required" format:"email"`
	JSON  accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEmailRuleEmailJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEmailRuleEmail]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEmailListRule struct {
	EmailList AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEmailListRuleEmailList `json:"email_list,required"`
	JSON      accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEmailListRuleJSON      `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEmailListRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEmailListRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                                `json:"id,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEmailListRuleEmailListJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEmailListRuleEmailList]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqDomainRule struct {
	EmailDomain AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqDomainRuleJSON        `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqDomainRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqDomainRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqDomainRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                               `json:"domain,required"`
	JSON   accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqDomainRuleEmailDomainJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqDomainRuleEmailDomain]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                                 `json:"everyone,required"`
	JSON     accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEveryoneRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEveryoneRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqEveryoneRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

// Matches an IP address block.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqIPRule struct {
	IP   AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqIPRuleIP   `json:"ip,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqIPRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqIPRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqIPRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqIPRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                                  `json:"ip,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqIPRuleIPJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqIPRuleIP]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqIPListRule struct {
	IPList AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqIPListRuleIPList `json:"ip_list,required"`
	JSON   accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqIPListRuleJSON   `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqIPListRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqIPListRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqIPListRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                                          `json:"id,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqIPListRuleIPListJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqIPListRuleIPList]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqCertificateRule struct {
	Certificate interface{}                                                                                    `json:"certificate,required"`
	JSON        accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqCertificateRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqCertificateRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqCertificateRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

// Matches an Access group.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAccessGroupRule struct {
	Group AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAccessGroupRuleGroup `json:"group,required"`
	JSON  accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAccessGroupRuleJSON  `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAccessGroupRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAccessGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                                              `json:"id,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAccessGroupRuleGroupJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAccessGroupRuleGroup]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAzureGroupRule struct {
	AzureAd AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAzureGroupRuleJSON    `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAzureGroupRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAzureGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                               `json:"connection_id,required"`
	JSON         accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAzureGroupRuleAzureAdJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAzureGroupRuleAzureAd]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqGitHubOrganizationRuleJSON               `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqGitHubOrganizationRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqGitHubOrganizationRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                                  `json:"name,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqGsuiteGroupRule struct {
	Gsuite AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqGsuiteGroupRuleJSON   `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqGsuiteGroupRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqGsuiteGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                               `json:"email,required"`
	JSON  accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqGsuiteGroupRuleGsuite]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqOktaGroupRule struct {
	Okta AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqOktaGroupRuleOkta `json:"okta,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqOktaGroupRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqOktaGroupRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqOktaGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                                           `json:"email,required"`
	JSON  accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqOktaGroupRuleOktaJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqOktaGroupRuleOkta]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqSamlGroupRule struct {
	Saml AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqSamlGroupRuleSaml `json:"saml,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqSamlGroupRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqSamlGroupRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqSamlGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                                           `json:"attribute_value,required"`
	JSON           accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqSamlGroupRuleSamlJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqSamlGroupRuleSaml]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqServiceTokenRule struct {
	ServiceToken AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqServiceTokenRuleJSON         `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqServiceTokenRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqServiceTokenRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                                      `json:"token_id,required"`
	JSON    accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqServiceTokenRuleServiceToken]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                                             `json:"any_valid_service_token,required"`
	JSON                 accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAnyValidServiceTokenRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAnyValidServiceTokenRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAnyValidServiceTokenRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqExternalEvaluationRuleJSON               `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqExternalEvaluationRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqExternalEvaluationRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                                  `json:"keys_url,required"`
	JSON    accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqCountryRule struct {
	Geo  AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqCountryRuleGeo  `json:"geo,required"`
	JSON accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqCountryRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqCountryRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqCountryRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqCountryRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                                        `json:"country_code,required"`
	JSON        accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqCountryRuleGeoJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqCountryRuleGeo]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAuthenticationMethodRule struct {
	AuthMethod AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAuthenticationMethodRuleJSON       `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAuthenticationMethodRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAuthenticationMethodRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                                            `json:"auth_method,required"`
	JSON       accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqDevicePostureRule struct {
	DevicePosture AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqDevicePostureRuleJSON          `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqDevicePostureRule]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqDevicePostureRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                                        `json:"integration_uid,required"`
	JSON           accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqDevicePostureRuleDevicePosture]
type accountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsNewAnAccessGroupResponseResultRequirePajwohLqDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessGroupAccessGroupsNewAnAccessGroupResponseSuccess bool

const (
	AccountAccessGroupAccessGroupsNewAnAccessGroupResponseSuccessTrue AccountAccessGroupAccessGroupsNewAnAccessGroupResponseSuccess = true
)

type AccountAccessGroupAccessGroupsListAccessGroupsResponse struct {
	Errors     []AccountAccessGroupAccessGroupsListAccessGroupsResponseError    `json:"errors"`
	Messages   []AccountAccessGroupAccessGroupsListAccessGroupsResponseMessage  `json:"messages"`
	Result     []AccountAccessGroupAccessGroupsListAccessGroupsResponseResult   `json:"result"`
	ResultInfo AccountAccessGroupAccessGroupsListAccessGroupsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountAccessGroupAccessGroupsListAccessGroupsResponseSuccess `json:"success"`
	JSON    accountAccessGroupAccessGroupsListAccessGroupsResponseJSON    `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseJSON contains the JSON
// metadata for the struct [AccountAccessGroupAccessGroupsListAccessGroupsResponse]
type accountAccessGroupAccessGroupsListAccessGroupsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseError struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    accountAccessGroupAccessGroupsListAccessGroupsResponseErrorJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseError]
type accountAccessGroupAccessGroupsListAccessGroupsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseMessage struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    accountAccessGroupAccessGroupsListAccessGroupsResponseMessageJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseMessageJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseMessage]
type accountAccessGroupAccessGroupsListAccessGroupsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResult struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccountAccessGroupAccessGroupsListAccessGroupsResponseResultInclude `json:"include"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	IsDefault []AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault `json:"is_default"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequire `json:"require"`
	UpdatedAt time.Time                                                             `json:"updated_at" format:"date-time"`
	JSON      accountAccessGroupAccessGroupsListAccessGroupsResponseResultJSON      `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultJSON contains the
// JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResult]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultJSON struct {
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

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEmailRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEmailListRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqDomainRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEveryoneRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqIPRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqIPListRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqCertificateRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAccessGroupRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAzureGroupRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqGitHubOrganizationRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqGsuiteGroupRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqOktaGroupRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqSamlGroupRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqServiceTokenRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAnyValidServiceTokenRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqExternalEvaluationRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqCountryRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAuthenticationMethodRule]
// or
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqDevicePostureRule].
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExclude interface {
	implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEmailRule struct {
	Email AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEmailRuleEmail `json:"email,required"`
	JSON  accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEmailRuleJSON  `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEmailRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEmailRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEmailRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                                        `json:"email,required" format:"email"`
	JSON  accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEmailRuleEmailJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEmailRuleEmail]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEmailListRule struct {
	EmailList AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEmailListRuleEmailList `json:"email_list,required"`
	JSON      accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEmailListRuleJSON      `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEmailListRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEmailListRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                                `json:"id,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEmailListRuleEmailListJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEmailListRuleEmailList]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqDomainRule struct {
	EmailDomain AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqDomainRuleJSON        `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqDomainRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqDomainRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqDomainRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                               `json:"domain,required"`
	JSON   accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqDomainRuleEmailDomainJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqDomainRuleEmailDomain]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                                 `json:"everyone,required"`
	JSON     accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEveryoneRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEveryoneRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqEveryoneRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

// Matches an IP address block.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqIPRule struct {
	IP   AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqIPRuleIP   `json:"ip,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqIPRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqIPRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqIPRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqIPRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                                  `json:"ip,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqIPRuleIPJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqIPRuleIP]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqIPListRule struct {
	IPList AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqIPListRuleIPList `json:"ip_list,required"`
	JSON   accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqIPListRuleJSON   `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqIPListRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqIPListRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqIPListRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                                          `json:"id,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqIPListRuleIPListJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqIPListRuleIPList]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqCertificateRule struct {
	Certificate interface{}                                                                                    `json:"certificate,required"`
	JSON        accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqCertificateRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqCertificateRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqCertificateRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

// Matches an Access group.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAccessGroupRule struct {
	Group AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAccessGroupRuleGroup `json:"group,required"`
	JSON  accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAccessGroupRuleJSON  `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAccessGroupRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAccessGroupRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                                              `json:"id,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAccessGroupRuleGroupJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAccessGroupRuleGroup]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAzureGroupRule struct {
	AzureAd AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAzureGroupRuleJSON    `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAzureGroupRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAzureGroupRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                               `json:"connection_id,required"`
	JSON         accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAzureGroupRuleAzureAdJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAzureGroupRuleAzureAd]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqGitHubOrganizationRuleJSON               `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqGitHubOrganizationRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqGitHubOrganizationRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                                  `json:"name,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqGsuiteGroupRule struct {
	Gsuite AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqGsuiteGroupRuleJSON   `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqGsuiteGroupRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqGsuiteGroupRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                               `json:"email,required"`
	JSON  accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqGsuiteGroupRuleGsuite]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqOktaGroupRule struct {
	Okta AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqOktaGroupRuleOkta `json:"okta,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqOktaGroupRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqOktaGroupRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqOktaGroupRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                                           `json:"email,required"`
	JSON  accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqOktaGroupRuleOktaJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqOktaGroupRuleOkta]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqSamlGroupRule struct {
	Saml AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqSamlGroupRuleSaml `json:"saml,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqSamlGroupRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqSamlGroupRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqSamlGroupRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                                           `json:"attribute_value,required"`
	JSON           accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqSamlGroupRuleSamlJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqSamlGroupRuleSaml]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqServiceTokenRule struct {
	ServiceToken AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqServiceTokenRuleJSON         `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqServiceTokenRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqServiceTokenRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                                      `json:"token_id,required"`
	JSON    accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqServiceTokenRuleServiceToken]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                                             `json:"any_valid_service_token,required"`
	JSON                 accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAnyValidServiceTokenRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAnyValidServiceTokenRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAnyValidServiceTokenRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqExternalEvaluationRuleJSON               `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqExternalEvaluationRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqExternalEvaluationRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                                  `json:"keys_url,required"`
	JSON    accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqCountryRule struct {
	Geo  AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqCountryRuleGeo  `json:"geo,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqCountryRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqCountryRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqCountryRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqCountryRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                                        `json:"country_code,required"`
	JSON        accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqCountryRuleGeoJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqCountryRuleGeo]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAuthenticationMethodRule struct {
	AuthMethod AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAuthenticationMethodRuleJSON       `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAuthenticationMethodRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAuthenticationMethodRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                                            `json:"auth_method,required"`
	JSON       accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqDevicePostureRule struct {
	DevicePosture AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqDevicePostureRuleJSON          `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqDevicePostureRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqDevicePostureRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultExclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                                        `json:"integration_uid,required"`
	JSON           accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqDevicePostureRuleDevicePosture]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultExcludePajwohLqDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEmailRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEmailListRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqDomainRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEveryoneRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqIPRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqIPListRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqCertificateRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAccessGroupRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAzureGroupRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqGitHubOrganizationRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqGsuiteGroupRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqOktaGroupRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqSamlGroupRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqServiceTokenRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAnyValidServiceTokenRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqExternalEvaluationRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqCountryRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAuthenticationMethodRule]
// or
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqDevicePostureRule].
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultInclude interface {
	implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountAccessGroupAccessGroupsListAccessGroupsResponseResultInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEmailRule struct {
	Email AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEmailRuleEmail `json:"email,required"`
	JSON  accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEmailRuleJSON  `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEmailRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEmailRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEmailRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                                        `json:"email,required" format:"email"`
	JSON  accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEmailRuleEmailJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEmailRuleEmail]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEmailListRule struct {
	EmailList AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEmailListRuleEmailList `json:"email_list,required"`
	JSON      accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEmailListRuleJSON      `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEmailListRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEmailListRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                                `json:"id,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEmailListRuleEmailListJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEmailListRuleEmailList]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqDomainRule struct {
	EmailDomain AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqDomainRuleJSON        `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqDomainRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqDomainRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqDomainRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                               `json:"domain,required"`
	JSON   accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqDomainRuleEmailDomainJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqDomainRuleEmailDomain]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                                 `json:"everyone,required"`
	JSON     accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEveryoneRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEveryoneRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqEveryoneRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

// Matches an IP address block.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqIPRule struct {
	IP   AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqIPRuleIP   `json:"ip,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqIPRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqIPRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqIPRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqIPRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                                  `json:"ip,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqIPRuleIPJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqIPRuleIP]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqIPListRule struct {
	IPList AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqIPListRuleIPList `json:"ip_list,required"`
	JSON   accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqIPListRuleJSON   `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqIPListRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqIPListRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqIPListRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                                          `json:"id,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqIPListRuleIPListJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqIPListRuleIPList]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqCertificateRule struct {
	Certificate interface{}                                                                                    `json:"certificate,required"`
	JSON        accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqCertificateRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqCertificateRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqCertificateRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

// Matches an Access group.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAccessGroupRule struct {
	Group AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAccessGroupRuleGroup `json:"group,required"`
	JSON  accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAccessGroupRuleJSON  `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAccessGroupRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAccessGroupRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                                              `json:"id,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAccessGroupRuleGroupJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAccessGroupRuleGroup]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAzureGroupRule struct {
	AzureAd AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAzureGroupRuleJSON    `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAzureGroupRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAzureGroupRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                               `json:"connection_id,required"`
	JSON         accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAzureGroupRuleAzureAdJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAzureGroupRuleAzureAd]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqGitHubOrganizationRuleJSON               `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqGitHubOrganizationRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqGitHubOrganizationRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                                  `json:"name,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqGsuiteGroupRule struct {
	Gsuite AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqGsuiteGroupRuleJSON   `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqGsuiteGroupRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqGsuiteGroupRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                               `json:"email,required"`
	JSON  accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqGsuiteGroupRuleGsuite]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqOktaGroupRule struct {
	Okta AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqOktaGroupRuleOkta `json:"okta,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqOktaGroupRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqOktaGroupRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqOktaGroupRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                                           `json:"email,required"`
	JSON  accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqOktaGroupRuleOktaJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqOktaGroupRuleOkta]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqSamlGroupRule struct {
	Saml AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqSamlGroupRuleSaml `json:"saml,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqSamlGroupRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqSamlGroupRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqSamlGroupRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                                           `json:"attribute_value,required"`
	JSON           accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqSamlGroupRuleSamlJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqSamlGroupRuleSaml]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqServiceTokenRule struct {
	ServiceToken AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqServiceTokenRuleJSON         `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqServiceTokenRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqServiceTokenRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                                      `json:"token_id,required"`
	JSON    accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqServiceTokenRuleServiceToken]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                                             `json:"any_valid_service_token,required"`
	JSON                 accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAnyValidServiceTokenRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAnyValidServiceTokenRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAnyValidServiceTokenRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqExternalEvaluationRuleJSON               `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqExternalEvaluationRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqExternalEvaluationRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                                  `json:"keys_url,required"`
	JSON    accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqCountryRule struct {
	Geo  AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqCountryRuleGeo  `json:"geo,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqCountryRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqCountryRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqCountryRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqCountryRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                                        `json:"country_code,required"`
	JSON        accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqCountryRuleGeoJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqCountryRuleGeo]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAuthenticationMethodRule struct {
	AuthMethod AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAuthenticationMethodRuleJSON       `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAuthenticationMethodRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAuthenticationMethodRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                                            `json:"auth_method,required"`
	JSON       accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqDevicePostureRule struct {
	DevicePosture AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqDevicePostureRuleJSON          `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqDevicePostureRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqDevicePostureRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultInclude() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                                        `json:"integration_uid,required"`
	JSON           accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqDevicePostureRuleDevicePosture]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIncludePajwohLqDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEmailRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEmailListRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqDomainRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEveryoneRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqIPRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqIPListRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqCertificateRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAccessGroupRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAzureGroupRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqGitHubOrganizationRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqGsuiteGroupRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqOktaGroupRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqSamlGroupRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqServiceTokenRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAnyValidServiceTokenRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqExternalEvaluationRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqCountryRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAuthenticationMethodRule]
// or
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqDevicePostureRule].
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault interface {
	implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault)(nil)).Elem(), "")
}

// Matches a specific email.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEmailRule struct {
	Email AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEmailRuleEmail `json:"email,required"`
	JSON  accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEmailRuleJSON  `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEmailRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEmailRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEmailRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                                          `json:"email,required" format:"email"`
	JSON  accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEmailRuleEmailJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEmailRuleEmail]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEmailListRule struct {
	EmailList AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEmailListRuleEmailList `json:"email_list,required"`
	JSON      accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEmailListRuleJSON      `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEmailListRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEmailListRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                                  `json:"id,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEmailListRuleEmailListJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEmailListRuleEmailList]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqDomainRule struct {
	EmailDomain AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqDomainRuleJSON        `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqDomainRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqDomainRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqDomainRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                                 `json:"domain,required"`
	JSON   accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqDomainRuleEmailDomainJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqDomainRuleEmailDomain]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                                   `json:"everyone,required"`
	JSON     accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEveryoneRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEveryoneRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqEveryoneRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

// Matches an IP address block.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqIPRule struct {
	IP   AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqIPRuleIP   `json:"ip,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqIPRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqIPRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqIPRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqIPRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                                    `json:"ip,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqIPRuleIPJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqIPRuleIP]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqIPListRule struct {
	IPList AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqIPListRuleIPList `json:"ip_list,required"`
	JSON   accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqIPListRuleJSON   `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqIPListRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqIPListRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqIPListRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                                            `json:"id,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqIPListRuleIPListJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqIPListRuleIPList]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqCertificateRule struct {
	Certificate interface{}                                                                                      `json:"certificate,required"`
	JSON        accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqCertificateRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqCertificateRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqCertificateRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

// Matches an Access group.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAccessGroupRule struct {
	Group AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAccessGroupRuleGroup `json:"group,required"`
	JSON  accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAccessGroupRuleJSON  `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAccessGroupRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAccessGroupRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                                                `json:"id,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAccessGroupRuleGroupJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAccessGroupRuleGroup]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAzureGroupRule struct {
	AzureAd AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAzureGroupRuleJSON    `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAzureGroupRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAzureGroupRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                                 `json:"connection_id,required"`
	JSON         accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAzureGroupRuleAzureAdJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAzureGroupRuleAzureAd]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqGitHubOrganizationRule struct {
	GitHubOrganization AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqGitHubOrganizationRuleJSON               `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqGitHubOrganizationRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqGitHubOrganizationRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                                    `json:"name,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqGitHubOrganizationRuleGitHubOrganization]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqGsuiteGroupRule struct {
	Gsuite AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqGsuiteGroupRuleJSON   `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqGsuiteGroupRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqGsuiteGroupRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                                 `json:"email,required"`
	JSON  accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqGsuiteGroupRuleGsuite]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqOktaGroupRule struct {
	Okta AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqOktaGroupRuleOkta `json:"okta,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqOktaGroupRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqOktaGroupRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqOktaGroupRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                                             `json:"email,required"`
	JSON  accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqOktaGroupRuleOktaJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqOktaGroupRuleOkta]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqSamlGroupRule struct {
	Saml AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqSamlGroupRuleSaml `json:"saml,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqSamlGroupRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqSamlGroupRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqSamlGroupRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                                             `json:"attribute_value,required"`
	JSON           accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqSamlGroupRuleSamlJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqSamlGroupRuleSaml]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqServiceTokenRule struct {
	ServiceToken AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqServiceTokenRuleJSON         `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqServiceTokenRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqServiceTokenRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                                        `json:"token_id,required"`
	JSON    accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqServiceTokenRuleServiceToken]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                                               `json:"any_valid_service_token,required"`
	JSON                 accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAnyValidServiceTokenRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAnyValidServiceTokenRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAnyValidServiceTokenRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqExternalEvaluationRule struct {
	ExternalEvaluation AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqExternalEvaluationRuleJSON               `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqExternalEvaluationRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqExternalEvaluationRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                                    `json:"keys_url,required"`
	JSON    accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqExternalEvaluationRuleExternalEvaluation]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqCountryRule struct {
	Geo  AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqCountryRuleGeo  `json:"geo,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqCountryRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqCountryRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqCountryRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqCountryRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                                          `json:"country_code,required"`
	JSON        accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqCountryRuleGeoJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqCountryRuleGeo]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAuthenticationMethodRule struct {
	AuthMethod AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAuthenticationMethodRuleJSON       `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAuthenticationMethodRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAuthenticationMethodRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                                              `json:"auth_method,required"`
	JSON       accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAuthenticationMethodRuleAuthMethod]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqDevicePostureRule struct {
	DevicePosture AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqDevicePostureRuleJSON          `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqDevicePostureRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqDevicePostureRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefault() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                                          `json:"integration_uid,required"`
	JSON           accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqDevicePostureRuleDevicePosture]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultIsDefaultPajwohLqDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEmailRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEmailListRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqDomainRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEveryoneRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqIPRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqIPListRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqCertificateRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAccessGroupRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAzureGroupRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqGitHubOrganizationRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqGsuiteGroupRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqOktaGroupRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqSamlGroupRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqServiceTokenRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAnyValidServiceTokenRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqExternalEvaluationRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqCountryRule],
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAuthenticationMethodRule]
// or
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqDevicePostureRule].
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequire interface {
	implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEmailRule struct {
	Email AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEmailRuleEmail `json:"email,required"`
	JSON  accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEmailRuleJSON  `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEmailRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEmailRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEmailRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email string                                                                                        `json:"email,required" format:"email"`
	JSON  accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEmailRuleEmailJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEmailRuleEmailJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEmailRuleEmail]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEmailListRule struct {
	EmailList AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEmailListRuleEmailList `json:"email_list,required"`
	JSON      accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEmailListRuleJSON      `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEmailListRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEmailListRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEmailListRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                                                `json:"id,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEmailListRuleEmailListJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEmailListRuleEmailList]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqDomainRule struct {
	EmailDomain AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqDomainRuleJSON        `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqDomainRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqDomainRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqDomainRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                                               `json:"domain,required"`
	JSON   accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqDomainRuleEmailDomainJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqDomainRuleEmailDomain]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                                                 `json:"everyone,required"`
	JSON     accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEveryoneRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEveryoneRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEveryoneRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqEveryoneRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

// Matches an IP address block.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqIPRule struct {
	IP   AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqIPRuleIP   `json:"ip,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqIPRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqIPRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqIPRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqIPRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                                                  `json:"ip,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqIPRuleIPJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqIPRuleIPJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqIPRuleIP]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqIPListRule struct {
	IPList AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqIPListRuleIPList `json:"ip_list,required"`
	JSON   accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqIPListRuleJSON   `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqIPListRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqIPListRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqIPListRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                                          `json:"id,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqIPListRuleIPListJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqIPListRuleIPListJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqIPListRuleIPList]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqCertificateRule struct {
	Certificate interface{}                                                                                    `json:"certificate,required"`
	JSON        accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqCertificateRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqCertificateRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqCertificateRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqCertificateRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

// Matches an Access group.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAccessGroupRule struct {
	Group AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAccessGroupRuleGroup `json:"group,required"`
	JSON  accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAccessGroupRuleJSON  `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAccessGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAccessGroupRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAccessGroupRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                                              `json:"id,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAccessGroupRuleGroupJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAccessGroupRuleGroup]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAzureGroupRule struct {
	AzureAd AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAzureGroupRuleJSON    `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAzureGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAzureGroupRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAzureGroupRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                                               `json:"connection_id,required"`
	JSON         accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAzureGroupRuleAzureAdJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAzureGroupRuleAzureAd]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqGitHubOrganizationRuleJSON               `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqGitHubOrganizationRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqGitHubOrganizationRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                                                  `json:"name,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqGsuiteGroupRule struct {
	Gsuite AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqGsuiteGroupRuleJSON   `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqGsuiteGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqGsuiteGroupRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqGsuiteGroupRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                                               `json:"email,required"`
	JSON  accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqGsuiteGroupRuleGsuite]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqOktaGroupRule struct {
	Okta AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqOktaGroupRuleOkta `json:"okta,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqOktaGroupRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqOktaGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqOktaGroupRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqOktaGroupRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                                           `json:"email,required"`
	JSON  accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqOktaGroupRuleOktaJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqOktaGroupRuleOktaJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqOktaGroupRuleOkta]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqSamlGroupRule struct {
	Saml AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqSamlGroupRuleSaml `json:"saml,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqSamlGroupRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqSamlGroupRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqSamlGroupRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqSamlGroupRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                                           `json:"attribute_value,required"`
	JSON           accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqSamlGroupRuleSamlJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqSamlGroupRuleSamlJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqSamlGroupRuleSaml]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqServiceTokenRule struct {
	ServiceToken AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqServiceTokenRuleJSON         `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqServiceTokenRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqServiceTokenRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                                                      `json:"token_id,required"`
	JSON    accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqServiceTokenRuleServiceToken]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                                             `json:"any_valid_service_token,required"`
	JSON                 accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAnyValidServiceTokenRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAnyValidServiceTokenRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAnyValidServiceTokenRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqExternalEvaluationRuleJSON               `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqExternalEvaluationRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqExternalEvaluationRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                                                  `json:"keys_url,required"`
	JSON    accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqCountryRule struct {
	Geo  AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqCountryRuleGeo  `json:"geo,required"`
	JSON accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqCountryRuleJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqCountryRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqCountryRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqCountryRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                                                        `json:"country_code,required"`
	JSON        accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqCountryRuleGeoJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqCountryRuleGeoJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqCountryRuleGeo]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAuthenticationMethodRule struct {
	AuthMethod AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAuthenticationMethodRuleJSON       `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAuthenticationMethodRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAuthenticationMethodRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                                            `json:"auth_method,required"`
	JSON       accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqDevicePostureRule struct {
	DevicePosture AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqDevicePostureRuleJSON          `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqDevicePostureRuleJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqDevicePostureRule]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqDevicePostureRule) implementsAccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequire() {
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                                                        `json:"integration_uid,required"`
	JSON           accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqDevicePostureRuleDevicePosture]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultRequirePajwohLqDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessGroupAccessGroupsListAccessGroupsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                              `json:"total_count"`
	JSON       accountAccessGroupAccessGroupsListAccessGroupsResponseResultInfoJSON `json:"-"`
}

// accountAccessGroupAccessGroupsListAccessGroupsResponseResultInfoJSON contains
// the JSON metadata for the struct
// [AccountAccessGroupAccessGroupsListAccessGroupsResponseResultInfo]
type accountAccessGroupAccessGroupsListAccessGroupsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessGroupAccessGroupsListAccessGroupsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessGroupAccessGroupsListAccessGroupsResponseSuccess bool

const (
	AccountAccessGroupAccessGroupsListAccessGroupsResponseSuccessTrue AccountAccessGroupAccessGroupsListAccessGroupsResponseSuccess = true
)

type AccountAccessGroupUpdateParams struct {
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccountAccessGroupUpdateParamsInclude] `json:"include,required"`
	// The name of the Access group.
	Name param.Field[string] `json:"name,required"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccountAccessGroupUpdateParamsExclude] `json:"exclude"`
	// Whether this is the default group
	IsDefault param.Field[bool] `json:"is_default"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccountAccessGroupUpdateParamsRequire] `json:"require"`
}

func (r AccountAccessGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [AccountAccessGroupUpdateParamsIncludePajwohLqEmailRule],
// [AccountAccessGroupUpdateParamsIncludePajwohLqEmailListRule],
// [AccountAccessGroupUpdateParamsIncludePajwohLqDomainRule],
// [AccountAccessGroupUpdateParamsIncludePajwohLqEveryoneRule],
// [AccountAccessGroupUpdateParamsIncludePajwohLqIPRule],
// [AccountAccessGroupUpdateParamsIncludePajwohLqIPListRule],
// [AccountAccessGroupUpdateParamsIncludePajwohLqCertificateRule],
// [AccountAccessGroupUpdateParamsIncludePajwohLqAccessGroupRule],
// [AccountAccessGroupUpdateParamsIncludePajwohLqAzureGroupRule],
// [AccountAccessGroupUpdateParamsIncludePajwohLqGitHubOrganizationRule],
// [AccountAccessGroupUpdateParamsIncludePajwohLqGsuiteGroupRule],
// [AccountAccessGroupUpdateParamsIncludePajwohLqOktaGroupRule],
// [AccountAccessGroupUpdateParamsIncludePajwohLqSamlGroupRule],
// [AccountAccessGroupUpdateParamsIncludePajwohLqServiceTokenRule],
// [AccountAccessGroupUpdateParamsIncludePajwohLqAnyValidServiceTokenRule],
// [AccountAccessGroupUpdateParamsIncludePajwohLqExternalEvaluationRule],
// [AccountAccessGroupUpdateParamsIncludePajwohLqCountryRule],
// [AccountAccessGroupUpdateParamsIncludePajwohLqAuthenticationMethodRule],
// [AccountAccessGroupUpdateParamsIncludePajwohLqDevicePostureRule].
type AccountAccessGroupUpdateParamsInclude interface {
	implementsAccountAccessGroupUpdateParamsInclude()
}

// Matches a specific email.
type AccountAccessGroupUpdateParamsIncludePajwohLqEmailRule struct {
	Email param.Field[AccountAccessGroupUpdateParamsIncludePajwohLqEmailRuleEmail] `json:"email,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqEmailRule) implementsAccountAccessGroupUpdateParamsInclude() {
}

type AccountAccessGroupUpdateParamsIncludePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccountAccessGroupUpdateParamsIncludePajwohLqEmailListRule struct {
	EmailList param.Field[AccountAccessGroupUpdateParamsIncludePajwohLqEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqEmailListRule) implementsAccountAccessGroupUpdateParamsInclude() {
}

type AccountAccessGroupUpdateParamsIncludePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccountAccessGroupUpdateParamsIncludePajwohLqDomainRule struct {
	EmailDomain param.Field[AccountAccessGroupUpdateParamsIncludePajwohLqDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqDomainRule) implementsAccountAccessGroupUpdateParamsInclude() {
}

type AccountAccessGroupUpdateParamsIncludePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccountAccessGroupUpdateParamsIncludePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqEveryoneRule) implementsAccountAccessGroupUpdateParamsInclude() {
}

// Matches an IP address block.
type AccountAccessGroupUpdateParamsIncludePajwohLqIPRule struct {
	IP param.Field[AccountAccessGroupUpdateParamsIncludePajwohLqIPRuleIP] `json:"ip,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqIPRule) implementsAccountAccessGroupUpdateParamsInclude() {
}

type AccountAccessGroupUpdateParamsIncludePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccountAccessGroupUpdateParamsIncludePajwohLqIPListRule struct {
	IPList param.Field[AccountAccessGroupUpdateParamsIncludePajwohLqIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqIPListRule) implementsAccountAccessGroupUpdateParamsInclude() {
}

type AccountAccessGroupUpdateParamsIncludePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccountAccessGroupUpdateParamsIncludePajwohLqCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqCertificateRule) implementsAccountAccessGroupUpdateParamsInclude() {
}

// Matches an Access group.
type AccountAccessGroupUpdateParamsIncludePajwohLqAccessGroupRule struct {
	Group param.Field[AccountAccessGroupUpdateParamsIncludePajwohLqAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqAccessGroupRule) implementsAccountAccessGroupUpdateParamsInclude() {
}

type AccountAccessGroupUpdateParamsIncludePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccountAccessGroupUpdateParamsIncludePajwohLqAzureGroupRule struct {
	AzureAd param.Field[AccountAccessGroupUpdateParamsIncludePajwohLqAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqAzureGroupRule) implementsAccountAccessGroupUpdateParamsInclude() {
}

type AccountAccessGroupUpdateParamsIncludePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccountAccessGroupUpdateParamsIncludePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccountAccessGroupUpdateParamsIncludePajwohLqGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqGitHubOrganizationRule) implementsAccountAccessGroupUpdateParamsInclude() {
}

type AccountAccessGroupUpdateParamsIncludePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccountAccessGroupUpdateParamsIncludePajwohLqGsuiteGroupRule struct {
	Gsuite param.Field[AccountAccessGroupUpdateParamsIncludePajwohLqGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqGsuiteGroupRule) implementsAccountAccessGroupUpdateParamsInclude() {
}

type AccountAccessGroupUpdateParamsIncludePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccountAccessGroupUpdateParamsIncludePajwohLqOktaGroupRule struct {
	Okta param.Field[AccountAccessGroupUpdateParamsIncludePajwohLqOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqOktaGroupRule) implementsAccountAccessGroupUpdateParamsInclude() {
}

type AccountAccessGroupUpdateParamsIncludePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccountAccessGroupUpdateParamsIncludePajwohLqSamlGroupRule struct {
	Saml param.Field[AccountAccessGroupUpdateParamsIncludePajwohLqSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqSamlGroupRule) implementsAccountAccessGroupUpdateParamsInclude() {
}

type AccountAccessGroupUpdateParamsIncludePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccountAccessGroupUpdateParamsIncludePajwohLqServiceTokenRule struct {
	ServiceToken param.Field[AccountAccessGroupUpdateParamsIncludePajwohLqServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqServiceTokenRule) implementsAccountAccessGroupUpdateParamsInclude() {
}

type AccountAccessGroupUpdateParamsIncludePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccountAccessGroupUpdateParamsIncludePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqAnyValidServiceTokenRule) implementsAccountAccessGroupUpdateParamsInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccountAccessGroupUpdateParamsIncludePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccountAccessGroupUpdateParamsIncludePajwohLqExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqExternalEvaluationRule) implementsAccountAccessGroupUpdateParamsInclude() {
}

type AccountAccessGroupUpdateParamsIncludePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccountAccessGroupUpdateParamsIncludePajwohLqCountryRule struct {
	Geo param.Field[AccountAccessGroupUpdateParamsIncludePajwohLqCountryRuleGeo] `json:"geo,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqCountryRule) implementsAccountAccessGroupUpdateParamsInclude() {
}

type AccountAccessGroupUpdateParamsIncludePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccountAccessGroupUpdateParamsIncludePajwohLqAuthenticationMethodRule struct {
	AuthMethod param.Field[AccountAccessGroupUpdateParamsIncludePajwohLqAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqAuthenticationMethodRule) implementsAccountAccessGroupUpdateParamsInclude() {
}

type AccountAccessGroupUpdateParamsIncludePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccountAccessGroupUpdateParamsIncludePajwohLqDevicePostureRule struct {
	DevicePosture param.Field[AccountAccessGroupUpdateParamsIncludePajwohLqDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqDevicePostureRule) implementsAccountAccessGroupUpdateParamsInclude() {
}

type AccountAccessGroupUpdateParamsIncludePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccountAccessGroupUpdateParamsIncludePajwohLqDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [AccountAccessGroupUpdateParamsExcludePajwohLqEmailRule],
// [AccountAccessGroupUpdateParamsExcludePajwohLqEmailListRule],
// [AccountAccessGroupUpdateParamsExcludePajwohLqDomainRule],
// [AccountAccessGroupUpdateParamsExcludePajwohLqEveryoneRule],
// [AccountAccessGroupUpdateParamsExcludePajwohLqIPRule],
// [AccountAccessGroupUpdateParamsExcludePajwohLqIPListRule],
// [AccountAccessGroupUpdateParamsExcludePajwohLqCertificateRule],
// [AccountAccessGroupUpdateParamsExcludePajwohLqAccessGroupRule],
// [AccountAccessGroupUpdateParamsExcludePajwohLqAzureGroupRule],
// [AccountAccessGroupUpdateParamsExcludePajwohLqGitHubOrganizationRule],
// [AccountAccessGroupUpdateParamsExcludePajwohLqGsuiteGroupRule],
// [AccountAccessGroupUpdateParamsExcludePajwohLqOktaGroupRule],
// [AccountAccessGroupUpdateParamsExcludePajwohLqSamlGroupRule],
// [AccountAccessGroupUpdateParamsExcludePajwohLqServiceTokenRule],
// [AccountAccessGroupUpdateParamsExcludePajwohLqAnyValidServiceTokenRule],
// [AccountAccessGroupUpdateParamsExcludePajwohLqExternalEvaluationRule],
// [AccountAccessGroupUpdateParamsExcludePajwohLqCountryRule],
// [AccountAccessGroupUpdateParamsExcludePajwohLqAuthenticationMethodRule],
// [AccountAccessGroupUpdateParamsExcludePajwohLqDevicePostureRule].
type AccountAccessGroupUpdateParamsExclude interface {
	implementsAccountAccessGroupUpdateParamsExclude()
}

// Matches a specific email.
type AccountAccessGroupUpdateParamsExcludePajwohLqEmailRule struct {
	Email param.Field[AccountAccessGroupUpdateParamsExcludePajwohLqEmailRuleEmail] `json:"email,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqEmailRule) implementsAccountAccessGroupUpdateParamsExclude() {
}

type AccountAccessGroupUpdateParamsExcludePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccountAccessGroupUpdateParamsExcludePajwohLqEmailListRule struct {
	EmailList param.Field[AccountAccessGroupUpdateParamsExcludePajwohLqEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqEmailListRule) implementsAccountAccessGroupUpdateParamsExclude() {
}

type AccountAccessGroupUpdateParamsExcludePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccountAccessGroupUpdateParamsExcludePajwohLqDomainRule struct {
	EmailDomain param.Field[AccountAccessGroupUpdateParamsExcludePajwohLqDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqDomainRule) implementsAccountAccessGroupUpdateParamsExclude() {
}

type AccountAccessGroupUpdateParamsExcludePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccountAccessGroupUpdateParamsExcludePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqEveryoneRule) implementsAccountAccessGroupUpdateParamsExclude() {
}

// Matches an IP address block.
type AccountAccessGroupUpdateParamsExcludePajwohLqIPRule struct {
	IP param.Field[AccountAccessGroupUpdateParamsExcludePajwohLqIPRuleIP] `json:"ip,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqIPRule) implementsAccountAccessGroupUpdateParamsExclude() {
}

type AccountAccessGroupUpdateParamsExcludePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccountAccessGroupUpdateParamsExcludePajwohLqIPListRule struct {
	IPList param.Field[AccountAccessGroupUpdateParamsExcludePajwohLqIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqIPListRule) implementsAccountAccessGroupUpdateParamsExclude() {
}

type AccountAccessGroupUpdateParamsExcludePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccountAccessGroupUpdateParamsExcludePajwohLqCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqCertificateRule) implementsAccountAccessGroupUpdateParamsExclude() {
}

// Matches an Access group.
type AccountAccessGroupUpdateParamsExcludePajwohLqAccessGroupRule struct {
	Group param.Field[AccountAccessGroupUpdateParamsExcludePajwohLqAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqAccessGroupRule) implementsAccountAccessGroupUpdateParamsExclude() {
}

type AccountAccessGroupUpdateParamsExcludePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccountAccessGroupUpdateParamsExcludePajwohLqAzureGroupRule struct {
	AzureAd param.Field[AccountAccessGroupUpdateParamsExcludePajwohLqAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqAzureGroupRule) implementsAccountAccessGroupUpdateParamsExclude() {
}

type AccountAccessGroupUpdateParamsExcludePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccountAccessGroupUpdateParamsExcludePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccountAccessGroupUpdateParamsExcludePajwohLqGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqGitHubOrganizationRule) implementsAccountAccessGroupUpdateParamsExclude() {
}

type AccountAccessGroupUpdateParamsExcludePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccountAccessGroupUpdateParamsExcludePajwohLqGsuiteGroupRule struct {
	Gsuite param.Field[AccountAccessGroupUpdateParamsExcludePajwohLqGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqGsuiteGroupRule) implementsAccountAccessGroupUpdateParamsExclude() {
}

type AccountAccessGroupUpdateParamsExcludePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccountAccessGroupUpdateParamsExcludePajwohLqOktaGroupRule struct {
	Okta param.Field[AccountAccessGroupUpdateParamsExcludePajwohLqOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqOktaGroupRule) implementsAccountAccessGroupUpdateParamsExclude() {
}

type AccountAccessGroupUpdateParamsExcludePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccountAccessGroupUpdateParamsExcludePajwohLqSamlGroupRule struct {
	Saml param.Field[AccountAccessGroupUpdateParamsExcludePajwohLqSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqSamlGroupRule) implementsAccountAccessGroupUpdateParamsExclude() {
}

type AccountAccessGroupUpdateParamsExcludePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccountAccessGroupUpdateParamsExcludePajwohLqServiceTokenRule struct {
	ServiceToken param.Field[AccountAccessGroupUpdateParamsExcludePajwohLqServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqServiceTokenRule) implementsAccountAccessGroupUpdateParamsExclude() {
}

type AccountAccessGroupUpdateParamsExcludePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccountAccessGroupUpdateParamsExcludePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqAnyValidServiceTokenRule) implementsAccountAccessGroupUpdateParamsExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccountAccessGroupUpdateParamsExcludePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccountAccessGroupUpdateParamsExcludePajwohLqExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqExternalEvaluationRule) implementsAccountAccessGroupUpdateParamsExclude() {
}

type AccountAccessGroupUpdateParamsExcludePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccountAccessGroupUpdateParamsExcludePajwohLqCountryRule struct {
	Geo param.Field[AccountAccessGroupUpdateParamsExcludePajwohLqCountryRuleGeo] `json:"geo,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqCountryRule) implementsAccountAccessGroupUpdateParamsExclude() {
}

type AccountAccessGroupUpdateParamsExcludePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccountAccessGroupUpdateParamsExcludePajwohLqAuthenticationMethodRule struct {
	AuthMethod param.Field[AccountAccessGroupUpdateParamsExcludePajwohLqAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqAuthenticationMethodRule) implementsAccountAccessGroupUpdateParamsExclude() {
}

type AccountAccessGroupUpdateParamsExcludePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccountAccessGroupUpdateParamsExcludePajwohLqDevicePostureRule struct {
	DevicePosture param.Field[AccountAccessGroupUpdateParamsExcludePajwohLqDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqDevicePostureRule) implementsAccountAccessGroupUpdateParamsExclude() {
}

type AccountAccessGroupUpdateParamsExcludePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccountAccessGroupUpdateParamsExcludePajwohLqDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [AccountAccessGroupUpdateParamsRequirePajwohLqEmailRule],
// [AccountAccessGroupUpdateParamsRequirePajwohLqEmailListRule],
// [AccountAccessGroupUpdateParamsRequirePajwohLqDomainRule],
// [AccountAccessGroupUpdateParamsRequirePajwohLqEveryoneRule],
// [AccountAccessGroupUpdateParamsRequirePajwohLqIPRule],
// [AccountAccessGroupUpdateParamsRequirePajwohLqIPListRule],
// [AccountAccessGroupUpdateParamsRequirePajwohLqCertificateRule],
// [AccountAccessGroupUpdateParamsRequirePajwohLqAccessGroupRule],
// [AccountAccessGroupUpdateParamsRequirePajwohLqAzureGroupRule],
// [AccountAccessGroupUpdateParamsRequirePajwohLqGitHubOrganizationRule],
// [AccountAccessGroupUpdateParamsRequirePajwohLqGsuiteGroupRule],
// [AccountAccessGroupUpdateParamsRequirePajwohLqOktaGroupRule],
// [AccountAccessGroupUpdateParamsRequirePajwohLqSamlGroupRule],
// [AccountAccessGroupUpdateParamsRequirePajwohLqServiceTokenRule],
// [AccountAccessGroupUpdateParamsRequirePajwohLqAnyValidServiceTokenRule],
// [AccountAccessGroupUpdateParamsRequirePajwohLqExternalEvaluationRule],
// [AccountAccessGroupUpdateParamsRequirePajwohLqCountryRule],
// [AccountAccessGroupUpdateParamsRequirePajwohLqAuthenticationMethodRule],
// [AccountAccessGroupUpdateParamsRequirePajwohLqDevicePostureRule].
type AccountAccessGroupUpdateParamsRequire interface {
	implementsAccountAccessGroupUpdateParamsRequire()
}

// Matches a specific email.
type AccountAccessGroupUpdateParamsRequirePajwohLqEmailRule struct {
	Email param.Field[AccountAccessGroupUpdateParamsRequirePajwohLqEmailRuleEmail] `json:"email,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqEmailRule) implementsAccountAccessGroupUpdateParamsRequire() {
}

type AccountAccessGroupUpdateParamsRequirePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccountAccessGroupUpdateParamsRequirePajwohLqEmailListRule struct {
	EmailList param.Field[AccountAccessGroupUpdateParamsRequirePajwohLqEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqEmailListRule) implementsAccountAccessGroupUpdateParamsRequire() {
}

type AccountAccessGroupUpdateParamsRequirePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccountAccessGroupUpdateParamsRequirePajwohLqDomainRule struct {
	EmailDomain param.Field[AccountAccessGroupUpdateParamsRequirePajwohLqDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqDomainRule) implementsAccountAccessGroupUpdateParamsRequire() {
}

type AccountAccessGroupUpdateParamsRequirePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccountAccessGroupUpdateParamsRequirePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqEveryoneRule) implementsAccountAccessGroupUpdateParamsRequire() {
}

// Matches an IP address block.
type AccountAccessGroupUpdateParamsRequirePajwohLqIPRule struct {
	IP param.Field[AccountAccessGroupUpdateParamsRequirePajwohLqIPRuleIP] `json:"ip,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqIPRule) implementsAccountAccessGroupUpdateParamsRequire() {
}

type AccountAccessGroupUpdateParamsRequirePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccountAccessGroupUpdateParamsRequirePajwohLqIPListRule struct {
	IPList param.Field[AccountAccessGroupUpdateParamsRequirePajwohLqIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqIPListRule) implementsAccountAccessGroupUpdateParamsRequire() {
}

type AccountAccessGroupUpdateParamsRequirePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccountAccessGroupUpdateParamsRequirePajwohLqCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqCertificateRule) implementsAccountAccessGroupUpdateParamsRequire() {
}

// Matches an Access group.
type AccountAccessGroupUpdateParamsRequirePajwohLqAccessGroupRule struct {
	Group param.Field[AccountAccessGroupUpdateParamsRequirePajwohLqAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqAccessGroupRule) implementsAccountAccessGroupUpdateParamsRequire() {
}

type AccountAccessGroupUpdateParamsRequirePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccountAccessGroupUpdateParamsRequirePajwohLqAzureGroupRule struct {
	AzureAd param.Field[AccountAccessGroupUpdateParamsRequirePajwohLqAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqAzureGroupRule) implementsAccountAccessGroupUpdateParamsRequire() {
}

type AccountAccessGroupUpdateParamsRequirePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccountAccessGroupUpdateParamsRequirePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccountAccessGroupUpdateParamsRequirePajwohLqGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqGitHubOrganizationRule) implementsAccountAccessGroupUpdateParamsRequire() {
}

type AccountAccessGroupUpdateParamsRequirePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccountAccessGroupUpdateParamsRequirePajwohLqGsuiteGroupRule struct {
	Gsuite param.Field[AccountAccessGroupUpdateParamsRequirePajwohLqGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqGsuiteGroupRule) implementsAccountAccessGroupUpdateParamsRequire() {
}

type AccountAccessGroupUpdateParamsRequirePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccountAccessGroupUpdateParamsRequirePajwohLqOktaGroupRule struct {
	Okta param.Field[AccountAccessGroupUpdateParamsRequirePajwohLqOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqOktaGroupRule) implementsAccountAccessGroupUpdateParamsRequire() {
}

type AccountAccessGroupUpdateParamsRequirePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccountAccessGroupUpdateParamsRequirePajwohLqSamlGroupRule struct {
	Saml param.Field[AccountAccessGroupUpdateParamsRequirePajwohLqSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqSamlGroupRule) implementsAccountAccessGroupUpdateParamsRequire() {
}

type AccountAccessGroupUpdateParamsRequirePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccountAccessGroupUpdateParamsRequirePajwohLqServiceTokenRule struct {
	ServiceToken param.Field[AccountAccessGroupUpdateParamsRequirePajwohLqServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqServiceTokenRule) implementsAccountAccessGroupUpdateParamsRequire() {
}

type AccountAccessGroupUpdateParamsRequirePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccountAccessGroupUpdateParamsRequirePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqAnyValidServiceTokenRule) implementsAccountAccessGroupUpdateParamsRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccountAccessGroupUpdateParamsRequirePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccountAccessGroupUpdateParamsRequirePajwohLqExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqExternalEvaluationRule) implementsAccountAccessGroupUpdateParamsRequire() {
}

type AccountAccessGroupUpdateParamsRequirePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccountAccessGroupUpdateParamsRequirePajwohLqCountryRule struct {
	Geo param.Field[AccountAccessGroupUpdateParamsRequirePajwohLqCountryRuleGeo] `json:"geo,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqCountryRule) implementsAccountAccessGroupUpdateParamsRequire() {
}

type AccountAccessGroupUpdateParamsRequirePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccountAccessGroupUpdateParamsRequirePajwohLqAuthenticationMethodRule struct {
	AuthMethod param.Field[AccountAccessGroupUpdateParamsRequirePajwohLqAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqAuthenticationMethodRule) implementsAccountAccessGroupUpdateParamsRequire() {
}

type AccountAccessGroupUpdateParamsRequirePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccountAccessGroupUpdateParamsRequirePajwohLqDevicePostureRule struct {
	DevicePosture param.Field[AccountAccessGroupUpdateParamsRequirePajwohLqDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqDevicePostureRule) implementsAccountAccessGroupUpdateParamsRequire() {
}

type AccountAccessGroupUpdateParamsRequirePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccountAccessGroupUpdateParamsRequirePajwohLqDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParams struct {
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude] `json:"include,required"`
	// The name of the Access group.
	Name param.Field[string] `json:"name,required"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude] `json:"exclude"`
	// Whether this is the default group
	IsDefault param.Field[bool] `json:"is_default"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire] `json:"require"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqEmailRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqEmailListRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqDomainRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqEveryoneRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqIPRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqIPListRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqCertificateRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqAccessGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqAzureGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqGitHubOrganizationRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqGsuiteGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqOktaGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqSamlGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqServiceTokenRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqAnyValidServiceTokenRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqExternalEvaluationRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqCountryRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqAuthenticationMethodRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqDevicePostureRule].
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude interface {
	implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude()
}

// Matches a specific email.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqEmailRule struct {
	Email param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqEmailRuleEmail] `json:"email,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqEmailRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqEmailListRule struct {
	EmailList param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqEmailListRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqDomainRule struct {
	EmailDomain param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqDomainRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqEveryoneRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

// Matches an IP address block.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqIPRule struct {
	IP param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqIPRuleIP] `json:"ip,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqIPRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqIPListRule struct {
	IPList param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqIPListRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqCertificateRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

// Matches an Access group.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqAccessGroupRule struct {
	Group param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqAccessGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqAzureGroupRule struct {
	AzureAd param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqAzureGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqGitHubOrganizationRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqGsuiteGroupRule struct {
	Gsuite param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqGsuiteGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqOktaGroupRule struct {
	Okta param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqOktaGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqSamlGroupRule struct {
	Saml param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqSamlGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqServiceTokenRule struct {
	ServiceToken param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqServiceTokenRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqAnyValidServiceTokenRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqExternalEvaluationRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqCountryRule struct {
	Geo param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqCountryRuleGeo] `json:"geo,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqCountryRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqAuthenticationMethodRule struct {
	AuthMethod param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqAuthenticationMethodRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqDevicePostureRule struct {
	DevicePosture param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqDevicePostureRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludePajwohLqDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqEmailRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqEmailListRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqDomainRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqEveryoneRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqIPRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqIPListRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqCertificateRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqAccessGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqAzureGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqGitHubOrganizationRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqGsuiteGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqOktaGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqSamlGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqServiceTokenRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqAnyValidServiceTokenRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqExternalEvaluationRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqCountryRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqAuthenticationMethodRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqDevicePostureRule].
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude interface {
	implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude()
}

// Matches a specific email.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqEmailRule struct {
	Email param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqEmailRuleEmail] `json:"email,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqEmailRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqEmailListRule struct {
	EmailList param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqEmailListRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqDomainRule struct {
	EmailDomain param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqDomainRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqEveryoneRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

// Matches an IP address block.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqIPRule struct {
	IP param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqIPRuleIP] `json:"ip,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqIPRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqIPListRule struct {
	IPList param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqIPListRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqCertificateRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

// Matches an Access group.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqAccessGroupRule struct {
	Group param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqAccessGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqAzureGroupRule struct {
	AzureAd param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqAzureGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqGitHubOrganizationRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqGsuiteGroupRule struct {
	Gsuite param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqGsuiteGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqOktaGroupRule struct {
	Okta param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqOktaGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqSamlGroupRule struct {
	Saml param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqSamlGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqServiceTokenRule struct {
	ServiceToken param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqServiceTokenRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqAnyValidServiceTokenRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqExternalEvaluationRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqCountryRule struct {
	Geo param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqCountryRuleGeo] `json:"geo,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqCountryRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqAuthenticationMethodRule struct {
	AuthMethod param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqAuthenticationMethodRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqDevicePostureRule struct {
	DevicePosture param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqDevicePostureRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludePajwohLqDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqEmailRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqEmailListRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqDomainRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqEveryoneRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqIPRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqIPListRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqCertificateRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqAccessGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqAzureGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqGitHubOrganizationRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqGsuiteGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqOktaGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqSamlGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqServiceTokenRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqAnyValidServiceTokenRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqExternalEvaluationRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqCountryRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqAuthenticationMethodRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqDevicePostureRule].
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire interface {
	implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire()
}

// Matches a specific email.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqEmailRule struct {
	Email param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqEmailRuleEmail] `json:"email,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqEmailRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqEmailListRule struct {
	EmailList param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqEmailListRuleEmailList] `json:"email_list,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqEmailListRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqDomainRule struct {
	EmailDomain param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqDomainRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqEveryoneRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

// Matches an IP address block.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqIPRule struct {
	IP param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqIPRuleIP] `json:"ip,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqIPRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqIPListRule struct {
	IPList param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqIPListRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqCertificateRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

// Matches an Access group.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqAccessGroupRule struct {
	Group param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqAccessGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqAzureGroupRule struct {
	AzureAd param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqAzureGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqGitHubOrganizationRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqGsuiteGroupRule struct {
	Gsuite param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqGsuiteGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqOktaGroupRule struct {
	Okta param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqOktaGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqSamlGroupRule struct {
	Saml param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqSamlGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqServiceTokenRule struct {
	ServiceToken param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqServiceTokenRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqAnyValidServiceTokenRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqExternalEvaluationRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqCountryRule struct {
	Geo param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqCountryRuleGeo] `json:"geo,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqCountryRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqAuthenticationMethodRule struct {
	AuthMethod param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqAuthenticationMethodRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqDevicePostureRule struct {
	DevicePosture param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqDevicePostureRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequirePajwohLqDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
