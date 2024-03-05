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
func (r *R2BucketService) New(ctx context.Context, params R2BucketNewParams, opts ...option.RequestOption) (res *R2Bucket, err error) {
	opts = append(r.Options[:], opts...)
	var env R2BucketNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/r2/buckets", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all R2 buckets on your account
func (r *R2BucketService) List(ctx context.Context, params R2BucketListParams, opts ...option.RequestOption) (res *[]R2Bucket, err error) {
	opts = append(r.Options[:], opts...)
	var env R2BucketListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/r2/buckets", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an existing R2 bucket.
func (r *R2BucketService) Delete(ctx context.Context, bucketName string, body R2BucketDeleteParams, opts ...option.RequestOption) (res *R2BucketDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env R2BucketDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s", body.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets metadata for an existing R2 bucket.
func (r *R2BucketService) Get(ctx context.Context, bucketName string, query R2BucketGetParams, opts ...option.RequestOption) (res *R2Bucket, err error) {
	opts = append(r.Options[:], opts...)
	var env R2BucketGetResponseEnvelope
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

// Location of the bucket
type R2BucketLocation string

const (
	R2BucketLocationApac R2BucketLocation = "apac"
	R2BucketLocationEeur R2BucketLocation = "eeur"
	R2BucketLocationEnam R2BucketLocation = "enam"
	R2BucketLocationWeur R2BucketLocation = "weur"
	R2BucketLocationWnam R2BucketLocation = "wnam"
)

type R2BucketDeleteResponse = interface{}

type R2BucketNewParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
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

type R2BucketNewResponseEnvelope struct {
	Errors   []R2BucketNewResponseEnvelopeErrors `json:"errors,required"`
	Messages []string                            `json:"messages,required"`
	// A single R2 bucket
	Result R2Bucket `json:"result,required"`
	// Whether the API call was successful
	Success R2BucketNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    r2BucketNewResponseEnvelopeJSON    `json:"-"`
}

// r2BucketNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [R2BucketNewResponseEnvelope]
type r2BucketNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2BucketNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type R2BucketNewResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    r2BucketNewResponseEnvelopeErrorsJSON `json:"-"`
}

// r2BucketNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [R2BucketNewResponseEnvelopeErrors]
type r2BucketNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2BucketNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type R2BucketNewResponseEnvelopeSuccess bool

const (
	R2BucketNewResponseEnvelopeSuccessTrue R2BucketNewResponseEnvelopeSuccess = true
)

type R2BucketListParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
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

type R2BucketListResponseEnvelope struct {
	Errors   []R2BucketListResponseEnvelopeErrors `json:"errors,required"`
	Messages []string                             `json:"messages,required"`
	Result   []R2Bucket                           `json:"result,required"`
	// Whether the API call was successful
	Success    R2BucketListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo R2BucketListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       r2BucketListResponseEnvelopeJSON       `json:"-"`
}

// r2BucketListResponseEnvelopeJSON contains the JSON metadata for the struct
// [R2BucketListResponseEnvelope]
type r2BucketListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2BucketListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type R2BucketListResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    r2BucketListResponseEnvelopeErrorsJSON `json:"-"`
}

// r2BucketListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [R2BucketListResponseEnvelopeErrors]
type r2BucketListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2BucketListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type R2BucketListResponseEnvelopeSuccess bool

const (
	R2BucketListResponseEnvelopeSuccessTrue R2BucketListResponseEnvelopeSuccess = true
)

type R2BucketListResponseEnvelopeResultInfo struct {
	// A continuation token that should be used to fetch the next page of results
	Cursor string `json:"cursor"`
	// Maximum number of results on this page
	PerPage float64                                    `json:"per_page"`
	JSON    r2BucketListResponseEnvelopeResultInfoJSON `json:"-"`
}

// r2BucketListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [R2BucketListResponseEnvelopeResultInfo]
type r2BucketListResponseEnvelopeResultInfoJSON struct {
	Cursor      apijson.Field
	PerPage     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2BucketListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type R2BucketDeleteParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type R2BucketDeleteResponseEnvelope struct {
	Errors   []R2BucketDeleteResponseEnvelopeErrors `json:"errors,required"`
	Messages []string                               `json:"messages,required"`
	Result   R2BucketDeleteResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success R2BucketDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    r2BucketDeleteResponseEnvelopeJSON    `json:"-"`
}

// r2BucketDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [R2BucketDeleteResponseEnvelope]
type r2BucketDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2BucketDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type R2BucketDeleteResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    r2BucketDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// r2BucketDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [R2BucketDeleteResponseEnvelopeErrors]
type r2BucketDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2BucketDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type R2BucketDeleteResponseEnvelopeSuccess bool

const (
	R2BucketDeleteResponseEnvelopeSuccessTrue R2BucketDeleteResponseEnvelopeSuccess = true
)

type R2BucketGetParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type R2BucketGetResponseEnvelope struct {
	Errors   []R2BucketGetResponseEnvelopeErrors `json:"errors,required"`
	Messages []string                            `json:"messages,required"`
	// A single R2 bucket
	Result R2Bucket `json:"result,required"`
	// Whether the API call was successful
	Success R2BucketGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    r2BucketGetResponseEnvelopeJSON    `json:"-"`
}

// r2BucketGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [R2BucketGetResponseEnvelope]
type r2BucketGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2BucketGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type R2BucketGetResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    r2BucketGetResponseEnvelopeErrorsJSON `json:"-"`
}

// r2BucketGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [R2BucketGetResponseEnvelopeErrors]
type r2BucketGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2BucketGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type R2BucketGetResponseEnvelopeSuccess bool

const (
	R2BucketGetResponseEnvelopeSuccessTrue R2BucketGetResponseEnvelopeSuccess = true
)
