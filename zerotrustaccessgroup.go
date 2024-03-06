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
func (r *ZeroTrustAccessGroupService) New(ctx context.Context, params ZeroTrustAccessGroupNewParams, opts ...option.RequestOption) (res *ZeroTrustAccessGroupNewResponse, err error) {
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
func (r *ZeroTrustAccessGroupService) Update(ctx context.Context, uuid string, params ZeroTrustAccessGroupUpdateParams, opts ...option.RequestOption) (res *ZeroTrustAccessGroupUpdateResponse, err error) {
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
func (r *ZeroTrustAccessGroupService) List(ctx context.Context, query ZeroTrustAccessGroupListParams, opts ...option.RequestOption) (res *[]ZeroTrustAccessGroupListResponse, err error) {
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
func (r *ZeroTrustAccessGroupService) Get(ctx context.Context, uuid string, query ZeroTrustAccessGroupGetParams, opts ...option.RequestOption) (res *ZeroTrustAccessGroupGetResponse, err error) {
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

type ZeroTrustAccessGroupNewResponse struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []ZeroTrustAccessGroupNewResponseExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []ZeroTrustAccessGroupNewResponseInclude `json:"include"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	IsDefault []ZeroTrustAccessGroupNewResponseIsDefault `json:"is_default"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []ZeroTrustAccessGroupNewResponseRequire `json:"require"`
	UpdatedAt time.Time                                `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessGroupNewResponseJSON      `json:"-"`
}

// zeroTrustAccessGroupNewResponseJSON contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponse]
type zeroTrustAccessGroupNewResponseJSON struct {
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

func (r *ZeroTrustAccessGroupNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by [ZeroTrustAccessGroupNewResponseExcludeAccessEmailRule],
// [ZeroTrustAccessGroupNewResponseExcludeAccessEmailListRule],
// [ZeroTrustAccessGroupNewResponseExcludeAccessDomainRule],
// [ZeroTrustAccessGroupNewResponseExcludeAccessEveryoneRule],
// [ZeroTrustAccessGroupNewResponseExcludeAccessIPRule],
// [ZeroTrustAccessGroupNewResponseExcludeAccessIPListRule],
// [ZeroTrustAccessGroupNewResponseExcludeAccessCertificateRule],
// [ZeroTrustAccessGroupNewResponseExcludeAccessAccessGroupRule],
// [ZeroTrustAccessGroupNewResponseExcludeAccessAzureGroupRule],
// [ZeroTrustAccessGroupNewResponseExcludeAccessGitHubOrganizationRule],
// [ZeroTrustAccessGroupNewResponseExcludeAccessGsuiteGroupRule],
// [ZeroTrustAccessGroupNewResponseExcludeAccessOktaGroupRule],
// [ZeroTrustAccessGroupNewResponseExcludeAccessSamlGroupRule],
// [ZeroTrustAccessGroupNewResponseExcludeAccessServiceTokenRule],
// [ZeroTrustAccessGroupNewResponseExcludeAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessGroupNewResponseExcludeAccessExternalEvaluationRule],
// [ZeroTrustAccessGroupNewResponseExcludeAccessCountryRule],
// [ZeroTrustAccessGroupNewResponseExcludeAccessAuthenticationMethodRule] or
// [ZeroTrustAccessGroupNewResponseExcludeAccessDevicePostureRule].
type ZeroTrustAccessGroupNewResponseExclude interface {
	implementsZeroTrustAccessGroupNewResponseExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustAccessGroupNewResponseExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type ZeroTrustAccessGroupNewResponseExcludeAccessEmailRule struct {
	Email ZeroTrustAccessGroupNewResponseExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustAccessGroupNewResponseExcludeAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessEmailRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupNewResponseExcludeAccessEmailRule]
type zeroTrustAccessGroupNewResponseExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseExcludeAccessEmailRule) implementsZeroTrustAccessGroupNewResponseExclude() {
}

type ZeroTrustAccessGroupNewResponseExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                         `json:"email,required" format:"email"`
	JSON  zeroTrustAccessGroupNewResponseExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupNewResponseExcludeAccessEmailRuleEmail]
type zeroTrustAccessGroupNewResponseExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustAccessGroupNewResponseExcludeAccessEmailListRule struct {
	EmailList ZeroTrustAccessGroupNewResponseExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustAccessGroupNewResponseExcludeAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessEmailListRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupNewResponseExcludeAccessEmailListRule]
type zeroTrustAccessGroupNewResponseExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseExcludeAccessEmailListRule) implementsZeroTrustAccessGroupNewResponseExclude() {
}

type ZeroTrustAccessGroupNewResponseExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                 `json:"id,required"`
	JSON zeroTrustAccessGroupNewResponseExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessEmailListRuleEmailListJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseExcludeAccessEmailListRuleEmailList]
type zeroTrustAccessGroupNewResponseExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustAccessGroupNewResponseExcludeAccessDomainRule struct {
	EmailDomain ZeroTrustAccessGroupNewResponseExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustAccessGroupNewResponseExcludeAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessDomainRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupNewResponseExcludeAccessDomainRule]
type zeroTrustAccessGroupNewResponseExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseExcludeAccessDomainRule) implementsZeroTrustAccessGroupNewResponseExclude() {
}

type ZeroTrustAccessGroupNewResponseExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                `json:"domain,required"`
	JSON   zeroTrustAccessGroupNewResponseExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessDomainRuleEmailDomainJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseExcludeAccessDomainRuleEmailDomain]
type zeroTrustAccessGroupNewResponseExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustAccessGroupNewResponseExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                  `json:"everyone,required"`
	JSON     zeroTrustAccessGroupNewResponseExcludeAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessEveryoneRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupNewResponseExcludeAccessEveryoneRule]
type zeroTrustAccessGroupNewResponseExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseExcludeAccessEveryoneRule) implementsZeroTrustAccessGroupNewResponseExclude() {
}

// Matches an IP address block.
type ZeroTrustAccessGroupNewResponseExcludeAccessIPRule struct {
	IP   ZeroTrustAccessGroupNewResponseExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustAccessGroupNewResponseExcludeAccessIPRuleJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessIPRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupNewResponseExcludeAccessIPRule]
type zeroTrustAccessGroupNewResponseExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseExcludeAccessIPRule) implementsZeroTrustAccessGroupNewResponseExclude() {
}

type ZeroTrustAccessGroupNewResponseExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                   `json:"ip,required"`
	JSON zeroTrustAccessGroupNewResponseExcludeAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessIPRuleIPJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupNewResponseExcludeAccessIPRuleIP]
type zeroTrustAccessGroupNewResponseExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustAccessGroupNewResponseExcludeAccessIPListRule struct {
	IPList ZeroTrustAccessGroupNewResponseExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustAccessGroupNewResponseExcludeAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessIPListRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupNewResponseExcludeAccessIPListRule]
type zeroTrustAccessGroupNewResponseExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseExcludeAccessIPListRule) implementsZeroTrustAccessGroupNewResponseExclude() {
}

type ZeroTrustAccessGroupNewResponseExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                           `json:"id,required"`
	JSON zeroTrustAccessGroupNewResponseExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseExcludeAccessIPListRuleIPList]
type zeroTrustAccessGroupNewResponseExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustAccessGroupNewResponseExcludeAccessCertificateRule struct {
	Certificate interface{}                                                     `json:"certificate,required"`
	JSON        zeroTrustAccessGroupNewResponseExcludeAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseExcludeAccessCertificateRule]
type zeroTrustAccessGroupNewResponseExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseExcludeAccessCertificateRule) implementsZeroTrustAccessGroupNewResponseExclude() {
}

// Matches an Access group.
type ZeroTrustAccessGroupNewResponseExcludeAccessAccessGroupRule struct {
	Group ZeroTrustAccessGroupNewResponseExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  zeroTrustAccessGroupNewResponseExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseExcludeAccessAccessGroupRule]
type zeroTrustAccessGroupNewResponseExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseExcludeAccessAccessGroupRule) implementsZeroTrustAccessGroupNewResponseExclude() {
}

type ZeroTrustAccessGroupNewResponseExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                               `json:"id,required"`
	JSON zeroTrustAccessGroupNewResponseExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseExcludeAccessAccessGroupRuleGroup]
type zeroTrustAccessGroupNewResponseExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessGroupNewResponseExcludeAccessAzureGroupRule struct {
	AzureAd ZeroTrustAccessGroupNewResponseExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustAccessGroupNewResponseExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupNewResponseExcludeAccessAzureGroupRule]
type zeroTrustAccessGroupNewResponseExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseExcludeAccessAzureGroupRule) implementsZeroTrustAccessGroupNewResponseExclude() {
}

type ZeroTrustAccessGroupNewResponseExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                `json:"connection_id,required"`
	JSON         zeroTrustAccessGroupNewResponseExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessAzureGroupRuleAzureAdJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseExcludeAccessAzureGroupRuleAzureAd]
type zeroTrustAccessGroupNewResponseExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessGroupNewResponseExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustAccessGroupNewResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustAccessGroupNewResponseExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessGitHubOrganizationRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseExcludeAccessGitHubOrganizationRule]
type zeroTrustAccessGroupNewResponseExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseExcludeAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupNewResponseExclude() {
}

type ZeroTrustAccessGroupNewResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                   `json:"name,required"`
	JSON zeroTrustAccessGroupNewResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustAccessGroupNewResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessGroupNewResponseExcludeAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustAccessGroupNewResponseExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustAccessGroupNewResponseExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseExcludeAccessGsuiteGroupRule]
type zeroTrustAccessGroupNewResponseExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseExcludeAccessGsuiteGroupRule) implementsZeroTrustAccessGroupNewResponseExclude() {
}

type ZeroTrustAccessGroupNewResponseExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                `json:"email,required"`
	JSON  zeroTrustAccessGroupNewResponseExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessGsuiteGroupRuleGsuiteJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseExcludeAccessGsuiteGroupRuleGsuite]
type zeroTrustAccessGroupNewResponseExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessGroupNewResponseExcludeAccessOktaGroupRule struct {
	Okta ZeroTrustAccessGroupNewResponseExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustAccessGroupNewResponseExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupNewResponseExcludeAccessOktaGroupRule]
type zeroTrustAccessGroupNewResponseExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseExcludeAccessOktaGroupRule) implementsZeroTrustAccessGroupNewResponseExclude() {
}

type ZeroTrustAccessGroupNewResponseExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                            `json:"email,required"`
	JSON  zeroTrustAccessGroupNewResponseExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessOktaGroupRuleOktaJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseExcludeAccessOktaGroupRuleOkta]
type zeroTrustAccessGroupNewResponseExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessGroupNewResponseExcludeAccessSamlGroupRule struct {
	Saml ZeroTrustAccessGroupNewResponseExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustAccessGroupNewResponseExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupNewResponseExcludeAccessSamlGroupRule]
type zeroTrustAccessGroupNewResponseExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseExcludeAccessSamlGroupRule) implementsZeroTrustAccessGroupNewResponseExclude() {
}

type ZeroTrustAccessGroupNewResponseExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                            `json:"attribute_value,required"`
	JSON           zeroTrustAccessGroupNewResponseExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessSamlGroupRuleSamlJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseExcludeAccessSamlGroupRuleSaml]
type zeroTrustAccessGroupNewResponseExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustAccessGroupNewResponseExcludeAccessServiceTokenRule struct {
	ServiceToken ZeroTrustAccessGroupNewResponseExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustAccessGroupNewResponseExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseExcludeAccessServiceTokenRule]
type zeroTrustAccessGroupNewResponseExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseExcludeAccessServiceTokenRule) implementsZeroTrustAccessGroupNewResponseExclude() {
}

type ZeroTrustAccessGroupNewResponseExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                       `json:"token_id,required"`
	JSON    zeroTrustAccessGroupNewResponseExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseExcludeAccessServiceTokenRuleServiceToken]
type zeroTrustAccessGroupNewResponseExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustAccessGroupNewResponseExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                              `json:"any_valid_service_token,required"`
	JSON                 zeroTrustAccessGroupNewResponseExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseExcludeAccessAnyValidServiceTokenRule]
type zeroTrustAccessGroupNewResponseExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseExcludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupNewResponseExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessGroupNewResponseExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustAccessGroupNewResponseExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustAccessGroupNewResponseExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessExternalEvaluationRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseExcludeAccessExternalEvaluationRule]
type zeroTrustAccessGroupNewResponseExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseExcludeAccessExternalEvaluationRule) implementsZeroTrustAccessGroupNewResponseExclude() {
}

type ZeroTrustAccessGroupNewResponseExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                   `json:"keys_url,required"`
	JSON    zeroTrustAccessGroupNewResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseExcludeAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustAccessGroupNewResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustAccessGroupNewResponseExcludeAccessCountryRule struct {
	Geo  ZeroTrustAccessGroupNewResponseExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustAccessGroupNewResponseExcludeAccessCountryRuleJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessCountryRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupNewResponseExcludeAccessCountryRule]
type zeroTrustAccessGroupNewResponseExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseExcludeAccessCountryRule) implementsZeroTrustAccessGroupNewResponseExclude() {
}

type ZeroTrustAccessGroupNewResponseExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                         `json:"country_code,required"`
	JSON        zeroTrustAccessGroupNewResponseExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupNewResponseExcludeAccessCountryRuleGeo]
type zeroTrustAccessGroupNewResponseExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustAccessGroupNewResponseExcludeAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustAccessGroupNewResponseExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupNewResponseExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseExcludeAccessAuthenticationMethodRule]
type zeroTrustAccessGroupNewResponseExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseExcludeAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupNewResponseExclude() {
}

type ZeroTrustAccessGroupNewResponseExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                             `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupNewResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseExcludeAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustAccessGroupNewResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessGroupNewResponseExcludeAccessDevicePostureRule struct {
	DevicePosture ZeroTrustAccessGroupNewResponseExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustAccessGroupNewResponseExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessDevicePostureRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseExcludeAccessDevicePostureRule]
type zeroTrustAccessGroupNewResponseExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseExcludeAccessDevicePostureRule) implementsZeroTrustAccessGroupNewResponseExclude() {
}

type ZeroTrustAccessGroupNewResponseExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                         `json:"integration_uid,required"`
	JSON           zeroTrustAccessGroupNewResponseExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseExcludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseExcludeAccessDevicePostureRuleDevicePosture]
type zeroTrustAccessGroupNewResponseExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseExcludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by [ZeroTrustAccessGroupNewResponseIncludeAccessEmailRule],
// [ZeroTrustAccessGroupNewResponseIncludeAccessEmailListRule],
// [ZeroTrustAccessGroupNewResponseIncludeAccessDomainRule],
// [ZeroTrustAccessGroupNewResponseIncludeAccessEveryoneRule],
// [ZeroTrustAccessGroupNewResponseIncludeAccessIPRule],
// [ZeroTrustAccessGroupNewResponseIncludeAccessIPListRule],
// [ZeroTrustAccessGroupNewResponseIncludeAccessCertificateRule],
// [ZeroTrustAccessGroupNewResponseIncludeAccessAccessGroupRule],
// [ZeroTrustAccessGroupNewResponseIncludeAccessAzureGroupRule],
// [ZeroTrustAccessGroupNewResponseIncludeAccessGitHubOrganizationRule],
// [ZeroTrustAccessGroupNewResponseIncludeAccessGsuiteGroupRule],
// [ZeroTrustAccessGroupNewResponseIncludeAccessOktaGroupRule],
// [ZeroTrustAccessGroupNewResponseIncludeAccessSamlGroupRule],
// [ZeroTrustAccessGroupNewResponseIncludeAccessServiceTokenRule],
// [ZeroTrustAccessGroupNewResponseIncludeAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessGroupNewResponseIncludeAccessExternalEvaluationRule],
// [ZeroTrustAccessGroupNewResponseIncludeAccessCountryRule],
// [ZeroTrustAccessGroupNewResponseIncludeAccessAuthenticationMethodRule] or
// [ZeroTrustAccessGroupNewResponseIncludeAccessDevicePostureRule].
type ZeroTrustAccessGroupNewResponseInclude interface {
	implementsZeroTrustAccessGroupNewResponseInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustAccessGroupNewResponseInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type ZeroTrustAccessGroupNewResponseIncludeAccessEmailRule struct {
	Email ZeroTrustAccessGroupNewResponseIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustAccessGroupNewResponseIncludeAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessEmailRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupNewResponseIncludeAccessEmailRule]
type zeroTrustAccessGroupNewResponseIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIncludeAccessEmailRule) implementsZeroTrustAccessGroupNewResponseInclude() {
}

type ZeroTrustAccessGroupNewResponseIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                         `json:"email,required" format:"email"`
	JSON  zeroTrustAccessGroupNewResponseIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupNewResponseIncludeAccessEmailRuleEmail]
type zeroTrustAccessGroupNewResponseIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustAccessGroupNewResponseIncludeAccessEmailListRule struct {
	EmailList ZeroTrustAccessGroupNewResponseIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustAccessGroupNewResponseIncludeAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessEmailListRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupNewResponseIncludeAccessEmailListRule]
type zeroTrustAccessGroupNewResponseIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIncludeAccessEmailListRule) implementsZeroTrustAccessGroupNewResponseInclude() {
}

type ZeroTrustAccessGroupNewResponseIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                 `json:"id,required"`
	JSON zeroTrustAccessGroupNewResponseIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessEmailListRuleEmailListJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIncludeAccessEmailListRuleEmailList]
type zeroTrustAccessGroupNewResponseIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustAccessGroupNewResponseIncludeAccessDomainRule struct {
	EmailDomain ZeroTrustAccessGroupNewResponseIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustAccessGroupNewResponseIncludeAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessDomainRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupNewResponseIncludeAccessDomainRule]
type zeroTrustAccessGroupNewResponseIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIncludeAccessDomainRule) implementsZeroTrustAccessGroupNewResponseInclude() {
}

type ZeroTrustAccessGroupNewResponseIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                `json:"domain,required"`
	JSON   zeroTrustAccessGroupNewResponseIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessDomainRuleEmailDomainJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIncludeAccessDomainRuleEmailDomain]
type zeroTrustAccessGroupNewResponseIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustAccessGroupNewResponseIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                  `json:"everyone,required"`
	JSON     zeroTrustAccessGroupNewResponseIncludeAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessEveryoneRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupNewResponseIncludeAccessEveryoneRule]
type zeroTrustAccessGroupNewResponseIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIncludeAccessEveryoneRule) implementsZeroTrustAccessGroupNewResponseInclude() {
}

// Matches an IP address block.
type ZeroTrustAccessGroupNewResponseIncludeAccessIPRule struct {
	IP   ZeroTrustAccessGroupNewResponseIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustAccessGroupNewResponseIncludeAccessIPRuleJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessIPRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupNewResponseIncludeAccessIPRule]
type zeroTrustAccessGroupNewResponseIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIncludeAccessIPRule) implementsZeroTrustAccessGroupNewResponseInclude() {
}

type ZeroTrustAccessGroupNewResponseIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                   `json:"ip,required"`
	JSON zeroTrustAccessGroupNewResponseIncludeAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessIPRuleIPJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupNewResponseIncludeAccessIPRuleIP]
type zeroTrustAccessGroupNewResponseIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustAccessGroupNewResponseIncludeAccessIPListRule struct {
	IPList ZeroTrustAccessGroupNewResponseIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustAccessGroupNewResponseIncludeAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessIPListRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupNewResponseIncludeAccessIPListRule]
type zeroTrustAccessGroupNewResponseIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIncludeAccessIPListRule) implementsZeroTrustAccessGroupNewResponseInclude() {
}

type ZeroTrustAccessGroupNewResponseIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                           `json:"id,required"`
	JSON zeroTrustAccessGroupNewResponseIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIncludeAccessIPListRuleIPList]
type zeroTrustAccessGroupNewResponseIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustAccessGroupNewResponseIncludeAccessCertificateRule struct {
	Certificate interface{}                                                     `json:"certificate,required"`
	JSON        zeroTrustAccessGroupNewResponseIncludeAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIncludeAccessCertificateRule]
type zeroTrustAccessGroupNewResponseIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIncludeAccessCertificateRule) implementsZeroTrustAccessGroupNewResponseInclude() {
}

// Matches an Access group.
type ZeroTrustAccessGroupNewResponseIncludeAccessAccessGroupRule struct {
	Group ZeroTrustAccessGroupNewResponseIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  zeroTrustAccessGroupNewResponseIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIncludeAccessAccessGroupRule]
type zeroTrustAccessGroupNewResponseIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIncludeAccessAccessGroupRule) implementsZeroTrustAccessGroupNewResponseInclude() {
}

type ZeroTrustAccessGroupNewResponseIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                               `json:"id,required"`
	JSON zeroTrustAccessGroupNewResponseIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIncludeAccessAccessGroupRuleGroup]
type zeroTrustAccessGroupNewResponseIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessGroupNewResponseIncludeAccessAzureGroupRule struct {
	AzureAd ZeroTrustAccessGroupNewResponseIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustAccessGroupNewResponseIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupNewResponseIncludeAccessAzureGroupRule]
type zeroTrustAccessGroupNewResponseIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIncludeAccessAzureGroupRule) implementsZeroTrustAccessGroupNewResponseInclude() {
}

type ZeroTrustAccessGroupNewResponseIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                `json:"connection_id,required"`
	JSON         zeroTrustAccessGroupNewResponseIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessAzureGroupRuleAzureAdJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIncludeAccessAzureGroupRuleAzureAd]
type zeroTrustAccessGroupNewResponseIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessGroupNewResponseIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustAccessGroupNewResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustAccessGroupNewResponseIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessGitHubOrganizationRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIncludeAccessGitHubOrganizationRule]
type zeroTrustAccessGroupNewResponseIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIncludeAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupNewResponseInclude() {
}

type ZeroTrustAccessGroupNewResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                   `json:"name,required"`
	JSON zeroTrustAccessGroupNewResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustAccessGroupNewResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessGroupNewResponseIncludeAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustAccessGroupNewResponseIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustAccessGroupNewResponseIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIncludeAccessGsuiteGroupRule]
type zeroTrustAccessGroupNewResponseIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIncludeAccessGsuiteGroupRule) implementsZeroTrustAccessGroupNewResponseInclude() {
}

type ZeroTrustAccessGroupNewResponseIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                `json:"email,required"`
	JSON  zeroTrustAccessGroupNewResponseIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessGsuiteGroupRuleGsuiteJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIncludeAccessGsuiteGroupRuleGsuite]
type zeroTrustAccessGroupNewResponseIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessGroupNewResponseIncludeAccessOktaGroupRule struct {
	Okta ZeroTrustAccessGroupNewResponseIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustAccessGroupNewResponseIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupNewResponseIncludeAccessOktaGroupRule]
type zeroTrustAccessGroupNewResponseIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIncludeAccessOktaGroupRule) implementsZeroTrustAccessGroupNewResponseInclude() {
}

type ZeroTrustAccessGroupNewResponseIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                            `json:"email,required"`
	JSON  zeroTrustAccessGroupNewResponseIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessOktaGroupRuleOktaJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIncludeAccessOktaGroupRuleOkta]
type zeroTrustAccessGroupNewResponseIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessGroupNewResponseIncludeAccessSamlGroupRule struct {
	Saml ZeroTrustAccessGroupNewResponseIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustAccessGroupNewResponseIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupNewResponseIncludeAccessSamlGroupRule]
type zeroTrustAccessGroupNewResponseIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIncludeAccessSamlGroupRule) implementsZeroTrustAccessGroupNewResponseInclude() {
}

type ZeroTrustAccessGroupNewResponseIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                            `json:"attribute_value,required"`
	JSON           zeroTrustAccessGroupNewResponseIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessSamlGroupRuleSamlJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIncludeAccessSamlGroupRuleSaml]
type zeroTrustAccessGroupNewResponseIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustAccessGroupNewResponseIncludeAccessServiceTokenRule struct {
	ServiceToken ZeroTrustAccessGroupNewResponseIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustAccessGroupNewResponseIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIncludeAccessServiceTokenRule]
type zeroTrustAccessGroupNewResponseIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIncludeAccessServiceTokenRule) implementsZeroTrustAccessGroupNewResponseInclude() {
}

type ZeroTrustAccessGroupNewResponseIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                       `json:"token_id,required"`
	JSON    zeroTrustAccessGroupNewResponseIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIncludeAccessServiceTokenRuleServiceToken]
type zeroTrustAccessGroupNewResponseIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustAccessGroupNewResponseIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                              `json:"any_valid_service_token,required"`
	JSON                 zeroTrustAccessGroupNewResponseIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIncludeAccessAnyValidServiceTokenRule]
type zeroTrustAccessGroupNewResponseIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIncludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupNewResponseInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessGroupNewResponseIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustAccessGroupNewResponseIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustAccessGroupNewResponseIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessExternalEvaluationRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIncludeAccessExternalEvaluationRule]
type zeroTrustAccessGroupNewResponseIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIncludeAccessExternalEvaluationRule) implementsZeroTrustAccessGroupNewResponseInclude() {
}

type ZeroTrustAccessGroupNewResponseIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                   `json:"keys_url,required"`
	JSON    zeroTrustAccessGroupNewResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIncludeAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustAccessGroupNewResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustAccessGroupNewResponseIncludeAccessCountryRule struct {
	Geo  ZeroTrustAccessGroupNewResponseIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustAccessGroupNewResponseIncludeAccessCountryRuleJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessCountryRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupNewResponseIncludeAccessCountryRule]
type zeroTrustAccessGroupNewResponseIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIncludeAccessCountryRule) implementsZeroTrustAccessGroupNewResponseInclude() {
}

type ZeroTrustAccessGroupNewResponseIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                         `json:"country_code,required"`
	JSON        zeroTrustAccessGroupNewResponseIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupNewResponseIncludeAccessCountryRuleGeo]
type zeroTrustAccessGroupNewResponseIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustAccessGroupNewResponseIncludeAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustAccessGroupNewResponseIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupNewResponseIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIncludeAccessAuthenticationMethodRule]
type zeroTrustAccessGroupNewResponseIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIncludeAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupNewResponseInclude() {
}

type ZeroTrustAccessGroupNewResponseIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                             `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupNewResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIncludeAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustAccessGroupNewResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessGroupNewResponseIncludeAccessDevicePostureRule struct {
	DevicePosture ZeroTrustAccessGroupNewResponseIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustAccessGroupNewResponseIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessDevicePostureRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIncludeAccessDevicePostureRule]
type zeroTrustAccessGroupNewResponseIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIncludeAccessDevicePostureRule) implementsZeroTrustAccessGroupNewResponseInclude() {
}

type ZeroTrustAccessGroupNewResponseIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                         `json:"integration_uid,required"`
	JSON           zeroTrustAccessGroupNewResponseIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIncludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIncludeAccessDevicePostureRuleDevicePosture]
type zeroTrustAccessGroupNewResponseIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIncludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by [ZeroTrustAccessGroupNewResponseIsDefaultAccessEmailRule],
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessEmailListRule],
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessDomainRule],
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessEveryoneRule],
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessIPRule],
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessIPListRule],
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessCertificateRule],
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessAccessGroupRule],
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessAzureGroupRule],
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessGitHubOrganizationRule],
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessGsuiteGroupRule],
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessOktaGroupRule],
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessSamlGroupRule],
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessServiceTokenRule],
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessExternalEvaluationRule],
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessCountryRule],
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessAuthenticationMethodRule] or
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessDevicePostureRule].
type ZeroTrustAccessGroupNewResponseIsDefault interface {
	implementsZeroTrustAccessGroupNewResponseIsDefault()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustAccessGroupNewResponseIsDefault)(nil)).Elem(), "")
}

// Matches a specific email.
type ZeroTrustAccessGroupNewResponseIsDefaultAccessEmailRule struct {
	Email ZeroTrustAccessGroupNewResponseIsDefaultAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustAccessGroupNewResponseIsDefaultAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessEmailRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessEmailRule]
type zeroTrustAccessGroupNewResponseIsDefaultAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIsDefaultAccessEmailRule) implementsZeroTrustAccessGroupNewResponseIsDefault() {
}

type ZeroTrustAccessGroupNewResponseIsDefaultAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                           `json:"email,required" format:"email"`
	JSON  zeroTrustAccessGroupNewResponseIsDefaultAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessEmailRuleEmail]
type zeroTrustAccessGroupNewResponseIsDefaultAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustAccessGroupNewResponseIsDefaultAccessEmailListRule struct {
	EmailList ZeroTrustAccessGroupNewResponseIsDefaultAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustAccessGroupNewResponseIsDefaultAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessEmailListRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessEmailListRule]
type zeroTrustAccessGroupNewResponseIsDefaultAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIsDefaultAccessEmailListRule) implementsZeroTrustAccessGroupNewResponseIsDefault() {
}

type ZeroTrustAccessGroupNewResponseIsDefaultAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                   `json:"id,required"`
	JSON zeroTrustAccessGroupNewResponseIsDefaultAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessEmailListRuleEmailList]
type zeroTrustAccessGroupNewResponseIsDefaultAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustAccessGroupNewResponseIsDefaultAccessDomainRule struct {
	EmailDomain ZeroTrustAccessGroupNewResponseIsDefaultAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustAccessGroupNewResponseIsDefaultAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessDomainRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessDomainRule]
type zeroTrustAccessGroupNewResponseIsDefaultAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIsDefaultAccessDomainRule) implementsZeroTrustAccessGroupNewResponseIsDefault() {
}

type ZeroTrustAccessGroupNewResponseIsDefaultAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                  `json:"domain,required"`
	JSON   zeroTrustAccessGroupNewResponseIsDefaultAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessDomainRuleEmailDomainJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessDomainRuleEmailDomain]
type zeroTrustAccessGroupNewResponseIsDefaultAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustAccessGroupNewResponseIsDefaultAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                    `json:"everyone,required"`
	JSON     zeroTrustAccessGroupNewResponseIsDefaultAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessEveryoneRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessEveryoneRule]
type zeroTrustAccessGroupNewResponseIsDefaultAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIsDefaultAccessEveryoneRule) implementsZeroTrustAccessGroupNewResponseIsDefault() {
}

// Matches an IP address block.
type ZeroTrustAccessGroupNewResponseIsDefaultAccessIPRule struct {
	IP   ZeroTrustAccessGroupNewResponseIsDefaultAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustAccessGroupNewResponseIsDefaultAccessIPRuleJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessIPRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupNewResponseIsDefaultAccessIPRule]
type zeroTrustAccessGroupNewResponseIsDefaultAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIsDefaultAccessIPRule) implementsZeroTrustAccessGroupNewResponseIsDefault() {
}

type ZeroTrustAccessGroupNewResponseIsDefaultAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                     `json:"ip,required"`
	JSON zeroTrustAccessGroupNewResponseIsDefaultAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessIPRuleIPJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupNewResponseIsDefaultAccessIPRuleIP]
type zeroTrustAccessGroupNewResponseIsDefaultAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustAccessGroupNewResponseIsDefaultAccessIPListRule struct {
	IPList ZeroTrustAccessGroupNewResponseIsDefaultAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustAccessGroupNewResponseIsDefaultAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessIPListRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessIPListRule]
type zeroTrustAccessGroupNewResponseIsDefaultAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIsDefaultAccessIPListRule) implementsZeroTrustAccessGroupNewResponseIsDefault() {
}

type ZeroTrustAccessGroupNewResponseIsDefaultAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                             `json:"id,required"`
	JSON zeroTrustAccessGroupNewResponseIsDefaultAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessIPListRuleIPList]
type zeroTrustAccessGroupNewResponseIsDefaultAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustAccessGroupNewResponseIsDefaultAccessCertificateRule struct {
	Certificate interface{}                                                       `json:"certificate,required"`
	JSON        zeroTrustAccessGroupNewResponseIsDefaultAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessCertificateRule]
type zeroTrustAccessGroupNewResponseIsDefaultAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIsDefaultAccessCertificateRule) implementsZeroTrustAccessGroupNewResponseIsDefault() {
}

// Matches an Access group.
type ZeroTrustAccessGroupNewResponseIsDefaultAccessAccessGroupRule struct {
	Group ZeroTrustAccessGroupNewResponseIsDefaultAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  zeroTrustAccessGroupNewResponseIsDefaultAccessAccessGroupRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessAccessGroupRule]
type zeroTrustAccessGroupNewResponseIsDefaultAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIsDefaultAccessAccessGroupRule) implementsZeroTrustAccessGroupNewResponseIsDefault() {
}

type ZeroTrustAccessGroupNewResponseIsDefaultAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                 `json:"id,required"`
	JSON zeroTrustAccessGroupNewResponseIsDefaultAccessAccessGroupRuleGroupJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessAccessGroupRuleGroup]
type zeroTrustAccessGroupNewResponseIsDefaultAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessGroupNewResponseIsDefaultAccessAzureGroupRule struct {
	AzureAd ZeroTrustAccessGroupNewResponseIsDefaultAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustAccessGroupNewResponseIsDefaultAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessAzureGroupRule]
type zeroTrustAccessGroupNewResponseIsDefaultAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIsDefaultAccessAzureGroupRule) implementsZeroTrustAccessGroupNewResponseIsDefault() {
}

type ZeroTrustAccessGroupNewResponseIsDefaultAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                  `json:"connection_id,required"`
	JSON         zeroTrustAccessGroupNewResponseIsDefaultAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessAzureGroupRuleAzureAdJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessAzureGroupRuleAzureAd]
type zeroTrustAccessGroupNewResponseIsDefaultAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessGroupNewResponseIsDefaultAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustAccessGroupNewResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustAccessGroupNewResponseIsDefaultAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessGitHubOrganizationRule]
type zeroTrustAccessGroupNewResponseIsDefaultAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIsDefaultAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupNewResponseIsDefault() {
}

type ZeroTrustAccessGroupNewResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                     `json:"name,required"`
	JSON zeroTrustAccessGroupNewResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustAccessGroupNewResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessGroupNewResponseIsDefaultAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustAccessGroupNewResponseIsDefaultAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustAccessGroupNewResponseIsDefaultAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessGsuiteGroupRule]
type zeroTrustAccessGroupNewResponseIsDefaultAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIsDefaultAccessGsuiteGroupRule) implementsZeroTrustAccessGroupNewResponseIsDefault() {
}

type ZeroTrustAccessGroupNewResponseIsDefaultAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                  `json:"email,required"`
	JSON  zeroTrustAccessGroupNewResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessGsuiteGroupRuleGsuite]
type zeroTrustAccessGroupNewResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessGroupNewResponseIsDefaultAccessOktaGroupRule struct {
	Okta ZeroTrustAccessGroupNewResponseIsDefaultAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustAccessGroupNewResponseIsDefaultAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessOktaGroupRule]
type zeroTrustAccessGroupNewResponseIsDefaultAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIsDefaultAccessOktaGroupRule) implementsZeroTrustAccessGroupNewResponseIsDefault() {
}

type ZeroTrustAccessGroupNewResponseIsDefaultAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                              `json:"email,required"`
	JSON  zeroTrustAccessGroupNewResponseIsDefaultAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessOktaGroupRuleOktaJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessOktaGroupRuleOkta]
type zeroTrustAccessGroupNewResponseIsDefaultAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessGroupNewResponseIsDefaultAccessSamlGroupRule struct {
	Saml ZeroTrustAccessGroupNewResponseIsDefaultAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustAccessGroupNewResponseIsDefaultAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessSamlGroupRule]
type zeroTrustAccessGroupNewResponseIsDefaultAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIsDefaultAccessSamlGroupRule) implementsZeroTrustAccessGroupNewResponseIsDefault() {
}

type ZeroTrustAccessGroupNewResponseIsDefaultAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                              `json:"attribute_value,required"`
	JSON           zeroTrustAccessGroupNewResponseIsDefaultAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessSamlGroupRuleSamlJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessSamlGroupRuleSaml]
type zeroTrustAccessGroupNewResponseIsDefaultAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustAccessGroupNewResponseIsDefaultAccessServiceTokenRule struct {
	ServiceToken ZeroTrustAccessGroupNewResponseIsDefaultAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustAccessGroupNewResponseIsDefaultAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessServiceTokenRule]
type zeroTrustAccessGroupNewResponseIsDefaultAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIsDefaultAccessServiceTokenRule) implementsZeroTrustAccessGroupNewResponseIsDefault() {
}

type ZeroTrustAccessGroupNewResponseIsDefaultAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                         `json:"token_id,required"`
	JSON    zeroTrustAccessGroupNewResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessServiceTokenRuleServiceToken]
type zeroTrustAccessGroupNewResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustAccessGroupNewResponseIsDefaultAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                `json:"any_valid_service_token,required"`
	JSON                 zeroTrustAccessGroupNewResponseIsDefaultAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessAnyValidServiceTokenRule]
type zeroTrustAccessGroupNewResponseIsDefaultAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIsDefaultAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupNewResponseIsDefault() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessGroupNewResponseIsDefaultAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustAccessGroupNewResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustAccessGroupNewResponseIsDefaultAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessExternalEvaluationRule]
type zeroTrustAccessGroupNewResponseIsDefaultAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIsDefaultAccessExternalEvaluationRule) implementsZeroTrustAccessGroupNewResponseIsDefault() {
}

type ZeroTrustAccessGroupNewResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                     `json:"keys_url,required"`
	JSON    zeroTrustAccessGroupNewResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustAccessGroupNewResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustAccessGroupNewResponseIsDefaultAccessCountryRule struct {
	Geo  ZeroTrustAccessGroupNewResponseIsDefaultAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustAccessGroupNewResponseIsDefaultAccessCountryRuleJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessCountryRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessCountryRule]
type zeroTrustAccessGroupNewResponseIsDefaultAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIsDefaultAccessCountryRule) implementsZeroTrustAccessGroupNewResponseIsDefault() {
}

type ZeroTrustAccessGroupNewResponseIsDefaultAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                           `json:"country_code,required"`
	JSON        zeroTrustAccessGroupNewResponseIsDefaultAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessCountryRuleGeo]
type zeroTrustAccessGroupNewResponseIsDefaultAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustAccessGroupNewResponseIsDefaultAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustAccessGroupNewResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupNewResponseIsDefaultAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessAuthenticationMethodRule]
type zeroTrustAccessGroupNewResponseIsDefaultAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIsDefaultAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupNewResponseIsDefault() {
}

type ZeroTrustAccessGroupNewResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                               `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupNewResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustAccessGroupNewResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessGroupNewResponseIsDefaultAccessDevicePostureRule struct {
	DevicePosture ZeroTrustAccessGroupNewResponseIsDefaultAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustAccessGroupNewResponseIsDefaultAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessDevicePostureRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessDevicePostureRule]
type zeroTrustAccessGroupNewResponseIsDefaultAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseIsDefaultAccessDevicePostureRule) implementsZeroTrustAccessGroupNewResponseIsDefault() {
}

type ZeroTrustAccessGroupNewResponseIsDefaultAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                           `json:"integration_uid,required"`
	JSON           zeroTrustAccessGroupNewResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseIsDefaultAccessDevicePostureRuleDevicePosture]
type zeroTrustAccessGroupNewResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseIsDefaultAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by [ZeroTrustAccessGroupNewResponseRequireAccessEmailRule],
// [ZeroTrustAccessGroupNewResponseRequireAccessEmailListRule],
// [ZeroTrustAccessGroupNewResponseRequireAccessDomainRule],
// [ZeroTrustAccessGroupNewResponseRequireAccessEveryoneRule],
// [ZeroTrustAccessGroupNewResponseRequireAccessIPRule],
// [ZeroTrustAccessGroupNewResponseRequireAccessIPListRule],
// [ZeroTrustAccessGroupNewResponseRequireAccessCertificateRule],
// [ZeroTrustAccessGroupNewResponseRequireAccessAccessGroupRule],
// [ZeroTrustAccessGroupNewResponseRequireAccessAzureGroupRule],
// [ZeroTrustAccessGroupNewResponseRequireAccessGitHubOrganizationRule],
// [ZeroTrustAccessGroupNewResponseRequireAccessGsuiteGroupRule],
// [ZeroTrustAccessGroupNewResponseRequireAccessOktaGroupRule],
// [ZeroTrustAccessGroupNewResponseRequireAccessSamlGroupRule],
// [ZeroTrustAccessGroupNewResponseRequireAccessServiceTokenRule],
// [ZeroTrustAccessGroupNewResponseRequireAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessGroupNewResponseRequireAccessExternalEvaluationRule],
// [ZeroTrustAccessGroupNewResponseRequireAccessCountryRule],
// [ZeroTrustAccessGroupNewResponseRequireAccessAuthenticationMethodRule] or
// [ZeroTrustAccessGroupNewResponseRequireAccessDevicePostureRule].
type ZeroTrustAccessGroupNewResponseRequire interface {
	implementsZeroTrustAccessGroupNewResponseRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustAccessGroupNewResponseRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type ZeroTrustAccessGroupNewResponseRequireAccessEmailRule struct {
	Email ZeroTrustAccessGroupNewResponseRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustAccessGroupNewResponseRequireAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessEmailRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupNewResponseRequireAccessEmailRule]
type zeroTrustAccessGroupNewResponseRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseRequireAccessEmailRule) implementsZeroTrustAccessGroupNewResponseRequire() {
}

type ZeroTrustAccessGroupNewResponseRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                         `json:"email,required" format:"email"`
	JSON  zeroTrustAccessGroupNewResponseRequireAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupNewResponseRequireAccessEmailRuleEmail]
type zeroTrustAccessGroupNewResponseRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustAccessGroupNewResponseRequireAccessEmailListRule struct {
	EmailList ZeroTrustAccessGroupNewResponseRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustAccessGroupNewResponseRequireAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessEmailListRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupNewResponseRequireAccessEmailListRule]
type zeroTrustAccessGroupNewResponseRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseRequireAccessEmailListRule) implementsZeroTrustAccessGroupNewResponseRequire() {
}

type ZeroTrustAccessGroupNewResponseRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                 `json:"id,required"`
	JSON zeroTrustAccessGroupNewResponseRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessEmailListRuleEmailListJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseRequireAccessEmailListRuleEmailList]
type zeroTrustAccessGroupNewResponseRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustAccessGroupNewResponseRequireAccessDomainRule struct {
	EmailDomain ZeroTrustAccessGroupNewResponseRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustAccessGroupNewResponseRequireAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessDomainRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupNewResponseRequireAccessDomainRule]
type zeroTrustAccessGroupNewResponseRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseRequireAccessDomainRule) implementsZeroTrustAccessGroupNewResponseRequire() {
}

type ZeroTrustAccessGroupNewResponseRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                `json:"domain,required"`
	JSON   zeroTrustAccessGroupNewResponseRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessDomainRuleEmailDomainJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseRequireAccessDomainRuleEmailDomain]
type zeroTrustAccessGroupNewResponseRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustAccessGroupNewResponseRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                  `json:"everyone,required"`
	JSON     zeroTrustAccessGroupNewResponseRequireAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessEveryoneRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupNewResponseRequireAccessEveryoneRule]
type zeroTrustAccessGroupNewResponseRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseRequireAccessEveryoneRule) implementsZeroTrustAccessGroupNewResponseRequire() {
}

// Matches an IP address block.
type ZeroTrustAccessGroupNewResponseRequireAccessIPRule struct {
	IP   ZeroTrustAccessGroupNewResponseRequireAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustAccessGroupNewResponseRequireAccessIPRuleJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessIPRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupNewResponseRequireAccessIPRule]
type zeroTrustAccessGroupNewResponseRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseRequireAccessIPRule) implementsZeroTrustAccessGroupNewResponseRequire() {
}

type ZeroTrustAccessGroupNewResponseRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                   `json:"ip,required"`
	JSON zeroTrustAccessGroupNewResponseRequireAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessIPRuleIPJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupNewResponseRequireAccessIPRuleIP]
type zeroTrustAccessGroupNewResponseRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustAccessGroupNewResponseRequireAccessIPListRule struct {
	IPList ZeroTrustAccessGroupNewResponseRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustAccessGroupNewResponseRequireAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessIPListRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupNewResponseRequireAccessIPListRule]
type zeroTrustAccessGroupNewResponseRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseRequireAccessIPListRule) implementsZeroTrustAccessGroupNewResponseRequire() {
}

type ZeroTrustAccessGroupNewResponseRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                           `json:"id,required"`
	JSON zeroTrustAccessGroupNewResponseRequireAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseRequireAccessIPListRuleIPList]
type zeroTrustAccessGroupNewResponseRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustAccessGroupNewResponseRequireAccessCertificateRule struct {
	Certificate interface{}                                                     `json:"certificate,required"`
	JSON        zeroTrustAccessGroupNewResponseRequireAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseRequireAccessCertificateRule]
type zeroTrustAccessGroupNewResponseRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseRequireAccessCertificateRule) implementsZeroTrustAccessGroupNewResponseRequire() {
}

// Matches an Access group.
type ZeroTrustAccessGroupNewResponseRequireAccessAccessGroupRule struct {
	Group ZeroTrustAccessGroupNewResponseRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  zeroTrustAccessGroupNewResponseRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseRequireAccessAccessGroupRule]
type zeroTrustAccessGroupNewResponseRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseRequireAccessAccessGroupRule) implementsZeroTrustAccessGroupNewResponseRequire() {
}

type ZeroTrustAccessGroupNewResponseRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                               `json:"id,required"`
	JSON zeroTrustAccessGroupNewResponseRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseRequireAccessAccessGroupRuleGroup]
type zeroTrustAccessGroupNewResponseRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessGroupNewResponseRequireAccessAzureGroupRule struct {
	AzureAd ZeroTrustAccessGroupNewResponseRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustAccessGroupNewResponseRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupNewResponseRequireAccessAzureGroupRule]
type zeroTrustAccessGroupNewResponseRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseRequireAccessAzureGroupRule) implementsZeroTrustAccessGroupNewResponseRequire() {
}

type ZeroTrustAccessGroupNewResponseRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                `json:"connection_id,required"`
	JSON         zeroTrustAccessGroupNewResponseRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessAzureGroupRuleAzureAdJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseRequireAccessAzureGroupRuleAzureAd]
type zeroTrustAccessGroupNewResponseRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessGroupNewResponseRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustAccessGroupNewResponseRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustAccessGroupNewResponseRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessGitHubOrganizationRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseRequireAccessGitHubOrganizationRule]
type zeroTrustAccessGroupNewResponseRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseRequireAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupNewResponseRequire() {
}

type ZeroTrustAccessGroupNewResponseRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                   `json:"name,required"`
	JSON zeroTrustAccessGroupNewResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseRequireAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustAccessGroupNewResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessGroupNewResponseRequireAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustAccessGroupNewResponseRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustAccessGroupNewResponseRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseRequireAccessGsuiteGroupRule]
type zeroTrustAccessGroupNewResponseRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseRequireAccessGsuiteGroupRule) implementsZeroTrustAccessGroupNewResponseRequire() {
}

type ZeroTrustAccessGroupNewResponseRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                `json:"email,required"`
	JSON  zeroTrustAccessGroupNewResponseRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessGsuiteGroupRuleGsuiteJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseRequireAccessGsuiteGroupRuleGsuite]
type zeroTrustAccessGroupNewResponseRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessGroupNewResponseRequireAccessOktaGroupRule struct {
	Okta ZeroTrustAccessGroupNewResponseRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustAccessGroupNewResponseRequireAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupNewResponseRequireAccessOktaGroupRule]
type zeroTrustAccessGroupNewResponseRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseRequireAccessOktaGroupRule) implementsZeroTrustAccessGroupNewResponseRequire() {
}

type ZeroTrustAccessGroupNewResponseRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                            `json:"email,required"`
	JSON  zeroTrustAccessGroupNewResponseRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessOktaGroupRuleOktaJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseRequireAccessOktaGroupRuleOkta]
type zeroTrustAccessGroupNewResponseRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessGroupNewResponseRequireAccessSamlGroupRule struct {
	Saml ZeroTrustAccessGroupNewResponseRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustAccessGroupNewResponseRequireAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupNewResponseRequireAccessSamlGroupRule]
type zeroTrustAccessGroupNewResponseRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseRequireAccessSamlGroupRule) implementsZeroTrustAccessGroupNewResponseRequire() {
}

type ZeroTrustAccessGroupNewResponseRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                            `json:"attribute_value,required"`
	JSON           zeroTrustAccessGroupNewResponseRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessSamlGroupRuleSamlJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseRequireAccessSamlGroupRuleSaml]
type zeroTrustAccessGroupNewResponseRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustAccessGroupNewResponseRequireAccessServiceTokenRule struct {
	ServiceToken ZeroTrustAccessGroupNewResponseRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustAccessGroupNewResponseRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseRequireAccessServiceTokenRule]
type zeroTrustAccessGroupNewResponseRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseRequireAccessServiceTokenRule) implementsZeroTrustAccessGroupNewResponseRequire() {
}

type ZeroTrustAccessGroupNewResponseRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                       `json:"token_id,required"`
	JSON    zeroTrustAccessGroupNewResponseRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseRequireAccessServiceTokenRuleServiceToken]
type zeroTrustAccessGroupNewResponseRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustAccessGroupNewResponseRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                              `json:"any_valid_service_token,required"`
	JSON                 zeroTrustAccessGroupNewResponseRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseRequireAccessAnyValidServiceTokenRule]
type zeroTrustAccessGroupNewResponseRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseRequireAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupNewResponseRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessGroupNewResponseRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustAccessGroupNewResponseRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustAccessGroupNewResponseRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessExternalEvaluationRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseRequireAccessExternalEvaluationRule]
type zeroTrustAccessGroupNewResponseRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseRequireAccessExternalEvaluationRule) implementsZeroTrustAccessGroupNewResponseRequire() {
}

type ZeroTrustAccessGroupNewResponseRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                   `json:"keys_url,required"`
	JSON    zeroTrustAccessGroupNewResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseRequireAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustAccessGroupNewResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustAccessGroupNewResponseRequireAccessCountryRule struct {
	Geo  ZeroTrustAccessGroupNewResponseRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustAccessGroupNewResponseRequireAccessCountryRuleJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessCountryRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupNewResponseRequireAccessCountryRule]
type zeroTrustAccessGroupNewResponseRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseRequireAccessCountryRule) implementsZeroTrustAccessGroupNewResponseRequire() {
}

type ZeroTrustAccessGroupNewResponseRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                         `json:"country_code,required"`
	JSON        zeroTrustAccessGroupNewResponseRequireAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupNewResponseRequireAccessCountryRuleGeo]
type zeroTrustAccessGroupNewResponseRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustAccessGroupNewResponseRequireAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustAccessGroupNewResponseRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupNewResponseRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseRequireAccessAuthenticationMethodRule]
type zeroTrustAccessGroupNewResponseRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseRequireAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupNewResponseRequire() {
}

type ZeroTrustAccessGroupNewResponseRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                             `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupNewResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseRequireAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustAccessGroupNewResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessGroupNewResponseRequireAccessDevicePostureRule struct {
	DevicePosture ZeroTrustAccessGroupNewResponseRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustAccessGroupNewResponseRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessDevicePostureRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseRequireAccessDevicePostureRule]
type zeroTrustAccessGroupNewResponseRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupNewResponseRequireAccessDevicePostureRule) implementsZeroTrustAccessGroupNewResponseRequire() {
}

type ZeroTrustAccessGroupNewResponseRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                         `json:"integration_uid,required"`
	JSON           zeroTrustAccessGroupNewResponseRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustAccessGroupNewResponseRequireAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupNewResponseRequireAccessDevicePostureRuleDevicePosture]
type zeroTrustAccessGroupNewResponseRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupNewResponseRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupNewResponseRequireAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessGroupUpdateResponse struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []ZeroTrustAccessGroupUpdateResponseExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []ZeroTrustAccessGroupUpdateResponseInclude `json:"include"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	IsDefault []ZeroTrustAccessGroupUpdateResponseIsDefault `json:"is_default"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []ZeroTrustAccessGroupUpdateResponseRequire `json:"require"`
	UpdatedAt time.Time                                   `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessGroupUpdateResponseJSON      `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseJSON contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponse]
type zeroTrustAccessGroupUpdateResponseJSON struct {
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

func (r *ZeroTrustAccessGroupUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by [ZeroTrustAccessGroupUpdateResponseExcludeAccessEmailRule],
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessEmailListRule],
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessDomainRule],
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessEveryoneRule],
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessIPRule],
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessIPListRule],
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessCertificateRule],
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessAccessGroupRule],
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessAzureGroupRule],
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessGitHubOrganizationRule],
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessGsuiteGroupRule],
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessOktaGroupRule],
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessSamlGroupRule],
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessServiceTokenRule],
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessExternalEvaluationRule],
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessCountryRule],
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessAuthenticationMethodRule] or
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessDevicePostureRule].
type ZeroTrustAccessGroupUpdateResponseExclude interface {
	implementsZeroTrustAccessGroupUpdateResponseExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustAccessGroupUpdateResponseExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type ZeroTrustAccessGroupUpdateResponseExcludeAccessEmailRule struct {
	Email ZeroTrustAccessGroupUpdateResponseExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustAccessGroupUpdateResponseExcludeAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessEmailRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessEmailRule]
type zeroTrustAccessGroupUpdateResponseExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseExcludeAccessEmailRule) implementsZeroTrustAccessGroupUpdateResponseExclude() {
}

type ZeroTrustAccessGroupUpdateResponseExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                            `json:"email,required" format:"email"`
	JSON  zeroTrustAccessGroupUpdateResponseExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessEmailRuleEmail]
type zeroTrustAccessGroupUpdateResponseExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustAccessGroupUpdateResponseExcludeAccessEmailListRule struct {
	EmailList ZeroTrustAccessGroupUpdateResponseExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustAccessGroupUpdateResponseExcludeAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessEmailListRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessEmailListRule]
type zeroTrustAccessGroupUpdateResponseExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseExcludeAccessEmailListRule) implementsZeroTrustAccessGroupUpdateResponseExclude() {
}

type ZeroTrustAccessGroupUpdateResponseExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                    `json:"id,required"`
	JSON zeroTrustAccessGroupUpdateResponseExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessEmailListRuleEmailList]
type zeroTrustAccessGroupUpdateResponseExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustAccessGroupUpdateResponseExcludeAccessDomainRule struct {
	EmailDomain ZeroTrustAccessGroupUpdateResponseExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustAccessGroupUpdateResponseExcludeAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessDomainRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessDomainRule]
type zeroTrustAccessGroupUpdateResponseExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseExcludeAccessDomainRule) implementsZeroTrustAccessGroupUpdateResponseExclude() {
}

type ZeroTrustAccessGroupUpdateResponseExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                   `json:"domain,required"`
	JSON   zeroTrustAccessGroupUpdateResponseExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessDomainRuleEmailDomain]
type zeroTrustAccessGroupUpdateResponseExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustAccessGroupUpdateResponseExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                     `json:"everyone,required"`
	JSON     zeroTrustAccessGroupUpdateResponseExcludeAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessEveryoneRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessEveryoneRule]
type zeroTrustAccessGroupUpdateResponseExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseExcludeAccessEveryoneRule) implementsZeroTrustAccessGroupUpdateResponseExclude() {
}

// Matches an IP address block.
type ZeroTrustAccessGroupUpdateResponseExcludeAccessIPRule struct {
	IP   ZeroTrustAccessGroupUpdateResponseExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustAccessGroupUpdateResponseExcludeAccessIPRuleJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessIPRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupUpdateResponseExcludeAccessIPRule]
type zeroTrustAccessGroupUpdateResponseExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseExcludeAccessIPRule) implementsZeroTrustAccessGroupUpdateResponseExclude() {
}

type ZeroTrustAccessGroupUpdateResponseExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                      `json:"ip,required"`
	JSON zeroTrustAccessGroupUpdateResponseExcludeAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessIPRuleIPJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessIPRuleIP]
type zeroTrustAccessGroupUpdateResponseExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustAccessGroupUpdateResponseExcludeAccessIPListRule struct {
	IPList ZeroTrustAccessGroupUpdateResponseExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustAccessGroupUpdateResponseExcludeAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessIPListRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessIPListRule]
type zeroTrustAccessGroupUpdateResponseExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseExcludeAccessIPListRule) implementsZeroTrustAccessGroupUpdateResponseExclude() {
}

type ZeroTrustAccessGroupUpdateResponseExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                              `json:"id,required"`
	JSON zeroTrustAccessGroupUpdateResponseExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessIPListRuleIPList]
type zeroTrustAccessGroupUpdateResponseExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustAccessGroupUpdateResponseExcludeAccessCertificateRule struct {
	Certificate interface{}                                                        `json:"certificate,required"`
	JSON        zeroTrustAccessGroupUpdateResponseExcludeAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessCertificateRule]
type zeroTrustAccessGroupUpdateResponseExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseExcludeAccessCertificateRule) implementsZeroTrustAccessGroupUpdateResponseExclude() {
}

// Matches an Access group.
type ZeroTrustAccessGroupUpdateResponseExcludeAccessAccessGroupRule struct {
	Group ZeroTrustAccessGroupUpdateResponseExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  zeroTrustAccessGroupUpdateResponseExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessAccessGroupRule]
type zeroTrustAccessGroupUpdateResponseExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseExcludeAccessAccessGroupRule) implementsZeroTrustAccessGroupUpdateResponseExclude() {
}

type ZeroTrustAccessGroupUpdateResponseExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                  `json:"id,required"`
	JSON zeroTrustAccessGroupUpdateResponseExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessAccessGroupRuleGroup]
type zeroTrustAccessGroupUpdateResponseExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessGroupUpdateResponseExcludeAccessAzureGroupRule struct {
	AzureAd ZeroTrustAccessGroupUpdateResponseExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustAccessGroupUpdateResponseExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessAzureGroupRule]
type zeroTrustAccessGroupUpdateResponseExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseExcludeAccessAzureGroupRule) implementsZeroTrustAccessGroupUpdateResponseExclude() {
}

type ZeroTrustAccessGroupUpdateResponseExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                   `json:"connection_id,required"`
	JSON         zeroTrustAccessGroupUpdateResponseExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessAzureGroupRuleAzureAd]
type zeroTrustAccessGroupUpdateResponseExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessGroupUpdateResponseExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustAccessGroupUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustAccessGroupUpdateResponseExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessGitHubOrganizationRule]
type zeroTrustAccessGroupUpdateResponseExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseExcludeAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupUpdateResponseExclude() {
}

type ZeroTrustAccessGroupUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                      `json:"name,required"`
	JSON zeroTrustAccessGroupUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustAccessGroupUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessGroupUpdateResponseExcludeAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustAccessGroupUpdateResponseExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustAccessGroupUpdateResponseExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessGsuiteGroupRule]
type zeroTrustAccessGroupUpdateResponseExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseExcludeAccessGsuiteGroupRule) implementsZeroTrustAccessGroupUpdateResponseExclude() {
}

type ZeroTrustAccessGroupUpdateResponseExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                   `json:"email,required"`
	JSON  zeroTrustAccessGroupUpdateResponseExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessGsuiteGroupRuleGsuite]
type zeroTrustAccessGroupUpdateResponseExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessGroupUpdateResponseExcludeAccessOktaGroupRule struct {
	Okta ZeroTrustAccessGroupUpdateResponseExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustAccessGroupUpdateResponseExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessOktaGroupRule]
type zeroTrustAccessGroupUpdateResponseExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseExcludeAccessOktaGroupRule) implementsZeroTrustAccessGroupUpdateResponseExclude() {
}

type ZeroTrustAccessGroupUpdateResponseExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                               `json:"email,required"`
	JSON  zeroTrustAccessGroupUpdateResponseExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessOktaGroupRuleOkta]
type zeroTrustAccessGroupUpdateResponseExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessGroupUpdateResponseExcludeAccessSamlGroupRule struct {
	Saml ZeroTrustAccessGroupUpdateResponseExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustAccessGroupUpdateResponseExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessSamlGroupRule]
type zeroTrustAccessGroupUpdateResponseExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseExcludeAccessSamlGroupRule) implementsZeroTrustAccessGroupUpdateResponseExclude() {
}

type ZeroTrustAccessGroupUpdateResponseExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                               `json:"attribute_value,required"`
	JSON           zeroTrustAccessGroupUpdateResponseExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessSamlGroupRuleSaml]
type zeroTrustAccessGroupUpdateResponseExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustAccessGroupUpdateResponseExcludeAccessServiceTokenRule struct {
	ServiceToken ZeroTrustAccessGroupUpdateResponseExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustAccessGroupUpdateResponseExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessServiceTokenRule]
type zeroTrustAccessGroupUpdateResponseExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseExcludeAccessServiceTokenRule) implementsZeroTrustAccessGroupUpdateResponseExclude() {
}

type ZeroTrustAccessGroupUpdateResponseExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                          `json:"token_id,required"`
	JSON    zeroTrustAccessGroupUpdateResponseExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessServiceTokenRuleServiceToken]
type zeroTrustAccessGroupUpdateResponseExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustAccessGroupUpdateResponseExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                 `json:"any_valid_service_token,required"`
	JSON                 zeroTrustAccessGroupUpdateResponseExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessAnyValidServiceTokenRule]
type zeroTrustAccessGroupUpdateResponseExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseExcludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupUpdateResponseExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessGroupUpdateResponseExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustAccessGroupUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustAccessGroupUpdateResponseExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessExternalEvaluationRule]
type zeroTrustAccessGroupUpdateResponseExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseExcludeAccessExternalEvaluationRule) implementsZeroTrustAccessGroupUpdateResponseExclude() {
}

type ZeroTrustAccessGroupUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                      `json:"keys_url,required"`
	JSON    zeroTrustAccessGroupUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustAccessGroupUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustAccessGroupUpdateResponseExcludeAccessCountryRule struct {
	Geo  ZeroTrustAccessGroupUpdateResponseExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustAccessGroupUpdateResponseExcludeAccessCountryRuleJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessCountryRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessCountryRule]
type zeroTrustAccessGroupUpdateResponseExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseExcludeAccessCountryRule) implementsZeroTrustAccessGroupUpdateResponseExclude() {
}

type ZeroTrustAccessGroupUpdateResponseExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                            `json:"country_code,required"`
	JSON        zeroTrustAccessGroupUpdateResponseExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessCountryRuleGeo]
type zeroTrustAccessGroupUpdateResponseExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustAccessGroupUpdateResponseExcludeAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustAccessGroupUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupUpdateResponseExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessAuthenticationMethodRule]
type zeroTrustAccessGroupUpdateResponseExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseExcludeAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupUpdateResponseExclude() {
}

type ZeroTrustAccessGroupUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustAccessGroupUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessGroupUpdateResponseExcludeAccessDevicePostureRule struct {
	DevicePosture ZeroTrustAccessGroupUpdateResponseExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustAccessGroupUpdateResponseExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessDevicePostureRule]
type zeroTrustAccessGroupUpdateResponseExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseExcludeAccessDevicePostureRule) implementsZeroTrustAccessGroupUpdateResponseExclude() {
}

type ZeroTrustAccessGroupUpdateResponseExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                            `json:"integration_uid,required"`
	JSON           zeroTrustAccessGroupUpdateResponseExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseExcludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseExcludeAccessDevicePostureRuleDevicePosture]
type zeroTrustAccessGroupUpdateResponseExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseExcludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by [ZeroTrustAccessGroupUpdateResponseIncludeAccessEmailRule],
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessEmailListRule],
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessDomainRule],
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessEveryoneRule],
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessIPRule],
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessIPListRule],
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessCertificateRule],
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessAccessGroupRule],
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessAzureGroupRule],
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessGitHubOrganizationRule],
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessGsuiteGroupRule],
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessOktaGroupRule],
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessSamlGroupRule],
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessServiceTokenRule],
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessExternalEvaluationRule],
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessCountryRule],
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessAuthenticationMethodRule] or
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessDevicePostureRule].
type ZeroTrustAccessGroupUpdateResponseInclude interface {
	implementsZeroTrustAccessGroupUpdateResponseInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustAccessGroupUpdateResponseInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type ZeroTrustAccessGroupUpdateResponseIncludeAccessEmailRule struct {
	Email ZeroTrustAccessGroupUpdateResponseIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustAccessGroupUpdateResponseIncludeAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessEmailRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessEmailRule]
type zeroTrustAccessGroupUpdateResponseIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIncludeAccessEmailRule) implementsZeroTrustAccessGroupUpdateResponseInclude() {
}

type ZeroTrustAccessGroupUpdateResponseIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                            `json:"email,required" format:"email"`
	JSON  zeroTrustAccessGroupUpdateResponseIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessEmailRuleEmail]
type zeroTrustAccessGroupUpdateResponseIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustAccessGroupUpdateResponseIncludeAccessEmailListRule struct {
	EmailList ZeroTrustAccessGroupUpdateResponseIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustAccessGroupUpdateResponseIncludeAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessEmailListRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessEmailListRule]
type zeroTrustAccessGroupUpdateResponseIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIncludeAccessEmailListRule) implementsZeroTrustAccessGroupUpdateResponseInclude() {
}

type ZeroTrustAccessGroupUpdateResponseIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                    `json:"id,required"`
	JSON zeroTrustAccessGroupUpdateResponseIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessEmailListRuleEmailList]
type zeroTrustAccessGroupUpdateResponseIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustAccessGroupUpdateResponseIncludeAccessDomainRule struct {
	EmailDomain ZeroTrustAccessGroupUpdateResponseIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustAccessGroupUpdateResponseIncludeAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessDomainRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessDomainRule]
type zeroTrustAccessGroupUpdateResponseIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIncludeAccessDomainRule) implementsZeroTrustAccessGroupUpdateResponseInclude() {
}

type ZeroTrustAccessGroupUpdateResponseIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                   `json:"domain,required"`
	JSON   zeroTrustAccessGroupUpdateResponseIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessDomainRuleEmailDomain]
type zeroTrustAccessGroupUpdateResponseIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustAccessGroupUpdateResponseIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                     `json:"everyone,required"`
	JSON     zeroTrustAccessGroupUpdateResponseIncludeAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessEveryoneRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessEveryoneRule]
type zeroTrustAccessGroupUpdateResponseIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIncludeAccessEveryoneRule) implementsZeroTrustAccessGroupUpdateResponseInclude() {
}

// Matches an IP address block.
type ZeroTrustAccessGroupUpdateResponseIncludeAccessIPRule struct {
	IP   ZeroTrustAccessGroupUpdateResponseIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustAccessGroupUpdateResponseIncludeAccessIPRuleJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessIPRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupUpdateResponseIncludeAccessIPRule]
type zeroTrustAccessGroupUpdateResponseIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIncludeAccessIPRule) implementsZeroTrustAccessGroupUpdateResponseInclude() {
}

type ZeroTrustAccessGroupUpdateResponseIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                      `json:"ip,required"`
	JSON zeroTrustAccessGroupUpdateResponseIncludeAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessIPRuleIPJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessIPRuleIP]
type zeroTrustAccessGroupUpdateResponseIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustAccessGroupUpdateResponseIncludeAccessIPListRule struct {
	IPList ZeroTrustAccessGroupUpdateResponseIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustAccessGroupUpdateResponseIncludeAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessIPListRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessIPListRule]
type zeroTrustAccessGroupUpdateResponseIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIncludeAccessIPListRule) implementsZeroTrustAccessGroupUpdateResponseInclude() {
}

type ZeroTrustAccessGroupUpdateResponseIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                              `json:"id,required"`
	JSON zeroTrustAccessGroupUpdateResponseIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessIPListRuleIPList]
type zeroTrustAccessGroupUpdateResponseIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustAccessGroupUpdateResponseIncludeAccessCertificateRule struct {
	Certificate interface{}                                                        `json:"certificate,required"`
	JSON        zeroTrustAccessGroupUpdateResponseIncludeAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessCertificateRule]
type zeroTrustAccessGroupUpdateResponseIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIncludeAccessCertificateRule) implementsZeroTrustAccessGroupUpdateResponseInclude() {
}

// Matches an Access group.
type ZeroTrustAccessGroupUpdateResponseIncludeAccessAccessGroupRule struct {
	Group ZeroTrustAccessGroupUpdateResponseIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  zeroTrustAccessGroupUpdateResponseIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessAccessGroupRule]
type zeroTrustAccessGroupUpdateResponseIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIncludeAccessAccessGroupRule) implementsZeroTrustAccessGroupUpdateResponseInclude() {
}

type ZeroTrustAccessGroupUpdateResponseIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                  `json:"id,required"`
	JSON zeroTrustAccessGroupUpdateResponseIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessAccessGroupRuleGroup]
type zeroTrustAccessGroupUpdateResponseIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessGroupUpdateResponseIncludeAccessAzureGroupRule struct {
	AzureAd ZeroTrustAccessGroupUpdateResponseIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustAccessGroupUpdateResponseIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessAzureGroupRule]
type zeroTrustAccessGroupUpdateResponseIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIncludeAccessAzureGroupRule) implementsZeroTrustAccessGroupUpdateResponseInclude() {
}

type ZeroTrustAccessGroupUpdateResponseIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                   `json:"connection_id,required"`
	JSON         zeroTrustAccessGroupUpdateResponseIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessAzureGroupRuleAzureAd]
type zeroTrustAccessGroupUpdateResponseIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessGroupUpdateResponseIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustAccessGroupUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustAccessGroupUpdateResponseIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessGitHubOrganizationRule]
type zeroTrustAccessGroupUpdateResponseIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIncludeAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupUpdateResponseInclude() {
}

type ZeroTrustAccessGroupUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                      `json:"name,required"`
	JSON zeroTrustAccessGroupUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustAccessGroupUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessGroupUpdateResponseIncludeAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustAccessGroupUpdateResponseIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustAccessGroupUpdateResponseIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessGsuiteGroupRule]
type zeroTrustAccessGroupUpdateResponseIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIncludeAccessGsuiteGroupRule) implementsZeroTrustAccessGroupUpdateResponseInclude() {
}

type ZeroTrustAccessGroupUpdateResponseIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                   `json:"email,required"`
	JSON  zeroTrustAccessGroupUpdateResponseIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessGsuiteGroupRuleGsuite]
type zeroTrustAccessGroupUpdateResponseIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessGroupUpdateResponseIncludeAccessOktaGroupRule struct {
	Okta ZeroTrustAccessGroupUpdateResponseIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustAccessGroupUpdateResponseIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessOktaGroupRule]
type zeroTrustAccessGroupUpdateResponseIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIncludeAccessOktaGroupRule) implementsZeroTrustAccessGroupUpdateResponseInclude() {
}

type ZeroTrustAccessGroupUpdateResponseIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                               `json:"email,required"`
	JSON  zeroTrustAccessGroupUpdateResponseIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessOktaGroupRuleOkta]
type zeroTrustAccessGroupUpdateResponseIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessGroupUpdateResponseIncludeAccessSamlGroupRule struct {
	Saml ZeroTrustAccessGroupUpdateResponseIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustAccessGroupUpdateResponseIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessSamlGroupRule]
type zeroTrustAccessGroupUpdateResponseIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIncludeAccessSamlGroupRule) implementsZeroTrustAccessGroupUpdateResponseInclude() {
}

type ZeroTrustAccessGroupUpdateResponseIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                               `json:"attribute_value,required"`
	JSON           zeroTrustAccessGroupUpdateResponseIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessSamlGroupRuleSaml]
type zeroTrustAccessGroupUpdateResponseIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustAccessGroupUpdateResponseIncludeAccessServiceTokenRule struct {
	ServiceToken ZeroTrustAccessGroupUpdateResponseIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustAccessGroupUpdateResponseIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessServiceTokenRule]
type zeroTrustAccessGroupUpdateResponseIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIncludeAccessServiceTokenRule) implementsZeroTrustAccessGroupUpdateResponseInclude() {
}

type ZeroTrustAccessGroupUpdateResponseIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                          `json:"token_id,required"`
	JSON    zeroTrustAccessGroupUpdateResponseIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessServiceTokenRuleServiceToken]
type zeroTrustAccessGroupUpdateResponseIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustAccessGroupUpdateResponseIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                 `json:"any_valid_service_token,required"`
	JSON                 zeroTrustAccessGroupUpdateResponseIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessAnyValidServiceTokenRule]
type zeroTrustAccessGroupUpdateResponseIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIncludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupUpdateResponseInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessGroupUpdateResponseIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustAccessGroupUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustAccessGroupUpdateResponseIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessExternalEvaluationRule]
type zeroTrustAccessGroupUpdateResponseIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIncludeAccessExternalEvaluationRule) implementsZeroTrustAccessGroupUpdateResponseInclude() {
}

type ZeroTrustAccessGroupUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                      `json:"keys_url,required"`
	JSON    zeroTrustAccessGroupUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustAccessGroupUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustAccessGroupUpdateResponseIncludeAccessCountryRule struct {
	Geo  ZeroTrustAccessGroupUpdateResponseIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustAccessGroupUpdateResponseIncludeAccessCountryRuleJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessCountryRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessCountryRule]
type zeroTrustAccessGroupUpdateResponseIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIncludeAccessCountryRule) implementsZeroTrustAccessGroupUpdateResponseInclude() {
}

type ZeroTrustAccessGroupUpdateResponseIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                            `json:"country_code,required"`
	JSON        zeroTrustAccessGroupUpdateResponseIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessCountryRuleGeo]
type zeroTrustAccessGroupUpdateResponseIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustAccessGroupUpdateResponseIncludeAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustAccessGroupUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupUpdateResponseIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessAuthenticationMethodRule]
type zeroTrustAccessGroupUpdateResponseIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIncludeAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupUpdateResponseInclude() {
}

type ZeroTrustAccessGroupUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustAccessGroupUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessGroupUpdateResponseIncludeAccessDevicePostureRule struct {
	DevicePosture ZeroTrustAccessGroupUpdateResponseIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustAccessGroupUpdateResponseIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessDevicePostureRule]
type zeroTrustAccessGroupUpdateResponseIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIncludeAccessDevicePostureRule) implementsZeroTrustAccessGroupUpdateResponseInclude() {
}

type ZeroTrustAccessGroupUpdateResponseIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                            `json:"integration_uid,required"`
	JSON           zeroTrustAccessGroupUpdateResponseIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIncludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIncludeAccessDevicePostureRuleDevicePosture]
type zeroTrustAccessGroupUpdateResponseIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIncludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailRule],
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailListRule],
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessDomainRule],
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessEveryoneRule],
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessIPRule],
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessIPListRule],
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessCertificateRule],
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAccessGroupRule],
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAzureGroupRule],
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRule],
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessGsuiteGroupRule],
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessOktaGroupRule],
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessSamlGroupRule],
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessServiceTokenRule],
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessExternalEvaluationRule],
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessCountryRule],
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRule] or
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessDevicePostureRule].
type ZeroTrustAccessGroupUpdateResponseIsDefault interface {
	implementsZeroTrustAccessGroupUpdateResponseIsDefault()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustAccessGroupUpdateResponseIsDefault)(nil)).Elem(), "")
}

// Matches a specific email.
type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailRule struct {
	Email ZeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailRule]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailRule) implementsZeroTrustAccessGroupUpdateResponseIsDefault() {
}

type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                              `json:"email,required" format:"email"`
	JSON  zeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailRuleEmail]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailListRule struct {
	EmailList ZeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailListRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailListRule]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailListRule) implementsZeroTrustAccessGroupUpdateResponseIsDefault() {
}

type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                      `json:"id,required"`
	JSON zeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailListRuleEmailList]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessDomainRule struct {
	EmailDomain ZeroTrustAccessGroupUpdateResponseIsDefaultAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustAccessGroupUpdateResponseIsDefaultAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessDomainRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessDomainRule]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIsDefaultAccessDomainRule) implementsZeroTrustAccessGroupUpdateResponseIsDefault() {
}

type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                     `json:"domain,required"`
	JSON   zeroTrustAccessGroupUpdateResponseIsDefaultAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessDomainRuleEmailDomain]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                       `json:"everyone,required"`
	JSON     zeroTrustAccessGroupUpdateResponseIsDefaultAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessEveryoneRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessEveryoneRule]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIsDefaultAccessEveryoneRule) implementsZeroTrustAccessGroupUpdateResponseIsDefault() {
}

// Matches an IP address block.
type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessIPRule struct {
	IP   ZeroTrustAccessGroupUpdateResponseIsDefaultAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustAccessGroupUpdateResponseIsDefaultAccessIPRuleJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessIPRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessIPRule]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIsDefaultAccessIPRule) implementsZeroTrustAccessGroupUpdateResponseIsDefault() {
}

type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                        `json:"ip,required"`
	JSON zeroTrustAccessGroupUpdateResponseIsDefaultAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessIPRuleIPJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessIPRuleIP]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessIPListRule struct {
	IPList ZeroTrustAccessGroupUpdateResponseIsDefaultAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustAccessGroupUpdateResponseIsDefaultAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessIPListRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessIPListRule]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIsDefaultAccessIPListRule) implementsZeroTrustAccessGroupUpdateResponseIsDefault() {
}

type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                                `json:"id,required"`
	JSON zeroTrustAccessGroupUpdateResponseIsDefaultAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessIPListRuleIPListJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessIPListRuleIPList]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessCertificateRule struct {
	Certificate interface{}                                                          `json:"certificate,required"`
	JSON        zeroTrustAccessGroupUpdateResponseIsDefaultAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessCertificateRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessCertificateRule]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIsDefaultAccessCertificateRule) implementsZeroTrustAccessGroupUpdateResponseIsDefault() {
}

// Matches an Access group.
type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAccessGroupRule struct {
	Group ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  zeroTrustAccessGroupUpdateResponseIsDefaultAccessAccessGroupRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessAccessGroupRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAccessGroupRule]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAccessGroupRule) implementsZeroTrustAccessGroupUpdateResponseIsDefault() {
}

type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                    `json:"id,required"`
	JSON zeroTrustAccessGroupUpdateResponseIsDefaultAccessAccessGroupRuleGroupJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessAccessGroupRuleGroupJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAccessGroupRuleGroup]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAzureGroupRule struct {
	AzureAd ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustAccessGroupUpdateResponseIsDefaultAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAzureGroupRule]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAzureGroupRule) implementsZeroTrustAccessGroupUpdateResponseIsDefault() {
}

type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                     `json:"connection_id,required"`
	JSON         zeroTrustAccessGroupUpdateResponseIsDefaultAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAzureGroupRuleAzureAd]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustAccessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustAccessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRule]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupUpdateResponseIsDefault() {
}

type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                        `json:"name,required"`
	JSON zeroTrustAccessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustAccessGroupUpdateResponseIsDefaultAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustAccessGroupUpdateResponseIsDefaultAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessGsuiteGroupRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessGsuiteGroupRule]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIsDefaultAccessGsuiteGroupRule) implementsZeroTrustAccessGroupUpdateResponseIsDefault() {
}

type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                     `json:"email,required"`
	JSON  zeroTrustAccessGroupUpdateResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessGsuiteGroupRuleGsuite]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessOktaGroupRule struct {
	Okta ZeroTrustAccessGroupUpdateResponseIsDefaultAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustAccessGroupUpdateResponseIsDefaultAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessOktaGroupRule]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIsDefaultAccessOktaGroupRule) implementsZeroTrustAccessGroupUpdateResponseIsDefault() {
}

type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                                 `json:"email,required"`
	JSON  zeroTrustAccessGroupUpdateResponseIsDefaultAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessOktaGroupRuleOkta]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessSamlGroupRule struct {
	Saml ZeroTrustAccessGroupUpdateResponseIsDefaultAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustAccessGroupUpdateResponseIsDefaultAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessSamlGroupRule]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIsDefaultAccessSamlGroupRule) implementsZeroTrustAccessGroupUpdateResponseIsDefault() {
}

type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                                 `json:"attribute_value,required"`
	JSON           zeroTrustAccessGroupUpdateResponseIsDefaultAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessSamlGroupRuleSaml]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessServiceTokenRule struct {
	ServiceToken ZeroTrustAccessGroupUpdateResponseIsDefaultAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustAccessGroupUpdateResponseIsDefaultAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessServiceTokenRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessServiceTokenRule]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIsDefaultAccessServiceTokenRule) implementsZeroTrustAccessGroupUpdateResponseIsDefault() {
}

type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                            `json:"token_id,required"`
	JSON    zeroTrustAccessGroupUpdateResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessServiceTokenRuleServiceToken]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                   `json:"any_valid_service_token,required"`
	JSON                 zeroTrustAccessGroupUpdateResponseIsDefaultAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAnyValidServiceTokenRule]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupUpdateResponseIsDefault() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustAccessGroupUpdateResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustAccessGroupUpdateResponseIsDefaultAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessExternalEvaluationRule]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIsDefaultAccessExternalEvaluationRule) implementsZeroTrustAccessGroupUpdateResponseIsDefault() {
}

type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                        `json:"keys_url,required"`
	JSON    zeroTrustAccessGroupUpdateResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessCountryRule struct {
	Geo  ZeroTrustAccessGroupUpdateResponseIsDefaultAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustAccessGroupUpdateResponseIsDefaultAccessCountryRuleJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessCountryRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessCountryRule]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIsDefaultAccessCountryRule) implementsZeroTrustAccessGroupUpdateResponseIsDefault() {
}

type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                              `json:"country_code,required"`
	JSON        zeroTrustAccessGroupUpdateResponseIsDefaultAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessCountryRuleGeo]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRule]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupUpdateResponseIsDefault() {
}

type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                  `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessDevicePostureRule struct {
	DevicePosture ZeroTrustAccessGroupUpdateResponseIsDefaultAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustAccessGroupUpdateResponseIsDefaultAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessDevicePostureRule]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseIsDefaultAccessDevicePostureRule) implementsZeroTrustAccessGroupUpdateResponseIsDefault() {
}

