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

// LogpushValidateDestinationExistService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewLogpushValidateDestinationExistService] method instead.
type LogpushValidateDestinationExistService struct {
	Options []option.RequestOption
}

// NewLogpushValidateDestinationExistService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewLogpushValidateDestinationExistService(opts ...option.RequestOption) (r *LogpushValidateDestinationExistService) {
	r = &LogpushValidateDestinationExistService{}
	r.Options = opts
	return
}

// Checks if there is an existing job with a destination.
func (r *LogpushValidateDestinationExistService) DeleteAccountsAccountIdentifierLogpushValidateDestinationExists(ctx context.Context, accountOrZone string, accountOrZoneID string, body LogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsParams, opts ...option.RequestOption) (res *LogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseEnvelope
	path := fmt.Sprintf("%s/%s/logpush/validate/destination/exists", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponse struct {
	Exists bool                                                                                                       `json:"exists"`
	JSON   logpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseJSON `json:"-"`
}

// logpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseJSON
// contains the JSON metadata for the struct
// [LogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponse]
type logpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseJSON struct {
	Exists      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf,required" format:"uri"`
}

func (r LogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseEnvelope struct {
	Errors   []LogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseEnvelopeMessages `json:"messages,required"`
	Result   LogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success LogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseEnvelopeSuccess `json:"success,required"`
	JSON    logpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseEnvelopeJSON    `json:"-"`
}

// logpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [LogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseEnvelope]
type logpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseEnvelopeErrors struct {
	Code    int64                                                                                                                    `json:"code,required"`
	Message string                                                                                                                   `json:"message,required"`
	JSON    logpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseEnvelopeErrorsJSON `json:"-"`
}

// logpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [LogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseEnvelopeErrors]
type logpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseEnvelopeMessages struct {
	Code    int64                                                                                                                      `json:"code,required"`
	Message string                                                                                                                     `json:"message,required"`
	JSON    logpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseEnvelopeMessagesJSON `json:"-"`
}

// logpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [LogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseEnvelopeMessages]
type logpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseEnvelopeSuccess bool

const (
	LogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseEnvelopeSuccessTrue LogpushValidateDestinationExistDeleteAccountsAccountIdentifierLogpushValidateDestinationExistsResponseEnvelopeSuccess = true
)
