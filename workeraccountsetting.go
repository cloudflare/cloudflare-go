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
func (r *WorkerAccountSettingService) Update(ctx context.Context, accountID string, body WorkerAccountSettingUpdateParams, opts ...option.RequestOption) (res *WorkerAccountSettingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerAccountSettingUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/account-settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
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

type WorkerAccountSettingUpdateResponse struct {
	DefaultUsageModel interface{}                            `json:"default_usage_model"`
	GreenCompute      interface{}                            `json:"green_compute"`
	JSON              workerAccountSettingUpdateResponseJSON `json:"-"`
}

// workerAccountSettingUpdateResponseJSON contains the JSON metadata for the struct
// [WorkerAccountSettingUpdateResponse]
type workerAccountSettingUpdateResponseJSON struct {
	DefaultUsageModel apijson.Field
	GreenCompute      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *WorkerAccountSettingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type WorkerAccountSettingUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r WorkerAccountSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type WorkerAccountSettingUpdateResponseEnvelope struct {
	Errors   []WorkerAccountSettingUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerAccountSettingUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerAccountSettingUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerAccountSettingUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerAccountSettingUpdateResponseEnvelopeJSON    `json:"-"`
}

// workerAccountSettingUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [WorkerAccountSettingUpdateResponseEnvelope]
type workerAccountSettingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerAccountSettingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerAccountSettingUpdateResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    workerAccountSettingUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// workerAccountSettingUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [WorkerAccountSettingUpdateResponseEnvelopeErrors]
type workerAccountSettingUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerAccountSettingUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerAccountSettingUpdateResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    workerAccountSettingUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// workerAccountSettingUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [WorkerAccountSettingUpdateResponseEnvelopeMessages]
type workerAccountSettingUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerAccountSettingUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerAccountSettingUpdateResponseEnvelopeSuccess bool

const (
	WorkerAccountSettingUpdateResponseEnvelopeSuccessTrue WorkerAccountSettingUpdateResponseEnvelopeSuccess = true
)

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
