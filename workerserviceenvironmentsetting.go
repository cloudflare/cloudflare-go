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

// WorkerServiceEnvironmentSettingService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewWorkerServiceEnvironmentSettingService] method instead.
type WorkerServiceEnvironmentSettingService struct {
	Options []option.RequestOption
}

// NewWorkerServiceEnvironmentSettingService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWorkerServiceEnvironmentSettingService(opts ...option.RequestOption) (r *WorkerServiceEnvironmentSettingService) {
	r = &WorkerServiceEnvironmentSettingService{}
	r.Options = opts
	return
}

// Get script settings from a worker with an environment
func (r *WorkerServiceEnvironmentSettingService) Get(ctx context.Context, accountID string, serviceName string, environmentName string, opts ...option.RequestOption) (res *WorkerServiceEnvironmentSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerServiceEnvironmentSettingGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/services/%s/environments/%s/settings", accountID, serviceName, environmentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Patch script metadata, such as bindings
func (r *WorkerServiceEnvironmentSettingService) Modify(ctx context.Context, accountID string, serviceName string, environmentName string, body WorkerServiceEnvironmentSettingModifyParams, opts ...option.RequestOption) (res *WorkerServiceEnvironmentSettingModifyResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerServiceEnvironmentSettingModifyResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/services/%s/environments/%s/settings", accountID, serviceName, environmentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkerServiceEnvironmentSettingGetResponse struct {
	// List of bindings attached to this Worker
	Bindings []WorkerServiceEnvironmentSettingGetResponseBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations WorkerServiceEnvironmentSettingGetResponseMigrations `json:"migrations"`
	Placement  WorkerServiceEnvironmentSettingGetResponsePlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []WorkerServiceEnvironmentSettingGetResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                         `json:"usage_model"`
	JSON       workerServiceEnvironmentSettingGetResponseJSON `json:"-"`
}

// workerServiceEnvironmentSettingGetResponseJSON contains the JSON metadata for
// the struct [WorkerServiceEnvironmentSettingGetResponse]
type workerServiceEnvironmentSettingGetResponseJSON struct {
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

func (r *WorkerServiceEnvironmentSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [WorkerServiceEnvironmentSettingGetResponseBindingsWorkersKvNamespaceBinding],
// [WorkerServiceEnvironmentSettingGetResponseBindingsWorkersServiceBinding],
// [WorkerServiceEnvironmentSettingGetResponseBindingsWorkersDoBinding],
// [WorkerServiceEnvironmentSettingGetResponseBindingsWorkersR2Binding],
// [WorkerServiceEnvironmentSettingGetResponseBindingsWorkersQueueBinding],
// [WorkerServiceEnvironmentSettingGetResponseBindingsWorkersD1Binding],
// [WorkerServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBinding]
// or [WorkerServiceEnvironmentSettingGetResponseBindingsWorkersMtlsCertBinding].
type WorkerServiceEnvironmentSettingGetResponseBinding interface {
	implementsWorkerServiceEnvironmentSettingGetResponseBinding()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkerServiceEnvironmentSettingGetResponseBinding)(nil)).Elem(), "")
}

type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersKvNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersKvNamespaceBindingType `json:"type,required"`
	JSON workerServiceEnvironmentSettingGetResponseBindingsWorkersKvNamespaceBindingJSON `json:"-"`
}

// workerServiceEnvironmentSettingGetResponseBindingsWorkersKvNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingGetResponseBindingsWorkersKvNamespaceBinding]
type workerServiceEnvironmentSettingGetResponseBindingsWorkersKvNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingGetResponseBindingsWorkersKvNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerServiceEnvironmentSettingGetResponseBindingsWorkersKvNamespaceBinding) implementsWorkerServiceEnvironmentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersKvNamespaceBindingType string

const (
	WorkerServiceEnvironmentSettingGetResponseBindingsWorkersKvNamespaceBindingTypeKvNamespace WorkerServiceEnvironmentSettingGetResponseBindingsWorkersKvNamespaceBindingType = "kv_namespace"
)

type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersServiceBindingType `json:"type,required"`
	JSON workerServiceEnvironmentSettingGetResponseBindingsWorkersServiceBindingJSON `json:"-"`
}

// workerServiceEnvironmentSettingGetResponseBindingsWorkersServiceBindingJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingGetResponseBindingsWorkersServiceBinding]
type workerServiceEnvironmentSettingGetResponseBindingsWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingGetResponseBindingsWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerServiceEnvironmentSettingGetResponseBindingsWorkersServiceBinding) implementsWorkerServiceEnvironmentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersServiceBindingType string

