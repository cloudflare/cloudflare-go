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
func (r *ZoneAccessGroupService) New(ctx context.Context, identifier string, body ZoneAccessGroupNewParams, opts ...option.RequestOption) (res *ZoneAccessGroupNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/groups", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a single Access group.
func (r *ZoneAccessGroupService) Get(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *ZoneAccessGroupGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/groups/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured Access group.
func (r *ZoneAccessGroupService) Update(ctx context.Context, identifier string, uuid string, body ZoneAccessGroupUpdateParams, opts ...option.RequestOption) (res *ZoneAccessGroupUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/groups/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists all Access groups.
func (r *ZoneAccessGroupService) List(ctx context.Context, identifier string, opts ...option.RequestOption) (res *ZoneAccessGroupListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/groups", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an Access group.
func (r *ZoneAccessGroupService) Delete(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *ZoneAccessGroupDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/groups/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type ZoneAccessGroupNewResponse struct {
	Errors   []ZoneAccessGroupNewResponseError   `json:"errors"`
	Messages []ZoneAccessGroupNewResponseMessage `json:"messages"`
	Result   ZoneAccessGroupNewResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneAccessGroupNewResponseSuccess `json:"success"`
	JSON    zoneAccessGroupNewResponseJSON    `json:"-"`
}

// zoneAccessGroupNewResponseJSON contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponse]
type zoneAccessGroupNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessGroupNewResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    zoneAccessGroupNewResponseErrorJSON `json:"-"`
}

// zoneAccessGroupNewResponseErrorJSON contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseError]
type zoneAccessGroupNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessGroupNewResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zoneAccessGroupNewResponseMessageJSON `json:"-"`
}

// zoneAccessGroupNewResponseMessageJSON contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseMessage]
type zoneAccessGroupNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessGroupNewResponseResult struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []ZoneAccessGroupNewResponseResultExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []ZoneAccessGroupNewResponseResultInclude `json:"include"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []ZoneAccessGroupNewResponseResultRequire `json:"require"`
	UpdatedAt time.Time                                 `json:"updated_at" format:"date-time"`
	JSON      zoneAccessGroupNewResponseResultJSON      `json:"-"`
}

// zoneAccessGroupNewResponseResultJSON contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResult]
type zoneAccessGroupNewResponseResultJSON struct {
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

func (r *ZoneAccessGroupNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [ZoneAccessGroupNewResponseResultExcludePajwohLqEmailRule],
// [ZoneAccessGroupNewResponseResultExcludePajwohLqEmailListRule],
// [ZoneAccessGroupNewResponseResultExcludePajwohLqDomainRule],
// [ZoneAccessGroupNewResponseResultExcludePajwohLqEveryoneRule],
// [ZoneAccessGroupNewResponseResultExcludePajwohLqIPRule],
// [ZoneAccessGroupNewResponseResultExcludePajwohLqIPListRule],
// [ZoneAccessGroupNewResponseResultExcludePajwohLqCertificateRule],
// [ZoneAccessGroupNewResponseResultExcludePajwohLqAccessGroupRule],
// [ZoneAccessGroupNewResponseResultExcludePajwohLqAzureGroupRule],
// [ZoneAccessGroupNewResponseResultExcludePajwohLqGitHubOrganizationRule],
// [ZoneAccessGroupNewResponseResultExcludePajwohLqGsuiteGroupRule],
// [ZoneAccessGroupNewResponseResultExcludePajwohLqOktaGroupRule],
// [ZoneAccessGroupNewResponseResultExcludePajwohLqSamlGroupRule],
// [ZoneAccessGroupNewResponseResultExcludePajwohLqServiceTokenRule],
// [ZoneAccessGroupNewResponseResultExcludePajwohLqAnyValidServiceTokenRule],
// [ZoneAccessGroupNewResponseResultExcludePajwohLqExternalEvaluationRule],
// [ZoneAccessGroupNewResponseResultExcludePajwohLqCountryRule],
// [ZoneAccessGroupNewResponseResultExcludePajwohLqAuthenticationMethodRule] or
// [ZoneAccessGroupNewResponseResultExcludePajwohLqDevicePostureRule].
type ZoneAccessGroupNewResponseResultExclude interface {
	implementsZoneAccessGroupNewResponseResultExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneAccessGroupNewResponseResultExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type ZoneAccessGroupNewResponseResultExcludePajwohLqEmailRule struct {
	Email ZoneAccessGroupNewResponseResultExcludePajwohLqEmailRuleEmail `json:"email,required"`
	JSON  zoneAccessGroupNewResponseResultExcludePajwohLqEmailRuleJSON  `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqEmailRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqEmailRule]
type zoneAccessGroupNewResponseResultExcludePajwohLqEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultExcludePajwohLqEmailRule) implementsZoneAccessGroupNewResponseResultExclude() {
}

type ZoneAccessGroupNewResponseResultExcludePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email string                                                            `json:"email,required" format:"email"`
	JSON  zoneAccessGroupNewResponseResultExcludePajwohLqEmailRuleEmailJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqEmailRuleEmail]
type zoneAccessGroupNewResponseResultExcludePajwohLqEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type ZoneAccessGroupNewResponseResultExcludePajwohLqEmailListRule struct {
	EmailList ZoneAccessGroupNewResponseResultExcludePajwohLqEmailListRuleEmailList `json:"email_list,required"`
	JSON      zoneAccessGroupNewResponseResultExcludePajwohLqEmailListRuleJSON      `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqEmailListRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqEmailListRule]
type zoneAccessGroupNewResponseResultExcludePajwohLqEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultExcludePajwohLqEmailListRule) implementsZoneAccessGroupNewResponseResultExclude() {
}

type ZoneAccessGroupNewResponseResultExcludePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                    `json:"id,required"`
	JSON zoneAccessGroupNewResponseResultExcludePajwohLqEmailListRuleEmailListJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqEmailListRuleEmailList]
type zoneAccessGroupNewResponseResultExcludePajwohLqEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type ZoneAccessGroupNewResponseResultExcludePajwohLqDomainRule struct {
	EmailDomain ZoneAccessGroupNewResponseResultExcludePajwohLqDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zoneAccessGroupNewResponseResultExcludePajwohLqDomainRuleJSON        `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqDomainRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqDomainRule]
type zoneAccessGroupNewResponseResultExcludePajwohLqDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultExcludePajwohLqDomainRule) implementsZoneAccessGroupNewResponseResultExclude() {
}

type ZoneAccessGroupNewResponseResultExcludePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                   `json:"domain,required"`
	JSON   zoneAccessGroupNewResponseResultExcludePajwohLqDomainRuleEmailDomainJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqDomainRuleEmailDomain]
type zoneAccessGroupNewResponseResultExcludePajwohLqDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type ZoneAccessGroupNewResponseResultExcludePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                     `json:"everyone,required"`
	JSON     zoneAccessGroupNewResponseResultExcludePajwohLqEveryoneRuleJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqEveryoneRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqEveryoneRule]
type zoneAccessGroupNewResponseResultExcludePajwohLqEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultExcludePajwohLqEveryoneRule) implementsZoneAccessGroupNewResponseResultExclude() {
}

// Matches an IP address block.
type ZoneAccessGroupNewResponseResultExcludePajwohLqIPRule struct {
	IP   ZoneAccessGroupNewResponseResultExcludePajwohLqIPRuleIP   `json:"ip,required"`
	JSON zoneAccessGroupNewResponseResultExcludePajwohLqIPRuleJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqIPRuleJSON contains the JSON
// metadata for the struct [ZoneAccessGroupNewResponseResultExcludePajwohLqIPRule]
type zoneAccessGroupNewResponseResultExcludePajwohLqIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultExcludePajwohLqIPRule) implementsZoneAccessGroupNewResponseResultExclude() {
}

type ZoneAccessGroupNewResponseResultExcludePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                      `json:"ip,required"`
	JSON zoneAccessGroupNewResponseResultExcludePajwohLqIPRuleIPJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqIPRuleIPJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqIPRuleIP]
type zoneAccessGroupNewResponseResultExcludePajwohLqIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type ZoneAccessGroupNewResponseResultExcludePajwohLqIPListRule struct {
	IPList ZoneAccessGroupNewResponseResultExcludePajwohLqIPListRuleIPList `json:"ip_list,required"`
	JSON   zoneAccessGroupNewResponseResultExcludePajwohLqIPListRuleJSON   `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqIPListRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqIPListRule]
type zoneAccessGroupNewResponseResultExcludePajwohLqIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultExcludePajwohLqIPListRule) implementsZoneAccessGroupNewResponseResultExclude() {
}

type ZoneAccessGroupNewResponseResultExcludePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                              `json:"id,required"`
	JSON zoneAccessGroupNewResponseResultExcludePajwohLqIPListRuleIPListJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqIPListRuleIPList]
type zoneAccessGroupNewResponseResultExcludePajwohLqIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type ZoneAccessGroupNewResponseResultExcludePajwohLqCertificateRule struct {
	Certificate interface{}                                                        `json:"certificate,required"`
	JSON        zoneAccessGroupNewResponseResultExcludePajwohLqCertificateRuleJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqCertificateRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqCertificateRule]
type zoneAccessGroupNewResponseResultExcludePajwohLqCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultExcludePajwohLqCertificateRule) implementsZoneAccessGroupNewResponseResultExclude() {
}

// Matches an Access group.
type ZoneAccessGroupNewResponseResultExcludePajwohLqAccessGroupRule struct {
	Group ZoneAccessGroupNewResponseResultExcludePajwohLqAccessGroupRuleGroup `json:"group,required"`
	JSON  zoneAccessGroupNewResponseResultExcludePajwohLqAccessGroupRuleJSON  `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqAccessGroupRule]
type zoneAccessGroupNewResponseResultExcludePajwohLqAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultExcludePajwohLqAccessGroupRule) implementsZoneAccessGroupNewResponseResultExclude() {
}

type ZoneAccessGroupNewResponseResultExcludePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                  `json:"id,required"`
	JSON zoneAccessGroupNewResponseResultExcludePajwohLqAccessGroupRuleGroupJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqAccessGroupRuleGroup]
type zoneAccessGroupNewResponseResultExcludePajwohLqAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZoneAccessGroupNewResponseResultExcludePajwohLqAzureGroupRule struct {
	AzureAd ZoneAccessGroupNewResponseResultExcludePajwohLqAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zoneAccessGroupNewResponseResultExcludePajwohLqAzureGroupRuleJSON    `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqAzureGroupRule]
type zoneAccessGroupNewResponseResultExcludePajwohLqAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultExcludePajwohLqAzureGroupRule) implementsZoneAccessGroupNewResponseResultExclude() {
}

type ZoneAccessGroupNewResponseResultExcludePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                   `json:"connection_id,required"`
	JSON         zoneAccessGroupNewResponseResultExcludePajwohLqAzureGroupRuleAzureAdJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqAzureGroupRuleAzureAd]
type zoneAccessGroupNewResponseResultExcludePajwohLqAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZoneAccessGroupNewResponseResultExcludePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization ZoneAccessGroupNewResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zoneAccessGroupNewResponseResultExcludePajwohLqGitHubOrganizationRuleJSON               `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqGitHubOrganizationRule]
type zoneAccessGroupNewResponseResultExcludePajwohLqGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultExcludePajwohLqGitHubOrganizationRule) implementsZoneAccessGroupNewResponseResultExclude() {
}

type ZoneAccessGroupNewResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                      `json:"name,required"`
	JSON zoneAccessGroupNewResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization]
type zoneAccessGroupNewResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZoneAccessGroupNewResponseResultExcludePajwohLqGsuiteGroupRule struct {
	Gsuite ZoneAccessGroupNewResponseResultExcludePajwohLqGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zoneAccessGroupNewResponseResultExcludePajwohLqGsuiteGroupRuleJSON   `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqGsuiteGroupRule]
type zoneAccessGroupNewResponseResultExcludePajwohLqGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultExcludePajwohLqGsuiteGroupRule) implementsZoneAccessGroupNewResponseResultExclude() {
}

type ZoneAccessGroupNewResponseResultExcludePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                   `json:"email,required"`
	JSON  zoneAccessGroupNewResponseResultExcludePajwohLqGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqGsuiteGroupRuleGsuite]
type zoneAccessGroupNewResponseResultExcludePajwohLqGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZoneAccessGroupNewResponseResultExcludePajwohLqOktaGroupRule struct {
	Okta ZoneAccessGroupNewResponseResultExcludePajwohLqOktaGroupRuleOkta `json:"okta,required"`
	JSON zoneAccessGroupNewResponseResultExcludePajwohLqOktaGroupRuleJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqOktaGroupRule]
type zoneAccessGroupNewResponseResultExcludePajwohLqOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultExcludePajwohLqOktaGroupRule) implementsZoneAccessGroupNewResponseResultExclude() {
}

type ZoneAccessGroupNewResponseResultExcludePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                               `json:"email,required"`
	JSON  zoneAccessGroupNewResponseResultExcludePajwohLqOktaGroupRuleOktaJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqOktaGroupRuleOkta]
type zoneAccessGroupNewResponseResultExcludePajwohLqOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZoneAccessGroupNewResponseResultExcludePajwohLqSamlGroupRule struct {
	Saml ZoneAccessGroupNewResponseResultExcludePajwohLqSamlGroupRuleSaml `json:"saml,required"`
	JSON zoneAccessGroupNewResponseResultExcludePajwohLqSamlGroupRuleJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqSamlGroupRule]
type zoneAccessGroupNewResponseResultExcludePajwohLqSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultExcludePajwohLqSamlGroupRule) implementsZoneAccessGroupNewResponseResultExclude() {
}

type ZoneAccessGroupNewResponseResultExcludePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                               `json:"attribute_value,required"`
	JSON           zoneAccessGroupNewResponseResultExcludePajwohLqSamlGroupRuleSamlJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqSamlGroupRuleSaml]
type zoneAccessGroupNewResponseResultExcludePajwohLqSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type ZoneAccessGroupNewResponseResultExcludePajwohLqServiceTokenRule struct {
	ServiceToken ZoneAccessGroupNewResponseResultExcludePajwohLqServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zoneAccessGroupNewResponseResultExcludePajwohLqServiceTokenRuleJSON         `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqServiceTokenRule]
type zoneAccessGroupNewResponseResultExcludePajwohLqServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultExcludePajwohLqServiceTokenRule) implementsZoneAccessGroupNewResponseResultExclude() {
}

type ZoneAccessGroupNewResponseResultExcludePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                          `json:"token_id,required"`
	JSON    zoneAccessGroupNewResponseResultExcludePajwohLqServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqServiceTokenRuleServiceToken]
type zoneAccessGroupNewResponseResultExcludePajwohLqServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type ZoneAccessGroupNewResponseResultExcludePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                 `json:"any_valid_service_token,required"`
	JSON                 zoneAccessGroupNewResponseResultExcludePajwohLqAnyValidServiceTokenRuleJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqAnyValidServiceTokenRule]
type zoneAccessGroupNewResponseResultExcludePajwohLqAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultExcludePajwohLqAnyValidServiceTokenRule) implementsZoneAccessGroupNewResponseResultExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZoneAccessGroupNewResponseResultExcludePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation ZoneAccessGroupNewResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zoneAccessGroupNewResponseResultExcludePajwohLqExternalEvaluationRuleJSON               `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqExternalEvaluationRule]
type zoneAccessGroupNewResponseResultExcludePajwohLqExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultExcludePajwohLqExternalEvaluationRule) implementsZoneAccessGroupNewResponseResultExclude() {
}

type ZoneAccessGroupNewResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                      `json:"keys_url,required"`
	JSON    zoneAccessGroupNewResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation]
type zoneAccessGroupNewResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type ZoneAccessGroupNewResponseResultExcludePajwohLqCountryRule struct {
	Geo  ZoneAccessGroupNewResponseResultExcludePajwohLqCountryRuleGeo  `json:"geo,required"`
	JSON zoneAccessGroupNewResponseResultExcludePajwohLqCountryRuleJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqCountryRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqCountryRule]
type zoneAccessGroupNewResponseResultExcludePajwohLqCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultExcludePajwohLqCountryRule) implementsZoneAccessGroupNewResponseResultExclude() {
}

type ZoneAccessGroupNewResponseResultExcludePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                            `json:"country_code,required"`
	JSON        zoneAccessGroupNewResponseResultExcludePajwohLqCountryRuleGeoJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqCountryRuleGeo]
type zoneAccessGroupNewResponseResultExcludePajwohLqCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type ZoneAccessGroupNewResponseResultExcludePajwohLqAuthenticationMethodRule struct {
	AuthMethod ZoneAccessGroupNewResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zoneAccessGroupNewResponseResultExcludePajwohLqAuthenticationMethodRuleJSON       `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqAuthenticationMethodRule]
type zoneAccessGroupNewResponseResultExcludePajwohLqAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultExcludePajwohLqAuthenticationMethodRule) implementsZoneAccessGroupNewResponseResultExclude() {
}

type ZoneAccessGroupNewResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                `json:"auth_method,required"`
	JSON       zoneAccessGroupNewResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod]
type zoneAccessGroupNewResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type ZoneAccessGroupNewResponseResultExcludePajwohLqDevicePostureRule struct {
	DevicePosture ZoneAccessGroupNewResponseResultExcludePajwohLqDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zoneAccessGroupNewResponseResultExcludePajwohLqDevicePostureRuleJSON          `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqDevicePostureRule]
type zoneAccessGroupNewResponseResultExcludePajwohLqDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultExcludePajwohLqDevicePostureRule) implementsZoneAccessGroupNewResponseResultExclude() {
}

type ZoneAccessGroupNewResponseResultExcludePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                            `json:"integration_uid,required"`
	JSON           zoneAccessGroupNewResponseResultExcludePajwohLqDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultExcludePajwohLqDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultExcludePajwohLqDevicePostureRuleDevicePosture]
type zoneAccessGroupNewResponseResultExcludePajwohLqDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultExcludePajwohLqDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [ZoneAccessGroupNewResponseResultIncludePajwohLqEmailRule],
// [ZoneAccessGroupNewResponseResultIncludePajwohLqEmailListRule],
// [ZoneAccessGroupNewResponseResultIncludePajwohLqDomainRule],
// [ZoneAccessGroupNewResponseResultIncludePajwohLqEveryoneRule],
// [ZoneAccessGroupNewResponseResultIncludePajwohLqIPRule],
// [ZoneAccessGroupNewResponseResultIncludePajwohLqIPListRule],
// [ZoneAccessGroupNewResponseResultIncludePajwohLqCertificateRule],
// [ZoneAccessGroupNewResponseResultIncludePajwohLqAccessGroupRule],
// [ZoneAccessGroupNewResponseResultIncludePajwohLqAzureGroupRule],
// [ZoneAccessGroupNewResponseResultIncludePajwohLqGitHubOrganizationRule],
// [ZoneAccessGroupNewResponseResultIncludePajwohLqGsuiteGroupRule],
// [ZoneAccessGroupNewResponseResultIncludePajwohLqOktaGroupRule],
// [ZoneAccessGroupNewResponseResultIncludePajwohLqSamlGroupRule],
// [ZoneAccessGroupNewResponseResultIncludePajwohLqServiceTokenRule],
// [ZoneAccessGroupNewResponseResultIncludePajwohLqAnyValidServiceTokenRule],
// [ZoneAccessGroupNewResponseResultIncludePajwohLqExternalEvaluationRule],
// [ZoneAccessGroupNewResponseResultIncludePajwohLqCountryRule],
// [ZoneAccessGroupNewResponseResultIncludePajwohLqAuthenticationMethodRule] or
// [ZoneAccessGroupNewResponseResultIncludePajwohLqDevicePostureRule].
type ZoneAccessGroupNewResponseResultInclude interface {
	implementsZoneAccessGroupNewResponseResultInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneAccessGroupNewResponseResultInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type ZoneAccessGroupNewResponseResultIncludePajwohLqEmailRule struct {
	Email ZoneAccessGroupNewResponseResultIncludePajwohLqEmailRuleEmail `json:"email,required"`
	JSON  zoneAccessGroupNewResponseResultIncludePajwohLqEmailRuleJSON  `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqEmailRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqEmailRule]
type zoneAccessGroupNewResponseResultIncludePajwohLqEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultIncludePajwohLqEmailRule) implementsZoneAccessGroupNewResponseResultInclude() {
}

type ZoneAccessGroupNewResponseResultIncludePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email string                                                            `json:"email,required" format:"email"`
	JSON  zoneAccessGroupNewResponseResultIncludePajwohLqEmailRuleEmailJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqEmailRuleEmail]
type zoneAccessGroupNewResponseResultIncludePajwohLqEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type ZoneAccessGroupNewResponseResultIncludePajwohLqEmailListRule struct {
	EmailList ZoneAccessGroupNewResponseResultIncludePajwohLqEmailListRuleEmailList `json:"email_list,required"`
	JSON      zoneAccessGroupNewResponseResultIncludePajwohLqEmailListRuleJSON      `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqEmailListRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqEmailListRule]
type zoneAccessGroupNewResponseResultIncludePajwohLqEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultIncludePajwohLqEmailListRule) implementsZoneAccessGroupNewResponseResultInclude() {
}

type ZoneAccessGroupNewResponseResultIncludePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                    `json:"id,required"`
	JSON zoneAccessGroupNewResponseResultIncludePajwohLqEmailListRuleEmailListJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqEmailListRuleEmailList]
type zoneAccessGroupNewResponseResultIncludePajwohLqEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type ZoneAccessGroupNewResponseResultIncludePajwohLqDomainRule struct {
	EmailDomain ZoneAccessGroupNewResponseResultIncludePajwohLqDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zoneAccessGroupNewResponseResultIncludePajwohLqDomainRuleJSON        `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqDomainRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqDomainRule]
type zoneAccessGroupNewResponseResultIncludePajwohLqDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultIncludePajwohLqDomainRule) implementsZoneAccessGroupNewResponseResultInclude() {
}

type ZoneAccessGroupNewResponseResultIncludePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                   `json:"domain,required"`
	JSON   zoneAccessGroupNewResponseResultIncludePajwohLqDomainRuleEmailDomainJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqDomainRuleEmailDomain]
type zoneAccessGroupNewResponseResultIncludePajwohLqDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type ZoneAccessGroupNewResponseResultIncludePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                     `json:"everyone,required"`
	JSON     zoneAccessGroupNewResponseResultIncludePajwohLqEveryoneRuleJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqEveryoneRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqEveryoneRule]
type zoneAccessGroupNewResponseResultIncludePajwohLqEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultIncludePajwohLqEveryoneRule) implementsZoneAccessGroupNewResponseResultInclude() {
}

// Matches an IP address block.
type ZoneAccessGroupNewResponseResultIncludePajwohLqIPRule struct {
	IP   ZoneAccessGroupNewResponseResultIncludePajwohLqIPRuleIP   `json:"ip,required"`
	JSON zoneAccessGroupNewResponseResultIncludePajwohLqIPRuleJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqIPRuleJSON contains the JSON
// metadata for the struct [ZoneAccessGroupNewResponseResultIncludePajwohLqIPRule]
type zoneAccessGroupNewResponseResultIncludePajwohLqIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultIncludePajwohLqIPRule) implementsZoneAccessGroupNewResponseResultInclude() {
}

type ZoneAccessGroupNewResponseResultIncludePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                      `json:"ip,required"`
	JSON zoneAccessGroupNewResponseResultIncludePajwohLqIPRuleIPJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqIPRuleIPJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqIPRuleIP]
type zoneAccessGroupNewResponseResultIncludePajwohLqIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type ZoneAccessGroupNewResponseResultIncludePajwohLqIPListRule struct {
	IPList ZoneAccessGroupNewResponseResultIncludePajwohLqIPListRuleIPList `json:"ip_list,required"`
	JSON   zoneAccessGroupNewResponseResultIncludePajwohLqIPListRuleJSON   `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqIPListRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqIPListRule]
type zoneAccessGroupNewResponseResultIncludePajwohLqIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultIncludePajwohLqIPListRule) implementsZoneAccessGroupNewResponseResultInclude() {
}

type ZoneAccessGroupNewResponseResultIncludePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                              `json:"id,required"`
	JSON zoneAccessGroupNewResponseResultIncludePajwohLqIPListRuleIPListJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqIPListRuleIPList]
type zoneAccessGroupNewResponseResultIncludePajwohLqIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type ZoneAccessGroupNewResponseResultIncludePajwohLqCertificateRule struct {
	Certificate interface{}                                                        `json:"certificate,required"`
	JSON        zoneAccessGroupNewResponseResultIncludePajwohLqCertificateRuleJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqCertificateRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqCertificateRule]
type zoneAccessGroupNewResponseResultIncludePajwohLqCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultIncludePajwohLqCertificateRule) implementsZoneAccessGroupNewResponseResultInclude() {
}

// Matches an Access group.
type ZoneAccessGroupNewResponseResultIncludePajwohLqAccessGroupRule struct {
	Group ZoneAccessGroupNewResponseResultIncludePajwohLqAccessGroupRuleGroup `json:"group,required"`
	JSON  zoneAccessGroupNewResponseResultIncludePajwohLqAccessGroupRuleJSON  `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqAccessGroupRule]
type zoneAccessGroupNewResponseResultIncludePajwohLqAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultIncludePajwohLqAccessGroupRule) implementsZoneAccessGroupNewResponseResultInclude() {
}

type ZoneAccessGroupNewResponseResultIncludePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                  `json:"id,required"`
	JSON zoneAccessGroupNewResponseResultIncludePajwohLqAccessGroupRuleGroupJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqAccessGroupRuleGroup]
type zoneAccessGroupNewResponseResultIncludePajwohLqAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZoneAccessGroupNewResponseResultIncludePajwohLqAzureGroupRule struct {
	AzureAd ZoneAccessGroupNewResponseResultIncludePajwohLqAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zoneAccessGroupNewResponseResultIncludePajwohLqAzureGroupRuleJSON    `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqAzureGroupRule]
type zoneAccessGroupNewResponseResultIncludePajwohLqAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultIncludePajwohLqAzureGroupRule) implementsZoneAccessGroupNewResponseResultInclude() {
}

type ZoneAccessGroupNewResponseResultIncludePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                   `json:"connection_id,required"`
	JSON         zoneAccessGroupNewResponseResultIncludePajwohLqAzureGroupRuleAzureAdJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqAzureGroupRuleAzureAd]
type zoneAccessGroupNewResponseResultIncludePajwohLqAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZoneAccessGroupNewResponseResultIncludePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization ZoneAccessGroupNewResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zoneAccessGroupNewResponseResultIncludePajwohLqGitHubOrganizationRuleJSON               `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqGitHubOrganizationRule]
type zoneAccessGroupNewResponseResultIncludePajwohLqGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultIncludePajwohLqGitHubOrganizationRule) implementsZoneAccessGroupNewResponseResultInclude() {
}

type ZoneAccessGroupNewResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                      `json:"name,required"`
	JSON zoneAccessGroupNewResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization]
type zoneAccessGroupNewResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZoneAccessGroupNewResponseResultIncludePajwohLqGsuiteGroupRule struct {
	Gsuite ZoneAccessGroupNewResponseResultIncludePajwohLqGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zoneAccessGroupNewResponseResultIncludePajwohLqGsuiteGroupRuleJSON   `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqGsuiteGroupRule]
type zoneAccessGroupNewResponseResultIncludePajwohLqGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultIncludePajwohLqGsuiteGroupRule) implementsZoneAccessGroupNewResponseResultInclude() {
}

type ZoneAccessGroupNewResponseResultIncludePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                   `json:"email,required"`
	JSON  zoneAccessGroupNewResponseResultIncludePajwohLqGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqGsuiteGroupRuleGsuite]
type zoneAccessGroupNewResponseResultIncludePajwohLqGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZoneAccessGroupNewResponseResultIncludePajwohLqOktaGroupRule struct {
	Okta ZoneAccessGroupNewResponseResultIncludePajwohLqOktaGroupRuleOkta `json:"okta,required"`
	JSON zoneAccessGroupNewResponseResultIncludePajwohLqOktaGroupRuleJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqOktaGroupRule]
type zoneAccessGroupNewResponseResultIncludePajwohLqOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultIncludePajwohLqOktaGroupRule) implementsZoneAccessGroupNewResponseResultInclude() {
}

type ZoneAccessGroupNewResponseResultIncludePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                               `json:"email,required"`
	JSON  zoneAccessGroupNewResponseResultIncludePajwohLqOktaGroupRuleOktaJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqOktaGroupRuleOkta]
type zoneAccessGroupNewResponseResultIncludePajwohLqOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZoneAccessGroupNewResponseResultIncludePajwohLqSamlGroupRule struct {
	Saml ZoneAccessGroupNewResponseResultIncludePajwohLqSamlGroupRuleSaml `json:"saml,required"`
	JSON zoneAccessGroupNewResponseResultIncludePajwohLqSamlGroupRuleJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqSamlGroupRule]
type zoneAccessGroupNewResponseResultIncludePajwohLqSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultIncludePajwohLqSamlGroupRule) implementsZoneAccessGroupNewResponseResultInclude() {
}

type ZoneAccessGroupNewResponseResultIncludePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                               `json:"attribute_value,required"`
	JSON           zoneAccessGroupNewResponseResultIncludePajwohLqSamlGroupRuleSamlJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqSamlGroupRuleSaml]
type zoneAccessGroupNewResponseResultIncludePajwohLqSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type ZoneAccessGroupNewResponseResultIncludePajwohLqServiceTokenRule struct {
	ServiceToken ZoneAccessGroupNewResponseResultIncludePajwohLqServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zoneAccessGroupNewResponseResultIncludePajwohLqServiceTokenRuleJSON         `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqServiceTokenRule]
