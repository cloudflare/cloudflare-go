// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/shared"
	"github.com/tidwall/gjson"
)

// BucketJobService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBucketJobService] method instead.
type BucketJobService struct {
	Options []option.RequestOption
}

// NewBucketJobService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBucketJobService(opts ...option.RequestOption) (r *BucketJobService) {
	r = &BucketJobService{}
	r.Options = opts
	return
}

// Creates a background job for an R2 bucket. The `jobType` field selects the job:
//
//   - **`prefixDelete`**: deletes every object whose key begins with `prefix`. A
//     non-empty prefix must end in `/`. An empty prefix (`""`) empties the entire
//     bucket.
//   - **`storageClassMigration`**: migrates every object in the bucket between
//     storage classes. `sourceStorageClass` and `destinationStorageClass` must be
//     provided together and must differ; omitting both migrates from
//     `InfrequentAccess` to `Standard`.
//
// Poll the returned `id` with the Get Bucket Job endpoint. Small prefix-delete
// jobs can finish synchronously and return `COMPLETED`. Objects uploaded after a
// background job starts are not affected by that job.
//
// For prefix-delete jobs: abort active multipart uploads before submitting the
// request, since a synchronously completed job does not abort them, and avoid
// writing objects or starting multipart uploads while a bucket-emptying job is in
// progress. Each request creates a distinct job, and the number of active
// prefix-delete jobs is limited per bucket; wait for an existing job to finish
// before retrying a request rejected with HTTP 429. A bucket cannot be emptied
// while event notifications are configured (HTTP 409 / error code 10083). To
// protect a bucket with R2 Data Catalog enabled, send the
// `cf-r2-data-catalog-check` header; a conflict is returned with HTTP 409 / error
// code 10081.
//
// Prefix-delete jobs require permission to delete objects; storage-class migration
// jobs require permission to write to the bucket.
func (r *BucketJobService) New(ctx context.Context, bucketName string, params BucketJobNewParams, opts ...option.RequestOption) (res *BucketJobNewResponse, err error) {
	var env BucketJobNewResponseEnvelope
	if params.CfR2DataCatalogCheck.Present {
		opts = append(opts, option.WithHeader("cf-r2-data-catalog-check", fmt.Sprintf("%v", params.CfR2DataCatalogCheck)))
	}
	if params.CfR2Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%v", params.CfR2Jurisdiction)))
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
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/jobs", params.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists background jobs for an R2 bucket, including prefix-delete (and
// bucket-emptying) jobs and storage-class migration jobs. Jobs of every type are
// returned unless `jobType` is provided.
func (r *BucketJobService) List(ctx context.Context, bucketName string, params BucketJobListParams, opts ...option.RequestOption) (res *BucketJobListResponse, err error) {
	var env BucketJobListResponseEnvelope
	if params.CfR2Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%v", params.CfR2Jurisdiction)))
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
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/jobs", params.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Gets the current status of a background job of any type for an R2 bucket. Poll
// this endpoint with the job identifier returned by Create Bucket Job until the
// status is `COMPLETED`, `FAILED`, or `CANCELLED`.
func (r *BucketJobService) Get(ctx context.Context, bucketName string, jobID string, params BucketJobGetParams, opts ...option.RequestOption) (res *BucketJobGetResponse, err error) {
	var env BucketJobGetResponseEnvelope
	if params.CfR2Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%v", params.CfR2Jurisdiction)))
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
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/r2/buckets/%s/jobs/%s", params.AccountID, bucketName, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Descriptor of a bucket background job, discriminated by `jobType`. Returned by
// the Create Bucket Job, Get Bucket Job and List Bucket Jobs endpoints. A small
// operation can already be `COMPLETED` when it is returned.
type BucketJobNewResponse struct {
	// Unique identifier used to poll the job.
	ID string `json:"id" api:"required"`
	// The job kind.
	JobType BucketJobNewResponseJobType `json:"jobType" api:"required"`
	// When the job was created.
	StartTime time.Time `json:"startTime" api:"required" format:"date-time"`
	// Lifecycle status of a bucket background job.
	Status BucketJobNewResponseStatus `json:"status" api:"required"`
	// When the job finished. Absent while the job is still `ENQUEUED` or `RUNNING`.
	EndTime time.Time `json:"endTime" format:"date-time"`
	// This field can have the runtime type of
	// [BucketJobNewResponseR2PrefixDeleteJobPrefixDelete].
	PrefixDelete interface{} `json:"prefixDelete"`
	// This field can have the runtime type of
	// [BucketJobNewResponseR2StorageClassMigrationJobStorageClassMigration].
	StorageClassMigration interface{}              `json:"storageClassMigration"`
	JSON                  bucketJobNewResponseJSON `json:"-"`
	union                 BucketJobNewResponseUnion
}

// bucketJobNewResponseJSON contains the JSON metadata for the struct
// [BucketJobNewResponse]
type bucketJobNewResponseJSON struct {
	ID                    apijson.Field
	JobType               apijson.Field
	StartTime             apijson.Field
	Status                apijson.Field
	EndTime               apijson.Field
	PrefixDelete          apijson.Field
	StorageClassMigration apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r bucketJobNewResponseJSON) RawJSON() string {
	return r.raw
}

func (r *BucketJobNewResponse) UnmarshalJSON(data []byte) (err error) {
	*r = BucketJobNewResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [BucketJobNewResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [BucketJobNewResponseR2PrefixDeleteJob],
// [BucketJobNewResponseR2StorageClassMigrationJob].
func (r BucketJobNewResponse) AsUnion() BucketJobNewResponseUnion {
	return r.union
}

// Descriptor of a bucket background job, discriminated by `jobType`. Returned by
// the Create Bucket Job, Get Bucket Job and List Bucket Jobs endpoints. A small
// operation can already be `COMPLETED` when it is returned.
//
// Union satisfied by [BucketJobNewResponseR2PrefixDeleteJob] or
// [BucketJobNewResponseR2StorageClassMigrationJob].
type BucketJobNewResponseUnion interface {
	implementsBucketJobNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BucketJobNewResponseUnion)(nil)).Elem(),
		"jobType",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BucketJobNewResponseR2PrefixDeleteJob{}),
			DiscriminatorValue: "prefixDelete",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BucketJobNewResponseR2StorageClassMigrationJob{}),
			DiscriminatorValue: "storageClassMigration",
		},
	)
}

// Descriptor of a prefix-delete job. An empty prefix indicates a job that empties
// the whole bucket.
type BucketJobNewResponseR2PrefixDeleteJob struct {
	// Unique identifier used to poll the job.
	ID string `json:"id" api:"required"`
	// The job kind.
	JobType BucketJobNewResponseR2PrefixDeleteJobJobType `json:"jobType" api:"required"`
	// Details specific to the prefix-delete job.
	PrefixDelete BucketJobNewResponseR2PrefixDeleteJobPrefixDelete `json:"prefixDelete" api:"required"`
	// When the job was created.
	StartTime time.Time `json:"startTime" api:"required" format:"date-time"`
	// Lifecycle status of a bucket background job.
	Status BucketJobNewResponseR2PrefixDeleteJobStatus `json:"status" api:"required"`
	// When the job finished. Absent while the job is still `ENQUEUED` or `RUNNING`.
	EndTime time.Time                                 `json:"endTime" format:"date-time"`
	JSON    bucketJobNewResponseR2PrefixDeleteJobJSON `json:"-"`
}

// bucketJobNewResponseR2PrefixDeleteJobJSON contains the JSON metadata for the
// struct [BucketJobNewResponseR2PrefixDeleteJob]
type bucketJobNewResponseR2PrefixDeleteJobJSON struct {
	ID           apijson.Field
	JobType      apijson.Field
	PrefixDelete apijson.Field
	StartTime    apijson.Field
	Status       apijson.Field
	EndTime      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *BucketJobNewResponseR2PrefixDeleteJob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketJobNewResponseR2PrefixDeleteJobJSON) RawJSON() string {
	return r.raw
}

func (r BucketJobNewResponseR2PrefixDeleteJob) implementsBucketJobNewResponse() {}

// The job kind.
type BucketJobNewResponseR2PrefixDeleteJobJobType string

const (
	BucketJobNewResponseR2PrefixDeleteJobJobTypePrefixDelete BucketJobNewResponseR2PrefixDeleteJobJobType = "prefixDelete"
)

func (r BucketJobNewResponseR2PrefixDeleteJobJobType) IsKnown() bool {
	switch r {
	case BucketJobNewResponseR2PrefixDeleteJobJobTypePrefixDelete:
		return true
	}
	return false
}

// Details specific to the prefix-delete job.
type BucketJobNewResponseR2PrefixDeleteJobPrefixDelete struct {
	// Number of objects deleted by the job so far.
	DeletedObjects int64 `json:"deletedObjects" api:"required"`
	// The prefix matched against object keys. An empty string means the entire bucket.
	Prefix string `json:"prefix" api:"required"`
	// True when the job was created to clear the entire bucket. This field can be
	// absent on jobs created before Empty Bucket job tracking was introduced.
	IsBucketClear bool                                                  `json:"isBucketClear"`
	JSON          bucketJobNewResponseR2PrefixDeleteJobPrefixDeleteJSON `json:"-"`
}

