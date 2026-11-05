// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"
	"net/http"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v7/shared"
	"github.com/tidwall/gjson"
)

// TenantService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTenantService] method instead.
type TenantService struct {
	Options []option.RequestOption
}

// NewTenantService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTenantService(opts ...option.RequestOption) (r *TenantService) {
	r = &TenantService{}
	r.Options = opts
	return
}

// Lists the Cloudflare tenants in which the authenticated user has direct
// membership using the legacy Tenant response shape. Use `GET /organizations` to
// list Organizations the user can access, including through hierarchy, using the
// modern Organization response shape.
func (r *TenantService) List(ctx context.Context, opts ...option.RequestOption) (res *pagination.SinglePage[TenantListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "user/tenants"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Lists the Cloudflare tenants in which the authenticated user has direct
// membership using the legacy Tenant response shape. Use `GET /organizations` to
// list Organizations the user can access, including through hierarchy, using the
// modern Organization response shape.
func (r *TenantService) ListAutoPaging(ctx context.Context, opts ...option.RequestOption) *pagination.SinglePageAutoPager[TenantListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, opts...))
}

type TenantListResponse struct {
	Cdate          time.Time                        `json:"cdate" api:"required" format:"date-time"`
	CustomerID     string                           `json:"customer_id" api:"required,nullable"`
	Edate          time.Time                        `json:"edate" api:"required" format:"date-time"`
	TenantContacts TenantListResponseTenantContacts `json:"tenant_contacts" api:"required"`
	TenantLabels   []string                         `json:"tenant_labels" api:"required,nullable"`
	TenantMetadata TenantListResponseTenantMetadata `json:"tenant_metadata" api:"required"`
	TenantName     string                           `json:"tenant_name" api:"required"`
	TenantNetwork  TenantListResponseTenantNetwork  `json:"tenant_network" api:"required"`
	TenantStatus   string                           `json:"tenant_status" api:"required"`
	TenantTag      string                           `json:"tenant_tag" api:"required"`
	TenantType     string                           `json:"tenant_type" api:"required"`
	TenantUnits    []TenantListResponseTenantUnit   `json:"tenant_units" api:"required,nullable"`
	// Which top-level dashboard experiences to render for this tenant; both may be
	// true.
	UIFlags TenantListResponseUIFlags `json:"ui_flags" api:"required"`
	JSON    tenantListResponseJSON    `json:"-"`
}

// tenantListResponseJSON contains the JSON metadata for the struct
// [TenantListResponse]
type tenantListResponseJSON struct {
	Cdate          apijson.Field
	CustomerID     apijson.Field
	Edate          apijson.Field
	TenantContacts apijson.Field
	TenantLabels   apijson.Field
	TenantMetadata apijson.Field
	TenantName     apijson.Field
	TenantNetwork  apijson.Field
	TenantStatus   apijson.Field
	TenantTag      apijson.Field
	TenantType     apijson.Field
	TenantUnits    apijson.Field
	UIFlags        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TenantListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantListResponseJSON) RawJSON() string {
	return r.raw
}

type TenantListResponseTenantContacts struct {
	Email   string                               `json:"email"`
	Website string                               `json:"website"`
	JSON    tenantListResponseTenantContactsJSON `json:"-"`
}

// tenantListResponseTenantContactsJSON contains the JSON metadata for the struct
// [TenantListResponseTenantContacts]
type tenantListResponseTenantContactsJSON struct {
	Email       apijson.Field
	Website     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantListResponseTenantContacts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantListResponseTenantContactsJSON) RawJSON() string {
	return r.raw
}

type TenantListResponseTenantMetadata struct {
	BillingAccountType        string                                                     `json:"billing_account_type"`
	ConfigurationAccount      string                                                     `json:"configuration_account"`
	CreatedBy                 string                                                     `json:"created_by"`
	DNS                       TenantListResponseTenantMetadataDNS                        `json:"dns"`
	Flags                     map[string]string                                          `json:"flags"`
	InnateAccountEntitlements []TenantListResponseTenantMetadataInnateAccountEntitlement `json:"innate_account_entitlements"`
	TenantEntitlements        []TenantListResponseTenantMetadataTenantEntitlement        `json:"tenant_entitlements"`
	JSON                      tenantListResponseTenantMetadataJSON                       `json:"-"`
}

