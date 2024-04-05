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
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
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
func (r *ScriptBindingService) Get(ctx context.Context, query ScriptBindingGetParams, opts ...option.RequestOption) (res *[]Binding, err error) {
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

type Binding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The class of resource that the binding provides.
	Type  BindingType `json:"type,required"`
	JSON  bindingJSON `json:"-"`
	union BindingUnion
}

// bindingJSON contains the JSON metadata for the struct [Binding]
type bindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r bindingJSON) RawJSON() string {
	return r.raw
}

func (r *Binding) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r Binding) AsUnion() BindingUnion {
	return r.union
}

// Union satisfied by [workers.KVNamespaceBinding] or
// [workers.BindingWorkersWasmModuleBinding].
type BindingUnion interface {
	implementsWorkersBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BindingUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(KVNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BindingWorkersWasmModuleBinding{}),
		},
	)
}

type BindingWorkersWasmModuleBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type BindingWorkersWasmModuleBindingType `json:"type,required"`
	JSON bindingWorkersWasmModuleBindingJSON `json:"-"`
}

// bindingWorkersWasmModuleBindingJSON contains the JSON metadata for the struct
// [BindingWorkersWasmModuleBinding]
type bindingWorkersWasmModuleBindingJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingWorkersWasmModuleBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingWorkersWasmModuleBindingJSON) RawJSON() string {
	return r.raw
}

func (r BindingWorkersWasmModuleBinding) implementsWorkersBinding() {}

// The class of resource that the binding provides.
type BindingWorkersWasmModuleBindingType string

const (
	BindingWorkersWasmModuleBindingTypeWasmModule BindingWorkersWasmModuleBindingType = "wasm_module"
)

func (r BindingWorkersWasmModuleBindingType) IsKnown() bool {
	switch r {
	case BindingWorkersWasmModuleBindingTypeWasmModule:
		return true
	}
	return false
}

// The class of resource that the binding provides.
type BindingType string

const (
	BindingTypeKVNamespace BindingType = "kv_namespace"
	BindingTypeWasmModule  BindingType = "wasm_module"
)

func (r BindingType) IsKnown() bool {
	switch r {
	case BindingTypeKVNamespace, BindingTypeWasmModule:
		return true
	}
	return false
}

type ScriptBindingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ScriptBindingGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   []Binding                                                 `json:"result,required"`
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

// Whether the API call was successful
type ScriptBindingGetResponseEnvelopeSuccess bool

const (
	ScriptBindingGetResponseEnvelopeSuccessTrue ScriptBindingGetResponseEnvelopeSuccess = true
)

func (r ScriptBindingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptBindingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