// bucketJobNewResponseR2PrefixDeleteJobPrefixDeleteJSON contains the JSON metadata
// for the struct [BucketJobNewResponseR2PrefixDeleteJobPrefixDelete]
type bucketJobNewResponseR2PrefixDeleteJobPrefixDeleteJSON struct {
	DeletedObjects apijson.Field
	Prefix         apijson.Field
	IsBucketClear  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BucketJobNewResponseR2PrefixDeleteJobPrefixDelete) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketJobNewResponseR2PrefixDeleteJobPrefixDeleteJSON) RawJSON() string {
	return r.raw
}

// Lifecycle status of a bucket background job.
type BucketJobNewResponseR2PrefixDeleteJobStatus string

const (
	BucketJobNewResponseR2PrefixDeleteJobStatusEnqueued  BucketJobNewResponseR2PrefixDeleteJobStatus = "ENQUEUED"
	BucketJobNewResponseR2PrefixDeleteJobStatusRunning   BucketJobNewResponseR2PrefixDeleteJobStatus = "RUNNING"
	BucketJobNewResponseR2PrefixDeleteJobStatusCompleted BucketJobNewResponseR2PrefixDeleteJobStatus = "COMPLETED"
	BucketJobNewResponseR2PrefixDeleteJobStatusFailed    BucketJobNewResponseR2PrefixDeleteJobStatus = "FAILED"
	BucketJobNewResponseR2PrefixDeleteJobStatusCancelled BucketJobNewResponseR2PrefixDeleteJobStatus = "CANCELLED"
)

func (r BucketJobNewResponseR2PrefixDeleteJobStatus) IsKnown() bool {
	switch r {
	case BucketJobNewResponseR2PrefixDeleteJobStatusEnqueued, BucketJobNewResponseR2PrefixDeleteJobStatusRunning, BucketJobNewResponseR2PrefixDeleteJobStatusCompleted, BucketJobNewResponseR2PrefixDeleteJobStatusFailed, BucketJobNewResponseR2PrefixDeleteJobStatusCancelled:
		return true
	}
	return false
}

// Descriptor of a bucket-wide storage-class migration job.
type BucketJobNewResponseR2StorageClassMigrationJob struct {
	// Unique identifier used to poll the job.
	ID string `json:"id" api:"required"`
	// The job kind.
	JobType BucketJobNewResponseR2StorageClassMigrationJobJobType `json:"jobType" api:"required"`
	// When the job was created.
	StartTime time.Time `json:"startTime" api:"required" format:"date-time"`
	// Lifecycle status of a bucket background job.
	Status                BucketJobNewResponseR2StorageClassMigrationJobStatus                `json:"status" api:"required"`
	StorageClassMigration BucketJobNewResponseR2StorageClassMigrationJobStorageClassMigration `json:"storageClassMigration" api:"required"`
	// When the job finished. Absent while the job is still `ENQUEUED` or `RUNNING`.
	EndTime time.Time                                          `json:"endTime" format:"date-time"`
	JSON    bucketJobNewResponseR2StorageClassMigrationJobJSON `json:"-"`
}

// bucketJobNewResponseR2StorageClassMigrationJobJSON contains the JSON metadata
// for the struct [BucketJobNewResponseR2StorageClassMigrationJob]
type bucketJobNewResponseR2StorageClassMigrationJobJSON struct {
	ID                    apijson.Field
	JobType               apijson.Field
	StartTime             apijson.Field
	Status                apijson.Field
	StorageClassMigration apijson.Field
	EndTime               apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *BucketJobNewResponseR2StorageClassMigrationJob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketJobNewResponseR2StorageClassMigrationJobJSON) RawJSON() string {
	return r.raw
}

func (r BucketJobNewResponseR2StorageClassMigrationJob) implementsBucketJobNewResponse() {}

// The job kind.
type BucketJobNewResponseR2StorageClassMigrationJobJobType string

const (
	BucketJobNewResponseR2StorageClassMigrationJobJobTypeStorageClassMigration BucketJobNewResponseR2StorageClassMigrationJobJobType = "storageClassMigration"
)

func (r BucketJobNewResponseR2StorageClassMigrationJobJobType) IsKnown() bool {
	switch r {
	case BucketJobNewResponseR2StorageClassMigrationJobJobTypeStorageClassMigration:
		return true
	}
	return false
}

// Lifecycle status of a bucket background job.
type BucketJobNewResponseR2StorageClassMigrationJobStatus string

const (
	BucketJobNewResponseR2StorageClassMigrationJobStatusEnqueued  BucketJobNewResponseR2StorageClassMigrationJobStatus = "ENQUEUED"
	BucketJobNewResponseR2StorageClassMigrationJobStatusRunning   BucketJobNewResponseR2StorageClassMigrationJobStatus = "RUNNING"
	BucketJobNewResponseR2StorageClassMigrationJobStatusCompleted BucketJobNewResponseR2StorageClassMigrationJobStatus = "COMPLETED"
	BucketJobNewResponseR2StorageClassMigrationJobStatusFailed    BucketJobNewResponseR2StorageClassMigrationJobStatus = "FAILED"
	BucketJobNewResponseR2StorageClassMigrationJobStatusCancelled BucketJobNewResponseR2StorageClassMigrationJobStatus = "CANCELLED"
)

func (r BucketJobNewResponseR2StorageClassMigrationJobStatus) IsKnown() bool {
	switch r {
	case BucketJobNewResponseR2StorageClassMigrationJobStatusEnqueued, BucketJobNewResponseR2StorageClassMigrationJobStatusRunning, BucketJobNewResponseR2StorageClassMigrationJobStatusCompleted, BucketJobNewResponseR2StorageClassMigrationJobStatusFailed, BucketJobNewResponseR2StorageClassMigrationJobStatusCancelled:
		return true
	}
	return false
}

type BucketJobNewResponseR2StorageClassMigrationJobStorageClassMigration struct {
	// The storage class objects are migrated to. The source and destination storage
	// classes always differ.
	DestinationStorageClass BucketJobNewResponseR2StorageClassMigrationJobStorageClassMigrationDestinationStorageClass `json:"destinationStorageClass" api:"required"`
	// Number of object versions migrated successfully.
	MigratedObjects int64 `json:"migratedObjects" api:"required"`
	// The storage class objects are migrated from.
	SourceStorageClass BucketJobNewResponseR2StorageClassMigrationJobStorageClassMigrationSourceStorageClass `json:"sourceStorageClass" api:"required"`
	// Payload bytes belonging to failed object versions.
	FailedBytes int64 `json:"failedBytes"`
	// Number of object versions that failed migration.
	FailedObjects int64 `json:"failedObjects"`
	// Object payload bytes migrated successfully. Present for Metadata Worker jobs.
	MigratedBytes int64 `json:"migratedBytes"`
	// Object metadata bytes migrated successfully. Present for Metadata Worker jobs.
	MigratedMetadataBytes int64 `json:"migratedMetadataBytes"`
	// Payload bytes belonging to skipped object versions.
	SkippedBytes int64 `json:"skippedBytes"`
	// Number of object versions skipped.
	SkippedObjects int64                                                                   `json:"skippedObjects"`
	JSON           bucketJobNewResponseR2StorageClassMigrationJobStorageClassMigrationJSON `json:"-"`
}

// bucketJobNewResponseR2StorageClassMigrationJobStorageClassMigrationJSON contains
// the JSON metadata for the struct
// [BucketJobNewResponseR2StorageClassMigrationJobStorageClassMigration]
type bucketJobNewResponseR2StorageClassMigrationJobStorageClassMigrationJSON struct {
	DestinationStorageClass apijson.Field
	MigratedObjects         apijson.Field
	SourceStorageClass      apijson.Field
	FailedBytes             apijson.Field
	FailedObjects           apijson.Field
	MigratedBytes           apijson.Field
	MigratedMetadataBytes   apijson.Field
	SkippedBytes            apijson.Field
	SkippedObjects          apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *BucketJobNewResponseR2StorageClassMigrationJobStorageClassMigration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketJobNewResponseR2StorageClassMigrationJobStorageClassMigrationJSON) RawJSON() string {
	return r.raw
}

// The storage class objects are migrated to. The source and destination storage
// classes always differ.
type BucketJobNewResponseR2StorageClassMigrationJobStorageClassMigrationDestinationStorageClass string

const (
	BucketJobNewResponseR2StorageClassMigrationJobStorageClassMigrationDestinationStorageClassStandard         BucketJobNewResponseR2StorageClassMigrationJobStorageClassMigrationDestinationStorageClass = "Standard"
	BucketJobNewResponseR2StorageClassMigrationJobStorageClassMigrationDestinationStorageClassInfrequentAccess BucketJobNewResponseR2StorageClassMigrationJobStorageClassMigrationDestinationStorageClass = "InfrequentAccess"
)

func (r BucketJobNewResponseR2StorageClassMigrationJobStorageClassMigrationDestinationStorageClass) IsKnown() bool {
	switch r {
	case BucketJobNewResponseR2StorageClassMigrationJobStorageClassMigrationDestinationStorageClassStandard, BucketJobNewResponseR2StorageClassMigrationJobStorageClassMigrationDestinationStorageClassInfrequentAccess:
		return true
	}
	return false
}

