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

// ScriptSettingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewScriptSettingService] method
// instead.
type ScriptSettingService struct {
	Options []option.RequestOption
}

// NewScriptSettingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScriptSettingService(opts ...option.RequestOption) (r *ScriptSettingService) {
	r = &ScriptSettingService{}
	r.Options = opts
	return
}

// Patch metadata or config, such as bindings or usage model
func (r *ScriptSettingService) Edit(ctx context.Context, scriptName string, params ScriptSettingEditParams, opts ...option.RequestOption) (res *ScriptSettingEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ScriptSettingEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/settings", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get metadata and config, such as bindings or usage model
func (r *ScriptSettingService) Get(ctx context.Context, scriptName string, query ScriptSettingGetParams, opts ...option.RequestOption) (res *ScriptSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ScriptSettingGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/settings", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ScriptSettingEditResponse struct {
	// List of bindings attached to this Worker
	Bindings []ScriptSettingEditResponseBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations ScriptSettingEditResponseMigrations `json:"migrations"`
	Placement  ScriptSettingEditResponsePlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []ScriptSettingEditResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                        `json:"usage_model"`
	JSON       scriptSettingEditResponseJSON `json:"-"`
}

// scriptSettingEditResponseJSON contains the JSON metadata for the struct
// [ScriptSettingEditResponse]
type scriptSettingEditResponseJSON struct {
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

func (r *ScriptSettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseJSON) RawJSON() string {
	return r.raw
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [workers.ScriptSettingEditResponseBindingsWorkersKVNamespaceBinding],
// [workers.ScriptSettingEditResponseBindingsWorkersServiceBinding],
// [workers.ScriptSettingEditResponseBindingsWorkersDoBinding],
// [workers.ScriptSettingEditResponseBindingsWorkersR2Binding],
// [workers.ScriptSettingEditResponseBindingsWorkersQueueBinding],
// [workers.ScriptSettingEditResponseBindingsWorkersD1Binding],
// [workers.ScriptSettingEditResponseBindingsWorkersDispatchNamespaceBinding] or
// [workers.ScriptSettingEditResponseBindingsWorkersMTLSCERTBinding].
type ScriptSettingEditResponseBinding interface {
	implementsWorkersScriptSettingEditResponseBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScriptSettingEditResponseBinding)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptSettingEditResponseBindingsWorkersKVNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptSettingEditResponseBindingsWorkersServiceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptSettingEditResponseBindingsWorkersDoBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptSettingEditResponseBindingsWorkersR2Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptSettingEditResponseBindingsWorkersQueueBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptSettingEditResponseBindingsWorkersD1Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptSettingEditResponseBindingsWorkersDispatchNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptSettingEditResponseBindingsWorkersMTLSCERTBinding{}),
		},
	)
}

type ScriptSettingEditResponseBindingsWorkersKVNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type ScriptSettingEditResponseBindingsWorkersKVNamespaceBindingType `json:"type,required"`
	JSON scriptSettingEditResponseBindingsWorkersKVNamespaceBindingJSON `json:"-"`
}

// scriptSettingEditResponseBindingsWorkersKVNamespaceBindingJSON contains the JSON
// metadata for the struct
// [ScriptSettingEditResponseBindingsWorkersKVNamespaceBinding]
type scriptSettingEditResponseBindingsWorkersKVNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingEditResponseBindingsWorkersKVNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseBindingsWorkersKVNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptSettingEditResponseBindingsWorkersKVNamespaceBinding) implementsWorkersScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingEditResponseBindingsWorkersKVNamespaceBindingType string

const (
	ScriptSettingEditResponseBindingsWorkersKVNamespaceBindingTypeKVNamespace ScriptSettingEditResponseBindingsWorkersKVNamespaceBindingType = "kv_namespace"
)

func (r ScriptSettingEditResponseBindingsWorkersKVNamespaceBindingType) IsKnown() bool {
	switch r {
	case ScriptSettingEditResponseBindingsWorkersKVNamespaceBindingTypeKVNamespace:
		return true
	}
	return false
}

type ScriptSettingEditResponseBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type ScriptSettingEditResponseBindingsWorkersServiceBindingType `json:"type,required"`
	JSON scriptSettingEditResponseBindingsWorkersServiceBindingJSON `json:"-"`
}

// scriptSettingEditResponseBindingsWorkersServiceBindingJSON contains the JSON
// metadata for the struct [ScriptSettingEditResponseBindingsWorkersServiceBinding]
type scriptSettingEditResponseBindingsWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingEditResponseBindingsWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseBindingsWorkersServiceBindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptSettingEditResponseBindingsWorkersServiceBinding) implementsWorkersScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingEditResponseBindingsWorkersServiceBindingType string

const (
	ScriptSettingEditResponseBindingsWorkersServiceBindingTypeService ScriptSettingEditResponseBindingsWorkersServiceBindingType = "service"
)

func (r ScriptSettingEditResponseBindingsWorkersServiceBindingType) IsKnown() bool {
	switch r {
	case ScriptSettingEditResponseBindingsWorkersServiceBindingTypeService:
		return true
	}
	return false
}

type ScriptSettingEditResponseBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type ScriptSettingEditResponseBindingsWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                `json:"script_name"`
	JSON       scriptSettingEditResponseBindingsWorkersDoBindingJSON `json:"-"`
}

// scriptSettingEditResponseBindingsWorkersDoBindingJSON contains the JSON metadata
// for the struct [ScriptSettingEditResponseBindingsWorkersDoBinding]
type scriptSettingEditResponseBindingsWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingEditResponseBindingsWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseBindingsWorkersDoBindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptSettingEditResponseBindingsWorkersDoBinding) implementsWorkersScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingEditResponseBindingsWorkersDoBindingType string

