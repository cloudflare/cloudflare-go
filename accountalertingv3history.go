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

// AccountAlertingV3HistoryService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountAlertingV3HistoryService] method instead.
type AccountAlertingV3HistoryService struct {
	Options []option.RequestOption
}

// NewAccountAlertingV3HistoryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAlertingV3HistoryService(opts ...option.RequestOption) (r *AccountAlertingV3HistoryService) {
	r = &AccountAlertingV3HistoryService{}
	r.Options = opts
	return
}

// Gets a list of history records for notifications sent to an account. The records
// are displayed for last `x` number of days based on the zone plan (free = 30, pro
// = 30, biz = 30, ent = 90).
func (r *AccountAlertingV3HistoryService) NotificationHistoryListHistory(ctx context.Context, identifier string, query AccountAlertingV3HistoryNotificationHistoryListHistoryParams, opts ...option.RequestOption) (res *HistoryResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/history", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type HistoryResponseCollection struct {
	Errors     []HistoryResponseCollectionError   `json:"errors"`
	Messages   []HistoryResponseCollectionMessage `json:"messages"`
	Result     []HistoryResponseCollectionResult  `json:"result"`
	ResultInfo interface{}                        `json:"result_info"`
	// Whether the API call was successful
	Success HistoryResponseCollectionSuccess `json:"success"`
	JSON    historyResponseCollectionJSON    `json:"-"`
}

// historyResponseCollectionJSON contains the JSON metadata for the struct
// [HistoryResponseCollection]
type historyResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HistoryResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HistoryResponseCollectionError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    historyResponseCollectionErrorJSON `json:"-"`
}

// historyResponseCollectionErrorJSON contains the JSON metadata for the struct
// [HistoryResponseCollectionError]
type historyResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HistoryResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HistoryResponseCollectionMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    historyResponseCollectionMessageJSON `json:"-"`
}

// historyResponseCollectionMessageJSON contains the JSON metadata for the struct
// [HistoryResponseCollectionMessage]
type historyResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HistoryResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HistoryResponseCollectionResult struct {
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
	MechanismType HistoryResponseCollectionResultMechanismType `json:"mechanism_type"`
	// Name of the policy.
	Name string `json:"name"`
	// Timestamp of when the notification was dispatched in ISO 8601 format.
	Sent time.Time                           `json:"sent" format:"date-time"`
	JSON historyResponseCollectionResultJSON `json:"-"`
}

// historyResponseCollectionResultJSON contains the JSON metadata for the struct
// [HistoryResponseCollectionResult]
type historyResponseCollectionResultJSON struct {
	ID            apijson.Field
	AlertBody     apijson.Field
	AlertType     apijson.Field
	Description   apijson.Field
	Mechanism     apijson.Field
	MechanismType apijson.Field
	Name          apijson.Field
	Sent          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *HistoryResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of mechanism to which the notification has been dispatched. This can be
// email/pagerduty/webhook based on the mechanism configured.
type HistoryResponseCollectionResultMechanismType string

const (
	HistoryResponseCollectionResultMechanismTypeEmail     HistoryResponseCollectionResultMechanismType = "email"
	HistoryResponseCollectionResultMechanismTypePagerduty HistoryResponseCollectionResultMechanismType = "pagerduty"
	HistoryResponseCollectionResultMechanismTypeWebhook   HistoryResponseCollectionResultMechanismType = "webhook"
)

// Whether the API call was successful
type HistoryResponseCollectionSuccess bool

const (
	HistoryResponseCollectionSuccessTrue HistoryResponseCollectionSuccess = true
)

type AccountAlertingV3HistoryNotificationHistoryListHistoryParams struct {
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

// URLQuery serializes
// [AccountAlertingV3HistoryNotificationHistoryListHistoryParams]'s query
// parameters as `url.Values`.
func (r AccountAlertingV3HistoryNotificationHistoryListHistoryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
