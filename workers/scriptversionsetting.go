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

// ScriptVersionSettingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewScriptVersionSettingService]
// method instead.
type ScriptVersionSettingService struct {
	Options []option.RequestOption
}

// NewScriptVersionSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewScriptVersionSettingService(opts ...option.RequestOption) (r *ScriptVersionSettingService) {
	r = &ScriptVersionSettingService{}
	r.Options = opts
	return
}

// Patch metadata or config, such as bindings or usage model
func (r *ScriptVersionSettingService) Edit(ctx context.Context, scriptName string, params ScriptVersionSettingEditParams, opts ...option.RequestOption) (res *ScriptVersionSettingEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ScriptVersionSettingEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/settings", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get metadata and config, such as bindings or usage model
func (r *ScriptVersionSettingService) Get(ctx context.Context, scriptName string, query ScriptVersionSettingGetParams, opts ...option.RequestOption) (res *ScriptVersionSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ScriptVersionSettingGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/settings", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ScriptVersionSettingEditResponse struct {
	// List of bindings attached to this Worker
	Bindings []ScriptVersionSettingEditResponseBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations ScriptVersionSettingEditResponseMigrations `json:"migrations"`
	Placement  ScriptVersionSettingEditResponsePlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []ScriptVersionSettingEditResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                               `json:"usage_model"`
	JSON       scriptVersionSettingEditResponseJSON `json:"-"`
}

// scriptVersionSettingEditResponseJSON contains the JSON metadata for the struct
// [ScriptVersionSettingEditResponse]
type scriptVersionSettingEditResponseJSON struct {
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

func (r *ScriptVersionSettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingEditResponseJSON) RawJSON() string {
	return r.raw
}

// A binding to allow the Worker to communicate with resources
type ScriptVersionSettingEditResponseBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The class of resource that the binding provides.
	Type ScriptVersionSettingEditResponseBindingsType `json:"type,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Name of Worker to bind to
	Service string `json:"service"`
	// The exported class name of the Durable Object
	ClassName string `json:"class_name"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string `json:"script_name"`
	// R2 bucket to bind to
	BucketName string `json:"bucket_name"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding"`
	// ID of the D1 database to bind to
	ID string `json:"id"`
	// Namespace to bind to
	Namespace string      `json:"namespace"`
	Outbound  interface{} `json:"outbound,required"`
	// ID of the certificate to bind to
	CertificateID string                                      `json:"certificate_id"`
	Certificate   interface{}                                 `json:"certificate,required"`
	JSON          scriptVersionSettingEditResponseBindingJSON `json:"-"`
	union         ScriptVersionSettingEditResponseBindingsUnion
}

// scriptVersionSettingEditResponseBindingJSON contains the JSON metadata for the
// struct [ScriptVersionSettingEditResponseBinding]
type scriptVersionSettingEditResponseBindingJSON struct {
	Name          apijson.Field
	NamespaceID   apijson.Field
	Type          apijson.Field
	Environment   apijson.Field
	Service       apijson.Field
	ClassName     apijson.Field
	ScriptName    apijson.Field
	BucketName    apijson.Field
	QueueName     apijson.Field
	Binding       apijson.Field
	ID            apijson.Field
	Namespace     apijson.Field
	Outbound      apijson.Field
	CertificateID apijson.Field
	Certificate   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r scriptVersionSettingEditResponseBindingJSON) RawJSON() string {
	return r.raw
}

func (r *ScriptVersionSettingEditResponseBinding) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r ScriptVersionSettingEditResponseBinding) AsUnion() ScriptVersionSettingEditResponseBindingsUnion {
	return r.union
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [workers.ScriptVersionSettingEditResponseBindingsWorkersKVNamespaceBinding],
// [workers.ScriptVersionSettingEditResponseBindingsWorkersServiceBinding],
// [workers.ScriptVersionSettingEditResponseBindingsWorkersDoBinding],
// [workers.ScriptVersionSettingEditResponseBindingsWorkersR2Binding],
// [workers.ScriptVersionSettingEditResponseBindingsWorkersQueueBinding],
// [workers.ScriptVersionSettingEditResponseBindingsWorkersD1Binding],
// [workers.ScriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBinding]
// or [workers.ScriptVersionSettingEditResponseBindingsWorkersMTLSCERTBinding].
type ScriptVersionSettingEditResponseBindingsUnion interface {
	implementsWorkersScriptVersionSettingEditResponseBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScriptVersionSettingEditResponseBindingsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptVersionSettingEditResponseBindingsWorkersKVNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptVersionSettingEditResponseBindingsWorkersServiceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptVersionSettingEditResponseBindingsWorkersDoBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptVersionSettingEditResponseBindingsWorkersR2Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptVersionSettingEditResponseBindingsWorkersQueueBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptVersionSettingEditResponseBindingsWorkersD1Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptVersionSettingEditResponseBindingsWorkersMTLSCERTBinding{}),
		},
	)
}

type ScriptVersionSettingEditResponseBindingsWorkersKVNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type ScriptVersionSettingEditResponseBindingsWorkersKVNamespaceBindingType `json:"type,required"`
	JSON scriptVersionSettingEditResponseBindingsWorkersKVNamespaceBindingJSON `json:"-"`
}

// scriptVersionSettingEditResponseBindingsWorkersKVNamespaceBindingJSON contains
// the JSON metadata for the struct
// [ScriptVersionSettingEditResponseBindingsWorkersKVNamespaceBinding]
type scriptVersionSettingEditResponseBindingsWorkersKVNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingEditResponseBindingsWorkersKVNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingEditResponseBindingsWorkersKVNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionSettingEditResponseBindingsWorkersKVNamespaceBinding) implementsWorkersScriptVersionSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptVersionSettingEditResponseBindingsWorkersKVNamespaceBindingType string

const (
	ScriptVersionSettingEditResponseBindingsWorkersKVNamespaceBindingTypeKVNamespace ScriptVersionSettingEditResponseBindingsWorkersKVNamespaceBindingType = "kv_namespace"
)

func (r ScriptVersionSettingEditResponseBindingsWorkersKVNamespaceBindingType) IsKnown() bool {
	switch r {
	case ScriptVersionSettingEditResponseBindingsWorkersKVNamespaceBindingTypeKVNamespace:
		return true
	}
	return false
}

type ScriptVersionSettingEditResponseBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type ScriptVersionSettingEditResponseBindingsWorkersServiceBindingType `json:"type,required"`
	JSON scriptVersionSettingEditResponseBindingsWorkersServiceBindingJSON `json:"-"`
}

// scriptVersionSettingEditResponseBindingsWorkersServiceBindingJSON contains the
// JSON metadata for the struct
// [ScriptVersionSettingEditResponseBindingsWorkersServiceBinding]
type scriptVersionSettingEditResponseBindingsWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingEditResponseBindingsWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingEditResponseBindingsWorkersServiceBindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionSettingEditResponseBindingsWorkersServiceBinding) implementsWorkersScriptVersionSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptVersionSettingEditResponseBindingsWorkersServiceBindingType string

const (
	ScriptVersionSettingEditResponseBindingsWorkersServiceBindingTypeService ScriptVersionSettingEditResponseBindingsWorkersServiceBindingType = "service"
)

func (r ScriptVersionSettingEditResponseBindingsWorkersServiceBindingType) IsKnown() bool {
	switch r {
	case ScriptVersionSettingEditResponseBindingsWorkersServiceBindingTypeService:
		return true
	}
	return false
}

type ScriptVersionSettingEditResponseBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type ScriptVersionSettingEditResponseBindingsWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                       `json:"script_name"`
	JSON       scriptVersionSettingEditResponseBindingsWorkersDoBindingJSON `json:"-"`
}

