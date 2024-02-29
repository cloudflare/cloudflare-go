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

// LogpushDatasetJobService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLogpushDatasetJobService] method
// instead.
type LogpushDatasetJobService struct {
	Options []option.RequestOption
}

// NewLogpushDatasetJobService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLogpushDatasetJobService(opts ...option.RequestOption) (r *LogpushDatasetJobService) {
	r = &LogpushDatasetJobService{}
	r.Options = opts
	return
}

// Lists Logpush jobs for an account or zone for a dataset.
func (r *LogpushDatasetJobService) List(ctx context.Context, datasetID string, query LogpushDatasetJobListParams, opts ...option.RequestOption) (res *[]LogpushDatasetJobListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LogpushDatasetJobListResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/logpush/datasets/%s/jobs", accountOrZone, accountOrZoneID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LogpushDatasetJobListResponse struct {
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
	Frequency LogpushDatasetJobListResponseFrequency `json:"frequency,nullable"`
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
	OutputOptions LogpushDatasetJobListResponseOutputOptions `json:"output_options,nullable"`
	JSON          logpushDatasetJobListResponseJSON          `json:"-"`
}

// logpushDatasetJobListResponseJSON contains the JSON metadata for the struct
// [LogpushDatasetJobListResponse]
type logpushDatasetJobListResponseJSON struct {
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

func (r *LogpushDatasetJobListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The frequency at which Cloudflare sends batches of logs to your destination.
// Setting frequency to high sends your logs in larger quantities of smaller files.
// Setting frequency to low sends logs in smaller quantities of larger files.
type LogpushDatasetJobListResponseFrequency string

const (
	LogpushDatasetJobListResponseFrequencyHigh LogpushDatasetJobListResponseFrequency = "high"
	LogpushDatasetJobListResponseFrequencyLow  LogpushDatasetJobListResponseFrequency = "low"
)

// The structured replacement for `logpull_options`. When including this field, the
// `logpull_option` field will be ignored.
type LogpushDatasetJobListResponseOutputOptions struct {
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
	OutputType LogpushDatasetJobListResponseOutputOptionsOutputType `json:"output_type"`
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
	TimestampFormat LogpushDatasetJobListResponseOutputOptionsTimestampFormat `json:"timestamp_format"`
	JSON            logpushDatasetJobListResponseOutputOptionsJSON            `json:"-"`
}

// logpushDatasetJobListResponseOutputOptionsJSON contains the JSON metadata for
// the struct [LogpushDatasetJobListResponseOutputOptions]
type logpushDatasetJobListResponseOutputOptionsJSON struct {
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

func (r *LogpushDatasetJobListResponseOutputOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the output type, such as `ndjson` or `csv`. This sets default values
// for the rest of the settings, depending on the chosen output type. Some
// formatting rules, like string quoting, are different between output types.
type LogpushDatasetJobListResponseOutputOptionsOutputType string

const (
	LogpushDatasetJobListResponseOutputOptionsOutputTypeNdjson LogpushDatasetJobListResponseOutputOptionsOutputType = "ndjson"
	LogpushDatasetJobListResponseOutputOptionsOutputTypeCsv    LogpushDatasetJobListResponseOutputOptionsOutputType = "csv"
)

// String to specify the format for timestamps, such as `unixnano`, `unix`, or
// `rfc3339`.
type LogpushDatasetJobListResponseOutputOptionsTimestampFormat string

const (
	LogpushDatasetJobListResponseOutputOptionsTimestampFormatUnixnano LogpushDatasetJobListResponseOutputOptionsTimestampFormat = "unixnano"
	LogpushDatasetJobListResponseOutputOptionsTimestampFormatUnix     LogpushDatasetJobListResponseOutputOptionsTimestampFormat = "unix"
	LogpushDatasetJobListResponseOutputOptionsTimestampFormatRfc3339  LogpushDatasetJobListResponseOutputOptionsTimestampFormat = "rfc3339"
)

type LogpushDatasetJobListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type LogpushDatasetJobListResponseEnvelope struct {
	Errors   []LogpushDatasetJobListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LogpushDatasetJobListResponseEnvelopeMessages `json:"messages,required"`
	Result   []LogpushDatasetJobListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success LogpushDatasetJobListResponseEnvelopeSuccess `json:"success,required"`
	JSON    logpushDatasetJobListResponseEnvelopeJSON    `json:"-"`
}

// logpushDatasetJobListResponseEnvelopeJSON contains the JSON metadata for the
// struct [LogpushDatasetJobListResponseEnvelope]
type logpushDatasetJobListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushDatasetJobListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushDatasetJobListResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    logpushDatasetJobListResponseEnvelopeErrorsJSON `json:"-"`
}

// logpushDatasetJobListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [LogpushDatasetJobListResponseEnvelopeErrors]
type logpushDatasetJobListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushDatasetJobListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushDatasetJobListResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    logpushDatasetJobListResponseEnvelopeMessagesJSON `json:"-"`
}

// logpushDatasetJobListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [LogpushDatasetJobListResponseEnvelopeMessages]
type logpushDatasetJobListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushDatasetJobListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LogpushDatasetJobListResponseEnvelopeSuccess bool

const (
	LogpushDatasetJobListResponseEnvelopeSuccessTrue LogpushDatasetJobListResponseEnvelopeSuccess = true
)
