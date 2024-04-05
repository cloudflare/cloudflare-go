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
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// ScriptVersionSettingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewScriptVersionSettingService]
// method instead.
type ScriptVersionSettingService struct {
	Options []option.RequestOption
}

// NewScriptVersionSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewScriptVersionSettingService(opts ...option.RequestOption) (r *ScriptVersionSettingService) {
	r = &ScriptVersionSettingService{}
	r.Options = opts
	return
}

// Patch metadata or config, such as bindings or usage model
func (r *ScriptVersionSettingService) Edit(ctx context.Context, scriptName string, params ScriptVersionSettingEditParams, opts ...option.RequestOption) (res *SettingsItem, err error) {
	opts = append(r.Options[:], opts...)
	var env ScriptVersionSettingEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/settings", params.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get metadata and config, such as bindings or usage model
func (r *ScriptVersionSettingService) Get(ctx context.Context, scriptName string, query ScriptVersionSettingGetParams, opts ...option.RequestOption) (res *SettingsItem, err error) {
	opts = append(r.Options[:], opts...)
	var env ScriptVersionSettingGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/settings", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A binding to allow the Worker to communicate with resources
type BindingItem struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The class of resource that the binding provides.
	Type BindingItemType `json:"type,required"`
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
	CertificateID string          `json:"certificate_id"`
	Certificate   interface{}     `json:"certificate,required"`
	JSON          bindingItemJSON `json:"-"`
	union         BindingItemUnion
}

// bindingItemJSON contains the JSON metadata for the struct [BindingItem]
type bindingItemJSON struct {
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

func (r bindingItemJSON) RawJSON() string {
	return r.raw
}

func (r *BindingItem) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r BindingItem) AsUnion() BindingItemUnion {
	return r.union
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by [workers.BindingItemWorkersKVNamespaceBinding],
// [workers.BindingItemWorkersServiceBinding],
// [workers.BindingItemWorkersDoBinding], [workers.BindingItemWorkersR2Binding],
// [workers.BindingItemWorkersQueueBinding], [workers.BindingItemWorkersD1Binding],
// [workers.DispatchNamespaceBinding] or [workers.MTLSCERTBinding].
type BindingItemUnion interface {
	implementsWorkersBindingItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BindingItemUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BindingItemWorkersKVNamespaceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BindingItemWorkersServiceBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BindingItemWorkersDoBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BindingItemWorkersR2Binding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BindingItemWorkersQueueBinding{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BindingItemWorkersD1Binding{}),
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

type BindingItemWorkersKVNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type BindingItemWorkersKVNamespaceBindingType `json:"type,required"`
	JSON bindingItemWorkersKVNamespaceBindingJSON `json:"-"`
}

// bindingItemWorkersKVNamespaceBindingJSON contains the JSON metadata for the
// struct [BindingItemWorkersKVNamespaceBinding]
type bindingItemWorkersKVNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingItemWorkersKVNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersKVNamespaceBindingJSON) RawJSON() string {
	return r.raw
}

func (r BindingItemWorkersKVNamespaceBinding) implementsWorkersBindingItem() {}

// The class of resource that the binding provides.
type BindingItemWorkersKVNamespaceBindingType string

const (
	BindingItemWorkersKVNamespaceBindingTypeKVNamespace BindingItemWorkersKVNamespaceBindingType = "kv_namespace"
)

func (r BindingItemWorkersKVNamespaceBindingType) IsKnown() bool {
	switch r {
	case BindingItemWorkersKVNamespaceBindingTypeKVNamespace:
		return true
	}
	return false
}

type BindingItemWorkersServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type BindingItemWorkersServiceBindingType `json:"type,required"`
	JSON bindingItemWorkersServiceBindingJSON `json:"-"`
}

// bindingItemWorkersServiceBindingJSON contains the JSON metadata for the struct
// [BindingItemWorkersServiceBinding]
type bindingItemWorkersServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingItemWorkersServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersServiceBindingJSON) RawJSON() string {
	return r.raw
}

func (r BindingItemWorkersServiceBinding) implementsWorkersBindingItem() {}

// The class of resource that the binding provides.
type BindingItemWorkersServiceBindingType string

const (
	BindingItemWorkersServiceBindingTypeService BindingItemWorkersServiceBindingType = "service"
)

func (r BindingItemWorkersServiceBindingType) IsKnown() bool {
	switch r {
	case BindingItemWorkersServiceBindingTypeService:
		return true
	}
	return false
}

type BindingItemWorkersDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type BindingItemWorkersDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                          `json:"script_name"`
	JSON       bindingItemWorkersDoBindingJSON `json:"-"`
}

// bindingItemWorkersDoBindingJSON contains the JSON metadata for the struct
// [BindingItemWorkersDoBinding]
type bindingItemWorkersDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingItemWorkersDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersDoBindingJSON) RawJSON() string {
	return r.raw
}

func (r BindingItemWorkersDoBinding) implementsWorkersBindingItem() {}

// The class of resource that the binding provides.
type BindingItemWorkersDoBindingType string

const (
	BindingItemWorkersDoBindingTypeDurableObjectNamespace BindingItemWorkersDoBindingType = "durable_object_namespace"
)

func (r BindingItemWorkersDoBindingType) IsKnown() bool {
	switch r {
	case BindingItemWorkersDoBindingTypeDurableObjectNamespace:
		return true
	}
	return false
}

type BindingItemWorkersR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type BindingItemWorkersR2BindingType `json:"type,required"`
	JSON bindingItemWorkersR2BindingJSON `json:"-"`
}

// bindingItemWorkersR2BindingJSON contains the JSON metadata for the struct
// [BindingItemWorkersR2Binding]
type bindingItemWorkersR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingItemWorkersR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersR2BindingJSON) RawJSON() string {
	return r.raw
}

func (r BindingItemWorkersR2Binding) implementsWorkersBindingItem() {}

// The class of resource that the binding provides.
type BindingItemWorkersR2BindingType string

const (
	BindingItemWorkersR2BindingTypeR2Bucket BindingItemWorkersR2BindingType = "r2_bucket"
)

func (r BindingItemWorkersR2BindingType) IsKnown() bool {
	switch r {
	case BindingItemWorkersR2BindingTypeR2Bucket:
		return true
	}
	return false
}

type BindingItemWorkersQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type BindingItemWorkersQueueBindingType `json:"type,required"`
	JSON bindingItemWorkersQueueBindingJSON `json:"-"`
}

// bindingItemWorkersQueueBindingJSON contains the JSON metadata for the struct
// [BindingItemWorkersQueueBinding]
type bindingItemWorkersQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingItemWorkersQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersQueueBindingJSON) RawJSON() string {
	return r.raw
}

func (r BindingItemWorkersQueueBinding) implementsWorkersBindingItem() {}

// The class of resource that the binding provides.
type BindingItemWorkersQueueBindingType string

const (
	BindingItemWorkersQueueBindingTypeQueue BindingItemWorkersQueueBindingType = "queue"
)

func (r BindingItemWorkersQueueBindingType) IsKnown() bool {
	switch r {
	case BindingItemWorkersQueueBindingTypeQueue:
		return true
	}
	return false
}

type BindingItemWorkersD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type BindingItemWorkersD1BindingType `json:"type,required"`
	JSON bindingItemWorkersD1BindingJSON `json:"-"`
}

// bindingItemWorkersD1BindingJSON contains the JSON metadata for the struct
// [BindingItemWorkersD1Binding]
type bindingItemWorkersD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BindingItemWorkersD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bindingItemWorkersD1BindingJSON) RawJSON() string {
	return r.raw
}

func (r BindingItemWorkersD1Binding) implementsWorkersBindingItem() {}

// The class of resource that the binding provides.
type BindingItemWorkersD1BindingType string

const (
	BindingItemWorkersD1BindingTypeD1 BindingItemWorkersD1BindingType = "d1"
)

func (r BindingItemWorkersD1BindingType) IsKnown() bool {
	switch r {
	case BindingItemWorkersD1BindingTypeD1:
		return true
	}
	return false
}

// The class of resource that the binding provides.
type BindingItemType string

const (
	BindingItemTypeKVNamespace            BindingItemType = "kv_namespace"
	BindingItemTypeService                BindingItemType = "service"
	BindingItemTypeDurableObjectNamespace BindingItemType = "durable_object_namespace"
	BindingItemTypeR2Bucket               BindingItemType = "r2_bucket"
	BindingItemTypeQueue                  BindingItemType = "queue"
	BindingItemTypeD1                     BindingItemType = "d1"
	BindingItemTypeDispatchNamespace      BindingItemType = "dispatch_namespace"
	BindingItemTypeMTLSCertificate        BindingItemType = "mtls_certificate"
)

func (r BindingItemType) IsKnown() bool {
	switch r {
	case BindingItemTypeKVNamespace, BindingItemTypeService, BindingItemTypeDurableObjectNamespace, BindingItemTypeR2Bucket, BindingItemTypeQueue, BindingItemTypeD1, BindingItemTypeDispatchNamespace, BindingItemTypeMTLSCertificate:
		return true
	}
	return false
}

