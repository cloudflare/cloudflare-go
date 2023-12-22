// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// UserBillingHistoryService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewUserBillingHistoryService] method
// instead.
type UserBillingHistoryService struct {
	Options []option.RequestOption
}

// NewUserBillingHistoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserBillingHistoryService(opts ...option.RequestOption) (r *UserBillingHistoryService) {
	r = &UserBillingHistoryService{}
	r.Options = opts
	return
}

// Accesses your billing history object.
func (r *UserBillingHistoryService) UserBillingHistoryBillingHistoryDetails(ctx context.Context, query UserBillingHistoryUserBillingHistoryBillingHistoryDetailsParams, opts ...option.RequestOption) (res *BillingHistoryCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/billing/history"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type BillingHistoryCollection struct {
	Errors     []BillingHistoryCollectionError    `json:"errors"`
	Messages   []BillingHistoryCollectionMessage  `json:"messages"`
	Result     []BillingHistoryCollectionResult   `json:"result"`
	ResultInfo BillingHistoryCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success BillingHistoryCollectionSuccess `json:"success"`
	JSON    billingHistoryCollectionJSON    `json:"-"`
}

// billingHistoryCollectionJSON contains the JSON metadata for the struct
// [BillingHistoryCollection]
type billingHistoryCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingHistoryCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BillingHistoryCollectionError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    billingHistoryCollectionErrorJSON `json:"-"`
}

// billingHistoryCollectionErrorJSON contains the JSON metadata for the struct
// [BillingHistoryCollectionError]
type billingHistoryCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingHistoryCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BillingHistoryCollectionMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    billingHistoryCollectionMessageJSON `json:"-"`
}

// billingHistoryCollectionMessageJSON contains the JSON metadata for the struct
// [BillingHistoryCollectionMessage]
type billingHistoryCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingHistoryCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BillingHistoryCollectionResult struct {
	// Billing item identifier tag.
	ID string `json:"id,required"`
	// The billing item action.
	Action string `json:"action,required"`
	// The amount associated with this billing item.
	Amount float64 `json:"amount,required"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency,required"`
	// The billing item description.
	Description string `json:"description,required"`
	// When the billing item was created.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// The billing item type.
	Type string                             `json:"type,required"`
	Zone BillingHistoryCollectionResultZone `json:"zone,required"`
	JSON billingHistoryCollectionResultJSON `json:"-"`
}

// billingHistoryCollectionResultJSON contains the JSON metadata for the struct
// [BillingHistoryCollectionResult]
type billingHistoryCollectionResultJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Amount      apijson.Field
	Currency    apijson.Field
	Description apijson.Field
	OccurredAt  apijson.Field
	Type        apijson.Field
	Zone        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingHistoryCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BillingHistoryCollectionResultZone struct {
	Name interface{}                            `json:"name"`
	JSON billingHistoryCollectionResultZoneJSON `json:"-"`
}

// billingHistoryCollectionResultZoneJSON contains the JSON metadata for the struct
// [BillingHistoryCollectionResultZone]
type billingHistoryCollectionResultZoneJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingHistoryCollectionResultZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BillingHistoryCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                `json:"total_count"`
	JSON       billingHistoryCollectionResultInfoJSON `json:"-"`
}

// billingHistoryCollectionResultInfoJSON contains the JSON metadata for the struct
// [BillingHistoryCollectionResultInfo]
type billingHistoryCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingHistoryCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type BillingHistoryCollectionSuccess bool

const (
	BillingHistoryCollectionSuccessTrue BillingHistoryCollectionSuccess = true
)

type UserBillingHistoryUserBillingHistoryBillingHistoryDetailsParams struct {
	// Field to order billing history by.
	Order param.Field[UserBillingHistoryUserBillingHistoryBillingHistoryDetailsParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes
// [UserBillingHistoryUserBillingHistoryBillingHistoryDetailsParams]'s query
// parameters as `url.Values`.
func (r UserBillingHistoryUserBillingHistoryBillingHistoryDetailsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Field to order billing history by.
type UserBillingHistoryUserBillingHistoryBillingHistoryDetailsParamsOrder string

const (
	UserBillingHistoryUserBillingHistoryBillingHistoryDetailsParamsOrderType      UserBillingHistoryUserBillingHistoryBillingHistoryDetailsParamsOrder = "type"
	UserBillingHistoryUserBillingHistoryBillingHistoryDetailsParamsOrderOccuredAt UserBillingHistoryUserBillingHistoryBillingHistoryDetailsParamsOrder = "occured_at"
	UserBillingHistoryUserBillingHistoryBillingHistoryDetailsParamsOrderAction    UserBillingHistoryUserBillingHistoryBillingHistoryDetailsParamsOrder = "action"
)
