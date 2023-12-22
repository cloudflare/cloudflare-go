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

// AccountStreamWebhookService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountStreamWebhookService]
// method instead.
type AccountStreamWebhookService struct {
	Options []option.RequestOption
}

// NewAccountStreamWebhookService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountStreamWebhookService(opts ...option.RequestOption) (r *AccountStreamWebhookService) {
	r = &AccountStreamWebhookService{}
	r.Options = opts
	return
}

// Deletes a webhook.
func (r *AccountStreamWebhookService) Delete(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *DeletedResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/webhook", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a webhook notification.
func (r *AccountStreamWebhookService) StreamWebhookNewWebhooks(ctx context.Context, accountIdentifier string, body AccountStreamWebhookStreamWebhookNewWebhooksParams, opts ...option.RequestOption) (res *WebhookResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/webhook", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieves a list of webhooks.
func (r *AccountStreamWebhookService) StreamWebhookViewWebhooks(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *WebhookResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/webhook", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type WebhookResponseSingle struct {
	Errors   []WebhookResponseSingleError   `json:"errors"`
	Messages []WebhookResponseSingleMessage `json:"messages"`
	Result   interface{}                    `json:"result"`
	// Whether the API call was successful
	Success WebhookResponseSingleSuccess `json:"success"`
	JSON    webhookResponseSingleJSON    `json:"-"`
}

// webhookResponseSingleJSON contains the JSON metadata for the struct
// [WebhookResponseSingle]
type webhookResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookResponseSingleError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    webhookResponseSingleErrorJSON `json:"-"`
}

// webhookResponseSingleErrorJSON contains the JSON metadata for the struct
// [WebhookResponseSingleError]
type webhookResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookResponseSingleMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    webhookResponseSingleMessageJSON `json:"-"`
}

// webhookResponseSingleMessageJSON contains the JSON metadata for the struct
// [WebhookResponseSingleMessage]
type webhookResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WebhookResponseSingleSuccess bool

const (
	WebhookResponseSingleSuccessTrue WebhookResponseSingleSuccess = true
)

type AccountStreamWebhookStreamWebhookNewWebhooksParams struct {
	// The URL where webhooks will be sent.
	NotificationURL param.Field[string] `json:"notificationUrl,required" format:"uri"`
}

func (r AccountStreamWebhookStreamWebhookNewWebhooksParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
