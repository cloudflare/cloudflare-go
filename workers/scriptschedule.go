// File generated from our OpenAPI spec by Stainless.

package workers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// ScriptScheduleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewScriptScheduleService] method
// instead.
type ScriptScheduleService struct {
	Options []option.RequestOption
}

// NewScriptScheduleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScriptScheduleService(opts ...option.RequestOption) (r *ScriptScheduleService) {
	r = &ScriptScheduleService{}
	r.Options = opts
	return
}

// Updates Cron Triggers for a Worker.
func (r *ScriptScheduleService) Update(ctx context.Context, scriptName string, params ScriptScheduleUpdateParams, opts ...option.RequestOption) (res *ScriptScheduleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ScriptScheduleUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/schedules", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches Cron Triggers for a Worker.
func (r *ScriptScheduleService) Get(ctx context.Context, scriptName string, query ScriptScheduleGetParams, opts ...option.RequestOption) (res *ScriptScheduleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ScriptScheduleGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/schedules", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ScriptScheduleUpdateResponse struct {
	Schedules []ScriptScheduleUpdateResponseSchedule `json:"schedules"`
	JSON      scriptScheduleUpdateResponseJSON       `json:"-"`
}

// scriptScheduleUpdateResponseJSON contains the JSON metadata for the struct
// [ScriptScheduleUpdateResponse]
type scriptScheduleUpdateResponseJSON struct {
	Schedules   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScheduleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScheduleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptScheduleUpdateResponseSchedule struct {
	CreatedOn  interface{}                              `json:"created_on"`
	Cron       interface{}                              `json:"cron"`
	ModifiedOn interface{}                              `json:"modified_on"`
	JSON       scriptScheduleUpdateResponseScheduleJSON `json:"-"`
}

// scriptScheduleUpdateResponseScheduleJSON contains the JSON metadata for the
// struct [ScriptScheduleUpdateResponseSchedule]
type scriptScheduleUpdateResponseScheduleJSON struct {
	CreatedOn   apijson.Field
	Cron        apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScheduleUpdateResponseSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScheduleUpdateResponseScheduleJSON) RawJSON() string {
	return r.raw
}

type ScriptScheduleGetResponse struct {
	Schedules []ScriptScheduleGetResponseSchedule `json:"schedules"`
	JSON      scriptScheduleGetResponseJSON       `json:"-"`
}

// scriptScheduleGetResponseJSON contains the JSON metadata for the struct
// [ScriptScheduleGetResponse]
type scriptScheduleGetResponseJSON struct {
	Schedules   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScheduleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScheduleGetResponseJSON) RawJSON() string {
	return r.raw
}

type ScriptScheduleGetResponseSchedule struct {
	CreatedOn  interface{}                           `json:"created_on"`
	Cron       interface{}                           `json:"cron"`
	ModifiedOn interface{}                           `json:"modified_on"`
	JSON       scriptScheduleGetResponseScheduleJSON `json:"-"`
}

// scriptScheduleGetResponseScheduleJSON contains the JSON metadata for the struct
// [ScriptScheduleGetResponseSchedule]
type scriptScheduleGetResponseScheduleJSON struct {
	CreatedOn   apijson.Field
	Cron        apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScheduleGetResponseSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScheduleGetResponseScheduleJSON) RawJSON() string {
	return r.raw
}

type ScriptScheduleUpdateParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r ScriptScheduleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ScriptScheduleUpdateResponseEnvelope struct {
	Errors   []ScriptScheduleUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptScheduleUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ScriptScheduleUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ScriptScheduleUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptScheduleUpdateResponseEnvelopeJSON    `json:"-"`
}

// scriptScheduleUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [ScriptScheduleUpdateResponseEnvelope]
type scriptScheduleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScheduleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScheduleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptScheduleUpdateResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    scriptScheduleUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// scriptScheduleUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ScriptScheduleUpdateResponseEnvelopeErrors]
type scriptScheduleUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScheduleUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScheduleUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptScheduleUpdateResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    scriptScheduleUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// scriptScheduleUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ScriptScheduleUpdateResponseEnvelopeMessages]
type scriptScheduleUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScheduleUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScheduleUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptScheduleUpdateResponseEnvelopeSuccess bool

const (
	ScriptScheduleUpdateResponseEnvelopeSuccessTrue ScriptScheduleUpdateResponseEnvelopeSuccess = true
)

type ScriptScheduleGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptScheduleGetResponseEnvelope struct {
	Errors   []ScriptScheduleGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptScheduleGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ScriptScheduleGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ScriptScheduleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptScheduleGetResponseEnvelopeJSON    `json:"-"`
}

// scriptScheduleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptScheduleGetResponseEnvelope]
type scriptScheduleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScheduleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScheduleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptScheduleGetResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    scriptScheduleGetResponseEnvelopeErrorsJSON `json:"-"`
}

// scriptScheduleGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ScriptScheduleGetResponseEnvelopeErrors]
type scriptScheduleGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScheduleGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScheduleGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptScheduleGetResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    scriptScheduleGetResponseEnvelopeMessagesJSON `json:"-"`
}

// scriptScheduleGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ScriptScheduleGetResponseEnvelopeMessages]
type scriptScheduleGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScheduleGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScheduleGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptScheduleGetResponseEnvelopeSuccess bool

const (
	ScriptScheduleGetResponseEnvelopeSuccessTrue ScriptScheduleGetResponseEnvelopeSuccess = true
)
