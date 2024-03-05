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

// ZeroTrustAccessGroupService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustAccessGroupService]
// method instead.
type ZeroTrustAccessGroupService struct {
	Options []option.RequestOption
}

// NewZeroTrustAccessGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustAccessGroupService(opts ...option.RequestOption) (r *ZeroTrustAccessGroupService) {
	r = &ZeroTrustAccessGroupService{}
	r.Options = opts
	return
}

// Creates a new Access group.
func (r *ZeroTrustAccessGroupService) New(ctx context.Context, params ZeroTrustAccessGroupNewParams, opts ...option.RequestOption) (res *AccessGroups, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessGroupNewResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/groups", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured Access group.
func (r *ZeroTrustAccessGroupService) Update(ctx context.Context, uuid string, params ZeroTrustAccessGroupUpdateParams, opts ...option.RequestOption) (res *AccessGroups, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessGroupUpdateResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/groups/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all Access groups.
func (r *ZeroTrustAccessGroupService) List(ctx context.Context, query ZeroTrustAccessGroupListParams, opts ...option.RequestOption) (res *[]AccessGroups, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessGroupListResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/groups", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an Access group.
func (r *ZeroTrustAccessGroupService) Delete(ctx context.Context, uuid string, body ZeroTrustAccessGroupDeleteParams, opts ...option.RequestOption) (res *ZeroTrustAccessGroupDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessGroupDeleteResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/groups/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Access group.
func (r *ZeroTrustAccessGroupService) Get(ctx context.Context, uuid string, query ZeroTrustAccessGroupGetParams, opts ...option.RequestOption) (res *AccessGroups, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessGroupGetResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/groups/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessGroups struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessGroupsExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessGroupsInclude `json:"include"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	IsDefault []AccessGroupsIsDefault `json:"is_default"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []AccessGroupsRequire `json:"require"`
	UpdatedAt time.Time             `json:"updated_at" format:"date-time"`
	JSON      accessGroupsJSON      `json:"-"`
}

// accessGroupsJSON contains the JSON metadata for the struct [AccessGroups]
type accessGroupsJSON struct {
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

func (r *AccessGroups) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupsExcludeAccessEmailRule],
// [AccessGroupsExcludeAccessEmailListRule], [AccessGroupsExcludeAccessDomainRule],
// [AccessGroupsExcludeAccessEveryoneRule], [AccessGroupsExcludeAccessIPRule],
// [AccessGroupsExcludeAccessIPListRule],
// [AccessGroupsExcludeAccessCertificateRule],
// [AccessGroupsExcludeAccessAccessGroupRule],
// [AccessGroupsExcludeAccessAzureGroupRule],
// [AccessGroupsExcludeAccessGitHubOrganizationRule],
// [AccessGroupsExcludeAccessGsuiteGroupRule],
// [AccessGroupsExcludeAccessOktaGroupRule],
// [AccessGroupsExcludeAccessSamlGroupRule],
// [AccessGroupsExcludeAccessServiceTokenRule],
// [AccessGroupsExcludeAccessAnyValidServiceTokenRule],
// [AccessGroupsExcludeAccessExternalEvaluationRule],
// [AccessGroupsExcludeAccessCountryRule],
// [AccessGroupsExcludeAccessAuthenticationMethodRule] or
// [AccessGroupsExcludeAccessDevicePostureRule].
type AccessGroupsExclude interface {
	implementsAccessGroupsExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupsExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupsExcludeAccessEmailRule struct {
	Email AccessGroupsExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupsExcludeAccessEmailRuleJSON  `json:"-"`
}

// accessGroupsExcludeAccessEmailRuleJSON contains the JSON metadata for the struct
// [AccessGroupsExcludeAccessEmailRule]
type accessGroupsExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsExcludeAccessEmailRule) implementsAccessGroupsExclude() {}

type AccessGroupsExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                      `json:"email,required" format:"email"`
	JSON  accessGroupsExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupsExcludeAccessEmailRuleEmailJSON contains the JSON metadata for the
// struct [AccessGroupsExcludeAccessEmailRuleEmail]
type accessGroupsExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupsExcludeAccessEmailListRule struct {
	EmailList AccessGroupsExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupsExcludeAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupsExcludeAccessEmailListRuleJSON contains the JSON metadata for the
// struct [AccessGroupsExcludeAccessEmailListRule]
type accessGroupsExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsExcludeAccessEmailListRule) implementsAccessGroupsExclude() {}

type AccessGroupsExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                              `json:"id,required"`
	JSON accessGroupsExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupsExcludeAccessEmailListRuleEmailListJSON contains the JSON metadata
// for the struct [AccessGroupsExcludeAccessEmailListRuleEmailList]
type accessGroupsExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupsExcludeAccessDomainRule struct {
	EmailDomain AccessGroupsExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupsExcludeAccessDomainRuleJSON        `json:"-"`
}

// accessGroupsExcludeAccessDomainRuleJSON contains the JSON metadata for the
// struct [AccessGroupsExcludeAccessDomainRule]
type accessGroupsExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsExcludeAccessDomainRule) implementsAccessGroupsExclude() {}

type AccessGroupsExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                             `json:"domain,required"`
	JSON   accessGroupsExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupsExcludeAccessDomainRuleEmailDomainJSON contains the JSON metadata
// for the struct [AccessGroupsExcludeAccessDomainRuleEmailDomain]
type accessGroupsExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupsExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                               `json:"everyone,required"`
	JSON     accessGroupsExcludeAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupsExcludeAccessEveryoneRuleJSON contains the JSON metadata for the
// struct [AccessGroupsExcludeAccessEveryoneRule]
type accessGroupsExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsExcludeAccessEveryoneRule) implementsAccessGroupsExclude() {}

// Matches an IP address block.
type AccessGroupsExcludeAccessIPRule struct {
	IP   AccessGroupsExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupsExcludeAccessIPRuleJSON `json:"-"`
}

// accessGroupsExcludeAccessIPRuleJSON contains the JSON metadata for the struct
// [AccessGroupsExcludeAccessIPRule]
type accessGroupsExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsExcludeAccessIPRule) implementsAccessGroupsExclude() {}

type AccessGroupsExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                `json:"ip,required"`
	JSON accessGroupsExcludeAccessIPRuleIPJSON `json:"-"`
}

// accessGroupsExcludeAccessIPRuleIPJSON contains the JSON metadata for the struct
// [AccessGroupsExcludeAccessIPRuleIP]
type accessGroupsExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupsExcludeAccessIPListRule struct {
	IPList AccessGroupsExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupsExcludeAccessIPListRuleJSON   `json:"-"`
}

// accessGroupsExcludeAccessIPListRuleJSON contains the JSON metadata for the
// struct [AccessGroupsExcludeAccessIPListRule]
type accessGroupsExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsExcludeAccessIPListRule) implementsAccessGroupsExclude() {}

type AccessGroupsExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                        `json:"id,required"`
	JSON accessGroupsExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupsExcludeAccessIPListRuleIPListJSON contains the JSON metadata for the
// struct [AccessGroupsExcludeAccessIPListRuleIPList]
type accessGroupsExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupsExcludeAccessCertificateRule struct {
	Certificate interface{}                                  `json:"certificate,required"`
	JSON        accessGroupsExcludeAccessCertificateRuleJSON `json:"-"`
}

// accessGroupsExcludeAccessCertificateRuleJSON contains the JSON metadata for the
// struct [AccessGroupsExcludeAccessCertificateRule]
type accessGroupsExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsExcludeAccessCertificateRule) implementsAccessGroupsExclude() {}

// Matches an Access group.
type AccessGroupsExcludeAccessAccessGroupRule struct {
	Group AccessGroupsExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupsExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupsExcludeAccessAccessGroupRuleJSON contains the JSON metadata for the
// struct [AccessGroupsExcludeAccessAccessGroupRule]
type accessGroupsExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsExcludeAccessAccessGroupRule) implementsAccessGroupsExclude() {}

type AccessGroupsExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                            `json:"id,required"`
	JSON accessGroupsExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupsExcludeAccessAccessGroupRuleGroupJSON contains the JSON metadata for
// the struct [AccessGroupsExcludeAccessAccessGroupRuleGroup]
type accessGroupsExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupsExcludeAccessAzureGroupRule struct {
	AzureAd AccessGroupsExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupsExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupsExcludeAccessAzureGroupRuleJSON contains the JSON metadata for the
// struct [AccessGroupsExcludeAccessAzureGroupRule]
type accessGroupsExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsExcludeAccessAzureGroupRule) implementsAccessGroupsExclude() {}

type AccessGroupsExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                             `json:"connection_id,required"`
	JSON         accessGroupsExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupsExcludeAccessAzureGroupRuleAzureAdJSON contains the JSON metadata
// for the struct [AccessGroupsExcludeAccessAzureGroupRuleAzureAd]
type accessGroupsExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupsExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupsExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupsExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupsExcludeAccessGitHubOrganizationRuleJSON contains the JSON metadata
// for the struct [AccessGroupsExcludeAccessGitHubOrganizationRule]
type accessGroupsExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsExcludeAccessGitHubOrganizationRule) implementsAccessGroupsExclude() {}

type AccessGroupsExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                `json:"name,required"`
	JSON accessGroupsExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupsExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON contains
// the JSON metadata for the struct
// [AccessGroupsExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupsExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupsExcludeAccessGsuiteGroupRule struct {
	Gsuite AccessGroupsExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupsExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupsExcludeAccessGsuiteGroupRuleJSON contains the JSON metadata for the
// struct [AccessGroupsExcludeAccessGsuiteGroupRule]
type accessGroupsExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsExcludeAccessGsuiteGroupRule) implementsAccessGroupsExclude() {}

type AccessGroupsExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                             `json:"email,required"`
	JSON  accessGroupsExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupsExcludeAccessGsuiteGroupRuleGsuiteJSON contains the JSON metadata
// for the struct [AccessGroupsExcludeAccessGsuiteGroupRuleGsuite]
type accessGroupsExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupsExcludeAccessOktaGroupRule struct {
	Okta AccessGroupsExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupsExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupsExcludeAccessOktaGroupRuleJSON contains the JSON metadata for the
// struct [AccessGroupsExcludeAccessOktaGroupRule]
type accessGroupsExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsExcludeAccessOktaGroupRule) implementsAccessGroupsExclude() {}

type AccessGroupsExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                         `json:"email,required"`
	JSON  accessGroupsExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupsExcludeAccessOktaGroupRuleOktaJSON contains the JSON metadata for
// the struct [AccessGroupsExcludeAccessOktaGroupRuleOkta]
type accessGroupsExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupsExcludeAccessSamlGroupRule struct {
	Saml AccessGroupsExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupsExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupsExcludeAccessSamlGroupRuleJSON contains the JSON metadata for the
// struct [AccessGroupsExcludeAccessSamlGroupRule]
type accessGroupsExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsExcludeAccessSamlGroupRule) implementsAccessGroupsExclude() {}

type AccessGroupsExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                         `json:"attribute_value,required"`
	JSON           accessGroupsExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupsExcludeAccessSamlGroupRuleSamlJSON contains the JSON metadata for
// the struct [AccessGroupsExcludeAccessSamlGroupRuleSaml]
type accessGroupsExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupsExcludeAccessServiceTokenRule struct {
	ServiceToken AccessGroupsExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupsExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupsExcludeAccessServiceTokenRuleJSON contains the JSON metadata for the
// struct [AccessGroupsExcludeAccessServiceTokenRule]
type accessGroupsExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsExcludeAccessServiceTokenRule) implementsAccessGroupsExclude() {}

type AccessGroupsExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                    `json:"token_id,required"`
	JSON    accessGroupsExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupsExcludeAccessServiceTokenRuleServiceTokenJSON contains the JSON
// metadata for the struct [AccessGroupsExcludeAccessServiceTokenRuleServiceToken]
type accessGroupsExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupsExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                           `json:"any_valid_service_token,required"`
	JSON                 accessGroupsExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupsExcludeAccessAnyValidServiceTokenRuleJSON contains the JSON metadata
// for the struct [AccessGroupsExcludeAccessAnyValidServiceTokenRule]
type accessGroupsExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsExcludeAccessAnyValidServiceTokenRule) implementsAccessGroupsExclude() {}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupsExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupsExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupsExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupsExcludeAccessExternalEvaluationRuleJSON contains the JSON metadata
// for the struct [AccessGroupsExcludeAccessExternalEvaluationRule]
type accessGroupsExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsExcludeAccessExternalEvaluationRule) implementsAccessGroupsExclude() {}

type AccessGroupsExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                `json:"keys_url,required"`
	JSON    accessGroupsExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupsExcludeAccessExternalEvaluationRuleExternalEvaluationJSON contains
// the JSON metadata for the struct
// [AccessGroupsExcludeAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupsExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupsExcludeAccessCountryRule struct {
	Geo  AccessGroupsExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupsExcludeAccessCountryRuleJSON `json:"-"`
}

// accessGroupsExcludeAccessCountryRuleJSON contains the JSON metadata for the
// struct [AccessGroupsExcludeAccessCountryRule]
type accessGroupsExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsExcludeAccessCountryRule) implementsAccessGroupsExclude() {}

type AccessGroupsExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                      `json:"country_code,required"`
	JSON        accessGroupsExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupsExcludeAccessCountryRuleGeoJSON contains the JSON metadata for the
// struct [AccessGroupsExcludeAccessCountryRuleGeo]
type accessGroupsExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupsExcludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupsExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupsExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupsExcludeAccessAuthenticationMethodRuleJSON contains the JSON metadata
// for the struct [AccessGroupsExcludeAccessAuthenticationMethodRule]
type accessGroupsExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsExcludeAccessAuthenticationMethodRule) implementsAccessGroupsExclude() {}

type AccessGroupsExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                          `json:"auth_method,required"`
	JSON       accessGroupsExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupsExcludeAccessAuthenticationMethodRuleAuthMethodJSON contains the
// JSON metadata for the struct
// [AccessGroupsExcludeAccessAuthenticationMethodRuleAuthMethod]
type accessGroupsExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupsExcludeAccessDevicePostureRule struct {
	DevicePosture AccessGroupsExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupsExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupsExcludeAccessDevicePostureRuleJSON contains the JSON metadata for
// the struct [AccessGroupsExcludeAccessDevicePostureRule]
type accessGroupsExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsExcludeAccessDevicePostureRule) implementsAccessGroupsExclude() {}

type AccessGroupsExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                      `json:"integration_uid,required"`
	JSON           accessGroupsExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupsExcludeAccessDevicePostureRuleDevicePostureJSON contains the JSON
// metadata for the struct
// [AccessGroupsExcludeAccessDevicePostureRuleDevicePosture]
type accessGroupsExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupsExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupsIncludeAccessEmailRule],
// [AccessGroupsIncludeAccessEmailListRule], [AccessGroupsIncludeAccessDomainRule],
// [AccessGroupsIncludeAccessEveryoneRule], [AccessGroupsIncludeAccessIPRule],
// [AccessGroupsIncludeAccessIPListRule],
// [AccessGroupsIncludeAccessCertificateRule],
// [AccessGroupsIncludeAccessAccessGroupRule],
// [AccessGroupsIncludeAccessAzureGroupRule],
// [AccessGroupsIncludeAccessGitHubOrganizationRule],
// [AccessGroupsIncludeAccessGsuiteGroupRule],
// [AccessGroupsIncludeAccessOktaGroupRule],
// [AccessGroupsIncludeAccessSamlGroupRule],
// [AccessGroupsIncludeAccessServiceTokenRule],
// [AccessGroupsIncludeAccessAnyValidServiceTokenRule],
// [AccessGroupsIncludeAccessExternalEvaluationRule],
// [AccessGroupsIncludeAccessCountryRule],
// [AccessGroupsIncludeAccessAuthenticationMethodRule] or
// [AccessGroupsIncludeAccessDevicePostureRule].
type AccessGroupsInclude interface {
	implementsAccessGroupsInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupsInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupsIncludeAccessEmailRule struct {
	Email AccessGroupsIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupsIncludeAccessEmailRuleJSON  `json:"-"`
}

// accessGroupsIncludeAccessEmailRuleJSON contains the JSON metadata for the struct
// [AccessGroupsIncludeAccessEmailRule]
type accessGroupsIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIncludeAccessEmailRule) implementsAccessGroupsInclude() {}

type AccessGroupsIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                      `json:"email,required" format:"email"`
	JSON  accessGroupsIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupsIncludeAccessEmailRuleEmailJSON contains the JSON metadata for the
// struct [AccessGroupsIncludeAccessEmailRuleEmail]
type accessGroupsIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupsIncludeAccessEmailListRule struct {
	EmailList AccessGroupsIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupsIncludeAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupsIncludeAccessEmailListRuleJSON contains the JSON metadata for the
// struct [AccessGroupsIncludeAccessEmailListRule]
type accessGroupsIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIncludeAccessEmailListRule) implementsAccessGroupsInclude() {}

type AccessGroupsIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                              `json:"id,required"`
	JSON accessGroupsIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupsIncludeAccessEmailListRuleEmailListJSON contains the JSON metadata
// for the struct [AccessGroupsIncludeAccessEmailListRuleEmailList]
type accessGroupsIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupsIncludeAccessDomainRule struct {
	EmailDomain AccessGroupsIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupsIncludeAccessDomainRuleJSON        `json:"-"`
}

// accessGroupsIncludeAccessDomainRuleJSON contains the JSON metadata for the
// struct [AccessGroupsIncludeAccessDomainRule]
type accessGroupsIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIncludeAccessDomainRule) implementsAccessGroupsInclude() {}

type AccessGroupsIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                             `json:"domain,required"`
	JSON   accessGroupsIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupsIncludeAccessDomainRuleEmailDomainJSON contains the JSON metadata
// for the struct [AccessGroupsIncludeAccessDomainRuleEmailDomain]
type accessGroupsIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupsIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                               `json:"everyone,required"`
	JSON     accessGroupsIncludeAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupsIncludeAccessEveryoneRuleJSON contains the JSON metadata for the
// struct [AccessGroupsIncludeAccessEveryoneRule]
type accessGroupsIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIncludeAccessEveryoneRule) implementsAccessGroupsInclude() {}

// Matches an IP address block.
type AccessGroupsIncludeAccessIPRule struct {
	IP   AccessGroupsIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupsIncludeAccessIPRuleJSON `json:"-"`
}

// accessGroupsIncludeAccessIPRuleJSON contains the JSON metadata for the struct
// [AccessGroupsIncludeAccessIPRule]
type accessGroupsIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIncludeAccessIPRule) implementsAccessGroupsInclude() {}

type AccessGroupsIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                `json:"ip,required"`
	JSON accessGroupsIncludeAccessIPRuleIPJSON `json:"-"`
}

// accessGroupsIncludeAccessIPRuleIPJSON contains the JSON metadata for the struct
// [AccessGroupsIncludeAccessIPRuleIP]
type accessGroupsIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupsIncludeAccessIPListRule struct {
	IPList AccessGroupsIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupsIncludeAccessIPListRuleJSON   `json:"-"`
}

// accessGroupsIncludeAccessIPListRuleJSON contains the JSON metadata for the
// struct [AccessGroupsIncludeAccessIPListRule]
type accessGroupsIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIncludeAccessIPListRule) implementsAccessGroupsInclude() {}

type AccessGroupsIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                        `json:"id,required"`
	JSON accessGroupsIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupsIncludeAccessIPListRuleIPListJSON contains the JSON metadata for the
// struct [AccessGroupsIncludeAccessIPListRuleIPList]
type accessGroupsIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupsIncludeAccessCertificateRule struct {
	Certificate interface{}                                  `json:"certificate,required"`
	JSON        accessGroupsIncludeAccessCertificateRuleJSON `json:"-"`
}

// accessGroupsIncludeAccessCertificateRuleJSON contains the JSON metadata for the
// struct [AccessGroupsIncludeAccessCertificateRule]
type accessGroupsIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIncludeAccessCertificateRule) implementsAccessGroupsInclude() {}

// Matches an Access group.
type AccessGroupsIncludeAccessAccessGroupRule struct {
	Group AccessGroupsIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupsIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupsIncludeAccessAccessGroupRuleJSON contains the JSON metadata for the
// struct [AccessGroupsIncludeAccessAccessGroupRule]
type accessGroupsIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIncludeAccessAccessGroupRule) implementsAccessGroupsInclude() {}

type AccessGroupsIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                            `json:"id,required"`
	JSON accessGroupsIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupsIncludeAccessAccessGroupRuleGroupJSON contains the JSON metadata for
// the struct [AccessGroupsIncludeAccessAccessGroupRuleGroup]
type accessGroupsIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupsIncludeAccessAzureGroupRule struct {
	AzureAd AccessGroupsIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupsIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupsIncludeAccessAzureGroupRuleJSON contains the JSON metadata for the
// struct [AccessGroupsIncludeAccessAzureGroupRule]
type accessGroupsIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIncludeAccessAzureGroupRule) implementsAccessGroupsInclude() {}

type AccessGroupsIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                             `json:"connection_id,required"`
	JSON         accessGroupsIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupsIncludeAccessAzureGroupRuleAzureAdJSON contains the JSON metadata
// for the struct [AccessGroupsIncludeAccessAzureGroupRuleAzureAd]
type accessGroupsIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupsIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupsIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupsIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupsIncludeAccessGitHubOrganizationRuleJSON contains the JSON metadata
// for the struct [AccessGroupsIncludeAccessGitHubOrganizationRule]
type accessGroupsIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIncludeAccessGitHubOrganizationRule) implementsAccessGroupsInclude() {}

type AccessGroupsIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                `json:"name,required"`
	JSON accessGroupsIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupsIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON contains
// the JSON metadata for the struct
// [AccessGroupsIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupsIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupsIncludeAccessGsuiteGroupRule struct {
	Gsuite AccessGroupsIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupsIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupsIncludeAccessGsuiteGroupRuleJSON contains the JSON metadata for the
// struct [AccessGroupsIncludeAccessGsuiteGroupRule]
type accessGroupsIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIncludeAccessGsuiteGroupRule) implementsAccessGroupsInclude() {}

type AccessGroupsIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                             `json:"email,required"`
	JSON  accessGroupsIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupsIncludeAccessGsuiteGroupRuleGsuiteJSON contains the JSON metadata
// for the struct [AccessGroupsIncludeAccessGsuiteGroupRuleGsuite]
type accessGroupsIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupsIncludeAccessOktaGroupRule struct {
	Okta AccessGroupsIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupsIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupsIncludeAccessOktaGroupRuleJSON contains the JSON metadata for the
// struct [AccessGroupsIncludeAccessOktaGroupRule]
type accessGroupsIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIncludeAccessOktaGroupRule) implementsAccessGroupsInclude() {}

type AccessGroupsIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                         `json:"email,required"`
	JSON  accessGroupsIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupsIncludeAccessOktaGroupRuleOktaJSON contains the JSON metadata for
// the struct [AccessGroupsIncludeAccessOktaGroupRuleOkta]
type accessGroupsIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupsIncludeAccessSamlGroupRule struct {
	Saml AccessGroupsIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupsIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupsIncludeAccessSamlGroupRuleJSON contains the JSON metadata for the
// struct [AccessGroupsIncludeAccessSamlGroupRule]
type accessGroupsIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIncludeAccessSamlGroupRule) implementsAccessGroupsInclude() {}

type AccessGroupsIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                         `json:"attribute_value,required"`
	JSON           accessGroupsIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupsIncludeAccessSamlGroupRuleSamlJSON contains the JSON metadata for
// the struct [AccessGroupsIncludeAccessSamlGroupRuleSaml]
type accessGroupsIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupsIncludeAccessServiceTokenRule struct {
	ServiceToken AccessGroupsIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupsIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupsIncludeAccessServiceTokenRuleJSON contains the JSON metadata for the
// struct [AccessGroupsIncludeAccessServiceTokenRule]
type accessGroupsIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIncludeAccessServiceTokenRule) implementsAccessGroupsInclude() {}

type AccessGroupsIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                    `json:"token_id,required"`
	JSON    accessGroupsIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupsIncludeAccessServiceTokenRuleServiceTokenJSON contains the JSON
// metadata for the struct [AccessGroupsIncludeAccessServiceTokenRuleServiceToken]
type accessGroupsIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupsIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                           `json:"any_valid_service_token,required"`
	JSON                 accessGroupsIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupsIncludeAccessAnyValidServiceTokenRuleJSON contains the JSON metadata
