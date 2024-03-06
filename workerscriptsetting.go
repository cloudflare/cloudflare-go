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

// WorkerScriptSettingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWorkerScriptSettingService]
// method instead.
type WorkerScriptSettingService struct {
	Options []option.RequestOption
}

// NewWorkerScriptSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkerScriptSettingService(opts ...option.RequestOption) (r *WorkerScriptSettingService) {
	r = &WorkerScriptSettingService{}
	r.Options = opts
	return
}

// Patch script metadata or config, such as bindings or usage model
func (r *WorkerScriptSettingService) Edit(ctx context.Context, scriptName string, params WorkerScriptSettingEditParams, opts ...option.RequestOption) (res *WorkerScriptSettingEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerScriptSettingEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/settings", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get script metadata and config, such as bindings or usage model
func (r *WorkerScriptSettingService) Get(ctx context.Context, scriptName string, query WorkerScriptSettingGetParams, opts ...option.RequestOption) (res *WorkerScriptSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerScriptSettingGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/settings", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkerScriptSettingEditResponse struct {
	// List of bindings attached to this Worker
	Bindings []WorkerScriptSettingEditResponseBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations WorkerScriptSettingEditResponseMigrations `json:"migrations"`
	Placement  WorkerScriptSettingEditResponsePlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []WorkerScriptSettingEditResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                              `json:"usage_model"`
	JSON       workerScriptSettingEditResponseJSON `json:"-"`
}

// workerScriptSettingEditResponseJSON contains the JSON metadata for the struct
// [WorkerScriptSettingEditResponse]
type workerScriptSettingEditResponseJSON struct {
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

func (r *WorkerScriptSettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingEditResponseJSON) RawJSON() string {
	return r.raw
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [WorkerScriptSettingEditResponseBindingsWorkersKVNamespaceBinding],
// [WorkerScriptSettingEditResponseBindingsWorkersServiceBinding],
// [WorkerScriptSettingEditResponseBindingsWorkersDoBinding],
// [WorkerScriptSettingEditResponseBindingsWorkersR2Binding],
// [WorkerScriptSettingEditResponseBindingsWorkersQueueBinding],
// [WorkerScriptSettingEditResponseBindingsWorkersD1Binding],
// [WorkerScriptSettingEditResponseBindingsWorkersDispatchNamespaceBinding] or
// [WorkerScriptSettingEditResponseBindingsWorkersMTLSCertBinding].
type WorkerScriptSettingEditResponseBinding interface {
	implementsWorkerScriptSettingEditResponseBinding()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkerScriptSettingEditResponseBinding)(nil)).Elem(), "")
}

type WorkerScriptSettingEditResponseBindingsWorkersKVNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingEditResponseBindingsWorkersKVNamespaceBindingType `json:"type,required"`
	JSON workerScriptSettingEditResponseBindingsWorkersKVNamespaceBindingJSON `json:"-"`
}

// workerScriptSettingEditResponseBindingsWorkersKVNamespaceBindingJSON contains
// the JSON metadata for the struct
// [WorkerScriptSettingEditResponseBindingsWorkersKVNamespaceBinding]
type workerScriptSettingEditResponseBindingsWorkersKVNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingEditResponseBindingsWorkersKVNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingEditResponseBindingsWorkersKVNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkerScriptSettingEditResponseBindingsWorkersKVNamespaceBinding) implementsWorkerScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingEditResponseBindingsWorkersKVNamespaceBindingType string

const (
	WorkerScriptSettingEditResponseBindingsWorkersKVNamespaceBindingTypeKVNamespace WorkerScriptSettingEditResponseBindingsWorkersKVNamespaceBindingType = "kv_namespace"
)

type WorkerScriptSettingEditResponseBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingEditResponseBindingsWorkersServiceBindingType `json:"type,required"`
	JSON workerScriptSettingEditResponseBindingsWorkersServiceBindingJSON `json:"-"`
}

// workerScriptSettingEditResponseBindingsWorkersServiceBindingJSON contains the
// JSON metadata for the struct
// [WorkerScriptSettingEditResponseBindingsWorkersServiceBinding]
type workerScriptSettingEditResponseBindingsWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingEditResponseBindingsWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingEditResponseBindingsWorkersServiceBindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkerScriptSettingEditResponseBindingsWorkersServiceBinding) implementsWorkerScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingEditResponseBindingsWorkersServiceBindingType string

const (
	WorkerScriptSettingEditResponseBindingsWorkersServiceBindingTypeService WorkerScriptSettingEditResponseBindingsWorkersServiceBindingType = "service"
)

type WorkerScriptSettingEditResponseBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingEditResponseBindingsWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                      `json:"script_name"`
	JSON       workerScriptSettingEditResponseBindingsWorkersDoBindingJSON `json:"-"`
}

// workerScriptSettingEditResponseBindingsWorkersDoBindingJSON contains the JSON
// metadata for the struct
// [WorkerScriptSettingEditResponseBindingsWorkersDoBinding]
type workerScriptSettingEditResponseBindingsWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingEditResponseBindingsWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingEditResponseBindingsWorkersDoBindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkerScriptSettingEditResponseBindingsWorkersDoBinding) implementsWorkerScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingEditResponseBindingsWorkersDoBindingType string

