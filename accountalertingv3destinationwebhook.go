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

// Get details for a single webhooks destination.
func (r *AccountAlertingV3DestinationWebhookService) Get(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *WebhooksSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a webhook destination.
func (r *AccountAlertingV3DestinationWebhookService) Update(ctx context.Context, identifier string, uuid string, body AccountAlertingV3DestinationWebhookUpdateParams, opts ...option.RequestOption) (res *WebhooksIDResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete a configured webhook destination.
func (r *AccountAlertingV3DestinationWebhookService) Delete(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *APIResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a new webhook destination.
func (r *AccountAlertingV3DestinationWebhookService) NotificationWebhooksNewAWebhook(ctx context.Context, identifier string, body AccountAlertingV3DestinationWebhookNotificationWebhooksNewAWebhookParams, opts ...option.RequestOption) (res *WebhooksIDResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets a list of all configured webhook destinations.
func (r *AccountAlertingV3DestinationWebhookService) NotificationWebhooksListWebhooks(ctx context.Context, identifier string, opts ...option.RequestOption) (res *WebhooksResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type WebhooksIDResponse struct {
	Errors   []WebhooksIDResponseError   `json:"errors"`
	Messages []WebhooksIDResponseMessage `json:"messages"`
	Result   WebhooksIDResponseResult    `json:"result"`
	// Whether the API call was successful
	Success WebhooksIDResponseSuccess `json:"success"`
	JSON    webhooksIDResponseJSON    `json:"-"`
}

// webhooksIDResponseJSON contains the JSON metadata for the struct
// [WebhooksIDResponse]
type webhooksIDResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhooksIDResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebhooksIDResponseError struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    webhooksIDResponseErrorJSON `json:"-"`
}

// webhooksIDResponseErrorJSON contains the JSON metadata for the struct
// [WebhooksIDResponseError]
type webhooksIDResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhooksIDResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebhooksIDResponseMessage struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    webhooksIDResponseMessageJSON `json:"-"`
}

// webhooksIDResponseMessageJSON contains the JSON metadata for the struct
// [WebhooksIDResponseMessage]
type webhooksIDResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhooksIDResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebhooksIDResponseResult struct {
	// UUID
	ID   string                       `json:"id"`
	JSON webhooksIDResponseResultJSON `json:"-"`
}

// webhooksIDResponseResultJSON contains the JSON metadata for the struct
// [WebhooksIDResponseResult]
type webhooksIDResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhooksIDResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WebhooksIDResponseSuccess bool

const (
	WebhooksIDResponseSuccessTrue WebhooksIDResponseSuccess = true
)

type WebhooksResponseCollection struct {
	Errors     []WebhooksResponseCollectionError    `json:"errors"`
	Messages   []WebhooksResponseCollectionMessage  `json:"messages"`
	Result     []WebhooksResponseCollectionResult   `json:"result"`
	ResultInfo WebhooksResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success WebhooksResponseCollectionSuccess `json:"success"`
	JSON    webhooksResponseCollectionJSON    `json:"-"`
}

// webhooksResponseCollectionJSON contains the JSON metadata for the struct
// [WebhooksResponseCollection]
type webhooksResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhooksResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebhooksResponseCollectionError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    webhooksResponseCollectionErrorJSON `json:"-"`
}

// webhooksResponseCollectionErrorJSON contains the JSON metadata for the struct
// [WebhooksResponseCollectionError]
type webhooksResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhooksResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebhooksResponseCollectionMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    webhooksResponseCollectionMessageJSON `json:"-"`
}

// webhooksResponseCollectionMessageJSON contains the JSON metadata for the struct
// [WebhooksResponseCollectionMessage]
type webhooksResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhooksResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebhooksResponseCollectionResult struct {
	// UUID
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
	// dispatching a webhook notification. Secrets are not returned in any API response
	// body.
	Secret string `json:"secret"`
	// Type of webhook endpoint.
	Type WebhooksResponseCollectionResultType `json:"type"`
	// The POST endpoint to call when dispatching a notification.
	URL  string                               `json:"url"`
	JSON webhooksResponseCollectionResultJSON `json:"-"`
}

// webhooksResponseCollectionResultJSON contains the JSON metadata for the struct
// [WebhooksResponseCollectionResult]
type webhooksResponseCollectionResultJSON struct {
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

func (r *WebhooksResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of webhook endpoint.
type WebhooksResponseCollectionResultType string

const (
	WebhooksResponseCollectionResultTypeSlack   WebhooksResponseCollectionResultType = "slack"
	WebhooksResponseCollectionResultTypeGeneric WebhooksResponseCollectionResultType = "generic"
	WebhooksResponseCollectionResultTypeGchat   WebhooksResponseCollectionResultType = "gchat"
)

type WebhooksResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       webhooksResponseCollectionResultInfoJSON `json:"-"`
}

// webhooksResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [WebhooksResponseCollectionResultInfo]
type webhooksResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhooksResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WebhooksResponseCollectionSuccess bool

const (
	WebhooksResponseCollectionSuccessTrue WebhooksResponseCollectionSuccess = true
)

type WebhooksSingleResponse struct {
	Errors   []WebhooksSingleResponseError   `json:"errors"`
	Messages []WebhooksSingleResponseMessage `json:"messages"`
	Result   WebhooksSingleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success WebhooksSingleResponseSuccess `json:"success"`
	JSON    webhooksSingleResponseJSON    `json:"-"`
}

// webhooksSingleResponseJSON contains the JSON metadata for the struct
// [WebhooksSingleResponse]
type webhooksSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhooksSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebhooksSingleResponseError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    webhooksSingleResponseErrorJSON `json:"-"`
}

// webhooksSingleResponseErrorJSON contains the JSON metadata for the struct
// [WebhooksSingleResponseError]
type webhooksSingleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhooksSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebhooksSingleResponseMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    webhooksSingleResponseMessageJSON `json:"-"`
}

