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

// Patch script metadata, such as bindings
func (r *WorkerServiceEnvironmentSettingService) Update(ctx context.Context, accountID string, serviceName string, environmentName string, body WorkerServiceEnvironmentSettingUpdateParams, opts ...option.RequestOption) (res *WorkerServiceEnvironmentSettingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerServiceEnvironmentSettingUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/services/%s/environments/%s/settings", accountID, serviceName, environmentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
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

type WorkerServiceEnvironmentSettingUpdateResponse struct {
	// List of bindings attached to this Worker
	Bindings []WorkerServiceEnvironmentSettingUpdateResponseBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations WorkerServiceEnvironmentSettingUpdateResponseMigrations `json:"migrations"`
	Placement  WorkerServiceEnvironmentSettingUpdateResponsePlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []WorkerServiceEnvironmentSettingUpdateResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                            `json:"usage_model"`
	JSON       workerServiceEnvironmentSettingUpdateResponseJSON `json:"-"`
}

// workerServiceEnvironmentSettingUpdateResponseJSON contains the JSON metadata for
// the struct [WorkerServiceEnvironmentSettingUpdateResponse]
type workerServiceEnvironmentSettingUpdateResponseJSON struct {
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

func (r *WorkerServiceEnvironmentSettingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersKvNamespaceBinding],
// [WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersServiceBinding],
// [WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersDoBinding],
// [WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersR2Binding],
// [WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersQueueBinding],
// [WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersD1Binding],
// [WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersDispatchNamespaceBinding]
// or
// [WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersMtlsCertBinding].
type WorkerServiceEnvironmentSettingUpdateResponseBinding interface {
	implementsWorkerServiceEnvironmentSettingUpdateResponseBinding()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkerServiceEnvironmentSettingUpdateResponseBinding)(nil)).Elem(), "")
}

type WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersKvNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersKvNamespaceBindingType `json:"type,required"`
	JSON workerServiceEnvironmentSettingUpdateResponseBindingsWorkersKvNamespaceBindingJSON `json:"-"`
}

// workerServiceEnvironmentSettingUpdateResponseBindingsWorkersKvNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersKvNamespaceBinding]
type workerServiceEnvironmentSettingUpdateResponseBindingsWorkersKvNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersKvNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersKvNamespaceBinding) implementsWorkerServiceEnvironmentSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersKvNamespaceBindingType string

const (
	WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersKvNamespaceBindingTypeKvNamespace WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersKvNamespaceBindingType = "kv_namespace"
)

type WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersServiceBindingType `json:"type,required"`
	JSON workerServiceEnvironmentSettingUpdateResponseBindingsWorkersServiceBindingJSON `json:"-"`
}

// workerServiceEnvironmentSettingUpdateResponseBindingsWorkersServiceBindingJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersServiceBinding]
type workerServiceEnvironmentSettingUpdateResponseBindingsWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersServiceBinding) implementsWorkerServiceEnvironmentSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersServiceBindingType string

const (
	WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersServiceBindingTypeService WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersServiceBindingType = "service"
)

type WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                                    `json:"script_name"`
	JSON       workerServiceEnvironmentSettingUpdateResponseBindingsWorkersDoBindingJSON `json:"-"`
}

// workerServiceEnvironmentSettingUpdateResponseBindingsWorkersDoBindingJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersDoBinding]
type workerServiceEnvironmentSettingUpdateResponseBindingsWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersDoBinding) implementsWorkerServiceEnvironmentSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersDoBindingType string

const (
	WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersDoBindingTypeDurableObjectNamespace WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersDoBindingType = "durable_object_namespace"
)

type WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersR2BindingType `json:"type,required"`
	JSON workerServiceEnvironmentSettingUpdateResponseBindingsWorkersR2BindingJSON `json:"-"`
}

// workerServiceEnvironmentSettingUpdateResponseBindingsWorkersR2BindingJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersR2Binding]
type workerServiceEnvironmentSettingUpdateResponseBindingsWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersR2Binding) implementsWorkerServiceEnvironmentSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersR2BindingType string

const (
	WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersR2BindingTypeR2Bucket WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersR2BindingType = "r2_bucket"
)

type WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersQueueBindingType `json:"type,required"`
	JSON workerServiceEnvironmentSettingUpdateResponseBindingsWorkersQueueBindingJSON `json:"-"`
}

// workerServiceEnvironmentSettingUpdateResponseBindingsWorkersQueueBindingJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersQueueBinding]
type workerServiceEnvironmentSettingUpdateResponseBindingsWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersQueueBinding) implementsWorkerServiceEnvironmentSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersQueueBindingType string

const (
	WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersQueueBindingTypeQueue WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersQueueBindingType = "queue"
)

type WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersD1BindingType `json:"type,required"`
	JSON workerServiceEnvironmentSettingUpdateResponseBindingsWorkersD1BindingJSON `json:"-"`
}

// workerServiceEnvironmentSettingUpdateResponseBindingsWorkersD1BindingJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersD1Binding]
type workerServiceEnvironmentSettingUpdateResponseBindingsWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersD1Binding) implementsWorkerServiceEnvironmentSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersD1BindingType string

const (
	WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersD1BindingTypeD1 WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersD1BindingType = "d1"
)

type WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     workerServiceEnvironmentSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingJSON     `json:"-"`
}

// workerServiceEnvironmentSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersDispatchNamespaceBinding]
type workerServiceEnvironmentSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersDispatchNamespaceBinding) implementsWorkerServiceEnvironmentSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingType string

const (
	WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   workerServiceEnvironmentSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// workerServiceEnvironmentSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutbound]
type workerServiceEnvironmentSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Outbound worker
type WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                                 `json:"service"`
	JSON    workerServiceEnvironmentSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// workerServiceEnvironmentSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker]
type workerServiceEnvironmentSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersMtlsCertBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersMtlsCertBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                                          `json:"certificate_id"`
	JSON          workerServiceEnvironmentSettingUpdateResponseBindingsWorkersMtlsCertBindingJSON `json:"-"`
}

// workerServiceEnvironmentSettingUpdateResponseBindingsWorkersMtlsCertBindingJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersMtlsCertBinding]
type workerServiceEnvironmentSettingUpdateResponseBindingsWorkersMtlsCertBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersMtlsCertBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersMtlsCertBinding) implementsWorkerServiceEnvironmentSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersMtlsCertBindingType string

const (
	WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersMtlsCertBindingTypeMtlsCertificate WorkerServiceEnvironmentSettingUpdateResponseBindingsWorkersMtlsCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [WorkerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSingleStepMigrations]
// or
// [WorkerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSteppedMigrations].
type WorkerServiceEnvironmentSettingUpdateResponseMigrations interface {
	implementsWorkerServiceEnvironmentSettingUpdateResponseMigrations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkerServiceEnvironmentSettingUpdateResponseMigrations)(nil)).Elem(), "")
}

// A single set of migrations to apply.
type WorkerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses []WorkerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               workerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSingleStepMigrationsJSON               `json:"-"`
}

// workerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSingleStepMigrationsJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSingleStepMigrations]
type workerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSingleStepMigrations) implementsWorkerServiceEnvironmentSettingUpdateResponseMigrations() {
}

type WorkerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From string                                                                                             `json:"from"`
	To   string                                                                                             `json:"to"`
	JSON workerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// workerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSingleStepMigrationsRenamedClass]
type workerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       string                                                                                                 `json:"from"`
	FromScript string                                                                                                 `json:"from_script"`
	To         string                                                                                                 `json:"to"`
	JSON       workerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// workerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSingleStepMigrationsTransferredClass]
type workerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []WorkerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSteppedMigrationsStep `json:"steps"`
	JSON  workerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSteppedMigrationsJSON   `json:"-"`
}

// workerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSteppedMigrationsJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSteppedMigrations]
type workerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSteppedMigrations) implementsWorkerServiceEnvironmentSettingUpdateResponseMigrations() {
}

type WorkerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []WorkerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               workerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepJSON                `json:"-"`
}

// workerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSteppedMigrationsStep]
type workerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From string                                                                                               `json:"from"`
	To   string                                                                                               `json:"to"`
	JSON workerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// workerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass]
type workerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                                                   `json:"from"`
	FromScript string                                                                                                   `json:"from_script"`
	To         string                                                                                                   `json:"to"`
	JSON       workerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// workerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass]
type workerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerServiceEnvironmentSettingUpdateResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode WorkerServiceEnvironmentSettingUpdateResponsePlacementMode `json:"mode"`
	JSON workerServiceEnvironmentSettingUpdateResponsePlacementJSON `json:"-"`
}

// workerServiceEnvironmentSettingUpdateResponsePlacementJSON contains the JSON
// metadata for the struct [WorkerServiceEnvironmentSettingUpdateResponsePlacement]
type workerServiceEnvironmentSettingUpdateResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingUpdateResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type WorkerServiceEnvironmentSettingUpdateResponsePlacementMode string

const (
	WorkerServiceEnvironmentSettingUpdateResponsePlacementModeSmart WorkerServiceEnvironmentSettingUpdateResponsePlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type WorkerServiceEnvironmentSettingUpdateResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                        `json:"namespace"`
	JSON      workerServiceEnvironmentSettingUpdateResponseTailConsumerJSON `json:"-"`
}

// workerServiceEnvironmentSettingUpdateResponseTailConsumerJSON contains the JSON
// metadata for the struct
// [WorkerServiceEnvironmentSettingUpdateResponseTailConsumer]
type workerServiceEnvironmentSettingUpdateResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingUpdateResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type WorkerServiceEnvironmentSettingUpdateParams struct {
	Errors   param.Field[[]WorkerServiceEnvironmentSettingUpdateParamsError]   `json:"errors,required"`
	Messages param.Field[[]WorkerServiceEnvironmentSettingUpdateParamsMessage] `json:"messages,required"`
	Result   param.Field[WorkerServiceEnvironmentSettingUpdateParamsResult]    `json:"result,required"`
	// Whether the API call was successful
	Success param.Field[WorkerServiceEnvironmentSettingUpdateParamsSuccess] `json:"success,required"`
}

func (r WorkerServiceEnvironmentSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerServiceEnvironmentSettingUpdateParamsError struct {
	Code    param.Field[int64]  `json:"code,required"`
	Message param.Field[string] `json:"message,required"`
}

func (r WorkerServiceEnvironmentSettingUpdateParamsError) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerServiceEnvironmentSettingUpdateParamsMessage struct {
	Code    param.Field[int64]  `json:"code,required"`
	Message param.Field[string] `json:"message,required"`
}

func (r WorkerServiceEnvironmentSettingUpdateParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerServiceEnvironmentSettingUpdateParamsResult struct {
	// List of bindings attached to this Worker
	Bindings param.Field[[]WorkerServiceEnvironmentSettingUpdateParamsResultBinding] `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations param.Field[WorkerServiceEnvironmentSettingUpdateParamsResultMigrations] `json:"migrations"`
	Placement  param.Field[WorkerServiceEnvironmentSettingUpdateParamsResultPlacement]  `json:"placement"`
	// Tags to help you manage your Workers
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]WorkerServiceEnvironmentSettingUpdateParamsResultTailConsumer] `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel param.Field[string] `json:"usage_model"`
}

func (r WorkerServiceEnvironmentSettingUpdateParamsResult) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A binding to allow the Worker to communicate with resources
//
// Satisfied by
// [WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersKvNamespaceBinding],
// [WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersServiceBinding],
// [WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersDoBinding],
// [WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersR2Binding],
// [WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersQueueBinding],
// [WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersD1Binding],
// [WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersDispatchNamespaceBinding],
// [WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersMtlsCertBinding].
type WorkerServiceEnvironmentSettingUpdateParamsResultBinding interface {
	implementsWorkerServiceEnvironmentSettingUpdateParamsResultBinding()
}

type WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersKvNamespaceBinding struct {
	// The class of resource that the binding provides.
	Type param.Field[WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersKvNamespaceBindingType] `json:"type,required"`
}

func (r WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersKvNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersKvNamespaceBinding) implementsWorkerServiceEnvironmentSettingUpdateParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersKvNamespaceBindingType string

const (
	WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersKvNamespaceBindingTypeKvNamespace WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersKvNamespaceBindingType = "kv_namespace"
)

type WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// Name of Worker to bind to
	Service param.Field[string] `json:"service,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersServiceBindingType] `json:"type,required"`
}

func (r WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersServiceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersServiceBinding) implementsWorkerServiceEnvironmentSettingUpdateParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersServiceBindingType string

const (
	WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersServiceBindingTypeService WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersServiceBindingType = "service"
)

type WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName param.Field[string] `json:"class_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersDoBindingType] `json:"type,required"`
	// The environment of the script_name to bind to
	Environment param.Field[string] `json:"environment"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName param.Field[string] `json:"script_name"`
}

func (r WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersDoBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersDoBinding) implementsWorkerServiceEnvironmentSettingUpdateParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersDoBindingType string

const (
	WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersDoBindingTypeDurableObjectNamespace WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersDoBindingType = "durable_object_namespace"
)

type WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName param.Field[string] `json:"bucket_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersR2BindingType] `json:"type,required"`
}

func (r WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersR2Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersR2Binding) implementsWorkerServiceEnvironmentSettingUpdateParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersR2BindingType string

const (
	WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersR2BindingTypeR2Bucket WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersR2BindingType = "r2_bucket"
)

type WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersQueueBinding struct {
	// Name of the Queue to bind to
	QueueName param.Field[string] `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersQueueBindingType] `json:"type,required"`
}

func (r WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersQueueBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersQueueBinding) implementsWorkerServiceEnvironmentSettingUpdateParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersQueueBindingType string

const (
	WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersQueueBindingTypeQueue WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersQueueBindingType = "queue"
)

type WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID param.Field[string] `json:"id,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name param.Field[string] `json:"name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersD1BindingType] `json:"type,required"`
}

