// File generated from our OpenAPI spec by Stainless.

package user

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// BillingHistoryService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewBillingHistoryService] method
// instead.
type BillingHistoryService struct {
	Options []option.RequestOption
}

// NewBillingHistoryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBillingHistoryService(opts ...option.RequestOption) (r *BillingHistoryService) {
	r = &BillingHistoryService{}
	r.Options = opts
	return
}

// Accesses your billing history object.
func (r *BillingHistoryService) Get(ctx context.Context, query BillingHistoryGetParams, opts ...option.RequestOption) (res *[]BillingHistoryGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env BillingHistoryGetResponseEnvelope
	path := "user/billing/history"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BillingHistoryGetResponse struct {
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
	Type string                        `json:"type,required"`
	Zone BillingHistoryGetResponseZone `json:"zone,required"`
	JSON billingHistoryGetResponseJSON `json:"-"`
}

// billingHistoryGetResponseJSON contains the JSON metadata for the struct
// [BillingHistoryGetResponse]
type billingHistoryGetResponseJSON struct {
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

func (r *BillingHistoryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingHistoryGetResponseJSON) RawJSON() string {
	return r.raw
}

type BillingHistoryGetResponseZone struct {
	Name interface{}                       `json:"name"`
	JSON billingHistoryGetResponseZoneJSON `json:"-"`
}

// billingHistoryGetResponseZoneJSON contains the JSON metadata for the struct
// [BillingHistoryGetResponseZone]
type billingHistoryGetResponseZoneJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingHistoryGetResponseZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingHistoryGetResponseZoneJSON) RawJSON() string {
	return r.raw
}

type BillingHistoryGetParams struct {
	// Field to order billing history by.
	Order param.Field[BillingHistoryGetParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [BillingHistoryGetParams]'s query parameters as
// `url.Values`.
func (r BillingHistoryGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Field to order billing history by.
type BillingHistoryGetParamsOrder string

const (
	BillingHistoryGetParamsOrderType      BillingHistoryGetParamsOrder = "type"
	BillingHistoryGetParamsOrderOccuredAt BillingHistoryGetParamsOrder = "occured_at"
	BillingHistoryGetParamsOrderAction    BillingHistoryGetParamsOrder = "action"
)

type BillingHistoryGetResponseEnvelope struct {
	Errors   []BillingHistoryGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []BillingHistoryGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []BillingHistoryGetResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    BillingHistoryGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo BillingHistoryGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       billingHistoryGetResponseEnvelopeJSON       `json:"-"`
}

// billingHistoryGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [BillingHistoryGetResponseEnvelope]
type billingHistoryGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingHistoryGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingHistoryGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BillingHistoryGetResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    billingHistoryGetResponseEnvelopeErrorsJSON `json:"-"`
}

// billingHistoryGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [BillingHistoryGetResponseEnvelopeErrors]
type billingHistoryGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingHistoryGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingHistoryGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type BillingHistoryGetResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    billingHistoryGetResponseEnvelopeMessagesJSON `json:"-"`
}

// billingHistoryGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [BillingHistoryGetResponseEnvelopeMessages]
type billingHistoryGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingHistoryGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingHistoryGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type BillingHistoryGetResponseEnvelopeSuccess bool

const (
	BillingHistoryGetResponseEnvelopeSuccessTrue BillingHistoryGetResponseEnvelopeSuccess = true
)

type BillingHistoryGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                         `json:"total_count"`
	JSON       billingHistoryGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// billingHistoryGetResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [BillingHistoryGetResponseEnvelopeResultInfo]
type billingHistoryGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingHistoryGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingHistoryGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