type zoneAccessGroupNewResponseResultIncludePajwohLqServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultIncludePajwohLqServiceTokenRule) implementsZoneAccessGroupNewResponseResultInclude() {
}

type ZoneAccessGroupNewResponseResultIncludePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                          `json:"token_id,required"`
	JSON    zoneAccessGroupNewResponseResultIncludePajwohLqServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqServiceTokenRuleServiceToken]
type zoneAccessGroupNewResponseResultIncludePajwohLqServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type ZoneAccessGroupNewResponseResultIncludePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                 `json:"any_valid_service_token,required"`
	JSON                 zoneAccessGroupNewResponseResultIncludePajwohLqAnyValidServiceTokenRuleJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqAnyValidServiceTokenRule]
type zoneAccessGroupNewResponseResultIncludePajwohLqAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultIncludePajwohLqAnyValidServiceTokenRule) implementsZoneAccessGroupNewResponseResultInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZoneAccessGroupNewResponseResultIncludePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation ZoneAccessGroupNewResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zoneAccessGroupNewResponseResultIncludePajwohLqExternalEvaluationRuleJSON               `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqExternalEvaluationRule]
type zoneAccessGroupNewResponseResultIncludePajwohLqExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultIncludePajwohLqExternalEvaluationRule) implementsZoneAccessGroupNewResponseResultInclude() {
}

type ZoneAccessGroupNewResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                      `json:"keys_url,required"`
	JSON    zoneAccessGroupNewResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation]
type zoneAccessGroupNewResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type ZoneAccessGroupNewResponseResultIncludePajwohLqCountryRule struct {
	Geo  ZoneAccessGroupNewResponseResultIncludePajwohLqCountryRuleGeo  `json:"geo,required"`
	JSON zoneAccessGroupNewResponseResultIncludePajwohLqCountryRuleJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqCountryRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqCountryRule]
type zoneAccessGroupNewResponseResultIncludePajwohLqCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultIncludePajwohLqCountryRule) implementsZoneAccessGroupNewResponseResultInclude() {
}

type ZoneAccessGroupNewResponseResultIncludePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                            `json:"country_code,required"`
	JSON        zoneAccessGroupNewResponseResultIncludePajwohLqCountryRuleGeoJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqCountryRuleGeo]
type zoneAccessGroupNewResponseResultIncludePajwohLqCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type ZoneAccessGroupNewResponseResultIncludePajwohLqAuthenticationMethodRule struct {
	AuthMethod ZoneAccessGroupNewResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zoneAccessGroupNewResponseResultIncludePajwohLqAuthenticationMethodRuleJSON       `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqAuthenticationMethodRule]
type zoneAccessGroupNewResponseResultIncludePajwohLqAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultIncludePajwohLqAuthenticationMethodRule) implementsZoneAccessGroupNewResponseResultInclude() {
}

type ZoneAccessGroupNewResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                `json:"auth_method,required"`
	JSON       zoneAccessGroupNewResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod]
type zoneAccessGroupNewResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type ZoneAccessGroupNewResponseResultIncludePajwohLqDevicePostureRule struct {
	DevicePosture ZoneAccessGroupNewResponseResultIncludePajwohLqDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zoneAccessGroupNewResponseResultIncludePajwohLqDevicePostureRuleJSON          `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqDevicePostureRule]
type zoneAccessGroupNewResponseResultIncludePajwohLqDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultIncludePajwohLqDevicePostureRule) implementsZoneAccessGroupNewResponseResultInclude() {
}

type ZoneAccessGroupNewResponseResultIncludePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                            `json:"integration_uid,required"`
	JSON           zoneAccessGroupNewResponseResultIncludePajwohLqDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultIncludePajwohLqDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultIncludePajwohLqDevicePostureRuleDevicePosture]
type zoneAccessGroupNewResponseResultIncludePajwohLqDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultIncludePajwohLqDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [ZoneAccessGroupNewResponseResultRequirePajwohLqEmailRule],
// [ZoneAccessGroupNewResponseResultRequirePajwohLqEmailListRule],
// [ZoneAccessGroupNewResponseResultRequirePajwohLqDomainRule],
// [ZoneAccessGroupNewResponseResultRequirePajwohLqEveryoneRule],
// [ZoneAccessGroupNewResponseResultRequirePajwohLqIPRule],
// [ZoneAccessGroupNewResponseResultRequirePajwohLqIPListRule],
// [ZoneAccessGroupNewResponseResultRequirePajwohLqCertificateRule],
// [ZoneAccessGroupNewResponseResultRequirePajwohLqAccessGroupRule],
// [ZoneAccessGroupNewResponseResultRequirePajwohLqAzureGroupRule],
// [ZoneAccessGroupNewResponseResultRequirePajwohLqGitHubOrganizationRule],
// [ZoneAccessGroupNewResponseResultRequirePajwohLqGsuiteGroupRule],
// [ZoneAccessGroupNewResponseResultRequirePajwohLqOktaGroupRule],
// [ZoneAccessGroupNewResponseResultRequirePajwohLqSamlGroupRule],
// [ZoneAccessGroupNewResponseResultRequirePajwohLqServiceTokenRule],
// [ZoneAccessGroupNewResponseResultRequirePajwohLqAnyValidServiceTokenRule],
// [ZoneAccessGroupNewResponseResultRequirePajwohLqExternalEvaluationRule],
// [ZoneAccessGroupNewResponseResultRequirePajwohLqCountryRule],
// [ZoneAccessGroupNewResponseResultRequirePajwohLqAuthenticationMethodRule] or
// [ZoneAccessGroupNewResponseResultRequirePajwohLqDevicePostureRule].
type ZoneAccessGroupNewResponseResultRequire interface {
	implementsZoneAccessGroupNewResponseResultRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneAccessGroupNewResponseResultRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type ZoneAccessGroupNewResponseResultRequirePajwohLqEmailRule struct {
	Email ZoneAccessGroupNewResponseResultRequirePajwohLqEmailRuleEmail `json:"email,required"`
	JSON  zoneAccessGroupNewResponseResultRequirePajwohLqEmailRuleJSON  `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqEmailRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqEmailRule]
type zoneAccessGroupNewResponseResultRequirePajwohLqEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultRequirePajwohLqEmailRule) implementsZoneAccessGroupNewResponseResultRequire() {
}

type ZoneAccessGroupNewResponseResultRequirePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email string                                                            `json:"email,required" format:"email"`
	JSON  zoneAccessGroupNewResponseResultRequirePajwohLqEmailRuleEmailJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqEmailRuleEmail]
type zoneAccessGroupNewResponseResultRequirePajwohLqEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type ZoneAccessGroupNewResponseResultRequirePajwohLqEmailListRule struct {
	EmailList ZoneAccessGroupNewResponseResultRequirePajwohLqEmailListRuleEmailList `json:"email_list,required"`
	JSON      zoneAccessGroupNewResponseResultRequirePajwohLqEmailListRuleJSON      `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqEmailListRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqEmailListRule]
type zoneAccessGroupNewResponseResultRequirePajwohLqEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultRequirePajwohLqEmailListRule) implementsZoneAccessGroupNewResponseResultRequire() {
}

type ZoneAccessGroupNewResponseResultRequirePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                    `json:"id,required"`
	JSON zoneAccessGroupNewResponseResultRequirePajwohLqEmailListRuleEmailListJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqEmailListRuleEmailList]
type zoneAccessGroupNewResponseResultRequirePajwohLqEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type ZoneAccessGroupNewResponseResultRequirePajwohLqDomainRule struct {
	EmailDomain ZoneAccessGroupNewResponseResultRequirePajwohLqDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zoneAccessGroupNewResponseResultRequirePajwohLqDomainRuleJSON        `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqDomainRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqDomainRule]
type zoneAccessGroupNewResponseResultRequirePajwohLqDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultRequirePajwohLqDomainRule) implementsZoneAccessGroupNewResponseResultRequire() {
}

type ZoneAccessGroupNewResponseResultRequirePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                   `json:"domain,required"`
	JSON   zoneAccessGroupNewResponseResultRequirePajwohLqDomainRuleEmailDomainJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqDomainRuleEmailDomain]
type zoneAccessGroupNewResponseResultRequirePajwohLqDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type ZoneAccessGroupNewResponseResultRequirePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                     `json:"everyone,required"`
	JSON     zoneAccessGroupNewResponseResultRequirePajwohLqEveryoneRuleJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqEveryoneRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqEveryoneRule]
type zoneAccessGroupNewResponseResultRequirePajwohLqEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultRequirePajwohLqEveryoneRule) implementsZoneAccessGroupNewResponseResultRequire() {
}

// Matches an IP address block.
type ZoneAccessGroupNewResponseResultRequirePajwohLqIPRule struct {
	IP   ZoneAccessGroupNewResponseResultRequirePajwohLqIPRuleIP   `json:"ip,required"`
	JSON zoneAccessGroupNewResponseResultRequirePajwohLqIPRuleJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqIPRuleJSON contains the JSON
// metadata for the struct [ZoneAccessGroupNewResponseResultRequirePajwohLqIPRule]
type zoneAccessGroupNewResponseResultRequirePajwohLqIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultRequirePajwohLqIPRule) implementsZoneAccessGroupNewResponseResultRequire() {
}

type ZoneAccessGroupNewResponseResultRequirePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                      `json:"ip,required"`
	JSON zoneAccessGroupNewResponseResultRequirePajwohLqIPRuleIPJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqIPRuleIPJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqIPRuleIP]
type zoneAccessGroupNewResponseResultRequirePajwohLqIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type ZoneAccessGroupNewResponseResultRequirePajwohLqIPListRule struct {
	IPList ZoneAccessGroupNewResponseResultRequirePajwohLqIPListRuleIPList `json:"ip_list,required"`
	JSON   zoneAccessGroupNewResponseResultRequirePajwohLqIPListRuleJSON   `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqIPListRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqIPListRule]
type zoneAccessGroupNewResponseResultRequirePajwohLqIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultRequirePajwohLqIPListRule) implementsZoneAccessGroupNewResponseResultRequire() {
}

type ZoneAccessGroupNewResponseResultRequirePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                              `json:"id,required"`
	JSON zoneAccessGroupNewResponseResultRequirePajwohLqIPListRuleIPListJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqIPListRuleIPList]
type zoneAccessGroupNewResponseResultRequirePajwohLqIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type ZoneAccessGroupNewResponseResultRequirePajwohLqCertificateRule struct {
	Certificate interface{}                                                        `json:"certificate,required"`
	JSON        zoneAccessGroupNewResponseResultRequirePajwohLqCertificateRuleJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqCertificateRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqCertificateRule]
type zoneAccessGroupNewResponseResultRequirePajwohLqCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultRequirePajwohLqCertificateRule) implementsZoneAccessGroupNewResponseResultRequire() {
}

// Matches an Access group.
type ZoneAccessGroupNewResponseResultRequirePajwohLqAccessGroupRule struct {
	Group ZoneAccessGroupNewResponseResultRequirePajwohLqAccessGroupRuleGroup `json:"group,required"`
	JSON  zoneAccessGroupNewResponseResultRequirePajwohLqAccessGroupRuleJSON  `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqAccessGroupRule]
type zoneAccessGroupNewResponseResultRequirePajwohLqAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultRequirePajwohLqAccessGroupRule) implementsZoneAccessGroupNewResponseResultRequire() {
}

type ZoneAccessGroupNewResponseResultRequirePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                  `json:"id,required"`
	JSON zoneAccessGroupNewResponseResultRequirePajwohLqAccessGroupRuleGroupJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqAccessGroupRuleGroup]
type zoneAccessGroupNewResponseResultRequirePajwohLqAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZoneAccessGroupNewResponseResultRequirePajwohLqAzureGroupRule struct {
	AzureAd ZoneAccessGroupNewResponseResultRequirePajwohLqAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zoneAccessGroupNewResponseResultRequirePajwohLqAzureGroupRuleJSON    `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqAzureGroupRule]
type zoneAccessGroupNewResponseResultRequirePajwohLqAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultRequirePajwohLqAzureGroupRule) implementsZoneAccessGroupNewResponseResultRequire() {
}

type ZoneAccessGroupNewResponseResultRequirePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                   `json:"connection_id,required"`
	JSON         zoneAccessGroupNewResponseResultRequirePajwohLqAzureGroupRuleAzureAdJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqAzureGroupRuleAzureAd]
type zoneAccessGroupNewResponseResultRequirePajwohLqAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZoneAccessGroupNewResponseResultRequirePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization ZoneAccessGroupNewResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zoneAccessGroupNewResponseResultRequirePajwohLqGitHubOrganizationRuleJSON               `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqGitHubOrganizationRule]
type zoneAccessGroupNewResponseResultRequirePajwohLqGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultRequirePajwohLqGitHubOrganizationRule) implementsZoneAccessGroupNewResponseResultRequire() {
}

type ZoneAccessGroupNewResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                      `json:"name,required"`
	JSON zoneAccessGroupNewResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization]
type zoneAccessGroupNewResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZoneAccessGroupNewResponseResultRequirePajwohLqGsuiteGroupRule struct {
	Gsuite ZoneAccessGroupNewResponseResultRequirePajwohLqGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zoneAccessGroupNewResponseResultRequirePajwohLqGsuiteGroupRuleJSON   `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqGsuiteGroupRule]
type zoneAccessGroupNewResponseResultRequirePajwohLqGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultRequirePajwohLqGsuiteGroupRule) implementsZoneAccessGroupNewResponseResultRequire() {
}

type ZoneAccessGroupNewResponseResultRequirePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                   `json:"email,required"`
	JSON  zoneAccessGroupNewResponseResultRequirePajwohLqGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqGsuiteGroupRuleGsuite]
type zoneAccessGroupNewResponseResultRequirePajwohLqGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZoneAccessGroupNewResponseResultRequirePajwohLqOktaGroupRule struct {
	Okta ZoneAccessGroupNewResponseResultRequirePajwohLqOktaGroupRuleOkta `json:"okta,required"`
	JSON zoneAccessGroupNewResponseResultRequirePajwohLqOktaGroupRuleJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqOktaGroupRule]
type zoneAccessGroupNewResponseResultRequirePajwohLqOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultRequirePajwohLqOktaGroupRule) implementsZoneAccessGroupNewResponseResultRequire() {
}

type ZoneAccessGroupNewResponseResultRequirePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                               `json:"email,required"`
	JSON  zoneAccessGroupNewResponseResultRequirePajwohLqOktaGroupRuleOktaJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqOktaGroupRuleOkta]
type zoneAccessGroupNewResponseResultRequirePajwohLqOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZoneAccessGroupNewResponseResultRequirePajwohLqSamlGroupRule struct {
	Saml ZoneAccessGroupNewResponseResultRequirePajwohLqSamlGroupRuleSaml `json:"saml,required"`
	JSON zoneAccessGroupNewResponseResultRequirePajwohLqSamlGroupRuleJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqSamlGroupRule]
type zoneAccessGroupNewResponseResultRequirePajwohLqSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultRequirePajwohLqSamlGroupRule) implementsZoneAccessGroupNewResponseResultRequire() {
}

type ZoneAccessGroupNewResponseResultRequirePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                               `json:"attribute_value,required"`
	JSON           zoneAccessGroupNewResponseResultRequirePajwohLqSamlGroupRuleSamlJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqSamlGroupRuleSaml]
type zoneAccessGroupNewResponseResultRequirePajwohLqSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type ZoneAccessGroupNewResponseResultRequirePajwohLqServiceTokenRule struct {
	ServiceToken ZoneAccessGroupNewResponseResultRequirePajwohLqServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zoneAccessGroupNewResponseResultRequirePajwohLqServiceTokenRuleJSON         `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqServiceTokenRule]
type zoneAccessGroupNewResponseResultRequirePajwohLqServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultRequirePajwohLqServiceTokenRule) implementsZoneAccessGroupNewResponseResultRequire() {
}

type ZoneAccessGroupNewResponseResultRequirePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                          `json:"token_id,required"`
	JSON    zoneAccessGroupNewResponseResultRequirePajwohLqServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqServiceTokenRuleServiceToken]
type zoneAccessGroupNewResponseResultRequirePajwohLqServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type ZoneAccessGroupNewResponseResultRequirePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                 `json:"any_valid_service_token,required"`
	JSON                 zoneAccessGroupNewResponseResultRequirePajwohLqAnyValidServiceTokenRuleJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqAnyValidServiceTokenRule]
type zoneAccessGroupNewResponseResultRequirePajwohLqAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultRequirePajwohLqAnyValidServiceTokenRule) implementsZoneAccessGroupNewResponseResultRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZoneAccessGroupNewResponseResultRequirePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation ZoneAccessGroupNewResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zoneAccessGroupNewResponseResultRequirePajwohLqExternalEvaluationRuleJSON               `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqExternalEvaluationRule]
type zoneAccessGroupNewResponseResultRequirePajwohLqExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultRequirePajwohLqExternalEvaluationRule) implementsZoneAccessGroupNewResponseResultRequire() {
}

type ZoneAccessGroupNewResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                      `json:"keys_url,required"`
	JSON    zoneAccessGroupNewResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation]
type zoneAccessGroupNewResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type ZoneAccessGroupNewResponseResultRequirePajwohLqCountryRule struct {
	Geo  ZoneAccessGroupNewResponseResultRequirePajwohLqCountryRuleGeo  `json:"geo,required"`
	JSON zoneAccessGroupNewResponseResultRequirePajwohLqCountryRuleJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqCountryRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqCountryRule]
type zoneAccessGroupNewResponseResultRequirePajwohLqCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultRequirePajwohLqCountryRule) implementsZoneAccessGroupNewResponseResultRequire() {
}

type ZoneAccessGroupNewResponseResultRequirePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                            `json:"country_code,required"`
	JSON        zoneAccessGroupNewResponseResultRequirePajwohLqCountryRuleGeoJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqCountryRuleGeo]
type zoneAccessGroupNewResponseResultRequirePajwohLqCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type ZoneAccessGroupNewResponseResultRequirePajwohLqAuthenticationMethodRule struct {
	AuthMethod ZoneAccessGroupNewResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zoneAccessGroupNewResponseResultRequirePajwohLqAuthenticationMethodRuleJSON       `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqAuthenticationMethodRule]
type zoneAccessGroupNewResponseResultRequirePajwohLqAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultRequirePajwohLqAuthenticationMethodRule) implementsZoneAccessGroupNewResponseResultRequire() {
}

type ZoneAccessGroupNewResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                `json:"auth_method,required"`
	JSON       zoneAccessGroupNewResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod]
type zoneAccessGroupNewResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type ZoneAccessGroupNewResponseResultRequirePajwohLqDevicePostureRule struct {
	DevicePosture ZoneAccessGroupNewResponseResultRequirePajwohLqDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zoneAccessGroupNewResponseResultRequirePajwohLqDevicePostureRuleJSON          `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqDevicePostureRule]
type zoneAccessGroupNewResponseResultRequirePajwohLqDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupNewResponseResultRequirePajwohLqDevicePostureRule) implementsZoneAccessGroupNewResponseResultRequire() {
}

type ZoneAccessGroupNewResponseResultRequirePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                            `json:"integration_uid,required"`
	JSON           zoneAccessGroupNewResponseResultRequirePajwohLqDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zoneAccessGroupNewResponseResultRequirePajwohLqDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupNewResponseResultRequirePajwohLqDevicePostureRuleDevicePosture]
type zoneAccessGroupNewResponseResultRequirePajwohLqDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZoneAccessGroupNewResponseResultRequirePajwohLqDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneAccessGroupNewResponseSuccess bool

const (
	ZoneAccessGroupNewResponseSuccessTrue ZoneAccessGroupNewResponseSuccess = true
)

type ZoneAccessGroupGetResponse struct {
	Errors   []ZoneAccessGroupGetResponseError   `json:"errors"`
	Messages []ZoneAccessGroupGetResponseMessage `json:"messages"`
	Result   ZoneAccessGroupGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneAccessGroupGetResponseSuccess `json:"success"`
	JSON    zoneAccessGroupGetResponseJSON    `json:"-"`
}

// zoneAccessGroupGetResponseJSON contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponse]
type zoneAccessGroupGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessGroupGetResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    zoneAccessGroupGetResponseErrorJSON `json:"-"`
}

// zoneAccessGroupGetResponseErrorJSON contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseError]
type zoneAccessGroupGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessGroupGetResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zoneAccessGroupGetResponseMessageJSON `json:"-"`
}

// zoneAccessGroupGetResponseMessageJSON contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseMessage]
type zoneAccessGroupGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessGroupGetResponseResult struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []ZoneAccessGroupGetResponseResultExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []ZoneAccessGroupGetResponseResultInclude `json:"include"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []ZoneAccessGroupGetResponseResultRequire `json:"require"`
	UpdatedAt time.Time                                 `json:"updated_at" format:"date-time"`
	JSON      zoneAccessGroupGetResponseResultJSON      `json:"-"`
}

// zoneAccessGroupGetResponseResultJSON contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResult]
type zoneAccessGroupGetResponseResultJSON struct {
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

func (r *ZoneAccessGroupGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [ZoneAccessGroupGetResponseResultExcludePajwohLqEmailRule],
// [ZoneAccessGroupGetResponseResultExcludePajwohLqEmailListRule],
// [ZoneAccessGroupGetResponseResultExcludePajwohLqDomainRule],
// [ZoneAccessGroupGetResponseResultExcludePajwohLqEveryoneRule],
// [ZoneAccessGroupGetResponseResultExcludePajwohLqIPRule],
// [ZoneAccessGroupGetResponseResultExcludePajwohLqIPListRule],
// [ZoneAccessGroupGetResponseResultExcludePajwohLqCertificateRule],
// [ZoneAccessGroupGetResponseResultExcludePajwohLqAccessGroupRule],
// [ZoneAccessGroupGetResponseResultExcludePajwohLqAzureGroupRule],
// [ZoneAccessGroupGetResponseResultExcludePajwohLqGitHubOrganizationRule],
// [ZoneAccessGroupGetResponseResultExcludePajwohLqGsuiteGroupRule],
// [ZoneAccessGroupGetResponseResultExcludePajwohLqOktaGroupRule],
// [ZoneAccessGroupGetResponseResultExcludePajwohLqSamlGroupRule],
// [ZoneAccessGroupGetResponseResultExcludePajwohLqServiceTokenRule],
// [ZoneAccessGroupGetResponseResultExcludePajwohLqAnyValidServiceTokenRule],
// [ZoneAccessGroupGetResponseResultExcludePajwohLqExternalEvaluationRule],
// [ZoneAccessGroupGetResponseResultExcludePajwohLqCountryRule],
// [ZoneAccessGroupGetResponseResultExcludePajwohLqAuthenticationMethodRule] or
// [ZoneAccessGroupGetResponseResultExcludePajwohLqDevicePostureRule].
type ZoneAccessGroupGetResponseResultExclude interface {
	implementsZoneAccessGroupGetResponseResultExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneAccessGroupGetResponseResultExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type ZoneAccessGroupGetResponseResultExcludePajwohLqEmailRule struct {
	Email ZoneAccessGroupGetResponseResultExcludePajwohLqEmailRuleEmail `json:"email,required"`
	JSON  zoneAccessGroupGetResponseResultExcludePajwohLqEmailRuleJSON  `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqEmailRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqEmailRule]
type zoneAccessGroupGetResponseResultExcludePajwohLqEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultExcludePajwohLqEmailRule) implementsZoneAccessGroupGetResponseResultExclude() {
}

type ZoneAccessGroupGetResponseResultExcludePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email string                                                            `json:"email,required" format:"email"`
	JSON  zoneAccessGroupGetResponseResultExcludePajwohLqEmailRuleEmailJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqEmailRuleEmail]
type zoneAccessGroupGetResponseResultExcludePajwohLqEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type ZoneAccessGroupGetResponseResultExcludePajwohLqEmailListRule struct {
	EmailList ZoneAccessGroupGetResponseResultExcludePajwohLqEmailListRuleEmailList `json:"email_list,required"`
	JSON      zoneAccessGroupGetResponseResultExcludePajwohLqEmailListRuleJSON      `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqEmailListRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqEmailListRule]
type zoneAccessGroupGetResponseResultExcludePajwohLqEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultExcludePajwohLqEmailListRule) implementsZoneAccessGroupGetResponseResultExclude() {
}

type ZoneAccessGroupGetResponseResultExcludePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                    `json:"id,required"`
	JSON zoneAccessGroupGetResponseResultExcludePajwohLqEmailListRuleEmailListJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqEmailListRuleEmailList]
type zoneAccessGroupGetResponseResultExcludePajwohLqEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type ZoneAccessGroupGetResponseResultExcludePajwohLqDomainRule struct {
	EmailDomain ZoneAccessGroupGetResponseResultExcludePajwohLqDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zoneAccessGroupGetResponseResultExcludePajwohLqDomainRuleJSON        `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqDomainRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqDomainRule]
type zoneAccessGroupGetResponseResultExcludePajwohLqDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultExcludePajwohLqDomainRule) implementsZoneAccessGroupGetResponseResultExclude() {
}

type ZoneAccessGroupGetResponseResultExcludePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                   `json:"domain,required"`
	JSON   zoneAccessGroupGetResponseResultExcludePajwohLqDomainRuleEmailDomainJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqDomainRuleEmailDomain]
type zoneAccessGroupGetResponseResultExcludePajwohLqDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type ZoneAccessGroupGetResponseResultExcludePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                     `json:"everyone,required"`
	JSON     zoneAccessGroupGetResponseResultExcludePajwohLqEveryoneRuleJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqEveryoneRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqEveryoneRule]
type zoneAccessGroupGetResponseResultExcludePajwohLqEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultExcludePajwohLqEveryoneRule) implementsZoneAccessGroupGetResponseResultExclude() {
}

// Matches an IP address block.
type ZoneAccessGroupGetResponseResultExcludePajwohLqIPRule struct {
	IP   ZoneAccessGroupGetResponseResultExcludePajwohLqIPRuleIP   `json:"ip,required"`
	JSON zoneAccessGroupGetResponseResultExcludePajwohLqIPRuleJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqIPRuleJSON contains the JSON
// metadata for the struct [ZoneAccessGroupGetResponseResultExcludePajwohLqIPRule]
type zoneAccessGroupGetResponseResultExcludePajwohLqIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultExcludePajwohLqIPRule) implementsZoneAccessGroupGetResponseResultExclude() {
}

type ZoneAccessGroupGetResponseResultExcludePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                      `json:"ip,required"`
	JSON zoneAccessGroupGetResponseResultExcludePajwohLqIPRuleIPJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqIPRuleIPJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqIPRuleIP]
type zoneAccessGroupGetResponseResultExcludePajwohLqIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type ZoneAccessGroupGetResponseResultExcludePajwohLqIPListRule struct {
	IPList ZoneAccessGroupGetResponseResultExcludePajwohLqIPListRuleIPList `json:"ip_list,required"`
	JSON   zoneAccessGroupGetResponseResultExcludePajwohLqIPListRuleJSON   `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqIPListRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqIPListRule]
type zoneAccessGroupGetResponseResultExcludePajwohLqIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultExcludePajwohLqIPListRule) implementsZoneAccessGroupGetResponseResultExclude() {
}

type ZoneAccessGroupGetResponseResultExcludePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                              `json:"id,required"`
	JSON zoneAccessGroupGetResponseResultExcludePajwohLqIPListRuleIPListJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqIPListRuleIPList]
type zoneAccessGroupGetResponseResultExcludePajwohLqIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type ZoneAccessGroupGetResponseResultExcludePajwohLqCertificateRule struct {
	Certificate interface{}                                                        `json:"certificate,required"`
	JSON        zoneAccessGroupGetResponseResultExcludePajwohLqCertificateRuleJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqCertificateRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqCertificateRule]
type zoneAccessGroupGetResponseResultExcludePajwohLqCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultExcludePajwohLqCertificateRule) implementsZoneAccessGroupGetResponseResultExclude() {
}

