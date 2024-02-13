// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// WorkerDurableObjectNamespaceObjectService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewWorkerDurableObjectNamespaceObjectService] method instead.
type WorkerDurableObjectNamespaceObjectService struct {
	Options []option.RequestOption
}

// NewWorkerDurableObjectNamespaceObjectService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewWorkerDurableObjectNamespaceObjectService(opts ...option.RequestOption) (r *WorkerDurableObjectNamespaceObjectService) {
	r = &WorkerDurableObjectNamespaceObjectService{}
	r.Options = opts
	return
}

// Returns the Durable Objects in a given namespace.
func (r *WorkerDurableObjectNamespaceObjectService) List(ctx context.Context, accountID string, id string, query WorkerDurableObjectNamespaceObjectListParams, opts ...option.RequestOption) (res *[]WorkerDurableObjectNamespaceObjectListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerDurableObjectNamespaceObjectListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/durable_objects/namespaces/%s/objects", accountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkerDurableObjectNamespaceObjectListResponse struct {
	// ID of the Durable Object.
	ID string `json:"id"`
	// Whether the Durable Object has stored data.
	HasStoredData bool                                               `json:"hasStoredData"`
	JSON          workerDurableObjectNamespaceObjectListResponseJSON `json:"-"`
}

// workerDurableObjectNamespaceObjectListResponseJSON contains the JSON metadata
// for the struct [WorkerDurableObjectNamespaceObjectListResponse]
type workerDurableObjectNamespaceObjectListResponseJSON struct {
	ID            apijson.Field
	HasStoredData apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkerDurableObjectNamespaceObjectListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDurableObjectNamespaceObjectListParams struct {
	// Opaque token indicating the position from which to continue when requesting the
	// next set of records. A valid value for the cursor can be obtained from the
	// cursors object in the result_info structure.
	Cursor param.Field[string] `query:"cursor"`
	// The number of objects to return. The cursor attribute may be used to iterate
	// over the next batch of objects if there are more than the limit.
	Limit param.Field[float64] `query:"limit"`
}

// URLQuery serializes [WorkerDurableObjectNamespaceObjectListParams]'s query
// parameters as `url.Values`.
func (r WorkerDurableObjectNamespaceObjectListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkerDurableObjectNamespaceObjectListResponseEnvelope struct {
	Errors   []WorkerDurableObjectNamespaceObjectListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerDurableObjectNamespaceObjectListResponseEnvelopeMessages `json:"messages,required"`
	Result   []WorkerDurableObjectNamespaceObjectListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    WorkerDurableObjectNamespaceObjectListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo WorkerDurableObjectNamespaceObjectListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       workerDurableObjectNamespaceObjectListResponseEnvelopeJSON       `json:"-"`
}

// workerDurableObjectNamespaceObjectListResponseEnvelopeJSON contains the JSON
// metadata for the struct [WorkerDurableObjectNamespaceObjectListResponseEnvelope]
type workerDurableObjectNamespaceObjectListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDurableObjectNamespaceObjectListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDurableObjectNamespaceObjectListResponseEnvelopeErrors struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    workerDurableObjectNamespaceObjectListResponseEnvelopeErrorsJSON `json:"-"`
}

// workerDurableObjectNamespaceObjectListResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [WorkerDurableObjectNamespaceObjectListResponseEnvelopeErrors]
type workerDurableObjectNamespaceObjectListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDurableObjectNamespaceObjectListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDurableObjectNamespaceObjectListResponseEnvelopeMessages struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    workerDurableObjectNamespaceObjectListResponseEnvelopeMessagesJSON `json:"-"`
}

// workerDurableObjectNamespaceObjectListResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [WorkerDurableObjectNamespaceObjectListResponseEnvelopeMessages]
type workerDurableObjectNamespaceObjectListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDurableObjectNamespaceObjectListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerDurableObjectNamespaceObjectListResponseEnvelopeSuccess bool

const (
	WorkerDurableObjectNamespaceObjectListResponseEnvelopeSuccessTrue WorkerDurableObjectNamespaceObjectListResponseEnvelopeSuccess = true
)

type WorkerDurableObjectNamespaceObjectListResponseEnvelopeResultInfo struct {
	// Total results returned based on your list parameters.
	Count float64 `json:"count"`
	// Opaque token indicating the position from which to continue when requesting the
	// next set of records. A valid value for the cursor can be obtained from the
	// cursors object in the result_info structure.
	Cursor string `json:"cursor"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                              `json:"total_count"`
	JSON       workerDurableObjectNamespaceObjectListResponseEnvelopeResultInfoJSON `json:"-"`
}

// workerDurableObjectNamespaceObjectListResponseEnvelopeResultInfoJSON contains
// the JSON metadata for the struct
// [WorkerDurableObjectNamespaceObjectListResponseEnvelopeResultInfo]
type workerDurableObjectNamespaceObjectListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Cursor      apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDurableObjectNamespaceObjectListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
