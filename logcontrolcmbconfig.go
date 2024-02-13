// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// LogControlCmbConfigService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLogControlCmbConfigService]
// method instead.
type LogControlCmbConfigService struct {
	Options []option.RequestOption
}

// NewLogControlCmbConfigService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLogControlCmbConfigService(opts ...option.RequestOption) (r *LogControlCmbConfigService) {
	r = &LogControlCmbConfigService{}
	r.Options = opts
	return
}

// Deletes CMB config.
func (r *LogControlCmbConfigService) Delete(ctx context.Context, accountID string, opts ...option.RequestOption) (res *LogControlCmbConfigDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LogControlCmbConfigDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/logs/control/cmb/config", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets CMB config.
func (r *LogControlCmbConfigService) GetAccountsAccountIdentifierLogsControlCmbConfig(ctx context.Context, accountID string, opts ...option.RequestOption) (res *LogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelope
	path := fmt.Sprintf("accounts/%s/logs/control/cmb/config", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates CMB config.
func (r *LogControlCmbConfigService) PutAccountsAccountIdentifierLogsControlCmbConfig(ctx context.Context, accountID string, body LogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigParams, opts ...option.RequestOption) (res *LogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelope
	path := fmt.Sprintf("accounts/%s/logs/control/cmb/config", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [LogControlCmbConfigDeleteResponseUnknown],
// [LogControlCmbConfigDeleteResponseArray] or [shared.UnionString].
type LogControlCmbConfigDeleteResponse interface {
	ImplementsLogControlCmbConfigDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*LogControlCmbConfigDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type LogControlCmbConfigDeleteResponseArray []interface{}

func (r LogControlCmbConfigDeleteResponseArray) ImplementsLogControlCmbConfigDeleteResponse() {}

type LogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponse struct {
	// Comma-separated list of regions.
	Regions string                                                                          `json:"regions"`
	JSON    logControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseJSON `json:"-"`
}

// logControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseJSON
// contains the JSON metadata for the struct
// [LogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponse]
type logControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseJSON struct {
	Regions     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponse struct {
	// Comma-separated list of regions.
	Regions string                                                                          `json:"regions"`
	JSON    logControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseJSON `json:"-"`
}

// logControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseJSON
// contains the JSON metadata for the struct
// [LogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponse]
type logControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseJSON struct {
	Regions     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogControlCmbConfigDeleteResponseEnvelope struct {
	Errors   []LogControlCmbConfigDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LogControlCmbConfigDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   LogControlCmbConfigDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success LogControlCmbConfigDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    logControlCmbConfigDeleteResponseEnvelopeJSON    `json:"-"`
}

// logControlCmbConfigDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [LogControlCmbConfigDeleteResponseEnvelope]
type logControlCmbConfigDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlCmbConfigDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogControlCmbConfigDeleteResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    logControlCmbConfigDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// logControlCmbConfigDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [LogControlCmbConfigDeleteResponseEnvelopeErrors]
type logControlCmbConfigDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlCmbConfigDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogControlCmbConfigDeleteResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    logControlCmbConfigDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// logControlCmbConfigDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [LogControlCmbConfigDeleteResponseEnvelopeMessages]
type logControlCmbConfigDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlCmbConfigDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LogControlCmbConfigDeleteResponseEnvelopeSuccess bool

const (
	LogControlCmbConfigDeleteResponseEnvelopeSuccessTrue LogControlCmbConfigDeleteResponseEnvelopeSuccess = true
)

type LogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelope struct {
	Errors   []LogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeMessages `json:"messages,required"`
	Result   LogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success LogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeSuccess `json:"success,required"`
	JSON    logControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeJSON    `json:"-"`
}

// logControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [LogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelope]
type logControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeErrors struct {
	Code    int64                                                                                         `json:"code,required"`
	Message string                                                                                        `json:"message,required"`
	JSON    logControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeErrorsJSON `json:"-"`
}

// logControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [LogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeErrors]
type logControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeMessages struct {
	Code    int64                                                                                           `json:"code,required"`
	Message string                                                                                          `json:"message,required"`
	JSON    logControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeMessagesJSON `json:"-"`
}

// logControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [LogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeMessages]
type logControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeSuccess bool

const (
	LogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeSuccessTrue LogControlCmbConfigGetAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeSuccess = true
)

type LogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigParams struct {
	// Comma-separated list of regions.
	Regions param.Field[string] `json:"regions"`
}

func (r LogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelope struct {
	Errors   []LogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeMessages `json:"messages,required"`
	Result   LogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success LogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeSuccess `json:"success,required"`
	JSON    logControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeJSON    `json:"-"`
}

// logControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [LogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelope]
type logControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeErrors struct {
	Code    int64                                                                                         `json:"code,required"`
	Message string                                                                                        `json:"message,required"`
	JSON    logControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeErrorsJSON `json:"-"`
}

// logControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [LogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeErrors]
type logControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeMessages struct {
	Code    int64                                                                                           `json:"code,required"`
	Message string                                                                                          `json:"message,required"`
	JSON    logControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeMessagesJSON `json:"-"`
}

// logControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [LogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeMessages]
type logControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeSuccess bool

const (
	LogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeSuccessTrue LogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigResponseEnvelopeSuccess = true
)
