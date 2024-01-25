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

// AccountWorkerScriptUsageModelService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountWorkerScriptUsageModelService] method instead.
type AccountWorkerScriptUsageModelService struct {
	Options []option.RequestOption
}

// NewAccountWorkerScriptUsageModelService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountWorkerScriptUsageModelService(opts ...option.RequestOption) (r *AccountWorkerScriptUsageModelService) {
	r = &AccountWorkerScriptUsageModelService{}
	r.Options = opts
	return
}

// Fetches the Usage Model for a given Worker.
func (r *AccountWorkerScriptUsageModelService) WorkerScriptFetchUsageModel(ctx context.Context, accountIdentifier string, scriptName string, opts ...option.RequestOption) (res *AccountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/usage-model", accountIdentifier, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the Usage Model for a given Worker. Requires a Workers Paid
// subscription.
func (r *AccountWorkerScriptUsageModelService) WorkerScriptUpdateUsageModel(ctx context.Context, accountIdentifier string, scriptName string, body AccountWorkerScriptUsageModelWorkerScriptUpdateUsageModelParams, opts ...option.RequestOption) (res *AccountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/usage-model", accountIdentifier, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponse struct {
	Errors   []AccountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponseError   `json:"errors"`
	Messages []AccountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponseMessage `json:"messages"`
	Result   AccountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponseSuccess `json:"success"`
	JSON    accountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponseJSON    `json:"-"`
}

// accountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponseJSON contains
// the JSON metadata for the struct
// [AccountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponse]
type accountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponseError struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    accountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponseErrorJSON `json:"-"`
}

// accountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponseError]
type accountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponseMessage struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    accountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponseMessageJSON `json:"-"`
}

// accountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponseMessage]
type accountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponseResult struct {
	UsageModel interface{}                                                                `json:"usage_model"`
	JSON       accountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponseResultJSON `json:"-"`
}

// accountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponseResultJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponseResult]
type accountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponseResultJSON struct {
	UsageModel  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponseSuccess bool

const (
	AccountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponseSuccessTrue AccountWorkerScriptUsageModelWorkerScriptFetchUsageModelResponseSuccess = true
)

type AccountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponse struct {
	Errors   []AccountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseError   `json:"errors"`
	Messages []AccountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseMessage `json:"messages"`
	Result   AccountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseSuccess `json:"success"`
	JSON    accountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseJSON    `json:"-"`
}

// accountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseJSON contains
// the JSON metadata for the struct
// [AccountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponse]
type accountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseError struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    accountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseErrorJSON `json:"-"`
}

// accountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseError]
type accountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseMessage struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    accountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseMessageJSON `json:"-"`
}

// accountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseMessage]
type accountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseResult struct {
	UsageModel interface{}                                                                 `json:"usage_model"`
	JSON       accountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseResultJSON `json:"-"`
}

// accountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseResultJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseResult]
type accountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseResultJSON struct {
	UsageModel  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseSuccess bool

const (
	AccountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseSuccessTrue AccountWorkerScriptUsageModelWorkerScriptUpdateUsageModelResponseSuccess = true
)

type AccountWorkerScriptUsageModelWorkerScriptUpdateUsageModelParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountWorkerScriptUsageModelWorkerScriptUpdateUsageModelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
