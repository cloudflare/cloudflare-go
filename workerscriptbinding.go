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
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// WorkerScriptBindingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWorkerScriptBindingService]
// method instead.
type WorkerScriptBindingService struct {
	Options []option.RequestOption
}

// NewWorkerScriptBindingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkerScriptBindingService(opts ...option.RequestOption) (r *WorkerScriptBindingService) {
	r = &WorkerScriptBindingService{}
	r.Options = opts
	return
}

// List the bindings for a Workers script.
func (r *WorkerScriptBindingService) Get(ctx context.Context, query WorkerScriptBindingGetParams, opts ...option.RequestOption) (res *[]WorkersSchemasBinding, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerScriptBindingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/script/bindings", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [WorkersSchemasBindingWorkersKVNamespaceBinding] or
// [WorkersSchemasBindingWorkersWasmModuleBinding].
type WorkersSchemasBinding interface {
	implementsWorkersSchemasBinding()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkersSchemasBinding)(nil)).Elem(), "")
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

func (r WorkersSchemasBindingWorkersKVNamespaceBinding) implementsWorkersSchemasBinding() {}

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

func (r WorkersSchemasBindingWorkersWasmModuleBinding) implementsWorkersSchemasBinding() {}

// The class of resource that the binding provides.
type WorkersSchemasBindingWorkersWasmModuleBindingType string

const (
	WorkersSchemasBindingWorkersWasmModuleBindingTypeWasmModule WorkersSchemasBindingWorkersWasmModuleBindingType = "wasm_module"
)

type WorkerScriptBindingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type WorkerScriptBindingGetResponseEnvelope struct {
	Errors   []WorkerScriptBindingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerScriptBindingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []WorkersSchemasBinding                          `json:"result,required"`
	// Whether the API call was successful
	Success WorkerScriptBindingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerScriptBindingGetResponseEnvelopeJSON    `json:"-"`
}

// workerScriptBindingGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [WorkerScriptBindingGetResponseEnvelope]
type workerScriptBindingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptBindingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptBindingGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    workerScriptBindingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// workerScriptBindingGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [WorkerScriptBindingGetResponseEnvelopeErrors]
type workerScriptBindingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptBindingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptBindingGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    workerScriptBindingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// workerScriptBindingGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [WorkerScriptBindingGetResponseEnvelopeMessages]
type workerScriptBindingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptBindingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerScriptBindingGetResponseEnvelopeSuccess bool

const (
	WorkerScriptBindingGetResponseEnvelopeSuccessTrue WorkerScriptBindingGetResponseEnvelopeSuccess = true
)
