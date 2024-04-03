// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms

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
	// List of bindings attached to this Worker
	Bindings []DispatchNamespaceScriptSettingEditResponseBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations DispatchNamespaceScriptSettingEditResponseMigrations `json:"migrations"`
	Placement  DispatchNamespaceScriptSettingEditResponsePlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []DispatchNamespaceScriptSettingEditResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                         `json:"usage_model"`
	JSON       dispatchNamespaceScriptSettingEditResponseJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseJSON contains the JSON metadata for
// the struct [DispatchNamespaceScriptSettingEditResponse]
type dispatchNamespaceScriptSettingEditResponseJSON struct {
	Bindings           apijson.Field
	CompatibilityDate  apijson.Field
	CompatibilityFlags apijson.Field
	Logpush            apijson.Field
	Migrations         apijson.Field
	Placement          apijson.Field
	Tags               apijson.Field
	TailConsumers      apijson.Field
	UsageModel         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseJSON) RawJSON() string {
	return r.raw
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersKVNamespaceBinding],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersServiceBinding],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersDoBinding],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersR2Binding],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersQueueBinding],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersD1Binding],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBinding]
// or
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseBindingsWorkersMTLSCERTBinding].
type DispatchNamespaceScriptSettingEditResponseBinding interface {
	implementsWorkersForPlatformsDispatchNamespaceScriptSettingEditResponseBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DispatchNamespaceScriptSettingEditResponseBinding)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersKVNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersServiceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersDoBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersR2Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersQueueBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersD1Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseBindingsWorkersMTLSCERTBinding{}),
		},
	)
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersKVNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersKVNamespaceBindingType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersKVNamespaceBindingJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersKVNamespaceBindingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersKVNamespaceBinding]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersKVNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersKVNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersKVNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersKVNamespaceBinding) implementsWorkersForPlatformsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersKVNamespaceBindingType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersKVNamespaceBindingTypeKVNamespace DispatchNamespaceScriptSettingEditResponseBindingsWorkersKVNamespaceBindingType = "kv_namespace"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersKVNamespaceBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersKVNamespaceBindingTypeKVNamespace:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersServiceBindingType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersServiceBindingJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersServiceBindingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersServiceBinding]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersServiceBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersServiceBinding) implementsWorkersForPlatformsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersServiceBindingType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersServiceBindingTypeService DispatchNamespaceScriptSettingEditResponseBindingsWorkersServiceBindingType = "service"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersServiceBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersServiceBindingTypeService:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                                 `json:"script_name"`
	JSON       dispatchNamespaceScriptSettingEditResponseBindingsWorkersDoBindingJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersDoBindingJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersDoBinding]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersDoBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersDoBinding) implementsWorkersForPlatformsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersDoBindingType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersDoBindingTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditResponseBindingsWorkersDoBindingType = "durable_object_namespace"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersDoBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersDoBindingTypeDurableObjectNamespace:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersR2BindingType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersR2BindingJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersR2BindingJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersR2Binding]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersR2BindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersR2Binding) implementsWorkersForPlatformsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersR2BindingType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersR2BindingTypeR2Bucket DispatchNamespaceScriptSettingEditResponseBindingsWorkersR2BindingType = "r2_bucket"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersR2BindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersR2BindingTypeR2Bucket:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersQueueBindingType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersQueueBindingJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersQueueBindingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersQueueBinding]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersQueueBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersQueueBinding) implementsWorkersForPlatformsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersQueueBindingType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersQueueBindingTypeQueue DispatchNamespaceScriptSettingEditResponseBindingsWorkersQueueBindingType = "queue"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersQueueBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersQueueBindingTypeQueue:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersD1BindingType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingEditResponseBindingsWorkersD1BindingJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersD1BindingJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersD1Binding]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersD1BindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersD1Binding) implementsWorkersForPlatformsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersD1BindingType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersD1BindingTypeD1 DispatchNamespaceScriptSettingEditResponseBindingsWorkersD1BindingType = "d1"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersD1BindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersD1BindingTypeD1:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound DispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     dispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON     `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBinding]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBinding) implementsWorkersForPlatformsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace DispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker DispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   dispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON) RawJSON() string {
	return r.raw
}

// Outbound worker
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                              `json:"service"`
	JSON    dispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSettingEditResponseBindingsWorkersMTLSCERTBinding struct {
	Certificate interface{} `json:"certificate,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptSettingEditResponseBindingsWorkersMTLSCERTBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                                       `json:"certificate_id"`
	JSON          dispatchNamespaceScriptSettingEditResponseBindingsWorkersMtlscertBindingJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseBindingsWorkersMtlscertBindingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseBindingsWorkersMTLSCERTBinding]
type dispatchNamespaceScriptSettingEditResponseBindingsWorkersMtlscertBindingJSON struct {
	Certificate   apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseBindingsWorkersMTLSCERTBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseBindingsWorkersMtlscertBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersMTLSCERTBinding) implementsWorkersForPlatformsDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptSettingEditResponseBindingsWorkersMTLSCERTBindingType string

const (
	DispatchNamespaceScriptSettingEditResponseBindingsWorkersMTLSCERTBindingTypeMTLSCertificate DispatchNamespaceScriptSettingEditResponseBindingsWorkersMTLSCERTBindingType = "mtls_certificate"
)

func (r DispatchNamespaceScriptSettingEditResponseBindingsWorkersMTLSCERTBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseBindingsWorkersMTLSCERTBindingTypeMTLSCertificate:
		return true
	}
	return false
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrations]
// or
// [workers_for_platforms.DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrations].
type DispatchNamespaceScriptSettingEditResponseMigrations interface {
	implementsWorkersForPlatformsDispatchNamespaceScriptSettingEditResponseMigrations()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DispatchNamespaceScriptSettingEditResponseMigrations)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrations{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrations{}),
		},
	)
}

// A single set of migrations to apply.
type DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrations struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               dispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON               `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrations]
type dispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrations) implementsWorkersForPlatformsDispatchNamespaceScriptSettingEditResponseMigrations() {
}

type DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From string                                                                                          `json:"from"`
	To   string                                                                                          `json:"to"`
	JSON dispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass]
type dispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       string                                                                                              `json:"from"`
	FromScript string                                                                                              `json:"from_script"`
	To         string                                                                                              `json:"to"`
	JSON       dispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass]
type dispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStep `json:"steps"`
	JSON  dispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsJSON   `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrations]
type dispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrations) implementsWorkersForPlatformsDispatchNamespaceScriptSettingEditResponseMigrations() {
}

type DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               dispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON                `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStep]
type dispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From string                                                                                            `json:"from"`
	To   string                                                                                            `json:"to"`
	JSON dispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass]
type dispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                                                `json:"from"`
	FromScript string                                                                                                `json:"from_script"`
	To         string                                                                                                `json:"to"`
	JSON       dispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass]
type dispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSettingEditResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode DispatchNamespaceScriptSettingEditResponsePlacementMode `json:"mode"`
	JSON dispatchNamespaceScriptSettingEditResponsePlacementJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponsePlacementJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptSettingEditResponsePlacement]
type dispatchNamespaceScriptSettingEditResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponsePlacementJSON) RawJSON() string {
	return r.raw
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type DispatchNamespaceScriptSettingEditResponsePlacementMode string

const (
	DispatchNamespaceScriptSettingEditResponsePlacementModeSmart DispatchNamespaceScriptSettingEditResponsePlacementMode = "smart"
)

func (r DispatchNamespaceScriptSettingEditResponsePlacementMode) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponsePlacementModeSmart:
		return true
	}
	return false
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
	// List of bindings attached to this Worker
	Bindings []DispatchNamespaceScriptSettingGetResponseBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations DispatchNamespaceScriptSettingGetResponseMigrations `json:"migrations"`
	Placement  DispatchNamespaceScriptSettingGetResponsePlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []DispatchNamespaceScriptSettingGetResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                        `json:"usage_model"`
	JSON       dispatchNamespaceScriptSettingGetResponseJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseJSON contains the JSON metadata for the
// struct [DispatchNamespaceScriptSettingGetResponse]
type dispatchNamespaceScriptSettingGetResponseJSON struct {
	Bindings           apijson.Field
	CompatibilityDate  apijson.Field
	CompatibilityFlags apijson.Field
	Logpush            apijson.Field
	Migrations         apijson.Field
	Placement          apijson.Field
	Tags               apijson.Field
	TailConsumers      apijson.Field
	UsageModel         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseJSON) RawJSON() string {
	return r.raw
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersKVNamespaceBinding],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersServiceBinding],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersDoBinding],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersR2Binding],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersQueueBinding],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersD1Binding],
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBinding]
// or
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseBindingsWorkersMTLSCERTBinding].
type DispatchNamespaceScriptSettingGetResponseBinding interface {
	implementsWorkersForPlatformsDispatchNamespaceScriptSettingGetResponseBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DispatchNamespaceScriptSettingGetResponseBinding)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersKVNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersServiceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersDoBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersR2Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersQueueBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersD1Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseBindingsWorkersMTLSCERTBinding{}),
		},
	)
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersKVNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersKVNamespaceBindingType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersKVNamespaceBindingJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersKVNamespaceBindingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersKVNamespaceBinding]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersKVNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersKVNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersKVNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersKVNamespaceBinding) implementsWorkersForPlatformsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersKVNamespaceBindingType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersKVNamespaceBindingTypeKVNamespace DispatchNamespaceScriptSettingGetResponseBindingsWorkersKVNamespaceBindingType = "kv_namespace"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersKVNamespaceBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersKVNamespaceBindingTypeKVNamespace:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersServiceBindingType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersServiceBindingJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersServiceBindingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersServiceBinding]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersServiceBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersServiceBinding) implementsWorkersForPlatformsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersServiceBindingType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersServiceBindingTypeService DispatchNamespaceScriptSettingGetResponseBindingsWorkersServiceBindingType = "service"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersServiceBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersServiceBindingTypeService:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                                `json:"script_name"`
	JSON       dispatchNamespaceScriptSettingGetResponseBindingsWorkersDoBindingJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersDoBindingJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersDoBinding]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersDoBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersDoBinding) implementsWorkersForPlatformsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersDoBindingType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersDoBindingTypeDurableObjectNamespace DispatchNamespaceScriptSettingGetResponseBindingsWorkersDoBindingType = "durable_object_namespace"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersDoBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersDoBindingTypeDurableObjectNamespace:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersR2BindingType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersR2BindingJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersR2BindingJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersR2Binding]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersR2BindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersR2Binding) implementsWorkersForPlatformsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersR2BindingType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersR2BindingTypeR2Bucket DispatchNamespaceScriptSettingGetResponseBindingsWorkersR2BindingType = "r2_bucket"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersR2BindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersR2BindingTypeR2Bucket:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersQueueBindingType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersQueueBindingJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersQueueBindingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersQueueBinding]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersQueueBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersQueueBinding) implementsWorkersForPlatformsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersQueueBindingType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersQueueBindingTypeQueue DispatchNamespaceScriptSettingGetResponseBindingsWorkersQueueBindingType = "queue"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersQueueBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersQueueBindingTypeQueue:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersD1BindingType `json:"type,required"`
	JSON dispatchNamespaceScriptSettingGetResponseBindingsWorkersD1BindingJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersD1BindingJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersD1Binding]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersD1BindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersD1Binding) implementsWorkersForPlatformsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersD1BindingType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersD1BindingTypeD1 DispatchNamespaceScriptSettingGetResponseBindingsWorkersD1BindingType = "d1"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersD1BindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersD1BindingTypeD1:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound DispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     dispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON     `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBinding]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBinding) implementsWorkersForPlatformsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace DispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker DispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   dispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON) RawJSON() string {
	return r.raw
}

