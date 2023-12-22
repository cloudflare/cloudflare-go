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

// Deletes an existing R2 bucket.
func (r *AccountR2BucketService) Delete(ctx context.Context, accountIdentifier string, bucketName string, opts ...option.RequestOption) (res *R2SingleBucketOperationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s", accountIdentifier, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a new R2 bucket.
func (r *AccountR2BucketService) R2BucketNewBucket(ctx context.Context, accountIdentifier string, body AccountR2BucketR2BucketNewBucketParams, opts ...option.RequestOption) (res *R2SingleBucketOperationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/r2/buckets", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all R2 buckets on your account
func (r *AccountR2BucketService) R2BucketListBuckets(ctx context.Context, accountIdentifier string, query AccountR2BucketR2BucketListBucketsParams, opts ...option.RequestOption) (res *R2SingleBucketOperationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/r2/buckets", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type R2SingleBucketOperationResponse struct {
	Errors   []R2SingleBucketOperationResponseError   `json:"errors"`
	Messages []R2SingleBucketOperationResponseMessage `json:"messages"`
	Result   interface{}                              `json:"result"`
	// Whether the API call was successful
	Success R2SingleBucketOperationResponseSuccess `json:"success"`
	JSON    r2SingleBucketOperationResponseJSON    `json:"-"`
}

// r2SingleBucketOperationResponseJSON contains the JSON metadata for the struct
// [R2SingleBucketOperationResponse]
type r2SingleBucketOperationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2SingleBucketOperationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type R2SingleBucketOperationResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    r2SingleBucketOperationResponseErrorJSON `json:"-"`
}

// r2SingleBucketOperationResponseErrorJSON contains the JSON metadata for the
// struct [R2SingleBucketOperationResponseError]
type r2SingleBucketOperationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2SingleBucketOperationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type R2SingleBucketOperationResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    r2SingleBucketOperationResponseMessageJSON `json:"-"`
}

// r2SingleBucketOperationResponseMessageJSON contains the JSON metadata for the
// struct [R2SingleBucketOperationResponseMessage]
type r2SingleBucketOperationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2SingleBucketOperationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type R2SingleBucketOperationResponseSuccess bool

const (
	R2SingleBucketOperationResponseSuccessTrue R2SingleBucketOperationResponseSuccess = true
)

type AccountR2BucketR2BucketNewBucketParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountR2BucketR2BucketNewBucketParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

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
