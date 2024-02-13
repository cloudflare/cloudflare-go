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

// WorkerAccountSettingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWorkerAccountSettingService]
// method instead.
type WorkerAccountSettingService struct {
	Options []option.RequestOption
}

// NewWorkerAccountSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkerAccountSettingService(opts ...option.RequestOption) (r *WorkerAccountSettingService) {
	r = &WorkerAccountSettingService{}
	r.Options = opts
	return
}

// Creates Worker account settings for an account.
func (r *WorkerAccountSettingService) WorkerAccountSettingsNewWorkerAccountSettings(ctx context.Context, accountID string, body WorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsParams, opts ...option.RequestOption) (res *WorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/account-settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches Worker account settings for an account.
func (r *WorkerAccountSettingService) WorkerAccountSettingsFetchWorkerAccountSettings(ctx context.Context, accountID string, opts ...option.RequestOption) (res *WorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/account-settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponse struct {
	DefaultUsageModel interface{}                                                                   `json:"default_usage_model"`
	GreenCompute      interface{}                                                                   `json:"green_compute"`
	JSON              workerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseJSON `json:"-"`
}

// workerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseJSON
// contains the JSON metadata for the struct
// [WorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponse]
type workerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseJSON struct {
	DefaultUsageModel apijson.Field
	GreenCompute      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *WorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponse struct {
	DefaultUsageModel interface{}                                                                     `json:"default_usage_model"`
	GreenCompute      interface{}                                                                     `json:"green_compute"`
	JSON              workerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseJSON `json:"-"`
}

// workerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseJSON
// contains the JSON metadata for the struct
// [WorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponse]
type workerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseJSON struct {
	DefaultUsageModel apijson.Field
	GreenCompute      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *WorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r WorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type WorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseEnvelope struct {
	Errors   []WorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseEnvelopeJSON    `json:"-"`
}

// workerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [WorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseEnvelope]
type workerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseEnvelopeErrors struct {
	Code    int64                                                                                       `json:"code,required"`
	Message string                                                                                      `json:"message,required"`
	JSON    workerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseEnvelopeErrorsJSON `json:"-"`
}

// workerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [WorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseEnvelopeErrors]
type workerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseEnvelopeMessages struct {
	Code    int64                                                                                         `json:"code,required"`
	Message string                                                                                        `json:"message,required"`
	JSON    workerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseEnvelopeMessagesJSON `json:"-"`
}

// workerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [WorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseEnvelopeMessages]
type workerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseEnvelopeSuccess bool

const (
	WorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseEnvelopeSuccessTrue WorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseEnvelopeSuccess = true
)

type WorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseEnvelope struct {
	Errors   []WorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseEnvelopeJSON    `json:"-"`
}

// workerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [WorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseEnvelope]
type workerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseEnvelopeErrors struct {
	Code    int64                                                                                         `json:"code,required"`
	Message string                                                                                        `json:"message,required"`
	JSON    workerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseEnvelopeErrorsJSON `json:"-"`
}

// workerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [WorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseEnvelopeErrors]
type workerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseEnvelopeMessages struct {
	Code    int64                                                                                           `json:"code,required"`
	Message string                                                                                          `json:"message,required"`
	JSON    workerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseEnvelopeMessagesJSON `json:"-"`
}

// workerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [WorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseEnvelopeMessages]
type workerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseEnvelopeSuccess bool

const (
	WorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseEnvelopeSuccessTrue WorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseEnvelopeSuccess = true
)
