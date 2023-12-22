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
func (r *AccountAccessGroupService) Get(ctx context.Context, identifier string, uuid interface{}, opts ...option.RequestOption) (res *SingleResponseLef4sow9, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/groups/%v", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured Access group.
func (r *AccountAccessGroupService) Update(ctx context.Context, identifier string, uuid interface{}, body AccountAccessGroupUpdateParams, opts ...option.RequestOption) (res *SingleResponseLef4sow9, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/groups/%v", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes an Access group.
func (r *AccountAccessGroupService) Delete(ctx context.Context, identifier string, uuid interface{}, opts ...option.RequestOption) (res *AccountAccessGroupDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/groups/%v", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a new Access group.
func (r *AccountAccessGroupService) AccessGroupsNewAnAccessGroup(ctx context.Context, identifier string, body AccountAccessGroupAccessGroupsNewAnAccessGroupParams, opts ...option.RequestOption) (res *SingleResponseLef4sow9, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/groups", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all Access groups.
func (r *AccountAccessGroupService) AccessGroupsListAccessGroups(ctx context.Context, identifier string, opts ...option.RequestOption) (res *SchemasResponseCollectionF4F9gHlN, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/groups", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SchemasResponseCollectionF4F9gHlN struct {
	Errors     []SchemasResponseCollectionF4F9gHlNError    `json:"errors"`
	Messages   []SchemasResponseCollectionF4F9gHlNMessage  `json:"messages"`
	Result     []SchemasResponseCollectionF4F9gHlNResult   `json:"result"`
	ResultInfo SchemasResponseCollectionF4F9gHlNResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success SchemasResponseCollectionF4F9gHlNSuccess `json:"success"`
	JSON    schemasResponseCollectionF4F9gHlNJSON    `json:"-"`
}

// schemasResponseCollectionF4F9gHlNJSON contains the JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlN]
type schemasResponseCollectionF4F9gHlNJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasResponseCollectionF4F9gHlNError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    schemasResponseCollectionF4F9gHlNErrorJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNErrorJSON contains the JSON metadata for the
// struct [SchemasResponseCollectionF4F9gHlNError]
type schemasResponseCollectionF4F9gHlNErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasResponseCollectionF4F9gHlNMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    schemasResponseCollectionF4F9gHlNMessageJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNMessageJSON contains the JSON metadata for the
// struct [SchemasResponseCollectionF4F9gHlNMessage]
type schemasResponseCollectionF4F9gHlNMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasResponseCollectionF4F9gHlNResult struct {
	// The unique identifier for the Access group.
	ID        interface{} `json:"id"`
	CreatedAt time.Time   `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []SchemasResponseCollectionF4F9gHlNResultExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []SchemasResponseCollectionF4F9gHlNResultInclude `json:"include"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []SchemasResponseCollectionF4F9gHlNResultRequire `json:"require"`
	UpdatedAt time.Time                                        `json:"updated_at" format:"date-time"`
	JSON      schemasResponseCollectionF4F9gHlNResultJSON      `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultJSON contains the JSON metadata for the
// struct [SchemasResponseCollectionF4F9gHlNResult]
type schemasResponseCollectionF4F9gHlNResultJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Exclude     apijson.Field
	Include     apijson.Field
	Name        apijson.Field
	Require     apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [SchemasResponseCollectionF4F9gHlNResultExcludeEmailRule],
// [SchemasResponseCollectionF4F9gHlNResultExcludeDomainRule],
// [SchemasResponseCollectionF4F9gHlNResultExcludeEveryoneRule],
// [SchemasResponseCollectionF4F9gHlNResultExcludeIPRule],
// [SchemasResponseCollectionF4F9gHlNResultExcludeIPListRule],
// [SchemasResponseCollectionF4F9gHlNResultExcludeCertificateRule],
// [SchemasResponseCollectionF4F9gHlNResultExcludeAccessGroupRule],
// [SchemasResponseCollectionF4F9gHlNResultExcludeAzureGroupRule],
// [SchemasResponseCollectionF4F9gHlNResultExcludeGitHubOrganizationRule],
// [SchemasResponseCollectionF4F9gHlNResultExcludeGsuiteGroupRule],
// [SchemasResponseCollectionF4F9gHlNResultExcludeOktaGroupRule] or
// [SchemasResponseCollectionF4F9gHlNResultExcludeSamlGroupRule].
type SchemasResponseCollectionF4F9gHlNResultExclude interface {
	implementsSchemasResponseCollectionF4F9gHlNResultExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*SchemasResponseCollectionF4F9gHlNResultExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type SchemasResponseCollectionF4F9gHlNResultExcludeEmailRule struct {
	Email SchemasResponseCollectionF4F9gHlNResultExcludeEmailRuleEmail `json:"email,required"`
	JSON  schemasResponseCollectionF4F9gHlNResultExcludeEmailRuleJSON  `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultExcludeEmailRuleJSON contains the JSON
// metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultExcludeEmailRule]
type schemasResponseCollectionF4F9gHlNResultExcludeEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultExcludeEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultExcludeEmailRule) implementsSchemasResponseCollectionF4F9gHlNResultExclude() {
}

type SchemasResponseCollectionF4F9gHlNResultExcludeEmailRuleEmail struct {
	// The email of the user.
	Email string                                                           `json:"email,required" format:"email"`
	JSON  schemasResponseCollectionF4F9gHlNResultExcludeEmailRuleEmailJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultExcludeEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultExcludeEmailRuleEmail]
type schemasResponseCollectionF4F9gHlNResultExcludeEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultExcludeEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type SchemasResponseCollectionF4F9gHlNResultExcludeDomainRule struct {
	EmailDomain SchemasResponseCollectionF4F9gHlNResultExcludeDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        schemasResponseCollectionF4F9gHlNResultExcludeDomainRuleJSON        `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultExcludeDomainRuleJSON contains the JSON
// metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultExcludeDomainRule]
type schemasResponseCollectionF4F9gHlNResultExcludeDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultExcludeDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultExcludeDomainRule) implementsSchemasResponseCollectionF4F9gHlNResultExclude() {
}

type SchemasResponseCollectionF4F9gHlNResultExcludeDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                  `json:"domain,required"`
	JSON   schemasResponseCollectionF4F9gHlNResultExcludeDomainRuleEmailDomainJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultExcludeDomainRuleEmailDomainJSON contains
// the JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultExcludeDomainRuleEmailDomain]
type schemasResponseCollectionF4F9gHlNResultExcludeDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultExcludeDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type SchemasResponseCollectionF4F9gHlNResultExcludeEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                    `json:"everyone,required"`
	JSON     schemasResponseCollectionF4F9gHlNResultExcludeEveryoneRuleJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultExcludeEveryoneRuleJSON contains the JSON
// metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultExcludeEveryoneRule]
type schemasResponseCollectionF4F9gHlNResultExcludeEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultExcludeEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultExcludeEveryoneRule) implementsSchemasResponseCollectionF4F9gHlNResultExclude() {
}

// Matches an IP address block.
type SchemasResponseCollectionF4F9gHlNResultExcludeIPRule struct {
	IP   SchemasResponseCollectionF4F9gHlNResultExcludeIPRuleIP   `json:"ip,required"`
	JSON schemasResponseCollectionF4F9gHlNResultExcludeIPRuleJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultExcludeIPRuleJSON contains the JSON
// metadata for the struct [SchemasResponseCollectionF4F9gHlNResultExcludeIPRule]
type schemasResponseCollectionF4F9gHlNResultExcludeIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultExcludeIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultExcludeIPRule) implementsSchemasResponseCollectionF4F9gHlNResultExclude() {
}

type SchemasResponseCollectionF4F9gHlNResultExcludeIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                     `json:"ip,required"`
	JSON schemasResponseCollectionF4F9gHlNResultExcludeIPRuleIPJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultExcludeIPRuleIPJSON contains the JSON
// metadata for the struct [SchemasResponseCollectionF4F9gHlNResultExcludeIPRuleIP]
type schemasResponseCollectionF4F9gHlNResultExcludeIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultExcludeIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type SchemasResponseCollectionF4F9gHlNResultExcludeIPListRule struct {
	IPList SchemasResponseCollectionF4F9gHlNResultExcludeIPListRuleIPList `json:"ip_list,required"`
	JSON   schemasResponseCollectionF4F9gHlNResultExcludeIPListRuleJSON   `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultExcludeIPListRuleJSON contains the JSON
// metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultExcludeIPListRule]
type schemasResponseCollectionF4F9gHlNResultExcludeIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultExcludeIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultExcludeIPListRule) implementsSchemasResponseCollectionF4F9gHlNResultExclude() {
}

type SchemasResponseCollectionF4F9gHlNResultExcludeIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                             `json:"id,required"`
	JSON schemasResponseCollectionF4F9gHlNResultExcludeIPListRuleIPListJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultExcludeIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultExcludeIPListRuleIPList]
type schemasResponseCollectionF4F9gHlNResultExcludeIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultExcludeIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type SchemasResponseCollectionF4F9gHlNResultExcludeCertificateRule struct {
	Certificate interface{}                                                       `json:"certificate,required"`
	JSON        schemasResponseCollectionF4F9gHlNResultExcludeCertificateRuleJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultExcludeCertificateRuleJSON contains the
// JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultExcludeCertificateRule]
type schemasResponseCollectionF4F9gHlNResultExcludeCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultExcludeCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultExcludeCertificateRule) implementsSchemasResponseCollectionF4F9gHlNResultExclude() {
}

// Matches an Access group.
type SchemasResponseCollectionF4F9gHlNResultExcludeAccessGroupRule struct {
	Group SchemasResponseCollectionF4F9gHlNResultExcludeAccessGroupRuleGroup `json:"group,required"`
	JSON  schemasResponseCollectionF4F9gHlNResultExcludeAccessGroupRuleJSON  `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultExcludeAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultExcludeAccessGroupRule]
type schemasResponseCollectionF4F9gHlNResultExcludeAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultExcludeAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultExcludeAccessGroupRule) implementsSchemasResponseCollectionF4F9gHlNResultExclude() {
}

type SchemasResponseCollectionF4F9gHlNResultExcludeAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                 `json:"id,required"`
	JSON schemasResponseCollectionF4F9gHlNResultExcludeAccessGroupRuleGroupJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultExcludeAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultExcludeAccessGroupRuleGroup]
type schemasResponseCollectionF4F9gHlNResultExcludeAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultExcludeAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type SchemasResponseCollectionF4F9gHlNResultExcludeAzureGroupRule struct {
	AzureAd SchemasResponseCollectionF4F9gHlNResultExcludeAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    schemasResponseCollectionF4F9gHlNResultExcludeAzureGroupRuleJSON    `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultExcludeAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultExcludeAzureGroupRule]
type schemasResponseCollectionF4F9gHlNResultExcludeAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultExcludeAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultExcludeAzureGroupRule) implementsSchemasResponseCollectionF4F9gHlNResultExclude() {
}

type SchemasResponseCollectionF4F9gHlNResultExcludeAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                  `json:"connection_id,required"`
	JSON         schemasResponseCollectionF4F9gHlNResultExcludeAzureGroupRuleAzureAdJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultExcludeAzureGroupRuleAzureAdJSON contains
// the JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultExcludeAzureGroupRuleAzureAd]
type schemasResponseCollectionF4F9gHlNResultExcludeAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultExcludeAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type SchemasResponseCollectionF4F9gHlNResultExcludeGitHubOrganizationRule struct {
	GitHubOrganization SchemasResponseCollectionF4F9gHlNResultExcludeGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               schemasResponseCollectionF4F9gHlNResultExcludeGitHubOrganizationRuleJSON               `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultExcludeGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultExcludeGitHubOrganizationRule]
type schemasResponseCollectionF4F9gHlNResultExcludeGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultExcludeGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultExcludeGitHubOrganizationRule) implementsSchemasResponseCollectionF4F9gHlNResultExclude() {
}

type SchemasResponseCollectionF4F9gHlNResultExcludeGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                     `json:"name,required"`
	JSON schemasResponseCollectionF4F9gHlNResultExcludeGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultExcludeGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultExcludeGitHubOrganizationRuleGitHubOrganization]
type schemasResponseCollectionF4F9gHlNResultExcludeGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultExcludeGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type SchemasResponseCollectionF4F9gHlNResultExcludeGsuiteGroupRule struct {
	Gsuite SchemasResponseCollectionF4F9gHlNResultExcludeGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   schemasResponseCollectionF4F9gHlNResultExcludeGsuiteGroupRuleJSON   `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultExcludeGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultExcludeGsuiteGroupRule]
type schemasResponseCollectionF4F9gHlNResultExcludeGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultExcludeGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultExcludeGsuiteGroupRule) implementsSchemasResponseCollectionF4F9gHlNResultExclude() {
}

type SchemasResponseCollectionF4F9gHlNResultExcludeGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                  `json:"email,required"`
	JSON  schemasResponseCollectionF4F9gHlNResultExcludeGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultExcludeGsuiteGroupRuleGsuiteJSON contains
// the JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultExcludeGsuiteGroupRuleGsuite]
type schemasResponseCollectionF4F9gHlNResultExcludeGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultExcludeGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type SchemasResponseCollectionF4F9gHlNResultExcludeOktaGroupRule struct {
	Okta SchemasResponseCollectionF4F9gHlNResultExcludeOktaGroupRuleOkta `json:"okta,required"`
	JSON schemasResponseCollectionF4F9gHlNResultExcludeOktaGroupRuleJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultExcludeOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultExcludeOktaGroupRule]
type schemasResponseCollectionF4F9gHlNResultExcludeOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultExcludeOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultExcludeOktaGroupRule) implementsSchemasResponseCollectionF4F9gHlNResultExclude() {
}

