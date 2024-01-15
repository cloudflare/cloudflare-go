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

// ZoneLogpushJobService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneLogpushJobService] method
// instead.
type ZoneLogpushJobService struct {
	Options []option.RequestOption
}

// NewZoneLogpushJobService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneLogpushJobService(opts ...option.RequestOption) (r *ZoneLogpushJobService) {
	r = &ZoneLogpushJobService{}
	r.Options = opts
	return
}

// Gets the details of a Logpush job.
func (r *ZoneLogpushJobService) Get(ctx context.Context, zoneIdentifier string, jobIdentifier int64, opts ...option.RequestOption) (res *ZoneLogpushJobGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/logpush/jobs/%v", zoneIdentifier, jobIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a Logpush job.
func (r *ZoneLogpushJobService) Update(ctx context.Context, zoneIdentifier string, jobIdentifier int64, body ZoneLogpushJobUpdateParams, opts ...option.RequestOption) (res *ZoneLogpushJobUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/logpush/jobs/%v", zoneIdentifier, jobIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes a Logpush job.
func (r *ZoneLogpushJobService) Delete(ctx context.Context, zoneIdentifier string, jobIdentifier int64, opts ...option.RequestOption) (res *ZoneLogpushJobDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/logpush/jobs/%v", zoneIdentifier, jobIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Lists Logpush jobs for a zone.
func (r *ZoneLogpushJobService) GetZonesZoneIdentifierLogpushJobs(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/logpush/jobs", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Creates a new Logpush job for a zone.
func (r *ZoneLogpushJobService) PostZonesZoneIdentifierLogpushJobs(ctx context.Context, zoneIdentifier string, body ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsParams, opts ...option.RequestOption) (res *ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/logpush/jobs", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ZoneLogpushJobGetResponse struct {
	Errors   []ZoneLogpushJobGetResponseError   `json:"errors"`
	Messages []ZoneLogpushJobGetResponseMessage `json:"messages"`
	Result   ZoneLogpushJobGetResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success ZoneLogpushJobGetResponseSuccess `json:"success"`
	JSON    zoneLogpushJobGetResponseJSON    `json:"-"`
}

// zoneLogpushJobGetResponseJSON contains the JSON metadata for the struct
// [ZoneLogpushJobGetResponse]
type zoneLogpushJobGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushJobGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogpushJobGetResponseError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    zoneLogpushJobGetResponseErrorJSON `json:"-"`
}

// zoneLogpushJobGetResponseErrorJSON contains the JSON metadata for the struct
// [ZoneLogpushJobGetResponseError]
type zoneLogpushJobGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushJobGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogpushJobGetResponseMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    zoneLogpushJobGetResponseMessageJSON `json:"-"`
}

// zoneLogpushJobGetResponseMessageJSON contains the JSON metadata for the struct
// [ZoneLogpushJobGetResponseMessage]
type zoneLogpushJobGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushJobGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogpushJobGetResponseResult struct {
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
	Frequency ZoneLogpushJobGetResponseResultFrequency `json:"frequency,nullable"`
	// Records the last time for which logs have been successfully pushed. If the last
	// successful push was for logs range 2018-07-23T10:00:00Z to 2018-07-23T10:01:00Z
	// then the value of this field will be 2018-07-23T10:01:00Z. If the job has never
	// run or has just been enabled and hasn't run yet then the field will be empty.
	LastComplete time.Time `json:"last_complete,nullable" format:"date-time"`
	// Records the last time the job failed. If not null, the job is currently failing.
	// If null, the job has either never failed or has run successfully at least once
	// since last failure. See also the error_message field.
	LastError time.Time `json:"last_error,nullable" format:"date-time"`
	// Configuration string. It specifies things like requested fields and timestamp
	// formats. If migrating from the logpull api, copy the url (full url or just the
	// query string) of your call here, and logpush will keep on making this call for
	// you, setting start and end times appropriately.
	LogpullOptions string `json:"logpull_options,nullable" format:"uri-reference"`
	// Optional human readable job name. Not unique. Cloudflare suggests that you set
	// this to a meaningful string, like the domain name, to make it easier to identify
	// your job.
	Name string                              `json:"name,nullable"`
	JSON zoneLogpushJobGetResponseResultJSON `json:"-"`
}

// zoneLogpushJobGetResponseResultJSON contains the JSON metadata for the struct
// [ZoneLogpushJobGetResponseResult]
type zoneLogpushJobGetResponseResultJSON struct {
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
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZoneLogpushJobGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The frequency at which Cloudflare sends batches of logs to your destination.
// Setting frequency to high sends your logs in larger quantities of smaller files.
// Setting frequency to low sends logs in smaller quantities of larger files.
type ZoneLogpushJobGetResponseResultFrequency string

const (
	ZoneLogpushJobGetResponseResultFrequencyHigh ZoneLogpushJobGetResponseResultFrequency = "high"
	ZoneLogpushJobGetResponseResultFrequencyLow  ZoneLogpushJobGetResponseResultFrequency = "low"
)

// Whether the API call was successful
type ZoneLogpushJobGetResponseSuccess bool

const (
	ZoneLogpushJobGetResponseSuccessTrue ZoneLogpushJobGetResponseSuccess = true
)

type ZoneLogpushJobUpdateResponse struct {
	Errors   []ZoneLogpushJobUpdateResponseError   `json:"errors"`
	Messages []ZoneLogpushJobUpdateResponseMessage `json:"messages"`
	Result   ZoneLogpushJobUpdateResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success ZoneLogpushJobUpdateResponseSuccess `json:"success"`
	JSON    zoneLogpushJobUpdateResponseJSON    `json:"-"`
}

// zoneLogpushJobUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneLogpushJobUpdateResponse]
type zoneLogpushJobUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushJobUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogpushJobUpdateResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zoneLogpushJobUpdateResponseErrorJSON `json:"-"`
}

// zoneLogpushJobUpdateResponseErrorJSON contains the JSON metadata for the struct
// [ZoneLogpushJobUpdateResponseError]
type zoneLogpushJobUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushJobUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogpushJobUpdateResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneLogpushJobUpdateResponseMessageJSON `json:"-"`
}

// zoneLogpushJobUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneLogpushJobUpdateResponseMessage]
type zoneLogpushJobUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushJobUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogpushJobUpdateResponseResult struct {
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
	Frequency ZoneLogpushJobUpdateResponseResultFrequency `json:"frequency,nullable"`
	// Records the last time for which logs have been successfully pushed. If the last
	// successful push was for logs range 2018-07-23T10:00:00Z to 2018-07-23T10:01:00Z
	// then the value of this field will be 2018-07-23T10:01:00Z. If the job has never
	// run or has just been enabled and hasn't run yet then the field will be empty.
	LastComplete time.Time `json:"last_complete,nullable" format:"date-time"`
	// Records the last time the job failed. If not null, the job is currently failing.
	// If null, the job has either never failed or has run successfully at least once
	// since last failure. See also the error_message field.
	LastError time.Time `json:"last_error,nullable" format:"date-time"`
	// Configuration string. It specifies things like requested fields and timestamp
	// formats. If migrating from the logpull api, copy the url (full url or just the
	// query string) of your call here, and logpush will keep on making this call for
	// you, setting start and end times appropriately.
	LogpullOptions string `json:"logpull_options,nullable" format:"uri-reference"`
	// Optional human readable job name. Not unique. Cloudflare suggests that you set
	// this to a meaningful string, like the domain name, to make it easier to identify
	// your job.
	Name string                                 `json:"name,nullable"`
	JSON zoneLogpushJobUpdateResponseResultJSON `json:"-"`
}

// zoneLogpushJobUpdateResponseResultJSON contains the JSON metadata for the struct
// [ZoneLogpushJobUpdateResponseResult]
type zoneLogpushJobUpdateResponseResultJSON struct {
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
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZoneLogpushJobUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The frequency at which Cloudflare sends batches of logs to your destination.
// Setting frequency to high sends your logs in larger quantities of smaller files.
// Setting frequency to low sends logs in smaller quantities of larger files.
type ZoneLogpushJobUpdateResponseResultFrequency string

const (
	ZoneLogpushJobUpdateResponseResultFrequencyHigh ZoneLogpushJobUpdateResponseResultFrequency = "high"
	ZoneLogpushJobUpdateResponseResultFrequencyLow  ZoneLogpushJobUpdateResponseResultFrequency = "low"
)

// Whether the API call was successful
type ZoneLogpushJobUpdateResponseSuccess bool

const (
	ZoneLogpushJobUpdateResponseSuccessTrue ZoneLogpushJobUpdateResponseSuccess = true
)

type ZoneLogpushJobDeleteResponse struct {
	Errors   []ZoneLogpushJobDeleteResponseError   `json:"errors"`
	Messages []ZoneLogpushJobDeleteResponseMessage `json:"messages"`
	Result   interface{}                           `json:"result,nullable"`
	// Whether the API call was successful
	Success ZoneLogpushJobDeleteResponseSuccess `json:"success"`
	JSON    zoneLogpushJobDeleteResponseJSON    `json:"-"`
}

// zoneLogpushJobDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneLogpushJobDeleteResponse]
type zoneLogpushJobDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushJobDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogpushJobDeleteResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zoneLogpushJobDeleteResponseErrorJSON `json:"-"`
}

// zoneLogpushJobDeleteResponseErrorJSON contains the JSON metadata for the struct
// [ZoneLogpushJobDeleteResponseError]
type zoneLogpushJobDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushJobDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogpushJobDeleteResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneLogpushJobDeleteResponseMessageJSON `json:"-"`
}

