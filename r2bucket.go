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

// R2BucketService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewR2BucketService] method instead.
type R2BucketService struct {
	Options []option.RequestOption
}

// NewR2BucketService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewR2BucketService(opts ...option.RequestOption) (r *R2BucketService) {
	r = &R2BucketService{}
	r.Options = opts
	return
}

// Creates a new R2 bucket.
func (r *R2BucketService) New(ctx context.Context, accountID string, body R2BucketNewParams, opts ...option.RequestOption) (res *R2BucketNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/r2/buckets", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets metadata for an existing R2 bucket.
func (r *R2BucketService) Get(ctx context.Context, accountID string, bucketName string, opts ...option.RequestOption) (res *R2BucketGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s", accountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists all R2 buckets on your account
func (r *R2BucketService) List(ctx context.Context, accountID string, query R2BucketListParams, opts ...option.RequestOption) (res *R2BucketListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/r2/buckets", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes an existing R2 bucket.
func (r *R2BucketService) Delete(ctx context.Context, accountID string, bucketName string, opts ...option.RequestOption) (res *R2BucketDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s", accountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type R2BucketNewResponse struct {
	Errors   []R2BucketNewResponseError `json:"errors"`
	Messages []string                   `json:"messages"`
	// A single R2 bucket
	Result R2BucketNewResponseResult `json:"result"`
	// Whether the API call was successful
	Success R2BucketNewResponseSuccess `json:"success"`
	JSON    r2BucketNewResponseJSON    `json:"-"`
}

// r2BucketNewResponseJSON contains the JSON metadata for the struct
// [R2BucketNewResponse]
type r2BucketNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2BucketNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type R2BucketNewResponseError struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    r2BucketNewResponseErrorJSON `json:"-"`
}

// r2BucketNewResponseErrorJSON contains the JSON metadata for the struct
// [R2BucketNewResponseError]
type r2BucketNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2BucketNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A single R2 bucket
type R2BucketNewResponseResult struct {
	// Creation timestamp
	CreationDate string `json:"creation_date"`
	// Location of the bucket
	Location R2BucketNewResponseResultLocation `json:"location"`
	// Name of the bucket
	Name string                        `json:"name"`
	JSON r2BucketNewResponseResultJSON `json:"-"`
}

// r2BucketNewResponseResultJSON contains the JSON metadata for the struct
// [R2BucketNewResponseResult]
type r2BucketNewResponseResultJSON struct {
	CreationDate apijson.Field
	Location     apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *R2BucketNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Location of the bucket
type R2BucketNewResponseResultLocation string

const (
	R2BucketNewResponseResultLocationApac R2BucketNewResponseResultLocation = "apac"
	R2BucketNewResponseResultLocationEeur R2BucketNewResponseResultLocation = "eeur"
	R2BucketNewResponseResultLocationEnam R2BucketNewResponseResultLocation = "enam"
	R2BucketNewResponseResultLocationWeur R2BucketNewResponseResultLocation = "weur"
	R2BucketNewResponseResultLocationWnam R2BucketNewResponseResultLocation = "wnam"
)

// Whether the API call was successful
type R2BucketNewResponseSuccess bool

const (
	R2BucketNewResponseSuccessTrue R2BucketNewResponseSuccess = true
)

type R2BucketGetResponse struct {
	Errors   []R2BucketGetResponseError `json:"errors"`
	Messages []string                   `json:"messages"`
	// A single R2 bucket
	Result R2BucketGetResponseResult `json:"result"`
	// Whether the API call was successful
	Success R2BucketGetResponseSuccess `json:"success"`
	JSON    r2BucketGetResponseJSON    `json:"-"`
}

// r2BucketGetResponseJSON contains the JSON metadata for the struct
// [R2BucketGetResponse]
type r2BucketGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2BucketGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type R2BucketGetResponseError struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    r2BucketGetResponseErrorJSON `json:"-"`
}

// r2BucketGetResponseErrorJSON contains the JSON metadata for the struct
// [R2BucketGetResponseError]
type r2BucketGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2BucketGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A single R2 bucket
type R2BucketGetResponseResult struct {
	// Creation timestamp
	CreationDate string `json:"creation_date"`
	// Location of the bucket
	Location R2BucketGetResponseResultLocation `json:"location"`
	// Name of the bucket
	Name string                        `json:"name"`
	JSON r2BucketGetResponseResultJSON `json:"-"`
}

// r2BucketGetResponseResultJSON contains the JSON metadata for the struct
// [R2BucketGetResponseResult]
type r2BucketGetResponseResultJSON struct {
	CreationDate apijson.Field
	Location     apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *R2BucketGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Location of the bucket
type R2BucketGetResponseResultLocation string

const (
	R2BucketGetResponseResultLocationApac R2BucketGetResponseResultLocation = "apac"
	R2BucketGetResponseResultLocationEeur R2BucketGetResponseResultLocation = "eeur"
	R2BucketGetResponseResultLocationEnam R2BucketGetResponseResultLocation = "enam"
	R2BucketGetResponseResultLocationWeur R2BucketGetResponseResultLocation = "weur"
	R2BucketGetResponseResultLocationWnam R2BucketGetResponseResultLocation = "wnam"
)

// Whether the API call was successful
type R2BucketGetResponseSuccess bool

const (
	R2BucketGetResponseSuccessTrue R2BucketGetResponseSuccess = true
)

type R2BucketListResponse struct {
	Errors     []R2BucketListResponseError    `json:"errors"`
	Messages   []string                       `json:"messages"`
	Result     []R2BucketListResponseResult   `json:"result"`
	ResultInfo R2BucketListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success R2BucketListResponseSuccess `json:"success"`
	JSON    r2BucketListResponseJSON    `json:"-"`
}

// r2BucketListResponseJSON contains the JSON metadata for the struct
// [R2BucketListResponse]
type r2BucketListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2BucketListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type R2BucketListResponseError struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    r2BucketListResponseErrorJSON `json:"-"`
}

// r2BucketListResponseErrorJSON contains the JSON metadata for the struct
// [R2BucketListResponseError]
type r2BucketListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2BucketListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A single R2 bucket
type R2BucketListResponseResult struct {
	// Creation timestamp
	CreationDate string `json:"creation_date"`
	// Location of the bucket
	Location R2BucketListResponseResultLocation `json:"location"`
	// Name of the bucket
	Name string                         `json:"name"`
	JSON r2BucketListResponseResultJSON `json:"-"`
}

// r2BucketListResponseResultJSON contains the JSON metadata for the struct
// [R2BucketListResponseResult]
type r2BucketListResponseResultJSON struct {
	CreationDate apijson.Field
	Location     apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *R2BucketListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Location of the bucket
type R2BucketListResponseResultLocation string

const (
	R2BucketListResponseResultLocationApac R2BucketListResponseResultLocation = "apac"
	R2BucketListResponseResultLocationEeur R2BucketListResponseResultLocation = "eeur"
	R2BucketListResponseResultLocationEnam R2BucketListResponseResultLocation = "enam"
	R2BucketListResponseResultLocationWeur R2BucketListResponseResultLocation = "weur"
	R2BucketListResponseResultLocationWnam R2BucketListResponseResultLocation = "wnam"
)

type R2BucketListResponseResultInfo struct {
	// A continuation token that should be used to fetch the next page of results
	Cursor string `json:"cursor"`
	// Maximum number of results on this page
	PerPage float64                            `json:"per_page"`
	JSON    r2BucketListResponseResultInfoJSON `json:"-"`
}

// r2BucketListResponseResultInfoJSON contains the JSON metadata for the struct
// [R2BucketListResponseResultInfo]
type r2BucketListResponseResultInfoJSON struct {
	Cursor      apijson.Field
	PerPage     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2BucketListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type R2BucketListResponseSuccess bool

const (
	R2BucketListResponseSuccessTrue R2BucketListResponseSuccess = true
)

type R2BucketDeleteResponse struct {
	Errors   []R2BucketDeleteResponseError `json:"errors,required"`
	Messages []string                      `json:"messages,required"`
	Result   interface{}                   `json:"result,required"`
	// Whether the API call was successful
	Success R2BucketDeleteResponseSuccess `json:"success,required"`
	JSON    r2BucketDeleteResponseJSON    `json:"-"`
}

// r2BucketDeleteResponseJSON contains the JSON metadata for the struct
// [R2BucketDeleteResponse]
type r2BucketDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2BucketDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type R2BucketDeleteResponseError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    r2BucketDeleteResponseErrorJSON `json:"-"`
}

// r2BucketDeleteResponseErrorJSON contains the JSON metadata for the struct
// [R2BucketDeleteResponseError]
type r2BucketDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2BucketDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type R2BucketDeleteResponseSuccess bool

const (
	R2BucketDeleteResponseSuccessTrue R2BucketDeleteResponseSuccess = true
)

type R2BucketNewParams struct {
	// Name of the bucket
	Name param.Field[string] `json:"name,required"`
	// Location of the bucket
	LocationHint param.Field[R2BucketNewParamsLocationHint] `json:"locationHint"`
}

func (r R2BucketNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Location of the bucket
type R2BucketNewParamsLocationHint string

const (
	R2BucketNewParamsLocationHintApac R2BucketNewParamsLocationHint = "apac"
	R2BucketNewParamsLocationHintEeur R2BucketNewParamsLocationHint = "eeur"
	R2BucketNewParamsLocationHintEnam R2BucketNewParamsLocationHint = "enam"
	R2BucketNewParamsLocationHintWeur R2BucketNewParamsLocationHint = "weur"
	R2BucketNewParamsLocationHintWnam R2BucketNewParamsLocationHint = "wnam"
)

type R2BucketListParams struct {
	// Pagination cursor received during the last List Buckets call. R2 buckets are
	// paginated using cursors instead of page numbers.
	Cursor param.Field[string] `query:"cursor"`
	// Direction to order buckets
	Direction param.Field[R2BucketListParamsDirection] `query:"direction"`
	// Bucket names to filter by. Only buckets with this phrase in their name will be
	// returned.
	NameContains param.Field[string] `query:"name_contains"`
	// Field to order buckets by
	Order param.Field[R2BucketListParamsOrder] `query:"order"`
	// Maximum number of buckets to return in a single call
	PerPage param.Field[float64] `query:"per_page"`
	// Bucket name to start searching after. Buckets are ordered lexicographically.
	StartAfter param.Field[string] `query:"start_after"`
}

// URLQuery serializes [R2BucketListParams]'s query parameters as `url.Values`.
func (r R2BucketListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order buckets
type R2BucketListParamsDirection string

const (
	R2BucketListParamsDirectionAsc  R2BucketListParamsDirection = "asc"
	R2BucketListParamsDirectionDesc R2BucketListParamsDirection = "desc"
)

// Field to order buckets by
type R2BucketListParamsOrder string

const (
	R2BucketListParamsOrderName R2BucketListParamsOrder = "name"
)
