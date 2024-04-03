// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

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

// ServiceEnvironmentSettingService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewServiceEnvironmentSettingService] method instead.
type ServiceEnvironmentSettingService struct {
	Options []option.RequestOption
}

// NewServiceEnvironmentSettingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewServiceEnvironmentSettingService(opts ...option.RequestOption) (r *ServiceEnvironmentSettingService) {
	r = &ServiceEnvironmentSettingService{}
	r.Options = opts
	return
}

// Patch script metadata, such as bindings
func (r *ServiceEnvironmentSettingService) Edit(ctx context.Context, serviceName string, environmentName string, params ServiceEnvironmentSettingEditParams, opts ...option.RequestOption) (res *ServiceEnvironmentSettingEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ServiceEnvironmentSettingEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/services/%s/environments/%s/settings", params.AccountID, serviceName, environmentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get script settings from a worker with an environment
func (r *ServiceEnvironmentSettingService) Get(ctx context.Context, serviceName string, environmentName string, query ServiceEnvironmentSettingGetParams, opts ...option.RequestOption) (res *ServiceEnvironmentSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ServiceEnvironmentSettingGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/services/%s/environments/%s/settings", query.AccountID, serviceName, environmentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ServiceEnvironmentSettingEditResponse struct {
	// List of bindings attached to this Worker
	Bindings []ServiceEnvironmentSettingEditResponseBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations ServiceEnvironmentSettingEditResponseMigrations `json:"migrations"`
	Placement  ServiceEnvironmentSettingEditResponsePlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []ServiceEnvironmentSettingEditResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                    `json:"usage_model"`
	JSON       serviceEnvironmentSettingEditResponseJSON `json:"-"`
}

// serviceEnvironmentSettingEditResponseJSON contains the JSON metadata for the
// struct [ServiceEnvironmentSettingEditResponse]
type serviceEnvironmentSettingEditResponseJSON struct {
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

func (r *ServiceEnvironmentSettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingEditResponseJSON) RawJSON() string {
	return r.raw
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [workers.ServiceEnvironmentSettingEditResponseBindingsWorkersKVNamespaceBinding],
// [workers.ServiceEnvironmentSettingEditResponseBindingsWorkersServiceBinding],
// [workers.ServiceEnvironmentSettingEditResponseBindingsWorkersDoBinding],
// [workers.ServiceEnvironmentSettingEditResponseBindingsWorkersR2Binding],
// [workers.ServiceEnvironmentSettingEditResponseBindingsWorkersQueueBinding],
// [workers.ServiceEnvironmentSettingEditResponseBindingsWorkersD1Binding],
// [workers.ServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBinding]
// or
// [workers.ServiceEnvironmentSettingEditResponseBindingsWorkersMTLSCERTBinding].
type ServiceEnvironmentSettingEditResponseBinding interface {
	implementsWorkersServiceEnvironmentSettingEditResponseBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ServiceEnvironmentSettingEditResponseBinding)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ServiceEnvironmentSettingEditResponseBindingsWorkersKVNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ServiceEnvironmentSettingEditResponseBindingsWorkersServiceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ServiceEnvironmentSettingEditResponseBindingsWorkersDoBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ServiceEnvironmentSettingEditResponseBindingsWorkersR2Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ServiceEnvironmentSettingEditResponseBindingsWorkersQueueBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ServiceEnvironmentSettingEditResponseBindingsWorkersD1Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ServiceEnvironmentSettingEditResponseBindingsWorkersMTLSCERTBinding{}),
		},
	)
}

type ServiceEnvironmentSettingEditResponseBindingsWorkersKVNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type ServiceEnvironmentSettingEditResponseBindingsWorkersKVNamespaceBindingType `json:"type,required"`
	JSON serviceEnvironmentSettingEditResponseBindingsWorkersKVNamespaceBindingJSON `json:"-"`
}

// serviceEnvironmentSettingEditResponseBindingsWorkersKVNamespaceBindingJSON
// contains the JSON metadata for the struct
// [ServiceEnvironmentSettingEditResponseBindingsWorkersKVNamespaceBinding]
type serviceEnvironmentSettingEditResponseBindingsWorkersKVNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingEditResponseBindingsWorkersKVNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingEditResponseBindingsWorkersKVNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r ServiceEnvironmentSettingEditResponseBindingsWorkersKVNamespaceBinding) implementsWorkersServiceEnvironmentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type ServiceEnvironmentSettingEditResponseBindingsWorkersKVNamespaceBindingType string

const (
	ServiceEnvironmentSettingEditResponseBindingsWorkersKVNamespaceBindingTypeKVNamespace ServiceEnvironmentSettingEditResponseBindingsWorkersKVNamespaceBindingType = "kv_namespace"
)

func (r ServiceEnvironmentSettingEditResponseBindingsWorkersKVNamespaceBindingType) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingEditResponseBindingsWorkersKVNamespaceBindingTypeKVNamespace:
		return true
	}
	return false
}

type ServiceEnvironmentSettingEditResponseBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type ServiceEnvironmentSettingEditResponseBindingsWorkersServiceBindingType `json:"type,required"`
	JSON serviceEnvironmentSettingEditResponseBindingsWorkersServiceBindingJSON `json:"-"`
}

// serviceEnvironmentSettingEditResponseBindingsWorkersServiceBindingJSON contains
// the JSON metadata for the struct
// [ServiceEnvironmentSettingEditResponseBindingsWorkersServiceBinding]
type serviceEnvironmentSettingEditResponseBindingsWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingEditResponseBindingsWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingEditResponseBindingsWorkersServiceBindingJSON) RawJSON() string {
	return r.raw
}

func (r ServiceEnvironmentSettingEditResponseBindingsWorkersServiceBinding) implementsWorkersServiceEnvironmentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type ServiceEnvironmentSettingEditResponseBindingsWorkersServiceBindingType string

const (
	ServiceEnvironmentSettingEditResponseBindingsWorkersServiceBindingTypeService ServiceEnvironmentSettingEditResponseBindingsWorkersServiceBindingType = "service"
)

func (r ServiceEnvironmentSettingEditResponseBindingsWorkersServiceBindingType) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingEditResponseBindingsWorkersServiceBindingTypeService:
		return true
	}
	return false
}

type ServiceEnvironmentSettingEditResponseBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type ServiceEnvironmentSettingEditResponseBindingsWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                            `json:"script_name"`
	JSON       serviceEnvironmentSettingEditResponseBindingsWorkersDoBindingJSON `json:"-"`
}

