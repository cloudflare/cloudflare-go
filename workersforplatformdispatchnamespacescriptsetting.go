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
func (r *WorkersForPlatformDispatchNamespaceScriptSettingService) Edit(ctx context.Context, accountID string, dispatchNamespace string, scriptName string, body WorkersForPlatformDispatchNamespaceScriptSettingEditParams, opts ...option.RequestOption) (res *WorkersForPlatformDispatchNamespaceScriptSettingEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkersForPlatformDispatchNamespaceScriptSettingEditResponseEnvelope
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

type WorkersForPlatformDispatchNamespaceScriptSettingEditResponse struct {
	// List of bindings attached to this Worker
	Bindings []WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrations `json:"migrations"`
	Placement  WorkersForPlatformDispatchNamespaceScriptSettingEditResponsePlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []WorkersForPlatformDispatchNamespaceScriptSettingEditResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                                           `json:"usage_model"`
	JSON       workersForPlatformDispatchNamespaceScriptSettingEditResponseJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingEditResponseJSON contains the
// JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponse]
type workersForPlatformDispatchNamespaceScriptSettingEditResponseJSON struct {
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

func (r *WorkersForPlatformDispatchNamespaceScriptSettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersKvNamespaceBinding],
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersServiceBinding],
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDoBinding],
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersR2Binding],
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersQueueBinding],
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersD1Binding],
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBinding]
// or
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersMtlsCertBinding].
type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBinding interface {
	implementsWorkersForPlatformDispatchNamespaceScriptSettingEditResponseBinding()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBinding)(nil)).Elem(), "")
}

type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersKvNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersKvNamespaceBindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersKvNamespaceBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersKvNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersKvNamespaceBinding]
type workersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersKvNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersKvNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersKvNamespaceBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersKvNamespaceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersKvNamespaceBindingTypeKvNamespace WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersKvNamespaceBindingType = "kv_namespace"
)

type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersServiceBindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersServiceBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersServiceBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersServiceBinding]
type workersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersServiceBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersServiceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersServiceBindingTypeService WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersServiceBindingType = "service"
)

type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                                                   `json:"script_name"`
	JSON       workersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDoBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDoBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDoBinding]
type workersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDoBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDoBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDoBindingTypeDurableObjectNamespace WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDoBindingType = "durable_object_namespace"
)

type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersR2BindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersR2BindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersR2BindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersR2Binding]
type workersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersR2Binding) implementsWorkersForPlatformDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersR2BindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersR2BindingTypeR2Bucket WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersR2BindingType = "r2_bucket"
)

type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersQueueBindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersQueueBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersQueueBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersQueueBinding]
type workersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersQueueBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersQueueBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersQueueBindingTypeQueue WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersQueueBindingType = "queue"
)

type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersD1BindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersD1BindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersD1BindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersD1Binding]
type workersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersD1Binding) implementsWorkersForPlatformDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersD1BindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersD1BindingTypeD1 WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersD1BindingType = "d1"
)

type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     workersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON     `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBinding]
type workersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   workersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound]
type workersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Outbound worker
type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                                                `json:"service"`
	JSON    workersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker]
type workersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersMtlsCertBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersMtlsCertBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                                                         `json:"certificate_id"`
	JSON          workersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersMtlsCertBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersMtlsCertBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersMtlsCertBinding]
type workersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersMtlsCertBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersMtlsCertBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersMtlsCertBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersMtlsCertBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersMtlsCertBindingTypeMtlsCertificate WorkersForPlatformDispatchNamespaceScriptSettingEditResponseBindingsWorkersMtlsCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrations]
// or
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrations].
type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrations interface {
	implementsWorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrations)(nil)).Elem(), "")
}

// A single set of migrations to apply.
type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses []WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               workersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON               `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrations]
type workersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrations) implementsWorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrations() {
}

type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From string                                                                                                            `json:"from"`
	To   string                                                                                                            `json:"to"`
	JSON workersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass]
type workersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       string                                                                                                                `json:"from"`
	FromScript string                                                                                                                `json:"from_script"`
	To         string                                                                                                                `json:"to"`
	JSON       workersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass]
type workersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStep `json:"steps"`
	JSON  workersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsJSON   `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrations]
type workersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrations) implementsWorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrations() {
}

type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               workersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON                `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStep]
type workersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From string                                                                                                              `json:"from"`
	To   string                                                                                                              `json:"to"`
	JSON workersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass]
type workersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                                                                  `json:"from"`
	FromScript string                                                                                                                  `json:"from_script"`
	To         string                                                                                                                  `json:"to"`
	JSON       workersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass]
type workersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingEditResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode WorkersForPlatformDispatchNamespaceScriptSettingEditResponsePlacementMode `json:"mode"`
	JSON workersForPlatformDispatchNamespaceScriptSettingEditResponsePlacementJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingEditResponsePlacementJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponsePlacement]
type workersForPlatformDispatchNamespaceScriptSettingEditResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingEditResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type WorkersForPlatformDispatchNamespaceScriptSettingEditResponsePlacementMode string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingEditResponsePlacementModeSmart WorkersForPlatformDispatchNamespaceScriptSettingEditResponsePlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                                       `json:"namespace"`
	JSON      workersForPlatformDispatchNamespaceScriptSettingEditResponseTailConsumerJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingEditResponseTailConsumerJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponseTailConsumer]
type workersForPlatformDispatchNamespaceScriptSettingEditResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingEditResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
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

type WorkersForPlatformDispatchNamespaceScriptSettingEditParams struct {
	Errors   param.Field[[]WorkersForPlatformDispatchNamespaceScriptSettingEditParamsError]   `json:"errors,required"`
	Messages param.Field[[]WorkersForPlatformDispatchNamespaceScriptSettingEditParamsMessage] `json:"messages,required"`
	Result   param.Field[WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResult]    `json:"result,required"`
	// Whether the API call was successful
	Success param.Field[WorkersForPlatformDispatchNamespaceScriptSettingEditParamsSuccess] `json:"success,required"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsError struct {
	Code    param.Field[int64]  `json:"code,required"`
	Message param.Field[string] `json:"message,required"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsError) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsMessage struct {
	Code    param.Field[int64]  `json:"code,required"`
	Message param.Field[string] `json:"message,required"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResult struct {
	// List of bindings attached to this Worker
	Bindings param.Field[[]WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBinding] `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations param.Field[WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultMigrations] `json:"migrations"`
	Placement  param.Field[WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultPlacement]  `json:"placement"`
	// Tags to help you manage your Workers
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultTailConsumer] `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel param.Field[string] `json:"usage_model"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResult) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A binding to allow the Worker to communicate with resources
//
// Satisfied by
// [WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersKvNamespaceBinding],
// [WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersServiceBinding],
// [WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDoBinding],
// [WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersR2Binding],
// [WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersQueueBinding],
// [WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersD1Binding],
// [WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDispatchNamespaceBinding],
// [WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersMtlsCertBinding].
type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBinding interface {
	implementsWorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBinding()
}

type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersKvNamespaceBinding struct {
	// The class of resource that the binding provides.
	Type param.Field[WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersKvNamespaceBindingType] `json:"type,required"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersKvNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersKvNamespaceBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersKvNamespaceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersKvNamespaceBindingTypeKvNamespace WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersKvNamespaceBindingType = "kv_namespace"
)

type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// Name of Worker to bind to
	Service param.Field[string] `json:"service,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersServiceBindingType] `json:"type,required"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersServiceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersServiceBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersServiceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersServiceBindingTypeService WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersServiceBindingType = "service"
)

type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName param.Field[string] `json:"class_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDoBindingType] `json:"type,required"`
	// The environment of the script_name to bind to
	Environment param.Field[string] `json:"environment"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName param.Field[string] `json:"script_name"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDoBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDoBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDoBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDoBindingTypeDurableObjectNamespace WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDoBindingType = "durable_object_namespace"
)

type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName param.Field[string] `json:"bucket_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersR2BindingType] `json:"type,required"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersR2Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersR2Binding) implementsWorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersR2BindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersR2BindingTypeR2Bucket WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersR2BindingType = "r2_bucket"
)