// scriptVersionSettingEditResponseBindingsWorkersDoBindingJSON contains the JSON
// metadata for the struct
// [ScriptVersionSettingEditResponseBindingsWorkersDoBinding]
type scriptVersionSettingEditResponseBindingsWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingEditResponseBindingsWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingEditResponseBindingsWorkersDoBindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionSettingEditResponseBindingsWorkersDoBinding) implementsWorkersScriptVersionSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptVersionSettingEditResponseBindingsWorkersDoBindingType string

const (
	ScriptVersionSettingEditResponseBindingsWorkersDoBindingTypeDurableObjectNamespace ScriptVersionSettingEditResponseBindingsWorkersDoBindingType = "durable_object_namespace"
)

func (r ScriptVersionSettingEditResponseBindingsWorkersDoBindingType) IsKnown() bool {
	switch r {
	case ScriptVersionSettingEditResponseBindingsWorkersDoBindingTypeDurableObjectNamespace:
		return true
	}
	return false
}

type ScriptVersionSettingEditResponseBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type ScriptVersionSettingEditResponseBindingsWorkersR2BindingType `json:"type,required"`
	JSON scriptVersionSettingEditResponseBindingsWorkersR2BindingJSON `json:"-"`
}

// scriptVersionSettingEditResponseBindingsWorkersR2BindingJSON contains the JSON
// metadata for the struct
// [ScriptVersionSettingEditResponseBindingsWorkersR2Binding]
type scriptVersionSettingEditResponseBindingsWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingEditResponseBindingsWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingEditResponseBindingsWorkersR2BindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionSettingEditResponseBindingsWorkersR2Binding) implementsWorkersScriptVersionSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptVersionSettingEditResponseBindingsWorkersR2BindingType string

const (
	ScriptVersionSettingEditResponseBindingsWorkersR2BindingTypeR2Bucket ScriptVersionSettingEditResponseBindingsWorkersR2BindingType = "r2_bucket"
)

func (r ScriptVersionSettingEditResponseBindingsWorkersR2BindingType) IsKnown() bool {
	switch r {
	case ScriptVersionSettingEditResponseBindingsWorkersR2BindingTypeR2Bucket:
		return true
	}
	return false
}

type ScriptVersionSettingEditResponseBindingsWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type ScriptVersionSettingEditResponseBindingsWorkersQueueBindingType `json:"type,required"`
	JSON scriptVersionSettingEditResponseBindingsWorkersQueueBindingJSON `json:"-"`
}

// scriptVersionSettingEditResponseBindingsWorkersQueueBindingJSON contains the
// JSON metadata for the struct
// [ScriptVersionSettingEditResponseBindingsWorkersQueueBinding]
type scriptVersionSettingEditResponseBindingsWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingEditResponseBindingsWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingEditResponseBindingsWorkersQueueBindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionSettingEditResponseBindingsWorkersQueueBinding) implementsWorkersScriptVersionSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptVersionSettingEditResponseBindingsWorkersQueueBindingType string

const (
	ScriptVersionSettingEditResponseBindingsWorkersQueueBindingTypeQueue ScriptVersionSettingEditResponseBindingsWorkersQueueBindingType = "queue"
)

func (r ScriptVersionSettingEditResponseBindingsWorkersQueueBindingType) IsKnown() bool {
	switch r {
	case ScriptVersionSettingEditResponseBindingsWorkersQueueBindingTypeQueue:
		return true
	}
	return false
}

type ScriptVersionSettingEditResponseBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type ScriptVersionSettingEditResponseBindingsWorkersD1BindingType `json:"type,required"`
	JSON scriptVersionSettingEditResponseBindingsWorkersD1BindingJSON `json:"-"`
}

// scriptVersionSettingEditResponseBindingsWorkersD1BindingJSON contains the JSON
// metadata for the struct
// [ScriptVersionSettingEditResponseBindingsWorkersD1Binding]
type scriptVersionSettingEditResponseBindingsWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingEditResponseBindingsWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingEditResponseBindingsWorkersD1BindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionSettingEditResponseBindingsWorkersD1Binding) implementsWorkersScriptVersionSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptVersionSettingEditResponseBindingsWorkersD1BindingType string

const (
	ScriptVersionSettingEditResponseBindingsWorkersD1BindingTypeD1 ScriptVersionSettingEditResponseBindingsWorkersD1BindingType = "d1"
)

func (r ScriptVersionSettingEditResponseBindingsWorkersD1BindingType) IsKnown() bool {
	switch r {
	case ScriptVersionSettingEditResponseBindingsWorkersD1BindingTypeD1:
		return true
	}
	return false
}

type ScriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type ScriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound ScriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     scriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON     `json:"-"`
}

// scriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [ScriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBinding]
type scriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBinding) implementsWorkersScriptVersionSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBindingType string

