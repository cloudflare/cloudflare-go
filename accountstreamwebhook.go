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
func (r *AccountStreamWebhookService) Delete(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountStreamWebhookDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/webhook", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a webhook notification.
func (r *AccountStreamWebhookService) StreamWebhookNewWebhooks(ctx context.Context, accountIdentifier string, body AccountStreamWebhookStreamWebhookNewWebhooksParams, opts ...option.RequestOption) (res *AccountStreamWebhookStreamWebhookNewWebhooksResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/webhook", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieves a list of webhooks.
func (r *AccountStreamWebhookService) StreamWebhookViewWebhooks(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountStreamWebhookStreamWebhookViewWebhooksResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/webhook", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountStreamWebhookDeleteResponse struct {
	Errors   []AccountStreamWebhookDeleteResponseError   `json:"errors"`
	Messages []AccountStreamWebhookDeleteResponseMessage `json:"messages"`
	Result   string                                      `json:"result"`
	// Whether the API call was successful
	Success AccountStreamWebhookDeleteResponseSuccess `json:"success"`
	JSON    accountStreamWebhookDeleteResponseJSON    `json:"-"`
}

// accountStreamWebhookDeleteResponseJSON contains the JSON metadata for the struct
// [AccountStreamWebhookDeleteResponse]
type accountStreamWebhookDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamWebhookDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamWebhookDeleteResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountStreamWebhookDeleteResponseErrorJSON `json:"-"`
}

// accountStreamWebhookDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountStreamWebhookDeleteResponseError]
type accountStreamWebhookDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamWebhookDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamWebhookDeleteResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountStreamWebhookDeleteResponseMessageJSON `json:"-"`
}

// accountStreamWebhookDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountStreamWebhookDeleteResponseMessage]
type accountStreamWebhookDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamWebhookDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStreamWebhookDeleteResponseSuccess bool

const (
	AccountStreamWebhookDeleteResponseSuccessTrue AccountStreamWebhookDeleteResponseSuccess = true
)

type AccountStreamWebhookStreamWebhookNewWebhooksResponse struct {
	Errors   []AccountStreamWebhookStreamWebhookNewWebhooksResponseError   `json:"errors"`
	Messages []AccountStreamWebhookStreamWebhookNewWebhooksResponseMessage `json:"messages"`
	Result   interface{}                                                   `json:"result"`
	// Whether the API call was successful
	Success AccountStreamWebhookStreamWebhookNewWebhooksResponseSuccess `json:"success"`
	JSON    accountStreamWebhookStreamWebhookNewWebhooksResponseJSON    `json:"-"`
}

// accountStreamWebhookStreamWebhookNewWebhooksResponseJSON contains the JSON
// metadata for the struct [AccountStreamWebhookStreamWebhookNewWebhooksResponse]
type accountStreamWebhookStreamWebhookNewWebhooksResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamWebhookStreamWebhookNewWebhooksResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamWebhookStreamWebhookNewWebhooksResponseError struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    accountStreamWebhookStreamWebhookNewWebhooksResponseErrorJSON `json:"-"`
}

// accountStreamWebhookStreamWebhookNewWebhooksResponseErrorJSON contains the JSON
// metadata for the struct
// [AccountStreamWebhookStreamWebhookNewWebhooksResponseError]
type accountStreamWebhookStreamWebhookNewWebhooksResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamWebhookStreamWebhookNewWebhooksResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamWebhookStreamWebhookNewWebhooksResponseMessage struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    accountStreamWebhookStreamWebhookNewWebhooksResponseMessageJSON `json:"-"`
}

// accountStreamWebhookStreamWebhookNewWebhooksResponseMessageJSON contains the
// JSON metadata for the struct
// [AccountStreamWebhookStreamWebhookNewWebhooksResponseMessage]
type accountStreamWebhookStreamWebhookNewWebhooksResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamWebhookStreamWebhookNewWebhooksResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStreamWebhookStreamWebhookNewWebhooksResponseSuccess bool

const (
	AccountStreamWebhookStreamWebhookNewWebhooksResponseSuccessTrue AccountStreamWebhookStreamWebhookNewWebhooksResponseSuccess = true
)

type AccountStreamWebhookStreamWebhookViewWebhooksResponse struct {
	Errors   []AccountStreamWebhookStreamWebhookViewWebhooksResponseError   `json:"errors"`
	Messages []AccountStreamWebhookStreamWebhookViewWebhooksResponseMessage `json:"messages"`
	Result   interface{}                                                    `json:"result"`
	// Whether the API call was successful
	Success AccountStreamWebhookStreamWebhookViewWebhooksResponseSuccess `json:"success"`
	JSON    accountStreamWebhookStreamWebhookViewWebhooksResponseJSON    `json:"-"`
}

// accountStreamWebhookStreamWebhookViewWebhooksResponseJSON contains the JSON
// metadata for the struct [AccountStreamWebhookStreamWebhookViewWebhooksResponse]
type accountStreamWebhookStreamWebhookViewWebhooksResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamWebhookStreamWebhookViewWebhooksResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamWebhookStreamWebhookViewWebhooksResponseError struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    accountStreamWebhookStreamWebhookViewWebhooksResponseErrorJSON `json:"-"`
}

// accountStreamWebhookStreamWebhookViewWebhooksResponseErrorJSON contains the JSON
// metadata for the struct
// [AccountStreamWebhookStreamWebhookViewWebhooksResponseError]
type accountStreamWebhookStreamWebhookViewWebhooksResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamWebhookStreamWebhookViewWebhooksResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamWebhookStreamWebhookViewWebhooksResponseMessage struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    accountStreamWebhookStreamWebhookViewWebhooksResponseMessageJSON `json:"-"`
}

// accountStreamWebhookStreamWebhookViewWebhooksResponseMessageJSON contains the
// JSON metadata for the struct
// [AccountStreamWebhookStreamWebhookViewWebhooksResponseMessage]
type accountStreamWebhookStreamWebhookViewWebhooksResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamWebhookStreamWebhookViewWebhooksResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStreamWebhookStreamWebhookViewWebhooksResponseSuccess bool

const (
	AccountStreamWebhookStreamWebhookViewWebhooksResponseSuccessTrue AccountStreamWebhookStreamWebhookViewWebhooksResponseSuccess = true
)

type AccountStreamWebhookStreamWebhookNewWebhooksParams struct {
	// The URL where webhooks will be sent.
	NotificationURL param.Field[string] `json:"notificationUrl,required" format:"uri"`
}

func (r AccountStreamWebhookStreamWebhookNewWebhooksParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