type ZeroTrustAccessGroupUpdateResponseIsDefaultAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                              `json:"integration_uid,required"`
	JSON           zeroTrustAccessGroupUpdateResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseIsDefaultAccessDevicePostureRuleDevicePosture]
type zeroTrustAccessGroupUpdateResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseIsDefaultAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by [ZeroTrustAccessGroupUpdateResponseRequireAccessEmailRule],
// [ZeroTrustAccessGroupUpdateResponseRequireAccessEmailListRule],
// [ZeroTrustAccessGroupUpdateResponseRequireAccessDomainRule],
// [ZeroTrustAccessGroupUpdateResponseRequireAccessEveryoneRule],
// [ZeroTrustAccessGroupUpdateResponseRequireAccessIPRule],
// [ZeroTrustAccessGroupUpdateResponseRequireAccessIPListRule],
// [ZeroTrustAccessGroupUpdateResponseRequireAccessCertificateRule],
// [ZeroTrustAccessGroupUpdateResponseRequireAccessAccessGroupRule],
// [ZeroTrustAccessGroupUpdateResponseRequireAccessAzureGroupRule],
// [ZeroTrustAccessGroupUpdateResponseRequireAccessGitHubOrganizationRule],
// [ZeroTrustAccessGroupUpdateResponseRequireAccessGsuiteGroupRule],
// [ZeroTrustAccessGroupUpdateResponseRequireAccessOktaGroupRule],
// [ZeroTrustAccessGroupUpdateResponseRequireAccessSamlGroupRule],
// [ZeroTrustAccessGroupUpdateResponseRequireAccessServiceTokenRule],
// [ZeroTrustAccessGroupUpdateResponseRequireAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessGroupUpdateResponseRequireAccessExternalEvaluationRule],
// [ZeroTrustAccessGroupUpdateResponseRequireAccessCountryRule],
// [ZeroTrustAccessGroupUpdateResponseRequireAccessAuthenticationMethodRule] or
// [ZeroTrustAccessGroupUpdateResponseRequireAccessDevicePostureRule].
type ZeroTrustAccessGroupUpdateResponseRequire interface {
	implementsZeroTrustAccessGroupUpdateResponseRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustAccessGroupUpdateResponseRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type ZeroTrustAccessGroupUpdateResponseRequireAccessEmailRule struct {
	Email ZeroTrustAccessGroupUpdateResponseRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustAccessGroupUpdateResponseRequireAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessEmailRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessEmailRule]
type zeroTrustAccessGroupUpdateResponseRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseRequireAccessEmailRule) implementsZeroTrustAccessGroupUpdateResponseRequire() {
}

type ZeroTrustAccessGroupUpdateResponseRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                            `json:"email,required" format:"email"`
	JSON  zeroTrustAccessGroupUpdateResponseRequireAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessEmailRuleEmail]
type zeroTrustAccessGroupUpdateResponseRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustAccessGroupUpdateResponseRequireAccessEmailListRule struct {
	EmailList ZeroTrustAccessGroupUpdateResponseRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustAccessGroupUpdateResponseRequireAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessEmailListRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessEmailListRule]
type zeroTrustAccessGroupUpdateResponseRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseRequireAccessEmailListRule) implementsZeroTrustAccessGroupUpdateResponseRequire() {
}

type ZeroTrustAccessGroupUpdateResponseRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                    `json:"id,required"`
	JSON zeroTrustAccessGroupUpdateResponseRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessEmailListRuleEmailList]
type zeroTrustAccessGroupUpdateResponseRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustAccessGroupUpdateResponseRequireAccessDomainRule struct {
	EmailDomain ZeroTrustAccessGroupUpdateResponseRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustAccessGroupUpdateResponseRequireAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessDomainRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessDomainRule]
type zeroTrustAccessGroupUpdateResponseRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseRequireAccessDomainRule) implementsZeroTrustAccessGroupUpdateResponseRequire() {
}

type ZeroTrustAccessGroupUpdateResponseRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                   `json:"domain,required"`
	JSON   zeroTrustAccessGroupUpdateResponseRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessDomainRuleEmailDomain]
type zeroTrustAccessGroupUpdateResponseRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustAccessGroupUpdateResponseRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                     `json:"everyone,required"`
	JSON     zeroTrustAccessGroupUpdateResponseRequireAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessEveryoneRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessEveryoneRule]
type zeroTrustAccessGroupUpdateResponseRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseRequireAccessEveryoneRule) implementsZeroTrustAccessGroupUpdateResponseRequire() {
}

// Matches an IP address block.
type ZeroTrustAccessGroupUpdateResponseRequireAccessIPRule struct {
	IP   ZeroTrustAccessGroupUpdateResponseRequireAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustAccessGroupUpdateResponseRequireAccessIPRuleJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessIPRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupUpdateResponseRequireAccessIPRule]
type zeroTrustAccessGroupUpdateResponseRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseRequireAccessIPRule) implementsZeroTrustAccessGroupUpdateResponseRequire() {
}

type ZeroTrustAccessGroupUpdateResponseRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                      `json:"ip,required"`
	JSON zeroTrustAccessGroupUpdateResponseRequireAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessIPRuleIPJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessIPRuleIP]
type zeroTrustAccessGroupUpdateResponseRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustAccessGroupUpdateResponseRequireAccessIPListRule struct {
	IPList ZeroTrustAccessGroupUpdateResponseRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustAccessGroupUpdateResponseRequireAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessIPListRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessIPListRule]
type zeroTrustAccessGroupUpdateResponseRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseRequireAccessIPListRule) implementsZeroTrustAccessGroupUpdateResponseRequire() {
}

type ZeroTrustAccessGroupUpdateResponseRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                              `json:"id,required"`
	JSON zeroTrustAccessGroupUpdateResponseRequireAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessIPListRuleIPList]
type zeroTrustAccessGroupUpdateResponseRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustAccessGroupUpdateResponseRequireAccessCertificateRule struct {
	Certificate interface{}                                                        `json:"certificate,required"`
	JSON        zeroTrustAccessGroupUpdateResponseRequireAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessCertificateRule]
type zeroTrustAccessGroupUpdateResponseRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseRequireAccessCertificateRule) implementsZeroTrustAccessGroupUpdateResponseRequire() {
}

// Matches an Access group.
type ZeroTrustAccessGroupUpdateResponseRequireAccessAccessGroupRule struct {
	Group ZeroTrustAccessGroupUpdateResponseRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  zeroTrustAccessGroupUpdateResponseRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessAccessGroupRule]
type zeroTrustAccessGroupUpdateResponseRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseRequireAccessAccessGroupRule) implementsZeroTrustAccessGroupUpdateResponseRequire() {
}

type ZeroTrustAccessGroupUpdateResponseRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                  `json:"id,required"`
	JSON zeroTrustAccessGroupUpdateResponseRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessAccessGroupRuleGroup]
type zeroTrustAccessGroupUpdateResponseRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessGroupUpdateResponseRequireAccessAzureGroupRule struct {
	AzureAd ZeroTrustAccessGroupUpdateResponseRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustAccessGroupUpdateResponseRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessAzureGroupRule]
type zeroTrustAccessGroupUpdateResponseRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseRequireAccessAzureGroupRule) implementsZeroTrustAccessGroupUpdateResponseRequire() {
}

type ZeroTrustAccessGroupUpdateResponseRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                   `json:"connection_id,required"`
	JSON         zeroTrustAccessGroupUpdateResponseRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessAzureGroupRuleAzureAd]
type zeroTrustAccessGroupUpdateResponseRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessGroupUpdateResponseRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustAccessGroupUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustAccessGroupUpdateResponseRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessGitHubOrganizationRule]
type zeroTrustAccessGroupUpdateResponseRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseRequireAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupUpdateResponseRequire() {
}

type ZeroTrustAccessGroupUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                      `json:"name,required"`
	JSON zeroTrustAccessGroupUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustAccessGroupUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessGroupUpdateResponseRequireAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustAccessGroupUpdateResponseRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustAccessGroupUpdateResponseRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessGsuiteGroupRule]
type zeroTrustAccessGroupUpdateResponseRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseRequireAccessGsuiteGroupRule) implementsZeroTrustAccessGroupUpdateResponseRequire() {
}

type ZeroTrustAccessGroupUpdateResponseRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                   `json:"email,required"`
	JSON  zeroTrustAccessGroupUpdateResponseRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessGsuiteGroupRuleGsuite]
type zeroTrustAccessGroupUpdateResponseRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessGroupUpdateResponseRequireAccessOktaGroupRule struct {
	Okta ZeroTrustAccessGroupUpdateResponseRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustAccessGroupUpdateResponseRequireAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessOktaGroupRule]
type zeroTrustAccessGroupUpdateResponseRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseRequireAccessOktaGroupRule) implementsZeroTrustAccessGroupUpdateResponseRequire() {
}

type ZeroTrustAccessGroupUpdateResponseRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                               `json:"email,required"`
	JSON  zeroTrustAccessGroupUpdateResponseRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessOktaGroupRuleOkta]
type zeroTrustAccessGroupUpdateResponseRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessGroupUpdateResponseRequireAccessSamlGroupRule struct {
	Saml ZeroTrustAccessGroupUpdateResponseRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustAccessGroupUpdateResponseRequireAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessSamlGroupRule]
type zeroTrustAccessGroupUpdateResponseRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseRequireAccessSamlGroupRule) implementsZeroTrustAccessGroupUpdateResponseRequire() {
}

type ZeroTrustAccessGroupUpdateResponseRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                               `json:"attribute_value,required"`
	JSON           zeroTrustAccessGroupUpdateResponseRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessSamlGroupRuleSaml]
type zeroTrustAccessGroupUpdateResponseRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustAccessGroupUpdateResponseRequireAccessServiceTokenRule struct {
	ServiceToken ZeroTrustAccessGroupUpdateResponseRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustAccessGroupUpdateResponseRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessServiceTokenRule]
type zeroTrustAccessGroupUpdateResponseRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseRequireAccessServiceTokenRule) implementsZeroTrustAccessGroupUpdateResponseRequire() {
}

type ZeroTrustAccessGroupUpdateResponseRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                          `json:"token_id,required"`
	JSON    zeroTrustAccessGroupUpdateResponseRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessServiceTokenRuleServiceToken]
type zeroTrustAccessGroupUpdateResponseRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustAccessGroupUpdateResponseRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                 `json:"any_valid_service_token,required"`
	JSON                 zeroTrustAccessGroupUpdateResponseRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessAnyValidServiceTokenRule]
type zeroTrustAccessGroupUpdateResponseRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseRequireAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupUpdateResponseRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessGroupUpdateResponseRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustAccessGroupUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustAccessGroupUpdateResponseRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessExternalEvaluationRule]
type zeroTrustAccessGroupUpdateResponseRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseRequireAccessExternalEvaluationRule) implementsZeroTrustAccessGroupUpdateResponseRequire() {
}

type ZeroTrustAccessGroupUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                      `json:"keys_url,required"`
	JSON    zeroTrustAccessGroupUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustAccessGroupUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustAccessGroupUpdateResponseRequireAccessCountryRule struct {
	Geo  ZeroTrustAccessGroupUpdateResponseRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustAccessGroupUpdateResponseRequireAccessCountryRuleJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessCountryRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessCountryRule]
type zeroTrustAccessGroupUpdateResponseRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseRequireAccessCountryRule) implementsZeroTrustAccessGroupUpdateResponseRequire() {
}

type ZeroTrustAccessGroupUpdateResponseRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                            `json:"country_code,required"`
	JSON        zeroTrustAccessGroupUpdateResponseRequireAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessCountryRuleGeo]
type zeroTrustAccessGroupUpdateResponseRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustAccessGroupUpdateResponseRequireAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustAccessGroupUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupUpdateResponseRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessAuthenticationMethodRule]
type zeroTrustAccessGroupUpdateResponseRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseRequireAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupUpdateResponseRequire() {
}

type ZeroTrustAccessGroupUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustAccessGroupUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessGroupUpdateResponseRequireAccessDevicePostureRule struct {
	DevicePosture ZeroTrustAccessGroupUpdateResponseRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustAccessGroupUpdateResponseRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessDevicePostureRule]
type zeroTrustAccessGroupUpdateResponseRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupUpdateResponseRequireAccessDevicePostureRule) implementsZeroTrustAccessGroupUpdateResponseRequire() {
}

type ZeroTrustAccessGroupUpdateResponseRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                            `json:"integration_uid,required"`
	JSON           zeroTrustAccessGroupUpdateResponseRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustAccessGroupUpdateResponseRequireAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupUpdateResponseRequireAccessDevicePostureRuleDevicePosture]
type zeroTrustAccessGroupUpdateResponseRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupUpdateResponseRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupUpdateResponseRequireAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessGroupListResponse struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []ZeroTrustAccessGroupListResponseExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []ZeroTrustAccessGroupListResponseInclude `json:"include"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	IsDefault []ZeroTrustAccessGroupListResponseIsDefault `json:"is_default"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []ZeroTrustAccessGroupListResponseRequire `json:"require"`
	UpdatedAt time.Time                                 `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessGroupListResponseJSON      `json:"-"`
}

// zeroTrustAccessGroupListResponseJSON contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponse]
type zeroTrustAccessGroupListResponseJSON struct {
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

func (r *ZeroTrustAccessGroupListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by [ZeroTrustAccessGroupListResponseExcludeAccessEmailRule],
// [ZeroTrustAccessGroupListResponseExcludeAccessEmailListRule],
// [ZeroTrustAccessGroupListResponseExcludeAccessDomainRule],
// [ZeroTrustAccessGroupListResponseExcludeAccessEveryoneRule],
// [ZeroTrustAccessGroupListResponseExcludeAccessIPRule],
// [ZeroTrustAccessGroupListResponseExcludeAccessIPListRule],
// [ZeroTrustAccessGroupListResponseExcludeAccessCertificateRule],
// [ZeroTrustAccessGroupListResponseExcludeAccessAccessGroupRule],
// [ZeroTrustAccessGroupListResponseExcludeAccessAzureGroupRule],
// [ZeroTrustAccessGroupListResponseExcludeAccessGitHubOrganizationRule],
// [ZeroTrustAccessGroupListResponseExcludeAccessGsuiteGroupRule],
// [ZeroTrustAccessGroupListResponseExcludeAccessOktaGroupRule],
// [ZeroTrustAccessGroupListResponseExcludeAccessSamlGroupRule],
// [ZeroTrustAccessGroupListResponseExcludeAccessServiceTokenRule],
// [ZeroTrustAccessGroupListResponseExcludeAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessGroupListResponseExcludeAccessExternalEvaluationRule],
// [ZeroTrustAccessGroupListResponseExcludeAccessCountryRule],
// [ZeroTrustAccessGroupListResponseExcludeAccessAuthenticationMethodRule] or
// [ZeroTrustAccessGroupListResponseExcludeAccessDevicePostureRule].
type ZeroTrustAccessGroupListResponseExclude interface {
	implementsZeroTrustAccessGroupListResponseExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustAccessGroupListResponseExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type ZeroTrustAccessGroupListResponseExcludeAccessEmailRule struct {
	Email ZeroTrustAccessGroupListResponseExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustAccessGroupListResponseExcludeAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessEmailRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupListResponseExcludeAccessEmailRule]
type zeroTrustAccessGroupListResponseExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseExcludeAccessEmailRule) implementsZeroTrustAccessGroupListResponseExclude() {
}

type ZeroTrustAccessGroupListResponseExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                          `json:"email,required" format:"email"`
	JSON  zeroTrustAccessGroupListResponseExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessEmailRuleEmail]
type zeroTrustAccessGroupListResponseExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustAccessGroupListResponseExcludeAccessEmailListRule struct {
	EmailList ZeroTrustAccessGroupListResponseExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustAccessGroupListResponseExcludeAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessEmailListRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessEmailListRule]
type zeroTrustAccessGroupListResponseExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseExcludeAccessEmailListRule) implementsZeroTrustAccessGroupListResponseExclude() {
}

type ZeroTrustAccessGroupListResponseExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                  `json:"id,required"`
	JSON zeroTrustAccessGroupListResponseExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessEmailListRuleEmailListJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessEmailListRuleEmailList]
type zeroTrustAccessGroupListResponseExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustAccessGroupListResponseExcludeAccessDomainRule struct {
	EmailDomain ZeroTrustAccessGroupListResponseExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustAccessGroupListResponseExcludeAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessDomainRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessDomainRule]
type zeroTrustAccessGroupListResponseExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseExcludeAccessDomainRule) implementsZeroTrustAccessGroupListResponseExclude() {
}

type ZeroTrustAccessGroupListResponseExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                 `json:"domain,required"`
	JSON   zeroTrustAccessGroupListResponseExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessDomainRuleEmailDomainJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessDomainRuleEmailDomain]
type zeroTrustAccessGroupListResponseExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustAccessGroupListResponseExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                   `json:"everyone,required"`
	JSON     zeroTrustAccessGroupListResponseExcludeAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessEveryoneRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessEveryoneRule]
type zeroTrustAccessGroupListResponseExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseExcludeAccessEveryoneRule) implementsZeroTrustAccessGroupListResponseExclude() {
}

// Matches an IP address block.
type ZeroTrustAccessGroupListResponseExcludeAccessIPRule struct {
	IP   ZeroTrustAccessGroupListResponseExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustAccessGroupListResponseExcludeAccessIPRuleJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessIPRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupListResponseExcludeAccessIPRule]
type zeroTrustAccessGroupListResponseExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseExcludeAccessIPRule) implementsZeroTrustAccessGroupListResponseExclude() {
}

type ZeroTrustAccessGroupListResponseExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                    `json:"ip,required"`
	JSON zeroTrustAccessGroupListResponseExcludeAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessIPRuleIPJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupListResponseExcludeAccessIPRuleIP]
type zeroTrustAccessGroupListResponseExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustAccessGroupListResponseExcludeAccessIPListRule struct {
	IPList ZeroTrustAccessGroupListResponseExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustAccessGroupListResponseExcludeAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessIPListRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessIPListRule]
type zeroTrustAccessGroupListResponseExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseExcludeAccessIPListRule) implementsZeroTrustAccessGroupListResponseExclude() {
}

type ZeroTrustAccessGroupListResponseExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                            `json:"id,required"`
	JSON zeroTrustAccessGroupListResponseExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessIPListRuleIPList]
type zeroTrustAccessGroupListResponseExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustAccessGroupListResponseExcludeAccessCertificateRule struct {
	Certificate interface{}                                                      `json:"certificate,required"`
	JSON        zeroTrustAccessGroupListResponseExcludeAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessCertificateRule]
type zeroTrustAccessGroupListResponseExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseExcludeAccessCertificateRule) implementsZeroTrustAccessGroupListResponseExclude() {
}

// Matches an Access group.
type ZeroTrustAccessGroupListResponseExcludeAccessAccessGroupRule struct {
	Group ZeroTrustAccessGroupListResponseExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  zeroTrustAccessGroupListResponseExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessAccessGroupRule]
type zeroTrustAccessGroupListResponseExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseExcludeAccessAccessGroupRule) implementsZeroTrustAccessGroupListResponseExclude() {
}

type ZeroTrustAccessGroupListResponseExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                `json:"id,required"`
	JSON zeroTrustAccessGroupListResponseExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessAccessGroupRuleGroup]
type zeroTrustAccessGroupListResponseExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessGroupListResponseExcludeAccessAzureGroupRule struct {
	AzureAd ZeroTrustAccessGroupListResponseExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustAccessGroupListResponseExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessAzureGroupRule]
type zeroTrustAccessGroupListResponseExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseExcludeAccessAzureGroupRule) implementsZeroTrustAccessGroupListResponseExclude() {
}

type ZeroTrustAccessGroupListResponseExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                 `json:"connection_id,required"`
	JSON         zeroTrustAccessGroupListResponseExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessAzureGroupRuleAzureAdJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessAzureGroupRuleAzureAd]
type zeroTrustAccessGroupListResponseExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessGroupListResponseExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustAccessGroupListResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustAccessGroupListResponseExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessGitHubOrganizationRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessGitHubOrganizationRule]
type zeroTrustAccessGroupListResponseExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseExcludeAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupListResponseExclude() {
}

type ZeroTrustAccessGroupListResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                    `json:"name,required"`
	JSON zeroTrustAccessGroupListResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustAccessGroupListResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessGroupListResponseExcludeAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustAccessGroupListResponseExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustAccessGroupListResponseExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessGsuiteGroupRule]
type zeroTrustAccessGroupListResponseExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseExcludeAccessGsuiteGroupRule) implementsZeroTrustAccessGroupListResponseExclude() {
}

type ZeroTrustAccessGroupListResponseExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                 `json:"email,required"`
	JSON  zeroTrustAccessGroupListResponseExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessGsuiteGroupRuleGsuiteJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessGsuiteGroupRuleGsuite]
type zeroTrustAccessGroupListResponseExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessGroupListResponseExcludeAccessOktaGroupRule struct {
	Okta ZeroTrustAccessGroupListResponseExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustAccessGroupListResponseExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessOktaGroupRule]
type zeroTrustAccessGroupListResponseExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseExcludeAccessOktaGroupRule) implementsZeroTrustAccessGroupListResponseExclude() {
}

type ZeroTrustAccessGroupListResponseExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                             `json:"email,required"`
	JSON  zeroTrustAccessGroupListResponseExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessOktaGroupRuleOktaJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessOktaGroupRuleOkta]
type zeroTrustAccessGroupListResponseExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessGroupListResponseExcludeAccessSamlGroupRule struct {
	Saml ZeroTrustAccessGroupListResponseExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustAccessGroupListResponseExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessSamlGroupRule]
type zeroTrustAccessGroupListResponseExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseExcludeAccessSamlGroupRule) implementsZeroTrustAccessGroupListResponseExclude() {
}

type ZeroTrustAccessGroupListResponseExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                             `json:"attribute_value,required"`
	JSON           zeroTrustAccessGroupListResponseExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessSamlGroupRuleSamlJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessSamlGroupRuleSaml]
type zeroTrustAccessGroupListResponseExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustAccessGroupListResponseExcludeAccessServiceTokenRule struct {
	ServiceToken ZeroTrustAccessGroupListResponseExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustAccessGroupListResponseExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessServiceTokenRule]
type zeroTrustAccessGroupListResponseExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseExcludeAccessServiceTokenRule) implementsZeroTrustAccessGroupListResponseExclude() {
}

type ZeroTrustAccessGroupListResponseExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                        `json:"token_id,required"`
	JSON    zeroTrustAccessGroupListResponseExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessServiceTokenRuleServiceToken]
type zeroTrustAccessGroupListResponseExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustAccessGroupListResponseExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                               `json:"any_valid_service_token,required"`
	JSON                 zeroTrustAccessGroupListResponseExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessAnyValidServiceTokenRule]
type zeroTrustAccessGroupListResponseExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseExcludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupListResponseExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessGroupListResponseExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustAccessGroupListResponseExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustAccessGroupListResponseExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessExternalEvaluationRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessExternalEvaluationRule]
type zeroTrustAccessGroupListResponseExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseExcludeAccessExternalEvaluationRule) implementsZeroTrustAccessGroupListResponseExclude() {
}

type ZeroTrustAccessGroupListResponseExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                    `json:"keys_url,required"`
	JSON    zeroTrustAccessGroupListResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustAccessGroupListResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustAccessGroupListResponseExcludeAccessCountryRule struct {
	Geo  ZeroTrustAccessGroupListResponseExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustAccessGroupListResponseExcludeAccessCountryRuleJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessCountryRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessCountryRule]
type zeroTrustAccessGroupListResponseExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseExcludeAccessCountryRule) implementsZeroTrustAccessGroupListResponseExclude() {
}

type ZeroTrustAccessGroupListResponseExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                          `json:"country_code,required"`
	JSON        zeroTrustAccessGroupListResponseExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessCountryRuleGeo]
type zeroTrustAccessGroupListResponseExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustAccessGroupListResponseExcludeAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustAccessGroupListResponseExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupListResponseExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessAuthenticationMethodRule]
type zeroTrustAccessGroupListResponseExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseExcludeAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupListResponseExclude() {
}

type ZeroTrustAccessGroupListResponseExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                              `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupListResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustAccessGroupListResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessGroupListResponseExcludeAccessDevicePostureRule struct {
	DevicePosture ZeroTrustAccessGroupListResponseExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustAccessGroupListResponseExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessDevicePostureRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessDevicePostureRule]
type zeroTrustAccessGroupListResponseExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseExcludeAccessDevicePostureRule) implementsZeroTrustAccessGroupListResponseExclude() {
}

type ZeroTrustAccessGroupListResponseExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                          `json:"integration_uid,required"`
	JSON           zeroTrustAccessGroupListResponseExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseExcludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseExcludeAccessDevicePostureRuleDevicePosture]
type zeroTrustAccessGroupListResponseExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseExcludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by [ZeroTrustAccessGroupListResponseIncludeAccessEmailRule],
// [ZeroTrustAccessGroupListResponseIncludeAccessEmailListRule],
// [ZeroTrustAccessGroupListResponseIncludeAccessDomainRule],
// [ZeroTrustAccessGroupListResponseIncludeAccessEveryoneRule],
// [ZeroTrustAccessGroupListResponseIncludeAccessIPRule],
// [ZeroTrustAccessGroupListResponseIncludeAccessIPListRule],
// [ZeroTrustAccessGroupListResponseIncludeAccessCertificateRule],
// [ZeroTrustAccessGroupListResponseIncludeAccessAccessGroupRule],
// [ZeroTrustAccessGroupListResponseIncludeAccessAzureGroupRule],
// [ZeroTrustAccessGroupListResponseIncludeAccessGitHubOrganizationRule],
// [ZeroTrustAccessGroupListResponseIncludeAccessGsuiteGroupRule],
// [ZeroTrustAccessGroupListResponseIncludeAccessOktaGroupRule],
// [ZeroTrustAccessGroupListResponseIncludeAccessSamlGroupRule],
// [ZeroTrustAccessGroupListResponseIncludeAccessServiceTokenRule],
// [ZeroTrustAccessGroupListResponseIncludeAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessGroupListResponseIncludeAccessExternalEvaluationRule],
// [ZeroTrustAccessGroupListResponseIncludeAccessCountryRule],
// [ZeroTrustAccessGroupListResponseIncludeAccessAuthenticationMethodRule] or
// [ZeroTrustAccessGroupListResponseIncludeAccessDevicePostureRule].
type ZeroTrustAccessGroupListResponseInclude interface {
	implementsZeroTrustAccessGroupListResponseInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustAccessGroupListResponseInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type ZeroTrustAccessGroupListResponseIncludeAccessEmailRule struct {
	Email ZeroTrustAccessGroupListResponseIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustAccessGroupListResponseIncludeAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessEmailRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupListResponseIncludeAccessEmailRule]
type zeroTrustAccessGroupListResponseIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIncludeAccessEmailRule) implementsZeroTrustAccessGroupListResponseInclude() {
}

type ZeroTrustAccessGroupListResponseIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                          `json:"email,required" format:"email"`
	JSON  zeroTrustAccessGroupListResponseIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessEmailRuleEmail]
type zeroTrustAccessGroupListResponseIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustAccessGroupListResponseIncludeAccessEmailListRule struct {
	EmailList ZeroTrustAccessGroupListResponseIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustAccessGroupListResponseIncludeAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessEmailListRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessEmailListRule]
type zeroTrustAccessGroupListResponseIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIncludeAccessEmailListRule) implementsZeroTrustAccessGroupListResponseInclude() {
}

type ZeroTrustAccessGroupListResponseIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                  `json:"id,required"`
	JSON zeroTrustAccessGroupListResponseIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessEmailListRuleEmailListJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessEmailListRuleEmailList]
type zeroTrustAccessGroupListResponseIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustAccessGroupListResponseIncludeAccessDomainRule struct {
	EmailDomain ZeroTrustAccessGroupListResponseIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustAccessGroupListResponseIncludeAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessDomainRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessDomainRule]
type zeroTrustAccessGroupListResponseIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIncludeAccessDomainRule) implementsZeroTrustAccessGroupListResponseInclude() {
}

type ZeroTrustAccessGroupListResponseIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                 `json:"domain,required"`
	JSON   zeroTrustAccessGroupListResponseIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessDomainRuleEmailDomainJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessDomainRuleEmailDomain]
type zeroTrustAccessGroupListResponseIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustAccessGroupListResponseIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                   `json:"everyone,required"`
	JSON     zeroTrustAccessGroupListResponseIncludeAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessEveryoneRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessEveryoneRule]
type zeroTrustAccessGroupListResponseIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIncludeAccessEveryoneRule) implementsZeroTrustAccessGroupListResponseInclude() {
}

// Matches an IP address block.
type ZeroTrustAccessGroupListResponseIncludeAccessIPRule struct {
	IP   ZeroTrustAccessGroupListResponseIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustAccessGroupListResponseIncludeAccessIPRuleJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessIPRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupListResponseIncludeAccessIPRule]
type zeroTrustAccessGroupListResponseIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIncludeAccessIPRule) implementsZeroTrustAccessGroupListResponseInclude() {
}

type ZeroTrustAccessGroupListResponseIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                    `json:"ip,required"`
	JSON zeroTrustAccessGroupListResponseIncludeAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessIPRuleIPJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupListResponseIncludeAccessIPRuleIP]
type zeroTrustAccessGroupListResponseIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustAccessGroupListResponseIncludeAccessIPListRule struct {
	IPList ZeroTrustAccessGroupListResponseIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustAccessGroupListResponseIncludeAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessIPListRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessIPListRule]
type zeroTrustAccessGroupListResponseIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIncludeAccessIPListRule) implementsZeroTrustAccessGroupListResponseInclude() {
}

type ZeroTrustAccessGroupListResponseIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                            `json:"id,required"`
	JSON zeroTrustAccessGroupListResponseIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessIPListRuleIPList]
type zeroTrustAccessGroupListResponseIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustAccessGroupListResponseIncludeAccessCertificateRule struct {
	Certificate interface{}                                                      `json:"certificate,required"`
	JSON        zeroTrustAccessGroupListResponseIncludeAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessCertificateRule]
type zeroTrustAccessGroupListResponseIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIncludeAccessCertificateRule) implementsZeroTrustAccessGroupListResponseInclude() {
}

// Matches an Access group.
type ZeroTrustAccessGroupListResponseIncludeAccessAccessGroupRule struct {
	Group ZeroTrustAccessGroupListResponseIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  zeroTrustAccessGroupListResponseIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessAccessGroupRule]
type zeroTrustAccessGroupListResponseIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIncludeAccessAccessGroupRule) implementsZeroTrustAccessGroupListResponseInclude() {
}

type ZeroTrustAccessGroupListResponseIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                `json:"id,required"`
	JSON zeroTrustAccessGroupListResponseIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessAccessGroupRuleGroup]
type zeroTrustAccessGroupListResponseIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessGroupListResponseIncludeAccessAzureGroupRule struct {
	AzureAd ZeroTrustAccessGroupListResponseIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustAccessGroupListResponseIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessAzureGroupRule]
type zeroTrustAccessGroupListResponseIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIncludeAccessAzureGroupRule) implementsZeroTrustAccessGroupListResponseInclude() {
}

type ZeroTrustAccessGroupListResponseIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                 `json:"connection_id,required"`
	JSON         zeroTrustAccessGroupListResponseIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessAzureGroupRuleAzureAdJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessAzureGroupRuleAzureAd]
type zeroTrustAccessGroupListResponseIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessGroupListResponseIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustAccessGroupListResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustAccessGroupListResponseIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessGitHubOrganizationRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessGitHubOrganizationRule]
type zeroTrustAccessGroupListResponseIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIncludeAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupListResponseInclude() {
}

type ZeroTrustAccessGroupListResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                    `json:"name,required"`
	JSON zeroTrustAccessGroupListResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustAccessGroupListResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessGroupListResponseIncludeAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustAccessGroupListResponseIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustAccessGroupListResponseIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessGsuiteGroupRule]
type zeroTrustAccessGroupListResponseIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIncludeAccessGsuiteGroupRule) implementsZeroTrustAccessGroupListResponseInclude() {
}

type ZeroTrustAccessGroupListResponseIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                 `json:"email,required"`
	JSON  zeroTrustAccessGroupListResponseIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessGsuiteGroupRuleGsuiteJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessGsuiteGroupRuleGsuite]
type zeroTrustAccessGroupListResponseIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessGroupListResponseIncludeAccessOktaGroupRule struct {
	Okta ZeroTrustAccessGroupListResponseIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustAccessGroupListResponseIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessOktaGroupRule]
type zeroTrustAccessGroupListResponseIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIncludeAccessOktaGroupRule) implementsZeroTrustAccessGroupListResponseInclude() {
}

type ZeroTrustAccessGroupListResponseIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                             `json:"email,required"`
	JSON  zeroTrustAccessGroupListResponseIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessOktaGroupRuleOktaJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessOktaGroupRuleOkta]
type zeroTrustAccessGroupListResponseIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessGroupListResponseIncludeAccessSamlGroupRule struct {
	Saml ZeroTrustAccessGroupListResponseIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustAccessGroupListResponseIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessSamlGroupRule]
type zeroTrustAccessGroupListResponseIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIncludeAccessSamlGroupRule) implementsZeroTrustAccessGroupListResponseInclude() {
}

type ZeroTrustAccessGroupListResponseIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                             `json:"attribute_value,required"`
	JSON           zeroTrustAccessGroupListResponseIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessSamlGroupRuleSamlJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessSamlGroupRuleSaml]
type zeroTrustAccessGroupListResponseIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustAccessGroupListResponseIncludeAccessServiceTokenRule struct {
	ServiceToken ZeroTrustAccessGroupListResponseIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustAccessGroupListResponseIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessServiceTokenRule]
type zeroTrustAccessGroupListResponseIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIncludeAccessServiceTokenRule) implementsZeroTrustAccessGroupListResponseInclude() {
}

type ZeroTrustAccessGroupListResponseIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                        `json:"token_id,required"`
	JSON    zeroTrustAccessGroupListResponseIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessServiceTokenRuleServiceToken]
type zeroTrustAccessGroupListResponseIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustAccessGroupListResponseIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                               `json:"any_valid_service_token,required"`
	JSON                 zeroTrustAccessGroupListResponseIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessAnyValidServiceTokenRule]
type zeroTrustAccessGroupListResponseIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIncludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupListResponseInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessGroupListResponseIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustAccessGroupListResponseIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustAccessGroupListResponseIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessExternalEvaluationRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessExternalEvaluationRule]
type zeroTrustAccessGroupListResponseIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIncludeAccessExternalEvaluationRule) implementsZeroTrustAccessGroupListResponseInclude() {
}

type ZeroTrustAccessGroupListResponseIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                    `json:"keys_url,required"`
	JSON    zeroTrustAccessGroupListResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustAccessGroupListResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustAccessGroupListResponseIncludeAccessCountryRule struct {
	Geo  ZeroTrustAccessGroupListResponseIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustAccessGroupListResponseIncludeAccessCountryRuleJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessCountryRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessCountryRule]
type zeroTrustAccessGroupListResponseIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIncludeAccessCountryRule) implementsZeroTrustAccessGroupListResponseInclude() {
}

type ZeroTrustAccessGroupListResponseIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                          `json:"country_code,required"`
	JSON        zeroTrustAccessGroupListResponseIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessCountryRuleGeo]
type zeroTrustAccessGroupListResponseIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustAccessGroupListResponseIncludeAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustAccessGroupListResponseIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupListResponseIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessAuthenticationMethodRule]
type zeroTrustAccessGroupListResponseIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIncludeAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupListResponseInclude() {
}

type ZeroTrustAccessGroupListResponseIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                              `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupListResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustAccessGroupListResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessGroupListResponseIncludeAccessDevicePostureRule struct {
	DevicePosture ZeroTrustAccessGroupListResponseIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustAccessGroupListResponseIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessDevicePostureRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessDevicePostureRule]
type zeroTrustAccessGroupListResponseIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIncludeAccessDevicePostureRule) implementsZeroTrustAccessGroupListResponseInclude() {
}

type ZeroTrustAccessGroupListResponseIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                          `json:"integration_uid,required"`
	JSON           zeroTrustAccessGroupListResponseIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIncludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIncludeAccessDevicePostureRuleDevicePosture]
type zeroTrustAccessGroupListResponseIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIncludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by [ZeroTrustAccessGroupListResponseIsDefaultAccessEmailRule],
// [ZeroTrustAccessGroupListResponseIsDefaultAccessEmailListRule],
// [ZeroTrustAccessGroupListResponseIsDefaultAccessDomainRule],
// [ZeroTrustAccessGroupListResponseIsDefaultAccessEveryoneRule],
// [ZeroTrustAccessGroupListResponseIsDefaultAccessIPRule],
// [ZeroTrustAccessGroupListResponseIsDefaultAccessIPListRule],
// [ZeroTrustAccessGroupListResponseIsDefaultAccessCertificateRule],
// [ZeroTrustAccessGroupListResponseIsDefaultAccessAccessGroupRule],
// [ZeroTrustAccessGroupListResponseIsDefaultAccessAzureGroupRule],
// [ZeroTrustAccessGroupListResponseIsDefaultAccessGitHubOrganizationRule],
// [ZeroTrustAccessGroupListResponseIsDefaultAccessGsuiteGroupRule],
// [ZeroTrustAccessGroupListResponseIsDefaultAccessOktaGroupRule],
// [ZeroTrustAccessGroupListResponseIsDefaultAccessSamlGroupRule],
// [ZeroTrustAccessGroupListResponseIsDefaultAccessServiceTokenRule],
// [ZeroTrustAccessGroupListResponseIsDefaultAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessGroupListResponseIsDefaultAccessExternalEvaluationRule],
// [ZeroTrustAccessGroupListResponseIsDefaultAccessCountryRule],
// [ZeroTrustAccessGroupListResponseIsDefaultAccessAuthenticationMethodRule] or
// [ZeroTrustAccessGroupListResponseIsDefaultAccessDevicePostureRule].
type ZeroTrustAccessGroupListResponseIsDefault interface {
	implementsZeroTrustAccessGroupListResponseIsDefault()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustAccessGroupListResponseIsDefault)(nil)).Elem(), "")
}

// Matches a specific email.
type ZeroTrustAccessGroupListResponseIsDefaultAccessEmailRule struct {
	Email ZeroTrustAccessGroupListResponseIsDefaultAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustAccessGroupListResponseIsDefaultAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessEmailRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessEmailRule]
type zeroTrustAccessGroupListResponseIsDefaultAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIsDefaultAccessEmailRule) implementsZeroTrustAccessGroupListResponseIsDefault() {
}

type ZeroTrustAccessGroupListResponseIsDefaultAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                            `json:"email,required" format:"email"`
	JSON  zeroTrustAccessGroupListResponseIsDefaultAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessEmailRuleEmail]
type zeroTrustAccessGroupListResponseIsDefaultAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustAccessGroupListResponseIsDefaultAccessEmailListRule struct {
	EmailList ZeroTrustAccessGroupListResponseIsDefaultAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustAccessGroupListResponseIsDefaultAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessEmailListRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessEmailListRule]
type zeroTrustAccessGroupListResponseIsDefaultAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIsDefaultAccessEmailListRule) implementsZeroTrustAccessGroupListResponseIsDefault() {
}

type ZeroTrustAccessGroupListResponseIsDefaultAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                    `json:"id,required"`
	JSON zeroTrustAccessGroupListResponseIsDefaultAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessEmailListRuleEmailList]
type zeroTrustAccessGroupListResponseIsDefaultAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustAccessGroupListResponseIsDefaultAccessDomainRule struct {
	EmailDomain ZeroTrustAccessGroupListResponseIsDefaultAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustAccessGroupListResponseIsDefaultAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessDomainRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessDomainRule]
type zeroTrustAccessGroupListResponseIsDefaultAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIsDefaultAccessDomainRule) implementsZeroTrustAccessGroupListResponseIsDefault() {
}

type ZeroTrustAccessGroupListResponseIsDefaultAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                   `json:"domain,required"`
	JSON   zeroTrustAccessGroupListResponseIsDefaultAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessDomainRuleEmailDomainJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessDomainRuleEmailDomain]
type zeroTrustAccessGroupListResponseIsDefaultAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustAccessGroupListResponseIsDefaultAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                     `json:"everyone,required"`
	JSON     zeroTrustAccessGroupListResponseIsDefaultAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessEveryoneRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessEveryoneRule]
type zeroTrustAccessGroupListResponseIsDefaultAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIsDefaultAccessEveryoneRule) implementsZeroTrustAccessGroupListResponseIsDefault() {
}

// Matches an IP address block.
type ZeroTrustAccessGroupListResponseIsDefaultAccessIPRule struct {
	IP   ZeroTrustAccessGroupListResponseIsDefaultAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustAccessGroupListResponseIsDefaultAccessIPRuleJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessIPRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupListResponseIsDefaultAccessIPRule]
type zeroTrustAccessGroupListResponseIsDefaultAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIsDefaultAccessIPRule) implementsZeroTrustAccessGroupListResponseIsDefault() {
}

type ZeroTrustAccessGroupListResponseIsDefaultAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                      `json:"ip,required"`
	JSON zeroTrustAccessGroupListResponseIsDefaultAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessIPRuleIPJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessIPRuleIP]
type zeroTrustAccessGroupListResponseIsDefaultAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustAccessGroupListResponseIsDefaultAccessIPListRule struct {
	IPList ZeroTrustAccessGroupListResponseIsDefaultAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustAccessGroupListResponseIsDefaultAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessIPListRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessIPListRule]
type zeroTrustAccessGroupListResponseIsDefaultAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIsDefaultAccessIPListRule) implementsZeroTrustAccessGroupListResponseIsDefault() {
}

type ZeroTrustAccessGroupListResponseIsDefaultAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                              `json:"id,required"`
	JSON zeroTrustAccessGroupListResponseIsDefaultAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessIPListRuleIPList]
type zeroTrustAccessGroupListResponseIsDefaultAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustAccessGroupListResponseIsDefaultAccessCertificateRule struct {
	Certificate interface{}                                                        `json:"certificate,required"`
	JSON        zeroTrustAccessGroupListResponseIsDefaultAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessCertificateRule]
type zeroTrustAccessGroupListResponseIsDefaultAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIsDefaultAccessCertificateRule) implementsZeroTrustAccessGroupListResponseIsDefault() {
}

// Matches an Access group.
type ZeroTrustAccessGroupListResponseIsDefaultAccessAccessGroupRule struct {
	Group ZeroTrustAccessGroupListResponseIsDefaultAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  zeroTrustAccessGroupListResponseIsDefaultAccessAccessGroupRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessAccessGroupRule]
type zeroTrustAccessGroupListResponseIsDefaultAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIsDefaultAccessAccessGroupRule) implementsZeroTrustAccessGroupListResponseIsDefault() {
}

type ZeroTrustAccessGroupListResponseIsDefaultAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                  `json:"id,required"`
	JSON zeroTrustAccessGroupListResponseIsDefaultAccessAccessGroupRuleGroupJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessAccessGroupRuleGroup]
type zeroTrustAccessGroupListResponseIsDefaultAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessGroupListResponseIsDefaultAccessAzureGroupRule struct {
	AzureAd ZeroTrustAccessGroupListResponseIsDefaultAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustAccessGroupListResponseIsDefaultAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessAzureGroupRule]
type zeroTrustAccessGroupListResponseIsDefaultAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIsDefaultAccessAzureGroupRule) implementsZeroTrustAccessGroupListResponseIsDefault() {
}

type ZeroTrustAccessGroupListResponseIsDefaultAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                   `json:"connection_id,required"`
	JSON         zeroTrustAccessGroupListResponseIsDefaultAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessAzureGroupRuleAzureAdJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessAzureGroupRuleAzureAd]
type zeroTrustAccessGroupListResponseIsDefaultAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessGroupListResponseIsDefaultAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustAccessGroupListResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustAccessGroupListResponseIsDefaultAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessGitHubOrganizationRule]
type zeroTrustAccessGroupListResponseIsDefaultAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIsDefaultAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupListResponseIsDefault() {
}

type ZeroTrustAccessGroupListResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                      `json:"name,required"`
	JSON zeroTrustAccessGroupListResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustAccessGroupListResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessGroupListResponseIsDefaultAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustAccessGroupListResponseIsDefaultAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustAccessGroupListResponseIsDefaultAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessGsuiteGroupRule]
type zeroTrustAccessGroupListResponseIsDefaultAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIsDefaultAccessGsuiteGroupRule) implementsZeroTrustAccessGroupListResponseIsDefault() {
}

type ZeroTrustAccessGroupListResponseIsDefaultAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                   `json:"email,required"`
	JSON  zeroTrustAccessGroupListResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessGsuiteGroupRuleGsuite]
type zeroTrustAccessGroupListResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessGroupListResponseIsDefaultAccessOktaGroupRule struct {
	Okta ZeroTrustAccessGroupListResponseIsDefaultAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustAccessGroupListResponseIsDefaultAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessOktaGroupRule]
type zeroTrustAccessGroupListResponseIsDefaultAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIsDefaultAccessOktaGroupRule) implementsZeroTrustAccessGroupListResponseIsDefault() {
}

type ZeroTrustAccessGroupListResponseIsDefaultAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                               `json:"email,required"`
	JSON  zeroTrustAccessGroupListResponseIsDefaultAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessOktaGroupRuleOktaJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessOktaGroupRuleOkta]
type zeroTrustAccessGroupListResponseIsDefaultAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessGroupListResponseIsDefaultAccessSamlGroupRule struct {
	Saml ZeroTrustAccessGroupListResponseIsDefaultAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustAccessGroupListResponseIsDefaultAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessSamlGroupRule]
type zeroTrustAccessGroupListResponseIsDefaultAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIsDefaultAccessSamlGroupRule) implementsZeroTrustAccessGroupListResponseIsDefault() {
}

type ZeroTrustAccessGroupListResponseIsDefaultAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                               `json:"attribute_value,required"`
	JSON           zeroTrustAccessGroupListResponseIsDefaultAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessSamlGroupRuleSamlJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessSamlGroupRuleSaml]
type zeroTrustAccessGroupListResponseIsDefaultAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustAccessGroupListResponseIsDefaultAccessServiceTokenRule struct {
	ServiceToken ZeroTrustAccessGroupListResponseIsDefaultAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustAccessGroupListResponseIsDefaultAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessServiceTokenRule]
type zeroTrustAccessGroupListResponseIsDefaultAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIsDefaultAccessServiceTokenRule) implementsZeroTrustAccessGroupListResponseIsDefault() {
}

type ZeroTrustAccessGroupListResponseIsDefaultAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                          `json:"token_id,required"`
	JSON    zeroTrustAccessGroupListResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessServiceTokenRuleServiceToken]
type zeroTrustAccessGroupListResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustAccessGroupListResponseIsDefaultAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                 `json:"any_valid_service_token,required"`
	JSON                 zeroTrustAccessGroupListResponseIsDefaultAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessAnyValidServiceTokenRule]
type zeroTrustAccessGroupListResponseIsDefaultAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIsDefaultAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupListResponseIsDefault() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessGroupListResponseIsDefaultAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustAccessGroupListResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustAccessGroupListResponseIsDefaultAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessExternalEvaluationRule]
type zeroTrustAccessGroupListResponseIsDefaultAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIsDefaultAccessExternalEvaluationRule) implementsZeroTrustAccessGroupListResponseIsDefault() {
}

type ZeroTrustAccessGroupListResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                      `json:"keys_url,required"`
	JSON    zeroTrustAccessGroupListResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustAccessGroupListResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustAccessGroupListResponseIsDefaultAccessCountryRule struct {
	Geo  ZeroTrustAccessGroupListResponseIsDefaultAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustAccessGroupListResponseIsDefaultAccessCountryRuleJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessCountryRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessCountryRule]
type zeroTrustAccessGroupListResponseIsDefaultAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIsDefaultAccessCountryRule) implementsZeroTrustAccessGroupListResponseIsDefault() {
}

type ZeroTrustAccessGroupListResponseIsDefaultAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                            `json:"country_code,required"`
	JSON        zeroTrustAccessGroupListResponseIsDefaultAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessCountryRuleGeo]
type zeroTrustAccessGroupListResponseIsDefaultAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustAccessGroupListResponseIsDefaultAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustAccessGroupListResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupListResponseIsDefaultAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessAuthenticationMethodRule]
type zeroTrustAccessGroupListResponseIsDefaultAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIsDefaultAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupListResponseIsDefault() {
}

type ZeroTrustAccessGroupListResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                                `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupListResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustAccessGroupListResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessGroupListResponseIsDefaultAccessDevicePostureRule struct {
	DevicePosture ZeroTrustAccessGroupListResponseIsDefaultAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustAccessGroupListResponseIsDefaultAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessDevicePostureRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessDevicePostureRule]
type zeroTrustAccessGroupListResponseIsDefaultAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseIsDefaultAccessDevicePostureRule) implementsZeroTrustAccessGroupListResponseIsDefault() {
}

type ZeroTrustAccessGroupListResponseIsDefaultAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                            `json:"integration_uid,required"`
	JSON           zeroTrustAccessGroupListResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseIsDefaultAccessDevicePostureRuleDevicePosture]
type zeroTrustAccessGroupListResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseIsDefaultAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by [ZeroTrustAccessGroupListResponseRequireAccessEmailRule],
// [ZeroTrustAccessGroupListResponseRequireAccessEmailListRule],
// [ZeroTrustAccessGroupListResponseRequireAccessDomainRule],
// [ZeroTrustAccessGroupListResponseRequireAccessEveryoneRule],
// [ZeroTrustAccessGroupListResponseRequireAccessIPRule],
// [ZeroTrustAccessGroupListResponseRequireAccessIPListRule],
// [ZeroTrustAccessGroupListResponseRequireAccessCertificateRule],
// [ZeroTrustAccessGroupListResponseRequireAccessAccessGroupRule],
// [ZeroTrustAccessGroupListResponseRequireAccessAzureGroupRule],
// [ZeroTrustAccessGroupListResponseRequireAccessGitHubOrganizationRule],
// [ZeroTrustAccessGroupListResponseRequireAccessGsuiteGroupRule],
// [ZeroTrustAccessGroupListResponseRequireAccessOktaGroupRule],
// [ZeroTrustAccessGroupListResponseRequireAccessSamlGroupRule],
// [ZeroTrustAccessGroupListResponseRequireAccessServiceTokenRule],
// [ZeroTrustAccessGroupListResponseRequireAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessGroupListResponseRequireAccessExternalEvaluationRule],
// [ZeroTrustAccessGroupListResponseRequireAccessCountryRule],
// [ZeroTrustAccessGroupListResponseRequireAccessAuthenticationMethodRule] or
// [ZeroTrustAccessGroupListResponseRequireAccessDevicePostureRule].
type ZeroTrustAccessGroupListResponseRequire interface {
	implementsZeroTrustAccessGroupListResponseRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustAccessGroupListResponseRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type ZeroTrustAccessGroupListResponseRequireAccessEmailRule struct {
	Email ZeroTrustAccessGroupListResponseRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustAccessGroupListResponseRequireAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessEmailRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupListResponseRequireAccessEmailRule]
type zeroTrustAccessGroupListResponseRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseRequireAccessEmailRule) implementsZeroTrustAccessGroupListResponseRequire() {
}

type ZeroTrustAccessGroupListResponseRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                          `json:"email,required" format:"email"`
	JSON  zeroTrustAccessGroupListResponseRequireAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessEmailRuleEmail]
type zeroTrustAccessGroupListResponseRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustAccessGroupListResponseRequireAccessEmailListRule struct {
	EmailList ZeroTrustAccessGroupListResponseRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustAccessGroupListResponseRequireAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessEmailListRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessEmailListRule]
type zeroTrustAccessGroupListResponseRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseRequireAccessEmailListRule) implementsZeroTrustAccessGroupListResponseRequire() {
}

type ZeroTrustAccessGroupListResponseRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                  `json:"id,required"`
	JSON zeroTrustAccessGroupListResponseRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessEmailListRuleEmailListJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessEmailListRuleEmailList]
type zeroTrustAccessGroupListResponseRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustAccessGroupListResponseRequireAccessDomainRule struct {
	EmailDomain ZeroTrustAccessGroupListResponseRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustAccessGroupListResponseRequireAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessDomainRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessDomainRule]
type zeroTrustAccessGroupListResponseRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseRequireAccessDomainRule) implementsZeroTrustAccessGroupListResponseRequire() {
}

type ZeroTrustAccessGroupListResponseRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                 `json:"domain,required"`
	JSON   zeroTrustAccessGroupListResponseRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessDomainRuleEmailDomainJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessDomainRuleEmailDomain]
type zeroTrustAccessGroupListResponseRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustAccessGroupListResponseRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                   `json:"everyone,required"`
	JSON     zeroTrustAccessGroupListResponseRequireAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessEveryoneRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessEveryoneRule]
type zeroTrustAccessGroupListResponseRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseRequireAccessEveryoneRule) implementsZeroTrustAccessGroupListResponseRequire() {
}

// Matches an IP address block.
type ZeroTrustAccessGroupListResponseRequireAccessIPRule struct {
	IP   ZeroTrustAccessGroupListResponseRequireAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustAccessGroupListResponseRequireAccessIPRuleJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessIPRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupListResponseRequireAccessIPRule]
type zeroTrustAccessGroupListResponseRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseRequireAccessIPRule) implementsZeroTrustAccessGroupListResponseRequire() {
}

type ZeroTrustAccessGroupListResponseRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                    `json:"ip,required"`
	JSON zeroTrustAccessGroupListResponseRequireAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessIPRuleIPJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupListResponseRequireAccessIPRuleIP]
type zeroTrustAccessGroupListResponseRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustAccessGroupListResponseRequireAccessIPListRule struct {
	IPList ZeroTrustAccessGroupListResponseRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustAccessGroupListResponseRequireAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessIPListRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessIPListRule]
type zeroTrustAccessGroupListResponseRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseRequireAccessIPListRule) implementsZeroTrustAccessGroupListResponseRequire() {
}

type ZeroTrustAccessGroupListResponseRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                            `json:"id,required"`
	JSON zeroTrustAccessGroupListResponseRequireAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessIPListRuleIPList]
type zeroTrustAccessGroupListResponseRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustAccessGroupListResponseRequireAccessCertificateRule struct {
	Certificate interface{}                                                      `json:"certificate,required"`
	JSON        zeroTrustAccessGroupListResponseRequireAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessCertificateRule]
type zeroTrustAccessGroupListResponseRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseRequireAccessCertificateRule) implementsZeroTrustAccessGroupListResponseRequire() {
}

// Matches an Access group.
type ZeroTrustAccessGroupListResponseRequireAccessAccessGroupRule struct {
	Group ZeroTrustAccessGroupListResponseRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  zeroTrustAccessGroupListResponseRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessAccessGroupRule]
type zeroTrustAccessGroupListResponseRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseRequireAccessAccessGroupRule) implementsZeroTrustAccessGroupListResponseRequire() {
}

type ZeroTrustAccessGroupListResponseRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                `json:"id,required"`
	JSON zeroTrustAccessGroupListResponseRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessAccessGroupRuleGroup]
type zeroTrustAccessGroupListResponseRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessGroupListResponseRequireAccessAzureGroupRule struct {
	AzureAd ZeroTrustAccessGroupListResponseRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustAccessGroupListResponseRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessAzureGroupRule]
type zeroTrustAccessGroupListResponseRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseRequireAccessAzureGroupRule) implementsZeroTrustAccessGroupListResponseRequire() {
}

type ZeroTrustAccessGroupListResponseRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                 `json:"connection_id,required"`
	JSON         zeroTrustAccessGroupListResponseRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessAzureGroupRuleAzureAdJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessAzureGroupRuleAzureAd]
type zeroTrustAccessGroupListResponseRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessGroupListResponseRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustAccessGroupListResponseRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustAccessGroupListResponseRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessGitHubOrganizationRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessGitHubOrganizationRule]
type zeroTrustAccessGroupListResponseRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseRequireAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupListResponseRequire() {
}

type ZeroTrustAccessGroupListResponseRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                    `json:"name,required"`
	JSON zeroTrustAccessGroupListResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustAccessGroupListResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessGroupListResponseRequireAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustAccessGroupListResponseRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustAccessGroupListResponseRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessGsuiteGroupRule]
type zeroTrustAccessGroupListResponseRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseRequireAccessGsuiteGroupRule) implementsZeroTrustAccessGroupListResponseRequire() {
}

type ZeroTrustAccessGroupListResponseRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                 `json:"email,required"`
	JSON  zeroTrustAccessGroupListResponseRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessGsuiteGroupRuleGsuiteJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessGsuiteGroupRuleGsuite]
type zeroTrustAccessGroupListResponseRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessGroupListResponseRequireAccessOktaGroupRule struct {
	Okta ZeroTrustAccessGroupListResponseRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustAccessGroupListResponseRequireAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessOktaGroupRule]
type zeroTrustAccessGroupListResponseRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseRequireAccessOktaGroupRule) implementsZeroTrustAccessGroupListResponseRequire() {
}

type ZeroTrustAccessGroupListResponseRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                             `json:"email,required"`
	JSON  zeroTrustAccessGroupListResponseRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessOktaGroupRuleOktaJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessOktaGroupRuleOkta]
type zeroTrustAccessGroupListResponseRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessGroupListResponseRequireAccessSamlGroupRule struct {
	Saml ZeroTrustAccessGroupListResponseRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustAccessGroupListResponseRequireAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessSamlGroupRule]
type zeroTrustAccessGroupListResponseRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseRequireAccessSamlGroupRule) implementsZeroTrustAccessGroupListResponseRequire() {
}

type ZeroTrustAccessGroupListResponseRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                             `json:"attribute_value,required"`
	JSON           zeroTrustAccessGroupListResponseRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessSamlGroupRuleSamlJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessSamlGroupRuleSaml]
type zeroTrustAccessGroupListResponseRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustAccessGroupListResponseRequireAccessServiceTokenRule struct {
	ServiceToken ZeroTrustAccessGroupListResponseRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustAccessGroupListResponseRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessServiceTokenRule]
type zeroTrustAccessGroupListResponseRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseRequireAccessServiceTokenRule) implementsZeroTrustAccessGroupListResponseRequire() {
}

type ZeroTrustAccessGroupListResponseRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                        `json:"token_id,required"`
	JSON    zeroTrustAccessGroupListResponseRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessServiceTokenRuleServiceToken]
type zeroTrustAccessGroupListResponseRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustAccessGroupListResponseRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                               `json:"any_valid_service_token,required"`
	JSON                 zeroTrustAccessGroupListResponseRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessAnyValidServiceTokenRule]
type zeroTrustAccessGroupListResponseRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseRequireAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupListResponseRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessGroupListResponseRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustAccessGroupListResponseRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustAccessGroupListResponseRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessExternalEvaluationRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessExternalEvaluationRule]
type zeroTrustAccessGroupListResponseRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseRequireAccessExternalEvaluationRule) implementsZeroTrustAccessGroupListResponseRequire() {
}

type ZeroTrustAccessGroupListResponseRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                    `json:"keys_url,required"`
	JSON    zeroTrustAccessGroupListResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustAccessGroupListResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustAccessGroupListResponseRequireAccessCountryRule struct {
	Geo  ZeroTrustAccessGroupListResponseRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustAccessGroupListResponseRequireAccessCountryRuleJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessCountryRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessCountryRule]
type zeroTrustAccessGroupListResponseRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseRequireAccessCountryRule) implementsZeroTrustAccessGroupListResponseRequire() {
}

type ZeroTrustAccessGroupListResponseRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                          `json:"country_code,required"`
	JSON        zeroTrustAccessGroupListResponseRequireAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessCountryRuleGeo]
type zeroTrustAccessGroupListResponseRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustAccessGroupListResponseRequireAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustAccessGroupListResponseRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupListResponseRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessAuthenticationMethodRule]
type zeroTrustAccessGroupListResponseRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseRequireAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupListResponseRequire() {
}

type ZeroTrustAccessGroupListResponseRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                              `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupListResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustAccessGroupListResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessGroupListResponseRequireAccessDevicePostureRule struct {
	DevicePosture ZeroTrustAccessGroupListResponseRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustAccessGroupListResponseRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessDevicePostureRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessDevicePostureRule]
type zeroTrustAccessGroupListResponseRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupListResponseRequireAccessDevicePostureRule) implementsZeroTrustAccessGroupListResponseRequire() {
}

type ZeroTrustAccessGroupListResponseRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                          `json:"integration_uid,required"`
	JSON           zeroTrustAccessGroupListResponseRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustAccessGroupListResponseRequireAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupListResponseRequireAccessDevicePostureRuleDevicePosture]
type zeroTrustAccessGroupListResponseRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupListResponseRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupListResponseRequireAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessGroupDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessGroupGetResponse struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []ZeroTrustAccessGroupGetResponseExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []ZeroTrustAccessGroupGetResponseInclude `json:"include"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	IsDefault []ZeroTrustAccessGroupGetResponseIsDefault `json:"is_default"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []ZeroTrustAccessGroupGetResponseRequire `json:"require"`
	UpdatedAt time.Time                                `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessGroupGetResponseJSON      `json:"-"`
}

// zeroTrustAccessGroupGetResponseJSON contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponse]
type zeroTrustAccessGroupGetResponseJSON struct {
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

func (r *ZeroTrustAccessGroupGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by [ZeroTrustAccessGroupGetResponseExcludeAccessEmailRule],
// [ZeroTrustAccessGroupGetResponseExcludeAccessEmailListRule],
// [ZeroTrustAccessGroupGetResponseExcludeAccessDomainRule],
// [ZeroTrustAccessGroupGetResponseExcludeAccessEveryoneRule],
// [ZeroTrustAccessGroupGetResponseExcludeAccessIPRule],
// [ZeroTrustAccessGroupGetResponseExcludeAccessIPListRule],
// [ZeroTrustAccessGroupGetResponseExcludeAccessCertificateRule],
// [ZeroTrustAccessGroupGetResponseExcludeAccessAccessGroupRule],
// [ZeroTrustAccessGroupGetResponseExcludeAccessAzureGroupRule],
// [ZeroTrustAccessGroupGetResponseExcludeAccessGitHubOrganizationRule],
// [ZeroTrustAccessGroupGetResponseExcludeAccessGsuiteGroupRule],
// [ZeroTrustAccessGroupGetResponseExcludeAccessOktaGroupRule],
// [ZeroTrustAccessGroupGetResponseExcludeAccessSamlGroupRule],
// [ZeroTrustAccessGroupGetResponseExcludeAccessServiceTokenRule],
// [ZeroTrustAccessGroupGetResponseExcludeAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessGroupGetResponseExcludeAccessExternalEvaluationRule],
// [ZeroTrustAccessGroupGetResponseExcludeAccessCountryRule],
// [ZeroTrustAccessGroupGetResponseExcludeAccessAuthenticationMethodRule] or
// [ZeroTrustAccessGroupGetResponseExcludeAccessDevicePostureRule].
type ZeroTrustAccessGroupGetResponseExclude interface {
	implementsZeroTrustAccessGroupGetResponseExclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustAccessGroupGetResponseExclude)(nil)).Elem(), "")
}

// Matches a specific email.
type ZeroTrustAccessGroupGetResponseExcludeAccessEmailRule struct {
	Email ZeroTrustAccessGroupGetResponseExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustAccessGroupGetResponseExcludeAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessEmailRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupGetResponseExcludeAccessEmailRule]
type zeroTrustAccessGroupGetResponseExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseExcludeAccessEmailRule) implementsZeroTrustAccessGroupGetResponseExclude() {
}

type ZeroTrustAccessGroupGetResponseExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                         `json:"email,required" format:"email"`
	JSON  zeroTrustAccessGroupGetResponseExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupGetResponseExcludeAccessEmailRuleEmail]
type zeroTrustAccessGroupGetResponseExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustAccessGroupGetResponseExcludeAccessEmailListRule struct {
	EmailList ZeroTrustAccessGroupGetResponseExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustAccessGroupGetResponseExcludeAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessEmailListRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupGetResponseExcludeAccessEmailListRule]
type zeroTrustAccessGroupGetResponseExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseExcludeAccessEmailListRule) implementsZeroTrustAccessGroupGetResponseExclude() {
}

type ZeroTrustAccessGroupGetResponseExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                 `json:"id,required"`
	JSON zeroTrustAccessGroupGetResponseExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessEmailListRuleEmailListJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseExcludeAccessEmailListRuleEmailList]
type zeroTrustAccessGroupGetResponseExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustAccessGroupGetResponseExcludeAccessDomainRule struct {
	EmailDomain ZeroTrustAccessGroupGetResponseExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustAccessGroupGetResponseExcludeAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessDomainRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupGetResponseExcludeAccessDomainRule]
type zeroTrustAccessGroupGetResponseExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseExcludeAccessDomainRule) implementsZeroTrustAccessGroupGetResponseExclude() {
}

type ZeroTrustAccessGroupGetResponseExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                `json:"domain,required"`
	JSON   zeroTrustAccessGroupGetResponseExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessDomainRuleEmailDomainJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseExcludeAccessDomainRuleEmailDomain]
type zeroTrustAccessGroupGetResponseExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustAccessGroupGetResponseExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                  `json:"everyone,required"`
	JSON     zeroTrustAccessGroupGetResponseExcludeAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessEveryoneRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupGetResponseExcludeAccessEveryoneRule]
type zeroTrustAccessGroupGetResponseExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseExcludeAccessEveryoneRule) implementsZeroTrustAccessGroupGetResponseExclude() {
}

// Matches an IP address block.
type ZeroTrustAccessGroupGetResponseExcludeAccessIPRule struct {
	IP   ZeroTrustAccessGroupGetResponseExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustAccessGroupGetResponseExcludeAccessIPRuleJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessIPRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupGetResponseExcludeAccessIPRule]
type zeroTrustAccessGroupGetResponseExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseExcludeAccessIPRule) implementsZeroTrustAccessGroupGetResponseExclude() {
}

type ZeroTrustAccessGroupGetResponseExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                   `json:"ip,required"`
	JSON zeroTrustAccessGroupGetResponseExcludeAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessIPRuleIPJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupGetResponseExcludeAccessIPRuleIP]
type zeroTrustAccessGroupGetResponseExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustAccessGroupGetResponseExcludeAccessIPListRule struct {
	IPList ZeroTrustAccessGroupGetResponseExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustAccessGroupGetResponseExcludeAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessIPListRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupGetResponseExcludeAccessIPListRule]
type zeroTrustAccessGroupGetResponseExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseExcludeAccessIPListRule) implementsZeroTrustAccessGroupGetResponseExclude() {
}

type ZeroTrustAccessGroupGetResponseExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                           `json:"id,required"`
	JSON zeroTrustAccessGroupGetResponseExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseExcludeAccessIPListRuleIPList]
type zeroTrustAccessGroupGetResponseExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustAccessGroupGetResponseExcludeAccessCertificateRule struct {
	Certificate interface{}                                                     `json:"certificate,required"`
	JSON        zeroTrustAccessGroupGetResponseExcludeAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseExcludeAccessCertificateRule]
type zeroTrustAccessGroupGetResponseExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseExcludeAccessCertificateRule) implementsZeroTrustAccessGroupGetResponseExclude() {
}

// Matches an Access group.
type ZeroTrustAccessGroupGetResponseExcludeAccessAccessGroupRule struct {
	Group ZeroTrustAccessGroupGetResponseExcludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  zeroTrustAccessGroupGetResponseExcludeAccessAccessGroupRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseExcludeAccessAccessGroupRule]
type zeroTrustAccessGroupGetResponseExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseExcludeAccessAccessGroupRule) implementsZeroTrustAccessGroupGetResponseExclude() {
}

type ZeroTrustAccessGroupGetResponseExcludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                               `json:"id,required"`
	JSON zeroTrustAccessGroupGetResponseExcludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseExcludeAccessAccessGroupRuleGroup]
type zeroTrustAccessGroupGetResponseExcludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessGroupGetResponseExcludeAccessAzureGroupRule struct {
	AzureAd ZeroTrustAccessGroupGetResponseExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustAccessGroupGetResponseExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupGetResponseExcludeAccessAzureGroupRule]
type zeroTrustAccessGroupGetResponseExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseExcludeAccessAzureGroupRule) implementsZeroTrustAccessGroupGetResponseExclude() {
}

type ZeroTrustAccessGroupGetResponseExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                `json:"connection_id,required"`
	JSON         zeroTrustAccessGroupGetResponseExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessAzureGroupRuleAzureAdJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseExcludeAccessAzureGroupRuleAzureAd]
type zeroTrustAccessGroupGetResponseExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessGroupGetResponseExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustAccessGroupGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustAccessGroupGetResponseExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessGitHubOrganizationRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseExcludeAccessGitHubOrganizationRule]
type zeroTrustAccessGroupGetResponseExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseExcludeAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupGetResponseExclude() {
}

type ZeroTrustAccessGroupGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                   `json:"name,required"`
	JSON zeroTrustAccessGroupGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustAccessGroupGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessGroupGetResponseExcludeAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustAccessGroupGetResponseExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustAccessGroupGetResponseExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseExcludeAccessGsuiteGroupRule]
type zeroTrustAccessGroupGetResponseExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseExcludeAccessGsuiteGroupRule) implementsZeroTrustAccessGroupGetResponseExclude() {
}

type ZeroTrustAccessGroupGetResponseExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                `json:"email,required"`
	JSON  zeroTrustAccessGroupGetResponseExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessGsuiteGroupRuleGsuiteJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseExcludeAccessGsuiteGroupRuleGsuite]
type zeroTrustAccessGroupGetResponseExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessGroupGetResponseExcludeAccessOktaGroupRule struct {
	Okta ZeroTrustAccessGroupGetResponseExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustAccessGroupGetResponseExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupGetResponseExcludeAccessOktaGroupRule]
type zeroTrustAccessGroupGetResponseExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseExcludeAccessOktaGroupRule) implementsZeroTrustAccessGroupGetResponseExclude() {
}

type ZeroTrustAccessGroupGetResponseExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                            `json:"email,required"`
	JSON  zeroTrustAccessGroupGetResponseExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessOktaGroupRuleOktaJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseExcludeAccessOktaGroupRuleOkta]
type zeroTrustAccessGroupGetResponseExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessGroupGetResponseExcludeAccessSamlGroupRule struct {
	Saml ZeroTrustAccessGroupGetResponseExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustAccessGroupGetResponseExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupGetResponseExcludeAccessSamlGroupRule]
type zeroTrustAccessGroupGetResponseExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseExcludeAccessSamlGroupRule) implementsZeroTrustAccessGroupGetResponseExclude() {
}

type ZeroTrustAccessGroupGetResponseExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                            `json:"attribute_value,required"`
	JSON           zeroTrustAccessGroupGetResponseExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessSamlGroupRuleSamlJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseExcludeAccessSamlGroupRuleSaml]
type zeroTrustAccessGroupGetResponseExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustAccessGroupGetResponseExcludeAccessServiceTokenRule struct {
	ServiceToken ZeroTrustAccessGroupGetResponseExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustAccessGroupGetResponseExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseExcludeAccessServiceTokenRule]
type zeroTrustAccessGroupGetResponseExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseExcludeAccessServiceTokenRule) implementsZeroTrustAccessGroupGetResponseExclude() {
}

type ZeroTrustAccessGroupGetResponseExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                       `json:"token_id,required"`
	JSON    zeroTrustAccessGroupGetResponseExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseExcludeAccessServiceTokenRuleServiceToken]
type zeroTrustAccessGroupGetResponseExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustAccessGroupGetResponseExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                              `json:"any_valid_service_token,required"`
	JSON                 zeroTrustAccessGroupGetResponseExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseExcludeAccessAnyValidServiceTokenRule]
type zeroTrustAccessGroupGetResponseExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseExcludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupGetResponseExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessGroupGetResponseExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustAccessGroupGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustAccessGroupGetResponseExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessExternalEvaluationRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseExcludeAccessExternalEvaluationRule]
type zeroTrustAccessGroupGetResponseExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseExcludeAccessExternalEvaluationRule) implementsZeroTrustAccessGroupGetResponseExclude() {
}

type ZeroTrustAccessGroupGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                   `json:"keys_url,required"`
	JSON    zeroTrustAccessGroupGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustAccessGroupGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustAccessGroupGetResponseExcludeAccessCountryRule struct {
	Geo  ZeroTrustAccessGroupGetResponseExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustAccessGroupGetResponseExcludeAccessCountryRuleJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessCountryRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupGetResponseExcludeAccessCountryRule]
type zeroTrustAccessGroupGetResponseExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseExcludeAccessCountryRule) implementsZeroTrustAccessGroupGetResponseExclude() {
}

type ZeroTrustAccessGroupGetResponseExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                         `json:"country_code,required"`
	JSON        zeroTrustAccessGroupGetResponseExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupGetResponseExcludeAccessCountryRuleGeo]
type zeroTrustAccessGroupGetResponseExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustAccessGroupGetResponseExcludeAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustAccessGroupGetResponseExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupGetResponseExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseExcludeAccessAuthenticationMethodRule]
type zeroTrustAccessGroupGetResponseExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseExcludeAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupGetResponseExclude() {
}

type ZeroTrustAccessGroupGetResponseExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                             `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupGetResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseExcludeAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustAccessGroupGetResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessGroupGetResponseExcludeAccessDevicePostureRule struct {
	DevicePosture ZeroTrustAccessGroupGetResponseExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustAccessGroupGetResponseExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessDevicePostureRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseExcludeAccessDevicePostureRule]
type zeroTrustAccessGroupGetResponseExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseExcludeAccessDevicePostureRule) implementsZeroTrustAccessGroupGetResponseExclude() {
}

type ZeroTrustAccessGroupGetResponseExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                         `json:"integration_uid,required"`
	JSON           zeroTrustAccessGroupGetResponseExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseExcludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseExcludeAccessDevicePostureRuleDevicePosture]
type zeroTrustAccessGroupGetResponseExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseExcludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by [ZeroTrustAccessGroupGetResponseIncludeAccessEmailRule],
// [ZeroTrustAccessGroupGetResponseIncludeAccessEmailListRule],
// [ZeroTrustAccessGroupGetResponseIncludeAccessDomainRule],
// [ZeroTrustAccessGroupGetResponseIncludeAccessEveryoneRule],
// [ZeroTrustAccessGroupGetResponseIncludeAccessIPRule],
// [ZeroTrustAccessGroupGetResponseIncludeAccessIPListRule],
// [ZeroTrustAccessGroupGetResponseIncludeAccessCertificateRule],
// [ZeroTrustAccessGroupGetResponseIncludeAccessAccessGroupRule],
// [ZeroTrustAccessGroupGetResponseIncludeAccessAzureGroupRule],
// [ZeroTrustAccessGroupGetResponseIncludeAccessGitHubOrganizationRule],
// [ZeroTrustAccessGroupGetResponseIncludeAccessGsuiteGroupRule],
// [ZeroTrustAccessGroupGetResponseIncludeAccessOktaGroupRule],
// [ZeroTrustAccessGroupGetResponseIncludeAccessSamlGroupRule],
// [ZeroTrustAccessGroupGetResponseIncludeAccessServiceTokenRule],
// [ZeroTrustAccessGroupGetResponseIncludeAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessGroupGetResponseIncludeAccessExternalEvaluationRule],
// [ZeroTrustAccessGroupGetResponseIncludeAccessCountryRule],
// [ZeroTrustAccessGroupGetResponseIncludeAccessAuthenticationMethodRule] or
// [ZeroTrustAccessGroupGetResponseIncludeAccessDevicePostureRule].
type ZeroTrustAccessGroupGetResponseInclude interface {
	implementsZeroTrustAccessGroupGetResponseInclude()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustAccessGroupGetResponseInclude)(nil)).Elem(), "")
}

// Matches a specific email.
type ZeroTrustAccessGroupGetResponseIncludeAccessEmailRule struct {
	Email ZeroTrustAccessGroupGetResponseIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustAccessGroupGetResponseIncludeAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessEmailRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupGetResponseIncludeAccessEmailRule]
type zeroTrustAccessGroupGetResponseIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIncludeAccessEmailRule) implementsZeroTrustAccessGroupGetResponseInclude() {
}

type ZeroTrustAccessGroupGetResponseIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                         `json:"email,required" format:"email"`
	JSON  zeroTrustAccessGroupGetResponseIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupGetResponseIncludeAccessEmailRuleEmail]
type zeroTrustAccessGroupGetResponseIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustAccessGroupGetResponseIncludeAccessEmailListRule struct {
	EmailList ZeroTrustAccessGroupGetResponseIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustAccessGroupGetResponseIncludeAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessEmailListRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupGetResponseIncludeAccessEmailListRule]
type zeroTrustAccessGroupGetResponseIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIncludeAccessEmailListRule) implementsZeroTrustAccessGroupGetResponseInclude() {
}

type ZeroTrustAccessGroupGetResponseIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                 `json:"id,required"`
	JSON zeroTrustAccessGroupGetResponseIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessEmailListRuleEmailListJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIncludeAccessEmailListRuleEmailList]
type zeroTrustAccessGroupGetResponseIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustAccessGroupGetResponseIncludeAccessDomainRule struct {
	EmailDomain ZeroTrustAccessGroupGetResponseIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustAccessGroupGetResponseIncludeAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessDomainRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupGetResponseIncludeAccessDomainRule]
type zeroTrustAccessGroupGetResponseIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIncludeAccessDomainRule) implementsZeroTrustAccessGroupGetResponseInclude() {
}

type ZeroTrustAccessGroupGetResponseIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                `json:"domain,required"`
	JSON   zeroTrustAccessGroupGetResponseIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessDomainRuleEmailDomainJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIncludeAccessDomainRuleEmailDomain]
type zeroTrustAccessGroupGetResponseIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustAccessGroupGetResponseIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                  `json:"everyone,required"`
	JSON     zeroTrustAccessGroupGetResponseIncludeAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessEveryoneRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupGetResponseIncludeAccessEveryoneRule]
type zeroTrustAccessGroupGetResponseIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIncludeAccessEveryoneRule) implementsZeroTrustAccessGroupGetResponseInclude() {
}

// Matches an IP address block.
type ZeroTrustAccessGroupGetResponseIncludeAccessIPRule struct {
	IP   ZeroTrustAccessGroupGetResponseIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustAccessGroupGetResponseIncludeAccessIPRuleJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessIPRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupGetResponseIncludeAccessIPRule]
type zeroTrustAccessGroupGetResponseIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIncludeAccessIPRule) implementsZeroTrustAccessGroupGetResponseInclude() {
}

type ZeroTrustAccessGroupGetResponseIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                   `json:"ip,required"`
	JSON zeroTrustAccessGroupGetResponseIncludeAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessIPRuleIPJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupGetResponseIncludeAccessIPRuleIP]
type zeroTrustAccessGroupGetResponseIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustAccessGroupGetResponseIncludeAccessIPListRule struct {
	IPList ZeroTrustAccessGroupGetResponseIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustAccessGroupGetResponseIncludeAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessIPListRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupGetResponseIncludeAccessIPListRule]
type zeroTrustAccessGroupGetResponseIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIncludeAccessIPListRule) implementsZeroTrustAccessGroupGetResponseInclude() {
}

type ZeroTrustAccessGroupGetResponseIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                           `json:"id,required"`
	JSON zeroTrustAccessGroupGetResponseIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIncludeAccessIPListRuleIPList]
type zeroTrustAccessGroupGetResponseIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustAccessGroupGetResponseIncludeAccessCertificateRule struct {
	Certificate interface{}                                                     `json:"certificate,required"`
	JSON        zeroTrustAccessGroupGetResponseIncludeAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIncludeAccessCertificateRule]
type zeroTrustAccessGroupGetResponseIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIncludeAccessCertificateRule) implementsZeroTrustAccessGroupGetResponseInclude() {
}

// Matches an Access group.
type ZeroTrustAccessGroupGetResponseIncludeAccessAccessGroupRule struct {
	Group ZeroTrustAccessGroupGetResponseIncludeAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  zeroTrustAccessGroupGetResponseIncludeAccessAccessGroupRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIncludeAccessAccessGroupRule]
type zeroTrustAccessGroupGetResponseIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIncludeAccessAccessGroupRule) implementsZeroTrustAccessGroupGetResponseInclude() {
}

type ZeroTrustAccessGroupGetResponseIncludeAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                               `json:"id,required"`
	JSON zeroTrustAccessGroupGetResponseIncludeAccessAccessGroupRuleGroupJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIncludeAccessAccessGroupRuleGroup]
type zeroTrustAccessGroupGetResponseIncludeAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessGroupGetResponseIncludeAccessAzureGroupRule struct {
	AzureAd ZeroTrustAccessGroupGetResponseIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustAccessGroupGetResponseIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupGetResponseIncludeAccessAzureGroupRule]
type zeroTrustAccessGroupGetResponseIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIncludeAccessAzureGroupRule) implementsZeroTrustAccessGroupGetResponseInclude() {
}

type ZeroTrustAccessGroupGetResponseIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                `json:"connection_id,required"`
	JSON         zeroTrustAccessGroupGetResponseIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessAzureGroupRuleAzureAdJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIncludeAccessAzureGroupRuleAzureAd]
type zeroTrustAccessGroupGetResponseIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessGroupGetResponseIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustAccessGroupGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustAccessGroupGetResponseIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessGitHubOrganizationRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIncludeAccessGitHubOrganizationRule]
type zeroTrustAccessGroupGetResponseIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIncludeAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupGetResponseInclude() {
}

type ZeroTrustAccessGroupGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                   `json:"name,required"`
	JSON zeroTrustAccessGroupGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustAccessGroupGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessGroupGetResponseIncludeAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustAccessGroupGetResponseIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustAccessGroupGetResponseIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIncludeAccessGsuiteGroupRule]
type zeroTrustAccessGroupGetResponseIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIncludeAccessGsuiteGroupRule) implementsZeroTrustAccessGroupGetResponseInclude() {
}

type ZeroTrustAccessGroupGetResponseIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                `json:"email,required"`
	JSON  zeroTrustAccessGroupGetResponseIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessGsuiteGroupRuleGsuiteJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIncludeAccessGsuiteGroupRuleGsuite]
type zeroTrustAccessGroupGetResponseIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessGroupGetResponseIncludeAccessOktaGroupRule struct {
	Okta ZeroTrustAccessGroupGetResponseIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustAccessGroupGetResponseIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupGetResponseIncludeAccessOktaGroupRule]
type zeroTrustAccessGroupGetResponseIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIncludeAccessOktaGroupRule) implementsZeroTrustAccessGroupGetResponseInclude() {
}

type ZeroTrustAccessGroupGetResponseIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                            `json:"email,required"`
	JSON  zeroTrustAccessGroupGetResponseIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessOktaGroupRuleOktaJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIncludeAccessOktaGroupRuleOkta]
type zeroTrustAccessGroupGetResponseIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessGroupGetResponseIncludeAccessSamlGroupRule struct {
	Saml ZeroTrustAccessGroupGetResponseIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustAccessGroupGetResponseIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupGetResponseIncludeAccessSamlGroupRule]
type zeroTrustAccessGroupGetResponseIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIncludeAccessSamlGroupRule) implementsZeroTrustAccessGroupGetResponseInclude() {
}

type ZeroTrustAccessGroupGetResponseIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                            `json:"attribute_value,required"`
	JSON           zeroTrustAccessGroupGetResponseIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessSamlGroupRuleSamlJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIncludeAccessSamlGroupRuleSaml]
type zeroTrustAccessGroupGetResponseIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustAccessGroupGetResponseIncludeAccessServiceTokenRule struct {
	ServiceToken ZeroTrustAccessGroupGetResponseIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustAccessGroupGetResponseIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIncludeAccessServiceTokenRule]
type zeroTrustAccessGroupGetResponseIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIncludeAccessServiceTokenRule) implementsZeroTrustAccessGroupGetResponseInclude() {
}

type ZeroTrustAccessGroupGetResponseIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                       `json:"token_id,required"`
	JSON    zeroTrustAccessGroupGetResponseIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIncludeAccessServiceTokenRuleServiceToken]
type zeroTrustAccessGroupGetResponseIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustAccessGroupGetResponseIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                              `json:"any_valid_service_token,required"`
	JSON                 zeroTrustAccessGroupGetResponseIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIncludeAccessAnyValidServiceTokenRule]
type zeroTrustAccessGroupGetResponseIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIncludeAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupGetResponseInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessGroupGetResponseIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustAccessGroupGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustAccessGroupGetResponseIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessExternalEvaluationRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIncludeAccessExternalEvaluationRule]
type zeroTrustAccessGroupGetResponseIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIncludeAccessExternalEvaluationRule) implementsZeroTrustAccessGroupGetResponseInclude() {
}

type ZeroTrustAccessGroupGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                   `json:"keys_url,required"`
	JSON    zeroTrustAccessGroupGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustAccessGroupGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustAccessGroupGetResponseIncludeAccessCountryRule struct {
	Geo  ZeroTrustAccessGroupGetResponseIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustAccessGroupGetResponseIncludeAccessCountryRuleJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessCountryRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupGetResponseIncludeAccessCountryRule]
type zeroTrustAccessGroupGetResponseIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIncludeAccessCountryRule) implementsZeroTrustAccessGroupGetResponseInclude() {
}

type ZeroTrustAccessGroupGetResponseIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                         `json:"country_code,required"`
	JSON        zeroTrustAccessGroupGetResponseIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupGetResponseIncludeAccessCountryRuleGeo]
type zeroTrustAccessGroupGetResponseIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustAccessGroupGetResponseIncludeAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustAccessGroupGetResponseIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupGetResponseIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIncludeAccessAuthenticationMethodRule]
type zeroTrustAccessGroupGetResponseIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIncludeAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupGetResponseInclude() {
}

type ZeroTrustAccessGroupGetResponseIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                             `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupGetResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIncludeAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustAccessGroupGetResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessGroupGetResponseIncludeAccessDevicePostureRule struct {
	DevicePosture ZeroTrustAccessGroupGetResponseIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustAccessGroupGetResponseIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessDevicePostureRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIncludeAccessDevicePostureRule]
type zeroTrustAccessGroupGetResponseIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIncludeAccessDevicePostureRule) implementsZeroTrustAccessGroupGetResponseInclude() {
}

type ZeroTrustAccessGroupGetResponseIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                         `json:"integration_uid,required"`
	JSON           zeroTrustAccessGroupGetResponseIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIncludeAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIncludeAccessDevicePostureRuleDevicePosture]
type zeroTrustAccessGroupGetResponseIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIncludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by [ZeroTrustAccessGroupGetResponseIsDefaultAccessEmailRule],
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessEmailListRule],
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessDomainRule],
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessEveryoneRule],
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessIPRule],
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessIPListRule],
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessCertificateRule],
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessAccessGroupRule],
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessAzureGroupRule],
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessGitHubOrganizationRule],
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessGsuiteGroupRule],
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessOktaGroupRule],
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessSamlGroupRule],
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessServiceTokenRule],
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessExternalEvaluationRule],
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessCountryRule],
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessAuthenticationMethodRule] or
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessDevicePostureRule].
type ZeroTrustAccessGroupGetResponseIsDefault interface {
	implementsZeroTrustAccessGroupGetResponseIsDefault()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustAccessGroupGetResponseIsDefault)(nil)).Elem(), "")
}

// Matches a specific email.
type ZeroTrustAccessGroupGetResponseIsDefaultAccessEmailRule struct {
	Email ZeroTrustAccessGroupGetResponseIsDefaultAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustAccessGroupGetResponseIsDefaultAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessEmailRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessEmailRule]
type zeroTrustAccessGroupGetResponseIsDefaultAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIsDefaultAccessEmailRule) implementsZeroTrustAccessGroupGetResponseIsDefault() {
}

type ZeroTrustAccessGroupGetResponseIsDefaultAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                           `json:"email,required" format:"email"`
	JSON  zeroTrustAccessGroupGetResponseIsDefaultAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessEmailRuleEmailJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessEmailRuleEmail]
type zeroTrustAccessGroupGetResponseIsDefaultAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustAccessGroupGetResponseIsDefaultAccessEmailListRule struct {
	EmailList ZeroTrustAccessGroupGetResponseIsDefaultAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustAccessGroupGetResponseIsDefaultAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessEmailListRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessEmailListRule]
type zeroTrustAccessGroupGetResponseIsDefaultAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIsDefaultAccessEmailListRule) implementsZeroTrustAccessGroupGetResponseIsDefault() {
}

type ZeroTrustAccessGroupGetResponseIsDefaultAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                   `json:"id,required"`
	JSON zeroTrustAccessGroupGetResponseIsDefaultAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessEmailListRuleEmailListJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessEmailListRuleEmailList]
type zeroTrustAccessGroupGetResponseIsDefaultAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustAccessGroupGetResponseIsDefaultAccessDomainRule struct {
	EmailDomain ZeroTrustAccessGroupGetResponseIsDefaultAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustAccessGroupGetResponseIsDefaultAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessDomainRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessDomainRule]
type zeroTrustAccessGroupGetResponseIsDefaultAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIsDefaultAccessDomainRule) implementsZeroTrustAccessGroupGetResponseIsDefault() {
}

