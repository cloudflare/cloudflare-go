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

// LogpushValidateService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLogpushValidateService] method
// instead.
type LogpushValidateService struct {
	Options []option.RequestOption
}

// NewLogpushValidateService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLogpushValidateService(opts ...option.RequestOption) (r *LogpushValidateService) {
	r = &LogpushValidateService{}
	r.Options = opts
	return
}

// Checks if there is an existing job with a destination.
func (r *LogpushValidateService) Destination(ctx context.Context, params LogpushValidateDestinationParams, opts ...option.RequestOption) (res *LogpushValidateDestinationResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LogpushValidateDestinationResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/logpush/validate/destination/exists", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Validates logpull origin with logpull_options.
func (r *LogpushValidateService) Origin(ctx context.Context, params LogpushValidateOriginParams, opts ...option.RequestOption) (res *LogpushValidateOriginResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LogpushValidateOriginResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/logpush/validate/origin", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LogpushValidateDestinationResponse struct {
	Exists bool                                   `json:"exists"`
	JSON   logpushValidateDestinationResponseJSON `json:"-"`
}

// logpushValidateDestinationResponseJSON contains the JSON metadata for the struct
// [LogpushValidateDestinationResponse]
type logpushValidateDestinationResponseJSON struct {
	Exists      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushValidateDestinationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushValidateOriginResponse struct {
	Message string                            `json:"message"`
	Valid   bool                              `json:"valid"`
	JSON    logpushValidateOriginResponseJSON `json:"-"`
}

// logpushValidateOriginResponseJSON contains the JSON metadata for the struct
// [LogpushValidateOriginResponse]
type logpushValidateOriginResponseJSON struct {
	Message     apijson.Field
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushValidateOriginResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushValidateDestinationParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf,required" format:"uri"`
}

func (r LogpushValidateDestinationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LogpushValidateDestinationResponseEnvelope struct {
	Errors   []LogpushValidateDestinationResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LogpushValidateDestinationResponseEnvelopeMessages `json:"messages,required"`
	Result   LogpushValidateDestinationResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success LogpushValidateDestinationResponseEnvelopeSuccess `json:"success,required"`
	JSON    logpushValidateDestinationResponseEnvelopeJSON    `json:"-"`
}

// logpushValidateDestinationResponseEnvelopeJSON contains the JSON metadata for
// the struct [LogpushValidateDestinationResponseEnvelope]
type logpushValidateDestinationResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushValidateDestinationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushValidateDestinationResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    logpushValidateDestinationResponseEnvelopeErrorsJSON `json:"-"`
}

// logpushValidateDestinationResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [LogpushValidateDestinationResponseEnvelopeErrors]
type logpushValidateDestinationResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushValidateDestinationResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushValidateDestinationResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    logpushValidateDestinationResponseEnvelopeMessagesJSON `json:"-"`
}

// logpushValidateDestinationResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [LogpushValidateDestinationResponseEnvelopeMessages]
type logpushValidateDestinationResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushValidateDestinationResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LogpushValidateDestinationResponseEnvelopeSuccess bool

const (
	LogpushValidateDestinationResponseEnvelopeSuccessTrue LogpushValidateDestinationResponseEnvelopeSuccess = true
)

type LogpushValidateOriginParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// This field is deprecated. Use `output_options` instead. Configuration string. It
	// specifies things like requested fields and timestamp formats. If migrating from
	// the logpull api, copy the url (full url or just the query string) of your call
	// here, and logpush will keep on making this call for you, setting start and end
	// times appropriately.
	LogpullOptions param.Field[string] `json:"logpull_options,required" format:"uri-reference"`
}

func (r LogpushValidateOriginParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LogpushValidateOriginResponseEnvelope struct {
	Errors   []LogpushValidateOriginResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LogpushValidateOriginResponseEnvelopeMessages `json:"messages,required"`
	Result   LogpushValidateOriginResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success LogpushValidateOriginResponseEnvelopeSuccess `json:"success,required"`
	JSON    logpushValidateOriginResponseEnvelopeJSON    `json:"-"`
}

// logpushValidateOriginResponseEnvelopeJSON contains the JSON metadata for the
// struct [LogpushValidateOriginResponseEnvelope]
type logpushValidateOriginResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushValidateOriginResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushValidateOriginResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    logpushValidateOriginResponseEnvelopeErrorsJSON `json:"-"`
}

// logpushValidateOriginResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [LogpushValidateOriginResponseEnvelopeErrors]
type logpushValidateOriginResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushValidateOriginResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushValidateOriginResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    logpushValidateOriginResponseEnvelopeMessagesJSON `json:"-"`
}

// logpushValidateOriginResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [LogpushValidateOriginResponseEnvelopeMessages]
type logpushValidateOriginResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushValidateOriginResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LogpushValidateOriginResponseEnvelopeSuccess bool

const (
	LogpushValidateOriginResponseEnvelopeSuccessTrue LogpushValidateOriginResponseEnvelopeSuccess = true
)
