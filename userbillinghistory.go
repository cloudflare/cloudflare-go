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
func (r *UserBillingHistoryService) UserBillingHistoryBillingHistoryDetails(ctx context.Context, query UserBillingHistoryUserBillingHistoryBillingHistoryDetailsParams, opts ...option.RequestOption) (res *[]UserBillingHistoryUserBillingHistoryBillingHistoryDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelope
	path := "user/billing/history"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type UserBillingHistoryUserBillingHistoryBillingHistoryDetailsResponse struct {
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
	Type string                                                                `json:"type,required"`
	Zone UserBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseZone `json:"zone,required"`
	JSON userBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseJSON `json:"-"`
}

// userBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseJSON contains
// the JSON metadata for the struct
// [UserBillingHistoryUserBillingHistoryBillingHistoryDetailsResponse]
type userBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseJSON struct {
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

func (r *UserBillingHistoryUserBillingHistoryBillingHistoryDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseZone struct {
	Name interface{}                                                               `json:"name"`
	JSON userBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseZoneJSON `json:"-"`
}

// userBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseZoneJSON
// contains the JSON metadata for the struct
// [UserBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseZone]
type userBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseZoneJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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

type UserBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelope struct {
	Errors   []UserBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelopeMessages `json:"messages,required"`
	Result   []UserBillingHistoryUserBillingHistoryBillingHistoryDetailsResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    UserBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo UserBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelopeResultInfo `json:"result_info"`
	JSON       userBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelopeJSON       `json:"-"`
}

// userBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [UserBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelope]
type userBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelopeErrors struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    userBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelopeErrorsJSON `json:"-"`
}

// userBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [UserBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelopeErrors]
type userBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelopeMessages struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    userBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelopeMessagesJSON `json:"-"`
}

// userBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [UserBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelopeMessages]
type userBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelopeSuccess bool

const (
	UserBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelopeSuccessTrue UserBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelopeSuccess = true
)

type UserBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                 `json:"total_count"`
	JSON       userBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelopeResultInfoJSON `json:"-"`
}

// userBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [UserBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelopeResultInfo]
type userBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBillingHistoryUserBillingHistoryBillingHistoryDetailsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
