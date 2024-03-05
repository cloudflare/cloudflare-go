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
func (r *UserBillingHistoryService) Get(ctx context.Context, query UserBillingHistoryGetParams, opts ...option.RequestOption) (res *[]UserBillingHistoryGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserBillingHistoryGetResponseEnvelope
	path := "user/billing/history"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type UserBillingHistoryGetResponse struct {
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
	Type string                            `json:"type,required"`
	Zone UserBillingHistoryGetResponseZone `json:"zone,required"`
	JSON userBillingHistoryGetResponseJSON `json:"-"`
}

// userBillingHistoryGetResponseJSON contains the JSON metadata for the struct
// [UserBillingHistoryGetResponse]
type userBillingHistoryGetResponseJSON struct {
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

func (r *UserBillingHistoryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserBillingHistoryGetResponseZone struct {
	Name interface{}                           `json:"name"`
	JSON userBillingHistoryGetResponseZoneJSON `json:"-"`
}

// userBillingHistoryGetResponseZoneJSON contains the JSON metadata for the struct
// [UserBillingHistoryGetResponseZone]
type userBillingHistoryGetResponseZoneJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBillingHistoryGetResponseZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserBillingHistoryGetParams struct {
	// Field to order billing history by.
	Order param.Field[UserBillingHistoryGetParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [UserBillingHistoryGetParams]'s query parameters as
// `url.Values`.
func (r UserBillingHistoryGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Field to order billing history by.
type UserBillingHistoryGetParamsOrder string

const (
	UserBillingHistoryGetParamsOrderType      UserBillingHistoryGetParamsOrder = "type"
	UserBillingHistoryGetParamsOrderOccuredAt UserBillingHistoryGetParamsOrder = "occured_at"
	UserBillingHistoryGetParamsOrderAction    UserBillingHistoryGetParamsOrder = "action"
)

type UserBillingHistoryGetResponseEnvelope struct {
	Errors   []UserBillingHistoryGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserBillingHistoryGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []UserBillingHistoryGetResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    UserBillingHistoryGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo UserBillingHistoryGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       userBillingHistoryGetResponseEnvelopeJSON       `json:"-"`
}

// userBillingHistoryGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [UserBillingHistoryGetResponseEnvelope]
type userBillingHistoryGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBillingHistoryGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserBillingHistoryGetResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    userBillingHistoryGetResponseEnvelopeErrorsJSON `json:"-"`
}

// userBillingHistoryGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [UserBillingHistoryGetResponseEnvelopeErrors]
type userBillingHistoryGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBillingHistoryGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserBillingHistoryGetResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    userBillingHistoryGetResponseEnvelopeMessagesJSON `json:"-"`
}

// userBillingHistoryGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [UserBillingHistoryGetResponseEnvelopeMessages]
type userBillingHistoryGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBillingHistoryGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserBillingHistoryGetResponseEnvelopeSuccess bool

const (
	UserBillingHistoryGetResponseEnvelopeSuccessTrue UserBillingHistoryGetResponseEnvelopeSuccess = true
)

type UserBillingHistoryGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                             `json:"total_count"`
	JSON       userBillingHistoryGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// userBillingHistoryGetResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [UserBillingHistoryGetResponseEnvelopeResultInfo]
type userBillingHistoryGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBillingHistoryGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
