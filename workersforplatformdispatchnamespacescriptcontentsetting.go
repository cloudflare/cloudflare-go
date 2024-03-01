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

// WorkersForPlatformDispatchNamespaceScriptContentSettingService contains methods
// and other services that help with interacting with the cloudflare API. Note,
// unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkersForPlatformDispatchNamespaceScriptContentSettingService] method
// instead.
type WorkersForPlatformDispatchNamespaceScriptContentSettingService struct {
	Options []option.RequestOption
}

// NewWorkersForPlatformDispatchNamespaceScriptContentSettingService generates a
// new service that applies the given options to each request. These options are
// applied after the parent client's options (if there is one), and before any
// request-specific options.
func NewWorkersForPlatformDispatchNamespaceScriptContentSettingService(opts ...option.RequestOption) (r *WorkersForPlatformDispatchNamespaceScriptContentSettingService) {
	r = &WorkersForPlatformDispatchNamespaceScriptContentSettingService{}
	r.Options = opts
	return
}

// Patch script metadata, such as bindings
func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingService) Edit(ctx context.Context, dispatchNamespace string, scriptName string, params WorkersForPlatformDispatchNamespaceScriptContentSettingEditParams, opts ...option.RequestOption) (res *WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/settings", params.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get script settings from a script uploaded to a Workers for Platforms namespace.
func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingService) Get(ctx context.Context, dispatchNamespace string, scriptName string, query WorkersForPlatformDispatchNamespaceScriptContentSettingGetParams, opts ...option.RequestOption) (res *WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/settings", query.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponse struct {
	// List of bindings attached to this Worker
	Bindings []WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrations `json:"migrations"`
	Placement  WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponsePlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                                                  `json:"usage_model"`
	JSON       workersForPlatformDispatchNamespaceScriptContentSettingEditResponseJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingEditResponseJSON contains
// the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponse]
type workersForPlatformDispatchNamespaceScriptContentSettingEditResponseJSON struct {
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

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersKVNamespaceBinding],
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersServiceBinding],
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDoBinding],
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersR2Binding],
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersQueueBinding],
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersD1Binding],
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBinding]
// or
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersMTLSCertBinding].
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBinding interface {
	implementsWorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBinding()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBinding)(nil)).Elem(), "")
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersKVNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersKVNamespaceBindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersKVNamespaceBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersKVNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersKVNamespaceBinding]
type workersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersKVNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersKVNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersKVNamespaceBinding) implementsWorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersKVNamespaceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersKVNamespaceBindingTypeKVNamespace WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersKVNamespaceBindingType = "kv_namespace"
)

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersServiceBindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersServiceBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersServiceBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersServiceBinding]
type workersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersServiceBinding) implementsWorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersServiceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersServiceBindingTypeService WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersServiceBindingType = "service"
)

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                                                          `json:"script_name"`
	JSON       workersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDoBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDoBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDoBinding]
type workersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDoBinding) implementsWorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDoBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDoBindingTypeDurableObjectNamespace WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDoBindingType = "durable_object_namespace"
)

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersR2BindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersR2BindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersR2BindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersR2Binding]
type workersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersR2Binding) implementsWorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersR2BindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersR2BindingTypeR2Bucket WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersR2BindingType = "r2_bucket"
)

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersQueueBindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersQueueBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersQueueBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersQueueBinding]
type workersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersQueueBinding) implementsWorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersQueueBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersQueueBindingTypeQueue WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersQueueBindingType = "queue"
)

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersD1BindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersD1BindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersD1BindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersD1Binding]
type workersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersD1Binding) implementsWorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersD1BindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersD1BindingTypeD1 WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersD1BindingType = "d1"
)

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     workersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON     `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBinding]
type workersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBinding) implementsWorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   workersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound]
type workersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Outbound worker
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                                                       `json:"service"`
	JSON    workersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker]
type workersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersMTLSCertBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersMTLSCertBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                                                                `json:"certificate_id"`
	JSON          workersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersMTLSCertBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersMTLSCertBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersMTLSCertBinding]
type workersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersMTLSCertBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersMTLSCertBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersMTLSCertBinding) implementsWorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersMTLSCertBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersMTLSCertBindingTypeMTLSCertificate WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseBindingsWorkersMTLSCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrations]
// or
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrations].
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrations interface {
	implementsWorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrations)(nil)).Elem(), "")
}

// A single set of migrations to apply.
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses []WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               workersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON               `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrations]
type workersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrations) implementsWorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrations() {
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From string                                                                                                                   `json:"from"`
	To   string                                                                                                                   `json:"to"`
	JSON workersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass]
type workersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       string                                                                                                                       `json:"from"`
	FromScript string                                                                                                                       `json:"from_script"`
	To         string                                                                                                                       `json:"to"`
	JSON       workersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass]
type workersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStep `json:"steps"`
	JSON  workersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsJSON   `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrations]
type workersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrations) implementsWorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrations() {
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               workersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON                `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStep]
type workersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From string                                                                                                                     `json:"from"`
	To   string                                                                                                                     `json:"to"`
	JSON workersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass]
type workersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                                                                         `json:"from"`
	FromScript string                                                                                                                         `json:"from_script"`
	To         string                                                                                                                         `json:"to"`
	JSON       workersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass]
type workersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponsePlacementMode `json:"mode"`
	JSON workersForPlatformDispatchNamespaceScriptContentSettingEditResponsePlacementJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingEditResponsePlacementJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponsePlacement]
type workersForPlatformDispatchNamespaceScriptContentSettingEditResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponsePlacementMode string

const (
	WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponsePlacementModeSmart WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponsePlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                                              `json:"namespace"`
	JSON      workersForPlatformDispatchNamespaceScriptContentSettingEditResponseTailConsumerJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingEditResponseTailConsumerJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseTailConsumer]
type workersForPlatformDispatchNamespaceScriptContentSettingEditResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponse struct {
	// List of bindings attached to this Worker
	Bindings []WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrations `json:"migrations"`
	Placement  WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponsePlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                                                 `json:"usage_model"`
	JSON       workersForPlatformDispatchNamespaceScriptContentSettingGetResponseJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingGetResponseJSON contains
// the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponse]
type workersForPlatformDispatchNamespaceScriptContentSettingGetResponseJSON struct {
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

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersKVNamespaceBinding],
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersServiceBinding],
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDoBinding],
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersR2Binding],
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersQueueBinding],
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersD1Binding],
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBinding]
// or
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersMTLSCertBinding].
type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBinding interface {
	implementsWorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBinding()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBinding)(nil)).Elem(), "")
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersKVNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersKVNamespaceBindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersKVNamespaceBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersKVNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersKVNamespaceBinding]
type workersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersKVNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersKVNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersKVNamespaceBinding) implementsWorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersKVNamespaceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersKVNamespaceBindingTypeKVNamespace WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersKVNamespaceBindingType = "kv_namespace"
)

type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersServiceBindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersServiceBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersServiceBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersServiceBinding]
type workersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersServiceBinding) implementsWorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersServiceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersServiceBindingTypeService WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersServiceBindingType = "service"
)

type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                                                         `json:"script_name"`
	JSON       workersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDoBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDoBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDoBinding]
type workersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDoBinding) implementsWorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDoBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDoBindingTypeDurableObjectNamespace WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDoBindingType = "durable_object_namespace"
)

type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersR2BindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersR2BindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersR2BindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersR2Binding]
type workersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersR2Binding) implementsWorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersR2BindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersR2BindingTypeR2Bucket WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersR2BindingType = "r2_bucket"
)

type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersQueueBindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersQueueBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersQueueBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersQueueBinding]
type workersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersQueueBinding) implementsWorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersQueueBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersQueueBindingTypeQueue WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersQueueBindingType = "queue"
)

type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersD1BindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersD1BindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersD1BindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersD1Binding]
type workersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersD1Binding) implementsWorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersD1BindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersD1BindingTypeD1 WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersD1BindingType = "d1"
)

type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     workersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON     `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBinding]
type workersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBinding) implementsWorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   workersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound]
type workersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Outbound worker
type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                                                      `json:"service"`
	JSON    workersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker]
type workersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersMTLSCertBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersMTLSCertBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                                                               `json:"certificate_id"`
	JSON          workersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersMTLSCertBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersMTLSCertBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersMTLSCertBinding]
type workersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersMTLSCertBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersMTLSCertBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersMTLSCertBinding) implementsWorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersMTLSCertBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersMTLSCertBindingTypeMTLSCertificate WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseBindingsWorkersMTLSCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrations]
// or
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrations].
type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrations interface {
	implementsWorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrations)(nil)).Elem(), "")
}

// A single set of migrations to apply.
type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses []WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               workersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON               `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrations]
type workersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrations) implementsWorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrations() {
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From string                                                                                                                  `json:"from"`
	To   string                                                                                                                  `json:"to"`
	JSON workersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass]
type workersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       string                                                                                                                      `json:"from"`
	FromScript string                                                                                                                      `json:"from_script"`
	To         string                                                                                                                      `json:"to"`
	JSON       workersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass]
type workersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStep `json:"steps"`
	JSON  workersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsJSON   `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrations]
type workersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrations) implementsWorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrations() {
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               workersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON                `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStep]
type workersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From string                                                                                                                    `json:"from"`
	To   string                                                                                                                    `json:"to"`
	JSON workersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass]
type workersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                                                                        `json:"from"`
	FromScript string                                                                                                                        `json:"from_script"`
	To         string                                                                                                                        `json:"to"`
	JSON       workersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass]
type workersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponsePlacementMode `json:"mode"`
	JSON workersForPlatformDispatchNamespaceScriptContentSettingGetResponsePlacementJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingGetResponsePlacementJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponsePlacement]
type workersForPlatformDispatchNamespaceScriptContentSettingGetResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponsePlacementMode string

const (
	WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponsePlacementModeSmart WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponsePlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                                             `json:"namespace"`
	JSON      workersForPlatformDispatchNamespaceScriptContentSettingGetResponseTailConsumerJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingGetResponseTailConsumerJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseTailConsumer]
type workersForPlatformDispatchNamespaceScriptContentSettingGetResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParams struct {
	// Identifier
	AccountID param.Field[string]                                                                     `path:"account_id,required"`
	Errors    param.Field[[]WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsError]   `json:"errors,required"`
	Messages  param.Field[[]WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsMessage] `json:"messages,required"`
	Result    param.Field[WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResult]    `json:"result,required"`
	// Whether the API call was successful
	Success param.Field[WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsSuccess] `json:"success,required"`
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsError struct {
	Code    param.Field[int64]  `json:"code,required"`
	Message param.Field[string] `json:"message,required"`
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsError) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsMessage struct {
	Code    param.Field[int64]  `json:"code,required"`
	Message param.Field[string] `json:"message,required"`
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResult struct {
	// List of bindings attached to this Worker
	Bindings param.Field[[]WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBinding] `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations param.Field[WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrations] `json:"migrations"`
	Placement  param.Field[WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultPlacement]  `json:"placement"`
	// Tags to help you manage your Workers
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultTailConsumer] `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel param.Field[string] `json:"usage_model"`
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResult) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A binding to allow the Worker to communicate with resources
//
// Satisfied by
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersKVNamespaceBinding],
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersServiceBinding],
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDoBinding],
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersR2Binding],
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersQueueBinding],
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersD1Binding],
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDispatchNamespaceBinding],
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersMTLSCertBinding].
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBinding interface {
	implementsWorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBinding()
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersKVNamespaceBinding struct {
	// The class of resource that the binding provides.
	Type param.Field[WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersKVNamespaceBindingType] `json:"type,required"`
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersKVNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersKVNamespaceBinding) implementsWorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersKVNamespaceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersKVNamespaceBindingTypeKVNamespace WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersKVNamespaceBindingType = "kv_namespace"
)

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// Name of Worker to bind to
	Service param.Field[string] `json:"service,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersServiceBindingType] `json:"type,required"`
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersServiceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersServiceBinding) implementsWorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersServiceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersServiceBindingTypeService WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersServiceBindingType = "service"
)

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName param.Field[string] `json:"class_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDoBindingType] `json:"type,required"`
	// The environment of the script_name to bind to
	Environment param.Field[string] `json:"environment"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName param.Field[string] `json:"script_name"`
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDoBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDoBinding) implementsWorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDoBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDoBindingTypeDurableObjectNamespace WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDoBindingType = "durable_object_namespace"
)

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName param.Field[string] `json:"bucket_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersR2BindingType] `json:"type,required"`
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersR2Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersR2Binding) implementsWorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersR2BindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersR2BindingTypeR2Bucket WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersR2BindingType = "r2_bucket"
)

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersQueueBinding struct {
	// Name of the Queue to bind to
	QueueName param.Field[string] `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersQueueBindingType] `json:"type,required"`
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersQueueBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersQueueBinding) implementsWorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersQueueBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersQueueBindingTypeQueue WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersQueueBindingType = "queue"
)

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID param.Field[string] `json:"id,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name param.Field[string] `json:"name,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersD1BindingType] `json:"type,required"`
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersD1Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersD1Binding) implementsWorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersD1BindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersD1BindingTypeD1 WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersD1BindingType = "d1"
)

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDispatchNamespaceBinding struct {
	// Namespace to bind to
	Namespace param.Field[string] `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type param.Field[WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingType] `json:"type,required"`
	// Outbound worker
	Outbound param.Field[WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutbound] `json:"outbound"`
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDispatchNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDispatchNamespaceBinding) implementsWorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params param.Field[[]string] `json:"params"`
	// Outbound worker
	Worker param.Field[WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker] `json:"worker"`
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutbound) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker
	Service param.Field[string] `json:"service"`
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersMTLSCertBinding struct {
	// The class of resource that the binding provides.
	Type param.Field[WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersMTLSCertBindingType] `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID param.Field[string] `json:"certificate_id"`
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersMTLSCertBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersMTLSCertBinding) implementsWorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersMTLSCertBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersMTLSCertBindingTypeMTLSCertificate WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultBindingsWorkersMTLSCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSingleStepMigrations],
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSteppedMigrations].
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrations interface {
	implementsWorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrations()
}

