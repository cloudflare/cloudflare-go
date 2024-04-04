// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
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
func (r *JobService) New(ctx context.Context, params JobNewParams, opts ...option.RequestOption) (res *LogpushJob, err error) {
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
func (r *JobService) Update(ctx context.Context, jobID int64, params JobUpdateParams, opts ...option.RequestOption) (res *LogpushJob, err error) {
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
func (r *JobService) List(ctx context.Context, query JobListParams, opts ...option.RequestOption) (res *pagination.SinglePage[LogpushJob], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// Lists Logpush jobs for an account or zone.
func (r *JobService) ListAutoPaging(ctx context.Context, query JobListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[LogpushJob] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a Logpush job.
func (r *JobService) Delete(ctx context.Context, jobID int64, params JobDeleteParams, opts ...option.RequestOption) (res *shared.UnnamedSchemaRef161, err error) {
	opts = append(r.Options[:], opts...)
	var env JobDeleteResponseEnvelope
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets the details of a Logpush job.
func (r *JobService) Get(ctx context.Context, jobID int64, query JobGetParams, opts ...option.RequestOption) (res *LogpushJob, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

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

func (r JobNewParamsFrequency) IsKnown() bool {
	switch r {
	case JobNewParamsFrequencyHigh, JobNewParamsFrequencyLow:
		return true
	}
	return false
}

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

func (r JobNewParamsOutputOptionsOutputType) IsKnown() bool {
	switch r {
	case JobNewParamsOutputOptionsOutputTypeNdjson, JobNewParamsOutputOptionsOutputTypeCsv:
		return true
	}
	return false
}

// String to specify the format for timestamps, such as `unixnano`, `unix`, or
// `rfc3339`.
type JobNewParamsOutputOptionsTimestampFormat string

const (
	JobNewParamsOutputOptionsTimestampFormatUnixnano JobNewParamsOutputOptionsTimestampFormat = "unixnano"
	JobNewParamsOutputOptionsTimestampFormatUnix     JobNewParamsOutputOptionsTimestampFormat = "unix"
	JobNewParamsOutputOptionsTimestampFormatRfc3339  JobNewParamsOutputOptionsTimestampFormat = "rfc3339"
)

func (r JobNewParamsOutputOptionsTimestampFormat) IsKnown() bool {
	switch r {
	case JobNewParamsOutputOptionsTimestampFormatUnixnano, JobNewParamsOutputOptionsTimestampFormatUnix, JobNewParamsOutputOptionsTimestampFormatRfc3339:
		return true
	}
	return false
}

type JobNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   LogpushJob            `json:"result,required,nullable"`
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

// Whether the API call was successful
type JobNewResponseEnvelopeSuccess bool

const (
	JobNewResponseEnvelopeSuccessTrue JobNewResponseEnvelopeSuccess = true
)

func (r JobNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case JobNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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

func (r JobUpdateParamsFrequency) IsKnown() bool {
	switch r {
	case JobUpdateParamsFrequencyHigh, JobUpdateParamsFrequencyLow:
		return true
	}
	return false
}

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

func (r JobUpdateParamsOutputOptionsOutputType) IsKnown() bool {
	switch r {
	case JobUpdateParamsOutputOptionsOutputTypeNdjson, JobUpdateParamsOutputOptionsOutputTypeCsv:
		return true
	}
	return false
}

// String to specify the format for timestamps, such as `unixnano`, `unix`, or
// `rfc3339`.
type JobUpdateParamsOutputOptionsTimestampFormat string

const (
	JobUpdateParamsOutputOptionsTimestampFormatUnixnano JobUpdateParamsOutputOptionsTimestampFormat = "unixnano"
	JobUpdateParamsOutputOptionsTimestampFormatUnix     JobUpdateParamsOutputOptionsTimestampFormat = "unix"
	JobUpdateParamsOutputOptionsTimestampFormatRfc3339  JobUpdateParamsOutputOptionsTimestampFormat = "rfc3339"
)

func (r JobUpdateParamsOutputOptionsTimestampFormat) IsKnown() bool {
	switch r {
	case JobUpdateParamsOutputOptionsTimestampFormatUnixnano, JobUpdateParamsOutputOptionsTimestampFormatUnix, JobUpdateParamsOutputOptionsTimestampFormatRfc3339:
		return true
	}
	return false
}

type JobUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   LogpushJob            `json:"result,required,nullable"`
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

// Whether the API call was successful
type JobUpdateResponseEnvelopeSuccess bool

const (
	JobUpdateResponseEnvelopeSuccessTrue JobUpdateResponseEnvelopeSuccess = true
)

func (r JobUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case JobUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type JobListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type JobDeleteParams struct {
	Body param.Field[interface{}] `json:"body,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r JobDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type JobDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo      `json:"errors,required"`
	Messages []shared.ResponseInfo      `json:"messages,required"`
	Result   shared.UnnamedSchemaRef161 `json:"result,required,nullable"`
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

// Whether the API call was successful
type JobDeleteResponseEnvelopeSuccess bool

const (
	JobDeleteResponseEnvelopeSuccessTrue JobDeleteResponseEnvelopeSuccess = true
)

func (r JobDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case JobDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type JobGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type JobGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   LogpushJob            `json:"result,required,nullable"`
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

// Whether the API call was successful
type JobGetResponseEnvelopeSuccess bool

const (
	JobGetResponseEnvelopeSuccessTrue JobGetResponseEnvelopeSuccess = true
)

func (r JobGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case JobGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