type SchemasResponseCollectionF4F9gHlNResultExcludeOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                              `json:"email,required"`
	JSON  schemasResponseCollectionF4F9gHlNResultExcludeOktaGroupRuleOktaJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultExcludeOktaGroupRuleOktaJSON contains the
// JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultExcludeOktaGroupRuleOkta]
type schemasResponseCollectionF4F9gHlNResultExcludeOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultExcludeOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type SchemasResponseCollectionF4F9gHlNResultExcludeSamlGroupRule struct {
	Saml SchemasResponseCollectionF4F9gHlNResultExcludeSamlGroupRuleSaml `json:"saml,required"`
	JSON schemasResponseCollectionF4F9gHlNResultExcludeSamlGroupRuleJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultExcludeSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultExcludeSamlGroupRule]
type schemasResponseCollectionF4F9gHlNResultExcludeSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultExcludeSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultExcludeSamlGroupRule) implementsSchemasResponseCollectionF4F9gHlNResultExclude() {
}

type SchemasResponseCollectionF4F9gHlNResultExcludeSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                              `json:"attribute_value,required"`
	JSON           schemasResponseCollectionF4F9gHlNResultExcludeSamlGroupRuleSamlJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultExcludeSamlGroupRuleSamlJSON contains the
// JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultExcludeSamlGroupRuleSaml]
type schemasResponseCollectionF4F9gHlNResultExcludeSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultExcludeSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [SchemasResponseCollectionF4F9gHlNResultIncludeEmailRule],
// [SchemasResponseCollectionF4F9gHlNResultIncludeDomainRule],
// [SchemasResponseCollectionF4F9gHlNResultIncludeEveryoneRule],
// [SchemasResponseCollectionF4F9gHlNResultIncludeIPRule],
// [SchemasResponseCollectionF4F9gHlNResultIncludeIPListRule],
// [SchemasResponseCollectionF4F9gHlNResultIncludeCertificateRule],
// [SchemasResponseCollectionF4F9gHlNResultIncludeAccessGroupRule],
// [SchemasResponseCollectionF4F9gHlNResultIncludeAzureGroupRule],
// [SchemasResponseCollectionF4F9gHlNResultIncludeGitHubOrganizationRule],
// [SchemasResponseCollectionF4F9gHlNResultIncludeGsuiteGroupRule],
// [SchemasResponseCollectionF4F9gHlNResultIncludeOktaGroupRule] or
// [SchemasResponseCollectionF4F9gHlNResultIncludeSamlGroupRule].
type SchemasResponseCollectionF4F9gHlNResultInclude interface {
	implementsSchemasResponseCollectionF4F9gHlNResultInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*SchemasResponseCollectionF4F9gHlNResultInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type SchemasResponseCollectionF4F9gHlNResultIncludeEmailRule struct {
	Email SchemasResponseCollectionF4F9gHlNResultIncludeEmailRuleEmail `json:"email,required"`
	JSON  schemasResponseCollectionF4F9gHlNResultIncludeEmailRuleJSON  `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultIncludeEmailRuleJSON contains the JSON
// metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultIncludeEmailRule]
type schemasResponseCollectionF4F9gHlNResultIncludeEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultIncludeEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultIncludeEmailRule) implementsSchemasResponseCollectionF4F9gHlNResultInclude() {
}

type SchemasResponseCollectionF4F9gHlNResultIncludeEmailRuleEmail struct {
	// The email of the user.
	Email string                                                           `json:"email,required" format:"email"`
	JSON  schemasResponseCollectionF4F9gHlNResultIncludeEmailRuleEmailJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultIncludeEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultIncludeEmailRuleEmail]
type schemasResponseCollectionF4F9gHlNResultIncludeEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultIncludeEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type SchemasResponseCollectionF4F9gHlNResultIncludeDomainRule struct {
	EmailDomain SchemasResponseCollectionF4F9gHlNResultIncludeDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        schemasResponseCollectionF4F9gHlNResultIncludeDomainRuleJSON        `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultIncludeDomainRuleJSON contains the JSON
// metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultIncludeDomainRule]
type schemasResponseCollectionF4F9gHlNResultIncludeDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultIncludeDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultIncludeDomainRule) implementsSchemasResponseCollectionF4F9gHlNResultInclude() {
}

type SchemasResponseCollectionF4F9gHlNResultIncludeDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                  `json:"domain,required"`
	JSON   schemasResponseCollectionF4F9gHlNResultIncludeDomainRuleEmailDomainJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultIncludeDomainRuleEmailDomainJSON contains
// the JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultIncludeDomainRuleEmailDomain]
type schemasResponseCollectionF4F9gHlNResultIncludeDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultIncludeDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type SchemasResponseCollectionF4F9gHlNResultIncludeEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                    `json:"everyone,required"`
	JSON     schemasResponseCollectionF4F9gHlNResultIncludeEveryoneRuleJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultIncludeEveryoneRuleJSON contains the JSON
// metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultIncludeEveryoneRule]
type schemasResponseCollectionF4F9gHlNResultIncludeEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultIncludeEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultIncludeEveryoneRule) implementsSchemasResponseCollectionF4F9gHlNResultInclude() {
}

// Matches an IP address block.
type SchemasResponseCollectionF4F9gHlNResultIncludeIPRule struct {
	IP   SchemasResponseCollectionF4F9gHlNResultIncludeIPRuleIP   `json:"ip,required"`
	JSON schemasResponseCollectionF4F9gHlNResultIncludeIPRuleJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultIncludeIPRuleJSON contains the JSON
// metadata for the struct [SchemasResponseCollectionF4F9gHlNResultIncludeIPRule]
type schemasResponseCollectionF4F9gHlNResultIncludeIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultIncludeIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultIncludeIPRule) implementsSchemasResponseCollectionF4F9gHlNResultInclude() {
}

type SchemasResponseCollectionF4F9gHlNResultIncludeIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                     `json:"ip,required"`
	JSON schemasResponseCollectionF4F9gHlNResultIncludeIPRuleIPJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultIncludeIPRuleIPJSON contains the JSON
// metadata for the struct [SchemasResponseCollectionF4F9gHlNResultIncludeIPRuleIP]
type schemasResponseCollectionF4F9gHlNResultIncludeIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultIncludeIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type SchemasResponseCollectionF4F9gHlNResultIncludeIPListRule struct {
	IPList SchemasResponseCollectionF4F9gHlNResultIncludeIPListRuleIPList `json:"ip_list,required"`
	JSON   schemasResponseCollectionF4F9gHlNResultIncludeIPListRuleJSON   `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultIncludeIPListRuleJSON contains the JSON
// metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultIncludeIPListRule]
type schemasResponseCollectionF4F9gHlNResultIncludeIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultIncludeIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultIncludeIPListRule) implementsSchemasResponseCollectionF4F9gHlNResultInclude() {
}

type SchemasResponseCollectionF4F9gHlNResultIncludeIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                             `json:"id,required"`
	JSON schemasResponseCollectionF4F9gHlNResultIncludeIPListRuleIPListJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultIncludeIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultIncludeIPListRuleIPList]
type schemasResponseCollectionF4F9gHlNResultIncludeIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultIncludeIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type SchemasResponseCollectionF4F9gHlNResultIncludeCertificateRule struct {
	Certificate interface{}                                                       `json:"certificate,required"`
	JSON        schemasResponseCollectionF4F9gHlNResultIncludeCertificateRuleJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultIncludeCertificateRuleJSON contains the
// JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultIncludeCertificateRule]
type schemasResponseCollectionF4F9gHlNResultIncludeCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultIncludeCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultIncludeCertificateRule) implementsSchemasResponseCollectionF4F9gHlNResultInclude() {
}

// Matches an Access group.
type SchemasResponseCollectionF4F9gHlNResultIncludeAccessGroupRule struct {
	Group SchemasResponseCollectionF4F9gHlNResultIncludeAccessGroupRuleGroup `json:"group,required"`
	JSON  schemasResponseCollectionF4F9gHlNResultIncludeAccessGroupRuleJSON  `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultIncludeAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultIncludeAccessGroupRule]
type schemasResponseCollectionF4F9gHlNResultIncludeAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultIncludeAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultIncludeAccessGroupRule) implementsSchemasResponseCollectionF4F9gHlNResultInclude() {
}

type SchemasResponseCollectionF4F9gHlNResultIncludeAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                 `json:"id,required"`
	JSON schemasResponseCollectionF4F9gHlNResultIncludeAccessGroupRuleGroupJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultIncludeAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultIncludeAccessGroupRuleGroup]
type schemasResponseCollectionF4F9gHlNResultIncludeAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultIncludeAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type SchemasResponseCollectionF4F9gHlNResultIncludeAzureGroupRule struct {
	AzureAd SchemasResponseCollectionF4F9gHlNResultIncludeAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    schemasResponseCollectionF4F9gHlNResultIncludeAzureGroupRuleJSON    `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultIncludeAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultIncludeAzureGroupRule]
type schemasResponseCollectionF4F9gHlNResultIncludeAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultIncludeAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultIncludeAzureGroupRule) implementsSchemasResponseCollectionF4F9gHlNResultInclude() {
}

type SchemasResponseCollectionF4F9gHlNResultIncludeAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                  `json:"connection_id,required"`
	JSON         schemasResponseCollectionF4F9gHlNResultIncludeAzureGroupRuleAzureAdJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultIncludeAzureGroupRuleAzureAdJSON contains
// the JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultIncludeAzureGroupRuleAzureAd]
type schemasResponseCollectionF4F9gHlNResultIncludeAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultIncludeAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type SchemasResponseCollectionF4F9gHlNResultIncludeGitHubOrganizationRule struct {
	GitHubOrganization SchemasResponseCollectionF4F9gHlNResultIncludeGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               schemasResponseCollectionF4F9gHlNResultIncludeGitHubOrganizationRuleJSON               `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultIncludeGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultIncludeGitHubOrganizationRule]
type schemasResponseCollectionF4F9gHlNResultIncludeGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultIncludeGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultIncludeGitHubOrganizationRule) implementsSchemasResponseCollectionF4F9gHlNResultInclude() {
}

type SchemasResponseCollectionF4F9gHlNResultIncludeGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                     `json:"name,required"`
	JSON schemasResponseCollectionF4F9gHlNResultIncludeGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultIncludeGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultIncludeGitHubOrganizationRuleGitHubOrganization]
