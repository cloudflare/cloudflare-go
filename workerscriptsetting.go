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

// Get script metadata and config, such as bindings or usage model
func (r *WorkerScriptSettingService) Get(ctx context.Context, accountID string, scriptName string, opts ...option.RequestOption) (res *WorkerScriptSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerScriptSettingGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/settings", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Patch script metadata or config, such as bindings or usage model
func (r *WorkerScriptSettingService) Update(ctx context.Context, accountID string, scriptName string, body WorkerScriptSettingUpdateParams, opts ...option.RequestOption) (res *WorkerScriptSettingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerScriptSettingUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/settings", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
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

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [WorkerScriptSettingGetResponseBindingsWorkersKvNamespaceBinding],
// [WorkerScriptSettingGetResponseBindingsWorkersServiceBinding],
// [WorkerScriptSettingGetResponseBindingsWorkersDoBinding],
// [WorkerScriptSettingGetResponseBindingsWorkersR2Binding],
// [WorkerScriptSettingGetResponseBindingsWorkersQueueBinding],
// [WorkerScriptSettingGetResponseBindingsWorkersD1Binding],
// [WorkerScriptSettingGetResponseBindingsWorkersDispatchNamespaceBinding] or
// [WorkerScriptSettingGetResponseBindingsWorkersMtlsCertBinding].
type WorkerScriptSettingGetResponseBinding interface {
	implementsWorkerScriptSettingGetResponseBinding()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkerScriptSettingGetResponseBinding)(nil)).Elem(), "")
}

type WorkerScriptSettingGetResponseBindingsWorkersKvNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingGetResponseBindingsWorkersKvNamespaceBindingType `json:"type,required"`
	JSON workerScriptSettingGetResponseBindingsWorkersKvNamespaceBindingJSON `json:"-"`
}

// workerScriptSettingGetResponseBindingsWorkersKvNamespaceBindingJSON contains the
// JSON metadata for the struct
// [WorkerScriptSettingGetResponseBindingsWorkersKvNamespaceBinding]
type workerScriptSettingGetResponseBindingsWorkersKvNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseBindingsWorkersKvNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingGetResponseBindingsWorkersKvNamespaceBinding) implementsWorkerScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingGetResponseBindingsWorkersKvNamespaceBindingType string

const (
	WorkerScriptSettingGetResponseBindingsWorkersKvNamespaceBindingTypeKvNamespace WorkerScriptSettingGetResponseBindingsWorkersKvNamespaceBindingType = "kv_namespace"
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

type WorkerScriptSettingGetResponseBindingsWorkersMtlsCertBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingGetResponseBindingsWorkersMtlsCertBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                           `json:"certificate_id"`
	JSON          workerScriptSettingGetResponseBindingsWorkersMtlsCertBindingJSON `json:"-"`
}

// workerScriptSettingGetResponseBindingsWorkersMtlsCertBindingJSON contains the
// JSON metadata for the struct
// [WorkerScriptSettingGetResponseBindingsWorkersMtlsCertBinding]
type workerScriptSettingGetResponseBindingsWorkersMtlsCertBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseBindingsWorkersMtlsCertBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingGetResponseBindingsWorkersMtlsCertBinding) implementsWorkerScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingGetResponseBindingsWorkersMtlsCertBindingType string

const (
	WorkerScriptSettingGetResponseBindingsWorkersMtlsCertBindingTypeMtlsCertificate WorkerScriptSettingGetResponseBindingsWorkersMtlsCertBindingType = "mtls_certificate"
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

type WorkerScriptSettingUpdateResponse struct {
	// List of bindings attached to this Worker
	Bindings []WorkerScriptSettingUpdateResponseBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations WorkerScriptSettingUpdateResponseMigrations `json:"migrations"`
	Placement  WorkerScriptSettingUpdateResponsePlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []WorkerScriptSettingUpdateResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                `json:"usage_model"`
	JSON       workerScriptSettingUpdateResponseJSON `json:"-"`
}

// workerScriptSettingUpdateResponseJSON contains the JSON metadata for the struct
// [WorkerScriptSettingUpdateResponse]
type workerScriptSettingUpdateResponseJSON struct {
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

func (r *WorkerScriptSettingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [WorkerScriptSettingUpdateResponseBindingsWorkersKvNamespaceBinding],
// [WorkerScriptSettingUpdateResponseBindingsWorkersServiceBinding],
// [WorkerScriptSettingUpdateResponseBindingsWorkersDoBinding],
// [WorkerScriptSettingUpdateResponseBindingsWorkersR2Binding],
// [WorkerScriptSettingUpdateResponseBindingsWorkersQueueBinding],
// [WorkerScriptSettingUpdateResponseBindingsWorkersD1Binding],
// [WorkerScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBinding] or
// [WorkerScriptSettingUpdateResponseBindingsWorkersMtlsCertBinding].
type WorkerScriptSettingUpdateResponseBinding interface {
	implementsWorkerScriptSettingUpdateResponseBinding()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkerScriptSettingUpdateResponseBinding)(nil)).Elem(), "")
}

