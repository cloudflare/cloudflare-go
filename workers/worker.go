// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"reflect"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/tidwall/gjson"
)

// WorkerService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkerService] method instead.
type WorkerService struct {
	Options         []option.RequestOption
	Routes          *RouteService
	Assets          *AssetService
	Scripts         *ScriptService
	AccountSettings *AccountSettingService
	Domains         *DomainService
	Subdomains      *SubdomainService
}

// NewWorkerService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWorkerService(opts ...option.RequestOption) (r *WorkerService) {
	r = &WorkerService{}
	r.Options = opts
	r.Routes = NewRouteService(opts...)
	r.Assets = NewAssetService(opts...)
	r.Scripts = NewScriptService(opts...)
	r.AccountSettings = NewAccountSettingService(opts...)
	r.Domains = NewDomainService(opts...)
	r.Subdomains = NewSubdomainService(opts...)
	return
}

// A binding to allow the Worker to communicate with resources
type Binding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type BindingType `json:"type,required"`
	// ID of the D1 database to bind to
	ID string `json:"id"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding"`
	// R2 bucket to bind to
	BucketName string `json:"bucket_name"`
	// ID of the certificate to bind to
	CertificateID string `json:"certificate_id"`
	// The exported class name of the Durable Object
	ClassName string `json:"class_name"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Namespace to bind to
	Namespace string `json:"namespace"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// This field can have the runtime type of [DispatchNamespaceBindingOutbound].
	Outbound interface{} `json:"outbound"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string `json:"script_name"`
	// Name of Worker to bind to
	Service string      `json:"service"`
	JSON    bindingJSON `json:"-"`
	union   BindingUnion
}

// bindingJSON contains the JSON metadata for the struct [Binding]
type bindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	ID            apijson.Field
	Binding       apijson.Field
	BucketName    apijson.Field
	CertificateID apijson.Field
	ClassName     apijson.Field
	Environment   apijson.Field
	Namespace     apijson.Field
	NamespaceID   apijson.Field
	Outbound      apijson.Field
	QueueName     apijson.Field
	ScriptName    apijson.Field
	Service       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r bindingJSON) RawJSON() string {
	return r.raw
}

func (r *Binding) UnmarshalJSON(data []byte) (err error) {
	*r = Binding{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [BindingUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [workers.KVNamespaceBinding],
// [workers.ServiceBinding], [workers.DurableObjectBinding], [workers.R2Binding],
// [workers.BindingWorkersQueueBinding], [workers.D1Binding],
// [workers.DispatchNamespaceBinding], [workers.MTLSCERTBinding],
// [workers.BindingWorkersAssetsBinding].
func (r Binding) AsUnion() BindingUnion {
	return r.union
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by [workers.KVNamespaceBinding], [workers.ServiceBinding],
// [workers.DurableObjectBinding], [workers.R2Binding],
// [workers.BindingWorkersQueueBinding], [workers.D1Binding],
// [workers.DispatchNamespaceBinding], [workers.MTLSCERTBinding] or
// [workers.BindingWorkersAssetsBinding].
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
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BindingWorkersAssetsBinding{}),
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

type BindingWorkersAssetsBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type BindingWorkersAssetsBindingType `json:"type,required"`
	JSON bindingWorkersAssetsBindingJSON `json:"-"`
}

// bindingWorkersAssetsBindingJSON contains the JSON metadata for the struct
// [BindingWorkersAssetsBinding]
type bindingWorkersAssetsBindingJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingWorkersAssetsBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingWorkersAssetsBindingJSON) RawJSON() string {
	return r.raw
}

func (r BindingWorkersAssetsBinding) implementsWorkersBinding() {}

// The class of resource that the binding provides.
type BindingWorkersAssetsBindingType string

const (
	BindingWorkersAssetsBindingTypeAssets BindingWorkersAssetsBindingType = "assets"
)

func (r BindingWorkersAssetsBindingType) IsKnown() bool {
	switch r {
	case BindingWorkersAssetsBindingTypeAssets:
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
	BindingTypeAssets                 BindingType = "assets"
)

func (r BindingType) IsKnown() bool {
	switch r {
	case BindingTypeKVNamespace, BindingTypeService, BindingTypeDurableObjectNamespace, BindingTypeR2Bucket, BindingTypeQueue, BindingTypeD1, BindingTypeDispatchNamespace, BindingTypeMTLSCertificate, BindingTypeAssets:
		return true
	}
	return false
}

// A binding to allow the Worker to communicate with resources
type BindingParam struct {
	// The class of resource that the binding provides.
	Type param.Field[BindingType] `json:"type,required"`
	// ID of the D1 database to bind to
	ID param.Field[string] `json:"id"`
	// R2 bucket to bind to
	BucketName param.Field[string] `json:"bucket_name"`
	// ID of the certificate to bind to
	CertificateID param.Field[string] `json:"certificate_id"`
	// The exported class name of the Durable Object
	ClassName param.Field[string] `json:"class_name"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Namespace to bind to
	Namespace param.Field[string]      `json:"namespace"`
	Outbound  param.Field[interface{}] `json:"outbound"`
	// Name of the Queue to bind to
	QueueName param.Field[string] `json:"queue_name"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName param.Field[string] `json:"script_name"`
	// Name of Worker to bind to
	Service param.Field[string] `json:"service"`
}

func (r BindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingParam) implementsWorkersBindingUnionParam() {}

// A binding to allow the Worker to communicate with resources
//
// Satisfied by [workers.KVNamespaceBindingParam], [workers.ServiceBindingParam],
// [workers.DurableObjectBindingParam], [workers.R2BindingParam],
// [workers.BindingWorkersQueueBindingParam], [workers.D1BindingParam],
// [workers.DispatchNamespaceBindingParam], [workers.MTLSCERTBindingParam],
// [workers.BindingWorkersAssetsBindingParam], [BindingParam].
type BindingUnionParam interface {
	implementsWorkersBindingUnionParam()
}

type BindingWorkersQueueBindingParam struct {
	// Name of the Queue to bind to
	QueueName param.Field[string] `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[BindingWorkersQueueBindingType] `json:"type,required"`
}

