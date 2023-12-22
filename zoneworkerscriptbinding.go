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

// ZoneWorkerScriptBindingService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneWorkerScriptBindingService] method instead.
type ZoneWorkerScriptBindingService struct {
	Options []option.RequestOption
}

// NewZoneWorkerScriptBindingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneWorkerScriptBindingService(opts ...option.RequestOption) (r *ZoneWorkerScriptBindingService) {
	r = &ZoneWorkerScriptBindingService{}
	r.Options = opts
	return
}

// List the bindings for a Workers script.
func (r *ZoneWorkerScriptBindingService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneWorkerScriptBindingListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/workers/script/bindings", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneWorkerScriptBindingListResponse struct {
	Errors   []ZoneWorkerScriptBindingListResponseError   `json:"errors"`
	Messages []ZoneWorkerScriptBindingListResponseMessage `json:"messages"`
	Result   []ZoneWorkerScriptBindingListResponseResult  `json:"result"`
	// Whether the API call was successful
	Success ZoneWorkerScriptBindingListResponseSuccess `json:"success"`
	JSON    zoneWorkerScriptBindingListResponseJSON    `json:"-"`
}

// zoneWorkerScriptBindingListResponseJSON contains the JSON metadata for the
// struct [ZoneWorkerScriptBindingListResponse]
type zoneWorkerScriptBindingListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerScriptBindingListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWorkerScriptBindingListResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zoneWorkerScriptBindingListResponseErrorJSON `json:"-"`
}

// zoneWorkerScriptBindingListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneWorkerScriptBindingListResponseError]
type zoneWorkerScriptBindingListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerScriptBindingListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWorkerScriptBindingListResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneWorkerScriptBindingListResponseMessageJSON `json:"-"`
}

// zoneWorkerScriptBindingListResponseMessageJSON contains the JSON metadata for
// the struct [ZoneWorkerScriptBindingListResponseMessage]
type zoneWorkerScriptBindingListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerScriptBindingListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZoneWorkerScriptBindingListResponseResultKvNamespaceBinding]
// or [ZoneWorkerScriptBindingListResponseResultWasmModuleBinding].
type ZoneWorkerScriptBindingListResponseResult interface {
	implementsZoneWorkerScriptBindingListResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneWorkerScriptBindingListResponseResult)(nil)).Elem(), "")
}

type ZoneWorkerScriptBindingListResponseResultKvNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type ZoneWorkerScriptBindingListResponseResultKvNamespaceBindingType `json:"type,required"`
	JSON zoneWorkerScriptBindingListResponseResultKvNamespaceBindingJSON `json:"-"`
}

// zoneWorkerScriptBindingListResponseResultKvNamespaceBindingJSON contains the
// JSON metadata for the struct
// [ZoneWorkerScriptBindingListResponseResultKvNamespaceBinding]
type zoneWorkerScriptBindingListResponseResultKvNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerScriptBindingListResponseResultKvNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneWorkerScriptBindingListResponseResultKvNamespaceBinding) implementsZoneWorkerScriptBindingListResponseResult() {
}

// The class of resource that the binding provides.
type ZoneWorkerScriptBindingListResponseResultKvNamespaceBindingType string

const (
	ZoneWorkerScriptBindingListResponseResultKvNamespaceBindingTypeKvNamespace ZoneWorkerScriptBindingListResponseResultKvNamespaceBindingType = "kv_namespace"
)

type ZoneWorkerScriptBindingListResponseResultWasmModuleBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type ZoneWorkerScriptBindingListResponseResultWasmModuleBindingType `json:"type,required"`
	JSON zoneWorkerScriptBindingListResponseResultWasmModuleBindingJSON `json:"-"`
}

// zoneWorkerScriptBindingListResponseResultWasmModuleBindingJSON contains the JSON
// metadata for the struct
// [ZoneWorkerScriptBindingListResponseResultWasmModuleBinding]
type zoneWorkerScriptBindingListResponseResultWasmModuleBindingJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWorkerScriptBindingListResponseResultWasmModuleBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneWorkerScriptBindingListResponseResultWasmModuleBinding) implementsZoneWorkerScriptBindingListResponseResult() {
}

// The class of resource that the binding provides.
type ZoneWorkerScriptBindingListResponseResultWasmModuleBindingType string

const (
	ZoneWorkerScriptBindingListResponseResultWasmModuleBindingTypeWasmModule ZoneWorkerScriptBindingListResponseResultWasmModuleBindingType = "wasm_module"
)

// Whether the API call was successful
type ZoneWorkerScriptBindingListResponseSuccess bool

const (
	ZoneWorkerScriptBindingListResponseSuccessTrue ZoneWorkerScriptBindingListResponseSuccess = true
)