type WorkerScriptSettingUpdateResponseBindingsWorkersKvNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingUpdateResponseBindingsWorkersKvNamespaceBindingType `json:"type,required"`
	JSON workerScriptSettingUpdateResponseBindingsWorkersKvNamespaceBindingJSON `json:"-"`
}

// workerScriptSettingUpdateResponseBindingsWorkersKvNamespaceBindingJSON contains
// the JSON metadata for the struct
// [WorkerScriptSettingUpdateResponseBindingsWorkersKvNamespaceBinding]
type workerScriptSettingUpdateResponseBindingsWorkersKvNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseBindingsWorkersKvNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingUpdateResponseBindingsWorkersKvNamespaceBinding) implementsWorkerScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingUpdateResponseBindingsWorkersKvNamespaceBindingType string

const (
	WorkerScriptSettingUpdateResponseBindingsWorkersKvNamespaceBindingTypeKvNamespace WorkerScriptSettingUpdateResponseBindingsWorkersKvNamespaceBindingType = "kv_namespace"
)

type WorkerScriptSettingUpdateResponseBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingUpdateResponseBindingsWorkersServiceBindingType `json:"type,required"`
	JSON workerScriptSettingUpdateResponseBindingsWorkersServiceBindingJSON `json:"-"`
}

// workerScriptSettingUpdateResponseBindingsWorkersServiceBindingJSON contains the
// JSON metadata for the struct
// [WorkerScriptSettingUpdateResponseBindingsWorkersServiceBinding]
type workerScriptSettingUpdateResponseBindingsWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseBindingsWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingUpdateResponseBindingsWorkersServiceBinding) implementsWorkerScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingUpdateResponseBindingsWorkersServiceBindingType string

const (
	WorkerScriptSettingUpdateResponseBindingsWorkersServiceBindingTypeService WorkerScriptSettingUpdateResponseBindingsWorkersServiceBindingType = "service"
)

type WorkerScriptSettingUpdateResponseBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingUpdateResponseBindingsWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                        `json:"script_name"`
	JSON       workerScriptSettingUpdateResponseBindingsWorkersDoBindingJSON `json:"-"`
}

// workerScriptSettingUpdateResponseBindingsWorkersDoBindingJSON contains the JSON
// metadata for the struct
// [WorkerScriptSettingUpdateResponseBindingsWorkersDoBinding]
type workerScriptSettingUpdateResponseBindingsWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseBindingsWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingUpdateResponseBindingsWorkersDoBinding) implementsWorkerScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingUpdateResponseBindingsWorkersDoBindingType string

const (
	WorkerScriptSettingUpdateResponseBindingsWorkersDoBindingTypeDurableObjectNamespace WorkerScriptSettingUpdateResponseBindingsWorkersDoBindingType = "durable_object_namespace"
)

type WorkerScriptSettingUpdateResponseBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingUpdateResponseBindingsWorkersR2BindingType `json:"type,required"`
	JSON workerScriptSettingUpdateResponseBindingsWorkersR2BindingJSON `json:"-"`
}

// workerScriptSettingUpdateResponseBindingsWorkersR2BindingJSON contains the JSON
// metadata for the struct
// [WorkerScriptSettingUpdateResponseBindingsWorkersR2Binding]
type workerScriptSettingUpdateResponseBindingsWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseBindingsWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingUpdateResponseBindingsWorkersR2Binding) implementsWorkerScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingUpdateResponseBindingsWorkersR2BindingType string

const (
	WorkerScriptSettingUpdateResponseBindingsWorkersR2BindingTypeR2Bucket WorkerScriptSettingUpdateResponseBindingsWorkersR2BindingType = "r2_bucket"
)