const (
	WorkerScriptSettingEditResponseBindingsWorkersDoBindingTypeDurableObjectNamespace WorkerScriptSettingEditResponseBindingsWorkersDoBindingType = "durable_object_namespace"
)

type WorkerScriptSettingEditResponseBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingEditResponseBindingsWorkersR2BindingType `json:"type,required"`
	JSON workerScriptSettingEditResponseBindingsWorkersR2BindingJSON `json:"-"`
}

// workerScriptSettingEditResponseBindingsWorkersR2BindingJSON contains the JSON
// metadata for the struct
// [WorkerScriptSettingEditResponseBindingsWorkersR2Binding]
type workerScriptSettingEditResponseBindingsWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingEditResponseBindingsWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingEditResponseBindingsWorkersR2BindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkerScriptSettingEditResponseBindingsWorkersR2Binding) implementsWorkerScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingEditResponseBindingsWorkersR2BindingType string

const (
	WorkerScriptSettingEditResponseBindingsWorkersR2BindingTypeR2Bucket WorkerScriptSettingEditResponseBindingsWorkersR2BindingType = "r2_bucket"
)

type WorkerScriptSettingEditResponseBindingsWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingEditResponseBindingsWorkersQueueBindingType `json:"type,required"`
	JSON workerScriptSettingEditResponseBindingsWorkersQueueBindingJSON `json:"-"`
}

// workerScriptSettingEditResponseBindingsWorkersQueueBindingJSON contains the JSON
// metadata for the struct
// [WorkerScriptSettingEditResponseBindingsWorkersQueueBinding]
type workerScriptSettingEditResponseBindingsWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingEditResponseBindingsWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingEditResponseBindingsWorkersQueueBindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkerScriptSettingEditResponseBindingsWorkersQueueBinding) implementsWorkerScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingEditResponseBindingsWorkersQueueBindingType string

const (
	WorkerScriptSettingEditResponseBindingsWorkersQueueBindingTypeQueue WorkerScriptSettingEditResponseBindingsWorkersQueueBindingType = "queue"
)

type WorkerScriptSettingEditResponseBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingEditResponseBindingsWorkersD1BindingType `json:"type,required"`
	JSON workerScriptSettingEditResponseBindingsWorkersD1BindingJSON `json:"-"`
}

// workerScriptSettingEditResponseBindingsWorkersD1BindingJSON contains the JSON
// metadata for the struct
// [WorkerScriptSettingEditResponseBindingsWorkersD1Binding]
type workerScriptSettingEditResponseBindingsWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingEditResponseBindingsWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingEditResponseBindingsWorkersD1BindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkerScriptSettingEditResponseBindingsWorkersD1Binding) implementsWorkerScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingEditResponseBindingsWorkersD1BindingType string

const (
	WorkerScriptSettingEditResponseBindingsWorkersD1BindingTypeD1 WorkerScriptSettingEditResponseBindingsWorkersD1BindingType = "d1"
)

type WorkerScriptSettingEditResponseBindingsWorkersDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound WorkerScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     workerScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON     `json:"-"`
}

// workerScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingEditResponseBindingsWorkersDispatchNamespaceBinding]
type workerScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingEditResponseBindingsWorkersDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkerScriptSettingEditResponseBindingsWorkersDispatchNamespaceBinding) implementsWorkerScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingType string

const (
	WorkerScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace WorkerScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type WorkerScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker WorkerScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   workerScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// workerScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound]
type workerScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON) RawJSON() string {
	return r.raw
}

// Outbound worker
type WorkerScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                   `json:"service"`
	JSON    workerScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// workerScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker]
type workerScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptSettingEditResponseBindingsWorkersMTLSCertBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingEditResponseBindingsWorkersMTLSCertBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                            `json:"certificate_id"`
	JSON          workerScriptSettingEditResponseBindingsWorkersMTLSCertBindingJSON `json:"-"`
}

// workerScriptSettingEditResponseBindingsWorkersMTLSCertBindingJSON contains the
// JSON metadata for the struct
// [WorkerScriptSettingEditResponseBindingsWorkersMTLSCertBinding]
type workerScriptSettingEditResponseBindingsWorkersMTLSCertBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkerScriptSettingEditResponseBindingsWorkersMTLSCertBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingEditResponseBindingsWorkersMTLSCertBindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkerScriptSettingEditResponseBindingsWorkersMTLSCertBinding) implementsWorkerScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingEditResponseBindingsWorkersMTLSCertBindingType string

const (
	WorkerScriptSettingEditResponseBindingsWorkersMTLSCertBindingTypeMTLSCertificate WorkerScriptSettingEditResponseBindingsWorkersMTLSCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [WorkerScriptSettingEditResponseMigrationsWorkersSingleStepMigrations] or
// [WorkerScriptSettingEditResponseMigrationsWorkersSteppedMigrations].
type WorkerScriptSettingEditResponseMigrations interface {
	implementsWorkerScriptSettingEditResponseMigrations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkerScriptSettingEditResponseMigrations)(nil)).Elem(), "")
}