const (
	ScriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace ScriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

func (r ScriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBindingType) IsKnown() bool {
	switch r {
	case ScriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker
type ScriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker ScriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   scriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// scriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [ScriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound]
type scriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON) RawJSON() string {
	return r.raw
}

// Outbound worker
type ScriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                    `json:"service"`
	JSON    scriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// scriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [ScriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker]
type scriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionSettingEditResponseBindingsWorkersMTLSCERTBinding struct {
	Certificate interface{} `json:"certificate,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type ScriptVersionSettingEditResponseBindingsWorkersMTLSCERTBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                             `json:"certificate_id"`
	JSON          scriptVersionSettingEditResponseBindingsWorkersMtlscertBindingJSON `json:"-"`
}

// scriptVersionSettingEditResponseBindingsWorkersMtlscertBindingJSON contains the
// JSON metadata for the struct
// [ScriptVersionSettingEditResponseBindingsWorkersMTLSCERTBinding]
type scriptVersionSettingEditResponseBindingsWorkersMtlscertBindingJSON struct {
	Certificate   apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScriptVersionSettingEditResponseBindingsWorkersMTLSCERTBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingEditResponseBindingsWorkersMtlscertBindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionSettingEditResponseBindingsWorkersMTLSCERTBinding) implementsWorkersScriptVersionSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptVersionSettingEditResponseBindingsWorkersMTLSCERTBindingType string

const (
	ScriptVersionSettingEditResponseBindingsWorkersMTLSCERTBindingTypeMTLSCertificate ScriptVersionSettingEditResponseBindingsWorkersMTLSCERTBindingType = "mtls_certificate"
)

func (r ScriptVersionSettingEditResponseBindingsWorkersMTLSCERTBindingType) IsKnown() bool {
	switch r {
	case ScriptVersionSettingEditResponseBindingsWorkersMTLSCERTBindingTypeMTLSCertificate:
		return true
	}
	return false
}

// The class of resource that the binding provides.
type ScriptVersionSettingEditResponseBindingsType string

const (
	ScriptVersionSettingEditResponseBindingsTypeKVNamespace            ScriptVersionSettingEditResponseBindingsType = "kv_namespace"
	ScriptVersionSettingEditResponseBindingsTypeService                ScriptVersionSettingEditResponseBindingsType = "service"
	ScriptVersionSettingEditResponseBindingsTypeDurableObjectNamespace ScriptVersionSettingEditResponseBindingsType = "durable_object_namespace"
	ScriptVersionSettingEditResponseBindingsTypeR2Bucket               ScriptVersionSettingEditResponseBindingsType = "r2_bucket"
	ScriptVersionSettingEditResponseBindingsTypeQueue                  ScriptVersionSettingEditResponseBindingsType = "queue"
	ScriptVersionSettingEditResponseBindingsTypeD1                     ScriptVersionSettingEditResponseBindingsType = "d1"
	ScriptVersionSettingEditResponseBindingsTypeDispatchNamespace      ScriptVersionSettingEditResponseBindingsType = "dispatch_namespace"
	ScriptVersionSettingEditResponseBindingsTypeMTLSCertificate        ScriptVersionSettingEditResponseBindingsType = "mtls_certificate"
)

func (r ScriptVersionSettingEditResponseBindingsType) IsKnown() bool {
	switch r {
	case ScriptVersionSettingEditResponseBindingsTypeKVNamespace, ScriptVersionSettingEditResponseBindingsTypeService, ScriptVersionSettingEditResponseBindingsTypeDurableObjectNamespace, ScriptVersionSettingEditResponseBindingsTypeR2Bucket, ScriptVersionSettingEditResponseBindingsTypeQueue, ScriptVersionSettingEditResponseBindingsTypeD1, ScriptVersionSettingEditResponseBindingsTypeDispatchNamespace, ScriptVersionSettingEditResponseBindingsTypeMTLSCertificate:
		return true
	}
	return false
}

// Migrations to apply for Durable Objects associated with this Worker.
type ScriptVersionSettingEditResponseMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag             string                                         `json:"old_tag"`
	DeletedClasses     interface{}                                    `json:"deleted_classes,required"`
	NewClasses         interface{}                                    `json:"new_classes,required"`
	RenamedClasses     interface{}                                    `json:"renamed_classes,required"`
	TransferredClasses interface{}                                    `json:"transferred_classes,required"`
	Steps              interface{}                                    `json:"steps,required"`
	JSON               scriptVersionSettingEditResponseMigrationsJSON `json:"-"`
	union              ScriptVersionSettingEditResponseMigrationsUnion
}

// scriptVersionSettingEditResponseMigrationsJSON contains the JSON metadata for
// the struct [ScriptVersionSettingEditResponseMigrations]
type scriptVersionSettingEditResponseMigrationsJSON struct {
	NewTag             apijson.Field
	OldTag             apijson.Field
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	Steps              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r scriptVersionSettingEditResponseMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r *ScriptVersionSettingEditResponseMigrations) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r ScriptVersionSettingEditResponseMigrations) AsUnion() ScriptVersionSettingEditResponseMigrationsUnion {
	return r.union
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [workers.ScriptVersionSettingEditResponseMigrationsWorkersSingleStepMigrations]
// or [workers.ScriptVersionSettingEditResponseMigrationsWorkersSteppedMigrations].
type ScriptVersionSettingEditResponseMigrationsUnion interface {
	implementsWorkersScriptVersionSettingEditResponseMigrations()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScriptVersionSettingEditResponseMigrationsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptVersionSettingEditResponseMigrationsWorkersSingleStepMigrations{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptVersionSettingEditResponseMigrationsWorkersSteppedMigrations{}),
		},
	)
}

// A single set of migrations to apply.
type ScriptVersionSettingEditResponseMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses []ScriptVersionSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []ScriptVersionSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               scriptVersionSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON               `json:"-"`
}

// scriptVersionSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON
// contains the JSON metadata for the struct
// [ScriptVersionSettingEditResponseMigrationsWorkersSingleStepMigrations]
type scriptVersionSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ScriptVersionSettingEditResponseMigrationsWorkersSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionSettingEditResponseMigrationsWorkersSingleStepMigrations) implementsWorkersScriptVersionSettingEditResponseMigrations() {
}

type ScriptVersionSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From string                                                                                `json:"from"`
	To   string                                                                                `json:"to"`
	JSON scriptVersionSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// scriptVersionSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [ScriptVersionSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass]
type scriptVersionSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       string                                                                                    `json:"from"`
	FromScript string                                                                                    `json:"from_script"`
	To         string                                                                                    `json:"to"`
	JSON       scriptVersionSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// scriptVersionSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [ScriptVersionSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass]
type scriptVersionSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionSettingEditResponseMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []ScriptVersionSettingEditResponseMigrationsWorkersSteppedMigrationsStep `json:"steps"`
	JSON  scriptVersionSettingEditResponseMigrationsWorkersSteppedMigrationsJSON   `json:"-"`
}

// scriptVersionSettingEditResponseMigrationsWorkersSteppedMigrationsJSON contains
// the JSON metadata for the struct
// [ScriptVersionSettingEditResponseMigrationsWorkersSteppedMigrations]
type scriptVersionSettingEditResponseMigrationsWorkersSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingEditResponseMigrationsWorkersSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingEditResponseMigrationsWorkersSteppedMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionSettingEditResponseMigrationsWorkersSteppedMigrations) implementsWorkersScriptVersionSettingEditResponseMigrations() {
}

type ScriptVersionSettingEditResponseMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []ScriptVersionSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []ScriptVersionSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               scriptVersionSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON                `json:"-"`
}

// scriptVersionSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON
// contains the JSON metadata for the struct
// [ScriptVersionSettingEditResponseMigrationsWorkersSteppedMigrationsStep]
type scriptVersionSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ScriptVersionSettingEditResponseMigrationsWorkersSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From string                                                                                  `json:"from"`
	To   string                                                                                  `json:"to"`
	JSON scriptVersionSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// scriptVersionSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [ScriptVersionSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass]
type scriptVersionSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                                      `json:"from"`
	FromScript string                                                                                      `json:"from_script"`
	To         string                                                                                      `json:"to"`
	JSON       scriptVersionSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// scriptVersionSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [ScriptVersionSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass]
type scriptVersionSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionSettingEditResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode ScriptVersionSettingEditResponsePlacementMode `json:"mode"`
	JSON scriptVersionSettingEditResponsePlacementJSON `json:"-"`
}

// scriptVersionSettingEditResponsePlacementJSON contains the JSON metadata for the
// struct [ScriptVersionSettingEditResponsePlacement]
type scriptVersionSettingEditResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingEditResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingEditResponsePlacementJSON) RawJSON() string {
	return r.raw
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type ScriptVersionSettingEditResponsePlacementMode string

const (
	ScriptVersionSettingEditResponsePlacementModeSmart ScriptVersionSettingEditResponsePlacementMode = "smart"
)

func (r ScriptVersionSettingEditResponsePlacementMode) IsKnown() bool {
	switch r {
	case ScriptVersionSettingEditResponsePlacementModeSmart:
		return true
	}
	return false
}

// A reference to a script that will consume logs from the attached Worker.
type ScriptVersionSettingEditResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                           `json:"namespace"`
	JSON      scriptVersionSettingEditResponseTailConsumerJSON `json:"-"`
}

// scriptVersionSettingEditResponseTailConsumerJSON contains the JSON metadata for
// the struct [ScriptVersionSettingEditResponseTailConsumer]
type scriptVersionSettingEditResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingEditResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingEditResponseTailConsumerJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionSettingGetResponse struct {
	// List of bindings attached to this Worker
	Bindings []ScriptVersionSettingGetResponseBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations ScriptVersionSettingGetResponseMigrations `json:"migrations"`
	Placement  ScriptVersionSettingGetResponsePlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []ScriptVersionSettingGetResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                              `json:"usage_model"`
	JSON       scriptVersionSettingGetResponseJSON `json:"-"`
}

// scriptVersionSettingGetResponseJSON contains the JSON metadata for the struct
// [ScriptVersionSettingGetResponse]
type scriptVersionSettingGetResponseJSON struct {
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

func (r *ScriptVersionSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingGetResponseJSON) RawJSON() string {
	return r.raw
}

// A binding to allow the Worker to communicate with resources
type ScriptVersionSettingGetResponseBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The class of resource that the binding provides.
	Type ScriptVersionSettingGetResponseBindingsType `json:"type,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Name of Worker to bind to
	Service string `json:"service"`
	// The exported class name of the Durable Object
	ClassName string `json:"class_name"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string `json:"script_name"`
	// R2 bucket to bind to
	BucketName string `json:"bucket_name"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding"`
	// ID of the D1 database to bind to
	ID string `json:"id"`
	// Namespace to bind to
	Namespace string      `json:"namespace"`
	Outbound  interface{} `json:"outbound,required"`
	// ID of the certificate to bind to
	CertificateID string                                     `json:"certificate_id"`
	Certificate   interface{}                                `json:"certificate,required"`
	JSON          scriptVersionSettingGetResponseBindingJSON `json:"-"`
	union         ScriptVersionSettingGetResponseBindingsUnion
}

// scriptVersionSettingGetResponseBindingJSON contains the JSON metadata for the
// struct [ScriptVersionSettingGetResponseBinding]
type scriptVersionSettingGetResponseBindingJSON struct {
	Name          apijson.Field
	NamespaceID   apijson.Field
	Type          apijson.Field
	Environment   apijson.Field
	Service       apijson.Field
	ClassName     apijson.Field
	ScriptName    apijson.Field
	BucketName    apijson.Field
	QueueName     apijson.Field
	Binding       apijson.Field
	ID            apijson.Field
	Namespace     apijson.Field
	Outbound      apijson.Field
	CertificateID apijson.Field
	Certificate   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r scriptVersionSettingGetResponseBindingJSON) RawJSON() string {
	return r.raw
}

func (r *ScriptVersionSettingGetResponseBinding) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r ScriptVersionSettingGetResponseBinding) AsUnion() ScriptVersionSettingGetResponseBindingsUnion {
	return r.union
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [workers.ScriptVersionSettingGetResponseBindingsWorkersKVNamespaceBinding],
// [workers.ScriptVersionSettingGetResponseBindingsWorkersServiceBinding],
// [workers.ScriptVersionSettingGetResponseBindingsWorkersDoBinding],
// [workers.ScriptVersionSettingGetResponseBindingsWorkersR2Binding],
// [workers.ScriptVersionSettingGetResponseBindingsWorkersQueueBinding],
// [workers.ScriptVersionSettingGetResponseBindingsWorkersD1Binding],
// [workers.ScriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBinding]
// or [workers.ScriptVersionSettingGetResponseBindingsWorkersMTLSCERTBinding].
type ScriptVersionSettingGetResponseBindingsUnion interface {
	implementsWorkersScriptVersionSettingGetResponseBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScriptVersionSettingGetResponseBindingsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptVersionSettingGetResponseBindingsWorkersKVNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptVersionSettingGetResponseBindingsWorkersServiceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptVersionSettingGetResponseBindingsWorkersDoBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptVersionSettingGetResponseBindingsWorkersR2Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptVersionSettingGetResponseBindingsWorkersQueueBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptVersionSettingGetResponseBindingsWorkersD1Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptVersionSettingGetResponseBindingsWorkersMTLSCERTBinding{}),
		},
	)
}

