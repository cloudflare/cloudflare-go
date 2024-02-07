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
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/settings", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Patch script metadata or config, such as bindings or usage model
func (r *WorkerScriptSettingService) Update(ctx context.Context, accountID string, scriptName string, body WorkerScriptSettingUpdateParams, opts ...option.RequestOption) (res *WorkerScriptSettingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/settings", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type WorkerScriptSettingGetResponse struct {
	Errors   []WorkerScriptSettingGetResponseError   `json:"errors"`
	Messages []WorkerScriptSettingGetResponseMessage `json:"messages"`
	Result   WorkerScriptSettingGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success WorkerScriptSettingGetResponseSuccess `json:"success"`
	JSON    workerScriptSettingGetResponseJSON    `json:"-"`
}

// workerScriptSettingGetResponseJSON contains the JSON metadata for the struct
// [WorkerScriptSettingGetResponse]
type workerScriptSettingGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptSettingGetResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    workerScriptSettingGetResponseErrorJSON `json:"-"`
}

// workerScriptSettingGetResponseErrorJSON contains the JSON metadata for the
// struct [WorkerScriptSettingGetResponseError]
type workerScriptSettingGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptSettingGetResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    workerScriptSettingGetResponseMessageJSON `json:"-"`
}

// workerScriptSettingGetResponseMessageJSON contains the JSON metadata for the
// struct [WorkerScriptSettingGetResponseMessage]
type workerScriptSettingGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptSettingGetResponseResult struct {
	// List of bindings attached to this Worker
	Bindings []WorkerScriptSettingGetResponseResultBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations WorkerScriptSettingGetResponseResultMigrations `json:"migrations"`
	Placement  WorkerScriptSettingGetResponseResultPlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []WorkerScriptSettingGetResponseResultTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                   `json:"usage_model"`
	JSON       workerScriptSettingGetResponseResultJSON `json:"-"`
}

// workerScriptSettingGetResponseResultJSON contains the JSON metadata for the
// struct [WorkerScriptSettingGetResponseResult]
type workerScriptSettingGetResponseResultJSON struct {
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

func (r *WorkerScriptSettingGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [WorkerScriptSettingGetResponseResultBindingsWorkersKvNamespaceBinding],
// [WorkerScriptSettingGetResponseResultBindingsWorkersServiceBinding],
// [WorkerScriptSettingGetResponseResultBindingsWorkersDoBinding],
// [WorkerScriptSettingGetResponseResultBindingsWorkersR2Binding],
// [WorkerScriptSettingGetResponseResultBindingsWorkersQueueBinding],
// [WorkerScriptSettingGetResponseResultBindingsWorkersD1Binding],
// [WorkerScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBinding] or
// [WorkerScriptSettingGetResponseResultBindingsWorkersMtlsCertBinding].
type WorkerScriptSettingGetResponseResultBinding interface {
	implementsWorkerScriptSettingGetResponseResultBinding()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkerScriptSettingGetResponseResultBinding)(nil)).Elem(), "")
}

type WorkerScriptSettingGetResponseResultBindingsWorkersKvNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingGetResponseResultBindingsWorkersKvNamespaceBindingType `json:"type,required"`
	JSON workerScriptSettingGetResponseResultBindingsWorkersKvNamespaceBindingJSON `json:"-"`
}

// workerScriptSettingGetResponseResultBindingsWorkersKvNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingGetResponseResultBindingsWorkersKvNamespaceBinding]
type workerScriptSettingGetResponseResultBindingsWorkersKvNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseResultBindingsWorkersKvNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingGetResponseResultBindingsWorkersKvNamespaceBinding) implementsWorkerScriptSettingGetResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingGetResponseResultBindingsWorkersKvNamespaceBindingType string

const (
	WorkerScriptSettingGetResponseResultBindingsWorkersKvNamespaceBindingTypeKvNamespace WorkerScriptSettingGetResponseResultBindingsWorkersKvNamespaceBindingType = "kv_namespace"
)

type WorkerScriptSettingGetResponseResultBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingGetResponseResultBindingsWorkersServiceBindingType `json:"type,required"`
	JSON workerScriptSettingGetResponseResultBindingsWorkersServiceBindingJSON `json:"-"`
}

