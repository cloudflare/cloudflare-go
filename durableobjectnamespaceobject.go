// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// DurableObjectNamespaceObjectService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewDurableObjectNamespaceObjectService] method instead.
type DurableObjectNamespaceObjectService struct {
	Options []option.RequestOption
}

// NewDurableObjectNamespaceObjectService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDurableObjectNamespaceObjectService(opts ...option.RequestOption) (r *DurableObjectNamespaceObjectService) {
	r = &DurableObjectNamespaceObjectService{}
	r.Options = opts
	return
}

// Returns the Durable Objects in a given namespace.
func (r *DurableObjectNamespaceObjectService) List(ctx context.Context, id string, params DurableObjectNamespaceObjectListParams, opts ...option.RequestOption) (res *[]DurableObjectNamespaceObjectListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DurableObjectNamespaceObjectListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/durable_objects/namespaces/%s/objects", params.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DurableObjectNamespaceObjectListResponse struct {
	// ID of the Durable Object.
	ID string `json:"id"`
	// Whether the Durable Object has stored data.
	HasStoredData bool                                         `json:"hasStoredData"`
	JSON          durableObjectNamespaceObjectListResponseJSON `json:"-"`
}

// durableObjectNamespaceObjectListResponseJSON contains the JSON metadata for the
// struct [DurableObjectNamespaceObjectListResponse]
type durableObjectNamespaceObjectListResponseJSON struct {
	ID            apijson.Field
	HasStoredData apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DurableObjectNamespaceObjectListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r durableObjectNamespaceObjectListResponseJSON) RawJSON() string {
	return r.raw
}

type DurableObjectNamespaceObjectListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Opaque token indicating the position from which to continue when requesting the
	// next set of records. A valid value for the cursor can be obtained from the
	// cursors object in the result_info structure.
	Cursor param.Field[string] `query:"cursor"`
	// The number of objects to return. The cursor attribute may be used to iterate
	// over the next batch of objects if there are more than the limit.
	Limit param.Field[float64] `query:"limit"`
}

// URLQuery serializes [DurableObjectNamespaceObjectListParams]'s query parameters
// as `url.Values`.
func (r DurableObjectNamespaceObjectListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DurableObjectNamespaceObjectListResponseEnvelope struct {
	Errors   []DurableObjectNamespaceObjectListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DurableObjectNamespaceObjectListResponseEnvelopeMessages `json:"messages,required"`
	Result   []DurableObjectNamespaceObjectListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    DurableObjectNamespaceObjectListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DurableObjectNamespaceObjectListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       durableObjectNamespaceObjectListResponseEnvelopeJSON       `json:"-"`
}

// durableObjectNamespaceObjectListResponseEnvelopeJSON contains the JSON metadata
// for the struct [DurableObjectNamespaceObjectListResponseEnvelope]
type durableObjectNamespaceObjectListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DurableObjectNamespaceObjectListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r durableObjectNamespaceObjectListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DurableObjectNamespaceObjectListResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    durableObjectNamespaceObjectListResponseEnvelopeErrorsJSON `json:"-"`
}

// durableObjectNamespaceObjectListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DurableObjectNamespaceObjectListResponseEnvelopeErrors]
type durableObjectNamespaceObjectListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DurableObjectNamespaceObjectListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r durableObjectNamespaceObjectListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DurableObjectNamespaceObjectListResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    durableObjectNamespaceObjectListResponseEnvelopeMessagesJSON `json:"-"`
}

// durableObjectNamespaceObjectListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [DurableObjectNamespaceObjectListResponseEnvelopeMessages]
type durableObjectNamespaceObjectListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DurableObjectNamespaceObjectListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r durableObjectNamespaceObjectListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DurableObjectNamespaceObjectListResponseEnvelopeSuccess bool

const (
	DurableObjectNamespaceObjectListResponseEnvelopeSuccessTrue DurableObjectNamespaceObjectListResponseEnvelopeSuccess = true
)

type DurableObjectNamespaceObjectListResponseEnvelopeResultInfo struct {
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
	TotalCount float64                                                        `json:"total_count"`
	JSON       durableObjectNamespaceObjectListResponseEnvelopeResultInfoJSON `json:"-"`
}

// durableObjectNamespaceObjectListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [DurableObjectNamespaceObjectListResponseEnvelopeResultInfo]
type durableObjectNamespaceObjectListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Cursor      apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DurableObjectNamespaceObjectListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r durableObjectNamespaceObjectListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
