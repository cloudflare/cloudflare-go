// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v7/shared"
)

// BucketObjectService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBucketObjectService] method instead.
type BucketObjectService struct {
	Options []option.RequestOption
}

// NewBucketObjectService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBucketObjectService(opts ...option.RequestOption) (r *BucketObjectService) {
	r = &BucketObjectService{}
	r.Options = opts
	return
}

// Lists objects in an R2 bucket. Returns object metadata including key, size,
// etag, last modified date, HTTP metadata, and custom metadata.
//
// For most workloads, we recommend using R2's
// [S3-compatible API](https://developers.cloudflare.com/r2/api/s3/api/) or a
// [Worker with an R2 binding](https://developers.cloudflare.com/r2/api/workers/workers-api-reference/)
// instead.
func (r *BucketObjectService) List(ctx context.Context, bucketName string, params BucketObjectListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[BucketObjectListResponse], err error) {
	var raw *http.Response
	if params.Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%v", params.Jurisdiction)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/objects", params.AccountID, bucketName)
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

// Lists objects in an R2 bucket. Returns object metadata including key, size,
// etag, last modified date, HTTP metadata, and custom metadata.
//
// For most workloads, we recommend using R2's
// [S3-compatible API](https://developers.cloudflare.com/r2/api/s3/api/) or a
// [Worker with an R2 binding](https://developers.cloudflare.com/r2/api/workers/workers-api-reference/)
// instead.
func (r *BucketObjectService) ListAutoPaging(ctx context.Context, bucketName string, params BucketObjectListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[BucketObjectListResponse] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, bucketName, params, opts...))
}