// A binding to allow the Worker to communicate with resources
type BindingItemParam struct {
	// The class of resource that the binding provides.
	Type param.Field[BindingItemType] `json:"type,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Name of Worker to bind to
	Service param.Field[string] `json:"service"`
	// The exported class name of the Durable Object
	ClassName param.Field[string] `json:"class_name"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName param.Field[string] `json:"script_name"`
	// R2 bucket to bind to
	BucketName param.Field[string] `json:"bucket_name"`
	// Name of the Queue to bind to
	QueueName param.Field[string] `json:"queue_name"`
	// ID of the D1 database to bind to
	ID param.Field[string] `json:"id"`
	// Namespace to bind to
	Namespace param.Field[string]      `json:"namespace"`
	Outbound  param.Field[interface{}] `json:"outbound,required"`
	// ID of the certificate to bind to
	CertificateID param.Field[string]      `json:"certificate_id"`
	Certificate   param.Field[interface{}] `json:"certificate,required"`
}

func (r BindingItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemParam) implementsWorkersBindingItemUnionParam() {}

// A binding to allow the Worker to communicate with resources
//
// Satisfied by [workers.BindingItemWorkersKVNamespaceBindingParam],
// [workers.BindingItemWorkersServiceBindingParam],
// [workers.BindingItemWorkersDoBindingParam],
// [workers.BindingItemWorkersR2BindingParam],
// [workers.BindingItemWorkersQueueBindingParam],
// [workers.BindingItemWorkersD1BindingParam],
// [workers.DispatchNamespaceBindingParam], [workers.MTLSCERTBindingParam],
// [BindingItemParam].
type BindingItemUnionParam interface {
	implementsWorkersBindingItemUnionParam()
}

type BindingItemWorkersKVNamespaceBindingParam struct {
	// The class of resource that the binding provides.
	Type param.Field[BindingItemWorkersKVNamespaceBindingType] `json:"type,required"`
}

func (r BindingItemWorkersKVNamespaceBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersKVNamespaceBindingParam) implementsWorkersBindingItemUnionParam() {}

type BindingItemWorkersServiceBindingParam struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// Name of Worker to bind to
	Service param.Field[string] `json:"service,required"`
	// The class of resource that the binding provides.
	Type param.Field[BindingItemWorkersServiceBindingType] `json:"type,required"`
}

func (r BindingItemWorkersServiceBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersServiceBindingParam) implementsWorkersBindingItemUnionParam() {}

type BindingItemWorkersDoBindingParam struct {
	// The exported class name of the Durable Object
	ClassName param.Field[string] `json:"class_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[BindingItemWorkersDoBindingType] `json:"type,required"`
	// The environment of the script_name to bind to
	Environment param.Field[string] `json:"environment"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName param.Field[string] `json:"script_name"`
}

func (r BindingItemWorkersDoBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersDoBindingParam) implementsWorkersBindingItemUnionParam() {}

type BindingItemWorkersR2BindingParam struct {
	// R2 bucket to bind to
	BucketName param.Field[string] `json:"bucket_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[BindingItemWorkersR2BindingType] `json:"type,required"`
}

func (r BindingItemWorkersR2BindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersR2BindingParam) implementsWorkersBindingItemUnionParam() {}

type BindingItemWorkersQueueBindingParam struct {
	// Name of the Queue to bind to
	QueueName param.Field[string] `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[BindingItemWorkersQueueBindingType] `json:"type,required"`
}

func (r BindingItemWorkersQueueBindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersQueueBindingParam) implementsWorkersBindingItemUnionParam() {}