// A single set of migrations to apply.
type WorkerScriptSettingEditResponseMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses []WorkerScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkerScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               workerScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON               `json:"-"`
}

// workerScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingEditResponseMigrationsWorkersSingleStepMigrations]
type workerScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkerScriptSettingEditResponseMigrationsWorkersSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r WorkerScriptSettingEditResponseMigrationsWorkersSingleStepMigrations) implementsWorkerScriptSettingEditResponseMigrations() {
}

type WorkerScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From string                                                                               `json:"from"`
	To   string                                                                               `json:"to"`
	JSON workerScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// workerScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass]
type workerScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       string                                                                                   `json:"from"`
	FromScript string                                                                                   `json:"from_script"`
	To         string                                                                                   `json:"to"`
	JSON       workerScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// workerScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass]
type workerScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptSettingEditResponseMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []WorkerScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStep `json:"steps"`
	JSON  workerScriptSettingEditResponseMigrationsWorkersSteppedMigrationsJSON   `json:"-"`
}

// workerScriptSettingEditResponseMigrationsWorkersSteppedMigrationsJSON contains
// the JSON metadata for the struct
// [WorkerScriptSettingEditResponseMigrationsWorkersSteppedMigrations]
type workerScriptSettingEditResponseMigrationsWorkersSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingEditResponseMigrationsWorkersSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingEditResponseMigrationsWorkersSteppedMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r WorkerScriptSettingEditResponseMigrationsWorkersSteppedMigrations) implementsWorkerScriptSettingEditResponseMigrations() {
}

type WorkerScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []WorkerScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkerScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               workerScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON                `json:"-"`
}

// workerScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStep]
type workerScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkerScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From string                                                                                 `json:"from"`
	To   string                                                                                 `json:"to"`
	JSON workerScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// workerScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass]
type workerScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                                     `json:"from"`
	FromScript string                                                                                     `json:"from_script"`
	To         string                                                                                     `json:"to"`
	JSON       workerScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// workerScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass]
type workerScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptSettingEditResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode WorkerScriptSettingEditResponsePlacementMode `json:"mode"`
	JSON workerScriptSettingEditResponsePlacementJSON `json:"-"`
}

// workerScriptSettingEditResponsePlacementJSON contains the JSON metadata for the
// struct [WorkerScriptSettingEditResponsePlacement]
type workerScriptSettingEditResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingEditResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingEditResponsePlacementJSON) RawJSON() string {
	return r.raw
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type WorkerScriptSettingEditResponsePlacementMode string

const (
	WorkerScriptSettingEditResponsePlacementModeSmart WorkerScriptSettingEditResponsePlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type WorkerScriptSettingEditResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                          `json:"namespace"`
	JSON      workerScriptSettingEditResponseTailConsumerJSON `json:"-"`
}

// workerScriptSettingEditResponseTailConsumerJSON contains the JSON metadata for
// the struct [WorkerScriptSettingEditResponseTailConsumer]
type workerScriptSettingEditResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingEditResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingEditResponseTailConsumerJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptSettingGetResponse struct {
	// List of bindings attached to this Worker
	Bindings []WorkerScriptSettingGetResponseBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations WorkerScriptSettingGetResponseMigrations `json:"migrations"`
	Placement  WorkerScriptSettingGetResponsePlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []WorkerScriptSettingGetResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                             `json:"usage_model"`
	JSON       workerScriptSettingGetResponseJSON `json:"-"`
}

// workerScriptSettingGetResponseJSON contains the JSON metadata for the struct
// [WorkerScriptSettingGetResponse]
type workerScriptSettingGetResponseJSON struct {
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

func (r *WorkerScriptSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingGetResponseJSON) RawJSON() string {
	return r.raw
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [WorkerScriptSettingGetResponseBindingsWorkersKVNamespaceBinding],
// [WorkerScriptSettingGetResponseBindingsWorkersServiceBinding],
// [WorkerScriptSettingGetResponseBindingsWorkersDoBinding],
// [WorkerScriptSettingGetResponseBindingsWorkersR2Binding],
// [WorkerScriptSettingGetResponseBindingsWorkersQueueBinding],
// [WorkerScriptSettingGetResponseBindingsWorkersD1Binding],
// [WorkerScriptSettingGetResponseBindingsWorkersDispatchNamespaceBinding] or
// [WorkerScriptSettingGetResponseBindingsWorkersMTLSCertBinding].
type WorkerScriptSettingGetResponseBinding interface {
	implementsWorkerScriptSettingGetResponseBinding()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkerScriptSettingGetResponseBinding)(nil)).Elem(), "")
}

type WorkerScriptSettingGetResponseBindingsWorkersKVNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingGetResponseBindingsWorkersKVNamespaceBindingType `json:"type,required"`
	JSON workerScriptSettingGetResponseBindingsWorkersKVNamespaceBindingJSON `json:"-"`
}