// The storage class objects are migrated from.
type BucketJobNewResponseR2StorageClassMigrationJobStorageClassMigrationSourceStorageClass string

const (
	BucketJobNewResponseR2StorageClassMigrationJobStorageClassMigrationSourceStorageClassStandard         BucketJobNewResponseR2StorageClassMigrationJobStorageClassMigrationSourceStorageClass = "Standard"
	BucketJobNewResponseR2StorageClassMigrationJobStorageClassMigrationSourceStorageClassInfrequentAccess BucketJobNewResponseR2StorageClassMigrationJobStorageClassMigrationSourceStorageClass = "InfrequentAccess"
)

func (r BucketJobNewResponseR2StorageClassMigrationJobStorageClassMigrationSourceStorageClass) IsKnown() bool {
	switch r {
	case BucketJobNewResponseR2StorageClassMigrationJobStorageClassMigrationSourceStorageClassStandard, BucketJobNewResponseR2StorageClassMigrationJobStorageClassMigrationSourceStorageClassInfrequentAccess:
		return true
	}
	return false
}

// The job kind.
type BucketJobNewResponseJobType string

const (
	BucketJobNewResponseJobTypePrefixDelete          BucketJobNewResponseJobType = "prefixDelete"
	BucketJobNewResponseJobTypeStorageClassMigration BucketJobNewResponseJobType = "storageClassMigration"
)

func (r BucketJobNewResponseJobType) IsKnown() bool {
	switch r {
	case BucketJobNewResponseJobTypePrefixDelete, BucketJobNewResponseJobTypeStorageClassMigration:
		return true
	}
	return false
}

// Lifecycle status of a bucket background job.
type BucketJobNewResponseStatus string

const (
	BucketJobNewResponseStatusEnqueued  BucketJobNewResponseStatus = "ENQUEUED"
	BucketJobNewResponseStatusRunning   BucketJobNewResponseStatus = "RUNNING"
	BucketJobNewResponseStatusCompleted BucketJobNewResponseStatus = "COMPLETED"
	BucketJobNewResponseStatusFailed    BucketJobNewResponseStatus = "FAILED"
	BucketJobNewResponseStatusCancelled BucketJobNewResponseStatus = "CANCELLED"
)

func (r BucketJobNewResponseStatus) IsKnown() bool {
	switch r {
	case BucketJobNewResponseStatusEnqueued, BucketJobNewResponseStatusRunning, BucketJobNewResponseStatusCompleted, BucketJobNewResponseStatusFailed, BucketJobNewResponseStatusCancelled:
		return true
	}
	return false
}

type BucketJobListResponse struct {
	// Bucket jobs in the current page.
	Jobs []BucketJobListResponseJob `json:"jobs" api:"required"`
	// Token to pass as `continuationToken` to retrieve the next page.
	NextContinuationToken string                    `json:"nextContinuationToken"`
	JSON                  bucketJobListResponseJSON `json:"-"`
}

// bucketJobListResponseJSON contains the JSON metadata for the struct
// [BucketJobListResponse]
type bucketJobListResponseJSON struct {
	Jobs                  apijson.Field
	NextContinuationToken apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *BucketJobListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketJobListResponseJSON) RawJSON() string {
	return r.raw
}

// Descriptor of a bucket background job, discriminated by `jobType`. Returned by
// the Create Bucket Job, Get Bucket Job and List Bucket Jobs endpoints. A small
// operation can already be `COMPLETED` when it is returned.
type BucketJobListResponseJob struct {
	// Unique identifier used to poll the job.
	ID string `json:"id" api:"required"`
	// The job kind.
	JobType BucketJobListResponseJobsJobType `json:"jobType" api:"required"`
	// When the job was created.
	StartTime time.Time `json:"startTime" api:"required" format:"date-time"`
	// Lifecycle status of a bucket background job.
	Status BucketJobListResponseJobsStatus `json:"status" api:"required"`
	// When the job finished. Absent while the job is still `ENQUEUED` or `RUNNING`.
	EndTime time.Time `json:"endTime" format:"date-time"`
	// This field can have the runtime type of
	// [BucketJobListResponseJobsR2PrefixDeleteJobPrefixDelete].
	PrefixDelete interface{} `json:"prefixDelete"`
	// This field can have the runtime type of
	// [BucketJobListResponseJobsR2StorageClassMigrationJobStorageClassMigration].
	StorageClassMigration interface{}                  `json:"storageClassMigration"`
	JSON                  bucketJobListResponseJobJSON `json:"-"`
	union                 BucketJobListResponseJobsUnion
}

// bucketJobListResponseJobJSON contains the JSON metadata for the struct
// [BucketJobListResponseJob]
type bucketJobListResponseJobJSON struct {
	ID                    apijson.Field
	JobType               apijson.Field
	StartTime             apijson.Field
	Status                apijson.Field
	EndTime               apijson.Field
	PrefixDelete          apijson.Field
	StorageClassMigration apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r bucketJobListResponseJobJSON) RawJSON() string {
	return r.raw
}

func (r *BucketJobListResponseJob) UnmarshalJSON(data []byte) (err error) {
	*r = BucketJobListResponseJob{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [BucketJobListResponseJobsUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [BucketJobListResponseJobsR2PrefixDeleteJob],
// [BucketJobListResponseJobsR2StorageClassMigrationJob].
func (r BucketJobListResponseJob) AsUnion() BucketJobListResponseJobsUnion {
	return r.union
}

// Descriptor of a bucket background job, discriminated by `jobType`. Returned by
// the Create Bucket Job, Get Bucket Job and List Bucket Jobs endpoints. A small
// operation can already be `COMPLETED` when it is returned.
//
// Union satisfied by [BucketJobListResponseJobsR2PrefixDeleteJob] or
// [BucketJobListResponseJobsR2StorageClassMigrationJob].
type BucketJobListResponseJobsUnion interface {
	implementsBucketJobListResponseJob()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BucketJobListResponseJobsUnion)(nil)).Elem(),
		"jobType",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BucketJobListResponseJobsR2PrefixDeleteJob{}),
			DiscriminatorValue: "prefixDelete",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BucketJobListResponseJobsR2StorageClassMigrationJob{}),
			DiscriminatorValue: "storageClassMigration",
		},
	)
}

// Descriptor of a prefix-delete job. An empty prefix indicates a job that empties
// the whole bucket.
type BucketJobListResponseJobsR2PrefixDeleteJob struct {
	// Unique identifier used to poll the job.
	ID string `json:"id" api:"required"`
	// The job kind.
	JobType BucketJobListResponseJobsR2PrefixDeleteJobJobType `json:"jobType" api:"required"`
	// Details specific to the prefix-delete job.
	PrefixDelete BucketJobListResponseJobsR2PrefixDeleteJobPrefixDelete `json:"prefixDelete" api:"required"`
	// When the job was created.
	StartTime time.Time `json:"startTime" api:"required" format:"date-time"`
	// Lifecycle status of a bucket background job.
	Status BucketJobListResponseJobsR2PrefixDeleteJobStatus `json:"status" api:"required"`
	// When the job finished. Absent while the job is still `ENQUEUED` or `RUNNING`.
	EndTime time.Time                                      `json:"endTime" format:"date-time"`
	JSON    bucketJobListResponseJobsR2PrefixDeleteJobJSON `json:"-"`
}

// bucketJobListResponseJobsR2PrefixDeleteJobJSON contains the JSON metadata for
// the struct [BucketJobListResponseJobsR2PrefixDeleteJob]
type bucketJobListResponseJobsR2PrefixDeleteJobJSON struct {
	ID           apijson.Field
	JobType      apijson.Field
	PrefixDelete apijson.Field
	StartTime    apijson.Field
	Status       apijson.Field
	EndTime      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *BucketJobListResponseJobsR2PrefixDeleteJob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketJobListResponseJobsR2PrefixDeleteJobJSON) RawJSON() string {
	return r.raw
}

func (r BucketJobListResponseJobsR2PrefixDeleteJob) implementsBucketJobListResponseJob() {}

// The job kind.
type BucketJobListResponseJobsR2PrefixDeleteJobJobType string

const (
	BucketJobListResponseJobsR2PrefixDeleteJobJobTypePrefixDelete BucketJobListResponseJobsR2PrefixDeleteJobJobType = "prefixDelete"
)

func (r BucketJobListResponseJobsR2PrefixDeleteJobJobType) IsKnown() bool {
	switch r {
	case BucketJobListResponseJobsR2PrefixDeleteJobJobTypePrefixDelete:
		return true
	}
	return false
}

// Details specific to the prefix-delete job.
type BucketJobListResponseJobsR2PrefixDeleteJobPrefixDelete struct {
	// Number of objects deleted by the job so far.
	DeletedObjects int64 `json:"deletedObjects" api:"required"`
	// The prefix matched against object keys. An empty string means the entire bucket.
	Prefix string `json:"prefix" api:"required"`
	// True when the job was created to clear the entire bucket. This field can be
	// absent on jobs created before Empty Bucket job tracking was introduced.
	IsBucketClear bool                                                       `json:"isBucketClear"`
	JSON          bucketJobListResponseJobsR2PrefixDeleteJobPrefixDeleteJSON `json:"-"`
}