type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersQueueBinding struct {
	// Name of the Queue to bind to
	QueueName param.Field[string] `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersQueueBindingType] `json:"type,required"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersQueueBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersQueueBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersQueueBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersQueueBindingTypeQueue WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersQueueBindingType = "queue"
)

type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID param.Field[string] `json:"id,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name param.Field[string] `json:"name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersD1BindingType] `json:"type,required"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersD1Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersD1Binding) implementsWorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersD1BindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersD1BindingTypeD1 WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersD1BindingType = "d1"
)

type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDispatchNamespaceBinding struct {
	// Namespace to bind to
	Namespace param.Field[string] `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingType] `json:"type,required"`
	// Outbound worker
	Outbound param.Field[WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutbound] `json:"outbound"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDispatchNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDispatchNamespaceBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params param.Field[[]string] `json:"params"`
	// Outbound worker
	Worker param.Field[WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker] `json:"worker"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutbound) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker
type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker
	Service param.Field[string] `json:"service"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersMtlsCertBinding struct {
	// The class of resource that the binding provides.
	Type param.Field[WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersMtlsCertBindingType] `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID param.Field[string] `json:"certificate_id"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersMtlsCertBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersMtlsCertBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersMtlsCertBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersMtlsCertBindingTypeMtlsCertificate WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultBindingsWorkersMtlsCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by
// [WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSingleStepMigrations],
// [WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSteppedMigrations].
type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultMigrations interface {
	implementsWorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultMigrations()
}

// A single set of migrations to apply.
type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses param.Field[[]WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSingleStepMigrationsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSingleStepMigrationsTransferredClass] `json:"transferred_classes"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSingleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSingleStepMigrations) implementsWorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultMigrations() {
}

type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSingleStepMigrationsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSingleStepMigrationsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSteppedMigrationsStep] `json:"steps"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSteppedMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSteppedMigrations) implementsWorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultMigrations() {
}

type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass] `json:"transferred_classes"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSteppedMigrationsStep) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode param.Field[WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultPlacementMode] `json:"mode"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultPlacementMode string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultPlacementModeSmart WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultPlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service param.Field[string] `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace param.Field[string] `json:"namespace"`
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingEditParamsResultTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the API call was successful
type WorkersForPlatformDispatchNamespaceScriptSettingEditParamsSuccess bool

const (
	WorkersForPlatformDispatchNamespaceScriptSettingEditParamsSuccessTrue WorkersForPlatformDispatchNamespaceScriptSettingEditParamsSuccess = true
)

type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseEnvelope struct {
	Errors   []WorkersForPlatformDispatchNamespaceScriptSettingEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkersForPlatformDispatchNamespaceScriptSettingEditResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkersForPlatformDispatchNamespaceScriptSettingEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkersForPlatformDispatchNamespaceScriptSettingEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    workersForPlatformDispatchNamespaceScriptSettingEditResponseEnvelopeJSON    `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingEditResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponseEnvelope]
type workersForPlatformDispatchNamespaceScriptSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseEnvelopeErrors struct {
	Code    int64                                                                          `json:"code,required"`
	Message string                                                                         `json:"message,required"`
	JSON    workersForPlatformDispatchNamespaceScriptSettingEditResponseEnvelopeErrorsJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingEditResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponseEnvelopeErrors]
type workersForPlatformDispatchNamespaceScriptSettingEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseEnvelopeMessages struct {
	Code    int64                                                                            `json:"code,required"`
	Message string                                                                           `json:"message,required"`
	JSON    workersForPlatformDispatchNamespaceScriptSettingEditResponseEnvelopeMessagesJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingEditResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingEditResponseEnvelopeMessages]
type workersForPlatformDispatchNamespaceScriptSettingEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkersForPlatformDispatchNamespaceScriptSettingEditResponseEnvelopeSuccess bool

const (
	WorkersForPlatformDispatchNamespaceScriptSettingEditResponseEnvelopeSuccessTrue WorkersForPlatformDispatchNamespaceScriptSettingEditResponseEnvelopeSuccess = true
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
