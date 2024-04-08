// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// WorkerService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewWorkerService] method instead.
type WorkerService struct {
	Options         []option.RequestOption
	AI              *AIService
	Scripts         *ScriptService
	Filters         *FilterService
	Routes          *RouteService
	AccountSettings *AccountSettingService
	Domains         *DomainService
	Subdomains      *SubdomainService
	Services        *ServiceService
}

// NewWorkerService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWorkerService(opts ...option.RequestOption) (r *WorkerService) {
	r = &WorkerService{}
	r.Options = opts
	r.AI = NewAIService(opts...)
	r.Scripts = NewScriptService(opts...)
	r.Filters = NewFilterService(opts...)
	r.Routes = NewRouteService(opts...)
	r.AccountSettings = NewAccountSettingService(opts...)
	r.Domains = NewDomainService(opts...)
	r.Subdomains = NewSubdomainService(opts...)
	r.Services = NewServiceService(opts...)
	return
}

// A binding to allow the Worker to communicate with resources
type Binding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The class of resource that the binding provides.
	Type BindingType `json:"type,required"`
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
	CertificateID string      `json:"certificate_id"`
	Certificate   interface{} `json:"certificate,required"`
	JSON          bindingJSON `json:"-"`
	union         BindingUnion
}

// bindingJSON contains the JSON metadata for the struct [Binding]
type bindingJSON struct {
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

func (r bindingJSON) RawJSON() string {
	return r.raw
}

func (r *Binding) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r Binding) AsUnion() BindingUnion {
	return r.union
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by [workers.KVNamespaceBinding], [workers.ServiceBinding],
// [workers.DurableObjectBinding], [workers.R2Binding],
// [workers.BindingWorkersQueueBinding], [workers.D1Binding],
// [workers.DispatchNamespaceBinding] or [workers.MTLSCERTBinding].
type BindingUnion interface {
	implementsWorkersBinding()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BindingUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(KVNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ServiceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DurableObjectBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(R2Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BindingWorkersQueueBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(D1Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MTLSCERTBinding{}),
		},
	)
}

type BindingWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type BindingWorkersQueueBindingType `json:"type,required"`
	JSON bindingWorkersQueueBindingJSON `json:"-"`
}

// bindingWorkersQueueBindingJSON contains the JSON metadata for the struct
// [BindingWorkersQueueBinding]
type bindingWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingWorkersQueueBindingJSON) RawJSON() string {
	return r.raw
}

func (r BindingWorkersQueueBinding) implementsWorkersBinding() {}

// The class of resource that the binding provides.
type BindingWorkersQueueBindingType string

const (
	BindingWorkersQueueBindingTypeQueue BindingWorkersQueueBindingType = "queue"
)

func (r BindingWorkersQueueBindingType) IsKnown() bool {
	switch r {
	case BindingWorkersQueueBindingTypeQueue:
		return true
	}
	return false
}

// The class of resource that the binding provides.
type BindingType string

const (
	BindingTypeKVNamespace            BindingType = "kv_namespace"
	BindingTypeService                BindingType = "service"
	BindingTypeDurableObjectNamespace BindingType = "durable_object_namespace"
	BindingTypeR2Bucket               BindingType = "r2_bucket"
	BindingTypeQueue                  BindingType = "queue"
	BindingTypeD1                     BindingType = "d1"
	BindingTypeDispatchNamespace      BindingType = "dispatch_namespace"
	BindingTypeMTLSCertificate        BindingType = "mtls_certificate"
)

func (r BindingType) IsKnown() bool {
	switch r {
	case BindingTypeKVNamespace, BindingTypeService, BindingTypeDurableObjectNamespace, BindingTypeR2Bucket, BindingTypeQueue, BindingTypeD1, BindingTypeDispatchNamespace, BindingTypeMTLSCertificate:
		return true
	}
	return false
}

type D1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type D1BindingType `json:"type,required"`
	JSON d1BindingJSON `json:"-"`
}

// d1BindingJSON contains the JSON metadata for the struct [D1Binding]
type d1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *D1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r d1BindingJSON) RawJSON() string {
	return r.raw
}

func (r D1Binding) implementsWorkersBinding() {}

// The class of resource that the binding provides.
type D1BindingType string

const (
	D1BindingTypeD1 D1BindingType = "d1"
)

func (r D1BindingType) IsKnown() bool {
	switch r {
	case D1BindingTypeD1:
		return true
	}
	return false
}

type DispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type DispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound DispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     dispatchNamespaceBindingJSON     `json:"-"`
}

// dispatchNamespaceBindingJSON contains the JSON metadata for the struct
// [DispatchNamespaceBinding]
type dispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceBinding) implementsWorkersBinding() {}

// The class of resource that the binding provides.
type DispatchNamespaceBindingType string

const (
	DispatchNamespaceBindingTypeDispatchNamespace DispatchNamespaceBindingType = "dispatch_namespace"
)

func (r DispatchNamespaceBindingType) IsKnown() bool {
	switch r {
	case DispatchNamespaceBindingTypeDispatchNamespace:
		return true
	}
	return false
}

// Outbound worker
type DispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker DispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   dispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// dispatchNamespaceBindingOutboundJSON contains the JSON metadata for the struct
// [DispatchNamespaceBindingOutbound]
type dispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceBindingOutboundJSON) RawJSON() string {
	return r.raw
}

// Outbound worker
type DispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                     `json:"service"`
	JSON    dispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// dispatchNamespaceBindingOutboundWorkerJSON contains the JSON metadata for the
// struct [DispatchNamespaceBindingOutboundWorker]
type dispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceBindingOutboundWorkerJSON) RawJSON() string {
	return r.raw
}

type DurableObjectBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type DurableObjectBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                   `json:"script_name"`
	JSON       durableObjectBindingJSON `json:"-"`
}

// durableObjectBindingJSON contains the JSON metadata for the struct
// [DurableObjectBinding]
type durableObjectBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DurableObjectBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r durableObjectBindingJSON) RawJSON() string {
	return r.raw
}

func (r DurableObjectBinding) implementsWorkersBinding() {}

// The class of resource that the binding provides.
type DurableObjectBindingType string

const (
	DurableObjectBindingTypeDurableObjectNamespace DurableObjectBindingType = "durable_object_namespace"
)

func (r DurableObjectBindingType) IsKnown() bool {
	switch r {
	case DurableObjectBindingTypeDurableObjectNamespace:
		return true
	}
	return false
}

type KVNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type KVNamespaceBindingType `json:"type,required"`
	JSON kvNamespaceBindingJSON `json:"-"`
}

// kvNamespaceBindingJSON contains the JSON metadata for the struct
// [KVNamespaceBinding]
type kvNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KVNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r kvNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r KVNamespaceBinding) implementsWorkersBinding() {}

// The class of resource that the binding provides.
type KVNamespaceBindingType string

const (
	KVNamespaceBindingTypeKVNamespace KVNamespaceBindingType = "kv_namespace"
)

func (r KVNamespaceBindingType) IsKnown() bool {
	switch r {
	case KVNamespaceBindingTypeKVNamespace:
		return true
	}
	return false
}

type MigrationStepParam struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]MigrationStepRenamedClassParam] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]MigrationStepTransferredClassParam] `json:"transferred_classes"`
}

func (r MigrationStepParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MigrationStepRenamedClassParam struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r MigrationStepRenamedClassParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MigrationStepTransferredClassParam struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r MigrationStepTransferredClassParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MTLSCERTBinding struct {
	Certificate interface{} `json:"certificate,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type MTLSCERTBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string              `json:"certificate_id"`
	JSON          mtlscertBindingJSON `json:"-"`
}

// mtlscertBindingJSON contains the JSON metadata for the struct [MTLSCERTBinding]
type mtlscertBindingJSON struct {
	Certificate   apijson.Field
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *MTLSCERTBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mtlscertBindingJSON) RawJSON() string {
	return r.raw
}

func (r MTLSCERTBinding) implementsWorkersBinding() {}

// The class of resource that the binding provides.
type MTLSCERTBindingType string

const (
	MTLSCERTBindingTypeMTLSCertificate MTLSCERTBindingType = "mtls_certificate"
)

func (r MTLSCERTBindingType) IsKnown() bool {
	switch r {
	case MTLSCERTBindingTypeMTLSCertificate:
		return true
	}
	return false
}

type PlacementConfigurationParam struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode param.Field[PlacementConfigurationMode] `json:"mode"`
}

func (r PlacementConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type PlacementConfigurationMode string

const (
	PlacementConfigurationModeSmart PlacementConfigurationMode = "smart"
)

func (r PlacementConfigurationMode) IsKnown() bool {
	switch r {
	case PlacementConfigurationModeSmart:
		return true
	}
	return false
}

type R2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type R2BindingType `json:"type,required"`
	JSON r2BindingJSON `json:"-"`
}

// r2BindingJSON contains the JSON metadata for the struct [R2Binding]
type r2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2BindingJSON) RawJSON() string {
	return r.raw
}

func (r R2Binding) implementsWorkersBinding() {}

// The class of resource that the binding provides.
type R2BindingType string

const (
	R2BindingTypeR2Bucket R2BindingType = "r2_bucket"
)

func (r R2BindingType) IsKnown() bool {
	switch r {
	case R2BindingTypeR2Bucket:
		return true
	}
	return false
}

type ServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type ServiceBindingType `json:"type,required"`
	JSON serviceBindingJSON `json:"-"`
}

// serviceBindingJSON contains the JSON metadata for the struct [ServiceBinding]
type serviceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serviceBindingJSON) RawJSON() string {
	return r.raw
}

func (r ServiceBinding) implementsWorkersBinding() {}

// The class of resource that the binding provides.
type ServiceBindingType string

const (
	ServiceBindingTypeService ServiceBindingType = "service"
)

func (r ServiceBindingType) IsKnown() bool {
	switch r {
	case ServiceBindingTypeService:
		return true
	}
	return false
}

// A single set of migrations to apply.
type SingleStepMigrationParam struct {
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
	RenamedClasses param.Field[[]SingleStepMigrationRenamedClassParam] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]SingleStepMigrationTransferredClassParam] `json:"transferred_classes"`
}

func (r SingleStepMigrationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SingleStepMigrationParam) implementsWorkersScriptUpdateParamsVariant0MetadataMigrationsUnion() {
}

func (r SingleStepMigrationParam) implementsWorkersForPlatformsDispatchNamespaceScriptUpdateParamsVariant0MetadataMigrationsUnion() {
}

type SingleStepMigrationRenamedClassParam struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r SingleStepMigrationRenamedClassParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SingleStepMigrationTransferredClassParam struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r SingleStepMigrationTransferredClassParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SteppedMigrationParam struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]MigrationStepParam] `json:"steps"`
}

func (r SteppedMigrationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SteppedMigrationParam) implementsWorkersScriptUpdateParamsVariant0MetadataMigrationsUnion() {}

func (r SteppedMigrationParam) implementsWorkersForPlatformsDispatchNamespaceScriptUpdateParamsVariant0MetadataMigrationsUnion() {
}