func (r WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersD1Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersD1Binding) implementsWorkerServiceEnvironmentSettingUpdateParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersD1BindingType string

const (
	WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersD1BindingTypeD1 WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersD1BindingType = "d1"
)

type WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersDispatchNamespaceBinding struct {
	// Namespace to bind to
	Namespace param.Field[string] `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersDispatchNamespaceBindingType] `json:"type,required"`
	// Outbound worker
	Outbound param.Field[WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersDispatchNamespaceBindingOutbound] `json:"outbound"`
}

func (r WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersDispatchNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersDispatchNamespaceBinding) implementsWorkerServiceEnvironmentSettingUpdateParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersDispatchNamespaceBindingType string

const (
	WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params param.Field[[]string] `json:"params"`
	// Outbound worker
	Worker param.Field[WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker] `json:"worker"`
}

func (r WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersDispatchNamespaceBindingOutbound) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker
type WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker
	Service param.Field[string] `json:"service"`
}

func (r WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersMtlsCertBinding struct {
	// The class of resource that the binding provides.
	Type param.Field[WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersMtlsCertBindingType] `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID param.Field[string] `json:"certificate_id"`
}

func (r WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersMtlsCertBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersMtlsCertBinding) implementsWorkerServiceEnvironmentSettingUpdateParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersMtlsCertBindingType string

const (
	WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersMtlsCertBindingTypeMtlsCertificate WorkerServiceEnvironmentSettingUpdateParamsResultBindingsWorkersMtlsCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by
// [WorkerServiceEnvironmentSettingUpdateParamsResultMigrationsWorkersSingleStepMigrations],
// [WorkerServiceEnvironmentSettingUpdateParamsResultMigrationsWorkersSteppedMigrations].
type WorkerServiceEnvironmentSettingUpdateParamsResultMigrations interface {
	implementsWorkerServiceEnvironmentSettingUpdateParamsResultMigrations()
}

// A single set of migrations to apply.
type WorkerServiceEnvironmentSettingUpdateParamsResultMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses param.Field[[]WorkerServiceEnvironmentSettingUpdateParamsResultMigrationsWorkersSingleStepMigrationsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]WorkerServiceEnvironmentSettingUpdateParamsResultMigrationsWorkersSingleStepMigrationsTransferredClass] `json:"transferred_classes"`
}