// A single set of migrations to apply.
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses param.Field[[]WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSingleStepMigrationsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSingleStepMigrationsTransferredClass] `json:"transferred_classes"`
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSingleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSingleStepMigrations) implementsWorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrations() {
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSingleStepMigrationsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSingleStepMigrationsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStep] `json:"steps"`
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSteppedMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSteppedMigrations) implementsWorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrations() {
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass] `json:"transferred_classes"`
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStep) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode param.Field[WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultPlacementMode] `json:"mode"`
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultPlacementMode string

const (
	WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultPlacementModeSmart WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultPlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service param.Field[string] `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace param.Field[string] `json:"namespace"`
}

func (r WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsResultTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the API call was successful
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsSuccess bool

const (
	WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsSuccessTrue WorkersForPlatformDispatchNamespaceScriptContentSettingEditParamsSuccess = true
)

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseEnvelope struct {
	Errors   []WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    workersForPlatformDispatchNamespaceScriptContentSettingEditResponseEnvelopeJSON    `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingEditResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseEnvelope]
type workersForPlatformDispatchNamespaceScriptContentSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseEnvelopeErrors struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    workersForPlatformDispatchNamespaceScriptContentSettingEditResponseEnvelopeErrorsJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingEditResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseEnvelopeErrors]
type workersForPlatformDispatchNamespaceScriptContentSettingEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseEnvelopeMessages struct {
	Code    int64                                                                                   `json:"code,required"`
	Message string                                                                                  `json:"message,required"`
	JSON    workersForPlatformDispatchNamespaceScriptContentSettingEditResponseEnvelopeMessagesJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingEditResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseEnvelopeMessages]
type workersForPlatformDispatchNamespaceScriptContentSettingEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseEnvelopeSuccess bool

const (
	WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseEnvelopeSuccessTrue WorkersForPlatformDispatchNamespaceScriptContentSettingEditResponseEnvelopeSuccess = true
)

type WorkersForPlatformDispatchNamespaceScriptContentSettingGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseEnvelope struct {
	Errors   []WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    workersForPlatformDispatchNamespaceScriptContentSettingGetResponseEnvelopeJSON    `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingGetResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseEnvelope]
type workersForPlatformDispatchNamespaceScriptContentSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseEnvelopeErrors struct {
	Code    int64                                                                                `json:"code,required"`
	Message string                                                                               `json:"message,required"`
	JSON    workersForPlatformDispatchNamespaceScriptContentSettingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingGetResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseEnvelopeErrors]
type workersForPlatformDispatchNamespaceScriptContentSettingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseEnvelopeMessages struct {
	Code    int64                                                                                  `json:"code,required"`
	Message string                                                                                 `json:"message,required"`
	JSON    workersForPlatformDispatchNamespaceScriptContentSettingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentSettingGetResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseEnvelopeMessages]
type workersForPlatformDispatchNamespaceScriptContentSettingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseEnvelopeSuccess bool

const (
	WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseEnvelopeSuccessTrue WorkersForPlatformDispatchNamespaceScriptContentSettingGetResponseEnvelopeSuccess = true
)
