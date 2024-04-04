// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
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
func (r *AccessGroupService) New(ctx context.Context, params AccessGroupNewParams, opts ...option.RequestOption) (res *ZeroTrustGroups, err error) {
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
func (r *AccessGroupService) Update(ctx context.Context, uuid string, params AccessGroupUpdateParams, opts ...option.RequestOption) (res *ZeroTrustGroups, err error) {
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
func (r *AccessGroupService) List(ctx context.Context, query AccessGroupListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ZeroTrustGroups], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists all Access groups.
func (r *AccessGroupService) ListAutoPaging(ctx context.Context, query AccessGroupListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ZeroTrustGroups] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
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
func (r *AccessGroupService) Get(ctx context.Context, uuid string, query AccessGroupGetParams, opts ...option.RequestOption) (res *ZeroTrustGroups, err error) {
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

type ZeroTrustGroups struct {
	// UUID
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Rules evaluated with a NOT logical operator. To match a policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []ZeroTrustGroupsExclude `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []ZeroTrustGroupsInclude `json:"include"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	IsDefault []ZeroTrustGroupsIsDefault `json:"is_default"`
	// The name of the Access group.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match a policy, a user must
	// meet all of the Require rules.
	Require   []ZeroTrustGroupsRequire `json:"require"`
	UpdatedAt time.Time                `json:"updated_at" format:"date-time"`
	JSON      zeroTrustGroupsJSON      `json:"-"`
}

// zeroTrustGroupsJSON contains the JSON metadata for the struct [ZeroTrustGroups]
type zeroTrustGroupsJSON struct {
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

func (r *ZeroTrustGroups) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by [zero_trust.ZeroTrustGroupsExcludeAccessEmailRule],
// [zero_trust.ZeroTrustGroupsExcludeAccessEmailListRule],
// [zero_trust.ZeroTrustGroupsExcludeAccessDomainRule],
// [zero_trust.ZeroTrustGroupsExcludeAccessEveryoneRule],
// [zero_trust.ZeroTrustGroupsExcludeAccessIPRule],
// [zero_trust.ZeroTrustGroupsExcludeAccessIPListRule],
// [zero_trust.ZeroTrustGroupsExcludeAccessCertificateRule],
// [zero_trust.ZeroTrustGroupsExcludeAccessAccessGroupRule],
// [zero_trust.ZeroTrustGroupsExcludeAccessAzureGroupRule],
// [zero_trust.ZeroTrustGroupsExcludeAccessGitHubOrganizationRule],
// [zero_trust.ZeroTrustGroupsExcludeAccessGsuiteGroupRule],
// [zero_trust.ZeroTrustGroupsExcludeAccessOktaGroupRule],
// [zero_trust.ZeroTrustGroupsExcludeAccessSamlGroupRule],
// [zero_trust.ZeroTrustGroupsExcludeAccessServiceTokenRule],
// [zero_trust.ZeroTrustGroupsExcludeAccessAnyValidServiceTokenRule],
// [zero_trust.ZeroTrustGroupsExcludeAccessExternalEvaluationRule],
// [zero_trust.ZeroTrustGroupsExcludeAccessCountryRule],
// [zero_trust.ZeroTrustGroupsExcludeAccessAuthenticationMethodRule] or
// [zero_trust.ZeroTrustGroupsExcludeAccessDevicePostureRule].
type ZeroTrustGroupsExclude interface {
	implementsZeroTrustZeroTrustGroupsExclude()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustGroupsExclude)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsExcludeAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsExcludeAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsExcludeAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsExcludeAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsExcludeAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsExcludeAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsExcludeAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsExcludeAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsExcludeAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsExcludeAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsExcludeAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsExcludeAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsExcludeAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsExcludeAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsExcludeAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsExcludeAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsExcludeAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsExcludeAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsExcludeAccessDevicePostureRule{}),
		},
	)
}

// Matches a specific email.
type ZeroTrustGroupsExcludeAccessEmailRule struct {
	Email ZeroTrustGroupsExcludeAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustGroupsExcludeAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustGroupsExcludeAccessEmailRuleJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsExcludeAccessEmailRule]
type zeroTrustGroupsExcludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsExcludeAccessEmailRule) implementsZeroTrustZeroTrustGroupsExclude() {}

type ZeroTrustGroupsExcludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                         `json:"email,required" format:"email"`
	JSON  zeroTrustGroupsExcludeAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustGroupsExcludeAccessEmailRuleEmailJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsExcludeAccessEmailRuleEmail]
type zeroTrustGroupsExcludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustGroupsExcludeAccessEmailListRule struct {
	EmailList ZeroTrustGroupsExcludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustGroupsExcludeAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustGroupsExcludeAccessEmailListRuleJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsExcludeAccessEmailListRule]
type zeroTrustGroupsExcludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsExcludeAccessEmailListRule) implementsZeroTrustZeroTrustGroupsExclude() {}

type ZeroTrustGroupsExcludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                 `json:"id,required"`
	JSON zeroTrustGroupsExcludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustGroupsExcludeAccessEmailListRuleEmailListJSON contains the JSON
// metadata for the struct [ZeroTrustGroupsExcludeAccessEmailListRuleEmailList]
type zeroTrustGroupsExcludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustGroupsExcludeAccessDomainRule struct {
	EmailDomain ZeroTrustGroupsExcludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustGroupsExcludeAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustGroupsExcludeAccessDomainRuleJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsExcludeAccessDomainRule]
type zeroTrustGroupsExcludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsExcludeAccessDomainRule) implementsZeroTrustZeroTrustGroupsExclude() {}

type ZeroTrustGroupsExcludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                `json:"domain,required"`
	JSON   zeroTrustGroupsExcludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustGroupsExcludeAccessDomainRuleEmailDomainJSON contains the JSON metadata
// for the struct [ZeroTrustGroupsExcludeAccessDomainRuleEmailDomain]
type zeroTrustGroupsExcludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustGroupsExcludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                  `json:"everyone,required"`
	JSON     zeroTrustGroupsExcludeAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustGroupsExcludeAccessEveryoneRuleJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsExcludeAccessEveryoneRule]
type zeroTrustGroupsExcludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsExcludeAccessEveryoneRule) implementsZeroTrustZeroTrustGroupsExclude() {}

// Matches an IP address block.
type ZeroTrustGroupsExcludeAccessIPRule struct {
	IP   ZeroTrustGroupsExcludeAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustGroupsExcludeAccessIPRuleJSON `json:"-"`
}

// zeroTrustGroupsExcludeAccessIPRuleJSON contains the JSON metadata for the struct
// [ZeroTrustGroupsExcludeAccessIPRule]
type zeroTrustGroupsExcludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsExcludeAccessIPRule) implementsZeroTrustZeroTrustGroupsExclude() {}

type ZeroTrustGroupsExcludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                   `json:"ip,required"`
	JSON zeroTrustGroupsExcludeAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustGroupsExcludeAccessIPRuleIPJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsExcludeAccessIPRuleIP]
type zeroTrustGroupsExcludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustGroupsExcludeAccessIPListRule struct {
	IPList ZeroTrustGroupsExcludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustGroupsExcludeAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustGroupsExcludeAccessIPListRuleJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsExcludeAccessIPListRule]
type zeroTrustGroupsExcludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsExcludeAccessIPListRule) implementsZeroTrustZeroTrustGroupsExclude() {}

type ZeroTrustGroupsExcludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                           `json:"id,required"`
	JSON zeroTrustGroupsExcludeAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustGroupsExcludeAccessIPListRuleIPListJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsExcludeAccessIPListRuleIPList]
type zeroTrustGroupsExcludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustGroupsExcludeAccessCertificateRule struct {
	Certificate interface{}                                     `json:"certificate,required"`
	JSON        zeroTrustGroupsExcludeAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustGroupsExcludeAccessCertificateRuleJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsExcludeAccessCertificateRule]
type zeroTrustGroupsExcludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsExcludeAccessCertificateRule) implementsZeroTrustZeroTrustGroupsExclude() {}

// Matches an Access group.
type ZeroTrustGroupsExcludeAccessAccessGroupRule struct {
	Group shared.UnnamedSchemaRef131                      `json:"group,required"`
	JSON  zeroTrustGroupsExcludeAccessAccessGroupRuleJSON `json:"-"`
}

// zeroTrustGroupsExcludeAccessAccessGroupRuleJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsExcludeAccessAccessGroupRule]
type zeroTrustGroupsExcludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsExcludeAccessAccessGroupRule) implementsZeroTrustZeroTrustGroupsExclude() {}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustGroupsExcludeAccessAzureGroupRule struct {
	AzureAd ZeroTrustGroupsExcludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustGroupsExcludeAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustGroupsExcludeAccessAzureGroupRuleJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsExcludeAccessAzureGroupRule]
type zeroTrustGroupsExcludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsExcludeAccessAzureGroupRule) implementsZeroTrustZeroTrustGroupsExclude() {}

type ZeroTrustGroupsExcludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                `json:"connection_id,required"`
	JSON         zeroTrustGroupsExcludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustGroupsExcludeAccessAzureGroupRuleAzureAdJSON contains the JSON metadata
