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

// Creates a Workers subdomain for an account.
func (r *WorkerSubdomainService) WorkerSubdomainNewSubdomain(ctx context.Context, accountID string, body WorkerSubdomainWorkerSubdomainNewSubdomainParams, opts ...option.RequestOption) (res *WorkerSubdomainWorkerSubdomainNewSubdomainResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerSubdomainWorkerSubdomainNewSubdomainResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/subdomain", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns a Workers subdomain for an account.
func (r *WorkerSubdomainService) WorkerSubdomainGetSubdomain(ctx context.Context, accountID string, opts ...option.RequestOption) (res *WorkerSubdomainWorkerSubdomainGetSubdomainResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerSubdomainWorkerSubdomainGetSubdomainResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/subdomain", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkerSubdomainWorkerSubdomainNewSubdomainResponse struct {
	Name interface{}                                            `json:"name"`
	JSON workerSubdomainWorkerSubdomainNewSubdomainResponseJSON `json:"-"`
}

// workerSubdomainWorkerSubdomainNewSubdomainResponseJSON contains the JSON
// metadata for the struct [WorkerSubdomainWorkerSubdomainNewSubdomainResponse]
type workerSubdomainWorkerSubdomainNewSubdomainResponseJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerSubdomainWorkerSubdomainNewSubdomainResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerSubdomainWorkerSubdomainGetSubdomainResponse struct {
	Name interface{}                                            `json:"name"`
	JSON workerSubdomainWorkerSubdomainGetSubdomainResponseJSON `json:"-"`
}

// workerSubdomainWorkerSubdomainGetSubdomainResponseJSON contains the JSON
// metadata for the struct [WorkerSubdomainWorkerSubdomainGetSubdomainResponse]
type workerSubdomainWorkerSubdomainGetSubdomainResponseJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerSubdomainWorkerSubdomainGetSubdomainResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerSubdomainWorkerSubdomainNewSubdomainParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r WorkerSubdomainWorkerSubdomainNewSubdomainParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type WorkerSubdomainWorkerSubdomainNewSubdomainResponseEnvelope struct {
	Errors   []WorkerSubdomainWorkerSubdomainNewSubdomainResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerSubdomainWorkerSubdomainNewSubdomainResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerSubdomainWorkerSubdomainNewSubdomainResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerSubdomainWorkerSubdomainNewSubdomainResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerSubdomainWorkerSubdomainNewSubdomainResponseEnvelopeJSON    `json:"-"`
}

// workerSubdomainWorkerSubdomainNewSubdomainResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [WorkerSubdomainWorkerSubdomainNewSubdomainResponseEnvelope]
type workerSubdomainWorkerSubdomainNewSubdomainResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerSubdomainWorkerSubdomainNewSubdomainResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerSubdomainWorkerSubdomainNewSubdomainResponseEnvelopeErrors struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    workerSubdomainWorkerSubdomainNewSubdomainResponseEnvelopeErrorsJSON `json:"-"`
}

// workerSubdomainWorkerSubdomainNewSubdomainResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [WorkerSubdomainWorkerSubdomainNewSubdomainResponseEnvelopeErrors]
type workerSubdomainWorkerSubdomainNewSubdomainResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerSubdomainWorkerSubdomainNewSubdomainResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerSubdomainWorkerSubdomainNewSubdomainResponseEnvelopeMessages struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    workerSubdomainWorkerSubdomainNewSubdomainResponseEnvelopeMessagesJSON `json:"-"`
}

// workerSubdomainWorkerSubdomainNewSubdomainResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [WorkerSubdomainWorkerSubdomainNewSubdomainResponseEnvelopeMessages]
type workerSubdomainWorkerSubdomainNewSubdomainResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerSubdomainWorkerSubdomainNewSubdomainResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerSubdomainWorkerSubdomainNewSubdomainResponseEnvelopeSuccess bool

const (
	WorkerSubdomainWorkerSubdomainNewSubdomainResponseEnvelopeSuccessTrue WorkerSubdomainWorkerSubdomainNewSubdomainResponseEnvelopeSuccess = true
)

type WorkerSubdomainWorkerSubdomainGetSubdomainResponseEnvelope struct {
	Errors   []WorkerSubdomainWorkerSubdomainGetSubdomainResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerSubdomainWorkerSubdomainGetSubdomainResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerSubdomainWorkerSubdomainGetSubdomainResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerSubdomainWorkerSubdomainGetSubdomainResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerSubdomainWorkerSubdomainGetSubdomainResponseEnvelopeJSON    `json:"-"`
}

// workerSubdomainWorkerSubdomainGetSubdomainResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [WorkerSubdomainWorkerSubdomainGetSubdomainResponseEnvelope]
type workerSubdomainWorkerSubdomainGetSubdomainResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerSubdomainWorkerSubdomainGetSubdomainResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerSubdomainWorkerSubdomainGetSubdomainResponseEnvelopeErrors struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    workerSubdomainWorkerSubdomainGetSubdomainResponseEnvelopeErrorsJSON `json:"-"`
}

// workerSubdomainWorkerSubdomainGetSubdomainResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [WorkerSubdomainWorkerSubdomainGetSubdomainResponseEnvelopeErrors]
type workerSubdomainWorkerSubdomainGetSubdomainResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerSubdomainWorkerSubdomainGetSubdomainResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerSubdomainWorkerSubdomainGetSubdomainResponseEnvelopeMessages struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    workerSubdomainWorkerSubdomainGetSubdomainResponseEnvelopeMessagesJSON `json:"-"`
}

// workerSubdomainWorkerSubdomainGetSubdomainResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [WorkerSubdomainWorkerSubdomainGetSubdomainResponseEnvelopeMessages]
type workerSubdomainWorkerSubdomainGetSubdomainResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerSubdomainWorkerSubdomainGetSubdomainResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerSubdomainWorkerSubdomainGetSubdomainResponseEnvelopeSuccess bool

const (
	WorkerSubdomainWorkerSubdomainGetSubdomainResponseEnvelopeSuccessTrue WorkerSubdomainWorkerSubdomainGetSubdomainResponseEnvelopeSuccess = true
)
