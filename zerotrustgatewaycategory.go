// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZeroTrustGatewayCategoryService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZeroTrustGatewayCategoryService] method instead.
type ZeroTrustGatewayCategoryService struct {
	Options []option.RequestOption
}

// NewZeroTrustGatewayCategoryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustGatewayCategoryService(opts ...option.RequestOption) (r *ZeroTrustGatewayCategoryService) {
	r = &ZeroTrustGatewayCategoryService{}
	r.Options = opts
	return
}

// Fetches a list of all categories.
func (r *ZeroTrustGatewayCategoryService) List(ctx context.Context, query ZeroTrustGatewayCategoryListParams, opts ...option.RequestOption) (res *[]ZeroTrustGatewayCategories, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayCategoryListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/gateway/categories", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
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

type ZeroTrustGatewayCategoryListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZeroTrustGatewayCategoryListResponseEnvelope struct {
	Errors   []ZeroTrustGatewayCategoryListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayCategoryListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustGatewayCategories                           `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustGatewayCategoryListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustGatewayCategoryListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustGatewayCategoryListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustGatewayCategoryListResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayCategoryListResponseEnvelope]
type zeroTrustGatewayCategoryListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayCategoryListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayCategoryListResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustGatewayCategoryListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayCategoryListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayCategoryListResponseEnvelopeErrors]
type zeroTrustGatewayCategoryListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayCategoryListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayCategoryListResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zeroTrustGatewayCategoryListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayCategoryListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayCategoryListResponseEnvelopeMessages]
type zeroTrustGatewayCategoryListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayCategoryListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustGatewayCategoryListResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayCategoryListResponseEnvelopeSuccessTrue ZeroTrustGatewayCategoryListResponseEnvelopeSuccess = true
)

type ZeroTrustGatewayCategoryListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                    `json:"total_count"`
	JSON       zeroTrustGatewayCategoryListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustGatewayCategoryListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayCategoryListResponseEnvelopeResultInfo]
type zeroTrustGatewayCategoryListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayCategoryListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
