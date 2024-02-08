// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// VectorizeIndexService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewVectorizeIndexService] method
// instead.
type VectorizeIndexService struct {
	Options []option.RequestOption
}

// NewVectorizeIndexService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVectorizeIndexService(opts ...option.RequestOption) (r *VectorizeIndexService) {
	r = &VectorizeIndexService{}
	r.Options = opts
	return
}

// Returns a list of Vectorize Indexes
func (r *VectorizeIndexService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *[]VectorizeIndexListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env VectorizeIndexListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/vectorize/indexes", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type VectorizeIndexListResponse struct {
	Config VectorizeIndexListResponseConfig `json:"config"`
	// Specifies the timestamp the resource was created as an ISO8601 string.
	CreatedOn interface{} `json:"created_on"`
	// Specifies the description of the index.
	Description string `json:"description"`
	// Specifies the timestamp the resource was modified as an ISO8601 string.
	ModifiedOn interface{}                    `json:"modified_on"`
	Name       string                         `json:"name"`
	JSON       vectorizeIndexListResponseJSON `json:"-"`
}

// vectorizeIndexListResponseJSON contains the JSON metadata for the struct
// [VectorizeIndexListResponse]
type vectorizeIndexListResponseJSON struct {
	Config      apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VectorizeIndexListResponseConfig struct {
	// Specifies the number of dimensions for the index
	Dimensions int64 `json:"dimensions,required"`
	// Specifies the type of metric to use calculating distance.
	Metric VectorizeIndexListResponseConfigMetric `json:"metric,required"`
	JSON   vectorizeIndexListResponseConfigJSON   `json:"-"`
}

// vectorizeIndexListResponseConfigJSON contains the JSON metadata for the struct
// [VectorizeIndexListResponseConfig]
type vectorizeIndexListResponseConfigJSON struct {
	Dimensions  apijson.Field
	Metric      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexListResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the type of metric to use calculating distance.
type VectorizeIndexListResponseConfigMetric string

const (
	VectorizeIndexListResponseConfigMetricCosine     VectorizeIndexListResponseConfigMetric = "cosine"
	VectorizeIndexListResponseConfigMetricEuclidean  VectorizeIndexListResponseConfigMetric = "euclidean"
	VectorizeIndexListResponseConfigMetricDotProduct VectorizeIndexListResponseConfigMetric = "dot-product"
)

type VectorizeIndexListResponseEnvelope struct {
	Errors   []VectorizeIndexListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []VectorizeIndexListResponseEnvelopeMessages `json:"messages,required"`
	Result   []VectorizeIndexListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success VectorizeIndexListResponseEnvelopeSuccess `json:"success,required"`
	JSON    vectorizeIndexListResponseEnvelopeJSON    `json:"-"`
}

// vectorizeIndexListResponseEnvelopeJSON contains the JSON metadata for the struct
// [VectorizeIndexListResponseEnvelope]
type vectorizeIndexListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VectorizeIndexListResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    vectorizeIndexListResponseEnvelopeErrorsJSON `json:"-"`
}

// vectorizeIndexListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [VectorizeIndexListResponseEnvelopeErrors]
type vectorizeIndexListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VectorizeIndexListResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    vectorizeIndexListResponseEnvelopeMessagesJSON `json:"-"`
}

// vectorizeIndexListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [VectorizeIndexListResponseEnvelopeMessages]
type vectorizeIndexListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type VectorizeIndexListResponseEnvelopeSuccess bool

const (
	VectorizeIndexListResponseEnvelopeSuccessTrue VectorizeIndexListResponseEnvelopeSuccess = true
)
