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

// LogpushOwnershipService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLogpushOwnershipService] method
// instead.
type LogpushOwnershipService struct {
	Options []option.RequestOption
}

// NewLogpushOwnershipService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLogpushOwnershipService(opts ...option.RequestOption) (r *LogpushOwnershipService) {
	r = &LogpushOwnershipService{}
	r.Options = opts
	return
}

// Gets a new ownership challenge sent to your destination.
func (r *LogpushOwnershipService) New(ctx context.Context, params LogpushOwnershipNewParams, opts ...option.RequestOption) (res *LogpushOwnershipNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LogpushOwnershipNewResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/logpush/ownership", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Validates ownership challenge of the destination.
func (r *LogpushOwnershipService) Validate(ctx context.Context, params LogpushOwnershipValidateParams, opts ...option.RequestOption) (res *LogpushOwnershipValidateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LogpushOwnershipValidateResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/logpush/ownership/validate", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LogpushOwnershipNewResponse struct {
	Filename string                          `json:"filename"`
	Message  string                          `json:"message"`
	Valid    bool                            `json:"valid"`
	JSON     logpushOwnershipNewResponseJSON `json:"-"`
}

// logpushOwnershipNewResponseJSON contains the JSON metadata for the struct
// [LogpushOwnershipNewResponse]
type logpushOwnershipNewResponseJSON struct {
	Filename    apijson.Field
	Message     apijson.Field
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushOwnershipNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushOwnershipValidateResponse struct {
	Valid bool                                 `json:"valid"`
	JSON  logpushOwnershipValidateResponseJSON `json:"-"`
}

// logpushOwnershipValidateResponseJSON contains the JSON metadata for the struct
// [LogpushOwnershipValidateResponse]
type logpushOwnershipValidateResponseJSON struct {
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushOwnershipValidateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushOwnershipNewParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf,required" format:"uri"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r LogpushOwnershipNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LogpushOwnershipNewResponseEnvelope struct {
	Errors   []LogpushOwnershipNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LogpushOwnershipNewResponseEnvelopeMessages `json:"messages,required"`
	Result   LogpushOwnershipNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success LogpushOwnershipNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    logpushOwnershipNewResponseEnvelopeJSON    `json:"-"`
}

// logpushOwnershipNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [LogpushOwnershipNewResponseEnvelope]
type logpushOwnershipNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushOwnershipNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushOwnershipNewResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    logpushOwnershipNewResponseEnvelopeErrorsJSON `json:"-"`
}

// logpushOwnershipNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [LogpushOwnershipNewResponseEnvelopeErrors]
type logpushOwnershipNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushOwnershipNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushOwnershipNewResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    logpushOwnershipNewResponseEnvelopeMessagesJSON `json:"-"`
}

// logpushOwnershipNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [LogpushOwnershipNewResponseEnvelopeMessages]
type logpushOwnershipNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushOwnershipNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LogpushOwnershipNewResponseEnvelopeSuccess bool

const (
	LogpushOwnershipNewResponseEnvelopeSuccessTrue LogpushOwnershipNewResponseEnvelopeSuccess = true
)

type LogpushOwnershipValidateParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf,required" format:"uri"`
	// Ownership challenge token to prove destination ownership.
	OwnershipChallenge param.Field[string] `json:"ownership_challenge,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r LogpushOwnershipValidateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LogpushOwnershipValidateResponseEnvelope struct {
	Errors   []LogpushOwnershipValidateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LogpushOwnershipValidateResponseEnvelopeMessages `json:"messages,required"`
	Result   LogpushOwnershipValidateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success LogpushOwnershipValidateResponseEnvelopeSuccess `json:"success,required"`
	JSON    logpushOwnershipValidateResponseEnvelopeJSON    `json:"-"`
}

// logpushOwnershipValidateResponseEnvelopeJSON contains the JSON metadata for the
// struct [LogpushOwnershipValidateResponseEnvelope]
type logpushOwnershipValidateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushOwnershipValidateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushOwnershipValidateResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    logpushOwnershipValidateResponseEnvelopeErrorsJSON `json:"-"`
}

// logpushOwnershipValidateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [LogpushOwnershipValidateResponseEnvelopeErrors]
type logpushOwnershipValidateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushOwnershipValidateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushOwnershipValidateResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    logpushOwnershipValidateResponseEnvelopeMessagesJSON `json:"-"`
}

// logpushOwnershipValidateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [LogpushOwnershipValidateResponseEnvelopeMessages]
type logpushOwnershipValidateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushOwnershipValidateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LogpushOwnershipValidateResponseEnvelopeSuccess bool

const (
	LogpushOwnershipValidateResponseEnvelopeSuccessTrue LogpushOwnershipValidateResponseEnvelopeSuccess = true
)
