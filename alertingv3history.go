// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AlertingV3HistoryService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAlertingV3HistoryService] method
// instead.
type AlertingV3HistoryService struct {
	Options []option.RequestOption
}

// NewAlertingV3HistoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAlertingV3HistoryService(opts ...option.RequestOption) (r *AlertingV3HistoryService) {
	r = &AlertingV3HistoryService{}
	r.Options = opts
	return
}

// Gets a list of history records for notifications sent to an account. The records
// are displayed for last `x` number of days based on the zone plan (free = 30, pro
// = 30, biz = 30, ent = 90).
func (r *AlertingV3HistoryService) NotificationHistoryListHistory(ctx context.Context, accountID string, query AlertingV3HistoryNotificationHistoryListHistoryParams, opts ...option.RequestOption) (res *[]AlertingV3HistoryNotificationHistoryListHistoryResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AlertingV3HistoryNotificationHistoryListHistoryResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/history", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AlertingV3HistoryNotificationHistoryListHistoryResponse struct {
	// UUID
	ID string `json:"id"`
	// Message body included in the notification sent.
	AlertBody string `json:"alert_body"`
	// Type of notification that has been dispatched.
	AlertType string `json:"alert_type"`
	// Description of the notification policy (if present).
	Description string `json:"description"`
	// The mechanism to which the notification has been dispatched.
	Mechanism string `json:"mechanism"`
	// The type of mechanism to which the notification has been dispatched. This can be
	// email/pagerduty/webhook based on the mechanism configured.
	MechanismType AlertingV3HistoryNotificationHistoryListHistoryResponseMechanismType `json:"mechanism_type"`
	// Name of the policy.
	Name string `json:"name"`
	// The unique identifier of a notification policy
	PolicyID string `json:"policy_id"`
	// Timestamp of when the notification was dispatched in ISO 8601 format.
	Sent time.Time                                                   `json:"sent" format:"date-time"`
	JSON alertingV3HistoryNotificationHistoryListHistoryResponseJSON `json:"-"`
}

// alertingV3HistoryNotificationHistoryListHistoryResponseJSON contains the JSON
// metadata for the struct
// [AlertingV3HistoryNotificationHistoryListHistoryResponse]
type alertingV3HistoryNotificationHistoryListHistoryResponseJSON struct {
	ID            apijson.Field
	AlertBody     apijson.Field
	AlertType     apijson.Field
	Description   apijson.Field
	Mechanism     apijson.Field
	MechanismType apijson.Field
	Name          apijson.Field
	PolicyID      apijson.Field
	Sent          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AlertingV3HistoryNotificationHistoryListHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of mechanism to which the notification has been dispatched. This can be
// email/pagerduty/webhook based on the mechanism configured.
type AlertingV3HistoryNotificationHistoryListHistoryResponseMechanismType string

const (
	AlertingV3HistoryNotificationHistoryListHistoryResponseMechanismTypeEmail     AlertingV3HistoryNotificationHistoryListHistoryResponseMechanismType = "email"
	AlertingV3HistoryNotificationHistoryListHistoryResponseMechanismTypePagerduty AlertingV3HistoryNotificationHistoryListHistoryResponseMechanismType = "pagerduty"
	AlertingV3HistoryNotificationHistoryListHistoryResponseMechanismTypeWebhook   AlertingV3HistoryNotificationHistoryListHistoryResponseMechanismType = "webhook"
)

type AlertingV3HistoryNotificationHistoryListHistoryParams struct {
	// Limit the returned results to history records older than the specified date.
	// This must be a timestamp that conforms to RFC3339.
	Before param.Field[time.Time] `query:"before" format:"date-time"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Limit the returned results to history records newer than the specified date.
	// This must be a timestamp that conforms to RFC3339.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
}

// URLQuery serializes [AlertingV3HistoryNotificationHistoryListHistoryParams]'s
// query parameters as `url.Values`.
func (r AlertingV3HistoryNotificationHistoryListHistoryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AlertingV3HistoryNotificationHistoryListHistoryResponseEnvelope struct {
	Errors   []AlertingV3HistoryNotificationHistoryListHistoryResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AlertingV3HistoryNotificationHistoryListHistoryResponseEnvelopeMessages `json:"messages,required"`
	Result   []AlertingV3HistoryNotificationHistoryListHistoryResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AlertingV3HistoryNotificationHistoryListHistoryResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AlertingV3HistoryNotificationHistoryListHistoryResponseEnvelopeResultInfo `json:"result_info"`
	JSON       alertingV3HistoryNotificationHistoryListHistoryResponseEnvelopeJSON       `json:"-"`
}

// alertingV3HistoryNotificationHistoryListHistoryResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [AlertingV3HistoryNotificationHistoryListHistoryResponseEnvelope]
type alertingV3HistoryNotificationHistoryListHistoryResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3HistoryNotificationHistoryListHistoryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3HistoryNotificationHistoryListHistoryResponseEnvelopeErrors struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    alertingV3HistoryNotificationHistoryListHistoryResponseEnvelopeErrorsJSON `json:"-"`
}

// alertingV3HistoryNotificationHistoryListHistoryResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AlertingV3HistoryNotificationHistoryListHistoryResponseEnvelopeErrors]
type alertingV3HistoryNotificationHistoryListHistoryResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3HistoryNotificationHistoryListHistoryResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3HistoryNotificationHistoryListHistoryResponseEnvelopeMessages struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    alertingV3HistoryNotificationHistoryListHistoryResponseEnvelopeMessagesJSON `json:"-"`
}

// alertingV3HistoryNotificationHistoryListHistoryResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AlertingV3HistoryNotificationHistoryListHistoryResponseEnvelopeMessages]
type alertingV3HistoryNotificationHistoryListHistoryResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3HistoryNotificationHistoryListHistoryResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AlertingV3HistoryNotificationHistoryListHistoryResponseEnvelopeSuccess bool

const (
	AlertingV3HistoryNotificationHistoryListHistoryResponseEnvelopeSuccessTrue AlertingV3HistoryNotificationHistoryListHistoryResponseEnvelopeSuccess = true
)

type AlertingV3HistoryNotificationHistoryListHistoryResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                       `json:"total_count"`
	JSON       alertingV3HistoryNotificationHistoryListHistoryResponseEnvelopeResultInfoJSON `json:"-"`
}

// alertingV3HistoryNotificationHistoryListHistoryResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [AlertingV3HistoryNotificationHistoryListHistoryResponseEnvelopeResultInfo]
type alertingV3HistoryNotificationHistoryListHistoryResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3HistoryNotificationHistoryListHistoryResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
