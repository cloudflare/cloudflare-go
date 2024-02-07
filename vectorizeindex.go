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
func (r *VectorizeIndexService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *VectorizeIndexListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/vectorize/indexes", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type VectorizeIndexListResponse struct {
	Errors   []VectorizeIndexListResponseError   `json:"errors"`
	Messages []VectorizeIndexListResponseMessage `json:"messages"`
	Result   []VectorizeIndexListResponseResult  `json:"result"`
	// Whether the API call was successful
	Success VectorizeIndexListResponseSuccess `json:"success"`
	JSON    vectorizeIndexListResponseJSON    `json:"-"`
}

// vectorizeIndexListResponseJSON contains the JSON metadata for the struct
// [VectorizeIndexListResponse]
type vectorizeIndexListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VectorizeIndexListResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    vectorizeIndexListResponseErrorJSON `json:"-"`
}

// vectorizeIndexListResponseErrorJSON contains the JSON metadata for the struct
// [VectorizeIndexListResponseError]
type vectorizeIndexListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VectorizeIndexListResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    vectorizeIndexListResponseMessageJSON `json:"-"`
}

// vectorizeIndexListResponseMessageJSON contains the JSON metadata for the struct
// [VectorizeIndexListResponseMessage]
type vectorizeIndexListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VectorizeIndexListResponseResult struct {
	Config VectorizeIndexListResponseResultConfig `json:"config"`
	// Specifies the timestamp the resource was created as an ISO8601 string.
	CreatedOn interface{} `json:"created_on"`
	// Specifies the description of the index.
	Description string `json:"description"`
	// Specifies the timestamp the resource was modified as an ISO8601 string.
	ModifiedOn interface{}                          `json:"modified_on"`
	Name       string                               `json:"name"`
	JSON       vectorizeIndexListResponseResultJSON `json:"-"`
}

// vectorizeIndexListResponseResultJSON contains the JSON metadata for the struct
// [VectorizeIndexListResponseResult]
type vectorizeIndexListResponseResultJSON struct {
	Config      apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VectorizeIndexListResponseResultConfig struct {
	// Specifies the number of dimensions for the index
	Dimensions int64 `json:"dimensions,required"`
	// Specifies the type of metric to use calculating distance.
	Metric VectorizeIndexListResponseResultConfigMetric `json:"metric,required"`
	JSON   vectorizeIndexListResponseResultConfigJSON   `json:"-"`
}

// vectorizeIndexListResponseResultConfigJSON contains the JSON metadata for the
// struct [VectorizeIndexListResponseResultConfig]
type vectorizeIndexListResponseResultConfigJSON struct {
	Dimensions  apijson.Field
	Metric      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexListResponseResultConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the type of metric to use calculating distance.
type VectorizeIndexListResponseResultConfigMetric string

const (
	VectorizeIndexListResponseResultConfigMetricCosine     VectorizeIndexListResponseResultConfigMetric = "cosine"
	VectorizeIndexListResponseResultConfigMetricEuclidean  VectorizeIndexListResponseResultConfigMetric = "euclidean"
	VectorizeIndexListResponseResultConfigMetricDotProduct VectorizeIndexListResponseResultConfigMetric = "dot-product"
)

// Whether the API call was successful
type VectorizeIndexListResponseSuccess bool

const (
	VectorizeIndexListResponseSuccessTrue VectorizeIndexListResponseSuccess = true
)