// tenantListResponseTenantMetadataJSON contains the JSON metadata for the struct
// [TenantListResponseTenantMetadata]
type tenantListResponseTenantMetadataJSON struct {
	BillingAccountType        apijson.Field
	ConfigurationAccount      apijson.Field
	CreatedBy                 apijson.Field
	DNS                       apijson.Field
	Flags                     apijson.Field
	InnateAccountEntitlements apijson.Field
	TenantEntitlements        apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *TenantListResponseTenantMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantListResponseTenantMetadataJSON) RawJSON() string {
	return r.raw
}

type TenantListResponseTenantMetadataDNS struct {
	NSPool TenantListResponseTenantMetadataDNSNSPool `json:"ns_pool" api:"required"`
	JSON   tenantListResponseTenantMetadataDNSJSON   `json:"-"`
}

// tenantListResponseTenantMetadataDNSJSON contains the JSON metadata for the
// struct [TenantListResponseTenantMetadataDNS]
type tenantListResponseTenantMetadataDNSJSON struct {
	NSPool      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantListResponseTenantMetadataDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantListResponseTenantMetadataDNSJSON) RawJSON() string {
	return r.raw
}

type TenantListResponseTenantMetadataDNSNSPool struct {
	Primary   string                                        `json:"primary" api:"required,nullable"`
	Secondary string                                        `json:"secondary" api:"required,nullable"`
	JSON      tenantListResponseTenantMetadataDnsnsPoolJSON `json:"-"`
}

// tenantListResponseTenantMetadataDnsnsPoolJSON contains the JSON metadata for the
// struct [TenantListResponseTenantMetadataDNSNSPool]
type tenantListResponseTenantMetadataDnsnsPoolJSON struct {
	Primary     apijson.Field
	Secondary   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantListResponseTenantMetadataDNSNSPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantListResponseTenantMetadataDnsnsPoolJSON) RawJSON() string {
	return r.raw
}

type TenantListResponseTenantMetadataInnateAccountEntitlement struct {
	Allocation TenantListResponseTenantMetadataInnateAccountEntitlementsAllocation `json:"allocation" api:"required"`
	Feature    TenantListResponseTenantMetadataInnateAccountEntitlementsFeature    `json:"feature" api:"required"`
	JSON       tenantListResponseTenantMetadataInnateAccountEntitlementJSON        `json:"-"`
}

// tenantListResponseTenantMetadataInnateAccountEntitlementJSON contains the JSON
// metadata for the struct
// [TenantListResponseTenantMetadataInnateAccountEntitlement]
type tenantListResponseTenantMetadataInnateAccountEntitlementJSON struct {
	Allocation  apijson.Field
	Feature     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantListResponseTenantMetadataInnateAccountEntitlement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantListResponseTenantMetadataInnateAccountEntitlementJSON) RawJSON() string {
	return r.raw
}

type TenantListResponseTenantMetadataInnateAccountEntitlementsAllocation struct {
	Type  string                                                                        `json:"type" api:"required"`
	Value TenantListResponseTenantMetadataInnateAccountEntitlementsAllocationValueUnion `json:"value" api:"required"`
	JSON  tenantListResponseTenantMetadataInnateAccountEntitlementsAllocationJSON       `json:"-"`
}

// tenantListResponseTenantMetadataInnateAccountEntitlementsAllocationJSON contains
// the JSON metadata for the struct
// [TenantListResponseTenantMetadataInnateAccountEntitlementsAllocation]
type tenantListResponseTenantMetadataInnateAccountEntitlementsAllocationJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantListResponseTenantMetadataInnateAccountEntitlementsAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantListResponseTenantMetadataInnateAccountEntitlementsAllocationJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionBool], [shared.UnionFloat],
// [shared.UnionString],
// [TenantListResponseTenantMetadataInnateAccountEntitlementsAllocationValueMap] or
// [TenantListResponseTenantMetadataInnateAccountEntitlementsAllocationValueArray].
type TenantListResponseTenantMetadataInnateAccountEntitlementsAllocationValueUnion interface {
	ImplementsTenantListResponseTenantMetadataInnateAccountEntitlementsAllocationValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TenantListResponseTenantMetadataInnateAccountEntitlementsAllocationValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TenantListResponseTenantMetadataInnateAccountEntitlementsAllocationValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TenantListResponseTenantMetadataInnateAccountEntitlementsAllocationValueArray{}),
		},
	)
}

