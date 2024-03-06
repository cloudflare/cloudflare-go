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
	"github.com/tidwall/gjson"
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
func (r *WorkerServiceEnvironmentSettingService) Edit(ctx context.Context, serviceName string, environmentName string, params WorkerServiceEnvironmentSettingEditParams, opts ...option.RequestOption) (res *WorkerServiceEnvironmentSettingEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerServiceEnvironmentSettingEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/services/%s/environments/%s/settings", params.AccountID, serviceName, environmentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get script settings from a worker with an environment
func (r *WorkerServiceEnvironmentSettingService) Get(ctx context.Context, serviceName string, environmentName string, query WorkerServiceEnvironmentSettingGetParams, opts ...option.RequestOption) (res *WorkerServiceEnvironmentSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerServiceEnvironmentSettingGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/services/%s/environments/%s/settings", query.AccountID, serviceName, environmentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkerServiceEnvironmentSettingEditResponse struct {
	// List of bindings attached to this Worker
	Bindings []WorkerServiceEnvironmentSettingEditResponseBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations WorkerServiceEnvironmentSettingEditResponseMigrations `json:"migrations"`
	Placement  WorkerServiceEnvironmentSettingEditResponsePlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []WorkerServiceEnvironmentSettingEditResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                          `json:"usage_model"`
	JSON       workerServiceEnvironmentSettingEditResponseJSON `json:"-"`
}

// workerServiceEnvironmentSettingEditResponseJSON contains the JSON metadata for
// the struct [WorkerServiceEnvironmentSettingEditResponse]
type workerServiceEnvironmentSettingEditResponseJSON struct {
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

func (r *WorkerServiceEnvironmentSettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerServiceEnvironmentSettingEditResponseJSON) RawJSON() string {
	return r.raw
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [WorkerServiceEnvironmentSettingEditResponseBindingsWorkersKVNamespaceBinding],
// [WorkerServiceEnvironmentSettingEditResponseBindingsWorkersServiceBinding],
// [WorkerServiceEnvironmentSettingEditResponseBindingsWorkersDoBinding],
// [WorkerServiceEnvironmentSettingEditResponseBindingsWorkersR2Binding],
// [WorkerServiceEnvironmentSettingEditResponseBindingsWorkersQueueBinding],
// [WorkerServiceEnvironmentSettingEditResponseBindingsWorkersD1Binding],
// [WorkerServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBinding]
// or [WorkerServiceEnvironmentSettingEditResponseBindingsWorkersMTLSCertBinding].
type WorkerServiceEnvironmentSettingEditResponseBinding interface {
	implementsWorkerServiceEnvironmentSettingEditResponseBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WorkerServiceEnvironmentSettingEditResponseBinding)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WorkerServiceEnvironmentSettingEditResponseBindingsWorkersKVNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WorkerServiceEnvironmentSettingEditResponseBindingsWorkersServiceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WorkerServiceEnvironmentSettingEditResponseBindingsWorkersDoBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WorkerServiceEnvironmentSettingEditResponseBindingsWorkersR2Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WorkerServiceEnvironmentSettingEditResponseBindingsWorkersQueueBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WorkerServiceEnvironmentSettingEditResponseBindingsWorkersD1Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WorkerServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WorkerServiceEnvironmentSettingEditResponseBindingsWorkersMTLSCertBinding{}),
		},
	)
}

type WorkerServiceEnvironmentSettingEditResponseBindingsWorkersKVNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingEditResponseBindingsWorkersKVNamespaceBindingType `json:"type,required"`
	JSON workerServiceEnvironmentSettingEditResponseBindingsWorkersKVNamespaceBindingJSON `json:"-"`
}

// workerServiceEnvironmentSettingEditResponseBindingsWorkersKVNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingEditResponseBindingsWorkersKVNamespaceBinding]
type workerServiceEnvironmentSettingEditResponseBindingsWorkersKVNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingEditResponseBindingsWorkersKVNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerServiceEnvironmentSettingEditResponseBindingsWorkersKVNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkerServiceEnvironmentSettingEditResponseBindingsWorkersKVNamespaceBinding) implementsWorkerServiceEnvironmentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingEditResponseBindingsWorkersKVNamespaceBindingType string

const (
	WorkerServiceEnvironmentSettingEditResponseBindingsWorkersKVNamespaceBindingTypeKVNamespace WorkerServiceEnvironmentSettingEditResponseBindingsWorkersKVNamespaceBindingType = "kv_namespace"
)

