// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountAlertingV3DestinationWebhookService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountAlertingV3DestinationWebhookService] method instead.
type AccountAlertingV3DestinationWebhookService struct {
	Options []option.RequestOption
}

// NewAccountAlertingV3DestinationWebhookService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountAlertingV3DestinationWebhookService(opts ...option.RequestOption) (r *AccountAlertingV3DestinationWebhookService) {
	r = &AccountAlertingV3DestinationWebhookService{}
	r.Options = opts
	return
}

// Creates a new webhook destination.
func (r *AccountAlertingV3DestinationWebhookService) New(ctx context.Context, accountID string, body AccountAlertingV3DestinationWebhookNewParams, opts ...option.RequestOption) (res *AccountAlertingV3DestinationWebhookNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get details for a single webhooks destination.
func (r *AccountAlertingV3DestinationWebhookService) Get(ctx context.Context, accountID string, webhookID string, opts ...option.RequestOption) (res *AccountAlertingV3DestinationWebhookGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks/%s", accountID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Gets a list of all configured webhook destinations.
func (r *AccountAlertingV3DestinationWebhookService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountAlertingV3DestinationWebhookListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a configured webhook destination.
func (r *AccountAlertingV3DestinationWebhookService) Delete(ctx context.Context, accountID string, webhookID string, opts ...option.RequestOption) (res *AccountAlertingV3DestinationWebhookDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks/%s", accountID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AccountAlertingV3DestinationWebhookNewResponse struct {
	Errors   []AccountAlertingV3DestinationWebhookNewResponseError   `json:"errors"`
	Messages []AccountAlertingV3DestinationWebhookNewResponseMessage `json:"messages"`
	Result   AccountAlertingV3DestinationWebhookNewResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAlertingV3DestinationWebhookNewResponseSuccess `json:"success"`
	JSON    accountAlertingV3DestinationWebhookNewResponseJSON    `json:"-"`
}

// accountAlertingV3DestinationWebhookNewResponseJSON contains the JSON metadata
// for the struct [AccountAlertingV3DestinationWebhookNewResponse]
type accountAlertingV3DestinationWebhookNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationWebhookNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAlertingV3DestinationWebhookNewResponseError struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    accountAlertingV3DestinationWebhookNewResponseErrorJSON `json:"-"`
}

// accountAlertingV3DestinationWebhookNewResponseErrorJSON contains the JSON
// metadata for the struct [AccountAlertingV3DestinationWebhookNewResponseError]
type accountAlertingV3DestinationWebhookNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationWebhookNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAlertingV3DestinationWebhookNewResponseMessage struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    accountAlertingV3DestinationWebhookNewResponseMessageJSON `json:"-"`
}

// accountAlertingV3DestinationWebhookNewResponseMessageJSON contains the JSON
// metadata for the struct [AccountAlertingV3DestinationWebhookNewResponseMessage]
type accountAlertingV3DestinationWebhookNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationWebhookNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAlertingV3DestinationWebhookNewResponseResult struct {
	// UUID
	ID   string                                                   `json:"id"`
	JSON accountAlertingV3DestinationWebhookNewResponseResultJSON `json:"-"`
}

// accountAlertingV3DestinationWebhookNewResponseResultJSON contains the JSON
// metadata for the struct [AccountAlertingV3DestinationWebhookNewResponseResult]
type accountAlertingV3DestinationWebhookNewResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationWebhookNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAlertingV3DestinationWebhookNewResponseSuccess bool

const (
	AccountAlertingV3DestinationWebhookNewResponseSuccessTrue AccountAlertingV3DestinationWebhookNewResponseSuccess = true
)

type AccountAlertingV3DestinationWebhookGetResponse struct {
	Errors   []AccountAlertingV3DestinationWebhookGetResponseError   `json:"errors"`
	Messages []AccountAlertingV3DestinationWebhookGetResponseMessage `json:"messages"`
	Result   AccountAlertingV3DestinationWebhookGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAlertingV3DestinationWebhookGetResponseSuccess `json:"success"`
	JSON    accountAlertingV3DestinationWebhookGetResponseJSON    `json:"-"`
}

// accountAlertingV3DestinationWebhookGetResponseJSON contains the JSON metadata
// for the struct [AccountAlertingV3DestinationWebhookGetResponse]
type accountAlertingV3DestinationWebhookGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationWebhookGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAlertingV3DestinationWebhookGetResponseError struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    accountAlertingV3DestinationWebhookGetResponseErrorJSON `json:"-"`
}

// accountAlertingV3DestinationWebhookGetResponseErrorJSON contains the JSON
// metadata for the struct [AccountAlertingV3DestinationWebhookGetResponseError]
type accountAlertingV3DestinationWebhookGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationWebhookGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAlertingV3DestinationWebhookGetResponseMessage struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    accountAlertingV3DestinationWebhookGetResponseMessageJSON `json:"-"`
}