func (r WorkerServiceEnvironmentSettingUpdateParamsResultMigrationsWorkersSingleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerServiceEnvironmentSettingUpdateParamsResultMigrationsWorkersSingleStepMigrations) implementsWorkerServiceEnvironmentSettingUpdateParamsResultMigrations() {
}

type WorkerServiceEnvironmentSettingUpdateParamsResultMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r WorkerServiceEnvironmentSettingUpdateParamsResultMigrationsWorkersSingleStepMigrationsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerServiceEnvironmentSettingUpdateParamsResultMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r WorkerServiceEnvironmentSettingUpdateParamsResultMigrationsWorkersSingleStepMigrationsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerServiceEnvironmentSettingUpdateParamsResultMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]WorkerServiceEnvironmentSettingUpdateParamsResultMigrationsWorkersSteppedMigrationsStep] `json:"steps"`
}

func (r WorkerServiceEnvironmentSettingUpdateParamsResultMigrationsWorkersSteppedMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerServiceEnvironmentSettingUpdateParamsResultMigrationsWorkersSteppedMigrations) implementsWorkerServiceEnvironmentSettingUpdateParamsResultMigrations() {
}

type WorkerServiceEnvironmentSettingUpdateParamsResultMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]WorkerServiceEnvironmentSettingUpdateParamsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]WorkerServiceEnvironmentSettingUpdateParamsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass] `json:"transferred_classes"`
}

func (r WorkerServiceEnvironmentSettingUpdateParamsResultMigrationsWorkersSteppedMigrationsStep) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerServiceEnvironmentSettingUpdateParamsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r WorkerServiceEnvironmentSettingUpdateParamsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerServiceEnvironmentSettingUpdateParamsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r WorkerServiceEnvironmentSettingUpdateParamsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerServiceEnvironmentSettingUpdateParamsResultPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode param.Field[WorkerServiceEnvironmentSettingUpdateParamsResultPlacementMode] `json:"mode"`
}

func (r WorkerServiceEnvironmentSettingUpdateParamsResultPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type WorkerServiceEnvironmentSettingUpdateParamsResultPlacementMode string

const (
	WorkerServiceEnvironmentSettingUpdateParamsResultPlacementModeSmart WorkerServiceEnvironmentSettingUpdateParamsResultPlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type WorkerServiceEnvironmentSettingUpdateParamsResultTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service param.Field[string] `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace param.Field[string] `json:"namespace"`
}

func (r WorkerServiceEnvironmentSettingUpdateParamsResultTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the API call was successful
type WorkerServiceEnvironmentSettingUpdateParamsSuccess bool

const (
	WorkerServiceEnvironmentSettingUpdateParamsSuccessTrue WorkerServiceEnvironmentSettingUpdateParamsSuccess = true
)

type WorkerServiceEnvironmentSettingUpdateResponseEnvelope struct {
	Errors   []WorkerServiceEnvironmentSettingUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerServiceEnvironmentSettingUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerServiceEnvironmentSettingUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerServiceEnvironmentSettingUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerServiceEnvironmentSettingUpdateResponseEnvelopeJSON    `json:"-"`
}

// workerServiceEnvironmentSettingUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [WorkerServiceEnvironmentSettingUpdateResponseEnvelope]
type workerServiceEnvironmentSettingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerServiceEnvironmentSettingUpdateResponseEnvelopeErrors struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    workerServiceEnvironmentSettingUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// workerServiceEnvironmentSettingUpdateResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [WorkerServiceEnvironmentSettingUpdateResponseEnvelopeErrors]
type workerServiceEnvironmentSettingUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerServiceEnvironmentSettingUpdateResponseEnvelopeMessages struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    workerServiceEnvironmentSettingUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// workerServiceEnvironmentSettingUpdateResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [WorkerServiceEnvironmentSettingUpdateResponseEnvelopeMessages]
type workerServiceEnvironmentSettingUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerServiceEnvironmentSettingUpdateResponseEnvelopeSuccess bool

const (
	WorkerServiceEnvironmentSettingUpdateResponseEnvelopeSuccessTrue WorkerServiceEnvironmentSettingUpdateResponseEnvelopeSuccess = true
)

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