// bucketJobListResponseJobsR2PrefixDeleteJobPrefixDeleteJSON contains the JSON
// metadata for the struct [BucketJobListResponseJobsR2PrefixDeleteJobPrefixDelete]
type bucketJobListResponseJobsR2PrefixDeleteJobPrefixDeleteJSON struct {
	DeletedObjects apijson.Field
	Prefix         apijson.Field
	IsBucketClear  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BucketJobListResponseJobsR2PrefixDeleteJobPrefixDelete) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketJobListResponseJobsR2PrefixDeleteJobPrefixDeleteJSON) RawJSON() string {
	return r.raw
}

// Lifecycle status of a bucket background job.
type BucketJobListResponseJobsR2PrefixDeleteJobStatus string

const (
	BucketJobListResponseJobsR2PrefixDeleteJobStatusEnqueued  BucketJobListResponseJobsR2PrefixDeleteJobStatus = "ENQUEUED"
	BucketJobListResponseJobsR2PrefixDeleteJobStatusRunning   BucketJobListResponseJobsR2PrefixDeleteJobStatus = "RUNNING"
	BucketJobListResponseJobsR2PrefixDeleteJobStatusCompleted BucketJobListResponseJobsR2PrefixDeleteJobStatus = "COMPLETED"
	BucketJobListResponseJobsR2PrefixDeleteJobStatusFailed    BucketJobListResponseJobsR2PrefixDeleteJobStatus = "FAILED"
	BucketJobListResponseJobsR2PrefixDeleteJobStatusCancelled BucketJobListResponseJobsR2PrefixDeleteJobStatus = "CANCELLED"
)

func (r BucketJobListResponseJobsR2PrefixDeleteJobStatus) IsKnown() bool {
	switch r {
	case BucketJobListResponseJobsR2PrefixDeleteJobStatusEnqueued, BucketJobListResponseJobsR2PrefixDeleteJobStatusRunning, BucketJobListResponseJobsR2PrefixDeleteJobStatusCompleted, BucketJobListResponseJobsR2PrefixDeleteJobStatusFailed, BucketJobListResponseJobsR2PrefixDeleteJobStatusCancelled:
		return true
	}
	return false
}

// Descriptor of a bucket-wide storage-class migration job.
type BucketJobListResponseJobsR2StorageClassMigrationJob struct {
	// Unique identifier used to poll the job.
	ID string `json:"id" api:"required"`
	// The job kind.
	JobType BucketJobListResponseJobsR2StorageClassMigrationJobJobType `json:"jobType" api:"required"`
	// When the job was created.
	StartTime time.Time `json:"startTime" api:"required" format:"date-time"`
	// Lifecycle status of a bucket background job.
	Status                BucketJobListResponseJobsR2StorageClassMigrationJobStatus                `json:"status" api:"required"`
	StorageClassMigration BucketJobListResponseJobsR2StorageClassMigrationJobStorageClassMigration `json:"storageClassMigration" api:"required"`
	// When the job finished. Absent while the job is still `ENQUEUED` or `RUNNING`.
	EndTime time.Time                                               `json:"endTime" format:"date-time"`
	JSON    bucketJobListResponseJobsR2StorageClassMigrationJobJSON `json:"-"`
}

// bucketJobListResponseJobsR2StorageClassMigrationJobJSON contains the JSON
// metadata for the struct [BucketJobListResponseJobsR2StorageClassMigrationJob]
type bucketJobListResponseJobsR2StorageClassMigrationJobJSON struct {
	ID                    apijson.Field
	JobType               apijson.Field
	StartTime             apijson.Field
	Status                apijson.Field
	StorageClassMigration apijson.Field
	EndTime               apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *BucketJobListResponseJobsR2StorageClassMigrationJob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketJobListResponseJobsR2StorageClassMigrationJobJSON) RawJSON() string {
	return r.raw
}

func (r BucketJobListResponseJobsR2StorageClassMigrationJob) implementsBucketJobListResponseJob() {}

// The job kind.
type BucketJobListResponseJobsR2StorageClassMigrationJobJobType string

const (
	BucketJobListResponseJobsR2StorageClassMigrationJobJobTypeStorageClassMigration BucketJobListResponseJobsR2StorageClassMigrationJobJobType = "storageClassMigration"
)

func (r BucketJobListResponseJobsR2StorageClassMigrationJobJobType) IsKnown() bool {
	switch r {
	case BucketJobListResponseJobsR2StorageClassMigrationJobJobTypeStorageClassMigration:
		return true
	}
	return false
}

// Lifecycle status of a bucket background job.
type BucketJobListResponseJobsR2StorageClassMigrationJobStatus string

const (
	BucketJobListResponseJobsR2StorageClassMigrationJobStatusEnqueued  BucketJobListResponseJobsR2StorageClassMigrationJobStatus = "ENQUEUED"
	BucketJobListResponseJobsR2StorageClassMigrationJobStatusRunning   BucketJobListResponseJobsR2StorageClassMigrationJobStatus = "RUNNING"
	BucketJobListResponseJobsR2StorageClassMigrationJobStatusCompleted BucketJobListResponseJobsR2StorageClassMigrationJobStatus = "COMPLETED"
	BucketJobListResponseJobsR2StorageClassMigrationJobStatusFailed    BucketJobListResponseJobsR2StorageClassMigrationJobStatus = "FAILED"
	BucketJobListResponseJobsR2StorageClassMigrationJobStatusCancelled BucketJobListResponseJobsR2StorageClassMigrationJobStatus = "CANCELLED"
)

func (r BucketJobListResponseJobsR2StorageClassMigrationJobStatus) IsKnown() bool {
	switch r {
	case BucketJobListResponseJobsR2StorageClassMigrationJobStatusEnqueued, BucketJobListResponseJobsR2StorageClassMigrationJobStatusRunning, BucketJobListResponseJobsR2StorageClassMigrationJobStatusCompleted, BucketJobListResponseJobsR2StorageClassMigrationJobStatusFailed, BucketJobListResponseJobsR2StorageClassMigrationJobStatusCancelled:
		return true
	}
	return false
}

type BucketJobListResponseJobsR2StorageClassMigrationJobStorageClassMigration struct {
	// The storage class objects are migrated to. The source and destination storage
	// classes always differ.
	DestinationStorageClass BucketJobListResponseJobsR2StorageClassMigrationJobStorageClassMigrationDestinationStorageClass `json:"destinationStorageClass" api:"required"`
	// Number of object versions migrated successfully.
	MigratedObjects int64 `json:"migratedObjects" api:"required"`
	// The storage class objects are migrated from.
	SourceStorageClass BucketJobListResponseJobsR2StorageClassMigrationJobStorageClassMigrationSourceStorageClass `json:"sourceStorageClass" api:"required"`
	// Payload bytes belonging to failed object versions.
	FailedBytes int64 `json:"failedBytes"`
	// Number of object versions that failed migration.
	FailedObjects int64 `json:"failedObjects"`
	// Object payload bytes migrated successfully. Present for Metadata Worker jobs.
	MigratedBytes int64 `json:"migratedBytes"`
	// Object metadata bytes migrated successfully. Present for Metadata Worker jobs.
	MigratedMetadataBytes int64 `json:"migratedMetadataBytes"`
	// Payload bytes belonging to skipped object versions.
	SkippedBytes int64 `json:"skippedBytes"`
	// Number of object versions skipped.
	SkippedObjects int64                                                                        `json:"skippedObjects"`
	JSON           bucketJobListResponseJobsR2StorageClassMigrationJobStorageClassMigrationJSON `json:"-"`
}

// bucketJobListResponseJobsR2StorageClassMigrationJobStorageClassMigrationJSON
// contains the JSON metadata for the struct
// [BucketJobListResponseJobsR2StorageClassMigrationJobStorageClassMigration]
type bucketJobListResponseJobsR2StorageClassMigrationJobStorageClassMigrationJSON struct {
	DestinationStorageClass apijson.Field
	MigratedObjects         apijson.Field
	SourceStorageClass      apijson.Field
	FailedBytes             apijson.Field
	FailedObjects           apijson.Field
	MigratedBytes           apijson.Field
	MigratedMetadataBytes   apijson.Field
	SkippedBytes            apijson.Field
	SkippedObjects          apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *BucketJobListResponseJobsR2StorageClassMigrationJobStorageClassMigration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketJobListResponseJobsR2StorageClassMigrationJobStorageClassMigrationJSON) RawJSON() string {
	return r.raw
}

// The storage class objects are migrated to. The source and destination storage
// classes always differ.
type BucketJobListResponseJobsR2StorageClassMigrationJobStorageClassMigrationDestinationStorageClass string

const (
	BucketJobListResponseJobsR2StorageClassMigrationJobStorageClassMigrationDestinationStorageClassStandard         BucketJobListResponseJobsR2StorageClassMigrationJobStorageClassMigrationDestinationStorageClass = "Standard"
	BucketJobListResponseJobsR2StorageClassMigrationJobStorageClassMigrationDestinationStorageClassInfrequentAccess BucketJobListResponseJobsR2StorageClassMigrationJobStorageClassMigrationDestinationStorageClass = "InfrequentAccess"
)