// workerScriptSettingGetResponseResultBindingsWorkersServiceBindingJSON contains
// the JSON metadata for the struct
// [WorkerScriptSettingGetResponseResultBindingsWorkersServiceBinding]
type workerScriptSettingGetResponseResultBindingsWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseResultBindingsWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingGetResponseResultBindingsWorkersServiceBinding) implementsWorkerScriptSettingGetResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingGetResponseResultBindingsWorkersServiceBindingType string

const (
	WorkerScriptSettingGetResponseResultBindingsWorkersServiceBindingTypeService WorkerScriptSettingGetResponseResultBindingsWorkersServiceBindingType = "service"
)

type WorkerScriptSettingGetResponseResultBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingGetResponseResultBindingsWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                           `json:"script_name"`
	JSON       workerScriptSettingGetResponseResultBindingsWorkersDoBindingJSON `json:"-"`
}

// workerScriptSettingGetResponseResultBindingsWorkersDoBindingJSON contains the
// JSON metadata for the struct
// [WorkerScriptSettingGetResponseResultBindingsWorkersDoBinding]
type workerScriptSettingGetResponseResultBindingsWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseResultBindingsWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingGetResponseResultBindingsWorkersDoBinding) implementsWorkerScriptSettingGetResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingGetResponseResultBindingsWorkersDoBindingType string

const (
	WorkerScriptSettingGetResponseResultBindingsWorkersDoBindingTypeDurableObjectNamespace WorkerScriptSettingGetResponseResultBindingsWorkersDoBindingType = "durable_object_namespace"
)

type WorkerScriptSettingGetResponseResultBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingGetResponseResultBindingsWorkersR2BindingType `json:"type,required"`
	JSON workerScriptSettingGetResponseResultBindingsWorkersR2BindingJSON `json:"-"`
}

// workerScriptSettingGetResponseResultBindingsWorkersR2BindingJSON contains the
// JSON metadata for the struct
// [WorkerScriptSettingGetResponseResultBindingsWorkersR2Binding]
type workerScriptSettingGetResponseResultBindingsWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseResultBindingsWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingGetResponseResultBindingsWorkersR2Binding) implementsWorkerScriptSettingGetResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingGetResponseResultBindingsWorkersR2BindingType string

const (
	WorkerScriptSettingGetResponseResultBindingsWorkersR2BindingTypeR2Bucket WorkerScriptSettingGetResponseResultBindingsWorkersR2BindingType = "r2_bucket"
)

type WorkerScriptSettingGetResponseResultBindingsWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingGetResponseResultBindingsWorkersQueueBindingType `json:"type,required"`
	JSON workerScriptSettingGetResponseResultBindingsWorkersQueueBindingJSON `json:"-"`
}

// workerScriptSettingGetResponseResultBindingsWorkersQueueBindingJSON contains the
// JSON metadata for the struct
// [WorkerScriptSettingGetResponseResultBindingsWorkersQueueBinding]
type workerScriptSettingGetResponseResultBindingsWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseResultBindingsWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingGetResponseResultBindingsWorkersQueueBinding) implementsWorkerScriptSettingGetResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingGetResponseResultBindingsWorkersQueueBindingType string

const (
	WorkerScriptSettingGetResponseResultBindingsWorkersQueueBindingTypeQueue WorkerScriptSettingGetResponseResultBindingsWorkersQueueBindingType = "queue"
)

type WorkerScriptSettingGetResponseResultBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingGetResponseResultBindingsWorkersD1BindingType `json:"type,required"`
	JSON workerScriptSettingGetResponseResultBindingsWorkersD1BindingJSON `json:"-"`
}

// workerScriptSettingGetResponseResultBindingsWorkersD1BindingJSON contains the
// JSON metadata for the struct
// [WorkerScriptSettingGetResponseResultBindingsWorkersD1Binding]
type workerScriptSettingGetResponseResultBindingsWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseResultBindingsWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingGetResponseResultBindingsWorkersD1Binding) implementsWorkerScriptSettingGetResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingGetResponseResultBindingsWorkersD1BindingType string

const (
	WorkerScriptSettingGetResponseResultBindingsWorkersD1BindingTypeD1 WorkerScriptSettingGetResponseResultBindingsWorkersD1BindingType = "d1"
)

type WorkerScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound WorkerScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     workerScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingJSON     `json:"-"`
}

// workerScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBinding]
type workerScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBinding) implementsWorkerScriptSettingGetResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingType string

