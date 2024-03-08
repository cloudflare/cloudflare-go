// File generated from our OpenAPI spec by Stainless.

package workers_for_platforms

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// DispatchNamespaceScriptContentBindingService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewDispatchNamespaceScriptContentBindingService] method instead.
type DispatchNamespaceScriptContentBindingService struct {
	Options []option.RequestOption
}

// NewDispatchNamespaceScriptContentBindingService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewDispatchNamespaceScriptContentBindingService(opts ...option.RequestOption) (r *DispatchNamespaceScriptContentBindingService) {
	r = &DispatchNamespaceScriptContentBindingService{}
	r.Options = opts
	return
}

// Fetch script bindings from a script uploaded to a Workers for Platforms
// namespace.
func (r *DispatchNamespaceScriptContentBindingService) Get(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptContentBindingGetParams, opts ...option.RequestOption) (res *DispatchNamespaceScriptContentBindingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/bindings", query.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [workers_for_platforms.DispatchNamespaceScriptContentBindingGetResponseWorkersKVNamespaceBinding],
// [workers_for_platforms.DispatchNamespaceScriptContentBindingGetResponseWorkersServiceBinding],
// [workers_for_platforms.DispatchNamespaceScriptContentBindingGetResponseWorkersDoBinding],
// [workers_for_platforms.DispatchNamespaceScriptContentBindingGetResponseWorkersR2Binding],
// [workers_for_platforms.DispatchNamespaceScriptContentBindingGetResponseWorkersQueueBinding],
// [workers_for_platforms.DispatchNamespaceScriptContentBindingGetResponseWorkersD1Binding],
// [workers_for_platforms.DispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBinding]
// or
// [workers_for_platforms.DispatchNamespaceScriptContentBindingGetResponseWorkersMTLSCertBinding].
type DispatchNamespaceScriptContentBindingGetResponse interface {
	implementsWorkersForPlatformsDispatchNamespaceScriptContentBindingGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DispatchNamespaceScriptContentBindingGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptContentBindingGetResponseWorkersKVNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptContentBindingGetResponseWorkersServiceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptContentBindingGetResponseWorkersDoBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptContentBindingGetResponseWorkersR2Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptContentBindingGetResponseWorkersQueueBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptContentBindingGetResponseWorkersD1Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptContentBindingGetResponseWorkersMTLSCertBinding{}),
		},
	)
}

type DispatchNamespaceScriptContentBindingGetResponseWorkersKVNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptContentBindingGetResponseWorkersKVNamespaceBindingType `json:"type,required"`
	JSON dispatchNamespaceScriptContentBindingGetResponseWorkersKVNamespaceBindingJSON `json:"-"`
}

// dispatchNamespaceScriptContentBindingGetResponseWorkersKVNamespaceBindingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentBindingGetResponseWorkersKVNamespaceBinding]
type dispatchNamespaceScriptContentBindingGetResponseWorkersKVNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentBindingGetResponseWorkersKVNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentBindingGetResponseWorkersKVNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptContentBindingGetResponseWorkersKVNamespaceBinding) implementsWorkersForPlatformsDispatchNamespaceScriptContentBindingGetResponse() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentBindingGetResponseWorkersKVNamespaceBindingType string

const (
	DispatchNamespaceScriptContentBindingGetResponseWorkersKVNamespaceBindingTypeKVNamespace DispatchNamespaceScriptContentBindingGetResponseWorkersKVNamespaceBindingType = "kv_namespace"
)

type DispatchNamespaceScriptContentBindingGetResponseWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptContentBindingGetResponseWorkersServiceBindingType `json:"type,required"`
	JSON dispatchNamespaceScriptContentBindingGetResponseWorkersServiceBindingJSON `json:"-"`
}

// dispatchNamespaceScriptContentBindingGetResponseWorkersServiceBindingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentBindingGetResponseWorkersServiceBinding]
type dispatchNamespaceScriptContentBindingGetResponseWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentBindingGetResponseWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentBindingGetResponseWorkersServiceBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptContentBindingGetResponseWorkersServiceBinding) implementsWorkersForPlatformsDispatchNamespaceScriptContentBindingGetResponse() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentBindingGetResponseWorkersServiceBindingType string

const (
	DispatchNamespaceScriptContentBindingGetResponseWorkersServiceBindingTypeService DispatchNamespaceScriptContentBindingGetResponseWorkersServiceBindingType = "service"
)

type DispatchNamespaceScriptContentBindingGetResponseWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptContentBindingGetResponseWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                               `json:"script_name"`
	JSON       dispatchNamespaceScriptContentBindingGetResponseWorkersDoBindingJSON `json:"-"`
}

// dispatchNamespaceScriptContentBindingGetResponseWorkersDoBindingJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptContentBindingGetResponseWorkersDoBinding]
type dispatchNamespaceScriptContentBindingGetResponseWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentBindingGetResponseWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentBindingGetResponseWorkersDoBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptContentBindingGetResponseWorkersDoBinding) implementsWorkersForPlatformsDispatchNamespaceScriptContentBindingGetResponse() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentBindingGetResponseWorkersDoBindingType string