// for the struct [AccessGroupsIncludeAccessAnyValidServiceTokenRule]
type accessGroupsIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIncludeAccessAnyValidServiceTokenRule) implementsAccessGroupsInclude() {}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupsIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupsIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupsIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupsIncludeAccessExternalEvaluationRuleJSON contains the JSON metadata
// for the struct [AccessGroupsIncludeAccessExternalEvaluationRule]
type accessGroupsIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIncludeAccessExternalEvaluationRule) implementsAccessGroupsInclude() {}

type AccessGroupsIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                `json:"keys_url,required"`
	JSON    accessGroupsIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupsIncludeAccessExternalEvaluationRuleExternalEvaluationJSON contains
// the JSON metadata for the struct
// [AccessGroupsIncludeAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupsIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupsIncludeAccessCountryRule struct {
	Geo  AccessGroupsIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupsIncludeAccessCountryRuleJSON `json:"-"`
}

// accessGroupsIncludeAccessCountryRuleJSON contains the JSON metadata for the
// struct [AccessGroupsIncludeAccessCountryRule]
type accessGroupsIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIncludeAccessCountryRule) implementsAccessGroupsInclude() {}

type AccessGroupsIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                      `json:"country_code,required"`
	JSON        accessGroupsIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupsIncludeAccessCountryRuleGeoJSON contains the JSON metadata for the
// struct [AccessGroupsIncludeAccessCountryRuleGeo]
type accessGroupsIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupsIncludeAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupsIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupsIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupsIncludeAccessAuthenticationMethodRuleJSON contains the JSON metadata
// for the struct [AccessGroupsIncludeAccessAuthenticationMethodRule]
type accessGroupsIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIncludeAccessAuthenticationMethodRule) implementsAccessGroupsInclude() {}

type AccessGroupsIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                          `json:"auth_method,required"`
	JSON       accessGroupsIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupsIncludeAccessAuthenticationMethodRuleAuthMethodJSON contains the
// JSON metadata for the struct
// [AccessGroupsIncludeAccessAuthenticationMethodRuleAuthMethod]
type accessGroupsIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupsIncludeAccessDevicePostureRule struct {
	DevicePosture AccessGroupsIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupsIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupsIncludeAccessDevicePostureRuleJSON contains the JSON metadata for
// the struct [AccessGroupsIncludeAccessDevicePostureRule]
type accessGroupsIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIncludeAccessDevicePostureRule) implementsAccessGroupsInclude() {}

type AccessGroupsIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                      `json:"integration_uid,required"`
	JSON           accessGroupsIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupsIncludeAccessDevicePostureRuleDevicePostureJSON contains the JSON
// metadata for the struct
// [AccessGroupsIncludeAccessDevicePostureRuleDevicePosture]
type accessGroupsIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupsIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupsIsDefaultAccessEmailRule],
// [AccessGroupsIsDefaultAccessEmailListRule],
// [AccessGroupsIsDefaultAccessDomainRule],
// [AccessGroupsIsDefaultAccessEveryoneRule], [AccessGroupsIsDefaultAccessIPRule],
// [AccessGroupsIsDefaultAccessIPListRule],
// [AccessGroupsIsDefaultAccessCertificateRule],
// [AccessGroupsIsDefaultAccessAccessGroupRule],
// [AccessGroupsIsDefaultAccessAzureGroupRule],
// [AccessGroupsIsDefaultAccessGitHubOrganizationRule],
// [AccessGroupsIsDefaultAccessGsuiteGroupRule],
// [AccessGroupsIsDefaultAccessOktaGroupRule],
// [AccessGroupsIsDefaultAccessSamlGroupRule],
// [AccessGroupsIsDefaultAccessServiceTokenRule],
// [AccessGroupsIsDefaultAccessAnyValidServiceTokenRule],
// [AccessGroupsIsDefaultAccessExternalEvaluationRule],
// [AccessGroupsIsDefaultAccessCountryRule],
// [AccessGroupsIsDefaultAccessAuthenticationMethodRule] or
// [AccessGroupsIsDefaultAccessDevicePostureRule].
type AccessGroupsIsDefault interface {
	implementsAccessGroupsIsDefault()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupsIsDefault)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupsIsDefaultAccessEmailRule struct {
	Email AccessGroupsIsDefaultAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupsIsDefaultAccessEmailRuleJSON  `json:"-"`
}

// accessGroupsIsDefaultAccessEmailRuleJSON contains the JSON metadata for the
// struct [AccessGroupsIsDefaultAccessEmailRule]
type accessGroupsIsDefaultAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIsDefaultAccessEmailRule) implementsAccessGroupsIsDefault() {}

type AccessGroupsIsDefaultAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                        `json:"email,required" format:"email"`
	JSON  accessGroupsIsDefaultAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupsIsDefaultAccessEmailRuleEmailJSON contains the JSON metadata for the
// struct [AccessGroupsIsDefaultAccessEmailRuleEmail]
type accessGroupsIsDefaultAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupsIsDefaultAccessEmailListRule struct {
	EmailList AccessGroupsIsDefaultAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupsIsDefaultAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupsIsDefaultAccessEmailListRuleJSON contains the JSON metadata for the
// struct [AccessGroupsIsDefaultAccessEmailListRule]
type accessGroupsIsDefaultAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIsDefaultAccessEmailListRule) implementsAccessGroupsIsDefault() {}

type AccessGroupsIsDefaultAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                `json:"id,required"`
	JSON accessGroupsIsDefaultAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupsIsDefaultAccessEmailListRuleEmailListJSON contains the JSON metadata
// for the struct [AccessGroupsIsDefaultAccessEmailListRuleEmailList]
type accessGroupsIsDefaultAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupsIsDefaultAccessDomainRule struct {
	EmailDomain AccessGroupsIsDefaultAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupsIsDefaultAccessDomainRuleJSON        `json:"-"`
}

// accessGroupsIsDefaultAccessDomainRuleJSON contains the JSON metadata for the
// struct [AccessGroupsIsDefaultAccessDomainRule]
type accessGroupsIsDefaultAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIsDefaultAccessDomainRule) implementsAccessGroupsIsDefault() {}

type AccessGroupsIsDefaultAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                               `json:"domain,required"`
	JSON   accessGroupsIsDefaultAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupsIsDefaultAccessDomainRuleEmailDomainJSON contains the JSON metadata
// for the struct [AccessGroupsIsDefaultAccessDomainRuleEmailDomain]
type accessGroupsIsDefaultAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupsIsDefaultAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                 `json:"everyone,required"`
	JSON     accessGroupsIsDefaultAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupsIsDefaultAccessEveryoneRuleJSON contains the JSON metadata for the
// struct [AccessGroupsIsDefaultAccessEveryoneRule]
type accessGroupsIsDefaultAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIsDefaultAccessEveryoneRule) implementsAccessGroupsIsDefault() {}

// Matches an IP address block.
type AccessGroupsIsDefaultAccessIPRule struct {
	IP   AccessGroupsIsDefaultAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupsIsDefaultAccessIPRuleJSON `json:"-"`
}

// accessGroupsIsDefaultAccessIPRuleJSON contains the JSON metadata for the struct
// [AccessGroupsIsDefaultAccessIPRule]
type accessGroupsIsDefaultAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIsDefaultAccessIPRule) implementsAccessGroupsIsDefault() {}

type AccessGroupsIsDefaultAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                  `json:"ip,required"`
	JSON accessGroupsIsDefaultAccessIPRuleIPJSON `json:"-"`
}

// accessGroupsIsDefaultAccessIPRuleIPJSON contains the JSON metadata for the
// struct [AccessGroupsIsDefaultAccessIPRuleIP]
type accessGroupsIsDefaultAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupsIsDefaultAccessIPListRule struct {
	IPList AccessGroupsIsDefaultAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupsIsDefaultAccessIPListRuleJSON   `json:"-"`
}

// accessGroupsIsDefaultAccessIPListRuleJSON contains the JSON metadata for the
// struct [AccessGroupsIsDefaultAccessIPListRule]
type accessGroupsIsDefaultAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIsDefaultAccessIPListRule) implementsAccessGroupsIsDefault() {}

type AccessGroupsIsDefaultAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                          `json:"id,required"`
	JSON accessGroupsIsDefaultAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupsIsDefaultAccessIPListRuleIPListJSON contains the JSON metadata for
// the struct [AccessGroupsIsDefaultAccessIPListRuleIPList]
type accessGroupsIsDefaultAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupsIsDefaultAccessCertificateRule struct {
	Certificate interface{}                                    `json:"certificate,required"`
	JSON        accessGroupsIsDefaultAccessCertificateRuleJSON `json:"-"`
}

// accessGroupsIsDefaultAccessCertificateRuleJSON contains the JSON metadata for
// the struct [AccessGroupsIsDefaultAccessCertificateRule]
type accessGroupsIsDefaultAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIsDefaultAccessCertificateRule) implementsAccessGroupsIsDefault() {}

// Matches an Access group.
type AccessGroupsIsDefaultAccessAccessGroupRule struct {
	Group AccessGroupsIsDefaultAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupsIsDefaultAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupsIsDefaultAccessAccessGroupRuleJSON contains the JSON metadata for
// the struct [AccessGroupsIsDefaultAccessAccessGroupRule]
type accessGroupsIsDefaultAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIsDefaultAccessAccessGroupRule) implementsAccessGroupsIsDefault() {}

type AccessGroupsIsDefaultAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                              `json:"id,required"`
	JSON accessGroupsIsDefaultAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupsIsDefaultAccessAccessGroupRuleGroupJSON contains the JSON metadata
// for the struct [AccessGroupsIsDefaultAccessAccessGroupRuleGroup]
type accessGroupsIsDefaultAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupsIsDefaultAccessAzureGroupRule struct {
	AzureAd AccessGroupsIsDefaultAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupsIsDefaultAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupsIsDefaultAccessAzureGroupRuleJSON contains the JSON metadata for the
// struct [AccessGroupsIsDefaultAccessAzureGroupRule]
type accessGroupsIsDefaultAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIsDefaultAccessAzureGroupRule) implementsAccessGroupsIsDefault() {}

type AccessGroupsIsDefaultAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                               `json:"connection_id,required"`
	JSON         accessGroupsIsDefaultAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupsIsDefaultAccessAzureGroupRuleAzureAdJSON contains the JSON metadata
// for the struct [AccessGroupsIsDefaultAccessAzureGroupRuleAzureAd]
type accessGroupsIsDefaultAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupsIsDefaultAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupsIsDefaultAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupsIsDefaultAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupsIsDefaultAccessGitHubOrganizationRuleJSON contains the JSON metadata
// for the struct [AccessGroupsIsDefaultAccessGitHubOrganizationRule]
type accessGroupsIsDefaultAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIsDefaultAccessGitHubOrganizationRule) implementsAccessGroupsIsDefault() {}

type AccessGroupsIsDefaultAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                  `json:"name,required"`
	JSON accessGroupsIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupsIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON contains
// the JSON metadata for the struct
// [AccessGroupsIsDefaultAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupsIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupsIsDefaultAccessGsuiteGroupRule struct {
	Gsuite AccessGroupsIsDefaultAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupsIsDefaultAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupsIsDefaultAccessGsuiteGroupRuleJSON contains the JSON metadata for
// the struct [AccessGroupsIsDefaultAccessGsuiteGroupRule]
type accessGroupsIsDefaultAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIsDefaultAccessGsuiteGroupRule) implementsAccessGroupsIsDefault() {}

type AccessGroupsIsDefaultAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                               `json:"email,required"`
	JSON  accessGroupsIsDefaultAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupsIsDefaultAccessGsuiteGroupRuleGsuiteJSON contains the JSON metadata
// for the struct [AccessGroupsIsDefaultAccessGsuiteGroupRuleGsuite]
type accessGroupsIsDefaultAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupsIsDefaultAccessOktaGroupRule struct {
	Okta AccessGroupsIsDefaultAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupsIsDefaultAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupsIsDefaultAccessOktaGroupRuleJSON contains the JSON metadata for the
// struct [AccessGroupsIsDefaultAccessOktaGroupRule]
type accessGroupsIsDefaultAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIsDefaultAccessOktaGroupRule) implementsAccessGroupsIsDefault() {}

type AccessGroupsIsDefaultAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                           `json:"email,required"`
	JSON  accessGroupsIsDefaultAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupsIsDefaultAccessOktaGroupRuleOktaJSON contains the JSON metadata for
// the struct [AccessGroupsIsDefaultAccessOktaGroupRuleOkta]
type accessGroupsIsDefaultAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupsIsDefaultAccessSamlGroupRule struct {
	Saml AccessGroupsIsDefaultAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupsIsDefaultAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupsIsDefaultAccessSamlGroupRuleJSON contains the JSON metadata for the
// struct [AccessGroupsIsDefaultAccessSamlGroupRule]
type accessGroupsIsDefaultAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIsDefaultAccessSamlGroupRule) implementsAccessGroupsIsDefault() {}

type AccessGroupsIsDefaultAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                           `json:"attribute_value,required"`
	JSON           accessGroupsIsDefaultAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupsIsDefaultAccessSamlGroupRuleSamlJSON contains the JSON metadata for
// the struct [AccessGroupsIsDefaultAccessSamlGroupRuleSaml]
type accessGroupsIsDefaultAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupsIsDefaultAccessServiceTokenRule struct {
	ServiceToken AccessGroupsIsDefaultAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupsIsDefaultAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupsIsDefaultAccessServiceTokenRuleJSON contains the JSON metadata for
// the struct [AccessGroupsIsDefaultAccessServiceTokenRule]
type accessGroupsIsDefaultAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIsDefaultAccessServiceTokenRule) implementsAccessGroupsIsDefault() {}

type AccessGroupsIsDefaultAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                      `json:"token_id,required"`
	JSON    accessGroupsIsDefaultAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupsIsDefaultAccessServiceTokenRuleServiceTokenJSON contains the JSON
// metadata for the struct
// [AccessGroupsIsDefaultAccessServiceTokenRuleServiceToken]
type accessGroupsIsDefaultAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupsIsDefaultAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                             `json:"any_valid_service_token,required"`
	JSON                 accessGroupsIsDefaultAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupsIsDefaultAccessAnyValidServiceTokenRuleJSON contains the JSON
// metadata for the struct [AccessGroupsIsDefaultAccessAnyValidServiceTokenRule]
type accessGroupsIsDefaultAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIsDefaultAccessAnyValidServiceTokenRule) implementsAccessGroupsIsDefault() {}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupsIsDefaultAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupsIsDefaultAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupsIsDefaultAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupsIsDefaultAccessExternalEvaluationRuleJSON contains the JSON metadata
// for the struct [AccessGroupsIsDefaultAccessExternalEvaluationRule]
type accessGroupsIsDefaultAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIsDefaultAccessExternalEvaluationRule) implementsAccessGroupsIsDefault() {}

type AccessGroupsIsDefaultAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                  `json:"keys_url,required"`
	JSON    accessGroupsIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupsIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON contains
// the JSON metadata for the struct
// [AccessGroupsIsDefaultAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupsIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupsIsDefaultAccessCountryRule struct {
	Geo  AccessGroupsIsDefaultAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupsIsDefaultAccessCountryRuleJSON `json:"-"`
}

// accessGroupsIsDefaultAccessCountryRuleJSON contains the JSON metadata for the
// struct [AccessGroupsIsDefaultAccessCountryRule]
type accessGroupsIsDefaultAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIsDefaultAccessCountryRule) implementsAccessGroupsIsDefault() {}

type AccessGroupsIsDefaultAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                        `json:"country_code,required"`
	JSON        accessGroupsIsDefaultAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupsIsDefaultAccessCountryRuleGeoJSON contains the JSON metadata for the
// struct [AccessGroupsIsDefaultAccessCountryRuleGeo]
type accessGroupsIsDefaultAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupsIsDefaultAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupsIsDefaultAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupsIsDefaultAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupsIsDefaultAccessAuthenticationMethodRuleJSON contains the JSON
// metadata for the struct [AccessGroupsIsDefaultAccessAuthenticationMethodRule]
type accessGroupsIsDefaultAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIsDefaultAccessAuthenticationMethodRule) implementsAccessGroupsIsDefault() {}

type AccessGroupsIsDefaultAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                            `json:"auth_method,required"`
	JSON       accessGroupsIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupsIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON contains the
// JSON metadata for the struct
// [AccessGroupsIsDefaultAccessAuthenticationMethodRuleAuthMethod]
type accessGroupsIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupsIsDefaultAccessDevicePostureRule struct {
	DevicePosture AccessGroupsIsDefaultAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupsIsDefaultAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupsIsDefaultAccessDevicePostureRuleJSON contains the JSON metadata for
// the struct [AccessGroupsIsDefaultAccessDevicePostureRule]
type accessGroupsIsDefaultAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsIsDefaultAccessDevicePostureRule) implementsAccessGroupsIsDefault() {}

type AccessGroupsIsDefaultAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                        `json:"integration_uid,required"`
	JSON           accessGroupsIsDefaultAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupsIsDefaultAccessDevicePostureRuleDevicePostureJSON contains the JSON
// metadata for the struct
// [AccessGroupsIsDefaultAccessDevicePostureRuleDevicePosture]
type accessGroupsIsDefaultAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupsIsDefaultAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific email.
//
// Union satisfied by [AccessGroupsRequireAccessEmailRule],
// [AccessGroupsRequireAccessEmailListRule], [AccessGroupsRequireAccessDomainRule],
// [AccessGroupsRequireAccessEveryoneRule], [AccessGroupsRequireAccessIPRule],
// [AccessGroupsRequireAccessIPListRule],
// [AccessGroupsRequireAccessCertificateRule],
// [AccessGroupsRequireAccessAccessGroupRule],
// [AccessGroupsRequireAccessAzureGroupRule],
// [AccessGroupsRequireAccessGitHubOrganizationRule],
// [AccessGroupsRequireAccessGsuiteGroupRule],
// [AccessGroupsRequireAccessOktaGroupRule],
// [AccessGroupsRequireAccessSamlGroupRule],
// [AccessGroupsRequireAccessServiceTokenRule],
// [AccessGroupsRequireAccessAnyValidServiceTokenRule],
// [AccessGroupsRequireAccessExternalEvaluationRule],
// [AccessGroupsRequireAccessCountryRule],
// [AccessGroupsRequireAccessAuthenticationMethodRule] or
// [AccessGroupsRequireAccessDevicePostureRule].
type AccessGroupsRequire interface {
	implementsAccessGroupsRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccessGroupsRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type AccessGroupsRequireAccessEmailRule struct {
	Email AccessGroupsRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  accessGroupsRequireAccessEmailRuleJSON  `json:"-"`
}

// accessGroupsRequireAccessEmailRuleJSON contains the JSON metadata for the struct
// [AccessGroupsRequireAccessEmailRule]
type accessGroupsRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsRequireAccessEmailRule) implementsAccessGroupsRequire() {}

type AccessGroupsRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                      `json:"email,required" format:"email"`
	JSON  accessGroupsRequireAccessEmailRuleEmailJSON `json:"-"`
}

// accessGroupsRequireAccessEmailRuleEmailJSON contains the JSON metadata for the
// struct [AccessGroupsRequireAccessEmailRuleEmail]
type accessGroupsRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an email address from a list.
type AccessGroupsRequireAccessEmailListRule struct {
	EmailList AccessGroupsRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      accessGroupsRequireAccessEmailListRuleJSON      `json:"-"`
}

// accessGroupsRequireAccessEmailListRuleJSON contains the JSON metadata for the
// struct [AccessGroupsRequireAccessEmailListRule]
type accessGroupsRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsRequireAccessEmailListRule) implementsAccessGroupsRequire() {}

type AccessGroupsRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                              `json:"id,required"`
	JSON accessGroupsRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// accessGroupsRequireAccessEmailListRuleEmailListJSON contains the JSON metadata
// for the struct [AccessGroupsRequireAccessEmailListRuleEmailList]
type accessGroupsRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Match an entire email domain.
type AccessGroupsRequireAccessDomainRule struct {
	EmailDomain AccessGroupsRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        accessGroupsRequireAccessDomainRuleJSON        `json:"-"`
}

// accessGroupsRequireAccessDomainRuleJSON contains the JSON metadata for the
// struct [AccessGroupsRequireAccessDomainRule]
type accessGroupsRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsRequireAccessDomainRule) implementsAccessGroupsRequire() {}

type AccessGroupsRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                             `json:"domain,required"`
	JSON   accessGroupsRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// accessGroupsRequireAccessDomainRuleEmailDomainJSON contains the JSON metadata
// for the struct [AccessGroupsRequireAccessDomainRuleEmailDomain]
type accessGroupsRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches everyone.
type AccessGroupsRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                               `json:"everyone,required"`
	JSON     accessGroupsRequireAccessEveryoneRuleJSON `json:"-"`
}

// accessGroupsRequireAccessEveryoneRuleJSON contains the JSON metadata for the
// struct [AccessGroupsRequireAccessEveryoneRule]
type accessGroupsRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsRequireAccessEveryoneRule) implementsAccessGroupsRequire() {}

// Matches an IP address block.
type AccessGroupsRequireAccessIPRule struct {
	IP   AccessGroupsRequireAccessIPRuleIP   `json:"ip,required"`
	JSON accessGroupsRequireAccessIPRuleJSON `json:"-"`
}

// accessGroupsRequireAccessIPRuleJSON contains the JSON metadata for the struct
// [AccessGroupsRequireAccessIPRule]
type accessGroupsRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsRequireAccessIPRule) implementsAccessGroupsRequire() {}

type AccessGroupsRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                `json:"ip,required"`
	JSON accessGroupsRequireAccessIPRuleIPJSON `json:"-"`
}

// accessGroupsRequireAccessIPRuleIPJSON contains the JSON metadata for the struct
// [AccessGroupsRequireAccessIPRuleIP]
type accessGroupsRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an IP address from a list.
type AccessGroupsRequireAccessIPListRule struct {
	IPList AccessGroupsRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   accessGroupsRequireAccessIPListRuleJSON   `json:"-"`
}

// accessGroupsRequireAccessIPListRuleJSON contains the JSON metadata for the
// struct [AccessGroupsRequireAccessIPListRule]
type accessGroupsRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsRequireAccessIPListRule) implementsAccessGroupsRequire() {}

type AccessGroupsRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                        `json:"id,required"`
	JSON accessGroupsRequireAccessIPListRuleIPListJSON `json:"-"`
}

// accessGroupsRequireAccessIPListRuleIPListJSON contains the JSON metadata for the
// struct [AccessGroupsRequireAccessIPListRuleIPList]
type accessGroupsRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid client certificate.
type AccessGroupsRequireAccessCertificateRule struct {
	Certificate interface{}                                  `json:"certificate,required"`
	JSON        accessGroupsRequireAccessCertificateRuleJSON `json:"-"`
}

// accessGroupsRequireAccessCertificateRuleJSON contains the JSON metadata for the
// struct [AccessGroupsRequireAccessCertificateRule]
type accessGroupsRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsRequireAccessCertificateRule) implementsAccessGroupsRequire() {}

// Matches an Access group.
type AccessGroupsRequireAccessAccessGroupRule struct {
	Group AccessGroupsRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  accessGroupsRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// accessGroupsRequireAccessAccessGroupRuleJSON contains the JSON metadata for the
// struct [AccessGroupsRequireAccessAccessGroupRule]
type accessGroupsRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsRequireAccessAccessGroupRule) implementsAccessGroupsRequire() {}

type AccessGroupsRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                            `json:"id,required"`
	JSON accessGroupsRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// accessGroupsRequireAccessAccessGroupRuleGroupJSON contains the JSON metadata for
// the struct [AccessGroupsRequireAccessAccessGroupRuleGroup]
type accessGroupsRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Azure group. Requires an Azure identity provider.
type AccessGroupsRequireAccessAzureGroupRule struct {
	AzureAd AccessGroupsRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    accessGroupsRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// accessGroupsRequireAccessAzureGroupRuleJSON contains the JSON metadata for the
// struct [AccessGroupsRequireAccessAzureGroupRule]
type accessGroupsRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsRequireAccessAzureGroupRule) implementsAccessGroupsRequire() {}

type AccessGroupsRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                             `json:"connection_id,required"`
	JSON         accessGroupsRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// accessGroupsRequireAccessAzureGroupRuleAzureAdJSON contains the JSON metadata
// for the struct [AccessGroupsRequireAccessAzureGroupRuleAzureAd]
type accessGroupsRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a Github organization. Requires a Github identity provider.
type AccessGroupsRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization AccessGroupsRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               accessGroupsRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// accessGroupsRequireAccessGitHubOrganizationRuleJSON contains the JSON metadata
// for the struct [AccessGroupsRequireAccessGitHubOrganizationRule]
type accessGroupsRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsRequireAccessGitHubOrganizationRule) implementsAccessGroupsRequire() {}

type AccessGroupsRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                `json:"name,required"`
	JSON accessGroupsRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// accessGroupsRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON contains
