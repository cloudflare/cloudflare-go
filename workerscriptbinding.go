// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
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
func (r *WorkerScriptBindingService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *[]WorkerScriptBindingListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerScriptBindingListResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/script/bindings", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [WorkerScriptBindingListResponseWorkersKVNamespaceBinding] or
// [WorkerScriptBindingListResponseWorkersWasmModuleBinding].
type WorkerScriptBindingListResponse interface {
	implementsWorkerScriptBindingListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkerScriptBindingListResponse)(nil)).Elem(), "")
}

type WorkerScriptBindingListResponseWorkersKVNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptBindingListResponseWorkersKVNamespaceBindingType `json:"type,required"`
	JSON workerScriptBindingListResponseWorkersKVNamespaceBindingJSON `json:"-"`
}

// workerScriptBindingListResponseWorkersKVNamespaceBindingJSON contains the JSON
// metadata for the struct
// [WorkerScriptBindingListResponseWorkersKVNamespaceBinding]
type workerScriptBindingListResponseWorkersKVNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptBindingListResponseWorkersKVNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptBindingListResponseWorkersKVNamespaceBinding) implementsWorkerScriptBindingListResponse() {
}

// The class of resource that the binding provides.
type WorkerScriptBindingListResponseWorkersKVNamespaceBindingType string

const (
	WorkerScriptBindingListResponseWorkersKVNamespaceBindingTypeKVNamespace WorkerScriptBindingListResponseWorkersKVNamespaceBindingType = "kv_namespace"
)

type WorkerScriptBindingListResponseWorkersWasmModuleBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptBindingListResponseWorkersWasmModuleBindingType `json:"type,required"`
	JSON workerScriptBindingListResponseWorkersWasmModuleBindingJSON `json:"-"`
}

// workerScriptBindingListResponseWorkersWasmModuleBindingJSON contains the JSON
// metadata for the struct
// [WorkerScriptBindingListResponseWorkersWasmModuleBinding]
type workerScriptBindingListResponseWorkersWasmModuleBindingJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptBindingListResponseWorkersWasmModuleBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptBindingListResponseWorkersWasmModuleBinding) implementsWorkerScriptBindingListResponse() {
}

// The class of resource that the binding provides.
type WorkerScriptBindingListResponseWorkersWasmModuleBindingType string

const (
	WorkerScriptBindingListResponseWorkersWasmModuleBindingTypeWasmModule WorkerScriptBindingListResponseWorkersWasmModuleBindingType = "wasm_module"
)

type WorkerScriptBindingListResponseEnvelope struct {
	Errors   []WorkerScriptBindingListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerScriptBindingListResponseEnvelopeMessages `json:"messages,required"`
	Result   []WorkerScriptBindingListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success WorkerScriptBindingListResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerScriptBindingListResponseEnvelopeJSON    `json:"-"`
}

// workerScriptBindingListResponseEnvelopeJSON contains the JSON metadata for the
// struct [WorkerScriptBindingListResponseEnvelope]
type workerScriptBindingListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptBindingListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptBindingListResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    workerScriptBindingListResponseEnvelopeErrorsJSON `json:"-"`
}

// workerScriptBindingListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [WorkerScriptBindingListResponseEnvelopeErrors]
type workerScriptBindingListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptBindingListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptBindingListResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    workerScriptBindingListResponseEnvelopeMessagesJSON `json:"-"`
}

// workerScriptBindingListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [WorkerScriptBindingListResponseEnvelopeMessages]
type workerScriptBindingListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptBindingListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerScriptBindingListResponseEnvelopeSuccess bool

const (
	WorkerScriptBindingListResponseEnvelopeSuccessTrue WorkerScriptBindingListResponseEnvelopeSuccess = true
)
