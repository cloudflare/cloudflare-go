// File generated from our OpenAPI spec by Stainless.

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

// Patch script metadata or config, such as bindings or usage model
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

// Get script metadata and config, such as bindings or usage model
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
// [workers.ScriptSettingEditResponseBindingsWorkersMTLSCertBinding].
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
			Type:       reflect.TypeOf(ScriptSettingEditResponseBindingsWorkersMTLSCertBinding{}),
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

type ScriptSettingEditResponseBindingsWorkersMTLSCertBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type ScriptSettingEditResponseBindingsWorkersMTLSCertBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                      `json:"certificate_id"`
	JSON          scriptSettingEditResponseBindingsWorkersMTLSCertBindingJSON `json:"-"`
}

// scriptSettingEditResponseBindingsWorkersMTLSCertBindingJSON contains the JSON
// metadata for the struct
// [ScriptSettingEditResponseBindingsWorkersMTLSCertBinding]
type scriptSettingEditResponseBindingsWorkersMTLSCertBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScriptSettingEditResponseBindingsWorkersMTLSCertBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseBindingsWorkersMTLSCertBindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptSettingEditResponseBindingsWorkersMTLSCertBinding) implementsWorkersScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingEditResponseBindingsWorkersMTLSCertBindingType string

const (
	ScriptSettingEditResponseBindingsWorkersMTLSCertBindingTypeMTLSCertificate ScriptSettingEditResponseBindingsWorkersMTLSCertBindingType = "mtls_certificate"
)

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
// [workers.ScriptSettingGetResponseBindingsWorkersMTLSCertBinding].
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
			Type:       reflect.TypeOf(ScriptSettingGetResponseBindingsWorkersMTLSCertBinding{}),
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

type ScriptSettingGetResponseBindingsWorkersMTLSCertBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type ScriptSettingGetResponseBindingsWorkersMTLSCertBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                     `json:"certificate_id"`
	JSON          scriptSettingGetResponseBindingsWorkersMTLSCertBindingJSON `json:"-"`
}

// scriptSettingGetResponseBindingsWorkersMTLSCertBindingJSON contains the JSON
// metadata for the struct [ScriptSettingGetResponseBindingsWorkersMTLSCertBinding]
type scriptSettingGetResponseBindingsWorkersMTLSCertBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScriptSettingGetResponseBindingsWorkersMTLSCertBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseBindingsWorkersMTLSCertBindingJSON) RawJSON() string {
	return r.raw
}

func (r ScriptSettingGetResponseBindingsWorkersMTLSCertBinding) implementsWorkersScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingGetResponseBindingsWorkersMTLSCertBindingType string

const (
	ScriptSettingGetResponseBindingsWorkersMTLSCertBindingTypeMTLSCertificate ScriptSettingGetResponseBindingsWorkersMTLSCertBindingType = "mtls_certificate"
)

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
	Errors   param.Field[[]ScriptSettingEditParamsSettingsError]   `json:"errors,required"`
	Messages param.Field[[]ScriptSettingEditParamsSettingsMessage] `json:"messages,required"`
	Result   param.Field[ScriptSettingEditParamsSettingsResult]    `json:"result,required"`
	// Whether the API call was successful
	Success param.Field[ScriptSettingEditParamsSettingsSuccess] `json:"success,required"`
}