const (
	WorkerScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace WorkerScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type WorkerScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker WorkerScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   workerScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// workerScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingOutbound]
type workerScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Outbound worker
type WorkerScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                        `json:"service"`
	JSON    workerScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// workerScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingOutboundWorker]
type workerScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptSettingGetResponseResultBindingsWorkersMtlsCertBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingGetResponseResultBindingsWorkersMtlsCertBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                                 `json:"certificate_id"`
	JSON          workerScriptSettingGetResponseResultBindingsWorkersMtlsCertBindingJSON `json:"-"`
}

// workerScriptSettingGetResponseResultBindingsWorkersMtlsCertBindingJSON contains
// the JSON metadata for the struct
// [WorkerScriptSettingGetResponseResultBindingsWorkersMtlsCertBinding]
type workerScriptSettingGetResponseResultBindingsWorkersMtlsCertBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseResultBindingsWorkersMtlsCertBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingGetResponseResultBindingsWorkersMtlsCertBinding) implementsWorkerScriptSettingGetResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingGetResponseResultBindingsWorkersMtlsCertBindingType string

const (
	WorkerScriptSettingGetResponseResultBindingsWorkersMtlsCertBindingTypeMtlsCertificate WorkerScriptSettingGetResponseResultBindingsWorkersMtlsCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [WorkerScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrations] or
// [WorkerScriptSettingGetResponseResultMigrationsWorkersSteppedMigrations].
type WorkerScriptSettingGetResponseResultMigrations interface {
	implementsWorkerScriptSettingGetResponseResultMigrations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkerScriptSettingGetResponseResultMigrations)(nil)).Elem(), "")
}

// A single set of migrations to apply.
type WorkerScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses []WorkerScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkerScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               workerScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsJSON               `json:"-"`
}

// workerScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrations]
type workerScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrations) implementsWorkerScriptSettingGetResponseResultMigrations() {
}

type WorkerScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From string                                                                                    `json:"from"`
	To   string                                                                                    `json:"to"`
	JSON workerScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// workerScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsRenamedClass]
type workerScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       string                                                                                        `json:"from"`
	FromScript string                                                                                        `json:"from_script"`
	To         string                                                                                        `json:"to"`
	JSON       workerScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// workerScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsTransferredClass]
type workerScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptSettingGetResponseResultMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []WorkerScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStep `json:"steps"`
	JSON  workerScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsJSON   `json:"-"`
}

// workerScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingGetResponseResultMigrationsWorkersSteppedMigrations]
type workerScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseResultMigrationsWorkersSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingGetResponseResultMigrationsWorkersSteppedMigrations) implementsWorkerScriptSettingGetResponseResultMigrations() {
}

type WorkerScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []WorkerScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkerScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               workerScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepJSON                `json:"-"`
}

// workerScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStep]
type workerScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From string                                                                                      `json:"from"`
	To   string                                                                                      `json:"to"`
	JSON workerScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// workerScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepsRenamedClass]
type workerScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                                          `json:"from"`
	FromScript string                                                                                          `json:"from_script"`
	To         string                                                                                          `json:"to"`
	JSON       workerScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// workerScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepsTransferredClass]
type workerScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptSettingGetResponseResultPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode WorkerScriptSettingGetResponseResultPlacementMode `json:"mode"`
	JSON workerScriptSettingGetResponseResultPlacementJSON `json:"-"`
}

// workerScriptSettingGetResponseResultPlacementJSON contains the JSON metadata for
// the struct [WorkerScriptSettingGetResponseResultPlacement]
type workerScriptSettingGetResponseResultPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseResultPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type WorkerScriptSettingGetResponseResultPlacementMode string

const (
	WorkerScriptSettingGetResponseResultPlacementModeSmart WorkerScriptSettingGetResponseResultPlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type WorkerScriptSettingGetResponseResultTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                               `json:"namespace"`
	JSON      workerScriptSettingGetResponseResultTailConsumerJSON `json:"-"`
}

// workerScriptSettingGetResponseResultTailConsumerJSON contains the JSON metadata
// for the struct [WorkerScriptSettingGetResponseResultTailConsumer]
type workerScriptSettingGetResponseResultTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingGetResponseResultTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerScriptSettingGetResponseSuccess bool

const (
	WorkerScriptSettingGetResponseSuccessTrue WorkerScriptSettingGetResponseSuccess = true
)

