// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms

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

// DispatchNamespaceScriptSettingService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewDispatchNamespaceScriptSettingService] method instead.
type DispatchNamespaceScriptSettingService struct {
	Options []option.RequestOption
}

// NewDispatchNamespaceScriptSettingService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDispatchNamespaceScriptSettingService(opts ...option.RequestOption) (r *DispatchNamespaceScriptSettingService) {
	r = &DispatchNamespaceScriptSettingService{}
	r.Options = opts
	return
}

// Patch script metadata, such as bindings
func (r *DispatchNamespaceScriptSettingService) Edit(ctx context.Context, dispatchNamespace string, scriptName string, params DispatchNamespaceScriptSettingEditParams, opts ...option.RequestOption) (res *DispatchNamespaceScriptSettingEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DispatchNamespaceScriptSettingEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/settings", params.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get script settings from a script uploaded to a Workers for Platforms namespace.
func (r *DispatchNamespaceScriptSettingService) Get(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptSettingGetParams, opts ...option.RequestOption) (res *DispatchNamespaceScriptSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DispatchNamespaceScriptSettingGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/settings", query.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DispatchNamespaceScriptSettingEditResponse struct {
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []DispatchNamespaceScriptSettingEditResponseTailConsumer `json:"tail_consumers"`
	JSON          dispatchNamespaceScriptSettingEditResponseJSON           `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseJSON contains the JSON metadata for
// the struct [DispatchNamespaceScriptSettingEditResponse]
type dispatchNamespaceScriptSettingEditResponseJSON struct {
	Logpush       apijson.Field
	TailConsumers apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseJSON) RawJSON() string {
	return r.raw
}

// A reference to a script that will consume logs from the attached Worker.
type DispatchNamespaceScriptSettingEditResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                     `json:"namespace"`
	JSON      dispatchNamespaceScriptSettingEditResponseTailConsumerJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseTailConsumerJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptSettingEditResponseTailConsumer]
type dispatchNamespaceScriptSettingEditResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseTailConsumerJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSettingGetResponse struct {
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []DispatchNamespaceScriptSettingGetResponseTailConsumer `json:"tail_consumers"`
	JSON          dispatchNamespaceScriptSettingGetResponseJSON           `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseJSON contains the JSON metadata for the
// struct [DispatchNamespaceScriptSettingGetResponse]
type dispatchNamespaceScriptSettingGetResponseJSON struct {
	Logpush       apijson.Field
	TailConsumers apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseJSON) RawJSON() string {
	return r.raw
}

// A reference to a script that will consume logs from the attached Worker.
type DispatchNamespaceScriptSettingGetResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                    `json:"namespace"`
	JSON      dispatchNamespaceScriptSettingGetResponseTailConsumerJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseTailConsumerJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptSettingGetResponseTailConsumer]
type dispatchNamespaceScriptSettingGetResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseTailConsumerJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSettingEditParams struct {
	// Identifier
	AccountID param.Field[string]                                         `path:"account_id,required"`
	Errors    param.Field[[]shared.UnnamedSchemaRef172Param]              `json:"errors,required"`
	Messages  param.Field[[]shared.UnnamedSchemaRef172Param]              `json:"messages,required"`
	Result    param.Field[DispatchNamespaceScriptSettingEditParamsResult] `json:"result,required"`
	// Whether the API call was successful
	Success param.Field[DispatchNamespaceScriptSettingEditParamsSuccess] `json:"success,required"`
}

func (r DispatchNamespaceScriptSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceScriptSettingEditParamsResult struct {
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]DispatchNamespaceScriptSettingEditParamsResultTailConsumer] `json:"tail_consumers"`
}

func (r DispatchNamespaceScriptSettingEditParamsResult) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A reference to a script that will consume logs from the attached Worker.
type DispatchNamespaceScriptSettingEditParamsResultTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service param.Field[string] `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace param.Field[string] `json:"namespace"`
}

func (r DispatchNamespaceScriptSettingEditParamsResultTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the API call was successful
type DispatchNamespaceScriptSettingEditParamsSuccess bool

const (
	DispatchNamespaceScriptSettingEditParamsSuccessTrue DispatchNamespaceScriptSettingEditParamsSuccess = true
)

func (r DispatchNamespaceScriptSettingEditParamsSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsSuccessTrue:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172               `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172               `json:"messages,required"`
	Result   DispatchNamespaceScriptSettingEditResponse `json:"result,required"`
	// Whether the API call was successful
	Success DispatchNamespaceScriptSettingEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    dispatchNamespaceScriptSettingEditResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseEnvelopeJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptSettingEditResponseEnvelope]
type dispatchNamespaceScriptSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DispatchNamespaceScriptSettingEditResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptSettingEditResponseEnvelopeSuccessTrue DispatchNamespaceScriptSettingEditResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptSettingEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DispatchNamespaceScriptSettingGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172              `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172              `json:"messages,required"`
	Result   DispatchNamespaceScriptSettingGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success DispatchNamespaceScriptSettingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dispatchNamespaceScriptSettingGetResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [DispatchNamespaceScriptSettingGetResponseEnvelope]
type dispatchNamespaceScriptSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DispatchNamespaceScriptSettingGetResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptSettingGetResponseEnvelopeSuccessTrue DispatchNamespaceScriptSettingGetResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptSettingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
