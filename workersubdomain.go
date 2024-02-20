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

// WorkerSubdomainService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWorkerSubdomainService] method
// instead.
type WorkerSubdomainService struct {
	Options []option.RequestOption
}

// NewWorkerSubdomainService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkerSubdomainService(opts ...option.RequestOption) (r *WorkerSubdomainService) {
	r = &WorkerSubdomainService{}
	r.Options = opts
	return
}

// Returns a Workers subdomain for an account.
func (r *WorkerSubdomainService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *WorkerSubdomainGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerSubdomainGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/subdomain", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Creates a Workers subdomain for an account.
func (r *WorkerSubdomainService) Replace(ctx context.Context, accountID string, body WorkerSubdomainReplaceParams, opts ...option.RequestOption) (res *WorkerSubdomainReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerSubdomainReplaceResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/subdomain", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkerSubdomainGetResponse struct {
	Name interface{}                    `json:"name"`
	JSON workerSubdomainGetResponseJSON `json:"-"`
}

// workerSubdomainGetResponseJSON contains the JSON metadata for the struct
// [WorkerSubdomainGetResponse]
type workerSubdomainGetResponseJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerSubdomainGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerSubdomainReplaceResponse struct {
	Name interface{}                        `json:"name"`
	JSON workerSubdomainReplaceResponseJSON `json:"-"`
}

// workerSubdomainReplaceResponseJSON contains the JSON metadata for the struct
// [WorkerSubdomainReplaceResponse]
type workerSubdomainReplaceResponseJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerSubdomainReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerSubdomainGetResponseEnvelope struct {
	Errors   []WorkerSubdomainGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerSubdomainGetResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerSubdomainGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerSubdomainGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerSubdomainGetResponseEnvelopeJSON    `json:"-"`
}

// workerSubdomainGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [WorkerSubdomainGetResponseEnvelope]
type workerSubdomainGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerSubdomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerSubdomainGetResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    workerSubdomainGetResponseEnvelopeErrorsJSON `json:"-"`
}

// workerSubdomainGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WorkerSubdomainGetResponseEnvelopeErrors]
type workerSubdomainGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerSubdomainGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerSubdomainGetResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    workerSubdomainGetResponseEnvelopeMessagesJSON `json:"-"`
}

// workerSubdomainGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [WorkerSubdomainGetResponseEnvelopeMessages]
type workerSubdomainGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerSubdomainGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerSubdomainGetResponseEnvelopeSuccess bool

const (
	WorkerSubdomainGetResponseEnvelopeSuccessTrue WorkerSubdomainGetResponseEnvelopeSuccess = true
)

type WorkerSubdomainReplaceParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r WorkerSubdomainReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type WorkerSubdomainReplaceResponseEnvelope struct {
	Errors   []WorkerSubdomainReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerSubdomainReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerSubdomainReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerSubdomainReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerSubdomainReplaceResponseEnvelopeJSON    `json:"-"`
}

// workerSubdomainReplaceResponseEnvelopeJSON contains the JSON metadata for the
// struct [WorkerSubdomainReplaceResponseEnvelope]
type workerSubdomainReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerSubdomainReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerSubdomainReplaceResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    workerSubdomainReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// workerSubdomainReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [WorkerSubdomainReplaceResponseEnvelopeErrors]
type workerSubdomainReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerSubdomainReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerSubdomainReplaceResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    workerSubdomainReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// workerSubdomainReplaceResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [WorkerSubdomainReplaceResponseEnvelopeMessages]
type workerSubdomainReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerSubdomainReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerSubdomainReplaceResponseEnvelopeSuccess bool

const (
	WorkerSubdomainReplaceResponseEnvelopeSuccessTrue WorkerSubdomainReplaceResponseEnvelopeSuccess = true
)