type ScriptVersionSettingGetResponseBindingsWorkersKVNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type ScriptVersionSettingGetResponseBindingsWorkersKVNamespaceBindingType `json:"type,required"`
	JSON scriptVersionSettingGetResponseBindingsWorkersKVNamespaceBindingJSON `json:"-"`
}

// scriptVersionSettingGetResponseBindingsWorkersKVNamespaceBindingJSON contains
// the JSON metadata for the struct
// [ScriptVersionSettingGetResponseBindingsWorkersKVNamespaceBinding]
type scriptVersionSettingGetResponseBindingsWorkersKVNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingGetResponseBindingsWorkersKVNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingGetResponseBindingsWorkersKVNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionSettingGetResponseBindingsWorkersKVNamespaceBinding) implementsWorkersScriptVersionSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptVersionSettingGetResponseBindingsWorkersKVNamespaceBindingType string

const (
	ScriptVersionSettingGetResponseBindingsWorkersKVNamespaceBindingTypeKVNamespace ScriptVersionSettingGetResponseBindingsWorkersKVNamespaceBindingType = "kv_namespace"
)

func (r ScriptVersionSettingGetResponseBindingsWorkersKVNamespaceBindingType) IsKnown() bool {
	switch r {
	case ScriptVersionSettingGetResponseBindingsWorkersKVNamespaceBindingTypeKVNamespace:
		return true
	}
	return false
}

type ScriptVersionSettingGetResponseBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type ScriptVersionSettingGetResponseBindingsWorkersServiceBindingType `json:"type,required"`
	JSON scriptVersionSettingGetResponseBindingsWorkersServiceBindingJSON `json:"-"`
}

// scriptVersionSettingGetResponseBindingsWorkersServiceBindingJSON contains the
// JSON metadata for the struct
// [ScriptVersionSettingGetResponseBindingsWorkersServiceBinding]
type scriptVersionSettingGetResponseBindingsWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingGetResponseBindingsWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingGetResponseBindingsWorkersServiceBindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionSettingGetResponseBindingsWorkersServiceBinding) implementsWorkersScriptVersionSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptVersionSettingGetResponseBindingsWorkersServiceBindingType string

const (
	ScriptVersionSettingGetResponseBindingsWorkersServiceBindingTypeService ScriptVersionSettingGetResponseBindingsWorkersServiceBindingType = "service"
)

func (r ScriptVersionSettingGetResponseBindingsWorkersServiceBindingType) IsKnown() bool {
	switch r {
	case ScriptVersionSettingGetResponseBindingsWorkersServiceBindingTypeService:
		return true
	}
	return false
}

type ScriptVersionSettingGetResponseBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type ScriptVersionSettingGetResponseBindingsWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                      `json:"script_name"`
	JSON       scriptVersionSettingGetResponseBindingsWorkersDoBindingJSON `json:"-"`
}

// scriptVersionSettingGetResponseBindingsWorkersDoBindingJSON contains the JSON
// metadata for the struct
// [ScriptVersionSettingGetResponseBindingsWorkersDoBinding]
type scriptVersionSettingGetResponseBindingsWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingGetResponseBindingsWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingGetResponseBindingsWorkersDoBindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionSettingGetResponseBindingsWorkersDoBinding) implementsWorkersScriptVersionSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptVersionSettingGetResponseBindingsWorkersDoBindingType string

const (
	ScriptVersionSettingGetResponseBindingsWorkersDoBindingTypeDurableObjectNamespace ScriptVersionSettingGetResponseBindingsWorkersDoBindingType = "durable_object_namespace"
)

func (r ScriptVersionSettingGetResponseBindingsWorkersDoBindingType) IsKnown() bool {
	switch r {
	case ScriptVersionSettingGetResponseBindingsWorkersDoBindingTypeDurableObjectNamespace:
		return true
	}
	return false
}

type ScriptVersionSettingGetResponseBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type ScriptVersionSettingGetResponseBindingsWorkersR2BindingType `json:"type,required"`
	JSON scriptVersionSettingGetResponseBindingsWorkersR2BindingJSON `json:"-"`
}

// scriptVersionSettingGetResponseBindingsWorkersR2BindingJSON contains the JSON
// metadata for the struct
// [ScriptVersionSettingGetResponseBindingsWorkersR2Binding]
type scriptVersionSettingGetResponseBindingsWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingGetResponseBindingsWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingGetResponseBindingsWorkersR2BindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionSettingGetResponseBindingsWorkersR2Binding) implementsWorkersScriptVersionSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptVersionSettingGetResponseBindingsWorkersR2BindingType string

const (
	ScriptVersionSettingGetResponseBindingsWorkersR2BindingTypeR2Bucket ScriptVersionSettingGetResponseBindingsWorkersR2BindingType = "r2_bucket"
)

func (r ScriptVersionSettingGetResponseBindingsWorkersR2BindingType) IsKnown() bool {
	switch r {
	case ScriptVersionSettingGetResponseBindingsWorkersR2BindingTypeR2Bucket:
		return true
	}
	return false
}

type ScriptVersionSettingGetResponseBindingsWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type ScriptVersionSettingGetResponseBindingsWorkersQueueBindingType `json:"type,required"`
	JSON scriptVersionSettingGetResponseBindingsWorkersQueueBindingJSON `json:"-"`
}

// scriptVersionSettingGetResponseBindingsWorkersQueueBindingJSON contains the JSON
// metadata for the struct
// [ScriptVersionSettingGetResponseBindingsWorkersQueueBinding]
type scriptVersionSettingGetResponseBindingsWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingGetResponseBindingsWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingGetResponseBindingsWorkersQueueBindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionSettingGetResponseBindingsWorkersQueueBinding) implementsWorkersScriptVersionSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptVersionSettingGetResponseBindingsWorkersQueueBindingType string

