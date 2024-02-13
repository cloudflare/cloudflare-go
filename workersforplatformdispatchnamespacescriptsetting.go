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

// WorkersForPlatformDispatchNamespaceScriptSettingService contains methods and
// other services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkersForPlatformDispatchNamespaceScriptSettingService] method instead.
type WorkersForPlatformDispatchNamespaceScriptSettingService struct {
	Options []option.RequestOption
}

// NewWorkersForPlatformDispatchNamespaceScriptSettingService generates a new
// service that applies the given options to each request. These options are
// applied after the parent client's options (if there is one), and before any
// request-specific options.
func NewWorkersForPlatformDispatchNamespaceScriptSettingService(opts ...option.RequestOption) (r *WorkersForPlatformDispatchNamespaceScriptSettingService) {
	r = &WorkersForPlatformDispatchNamespaceScriptSettingService{}
	r.Options = opts
	return
}

// Patch script metadata, such as bindings
func (r *WorkersForPlatformDispatchNamespaceScriptSettingService) Update(ctx context.Context, accountID string, dispatchNamespace string, scriptName string, body WorkersForPlatformDispatchNamespaceScriptSettingUpdateParams, opts ...option.RequestOption) (res *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/settings", accountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get script settings from a script uploaded to a Workers for Platforms namespace.
func (r *WorkersForPlatformDispatchNamespaceScriptSettingService) Get(ctx context.Context, accountID string, dispatchNamespace string, scriptName string, opts ...option.RequestOption) (res *WorkersForPlatformDispatchNamespaceScriptSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkersForPlatformDispatchNamespaceScriptSettingGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/settings", accountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponse struct {
	// List of bindings attached to this Worker
	Bindings []WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrations `json:"migrations"`
	Placement  WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponsePlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                                             `json:"usage_model"`
	JSON       workersForPlatformDispatchNamespaceScriptSettingUpdateResponseJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseJSON contains the
// JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponse]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseJSON struct {
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

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersKvNamespaceBinding],
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersServiceBinding],
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDoBinding],
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersR2Binding],
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersQueueBinding],
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersD1Binding],
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBinding]
// or
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersMtlsCertBinding].
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBinding interface {
	implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBinding()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBinding)(nil)).Elem(), "")
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersKvNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersKvNamespaceBindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersKvNamespaceBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersKvNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersKvNamespaceBinding]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersKvNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersKvNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersKvNamespaceBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersKvNamespaceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersKvNamespaceBindingTypeKvNamespace WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersKvNamespaceBindingType = "kv_namespace"
)

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersServiceBindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersServiceBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersServiceBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersServiceBinding]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersServiceBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersServiceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersServiceBindingTypeService WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersServiceBindingType = "service"
)

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                                                     `json:"script_name"`
	JSON       workersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDoBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDoBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDoBinding]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDoBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDoBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDoBindingTypeDurableObjectNamespace WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDoBindingType = "durable_object_namespace"
)

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersR2BindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersR2BindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersR2BindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersR2Binding]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersR2Binding) implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersR2BindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersR2BindingTypeR2Bucket WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersR2BindingType = "r2_bucket"
)

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersQueueBindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersQueueBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersQueueBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersQueueBinding]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersQueueBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersQueueBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersQueueBindingTypeQueue WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersQueueBindingType = "queue"
)

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersD1BindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersD1BindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersD1BindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersD1Binding]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersD1Binding) implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersD1BindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersD1BindingTypeD1 WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersD1BindingType = "d1"
)

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     workersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingJSON     `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBinding]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   workersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutbound]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Outbound worker
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                                                  `json:"service"`
	JSON    workersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersMtlsCertBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersMtlsCertBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                                                           `json:"certificate_id"`
	JSON          workersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersMtlsCertBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersMtlsCertBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersMtlsCertBinding]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersMtlsCertBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersMtlsCertBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersMtlsCertBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersMtlsCertBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersMtlsCertBindingTypeMtlsCertificate WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseBindingsWorkersMtlsCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrations]
// or
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSteppedMigrations].
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrations interface {
	implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrations)(nil)).Elem(), "")
}

// A single set of migrations to apply.
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses []WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               workersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsJSON               `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrations]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrations) implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrations() {
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From string                                                                                                              `json:"from"`
	To   string                                                                                                              `json:"to"`
	JSON workersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsRenamedClass]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       string                                                                                                                  `json:"from"`
	FromScript string                                                                                                                  `json:"from_script"`
	To         string                                                                                                                  `json:"to"`
	JSON       workersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsTransferredClass]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStep `json:"steps"`
	JSON  workersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsJSON   `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSteppedMigrations]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSteppedMigrations) implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrations() {
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               workersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepJSON                `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStep]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From string                                                                                                                `json:"from"`
	To   string                                                                                                                `json:"to"`
	JSON workersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                                                                    `json:"from"`
	FromScript string                                                                                                                    `json:"from_script"`
	To         string                                                                                                                    `json:"to"`
	JSON       workersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponsePlacementMode `json:"mode"`
	JSON workersForPlatformDispatchNamespaceScriptSettingUpdateResponsePlacementJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponsePlacementJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponsePlacement]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponsePlacementMode string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponsePlacementModeSmart WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponsePlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                                         `json:"namespace"`
	JSON      workersForPlatformDispatchNamespaceScriptSettingUpdateResponseTailConsumerJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseTailConsumerJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseTailConsumer]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponse struct {
	// List of bindings attached to this Worker
	Bindings []WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrations `json:"migrations"`
	Placement  WorkersForPlatformDispatchNamespaceScriptSettingGetResponsePlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []WorkersForPlatformDispatchNamespaceScriptSettingGetResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                                          `json:"usage_model"`
	JSON       workersForPlatformDispatchNamespaceScriptSettingGetResponseJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseJSON contains the
// JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponse]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseJSON struct {
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

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersKvNamespaceBinding],
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersServiceBinding],
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDoBinding],
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersR2Binding],
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersQueueBinding],
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersD1Binding],
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBinding]
// or
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersMtlsCertBinding].
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBinding interface {
	implementsWorkersForPlatformDispatchNamespaceScriptSettingGetResponseBinding()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBinding)(nil)).Elem(), "")
}

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersKvNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersKvNamespaceBindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersKvNamespaceBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersKvNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersKvNamespaceBinding]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersKvNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersKvNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersKvNamespaceBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersKvNamespaceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersKvNamespaceBindingTypeKvNamespace WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersKvNamespaceBindingType = "kv_namespace"
)

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersServiceBindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersServiceBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersServiceBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersServiceBinding]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersServiceBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersServiceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersServiceBindingTypeService WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersServiceBindingType = "service"
)

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                                                  `json:"script_name"`
	JSON       workersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDoBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDoBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDoBinding]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDoBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDoBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDoBindingTypeDurableObjectNamespace WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDoBindingType = "durable_object_namespace"
)

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersR2BindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersR2BindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersR2BindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersR2Binding]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersR2Binding) implementsWorkersForPlatformDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersR2BindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersR2BindingTypeR2Bucket WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersR2BindingType = "r2_bucket"
)

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersQueueBindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersQueueBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersQueueBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersQueueBinding]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersQueueBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersQueueBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersQueueBindingTypeQueue WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersQueueBindingType = "queue"
)

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersD1BindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersD1BindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersD1BindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersD1Binding]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersD1Binding) implementsWorkersForPlatformDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersD1BindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersD1BindingTypeD1 WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersD1BindingType = "d1"
)

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     workersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON     `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBinding]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   workersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Outbound worker
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                                               `json:"service"`
	JSON    workersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersMtlsCertBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersMtlsCertBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                                                        `json:"certificate_id"`
	JSON          workersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersMtlsCertBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersMtlsCertBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersMtlsCertBinding]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersMtlsCertBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersMtlsCertBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersMtlsCertBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersMtlsCertBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersMtlsCertBindingTypeMtlsCertificate WorkersForPlatformDispatchNamespaceScriptSettingGetResponseBindingsWorkersMtlsCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrations]
// or
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrations].
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrations interface {
	implementsWorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrations)(nil)).Elem(), "")
}

// A single set of migrations to apply.
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses []WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               workersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON               `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrations]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrations) implementsWorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrations() {
}

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From string                                                                                                           `json:"from"`
	To   string                                                                                                           `json:"to"`
	JSON workersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       string                                                                                                               `json:"from"`
	FromScript string                                                                                                               `json:"from_script"`
	To         string                                                                                                               `json:"to"`
	JSON       workersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStep `json:"steps"`
	JSON  workersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsJSON   `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrations]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrations) implementsWorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrations() {
}

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               workersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON                `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStep]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From string                                                                                                             `json:"from"`
	To   string                                                                                                             `json:"to"`
	JSON workersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                                                                 `json:"from"`
	FromScript string                                                                                                                 `json:"from_script"`
	To         string                                                                                                                 `json:"to"`
	JSON       workersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode WorkersForPlatformDispatchNamespaceScriptSettingGetResponsePlacementMode `json:"mode"`
	JSON workersForPlatformDispatchNamespaceScriptSettingGetResponsePlacementJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponsePlacementJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponsePlacement]