const (
	DispatchNamespaceScriptContentBindingGetResponseWorkersDoBindingTypeDurableObjectNamespace DispatchNamespaceScriptContentBindingGetResponseWorkersDoBindingType = "durable_object_namespace"
)

type DispatchNamespaceScriptContentBindingGetResponseWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptContentBindingGetResponseWorkersR2BindingType `json:"type,required"`
	JSON dispatchNamespaceScriptContentBindingGetResponseWorkersR2BindingJSON `json:"-"`
}

// dispatchNamespaceScriptContentBindingGetResponseWorkersR2BindingJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptContentBindingGetResponseWorkersR2Binding]
type dispatchNamespaceScriptContentBindingGetResponseWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentBindingGetResponseWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentBindingGetResponseWorkersR2BindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptContentBindingGetResponseWorkersR2Binding) implementsWorkersForPlatformsDispatchNamespaceScriptContentBindingGetResponse() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentBindingGetResponseWorkersR2BindingType string

const (
	DispatchNamespaceScriptContentBindingGetResponseWorkersR2BindingTypeR2Bucket DispatchNamespaceScriptContentBindingGetResponseWorkersR2BindingType = "r2_bucket"
)

type DispatchNamespaceScriptContentBindingGetResponseWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptContentBindingGetResponseWorkersQueueBindingType `json:"type,required"`
	JSON dispatchNamespaceScriptContentBindingGetResponseWorkersQueueBindingJSON `json:"-"`
}

// dispatchNamespaceScriptContentBindingGetResponseWorkersQueueBindingJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptContentBindingGetResponseWorkersQueueBinding]
type dispatchNamespaceScriptContentBindingGetResponseWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentBindingGetResponseWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentBindingGetResponseWorkersQueueBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptContentBindingGetResponseWorkersQueueBinding) implementsWorkersForPlatformsDispatchNamespaceScriptContentBindingGetResponse() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentBindingGetResponseWorkersQueueBindingType string

const (
	DispatchNamespaceScriptContentBindingGetResponseWorkersQueueBindingTypeQueue DispatchNamespaceScriptContentBindingGetResponseWorkersQueueBindingType = "queue"
)

type DispatchNamespaceScriptContentBindingGetResponseWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptContentBindingGetResponseWorkersD1BindingType `json:"type,required"`
	JSON dispatchNamespaceScriptContentBindingGetResponseWorkersD1BindingJSON `json:"-"`
}

// dispatchNamespaceScriptContentBindingGetResponseWorkersD1BindingJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptContentBindingGetResponseWorkersD1Binding]
type dispatchNamespaceScriptContentBindingGetResponseWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentBindingGetResponseWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentBindingGetResponseWorkersD1BindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptContentBindingGetResponseWorkersD1Binding) implementsWorkersForPlatformsDispatchNamespaceScriptContentBindingGetResponse() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentBindingGetResponseWorkersD1BindingType string

const (
	DispatchNamespaceScriptContentBindingGetResponseWorkersD1BindingTypeD1 DispatchNamespaceScriptContentBindingGetResponseWorkersD1BindingType = "d1"
)

type DispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound DispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     dispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingJSON     `json:"-"`
}

// dispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBinding]
type dispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBinding) implementsWorkersForPlatformsDispatchNamespaceScriptContentBindingGetResponse() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingType string

const (
	DispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingTypeDispatchNamespace DispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type DispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker DispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   dispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// dispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingOutbound]
type dispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingOutboundJSON) RawJSON() string {
	return r.raw
}

// Outbound worker
type DispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                            `json:"service"`
	JSON    dispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// dispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingOutboundWorker]
type dispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentBindingGetResponseWorkersDispatchNamespaceBindingOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptContentBindingGetResponseWorkersMTLSCertBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptContentBindingGetResponseWorkersMTLSCertBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                                     `json:"certificate_id"`
	JSON          dispatchNamespaceScriptContentBindingGetResponseWorkersMTLSCertBindingJSON `json:"-"`
}

// dispatchNamespaceScriptContentBindingGetResponseWorkersMTLSCertBindingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptContentBindingGetResponseWorkersMTLSCertBinding]
type dispatchNamespaceScriptContentBindingGetResponseWorkersMTLSCertBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DispatchNamespaceScriptContentBindingGetResponseWorkersMTLSCertBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptContentBindingGetResponseWorkersMTLSCertBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptContentBindingGetResponseWorkersMTLSCertBinding) implementsWorkersForPlatformsDispatchNamespaceScriptContentBindingGetResponse() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptContentBindingGetResponseWorkersMTLSCertBindingType string

const (
	DispatchNamespaceScriptContentBindingGetResponseWorkersMTLSCertBindingTypeMTLSCertificate DispatchNamespaceScriptContentBindingGetResponseWorkersMTLSCertBindingType = "mtls_certificate"
)

type DispatchNamespaceScriptContentBindingGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
