// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

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
func (r *LogpushJobService) New(ctx context.Context, params LogpushJobNewParams, opts ...option.RequestOption) (res *LogpushLogpushJob, err error) {
	opts = append(r.Options[:], opts...)
	var env LogpushJobNewResponseEnvelope
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
func (r *LogpushJobService) Update(ctx context.Context, jobID int64, params LogpushJobUpdateParams, opts ...option.RequestOption) (res *LogpushLogpushJob, err error) {
	opts = append(r.Options[:], opts...)
	var env LogpushJobUpdateResponseEnvelope
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
func (r *LogpushJobService) List(ctx context.Context, query LogpushJobListParams, opts ...option.RequestOption) (res *[]LogpushLogpushJob, err error) {
	opts = append(r.Options[:], opts...)
	var env LogpushJobListResponseEnvelope
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
func (r *LogpushJobService) Delete(ctx context.Context, jobID int64, body LogpushJobDeleteParams, opts ...option.RequestOption) (res *LogpushJobDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LogpushJobDeleteResponseEnvelope
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
func (r *LogpushJobService) Get(ctx context.Context, jobID int64, query LogpushJobGetParams, opts ...option.RequestOption) (res *LogpushLogpushJob, err error) {
	opts = append(r.Options[:], opts...)
	var env LogpushJobGetResponseEnvelope
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

type LogpushJobNewParams struct {
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
	Result   LogpushLogpushJob                       `json:"result,required,nullable"`
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

type LogpushJobUpdateParams struct {
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
	Errors   []LogpushJobUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LogpushJobUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   LogpushLogpushJob                          `json:"result,required,nullable"`
	// Whether the API call was successful
	Success LogpushJobUpdateResponseEnvelopeSuccess `json:"success,required"`
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

type LogpushJobListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type LogpushJobListResponseEnvelope struct {
	Errors   []LogpushJobListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LogpushJobListResponseEnvelopeMessages `json:"messages,required"`
	Result   []LogpushLogpushJob                      `json:"result,required"`
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

type LogpushJobDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

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

type LogpushJobGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type LogpushJobGetResponseEnvelope struct {
	Errors   []LogpushJobGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LogpushJobGetResponseEnvelopeMessages `json:"messages,required"`
	Result   LogpushLogpushJob                       `json:"result,required,nullable"`
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