type schemasResponseCollectionF4F9gHlNResultIncludeGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultIncludeGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type SchemasResponseCollectionF4F9gHlNResultIncludeGsuiteGroupRule struct {
	Gsuite SchemasResponseCollectionF4F9gHlNResultIncludeGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   schemasResponseCollectionF4F9gHlNResultIncludeGsuiteGroupRuleJSON   `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultIncludeGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultIncludeGsuiteGroupRule]
type schemasResponseCollectionF4F9gHlNResultIncludeGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultIncludeGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultIncludeGsuiteGroupRule) implementsSchemasResponseCollectionF4F9gHlNResultInclude() {
}

type SchemasResponseCollectionF4F9gHlNResultIncludeGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                  `json:"email,required"`
	JSON  schemasResponseCollectionF4F9gHlNResultIncludeGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultIncludeGsuiteGroupRuleGsuiteJSON contains
// the JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultIncludeGsuiteGroupRuleGsuite]
type schemasResponseCollectionF4F9gHlNResultIncludeGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultIncludeGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type SchemasResponseCollectionF4F9gHlNResultIncludeOktaGroupRule struct {
	Okta SchemasResponseCollectionF4F9gHlNResultIncludeOktaGroupRuleOkta `json:"okta,required"`
	JSON schemasResponseCollectionF4F9gHlNResultIncludeOktaGroupRuleJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultIncludeOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultIncludeOktaGroupRule]
type schemasResponseCollectionF4F9gHlNResultIncludeOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultIncludeOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultIncludeOktaGroupRule) implementsSchemasResponseCollectionF4F9gHlNResultInclude() {
}

type SchemasResponseCollectionF4F9gHlNResultIncludeOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                              `json:"email,required"`
	JSON  schemasResponseCollectionF4F9gHlNResultIncludeOktaGroupRuleOktaJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultIncludeOktaGroupRuleOktaJSON contains the
// JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultIncludeOktaGroupRuleOkta]
type schemasResponseCollectionF4F9gHlNResultIncludeOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultIncludeOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type SchemasResponseCollectionF4F9gHlNResultIncludeSamlGroupRule struct {
	Saml SchemasResponseCollectionF4F9gHlNResultIncludeSamlGroupRuleSaml `json:"saml,required"`
	JSON schemasResponseCollectionF4F9gHlNResultIncludeSamlGroupRuleJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultIncludeSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultIncludeSamlGroupRule]
type schemasResponseCollectionF4F9gHlNResultIncludeSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultIncludeSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultIncludeSamlGroupRule) implementsSchemasResponseCollectionF4F9gHlNResultInclude() {
}

type SchemasResponseCollectionF4F9gHlNResultIncludeSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                              `json:"attribute_value,required"`
	JSON           schemasResponseCollectionF4F9gHlNResultIncludeSamlGroupRuleSamlJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultIncludeSamlGroupRuleSamlJSON contains the
// JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultIncludeSamlGroupRuleSaml]
type schemasResponseCollectionF4F9gHlNResultIncludeSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultIncludeSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [SchemasResponseCollectionF4F9gHlNResultRequireEmailRule],
// [SchemasResponseCollectionF4F9gHlNResultRequireDomainRule],
// [SchemasResponseCollectionF4F9gHlNResultRequireEveryoneRule],
// [SchemasResponseCollectionF4F9gHlNResultRequireIPRule],
// [SchemasResponseCollectionF4F9gHlNResultRequireIPListRule],
// [SchemasResponseCollectionF4F9gHlNResultRequireCertificateRule],
// [SchemasResponseCollectionF4F9gHlNResultRequireAccessGroupRule],
// [SchemasResponseCollectionF4F9gHlNResultRequireAzureGroupRule],
// [SchemasResponseCollectionF4F9gHlNResultRequireGitHubOrganizationRule],
// [SchemasResponseCollectionF4F9gHlNResultRequireGsuiteGroupRule],
// [SchemasResponseCollectionF4F9gHlNResultRequireOktaGroupRule] or
// [SchemasResponseCollectionF4F9gHlNResultRequireSamlGroupRule].
type SchemasResponseCollectionF4F9gHlNResultRequire interface {
	implementsSchemasResponseCollectionF4F9gHlNResultRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*SchemasResponseCollectionF4F9gHlNResultRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type SchemasResponseCollectionF4F9gHlNResultRequireEmailRule struct {
	Email SchemasResponseCollectionF4F9gHlNResultRequireEmailRuleEmail `json:"email,required"`
	JSON  schemasResponseCollectionF4F9gHlNResultRequireEmailRuleJSON  `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultRequireEmailRuleJSON contains the JSON
// metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultRequireEmailRule]
type schemasResponseCollectionF4F9gHlNResultRequireEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultRequireEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultRequireEmailRule) implementsSchemasResponseCollectionF4F9gHlNResultRequire() {
}

type SchemasResponseCollectionF4F9gHlNResultRequireEmailRuleEmail struct {
	// The email of the user.
	Email string                                                           `json:"email,required" format:"email"`
	JSON  schemasResponseCollectionF4F9gHlNResultRequireEmailRuleEmailJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultRequireEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultRequireEmailRuleEmail]
type schemasResponseCollectionF4F9gHlNResultRequireEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultRequireEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type SchemasResponseCollectionF4F9gHlNResultRequireDomainRule struct {
	EmailDomain SchemasResponseCollectionF4F9gHlNResultRequireDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        schemasResponseCollectionF4F9gHlNResultRequireDomainRuleJSON        `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultRequireDomainRuleJSON contains the JSON
// metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultRequireDomainRule]
type schemasResponseCollectionF4F9gHlNResultRequireDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultRequireDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultRequireDomainRule) implementsSchemasResponseCollectionF4F9gHlNResultRequire() {
}

type SchemasResponseCollectionF4F9gHlNResultRequireDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                  `json:"domain,required"`
	JSON   schemasResponseCollectionF4F9gHlNResultRequireDomainRuleEmailDomainJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultRequireDomainRuleEmailDomainJSON contains
// the JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultRequireDomainRuleEmailDomain]
type schemasResponseCollectionF4F9gHlNResultRequireDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultRequireDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type SchemasResponseCollectionF4F9gHlNResultRequireEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                    `json:"everyone,required"`
	JSON     schemasResponseCollectionF4F9gHlNResultRequireEveryoneRuleJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultRequireEveryoneRuleJSON contains the JSON
// metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultRequireEveryoneRule]
type schemasResponseCollectionF4F9gHlNResultRequireEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultRequireEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultRequireEveryoneRule) implementsSchemasResponseCollectionF4F9gHlNResultRequire() {
}

// Matches an IP address block.
type SchemasResponseCollectionF4F9gHlNResultRequireIPRule struct {
	IP   SchemasResponseCollectionF4F9gHlNResultRequireIPRuleIP   `json:"ip,required"`
	JSON schemasResponseCollectionF4F9gHlNResultRequireIPRuleJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultRequireIPRuleJSON contains the JSON
// metadata for the struct [SchemasResponseCollectionF4F9gHlNResultRequireIPRule]
type schemasResponseCollectionF4F9gHlNResultRequireIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultRequireIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultRequireIPRule) implementsSchemasResponseCollectionF4F9gHlNResultRequire() {
}

type SchemasResponseCollectionF4F9gHlNResultRequireIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                     `json:"ip,required"`
	JSON schemasResponseCollectionF4F9gHlNResultRequireIPRuleIPJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultRequireIPRuleIPJSON contains the JSON
// metadata for the struct [SchemasResponseCollectionF4F9gHlNResultRequireIPRuleIP]
type schemasResponseCollectionF4F9gHlNResultRequireIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultRequireIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type SchemasResponseCollectionF4F9gHlNResultRequireIPListRule struct {
	IPList SchemasResponseCollectionF4F9gHlNResultRequireIPListRuleIPList `json:"ip_list,required"`
	JSON   schemasResponseCollectionF4F9gHlNResultRequireIPListRuleJSON   `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultRequireIPListRuleJSON contains the JSON
// metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultRequireIPListRule]
type schemasResponseCollectionF4F9gHlNResultRequireIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultRequireIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultRequireIPListRule) implementsSchemasResponseCollectionF4F9gHlNResultRequire() {
}

type SchemasResponseCollectionF4F9gHlNResultRequireIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                             `json:"id,required"`
	JSON schemasResponseCollectionF4F9gHlNResultRequireIPListRuleIPListJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultRequireIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultRequireIPListRuleIPList]
type schemasResponseCollectionF4F9gHlNResultRequireIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultRequireIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type SchemasResponseCollectionF4F9gHlNResultRequireCertificateRule struct {
	Certificate interface{}                                                       `json:"certificate,required"`
	JSON        schemasResponseCollectionF4F9gHlNResultRequireCertificateRuleJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultRequireCertificateRuleJSON contains the
// JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultRequireCertificateRule]
type schemasResponseCollectionF4F9gHlNResultRequireCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultRequireCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultRequireCertificateRule) implementsSchemasResponseCollectionF4F9gHlNResultRequire() {
}

// Matches an Access group.
type SchemasResponseCollectionF4F9gHlNResultRequireAccessGroupRule struct {
	Group SchemasResponseCollectionF4F9gHlNResultRequireAccessGroupRuleGroup `json:"group,required"`
	JSON  schemasResponseCollectionF4F9gHlNResultRequireAccessGroupRuleJSON  `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultRequireAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultRequireAccessGroupRule]
type schemasResponseCollectionF4F9gHlNResultRequireAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultRequireAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultRequireAccessGroupRule) implementsSchemasResponseCollectionF4F9gHlNResultRequire() {
}

type SchemasResponseCollectionF4F9gHlNResultRequireAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                 `json:"id,required"`
	JSON schemasResponseCollectionF4F9gHlNResultRequireAccessGroupRuleGroupJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultRequireAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultRequireAccessGroupRuleGroup]
type schemasResponseCollectionF4F9gHlNResultRequireAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultRequireAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type SchemasResponseCollectionF4F9gHlNResultRequireAzureGroupRule struct {
	AzureAd SchemasResponseCollectionF4F9gHlNResultRequireAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    schemasResponseCollectionF4F9gHlNResultRequireAzureGroupRuleJSON    `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultRequireAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultRequireAzureGroupRule]
type schemasResponseCollectionF4F9gHlNResultRequireAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultRequireAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultRequireAzureGroupRule) implementsSchemasResponseCollectionF4F9gHlNResultRequire() {
}

type SchemasResponseCollectionF4F9gHlNResultRequireAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                  `json:"connection_id,required"`
	JSON         schemasResponseCollectionF4F9gHlNResultRequireAzureGroupRuleAzureAdJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultRequireAzureGroupRuleAzureAdJSON contains
// the JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultRequireAzureGroupRuleAzureAd]
type schemasResponseCollectionF4F9gHlNResultRequireAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultRequireAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type SchemasResponseCollectionF4F9gHlNResultRequireGitHubOrganizationRule struct {
	GitHubOrganization SchemasResponseCollectionF4F9gHlNResultRequireGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               schemasResponseCollectionF4F9gHlNResultRequireGitHubOrganizationRuleJSON               `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultRequireGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultRequireGitHubOrganizationRule]
type schemasResponseCollectionF4F9gHlNResultRequireGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultRequireGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultRequireGitHubOrganizationRule) implementsSchemasResponseCollectionF4F9gHlNResultRequire() {
}

type SchemasResponseCollectionF4F9gHlNResultRequireGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                     `json:"name,required"`
	JSON schemasResponseCollectionF4F9gHlNResultRequireGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultRequireGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultRequireGitHubOrganizationRuleGitHubOrganization]
type schemasResponseCollectionF4F9gHlNResultRequireGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultRequireGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type SchemasResponseCollectionF4F9gHlNResultRequireGsuiteGroupRule struct {
	Gsuite SchemasResponseCollectionF4F9gHlNResultRequireGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   schemasResponseCollectionF4F9gHlNResultRequireGsuiteGroupRuleJSON   `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultRequireGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultRequireGsuiteGroupRule]
type schemasResponseCollectionF4F9gHlNResultRequireGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultRequireGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultRequireGsuiteGroupRule) implementsSchemasResponseCollectionF4F9gHlNResultRequire() {
}

type SchemasResponseCollectionF4F9gHlNResultRequireGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                  `json:"email,required"`
	JSON  schemasResponseCollectionF4F9gHlNResultRequireGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultRequireGsuiteGroupRuleGsuiteJSON contains
// the JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultRequireGsuiteGroupRuleGsuite]
type schemasResponseCollectionF4F9gHlNResultRequireGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultRequireGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type SchemasResponseCollectionF4F9gHlNResultRequireOktaGroupRule struct {
	Okta SchemasResponseCollectionF4F9gHlNResultRequireOktaGroupRuleOkta `json:"okta,required"`
	JSON schemasResponseCollectionF4F9gHlNResultRequireOktaGroupRuleJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultRequireOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultRequireOktaGroupRule]
type schemasResponseCollectionF4F9gHlNResultRequireOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultRequireOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultRequireOktaGroupRule) implementsSchemasResponseCollectionF4F9gHlNResultRequire() {
}

type SchemasResponseCollectionF4F9gHlNResultRequireOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                              `json:"email,required"`
	JSON  schemasResponseCollectionF4F9gHlNResultRequireOktaGroupRuleOktaJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultRequireOktaGroupRuleOktaJSON contains the
// JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultRequireOktaGroupRuleOkta]
type schemasResponseCollectionF4F9gHlNResultRequireOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultRequireOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type SchemasResponseCollectionF4F9gHlNResultRequireSamlGroupRule struct {
	Saml SchemasResponseCollectionF4F9gHlNResultRequireSamlGroupRuleSaml `json:"saml,required"`
	JSON schemasResponseCollectionF4F9gHlNResultRequireSamlGroupRuleJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultRequireSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultRequireSamlGroupRule]