func (r ScriptSettingEditParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptSettingEditParamsSettingsError struct {
	Code    param.Field[int64]  `json:"code,required"`
	Message param.Field[string] `json:"message,required"`
}

func (r ScriptSettingEditParamsSettingsError) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptSettingEditParamsSettingsMessage struct {
	Code    param.Field[int64]  `json:"code,required"`
	Message param.Field[string] `json:"message,required"`
}

func (r ScriptSettingEditParamsSettingsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptSettingEditParamsSettingsResult struct {
	// List of bindings attached to this Worker
	Bindings param.Field[[]ScriptSettingEditParamsSettingsResultBinding] `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations param.Field[ScriptSettingEditParamsSettingsResultMigrations] `json:"migrations"`
	Placement  param.Field[ScriptSettingEditParamsSettingsResultPlacement]  `json:"placement"`
	// Tags to help you manage your Workers
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]ScriptSettingEditParamsSettingsResultTailConsumer] `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel param.Field[string] `json:"usage_model"`
}

func (r ScriptSettingEditParamsSettingsResult) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A binding to allow the Worker to communicate with resources
//
// Satisfied by
// [workers.ScriptSettingEditParamsSettingsResultBindingsWorkersKVNamespaceBinding],
// [workers.ScriptSettingEditParamsSettingsResultBindingsWorkersServiceBinding],
// [workers.ScriptSettingEditParamsSettingsResultBindingsWorkersDoBinding],
// [workers.ScriptSettingEditParamsSettingsResultBindingsWorkersR2Binding],
// [workers.ScriptSettingEditParamsSettingsResultBindingsWorkersQueueBinding],
// [workers.ScriptSettingEditParamsSettingsResultBindingsWorkersD1Binding],
// [workers.ScriptSettingEditParamsSettingsResultBindingsWorkersDispatchNamespaceBinding],
// [workers.ScriptSettingEditParamsSettingsResultBindingsWorkersMTLSCertBinding].
type ScriptSettingEditParamsSettingsResultBinding interface {
	implementsWorkersScriptSettingEditParamsSettingsResultBinding()
}

type ScriptSettingEditParamsSettingsResultBindingsWorkersKVNamespaceBinding struct {
	// The class of resource that the binding provides.
	Type param.Field[ScriptSettingEditParamsSettingsResultBindingsWorkersKVNamespaceBindingType] `json:"type,required"`
}

func (r ScriptSettingEditParamsSettingsResultBindingsWorkersKVNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptSettingEditParamsSettingsResultBindingsWorkersKVNamespaceBinding) implementsWorkersScriptSettingEditParamsSettingsResultBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingEditParamsSettingsResultBindingsWorkersKVNamespaceBindingType string

const (
	ScriptSettingEditParamsSettingsResultBindingsWorkersKVNamespaceBindingTypeKVNamespace ScriptSettingEditParamsSettingsResultBindingsWorkersKVNamespaceBindingType = "kv_namespace"
)

type ScriptSettingEditParamsSettingsResultBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// Name of Worker to bind to
	Service param.Field[string] `json:"service,required"`
	// The class of resource that the binding provides.
	Type param.Field[ScriptSettingEditParamsSettingsResultBindingsWorkersServiceBindingType] `json:"type,required"`
}

func (r ScriptSettingEditParamsSettingsResultBindingsWorkersServiceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptSettingEditParamsSettingsResultBindingsWorkersServiceBinding) implementsWorkersScriptSettingEditParamsSettingsResultBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingEditParamsSettingsResultBindingsWorkersServiceBindingType string

const (
	ScriptSettingEditParamsSettingsResultBindingsWorkersServiceBindingTypeService ScriptSettingEditParamsSettingsResultBindingsWorkersServiceBindingType = "service"
)

type ScriptSettingEditParamsSettingsResultBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName param.Field[string] `json:"class_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[ScriptSettingEditParamsSettingsResultBindingsWorkersDoBindingType] `json:"type,required"`
	// The environment of the script_name to bind to
	Environment param.Field[string] `json:"environment"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName param.Field[string] `json:"script_name"`
}

func (r ScriptSettingEditParamsSettingsResultBindingsWorkersDoBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptSettingEditParamsSettingsResultBindingsWorkersDoBinding) implementsWorkersScriptSettingEditParamsSettingsResultBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingEditParamsSettingsResultBindingsWorkersDoBindingType string

const (
	ScriptSettingEditParamsSettingsResultBindingsWorkersDoBindingTypeDurableObjectNamespace ScriptSettingEditParamsSettingsResultBindingsWorkersDoBindingType = "durable_object_namespace"
)

type ScriptSettingEditParamsSettingsResultBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName param.Field[string] `json:"bucket_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[ScriptSettingEditParamsSettingsResultBindingsWorkersR2BindingType] `json:"type,required"`
}

func (r ScriptSettingEditParamsSettingsResultBindingsWorkersR2Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptSettingEditParamsSettingsResultBindingsWorkersR2Binding) implementsWorkersScriptSettingEditParamsSettingsResultBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingEditParamsSettingsResultBindingsWorkersR2BindingType string

const (
	ScriptSettingEditParamsSettingsResultBindingsWorkersR2BindingTypeR2Bucket ScriptSettingEditParamsSettingsResultBindingsWorkersR2BindingType = "r2_bucket"
)

type ScriptSettingEditParamsSettingsResultBindingsWorkersQueueBinding struct {
	// Name of the Queue to bind to
	QueueName param.Field[string] `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[ScriptSettingEditParamsSettingsResultBindingsWorkersQueueBindingType] `json:"type,required"`
}

