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
func (r *AccountWorkerScriptUsageModelService) WorkerScriptFetchUsageModel(ctx context.Context, accountIdentifier string, scriptName string, opts ...option.RequestOption) (res *UsageModelResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/usage-model", accountIdentifier, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the Usage Model for a given Worker. Requires a Workers Paid
// subscription.
func (r *AccountWorkerScriptUsageModelService) WorkerScriptUpdateUsageModel(ctx context.Context, accountIdentifier string, scriptName string, body AccountWorkerScriptUsageModelWorkerScriptUpdateUsageModelParams, opts ...option.RequestOption) (res *UsageModelResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/usage-model", accountIdentifier, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type UsageModelResponse struct {
	Errors   []UsageModelResponseError   `json:"errors"`
	Messages []UsageModelResponseMessage `json:"messages"`
	Result   UsageModelResponseResult    `json:"result"`
	// Whether the API call was successful
	Success UsageModelResponseSuccess `json:"success"`
	JSON    usageModelResponseJSON    `json:"-"`
}

// usageModelResponseJSON contains the JSON metadata for the struct
// [UsageModelResponse]
type usageModelResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageModelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UsageModelResponseError struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    usageModelResponseErrorJSON `json:"-"`
}

// usageModelResponseErrorJSON contains the JSON metadata for the struct
// [UsageModelResponseError]
type usageModelResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageModelResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UsageModelResponseMessage struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    usageModelResponseMessageJSON `json:"-"`
}

// usageModelResponseMessageJSON contains the JSON metadata for the struct
// [UsageModelResponseMessage]
type usageModelResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageModelResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UsageModelResponseResult struct {
	UsageModel interface{}                  `json:"usage_model"`
	JSON       usageModelResponseResultJSON `json:"-"`
}

// usageModelResponseResultJSON contains the JSON metadata for the struct
// [UsageModelResponseResult]
type usageModelResponseResultJSON struct {
	UsageModel  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsageModelResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UsageModelResponseSuccess bool

const (
	UsageModelResponseSuccessTrue UsageModelResponseSuccess = true
)

type AccountWorkerScriptUsageModelWorkerScriptUpdateUsageModelParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountWorkerScriptUsageModelWorkerScriptUpdateUsageModelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
