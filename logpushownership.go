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
	Options   []option.RequestOption
	Validates *LogpushOwnershipValidateService
}

// NewLogpushOwnershipService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLogpushOwnershipService(opts ...option.RequestOption) (r *LogpushOwnershipService) {
	r = &LogpushOwnershipService{}
	r.Options = opts
	r.Validates = NewLogpushOwnershipValidateService(opts...)
	return
}

// Gets a new ownership challenge sent to your destination.
func (r *LogpushOwnershipService) PostAccountsAccountIdentifierLogpushOwnership(ctx context.Context, accountOrZone string, accountOrZoneID string, body LogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipParams, opts ...option.RequestOption) (res *LogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseEnvelope
	path := fmt.Sprintf("%s/%s/logpush/ownership", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponse struct {
	Filename string                                                                    `json:"filename"`
	Message  string                                                                    `json:"message"`
	Valid    bool                                                                      `json:"valid"`
	JSON     logpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseJSON `json:"-"`
}

// logpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseJSON
// contains the JSON metadata for the struct
// [LogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponse]
type logpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseJSON struct {
	Filename    apijson.Field
	Message     apijson.Field
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf,required" format:"uri"`
}

func (r LogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseEnvelope struct {
	Errors   []LogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseEnvelopeMessages `json:"messages,required"`
	Result   LogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success LogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseEnvelopeSuccess `json:"success,required"`
	JSON    logpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseEnvelopeJSON    `json:"-"`
}

// logpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [LogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseEnvelope]
type logpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseEnvelopeErrors struct {
	Code    int64                                                                                   `json:"code,required"`
	Message string                                                                                  `json:"message,required"`
	JSON    logpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseEnvelopeErrorsJSON `json:"-"`
}

// logpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [LogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseEnvelopeErrors]
type logpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseEnvelopeMessages struct {
	Code    int64                                                                                     `json:"code,required"`
	Message string                                                                                    `json:"message,required"`
	JSON    logpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseEnvelopeMessagesJSON `json:"-"`
}

// logpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [LogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseEnvelopeMessages]
type logpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseEnvelopeSuccess bool

const (
	LogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseEnvelopeSuccessTrue LogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipResponseEnvelopeSuccess = true
)