type WorkerServiceEnvironmentSettingEditResponseBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingEditResponseBindingsWorkersServiceBindingType `json:"type,required"`
	JSON workerServiceEnvironmentSettingEditResponseBindingsWorkersServiceBindingJSON `json:"-"`
}

// workerServiceEnvironmentSettingEditResponseBindingsWorkersServiceBindingJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingEditResponseBindingsWorkersServiceBinding]
type workerServiceEnvironmentSettingEditResponseBindingsWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingEditResponseBindingsWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerServiceEnvironmentSettingEditResponseBindingsWorkersServiceBindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkerServiceEnvironmentSettingEditResponseBindingsWorkersServiceBinding) implementsWorkerServiceEnvironmentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingEditResponseBindingsWorkersServiceBindingType string

const (
	WorkerServiceEnvironmentSettingEditResponseBindingsWorkersServiceBindingTypeService WorkerServiceEnvironmentSettingEditResponseBindingsWorkersServiceBindingType = "service"
)

type WorkerServiceEnvironmentSettingEditResponseBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingEditResponseBindingsWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                                  `json:"script_name"`
	JSON       workerServiceEnvironmentSettingEditResponseBindingsWorkersDoBindingJSON `json:"-"`
}

// workerServiceEnvironmentSettingEditResponseBindingsWorkersDoBindingJSON contains
// the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingEditResponseBindingsWorkersDoBinding]
type workerServiceEnvironmentSettingEditResponseBindingsWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingEditResponseBindingsWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerServiceEnvironmentSettingEditResponseBindingsWorkersDoBindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkerServiceEnvironmentSettingEditResponseBindingsWorkersDoBinding) implementsWorkerServiceEnvironmentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingEditResponseBindingsWorkersDoBindingType string

const (
	WorkerServiceEnvironmentSettingEditResponseBindingsWorkersDoBindingTypeDurableObjectNamespace WorkerServiceEnvironmentSettingEditResponseBindingsWorkersDoBindingType = "durable_object_namespace"
)

type WorkerServiceEnvironmentSettingEditResponseBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingEditResponseBindingsWorkersR2BindingType `json:"type,required"`
	JSON workerServiceEnvironmentSettingEditResponseBindingsWorkersR2BindingJSON `json:"-"`
}

// workerServiceEnvironmentSettingEditResponseBindingsWorkersR2BindingJSON contains
// the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingEditResponseBindingsWorkersR2Binding]
type workerServiceEnvironmentSettingEditResponseBindingsWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingEditResponseBindingsWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerServiceEnvironmentSettingEditResponseBindingsWorkersR2BindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkerServiceEnvironmentSettingEditResponseBindingsWorkersR2Binding) implementsWorkerServiceEnvironmentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingEditResponseBindingsWorkersR2BindingType string

const (
	WorkerServiceEnvironmentSettingEditResponseBindingsWorkersR2BindingTypeR2Bucket WorkerServiceEnvironmentSettingEditResponseBindingsWorkersR2BindingType = "r2_bucket"
)

type WorkerServiceEnvironmentSettingEditResponseBindingsWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingEditResponseBindingsWorkersQueueBindingType `json:"type,required"`
	JSON workerServiceEnvironmentSettingEditResponseBindingsWorkersQueueBindingJSON `json:"-"`
}

// workerServiceEnvironmentSettingEditResponseBindingsWorkersQueueBindingJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingEditResponseBindingsWorkersQueueBinding]
type workerServiceEnvironmentSettingEditResponseBindingsWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingEditResponseBindingsWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerServiceEnvironmentSettingEditResponseBindingsWorkersQueueBindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkerServiceEnvironmentSettingEditResponseBindingsWorkersQueueBinding) implementsWorkerServiceEnvironmentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingEditResponseBindingsWorkersQueueBindingType string

const (
	WorkerServiceEnvironmentSettingEditResponseBindingsWorkersQueueBindingTypeQueue WorkerServiceEnvironmentSettingEditResponseBindingsWorkersQueueBindingType = "queue"
)

type WorkerServiceEnvironmentSettingEditResponseBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingEditResponseBindingsWorkersD1BindingType `json:"type,required"`
	JSON workerServiceEnvironmentSettingEditResponseBindingsWorkersD1BindingJSON `json:"-"`
}

