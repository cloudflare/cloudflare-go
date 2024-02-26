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

// WorkerDurableObjectNamespaceService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewWorkerDurableObjectNamespaceService] method instead.
type WorkerDurableObjectNamespaceService struct {
	Options []option.RequestOption
	Objects *WorkerDurableObjectNamespaceObjectService
}

// NewWorkerDurableObjectNamespaceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWorkerDurableObjectNamespaceService(opts ...option.RequestOption) (r *WorkerDurableObjectNamespaceService) {
	r = &WorkerDurableObjectNamespaceService{}
	r.Options = opts
	r.Objects = NewWorkerDurableObjectNamespaceObjectService(opts...)
	return
}

// Returns the Durable Object namespaces owned by an account.
func (r *WorkerDurableObjectNamespaceService) List(ctx context.Context, query WorkerDurableObjectNamespaceListParams, opts ...option.RequestOption) (res *[]WorkerDurableObjectNamespaceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerDurableObjectNamespaceListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/durable_objects/namespaces", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkerDurableObjectNamespaceListResponse struct {
	ID     interface{}                                  `json:"id"`
	Class  interface{}                                  `json:"class"`
	Name   interface{}                                  `json:"name"`
	Script interface{}                                  `json:"script"`
	JSON   workerDurableObjectNamespaceListResponseJSON `json:"-"`
}

// workerDurableObjectNamespaceListResponseJSON contains the JSON metadata for the
// struct [WorkerDurableObjectNamespaceListResponse]
type workerDurableObjectNamespaceListResponseJSON struct {
	ID          apijson.Field
	Class       apijson.Field
	Name        apijson.Field
	Script      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDurableObjectNamespaceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDurableObjectNamespaceListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type WorkerDurableObjectNamespaceListResponseEnvelope struct {
	Errors   []WorkerDurableObjectNamespaceListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerDurableObjectNamespaceListResponseEnvelopeMessages `json:"messages,required"`
	Result   []WorkerDurableObjectNamespaceListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    WorkerDurableObjectNamespaceListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo WorkerDurableObjectNamespaceListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       workerDurableObjectNamespaceListResponseEnvelopeJSON       `json:"-"`
}

// workerDurableObjectNamespaceListResponseEnvelopeJSON contains the JSON metadata
// for the struct [WorkerDurableObjectNamespaceListResponseEnvelope]
type workerDurableObjectNamespaceListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDurableObjectNamespaceListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDurableObjectNamespaceListResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    workerDurableObjectNamespaceListResponseEnvelopeErrorsJSON `json:"-"`
}

// workerDurableObjectNamespaceListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [WorkerDurableObjectNamespaceListResponseEnvelopeErrors]
type workerDurableObjectNamespaceListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDurableObjectNamespaceListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDurableObjectNamespaceListResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    workerDurableObjectNamespaceListResponseEnvelopeMessagesJSON `json:"-"`
}

// workerDurableObjectNamespaceListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [WorkerDurableObjectNamespaceListResponseEnvelopeMessages]
type workerDurableObjectNamespaceListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDurableObjectNamespaceListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerDurableObjectNamespaceListResponseEnvelopeSuccess bool

const (
	WorkerDurableObjectNamespaceListResponseEnvelopeSuccessTrue WorkerDurableObjectNamespaceListResponseEnvelopeSuccess = true
)

type WorkerDurableObjectNamespaceListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                        `json:"total_count"`
	JSON       workerDurableObjectNamespaceListResponseEnvelopeResultInfoJSON `json:"-"`
}

// workerDurableObjectNamespaceListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [WorkerDurableObjectNamespaceListResponseEnvelopeResultInfo]
type workerDurableObjectNamespaceListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDurableObjectNamespaceListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
