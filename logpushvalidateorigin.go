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

// LogpushValidateOriginService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLogpushValidateOriginService]
// method instead.
type LogpushValidateOriginService struct {
	Options []option.RequestOption
}

// NewLogpushValidateOriginService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLogpushValidateOriginService(opts ...option.RequestOption) (r *LogpushValidateOriginService) {
	r = &LogpushValidateOriginService{}
	r.Options = opts
	return
}

// Validates logpull origin with logpull_options.
func (r *LogpushValidateOriginService) PostAccountsAccountIdentifierLogpushValidateOrigin(ctx context.Context, accountOrZone string, accountOrZoneID string, body LogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginParams, opts ...option.RequestOption) (res *LogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseEnvelope
	path := fmt.Sprintf("%s/%s/logpush/validate/origin", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponse struct {
	Message string                                                                              `json:"message"`
	Valid   bool                                                                                `json:"valid"`
	JSON    logpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseJSON `json:"-"`
}

// logpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseJSON
// contains the JSON metadata for the struct
// [LogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponse]
type logpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseJSON struct {
	Message     apijson.Field
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginParams struct {
	// This field is deprecated. Use `output_options` instead. Configuration string. It
	// specifies things like requested fields and timestamp formats. If migrating from
	// the logpull api, copy the url (full url or just the query string) of your call
	// here, and logpush will keep on making this call for you, setting start and end
	// times appropriately.
	LogpullOptions param.Field[string] `json:"logpull_options,required" format:"uri-reference"`
}

func (r LogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseEnvelope struct {
	Errors   []LogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseEnvelopeMessages `json:"messages,required"`
	Result   LogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success LogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseEnvelopeSuccess `json:"success,required"`
	JSON    logpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseEnvelopeJSON    `json:"-"`
}

// logpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [LogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseEnvelope]
type logpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseEnvelopeErrors struct {
	Code    int64                                                                                             `json:"code,required"`
	Message string                                                                                            `json:"message,required"`
	JSON    logpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseEnvelopeErrorsJSON `json:"-"`
}

// logpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [LogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseEnvelopeErrors]
type logpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseEnvelopeMessages struct {
	Code    int64                                                                                               `json:"code,required"`
	Message string                                                                                              `json:"message,required"`
	JSON    logpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseEnvelopeMessagesJSON `json:"-"`
}

// logpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [LogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseEnvelopeMessages]
type logpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseEnvelopeSuccess bool

const (
	LogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseEnvelopeSuccessTrue LogpushValidateOriginPostAccountsAccountIdentifierLogpushValidateOriginResponseEnvelopeSuccess = true
)