type WorkerScriptSettingUpdateResponse struct {
	Errors   []WorkerScriptSettingUpdateResponseError   `json:"errors"`
	Messages []WorkerScriptSettingUpdateResponseMessage `json:"messages"`
	Result   WorkerScriptSettingUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success WorkerScriptSettingUpdateResponseSuccess `json:"success"`
	JSON    workerScriptSettingUpdateResponseJSON    `json:"-"`
}

// workerScriptSettingUpdateResponseJSON contains the JSON metadata for the struct
// [WorkerScriptSettingUpdateResponse]
type workerScriptSettingUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptSettingUpdateResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    workerScriptSettingUpdateResponseErrorJSON `json:"-"`
}

// workerScriptSettingUpdateResponseErrorJSON contains the JSON metadata for the
// struct [WorkerScriptSettingUpdateResponseError]
type workerScriptSettingUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptSettingUpdateResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    workerScriptSettingUpdateResponseMessageJSON `json:"-"`
}

// workerScriptSettingUpdateResponseMessageJSON contains the JSON metadata for the
// struct [WorkerScriptSettingUpdateResponseMessage]
type workerScriptSettingUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptSettingUpdateResponseResult struct {
	// List of bindings attached to this Worker
	Bindings []WorkerScriptSettingUpdateResponseResultBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations WorkerScriptSettingUpdateResponseResultMigrations `json:"migrations"`
	Placement  WorkerScriptSettingUpdateResponseResultPlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []WorkerScriptSettingUpdateResponseResultTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                      `json:"usage_model"`
	JSON       workerScriptSettingUpdateResponseResultJSON `json:"-"`
}

// workerScriptSettingUpdateResponseResultJSON contains the JSON metadata for the
// struct [WorkerScriptSettingUpdateResponseResult]
type workerScriptSettingUpdateResponseResultJSON struct {
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

func (r *WorkerScriptSettingUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [WorkerScriptSettingUpdateResponseResultBindingsWorkersKvNamespaceBinding],
// [WorkerScriptSettingUpdateResponseResultBindingsWorkersServiceBinding],
// [WorkerScriptSettingUpdateResponseResultBindingsWorkersDoBinding],
// [WorkerScriptSettingUpdateResponseResultBindingsWorkersR2Binding],
// [WorkerScriptSettingUpdateResponseResultBindingsWorkersQueueBinding],
// [WorkerScriptSettingUpdateResponseResultBindingsWorkersD1Binding],
// [WorkerScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBinding]
// or [WorkerScriptSettingUpdateResponseResultBindingsWorkersMtlsCertBinding].
type WorkerScriptSettingUpdateResponseResultBinding interface {
	implementsWorkerScriptSettingUpdateResponseResultBinding()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkerScriptSettingUpdateResponseResultBinding)(nil)).Elem(), "")
}

type WorkerScriptSettingUpdateResponseResultBindingsWorkersKvNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingUpdateResponseResultBindingsWorkersKvNamespaceBindingType `json:"type,required"`
	JSON workerScriptSettingUpdateResponseResultBindingsWorkersKvNamespaceBindingJSON `json:"-"`
}

// workerScriptSettingUpdateResponseResultBindingsWorkersKvNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingUpdateResponseResultBindingsWorkersKvNamespaceBinding]
type workerScriptSettingUpdateResponseResultBindingsWorkersKvNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseResultBindingsWorkersKvNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingUpdateResponseResultBindingsWorkersKvNamespaceBinding) implementsWorkerScriptSettingUpdateResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingUpdateResponseResultBindingsWorkersKvNamespaceBindingType string

const (
	WorkerScriptSettingUpdateResponseResultBindingsWorkersKvNamespaceBindingTypeKvNamespace WorkerScriptSettingUpdateResponseResultBindingsWorkersKvNamespaceBindingType = "kv_namespace"
)

type WorkerScriptSettingUpdateResponseResultBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingUpdateResponseResultBindingsWorkersServiceBindingType `json:"type,required"`
	JSON workerScriptSettingUpdateResponseResultBindingsWorkersServiceBindingJSON `json:"-"`
}

// workerScriptSettingUpdateResponseResultBindingsWorkersServiceBindingJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingUpdateResponseResultBindingsWorkersServiceBinding]
type workerScriptSettingUpdateResponseResultBindingsWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseResultBindingsWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingUpdateResponseResultBindingsWorkersServiceBinding) implementsWorkerScriptSettingUpdateResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingUpdateResponseResultBindingsWorkersServiceBindingType string

