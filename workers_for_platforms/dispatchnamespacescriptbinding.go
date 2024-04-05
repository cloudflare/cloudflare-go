// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

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

// DispatchNamespaceScriptBindingService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewDispatchNamespaceScriptBindingService] method instead.
type DispatchNamespaceScriptBindingService struct {
	Options []option.RequestOption
}

// NewDispatchNamespaceScriptBindingService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDispatchNamespaceScriptBindingService(opts ...option.RequestOption) (r *DispatchNamespaceScriptBindingService) {
	r = &DispatchNamespaceScriptBindingService{}
	r.Options = opts
	return
}

// Fetch script bindings from a script uploaded to a Workers for Platforms
// namespace.
func (r *DispatchNamespaceScriptBindingService) Get(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptBindingGetParams, opts ...option.RequestOption) (res *DispatchNamespaceScriptBindingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/bindings", query.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// A binding to allow the Worker to communicate with resources
type DispatchNamespaceScriptBindingGetResponse struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseType `json:"type,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Name of Worker to bind to
	Service string `json:"service"`
	// The exported class name of the Durable Object
	ClassName string `json:"class_name"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string `json:"script_name"`
	// R2 bucket to bind to
	BucketName string `json:"bucket_name"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding"`
	// ID of the D1 database to bind to
	ID string `json:"id"`
	// Namespace to bind to
	Namespace string      `json:"namespace"`
	Outbound  interface{} `json:"outbound,required"`
	// ID of the certificate to bind to
	CertificateID string                                        `json:"certificate_id"`
	Certificate   interface{}                                   `json:"certificate,required"`
	JSON          dispatchNamespaceScriptBindingGetResponseJSON `json:"-"`
	union         DispatchNamespaceScriptBindingGetResponseUnion
}

// dispatchNamespaceScriptBindingGetResponseJSON contains the JSON metadata for the
// struct [DispatchNamespaceScriptBindingGetResponse]
type dispatchNamespaceScriptBindingGetResponseJSON struct {
	Name          apijson.Field
	NamespaceID   apijson.Field
	Type          apijson.Field
	Environment   apijson.Field
	Service       apijson.Field
	ClassName     apijson.Field
	ScriptName    apijson.Field
	BucketName    apijson.Field
	QueueName     apijson.Field
	Binding       apijson.Field
	ID            apijson.Field
	Namespace     apijson.Field
	Outbound      apijson.Field
	CertificateID apijson.Field
	Certificate   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r dispatchNamespaceScriptBindingGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DispatchNamespaceScriptBindingGetResponse) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r DispatchNamespaceScriptBindingGetResponse) AsUnion() DispatchNamespaceScriptBindingGetResponseUnion {
	return r.union
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersKVNamespaceBinding],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersServiceBinding],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersDoBinding],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersR2Binding],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersQueueBinding],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersD1Binding],
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBinding]
// or
// [workers_for_platforms.DispatchNamespaceScriptBindingGetResponseWorkersMTLSCERTBinding].
type DispatchNamespaceScriptBindingGetResponseUnion interface {
	implementsWorkersForPlatformsDispatchNamespaceScriptBindingGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DispatchNamespaceScriptBindingGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersKVNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersServiceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersDoBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersR2Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersQueueBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersD1Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptBindingGetResponseWorkersMTLSCERTBinding{}),
		},
	)
}

type DispatchNamespaceScriptBindingGetResponseWorkersKVNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersKVNamespaceBindingType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersKVNamespaceBindingJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersKVNamespaceBindingJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersKVNamespaceBinding]
type dispatchNamespaceScriptBindingGetResponseWorkersKVNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersKVNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersKVNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersKVNamespaceBinding) implementsWorkersForPlatformsDispatchNamespaceScriptBindingGetResponse() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersKVNamespaceBindingType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersKVNamespaceBindingTypeKVNamespace DispatchNamespaceScriptBindingGetResponseWorkersKVNamespaceBindingType = "kv_namespace"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersKVNamespaceBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersKVNamespaceBindingTypeKVNamespace:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersServiceBindingType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersServiceBindingJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersServiceBindingJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersServiceBinding]
type dispatchNamespaceScriptBindingGetResponseWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersServiceBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersServiceBinding) implementsWorkersForPlatformsDispatchNamespaceScriptBindingGetResponse() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersServiceBindingType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersServiceBindingTypeService DispatchNamespaceScriptBindingGetResponseWorkersServiceBindingType = "service"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersServiceBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersServiceBindingTypeService:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                        `json:"script_name"`
	JSON       dispatchNamespaceScriptBindingGetResponseWorkersDoBindingJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersDoBindingJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersDoBinding]
type dispatchNamespaceScriptBindingGetResponseWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersDoBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersDoBinding) implementsWorkersForPlatformsDispatchNamespaceScriptBindingGetResponse() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersDoBindingType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersDoBindingTypeDurableObjectNamespace DispatchNamespaceScriptBindingGetResponseWorkersDoBindingType = "durable_object_namespace"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersDoBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersDoBindingTypeDurableObjectNamespace:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersR2BindingType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersR2BindingJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersR2BindingJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersR2Binding]
type dispatchNamespaceScriptBindingGetResponseWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersR2BindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersR2Binding) implementsWorkersForPlatformsDispatchNamespaceScriptBindingGetResponse() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersR2BindingType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersR2BindingTypeR2Bucket DispatchNamespaceScriptBindingGetResponseWorkersR2BindingType = "r2_bucket"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersR2BindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersR2BindingTypeR2Bucket:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersQueueBindingType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersQueueBindingJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersQueueBindingJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersQueueBinding]
type dispatchNamespaceScriptBindingGetResponseWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersQueueBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersQueueBinding) implementsWorkersForPlatformsDispatchNamespaceScriptBindingGetResponse() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersQueueBindingType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersQueueBindingTypeQueue DispatchNamespaceScriptBindingGetResponseWorkersQueueBindingType = "queue"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersQueueBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersQueueBindingTypeQueue:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersD1BindingType `json:"type,required"`
	JSON dispatchNamespaceScriptBindingGetResponseWorkersD1BindingJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersD1BindingJSON contains the JSON
// metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersD1Binding]
type dispatchNamespaceScriptBindingGetResponseWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersD1BindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersD1Binding) implementsWorkersForPlatformsDispatchNamespaceScriptBindingGetResponse() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersD1BindingType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersD1BindingTypeD1 DispatchNamespaceScriptBindingGetResponseWorkersD1BindingType = "d1"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersD1BindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersD1BindingTypeD1:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound DispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     dispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingJSON     `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBinding]
type dispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBinding) implementsWorkersForPlatformsDispatchNamespaceScriptBindingGetResponse() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingTypeDispatchNamespace DispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker
type DispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker DispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   dispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingOutbound]
type dispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingOutboundJSON) RawJSON() string {
	return r.raw
}

// Outbound worker
type DispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                     `json:"service"`
	JSON    dispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingOutboundWorker]
type dispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersDispatchNamespaceBindingOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptBindingGetResponseWorkersMTLSCERTBinding struct {
	Certificate interface{} `json:"certificate,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceScriptBindingGetResponseWorkersMTLSCERTBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                              `json:"certificate_id"`
	JSON          dispatchNamespaceScriptBindingGetResponseWorkersMtlscertBindingJSON `json:"-"`
}

// dispatchNamespaceScriptBindingGetResponseWorkersMtlscertBindingJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptBindingGetResponseWorkersMTLSCERTBinding]
type dispatchNamespaceScriptBindingGetResponseWorkersMtlscertBindingJSON struct {
	Certificate   apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DispatchNamespaceScriptBindingGetResponseWorkersMTLSCERTBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptBindingGetResponseWorkersMtlscertBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptBindingGetResponseWorkersMTLSCERTBinding) implementsWorkersForPlatformsDispatchNamespaceScriptBindingGetResponse() {
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseWorkersMTLSCERTBindingType string

const (
	DispatchNamespaceScriptBindingGetResponseWorkersMTLSCERTBindingTypeMTLSCertificate DispatchNamespaceScriptBindingGetResponseWorkersMTLSCERTBindingType = "mtls_certificate"
)

func (r DispatchNamespaceScriptBindingGetResponseWorkersMTLSCERTBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseWorkersMTLSCERTBindingTypeMTLSCertificate:
		return true
	}
	return false
}

// The class of resource that the binding provides.
type DispatchNamespaceScriptBindingGetResponseType string

const (
	DispatchNamespaceScriptBindingGetResponseTypeKVNamespace            DispatchNamespaceScriptBindingGetResponseType = "kv_namespace"
	DispatchNamespaceScriptBindingGetResponseTypeService                DispatchNamespaceScriptBindingGetResponseType = "service"
	DispatchNamespaceScriptBindingGetResponseTypeDurableObjectNamespace DispatchNamespaceScriptBindingGetResponseType = "durable_object_namespace"
	DispatchNamespaceScriptBindingGetResponseTypeR2Bucket               DispatchNamespaceScriptBindingGetResponseType = "r2_bucket"
	DispatchNamespaceScriptBindingGetResponseTypeQueue                  DispatchNamespaceScriptBindingGetResponseType = "queue"
	DispatchNamespaceScriptBindingGetResponseTypeD1                     DispatchNamespaceScriptBindingGetResponseType = "d1"
	DispatchNamespaceScriptBindingGetResponseTypeDispatchNamespace      DispatchNamespaceScriptBindingGetResponseType = "dispatch_namespace"
	DispatchNamespaceScriptBindingGetResponseTypeMTLSCertificate        DispatchNamespaceScriptBindingGetResponseType = "mtls_certificate"
)

func (r DispatchNamespaceScriptBindingGetResponseType) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptBindingGetResponseTypeKVNamespace, DispatchNamespaceScriptBindingGetResponseTypeService, DispatchNamespaceScriptBindingGetResponseTypeDurableObjectNamespace, DispatchNamespaceScriptBindingGetResponseTypeR2Bucket, DispatchNamespaceScriptBindingGetResponseTypeQueue, DispatchNamespaceScriptBindingGetResponseTypeD1, DispatchNamespaceScriptBindingGetResponseTypeDispatchNamespace, DispatchNamespaceScriptBindingGetResponseTypeMTLSCertificate:
		return true
	}
	return false
}

type DispatchNamespaceScriptBindingGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
