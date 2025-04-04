// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// ScriptScheduleService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScriptScheduleService] method instead.
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
	var env ScriptScheduleUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
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
	var env ScriptScheduleGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/schedules", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Schedule struct {
	CreatedOn  string       `json:"created_on"`
	Cron       string       `json:"cron"`
	ModifiedOn string       `json:"modified_on"`
	JSON       scheduleJSON `json:"-"`
}

// scheduleJSON contains the JSON metadata for the struct [Schedule]
type scheduleJSON struct {
	CreatedOn   apijson.Field
	Cron        apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Schedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleJSON) RawJSON() string {
	return r.raw
}

type ScheduleParam struct {
	Cron param.Field[string] `json:"cron"`
}

func (r ScheduleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptScheduleUpdateResponse struct {
	Schedules []Schedule                       `json:"schedules"`
	JSON      scriptScheduleUpdateResponseJSON `json:"-"`
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

type ScriptScheduleGetResponse struct {
	Schedules []Schedule                    `json:"schedules"`
	JSON      scriptScheduleGetResponseJSON `json:"-"`
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

type ScriptScheduleUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Body      []ScheduleParam     `json:"body,required"`
}

func (r ScriptScheduleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ScriptScheduleUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ScriptScheduleUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  ScriptScheduleUpdateResponse                `json:"result"`
	JSON    scriptScheduleUpdateResponseEnvelopeJSON    `json:"-"`
}

// scriptScheduleUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [ScriptScheduleUpdateResponseEnvelope]
type scriptScheduleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScheduleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScheduleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptScheduleUpdateResponseEnvelopeSuccess bool

const (
	ScriptScheduleUpdateResponseEnvelopeSuccessTrue ScriptScheduleUpdateResponseEnvelopeSuccess = true
)

func (r ScriptScheduleUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptScheduleUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptScheduleGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptScheduleGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ScriptScheduleGetResponseEnvelopeSuccess `json:"success,required"`
	Result  ScriptScheduleGetResponse                `json:"result"`
	JSON    scriptScheduleGetResponseEnvelopeJSON    `json:"-"`
}

// scriptScheduleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptScheduleGetResponseEnvelope]
type scriptScheduleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptScheduleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptScheduleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptScheduleGetResponseEnvelopeSuccess bool

const (
	ScriptScheduleGetResponseEnvelopeSuccessTrue ScriptScheduleGetResponseEnvelopeSuccess = true
)

func (r ScriptScheduleGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptScheduleGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