const (
	WorkerScriptSettingUpdateResponseResultBindingsWorkersServiceBindingTypeService WorkerScriptSettingUpdateResponseResultBindingsWorkersServiceBindingType = "service"
)

type WorkerScriptSettingUpdateResponseResultBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingUpdateResponseResultBindingsWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                              `json:"script_name"`
	JSON       workerScriptSettingUpdateResponseResultBindingsWorkersDoBindingJSON `json:"-"`
}

// workerScriptSettingUpdateResponseResultBindingsWorkersDoBindingJSON contains the
// JSON metadata for the struct
// [WorkerScriptSettingUpdateResponseResultBindingsWorkersDoBinding]
type workerScriptSettingUpdateResponseResultBindingsWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseResultBindingsWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingUpdateResponseResultBindingsWorkersDoBinding) implementsWorkerScriptSettingUpdateResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingUpdateResponseResultBindingsWorkersDoBindingType string

const (
	WorkerScriptSettingUpdateResponseResultBindingsWorkersDoBindingTypeDurableObjectNamespace WorkerScriptSettingUpdateResponseResultBindingsWorkersDoBindingType = "durable_object_namespace"
)

type WorkerScriptSettingUpdateResponseResultBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingUpdateResponseResultBindingsWorkersR2BindingType `json:"type,required"`
	JSON workerScriptSettingUpdateResponseResultBindingsWorkersR2BindingJSON `json:"-"`
}

// workerScriptSettingUpdateResponseResultBindingsWorkersR2BindingJSON contains the
// JSON metadata for the struct
// [WorkerScriptSettingUpdateResponseResultBindingsWorkersR2Binding]
type workerScriptSettingUpdateResponseResultBindingsWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseResultBindingsWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingUpdateResponseResultBindingsWorkersR2Binding) implementsWorkerScriptSettingUpdateResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingUpdateResponseResultBindingsWorkersR2BindingType string

const (
	WorkerScriptSettingUpdateResponseResultBindingsWorkersR2BindingTypeR2Bucket WorkerScriptSettingUpdateResponseResultBindingsWorkersR2BindingType = "r2_bucket"
)

type WorkerScriptSettingUpdateResponseResultBindingsWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingUpdateResponseResultBindingsWorkersQueueBindingType `json:"type,required"`
	JSON workerScriptSettingUpdateResponseResultBindingsWorkersQueueBindingJSON `json:"-"`
}

// workerScriptSettingUpdateResponseResultBindingsWorkersQueueBindingJSON contains
// the JSON metadata for the struct
// [WorkerScriptSettingUpdateResponseResultBindingsWorkersQueueBinding]
type workerScriptSettingUpdateResponseResultBindingsWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseResultBindingsWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingUpdateResponseResultBindingsWorkersQueueBinding) implementsWorkerScriptSettingUpdateResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingUpdateResponseResultBindingsWorkersQueueBindingType string

const (
	WorkerScriptSettingUpdateResponseResultBindingsWorkersQueueBindingTypeQueue WorkerScriptSettingUpdateResponseResultBindingsWorkersQueueBindingType = "queue"
)

type WorkerScriptSettingUpdateResponseResultBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingUpdateResponseResultBindingsWorkersD1BindingType `json:"type,required"`
	JSON workerScriptSettingUpdateResponseResultBindingsWorkersD1BindingJSON `json:"-"`
}

// workerScriptSettingUpdateResponseResultBindingsWorkersD1BindingJSON contains the
// JSON metadata for the struct
// [WorkerScriptSettingUpdateResponseResultBindingsWorkersD1Binding]
type workerScriptSettingUpdateResponseResultBindingsWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseResultBindingsWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingUpdateResponseResultBindingsWorkersD1Binding) implementsWorkerScriptSettingUpdateResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingUpdateResponseResultBindingsWorkersD1BindingType string

const (
	WorkerScriptSettingUpdateResponseResultBindingsWorkersD1BindingTypeD1 WorkerScriptSettingUpdateResponseResultBindingsWorkersD1BindingType = "d1"
)

type WorkerScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound WorkerScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     workerScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingJSON     `json:"-"`
}

// workerScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBinding]
type workerScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBinding) implementsWorkerScriptSettingUpdateResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingType string