const (
	WorkerServiceEnvironmentSettingGetResponseBindingsWorkersServiceBindingTypeService WorkerServiceEnvironmentSettingGetResponseBindingsWorkersServiceBindingType = "service"
)

type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                                 `json:"script_name"`
	JSON       workerServiceEnvironmentSettingGetResponseBindingsWorkersDoBindingJSON `json:"-"`
}

// workerServiceEnvironmentSettingGetResponseBindingsWorkersDoBindingJSON contains
// the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingGetResponseBindingsWorkersDoBinding]
type workerServiceEnvironmentSettingGetResponseBindingsWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingGetResponseBindingsWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerServiceEnvironmentSettingGetResponseBindingsWorkersDoBinding) implementsWorkerServiceEnvironmentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersDoBindingType string

const (
	WorkerServiceEnvironmentSettingGetResponseBindingsWorkersDoBindingTypeDurableObjectNamespace WorkerServiceEnvironmentSettingGetResponseBindingsWorkersDoBindingType = "durable_object_namespace"
)

type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersR2BindingType `json:"type,required"`
	JSON workerServiceEnvironmentSettingGetResponseBindingsWorkersR2BindingJSON `json:"-"`
}

// workerServiceEnvironmentSettingGetResponseBindingsWorkersR2BindingJSON contains
// the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingGetResponseBindingsWorkersR2Binding]
type workerServiceEnvironmentSettingGetResponseBindingsWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingGetResponseBindingsWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerServiceEnvironmentSettingGetResponseBindingsWorkersR2Binding) implementsWorkerServiceEnvironmentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersR2BindingType string

const (
	WorkerServiceEnvironmentSettingGetResponseBindingsWorkersR2BindingTypeR2Bucket WorkerServiceEnvironmentSettingGetResponseBindingsWorkersR2BindingType = "r2_bucket"
)

type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersQueueBindingType `json:"type,required"`
	JSON workerServiceEnvironmentSettingGetResponseBindingsWorkersQueueBindingJSON `json:"-"`
}

// workerServiceEnvironmentSettingGetResponseBindingsWorkersQueueBindingJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingGetResponseBindingsWorkersQueueBinding]
type workerServiceEnvironmentSettingGetResponseBindingsWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingGetResponseBindingsWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerServiceEnvironmentSettingGetResponseBindingsWorkersQueueBinding) implementsWorkerServiceEnvironmentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersQueueBindingType string

const (
	WorkerServiceEnvironmentSettingGetResponseBindingsWorkersQueueBindingTypeQueue WorkerServiceEnvironmentSettingGetResponseBindingsWorkersQueueBindingType = "queue"
)

type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersD1BindingType `json:"type,required"`
	JSON workerServiceEnvironmentSettingGetResponseBindingsWorkersD1BindingJSON `json:"-"`
}

// workerServiceEnvironmentSettingGetResponseBindingsWorkersD1BindingJSON contains
// the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingGetResponseBindingsWorkersD1Binding]
type workerServiceEnvironmentSettingGetResponseBindingsWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingGetResponseBindingsWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerServiceEnvironmentSettingGetResponseBindingsWorkersD1Binding) implementsWorkerServiceEnvironmentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersD1BindingType string

const (
	WorkerServiceEnvironmentSettingGetResponseBindingsWorkersD1BindingTypeD1 WorkerServiceEnvironmentSettingGetResponseBindingsWorkersD1BindingType = "d1"
)

type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound WorkerServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     workerServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON     `json:"-"`
}

// workerServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBinding]
type workerServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBinding) implementsWorkerServiceEnvironmentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingType string

const (
	WorkerServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace WorkerServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker WorkerServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   workerServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// workerServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound]
type workerServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Outbound worker
type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                              `json:"service"`
	JSON    workerServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// workerServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker]
type workerServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersMtlsCertBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersMtlsCertBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                                       `json:"certificate_id"`
	JSON          workerServiceEnvironmentSettingGetResponseBindingsWorkersMtlsCertBindingJSON `json:"-"`
}

// workerServiceEnvironmentSettingGetResponseBindingsWorkersMtlsCertBindingJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingGetResponseBindingsWorkersMtlsCertBinding]
type workerServiceEnvironmentSettingGetResponseBindingsWorkersMtlsCertBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingGetResponseBindingsWorkersMtlsCertBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerServiceEnvironmentSettingGetResponseBindingsWorkersMtlsCertBinding) implementsWorkerServiceEnvironmentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersMtlsCertBindingType string

