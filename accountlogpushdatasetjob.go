// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountLogpushDatasetJobService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountLogpushDatasetJobService] method instead.
type AccountLogpushDatasetJobService struct {
	Options []option.RequestOption
}

// NewAccountLogpushDatasetJobService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountLogpushDatasetJobService(opts ...option.RequestOption) (r *AccountLogpushDatasetJobService) {
	r = &AccountLogpushDatasetJobService{}
	r.Options = opts
	return
}

// Lists Logpush jobs for an account for a dataset.
func (r *AccountLogpushDatasetJobService) GetAccountsAccountIdentifierLogpushDatasetsDatasetJobs(ctx context.Context, accountIdentifier string, dataset string, opts ...option.RequestOption) (res *LogpushJobResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/logpush/datasets/%s/jobs", accountIdentifier, dataset)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LogpushJobResponseCollection struct {
	Errors   []LogpushJobResponseCollectionError   `json:"errors"`
	Messages []LogpushJobResponseCollectionMessage `json:"messages"`
	Result   []LogpushJobResponseCollectionResult  `json:"result"`
	// Whether the API call was successful
	Success LogpushJobResponseCollectionSuccess `json:"success"`
	JSON    logpushJobResponseCollectionJSON    `json:"-"`
}

// logpushJobResponseCollectionJSON contains the JSON metadata for the struct
// [LogpushJobResponseCollection]
type logpushJobResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushJobResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushJobResponseCollectionError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    logpushJobResponseCollectionErrorJSON `json:"-"`
}

// logpushJobResponseCollectionErrorJSON contains the JSON metadata for the struct
// [LogpushJobResponseCollectionError]
type logpushJobResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushJobResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushJobResponseCollectionMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    logpushJobResponseCollectionMessageJSON `json:"-"`
}

// logpushJobResponseCollectionMessageJSON contains the JSON metadata for the
// struct [LogpushJobResponseCollectionMessage]
type logpushJobResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushJobResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushJobResponseCollectionResult struct {
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
	Frequency LogpushJobResponseCollectionResultFrequency `json:"frequency,nullable"`
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
	JSON logpushJobResponseCollectionResultJSON `json:"-"`
}

// logpushJobResponseCollectionResultJSON contains the JSON metadata for the struct
// [LogpushJobResponseCollectionResult]
type logpushJobResponseCollectionResultJSON struct {
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

func (r *LogpushJobResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The frequency at which Cloudflare sends batches of logs to your destination.
// Setting frequency to high sends your logs in larger quantities of smaller files.
// Setting frequency to low sends logs in smaller quantities of larger files.
type LogpushJobResponseCollectionResultFrequency string

const (
	LogpushJobResponseCollectionResultFrequencyHigh LogpushJobResponseCollectionResultFrequency = "high"
	LogpushJobResponseCollectionResultFrequencyLow  LogpushJobResponseCollectionResultFrequency = "low"
)

// Whether the API call was successful
type LogpushJobResponseCollectionSuccess bool

const (
	LogpushJobResponseCollectionSuccessTrue LogpushJobResponseCollectionSuccess = true
)