// for the struct [ZeroTrustGroupsExcludeAccessAzureGroupRuleAzureAd]
type zeroTrustGroupsExcludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustGroupsExcludeAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustGroupsExcludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustGroupsExcludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustGroupsExcludeAccessGitHubOrganizationRuleJSON contains the JSON
// metadata for the struct [ZeroTrustGroupsExcludeAccessGitHubOrganizationRule]
type zeroTrustGroupsExcludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsExcludeAccessGitHubOrganizationRule) implementsZeroTrustZeroTrustGroupsExclude() {
}

type ZeroTrustGroupsExcludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                   `json:"name,required"`
	JSON zeroTrustGroupsExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustGroupsExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustGroupsExcludeAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustGroupsExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustGroupsExcludeAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustGroupsExcludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustGroupsExcludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustGroupsExcludeAccessGsuiteGroupRuleJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsExcludeAccessGsuiteGroupRule]
type zeroTrustGroupsExcludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsExcludeAccessGsuiteGroupRule) implementsZeroTrustZeroTrustGroupsExclude() {}

type ZeroTrustGroupsExcludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                `json:"email,required"`
	JSON  zeroTrustGroupsExcludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustGroupsExcludeAccessGsuiteGroupRuleGsuiteJSON contains the JSON metadata
// for the struct [ZeroTrustGroupsExcludeAccessGsuiteGroupRuleGsuite]
type zeroTrustGroupsExcludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustGroupsExcludeAccessOktaGroupRule struct {
	Okta ZeroTrustGroupsExcludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustGroupsExcludeAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustGroupsExcludeAccessOktaGroupRuleJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsExcludeAccessOktaGroupRule]
type zeroTrustGroupsExcludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsExcludeAccessOktaGroupRule) implementsZeroTrustZeroTrustGroupsExclude() {}

type ZeroTrustGroupsExcludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                            `json:"email,required"`
	JSON  zeroTrustGroupsExcludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustGroupsExcludeAccessOktaGroupRuleOktaJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsExcludeAccessOktaGroupRuleOkta]
type zeroTrustGroupsExcludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustGroupsExcludeAccessSamlGroupRule struct {
	Saml ZeroTrustGroupsExcludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustGroupsExcludeAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustGroupsExcludeAccessSamlGroupRuleJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsExcludeAccessSamlGroupRule]
type zeroTrustGroupsExcludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsExcludeAccessSamlGroupRule) implementsZeroTrustZeroTrustGroupsExclude() {}

type ZeroTrustGroupsExcludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                            `json:"attribute_value,required"`
	JSON           zeroTrustGroupsExcludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustGroupsExcludeAccessSamlGroupRuleSamlJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsExcludeAccessSamlGroupRuleSaml]
type zeroTrustGroupsExcludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustGroupsExcludeAccessServiceTokenRule struct {
	ServiceToken ZeroTrustGroupsExcludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustGroupsExcludeAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustGroupsExcludeAccessServiceTokenRuleJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsExcludeAccessServiceTokenRule]
type zeroTrustGroupsExcludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsExcludeAccessServiceTokenRule) implementsZeroTrustZeroTrustGroupsExclude() {}

type ZeroTrustGroupsExcludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                       `json:"token_id,required"`
	JSON    zeroTrustGroupsExcludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustGroupsExcludeAccessServiceTokenRuleServiceTokenJSON contains the JSON
// metadata for the struct
// [ZeroTrustGroupsExcludeAccessServiceTokenRuleServiceToken]
type zeroTrustGroupsExcludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustGroupsExcludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                              `json:"any_valid_service_token,required"`
	JSON                 zeroTrustGroupsExcludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustGroupsExcludeAccessAnyValidServiceTokenRuleJSON contains the JSON
// metadata for the struct [ZeroTrustGroupsExcludeAccessAnyValidServiceTokenRule]
type zeroTrustGroupsExcludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsExcludeAccessAnyValidServiceTokenRule) implementsZeroTrustZeroTrustGroupsExclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustGroupsExcludeAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustGroupsExcludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustGroupsExcludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustGroupsExcludeAccessExternalEvaluationRuleJSON contains the JSON
// metadata for the struct [ZeroTrustGroupsExcludeAccessExternalEvaluationRule]
type zeroTrustGroupsExcludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsExcludeAccessExternalEvaluationRule) implementsZeroTrustZeroTrustGroupsExclude() {
}

type ZeroTrustGroupsExcludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                   `json:"keys_url,required"`
	JSON    zeroTrustGroupsExcludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustGroupsExcludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustGroupsExcludeAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustGroupsExcludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustGroupsExcludeAccessCountryRule struct {
	Geo  ZeroTrustGroupsExcludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustGroupsExcludeAccessCountryRuleJSON `json:"-"`
}

// zeroTrustGroupsExcludeAccessCountryRuleJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsExcludeAccessCountryRule]
type zeroTrustGroupsExcludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsExcludeAccessCountryRule) implementsZeroTrustZeroTrustGroupsExclude() {}

type ZeroTrustGroupsExcludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                         `json:"country_code,required"`
	JSON        zeroTrustGroupsExcludeAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustGroupsExcludeAccessCountryRuleGeoJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsExcludeAccessCountryRuleGeo]
type zeroTrustGroupsExcludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustGroupsExcludeAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustGroupsExcludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustGroupsExcludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustGroupsExcludeAccessAuthenticationMethodRuleJSON contains the JSON
// metadata for the struct [ZeroTrustGroupsExcludeAccessAuthenticationMethodRule]
type zeroTrustGroupsExcludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsExcludeAccessAuthenticationMethodRule) implementsZeroTrustZeroTrustGroupsExclude() {
}

type ZeroTrustGroupsExcludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                             `json:"auth_method,required"`
	JSON       zeroTrustGroupsExcludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustGroupsExcludeAccessAuthenticationMethodRuleAuthMethodJSON contains the
// JSON metadata for the struct
// [ZeroTrustGroupsExcludeAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustGroupsExcludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustGroupsExcludeAccessDevicePostureRule struct {
	DevicePosture ZeroTrustGroupsExcludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustGroupsExcludeAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustGroupsExcludeAccessDevicePostureRuleJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsExcludeAccessDevicePostureRule]
type zeroTrustGroupsExcludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsExcludeAccessDevicePostureRule) implementsZeroTrustZeroTrustGroupsExclude() {}

type ZeroTrustGroupsExcludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                         `json:"integration_uid,required"`
	JSON           zeroTrustGroupsExcludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustGroupsExcludeAccessDevicePostureRuleDevicePostureJSON contains the JSON
// metadata for the struct
// [ZeroTrustGroupsExcludeAccessDevicePostureRuleDevicePosture]
type zeroTrustGroupsExcludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustGroupsExcludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsExcludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by [zero_trust.ZeroTrustGroupsIncludeAccessEmailRule],
// [zero_trust.ZeroTrustGroupsIncludeAccessEmailListRule],
// [zero_trust.ZeroTrustGroupsIncludeAccessDomainRule],
// [zero_trust.ZeroTrustGroupsIncludeAccessEveryoneRule],
// [zero_trust.ZeroTrustGroupsIncludeAccessIPRule],
// [zero_trust.ZeroTrustGroupsIncludeAccessIPListRule],
// [zero_trust.ZeroTrustGroupsIncludeAccessCertificateRule],
// [zero_trust.ZeroTrustGroupsIncludeAccessAccessGroupRule],
// [zero_trust.ZeroTrustGroupsIncludeAccessAzureGroupRule],
// [zero_trust.ZeroTrustGroupsIncludeAccessGitHubOrganizationRule],
// [zero_trust.ZeroTrustGroupsIncludeAccessGsuiteGroupRule],
// [zero_trust.ZeroTrustGroupsIncludeAccessOktaGroupRule],
// [zero_trust.ZeroTrustGroupsIncludeAccessSamlGroupRule],
// [zero_trust.ZeroTrustGroupsIncludeAccessServiceTokenRule],
// [zero_trust.ZeroTrustGroupsIncludeAccessAnyValidServiceTokenRule],
// [zero_trust.ZeroTrustGroupsIncludeAccessExternalEvaluationRule],
// [zero_trust.ZeroTrustGroupsIncludeAccessCountryRule],
// [zero_trust.ZeroTrustGroupsIncludeAccessAuthenticationMethodRule] or
// [zero_trust.ZeroTrustGroupsIncludeAccessDevicePostureRule].
type ZeroTrustGroupsInclude interface {
	implementsZeroTrustZeroTrustGroupsInclude()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustGroupsInclude)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIncludeAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIncludeAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIncludeAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIncludeAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIncludeAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIncludeAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIncludeAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIncludeAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIncludeAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIncludeAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIncludeAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIncludeAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIncludeAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIncludeAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIncludeAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIncludeAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIncludeAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIncludeAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIncludeAccessDevicePostureRule{}),
		},
	)
}

// Matches a specific email.
type ZeroTrustGroupsIncludeAccessEmailRule struct {
	Email ZeroTrustGroupsIncludeAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustGroupsIncludeAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustGroupsIncludeAccessEmailRuleJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsIncludeAccessEmailRule]
type zeroTrustGroupsIncludeAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIncludeAccessEmailRule) implementsZeroTrustZeroTrustGroupsInclude() {}

type ZeroTrustGroupsIncludeAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                         `json:"email,required" format:"email"`
	JSON  zeroTrustGroupsIncludeAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustGroupsIncludeAccessEmailRuleEmailJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsIncludeAccessEmailRuleEmail]
type zeroTrustGroupsIncludeAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustGroupsIncludeAccessEmailListRule struct {
	EmailList ZeroTrustGroupsIncludeAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustGroupsIncludeAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustGroupsIncludeAccessEmailListRuleJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsIncludeAccessEmailListRule]
type zeroTrustGroupsIncludeAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIncludeAccessEmailListRule) implementsZeroTrustZeroTrustGroupsInclude() {}

type ZeroTrustGroupsIncludeAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                 `json:"id,required"`
	JSON zeroTrustGroupsIncludeAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustGroupsIncludeAccessEmailListRuleEmailListJSON contains the JSON
// metadata for the struct [ZeroTrustGroupsIncludeAccessEmailListRuleEmailList]
type zeroTrustGroupsIncludeAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustGroupsIncludeAccessDomainRule struct {
	EmailDomain ZeroTrustGroupsIncludeAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustGroupsIncludeAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustGroupsIncludeAccessDomainRuleJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsIncludeAccessDomainRule]
type zeroTrustGroupsIncludeAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIncludeAccessDomainRule) implementsZeroTrustZeroTrustGroupsInclude() {}

type ZeroTrustGroupsIncludeAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                `json:"domain,required"`
	JSON   zeroTrustGroupsIncludeAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustGroupsIncludeAccessDomainRuleEmailDomainJSON contains the JSON metadata
// for the struct [ZeroTrustGroupsIncludeAccessDomainRuleEmailDomain]
type zeroTrustGroupsIncludeAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustGroupsIncludeAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                  `json:"everyone,required"`
	JSON     zeroTrustGroupsIncludeAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustGroupsIncludeAccessEveryoneRuleJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsIncludeAccessEveryoneRule]
type zeroTrustGroupsIncludeAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIncludeAccessEveryoneRule) implementsZeroTrustZeroTrustGroupsInclude() {}

// Matches an IP address block.
type ZeroTrustGroupsIncludeAccessIPRule struct {
	IP   ZeroTrustGroupsIncludeAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustGroupsIncludeAccessIPRuleJSON `json:"-"`
}

// zeroTrustGroupsIncludeAccessIPRuleJSON contains the JSON metadata for the struct
// [ZeroTrustGroupsIncludeAccessIPRule]
type zeroTrustGroupsIncludeAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIncludeAccessIPRule) implementsZeroTrustZeroTrustGroupsInclude() {}

type ZeroTrustGroupsIncludeAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                   `json:"ip,required"`
	JSON zeroTrustGroupsIncludeAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustGroupsIncludeAccessIPRuleIPJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsIncludeAccessIPRuleIP]
type zeroTrustGroupsIncludeAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustGroupsIncludeAccessIPListRule struct {
	IPList ZeroTrustGroupsIncludeAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustGroupsIncludeAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustGroupsIncludeAccessIPListRuleJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsIncludeAccessIPListRule]
type zeroTrustGroupsIncludeAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIncludeAccessIPListRule) implementsZeroTrustZeroTrustGroupsInclude() {}

type ZeroTrustGroupsIncludeAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                           `json:"id,required"`
	JSON zeroTrustGroupsIncludeAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustGroupsIncludeAccessIPListRuleIPListJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsIncludeAccessIPListRuleIPList]
type zeroTrustGroupsIncludeAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustGroupsIncludeAccessCertificateRule struct {
	Certificate interface{}                                     `json:"certificate,required"`
	JSON        zeroTrustGroupsIncludeAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustGroupsIncludeAccessCertificateRuleJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsIncludeAccessCertificateRule]
type zeroTrustGroupsIncludeAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIncludeAccessCertificateRule) implementsZeroTrustZeroTrustGroupsInclude() {}

// Matches an Access group.
type ZeroTrustGroupsIncludeAccessAccessGroupRule struct {
	Group shared.UnnamedSchemaRef131                      `json:"group,required"`
	JSON  zeroTrustGroupsIncludeAccessAccessGroupRuleJSON `json:"-"`
}

// zeroTrustGroupsIncludeAccessAccessGroupRuleJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsIncludeAccessAccessGroupRule]
type zeroTrustGroupsIncludeAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIncludeAccessAccessGroupRule) implementsZeroTrustZeroTrustGroupsInclude() {}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustGroupsIncludeAccessAzureGroupRule struct {
	AzureAd ZeroTrustGroupsIncludeAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustGroupsIncludeAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustGroupsIncludeAccessAzureGroupRuleJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsIncludeAccessAzureGroupRule]
type zeroTrustGroupsIncludeAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIncludeAccessAzureGroupRule) implementsZeroTrustZeroTrustGroupsInclude() {}

type ZeroTrustGroupsIncludeAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                `json:"connection_id,required"`
	JSON         zeroTrustGroupsIncludeAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustGroupsIncludeAccessAzureGroupRuleAzureAdJSON contains the JSON metadata
// for the struct [ZeroTrustGroupsIncludeAccessAzureGroupRuleAzureAd]
type zeroTrustGroupsIncludeAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustGroupsIncludeAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustGroupsIncludeAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustGroupsIncludeAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustGroupsIncludeAccessGitHubOrganizationRuleJSON contains the JSON
// metadata for the struct [ZeroTrustGroupsIncludeAccessGitHubOrganizationRule]
type zeroTrustGroupsIncludeAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIncludeAccessGitHubOrganizationRule) implementsZeroTrustZeroTrustGroupsInclude() {
}

type ZeroTrustGroupsIncludeAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                   `json:"name,required"`
	JSON zeroTrustGroupsIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustGroupsIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustGroupsIncludeAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustGroupsIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustGroupsIncludeAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustGroupsIncludeAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustGroupsIncludeAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustGroupsIncludeAccessGsuiteGroupRuleJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsIncludeAccessGsuiteGroupRule]
type zeroTrustGroupsIncludeAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIncludeAccessGsuiteGroupRule) implementsZeroTrustZeroTrustGroupsInclude() {}

type ZeroTrustGroupsIncludeAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                `json:"email,required"`
	JSON  zeroTrustGroupsIncludeAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustGroupsIncludeAccessGsuiteGroupRuleGsuiteJSON contains the JSON metadata
// for the struct [ZeroTrustGroupsIncludeAccessGsuiteGroupRuleGsuite]
type zeroTrustGroupsIncludeAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustGroupsIncludeAccessOktaGroupRule struct {
	Okta ZeroTrustGroupsIncludeAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustGroupsIncludeAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustGroupsIncludeAccessOktaGroupRuleJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsIncludeAccessOktaGroupRule]
type zeroTrustGroupsIncludeAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIncludeAccessOktaGroupRule) implementsZeroTrustZeroTrustGroupsInclude() {}

type ZeroTrustGroupsIncludeAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                            `json:"email,required"`
	JSON  zeroTrustGroupsIncludeAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustGroupsIncludeAccessOktaGroupRuleOktaJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsIncludeAccessOktaGroupRuleOkta]
type zeroTrustGroupsIncludeAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustGroupsIncludeAccessSamlGroupRule struct {
	Saml ZeroTrustGroupsIncludeAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustGroupsIncludeAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustGroupsIncludeAccessSamlGroupRuleJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsIncludeAccessSamlGroupRule]
type zeroTrustGroupsIncludeAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIncludeAccessSamlGroupRule) implementsZeroTrustZeroTrustGroupsInclude() {}

type ZeroTrustGroupsIncludeAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                            `json:"attribute_value,required"`
	JSON           zeroTrustGroupsIncludeAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustGroupsIncludeAccessSamlGroupRuleSamlJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsIncludeAccessSamlGroupRuleSaml]
type zeroTrustGroupsIncludeAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustGroupsIncludeAccessServiceTokenRule struct {
	ServiceToken ZeroTrustGroupsIncludeAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustGroupsIncludeAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustGroupsIncludeAccessServiceTokenRuleJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsIncludeAccessServiceTokenRule]
type zeroTrustGroupsIncludeAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIncludeAccessServiceTokenRule) implementsZeroTrustZeroTrustGroupsInclude() {}

type ZeroTrustGroupsIncludeAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                       `json:"token_id,required"`
	JSON    zeroTrustGroupsIncludeAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustGroupsIncludeAccessServiceTokenRuleServiceTokenJSON contains the JSON
// metadata for the struct
// [ZeroTrustGroupsIncludeAccessServiceTokenRuleServiceToken]
type zeroTrustGroupsIncludeAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustGroupsIncludeAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                              `json:"any_valid_service_token,required"`
	JSON                 zeroTrustGroupsIncludeAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustGroupsIncludeAccessAnyValidServiceTokenRuleJSON contains the JSON
// metadata for the struct [ZeroTrustGroupsIncludeAccessAnyValidServiceTokenRule]
type zeroTrustGroupsIncludeAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIncludeAccessAnyValidServiceTokenRule) implementsZeroTrustZeroTrustGroupsInclude() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustGroupsIncludeAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustGroupsIncludeAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustGroupsIncludeAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustGroupsIncludeAccessExternalEvaluationRuleJSON contains the JSON
// metadata for the struct [ZeroTrustGroupsIncludeAccessExternalEvaluationRule]
type zeroTrustGroupsIncludeAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIncludeAccessExternalEvaluationRule) implementsZeroTrustZeroTrustGroupsInclude() {
}

type ZeroTrustGroupsIncludeAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                   `json:"keys_url,required"`
	JSON    zeroTrustGroupsIncludeAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustGroupsIncludeAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustGroupsIncludeAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustGroupsIncludeAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustGroupsIncludeAccessCountryRule struct {
	Geo  ZeroTrustGroupsIncludeAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustGroupsIncludeAccessCountryRuleJSON `json:"-"`
}

// zeroTrustGroupsIncludeAccessCountryRuleJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsIncludeAccessCountryRule]
type zeroTrustGroupsIncludeAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIncludeAccessCountryRule) implementsZeroTrustZeroTrustGroupsInclude() {}

type ZeroTrustGroupsIncludeAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                         `json:"country_code,required"`
	JSON        zeroTrustGroupsIncludeAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustGroupsIncludeAccessCountryRuleGeoJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsIncludeAccessCountryRuleGeo]
type zeroTrustGroupsIncludeAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustGroupsIncludeAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustGroupsIncludeAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustGroupsIncludeAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustGroupsIncludeAccessAuthenticationMethodRuleJSON contains the JSON
// metadata for the struct [ZeroTrustGroupsIncludeAccessAuthenticationMethodRule]
type zeroTrustGroupsIncludeAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIncludeAccessAuthenticationMethodRule) implementsZeroTrustZeroTrustGroupsInclude() {
}

type ZeroTrustGroupsIncludeAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                             `json:"auth_method,required"`
	JSON       zeroTrustGroupsIncludeAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustGroupsIncludeAccessAuthenticationMethodRuleAuthMethodJSON contains the
// JSON metadata for the struct
// [ZeroTrustGroupsIncludeAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustGroupsIncludeAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustGroupsIncludeAccessDevicePostureRule struct {
	DevicePosture ZeroTrustGroupsIncludeAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustGroupsIncludeAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustGroupsIncludeAccessDevicePostureRuleJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsIncludeAccessDevicePostureRule]
type zeroTrustGroupsIncludeAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIncludeAccessDevicePostureRule) implementsZeroTrustZeroTrustGroupsInclude() {}

type ZeroTrustGroupsIncludeAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                         `json:"integration_uid,required"`
	JSON           zeroTrustGroupsIncludeAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustGroupsIncludeAccessDevicePostureRuleDevicePostureJSON contains the JSON
// metadata for the struct
// [ZeroTrustGroupsIncludeAccessDevicePostureRuleDevicePosture]
type zeroTrustGroupsIncludeAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustGroupsIncludeAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIncludeAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by [zero_trust.ZeroTrustGroupsIsDefaultAccessEmailRule],
// [zero_trust.ZeroTrustGroupsIsDefaultAccessEmailListRule],
// [zero_trust.ZeroTrustGroupsIsDefaultAccessDomainRule],
// [zero_trust.ZeroTrustGroupsIsDefaultAccessEveryoneRule],
// [zero_trust.ZeroTrustGroupsIsDefaultAccessIPRule],
// [zero_trust.ZeroTrustGroupsIsDefaultAccessIPListRule],
// [zero_trust.ZeroTrustGroupsIsDefaultAccessCertificateRule],
// [zero_trust.ZeroTrustGroupsIsDefaultAccessAccessGroupRule],
// [zero_trust.ZeroTrustGroupsIsDefaultAccessAzureGroupRule],
// [zero_trust.ZeroTrustGroupsIsDefaultAccessGitHubOrganizationRule],
// [zero_trust.ZeroTrustGroupsIsDefaultAccessGsuiteGroupRule],
// [zero_trust.ZeroTrustGroupsIsDefaultAccessOktaGroupRule],
// [zero_trust.ZeroTrustGroupsIsDefaultAccessSamlGroupRule],
// [zero_trust.ZeroTrustGroupsIsDefaultAccessServiceTokenRule],
// [zero_trust.ZeroTrustGroupsIsDefaultAccessAnyValidServiceTokenRule],
// [zero_trust.ZeroTrustGroupsIsDefaultAccessExternalEvaluationRule],
// [zero_trust.ZeroTrustGroupsIsDefaultAccessCountryRule],
// [zero_trust.ZeroTrustGroupsIsDefaultAccessAuthenticationMethodRule] or
// [zero_trust.ZeroTrustGroupsIsDefaultAccessDevicePostureRule].
type ZeroTrustGroupsIsDefault interface {
	implementsZeroTrustZeroTrustGroupsIsDefault()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustGroupsIsDefault)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIsDefaultAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIsDefaultAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIsDefaultAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIsDefaultAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIsDefaultAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIsDefaultAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIsDefaultAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIsDefaultAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIsDefaultAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIsDefaultAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIsDefaultAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIsDefaultAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIsDefaultAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIsDefaultAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIsDefaultAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIsDefaultAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIsDefaultAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIsDefaultAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsIsDefaultAccessDevicePostureRule{}),
		},
	)
}

// Matches a specific email.
type ZeroTrustGroupsIsDefaultAccessEmailRule struct {
	Email ZeroTrustGroupsIsDefaultAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustGroupsIsDefaultAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessEmailRuleJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsIsDefaultAccessEmailRule]
type zeroTrustGroupsIsDefaultAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIsDefaultAccessEmailRule) implementsZeroTrustZeroTrustGroupsIsDefault() {}

type ZeroTrustGroupsIsDefaultAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                           `json:"email,required" format:"email"`
	JSON  zeroTrustGroupsIsDefaultAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessEmailRuleEmailJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsIsDefaultAccessEmailRuleEmail]
type zeroTrustGroupsIsDefaultAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustGroupsIsDefaultAccessEmailListRule struct {
	EmailList ZeroTrustGroupsIsDefaultAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustGroupsIsDefaultAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessEmailListRuleJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsIsDefaultAccessEmailListRule]
type zeroTrustGroupsIsDefaultAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIsDefaultAccessEmailListRule) implementsZeroTrustZeroTrustGroupsIsDefault() {}

type ZeroTrustGroupsIsDefaultAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                   `json:"id,required"`
	JSON zeroTrustGroupsIsDefaultAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessEmailListRuleEmailListJSON contains the JSON
// metadata for the struct [ZeroTrustGroupsIsDefaultAccessEmailListRuleEmailList]
type zeroTrustGroupsIsDefaultAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustGroupsIsDefaultAccessDomainRule struct {
	EmailDomain ZeroTrustGroupsIsDefaultAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustGroupsIsDefaultAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessDomainRuleJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsIsDefaultAccessDomainRule]
type zeroTrustGroupsIsDefaultAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIsDefaultAccessDomainRule) implementsZeroTrustZeroTrustGroupsIsDefault() {}

type ZeroTrustGroupsIsDefaultAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                  `json:"domain,required"`
	JSON   zeroTrustGroupsIsDefaultAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessDomainRuleEmailDomainJSON contains the JSON
// metadata for the struct [ZeroTrustGroupsIsDefaultAccessDomainRuleEmailDomain]
type zeroTrustGroupsIsDefaultAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustGroupsIsDefaultAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                    `json:"everyone,required"`
	JSON     zeroTrustGroupsIsDefaultAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessEveryoneRuleJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsIsDefaultAccessEveryoneRule]
type zeroTrustGroupsIsDefaultAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIsDefaultAccessEveryoneRule) implementsZeroTrustZeroTrustGroupsIsDefault() {}

// Matches an IP address block.
type ZeroTrustGroupsIsDefaultAccessIPRule struct {
	IP   ZeroTrustGroupsIsDefaultAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustGroupsIsDefaultAccessIPRuleJSON `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessIPRuleJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsIsDefaultAccessIPRule]
type zeroTrustGroupsIsDefaultAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIsDefaultAccessIPRule) implementsZeroTrustZeroTrustGroupsIsDefault() {}

type ZeroTrustGroupsIsDefaultAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                     `json:"ip,required"`
	JSON zeroTrustGroupsIsDefaultAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessIPRuleIPJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsIsDefaultAccessIPRuleIP]
type zeroTrustGroupsIsDefaultAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustGroupsIsDefaultAccessIPListRule struct {
	IPList ZeroTrustGroupsIsDefaultAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustGroupsIsDefaultAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessIPListRuleJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsIsDefaultAccessIPListRule]
type zeroTrustGroupsIsDefaultAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIsDefaultAccessIPListRule) implementsZeroTrustZeroTrustGroupsIsDefault() {}

type ZeroTrustGroupsIsDefaultAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                             `json:"id,required"`
	JSON zeroTrustGroupsIsDefaultAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessIPListRuleIPListJSON contains the JSON metadata
// for the struct [ZeroTrustGroupsIsDefaultAccessIPListRuleIPList]
type zeroTrustGroupsIsDefaultAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustGroupsIsDefaultAccessCertificateRule struct {
	Certificate interface{}                                       `json:"certificate,required"`
	JSON        zeroTrustGroupsIsDefaultAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessCertificateRuleJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsIsDefaultAccessCertificateRule]
type zeroTrustGroupsIsDefaultAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIsDefaultAccessCertificateRule) implementsZeroTrustZeroTrustGroupsIsDefault() {
}

// Matches an Access group.
type ZeroTrustGroupsIsDefaultAccessAccessGroupRule struct {
	Group shared.UnnamedSchemaRef131                        `json:"group,required"`
	JSON  zeroTrustGroupsIsDefaultAccessAccessGroupRuleJSON `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessAccessGroupRuleJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsIsDefaultAccessAccessGroupRule]
type zeroTrustGroupsIsDefaultAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIsDefaultAccessAccessGroupRule) implementsZeroTrustZeroTrustGroupsIsDefault() {
}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustGroupsIsDefaultAccessAzureGroupRule struct {
	AzureAd ZeroTrustGroupsIsDefaultAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustGroupsIsDefaultAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessAzureGroupRuleJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsIsDefaultAccessAzureGroupRule]
type zeroTrustGroupsIsDefaultAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIsDefaultAccessAzureGroupRule) implementsZeroTrustZeroTrustGroupsIsDefault() {}

type ZeroTrustGroupsIsDefaultAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                  `json:"connection_id,required"`
	JSON         zeroTrustGroupsIsDefaultAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessAzureGroupRuleAzureAdJSON contains the JSON
// metadata for the struct [ZeroTrustGroupsIsDefaultAccessAzureGroupRuleAzureAd]
type zeroTrustGroupsIsDefaultAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustGroupsIsDefaultAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustGroupsIsDefaultAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustGroupsIsDefaultAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessGitHubOrganizationRuleJSON contains the JSON
// metadata for the struct [ZeroTrustGroupsIsDefaultAccessGitHubOrganizationRule]
type zeroTrustGroupsIsDefaultAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIsDefaultAccessGitHubOrganizationRule) implementsZeroTrustZeroTrustGroupsIsDefault() {
}

type ZeroTrustGroupsIsDefaultAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                     `json:"name,required"`
	JSON zeroTrustGroupsIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustGroupsIsDefaultAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustGroupsIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustGroupsIsDefaultAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustGroupsIsDefaultAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustGroupsIsDefaultAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessGsuiteGroupRuleJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsIsDefaultAccessGsuiteGroupRule]
type zeroTrustGroupsIsDefaultAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIsDefaultAccessGsuiteGroupRule) implementsZeroTrustZeroTrustGroupsIsDefault() {
}

type ZeroTrustGroupsIsDefaultAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                  `json:"email,required"`
	JSON  zeroTrustGroupsIsDefaultAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessGsuiteGroupRuleGsuiteJSON contains the JSON
// metadata for the struct [ZeroTrustGroupsIsDefaultAccessGsuiteGroupRuleGsuite]
type zeroTrustGroupsIsDefaultAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustGroupsIsDefaultAccessOktaGroupRule struct {
	Okta ZeroTrustGroupsIsDefaultAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustGroupsIsDefaultAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessOktaGroupRuleJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsIsDefaultAccessOktaGroupRule]
type zeroTrustGroupsIsDefaultAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIsDefaultAccessOktaGroupRule) implementsZeroTrustZeroTrustGroupsIsDefault() {}

type ZeroTrustGroupsIsDefaultAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                              `json:"email,required"`
	JSON  zeroTrustGroupsIsDefaultAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessOktaGroupRuleOktaJSON contains the JSON metadata
// for the struct [ZeroTrustGroupsIsDefaultAccessOktaGroupRuleOkta]
type zeroTrustGroupsIsDefaultAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustGroupsIsDefaultAccessSamlGroupRule struct {
	Saml ZeroTrustGroupsIsDefaultAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustGroupsIsDefaultAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessSamlGroupRuleJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsIsDefaultAccessSamlGroupRule]
type zeroTrustGroupsIsDefaultAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIsDefaultAccessSamlGroupRule) implementsZeroTrustZeroTrustGroupsIsDefault() {}

type ZeroTrustGroupsIsDefaultAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                              `json:"attribute_value,required"`
	JSON           zeroTrustGroupsIsDefaultAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessSamlGroupRuleSamlJSON contains the JSON metadata
// for the struct [ZeroTrustGroupsIsDefaultAccessSamlGroupRuleSaml]
type zeroTrustGroupsIsDefaultAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustGroupsIsDefaultAccessServiceTokenRule struct {
	ServiceToken ZeroTrustGroupsIsDefaultAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustGroupsIsDefaultAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessServiceTokenRuleJSON contains the JSON metadata
// for the struct [ZeroTrustGroupsIsDefaultAccessServiceTokenRule]
type zeroTrustGroupsIsDefaultAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIsDefaultAccessServiceTokenRule) implementsZeroTrustZeroTrustGroupsIsDefault() {
}

type ZeroTrustGroupsIsDefaultAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                         `json:"token_id,required"`
	JSON    zeroTrustGroupsIsDefaultAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessServiceTokenRuleServiceTokenJSON contains the JSON
// metadata for the struct
// [ZeroTrustGroupsIsDefaultAccessServiceTokenRuleServiceToken]
type zeroTrustGroupsIsDefaultAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustGroupsIsDefaultAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                                `json:"any_valid_service_token,required"`
	JSON                 zeroTrustGroupsIsDefaultAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessAnyValidServiceTokenRuleJSON contains the JSON
// metadata for the struct [ZeroTrustGroupsIsDefaultAccessAnyValidServiceTokenRule]
type zeroTrustGroupsIsDefaultAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIsDefaultAccessAnyValidServiceTokenRule) implementsZeroTrustZeroTrustGroupsIsDefault() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustGroupsIsDefaultAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustGroupsIsDefaultAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustGroupsIsDefaultAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessExternalEvaluationRuleJSON contains the JSON
// metadata for the struct [ZeroTrustGroupsIsDefaultAccessExternalEvaluationRule]
type zeroTrustGroupsIsDefaultAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIsDefaultAccessExternalEvaluationRule) implementsZeroTrustZeroTrustGroupsIsDefault() {
}

type ZeroTrustGroupsIsDefaultAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                     `json:"keys_url,required"`
	JSON    zeroTrustGroupsIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustGroupsIsDefaultAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustGroupsIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustGroupsIsDefaultAccessCountryRule struct {
	Geo  ZeroTrustGroupsIsDefaultAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustGroupsIsDefaultAccessCountryRuleJSON `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessCountryRuleJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsIsDefaultAccessCountryRule]
type zeroTrustGroupsIsDefaultAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIsDefaultAccessCountryRule) implementsZeroTrustZeroTrustGroupsIsDefault() {}

type ZeroTrustGroupsIsDefaultAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                           `json:"country_code,required"`
	JSON        zeroTrustGroupsIsDefaultAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessCountryRuleGeoJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsIsDefaultAccessCountryRuleGeo]
type zeroTrustGroupsIsDefaultAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustGroupsIsDefaultAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustGroupsIsDefaultAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustGroupsIsDefaultAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessAuthenticationMethodRuleJSON contains the JSON
// metadata for the struct [ZeroTrustGroupsIsDefaultAccessAuthenticationMethodRule]
type zeroTrustGroupsIsDefaultAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIsDefaultAccessAuthenticationMethodRule) implementsZeroTrustZeroTrustGroupsIsDefault() {
}

type ZeroTrustGroupsIsDefaultAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                               `json:"auth_method,required"`
	JSON       zeroTrustGroupsIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON contains
// the JSON metadata for the struct
// [ZeroTrustGroupsIsDefaultAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustGroupsIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustGroupsIsDefaultAccessDevicePostureRule struct {
	DevicePosture ZeroTrustGroupsIsDefaultAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustGroupsIsDefaultAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessDevicePostureRuleJSON contains the JSON metadata
// for the struct [ZeroTrustGroupsIsDefaultAccessDevicePostureRule]
type zeroTrustGroupsIsDefaultAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsIsDefaultAccessDevicePostureRule) implementsZeroTrustZeroTrustGroupsIsDefault() {
}

type ZeroTrustGroupsIsDefaultAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                           `json:"integration_uid,required"`
	JSON           zeroTrustGroupsIsDefaultAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustGroupsIsDefaultAccessDevicePostureRuleDevicePostureJSON contains the
// JSON metadata for the struct
// [ZeroTrustGroupsIsDefaultAccessDevicePostureRuleDevicePosture]
type zeroTrustGroupsIsDefaultAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustGroupsIsDefaultAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsIsDefaultAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
	return r.raw
}

// Matches a specific email.
//
// Union satisfied by [zero_trust.ZeroTrustGroupsRequireAccessEmailRule],
// [zero_trust.ZeroTrustGroupsRequireAccessEmailListRule],
// [zero_trust.ZeroTrustGroupsRequireAccessDomainRule],
// [zero_trust.ZeroTrustGroupsRequireAccessEveryoneRule],
// [zero_trust.ZeroTrustGroupsRequireAccessIPRule],
// [zero_trust.ZeroTrustGroupsRequireAccessIPListRule],
// [zero_trust.ZeroTrustGroupsRequireAccessCertificateRule],
// [zero_trust.ZeroTrustGroupsRequireAccessAccessGroupRule],
// [zero_trust.ZeroTrustGroupsRequireAccessAzureGroupRule],
// [zero_trust.ZeroTrustGroupsRequireAccessGitHubOrganizationRule],
// [zero_trust.ZeroTrustGroupsRequireAccessGsuiteGroupRule],
// [zero_trust.ZeroTrustGroupsRequireAccessOktaGroupRule],
// [zero_trust.ZeroTrustGroupsRequireAccessSamlGroupRule],
// [zero_trust.ZeroTrustGroupsRequireAccessServiceTokenRule],
// [zero_trust.ZeroTrustGroupsRequireAccessAnyValidServiceTokenRule],
// [zero_trust.ZeroTrustGroupsRequireAccessExternalEvaluationRule],
// [zero_trust.ZeroTrustGroupsRequireAccessCountryRule],
// [zero_trust.ZeroTrustGroupsRequireAccessAuthenticationMethodRule] or
// [zero_trust.ZeroTrustGroupsRequireAccessDevicePostureRule].
type ZeroTrustGroupsRequire interface {
	implementsZeroTrustZeroTrustGroupsRequire()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustGroupsRequire)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsRequireAccessEmailRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsRequireAccessEmailListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsRequireAccessDomainRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsRequireAccessEveryoneRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsRequireAccessIPRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsRequireAccessIPListRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsRequireAccessCertificateRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsRequireAccessAccessGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsRequireAccessAzureGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsRequireAccessGitHubOrganizationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsRequireAccessGsuiteGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsRequireAccessOktaGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsRequireAccessSamlGroupRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsRequireAccessServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsRequireAccessAnyValidServiceTokenRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsRequireAccessExternalEvaluationRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsRequireAccessCountryRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsRequireAccessAuthenticationMethodRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustGroupsRequireAccessDevicePostureRule{}),
		},
	)
}

// Matches a specific email.
type ZeroTrustGroupsRequireAccessEmailRule struct {
	Email ZeroTrustGroupsRequireAccessEmailRuleEmail `json:"email,required"`
	JSON  zeroTrustGroupsRequireAccessEmailRuleJSON  `json:"-"`
}

// zeroTrustGroupsRequireAccessEmailRuleJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsRequireAccessEmailRule]
type zeroTrustGroupsRequireAccessEmailRuleJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessEmailRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessEmailRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsRequireAccessEmailRule) implementsZeroTrustZeroTrustGroupsRequire() {}

type ZeroTrustGroupsRequireAccessEmailRuleEmail struct {
	// The email of the user.
	Email string                                         `json:"email,required" format:"email"`
	JSON  zeroTrustGroupsRequireAccessEmailRuleEmailJSON `json:"-"`
}

// zeroTrustGroupsRequireAccessEmailRuleEmailJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsRequireAccessEmailRuleEmail]
type zeroTrustGroupsRequireAccessEmailRuleEmailJSON struct {
	Email       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessEmailRuleEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessEmailRuleEmailJSON) RawJSON() string {
	return r.raw
}

// Matches an email address from a list.
type ZeroTrustGroupsRequireAccessEmailListRule struct {
	EmailList ZeroTrustGroupsRequireAccessEmailListRuleEmailList `json:"email_list,required"`
	JSON      zeroTrustGroupsRequireAccessEmailListRuleJSON      `json:"-"`
}

// zeroTrustGroupsRequireAccessEmailListRuleJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsRequireAccessEmailListRule]
type zeroTrustGroupsRequireAccessEmailListRuleJSON struct {
	EmailList   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessEmailListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessEmailListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsRequireAccessEmailListRule) implementsZeroTrustZeroTrustGroupsRequire() {}

type ZeroTrustGroupsRequireAccessEmailListRuleEmailList struct {
	// The ID of a previously created email list.
	ID   string                                                 `json:"id,required"`
	JSON zeroTrustGroupsRequireAccessEmailListRuleEmailListJSON `json:"-"`
}

// zeroTrustGroupsRequireAccessEmailListRuleEmailListJSON contains the JSON
// metadata for the struct [ZeroTrustGroupsRequireAccessEmailListRuleEmailList]
type zeroTrustGroupsRequireAccessEmailListRuleEmailListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessEmailListRuleEmailList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessEmailListRuleEmailListJSON) RawJSON() string {
	return r.raw
}

// Match an entire email domain.
type ZeroTrustGroupsRequireAccessDomainRule struct {
	EmailDomain ZeroTrustGroupsRequireAccessDomainRuleEmailDomain `json:"email_domain,required"`
	JSON        zeroTrustGroupsRequireAccessDomainRuleJSON        `json:"-"`
}

// zeroTrustGroupsRequireAccessDomainRuleJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsRequireAccessDomainRule]
type zeroTrustGroupsRequireAccessDomainRuleJSON struct {
	EmailDomain apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessDomainRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessDomainRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsRequireAccessDomainRule) implementsZeroTrustZeroTrustGroupsRequire() {}

type ZeroTrustGroupsRequireAccessDomainRuleEmailDomain struct {
	// The email domain to match.
	Domain string                                                `json:"domain,required"`
	JSON   zeroTrustGroupsRequireAccessDomainRuleEmailDomainJSON `json:"-"`
}

// zeroTrustGroupsRequireAccessDomainRuleEmailDomainJSON contains the JSON metadata
// for the struct [ZeroTrustGroupsRequireAccessDomainRuleEmailDomain]
type zeroTrustGroupsRequireAccessDomainRuleEmailDomainJSON struct {
	Domain      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessDomainRuleEmailDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessDomainRuleEmailDomainJSON) RawJSON() string {
	return r.raw
}

// Matches everyone.
type ZeroTrustGroupsRequireAccessEveryoneRule struct {
	// An empty object which matches on all users.
	Everyone interface{}                                  `json:"everyone,required"`
	JSON     zeroTrustGroupsRequireAccessEveryoneRuleJSON `json:"-"`
}

// zeroTrustGroupsRequireAccessEveryoneRuleJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsRequireAccessEveryoneRule]
type zeroTrustGroupsRequireAccessEveryoneRuleJSON struct {
	Everyone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessEveryoneRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessEveryoneRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsRequireAccessEveryoneRule) implementsZeroTrustZeroTrustGroupsRequire() {}

// Matches an IP address block.
type ZeroTrustGroupsRequireAccessIPRule struct {
	IP   ZeroTrustGroupsRequireAccessIPRuleIP   `json:"ip,required"`
	JSON zeroTrustGroupsRequireAccessIPRuleJSON `json:"-"`
}

// zeroTrustGroupsRequireAccessIPRuleJSON contains the JSON metadata for the struct
// [ZeroTrustGroupsRequireAccessIPRule]
type zeroTrustGroupsRequireAccessIPRuleJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessIPRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessIPRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsRequireAccessIPRule) implementsZeroTrustZeroTrustGroupsRequire() {}

type ZeroTrustGroupsRequireAccessIPRuleIP struct {
	// An IPv4 or IPv6 CIDR block.
	IP   string                                   `json:"ip,required"`
	JSON zeroTrustGroupsRequireAccessIPRuleIPJSON `json:"-"`
}

// zeroTrustGroupsRequireAccessIPRuleIPJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsRequireAccessIPRuleIP]
type zeroTrustGroupsRequireAccessIPRuleIPJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessIPRuleIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessIPRuleIPJSON) RawJSON() string {
	return r.raw
}

// Matches an IP address from a list.
type ZeroTrustGroupsRequireAccessIPListRule struct {
	IPList ZeroTrustGroupsRequireAccessIPListRuleIPList `json:"ip_list,required"`
	JSON   zeroTrustGroupsRequireAccessIPListRuleJSON   `json:"-"`
}

// zeroTrustGroupsRequireAccessIPListRuleJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsRequireAccessIPListRule]
type zeroTrustGroupsRequireAccessIPListRuleJSON struct {
	IPList      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessIPListRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessIPListRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsRequireAccessIPListRule) implementsZeroTrustZeroTrustGroupsRequire() {}

type ZeroTrustGroupsRequireAccessIPListRuleIPList struct {
	// The ID of a previously created IP list.
	ID   string                                           `json:"id,required"`
	JSON zeroTrustGroupsRequireAccessIPListRuleIPListJSON `json:"-"`
}

// zeroTrustGroupsRequireAccessIPListRuleIPListJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsRequireAccessIPListRuleIPList]
type zeroTrustGroupsRequireAccessIPListRuleIPListJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessIPListRuleIPList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessIPListRuleIPListJSON) RawJSON() string {
	return r.raw
}

// Matches any valid client certificate.
type ZeroTrustGroupsRequireAccessCertificateRule struct {
	Certificate interface{}                                     `json:"certificate,required"`
	JSON        zeroTrustGroupsRequireAccessCertificateRuleJSON `json:"-"`
}

// zeroTrustGroupsRequireAccessCertificateRuleJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsRequireAccessCertificateRule]
type zeroTrustGroupsRequireAccessCertificateRuleJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessCertificateRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessCertificateRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsRequireAccessCertificateRule) implementsZeroTrustZeroTrustGroupsRequire() {}

// Matches an Access group.
type ZeroTrustGroupsRequireAccessAccessGroupRule struct {
	Group shared.UnnamedSchemaRef131                      `json:"group,required"`
	JSON  zeroTrustGroupsRequireAccessAccessGroupRuleJSON `json:"-"`
}

// zeroTrustGroupsRequireAccessAccessGroupRuleJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsRequireAccessAccessGroupRule]
type zeroTrustGroupsRequireAccessAccessGroupRuleJSON struct {
	Group       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessAccessGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessAccessGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsRequireAccessAccessGroupRule) implementsZeroTrustZeroTrustGroupsRequire() {}

// Matches an Azure group. Requires an Azure identity provider.
type ZeroTrustGroupsRequireAccessAzureGroupRule struct {
	AzureAd ZeroTrustGroupsRequireAccessAzureGroupRuleAzureAd `json:"azureAD,required"`
	JSON    zeroTrustGroupsRequireAccessAzureGroupRuleJSON    `json:"-"`
}

// zeroTrustGroupsRequireAccessAzureGroupRuleJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsRequireAccessAzureGroupRule]
type zeroTrustGroupsRequireAccessAzureGroupRuleJSON struct {
	AzureAd     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessAzureGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessAzureGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsRequireAccessAzureGroupRule) implementsZeroTrustZeroTrustGroupsRequire() {}

type ZeroTrustGroupsRequireAccessAzureGroupRuleAzureAd struct {
	// The ID of an Azure group.
	ID string `json:"id,required"`
	// The ID of your Azure identity provider.
	ConnectionID string                                                `json:"connection_id,required"`
	JSON         zeroTrustGroupsRequireAccessAzureGroupRuleAzureAdJSON `json:"-"`
}

// zeroTrustGroupsRequireAccessAzureGroupRuleAzureAdJSON contains the JSON metadata
// for the struct [ZeroTrustGroupsRequireAccessAzureGroupRuleAzureAd]
type zeroTrustGroupsRequireAccessAzureGroupRuleAzureAdJSON struct {
	ID           apijson.Field
	ConnectionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessAzureGroupRuleAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessAzureGroupRuleAzureAdJSON) RawJSON() string {
	return r.raw
}

// Matches a Github organization. Requires a Github identity provider.
type ZeroTrustGroupsRequireAccessGitHubOrganizationRule struct {
	GitHubOrganization ZeroTrustGroupsRequireAccessGitHubOrganizationRuleGitHubOrganization `json:"github-organization,required"`
	JSON               zeroTrustGroupsRequireAccessGitHubOrganizationRuleJSON               `json:"-"`
}

// zeroTrustGroupsRequireAccessGitHubOrganizationRuleJSON contains the JSON
// metadata for the struct [ZeroTrustGroupsRequireAccessGitHubOrganizationRule]
type zeroTrustGroupsRequireAccessGitHubOrganizationRuleJSON struct {
	GitHubOrganization apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessGitHubOrganizationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessGitHubOrganizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsRequireAccessGitHubOrganizationRule) implementsZeroTrustZeroTrustGroupsRequire() {
}

type ZeroTrustGroupsRequireAccessGitHubOrganizationRuleGitHubOrganization struct {
	// The ID of your Github identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The name of the organization.
	Name string                                                                   `json:"name,required"`
	JSON zeroTrustGroupsRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON `json:"-"`
}

// zeroTrustGroupsRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON
// contains the JSON metadata for the struct
// [ZeroTrustGroupsRequireAccessGitHubOrganizationRuleGitHubOrganization]
type zeroTrustGroupsRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON struct {
	ConnectionID apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessGitHubOrganizationRuleGitHubOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessGitHubOrganizationRuleGitHubOrganizationJSON) RawJSON() string {
	return r.raw
}

// Matches a group in Google Workspace. Requires a Google Workspace identity
// provider.
type ZeroTrustGroupsRequireAccessGsuiteGroupRule struct {
	Gsuite ZeroTrustGroupsRequireAccessGsuiteGroupRuleGsuite `json:"gsuite,required"`
	JSON   zeroTrustGroupsRequireAccessGsuiteGroupRuleJSON   `json:"-"`
}

// zeroTrustGroupsRequireAccessGsuiteGroupRuleJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsRequireAccessGsuiteGroupRule]
type zeroTrustGroupsRequireAccessGsuiteGroupRuleJSON struct {
	Gsuite      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessGsuiteGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessGsuiteGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsRequireAccessGsuiteGroupRule) implementsZeroTrustZeroTrustGroupsRequire() {}

type ZeroTrustGroupsRequireAccessGsuiteGroupRuleGsuite struct {
	// The ID of your Google Workspace identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Google Workspace group.
	Email string                                                `json:"email,required"`
	JSON  zeroTrustGroupsRequireAccessGsuiteGroupRuleGsuiteJSON `json:"-"`
}

// zeroTrustGroupsRequireAccessGsuiteGroupRuleGsuiteJSON contains the JSON metadata
// for the struct [ZeroTrustGroupsRequireAccessGsuiteGroupRuleGsuite]
type zeroTrustGroupsRequireAccessGsuiteGroupRuleGsuiteJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessGsuiteGroupRuleGsuite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessGsuiteGroupRuleGsuiteJSON) RawJSON() string {
	return r.raw
}

// Matches an Okta group. Requires an Okta identity provider.
type ZeroTrustGroupsRequireAccessOktaGroupRule struct {
	Okta ZeroTrustGroupsRequireAccessOktaGroupRuleOkta `json:"okta,required"`
	JSON zeroTrustGroupsRequireAccessOktaGroupRuleJSON `json:"-"`
}

// zeroTrustGroupsRequireAccessOktaGroupRuleJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsRequireAccessOktaGroupRule]
type zeroTrustGroupsRequireAccessOktaGroupRuleJSON struct {
	Okta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessOktaGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessOktaGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsRequireAccessOktaGroupRule) implementsZeroTrustZeroTrustGroupsRequire() {}

type ZeroTrustGroupsRequireAccessOktaGroupRuleOkta struct {
	// The ID of your Okta identity provider.
	ConnectionID string `json:"connection_id,required"`
	// The email of the Okta group.
	Email string                                            `json:"email,required"`
	JSON  zeroTrustGroupsRequireAccessOktaGroupRuleOktaJSON `json:"-"`
}

// zeroTrustGroupsRequireAccessOktaGroupRuleOktaJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsRequireAccessOktaGroupRuleOkta]
type zeroTrustGroupsRequireAccessOktaGroupRuleOktaJSON struct {
	ConnectionID apijson.Field
	Email        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessOktaGroupRuleOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessOktaGroupRuleOktaJSON) RawJSON() string {
	return r.raw
}

// Matches a SAML group. Requires a SAML identity provider.
type ZeroTrustGroupsRequireAccessSamlGroupRule struct {
	Saml ZeroTrustGroupsRequireAccessSamlGroupRuleSaml `json:"saml,required"`
	JSON zeroTrustGroupsRequireAccessSamlGroupRuleJSON `json:"-"`
}

// zeroTrustGroupsRequireAccessSamlGroupRuleJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsRequireAccessSamlGroupRule]
type zeroTrustGroupsRequireAccessSamlGroupRuleJSON struct {
	Saml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessSamlGroupRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessSamlGroupRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsRequireAccessSamlGroupRule) implementsZeroTrustZeroTrustGroupsRequire() {}

type ZeroTrustGroupsRequireAccessSamlGroupRuleSaml struct {
	// The name of the SAML attribute.
	AttributeName string `json:"attribute_name,required"`
	// The SAML attribute value to look for.
	AttributeValue string                                            `json:"attribute_value,required"`
	JSON           zeroTrustGroupsRequireAccessSamlGroupRuleSamlJSON `json:"-"`
}

// zeroTrustGroupsRequireAccessSamlGroupRuleSamlJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsRequireAccessSamlGroupRuleSaml]
type zeroTrustGroupsRequireAccessSamlGroupRuleSamlJSON struct {
	AttributeName  apijson.Field
	AttributeValue apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessSamlGroupRuleSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessSamlGroupRuleSamlJSON) RawJSON() string {
	return r.raw
}

// Matches a specific Access Service Token
type ZeroTrustGroupsRequireAccessServiceTokenRule struct {
	ServiceToken ZeroTrustGroupsRequireAccessServiceTokenRuleServiceToken `json:"service_token,required"`
	JSON         zeroTrustGroupsRequireAccessServiceTokenRuleJSON         `json:"-"`
}

// zeroTrustGroupsRequireAccessServiceTokenRuleJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsRequireAccessServiceTokenRule]
type zeroTrustGroupsRequireAccessServiceTokenRuleJSON struct {
	ServiceToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsRequireAccessServiceTokenRule) implementsZeroTrustZeroTrustGroupsRequire() {}

type ZeroTrustGroupsRequireAccessServiceTokenRuleServiceToken struct {
	// The ID of a Service Token.
	TokenID string                                                       `json:"token_id,required"`
	JSON    zeroTrustGroupsRequireAccessServiceTokenRuleServiceTokenJSON `json:"-"`
}

// zeroTrustGroupsRequireAccessServiceTokenRuleServiceTokenJSON contains the JSON
// metadata for the struct
// [ZeroTrustGroupsRequireAccessServiceTokenRuleServiceToken]
type zeroTrustGroupsRequireAccessServiceTokenRuleServiceTokenJSON struct {
	TokenID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessServiceTokenRuleServiceToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessServiceTokenRuleServiceTokenJSON) RawJSON() string {
	return r.raw
}

// Matches any valid Access Service Token
type ZeroTrustGroupsRequireAccessAnyValidServiceTokenRule struct {
	// An empty object which matches on all service tokens.
	AnyValidServiceToken interface{}                                              `json:"any_valid_service_token,required"`
	JSON                 zeroTrustGroupsRequireAccessAnyValidServiceTokenRuleJSON `json:"-"`
}

// zeroTrustGroupsRequireAccessAnyValidServiceTokenRuleJSON contains the JSON
// metadata for the struct [ZeroTrustGroupsRequireAccessAnyValidServiceTokenRule]
type zeroTrustGroupsRequireAccessAnyValidServiceTokenRuleJSON struct {
	AnyValidServiceToken apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessAnyValidServiceTokenRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessAnyValidServiceTokenRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsRequireAccessAnyValidServiceTokenRule) implementsZeroTrustZeroTrustGroupsRequire() {
}

// Create Allow or Block policies which evaluate the user based on custom criteria.
type ZeroTrustGroupsRequireAccessExternalEvaluationRule struct {
	ExternalEvaluation ZeroTrustGroupsRequireAccessExternalEvaluationRuleExternalEvaluation `json:"external_evaluation,required"`
	JSON               zeroTrustGroupsRequireAccessExternalEvaluationRuleJSON               `json:"-"`
}

// zeroTrustGroupsRequireAccessExternalEvaluationRuleJSON contains the JSON
// metadata for the struct [ZeroTrustGroupsRequireAccessExternalEvaluationRule]
type zeroTrustGroupsRequireAccessExternalEvaluationRuleJSON struct {
	ExternalEvaluation apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessExternalEvaluationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessExternalEvaluationRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsRequireAccessExternalEvaluationRule) implementsZeroTrustZeroTrustGroupsRequire() {
}

type ZeroTrustGroupsRequireAccessExternalEvaluationRuleExternalEvaluation struct {
	// The API endpoint containing your business logic.
	EvaluateURL string `json:"evaluate_url,required"`
	// The API endpoint containing the key that Access uses to verify that the response
	// came from your API.
	KeysURL string                                                                   `json:"keys_url,required"`
	JSON    zeroTrustGroupsRequireAccessExternalEvaluationRuleExternalEvaluationJSON `json:"-"`
}

// zeroTrustGroupsRequireAccessExternalEvaluationRuleExternalEvaluationJSON
// contains the JSON metadata for the struct
// [ZeroTrustGroupsRequireAccessExternalEvaluationRuleExternalEvaluation]
type zeroTrustGroupsRequireAccessExternalEvaluationRuleExternalEvaluationJSON struct {
	EvaluateURL apijson.Field
	KeysURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessExternalEvaluationRuleExternalEvaluation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessExternalEvaluationRuleExternalEvaluationJSON) RawJSON() string {
	return r.raw
}

// Matches a specific country
type ZeroTrustGroupsRequireAccessCountryRule struct {
	Geo  ZeroTrustGroupsRequireAccessCountryRuleGeo  `json:"geo,required"`
	JSON zeroTrustGroupsRequireAccessCountryRuleJSON `json:"-"`
}

// zeroTrustGroupsRequireAccessCountryRuleJSON contains the JSON metadata for the
// struct [ZeroTrustGroupsRequireAccessCountryRule]
type zeroTrustGroupsRequireAccessCountryRuleJSON struct {
	Geo         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessCountryRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessCountryRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsRequireAccessCountryRule) implementsZeroTrustZeroTrustGroupsRequire() {}

type ZeroTrustGroupsRequireAccessCountryRuleGeo struct {
	// The country code that should be matched.
	CountryCode string                                         `json:"country_code,required"`
	JSON        zeroTrustGroupsRequireAccessCountryRuleGeoJSON `json:"-"`
}

// zeroTrustGroupsRequireAccessCountryRuleGeoJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsRequireAccessCountryRuleGeo]
type zeroTrustGroupsRequireAccessCountryRuleGeoJSON struct {
	CountryCode apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessCountryRuleGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessCountryRuleGeoJSON) RawJSON() string {
	return r.raw
}

// Enforce different MFA options
type ZeroTrustGroupsRequireAccessAuthenticationMethodRule struct {
	AuthMethod ZeroTrustGroupsRequireAccessAuthenticationMethodRuleAuthMethod `json:"auth_method,required"`
	JSON       zeroTrustGroupsRequireAccessAuthenticationMethodRuleJSON       `json:"-"`
}

// zeroTrustGroupsRequireAccessAuthenticationMethodRuleJSON contains the JSON
// metadata for the struct [ZeroTrustGroupsRequireAccessAuthenticationMethodRule]
type zeroTrustGroupsRequireAccessAuthenticationMethodRuleJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessAuthenticationMethodRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessAuthenticationMethodRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsRequireAccessAuthenticationMethodRule) implementsZeroTrustZeroTrustGroupsRequire() {
}

type ZeroTrustGroupsRequireAccessAuthenticationMethodRuleAuthMethod struct {
	// The type of authentication method https://datatracker.ietf.org/doc/html/rfc8176.
	AuthMethod string                                                             `json:"auth_method,required"`
	JSON       zeroTrustGroupsRequireAccessAuthenticationMethodRuleAuthMethodJSON `json:"-"`
}

// zeroTrustGroupsRequireAccessAuthenticationMethodRuleAuthMethodJSON contains the
// JSON metadata for the struct
// [ZeroTrustGroupsRequireAccessAuthenticationMethodRuleAuthMethod]
type zeroTrustGroupsRequireAccessAuthenticationMethodRuleAuthMethodJSON struct {
	AuthMethod  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessAuthenticationMethodRuleAuthMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessAuthenticationMethodRuleAuthMethodJSON) RawJSON() string {
	return r.raw
}

// Enforces a device posture rule has run successfully
type ZeroTrustGroupsRequireAccessDevicePostureRule struct {
	DevicePosture ZeroTrustGroupsRequireAccessDevicePostureRuleDevicePosture `json:"device_posture,required"`
	JSON          zeroTrustGroupsRequireAccessDevicePostureRuleJSON          `json:"-"`
}

// zeroTrustGroupsRequireAccessDevicePostureRuleJSON contains the JSON metadata for
// the struct [ZeroTrustGroupsRequireAccessDevicePostureRule]
type zeroTrustGroupsRequireAccessDevicePostureRuleJSON struct {
	DevicePosture apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessDevicePostureRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessDevicePostureRuleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustGroupsRequireAccessDevicePostureRule) implementsZeroTrustZeroTrustGroupsRequire() {}

type ZeroTrustGroupsRequireAccessDevicePostureRuleDevicePosture struct {
	// The ID of a device posture integration.
	IntegrationUid string                                                         `json:"integration_uid,required"`
	JSON           zeroTrustGroupsRequireAccessDevicePostureRuleDevicePostureJSON `json:"-"`
}

// zeroTrustGroupsRequireAccessDevicePostureRuleDevicePostureJSON contains the JSON
// metadata for the struct
// [ZeroTrustGroupsRequireAccessDevicePostureRuleDevicePosture]
type zeroTrustGroupsRequireAccessDevicePostureRuleDevicePostureJSON struct {
	IntegrationUid apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustGroupsRequireAccessDevicePostureRuleDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGroupsRequireAccessDevicePostureRuleDevicePostureJSON) RawJSON() string {
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
	Group param.Field[shared.UnnamedSchemaRef131Param] `json:"group,required"`
}

func (r AccessGroupNewParamsIncludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsIncludeAccessAccessGroupRule) implementsZeroTrustAccessGroupNewParamsInclude() {
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
	Group param.Field[shared.UnnamedSchemaRef131Param] `json:"group,required"`
}

func (r AccessGroupNewParamsExcludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsExcludeAccessAccessGroupRule) implementsZeroTrustAccessGroupNewParamsExclude() {
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
	Group param.Field[shared.UnnamedSchemaRef131Param] `json:"group,required"`
}

func (r AccessGroupNewParamsRequireAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupNewParamsRequireAccessAccessGroupRule) implementsZeroTrustAccessGroupNewParamsRequire() {
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ZeroTrustGroups       `json:"result,required"`
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
	Group param.Field[shared.UnnamedSchemaRef131Param] `json:"group,required"`
}

func (r AccessGroupUpdateParamsIncludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsIncludeAccessAccessGroupRule) implementsZeroTrustAccessGroupUpdateParamsInclude() {
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
	Group param.Field[shared.UnnamedSchemaRef131Param] `json:"group,required"`
}

func (r AccessGroupUpdateParamsExcludeAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsExcludeAccessAccessGroupRule) implementsZeroTrustAccessGroupUpdateParamsExclude() {
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
	Group param.Field[shared.UnnamedSchemaRef131Param] `json:"group,required"`
}

func (r AccessGroupUpdateParamsRequireAccessAccessGroupRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessGroupUpdateParamsRequireAccessAccessGroupRule) implementsZeroTrustAccessGroupUpdateParamsRequire() {
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ZeroTrustGroups       `json:"result,required"`
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

type AccessGroupDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessGroupDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo     `json:"errors,required"`
	Messages []shared.ResponseInfo     `json:"messages,required"`
	Result   AccessGroupDeleteResponse `json:"result,required"`
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ZeroTrustGroups       `json:"result,required"`
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
