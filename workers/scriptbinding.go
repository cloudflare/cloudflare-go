// File generated from our OpenAPI spec by Stainless.

package workers

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *ScriptBindingService) Get(ctx context.Context, query ScriptBindingGetParams, opts ...option.RequestOption) (res *[]ScriptBindingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ScriptBindingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/script/bindings", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [workers.ScriptBindingGetResponseWorkersKVNamespaceBinding]
// or [workers.ScriptBindingGetResponseWorkersWasmModuleBinding].
type ScriptBindingGetResponse interface {
	implementsWorkersScriptBindingGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScriptBindingGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptBindingGetResponseWorkersKVNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptBindingGetResponseWorkersWasmModuleBinding{}),
		},
	)
}

type ScriptBindingGetResponseWorkersKVNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type ScriptBindingGetResponseWorkersKVNamespaceBindingType `json:"type,required"`
	JSON scriptBindingGetResponseWorkersKVNamespaceBindingJSON `json:"-"`
}

// scriptBindingGetResponseWorkersKVNamespaceBindingJSON contains the JSON metadata
// for the struct [ScriptBindingGetResponseWorkersKVNamespaceBinding]
type scriptBindingGetResponseWorkersKVNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptBindingGetResponseWorkersKVNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptBindingGetResponseWorkersKVNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptBindingGetResponseWorkersKVNamespaceBinding) implementsWorkersScriptBindingGetResponse() {
}

// The class of resource that the binding provides.
type ScriptBindingGetResponseWorkersKVNamespaceBindingType string

const (
	ScriptBindingGetResponseWorkersKVNamespaceBindingTypeKVNamespace ScriptBindingGetResponseWorkersKVNamespaceBindingType = "kv_namespace"
)

type ScriptBindingGetResponseWorkersWasmModuleBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type ScriptBindingGetResponseWorkersWasmModuleBindingType `json:"type,required"`
	JSON scriptBindingGetResponseWorkersWasmModuleBindingJSON `json:"-"`
}

// scriptBindingGetResponseWorkersWasmModuleBindingJSON contains the JSON metadata
// for the struct [ScriptBindingGetResponseWorkersWasmModuleBinding]
type scriptBindingGetResponseWorkersWasmModuleBindingJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptBindingGetResponseWorkersWasmModuleBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptBindingGetResponseWorkersWasmModuleBindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptBindingGetResponseWorkersWasmModuleBinding) implementsWorkersScriptBindingGetResponse() {
}

// The class of resource that the binding provides.
type ScriptBindingGetResponseWorkersWasmModuleBindingType string

const (
	ScriptBindingGetResponseWorkersWasmModuleBindingTypeWasmModule ScriptBindingGetResponseWorkersWasmModuleBindingType = "wasm_module"
)

type ScriptBindingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ScriptBindingGetResponseEnvelope struct {
	Errors   []ScriptBindingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptBindingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []ScriptBindingGetResponse                 `json:"result,required"`
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
