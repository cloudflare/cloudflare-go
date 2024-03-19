// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// ScriptBindingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewScriptBindingService] method
// instead.
type ScriptBindingService struct {
	Options []option.RequestOption
}

// NewScriptBindingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScriptBindingService(opts ...option.RequestOption) (r *ScriptBindingService) {
	r = &ScriptBindingService{}
	r.Options = opts
	return
}

// List the bindings for a Workers script.
func (r *ScriptBindingService) Get(ctx context.Context, query ScriptBindingGetParams, opts ...option.RequestOption) (res *[]WorkersSchemasBinding, err error) {
	opts = append(r.Options[:], opts...)
	var env ScriptBindingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/script/bindings", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [workers.WorkersSchemasBindingWorkersKVNamespaceBinding] or
// [workers.WorkersSchemasBindingWorkersWasmModuleBinding].
type WorkersSchemasBinding interface {
	implementsWorkersWorkersSchemasBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WorkersSchemasBinding)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WorkersSchemasBindingWorkersKVNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WorkersSchemasBindingWorkersWasmModuleBinding{}),
		},
	)
}

type WorkersSchemasBindingWorkersKVNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type WorkersSchemasBindingWorkersKVNamespaceBindingType `json:"type,required"`
	JSON workersSchemasBindingWorkersKVNamespaceBindingJSON `json:"-"`
}

// workersSchemasBindingWorkersKVNamespaceBindingJSON contains the JSON metadata
// for the struct [WorkersSchemasBindingWorkersKVNamespaceBinding]
type workersSchemasBindingWorkersKVNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersSchemasBindingWorkersKVNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersSchemasBindingWorkersKVNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkersSchemasBindingWorkersKVNamespaceBinding) implementsWorkersWorkersSchemasBinding() {}

// The class of resource that the binding provides.
type WorkersSchemasBindingWorkersKVNamespaceBindingType string

const (
	WorkersSchemasBindingWorkersKVNamespaceBindingTypeKVNamespace WorkersSchemasBindingWorkersKVNamespaceBindingType = "kv_namespace"
)

type WorkersSchemasBindingWorkersWasmModuleBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersSchemasBindingWorkersWasmModuleBindingType `json:"type,required"`
	JSON workersSchemasBindingWorkersWasmModuleBindingJSON `json:"-"`
}

// workersSchemasBindingWorkersWasmModuleBindingJSON contains the JSON metadata for
// the struct [WorkersSchemasBindingWorkersWasmModuleBinding]
type workersSchemasBindingWorkersWasmModuleBindingJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersSchemasBindingWorkersWasmModuleBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersSchemasBindingWorkersWasmModuleBindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkersSchemasBindingWorkersWasmModuleBinding) implementsWorkersWorkersSchemasBinding() {}

// The class of resource that the binding provides.
type WorkersSchemasBindingWorkersWasmModuleBindingType string

const (
	WorkersSchemasBindingWorkersWasmModuleBindingTypeWasmModule WorkersSchemasBindingWorkersWasmModuleBindingType = "wasm_module"
)

type ScriptBindingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ScriptBindingGetResponseEnvelope struct {
	Errors   []ScriptBindingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptBindingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []WorkersSchemasBinding                    `json:"result,required"`
	// Whether the API call was successful
	Success ScriptBindingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptBindingGetResponseEnvelopeJSON    `json:"-"`
}

// scriptBindingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptBindingGetResponseEnvelope]
type scriptBindingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptBindingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptBindingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScriptBindingGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    scriptBindingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// scriptBindingGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ScriptBindingGetResponseEnvelopeErrors]
type scriptBindingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptBindingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptBindingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptBindingGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    scriptBindingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// scriptBindingGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ScriptBindingGetResponseEnvelopeMessages]
type scriptBindingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptBindingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptBindingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptBindingGetResponseEnvelopeSuccess bool

const (
	ScriptBindingGetResponseEnvelopeSuccessTrue ScriptBindingGetResponseEnvelopeSuccess = true
)