// Matches an Access group.
type ZoneAccessGroupGetResponseResultExcludePajwohLqAccessGroupRule struct {
	Group ZoneAccessGroupGetResponseResultExcludePajwohLqAccessGroupRuleGroup `json:"group,required"`
	JSON  zoneAccessGroupGetResponseResultExcludePajwohLqAccessGroupRuleJSON  `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqAccessGroupRule]
type zoneAccessGroupGetResponseResultExcludePajwohLqAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultExcludePajwohLqAccessGroupRule) implementsZoneAccessGroupGetResponseResultExclude() {
}

type ZoneAccessGroupGetResponseResultExcludePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                  `json:"id,required"`
	JSON zoneAccessGroupGetResponseResultExcludePajwohLqAccessGroupRuleGroupJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqAccessGroupRuleGroup]
type zoneAccessGroupGetResponseResultExcludePajwohLqAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZoneAccessGroupGetResponseResultExcludePajwohLqAzureGroupRule struct {
	AzureAd ZoneAccessGroupGetResponseResultExcludePajwohLqAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zoneAccessGroupGetResponseResultExcludePajwohLqAzureGroupRuleJSON    `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqAzureGroupRule]
type zoneAccessGroupGetResponseResultExcludePajwohLqAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultExcludePajwohLqAzureGroupRule) implementsZoneAccessGroupGetResponseResultExclude() {
}

type ZoneAccessGroupGetResponseResultExcludePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                   `json:"connection_id,required"`
	JSON         zoneAccessGroupGetResponseResultExcludePajwohLqAzureGroupRuleAzureAdJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqAzureGroupRuleAzureAd]
type zoneAccessGroupGetResponseResultExcludePajwohLqAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZoneAccessGroupGetResponseResultExcludePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization ZoneAccessGroupGetResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zoneAccessGroupGetResponseResultExcludePajwohLqGitHubOrganizationRuleJSON               `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqGitHubOrganizationRule]
type zoneAccessGroupGetResponseResultExcludePajwohLqGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultExcludePajwohLqGitHubOrganizationRule) implementsZoneAccessGroupGetResponseResultExclude() {
}

type ZoneAccessGroupGetResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                      `json:"name,required"`
	JSON zoneAccessGroupGetResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization]
type zoneAccessGroupGetResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZoneAccessGroupGetResponseResultExcludePajwohLqGsuiteGroupRule struct {
	Gsuite ZoneAccessGroupGetResponseResultExcludePajwohLqGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zoneAccessGroupGetResponseResultExcludePajwohLqGsuiteGroupRuleJSON   `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqGsuiteGroupRule]
type zoneAccessGroupGetResponseResultExcludePajwohLqGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultExcludePajwohLqGsuiteGroupRule) implementsZoneAccessGroupGetResponseResultExclude() {
}

type ZoneAccessGroupGetResponseResultExcludePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                   `json:"email,required"`
	JSON  zoneAccessGroupGetResponseResultExcludePajwohLqGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqGsuiteGroupRuleGsuite]
type zoneAccessGroupGetResponseResultExcludePajwohLqGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZoneAccessGroupGetResponseResultExcludePajwohLqOktaGroupRule struct {
	Okta ZoneAccessGroupGetResponseResultExcludePajwohLqOktaGroupRuleOkta `json:"okta,required"`
	JSON zoneAccessGroupGetResponseResultExcludePajwohLqOktaGroupRuleJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqOktaGroupRule]
type zoneAccessGroupGetResponseResultExcludePajwohLqOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultExcludePajwohLqOktaGroupRule) implementsZoneAccessGroupGetResponseResultExclude() {
}

type ZoneAccessGroupGetResponseResultExcludePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                               `json:"email,required"`
	JSON  zoneAccessGroupGetResponseResultExcludePajwohLqOktaGroupRuleOktaJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqOktaGroupRuleOkta]
type zoneAccessGroupGetResponseResultExcludePajwohLqOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZoneAccessGroupGetResponseResultExcludePajwohLqSamlGroupRule struct {
	Saml ZoneAccessGroupGetResponseResultExcludePajwohLqSamlGroupRuleSaml `json:"saml,required"`
	JSON zoneAccessGroupGetResponseResultExcludePajwohLqSamlGroupRuleJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqSamlGroupRule]
type zoneAccessGroupGetResponseResultExcludePajwohLqSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultExcludePajwohLqSamlGroupRule) implementsZoneAccessGroupGetResponseResultExclude() {
}

type ZoneAccessGroupGetResponseResultExcludePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                               `json:"attribute_value,required"`
	JSON           zoneAccessGroupGetResponseResultExcludePajwohLqSamlGroupRuleSamlJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqSamlGroupRuleSaml]
type zoneAccessGroupGetResponseResultExcludePajwohLqSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type ZoneAccessGroupGetResponseResultExcludePajwohLqServiceTokenRule struct {
	ServiceToken ZoneAccessGroupGetResponseResultExcludePajwohLqServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zoneAccessGroupGetResponseResultExcludePajwohLqServiceTokenRuleJSON         `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqServiceTokenRule]
type zoneAccessGroupGetResponseResultExcludePajwohLqServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultExcludePajwohLqServiceTokenRule) implementsZoneAccessGroupGetResponseResultExclude() {
}

type ZoneAccessGroupGetResponseResultExcludePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                          `json:"token_id,required"`
	JSON    zoneAccessGroupGetResponseResultExcludePajwohLqServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqServiceTokenRuleServiceToken]
type zoneAccessGroupGetResponseResultExcludePajwohLqServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type ZoneAccessGroupGetResponseResultExcludePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                 `json:"any_valid_service_token,required"`
	JSON                 zoneAccessGroupGetResponseResultExcludePajwohLqAnyValidServiceTokenRuleJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqAnyValidServiceTokenRule]
type zoneAccessGroupGetResponseResultExcludePajwohLqAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultExcludePajwohLqAnyValidServiceTokenRule) implementsZoneAccessGroupGetResponseResultExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZoneAccessGroupGetResponseResultExcludePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation ZoneAccessGroupGetResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zoneAccessGroupGetResponseResultExcludePajwohLqExternalEvaluationRuleJSON               `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqExternalEvaluationRule]
type zoneAccessGroupGetResponseResultExcludePajwohLqExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultExcludePajwohLqExternalEvaluationRule) implementsZoneAccessGroupGetResponseResultExclude() {
}

type ZoneAccessGroupGetResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                      `json:"keys_url,required"`
	JSON    zoneAccessGroupGetResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation]
type zoneAccessGroupGetResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type ZoneAccessGroupGetResponseResultExcludePajwohLqCountryRule struct {
	Geo  ZoneAccessGroupGetResponseResultExcludePajwohLqCountryRuleGeo  `json:"geo,required"`
	JSON zoneAccessGroupGetResponseResultExcludePajwohLqCountryRuleJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqCountryRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqCountryRule]
type zoneAccessGroupGetResponseResultExcludePajwohLqCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultExcludePajwohLqCountryRule) implementsZoneAccessGroupGetResponseResultExclude() {
}

type ZoneAccessGroupGetResponseResultExcludePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                            `json:"country_code,required"`
	JSON        zoneAccessGroupGetResponseResultExcludePajwohLqCountryRuleGeoJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqCountryRuleGeo]
type zoneAccessGroupGetResponseResultExcludePajwohLqCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type ZoneAccessGroupGetResponseResultExcludePajwohLqAuthenticationMethodRule struct {
	AuthMethod ZoneAccessGroupGetResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zoneAccessGroupGetResponseResultExcludePajwohLqAuthenticationMethodRuleJSON       `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqAuthenticationMethodRule]
type zoneAccessGroupGetResponseResultExcludePajwohLqAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultExcludePajwohLqAuthenticationMethodRule) implementsZoneAccessGroupGetResponseResultExclude() {
}

type ZoneAccessGroupGetResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                `json:"auth_method,required"`
	JSON       zoneAccessGroupGetResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod]
type zoneAccessGroupGetResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type ZoneAccessGroupGetResponseResultExcludePajwohLqDevicePostureRule struct {
	DevicePosture ZoneAccessGroupGetResponseResultExcludePajwohLqDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zoneAccessGroupGetResponseResultExcludePajwohLqDevicePostureRuleJSON          `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqDevicePostureRule]
type zoneAccessGroupGetResponseResultExcludePajwohLqDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultExcludePajwohLqDevicePostureRule) implementsZoneAccessGroupGetResponseResultExclude() {
}

type ZoneAccessGroupGetResponseResultExcludePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                            `json:"integration_uid,required"`
	JSON           zoneAccessGroupGetResponseResultExcludePajwohLqDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultExcludePajwohLqDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultExcludePajwohLqDevicePostureRuleDevicePosture]
type zoneAccessGroupGetResponseResultExcludePajwohLqDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultExcludePajwohLqDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [ZoneAccessGroupGetResponseResultIncludePajwohLqEmailRule],
// [ZoneAccessGroupGetResponseResultIncludePajwohLqEmailListRule],
// [ZoneAccessGroupGetResponseResultIncludePajwohLqDomainRule],
// [ZoneAccessGroupGetResponseResultIncludePajwohLqEveryoneRule],
// [ZoneAccessGroupGetResponseResultIncludePajwohLqIPRule],
// [ZoneAccessGroupGetResponseResultIncludePajwohLqIPListRule],
// [ZoneAccessGroupGetResponseResultIncludePajwohLqCertificateRule],
// [ZoneAccessGroupGetResponseResultIncludePajwohLqAccessGroupRule],
// [ZoneAccessGroupGetResponseResultIncludePajwohLqAzureGroupRule],
// [ZoneAccessGroupGetResponseResultIncludePajwohLqGitHubOrganizationRule],
// [ZoneAccessGroupGetResponseResultIncludePajwohLqGsuiteGroupRule],
// [ZoneAccessGroupGetResponseResultIncludePajwohLqOktaGroupRule],
// [ZoneAccessGroupGetResponseResultIncludePajwohLqSamlGroupRule],
// [ZoneAccessGroupGetResponseResultIncludePajwohLqServiceTokenRule],
// [ZoneAccessGroupGetResponseResultIncludePajwohLqAnyValidServiceTokenRule],
// [ZoneAccessGroupGetResponseResultIncludePajwohLqExternalEvaluationRule],
// [ZoneAccessGroupGetResponseResultIncludePajwohLqCountryRule],
// [ZoneAccessGroupGetResponseResultIncludePajwohLqAuthenticationMethodRule] or
// [ZoneAccessGroupGetResponseResultIncludePajwohLqDevicePostureRule].
type ZoneAccessGroupGetResponseResultInclude interface {
	implementsZoneAccessGroupGetResponseResultInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneAccessGroupGetResponseResultInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type ZoneAccessGroupGetResponseResultIncludePajwohLqEmailRule struct {
	Email ZoneAccessGroupGetResponseResultIncludePajwohLqEmailRuleEmail `json:"email,required"`
	JSON  zoneAccessGroupGetResponseResultIncludePajwohLqEmailRuleJSON  `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqEmailRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqEmailRule]
type zoneAccessGroupGetResponseResultIncludePajwohLqEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultIncludePajwohLqEmailRule) implementsZoneAccessGroupGetResponseResultInclude() {
}

type ZoneAccessGroupGetResponseResultIncludePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email string                                                            `json:"email,required" format:"email"`
	JSON  zoneAccessGroupGetResponseResultIncludePajwohLqEmailRuleEmailJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqEmailRuleEmail]
type zoneAccessGroupGetResponseResultIncludePajwohLqEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type ZoneAccessGroupGetResponseResultIncludePajwohLqEmailListRule struct {
	EmailList ZoneAccessGroupGetResponseResultIncludePajwohLqEmailListRuleEmailList `json:"email_list,required"`
	JSON      zoneAccessGroupGetResponseResultIncludePajwohLqEmailListRuleJSON      `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqEmailListRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqEmailListRule]
type zoneAccessGroupGetResponseResultIncludePajwohLqEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultIncludePajwohLqEmailListRule) implementsZoneAccessGroupGetResponseResultInclude() {
}

type ZoneAccessGroupGetResponseResultIncludePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                    `json:"id,required"`
	JSON zoneAccessGroupGetResponseResultIncludePajwohLqEmailListRuleEmailListJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqEmailListRuleEmailList]
type zoneAccessGroupGetResponseResultIncludePajwohLqEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type ZoneAccessGroupGetResponseResultIncludePajwohLqDomainRule struct {
	EmailDomain ZoneAccessGroupGetResponseResultIncludePajwohLqDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zoneAccessGroupGetResponseResultIncludePajwohLqDomainRuleJSON        `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqDomainRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqDomainRule]
type zoneAccessGroupGetResponseResultIncludePajwohLqDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultIncludePajwohLqDomainRule) implementsZoneAccessGroupGetResponseResultInclude() {
}

type ZoneAccessGroupGetResponseResultIncludePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                   `json:"domain,required"`
	JSON   zoneAccessGroupGetResponseResultIncludePajwohLqDomainRuleEmailDomainJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqDomainRuleEmailDomain]
type zoneAccessGroupGetResponseResultIncludePajwohLqDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type ZoneAccessGroupGetResponseResultIncludePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                     `json:"everyone,required"`
	JSON     zoneAccessGroupGetResponseResultIncludePajwohLqEveryoneRuleJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqEveryoneRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqEveryoneRule]
type zoneAccessGroupGetResponseResultIncludePajwohLqEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultIncludePajwohLqEveryoneRule) implementsZoneAccessGroupGetResponseResultInclude() {
}

// Matches an IP address block.
type ZoneAccessGroupGetResponseResultIncludePajwohLqIPRule struct {
	IP   ZoneAccessGroupGetResponseResultIncludePajwohLqIPRuleIP   `json:"ip,required"`
	JSON zoneAccessGroupGetResponseResultIncludePajwohLqIPRuleJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqIPRuleJSON contains the JSON
// metadata for the struct [ZoneAccessGroupGetResponseResultIncludePajwohLqIPRule]
type zoneAccessGroupGetResponseResultIncludePajwohLqIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultIncludePajwohLqIPRule) implementsZoneAccessGroupGetResponseResultInclude() {
}

type ZoneAccessGroupGetResponseResultIncludePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                      `json:"ip,required"`
	JSON zoneAccessGroupGetResponseResultIncludePajwohLqIPRuleIPJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqIPRuleIPJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqIPRuleIP]
type zoneAccessGroupGetResponseResultIncludePajwohLqIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type ZoneAccessGroupGetResponseResultIncludePajwohLqIPListRule struct {
	IPList ZoneAccessGroupGetResponseResultIncludePajwohLqIPListRuleIPList `json:"ip_list,required"`
	JSON   zoneAccessGroupGetResponseResultIncludePajwohLqIPListRuleJSON   `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqIPListRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqIPListRule]
type zoneAccessGroupGetResponseResultIncludePajwohLqIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultIncludePajwohLqIPListRule) implementsZoneAccessGroupGetResponseResultInclude() {
}

type ZoneAccessGroupGetResponseResultIncludePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                              `json:"id,required"`
	JSON zoneAccessGroupGetResponseResultIncludePajwohLqIPListRuleIPListJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqIPListRuleIPList]
type zoneAccessGroupGetResponseResultIncludePajwohLqIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type ZoneAccessGroupGetResponseResultIncludePajwohLqCertificateRule struct {
	Certificate interface{}                                                        `json:"certificate,required"`
	JSON        zoneAccessGroupGetResponseResultIncludePajwohLqCertificateRuleJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqCertificateRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqCertificateRule]
type zoneAccessGroupGetResponseResultIncludePajwohLqCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultIncludePajwohLqCertificateRule) implementsZoneAccessGroupGetResponseResultInclude() {
}

// Matches an Access group.
type ZoneAccessGroupGetResponseResultIncludePajwohLqAccessGroupRule struct {
	Group ZoneAccessGroupGetResponseResultIncludePajwohLqAccessGroupRuleGroup `json:"group,required"`
	JSON  zoneAccessGroupGetResponseResultIncludePajwohLqAccessGroupRuleJSON  `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqAccessGroupRule]
type zoneAccessGroupGetResponseResultIncludePajwohLqAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultIncludePajwohLqAccessGroupRule) implementsZoneAccessGroupGetResponseResultInclude() {
}

type ZoneAccessGroupGetResponseResultIncludePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                  `json:"id,required"`
	JSON zoneAccessGroupGetResponseResultIncludePajwohLqAccessGroupRuleGroupJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqAccessGroupRuleGroup]
type zoneAccessGroupGetResponseResultIncludePajwohLqAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZoneAccessGroupGetResponseResultIncludePajwohLqAzureGroupRule struct {
	AzureAd ZoneAccessGroupGetResponseResultIncludePajwohLqAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zoneAccessGroupGetResponseResultIncludePajwohLqAzureGroupRuleJSON    `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqAzureGroupRule]
type zoneAccessGroupGetResponseResultIncludePajwohLqAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultIncludePajwohLqAzureGroupRule) implementsZoneAccessGroupGetResponseResultInclude() {
}

type ZoneAccessGroupGetResponseResultIncludePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                   `json:"connection_id,required"`
	JSON         zoneAccessGroupGetResponseResultIncludePajwohLqAzureGroupRuleAzureAdJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqAzureGroupRuleAzureAd]
type zoneAccessGroupGetResponseResultIncludePajwohLqAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZoneAccessGroupGetResponseResultIncludePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization ZoneAccessGroupGetResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zoneAccessGroupGetResponseResultIncludePajwohLqGitHubOrganizationRuleJSON               `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqGitHubOrganizationRule]
type zoneAccessGroupGetResponseResultIncludePajwohLqGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultIncludePajwohLqGitHubOrganizationRule) implementsZoneAccessGroupGetResponseResultInclude() {
}

type ZoneAccessGroupGetResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                      `json:"name,required"`
	JSON zoneAccessGroupGetResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization]
type zoneAccessGroupGetResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZoneAccessGroupGetResponseResultIncludePajwohLqGsuiteGroupRule struct {
	Gsuite ZoneAccessGroupGetResponseResultIncludePajwohLqGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zoneAccessGroupGetResponseResultIncludePajwohLqGsuiteGroupRuleJSON   `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqGsuiteGroupRule]
type zoneAccessGroupGetResponseResultIncludePajwohLqGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultIncludePajwohLqGsuiteGroupRule) implementsZoneAccessGroupGetResponseResultInclude() {
}

type ZoneAccessGroupGetResponseResultIncludePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                   `json:"email,required"`
	JSON  zoneAccessGroupGetResponseResultIncludePajwohLqGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqGsuiteGroupRuleGsuite]
type zoneAccessGroupGetResponseResultIncludePajwohLqGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZoneAccessGroupGetResponseResultIncludePajwohLqOktaGroupRule struct {
	Okta ZoneAccessGroupGetResponseResultIncludePajwohLqOktaGroupRuleOkta `json:"okta,required"`
	JSON zoneAccessGroupGetResponseResultIncludePajwohLqOktaGroupRuleJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqOktaGroupRule]
type zoneAccessGroupGetResponseResultIncludePajwohLqOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultIncludePajwohLqOktaGroupRule) implementsZoneAccessGroupGetResponseResultInclude() {
}

type ZoneAccessGroupGetResponseResultIncludePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                               `json:"email,required"`
	JSON  zoneAccessGroupGetResponseResultIncludePajwohLqOktaGroupRuleOktaJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqOktaGroupRuleOkta]
type zoneAccessGroupGetResponseResultIncludePajwohLqOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZoneAccessGroupGetResponseResultIncludePajwohLqSamlGroupRule struct {
	Saml ZoneAccessGroupGetResponseResultIncludePajwohLqSamlGroupRuleSaml `json:"saml,required"`
	JSON zoneAccessGroupGetResponseResultIncludePajwohLqSamlGroupRuleJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqSamlGroupRule]
type zoneAccessGroupGetResponseResultIncludePajwohLqSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultIncludePajwohLqSamlGroupRule) implementsZoneAccessGroupGetResponseResultInclude() {
}

type ZoneAccessGroupGetResponseResultIncludePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                               `json:"attribute_value,required"`
	JSON           zoneAccessGroupGetResponseResultIncludePajwohLqSamlGroupRuleSamlJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqSamlGroupRuleSaml]
type zoneAccessGroupGetResponseResultIncludePajwohLqSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type ZoneAccessGroupGetResponseResultIncludePajwohLqServiceTokenRule struct {
	ServiceToken ZoneAccessGroupGetResponseResultIncludePajwohLqServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zoneAccessGroupGetResponseResultIncludePajwohLqServiceTokenRuleJSON         `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqServiceTokenRule]
type zoneAccessGroupGetResponseResultIncludePajwohLqServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultIncludePajwohLqServiceTokenRule) implementsZoneAccessGroupGetResponseResultInclude() {
}

type ZoneAccessGroupGetResponseResultIncludePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                          `json:"token_id,required"`
	JSON    zoneAccessGroupGetResponseResultIncludePajwohLqServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqServiceTokenRuleServiceToken]
type zoneAccessGroupGetResponseResultIncludePajwohLqServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type ZoneAccessGroupGetResponseResultIncludePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                 `json:"any_valid_service_token,required"`
	JSON                 zoneAccessGroupGetResponseResultIncludePajwohLqAnyValidServiceTokenRuleJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqAnyValidServiceTokenRule]
type zoneAccessGroupGetResponseResultIncludePajwohLqAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultIncludePajwohLqAnyValidServiceTokenRule) implementsZoneAccessGroupGetResponseResultInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZoneAccessGroupGetResponseResultIncludePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation ZoneAccessGroupGetResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zoneAccessGroupGetResponseResultIncludePajwohLqExternalEvaluationRuleJSON               `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqExternalEvaluationRule]
type zoneAccessGroupGetResponseResultIncludePajwohLqExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultIncludePajwohLqExternalEvaluationRule) implementsZoneAccessGroupGetResponseResultInclude() {
}

type ZoneAccessGroupGetResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                      `json:"keys_url,required"`
	JSON    zoneAccessGroupGetResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation]
type zoneAccessGroupGetResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type ZoneAccessGroupGetResponseResultIncludePajwohLqCountryRule struct {
	Geo  ZoneAccessGroupGetResponseResultIncludePajwohLqCountryRuleGeo  `json:"geo,required"`
	JSON zoneAccessGroupGetResponseResultIncludePajwohLqCountryRuleJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqCountryRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqCountryRule]
type zoneAccessGroupGetResponseResultIncludePajwohLqCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultIncludePajwohLqCountryRule) implementsZoneAccessGroupGetResponseResultInclude() {
}

type ZoneAccessGroupGetResponseResultIncludePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                            `json:"country_code,required"`
	JSON        zoneAccessGroupGetResponseResultIncludePajwohLqCountryRuleGeoJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqCountryRuleGeo]
type zoneAccessGroupGetResponseResultIncludePajwohLqCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type ZoneAccessGroupGetResponseResultIncludePajwohLqAuthenticationMethodRule struct {
	AuthMethod ZoneAccessGroupGetResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zoneAccessGroupGetResponseResultIncludePajwohLqAuthenticationMethodRuleJSON       `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqAuthenticationMethodRule]
type zoneAccessGroupGetResponseResultIncludePajwohLqAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultIncludePajwohLqAuthenticationMethodRule) implementsZoneAccessGroupGetResponseResultInclude() {
}

type ZoneAccessGroupGetResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                `json:"auth_method,required"`
	JSON       zoneAccessGroupGetResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod]
type zoneAccessGroupGetResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type ZoneAccessGroupGetResponseResultIncludePajwohLqDevicePostureRule struct {
	DevicePosture ZoneAccessGroupGetResponseResultIncludePajwohLqDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zoneAccessGroupGetResponseResultIncludePajwohLqDevicePostureRuleJSON          `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqDevicePostureRule]
type zoneAccessGroupGetResponseResultIncludePajwohLqDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultIncludePajwohLqDevicePostureRule) implementsZoneAccessGroupGetResponseResultInclude() {
}

type ZoneAccessGroupGetResponseResultIncludePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                            `json:"integration_uid,required"`
	JSON           zoneAccessGroupGetResponseResultIncludePajwohLqDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultIncludePajwohLqDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultIncludePajwohLqDevicePostureRuleDevicePosture]
type zoneAccessGroupGetResponseResultIncludePajwohLqDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultIncludePajwohLqDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [ZoneAccessGroupGetResponseResultRequirePajwohLqEmailRule],
// [ZoneAccessGroupGetResponseResultRequirePajwohLqEmailListRule],
// [ZoneAccessGroupGetResponseResultRequirePajwohLqDomainRule],
// [ZoneAccessGroupGetResponseResultRequirePajwohLqEveryoneRule],
// [ZoneAccessGroupGetResponseResultRequirePajwohLqIPRule],
// [ZoneAccessGroupGetResponseResultRequirePajwohLqIPListRule],
// [ZoneAccessGroupGetResponseResultRequirePajwohLqCertificateRule],
// [ZoneAccessGroupGetResponseResultRequirePajwohLqAccessGroupRule],
// [ZoneAccessGroupGetResponseResultRequirePajwohLqAzureGroupRule],
// [ZoneAccessGroupGetResponseResultRequirePajwohLqGitHubOrganizationRule],
// [ZoneAccessGroupGetResponseResultRequirePajwohLqGsuiteGroupRule],
// [ZoneAccessGroupGetResponseResultRequirePajwohLqOktaGroupRule],
// [ZoneAccessGroupGetResponseResultRequirePajwohLqSamlGroupRule],
// [ZoneAccessGroupGetResponseResultRequirePajwohLqServiceTokenRule],
// [ZoneAccessGroupGetResponseResultRequirePajwohLqAnyValidServiceTokenRule],
// [ZoneAccessGroupGetResponseResultRequirePajwohLqExternalEvaluationRule],
// [ZoneAccessGroupGetResponseResultRequirePajwohLqCountryRule],
// [ZoneAccessGroupGetResponseResultRequirePajwohLqAuthenticationMethodRule] or
// [ZoneAccessGroupGetResponseResultRequirePajwohLqDevicePostureRule].
type ZoneAccessGroupGetResponseResultRequire interface {
	implementsZoneAccessGroupGetResponseResultRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneAccessGroupGetResponseResultRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type ZoneAccessGroupGetResponseResultRequirePajwohLqEmailRule struct {
	Email ZoneAccessGroupGetResponseResultRequirePajwohLqEmailRuleEmail `json:"email,required"`
	JSON  zoneAccessGroupGetResponseResultRequirePajwohLqEmailRuleJSON  `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqEmailRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqEmailRule]
type zoneAccessGroupGetResponseResultRequirePajwohLqEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultRequirePajwohLqEmailRule) implementsZoneAccessGroupGetResponseResultRequire() {
}

type ZoneAccessGroupGetResponseResultRequirePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email string                                                            `json:"email,required" format:"email"`
	JSON  zoneAccessGroupGetResponseResultRequirePajwohLqEmailRuleEmailJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqEmailRuleEmail]
type zoneAccessGroupGetResponseResultRequirePajwohLqEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type ZoneAccessGroupGetResponseResultRequirePajwohLqEmailListRule struct {
	EmailList ZoneAccessGroupGetResponseResultRequirePajwohLqEmailListRuleEmailList `json:"email_list,required"`
	JSON      zoneAccessGroupGetResponseResultRequirePajwohLqEmailListRuleJSON      `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqEmailListRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqEmailListRule]
type zoneAccessGroupGetResponseResultRequirePajwohLqEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultRequirePajwohLqEmailListRule) implementsZoneAccessGroupGetResponseResultRequire() {
}

type ZoneAccessGroupGetResponseResultRequirePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                    `json:"id,required"`
	JSON zoneAccessGroupGetResponseResultRequirePajwohLqEmailListRuleEmailListJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqEmailListRuleEmailList]
type zoneAccessGroupGetResponseResultRequirePajwohLqEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type ZoneAccessGroupGetResponseResultRequirePajwohLqDomainRule struct {
	EmailDomain ZoneAccessGroupGetResponseResultRequirePajwohLqDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zoneAccessGroupGetResponseResultRequirePajwohLqDomainRuleJSON        `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqDomainRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqDomainRule]
type zoneAccessGroupGetResponseResultRequirePajwohLqDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultRequirePajwohLqDomainRule) implementsZoneAccessGroupGetResponseResultRequire() {
}

type ZoneAccessGroupGetResponseResultRequirePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                   `json:"domain,required"`
	JSON   zoneAccessGroupGetResponseResultRequirePajwohLqDomainRuleEmailDomainJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqDomainRuleEmailDomain]
type zoneAccessGroupGetResponseResultRequirePajwohLqDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type ZoneAccessGroupGetResponseResultRequirePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                     `json:"everyone,required"`
	JSON     zoneAccessGroupGetResponseResultRequirePajwohLqEveryoneRuleJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqEveryoneRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqEveryoneRule]
type zoneAccessGroupGetResponseResultRequirePajwohLqEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultRequirePajwohLqEveryoneRule) implementsZoneAccessGroupGetResponseResultRequire() {
}

// Matches an IP address block.
type ZoneAccessGroupGetResponseResultRequirePajwohLqIPRule struct {
	IP   ZoneAccessGroupGetResponseResultRequirePajwohLqIPRuleIP   `json:"ip,required"`
	JSON zoneAccessGroupGetResponseResultRequirePajwohLqIPRuleJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqIPRuleJSON contains the JSON
// metadata for the struct [ZoneAccessGroupGetResponseResultRequirePajwohLqIPRule]
type zoneAccessGroupGetResponseResultRequirePajwohLqIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultRequirePajwohLqIPRule) implementsZoneAccessGroupGetResponseResultRequire() {
}

type ZoneAccessGroupGetResponseResultRequirePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                      `json:"ip,required"`
	JSON zoneAccessGroupGetResponseResultRequirePajwohLqIPRuleIPJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqIPRuleIPJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqIPRuleIP]
type zoneAccessGroupGetResponseResultRequirePajwohLqIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type ZoneAccessGroupGetResponseResultRequirePajwohLqIPListRule struct {
	IPList ZoneAccessGroupGetResponseResultRequirePajwohLqIPListRuleIPList `json:"ip_list,required"`
	JSON   zoneAccessGroupGetResponseResultRequirePajwohLqIPListRuleJSON   `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqIPListRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqIPListRule]