const (
	WorkerScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace WorkerScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type WorkerScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker WorkerScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   workerScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// workerScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingOutbound]
type workerScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Outbound worker
type WorkerScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                           `json:"service"`
	JSON    workerScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// workerScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingOutboundWorker]
type workerScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptSettingUpdateResponseResultBindingsWorkersMtlsCertBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkerScriptSettingUpdateResponseResultBindingsWorkersMtlsCertBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                                    `json:"certificate_id"`
	JSON          workerScriptSettingUpdateResponseResultBindingsWorkersMtlsCertBindingJSON `json:"-"`
}

// workerScriptSettingUpdateResponseResultBindingsWorkersMtlsCertBindingJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingUpdateResponseResultBindingsWorkersMtlsCertBinding]
type workerScriptSettingUpdateResponseResultBindingsWorkersMtlsCertBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseResultBindingsWorkersMtlsCertBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingUpdateResponseResultBindingsWorkersMtlsCertBinding) implementsWorkerScriptSettingUpdateResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkerScriptSettingUpdateResponseResultBindingsWorkersMtlsCertBindingType string

const (
	WorkerScriptSettingUpdateResponseResultBindingsWorkersMtlsCertBindingTypeMtlsCertificate WorkerScriptSettingUpdateResponseResultBindingsWorkersMtlsCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [WorkerScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrations]
// or [WorkerScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrations].
type WorkerScriptSettingUpdateResponseResultMigrations interface {
	implementsWorkerScriptSettingUpdateResponseResultMigrations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkerScriptSettingUpdateResponseResultMigrations)(nil)).Elem(), "")
}

// A single set of migrations to apply.
type WorkerScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses []WorkerScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkerScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               workerScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsJSON               `json:"-"`
}

// workerScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrations]
type workerScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrations) implementsWorkerScriptSettingUpdateResponseResultMigrations() {
}

type WorkerScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From string                                                                                       `json:"from"`
	To   string                                                                                       `json:"to"`
	JSON workerScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// workerScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsRenamedClass]
type workerScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       string                                                                                           `json:"from"`
	FromScript string                                                                                           `json:"from_script"`
	To         string                                                                                           `json:"to"`
	JSON       workerScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// workerScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsTransferredClass]
type workerScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []WorkerScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStep `json:"steps"`
	JSON  workerScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsJSON   `json:"-"`
}

// workerScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrations]
type workerScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkerScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrations) implementsWorkerScriptSettingUpdateResponseResultMigrations() {
}

type WorkerScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []WorkerScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkerScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               workerScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepJSON                `json:"-"`
}

// workerScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStep]
type workerScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From string                                                                                         `json:"from"`
	To   string                                                                                         `json:"to"`
	JSON workerScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// workerScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepsRenamedClass]
type workerScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                                             `json:"from"`
	FromScript string                                                                                             `json:"from_script"`
	To         string                                                                                             `json:"to"`
	JSON       workerScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// workerScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkerScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepsTransferredClass]
type workerScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerScriptSettingUpdateResponseResultPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode WorkerScriptSettingUpdateResponseResultPlacementMode `json:"mode"`
	JSON workerScriptSettingUpdateResponseResultPlacementJSON `json:"-"`
}

// workerScriptSettingUpdateResponseResultPlacementJSON contains the JSON metadata
// for the struct [WorkerScriptSettingUpdateResponseResultPlacement]
type workerScriptSettingUpdateResponseResultPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseResultPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type WorkerScriptSettingUpdateResponseResultPlacementMode string

const (
	WorkerScriptSettingUpdateResponseResultPlacementModeSmart WorkerScriptSettingUpdateResponseResultPlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type WorkerScriptSettingUpdateResponseResultTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                  `json:"namespace"`
	JSON      workerScriptSettingUpdateResponseResultTailConsumerJSON `json:"-"`
}

// workerScriptSettingUpdateResponseResultTailConsumerJSON contains the JSON
// metadata for the struct [WorkerScriptSettingUpdateResponseResultTailConsumer]
type workerScriptSettingUpdateResponseResultTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerScriptSettingUpdateResponseResultTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerScriptSettingUpdateResponseSuccess bool

const (
	WorkerScriptSettingUpdateResponseSuccessTrue WorkerScriptSettingUpdateResponseSuccess = true
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