type WorkerScriptSettingUpdateResponseBindingsWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingUpdateResponseBindingsWorkersQueueBindingType `json:"type,required"`
	JSON workerScriptSettingUpdateResponseBindingsWorkersQueueBindingJSON `json:"-"`
}

// workerScriptSettingUpdateResponseBindingsWorkersQueueBindingJSON contains the
// JSON metadata for the struct
// [WorkerScriptSettingUpdateResponseBindingsWorkersQueueBinding]
type workerScriptSettingUpdateResponseBindingsWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseBindingsWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingUpdateResponseBindingsWorkersQueueBinding) implementsWorkerScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingUpdateResponseBindingsWorkersQueueBindingType string

const (
	WorkerScriptSettingUpdateResponseBindingsWorkersQueueBindingTypeQueue WorkerScriptSettingUpdateResponseBindingsWorkersQueueBindingType = "queue"
)

type WorkerScriptSettingUpdateResponseBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingUpdateResponseBindingsWorkersD1BindingType `json:"type,required"`
	JSON workerScriptSettingUpdateResponseBindingsWorkersD1BindingJSON `json:"-"`
}

// workerScriptSettingUpdateResponseBindingsWorkersD1BindingJSON contains the JSON
// metadata for the struct
// [WorkerScriptSettingUpdateResponseBindingsWorkersD1Binding]
type workerScriptSettingUpdateResponseBindingsWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseBindingsWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingUpdateResponseBindingsWorkersD1Binding) implementsWorkerScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingUpdateResponseBindingsWorkersD1BindingType string

const (
	WorkerScriptSettingUpdateResponseBindingsWorkersD1BindingTypeD1 WorkerScriptSettingUpdateResponseBindingsWorkersD1BindingType = "d1"
)

type WorkerScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound WorkerScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     workerScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingJSON     `json:"-"`
}

// workerScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBinding]
type workerScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBinding) implementsWorkerScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingType string

const (
	WorkerScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace WorkerScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type WorkerScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker WorkerScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   workerScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// workerScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutbound]
type workerScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Outbound worker
type WorkerScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                     `json:"service"`
	JSON    workerScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// workerScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker]
type workerScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptSettingUpdateResponseBindingsWorkersMtlsCertBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingUpdateResponseBindingsWorkersMtlsCertBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                              `json:"certificate_id"`
	JSON          workerScriptSettingUpdateResponseBindingsWorkersMtlsCertBindingJSON `json:"-"`
}

// workerScriptSettingUpdateResponseBindingsWorkersMtlsCertBindingJSON contains the
// JSON metadata for the struct
// [WorkerScriptSettingUpdateResponseBindingsWorkersMtlsCertBinding]
type workerScriptSettingUpdateResponseBindingsWorkersMtlsCertBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseBindingsWorkersMtlsCertBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingUpdateResponseBindingsWorkersMtlsCertBinding) implementsWorkerScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingUpdateResponseBindingsWorkersMtlsCertBindingType string

const (
	WorkerScriptSettingUpdateResponseBindingsWorkersMtlsCertBindingTypeMtlsCertificate WorkerScriptSettingUpdateResponseBindingsWorkersMtlsCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [WorkerScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrations] or
// [WorkerScriptSettingUpdateResponseMigrationsWorkersSteppedMigrations].
type WorkerScriptSettingUpdateResponseMigrations interface {
	implementsWorkerScriptSettingUpdateResponseMigrations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkerScriptSettingUpdateResponseMigrations)(nil)).Elem(), "")
}

// A single set of migrations to apply.
type WorkerScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses []WorkerScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkerScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               workerScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsJSON               `json:"-"`
}

// workerScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrations]
type workerScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrations) implementsWorkerScriptSettingUpdateResponseMigrations() {
}

type WorkerScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From string                                                                                 `json:"from"`
	To   string                                                                                 `json:"to"`
	JSON workerScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// workerScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsRenamedClass]
type workerScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       string                                                                                     `json:"from"`
	FromScript string                                                                                     `json:"from_script"`
	To         string                                                                                     `json:"to"`
	JSON       workerScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// workerScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsTransferredClass]
type workerScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptSettingUpdateResponseMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []WorkerScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStep `json:"steps"`
	JSON  workerScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsJSON   `json:"-"`
}

// workerScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsJSON contains
// the JSON metadata for the struct
// [WorkerScriptSettingUpdateResponseMigrationsWorkersSteppedMigrations]
type workerScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseMigrationsWorkersSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingUpdateResponseMigrationsWorkersSteppedMigrations) implementsWorkerScriptSettingUpdateResponseMigrations() {
}

type WorkerScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []WorkerScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkerScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               workerScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepJSON                `json:"-"`
}

// workerScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStep]
type workerScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From string                                                                                   `json:"from"`
	To   string                                                                                   `json:"to"`
	JSON workerScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// workerScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass]
type workerScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                                       `json:"from"`
	FromScript string                                                                                       `json:"from_script"`
	To         string                                                                                       `json:"to"`
	JSON       workerScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// workerScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass]
type workerScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptSettingUpdateResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode WorkerScriptSettingUpdateResponsePlacementMode `json:"mode"`
	JSON workerScriptSettingUpdateResponsePlacementJSON `json:"-"`
}

// workerScriptSettingUpdateResponsePlacementJSON contains the JSON metadata for
// the struct [WorkerScriptSettingUpdateResponsePlacement]
type workerScriptSettingUpdateResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type WorkerScriptSettingUpdateResponsePlacementMode string

const (
	WorkerScriptSettingUpdateResponsePlacementModeSmart WorkerScriptSettingUpdateResponsePlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type WorkerScriptSettingUpdateResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                            `json:"namespace"`
	JSON      workerScriptSettingUpdateResponseTailConsumerJSON `json:"-"`
}