const (
	ScriptVersionSettingGetResponseBindingsWorkersQueueBindingTypeQueue ScriptVersionSettingGetResponseBindingsWorkersQueueBindingType = "queue"
)

func (r ScriptVersionSettingGetResponseBindingsWorkersQueueBindingType) IsKnown() bool {
	switch r {
	case ScriptVersionSettingGetResponseBindingsWorkersQueueBindingTypeQueue:
		return true
	}
	return false
}

type ScriptVersionSettingGetResponseBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type ScriptVersionSettingGetResponseBindingsWorkersD1BindingType `json:"type,required"`
	JSON scriptVersionSettingGetResponseBindingsWorkersD1BindingJSON `json:"-"`
}

// scriptVersionSettingGetResponseBindingsWorkersD1BindingJSON contains the JSON
// metadata for the struct
// [ScriptVersionSettingGetResponseBindingsWorkersD1Binding]
type scriptVersionSettingGetResponseBindingsWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingGetResponseBindingsWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingGetResponseBindingsWorkersD1BindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionSettingGetResponseBindingsWorkersD1Binding) implementsWorkersScriptVersionSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptVersionSettingGetResponseBindingsWorkersD1BindingType string

const (
	ScriptVersionSettingGetResponseBindingsWorkersD1BindingTypeD1 ScriptVersionSettingGetResponseBindingsWorkersD1BindingType = "d1"
)

func (r ScriptVersionSettingGetResponseBindingsWorkersD1BindingType) IsKnown() bool {
	switch r {
	case ScriptVersionSettingGetResponseBindingsWorkersD1BindingTypeD1:
		return true
	}
	return false
}

type ScriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type ScriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound ScriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     scriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON     `json:"-"`
}

// scriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [ScriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBinding]
type scriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBinding) implementsWorkersScriptVersionSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBindingType string