// Outbound worker
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                             `json:"service"`
	JSON    dispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSettingGetResponseBindingsWorkersMTLSCERTBinding struct {
	Certificate interface{} `json:"certificate,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptSettingGetResponseBindingsWorkersMTLSCERTBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                                      `json:"certificate_id"`
	JSON          dispatchNamespaceScriptSettingGetResponseBindingsWorkersMtlscertBindingJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseBindingsWorkersMtlscertBindingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseBindingsWorkersMTLSCERTBinding]
type dispatchNamespaceScriptSettingGetResponseBindingsWorkersMtlscertBindingJSON struct {
	Certificate   apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseBindingsWorkersMTLSCERTBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseBindingsWorkersMtlscertBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersMTLSCERTBinding) implementsWorkersForPlatformsDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptSettingGetResponseBindingsWorkersMTLSCERTBindingType string

const (
	DispatchNamespaceScriptSettingGetResponseBindingsWorkersMTLSCERTBindingTypeMTLSCertificate DispatchNamespaceScriptSettingGetResponseBindingsWorkersMTLSCERTBindingType = "mtls_certificate"
)

func (r DispatchNamespaceScriptSettingGetResponseBindingsWorkersMTLSCERTBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseBindingsWorkersMTLSCERTBindingTypeMTLSCertificate:
		return true
	}
	return false
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrations]
// or
// [workers_for_platforms.DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrations].
type DispatchNamespaceScriptSettingGetResponseMigrations interface {
	implementsWorkersForPlatformsDispatchNamespaceScriptSettingGetResponseMigrations()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DispatchNamespaceScriptSettingGetResponseMigrations)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrations{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrations{}),
		},
	)
}

// A single set of migrations to apply.
type DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrations struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               dispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON               `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrations]
type dispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrations) implementsWorkersForPlatformsDispatchNamespaceScriptSettingGetResponseMigrations() {
}

type DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From string                                                                                         `json:"from"`
	To   string                                                                                         `json:"to"`
	JSON dispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass]
type dispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       string                                                                                             `json:"from"`
	FromScript string                                                                                             `json:"from_script"`
	To         string                                                                                             `json:"to"`
	JSON       dispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass]
type dispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStep `json:"steps"`
	JSON  dispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsJSON   `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrations]
type dispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrations) implementsWorkersForPlatformsDispatchNamespaceScriptSettingGetResponseMigrations() {
}

type DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               dispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON                `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStep]
type dispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From string                                                                                           `json:"from"`
	To   string                                                                                           `json:"to"`
	JSON dispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass]
type dispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                                               `json:"from"`
	FromScript string                                                                                               `json:"from_script"`
	To         string                                                                                               `json:"to"`
	JSON       dispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass]
type dispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSettingGetResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode DispatchNamespaceScriptSettingGetResponsePlacementMode `json:"mode"`
	JSON dispatchNamespaceScriptSettingGetResponsePlacementJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponsePlacementJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptSettingGetResponsePlacement]
type dispatchNamespaceScriptSettingGetResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponsePlacementJSON) RawJSON() string {
	return r.raw
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type DispatchNamespaceScriptSettingGetResponsePlacementMode string

const (
	DispatchNamespaceScriptSettingGetResponsePlacementModeSmart DispatchNamespaceScriptSettingGetResponsePlacementMode = "smart"
)

func (r DispatchNamespaceScriptSettingGetResponsePlacementMode) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponsePlacementModeSmart:
		return true
	}
	return false
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
	AccountID param.Field[string]                                            `path:"account_id,required"`
	Errors    param.Field[[]DispatchNamespaceScriptSettingEditParamsError]   `json:"errors,required"`
	Messages  param.Field[[]DispatchNamespaceScriptSettingEditParamsMessage] `json:"messages,required"`
	Result    param.Field[DispatchNamespaceScriptSettingEditParamsResult]    `json:"result,required"`
	// Whether the API call was successful
	Success param.Field[DispatchNamespaceScriptSettingEditParamsSuccess] `json:"success,required"`
}

func (r DispatchNamespaceScriptSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceScriptSettingEditParamsError struct {
	Code    param.Field[int64]  `json:"code,required"`
	Message param.Field[string] `json:"message,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsError) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceScriptSettingEditParamsMessage struct {
	Code    param.Field[int64]  `json:"code,required"`
	Message param.Field[string] `json:"message,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceScriptSettingEditParamsResult struct {
	// List of bindings attached to this Worker
	Bindings param.Field[[]DispatchNamespaceScriptSettingEditParamsResultBinding] `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations param.Field[DispatchNamespaceScriptSettingEditParamsResultMigrations] `json:"migrations"`
	Placement  param.Field[DispatchNamespaceScriptSettingEditParamsResultPlacement]  `json:"placement"`
	// Tags to help you manage your Workers
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]DispatchNamespaceScriptSettingEditParamsResultTailConsumer] `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel param.Field[string] `json:"usage_model"`
}

func (r DispatchNamespaceScriptSettingEditParamsResult) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A binding to allow the Worker to communicate with resources
//
// Satisfied by
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersKVNamespaceBinding],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersServiceBinding],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDoBinding],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersR2Binding],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersQueueBinding],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersD1Binding],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDispatchNamespaceBinding],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersMTLSCERTBinding].
type DispatchNamespaceScriptSettingEditParamsResultBinding interface {
	implementsWorkersForPlatformsDispatchNamespaceScriptSettingEditParamsResultBinding()
}

type DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersKVNamespaceBinding struct {
	// The class of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersKVNamespaceBindingType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersKVNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersKVNamespaceBinding) implementsWorkersForPlatformsDispatchNamespaceScriptSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersKVNamespaceBindingType string

const (
	DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersKVNamespaceBindingTypeKVNamespace DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersKVNamespaceBindingType = "kv_namespace"
)

func (r DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersKVNamespaceBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersKVNamespaceBindingTypeKVNamespace:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// Name of Worker to bind to
	Service param.Field[string] `json:"service,required"`
	// The class of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersServiceBindingType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersServiceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersServiceBinding) implementsWorkersForPlatformsDispatchNamespaceScriptSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersServiceBindingType string

const (
	DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersServiceBindingTypeService DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersServiceBindingType = "service"
)

func (r DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersServiceBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersServiceBindingTypeService:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName param.Field[string] `json:"class_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDoBindingType] `json:"type,required"`
	// The environment of the script_name to bind to
	Environment param.Field[string] `json:"environment"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName param.Field[string] `json:"script_name"`
}

func (r DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDoBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDoBinding) implementsWorkersForPlatformsDispatchNamespaceScriptSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDoBindingType string

const (
	DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDoBindingTypeDurableObjectNamespace DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDoBindingType = "durable_object_namespace"
)

func (r DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDoBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDoBindingTypeDurableObjectNamespace:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName param.Field[string] `json:"bucket_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersR2BindingType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersR2Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersR2Binding) implementsWorkersForPlatformsDispatchNamespaceScriptSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersR2BindingType string

const (
	DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersR2BindingTypeR2Bucket DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersR2BindingType = "r2_bucket"
)

func (r DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersR2BindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersR2BindingTypeR2Bucket:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersQueueBinding struct {
	// Name of the Queue to bind to
	QueueName param.Field[string] `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersQueueBindingType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersQueueBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersQueueBinding) implementsWorkersForPlatformsDispatchNamespaceScriptSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersQueueBindingType string

const (
	DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersQueueBindingTypeQueue DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersQueueBindingType = "queue"
)

func (r DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersQueueBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersQueueBindingTypeQueue:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID param.Field[string] `json:"id,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name param.Field[string] `json:"name,required"`
	// The class of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersD1BindingType] `json:"type,required"`
}

func (r DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersD1Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersD1Binding) implementsWorkersForPlatformsDispatchNamespaceScriptSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersD1BindingType string

const (
	DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersD1BindingTypeD1 DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersD1BindingType = "d1"
)

func (r DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersD1BindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersD1BindingTypeD1:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDispatchNamespaceBinding struct {
	// Namespace to bind to
	Namespace param.Field[string] `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingType] `json:"type,required"`
	// Outbound worker
	Outbound param.Field[DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutbound] `json:"outbound"`
}

func (r DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDispatchNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDispatchNamespaceBinding) implementsWorkersForPlatformsDispatchNamespaceScriptSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingType string

const (
	DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

func (r DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker
type DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params param.Field[[]string] `json:"params"`
	// Outbound worker
	Worker param.Field[DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker] `json:"worker"`
}

func (r DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutbound) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker
type DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker
	Service param.Field[string] `json:"service"`
}

func (r DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersMTLSCERTBinding struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
	// The class of resource that the binding provides.
	Type param.Field[DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersMTLSCERTBindingType] `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID param.Field[string] `json:"certificate_id"`
}

func (r DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersMTLSCERTBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersMTLSCERTBinding) implementsWorkersForPlatformsDispatchNamespaceScriptSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersMTLSCERTBindingType string

const (
	DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersMTLSCERTBindingTypeMTLSCertificate DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersMTLSCERTBindingType = "mtls_certificate"
)

func (r DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersMTLSCERTBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsResultBindingsWorkersMTLSCERTBindingTypeMTLSCertificate:
		return true
	}
	return false
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSingleStepMigrations],
// [workers_for_platforms.DispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSteppedMigrations].
type DispatchNamespaceScriptSettingEditParamsResultMigrations interface {
	implementsWorkersForPlatformsDispatchNamespaceScriptSettingEditParamsResultMigrations()
}

// A single set of migrations to apply.
type DispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSingleStepMigrations struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]DispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSingleStepMigrationsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]DispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSingleStepMigrationsTransferredClass] `json:"transferred_classes"`
}

func (r DispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSingleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSingleStepMigrations) implementsWorkersForPlatformsDispatchNamespaceScriptSettingEditParamsResultMigrations() {
}

type DispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r DispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSingleStepMigrationsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r DispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSingleStepMigrationsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]DispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSteppedMigrationsStep] `json:"steps"`
}

func (r DispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSteppedMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSteppedMigrations) implementsWorkersForPlatformsDispatchNamespaceScriptSettingEditParamsResultMigrations() {
}

type DispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]DispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]DispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass] `json:"transferred_classes"`
}

func (r DispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSteppedMigrationsStep) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r DispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r DispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceScriptSettingEditParamsResultPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode param.Field[DispatchNamespaceScriptSettingEditParamsResultPlacementMode] `json:"mode"`
}

func (r DispatchNamespaceScriptSettingEditParamsResultPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type DispatchNamespaceScriptSettingEditParamsResultPlacementMode string

const (
	DispatchNamespaceScriptSettingEditParamsResultPlacementModeSmart DispatchNamespaceScriptSettingEditParamsResultPlacementMode = "smart"
)

func (r DispatchNamespaceScriptSettingEditParamsResultPlacementMode) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditParamsResultPlacementModeSmart:
		return true
	}
	return false
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
	Errors   []DispatchNamespaceScriptSettingEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DispatchNamespaceScriptSettingEditResponseEnvelopeMessages `json:"messages,required"`
	Result   DispatchNamespaceScriptSettingEditResponse                   `json:"result,required"`
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

type DispatchNamespaceScriptSettingEditResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    dispatchNamespaceScriptSettingEditResponseEnvelopeErrorsJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseEnvelopeErrors]
type dispatchNamespaceScriptSettingEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSettingEditResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    dispatchNamespaceScriptSettingEditResponseEnvelopeMessagesJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptSettingEditResponseEnvelopeMessages]
type dispatchNamespaceScriptSettingEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseEnvelopeMessagesJSON) RawJSON() string {
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
	Errors   []DispatchNamespaceScriptSettingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DispatchNamespaceScriptSettingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DispatchNamespaceScriptSettingGetResponse                   `json:"result,required"`
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

type DispatchNamespaceScriptSettingGetResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    dispatchNamespaceScriptSettingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseEnvelopeErrors]
type dispatchNamespaceScriptSettingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptSettingGetResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    dispatchNamespaceScriptSettingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptSettingGetResponseEnvelopeMessages]
type dispatchNamespaceScriptSettingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseEnvelopeMessagesJSON) RawJSON() string {
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
