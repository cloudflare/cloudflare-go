// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// ScriptSettingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewScriptSettingService] method
// instead.
type ScriptSettingService struct {
	Options []option.RequestOption
}

// NewScriptSettingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScriptSettingService(opts ...option.RequestOption) (r *ScriptSettingService) {
	r = &ScriptSettingService{}
	r.Options = opts
	return
}

// Patch script-level settings when using
// [Worker Versions](https://developers.cloudflare.com/api/operations/worker-versions-list-versions).
// Includes Logpush and Tail Consumers.
func (r *ScriptSettingService) Edit(ctx context.Context, scriptName string, params ScriptSettingEditParams, opts ...option.RequestOption) (res *ScriptSettingEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ScriptSettingEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/script-settings", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get script-level settings when using
// [Worker Versions](https://developers.cloudflare.com/api/operations/worker-versions-list-versions).
// Includes Logpush and Tail Consumers.
func (r *ScriptSettingService) Get(ctx context.Context, scriptName string, query ScriptSettingGetParams, opts ...option.RequestOption) (res *ScriptSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ScriptSettingGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/script-settings", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ScriptSettingEditResponse struct {
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []ScriptSettingEditResponseTailConsumer `json:"tail_consumers"`
	JSON          scriptSettingEditResponseJSON           `json:"-"`
}

// scriptSettingEditResponseJSON contains the JSON metadata for the struct
// [ScriptSettingEditResponse]
type scriptSettingEditResponseJSON struct {
	Logpush       apijson.Field
	TailConsumers apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScriptSettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseJSON) RawJSON() string {
	return r.raw
}

// A reference to a script that will consume logs from the attached Worker.
type ScriptSettingEditResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                    `json:"namespace"`
	JSON      scriptSettingEditResponseTailConsumerJSON `json:"-"`
}

// scriptSettingEditResponseTailConsumerJSON contains the JSON metadata for the
// struct [ScriptSettingEditResponseTailConsumer]
type scriptSettingEditResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingEditResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseTailConsumerJSON) RawJSON() string {
	return r.raw
}

type ScriptSettingGetResponse struct {
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []ScriptSettingGetResponseTailConsumer `json:"tail_consumers"`
	JSON          scriptSettingGetResponseJSON           `json:"-"`
}

// scriptSettingGetResponseJSON contains the JSON metadata for the struct
// [ScriptSettingGetResponse]
type scriptSettingGetResponseJSON struct {
	Logpush       apijson.Field
	TailConsumers apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScriptSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseJSON) RawJSON() string {
	return r.raw
}

// A reference to a script that will consume logs from the attached Worker.
type ScriptSettingGetResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                   `json:"namespace"`
	JSON      scriptSettingGetResponseTailConsumerJSON `json:"-"`
}

// scriptSettingGetResponseTailConsumerJSON contains the JSON metadata for the
// struct [ScriptSettingGetResponseTailConsumer]
type scriptSettingGetResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingGetResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseTailConsumerJSON) RawJSON() string {
	return r.raw
}

type ScriptSettingEditParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]ScriptSettingEditParamsTailConsumer] `json:"tail_consumers"`
}

func (r ScriptSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A reference to a script that will consume logs from the attached Worker.
type ScriptSettingEditParamsTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service param.Field[string] `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace param.Field[string] `json:"namespace"`
}

func (r ScriptSettingEditParamsTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptSettingEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo     `json:"errors,required"`
	Messages []shared.ResponseInfo     `json:"messages,required"`
	Result   ScriptSettingEditResponse `json:"result,required"`
	// Whether the API call was successful
	Success ScriptSettingEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptSettingEditResponseEnvelopeJSON    `json:"-"`
}

// scriptSettingEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptSettingEditResponseEnvelope]
type scriptSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptSettingEditResponseEnvelopeSuccess bool

const (
	ScriptSettingEditResponseEnvelopeSuccessTrue ScriptSettingEditResponseEnvelopeSuccess = true
)

func (r ScriptSettingEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptSettingEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptSettingGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptSettingGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo    `json:"errors,required"`
	Messages []shared.ResponseInfo    `json:"messages,required"`
	Result   ScriptSettingGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success ScriptSettingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptSettingGetResponseEnvelopeJSON    `json:"-"`
}

// scriptSettingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptSettingGetResponseEnvelope]
type scriptSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptSettingGetResponseEnvelopeSuccess bool

const (
	ScriptSettingGetResponseEnvelopeSuccessTrue ScriptSettingGetResponseEnvelopeSuccess = true
)

func (r ScriptSettingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptSettingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