// accountAlertingV3DestinationWebhookGetResponseMessageJSON contains the JSON
// metadata for the struct [AccountAlertingV3DestinationWebhookGetResponseMessage]
type accountAlertingV3DestinationWebhookGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationWebhookGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAlertingV3DestinationWebhookGetResponseResult struct {
	// The unique identifier of a webhook
	ID string `json:"id"`
	// Timestamp of when the webhook destination was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of the last time an attempt to dispatch a notification to this webhook
	// failed.
	LastFailure time.Time `json:"last_failure" format:"date-time"`
	// Timestamp of the last time Cloudflare was able to successfully dispatch a
	// notification using this webhook.
	LastSuccess time.Time `json:"last_success" format:"date-time"`
	// The name of the webhook destination. This will be included in the request body
	// when you receive a webhook notification.
	Name string `json:"name"`
	// Optional secret that will be passed in the `cf-webhook-auth` header when
	// dispatching generic webhook notifications or formatted for supported
	// destinations. Secrets are not returned in any API response body.
	Secret string `json:"secret"`
	// Type of webhook endpoint.
	Type AccountAlertingV3DestinationWebhookGetResponseResultType `json:"type"`
	// The POST endpoint to call when dispatching a notification.
	URL  string                                                   `json:"url"`
	JSON accountAlertingV3DestinationWebhookGetResponseResultJSON `json:"-"`
}

// accountAlertingV3DestinationWebhookGetResponseResultJSON contains the JSON
// metadata for the struct [AccountAlertingV3DestinationWebhookGetResponseResult]
type accountAlertingV3DestinationWebhookGetResponseResultJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	LastFailure apijson.Field
	LastSuccess apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationWebhookGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of webhook endpoint.
type AccountAlertingV3DestinationWebhookGetResponseResultType string

const (
	AccountAlertingV3DestinationWebhookGetResponseResultTypeSlack   AccountAlertingV3DestinationWebhookGetResponseResultType = "slack"
	AccountAlertingV3DestinationWebhookGetResponseResultTypeGeneric AccountAlertingV3DestinationWebhookGetResponseResultType = "generic"
	AccountAlertingV3DestinationWebhookGetResponseResultTypeGchat   AccountAlertingV3DestinationWebhookGetResponseResultType = "gchat"
)

// Whether the API call was successful
type AccountAlertingV3DestinationWebhookGetResponseSuccess bool

const (
	AccountAlertingV3DestinationWebhookGetResponseSuccessTrue AccountAlertingV3DestinationWebhookGetResponseSuccess = true
)

type AccountAlertingV3DestinationWebhookListResponse struct {
	Errors     []AccountAlertingV3DestinationWebhookListResponseError    `json:"errors"`
	Messages   []AccountAlertingV3DestinationWebhookListResponseMessage  `json:"messages"`
	Result     []AccountAlertingV3DestinationWebhookListResponseResult   `json:"result"`
	ResultInfo AccountAlertingV3DestinationWebhookListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountAlertingV3DestinationWebhookListResponseSuccess `json:"success"`
	JSON    accountAlertingV3DestinationWebhookListResponseJSON    `json:"-"`
}

// accountAlertingV3DestinationWebhookListResponseJSON contains the JSON metadata
// for the struct [AccountAlertingV3DestinationWebhookListResponse]
type accountAlertingV3DestinationWebhookListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationWebhookListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAlertingV3DestinationWebhookListResponseError struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    accountAlertingV3DestinationWebhookListResponseErrorJSON `json:"-"`
}

// accountAlertingV3DestinationWebhookListResponseErrorJSON contains the JSON
// metadata for the struct [AccountAlertingV3DestinationWebhookListResponseError]
type accountAlertingV3DestinationWebhookListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationWebhookListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAlertingV3DestinationWebhookListResponseMessage struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    accountAlertingV3DestinationWebhookListResponseMessageJSON `json:"-"`
}

// accountAlertingV3DestinationWebhookListResponseMessageJSON contains the JSON
// metadata for the struct [AccountAlertingV3DestinationWebhookListResponseMessage]
type accountAlertingV3DestinationWebhookListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationWebhookListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAlertingV3DestinationWebhookListResponseResult struct {
	// The unique identifier of a webhook
	ID string `json:"id"`
	// Timestamp of when the webhook destination was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of the last time an attempt to dispatch a notification to this webhook
	// failed.
	LastFailure time.Time `json:"last_failure" format:"date-time"`
	// Timestamp of the last time Cloudflare was able to successfully dispatch a
	// notification using this webhook.
	LastSuccess time.Time `json:"last_success" format:"date-time"`
	// The name of the webhook destination. This will be included in the request body
	// when you receive a webhook notification.
	Name string `json:"name"`
	// Optional secret that will be passed in the `cf-webhook-auth` header when
	// dispatching generic webhook notifications or formatted for supported
	// destinations. Secrets are not returned in any API response body.
	Secret string `json:"secret"`
	// Type of webhook endpoint.
	Type AccountAlertingV3DestinationWebhookListResponseResultType `json:"type"`
	// The POST endpoint to call when dispatching a notification.
	URL  string                                                    `json:"url"`
	JSON accountAlertingV3DestinationWebhookListResponseResultJSON `json:"-"`
}