type schemasResponseCollectionF4F9gHlNResultRequireSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultRequireSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SchemasResponseCollectionF4F9gHlNResultRequireSamlGroupRule) implementsSchemasResponseCollectionF4F9gHlNResultRequire() {
}

type SchemasResponseCollectionF4F9gHlNResultRequireSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                              `json:"attribute_value,required"`
	JSON           schemasResponseCollectionF4F9gHlNResultRequireSamlGroupRuleSamlJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultRequireSamlGroupRuleSamlJSON contains the
// JSON metadata for the struct
// [SchemasResponseCollectionF4F9gHlNResultRequireSamlGroupRuleSaml]
type schemasResponseCollectionF4F9gHlNResultRequireSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultRequireSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasResponseCollectionF4F9gHlNResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                         `json:"total_count"`
	JSON       schemasResponseCollectionF4F9gHlNResultInfoJSON `json:"-"`
}

// schemasResponseCollectionF4F9gHlNResultInfoJSON contains the JSON metadata for
// the struct [SchemasResponseCollectionF4F9gHlNResultInfo]
type schemasResponseCollectionF4F9gHlNResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionF4F9gHlNResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SchemasResponseCollectionF4F9gHlNSuccess bool

const (
	SchemasResponseCollectionF4F9gHlNSuccessTrue SchemasResponseCollectionF4F9gHlNSuccess = true
)

type SingleResponseLef4sow9 struct {
	Errors   []SingleResponseLef4sow9Error   `json:"errors"`
	Messages []SingleResponseLef4sow9Message `json:"messages"`
	Result   SingleResponseLef4sow9Result    `json:"result"`
	// Whether the API call was successful
	Success SingleResponseLef4sow9Success `json:"success"`
	JSON    singleResponseLef4sow9JSON    `json:"-"`
}

// singleResponseLef4sow9JSON contains the JSON metadata for the struct
// [SingleResponseLef4sow9]
type singleResponseLef4sow9JSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseLef4sow9Error struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    singleResponseLef4sow9ErrorJSON `json:"-"`
}

// singleResponseLef4sow9ErrorJSON contains the JSON metadata for the struct
// [SingleResponseLef4sow9Error]
type singleResponseLef4sow9ErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9Error) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseLef4sow9Message struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    singleResponseLef4sow9MessageJSON `json:"-"`
}

// singleResponseLef4sow9MessageJSON contains the JSON metadata for the struct
// [SingleResponseLef4sow9Message]
type singleResponseLef4sow9MessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9Message) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseLef4sow9Result struct {
	// The unique identifier for the Access group.
	ID        interface{} `json:"id"`
	CreatedAt time.Time   `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []SingleResponseLef4sow9ResultExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []SingleResponseLef4sow9ResultInclude `json:"include"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []SingleResponseLef4sow9ResultRequire `json:"require"`
	UpdatedAt time.Time                             `json:"updated_at" format:"date-time"`
	JSON      singleResponseLef4sow9ResultJSON      `json:"-"`
}

// singleResponseLef4sow9ResultJSON contains the JSON metadata for the struct
// [SingleResponseLef4sow9Result]
type singleResponseLef4sow9ResultJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Exclude     apijson.Field
	Include     apijson.Field
	Name        apijson.Field
	Require     apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9Result) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [SingleResponseLef4sow9ResultExcludeEmailRule],
// [SingleResponseLef4sow9ResultExcludeDomainRule],
// [SingleResponseLef4sow9ResultExcludeEveryoneRule],
// [SingleResponseLef4sow9ResultExcludeIPRule],
// [SingleResponseLef4sow9ResultExcludeIPListRule],
// [SingleResponseLef4sow9ResultExcludeCertificateRule],
// [SingleResponseLef4sow9ResultExcludeAccessGroupRule],
// [SingleResponseLef4sow9ResultExcludeAzureGroupRule],
// [SingleResponseLef4sow9ResultExcludeGitHubOrganizationRule],
// [SingleResponseLef4sow9ResultExcludeGsuiteGroupRule],
// [SingleResponseLef4sow9ResultExcludeOktaGroupRule] or
// [SingleResponseLef4sow9ResultExcludeSamlGroupRule].
type SingleResponseLef4sow9ResultExclude interface {
	implementsSingleResponseLef4sow9ResultExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*SingleResponseLef4sow9ResultExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type SingleResponseLef4sow9ResultExcludeEmailRule struct {
	Email SingleResponseLef4sow9ResultExcludeEmailRuleEmail `json:"email,required"`
	JSON  singleResponseLef4sow9ResultExcludeEmailRuleJSON  `json:"-"`
}

// singleResponseLef4sow9ResultExcludeEmailRuleJSON contains the JSON metadata for
// the struct [SingleResponseLef4sow9ResultExcludeEmailRule]
type singleResponseLef4sow9ResultExcludeEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultExcludeEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultExcludeEmailRule) implementsSingleResponseLef4sow9ResultExclude() {
}

type SingleResponseLef4sow9ResultExcludeEmailRuleEmail struct {
	// The email of the user.
	Email string                                                `json:"email,required" format:"email"`
	JSON  singleResponseLef4sow9ResultExcludeEmailRuleEmailJSON `json:"-"`
}

// singleResponseLef4sow9ResultExcludeEmailRuleEmailJSON contains the JSON metadata
// for the struct [SingleResponseLef4sow9ResultExcludeEmailRuleEmail]
type singleResponseLef4sow9ResultExcludeEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultExcludeEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type SingleResponseLef4sow9ResultExcludeDomainRule struct {
	EmailDomain SingleResponseLef4sow9ResultExcludeDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        singleResponseLef4sow9ResultExcludeDomainRuleJSON        `json:"-"`
}

// singleResponseLef4sow9ResultExcludeDomainRuleJSON contains the JSON metadata for
// the struct [SingleResponseLef4sow9ResultExcludeDomainRule]
type singleResponseLef4sow9ResultExcludeDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultExcludeDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultExcludeDomainRule) implementsSingleResponseLef4sow9ResultExclude() {
}

type SingleResponseLef4sow9ResultExcludeDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                       `json:"domain,required"`
	JSON   singleResponseLef4sow9ResultExcludeDomainRuleEmailDomainJSON `json:"-"`
}

// singleResponseLef4sow9ResultExcludeDomainRuleEmailDomainJSON contains the JSON
// metadata for the struct
// [SingleResponseLef4sow9ResultExcludeDomainRuleEmailDomain]
type singleResponseLef4sow9ResultExcludeDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultExcludeDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type SingleResponseLef4sow9ResultExcludeEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                         `json:"everyone,required"`
	JSON     singleResponseLef4sow9ResultExcludeEveryoneRuleJSON `json:"-"`
}

// singleResponseLef4sow9ResultExcludeEveryoneRuleJSON contains the JSON metadata
// for the struct [SingleResponseLef4sow9ResultExcludeEveryoneRule]
type singleResponseLef4sow9ResultExcludeEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultExcludeEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultExcludeEveryoneRule) implementsSingleResponseLef4sow9ResultExclude() {
}

// Matches an IP address block.
type SingleResponseLef4sow9ResultExcludeIPRule struct {
	IP   SingleResponseLef4sow9ResultExcludeIPRuleIP   `json:"ip,required"`
	JSON singleResponseLef4sow9ResultExcludeIPRuleJSON `json:"-"`
}

// singleResponseLef4sow9ResultExcludeIPRuleJSON contains the JSON metadata for the
// struct [SingleResponseLef4sow9ResultExcludeIPRule]
type singleResponseLef4sow9ResultExcludeIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultExcludeIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultExcludeIPRule) implementsSingleResponseLef4sow9ResultExclude() {}

type SingleResponseLef4sow9ResultExcludeIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                          `json:"ip,required"`
	JSON singleResponseLef4sow9ResultExcludeIPRuleIPJSON `json:"-"`
}

// singleResponseLef4sow9ResultExcludeIPRuleIPJSON contains the JSON metadata for
// the struct [SingleResponseLef4sow9ResultExcludeIPRuleIP]
type singleResponseLef4sow9ResultExcludeIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultExcludeIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type SingleResponseLef4sow9ResultExcludeIPListRule struct {
	IPList SingleResponseLef4sow9ResultExcludeIPListRuleIPList `json:"ip_list,required"`
	JSON   singleResponseLef4sow9ResultExcludeIPListRuleJSON   `json:"-"`
}

// singleResponseLef4sow9ResultExcludeIPListRuleJSON contains the JSON metadata for
// the struct [SingleResponseLef4sow9ResultExcludeIPListRule]
type singleResponseLef4sow9ResultExcludeIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultExcludeIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultExcludeIPListRule) implementsSingleResponseLef4sow9ResultExclude() {
}

type SingleResponseLef4sow9ResultExcludeIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                  `json:"id,required"`
	JSON singleResponseLef4sow9ResultExcludeIPListRuleIPListJSON `json:"-"`
}

// singleResponseLef4sow9ResultExcludeIPListRuleIPListJSON contains the JSON
// metadata for the struct [SingleResponseLef4sow9ResultExcludeIPListRuleIPList]
type singleResponseLef4sow9ResultExcludeIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultExcludeIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type SingleResponseLef4sow9ResultExcludeCertificateRule struct {
	Certificate interface{}                                            `json:"certificate,required"`
	JSON        singleResponseLef4sow9ResultExcludeCertificateRuleJSON `json:"-"`
}

// singleResponseLef4sow9ResultExcludeCertificateRuleJSON contains the JSON
// metadata for the struct [SingleResponseLef4sow9ResultExcludeCertificateRule]
type singleResponseLef4sow9ResultExcludeCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultExcludeCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultExcludeCertificateRule) implementsSingleResponseLef4sow9ResultExclude() {
}

// Matches an Access group.
type SingleResponseLef4sow9ResultExcludeAccessGroupRule struct {
	Group SingleResponseLef4sow9ResultExcludeAccessGroupRuleGroup `json:"group,required"`
	JSON  singleResponseLef4sow9ResultExcludeAccessGroupRuleJSON  `json:"-"`
}

// singleResponseLef4sow9ResultExcludeAccessGroupRuleJSON contains the JSON
// metadata for the struct [SingleResponseLef4sow9ResultExcludeAccessGroupRule]
type singleResponseLef4sow9ResultExcludeAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultExcludeAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultExcludeAccessGroupRule) implementsSingleResponseLef4sow9ResultExclude() {
}

type SingleResponseLef4sow9ResultExcludeAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                      `json:"id,required"`
	JSON singleResponseLef4sow9ResultExcludeAccessGroupRuleGroupJSON `json:"-"`
}

// singleResponseLef4sow9ResultExcludeAccessGroupRuleGroupJSON contains the JSON
// metadata for the struct
// [SingleResponseLef4sow9ResultExcludeAccessGroupRuleGroup]
type singleResponseLef4sow9ResultExcludeAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultExcludeAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type SingleResponseLef4sow9ResultExcludeAzureGroupRule struct {
	AzureAd SingleResponseLef4sow9ResultExcludeAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    singleResponseLef4sow9ResultExcludeAzureGroupRuleJSON    `json:"-"`
}

// singleResponseLef4sow9ResultExcludeAzureGroupRuleJSON contains the JSON metadata
// for the struct [SingleResponseLef4sow9ResultExcludeAzureGroupRule]
type singleResponseLef4sow9ResultExcludeAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultExcludeAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultExcludeAzureGroupRule) implementsSingleResponseLef4sow9ResultExclude() {
}

type SingleResponseLef4sow9ResultExcludeAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                       `json:"connection_id,required"`
	JSON         singleResponseLef4sow9ResultExcludeAzureGroupRuleAzureAdJSON `json:"-"`
}

// singleResponseLef4sow9ResultExcludeAzureGroupRuleAzureAdJSON contains the JSON
// metadata for the struct
// [SingleResponseLef4sow9ResultExcludeAzureGroupRuleAzureAd]
type singleResponseLef4sow9ResultExcludeAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultExcludeAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type SingleResponseLef4sow9ResultExcludeGitHubOrganizationRule struct {
	GitHubOrganization SingleResponseLef4sow9ResultExcludeGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               singleResponseLef4sow9ResultExcludeGitHubOrganizationRuleJSON               `json:"-"`
}

// singleResponseLef4sow9ResultExcludeGitHubOrganizationRuleJSON contains the JSON
// metadata for the struct
// [SingleResponseLef4sow9ResultExcludeGitHubOrganizationRule]
type singleResponseLef4sow9ResultExcludeGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultExcludeGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultExcludeGitHubOrganizationRule) implementsSingleResponseLef4sow9ResultExclude() {
}