const (
	ScriptSettingEditResponseBindingsWorkersDoBindingTypeDurableObjectNamespace ScriptSettingEditResponseBindingsWorkersDoBindingType = "durable_object_namespace"
)

func (r ScriptSettingEditResponseBindingsWorkersDoBindingType) IsKnown() bool {
	switch r {
	case ScriptSettingEditResponseBindingsWorkersDoBindingTypeDurableObjectNamespace:
		return true
	}
	return false
}

type ScriptSettingEditResponseBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type ScriptSettingEditResponseBindingsWorkersR2BindingType `json:"type,required"`
	JSON scriptSettingEditResponseBindingsWorkersR2BindingJSON `json:"-"`
}

// scriptSettingEditResponseBindingsWorkersR2BindingJSON contains the JSON metadata
// for the struct [ScriptSettingEditResponseBindingsWorkersR2Binding]
type scriptSettingEditResponseBindingsWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingEditResponseBindingsWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseBindingsWorkersR2BindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptSettingEditResponseBindingsWorkersR2Binding) implementsWorkersScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingEditResponseBindingsWorkersR2BindingType string

const (
	ScriptSettingEditResponseBindingsWorkersR2BindingTypeR2Bucket ScriptSettingEditResponseBindingsWorkersR2BindingType = "r2_bucket"
)

func (r ScriptSettingEditResponseBindingsWorkersR2BindingType) IsKnown() bool {
	switch r {
	case ScriptSettingEditResponseBindingsWorkersR2BindingTypeR2Bucket:
		return true
	}
	return false
}

type ScriptSettingEditResponseBindingsWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type ScriptSettingEditResponseBindingsWorkersQueueBindingType `json:"type,required"`
	JSON scriptSettingEditResponseBindingsWorkersQueueBindingJSON `json:"-"`
}

// scriptSettingEditResponseBindingsWorkersQueueBindingJSON contains the JSON
// metadata for the struct [ScriptSettingEditResponseBindingsWorkersQueueBinding]
type scriptSettingEditResponseBindingsWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingEditResponseBindingsWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseBindingsWorkersQueueBindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptSettingEditResponseBindingsWorkersQueueBinding) implementsWorkersScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingEditResponseBindingsWorkersQueueBindingType string

const (
	ScriptSettingEditResponseBindingsWorkersQueueBindingTypeQueue ScriptSettingEditResponseBindingsWorkersQueueBindingType = "queue"
)

func (r ScriptSettingEditResponseBindingsWorkersQueueBindingType) IsKnown() bool {
	switch r {
	case ScriptSettingEditResponseBindingsWorkersQueueBindingTypeQueue:
		return true
	}
	return false
}

type ScriptSettingEditResponseBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type ScriptSettingEditResponseBindingsWorkersD1BindingType `json:"type,required"`
	JSON scriptSettingEditResponseBindingsWorkersD1BindingJSON `json:"-"`
}

// scriptSettingEditResponseBindingsWorkersD1BindingJSON contains the JSON metadata
// for the struct [ScriptSettingEditResponseBindingsWorkersD1Binding]
type scriptSettingEditResponseBindingsWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingEditResponseBindingsWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseBindingsWorkersD1BindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptSettingEditResponseBindingsWorkersD1Binding) implementsWorkersScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingEditResponseBindingsWorkersD1BindingType string

const (
	ScriptSettingEditResponseBindingsWorkersD1BindingTypeD1 ScriptSettingEditResponseBindingsWorkersD1BindingType = "d1"
)

func (r ScriptSettingEditResponseBindingsWorkersD1BindingType) IsKnown() bool {
	switch r {
	case ScriptSettingEditResponseBindingsWorkersD1BindingTypeD1:
		return true
	}
	return false
}

type ScriptSettingEditResponseBindingsWorkersDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type ScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound ScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     scriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON     `json:"-"`
}

// scriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON contains
// the JSON metadata for the struct
// [ScriptSettingEditResponseBindingsWorkersDispatchNamespaceBinding]
type scriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingEditResponseBindingsWorkersDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptSettingEditResponseBindingsWorkersDispatchNamespaceBinding) implementsWorkersScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingType string

