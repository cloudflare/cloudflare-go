// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// AlertingV3DestinationEligibleService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAlertingV3DestinationEligibleService] method instead.
type AlertingV3DestinationEligibleService struct {
	Options []option.RequestOption
}

// NewAlertingV3DestinationEligibleService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAlertingV3DestinationEligibleService(opts ...option.RequestOption) (r *AlertingV3DestinationEligibleService) {
	r = &AlertingV3DestinationEligibleService{}
	r.Options = opts
	return
}

// Get a list of all delivery mechanism types for which an account is eligible.
func (r *AlertingV3DestinationEligibleService) Get(ctx context.Context, query AlertingV3DestinationEligibleGetParams, opts ...option.RequestOption) (res *AlertingV3DestinationEligibleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AlertingV3DestinationEligibleGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/eligible", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [AlertingV3DestinationEligibleGetResponseUnknown],
// [AlertingV3DestinationEligibleGetResponseArray] or [shared.UnionString].
type AlertingV3DestinationEligibleGetResponse interface {
	ImplementsAlertingV3DestinationEligibleGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AlertingV3DestinationEligibleGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AlertingV3DestinationEligibleGetResponseArray []interface{}

func (r AlertingV3DestinationEligibleGetResponseArray) ImplementsAlertingV3DestinationEligibleGetResponse() {
}

type AlertingV3DestinationEligibleGetParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type AlertingV3DestinationEligibleGetResponseEnvelope struct {
	Errors   []AlertingV3DestinationEligibleGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AlertingV3DestinationEligibleGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AlertingV3DestinationEligibleGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AlertingV3DestinationEligibleGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AlertingV3DestinationEligibleGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       alertingV3DestinationEligibleGetResponseEnvelopeJSON       `json:"-"`
}

// alertingV3DestinationEligibleGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [AlertingV3DestinationEligibleGetResponseEnvelope]
type alertingV3DestinationEligibleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationEligibleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationEligibleGetResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    alertingV3DestinationEligibleGetResponseEnvelopeErrorsJSON `json:"-"`
}

// alertingV3DestinationEligibleGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AlertingV3DestinationEligibleGetResponseEnvelopeErrors]
type alertingV3DestinationEligibleGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationEligibleGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationEligibleGetResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    alertingV3DestinationEligibleGetResponseEnvelopeMessagesJSON `json:"-"`
}

// alertingV3DestinationEligibleGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [AlertingV3DestinationEligibleGetResponseEnvelopeMessages]
type alertingV3DestinationEligibleGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationEligibleGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AlertingV3DestinationEligibleGetResponseEnvelopeSuccess bool

const (
	AlertingV3DestinationEligibleGetResponseEnvelopeSuccessTrue AlertingV3DestinationEligibleGetResponseEnvelopeSuccess = true
)

type AlertingV3DestinationEligibleGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                        `json:"total_count"`
	JSON       alertingV3DestinationEligibleGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// alertingV3DestinationEligibleGetResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [AlertingV3DestinationEligibleGetResponseEnvelopeResultInfo]
type alertingV3DestinationEligibleGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationEligibleGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