// serviceEnvironmentSettingEditResponseBindingsWorkersDoBindingJSON contains the
// JSON metadata for the struct
// [ServiceEnvironmentSettingEditResponseBindingsWorkersDoBinding]
type serviceEnvironmentSettingEditResponseBindingsWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingEditResponseBindingsWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingEditResponseBindingsWorkersDoBindingJSON) RawJSON() string {
	return r.raw
}

func (r ServiceEnvironmentSettingEditResponseBindingsWorkersDoBinding) implementsWorkersServiceEnvironmentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type ServiceEnvironmentSettingEditResponseBindingsWorkersDoBindingType string

const (
	ServiceEnvironmentSettingEditResponseBindingsWorkersDoBindingTypeDurableObjectNamespace ServiceEnvironmentSettingEditResponseBindingsWorkersDoBindingType = "durable_object_namespace"
)

func (r ServiceEnvironmentSettingEditResponseBindingsWorkersDoBindingType) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingEditResponseBindingsWorkersDoBindingTypeDurableObjectNamespace:
		return true
	}
	return false
}

type ServiceEnvironmentSettingEditResponseBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type ServiceEnvironmentSettingEditResponseBindingsWorkersR2BindingType `json:"type,required"`
	JSON serviceEnvironmentSettingEditResponseBindingsWorkersR2BindingJSON `json:"-"`
}

// serviceEnvironmentSettingEditResponseBindingsWorkersR2BindingJSON contains the
// JSON metadata for the struct
// [ServiceEnvironmentSettingEditResponseBindingsWorkersR2Binding]
type serviceEnvironmentSettingEditResponseBindingsWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingEditResponseBindingsWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingEditResponseBindingsWorkersR2BindingJSON) RawJSON() string {
	return r.raw
}

func (r ServiceEnvironmentSettingEditResponseBindingsWorkersR2Binding) implementsWorkersServiceEnvironmentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type ServiceEnvironmentSettingEditResponseBindingsWorkersR2BindingType string

const (
	ServiceEnvironmentSettingEditResponseBindingsWorkersR2BindingTypeR2Bucket ServiceEnvironmentSettingEditResponseBindingsWorkersR2BindingType = "r2_bucket"
)

func (r ServiceEnvironmentSettingEditResponseBindingsWorkersR2BindingType) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingEditResponseBindingsWorkersR2BindingTypeR2Bucket:
		return true
	}
	return false
}

type ServiceEnvironmentSettingEditResponseBindingsWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type ServiceEnvironmentSettingEditResponseBindingsWorkersQueueBindingType `json:"type,required"`
	JSON serviceEnvironmentSettingEditResponseBindingsWorkersQueueBindingJSON `json:"-"`
}

// serviceEnvironmentSettingEditResponseBindingsWorkersQueueBindingJSON contains
// the JSON metadata for the struct
// [ServiceEnvironmentSettingEditResponseBindingsWorkersQueueBinding]
type serviceEnvironmentSettingEditResponseBindingsWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingEditResponseBindingsWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingEditResponseBindingsWorkersQueueBindingJSON) RawJSON() string {
	return r.raw
}

func (r ServiceEnvironmentSettingEditResponseBindingsWorkersQueueBinding) implementsWorkersServiceEnvironmentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type ServiceEnvironmentSettingEditResponseBindingsWorkersQueueBindingType string

const (
	ServiceEnvironmentSettingEditResponseBindingsWorkersQueueBindingTypeQueue ServiceEnvironmentSettingEditResponseBindingsWorkersQueueBindingType = "queue"
)

func (r ServiceEnvironmentSettingEditResponseBindingsWorkersQueueBindingType) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingEditResponseBindingsWorkersQueueBindingTypeQueue:
		return true
	}
	return false
}

type ServiceEnvironmentSettingEditResponseBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type ServiceEnvironmentSettingEditResponseBindingsWorkersD1BindingType `json:"type,required"`
	JSON serviceEnvironmentSettingEditResponseBindingsWorkersD1BindingJSON `json:"-"`
}

// serviceEnvironmentSettingEditResponseBindingsWorkersD1BindingJSON contains the
// JSON metadata for the struct
// [ServiceEnvironmentSettingEditResponseBindingsWorkersD1Binding]
type serviceEnvironmentSettingEditResponseBindingsWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingEditResponseBindingsWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingEditResponseBindingsWorkersD1BindingJSON) RawJSON() string {
	return r.raw
}

func (r ServiceEnvironmentSettingEditResponseBindingsWorkersD1Binding) implementsWorkersServiceEnvironmentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type ServiceEnvironmentSettingEditResponseBindingsWorkersD1BindingType string

const (
	ServiceEnvironmentSettingEditResponseBindingsWorkersD1BindingTypeD1 ServiceEnvironmentSettingEditResponseBindingsWorkersD1BindingType = "d1"
)

func (r ServiceEnvironmentSettingEditResponseBindingsWorkersD1BindingType) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingEditResponseBindingsWorkersD1BindingTypeD1:
		return true
	}
	return false
}

type ServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type ServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound ServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     serviceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON     `json:"-"`
}

// serviceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [ServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBinding]
type serviceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r ServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBinding) implementsWorkersServiceEnvironmentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type ServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingType string

const (
	ServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace ServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

func (r ServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingType) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker
type ServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker ServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   serviceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// serviceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [ServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound]
type serviceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON) RawJSON() string {
	return r.raw
}

// Outbound worker
type ServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                         `json:"service"`
	JSON    serviceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// serviceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [ServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker]
type serviceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingEditResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type ServiceEnvironmentSettingEditResponseBindingsWorkersMTLSCERTBinding struct {
	Certificate interface{} `json:"certificate,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type ServiceEnvironmentSettingEditResponseBindingsWorkersMTLSCERTBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                                  `json:"certificate_id"`
	JSON          serviceEnvironmentSettingEditResponseBindingsWorkersMtlscertBindingJSON `json:"-"`
}

// serviceEnvironmentSettingEditResponseBindingsWorkersMtlscertBindingJSON contains
// the JSON metadata for the struct
// [ServiceEnvironmentSettingEditResponseBindingsWorkersMTLSCERTBinding]
type serviceEnvironmentSettingEditResponseBindingsWorkersMtlscertBindingJSON struct {
	Certificate   apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingEditResponseBindingsWorkersMTLSCERTBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingEditResponseBindingsWorkersMtlscertBindingJSON) RawJSON() string {
	return r.raw
}

func (r ServiceEnvironmentSettingEditResponseBindingsWorkersMTLSCERTBinding) implementsWorkersServiceEnvironmentSettingEditResponseBinding() {
}

// The class of resource that the binding provides.
type ServiceEnvironmentSettingEditResponseBindingsWorkersMTLSCERTBindingType string

const (
	ServiceEnvironmentSettingEditResponseBindingsWorkersMTLSCERTBindingTypeMTLSCertificate ServiceEnvironmentSettingEditResponseBindingsWorkersMTLSCERTBindingType = "mtls_certificate"
)

func (r ServiceEnvironmentSettingEditResponseBindingsWorkersMTLSCERTBindingType) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingEditResponseBindingsWorkersMTLSCERTBindingTypeMTLSCertificate:
		return true
	}
	return false
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [workers.ServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrations]
// or
// [workers.ServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrations].
type ServiceEnvironmentSettingEditResponseMigrations interface {
	implementsWorkersServiceEnvironmentSettingEditResponseMigrations()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ServiceEnvironmentSettingEditResponseMigrations)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrations{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrations{}),
		},
	)
}

// A single set of migrations to apply.
type ServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses []ServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []ServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               serviceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON               `json:"-"`
}

// serviceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON
// contains the JSON metadata for the struct
// [ServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrations]
type serviceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r ServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrations) implementsWorkersServiceEnvironmentSettingEditResponseMigrations() {
}

type ServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From string                                                                                     `json:"from"`
	To   string                                                                                     `json:"to"`
	JSON serviceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// serviceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [ServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass]
type serviceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON) RawJSON() string {
	return r.raw
}

type ServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       string                                                                                         `json:"from"`
	FromScript string                                                                                         `json:"from_script"`
	To         string                                                                                         `json:"to"`
	JSON       serviceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// serviceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [ServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass]
type serviceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingEditResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON) RawJSON() string {
	return r.raw
}

type ServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []ServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStep `json:"steps"`
	JSON  serviceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsJSON   `json:"-"`
}

// serviceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsJSON
// contains the JSON metadata for the struct
// [ServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrations]
type serviceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r ServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrations) implementsWorkersServiceEnvironmentSettingEditResponseMigrations() {
}

type ServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []ServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []ServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               serviceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON                `json:"-"`
}

// serviceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON
// contains the JSON metadata for the struct
// [ServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStep]
type serviceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepJSON) RawJSON() string {
	return r.raw
}

type ServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From string                                                                                       `json:"from"`
	To   string                                                                                       `json:"to"`
	JSON serviceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// serviceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [ServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass]
type serviceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON) RawJSON() string {
	return r.raw
}

type ServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                                           `json:"from"`
	FromScript string                                                                                           `json:"from_script"`
	To         string                                                                                           `json:"to"`
	JSON       serviceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// serviceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [ServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass]
type serviceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingEditResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON) RawJSON() string {
	return r.raw
}

type ServiceEnvironmentSettingEditResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode ServiceEnvironmentSettingEditResponsePlacementMode `json:"mode"`
	JSON serviceEnvironmentSettingEditResponsePlacementJSON `json:"-"`
}

// serviceEnvironmentSettingEditResponsePlacementJSON contains the JSON metadata
// for the struct [ServiceEnvironmentSettingEditResponsePlacement]
type serviceEnvironmentSettingEditResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingEditResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingEditResponsePlacementJSON) RawJSON() string {
	return r.raw
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type ServiceEnvironmentSettingEditResponsePlacementMode string

const (
	ServiceEnvironmentSettingEditResponsePlacementModeSmart ServiceEnvironmentSettingEditResponsePlacementMode = "smart"
)

func (r ServiceEnvironmentSettingEditResponsePlacementMode) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingEditResponsePlacementModeSmart:
		return true
	}
	return false
}

// A reference to a script that will consume logs from the attached Worker.
type ServiceEnvironmentSettingEditResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                `json:"namespace"`
	JSON      serviceEnvironmentSettingEditResponseTailConsumerJSON `json:"-"`
}

// serviceEnvironmentSettingEditResponseTailConsumerJSON contains the JSON metadata
// for the struct [ServiceEnvironmentSettingEditResponseTailConsumer]
type serviceEnvironmentSettingEditResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingEditResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingEditResponseTailConsumerJSON) RawJSON() string {
	return r.raw
}

type ServiceEnvironmentSettingGetResponse struct {
	// List of bindings attached to this Worker
	Bindings []ServiceEnvironmentSettingGetResponseBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations ServiceEnvironmentSettingGetResponseMigrations `json:"migrations"`
	Placement  ServiceEnvironmentSettingGetResponsePlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []ServiceEnvironmentSettingGetResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                   `json:"usage_model"`
	JSON       serviceEnvironmentSettingGetResponseJSON `json:"-"`
}

// serviceEnvironmentSettingGetResponseJSON contains the JSON metadata for the
// struct [ServiceEnvironmentSettingGetResponse]
type serviceEnvironmentSettingGetResponseJSON struct {
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

func (r *ServiceEnvironmentSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingGetResponseJSON) RawJSON() string {
	return r.raw
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [workers.ServiceEnvironmentSettingGetResponseBindingsWorkersKVNamespaceBinding],
// [workers.ServiceEnvironmentSettingGetResponseBindingsWorkersServiceBinding],
// [workers.ServiceEnvironmentSettingGetResponseBindingsWorkersDoBinding],
// [workers.ServiceEnvironmentSettingGetResponseBindingsWorkersR2Binding],
// [workers.ServiceEnvironmentSettingGetResponseBindingsWorkersQueueBinding],
// [workers.ServiceEnvironmentSettingGetResponseBindingsWorkersD1Binding],
// [workers.ServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBinding]
// or [workers.ServiceEnvironmentSettingGetResponseBindingsWorkersMTLSCERTBinding].
type ServiceEnvironmentSettingGetResponseBinding interface {
	implementsWorkersServiceEnvironmentSettingGetResponseBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ServiceEnvironmentSettingGetResponseBinding)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ServiceEnvironmentSettingGetResponseBindingsWorkersKVNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ServiceEnvironmentSettingGetResponseBindingsWorkersServiceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ServiceEnvironmentSettingGetResponseBindingsWorkersDoBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ServiceEnvironmentSettingGetResponseBindingsWorkersR2Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ServiceEnvironmentSettingGetResponseBindingsWorkersQueueBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ServiceEnvironmentSettingGetResponseBindingsWorkersD1Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ServiceEnvironmentSettingGetResponseBindingsWorkersMTLSCERTBinding{}),
		},
	)
}

type ServiceEnvironmentSettingGetResponseBindingsWorkersKVNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type ServiceEnvironmentSettingGetResponseBindingsWorkersKVNamespaceBindingType `json:"type,required"`
	JSON serviceEnvironmentSettingGetResponseBindingsWorkersKVNamespaceBindingJSON `json:"-"`
}

// serviceEnvironmentSettingGetResponseBindingsWorkersKVNamespaceBindingJSON
// contains the JSON metadata for the struct
// [ServiceEnvironmentSettingGetResponseBindingsWorkersKVNamespaceBinding]
type serviceEnvironmentSettingGetResponseBindingsWorkersKVNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingGetResponseBindingsWorkersKVNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingGetResponseBindingsWorkersKVNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r ServiceEnvironmentSettingGetResponseBindingsWorkersKVNamespaceBinding) implementsWorkersServiceEnvironmentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type ServiceEnvironmentSettingGetResponseBindingsWorkersKVNamespaceBindingType string

const (
	ServiceEnvironmentSettingGetResponseBindingsWorkersKVNamespaceBindingTypeKVNamespace ServiceEnvironmentSettingGetResponseBindingsWorkersKVNamespaceBindingType = "kv_namespace"
)

func (r ServiceEnvironmentSettingGetResponseBindingsWorkersKVNamespaceBindingType) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingGetResponseBindingsWorkersKVNamespaceBindingTypeKVNamespace:
		return true
	}
	return false
}

type ServiceEnvironmentSettingGetResponseBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type ServiceEnvironmentSettingGetResponseBindingsWorkersServiceBindingType `json:"type,required"`
	JSON serviceEnvironmentSettingGetResponseBindingsWorkersServiceBindingJSON `json:"-"`
}

// serviceEnvironmentSettingGetResponseBindingsWorkersServiceBindingJSON contains
// the JSON metadata for the struct
// [ServiceEnvironmentSettingGetResponseBindingsWorkersServiceBinding]
type serviceEnvironmentSettingGetResponseBindingsWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingGetResponseBindingsWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingGetResponseBindingsWorkersServiceBindingJSON) RawJSON() string {
	return r.raw
}

func (r ServiceEnvironmentSettingGetResponseBindingsWorkersServiceBinding) implementsWorkersServiceEnvironmentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type ServiceEnvironmentSettingGetResponseBindingsWorkersServiceBindingType string

const (
	ServiceEnvironmentSettingGetResponseBindingsWorkersServiceBindingTypeService ServiceEnvironmentSettingGetResponseBindingsWorkersServiceBindingType = "service"
)

func (r ServiceEnvironmentSettingGetResponseBindingsWorkersServiceBindingType) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingGetResponseBindingsWorkersServiceBindingTypeService:
		return true
	}
	return false
}

type ServiceEnvironmentSettingGetResponseBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type ServiceEnvironmentSettingGetResponseBindingsWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                           `json:"script_name"`
	JSON       serviceEnvironmentSettingGetResponseBindingsWorkersDoBindingJSON `json:"-"`
}

// serviceEnvironmentSettingGetResponseBindingsWorkersDoBindingJSON contains the
// JSON metadata for the struct
// [ServiceEnvironmentSettingGetResponseBindingsWorkersDoBinding]
type serviceEnvironmentSettingGetResponseBindingsWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingGetResponseBindingsWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingGetResponseBindingsWorkersDoBindingJSON) RawJSON() string {
	return r.raw
}

func (r ServiceEnvironmentSettingGetResponseBindingsWorkersDoBinding) implementsWorkersServiceEnvironmentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type ServiceEnvironmentSettingGetResponseBindingsWorkersDoBindingType string

const (
	ServiceEnvironmentSettingGetResponseBindingsWorkersDoBindingTypeDurableObjectNamespace ServiceEnvironmentSettingGetResponseBindingsWorkersDoBindingType = "durable_object_namespace"
)

func (r ServiceEnvironmentSettingGetResponseBindingsWorkersDoBindingType) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingGetResponseBindingsWorkersDoBindingTypeDurableObjectNamespace:
		return true
	}
	return false
}

type ServiceEnvironmentSettingGetResponseBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type ServiceEnvironmentSettingGetResponseBindingsWorkersR2BindingType `json:"type,required"`
	JSON serviceEnvironmentSettingGetResponseBindingsWorkersR2BindingJSON `json:"-"`
}

// serviceEnvironmentSettingGetResponseBindingsWorkersR2BindingJSON contains the
// JSON metadata for the struct
// [ServiceEnvironmentSettingGetResponseBindingsWorkersR2Binding]
type serviceEnvironmentSettingGetResponseBindingsWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingGetResponseBindingsWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingGetResponseBindingsWorkersR2BindingJSON) RawJSON() string {
	return r.raw
}

func (r ServiceEnvironmentSettingGetResponseBindingsWorkersR2Binding) implementsWorkersServiceEnvironmentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type ServiceEnvironmentSettingGetResponseBindingsWorkersR2BindingType string

const (
	ServiceEnvironmentSettingGetResponseBindingsWorkersR2BindingTypeR2Bucket ServiceEnvironmentSettingGetResponseBindingsWorkersR2BindingType = "r2_bucket"
)

func (r ServiceEnvironmentSettingGetResponseBindingsWorkersR2BindingType) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingGetResponseBindingsWorkersR2BindingTypeR2Bucket:
		return true
	}
	return false
}

type ServiceEnvironmentSettingGetResponseBindingsWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type ServiceEnvironmentSettingGetResponseBindingsWorkersQueueBindingType `json:"type,required"`
	JSON serviceEnvironmentSettingGetResponseBindingsWorkersQueueBindingJSON `json:"-"`
}

// serviceEnvironmentSettingGetResponseBindingsWorkersQueueBindingJSON contains the
// JSON metadata for the struct
// [ServiceEnvironmentSettingGetResponseBindingsWorkersQueueBinding]
type serviceEnvironmentSettingGetResponseBindingsWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingGetResponseBindingsWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingGetResponseBindingsWorkersQueueBindingJSON) RawJSON() string {
	return r.raw
}

func (r ServiceEnvironmentSettingGetResponseBindingsWorkersQueueBinding) implementsWorkersServiceEnvironmentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type ServiceEnvironmentSettingGetResponseBindingsWorkersQueueBindingType string

const (
	ServiceEnvironmentSettingGetResponseBindingsWorkersQueueBindingTypeQueue ServiceEnvironmentSettingGetResponseBindingsWorkersQueueBindingType = "queue"
)

func (r ServiceEnvironmentSettingGetResponseBindingsWorkersQueueBindingType) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingGetResponseBindingsWorkersQueueBindingTypeQueue:
		return true
	}
	return false
}

type ServiceEnvironmentSettingGetResponseBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type ServiceEnvironmentSettingGetResponseBindingsWorkersD1BindingType `json:"type,required"`
	JSON serviceEnvironmentSettingGetResponseBindingsWorkersD1BindingJSON `json:"-"`
}

// serviceEnvironmentSettingGetResponseBindingsWorkersD1BindingJSON contains the
// JSON metadata for the struct
// [ServiceEnvironmentSettingGetResponseBindingsWorkersD1Binding]
type serviceEnvironmentSettingGetResponseBindingsWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingGetResponseBindingsWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingGetResponseBindingsWorkersD1BindingJSON) RawJSON() string {
	return r.raw
}

func (r ServiceEnvironmentSettingGetResponseBindingsWorkersD1Binding) implementsWorkersServiceEnvironmentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type ServiceEnvironmentSettingGetResponseBindingsWorkersD1BindingType string

const (
	ServiceEnvironmentSettingGetResponseBindingsWorkersD1BindingTypeD1 ServiceEnvironmentSettingGetResponseBindingsWorkersD1BindingType = "d1"
)

func (r ServiceEnvironmentSettingGetResponseBindingsWorkersD1BindingType) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingGetResponseBindingsWorkersD1BindingTypeD1:
		return true
	}
	return false
}

type ServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type ServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound ServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     serviceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON     `json:"-"`
}

// serviceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [ServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBinding]
type serviceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r ServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBinding) implementsWorkersServiceEnvironmentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type ServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingType string

const (
	ServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace ServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

func (r ServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingType) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker
type ServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker ServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   serviceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// serviceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [ServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound]
type serviceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundJSON) RawJSON() string {
	return r.raw
}

// Outbound worker
type ServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                        `json:"service"`
	JSON    serviceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// serviceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [ServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker]
type serviceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingGetResponseBindingsWorkersDispatchNamespaceBindingOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type ServiceEnvironmentSettingGetResponseBindingsWorkersMTLSCERTBinding struct {
	Certificate interface{} `json:"certificate,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type ServiceEnvironmentSettingGetResponseBindingsWorkersMTLSCERTBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                                 `json:"certificate_id"`
	JSON          serviceEnvironmentSettingGetResponseBindingsWorkersMtlscertBindingJSON `json:"-"`
}