const (
	ScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace ScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

func (r ScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingType) IsKnown() bool {
	switch r {
	case ScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker
type ScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker ScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   scriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// scriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [ScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound]
type scriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON) RawJSON() string {
	return r.raw
}

// Outbound worker
type ScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                             `json:"service"`
	JSON    scriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// scriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [ScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker]
type scriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type ScriptSettingEditResponseBindingsWorkersMTLSCERTBinding struct {
	Certificate interface{} `json:"certificate,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type ScriptSettingEditResponseBindingsWorkersMTLSCERTBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                      `json:"certificate_id"`
	JSON          scriptSettingEditResponseBindingsWorkersMtlscertBindingJSON `json:"-"`
}

// scriptSettingEditResponseBindingsWorkersMtlscertBindingJSON contains the JSON
// metadata for the struct
// [ScriptSettingEditResponseBindingsWorkersMTLSCERTBinding]
type scriptSettingEditResponseBindingsWorkersMtlscertBindingJSON struct {
	Certificate   apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScriptSettingEditResponseBindingsWorkersMTLSCERTBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseBindingsWorkersMtlscertBindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptSettingEditResponseBindingsWorkersMTLSCERTBinding) implementsWorkersScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingEditResponseBindingsWorkersMTLSCERTBindingType string

const (
	ScriptSettingEditResponseBindingsWorkersMTLSCERTBindingTypeMTLSCertificate ScriptSettingEditResponseBindingsWorkersMTLSCERTBindingType = "mtls_certificate"
)

func (r ScriptSettingEditResponseBindingsWorkersMTLSCERTBindingType) IsKnown() bool {
	switch r {
	case ScriptSettingEditResponseBindingsWorkersMTLSCERTBindingTypeMTLSCertificate:
		return true
	}
	return false
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [workers.ScriptSettingEditResponseMigrationsWorkersSingleStepMigrations] or
// [workers.ScriptSettingEditResponseMigrationsWorkersSteppedMigrations].
type ScriptSettingEditResponseMigrations interface {
	implementsWorkersScriptSettingEditResponseMigrations()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScriptSettingEditResponseMigrations)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptSettingEditResponseMigrationsWorkersSingleStepMigrations{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptSettingEditResponseMigrationsWorkersSteppedMigrations{}),
		},
	)
}

// A single set of migrations to apply.
type ScriptSettingEditResponseMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses []ScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []ScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               scriptSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON               `json:"-"`
}

// scriptSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON contains the
// JSON metadata for the struct
// [ScriptSettingEditResponseMigrationsWorkersSingleStepMigrations]
type scriptSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ScriptSettingEditResponseMigrationsWorkersSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r ScriptSettingEditResponseMigrationsWorkersSingleStepMigrations) implementsWorkersScriptSettingEditResponseMigrations() {
}

type ScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From string                                                                         `json:"from"`
	To   string                                                                         `json:"to"`
	JSON scriptSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// scriptSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [ScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass]
type scriptSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON) RawJSON() string {
	return r.raw
}

type ScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       string                                                                             `json:"from"`
	FromScript string                                                                             `json:"from_script"`
	To         string                                                                             `json:"to"`
	JSON       scriptSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// scriptSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [ScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass]
type scriptSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON) RawJSON() string {
	return r.raw
}

type ScriptSettingEditResponseMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []ScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStep `json:"steps"`
	JSON  scriptSettingEditResponseMigrationsWorkersSteppedMigrationsJSON   `json:"-"`
}

// scriptSettingEditResponseMigrationsWorkersSteppedMigrationsJSON contains the
// JSON metadata for the struct
// [ScriptSettingEditResponseMigrationsWorkersSteppedMigrations]
type scriptSettingEditResponseMigrationsWorkersSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingEditResponseMigrationsWorkersSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseMigrationsWorkersSteppedMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r ScriptSettingEditResponseMigrationsWorkersSteppedMigrations) implementsWorkersScriptSettingEditResponseMigrations() {
}

type ScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []ScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []ScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               scriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON                `json:"-"`
}

// scriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON contains the
// JSON metadata for the struct
// [ScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStep]
type scriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON) RawJSON() string {
	return r.raw
}

type ScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From string                                                                           `json:"from"`
	To   string                                                                           `json:"to"`
	JSON scriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// scriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [ScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass]
type scriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON) RawJSON() string {
	return r.raw
}

type ScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                               `json:"from"`
	FromScript string                                                                               `json:"from_script"`
	To         string                                                                               `json:"to"`
	JSON       scriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// scriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [ScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass]
type scriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON) RawJSON() string {
	return r.raw
}

type ScriptSettingEditResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode ScriptSettingEditResponsePlacementMode `json:"mode"`
	JSON scriptSettingEditResponsePlacementJSON `json:"-"`
}

// scriptSettingEditResponsePlacementJSON contains the JSON metadata for the struct
// [ScriptSettingEditResponsePlacement]
type scriptSettingEditResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingEditResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponsePlacementJSON) RawJSON() string {
	return r.raw
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type ScriptSettingEditResponsePlacementMode string

const (
	ScriptSettingEditResponsePlacementModeSmart ScriptSettingEditResponsePlacementMode = "smart"
)

func (r ScriptSettingEditResponsePlacementMode) IsKnown() bool {
	switch r {
	case ScriptSettingEditResponsePlacementModeSmart:
		return true
	}
	return false
}

// A reference to a script that will consume logs from the attached Worker.
type ScriptSettingEditResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                    `json:"namespace"`
	JSON      scriptSettingEditResponseTailConsumerJSON `json:"-"`
}

// scriptSettingEditResponseTailConsumerJSON contains the JSON metadata for the
// struct [ScriptSettingEditResponseTailConsumer]
type scriptSettingEditResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingEditResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseTailConsumerJSON) RawJSON() string {
	return r.raw
}

type ScriptSettingGetResponse struct {
	// List of bindings attached to this Worker
	Bindings []ScriptSettingGetResponseBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations ScriptSettingGetResponseMigrations `json:"migrations"`
	Placement  ScriptSettingGetResponsePlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []ScriptSettingGetResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                       `json:"usage_model"`
	JSON       scriptSettingGetResponseJSON `json:"-"`
}

// scriptSettingGetResponseJSON contains the JSON metadata for the struct
// [ScriptSettingGetResponse]
type scriptSettingGetResponseJSON struct {
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

func (r *ScriptSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseJSON) RawJSON() string {
	return r.raw
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [workers.ScriptSettingGetResponseBindingsWorkersKVNamespaceBinding],
// [workers.ScriptSettingGetResponseBindingsWorkersServiceBinding],
// [workers.ScriptSettingGetResponseBindingsWorkersDoBinding],
// [workers.ScriptSettingGetResponseBindingsWorkersR2Binding],
// [workers.ScriptSettingGetResponseBindingsWorkersQueueBinding],
// [workers.ScriptSettingGetResponseBindingsWorkersD1Binding],
// [workers.ScriptSettingGetResponseBindingsWorkersDispatchNamespaceBinding] or
// [workers.ScriptSettingGetResponseBindingsWorkersMTLSCERTBinding].
type ScriptSettingGetResponseBinding interface {
	implementsWorkersScriptSettingGetResponseBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScriptSettingGetResponseBinding)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptSettingGetResponseBindingsWorkersKVNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptSettingGetResponseBindingsWorkersServiceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptSettingGetResponseBindingsWorkersDoBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptSettingGetResponseBindingsWorkersR2Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptSettingGetResponseBindingsWorkersQueueBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptSettingGetResponseBindingsWorkersD1Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptSettingGetResponseBindingsWorkersDispatchNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptSettingGetResponseBindingsWorkersMTLSCERTBinding{}),
		},
	)
}

type ScriptSettingGetResponseBindingsWorkersKVNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type ScriptSettingGetResponseBindingsWorkersKVNamespaceBindingType `json:"type,required"`
	JSON scriptSettingGetResponseBindingsWorkersKVNamespaceBindingJSON `json:"-"`
}

// scriptSettingGetResponseBindingsWorkersKVNamespaceBindingJSON contains the JSON
// metadata for the struct
// [ScriptSettingGetResponseBindingsWorkersKVNamespaceBinding]
type scriptSettingGetResponseBindingsWorkersKVNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingGetResponseBindingsWorkersKVNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseBindingsWorkersKVNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptSettingGetResponseBindingsWorkersKVNamespaceBinding) implementsWorkersScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingGetResponseBindingsWorkersKVNamespaceBindingType string

const (
	ScriptSettingGetResponseBindingsWorkersKVNamespaceBindingTypeKVNamespace ScriptSettingGetResponseBindingsWorkersKVNamespaceBindingType = "kv_namespace"
)

func (r ScriptSettingGetResponseBindingsWorkersKVNamespaceBindingType) IsKnown() bool {
	switch r {
	case ScriptSettingGetResponseBindingsWorkersKVNamespaceBindingTypeKVNamespace:
		return true
	}
	return false
}

type ScriptSettingGetResponseBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type ScriptSettingGetResponseBindingsWorkersServiceBindingType `json:"type,required"`
	JSON scriptSettingGetResponseBindingsWorkersServiceBindingJSON `json:"-"`
}

// scriptSettingGetResponseBindingsWorkersServiceBindingJSON contains the JSON
// metadata for the struct [ScriptSettingGetResponseBindingsWorkersServiceBinding]
type scriptSettingGetResponseBindingsWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingGetResponseBindingsWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseBindingsWorkersServiceBindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptSettingGetResponseBindingsWorkersServiceBinding) implementsWorkersScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingGetResponseBindingsWorkersServiceBindingType string

const (
	ScriptSettingGetResponseBindingsWorkersServiceBindingTypeService ScriptSettingGetResponseBindingsWorkersServiceBindingType = "service"
)

func (r ScriptSettingGetResponseBindingsWorkersServiceBindingType) IsKnown() bool {
	switch r {
	case ScriptSettingGetResponseBindingsWorkersServiceBindingTypeService:
		return true
	}
	return false
}

type ScriptSettingGetResponseBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type ScriptSettingGetResponseBindingsWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                               `json:"script_name"`
	JSON       scriptSettingGetResponseBindingsWorkersDoBindingJSON `json:"-"`
}

// scriptSettingGetResponseBindingsWorkersDoBindingJSON contains the JSON metadata
// for the struct [ScriptSettingGetResponseBindingsWorkersDoBinding]
type scriptSettingGetResponseBindingsWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingGetResponseBindingsWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseBindingsWorkersDoBindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptSettingGetResponseBindingsWorkersDoBinding) implementsWorkersScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingGetResponseBindingsWorkersDoBindingType string

const (
	ScriptSettingGetResponseBindingsWorkersDoBindingTypeDurableObjectNamespace ScriptSettingGetResponseBindingsWorkersDoBindingType = "durable_object_namespace"
)

func (r ScriptSettingGetResponseBindingsWorkersDoBindingType) IsKnown() bool {
	switch r {
	case ScriptSettingGetResponseBindingsWorkersDoBindingTypeDurableObjectNamespace:
		return true
	}
	return false
}

type ScriptSettingGetResponseBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type ScriptSettingGetResponseBindingsWorkersR2BindingType `json:"type,required"`
	JSON scriptSettingGetResponseBindingsWorkersR2BindingJSON `json:"-"`
}

// scriptSettingGetResponseBindingsWorkersR2BindingJSON contains the JSON metadata
// for the struct [ScriptSettingGetResponseBindingsWorkersR2Binding]
type scriptSettingGetResponseBindingsWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingGetResponseBindingsWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseBindingsWorkersR2BindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptSettingGetResponseBindingsWorkersR2Binding) implementsWorkersScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingGetResponseBindingsWorkersR2BindingType string

const (
	ScriptSettingGetResponseBindingsWorkersR2BindingTypeR2Bucket ScriptSettingGetResponseBindingsWorkersR2BindingType = "r2_bucket"
)

func (r ScriptSettingGetResponseBindingsWorkersR2BindingType) IsKnown() bool {
	switch r {
	case ScriptSettingGetResponseBindingsWorkersR2BindingTypeR2Bucket:
		return true
	}
	return false
}

type ScriptSettingGetResponseBindingsWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type ScriptSettingGetResponseBindingsWorkersQueueBindingType `json:"type,required"`
	JSON scriptSettingGetResponseBindingsWorkersQueueBindingJSON `json:"-"`
}

// scriptSettingGetResponseBindingsWorkersQueueBindingJSON contains the JSON
// metadata for the struct [ScriptSettingGetResponseBindingsWorkersQueueBinding]
type scriptSettingGetResponseBindingsWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingGetResponseBindingsWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseBindingsWorkersQueueBindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptSettingGetResponseBindingsWorkersQueueBinding) implementsWorkersScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingGetResponseBindingsWorkersQueueBindingType string

const (
	ScriptSettingGetResponseBindingsWorkersQueueBindingTypeQueue ScriptSettingGetResponseBindingsWorkersQueueBindingType = "queue"
)

func (r ScriptSettingGetResponseBindingsWorkersQueueBindingType) IsKnown() bool {
	switch r {
	case ScriptSettingGetResponseBindingsWorkersQueueBindingTypeQueue:
		return true
	}
	return false
}

type ScriptSettingGetResponseBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type ScriptSettingGetResponseBindingsWorkersD1BindingType `json:"type,required"`
	JSON scriptSettingGetResponseBindingsWorkersD1BindingJSON `json:"-"`
}

// scriptSettingGetResponseBindingsWorkersD1BindingJSON contains the JSON metadata
// for the struct [ScriptSettingGetResponseBindingsWorkersD1Binding]
type scriptSettingGetResponseBindingsWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingGetResponseBindingsWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseBindingsWorkersD1BindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptSettingGetResponseBindingsWorkersD1Binding) implementsWorkersScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingGetResponseBindingsWorkersD1BindingType string

const (
	ScriptSettingGetResponseBindingsWorkersD1BindingTypeD1 ScriptSettingGetResponseBindingsWorkersD1BindingType = "d1"
)

func (r ScriptSettingGetResponseBindingsWorkersD1BindingType) IsKnown() bool {
	switch r {
	case ScriptSettingGetResponseBindingsWorkersD1BindingTypeD1:
		return true
	}
	return false
}

type ScriptSettingGetResponseBindingsWorkersDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type ScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound ScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     scriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON     `json:"-"`
}

// scriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON contains the
// JSON metadata for the struct
// [ScriptSettingGetResponseBindingsWorkersDispatchNamespaceBinding]
type scriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingGetResponseBindingsWorkersDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptSettingGetResponseBindingsWorkersDispatchNamespaceBinding) implementsWorkersScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingType string

const (
	ScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace ScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

func (r ScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingType) IsKnown() bool {
	switch r {
	case ScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker
type ScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker ScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   scriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// scriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [ScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound]
type scriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON) RawJSON() string {
	return r.raw
}

// Outbound worker
type ScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                            `json:"service"`
	JSON    scriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// scriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [ScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker]
type scriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type ScriptSettingGetResponseBindingsWorkersMTLSCERTBinding struct {
	Certificate interface{} `json:"certificate,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type ScriptSettingGetResponseBindingsWorkersMTLSCERTBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                     `json:"certificate_id"`
	JSON          scriptSettingGetResponseBindingsWorkersMtlscertBindingJSON `json:"-"`
}

// scriptSettingGetResponseBindingsWorkersMtlscertBindingJSON contains the JSON
// metadata for the struct [ScriptSettingGetResponseBindingsWorkersMTLSCERTBinding]
type scriptSettingGetResponseBindingsWorkersMtlscertBindingJSON struct {
	Certificate   apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScriptSettingGetResponseBindingsWorkersMTLSCERTBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseBindingsWorkersMtlscertBindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptSettingGetResponseBindingsWorkersMTLSCERTBinding) implementsWorkersScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingGetResponseBindingsWorkersMTLSCERTBindingType string

const (
	ScriptSettingGetResponseBindingsWorkersMTLSCERTBindingTypeMTLSCertificate ScriptSettingGetResponseBindingsWorkersMTLSCERTBindingType = "mtls_certificate"
)

func (r ScriptSettingGetResponseBindingsWorkersMTLSCERTBindingType) IsKnown() bool {
	switch r {
	case ScriptSettingGetResponseBindingsWorkersMTLSCERTBindingTypeMTLSCertificate:
		return true
	}
	return false
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [workers.ScriptSettingGetResponseMigrationsWorkersSingleStepMigrations] or
// [workers.ScriptSettingGetResponseMigrationsWorkersSteppedMigrations].
type ScriptSettingGetResponseMigrations interface {
	implementsWorkersScriptSettingGetResponseMigrations()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ScriptSettingGetResponseMigrations)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptSettingGetResponseMigrationsWorkersSingleStepMigrations{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ScriptSettingGetResponseMigrationsWorkersSteppedMigrations{}),
		},
	)
}

