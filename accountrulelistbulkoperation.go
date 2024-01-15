// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountRuleListBulkOperationService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountRuleListBulkOperationService] method instead.
type AccountRuleListBulkOperationService struct {
	Options []option.RequestOption
}

// NewAccountRuleListBulkOperationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountRuleListBulkOperationService(opts ...option.RequestOption) (r *AccountRuleListBulkOperationService) {
	r = &AccountRuleListBulkOperationService{}
	r.Options = opts
	return
}

// Gets the current status of an asynchronous operation on a list.
//
// The `status` property can have one of the following values: `pending`,
// `running`, `completed`, or `failed`. If the status is `failed`, the `error`
// property will contain a message describing the error.
func (r *AccountRuleListBulkOperationService) Get(ctx context.Context, accountIdentifier string, operationID string, opts ...option.RequestOption) (res *AccountRuleListBulkOperationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rules/lists/bulk_operations/%s", accountIdentifier, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountRuleListBulkOperationGetResponse struct {
	Errors   []AccountRuleListBulkOperationGetResponseError   `json:"errors"`
	Messages []AccountRuleListBulkOperationGetResponseMessage `json:"messages"`
	Result   AccountRuleListBulkOperationGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountRuleListBulkOperationGetResponseSuccess `json:"success"`
	JSON    accountRuleListBulkOperationGetResponseJSON    `json:"-"`
}

// accountRuleListBulkOperationGetResponseJSON contains the JSON metadata for the
// struct [AccountRuleListBulkOperationGetResponse]
type accountRuleListBulkOperationGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListBulkOperationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListBulkOperationGetResponseError struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accountRuleListBulkOperationGetResponseErrorJSON `json:"-"`
}

// accountRuleListBulkOperationGetResponseErrorJSON contains the JSON metadata for
// the struct [AccountRuleListBulkOperationGetResponseError]
type accountRuleListBulkOperationGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListBulkOperationGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListBulkOperationGetResponseMessage struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    accountRuleListBulkOperationGetResponseMessageJSON `json:"-"`
}

// accountRuleListBulkOperationGetResponseMessageJSON contains the JSON metadata
// for the struct [AccountRuleListBulkOperationGetResponseMessage]
type accountRuleListBulkOperationGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListBulkOperationGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRuleListBulkOperationGetResponseResult struct {
	// The unique operation ID of the asynchronous action.
	ID string `json:"id,required"`
	// The current status of the asynchronous operation.
	Status AccountRuleListBulkOperationGetResponseResultStatus `json:"status,required"`
	// The RFC 3339 timestamp of when the operation was completed.
	Completed string `json:"completed"`
	// A message describing the error when the status is `failed`.
	Error string                                            `json:"error"`
	JSON  accountRuleListBulkOperationGetResponseResultJSON `json:"-"`
}

// accountRuleListBulkOperationGetResponseResultJSON contains the JSON metadata for
// the struct [AccountRuleListBulkOperationGetResponseResult]
type accountRuleListBulkOperationGetResponseResultJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	Completed   apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListBulkOperationGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the asynchronous operation.
type AccountRuleListBulkOperationGetResponseResultStatus string

const (
	AccountRuleListBulkOperationGetResponseResultStatusPending   AccountRuleListBulkOperationGetResponseResultStatus = "pending"
	AccountRuleListBulkOperationGetResponseResultStatusRunning   AccountRuleListBulkOperationGetResponseResultStatus = "running"
	AccountRuleListBulkOperationGetResponseResultStatusCompleted AccountRuleListBulkOperationGetResponseResultStatus = "completed"
	AccountRuleListBulkOperationGetResponseResultStatusFailed    AccountRuleListBulkOperationGetResponseResultStatus = "failed"
)

// Whether the API call was successful
type AccountRuleListBulkOperationGetResponseSuccess bool

const (
	AccountRuleListBulkOperationGetResponseSuccessTrue AccountRuleListBulkOperationGetResponseSuccess = true
)