// serviceEnvironmentSettingGetResponseBindingsWorkersMtlscertBindingJSON contains
// the JSON metadata for the struct
// [ServiceEnvironmentSettingGetResponseBindingsWorkersMTLSCERTBinding]
type serviceEnvironmentSettingGetResponseBindingsWorkersMtlscertBindingJSON struct {
	Certificate   apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingGetResponseBindingsWorkersMTLSCERTBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingGetResponseBindingsWorkersMtlscertBindingJSON) RawJSON() string {
	return r.raw
}

func (r ServiceEnvironmentSettingGetResponseBindingsWorkersMTLSCERTBinding) implementsWorkersServiceEnvironmentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type ServiceEnvironmentSettingGetResponseBindingsWorkersMTLSCERTBindingType string

const (
	ServiceEnvironmentSettingGetResponseBindingsWorkersMTLSCERTBindingTypeMTLSCertificate ServiceEnvironmentSettingGetResponseBindingsWorkersMTLSCERTBindingType = "mtls_certificate"
)

func (r ServiceEnvironmentSettingGetResponseBindingsWorkersMTLSCERTBindingType) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingGetResponseBindingsWorkersMTLSCERTBindingTypeMTLSCertificate:
		return true
	}
	return false
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [workers.ServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrations]
// or
// [workers.ServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrations].
type ServiceEnvironmentSettingGetResponseMigrations interface {
	implementsWorkersServiceEnvironmentSettingGetResponseMigrations()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ServiceEnvironmentSettingGetResponseMigrations)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrations{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrations{}),
		},
	)
}

