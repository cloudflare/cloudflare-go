// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
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
func (r *GatewayCategoryService) List(ctx context.Context, query GatewayCategoryListParams, opts ...option.RequestOption) (res *[]ZeroTrustGatewayCategories, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayCategoryListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/gateway/categories", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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

type GatewayCategoryListResponseEnvelope struct {
	Errors   []GatewayCategoryListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayCategoryListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustGatewayCategories                  `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    GatewayCategoryListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo GatewayCategoryListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       gatewayCategoryListResponseEnvelopeJSON       `json:"-"`
}

// gatewayCategoryListResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayCategoryListResponseEnvelope]
type gatewayCategoryListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayCategoryListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayCategoryListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type GatewayCategoryListResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    gatewayCategoryListResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayCategoryListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [GatewayCategoryListResponseEnvelopeErrors]
type gatewayCategoryListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayCategoryListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayCategoryListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type GatewayCategoryListResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    gatewayCategoryListResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayCategoryListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [GatewayCategoryListResponseEnvelopeMessages]
type gatewayCategoryListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayCategoryListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayCategoryListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayCategoryListResponseEnvelopeSuccess bool

const (
	GatewayCategoryListResponseEnvelopeSuccessTrue GatewayCategoryListResponseEnvelopeSuccess = true
)

func (r GatewayCategoryListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayCategoryListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayCategoryListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                           `json:"total_count"`
	JSON       gatewayCategoryListResponseEnvelopeResultInfoJSON `json:"-"`
}

// gatewayCategoryListResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [GatewayCategoryListResponseEnvelopeResultInfo]
type gatewayCategoryListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayCategoryListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayCategoryListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