type workersForPlatformDispatchNamespaceScriptSettingGetResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponsePlacementMode string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingGetResponsePlacementModeSmart WorkersForPlatformDispatchNamespaceScriptSettingGetResponsePlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                                      `json:"namespace"`
	JSON      workersForPlatformDispatchNamespaceScriptSettingGetResponseTailConsumerJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseTailConsumerJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseTailConsumer]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParams struct {
	Errors   param.Field[[]WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsError]   `json:"errors,required"`
	Messages param.Field[[]WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsMessage] `json:"messages,required"`
	Result   param.Field[WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResult]    `json:"result,required"`
	// Whether the API call was successful
	Success param.Field[WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsSuccess] `json:"success,required"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsError struct {
	Code    param.Field[int64]  `json:"code,required"`
	Message param.Field[string] `json:"message,required"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsError) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsMessage struct {
	Code    param.Field[int64]  `json:"code,required"`
	Message param.Field[string] `json:"message,required"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResult struct {
	// List of bindings attached to this Worker
	Bindings param.Field[[]WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBinding] `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations param.Field[WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultMigrations] `json:"migrations"`
	Placement  param.Field[WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultPlacement]  `json:"placement"`
	// Tags to help you manage your Workers
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultTailConsumer] `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel param.Field[string] `json:"usage_model"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResult) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A binding to allow the Worker to communicate with resources
//
// Satisfied by
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersKvNamespaceBinding],
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersServiceBinding],
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersDoBinding],
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersR2Binding],
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersQueueBinding],
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersD1Binding],
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersDispatchNamespaceBinding],
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersMtlsCertBinding].
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBinding interface {
	implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBinding()
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersKvNamespaceBinding struct {
	// The class of resource that the binding provides.
	Type param.Field[WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersKvNamespaceBindingType] `json:"type,required"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersKvNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersKvNamespaceBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersKvNamespaceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersKvNamespaceBindingTypeKvNamespace WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersKvNamespaceBindingType = "kv_namespace"
)

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// Name of Worker to bind to
	Service param.Field[string] `json:"service,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersServiceBindingType] `json:"type,required"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersServiceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersServiceBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersServiceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersServiceBindingTypeService WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersServiceBindingType = "service"
)

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName param.Field[string] `json:"class_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersDoBindingType] `json:"type,required"`
	// The environment of the script_name to bind to
	Environment param.Field[string] `json:"environment"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName param.Field[string] `json:"script_name"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersDoBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersDoBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersDoBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersDoBindingTypeDurableObjectNamespace WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersDoBindingType = "durable_object_namespace"
)

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName param.Field[string] `json:"bucket_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersR2BindingType] `json:"type,required"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersR2Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersR2Binding) implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersR2BindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersR2BindingTypeR2Bucket WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersR2BindingType = "r2_bucket"
)

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersQueueBinding struct {
	// Name of the Queue to bind to
	QueueName param.Field[string] `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersQueueBindingType] `json:"type,required"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersQueueBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersQueueBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersQueueBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersQueueBindingTypeQueue WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersQueueBindingType = "queue"
)

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID param.Field[string] `json:"id,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name param.Field[string] `json:"name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersD1BindingType] `json:"type,required"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersD1Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersD1Binding) implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersD1BindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersD1BindingTypeD1 WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersD1BindingType = "d1"
)

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersDispatchNamespaceBinding struct {
	// Namespace to bind to
	Namespace param.Field[string] `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersDispatchNamespaceBindingType] `json:"type,required"`
	// Outbound worker
	Outbound param.Field[WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersDispatchNamespaceBindingOutbound] `json:"outbound"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersDispatchNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersDispatchNamespaceBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersDispatchNamespaceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params param.Field[[]string] `json:"params"`
	// Outbound worker
	Worker param.Field[WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker] `json:"worker"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersDispatchNamespaceBindingOutbound) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker
	Service param.Field[string] `json:"service"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersMtlsCertBinding struct {
	// The class of resource that the binding provides.
	Type param.Field[WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersMtlsCertBindingType] `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID param.Field[string] `json:"certificate_id"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersMtlsCertBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersMtlsCertBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersMtlsCertBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersMtlsCertBindingTypeMtlsCertificate WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultBindingsWorkersMtlsCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultMigrationsWorkersSingleStepMigrations],
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultMigrationsWorkersSteppedMigrations].
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultMigrations interface {
	implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultMigrations()
}

