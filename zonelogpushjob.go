// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

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
func (r *ZoneLogpushJobService) Get(ctx context.Context, zoneIdentifier string, jobIdentifier int64, opts ...option.RequestOption) (res *LogpushJobResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/logpush/jobs/%v", zoneIdentifier, jobIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a Logpush job.
func (r *ZoneLogpushJobService) Update(ctx context.Context, zoneIdentifier string, jobIdentifier int64, body ZoneLogpushJobUpdateParams, opts ...option.RequestOption) (res *LogpushJobResponseSingle, err error) {
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
func (r *ZoneLogpushJobService) GetZonesZoneIdentifierLogpushJobs(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *LogpushJobResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/logpush/jobs", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Creates a new Logpush job for a zone.
func (r *ZoneLogpushJobService) PostZonesZoneIdentifierLogpushJobs(ctx context.Context, zoneIdentifier string, body ZoneLogpushJobPostZonesZoneIdentifierLogpushJobsParams, opts ...option.RequestOption) (res *LogpushJobResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/logpush/jobs", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

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
