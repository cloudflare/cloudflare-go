// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tenants

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/shared"
)

// TenantService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTenantService] method instead.
type TenantService struct {
	Options      []option.RequestOption
	AccountTypes *AccountTypeService
	Accounts     *AccountService
	Entitlements *EntitlementService
	Memberships  *MembershipService
}

// NewTenantService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTenantService(opts ...option.RequestOption) (r *TenantService) {
	r = &TenantService{}
	r.Options = opts
	r.AccountTypes = NewAccountTypeService(opts...)
	r.Accounts = NewAccountService(opts...)
	r.Entitlements = NewEntitlementService(opts...)
	r.Memberships = NewMembershipService(opts...)
	return
}

// Retrieves a Tenant by Tenant ID.
func (r *TenantService) Get(ctx context.Context, tenantID string, opts ...option.RequestOption) (res *Tenant, err error) {
	var env TenantGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if tenantID == "" {
		err = errors.New("missing required tenant_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tenants/%s", tenantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type Tenant struct {
	Cdate          time.Time            `json:"cdate" api:"required" format:"date-time"`
	Edate          time.Time            `json:"edate" api:"required" format:"date-time"`
	TenantContacts TenantTenantContacts `json:"tenant_contacts" api:"required"`
	TenantLabels   []string             `json:"tenant_labels" api:"required"`
	TenantMetadata TenantTenantMetadata `json:"tenant_metadata" api:"required"`
	TenantName     string               `json:"tenant_name" api:"required"`
	TenantNetwork  interface{}          `json:"tenant_network" api:"required"`
	TenantStatus   string               `json:"tenant_status" api:"required"`
	TenantTag      string               `json:"tenant_tag" api:"required"`
	TenantType     string               `json:"tenant_type" api:"required"`
	TenantUnits    []TenantTenantUnit   `json:"tenant_units" api:"required"`
	CustomerID     string               `json:"customer_id"`
	JSON           tenantJSON           `json:"-"`
}

// tenantJSON contains the JSON metadata for the struct [Tenant]
type tenantJSON struct {
	Cdate          apijson.Field
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
	CustomerID     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Tenant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantJSON) RawJSON() string {
	return r.raw
}

type TenantTenantContacts struct {
	Email   string                   `json:"email"`
	Website string                   `json:"website"`
	JSON    tenantTenantContactsJSON `json:"-"`
}

// tenantTenantContactsJSON contains the JSON metadata for the struct
// [TenantTenantContacts]
type tenantTenantContactsJSON struct {
	Email       apijson.Field
	Website     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantTenantContacts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantTenantContactsJSON) RawJSON() string {
	return r.raw
}

type TenantTenantMetadata struct {
	DNS  TenantTenantMetadataDNS  `json:"dns"`
	JSON tenantTenantMetadataJSON `json:"-"`
}

// tenantTenantMetadataJSON contains the JSON metadata for the struct
// [TenantTenantMetadata]
type tenantTenantMetadataJSON struct {
	DNS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantTenantMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantTenantMetadataJSON) RawJSON() string {
	return r.raw
}

type TenantTenantMetadataDNS struct {
	NSPool TenantTenantMetadataDNSNSPool `json:"ns_pool" api:"required"`
	JSON   tenantTenantMetadataDNSJSON   `json:"-"`
}

// tenantTenantMetadataDNSJSON contains the JSON metadata for the struct
// [TenantTenantMetadataDNS]
type tenantTenantMetadataDNSJSON struct {
	NSPool      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantTenantMetadataDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantTenantMetadataDNSJSON) RawJSON() string {
	return r.raw
}

type TenantTenantMetadataDNSNSPool struct {
	Primary   string                            `json:"primary"`
	Secondary string                            `json:"secondary"`
	JSON      tenantTenantMetadataDnsnsPoolJSON `json:"-"`
}

// tenantTenantMetadataDnsnsPoolJSON contains the JSON metadata for the struct
// [TenantTenantMetadataDNSNSPool]
type tenantTenantMetadataDnsnsPoolJSON struct {
	Primary     apijson.Field
	Secondary   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantTenantMetadataDNSNSPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantTenantMetadataDnsnsPoolJSON) RawJSON() string {
	return r.raw
}

type TenantTenantUnit struct {
	UnitMemberships []interface{}        `json:"unit_memberships" api:"required"`
	UnitMetadata    interface{}          `json:"unit_metadata" api:"required"`
	UnitName        string               `json:"unit_name" api:"required"`
	UnitStatus      string               `json:"unit_status" api:"required"`
	UnitTag         string               `json:"unit_tag" api:"required"`
	JSON            tenantTenantUnitJSON `json:"-"`
}

// tenantTenantUnitJSON contains the JSON metadata for the struct
// [TenantTenantUnit]
type tenantTenantUnitJSON struct {
	UnitMemberships apijson.Field
	UnitMetadata    apijson.Field
	UnitName        apijson.Field
	UnitStatus      apijson.Field
	UnitTag         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TenantTenantUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantTenantUnitJSON) RawJSON() string {
	return r.raw
}

type TenantGetResponseEnvelope struct {
	Errors   []interface{}                    `json:"errors" api:"required"`
	Messages []shared.ResponseInfo            `json:"messages" api:"required"`
	Result   Tenant                           `json:"result" api:"required"`
	Success  TenantGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON     tenantGetResponseEnvelopeJSON    `json:"-"`
}

// tenantGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [TenantGetResponseEnvelope]
type tenantGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TenantGetResponseEnvelopeSuccess bool

const (
	TenantGetResponseEnvelopeSuccessTrue TenantGetResponseEnvelopeSuccess = true
)

func (r TenantGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TenantGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