// workerServiceEnvironmentSettingEditResponseBindingsWorkersD1BindingJSON contains
// the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingEditResponseBindingsWorkersD1Binding]
type workerServiceEnvironmentSettingEditResponseBindingsWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingEditResponseBindingsWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerServiceEnvironmentSettingEditResponseBindingsWorkersD1BindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkerServiceEnvironmentSettingEditResponseBindingsWorkersD1Binding) implementsWorkerServiceEnvironmentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingEditResponseBindingsWorkersD1BindingType string

const (
	WorkerServiceEnvironmentSettingEditResponseBindingsWorkersD1BindingTypeD1 WorkerServiceEnvironmentSettingEditResponseBindingsWorkersD1BindingType = "d1"
)

type WorkerServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound WorkerServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     workerServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON     `json:"-"`
}

// workerServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBinding]
type workerServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkerServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBinding) implementsWorkerServiceEnvironmentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingType string

const (
	WorkerServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace WorkerServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type WorkerServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker WorkerServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   workerServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// workerServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound]
type workerServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON) RawJSON() string {
	return r.raw
}

// Outbound worker
type WorkerServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                               `json:"service"`
	JSON    workerServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// workerServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker]
type workerServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type WorkerServiceEnvironmentSettingEditResponseBindingsWorkersMTLSCertBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingEditResponseBindingsWorkersMTLSCertBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                                        `json:"certificate_id"`
	JSON          workerServiceEnvironmentSettingEditResponseBindingsWorkersMTLSCertBindingJSON `json:"-"`
}

// workerServiceEnvironmentSettingEditResponseBindingsWorkersMTLSCertBindingJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingEditResponseBindingsWorkersMTLSCertBinding]
type workerServiceEnvironmentSettingEditResponseBindingsWorkersMTLSCertBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingEditResponseBindingsWorkersMTLSCertBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerServiceEnvironmentSettingEditResponseBindingsWorkersMTLSCertBindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkerServiceEnvironmentSettingEditResponseBindingsWorkersMTLSCertBinding) implementsWorkerServiceEnvironmentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingEditResponseBindingsWorkersMTLSCertBindingType string

const (
	WorkerServiceEnvironmentSettingEditResponseBindingsWorkersMTLSCertBindingTypeMTLSCertificate WorkerServiceEnvironmentSettingEditResponseBindingsWorkersMTLSCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrations]
// or
// [WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrations].
type WorkerServiceEnvironmentSettingEditResponseMigrations interface {
	implementsWorkerServiceEnvironmentSettingEditResponseMigrations()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WorkerServiceEnvironmentSettingEditResponseMigrations)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrations{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrations{}),
		},
	)
}

// A single set of migrations to apply.
type WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses []WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               workerServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON               `json:"-"`
}

// workerServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrations]
type workerServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrations) implementsWorkerServiceEnvironmentSettingEditResponseMigrations() {
}

type WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From string                                                                                           `json:"from"`
	To   string                                                                                           `json:"to"`
	JSON workerServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// workerServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass]
type workerServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON) RawJSON() string {
	return r.raw
}

type WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       string                                                                                               `json:"from"`
	FromScript string                                                                                               `json:"from_script"`
	To         string                                                                                               `json:"to"`
	JSON       workerServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// workerServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass]
type workerServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON) RawJSON() string {
	return r.raw
}

type WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStep `json:"steps"`
	JSON  workerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsJSON   `json:"-"`
}

// workerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrations]
type workerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrations) implementsWorkerServiceEnvironmentSettingEditResponseMigrations() {
}

type WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               workerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON                `json:"-"`
}

// workerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStep]
type workerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON) RawJSON() string {
	return r.raw
}

type WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From string                                                                                             `json:"from"`
	To   string                                                                                             `json:"to"`
	JSON workerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// workerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass]
type workerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON) RawJSON() string {
	return r.raw
}

type WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                                                 `json:"from"`
	FromScript string                                                                                                 `json:"from_script"`
	To         string                                                                                                 `json:"to"`
	JSON       workerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// workerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass]
type workerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON) RawJSON() string {
	return r.raw
}

type WorkerServiceEnvironmentSettingEditResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode WorkerServiceEnvironmentSettingEditResponsePlacementMode `json:"mode"`
	JSON workerServiceEnvironmentSettingEditResponsePlacementJSON `json:"-"`
}