// workerScriptSettingGetResponseBindingsWorkersKVNamespaceBindingJSON contains the
// JSON metadata for the struct
// [WorkerScriptSettingGetResponseBindingsWorkersKVNamespaceBinding]
type workerScriptSettingGetResponseBindingsWorkersKVNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseBindingsWorkersKVNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingGetResponseBindingsWorkersKVNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkerScriptSettingGetResponseBindingsWorkersKVNamespaceBinding) implementsWorkerScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingGetResponseBindingsWorkersKVNamespaceBindingType string

const (
	WorkerScriptSettingGetResponseBindingsWorkersKVNamespaceBindingTypeKVNamespace WorkerScriptSettingGetResponseBindingsWorkersKVNamespaceBindingType = "kv_namespace"
)

type WorkerScriptSettingGetResponseBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingGetResponseBindingsWorkersServiceBindingType `json:"type,required"`
	JSON workerScriptSettingGetResponseBindingsWorkersServiceBindingJSON `json:"-"`
}

// workerScriptSettingGetResponseBindingsWorkersServiceBindingJSON contains the
// JSON metadata for the struct
// [WorkerScriptSettingGetResponseBindingsWorkersServiceBinding]
type workerScriptSettingGetResponseBindingsWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseBindingsWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingGetResponseBindingsWorkersServiceBindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkerScriptSettingGetResponseBindingsWorkersServiceBinding) implementsWorkerScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingGetResponseBindingsWorkersServiceBindingType string

const (
	WorkerScriptSettingGetResponseBindingsWorkersServiceBindingTypeService WorkerScriptSettingGetResponseBindingsWorkersServiceBindingType = "service"
)

type WorkerScriptSettingGetResponseBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingGetResponseBindingsWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                     `json:"script_name"`
	JSON       workerScriptSettingGetResponseBindingsWorkersDoBindingJSON `json:"-"`
}

// workerScriptSettingGetResponseBindingsWorkersDoBindingJSON contains the JSON
// metadata for the struct [WorkerScriptSettingGetResponseBindingsWorkersDoBinding]
type workerScriptSettingGetResponseBindingsWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseBindingsWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingGetResponseBindingsWorkersDoBindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkerScriptSettingGetResponseBindingsWorkersDoBinding) implementsWorkerScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingGetResponseBindingsWorkersDoBindingType string

const (
	WorkerScriptSettingGetResponseBindingsWorkersDoBindingTypeDurableObjectNamespace WorkerScriptSettingGetResponseBindingsWorkersDoBindingType = "durable_object_namespace"
)

type WorkerScriptSettingGetResponseBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingGetResponseBindingsWorkersR2BindingType `json:"type,required"`
	JSON workerScriptSettingGetResponseBindingsWorkersR2BindingJSON `json:"-"`
}

// workerScriptSettingGetResponseBindingsWorkersR2BindingJSON contains the JSON
// metadata for the struct [WorkerScriptSettingGetResponseBindingsWorkersR2Binding]
type workerScriptSettingGetResponseBindingsWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseBindingsWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingGetResponseBindingsWorkersR2BindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkerScriptSettingGetResponseBindingsWorkersR2Binding) implementsWorkerScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingGetResponseBindingsWorkersR2BindingType string

const (
	WorkerScriptSettingGetResponseBindingsWorkersR2BindingTypeR2Bucket WorkerScriptSettingGetResponseBindingsWorkersR2BindingType = "r2_bucket"
)

type WorkerScriptSettingGetResponseBindingsWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingGetResponseBindingsWorkersQueueBindingType `json:"type,required"`
	JSON workerScriptSettingGetResponseBindingsWorkersQueueBindingJSON `json:"-"`
}

// workerScriptSettingGetResponseBindingsWorkersQueueBindingJSON contains the JSON
// metadata for the struct
// [WorkerScriptSettingGetResponseBindingsWorkersQueueBinding]
type workerScriptSettingGetResponseBindingsWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseBindingsWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingGetResponseBindingsWorkersQueueBindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkerScriptSettingGetResponseBindingsWorkersQueueBinding) implementsWorkerScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingGetResponseBindingsWorkersQueueBindingType string

const (
	WorkerScriptSettingGetResponseBindingsWorkersQueueBindingTypeQueue WorkerScriptSettingGetResponseBindingsWorkersQueueBindingType = "queue"
)

type WorkerScriptSettingGetResponseBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingGetResponseBindingsWorkersD1BindingType `json:"type,required"`
	JSON workerScriptSettingGetResponseBindingsWorkersD1BindingJSON `json:"-"`
}

// workerScriptSettingGetResponseBindingsWorkersD1BindingJSON contains the JSON
// metadata for the struct [WorkerScriptSettingGetResponseBindingsWorkersD1Binding]
type workerScriptSettingGetResponseBindingsWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseBindingsWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingGetResponseBindingsWorkersD1BindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkerScriptSettingGetResponseBindingsWorkersD1Binding) implementsWorkerScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingGetResponseBindingsWorkersD1BindingType string

const (
	WorkerScriptSettingGetResponseBindingsWorkersD1BindingTypeD1 WorkerScriptSettingGetResponseBindingsWorkersD1BindingType = "d1"
)