// A single set of migrations to apply.
type ServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses []ServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []ServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               serviceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON               `json:"-"`
}

// serviceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON
// contains the JSON metadata for the struct
// [ServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrations]
type serviceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r ServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrations) implementsWorkersServiceEnvironmentSettingGetResponseMigrations() {
}

type ServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From string                                                                                    `json:"from"`
	To   string                                                                                    `json:"to"`
	JSON serviceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// serviceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [ServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass]
type serviceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsRenamedClassJSON) RawJSON() string {
	return r.raw
}

type ServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       string                                                                                        `json:"from"`
	FromScript string                                                                                        `json:"from_script"`
	To         string                                                                                        `json:"to"`
	JSON       serviceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// serviceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [ServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass]
type serviceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingGetResponseMigrationsWorkersSingleStepMigrationsTransferredClassJSON) RawJSON() string {
	return r.raw
}

type ServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []ServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStep `json:"steps"`
	JSON  serviceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsJSON   `json:"-"`
}

// serviceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsJSON
// contains the JSON metadata for the struct
// [ServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrations]
type serviceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r ServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrations) implementsWorkersServiceEnvironmentSettingGetResponseMigrations() {
}

type ServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []ServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []ServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               serviceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON                `json:"-"`
}

// serviceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON
// contains the JSON metadata for the struct
// [ServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStep]
type serviceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepJSON) RawJSON() string {
	return r.raw
}

type ServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From string                                                                                      `json:"from"`
	To   string                                                                                      `json:"to"`
	JSON serviceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// serviceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [ServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass]
type serviceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsRenamedClassJSON) RawJSON() string {
	return r.raw
}

type ServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                                          `json:"from"`
	FromScript string                                                                                          `json:"from_script"`
	To         string                                                                                          `json:"to"`
	JSON       serviceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// serviceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [ServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass]
type serviceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingGetResponseMigrationsWorkersSteppedMigrationsStepsTransferredClassJSON) RawJSON() string {
	return r.raw
}

type ServiceEnvironmentSettingGetResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode ServiceEnvironmentSettingGetResponsePlacementMode `json:"mode"`
	JSON serviceEnvironmentSettingGetResponsePlacementJSON `json:"-"`
}

// serviceEnvironmentSettingGetResponsePlacementJSON contains the JSON metadata for
// the struct [ServiceEnvironmentSettingGetResponsePlacement]
type serviceEnvironmentSettingGetResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingGetResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingGetResponsePlacementJSON) RawJSON() string {
	return r.raw
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type ServiceEnvironmentSettingGetResponsePlacementMode string

const (
	ServiceEnvironmentSettingGetResponsePlacementModeSmart ServiceEnvironmentSettingGetResponsePlacementMode = "smart"
)

func (r ServiceEnvironmentSettingGetResponsePlacementMode) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingGetResponsePlacementModeSmart:
		return true
	}
	return false
}

// A reference to a script that will consume logs from the attached Worker.
type ServiceEnvironmentSettingGetResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                               `json:"namespace"`
	JSON      serviceEnvironmentSettingGetResponseTailConsumerJSON `json:"-"`
}

// serviceEnvironmentSettingGetResponseTailConsumerJSON contains the JSON metadata
// for the struct [ServiceEnvironmentSettingGetResponseTailConsumer]
type serviceEnvironmentSettingGetResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingGetResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingGetResponseTailConsumerJSON) RawJSON() string {
	return r.raw
}

type ServiceEnvironmentSettingEditParams struct {
	// Identifier
	AccountID param.Field[string]                                       `path:"account_id,required"`
	Errors    param.Field[[]ServiceEnvironmentSettingEditParamsError]   `json:"errors,required"`
	Messages  param.Field[[]ServiceEnvironmentSettingEditParamsMessage] `json:"messages,required"`
	Result    param.Field[ServiceEnvironmentSettingEditParamsResult]    `json:"result,required"`
	// Whether the API call was successful
	Success param.Field[ServiceEnvironmentSettingEditParamsSuccess] `json:"success,required"`
}

func (r ServiceEnvironmentSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ServiceEnvironmentSettingEditParamsError struct {
	Code    param.Field[int64]  `json:"code,required"`
	Message param.Field[string] `json:"message,required"`
}

func (r ServiceEnvironmentSettingEditParamsError) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ServiceEnvironmentSettingEditParamsMessage struct {
	Code    param.Field[int64]  `json:"code,required"`
	Message param.Field[string] `json:"message,required"`
}

func (r ServiceEnvironmentSettingEditParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ServiceEnvironmentSettingEditParamsResult struct {
	// List of bindings attached to this Worker
	Bindings param.Field[[]ServiceEnvironmentSettingEditParamsResultBinding] `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations param.Field[ServiceEnvironmentSettingEditParamsResultMigrations] `json:"migrations"`
	Placement  param.Field[ServiceEnvironmentSettingEditParamsResultPlacement]  `json:"placement"`
	// Tags to help you manage your Workers
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]ServiceEnvironmentSettingEditParamsResultTailConsumer] `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel param.Field[string] `json:"usage_model"`
}

func (r ServiceEnvironmentSettingEditParamsResult) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A binding to allow the Worker to communicate with resources
//
// Satisfied by
// [workers.ServiceEnvironmentSettingEditParamsResultBindingsWorkersKVNamespaceBinding],
// [workers.ServiceEnvironmentSettingEditParamsResultBindingsWorkersServiceBinding],
// [workers.ServiceEnvironmentSettingEditParamsResultBindingsWorkersDoBinding],
// [workers.ServiceEnvironmentSettingEditParamsResultBindingsWorkersR2Binding],
// [workers.ServiceEnvironmentSettingEditParamsResultBindingsWorkersQueueBinding],
// [workers.ServiceEnvironmentSettingEditParamsResultBindingsWorkersD1Binding],
// [workers.ServiceEnvironmentSettingEditParamsResultBindingsWorkersDispatchNamespaceBinding],
// [workers.ServiceEnvironmentSettingEditParamsResultBindingsWorkersMTLSCERTBinding].
type ServiceEnvironmentSettingEditParamsResultBinding interface {
	implementsWorkersServiceEnvironmentSettingEditParamsResultBinding()
}

type ServiceEnvironmentSettingEditParamsResultBindingsWorkersKVNamespaceBinding struct {
	// The class of resource that the binding provides.
	Type param.Field[ServiceEnvironmentSettingEditParamsResultBindingsWorkersKVNamespaceBindingType] `json:"type,required"`
}

func (r ServiceEnvironmentSettingEditParamsResultBindingsWorkersKVNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ServiceEnvironmentSettingEditParamsResultBindingsWorkersKVNamespaceBinding) implementsWorkersServiceEnvironmentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type ServiceEnvironmentSettingEditParamsResultBindingsWorkersKVNamespaceBindingType string

const (
	ServiceEnvironmentSettingEditParamsResultBindingsWorkersKVNamespaceBindingTypeKVNamespace ServiceEnvironmentSettingEditParamsResultBindingsWorkersKVNamespaceBindingType = "kv_namespace"
)

func (r ServiceEnvironmentSettingEditParamsResultBindingsWorkersKVNamespaceBindingType) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingEditParamsResultBindingsWorkersKVNamespaceBindingTypeKVNamespace:
		return true
	}
	return false
}

type ServiceEnvironmentSettingEditParamsResultBindingsWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// Name of Worker to bind to
	Service param.Field[string] `json:"service,required"`
	// The class of resource that the binding provides.
	Type param.Field[ServiceEnvironmentSettingEditParamsResultBindingsWorkersServiceBindingType] `json:"type,required"`
}

func (r ServiceEnvironmentSettingEditParamsResultBindingsWorkersServiceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ServiceEnvironmentSettingEditParamsResultBindingsWorkersServiceBinding) implementsWorkersServiceEnvironmentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type ServiceEnvironmentSettingEditParamsResultBindingsWorkersServiceBindingType string

const (
	ServiceEnvironmentSettingEditParamsResultBindingsWorkersServiceBindingTypeService ServiceEnvironmentSettingEditParamsResultBindingsWorkersServiceBindingType = "service"
)

func (r ServiceEnvironmentSettingEditParamsResultBindingsWorkersServiceBindingType) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingEditParamsResultBindingsWorkersServiceBindingTypeService:
		return true
	}
	return false
}