const (
	WorkerServiceEnvironmentSettingGetResponseBindingsWorkersMtlsCertBindingTypeMtlsCertificate WorkerServiceEnvironmentSettingGetResponseBindingsWorkersMtlsCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrations]
// or
// [WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrations].
type WorkerServiceEnvironmentSettingGetResponseMigrations interface {
	implementsWorkerServiceEnvironmentSettingGetResponseMigrations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkerServiceEnvironmentSettingGetResponseMigrations)(nil)).Elem(), "")
}

// A single set of migrations to apply.
type WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses []WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               workerServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON               `json:"-"`
}

// workerServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrations]
type workerServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrations) implementsWorkerServiceEnvironmentSettingGetResponseMigrations() {
}

type WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From string                                                                                          `json:"from"`
	To   string                                                                                          `json:"to"`
	JSON workerServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// workerServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass]
type workerServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       string                                                                                              `json:"from"`
	FromScript string                                                                                              `json:"from_script"`
	To         string                                                                                              `json:"to"`
	JSON       workerServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// workerServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass]
type workerServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStep `json:"steps"`
	JSON  workerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsJSON   `json:"-"`
}

// workerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrations]
type workerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrations) implementsWorkerServiceEnvironmentSettingGetResponseMigrations() {
}

type WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               workerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON                `json:"-"`
}

// workerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStep]
type workerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From string                                                                                            `json:"from"`
	To   string                                                                                            `json:"to"`
	JSON workerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// workerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass]
type workerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                                                `json:"from"`
	FromScript string                                                                                                `json:"from_script"`
	To         string                                                                                                `json:"to"`
	JSON       workerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// workerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass]
type workerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerServiceEnvironmentSettingGetResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode WorkerServiceEnvironmentSettingGetResponsePlacementMode `json:"mode"`
	JSON workerServiceEnvironmentSettingGetResponsePlacementJSON `json:"-"`
}

// workerServiceEnvironmentSettingGetResponsePlacementJSON contains the JSON
// metadata for the struct [WorkerServiceEnvironmentSettingGetResponsePlacement]
type workerServiceEnvironmentSettingGetResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingGetResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type WorkerServiceEnvironmentSettingGetResponsePlacementMode string

const (
	WorkerServiceEnvironmentSettingGetResponsePlacementModeSmart WorkerServiceEnvironmentSettingGetResponsePlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type WorkerServiceEnvironmentSettingGetResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                     `json:"namespace"`
	JSON      workerServiceEnvironmentSettingGetResponseTailConsumerJSON `json:"-"`
}

// workerServiceEnvironmentSettingGetResponseTailConsumerJSON contains the JSON
// metadata for the struct [WorkerServiceEnvironmentSettingGetResponseTailConsumer]
type workerServiceEnvironmentSettingGetResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingGetResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerServiceEnvironmentSettingModifyResponse struct {
	// List of bindings attached to this Worker
	Bindings []WorkerServiceEnvironmentSettingModifyResponseBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations WorkerServiceEnvironmentSettingModifyResponseMigrations `json:"migrations"`
	Placement  WorkerServiceEnvironmentSettingModifyResponsePlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []WorkerServiceEnvironmentSettingModifyResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                            `json:"usage_model"`
	JSON       workerServiceEnvironmentSettingModifyResponseJSON `json:"-"`
}

// workerServiceEnvironmentSettingModifyResponseJSON contains the JSON metadata for
// the struct [WorkerServiceEnvironmentSettingModifyResponse]
type workerServiceEnvironmentSettingModifyResponseJSON struct {
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

func (r *WorkerServiceEnvironmentSettingModifyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersKvNamespaceBinding],
// [WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersServiceBinding],
// [WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersDoBinding],
// [WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersR2Binding],
// [WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersQueueBinding],
// [WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersD1Binding],
// [WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersDispatchNamespaceBinding]
// or
// [WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersMtlsCertBinding].
type WorkerServiceEnvironmentSettingModifyResponseBinding interface {
	implementsWorkerServiceEnvironmentSettingModifyResponseBinding()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkerServiceEnvironmentSettingModifyResponseBinding)(nil)).Elem(), "")
}

type WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersKvNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersKvNamespaceBindingType `json:"type,required"`
	JSON workerServiceEnvironmentSettingModifyResponseBindingsWorkersKvNamespaceBindingJSON `json:"-"`
}

// workerServiceEnvironmentSettingModifyResponseBindingsWorkersKvNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersKvNamespaceBinding]
type workerServiceEnvironmentSettingModifyResponseBindingsWorkersKvNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersKvNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersKvNamespaceBinding) implementsWorkerServiceEnvironmentSettingModifyResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersKvNamespaceBindingType string

