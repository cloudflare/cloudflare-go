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

// AccountLogpushJobService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountLogpushJobService] method
// instead.
type AccountLogpushJobService struct {
	Options []option.RequestOption
}

// NewAccountLogpushJobService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountLogpushJobService(opts ...option.RequestOption) (r *AccountLogpushJobService) {
	r = &AccountLogpushJobService{}
	r.Options = opts
	return
}

// Gets the details of a Logpush job.
func (r *AccountLogpushJobService) Get(ctx context.Context, accountIdentifier string, jobIdentifier int64, opts ...option.RequestOption) (res *LogpushJobResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/logpush/jobs/%v", accountIdentifier, jobIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a Logpush job.
func (r *AccountLogpushJobService) Update(ctx context.Context, accountIdentifier string, jobIdentifier int64, body AccountLogpushJobUpdateParams, opts ...option.RequestOption) (res *LogpushJobResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/logpush/jobs/%v", accountIdentifier, jobIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes a Logpush job.
func (r *AccountLogpushJobService) Delete(ctx context.Context, accountIdentifier string, jobIdentifier int64, opts ...option.RequestOption) (res *AccountLogpushJobDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/logpush/jobs/%v", accountIdentifier, jobIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Lists Logpush jobs for an account.
func (r *AccountLogpushJobService) GetAccountsAccountIdentifierLogpushJobs(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *LogpushJobResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/logpush/jobs", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Creates a new Logpush job for an account.
func (r *AccountLogpushJobService) PostAccountsAccountIdentifierLogpushJobs(ctx context.Context, accountIdentifier string, body AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsParams, opts ...option.RequestOption) (res *LogpushJobResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/logpush/jobs", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type LogpushJobResponseSingle struct {
	Errors   []LogpushJobResponseSingleError   `json:"errors"`
	Messages []LogpushJobResponseSingleMessage `json:"messages"`
	Result   LogpushJobResponseSingleResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success LogpushJobResponseSingleSuccess `json:"success"`
	JSON    logpushJobResponseSingleJSON    `json:"-"`
}

// logpushJobResponseSingleJSON contains the JSON metadata for the struct
// [LogpushJobResponseSingle]
type logpushJobResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushJobResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushJobResponseSingleError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    logpushJobResponseSingleErrorJSON `json:"-"`
}

// logpushJobResponseSingleErrorJSON contains the JSON metadata for the struct
// [LogpushJobResponseSingleError]
type logpushJobResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushJobResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushJobResponseSingleMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    logpushJobResponseSingleMessageJSON `json:"-"`
}

// logpushJobResponseSingleMessageJSON contains the JSON metadata for the struct
// [LogpushJobResponseSingleMessage]
type logpushJobResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushJobResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushJobResponseSingleResult struct {
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
	Frequency LogpushJobResponseSingleResultFrequency `json:"frequency,nullable"`
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
	Name string                             `json:"name,nullable"`
	JSON logpushJobResponseSingleResultJSON `json:"-"`
}

// logpushJobResponseSingleResultJSON contains the JSON metadata for the struct
// [LogpushJobResponseSingleResult]
type logpushJobResponseSingleResultJSON struct {
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

func (r *LogpushJobResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The frequency at which Cloudflare sends batches of logs to your destination.
// Setting frequency to high sends your logs in larger quantities of smaller files.
// Setting frequency to low sends logs in smaller quantities of larger files.
type LogpushJobResponseSingleResultFrequency string

const (
	LogpushJobResponseSingleResultFrequencyHigh LogpushJobResponseSingleResultFrequency = "high"
	LogpushJobResponseSingleResultFrequencyLow  LogpushJobResponseSingleResultFrequency = "low"
)

// Whether the API call was successful
type LogpushJobResponseSingleSuccess bool

const (
	LogpushJobResponseSingleSuccessTrue LogpushJobResponseSingleSuccess = true
)

type AccountLogpushJobDeleteResponse struct {
	Errors   []AccountLogpushJobDeleteResponseError   `json:"errors"`
	Messages []AccountLogpushJobDeleteResponseMessage `json:"messages"`
	Result   interface{}                              `json:"result,nullable"`
	// Whether the API call was successful
	Success AccountLogpushJobDeleteResponseSuccess `json:"success"`
	JSON    accountLogpushJobDeleteResponseJSON    `json:"-"`
}

// accountLogpushJobDeleteResponseJSON contains the JSON metadata for the struct
// [AccountLogpushJobDeleteResponse]
type accountLogpushJobDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogpushJobDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogpushJobDeleteResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accountLogpushJobDeleteResponseErrorJSON `json:"-"`
}

// accountLogpushJobDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountLogpushJobDeleteResponseError]
type accountLogpushJobDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogpushJobDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogpushJobDeleteResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountLogpushJobDeleteResponseMessageJSON `json:"-"`
}

// accountLogpushJobDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountLogpushJobDeleteResponseMessage]
type accountLogpushJobDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogpushJobDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountLogpushJobDeleteResponseSuccess bool

const (
	AccountLogpushJobDeleteResponseSuccessTrue AccountLogpushJobDeleteResponseSuccess = true
)

type AccountLogpushJobUpdateParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf" format:"uri"`
	// Flag that indicates if the job is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// The frequency at which Cloudflare sends batches of logs to your destination.
	// Setting frequency to high sends your logs in larger quantities of smaller files.
	// Setting frequency to low sends logs in smaller quantities of larger files.
	Frequency param.Field[AccountLogpushJobUpdateParamsFrequency] `json:"frequency"`
	// Configuration string. It specifies things like requested fields and timestamp
	// formats. If migrating from the logpull api, copy the url (full url or just the
	// query string) of your call here, and logpush will keep on making this call for
	// you, setting start and end times appropriately.
	LogpullOptions param.Field[string] `json:"logpull_options" format:"uri-reference"`
	// Ownership challenge token to prove destination ownership.
	OwnershipChallenge param.Field[string] `json:"ownership_challenge"`
}

func (r AccountLogpushJobUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The frequency at which Cloudflare sends batches of logs to your destination.
// Setting frequency to high sends your logs in larger quantities of smaller files.
// Setting frequency to low sends logs in smaller quantities of larger files.
type AccountLogpushJobUpdateParamsFrequency string

const (
	AccountLogpushJobUpdateParamsFrequencyHigh AccountLogpushJobUpdateParamsFrequency = "high"
	AccountLogpushJobUpdateParamsFrequencyLow  AccountLogpushJobUpdateParamsFrequency = "low"
)

type AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsParams struct {
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
	Frequency param.Field[AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsParamsFrequency] `json:"frequency"`
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

func (r AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The frequency at which Cloudflare sends batches of logs to your destination.
// Setting frequency to high sends your logs in larger quantities of smaller files.
// Setting frequency to low sends logs in smaller quantities of larger files.
type AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsParamsFrequency string

const (
	AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsParamsFrequencyHigh AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsParamsFrequency = "high"
	AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsParamsFrequencyLow  AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsParamsFrequency = "low"
)
