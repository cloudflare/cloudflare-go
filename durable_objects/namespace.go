// File generated from our OpenAPI spec by Stainless.

package durable_objects

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// NamespaceService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewNamespaceService] method instead.
type NamespaceService struct {
	Options []option.RequestOption
	Objects *NamespaceObjectService
}

// NewNamespaceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNamespaceService(opts ...option.RequestOption) (r *NamespaceService) {
	r = &NamespaceService{}
	r.Options = opts
	r.Objects = NewNamespaceObjectService(opts...)
	return
}

// Returns the Durable Object namespaces owned by an account.
func (r *NamespaceService) List(ctx context.Context, query NamespaceListParams, opts ...option.RequestOption) (res *[]WorkersNamespace, err error) {
	opts = append(r.Options[:], opts...)
	var env NamespaceListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/durable_objects/namespaces", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkersNamespace struct {
	ID     interface{}          `json:"id"`
	Class  interface{}          `json:"class"`
	Name   interface{}          `json:"name"`
	Script interface{}          `json:"script"`
	JSON   workersNamespaceJSON `json:"-"`
}

// workersNamespaceJSON contains the JSON metadata for the struct
// [WorkersNamespace]
type workersNamespaceJSON struct {
	ID          apijson.Field
	Class       apijson.Field
	Name        apijson.Field
	Script      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersNamespaceJSON) RawJSON() string {
	return r.raw
}

type NamespaceListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type NamespaceListResponseEnvelope struct {
	Errors   []NamespaceListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []NamespaceListResponseEnvelopeMessages `json:"messages,required"`
	Result   []WorkersNamespace                      `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    NamespaceListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo NamespaceListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       namespaceListResponseEnvelopeJSON       `json:"-"`
}

// namespaceListResponseEnvelopeJSON contains the JSON metadata for the struct
// [NamespaceListResponseEnvelope]
type namespaceListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceListResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    namespaceListResponseEnvelopeErrorsJSON `json:"-"`
}

// namespaceListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [NamespaceListResponseEnvelopeErrors]
type namespaceListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type NamespaceListResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    namespaceListResponseEnvelopeMessagesJSON `json:"-"`
}

// namespaceListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [NamespaceListResponseEnvelopeMessages]
type namespaceListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NamespaceListResponseEnvelopeSuccess bool

const (
	NamespaceListResponseEnvelopeSuccessTrue NamespaceListResponseEnvelopeSuccess = true
)

type NamespaceListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                     `json:"total_count"`
	JSON       namespaceListResponseEnvelopeResultInfoJSON `json:"-"`
}

// namespaceListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [NamespaceListResponseEnvelopeResultInfo]
type namespaceListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