const (
	ScriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace ScriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

func (r ScriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBindingType) IsKnown() bool {
	switch r {
	case ScriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker
type ScriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker ScriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   scriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// scriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [ScriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound]
type scriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON) RawJSON() string {
	return r.raw
}

// Outbound worker
type ScriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                   `json:"service"`
	JSON    scriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// scriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [ScriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker]
type scriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionSettingGetResponseBindingsWorkersMTLSCERTBinding struct {
	Certificate interface{} `json:"certificate,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type ScriptVersionSettingGetResponseBindingsWorkersMTLSCERTBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                            `json:"certificate_id"`
	JSON          scriptVersionSettingGetResponseBindingsWorkersMtlscertBindingJSON `json:"-"`
}

// scriptVersionSettingGetResponseBindingsWorkersMtlscertBindingJSON contains the
// JSON metadata for the struct
// [ScriptVersionSettingGetResponseBindingsWorkersMTLSCERTBinding]
type scriptVersionSettingGetResponseBindingsWorkersMtlscertBindingJSON struct {
	Certificate   apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScriptVersionSettingGetResponseBindingsWorkersMTLSCERTBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingGetResponseBindingsWorkersMtlscertBindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionSettingGetResponseBindingsWorkersMTLSCERTBinding) implementsWorkersScriptVersionSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptVersionSettingGetResponseBindingsWorkersMTLSCERTBindingType string

const (
	ScriptVersionSettingGetResponseBindingsWorkersMTLSCERTBindingTypeMTLSCertificate ScriptVersionSettingGetResponseBindingsWorkersMTLSCERTBindingType = "mtls_certificate"
)

func (r ScriptVersionSettingGetResponseBindingsWorkersMTLSCERTBindingType) IsKnown() bool {
	switch r {
	case ScriptVersionSettingGetResponseBindingsWorkersMTLSCERTBindingTypeMTLSCertificate:
		return true
	}
	return false
}

// The class of resource that the binding provides.
type ScriptVersionSettingGetResponseBindingsType string

const (
	ScriptVersionSettingGetResponseBindingsTypeKVNamespace            ScriptVersionSettingGetResponseBindingsType = "kv_namespace"
	ScriptVersionSettingGetResponseBindingsTypeService                ScriptVersionSettingGetResponseBindingsType = "service"
	ScriptVersionSettingGetResponseBindingsTypeDurableObjectNamespace ScriptVersionSettingGetResponseBindingsType = "durable_object_namespace"
	ScriptVersionSettingGetResponseBindingsTypeR2Bucket               ScriptVersionSettingGetResponseBindingsType = "r2_bucket"
	ScriptVersionSettingGetResponseBindingsTypeQueue                  ScriptVersionSettingGetResponseBindingsType = "queue"
	ScriptVersionSettingGetResponseBindingsTypeD1                     ScriptVersionSettingGetResponseBindingsType = "d1"
	ScriptVersionSettingGetResponseBindingsTypeDispatchNamespace      ScriptVersionSettingGetResponseBindingsType = "dispatch_namespace"
	ScriptVersionSettingGetResponseBindingsTypeMTLSCertificate        ScriptVersionSettingGetResponseBindingsType = "mtls_certificate"
)

func (r ScriptVersionSettingGetResponseBindingsType) IsKnown() bool {
	switch r {
	case ScriptVersionSettingGetResponseBindingsTypeKVNamespace, ScriptVersionSettingGetResponseBindingsTypeService, ScriptVersionSettingGetResponseBindingsTypeDurableObjectNamespace, ScriptVersionSettingGetResponseBindingsTypeR2Bucket, ScriptVersionSettingGetResponseBindingsTypeQueue, ScriptVersionSettingGetResponseBindingsTypeD1, ScriptVersionSettingGetResponseBindingsTypeDispatchNamespace, ScriptVersionSettingGetResponseBindingsTypeMTLSCertificate:
		return true
	}
	return false
}

// Migrations to apply for Durable Objects associated with this Worker.
type ScriptVersionSettingGetResponseMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag             string                                        `json:"old_tag"`
	DeletedClasses     interface{}                                   `json:"deleted_classes,required"`
	NewClasses         interface{}                                   `json:"new_classes,required"`
	RenamedClasses     interface{}                                   `json:"renamed_classes,required"`
	TransferredClasses interface{}                                   `json:"transferred_classes,required"`
	Steps              interface{}                                   `json:"steps,required"`
	JSON               scriptVersionSettingGetResponseMigrationsJSON `json:"-"`
	union              ScriptVersionSettingGetResponseMigrationsUnion
}

// scriptVersionSettingGetResponseMigrationsJSON contains the JSON metadata for the
// struct [ScriptVersionSettingGetResponseMigrations]
type scriptVersionSettingGetResponseMigrationsJSON struct {
	NewTag             apijson.Field
	OldTag             apijson.Field
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	Steps              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r scriptVersionSettingGetResponseMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r *ScriptVersionSettingGetResponseMigrations) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r ScriptVersionSettingGetResponseMigrations) AsUnion() ScriptVersionSettingGetResponseMigrationsUnion {
	return r.union
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [workers.ScriptVersionSettingGetResponseMigrationsWorkersSingleStepMigrations]
// or [workers.ScriptVersionSettingGetResponseMigrationsWorkersSteppedMigrations].
type ScriptVersionSettingGetResponseMigrationsUnion interface {
	implementsWorkersScriptVersionSettingGetResponseMigrations()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScriptVersionSettingGetResponseMigrationsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptVersionSettingGetResponseMigrationsWorkersSingleStepMigrations{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptVersionSettingGetResponseMigrationsWorkersSteppedMigrations{}),
		},
	)
}

// A single set of migrations to apply.
type ScriptVersionSettingGetResponseMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses []ScriptVersionSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []ScriptVersionSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               scriptVersionSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON               `json:"-"`
}

// scriptVersionSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON
// contains the JSON metadata for the struct
// [ScriptVersionSettingGetResponseMigrationsWorkersSingleStepMigrations]
type scriptVersionSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ScriptVersionSettingGetResponseMigrationsWorkersSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionSettingGetResponseMigrationsWorkersSingleStepMigrations) implementsWorkersScriptVersionSettingGetResponseMigrations() {
}

type ScriptVersionSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From string                                                                               `json:"from"`
	To   string                                                                               `json:"to"`
	JSON scriptVersionSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// scriptVersionSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [ScriptVersionSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass]
type scriptVersionSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       string                                                                                   `json:"from"`
	FromScript string                                                                                   `json:"from_script"`
	To         string                                                                                   `json:"to"`
	JSON       scriptVersionSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// scriptVersionSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [ScriptVersionSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass]
type scriptVersionSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionSettingGetResponseMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []ScriptVersionSettingGetResponseMigrationsWorkersSteppedMigrationsStep `json:"steps"`
	JSON  scriptVersionSettingGetResponseMigrationsWorkersSteppedMigrationsJSON   `json:"-"`
}

// scriptVersionSettingGetResponseMigrationsWorkersSteppedMigrationsJSON contains
// the JSON metadata for the struct
// [ScriptVersionSettingGetResponseMigrationsWorkersSteppedMigrations]
type scriptVersionSettingGetResponseMigrationsWorkersSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingGetResponseMigrationsWorkersSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingGetResponseMigrationsWorkersSteppedMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r ScriptVersionSettingGetResponseMigrationsWorkersSteppedMigrations) implementsWorkersScriptVersionSettingGetResponseMigrations() {
}

type ScriptVersionSettingGetResponseMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []ScriptVersionSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []ScriptVersionSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               scriptVersionSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON                `json:"-"`
}

// scriptVersionSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON
// contains the JSON metadata for the struct
// [ScriptVersionSettingGetResponseMigrationsWorkersSteppedMigrationsStep]
type scriptVersionSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ScriptVersionSettingGetResponseMigrationsWorkersSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From string                                                                                 `json:"from"`
	To   string                                                                                 `json:"to"`
	JSON scriptVersionSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// scriptVersionSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [ScriptVersionSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass]
type scriptVersionSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                                     `json:"from"`
	FromScript string                                                                                     `json:"from_script"`
	To         string                                                                                     `json:"to"`
	JSON       scriptVersionSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// scriptVersionSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [ScriptVersionSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass]
type scriptVersionSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionSettingGetResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode ScriptVersionSettingGetResponsePlacementMode `json:"mode"`
	JSON scriptVersionSettingGetResponsePlacementJSON `json:"-"`
}

// scriptVersionSettingGetResponsePlacementJSON contains the JSON metadata for the
// struct [ScriptVersionSettingGetResponsePlacement]
type scriptVersionSettingGetResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingGetResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingGetResponsePlacementJSON) RawJSON() string {
	return r.raw
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type ScriptVersionSettingGetResponsePlacementMode string

const (
	ScriptVersionSettingGetResponsePlacementModeSmart ScriptVersionSettingGetResponsePlacementMode = "smart"
)

func (r ScriptVersionSettingGetResponsePlacementMode) IsKnown() bool {
	switch r {
	case ScriptVersionSettingGetResponsePlacementModeSmart:
		return true
	}
	return false
}

// A reference to a script that will consume logs from the attached Worker.
type ScriptVersionSettingGetResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                          `json:"namespace"`
	JSON      scriptVersionSettingGetResponseTailConsumerJSON `json:"-"`
}

// scriptVersionSettingGetResponseTailConsumerJSON contains the JSON metadata for
// the struct [ScriptVersionSettingGetResponseTailConsumer]
type scriptVersionSettingGetResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingGetResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingGetResponseTailConsumerJSON) RawJSON() string {
	return r.raw
}

type ScriptVersionSettingEditParams struct {
	// Identifier
	AccountID param.Field[string]                                 `path:"account_id,required"`
	Settings  param.Field[ScriptVersionSettingEditParamsSettings] `json:"settings"`
}

func (r ScriptVersionSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptVersionSettingEditParamsSettings struct {
	// List of bindings attached to this Worker
	Bindings param.Field[[]ScriptVersionSettingEditParamsSettingsBindingUnion] `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations param.Field[ScriptVersionSettingEditParamsSettingsMigrationsUnion] `json:"migrations"`
	Placement  param.Field[ScriptVersionSettingEditParamsSettingsPlacement]       `json:"placement"`
	// Tags to help you manage your Workers
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]ScriptVersionSettingEditParamsSettingsTailConsumer] `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel param.Field[string] `json:"usage_model"`
}

func (r ScriptVersionSettingEditParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A binding to allow the Worker to communicate with resources
type ScriptVersionSettingEditParamsSettingsBinding struct {
	// The class of resource that the binding provides.
	Type param.Field[ScriptVersionSettingEditParamsSettingsBindingsType] `json:"type,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Name of Worker to bind to
	Service param.Field[string] `json:"service"`
	// The exported class name of the Durable Object
	ClassName param.Field[string] `json:"class_name"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName param.Field[string] `json:"script_name"`
	// R2 bucket to bind to
	BucketName param.Field[string] `json:"bucket_name"`
	// Name of the Queue to bind to
	QueueName param.Field[string] `json:"queue_name"`
	// ID of the D1 database to bind to
	ID param.Field[string] `json:"id"`
	// Namespace to bind to
	Namespace param.Field[string]      `json:"namespace"`
	Outbound  param.Field[interface{}] `json:"outbound,required"`
	// ID of the certificate to bind to
	CertificateID param.Field[string]      `json:"certificate_id"`
	Certificate   param.Field[interface{}] `json:"certificate,required"`
}

func (r ScriptVersionSettingEditParamsSettingsBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionSettingEditParamsSettingsBinding) implementsWorkersScriptVersionSettingEditParamsSettingsBindingUnion() {
}

// A binding to allow the Worker to communicate with resources
//
// Satisfied by
// [workers.ScriptVersionSettingEditParamsSettingsBindingsWorkersKVNamespaceBinding],
// [workers.ScriptVersionSettingEditParamsSettingsBindingsWorkersServiceBinding],
// [workers.ScriptVersionSettingEditParamsSettingsBindingsWorkersDoBinding],
// [workers.ScriptVersionSettingEditParamsSettingsBindingsWorkersR2Binding],
// [workers.ScriptVersionSettingEditParamsSettingsBindingsWorkersQueueBinding],
// [workers.ScriptVersionSettingEditParamsSettingsBindingsWorkersD1Binding],
// [workers.ScriptVersionSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBinding],
// [workers.ScriptVersionSettingEditParamsSettingsBindingsWorkersMTLSCERTBinding],
// [ScriptVersionSettingEditParamsSettingsBinding].
type ScriptVersionSettingEditParamsSettingsBindingUnion interface {
	implementsWorkersScriptVersionSettingEditParamsSettingsBindingUnion()
}

type ScriptVersionSettingEditParamsSettingsBindingsWorkersKVNamespaceBinding struct {
	// The class of resource that the binding provides.
	Type param.Field[ScriptVersionSettingEditParamsSettingsBindingsWorkersKVNamespaceBindingType] `json:"type,required"`
}

func (r ScriptVersionSettingEditParamsSettingsBindingsWorkersKVNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionSettingEditParamsSettingsBindingsWorkersKVNamespaceBinding) implementsWorkersScriptVersionSettingEditParamsSettingsBindingUnion() {
}

// The class of resource that the binding provides.
type ScriptVersionSettingEditParamsSettingsBindingsWorkersKVNamespaceBindingType string

const (
	ScriptVersionSettingEditParamsSettingsBindingsWorkersKVNamespaceBindingTypeKVNamespace ScriptVersionSettingEditParamsSettingsBindingsWorkersKVNamespaceBindingType = "kv_namespace"
)

func (r ScriptVersionSettingEditParamsSettingsBindingsWorkersKVNamespaceBindingType) IsKnown() bool {
	switch r {
	case ScriptVersionSettingEditParamsSettingsBindingsWorkersKVNamespaceBindingTypeKVNamespace:
		return true
	}
	return false
}

type ScriptVersionSettingEditParamsSettingsBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// Name of Worker to bind to
	Service param.Field[string] `json:"service,required"`
	// The class of resource that the binding provides.
	Type param.Field[ScriptVersionSettingEditParamsSettingsBindingsWorkersServiceBindingType] `json:"type,required"`
}

func (r ScriptVersionSettingEditParamsSettingsBindingsWorkersServiceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionSettingEditParamsSettingsBindingsWorkersServiceBinding) implementsWorkersScriptVersionSettingEditParamsSettingsBindingUnion() {
}

// The class of resource that the binding provides.
type ScriptVersionSettingEditParamsSettingsBindingsWorkersServiceBindingType string

const (
	ScriptVersionSettingEditParamsSettingsBindingsWorkersServiceBindingTypeService ScriptVersionSettingEditParamsSettingsBindingsWorkersServiceBindingType = "service"
)

func (r ScriptVersionSettingEditParamsSettingsBindingsWorkersServiceBindingType) IsKnown() bool {
	switch r {
	case ScriptVersionSettingEditParamsSettingsBindingsWorkersServiceBindingTypeService:
		return true
	}
	return false
}

type ScriptVersionSettingEditParamsSettingsBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName param.Field[string] `json:"class_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[ScriptVersionSettingEditParamsSettingsBindingsWorkersDoBindingType] `json:"type,required"`
	// The environment of the script_name to bind to
	Environment param.Field[string] `json:"environment"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName param.Field[string] `json:"script_name"`
}

func (r ScriptVersionSettingEditParamsSettingsBindingsWorkersDoBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionSettingEditParamsSettingsBindingsWorkersDoBinding) implementsWorkersScriptVersionSettingEditParamsSettingsBindingUnion() {
}

// The class of resource that the binding provides.
type ScriptVersionSettingEditParamsSettingsBindingsWorkersDoBindingType string

const (
	ScriptVersionSettingEditParamsSettingsBindingsWorkersDoBindingTypeDurableObjectNamespace ScriptVersionSettingEditParamsSettingsBindingsWorkersDoBindingType = "durable_object_namespace"
)

func (r ScriptVersionSettingEditParamsSettingsBindingsWorkersDoBindingType) IsKnown() bool {
	switch r {
	case ScriptVersionSettingEditParamsSettingsBindingsWorkersDoBindingTypeDurableObjectNamespace:
		return true
	}
	return false
}

type ScriptVersionSettingEditParamsSettingsBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName param.Field[string] `json:"bucket_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[ScriptVersionSettingEditParamsSettingsBindingsWorkersR2BindingType] `json:"type,required"`
}