// Deletes an object from an R2 bucket.
//
// For most workloads, we recommend using R2's
// [S3-compatible API](https://developers.cloudflare.com/r2/api/s3/api/) or a
// [Worker with an R2 binding](https://developers.cloudflare.com/r2/api/workers/workers-api-reference/)
// instead.
func (r *BucketObjectService) Delete(ctx context.Context, bucketName string, objectKey string, params BucketObjectDeleteParams, opts ...option.RequestOption) (res *BucketObjectDeleteResponse, err error) {
	var env BucketObjectDeleteResponseEnvelope
	if params.Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%v", params.Jurisdiction)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return nil, err
	}
	if objectKey == "" {
		err = errors.New("missing required object_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/objects/%s", params.AccountID, bucketName, objectKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieves an object from an R2 bucket. Returns the object body along with
// metadata headers.
//
// For most workloads, we recommend using R2's
// [S3-compatible API](https://developers.cloudflare.com/r2/api/s3/api/) or a
// [Worker with an R2 binding](https://developers.cloudflare.com/r2/api/workers/workers-api-reference/)
// instead.
func (r *BucketObjectService) Get(ctx context.Context, bucketName string, objectKey string, params BucketObjectGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if params.Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%v", params.Jurisdiction)))
	}
	if params.IfModifiedSince.Present {
		opts = append(opts, option.WithHeader("If-Modified-Since", fmt.Sprintf("%v", params.IfModifiedSince)))
	}
	if params.IfNoneMatch.Present {
		opts = append(opts, option.WithHeader("If-None-Match", fmt.Sprintf("%v", params.IfNoneMatch)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return nil, err
	}
	if objectKey == "" {
		err = errors.New("missing required object_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/objects/%s", params.AccountID, bucketName, objectKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Uploads an object to an R2 bucket. The object body is provided as the request
// body. Returns metadata about the uploaded object.
//
// The maximum upload size for this endpoint is 300 MB. For most workloads, we
// recommend using R2's
// [S3-compatible API](https://developers.cloudflare.com/r2/api/s3/api/) or a
// [Worker with an R2 binding](https://developers.cloudflare.com/r2/api/workers/workers-api-reference/)
// instead.
func (r *BucketObjectService) Upload(ctx context.Context, bucketName string, objectKey string, body io.Reader, params BucketObjectUploadParams, opts ...option.RequestOption) (res *BucketObjectUploadResponse, err error) {
	var env BucketObjectUploadResponseEnvelope
	if params.Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%v", params.Jurisdiction)))
	}
	if params.CfR2StorageClass.Present {
		opts = append(opts, option.WithHeader("cf-r2-storage-class", fmt.Sprintf("%v", params.CfR2StorageClass)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithRequestBody("application/octet-stream", body)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return nil, err
	}
	if objectKey == "" {
		err = errors.New("missing required object_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/objects/%s", params.AccountID, bucketName, objectKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Metadata for an R2 object.
type BucketObjectListResponse struct {
	// Custom metadata key-value pairs associated with the object.
	CustomMetadata map[string]string `json:"custom_metadata"`
	// The entity tag for the object. In JSON list/get responses this is the raw hex
	// digest (without surrounding quotes). The HTTP `ETag` response header on Get
	// Object follows RFC 7232 and IS wrapped in surrounding double-quotes.
	Etag string `json:"etag"`
	// HTTP metadata associated with an R2 object.
	HTTPMetadata BucketObjectListResponseHTTPMetadata `json:"http_metadata"`
	// The object key (name).
	Key string `json:"key"`
	// The date and time the object was last modified.
	LastModified time.Time `json:"last_modified" format:"date-time"`
	// The size of the object in bytes.
	Size int64 `json:"size"`
	// Whether the object is encrypted with a customer-supplied encryption key.
	Ssec bool `json:"ssec"`
	// Storage class for newly uploaded objects, unless specified otherwise.
	StorageClass BucketObjectListResponseStorageClass `json:"storage_class"`
	JSON         bucketObjectListResponseJSON         `json:"-"`
}

// bucketObjectListResponseJSON contains the JSON metadata for the struct
// [BucketObjectListResponse]
type bucketObjectListResponseJSON struct {
	CustomMetadata apijson.Field
	Etag           apijson.Field
	HTTPMetadata   apijson.Field
	Key            apijson.Field
	LastModified   apijson.Field
	Size           apijson.Field
	Ssec           apijson.Field
	StorageClass   apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BucketObjectListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketObjectListResponseJSON) RawJSON() string {
	return r.raw
}

// HTTP metadata associated with an R2 object.
type BucketObjectListResponseHTTPMetadata struct {
	// Specifies caching behavior for the object.
	CacheControl string `json:"cacheControl"`
	// The date and time at which the object's cache entry expires.
	CacheExpiry time.Time `json:"cacheExpiry" format:"date-time"`
	// Specifies presentational information for the object.
	ContentDisposition string `json:"contentDisposition"`
	// Specifies the content encoding applied to the object.
	ContentEncoding string `json:"contentEncoding"`
	// The language of the object content.
	ContentLanguage string `json:"contentLanguage"`
	// The MIME type of the object.
	ContentType string                                   `json:"contentType"`
	JSON        bucketObjectListResponseHTTPMetadataJSON `json:"-"`
}

// bucketObjectListResponseHTTPMetadataJSON contains the JSON metadata for the
// struct [BucketObjectListResponseHTTPMetadata]
type bucketObjectListResponseHTTPMetadataJSON struct {
	CacheControl       apijson.Field
	CacheExpiry        apijson.Field
	ContentDisposition apijson.Field
	ContentEncoding    apijson.Field
	ContentLanguage    apijson.Field
	ContentType        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *BucketObjectListResponseHTTPMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketObjectListResponseHTTPMetadataJSON) RawJSON() string {
	return r.raw
}

// Storage class for newly uploaded objects, unless specified otherwise.
type BucketObjectListResponseStorageClass string

const (
	BucketObjectListResponseStorageClassStandard         BucketObjectListResponseStorageClass = "Standard"
	BucketObjectListResponseStorageClassInfrequentAccess BucketObjectListResponseStorageClass = "InfrequentAccess"
)

func (r BucketObjectListResponseStorageClass) IsKnown() bool {
	switch r {
	case BucketObjectListResponseStorageClassStandard, BucketObjectListResponseStorageClassInfrequentAccess:
		return true
	}
	return false
}

// Result of a successful object deletion.
type BucketObjectDeleteResponse struct {
	// The key (name) of the deleted object.
	Key  string                         `json:"key"`
	JSON bucketObjectDeleteResponseJSON `json:"-"`
}

// bucketObjectDeleteResponseJSON contains the JSON metadata for the struct
// [BucketObjectDeleteResponse]
type bucketObjectDeleteResponseJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketObjectDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketObjectDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Result of a successful object upload.
type BucketObjectUploadResponse struct {
	// The entity tag for the uploaded object.
	Etag string `json:"etag"`
	// The key (name) of the uploaded object.
	Key string `json:"key"`
	// The size of the uploaded object in bytes (as a string).
	Size string `json:"size"`
	// Storage class for newly uploaded objects, unless specified otherwise.
	StorageClass BucketObjectUploadResponseStorageClass `json:"storage_class"`
	// The date and time the object was uploaded.
	Uploaded time.Time `json:"uploaded" format:"date-time"`
	// The version UUID of the uploaded object.
	Version string                         `json:"version"`
	JSON    bucketObjectUploadResponseJSON `json:"-"`
}

// bucketObjectUploadResponseJSON contains the JSON metadata for the struct
// [BucketObjectUploadResponse]
type bucketObjectUploadResponseJSON struct {
	Etag         apijson.Field
	Key          apijson.Field
	Size         apijson.Field
	StorageClass apijson.Field
	Uploaded     apijson.Field
	Version      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *BucketObjectUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketObjectUploadResponseJSON) RawJSON() string {
	return r.raw
}

// Storage class for newly uploaded objects, unless specified otherwise.
type BucketObjectUploadResponseStorageClass string

const (
	BucketObjectUploadResponseStorageClassStandard         BucketObjectUploadResponseStorageClass = "Standard"
	BucketObjectUploadResponseStorageClassInfrequentAccess BucketObjectUploadResponseStorageClass = "InfrequentAccess"
)

func (r BucketObjectUploadResponseStorageClass) IsKnown() bool {
	switch r {
	case BucketObjectUploadResponseStorageClassStandard, BucketObjectUploadResponseStorageClassInfrequentAccess:
		return true
	}
	return false
}

type BucketObjectListParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Pagination cursor received from a previous List Objects call. Used to retrieve
	// the next page of results.
	Cursor param.Field[string] `query:"cursor"`
	// A single character used to group keys. All keys that contain the delimiter
	// between the prefix and the first occurrence of the delimiter after the prefix
	// are grouped under a single result element.
	Delimiter param.Field[string] `query:"delimiter"`
	// Maximum number of objects to return per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Restricts results to only those objects whose keys begin with the specified
	// prefix.
	Prefix param.Field[string] `query:"prefix"`
	// Returns objects with keys that come after the specified key in lexicographic
	// order.
	StartAfter param.Field[string] `query:"start_after"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[BucketObjectListParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// URLQuery serializes [BucketObjectListParams]'s query parameters as `url.Values`.
func (r BucketObjectListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type BucketObjectListParamsCfR2Jurisdiction string

const (
	BucketObjectListParamsCfR2JurisdictionDefault BucketObjectListParamsCfR2Jurisdiction = "default"
	BucketObjectListParamsCfR2JurisdictionEu      BucketObjectListParamsCfR2Jurisdiction = "eu"
	BucketObjectListParamsCfR2JurisdictionFedramp BucketObjectListParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketObjectListParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketObjectListParamsCfR2JurisdictionDefault, BucketObjectListParamsCfR2JurisdictionEu, BucketObjectListParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketObjectDeleteParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[BucketObjectDeleteParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type BucketObjectDeleteParamsCfR2Jurisdiction string

const (
	BucketObjectDeleteParamsCfR2JurisdictionDefault BucketObjectDeleteParamsCfR2Jurisdiction = "default"
	BucketObjectDeleteParamsCfR2JurisdictionEu      BucketObjectDeleteParamsCfR2Jurisdiction = "eu"
	BucketObjectDeleteParamsCfR2JurisdictionFedramp BucketObjectDeleteParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketObjectDeleteParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketObjectDeleteParamsCfR2JurisdictionDefault, BucketObjectDeleteParamsCfR2JurisdictionEu, BucketObjectDeleteParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketObjectDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []string              `json:"messages" api:"required"`
	// Result of a successful object deletion.
	Result BucketObjectDeleteResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success BucketObjectDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    bucketObjectDeleteResponseEnvelopeJSON    `json:"-"`
}

// bucketObjectDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [BucketObjectDeleteResponseEnvelope]
type bucketObjectDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketObjectDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketObjectDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketObjectDeleteResponseEnvelopeSuccess bool

const (
	BucketObjectDeleteResponseEnvelopeSuccessTrue BucketObjectDeleteResponseEnvelopeSuccess = true
)

func (r BucketObjectDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketObjectDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BucketObjectGetParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[BucketObjectGetParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
	// Returns the object only if it has been modified since the specified time. Must
	// be formatted as an HTTP-date (RFC 7231), e.g. `Tue, 15 Jan 2024 10:30:00 GMT`.
	IfModifiedSince param.Field[string] `header:"If-Modified-Since"`
	// Returns the object only if its ETag does not match the given value.
	IfNoneMatch param.Field[string] `header:"If-None-Match"`
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type BucketObjectGetParamsCfR2Jurisdiction string

const (
	BucketObjectGetParamsCfR2JurisdictionDefault BucketObjectGetParamsCfR2Jurisdiction = "default"
	BucketObjectGetParamsCfR2JurisdictionEu      BucketObjectGetParamsCfR2Jurisdiction = "eu"
	BucketObjectGetParamsCfR2JurisdictionFedramp BucketObjectGetParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketObjectGetParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketObjectGetParamsCfR2JurisdictionDefault, BucketObjectGetParamsCfR2JurisdictionEu, BucketObjectGetParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketObjectUploadParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[BucketObjectUploadParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
	// Storage class for newly uploaded objects, unless specified otherwise.
	CfR2StorageClass param.Field[BucketObjectUploadParamsCfR2StorageClass] `header:"cf-r2-storage-class"`
}

func (r BucketObjectUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type BucketObjectUploadParamsCfR2Jurisdiction string

const (
	BucketObjectUploadParamsCfR2JurisdictionDefault BucketObjectUploadParamsCfR2Jurisdiction = "default"
	BucketObjectUploadParamsCfR2JurisdictionEu      BucketObjectUploadParamsCfR2Jurisdiction = "eu"
	BucketObjectUploadParamsCfR2JurisdictionFedramp BucketObjectUploadParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketObjectUploadParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketObjectUploadParamsCfR2JurisdictionDefault, BucketObjectUploadParamsCfR2JurisdictionEu, BucketObjectUploadParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

// Storage class for newly uploaded objects, unless specified otherwise.
type BucketObjectUploadParamsCfR2StorageClass string

const (
	BucketObjectUploadParamsCfR2StorageClassStandard         BucketObjectUploadParamsCfR2StorageClass = "Standard"
	BucketObjectUploadParamsCfR2StorageClassInfrequentAccess BucketObjectUploadParamsCfR2StorageClass = "InfrequentAccess"
)

func (r BucketObjectUploadParamsCfR2StorageClass) IsKnown() bool {
	switch r {
	case BucketObjectUploadParamsCfR2StorageClassStandard, BucketObjectUploadParamsCfR2StorageClassInfrequentAccess:
		return true
	}
	return false
}

type BucketObjectUploadResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []string              `json:"messages" api:"required"`
	// Result of a successful object upload.
	Result BucketObjectUploadResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success BucketObjectUploadResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    bucketObjectUploadResponseEnvelopeJSON    `json:"-"`
}

// bucketObjectUploadResponseEnvelopeJSON contains the JSON metadata for the struct
// [BucketObjectUploadResponseEnvelope]
type bucketObjectUploadResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketObjectUploadResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketObjectUploadResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketObjectUploadResponseEnvelopeSuccess bool

const (
	BucketObjectUploadResponseEnvelopeSuccessTrue BucketObjectUploadResponseEnvelopeSuccess = true
)

func (r BucketObjectUploadResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketObjectUploadResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