// A single set of migrations to apply.
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses param.Field[[]WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultMigrationsWorkersSingleStepMigrationsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultMigrationsWorkersSingleStepMigrationsTransferredClass] `json:"transferred_classes"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultMigrationsWorkersSingleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultMigrationsWorkersSingleStepMigrations) implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultMigrations() {
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultMigrationsWorkersSingleStepMigrationsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultMigrationsWorkersSingleStepMigrationsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultMigrationsWorkersSteppedMigrationsStep] `json:"steps"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultMigrationsWorkersSteppedMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultMigrationsWorkersSteppedMigrations) implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultMigrations() {
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass] `json:"transferred_classes"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultMigrationsWorkersSteppedMigrationsStep) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode param.Field[WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultPlacementMode] `json:"mode"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultPlacementMode string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultPlacementModeSmart WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultPlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service param.Field[string] `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace param.Field[string] `json:"namespace"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResultTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the API call was successful
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsSuccess bool

const (
	WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsSuccessTrue WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsSuccess = true
)

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseEnvelope struct {
	Errors   []WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    workersForPlatformDispatchNamespaceScriptSettingUpdateResponseEnvelopeJSON    `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseEnvelope]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseEnvelopeErrors struct {
	Code    int64                                                                            `json:"code,required"`
	Message string                                                                           `json:"message,required"`
	JSON    workersForPlatformDispatchNamespaceScriptSettingUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseEnvelopeErrors]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseEnvelopeMessages struct {
	Code    int64                                                                              `json:"code,required"`
	Message string                                                                             `json:"message,required"`
	JSON    workersForPlatformDispatchNamespaceScriptSettingUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseEnvelopeMessages]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseEnvelopeSuccess bool

const (
	WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseEnvelopeSuccessTrue WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseEnvelopeSuccess = true
)

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseEnvelope struct {
	Errors   []WorkersForPlatformDispatchNamespaceScriptSettingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkersForPlatformDispatchNamespaceScriptSettingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkersForPlatformDispatchNamespaceScriptSettingGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkersForPlatformDispatchNamespaceScriptSettingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    workersForPlatformDispatchNamespaceScriptSettingGetResponseEnvelopeJSON    `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseEnvelope]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseEnvelopeErrors struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    workersForPlatformDispatchNamespaceScriptSettingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseEnvelopeErrors]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseEnvelopeMessages struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    workersForPlatformDispatchNamespaceScriptSettingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseEnvelopeMessages]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseEnvelopeSuccess bool

const (
	WorkersForPlatformDispatchNamespaceScriptSettingGetResponseEnvelopeSuccessTrue WorkersForPlatformDispatchNamespaceScriptSettingGetResponseEnvelopeSuccess = true
)
