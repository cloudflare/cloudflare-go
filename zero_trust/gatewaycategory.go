// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// GatewayCategoryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGatewayCategoryService] method instead.
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

// List all categories.
func (r *GatewayCategoryService) List(ctx context.Context, query GatewayCategoryListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Category], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/categories", query.AccountID)
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

// List all categories.
func (r *GatewayCategoryService) ListAutoPaging(ctx context.Context, query GatewayCategoryListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Category] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

type Category struct {
	// Identify this category. Only one category per ID.
	ID int64 `json:"id"`
	// Indicate whether the category is in beta and subject to change.
	Beta bool `json:"beta"`
	// Specify which account types can create policies for this category. `blocked`
	// Blocks unconditionally for all accounts. `removalPending` Allows removal from
	// policies but disables addition. `noBlock` Prevents blocking.
	Class CategoryClass `json:"class"`
	// Provide a short summary of domains in the category.
	Description string `json:"description"`
	// Specify the category name.
	Name string `json:"name"`
	// Provide all subcategories for this category.
	Subcategories []CategorySubcategory `json:"subcategories"`
	JSON          categoryJSON          `json:"-"`
}

// categoryJSON contains the JSON metadata for the struct [Category]
type categoryJSON struct {
	ID            apijson.Field
	Beta          apijson.Field
	Class         apijson.Field
	Description   apijson.Field
	Name          apijson.Field
	Subcategories apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Category) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r categoryJSON) RawJSON() string {
	return r.raw
}

// Specify which account types can create policies for this category. `blocked`
// Blocks unconditionally for all accounts. `removalPending` Allows removal from
// policies but disables addition. `noBlock` Prevents blocking.
type CategoryClass string

const (
	CategoryClassFree           CategoryClass = "free"
	CategoryClassPremium        CategoryClass = "premium"
	CategoryClassBlocked        CategoryClass = "blocked"
	CategoryClassRemovalPending CategoryClass = "removalPending"
	CategoryClassNoBlock        CategoryClass = "noBlock"
)

func (r CategoryClass) IsKnown() bool {
	switch r {
	case CategoryClassFree, CategoryClassPremium, CategoryClassBlocked, CategoryClassRemovalPending, CategoryClassNoBlock:
		return true
	}
	return false
}

type CategorySubcategory struct {
	// Identify this category. Only one category per ID.
	ID int64 `json:"id"`
	// Indicate whether the category is in beta and subject to change.
	Beta bool `json:"beta"`
	// Specify which account types can create policies for this category. `blocked`
	// Blocks unconditionally for all accounts. `removalPending` Allows removal from
	// policies but disables addition. `noBlock` Prevents blocking.
	Class CategorySubcategoriesClass `json:"class"`
	// Provide a short summary of domains in the category.
	Description string `json:"description"`
	// Specify the category name.
	Name string                  `json:"name"`
	JSON categorySubcategoryJSON `json:"-"`
}

// categorySubcategoryJSON contains the JSON metadata for the struct
// [CategorySubcategory]
type categorySubcategoryJSON struct {
	ID          apijson.Field
	Beta        apijson.Field
	Class       apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CategorySubcategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r categorySubcategoryJSON) RawJSON() string {
	return r.raw
}

// Specify which account types can create policies for this category. `blocked`
// Blocks unconditionally for all accounts. `removalPending` Allows removal from
// policies but disables addition. `noBlock` Prevents blocking.
type CategorySubcategoriesClass string

const (
	CategorySubcategoriesClassFree           CategorySubcategoriesClass = "free"
	CategorySubcategoriesClassPremium        CategorySubcategoriesClass = "premium"
	CategorySubcategoriesClassBlocked        CategorySubcategoriesClass = "blocked"
	CategorySubcategoriesClassRemovalPending CategorySubcategoriesClass = "removalPending"
	CategorySubcategoriesClassNoBlock        CategorySubcategoriesClass = "noBlock"
)

func (r CategorySubcategoriesClass) IsKnown() bool {
	switch r {
	case CategorySubcategoriesClassFree, CategorySubcategoriesClassPremium, CategorySubcategoriesClassBlocked, CategorySubcategoriesClassRemovalPending, CategorySubcategoriesClassNoBlock:
		return true
	}
	return false
}

type GatewayCategoryListParams struct {
	// Provide the identifier string.
	AccountID param.Field[string] `path:"account_id,required"`
}
