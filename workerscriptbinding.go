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
	"github.com/tidwall/gjson"
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
func (r *WorkerScriptBindingService) Get(ctx context.Context, query WorkerScriptBindingGetParams, opts ...option.RequestOption) (res *[]WorkerScriptBindingGetResponse, err error) {
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

// Union satisfied by [WorkerScriptBindingGetResponseWorkersKVNamespaceBinding] or
// [WorkerScriptBindingGetResponseWorkersWasmModuleBinding].
type WorkerScriptBindingGetResponse interface {
	implementsWorkerScriptBindingGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WorkerScriptBindingGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WorkerScriptBindingGetResponseWorkersKVNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WorkerScriptBindingGetResponseWorkersWasmModuleBinding{}),
		},
	)
}

type WorkerScriptBindingGetResponseWorkersKVNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptBindingGetResponseWorkersKVNamespaceBindingType `json:"type,required"`
	JSON workerScriptBindingGetResponseWorkersKVNamespaceBindingJSON `json:"-"`
}

// workerScriptBindingGetResponseWorkersKVNamespaceBindingJSON contains the JSON
// metadata for the struct
// [WorkerScriptBindingGetResponseWorkersKVNamespaceBinding]
type workerScriptBindingGetResponseWorkersKVNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptBindingGetResponseWorkersKVNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptBindingGetResponseWorkersKVNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkerScriptBindingGetResponseWorkersKVNamespaceBinding) implementsWorkerScriptBindingGetResponse() {
}

// The class of resource that the binding provides.
type WorkerScriptBindingGetResponseWorkersKVNamespaceBindingType string

const (
	WorkerScriptBindingGetResponseWorkersKVNamespaceBindingTypeKVNamespace WorkerScriptBindingGetResponseWorkersKVNamespaceBindingType = "kv_namespace"
)

type WorkerScriptBindingGetResponseWorkersWasmModuleBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptBindingGetResponseWorkersWasmModuleBindingType `json:"type,required"`
	JSON workerScriptBindingGetResponseWorkersWasmModuleBindingJSON `json:"-"`
}

// workerScriptBindingGetResponseWorkersWasmModuleBindingJSON contains the JSON
// metadata for the struct [WorkerScriptBindingGetResponseWorkersWasmModuleBinding]
type workerScriptBindingGetResponseWorkersWasmModuleBindingJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptBindingGetResponseWorkersWasmModuleBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptBindingGetResponseWorkersWasmModuleBindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkerScriptBindingGetResponseWorkersWasmModuleBinding) implementsWorkerScriptBindingGetResponse() {
}

// The class of resource that the binding provides.
type WorkerScriptBindingGetResponseWorkersWasmModuleBindingType string

const (
	WorkerScriptBindingGetResponseWorkersWasmModuleBindingTypeWasmModule WorkerScriptBindingGetResponseWorkersWasmModuleBindingType = "wasm_module"
)

type WorkerScriptBindingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type WorkerScriptBindingGetResponseEnvelope struct {
	Errors   []WorkerScriptBindingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerScriptBindingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []WorkerScriptBindingGetResponse                 `json:"result,required"`
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

func (r workerScriptBindingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r workerScriptBindingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r workerScriptBindingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WorkerScriptBindingGetResponseEnvelopeSuccess bool

const (
	WorkerScriptBindingGetResponseEnvelopeSuccessTrue WorkerScriptBindingGetResponseEnvelopeSuccess = true
)
