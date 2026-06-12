// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package tenants

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/shared"
	"github.com/tidwall/gjson"
)

// EntitlementService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntitlementService] method instead.
type EntitlementService struct {
	Options []option.RequestOption
}

// NewEntitlementService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEntitlementService(opts ...option.RequestOption) (r *EntitlementService) {
	r = &EntitlementService{}
	r.Options = opts
	return
}

// List of innate entitlements available for the Tenant.
func (r *EntitlementService) Get(ctx context.Context, tenantID string, opts ...option.RequestOption) (res *TenantEntitlements, err error) {
	var env EntitlementGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if tenantID == "" {
		err = errors.New("missing required tenant_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("tenants/%s/entitlements", tenantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type TenantEntitlements struct {
	AllowAddSubdomain      TenantEntitlementsAllowAddSubdomain      `json:"allow_add_subdomain" api:"required"`
	AllowAutoAcceptInvites TenantEntitlementsAllowAutoAcceptInvites `json:"allow_auto_accept_invites" api:"required"`
	CNAMESetupAllowed      TenantEntitlementsCNAMESetupAllowed      `json:"cname_setup_allowed" api:"required"`
	CustomEntitlements     []TenantEntitlementsCustomEntitlement    `json:"custom_entitlements" api:"required,nullable"`
	MhsCertificateCount    TenantEntitlementsMhsCertificateCount    `json:"mhs_certificate_count" api:"required"`
	PartialSetupAllowed    TenantEntitlementsPartialSetupAllowed    `json:"partial_setup_allowed" api:"required"`
	JSON                   tenantEntitlementsJSON                   `json:"-"`
}

// tenantEntitlementsJSON contains the JSON metadata for the struct
// [TenantEntitlements]
type tenantEntitlementsJSON struct {
	AllowAddSubdomain      apijson.Field
	AllowAutoAcceptInvites apijson.Field
	CNAMESetupAllowed      apijson.Field
	CustomEntitlements     apijson.Field
	MhsCertificateCount    apijson.Field
	PartialSetupAllowed    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *TenantEntitlements) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantEntitlementsJSON) RawJSON() string {
	return r.raw
}

type TenantEntitlementsAllowAddSubdomain struct {
	Type  TenantEntitlementsAllowAddSubdomainType `json:"type" api:"required"`
	Value bool                                    `json:"value" api:"required"`
	JSON  tenantEntitlementsAllowAddSubdomainJSON `json:"-"`
}

// tenantEntitlementsAllowAddSubdomainJSON contains the JSON metadata for the
// struct [TenantEntitlementsAllowAddSubdomain]
type tenantEntitlementsAllowAddSubdomainJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantEntitlementsAllowAddSubdomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantEntitlementsAllowAddSubdomainJSON) RawJSON() string {
	return r.raw
}

type TenantEntitlementsAllowAddSubdomainType string

const (
	TenantEntitlementsAllowAddSubdomainTypeBool TenantEntitlementsAllowAddSubdomainType = "bool"
)

func (r TenantEntitlementsAllowAddSubdomainType) IsKnown() bool {
	switch r {
	case TenantEntitlementsAllowAddSubdomainTypeBool:
		return true
	}
	return false
}

type TenantEntitlementsAllowAutoAcceptInvites struct {
	Type  TenantEntitlementsAllowAutoAcceptInvitesType `json:"type" api:"required"`
	Value bool                                         `json:"value" api:"required"`
	JSON  tenantEntitlementsAllowAutoAcceptInvitesJSON `json:"-"`
}

// tenantEntitlementsAllowAutoAcceptInvitesJSON contains the JSON metadata for the
// struct [TenantEntitlementsAllowAutoAcceptInvites]
type tenantEntitlementsAllowAutoAcceptInvitesJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantEntitlementsAllowAutoAcceptInvites) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantEntitlementsAllowAutoAcceptInvitesJSON) RawJSON() string {
	return r.raw
}

type TenantEntitlementsAllowAutoAcceptInvitesType string

const (
	TenantEntitlementsAllowAutoAcceptInvitesTypeBool TenantEntitlementsAllowAutoAcceptInvitesType = "bool"
)

func (r TenantEntitlementsAllowAutoAcceptInvitesType) IsKnown() bool {
	switch r {
	case TenantEntitlementsAllowAutoAcceptInvitesTypeBool:
		return true
	}
	return false
}

type TenantEntitlementsCNAMESetupAllowed struct {
	Type  TenantEntitlementsCNAMESetupAllowedType `json:"type" api:"required"`
	Value bool                                    `json:"value" api:"required"`
	JSON  tenantEntitlementsCNAMESetupAllowedJSON `json:"-"`
}

// tenantEntitlementsCNAMESetupAllowedJSON contains the JSON metadata for the
// struct [TenantEntitlementsCNAMESetupAllowed]
type tenantEntitlementsCNAMESetupAllowedJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantEntitlementsCNAMESetupAllowed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantEntitlementsCNAMESetupAllowedJSON) RawJSON() string {
	return r.raw
}

type TenantEntitlementsCNAMESetupAllowedType string

