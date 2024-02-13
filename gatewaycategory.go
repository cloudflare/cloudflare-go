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
func (r *GatewayCategoryService) ZeroTrustGatewayCategoriesListCategories(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelope
	path := fmt.Sprintf("accounts/%s/gateway/categories", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponse struct {
	// The identifier for this category. There is only one category per ID.
	ID int64 `json:"id"`
	// True if the category is in beta and subject to change.
	Beta bool `json:"beta"`
	// Which account types are allowed to create policies based on this category.
	// `blocked` categories are blocked unconditionally for all accounts.
	// `removalPending` categories can be removed from policies but not added.
	// `noBlock` categories cannot be blocked.
	Class GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseClass `json:"class"`
	// A short summary of domains in the category.
	Description string `json:"description"`
	// The name of the category.
	Name string `json:"name"`
	// All subcategories for this category.
	Subcategories []GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseSubcategory `json:"subcategories"`
	JSON          gatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseJSON          `json:"-"`
}

// gatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseJSON contains the
// JSON metadata for the struct
// [GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponse]
type gatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseJSON struct {
	ID            apijson.Field
	Beta          apijson.Field
	Class         apijson.Field
	Description   apijson.Field
	Name          apijson.Field
	Subcategories apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Which account types are allowed to create policies based on this category.
// `blocked` categories are blocked unconditionally for all accounts.
// `removalPending` categories can be removed from policies but not added.
// `noBlock` categories cannot be blocked.
type GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseClass string

const (
	GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseClassFree           GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseClass = "free"
	GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseClassPremium        GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseClass = "premium"
	GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseClassBlocked        GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseClass = "blocked"
	GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseClassRemovalPending GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseClass = "removalPending"
	GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseClassNoBlock        GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseClass = "noBlock"
)

type GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseSubcategory struct {
	// The identifier for this category. There is only one category per ID.
	ID int64 `json:"id"`
	// True if the category is in beta and subject to change.
	Beta bool `json:"beta"`
	// Which account types are allowed to create policies based on this category.
	// `blocked` categories are blocked unconditionally for all accounts.
	// `removalPending` categories can be removed from policies but not added.
	// `noBlock` categories cannot be blocked.
	Class GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseSubcategoriesClass `json:"class"`
	// A short summary of domains in the category.
	Description string `json:"description"`
	// The name of the category.
	Name string                                                                         `json:"name"`
	JSON gatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseSubcategoryJSON `json:"-"`
}

// gatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseSubcategoryJSON
// contains the JSON metadata for the struct
// [GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseSubcategory]
type gatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseSubcategoryJSON struct {
	ID          apijson.Field
	Beta        apijson.Field
	Class       apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseSubcategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Which account types are allowed to create policies based on this category.
// `blocked` categories are blocked unconditionally for all accounts.
// `removalPending` categories can be removed from policies but not added.
// `noBlock` categories cannot be blocked.
type GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseSubcategoriesClass string

const (
	GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseSubcategoriesClassFree           GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseSubcategoriesClass = "free"
	GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseSubcategoriesClassPremium        GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseSubcategoriesClass = "premium"
	GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseSubcategoriesClassBlocked        GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseSubcategoriesClass = "blocked"
	GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseSubcategoriesClassRemovalPending GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseSubcategoriesClass = "removalPending"
	GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseSubcategoriesClassNoBlock        GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseSubcategoriesClass = "noBlock"
)

type GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelope struct {
	Errors   []GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelopeMessages `json:"messages,required"`
	Result   []GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       gatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelopeJSON       `json:"-"`
}

// gatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelope]
type gatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelopeErrors struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    gatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelopeErrors]
type gatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelopeMessages struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    gatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelopeMessages]
type gatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelopeSuccess bool

const (
	GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelopeSuccessTrue GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelopeSuccess = true
)

type GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                               `json:"total_count"`
	JSON       gatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelopeResultInfoJSON `json:"-"`
}

// gatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelopeResultInfo]
type gatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayCategoryZeroTrustGatewayCategoriesListCategoriesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