// A single set of migrations to apply.
type ScriptSettingGetResponseMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses []ScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []ScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               scriptSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON               `json:"-"`
}

// scriptSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON contains the
// JSON metadata for the struct
// [ScriptSettingGetResponseMigrationsWorkersSingleStepMigrations]
type scriptSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ScriptSettingGetResponseMigrationsWorkersSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r ScriptSettingGetResponseMigrationsWorkersSingleStepMigrations) implementsWorkersScriptSettingGetResponseMigrations() {
}

type ScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From string                                                                        `json:"from"`
	To   string                                                                        `json:"to"`
	JSON scriptSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// scriptSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [ScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass]
type scriptSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON) RawJSON() string {
	return r.raw
}

type ScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       string                                                                            `json:"from"`
	FromScript string                                                                            `json:"from_script"`
	To         string                                                                            `json:"to"`
	JSON       scriptSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// scriptSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [ScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass]
type scriptSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON) RawJSON() string {
	return r.raw
}

type ScriptSettingGetResponseMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []ScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStep `json:"steps"`
	JSON  scriptSettingGetResponseMigrationsWorkersSteppedMigrationsJSON   `json:"-"`
}

// scriptSettingGetResponseMigrationsWorkersSteppedMigrationsJSON contains the JSON
// metadata for the struct
// [ScriptSettingGetResponseMigrationsWorkersSteppedMigrations]
type scriptSettingGetResponseMigrationsWorkersSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingGetResponseMigrationsWorkersSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseMigrationsWorkersSteppedMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r ScriptSettingGetResponseMigrationsWorkersSteppedMigrations) implementsWorkersScriptSettingGetResponseMigrations() {
}

type ScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []ScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []ScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               scriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON                `json:"-"`
}

// scriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON contains the
// JSON metadata for the struct
// [ScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStep]
type scriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON) RawJSON() string {
	return r.raw
}

type ScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From string                                                                          `json:"from"`
	To   string                                                                          `json:"to"`
	JSON scriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// scriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [ScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass]
type scriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON) RawJSON() string {
	return r.raw
}

type ScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                              `json:"from"`
	FromScript string                                                                              `json:"from_script"`
	To         string                                                                              `json:"to"`
	JSON       scriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// scriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [ScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass]
type scriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON) RawJSON() string {
	return r.raw
}

type ScriptSettingGetResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode ScriptSettingGetResponsePlacementMode `json:"mode"`
	JSON scriptSettingGetResponsePlacementJSON `json:"-"`
}

// scriptSettingGetResponsePlacementJSON contains the JSON metadata for the struct
// [ScriptSettingGetResponsePlacement]
type scriptSettingGetResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingGetResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponsePlacementJSON) RawJSON() string {
	return r.raw
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type ScriptSettingGetResponsePlacementMode string

const (
	ScriptSettingGetResponsePlacementModeSmart ScriptSettingGetResponsePlacementMode = "smart"
)

func (r ScriptSettingGetResponsePlacementMode) IsKnown() bool {
	switch r {
	case ScriptSettingGetResponsePlacementModeSmart:
		return true
	}
	return false
}

// A reference to a script that will consume logs from the attached Worker.
type ScriptSettingGetResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                   `json:"namespace"`
	JSON      scriptSettingGetResponseTailConsumerJSON `json:"-"`
}

