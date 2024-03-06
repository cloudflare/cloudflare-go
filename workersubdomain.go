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
func (r *WorkerSubdomainService) Update(ctx context.Context, params WorkerSubdomainUpdateParams, opts ...option.RequestOption) (res *WorkerSubdomainUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerSubdomainUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/subdomain", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns a Workers subdomain for an account.
func (r *WorkerSubdomainService) Get(ctx context.Context, query WorkerSubdomainGetParams, opts ...option.RequestOption) (res *WorkerSubdomainGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerSubdomainGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/subdomain", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkerSubdomainUpdateResponse struct {
	Name interface{}                       `json:"name"`
	JSON workerSubdomainUpdateResponseJSON `json:"-"`
}

// workerSubdomainUpdateResponseJSON contains the JSON metadata for the struct
// [WorkerSubdomainUpdateResponse]
type workerSubdomainUpdateResponseJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerSubdomainUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerSubdomainUpdateResponseJSON) RawJSON() string {
	return r.raw
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

func (r workerSubdomainGetResponseJSON) RawJSON() string {
	return r.raw
}

type WorkerSubdomainUpdateParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r WorkerSubdomainUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type WorkerSubdomainUpdateResponseEnvelope struct {
	Errors   []WorkerSubdomainUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerSubdomainUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerSubdomainUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerSubdomainUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerSubdomainUpdateResponseEnvelopeJSON    `json:"-"`
}

// workerSubdomainUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [WorkerSubdomainUpdateResponseEnvelope]
type workerSubdomainUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerSubdomainUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerSubdomainUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WorkerSubdomainUpdateResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    workerSubdomainUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// workerSubdomainUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [WorkerSubdomainUpdateResponseEnvelopeErrors]
type workerSubdomainUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerSubdomainUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerSubdomainUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WorkerSubdomainUpdateResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    workerSubdomainUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// workerSubdomainUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [WorkerSubdomainUpdateResponseEnvelopeMessages]
type workerSubdomainUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerSubdomainUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerSubdomainUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WorkerSubdomainUpdateResponseEnvelopeSuccess bool

const (
	WorkerSubdomainUpdateResponseEnvelopeSuccessTrue WorkerSubdomainUpdateResponseEnvelopeSuccess = true
)

type WorkerSubdomainGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
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

func (r workerSubdomainGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r workerSubdomainGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r workerSubdomainGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WorkerSubdomainGetResponseEnvelopeSuccess bool

const (
	WorkerSubdomainGetResponseEnvelopeSuccessTrue WorkerSubdomainGetResponseEnvelopeSuccess = true
)