const (
	WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersKvNamespaceBindingTypeKvNamespace WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersKvNamespaceBindingType = "kv_namespace"
)

type WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersServiceBindingType `json:"type,required"`
	JSON workerServiceEnvironmentSettingModifyResponseBindingsWorkersServiceBindingJSON `json:"-"`
}

// workerServiceEnvironmentSettingModifyResponseBindingsWorkersServiceBindingJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersServiceBinding]
type workerServiceEnvironmentSettingModifyResponseBindingsWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersServiceBinding) implementsWorkerServiceEnvironmentSettingModifyResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersServiceBindingType string

const (
	WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersServiceBindingTypeService WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersServiceBindingType = "service"
)

type WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                                    `json:"script_name"`
	JSON       workerServiceEnvironmentSettingModifyResponseBindingsWorkersDoBindingJSON `json:"-"`
}

// workerServiceEnvironmentSettingModifyResponseBindingsWorkersDoBindingJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersDoBinding]
type workerServiceEnvironmentSettingModifyResponseBindingsWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersDoBinding) implementsWorkerServiceEnvironmentSettingModifyResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersDoBindingType string

const (
	WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersDoBindingTypeDurableObjectNamespace WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersDoBindingType = "durable_object_namespace"
)

type WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersR2BindingType `json:"type,required"`
	JSON workerServiceEnvironmentSettingModifyResponseBindingsWorkersR2BindingJSON `json:"-"`
}

// workerServiceEnvironmentSettingModifyResponseBindingsWorkersR2BindingJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersR2Binding]
type workerServiceEnvironmentSettingModifyResponseBindingsWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersR2Binding) implementsWorkerServiceEnvironmentSettingModifyResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersR2BindingType string

const (
	WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersR2BindingTypeR2Bucket WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersR2BindingType = "r2_bucket"
)

type WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersQueueBindingType `json:"type,required"`
	JSON workerServiceEnvironmentSettingModifyResponseBindingsWorkersQueueBindingJSON `json:"-"`
}

// workerServiceEnvironmentSettingModifyResponseBindingsWorkersQueueBindingJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersQueueBinding]
type workerServiceEnvironmentSettingModifyResponseBindingsWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersQueueBinding) implementsWorkerServiceEnvironmentSettingModifyResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersQueueBindingType string

const (
	WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersQueueBindingTypeQueue WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersQueueBindingType = "queue"
)

type WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersD1BindingType `json:"type,required"`
	JSON workerServiceEnvironmentSettingModifyResponseBindingsWorkersD1BindingJSON `json:"-"`
}

// workerServiceEnvironmentSettingModifyResponseBindingsWorkersD1BindingJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersD1Binding]
type workerServiceEnvironmentSettingModifyResponseBindingsWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersD1Binding) implementsWorkerServiceEnvironmentSettingModifyResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersD1BindingType string

const (
	WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersD1BindingTypeD1 WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersD1BindingType = "d1"
)

type WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     workerServiceEnvironmentSettingModifyResponseBindingsWorkersDispatchNamespaceBindingJSON     `json:"-"`
}

// workerServiceEnvironmentSettingModifyResponseBindingsWorkersDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersDispatchNamespaceBinding]
type workerServiceEnvironmentSettingModifyResponseBindingsWorkersDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersDispatchNamespaceBinding) implementsWorkerServiceEnvironmentSettingModifyResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersDispatchNamespaceBindingType string

const (
	WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   workerServiceEnvironmentSettingModifyResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// workerServiceEnvironmentSettingModifyResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersDispatchNamespaceBindingOutbound]
type workerServiceEnvironmentSettingModifyResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Outbound worker
type WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                                 `json:"service"`
	JSON    workerServiceEnvironmentSettingModifyResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// workerServiceEnvironmentSettingModifyResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker]
type workerServiceEnvironmentSettingModifyResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersMtlsCertBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersMtlsCertBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                                          `json:"certificate_id"`
	JSON          workerServiceEnvironmentSettingModifyResponseBindingsWorkersMtlsCertBindingJSON `json:"-"`
}

// workerServiceEnvironmentSettingModifyResponseBindingsWorkersMtlsCertBindingJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersMtlsCertBinding]
type workerServiceEnvironmentSettingModifyResponseBindingsWorkersMtlsCertBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersMtlsCertBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersMtlsCertBinding) implementsWorkerServiceEnvironmentSettingModifyResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersMtlsCertBindingType string

