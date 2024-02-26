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

// AlertingV3Service contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAlertingV3Service] method instead.
type AlertingV3Service struct {
	Options      []option.RequestOption
	Destinations *AlertingV3DestinationService
	Histories    *AlertingV3HistoryService
	Policies     *AlertingV3PolicyService
}

// NewAlertingV3Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAlertingV3Service(opts ...option.RequestOption) (r *AlertingV3Service) {
	r = &AlertingV3Service{}
	r.Options = opts
	r.Destinations = NewAlertingV3DestinationService(opts...)
	r.Histories = NewAlertingV3HistoryService(opts...)
	r.Policies = NewAlertingV3PolicyService(opts...)
	return
}

// Gets a list of all alert types for which an account is eligible.
func (r *AlertingV3Service) List(ctx context.Context, query AlertingV3ListParams, opts ...option.RequestOption) (res *AlertingV3ListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AlertingV3ListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/available_alerts", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [AlertingV3ListResponseUnknown],
// [AlertingV3ListResponseArray] or [shared.UnionString].
type AlertingV3ListResponse interface {
	ImplementsAlertingV3ListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AlertingV3ListResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AlertingV3ListResponseArray []interface{}

func (r AlertingV3ListResponseArray) ImplementsAlertingV3ListResponse() {}

type AlertingV3ListParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type AlertingV3ListResponseEnvelope struct {
	Errors   []AlertingV3ListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AlertingV3ListResponseEnvelopeMessages `json:"messages,required"`
	Result   AlertingV3ListResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AlertingV3ListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AlertingV3ListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       alertingV3ListResponseEnvelopeJSON       `json:"-"`
}

// alertingV3ListResponseEnvelopeJSON contains the JSON metadata for the struct
// [AlertingV3ListResponseEnvelope]
type alertingV3ListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3ListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3ListResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    alertingV3ListResponseEnvelopeErrorsJSON `json:"-"`
}

// alertingV3ListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AlertingV3ListResponseEnvelopeErrors]
type alertingV3ListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3ListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3ListResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    alertingV3ListResponseEnvelopeMessagesJSON `json:"-"`
}

// alertingV3ListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AlertingV3ListResponseEnvelopeMessages]
type alertingV3ListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3ListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AlertingV3ListResponseEnvelopeSuccess bool

const (
	AlertingV3ListResponseEnvelopeSuccessTrue AlertingV3ListResponseEnvelopeSuccess = true
)

type AlertingV3ListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                      `json:"total_count"`
	JSON       alertingV3ListResponseEnvelopeResultInfoJSON `json:"-"`
}

// alertingV3ListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [AlertingV3ListResponseEnvelopeResultInfo]
type alertingV3ListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3ListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
