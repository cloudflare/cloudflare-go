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

// Fetches a list of all categories.
func (r *AccountGatewayCategoryService) ZeroTrustGatewayCategoriesListCategories(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/gateway/categories", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponse struct {
	Errors     []AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseError    `json:"errors"`
	Messages   []AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseMessage  `json:"messages"`
	Result     []AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResult   `json:"result"`
	ResultInfo AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseSuccess `json:"success"`
	JSON    accountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseJSON    `json:"-"`
}

// accountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseJSON
// contains the JSON metadata for the struct
// [AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponse]
type accountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseError struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    accountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseErrorJSON `json:"-"`
}

// accountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseError]
type accountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseMessage struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    accountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseMessageJSON `json:"-"`
}

// accountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseMessage]
type accountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResult struct {
	// The identifier for this category. There is only one category per ID.
	ID int64 `json:"id"`
	// True if the category is in beta and subject to change.
	Beta bool `json:"beta"`
	// Which account types are allowed to create policies based on this category.
	// `blocked` categories are blocked unconditionally for all accounts.
	// `removalPending` categories can be removed from policies but not added.
	// `noBlock` categories cannot be blocked.
	Class AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultClass `json:"class"`
	// A short summary of domains in the category.
	Description string `json:"description"`
	// The name of the category.
	Name string `json:"name"`
	// All subcategories for this category.
	Subcategories []AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultSubcategory `json:"subcategories"`
	JSON          accountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultJSON          `json:"-"`
}

// accountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultJSON
// contains the JSON metadata for the struct
// [AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResult]
type accountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultJSON struct {
	ID            apijson.Field
	Beta          apijson.Field
	Class         apijson.Field
	Description   apijson.Field
	Name          apijson.Field
	Subcategories apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Which account types are allowed to create policies based on this category.
// `blocked` categories are blocked unconditionally for all accounts.
// `removalPending` categories can be removed from policies but not added.
// `noBlock` categories cannot be blocked.
type AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultClass string

const (
	AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultClassFree           AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultClass = "free"
	AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultClassPremium        AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultClass = "premium"
	AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultClassBlocked        AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultClass = "blocked"
	AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultClassRemovalPending AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultClass = "removalPending"
	AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultClassNoBlock        AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultClass = "noBlock"
)

type AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultSubcategory struct {
	// The identifier for this category. There is only one category per ID.
	ID int64 `json:"id"`
	// True if the category is in beta and subject to change.
	Beta bool `json:"beta"`
	// Which account types are allowed to create policies based on this category.
	// `blocked` categories are blocked unconditionally for all accounts.
	// `removalPending` categories can be removed from policies but not added.
	// `noBlock` categories cannot be blocked.
	Class AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultSubcategoriesClass `json:"class"`
	// A short summary of domains in the category.
	Description string `json:"description"`
	// The name of the category.
	Name string                                                                                      `json:"name"`
	JSON accountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultSubcategoryJSON `json:"-"`
}

// accountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultSubcategoryJSON
// contains the JSON metadata for the struct
// [AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultSubcategory]
type accountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultSubcategoryJSON struct {
	ID          apijson.Field
	Beta        apijson.Field
	Class       apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultSubcategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Which account types are allowed to create policies based on this category.
// `blocked` categories are blocked unconditionally for all accounts.
// `removalPending` categories can be removed from policies but not added.
// `noBlock` categories cannot be blocked.
type AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultSubcategoriesClass string

const (
	AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultSubcategoriesClassFree           AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultSubcategoriesClass = "free"
	AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultSubcategoriesClassPremium        AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultSubcategoriesClass = "premium"
	AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultSubcategoriesClassBlocked        AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultSubcategoriesClass = "blocked"
	AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultSubcategoriesClassRemovalPending AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultSubcategoriesClass = "removalPending"
	AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultSubcategoriesClassNoBlock        AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultSubcategoriesClass = "noBlock"
)

type AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                              `json:"total_count"`
	JSON       accountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultInfoJSON `json:"-"`
}

// accountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultInfo]
type accountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseSuccess bool

const (
	AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseSuccessTrue AccountGatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseSuccess = true
)