// workerServiceEnvironmentSettingEditResponsePlacementJSON contains the JSON
// metadata for the struct [WorkerServiceEnvironmentSettingEditResponsePlacement]
type workerServiceEnvironmentSettingEditResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingEditResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerServiceEnvironmentSettingEditResponsePlacementJSON) RawJSON() string {
	return r.raw
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type WorkerServiceEnvironmentSettingEditResponsePlacementMode string

const (
	WorkerServiceEnvironmentSettingEditResponsePlacementModeSmart WorkerServiceEnvironmentSettingEditResponsePlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type WorkerServiceEnvironmentSettingEditResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                      `json:"namespace"`
	JSON      workerServiceEnvironmentSettingEditResponseTailConsumerJSON `json:"-"`
}

// workerServiceEnvironmentSettingEditResponseTailConsumerJSON contains the JSON
// metadata for the struct
// [WorkerServiceEnvironmentSettingEditResponseTailConsumer]
type workerServiceEnvironmentSettingEditResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingEditResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerServiceEnvironmentSettingEditResponseTailConsumerJSON) RawJSON() string {
	return r.raw
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

func (r workerServiceEnvironmentSettingGetResponseJSON) RawJSON() string {
	return r.raw
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [WorkerServiceEnvironmentSettingGetResponseBindingsWorkersKVNamespaceBinding],
// [WorkerServiceEnvironmentSettingGetResponseBindingsWorkersServiceBinding],
// [WorkerServiceEnvironmentSettingGetResponseBindingsWorkersDoBinding],
// [WorkerServiceEnvironmentSettingGetResponseBindingsWorkersR2Binding],
// [WorkerServiceEnvironmentSettingGetResponseBindingsWorkersQueueBinding],
// [WorkerServiceEnvironmentSettingGetResponseBindingsWorkersD1Binding],
// [WorkerServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBinding]
// or [WorkerServiceEnvironmentSettingGetResponseBindingsWorkersMTLSCertBinding].
type WorkerServiceEnvironmentSettingGetResponseBinding interface {
	implementsWorkerServiceEnvironmentSettingGetResponseBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WorkerServiceEnvironmentSettingGetResponseBinding)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WorkerServiceEnvironmentSettingGetResponseBindingsWorkersKVNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WorkerServiceEnvironmentSettingGetResponseBindingsWorkersServiceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WorkerServiceEnvironmentSettingGetResponseBindingsWorkersDoBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WorkerServiceEnvironmentSettingGetResponseBindingsWorkersR2Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WorkerServiceEnvironmentSettingGetResponseBindingsWorkersQueueBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WorkerServiceEnvironmentSettingGetResponseBindingsWorkersD1Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WorkerServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WorkerServiceEnvironmentSettingGetResponseBindingsWorkersMTLSCertBinding{}),
		},
	)
}

type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersKVNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersKVNamespaceBindingType `json:"type,required"`
	JSON workerServiceEnvironmentSettingGetResponseBindingsWorkersKVNamespaceBindingJSON `json:"-"`
}

// workerServiceEnvironmentSettingGetResponseBindingsWorkersKVNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingGetResponseBindingsWorkersKVNamespaceBinding]
type workerServiceEnvironmentSettingGetResponseBindingsWorkersKVNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingGetResponseBindingsWorkersKVNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerServiceEnvironmentSettingGetResponseBindingsWorkersKVNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkerServiceEnvironmentSettingGetResponseBindingsWorkersKVNamespaceBinding) implementsWorkerServiceEnvironmentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersKVNamespaceBindingType string

