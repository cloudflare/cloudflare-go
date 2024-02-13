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
func (r *AlertingV3DestinationEligibleService) NotificationMechanismEligibilityGetDeliveryMechanismEligibility(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AlertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AlertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/eligible", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by
// [AlertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseUnknown],
// [AlertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseArray]
// or [shared.UnionString].
type AlertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponse interface {
	ImplementsAlertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AlertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AlertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseArray []interface{}

func (r AlertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseArray) ImplementsAlertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponse() {
}

type AlertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelope struct {
	Errors   []AlertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AlertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelopeMessages `json:"messages,required"`
	Result   AlertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AlertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AlertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelopeResultInfo `json:"result_info"`
	JSON       alertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelopeJSON       `json:"-"`
}

// alertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AlertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelope]
type alertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelopeErrors struct {
	Code    int64                                                                                                                  `json:"code,required"`
	Message string                                                                                                                 `json:"message,required"`
	JSON    alertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelopeErrorsJSON `json:"-"`
}

// alertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AlertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelopeErrors]
type alertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelopeMessages struct {
	Code    int64                                                                                                                    `json:"code,required"`
	Message string                                                                                                                   `json:"message,required"`
	JSON    alertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelopeMessagesJSON `json:"-"`
}

// alertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AlertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelopeMessages]
type alertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AlertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelopeSuccess bool

const (
	AlertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelopeSuccessTrue AlertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelopeSuccess = true
)

type AlertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                                    `json:"total_count"`
	JSON       alertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelopeResultInfoJSON `json:"-"`
}

// alertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [AlertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelopeResultInfo]
type alertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationEligibleNotificationMechanismEligibilityGetDeliveryMechanismEligibilityResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
