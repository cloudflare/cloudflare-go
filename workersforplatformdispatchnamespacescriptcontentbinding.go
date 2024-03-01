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

// WorkersForPlatformDispatchNamespaceScriptContentBindingService contains methods
// and other services that help with interacting with the cloudflare API. Note,
// unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkersForPlatformDispatchNamespaceScriptContentBindingService] method
// instead.
type WorkersForPlatformDispatchNamespaceScriptContentBindingService struct {
	Options []option.RequestOption
}

// NewWorkersForPlatformDispatchNamespaceScriptContentBindingService generates a
// new service that applies the given options to each request. These options are
// applied after the parent client's options (if there is one), and before any
// request-specific options.
func NewWorkersForPlatformDispatchNamespaceScriptContentBindingService(opts ...option.RequestOption) (r *WorkersForPlatformDispatchNamespaceScriptContentBindingService) {
	r = &WorkersForPlatformDispatchNamespaceScriptContentBindingService{}
	r.Options = opts
	return
}

// Fetch script bindings from a script uploaded to a Workers for Platforms
// namespace.
func (r *WorkersForPlatformDispatchNamespaceScriptContentBindingService) Get(ctx context.Context, dispatchNamespace string, scriptName string, query WorkersForPlatformDispatchNamespaceScriptContentBindingGetParams, opts ...option.RequestOption) (res *WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/bindings", query.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersKVNamespaceBinding],
// [WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersServiceBinding],
// [WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDoBinding],
// [WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersR2Binding],
// [WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersQueueBinding],
// [WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersD1Binding],
// [WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBinding]
// or
// [WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersMTLSCertBinding].
type WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponse interface {
	implementsWorkersForPlatformDispatchNamespaceScriptContentBindingGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponse)(nil)).Elem(), "")
}

type WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersKVNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersKVNamespaceBindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersKVNamespaceBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersKVNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersKVNamespaceBinding]
type workersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersKVNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersKVNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersKVNamespaceBinding) implementsWorkersForPlatformDispatchNamespaceScriptContentBindingGetResponse() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersKVNamespaceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersKVNamespaceBindingTypeKVNamespace WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersKVNamespaceBindingType = "kv_namespace"
)

type WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersServiceBindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersServiceBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersServiceBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersServiceBinding]
type workersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersServiceBinding) implementsWorkersForPlatformDispatchNamespaceScriptContentBindingGetResponse() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersServiceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersServiceBindingTypeService WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersServiceBindingType = "service"
)

type WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                                                 `json:"script_name"`
	JSON       workersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDoBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDoBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDoBinding]
type workersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDoBinding) implementsWorkersForPlatformDispatchNamespaceScriptContentBindingGetResponse() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDoBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDoBindingTypeDurableObjectNamespace WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDoBindingType = "durable_object_namespace"
)

type WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersR2BindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersR2BindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersR2BindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersR2Binding]
type workersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersR2Binding) implementsWorkersForPlatformDispatchNamespaceScriptContentBindingGetResponse() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersR2BindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersR2BindingTypeR2Bucket WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersR2BindingType = "r2_bucket"
)

type WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersQueueBindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersQueueBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersQueueBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersQueueBinding]
type workersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersQueueBinding) implementsWorkersForPlatformDispatchNamespaceScriptContentBindingGetResponse() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersQueueBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersQueueBindingTypeQueue WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersQueueBindingType = "queue"
)

type WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersD1BindingType `json:"type,required"`
	JSON workersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersD1BindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersD1BindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersD1Binding]
type workersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersD1Binding) implementsWorkersForPlatformDispatchNamespaceScriptContentBindingGetResponse() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersD1BindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersD1BindingTypeD1 WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersD1BindingType = "d1"
)

type WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     workersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingJSON     `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBinding]
type workersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBinding) implementsWorkersForPlatformDispatchNamespaceScriptContentBindingGetResponse() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingTypeDispatchNamespace WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   workersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingOutbound]
type workersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Outbound worker
type WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                                              `json:"service"`
	JSON    workersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingOutboundWorker]
type workersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersMTLSCertBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersMTLSCertBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                                                       `json:"certificate_id"`
	JSON          workersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersMTLSCertBindingJSON `json:"-"`
}

// workersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersMTLSCertBindingJSON
// contains the JSON metadata for the struct
// [WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersMTLSCertBinding]
type workersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersMTLSCertBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersMTLSCertBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersMTLSCertBinding) implementsWorkersForPlatformDispatchNamespaceScriptContentBindingGetResponse() {
}

// The class of resource that the binding provides.
type WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersMTLSCertBindingType string

const (
	WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersMTLSCertBindingTypeMTLSCertificate WorkersForPlatformDispatchNamespaceScriptContentBindingGetResponseWorkersMTLSCertBindingType = "mtls_certificate"
)

type WorkersForPlatformDispatchNamespaceScriptContentBindingGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