type WorkerScriptSettingGetResponseBindingsWorkersDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound WorkerScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     workerScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON     `json:"-"`
}

// workerScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingGetResponseBindingsWorkersDispatchNamespaceBinding]
type workerScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseBindingsWorkersDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkerScriptSettingGetResponseBindingsWorkersDispatchNamespaceBinding) implementsWorkerScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingType string

const (
	WorkerScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace WorkerScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type WorkerScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker WorkerScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   workerScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// workerScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound]
type workerScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON) RawJSON() string {
	return r.raw
}

// Outbound worker
type WorkerScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                  `json:"service"`
	JSON    workerScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// workerScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker]
type workerScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptSettingGetResponseBindingsWorkersMTLSCertBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingGetResponseBindingsWorkersMTLSCertBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                           `json:"certificate_id"`
	JSON          workerScriptSettingGetResponseBindingsWorkersMTLSCertBindingJSON `json:"-"`
}

// workerScriptSettingGetResponseBindingsWorkersMTLSCertBindingJSON contains the
// JSON metadata for the struct
// [WorkerScriptSettingGetResponseBindingsWorkersMTLSCertBinding]
type workerScriptSettingGetResponseBindingsWorkersMTLSCertBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseBindingsWorkersMTLSCertBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingGetResponseBindingsWorkersMTLSCertBindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkerScriptSettingGetResponseBindingsWorkersMTLSCertBinding) implementsWorkerScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingGetResponseBindingsWorkersMTLSCertBindingType string

const (
	WorkerScriptSettingGetResponseBindingsWorkersMTLSCertBindingTypeMTLSCertificate WorkerScriptSettingGetResponseBindingsWorkersMTLSCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [WorkerScriptSettingGetResponseMigrationsWorkersSingleStepMigrations] or
// [WorkerScriptSettingGetResponseMigrationsWorkersSteppedMigrations].
type WorkerScriptSettingGetResponseMigrations interface {
	implementsWorkerScriptSettingGetResponseMigrations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkerScriptSettingGetResponseMigrations)(nil)).Elem(), "")
}

// A single set of migrations to apply.
type WorkerScriptSettingGetResponseMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses []WorkerScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkerScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               workerScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON               `json:"-"`
}

// workerScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON contains
// the JSON metadata for the struct
// [WorkerScriptSettingGetResponseMigrationsWorkersSingleStepMigrations]
type workerScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseMigrationsWorkersSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r WorkerScriptSettingGetResponseMigrationsWorkersSingleStepMigrations) implementsWorkerScriptSettingGetResponseMigrations() {
}

type WorkerScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From string                                                                              `json:"from"`
	To   string                                                                              `json:"to"`
	JSON workerScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// workerScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass]
type workerScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       string                                                                                  `json:"from"`
	FromScript string                                                                                  `json:"from_script"`
	To         string                                                                                  `json:"to"`
	JSON       workerScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// workerScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass]
type workerScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptSettingGetResponseMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []WorkerScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStep `json:"steps"`
	JSON  workerScriptSettingGetResponseMigrationsWorkersSteppedMigrationsJSON   `json:"-"`
}

// workerScriptSettingGetResponseMigrationsWorkersSteppedMigrationsJSON contains
// the JSON metadata for the struct
// [WorkerScriptSettingGetResponseMigrationsWorkersSteppedMigrations]
type workerScriptSettingGetResponseMigrationsWorkersSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseMigrationsWorkersSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingGetResponseMigrationsWorkersSteppedMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r WorkerScriptSettingGetResponseMigrationsWorkersSteppedMigrations) implementsWorkerScriptSettingGetResponseMigrations() {
}

type WorkerScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []WorkerScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkerScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               workerScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON                `json:"-"`
}

// workerScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStep]
type workerScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From string                                                                                `json:"from"`
	To   string                                                                                `json:"to"`
	JSON workerScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// workerScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass]
type workerScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                                    `json:"from"`
	FromScript string                                                                                    `json:"from_script"`
	To         string                                                                                    `json:"to"`
	JSON       workerScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// workerScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass]
type workerScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptSettingGetResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode WorkerScriptSettingGetResponsePlacementMode `json:"mode"`
	JSON workerScriptSettingGetResponsePlacementJSON `json:"-"`
}

