// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// GatewayCategoryService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewGatewayCategoryService] method
// instead.
type GatewayCategoryService struct {
	Options []option.RequestOption
}

// NewGatewayCategoryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGatewayCategoryService(opts ...option.RequestOption) (r *GatewayCategoryService) {
	r = &GatewayCategoryService{}
	r.Options = opts
	return
}

// Fetches a list of all categories.
func (r *GatewayCategoryService) List(ctx context.Context, query GatewayCategoryListParams, opts ...option.RequestOption) (res *shared.SinglePage[ZeroTrustGatewayCategories], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/gateway/categories", query.AccountID)
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

// Fetches a list of all categories.
func (r *GatewayCategoryService) ListAutoPaging(ctx context.Context, query GatewayCategoryListParams, opts ...option.RequestOption) *shared.SinglePageAutoPager[ZeroTrustGatewayCategories] {
	return shared.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

type ZeroTrustGatewayCategories struct {
	// The identifier for this category. There is only one category per ID.
	ID int64 `json:"id"`
	// True if the category is in beta and subject to change.
	Beta bool `json:"beta"`
	// Which account types are allowed to create policies based on this category.
	// `blocked` categories are blocked unconditionally for all accounts.
	// `removalPending` categories can be removed from policies but not added.
	// `noBlock` categories cannot be blocked.
	Class ZeroTrustGatewayCategoriesClass `json:"class"`
	// A short summary of domains in the category.
	Description string `json:"description"`
	// The name of the category.
	Name string `json:"name"`
	// All subcategories for this category.
	Subcategories []ZeroTrustGatewayCategoriesSubcategory `json:"subcategories"`
	JSON          zeroTrustGatewayCategoriesJSON          `json:"-"`
}

// zeroTrustGatewayCategoriesJSON contains the JSON metadata for the struct
// [ZeroTrustGatewayCategories]
type zeroTrustGatewayCategoriesJSON struct {
	ID            apijson.Field
	Beta          apijson.Field
	Class         apijson.Field
	Description   apijson.Field
	Name          apijson.Field
	Subcategories apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustGatewayCategories) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayCategoriesJSON) RawJSON() string {
	return r.raw
}

// Which account types are allowed to create policies based on this category.
// `blocked` categories are blocked unconditionally for all accounts.
// `removalPending` categories can be removed from policies but not added.
// `noBlock` categories cannot be blocked.
type ZeroTrustGatewayCategoriesClass string

const (
	ZeroTrustGatewayCategoriesClassFree           ZeroTrustGatewayCategoriesClass = "free"
	ZeroTrustGatewayCategoriesClassPremium        ZeroTrustGatewayCategoriesClass = "premium"
	ZeroTrustGatewayCategoriesClassBlocked        ZeroTrustGatewayCategoriesClass = "blocked"
	ZeroTrustGatewayCategoriesClassRemovalPending ZeroTrustGatewayCategoriesClass = "removalPending"
	ZeroTrustGatewayCategoriesClassNoBlock        ZeroTrustGatewayCategoriesClass = "noBlock"
)

func (r ZeroTrustGatewayCategoriesClass) IsKnown() bool {
	switch r {
	case ZeroTrustGatewayCategoriesClassFree, ZeroTrustGatewayCategoriesClassPremium, ZeroTrustGatewayCategoriesClassBlocked, ZeroTrustGatewayCategoriesClassRemovalPending, ZeroTrustGatewayCategoriesClassNoBlock:
		return true
	}
	return false
}

type ZeroTrustGatewayCategoriesSubcategory struct {
	// The identifier for this category. There is only one category per ID.
	ID int64 `json:"id"`
	// True if the category is in beta and subject to change.
	Beta bool `json:"beta"`
	// Which account types are allowed to create policies based on this category.
	// `blocked` categories are blocked unconditionally for all accounts.
	// `removalPending` categories can be removed from policies but not added.
	// `noBlock` categories cannot be blocked.
	Class ZeroTrustGatewayCategoriesSubcategoriesClass `json:"class"`
	// A short summary of domains in the category.
	Description string `json:"description"`
	// The name of the category.
	Name string                                    `json:"name"`
	JSON zeroTrustGatewayCategoriesSubcategoryJSON `json:"-"`
}

// zeroTrustGatewayCategoriesSubcategoryJSON contains the JSON metadata for the
// struct [ZeroTrustGatewayCategoriesSubcategory]
type zeroTrustGatewayCategoriesSubcategoryJSON struct {
	ID          apijson.Field
	Beta        apijson.Field
	Class       apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayCategoriesSubcategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayCategoriesSubcategoryJSON) RawJSON() string {
	return r.raw
}

// Which account types are allowed to create policies based on this category.
// `blocked` categories are blocked unconditionally for all accounts.
// `removalPending` categories can be removed from policies but not added.
// `noBlock` categories cannot be blocked.
type ZeroTrustGatewayCategoriesSubcategoriesClass string

const (
	ZeroTrustGatewayCategoriesSubcategoriesClassFree           ZeroTrustGatewayCategoriesSubcategoriesClass = "free"
	ZeroTrustGatewayCategoriesSubcategoriesClassPremium        ZeroTrustGatewayCategoriesSubcategoriesClass = "premium"
	ZeroTrustGatewayCategoriesSubcategoriesClassBlocked        ZeroTrustGatewayCategoriesSubcategoriesClass = "blocked"
	ZeroTrustGatewayCategoriesSubcategoriesClassRemovalPending ZeroTrustGatewayCategoriesSubcategoriesClass = "removalPending"
	ZeroTrustGatewayCategoriesSubcategoriesClassNoBlock        ZeroTrustGatewayCategoriesSubcategoriesClass = "noBlock"
)

func (r ZeroTrustGatewayCategoriesSubcategoriesClass) IsKnown() bool {
	switch r {
	case ZeroTrustGatewayCategoriesSubcategoriesClassFree, ZeroTrustGatewayCategoriesSubcategoriesClassPremium, ZeroTrustGatewayCategoriesSubcategoriesClassBlocked, ZeroTrustGatewayCategoriesSubcategoriesClassRemovalPending, ZeroTrustGatewayCategoriesSubcategoriesClassNoBlock:
		return true
	}
	return false
}

type GatewayCategoryListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