type ServiceEnvironmentSettingEditParamsResultBindingsWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName param.Field[string] `json:"class_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[ServiceEnvironmentSettingEditParamsResultBindingsWorkersDoBindingType] `json:"type,required"`
	// The environment of the script_name to bind to
	Environment param.Field[string] `json:"environment"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName param.Field[string] `json:"script_name"`
}

func (r ServiceEnvironmentSettingEditParamsResultBindingsWorkersDoBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ServiceEnvironmentSettingEditParamsResultBindingsWorkersDoBinding) implementsWorkersServiceEnvironmentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type ServiceEnvironmentSettingEditParamsResultBindingsWorkersDoBindingType string

const (
	ServiceEnvironmentSettingEditParamsResultBindingsWorkersDoBindingTypeDurableObjectNamespace ServiceEnvironmentSettingEditParamsResultBindingsWorkersDoBindingType = "durable_object_namespace"
)

func (r ServiceEnvironmentSettingEditParamsResultBindingsWorkersDoBindingType) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingEditParamsResultBindingsWorkersDoBindingTypeDurableObjectNamespace:
		return true
	}
	return false
}

type ServiceEnvironmentSettingEditParamsResultBindingsWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName param.Field[string] `json:"bucket_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[ServiceEnvironmentSettingEditParamsResultBindingsWorkersR2BindingType] `json:"type,required"`
}

func (r ServiceEnvironmentSettingEditParamsResultBindingsWorkersR2Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ServiceEnvironmentSettingEditParamsResultBindingsWorkersR2Binding) implementsWorkersServiceEnvironmentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type ServiceEnvironmentSettingEditParamsResultBindingsWorkersR2BindingType string

const (
	ServiceEnvironmentSettingEditParamsResultBindingsWorkersR2BindingTypeR2Bucket ServiceEnvironmentSettingEditParamsResultBindingsWorkersR2BindingType = "r2_bucket"
)

func (r ServiceEnvironmentSettingEditParamsResultBindingsWorkersR2BindingType) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingEditParamsResultBindingsWorkersR2BindingTypeR2Bucket:
		return true
	}
	return false
}

type ServiceEnvironmentSettingEditParamsResultBindingsWorkersQueueBinding struct {
	// Name of the Queue to bind to
	QueueName param.Field[string] `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[ServiceEnvironmentSettingEditParamsResultBindingsWorkersQueueBindingType] `json:"type,required"`
}

func (r ServiceEnvironmentSettingEditParamsResultBindingsWorkersQueueBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ServiceEnvironmentSettingEditParamsResultBindingsWorkersQueueBinding) implementsWorkersServiceEnvironmentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type ServiceEnvironmentSettingEditParamsResultBindingsWorkersQueueBindingType string

const (
	ServiceEnvironmentSettingEditParamsResultBindingsWorkersQueueBindingTypeQueue ServiceEnvironmentSettingEditParamsResultBindingsWorkersQueueBindingType = "queue"
)

func (r ServiceEnvironmentSettingEditParamsResultBindingsWorkersQueueBindingType) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingEditParamsResultBindingsWorkersQueueBindingTypeQueue:
		return true
	}
	return false
}

type ServiceEnvironmentSettingEditParamsResultBindingsWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID param.Field[string] `json:"id,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name param.Field[string] `json:"name,required"`
	// The class of resource that the binding provides.
	Type param.Field[ServiceEnvironmentSettingEditParamsResultBindingsWorkersD1BindingType] `json:"type,required"`
}

func (r ServiceEnvironmentSettingEditParamsResultBindingsWorkersD1Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ServiceEnvironmentSettingEditParamsResultBindingsWorkersD1Binding) implementsWorkersServiceEnvironmentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type ServiceEnvironmentSettingEditParamsResultBindingsWorkersD1BindingType string

const (
	ServiceEnvironmentSettingEditParamsResultBindingsWorkersD1BindingTypeD1 ServiceEnvironmentSettingEditParamsResultBindingsWorkersD1BindingType = "d1"
)

func (r ServiceEnvironmentSettingEditParamsResultBindingsWorkersD1BindingType) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingEditParamsResultBindingsWorkersD1BindingTypeD1:
		return true
	}
	return false
}

type ServiceEnvironmentSettingEditParamsResultBindingsWorkersDispatchNamespaceBinding struct {
	// Namespace to bind to
	Namespace param.Field[string] `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type param.Field[ServiceEnvironmentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingType] `json:"type,required"`
	// Outbound worker
	Outbound param.Field[ServiceEnvironmentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutbound] `json:"outbound"`
}

func (r ServiceEnvironmentSettingEditParamsResultBindingsWorkersDispatchNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ServiceEnvironmentSettingEditParamsResultBindingsWorkersDispatchNamespaceBinding) implementsWorkersServiceEnvironmentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type ServiceEnvironmentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingType string

const (
	ServiceEnvironmentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace ServiceEnvironmentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingType = "dispatch_namespace"
)

func (r ServiceEnvironmentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingType) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker
type ServiceEnvironmentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params param.Field[[]string] `json:"params"`
	// Outbound worker
	Worker param.Field[ServiceEnvironmentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker] `json:"worker"`
}

func (r ServiceEnvironmentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutbound) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker
type ServiceEnvironmentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker
	Service param.Field[string] `json:"service"`
}