// workerScriptSettingGetResponsePlacementJSON contains the JSON metadata for the
// struct [WorkerScriptSettingGetResponsePlacement]
type workerScriptSettingGetResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingGetResponsePlacementJSON) RawJSON() string {
	return r.raw
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type WorkerScriptSettingGetResponsePlacementMode string

const (
	WorkerScriptSettingGetResponsePlacementModeSmart WorkerScriptSettingGetResponsePlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type WorkerScriptSettingGetResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                         `json:"namespace"`
	JSON      workerScriptSettingGetResponseTailConsumerJSON `json:"-"`
}

// workerScriptSettingGetResponseTailConsumerJSON contains the JSON metadata for
// the struct [WorkerScriptSettingGetResponseTailConsumer]
type workerScriptSettingGetResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingGetResponseTailConsumerJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptSettingEditParams struct {
	// Identifier
	AccountID param.Field[string]                                `path:"account_id,required"`
	Settings  param.Field[WorkerScriptSettingEditParamsSettings] `json:"settings"`
}

func (r WorkerScriptSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerScriptSettingEditParamsSettings struct {
	Errors   param.Field[[]WorkerScriptSettingEditParamsSettingsError]   `json:"errors,required"`
	Messages param.Field[[]WorkerScriptSettingEditParamsSettingsMessage] `json:"messages,required"`
	Result   param.Field[WorkerScriptSettingEditParamsSettingsResult]    `json:"result,required"`
	// Whether the API call was successful
	Success param.Field[WorkerScriptSettingEditParamsSettingsSuccess] `json:"success,required"`
}

func (r WorkerScriptSettingEditParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerScriptSettingEditParamsSettingsError struct {
	Code    param.Field[int64]  `json:"code,required"`
	Message param.Field[string] `json:"message,required"`
}

func (r WorkerScriptSettingEditParamsSettingsError) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerScriptSettingEditParamsSettingsMessage struct {
	Code    param.Field[int64]  `json:"code,required"`
	Message param.Field[string] `json:"message,required"`
}

func (r WorkerScriptSettingEditParamsSettingsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerScriptSettingEditParamsSettingsResult struct {
	// List of bindings attached to this Worker
	Bindings param.Field[[]WorkerScriptSettingEditParamsSettingsResultBinding] `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations param.Field[WorkerScriptSettingEditParamsSettingsResultMigrations] `json:"migrations"`
	Placement  param.Field[WorkerScriptSettingEditParamsSettingsResultPlacement]  `json:"placement"`
	// Tags to help you manage your Workers
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]WorkerScriptSettingEditParamsSettingsResultTailConsumer] `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel param.Field[string] `json:"usage_model"`
}

func (r WorkerScriptSettingEditParamsSettingsResult) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A binding to allow the Worker to communicate with resources
//
// Satisfied by
// [WorkerScriptSettingEditParamsSettingsResultBindingsWorkersKVNamespaceBinding],
// [WorkerScriptSettingEditParamsSettingsResultBindingsWorkersServiceBinding],
// [WorkerScriptSettingEditParamsSettingsResultBindingsWorkersDoBinding],
// [WorkerScriptSettingEditParamsSettingsResultBindingsWorkersR2Binding],
// [WorkerScriptSettingEditParamsSettingsResultBindingsWorkersQueueBinding],
// [WorkerScriptSettingEditParamsSettingsResultBindingsWorkersD1Binding],
// [WorkerScriptSettingEditParamsSettingsResultBindingsWorkersDispatchNamespaceBinding],
// [WorkerScriptSettingEditParamsSettingsResultBindingsWorkersMTLSCertBinding].
type WorkerScriptSettingEditParamsSettingsResultBinding interface {
	implementsWorkerScriptSettingEditParamsSettingsResultBinding()
}

type WorkerScriptSettingEditParamsSettingsResultBindingsWorkersKVNamespaceBinding struct {
	// The class of resource that the binding provides.
	Type param.Field[WorkerScriptSettingEditParamsSettingsResultBindingsWorkersKVNamespaceBindingType] `json:"type,required"`
}

func (r WorkerScriptSettingEditParamsSettingsResultBindingsWorkersKVNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerScriptSettingEditParamsSettingsResultBindingsWorkersKVNamespaceBinding) implementsWorkerScriptSettingEditParamsSettingsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingEditParamsSettingsResultBindingsWorkersKVNamespaceBindingType string

const (
	WorkerScriptSettingEditParamsSettingsResultBindingsWorkersKVNamespaceBindingTypeKVNamespace WorkerScriptSettingEditParamsSettingsResultBindingsWorkersKVNamespaceBindingType = "kv_namespace"
)

type WorkerScriptSettingEditParamsSettingsResultBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// Name of Worker to bind to
	Service param.Field[string] `json:"service,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkerScriptSettingEditParamsSettingsResultBindingsWorkersServiceBindingType] `json:"type,required"`
}

func (r WorkerScriptSettingEditParamsSettingsResultBindingsWorkersServiceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerScriptSettingEditParamsSettingsResultBindingsWorkersServiceBinding) implementsWorkerScriptSettingEditParamsSettingsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingEditParamsSettingsResultBindingsWorkersServiceBindingType string

const (
	WorkerScriptSettingEditParamsSettingsResultBindingsWorkersServiceBindingTypeService WorkerScriptSettingEditParamsSettingsResultBindingsWorkersServiceBindingType = "service"
)

type WorkerScriptSettingEditParamsSettingsResultBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName param.Field[string] `json:"class_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkerScriptSettingEditParamsSettingsResultBindingsWorkersDoBindingType] `json:"type,required"`
	// The environment of the script_name to bind to
	Environment param.Field[string] `json:"environment"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName param.Field[string] `json:"script_name"`
}

func (r WorkerScriptSettingEditParamsSettingsResultBindingsWorkersDoBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerScriptSettingEditParamsSettingsResultBindingsWorkersDoBinding) implementsWorkerScriptSettingEditParamsSettingsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingEditParamsSettingsResultBindingsWorkersDoBindingType string

const (
	WorkerScriptSettingEditParamsSettingsResultBindingsWorkersDoBindingTypeDurableObjectNamespace WorkerScriptSettingEditParamsSettingsResultBindingsWorkersDoBindingType = "durable_object_namespace"
)

type WorkerScriptSettingEditParamsSettingsResultBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName param.Field[string] `json:"bucket_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkerScriptSettingEditParamsSettingsResultBindingsWorkersR2BindingType] `json:"type,required"`
}

func (r WorkerScriptSettingEditParamsSettingsResultBindingsWorkersR2Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerScriptSettingEditParamsSettingsResultBindingsWorkersR2Binding) implementsWorkerScriptSettingEditParamsSettingsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingEditParamsSettingsResultBindingsWorkersR2BindingType string

const (
	WorkerScriptSettingEditParamsSettingsResultBindingsWorkersR2BindingTypeR2Bucket WorkerScriptSettingEditParamsSettingsResultBindingsWorkersR2BindingType = "r2_bucket"
)

type WorkerScriptSettingEditParamsSettingsResultBindingsWorkersQueueBinding struct {
	// Name of the Queue to bind to
	QueueName param.Field[string] `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkerScriptSettingEditParamsSettingsResultBindingsWorkersQueueBindingType] `json:"type,required"`
}

func (r WorkerScriptSettingEditParamsSettingsResultBindingsWorkersQueueBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerScriptSettingEditParamsSettingsResultBindingsWorkersQueueBinding) implementsWorkerScriptSettingEditParamsSettingsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingEditParamsSettingsResultBindingsWorkersQueueBindingType string

const (
	WorkerScriptSettingEditParamsSettingsResultBindingsWorkersQueueBindingTypeQueue WorkerScriptSettingEditParamsSettingsResultBindingsWorkersQueueBindingType = "queue"
)

type WorkerScriptSettingEditParamsSettingsResultBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID param.Field[string] `json:"id,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name param.Field[string] `json:"name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkerScriptSettingEditParamsSettingsResultBindingsWorkersD1BindingType] `json:"type,required"`
}

func (r WorkerScriptSettingEditParamsSettingsResultBindingsWorkersD1Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerScriptSettingEditParamsSettingsResultBindingsWorkersD1Binding) implementsWorkerScriptSettingEditParamsSettingsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingEditParamsSettingsResultBindingsWorkersD1BindingType string

const (
	WorkerScriptSettingEditParamsSettingsResultBindingsWorkersD1BindingTypeD1 WorkerScriptSettingEditParamsSettingsResultBindingsWorkersD1BindingType = "d1"
)

type WorkerScriptSettingEditParamsSettingsResultBindingsWorkersDispatchNamespaceBinding struct {
	// Namespace to bind to
	Namespace param.Field[string] `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkerScriptSettingEditParamsSettingsResultBindingsWorkersDispatchNamespaceBindingType] `json:"type,required"`
	// Outbound worker
	Outbound param.Field[WorkerScriptSettingEditParamsSettingsResultBindingsWorkersDispatchNamespaceBindingOutbound] `json:"outbound"`
}

func (r WorkerScriptSettingEditParamsSettingsResultBindingsWorkersDispatchNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerScriptSettingEditParamsSettingsResultBindingsWorkersDispatchNamespaceBinding) implementsWorkerScriptSettingEditParamsSettingsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingEditParamsSettingsResultBindingsWorkersDispatchNamespaceBindingType string

const (
	WorkerScriptSettingEditParamsSettingsResultBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace WorkerScriptSettingEditParamsSettingsResultBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type WorkerScriptSettingEditParamsSettingsResultBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params param.Field[[]string] `json:"params"`
	// Outbound worker
	Worker param.Field[WorkerScriptSettingEditParamsSettingsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker] `json:"worker"`
}

func (r WorkerScriptSettingEditParamsSettingsResultBindingsWorkersDispatchNamespaceBindingOutbound) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker
type WorkerScriptSettingEditParamsSettingsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker
	Service param.Field[string] `json:"service"`
}

func (r WorkerScriptSettingEditParamsSettingsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerScriptSettingEditParamsSettingsResultBindingsWorkersMTLSCertBinding struct {
	// The class of resource that the binding provides.
	Type param.Field[WorkerScriptSettingEditParamsSettingsResultBindingsWorkersMTLSCertBindingType] `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID param.Field[string] `json:"certificate_id"`
}

func (r WorkerScriptSettingEditParamsSettingsResultBindingsWorkersMTLSCertBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerScriptSettingEditParamsSettingsResultBindingsWorkersMTLSCertBinding) implementsWorkerScriptSettingEditParamsSettingsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingEditParamsSettingsResultBindingsWorkersMTLSCertBindingType string

const (
	WorkerScriptSettingEditParamsSettingsResultBindingsWorkersMTLSCertBindingTypeMTLSCertificate WorkerScriptSettingEditParamsSettingsResultBindingsWorkersMTLSCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by
// [WorkerScriptSettingEditParamsSettingsResultMigrationsWorkersSingleStepMigrations],
// [WorkerScriptSettingEditParamsSettingsResultMigrationsWorkersSteppedMigrations].
type WorkerScriptSettingEditParamsSettingsResultMigrations interface {
	implementsWorkerScriptSettingEditParamsSettingsResultMigrations()
}

// A single set of migrations to apply.
type WorkerScriptSettingEditParamsSettingsResultMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses param.Field[[]WorkerScriptSettingEditParamsSettingsResultMigrationsWorkersSingleStepMigrationsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]WorkerScriptSettingEditParamsSettingsResultMigrationsWorkersSingleStepMigrationsTransferredClass] `json:"transferred_classes"`
}

func (r WorkerScriptSettingEditParamsSettingsResultMigrationsWorkersSingleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerScriptSettingEditParamsSettingsResultMigrationsWorkersSingleStepMigrations) implementsWorkerScriptSettingEditParamsSettingsResultMigrations() {
}

type WorkerScriptSettingEditParamsSettingsResultMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r WorkerScriptSettingEditParamsSettingsResultMigrationsWorkersSingleStepMigrationsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerScriptSettingEditParamsSettingsResultMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r WorkerScriptSettingEditParamsSettingsResultMigrationsWorkersSingleStepMigrationsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerScriptSettingEditParamsSettingsResultMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]WorkerScriptSettingEditParamsSettingsResultMigrationsWorkersSteppedMigrationsStep] `json:"steps"`
}