// webhooksSingleResponseMessageJSON contains the JSON metadata for the struct
// [WebhooksSingleResponseMessage]
type webhooksSingleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhooksSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebhooksSingleResponseResult struct {
	// UUID
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
	// dispatching a webhook notification. Secrets are not returned in any API response
	// body.
	Secret string `json:"secret"`
	// Type of webhook endpoint.
	Type WebhooksSingleResponseResultType `json:"type"`
	// The POST endpoint to call when dispatching a notification.
	URL  string                           `json:"url"`
	JSON webhooksSingleResponseResultJSON `json:"-"`
}

// webhooksSingleResponseResultJSON contains the JSON metadata for the struct
// [WebhooksSingleResponseResult]
type webhooksSingleResponseResultJSON struct {
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

func (r *WebhooksSingleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of webhook endpoint.
type WebhooksSingleResponseResultType string

const (
	WebhooksSingleResponseResultTypeSlack   WebhooksSingleResponseResultType = "slack"
	WebhooksSingleResponseResultTypeGeneric WebhooksSingleResponseResultType = "generic"
	WebhooksSingleResponseResultTypeGchat   WebhooksSingleResponseResultType = "gchat"
)

// Whether the API call was successful
type WebhooksSingleResponseSuccess bool

const (
	WebhooksSingleResponseSuccessTrue WebhooksSingleResponseSuccess = true
)

type AccountAlertingV3DestinationWebhookUpdateParams struct {
	// The name of the webhook destination. This will be included in the request body
	// when you receive a webhook notification.
	Name param.Field[string] `json:"name,required"`
	// The POST endpoint to call when dispatching a notification.
	URL param.Field[string] `json:"url,required"`
	// Optional secret that will be passed in the `cf-webhook-auth` header when
	// dispatching a webhook notification. Secrets are not returned in any API response
	// body.
	Secret param.Field[string] `json:"secret"`
}

func (r AccountAlertingV3DestinationWebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAlertingV3DestinationWebhookNotificationWebhooksNewAWebhookParams struct {
	// The name of the webhook destination. This will be included in the request body
	// when you receive a webhook notification.
	Name param.Field[string] `json:"name,required"`
	// The POST endpoint to call when dispatching a notification.
	URL param.Field[string] `json:"url,required"`
	// Optional secret that will be passed in the `cf-webhook-auth` header when
	// dispatching a webhook notification. Secrets are not returned in any API response
	// body.
	Secret param.Field[string] `json:"secret"`
}

func (r AccountAlertingV3DestinationWebhookNotificationWebhooksNewAWebhookParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