type TenantListResponseTenantMetadataInnateAccountEntitlementsAllocationValueMap map[string]interface{}

func (r TenantListResponseTenantMetadataInnateAccountEntitlementsAllocationValueMap) ImplementsTenantListResponseTenantMetadataInnateAccountEntitlementsAllocationValueUnion() {
}

type TenantListResponseTenantMetadataInnateAccountEntitlementsAllocationValueArray []interface{}

func (r TenantListResponseTenantMetadataInnateAccountEntitlementsAllocationValueArray) ImplementsTenantListResponseTenantMetadataInnateAccountEntitlementsAllocationValueUnion() {
}

type TenantListResponseTenantMetadataInnateAccountEntitlementsFeature struct {
	Key  string                                                               `json:"key" api:"required"`
	JSON tenantListResponseTenantMetadataInnateAccountEntitlementsFeatureJSON `json:"-"`
}

// tenantListResponseTenantMetadataInnateAccountEntitlementsFeatureJSON contains
// the JSON metadata for the struct
// [TenantListResponseTenantMetadataInnateAccountEntitlementsFeature]
type tenantListResponseTenantMetadataInnateAccountEntitlementsFeatureJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantListResponseTenantMetadataInnateAccountEntitlementsFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantListResponseTenantMetadataInnateAccountEntitlementsFeatureJSON) RawJSON() string {
	return r.raw
}

type TenantListResponseTenantMetadataTenantEntitlement struct {
	Allocation TenantListResponseTenantMetadataTenantEntitlementsAllocation `json:"allocation" api:"required"`
	Feature    TenantListResponseTenantMetadataTenantEntitlementsFeature    `json:"feature" api:"required"`
	JSON       tenantListResponseTenantMetadataTenantEntitlementJSON        `json:"-"`
}

// tenantListResponseTenantMetadataTenantEntitlementJSON contains the JSON metadata
// for the struct [TenantListResponseTenantMetadataTenantEntitlement]
type tenantListResponseTenantMetadataTenantEntitlementJSON struct {
	Allocation  apijson.Field
	Feature     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantListResponseTenantMetadataTenantEntitlement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantListResponseTenantMetadataTenantEntitlementJSON) RawJSON() string {
	return r.raw
}

type TenantListResponseTenantMetadataTenantEntitlementsAllocation struct {
	Type  string                                                                 `json:"type" api:"required"`
	Value TenantListResponseTenantMetadataTenantEntitlementsAllocationValueUnion `json:"value" api:"required"`
	JSON  tenantListResponseTenantMetadataTenantEntitlementsAllocationJSON       `json:"-"`
}

// tenantListResponseTenantMetadataTenantEntitlementsAllocationJSON contains the
// JSON metadata for the struct
// [TenantListResponseTenantMetadataTenantEntitlementsAllocation]
type tenantListResponseTenantMetadataTenantEntitlementsAllocationJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantListResponseTenantMetadataTenantEntitlementsAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantListResponseTenantMetadataTenantEntitlementsAllocationJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionBool], [shared.UnionFloat],
// [shared.UnionString],
// [TenantListResponseTenantMetadataTenantEntitlementsAllocationValueMap] or
// [TenantListResponseTenantMetadataTenantEntitlementsAllocationValueArray].
type TenantListResponseTenantMetadataTenantEntitlementsAllocationValueUnion interface {
	ImplementsTenantListResponseTenantMetadataTenantEntitlementsAllocationValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TenantListResponseTenantMetadataTenantEntitlementsAllocationValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TenantListResponseTenantMetadataTenantEntitlementsAllocationValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TenantListResponseTenantMetadataTenantEntitlementsAllocationValueArray{}),
		},
	)
}

type TenantListResponseTenantMetadataTenantEntitlementsAllocationValueMap map[string]interface{}

func (r TenantListResponseTenantMetadataTenantEntitlementsAllocationValueMap) ImplementsTenantListResponseTenantMetadataTenantEntitlementsAllocationValueUnion() {
}