// scriptSettingGetResponseTailConsumerJSON contains the JSON metadata for the
// struct [ScriptSettingGetResponseTailConsumer]
type scriptSettingGetResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingGetResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseTailConsumerJSON) RawJSON() string {
	return r.raw
}

type ScriptSettingEditParams struct {
	// Identifier
	AccountID param.Field[string]                          `path:"account_id,required"`
	Settings  param.Field[ScriptSettingEditParamsSettings] `json:"settings"`
}

func (r ScriptSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptSettingEditParamsSettings struct {
	// List of bindings attached to this Worker
	Bindings param.Field[[]ScriptSettingEditParamsSettingsBinding] `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations param.Field[ScriptSettingEditParamsSettingsMigrations] `json:"migrations"`
	Placement  param.Field[ScriptSettingEditParamsSettingsPlacement]  `json:"placement"`
	// Tags to help you manage your Workers
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]ScriptSettingEditParamsSettingsTailConsumer] `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel param.Field[string] `json:"usage_model"`
}

func (r ScriptSettingEditParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A binding to allow the Worker to communicate with resources
//
// Satisfied by
// [workers.ScriptSettingEditParamsSettingsBindingsWorkersKVNamespaceBinding],
// [workers.ScriptSettingEditParamsSettingsBindingsWorkersServiceBinding],
// [workers.ScriptSettingEditParamsSettingsBindingsWorkersDoBinding],
// [workers.ScriptSettingEditParamsSettingsBindingsWorkersR2Binding],
// [workers.ScriptSettingEditParamsSettingsBindingsWorkersQueueBinding],
// [workers.ScriptSettingEditParamsSettingsBindingsWorkersD1Binding],
// [workers.ScriptSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBinding],
// [workers.ScriptSettingEditParamsSettingsBindingsWorkersMTLSCERTBinding].
type ScriptSettingEditParamsSettingsBinding interface {
	implementsWorkersScriptSettingEditParamsSettingsBinding()
}

type ScriptSettingEditParamsSettingsBindingsWorkersKVNamespaceBinding struct {
	// The class of resource that the binding provides.
	Type param.Field[ScriptSettingEditParamsSettingsBindingsWorkersKVNamespaceBindingType] `json:"type,required"`
}

func (r ScriptSettingEditParamsSettingsBindingsWorkersKVNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptSettingEditParamsSettingsBindingsWorkersKVNamespaceBinding) implementsWorkersScriptSettingEditParamsSettingsBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingEditParamsSettingsBindingsWorkersKVNamespaceBindingType string

const (
	ScriptSettingEditParamsSettingsBindingsWorkersKVNamespaceBindingTypeKVNamespace ScriptSettingEditParamsSettingsBindingsWorkersKVNamespaceBindingType = "kv_namespace"
)

func (r ScriptSettingEditParamsSettingsBindingsWorkersKVNamespaceBindingType) IsKnown() bool {
	switch r {
	case ScriptSettingEditParamsSettingsBindingsWorkersKVNamespaceBindingTypeKVNamespace:
		return true
	}
	return false
}

type ScriptSettingEditParamsSettingsBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// Name of Worker to bind to
	Service param.Field[string] `json:"service,required"`
	// The class of resource that the binding provides.
	Type param.Field[ScriptSettingEditParamsSettingsBindingsWorkersServiceBindingType] `json:"type,required"`
}

func (r ScriptSettingEditParamsSettingsBindingsWorkersServiceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptSettingEditParamsSettingsBindingsWorkersServiceBinding) implementsWorkersScriptSettingEditParamsSettingsBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingEditParamsSettingsBindingsWorkersServiceBindingType string

const (
	ScriptSettingEditParamsSettingsBindingsWorkersServiceBindingTypeService ScriptSettingEditParamsSettingsBindingsWorkersServiceBindingType = "service"
)

func (r ScriptSettingEditParamsSettingsBindingsWorkersServiceBindingType) IsKnown() bool {
	switch r {
	case ScriptSettingEditParamsSettingsBindingsWorkersServiceBindingTypeService:
		return true
	}
	return false
}

type ScriptSettingEditParamsSettingsBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName param.Field[string] `json:"class_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[ScriptSettingEditParamsSettingsBindingsWorkersDoBindingType] `json:"type,required"`
	// The environment of the script_name to bind to
	Environment param.Field[string] `json:"environment"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName param.Field[string] `json:"script_name"`
}

