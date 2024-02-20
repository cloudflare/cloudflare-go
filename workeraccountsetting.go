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

// Fetches Worker account settings for an account.
func (r *WorkerAccountSettingService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *WorkerAccountSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerAccountSettingGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/account-settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Creates Worker account settings for an account.
func (r *WorkerAccountSettingService) Replace(ctx context.Context, accountID string, body WorkerAccountSettingReplaceParams, opts ...option.RequestOption) (res *WorkerAccountSettingReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerAccountSettingReplaceResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/account-settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkerAccountSettingGetResponse struct {
	DefaultUsageModel interface{}                         `json:"default_usage_model"`
	GreenCompute      interface{}                         `json:"green_compute"`
	JSON              workerAccountSettingGetResponseJSON `json:"-"`
}

// workerAccountSettingGetResponseJSON contains the JSON metadata for the struct
// [WorkerAccountSettingGetResponse]
type workerAccountSettingGetResponseJSON struct {
	DefaultUsageModel apijson.Field
	GreenCompute      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *WorkerAccountSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerAccountSettingReplaceResponse struct {
	DefaultUsageModel interface{}                             `json:"default_usage_model"`
	GreenCompute      interface{}                             `json:"green_compute"`
	JSON              workerAccountSettingReplaceResponseJSON `json:"-"`
}

// workerAccountSettingReplaceResponseJSON contains the JSON metadata for the
// struct [WorkerAccountSettingReplaceResponse]
type workerAccountSettingReplaceResponseJSON struct {
	DefaultUsageModel apijson.Field
	GreenCompute      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *WorkerAccountSettingReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerAccountSettingGetResponseEnvelope struct {
	Errors   []WorkerAccountSettingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerAccountSettingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerAccountSettingGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerAccountSettingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerAccountSettingGetResponseEnvelopeJSON    `json:"-"`
}

// workerAccountSettingGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [WorkerAccountSettingGetResponseEnvelope]
type workerAccountSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerAccountSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerAccountSettingGetResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    workerAccountSettingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// workerAccountSettingGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [WorkerAccountSettingGetResponseEnvelopeErrors]
type workerAccountSettingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerAccountSettingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerAccountSettingGetResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    workerAccountSettingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// workerAccountSettingGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [WorkerAccountSettingGetResponseEnvelopeMessages]
type workerAccountSettingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerAccountSettingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerAccountSettingGetResponseEnvelopeSuccess bool

const (
	WorkerAccountSettingGetResponseEnvelopeSuccessTrue WorkerAccountSettingGetResponseEnvelopeSuccess = true
)

type WorkerAccountSettingReplaceParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r WorkerAccountSettingReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type WorkerAccountSettingReplaceResponseEnvelope struct {
	Errors   []WorkerAccountSettingReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerAccountSettingReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerAccountSettingReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerAccountSettingReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerAccountSettingReplaceResponseEnvelopeJSON    `json:"-"`
}

// workerAccountSettingReplaceResponseEnvelopeJSON contains the JSON metadata for
// the struct [WorkerAccountSettingReplaceResponseEnvelope]
type workerAccountSettingReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerAccountSettingReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerAccountSettingReplaceResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    workerAccountSettingReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// workerAccountSettingReplaceResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [WorkerAccountSettingReplaceResponseEnvelopeErrors]
type workerAccountSettingReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerAccountSettingReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerAccountSettingReplaceResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    workerAccountSettingReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// workerAccountSettingReplaceResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [WorkerAccountSettingReplaceResponseEnvelopeMessages]
type workerAccountSettingReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerAccountSettingReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerAccountSettingReplaceResponseEnvelopeSuccess bool

const (
	WorkerAccountSettingReplaceResponseEnvelopeSuccessTrue WorkerAccountSettingReplaceResponseEnvelopeSuccess = true
)