func (r BucketJobListResponseJobsR2StorageClassMigrationJobStorageClassMigrationDestinationStorageClass) IsKnown() bool {
	switch r {
	case BucketJobListResponseJobsR2StorageClassMigrationJobStorageClassMigrationDestinationStorageClassStandard, BucketJobListResponseJobsR2StorageClassMigrationJobStorageClassMigrationDestinationStorageClassInfrequentAccess:
		return true
	}
	return false
}

// The storage class objects are migrated from.
type BucketJobListResponseJobsR2StorageClassMigrationJobStorageClassMigrationSourceStorageClass string

const (
	BucketJobListResponseJobsR2StorageClassMigrationJobStorageClassMigrationSourceStorageClassStandard         BucketJobListResponseJobsR2StorageClassMigrationJobStorageClassMigrationSourceStorageClass = "Standard"
	BucketJobListResponseJobsR2StorageClassMigrationJobStorageClassMigrationSourceStorageClassInfrequentAccess BucketJobListResponseJobsR2StorageClassMigrationJobStorageClassMigrationSourceStorageClass = "InfrequentAccess"
)

func (r BucketJobListResponseJobsR2StorageClassMigrationJobStorageClassMigrationSourceStorageClass) IsKnown() bool {
	switch r {
	case BucketJobListResponseJobsR2StorageClassMigrationJobStorageClassMigrationSourceStorageClassStandard, BucketJobListResponseJobsR2StorageClassMigrationJobStorageClassMigrationSourceStorageClassInfrequentAccess:
		return true
	}
	return false
}

// The job kind.
type BucketJobListResponseJobsJobType string

const (
	BucketJobListResponseJobsJobTypePrefixDelete          BucketJobListResponseJobsJobType = "prefixDelete"
	BucketJobListResponseJobsJobTypeStorageClassMigration BucketJobListResponseJobsJobType = "storageClassMigration"
)

func (r BucketJobListResponseJobsJobType) IsKnown() bool {
	switch r {
	case BucketJobListResponseJobsJobTypePrefixDelete, BucketJobListResponseJobsJobTypeStorageClassMigration:
		return true
	}
	return false
}

// Lifecycle status of a bucket background job.
type BucketJobListResponseJobsStatus string

const (
	BucketJobListResponseJobsStatusEnqueued  BucketJobListResponseJobsStatus = "ENQUEUED"
	BucketJobListResponseJobsStatusRunning   BucketJobListResponseJobsStatus = "RUNNING"
	BucketJobListResponseJobsStatusCompleted BucketJobListResponseJobsStatus = "COMPLETED"
	BucketJobListResponseJobsStatusFailed    BucketJobListResponseJobsStatus = "FAILED"
	BucketJobListResponseJobsStatusCancelled BucketJobListResponseJobsStatus = "CANCELLED"
)

func (r BucketJobListResponseJobsStatus) IsKnown() bool {
	switch r {
	case BucketJobListResponseJobsStatusEnqueued, BucketJobListResponseJobsStatusRunning, BucketJobListResponseJobsStatusCompleted, BucketJobListResponseJobsStatusFailed, BucketJobListResponseJobsStatusCancelled:
		return true
	}
	return false
}

// Descriptor of a bucket background job, discriminated by `jobType`. Returned by
// the Create Bucket Job, Get Bucket Job and List Bucket Jobs endpoints. A small
// operation can already be `COMPLETED` when it is returned.
type BucketJobGetResponse struct {
	// Unique identifier used to poll the job.
	ID string `json:"id" api:"required"`
	// The job kind.
	JobType BucketJobGetResponseJobType `json:"jobType" api:"required"`
	// When the job was created.
	StartTime time.Time `json:"startTime" api:"required" format:"date-time"`
	// Lifecycle status of a bucket background job.
	Status BucketJobGetResponseStatus `json:"status" api:"required"`
	// When the job finished. Absent while the job is still `ENQUEUED` or `RUNNING`.
	EndTime time.Time `json:"endTime" format:"date-time"`
	// This field can have the runtime type of
	// [BucketJobGetResponseR2PrefixDeleteJobPrefixDelete].
	PrefixDelete interface{} `json:"prefixDelete"`
	// This field can have the runtime type of
	// [BucketJobGetResponseR2StorageClassMigrationJobStorageClassMigration].
	StorageClassMigration interface{}              `json:"storageClassMigration"`
	JSON                  bucketJobGetResponseJSON `json:"-"`
	union                 BucketJobGetResponseUnion
}

// bucketJobGetResponseJSON contains the JSON metadata for the struct
// [BucketJobGetResponse]
type bucketJobGetResponseJSON struct {
	ID                    apijson.Field
	JobType               apijson.Field
	StartTime             apijson.Field
	Status                apijson.Field
	EndTime               apijson.Field
	PrefixDelete          apijson.Field
	StorageClassMigration apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r bucketJobGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *BucketJobGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = BucketJobGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [BucketJobGetResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [BucketJobGetResponseR2PrefixDeleteJob],
// [BucketJobGetResponseR2StorageClassMigrationJob].
func (r BucketJobGetResponse) AsUnion() BucketJobGetResponseUnion {
	return r.union
}

// Descriptor of a bucket background job, discriminated by `jobType`. Returned by
// the Create Bucket Job, Get Bucket Job and List Bucket Jobs endpoints. A small
// operation can already be `COMPLETED` when it is returned.
//
// Union satisfied by [BucketJobGetResponseR2PrefixDeleteJob] or
// [BucketJobGetResponseR2StorageClassMigrationJob].
type BucketJobGetResponseUnion interface {
	implementsBucketJobGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BucketJobGetResponseUnion)(nil)).Elem(),
		"jobType",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BucketJobGetResponseR2PrefixDeleteJob{}),
			DiscriminatorValue: "prefixDelete",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BucketJobGetResponseR2StorageClassMigrationJob{}),
			DiscriminatorValue: "storageClassMigration",
		},
	)
}

// Descriptor of a prefix-delete job. An empty prefix indicates a job that empties
// the whole bucket.
type BucketJobGetResponseR2PrefixDeleteJob struct {
	// Unique identifier used to poll the job.
	ID string `json:"id" api:"required"`
	// The job kind.
	JobType BucketJobGetResponseR2PrefixDeleteJobJobType `json:"jobType" api:"required"`
	// Details specific to the prefix-delete job.
	PrefixDelete BucketJobGetResponseR2PrefixDeleteJobPrefixDelete `json:"prefixDelete" api:"required"`
	// When the job was created.
	StartTime time.Time `json:"startTime" api:"required" format:"date-time"`
	// Lifecycle status of a bucket background job.
	Status BucketJobGetResponseR2PrefixDeleteJobStatus `json:"status" api:"required"`
	// When the job finished. Absent while the job is still `ENQUEUED` or `RUNNING`.
	EndTime time.Time                                 `json:"endTime" format:"date-time"`
	JSON    bucketJobGetResponseR2PrefixDeleteJobJSON `json:"-"`
}

// bucketJobGetResponseR2PrefixDeleteJobJSON contains the JSON metadata for the
// struct [BucketJobGetResponseR2PrefixDeleteJob]
type bucketJobGetResponseR2PrefixDeleteJobJSON struct {
	ID           apijson.Field
	JobType      apijson.Field
	PrefixDelete apijson.Field
	StartTime    apijson.Field
	Status       apijson.Field
	EndTime      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *BucketJobGetResponseR2PrefixDeleteJob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketJobGetResponseR2PrefixDeleteJobJSON) RawJSON() string {
	return r.raw
}

func (r BucketJobGetResponseR2PrefixDeleteJob) implementsBucketJobGetResponse() {}

// The job kind.
type BucketJobGetResponseR2PrefixDeleteJobJobType string

const (
	BucketJobGetResponseR2PrefixDeleteJobJobTypePrefixDelete BucketJobGetResponseR2PrefixDeleteJobJobType = "prefixDelete"
)

func (r BucketJobGetResponseR2PrefixDeleteJobJobType) IsKnown() bool {
	switch r {
	case BucketJobGetResponseR2PrefixDeleteJobJobTypePrefixDelete:
		return true
	}
	return false
}

// Details specific to the prefix-delete job.
type BucketJobGetResponseR2PrefixDeleteJobPrefixDelete struct {
	// Number of objects deleted by the job so far.
	DeletedObjects int64 `json:"deletedObjects" api:"required"`
	// The prefix matched against object keys. An empty string means the entire bucket.
	Prefix string `json:"prefix" api:"required"`
	// True when the job was created to clear the entire bucket. This field can be
	// absent on jobs created before Empty Bucket job tracking was introduced.
	IsBucketClear bool                                                  `json:"isBucketClear"`
	JSON          bucketJobGetResponseR2PrefixDeleteJobPrefixDeleteJSON `json:"-"`
}

// bucketJobGetResponseR2PrefixDeleteJobPrefixDeleteJSON contains the JSON metadata
// for the struct [BucketJobGetResponseR2PrefixDeleteJobPrefixDelete]
type bucketJobGetResponseR2PrefixDeleteJobPrefixDeleteJSON struct {
	DeletedObjects apijson.Field
	Prefix         apijson.Field
	IsBucketClear  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BucketJobGetResponseR2PrefixDeleteJobPrefixDelete) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketJobGetResponseR2PrefixDeleteJobPrefixDeleteJSON) RawJSON() string {
	return r.raw
}