func (r ScriptVersionSettingEditParamsSettingsBindingsWorkersR2Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionSettingEditParamsSettingsBindingsWorkersR2Binding) implementsWorkersScriptVersionSettingEditParamsSettingsBindingUnion() {
}

// The class of resource that the binding provides.
type ScriptVersionSettingEditParamsSettingsBindingsWorkersR2BindingType string

const (
	ScriptVersionSettingEditParamsSettingsBindingsWorkersR2BindingTypeR2Bucket ScriptVersionSettingEditParamsSettingsBindingsWorkersR2BindingType = "r2_bucket"
)

func (r ScriptVersionSettingEditParamsSettingsBindingsWorkersR2BindingType) IsKnown() bool {
	switch r {
	case ScriptVersionSettingEditParamsSettingsBindingsWorkersR2BindingTypeR2Bucket:
		return true
	}
	return false
}

type ScriptVersionSettingEditParamsSettingsBindingsWorkersQueueBinding struct {
	// Name of the Queue to bind to
	QueueName param.Field[string] `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[ScriptVersionSettingEditParamsSettingsBindingsWorkersQueueBindingType] `json:"type,required"`
}

func (r ScriptVersionSettingEditParamsSettingsBindingsWorkersQueueBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionSettingEditParamsSettingsBindingsWorkersQueueBinding) implementsWorkersScriptVersionSettingEditParamsSettingsBindingUnion() {
}

// The class of resource that the binding provides.
type ScriptVersionSettingEditParamsSettingsBindingsWorkersQueueBindingType string

const (
	ScriptVersionSettingEditParamsSettingsBindingsWorkersQueueBindingTypeQueue ScriptVersionSettingEditParamsSettingsBindingsWorkersQueueBindingType = "queue"
)

func (r ScriptVersionSettingEditParamsSettingsBindingsWorkersQueueBindingType) IsKnown() bool {
	switch r {
	case ScriptVersionSettingEditParamsSettingsBindingsWorkersQueueBindingTypeQueue:
		return true
	}
	return false
}

type ScriptVersionSettingEditParamsSettingsBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID param.Field[string] `json:"id,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name param.Field[string] `json:"name,required"`
	// The class of resource that the binding provides.
	Type param.Field[ScriptVersionSettingEditParamsSettingsBindingsWorkersD1BindingType] `json:"type,required"`
}

func (r ScriptVersionSettingEditParamsSettingsBindingsWorkersD1Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionSettingEditParamsSettingsBindingsWorkersD1Binding) implementsWorkersScriptVersionSettingEditParamsSettingsBindingUnion() {
}

// The class of resource that the binding provides.
type ScriptVersionSettingEditParamsSettingsBindingsWorkersD1BindingType string

const (
	ScriptVersionSettingEditParamsSettingsBindingsWorkersD1BindingTypeD1 ScriptVersionSettingEditParamsSettingsBindingsWorkersD1BindingType = "d1"
)

func (r ScriptVersionSettingEditParamsSettingsBindingsWorkersD1BindingType) IsKnown() bool {
	switch r {
	case ScriptVersionSettingEditParamsSettingsBindingsWorkersD1BindingTypeD1:
		return true
	}
	return false
}

type ScriptVersionSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBinding struct {
	// Namespace to bind to
	Namespace param.Field[string] `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type param.Field[ScriptVersionSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBindingType] `json:"type,required"`
	// Outbound worker
	Outbound param.Field[ScriptVersionSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBindingOutbound] `json:"outbound"`
}

func (r ScriptVersionSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBinding) implementsWorkersScriptVersionSettingEditParamsSettingsBindingUnion() {
}

// The class of resource that the binding provides.
type ScriptVersionSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBindingType string

const (
	ScriptVersionSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace ScriptVersionSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

func (r ScriptVersionSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBindingType) IsKnown() bool {
	switch r {
	case ScriptVersionSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker
type ScriptVersionSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params param.Field[[]string] `json:"params"`
	// Outbound worker
	Worker param.Field[ScriptVersionSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBindingOutboundWorker] `json:"worker"`
}

func (r ScriptVersionSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBindingOutbound) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker
type ScriptVersionSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker
	Service param.Field[string] `json:"service"`
}

func (r ScriptVersionSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBindingOutboundWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptVersionSettingEditParamsSettingsBindingsWorkersMTLSCERTBinding struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
	// The class of resource that the binding provides.
	Type param.Field[ScriptVersionSettingEditParamsSettingsBindingsWorkersMTLSCERTBindingType] `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID param.Field[string] `json:"certificate_id"`
}

func (r ScriptVersionSettingEditParamsSettingsBindingsWorkersMTLSCERTBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionSettingEditParamsSettingsBindingsWorkersMTLSCERTBinding) implementsWorkersScriptVersionSettingEditParamsSettingsBindingUnion() {
}

// The class of resource that the binding provides.
type ScriptVersionSettingEditParamsSettingsBindingsWorkersMTLSCERTBindingType string

const (
	ScriptVersionSettingEditParamsSettingsBindingsWorkersMTLSCERTBindingTypeMTLSCertificate ScriptVersionSettingEditParamsSettingsBindingsWorkersMTLSCERTBindingType = "mtls_certificate"
)