const (
	TenantEntitlementsCNAMESetupAllowedTypeBool TenantEntitlementsCNAMESetupAllowedType = "bool"
)

func (r TenantEntitlementsCNAMESetupAllowedType) IsKnown() bool {
	switch r {
	case TenantEntitlementsCNAMESetupAllowedTypeBool:
		return true
	}
	return false
}

type TenantEntitlementsCustomEntitlement struct {
	Allocation TenantEntitlementsCustomEntitlementsAllocation `json:"allocation" api:"required"`
	Feature    TenantEntitlementsCustomEntitlementsFeature    `json:"feature" api:"required"`
	JSON       tenantEntitlementsCustomEntitlementJSON        `json:"-"`
}

// tenantEntitlementsCustomEntitlementJSON contains the JSON metadata for the
// struct [TenantEntitlementsCustomEntitlement]
type tenantEntitlementsCustomEntitlementJSON struct {
	Allocation  apijson.Field
	Feature     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantEntitlementsCustomEntitlement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantEntitlementsCustomEntitlementJSON) RawJSON() string {
	return r.raw
}

type TenantEntitlementsCustomEntitlementsAllocation struct {
	Type TenantEntitlementsCustomEntitlementsAllocationType `json:"type" api:"required"`
	// This field can have the runtime type of [int64], [bool], [interface{}].
	Value interface{}                                        `json:"value"`
	JSON  tenantEntitlementsCustomEntitlementsAllocationJSON `json:"-"`
	union TenantEntitlementsCustomEntitlementsAllocationUnion
}

// tenantEntitlementsCustomEntitlementsAllocationJSON contains the JSON metadata
// for the struct [TenantEntitlementsCustomEntitlementsAllocation]
type tenantEntitlementsCustomEntitlementsAllocationJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r tenantEntitlementsCustomEntitlementsAllocationJSON) RawJSON() string {
	return r.raw
}

func (r *TenantEntitlementsCustomEntitlementsAllocation) UnmarshalJSON(data []byte) (err error) {
	*r = TenantEntitlementsCustomEntitlementsAllocation{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TenantEntitlementsCustomEntitlementsAllocationUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIMaxCountAllocation],
// [TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIBoolAllocation],
// [TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPINullAllocation].
func (r TenantEntitlementsCustomEntitlementsAllocation) AsUnion() TenantEntitlementsCustomEntitlementsAllocationUnion {
	return r.union
}

// Union satisfied by
// [TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIMaxCountAllocation],
// [TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIBoolAllocation]
// or
// [TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPINullAllocation].
type TenantEntitlementsCustomEntitlementsAllocationUnion interface {
	implementsTenantEntitlementsCustomEntitlementsAllocation()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TenantEntitlementsCustomEntitlementsAllocationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIMaxCountAllocation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIBoolAllocation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPINullAllocation{}),
		},
	)
}

type TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIMaxCountAllocation struct {
	Type  TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIMaxCountAllocationType `json:"type" api:"required"`
	Value int64                                                                                `json:"value" api:"required"`
	JSON  tenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIMaxCountAllocationJSON `json:"-"`
}

// tenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIMaxCountAllocationJSON
// contains the JSON metadata for the struct
// [TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIMaxCountAllocation]
type tenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIMaxCountAllocationJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIMaxCountAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIMaxCountAllocationJSON) RawJSON() string {
	return r.raw
}

func (r TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIMaxCountAllocation) implementsTenantEntitlementsCustomEntitlementsAllocation() {
}

type TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIMaxCountAllocationType string

const (
	TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIMaxCountAllocationTypeMaxCount TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIMaxCountAllocationType = "max_count"
)

func (r TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIMaxCountAllocationType) IsKnown() bool {
	switch r {
	case TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIMaxCountAllocationTypeMaxCount:
		return true
	}
	return false
}

type TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIBoolAllocation struct {
	Type  TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIBoolAllocationType `json:"type" api:"required"`
	Value bool                                                                             `json:"value" api:"required"`
	JSON  tenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIBoolAllocationJSON `json:"-"`
}

// tenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIBoolAllocationJSON
// contains the JSON metadata for the struct
// [TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIBoolAllocation]
type tenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIBoolAllocationJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIBoolAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIBoolAllocationJSON) RawJSON() string {
	return r.raw
}

func (r TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIBoolAllocation) implementsTenantEntitlementsCustomEntitlementsAllocation() {
}

type TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIBoolAllocationType string

const (
	TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIBoolAllocationTypeBool TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIBoolAllocationType = "bool"
)

func (r TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIBoolAllocationType) IsKnown() bool {
	switch r {
	case TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPIBoolAllocationTypeBool:
		return true
	}
	return false
}

type TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPINullAllocation struct {
	Type  TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPINullAllocationType `json:"type" api:"required"`
	Value interface{}                                                                      `json:"value"`
	JSON  tenantEntitlementsCustomEntitlementsAllocationOrganizationsAPINullAllocationJSON `json:"-"`
}