type zoneAccessGroupGetResponseResultRequirePajwohLqIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultRequirePajwohLqIPListRule) implementsZoneAccessGroupGetResponseResultRequire() {
}

type ZoneAccessGroupGetResponseResultRequirePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                              `json:"id,required"`
	JSON zoneAccessGroupGetResponseResultRequirePajwohLqIPListRuleIPListJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqIPListRuleIPList]
type zoneAccessGroupGetResponseResultRequirePajwohLqIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type ZoneAccessGroupGetResponseResultRequirePajwohLqCertificateRule struct {
	Certificate interface{}                                                        `json:"certificate,required"`
	JSON        zoneAccessGroupGetResponseResultRequirePajwohLqCertificateRuleJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqCertificateRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqCertificateRule]
type zoneAccessGroupGetResponseResultRequirePajwohLqCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultRequirePajwohLqCertificateRule) implementsZoneAccessGroupGetResponseResultRequire() {
}

// Matches an Access group.
type ZoneAccessGroupGetResponseResultRequirePajwohLqAccessGroupRule struct {
	Group ZoneAccessGroupGetResponseResultRequirePajwohLqAccessGroupRuleGroup `json:"group,required"`
	JSON  zoneAccessGroupGetResponseResultRequirePajwohLqAccessGroupRuleJSON  `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqAccessGroupRule]
type zoneAccessGroupGetResponseResultRequirePajwohLqAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultRequirePajwohLqAccessGroupRule) implementsZoneAccessGroupGetResponseResultRequire() {
}

type ZoneAccessGroupGetResponseResultRequirePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                  `json:"id,required"`
	JSON zoneAccessGroupGetResponseResultRequirePajwohLqAccessGroupRuleGroupJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqAccessGroupRuleGroup]
type zoneAccessGroupGetResponseResultRequirePajwohLqAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZoneAccessGroupGetResponseResultRequirePajwohLqAzureGroupRule struct {
	AzureAd ZoneAccessGroupGetResponseResultRequirePajwohLqAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zoneAccessGroupGetResponseResultRequirePajwohLqAzureGroupRuleJSON    `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqAzureGroupRule]
type zoneAccessGroupGetResponseResultRequirePajwohLqAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultRequirePajwohLqAzureGroupRule) implementsZoneAccessGroupGetResponseResultRequire() {
}

type ZoneAccessGroupGetResponseResultRequirePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                   `json:"connection_id,required"`
	JSON         zoneAccessGroupGetResponseResultRequirePajwohLqAzureGroupRuleAzureAdJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqAzureGroupRuleAzureAd]
type zoneAccessGroupGetResponseResultRequirePajwohLqAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZoneAccessGroupGetResponseResultRequirePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization ZoneAccessGroupGetResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zoneAccessGroupGetResponseResultRequirePajwohLqGitHubOrganizationRuleJSON               `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqGitHubOrganizationRule]
type zoneAccessGroupGetResponseResultRequirePajwohLqGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultRequirePajwohLqGitHubOrganizationRule) implementsZoneAccessGroupGetResponseResultRequire() {
}

type ZoneAccessGroupGetResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                      `json:"name,required"`
	JSON zoneAccessGroupGetResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization]
type zoneAccessGroupGetResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZoneAccessGroupGetResponseResultRequirePajwohLqGsuiteGroupRule struct {
	Gsuite ZoneAccessGroupGetResponseResultRequirePajwohLqGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zoneAccessGroupGetResponseResultRequirePajwohLqGsuiteGroupRuleJSON   `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqGsuiteGroupRule]
type zoneAccessGroupGetResponseResultRequirePajwohLqGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultRequirePajwohLqGsuiteGroupRule) implementsZoneAccessGroupGetResponseResultRequire() {
}

type ZoneAccessGroupGetResponseResultRequirePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                   `json:"email,required"`
	JSON  zoneAccessGroupGetResponseResultRequirePajwohLqGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqGsuiteGroupRuleGsuite]
type zoneAccessGroupGetResponseResultRequirePajwohLqGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZoneAccessGroupGetResponseResultRequirePajwohLqOktaGroupRule struct {
	Okta ZoneAccessGroupGetResponseResultRequirePajwohLqOktaGroupRuleOkta `json:"okta,required"`
	JSON zoneAccessGroupGetResponseResultRequirePajwohLqOktaGroupRuleJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqOktaGroupRule]
type zoneAccessGroupGetResponseResultRequirePajwohLqOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultRequirePajwohLqOktaGroupRule) implementsZoneAccessGroupGetResponseResultRequire() {
}

type ZoneAccessGroupGetResponseResultRequirePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                               `json:"email,required"`
	JSON  zoneAccessGroupGetResponseResultRequirePajwohLqOktaGroupRuleOktaJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqOktaGroupRuleOkta]
type zoneAccessGroupGetResponseResultRequirePajwohLqOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZoneAccessGroupGetResponseResultRequirePajwohLqSamlGroupRule struct {
	Saml ZoneAccessGroupGetResponseResultRequirePajwohLqSamlGroupRuleSaml `json:"saml,required"`
	JSON zoneAccessGroupGetResponseResultRequirePajwohLqSamlGroupRuleJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqSamlGroupRule]
type zoneAccessGroupGetResponseResultRequirePajwohLqSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultRequirePajwohLqSamlGroupRule) implementsZoneAccessGroupGetResponseResultRequire() {
}

type ZoneAccessGroupGetResponseResultRequirePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                               `json:"attribute_value,required"`
	JSON           zoneAccessGroupGetResponseResultRequirePajwohLqSamlGroupRuleSamlJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqSamlGroupRuleSaml]
type zoneAccessGroupGetResponseResultRequirePajwohLqSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type ZoneAccessGroupGetResponseResultRequirePajwohLqServiceTokenRule struct {
	ServiceToken ZoneAccessGroupGetResponseResultRequirePajwohLqServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zoneAccessGroupGetResponseResultRequirePajwohLqServiceTokenRuleJSON         `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqServiceTokenRule]
type zoneAccessGroupGetResponseResultRequirePajwohLqServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultRequirePajwohLqServiceTokenRule) implementsZoneAccessGroupGetResponseResultRequire() {
}

type ZoneAccessGroupGetResponseResultRequirePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                          `json:"token_id,required"`
	JSON    zoneAccessGroupGetResponseResultRequirePajwohLqServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqServiceTokenRuleServiceToken]
type zoneAccessGroupGetResponseResultRequirePajwohLqServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type ZoneAccessGroupGetResponseResultRequirePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                 `json:"any_valid_service_token,required"`
	JSON                 zoneAccessGroupGetResponseResultRequirePajwohLqAnyValidServiceTokenRuleJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqAnyValidServiceTokenRule]
type zoneAccessGroupGetResponseResultRequirePajwohLqAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultRequirePajwohLqAnyValidServiceTokenRule) implementsZoneAccessGroupGetResponseResultRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZoneAccessGroupGetResponseResultRequirePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation ZoneAccessGroupGetResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zoneAccessGroupGetResponseResultRequirePajwohLqExternalEvaluationRuleJSON               `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqExternalEvaluationRule]
type zoneAccessGroupGetResponseResultRequirePajwohLqExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultRequirePajwohLqExternalEvaluationRule) implementsZoneAccessGroupGetResponseResultRequire() {
}

type ZoneAccessGroupGetResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                      `json:"keys_url,required"`
	JSON    zoneAccessGroupGetResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation]
type zoneAccessGroupGetResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type ZoneAccessGroupGetResponseResultRequirePajwohLqCountryRule struct {
	Geo  ZoneAccessGroupGetResponseResultRequirePajwohLqCountryRuleGeo  `json:"geo,required"`
	JSON zoneAccessGroupGetResponseResultRequirePajwohLqCountryRuleJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqCountryRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqCountryRule]
type zoneAccessGroupGetResponseResultRequirePajwohLqCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultRequirePajwohLqCountryRule) implementsZoneAccessGroupGetResponseResultRequire() {
}

type ZoneAccessGroupGetResponseResultRequirePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                            `json:"country_code,required"`
	JSON        zoneAccessGroupGetResponseResultRequirePajwohLqCountryRuleGeoJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqCountryRuleGeo]
type zoneAccessGroupGetResponseResultRequirePajwohLqCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type ZoneAccessGroupGetResponseResultRequirePajwohLqAuthenticationMethodRule struct {
	AuthMethod ZoneAccessGroupGetResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zoneAccessGroupGetResponseResultRequirePajwohLqAuthenticationMethodRuleJSON       `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqAuthenticationMethodRule]
type zoneAccessGroupGetResponseResultRequirePajwohLqAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultRequirePajwohLqAuthenticationMethodRule) implementsZoneAccessGroupGetResponseResultRequire() {
}

type ZoneAccessGroupGetResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                `json:"auth_method,required"`
	JSON       zoneAccessGroupGetResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod]
type zoneAccessGroupGetResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type ZoneAccessGroupGetResponseResultRequirePajwohLqDevicePostureRule struct {
	DevicePosture ZoneAccessGroupGetResponseResultRequirePajwohLqDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zoneAccessGroupGetResponseResultRequirePajwohLqDevicePostureRuleJSON          `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqDevicePostureRule]
type zoneAccessGroupGetResponseResultRequirePajwohLqDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupGetResponseResultRequirePajwohLqDevicePostureRule) implementsZoneAccessGroupGetResponseResultRequire() {
}

type ZoneAccessGroupGetResponseResultRequirePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                            `json:"integration_uid,required"`
	JSON           zoneAccessGroupGetResponseResultRequirePajwohLqDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zoneAccessGroupGetResponseResultRequirePajwohLqDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupGetResponseResultRequirePajwohLqDevicePostureRuleDevicePosture]
type zoneAccessGroupGetResponseResultRequirePajwohLqDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZoneAccessGroupGetResponseResultRequirePajwohLqDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneAccessGroupGetResponseSuccess bool

const (
	ZoneAccessGroupGetResponseSuccessTrue ZoneAccessGroupGetResponseSuccess = true
)

type ZoneAccessGroupUpdateResponse struct {
	Errors   []ZoneAccessGroupUpdateResponseError   `json:"errors"`
	Messages []ZoneAccessGroupUpdateResponseMessage `json:"messages"`
	Result   ZoneAccessGroupUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneAccessGroupUpdateResponseSuccess `json:"success"`
	JSON    zoneAccessGroupUpdateResponseJSON    `json:"-"`
}

// zoneAccessGroupUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponse]
type zoneAccessGroupUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessGroupUpdateResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneAccessGroupUpdateResponseErrorJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseErrorJSON contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseError]
type zoneAccessGroupUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessGroupUpdateResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneAccessGroupUpdateResponseMessageJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneAccessGroupUpdateResponseMessage]
type zoneAccessGroupUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessGroupUpdateResponseResult struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []ZoneAccessGroupUpdateResponseResultExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []ZoneAccessGroupUpdateResponseResultInclude `json:"include"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []ZoneAccessGroupUpdateResponseResultRequire `json:"require"`
	UpdatedAt time.Time                                    `json:"updated_at" format:"date-time"`
	JSON      zoneAccessGroupUpdateResponseResultJSON      `json:"-"`
}

// zoneAccessGroupUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneAccessGroupUpdateResponseResult]
type zoneAccessGroupUpdateResponseResultJSON struct {
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

func (r *ZoneAccessGroupUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqEmailRule],
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqEmailListRule],
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqDomainRule],
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqEveryoneRule],
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqIPRule],
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqIPListRule],
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqCertificateRule],
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqAccessGroupRule],
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqAzureGroupRule],
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqGitHubOrganizationRule],
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqGsuiteGroupRule],
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqOktaGroupRule],
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqSamlGroupRule],
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqServiceTokenRule],
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqAnyValidServiceTokenRule],
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqExternalEvaluationRule],
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqCountryRule],
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqAuthenticationMethodRule] or
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqDevicePostureRule].
type ZoneAccessGroupUpdateResponseResultExclude interface {
	implementsZoneAccessGroupUpdateResponseResultExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneAccessGroupUpdateResponseResultExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type ZoneAccessGroupUpdateResponseResultExcludePajwohLqEmailRule struct {
	Email ZoneAccessGroupUpdateResponseResultExcludePajwohLqEmailRuleEmail `json:"email,required"`
	JSON  zoneAccessGroupUpdateResponseResultExcludePajwohLqEmailRuleJSON  `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqEmailRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqEmailRule]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultExcludePajwohLqEmailRule) implementsZoneAccessGroupUpdateResponseResultExclude() {
}

type ZoneAccessGroupUpdateResponseResultExcludePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email string                                                               `json:"email,required" format:"email"`
	JSON  zoneAccessGroupUpdateResponseResultExcludePajwohLqEmailRuleEmailJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqEmailRuleEmailJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqEmailRuleEmail]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type ZoneAccessGroupUpdateResponseResultExcludePajwohLqEmailListRule struct {
	EmailList ZoneAccessGroupUpdateResponseResultExcludePajwohLqEmailListRuleEmailList `json:"email_list,required"`
	JSON      zoneAccessGroupUpdateResponseResultExcludePajwohLqEmailListRuleJSON      `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqEmailListRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqEmailListRule]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultExcludePajwohLqEmailListRule) implementsZoneAccessGroupUpdateResponseResultExclude() {
}

type ZoneAccessGroupUpdateResponseResultExcludePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                       `json:"id,required"`
	JSON zoneAccessGroupUpdateResponseResultExcludePajwohLqEmailListRuleEmailListJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqEmailListRuleEmailList]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type ZoneAccessGroupUpdateResponseResultExcludePajwohLqDomainRule struct {
	EmailDomain ZoneAccessGroupUpdateResponseResultExcludePajwohLqDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zoneAccessGroupUpdateResponseResultExcludePajwohLqDomainRuleJSON        `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqDomainRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqDomainRule]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultExcludePajwohLqDomainRule) implementsZoneAccessGroupUpdateResponseResultExclude() {
}

type ZoneAccessGroupUpdateResponseResultExcludePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                      `json:"domain,required"`
	JSON   zoneAccessGroupUpdateResponseResultExcludePajwohLqDomainRuleEmailDomainJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqDomainRuleEmailDomain]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type ZoneAccessGroupUpdateResponseResultExcludePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                        `json:"everyone,required"`
	JSON     zoneAccessGroupUpdateResponseResultExcludePajwohLqEveryoneRuleJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqEveryoneRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqEveryoneRule]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultExcludePajwohLqEveryoneRule) implementsZoneAccessGroupUpdateResponseResultExclude() {
}

// Matches an IP address block.
type ZoneAccessGroupUpdateResponseResultExcludePajwohLqIPRule struct {
	IP   ZoneAccessGroupUpdateResponseResultExcludePajwohLqIPRuleIP   `json:"ip,required"`
	JSON zoneAccessGroupUpdateResponseResultExcludePajwohLqIPRuleJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqIPRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqIPRule]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultExcludePajwohLqIPRule) implementsZoneAccessGroupUpdateResponseResultExclude() {
}

type ZoneAccessGroupUpdateResponseResultExcludePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                         `json:"ip,required"`
	JSON zoneAccessGroupUpdateResponseResultExcludePajwohLqIPRuleIPJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqIPRuleIPJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqIPRuleIP]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type ZoneAccessGroupUpdateResponseResultExcludePajwohLqIPListRule struct {
	IPList ZoneAccessGroupUpdateResponseResultExcludePajwohLqIPListRuleIPList `json:"ip_list,required"`
	JSON   zoneAccessGroupUpdateResponseResultExcludePajwohLqIPListRuleJSON   `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqIPListRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqIPListRule]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultExcludePajwohLqIPListRule) implementsZoneAccessGroupUpdateResponseResultExclude() {
}

type ZoneAccessGroupUpdateResponseResultExcludePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                 `json:"id,required"`
	JSON zoneAccessGroupUpdateResponseResultExcludePajwohLqIPListRuleIPListJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqIPListRuleIPListJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqIPListRuleIPList]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type ZoneAccessGroupUpdateResponseResultExcludePajwohLqCertificateRule struct {
	Certificate interface{}                                                           `json:"certificate,required"`
	JSON        zoneAccessGroupUpdateResponseResultExcludePajwohLqCertificateRuleJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqCertificateRuleJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqCertificateRule]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultExcludePajwohLqCertificateRule) implementsZoneAccessGroupUpdateResponseResultExclude() {
}

// Matches an Access group.
type ZoneAccessGroupUpdateResponseResultExcludePajwohLqAccessGroupRule struct {
	Group ZoneAccessGroupUpdateResponseResultExcludePajwohLqAccessGroupRuleGroup `json:"group,required"`
	JSON  zoneAccessGroupUpdateResponseResultExcludePajwohLqAccessGroupRuleJSON  `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqAccessGroupRuleJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqAccessGroupRule]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultExcludePajwohLqAccessGroupRule) implementsZoneAccessGroupUpdateResponseResultExclude() {
}

type ZoneAccessGroupUpdateResponseResultExcludePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                     `json:"id,required"`
	JSON zoneAccessGroupUpdateResponseResultExcludePajwohLqAccessGroupRuleGroupJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqAccessGroupRuleGroup]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZoneAccessGroupUpdateResponseResultExcludePajwohLqAzureGroupRule struct {
	AzureAd ZoneAccessGroupUpdateResponseResultExcludePajwohLqAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zoneAccessGroupUpdateResponseResultExcludePajwohLqAzureGroupRuleJSON    `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqAzureGroupRuleJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqAzureGroupRule]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultExcludePajwohLqAzureGroupRule) implementsZoneAccessGroupUpdateResponseResultExclude() {
}

type ZoneAccessGroupUpdateResponseResultExcludePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                      `json:"connection_id,required"`
	JSON         zoneAccessGroupUpdateResponseResultExcludePajwohLqAzureGroupRuleAzureAdJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqAzureGroupRuleAzureAd]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZoneAccessGroupUpdateResponseResultExcludePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization ZoneAccessGroupUpdateResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zoneAccessGroupUpdateResponseResultExcludePajwohLqGitHubOrganizationRuleJSON               `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqGitHubOrganizationRule]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultExcludePajwohLqGitHubOrganizationRule) implementsZoneAccessGroupUpdateResponseResultExclude() {
}

type ZoneAccessGroupUpdateResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                         `json:"name,required"`
	JSON zoneAccessGroupUpdateResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZoneAccessGroupUpdateResponseResultExcludePajwohLqGsuiteGroupRule struct {
	Gsuite ZoneAccessGroupUpdateResponseResultExcludePajwohLqGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zoneAccessGroupUpdateResponseResultExcludePajwohLqGsuiteGroupRuleJSON   `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqGsuiteGroupRuleJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqGsuiteGroupRule]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultExcludePajwohLqGsuiteGroupRule) implementsZoneAccessGroupUpdateResponseResultExclude() {
}

type ZoneAccessGroupUpdateResponseResultExcludePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                      `json:"email,required"`
	JSON  zoneAccessGroupUpdateResponseResultExcludePajwohLqGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqGsuiteGroupRuleGsuite]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZoneAccessGroupUpdateResponseResultExcludePajwohLqOktaGroupRule struct {
	Okta ZoneAccessGroupUpdateResponseResultExcludePajwohLqOktaGroupRuleOkta `json:"okta,required"`
	JSON zoneAccessGroupUpdateResponseResultExcludePajwohLqOktaGroupRuleJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqOktaGroupRule]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultExcludePajwohLqOktaGroupRule) implementsZoneAccessGroupUpdateResponseResultExclude() {
}

type ZoneAccessGroupUpdateResponseResultExcludePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                  `json:"email,required"`
	JSON  zoneAccessGroupUpdateResponseResultExcludePajwohLqOktaGroupRuleOktaJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqOktaGroupRuleOkta]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZoneAccessGroupUpdateResponseResultExcludePajwohLqSamlGroupRule struct {
	Saml ZoneAccessGroupUpdateResponseResultExcludePajwohLqSamlGroupRuleSaml `json:"saml,required"`
	JSON zoneAccessGroupUpdateResponseResultExcludePajwohLqSamlGroupRuleJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqSamlGroupRule]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultExcludePajwohLqSamlGroupRule) implementsZoneAccessGroupUpdateResponseResultExclude() {
}

type ZoneAccessGroupUpdateResponseResultExcludePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                  `json:"attribute_value,required"`
	JSON           zoneAccessGroupUpdateResponseResultExcludePajwohLqSamlGroupRuleSamlJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqSamlGroupRuleSaml]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type ZoneAccessGroupUpdateResponseResultExcludePajwohLqServiceTokenRule struct {
	ServiceToken ZoneAccessGroupUpdateResponseResultExcludePajwohLqServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zoneAccessGroupUpdateResponseResultExcludePajwohLqServiceTokenRuleJSON         `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqServiceTokenRuleJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqServiceTokenRule]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultExcludePajwohLqServiceTokenRule) implementsZoneAccessGroupUpdateResponseResultExclude() {
}

type ZoneAccessGroupUpdateResponseResultExcludePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                             `json:"token_id,required"`
	JSON    zoneAccessGroupUpdateResponseResultExcludePajwohLqServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqServiceTokenRuleServiceToken]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type ZoneAccessGroupUpdateResponseResultExcludePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                    `json:"any_valid_service_token,required"`
	JSON                 zoneAccessGroupUpdateResponseResultExcludePajwohLqAnyValidServiceTokenRuleJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqAnyValidServiceTokenRule]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultExcludePajwohLqAnyValidServiceTokenRule) implementsZoneAccessGroupUpdateResponseResultExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZoneAccessGroupUpdateResponseResultExcludePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation ZoneAccessGroupUpdateResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zoneAccessGroupUpdateResponseResultExcludePajwohLqExternalEvaluationRuleJSON               `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqExternalEvaluationRule]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultExcludePajwohLqExternalEvaluationRule) implementsZoneAccessGroupUpdateResponseResultExclude() {
}

type ZoneAccessGroupUpdateResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                         `json:"keys_url,required"`
	JSON    zoneAccessGroupUpdateResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type ZoneAccessGroupUpdateResponseResultExcludePajwohLqCountryRule struct {
	Geo  ZoneAccessGroupUpdateResponseResultExcludePajwohLqCountryRuleGeo  `json:"geo,required"`
	JSON zoneAccessGroupUpdateResponseResultExcludePajwohLqCountryRuleJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqCountryRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqCountryRule]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultExcludePajwohLqCountryRule) implementsZoneAccessGroupUpdateResponseResultExclude() {
}

type ZoneAccessGroupUpdateResponseResultExcludePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                               `json:"country_code,required"`
	JSON        zoneAccessGroupUpdateResponseResultExcludePajwohLqCountryRuleGeoJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqCountryRuleGeoJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqCountryRuleGeo]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type ZoneAccessGroupUpdateResponseResultExcludePajwohLqAuthenticationMethodRule struct {
	AuthMethod ZoneAccessGroupUpdateResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zoneAccessGroupUpdateResponseResultExcludePajwohLqAuthenticationMethodRuleJSON       `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqAuthenticationMethodRule]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultExcludePajwohLqAuthenticationMethodRule) implementsZoneAccessGroupUpdateResponseResultExclude() {
}

type ZoneAccessGroupUpdateResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                   `json:"auth_method,required"`
	JSON       zoneAccessGroupUpdateResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type ZoneAccessGroupUpdateResponseResultExcludePajwohLqDevicePostureRule struct {
	DevicePosture ZoneAccessGroupUpdateResponseResultExcludePajwohLqDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zoneAccessGroupUpdateResponseResultExcludePajwohLqDevicePostureRuleJSON          `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqDevicePostureRule]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultExcludePajwohLqDevicePostureRule) implementsZoneAccessGroupUpdateResponseResultExclude() {
}

type ZoneAccessGroupUpdateResponseResultExcludePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                               `json:"integration_uid,required"`
	JSON           zoneAccessGroupUpdateResponseResultExcludePajwohLqDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultExcludePajwohLqDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultExcludePajwohLqDevicePostureRuleDevicePosture]
type zoneAccessGroupUpdateResponseResultExcludePajwohLqDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultExcludePajwohLqDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqEmailRule],
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqEmailListRule],
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqDomainRule],
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqEveryoneRule],
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqIPRule],
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqIPListRule],
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqCertificateRule],
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqAccessGroupRule],
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqAzureGroupRule],
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqGitHubOrganizationRule],
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqGsuiteGroupRule],
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqOktaGroupRule],
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqSamlGroupRule],
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqServiceTokenRule],
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqAnyValidServiceTokenRule],
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqExternalEvaluationRule],
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqCountryRule],
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqAuthenticationMethodRule] or
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqDevicePostureRule].
type ZoneAccessGroupUpdateResponseResultInclude interface {
	implementsZoneAccessGroupUpdateResponseResultInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneAccessGroupUpdateResponseResultInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type ZoneAccessGroupUpdateResponseResultIncludePajwohLqEmailRule struct {
	Email ZoneAccessGroupUpdateResponseResultIncludePajwohLqEmailRuleEmail `json:"email,required"`
	JSON  zoneAccessGroupUpdateResponseResultIncludePajwohLqEmailRuleJSON  `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqEmailRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqEmailRule]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultIncludePajwohLqEmailRule) implementsZoneAccessGroupUpdateResponseResultInclude() {
}

type ZoneAccessGroupUpdateResponseResultIncludePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email string                                                               `json:"email,required" format:"email"`
	JSON  zoneAccessGroupUpdateResponseResultIncludePajwohLqEmailRuleEmailJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqEmailRuleEmailJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqEmailRuleEmail]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type ZoneAccessGroupUpdateResponseResultIncludePajwohLqEmailListRule struct {
	EmailList ZoneAccessGroupUpdateResponseResultIncludePajwohLqEmailListRuleEmailList `json:"email_list,required"`
	JSON      zoneAccessGroupUpdateResponseResultIncludePajwohLqEmailListRuleJSON      `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqEmailListRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqEmailListRule]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultIncludePajwohLqEmailListRule) implementsZoneAccessGroupUpdateResponseResultInclude() {
}

type ZoneAccessGroupUpdateResponseResultIncludePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                       `json:"id,required"`
	JSON zoneAccessGroupUpdateResponseResultIncludePajwohLqEmailListRuleEmailListJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqEmailListRuleEmailList]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type ZoneAccessGroupUpdateResponseResultIncludePajwohLqDomainRule struct {
	EmailDomain ZoneAccessGroupUpdateResponseResultIncludePajwohLqDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zoneAccessGroupUpdateResponseResultIncludePajwohLqDomainRuleJSON        `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqDomainRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqDomainRule]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultIncludePajwohLqDomainRule) implementsZoneAccessGroupUpdateResponseResultInclude() {
}

type ZoneAccessGroupUpdateResponseResultIncludePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                      `json:"domain,required"`
	JSON   zoneAccessGroupUpdateResponseResultIncludePajwohLqDomainRuleEmailDomainJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqDomainRuleEmailDomain]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type ZoneAccessGroupUpdateResponseResultIncludePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                        `json:"everyone,required"`
	JSON     zoneAccessGroupUpdateResponseResultIncludePajwohLqEveryoneRuleJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqEveryoneRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqEveryoneRule]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultIncludePajwohLqEveryoneRule) implementsZoneAccessGroupUpdateResponseResultInclude() {
}

// Matches an IP address block.
type ZoneAccessGroupUpdateResponseResultIncludePajwohLqIPRule struct {
	IP   ZoneAccessGroupUpdateResponseResultIncludePajwohLqIPRuleIP   `json:"ip,required"`
	JSON zoneAccessGroupUpdateResponseResultIncludePajwohLqIPRuleJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqIPRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqIPRule]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultIncludePajwohLqIPRule) implementsZoneAccessGroupUpdateResponseResultInclude() {
}