func (r ScriptSettingEditParamsSettingsResultBindingsWorkersQueueBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptSettingEditParamsSettingsResultBindingsWorkersQueueBinding) implementsWorkersScriptSettingEditParamsSettingsResultBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingEditParamsSettingsResultBindingsWorkersQueueBindingType string

const (
	ScriptSettingEditParamsSettingsResultBindingsWorkersQueueBindingTypeQueue ScriptSettingEditParamsSettingsResultBindingsWorkersQueueBindingType = "queue"
)

type ScriptSettingEditParamsSettingsResultBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID param.Field[string] `json:"id,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name param.Field[string] `json:"name,required"`
	// The class of resource that the binding provides.
	Type param.Field[ScriptSettingEditParamsSettingsResultBindingsWorkersD1BindingType] `json:"type,required"`
}

func (r ScriptSettingEditParamsSettingsResultBindingsWorkersD1Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptSettingEditParamsSettingsResultBindingsWorkersD1Binding) implementsWorkersScriptSettingEditParamsSettingsResultBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingEditParamsSettingsResultBindingsWorkersD1BindingType string

const (
	ScriptSettingEditParamsSettingsResultBindingsWorkersD1BindingTypeD1 ScriptSettingEditParamsSettingsResultBindingsWorkersD1BindingType = "d1"
)

type ScriptSettingEditParamsSettingsResultBindingsWorkersDispatchNamespaceBinding struct {
	// Namespace to bind to
	Namespace param.Field[string] `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type param.Field[ScriptSettingEditParamsSettingsResultBindingsWorkersDispatchNamespaceBindingType] `json:"type,required"`
	// Outbound worker
	Outbound param.Field[ScriptSettingEditParamsSettingsResultBindingsWorkersDispatchNamespaceBindingOutbound] `json:"outbound"`
}

func (r ScriptSettingEditParamsSettingsResultBindingsWorkersDispatchNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptSettingEditParamsSettingsResultBindingsWorkersDispatchNamespaceBinding) implementsWorkersScriptSettingEditParamsSettingsResultBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingEditParamsSettingsResultBindingsWorkersDispatchNamespaceBindingType string

const (
	ScriptSettingEditParamsSettingsResultBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace ScriptSettingEditParamsSettingsResultBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type ScriptSettingEditParamsSettingsResultBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params param.Field[[]string] `json:"params"`
	// Outbound worker
	Worker param.Field[ScriptSettingEditParamsSettingsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker] `json:"worker"`
}

func (r ScriptSettingEditParamsSettingsResultBindingsWorkersDispatchNamespaceBindingOutbound) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker
type ScriptSettingEditParamsSettingsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker
	Service param.Field[string] `json:"service"`
}

func (r ScriptSettingEditParamsSettingsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptSettingEditParamsSettingsResultBindingsWorkersMTLSCertBinding struct {
	// The class of resource that the binding provides.
	Type param.Field[ScriptSettingEditParamsSettingsResultBindingsWorkersMTLSCertBindingType] `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID param.Field[string] `json:"certificate_id"`
}

func (r ScriptSettingEditParamsSettingsResultBindingsWorkersMTLSCertBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptSettingEditParamsSettingsResultBindingsWorkersMTLSCertBinding) implementsWorkersScriptSettingEditParamsSettingsResultBinding() {
}

// The class of resource that the binding provides.
type ScriptSettingEditParamsSettingsResultBindingsWorkersMTLSCertBindingType string

const (
	ScriptSettingEditParamsSettingsResultBindingsWorkersMTLSCertBindingTypeMTLSCertificate ScriptSettingEditParamsSettingsResultBindingsWorkersMTLSCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by
// [workers.ScriptSettingEditParamsSettingsResultMigrationsWorkersSingleStepMigrations],
// [workers.ScriptSettingEditParamsSettingsResultMigrationsWorkersSteppedMigrations].
type ScriptSettingEditParamsSettingsResultMigrations interface {
	implementsWorkersScriptSettingEditParamsSettingsResultMigrations()
}

// A single set of migrations to apply.
type ScriptSettingEditParamsSettingsResultMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses param.Field[[]ScriptSettingEditParamsSettingsResultMigrationsWorkersSingleStepMigrationsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]ScriptSettingEditParamsSettingsResultMigrationsWorkersSingleStepMigrationsTransferredClass] `json:"transferred_classes"`
}

func (r ScriptSettingEditParamsSettingsResultMigrationsWorkersSingleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptSettingEditParamsSettingsResultMigrationsWorkersSingleStepMigrations) implementsWorkersScriptSettingEditParamsSettingsResultMigrations() {
}

type ScriptSettingEditParamsSettingsResultMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r ScriptSettingEditParamsSettingsResultMigrationsWorkersSingleStepMigrationsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptSettingEditParamsSettingsResultMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r ScriptSettingEditParamsSettingsResultMigrationsWorkersSingleStepMigrationsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptSettingEditParamsSettingsResultMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]ScriptSettingEditParamsSettingsResultMigrationsWorkersSteppedMigrationsStep] `json:"steps"`
}

func (r ScriptSettingEditParamsSettingsResultMigrationsWorkersSteppedMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScriptSettingEditParamsSettingsResultMigrationsWorkersSteppedMigrations) implementsWorkersScriptSettingEditParamsSettingsResultMigrations() {
}

type ScriptSettingEditParamsSettingsResultMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]ScriptSettingEditParamsSettingsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]ScriptSettingEditParamsSettingsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass] `json:"transferred_classes"`
}

func (r ScriptSettingEditParamsSettingsResultMigrationsWorkersSteppedMigrationsStep) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptSettingEditParamsSettingsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r ScriptSettingEditParamsSettingsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptSettingEditParamsSettingsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r ScriptSettingEditParamsSettingsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptSettingEditParamsSettingsResultPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode param.Field[ScriptSettingEditParamsSettingsResultPlacementMode] `json:"mode"`
}

func (r ScriptSettingEditParamsSettingsResultPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type ScriptSettingEditParamsSettingsResultPlacementMode string

const (
	ScriptSettingEditParamsSettingsResultPlacementModeSmart ScriptSettingEditParamsSettingsResultPlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type ScriptSettingEditParamsSettingsResultTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service param.Field[string] `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace param.Field[string] `json:"namespace"`
}

func (r ScriptSettingEditParamsSettingsResultTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the API call was successful
type ScriptSettingEditParamsSettingsSuccess bool

const (
	ScriptSettingEditParamsSettingsSuccessTrue ScriptSettingEditParamsSettingsSuccess = true
)

type ScriptSettingEditResponseEnvelope struct {
	Errors   []ScriptSettingEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptSettingEditResponseEnvelopeMessages `json:"messages,required"`
	Result   ScriptSettingEditResponse                   `json:"result,required"`
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

type ScriptSettingEditResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    scriptSettingEditResponseEnvelopeErrorsJSON `json:"-"`
}

// scriptSettingEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ScriptSettingEditResponseEnvelopeErrors]
type scriptSettingEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptSettingEditResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    scriptSettingEditResponseEnvelopeMessagesJSON `json:"-"`
}

// scriptSettingEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ScriptSettingEditResponseEnvelopeMessages]
type scriptSettingEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptSettingEditResponseEnvelopeSuccess bool

const (
	ScriptSettingEditResponseEnvelopeSuccessTrue ScriptSettingEditResponseEnvelopeSuccess = true
)

type ScriptSettingGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptSettingGetResponseEnvelope struct {
	Errors   []ScriptSettingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScriptSettingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ScriptSettingGetResponse                   `json:"result,required"`
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

type ScriptSettingGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    scriptSettingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// scriptSettingGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ScriptSettingGetResponseEnvelopeErrors]
type scriptSettingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScriptSettingGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    scriptSettingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// scriptSettingGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ScriptSettingGetResponseEnvelopeMessages]
type scriptSettingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptSettingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptSettingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptSettingGetResponseEnvelopeSuccess bool

const (
	ScriptSettingGetResponseEnvelopeSuccessTrue ScriptSettingGetResponseEnvelopeSuccess = true
)