// workerScriptSettingUpdateResponseTailConsumerJSON contains the JSON metadata for
// the struct [WorkerScriptSettingUpdateResponseTailConsumer]
type workerScriptSettingUpdateResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptSettingGetResponseEnvelope struct {
	Errors   []WorkerScriptSettingGetResponseEnvelopeErrors   `json:"errors"`
	Messages []WorkerScriptSettingGetResponseEnvelopeMessages `json:"messages"`
	Result   WorkerScriptSettingGetResponse                   `json:"result"`
	// Whether the API call was successful
	Success WorkerScriptSettingGetResponseEnvelopeSuccess `json:"success"`
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

// Whether the API call was successful
type WorkerScriptSettingGetResponseEnvelopeSuccess bool

const (
	WorkerScriptSettingGetResponseEnvelopeSuccessTrue WorkerScriptSettingGetResponseEnvelopeSuccess = true
)

type WorkerScriptSettingUpdateParams struct {
	Settings param.Field[WorkerScriptSettingUpdateParamsSettings] `json:"settings"`
}

func (r WorkerScriptSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerScriptSettingUpdateParamsSettings struct {
	Errors   param.Field[[]WorkerScriptSettingUpdateParamsSettingsError]   `json:"errors"`
	Messages param.Field[[]WorkerScriptSettingUpdateParamsSettingsMessage] `json:"messages"`
	Result   param.Field[WorkerScriptSettingUpdateParamsSettingsResult]    `json:"result"`
	// Whether the API call was successful
	Success param.Field[WorkerScriptSettingUpdateParamsSettingsSuccess] `json:"success"`
}

func (r WorkerScriptSettingUpdateParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerScriptSettingUpdateParamsSettingsError struct {
	Code    param.Field[int64]  `json:"code,required"`
	Message param.Field[string] `json:"message,required"`
}

func (r WorkerScriptSettingUpdateParamsSettingsError) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerScriptSettingUpdateParamsSettingsMessage struct {
	Code    param.Field[int64]  `json:"code,required"`
	Message param.Field[string] `json:"message,required"`
}

func (r WorkerScriptSettingUpdateParamsSettingsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerScriptSettingUpdateParamsSettingsResult struct {
	// List of bindings attached to this Worker
	Bindings param.Field[[]WorkerScriptSettingUpdateParamsSettingsResultBinding] `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations param.Field[WorkerScriptSettingUpdateParamsSettingsResultMigrations] `json:"migrations"`
	Placement  param.Field[WorkerScriptSettingUpdateParamsSettingsResultPlacement]  `json:"placement"`
	// Tags to help you manage your Workers
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]WorkerScriptSettingUpdateParamsSettingsResultTailConsumer] `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel param.Field[string] `json:"usage_model"`
}

func (r WorkerScriptSettingUpdateParamsSettingsResult) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A binding to allow the Worker to communicate with resources
//
// Satisfied by
// [WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersKvNamespaceBinding],
// [WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersServiceBinding],
// [WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersDoBinding],
// [WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersR2Binding],
// [WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersQueueBinding],
// [WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersD1Binding],
// [WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersDispatchNamespaceBinding],
// [WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersMtlsCertBinding].
type WorkerScriptSettingUpdateParamsSettingsResultBinding interface {
	implementsWorkerScriptSettingUpdateParamsSettingsResultBinding()
}

type WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersKvNamespaceBinding struct {
	// The class of resource that the binding provides.
	Type param.Field[WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersKvNamespaceBindingType] `json:"type,required"`
}

func (r WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersKvNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersKvNamespaceBinding) implementsWorkerScriptSettingUpdateParamsSettingsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersKvNamespaceBindingType string

const (
	WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersKvNamespaceBindingTypeKvNamespace WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersKvNamespaceBindingType = "kv_namespace"
)

type WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// Name of Worker to bind to
	Service param.Field[string] `json:"service,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersServiceBindingType] `json:"type,required"`
}

func (r WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersServiceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersServiceBinding) implementsWorkerScriptSettingUpdateParamsSettingsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersServiceBindingType string

const (
	WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersServiceBindingTypeService WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersServiceBindingType = "service"
)

type WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName param.Field[string] `json:"class_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersDoBindingType] `json:"type,required"`
	// The environment of the script_name to bind to
	Environment param.Field[string] `json:"environment"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName param.Field[string] `json:"script_name"`
}

func (r WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersDoBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersDoBinding) implementsWorkerScriptSettingUpdateParamsSettingsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersDoBindingType string

const (
	WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersDoBindingTypeDurableObjectNamespace WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersDoBindingType = "durable_object_namespace"
)

type WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName param.Field[string] `json:"bucket_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersR2BindingType] `json:"type,required"`
}

func (r WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersR2Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersR2Binding) implementsWorkerScriptSettingUpdateParamsSettingsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersR2BindingType string

const (
	WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersR2BindingTypeR2Bucket WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersR2BindingType = "r2_bucket"
)

type WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersQueueBinding struct {
	// Name of the Queue to bind to
	QueueName param.Field[string] `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersQueueBindingType] `json:"type,required"`
}

func (r WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersQueueBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersQueueBinding) implementsWorkerScriptSettingUpdateParamsSettingsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersQueueBindingType string

const (
	WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersQueueBindingTypeQueue WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersQueueBindingType = "queue"
)

type WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID param.Field[string] `json:"id,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name param.Field[string] `json:"name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersD1BindingType] `json:"type,required"`
}

func (r WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersD1Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersD1Binding) implementsWorkerScriptSettingUpdateParamsSettingsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersD1BindingType string

const (
	WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersD1BindingTypeD1 WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersD1BindingType = "d1"
)

type WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersDispatchNamespaceBinding struct {
	// Namespace to bind to
	Namespace param.Field[string] `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersDispatchNamespaceBindingType] `json:"type,required"`
	// Outbound worker
	Outbound param.Field[WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersDispatchNamespaceBindingOutbound] `json:"outbound"`
}

func (r WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersDispatchNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersDispatchNamespaceBinding) implementsWorkerScriptSettingUpdateParamsSettingsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersDispatchNamespaceBindingType string

const (
	WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params param.Field[[]string] `json:"params"`
	// Outbound worker
	Worker param.Field[WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker] `json:"worker"`
}

func (r WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersDispatchNamespaceBindingOutbound) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker
type WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker
	Service param.Field[string] `json:"service"`
}

func (r WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersMtlsCertBinding struct {
	// The class of resource that the binding provides.
	Type param.Field[WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersMtlsCertBindingType] `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID param.Field[string] `json:"certificate_id"`
}

func (r WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersMtlsCertBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersMtlsCertBinding) implementsWorkerScriptSettingUpdateParamsSettingsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersMtlsCertBindingType string

const (
	WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersMtlsCertBindingTypeMtlsCertificate WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersMtlsCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by
// [WorkerScriptSettingUpdateParamsSettingsResultMigrationsWorkersSingleStepMigrations],
// [WorkerScriptSettingUpdateParamsSettingsResultMigrationsWorkersSteppedMigrations].
type WorkerScriptSettingUpdateParamsSettingsResultMigrations interface {
	implementsWorkerScriptSettingUpdateParamsSettingsResultMigrations()
}

// A single set of migrations to apply.
type WorkerScriptSettingUpdateParamsSettingsResultMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses param.Field[[]WorkerScriptSettingUpdateParamsSettingsResultMigrationsWorkersSingleStepMigrationsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]WorkerScriptSettingUpdateParamsSettingsResultMigrationsWorkersSingleStepMigrationsTransferredClass] `json:"transferred_classes"`
}

func (r WorkerScriptSettingUpdateParamsSettingsResultMigrationsWorkersSingleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerScriptSettingUpdateParamsSettingsResultMigrationsWorkersSingleStepMigrations) implementsWorkerScriptSettingUpdateParamsSettingsResultMigrations() {
}

type WorkerScriptSettingUpdateParamsSettingsResultMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r WorkerScriptSettingUpdateParamsSettingsResultMigrationsWorkersSingleStepMigrationsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerScriptSettingUpdateParamsSettingsResultMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r WorkerScriptSettingUpdateParamsSettingsResultMigrationsWorkersSingleStepMigrationsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerScriptSettingUpdateParamsSettingsResultMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]WorkerScriptSettingUpdateParamsSettingsResultMigrationsWorkersSteppedMigrationsStep] `json:"steps"`
}