func (r ScriptSettingEditParamsSettingsBindingsWorkersDoBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptSettingEditParamsSettingsBindingsWorkersDoBinding) implementsWorkersScriptSettingEditParamsSettingsBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingEditParamsSettingsBindingsWorkersDoBindingType string

const (
	ScriptSettingEditParamsSettingsBindingsWorkersDoBindingTypeDurableObjectNamespace ScriptSettingEditParamsSettingsBindingsWorkersDoBindingType = "durable_object_namespace"
)

func (r ScriptSettingEditParamsSettingsBindingsWorkersDoBindingType) IsKnown() bool {
	switch r {
	case ScriptSettingEditParamsSettingsBindingsWorkersDoBindingTypeDurableObjectNamespace:
		return true
	}
	return false
}

type ScriptSettingEditParamsSettingsBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName param.Field[string] `json:"bucket_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[ScriptSettingEditParamsSettingsBindingsWorkersR2BindingType] `json:"type,required"`
}

func (r ScriptSettingEditParamsSettingsBindingsWorkersR2Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptSettingEditParamsSettingsBindingsWorkersR2Binding) implementsWorkersScriptSettingEditParamsSettingsBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingEditParamsSettingsBindingsWorkersR2BindingType string

const (
	ScriptSettingEditParamsSettingsBindingsWorkersR2BindingTypeR2Bucket ScriptSettingEditParamsSettingsBindingsWorkersR2BindingType = "r2_bucket"
)

func (r ScriptSettingEditParamsSettingsBindingsWorkersR2BindingType) IsKnown() bool {
	switch r {
	case ScriptSettingEditParamsSettingsBindingsWorkersR2BindingTypeR2Bucket:
		return true
	}
	return false
}

type ScriptSettingEditParamsSettingsBindingsWorkersQueueBinding struct {
	// Name of the Queue to bind to
	QueueName param.Field[string] `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[ScriptSettingEditParamsSettingsBindingsWorkersQueueBindingType] `json:"type,required"`
}

func (r ScriptSettingEditParamsSettingsBindingsWorkersQueueBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptSettingEditParamsSettingsBindingsWorkersQueueBinding) implementsWorkersScriptSettingEditParamsSettingsBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingEditParamsSettingsBindingsWorkersQueueBindingType string

const (
	ScriptSettingEditParamsSettingsBindingsWorkersQueueBindingTypeQueue ScriptSettingEditParamsSettingsBindingsWorkersQueueBindingType = "queue"
)

func (r ScriptSettingEditParamsSettingsBindingsWorkersQueueBindingType) IsKnown() bool {
	switch r {
	case ScriptSettingEditParamsSettingsBindingsWorkersQueueBindingTypeQueue:
		return true
	}
	return false
}

type ScriptSettingEditParamsSettingsBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID param.Field[string] `json:"id,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name param.Field[string] `json:"name,required"`
	// The class of resource that the binding provides.
	Type param.Field[ScriptSettingEditParamsSettingsBindingsWorkersD1BindingType] `json:"type,required"`
}

func (r ScriptSettingEditParamsSettingsBindingsWorkersD1Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptSettingEditParamsSettingsBindingsWorkersD1Binding) implementsWorkersScriptSettingEditParamsSettingsBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingEditParamsSettingsBindingsWorkersD1BindingType string

const (
	ScriptSettingEditParamsSettingsBindingsWorkersD1BindingTypeD1 ScriptSettingEditParamsSettingsBindingsWorkersD1BindingType = "d1"
)

func (r ScriptSettingEditParamsSettingsBindingsWorkersD1BindingType) IsKnown() bool {
	switch r {
	case ScriptSettingEditParamsSettingsBindingsWorkersD1BindingTypeD1:
		return true
	}
	return false
}

type ScriptSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBinding struct {
	// Namespace to bind to
	Namespace param.Field[string] `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type param.Field[ScriptSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBindingType] `json:"type,required"`
	// Outbound worker
	Outbound param.Field[ScriptSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBindingOutbound] `json:"outbound"`
}

func (r ScriptSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBinding) implementsWorkersScriptSettingEditParamsSettingsBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBindingType string

