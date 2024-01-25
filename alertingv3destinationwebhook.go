// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AlertingV3DestinationWebhookService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAlertingV3DestinationWebhookService] method instead.
type AlertingV3DestinationWebhookService struct {
	Options []option.RequestOption
}

// NewAlertingV3DestinationWebhookService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAlertingV3DestinationWebhookService(opts ...option.RequestOption) (r *AlertingV3DestinationWebhookService) {
	r = &AlertingV3DestinationWebhookService{}
	r.Options = opts
	return
}

// Update a webhook destination.
func (r *AlertingV3DestinationWebhookService) Update(ctx context.Context, accountID string, webhookID string, body AlertingV3DestinationWebhookUpdateParams, opts ...option.RequestOption) (res *AlertingV3DestinationWebhookUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks/%s", accountID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AlertingV3DestinationWebhookUpdateResponse struct {
	Errors   []AlertingV3DestinationWebhookUpdateResponseError   `json:"errors"`
	Messages []AlertingV3DestinationWebhookUpdateResponseMessage `json:"messages"`
	Result   AlertingV3DestinationWebhookUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AlertingV3DestinationWebhookUpdateResponseSuccess `json:"success"`
	JSON    alertingV3DestinationWebhookUpdateResponseJSON    `json:"-"`
}

// alertingV3DestinationWebhookUpdateResponseJSON contains the JSON metadata for
// the struct [AlertingV3DestinationWebhookUpdateResponse]
type alertingV3DestinationWebhookUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationWebhookUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationWebhookUpdateResponseError struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    alertingV3DestinationWebhookUpdateResponseErrorJSON `json:"-"`
}

// alertingV3DestinationWebhookUpdateResponseErrorJSON contains the JSON metadata
// for the struct [AlertingV3DestinationWebhookUpdateResponseError]
type alertingV3DestinationWebhookUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationWebhookUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationWebhookUpdateResponseMessage struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    alertingV3DestinationWebhookUpdateResponseMessageJSON `json:"-"`
}

// alertingV3DestinationWebhookUpdateResponseMessageJSON contains the JSON metadata
// for the struct [AlertingV3DestinationWebhookUpdateResponseMessage]
type alertingV3DestinationWebhookUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationWebhookUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationWebhookUpdateResponseResult struct {
	// UUID
	ID   string                                               `json:"id"`
	JSON alertingV3DestinationWebhookUpdateResponseResultJSON `json:"-"`
}

// alertingV3DestinationWebhookUpdateResponseResultJSON contains the JSON metadata
// for the struct [AlertingV3DestinationWebhookUpdateResponseResult]
type alertingV3DestinationWebhookUpdateResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationWebhookUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AlertingV3DestinationWebhookUpdateResponseSuccess bool

const (
	AlertingV3DestinationWebhookUpdateResponseSuccessTrue AlertingV3DestinationWebhookUpdateResponseSuccess = true
)

type AlertingV3DestinationWebhookUpdateParams struct {
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

func (r AlertingV3DestinationWebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
