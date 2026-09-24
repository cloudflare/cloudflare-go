// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
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

// Returns the list of entitlements (features and their allocations) for a given
// account. Each entitlement describes a product feature the account is permitted
// to use and the allocation value (boolean, count, range, enum, or string) that
// governs its behaviour.
func (r *EntitlementService) List(ctx context.Context, query EntitlementListParams, opts ...option.RequestOption) (res *pagination.SinglePage[EntitlementListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/entitlements", query.AccountID)
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

// Returns the list of entitlements (features and their allocations) for a given
// account. Each entitlement describes a product feature the account is permitted
// to use and the allocation value (boolean, count, range, enum, or string) that
// governs its behaviour.
func (r *EntitlementService) ListAutoPaging(ctx context.Context, query EntitlementListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[EntitlementListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// A single entitlement record for a zone or account.
type EntitlementListResponse struct {
	// Entitlement identifier — equal to the feature key.
	ID string `json:"id" api:"required"`
	// Represents the allocation value for an entitlement. The shape of `value` depends
	// on `type`: `bool` uses a boolean, `max_count` uses an integer, `enum_number`
	// uses an array of numbers, `range` uses an object with `min` and `max` integer
	// fields, and `string` uses a string.
	Allocation EntitlementListResponseAllocation `json:"allocation" api:"required"`
	// ISO 8601 timestamp (microsecond precision, no timezone offset) when the
	// entitlement was created. Format: `YYYY-MM-DDTHH:MM:SS.ffffff`.
	CreatedDate string `json:"created_date" api:"required"`
	// ISO 8601 timestamp when the entitlement was deleted, or empty string if not
	// deleted.
	DeletedDate string `json:"deleted_date" api:"required"`
	// ISO 8601 timestamp (microsecond precision, no timezone offset) when the
	// entitlement was last edited.
	EditedDate string `json:"edited_date" api:"required"`
	// Describes a product feature associated with an entitlement.
	Feature EntitlementListResponseFeature `json:"feature" api:"required"`
	JSON    entitlementListResponseJSON    `json:"-"`
}

// entitlementListResponseJSON contains the JSON metadata for the struct
// [EntitlementListResponse]
type entitlementListResponseJSON struct {
	ID          apijson.Field
	Allocation  apijson.Field
	CreatedDate apijson.Field
	DeletedDate apijson.Field
	EditedDate  apijson.Field
	Feature     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntitlementListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementListResponseJSON) RawJSON() string {
	return r.raw
}

// Represents the allocation value for an entitlement. The shape of `value` depends
// on `type`: `bool` uses a boolean, `max_count` uses an integer, `enum_number`
// uses an array of numbers, `range` uses an object with `min` and `max` integer
// fields, and `string` uses a string.
type EntitlementListResponseAllocation struct {
	// Allocation type discriminator.
	Type EntitlementListResponseAllocationType `json:"type" api:"required"`
	// Contains the allocation value whose concrete type the `type` field determines:
	// bool yields a boolean, max_count yields an integer, enum_number yields an array
	// of numbers, range yields an object with `min` and `max`, and string yields a
	// string.
	Value EntitlementListResponseAllocationValueUnion `json:"value" api:"required"`
	JSON  entitlementListResponseAllocationJSON       `json:"-"`
}

// entitlementListResponseAllocationJSON contains the JSON metadata for the struct
// [EntitlementListResponseAllocation]
type entitlementListResponseAllocationJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntitlementListResponseAllocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementListResponseAllocationJSON) RawJSON() string {
	return r.raw
}

// Allocation type discriminator.
type EntitlementListResponseAllocationType string

const (
	EntitlementListResponseAllocationTypeBool       EntitlementListResponseAllocationType = "bool"
	EntitlementListResponseAllocationTypeMaxCount   EntitlementListResponseAllocationType = "max_count"
	EntitlementListResponseAllocationTypeEnumNumber EntitlementListResponseAllocationType = "enum_number"
	EntitlementListResponseAllocationTypeRange      EntitlementListResponseAllocationType = "range"
	EntitlementListResponseAllocationTypeString     EntitlementListResponseAllocationType = "string"
)

func (r EntitlementListResponseAllocationType) IsKnown() bool {
	switch r {
	case EntitlementListResponseAllocationTypeBool, EntitlementListResponseAllocationTypeMaxCount, EntitlementListResponseAllocationTypeEnumNumber, EntitlementListResponseAllocationTypeRange, EntitlementListResponseAllocationTypeString:
		return true
	}
	return false
}

// Contains the allocation value whose concrete type the `type` field determines:
// bool yields a boolean, max_count yields an integer, enum_number yields an array
// of numbers, range yields an object with `min` and `max`, and string yields a
// string.
//
// Union satisfied by [shared.UnionBool], [shared.UnionInt], [shared.UnionString],
// [EntitlementListResponseAllocationValueArray] or
// [EntitlementListResponseAllocationValueObject].
type EntitlementListResponseAllocationValueUnion interface {
	ImplementsEntitlementListResponseAllocationValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EntitlementListResponseAllocationValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementListResponseAllocationValueArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EntitlementListResponseAllocationValueObject{}),
		},
	)
}

type EntitlementListResponseAllocationValueArray []float64

func (r EntitlementListResponseAllocationValueArray) ImplementsEntitlementListResponseAllocationValueUnion() {
}

type EntitlementListResponseAllocationValueObject struct {
	Max  int64                                            `json:"max" api:"required"`
	Min  int64                                            `json:"min" api:"required"`
	JSON entitlementListResponseAllocationValueObjectJSON `json:"-"`
}

// entitlementListResponseAllocationValueObjectJSON contains the JSON metadata for
// the struct [EntitlementListResponseAllocationValueObject]
type entitlementListResponseAllocationValueObjectJSON struct {
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntitlementListResponseAllocationValueObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementListResponseAllocationValueObjectJSON) RawJSON() string {
	return r.raw
}

func (r EntitlementListResponseAllocationValueObject) ImplementsEntitlementListResponseAllocationValueUnion() {
}

// Describes a product feature associated with an entitlement.
type EntitlementListResponseFeature struct {
	// Numeric identifier of the feature.
	ID int64 `json:"id" api:"required"`
	// The logical grouping (set) this feature belongs to.
	FeatureSet string `json:"feature_set" api:"required"`
	// Unique string key for the feature.
	Key string `json:"key" api:"required"`
	// Human-readable name of the feature.
	Name string                             `json:"name" api:"required"`
	JSON entitlementListResponseFeatureJSON `json:"-"`
}

// entitlementListResponseFeatureJSON contains the JSON metadata for the struct
// [EntitlementListResponseFeature]
type entitlementListResponseFeatureJSON struct {
	ID          apijson.Field
	FeatureSet  apijson.Field
	Key         apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntitlementListResponseFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entitlementListResponseFeatureJSON) RawJSON() string {
	return r.raw
}

type EntitlementListParams struct {
	// Identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}
