// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
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
func (r *AccessGroupService) New(ctx context.Context, params AccessGroupNewParams, opts ...option.RequestOption) (res *AccessGroups, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessGroupNewResponseEnvelope
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
func (r *AccessGroupService) Update(ctx context.Context, uuid string, params AccessGroupUpdateParams, opts ...option.RequestOption) (res *AccessGroups, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessGroupUpdateResponseEnvelope
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
func (r *AccessGroupService) List(ctx context.Context, query AccessGroupListParams, opts ...option.RequestOption) (res *[]AccessGroups, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessGroupListResponseEnvelope
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an Access group.
func (r *AccessGroupService) Delete(ctx context.Context, uuid string, body AccessGroupDeleteParams, opts ...option.RequestOption) (res *AccessGroupDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessGroupDeleteResponseEnvelope
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Access group.
func (r *AccessGroupService) Get(ctx context.Context, uuid string, query AccessGroupGetParams, opts ...option.RequestOption) (res *AccessGroups, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessGroupGetResponseEnvelope
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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

func (r accessGroupsJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by [zero_trust.AccessGroupsExcludeAccessEmailRule],
// [zero_trust.AccessGroupsExcludeAccessEmailListRule],
// [zero_trust.AccessGroupsExcludeAccessDomainRule],
// [zero_trust.AccessGroupsExcludeAccessEveryoneRule],
// [zero_trust.AccessGroupsExcludeAccessIPRule],
// [zero_trust.AccessGroupsExcludeAccessIPListRule],
// [zero_trust.AccessGroupsExcludeAccessCertificateRule],
// [zero_trust.AccessGroupsExcludeAccessAccessGroupRule],
// [zero_trust.AccessGroupsExcludeAccessAzureGroupRule],
// [zero_trust.AccessGroupsExcludeAccessGitHubOrganizationRule],
// [zero_trust.AccessGroupsExcludeAccessGsuiteGroupRule],
// [zero_trust.AccessGroupsExcludeAccessOktaGroupRule],
// [zero_trust.AccessGroupsExcludeAccessSamlGroupRule],
// [zero_trust.AccessGroupsExcludeAccessServiceTokenRule],
// [zero_trust.AccessGroupsExcludeAccessAnyValidServiceTokenRule],
// [zero_trust.AccessGroupsExcludeAccessExternalEvaluationRule],
// [zero_trust.AccessGroupsExcludeAccessCountryRule],
// [zero_trust.AccessGroupsExcludeAccessAuthenticationMethodRule] or
// [zero_trust.AccessGroupsExcludeAccessDevicePostureRule].
type AccessGroupsExclude interface {
	implementsZeroTrustAccessGroupsExclude()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessGroupsExclude)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsExcludeAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsExcludeAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsExcludeAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsExcludeAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsExcludeAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsExcludeAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsExcludeAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsExcludeAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsExcludeAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsExcludeAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsExcludeAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsExcludeAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsExcludeAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsExcludeAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsExcludeAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsExcludeAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsExcludeAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsExcludeAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsExcludeAccessDevicePostureRule{}),
		},
	)
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

func (r accessGroupsExcludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsExcludeAccessEmailRule) implementsZeroTrustAccessGroupsExclude() {}

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

func (r accessGroupsExcludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsExcludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsExcludeAccessEmailListRule) implementsZeroTrustAccessGroupsExclude() {}

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

func (r accessGroupsExcludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsExcludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsExcludeAccessDomainRule) implementsZeroTrustAccessGroupsExclude() {}

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

func (r accessGroupsExcludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsExcludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsExcludeAccessEveryoneRule) implementsZeroTrustAccessGroupsExclude() {}

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

func (r accessGroupsExcludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsExcludeAccessIPRule) implementsZeroTrustAccessGroupsExclude() {}

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

func (r accessGroupsExcludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsExcludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsExcludeAccessIPListRule) implementsZeroTrustAccessGroupsExclude() {}

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

func (r accessGroupsExcludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsExcludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsExcludeAccessCertificateRule) implementsZeroTrustAccessGroupsExclude() {}

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

func (r accessGroupsExcludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsExcludeAccessAccessGroupRule) implementsZeroTrustAccessGroupsExclude() {}

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

func (r accessGroupsExcludeAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsExcludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsExcludeAccessAzureGroupRule) implementsZeroTrustAccessGroupsExclude() {}

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

func (r accessGroupsExcludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsExcludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsExcludeAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupsExclude() {}

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

func (r accessGroupsExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsExcludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsExcludeAccessGsuiteGroupRule) implementsZeroTrustAccessGroupsExclude() {}

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

func (r accessGroupsExcludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsExcludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsExcludeAccessOktaGroupRule) implementsZeroTrustAccessGroupsExclude() {}

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

func (r accessGroupsExcludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsExcludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsExcludeAccessSamlGroupRule) implementsZeroTrustAccessGroupsExclude() {}

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

func (r accessGroupsExcludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsExcludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsExcludeAccessServiceTokenRule) implementsZeroTrustAccessGroupsExclude() {}

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

func (r accessGroupsExcludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsExcludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsExcludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupsExclude() {}

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

func (r accessGroupsExcludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsExcludeAccessExternalEvaluationRule) implementsZeroTrustAccessGroupsExclude() {}

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

func (r accessGroupsExcludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsExcludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsExcludeAccessCountryRule) implementsZeroTrustAccessGroupsExclude() {}

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

func (r accessGroupsExcludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsExcludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsExcludeAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupsExclude() {}

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

func (r accessGroupsExcludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsExcludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsExcludeAccessDevicePostureRule) implementsZeroTrustAccessGroupsExclude() {}

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

func (r accessGroupsExcludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by [zero_trust.AccessGroupsIncludeAccessEmailRule],
// [zero_trust.AccessGroupsIncludeAccessEmailListRule],
// [zero_trust.AccessGroupsIncludeAccessDomainRule],
// [zero_trust.AccessGroupsIncludeAccessEveryoneRule],
// [zero_trust.AccessGroupsIncludeAccessIPRule],
// [zero_trust.AccessGroupsIncludeAccessIPListRule],
// [zero_trust.AccessGroupsIncludeAccessCertificateRule],
// [zero_trust.AccessGroupsIncludeAccessAccessGroupRule],
// [zero_trust.AccessGroupsIncludeAccessAzureGroupRule],
// [zero_trust.AccessGroupsIncludeAccessGitHubOrganizationRule],
// [zero_trust.AccessGroupsIncludeAccessGsuiteGroupRule],
// [zero_trust.AccessGroupsIncludeAccessOktaGroupRule],
// [zero_trust.AccessGroupsIncludeAccessSamlGroupRule],
// [zero_trust.AccessGroupsIncludeAccessServiceTokenRule],
// [zero_trust.AccessGroupsIncludeAccessAnyValidServiceTokenRule],
// [zero_trust.AccessGroupsIncludeAccessExternalEvaluationRule],
// [zero_trust.AccessGroupsIncludeAccessCountryRule],
// [zero_trust.AccessGroupsIncludeAccessAuthenticationMethodRule] or
// [zero_trust.AccessGroupsIncludeAccessDevicePostureRule].
type AccessGroupsInclude interface {
	implementsZeroTrustAccessGroupsInclude()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessGroupsInclude)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIncludeAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIncludeAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIncludeAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIncludeAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIncludeAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIncludeAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIncludeAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIncludeAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIncludeAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIncludeAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIncludeAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIncludeAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIncludeAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIncludeAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIncludeAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIncludeAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIncludeAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIncludeAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIncludeAccessDevicePostureRule{}),
		},
	)
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

func (r accessGroupsIncludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIncludeAccessEmailRule) implementsZeroTrustAccessGroupsInclude() {}

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

func (r accessGroupsIncludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsIncludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIncludeAccessEmailListRule) implementsZeroTrustAccessGroupsInclude() {}

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

func (r accessGroupsIncludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsIncludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIncludeAccessDomainRule) implementsZeroTrustAccessGroupsInclude() {}

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

func (r accessGroupsIncludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsIncludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIncludeAccessEveryoneRule) implementsZeroTrustAccessGroupsInclude() {}

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

func (r accessGroupsIncludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIncludeAccessIPRule) implementsZeroTrustAccessGroupsInclude() {}

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

func (r accessGroupsIncludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsIncludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIncludeAccessIPListRule) implementsZeroTrustAccessGroupsInclude() {}

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

func (r accessGroupsIncludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsIncludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIncludeAccessCertificateRule) implementsZeroTrustAccessGroupsInclude() {}

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

func (r accessGroupsIncludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIncludeAccessAccessGroupRule) implementsZeroTrustAccessGroupsInclude() {}

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

func (r accessGroupsIncludeAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsIncludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIncludeAccessAzureGroupRule) implementsZeroTrustAccessGroupsInclude() {}

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

func (r accessGroupsIncludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsIncludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIncludeAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupsInclude() {}

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

func (r accessGroupsIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsIncludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIncludeAccessGsuiteGroupRule) implementsZeroTrustAccessGroupsInclude() {}

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

func (r accessGroupsIncludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsIncludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIncludeAccessOktaGroupRule) implementsZeroTrustAccessGroupsInclude() {}

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

func (r accessGroupsIncludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsIncludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIncludeAccessSamlGroupRule) implementsZeroTrustAccessGroupsInclude() {}

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

func (r accessGroupsIncludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsIncludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIncludeAccessServiceTokenRule) implementsZeroTrustAccessGroupsInclude() {}

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

func (r accessGroupsIncludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsIncludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIncludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupsInclude() {}

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

func (r accessGroupsIncludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIncludeAccessExternalEvaluationRule) implementsZeroTrustAccessGroupsInclude() {}

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

func (r accessGroupsIncludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsIncludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIncludeAccessCountryRule) implementsZeroTrustAccessGroupsInclude() {}

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

func (r accessGroupsIncludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsIncludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIncludeAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupsInclude() {}

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

func (r accessGroupsIncludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsIncludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIncludeAccessDevicePostureRule) implementsZeroTrustAccessGroupsInclude() {}

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

func (r accessGroupsIncludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by [zero_trust.AccessGroupsIsDefaultAccessEmailRule],
// [zero_trust.AccessGroupsIsDefaultAccessEmailListRule],
// [zero_trust.AccessGroupsIsDefaultAccessDomainRule],
// [zero_trust.AccessGroupsIsDefaultAccessEveryoneRule],
// [zero_trust.AccessGroupsIsDefaultAccessIPRule],
// [zero_trust.AccessGroupsIsDefaultAccessIPListRule],
// [zero_trust.AccessGroupsIsDefaultAccessCertificateRule],
// [zero_trust.AccessGroupsIsDefaultAccessAccessGroupRule],
// [zero_trust.AccessGroupsIsDefaultAccessAzureGroupRule],
// [zero_trust.AccessGroupsIsDefaultAccessGitHubOrganizationRule],
// [zero_trust.AccessGroupsIsDefaultAccessGsuiteGroupRule],
// [zero_trust.AccessGroupsIsDefaultAccessOktaGroupRule],
// [zero_trust.AccessGroupsIsDefaultAccessSamlGroupRule],
// [zero_trust.AccessGroupsIsDefaultAccessServiceTokenRule],
// [zero_trust.AccessGroupsIsDefaultAccessAnyValidServiceTokenRule],
// [zero_trust.AccessGroupsIsDefaultAccessExternalEvaluationRule],
// [zero_trust.AccessGroupsIsDefaultAccessCountryRule],
// [zero_trust.AccessGroupsIsDefaultAccessAuthenticationMethodRule] or
// [zero_trust.AccessGroupsIsDefaultAccessDevicePostureRule].
type AccessGroupsIsDefault interface {
	implementsZeroTrustAccessGroupsIsDefault()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessGroupsIsDefault)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIsDefaultAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIsDefaultAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIsDefaultAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIsDefaultAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIsDefaultAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIsDefaultAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIsDefaultAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIsDefaultAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIsDefaultAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIsDefaultAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIsDefaultAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIsDefaultAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIsDefaultAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIsDefaultAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIsDefaultAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIsDefaultAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIsDefaultAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIsDefaultAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsIsDefaultAccessDevicePostureRule{}),
		},
	)
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

func (r accessGroupsIsDefaultAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIsDefaultAccessEmailRule) implementsZeroTrustAccessGroupsIsDefault() {}

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

func (r accessGroupsIsDefaultAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsIsDefaultAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIsDefaultAccessEmailListRule) implementsZeroTrustAccessGroupsIsDefault() {}

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

func (r accessGroupsIsDefaultAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsIsDefaultAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIsDefaultAccessDomainRule) implementsZeroTrustAccessGroupsIsDefault() {}

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

func (r accessGroupsIsDefaultAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsIsDefaultAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIsDefaultAccessEveryoneRule) implementsZeroTrustAccessGroupsIsDefault() {}

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

func (r accessGroupsIsDefaultAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIsDefaultAccessIPRule) implementsZeroTrustAccessGroupsIsDefault() {}

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

func (r accessGroupsIsDefaultAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsIsDefaultAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIsDefaultAccessIPListRule) implementsZeroTrustAccessGroupsIsDefault() {}

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

func (r accessGroupsIsDefaultAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsIsDefaultAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIsDefaultAccessCertificateRule) implementsZeroTrustAccessGroupsIsDefault() {}

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

func (r accessGroupsIsDefaultAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIsDefaultAccessAccessGroupRule) implementsZeroTrustAccessGroupsIsDefault() {}

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

func (r accessGroupsIsDefaultAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsIsDefaultAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIsDefaultAccessAzureGroupRule) implementsZeroTrustAccessGroupsIsDefault() {}

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

func (r accessGroupsIsDefaultAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsIsDefaultAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIsDefaultAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupsIsDefault() {
}

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

func (r accessGroupsIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsIsDefaultAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIsDefaultAccessGsuiteGroupRule) implementsZeroTrustAccessGroupsIsDefault() {}

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

func (r accessGroupsIsDefaultAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsIsDefaultAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIsDefaultAccessOktaGroupRule) implementsZeroTrustAccessGroupsIsDefault() {}

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

func (r accessGroupsIsDefaultAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsIsDefaultAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIsDefaultAccessSamlGroupRule) implementsZeroTrustAccessGroupsIsDefault() {}

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

func (r accessGroupsIsDefaultAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsIsDefaultAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIsDefaultAccessServiceTokenRule) implementsZeroTrustAccessGroupsIsDefault() {}

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

func (r accessGroupsIsDefaultAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsIsDefaultAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIsDefaultAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupsIsDefault() {
}

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

func (r accessGroupsIsDefaultAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIsDefaultAccessExternalEvaluationRule) implementsZeroTrustAccessGroupsIsDefault() {
}

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

func (r accessGroupsIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsIsDefaultAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIsDefaultAccessCountryRule) implementsZeroTrustAccessGroupsIsDefault() {}

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

func (r accessGroupsIsDefaultAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsIsDefaultAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIsDefaultAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupsIsDefault() {
}

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

func (r accessGroupsIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsIsDefaultAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsIsDefaultAccessDevicePostureRule) implementsZeroTrustAccessGroupsIsDefault() {}

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

func (r accessGroupsIsDefaultAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by [zero_trust.AccessGroupsRequireAccessEmailRule],
// [zero_trust.AccessGroupsRequireAccessEmailListRule],
// [zero_trust.AccessGroupsRequireAccessDomainRule],
// [zero_trust.AccessGroupsRequireAccessEveryoneRule],
// [zero_trust.AccessGroupsRequireAccessIPRule],
// [zero_trust.AccessGroupsRequireAccessIPListRule],
// [zero_trust.AccessGroupsRequireAccessCertificateRule],
// [zero_trust.AccessGroupsRequireAccessAccessGroupRule],
// [zero_trust.AccessGroupsRequireAccessAzureGroupRule],
// [zero_trust.AccessGroupsRequireAccessGitHubOrganizationRule],
// [zero_trust.AccessGroupsRequireAccessGsuiteGroupRule],
// [zero_trust.AccessGroupsRequireAccessOktaGroupRule],
// [zero_trust.AccessGroupsRequireAccessSamlGroupRule],
// [zero_trust.AccessGroupsRequireAccessServiceTokenRule],
// [zero_trust.AccessGroupsRequireAccessAnyValidServiceTokenRule],
// [zero_trust.AccessGroupsRequireAccessExternalEvaluationRule],
// [zero_trust.AccessGroupsRequireAccessCountryRule],
// [zero_trust.AccessGroupsRequireAccessAuthenticationMethodRule] or
// [zero_trust.AccessGroupsRequireAccessDevicePostureRule].
type AccessGroupsRequire interface {
	implementsZeroTrustAccessGroupsRequire()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessGroupsRequire)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsRequireAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsRequireAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsRequireAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsRequireAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsRequireAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsRequireAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsRequireAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsRequireAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsRequireAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsRequireAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsRequireAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsRequireAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsRequireAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsRequireAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsRequireAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsRequireAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsRequireAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsRequireAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessGroupsRequireAccessDevicePostureRule{}),
		},
	)
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

func (r accessGroupsRequireAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsRequireAccessEmailRule) implementsZeroTrustAccessGroupsRequire() {}

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

func (r accessGroupsRequireAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsRequireAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsRequireAccessEmailListRule) implementsZeroTrustAccessGroupsRequire() {}

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

func (r accessGroupsRequireAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsRequireAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsRequireAccessDomainRule) implementsZeroTrustAccessGroupsRequire() {}

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

func (r accessGroupsRequireAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsRequireAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsRequireAccessEveryoneRule) implementsZeroTrustAccessGroupsRequire() {}

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

func (r accessGroupsRequireAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsRequireAccessIPRule) implementsZeroTrustAccessGroupsRequire() {}

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

func (r accessGroupsRequireAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsRequireAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsRequireAccessIPListRule) implementsZeroTrustAccessGroupsRequire() {}

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

func (r accessGroupsRequireAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsRequireAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsRequireAccessCertificateRule) implementsZeroTrustAccessGroupsRequire() {}

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

func (r accessGroupsRequireAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsRequireAccessAccessGroupRule) implementsZeroTrustAccessGroupsRequire() {}

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

func (r accessGroupsRequireAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsRequireAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsRequireAccessAzureGroupRule) implementsZeroTrustAccessGroupsRequire() {}

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

func (r accessGroupsRequireAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsRequireAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsRequireAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupsRequire() {}

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

func (r accessGroupsRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsRequireAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsRequireAccessGsuiteGroupRule) implementsZeroTrustAccessGroupsRequire() {}

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

func (r accessGroupsRequireAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsRequireAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsRequireAccessOktaGroupRule) implementsZeroTrustAccessGroupsRequire() {}

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

func (r accessGroupsRequireAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsRequireAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsRequireAccessSamlGroupRule) implementsZeroTrustAccessGroupsRequire() {}

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

func (r accessGroupsRequireAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsRequireAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsRequireAccessServiceTokenRule) implementsZeroTrustAccessGroupsRequire() {}

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

func (r accessGroupsRequireAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsRequireAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsRequireAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupsRequire() {}

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

func (r accessGroupsRequireAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsRequireAccessExternalEvaluationRule) implementsZeroTrustAccessGroupsRequire() {}

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

func (r accessGroupsRequireAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsRequireAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsRequireAccessCountryRule) implementsZeroTrustAccessGroupsRequire() {}

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

func (r accessGroupsRequireAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsRequireAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsRequireAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupsRequire() {}

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

func (r accessGroupsRequireAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupsRequireAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r AccessGroupsRequireAccessDevicePostureRule) implementsZeroTrustAccessGroupsRequire() {}

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

func (r accessGroupsRequireAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccessGroupNewParams struct {
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessGroupNewParamsInclude] `json:"include,required"`
	// The name of the Access group.
	Name param.Field[string] `json:"name,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
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
// Satisfied by [zero_trust.AccessGroupNewParamsIncludeAccessEmailRule],
// [zero_trust.AccessGroupNewParamsIncludeAccessEmailListRule],
// [zero_trust.AccessGroupNewParamsIncludeAccessDomainRule],
// [zero_trust.AccessGroupNewParamsIncludeAccessEveryoneRule],
// [zero_trust.AccessGroupNewParamsIncludeAccessIPRule],
// [zero_trust.AccessGroupNewParamsIncludeAccessIPListRule],
// [zero_trust.AccessGroupNewParamsIncludeAccessCertificateRule],
// [zero_trust.AccessGroupNewParamsIncludeAccessAccessGroupRule],
// [zero_trust.AccessGroupNewParamsIncludeAccessAzureGroupRule],
// [zero_trust.AccessGroupNewParamsIncludeAccessGitHubOrganizationRule],
// [zero_trust.AccessGroupNewParamsIncludeAccessGsuiteGroupRule],
// [zero_trust.AccessGroupNewParamsIncludeAccessOktaGroupRule],
// [zero_trust.AccessGroupNewParamsIncludeAccessSamlGroupRule],
// [zero_trust.AccessGroupNewParamsIncludeAccessServiceTokenRule],
// [zero_trust.AccessGroupNewParamsIncludeAccessAnyValidServiceTokenRule],
// [zero_trust.AccessGroupNewParamsIncludeAccessExternalEvaluationRule],
// [zero_trust.AccessGroupNewParamsIncludeAccessCountryRule],
// [zero_trust.AccessGroupNewParamsIncludeAccessAuthenticationMethodRule],
// [zero_trust.AccessGroupNewParamsIncludeAccessDevicePostureRule].
type AccessGroupNewParamsInclude interface {
	implementsZeroTrustAccessGroupNewParamsInclude()
}

// Matches a specific email.
type AccessGroupNewParamsIncludeAccessEmailRule struct {
	Email param.Field[AccessGroupNewParamsIncludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessGroupNewParamsIncludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsIncludeAccessEmailRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

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

func (r AccessGroupNewParamsIncludeAccessEmailListRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

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

func (r AccessGroupNewParamsIncludeAccessDomainRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

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

func (r AccessGroupNewParamsIncludeAccessEveryoneRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

// Matches an IP address block.
type AccessGroupNewParamsIncludeAccessIPRule struct {
	IP param.Field[AccessGroupNewParamsIncludeAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessGroupNewParamsIncludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsIncludeAccessIPRule) implementsZeroTrustAccessGroupNewParamsInclude() {}

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

func (r AccessGroupNewParamsIncludeAccessIPListRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

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

func (r AccessGroupNewParamsIncludeAccessCertificateRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

// Matches an Access group.
type AccessGroupNewParamsIncludeAccessAccessGroupRule struct {
	Group param.Field[AccessGroupNewParamsIncludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessGroupNewParamsIncludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsIncludeAccessAccessGroupRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

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

func (r AccessGroupNewParamsIncludeAccessAzureGroupRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

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

func (r AccessGroupNewParamsIncludeAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupNewParamsInclude() {
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

func (r AccessGroupNewParamsIncludeAccessGsuiteGroupRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

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

func (r AccessGroupNewParamsIncludeAccessOktaGroupRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

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

func (r AccessGroupNewParamsIncludeAccessSamlGroupRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

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

func (r AccessGroupNewParamsIncludeAccessServiceTokenRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

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

func (r AccessGroupNewParamsIncludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupNewParamsIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessGroupNewParamsIncludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessGroupNewParamsIncludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsIncludeAccessExternalEvaluationRule) implementsZeroTrustAccessGroupNewParamsInclude() {
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

func (r AccessGroupNewParamsIncludeAccessCountryRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

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

func (r AccessGroupNewParamsIncludeAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupNewParamsInclude() {
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

func (r AccessGroupNewParamsIncludeAccessDevicePostureRule) implementsZeroTrustAccessGroupNewParamsInclude() {
}

type AccessGroupNewParamsIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccessGroupNewParamsIncludeAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [zero_trust.AccessGroupNewParamsExcludeAccessEmailRule],
// [zero_trust.AccessGroupNewParamsExcludeAccessEmailListRule],
// [zero_trust.AccessGroupNewParamsExcludeAccessDomainRule],
// [zero_trust.AccessGroupNewParamsExcludeAccessEveryoneRule],
// [zero_trust.AccessGroupNewParamsExcludeAccessIPRule],
// [zero_trust.AccessGroupNewParamsExcludeAccessIPListRule],
// [zero_trust.AccessGroupNewParamsExcludeAccessCertificateRule],
// [zero_trust.AccessGroupNewParamsExcludeAccessAccessGroupRule],
// [zero_trust.AccessGroupNewParamsExcludeAccessAzureGroupRule],
// [zero_trust.AccessGroupNewParamsExcludeAccessGitHubOrganizationRule],
// [zero_trust.AccessGroupNewParamsExcludeAccessGsuiteGroupRule],
// [zero_trust.AccessGroupNewParamsExcludeAccessOktaGroupRule],
// [zero_trust.AccessGroupNewParamsExcludeAccessSamlGroupRule],
// [zero_trust.AccessGroupNewParamsExcludeAccessServiceTokenRule],
// [zero_trust.AccessGroupNewParamsExcludeAccessAnyValidServiceTokenRule],
// [zero_trust.AccessGroupNewParamsExcludeAccessExternalEvaluationRule],
// [zero_trust.AccessGroupNewParamsExcludeAccessCountryRule],
// [zero_trust.AccessGroupNewParamsExcludeAccessAuthenticationMethodRule],
// [zero_trust.AccessGroupNewParamsExcludeAccessDevicePostureRule].
type AccessGroupNewParamsExclude interface {
	implementsZeroTrustAccessGroupNewParamsExclude()
}

// Matches a specific email.
type AccessGroupNewParamsExcludeAccessEmailRule struct {
	Email param.Field[AccessGroupNewParamsExcludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessGroupNewParamsExcludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsExcludeAccessEmailRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

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

func (r AccessGroupNewParamsExcludeAccessEmailListRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

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

func (r AccessGroupNewParamsExcludeAccessDomainRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

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

func (r AccessGroupNewParamsExcludeAccessEveryoneRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

// Matches an IP address block.
type AccessGroupNewParamsExcludeAccessIPRule struct {
	IP param.Field[AccessGroupNewParamsExcludeAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessGroupNewParamsExcludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsExcludeAccessIPRule) implementsZeroTrustAccessGroupNewParamsExclude() {}

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

func (r AccessGroupNewParamsExcludeAccessIPListRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

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

func (r AccessGroupNewParamsExcludeAccessCertificateRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

// Matches an Access group.
type AccessGroupNewParamsExcludeAccessAccessGroupRule struct {
	Group param.Field[AccessGroupNewParamsExcludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessGroupNewParamsExcludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsExcludeAccessAccessGroupRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

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

func (r AccessGroupNewParamsExcludeAccessAzureGroupRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

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

func (r AccessGroupNewParamsExcludeAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupNewParamsExclude() {
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

func (r AccessGroupNewParamsExcludeAccessGsuiteGroupRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

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

func (r AccessGroupNewParamsExcludeAccessOktaGroupRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

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

func (r AccessGroupNewParamsExcludeAccessSamlGroupRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

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

func (r AccessGroupNewParamsExcludeAccessServiceTokenRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

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

func (r AccessGroupNewParamsExcludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupNewParamsExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessGroupNewParamsExcludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessGroupNewParamsExcludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsExcludeAccessExternalEvaluationRule) implementsZeroTrustAccessGroupNewParamsExclude() {
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

func (r AccessGroupNewParamsExcludeAccessCountryRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

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

func (r AccessGroupNewParamsExcludeAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupNewParamsExclude() {
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

func (r AccessGroupNewParamsExcludeAccessDevicePostureRule) implementsZeroTrustAccessGroupNewParamsExclude() {
}

type AccessGroupNewParamsExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid param.Field[string] `json:"integration_uid,required"`
}

func (r AccessGroupNewParamsExcludeAccessDevicePostureRuleDevicePosture) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matches a specific email.
//
// Satisfied by [zero_trust.AccessGroupNewParamsRequireAccessEmailRule],
// [zero_trust.AccessGroupNewParamsRequireAccessEmailListRule],
// [zero_trust.AccessGroupNewParamsRequireAccessDomainRule],
// [zero_trust.AccessGroupNewParamsRequireAccessEveryoneRule],
// [zero_trust.AccessGroupNewParamsRequireAccessIPRule],
// [zero_trust.AccessGroupNewParamsRequireAccessIPListRule],
// [zero_trust.AccessGroupNewParamsRequireAccessCertificateRule],
// [zero_trust.AccessGroupNewParamsRequireAccessAccessGroupRule],
// [zero_trust.AccessGroupNewParamsRequireAccessAzureGroupRule],
// [zero_trust.AccessGroupNewParamsRequireAccessGitHubOrganizationRule],
// [zero_trust.AccessGroupNewParamsRequireAccessGsuiteGroupRule],
// [zero_trust.AccessGroupNewParamsRequireAccessOktaGroupRule],
// [zero_trust.AccessGroupNewParamsRequireAccessSamlGroupRule],
// [zero_trust.AccessGroupNewParamsRequireAccessServiceTokenRule],
// [zero_trust.AccessGroupNewParamsRequireAccessAnyValidServiceTokenRule],
// [zero_trust.AccessGroupNewParamsRequireAccessExternalEvaluationRule],
// [zero_trust.AccessGroupNewParamsRequireAccessCountryRule],
// [zero_trust.AccessGroupNewParamsRequireAccessAuthenticationMethodRule],
// [zero_trust.AccessGroupNewParamsRequireAccessDevicePostureRule].
type AccessGroupNewParamsRequire interface {
	implementsZeroTrustAccessGroupNewParamsRequire()
}

// Matches a specific email.
type AccessGroupNewParamsRequireAccessEmailRule struct {
	Email param.Field[AccessGroupNewParamsRequireAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessGroupNewParamsRequireAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsRequireAccessEmailRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

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

func (r AccessGroupNewParamsRequireAccessEmailListRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

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

func (r AccessGroupNewParamsRequireAccessDomainRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

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

func (r AccessGroupNewParamsRequireAccessEveryoneRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

// Matches an IP address block.
type AccessGroupNewParamsRequireAccessIPRule struct {
	IP param.Field[AccessGroupNewParamsRequireAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessGroupNewParamsRequireAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsRequireAccessIPRule) implementsZeroTrustAccessGroupNewParamsRequire() {}

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

func (r AccessGroupNewParamsRequireAccessIPListRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

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

func (r AccessGroupNewParamsRequireAccessCertificateRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

// Matches an Access group.
type AccessGroupNewParamsRequireAccessAccessGroupRule struct {
	Group param.Field[AccessGroupNewParamsRequireAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessGroupNewParamsRequireAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsRequireAccessAccessGroupRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

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

func (r AccessGroupNewParamsRequireAccessAzureGroupRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

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

func (r AccessGroupNewParamsRequireAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupNewParamsRequire() {
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

func (r AccessGroupNewParamsRequireAccessGsuiteGroupRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

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

func (r AccessGroupNewParamsRequireAccessOktaGroupRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

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

func (r AccessGroupNewParamsRequireAccessSamlGroupRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

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

func (r AccessGroupNewParamsRequireAccessServiceTokenRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

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

func (r AccessGroupNewParamsRequireAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupNewParamsRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessGroupNewParamsRequireAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessGroupNewParamsRequireAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsRequireAccessExternalEvaluationRule) implementsZeroTrustAccessGroupNewParamsRequire() {
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

func (r AccessGroupNewParamsRequireAccessCountryRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

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

func (r AccessGroupNewParamsRequireAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupNewParamsRequire() {
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

func (r AccessGroupNewParamsRequireAccessDevicePostureRule) implementsZeroTrustAccessGroupNewParamsRequire() {
}

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
	Result   AccessGroups                             `json:"result,required"`
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

func (r accessGroupNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessGroupNewResponseEnvelopeSuccess bool

const (
	AccessGroupNewResponseEnvelopeSuccessTrue AccessGroupNewResponseEnvelopeSuccess = true
)

func (r AccessGroupNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessGroupNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessGroupUpdateParams struct {
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessGroupUpdateParamsInclude] `json:"include,required"`
	// The name of the Access group.
	Name param.Field[string] `json:"name,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
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
// Satisfied by [zero_trust.AccessGroupUpdateParamsIncludeAccessEmailRule],
// [zero_trust.AccessGroupUpdateParamsIncludeAccessEmailListRule],
// [zero_trust.AccessGroupUpdateParamsIncludeAccessDomainRule],
// [zero_trust.AccessGroupUpdateParamsIncludeAccessEveryoneRule],
// [zero_trust.AccessGroupUpdateParamsIncludeAccessIPRule],
// [zero_trust.AccessGroupUpdateParamsIncludeAccessIPListRule],
// [zero_trust.AccessGroupUpdateParamsIncludeAccessCertificateRule],
// [zero_trust.AccessGroupUpdateParamsIncludeAccessAccessGroupRule],
// [zero_trust.AccessGroupUpdateParamsIncludeAccessAzureGroupRule],
// [zero_trust.AccessGroupUpdateParamsIncludeAccessGitHubOrganizationRule],
// [zero_trust.AccessGroupUpdateParamsIncludeAccessGsuiteGroupRule],
// [zero_trust.AccessGroupUpdateParamsIncludeAccessOktaGroupRule],
// [zero_trust.AccessGroupUpdateParamsIncludeAccessSamlGroupRule],
// [zero_trust.AccessGroupUpdateParamsIncludeAccessServiceTokenRule],
// [zero_trust.AccessGroupUpdateParamsIncludeAccessAnyValidServiceTokenRule],
// [zero_trust.AccessGroupUpdateParamsIncludeAccessExternalEvaluationRule],
// [zero_trust.AccessGroupUpdateParamsIncludeAccessCountryRule],
// [zero_trust.AccessGroupUpdateParamsIncludeAccessAuthenticationMethodRule],
// [zero_trust.AccessGroupUpdateParamsIncludeAccessDevicePostureRule].
type AccessGroupUpdateParamsInclude interface {
	implementsZeroTrustAccessGroupUpdateParamsInclude()
}

// Matches a specific email.
type AccessGroupUpdateParamsIncludeAccessEmailRule struct {
	Email param.Field[AccessGroupUpdateParamsIncludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsIncludeAccessEmailRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
}

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

func (r AccessGroupUpdateParamsIncludeAccessEmailListRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
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

func (r AccessGroupUpdateParamsIncludeAccessDomainRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
}

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

func (r AccessGroupUpdateParamsIncludeAccessEveryoneRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
}

// Matches an IP address block.
type AccessGroupUpdateParamsIncludeAccessIPRule struct {
	IP param.Field[AccessGroupUpdateParamsIncludeAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsIncludeAccessIPRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
}

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

func (r AccessGroupUpdateParamsIncludeAccessIPListRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
}

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

func (r AccessGroupUpdateParamsIncludeAccessCertificateRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
}

// Matches an Access group.
type AccessGroupUpdateParamsIncludeAccessAccessGroupRule struct {
	Group param.Field[AccessGroupUpdateParamsIncludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsIncludeAccessAccessGroupRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
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

func (r AccessGroupUpdateParamsIncludeAccessAzureGroupRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
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

func (r AccessGroupUpdateParamsIncludeAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
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

func (r AccessGroupUpdateParamsIncludeAccessGsuiteGroupRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
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

func (r AccessGroupUpdateParamsIncludeAccessOktaGroupRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
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

func (r AccessGroupUpdateParamsIncludeAccessSamlGroupRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
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

func (r AccessGroupUpdateParamsIncludeAccessServiceTokenRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
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

func (r AccessGroupUpdateParamsIncludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupUpdateParamsIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessGroupUpdateParamsIncludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsIncludeAccessExternalEvaluationRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
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

func (r AccessGroupUpdateParamsIncludeAccessCountryRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
}

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

func (r AccessGroupUpdateParamsIncludeAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
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

func (r AccessGroupUpdateParamsIncludeAccessDevicePostureRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
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
// Satisfied by [zero_trust.AccessGroupUpdateParamsExcludeAccessEmailRule],
// [zero_trust.AccessGroupUpdateParamsExcludeAccessEmailListRule],
// [zero_trust.AccessGroupUpdateParamsExcludeAccessDomainRule],
// [zero_trust.AccessGroupUpdateParamsExcludeAccessEveryoneRule],
// [zero_trust.AccessGroupUpdateParamsExcludeAccessIPRule],
// [zero_trust.AccessGroupUpdateParamsExcludeAccessIPListRule],
// [zero_trust.AccessGroupUpdateParamsExcludeAccessCertificateRule],
// [zero_trust.AccessGroupUpdateParamsExcludeAccessAccessGroupRule],
// [zero_trust.AccessGroupUpdateParamsExcludeAccessAzureGroupRule],
// [zero_trust.AccessGroupUpdateParamsExcludeAccessGitHubOrganizationRule],
// [zero_trust.AccessGroupUpdateParamsExcludeAccessGsuiteGroupRule],
// [zero_trust.AccessGroupUpdateParamsExcludeAccessOktaGroupRule],
// [zero_trust.AccessGroupUpdateParamsExcludeAccessSamlGroupRule],
// [zero_trust.AccessGroupUpdateParamsExcludeAccessServiceTokenRule],
// [zero_trust.AccessGroupUpdateParamsExcludeAccessAnyValidServiceTokenRule],
// [zero_trust.AccessGroupUpdateParamsExcludeAccessExternalEvaluationRule],
// [zero_trust.AccessGroupUpdateParamsExcludeAccessCountryRule],
// [zero_trust.AccessGroupUpdateParamsExcludeAccessAuthenticationMethodRule],
// [zero_trust.AccessGroupUpdateParamsExcludeAccessDevicePostureRule].
type AccessGroupUpdateParamsExclude interface {
	implementsZeroTrustAccessGroupUpdateParamsExclude()
}

// Matches a specific email.
type AccessGroupUpdateParamsExcludeAccessEmailRule struct {
	Email param.Field[AccessGroupUpdateParamsExcludeAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsExcludeAccessEmailRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
}

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

func (r AccessGroupUpdateParamsExcludeAccessEmailListRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
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

func (r AccessGroupUpdateParamsExcludeAccessDomainRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
}

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

func (r AccessGroupUpdateParamsExcludeAccessEveryoneRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
}

// Matches an IP address block.
type AccessGroupUpdateParamsExcludeAccessIPRule struct {
	IP param.Field[AccessGroupUpdateParamsExcludeAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsExcludeAccessIPRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
}

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

func (r AccessGroupUpdateParamsExcludeAccessIPListRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
}

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

func (r AccessGroupUpdateParamsExcludeAccessCertificateRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
}

// Matches an Access group.
type AccessGroupUpdateParamsExcludeAccessAccessGroupRule struct {
	Group param.Field[AccessGroupUpdateParamsExcludeAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsExcludeAccessAccessGroupRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
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

func (r AccessGroupUpdateParamsExcludeAccessAzureGroupRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
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

func (r AccessGroupUpdateParamsExcludeAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
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

func (r AccessGroupUpdateParamsExcludeAccessGsuiteGroupRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
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

func (r AccessGroupUpdateParamsExcludeAccessOktaGroupRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
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

func (r AccessGroupUpdateParamsExcludeAccessSamlGroupRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
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

func (r AccessGroupUpdateParamsExcludeAccessServiceTokenRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
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

func (r AccessGroupUpdateParamsExcludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupUpdateParamsExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessGroupUpdateParamsExcludeAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsExcludeAccessExternalEvaluationRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
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

func (r AccessGroupUpdateParamsExcludeAccessCountryRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
}

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

func (r AccessGroupUpdateParamsExcludeAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
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

func (r AccessGroupUpdateParamsExcludeAccessDevicePostureRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
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
// Satisfied by [zero_trust.AccessGroupUpdateParamsRequireAccessEmailRule],
// [zero_trust.AccessGroupUpdateParamsRequireAccessEmailListRule],
// [zero_trust.AccessGroupUpdateParamsRequireAccessDomainRule],
// [zero_trust.AccessGroupUpdateParamsRequireAccessEveryoneRule],
// [zero_trust.AccessGroupUpdateParamsRequireAccessIPRule],
// [zero_trust.AccessGroupUpdateParamsRequireAccessIPListRule],
// [zero_trust.AccessGroupUpdateParamsRequireAccessCertificateRule],
// [zero_trust.AccessGroupUpdateParamsRequireAccessAccessGroupRule],
// [zero_trust.AccessGroupUpdateParamsRequireAccessAzureGroupRule],
// [zero_trust.AccessGroupUpdateParamsRequireAccessGitHubOrganizationRule],
// [zero_trust.AccessGroupUpdateParamsRequireAccessGsuiteGroupRule],
// [zero_trust.AccessGroupUpdateParamsRequireAccessOktaGroupRule],
// [zero_trust.AccessGroupUpdateParamsRequireAccessSamlGroupRule],
// [zero_trust.AccessGroupUpdateParamsRequireAccessServiceTokenRule],
// [zero_trust.AccessGroupUpdateParamsRequireAccessAnyValidServiceTokenRule],
// [zero_trust.AccessGroupUpdateParamsRequireAccessExternalEvaluationRule],
// [zero_trust.AccessGroupUpdateParamsRequireAccessCountryRule],
// [zero_trust.AccessGroupUpdateParamsRequireAccessAuthenticationMethodRule],
// [zero_trust.AccessGroupUpdateParamsRequireAccessDevicePostureRule].
type AccessGroupUpdateParamsRequire interface {
	implementsZeroTrustAccessGroupUpdateParamsRequire()
}

// Matches a specific email.
type AccessGroupUpdateParamsRequireAccessEmailRule struct {
	Email param.Field[AccessGroupUpdateParamsRequireAccessEmailRuleEmail] `json:"email,required"`
}

func (r AccessGroupUpdateParamsRequireAccessEmailRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsRequireAccessEmailRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
}

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

func (r AccessGroupUpdateParamsRequireAccessEmailListRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
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

func (r AccessGroupUpdateParamsRequireAccessDomainRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
}

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

func (r AccessGroupUpdateParamsRequireAccessEveryoneRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
}

// Matches an IP address block.
type AccessGroupUpdateParamsRequireAccessIPRule struct {
	IP param.Field[AccessGroupUpdateParamsRequireAccessIPRuleIP] `json:"ip,required"`
}

func (r AccessGroupUpdateParamsRequireAccessIPRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsRequireAccessIPRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
}

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

func (r AccessGroupUpdateParamsRequireAccessIPListRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
}

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

func (r AccessGroupUpdateParamsRequireAccessCertificateRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
}

// Matches an Access group.
type AccessGroupUpdateParamsRequireAccessAccessGroupRule struct {
	Group param.Field[AccessGroupUpdateParamsRequireAccessAccessGroupRuleGroup] `json:"group,required"`
}

func (r AccessGroupUpdateParamsRequireAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsRequireAccessAccessGroupRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
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

func (r AccessGroupUpdateParamsRequireAccessAzureGroupRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
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

func (r AccessGroupUpdateParamsRequireAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
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

func (r AccessGroupUpdateParamsRequireAccessGsuiteGroupRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
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

func (r AccessGroupUpdateParamsRequireAccessOktaGroupRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
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

func (r AccessGroupUpdateParamsRequireAccessSamlGroupRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
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

func (r AccessGroupUpdateParamsRequireAccessServiceTokenRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
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

func (r AccessGroupUpdateParamsRequireAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type AccessGroupUpdateParamsRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation param.Field[AccessGroupUpdateParamsRequireAccessExternalEvaluationRuleExternalEvaluation] `json:"external_evaluation,required"`
}

func (r AccessGroupUpdateParamsRequireAccessExternalEvaluationRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsRequireAccessExternalEvaluationRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
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

func (r AccessGroupUpdateParamsRequireAccessCountryRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
}

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

func (r AccessGroupUpdateParamsRequireAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
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

func (r AccessGroupUpdateParamsRequireAccessDevicePostureRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
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
	Result   AccessGroups                                `json:"result,required"`
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

func (r accessGroupUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessGroupUpdateResponseEnvelopeSuccess bool

const (
	AccessGroupUpdateResponseEnvelopeSuccessTrue AccessGroupUpdateResponseEnvelopeSuccess = true
)

func (r AccessGroupUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessGroupUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessGroupListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessGroupListResponseEnvelope struct {
	Errors   []AccessGroupListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessGroupListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessGroups                            `json:"result,required,nullable"`
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

func (r accessGroupListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessGroupListResponseEnvelopeSuccess bool

const (
	AccessGroupListResponseEnvelopeSuccessTrue AccessGroupListResponseEnvelopeSuccess = true
)

func (r AccessGroupListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessGroupListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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

func (r accessGroupListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccessGroupDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
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

func (r accessGroupDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessGroupDeleteResponseEnvelopeSuccess bool

const (
	AccessGroupDeleteResponseEnvelopeSuccessTrue AccessGroupDeleteResponseEnvelopeSuccess = true
)

func (r AccessGroupDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessGroupDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessGroupGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessGroupGetResponseEnvelope struct {
	Errors   []AccessGroupGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessGroupGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessGroups                             `json:"result,required"`
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

func (r accessGroupGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r accessGroupGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessGroupGetResponseEnvelopeSuccess bool

const (
	AccessGroupGetResponseEnvelopeSuccessTrue AccessGroupGetResponseEnvelopeSuccess = true
)

func (r AccessGroupGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessGroupGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