type BindingItemWorkersD1BindingParam struct {
	// ID of the D1 database to bind to
	ID param.Field[string] `json:"id,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name param.Field[string] `json:"name,required"`
	// The class of resource that the binding provides.
	Type param.Field[BindingItemWorkersD1BindingType] `json:"type,required"`
}

func (r BindingItemWorkersD1BindingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BindingItemWorkersD1BindingParam) implementsWorkersBindingItemUnionParam() {}

type CompatibilityFlagsItem = string

type CompatibilityFlagsItemParam = string

type SettingsItem struct {
	// List of bindings attached to this Worker
	Bindings []BindingItem `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []CompatibilityFlagsItem `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations SettingsItemMigrations `json:"migrations"`
	Placement  PlacementConfiguration `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []ConsumerScript `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string           `json:"usage_model"`
	JSON       settingsItemJSON `json:"-"`
}

// settingsItemJSON contains the JSON metadata for the struct [SettingsItem]
type settingsItemJSON struct {
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

func (r *SettingsItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingsItemJSON) RawJSON() string {
	return r.raw
}

// Migrations to apply for Durable Objects associated with this Worker.
type SettingsItemMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag             string                     `json:"old_tag"`
	DeletedClasses     interface{}                `json:"deleted_classes,required"`
	NewClasses         interface{}                `json:"new_classes,required"`
	RenamedClasses     interface{}                `json:"renamed_classes,required"`
	TransferredClasses interface{}                `json:"transferred_classes,required"`
	Steps              interface{}                `json:"steps,required"`
	JSON               settingsItemMigrationsJSON `json:"-"`
	union              SettingsItemMigrationsUnion
}

// settingsItemMigrationsJSON contains the JSON metadata for the struct
// [SettingsItemMigrations]
type settingsItemMigrationsJSON struct {
	NewTag             apijson.Field
	OldTag             apijson.Field
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	Steps              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r settingsItemMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r *SettingsItemMigrations) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r SettingsItemMigrations) AsUnion() SettingsItemMigrationsUnion {
	return r.union
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by [workers.SingleStepMigration] or [workers.SteppedMigration].
type SettingsItemMigrationsUnion interface {
	implementsWorkersSettingsItemMigrations()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SettingsItemMigrationsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SingleStepMigration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SteppedMigration{}),
		},
	)
}

type SettingsItemParam struct {
	// List of bindings attached to this Worker
	Bindings param.Field[[]BindingItemUnionParam] `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags param.Field[[]CompatibilityFlagsItemParam] `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations param.Field[SettingsItemMigrationsUnionParam] `json:"migrations"`
	Placement  param.Field[PlacementConfigurationParam]      `json:"placement"`
	// Tags to help you manage your Workers
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]ConsumerScriptParam] `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel param.Field[string] `json:"usage_model"`
}

func (r SettingsItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Migrations to apply for Durable Objects associated with this Worker.
type SettingsItemMigrationsParam struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag             param.Field[string]      `json:"old_tag"`
	DeletedClasses     param.Field[interface{}] `json:"deleted_classes,required"`
	NewClasses         param.Field[interface{}] `json:"new_classes,required"`
	RenamedClasses     param.Field[interface{}] `json:"renamed_classes,required"`
	TransferredClasses param.Field[interface{}] `json:"transferred_classes,required"`
	Steps              param.Field[interface{}] `json:"steps,required"`
}

func (r SettingsItemMigrationsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingsItemMigrationsParam) implementsWorkersSettingsItemMigrationsUnionParam() {}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by [workers.SingleStepMigrationParam],
// [workers.SteppedMigrationParam], [SettingsItemMigrationsParam].
type SettingsItemMigrationsUnionParam interface {
	implementsWorkersSettingsItemMigrationsUnionParam()
}

type ScriptVersionSettingEditParams struct {
	// Identifier
	AccountID param.Field[string]            `path:"account_id,required"`
	Settings  param.Field[SettingsItemParam] `json:"settings"`
}

func (r ScriptVersionSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptVersionSettingEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   SettingsItem                                              `json:"result,required"`
	// Whether the API call was successful
	Success ScriptVersionSettingEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptVersionSettingEditResponseEnvelopeJSON    `json:"-"`
}

// scriptVersionSettingEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [ScriptVersionSettingEditResponseEnvelope]
type scriptVersionSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptVersionSettingEditResponseEnvelopeSuccess bool

const (
	ScriptVersionSettingEditResponseEnvelopeSuccessTrue ScriptVersionSettingEditResponseEnvelopeSuccess = true
)

func (r ScriptVersionSettingEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptVersionSettingEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScriptVersionSettingGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScriptVersionSettingGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   SettingsItem                                              `json:"result,required"`
	// Whether the API call was successful
	Success ScriptVersionSettingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    scriptVersionSettingGetResponseEnvelopeJSON    `json:"-"`
}

// scriptVersionSettingGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ScriptVersionSettingGetResponseEnvelope]
type scriptVersionSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScriptVersionSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptVersionSettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ScriptVersionSettingGetResponseEnvelopeSuccess bool

const (
	ScriptVersionSettingGetResponseEnvelopeSuccessTrue ScriptVersionSettingGetResponseEnvelopeSuccess = true
)

func (r ScriptVersionSettingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScriptVersionSettingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
