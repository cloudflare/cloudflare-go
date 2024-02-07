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

// Get script settings from a script uploaded to a Workers for Platforms namespace.
func (r *WorkersForPlatformDispatchNamespaceScriptSettingService) Get(ctx context.Context, accountID string, dispatchNamespace string, scriptName string, opts ...option.RequestOption) (res *WorkersForPlatformDispatchNamespaceScriptSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/settings", accountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Patch script metadata, such as bindings
func (r *WorkersForPlatformDispatchNamespaceScriptSettingService) Update(ctx context.Context, accountID string, dispatchNamespace string, scriptName string, body WorkersForPlatformDispatchNamespaceScriptSettingUpdateParams, opts ...option.RequestOption) (res *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/settings", accountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponse struct {
	Errors   []WorkersForPlatformDispatchNamespaceScriptSettingGetResponseError   `json:"errors"`
	Messages []WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMessage `json:"messages"`
	Result   WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success WorkersForPlatformDispatchNamespaceScriptSettingGetResponseSuccess `json:"success"`
	JSON    workersForPlatformDispatchNamespaceScriptSettingGetResponseJSON    `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseJSON contains the
// JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponse]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseError struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    workersForPlatformDispatchNamespaceScriptSettingGetResponseErrorJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseErrorJSON contains
// the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseError]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMessage struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    workersForPlatformDispatchNamespaceScriptSettingGetResponseMessageJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseMessageJSON contains
// the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMessage]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResult struct {
	// List of bindings attached to this Worker
	Bindings []WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrations `json:"migrations"`
	Placement  WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultPlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                                                `json:"usage_model"`
	JSON       workersForPlatformDispatchNamespaceScriptSettingGetResponseResultJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseResultJSON contains
// the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResult]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseResultJSON struct {
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

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersKvNamespaceBinding],
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersServiceBinding],
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDoBinding],
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersR2Binding],
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersQueueBinding],
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersD1Binding],
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBinding]
// or
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersMtlsCertBinding].
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBinding interface {
	implementsWorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBinding()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBinding)(nil)).Elem(), "")
}

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersKvNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersKvNamespaceBindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersKvNamespaceBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersKvNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersKvNamespaceBinding]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersKvNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersKvNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersKvNamespaceBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersKvNamespaceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersKvNamespaceBindingTypeKvNamespace WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersKvNamespaceBindingType = "kv_namespace"
)

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersServiceBindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersServiceBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersServiceBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersServiceBinding]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersServiceBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersServiceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersServiceBindingTypeService WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersServiceBindingType = "service"
)

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                                                        `json:"script_name"`
	JSON       workersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDoBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDoBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDoBinding]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDoBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDoBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDoBindingTypeDurableObjectNamespace WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDoBindingType = "durable_object_namespace"
)

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersR2BindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersR2BindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersR2BindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersR2Binding]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersR2Binding) implementsWorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersR2BindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersR2BindingTypeR2Bucket WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersR2BindingType = "r2_bucket"
)

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersQueueBindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersQueueBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersQueueBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersQueueBinding]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersQueueBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersQueueBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersQueueBindingTypeQueue WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersQueueBindingType = "queue"
)

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersD1BindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersD1BindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersD1BindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersD1Binding]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersD1Binding) implementsWorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersD1BindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersD1BindingTypeD1 WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersD1BindingType = "d1"
)

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     workersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingJSON     `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBinding]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   workersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingOutbound]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Outbound worker
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                                                     `json:"service"`
	JSON    workersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingOutboundWorker]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersMtlsCertBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersMtlsCertBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                                                              `json:"certificate_id"`
	JSON          workersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersMtlsCertBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersMtlsCertBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersMtlsCertBinding]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersMtlsCertBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersMtlsCertBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersMtlsCertBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersMtlsCertBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersMtlsCertBindingTypeMtlsCertificate WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultBindingsWorkersMtlsCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrations]
// or
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSteppedMigrations].
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrations interface {
	implementsWorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrations)(nil)).Elem(), "")
}

// A single set of migrations to apply.
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses []WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               workersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsJSON               `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrations]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrations) implementsWorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrations() {
}

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From string                                                                                                                 `json:"from"`
	To   string                                                                                                                 `json:"to"`
	JSON workersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsRenamedClass]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       string                                                                                                                     `json:"from"`
	FromScript string                                                                                                                     `json:"from_script"`
	To         string                                                                                                                     `json:"to"`
	JSON       workersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsTransferredClass]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStep `json:"steps"`
	JSON  workersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsJSON   `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSteppedMigrations]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSteppedMigrations) implementsWorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrations() {
}

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               workersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepJSON                `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStep]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From string                                                                                                                   `json:"from"`
	To   string                                                                                                                   `json:"to"`
	JSON workersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepsRenamedClass]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                                                                       `json:"from"`
	FromScript string                                                                                                                       `json:"from_script"`
	To         string                                                                                                                       `json:"to"`
	JSON       workersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepsTransferredClass]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultMigrationsWorkersSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultPlacementMode `json:"mode"`
	JSON workersForPlatformDispatchNamespaceScriptSettingGetResponseResultPlacementJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseResultPlacementJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultPlacement]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseResultPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultPlacementMode string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultPlacementModeSmart WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultPlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                                            `json:"namespace"`
	JSON      workersForPlatformDispatchNamespaceScriptSettingGetResponseResultTailConsumerJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingGetResponseResultTailConsumerJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultTailConsumer]
type workersForPlatformDispatchNamespaceScriptSettingGetResponseResultTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingGetResponseResultTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkersForPlatformDispatchNamespaceScriptSettingGetResponseSuccess bool

const (
	WorkersForPlatformDispatchNamespaceScriptSettingGetResponseSuccessTrue WorkersForPlatformDispatchNamespaceScriptSettingGetResponseSuccess = true
)

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponse struct {
	Errors   []WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseError   `json:"errors"`
	Messages []WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMessage `json:"messages"`
	Result   WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseSuccess `json:"success"`
	JSON    workersForPlatformDispatchNamespaceScriptSettingUpdateResponseJSON    `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseJSON contains the
// JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponse]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseError struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    workersForPlatformDispatchNamespaceScriptSettingUpdateResponseErrorJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseErrorJSON contains
// the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseError]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMessage struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    workersForPlatformDispatchNamespaceScriptSettingUpdateResponseMessageJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseMessageJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMessage]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResult struct {
	// List of bindings attached to this Worker
	Bindings []WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrations `json:"migrations"`
	Placement  WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultPlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                                                   `json:"usage_model"`
	JSON       workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResult]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultJSON struct {
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

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersKvNamespaceBinding],
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersServiceBinding],
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDoBinding],
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersR2Binding],
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersQueueBinding],
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersD1Binding],
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBinding]
// or
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersMtlsCertBinding].
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBinding interface {
	implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBinding()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBinding)(nil)).Elem(), "")
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersKvNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersKvNamespaceBindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersKvNamespaceBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersKvNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersKvNamespaceBinding]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersKvNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersKvNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersKvNamespaceBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersKvNamespaceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersKvNamespaceBindingTypeKvNamespace WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersKvNamespaceBindingType = "kv_namespace"
)

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersServiceBindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersServiceBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersServiceBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersServiceBinding]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersServiceBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersServiceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersServiceBindingTypeService WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersServiceBindingType = "service"
)

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                                                           `json:"script_name"`
	JSON       workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDoBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDoBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDoBinding]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDoBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDoBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDoBindingTypeDurableObjectNamespace WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDoBindingType = "durable_object_namespace"
)

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersR2BindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersR2BindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersR2BindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersR2Binding]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersR2Binding) implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersR2BindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersR2BindingTypeR2Bucket WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersR2BindingType = "r2_bucket"
)

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersQueueBindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersQueueBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersQueueBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersQueueBinding]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersQueueBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersQueueBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersQueueBindingTypeQueue WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersQueueBindingType = "queue"
)

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersD1BindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersD1BindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersD1BindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersD1Binding]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersD1Binding) implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersD1BindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersD1BindingTypeD1 WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersD1BindingType = "d1"
)

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingJSON     `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBinding]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingOutbound]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Outbound worker
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                                                        `json:"service"`
	JSON    workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingOutboundWorker]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersMtlsCertBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersMtlsCertBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                                                                 `json:"certificate_id"`
	JSON          workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersMtlsCertBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersMtlsCertBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersMtlsCertBinding]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersMtlsCertBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersMtlsCertBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersMtlsCertBinding) implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBinding() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersMtlsCertBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersMtlsCertBindingTypeMtlsCertificate WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultBindingsWorkersMtlsCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrations]
// or
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrations].
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrations interface {
	implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrations)(nil)).Elem(), "")
}

// A single set of migrations to apply.
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses []WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsJSON               `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrations]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrations) implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrations() {
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From string                                                                                                                    `json:"from"`
	To   string                                                                                                                    `json:"to"`
	JSON workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsRenamedClass]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       string                                                                                                                        `json:"from"`
	FromScript string                                                                                                                        `json:"from_script"`
	To         string                                                                                                                        `json:"to"`
	JSON       workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsTransferredClass]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStep `json:"steps"`
	JSON  workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsJSON   `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrations]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrations) implementsWorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrations() {
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepJSON                `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStep]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From string                                                                                                                      `json:"from"`
	To   string                                                                                                                      `json:"to"`
	JSON workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepsRenamedClass]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                                                                          `json:"from"`
	FromScript string                                                                                                                          `json:"from_script"`
	To         string                                                                                                                          `json:"to"`
	JSON       workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepsTransferredClass]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultMigrationsWorkersSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultPlacementMode `json:"mode"`
	JSON workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultPlacementJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultPlacementJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultPlacement]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultPlacementMode string

const (
	WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultPlacementModeSmart WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultPlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                                               `json:"namespace"`
	JSON      workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultTailConsumerJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultTailConsumerJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultTailConsumer]
type workersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseResultTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseSuccess bool

const (
	WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseSuccessTrue WorkersForPlatformDispatchNamespaceScriptSettingUpdateResponseSuccess = true
)

type WorkersForPlatformDispatchNamespaceScriptSettingUpdateParams struct {
	Errors   param.Field[[]WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsError]   `json:"errors"`
	Messages param.Field[[]WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsMessage] `json:"messages"`
	Result   param.Field[WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsResult]    `json:"result"`
	// Whether the API call was successful
	Success param.Field[WorkersForPlatformDispatchNamespaceScriptSettingUpdateParamsSuccess] `json:"success"`
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