func (r BindingWorkersQueueBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingWorkersQueueBindingParam) implementsWorkersBindingUnionParam() {}

type BindingWorkersAssetsBindingParam struct {
	// The class of resource that the binding provides.
	Type param.Field[BindingWorkersAssetsBindingType] `json:"type,required"`
}

func (r BindingWorkersAssetsBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingWorkersAssetsBindingParam) implementsWorkersBindingUnionParam() {}

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

type D1BindingParam struct {
	// ID of the D1 database to bind to
	ID param.Field[string] `json:"id,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name param.Field[string] `json:"name,required"`
	// The class of resource that the binding provides.
	Type param.Field[D1BindingType] `json:"type,required"`
}

func (r D1BindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r D1BindingParam) implementsWorkersBindingUnionParam() {}

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

type DispatchNamespaceBindingParam struct {
	// Namespace to bind to
	Namespace param.Field[string] `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type param.Field[DispatchNamespaceBindingType] `json:"type,required"`
	// Outbound worker
	Outbound param.Field[DispatchNamespaceBindingOutboundParam] `json:"outbound"`
}

func (r DispatchNamespaceBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceBindingParam) implementsWorkersBindingUnionParam() {}

// Outbound worker
type DispatchNamespaceBindingOutboundParam struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params param.Field[[]string] `json:"params"`
	// Outbound worker
	Worker param.Field[DispatchNamespaceBindingOutboundWorkerParam] `json:"worker"`
}

func (r DispatchNamespaceBindingOutboundParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker
type DispatchNamespaceBindingOutboundWorkerParam struct {
	// Environment of the outbound worker
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker
	Service param.Field[string] `json:"service"`
}

func (r DispatchNamespaceBindingOutboundWorkerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

type DurableObjectBindingParam struct {
	// The exported class name of the Durable Object
	ClassName param.Field[string] `json:"class_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[DurableObjectBindingType] `json:"type,required"`
	// The environment of the script_name to bind to
	Environment param.Field[string] `json:"environment"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName param.Field[string] `json:"script_name"`
}

func (r DurableObjectBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DurableObjectBindingParam) implementsWorkersBindingUnionParam() {}

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

type KVNamespaceBindingParam struct {
	// The class of resource that the binding provides.
	Type param.Field[KVNamespaceBindingType] `json:"type,required"`
}

func (r KVNamespaceBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r KVNamespaceBindingParam) implementsWorkersBindingUnionParam() {}

type MigrationStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes to create Durable Object namespaces with SQLite from.
	NewSqliteClasses []string `json:"new_sqlite_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []MigrationStepRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []MigrationStepTransferredClass `json:"transferred_classes"`
	JSON               migrationStepJSON               `json:"-"`
}

// migrationStepJSON contains the JSON metadata for the struct [MigrationStep]
type migrationStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewSqliteClasses   apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *MigrationStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r migrationStepJSON) RawJSON() string {
	return r.raw
}

type MigrationStepRenamedClass struct {
	From string                        `json:"from"`
	To   string                        `json:"to"`
	JSON migrationStepRenamedClassJSON `json:"-"`
}

// migrationStepRenamedClassJSON contains the JSON metadata for the struct
// [MigrationStepRenamedClass]
type migrationStepRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MigrationStepRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r migrationStepRenamedClassJSON) RawJSON() string {
	return r.raw
}

type MigrationStepTransferredClass struct {
	From       string                            `json:"from"`
	FromScript string                            `json:"from_script"`
	To         string                            `json:"to"`
	JSON       migrationStepTransferredClassJSON `json:"-"`
}

// migrationStepTransferredClassJSON contains the JSON metadata for the struct
// [MigrationStepTransferredClass]
type migrationStepTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MigrationStepTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r migrationStepTransferredClassJSON) RawJSON() string {
	return r.raw
}

type MigrationStepParam struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// A list of classes to create Durable Object namespaces with SQLite from.
	NewSqliteClasses param.Field[[]string] `json:"new_sqlite_classes"`
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
	// ID of the certificate to bind to
	CertificateID string `json:"certificate_id,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type MTLSCERTBindingType `json:"type,required"`
	JSON mtlscertBindingJSON `json:"-"`
}

// mtlscertBindingJSON contains the JSON metadata for the struct [MTLSCERTBinding]
type mtlscertBindingJSON struct {
	CertificateID apijson.Field
	Name          apijson.Field
	Type          apijson.Field
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

type MTLSCERTBindingParam struct {
	// ID of the certificate to bind to
	CertificateID param.Field[string] `json:"certificate_id,required"`
	// The class of resource that the binding provides.
	Type param.Field[MTLSCERTBindingType] `json:"type,required"`
}

func (r MTLSCERTBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MTLSCERTBindingParam) implementsWorkersBindingUnionParam() {}

type PlacementConfiguration struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode PlacementConfigurationMode `json:"mode"`
	JSON placementConfigurationJSON `json:"-"`
}

// placementConfigurationJSON contains the JSON metadata for the struct
// [PlacementConfiguration]
type placementConfigurationJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlacementConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r placementConfigurationJSON) RawJSON() string {
	return r.raw
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

type PlacementConfigurationParam struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode param.Field[PlacementConfigurationMode] `json:"mode"`
}