// Lifecycle status of a bucket background job.
type BucketJobGetResponseR2PrefixDeleteJobStatus string

const (
	BucketJobGetResponseR2PrefixDeleteJobStatusEnqueued  BucketJobGetResponseR2PrefixDeleteJobStatus = "ENQUEUED"
	BucketJobGetResponseR2PrefixDeleteJobStatusRunning   BucketJobGetResponseR2PrefixDeleteJobStatus = "RUNNING"
	BucketJobGetResponseR2PrefixDeleteJobStatusCompleted BucketJobGetResponseR2PrefixDeleteJobStatus = "COMPLETED"
	BucketJobGetResponseR2PrefixDeleteJobStatusFailed    BucketJobGetResponseR2PrefixDeleteJobStatus = "FAILED"
	BucketJobGetResponseR2PrefixDeleteJobStatusCancelled BucketJobGetResponseR2PrefixDeleteJobStatus = "CANCELLED"
)

func (r BucketJobGetResponseR2PrefixDeleteJobStatus) IsKnown() bool {
	switch r {
	case BucketJobGetResponseR2PrefixDeleteJobStatusEnqueued, BucketJobGetResponseR2PrefixDeleteJobStatusRunning, BucketJobGetResponseR2PrefixDeleteJobStatusCompleted, BucketJobGetResponseR2PrefixDeleteJobStatusFailed, BucketJobGetResponseR2PrefixDeleteJobStatusCancelled:
		return true
	}
	return false
}

// Descriptor of a bucket-wide storage-class migration job.
type BucketJobGetResponseR2StorageClassMigrationJob struct {
	// Unique identifier used to poll the job.
	ID string `json:"id" api:"required"`
	// The job kind.
	JobType BucketJobGetResponseR2StorageClassMigrationJobJobType `json:"jobType" api:"required"`
	// When the job was created.
	StartTime time.Time `json:"startTime" api:"required" format:"date-time"`
	// Lifecycle status of a bucket background job.
	Status                BucketJobGetResponseR2StorageClassMigrationJobStatus                `json:"status" api:"required"`
	StorageClassMigration BucketJobGetResponseR2StorageClassMigrationJobStorageClassMigration `json:"storageClassMigration" api:"required"`
	// When the job finished. Absent while the job is still `ENQUEUED` or `RUNNING`.
	EndTime time.Time                                          `json:"endTime" format:"date-time"`
	JSON    bucketJobGetResponseR2StorageClassMigrationJobJSON `json:"-"`
}

// bucketJobGetResponseR2StorageClassMigrationJobJSON contains the JSON metadata
// for the struct [BucketJobGetResponseR2StorageClassMigrationJob]
type bucketJobGetResponseR2StorageClassMigrationJobJSON struct {
	ID                    apijson.Field
	JobType               apijson.Field
	StartTime             apijson.Field
	Status                apijson.Field
	StorageClassMigration apijson.Field
	EndTime               apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *BucketJobGetResponseR2StorageClassMigrationJob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketJobGetResponseR2StorageClassMigrationJobJSON) RawJSON() string {
	return r.raw
}

func (r BucketJobGetResponseR2StorageClassMigrationJob) implementsBucketJobGetResponse() {}

// The job kind.
type BucketJobGetResponseR2StorageClassMigrationJobJobType string

const (
	BucketJobGetResponseR2StorageClassMigrationJobJobTypeStorageClassMigration BucketJobGetResponseR2StorageClassMigrationJobJobType = "storageClassMigration"
)

func (r BucketJobGetResponseR2StorageClassMigrationJobJobType) IsKnown() bool {
	switch r {
	case BucketJobGetResponseR2StorageClassMigrationJobJobTypeStorageClassMigration:
		return true
	}
	return false
}

// Lifecycle status of a bucket background job.
type BucketJobGetResponseR2StorageClassMigrationJobStatus string

const (
	BucketJobGetResponseR2StorageClassMigrationJobStatusEnqueued  BucketJobGetResponseR2StorageClassMigrationJobStatus = "ENQUEUED"
	BucketJobGetResponseR2StorageClassMigrationJobStatusRunning   BucketJobGetResponseR2StorageClassMigrationJobStatus = "RUNNING"
	BucketJobGetResponseR2StorageClassMigrationJobStatusCompleted BucketJobGetResponseR2StorageClassMigrationJobStatus = "COMPLETED"
	BucketJobGetResponseR2StorageClassMigrationJobStatusFailed    BucketJobGetResponseR2StorageClassMigrationJobStatus = "FAILED"
	BucketJobGetResponseR2StorageClassMigrationJobStatusCancelled BucketJobGetResponseR2StorageClassMigrationJobStatus = "CANCELLED"
)

func (r BucketJobGetResponseR2StorageClassMigrationJobStatus) IsKnown() bool {
	switch r {
	case BucketJobGetResponseR2StorageClassMigrationJobStatusEnqueued, BucketJobGetResponseR2StorageClassMigrationJobStatusRunning, BucketJobGetResponseR2StorageClassMigrationJobStatusCompleted, BucketJobGetResponseR2StorageClassMigrationJobStatusFailed, BucketJobGetResponseR2StorageClassMigrationJobStatusCancelled:
		return true
	}
	return false
}

type BucketJobGetResponseR2StorageClassMigrationJobStorageClassMigration struct {
	// The storage class objects are migrated to. The source and destination storage
	// classes always differ.
	DestinationStorageClass BucketJobGetResponseR2StorageClassMigrationJobStorageClassMigrationDestinationStorageClass `json:"destinationStorageClass" api:"required"`
	// Number of object versions migrated successfully.
	MigratedObjects int64 `json:"migratedObjects" api:"required"`
	// The storage class objects are migrated from.
	SourceStorageClass BucketJobGetResponseR2StorageClassMigrationJobStorageClassMigrationSourceStorageClass `json:"sourceStorageClass" api:"required"`
	// Payload bytes belonging to failed object versions.
	FailedBytes int64 `json:"failedBytes"`
	// Number of object versions that failed migration.
	FailedObjects int64 `json:"failedObjects"`
	// Object payload bytes migrated successfully. Present for Metadata Worker jobs.
	MigratedBytes int64 `json:"migratedBytes"`
	// Object metadata bytes migrated successfully. Present for Metadata Worker jobs.
	MigratedMetadataBytes int64 `json:"migratedMetadataBytes"`
	// Payload bytes belonging to skipped object versions.
	SkippedBytes int64 `json:"skippedBytes"`
	// Number of object versions skipped.
	SkippedObjects int64                                                                   `json:"skippedObjects"`
	JSON           bucketJobGetResponseR2StorageClassMigrationJobStorageClassMigrationJSON `json:"-"`
}

// bucketJobGetResponseR2StorageClassMigrationJobStorageClassMigrationJSON contains
// the JSON metadata for the struct
// [BucketJobGetResponseR2StorageClassMigrationJobStorageClassMigration]
type bucketJobGetResponseR2StorageClassMigrationJobStorageClassMigrationJSON struct {
	DestinationStorageClass apijson.Field
	MigratedObjects         apijson.Field
	SourceStorageClass      apijson.Field
	FailedBytes             apijson.Field
	FailedObjects           apijson.Field
	MigratedBytes           apijson.Field
	MigratedMetadataBytes   apijson.Field
	SkippedBytes            apijson.Field
	SkippedObjects          apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *BucketJobGetResponseR2StorageClassMigrationJobStorageClassMigration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketJobGetResponseR2StorageClassMigrationJobStorageClassMigrationJSON) RawJSON() string {
	return r.raw
}

// The storage class objects are migrated to. The source and destination storage
// classes always differ.
type BucketJobGetResponseR2StorageClassMigrationJobStorageClassMigrationDestinationStorageClass string

const (
	BucketJobGetResponseR2StorageClassMigrationJobStorageClassMigrationDestinationStorageClassStandard         BucketJobGetResponseR2StorageClassMigrationJobStorageClassMigrationDestinationStorageClass = "Standard"
	BucketJobGetResponseR2StorageClassMigrationJobStorageClassMigrationDestinationStorageClassInfrequentAccess BucketJobGetResponseR2StorageClassMigrationJobStorageClassMigrationDestinationStorageClass = "InfrequentAccess"
)

func (r BucketJobGetResponseR2StorageClassMigrationJobStorageClassMigrationDestinationStorageClass) IsKnown() bool {
	switch r {
	case BucketJobGetResponseR2StorageClassMigrationJobStorageClassMigrationDestinationStorageClassStandard, BucketJobGetResponseR2StorageClassMigrationJobStorageClassMigrationDestinationStorageClassInfrequentAccess:
		return true
	}
	return false
}

// The storage class objects are migrated from.
type BucketJobGetResponseR2StorageClassMigrationJobStorageClassMigrationSourceStorageClass string

