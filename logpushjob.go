// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
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
func (r *LogpushJobService) Update(ctx context.Context, accountOrZone string, accountOrZoneID string, jobID int64, body LogpushJobUpdateParams, opts ...option.RequestOption) (res *LogpushJobUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LogpushJobUpdateResponseEnvelope
	path := fmt.Sprintf("%s/%s/logpush/jobs/%v", accountOrZone, accountOrZoneID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
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

type LogpushJobUpdateResponse struct {
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
	Frequency LogpushJobUpdateResponseFrequency `json:"frequency,nullable"`
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
	OutputOptions LogpushJobUpdateResponseOutputOptions `json:"output_options,nullable"`
	JSON          logpushJobUpdateResponseJSON          `json:"-"`
}

// logpushJobUpdateResponseJSON contains the JSON metadata for the struct
// [LogpushJobUpdateResponse]
type logpushJobUpdateResponseJSON struct {
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

func (r *LogpushJobUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The frequency at which Cloudflare sends batches of logs to your destination.
// Setting frequency to high sends your logs in larger quantities of smaller files.
// Setting frequency to low sends logs in smaller quantities of larger files.
type LogpushJobUpdateResponseFrequency string

const (
	LogpushJobUpdateResponseFrequencyHigh LogpushJobUpdateResponseFrequency = "high"
	LogpushJobUpdateResponseFrequencyLow  LogpushJobUpdateResponseFrequency = "low"
)

// The structured replacement for `logpull_options`. When including this field, the
// `logpull_option` field will be ignored.
type LogpushJobUpdateResponseOutputOptions struct {
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
	OutputType LogpushJobUpdateResponseOutputOptionsOutputType `json:"output_type"`
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
	TimestampFormat LogpushJobUpdateResponseOutputOptionsTimestampFormat `json:"timestamp_format"`
	JSON            logpushJobUpdateResponseOutputOptionsJSON            `json:"-"`
}

// logpushJobUpdateResponseOutputOptionsJSON contains the JSON metadata for the
// struct [LogpushJobUpdateResponseOutputOptions]
type logpushJobUpdateResponseOutputOptionsJSON struct {
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

func (r *LogpushJobUpdateResponseOutputOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the output type, such as `ndjson` or `csv`. This sets default values
// for the rest of the settings, depending on the chosen output type. Some
// formatting rules, like string quoting, are different between output types.
type LogpushJobUpdateResponseOutputOptionsOutputType string

const (
	LogpushJobUpdateResponseOutputOptionsOutputTypeNdjson LogpushJobUpdateResponseOutputOptionsOutputType = "ndjson"
	LogpushJobUpdateResponseOutputOptionsOutputTypeCsv    LogpushJobUpdateResponseOutputOptionsOutputType = "csv"
)

// String to specify the format for timestamps, such as `unixnano`, `unix`, or
// `rfc3339`.
type LogpushJobUpdateResponseOutputOptionsTimestampFormat string

const (
	LogpushJobUpdateResponseOutputOptionsTimestampFormatUnixnano LogpushJobUpdateResponseOutputOptionsTimestampFormat = "unixnano"
	LogpushJobUpdateResponseOutputOptionsTimestampFormatUnix     LogpushJobUpdateResponseOutputOptionsTimestampFormat = "unix"
	LogpushJobUpdateResponseOutputOptionsTimestampFormatRfc3339  LogpushJobUpdateResponseOutputOptionsTimestampFormat = "rfc3339"
)

type LogpushJobDeleteResponse = interface{}

type LogpushJobGetResponseEnvelope struct {
	Errors   []LogpushJobGetResponseEnvelopeErrors   `json:"errors"`
	Messages []LogpushJobGetResponseEnvelopeMessages `json:"messages"`
	Result   LogpushJobGetResponse                   `json:"result,nullable"`
	// Whether the API call was successful
	Success LogpushJobGetResponseEnvelopeSuccess `json:"success"`
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

type LogpushJobUpdateParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf" format:"uri"`
	// Flag that indicates if the job is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// The frequency at which Cloudflare sends batches of logs to your destination.
	// Setting frequency to high sends your logs in larger quantities of smaller files.
	// Setting frequency to low sends logs in smaller quantities of larger files.
	Frequency param.Field[LogpushJobUpdateParamsFrequency] `json:"frequency"`
	// This field is deprecated. Use `output_options` instead. Configuration string. It
	// specifies things like requested fields and timestamp formats. If migrating from
	// the logpull api, copy the url (full url or just the query string) of your call
	// here, and logpush will keep on making this call for you, setting start and end
	// times appropriately.
	LogpullOptions param.Field[string] `json:"logpull_options" format:"uri-reference"`
	// The structured replacement for `logpull_options`. When including this field, the
	// `logpull_option` field will be ignored.
	OutputOptions param.Field[LogpushJobUpdateParamsOutputOptions] `json:"output_options"`
	// Ownership challenge token to prove destination ownership.
	OwnershipChallenge param.Field[string] `json:"ownership_challenge"`
}

func (r LogpushJobUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The frequency at which Cloudflare sends batches of logs to your destination.
// Setting frequency to high sends your logs in larger quantities of smaller files.
// Setting frequency to low sends logs in smaller quantities of larger files.
type LogpushJobUpdateParamsFrequency string

const (
	LogpushJobUpdateParamsFrequencyHigh LogpushJobUpdateParamsFrequency = "high"
	LogpushJobUpdateParamsFrequencyLow  LogpushJobUpdateParamsFrequency = "low"
)

// The structured replacement for `logpull_options`. When including this field, the
// `logpull_option` field will be ignored.
type LogpushJobUpdateParamsOutputOptions struct {
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
	OutputType param.Field[LogpushJobUpdateParamsOutputOptionsOutputType] `json:"output_type"`
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
	TimestampFormat param.Field[LogpushJobUpdateParamsOutputOptionsTimestampFormat] `json:"timestamp_format"`
}

func (r LogpushJobUpdateParamsOutputOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the output type, such as `ndjson` or `csv`. This sets default values
// for the rest of the settings, depending on the chosen output type. Some
// formatting rules, like string quoting, are different between output types.
type LogpushJobUpdateParamsOutputOptionsOutputType string

const (
	LogpushJobUpdateParamsOutputOptionsOutputTypeNdjson LogpushJobUpdateParamsOutputOptionsOutputType = "ndjson"
	LogpushJobUpdateParamsOutputOptionsOutputTypeCsv    LogpushJobUpdateParamsOutputOptionsOutputType = "csv"
)

// String to specify the format for timestamps, such as `unixnano`, `unix`, or
// `rfc3339`.
type LogpushJobUpdateParamsOutputOptionsTimestampFormat string

const (
	LogpushJobUpdateParamsOutputOptionsTimestampFormatUnixnano LogpushJobUpdateParamsOutputOptionsTimestampFormat = "unixnano"
	LogpushJobUpdateParamsOutputOptionsTimestampFormatUnix     LogpushJobUpdateParamsOutputOptionsTimestampFormat = "unix"
	LogpushJobUpdateParamsOutputOptionsTimestampFormatRfc3339  LogpushJobUpdateParamsOutputOptionsTimestampFormat = "rfc3339"
)

type LogpushJobUpdateResponseEnvelope struct {
	Errors   []LogpushJobUpdateResponseEnvelopeErrors   `json:"errors"`
	Messages []LogpushJobUpdateResponseEnvelopeMessages `json:"messages"`
	Result   LogpushJobUpdateResponse                   `json:"result,nullable"`
	// Whether the API call was successful
	Success LogpushJobUpdateResponseEnvelopeSuccess `json:"success"`
	JSON    logpushJobUpdateResponseEnvelopeJSON    `json:"-"`
}

// logpushJobUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [LogpushJobUpdateResponseEnvelope]
type logpushJobUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushJobUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushJobUpdateResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    logpushJobUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// logpushJobUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [LogpushJobUpdateResponseEnvelopeErrors]
type logpushJobUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushJobUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushJobUpdateResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    logpushJobUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// logpushJobUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [LogpushJobUpdateResponseEnvelopeMessages]
type logpushJobUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushJobUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LogpushJobUpdateResponseEnvelopeSuccess bool

const (
	LogpushJobUpdateResponseEnvelopeSuccessTrue LogpushJobUpdateResponseEnvelopeSuccess = true
)

type LogpushJobDeleteResponseEnvelope struct {
	Errors   []LogpushJobDeleteResponseEnvelopeErrors   `json:"errors"`
	Messages []LogpushJobDeleteResponseEnvelopeMessages `json:"messages"`
	Result   LogpushJobDeleteResponse                   `json:"result,nullable"`
	// Whether the API call was successful
	Success LogpushJobDeleteResponseEnvelopeSuccess `json:"success"`
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