func (r ScriptVersionSettingEditParamsSettingsBindingsWorkersMTLSCERTBindingType) IsKnown() bool {
	switch r {
	case ScriptVersionSettingEditParamsSettingsBindingsWorkersMTLSCERTBindingTypeMTLSCertificate:
		return true
	}
	return false
}

// The class of resource that the binding provides.
type ScriptVersionSettingEditParamsSettingsBindingsType string

const (
	ScriptVersionSettingEditParamsSettingsBindingsTypeKVNamespace            ScriptVersionSettingEditParamsSettingsBindingsType = "kv_namespace"
	ScriptVersionSettingEditParamsSettingsBindingsTypeService                ScriptVersionSettingEditParamsSettingsBindingsType = "service"
	ScriptVersionSettingEditParamsSettingsBindingsTypeDurableObjectNamespace ScriptVersionSettingEditParamsSettingsBindingsType = "durable_object_namespace"
	ScriptVersionSettingEditParamsSettingsBindingsTypeR2Bucket               ScriptVersionSettingEditParamsSettingsBindingsType = "r2_bucket"
	ScriptVersionSettingEditParamsSettingsBindingsTypeQueue                  ScriptVersionSettingEditParamsSettingsBindingsType = "queue"
	ScriptVersionSettingEditParamsSettingsBindingsTypeD1                     ScriptVersionSettingEditParamsSettingsBindingsType = "d1"
	ScriptVersionSettingEditParamsSettingsBindingsTypeDispatchNamespace      ScriptVersionSettingEditParamsSettingsBindingsType = "dispatch_namespace"
	ScriptVersionSettingEditParamsSettingsBindingsTypeMTLSCertificate        ScriptVersionSettingEditParamsSettingsBindingsType = "mtls_certificate"
)

func (r ScriptVersionSettingEditParamsSettingsBindingsType) IsKnown() bool {
	switch r {
	case ScriptVersionSettingEditParamsSettingsBindingsTypeKVNamespace, ScriptVersionSettingEditParamsSettingsBindingsTypeService, ScriptVersionSettingEditParamsSettingsBindingsTypeDurableObjectNamespace, ScriptVersionSettingEditParamsSettingsBindingsTypeR2Bucket, ScriptVersionSettingEditParamsSettingsBindingsTypeQueue, ScriptVersionSettingEditParamsSettingsBindingsTypeD1, ScriptVersionSettingEditParamsSettingsBindingsTypeDispatchNamespace, ScriptVersionSettingEditParamsSettingsBindingsTypeMTLSCertificate:
		return true
	}
	return false
}

// Migrations to apply for Durable Objects associated with this Worker.
type ScriptVersionSettingEditParamsSettingsMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag             param.Field[string]      `json:"old_tag"`
	DeletedClasses     param.Field[interface{}] `json:"deleted_classes,required"`
	NewClasses         param.Field[interface{}] `json:"new_classes,required"`
	RenamedClasses     param.Field[interface{}] `json:"renamed_classes,required"`
	TransferredClasses param.Field[interface{}] `json:"transferred_classes,required"`
	Steps              param.Field[interface{}] `json:"steps,required"`
}

func (r ScriptVersionSettingEditParamsSettingsMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionSettingEditParamsSettingsMigrations) implementsWorkersScriptVersionSettingEditParamsSettingsMigrationsUnion() {
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by
// [workers.ScriptVersionSettingEditParamsSettingsMigrationsWorkersSingleStepMigrations],
// [workers.ScriptVersionSettingEditParamsSettingsMigrationsWorkersSteppedMigrations],
// [ScriptVersionSettingEditParamsSettingsMigrations].
type ScriptVersionSettingEditParamsSettingsMigrationsUnion interface {
	implementsWorkersScriptVersionSettingEditParamsSettingsMigrationsUnion()
}

// A single set of migrations to apply.
type ScriptVersionSettingEditParamsSettingsMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses param.Field[[]ScriptVersionSettingEditParamsSettingsMigrationsWorkersSingleStepMigrationsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]ScriptVersionSettingEditParamsSettingsMigrationsWorkersSingleStepMigrationsTransferredClass] `json:"transferred_classes"`
}

func (r ScriptVersionSettingEditParamsSettingsMigrationsWorkersSingleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionSettingEditParamsSettingsMigrationsWorkersSingleStepMigrations) implementsWorkersScriptVersionSettingEditParamsSettingsMigrationsUnion() {
}

type ScriptVersionSettingEditParamsSettingsMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r ScriptVersionSettingEditParamsSettingsMigrationsWorkersSingleStepMigrationsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptVersionSettingEditParamsSettingsMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r ScriptVersionSettingEditParamsSettingsMigrationsWorkersSingleStepMigrationsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptVersionSettingEditParamsSettingsMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]ScriptVersionSettingEditParamsSettingsMigrationsWorkersSteppedMigrationsStep] `json:"steps"`
}

func (r ScriptVersionSettingEditParamsSettingsMigrationsWorkersSteppedMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptVersionSettingEditParamsSettingsMigrationsWorkersSteppedMigrations) implementsWorkersScriptVersionSettingEditParamsSettingsMigrationsUnion() {
}

type ScriptVersionSettingEditParamsSettingsMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]ScriptVersionSettingEditParamsSettingsMigrationsWorkersSteppedMigrationsStepsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]ScriptVersionSettingEditParamsSettingsMigrationsWorkersSteppedMigrationsStepsTransferredClass] `json:"transferred_classes"`
}

func (r ScriptVersionSettingEditParamsSettingsMigrationsWorkersSteppedMigrationsStep) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptVersionSettingEditParamsSettingsMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r ScriptVersionSettingEditParamsSettingsMigrationsWorkersSteppedMigrationsStepsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptVersionSettingEditParamsSettingsMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r ScriptVersionSettingEditParamsSettingsMigrationsWorkersSteppedMigrationsStepsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptVersionSettingEditParamsSettingsPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode param.Field[ScriptVersionSettingEditParamsSettingsPlacementMode] `json:"mode"`
}

func (r ScriptVersionSettingEditParamsSettingsPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type ScriptVersionSettingEditParamsSettingsPlacementMode string

const (
	ScriptVersionSettingEditParamsSettingsPlacementModeSmart ScriptVersionSettingEditParamsSettingsPlacementMode = "smart"
)

func (r ScriptVersionSettingEditParamsSettingsPlacementMode) IsKnown() bool {
	switch r {
	case ScriptVersionSettingEditParamsSettingsPlacementModeSmart:
		return true
	}
	return false
}

// A reference to a script that will consume logs from the attached Worker.
type ScriptVersionSettingEditParamsSettingsTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service param.Field[string] `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace param.Field[string] `json:"namespace"`
}

func (r ScriptVersionSettingEditParamsSettingsTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptVersionSettingEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   ScriptVersionSettingEditResponse                          `json:"result,required"`
	// Whether the API call was successful
	Success ScriptVersionSettingEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptVersionSettingEditResponseEnvelopeJSON    `json:"-"`
}

// scriptVersionSettingEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [ScriptVersionSettingEditResponseEnvelope]
type scriptVersionSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptVersionSettingEditResponseEnvelopeSuccess bool

const (
	ScriptVersionSettingEditResponseEnvelopeSuccessTrue ScriptVersionSettingEditResponseEnvelopeSuccess = true
)

func (r ScriptVersionSettingEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptVersionSettingEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptVersionSettingGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptVersionSettingGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   ScriptVersionSettingGetResponse                           `json:"result,required"`
	// Whether the API call was successful
	Success ScriptVersionSettingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptVersionSettingGetResponseEnvelopeJSON    `json:"-"`
}

// scriptVersionSettingGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ScriptVersionSettingGetResponseEnvelope]
type scriptVersionSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptVersionSettingGetResponseEnvelopeSuccess bool

const (
	ScriptVersionSettingGetResponseEnvelopeSuccessTrue ScriptVersionSettingGetResponseEnvelopeSuccess = true
)

func (r ScriptVersionSettingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptVersionSettingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