type ZeroTrustAccessGroupGetResponseIsDefaultAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                  `json:"domain,required"`
	JSON   zeroTrustAccessGroupGetResponseIsDefaultAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessDomainRuleEmailDomainJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessDomainRuleEmailDomain]
type zeroTrustAccessGroupGetResponseIsDefaultAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustAccessGroupGetResponseIsDefaultAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                    `json:"everyone,required"`
	JSON     zeroTrustAccessGroupGetResponseIsDefaultAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessEveryoneRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessEveryoneRule]
type zeroTrustAccessGroupGetResponseIsDefaultAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIsDefaultAccessEveryoneRule) implementsZeroTrustAccessGroupGetResponseIsDefault() {
}

// Matches an IP address block.
type ZeroTrustAccessGroupGetResponseIsDefaultAccessIPRule struct {
	IP   ZeroTrustAccessGroupGetResponseIsDefaultAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustAccessGroupGetResponseIsDefaultAccessIPRuleJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessIPRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupGetResponseIsDefaultAccessIPRule]
type zeroTrustAccessGroupGetResponseIsDefaultAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIsDefaultAccessIPRule) implementsZeroTrustAccessGroupGetResponseIsDefault() {
}

type ZeroTrustAccessGroupGetResponseIsDefaultAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                     `json:"ip,required"`
	JSON zeroTrustAccessGroupGetResponseIsDefaultAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessIPRuleIPJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupGetResponseIsDefaultAccessIPRuleIP]
type zeroTrustAccessGroupGetResponseIsDefaultAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustAccessGroupGetResponseIsDefaultAccessIPListRule struct {
	IPList ZeroTrustAccessGroupGetResponseIsDefaultAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustAccessGroupGetResponseIsDefaultAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessIPListRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessIPListRule]
type zeroTrustAccessGroupGetResponseIsDefaultAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIsDefaultAccessIPListRule) implementsZeroTrustAccessGroupGetResponseIsDefault() {
}

type ZeroTrustAccessGroupGetResponseIsDefaultAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                             `json:"id,required"`
	JSON zeroTrustAccessGroupGetResponseIsDefaultAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessIPListRuleIPList]
type zeroTrustAccessGroupGetResponseIsDefaultAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustAccessGroupGetResponseIsDefaultAccessCertificateRule struct {
	Certificate interface{}                                                       `json:"certificate,required"`
	JSON        zeroTrustAccessGroupGetResponseIsDefaultAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessCertificateRule]
type zeroTrustAccessGroupGetResponseIsDefaultAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIsDefaultAccessCertificateRule) implementsZeroTrustAccessGroupGetResponseIsDefault() {
}

// Matches an Access group.
type ZeroTrustAccessGroupGetResponseIsDefaultAccessAccessGroupRule struct {
	Group ZeroTrustAccessGroupGetResponseIsDefaultAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  zeroTrustAccessGroupGetResponseIsDefaultAccessAccessGroupRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessAccessGroupRule]
type zeroTrustAccessGroupGetResponseIsDefaultAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIsDefaultAccessAccessGroupRule) implementsZeroTrustAccessGroupGetResponseIsDefault() {
}

type ZeroTrustAccessGroupGetResponseIsDefaultAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                                 `json:"id,required"`
	JSON zeroTrustAccessGroupGetResponseIsDefaultAccessAccessGroupRuleGroupJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessAccessGroupRuleGroup]
type zeroTrustAccessGroupGetResponseIsDefaultAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessGroupGetResponseIsDefaultAccessAzureGroupRule struct {
	AzureAd ZeroTrustAccessGroupGetResponseIsDefaultAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustAccessGroupGetResponseIsDefaultAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessAzureGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessAzureGroupRule]
type zeroTrustAccessGroupGetResponseIsDefaultAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIsDefaultAccessAzureGroupRule) implementsZeroTrustAccessGroupGetResponseIsDefault() {
}

type ZeroTrustAccessGroupGetResponseIsDefaultAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                  `json:"connection_id,required"`
	JSON         zeroTrustAccessGroupGetResponseIsDefaultAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessAzureGroupRuleAzureAdJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessAzureGroupRuleAzureAd]
type zeroTrustAccessGroupGetResponseIsDefaultAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessGroupGetResponseIsDefaultAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustAccessGroupGetResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustAccessGroupGetResponseIsDefaultAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessGitHubOrganizationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessGitHubOrganizationRule]
type zeroTrustAccessGroupGetResponseIsDefaultAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIsDefaultAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupGetResponseIsDefault() {
}

type ZeroTrustAccessGroupGetResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                     `json:"name,required"`
	JSON zeroTrustAccessGroupGetResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustAccessGroupGetResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessGroupGetResponseIsDefaultAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustAccessGroupGetResponseIsDefaultAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustAccessGroupGetResponseIsDefaultAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessGsuiteGroupRule]
type zeroTrustAccessGroupGetResponseIsDefaultAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIsDefaultAccessGsuiteGroupRule) implementsZeroTrustAccessGroupGetResponseIsDefault() {
}

type ZeroTrustAccessGroupGetResponseIsDefaultAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                  `json:"email,required"`
	JSON  zeroTrustAccessGroupGetResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessGsuiteGroupRuleGsuite]
type zeroTrustAccessGroupGetResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessGroupGetResponseIsDefaultAccessOktaGroupRule struct {
	Okta ZeroTrustAccessGroupGetResponseIsDefaultAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustAccessGroupGetResponseIsDefaultAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessOktaGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessOktaGroupRule]
type zeroTrustAccessGroupGetResponseIsDefaultAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIsDefaultAccessOktaGroupRule) implementsZeroTrustAccessGroupGetResponseIsDefault() {
}

type ZeroTrustAccessGroupGetResponseIsDefaultAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                              `json:"email,required"`
	JSON  zeroTrustAccessGroupGetResponseIsDefaultAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessOktaGroupRuleOktaJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessOktaGroupRuleOkta]
type zeroTrustAccessGroupGetResponseIsDefaultAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessGroupGetResponseIsDefaultAccessSamlGroupRule struct {
	Saml ZeroTrustAccessGroupGetResponseIsDefaultAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustAccessGroupGetResponseIsDefaultAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessSamlGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessSamlGroupRule]
type zeroTrustAccessGroupGetResponseIsDefaultAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIsDefaultAccessSamlGroupRule) implementsZeroTrustAccessGroupGetResponseIsDefault() {
}

type ZeroTrustAccessGroupGetResponseIsDefaultAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                              `json:"attribute_value,required"`
	JSON           zeroTrustAccessGroupGetResponseIsDefaultAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessSamlGroupRuleSamlJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessSamlGroupRuleSaml]
type zeroTrustAccessGroupGetResponseIsDefaultAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustAccessGroupGetResponseIsDefaultAccessServiceTokenRule struct {
	ServiceToken ZeroTrustAccessGroupGetResponseIsDefaultAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustAccessGroupGetResponseIsDefaultAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessServiceTokenRule]
type zeroTrustAccessGroupGetResponseIsDefaultAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIsDefaultAccessServiceTokenRule) implementsZeroTrustAccessGroupGetResponseIsDefault() {
}

type ZeroTrustAccessGroupGetResponseIsDefaultAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                         `json:"token_id,required"`
	JSON    zeroTrustAccessGroupGetResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessServiceTokenRuleServiceToken]
type zeroTrustAccessGroupGetResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustAccessGroupGetResponseIsDefaultAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                                `json:"any_valid_service_token,required"`
	JSON                 zeroTrustAccessGroupGetResponseIsDefaultAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessAnyValidServiceTokenRule]
type zeroTrustAccessGroupGetResponseIsDefaultAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIsDefaultAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupGetResponseIsDefault() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessGroupGetResponseIsDefaultAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustAccessGroupGetResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustAccessGroupGetResponseIsDefaultAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessExternalEvaluationRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessExternalEvaluationRule]
type zeroTrustAccessGroupGetResponseIsDefaultAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIsDefaultAccessExternalEvaluationRule) implementsZeroTrustAccessGroupGetResponseIsDefault() {
}

type ZeroTrustAccessGroupGetResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                     `json:"keys_url,required"`
	JSON    zeroTrustAccessGroupGetResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustAccessGroupGetResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustAccessGroupGetResponseIsDefaultAccessCountryRule struct {
	Geo  ZeroTrustAccessGroupGetResponseIsDefaultAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustAccessGroupGetResponseIsDefaultAccessCountryRuleJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessCountryRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessCountryRule]
type zeroTrustAccessGroupGetResponseIsDefaultAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIsDefaultAccessCountryRule) implementsZeroTrustAccessGroupGetResponseIsDefault() {
}

type ZeroTrustAccessGroupGetResponseIsDefaultAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                           `json:"country_code,required"`
	JSON        zeroTrustAccessGroupGetResponseIsDefaultAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessCountryRuleGeoJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessCountryRuleGeo]
type zeroTrustAccessGroupGetResponseIsDefaultAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustAccessGroupGetResponseIsDefaultAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustAccessGroupGetResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupGetResponseIsDefaultAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessAuthenticationMethodRule]
type zeroTrustAccessGroupGetResponseIsDefaultAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIsDefaultAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupGetResponseIsDefault() {
}

type ZeroTrustAccessGroupGetResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                               `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupGetResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustAccessGroupGetResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessGroupGetResponseIsDefaultAccessDevicePostureRule struct {
	DevicePosture ZeroTrustAccessGroupGetResponseIsDefaultAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustAccessGroupGetResponseIsDefaultAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessDevicePostureRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessDevicePostureRule]
type zeroTrustAccessGroupGetResponseIsDefaultAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseIsDefaultAccessDevicePostureRule) implementsZeroTrustAccessGroupGetResponseIsDefault() {
}

type ZeroTrustAccessGroupGetResponseIsDefaultAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                           `json:"integration_uid,required"`
	JSON           zeroTrustAccessGroupGetResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseIsDefaultAccessDevicePostureRuleDevicePosture]
type zeroTrustAccessGroupGetResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseIsDefaultAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseIsDefaultAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by [ZeroTrustAccessGroupGetResponseRequireAccessEmailRule],
// [ZeroTrustAccessGroupGetResponseRequireAccessEmailListRule],
// [ZeroTrustAccessGroupGetResponseRequireAccessDomainRule],
// [ZeroTrustAccessGroupGetResponseRequireAccessEveryoneRule],
// [ZeroTrustAccessGroupGetResponseRequireAccessIPRule],
// [ZeroTrustAccessGroupGetResponseRequireAccessIPListRule],
// [ZeroTrustAccessGroupGetResponseRequireAccessCertificateRule],
// [ZeroTrustAccessGroupGetResponseRequireAccessAccessGroupRule],
// [ZeroTrustAccessGroupGetResponseRequireAccessAzureGroupRule],
// [ZeroTrustAccessGroupGetResponseRequireAccessGitHubOrganizationRule],
// [ZeroTrustAccessGroupGetResponseRequireAccessGsuiteGroupRule],
// [ZeroTrustAccessGroupGetResponseRequireAccessOktaGroupRule],
// [ZeroTrustAccessGroupGetResponseRequireAccessSamlGroupRule],
// [ZeroTrustAccessGroupGetResponseRequireAccessServiceTokenRule],
// [ZeroTrustAccessGroupGetResponseRequireAccessAnyValidServiceTokenRule],
// [ZeroTrustAccessGroupGetResponseRequireAccessExternalEvaluationRule],
// [ZeroTrustAccessGroupGetResponseRequireAccessCountryRule],
// [ZeroTrustAccessGroupGetResponseRequireAccessAuthenticationMethodRule] or
// [ZeroTrustAccessGroupGetResponseRequireAccessDevicePostureRule].
type ZeroTrustAccessGroupGetResponseRequire interface {
	implementsZeroTrustAccessGroupGetResponseRequire()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustAccessGroupGetResponseRequire)(nil)).Elem(), "")
}

// Matches a specific email.
type ZeroTrustAccessGroupGetResponseRequireAccessEmailRule struct {
	Email ZeroTrustAccessGroupGetResponseRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustAccessGroupGetResponseRequireAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessEmailRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupGetResponseRequireAccessEmailRule]
type zeroTrustAccessGroupGetResponseRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseRequireAccessEmailRule) implementsZeroTrustAccessGroupGetResponseRequire() {
}

type ZeroTrustAccessGroupGetResponseRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                                         `json:"email,required" format:"email"`
	JSON  zeroTrustAccessGroupGetResponseRequireAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessEmailRuleEmailJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupGetResponseRequireAccessEmailRuleEmail]
type zeroTrustAccessGroupGetResponseRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustAccessGroupGetResponseRequireAccessEmailListRule struct {
	EmailList ZeroTrustAccessGroupGetResponseRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustAccessGroupGetResponseRequireAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessEmailListRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupGetResponseRequireAccessEmailListRule]
type zeroTrustAccessGroupGetResponseRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseRequireAccessEmailListRule) implementsZeroTrustAccessGroupGetResponseRequire() {
}

type ZeroTrustAccessGroupGetResponseRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                                 `json:"id,required"`
	JSON zeroTrustAccessGroupGetResponseRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessEmailListRuleEmailListJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseRequireAccessEmailListRuleEmailList]
type zeroTrustAccessGroupGetResponseRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustAccessGroupGetResponseRequireAccessDomainRule struct {
	EmailDomain ZeroTrustAccessGroupGetResponseRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustAccessGroupGetResponseRequireAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessDomainRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupGetResponseRequireAccessDomainRule]
type zeroTrustAccessGroupGetResponseRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseRequireAccessDomainRule) implementsZeroTrustAccessGroupGetResponseRequire() {
}

type ZeroTrustAccessGroupGetResponseRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                                `json:"domain,required"`
	JSON   zeroTrustAccessGroupGetResponseRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessDomainRuleEmailDomainJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseRequireAccessDomainRuleEmailDomain]
type zeroTrustAccessGroupGetResponseRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustAccessGroupGetResponseRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                                  `json:"everyone,required"`
	JSON     zeroTrustAccessGroupGetResponseRequireAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessEveryoneRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupGetResponseRequireAccessEveryoneRule]
type zeroTrustAccessGroupGetResponseRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseRequireAccessEveryoneRule) implementsZeroTrustAccessGroupGetResponseRequire() {
}

// Matches an IP address block.
type ZeroTrustAccessGroupGetResponseRequireAccessIPRule struct {
	IP   ZeroTrustAccessGroupGetResponseRequireAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustAccessGroupGetResponseRequireAccessIPRuleJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessIPRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupGetResponseRequireAccessIPRule]
type zeroTrustAccessGroupGetResponseRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseRequireAccessIPRule) implementsZeroTrustAccessGroupGetResponseRequire() {
}

type ZeroTrustAccessGroupGetResponseRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                                   `json:"ip,required"`
	JSON zeroTrustAccessGroupGetResponseRequireAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessIPRuleIPJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupGetResponseRequireAccessIPRuleIP]
type zeroTrustAccessGroupGetResponseRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustAccessGroupGetResponseRequireAccessIPListRule struct {
	IPList ZeroTrustAccessGroupGetResponseRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustAccessGroupGetResponseRequireAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessIPListRuleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessGroupGetResponseRequireAccessIPListRule]
type zeroTrustAccessGroupGetResponseRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseRequireAccessIPListRule) implementsZeroTrustAccessGroupGetResponseRequire() {
}

type ZeroTrustAccessGroupGetResponseRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                                           `json:"id,required"`
	JSON zeroTrustAccessGroupGetResponseRequireAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessIPListRuleIPListJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseRequireAccessIPListRuleIPList]
type zeroTrustAccessGroupGetResponseRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustAccessGroupGetResponseRequireAccessCertificateRule struct {
	Certificate interface{}                                                     `json:"certificate,required"`
	JSON        zeroTrustAccessGroupGetResponseRequireAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessCertificateRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseRequireAccessCertificateRule]
type zeroTrustAccessGroupGetResponseRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseRequireAccessCertificateRule) implementsZeroTrustAccessGroupGetResponseRequire() {
}

// Matches an Access group.
type ZeroTrustAccessGroupGetResponseRequireAccessAccessGroupRule struct {
	Group ZeroTrustAccessGroupGetResponseRequireAccessAccessGroupRuleGroup `json:"group,required"`
	JSON  zeroTrustAccessGroupGetResponseRequireAccessAccessGroupRuleJSON  `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessAccessGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseRequireAccessAccessGroupRule]
type zeroTrustAccessGroupGetResponseRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseRequireAccessAccessGroupRule) implementsZeroTrustAccessGroupGetResponseRequire() {
}

type ZeroTrustAccessGroupGetResponseRequireAccessAccessGroupRuleGroup struct {
	// The ID of a previously created Access group.
	ID   string                                                               `json:"id,required"`
	JSON zeroTrustAccessGroupGetResponseRequireAccessAccessGroupRuleGroupJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessAccessGroupRuleGroupJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseRequireAccessAccessGroupRuleGroup]
type zeroTrustAccessGroupGetResponseRequireAccessAccessGroupRuleGroupJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessAccessGroupRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessAccessGroupRuleGroupJSON) RawJSON() string {
	return r.raw
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustAccessGroupGetResponseRequireAccessAzureGroupRule struct {
	AzureAd ZeroTrustAccessGroupGetResponseRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustAccessGroupGetResponseRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessAzureGroupRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupGetResponseRequireAccessAzureGroupRule]
type zeroTrustAccessGroupGetResponseRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseRequireAccessAzureGroupRule) implementsZeroTrustAccessGroupGetResponseRequire() {
}

type ZeroTrustAccessGroupGetResponseRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                                `json:"connection_id,required"`
	JSON         zeroTrustAccessGroupGetResponseRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessAzureGroupRuleAzureAdJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseRequireAccessAzureGroupRuleAzureAd]
type zeroTrustAccessGroupGetResponseRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustAccessGroupGetResponseRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustAccessGroupGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustAccessGroupGetResponseRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessGitHubOrganizationRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseRequireAccessGitHubOrganizationRule]
type zeroTrustAccessGroupGetResponseRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseRequireAccessGitHubOrganizationRule) implementsZeroTrustAccessGroupGetResponseRequire() {
}

type ZeroTrustAccessGroupGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                                   `json:"name,required"`
	JSON zeroTrustAccessGroupGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustAccessGroupGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustAccessGroupGetResponseRequireAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustAccessGroupGetResponseRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustAccessGroupGetResponseRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessGsuiteGroupRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseRequireAccessGsuiteGroupRule]
type zeroTrustAccessGroupGetResponseRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseRequireAccessGsuiteGroupRule) implementsZeroTrustAccessGroupGetResponseRequire() {
}

type ZeroTrustAccessGroupGetResponseRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                                `json:"email,required"`
	JSON  zeroTrustAccessGroupGetResponseRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessGsuiteGroupRuleGsuiteJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseRequireAccessGsuiteGroupRuleGsuite]
type zeroTrustAccessGroupGetResponseRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustAccessGroupGetResponseRequireAccessOktaGroupRule struct {
	Okta ZeroTrustAccessGroupGetResponseRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustAccessGroupGetResponseRequireAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessOktaGroupRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupGetResponseRequireAccessOktaGroupRule]
type zeroTrustAccessGroupGetResponseRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseRequireAccessOktaGroupRule) implementsZeroTrustAccessGroupGetResponseRequire() {
}

type ZeroTrustAccessGroupGetResponseRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                                            `json:"email,required"`
	JSON  zeroTrustAccessGroupGetResponseRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessOktaGroupRuleOktaJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseRequireAccessOktaGroupRuleOkta]
type zeroTrustAccessGroupGetResponseRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustAccessGroupGetResponseRequireAccessSamlGroupRule struct {
	Saml ZeroTrustAccessGroupGetResponseRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustAccessGroupGetResponseRequireAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessSamlGroupRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupGetResponseRequireAccessSamlGroupRule]
type zeroTrustAccessGroupGetResponseRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseRequireAccessSamlGroupRule) implementsZeroTrustAccessGroupGetResponseRequire() {
}

type ZeroTrustAccessGroupGetResponseRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                                            `json:"attribute_value,required"`
	JSON           zeroTrustAccessGroupGetResponseRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessSamlGroupRuleSamlJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseRequireAccessSamlGroupRuleSaml]
type zeroTrustAccessGroupGetResponseRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustAccessGroupGetResponseRequireAccessServiceTokenRule struct {
	ServiceToken ZeroTrustAccessGroupGetResponseRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustAccessGroupGetResponseRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessServiceTokenRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseRequireAccessServiceTokenRule]
type zeroTrustAccessGroupGetResponseRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseRequireAccessServiceTokenRule) implementsZeroTrustAccessGroupGetResponseRequire() {
}

type ZeroTrustAccessGroupGetResponseRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                                       `json:"token_id,required"`
	JSON    zeroTrustAccessGroupGetResponseRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessServiceTokenRuleServiceTokenJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseRequireAccessServiceTokenRuleServiceToken]
type zeroTrustAccessGroupGetResponseRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustAccessGroupGetResponseRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                              `json:"any_valid_service_token,required"`
	JSON                 zeroTrustAccessGroupGetResponseRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessAnyValidServiceTokenRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseRequireAccessAnyValidServiceTokenRule]
type zeroTrustAccessGroupGetResponseRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseRequireAccessAnyValidServiceTokenRule) implementsZeroTrustAccessGroupGetResponseRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustAccessGroupGetResponseRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustAccessGroupGetResponseRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustAccessGroupGetResponseRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessExternalEvaluationRuleJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseRequireAccessExternalEvaluationRule]
type zeroTrustAccessGroupGetResponseRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseRequireAccessExternalEvaluationRule) implementsZeroTrustAccessGroupGetResponseRequire() {
}

type ZeroTrustAccessGroupGetResponseRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                                   `json:"keys_url,required"`
	JSON    zeroTrustAccessGroupGetResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseRequireAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustAccessGroupGetResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustAccessGroupGetResponseRequireAccessCountryRule struct {
	Geo  ZeroTrustAccessGroupGetResponseRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustAccessGroupGetResponseRequireAccessCountryRuleJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessCountryRuleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupGetResponseRequireAccessCountryRule]
type zeroTrustAccessGroupGetResponseRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseRequireAccessCountryRule) implementsZeroTrustAccessGroupGetResponseRequire() {
}

type ZeroTrustAccessGroupGetResponseRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                                         `json:"country_code,required"`
	JSON        zeroTrustAccessGroupGetResponseRequireAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessCountryRuleGeoJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessGroupGetResponseRequireAccessCountryRuleGeo]
type zeroTrustAccessGroupGetResponseRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustAccessGroupGetResponseRequireAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustAccessGroupGetResponseRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupGetResponseRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessAuthenticationMethodRuleJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseRequireAccessAuthenticationMethodRule]
type zeroTrustAccessGroupGetResponseRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseRequireAccessAuthenticationMethodRule) implementsZeroTrustAccessGroupGetResponseRequire() {
}

type ZeroTrustAccessGroupGetResponseRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                                             `json:"auth_method,required"`
	JSON       zeroTrustAccessGroupGetResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseRequireAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustAccessGroupGetResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustAccessGroupGetResponseRequireAccessDevicePostureRule struct {
	DevicePosture ZeroTrustAccessGroupGetResponseRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustAccessGroupGetResponseRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessDevicePostureRuleJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseRequireAccessDevicePostureRule]
type zeroTrustAccessGroupGetResponseRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustAccessGroupGetResponseRequireAccessDevicePostureRule) implementsZeroTrustAccessGroupGetResponseRequire() {
}

type ZeroTrustAccessGroupGetResponseRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                                         `json:"integration_uid,required"`
	JSON           zeroTrustAccessGroupGetResponseRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustAccessGroupGetResponseRequireAccessDevicePostureRuleDevicePostureJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessGroupGetResponseRequireAccessDevicePostureRuleDevicePosture]
type zeroTrustAccessGroupGetResponseRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessGroupGetResponseRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessGroupGetResponseRequireAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
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
	Result   ZeroTrustAccessGroupNewResponse                   `json:"result,required"`
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

func (r zeroTrustAccessGroupNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessGroupNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessGroupNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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
	Result   ZeroTrustAccessGroupUpdateResponse                   `json:"result,required"`
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

func (r zeroTrustAccessGroupUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessGroupUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessGroupUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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
	Result   []ZeroTrustAccessGroupListResponse                 `json:"result,required,nullable"`
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

func (r zeroTrustAccessGroupListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessGroupListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessGroupListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessGroupListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessGroupDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessGroupDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessGroupDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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
	Result   ZeroTrustAccessGroupGetResponse                   `json:"result,required"`
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

func (r zeroTrustAccessGroupGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessGroupGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustAccessGroupGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustAccessGroupGetResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessGroupGetResponseEnvelopeSuccessTrue ZeroTrustAccessGroupGetResponseEnvelopeSuccess = true
)
