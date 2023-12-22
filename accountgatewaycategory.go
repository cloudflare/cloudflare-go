// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountGatewayCategoryService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountGatewayCategoryService]
// method instead.
type AccountGatewayCategoryService struct {
	Options []option.RequestOption
}

// NewAccountGatewayCategoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountGatewayCategoryService(opts ...option.RequestOption) (r *AccountGatewayCategoryService) {
	r = &AccountGatewayCategoryService{}
	r.Options = opts
	return
}

// List all Categories.
func (r *AccountGatewayCategoryService) ZeroTrustGatewayCategoriesListCategories(ctx context.Context, accountID string, opts ...option.RequestOption) (res *CategoriesResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/gateway/categories", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CategoriesResponseCollection struct {
	Errors     []CategoriesResponseCollectionError    `json:"errors"`
	Messages   []CategoriesResponseCollectionMessage  `json:"messages"`
	Result     []CategoriesResponseCollectionResult   `json:"result"`
	ResultInfo CategoriesResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success CategoriesResponseCollectionSuccess `json:"success"`
	JSON    categoriesResponseCollectionJSON    `json:"-"`
}

// categoriesResponseCollectionJSON contains the JSON metadata for the struct
// [CategoriesResponseCollection]
type categoriesResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CategoriesResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CategoriesResponseCollectionError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    categoriesResponseCollectionErrorJSON `json:"-"`
}

// categoriesResponseCollectionErrorJSON contains the JSON metadata for the struct
// [CategoriesResponseCollectionError]
type categoriesResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CategoriesResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CategoriesResponseCollectionMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    categoriesResponseCollectionMessageJSON `json:"-"`
}

// categoriesResponseCollectionMessageJSON contains the JSON metadata for the
// struct [CategoriesResponseCollectionMessage]
type categoriesResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CategoriesResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CategoriesResponseCollectionResult struct {
	// The identifier for this category. There is only one category per id.
	ID int64 `json:"id"`
	// Whether the category is in beta and subject to change.
	Beta bool `json:"beta"`
	// Which account types are allowed to create policies based on this categories.
	// `blocked` categories are blocked unconditionally for all accounts.
	// `removalPending` categories can be removed from policies but not added.
	// `noBlock` categories cannot be blocked.
	Class CategoriesResponseCollectionResultClass `json:"class"`
	// A short summary of domains in the category.
	Description string `json:"description"`
	// The name of the category.
	Name string `json:"name"`
	// All subcategories for this category.
	Subcategories []CategoriesResponseCollectionResultSubcategory `json:"subcategories"`
	JSON          categoriesResponseCollectionResultJSON          `json:"-"`
}

// categoriesResponseCollectionResultJSON contains the JSON metadata for the struct
// [CategoriesResponseCollectionResult]
type categoriesResponseCollectionResultJSON struct {
	ID            apijson.Field
	Beta          apijson.Field
	Class         apijson.Field
	Description   apijson.Field
	Name          apijson.Field
	Subcategories apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CategoriesResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Which account types are allowed to create policies based on this categories.
// `blocked` categories are blocked unconditionally for all accounts.
// `removalPending` categories can be removed from policies but not added.
// `noBlock` categories cannot be blocked.
type CategoriesResponseCollectionResultClass string

const (
	CategoriesResponseCollectionResultClassFree           CategoriesResponseCollectionResultClass = "free"
	CategoriesResponseCollectionResultClassPremium        CategoriesResponseCollectionResultClass = "premium"
	CategoriesResponseCollectionResultClassBlocked        CategoriesResponseCollectionResultClass = "blocked"
	CategoriesResponseCollectionResultClassRemovalPending CategoriesResponseCollectionResultClass = "removalPending"
	CategoriesResponseCollectionResultClassNoBlock        CategoriesResponseCollectionResultClass = "noBlock"
)

type CategoriesResponseCollectionResultSubcategory struct {
	// The identifier for this category. There is only one category per id.
	ID int64 `json:"id"`
	// Whether the category is in beta and subject to change.
	Beta bool `json:"beta"`
	// Which account types are allowed to create policies based on this categories.
	// `blocked` categories are blocked unconditionally for all accounts.
	// `removalPending` categories can be removed from policies but not added.
	// `noBlock` categories cannot be blocked.
	Class CategoriesResponseCollectionResultSubcategoriesClass `json:"class"`
	// A short summary of domains in the category.
	Description string `json:"description"`
	// The name of the category.
	Name string                                            `json:"name"`
	JSON categoriesResponseCollectionResultSubcategoryJSON `json:"-"`
}

// categoriesResponseCollectionResultSubcategoryJSON contains the JSON metadata for
// the struct [CategoriesResponseCollectionResultSubcategory]
type categoriesResponseCollectionResultSubcategoryJSON struct {
	ID          apijson.Field
	Beta        apijson.Field
	Class       apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CategoriesResponseCollectionResultSubcategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Which account types are allowed to create policies based on this categories.
// `blocked` categories are blocked unconditionally for all accounts.
// `removalPending` categories can be removed from policies but not added.
// `noBlock` categories cannot be blocked.
type CategoriesResponseCollectionResultSubcategoriesClass string

const (
	CategoriesResponseCollectionResultSubcategoriesClassFree           CategoriesResponseCollectionResultSubcategoriesClass = "free"
	CategoriesResponseCollectionResultSubcategoriesClassPremium        CategoriesResponseCollectionResultSubcategoriesClass = "premium"
	CategoriesResponseCollectionResultSubcategoriesClassBlocked        CategoriesResponseCollectionResultSubcategoriesClass = "blocked"
	CategoriesResponseCollectionResultSubcategoriesClassRemovalPending CategoriesResponseCollectionResultSubcategoriesClass = "removalPending"
	CategoriesResponseCollectionResultSubcategoriesClassNoBlock        CategoriesResponseCollectionResultSubcategoriesClass = "noBlock"
)

type CategoriesResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                    `json:"total_count"`
	JSON       categoriesResponseCollectionResultInfoJSON `json:"-"`
}

// categoriesResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [CategoriesResponseCollectionResultInfo]
type categoriesResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CategoriesResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CategoriesResponseCollectionSuccess bool

const (
	CategoriesResponseCollectionSuccessTrue CategoriesResponseCollectionSuccess = true
)