const (
	WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersMtlsCertBindingTypeMtlsCertificate WorkerServiceEnvironmentSettingModifyResponseBindingsWorkersMtlsCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [WorkerServiceEnvironmentSettingModifyResponseMigrationsWorkersSingleStepMigrations]
// or
// [WorkerServiceEnvironmentSettingModifyResponseMigrationsWorkersSteppedMigrations].
type WorkerServiceEnvironmentSettingModifyResponseMigrations interface {
	implementsWorkerServiceEnvironmentSettingModifyResponseMigrations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkerServiceEnvironmentSettingModifyResponseMigrations)(nil)).Elem(), "")
}

// A single set of migrations to apply.
type WorkerServiceEnvironmentSettingModifyResponseMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses []WorkerServiceEnvironmentSettingModifyResponseMigrationsWorkersSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkerServiceEnvironmentSettingModifyResponseMigrationsWorkersSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               workerServiceEnvironmentSettingModifyResponseMigrationsWorkersSingleStepMigrationsJSON               `json:"-"`
}

// workerServiceEnvironmentSettingModifyResponseMigrationsWorkersSingleStepMigrationsJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingModifyResponseMigrationsWorkersSingleStepMigrations]
type workerServiceEnvironmentSettingModifyResponseMigrationsWorkersSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingModifyResponseMigrationsWorkersSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerServiceEnvironmentSettingModifyResponseMigrationsWorkersSingleStepMigrations) implementsWorkerServiceEnvironmentSettingModifyResponseMigrations() {
}

type WorkerServiceEnvironmentSettingModifyResponseMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From string                                                                                             `json:"from"`
	To   string                                                                                             `json:"to"`
	JSON workerServiceEnvironmentSettingModifyResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// workerServiceEnvironmentSettingModifyResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingModifyResponseMigrationsWorkersSingleStepMigrationsRenamedClass]
type workerServiceEnvironmentSettingModifyResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingModifyResponseMigrationsWorkersSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerServiceEnvironmentSettingModifyResponseMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       string                                                                                                 `json:"from"`
	FromScript string                                                                                                 `json:"from_script"`
	To         string                                                                                                 `json:"to"`
	JSON       workerServiceEnvironmentSettingModifyResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// workerServiceEnvironmentSettingModifyResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingModifyResponseMigrationsWorkersSingleStepMigrationsTransferredClass]
type workerServiceEnvironmentSettingModifyResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingModifyResponseMigrationsWorkersSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerServiceEnvironmentSettingModifyResponseMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []WorkerServiceEnvironmentSettingModifyResponseMigrationsWorkersSteppedMigrationsStep `json:"steps"`
	JSON  workerServiceEnvironmentSettingModifyResponseMigrationsWorkersSteppedMigrationsJSON   `json:"-"`
}

// workerServiceEnvironmentSettingModifyResponseMigrationsWorkersSteppedMigrationsJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingModifyResponseMigrationsWorkersSteppedMigrations]
type workerServiceEnvironmentSettingModifyResponseMigrationsWorkersSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingModifyResponseMigrationsWorkersSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerServiceEnvironmentSettingModifyResponseMigrationsWorkersSteppedMigrations) implementsWorkerServiceEnvironmentSettingModifyResponseMigrations() {
}

type WorkerServiceEnvironmentSettingModifyResponseMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []WorkerServiceEnvironmentSettingModifyResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkerServiceEnvironmentSettingModifyResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               workerServiceEnvironmentSettingModifyResponseMigrationsWorkersSteppedMigrationsStepJSON                `json:"-"`
}

// workerServiceEnvironmentSettingModifyResponseMigrationsWorkersSteppedMigrationsStepJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingModifyResponseMigrationsWorkersSteppedMigrationsStep]
type workerServiceEnvironmentSettingModifyResponseMigrationsWorkersSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingModifyResponseMigrationsWorkersSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerServiceEnvironmentSettingModifyResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From string                                                                                               `json:"from"`
	To   string                                                                                               `json:"to"`
	JSON workerServiceEnvironmentSettingModifyResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// workerServiceEnvironmentSettingModifyResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingModifyResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass]
type workerServiceEnvironmentSettingModifyResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingModifyResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerServiceEnvironmentSettingModifyResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                                                   `json:"from"`
	FromScript string                                                                                                   `json:"from_script"`
	To         string                                                                                                   `json:"to"`
	JSON       workerServiceEnvironmentSettingModifyResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// workerServiceEnvironmentSettingModifyResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingModifyResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass]
type workerServiceEnvironmentSettingModifyResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingModifyResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerServiceEnvironmentSettingModifyResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode WorkerServiceEnvironmentSettingModifyResponsePlacementMode `json:"mode"`
	JSON workerServiceEnvironmentSettingModifyResponsePlacementJSON `json:"-"`
}

// workerServiceEnvironmentSettingModifyResponsePlacementJSON contains the JSON
// metadata for the struct [WorkerServiceEnvironmentSettingModifyResponsePlacement]
type workerServiceEnvironmentSettingModifyResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingModifyResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type WorkerServiceEnvironmentSettingModifyResponsePlacementMode string

const (
	WorkerServiceEnvironmentSettingModifyResponsePlacementModeSmart WorkerServiceEnvironmentSettingModifyResponsePlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type WorkerServiceEnvironmentSettingModifyResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                        `json:"namespace"`
	JSON      workerServiceEnvironmentSettingModifyResponseTailConsumerJSON `json:"-"`
}

// workerServiceEnvironmentSettingModifyResponseTailConsumerJSON contains the JSON
// metadata for the struct
// [WorkerServiceEnvironmentSettingModifyResponseTailConsumer]
type workerServiceEnvironmentSettingModifyResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingModifyResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerServiceEnvironmentSettingGetResponseEnvelope struct {
	Errors   []WorkerServiceEnvironmentSettingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerServiceEnvironmentSettingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerServiceEnvironmentSettingGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerServiceEnvironmentSettingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerServiceEnvironmentSettingGetResponseEnvelopeJSON    `json:"-"`
}

// workerServiceEnvironmentSettingGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [WorkerServiceEnvironmentSettingGetResponseEnvelope]
type workerServiceEnvironmentSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerServiceEnvironmentSettingGetResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    workerServiceEnvironmentSettingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// workerServiceEnvironmentSettingGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [WorkerServiceEnvironmentSettingGetResponseEnvelopeErrors]
type workerServiceEnvironmentSettingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerServiceEnvironmentSettingGetResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    workerServiceEnvironmentSettingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// workerServiceEnvironmentSettingGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [WorkerServiceEnvironmentSettingGetResponseEnvelopeMessages]
type workerServiceEnvironmentSettingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerServiceEnvironmentSettingGetResponseEnvelopeSuccess bool

const (
	WorkerServiceEnvironmentSettingGetResponseEnvelopeSuccessTrue WorkerServiceEnvironmentSettingGetResponseEnvelopeSuccess = true
)

type WorkerServiceEnvironmentSettingModifyParams struct {
	Errors   param.Field[[]WorkerServiceEnvironmentSettingModifyParamsError]   `json:"errors,required"`
	Messages param.Field[[]WorkerServiceEnvironmentSettingModifyParamsMessage] `json:"messages,required"`
	Result   param.Field[WorkerServiceEnvironmentSettingModifyParamsResult]    `json:"result,required"`
	// Whether the API call was successful
	Success param.Field[WorkerServiceEnvironmentSettingModifyParamsSuccess] `json:"success,required"`
}

func (r WorkerServiceEnvironmentSettingModifyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerServiceEnvironmentSettingModifyParamsError struct {
	Code    param.Field[int64]  `json:"code,required"`
	Message param.Field[string] `json:"message,required"`
}

func (r WorkerServiceEnvironmentSettingModifyParamsError) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerServiceEnvironmentSettingModifyParamsMessage struct {
	Code    param.Field[int64]  `json:"code,required"`
	Message param.Field[string] `json:"message,required"`
}

func (r WorkerServiceEnvironmentSettingModifyParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerServiceEnvironmentSettingModifyParamsResult struct {
	// List of bindings attached to this Worker
	Bindings param.Field[[]WorkerServiceEnvironmentSettingModifyParamsResultBinding] `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations param.Field[WorkerServiceEnvironmentSettingModifyParamsResultMigrations] `json:"migrations"`
	Placement  param.Field[WorkerServiceEnvironmentSettingModifyParamsResultPlacement]  `json:"placement"`
	// Tags to help you manage your Workers
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]WorkerServiceEnvironmentSettingModifyParamsResultTailConsumer] `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel param.Field[string] `json:"usage_model"`
}

func (r WorkerServiceEnvironmentSettingModifyParamsResult) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A binding to allow the Worker to communicate with resources
//
// Satisfied by
// [WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersKvNamespaceBinding],
// [WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersServiceBinding],
// [WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersDoBinding],
// [WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersR2Binding],
// [WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersQueueBinding],
// [WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersD1Binding],
// [WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersDispatchNamespaceBinding],
// [WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersMtlsCertBinding].
type WorkerServiceEnvironmentSettingModifyParamsResultBinding interface {
	implementsWorkerServiceEnvironmentSettingModifyParamsResultBinding()
}

type WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersKvNamespaceBinding struct {
	// The class of resource that the binding provides.
	Type param.Field[WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersKvNamespaceBindingType] `json:"type,required"`
}

func (r WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersKvNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersKvNamespaceBinding) implementsWorkerServiceEnvironmentSettingModifyParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersKvNamespaceBindingType string

const (
	WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersKvNamespaceBindingTypeKvNamespace WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersKvNamespaceBindingType = "kv_namespace"
)

type WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// Name of Worker to bind to
	Service param.Field[string] `json:"service,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersServiceBindingType] `json:"type,required"`
}

func (r WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersServiceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersServiceBinding) implementsWorkerServiceEnvironmentSettingModifyParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersServiceBindingType string

const (
	WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersServiceBindingTypeService WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersServiceBindingType = "service"
)

type WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName param.Field[string] `json:"class_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersDoBindingType] `json:"type,required"`
	// The environment of the script_name to bind to
	Environment param.Field[string] `json:"environment"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName param.Field[string] `json:"script_name"`
}

func (r WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersDoBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersDoBinding) implementsWorkerServiceEnvironmentSettingModifyParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersDoBindingType string

const (
	WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersDoBindingTypeDurableObjectNamespace WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersDoBindingType = "durable_object_namespace"
)

type WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName param.Field[string] `json:"bucket_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersR2BindingType] `json:"type,required"`
}

func (r WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersR2Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersR2Binding) implementsWorkerServiceEnvironmentSettingModifyParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersR2BindingType string

const (
	WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersR2BindingTypeR2Bucket WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersR2BindingType = "r2_bucket"
)

type WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersQueueBinding struct {
	// Name of the Queue to bind to
	QueueName param.Field[string] `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersQueueBindingType] `json:"type,required"`
}

func (r WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersQueueBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersQueueBinding) implementsWorkerServiceEnvironmentSettingModifyParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersQueueBindingType string

const (
	WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersQueueBindingTypeQueue WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersQueueBindingType = "queue"
)

type WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID param.Field[string] `json:"id,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name param.Field[string] `json:"name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersD1BindingType] `json:"type,required"`
}

func (r WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersD1Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersD1Binding) implementsWorkerServiceEnvironmentSettingModifyParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersD1BindingType string

const (
	WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersD1BindingTypeD1 WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersD1BindingType = "d1"
)

type WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersDispatchNamespaceBinding struct {
	// Namespace to bind to
	Namespace param.Field[string] `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersDispatchNamespaceBindingType] `json:"type,required"`
	// Outbound worker
	Outbound param.Field[WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersDispatchNamespaceBindingOutbound] `json:"outbound"`
}

func (r WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersDispatchNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersDispatchNamespaceBinding) implementsWorkerServiceEnvironmentSettingModifyParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersDispatchNamespaceBindingType string

const (
	WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params param.Field[[]string] `json:"params"`
	// Outbound worker
	Worker param.Field[WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker] `json:"worker"`
}

func (r WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersDispatchNamespaceBindingOutbound) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker
type WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker
	Service param.Field[string] `json:"service"`
}

func (r WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersMtlsCertBinding struct {
	// The class of resource that the binding provides.
	Type param.Field[WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersMtlsCertBindingType] `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID param.Field[string] `json:"certificate_id"`
}

func (r WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersMtlsCertBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersMtlsCertBinding) implementsWorkerServiceEnvironmentSettingModifyParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersMtlsCertBindingType string

const (
	WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersMtlsCertBindingTypeMtlsCertificate WorkerServiceEnvironmentSettingModifyParamsResultBindingsWorkersMtlsCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by
// [WorkerServiceEnvironmentSettingModifyParamsResultMigrationsWorkersSingleStepMigrations],
// [WorkerServiceEnvironmentSettingModifyParamsResultMigrationsWorkersSteppedMigrations].
type WorkerServiceEnvironmentSettingModifyParamsResultMigrations interface {
	implementsWorkerServiceEnvironmentSettingModifyParamsResultMigrations()
}

// A single set of migrations to apply.
type WorkerServiceEnvironmentSettingModifyParamsResultMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses param.Field[[]WorkerServiceEnvironmentSettingModifyParamsResultMigrationsWorkersSingleStepMigrationsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]WorkerServiceEnvironmentSettingModifyParamsResultMigrationsWorkersSingleStepMigrationsTransferredClass] `json:"transferred_classes"`
}