type SingleResponseLef4sow9ResultExcludeGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                          `json:"name,required"`
	JSON singleResponseLef4sow9ResultExcludeGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// singleResponseLef4sow9ResultExcludeGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [SingleResponseLef4sow9ResultExcludeGitHubOrganizationRuleGitHubOrganization]
type singleResponseLef4sow9ResultExcludeGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultExcludeGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type SingleResponseLef4sow9ResultExcludeGsuiteGroupRule struct {
	Gsuite SingleResponseLef4sow9ResultExcludeGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   singleResponseLef4sow9ResultExcludeGsuiteGroupRuleJSON   `json:"-"`
}

// singleResponseLef4sow9ResultExcludeGsuiteGroupRuleJSON contains the JSON
// metadata for the struct [SingleResponseLef4sow9ResultExcludeGsuiteGroupRule]
type singleResponseLef4sow9ResultExcludeGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultExcludeGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultExcludeGsuiteGroupRule) implementsSingleResponseLef4sow9ResultExclude() {
}

type SingleResponseLef4sow9ResultExcludeGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                       `json:"email,required"`
	JSON  singleResponseLef4sow9ResultExcludeGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// singleResponseLef4sow9ResultExcludeGsuiteGroupRuleGsuiteJSON contains the JSON
// metadata for the struct
// [SingleResponseLef4sow9ResultExcludeGsuiteGroupRuleGsuite]
type singleResponseLef4sow9ResultExcludeGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultExcludeGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type SingleResponseLef4sow9ResultExcludeOktaGroupRule struct {
	Okta SingleResponseLef4sow9ResultExcludeOktaGroupRuleOkta `json:"okta,required"`
	JSON singleResponseLef4sow9ResultExcludeOktaGroupRuleJSON `json:"-"`
}

// singleResponseLef4sow9ResultExcludeOktaGroupRuleJSON contains the JSON metadata
// for the struct [SingleResponseLef4sow9ResultExcludeOktaGroupRule]
type singleResponseLef4sow9ResultExcludeOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultExcludeOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultExcludeOktaGroupRule) implementsSingleResponseLef4sow9ResultExclude() {
}

type SingleResponseLef4sow9ResultExcludeOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                   `json:"email,required"`
	JSON  singleResponseLef4sow9ResultExcludeOktaGroupRuleOktaJSON `json:"-"`
}

// singleResponseLef4sow9ResultExcludeOktaGroupRuleOktaJSON contains the JSON
// metadata for the struct [SingleResponseLef4sow9ResultExcludeOktaGroupRuleOkta]
type singleResponseLef4sow9ResultExcludeOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultExcludeOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type SingleResponseLef4sow9ResultExcludeSamlGroupRule struct {
	Saml SingleResponseLef4sow9ResultExcludeSamlGroupRuleSaml `json:"saml,required"`
	JSON singleResponseLef4sow9ResultExcludeSamlGroupRuleJSON `json:"-"`
}

// singleResponseLef4sow9ResultExcludeSamlGroupRuleJSON contains the JSON metadata
// for the struct [SingleResponseLef4sow9ResultExcludeSamlGroupRule]
type singleResponseLef4sow9ResultExcludeSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultExcludeSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultExcludeSamlGroupRule) implementsSingleResponseLef4sow9ResultExclude() {
}

type SingleResponseLef4sow9ResultExcludeSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                   `json:"attribute_value,required"`
	JSON           singleResponseLef4sow9ResultExcludeSamlGroupRuleSamlJSON `json:"-"`
}

// singleResponseLef4sow9ResultExcludeSamlGroupRuleSamlJSON contains the JSON
// metadata for the struct [SingleResponseLef4sow9ResultExcludeSamlGroupRuleSaml]
type singleResponseLef4sow9ResultExcludeSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultExcludeSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [SingleResponseLef4sow9ResultIncludeEmailRule],
// [SingleResponseLef4sow9ResultIncludeDomainRule],
// [SingleResponseLef4sow9ResultIncludeEveryoneRule],
// [SingleResponseLef4sow9ResultIncludeIPRule],
// [SingleResponseLef4sow9ResultIncludeIPListRule],
// [SingleResponseLef4sow9ResultIncludeCertificateRule],
// [SingleResponseLef4sow9ResultIncludeAccessGroupRule],
// [SingleResponseLef4sow9ResultIncludeAzureGroupRule],
// [SingleResponseLef4sow9ResultIncludeGitHubOrganizationRule],
// [SingleResponseLef4sow9ResultIncludeGsuiteGroupRule],
// [SingleResponseLef4sow9ResultIncludeOktaGroupRule] or
// [SingleResponseLef4sow9ResultIncludeSamlGroupRule].
type SingleResponseLef4sow9ResultInclude interface {
	implementsSingleResponseLef4sow9ResultInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*SingleResponseLef4sow9ResultInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type SingleResponseLef4sow9ResultIncludeEmailRule struct {
	Email SingleResponseLef4sow9ResultIncludeEmailRuleEmail `json:"email,required"`
	JSON  singleResponseLef4sow9ResultIncludeEmailRuleJSON  `json:"-"`
}

// singleResponseLef4sow9ResultIncludeEmailRuleJSON contains the JSON metadata for
// the struct [SingleResponseLef4sow9ResultIncludeEmailRule]
type singleResponseLef4sow9ResultIncludeEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultIncludeEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultIncludeEmailRule) implementsSingleResponseLef4sow9ResultInclude() {
}

type SingleResponseLef4sow9ResultIncludeEmailRuleEmail struct {
	// The email of the user.
	Email string                                                `json:"email,required" format:"email"`
	JSON  singleResponseLef4sow9ResultIncludeEmailRuleEmailJSON `json:"-"`
}

// singleResponseLef4sow9ResultIncludeEmailRuleEmailJSON contains the JSON metadata
// for the struct [SingleResponseLef4sow9ResultIncludeEmailRuleEmail]
type singleResponseLef4sow9ResultIncludeEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultIncludeEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type SingleResponseLef4sow9ResultIncludeDomainRule struct {
	EmailDomain SingleResponseLef4sow9ResultIncludeDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        singleResponseLef4sow9ResultIncludeDomainRuleJSON        `json:"-"`
}

// singleResponseLef4sow9ResultIncludeDomainRuleJSON contains the JSON metadata for
// the struct [SingleResponseLef4sow9ResultIncludeDomainRule]
type singleResponseLef4sow9ResultIncludeDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultIncludeDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultIncludeDomainRule) implementsSingleResponseLef4sow9ResultInclude() {
}

type SingleResponseLef4sow9ResultIncludeDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                       `json:"domain,required"`
	JSON   singleResponseLef4sow9ResultIncludeDomainRuleEmailDomainJSON `json:"-"`
}

// singleResponseLef4sow9ResultIncludeDomainRuleEmailDomainJSON contains the JSON
// metadata for the struct
// [SingleResponseLef4sow9ResultIncludeDomainRuleEmailDomain]
type singleResponseLef4sow9ResultIncludeDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultIncludeDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type SingleResponseLef4sow9ResultIncludeEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                         `json:"everyone,required"`
	JSON     singleResponseLef4sow9ResultIncludeEveryoneRuleJSON `json:"-"`
}

// singleResponseLef4sow9ResultIncludeEveryoneRuleJSON contains the JSON metadata
// for the struct [SingleResponseLef4sow9ResultIncludeEveryoneRule]
type singleResponseLef4sow9ResultIncludeEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultIncludeEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultIncludeEveryoneRule) implementsSingleResponseLef4sow9ResultInclude() {
}

// Matches an IP address block.
type SingleResponseLef4sow9ResultIncludeIPRule struct {
	IP   SingleResponseLef4sow9ResultIncludeIPRuleIP   `json:"ip,required"`
	JSON singleResponseLef4sow9ResultIncludeIPRuleJSON `json:"-"`
}

// singleResponseLef4sow9ResultIncludeIPRuleJSON contains the JSON metadata for the
// struct [SingleResponseLef4sow9ResultIncludeIPRule]
type singleResponseLef4sow9ResultIncludeIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultIncludeIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultIncludeIPRule) implementsSingleResponseLef4sow9ResultInclude() {}

type SingleResponseLef4sow9ResultIncludeIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                          `json:"ip,required"`
	JSON singleResponseLef4sow9ResultIncludeIPRuleIPJSON `json:"-"`
}

// singleResponseLef4sow9ResultIncludeIPRuleIPJSON contains the JSON metadata for
// the struct [SingleResponseLef4sow9ResultIncludeIPRuleIP]
type singleResponseLef4sow9ResultIncludeIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultIncludeIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type SingleResponseLef4sow9ResultIncludeIPListRule struct {
	IPList SingleResponseLef4sow9ResultIncludeIPListRuleIPList `json:"ip_list,required"`
	JSON   singleResponseLef4sow9ResultIncludeIPListRuleJSON   `json:"-"`
}

// singleResponseLef4sow9ResultIncludeIPListRuleJSON contains the JSON metadata for
// the struct [SingleResponseLef4sow9ResultIncludeIPListRule]
type singleResponseLef4sow9ResultIncludeIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultIncludeIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultIncludeIPListRule) implementsSingleResponseLef4sow9ResultInclude() {
}

type SingleResponseLef4sow9ResultIncludeIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                  `json:"id,required"`
	JSON singleResponseLef4sow9ResultIncludeIPListRuleIPListJSON `json:"-"`
}

// singleResponseLef4sow9ResultIncludeIPListRuleIPListJSON contains the JSON
// metadata for the struct [SingleResponseLef4sow9ResultIncludeIPListRuleIPList]
type singleResponseLef4sow9ResultIncludeIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultIncludeIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type SingleResponseLef4sow9ResultIncludeCertificateRule struct {
	Certificate interface{}                                            `json:"certificate,required"`
	JSON        singleResponseLef4sow9ResultIncludeCertificateRuleJSON `json:"-"`
}

// singleResponseLef4sow9ResultIncludeCertificateRuleJSON contains the JSON
// metadata for the struct [SingleResponseLef4sow9ResultIncludeCertificateRule]
type singleResponseLef4sow9ResultIncludeCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultIncludeCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultIncludeCertificateRule) implementsSingleResponseLef4sow9ResultInclude() {
}

// Matches an Access group.
type SingleResponseLef4sow9ResultIncludeAccessGroupRule struct {
	Group SingleResponseLef4sow9ResultIncludeAccessGroupRuleGroup `json:"group,required"`
	JSON  singleResponseLef4sow9ResultIncludeAccessGroupRuleJSON  `json:"-"`
}

// singleResponseLef4sow9ResultIncludeAccessGroupRuleJSON contains the JSON
// metadata for the struct [SingleResponseLef4sow9ResultIncludeAccessGroupRule]
type singleResponseLef4sow9ResultIncludeAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultIncludeAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultIncludeAccessGroupRule) implementsSingleResponseLef4sow9ResultInclude() {
}

type SingleResponseLef4sow9ResultIncludeAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                      `json:"id,required"`
	JSON singleResponseLef4sow9ResultIncludeAccessGroupRuleGroupJSON `json:"-"`
}

// singleResponseLef4sow9ResultIncludeAccessGroupRuleGroupJSON contains the JSON
// metadata for the struct
// [SingleResponseLef4sow9ResultIncludeAccessGroupRuleGroup]
type singleResponseLef4sow9ResultIncludeAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultIncludeAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type SingleResponseLef4sow9ResultIncludeAzureGroupRule struct {
	AzureAd SingleResponseLef4sow9ResultIncludeAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    singleResponseLef4sow9ResultIncludeAzureGroupRuleJSON    `json:"-"`
}

// singleResponseLef4sow9ResultIncludeAzureGroupRuleJSON contains the JSON metadata
// for the struct [SingleResponseLef4sow9ResultIncludeAzureGroupRule]
type singleResponseLef4sow9ResultIncludeAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultIncludeAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultIncludeAzureGroupRule) implementsSingleResponseLef4sow9ResultInclude() {
}

type SingleResponseLef4sow9ResultIncludeAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                       `json:"connection_id,required"`
	JSON         singleResponseLef4sow9ResultIncludeAzureGroupRuleAzureAdJSON `json:"-"`
}

// singleResponseLef4sow9ResultIncludeAzureGroupRuleAzureAdJSON contains the JSON
// metadata for the struct
// [SingleResponseLef4sow9ResultIncludeAzureGroupRuleAzureAd]
type singleResponseLef4sow9ResultIncludeAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultIncludeAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type SingleResponseLef4sow9ResultIncludeGitHubOrganizationRule struct {
	GitHubOrganization SingleResponseLef4sow9ResultIncludeGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               singleResponseLef4sow9ResultIncludeGitHubOrganizationRuleJSON               `json:"-"`
}