type TenantListResponseTenantMetadataTenantEntitlementsAllocationValueArray []interface{}

func (r TenantListResponseTenantMetadataTenantEntitlementsAllocationValueArray) ImplementsTenantListResponseTenantMetadataTenantEntitlementsAllocationValueUnion() {
}

type TenantListResponseTenantMetadataTenantEntitlementsFeature struct {
	Key  string                                                        `json:"key" api:"required"`
	JSON tenantListResponseTenantMetadataTenantEntitlementsFeatureJSON `json:"-"`
}

// tenantListResponseTenantMetadataTenantEntitlementsFeatureJSON contains the JSON
// metadata for the struct
// [TenantListResponseTenantMetadataTenantEntitlementsFeature]
type tenantListResponseTenantMetadataTenantEntitlementsFeatureJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantListResponseTenantMetadataTenantEntitlementsFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantListResponseTenantMetadataTenantEntitlementsFeatureJSON) RawJSON() string {
	return r.raw
}

type TenantListResponseTenantNetwork struct {
	Choices map[string]interface{}              `json:"choices"`
	JSON    tenantListResponseTenantNetworkJSON `json:"-"`
}

// tenantListResponseTenantNetworkJSON contains the JSON metadata for the struct
// [TenantListResponseTenantNetwork]
type tenantListResponseTenantNetworkJSON struct {
	Choices     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantListResponseTenantNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantListResponseTenantNetworkJSON) RawJSON() string {
	return r.raw
}

type TenantListResponseTenantUnit struct {
	UnitMemberships []TenantListResponseTenantUnitsUnitMembership `json:"unit_memberships" api:"required,nullable"`
	UnitMetadata    interface{}                                   `json:"unit_metadata" api:"required"`
	UnitName        string                                        `json:"unit_name" api:"required"`
	UnitStatus      string                                        `json:"unit_status" api:"required"`
	UnitTag         string                                        `json:"unit_tag" api:"required"`
	JSON            tenantListResponseTenantUnitJSON              `json:"-"`
}

// tenantListResponseTenantUnitJSON contains the JSON metadata for the struct
// [TenantListResponseTenantUnit]
type tenantListResponseTenantUnitJSON struct {
	UnitMemberships apijson.Field
	UnitMetadata    apijson.Field
	UnitName        apijson.Field
	UnitStatus      apijson.Field
	UnitTag         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TenantListResponseTenantUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantListResponseTenantUnitJSON) RawJSON() string {
	return r.raw
}

type TenantListResponseTenantUnitsUnitMembership struct {
	MemberMetadata interface{}                                     `json:"member_metadata" api:"required"`
	MemberStatus   string                                          `json:"member_status" api:"required"`
	MemberTag      string                                          `json:"member_tag" api:"required"`
	JSON           tenantListResponseTenantUnitsUnitMembershipJSON `json:"-"`
}

// tenantListResponseTenantUnitsUnitMembershipJSON contains the JSON metadata for
// the struct [TenantListResponseTenantUnitsUnitMembership]
type tenantListResponseTenantUnitsUnitMembershipJSON struct {
	MemberMetadata apijson.Field
	MemberStatus   apijson.Field
	MemberTag      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TenantListResponseTenantUnitsUnitMembership) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantListResponseTenantUnitsUnitMembershipJSON) RawJSON() string {
	return r.raw
}

// Which top-level dashboard experiences to render for this tenant; both may be
// true.
type TenantListResponseUIFlags struct {
	// Render the Organizations experience for this tenant.
	ShowOrgsUI bool `json:"show_orgs_ui" api:"required"`
	// Render the legacy Tenants experience for this tenant.
	ShowTenantsUI bool                          `json:"show_tenants_ui" api:"required"`
	JSON          tenantListResponseUIFlagsJSON `json:"-"`
}

// tenantListResponseUIFlagsJSON contains the JSON metadata for the struct
// [TenantListResponseUIFlags]
type tenantListResponseUIFlagsJSON struct {
	ShowOrgsUI    apijson.Field
	ShowTenantsUI apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TenantListResponseUIFlags) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantListResponseUIFlagsJSON) RawJSON() string {
	return r.raw
}