func (r WorkerServiceEnvironmentSettingModifyParamsResultMigrationsWorkersSingleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerServiceEnvironmentSettingModifyParamsResultMigrationsWorkersSingleStepMigrations) implementsWorkerServiceEnvironmentSettingModifyParamsResultMigrations() {
}

type WorkerServiceEnvironmentSettingModifyParamsResultMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r WorkerServiceEnvironmentSettingModifyParamsResultMigrationsWorkersSingleStepMigrationsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerServiceEnvironmentSettingModifyParamsResultMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r WorkerServiceEnvironmentSettingModifyParamsResultMigrationsWorkersSingleStepMigrationsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerServiceEnvironmentSettingModifyParamsResultMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]WorkerServiceEnvironmentSettingModifyParamsResultMigrationsWorkersSteppedMigrationsStep] `json:"steps"`
}

func (r WorkerServiceEnvironmentSettingModifyParamsResultMigrationsWorkersSteppedMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerServiceEnvironmentSettingModifyParamsResultMigrationsWorkersSteppedMigrations) implementsWorkerServiceEnvironmentSettingModifyParamsResultMigrations() {
}

type WorkerServiceEnvironmentSettingModifyParamsResultMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]WorkerServiceEnvironmentSettingModifyParamsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]WorkerServiceEnvironmentSettingModifyParamsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass] `json:"transferred_classes"`
}

func (r WorkerServiceEnvironmentSettingModifyParamsResultMigrationsWorkersSteppedMigrationsStep) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerServiceEnvironmentSettingModifyParamsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r WorkerServiceEnvironmentSettingModifyParamsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerServiceEnvironmentSettingModifyParamsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r WorkerServiceEnvironmentSettingModifyParamsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerServiceEnvironmentSettingModifyParamsResultPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode param.Field[WorkerServiceEnvironmentSettingModifyParamsResultPlacementMode] `json:"mode"`
}

func (r WorkerServiceEnvironmentSettingModifyParamsResultPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type WorkerServiceEnvironmentSettingModifyParamsResultPlacementMode string

const (
	WorkerServiceEnvironmentSettingModifyParamsResultPlacementModeSmart WorkerServiceEnvironmentSettingModifyParamsResultPlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type WorkerServiceEnvironmentSettingModifyParamsResultTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service param.Field[string] `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace param.Field[string] `json:"namespace"`
}

func (r WorkerServiceEnvironmentSettingModifyParamsResultTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the API call was successful
type WorkerServiceEnvironmentSettingModifyParamsSuccess bool

const (
	WorkerServiceEnvironmentSettingModifyParamsSuccessTrue WorkerServiceEnvironmentSettingModifyParamsSuccess = true
)

type WorkerServiceEnvironmentSettingModifyResponseEnvelope struct {
	Errors   []WorkerServiceEnvironmentSettingModifyResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerServiceEnvironmentSettingModifyResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerServiceEnvironmentSettingModifyResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerServiceEnvironmentSettingModifyResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerServiceEnvironmentSettingModifyResponseEnvelopeJSON    `json:"-"`
}

// workerServiceEnvironmentSettingModifyResponseEnvelopeJSON contains the JSON
// metadata for the struct [WorkerServiceEnvironmentSettingModifyResponseEnvelope]
type workerServiceEnvironmentSettingModifyResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingModifyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerServiceEnvironmentSettingModifyResponseEnvelopeErrors struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    workerServiceEnvironmentSettingModifyResponseEnvelopeErrorsJSON `json:"-"`
}

// workerServiceEnvironmentSettingModifyResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [WorkerServiceEnvironmentSettingModifyResponseEnvelopeErrors]
type workerServiceEnvironmentSettingModifyResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingModifyResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerServiceEnvironmentSettingModifyResponseEnvelopeMessages struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    workerServiceEnvironmentSettingModifyResponseEnvelopeMessagesJSON `json:"-"`
}

// workerServiceEnvironmentSettingModifyResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [WorkerServiceEnvironmentSettingModifyResponseEnvelopeMessages]
type workerServiceEnvironmentSettingModifyResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingModifyResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerServiceEnvironmentSettingModifyResponseEnvelopeSuccess bool

const (
	WorkerServiceEnvironmentSettingModifyResponseEnvelopeSuccessTrue WorkerServiceEnvironmentSettingModifyResponseEnvelopeSuccess = true
)