// zoneLogpushJobDeleteResponseMessageJSON contains the JSON metadata for the
// struct [ZoneLogpushJobDeleteResponseMessage]
type zoneLogpushJobDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushJobDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneLogpushJobDeleteResponseSuccess bool

const (
	ZoneLogpushJobDeleteResponseSuccessTrue ZoneLogpushJobDeleteResponseSuccess = true
)

type ZoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponse struct {
	Errors   []ZoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseError   `json:"errors"`
	Messages []ZoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseMessage `json:"messages"`
	Result   []ZoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseResult  `json:"result"`
	// Whether the API call was successful
	Success ZoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseSuccess `json:"success"`
	JSON    zoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseJSON    `json:"-"`
}

// zoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseJSON contains the JSON
// metadata for the struct
// [ZoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponse]
type zoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseError struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    zoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseErrorJSON `json:"-"`
}

// zoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseErrorJSON contains the
// JSON metadata for the struct
// [ZoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseError]
type zoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseMessage struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    zoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseMessageJSON `json:"-"`
}

// zoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseMessageJSON contains the
// JSON metadata for the struct
// [ZoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseMessage]
type zoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseResult struct {
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
	Frequency ZoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseResultFrequency `json:"frequency,nullable"`
	// Records the last time for which logs have been successfully pushed. If the last
	// successful push was for logs range 2018-07-23T10:00:00Z to 2018-07-23T10:01:00Z
	// then the value of this field will be 2018-07-23T10:01:00Z. If the job has never
	// run or has just been enabled and hasn't run yet then the field will be empty.
	LastComplete time.Time `json:"last_complete,nullable" format:"date-time"`
	// Records the last time the job failed. If not null, the job is currently failing.
	// If null, the job has either never failed or has run successfully at least once
	// since last failure. See also the error_message field.
	LastError time.Time `json:"last_error,nullable" format:"date-time"`
	// Configuration string. It specifies things like requested fields and timestamp
	// formats. If migrating from the logpull api, copy the url (full url or just the
	// query string) of your call here, and logpush will keep on making this call for
	// you, setting start and end times appropriately.
	LogpullOptions string `json:"logpull_options,nullable" format:"uri-reference"`
	// Optional human readable job name. Not unique. Cloudflare suggests that you set
	// this to a meaningful string, like the domain name, to make it easier to identify
	// your job.
	Name string                                                            `json:"name,nullable"`
	JSON zoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseResultJSON `json:"-"`
}

// zoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseResultJSON contains the
// JSON metadata for the struct
// [ZoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseResult]
type zoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseResultJSON struct {
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
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The frequency at which Cloudflare sends batches of logs to your destination.
// Setting frequency to high sends your logs in larger quantities of smaller files.
// Setting frequency to low sends logs in smaller quantities of larger files.
type ZoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseResultFrequency string

const (
	ZoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseResultFrequencyHigh ZoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseResultFrequency = "high"
	ZoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseResultFrequencyLow  ZoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseResultFrequency = "low"
)

// Whether the API call was successful
type ZoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseSuccess bool

const (
	ZoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseSuccessTrue ZoneLogpushJobGetZonesZoneIdentifierLogpushJobsResponseSuccess = true
)

type ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponse struct {
	Errors   []ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseError   `json:"errors"`
	Messages []ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseMessage `json:"messages"`
	Result   ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseSuccess `json:"success"`
	JSON    zoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseJSON    `json:"-"`
}

// zoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseJSON contains the JSON
// metadata for the struct
// [ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponse]
type zoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseError struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    zoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseErrorJSON `json:"-"`
}

// zoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseErrorJSON contains the
// JSON metadata for the struct
// [ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseError]
type zoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseMessage struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    zoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseMessageJSON `json:"-"`
}

// zoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseMessageJSON contains the
// JSON metadata for the struct
// [ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseMessage]
type zoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseResult struct {
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
	Frequency ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseResultFrequency `json:"frequency,nullable"`
	// Records the last time for which logs have been successfully pushed. If the last
	// successful push was for logs range 2018-07-23T10:00:00Z to 2018-07-23T10:01:00Z
	// then the value of this field will be 2018-07-23T10:01:00Z. If the job has never
	// run or has just been enabled and hasn't run yet then the field will be empty.
	LastComplete time.Time `json:"last_complete,nullable" format:"date-time"`
	// Records the last time the job failed. If not null, the job is currently failing.
	// If null, the job has either never failed or has run successfully at least once
	// since last failure. See also the error_message field.
	LastError time.Time `json:"last_error,nullable" format:"date-time"`
	// Configuration string. It specifies things like requested fields and timestamp
	// formats. If migrating from the logpull api, copy the url (full url or just the
	// query string) of your call here, and logpush will keep on making this call for
	// you, setting start and end times appropriately.
	LogpullOptions string `json:"logpull_options,nullable" format:"uri-reference"`
	// Optional human readable job name. Not unique. Cloudflare suggests that you set
	// this to a meaningful string, like the domain name, to make it easier to identify
	// your job.
	Name string                                                             `json:"name,nullable"`
	JSON zoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseResultJSON `json:"-"`
}

// zoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseResultJSON contains the
// JSON metadata for the struct
// [ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseResult]
type zoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseResultJSON struct {
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
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The frequency at which Cloudflare sends batches of logs to your destination.
// Setting frequency to high sends your logs in larger quantities of smaller files.
// Setting frequency to low sends logs in smaller quantities of larger files.
type ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseResultFrequency string

const (
	ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseResultFrequencyHigh ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseResultFrequency = "high"
	ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseResultFrequencyLow  ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseResultFrequency = "low"
)

// Whether the API call was successful
type ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseSuccess bool

const (
	ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseSuccessTrue ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsResponseSuccess = true
)

type ZoneLogpushJobUpdateParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf" format:"uri"`
	// Flag that indicates if the job is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// The frequency at which Cloudflare sends batches of logs to your destination.
	// Setting frequency to high sends your logs in larger quantities of smaller files.
	// Setting frequency to low sends logs in smaller quantities of larger files.
	Frequency param.Field[ZoneLogpushJobUpdateParamsFrequency] `json:"frequency"`
	// Configuration string. It specifies things like requested fields and timestamp
	// formats. If migrating from the logpull api, copy the url (full url or just the
	// query string) of your call here, and logpush will keep on making this call for
	// you, setting start and end times appropriately.
	LogpullOptions param.Field[string] `json:"logpull_options" format:"uri-reference"`
	// Ownership challenge token to prove destination ownership.
	OwnershipChallenge param.Field[string] `json:"ownership_challenge"`
}

func (r ZoneLogpushJobUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The frequency at which Cloudflare sends batches of logs to your destination.
// Setting frequency to high sends your logs in larger quantities of smaller files.
// Setting frequency to low sends logs in smaller quantities of larger files.
type ZoneLogpushJobUpdateParamsFrequency string

const (
	ZoneLogpushJobUpdateParamsFrequencyHigh ZoneLogpushJobUpdateParamsFrequency = "high"
	ZoneLogpushJobUpdateParamsFrequencyLow  ZoneLogpushJobUpdateParamsFrequency = "low"
)

type ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsParams struct {
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
	Frequency param.Field[ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsParamsFrequency] `json:"frequency"`
	// Configuration string. It specifies things like requested fields and timestamp
	// formats. If migrating from the logpull api, copy the url (full url or just the
	// query string) of your call here, and logpush will keep on making this call for
	// you, setting start and end times appropriately.
	LogpullOptions param.Field[string] `json:"logpull_options" format:"uri-reference"`
	// Optional human readable job name. Not unique. Cloudflare suggests that you set
	// this to a meaningful string, like the domain name, to make it easier to identify
	// your job.
	Name param.Field[string] `json:"name"`
	// Ownership challenge token to prove destination ownership.
	OwnershipChallenge param.Field[string] `json:"ownership_challenge"`
}

func (r ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The frequency at which Cloudflare sends batches of logs to your destination.
// Setting frequency to high sends your logs in larger quantities of smaller files.
// Setting frequency to low sends logs in smaller quantities of larger files.
type ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsParamsFrequency string

const (
	ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsParamsFrequencyHigh ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsParamsFrequency = "high"
	ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsParamsFrequencyLow  ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsParamsFrequency = "low"
)
