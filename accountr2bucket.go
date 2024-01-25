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

// AccountR2BucketService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountR2BucketService] method
// instead.
type AccountR2BucketService struct {
	Options []option.RequestOption
}

// NewAccountR2BucketService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountR2BucketService(opts ...option.RequestOption) (r *AccountR2BucketService) {
	r = &AccountR2BucketService{}
	r.Options = opts
	return
}

// Gets metadata for an existing R2 bucket.
func (r *AccountR2BucketService) Get(ctx context.Context, accountIdentifier string, bucketName string, opts ...option.RequestOption) (res *AccountR2BucketGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s", accountIdentifier, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an existing R2 bucket.
func (r *AccountR2BucketService) Delete(ctx context.Context, accountIdentifier string, bucketName string, opts ...option.RequestOption) (res *AccountR2BucketDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s", accountIdentifier, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a new R2 bucket.
func (r *AccountR2BucketService) R2BucketNewBucket(ctx context.Context, accountIdentifier string, body AccountR2BucketR2BucketNewBucketParams, opts ...option.RequestOption) (res *AccountR2BucketR2BucketNewBucketResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/r2/buckets", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all R2 buckets on your account
func (r *AccountR2BucketService) R2BucketListBuckets(ctx context.Context, accountIdentifier string, query AccountR2BucketR2BucketListBucketsParams, opts ...option.RequestOption) (res *AccountR2BucketR2BucketListBucketsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/r2/buckets", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountR2BucketGetResponse struct {
	Errors   []AccountR2BucketGetResponseError `json:"errors"`
	Messages []string                          `json:"messages"`
	// A single R2 bucket
	Result AccountR2BucketGetResponseResult `json:"result"`
	// Whether the API call was successful
	Success AccountR2BucketGetResponseSuccess `json:"success"`
	JSON    accountR2BucketGetResponseJSON    `json:"-"`
}

// accountR2BucketGetResponseJSON contains the JSON metadata for the struct
// [AccountR2BucketGetResponse]
type accountR2BucketGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountR2BucketGetResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    accountR2BucketGetResponseErrorJSON `json:"-"`
}

// accountR2BucketGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountR2BucketGetResponseError]
type accountR2BucketGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A single R2 bucket
type AccountR2BucketGetResponseResult struct {
	// Creation timestamp
	CreationDate string `json:"creation_date"`
	// Location of the bucket
	Location AccountR2BucketGetResponseResultLocation `json:"location"`
	// Name of the bucket
	Name string                               `json:"name"`
	JSON accountR2BucketGetResponseResultJSON `json:"-"`
}

// accountR2BucketGetResponseResultJSON contains the JSON metadata for the struct
// [AccountR2BucketGetResponseResult]
type accountR2BucketGetResponseResultJSON struct {
	CreationDate apijson.Field
	Location     apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountR2BucketGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Location of the bucket
type AccountR2BucketGetResponseResultLocation string

const (
	AccountR2BucketGetResponseResultLocationApac AccountR2BucketGetResponseResultLocation = "apac"
	AccountR2BucketGetResponseResultLocationEeur AccountR2BucketGetResponseResultLocation = "eeur"
	AccountR2BucketGetResponseResultLocationEnam AccountR2BucketGetResponseResultLocation = "enam"
	AccountR2BucketGetResponseResultLocationWeur AccountR2BucketGetResponseResultLocation = "weur"
	AccountR2BucketGetResponseResultLocationWnam AccountR2BucketGetResponseResultLocation = "wnam"
)

// Whether the API call was successful
type AccountR2BucketGetResponseSuccess bool

const (
	AccountR2BucketGetResponseSuccessTrue AccountR2BucketGetResponseSuccess = true
)

type AccountR2BucketDeleteResponse struct {
	Errors   []AccountR2BucketDeleteResponseError `json:"errors,required"`
	Messages []string                             `json:"messages,required"`
	Result   interface{}                          `json:"result,required"`
	// Whether the API call was successful
	Success AccountR2BucketDeleteResponseSuccess `json:"success,required"`
	JSON    accountR2BucketDeleteResponseJSON    `json:"-"`
}

// accountR2BucketDeleteResponseJSON contains the JSON metadata for the struct
// [AccountR2BucketDeleteResponse]
type accountR2BucketDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountR2BucketDeleteResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    accountR2BucketDeleteResponseErrorJSON `json:"-"`
}

// accountR2BucketDeleteResponseErrorJSON contains the JSON metadata for the struct
// [AccountR2BucketDeleteResponseError]
type accountR2BucketDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountR2BucketDeleteResponseSuccess bool

const (
	AccountR2BucketDeleteResponseSuccessTrue AccountR2BucketDeleteResponseSuccess = true
)

type AccountR2BucketR2BucketNewBucketResponse struct {
	Errors   []AccountR2BucketR2BucketNewBucketResponseError `json:"errors"`
	Messages []string                                        `json:"messages"`
	// A single R2 bucket
	Result AccountR2BucketR2BucketNewBucketResponseResult `json:"result"`
	// Whether the API call was successful
	Success AccountR2BucketR2BucketNewBucketResponseSuccess `json:"success"`
	JSON    accountR2BucketR2BucketNewBucketResponseJSON    `json:"-"`
}

// accountR2BucketR2BucketNewBucketResponseJSON contains the JSON metadata for the
// struct [AccountR2BucketR2BucketNewBucketResponse]
type accountR2BucketR2BucketNewBucketResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketR2BucketNewBucketResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountR2BucketR2BucketNewBucketResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountR2BucketR2BucketNewBucketResponseErrorJSON `json:"-"`
}

// accountR2BucketR2BucketNewBucketResponseErrorJSON contains the JSON metadata for
// the struct [AccountR2BucketR2BucketNewBucketResponseError]
type accountR2BucketR2BucketNewBucketResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketR2BucketNewBucketResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A single R2 bucket
type AccountR2BucketR2BucketNewBucketResponseResult struct {
	// Creation timestamp
	CreationDate string `json:"creation_date"`
	// Location of the bucket
	Location AccountR2BucketR2BucketNewBucketResponseResultLocation `json:"location"`
	// Name of the bucket
	Name string                                             `json:"name"`
	JSON accountR2BucketR2BucketNewBucketResponseResultJSON `json:"-"`
}

// accountR2BucketR2BucketNewBucketResponseResultJSON contains the JSON metadata
// for the struct [AccountR2BucketR2BucketNewBucketResponseResult]
type accountR2BucketR2BucketNewBucketResponseResultJSON struct {
	CreationDate apijson.Field
	Location     apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountR2BucketR2BucketNewBucketResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Location of the bucket
type AccountR2BucketR2BucketNewBucketResponseResultLocation string

const (
	AccountR2BucketR2BucketNewBucketResponseResultLocationApac AccountR2BucketR2BucketNewBucketResponseResultLocation = "apac"
	AccountR2BucketR2BucketNewBucketResponseResultLocationEeur AccountR2BucketR2BucketNewBucketResponseResultLocation = "eeur"
	AccountR2BucketR2BucketNewBucketResponseResultLocationEnam AccountR2BucketR2BucketNewBucketResponseResultLocation = "enam"
	AccountR2BucketR2BucketNewBucketResponseResultLocationWeur AccountR2BucketR2BucketNewBucketResponseResultLocation = "weur"
	AccountR2BucketR2BucketNewBucketResponseResultLocationWnam AccountR2BucketR2BucketNewBucketResponseResultLocation = "wnam"
)

// Whether the API call was successful
type AccountR2BucketR2BucketNewBucketResponseSuccess bool

const (
	AccountR2BucketR2BucketNewBucketResponseSuccessTrue AccountR2BucketR2BucketNewBucketResponseSuccess = true
)

type AccountR2BucketR2BucketListBucketsResponse struct {
	Errors     []AccountR2BucketR2BucketListBucketsResponseError    `json:"errors"`
	Messages   []string                                             `json:"messages"`
	Result     []AccountR2BucketR2BucketListBucketsResponseResult   `json:"result"`
	ResultInfo AccountR2BucketR2BucketListBucketsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountR2BucketR2BucketListBucketsResponseSuccess `json:"success"`
	JSON    accountR2BucketR2BucketListBucketsResponseJSON    `json:"-"`
}

// accountR2BucketR2BucketListBucketsResponseJSON contains the JSON metadata for
// the struct [AccountR2BucketR2BucketListBucketsResponse]
type accountR2BucketR2BucketListBucketsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketR2BucketListBucketsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountR2BucketR2BucketListBucketsResponseError struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountR2BucketR2BucketListBucketsResponseErrorJSON `json:"-"`
}

// accountR2BucketR2BucketListBucketsResponseErrorJSON contains the JSON metadata
// for the struct [AccountR2BucketR2BucketListBucketsResponseError]
type accountR2BucketR2BucketListBucketsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketR2BucketListBucketsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A single R2 bucket
type AccountR2BucketR2BucketListBucketsResponseResult struct {
	// Creation timestamp
	CreationDate string `json:"creation_date"`
	// Location of the bucket
	Location AccountR2BucketR2BucketListBucketsResponseResultLocation `json:"location"`
	// Name of the bucket
	Name string                                               `json:"name"`
	JSON accountR2BucketR2BucketListBucketsResponseResultJSON `json:"-"`
}

// accountR2BucketR2BucketListBucketsResponseResultJSON contains the JSON metadata
// for the struct [AccountR2BucketR2BucketListBucketsResponseResult]
type accountR2BucketR2BucketListBucketsResponseResultJSON struct {
	CreationDate apijson.Field
	Location     apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountR2BucketR2BucketListBucketsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Location of the bucket
type AccountR2BucketR2BucketListBucketsResponseResultLocation string

const (
	AccountR2BucketR2BucketListBucketsResponseResultLocationApac AccountR2BucketR2BucketListBucketsResponseResultLocation = "apac"
	AccountR2BucketR2BucketListBucketsResponseResultLocationEeur AccountR2BucketR2BucketListBucketsResponseResultLocation = "eeur"
	AccountR2BucketR2BucketListBucketsResponseResultLocationEnam AccountR2BucketR2BucketListBucketsResponseResultLocation = "enam"
	AccountR2BucketR2BucketListBucketsResponseResultLocationWeur AccountR2BucketR2BucketListBucketsResponseResultLocation = "weur"
	AccountR2BucketR2BucketListBucketsResponseResultLocationWnam AccountR2BucketR2BucketListBucketsResponseResultLocation = "wnam"
)

type AccountR2BucketR2BucketListBucketsResponseResultInfo struct {
	// A continuation token that should be used to fetch the next page of results
	Cursor string `json:"cursor"`
	// Maximum number of results on this page
	PerPage float64                                                  `json:"per_page"`
	JSON    accountR2BucketR2BucketListBucketsResponseResultInfoJSON `json:"-"`
}

// accountR2BucketR2BucketListBucketsResponseResultInfoJSON contains the JSON
// metadata for the struct [AccountR2BucketR2BucketListBucketsResponseResultInfo]
type accountR2BucketR2BucketListBucketsResponseResultInfoJSON struct {
	Cursor      apijson.Field
	PerPage     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountR2BucketR2BucketListBucketsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountR2BucketR2BucketListBucketsResponseSuccess bool

const (
	AccountR2BucketR2BucketListBucketsResponseSuccessTrue AccountR2BucketR2BucketListBucketsResponseSuccess = true
)

type AccountR2BucketR2BucketNewBucketParams struct {
	// Name of the bucket
	Name param.Field[string] `json:"name,required"`
	// Location of the bucket
	LocationHint param.Field[AccountR2BucketR2BucketNewBucketParamsLocationHint] `json:"locationHint"`
}

func (r AccountR2BucketR2BucketNewBucketParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Location of the bucket
type AccountR2BucketR2BucketNewBucketParamsLocationHint string

const (
	AccountR2BucketR2BucketNewBucketParamsLocationHintApac AccountR2BucketR2BucketNewBucketParamsLocationHint = "apac"
	AccountR2BucketR2BucketNewBucketParamsLocationHintEeur AccountR2BucketR2BucketNewBucketParamsLocationHint = "eeur"
	AccountR2BucketR2BucketNewBucketParamsLocationHintEnam AccountR2BucketR2BucketNewBucketParamsLocationHint = "enam"
	AccountR2BucketR2BucketNewBucketParamsLocationHintWeur AccountR2BucketR2BucketNewBucketParamsLocationHint = "weur"
	AccountR2BucketR2BucketNewBucketParamsLocationHintWnam AccountR2BucketR2BucketNewBucketParamsLocationHint = "wnam"
)

type AccountR2BucketR2BucketListBucketsParams struct {
	// Pagination cursor received during the last List Buckets call. R2 buckets are
	// paginated using cursors instead of page numbers.
	Cursor param.Field[string] `query:"cursor"`
	// Direction to order buckets
	Direction param.Field[AccountR2BucketR2BucketListBucketsParamsDirection] `query:"direction"`
	// Bucket names to filter by. Only buckets with this phrase in their name will be
	// returned.
	NameContains param.Field[string] `query:"name_contains"`
	// Field to order buckets by
	Order param.Field[AccountR2BucketR2BucketListBucketsParamsOrder] `query:"order"`
	// Maximum number of buckets to return in a single call
	PerPage param.Field[float64] `query:"per_page"`
	// Bucket name to start searching after. Buckets are ordered lexicographically.
	StartAfter param.Field[string] `query:"start_after"`
}

// URLQuery serializes [AccountR2BucketR2BucketListBucketsParams]'s query
// parameters as `url.Values`.
func (r AccountR2BucketR2BucketListBucketsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order buckets
type AccountR2BucketR2BucketListBucketsParamsDirection string

const (
	AccountR2BucketR2BucketListBucketsParamsDirectionAsc  AccountR2BucketR2BucketListBucketsParamsDirection = "asc"
	AccountR2BucketR2BucketListBucketsParamsDirectionDesc AccountR2BucketR2BucketListBucketsParamsDirection = "desc"
)

// Field to order buckets by
type AccountR2BucketR2BucketListBucketsParamsOrder string

const (
	AccountR2BucketR2BucketListBucketsParamsOrderName AccountR2BucketR2BucketListBucketsParamsOrder = "name"
)
