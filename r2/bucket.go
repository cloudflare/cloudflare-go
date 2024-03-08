// File generated from our OpenAPI spec by Stainless.

package r2

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

// BucketService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewBucketService] method instead.
type BucketService struct {
	Options []option.RequestOption
}

// NewBucketService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBucketService(opts ...option.RequestOption) (r *BucketService) {
	r = &BucketService{}
	r.Options = opts
	return
}

// Creates a new R2 bucket.
func (r *BucketService) New(ctx context.Context, params BucketNewParams, opts ...option.RequestOption) (res *R2Bucket, err error) {
	opts = append(r.Options[:], opts...)
	var env BucketNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/r2/buckets", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all R2 buckets on your account
func (r *BucketService) List(ctx context.Context, params BucketListParams, opts ...option.RequestOption) (res *[]R2Bucket, err error) {
	opts = append(r.Options[:], opts...)
	var env BucketListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/r2/buckets", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an existing R2 bucket.
func (r *BucketService) Delete(ctx context.Context, bucketName string, body BucketDeleteParams, opts ...option.RequestOption) (res *BucketDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env BucketDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s", body.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets metadata for an existing R2 bucket.
func (r *BucketService) Get(ctx context.Context, bucketName string, query BucketGetParams, opts ...option.RequestOption) (res *R2Bucket, err error) {
	opts = append(r.Options[:], opts...)
	var env BucketGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s", query.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A single R2 bucket
type R2Bucket struct {
	// Creation timestamp
	CreationDate string `json:"creation_date"`
	// Location of the bucket
	Location R2BucketLocation `json:"location"`
	// Name of the bucket
	Name string       `json:"name"`
	JSON r2BucketJSON `json:"-"`
}

// r2BucketJSON contains the JSON metadata for the struct [R2Bucket]
type r2BucketJSON struct {
	CreationDate apijson.Field
	Location     apijson.Field
	Name         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *R2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2BucketJSON) RawJSON() string {
	return r.raw
}

// Location of the bucket
type R2BucketLocation string

const (
	R2BucketLocationApac R2BucketLocation = "apac"
	R2BucketLocationEeur R2BucketLocation = "eeur"
	R2BucketLocationEnam R2BucketLocation = "enam"
	R2BucketLocationWeur R2BucketLocation = "weur"
	R2BucketLocationWnam R2BucketLocation = "wnam"
)

type BucketDeleteResponse = interface{}

type BucketNewParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// Name of the bucket
	Name param.Field[string] `json:"name,required"`
	// Location of the bucket
	LocationHint param.Field[BucketNewParamsLocationHint] `json:"locationHint"`
}

func (r BucketNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Location of the bucket
type BucketNewParamsLocationHint string

const (
	BucketNewParamsLocationHintApac BucketNewParamsLocationHint = "apac"
	BucketNewParamsLocationHintEeur BucketNewParamsLocationHint = "eeur"
	BucketNewParamsLocationHintEnam BucketNewParamsLocationHint = "enam"
	BucketNewParamsLocationHintWeur BucketNewParamsLocationHint = "weur"
	BucketNewParamsLocationHintWnam BucketNewParamsLocationHint = "wnam"
)

type BucketNewResponseEnvelope struct {
	Errors   []BucketNewResponseEnvelopeErrors `json:"errors,required"`
	Messages []string                          `json:"messages,required"`
	// A single R2 bucket
	Result R2Bucket `json:"result,required"`
	// Whether the API call was successful
	Success BucketNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketNewResponseEnvelopeJSON    `json:"-"`
}

// bucketNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [BucketNewResponseEnvelope]
type bucketNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BucketNewResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    bucketNewResponseEnvelopeErrorsJSON `json:"-"`
}

// bucketNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [BucketNewResponseEnvelopeErrors]
type bucketNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type BucketNewResponseEnvelopeSuccess bool

const (
	BucketNewResponseEnvelopeSuccessTrue BucketNewResponseEnvelopeSuccess = true
)

type BucketListParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// Pagination cursor received during the last List Buckets call. R2 buckets are
	// paginated using cursors instead of page numbers.
	Cursor param.Field[string] `query:"cursor"`
	// Direction to order buckets
	Direction param.Field[BucketListParamsDirection] `query:"direction"`
	// Bucket names to filter by. Only buckets with this phrase in their name will be
	// returned.
	NameContains param.Field[string] `query:"name_contains"`
	// Field to order buckets by
	Order param.Field[BucketListParamsOrder] `query:"order"`
	// Maximum number of buckets to return in a single call
	PerPage param.Field[float64] `query:"per_page"`
	// Bucket name to start searching after. Buckets are ordered lexicographically.
	StartAfter param.Field[string] `query:"start_after"`
}

// URLQuery serializes [BucketListParams]'s query parameters as `url.Values`.
func (r BucketListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order buckets
type BucketListParamsDirection string

const (
	BucketListParamsDirectionAsc  BucketListParamsDirection = "asc"
	BucketListParamsDirectionDesc BucketListParamsDirection = "desc"
)

// Field to order buckets by
type BucketListParamsOrder string

const (
	BucketListParamsOrderName BucketListParamsOrder = "name"
)

type BucketListResponseEnvelope struct {
	Errors   []BucketListResponseEnvelopeErrors `json:"errors,required"`
	Messages []string                           `json:"messages,required"`
	Result   []R2Bucket                         `json:"result,required"`
	// Whether the API call was successful
	Success    BucketListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo BucketListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       bucketListResponseEnvelopeJSON       `json:"-"`
}

// bucketListResponseEnvelopeJSON contains the JSON metadata for the struct
// [BucketListResponseEnvelope]
type bucketListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BucketListResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    bucketListResponseEnvelopeErrorsJSON `json:"-"`
}

// bucketListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [BucketListResponseEnvelopeErrors]
type bucketListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type BucketListResponseEnvelopeSuccess bool

const (
	BucketListResponseEnvelopeSuccessTrue BucketListResponseEnvelopeSuccess = true
)

type BucketListResponseEnvelopeResultInfo struct {
	// A continuation token that should be used to fetch the next page of results
	Cursor string `json:"cursor"`
	// Maximum number of results on this page
	PerPage float64                                  `json:"per_page"`
	JSON    bucketListResponseEnvelopeResultInfoJSON `json:"-"`
}

// bucketListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [BucketListResponseEnvelopeResultInfo]
type bucketListResponseEnvelopeResultInfoJSON struct {
	Cursor      apijson.Field
	PerPage     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type BucketDeleteParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type BucketDeleteResponseEnvelope struct {
	Errors   []BucketDeleteResponseEnvelopeErrors `json:"errors,required"`
	Messages []string                             `json:"messages,required"`
	Result   BucketDeleteResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success BucketDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketDeleteResponseEnvelopeJSON    `json:"-"`
}

// bucketDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [BucketDeleteResponseEnvelope]
type bucketDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BucketDeleteResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    bucketDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// bucketDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [BucketDeleteResponseEnvelopeErrors]
type bucketDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type BucketDeleteResponseEnvelopeSuccess bool

const (
	BucketDeleteResponseEnvelopeSuccessTrue BucketDeleteResponseEnvelopeSuccess = true
)

type BucketGetParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type BucketGetResponseEnvelope struct {
	Errors   []BucketGetResponseEnvelopeErrors `json:"errors,required"`
	Messages []string                          `json:"messages,required"`
	// A single R2 bucket
	Result R2Bucket `json:"result,required"`
	// Whether the API call was successful
	Success BucketGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketGetResponseEnvelopeJSON    `json:"-"`
}

// bucketGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [BucketGetResponseEnvelope]
type bucketGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BucketGetResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    bucketGetResponseEnvelopeErrorsJSON `json:"-"`
}

// bucketGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [BucketGetResponseEnvelopeErrors]
type bucketGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type BucketGetResponseEnvelopeSuccess bool

const (
	BucketGetResponseEnvelopeSuccessTrue BucketGetResponseEnvelopeSuccess = true
)
