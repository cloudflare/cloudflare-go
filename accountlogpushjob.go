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
func (r *AccountLogpushJobService) Get(ctx context.Context, accountIdentifier string, jobIdentifier int64, opts ...option.RequestOption) (res *AccountLogpushJobGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/logpush/jobs/%v", accountIdentifier, jobIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a Logpush job.
func (r *AccountLogpushJobService) Update(ctx context.Context, accountIdentifier string, jobIdentifier int64, body AccountLogpushJobUpdateParams, opts ...option.RequestOption) (res *AccountLogpushJobUpdateResponse, err error) {
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
func (r *AccountLogpushJobService) GetAccountsAccountIdentifierLogpushJobs(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/logpush/jobs", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Creates a new Logpush job for an account.
func (r *AccountLogpushJobService) PostAccountsAccountIdentifierLogpushJobs(ctx context.Context, accountIdentifier string, body AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsParams, opts ...option.RequestOption) (res *AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/logpush/jobs", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountLogpushJobGetResponse struct {
	Errors   []AccountLogpushJobGetResponseError   `json:"errors"`
	Messages []AccountLogpushJobGetResponseMessage `json:"messages"`
	Result   AccountLogpushJobGetResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success AccountLogpushJobGetResponseSuccess `json:"success"`
	JSON    accountLogpushJobGetResponseJSON    `json:"-"`
}

// accountLogpushJobGetResponseJSON contains the JSON metadata for the struct
// [AccountLogpushJobGetResponse]
type accountLogpushJobGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogpushJobGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogpushJobGetResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    accountLogpushJobGetResponseErrorJSON `json:"-"`
}

// accountLogpushJobGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountLogpushJobGetResponseError]
type accountLogpushJobGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogpushJobGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogpushJobGetResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accountLogpushJobGetResponseMessageJSON `json:"-"`
}

// accountLogpushJobGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountLogpushJobGetResponseMessage]
type accountLogpushJobGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogpushJobGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogpushJobGetResponseResult struct {
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
	Frequency AccountLogpushJobGetResponseResultFrequency `json:"frequency,nullable"`
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
	JSON accountLogpushJobGetResponseResultJSON `json:"-"`
}

// accountLogpushJobGetResponseResultJSON contains the JSON metadata for the struct
// [AccountLogpushJobGetResponseResult]
type accountLogpushJobGetResponseResultJSON struct {
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

func (r *AccountLogpushJobGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The frequency at which Cloudflare sends batches of logs to your destination.
// Setting frequency to high sends your logs in larger quantities of smaller files.
// Setting frequency to low sends logs in smaller quantities of larger files.
type AccountLogpushJobGetResponseResultFrequency string

const (
	AccountLogpushJobGetResponseResultFrequencyHigh AccountLogpushJobGetResponseResultFrequency = "high"
	AccountLogpushJobGetResponseResultFrequencyLow  AccountLogpushJobGetResponseResultFrequency = "low"
)

// Whether the API call was successful
type AccountLogpushJobGetResponseSuccess bool

const (
	AccountLogpushJobGetResponseSuccessTrue AccountLogpushJobGetResponseSuccess = true
)

type AccountLogpushJobUpdateResponse struct {
	Errors   []AccountLogpushJobUpdateResponseError   `json:"errors"`
	Messages []AccountLogpushJobUpdateResponseMessage `json:"messages"`
	Result   AccountLogpushJobUpdateResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success AccountLogpushJobUpdateResponseSuccess `json:"success"`
	JSON    accountLogpushJobUpdateResponseJSON    `json:"-"`
}

// accountLogpushJobUpdateResponseJSON contains the JSON metadata for the struct
// [AccountLogpushJobUpdateResponse]
type accountLogpushJobUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogpushJobUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogpushJobUpdateResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accountLogpushJobUpdateResponseErrorJSON `json:"-"`
}

// accountLogpushJobUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccountLogpushJobUpdateResponseError]
type accountLogpushJobUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogpushJobUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogpushJobUpdateResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountLogpushJobUpdateResponseMessageJSON `json:"-"`
}

// accountLogpushJobUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccountLogpushJobUpdateResponseMessage]
type accountLogpushJobUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogpushJobUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogpushJobUpdateResponseResult struct {
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
	Frequency AccountLogpushJobUpdateResponseResultFrequency `json:"frequency,nullable"`
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
	Name string                                    `json:"name,nullable"`
	JSON accountLogpushJobUpdateResponseResultJSON `json:"-"`
}

// accountLogpushJobUpdateResponseResultJSON contains the JSON metadata for the
// struct [AccountLogpushJobUpdateResponseResult]
type accountLogpushJobUpdateResponseResultJSON struct {
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

func (r *AccountLogpushJobUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The frequency at which Cloudflare sends batches of logs to your destination.
// Setting frequency to high sends your logs in larger quantities of smaller files.
// Setting frequency to low sends logs in smaller quantities of larger files.
type AccountLogpushJobUpdateResponseResultFrequency string

const (
	AccountLogpushJobUpdateResponseResultFrequencyHigh AccountLogpushJobUpdateResponseResultFrequency = "high"
	AccountLogpushJobUpdateResponseResultFrequencyLow  AccountLogpushJobUpdateResponseResultFrequency = "low"
)

// Whether the API call was successful
type AccountLogpushJobUpdateResponseSuccess bool

const (
	AccountLogpushJobUpdateResponseSuccessTrue AccountLogpushJobUpdateResponseSuccess = true
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

type AccountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponse struct {
	Errors   []AccountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseError   `json:"errors"`
	Messages []AccountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseMessage `json:"messages"`
	Result   []AccountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseResult  `json:"result"`
	// Whether the API call was successful
	Success AccountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseSuccess `json:"success"`
	JSON    accountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseJSON    `json:"-"`
}

// accountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseJSON contains
// the JSON metadata for the struct
// [AccountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponse]
type accountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseError struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    accountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseErrorJSON `json:"-"`
}

// accountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseError]
type accountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseMessage struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    accountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseMessageJSON `json:"-"`
}

// accountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseMessage]
type accountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseResult struct {
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
	Frequency AccountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseResultFrequency `json:"frequency,nullable"`
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
	Name string                                                                     `json:"name,nullable"`
	JSON accountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseResultJSON `json:"-"`
}

// accountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseResult]
type accountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseResultJSON struct {
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

func (r *AccountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The frequency at which Cloudflare sends batches of logs to your destination.
// Setting frequency to high sends your logs in larger quantities of smaller files.
// Setting frequency to low sends logs in smaller quantities of larger files.
type AccountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseResultFrequency string

const (
	AccountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseResultFrequencyHigh AccountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseResultFrequency = "high"
	AccountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseResultFrequencyLow  AccountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseResultFrequency = "low"
)

// Whether the API call was successful
type AccountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseSuccess bool

const (
	AccountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseSuccessTrue AccountLogpushJobGetAccountsAccountIdentifierLogpushJobsResponseSuccess = true
)

type AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponse struct {
	Errors   []AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseError   `json:"errors"`
	Messages []AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseMessage `json:"messages"`
	Result   AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseSuccess `json:"success"`
	JSON    accountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseJSON    `json:"-"`
}

// accountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseJSON contains
// the JSON metadata for the struct
// [AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponse]
type accountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseError struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    accountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseErrorJSON `json:"-"`
}

// accountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseError]
type accountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseMessage struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    accountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseMessageJSON `json:"-"`
}

// accountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseMessage]
type accountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseResult struct {
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
	Frequency AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseResultFrequency `json:"frequency,nullable"`
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
	Name string                                                                      `json:"name,nullable"`
	JSON accountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseResultJSON `json:"-"`
}

// accountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseResult]
type accountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseResultJSON struct {
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

func (r *AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The frequency at which Cloudflare sends batches of logs to your destination.
// Setting frequency to high sends your logs in larger quantities of smaller files.
// Setting frequency to low sends logs in smaller quantities of larger files.
type AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseResultFrequency string

const (
	AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseResultFrequencyHigh AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseResultFrequency = "high"
	AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseResultFrequencyLow  AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseResultFrequency = "low"
)

// Whether the API call was successful
type AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseSuccess bool

const (
	AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseSuccessTrue AccountLogpushJobPostAccountsAccountIdentifierLogpushJobsResponseSuccess = true
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
