// File generated from our OpenAPI spec by Stainless.

package logpush

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
)

// JobService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewJobService] method instead.
type JobService struct {
	Options []option.RequestOption
}

// NewJobService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewJobService(opts ...option.RequestOption) (r *JobService) {
	r = &JobService{}
	r.Options = opts
	return
}

// Creates a new Logpush job for an account or zone.
func (r *JobService) New(ctx context.Context, params JobNewParams, opts ...option.RequestOption) (res *JobNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env JobNewResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/logpush/jobs", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a Logpush job.
func (r *JobService) Update(ctx context.Context, jobID int64, params JobUpdateParams, opts ...option.RequestOption) (res *JobUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env JobUpdateResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/logpush/jobs/%v", accountOrZone, accountOrZoneID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists Logpush jobs for an account or zone.
func (r *JobService) List(ctx context.Context, query JobListParams, opts ...option.RequestOption) (res *[]JobListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env JobListResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/logpush/jobs", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a Logpush job.
func (r *JobService) Delete(ctx context.Context, jobID int64, body JobDeleteParams, opts ...option.RequestOption) (res *JobDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env JobDeleteResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	path := fmt.Sprintf("%s/%s/logpush/jobs/%v", accountOrZone, accountOrZoneID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets the details of a Logpush job.
func (r *JobService) Get(ctx context.Context, jobID int64, query JobGetParams, opts ...option.RequestOption) (res *JobGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env JobGetResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/logpush/jobs/%v", accountOrZone, accountOrZoneID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type JobNewResponse struct {
	// Unique id of the job.
	ID int64 `json:"id"`
	// Name of the dataset.
	Dataset string `json:"dataset,nullable"`
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf string `json:"destination_conf" format:"uri"`
	// Flag that indicates if the job is enabled.
	Enabled bool `json:"enabled"`
	// If not null, the job is currently failing. Failures are usually repetitive
	// (example: no permissions to write to destination bucket). Only the last failure
	// is recorded. On successful execution of a job the error_message and last_error
	// are set to null.
	ErrorMessage time.Time `json:"error_message,nullable" format:"date-time"`
	// The frequency at which Cloudflare sends batches of logs to your destination.
	// Setting frequency to high sends your logs in larger quantities of smaller files.
	// Setting frequency to low sends logs in smaller quantities of larger files.
	Frequency JobNewResponseFrequency `json:"frequency,nullable"`
	// Records the last time for which logs have been successfully pushed. If the last
	// successful push was for logs range 2018-07-23T10:00:00Z to 2018-07-23T10:01:00Z
	// then the value of this field will be 2018-07-23T10:01:00Z. If the job has never
	// run or has just been enabled and hasn't run yet then the field will be empty.
	LastComplete time.Time `json:"last_complete,nullable" format:"date-time"`
	// Records the last time the job failed. If not null, the job is currently failing.
	// If null, the job has either never failed or has run successfully at least once
	// since last failure. See also the error_message field.
	LastError time.Time `json:"last_error,nullable" format:"date-time"`
	// This field is deprecated. Use `output_options` instead. Configuration string. It
	// specifies things like requested fields and timestamp formats. If migrating from
	// the logpull api, copy the url (full url or just the query string) of your call
	// here, and logpush will keep on making this call for you, setting start and end
	// times appropriately.
	LogpullOptions string `json:"logpull_options,nullable" format:"uri-reference"`
	// Optional human readable job name. Not unique. Cloudflare suggests that you set
	// this to a meaningful string, like the domain name, to make it easier to identify
	// your job.
	Name string `json:"name,nullable"`
	// The structured replacement for `logpull_options`. When including this field, the
	// `logpull_option` field will be ignored.
	OutputOptions JobNewResponseOutputOptions `json:"output_options,nullable"`
	JSON          jobNewResponseJSON          `json:"-"`
}

// jobNewResponseJSON contains the JSON metadata for the struct [JobNewResponse]
type jobNewResponseJSON struct {
	ID              apijson.Field
	Dataset         apijson.Field
	DestinationConf apijson.Field
	Enabled         apijson.Field
	ErrorMessage    apijson.Field
	Frequency       apijson.Field
	LastComplete    apijson.Field
	LastError       apijson.Field
	LogpullOptions  apijson.Field
	Name            apijson.Field
	OutputOptions   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *JobNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobNewResponseJSON) RawJSON() string {
	return r.raw
}

// The frequency at which Cloudflare sends batches of logs to your destination.
// Setting frequency to high sends your logs in larger quantities of smaller files.
// Setting frequency to low sends logs in smaller quantities of larger files.
type JobNewResponseFrequency string

const (
	JobNewResponseFrequencyHigh JobNewResponseFrequency = "high"
	JobNewResponseFrequencyLow  JobNewResponseFrequency = "low"
)

// The structured replacement for `logpull_options`. When including this field, the
// `logpull_option` field will be ignored.
type JobNewResponseOutputOptions struct {
	// String to be prepended before each batch.
	BatchPrefix string `json:"batch_prefix,nullable"`
	// String to be appended after each batch.
	BatchSuffix string `json:"batch_suffix,nullable"`
	// If set to true, will cause all occurrences of `${` in the generated files to be
	// replaced with `x{`.
	Cve2021_4428 bool `json:"CVE-2021-4428,nullable"`
	// String to join fields. This field be ignored when `record_template` is set.
	FieldDelimiter string `json:"field_delimiter,nullable"`
	// List of field names to be included in the Logpush output. For the moment, there
	// is no option to add all fields at once, so you must specify all the fields names
	// you are interested in.
	FieldNames []string `json:"field_names"`
	// Specifies the output type, such as `ndjson` or `csv`. This sets default values
	// for the rest of the settings, depending on the chosen output type. Some
	// formatting rules, like string quoting, are different between output types.
	OutputType JobNewResponseOutputOptionsOutputType `json:"output_type"`
	// String to be inserted in-between the records as separator.
	RecordDelimiter string `json:"record_delimiter,nullable"`
	// String to be prepended before each record.
	RecordPrefix string `json:"record_prefix,nullable"`
	// String to be appended after each record.
	RecordSuffix string `json:"record_suffix,nullable"`
	// String to use as template for each record instead of the default comma-separated
	// list. All fields used in the template must be present in `field_names` as well,
	// otherwise they will end up as null. Format as a Go `text/template` without any
	// standard functions, like conditionals, loops, sub-templates, etc.
	RecordTemplate string `json:"record_template,nullable"`
	// Floating number to specify sampling rate. Sampling is applied on top of
	// filtering, and regardless of the current `sample_interval` of the data.
	SampleRate float64 `json:"sample_rate,nullable"`
	// String to specify the format for timestamps, such as `unixnano`, `unix`, or
	// `rfc3339`.
	TimestampFormat JobNewResponseOutputOptionsTimestampFormat `json:"timestamp_format"`
	JSON            jobNewResponseOutputOptionsJSON            `json:"-"`
}

// jobNewResponseOutputOptionsJSON contains the JSON metadata for the struct
// [JobNewResponseOutputOptions]
type jobNewResponseOutputOptionsJSON struct {
	BatchPrefix     apijson.Field
	BatchSuffix     apijson.Field
	Cve2021_4428    apijson.Field
	FieldDelimiter  apijson.Field
	FieldNames      apijson.Field
	OutputType      apijson.Field
	RecordDelimiter apijson.Field
	RecordPrefix    apijson.Field
	RecordSuffix    apijson.Field
	RecordTemplate  apijson.Field
	SampleRate      apijson.Field
	TimestampFormat apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *JobNewResponseOutputOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobNewResponseOutputOptionsJSON) RawJSON() string {
	return r.raw
}

// Specifies the output type, such as `ndjson` or `csv`. This sets default values
// for the rest of the settings, depending on the chosen output type. Some
// formatting rules, like string quoting, are different between output types.
type JobNewResponseOutputOptionsOutputType string

const (
	JobNewResponseOutputOptionsOutputTypeNdjson JobNewResponseOutputOptionsOutputType = "ndjson"
	JobNewResponseOutputOptionsOutputTypeCsv    JobNewResponseOutputOptionsOutputType = "csv"
)

// String to specify the format for timestamps, such as `unixnano`, `unix`, or
// `rfc3339`.
type JobNewResponseOutputOptionsTimestampFormat string

const (
	JobNewResponseOutputOptionsTimestampFormatUnixnano JobNewResponseOutputOptionsTimestampFormat = "unixnano"
	JobNewResponseOutputOptionsTimestampFormatUnix     JobNewResponseOutputOptionsTimestampFormat = "unix"
	JobNewResponseOutputOptionsTimestampFormatRfc3339  JobNewResponseOutputOptionsTimestampFormat = "rfc3339"
)

type JobUpdateResponse struct {
	// Unique id of the job.
	ID int64 `json:"id"`
	// Name of the dataset.
	Dataset string `json:"dataset,nullable"`
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf string `json:"destination_conf" format:"uri"`
	// Flag that indicates if the job is enabled.
	Enabled bool `json:"enabled"`
	// If not null, the job is currently failing. Failures are usually repetitive
	// (example: no permissions to write to destination bucket). Only the last failure
	// is recorded. On successful execution of a job the error_message and last_error
	// are set to null.
	ErrorMessage time.Time `json:"error_message,nullable" format:"date-time"`
	// The frequency at which Cloudflare sends batches of logs to your destination.
	// Setting frequency to high sends your logs in larger quantities of smaller files.
	// Setting frequency to low sends logs in smaller quantities of larger files.
	Frequency JobUpdateResponseFrequency `json:"frequency,nullable"`
	// Records the last time for which logs have been successfully pushed. If the last
	// successful push was for logs range 2018-07-23T10:00:00Z to 2018-07-23T10:01:00Z
	// then the value of this field will be 2018-07-23T10:01:00Z. If the job has never
	// run or has just been enabled and hasn't run yet then the field will be empty.
	LastComplete time.Time `json:"last_complete,nullable" format:"date-time"`
	// Records the last time the job failed. If not null, the job is currently failing.
	// If null, the job has either never failed or has run successfully at least once
	// since last failure. See also the error_message field.
	LastError time.Time `json:"last_error,nullable" format:"date-time"`
	// This field is deprecated. Use `output_options` instead. Configuration string. It
	// specifies things like requested fields and timestamp formats. If migrating from
	// the logpull api, copy the url (full url or just the query string) of your call
	// here, and logpush will keep on making this call for you, setting start and end
	// times appropriately.
	LogpullOptions string `json:"logpull_options,nullable" format:"uri-reference"`
	// Optional human readable job name. Not unique. Cloudflare suggests that you set
	// this to a meaningful string, like the domain name, to make it easier to identify
	// your job.
	Name string `json:"name,nullable"`
	// The structured replacement for `logpull_options`. When including this field, the
	// `logpull_option` field will be ignored.
	OutputOptions JobUpdateResponseOutputOptions `json:"output_options,nullable"`
	JSON          jobUpdateResponseJSON          `json:"-"`
}

// jobUpdateResponseJSON contains the JSON metadata for the struct
// [JobUpdateResponse]
type jobUpdateResponseJSON struct {
	ID              apijson.Field
	Dataset         apijson.Field
	DestinationConf apijson.Field
	Enabled         apijson.Field
	ErrorMessage    apijson.Field
	Frequency       apijson.Field
	LastComplete    apijson.Field
	LastError       apijson.Field
	LogpullOptions  apijson.Field
	Name            apijson.Field
	OutputOptions   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *JobUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The frequency at which Cloudflare sends batches of logs to your destination.
// Setting frequency to high sends your logs in larger quantities of smaller files.
// Setting frequency to low sends logs in smaller quantities of larger files.
type JobUpdateResponseFrequency string

const (
	JobUpdateResponseFrequencyHigh JobUpdateResponseFrequency = "high"
	JobUpdateResponseFrequencyLow  JobUpdateResponseFrequency = "low"
)

// The structured replacement for `logpull_options`. When including this field, the
// `logpull_option` field will be ignored.
type JobUpdateResponseOutputOptions struct {
	// String to be prepended before each batch.
	BatchPrefix string `json:"batch_prefix,nullable"`
	// String to be appended after each batch.
	BatchSuffix string `json:"batch_suffix,nullable"`
	// If set to true, will cause all occurrences of `${` in the generated files to be
	// replaced with `x{`.
	Cve2021_4428 bool `json:"CVE-2021-4428,nullable"`
	// String to join fields. This field be ignored when `record_template` is set.
	FieldDelimiter string `json:"field_delimiter,nullable"`
	// List of field names to be included in the Logpush output. For the moment, there
	// is no option to add all fields at once, so you must specify all the fields names
	// you are interested in.
	FieldNames []string `json:"field_names"`
	// Specifies the output type, such as `ndjson` or `csv`. This sets default values
	// for the rest of the settings, depending on the chosen output type. Some
	// formatting rules, like string quoting, are different between output types.
	OutputType JobUpdateResponseOutputOptionsOutputType `json:"output_type"`
	// String to be inserted in-between the records as separator.
	RecordDelimiter string `json:"record_delimiter,nullable"`
	// String to be prepended before each record.
	RecordPrefix string `json:"record_prefix,nullable"`
	// String to be appended after each record.
	RecordSuffix string `json:"record_suffix,nullable"`
	// String to use as template for each record instead of the default comma-separated
	// list. All fields used in the template must be present in `field_names` as well,
	// otherwise they will end up as null. Format as a Go `text/template` without any
	// standard functions, like conditionals, loops, sub-templates, etc.
	RecordTemplate string `json:"record_template,nullable"`
	// Floating number to specify sampling rate. Sampling is applied on top of
	// filtering, and regardless of the current `sample_interval` of the data.
	SampleRate float64 `json:"sample_rate,nullable"`
	// String to specify the format for timestamps, such as `unixnano`, `unix`, or
	// `rfc3339`.
	TimestampFormat JobUpdateResponseOutputOptionsTimestampFormat `json:"timestamp_format"`
	JSON            jobUpdateResponseOutputOptionsJSON            `json:"-"`
}

// jobUpdateResponseOutputOptionsJSON contains the JSON metadata for the struct
// [JobUpdateResponseOutputOptions]
type jobUpdateResponseOutputOptionsJSON struct {
	BatchPrefix     apijson.Field
	BatchSuffix     apijson.Field
	Cve2021_4428    apijson.Field
	FieldDelimiter  apijson.Field
	FieldNames      apijson.Field
	OutputType      apijson.Field
	RecordDelimiter apijson.Field
	RecordPrefix    apijson.Field
	RecordSuffix    apijson.Field
	RecordTemplate  apijson.Field
	SampleRate      apijson.Field
	TimestampFormat apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *JobUpdateResponseOutputOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobUpdateResponseOutputOptionsJSON) RawJSON() string {
	return r.raw
}

// Specifies the output type, such as `ndjson` or `csv`. This sets default values
// for the rest of the settings, depending on the chosen output type. Some
// formatting rules, like string quoting, are different between output types.
type JobUpdateResponseOutputOptionsOutputType string

const (
	JobUpdateResponseOutputOptionsOutputTypeNdjson JobUpdateResponseOutputOptionsOutputType = "ndjson"
	JobUpdateResponseOutputOptionsOutputTypeCsv    JobUpdateResponseOutputOptionsOutputType = "csv"
)

// String to specify the format for timestamps, such as `unixnano`, `unix`, or
// `rfc3339`.
type JobUpdateResponseOutputOptionsTimestampFormat string

const (
	JobUpdateResponseOutputOptionsTimestampFormatUnixnano JobUpdateResponseOutputOptionsTimestampFormat = "unixnano"
	JobUpdateResponseOutputOptionsTimestampFormatUnix     JobUpdateResponseOutputOptionsTimestampFormat = "unix"
	JobUpdateResponseOutputOptionsTimestampFormatRfc3339  JobUpdateResponseOutputOptionsTimestampFormat = "rfc3339"
)

type JobListResponse struct {
	// Unique id of the job.
	ID int64 `json:"id"`
	// Name of the dataset.
	Dataset string `json:"dataset,nullable"`
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf string `json:"destination_conf" format:"uri"`
	// Flag that indicates if the job is enabled.
	Enabled bool `json:"enabled"`
	// If not null, the job is currently failing. Failures are usually repetitive
	// (example: no permissions to write to destination bucket). Only the last failure
	// is recorded. On successful execution of a job the error_message and last_error
	// are set to null.
	ErrorMessage time.Time `json:"error_message,nullable" format:"date-time"`
	// The frequency at which Cloudflare sends batches of logs to your destination.
	// Setting frequency to high sends your logs in larger quantities of smaller files.
	// Setting frequency to low sends logs in smaller quantities of larger files.
	Frequency JobListResponseFrequency `json:"frequency,nullable"`
	// Records the last time for which logs have been successfully pushed. If the last
	// successful push was for logs range 2018-07-23T10:00:00Z to 2018-07-23T10:01:00Z
	// then the value of this field will be 2018-07-23T10:01:00Z. If the job has never
	// run or has just been enabled and hasn't run yet then the field will be empty.
	LastComplete time.Time `json:"last_complete,nullable" format:"date-time"`
	// Records the last time the job failed. If not null, the job is currently failing.
	// If null, the job has either never failed or has run successfully at least once
	// since last failure. See also the error_message field.
	LastError time.Time `json:"last_error,nullable" format:"date-time"`
	// This field is deprecated. Use `output_options` instead. Configuration string. It
	// specifies things like requested fields and timestamp formats. If migrating from
	// the logpull api, copy the url (full url or just the query string) of your call
	// here, and logpush will keep on making this call for you, setting start and end
	// times appropriately.
	LogpullOptions string `json:"logpull_options,nullable" format:"uri-reference"`
	// Optional human readable job name. Not unique. Cloudflare suggests that you set
	// this to a meaningful string, like the domain name, to make it easier to identify
	// your job.
	Name string `json:"name,nullable"`
	// The structured replacement for `logpull_options`. When including this field, the
	// `logpull_option` field will be ignored.
	OutputOptions JobListResponseOutputOptions `json:"output_options,nullable"`
	JSON          jobListResponseJSON          `json:"-"`
}

// jobListResponseJSON contains the JSON metadata for the struct [JobListResponse]
type jobListResponseJSON struct {
	ID              apijson.Field
	Dataset         apijson.Field
	DestinationConf apijson.Field
	Enabled         apijson.Field
	ErrorMessage    apijson.Field
	Frequency       apijson.Field
	LastComplete    apijson.Field
	LastError       apijson.Field
	LogpullOptions  apijson.Field
	Name            apijson.Field
	OutputOptions   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *JobListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobListResponseJSON) RawJSON() string {
	return r.raw
}

// The frequency at which Cloudflare sends batches of logs to your destination.
// Setting frequency to high sends your logs in larger quantities of smaller files.
// Setting frequency to low sends logs in smaller quantities of larger files.
type JobListResponseFrequency string

const (
	JobListResponseFrequencyHigh JobListResponseFrequency = "high"
	JobListResponseFrequencyLow  JobListResponseFrequency = "low"
)

// The structured replacement for `logpull_options`. When including this field, the
// `logpull_option` field will be ignored.
type JobListResponseOutputOptions struct {
	// String to be prepended before each batch.
	BatchPrefix string `json:"batch_prefix,nullable"`
	// String to be appended after each batch.
	BatchSuffix string `json:"batch_suffix,nullable"`
	// If set to true, will cause all occurrences of `${` in the generated files to be
	// replaced with `x{`.
	Cve2021_4428 bool `json:"CVE-2021-4428,nullable"`
	// String to join fields. This field be ignored when `record_template` is set.
	FieldDelimiter string `json:"field_delimiter,nullable"`
	// List of field names to be included in the Logpush output. For the moment, there
	// is no option to add all fields at once, so you must specify all the fields names
	// you are interested in.
	FieldNames []string `json:"field_names"`
	// Specifies the output type, such as `ndjson` or `csv`. This sets default values
	// for the rest of the settings, depending on the chosen output type. Some
	// formatting rules, like string quoting, are different between output types.
	OutputType JobListResponseOutputOptionsOutputType `json:"output_type"`
	// String to be inserted in-between the records as separator.
	RecordDelimiter string `json:"record_delimiter,nullable"`
	// String to be prepended before each record.
	RecordPrefix string `json:"record_prefix,nullable"`
	// String to be appended after each record.
	RecordSuffix string `json:"record_suffix,nullable"`
	// String to use as template for each record instead of the default comma-separated
	// list. All fields used in the template must be present in `field_names` as well,
	// otherwise they will end up as null. Format as a Go `text/template` without any
	// standard functions, like conditionals, loops, sub-templates, etc.
	RecordTemplate string `json:"record_template,nullable"`
	// Floating number to specify sampling rate. Sampling is applied on top of
	// filtering, and regardless of the current `sample_interval` of the data.
	SampleRate float64 `json:"sample_rate,nullable"`
	// String to specify the format for timestamps, such as `unixnano`, `unix`, or
	// `rfc3339`.
	TimestampFormat JobListResponseOutputOptionsTimestampFormat `json:"timestamp_format"`
	JSON            jobListResponseOutputOptionsJSON            `json:"-"`
}

// jobListResponseOutputOptionsJSON contains the JSON metadata for the struct
// [JobListResponseOutputOptions]
type jobListResponseOutputOptionsJSON struct {
	BatchPrefix     apijson.Field
	BatchSuffix     apijson.Field
	Cve2021_4428    apijson.Field
	FieldDelimiter  apijson.Field
	FieldNames      apijson.Field
	OutputType      apijson.Field
	RecordDelimiter apijson.Field
	RecordPrefix    apijson.Field
	RecordSuffix    apijson.Field
	RecordTemplate  apijson.Field
	SampleRate      apijson.Field
	TimestampFormat apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *JobListResponseOutputOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobListResponseOutputOptionsJSON) RawJSON() string {
	return r.raw
}

// Specifies the output type, such as `ndjson` or `csv`. This sets default values
// for the rest of the settings, depending on the chosen output type. Some
// formatting rules, like string quoting, are different between output types.
type JobListResponseOutputOptionsOutputType string

const (
	JobListResponseOutputOptionsOutputTypeNdjson JobListResponseOutputOptionsOutputType = "ndjson"
	JobListResponseOutputOptionsOutputTypeCsv    JobListResponseOutputOptionsOutputType = "csv"
)

// String to specify the format for timestamps, such as `unixnano`, `unix`, or
// `rfc3339`.
type JobListResponseOutputOptionsTimestampFormat string

const (
	JobListResponseOutputOptionsTimestampFormatUnixnano JobListResponseOutputOptionsTimestampFormat = "unixnano"
	JobListResponseOutputOptionsTimestampFormatUnix     JobListResponseOutputOptionsTimestampFormat = "unix"
	JobListResponseOutputOptionsTimestampFormatRfc3339  JobListResponseOutputOptionsTimestampFormat = "rfc3339"
)

// Union satisfied by [logpush.JobDeleteResponseUnknown],
// [logpush.JobDeleteResponseArray] or [shared.UnionString].
type JobDeleteResponse interface {
	ImplementsLogpushJobDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*JobDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(JobDeleteResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type JobDeleteResponseArray []interface{}

func (r JobDeleteResponseArray) ImplementsLogpushJobDeleteResponse() {}

type JobGetResponse struct {
	// Unique id of the job.
	ID int64 `json:"id"`
	// Name of the dataset.
	Dataset string `json:"dataset,nullable"`
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf string `json:"destination_conf" format:"uri"`
	// Flag that indicates if the job is enabled.
	Enabled bool `json:"enabled"`
	// If not null, the job is currently failing. Failures are usually repetitive
	// (example: no permissions to write to destination bucket). Only the last failure
	// is recorded. On successful execution of a job the error_message and last_error
	// are set to null.
	ErrorMessage time.Time `json:"error_message,nullable" format:"date-time"`
	// The frequency at which Cloudflare sends batches of logs to your destination.
	// Setting frequency to high sends your logs in larger quantities of smaller files.
	// Setting frequency to low sends logs in smaller quantities of larger files.
	Frequency JobGetResponseFrequency `json:"frequency,nullable"`
	// Records the last time for which logs have been successfully pushed. If the last
	// successful push was for logs range 2018-07-23T10:00:00Z to 2018-07-23T10:01:00Z
	// then the value of this field will be 2018-07-23T10:01:00Z. If the job has never
	// run or has just been enabled and hasn't run yet then the field will be empty.
	LastComplete time.Time `json:"last_complete,nullable" format:"date-time"`
	// Records the last time the job failed. If not null, the job is currently failing.
	// If null, the job has either never failed or has run successfully at least once
	// since last failure. See also the error_message field.
	LastError time.Time `json:"last_error,nullable" format:"date-time"`
	// This field is deprecated. Use `output_options` instead. Configuration string. It
	// specifies things like requested fields and timestamp formats. If migrating from
	// the logpull api, copy the url (full url or just the query string) of your call
	// here, and logpush will keep on making this call for you, setting start and end
	// times appropriately.
	LogpullOptions string `json:"logpull_options,nullable" format:"uri-reference"`
	// Optional human readable job name. Not unique. Cloudflare suggests that you set
	// this to a meaningful string, like the domain name, to make it easier to identify
	// your job.
	Name string `json:"name,nullable"`
	// The structured replacement for `logpull_options`. When including this field, the
	// `logpull_option` field will be ignored.
	OutputOptions JobGetResponseOutputOptions `json:"output_options,nullable"`
	JSON          jobGetResponseJSON          `json:"-"`
}

// jobGetResponseJSON contains the JSON metadata for the struct [JobGetResponse]
type jobGetResponseJSON struct {
	ID              apijson.Field
	Dataset         apijson.Field
	DestinationConf apijson.Field
	Enabled         apijson.Field
	ErrorMessage    apijson.Field
	Frequency       apijson.Field
	LastComplete    apijson.Field
	LastError       apijson.Field
	LogpullOptions  apijson.Field
	Name            apijson.Field
	OutputOptions   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *JobGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseJSON) RawJSON() string {
	return r.raw
}

// The frequency at which Cloudflare sends batches of logs to your destination.
// Setting frequency to high sends your logs in larger quantities of smaller files.
// Setting frequency to low sends logs in smaller quantities of larger files.
type JobGetResponseFrequency string

const (
	JobGetResponseFrequencyHigh JobGetResponseFrequency = "high"
	JobGetResponseFrequencyLow  JobGetResponseFrequency = "low"
)

// The structured replacement for `logpull_options`. When including this field, the
// `logpull_option` field will be ignored.
type JobGetResponseOutputOptions struct {
	// String to be prepended before each batch.
	BatchPrefix string `json:"batch_prefix,nullable"`
	// String to be appended after each batch.
	BatchSuffix string `json:"batch_suffix,nullable"`
	// If set to true, will cause all occurrences of `${` in the generated files to be
	// replaced with `x{`.
	Cve2021_4428 bool `json:"CVE-2021-4428,nullable"`
	// String to join fields. This field be ignored when `record_template` is set.
	FieldDelimiter string `json:"field_delimiter,nullable"`
	// List of field names to be included in the Logpush output. For the moment, there
	// is no option to add all fields at once, so you must specify all the fields names
	// you are interested in.
	FieldNames []string `json:"field_names"`
	// Specifies the output type, such as `ndjson` or `csv`. This sets default values
	// for the rest of the settings, depending on the chosen output type. Some
	// formatting rules, like string quoting, are different between output types.
	OutputType JobGetResponseOutputOptionsOutputType `json:"output_type"`
	// String to be inserted in-between the records as separator.
	RecordDelimiter string `json:"record_delimiter,nullable"`
	// String to be prepended before each record.
	RecordPrefix string `json:"record_prefix,nullable"`
	// String to be appended after each record.
	RecordSuffix string `json:"record_suffix,nullable"`
	// String to use as template for each record instead of the default comma-separated
	// list. All fields used in the template must be present in `field_names` as well,
	// otherwise they will end up as null. Format as a Go `text/template` without any
	// standard functions, like conditionals, loops, sub-templates, etc.
	RecordTemplate string `json:"record_template,nullable"`
	// Floating number to specify sampling rate. Sampling is applied on top of
	// filtering, and regardless of the current `sample_interval` of the data.
	SampleRate float64 `json:"sample_rate,nullable"`
	// String to specify the format for timestamps, such as `unixnano`, `unix`, or
	// `rfc3339`.
	TimestampFormat JobGetResponseOutputOptionsTimestampFormat `json:"timestamp_format"`
	JSON            jobGetResponseOutputOptionsJSON            `json:"-"`
}

// jobGetResponseOutputOptionsJSON contains the JSON metadata for the struct
// [JobGetResponseOutputOptions]
type jobGetResponseOutputOptionsJSON struct {
	BatchPrefix     apijson.Field
	BatchSuffix     apijson.Field
	Cve2021_4428    apijson.Field
	FieldDelimiter  apijson.Field
	FieldNames      apijson.Field
	OutputType      apijson.Field
	RecordDelimiter apijson.Field
	RecordPrefix    apijson.Field
	RecordSuffix    apijson.Field
	RecordTemplate  apijson.Field
	SampleRate      apijson.Field
	TimestampFormat apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *JobGetResponseOutputOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseOutputOptionsJSON) RawJSON() string {
	return r.raw
}

// Specifies the output type, such as `ndjson` or `csv`. This sets default values
// for the rest of the settings, depending on the chosen output type. Some
// formatting rules, like string quoting, are different between output types.
type JobGetResponseOutputOptionsOutputType string

const (
	JobGetResponseOutputOptionsOutputTypeNdjson JobGetResponseOutputOptionsOutputType = "ndjson"
	JobGetResponseOutputOptionsOutputTypeCsv    JobGetResponseOutputOptionsOutputType = "csv"
)

// String to specify the format for timestamps, such as `unixnano`, `unix`, or
// `rfc3339`.
type JobGetResponseOutputOptionsTimestampFormat string

const (
	JobGetResponseOutputOptionsTimestampFormatUnixnano JobGetResponseOutputOptionsTimestampFormat = "unixnano"
	JobGetResponseOutputOptionsTimestampFormatUnix     JobGetResponseOutputOptionsTimestampFormat = "unix"
	JobGetResponseOutputOptionsTimestampFormatRfc3339  JobGetResponseOutputOptionsTimestampFormat = "rfc3339"
)

type JobNewParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf,required" format:"uri"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// Name of the dataset.
	Dataset param.Field[string] `json:"dataset"`
	// Flag that indicates if the job is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// The frequency at which Cloudflare sends batches of logs to your destination.
	// Setting frequency to high sends your logs in larger quantities of smaller files.
	// Setting frequency to low sends logs in smaller quantities of larger files.
	Frequency param.Field[JobNewParamsFrequency] `json:"frequency"`
	// This field is deprecated. Use `output_options` instead. Configuration string. It
	// specifies things like requested fields and timestamp formats. If migrating from
	// the logpull api, copy the url (full url or just the query string) of your call
	// here, and logpush will keep on making this call for you, setting start and end
	// times appropriately.
	LogpullOptions param.Field[string] `json:"logpull_options" format:"uri-reference"`
	// Optional human readable job name. Not unique. Cloudflare suggests that you set
	// this to a meaningful string, like the domain name, to make it easier to identify
	// your job.
	Name param.Field[string] `json:"name"`
	// The structured replacement for `logpull_options`. When including this field, the
	// `logpull_option` field will be ignored.
	OutputOptions param.Field[JobNewParamsOutputOptions] `json:"output_options"`
	// Ownership challenge token to prove destination ownership.
	OwnershipChallenge param.Field[string] `json:"ownership_challenge"`
}

func (r JobNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The frequency at which Cloudflare sends batches of logs to your destination.
// Setting frequency to high sends your logs in larger quantities of smaller files.
// Setting frequency to low sends logs in smaller quantities of larger files.
type JobNewParamsFrequency string

const (
	JobNewParamsFrequencyHigh JobNewParamsFrequency = "high"
	JobNewParamsFrequencyLow  JobNewParamsFrequency = "low"
)

// The structured replacement for `logpull_options`. When including this field, the
// `logpull_option` field will be ignored.
type JobNewParamsOutputOptions struct {
	// String to be prepended before each batch.
	BatchPrefix param.Field[string] `json:"batch_prefix"`
	// String to be appended after each batch.
	BatchSuffix param.Field[string] `json:"batch_suffix"`
	// If set to true, will cause all occurrences of `${` in the generated files to be
	// replaced with `x{`.
	Cve2021_4428 param.Field[bool] `json:"CVE-2021-4428"`
	// String to join fields. This field be ignored when `record_template` is set.
	FieldDelimiter param.Field[string] `json:"field_delimiter"`
	// List of field names to be included in the Logpush output. For the moment, there
	// is no option to add all fields at once, so you must specify all the fields names
	// you are interested in.
	FieldNames param.Field[[]string] `json:"field_names"`
	// Specifies the output type, such as `ndjson` or `csv`. This sets default values
	// for the rest of the settings, depending on the chosen output type. Some
	// formatting rules, like string quoting, are different between output types.
	OutputType param.Field[JobNewParamsOutputOptionsOutputType] `json:"output_type"`
	// String to be inserted in-between the records as separator.
	RecordDelimiter param.Field[string] `json:"record_delimiter"`
	// String to be prepended before each record.
	RecordPrefix param.Field[string] `json:"record_prefix"`
	// String to be appended after each record.
	RecordSuffix param.Field[string] `json:"record_suffix"`
	// String to use as template for each record instead of the default comma-separated
	// list. All fields used in the template must be present in `field_names` as well,
	// otherwise they will end up as null. Format as a Go `text/template` without any
	// standard functions, like conditionals, loops, sub-templates, etc.
	RecordTemplate param.Field[string] `json:"record_template"`
	// Floating number to specify sampling rate. Sampling is applied on top of
	// filtering, and regardless of the current `sample_interval` of the data.
	SampleRate param.Field[float64] `json:"sample_rate"`
	// String to specify the format for timestamps, such as `unixnano`, `unix`, or
	// `rfc3339`.
	TimestampFormat param.Field[JobNewParamsOutputOptionsTimestampFormat] `json:"timestamp_format"`
}

func (r JobNewParamsOutputOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the output type, such as `ndjson` or `csv`. This sets default values
// for the rest of the settings, depending on the chosen output type. Some
// formatting rules, like string quoting, are different between output types.
type JobNewParamsOutputOptionsOutputType string

const (
	JobNewParamsOutputOptionsOutputTypeNdjson JobNewParamsOutputOptionsOutputType = "ndjson"
	JobNewParamsOutputOptionsOutputTypeCsv    JobNewParamsOutputOptionsOutputType = "csv"
)

// String to specify the format for timestamps, such as `unixnano`, `unix`, or
// `rfc3339`.
type JobNewParamsOutputOptionsTimestampFormat string

const (
	JobNewParamsOutputOptionsTimestampFormatUnixnano JobNewParamsOutputOptionsTimestampFormat = "unixnano"
	JobNewParamsOutputOptionsTimestampFormatUnix     JobNewParamsOutputOptionsTimestampFormat = "unix"
	JobNewParamsOutputOptionsTimestampFormatRfc3339  JobNewParamsOutputOptionsTimestampFormat = "rfc3339"
)

type JobNewResponseEnvelope struct {
	Errors   []JobNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []JobNewResponseEnvelopeMessages `json:"messages,required"`
	Result   JobNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success JobNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    jobNewResponseEnvelopeJSON    `json:"-"`
}

// jobNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [JobNewResponseEnvelope]
type jobNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type JobNewResponseEnvelopeErrors struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    jobNewResponseEnvelopeErrorsJSON `json:"-"`
}

// jobNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [JobNewResponseEnvelopeErrors]
type jobNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type JobNewResponseEnvelopeMessages struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    jobNewResponseEnvelopeMessagesJSON `json:"-"`
}

// jobNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [JobNewResponseEnvelopeMessages]
type jobNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type JobNewResponseEnvelopeSuccess bool

const (
	JobNewResponseEnvelopeSuccessTrue JobNewResponseEnvelopeSuccess = true
)

type JobUpdateParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf" format:"uri"`
	// Flag that indicates if the job is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// The frequency at which Cloudflare sends batches of logs to your destination.
	// Setting frequency to high sends your logs in larger quantities of smaller files.
	// Setting frequency to low sends logs in smaller quantities of larger files.
	Frequency param.Field[JobUpdateParamsFrequency] `json:"frequency"`
	// This field is deprecated. Use `output_options` instead. Configuration string. It
	// specifies things like requested fields and timestamp formats. If migrating from
	// the logpull api, copy the url (full url or just the query string) of your call
	// here, and logpush will keep on making this call for you, setting start and end
	// times appropriately.
	LogpullOptions param.Field[string] `json:"logpull_options" format:"uri-reference"`
	// The structured replacement for `logpull_options`. When including this field, the
	// `logpull_option` field will be ignored.
	OutputOptions param.Field[JobUpdateParamsOutputOptions] `json:"output_options"`
	// Ownership challenge token to prove destination ownership.
	OwnershipChallenge param.Field[string] `json:"ownership_challenge"`
}

func (r JobUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The frequency at which Cloudflare sends batches of logs to your destination.
// Setting frequency to high sends your logs in larger quantities of smaller files.
// Setting frequency to low sends logs in smaller quantities of larger files.
type JobUpdateParamsFrequency string

const (
	JobUpdateParamsFrequencyHigh JobUpdateParamsFrequency = "high"
	JobUpdateParamsFrequencyLow  JobUpdateParamsFrequency = "low"
)

// The structured replacement for `logpull_options`. When including this field, the
// `logpull_option` field will be ignored.
type JobUpdateParamsOutputOptions struct {
	// String to be prepended before each batch.
	BatchPrefix param.Field[string] `json:"batch_prefix"`
	// String to be appended after each batch.
	BatchSuffix param.Field[string] `json:"batch_suffix"`
	// If set to true, will cause all occurrences of `${` in the generated files to be
	// replaced with `x{`.
	Cve2021_4428 param.Field[bool] `json:"CVE-2021-4428"`
	// String to join fields. This field be ignored when `record_template` is set.
	FieldDelimiter param.Field[string] `json:"field_delimiter"`
	// List of field names to be included in the Logpush output. For the moment, there
	// is no option to add all fields at once, so you must specify all the fields names
	// you are interested in.
	FieldNames param.Field[[]string] `json:"field_names"`
	// Specifies the output type, such as `ndjson` or `csv`. This sets default values
	// for the rest of the settings, depending on the chosen output type. Some
	// formatting rules, like string quoting, are different between output types.
	OutputType param.Field[JobUpdateParamsOutputOptionsOutputType] `json:"output_type"`
	// String to be inserted in-between the records as separator.
	RecordDelimiter param.Field[string] `json:"record_delimiter"`
	// String to be prepended before each record.
	RecordPrefix param.Field[string] `json:"record_prefix"`
	// String to be appended after each record.
	RecordSuffix param.Field[string] `json:"record_suffix"`
	// String to use as template for each record instead of the default comma-separated
	// list. All fields used in the template must be present in `field_names` as well,
	// otherwise they will end up as null. Format as a Go `text/template` without any
	// standard functions, like conditionals, loops, sub-templates, etc.
	RecordTemplate param.Field[string] `json:"record_template"`
	// Floating number to specify sampling rate. Sampling is applied on top of
	// filtering, and regardless of the current `sample_interval` of the data.
	SampleRate param.Field[float64] `json:"sample_rate"`
	// String to specify the format for timestamps, such as `unixnano`, `unix`, or
	// `rfc3339`.
	TimestampFormat param.Field[JobUpdateParamsOutputOptionsTimestampFormat] `json:"timestamp_format"`
}

func (r JobUpdateParamsOutputOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the output type, such as `ndjson` or `csv`. This sets default values
// for the rest of the settings, depending on the chosen output type. Some
// formatting rules, like string quoting, are different between output types.
type JobUpdateParamsOutputOptionsOutputType string

const (
	JobUpdateParamsOutputOptionsOutputTypeNdjson JobUpdateParamsOutputOptionsOutputType = "ndjson"
	JobUpdateParamsOutputOptionsOutputTypeCsv    JobUpdateParamsOutputOptionsOutputType = "csv"
)

// String to specify the format for timestamps, such as `unixnano`, `unix`, or
// `rfc3339`.
type JobUpdateParamsOutputOptionsTimestampFormat string

const (
	JobUpdateParamsOutputOptionsTimestampFormatUnixnano JobUpdateParamsOutputOptionsTimestampFormat = "unixnano"
	JobUpdateParamsOutputOptionsTimestampFormatUnix     JobUpdateParamsOutputOptionsTimestampFormat = "unix"
	JobUpdateParamsOutputOptionsTimestampFormatRfc3339  JobUpdateParamsOutputOptionsTimestampFormat = "rfc3339"
)

type JobUpdateResponseEnvelope struct {
	Errors   []JobUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []JobUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   JobUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success JobUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    jobUpdateResponseEnvelopeJSON    `json:"-"`
}

// jobUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [JobUpdateResponseEnvelope]
type jobUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type JobUpdateResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    jobUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// jobUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [JobUpdateResponseEnvelopeErrors]
type jobUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type JobUpdateResponseEnvelopeMessages struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    jobUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// jobUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [JobUpdateResponseEnvelopeMessages]
type jobUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type JobUpdateResponseEnvelopeSuccess bool

const (
	JobUpdateResponseEnvelopeSuccessTrue JobUpdateResponseEnvelopeSuccess = true
)

type JobListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type JobListResponseEnvelope struct {
	Errors   []JobListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []JobListResponseEnvelopeMessages `json:"messages,required"`
	Result   []JobListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success JobListResponseEnvelopeSuccess `json:"success,required"`
	JSON    jobListResponseEnvelopeJSON    `json:"-"`
}

// jobListResponseEnvelopeJSON contains the JSON metadata for the struct
// [JobListResponseEnvelope]
type jobListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type JobListResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    jobListResponseEnvelopeErrorsJSON `json:"-"`
}

// jobListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [JobListResponseEnvelopeErrors]
type jobListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type JobListResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    jobListResponseEnvelopeMessagesJSON `json:"-"`
}

// jobListResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [JobListResponseEnvelopeMessages]
type jobListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type JobListResponseEnvelopeSuccess bool

const (
	JobListResponseEnvelopeSuccessTrue JobListResponseEnvelopeSuccess = true
)

type JobDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type JobDeleteResponseEnvelope struct {
	Errors   []JobDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []JobDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   JobDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success JobDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    jobDeleteResponseEnvelopeJSON    `json:"-"`
}

// jobDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [JobDeleteResponseEnvelope]
type jobDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type JobDeleteResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    jobDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// jobDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [JobDeleteResponseEnvelopeErrors]
type jobDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type JobDeleteResponseEnvelopeMessages struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    jobDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// jobDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [JobDeleteResponseEnvelopeMessages]
type jobDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type JobDeleteResponseEnvelopeSuccess bool

const (
	JobDeleteResponseEnvelopeSuccessTrue JobDeleteResponseEnvelopeSuccess = true
)

type JobGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type JobGetResponseEnvelope struct {
	Errors   []JobGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []JobGetResponseEnvelopeMessages `json:"messages,required"`
	Result   JobGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success JobGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    jobGetResponseEnvelopeJSON    `json:"-"`
}

// jobGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [JobGetResponseEnvelope]
type jobGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type JobGetResponseEnvelopeErrors struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    jobGetResponseEnvelopeErrorsJSON `json:"-"`
}

// jobGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [JobGetResponseEnvelopeErrors]
type jobGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type JobGetResponseEnvelopeMessages struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    jobGetResponseEnvelopeMessagesJSON `json:"-"`
}

// jobGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [JobGetResponseEnvelopeMessages]
type jobGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type JobGetResponseEnvelopeSuccess bool

const (
	JobGetResponseEnvelopeSuccessTrue JobGetResponseEnvelopeSuccess = true
)
