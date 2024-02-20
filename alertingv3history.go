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
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
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
func (r *AlertingV3HistoryService) List(ctx context.Context, accountID string, query AlertingV3HistoryListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[AlertingV3HistoryListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/history", accountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Gets a list of history records for notifications sent to an account. The records
// are displayed for last `x` number of days based on the zone plan (free = 30, pro
// = 30, biz = 30, ent = 90).
func (r *AlertingV3HistoryService) ListAutoPaging(ctx context.Context, accountID string, query AlertingV3HistoryListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[AlertingV3HistoryListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, accountID, query, opts...))
}

type AlertingV3HistoryListResponse struct {
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
	MechanismType AlertingV3HistoryListResponseMechanismType `json:"mechanism_type"`
	// Name of the policy.
	Name string `json:"name"`
	// The unique identifier of a notification policy
	PolicyID string `json:"policy_id"`
	// Timestamp of when the notification was dispatched in ISO 8601 format.
	Sent time.Time                         `json:"sent" format:"date-time"`
	JSON alertingV3HistoryListResponseJSON `json:"-"`
}

// alertingV3HistoryListResponseJSON contains the JSON metadata for the struct
// [AlertingV3HistoryListResponse]
type alertingV3HistoryListResponseJSON struct {
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

func (r *AlertingV3HistoryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of mechanism to which the notification has been dispatched. This can be
// email/pagerduty/webhook based on the mechanism configured.
type AlertingV3HistoryListResponseMechanismType string

const (
	AlertingV3HistoryListResponseMechanismTypeEmail     AlertingV3HistoryListResponseMechanismType = "email"
	AlertingV3HistoryListResponseMechanismTypePagerduty AlertingV3HistoryListResponseMechanismType = "pagerduty"
	AlertingV3HistoryListResponseMechanismTypeWebhook   AlertingV3HistoryListResponseMechanismType = "webhook"
)

type AlertingV3HistoryListParams struct {
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

// URLQuery serializes [AlertingV3HistoryListParams]'s query parameters as
// `url.Values`.
func (r AlertingV3HistoryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