func (r PlacementConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

type R2BindingParam struct {
	// R2 bucket to bind to
	BucketName param.Field[string] `json:"bucket_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[R2BindingType] `json:"type,required"`
}

func (r R2BindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r R2BindingParam) implementsWorkersBindingUnionParam() {}

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

type ServiceBindingParam struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// Name of Worker to bind to
	Service param.Field[string] `json:"service,required"`
	// The class of resource that the binding provides.
	Type param.Field[ServiceBindingType] `json:"type,required"`
}

func (r ServiceBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ServiceBindingParam) implementsWorkersBindingUnionParam() {}

// A single set of migrations to apply.
type SingleStepMigration struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes to create Durable Object namespaces with SQLite from.
	NewSqliteClasses []string `json:"new_sqlite_classes"`
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []SingleStepMigrationRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []SingleStepMigrationTransferredClass `json:"transferred_classes"`
	JSON               singleStepMigrationJSON               `json:"-"`
}

// singleStepMigrationJSON contains the JSON metadata for the struct
// [SingleStepMigration]
type singleStepMigrationJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewSqliteClasses   apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SingleStepMigration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleStepMigrationJSON) RawJSON() string {
	return r.raw
}

func (r SingleStepMigration) ImplementsWorkersForPlatformsDispatchNamespaceScriptSettingEditResponseMigrations() {
}

func (r SingleStepMigration) ImplementsWorkersForPlatformsDispatchNamespaceScriptSettingGetResponseMigrations() {
}

type SingleStepMigrationRenamedClass struct {
	From string                              `json:"from"`
	To   string                              `json:"to"`
	JSON singleStepMigrationRenamedClassJSON `json:"-"`
}

// singleStepMigrationRenamedClassJSON contains the JSON metadata for the struct
// [SingleStepMigrationRenamedClass]
type singleStepMigrationRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleStepMigrationRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleStepMigrationRenamedClassJSON) RawJSON() string {
	return r.raw
}

type SingleStepMigrationTransferredClass struct {
	From       string                                  `json:"from"`
	FromScript string                                  `json:"from_script"`
	To         string                                  `json:"to"`
	JSON       singleStepMigrationTransferredClassJSON `json:"-"`
}

// singleStepMigrationTransferredClassJSON contains the JSON metadata for the
// struct [SingleStepMigrationTransferredClass]
type singleStepMigrationTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleStepMigrationTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleStepMigrationTransferredClassJSON) RawJSON() string {
	return r.raw
}

// A single set of migrations to apply.
type SingleStepMigrationParam struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// A list of classes to create Durable Object namespaces with SQLite from.
	NewSqliteClasses param.Field[[]string] `json:"new_sqlite_classes"`
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

func (r SingleStepMigrationParam) implementsWorkersScriptUpdateParamsBodyMetadataMetadataMigrationsUnion() {
}

func (r SingleStepMigrationParam) ImplementsWorkersForPlatformsDispatchNamespaceScriptUpdateParamsBodyMetadataMetadataMigrationsUnion() {
}

func (r SingleStepMigrationParam) ImplementsWorkersForPlatformsDispatchNamespaceScriptSettingEditParamsSettingsMigrationsUnion() {
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

// JSON encoded metadata about the uploaded parts and Worker configuration.
type WorkerMetadataParam struct {
	// Name of the part in the multipart request that contains the script (e.g. the
	// file adding a listener to the `fetch` event). Indicates a
	// `service worker syntax` Worker.
	BodyPart param.Field[string] `json:"body_part"`
	// Name of the part in the multipart request that contains the main module (e.g.
	// the file exporting a `fetch` handler). Indicates a `module syntax` Worker.
	MainModule param.Field[string] `json:"main_module"`
}

func (r WorkerMetadataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
