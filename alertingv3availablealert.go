// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// AlertingV3AvailableAlertService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAlertingV3AvailableAlertService] method instead.
type AlertingV3AvailableAlertService struct {
	Options []option.RequestOption
}

// NewAlertingV3AvailableAlertService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAlertingV3AvailableAlertService(opts ...option.RequestOption) (r *AlertingV3AvailableAlertService) {
	r = &AlertingV3AvailableAlertService{}
	r.Options = opts
	return
}

// Gets a list of all alert types for which an account is eligible.
func (r *AlertingV3AvailableAlertService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AlertingV3AvailableAlertListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AlertingV3AvailableAlertListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/available_alerts", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [AlertingV3AvailableAlertListResponseUnknown],
// [AlertingV3AvailableAlertListResponseArray] or [shared.UnionString].
type AlertingV3AvailableAlertListResponse interface {
	ImplementsAlertingV3AvailableAlertListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AlertingV3AvailableAlertListResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AlertingV3AvailableAlertListResponseArray []interface{}

func (r AlertingV3AvailableAlertListResponseArray) ImplementsAlertingV3AvailableAlertListResponse() {}

type AlertingV3AvailableAlertListResponseEnvelope struct {
	Errors   []AlertingV3AvailableAlertListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AlertingV3AvailableAlertListResponseEnvelopeMessages `json:"messages,required"`
	Result   AlertingV3AvailableAlertListResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AlertingV3AvailableAlertListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AlertingV3AvailableAlertListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       alertingV3AvailableAlertListResponseEnvelopeJSON       `json:"-"`
}

// alertingV3AvailableAlertListResponseEnvelopeJSON contains the JSON metadata for
// the struct [AlertingV3AvailableAlertListResponseEnvelope]
type alertingV3AvailableAlertListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3AvailableAlertListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3AvailableAlertListResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    alertingV3AvailableAlertListResponseEnvelopeErrorsJSON `json:"-"`
}

// alertingV3AvailableAlertListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AlertingV3AvailableAlertListResponseEnvelopeErrors]
type alertingV3AvailableAlertListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3AvailableAlertListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3AvailableAlertListResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    alertingV3AvailableAlertListResponseEnvelopeMessagesJSON `json:"-"`
}

// alertingV3AvailableAlertListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AlertingV3AvailableAlertListResponseEnvelopeMessages]
type alertingV3AvailableAlertListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3AvailableAlertListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AlertingV3AvailableAlertListResponseEnvelopeSuccess bool

const (
	AlertingV3AvailableAlertListResponseEnvelopeSuccessTrue AlertingV3AvailableAlertListResponseEnvelopeSuccess = true
)

type AlertingV3AvailableAlertListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                    `json:"total_count"`
	JSON       alertingV3AvailableAlertListResponseEnvelopeResultInfoJSON `json:"-"`
}

// alertingV3AvailableAlertListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [AlertingV3AvailableAlertListResponseEnvelopeResultInfo]
type alertingV3AvailableAlertListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3AvailableAlertListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