func (r WorkerScriptSettingEditParamsSettingsResultMigrationsWorkersSteppedMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerScriptSettingEditParamsSettingsResultMigrationsWorkersSteppedMigrations) implementsWorkerScriptSettingEditParamsSettingsResultMigrations() {
}

type WorkerScriptSettingEditParamsSettingsResultMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]WorkerScriptSettingEditParamsSettingsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]WorkerScriptSettingEditParamsSettingsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass] `json:"transferred_classes"`
}

func (r WorkerScriptSettingEditParamsSettingsResultMigrationsWorkersSteppedMigrationsStep) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerScriptSettingEditParamsSettingsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r WorkerScriptSettingEditParamsSettingsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerScriptSettingEditParamsSettingsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r WorkerScriptSettingEditParamsSettingsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerScriptSettingEditParamsSettingsResultPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode param.Field[WorkerScriptSettingEditParamsSettingsResultPlacementMode] `json:"mode"`
}

func (r WorkerScriptSettingEditParamsSettingsResultPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type WorkerScriptSettingEditParamsSettingsResultPlacementMode string

const (
	WorkerScriptSettingEditParamsSettingsResultPlacementModeSmart WorkerScriptSettingEditParamsSettingsResultPlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type WorkerScriptSettingEditParamsSettingsResultTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service param.Field[string] `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace param.Field[string] `json:"namespace"`
}

func (r WorkerScriptSettingEditParamsSettingsResultTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the API call was successful
type WorkerScriptSettingEditParamsSettingsSuccess bool

const (
	WorkerScriptSettingEditParamsSettingsSuccessTrue WorkerScriptSettingEditParamsSettingsSuccess = true
)

type WorkerScriptSettingEditResponseEnvelope struct {
	Errors   []WorkerScriptSettingEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerScriptSettingEditResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerScriptSettingEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerScriptSettingEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerScriptSettingEditResponseEnvelopeJSON    `json:"-"`
}

// workerScriptSettingEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [WorkerScriptSettingEditResponseEnvelope]
type workerScriptSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptSettingEditResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    workerScriptSettingEditResponseEnvelopeErrorsJSON `json:"-"`
}

// workerScriptSettingEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [WorkerScriptSettingEditResponseEnvelopeErrors]
type workerScriptSettingEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptSettingEditResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    workerScriptSettingEditResponseEnvelopeMessagesJSON `json:"-"`
}

// workerScriptSettingEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [WorkerScriptSettingEditResponseEnvelopeMessages]
type workerScriptSettingEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WorkerScriptSettingEditResponseEnvelopeSuccess bool

const (
	WorkerScriptSettingEditResponseEnvelopeSuccessTrue WorkerScriptSettingEditResponseEnvelopeSuccess = true
)

type WorkerScriptSettingGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type WorkerScriptSettingGetResponseEnvelope struct {
	Errors   []WorkerScriptSettingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerScriptSettingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerScriptSettingGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerScriptSettingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerScriptSettingGetResponseEnvelopeJSON    `json:"-"`
}

// workerScriptSettingGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [WorkerScriptSettingGetResponseEnvelope]
type workerScriptSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptSettingGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    workerScriptSettingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// workerScriptSettingGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [WorkerScriptSettingGetResponseEnvelopeErrors]
type workerScriptSettingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WorkerScriptSettingGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    workerScriptSettingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// workerScriptSettingGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [WorkerScriptSettingGetResponseEnvelopeMessages]
type workerScriptSettingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerScriptSettingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WorkerScriptSettingGetResponseEnvelopeSuccess bool

const (
	WorkerScriptSettingGetResponseEnvelopeSuccessTrue WorkerScriptSettingGetResponseEnvelopeSuccess = true
)
