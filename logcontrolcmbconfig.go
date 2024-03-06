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

// Updates CMB config.
func (r *LogControlCmbConfigService) New(ctx context.Context, params LogControlCmbConfigNewParams, opts ...option.RequestOption) (res *LogControlCmbConfigNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LogControlCmbConfigNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/logs/control/cmb/config", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes CMB config.
func (r *LogControlCmbConfigService) Delete(ctx context.Context, body LogControlCmbConfigDeleteParams, opts ...option.RequestOption) (res *LogControlCmbConfigDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LogControlCmbConfigDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/logs/control/cmb/config", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets CMB config.
func (r *LogControlCmbConfigService) Get(ctx context.Context, query LogControlCmbConfigGetParams, opts ...option.RequestOption) (res *LogControlCmbConfigGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LogControlCmbConfigGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/logs/control/cmb/config", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LogControlCmbConfigNewResponse struct {
	// Comma-separated list of regions.
	Regions string                             `json:"regions"`
	JSON    logControlCmbConfigNewResponseJSON `json:"-"`
}

// logControlCmbConfigNewResponseJSON contains the JSON metadata for the struct
// [LogControlCmbConfigNewResponse]
type logControlCmbConfigNewResponseJSON struct {
	Regions     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlCmbConfigNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logControlCmbConfigNewResponseJSON) RawJSON() string {
	return r.raw
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
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(LogControlCmbConfigDeleteResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type LogControlCmbConfigDeleteResponseArray []interface{}

func (r LogControlCmbConfigDeleteResponseArray) ImplementsLogControlCmbConfigDeleteResponse() {}

type LogControlCmbConfigGetResponse struct {
	// Comma-separated list of regions.
	Regions string                             `json:"regions"`
	JSON    logControlCmbConfigGetResponseJSON `json:"-"`
}

// logControlCmbConfigGetResponseJSON contains the JSON metadata for the struct
// [LogControlCmbConfigGetResponse]
type logControlCmbConfigGetResponseJSON struct {
	Regions     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlCmbConfigGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logControlCmbConfigGetResponseJSON) RawJSON() string {
	return r.raw
}

type LogControlCmbConfigNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Comma-separated list of regions.
	Regions param.Field[string] `json:"regions"`
}

func (r LogControlCmbConfigNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LogControlCmbConfigNewResponseEnvelope struct {
	Errors   []LogControlCmbConfigNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LogControlCmbConfigNewResponseEnvelopeMessages `json:"messages,required"`
	Result   LogControlCmbConfigNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success LogControlCmbConfigNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    logControlCmbConfigNewResponseEnvelopeJSON    `json:"-"`
}

// logControlCmbConfigNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [LogControlCmbConfigNewResponseEnvelope]
type logControlCmbConfigNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlCmbConfigNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logControlCmbConfigNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LogControlCmbConfigNewResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    logControlCmbConfigNewResponseEnvelopeErrorsJSON `json:"-"`
}

// logControlCmbConfigNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [LogControlCmbConfigNewResponseEnvelopeErrors]
type logControlCmbConfigNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlCmbConfigNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logControlCmbConfigNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type LogControlCmbConfigNewResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    logControlCmbConfigNewResponseEnvelopeMessagesJSON `json:"-"`
}

// logControlCmbConfigNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [LogControlCmbConfigNewResponseEnvelopeMessages]
type logControlCmbConfigNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlCmbConfigNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logControlCmbConfigNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LogControlCmbConfigNewResponseEnvelopeSuccess bool

const (
	LogControlCmbConfigNewResponseEnvelopeSuccessTrue LogControlCmbConfigNewResponseEnvelopeSuccess = true
)

type LogControlCmbConfigDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
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

func (r logControlCmbConfigDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r logControlCmbConfigDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r logControlCmbConfigDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LogControlCmbConfigDeleteResponseEnvelopeSuccess bool

const (
	LogControlCmbConfigDeleteResponseEnvelopeSuccessTrue LogControlCmbConfigDeleteResponseEnvelopeSuccess = true
)

type LogControlCmbConfigGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type LogControlCmbConfigGetResponseEnvelope struct {
	Errors   []LogControlCmbConfigGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LogControlCmbConfigGetResponseEnvelopeMessages `json:"messages,required"`
	Result   LogControlCmbConfigGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success LogControlCmbConfigGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    logControlCmbConfigGetResponseEnvelopeJSON    `json:"-"`
}

// logControlCmbConfigGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [LogControlCmbConfigGetResponseEnvelope]
type logControlCmbConfigGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlCmbConfigGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logControlCmbConfigGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LogControlCmbConfigGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    logControlCmbConfigGetResponseEnvelopeErrorsJSON `json:"-"`
}

// logControlCmbConfigGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [LogControlCmbConfigGetResponseEnvelopeErrors]
type logControlCmbConfigGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlCmbConfigGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logControlCmbConfigGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type LogControlCmbConfigGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    logControlCmbConfigGetResponseEnvelopeMessagesJSON `json:"-"`
}

// logControlCmbConfigGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [LogControlCmbConfigGetResponseEnvelopeMessages]
type logControlCmbConfigGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogControlCmbConfigGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logControlCmbConfigGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LogControlCmbConfigGetResponseEnvelopeSuccess bool

const (
	LogControlCmbConfigGetResponseEnvelopeSuccessTrue LogControlCmbConfigGetResponseEnvelopeSuccess = true
)
