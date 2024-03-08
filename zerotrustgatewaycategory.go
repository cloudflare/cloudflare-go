// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *ZeroTrustGatewayCategoryService) List(ctx context.Context, query ZeroTrustGatewayCategoryListParams, opts ...option.RequestOption) (res *[]ZeroTrustGatewayCategoryListResponse, err error) {
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

type ZeroTrustGatewayCategoryListResponse struct {
	// The identifier for this category. There is only one category per ID.
	ID int64 `json:"id"`
	// True if the category is in beta and subject to change.
	Beta bool `json:"beta"`
	// Which account types are allowed to create policies based on this category.
	// `blocked` categories are blocked unconditionally for all accounts.
	// `removalPending` categories can be removed from policies but not added.
	// `noBlock` categories cannot be blocked.
	Class ZeroTrustGatewayCategoryListResponseClass `json:"class"`
	// A short summary of domains in the category.
	Description string `json:"description"`
	// The name of the category.
	Name string `json:"name"`
	// All subcategories for this category.
	Subcategories []ZeroTrustGatewayCategoryListResponseSubcategory `json:"subcategories"`
	JSON          zeroTrustGatewayCategoryListResponseJSON          `json:"-"`
}

// zeroTrustGatewayCategoryListResponseJSON contains the JSON metadata for the
// struct [ZeroTrustGatewayCategoryListResponse]
type zeroTrustGatewayCategoryListResponseJSON struct {
	ID            apijson.Field
	Beta          apijson.Field
	Class         apijson.Field
	Description   apijson.Field
	Name          apijson.Field
	Subcategories apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustGatewayCategoryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayCategoryListResponseJSON) RawJSON() string {
	return r.raw
}

// Which account types are allowed to create policies based on this category.
// `blocked` categories are blocked unconditionally for all accounts.
// `removalPending` categories can be removed from policies but not added.
// `noBlock` categories cannot be blocked.
type ZeroTrustGatewayCategoryListResponseClass string

const (
	ZeroTrustGatewayCategoryListResponseClassFree           ZeroTrustGatewayCategoryListResponseClass = "free"
	ZeroTrustGatewayCategoryListResponseClassPremium        ZeroTrustGatewayCategoryListResponseClass = "premium"
	ZeroTrustGatewayCategoryListResponseClassBlocked        ZeroTrustGatewayCategoryListResponseClass = "blocked"
	ZeroTrustGatewayCategoryListResponseClassRemovalPending ZeroTrustGatewayCategoryListResponseClass = "removalPending"
	ZeroTrustGatewayCategoryListResponseClassNoBlock        ZeroTrustGatewayCategoryListResponseClass = "noBlock"
)

type ZeroTrustGatewayCategoryListResponseSubcategory struct {
	// The identifier for this category. There is only one category per ID.
	ID int64 `json:"id"`
	// True if the category is in beta and subject to change.
	Beta bool `json:"beta"`
	// Which account types are allowed to create policies based on this category.
	// `blocked` categories are blocked unconditionally for all accounts.
	// `removalPending` categories can be removed from policies but not added.
	// `noBlock` categories cannot be blocked.
	Class ZeroTrustGatewayCategoryListResponseSubcategoriesClass `json:"class"`
	// A short summary of domains in the category.
	Description string `json:"description"`
	// The name of the category.
	Name string                                              `json:"name"`
	JSON zeroTrustGatewayCategoryListResponseSubcategoryJSON `json:"-"`
}

// zeroTrustGatewayCategoryListResponseSubcategoryJSON contains the JSON metadata
// for the struct [ZeroTrustGatewayCategoryListResponseSubcategory]
type zeroTrustGatewayCategoryListResponseSubcategoryJSON struct {
	ID          apijson.Field
	Beta        apijson.Field
	Class       apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayCategoryListResponseSubcategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayCategoryListResponseSubcategoryJSON) RawJSON() string {
	return r.raw
}

// Which account types are allowed to create policies based on this category.
// `blocked` categories are blocked unconditionally for all accounts.
// `removalPending` categories can be removed from policies but not added.
// `noBlock` categories cannot be blocked.
type ZeroTrustGatewayCategoryListResponseSubcategoriesClass string

const (
	ZeroTrustGatewayCategoryListResponseSubcategoriesClassFree           ZeroTrustGatewayCategoryListResponseSubcategoriesClass = "free"
	ZeroTrustGatewayCategoryListResponseSubcategoriesClassPremium        ZeroTrustGatewayCategoryListResponseSubcategoriesClass = "premium"
	ZeroTrustGatewayCategoryListResponseSubcategoriesClassBlocked        ZeroTrustGatewayCategoryListResponseSubcategoriesClass = "blocked"
	ZeroTrustGatewayCategoryListResponseSubcategoriesClassRemovalPending ZeroTrustGatewayCategoryListResponseSubcategoriesClass = "removalPending"
	ZeroTrustGatewayCategoryListResponseSubcategoriesClassNoBlock        ZeroTrustGatewayCategoryListResponseSubcategoriesClass = "noBlock"
)

type ZeroTrustGatewayCategoryListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZeroTrustGatewayCategoryListResponseEnvelope struct {
	Errors   []ZeroTrustGatewayCategoryListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayCategoryListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustGatewayCategoryListResponse                 `json:"result,required,nullable"`
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

func (r zeroTrustGatewayCategoryListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustGatewayCategoryListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustGatewayCategoryListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustGatewayCategoryListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