func (r ServiceEnvironmentSettingEditParamsResultBindingsWorkersDispatchNamespaceBindingOutboundWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ServiceEnvironmentSettingEditParamsResultBindingsWorkersMTLSCERTBinding struct {
	Certificate param.Field[interface{}] `json:"certificate,required"`
	// The class of resource that the binding provides.
	Type param.Field[ServiceEnvironmentSettingEditParamsResultBindingsWorkersMTLSCERTBindingType] `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID param.Field[string] `json:"certificate_id"`
}

func (r ServiceEnvironmentSettingEditParamsResultBindingsWorkersMTLSCERTBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ServiceEnvironmentSettingEditParamsResultBindingsWorkersMTLSCERTBinding) implementsWorkersServiceEnvironmentSettingEditParamsResultBinding() {
}

// The class of resource that the binding provides.
type ServiceEnvironmentSettingEditParamsResultBindingsWorkersMTLSCERTBindingType string

const (
	ServiceEnvironmentSettingEditParamsResultBindingsWorkersMTLSCERTBindingTypeMTLSCertificate ServiceEnvironmentSettingEditParamsResultBindingsWorkersMTLSCERTBindingType = "mtls_certificate"
)

func (r ServiceEnvironmentSettingEditParamsResultBindingsWorkersMTLSCERTBindingType) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingEditParamsResultBindingsWorkersMTLSCERTBindingTypeMTLSCertificate:
		return true
	}
	return false
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by
// [workers.ServiceEnvironmentSettingEditParamsResultMigrationsWorkersSingleStepMigrations],
// [workers.ServiceEnvironmentSettingEditParamsResultMigrationsWorkersSteppedMigrations].
type ServiceEnvironmentSettingEditParamsResultMigrations interface {
	implementsWorkersServiceEnvironmentSettingEditParamsResultMigrations()
}

// A single set of migrations to apply.
type ServiceEnvironmentSettingEditParamsResultMigrationsWorkersSingleStepMigrations struct {
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
	RenamedClasses param.Field[[]ServiceEnvironmentSettingEditParamsResultMigrationsWorkersSingleStepMigrationsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]ServiceEnvironmentSettingEditParamsResultMigrationsWorkersSingleStepMigrationsTransferredClass] `json:"transferred_classes"`
}

func (r ServiceEnvironmentSettingEditParamsResultMigrationsWorkersSingleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ServiceEnvironmentSettingEditParamsResultMigrationsWorkersSingleStepMigrations) implementsWorkersServiceEnvironmentSettingEditParamsResultMigrations() {
}

type ServiceEnvironmentSettingEditParamsResultMigrationsWorkersSingleStepMigrationsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r ServiceEnvironmentSettingEditParamsResultMigrationsWorkersSingleStepMigrationsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ServiceEnvironmentSettingEditParamsResultMigrationsWorkersSingleStepMigrationsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r ServiceEnvironmentSettingEditParamsResultMigrationsWorkersSingleStepMigrationsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ServiceEnvironmentSettingEditParamsResultMigrationsWorkersSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]ServiceEnvironmentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStep] `json:"steps"`
}

func (r ServiceEnvironmentSettingEditParamsResultMigrationsWorkersSteppedMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ServiceEnvironmentSettingEditParamsResultMigrationsWorkersSteppedMigrations) implementsWorkersServiceEnvironmentSettingEditParamsResultMigrations() {
}

type ServiceEnvironmentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]ServiceEnvironmentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]ServiceEnvironmentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass] `json:"transferred_classes"`
}

func (r ServiceEnvironmentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStep) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ServiceEnvironmentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r ServiceEnvironmentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ServiceEnvironmentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r ServiceEnvironmentSettingEditParamsResultMigrationsWorkersSteppedMigrationsStepsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ServiceEnvironmentSettingEditParamsResultPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode param.Field[ServiceEnvironmentSettingEditParamsResultPlacementMode] `json:"mode"`
}

func (r ServiceEnvironmentSettingEditParamsResultPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type ServiceEnvironmentSettingEditParamsResultPlacementMode string

const (
	ServiceEnvironmentSettingEditParamsResultPlacementModeSmart ServiceEnvironmentSettingEditParamsResultPlacementMode = "smart"
)

func (r ServiceEnvironmentSettingEditParamsResultPlacementMode) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingEditParamsResultPlacementModeSmart:
		return true
	}
	return false
}

// A reference to a script that will consume logs from the attached Worker.
type ServiceEnvironmentSettingEditParamsResultTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service param.Field[string] `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace param.Field[string] `json:"namespace"`
}

func (r ServiceEnvironmentSettingEditParamsResultTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the API call was successful
type ServiceEnvironmentSettingEditParamsSuccess bool

const (
	ServiceEnvironmentSettingEditParamsSuccessTrue ServiceEnvironmentSettingEditParamsSuccess = true
)

func (r ServiceEnvironmentSettingEditParamsSuccess) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingEditParamsSuccessTrue:
		return true
	}
	return false
}

type ServiceEnvironmentSettingEditResponseEnvelope struct {
	Errors   []ServiceEnvironmentSettingEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ServiceEnvironmentSettingEditResponseEnvelopeMessages `json:"messages,required"`
	Result   ServiceEnvironmentSettingEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ServiceEnvironmentSettingEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    serviceEnvironmentSettingEditResponseEnvelopeJSON    `json:"-"`
}

// serviceEnvironmentSettingEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [ServiceEnvironmentSettingEditResponseEnvelope]
type serviceEnvironmentSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ServiceEnvironmentSettingEditResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    serviceEnvironmentSettingEditResponseEnvelopeErrorsJSON `json:"-"`
}

// serviceEnvironmentSettingEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ServiceEnvironmentSettingEditResponseEnvelopeErrors]
type serviceEnvironmentSettingEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ServiceEnvironmentSettingEditResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    serviceEnvironmentSettingEditResponseEnvelopeMessagesJSON `json:"-"`
}

// serviceEnvironmentSettingEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ServiceEnvironmentSettingEditResponseEnvelopeMessages]
type serviceEnvironmentSettingEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ServiceEnvironmentSettingEditResponseEnvelopeSuccess bool

const (
	ServiceEnvironmentSettingEditResponseEnvelopeSuccessTrue ServiceEnvironmentSettingEditResponseEnvelopeSuccess = true
)

func (r ServiceEnvironmentSettingEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ServiceEnvironmentSettingGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ServiceEnvironmentSettingGetResponseEnvelope struct {
	Errors   []ServiceEnvironmentSettingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ServiceEnvironmentSettingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ServiceEnvironmentSettingGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ServiceEnvironmentSettingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    serviceEnvironmentSettingGetResponseEnvelopeJSON    `json:"-"`
}

// serviceEnvironmentSettingGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ServiceEnvironmentSettingGetResponseEnvelope]
type serviceEnvironmentSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ServiceEnvironmentSettingGetResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    serviceEnvironmentSettingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// serviceEnvironmentSettingGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ServiceEnvironmentSettingGetResponseEnvelopeErrors]
type serviceEnvironmentSettingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ServiceEnvironmentSettingGetResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    serviceEnvironmentSettingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// serviceEnvironmentSettingGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ServiceEnvironmentSettingGetResponseEnvelopeMessages]
type serviceEnvironmentSettingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceEnvironmentSettingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceEnvironmentSettingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ServiceEnvironmentSettingGetResponseEnvelopeSuccess bool

const (
	ServiceEnvironmentSettingGetResponseEnvelopeSuccessTrue ServiceEnvironmentSettingGetResponseEnvelopeSuccess = true
)

func (r ServiceEnvironmentSettingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ServiceEnvironmentSettingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