// accountAlertingV3DestinationWebhookListResponseResultJSON contains the JSON
// metadata for the struct [AccountAlertingV3DestinationWebhookListResponseResult]
type accountAlertingV3DestinationWebhookListResponseResultJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	LastFailure apijson.Field
	LastSuccess apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationWebhookListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of webhook endpoint.
type AccountAlertingV3DestinationWebhookListResponseResultType string

const (
	AccountAlertingV3DestinationWebhookListResponseResultTypeSlack   AccountAlertingV3DestinationWebhookListResponseResultType = "slack"
	AccountAlertingV3DestinationWebhookListResponseResultTypeGeneric AccountAlertingV3DestinationWebhookListResponseResultType = "generic"
	AccountAlertingV3DestinationWebhookListResponseResultTypeGchat   AccountAlertingV3DestinationWebhookListResponseResultType = "gchat"
)

type AccountAlertingV3DestinationWebhookListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                       `json:"total_count"`
	JSON       accountAlertingV3DestinationWebhookListResponseResultInfoJSON `json:"-"`
}

// accountAlertingV3DestinationWebhookListResponseResultInfoJSON contains the JSON
// metadata for the struct
// [AccountAlertingV3DestinationWebhookListResponseResultInfo]
type accountAlertingV3DestinationWebhookListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationWebhookListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAlertingV3DestinationWebhookListResponseSuccess bool

const (
	AccountAlertingV3DestinationWebhookListResponseSuccessTrue AccountAlertingV3DestinationWebhookListResponseSuccess = true
)

type AccountAlertingV3DestinationWebhookDeleteResponse struct {
	Errors     []AccountAlertingV3DestinationWebhookDeleteResponseError    `json:"errors"`
	Messages   []AccountAlertingV3DestinationWebhookDeleteResponseMessage  `json:"messages"`
	Result     []interface{}                                               `json:"result,nullable"`
	ResultInfo AccountAlertingV3DestinationWebhookDeleteResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountAlertingV3DestinationWebhookDeleteResponseSuccess `json:"success"`
	JSON    accountAlertingV3DestinationWebhookDeleteResponseJSON    `json:"-"`
}

// accountAlertingV3DestinationWebhookDeleteResponseJSON contains the JSON metadata
// for the struct [AccountAlertingV3DestinationWebhookDeleteResponse]
type accountAlertingV3DestinationWebhookDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationWebhookDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAlertingV3DestinationWebhookDeleteResponseError struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    accountAlertingV3DestinationWebhookDeleteResponseErrorJSON `json:"-"`
}

// accountAlertingV3DestinationWebhookDeleteResponseErrorJSON contains the JSON
// metadata for the struct [AccountAlertingV3DestinationWebhookDeleteResponseError]
type accountAlertingV3DestinationWebhookDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationWebhookDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAlertingV3DestinationWebhookDeleteResponseMessage struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    accountAlertingV3DestinationWebhookDeleteResponseMessageJSON `json:"-"`
}

// accountAlertingV3DestinationWebhookDeleteResponseMessageJSON contains the JSON
// metadata for the struct
// [AccountAlertingV3DestinationWebhookDeleteResponseMessage]
type accountAlertingV3DestinationWebhookDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationWebhookDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAlertingV3DestinationWebhookDeleteResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                         `json:"total_count"`
	JSON       accountAlertingV3DestinationWebhookDeleteResponseResultInfoJSON `json:"-"`
}

// accountAlertingV3DestinationWebhookDeleteResponseResultInfoJSON contains the
// JSON metadata for the struct
// [AccountAlertingV3DestinationWebhookDeleteResponseResultInfo]
type accountAlertingV3DestinationWebhookDeleteResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationWebhookDeleteResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAlertingV3DestinationWebhookDeleteResponseSuccess bool

const (
	AccountAlertingV3DestinationWebhookDeleteResponseSuccessTrue AccountAlertingV3DestinationWebhookDeleteResponseSuccess = true
)

type AccountAlertingV3DestinationWebhookNewParams struct {
	// The name of the webhook destination. This will be included in the request body
	// when you receive a webhook notification.
	Name param.Field[string] `json:"name,required"`
	// The POST endpoint to call when dispatching a notification.
	URL param.Field[string] `json:"url,required"`
	// Optional secret that will be passed in the `cf-webhook-auth` header when
	// dispatching generic webhook notifications or formatted for supported
	// destinations. Secrets are not returned in any API response body.
	Secret param.Field[string] `json:"secret"`
}

func (r AccountAlertingV3DestinationWebhookNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