const (
	WorkerServiceEnvironmentSettingGetResponseBindingsWorkersKVNamespaceBindingTypeKVNamespace WorkerServiceEnvironmentSettingGetResponseBindingsWorkersKVNamespaceBindingType = "kv_namespace"
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

func (r workerServiceEnvironmentSettingGetResponseBindingsWorkersServiceBindingJSON) RawJSON() string {
	return r.raw
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

func (r workerServiceEnvironmentSettingGetResponseBindingsWorkersDoBindingJSON) RawJSON() string {
	return r.raw
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

func (r workerServiceEnvironmentSettingGetResponseBindingsWorkersR2BindingJSON) RawJSON() string {
	return r.raw
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

func (r workerServiceEnvironmentSettingGetResponseBindingsWorkersQueueBindingJSON) RawJSON() string {
	return r.raw
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

func (r workerServiceEnvironmentSettingGetResponseBindingsWorkersD1BindingJSON) RawJSON() string {
	return r.raw
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

func (r workerServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON) RawJSON() string {
	return r.raw
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

func (r workerServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON) RawJSON() string {
	return r.raw
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

func (r workerServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersMTLSCertBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersMTLSCertBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                                       `json:"certificate_id"`
	JSON          workerServiceEnvironmentSettingGetResponseBindingsWorkersMTLSCertBindingJSON `json:"-"`
}

// workerServiceEnvironmentSettingGetResponseBindingsWorkersMTLSCertBindingJSON
// contains the JSON metadata for the struct
// [WorkerServiceEnvironmentSettingGetResponseBindingsWorkersMTLSCertBinding]
type workerServiceEnvironmentSettingGetResponseBindingsWorkersMTLSCertBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingGetResponseBindingsWorkersMTLSCertBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerServiceEnvironmentSettingGetResponseBindingsWorkersMTLSCertBindingJSON) RawJSON() string {
	return r.raw
}

func (r WorkerServiceEnvironmentSettingGetResponseBindingsWorkersMTLSCertBinding) implementsWorkerServiceEnvironmentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingGetResponseBindingsWorkersMTLSCertBindingType string

const (
	WorkerServiceEnvironmentSettingGetResponseBindingsWorkersMTLSCertBindingTypeMTLSCertificate WorkerServiceEnvironmentSettingGetResponseBindingsWorkersMTLSCertBindingType = "mtls_certificate"
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
	apijson.RegisterUnion(
		reflect.TypeOf((*WorkerServiceEnvironmentSettingGetResponseMigrations)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrations{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WorkerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrations{}),
		},
	)
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

func (r workerServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON) RawJSON() string {
	return r.raw
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

func (r workerServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON) RawJSON() string {
	return r.raw
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

func (r workerServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON) RawJSON() string {
	return r.raw
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

func (r workerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsJSON) RawJSON() string {
	return r.raw
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

func (r workerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON) RawJSON() string {
	return r.raw
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

func (r workerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON) RawJSON() string {
	return r.raw
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

func (r workerServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON) RawJSON() string {
	return r.raw
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

func (r workerServiceEnvironmentSettingGetResponsePlacementJSON) RawJSON() string {
	return r.raw
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

func (r workerServiceEnvironmentSettingGetResponseTailConsumerJSON) RawJSON() string {
	return r.raw
}

type WorkerServiceEnvironmentSettingEditParams struct {
	// Identifier
	AccountID param.Field[string]                                             `path:"account_id,required"`
	Errors    param.Field[[]WorkerServiceEnvironmentSettingEditParamsError]   `json:"errors,required"`
	Messages  param.Field[[]WorkerServiceEnvironmentSettingEditParamsMessage] `json:"messages,required"`
	Result    param.Field[WorkerServiceEnvironmentSettingEditParamsResult]    `json:"result,required"`
	// Whether the API call was successful
	Success param.Field[WorkerServiceEnvironmentSettingEditParamsSuccess] `json:"success,required"`
}

func (r WorkerServiceEnvironmentSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerServiceEnvironmentSettingEditParamsError struct {
	Code    param.Field[int64]  `json:"code,required"`
	Message param.Field[string] `json:"message,required"`
}

func (r WorkerServiceEnvironmentSettingEditParamsError) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerServiceEnvironmentSettingEditParamsMessage struct {
	Code    param.Field[int64]  `json:"code,required"`
	Message param.Field[string] `json:"message,required"`
}

func (r WorkerServiceEnvironmentSettingEditParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerServiceEnvironmentSettingEditParamsResult struct {
	// List of bindings attached to this Worker
	Bindings param.Field[[]WorkerServiceEnvironmentSettingEditParamsResultBinding] `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations param.Field[WorkerServiceEnvironmentSettingEditParamsResultMigrations] `json:"migrations"`
	Placement  param.Field[WorkerServiceEnvironmentSettingEditParamsResultPlacement]  `json:"placement"`
	// Tags to help you manage your Workers
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]WorkerServiceEnvironmentSettingEditParamsResultTailConsumer] `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel param.Field[string] `json:"usage_model"`
}

func (r WorkerServiceEnvironmentSettingEditParamsResult) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A binding to allow the Worker to communicate with resources
//
// Satisfied by
// [WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersKVNamespaceBinding],
// [WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersServiceBinding],
// [WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersDoBinding],
// [WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersR2Binding],
// [WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersQueueBinding],
// [WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersD1Binding],
// [WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersDispatchNamespaceBinding],
// [WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersMTLSCertBinding].
type WorkerServiceEnvironmentSettingEditParamsResultBinding interface {
	implementsWorkerServiceEnvironmentSettingEditParamsResultBinding()
}

type WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersKVNamespaceBinding struct {
	// The class of resource that the binding provides.
	Type param.Field[WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersKVNamespaceBindingType] `json:"type,required"`
}

func (r WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersKVNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersKVNamespaceBinding) implementsWorkerServiceEnvironmentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersKVNamespaceBindingType string

const (
	WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersKVNamespaceBindingTypeKVNamespace WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersKVNamespaceBindingType = "kv_namespace"
)

type WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// Name of Worker to bind to
	Service param.Field[string] `json:"service,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersServiceBindingType] `json:"type,required"`
}

func (r WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersServiceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersServiceBinding) implementsWorkerServiceEnvironmentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersServiceBindingType string

const (
	WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersServiceBindingTypeService WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersServiceBindingType = "service"
)

type WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName param.Field[string] `json:"class_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersDoBindingType] `json:"type,required"`
	// The environment of the script_name to bind to
	Environment param.Field[string] `json:"environment"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName param.Field[string] `json:"script_name"`
}

func (r WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersDoBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersDoBinding) implementsWorkerServiceEnvironmentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersDoBindingType string

const (
	WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersDoBindingTypeDurableObjectNamespace WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersDoBindingType = "durable_object_namespace"
)

type WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName param.Field[string] `json:"bucket_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersR2BindingType] `json:"type,required"`
}

func (r WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersR2Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersR2Binding) implementsWorkerServiceEnvironmentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersR2BindingType string

const (
	WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersR2BindingTypeR2Bucket WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersR2BindingType = "r2_bucket"
)

type WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersQueueBinding struct {
	// Name of the Queue to bind to
	QueueName param.Field[string] `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersQueueBindingType] `json:"type,required"`
}

func (r WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersQueueBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersQueueBinding) implementsWorkerServiceEnvironmentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersQueueBindingType string

const (
	WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersQueueBindingTypeQueue WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersQueueBindingType = "queue"
)

type WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID param.Field[string] `json:"id,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name param.Field[string] `json:"name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersD1BindingType] `json:"type,required"`
}

func (r WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersD1Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersD1Binding) implementsWorkerServiceEnvironmentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersD1BindingType string

const (
	WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersD1BindingTypeD1 WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersD1BindingType = "d1"
)

type WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersDispatchNamespaceBinding struct {
	// Namespace to bind to
	Namespace param.Field[string] `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingType] `json:"type,required"`
	// Outbound worker
	Outbound param.Field[WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutbound] `json:"outbound"`
}

func (r WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersDispatchNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersDispatchNamespaceBinding) implementsWorkerServiceEnvironmentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingType string

const (
	WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params param.Field[[]string] `json:"params"`
	// Outbound worker
	Worker param.Field[WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker] `json:"worker"`
}

func (r WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutbound) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker
type WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker
	Service param.Field[string] `json:"service"`
}

func (r WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersMTLSCertBinding struct {
	// The class of resource that the binding provides.
	Type param.Field[WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersMTLSCertBindingType] `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID param.Field[string] `json:"certificate_id"`
}

func (r WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersMTLSCertBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersMTLSCertBinding) implementsWorkerServiceEnvironmentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersMTLSCertBindingType string

const (
	WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersMTLSCertBindingTypeMTLSCertificate WorkerServiceEnvironmentSettingEditParamsResultBindingsWorkersMTLSCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by
// [WorkerServiceEnvironmentSettingEditParamsResultMigrationsWorkersSingleStepMigrations],
// [WorkerServiceEnvironmentSettingEditParamsResultMigrationsWorkersSteppedMigrations].
type WorkerServiceEnvironmentSettingEditParamsResultMigrations interface {
	implementsWorkerServiceEnvironmentSettingEditParamsResultMigrations()
}

// A single set of migrations to apply.
type WorkerServiceEnvironmentSettingEditParamsResultMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses param.Field[[]WorkerServiceEnvironmentSettingEditParamsResultMigrationsWorkersSingleStepMigrationsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]WorkerServiceEnvironmentSettingEditParamsResultMigrationsWorkersSingleStepMigrationsTransferredClass] `json:"transferred_classes"`
}

func (r WorkerServiceEnvironmentSettingEditParamsResultMigrationsWorkersSingleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerServiceEnvironmentSettingEditParamsResultMigrationsWorkersSingleStepMigrations) implementsWorkerServiceEnvironmentSettingEditParamsResultMigrations() {
}

type WorkerServiceEnvironmentSettingEditParamsResultMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r WorkerServiceEnvironmentSettingEditParamsResultMigrationsWorkersSingleStepMigrationsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerServiceEnvironmentSettingEditParamsResultMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r WorkerServiceEnvironmentSettingEditParamsResultMigrationsWorkersSingleStepMigrationsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerServiceEnvironmentSettingEditParamsResultMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]WorkerServiceEnvironmentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStep] `json:"steps"`
}

func (r WorkerServiceEnvironmentSettingEditParamsResultMigrationsWorkersSteppedMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkerServiceEnvironmentSettingEditParamsResultMigrationsWorkersSteppedMigrations) implementsWorkerServiceEnvironmentSettingEditParamsResultMigrations() {
}

type WorkerServiceEnvironmentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]WorkerServiceEnvironmentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]WorkerServiceEnvironmentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass] `json:"transferred_classes"`
}

func (r WorkerServiceEnvironmentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStep) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerServiceEnvironmentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r WorkerServiceEnvironmentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerServiceEnvironmentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r WorkerServiceEnvironmentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerServiceEnvironmentSettingEditParamsResultPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode param.Field[WorkerServiceEnvironmentSettingEditParamsResultPlacementMode] `json:"mode"`
}

func (r WorkerServiceEnvironmentSettingEditParamsResultPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type WorkerServiceEnvironmentSettingEditParamsResultPlacementMode string

const (
	WorkerServiceEnvironmentSettingEditParamsResultPlacementModeSmart WorkerServiceEnvironmentSettingEditParamsResultPlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type WorkerServiceEnvironmentSettingEditParamsResultTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service param.Field[string] `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace param.Field[string] `json:"namespace"`
}

func (r WorkerServiceEnvironmentSettingEditParamsResultTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the API call was successful
type WorkerServiceEnvironmentSettingEditParamsSuccess bool

const (
	WorkerServiceEnvironmentSettingEditParamsSuccessTrue WorkerServiceEnvironmentSettingEditParamsSuccess = true
)

type WorkerServiceEnvironmentSettingEditResponseEnvelope struct {
	Errors   []WorkerServiceEnvironmentSettingEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerServiceEnvironmentSettingEditResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerServiceEnvironmentSettingEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerServiceEnvironmentSettingEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerServiceEnvironmentSettingEditResponseEnvelopeJSON    `json:"-"`
}

// workerServiceEnvironmentSettingEditResponseEnvelopeJSON contains the JSON
// metadata for the struct [WorkerServiceEnvironmentSettingEditResponseEnvelope]
type workerServiceEnvironmentSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerServiceEnvironmentSettingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WorkerServiceEnvironmentSettingEditResponseEnvelopeErrors struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    workerServiceEnvironmentSettingEditResponseEnvelopeErrorsJSON `json:"-"`
}

// workerServiceEnvironmentSettingEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [WorkerServiceEnvironmentSettingEditResponseEnvelopeErrors]
type workerServiceEnvironmentSettingEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerServiceEnvironmentSettingEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WorkerServiceEnvironmentSettingEditResponseEnvelopeMessages struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    workerServiceEnvironmentSettingEditResponseEnvelopeMessagesJSON `json:"-"`
}

// workerServiceEnvironmentSettingEditResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [WorkerServiceEnvironmentSettingEditResponseEnvelopeMessages]
type workerServiceEnvironmentSettingEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerServiceEnvironmentSettingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workerServiceEnvironmentSettingEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WorkerServiceEnvironmentSettingEditResponseEnvelopeSuccess bool

const (
	WorkerServiceEnvironmentSettingEditResponseEnvelopeSuccessTrue WorkerServiceEnvironmentSettingEditResponseEnvelopeSuccess = true
)

type WorkerServiceEnvironmentSettingGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
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

func (r workerServiceEnvironmentSettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r workerServiceEnvironmentSettingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r workerServiceEnvironmentSettingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WorkerServiceEnvironmentSettingGetResponseEnvelopeSuccess bool

const (
	WorkerServiceEnvironmentSettingGetResponseEnvelopeSuccessTrue WorkerServiceEnvironmentSettingGetResponseEnvelopeSuccess = true
)
