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

// LogpushOwnershipValidateService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewLogpushOwnershipValidateService] method instead.
type LogpushOwnershipValidateService struct {
	Options []option.RequestOption
}

// NewLogpushOwnershipValidateService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewLogpushOwnershipValidateService(opts ...option.RequestOption) (r *LogpushOwnershipValidateService) {
	r = &LogpushOwnershipValidateService{}
	r.Options = opts
	return
}

// Validates ownership challenge of the destination.
func (r *LogpushOwnershipValidateService) PostAccountsAccountIdentifierLogpushOwnershipValidate(ctx context.Context, accountOrZone string, accountOrZoneID string, body LogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateParams, opts ...option.RequestOption) (res *LogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseEnvelope
	path := fmt.Sprintf("%s/%s/logpush/ownership/validate", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponse struct {
	Valid bool                                                                                      `json:"valid"`
	JSON  logpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseJSON `json:"-"`
}

// logpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseJSON
// contains the JSON metadata for the struct
// [LogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponse]
type logpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseJSON struct {
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf,required" format:"uri"`
	// Ownership challenge token to prove destination ownership.
	OwnershipChallenge param.Field[string] `json:"ownership_challenge,required"`
}

func (r LogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseEnvelope struct {
	Errors   []LogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseEnvelopeMessages `json:"messages,required"`
	Result   LogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success LogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseEnvelopeSuccess `json:"success,required"`
	JSON    logpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseEnvelopeJSON    `json:"-"`
}

// logpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [LogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseEnvelope]
type logpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseEnvelopeErrors struct {
	Code    int64                                                                                                   `json:"code,required"`
	Message string                                                                                                  `json:"message,required"`
	JSON    logpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseEnvelopeErrorsJSON `json:"-"`
}

// logpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [LogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseEnvelopeErrors]
type logpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseEnvelopeMessages struct {
	Code    int64                                                                                                     `json:"code,required"`
	Message string                                                                                                    `json:"message,required"`
	JSON    logpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseEnvelopeMessagesJSON `json:"-"`
}

// logpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [LogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseEnvelopeMessages]
type logpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseEnvelopeSuccess bool

const (
	LogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseEnvelopeSuccessTrue LogpushOwnershipValidatePostAccountsAccountIdentifierLogpushOwnershipValidateResponseEnvelopeSuccess = true
)