func (r WorkerScriptSettingUpdateParamsSettingsResultMigrationsWorkersSteppedMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerScriptSettingUpdateParamsSettingsResultMigrationsWorkersSteppedMigrations) implementsWorkerScriptSettingUpdateParamsSettingsResultMigrations() {
}

type WorkerScriptSettingUpdateParamsSettingsResultMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]WorkerScriptSettingUpdateParamsSettingsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]WorkerScriptSettingUpdateParamsSettingsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass] `json:"transferred_classes"`
}

func (r WorkerScriptSettingUpdateParamsSettingsResultMigrationsWorkersSteppedMigrationsStep) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerScriptSettingUpdateParamsSettingsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r WorkerScriptSettingUpdateParamsSettingsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerScriptSettingUpdateParamsSettingsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r WorkerScriptSettingUpdateParamsSettingsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerScriptSettingUpdateParamsSettingsResultPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode param.Field[WorkerScriptSettingUpdateParamsSettingsResultPlacementMode] `json:"mode"`
}

func (r WorkerScriptSettingUpdateParamsSettingsResultPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type WorkerScriptSettingUpdateParamsSettingsResultPlacementMode string

const (
	WorkerScriptSettingUpdateParamsSettingsResultPlacementModeSmart WorkerScriptSettingUpdateParamsSettingsResultPlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type WorkerScriptSettingUpdateParamsSettingsResultTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service param.Field[string] `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace param.Field[string] `json:"namespace"`
}

func (r WorkerScriptSettingUpdateParamsSettingsResultTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the API call was successful
type WorkerScriptSettingUpdateParamsSettingsSuccess bool

const (
	WorkerScriptSettingUpdateParamsSettingsSuccessTrue WorkerScriptSettingUpdateParamsSettingsSuccess = true
)

type WorkerScriptSettingUpdateResponseEnvelope struct {
	Errors   []WorkerScriptSettingUpdateResponseEnvelopeErrors   `json:"errors"`
	Messages []WorkerScriptSettingUpdateResponseEnvelopeMessages `json:"messages"`
	Result   WorkerScriptSettingUpdateResponse                   `json:"result"`
	// Whether the API call was successful
	Success WorkerScriptSettingUpdateResponseEnvelopeSuccess `json:"success"`
	JSON    workerScriptSettingUpdateResponseEnvelopeJSON    `json:"-"`
}

// workerScriptSettingUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [WorkerScriptSettingUpdateResponseEnvelope]
type workerScriptSettingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptSettingUpdateResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    workerScriptSettingUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// workerScriptSettingUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [WorkerScriptSettingUpdateResponseEnvelopeErrors]
type workerScriptSettingUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptSettingUpdateResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    workerScriptSettingUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// workerScriptSettingUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [WorkerScriptSettingUpdateResponseEnvelopeMessages]
type workerScriptSettingUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerScriptSettingUpdateResponseEnvelopeSuccess bool

const (
	WorkerScriptSettingUpdateResponseEnvelopeSuccessTrue WorkerScriptSettingUpdateResponseEnvelopeSuccess = true
)