const (
	BucketJobGetResponseR2StorageClassMigrationJobStorageClassMigrationSourceStorageClassStandard         BucketJobGetResponseR2StorageClassMigrationJobStorageClassMigrationSourceStorageClass = "Standard"
	BucketJobGetResponseR2StorageClassMigrationJobStorageClassMigrationSourceStorageClassInfrequentAccess BucketJobGetResponseR2StorageClassMigrationJobStorageClassMigrationSourceStorageClass = "InfrequentAccess"
)

func (r BucketJobGetResponseR2StorageClassMigrationJobStorageClassMigrationSourceStorageClass) IsKnown() bool {
	switch r {
	case BucketJobGetResponseR2StorageClassMigrationJobStorageClassMigrationSourceStorageClassStandard, BucketJobGetResponseR2StorageClassMigrationJobStorageClassMigrationSourceStorageClassInfrequentAccess:
		return true
	}
	return false
}

// The job kind.
type BucketJobGetResponseJobType string

const (
	BucketJobGetResponseJobTypePrefixDelete          BucketJobGetResponseJobType = "prefixDelete"
	BucketJobGetResponseJobTypeStorageClassMigration BucketJobGetResponseJobType = "storageClassMigration"
)

func (r BucketJobGetResponseJobType) IsKnown() bool {
	switch r {
	case BucketJobGetResponseJobTypePrefixDelete, BucketJobGetResponseJobTypeStorageClassMigration:
		return true
	}
	return false
}

// Lifecycle status of a bucket background job.
type BucketJobGetResponseStatus string

const (
	BucketJobGetResponseStatusEnqueued  BucketJobGetResponseStatus = "ENQUEUED"
	BucketJobGetResponseStatusRunning   BucketJobGetResponseStatus = "RUNNING"
	BucketJobGetResponseStatusCompleted BucketJobGetResponseStatus = "COMPLETED"
	BucketJobGetResponseStatusFailed    BucketJobGetResponseStatus = "FAILED"
	BucketJobGetResponseStatusCancelled BucketJobGetResponseStatus = "CANCELLED"
)

func (r BucketJobGetResponseStatus) IsKnown() bool {
	switch r {
	case BucketJobGetResponseStatusEnqueued, BucketJobGetResponseStatusRunning, BucketJobGetResponseStatusCompleted, BucketJobGetResponseStatusFailed, BucketJobGetResponseStatusCancelled:
		return true
	}
	return false
}