// the JSON metadata for the struct
// [AccessGroupsRequireAccessGitHubOrganizationRuleGitHubOrganization]
type accessGroupsRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type AccessGroupsRequireAccessGsuiteGroupRule struct {
	Gsuite AccessGroupsRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   accessGroupsRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// accessGroupsRequireAccessGsuiteGroupRuleJSON contains the JSON metadata for the
// struct [AccessGroupsRequireAccessGsuiteGroupRule]
type accessGroupsRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsRequireAccessGsuiteGroupRule) implementsAccessGroupsRequire() {}

type AccessGroupsRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                             `json:"email,required"`
	JSON  accessGroupsRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// accessGroupsRequireAccessGsuiteGroupRuleGsuiteJSON contains the JSON metadata
// for the struct [AccessGroupsRequireAccessGsuiteGroupRuleGsuite]
type accessGroupsRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches an Okta group. Requires an Okta identity provider.
type AccessGroupsRequireAccessOktaGroupRule struct {
	Okta AccessGroupsRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON accessGroupsRequireAccessOktaGroupRuleJSON `json:"-"`
}

// accessGroupsRequireAccessOktaGroupRuleJSON contains the JSON metadata for the
// struct [AccessGroupsRequireAccessOktaGroupRule]
type accessGroupsRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsRequireAccessOktaGroupRule) implementsAccessGroupsRequire() {}

type AccessGroupsRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                         `json:"email,required"`
	JSON  accessGroupsRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// accessGroupsRequireAccessOktaGroupRuleOktaJSON contains the JSON metadata for
// the struct [AccessGroupsRequireAccessOktaGroupRuleOkta]
type accessGroupsRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a SAML group. Requires a SAML identity provider.
type AccessGroupsRequireAccessSamlGroupRule struct {
	Saml AccessGroupsRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON accessGroupsRequireAccessSamlGroupRuleJSON `json:"-"`
}

// accessGroupsRequireAccessSamlGroupRuleJSON contains the JSON metadata for the
// struct [AccessGroupsRequireAccessSamlGroupRule]
type accessGroupsRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsRequireAccessSamlGroupRule) implementsAccessGroupsRequire() {}

type AccessGroupsRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                         `json:"attribute_value,required"`
	JSON           accessGroupsRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// accessGroupsRequireAccessSamlGroupRuleSamlJSON contains the JSON metadata for
// the struct [AccessGroupsRequireAccessSamlGroupRuleSaml]
type accessGroupsRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific Access Service Token
type AccessGroupsRequireAccessServiceTokenRule struct {
	ServiceToken AccessGroupsRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         accessGroupsRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// accessGroupsRequireAccessServiceTokenRuleJSON contains the JSON metadata for the
// struct [AccessGroupsRequireAccessServiceTokenRule]
type accessGroupsRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsRequireAccessServiceTokenRule) implementsAccessGroupsRequire() {}

type AccessGroupsRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                    `json:"token_id,required"`
	JSON    accessGroupsRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// accessGroupsRequireAccessServiceTokenRuleServiceTokenJSON contains the JSON
// metadata for the struct [AccessGroupsRequireAccessServiceTokenRuleServiceToken]
type accessGroupsRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches any valid Access Service Token
type AccessGroupsRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                           `json:"any_valid_service_token,required"`
	JSON                 accessGroupsRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// accessGroupsRequireAccessAnyValidServiceTokenRuleJSON contains the JSON metadata
// for the struct [AccessGroupsRequireAccessAnyValidServiceTokenRule]
type accessGroupsRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsRequireAccessAnyValidServiceTokenRule) implementsAccessGroupsRequire() {}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupsRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation AccessGroupsRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               accessGroupsRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// accessGroupsRequireAccessExternalEvaluationRuleJSON contains the JSON metadata
// for the struct [AccessGroupsRequireAccessExternalEvaluationRule]
type accessGroupsRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsRequireAccessExternalEvaluationRule) implementsAccessGroupsRequire() {}

type AccessGroupsRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                `json:"keys_url,required"`
	JSON    accessGroupsRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// accessGroupsRequireAccessExternalEvaluationRuleExternalEvaluationJSON contains
// the JSON metadata for the struct
// [AccessGroupsRequireAccessExternalEvaluationRuleExternalEvaluation]
type accessGroupsRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Matches a specific country
type AccessGroupsRequireAccessCountryRule struct {
	Geo  AccessGroupsRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON accessGroupsRequireAccessCountryRuleJSON `json:"-"`
}

// accessGroupsRequireAccessCountryRuleJSON contains the JSON metadata for the
// struct [AccessGroupsRequireAccessCountryRule]
type accessGroupsRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsRequireAccessCountryRule) implementsAccessGroupsRequire() {}

type AccessGroupsRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                      `json:"country_code,required"`
	JSON        accessGroupsRequireAccessCountryRuleGeoJSON `json:"-"`
}

// accessGroupsRequireAccessCountryRuleGeoJSON contains the JSON metadata for the
// struct [AccessGroupsRequireAccessCountryRuleGeo]
type accessGroupsRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforce different MFA options
type AccessGroupsRequireAccessAuthenticationMethodRule struct {
	AuthMethod AccessGroupsRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       accessGroupsRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// accessGroupsRequireAccessAuthenticationMethodRuleJSON contains the JSON metadata
// for the struct [AccessGroupsRequireAccessAuthenticationMethodRule]
type accessGroupsRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsRequireAccessAuthenticationMethodRule) implementsAccessGroupsRequire() {}

type AccessGroupsRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                          `json:"auth_method,required"`
	JSON       accessGroupsRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// accessGroupsRequireAccessAuthenticationMethodRuleAuthMethodJSON contains the
// JSON metadata for the struct
// [AccessGroupsRequireAccessAuthenticationMethodRuleAuthMethod]
type accessGroupsRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enforces a device posture rule has run successfully
type AccessGroupsRequireAccessDevicePostureRule struct {
	DevicePosture AccessGroupsRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          accessGroupsRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// accessGroupsRequireAccessDevicePostureRuleJSON contains the JSON metadata for
// the struct [AccessGroupsRequireAccessDevicePostureRule]
type accessGroupsRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccessGroupsRequireAccessDevicePostureRule) implementsAccessGroupsRequire() {}

type AccessGroupsRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                      `json:"integration_uid,required"`
	JSON           accessGroupsRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// accessGroupsRequireAccessDevicePostureRuleDevicePostureJSON contains the JSON
// metadata for the struct
// [AccessGroupsRequireAccessDevicePostureRuleDevicePosture]
type accessGroupsRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessGroupsRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessGroupDeleteResponse struct {
	// UUID
	ID   string                                 `json:"id"`
	JSON zeroTrustAccessGroupDeleteResponseJSON `json:"-"`
}

// zeroTrustAccessGroupDeleteResponseJSON contains the JSON metadata for the struct
// [ZeroTrustAccessGroupDeleteResponse]
type zeroTrustAccessGroupDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessGroupNewParams struct {
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]ZeroTrustAccessGroupNewParamsInclude] `json:"include,required"`
	// The name of the Access group.
	Name param.Field[string] `json:"name,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]ZeroTrustAccessGroupNewParamsExclude] `json:"exclude"`
	// Whether this is the default group
	IsDefault param.Field[bool] `json:"is_default"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]ZeroTrustAccessGroupNewParamsRequire] `json:"require"`
}

func (r ZeroTrustAccessGroupNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [ZeroTrustAccessGroupNewParamsIncludeAccessEmailRule],
// [ZeroTrustAccessGroupNewParamsIncludeAccessEmailListRule],
// [ZeroTrustAccessGroupNewParamsIncludeAccessDomainRule],
// [ZeroTrustAccessGroupNewParamsIncludeAccessEveryoneRule],
// [ZeroTrustAccessGroupNewParamsIncludeAccessIPRule],
// [ZeroTrustAccessGroupNewParamsIncludeAccessIPListRule],
// [ZeroTrustAccessGroupNewParamsIncludeAccessCertificateRule],
// [ZeroTrustAccessGroupNewParamsIncludeAccessAccessGroupRule],
// [ZeroTrustAccessGroupNewParamsIncludeAccessAzureGroupRule],
// [ZeroTrustAccessGroupNewParamsIncludeAccessGitHubOrganizationRule],
// [ZeroTrustAccessGroupNewParamsIncludeAccessGsuiteGroupRule],
// [ZeroTrustAccessGroupNewParamsIncludeAccessOktaGroupRule],
// [ZeroTrustAccessGroupNewParamsIncludeAccessSamlGroupRule],
// [ZeroTrustAccessGroupNewParamsIncludeAccessServiceTokenRule],
// [ZeroTrustAccessGroupNewParamsIncludeAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessGroupNewParamsIncludeAccessExternalEvaluationRule],
// [ZeroTrustAccessGroupNewParamsIncludeAccessCountryRule],
// [ZeroTrustAccessGroupNewParamsIncludeAccessAuthenticationMethodRule],
// [ZeroTrustAccessGroupNewParamsIncludeAccessDevicePostureRule].
type ZeroTrustAccessGroupNewParamsInclude interface {
	implementsZeroTrustAccessGroupNewParamsInclude()
}

// Matches a specific email.
type ZeroTrustAccessGroupNewParamsIncludeAccessEmailRule struct {
	Email param.Field[ZeroTrustAccessGroupNewParamsIncludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessEmailRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

type ZeroTrustAccessGroupNewParamsIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type ZeroTrustAccessGroupNewParamsIncludeAccessEmailListRule struct {
	EmailList param.Field[ZeroTrustAccessGroupNewParamsIncludeAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessEmailListRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

type ZeroTrustAccessGroupNewParamsIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type ZeroTrustAccessGroupNewParamsIncludeAccessDomainRule struct {
	EmailDomain param.Field[ZeroTrustAccessGroupNewParamsIncludeAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessDomainRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

type ZeroTrustAccessGroupNewParamsIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type ZeroTrustAccessGroupNewParamsIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessEveryoneRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

// Matches an IP address block.
type ZeroTrustAccessGroupNewParamsIncludeAccessIPRule struct {
	IP param.Field[ZeroTrustAccessGroupNewParamsIncludeAccessIPRuleIP] `json:"ip,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessIPRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

type ZeroTrustAccessGroupNewParamsIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type ZeroTrustAccessGroupNewParamsIncludeAccessIPListRule struct {
	IPList param.Field[ZeroTrustAccessGroupNewParamsIncludeAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessIPListRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

type ZeroTrustAccessGroupNewParamsIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type ZeroTrustAccessGroupNewParamsIncludeAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessCertificateRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

// Matches an Access group.
type ZeroTrustAccessGroupNewParamsIncludeAccessAccessGroupRule struct {
	Group param.Field[ZeroTrustAccessGroupNewParamsIncludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessAccessGroupRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

type ZeroTrustAccessGroupNewParamsIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessGroupNewParamsIncludeAccessAzureGroupRule struct {
	AzureAd param.Field[ZeroTrustAccessGroupNewParamsIncludeAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessAzureGroupRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

type ZeroTrustAccessGroupNewParamsIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessGroupNewParamsIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[ZeroTrustAccessGroupNewParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

type ZeroTrustAccessGroupNewParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessGroupNewParamsIncludeAccessGsuiteGroupRule struct {
	Gsuite param.Field[ZeroTrustAccessGroupNewParamsIncludeAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessGsuiteGroupRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

type ZeroTrustAccessGroupNewParamsIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessGroupNewParamsIncludeAccessOktaGroupRule struct {
	Okta param.Field[ZeroTrustAccessGroupNewParamsIncludeAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessOktaGroupRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

type ZeroTrustAccessGroupNewParamsIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessGroupNewParamsIncludeAccessSamlGroupRule struct {
	Saml param.Field[ZeroTrustAccessGroupNewParamsIncludeAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessSamlGroupRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

type ZeroTrustAccessGroupNewParamsIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type ZeroTrustAccessGroupNewParamsIncludeAccessServiceTokenRule struct {
	ServiceToken param.Field[ZeroTrustAccessGroupNewParamsIncludeAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessServiceTokenRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

type ZeroTrustAccessGroupNewParamsIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type ZeroTrustAccessGroupNewParamsIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessGroupNewParamsIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[ZeroTrustAccessGroupNewParamsIncludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessExternalEvaluationRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

type ZeroTrustAccessGroupNewParamsIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type ZeroTrustAccessGroupNewParamsIncludeAccessCountryRule struct {
	Geo param.Field[ZeroTrustAccessGroupNewParamsIncludeAccessCountryRuleGeo] `json:"geo,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessCountryRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

type ZeroTrustAccessGroupNewParamsIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type ZeroTrustAccessGroupNewParamsIncludeAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[ZeroTrustAccessGroupNewParamsIncludeAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

type ZeroTrustAccessGroupNewParamsIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessGroupNewParamsIncludeAccessDevicePostureRule struct {
	DevicePosture param.Field[ZeroTrustAccessGroupNewParamsIncludeAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessDevicePostureRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

type ZeroTrustAccessGroupNewParamsIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r ZeroTrustAccessGroupNewParamsIncludeAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [ZeroTrustAccessGroupNewParamsExcludeAccessEmailRule],
// [ZeroTrustAccessGroupNewParamsExcludeAccessEmailListRule],
// [ZeroTrustAccessGroupNewParamsExcludeAccessDomainRule],
// [ZeroTrustAccessGroupNewParamsExcludeAccessEveryoneRule],
// [ZeroTrustAccessGroupNewParamsExcludeAccessIPRule],
// [ZeroTrustAccessGroupNewParamsExcludeAccessIPListRule],
// [ZeroTrustAccessGroupNewParamsExcludeAccessCertificateRule],
// [ZeroTrustAccessGroupNewParamsExcludeAccessAccessGroupRule],
// [ZeroTrustAccessGroupNewParamsExcludeAccessAzureGroupRule],
// [ZeroTrustAccessGroupNewParamsExcludeAccessGitHubOrganizationRule],
// [ZeroTrustAccessGroupNewParamsExcludeAccessGsuiteGroupRule],
// [ZeroTrustAccessGroupNewParamsExcludeAccessOktaGroupRule],
// [ZeroTrustAccessGroupNewParamsExcludeAccessSamlGroupRule],
// [ZeroTrustAccessGroupNewParamsExcludeAccessServiceTokenRule],
// [ZeroTrustAccessGroupNewParamsExcludeAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessGroupNewParamsExcludeAccessExternalEvaluationRule],
// [ZeroTrustAccessGroupNewParamsExcludeAccessCountryRule],
// [ZeroTrustAccessGroupNewParamsExcludeAccessAuthenticationMethodRule],
// [ZeroTrustAccessGroupNewParamsExcludeAccessDevicePostureRule].
type ZeroTrustAccessGroupNewParamsExclude interface {
	implementsZeroTrustAccessGroupNewParamsExclude()
}

// Matches a specific email.
type ZeroTrustAccessGroupNewParamsExcludeAccessEmailRule struct {
	Email param.Field[ZeroTrustAccessGroupNewParamsExcludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessEmailRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

type ZeroTrustAccessGroupNewParamsExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type ZeroTrustAccessGroupNewParamsExcludeAccessEmailListRule struct {
	EmailList param.Field[ZeroTrustAccessGroupNewParamsExcludeAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessEmailListRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

type ZeroTrustAccessGroupNewParamsExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type ZeroTrustAccessGroupNewParamsExcludeAccessDomainRule struct {
	EmailDomain param.Field[ZeroTrustAccessGroupNewParamsExcludeAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessDomainRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

type ZeroTrustAccessGroupNewParamsExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type ZeroTrustAccessGroupNewParamsExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessEveryoneRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

// Matches an IP address block.
type ZeroTrustAccessGroupNewParamsExcludeAccessIPRule struct {
	IP param.Field[ZeroTrustAccessGroupNewParamsExcludeAccessIPRuleIP] `json:"ip,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessIPRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

type ZeroTrustAccessGroupNewParamsExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type ZeroTrustAccessGroupNewParamsExcludeAccessIPListRule struct {
	IPList param.Field[ZeroTrustAccessGroupNewParamsExcludeAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessIPListRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

type ZeroTrustAccessGroupNewParamsExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type ZeroTrustAccessGroupNewParamsExcludeAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessCertificateRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

// Matches an Access group.
type ZeroTrustAccessGroupNewParamsExcludeAccessAccessGroupRule struct {
	Group param.Field[ZeroTrustAccessGroupNewParamsExcludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessAccessGroupRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

type ZeroTrustAccessGroupNewParamsExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessGroupNewParamsExcludeAccessAzureGroupRule struct {
	AzureAd param.Field[ZeroTrustAccessGroupNewParamsExcludeAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessAzureGroupRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

type ZeroTrustAccessGroupNewParamsExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessGroupNewParamsExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[ZeroTrustAccessGroupNewParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

type ZeroTrustAccessGroupNewParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessGroupNewParamsExcludeAccessGsuiteGroupRule struct {
	Gsuite param.Field[ZeroTrustAccessGroupNewParamsExcludeAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessGsuiteGroupRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

type ZeroTrustAccessGroupNewParamsExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessGroupNewParamsExcludeAccessOktaGroupRule struct {
	Okta param.Field[ZeroTrustAccessGroupNewParamsExcludeAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessOktaGroupRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

type ZeroTrustAccessGroupNewParamsExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessGroupNewParamsExcludeAccessSamlGroupRule struct {
	Saml param.Field[ZeroTrustAccessGroupNewParamsExcludeAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessSamlGroupRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

type ZeroTrustAccessGroupNewParamsExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type ZeroTrustAccessGroupNewParamsExcludeAccessServiceTokenRule struct {
	ServiceToken param.Field[ZeroTrustAccessGroupNewParamsExcludeAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessServiceTokenRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

type ZeroTrustAccessGroupNewParamsExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type ZeroTrustAccessGroupNewParamsExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessGroupNewParamsExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[ZeroTrustAccessGroupNewParamsExcludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessExternalEvaluationRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

type ZeroTrustAccessGroupNewParamsExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type ZeroTrustAccessGroupNewParamsExcludeAccessCountryRule struct {
	Geo param.Field[ZeroTrustAccessGroupNewParamsExcludeAccessCountryRuleGeo] `json:"geo,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessCountryRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

type ZeroTrustAccessGroupNewParamsExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type ZeroTrustAccessGroupNewParamsExcludeAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[ZeroTrustAccessGroupNewParamsExcludeAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

type ZeroTrustAccessGroupNewParamsExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessGroupNewParamsExcludeAccessDevicePostureRule struct {
	DevicePosture param.Field[ZeroTrustAccessGroupNewParamsExcludeAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessDevicePostureRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

type ZeroTrustAccessGroupNewParamsExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r ZeroTrustAccessGroupNewParamsExcludeAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [ZeroTrustAccessGroupNewParamsRequireAccessEmailRule],
// [ZeroTrustAccessGroupNewParamsRequireAccessEmailListRule],
// [ZeroTrustAccessGroupNewParamsRequireAccessDomainRule],
// [ZeroTrustAccessGroupNewParamsRequireAccessEveryoneRule],
// [ZeroTrustAccessGroupNewParamsRequireAccessIPRule],
// [ZeroTrustAccessGroupNewParamsRequireAccessIPListRule],
// [ZeroTrustAccessGroupNewParamsRequireAccessCertificateRule],
// [ZeroTrustAccessGroupNewParamsRequireAccessAccessGroupRule],
// [ZeroTrustAccessGroupNewParamsRequireAccessAzureGroupRule],
// [ZeroTrustAccessGroupNewParamsRequireAccessGitHubOrganizationRule],
// [ZeroTrustAccessGroupNewParamsRequireAccessGsuiteGroupRule],
// [ZeroTrustAccessGroupNewParamsRequireAccessOktaGroupRule],
// [ZeroTrustAccessGroupNewParamsRequireAccessSamlGroupRule],
// [ZeroTrustAccessGroupNewParamsRequireAccessServiceTokenRule],
// [ZeroTrustAccessGroupNewParamsRequireAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessGroupNewParamsRequireAccessExternalEvaluationRule],
// [ZeroTrustAccessGroupNewParamsRequireAccessCountryRule],
// [ZeroTrustAccessGroupNewParamsRequireAccessAuthenticationMethodRule],
// [ZeroTrustAccessGroupNewParamsRequireAccessDevicePostureRule].
type ZeroTrustAccessGroupNewParamsRequire interface {
	implementsZeroTrustAccessGroupNewParamsRequire()
}

// Matches a specific email.
type ZeroTrustAccessGroupNewParamsRequireAccessEmailRule struct {
	Email param.Field[ZeroTrustAccessGroupNewParamsRequireAccessEmailRuleEmail] `json:"email,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessEmailRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

type ZeroTrustAccessGroupNewParamsRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type ZeroTrustAccessGroupNewParamsRequireAccessEmailListRule struct {
	EmailList param.Field[ZeroTrustAccessGroupNewParamsRequireAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessEmailListRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

type ZeroTrustAccessGroupNewParamsRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type ZeroTrustAccessGroupNewParamsRequireAccessDomainRule struct {
	EmailDomain param.Field[ZeroTrustAccessGroupNewParamsRequireAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessDomainRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

type ZeroTrustAccessGroupNewParamsRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type ZeroTrustAccessGroupNewParamsRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessEveryoneRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

// Matches an IP address block.
type ZeroTrustAccessGroupNewParamsRequireAccessIPRule struct {
	IP param.Field[ZeroTrustAccessGroupNewParamsRequireAccessIPRuleIP] `json:"ip,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessIPRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

type ZeroTrustAccessGroupNewParamsRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type ZeroTrustAccessGroupNewParamsRequireAccessIPListRule struct {
	IPList param.Field[ZeroTrustAccessGroupNewParamsRequireAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessIPListRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

type ZeroTrustAccessGroupNewParamsRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type ZeroTrustAccessGroupNewParamsRequireAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessCertificateRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

// Matches an Access group.
type ZeroTrustAccessGroupNewParamsRequireAccessAccessGroupRule struct {
	Group param.Field[ZeroTrustAccessGroupNewParamsRequireAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessAccessGroupRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

type ZeroTrustAccessGroupNewParamsRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessGroupNewParamsRequireAccessAzureGroupRule struct {
	AzureAd param.Field[ZeroTrustAccessGroupNewParamsRequireAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessAzureGroupRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

type ZeroTrustAccessGroupNewParamsRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessGroupNewParamsRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[ZeroTrustAccessGroupNewParamsRequireAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

type ZeroTrustAccessGroupNewParamsRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessGroupNewParamsRequireAccessGsuiteGroupRule struct {
	Gsuite param.Field[ZeroTrustAccessGroupNewParamsRequireAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessGsuiteGroupRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

type ZeroTrustAccessGroupNewParamsRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessGroupNewParamsRequireAccessOktaGroupRule struct {
	Okta param.Field[ZeroTrustAccessGroupNewParamsRequireAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessOktaGroupRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

type ZeroTrustAccessGroupNewParamsRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessGroupNewParamsRequireAccessSamlGroupRule struct {
	Saml param.Field[ZeroTrustAccessGroupNewParamsRequireAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessSamlGroupRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

type ZeroTrustAccessGroupNewParamsRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type ZeroTrustAccessGroupNewParamsRequireAccessServiceTokenRule struct {
	ServiceToken param.Field[ZeroTrustAccessGroupNewParamsRequireAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessServiceTokenRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

type ZeroTrustAccessGroupNewParamsRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type ZeroTrustAccessGroupNewParamsRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessGroupNewParamsRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[ZeroTrustAccessGroupNewParamsRequireAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessExternalEvaluationRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

type ZeroTrustAccessGroupNewParamsRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type ZeroTrustAccessGroupNewParamsRequireAccessCountryRule struct {
	Geo param.Field[ZeroTrustAccessGroupNewParamsRequireAccessCountryRuleGeo] `json:"geo,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessCountryRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

type ZeroTrustAccessGroupNewParamsRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type ZeroTrustAccessGroupNewParamsRequireAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[ZeroTrustAccessGroupNewParamsRequireAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

type ZeroTrustAccessGroupNewParamsRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessGroupNewParamsRequireAccessDevicePostureRule struct {
	DevicePosture param.Field[ZeroTrustAccessGroupNewParamsRequireAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessDevicePostureRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

type ZeroTrustAccessGroupNewParamsRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r ZeroTrustAccessGroupNewParamsRequireAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustAccessGroupNewResponseEnvelope struct {
	Errors   []ZeroTrustAccessGroupNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessGroupNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessGroups                                      `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessGroupNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessGroupNewResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessGroupNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustAccessGroupNewResponseEnvelope]
type zeroTrustAccessGroupNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessGroupNewResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zeroTrustAccessGroupNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZeroTrustAccessGroupNewResponseEnvelopeErrors]
type zeroTrustAccessGroupNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessGroupNewResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zeroTrustAccessGroupNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustAccessGroupNewResponseEnvelopeMessages]
type zeroTrustAccessGroupNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessGroupNewResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessGroupNewResponseEnvelopeSuccessTrue ZeroTrustAccessGroupNewResponseEnvelopeSuccess = true
)

type ZeroTrustAccessGroupUpdateParams struct {
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]ZeroTrustAccessGroupUpdateParamsInclude] `json:"include,required"`
	// The name of the Access group.
	Name param.Field[string] `json:"name,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]ZeroTrustAccessGroupUpdateParamsExclude] `json:"exclude"`
	// Whether this is the default group
	IsDefault param.Field[bool] `json:"is_default"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]ZeroTrustAccessGroupUpdateParamsRequire] `json:"require"`
}

func (r ZeroTrustAccessGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [ZeroTrustAccessGroupUpdateParamsIncludeAccessEmailRule],
// [ZeroTrustAccessGroupUpdateParamsIncludeAccessEmailListRule],
// [ZeroTrustAccessGroupUpdateParamsIncludeAccessDomainRule],
// [ZeroTrustAccessGroupUpdateParamsIncludeAccessEveryoneRule],
// [ZeroTrustAccessGroupUpdateParamsIncludeAccessIPRule],
// [ZeroTrustAccessGroupUpdateParamsIncludeAccessIPListRule],
// [ZeroTrustAccessGroupUpdateParamsIncludeAccessCertificateRule],
// [ZeroTrustAccessGroupUpdateParamsIncludeAccessAccessGroupRule],
// [ZeroTrustAccessGroupUpdateParamsIncludeAccessAzureGroupRule],
// [ZeroTrustAccessGroupUpdateParamsIncludeAccessGitHubOrganizationRule],
// [ZeroTrustAccessGroupUpdateParamsIncludeAccessGsuiteGroupRule],
// [ZeroTrustAccessGroupUpdateParamsIncludeAccessOktaGroupRule],
// [ZeroTrustAccessGroupUpdateParamsIncludeAccessSamlGroupRule],
// [ZeroTrustAccessGroupUpdateParamsIncludeAccessServiceTokenRule],
// [ZeroTrustAccessGroupUpdateParamsIncludeAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessGroupUpdateParamsIncludeAccessExternalEvaluationRule],
// [ZeroTrustAccessGroupUpdateParamsIncludeAccessCountryRule],
// [ZeroTrustAccessGroupUpdateParamsIncludeAccessAuthenticationMethodRule],
// [ZeroTrustAccessGroupUpdateParamsIncludeAccessDevicePostureRule].
type ZeroTrustAccessGroupUpdateParamsInclude interface {
	implementsZeroTrustAccessGroupUpdateParamsInclude()
}

// Matches a specific email.
type ZeroTrustAccessGroupUpdateParamsIncludeAccessEmailRule struct {
	Email param.Field[ZeroTrustAccessGroupUpdateParamsIncludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessEmailRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
}

type ZeroTrustAccessGroupUpdateParamsIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type ZeroTrustAccessGroupUpdateParamsIncludeAccessEmailListRule struct {
	EmailList param.Field[ZeroTrustAccessGroupUpdateParamsIncludeAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessEmailListRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
}

type ZeroTrustAccessGroupUpdateParamsIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type ZeroTrustAccessGroupUpdateParamsIncludeAccessDomainRule struct {
	EmailDomain param.Field[ZeroTrustAccessGroupUpdateParamsIncludeAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessDomainRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
}

type ZeroTrustAccessGroupUpdateParamsIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type ZeroTrustAccessGroupUpdateParamsIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessEveryoneRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
}

// Matches an IP address block.
type ZeroTrustAccessGroupUpdateParamsIncludeAccessIPRule struct {
	IP param.Field[ZeroTrustAccessGroupUpdateParamsIncludeAccessIPRuleIP] `json:"ip,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessIPRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
}

type ZeroTrustAccessGroupUpdateParamsIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type ZeroTrustAccessGroupUpdateParamsIncludeAccessIPListRule struct {
	IPList param.Field[ZeroTrustAccessGroupUpdateParamsIncludeAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessIPListRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
}

type ZeroTrustAccessGroupUpdateParamsIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type ZeroTrustAccessGroupUpdateParamsIncludeAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessCertificateRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
}

// Matches an Access group.
type ZeroTrustAccessGroupUpdateParamsIncludeAccessAccessGroupRule struct {
	Group param.Field[ZeroTrustAccessGroupUpdateParamsIncludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessAccessGroupRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
}

type ZeroTrustAccessGroupUpdateParamsIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessGroupUpdateParamsIncludeAccessAzureGroupRule struct {
	AzureAd param.Field[ZeroTrustAccessGroupUpdateParamsIncludeAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessAzureGroupRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
}

type ZeroTrustAccessGroupUpdateParamsIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessGroupUpdateParamsIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[ZeroTrustAccessGroupUpdateParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
}

type ZeroTrustAccessGroupUpdateParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessGroupUpdateParamsIncludeAccessGsuiteGroupRule struct {
	Gsuite param.Field[ZeroTrustAccessGroupUpdateParamsIncludeAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessGsuiteGroupRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
}

type ZeroTrustAccessGroupUpdateParamsIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessGroupUpdateParamsIncludeAccessOktaGroupRule struct {
	Okta param.Field[ZeroTrustAccessGroupUpdateParamsIncludeAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessOktaGroupRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
}

type ZeroTrustAccessGroupUpdateParamsIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessGroupUpdateParamsIncludeAccessSamlGroupRule struct {
	Saml param.Field[ZeroTrustAccessGroupUpdateParamsIncludeAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessSamlGroupRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
}

type ZeroTrustAccessGroupUpdateParamsIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type ZeroTrustAccessGroupUpdateParamsIncludeAccessServiceTokenRule struct {
	ServiceToken param.Field[ZeroTrustAccessGroupUpdateParamsIncludeAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessServiceTokenRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
}

type ZeroTrustAccessGroupUpdateParamsIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type ZeroTrustAccessGroupUpdateParamsIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessGroupUpdateParamsIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[ZeroTrustAccessGroupUpdateParamsIncludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessExternalEvaluationRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
}

type ZeroTrustAccessGroupUpdateParamsIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type ZeroTrustAccessGroupUpdateParamsIncludeAccessCountryRule struct {
	Geo param.Field[ZeroTrustAccessGroupUpdateParamsIncludeAccessCountryRuleGeo] `json:"geo,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessCountryRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
}

type ZeroTrustAccessGroupUpdateParamsIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type ZeroTrustAccessGroupUpdateParamsIncludeAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[ZeroTrustAccessGroupUpdateParamsIncludeAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
}

type ZeroTrustAccessGroupUpdateParamsIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessGroupUpdateParamsIncludeAccessDevicePostureRule struct {
	DevicePosture param.Field[ZeroTrustAccessGroupUpdateParamsIncludeAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessDevicePostureRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
}

type ZeroTrustAccessGroupUpdateParamsIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsIncludeAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [ZeroTrustAccessGroupUpdateParamsExcludeAccessEmailRule],
// [ZeroTrustAccessGroupUpdateParamsExcludeAccessEmailListRule],
// [ZeroTrustAccessGroupUpdateParamsExcludeAccessDomainRule],
// [ZeroTrustAccessGroupUpdateParamsExcludeAccessEveryoneRule],
// [ZeroTrustAccessGroupUpdateParamsExcludeAccessIPRule],
// [ZeroTrustAccessGroupUpdateParamsExcludeAccessIPListRule],
// [ZeroTrustAccessGroupUpdateParamsExcludeAccessCertificateRule],
// [ZeroTrustAccessGroupUpdateParamsExcludeAccessAccessGroupRule],
// [ZeroTrustAccessGroupUpdateParamsExcludeAccessAzureGroupRule],
// [ZeroTrustAccessGroupUpdateParamsExcludeAccessGitHubOrganizationRule],
// [ZeroTrustAccessGroupUpdateParamsExcludeAccessGsuiteGroupRule],
// [ZeroTrustAccessGroupUpdateParamsExcludeAccessOktaGroupRule],
// [ZeroTrustAccessGroupUpdateParamsExcludeAccessSamlGroupRule],
// [ZeroTrustAccessGroupUpdateParamsExcludeAccessServiceTokenRule],
// [ZeroTrustAccessGroupUpdateParamsExcludeAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessGroupUpdateParamsExcludeAccessExternalEvaluationRule],
// [ZeroTrustAccessGroupUpdateParamsExcludeAccessCountryRule],
// [ZeroTrustAccessGroupUpdateParamsExcludeAccessAuthenticationMethodRule],
// [ZeroTrustAccessGroupUpdateParamsExcludeAccessDevicePostureRule].
type ZeroTrustAccessGroupUpdateParamsExclude interface {
	implementsZeroTrustAccessGroupUpdateParamsExclude()
}

// Matches a specific email.
type ZeroTrustAccessGroupUpdateParamsExcludeAccessEmailRule struct {
	Email param.Field[ZeroTrustAccessGroupUpdateParamsExcludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessEmailRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
}

type ZeroTrustAccessGroupUpdateParamsExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type ZeroTrustAccessGroupUpdateParamsExcludeAccessEmailListRule struct {
	EmailList param.Field[ZeroTrustAccessGroupUpdateParamsExcludeAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessEmailListRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
}

type ZeroTrustAccessGroupUpdateParamsExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type ZeroTrustAccessGroupUpdateParamsExcludeAccessDomainRule struct {
	EmailDomain param.Field[ZeroTrustAccessGroupUpdateParamsExcludeAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessDomainRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
}

type ZeroTrustAccessGroupUpdateParamsExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type ZeroTrustAccessGroupUpdateParamsExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessEveryoneRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
}

// Matches an IP address block.
type ZeroTrustAccessGroupUpdateParamsExcludeAccessIPRule struct {
	IP param.Field[ZeroTrustAccessGroupUpdateParamsExcludeAccessIPRuleIP] `json:"ip,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessIPRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
}

type ZeroTrustAccessGroupUpdateParamsExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type ZeroTrustAccessGroupUpdateParamsExcludeAccessIPListRule struct {
	IPList param.Field[ZeroTrustAccessGroupUpdateParamsExcludeAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessIPListRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
}

type ZeroTrustAccessGroupUpdateParamsExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type ZeroTrustAccessGroupUpdateParamsExcludeAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessCertificateRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
}

// Matches an Access group.
type ZeroTrustAccessGroupUpdateParamsExcludeAccessAccessGroupRule struct {
	Group param.Field[ZeroTrustAccessGroupUpdateParamsExcludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessAccessGroupRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
}

type ZeroTrustAccessGroupUpdateParamsExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessGroupUpdateParamsExcludeAccessAzureGroupRule struct {
	AzureAd param.Field[ZeroTrustAccessGroupUpdateParamsExcludeAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessAzureGroupRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
}

type ZeroTrustAccessGroupUpdateParamsExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessGroupUpdateParamsExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[ZeroTrustAccessGroupUpdateParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
}

type ZeroTrustAccessGroupUpdateParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessGroupUpdateParamsExcludeAccessGsuiteGroupRule struct {
	Gsuite param.Field[ZeroTrustAccessGroupUpdateParamsExcludeAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessGsuiteGroupRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
}

type ZeroTrustAccessGroupUpdateParamsExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessGroupUpdateParamsExcludeAccessOktaGroupRule struct {
	Okta param.Field[ZeroTrustAccessGroupUpdateParamsExcludeAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessOktaGroupRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
}

type ZeroTrustAccessGroupUpdateParamsExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessGroupUpdateParamsExcludeAccessSamlGroupRule struct {
	Saml param.Field[ZeroTrustAccessGroupUpdateParamsExcludeAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessSamlGroupRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
}

type ZeroTrustAccessGroupUpdateParamsExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type ZeroTrustAccessGroupUpdateParamsExcludeAccessServiceTokenRule struct {
	ServiceToken param.Field[ZeroTrustAccessGroupUpdateParamsExcludeAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessServiceTokenRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
}

type ZeroTrustAccessGroupUpdateParamsExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type ZeroTrustAccessGroupUpdateParamsExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessGroupUpdateParamsExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[ZeroTrustAccessGroupUpdateParamsExcludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessExternalEvaluationRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
}

type ZeroTrustAccessGroupUpdateParamsExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type ZeroTrustAccessGroupUpdateParamsExcludeAccessCountryRule struct {
	Geo param.Field[ZeroTrustAccessGroupUpdateParamsExcludeAccessCountryRuleGeo] `json:"geo,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessCountryRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
}

type ZeroTrustAccessGroupUpdateParamsExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type ZeroTrustAccessGroupUpdateParamsExcludeAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[ZeroTrustAccessGroupUpdateParamsExcludeAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
}

type ZeroTrustAccessGroupUpdateParamsExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessGroupUpdateParamsExcludeAccessDevicePostureRule struct {
	DevicePosture param.Field[ZeroTrustAccessGroupUpdateParamsExcludeAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessDevicePostureRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
}

type ZeroTrustAccessGroupUpdateParamsExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsExcludeAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [ZeroTrustAccessGroupUpdateParamsRequireAccessEmailRule],
// [ZeroTrustAccessGroupUpdateParamsRequireAccessEmailListRule],
// [ZeroTrustAccessGroupUpdateParamsRequireAccessDomainRule],
// [ZeroTrustAccessGroupUpdateParamsRequireAccessEveryoneRule],
// [ZeroTrustAccessGroupUpdateParamsRequireAccessIPRule],
// [ZeroTrustAccessGroupUpdateParamsRequireAccessIPListRule],
// [ZeroTrustAccessGroupUpdateParamsRequireAccessCertificateRule],
// [ZeroTrustAccessGroupUpdateParamsRequireAccessAccessGroupRule],
// [ZeroTrustAccessGroupUpdateParamsRequireAccessAzureGroupRule],
// [ZeroTrustAccessGroupUpdateParamsRequireAccessGitHubOrganizationRule],
// [ZeroTrustAccessGroupUpdateParamsRequireAccessGsuiteGroupRule],
// [ZeroTrustAccessGroupUpdateParamsRequireAccessOktaGroupRule],
// [ZeroTrustAccessGroupUpdateParamsRequireAccessSamlGroupRule],
// [ZeroTrustAccessGroupUpdateParamsRequireAccessServiceTokenRule],
// [ZeroTrustAccessGroupUpdateParamsRequireAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessGroupUpdateParamsRequireAccessExternalEvaluationRule],
// [ZeroTrustAccessGroupUpdateParamsRequireAccessCountryRule],
// [ZeroTrustAccessGroupUpdateParamsRequireAccessAuthenticationMethodRule],
// [ZeroTrustAccessGroupUpdateParamsRequireAccessDevicePostureRule].
type ZeroTrustAccessGroupUpdateParamsRequire interface {
	implementsZeroTrustAccessGroupUpdateParamsRequire()
}

// Matches a specific email.
type ZeroTrustAccessGroupUpdateParamsRequireAccessEmailRule struct {
	Email param.Field[ZeroTrustAccessGroupUpdateParamsRequireAccessEmailRuleEmail] `json:"email,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessEmailRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
}

type ZeroTrustAccessGroupUpdateParamsRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email param.Field[string] `json:"email,required" format:"email"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessEmailRuleEmail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an email address from a list.
type ZeroTrustAccessGroupUpdateParamsRequireAccessEmailListRule struct {
	EmailList param.Field[ZeroTrustAccessGroupUpdateParamsRequireAccessEmailListRuleEmailList] `json:"email_list,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessEmailListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessEmailListRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
}

type ZeroTrustAccessGroupUpdateParamsRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessEmailListRuleEmailList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Match an entire email domain.
type ZeroTrustAccessGroupUpdateParamsRequireAccessDomainRule struct {
	EmailDomain param.Field[ZeroTrustAccessGroupUpdateParamsRequireAccessDomainRuleEmailDomain] `json:"email_domain,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessDomainRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessDomainRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
}

type ZeroTrustAccessGroupUpdateParamsRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain param.Field[string] `json:"domain,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessDomainRuleEmailDomain) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches everyone.
type ZeroTrustAccessGroupUpdateParamsRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone param.Field[interface{}] `json:"everyone,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessEveryoneRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessEveryoneRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
}

// Matches an IP address block.
type ZeroTrustAccessGroupUpdateParamsRequireAccessIPRule struct {
	IP param.Field[ZeroTrustAccessGroupUpdateParamsRequireAccessIPRuleIP] `json:"ip,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessIPRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
}

type ZeroTrustAccessGroupUpdateParamsRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP param.Field[string] `json:"ip,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessIPRuleIP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an IP address from a list.
type ZeroTrustAccessGroupUpdateParamsRequireAccessIPListRule struct {
	IPList param.Field[ZeroTrustAccessGroupUpdateParamsRequireAccessIPListRuleIPList] `json:"ip_list,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessIPListRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessIPListRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
}

type ZeroTrustAccessGroupUpdateParamsRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessIPListRuleIPList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid client certificate.
type ZeroTrustAccessGroupUpdateParamsRequireAccessCertificateRule struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessCertificateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessCertificateRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
}

// Matches an Access group.
type ZeroTrustAccessGroupUpdateParamsRequireAccessAccessGroupRule struct {
	Group param.Field[ZeroTrustAccessGroupUpdateParamsRequireAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessAccessGroupRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
}

type ZeroTrustAccessGroupUpdateParamsRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID param.Field[string] `json:"id,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessAccessGroupRuleGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessGroupUpdateParamsRequireAccessAzureGroupRule struct {
	AzureAd param.Field[ZeroTrustAccessGroupUpdateParamsRequireAccessAzureGroupRuleAzureAd] `json:"azureAD,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessAzureGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessAzureGroupRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
}

type ZeroTrustAccessGroupUpdateParamsRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID param.Field[string] `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessAzureGroupRuleAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessGroupUpdateParamsRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization param.Field[ZeroTrustAccessGroupUpdateParamsRequireAccessGitHubOrganizationRuleGitHubOrganization] `json:"github-organization,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessGitHubOrganizationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
}

type ZeroTrustAccessGroupUpdateParamsRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The name of the organization.
	Name param.Field[string] `json:"name,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessGitHubOrganizationRuleGitHubOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessGroupUpdateParamsRequireAccessGsuiteGroupRule struct {
	Gsuite param.Field[ZeroTrustAccessGroupUpdateParamsRequireAccessGsuiteGroupRuleGsuite] `json:"gsuite,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessGsuiteGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessGsuiteGroupRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
}

type ZeroTrustAccessGroupUpdateParamsRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessGsuiteGroupRuleGsuite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessGroupUpdateParamsRequireAccessOktaGroupRule struct {
	Okta param.Field[ZeroTrustAccessGroupUpdateParamsRequireAccessOktaGroupRuleOkta] `json:"okta,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessOktaGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessOktaGroupRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
}

type ZeroTrustAccessGroupUpdateParamsRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID param.Field[string] `json:"connection_id,required"`
	// The email of the Okta group.
	Email param.Field[string] `json:"email,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessOktaGroupRuleOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessGroupUpdateParamsRequireAccessSamlGroupRule struct {
	Saml param.Field[ZeroTrustAccessGroupUpdateParamsRequireAccessSamlGroupRuleSaml] `json:"saml,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessSamlGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessSamlGroupRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
}

type ZeroTrustAccessGroupUpdateParamsRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName param.Field[string] `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue param.Field[string] `json:"attribute_value,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessSamlGroupRuleSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific Access Service Token
type ZeroTrustAccessGroupUpdateParamsRequireAccessServiceTokenRule struct {
	ServiceToken param.Field[ZeroTrustAccessGroupUpdateParamsRequireAccessServiceTokenRuleServiceToken] `json:"service_token,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessServiceTokenRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
}

type ZeroTrustAccessGroupUpdateParamsRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID param.Field[string] `json:"token_id,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessServiceTokenRuleServiceToken) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches any valid Access Service Token
type ZeroTrustAccessGroupUpdateParamsRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken param.Field[interface{}] `json:"any_valid_service_token,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessAnyValidServiceTokenRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessGroupUpdateParamsRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[ZeroTrustAccessGroupUpdateParamsRequireAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessExternalEvaluationRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
}

type ZeroTrustAccessGroupUpdateParamsRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL param.Field[string] `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL param.Field[string] `json:"keys_url,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessExternalEvaluationRuleExternalEvaluation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific country
type ZeroTrustAccessGroupUpdateParamsRequireAccessCountryRule struct {
	Geo param.Field[ZeroTrustAccessGroupUpdateParamsRequireAccessCountryRuleGeo] `json:"geo,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessCountryRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessCountryRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
}

type ZeroTrustAccessGroupUpdateParamsRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode param.Field[string] `json:"country_code,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessCountryRuleGeo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforce different MFA options
type ZeroTrustAccessGroupUpdateParamsRequireAccessAuthenticationMethodRule struct {
	AuthMethod param.Field[ZeroTrustAccessGroupUpdateParamsRequireAccessAuthenticationMethodRuleAuthMethod] `json:"auth_method,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessAuthenticationMethodRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
}

type ZeroTrustAccessGroupUpdateParamsRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod param.Field[string] `json:"auth_method,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessAuthenticationMethodRuleAuthMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessGroupUpdateParamsRequireAccessDevicePostureRule struct {
	DevicePosture param.Field[ZeroTrustAccessGroupUpdateParamsRequireAccessDevicePostureRuleDevicePosture] `json:"device_posture,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessDevicePostureRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessDevicePostureRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
}

type ZeroTrustAccessGroupUpdateParamsRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r ZeroTrustAccessGroupUpdateParamsRequireAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustAccessGroupUpdateResponseEnvelope struct {
	Errors   []ZeroTrustAccessGroupUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessGroupUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessGroups                                         `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessGroupUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessGroupUpdateResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustAccessGroupUpdateResponseEnvelope]
type zeroTrustAccessGroupUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessGroupUpdateResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zeroTrustAccessGroupUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustAccessGroupUpdateResponseEnvelopeErrors]
type zeroTrustAccessGroupUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessGroupUpdateResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustAccessGroupUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupUpdateResponseEnvelopeMessages]
type zeroTrustAccessGroupUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessGroupUpdateResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessGroupUpdateResponseEnvelopeSuccessTrue ZeroTrustAccessGroupUpdateResponseEnvelopeSuccess = true
)

type ZeroTrustAccessGroupListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type ZeroTrustAccessGroupListResponseEnvelope struct {
	Errors   []ZeroTrustAccessGroupListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessGroupListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessGroups                                     `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustAccessGroupListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustAccessGroupListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustAccessGroupListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustAccessGroupListResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustAccessGroupListResponseEnvelope]
type zeroTrustAccessGroupListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessGroupListResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zeroTrustAccessGroupListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustAccessGroupListResponseEnvelopeErrors]
type zeroTrustAccessGroupListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessGroupListResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zeroTrustAccessGroupListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustAccessGroupListResponseEnvelopeMessages]
type zeroTrustAccessGroupListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessGroupListResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessGroupListResponseEnvelopeSuccessTrue ZeroTrustAccessGroupListResponseEnvelopeSuccess = true
)

type ZeroTrustAccessGroupListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                `json:"total_count"`
	JSON       zeroTrustAccessGroupListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupListResponseEnvelopeResultInfo]
type zeroTrustAccessGroupListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessGroupDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type ZeroTrustAccessGroupDeleteResponseEnvelope struct {
	Errors   []ZeroTrustAccessGroupDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessGroupDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessGroupDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessGroupDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessGroupDeleteResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessGroupDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustAccessGroupDeleteResponseEnvelope]
type zeroTrustAccessGroupDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessGroupDeleteResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zeroTrustAccessGroupDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessGroupDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustAccessGroupDeleteResponseEnvelopeErrors]
type zeroTrustAccessGroupDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessGroupDeleteResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustAccessGroupDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessGroupDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupDeleteResponseEnvelopeMessages]
type zeroTrustAccessGroupDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessGroupDeleteResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessGroupDeleteResponseEnvelopeSuccessTrue ZeroTrustAccessGroupDeleteResponseEnvelopeSuccess = true
)

type ZeroTrustAccessGroupGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type ZeroTrustAccessGroupGetResponseEnvelope struct {
	Errors   []ZeroTrustAccessGroupGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessGroupGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessGroups                                      `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessGroupGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessGroupGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessGroupGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustAccessGroupGetResponseEnvelope]
type zeroTrustAccessGroupGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessGroupGetResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zeroTrustAccessGroupGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZeroTrustAccessGroupGetResponseEnvelopeErrors]
type zeroTrustAccessGroupGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessGroupGetResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zeroTrustAccessGroupGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustAccessGroupGetResponseEnvelopeMessages]
type zeroTrustAccessGroupGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessGroupGetResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessGroupGetResponseEnvelopeSuccessTrue ZeroTrustAccessGroupGetResponseEnvelopeSuccess = true
)