// tenantEntitlementsCustomEntitlementsAllocationOrganizationsAPINullAllocationJSON
// contains the JSON metadata for the struct
// [TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPINullAllocation]
type tenantEntitlementsCustomEntitlementsAllocationOrganizationsAPINullAllocationJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPINullAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantEntitlementsCustomEntitlementsAllocationOrganizationsAPINullAllocationJSON) RawJSON() string {
	return r.raw
}

func (r TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPINullAllocation) implementsTenantEntitlementsCustomEntitlementsAllocation() {
}

type TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPINullAllocationType string

const (
	TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPINullAllocationTypeEmpty TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPINullAllocationType = ""
)

func (r TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPINullAllocationType) IsKnown() bool {
	switch r {
	case TenantEntitlementsCustomEntitlementsAllocationOrganizationsAPINullAllocationTypeEmpty:
		return true
	}
	return false
}

type TenantEntitlementsCustomEntitlementsAllocationType string

const (
	TenantEntitlementsCustomEntitlementsAllocationTypeMaxCount TenantEntitlementsCustomEntitlementsAllocationType = "max_count"
	TenantEntitlementsCustomEntitlementsAllocationTypeBool     TenantEntitlementsCustomEntitlementsAllocationType = "bool"
	TenantEntitlementsCustomEntitlementsAllocationTypeEmpty    TenantEntitlementsCustomEntitlementsAllocationType = ""
)

func (r TenantEntitlementsCustomEntitlementsAllocationType) IsKnown() bool {
	switch r {
	case TenantEntitlementsCustomEntitlementsAllocationTypeMaxCount, TenantEntitlementsCustomEntitlementsAllocationTypeBool, TenantEntitlementsCustomEntitlementsAllocationTypeEmpty:
		return true
	}
	return false
}

type TenantEntitlementsCustomEntitlementsFeature struct {
	Key  string                                          `json:"key" api:"required"`
	JSON tenantEntitlementsCustomEntitlementsFeatureJSON `json:"-"`
}

// tenantEntitlementsCustomEntitlementsFeatureJSON contains the JSON metadata for
// the struct [TenantEntitlementsCustomEntitlementsFeature]
type tenantEntitlementsCustomEntitlementsFeatureJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantEntitlementsCustomEntitlementsFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantEntitlementsCustomEntitlementsFeatureJSON) RawJSON() string {
	return r.raw
}

type TenantEntitlementsMhsCertificateCount struct {
	Type  TenantEntitlementsMhsCertificateCountType `json:"type" api:"required"`
	Value int64                                     `json:"value" api:"required"`
	JSON  tenantEntitlementsMhsCertificateCountJSON `json:"-"`
}

// tenantEntitlementsMhsCertificateCountJSON contains the JSON metadata for the
// struct [TenantEntitlementsMhsCertificateCount]
type tenantEntitlementsMhsCertificateCountJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantEntitlementsMhsCertificateCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantEntitlementsMhsCertificateCountJSON) RawJSON() string {
	return r.raw
}

type TenantEntitlementsMhsCertificateCountType string

const (
	TenantEntitlementsMhsCertificateCountTypeMaxCount TenantEntitlementsMhsCertificateCountType = "max_count"
)

func (r TenantEntitlementsMhsCertificateCountType) IsKnown() bool {
	switch r {
	case TenantEntitlementsMhsCertificateCountTypeMaxCount:
		return true
	}
	return false
}

type TenantEntitlementsPartialSetupAllowed struct {
	Type  TenantEntitlementsPartialSetupAllowedType `json:"type" api:"required"`
	Value bool                                      `json:"value" api:"required"`
	JSON  tenantEntitlementsPartialSetupAllowedJSON `json:"-"`
}

// tenantEntitlementsPartialSetupAllowedJSON contains the JSON metadata for the
// struct [TenantEntitlementsPartialSetupAllowed]
type tenantEntitlementsPartialSetupAllowedJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TenantEntitlementsPartialSetupAllowed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tenantEntitlementsPartialSetupAllowedJSON) RawJSON() string {
	return r.raw
}

type TenantEntitlementsPartialSetupAllowedType string

const (
	TenantEntitlementsPartialSetupAllowedTypeBool TenantEntitlementsPartialSetupAllowedType = "bool"
)

func (r TenantEntitlementsPartialSetupAllowedType) IsKnown() bool {
	switch r {
	case TenantEntitlementsPartialSetupAllowedTypeBool:
		return true
	}
	return false
}

type EntitlementGetResponseEnvelope struct {
	Errors   []interface{}                         `json:"errors" api:"required"`
	Messages []shared.ResponseInfo                 `json:"messages" api:"required"`
	Result   TenantEntitlements                    `json:"result" api:"required"`
	Success  EntitlementGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON     entitlementGetResponseEnvelopeJSON    `json:"-"`
}

// entitlementGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [EntitlementGetResponseEnvelope]
type entitlementGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntitlementGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EntitlementGetResponseEnvelopeSuccess bool

const (
	EntitlementGetResponseEnvelopeSuccessTrue EntitlementGetResponseEnvelopeSuccess = true
)

func (r EntitlementGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case EntitlementGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