type BucketJobNewParams struct {
	// Cloudflare account ID that owns the R2 resource.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Request body that creates a prefix-delete job.
	Body                 BucketJobNewParamsBodyUnion                     `json:"body" api:"required"`
	CfR2DataCatalogCheck param.Field[string]                             `header:"cf-r2-data-catalog-check"`
	CfR2Jurisdiction     param.Field[BucketJobNewParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

func (r BucketJobNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Request body that creates a prefix-delete job.
type BucketJobNewParamsBody struct {
	JobType                 param.Field[BucketJobNewParamsBodyJobType]                 `json:"jobType" api:"required"`
	DestinationStorageClass param.Field[BucketJobNewParamsBodyDestinationStorageClass] `json:"destinationStorageClass"`
	// Prefix matched against object keys. A non-empty value must end in `/`. An empty
	// string empties the entire bucket.
	Prefix             param.Field[string]                                   `json:"prefix"`
	SourceStorageClass param.Field[BucketJobNewParamsBodySourceStorageClass] `json:"sourceStorageClass"`
}

func (r BucketJobNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BucketJobNewParamsBody) implementsBucketJobNewParamsBodyUnion() {}

// Request body that creates a prefix-delete job.
//
// Satisfied by [r2.BucketJobNewParamsBodyR2CreatePrefixDeleteJobRequest],
// [r2.BucketJobNewParamsBodyR2CreateStorageClassMigrationJobRequest],
// [BucketJobNewParamsBody].
type BucketJobNewParamsBodyUnion interface {
	implementsBucketJobNewParamsBodyUnion()
}

// Request body that creates a prefix-delete job.
type BucketJobNewParamsBodyR2CreatePrefixDeleteJobRequest struct {
	JobType param.Field[BucketJobNewParamsBodyR2CreatePrefixDeleteJobRequestJobType] `json:"jobType" api:"required"`
	// Prefix matched against object keys. A non-empty value must end in `/`. An empty
	// string empties the entire bucket.
	Prefix param.Field[string] `json:"prefix" api:"required"`
}

func (r BucketJobNewParamsBodyR2CreatePrefixDeleteJobRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BucketJobNewParamsBodyR2CreatePrefixDeleteJobRequest) implementsBucketJobNewParamsBodyUnion() {
}

type BucketJobNewParamsBodyR2CreatePrefixDeleteJobRequestJobType string

const (
	BucketJobNewParamsBodyR2CreatePrefixDeleteJobRequestJobTypePrefixDelete BucketJobNewParamsBodyR2CreatePrefixDeleteJobRequestJobType = "prefixDelete"
)

func (r BucketJobNewParamsBodyR2CreatePrefixDeleteJobRequestJobType) IsKnown() bool {
	switch r {
	case BucketJobNewParamsBodyR2CreatePrefixDeleteJobRequestJobTypePrefixDelete:
		return true
	}
	return false
}

// Request body that creates a storage-class migration job.
type BucketJobNewParamsBodyR2CreateStorageClassMigrationJobRequest struct {
	JobType                 param.Field[BucketJobNewParamsBodyR2CreateStorageClassMigrationJobRequestJobType]                 `json:"jobType" api:"required"`
	DestinationStorageClass param.Field[BucketJobNewParamsBodyR2CreateStorageClassMigrationJobRequestDestinationStorageClass] `json:"destinationStorageClass"`
	SourceStorageClass      param.Field[BucketJobNewParamsBodyR2CreateStorageClassMigrationJobRequestSourceStorageClass]      `json:"sourceStorageClass"`
}

func (r BucketJobNewParamsBodyR2CreateStorageClassMigrationJobRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BucketJobNewParamsBodyR2CreateStorageClassMigrationJobRequest) implementsBucketJobNewParamsBodyUnion() {
}

type BucketJobNewParamsBodyR2CreateStorageClassMigrationJobRequestJobType string

const (
	BucketJobNewParamsBodyR2CreateStorageClassMigrationJobRequestJobTypeStorageClassMigration BucketJobNewParamsBodyR2CreateStorageClassMigrationJobRequestJobType = "storageClassMigration"
)

func (r BucketJobNewParamsBodyR2CreateStorageClassMigrationJobRequestJobType) IsKnown() bool {
	switch r {
	case BucketJobNewParamsBodyR2CreateStorageClassMigrationJobRequestJobTypeStorageClassMigration:
		return true
	}
	return false
}

type BucketJobNewParamsBodyR2CreateStorageClassMigrationJobRequestDestinationStorageClass string

const (
	BucketJobNewParamsBodyR2CreateStorageClassMigrationJobRequestDestinationStorageClassStandard         BucketJobNewParamsBodyR2CreateStorageClassMigrationJobRequestDestinationStorageClass = "Standard"
	BucketJobNewParamsBodyR2CreateStorageClassMigrationJobRequestDestinationStorageClassInfrequentAccess BucketJobNewParamsBodyR2CreateStorageClassMigrationJobRequestDestinationStorageClass = "InfrequentAccess"
)

func (r BucketJobNewParamsBodyR2CreateStorageClassMigrationJobRequestDestinationStorageClass) IsKnown() bool {
	switch r {
	case BucketJobNewParamsBodyR2CreateStorageClassMigrationJobRequestDestinationStorageClassStandard, BucketJobNewParamsBodyR2CreateStorageClassMigrationJobRequestDestinationStorageClassInfrequentAccess:
		return true
	}
	return false
}

type BucketJobNewParamsBodyR2CreateStorageClassMigrationJobRequestSourceStorageClass string

const (
	BucketJobNewParamsBodyR2CreateStorageClassMigrationJobRequestSourceStorageClassStandard         BucketJobNewParamsBodyR2CreateStorageClassMigrationJobRequestSourceStorageClass = "Standard"
	BucketJobNewParamsBodyR2CreateStorageClassMigrationJobRequestSourceStorageClassInfrequentAccess BucketJobNewParamsBodyR2CreateStorageClassMigrationJobRequestSourceStorageClass = "InfrequentAccess"
)

func (r BucketJobNewParamsBodyR2CreateStorageClassMigrationJobRequestSourceStorageClass) IsKnown() bool {
	switch r {
	case BucketJobNewParamsBodyR2CreateStorageClassMigrationJobRequestSourceStorageClassStandard, BucketJobNewParamsBodyR2CreateStorageClassMigrationJobRequestSourceStorageClassInfrequentAccess:
		return true
	}
	return false
}

type BucketJobNewParamsBodyJobType string

const (
	BucketJobNewParamsBodyJobTypePrefixDelete          BucketJobNewParamsBodyJobType = "prefixDelete"
	BucketJobNewParamsBodyJobTypeStorageClassMigration BucketJobNewParamsBodyJobType = "storageClassMigration"
)

func (r BucketJobNewParamsBodyJobType) IsKnown() bool {
	switch r {
	case BucketJobNewParamsBodyJobTypePrefixDelete, BucketJobNewParamsBodyJobTypeStorageClassMigration:
		return true
	}
	return false
}

type BucketJobNewParamsBodyDestinationStorageClass string

const (
	BucketJobNewParamsBodyDestinationStorageClassStandard         BucketJobNewParamsBodyDestinationStorageClass = "Standard"
	BucketJobNewParamsBodyDestinationStorageClassInfrequentAccess BucketJobNewParamsBodyDestinationStorageClass = "InfrequentAccess"
)

func (r BucketJobNewParamsBodyDestinationStorageClass) IsKnown() bool {
	switch r {
	case BucketJobNewParamsBodyDestinationStorageClassStandard, BucketJobNewParamsBodyDestinationStorageClassInfrequentAccess:
		return true
	}
	return false
}

type BucketJobNewParamsBodySourceStorageClass string

const (
	BucketJobNewParamsBodySourceStorageClassStandard         BucketJobNewParamsBodySourceStorageClass = "Standard"
	BucketJobNewParamsBodySourceStorageClassInfrequentAccess BucketJobNewParamsBodySourceStorageClass = "InfrequentAccess"
)

func (r BucketJobNewParamsBodySourceStorageClass) IsKnown() bool {
	switch r {
	case BucketJobNewParamsBodySourceStorageClassStandard, BucketJobNewParamsBodySourceStorageClassInfrequentAccess:
		return true
	}
	return false
}

type BucketJobNewParamsCfR2Jurisdiction string

const (
	BucketJobNewParamsCfR2JurisdictionDefault     BucketJobNewParamsCfR2Jurisdiction = "default"
	BucketJobNewParamsCfR2JurisdictionEu          BucketJobNewParamsCfR2Jurisdiction = "eu"
	BucketJobNewParamsCfR2JurisdictionUs          BucketJobNewParamsCfR2Jurisdiction = "us"
	BucketJobNewParamsCfR2JurisdictionFedramp     BucketJobNewParamsCfR2Jurisdiction = "fedramp"
	BucketJobNewParamsCfR2JurisdictionFedrampHigh BucketJobNewParamsCfR2Jurisdiction = "fedramp-high"
)

func (r BucketJobNewParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketJobNewParamsCfR2JurisdictionDefault, BucketJobNewParamsCfR2JurisdictionEu, BucketJobNewParamsCfR2JurisdictionUs, BucketJobNewParamsCfR2JurisdictionFedramp, BucketJobNewParamsCfR2JurisdictionFedrampHigh:
		return true
	}
	return false
}

type BucketJobNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []string              `json:"messages" api:"required"`
	// Descriptor of a bucket background job, discriminated by `jobType`. Returned by
	// the Create Bucket Job, Get Bucket Job and List Bucket Jobs endpoints. A small
	// operation can already be `COMPLETED` when it is returned.
	Result BucketJobNewResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success BucketJobNewResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    bucketJobNewResponseEnvelopeJSON    `json:"-"`
}

// bucketJobNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [BucketJobNewResponseEnvelope]
type bucketJobNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketJobNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketJobNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketJobNewResponseEnvelopeSuccess bool

const (
	BucketJobNewResponseEnvelopeSuccessTrue BucketJobNewResponseEnvelopeSuccess = true
)

func (r BucketJobNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketJobNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BucketJobListParams struct {
	// Cloudflare account ID that owns the R2 resource.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Pagination token received as `nextContinuationToken` in the previous response.
	ContinuationToken param.Field[string] `query:"continuationToken"`
	// Restricts results to jobs of the specified type.
	JobType param.Field[BucketJobListParamsJobType] `query:"jobType"`
	// Maximum number of jobs to return.
	MaxKeys param.Field[int64] `query:"maxKeys"`
	// Restricts results to jobs with the specified status. `jobType` is required when
	// this parameter is provided.
	Status           param.Field[BucketJobListParamsStatus]           `query:"status"`
	CfR2Jurisdiction param.Field[BucketJobListParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// URLQuery serializes [BucketJobListParams]'s query parameters as `url.Values`.
func (r BucketJobListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Restricts results to jobs of the specified type.
type BucketJobListParamsJobType string

const (
	BucketJobListParamsJobTypePrefixDelete          BucketJobListParamsJobType = "prefixDelete"
	BucketJobListParamsJobTypeStorageClassMigration BucketJobListParamsJobType = "storageClassMigration"
)

func (r BucketJobListParamsJobType) IsKnown() bool {
	switch r {
	case BucketJobListParamsJobTypePrefixDelete, BucketJobListParamsJobTypeStorageClassMigration:
		return true
	}
	return false
}

// Restricts results to jobs with the specified status. `jobType` is required when
// this parameter is provided.
type BucketJobListParamsStatus string

const (
	BucketJobListParamsStatusEnqueued  BucketJobListParamsStatus = "ENQUEUED"
	BucketJobListParamsStatusRunning   BucketJobListParamsStatus = "RUNNING"
	BucketJobListParamsStatusCompleted BucketJobListParamsStatus = "COMPLETED"
	BucketJobListParamsStatusFailed    BucketJobListParamsStatus = "FAILED"
	BucketJobListParamsStatusCancelled BucketJobListParamsStatus = "CANCELLED"
)

func (r BucketJobListParamsStatus) IsKnown() bool {
	switch r {
	case BucketJobListParamsStatusEnqueued, BucketJobListParamsStatusRunning, BucketJobListParamsStatusCompleted, BucketJobListParamsStatusFailed, BucketJobListParamsStatusCancelled:
		return true
	}
	return false
}

type BucketJobListParamsCfR2Jurisdiction string

const (
	BucketJobListParamsCfR2JurisdictionDefault     BucketJobListParamsCfR2Jurisdiction = "default"
	BucketJobListParamsCfR2JurisdictionEu          BucketJobListParamsCfR2Jurisdiction = "eu"
	BucketJobListParamsCfR2JurisdictionUs          BucketJobListParamsCfR2Jurisdiction = "us"
	BucketJobListParamsCfR2JurisdictionFedramp     BucketJobListParamsCfR2Jurisdiction = "fedramp"
	BucketJobListParamsCfR2JurisdictionFedrampHigh BucketJobListParamsCfR2Jurisdiction = "fedramp-high"
)

func (r BucketJobListParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketJobListParamsCfR2JurisdictionDefault, BucketJobListParamsCfR2JurisdictionEu, BucketJobListParamsCfR2JurisdictionUs, BucketJobListParamsCfR2JurisdictionFedramp, BucketJobListParamsCfR2JurisdictionFedrampHigh:
		return true
	}
	return false
}

type BucketJobListResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []string              `json:"messages" api:"required"`
	Result   BucketJobListResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success BucketJobListResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    bucketJobListResponseEnvelopeJSON    `json:"-"`
}

// bucketJobListResponseEnvelopeJSON contains the JSON metadata for the struct
// [BucketJobListResponseEnvelope]
type bucketJobListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketJobListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketJobListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketJobListResponseEnvelopeSuccess bool

const (
	BucketJobListResponseEnvelopeSuccessTrue BucketJobListResponseEnvelopeSuccess = true
)

func (r BucketJobListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketJobListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BucketJobGetParams struct {
	// Cloudflare account ID that owns the R2 resource.
	AccountID        param.Field[string]                             `path:"account_id" api:"required"`
	CfR2Jurisdiction param.Field[BucketJobGetParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

type BucketJobGetParamsCfR2Jurisdiction string

const (
	BucketJobGetParamsCfR2JurisdictionDefault     BucketJobGetParamsCfR2Jurisdiction = "default"
	BucketJobGetParamsCfR2JurisdictionEu          BucketJobGetParamsCfR2Jurisdiction = "eu"
	BucketJobGetParamsCfR2JurisdictionUs          BucketJobGetParamsCfR2Jurisdiction = "us"
	BucketJobGetParamsCfR2JurisdictionFedramp     BucketJobGetParamsCfR2Jurisdiction = "fedramp"
	BucketJobGetParamsCfR2JurisdictionFedrampHigh BucketJobGetParamsCfR2Jurisdiction = "fedramp-high"
)

func (r BucketJobGetParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketJobGetParamsCfR2JurisdictionDefault, BucketJobGetParamsCfR2JurisdictionEu, BucketJobGetParamsCfR2JurisdictionUs, BucketJobGetParamsCfR2JurisdictionFedramp, BucketJobGetParamsCfR2JurisdictionFedrampHigh:
		return true
	}
	return false
}

type BucketJobGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []string              `json:"messages" api:"required"`
	// Descriptor of a bucket background job, discriminated by `jobType`. Returned by
	// the Create Bucket Job, Get Bucket Job and List Bucket Jobs endpoints. A small
	// operation can already be `COMPLETED` when it is returned.
	Result BucketJobGetResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success BucketJobGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    bucketJobGetResponseEnvelopeJSON    `json:"-"`
}

// bucketJobGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [BucketJobGetResponseEnvelope]
type bucketJobGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketJobGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketJobGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketJobGetResponseEnvelopeSuccess bool

const (
	BucketJobGetResponseEnvelopeSuccessTrue BucketJobGetResponseEnvelopeSuccess = true
)

func (r BucketJobGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketJobGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