// singleResponseLef4sow9ResultIncludeGitHubOrganizationRuleJSON contains the JSON
// metadata for the struct
// [SingleResponseLef4sow9ResultIncludeGitHubOrganizationRule]
type singleResponseLef4sow9ResultIncludeGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultIncludeGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultIncludeGitHubOrganizationRule) implementsSingleResponseLef4sow9ResultInclude() {
}

type SingleResponseLef4sow9ResultIncludeGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                          `json:"name,required"`
	JSON singleResponseLef4sow9ResultIncludeGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// singleResponseLef4sow9ResultIncludeGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [SingleResponseLef4sow9ResultIncludeGitHubOrganizationRuleGitHubOrganization]
type singleResponseLef4sow9ResultIncludeGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultIncludeGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type SingleResponseLef4sow9ResultIncludeGsuiteGroupRule struct {
	Gsuite SingleResponseLef4sow9ResultIncludeGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   singleResponseLef4sow9ResultIncludeGsuiteGroupRuleJSON   `json:"-"`
}

// singleResponseLef4sow9ResultIncludeGsuiteGroupRuleJSON contains the JSON
// metadata for the struct [SingleResponseLef4sow9ResultIncludeGsuiteGroupRule]
type singleResponseLef4sow9ResultIncludeGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultIncludeGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultIncludeGsuiteGroupRule) implementsSingleResponseLef4sow9ResultInclude() {
}

type SingleResponseLef4sow9ResultIncludeGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                       `json:"email,required"`
	JSON  singleResponseLef4sow9ResultIncludeGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// singleResponseLef4sow9ResultIncludeGsuiteGroupRuleGsuiteJSON contains the JSON
// metadata for the struct
// [SingleResponseLef4sow9ResultIncludeGsuiteGroupRuleGsuite]
type singleResponseLef4sow9ResultIncludeGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultIncludeGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type SingleResponseLef4sow9ResultIncludeOktaGroupRule struct {
	Okta SingleResponseLef4sow9ResultIncludeOktaGroupRuleOkta `json:"okta,required"`
	JSON singleResponseLef4sow9ResultIncludeOktaGroupRuleJSON `json:"-"`
}

// singleResponseLef4sow9ResultIncludeOktaGroupRuleJSON contains the JSON metadata
// for the struct [SingleResponseLef4sow9ResultIncludeOktaGroupRule]
type singleResponseLef4sow9ResultIncludeOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultIncludeOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultIncludeOktaGroupRule) implementsSingleResponseLef4sow9ResultInclude() {
}

type SingleResponseLef4sow9ResultIncludeOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                   `json:"email,required"`
	JSON  singleResponseLef4sow9ResultIncludeOktaGroupRuleOktaJSON `json:"-"`
}

// singleResponseLef4sow9ResultIncludeOktaGroupRuleOktaJSON contains the JSON
// metadata for the struct [SingleResponseLef4sow9ResultIncludeOktaGroupRuleOkta]
type singleResponseLef4sow9ResultIncludeOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultIncludeOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type SingleResponseLef4sow9ResultIncludeSamlGroupRule struct {
	Saml SingleResponseLef4sow9ResultIncludeSamlGroupRuleSaml `json:"saml,required"`
	JSON singleResponseLef4sow9ResultIncludeSamlGroupRuleJSON `json:"-"`
}

// singleResponseLef4sow9ResultIncludeSamlGroupRuleJSON contains the JSON metadata
// for the struct [SingleResponseLef4sow9ResultIncludeSamlGroupRule]
type singleResponseLef4sow9ResultIncludeSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultIncludeSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultIncludeSamlGroupRule) implementsSingleResponseLef4sow9ResultInclude() {
}

type SingleResponseLef4sow9ResultIncludeSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                   `json:"attribute_value,required"`
	JSON           singleResponseLef4sow9ResultIncludeSamlGroupRuleSamlJSON `json:"-"`
}

// singleResponseLef4sow9ResultIncludeSamlGroupRuleSamlJSON contains the JSON
// metadata for the struct [SingleResponseLef4sow9ResultIncludeSamlGroupRuleSaml]
type singleResponseLef4sow9ResultIncludeSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultIncludeSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [SingleResponseLef4sow9ResultRequireEmailRule],
// [SingleResponseLef4sow9ResultRequireDomainRule],
// [SingleResponseLef4sow9ResultRequireEveryoneRule],
// [SingleResponseLef4sow9ResultRequireIPRule],
// [SingleResponseLef4sow9ResultRequireIPListRule],
// [SingleResponseLef4sow9ResultRequireCertificateRule],
// [SingleResponseLef4sow9ResultRequireAccessGroupRule],
// [SingleResponseLef4sow9ResultRequireAzureGroupRule],
// [SingleResponseLef4sow9ResultRequireGitHubOrganizationRule],
// [SingleResponseLef4sow9ResultRequireGsuiteGroupRule],
// [SingleResponseLef4sow9ResultRequireOktaGroupRule] or
// [SingleResponseLef4sow9ResultRequireSamlGroupRule].
type SingleResponseLef4sow9ResultRequire interface {
	implementsSingleResponseLef4sow9ResultRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*SingleResponseLef4sow9ResultRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type SingleResponseLef4sow9ResultRequireEmailRule struct {
	Email SingleResponseLef4sow9ResultRequireEmailRuleEmail `json:"email,required"`
	JSON  singleResponseLef4sow9ResultRequireEmailRuleJSON  `json:"-"`
}

// singleResponseLef4sow9ResultRequireEmailRuleJSON contains the JSON metadata for
// the struct [SingleResponseLef4sow9ResultRequireEmailRule]
type singleResponseLef4sow9ResultRequireEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultRequireEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultRequireEmailRule) implementsSingleResponseLef4sow9ResultRequire() {
}

type SingleResponseLef4sow9ResultRequireEmailRuleEmail struct {
	// The email of the user.
	Email string                                                `json:"email,required" format:"email"`
	JSON  singleResponseLef4sow9ResultRequireEmailRuleEmailJSON `json:"-"`
}

// singleResponseLef4sow9ResultRequireEmailRuleEmailJSON contains the JSON metadata
// for the struct [SingleResponseLef4sow9ResultRequireEmailRuleEmail]
type singleResponseLef4sow9ResultRequireEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultRequireEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type SingleResponseLef4sow9ResultRequireDomainRule struct {
	EmailDomain SingleResponseLef4sow9ResultRequireDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        singleResponseLef4sow9ResultRequireDomainRuleJSON        `json:"-"`
}

// singleResponseLef4sow9ResultRequireDomainRuleJSON contains the JSON metadata for
// the struct [SingleResponseLef4sow9ResultRequireDomainRule]
type singleResponseLef4sow9ResultRequireDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultRequireDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultRequireDomainRule) implementsSingleResponseLef4sow9ResultRequire() {
}

type SingleResponseLef4sow9ResultRequireDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                       `json:"domain,required"`
	JSON   singleResponseLef4sow9ResultRequireDomainRuleEmailDomainJSON `json:"-"`
}

// singleResponseLef4sow9ResultRequireDomainRuleEmailDomainJSON contains the JSON
// metadata for the struct
// [SingleResponseLef4sow9ResultRequireDomainRuleEmailDomain]
type singleResponseLef4sow9ResultRequireDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultRequireDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type SingleResponseLef4sow9ResultRequireEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                         `json:"everyone,required"`
	JSON     singleResponseLef4sow9ResultRequireEveryoneRuleJSON `json:"-"`
}

// singleResponseLef4sow9ResultRequireEveryoneRuleJSON contains the JSON metadata
// for the struct [SingleResponseLef4sow9ResultRequireEveryoneRule]
type singleResponseLef4sow9ResultRequireEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultRequireEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultRequireEveryoneRule) implementsSingleResponseLef4sow9ResultRequire() {
}

// Matches an IP address block.
type SingleResponseLef4sow9ResultRequireIPRule struct {
	IP   SingleResponseLef4sow9ResultRequireIPRuleIP   `json:"ip,required"`
	JSON singleResponseLef4sow9ResultRequireIPRuleJSON `json:"-"`
}

// singleResponseLef4sow9ResultRequireIPRuleJSON contains the JSON metadata for the
// struct [SingleResponseLef4sow9ResultRequireIPRule]
type singleResponseLef4sow9ResultRequireIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultRequireIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultRequireIPRule) implementsSingleResponseLef4sow9ResultRequire() {}

type SingleResponseLef4sow9ResultRequireIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                          `json:"ip,required"`
	JSON singleResponseLef4sow9ResultRequireIPRuleIPJSON `json:"-"`
}

// singleResponseLef4sow9ResultRequireIPRuleIPJSON contains the JSON metadata for
// the struct [SingleResponseLef4sow9ResultRequireIPRuleIP]
type singleResponseLef4sow9ResultRequireIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultRequireIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type SingleResponseLef4sow9ResultRequireIPListRule struct {
	IPList SingleResponseLef4sow9ResultRequireIPListRuleIPList `json:"ip_list,required"`
	JSON   singleResponseLef4sow9ResultRequireIPListRuleJSON   `json:"-"`
}

// singleResponseLef4sow9ResultRequireIPListRuleJSON contains the JSON metadata for
// the struct [SingleResponseLef4sow9ResultRequireIPListRule]
type singleResponseLef4sow9ResultRequireIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultRequireIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultRequireIPListRule) implementsSingleResponseLef4sow9ResultRequire() {
}

type SingleResponseLef4sow9ResultRequireIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                  `json:"id,required"`
	JSON singleResponseLef4sow9ResultRequireIPListRuleIPListJSON `json:"-"`
}

// singleResponseLef4sow9ResultRequireIPListRuleIPListJSON contains the JSON
// metadata for the struct [SingleResponseLef4sow9ResultRequireIPListRuleIPList]
type singleResponseLef4sow9ResultRequireIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultRequireIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type SingleResponseLef4sow9ResultRequireCertificateRule struct {
	Certificate interface{}                                            `json:"certificate,required"`
	JSON        singleResponseLef4sow9ResultRequireCertificateRuleJSON `json:"-"`
}

// singleResponseLef4sow9ResultRequireCertificateRuleJSON contains the JSON
// metadata for the struct [SingleResponseLef4sow9ResultRequireCertificateRule]
type singleResponseLef4sow9ResultRequireCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultRequireCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultRequireCertificateRule) implementsSingleResponseLef4sow9ResultRequire() {
}

// Matches an Access group.
type SingleResponseLef4sow9ResultRequireAccessGroupRule struct {
	Group SingleResponseLef4sow9ResultRequireAccessGroupRuleGroup `json:"group,required"`
	JSON  singleResponseLef4sow9ResultRequireAccessGroupRuleJSON  `json:"-"`
}

// singleResponseLef4sow9ResultRequireAccessGroupRuleJSON contains the JSON
// metadata for the struct [SingleResponseLef4sow9ResultRequireAccessGroupRule]
type singleResponseLef4sow9ResultRequireAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultRequireAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultRequireAccessGroupRule) implementsSingleResponseLef4sow9ResultRequire() {
}

type SingleResponseLef4sow9ResultRequireAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                      `json:"id,required"`
	JSON singleResponseLef4sow9ResultRequireAccessGroupRuleGroupJSON `json:"-"`
}

// singleResponseLef4sow9ResultRequireAccessGroupRuleGroupJSON contains the JSON
// metadata for the struct
// [SingleResponseLef4sow9ResultRequireAccessGroupRuleGroup]
type singleResponseLef4sow9ResultRequireAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultRequireAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type SingleResponseLef4sow9ResultRequireAzureGroupRule struct {
	AzureAd SingleResponseLef4sow9ResultRequireAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    singleResponseLef4sow9ResultRequireAzureGroupRuleJSON    `json:"-"`
}

// singleResponseLef4sow9ResultRequireAzureGroupRuleJSON contains the JSON metadata
// for the struct [SingleResponseLef4sow9ResultRequireAzureGroupRule]
type singleResponseLef4sow9ResultRequireAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultRequireAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultRequireAzureGroupRule) implementsSingleResponseLef4sow9ResultRequire() {
}

type SingleResponseLef4sow9ResultRequireAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                       `json:"connection_id,required"`
	JSON         singleResponseLef4sow9ResultRequireAzureGroupRuleAzureAdJSON `json:"-"`
}

