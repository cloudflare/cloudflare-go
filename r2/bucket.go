// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// BucketService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBucketService] method instead.
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
func (r *BucketService) New(ctx context.Context, params BucketNewParams, opts ...option.RequestOption) (res *Bucket, err error) {
	var env BucketNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all R2 buckets on your account
func (r *BucketService) List(ctx context.Context, params BucketListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Bucket], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists all R2 buckets on your account
func (r *BucketService) ListAutoPaging(ctx context.Context, params BucketListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Bucket] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Deletes an existing R2 bucket.
func (r *BucketService) Delete(ctx context.Context, bucketName string, body BucketDeleteParams, opts ...option.RequestOption) (res *BucketDeleteResponse, err error) {
	var env BucketDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s", body.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets metadata for an existing R2 bucket.
func (r *BucketService) Get(ctx context.Context, bucketName string, query BucketGetParams, opts ...option.RequestOption) (res *Bucket, err error) {
	var env BucketGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s", query.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A single R2 bucket
type Bucket struct {
	// Creation timestamp
	CreationDate string `json:"creation_date"`
	// Location of the bucket
	Location BucketLocation `json:"location"`
	// Name of the bucket
	Name string `json:"name"`
	// Storage class for newly uploaded objects, unless specified otherwise.
	StorageClass BucketStorageClass `json:"storage_class"`
	JSON         bucketJSON         `json:"-"`
}

// bucketJSON contains the JSON metadata for the struct [Bucket]
type bucketJSON struct {
	CreationDate apijson.Field
	Location     apijson.Field
	Name         apijson.Field
	StorageClass apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketJSON) RawJSON() string {
	return r.raw
}

// Location of the bucket
type BucketLocation string

const (
	BucketLocationApac BucketLocation = "apac"
	BucketLocationEeur BucketLocation = "eeur"
	BucketLocationEnam BucketLocation = "enam"
	BucketLocationWeur BucketLocation = "weur"
	BucketLocationWnam BucketLocation = "wnam"
)

func (r BucketLocation) IsKnown() bool {
	switch r {
	case BucketLocationApac, BucketLocationEeur, BucketLocationEnam, BucketLocationWeur, BucketLocationWnam:
		return true
	}
	return false
}

// Storage class for newly uploaded objects, unless specified otherwise.
type BucketStorageClass string

const (
	BucketStorageClassStandard         BucketStorageClass = "Standard"
	BucketStorageClassInfrequentAccess BucketStorageClass = "InfrequentAccess"
)

func (r BucketStorageClass) IsKnown() bool {
	switch r {
	case BucketStorageClassStandard, BucketStorageClassInfrequentAccess:
		return true
	}
	return false
}

type BucketDeleteResponse = interface{}

type BucketNewParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// Name of the bucket
	Name param.Field[string] `json:"name,required"`
	// Location of the bucket
	LocationHint param.Field[BucketNewParamsLocationHint] `json:"locationHint"`
	// Storage class for newly uploaded objects, unless specified otherwise.
	StorageClass param.Field[BucketNewParamsStorageClass] `json:"storageClass"`
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

func (r BucketNewParamsLocationHint) IsKnown() bool {
	switch r {
	case BucketNewParamsLocationHintApac, BucketNewParamsLocationHintEeur, BucketNewParamsLocationHintEnam, BucketNewParamsLocationHintWeur, BucketNewParamsLocationHintWnam:
		return true
	}
	return false
}

// Storage class for newly uploaded objects, unless specified otherwise.
type BucketNewParamsStorageClass string

const (
	BucketNewParamsStorageClassStandard         BucketNewParamsStorageClass = "Standard"
	BucketNewParamsStorageClassInfrequentAccess BucketNewParamsStorageClass = "InfrequentAccess"
)

func (r BucketNewParamsStorageClass) IsKnown() bool {
	switch r {
	case BucketNewParamsStorageClassStandard, BucketNewParamsStorageClassInfrequentAccess:
		return true
	}
	return false
}

type BucketNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []string              `json:"messages,required"`
	// A single R2 bucket
	Result Bucket `json:"result,required"`
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

// Whether the API call was successful
type BucketNewResponseEnvelopeSuccess bool

const (
	BucketNewResponseEnvelopeSuccessTrue BucketNewResponseEnvelopeSuccess = true
)

func (r BucketNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Direction to order buckets
type BucketListParamsDirection string

const (
	BucketListParamsDirectionAsc  BucketListParamsDirection = "asc"
	BucketListParamsDirectionDesc BucketListParamsDirection = "desc"
)

func (r BucketListParamsDirection) IsKnown() bool {
	switch r {
	case BucketListParamsDirectionAsc, BucketListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order buckets by
type BucketListParamsOrder string

const (
	BucketListParamsOrderName BucketListParamsOrder = "name"
)

func (r BucketListParamsOrder) IsKnown() bool {
	switch r {
	case BucketListParamsOrderName:
		return true
	}
	return false
}

type BucketDeleteParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type BucketDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []string              `json:"messages,required"`
	Result   BucketDeleteResponse  `json:"result,required"`
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

// Whether the API call was successful
type BucketDeleteResponseEnvelopeSuccess bool

const (
	BucketDeleteResponseEnvelopeSuccessTrue BucketDeleteResponseEnvelopeSuccess = true
)

func (r BucketDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BucketGetParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type BucketGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []string              `json:"messages,required"`
	// A single R2 bucket
	Result Bucket `json:"result,required"`
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

// Whether the API call was successful
type BucketGetResponseEnvelopeSuccess bool

const (
	BucketGetResponseEnvelopeSuccessTrue BucketGetResponseEnvelopeSuccess = true
)

func (r BucketGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
