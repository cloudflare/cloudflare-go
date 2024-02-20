// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// LogpushJobService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewLogpushJobService] method instead.
type LogpushJobService struct {
	Options []option.RequestOption
}

// NewLogpushJobService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLogpushJobService(opts ...option.RequestOption) (r *LogpushJobService) {
	r = &LogpushJobService{}
	r.Options = opts
	return
}

// Creates a new Logpush job for an account or zone.
func (r *LogpushJobService) New(ctx context.Context, accountOrZone string, accountOrZoneID string, body LogpushJobNewParams, opts ...option.RequestOption) (res *LogpushJobNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LogpushJobNewResponseEnvelope
	path := fmt.Sprintf("%s/%s/logpush/jobs", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists Logpush jobs for an account or zone.
func (r *LogpushJobService) List(ctx context.Context, accountOrZone string, accountOrZoneID string, opts ...option.RequestOption) (res *[]LogpushJobListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LogpushJobListResponseEnvelope
	path := fmt.Sprintf("%s/%s/logpush/jobs", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a Logpush job.
func (r *LogpushJobService) Delete(ctx context.Context, accountOrZone string, accountOrZoneID string, jobID int64, opts ...option.RequestOption) (res *LogpushJobDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LogpushJobDeleteResponseEnvelope
	path := fmt.Sprintf("%s/%s/logpush/jobs/%v", accountOrZone, accountOrZoneID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets the details of a Logpush job.
func (r *LogpushJobService) Get(ctx context.Context, accountOrZone string, accountOrZoneID string, jobID int64, opts ...option.RequestOption) (res *LogpushJobGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LogpushJobGetResponseEnvelope
	path := fmt.Sprintf("%s/%s/logpush/jobs/%v", accountOrZone, accountOrZoneID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a Logpush job.
func (r *LogpushJobService) Replace(ctx context.Context, accountOrZone string, accountOrZoneID string, jobID int64, body LogpushJobReplaceParams, opts ...option.RequestOption) (res *LogpushJobReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LogpushJobReplaceResponseEnvelope
	path := fmt.Sprintf("%s/%s/logpush/jobs/%v", accountOrZone, accountOrZoneID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LogpushJobNewResponse struct {
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
	Frequency LogpushJobNewResponseFrequency `json:"frequency,nullable"`
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
	OutputOptions LogpushJobNewResponseOutputOptions `json:"output_options,nullable"`
	JSON          logpushJobNewResponseJSON          `json:"-"`
}

// logpushJobNewResponseJSON contains the JSON metadata for the struct
// [LogpushJobNewResponse]
type logpushJobNewResponseJSON struct {
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

func (r *LogpushJobNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The frequency at which Cloudflare sends batches of logs to your destination.
// Setting frequency to high sends your logs in larger quantities of smaller files.
// Setting frequency to low sends logs in smaller quantities of larger files.
type LogpushJobNewResponseFrequency string

const (
	LogpushJobNewResponseFrequencyHigh LogpushJobNewResponseFrequency = "high"
	LogpushJobNewResponseFrequencyLow  LogpushJobNewResponseFrequency = "low"
)

// The structured replacement for `logpull_options`. When including this field, the
// `logpull_option` field will be ignored.
type LogpushJobNewResponseOutputOptions struct {
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
	OutputType LogpushJobNewResponseOutputOptionsOutputType `json:"output_type"`
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
	TimestampFormat LogpushJobNewResponseOutputOptionsTimestampFormat `json:"timestamp_format"`
	JSON            logpushJobNewResponseOutputOptionsJSON            `json:"-"`
}

// logpushJobNewResponseOutputOptionsJSON contains the JSON metadata for the struct
// [LogpushJobNewResponseOutputOptions]
type logpushJobNewResponseOutputOptionsJSON struct {
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

func (r *LogpushJobNewResponseOutputOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the output type, such as `ndjson` or `csv`. This sets default values
// for the rest of the settings, depending on the chosen output type. Some
// formatting rules, like string quoting, are different between output types.
type LogpushJobNewResponseOutputOptionsOutputType string

const (
	LogpushJobNewResponseOutputOptionsOutputTypeNdjson LogpushJobNewResponseOutputOptionsOutputType = "ndjson"
	LogpushJobNewResponseOutputOptionsOutputTypeCsv    LogpushJobNewResponseOutputOptionsOutputType = "csv"
)

// String to specify the format for timestamps, such as `unixnano`, `unix`, or
// `rfc3339`.
type LogpushJobNewResponseOutputOptionsTimestampFormat string

const (
	LogpushJobNewResponseOutputOptionsTimestampFormatUnixnano LogpushJobNewResponseOutputOptionsTimestampFormat = "unixnano"
	LogpushJobNewResponseOutputOptionsTimestampFormatUnix     LogpushJobNewResponseOutputOptionsTimestampFormat = "unix"
	LogpushJobNewResponseOutputOptionsTimestampFormatRfc3339  LogpushJobNewResponseOutputOptionsTimestampFormat = "rfc3339"
)

type LogpushJobListResponse struct {
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
	Frequency LogpushJobListResponseFrequency `json:"frequency,nullable"`
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
	OutputOptions LogpushJobListResponseOutputOptions `json:"output_options,nullable"`
	JSON          logpushJobListResponseJSON          `json:"-"`
}

// logpushJobListResponseJSON contains the JSON metadata for the struct
// [LogpushJobListResponse]
type logpushJobListResponseJSON struct {
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

func (r *LogpushJobListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The frequency at which Cloudflare sends batches of logs to your destination.
// Setting frequency to high sends your logs in larger quantities of smaller files.
// Setting frequency to low sends logs in smaller quantities of larger files.
type LogpushJobListResponseFrequency string

const (
	LogpushJobListResponseFrequencyHigh LogpushJobListResponseFrequency = "high"
	LogpushJobListResponseFrequencyLow  LogpushJobListResponseFrequency = "low"
)

// The structured replacement for `logpull_options`. When including this field, the
// `logpull_option` field will be ignored.
type LogpushJobListResponseOutputOptions struct {
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
	OutputType LogpushJobListResponseOutputOptionsOutputType `json:"output_type"`
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
	TimestampFormat LogpushJobListResponseOutputOptionsTimestampFormat `json:"timestamp_format"`
	JSON            logpushJobListResponseOutputOptionsJSON            `json:"-"`
}

// logpushJobListResponseOutputOptionsJSON contains the JSON metadata for the
// struct [LogpushJobListResponseOutputOptions]
type logpushJobListResponseOutputOptionsJSON struct {
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

func (r *LogpushJobListResponseOutputOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the output type, such as `ndjson` or `csv`. This sets default values
// for the rest of the settings, depending on the chosen output type. Some
// formatting rules, like string quoting, are different between output types.
type LogpushJobListResponseOutputOptionsOutputType string

const (
	LogpushJobListResponseOutputOptionsOutputTypeNdjson LogpushJobListResponseOutputOptionsOutputType = "ndjson"
	LogpushJobListResponseOutputOptionsOutputTypeCsv    LogpushJobListResponseOutputOptionsOutputType = "csv"
)

// String to specify the format for timestamps, such as `unixnano`, `unix`, or
// `rfc3339`.
type LogpushJobListResponseOutputOptionsTimestampFormat string

const (
	LogpushJobListResponseOutputOptionsTimestampFormatUnixnano LogpushJobListResponseOutputOptionsTimestampFormat = "unixnano"
	LogpushJobListResponseOutputOptionsTimestampFormatUnix     LogpushJobListResponseOutputOptionsTimestampFormat = "unix"
	LogpushJobListResponseOutputOptionsTimestampFormatRfc3339  LogpushJobListResponseOutputOptionsTimestampFormat = "rfc3339"
)

// Union satisfied by [LogpushJobDeleteResponseUnknown],
// [LogpushJobDeleteResponseArray] or [shared.UnionString].
type LogpushJobDeleteResponse interface {
	ImplementsLogpushJobDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*LogpushJobDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type LogpushJobDeleteResponseArray []interface{}

func (r LogpushJobDeleteResponseArray) ImplementsLogpushJobDeleteResponse() {}

type LogpushJobGetResponse struct {
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
	Frequency LogpushJobGetResponseFrequency `json:"frequency,nullable"`
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
	OutputOptions LogpushJobGetResponseOutputOptions `json:"output_options,nullable"`
	JSON          logpushJobGetResponseJSON          `json:"-"`
}

// logpushJobGetResponseJSON contains the JSON metadata for the struct
// [LogpushJobGetResponse]
type logpushJobGetResponseJSON struct {
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

func (r *LogpushJobGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The frequency at which Cloudflare sends batches of logs to your destination.
// Setting frequency to high sends your logs in larger quantities of smaller files.
// Setting frequency to low sends logs in smaller quantities of larger files.
type LogpushJobGetResponseFrequency string

const (
	LogpushJobGetResponseFrequencyHigh LogpushJobGetResponseFrequency = "high"
	LogpushJobGetResponseFrequencyLow  LogpushJobGetResponseFrequency = "low"
)

// The structured replacement for `logpull_options`. When including this field, the
// `logpull_option` field will be ignored.
type LogpushJobGetResponseOutputOptions struct {
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
	OutputType LogpushJobGetResponseOutputOptionsOutputType `json:"output_type"`
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
	TimestampFormat LogpushJobGetResponseOutputOptionsTimestampFormat `json:"timestamp_format"`
	JSON            logpushJobGetResponseOutputOptionsJSON            `json:"-"`
}

// logpushJobGetResponseOutputOptionsJSON contains the JSON metadata for the struct
// [LogpushJobGetResponseOutputOptions]
type logpushJobGetResponseOutputOptionsJSON struct {
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

func (r *LogpushJobGetResponseOutputOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the output type, such as `ndjson` or `csv`. This sets default values
// for the rest of the settings, depending on the chosen output type. Some
// formatting rules, like string quoting, are different between output types.
type LogpushJobGetResponseOutputOptionsOutputType string

const (
	LogpushJobGetResponseOutputOptionsOutputTypeNdjson LogpushJobGetResponseOutputOptionsOutputType = "ndjson"
	LogpushJobGetResponseOutputOptionsOutputTypeCsv    LogpushJobGetResponseOutputOptionsOutputType = "csv"
)

// String to specify the format for timestamps, such as `unixnano`, `unix`, or
// `rfc3339`.
type LogpushJobGetResponseOutputOptionsTimestampFormat string

const (
	LogpushJobGetResponseOutputOptionsTimestampFormatUnixnano LogpushJobGetResponseOutputOptionsTimestampFormat = "unixnano"
	LogpushJobGetResponseOutputOptionsTimestampFormatUnix     LogpushJobGetResponseOutputOptionsTimestampFormat = "unix"
	LogpushJobGetResponseOutputOptionsTimestampFormatRfc3339  LogpushJobGetResponseOutputOptionsTimestampFormat = "rfc3339"
)

type LogpushJobReplaceResponse struct {
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
	Frequency LogpushJobReplaceResponseFrequency `json:"frequency,nullable"`
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
	OutputOptions LogpushJobReplaceResponseOutputOptions `json:"output_options,nullable"`
	JSON          logpushJobReplaceResponseJSON          `json:"-"`
}

// logpushJobReplaceResponseJSON contains the JSON metadata for the struct
// [LogpushJobReplaceResponse]
type logpushJobReplaceResponseJSON struct {
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

func (r *LogpushJobReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The frequency at which Cloudflare sends batches of logs to your destination.
// Setting frequency to high sends your logs in larger quantities of smaller files.
// Setting frequency to low sends logs in smaller quantities of larger files.
type LogpushJobReplaceResponseFrequency string

const (
	LogpushJobReplaceResponseFrequencyHigh LogpushJobReplaceResponseFrequency = "high"
	LogpushJobReplaceResponseFrequencyLow  LogpushJobReplaceResponseFrequency = "low"
)

// The structured replacement for `logpull_options`. When including this field, the
// `logpull_option` field will be ignored.
type LogpushJobReplaceResponseOutputOptions struct {
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
	OutputType LogpushJobReplaceResponseOutputOptionsOutputType `json:"output_type"`
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
	TimestampFormat LogpushJobReplaceResponseOutputOptionsTimestampFormat `json:"timestamp_format"`
	JSON            logpushJobReplaceResponseOutputOptionsJSON            `json:"-"`
}

// logpushJobReplaceResponseOutputOptionsJSON contains the JSON metadata for the
// struct [LogpushJobReplaceResponseOutputOptions]
type logpushJobReplaceResponseOutputOptionsJSON struct {
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

func (r *LogpushJobReplaceResponseOutputOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the output type, such as `ndjson` or `csv`. This sets default values
// for the rest of the settings, depending on the chosen output type. Some
// formatting rules, like string quoting, are different between output types.
type LogpushJobReplaceResponseOutputOptionsOutputType string

const (
	LogpushJobReplaceResponseOutputOptionsOutputTypeNdjson LogpushJobReplaceResponseOutputOptionsOutputType = "ndjson"
	LogpushJobReplaceResponseOutputOptionsOutputTypeCsv    LogpushJobReplaceResponseOutputOptionsOutputType = "csv"
)

// String to specify the format for timestamps, such as `unixnano`, `unix`, or
// `rfc3339`.
type LogpushJobReplaceResponseOutputOptionsTimestampFormat string

const (
	LogpushJobReplaceResponseOutputOptionsTimestampFormatUnixnano LogpushJobReplaceResponseOutputOptionsTimestampFormat = "unixnano"
	LogpushJobReplaceResponseOutputOptionsTimestampFormatUnix     LogpushJobReplaceResponseOutputOptionsTimestampFormat = "unix"
	LogpushJobReplaceResponseOutputOptionsTimestampFormatRfc3339  LogpushJobReplaceResponseOutputOptionsTimestampFormat = "rfc3339"
)

type LogpushJobNewParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf,required" format:"uri"`
	// Name of the dataset.
	Dataset param.Field[string] `json:"dataset"`
	// Flag that indicates if the job is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// The frequency at which Cloudflare sends batches of logs to your destination.
	// Setting frequency to high sends your logs in larger quantities of smaller files.
	// Setting frequency to low sends logs in smaller quantities of larger files.
	Frequency param.Field[LogpushJobNewParamsFrequency] `json:"frequency"`
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
	OutputOptions param.Field[LogpushJobNewParamsOutputOptions] `json:"output_options"`
	// Ownership challenge token to prove destination ownership.
	OwnershipChallenge param.Field[string] `json:"ownership_challenge"`
}

func (r LogpushJobNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The frequency at which Cloudflare sends batches of logs to your destination.
// Setting frequency to high sends your logs in larger quantities of smaller files.
// Setting frequency to low sends logs in smaller quantities of larger files.
type LogpushJobNewParamsFrequency string

const (
	LogpushJobNewParamsFrequencyHigh LogpushJobNewParamsFrequency = "high"
	LogpushJobNewParamsFrequencyLow  LogpushJobNewParamsFrequency = "low"
)

// The structured replacement for `logpull_options`. When including this field, the
// `logpull_option` field will be ignored.
type LogpushJobNewParamsOutputOptions struct {
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
	OutputType param.Field[LogpushJobNewParamsOutputOptionsOutputType] `json:"output_type"`
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
	TimestampFormat param.Field[LogpushJobNewParamsOutputOptionsTimestampFormat] `json:"timestamp_format"`
}

func (r LogpushJobNewParamsOutputOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the output type, such as `ndjson` or `csv`. This sets default values
// for the rest of the settings, depending on the chosen output type. Some
// formatting rules, like string quoting, are different between output types.
type LogpushJobNewParamsOutputOptionsOutputType string

const (
	LogpushJobNewParamsOutputOptionsOutputTypeNdjson LogpushJobNewParamsOutputOptionsOutputType = "ndjson"
	LogpushJobNewParamsOutputOptionsOutputTypeCsv    LogpushJobNewParamsOutputOptionsOutputType = "csv"
)

// String to specify the format for timestamps, such as `unixnano`, `unix`, or
// `rfc3339`.
type LogpushJobNewParamsOutputOptionsTimestampFormat string

const (
	LogpushJobNewParamsOutputOptionsTimestampFormatUnixnano LogpushJobNewParamsOutputOptionsTimestampFormat = "unixnano"
	LogpushJobNewParamsOutputOptionsTimestampFormatUnix     LogpushJobNewParamsOutputOptionsTimestampFormat = "unix"
	LogpushJobNewParamsOutputOptionsTimestampFormatRfc3339  LogpushJobNewParamsOutputOptionsTimestampFormat = "rfc3339"
)

type LogpushJobNewResponseEnvelope struct {
	Errors   []LogpushJobNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LogpushJobNewResponseEnvelopeMessages `json:"messages,required"`
	Result   LogpushJobNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success LogpushJobNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    logpushJobNewResponseEnvelopeJSON    `json:"-"`
}

// logpushJobNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [LogpushJobNewResponseEnvelope]
type logpushJobNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushJobNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushJobNewResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    logpushJobNewResponseEnvelopeErrorsJSON `json:"-"`
}

// logpushJobNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [LogpushJobNewResponseEnvelopeErrors]
type logpushJobNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushJobNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushJobNewResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    logpushJobNewResponseEnvelopeMessagesJSON `json:"-"`
}

// logpushJobNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [LogpushJobNewResponseEnvelopeMessages]
type logpushJobNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushJobNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LogpushJobNewResponseEnvelopeSuccess bool

const (
	LogpushJobNewResponseEnvelopeSuccessTrue LogpushJobNewResponseEnvelopeSuccess = true
)

type LogpushJobListResponseEnvelope struct {
	Errors   []LogpushJobListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LogpushJobListResponseEnvelopeMessages `json:"messages,required"`
	Result   []LogpushJobListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success LogpushJobListResponseEnvelopeSuccess `json:"success,required"`
	JSON    logpushJobListResponseEnvelopeJSON    `json:"-"`
}

// logpushJobListResponseEnvelopeJSON contains the JSON metadata for the struct
// [LogpushJobListResponseEnvelope]
type logpushJobListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushJobListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushJobListResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    logpushJobListResponseEnvelopeErrorsJSON `json:"-"`
}

// logpushJobListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [LogpushJobListResponseEnvelopeErrors]
type logpushJobListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushJobListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushJobListResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    logpushJobListResponseEnvelopeMessagesJSON `json:"-"`
}

// logpushJobListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [LogpushJobListResponseEnvelopeMessages]
type logpushJobListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushJobListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LogpushJobListResponseEnvelopeSuccess bool

const (
	LogpushJobListResponseEnvelopeSuccessTrue LogpushJobListResponseEnvelopeSuccess = true
)

type LogpushJobDeleteResponseEnvelope struct {
	Errors   []LogpushJobDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LogpushJobDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   LogpushJobDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success LogpushJobDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    logpushJobDeleteResponseEnvelopeJSON    `json:"-"`
}

// logpushJobDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [LogpushJobDeleteResponseEnvelope]
type logpushJobDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushJobDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushJobDeleteResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    logpushJobDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// logpushJobDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [LogpushJobDeleteResponseEnvelopeErrors]
type logpushJobDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushJobDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushJobDeleteResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    logpushJobDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// logpushJobDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [LogpushJobDeleteResponseEnvelopeMessages]
type logpushJobDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushJobDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LogpushJobDeleteResponseEnvelopeSuccess bool

const (
	LogpushJobDeleteResponseEnvelopeSuccessTrue LogpushJobDeleteResponseEnvelopeSuccess = true
)

type LogpushJobGetResponseEnvelope struct {
	Errors   []LogpushJobGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LogpushJobGetResponseEnvelopeMessages `json:"messages,required"`
	Result   LogpushJobGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success LogpushJobGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    logpushJobGetResponseEnvelopeJSON    `json:"-"`
}

// logpushJobGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [LogpushJobGetResponseEnvelope]
type logpushJobGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushJobGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushJobGetResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    logpushJobGetResponseEnvelopeErrorsJSON `json:"-"`
}

// logpushJobGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [LogpushJobGetResponseEnvelopeErrors]
type logpushJobGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushJobGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushJobGetResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    logpushJobGetResponseEnvelopeMessagesJSON `json:"-"`
}

// logpushJobGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [LogpushJobGetResponseEnvelopeMessages]
type logpushJobGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushJobGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LogpushJobGetResponseEnvelopeSuccess bool

const (
	LogpushJobGetResponseEnvelopeSuccessTrue LogpushJobGetResponseEnvelopeSuccess = true
)

type LogpushJobReplaceParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf" format:"uri"`
	// Flag that indicates if the job is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// The frequency at which Cloudflare sends batches of logs to your destination.
	// Setting frequency to high sends your logs in larger quantities of smaller files.
	// Setting frequency to low sends logs in smaller quantities of larger files.
	Frequency param.Field[LogpushJobReplaceParamsFrequency] `json:"frequency"`
	// This field is deprecated. Use `output_options` instead. Configuration string. It
	// specifies things like requested fields and timestamp formats. If migrating from
	// the logpull api, copy the url (full url or just the query string) of your call
	// here, and logpush will keep on making this call for you, setting start and end
	// times appropriately.
	LogpullOptions param.Field[string] `json:"logpull_options" format:"uri-reference"`
	// The structured replacement for `logpull_options`. When including this field, the
	// `logpull_option` field will be ignored.
	OutputOptions param.Field[LogpushJobReplaceParamsOutputOptions] `json:"output_options"`
	// Ownership challenge token to prove destination ownership.
	OwnershipChallenge param.Field[string] `json:"ownership_challenge"`
}

func (r LogpushJobReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The frequency at which Cloudflare sends batches of logs to your destination.
// Setting frequency to high sends your logs in larger quantities of smaller files.
// Setting frequency to low sends logs in smaller quantities of larger files.
type LogpushJobReplaceParamsFrequency string

const (
	LogpushJobReplaceParamsFrequencyHigh LogpushJobReplaceParamsFrequency = "high"
	LogpushJobReplaceParamsFrequencyLow  LogpushJobReplaceParamsFrequency = "low"
)

// The structured replacement for `logpull_options`. When including this field, the
// `logpull_option` field will be ignored.
type LogpushJobReplaceParamsOutputOptions struct {
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
	OutputType param.Field[LogpushJobReplaceParamsOutputOptionsOutputType] `json:"output_type"`
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
	TimestampFormat param.Field[LogpushJobReplaceParamsOutputOptionsTimestampFormat] `json:"timestamp_format"`
}

func (r LogpushJobReplaceParamsOutputOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the output type, such as `ndjson` or `csv`. This sets default values
// for the rest of the settings, depending on the chosen output type. Some
// formatting rules, like string quoting, are different between output types.
type LogpushJobReplaceParamsOutputOptionsOutputType string

const (
	LogpushJobReplaceParamsOutputOptionsOutputTypeNdjson LogpushJobReplaceParamsOutputOptionsOutputType = "ndjson"
	LogpushJobReplaceParamsOutputOptionsOutputTypeCsv    LogpushJobReplaceParamsOutputOptionsOutputType = "csv"
)

// String to specify the format for timestamps, such as `unixnano`, `unix`, or
// `rfc3339`.
type LogpushJobReplaceParamsOutputOptionsTimestampFormat string

const (
	LogpushJobReplaceParamsOutputOptionsTimestampFormatUnixnano LogpushJobReplaceParamsOutputOptionsTimestampFormat = "unixnano"
	LogpushJobReplaceParamsOutputOptionsTimestampFormatUnix     LogpushJobReplaceParamsOutputOptionsTimestampFormat = "unix"
	LogpushJobReplaceParamsOutputOptionsTimestampFormatRfc3339  LogpushJobReplaceParamsOutputOptionsTimestampFormat = "rfc3339"
)

type LogpushJobReplaceResponseEnvelope struct {
	Errors   []LogpushJobReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LogpushJobReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   LogpushJobReplaceResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success LogpushJobReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    logpushJobReplaceResponseEnvelopeJSON    `json:"-"`
}

// logpushJobReplaceResponseEnvelopeJSON contains the JSON metadata for the struct
// [LogpushJobReplaceResponseEnvelope]
type logpushJobReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushJobReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushJobReplaceResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    logpushJobReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// logpushJobReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [LogpushJobReplaceResponseEnvelopeErrors]
type logpushJobReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushJobReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushJobReplaceResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    logpushJobReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// logpushJobReplaceResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [LogpushJobReplaceResponseEnvelopeMessages]
type logpushJobReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushJobReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LogpushJobReplaceResponseEnvelopeSuccess bool

const (
	LogpushJobReplaceResponseEnvelopeSuccessTrue LogpushJobReplaceResponseEnvelopeSuccess = true
)