type ZoneAccessGroupUpdateResponseResultIncludePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                         `json:"ip,required"`
	JSON zoneAccessGroupUpdateResponseResultIncludePajwohLqIPRuleIPJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqIPRuleIPJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqIPRuleIP]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type ZoneAccessGroupUpdateResponseResultIncludePajwohLqIPListRule struct {
	IPList ZoneAccessGroupUpdateResponseResultIncludePajwohLqIPListRuleIPList `json:"ip_list,required"`
	JSON   zoneAccessGroupUpdateResponseResultIncludePajwohLqIPListRuleJSON   `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqIPListRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqIPListRule]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultIncludePajwohLqIPListRule) implementsZoneAccessGroupUpdateResponseResultInclude() {
}

type ZoneAccessGroupUpdateResponseResultIncludePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                 `json:"id,required"`
	JSON zoneAccessGroupUpdateResponseResultIncludePajwohLqIPListRuleIPListJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqIPListRuleIPListJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqIPListRuleIPList]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type ZoneAccessGroupUpdateResponseResultIncludePajwohLqCertificateRule struct {
	Certificate interface{}                                                           `json:"certificate,required"`
	JSON        zoneAccessGroupUpdateResponseResultIncludePajwohLqCertificateRuleJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqCertificateRuleJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqCertificateRule]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultIncludePajwohLqCertificateRule) implementsZoneAccessGroupUpdateResponseResultInclude() {
}

// Matches an Access group.
type ZoneAccessGroupUpdateResponseResultIncludePajwohLqAccessGroupRule struct {
	Group ZoneAccessGroupUpdateResponseResultIncludePajwohLqAccessGroupRuleGroup `json:"group,required"`
	JSON  zoneAccessGroupUpdateResponseResultIncludePajwohLqAccessGroupRuleJSON  `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqAccessGroupRuleJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqAccessGroupRule]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultIncludePajwohLqAccessGroupRule) implementsZoneAccessGroupUpdateResponseResultInclude() {
}

type ZoneAccessGroupUpdateResponseResultIncludePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                     `json:"id,required"`
	JSON zoneAccessGroupUpdateResponseResultIncludePajwohLqAccessGroupRuleGroupJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqAccessGroupRuleGroup]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZoneAccessGroupUpdateResponseResultIncludePajwohLqAzureGroupRule struct {
	AzureAd ZoneAccessGroupUpdateResponseResultIncludePajwohLqAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zoneAccessGroupUpdateResponseResultIncludePajwohLqAzureGroupRuleJSON    `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqAzureGroupRuleJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqAzureGroupRule]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultIncludePajwohLqAzureGroupRule) implementsZoneAccessGroupUpdateResponseResultInclude() {
}

type ZoneAccessGroupUpdateResponseResultIncludePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                      `json:"connection_id,required"`
	JSON         zoneAccessGroupUpdateResponseResultIncludePajwohLqAzureGroupRuleAzureAdJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqAzureGroupRuleAzureAd]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZoneAccessGroupUpdateResponseResultIncludePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization ZoneAccessGroupUpdateResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zoneAccessGroupUpdateResponseResultIncludePajwohLqGitHubOrganizationRuleJSON               `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqGitHubOrganizationRule]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultIncludePajwohLqGitHubOrganizationRule) implementsZoneAccessGroupUpdateResponseResultInclude() {
}

type ZoneAccessGroupUpdateResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                         `json:"name,required"`
	JSON zoneAccessGroupUpdateResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZoneAccessGroupUpdateResponseResultIncludePajwohLqGsuiteGroupRule struct {
	Gsuite ZoneAccessGroupUpdateResponseResultIncludePajwohLqGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zoneAccessGroupUpdateResponseResultIncludePajwohLqGsuiteGroupRuleJSON   `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqGsuiteGroupRuleJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqGsuiteGroupRule]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultIncludePajwohLqGsuiteGroupRule) implementsZoneAccessGroupUpdateResponseResultInclude() {
}

type ZoneAccessGroupUpdateResponseResultIncludePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                      `json:"email,required"`
	JSON  zoneAccessGroupUpdateResponseResultIncludePajwohLqGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqGsuiteGroupRuleGsuite]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZoneAccessGroupUpdateResponseResultIncludePajwohLqOktaGroupRule struct {
	Okta ZoneAccessGroupUpdateResponseResultIncludePajwohLqOktaGroupRuleOkta `json:"okta,required"`
	JSON zoneAccessGroupUpdateResponseResultIncludePajwohLqOktaGroupRuleJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqOktaGroupRule]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultIncludePajwohLqOktaGroupRule) implementsZoneAccessGroupUpdateResponseResultInclude() {
}

type ZoneAccessGroupUpdateResponseResultIncludePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                  `json:"email,required"`
	JSON  zoneAccessGroupUpdateResponseResultIncludePajwohLqOktaGroupRuleOktaJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqOktaGroupRuleOkta]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZoneAccessGroupUpdateResponseResultIncludePajwohLqSamlGroupRule struct {
	Saml ZoneAccessGroupUpdateResponseResultIncludePajwohLqSamlGroupRuleSaml `json:"saml,required"`
	JSON zoneAccessGroupUpdateResponseResultIncludePajwohLqSamlGroupRuleJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqSamlGroupRule]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultIncludePajwohLqSamlGroupRule) implementsZoneAccessGroupUpdateResponseResultInclude() {
}

type ZoneAccessGroupUpdateResponseResultIncludePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                  `json:"attribute_value,required"`
	JSON           zoneAccessGroupUpdateResponseResultIncludePajwohLqSamlGroupRuleSamlJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqSamlGroupRuleSaml]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type ZoneAccessGroupUpdateResponseResultIncludePajwohLqServiceTokenRule struct {
	ServiceToken ZoneAccessGroupUpdateResponseResultIncludePajwohLqServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zoneAccessGroupUpdateResponseResultIncludePajwohLqServiceTokenRuleJSON         `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqServiceTokenRuleJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqServiceTokenRule]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultIncludePajwohLqServiceTokenRule) implementsZoneAccessGroupUpdateResponseResultInclude() {
}

type ZoneAccessGroupUpdateResponseResultIncludePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                             `json:"token_id,required"`
	JSON    zoneAccessGroupUpdateResponseResultIncludePajwohLqServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqServiceTokenRuleServiceToken]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type ZoneAccessGroupUpdateResponseResultIncludePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                    `json:"any_valid_service_token,required"`
	JSON                 zoneAccessGroupUpdateResponseResultIncludePajwohLqAnyValidServiceTokenRuleJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqAnyValidServiceTokenRule]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultIncludePajwohLqAnyValidServiceTokenRule) implementsZoneAccessGroupUpdateResponseResultInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZoneAccessGroupUpdateResponseResultIncludePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation ZoneAccessGroupUpdateResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zoneAccessGroupUpdateResponseResultIncludePajwohLqExternalEvaluationRuleJSON               `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqExternalEvaluationRule]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultIncludePajwohLqExternalEvaluationRule) implementsZoneAccessGroupUpdateResponseResultInclude() {
}

type ZoneAccessGroupUpdateResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                         `json:"keys_url,required"`
	JSON    zoneAccessGroupUpdateResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type ZoneAccessGroupUpdateResponseResultIncludePajwohLqCountryRule struct {
	Geo  ZoneAccessGroupUpdateResponseResultIncludePajwohLqCountryRuleGeo  `json:"geo,required"`
	JSON zoneAccessGroupUpdateResponseResultIncludePajwohLqCountryRuleJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqCountryRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqCountryRule]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultIncludePajwohLqCountryRule) implementsZoneAccessGroupUpdateResponseResultInclude() {
}

type ZoneAccessGroupUpdateResponseResultIncludePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                               `json:"country_code,required"`
	JSON        zoneAccessGroupUpdateResponseResultIncludePajwohLqCountryRuleGeoJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqCountryRuleGeoJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqCountryRuleGeo]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type ZoneAccessGroupUpdateResponseResultIncludePajwohLqAuthenticationMethodRule struct {
	AuthMethod ZoneAccessGroupUpdateResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zoneAccessGroupUpdateResponseResultIncludePajwohLqAuthenticationMethodRuleJSON       `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqAuthenticationMethodRule]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultIncludePajwohLqAuthenticationMethodRule) implementsZoneAccessGroupUpdateResponseResultInclude() {
}

type ZoneAccessGroupUpdateResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                   `json:"auth_method,required"`
	JSON       zoneAccessGroupUpdateResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type ZoneAccessGroupUpdateResponseResultIncludePajwohLqDevicePostureRule struct {
	DevicePosture ZoneAccessGroupUpdateResponseResultIncludePajwohLqDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zoneAccessGroupUpdateResponseResultIncludePajwohLqDevicePostureRuleJSON          `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqDevicePostureRule]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultIncludePajwohLqDevicePostureRule) implementsZoneAccessGroupUpdateResponseResultInclude() {
}

type ZoneAccessGroupUpdateResponseResultIncludePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                               `json:"integration_uid,required"`
	JSON           zoneAccessGroupUpdateResponseResultIncludePajwohLqDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultIncludePajwohLqDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultIncludePajwohLqDevicePostureRuleDevicePosture]
type zoneAccessGroupUpdateResponseResultIncludePajwohLqDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultIncludePajwohLqDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqEmailRule],
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqEmailListRule],
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqDomainRule],
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqEveryoneRule],
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqIPRule],
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqIPListRule],
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqCertificateRule],
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqAccessGroupRule],
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqAzureGroupRule],
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqGitHubOrganizationRule],
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqGsuiteGroupRule],
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqOktaGroupRule],
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqSamlGroupRule],
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqServiceTokenRule],
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqAnyValidServiceTokenRule],
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqExternalEvaluationRule],
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqCountryRule],
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqAuthenticationMethodRule] or
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqDevicePostureRule].
type ZoneAccessGroupUpdateResponseResultRequire interface {
	implementsZoneAccessGroupUpdateResponseResultRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneAccessGroupUpdateResponseResultRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type ZoneAccessGroupUpdateResponseResultRequirePajwohLqEmailRule struct {
	Email ZoneAccessGroupUpdateResponseResultRequirePajwohLqEmailRuleEmail `json:"email,required"`
	JSON  zoneAccessGroupUpdateResponseResultRequirePajwohLqEmailRuleJSON  `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqEmailRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqEmailRule]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultRequirePajwohLqEmailRule) implementsZoneAccessGroupUpdateResponseResultRequire() {
}

type ZoneAccessGroupUpdateResponseResultRequirePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email string                                                               `json:"email,required" format:"email"`
	JSON  zoneAccessGroupUpdateResponseResultRequirePajwohLqEmailRuleEmailJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqEmailRuleEmailJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqEmailRuleEmail]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type ZoneAccessGroupUpdateResponseResultRequirePajwohLqEmailListRule struct {
	EmailList ZoneAccessGroupUpdateResponseResultRequirePajwohLqEmailListRuleEmailList `json:"email_list,required"`
	JSON      zoneAccessGroupUpdateResponseResultRequirePajwohLqEmailListRuleJSON      `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqEmailListRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqEmailListRule]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultRequirePajwohLqEmailListRule) implementsZoneAccessGroupUpdateResponseResultRequire() {
}

type ZoneAccessGroupUpdateResponseResultRequirePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                       `json:"id,required"`
	JSON zoneAccessGroupUpdateResponseResultRequirePajwohLqEmailListRuleEmailListJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqEmailListRuleEmailList]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type ZoneAccessGroupUpdateResponseResultRequirePajwohLqDomainRule struct {
	EmailDomain ZoneAccessGroupUpdateResponseResultRequirePajwohLqDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zoneAccessGroupUpdateResponseResultRequirePajwohLqDomainRuleJSON        `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqDomainRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqDomainRule]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultRequirePajwohLqDomainRule) implementsZoneAccessGroupUpdateResponseResultRequire() {
}

type ZoneAccessGroupUpdateResponseResultRequirePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                      `json:"domain,required"`
	JSON   zoneAccessGroupUpdateResponseResultRequirePajwohLqDomainRuleEmailDomainJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqDomainRuleEmailDomain]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type ZoneAccessGroupUpdateResponseResultRequirePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                        `json:"everyone,required"`
	JSON     zoneAccessGroupUpdateResponseResultRequirePajwohLqEveryoneRuleJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqEveryoneRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqEveryoneRule]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultRequirePajwohLqEveryoneRule) implementsZoneAccessGroupUpdateResponseResultRequire() {
}

// Matches an IP address block.
type ZoneAccessGroupUpdateResponseResultRequirePajwohLqIPRule struct {
	IP   ZoneAccessGroupUpdateResponseResultRequirePajwohLqIPRuleIP   `json:"ip,required"`
	JSON zoneAccessGroupUpdateResponseResultRequirePajwohLqIPRuleJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqIPRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqIPRule]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultRequirePajwohLqIPRule) implementsZoneAccessGroupUpdateResponseResultRequire() {
}

type ZoneAccessGroupUpdateResponseResultRequirePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                         `json:"ip,required"`
	JSON zoneAccessGroupUpdateResponseResultRequirePajwohLqIPRuleIPJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqIPRuleIPJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqIPRuleIP]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type ZoneAccessGroupUpdateResponseResultRequirePajwohLqIPListRule struct {
	IPList ZoneAccessGroupUpdateResponseResultRequirePajwohLqIPListRuleIPList `json:"ip_list,required"`
	JSON   zoneAccessGroupUpdateResponseResultRequirePajwohLqIPListRuleJSON   `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqIPListRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqIPListRule]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultRequirePajwohLqIPListRule) implementsZoneAccessGroupUpdateResponseResultRequire() {
}

type ZoneAccessGroupUpdateResponseResultRequirePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                 `json:"id,required"`
	JSON zoneAccessGroupUpdateResponseResultRequirePajwohLqIPListRuleIPListJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqIPListRuleIPListJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqIPListRuleIPList]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type ZoneAccessGroupUpdateResponseResultRequirePajwohLqCertificateRule struct {
	Certificate interface{}                                                           `json:"certificate,required"`
	JSON        zoneAccessGroupUpdateResponseResultRequirePajwohLqCertificateRuleJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqCertificateRuleJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqCertificateRule]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultRequirePajwohLqCertificateRule) implementsZoneAccessGroupUpdateResponseResultRequire() {
}

// Matches an Access group.
type ZoneAccessGroupUpdateResponseResultRequirePajwohLqAccessGroupRule struct {
	Group ZoneAccessGroupUpdateResponseResultRequirePajwohLqAccessGroupRuleGroup `json:"group,required"`
	JSON  zoneAccessGroupUpdateResponseResultRequirePajwohLqAccessGroupRuleJSON  `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqAccessGroupRuleJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqAccessGroupRule]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultRequirePajwohLqAccessGroupRule) implementsZoneAccessGroupUpdateResponseResultRequire() {
}

type ZoneAccessGroupUpdateResponseResultRequirePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                     `json:"id,required"`
	JSON zoneAccessGroupUpdateResponseResultRequirePajwohLqAccessGroupRuleGroupJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqAccessGroupRuleGroup]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZoneAccessGroupUpdateResponseResultRequirePajwohLqAzureGroupRule struct {
	AzureAd ZoneAccessGroupUpdateResponseResultRequirePajwohLqAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zoneAccessGroupUpdateResponseResultRequirePajwohLqAzureGroupRuleJSON    `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqAzureGroupRuleJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqAzureGroupRule]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultRequirePajwohLqAzureGroupRule) implementsZoneAccessGroupUpdateResponseResultRequire() {
}

type ZoneAccessGroupUpdateResponseResultRequirePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                      `json:"connection_id,required"`
	JSON         zoneAccessGroupUpdateResponseResultRequirePajwohLqAzureGroupRuleAzureAdJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqAzureGroupRuleAzureAd]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZoneAccessGroupUpdateResponseResultRequirePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization ZoneAccessGroupUpdateResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zoneAccessGroupUpdateResponseResultRequirePajwohLqGitHubOrganizationRuleJSON               `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqGitHubOrganizationRule]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultRequirePajwohLqGitHubOrganizationRule) implementsZoneAccessGroupUpdateResponseResultRequire() {
}

type ZoneAccessGroupUpdateResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                         `json:"name,required"`
	JSON zoneAccessGroupUpdateResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZoneAccessGroupUpdateResponseResultRequirePajwohLqGsuiteGroupRule struct {
	Gsuite ZoneAccessGroupUpdateResponseResultRequirePajwohLqGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zoneAccessGroupUpdateResponseResultRequirePajwohLqGsuiteGroupRuleJSON   `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqGsuiteGroupRuleJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqGsuiteGroupRule]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultRequirePajwohLqGsuiteGroupRule) implementsZoneAccessGroupUpdateResponseResultRequire() {
}

type ZoneAccessGroupUpdateResponseResultRequirePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                      `json:"email,required"`
	JSON  zoneAccessGroupUpdateResponseResultRequirePajwohLqGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqGsuiteGroupRuleGsuite]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZoneAccessGroupUpdateResponseResultRequirePajwohLqOktaGroupRule struct {
	Okta ZoneAccessGroupUpdateResponseResultRequirePajwohLqOktaGroupRuleOkta `json:"okta,required"`
	JSON zoneAccessGroupUpdateResponseResultRequirePajwohLqOktaGroupRuleJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqOktaGroupRule]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultRequirePajwohLqOktaGroupRule) implementsZoneAccessGroupUpdateResponseResultRequire() {
}

type ZoneAccessGroupUpdateResponseResultRequirePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                  `json:"email,required"`
	JSON  zoneAccessGroupUpdateResponseResultRequirePajwohLqOktaGroupRuleOktaJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqOktaGroupRuleOkta]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZoneAccessGroupUpdateResponseResultRequirePajwohLqSamlGroupRule struct {
	Saml ZoneAccessGroupUpdateResponseResultRequirePajwohLqSamlGroupRuleSaml `json:"saml,required"`
	JSON zoneAccessGroupUpdateResponseResultRequirePajwohLqSamlGroupRuleJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqSamlGroupRule]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultRequirePajwohLqSamlGroupRule) implementsZoneAccessGroupUpdateResponseResultRequire() {
}

type ZoneAccessGroupUpdateResponseResultRequirePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                  `json:"attribute_value,required"`
	JSON           zoneAccessGroupUpdateResponseResultRequirePajwohLqSamlGroupRuleSamlJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqSamlGroupRuleSaml]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type ZoneAccessGroupUpdateResponseResultRequirePajwohLqServiceTokenRule struct {
	ServiceToken ZoneAccessGroupUpdateResponseResultRequirePajwohLqServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zoneAccessGroupUpdateResponseResultRequirePajwohLqServiceTokenRuleJSON         `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqServiceTokenRuleJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqServiceTokenRule]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultRequirePajwohLqServiceTokenRule) implementsZoneAccessGroupUpdateResponseResultRequire() {
}

type ZoneAccessGroupUpdateResponseResultRequirePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                             `json:"token_id,required"`
	JSON    zoneAccessGroupUpdateResponseResultRequirePajwohLqServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqServiceTokenRuleServiceToken]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type ZoneAccessGroupUpdateResponseResultRequirePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                    `json:"any_valid_service_token,required"`
	JSON                 zoneAccessGroupUpdateResponseResultRequirePajwohLqAnyValidServiceTokenRuleJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqAnyValidServiceTokenRule]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultRequirePajwohLqAnyValidServiceTokenRule) implementsZoneAccessGroupUpdateResponseResultRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZoneAccessGroupUpdateResponseResultRequirePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation ZoneAccessGroupUpdateResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zoneAccessGroupUpdateResponseResultRequirePajwohLqExternalEvaluationRuleJSON               `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqExternalEvaluationRule]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultRequirePajwohLqExternalEvaluationRule) implementsZoneAccessGroupUpdateResponseResultRequire() {
}

type ZoneAccessGroupUpdateResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                         `json:"keys_url,required"`
	JSON    zoneAccessGroupUpdateResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type ZoneAccessGroupUpdateResponseResultRequirePajwohLqCountryRule struct {
	Geo  ZoneAccessGroupUpdateResponseResultRequirePajwohLqCountryRuleGeo  `json:"geo,required"`
	JSON zoneAccessGroupUpdateResponseResultRequirePajwohLqCountryRuleJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqCountryRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqCountryRule]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultRequirePajwohLqCountryRule) implementsZoneAccessGroupUpdateResponseResultRequire() {
}

type ZoneAccessGroupUpdateResponseResultRequirePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                               `json:"country_code,required"`
	JSON        zoneAccessGroupUpdateResponseResultRequirePajwohLqCountryRuleGeoJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqCountryRuleGeoJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqCountryRuleGeo]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type ZoneAccessGroupUpdateResponseResultRequirePajwohLqAuthenticationMethodRule struct {
	AuthMethod ZoneAccessGroupUpdateResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zoneAccessGroupUpdateResponseResultRequirePajwohLqAuthenticationMethodRuleJSON       `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqAuthenticationMethodRule]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultRequirePajwohLqAuthenticationMethodRule) implementsZoneAccessGroupUpdateResponseResultRequire() {
}

type ZoneAccessGroupUpdateResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                   `json:"auth_method,required"`
	JSON       zoneAccessGroupUpdateResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type ZoneAccessGroupUpdateResponseResultRequirePajwohLqDevicePostureRule struct {
	DevicePosture ZoneAccessGroupUpdateResponseResultRequirePajwohLqDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zoneAccessGroupUpdateResponseResultRequirePajwohLqDevicePostureRuleJSON          `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqDevicePostureRule]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupUpdateResponseResultRequirePajwohLqDevicePostureRule) implementsZoneAccessGroupUpdateResponseResultRequire() {
}

type ZoneAccessGroupUpdateResponseResultRequirePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                               `json:"integration_uid,required"`
	JSON           zoneAccessGroupUpdateResponseResultRequirePajwohLqDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zoneAccessGroupUpdateResponseResultRequirePajwohLqDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupUpdateResponseResultRequirePajwohLqDevicePostureRuleDevicePosture]
type zoneAccessGroupUpdateResponseResultRequirePajwohLqDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZoneAccessGroupUpdateResponseResultRequirePajwohLqDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneAccessGroupUpdateResponseSuccess bool

const (
	ZoneAccessGroupUpdateResponseSuccessTrue ZoneAccessGroupUpdateResponseSuccess = true
)

type ZoneAccessGroupListResponse struct {
	Errors     []ZoneAccessGroupListResponseError    `json:"errors"`
	Messages   []ZoneAccessGroupListResponseMessage  `json:"messages"`
	Result     []ZoneAccessGroupListResponseResult   `json:"result"`
	ResultInfo ZoneAccessGroupListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneAccessGroupListResponseSuccess `json:"success"`
	JSON    zoneAccessGroupListResponseJSON    `json:"-"`
}

// zoneAccessGroupListResponseJSON contains the JSON metadata for the struct
// [ZoneAccessGroupListResponse]
type zoneAccessGroupListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessGroupListResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    zoneAccessGroupListResponseErrorJSON `json:"-"`
}

// zoneAccessGroupListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseError]
type zoneAccessGroupListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessGroupListResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneAccessGroupListResponseMessageJSON `json:"-"`
}

// zoneAccessGroupListResponseMessageJSON contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseMessage]
type zoneAccessGroupListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessGroupListResponseResult struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []ZoneAccessGroupListResponseResultExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []ZoneAccessGroupListResponseResultInclude `json:"include"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []ZoneAccessGroupListResponseResultRequire `json:"require"`
	UpdatedAt time.Time                                  `json:"updated_at" format:"date-time"`
	JSON      zoneAccessGroupListResponseResultJSON      `json:"-"`
}

// zoneAccessGroupListResponseResultJSON contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResult]
type zoneAccessGroupListResponseResultJSON struct {
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

func (r *ZoneAccessGroupListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [ZoneAccessGroupListResponseResultExcludePajwohLqEmailRule],
// [ZoneAccessGroupListResponseResultExcludePajwohLqEmailListRule],
// [ZoneAccessGroupListResponseResultExcludePajwohLqDomainRule],
// [ZoneAccessGroupListResponseResultExcludePajwohLqEveryoneRule],
// [ZoneAccessGroupListResponseResultExcludePajwohLqIPRule],
// [ZoneAccessGroupListResponseResultExcludePajwohLqIPListRule],
// [ZoneAccessGroupListResponseResultExcludePajwohLqCertificateRule],
// [ZoneAccessGroupListResponseResultExcludePajwohLqAccessGroupRule],
// [ZoneAccessGroupListResponseResultExcludePajwohLqAzureGroupRule],
// [ZoneAccessGroupListResponseResultExcludePajwohLqGitHubOrganizationRule],
// [ZoneAccessGroupListResponseResultExcludePajwohLqGsuiteGroupRule],
// [ZoneAccessGroupListResponseResultExcludePajwohLqOktaGroupRule],
// [ZoneAccessGroupListResponseResultExcludePajwohLqSamlGroupRule],
// [ZoneAccessGroupListResponseResultExcludePajwohLqServiceTokenRule],
// [ZoneAccessGroupListResponseResultExcludePajwohLqAnyValidServiceTokenRule],
// [ZoneAccessGroupListResponseResultExcludePajwohLqExternalEvaluationRule],
// [ZoneAccessGroupListResponseResultExcludePajwohLqCountryRule],
// [ZoneAccessGroupListResponseResultExcludePajwohLqAuthenticationMethodRule] or
// [ZoneAccessGroupListResponseResultExcludePajwohLqDevicePostureRule].
type ZoneAccessGroupListResponseResultExclude interface {
	implementsZoneAccessGroupListResponseResultExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneAccessGroupListResponseResultExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type ZoneAccessGroupListResponseResultExcludePajwohLqEmailRule struct {
	Email ZoneAccessGroupListResponseResultExcludePajwohLqEmailRuleEmail `json:"email,required"`
	JSON  zoneAccessGroupListResponseResultExcludePajwohLqEmailRuleJSON  `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqEmailRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqEmailRule]
type zoneAccessGroupListResponseResultExcludePajwohLqEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultExcludePajwohLqEmailRule) implementsZoneAccessGroupListResponseResultExclude() {
}

type ZoneAccessGroupListResponseResultExcludePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email string                                                             `json:"email,required" format:"email"`
	JSON  zoneAccessGroupListResponseResultExcludePajwohLqEmailRuleEmailJSON `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqEmailRuleEmail]
type zoneAccessGroupListResponseResultExcludePajwohLqEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type ZoneAccessGroupListResponseResultExcludePajwohLqEmailListRule struct {
	EmailList ZoneAccessGroupListResponseResultExcludePajwohLqEmailListRuleEmailList `json:"email_list,required"`
	JSON      zoneAccessGroupListResponseResultExcludePajwohLqEmailListRuleJSON      `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqEmailListRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqEmailListRule]
type zoneAccessGroupListResponseResultExcludePajwohLqEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultExcludePajwohLqEmailListRule) implementsZoneAccessGroupListResponseResultExclude() {
}

type ZoneAccessGroupListResponseResultExcludePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                     `json:"id,required"`
	JSON zoneAccessGroupListResponseResultExcludePajwohLqEmailListRuleEmailListJSON `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqEmailListRuleEmailList]
type zoneAccessGroupListResponseResultExcludePajwohLqEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type ZoneAccessGroupListResponseResultExcludePajwohLqDomainRule struct {
	EmailDomain ZoneAccessGroupListResponseResultExcludePajwohLqDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zoneAccessGroupListResponseResultExcludePajwohLqDomainRuleJSON        `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqDomainRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqDomainRule]
type zoneAccessGroupListResponseResultExcludePajwohLqDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultExcludePajwohLqDomainRule) implementsZoneAccessGroupListResponseResultExclude() {
}

type ZoneAccessGroupListResponseResultExcludePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                    `json:"domain,required"`
	JSON   zoneAccessGroupListResponseResultExcludePajwohLqDomainRuleEmailDomainJSON `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqDomainRuleEmailDomain]
type zoneAccessGroupListResponseResultExcludePajwohLqDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type ZoneAccessGroupListResponseResultExcludePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                      `json:"everyone,required"`
	JSON     zoneAccessGroupListResponseResultExcludePajwohLqEveryoneRuleJSON `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqEveryoneRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqEveryoneRule]
type zoneAccessGroupListResponseResultExcludePajwohLqEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultExcludePajwohLqEveryoneRule) implementsZoneAccessGroupListResponseResultExclude() {
}

// Matches an IP address block.
type ZoneAccessGroupListResponseResultExcludePajwohLqIPRule struct {
	IP   ZoneAccessGroupListResponseResultExcludePajwohLqIPRuleIP   `json:"ip,required"`
	JSON zoneAccessGroupListResponseResultExcludePajwohLqIPRuleJSON `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqIPRuleJSON contains the JSON
// metadata for the struct [ZoneAccessGroupListResponseResultExcludePajwohLqIPRule]
type zoneAccessGroupListResponseResultExcludePajwohLqIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultExcludePajwohLqIPRule) implementsZoneAccessGroupListResponseResultExclude() {
}

type ZoneAccessGroupListResponseResultExcludePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                       `json:"ip,required"`
	JSON zoneAccessGroupListResponseResultExcludePajwohLqIPRuleIPJSON `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqIPRuleIPJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqIPRuleIP]
type zoneAccessGroupListResponseResultExcludePajwohLqIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type ZoneAccessGroupListResponseResultExcludePajwohLqIPListRule struct {
	IPList ZoneAccessGroupListResponseResultExcludePajwohLqIPListRuleIPList `json:"ip_list,required"`
	JSON   zoneAccessGroupListResponseResultExcludePajwohLqIPListRuleJSON   `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqIPListRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqIPListRule]
type zoneAccessGroupListResponseResultExcludePajwohLqIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultExcludePajwohLqIPListRule) implementsZoneAccessGroupListResponseResultExclude() {
}

type ZoneAccessGroupListResponseResultExcludePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                               `json:"id,required"`
	JSON zoneAccessGroupListResponseResultExcludePajwohLqIPListRuleIPListJSON `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqIPListRuleIPListJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqIPListRuleIPList]
type zoneAccessGroupListResponseResultExcludePajwohLqIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type ZoneAccessGroupListResponseResultExcludePajwohLqCertificateRule struct {
	Certificate interface{}                                                         `json:"certificate,required"`
	JSON        zoneAccessGroupListResponseResultExcludePajwohLqCertificateRuleJSON `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqCertificateRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqCertificateRule]
type zoneAccessGroupListResponseResultExcludePajwohLqCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultExcludePajwohLqCertificateRule) implementsZoneAccessGroupListResponseResultExclude() {
}

// Matches an Access group.
type ZoneAccessGroupListResponseResultExcludePajwohLqAccessGroupRule struct {
	Group ZoneAccessGroupListResponseResultExcludePajwohLqAccessGroupRuleGroup `json:"group,required"`
	JSON  zoneAccessGroupListResponseResultExcludePajwohLqAccessGroupRuleJSON  `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqAccessGroupRule]
type zoneAccessGroupListResponseResultExcludePajwohLqAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultExcludePajwohLqAccessGroupRule) implementsZoneAccessGroupListResponseResultExclude() {
}

type ZoneAccessGroupListResponseResultExcludePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                   `json:"id,required"`
	JSON zoneAccessGroupListResponseResultExcludePajwohLqAccessGroupRuleGroupJSON `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqAccessGroupRuleGroup]
type zoneAccessGroupListResponseResultExcludePajwohLqAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZoneAccessGroupListResponseResultExcludePajwohLqAzureGroupRule struct {
	AzureAd ZoneAccessGroupListResponseResultExcludePajwohLqAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zoneAccessGroupListResponseResultExcludePajwohLqAzureGroupRuleJSON    `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqAzureGroupRule]
type zoneAccessGroupListResponseResultExcludePajwohLqAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultExcludePajwohLqAzureGroupRule) implementsZoneAccessGroupListResponseResultExclude() {
}

type ZoneAccessGroupListResponseResultExcludePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                    `json:"connection_id,required"`
	JSON         zoneAccessGroupListResponseResultExcludePajwohLqAzureGroupRuleAzureAdJSON `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqAzureGroupRuleAzureAd]
type zoneAccessGroupListResponseResultExcludePajwohLqAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZoneAccessGroupListResponseResultExcludePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization ZoneAccessGroupListResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zoneAccessGroupListResponseResultExcludePajwohLqGitHubOrganizationRuleJSON               `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqGitHubOrganizationRule]
type zoneAccessGroupListResponseResultExcludePajwohLqGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultExcludePajwohLqGitHubOrganizationRule) implementsZoneAccessGroupListResponseResultExclude() {
}

type ZoneAccessGroupListResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                       `json:"name,required"`
	JSON zoneAccessGroupListResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization]
type zoneAccessGroupListResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZoneAccessGroupListResponseResultExcludePajwohLqGsuiteGroupRule struct {
	Gsuite ZoneAccessGroupListResponseResultExcludePajwohLqGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zoneAccessGroupListResponseResultExcludePajwohLqGsuiteGroupRuleJSON   `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqGsuiteGroupRule]
type zoneAccessGroupListResponseResultExcludePajwohLqGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultExcludePajwohLqGsuiteGroupRule) implementsZoneAccessGroupListResponseResultExclude() {
}

type ZoneAccessGroupListResponseResultExcludePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                    `json:"email,required"`
	JSON  zoneAccessGroupListResponseResultExcludePajwohLqGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqGsuiteGroupRuleGsuite]
type zoneAccessGroupListResponseResultExcludePajwohLqGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZoneAccessGroupListResponseResultExcludePajwohLqOktaGroupRule struct {
	Okta ZoneAccessGroupListResponseResultExcludePajwohLqOktaGroupRuleOkta `json:"okta,required"`
	JSON zoneAccessGroupListResponseResultExcludePajwohLqOktaGroupRuleJSON `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqOktaGroupRule]
type zoneAccessGroupListResponseResultExcludePajwohLqOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultExcludePajwohLqOktaGroupRule) implementsZoneAccessGroupListResponseResultExclude() {
}

type ZoneAccessGroupListResponseResultExcludePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                `json:"email,required"`
	JSON  zoneAccessGroupListResponseResultExcludePajwohLqOktaGroupRuleOktaJSON `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqOktaGroupRuleOkta]
type zoneAccessGroupListResponseResultExcludePajwohLqOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZoneAccessGroupListResponseResultExcludePajwohLqSamlGroupRule struct {
	Saml ZoneAccessGroupListResponseResultExcludePajwohLqSamlGroupRuleSaml `json:"saml,required"`
	JSON zoneAccessGroupListResponseResultExcludePajwohLqSamlGroupRuleJSON `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqSamlGroupRule]
type zoneAccessGroupListResponseResultExcludePajwohLqSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultExcludePajwohLqSamlGroupRule) implementsZoneAccessGroupListResponseResultExclude() {
}

type ZoneAccessGroupListResponseResultExcludePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                `json:"attribute_value,required"`
	JSON           zoneAccessGroupListResponseResultExcludePajwohLqSamlGroupRuleSamlJSON `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqSamlGroupRuleSaml]
type zoneAccessGroupListResponseResultExcludePajwohLqSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type ZoneAccessGroupListResponseResultExcludePajwohLqServiceTokenRule struct {
	ServiceToken ZoneAccessGroupListResponseResultExcludePajwohLqServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zoneAccessGroupListResponseResultExcludePajwohLqServiceTokenRuleJSON         `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqServiceTokenRuleJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqServiceTokenRule]
type zoneAccessGroupListResponseResultExcludePajwohLqServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultExcludePajwohLqServiceTokenRule) implementsZoneAccessGroupListResponseResultExclude() {
}

type ZoneAccessGroupListResponseResultExcludePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                           `json:"token_id,required"`
	JSON    zoneAccessGroupListResponseResultExcludePajwohLqServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqServiceTokenRuleServiceToken]
type zoneAccessGroupListResponseResultExcludePajwohLqServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type ZoneAccessGroupListResponseResultExcludePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                  `json:"any_valid_service_token,required"`
	JSON                 zoneAccessGroupListResponseResultExcludePajwohLqAnyValidServiceTokenRuleJSON `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqAnyValidServiceTokenRule]
type zoneAccessGroupListResponseResultExcludePajwohLqAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultExcludePajwohLqAnyValidServiceTokenRule) implementsZoneAccessGroupListResponseResultExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZoneAccessGroupListResponseResultExcludePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation ZoneAccessGroupListResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zoneAccessGroupListResponseResultExcludePajwohLqExternalEvaluationRuleJSON               `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqExternalEvaluationRule]
type zoneAccessGroupListResponseResultExcludePajwohLqExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultExcludePajwohLqExternalEvaluationRule) implementsZoneAccessGroupListResponseResultExclude() {
}

type ZoneAccessGroupListResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                       `json:"keys_url,required"`
	JSON    zoneAccessGroupListResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation]
type zoneAccessGroupListResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type ZoneAccessGroupListResponseResultExcludePajwohLqCountryRule struct {
	Geo  ZoneAccessGroupListResponseResultExcludePajwohLqCountryRuleGeo  `json:"geo,required"`
	JSON zoneAccessGroupListResponseResultExcludePajwohLqCountryRuleJSON `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqCountryRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqCountryRule]
type zoneAccessGroupListResponseResultExcludePajwohLqCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultExcludePajwohLqCountryRule) implementsZoneAccessGroupListResponseResultExclude() {
}

type ZoneAccessGroupListResponseResultExcludePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                             `json:"country_code,required"`
	JSON        zoneAccessGroupListResponseResultExcludePajwohLqCountryRuleGeoJSON `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqCountryRuleGeo]
type zoneAccessGroupListResponseResultExcludePajwohLqCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type ZoneAccessGroupListResponseResultExcludePajwohLqAuthenticationMethodRule struct {
	AuthMethod ZoneAccessGroupListResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zoneAccessGroupListResponseResultExcludePajwohLqAuthenticationMethodRuleJSON       `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqAuthenticationMethodRule]
type zoneAccessGroupListResponseResultExcludePajwohLqAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultExcludePajwohLqAuthenticationMethodRule) implementsZoneAccessGroupListResponseResultExclude() {
}

type ZoneAccessGroupListResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                 `json:"auth_method,required"`
	JSON       zoneAccessGroupListResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod]
type zoneAccessGroupListResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type ZoneAccessGroupListResponseResultExcludePajwohLqDevicePostureRule struct {
	DevicePosture ZoneAccessGroupListResponseResultExcludePajwohLqDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zoneAccessGroupListResponseResultExcludePajwohLqDevicePostureRuleJSON          `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqDevicePostureRule]
type zoneAccessGroupListResponseResultExcludePajwohLqDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultExcludePajwohLqDevicePostureRule) implementsZoneAccessGroupListResponseResultExclude() {
}

type ZoneAccessGroupListResponseResultExcludePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                             `json:"integration_uid,required"`
	JSON           zoneAccessGroupListResponseResultExcludePajwohLqDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zoneAccessGroupListResponseResultExcludePajwohLqDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultExcludePajwohLqDevicePostureRuleDevicePosture]
type zoneAccessGroupListResponseResultExcludePajwohLqDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultExcludePajwohLqDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [ZoneAccessGroupListResponseResultIncludePajwohLqEmailRule],
// [ZoneAccessGroupListResponseResultIncludePajwohLqEmailListRule],
// [ZoneAccessGroupListResponseResultIncludePajwohLqDomainRule],
// [ZoneAccessGroupListResponseResultIncludePajwohLqEveryoneRule],
// [ZoneAccessGroupListResponseResultIncludePajwohLqIPRule],
// [ZoneAccessGroupListResponseResultIncludePajwohLqIPListRule],
// [ZoneAccessGroupListResponseResultIncludePajwohLqCertificateRule],
// [ZoneAccessGroupListResponseResultIncludePajwohLqAccessGroupRule],
// [ZoneAccessGroupListResponseResultIncludePajwohLqAzureGroupRule],
// [ZoneAccessGroupListResponseResultIncludePajwohLqGitHubOrganizationRule],
// [ZoneAccessGroupListResponseResultIncludePajwohLqGsuiteGroupRule],
// [ZoneAccessGroupListResponseResultIncludePajwohLqOktaGroupRule],
// [ZoneAccessGroupListResponseResultIncludePajwohLqSamlGroupRule],
// [ZoneAccessGroupListResponseResultIncludePajwohLqServiceTokenRule],
// [ZoneAccessGroupListResponseResultIncludePajwohLqAnyValidServiceTokenRule],
// [ZoneAccessGroupListResponseResultIncludePajwohLqExternalEvaluationRule],
// [ZoneAccessGroupListResponseResultIncludePajwohLqCountryRule],
// [ZoneAccessGroupListResponseResultIncludePajwohLqAuthenticationMethodRule] or
// [ZoneAccessGroupListResponseResultIncludePajwohLqDevicePostureRule].
type ZoneAccessGroupListResponseResultInclude interface {
	implementsZoneAccessGroupListResponseResultInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneAccessGroupListResponseResultInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type ZoneAccessGroupListResponseResultIncludePajwohLqEmailRule struct {
	Email ZoneAccessGroupListResponseResultIncludePajwohLqEmailRuleEmail `json:"email,required"`
	JSON  zoneAccessGroupListResponseResultIncludePajwohLqEmailRuleJSON  `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqEmailRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqEmailRule]
type zoneAccessGroupListResponseResultIncludePajwohLqEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultIncludePajwohLqEmailRule) implementsZoneAccessGroupListResponseResultInclude() {
}

type ZoneAccessGroupListResponseResultIncludePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email string                                                             `json:"email,required" format:"email"`
	JSON  zoneAccessGroupListResponseResultIncludePajwohLqEmailRuleEmailJSON `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqEmailRuleEmail]
type zoneAccessGroupListResponseResultIncludePajwohLqEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type ZoneAccessGroupListResponseResultIncludePajwohLqEmailListRule struct {
	EmailList ZoneAccessGroupListResponseResultIncludePajwohLqEmailListRuleEmailList `json:"email_list,required"`
	JSON      zoneAccessGroupListResponseResultIncludePajwohLqEmailListRuleJSON      `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqEmailListRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqEmailListRule]
type zoneAccessGroupListResponseResultIncludePajwohLqEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultIncludePajwohLqEmailListRule) implementsZoneAccessGroupListResponseResultInclude() {
}

type ZoneAccessGroupListResponseResultIncludePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                     `json:"id,required"`
	JSON zoneAccessGroupListResponseResultIncludePajwohLqEmailListRuleEmailListJSON `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqEmailListRuleEmailList]
type zoneAccessGroupListResponseResultIncludePajwohLqEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type ZoneAccessGroupListResponseResultIncludePajwohLqDomainRule struct {
	EmailDomain ZoneAccessGroupListResponseResultIncludePajwohLqDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zoneAccessGroupListResponseResultIncludePajwohLqDomainRuleJSON        `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqDomainRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqDomainRule]
type zoneAccessGroupListResponseResultIncludePajwohLqDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultIncludePajwohLqDomainRule) implementsZoneAccessGroupListResponseResultInclude() {
}

type ZoneAccessGroupListResponseResultIncludePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                    `json:"domain,required"`
	JSON   zoneAccessGroupListResponseResultIncludePajwohLqDomainRuleEmailDomainJSON `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqDomainRuleEmailDomain]
type zoneAccessGroupListResponseResultIncludePajwohLqDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type ZoneAccessGroupListResponseResultIncludePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                      `json:"everyone,required"`
	JSON     zoneAccessGroupListResponseResultIncludePajwohLqEveryoneRuleJSON `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqEveryoneRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqEveryoneRule]
type zoneAccessGroupListResponseResultIncludePajwohLqEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultIncludePajwohLqEveryoneRule) implementsZoneAccessGroupListResponseResultInclude() {
}

// Matches an IP address block.
type ZoneAccessGroupListResponseResultIncludePajwohLqIPRule struct {
	IP   ZoneAccessGroupListResponseResultIncludePajwohLqIPRuleIP   `json:"ip,required"`
	JSON zoneAccessGroupListResponseResultIncludePajwohLqIPRuleJSON `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqIPRuleJSON contains the JSON
// metadata for the struct [ZoneAccessGroupListResponseResultIncludePajwohLqIPRule]
type zoneAccessGroupListResponseResultIncludePajwohLqIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultIncludePajwohLqIPRule) implementsZoneAccessGroupListResponseResultInclude() {
}

type ZoneAccessGroupListResponseResultIncludePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                       `json:"ip,required"`
	JSON zoneAccessGroupListResponseResultIncludePajwohLqIPRuleIPJSON `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqIPRuleIPJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqIPRuleIP]
type zoneAccessGroupListResponseResultIncludePajwohLqIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type ZoneAccessGroupListResponseResultIncludePajwohLqIPListRule struct {
	IPList ZoneAccessGroupListResponseResultIncludePajwohLqIPListRuleIPList `json:"ip_list,required"`
	JSON   zoneAccessGroupListResponseResultIncludePajwohLqIPListRuleJSON   `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqIPListRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqIPListRule]
type zoneAccessGroupListResponseResultIncludePajwohLqIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultIncludePajwohLqIPListRule) implementsZoneAccessGroupListResponseResultInclude() {
}

type ZoneAccessGroupListResponseResultIncludePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                               `json:"id,required"`
	JSON zoneAccessGroupListResponseResultIncludePajwohLqIPListRuleIPListJSON `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqIPListRuleIPListJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqIPListRuleIPList]
type zoneAccessGroupListResponseResultIncludePajwohLqIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type ZoneAccessGroupListResponseResultIncludePajwohLqCertificateRule struct {
	Certificate interface{}                                                         `json:"certificate,required"`
	JSON        zoneAccessGroupListResponseResultIncludePajwohLqCertificateRuleJSON `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqCertificateRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqCertificateRule]
type zoneAccessGroupListResponseResultIncludePajwohLqCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultIncludePajwohLqCertificateRule) implementsZoneAccessGroupListResponseResultInclude() {
}

// Matches an Access group.
type ZoneAccessGroupListResponseResultIncludePajwohLqAccessGroupRule struct {
	Group ZoneAccessGroupListResponseResultIncludePajwohLqAccessGroupRuleGroup `json:"group,required"`
	JSON  zoneAccessGroupListResponseResultIncludePajwohLqAccessGroupRuleJSON  `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqAccessGroupRule]
type zoneAccessGroupListResponseResultIncludePajwohLqAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultIncludePajwohLqAccessGroupRule) implementsZoneAccessGroupListResponseResultInclude() {
}

type ZoneAccessGroupListResponseResultIncludePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                   `json:"id,required"`
	JSON zoneAccessGroupListResponseResultIncludePajwohLqAccessGroupRuleGroupJSON `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqAccessGroupRuleGroup]
type zoneAccessGroupListResponseResultIncludePajwohLqAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZoneAccessGroupListResponseResultIncludePajwohLqAzureGroupRule struct {
	AzureAd ZoneAccessGroupListResponseResultIncludePajwohLqAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zoneAccessGroupListResponseResultIncludePajwohLqAzureGroupRuleJSON    `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqAzureGroupRule]
type zoneAccessGroupListResponseResultIncludePajwohLqAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultIncludePajwohLqAzureGroupRule) implementsZoneAccessGroupListResponseResultInclude() {
}

type ZoneAccessGroupListResponseResultIncludePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                    `json:"connection_id,required"`
	JSON         zoneAccessGroupListResponseResultIncludePajwohLqAzureGroupRuleAzureAdJSON `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqAzureGroupRuleAzureAd]
type zoneAccessGroupListResponseResultIncludePajwohLqAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZoneAccessGroupListResponseResultIncludePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization ZoneAccessGroupListResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zoneAccessGroupListResponseResultIncludePajwohLqGitHubOrganizationRuleJSON               `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqGitHubOrganizationRule]
type zoneAccessGroupListResponseResultIncludePajwohLqGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultIncludePajwohLqGitHubOrganizationRule) implementsZoneAccessGroupListResponseResultInclude() {
}

type ZoneAccessGroupListResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                       `json:"name,required"`
	JSON zoneAccessGroupListResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization]
type zoneAccessGroupListResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZoneAccessGroupListResponseResultIncludePajwohLqGsuiteGroupRule struct {
	Gsuite ZoneAccessGroupListResponseResultIncludePajwohLqGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zoneAccessGroupListResponseResultIncludePajwohLqGsuiteGroupRuleJSON   `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqGsuiteGroupRule]
type zoneAccessGroupListResponseResultIncludePajwohLqGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultIncludePajwohLqGsuiteGroupRule) implementsZoneAccessGroupListResponseResultInclude() {
}

type ZoneAccessGroupListResponseResultIncludePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                    `json:"email,required"`
	JSON  zoneAccessGroupListResponseResultIncludePajwohLqGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqGsuiteGroupRuleGsuite]
type zoneAccessGroupListResponseResultIncludePajwohLqGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZoneAccessGroupListResponseResultIncludePajwohLqOktaGroupRule struct {
	Okta ZoneAccessGroupListResponseResultIncludePajwohLqOktaGroupRuleOkta `json:"okta,required"`
	JSON zoneAccessGroupListResponseResultIncludePajwohLqOktaGroupRuleJSON `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqOktaGroupRule]
type zoneAccessGroupListResponseResultIncludePajwohLqOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultIncludePajwohLqOktaGroupRule) implementsZoneAccessGroupListResponseResultInclude() {
}

type ZoneAccessGroupListResponseResultIncludePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                `json:"email,required"`
	JSON  zoneAccessGroupListResponseResultIncludePajwohLqOktaGroupRuleOktaJSON `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqOktaGroupRuleOkta]
type zoneAccessGroupListResponseResultIncludePajwohLqOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZoneAccessGroupListResponseResultIncludePajwohLqSamlGroupRule struct {
	Saml ZoneAccessGroupListResponseResultIncludePajwohLqSamlGroupRuleSaml `json:"saml,required"`
	JSON zoneAccessGroupListResponseResultIncludePajwohLqSamlGroupRuleJSON `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqSamlGroupRule]
type zoneAccessGroupListResponseResultIncludePajwohLqSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultIncludePajwohLqSamlGroupRule) implementsZoneAccessGroupListResponseResultInclude() {
}

type ZoneAccessGroupListResponseResultIncludePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                `json:"attribute_value,required"`
	JSON           zoneAccessGroupListResponseResultIncludePajwohLqSamlGroupRuleSamlJSON `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqSamlGroupRuleSaml]
type zoneAccessGroupListResponseResultIncludePajwohLqSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type ZoneAccessGroupListResponseResultIncludePajwohLqServiceTokenRule struct {
	ServiceToken ZoneAccessGroupListResponseResultIncludePajwohLqServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zoneAccessGroupListResponseResultIncludePajwohLqServiceTokenRuleJSON         `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqServiceTokenRuleJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqServiceTokenRule]
type zoneAccessGroupListResponseResultIncludePajwohLqServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultIncludePajwohLqServiceTokenRule) implementsZoneAccessGroupListResponseResultInclude() {
}

type ZoneAccessGroupListResponseResultIncludePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                           `json:"token_id,required"`
	JSON    zoneAccessGroupListResponseResultIncludePajwohLqServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqServiceTokenRuleServiceToken]
type zoneAccessGroupListResponseResultIncludePajwohLqServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type ZoneAccessGroupListResponseResultIncludePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                  `json:"any_valid_service_token,required"`
	JSON                 zoneAccessGroupListResponseResultIncludePajwohLqAnyValidServiceTokenRuleJSON `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqAnyValidServiceTokenRule]
type zoneAccessGroupListResponseResultIncludePajwohLqAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultIncludePajwohLqAnyValidServiceTokenRule) implementsZoneAccessGroupListResponseResultInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZoneAccessGroupListResponseResultIncludePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation ZoneAccessGroupListResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zoneAccessGroupListResponseResultIncludePajwohLqExternalEvaluationRuleJSON               `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqExternalEvaluationRule]
type zoneAccessGroupListResponseResultIncludePajwohLqExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultIncludePajwohLqExternalEvaluationRule) implementsZoneAccessGroupListResponseResultInclude() {
}

type ZoneAccessGroupListResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                       `json:"keys_url,required"`
	JSON    zoneAccessGroupListResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation]
type zoneAccessGroupListResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type ZoneAccessGroupListResponseResultIncludePajwohLqCountryRule struct {
	Geo  ZoneAccessGroupListResponseResultIncludePajwohLqCountryRuleGeo  `json:"geo,required"`
	JSON zoneAccessGroupListResponseResultIncludePajwohLqCountryRuleJSON `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqCountryRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqCountryRule]
type zoneAccessGroupListResponseResultIncludePajwohLqCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultIncludePajwohLqCountryRule) implementsZoneAccessGroupListResponseResultInclude() {
}

type ZoneAccessGroupListResponseResultIncludePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                             `json:"country_code,required"`
	JSON        zoneAccessGroupListResponseResultIncludePajwohLqCountryRuleGeoJSON `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqCountryRuleGeo]
type zoneAccessGroupListResponseResultIncludePajwohLqCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type ZoneAccessGroupListResponseResultIncludePajwohLqAuthenticationMethodRule struct {
	AuthMethod ZoneAccessGroupListResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zoneAccessGroupListResponseResultIncludePajwohLqAuthenticationMethodRuleJSON       `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqAuthenticationMethodRule]
type zoneAccessGroupListResponseResultIncludePajwohLqAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultIncludePajwohLqAuthenticationMethodRule) implementsZoneAccessGroupListResponseResultInclude() {
}

type ZoneAccessGroupListResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                 `json:"auth_method,required"`
	JSON       zoneAccessGroupListResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod]
type zoneAccessGroupListResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type ZoneAccessGroupListResponseResultIncludePajwohLqDevicePostureRule struct {
	DevicePosture ZoneAccessGroupListResponseResultIncludePajwohLqDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zoneAccessGroupListResponseResultIncludePajwohLqDevicePostureRuleJSON          `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqDevicePostureRule]
type zoneAccessGroupListResponseResultIncludePajwohLqDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultIncludePajwohLqDevicePostureRule) implementsZoneAccessGroupListResponseResultInclude() {
}

type ZoneAccessGroupListResponseResultIncludePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                             `json:"integration_uid,required"`
	JSON           zoneAccessGroupListResponseResultIncludePajwohLqDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zoneAccessGroupListResponseResultIncludePajwohLqDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultIncludePajwohLqDevicePostureRuleDevicePosture]
type zoneAccessGroupListResponseResultIncludePajwohLqDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultIncludePajwohLqDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [ZoneAccessGroupListResponseResultRequirePajwohLqEmailRule],
// [ZoneAccessGroupListResponseResultRequirePajwohLqEmailListRule],
// [ZoneAccessGroupListResponseResultRequirePajwohLqDomainRule],
// [ZoneAccessGroupListResponseResultRequirePajwohLqEveryoneRule],
// [ZoneAccessGroupListResponseResultRequirePajwohLqIPRule],
// [ZoneAccessGroupListResponseResultRequirePajwohLqIPListRule],
// [ZoneAccessGroupListResponseResultRequirePajwohLqCertificateRule],
// [ZoneAccessGroupListResponseResultRequirePajwohLqAccessGroupRule],
// [ZoneAccessGroupListResponseResultRequirePajwohLqAzureGroupRule],
// [ZoneAccessGroupListResponseResultRequirePajwohLqGitHubOrganizationRule],
// [ZoneAccessGroupListResponseResultRequirePajwohLqGsuiteGroupRule],
// [ZoneAccessGroupListResponseResultRequirePajwohLqOktaGroupRule],
// [ZoneAccessGroupListResponseResultRequirePajwohLqSamlGroupRule],
// [ZoneAccessGroupListResponseResultRequirePajwohLqServiceTokenRule],
// [ZoneAccessGroupListResponseResultRequirePajwohLqAnyValidServiceTokenRule],
// [ZoneAccessGroupListResponseResultRequirePajwohLqExternalEvaluationRule],
// [ZoneAccessGroupListResponseResultRequirePajwohLqCountryRule],
// [ZoneAccessGroupListResponseResultRequirePajwohLqAuthenticationMethodRule] or
// [ZoneAccessGroupListResponseResultRequirePajwohLqDevicePostureRule].
type ZoneAccessGroupListResponseResultRequire interface {
	implementsZoneAccessGroupListResponseResultRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneAccessGroupListResponseResultRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type ZoneAccessGroupListResponseResultRequirePajwohLqEmailRule struct {
	Email ZoneAccessGroupListResponseResultRequirePajwohLqEmailRuleEmail `json:"email,required"`
	JSON  zoneAccessGroupListResponseResultRequirePajwohLqEmailRuleJSON  `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqEmailRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqEmailRule]
type zoneAccessGroupListResponseResultRequirePajwohLqEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultRequirePajwohLqEmailRule) implementsZoneAccessGroupListResponseResultRequire() {
}

type ZoneAccessGroupListResponseResultRequirePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email string                                                             `json:"email,required" format:"email"`
	JSON  zoneAccessGroupListResponseResultRequirePajwohLqEmailRuleEmailJSON `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqEmailRuleEmail]
type zoneAccessGroupListResponseResultRequirePajwohLqEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type ZoneAccessGroupListResponseResultRequirePajwohLqEmailListRule struct {
	EmailList ZoneAccessGroupListResponseResultRequirePajwohLqEmailListRuleEmailList `json:"email_list,required"`
	JSON      zoneAccessGroupListResponseResultRequirePajwohLqEmailListRuleJSON      `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqEmailListRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqEmailListRule]
type zoneAccessGroupListResponseResultRequirePajwohLqEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultRequirePajwohLqEmailListRule) implementsZoneAccessGroupListResponseResultRequire() {
}

type ZoneAccessGroupListResponseResultRequirePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                     `json:"id,required"`
	JSON zoneAccessGroupListResponseResultRequirePajwohLqEmailListRuleEmailListJSON `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqEmailListRuleEmailList]
type zoneAccessGroupListResponseResultRequirePajwohLqEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type ZoneAccessGroupListResponseResultRequirePajwohLqDomainRule struct {
	EmailDomain ZoneAccessGroupListResponseResultRequirePajwohLqDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zoneAccessGroupListResponseResultRequirePajwohLqDomainRuleJSON        `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqDomainRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqDomainRule]
type zoneAccessGroupListResponseResultRequirePajwohLqDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultRequirePajwohLqDomainRule) implementsZoneAccessGroupListResponseResultRequire() {
}

type ZoneAccessGroupListResponseResultRequirePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                    `json:"domain,required"`
	JSON   zoneAccessGroupListResponseResultRequirePajwohLqDomainRuleEmailDomainJSON `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqDomainRuleEmailDomain]
type zoneAccessGroupListResponseResultRequirePajwohLqDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type ZoneAccessGroupListResponseResultRequirePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                      `json:"everyone,required"`
	JSON     zoneAccessGroupListResponseResultRequirePajwohLqEveryoneRuleJSON `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqEveryoneRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqEveryoneRule]
type zoneAccessGroupListResponseResultRequirePajwohLqEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultRequirePajwohLqEveryoneRule) implementsZoneAccessGroupListResponseResultRequire() {
}

// Matches an IP address block.
type ZoneAccessGroupListResponseResultRequirePajwohLqIPRule struct {
	IP   ZoneAccessGroupListResponseResultRequirePajwohLqIPRuleIP   `json:"ip,required"`
	JSON zoneAccessGroupListResponseResultRequirePajwohLqIPRuleJSON `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqIPRuleJSON contains the JSON
// metadata for the struct [ZoneAccessGroupListResponseResultRequirePajwohLqIPRule]
type zoneAccessGroupListResponseResultRequirePajwohLqIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultRequirePajwohLqIPRule) implementsZoneAccessGroupListResponseResultRequire() {
}

type ZoneAccessGroupListResponseResultRequirePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                       `json:"ip,required"`
	JSON zoneAccessGroupListResponseResultRequirePajwohLqIPRuleIPJSON `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqIPRuleIPJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqIPRuleIP]
type zoneAccessGroupListResponseResultRequirePajwohLqIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type ZoneAccessGroupListResponseResultRequirePajwohLqIPListRule struct {
	IPList ZoneAccessGroupListResponseResultRequirePajwohLqIPListRuleIPList `json:"ip_list,required"`
	JSON   zoneAccessGroupListResponseResultRequirePajwohLqIPListRuleJSON   `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqIPListRuleJSON contains the JSON
// metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqIPListRule]
type zoneAccessGroupListResponseResultRequirePajwohLqIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultRequirePajwohLqIPListRule) implementsZoneAccessGroupListResponseResultRequire() {
}

type ZoneAccessGroupListResponseResultRequirePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                               `json:"id,required"`
	JSON zoneAccessGroupListResponseResultRequirePajwohLqIPListRuleIPListJSON `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqIPListRuleIPListJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqIPListRuleIPList]
type zoneAccessGroupListResponseResultRequirePajwohLqIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type ZoneAccessGroupListResponseResultRequirePajwohLqCertificateRule struct {
	Certificate interface{}                                                         `json:"certificate,required"`
	JSON        zoneAccessGroupListResponseResultRequirePajwohLqCertificateRuleJSON `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqCertificateRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqCertificateRule]
type zoneAccessGroupListResponseResultRequirePajwohLqCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultRequirePajwohLqCertificateRule) implementsZoneAccessGroupListResponseResultRequire() {
}

// Matches an Access group.
type ZoneAccessGroupListResponseResultRequirePajwohLqAccessGroupRule struct {
	Group ZoneAccessGroupListResponseResultRequirePajwohLqAccessGroupRuleGroup `json:"group,required"`
	JSON  zoneAccessGroupListResponseResultRequirePajwohLqAccessGroupRuleJSON  `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqAccessGroupRule]
type zoneAccessGroupListResponseResultRequirePajwohLqAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultRequirePajwohLqAccessGroupRule) implementsZoneAccessGroupListResponseResultRequire() {
}

type ZoneAccessGroupListResponseResultRequirePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                   `json:"id,required"`
	JSON zoneAccessGroupListResponseResultRequirePajwohLqAccessGroupRuleGroupJSON `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqAccessGroupRuleGroup]
type zoneAccessGroupListResponseResultRequirePajwohLqAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZoneAccessGroupListResponseResultRequirePajwohLqAzureGroupRule struct {
	AzureAd ZoneAccessGroupListResponseResultRequirePajwohLqAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zoneAccessGroupListResponseResultRequirePajwohLqAzureGroupRuleJSON    `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqAzureGroupRule]
type zoneAccessGroupListResponseResultRequirePajwohLqAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultRequirePajwohLqAzureGroupRule) implementsZoneAccessGroupListResponseResultRequire() {
}

type ZoneAccessGroupListResponseResultRequirePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                    `json:"connection_id,required"`
	JSON         zoneAccessGroupListResponseResultRequirePajwohLqAzureGroupRuleAzureAdJSON `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqAzureGroupRuleAzureAd]
type zoneAccessGroupListResponseResultRequirePajwohLqAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZoneAccessGroupListResponseResultRequirePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization ZoneAccessGroupListResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zoneAccessGroupListResponseResultRequirePajwohLqGitHubOrganizationRuleJSON               `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqGitHubOrganizationRule]
type zoneAccessGroupListResponseResultRequirePajwohLqGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultRequirePajwohLqGitHubOrganizationRule) implementsZoneAccessGroupListResponseResultRequire() {
}

type ZoneAccessGroupListResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                       `json:"name,required"`
	JSON zoneAccessGroupListResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization]
type zoneAccessGroupListResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZoneAccessGroupListResponseResultRequirePajwohLqGsuiteGroupRule struct {
	Gsuite ZoneAccessGroupListResponseResultRequirePajwohLqGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zoneAccessGroupListResponseResultRequirePajwohLqGsuiteGroupRuleJSON   `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqGsuiteGroupRule]
type zoneAccessGroupListResponseResultRequirePajwohLqGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultRequirePajwohLqGsuiteGroupRule) implementsZoneAccessGroupListResponseResultRequire() {
}

type ZoneAccessGroupListResponseResultRequirePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                    `json:"email,required"`
	JSON  zoneAccessGroupListResponseResultRequirePajwohLqGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqGsuiteGroupRuleGsuite]
type zoneAccessGroupListResponseResultRequirePajwohLqGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZoneAccessGroupListResponseResultRequirePajwohLqOktaGroupRule struct {
	Okta ZoneAccessGroupListResponseResultRequirePajwohLqOktaGroupRuleOkta `json:"okta,required"`
	JSON zoneAccessGroupListResponseResultRequirePajwohLqOktaGroupRuleJSON `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqOktaGroupRule]
type zoneAccessGroupListResponseResultRequirePajwohLqOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultRequirePajwohLqOktaGroupRule) implementsZoneAccessGroupListResponseResultRequire() {
}

type ZoneAccessGroupListResponseResultRequirePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                `json:"email,required"`
	JSON  zoneAccessGroupListResponseResultRequirePajwohLqOktaGroupRuleOktaJSON `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqOktaGroupRuleOkta]
type zoneAccessGroupListResponseResultRequirePajwohLqOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZoneAccessGroupListResponseResultRequirePajwohLqSamlGroupRule struct {
	Saml ZoneAccessGroupListResponseResultRequirePajwohLqSamlGroupRuleSaml `json:"saml,required"`
	JSON zoneAccessGroupListResponseResultRequirePajwohLqSamlGroupRuleJSON `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqSamlGroupRule]
type zoneAccessGroupListResponseResultRequirePajwohLqSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultRequirePajwohLqSamlGroupRule) implementsZoneAccessGroupListResponseResultRequire() {
}

type ZoneAccessGroupListResponseResultRequirePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                `json:"attribute_value,required"`
	JSON           zoneAccessGroupListResponseResultRequirePajwohLqSamlGroupRuleSamlJSON `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqSamlGroupRuleSaml]
type zoneAccessGroupListResponseResultRequirePajwohLqSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type ZoneAccessGroupListResponseResultRequirePajwohLqServiceTokenRule struct {
	ServiceToken ZoneAccessGroupListResponseResultRequirePajwohLqServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zoneAccessGroupListResponseResultRequirePajwohLqServiceTokenRuleJSON         `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqServiceTokenRuleJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqServiceTokenRule]
type zoneAccessGroupListResponseResultRequirePajwohLqServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultRequirePajwohLqServiceTokenRule) implementsZoneAccessGroupListResponseResultRequire() {
}

type ZoneAccessGroupListResponseResultRequirePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                           `json:"token_id,required"`
	JSON    zoneAccessGroupListResponseResultRequirePajwohLqServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqServiceTokenRuleServiceToken]
type zoneAccessGroupListResponseResultRequirePajwohLqServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type ZoneAccessGroupListResponseResultRequirePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                  `json:"any_valid_service_token,required"`
	JSON                 zoneAccessGroupListResponseResultRequirePajwohLqAnyValidServiceTokenRuleJSON `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqAnyValidServiceTokenRule]
type zoneAccessGroupListResponseResultRequirePajwohLqAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultRequirePajwohLqAnyValidServiceTokenRule) implementsZoneAccessGroupListResponseResultRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZoneAccessGroupListResponseResultRequirePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation ZoneAccessGroupListResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zoneAccessGroupListResponseResultRequirePajwohLqExternalEvaluationRuleJSON               `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqExternalEvaluationRule]
type zoneAccessGroupListResponseResultRequirePajwohLqExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultRequirePajwohLqExternalEvaluationRule) implementsZoneAccessGroupListResponseResultRequire() {
}

type ZoneAccessGroupListResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                       `json:"keys_url,required"`
	JSON    zoneAccessGroupListResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation]
type zoneAccessGroupListResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type ZoneAccessGroupListResponseResultRequirePajwohLqCountryRule struct {
	Geo  ZoneAccessGroupListResponseResultRequirePajwohLqCountryRuleGeo  `json:"geo,required"`
	JSON zoneAccessGroupListResponseResultRequirePajwohLqCountryRuleJSON `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqCountryRuleJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqCountryRule]
type zoneAccessGroupListResponseResultRequirePajwohLqCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultRequirePajwohLqCountryRule) implementsZoneAccessGroupListResponseResultRequire() {
}

type ZoneAccessGroupListResponseResultRequirePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                             `json:"country_code,required"`
	JSON        zoneAccessGroupListResponseResultRequirePajwohLqCountryRuleGeoJSON `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqCountryRuleGeo]
type zoneAccessGroupListResponseResultRequirePajwohLqCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type ZoneAccessGroupListResponseResultRequirePajwohLqAuthenticationMethodRule struct {
	AuthMethod ZoneAccessGroupListResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zoneAccessGroupListResponseResultRequirePajwohLqAuthenticationMethodRuleJSON       `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqAuthenticationMethodRule]
type zoneAccessGroupListResponseResultRequirePajwohLqAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultRequirePajwohLqAuthenticationMethodRule) implementsZoneAccessGroupListResponseResultRequire() {
}

type ZoneAccessGroupListResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                 `json:"auth_method,required"`
	JSON       zoneAccessGroupListResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod]
type zoneAccessGroupListResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type ZoneAccessGroupListResponseResultRequirePajwohLqDevicePostureRule struct {
	DevicePosture ZoneAccessGroupListResponseResultRequirePajwohLqDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zoneAccessGroupListResponseResultRequirePajwohLqDevicePostureRuleJSON          `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqDevicePostureRule]
type zoneAccessGroupListResponseResultRequirePajwohLqDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessGroupListResponseResultRequirePajwohLqDevicePostureRule) implementsZoneAccessGroupListResponseResultRequire() {
}

type ZoneAccessGroupListResponseResultRequirePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                             `json:"integration_uid,required"`
	JSON           zoneAccessGroupListResponseResultRequirePajwohLqDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zoneAccessGroupListResponseResultRequirePajwohLqDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZoneAccessGroupListResponseResultRequirePajwohLqDevicePostureRuleDevicePosture]
type zoneAccessGroupListResponseResultRequirePajwohLqDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultRequirePajwohLqDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessGroupListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                   `json:"total_count"`
	JSON       zoneAccessGroupListResponseResultInfoJSON `json:"-"`
}

// zoneAccessGroupListResponseResultInfoJSON contains the JSON metadata for the
// struct [ZoneAccessGroupListResponseResultInfo]
type zoneAccessGroupListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessGroupListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneAccessGroupListResponseSuccess bool

const (
	ZoneAccessGroupListResponseSuccessTrue ZoneAccessGroupListResponseSuccess = true
)

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
	// UUID
	ID   string                                  `json:"id"`
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
// Satisfied by [ZoneAccessGroupNewParamsIncludePajwohLqEmailRule],
// [ZoneAccessGroupNewParamsIncludePajwohLqEmailListRule],
// [ZoneAccessGroupNewParamsIncludePajwohLqDomainRule],
// [ZoneAccessGroupNewParamsIncludePajwohLqEveryoneRule],
// [ZoneAccessGroupNewParamsIncludePajwohLqIPRule],
// [ZoneAccessGroupNewParamsIncludePajwohLqIPListRule],
// [ZoneAccessGroupNewParamsIncludePajwohLqCertificateRule],
// [ZoneAccessGroupNewParamsIncludePajwohLqAccessGroupRule],
// [ZoneAccessGroupNewParamsIncludePajwohLqAzureGroupRule],
// [ZoneAccessGroupNewParamsIncludePajwohLqGitHubOrganizationRule],
// [ZoneAccessGroupNewParamsIncludePajwohLqGsuiteGroupRule],
// [ZoneAccessGroupNewParamsIncludePajwohLqOktaGroupRule],
// [ZoneAccessGroupNewParamsIncludePajwohLqSamlGroupRule],
// [ZoneAccessGroupNewParamsIncludePajwohLqServiceTokenRule],
// [ZoneAccessGroupNewParamsIncludePajwohLqAnyValidServiceTokenRule],
// [ZoneAccessGroupNewParamsIncludePajwohLqExternalEvaluationRule],
// [ZoneAccessGroupNewParamsIncludePajwohLqCountryRule],
// [ZoneAccessGroupNewParamsIncludePajwohLqAuthenticationMethodRule],
// [ZoneAccessGroupNewParamsIncludePajwohLqDevicePostureRule].
type ZoneAccessGroupNewParamsInclude interface {
	implementsZoneAccessGroupNewParamsInclude()
}

// Matches a specific email.
type ZoneAccessGroupNewParamsIncludePajwohLqEmailRule struct {
	Email param.Field[ZoneAccessGroupNewParamsIncludePajwohLqEmailRuleEmail] `json:"email,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqEmailRule) implementsZoneAccessGroupNewParamsInclude() {
}

type ZoneAccessGroupNewParamsIncludePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type ZoneAccessGroupNewParamsIncludePajwohLqEmailListRule struct {
	EmailList param.Field[ZoneAccessGroupNewParamsIncludePajwohLqEmailListRuleEmailList] `json:"email_list,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqEmailListRule) implementsZoneAccessGroupNewParamsInclude() {
}

type ZoneAccessGroupNewParamsIncludePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type ZoneAccessGroupNewParamsIncludePajwohLqDomainRule struct {
	EmailDomain param.Field[ZoneAccessGroupNewParamsIncludePajwohLqDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqDomainRule) implementsZoneAccessGroupNewParamsInclude() {
}

type ZoneAccessGroupNewParamsIncludePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type ZoneAccessGroupNewParamsIncludePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqEveryoneRule) implementsZoneAccessGroupNewParamsInclude() {
}

// Matches an IP address block.
type ZoneAccessGroupNewParamsIncludePajwohLqIPRule struct {
	IP param.Field[ZoneAccessGroupNewParamsIncludePajwohLqIPRuleIP] `json:"ip,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqIPRule) implementsZoneAccessGroupNewParamsInclude() {}

type ZoneAccessGroupNewParamsIncludePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type ZoneAccessGroupNewParamsIncludePajwohLqIPListRule struct {
	IPList param.Field[ZoneAccessGroupNewParamsIncludePajwohLqIPListRuleIPList] `json:"ip_list,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqIPListRule) implementsZoneAccessGroupNewParamsInclude() {
}

type ZoneAccessGroupNewParamsIncludePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type ZoneAccessGroupNewParamsIncludePajwohLqCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqCertificateRule) implementsZoneAccessGroupNewParamsInclude() {
}

// Matches an Access group.
type ZoneAccessGroupNewParamsIncludePajwohLqAccessGroupRule struct {
	Group param.Field[ZoneAccessGroupNewParamsIncludePajwohLqAccessGroupRuleGroup] `json:"group,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqAccessGroupRule) implementsZoneAccessGroupNewParamsInclude() {
}

type ZoneAccessGroupNewParamsIncludePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZoneAccessGroupNewParamsIncludePajwohLqAzureGroupRule struct {
	AzureAd param.Field[ZoneAccessGroupNewParamsIncludePajwohLqAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqAzureGroupRule) implementsZoneAccessGroupNewParamsInclude() {
}

type ZoneAccessGroupNewParamsIncludePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZoneAccessGroupNewParamsIncludePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization param.Field[ZoneAccessGroupNewParamsIncludePajwohLqGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqGitHubOrganizationRule) implementsZoneAccessGroupNewParamsInclude() {
}

type ZoneAccessGroupNewParamsIncludePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZoneAccessGroupNewParamsIncludePajwohLqGsuiteGroupRule struct {
	Gsuite param.Field[ZoneAccessGroupNewParamsIncludePajwohLqGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqGsuiteGroupRule) implementsZoneAccessGroupNewParamsInclude() {
}

type ZoneAccessGroupNewParamsIncludePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZoneAccessGroupNewParamsIncludePajwohLqOktaGroupRule struct {
	Okta param.Field[ZoneAccessGroupNewParamsIncludePajwohLqOktaGroupRuleOkta] `json:"okta,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqOktaGroupRule) implementsZoneAccessGroupNewParamsInclude() {
}

type ZoneAccessGroupNewParamsIncludePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZoneAccessGroupNewParamsIncludePajwohLqSamlGroupRule struct {
	Saml param.Field[ZoneAccessGroupNewParamsIncludePajwohLqSamlGroupRuleSaml] `json:"saml,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqSamlGroupRule) implementsZoneAccessGroupNewParamsInclude() {
}

type ZoneAccessGroupNewParamsIncludePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type ZoneAccessGroupNewParamsIncludePajwohLqServiceTokenRule struct {
	ServiceToken param.Field[ZoneAccessGroupNewParamsIncludePajwohLqServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqServiceTokenRule) implementsZoneAccessGroupNewParamsInclude() {
}

type ZoneAccessGroupNewParamsIncludePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type ZoneAccessGroupNewParamsIncludePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqAnyValidServiceTokenRule) implementsZoneAccessGroupNewParamsInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZoneAccessGroupNewParamsIncludePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation param.Field[ZoneAccessGroupNewParamsIncludePajwohLqExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqExternalEvaluationRule) implementsZoneAccessGroupNewParamsInclude() {
}

type ZoneAccessGroupNewParamsIncludePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type ZoneAccessGroupNewParamsIncludePajwohLqCountryRule struct {
	Geo param.Field[ZoneAccessGroupNewParamsIncludePajwohLqCountryRuleGeo] `json:"geo,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqCountryRule) implementsZoneAccessGroupNewParamsInclude() {
}

type ZoneAccessGroupNewParamsIncludePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type ZoneAccessGroupNewParamsIncludePajwohLqAuthenticationMethodRule struct {
	AuthMethod param.Field[ZoneAccessGroupNewParamsIncludePajwohLqAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqAuthenticationMethodRule) implementsZoneAccessGroupNewParamsInclude() {
}

type ZoneAccessGroupNewParamsIncludePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type ZoneAccessGroupNewParamsIncludePajwohLqDevicePostureRule struct {
	DevicePosture param.Field[ZoneAccessGroupNewParamsIncludePajwohLqDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqDevicePostureRule) implementsZoneAccessGroupNewParamsInclude() {
}

type ZoneAccessGroupNewParamsIncludePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r ZoneAccessGroupNewParamsIncludePajwohLqDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [ZoneAccessGroupNewParamsExcludePajwohLqEmailRule],
// [ZoneAccessGroupNewParamsExcludePajwohLqEmailListRule],
// [ZoneAccessGroupNewParamsExcludePajwohLqDomainRule],
// [ZoneAccessGroupNewParamsExcludePajwohLqEveryoneRule],
// [ZoneAccessGroupNewParamsExcludePajwohLqIPRule],
// [ZoneAccessGroupNewParamsExcludePajwohLqIPListRule],
// [ZoneAccessGroupNewParamsExcludePajwohLqCertificateRule],
// [ZoneAccessGroupNewParamsExcludePajwohLqAccessGroupRule],
// [ZoneAccessGroupNewParamsExcludePajwohLqAzureGroupRule],
// [ZoneAccessGroupNewParamsExcludePajwohLqGitHubOrganizationRule],
// [ZoneAccessGroupNewParamsExcludePajwohLqGsuiteGroupRule],
// [ZoneAccessGroupNewParamsExcludePajwohLqOktaGroupRule],
// [ZoneAccessGroupNewParamsExcludePajwohLqSamlGroupRule],
// [ZoneAccessGroupNewParamsExcludePajwohLqServiceTokenRule],
// [ZoneAccessGroupNewParamsExcludePajwohLqAnyValidServiceTokenRule],
// [ZoneAccessGroupNewParamsExcludePajwohLqExternalEvaluationRule],
// [ZoneAccessGroupNewParamsExcludePajwohLqCountryRule],
// [ZoneAccessGroupNewParamsExcludePajwohLqAuthenticationMethodRule],
// [ZoneAccessGroupNewParamsExcludePajwohLqDevicePostureRule].
type ZoneAccessGroupNewParamsExclude interface {
	implementsZoneAccessGroupNewParamsExclude()
}

// Matches a specific email.
type ZoneAccessGroupNewParamsExcludePajwohLqEmailRule struct {
	Email param.Field[ZoneAccessGroupNewParamsExcludePajwohLqEmailRuleEmail] `json:"email,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqEmailRule) implementsZoneAccessGroupNewParamsExclude() {
}

type ZoneAccessGroupNewParamsExcludePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type ZoneAccessGroupNewParamsExcludePajwohLqEmailListRule struct {
	EmailList param.Field[ZoneAccessGroupNewParamsExcludePajwohLqEmailListRuleEmailList] `json:"email_list,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqEmailListRule) implementsZoneAccessGroupNewParamsExclude() {
}

type ZoneAccessGroupNewParamsExcludePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type ZoneAccessGroupNewParamsExcludePajwohLqDomainRule struct {
	EmailDomain param.Field[ZoneAccessGroupNewParamsExcludePajwohLqDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqDomainRule) implementsZoneAccessGroupNewParamsExclude() {
}

type ZoneAccessGroupNewParamsExcludePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type ZoneAccessGroupNewParamsExcludePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqEveryoneRule) implementsZoneAccessGroupNewParamsExclude() {
}

// Matches an IP address block.
type ZoneAccessGroupNewParamsExcludePajwohLqIPRule struct {
	IP param.Field[ZoneAccessGroupNewParamsExcludePajwohLqIPRuleIP] `json:"ip,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqIPRule) implementsZoneAccessGroupNewParamsExclude() {}

type ZoneAccessGroupNewParamsExcludePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type ZoneAccessGroupNewParamsExcludePajwohLqIPListRule struct {
	IPList param.Field[ZoneAccessGroupNewParamsExcludePajwohLqIPListRuleIPList] `json:"ip_list,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqIPListRule) implementsZoneAccessGroupNewParamsExclude() {
}

type ZoneAccessGroupNewParamsExcludePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type ZoneAccessGroupNewParamsExcludePajwohLqCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqCertificateRule) implementsZoneAccessGroupNewParamsExclude() {
}

// Matches an Access group.
type ZoneAccessGroupNewParamsExcludePajwohLqAccessGroupRule struct {
	Group param.Field[ZoneAccessGroupNewParamsExcludePajwohLqAccessGroupRuleGroup] `json:"group,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqAccessGroupRule) implementsZoneAccessGroupNewParamsExclude() {
}

type ZoneAccessGroupNewParamsExcludePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZoneAccessGroupNewParamsExcludePajwohLqAzureGroupRule struct {
	AzureAd param.Field[ZoneAccessGroupNewParamsExcludePajwohLqAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqAzureGroupRule) implementsZoneAccessGroupNewParamsExclude() {
}

type ZoneAccessGroupNewParamsExcludePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZoneAccessGroupNewParamsExcludePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization param.Field[ZoneAccessGroupNewParamsExcludePajwohLqGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqGitHubOrganizationRule) implementsZoneAccessGroupNewParamsExclude() {
}

type ZoneAccessGroupNewParamsExcludePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZoneAccessGroupNewParamsExcludePajwohLqGsuiteGroupRule struct {
	Gsuite param.Field[ZoneAccessGroupNewParamsExcludePajwohLqGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqGsuiteGroupRule) implementsZoneAccessGroupNewParamsExclude() {
}

type ZoneAccessGroupNewParamsExcludePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZoneAccessGroupNewParamsExcludePajwohLqOktaGroupRule struct {
	Okta param.Field[ZoneAccessGroupNewParamsExcludePajwohLqOktaGroupRuleOkta] `json:"okta,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqOktaGroupRule) implementsZoneAccessGroupNewParamsExclude() {
}

type ZoneAccessGroupNewParamsExcludePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZoneAccessGroupNewParamsExcludePajwohLqSamlGroupRule struct {
	Saml param.Field[ZoneAccessGroupNewParamsExcludePajwohLqSamlGroupRuleSaml] `json:"saml,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqSamlGroupRule) implementsZoneAccessGroupNewParamsExclude() {
}

type ZoneAccessGroupNewParamsExcludePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type ZoneAccessGroupNewParamsExcludePajwohLqServiceTokenRule struct {
	ServiceToken param.Field[ZoneAccessGroupNewParamsExcludePajwohLqServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqServiceTokenRule) implementsZoneAccessGroupNewParamsExclude() {
}

type ZoneAccessGroupNewParamsExcludePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type ZoneAccessGroupNewParamsExcludePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqAnyValidServiceTokenRule) implementsZoneAccessGroupNewParamsExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZoneAccessGroupNewParamsExcludePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation param.Field[ZoneAccessGroupNewParamsExcludePajwohLqExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqExternalEvaluationRule) implementsZoneAccessGroupNewParamsExclude() {
}

type ZoneAccessGroupNewParamsExcludePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type ZoneAccessGroupNewParamsExcludePajwohLqCountryRule struct {
	Geo param.Field[ZoneAccessGroupNewParamsExcludePajwohLqCountryRuleGeo] `json:"geo,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqCountryRule) implementsZoneAccessGroupNewParamsExclude() {
}

type ZoneAccessGroupNewParamsExcludePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type ZoneAccessGroupNewParamsExcludePajwohLqAuthenticationMethodRule struct {
	AuthMethod param.Field[ZoneAccessGroupNewParamsExcludePajwohLqAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqAuthenticationMethodRule) implementsZoneAccessGroupNewParamsExclude() {
}

type ZoneAccessGroupNewParamsExcludePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type ZoneAccessGroupNewParamsExcludePajwohLqDevicePostureRule struct {
	DevicePosture param.Field[ZoneAccessGroupNewParamsExcludePajwohLqDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqDevicePostureRule) implementsZoneAccessGroupNewParamsExclude() {
}

type ZoneAccessGroupNewParamsExcludePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r ZoneAccessGroupNewParamsExcludePajwohLqDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [ZoneAccessGroupNewParamsRequirePajwohLqEmailRule],
// [ZoneAccessGroupNewParamsRequirePajwohLqEmailListRule],
// [ZoneAccessGroupNewParamsRequirePajwohLqDomainRule],
// [ZoneAccessGroupNewParamsRequirePajwohLqEveryoneRule],
// [ZoneAccessGroupNewParamsRequirePajwohLqIPRule],
// [ZoneAccessGroupNewParamsRequirePajwohLqIPListRule],
// [ZoneAccessGroupNewParamsRequirePajwohLqCertificateRule],
// [ZoneAccessGroupNewParamsRequirePajwohLqAccessGroupRule],
// [ZoneAccessGroupNewParamsRequirePajwohLqAzureGroupRule],
// [ZoneAccessGroupNewParamsRequirePajwohLqGitHubOrganizationRule],
// [ZoneAccessGroupNewParamsRequirePajwohLqGsuiteGroupRule],
// [ZoneAccessGroupNewParamsRequirePajwohLqOktaGroupRule],
// [ZoneAccessGroupNewParamsRequirePajwohLqSamlGroupRule],
// [ZoneAccessGroupNewParamsRequirePajwohLqServiceTokenRule],
// [ZoneAccessGroupNewParamsRequirePajwohLqAnyValidServiceTokenRule],
// [ZoneAccessGroupNewParamsRequirePajwohLqExternalEvaluationRule],
// [ZoneAccessGroupNewParamsRequirePajwohLqCountryRule],
// [ZoneAccessGroupNewParamsRequirePajwohLqAuthenticationMethodRule],
// [ZoneAccessGroupNewParamsRequirePajwohLqDevicePostureRule].
type ZoneAccessGroupNewParamsRequire interface {
	implementsZoneAccessGroupNewParamsRequire()
}

// Matches a specific email.
type ZoneAccessGroupNewParamsRequirePajwohLqEmailRule struct {
	Email param.Field[ZoneAccessGroupNewParamsRequirePajwohLqEmailRuleEmail] `json:"email,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqEmailRule) implementsZoneAccessGroupNewParamsRequire() {
}

type ZoneAccessGroupNewParamsRequirePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type ZoneAccessGroupNewParamsRequirePajwohLqEmailListRule struct {
	EmailList param.Field[ZoneAccessGroupNewParamsRequirePajwohLqEmailListRuleEmailList] `json:"email_list,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqEmailListRule) implementsZoneAccessGroupNewParamsRequire() {
}

type ZoneAccessGroupNewParamsRequirePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type ZoneAccessGroupNewParamsRequirePajwohLqDomainRule struct {
	EmailDomain param.Field[ZoneAccessGroupNewParamsRequirePajwohLqDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqDomainRule) implementsZoneAccessGroupNewParamsRequire() {
}

type ZoneAccessGroupNewParamsRequirePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type ZoneAccessGroupNewParamsRequirePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqEveryoneRule) implementsZoneAccessGroupNewParamsRequire() {
}

// Matches an IP address block.
type ZoneAccessGroupNewParamsRequirePajwohLqIPRule struct {
	IP param.Field[ZoneAccessGroupNewParamsRequirePajwohLqIPRuleIP] `json:"ip,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqIPRule) implementsZoneAccessGroupNewParamsRequire() {}

type ZoneAccessGroupNewParamsRequirePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type ZoneAccessGroupNewParamsRequirePajwohLqIPListRule struct {
	IPList param.Field[ZoneAccessGroupNewParamsRequirePajwohLqIPListRuleIPList] `json:"ip_list,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqIPListRule) implementsZoneAccessGroupNewParamsRequire() {
}

type ZoneAccessGroupNewParamsRequirePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type ZoneAccessGroupNewParamsRequirePajwohLqCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqCertificateRule) implementsZoneAccessGroupNewParamsRequire() {
}

// Matches an Access group.
type ZoneAccessGroupNewParamsRequirePajwohLqAccessGroupRule struct {
	Group param.Field[ZoneAccessGroupNewParamsRequirePajwohLqAccessGroupRuleGroup] `json:"group,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqAccessGroupRule) implementsZoneAccessGroupNewParamsRequire() {
}

type ZoneAccessGroupNewParamsRequirePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZoneAccessGroupNewParamsRequirePajwohLqAzureGroupRule struct {
	AzureAd param.Field[ZoneAccessGroupNewParamsRequirePajwohLqAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqAzureGroupRule) implementsZoneAccessGroupNewParamsRequire() {
}

type ZoneAccessGroupNewParamsRequirePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZoneAccessGroupNewParamsRequirePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization param.Field[ZoneAccessGroupNewParamsRequirePajwohLqGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqGitHubOrganizationRule) implementsZoneAccessGroupNewParamsRequire() {
}

type ZoneAccessGroupNewParamsRequirePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZoneAccessGroupNewParamsRequirePajwohLqGsuiteGroupRule struct {
	Gsuite param.Field[ZoneAccessGroupNewParamsRequirePajwohLqGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqGsuiteGroupRule) implementsZoneAccessGroupNewParamsRequire() {
}

type ZoneAccessGroupNewParamsRequirePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZoneAccessGroupNewParamsRequirePajwohLqOktaGroupRule struct {
	Okta param.Field[ZoneAccessGroupNewParamsRequirePajwohLqOktaGroupRuleOkta] `json:"okta,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqOktaGroupRule) implementsZoneAccessGroupNewParamsRequire() {
}

type ZoneAccessGroupNewParamsRequirePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZoneAccessGroupNewParamsRequirePajwohLqSamlGroupRule struct {
	Saml param.Field[ZoneAccessGroupNewParamsRequirePajwohLqSamlGroupRuleSaml] `json:"saml,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqSamlGroupRule) implementsZoneAccessGroupNewParamsRequire() {
}

type ZoneAccessGroupNewParamsRequirePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type ZoneAccessGroupNewParamsRequirePajwohLqServiceTokenRule struct {
	ServiceToken param.Field[ZoneAccessGroupNewParamsRequirePajwohLqServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqServiceTokenRule) implementsZoneAccessGroupNewParamsRequire() {
}

type ZoneAccessGroupNewParamsRequirePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type ZoneAccessGroupNewParamsRequirePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqAnyValidServiceTokenRule) implementsZoneAccessGroupNewParamsRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZoneAccessGroupNewParamsRequirePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation param.Field[ZoneAccessGroupNewParamsRequirePajwohLqExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqExternalEvaluationRule) implementsZoneAccessGroupNewParamsRequire() {
}

type ZoneAccessGroupNewParamsRequirePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type ZoneAccessGroupNewParamsRequirePajwohLqCountryRule struct {
	Geo param.Field[ZoneAccessGroupNewParamsRequirePajwohLqCountryRuleGeo] `json:"geo,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqCountryRule) implementsZoneAccessGroupNewParamsRequire() {
}

type ZoneAccessGroupNewParamsRequirePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type ZoneAccessGroupNewParamsRequirePajwohLqAuthenticationMethodRule struct {
	AuthMethod param.Field[ZoneAccessGroupNewParamsRequirePajwohLqAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqAuthenticationMethodRule) implementsZoneAccessGroupNewParamsRequire() {
}

type ZoneAccessGroupNewParamsRequirePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type ZoneAccessGroupNewParamsRequirePajwohLqDevicePostureRule struct {
	DevicePosture param.Field[ZoneAccessGroupNewParamsRequirePajwohLqDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqDevicePostureRule) implementsZoneAccessGroupNewParamsRequire() {
}

type ZoneAccessGroupNewParamsRequirePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r ZoneAccessGroupNewParamsRequirePajwohLqDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
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
// Satisfied by [ZoneAccessGroupUpdateParamsIncludePajwohLqEmailRule],
// [ZoneAccessGroupUpdateParamsIncludePajwohLqEmailListRule],
// [ZoneAccessGroupUpdateParamsIncludePajwohLqDomainRule],
// [ZoneAccessGroupUpdateParamsIncludePajwohLqEveryoneRule],
// [ZoneAccessGroupUpdateParamsIncludePajwohLqIPRule],
// [ZoneAccessGroupUpdateParamsIncludePajwohLqIPListRule],
// [ZoneAccessGroupUpdateParamsIncludePajwohLqCertificateRule],
// [ZoneAccessGroupUpdateParamsIncludePajwohLqAccessGroupRule],
// [ZoneAccessGroupUpdateParamsIncludePajwohLqAzureGroupRule],
// [ZoneAccessGroupUpdateParamsIncludePajwohLqGitHubOrganizationRule],
// [ZoneAccessGroupUpdateParamsIncludePajwohLqGsuiteGroupRule],
// [ZoneAccessGroupUpdateParamsIncludePajwohLqOktaGroupRule],
// [ZoneAccessGroupUpdateParamsIncludePajwohLqSamlGroupRule],
// [ZoneAccessGroupUpdateParamsIncludePajwohLqServiceTokenRule],
// [ZoneAccessGroupUpdateParamsIncludePajwohLqAnyValidServiceTokenRule],
// [ZoneAccessGroupUpdateParamsIncludePajwohLqExternalEvaluationRule],
// [ZoneAccessGroupUpdateParamsIncludePajwohLqCountryRule],
// [ZoneAccessGroupUpdateParamsIncludePajwohLqAuthenticationMethodRule],
// [ZoneAccessGroupUpdateParamsIncludePajwohLqDevicePostureRule].
type ZoneAccessGroupUpdateParamsInclude interface {
	implementsZoneAccessGroupUpdateParamsInclude()
}

// Matches a specific email.
type ZoneAccessGroupUpdateParamsIncludePajwohLqEmailRule struct {
	Email param.Field[ZoneAccessGroupUpdateParamsIncludePajwohLqEmailRuleEmail] `json:"email,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqEmailRule) implementsZoneAccessGroupUpdateParamsInclude() {
}

type ZoneAccessGroupUpdateParamsIncludePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type ZoneAccessGroupUpdateParamsIncludePajwohLqEmailListRule struct {
	EmailList param.Field[ZoneAccessGroupUpdateParamsIncludePajwohLqEmailListRuleEmailList] `json:"email_list,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqEmailListRule) implementsZoneAccessGroupUpdateParamsInclude() {
}

type ZoneAccessGroupUpdateParamsIncludePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type ZoneAccessGroupUpdateParamsIncludePajwohLqDomainRule struct {
	EmailDomain param.Field[ZoneAccessGroupUpdateParamsIncludePajwohLqDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqDomainRule) implementsZoneAccessGroupUpdateParamsInclude() {
}

type ZoneAccessGroupUpdateParamsIncludePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type ZoneAccessGroupUpdateParamsIncludePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqEveryoneRule) implementsZoneAccessGroupUpdateParamsInclude() {
}

// Matches an IP address block.
type ZoneAccessGroupUpdateParamsIncludePajwohLqIPRule struct {
	IP param.Field[ZoneAccessGroupUpdateParamsIncludePajwohLqIPRuleIP] `json:"ip,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqIPRule) implementsZoneAccessGroupUpdateParamsInclude() {
}

type ZoneAccessGroupUpdateParamsIncludePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type ZoneAccessGroupUpdateParamsIncludePajwohLqIPListRule struct {
	IPList param.Field[ZoneAccessGroupUpdateParamsIncludePajwohLqIPListRuleIPList] `json:"ip_list,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqIPListRule) implementsZoneAccessGroupUpdateParamsInclude() {
}

type ZoneAccessGroupUpdateParamsIncludePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type ZoneAccessGroupUpdateParamsIncludePajwohLqCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqCertificateRule) implementsZoneAccessGroupUpdateParamsInclude() {
}

// Matches an Access group.
type ZoneAccessGroupUpdateParamsIncludePajwohLqAccessGroupRule struct {
	Group param.Field[ZoneAccessGroupUpdateParamsIncludePajwohLqAccessGroupRuleGroup] `json:"group,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqAccessGroupRule) implementsZoneAccessGroupUpdateParamsInclude() {
}

type ZoneAccessGroupUpdateParamsIncludePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZoneAccessGroupUpdateParamsIncludePajwohLqAzureGroupRule struct {
	AzureAd param.Field[ZoneAccessGroupUpdateParamsIncludePajwohLqAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqAzureGroupRule) implementsZoneAccessGroupUpdateParamsInclude() {
}

type ZoneAccessGroupUpdateParamsIncludePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZoneAccessGroupUpdateParamsIncludePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization param.Field[ZoneAccessGroupUpdateParamsIncludePajwohLqGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqGitHubOrganizationRule) implementsZoneAccessGroupUpdateParamsInclude() {
}

type ZoneAccessGroupUpdateParamsIncludePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZoneAccessGroupUpdateParamsIncludePajwohLqGsuiteGroupRule struct {
	Gsuite param.Field[ZoneAccessGroupUpdateParamsIncludePajwohLqGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqGsuiteGroupRule) implementsZoneAccessGroupUpdateParamsInclude() {
}

type ZoneAccessGroupUpdateParamsIncludePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZoneAccessGroupUpdateParamsIncludePajwohLqOktaGroupRule struct {
	Okta param.Field[ZoneAccessGroupUpdateParamsIncludePajwohLqOktaGroupRuleOkta] `json:"okta,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqOktaGroupRule) implementsZoneAccessGroupUpdateParamsInclude() {
}

type ZoneAccessGroupUpdateParamsIncludePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZoneAccessGroupUpdateParamsIncludePajwohLqSamlGroupRule struct {
	Saml param.Field[ZoneAccessGroupUpdateParamsIncludePajwohLqSamlGroupRuleSaml] `json:"saml,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqSamlGroupRule) implementsZoneAccessGroupUpdateParamsInclude() {
}

type ZoneAccessGroupUpdateParamsIncludePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type ZoneAccessGroupUpdateParamsIncludePajwohLqServiceTokenRule struct {
	ServiceToken param.Field[ZoneAccessGroupUpdateParamsIncludePajwohLqServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqServiceTokenRule) implementsZoneAccessGroupUpdateParamsInclude() {
}

type ZoneAccessGroupUpdateParamsIncludePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type ZoneAccessGroupUpdateParamsIncludePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqAnyValidServiceTokenRule) implementsZoneAccessGroupUpdateParamsInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZoneAccessGroupUpdateParamsIncludePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation param.Field[ZoneAccessGroupUpdateParamsIncludePajwohLqExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqExternalEvaluationRule) implementsZoneAccessGroupUpdateParamsInclude() {
}

type ZoneAccessGroupUpdateParamsIncludePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type ZoneAccessGroupUpdateParamsIncludePajwohLqCountryRule struct {
	Geo param.Field[ZoneAccessGroupUpdateParamsIncludePajwohLqCountryRuleGeo] `json:"geo,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqCountryRule) implementsZoneAccessGroupUpdateParamsInclude() {
}

type ZoneAccessGroupUpdateParamsIncludePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type ZoneAccessGroupUpdateParamsIncludePajwohLqAuthenticationMethodRule struct {
	AuthMethod param.Field[ZoneAccessGroupUpdateParamsIncludePajwohLqAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqAuthenticationMethodRule) implementsZoneAccessGroupUpdateParamsInclude() {
}

type ZoneAccessGroupUpdateParamsIncludePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type ZoneAccessGroupUpdateParamsIncludePajwohLqDevicePostureRule struct {
	DevicePosture param.Field[ZoneAccessGroupUpdateParamsIncludePajwohLqDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqDevicePostureRule) implementsZoneAccessGroupUpdateParamsInclude() {
}

type ZoneAccessGroupUpdateParamsIncludePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r ZoneAccessGroupUpdateParamsIncludePajwohLqDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [ZoneAccessGroupUpdateParamsExcludePajwohLqEmailRule],
// [ZoneAccessGroupUpdateParamsExcludePajwohLqEmailListRule],
// [ZoneAccessGroupUpdateParamsExcludePajwohLqDomainRule],
// [ZoneAccessGroupUpdateParamsExcludePajwohLqEveryoneRule],
// [ZoneAccessGroupUpdateParamsExcludePajwohLqIPRule],
// [ZoneAccessGroupUpdateParamsExcludePajwohLqIPListRule],
// [ZoneAccessGroupUpdateParamsExcludePajwohLqCertificateRule],
// [ZoneAccessGroupUpdateParamsExcludePajwohLqAccessGroupRule],
// [ZoneAccessGroupUpdateParamsExcludePajwohLqAzureGroupRule],
// [ZoneAccessGroupUpdateParamsExcludePajwohLqGitHubOrganizationRule],
// [ZoneAccessGroupUpdateParamsExcludePajwohLqGsuiteGroupRule],
// [ZoneAccessGroupUpdateParamsExcludePajwohLqOktaGroupRule],
// [ZoneAccessGroupUpdateParamsExcludePajwohLqSamlGroupRule],
// [ZoneAccessGroupUpdateParamsExcludePajwohLqServiceTokenRule],
// [ZoneAccessGroupUpdateParamsExcludePajwohLqAnyValidServiceTokenRule],
// [ZoneAccessGroupUpdateParamsExcludePajwohLqExternalEvaluationRule],
// [ZoneAccessGroupUpdateParamsExcludePajwohLqCountryRule],
// [ZoneAccessGroupUpdateParamsExcludePajwohLqAuthenticationMethodRule],
// [ZoneAccessGroupUpdateParamsExcludePajwohLqDevicePostureRule].
type ZoneAccessGroupUpdateParamsExclude interface {
	implementsZoneAccessGroupUpdateParamsExclude()
}

// Matches a specific email.
type ZoneAccessGroupUpdateParamsExcludePajwohLqEmailRule struct {
	Email param.Field[ZoneAccessGroupUpdateParamsExcludePajwohLqEmailRuleEmail] `json:"email,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqEmailRule) implementsZoneAccessGroupUpdateParamsExclude() {
}

type ZoneAccessGroupUpdateParamsExcludePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type ZoneAccessGroupUpdateParamsExcludePajwohLqEmailListRule struct {
	EmailList param.Field[ZoneAccessGroupUpdateParamsExcludePajwohLqEmailListRuleEmailList] `json:"email_list,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqEmailListRule) implementsZoneAccessGroupUpdateParamsExclude() {
}

type ZoneAccessGroupUpdateParamsExcludePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type ZoneAccessGroupUpdateParamsExcludePajwohLqDomainRule struct {
	EmailDomain param.Field[ZoneAccessGroupUpdateParamsExcludePajwohLqDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqDomainRule) implementsZoneAccessGroupUpdateParamsExclude() {
}

type ZoneAccessGroupUpdateParamsExcludePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type ZoneAccessGroupUpdateParamsExcludePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqEveryoneRule) implementsZoneAccessGroupUpdateParamsExclude() {
}

// Matches an IP address block.
type ZoneAccessGroupUpdateParamsExcludePajwohLqIPRule struct {
	IP param.Field[ZoneAccessGroupUpdateParamsExcludePajwohLqIPRuleIP] `json:"ip,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqIPRule) implementsZoneAccessGroupUpdateParamsExclude() {
}

type ZoneAccessGroupUpdateParamsExcludePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type ZoneAccessGroupUpdateParamsExcludePajwohLqIPListRule struct {
	IPList param.Field[ZoneAccessGroupUpdateParamsExcludePajwohLqIPListRuleIPList] `json:"ip_list,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqIPListRule) implementsZoneAccessGroupUpdateParamsExclude() {
}

type ZoneAccessGroupUpdateParamsExcludePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type ZoneAccessGroupUpdateParamsExcludePajwohLqCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqCertificateRule) implementsZoneAccessGroupUpdateParamsExclude() {
}

// Matches an Access group.
type ZoneAccessGroupUpdateParamsExcludePajwohLqAccessGroupRule struct {
	Group param.Field[ZoneAccessGroupUpdateParamsExcludePajwohLqAccessGroupRuleGroup] `json:"group,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqAccessGroupRule) implementsZoneAccessGroupUpdateParamsExclude() {
}

type ZoneAccessGroupUpdateParamsExcludePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZoneAccessGroupUpdateParamsExcludePajwohLqAzureGroupRule struct {
	AzureAd param.Field[ZoneAccessGroupUpdateParamsExcludePajwohLqAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqAzureGroupRule) implementsZoneAccessGroupUpdateParamsExclude() {
}

type ZoneAccessGroupUpdateParamsExcludePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZoneAccessGroupUpdateParamsExcludePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization param.Field[ZoneAccessGroupUpdateParamsExcludePajwohLqGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqGitHubOrganizationRule) implementsZoneAccessGroupUpdateParamsExclude() {
}

type ZoneAccessGroupUpdateParamsExcludePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZoneAccessGroupUpdateParamsExcludePajwohLqGsuiteGroupRule struct {
	Gsuite param.Field[ZoneAccessGroupUpdateParamsExcludePajwohLqGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqGsuiteGroupRule) implementsZoneAccessGroupUpdateParamsExclude() {
}

type ZoneAccessGroupUpdateParamsExcludePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZoneAccessGroupUpdateParamsExcludePajwohLqOktaGroupRule struct {
	Okta param.Field[ZoneAccessGroupUpdateParamsExcludePajwohLqOktaGroupRuleOkta] `json:"okta,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqOktaGroupRule) implementsZoneAccessGroupUpdateParamsExclude() {
}

type ZoneAccessGroupUpdateParamsExcludePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZoneAccessGroupUpdateParamsExcludePajwohLqSamlGroupRule struct {
	Saml param.Field[ZoneAccessGroupUpdateParamsExcludePajwohLqSamlGroupRuleSaml] `json:"saml,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqSamlGroupRule) implementsZoneAccessGroupUpdateParamsExclude() {
}

type ZoneAccessGroupUpdateParamsExcludePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type ZoneAccessGroupUpdateParamsExcludePajwohLqServiceTokenRule struct {
	ServiceToken param.Field[ZoneAccessGroupUpdateParamsExcludePajwohLqServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqServiceTokenRule) implementsZoneAccessGroupUpdateParamsExclude() {
}

type ZoneAccessGroupUpdateParamsExcludePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type ZoneAccessGroupUpdateParamsExcludePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqAnyValidServiceTokenRule) implementsZoneAccessGroupUpdateParamsExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZoneAccessGroupUpdateParamsExcludePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation param.Field[ZoneAccessGroupUpdateParamsExcludePajwohLqExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqExternalEvaluationRule) implementsZoneAccessGroupUpdateParamsExclude() {
}

type ZoneAccessGroupUpdateParamsExcludePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type ZoneAccessGroupUpdateParamsExcludePajwohLqCountryRule struct {
	Geo param.Field[ZoneAccessGroupUpdateParamsExcludePajwohLqCountryRuleGeo] `json:"geo,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqCountryRule) implementsZoneAccessGroupUpdateParamsExclude() {
}

type ZoneAccessGroupUpdateParamsExcludePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type ZoneAccessGroupUpdateParamsExcludePajwohLqAuthenticationMethodRule struct {
	AuthMethod param.Field[ZoneAccessGroupUpdateParamsExcludePajwohLqAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqAuthenticationMethodRule) implementsZoneAccessGroupUpdateParamsExclude() {
}

type ZoneAccessGroupUpdateParamsExcludePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type ZoneAccessGroupUpdateParamsExcludePajwohLqDevicePostureRule struct {
	DevicePosture param.Field[ZoneAccessGroupUpdateParamsExcludePajwohLqDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqDevicePostureRule) implementsZoneAccessGroupUpdateParamsExclude() {
}

type ZoneAccessGroupUpdateParamsExcludePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r ZoneAccessGroupUpdateParamsExcludePajwohLqDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [ZoneAccessGroupUpdateParamsRequirePajwohLqEmailRule],
// [ZoneAccessGroupUpdateParamsRequirePajwohLqEmailListRule],
// [ZoneAccessGroupUpdateParamsRequirePajwohLqDomainRule],
// [ZoneAccessGroupUpdateParamsRequirePajwohLqEveryoneRule],
// [ZoneAccessGroupUpdateParamsRequirePajwohLqIPRule],
// [ZoneAccessGroupUpdateParamsRequirePajwohLqIPListRule],
// [ZoneAccessGroupUpdateParamsRequirePajwohLqCertificateRule],
// [ZoneAccessGroupUpdateParamsRequirePajwohLqAccessGroupRule],
// [ZoneAccessGroupUpdateParamsRequirePajwohLqAzureGroupRule],
// [ZoneAccessGroupUpdateParamsRequirePajwohLqGitHubOrganizationRule],
// [ZoneAccessGroupUpdateParamsRequirePajwohLqGsuiteGroupRule],
// [ZoneAccessGroupUpdateParamsRequirePajwohLqOktaGroupRule],
// [ZoneAccessGroupUpdateParamsRequirePajwohLqSamlGroupRule],
// [ZoneAccessGroupUpdateParamsRequirePajwohLqServiceTokenRule],
// [ZoneAccessGroupUpdateParamsRequirePajwohLqAnyValidServiceTokenRule],
// [ZoneAccessGroupUpdateParamsRequirePajwohLqExternalEvaluationRule],
// [ZoneAccessGroupUpdateParamsRequirePajwohLqCountryRule],
// [ZoneAccessGroupUpdateParamsRequirePajwohLqAuthenticationMethodRule],
// [ZoneAccessGroupUpdateParamsRequirePajwohLqDevicePostureRule].
type ZoneAccessGroupUpdateParamsRequire interface {
	implementsZoneAccessGroupUpdateParamsRequire()
}

// Matches a specific email.
type ZoneAccessGroupUpdateParamsRequirePajwohLqEmailRule struct {
	Email param.Field[ZoneAccessGroupUpdateParamsRequirePajwohLqEmailRuleEmail] `json:"email,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqEmailRule) implementsZoneAccessGroupUpdateParamsRequire() {
}

type ZoneAccessGroupUpdateParamsRequirePajwohLqEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type ZoneAccessGroupUpdateParamsRequirePajwohLqEmailListRule struct {
	EmailList param.Field[ZoneAccessGroupUpdateParamsRequirePajwohLqEmailListRuleEmailList] `json:"email_list,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqEmailListRule) implementsZoneAccessGroupUpdateParamsRequire() {
}

type ZoneAccessGroupUpdateParamsRequirePajwohLqEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type ZoneAccessGroupUpdateParamsRequirePajwohLqDomainRule struct {
	EmailDomain param.Field[ZoneAccessGroupUpdateParamsRequirePajwohLqDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqDomainRule) implementsZoneAccessGroupUpdateParamsRequire() {
}

type ZoneAccessGroupUpdateParamsRequirePajwohLqDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type ZoneAccessGroupUpdateParamsRequirePajwohLqEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqEveryoneRule) implementsZoneAccessGroupUpdateParamsRequire() {
}

// Matches an IP address block.
type ZoneAccessGroupUpdateParamsRequirePajwohLqIPRule struct {
	IP param.Field[ZoneAccessGroupUpdateParamsRequirePajwohLqIPRuleIP] `json:"ip,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqIPRule) implementsZoneAccessGroupUpdateParamsRequire() {
}

type ZoneAccessGroupUpdateParamsRequirePajwohLqIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type ZoneAccessGroupUpdateParamsRequirePajwohLqIPListRule struct {
	IPList param.Field[ZoneAccessGroupUpdateParamsRequirePajwohLqIPListRuleIPList] `json:"ip_list,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqIPListRule) implementsZoneAccessGroupUpdateParamsRequire() {
}

type ZoneAccessGroupUpdateParamsRequirePajwohLqIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type ZoneAccessGroupUpdateParamsRequirePajwohLqCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqCertificateRule) implementsZoneAccessGroupUpdateParamsRequire() {
}

// Matches an Access group.
type ZoneAccessGroupUpdateParamsRequirePajwohLqAccessGroupRule struct {
	Group param.Field[ZoneAccessGroupUpdateParamsRequirePajwohLqAccessGroupRuleGroup] `json:"group,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqAccessGroupRule) implementsZoneAccessGroupUpdateParamsRequire() {
}

type ZoneAccessGroupUpdateParamsRequirePajwohLqAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZoneAccessGroupUpdateParamsRequirePajwohLqAzureGroupRule struct {
	AzureAd param.Field[ZoneAccessGroupUpdateParamsRequirePajwohLqAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqAzureGroupRule) implementsZoneAccessGroupUpdateParamsRequire() {
}

type ZoneAccessGroupUpdateParamsRequirePajwohLqAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZoneAccessGroupUpdateParamsRequirePajwohLqGitHubOrganizationRule struct {
	GitHubOrganization param.Field[ZoneAccessGroupUpdateParamsRequirePajwohLqGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqGitHubOrganizationRule) implementsZoneAccessGroupUpdateParamsRequire() {
}

type ZoneAccessGroupUpdateParamsRequirePajwohLqGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZoneAccessGroupUpdateParamsRequirePajwohLqGsuiteGroupRule struct {
	Gsuite param.Field[ZoneAccessGroupUpdateParamsRequirePajwohLqGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqGsuiteGroupRule) implementsZoneAccessGroupUpdateParamsRequire() {
}

type ZoneAccessGroupUpdateParamsRequirePajwohLqGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZoneAccessGroupUpdateParamsRequirePajwohLqOktaGroupRule struct {
	Okta param.Field[ZoneAccessGroupUpdateParamsRequirePajwohLqOktaGroupRuleOkta] `json:"okta,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqOktaGroupRule) implementsZoneAccessGroupUpdateParamsRequire() {
}

type ZoneAccessGroupUpdateParamsRequirePajwohLqOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZoneAccessGroupUpdateParamsRequirePajwohLqSamlGroupRule struct {
	Saml param.Field[ZoneAccessGroupUpdateParamsRequirePajwohLqSamlGroupRuleSaml] `json:"saml,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqSamlGroupRule) implementsZoneAccessGroupUpdateParamsRequire() {
}

type ZoneAccessGroupUpdateParamsRequirePajwohLqSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type ZoneAccessGroupUpdateParamsRequirePajwohLqServiceTokenRule struct {
	ServiceToken param.Field[ZoneAccessGroupUpdateParamsRequirePajwohLqServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqServiceTokenRule) implementsZoneAccessGroupUpdateParamsRequire() {
}

type ZoneAccessGroupUpdateParamsRequirePajwohLqServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type ZoneAccessGroupUpdateParamsRequirePajwohLqAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqAnyValidServiceTokenRule) implementsZoneAccessGroupUpdateParamsRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZoneAccessGroupUpdateParamsRequirePajwohLqExternalEvaluationRule struct {
	ExternalEvaluation param.Field[ZoneAccessGroupUpdateParamsRequirePajwohLqExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqExternalEvaluationRule) implementsZoneAccessGroupUpdateParamsRequire() {
}

type ZoneAccessGroupUpdateParamsRequirePajwohLqExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type ZoneAccessGroupUpdateParamsRequirePajwohLqCountryRule struct {
	Geo param.Field[ZoneAccessGroupUpdateParamsRequirePajwohLqCountryRuleGeo] `json:"geo,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqCountryRule) implementsZoneAccessGroupUpdateParamsRequire() {
}

type ZoneAccessGroupUpdateParamsRequirePajwohLqCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type ZoneAccessGroupUpdateParamsRequirePajwohLqAuthenticationMethodRule struct {
	AuthMethod param.Field[ZoneAccessGroupUpdateParamsRequirePajwohLqAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqAuthenticationMethodRule) implementsZoneAccessGroupUpdateParamsRequire() {
}

type ZoneAccessGroupUpdateParamsRequirePajwohLqAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type ZoneAccessGroupUpdateParamsRequirePajwohLqDevicePostureRule struct {
	DevicePosture param.Field[ZoneAccessGroupUpdateParamsRequirePajwohLqDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqDevicePostureRule) implementsZoneAccessGroupUpdateParamsRequire() {
}

type ZoneAccessGroupUpdateParamsRequirePajwohLqDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r ZoneAccessGroupUpdateParamsRequirePajwohLqDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
