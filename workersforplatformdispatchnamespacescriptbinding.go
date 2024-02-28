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

// WorkersForPlatformDispatchNamespaceScriptBindingService contains methods and
// other services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkersForPlatformDispatchNamespaceScriptBindingService] method instead.
type WorkersForPlatformDispatchNamespaceScriptBindingService struct {
	Options []option.RequestOption
}

// NewWorkersForPlatformDispatchNamespaceScriptBindingService generates a new
// service that applies the given options to each request. These options are
// applied after the parent client's options (if there is one), and before any
// request-specific options.
func NewWorkersForPlatformDispatchNamespaceScriptBindingService(opts ...option.RequestOption) (r *WorkersForPlatformDispatchNamespaceScriptBindingService) {
	r = &WorkersForPlatformDispatchNamespaceScriptBindingService{}
	r.Options = opts
	return
}

// Fetch script bindings from a script uploaded to a Workers for Platforms
// namespace.
func (r *WorkersForPlatformDispatchNamespaceScriptBindingService) Get(ctx context.Context, dispatchNamespace string, scriptName string, query WorkersForPlatformDispatchNamespaceScriptBindingGetParams, opts ...option.RequestOption) (res *WorkersForPlatformDispatchNamespaceScriptBindingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/bindings", query.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersKVNamespaceBinding],
// [WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersServiceBinding],
// [WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDoBinding],
// [WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersR2Binding],
// [WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersQueueBinding],
// [WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersD1Binding],
// [WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBinding]
// or
// [WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersMTLSCertBinding].
type WorkersForPlatformDispatchNamespaceScriptBindingGetResponse interface {
	implementsWorkersForPlatformDispatchNamespaceScriptBindingGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkersForPlatformDispatchNamespaceScriptBindingGetResponse)(nil)).Elem(), "")
}

type WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersKVNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersKVNamespaceBindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersKVNamespaceBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersKVNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersKVNamespaceBinding]
type workersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersKVNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersKVNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersKVNamespaceBinding) implementsWorkersForPlatformDispatchNamespaceScriptBindingGetResponse() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersKVNamespaceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersKVNamespaceBindingTypeKVNamespace WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersKVNamespaceBindingType = "kv_namespace"
)

type WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersServiceBindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersServiceBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersServiceBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersServiceBinding]
type workersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersServiceBinding) implementsWorkersForPlatformDispatchNamespaceScriptBindingGetResponse() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersServiceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersServiceBindingTypeService WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersServiceBindingType = "service"
)

type WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                                          `json:"script_name"`
	JSON       workersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDoBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDoBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDoBinding]
type workersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDoBinding) implementsWorkersForPlatformDispatchNamespaceScriptBindingGetResponse() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDoBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDoBindingTypeDurableObjectNamespace WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDoBindingType = "durable_object_namespace"
)

type WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersR2BindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersR2BindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersR2BindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersR2Binding]
type workersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersR2Binding) implementsWorkersForPlatformDispatchNamespaceScriptBindingGetResponse() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersR2BindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersR2BindingTypeR2Bucket WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersR2BindingType = "r2_bucket"
)

type WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersQueueBindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersQueueBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersQueueBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersQueueBinding]
type workersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersQueueBinding) implementsWorkersForPlatformDispatchNamespaceScriptBindingGetResponse() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersQueueBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersQueueBindingTypeQueue WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersQueueBindingType = "queue"
)

type WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersD1BindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersD1BindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersD1BindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersD1Binding]
type workersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersD1Binding) implementsWorkersForPlatformDispatchNamespaceScriptBindingGetResponse() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersD1BindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersD1BindingTypeD1 WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersD1BindingType = "d1"
)

type WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     workersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingJSON     `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBinding]
type workersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBinding) implementsWorkersForPlatformDispatchNamespaceScriptBindingGetResponse() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingTypeDispatchNamespace WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   workersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingOutbound]
type workersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Outbound worker
type WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                                       `json:"service"`
	JSON    workersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingOutboundWorker]
type workersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersMTLSCertBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersMTLSCertBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                                                `json:"certificate_id"`
	JSON          workersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersMTLSCertBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersMTLSCertBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersMTLSCertBinding]
type workersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersMTLSCertBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersMTLSCertBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersMTLSCertBinding) implementsWorkersForPlatformDispatchNamespaceScriptBindingGetResponse() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersMTLSCertBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersMTLSCertBindingTypeMTLSCertificate WorkersForPlatformDispatchNamespaceScriptBindingGetResponseWorkersMTLSCertBindingType = "mtls_certificate"
)

type WorkersForPlatformDispatchNamespaceScriptBindingGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