const (
	ScriptSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace ScriptSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

func (r ScriptSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBindingType) IsKnown() bool {
	switch r {
	case ScriptSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker
type ScriptSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params param.Field[[]string] `json:"params"`
	// Outbound worker
	Worker param.Field[ScriptSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBindingOutboundWorker] `json:"worker"`
}

func (r ScriptSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBindingOutbound) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker
type ScriptSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker
	Service param.Field[string] `json:"service"`
}

func (r ScriptSettingEditParamsSettingsBindingsWorkersDispatchNamespaceBindingOutboundWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptSettingEditParamsSettingsBindingsWorkersMTLSCERTBinding struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
	// The class of resource that the binding provides.
	Type param.Field[ScriptSettingEditParamsSettingsBindingsWorkersMTLSCERTBindingType] `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID param.Field[string] `json:"certificate_id"`
}

func (r ScriptSettingEditParamsSettingsBindingsWorkersMTLSCERTBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptSettingEditParamsSettingsBindingsWorkersMTLSCERTBinding) implementsWorkersScriptSettingEditParamsSettingsBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingEditParamsSettingsBindingsWorkersMTLSCERTBindingType string

const (
	ScriptSettingEditParamsSettingsBindingsWorkersMTLSCERTBindingTypeMTLSCertificate ScriptSettingEditParamsSettingsBindingsWorkersMTLSCERTBindingType = "mtls_certificate"
)

func (r ScriptSettingEditParamsSettingsBindingsWorkersMTLSCERTBindingType) IsKnown() bool {
	switch r {
	case ScriptSettingEditParamsSettingsBindingsWorkersMTLSCERTBindingTypeMTLSCertificate:
		return true
	}
	return false
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by
// [workers.ScriptSettingEditParamsSettingsMigrationsWorkersSingleStepMigrations],
// [workers.ScriptSettingEditParamsSettingsMigrationsWorkersSteppedMigrations].
type ScriptSettingEditParamsSettingsMigrations interface {
	implementsWorkersScriptSettingEditParamsSettingsMigrations()
}

// A single set of migrations to apply.
type ScriptSettingEditParamsSettingsMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses param.Field[[]ScriptSettingEditParamsSettingsMigrationsWorkersSingleStepMigrationsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]ScriptSettingEditParamsSettingsMigrationsWorkersSingleStepMigrationsTransferredClass] `json:"transferred_classes"`
}

func (r ScriptSettingEditParamsSettingsMigrationsWorkersSingleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptSettingEditParamsSettingsMigrationsWorkersSingleStepMigrations) implementsWorkersScriptSettingEditParamsSettingsMigrations() {
}

type ScriptSettingEditParamsSettingsMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r ScriptSettingEditParamsSettingsMigrationsWorkersSingleStepMigrationsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptSettingEditParamsSettingsMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r ScriptSettingEditParamsSettingsMigrationsWorkersSingleStepMigrationsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptSettingEditParamsSettingsMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]ScriptSettingEditParamsSettingsMigrationsWorkersSteppedMigrationsStep] `json:"steps"`
}

func (r ScriptSettingEditParamsSettingsMigrationsWorkersSteppedMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptSettingEditParamsSettingsMigrationsWorkersSteppedMigrations) implementsWorkersScriptSettingEditParamsSettingsMigrations() {
}

type ScriptSettingEditParamsSettingsMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]ScriptSettingEditParamsSettingsMigrationsWorkersSteppedMigrationsStepsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]ScriptSettingEditParamsSettingsMigrationsWorkersSteppedMigrationsStepsTransferredClass] `json:"transferred_classes"`
}

func (r ScriptSettingEditParamsSettingsMigrationsWorkersSteppedMigrationsStep) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptSettingEditParamsSettingsMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r ScriptSettingEditParamsSettingsMigrationsWorkersSteppedMigrationsStepsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptSettingEditParamsSettingsMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r ScriptSettingEditParamsSettingsMigrationsWorkersSteppedMigrationsStepsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptSettingEditParamsSettingsPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode param.Field[ScriptSettingEditParamsSettingsPlacementMode] `json:"mode"`
}

func (r ScriptSettingEditParamsSettingsPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type ScriptSettingEditParamsSettingsPlacementMode string

const (
	ScriptSettingEditParamsSettingsPlacementModeSmart ScriptSettingEditParamsSettingsPlacementMode = "smart"
)

func (r ScriptSettingEditParamsSettingsPlacementMode) IsKnown() bool {
	switch r {
	case ScriptSettingEditParamsSettingsPlacementModeSmart:
		return true
	}
	return false
}

// A reference to a script that will consume logs from the attached Worker.
type ScriptSettingEditParamsSettingsTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service param.Field[string] `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace param.Field[string] `json:"namespace"`
}

func (r ScriptSettingEditParamsSettingsTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptSettingEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	Result   ScriptSettingEditResponse    `json:"result,required"`
	// Whether the API call was successful
	Success ScriptSettingEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptSettingEditResponseEnvelopeJSON    `json:"-"`
}

// scriptSettingEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptSettingEditResponseEnvelope]
type scriptSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptSettingEditResponseEnvelopeSuccess bool

const (
	ScriptSettingEditResponseEnvelopeSuccessTrue ScriptSettingEditResponseEnvelopeSuccess = true
)

func (r ScriptSettingEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptSettingEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptSettingGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptSettingGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	Result   ScriptSettingGetResponse     `json:"result,required"`
	// Whether the API call was successful
	Success ScriptSettingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptSettingGetResponseEnvelopeJSON    `json:"-"`
}

// scriptSettingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScriptSettingGetResponseEnvelope]
type scriptSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptSettingGetResponseEnvelopeSuccess bool

const (
	ScriptSettingGetResponseEnvelopeSuccessTrue ScriptSettingGetResponseEnvelopeSuccess = true
)

func (r ScriptSettingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptSettingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