// singleResponseLef4sow9ResultRequireAzureGroupRuleAzureAdJSON contains the JSON
// metadata for the struct
// [SingleResponseLef4sow9ResultRequireAzureGroupRuleAzureAd]
type singleResponseLef4sow9ResultRequireAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultRequireAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type SingleResponseLef4sow9ResultRequireGitHubOrganizationRule struct {
	GitHubOrganization SingleResponseLef4sow9ResultRequireGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               singleResponseLef4sow9ResultRequireGitHubOrganizationRuleJSON               `json:"-"`
}

// singleResponseLef4sow9ResultRequireGitHubOrganizationRuleJSON contains the JSON
// metadata for the struct
// [SingleResponseLef4sow9ResultRequireGitHubOrganizationRule]
type singleResponseLef4sow9ResultRequireGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultRequireGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultRequireGitHubOrganizationRule) implementsSingleResponseLef4sow9ResultRequire() {
}

type SingleResponseLef4sow9ResultRequireGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                          `json:"name,required"`
	JSON singleResponseLef4sow9ResultRequireGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// singleResponseLef4sow9ResultRequireGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [SingleResponseLef4sow9ResultRequireGitHubOrganizationRuleGitHubOrganization]
type singleResponseLef4sow9ResultRequireGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultRequireGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type SingleResponseLef4sow9ResultRequireGsuiteGroupRule struct {
	Gsuite SingleResponseLef4sow9ResultRequireGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   singleResponseLef4sow9ResultRequireGsuiteGroupRuleJSON   `json:"-"`
}

// singleResponseLef4sow9ResultRequireGsuiteGroupRuleJSON contains the JSON
// metadata for the struct [SingleResponseLef4sow9ResultRequireGsuiteGroupRule]
type singleResponseLef4sow9ResultRequireGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultRequireGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultRequireGsuiteGroupRule) implementsSingleResponseLef4sow9ResultRequire() {
}

type SingleResponseLef4sow9ResultRequireGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                       `json:"email,required"`
	JSON  singleResponseLef4sow9ResultRequireGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// singleResponseLef4sow9ResultRequireGsuiteGroupRuleGsuiteJSON contains the JSON
// metadata for the struct
// [SingleResponseLef4sow9ResultRequireGsuiteGroupRuleGsuite]
type singleResponseLef4sow9ResultRequireGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultRequireGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type SingleResponseLef4sow9ResultRequireOktaGroupRule struct {
	Okta SingleResponseLef4sow9ResultRequireOktaGroupRuleOkta `json:"okta,required"`
	JSON singleResponseLef4sow9ResultRequireOktaGroupRuleJSON `json:"-"`
}

// singleResponseLef4sow9ResultRequireOktaGroupRuleJSON contains the JSON metadata
// for the struct [SingleResponseLef4sow9ResultRequireOktaGroupRule]
type singleResponseLef4sow9ResultRequireOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultRequireOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultRequireOktaGroupRule) implementsSingleResponseLef4sow9ResultRequire() {
}

type SingleResponseLef4sow9ResultRequireOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                   `json:"email,required"`
	JSON  singleResponseLef4sow9ResultRequireOktaGroupRuleOktaJSON `json:"-"`
}

// singleResponseLef4sow9ResultRequireOktaGroupRuleOktaJSON contains the JSON
// metadata for the struct [SingleResponseLef4sow9ResultRequireOktaGroupRuleOkta]
type singleResponseLef4sow9ResultRequireOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultRequireOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type SingleResponseLef4sow9ResultRequireSamlGroupRule struct {
	Saml SingleResponseLef4sow9ResultRequireSamlGroupRuleSaml `json:"saml,required"`
	JSON singleResponseLef4sow9ResultRequireSamlGroupRuleJSON `json:"-"`
}

// singleResponseLef4sow9ResultRequireSamlGroupRuleJSON contains the JSON metadata
// for the struct [SingleResponseLef4sow9ResultRequireSamlGroupRule]
type singleResponseLef4sow9ResultRequireSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultRequireSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SingleResponseLef4sow9ResultRequireSamlGroupRule) implementsSingleResponseLef4sow9ResultRequire() {
}

type SingleResponseLef4sow9ResultRequireSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                   `json:"attribute_value,required"`
	JSON           singleResponseLef4sow9ResultRequireSamlGroupRuleSamlJSON `json:"-"`
}

// singleResponseLef4sow9ResultRequireSamlGroupRuleSamlJSON contains the JSON
// metadata for the struct [SingleResponseLef4sow9ResultRequireSamlGroupRuleSaml]
type singleResponseLef4sow9ResultRequireSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *SingleResponseLef4sow9ResultRequireSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SingleResponseLef4sow9Success bool

const (
	SingleResponseLef4sow9SuccessTrue SingleResponseLef4sow9Success = true
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
	// The unique identifier for the Access group.
	ID   interface{}                                `json:"id"`
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

type AccountAccessGroupUpdateParams struct {
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccountAccessGroupUpdateParamsInclude] `json:"include,required"`
	// The name of the Access group.
	Name param.Field[string] `json:"name,required"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccountAccessGroupUpdateParamsExclude] `json:"exclude"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccountAccessGroupUpdateParamsRequire] `json:"require"`
}

func (r AccountAccessGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [AccountAccessGroupUpdateParamsIncludeEmailRule],
// [AccountAccessGroupUpdateParamsIncludeDomainRule],
// [AccountAccessGroupUpdateParamsIncludeEveryoneRule],
// [AccountAccessGroupUpdateParamsIncludeIPRule],
// [AccountAccessGroupUpdateParamsIncludeIPListRule],
// [AccountAccessGroupUpdateParamsIncludeCertificateRule],
// [AccountAccessGroupUpdateParamsIncludeAccessGroupRule],
// [AccountAccessGroupUpdateParamsIncludeAzureGroupRule],
// [AccountAccessGroupUpdateParamsIncludeGitHubOrganizationRule],
// [AccountAccessGroupUpdateParamsIncludeGsuiteGroupRule],
// [AccountAccessGroupUpdateParamsIncludeOktaGroupRule],
// [AccountAccessGroupUpdateParamsIncludeSamlGroupRule].
type AccountAccessGroupUpdateParamsInclude interface {
	implementsAccountAccessGroupUpdateParamsInclude()
}

// Matches a specific email.
type AccountAccessGroupUpdateParamsIncludeEmailRule struct {
	Email param.Field[AccountAccessGroupUpdateParamsIncludeEmailRuleEmail] `json:"email,required"`
}

func (r AccountAccessGroupUpdateParamsIncludeEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsIncludeEmailRule) implementsAccountAccessGroupUpdateParamsInclude() {
}

type AccountAccessGroupUpdateParamsIncludeEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccountAccessGroupUpdateParamsIncludeEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccountAccessGroupUpdateParamsIncludeDomainRule struct {
	EmailDomain param.Field[AccountAccessGroupUpdateParamsIncludeDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccountAccessGroupUpdateParamsIncludeDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsIncludeDomainRule) implementsAccountAccessGroupUpdateParamsInclude() {
}

type AccountAccessGroupUpdateParamsIncludeDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccountAccessGroupUpdateParamsIncludeDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccountAccessGroupUpdateParamsIncludeEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccountAccessGroupUpdateParamsIncludeEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsIncludeEveryoneRule) implementsAccountAccessGroupUpdateParamsInclude() {
}

// Matches an IP address block.
type AccountAccessGroupUpdateParamsIncludeIPRule struct {
	IP param.Field[AccountAccessGroupUpdateParamsIncludeIPRuleIP] `json:"ip,required"`
}

func (r AccountAccessGroupUpdateParamsIncludeIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsIncludeIPRule) implementsAccountAccessGroupUpdateParamsInclude() {
}

type AccountAccessGroupUpdateParamsIncludeIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccountAccessGroupUpdateParamsIncludeIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccountAccessGroupUpdateParamsIncludeIPListRule struct {
	IPList param.Field[AccountAccessGroupUpdateParamsIncludeIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccountAccessGroupUpdateParamsIncludeIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsIncludeIPListRule) implementsAccountAccessGroupUpdateParamsInclude() {
}

type AccountAccessGroupUpdateParamsIncludeIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccountAccessGroupUpdateParamsIncludeIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccountAccessGroupUpdateParamsIncludeCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccountAccessGroupUpdateParamsIncludeCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsIncludeCertificateRule) implementsAccountAccessGroupUpdateParamsInclude() {
}

// Matches an Access group.
type AccountAccessGroupUpdateParamsIncludeAccessGroupRule struct {
	Group param.Field[AccountAccessGroupUpdateParamsIncludeAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccountAccessGroupUpdateParamsIncludeAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsIncludeAccessGroupRule) implementsAccountAccessGroupUpdateParamsInclude() {
}

type AccountAccessGroupUpdateParamsIncludeAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccountAccessGroupUpdateParamsIncludeAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccountAccessGroupUpdateParamsIncludeAzureGroupRule struct {
	AzureAd param.Field[AccountAccessGroupUpdateParamsIncludeAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccountAccessGroupUpdateParamsIncludeAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsIncludeAzureGroupRule) implementsAccountAccessGroupUpdateParamsInclude() {
}

type AccountAccessGroupUpdateParamsIncludeAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccountAccessGroupUpdateParamsIncludeAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccountAccessGroupUpdateParamsIncludeGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccountAccessGroupUpdateParamsIncludeGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccountAccessGroupUpdateParamsIncludeGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsIncludeGitHubOrganizationRule) implementsAccountAccessGroupUpdateParamsInclude() {
}

type AccountAccessGroupUpdateParamsIncludeGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAccessGroupUpdateParamsIncludeGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccountAccessGroupUpdateParamsIncludeGsuiteGroupRule struct {
	Gsuite param.Field[AccountAccessGroupUpdateParamsIncludeGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccountAccessGroupUpdateParamsIncludeGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsIncludeGsuiteGroupRule) implementsAccountAccessGroupUpdateParamsInclude() {
}

type AccountAccessGroupUpdateParamsIncludeGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccountAccessGroupUpdateParamsIncludeGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccountAccessGroupUpdateParamsIncludeOktaGroupRule struct {
	Okta param.Field[AccountAccessGroupUpdateParamsIncludeOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccountAccessGroupUpdateParamsIncludeOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsIncludeOktaGroupRule) implementsAccountAccessGroupUpdateParamsInclude() {
}

type AccountAccessGroupUpdateParamsIncludeOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccountAccessGroupUpdateParamsIncludeOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccountAccessGroupUpdateParamsIncludeSamlGroupRule struct {
	Saml param.Field[AccountAccessGroupUpdateParamsIncludeSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccountAccessGroupUpdateParamsIncludeSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsIncludeSamlGroupRule) implementsAccountAccessGroupUpdateParamsInclude() {
}

type AccountAccessGroupUpdateParamsIncludeSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccountAccessGroupUpdateParamsIncludeSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [AccountAccessGroupUpdateParamsExcludeEmailRule],
// [AccountAccessGroupUpdateParamsExcludeDomainRule],
// [AccountAccessGroupUpdateParamsExcludeEveryoneRule],
// [AccountAccessGroupUpdateParamsExcludeIPRule],
// [AccountAccessGroupUpdateParamsExcludeIPListRule],
// [AccountAccessGroupUpdateParamsExcludeCertificateRule],
// [AccountAccessGroupUpdateParamsExcludeAccessGroupRule],
// [AccountAccessGroupUpdateParamsExcludeAzureGroupRule],
// [AccountAccessGroupUpdateParamsExcludeGitHubOrganizationRule],
// [AccountAccessGroupUpdateParamsExcludeGsuiteGroupRule],
// [AccountAccessGroupUpdateParamsExcludeOktaGroupRule],
// [AccountAccessGroupUpdateParamsExcludeSamlGroupRule].
type AccountAccessGroupUpdateParamsExclude interface {
	implementsAccountAccessGroupUpdateParamsExclude()
}

// Matches a specific email.
type AccountAccessGroupUpdateParamsExcludeEmailRule struct {
	Email param.Field[AccountAccessGroupUpdateParamsExcludeEmailRuleEmail] `json:"email,required"`
}

func (r AccountAccessGroupUpdateParamsExcludeEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsExcludeEmailRule) implementsAccountAccessGroupUpdateParamsExclude() {
}

type AccountAccessGroupUpdateParamsExcludeEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccountAccessGroupUpdateParamsExcludeEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccountAccessGroupUpdateParamsExcludeDomainRule struct {
	EmailDomain param.Field[AccountAccessGroupUpdateParamsExcludeDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccountAccessGroupUpdateParamsExcludeDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsExcludeDomainRule) implementsAccountAccessGroupUpdateParamsExclude() {
}

type AccountAccessGroupUpdateParamsExcludeDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccountAccessGroupUpdateParamsExcludeDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccountAccessGroupUpdateParamsExcludeEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccountAccessGroupUpdateParamsExcludeEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsExcludeEveryoneRule) implementsAccountAccessGroupUpdateParamsExclude() {
}

// Matches an IP address block.
type AccountAccessGroupUpdateParamsExcludeIPRule struct {
	IP param.Field[AccountAccessGroupUpdateParamsExcludeIPRuleIP] `json:"ip,required"`
}

func (r AccountAccessGroupUpdateParamsExcludeIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsExcludeIPRule) implementsAccountAccessGroupUpdateParamsExclude() {
}

type AccountAccessGroupUpdateParamsExcludeIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccountAccessGroupUpdateParamsExcludeIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccountAccessGroupUpdateParamsExcludeIPListRule struct {
	IPList param.Field[AccountAccessGroupUpdateParamsExcludeIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccountAccessGroupUpdateParamsExcludeIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsExcludeIPListRule) implementsAccountAccessGroupUpdateParamsExclude() {
}

type AccountAccessGroupUpdateParamsExcludeIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccountAccessGroupUpdateParamsExcludeIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccountAccessGroupUpdateParamsExcludeCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccountAccessGroupUpdateParamsExcludeCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsExcludeCertificateRule) implementsAccountAccessGroupUpdateParamsExclude() {
}

// Matches an Access group.
type AccountAccessGroupUpdateParamsExcludeAccessGroupRule struct {
	Group param.Field[AccountAccessGroupUpdateParamsExcludeAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccountAccessGroupUpdateParamsExcludeAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsExcludeAccessGroupRule) implementsAccountAccessGroupUpdateParamsExclude() {
}

type AccountAccessGroupUpdateParamsExcludeAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccountAccessGroupUpdateParamsExcludeAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccountAccessGroupUpdateParamsExcludeAzureGroupRule struct {
	AzureAd param.Field[AccountAccessGroupUpdateParamsExcludeAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccountAccessGroupUpdateParamsExcludeAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsExcludeAzureGroupRule) implementsAccountAccessGroupUpdateParamsExclude() {
}

type AccountAccessGroupUpdateParamsExcludeAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccountAccessGroupUpdateParamsExcludeAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccountAccessGroupUpdateParamsExcludeGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccountAccessGroupUpdateParamsExcludeGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccountAccessGroupUpdateParamsExcludeGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsExcludeGitHubOrganizationRule) implementsAccountAccessGroupUpdateParamsExclude() {
}

type AccountAccessGroupUpdateParamsExcludeGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAccessGroupUpdateParamsExcludeGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccountAccessGroupUpdateParamsExcludeGsuiteGroupRule struct {
	Gsuite param.Field[AccountAccessGroupUpdateParamsExcludeGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccountAccessGroupUpdateParamsExcludeGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsExcludeGsuiteGroupRule) implementsAccountAccessGroupUpdateParamsExclude() {
}

type AccountAccessGroupUpdateParamsExcludeGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccountAccessGroupUpdateParamsExcludeGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccountAccessGroupUpdateParamsExcludeOktaGroupRule struct {
	Okta param.Field[AccountAccessGroupUpdateParamsExcludeOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccountAccessGroupUpdateParamsExcludeOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsExcludeOktaGroupRule) implementsAccountAccessGroupUpdateParamsExclude() {
}

type AccountAccessGroupUpdateParamsExcludeOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccountAccessGroupUpdateParamsExcludeOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccountAccessGroupUpdateParamsExcludeSamlGroupRule struct {
	Saml param.Field[AccountAccessGroupUpdateParamsExcludeSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccountAccessGroupUpdateParamsExcludeSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsExcludeSamlGroupRule) implementsAccountAccessGroupUpdateParamsExclude() {
}

type AccountAccessGroupUpdateParamsExcludeSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccountAccessGroupUpdateParamsExcludeSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [AccountAccessGroupUpdateParamsRequireEmailRule],
// [AccountAccessGroupUpdateParamsRequireDomainRule],
// [AccountAccessGroupUpdateParamsRequireEveryoneRule],
// [AccountAccessGroupUpdateParamsRequireIPRule],
// [AccountAccessGroupUpdateParamsRequireIPListRule],
// [AccountAccessGroupUpdateParamsRequireCertificateRule],
// [AccountAccessGroupUpdateParamsRequireAccessGroupRule],
// [AccountAccessGroupUpdateParamsRequireAzureGroupRule],
// [AccountAccessGroupUpdateParamsRequireGitHubOrganizationRule],
// [AccountAccessGroupUpdateParamsRequireGsuiteGroupRule],
// [AccountAccessGroupUpdateParamsRequireOktaGroupRule],
// [AccountAccessGroupUpdateParamsRequireSamlGroupRule].
type AccountAccessGroupUpdateParamsRequire interface {
	implementsAccountAccessGroupUpdateParamsRequire()
}

// Matches a specific email.
type AccountAccessGroupUpdateParamsRequireEmailRule struct {
	Email param.Field[AccountAccessGroupUpdateParamsRequireEmailRuleEmail] `json:"email,required"`
}

func (r AccountAccessGroupUpdateParamsRequireEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsRequireEmailRule) implementsAccountAccessGroupUpdateParamsRequire() {
}

type AccountAccessGroupUpdateParamsRequireEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccountAccessGroupUpdateParamsRequireEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccountAccessGroupUpdateParamsRequireDomainRule struct {
	EmailDomain param.Field[AccountAccessGroupUpdateParamsRequireDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccountAccessGroupUpdateParamsRequireDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsRequireDomainRule) implementsAccountAccessGroupUpdateParamsRequire() {
}

type AccountAccessGroupUpdateParamsRequireDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccountAccessGroupUpdateParamsRequireDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccountAccessGroupUpdateParamsRequireEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccountAccessGroupUpdateParamsRequireEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsRequireEveryoneRule) implementsAccountAccessGroupUpdateParamsRequire() {
}

// Matches an IP address block.
type AccountAccessGroupUpdateParamsRequireIPRule struct {
	IP param.Field[AccountAccessGroupUpdateParamsRequireIPRuleIP] `json:"ip,required"`
}

func (r AccountAccessGroupUpdateParamsRequireIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsRequireIPRule) implementsAccountAccessGroupUpdateParamsRequire() {
}

type AccountAccessGroupUpdateParamsRequireIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccountAccessGroupUpdateParamsRequireIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccountAccessGroupUpdateParamsRequireIPListRule struct {
	IPList param.Field[AccountAccessGroupUpdateParamsRequireIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccountAccessGroupUpdateParamsRequireIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsRequireIPListRule) implementsAccountAccessGroupUpdateParamsRequire() {
}

type AccountAccessGroupUpdateParamsRequireIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccountAccessGroupUpdateParamsRequireIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccountAccessGroupUpdateParamsRequireCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccountAccessGroupUpdateParamsRequireCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsRequireCertificateRule) implementsAccountAccessGroupUpdateParamsRequire() {
}

// Matches an Access group.
type AccountAccessGroupUpdateParamsRequireAccessGroupRule struct {
	Group param.Field[AccountAccessGroupUpdateParamsRequireAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccountAccessGroupUpdateParamsRequireAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsRequireAccessGroupRule) implementsAccountAccessGroupUpdateParamsRequire() {
}

type AccountAccessGroupUpdateParamsRequireAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccountAccessGroupUpdateParamsRequireAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccountAccessGroupUpdateParamsRequireAzureGroupRule struct {
	AzureAd param.Field[AccountAccessGroupUpdateParamsRequireAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccountAccessGroupUpdateParamsRequireAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsRequireAzureGroupRule) implementsAccountAccessGroupUpdateParamsRequire() {
}

type AccountAccessGroupUpdateParamsRequireAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccountAccessGroupUpdateParamsRequireAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccountAccessGroupUpdateParamsRequireGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccountAccessGroupUpdateParamsRequireGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccountAccessGroupUpdateParamsRequireGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsRequireGitHubOrganizationRule) implementsAccountAccessGroupUpdateParamsRequire() {
}

type AccountAccessGroupUpdateParamsRequireGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAccessGroupUpdateParamsRequireGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccountAccessGroupUpdateParamsRequireGsuiteGroupRule struct {
	Gsuite param.Field[AccountAccessGroupUpdateParamsRequireGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccountAccessGroupUpdateParamsRequireGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsRequireGsuiteGroupRule) implementsAccountAccessGroupUpdateParamsRequire() {
}

type AccountAccessGroupUpdateParamsRequireGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccountAccessGroupUpdateParamsRequireGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccountAccessGroupUpdateParamsRequireOktaGroupRule struct {
	Okta param.Field[AccountAccessGroupUpdateParamsRequireOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccountAccessGroupUpdateParamsRequireOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsRequireOktaGroupRule) implementsAccountAccessGroupUpdateParamsRequire() {
}

type AccountAccessGroupUpdateParamsRequireOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccountAccessGroupUpdateParamsRequireOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccountAccessGroupUpdateParamsRequireSamlGroupRule struct {
	Saml param.Field[AccountAccessGroupUpdateParamsRequireSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccountAccessGroupUpdateParamsRequireSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupUpdateParamsRequireSamlGroupRule) implementsAccountAccessGroupUpdateParamsRequire() {
}

type AccountAccessGroupUpdateParamsRequireSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccountAccessGroupUpdateParamsRequireSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
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
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeEmailRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeDomainRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeEveryoneRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeIPRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeIPListRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeCertificateRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAzureGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeGitHubOrganizationRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeGsuiteGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeOktaGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeSamlGroupRule].
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude interface {
	implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude()
}

// Matches a specific email.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeEmailRule struct {
	Email param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeEmailRuleEmail] `json:"email,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeEmailRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeDomainRule struct {
	EmailDomain param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeDomainRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeEveryoneRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

// Matches an IP address block.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeIPRule struct {
	IP param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeIPRuleIP] `json:"ip,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeIPRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeIPListRule struct {
	IPList param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeIPListRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeCertificateRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

// Matches an Access group.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessGroupRule struct {
	Group param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAzureGroupRule struct {
	AzureAd param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAzureGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeGitHubOrganizationRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeGsuiteGroupRule struct {
	Gsuite param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeGsuiteGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeOktaGroupRule struct {
	Okta param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeOktaGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeSamlGroupRule struct {
	Saml param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeSamlGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsInclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsIncludeSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeEmailRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeDomainRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeEveryoneRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeIPRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeIPListRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeCertificateRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAzureGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeGitHubOrganizationRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeGsuiteGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeOktaGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeSamlGroupRule].
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude interface {
	implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude()
}

// Matches a specific email.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeEmailRule struct {
	Email param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeEmailRuleEmail] `json:"email,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeEmailRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeDomainRule struct {
	EmailDomain param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeDomainRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeEveryoneRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

// Matches an IP address block.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeIPRule struct {
	IP param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeIPRuleIP] `json:"ip,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeIPRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeIPListRule struct {
	IPList param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeIPListRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeCertificateRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

// Matches an Access group.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessGroupRule struct {
	Group param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAzureGroupRule struct {
	AzureAd param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAzureGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeGitHubOrganizationRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeGsuiteGroupRule struct {
	Gsuite param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeGsuiteGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeOktaGroupRule struct {
	Okta param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeOktaGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeSamlGroupRule struct {
	Saml param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeSamlGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsExclude() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsExcludeSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireEmailRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireDomainRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireEveryoneRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireIPRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireIPListRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireCertificateRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireAzureGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireGitHubOrganizationRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireGsuiteGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireOktaGroupRule],
// [AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireSamlGroupRule].
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire interface {
	implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire()
}

// Matches a specific email.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireEmailRule struct {
	Email param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireEmailRuleEmail] `json:"email,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireEmailRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireDomainRule struct {
	EmailDomain param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireDomainRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireEveryoneRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

// Matches an IP address block.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireIPRule struct {
	IP param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireIPRuleIP] `json:"ip,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireIPRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireIPListRule struct {
	IPList param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireIPListRuleIPList] `json:"ip_list,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireIPListRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireCertificateRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

// Matches an Access group.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessGroupRule struct {
	Group param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireAzureGroupRule struct {
	AzureAd param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireAzureGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireGitHubOrganizationRule struct {
	GitHubOrganization param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireGitHubOrganizationRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireGsuiteGroupRule struct {
	Gsuite param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireGsuiteGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireOktaGroupRule struct {
	Okta param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireOktaGroupRuleOkta] `json:"okta,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireOktaGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireSamlGroupRule struct {
	Saml param.Field[AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireSamlGroupRuleSaml] `json:"saml,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireSamlGroupRule) implementsAccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequire() {
}

type AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r AccountAccessGroupAccessGroupsNewAnAccessGroupParamsRequireSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
