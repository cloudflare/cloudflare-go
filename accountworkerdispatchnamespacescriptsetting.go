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

// AccountWorkerDispatchNamespaceScriptSettingService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkerDispatchNamespaceScriptSettingService] method instead.
type AccountWorkerDispatchNamespaceScriptSettingService struct {
	Options []option.RequestOption
}

// NewAccountWorkerDispatchNamespaceScriptSettingService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewAccountWorkerDispatchNamespaceScriptSettingService(opts ...option.RequestOption) (r *AccountWorkerDispatchNamespaceScriptSettingService) {
	r = &AccountWorkerDispatchNamespaceScriptSettingService{}
	r.Options = opts
	return
}

// Get script settings from a script uploaded to a Workers for Platforms namespace.
func (r *AccountWorkerDispatchNamespaceScriptSettingService) Get(ctx context.Context, accountIdentifier string, dispatchNamespace string, scriptName string, opts ...option.RequestOption) (res *AccountWorkerDispatchNamespaceScriptSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/settings", accountIdentifier, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Patch script metadata, such as bindings
func (r *AccountWorkerDispatchNamespaceScriptSettingService) Update(ctx context.Context, accountIdentifier string, dispatchNamespace string, scriptName string, body AccountWorkerDispatchNamespaceScriptSettingUpdateParams, opts ...option.RequestOption) (res *AccountWorkerDispatchNamespaceScriptSettingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/settings", accountIdentifier, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type AccountWorkerDispatchNamespaceScriptSettingGetResponse struct {
	// List of bindings attached to this Worker
	Bindings []AccountWorkerDispatchNamespaceScriptSettingGetResponseBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrations `json:"migrations"`
	Placement  AccountWorkerDispatchNamespaceScriptSettingGetResponsePlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []AccountWorkerDispatchNamespaceScriptSettingGetResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                                     `json:"usage_model"`
	JSON       accountWorkerDispatchNamespaceScriptSettingGetResponseJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingGetResponseJSON contains the JSON
// metadata for the struct [AccountWorkerDispatchNamespaceScriptSettingGetResponse]
type accountWorkerDispatchNamespaceScriptSettingGetResponseJSON struct {
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

func (r *AccountWorkerDispatchNamespaceScriptSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uKvNamespaceBinding],
// [AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uServiceBinding],
// [AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDoBinding],
// [AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uR2Binding],
// [AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uQueueBinding],
// [AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uD1Binding],
// [AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBinding]
// or
// [AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uMtlsCertBinding].
type AccountWorkerDispatchNamespaceScriptSettingGetResponseBinding interface {
	implementsAccountWorkerDispatchNamespaceScriptSettingGetResponseBinding()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountWorkerDispatchNamespaceScriptSettingGetResponseBinding)(nil)).Elem(), "")
}

type AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uKvNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uKvNamespaceBindingType `json:"type,required"`
	JSON accountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uKvNamespaceBindingJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uKvNamespaceBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uKvNamespaceBinding]
type accountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uKvNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uKvNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uKvNamespaceBinding) implementsAccountWorkerDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uKvNamespaceBindingType string

const (
	AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uKvNamespaceBindingTypeKvNamespace AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uKvNamespaceBindingType = "kv_namespace"
)

type AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uServiceBindingType `json:"type,required"`
	JSON accountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uServiceBindingJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uServiceBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uServiceBinding]
type accountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uServiceBinding) implementsAccountWorkerDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uServiceBindingType string

const (
	AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uServiceBindingTypeService AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uServiceBindingType = "service"
)

type AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                                              `json:"script_name"`
	JSON       accountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDoBindingJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDoBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDoBinding]
type accountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDoBinding) implementsAccountWorkerDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDoBindingType string

const (
	AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDoBindingTypeDurableObjectNamespace AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDoBindingType = "durable_object_namespace"
)

type AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uR2BindingType `json:"type,required"`
	JSON accountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uR2BindingJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uR2BindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uR2Binding]
type accountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uR2Binding) implementsAccountWorkerDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uR2BindingType string

const (
	AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uR2BindingTypeR2Bucket AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uR2BindingType = "r2_bucket"
)

type AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uQueueBindingType `json:"type,required"`
	JSON accountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uQueueBindingJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uQueueBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uQueueBinding]
type accountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uQueueBinding) implementsAccountWorkerDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uQueueBindingType string

const (
	AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uQueueBindingTypeQueue AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uQueueBindingType = "queue"
)

type AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uD1BindingType `json:"type,required"`
	JSON accountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uD1BindingJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uD1BindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uD1Binding]
type accountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uD1Binding) implementsAccountWorkerDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uD1BindingType string

const (
	AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uD1BindingTypeD1 AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uD1BindingType = "d1"
)

type AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     accountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingJSON     `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBinding]
type accountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBinding) implementsAccountWorkerDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingType string

const (
	AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingTypeDispatchNamespace AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   accountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutbound]
type accountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Outbound worker
type AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                                           `json:"service"`
	JSON    accountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker]
type accountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uMtlsCertBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uMtlsCertBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                                                    `json:"certificate_id"`
	JSON          accountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uMtlsCertBindingJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uMtlsCertBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uMtlsCertBinding]
type accountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uMtlsCertBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uMtlsCertBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uMtlsCertBinding) implementsAccountWorkerDispatchNamespaceScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uMtlsCertBindingType string

const (
	AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uMtlsCertBindingTypeMtlsCertificate AccountWorkerDispatchNamespaceScriptSettingGetResponseBindingsAvYbsl2uMtlsCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrations]
// or
// [AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrations].
type AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrations interface {
	implementsAccountWorkerDispatchNamespaceScriptSettingGetResponseMigrations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrations)(nil)).Elem(), "")
}

// A single set of migrations to apply.
type AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrations struct {
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
	RenamedClasses []AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               accountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsJSON               `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrations]
type accountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrations) implementsAccountWorkerDispatchNamespaceScriptSettingGetResponseMigrations() {
}

type AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClass struct {
	From string                                                                                                       `json:"from"`
	To   string                                                                                                       `json:"to"`
	JSON accountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClass]
type accountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClass struct {
	From       string                                                                                                           `json:"from"`
	FromScript string                                                                                                           `json:"from_script"`
	To         string                                                                                                           `json:"to"`
	JSON       accountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClass]
type accountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStep `json:"steps"`
	JSON  accountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsJSON   `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrations]
type accountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrations) implementsAccountWorkerDispatchNamespaceScriptSettingGetResponseMigrations() {
}

type AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               accountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepJSON                `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStep]
type accountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass struct {
	From string                                                                                                         `json:"from"`
	To   string                                                                                                         `json:"to"`
	JSON accountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass]
type accountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                                                             `json:"from"`
	FromScript string                                                                                                             `json:"from_script"`
	To         string                                                                                                             `json:"to"`
	JSON       accountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass]
type accountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDispatchNamespaceScriptSettingGetResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode AccountWorkerDispatchNamespaceScriptSettingGetResponsePlacementMode `json:"mode"`
	JSON accountWorkerDispatchNamespaceScriptSettingGetResponsePlacementJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingGetResponsePlacementJSON contains the
// JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingGetResponsePlacement]
type accountWorkerDispatchNamespaceScriptSettingGetResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingGetResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type AccountWorkerDispatchNamespaceScriptSettingGetResponsePlacementMode string

const (
	AccountWorkerDispatchNamespaceScriptSettingGetResponsePlacementModeSmart AccountWorkerDispatchNamespaceScriptSettingGetResponsePlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type AccountWorkerDispatchNamespaceScriptSettingGetResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                                 `json:"namespace"`
	JSON      accountWorkerDispatchNamespaceScriptSettingGetResponseTailConsumerJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingGetResponseTailConsumerJSON contains
// the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingGetResponseTailConsumer]
type accountWorkerDispatchNamespaceScriptSettingGetResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingGetResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDispatchNamespaceScriptSettingUpdateResponse struct {
	// List of bindings attached to this Worker
	Bindings []AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrations `json:"migrations"`
	Placement  AccountWorkerDispatchNamespaceScriptSettingUpdateResponsePlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []AccountWorkerDispatchNamespaceScriptSettingUpdateResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                                        `json:"usage_model"`
	JSON       accountWorkerDispatchNamespaceScriptSettingUpdateResponseJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingUpdateResponseJSON contains the JSON
// metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingUpdateResponse]
type accountWorkerDispatchNamespaceScriptSettingUpdateResponseJSON struct {
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

func (r *AccountWorkerDispatchNamespaceScriptSettingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBinding],
// [AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uServiceBinding],
// [AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDoBinding],
// [AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uR2Binding],
// [AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uQueueBinding],
// [AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uD1Binding],
// [AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBinding]
// or
// [AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uMtlsCertBinding].
type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBinding interface {
	implementsAccountWorkerDispatchNamespaceScriptSettingUpdateResponseBinding()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBinding)(nil)).Elem(), "")
}

type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBindingType `json:"type,required"`
	JSON accountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBindingJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBinding]
type accountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBinding) implementsAccountWorkerDispatchNamespaceScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBindingType string

const (
	AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBindingTypeKvNamespace AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBindingType = "kv_namespace"
)

type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uServiceBindingType `json:"type,required"`
	JSON accountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uServiceBindingJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uServiceBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uServiceBinding]
type accountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uServiceBinding) implementsAccountWorkerDispatchNamespaceScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uServiceBindingType string

const (
	AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uServiceBindingTypeService AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uServiceBindingType = "service"
)

type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                                                 `json:"script_name"`
	JSON       accountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDoBindingJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDoBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDoBinding]
type accountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDoBinding) implementsAccountWorkerDispatchNamespaceScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDoBindingType string

const (
	AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDoBindingTypeDurableObjectNamespace AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDoBindingType = "durable_object_namespace"
)

type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uR2BindingType `json:"type,required"`
	JSON accountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uR2BindingJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uR2BindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uR2Binding]
type accountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uR2Binding) implementsAccountWorkerDispatchNamespaceScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uR2BindingType string

const (
	AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uR2BindingTypeR2Bucket AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uR2BindingType = "r2_bucket"
)

type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uQueueBindingType `json:"type,required"`
	JSON accountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uQueueBindingJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uQueueBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uQueueBinding]
type accountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uQueueBinding) implementsAccountWorkerDispatchNamespaceScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uQueueBindingType string

const (
	AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uQueueBindingTypeQueue AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uQueueBindingType = "queue"
)

type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uD1BindingType `json:"type,required"`
	JSON accountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uD1BindingJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uD1BindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uD1Binding]
type accountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uD1Binding) implementsAccountWorkerDispatchNamespaceScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uD1BindingType string

const (
	AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uD1BindingTypeD1 AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uD1BindingType = "d1"
)

type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     accountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingJSON     `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBinding]
type accountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBinding) implementsAccountWorkerDispatchNamespaceScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingType string

const (
	AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingTypeDispatchNamespace AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   accountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutbound]
type accountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Outbound worker
type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                                              `json:"service"`
	JSON    accountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker]
type accountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uMtlsCertBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uMtlsCertBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                                                       `json:"certificate_id"`
	JSON          accountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uMtlsCertBindingJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uMtlsCertBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uMtlsCertBinding]
type accountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uMtlsCertBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uMtlsCertBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uMtlsCertBinding) implementsAccountWorkerDispatchNamespaceScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uMtlsCertBindingType string

const (
	AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uMtlsCertBindingTypeMtlsCertificate AccountWorkerDispatchNamespaceScriptSettingUpdateResponseBindingsAvYbsl2uMtlsCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrations]
// or
// [AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrations].
type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrations interface {
	implementsAccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrations)(nil)).Elem(), "")
}

// A single set of migrations to apply.
type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrations struct {
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
	RenamedClasses []AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               accountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsJSON               `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrations]
type accountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrations) implementsAccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrations() {
}

type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClass struct {
	From string                                                                                                          `json:"from"`
	To   string                                                                                                          `json:"to"`
	JSON accountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClass]
type accountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClass struct {
	From       string                                                                                                              `json:"from"`
	FromScript string                                                                                                              `json:"from_script"`
	To         string                                                                                                              `json:"to"`
	JSON       accountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClass]
type accountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStep `json:"steps"`
	JSON  accountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsJSON   `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrations]
type accountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrations) implementsAccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrations() {
}

type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               accountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepJSON                `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStep]
type accountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass struct {
	From string                                                                                                            `json:"from"`
	To   string                                                                                                            `json:"to"`
	JSON accountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass]
type accountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                                                                `json:"from"`
	FromScript string                                                                                                                `json:"from_script"`
	To         string                                                                                                                `json:"to"`
	JSON       accountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass]
type accountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDispatchNamespaceScriptSettingUpdateResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode AccountWorkerDispatchNamespaceScriptSettingUpdateResponsePlacementMode `json:"mode"`
	JSON accountWorkerDispatchNamespaceScriptSettingUpdateResponsePlacementJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingUpdateResponsePlacementJSON contains
// the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingUpdateResponsePlacement]
type accountWorkerDispatchNamespaceScriptSettingUpdateResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingUpdateResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type AccountWorkerDispatchNamespaceScriptSettingUpdateResponsePlacementMode string

const (
	AccountWorkerDispatchNamespaceScriptSettingUpdateResponsePlacementModeSmart AccountWorkerDispatchNamespaceScriptSettingUpdateResponsePlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type AccountWorkerDispatchNamespaceScriptSettingUpdateResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                                    `json:"namespace"`
	JSON      accountWorkerDispatchNamespaceScriptSettingUpdateResponseTailConsumerJSON `json:"-"`
}

// accountWorkerDispatchNamespaceScriptSettingUpdateResponseTailConsumerJSON
// contains the JSON metadata for the struct
// [AccountWorkerDispatchNamespaceScriptSettingUpdateResponseTailConsumer]
type accountWorkerDispatchNamespaceScriptSettingUpdateResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDispatchNamespaceScriptSettingUpdateResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDispatchNamespaceScriptSettingUpdateParams struct {
	// List of bindings attached to this Worker
	Bindings param.Field[[]AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBinding] `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations param.Field[AccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrations] `json:"migrations"`
	Placement  param.Field[AccountWorkerDispatchNamespaceScriptSettingUpdateParamsPlacement]  `json:"placement"`
	// Tags to help you manage your Workers
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]AccountWorkerDispatchNamespaceScriptSettingUpdateParamsTailConsumer] `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel param.Field[string] `json:"usage_model"`
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A binding to allow the Worker to communicate with resources
//
// Satisfied by
// [AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBinding],
// [AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uServiceBinding],
// [AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uDoBinding],
// [AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uR2Binding],
// [AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uQueueBinding],
// [AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uD1Binding],
// [AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uDispatchNamespaceBinding],
// [AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uMtlsCertBinding].
type AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBinding interface {
	implementsAccountWorkerDispatchNamespaceScriptSettingUpdateParamsBinding()
}

type AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBinding struct {
	// The class of resource that the binding provides.
	Type param.Field[AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBindingType] `json:"type,required"`
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBinding) implementsAccountWorkerDispatchNamespaceScriptSettingUpdateParamsBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBindingType string

const (
	AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBindingTypeKvNamespace AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBindingType = "kv_namespace"
)

type AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// Name of Worker to bind to
	Service param.Field[string] `json:"service,required"`
	// The class of resource that the binding provides.
	Type param.Field[AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uServiceBindingType] `json:"type,required"`
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uServiceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uServiceBinding) implementsAccountWorkerDispatchNamespaceScriptSettingUpdateParamsBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uServiceBindingType string

const (
	AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uServiceBindingTypeService AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uServiceBindingType = "service"
)

type AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uDoBinding struct {
	// The exported class name of the Durable Object
	ClassName param.Field[string] `json:"class_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uDoBindingType] `json:"type,required"`
	// The environment of the script_name to bind to
	Environment param.Field[string] `json:"environment"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName param.Field[string] `json:"script_name"`
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uDoBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uDoBinding) implementsAccountWorkerDispatchNamespaceScriptSettingUpdateParamsBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uDoBindingType string

const (
	AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uDoBindingTypeDurableObjectNamespace AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uDoBindingType = "durable_object_namespace"
)

type AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uR2Binding struct {
	// R2 bucket to bind to
	BucketName param.Field[string] `json:"bucket_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uR2BindingType] `json:"type,required"`
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uR2Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uR2Binding) implementsAccountWorkerDispatchNamespaceScriptSettingUpdateParamsBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uR2BindingType string

const (
	AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uR2BindingTypeR2Bucket AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uR2BindingType = "r2_bucket"
)

type AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uQueueBinding struct {
	// Name of the Queue to bind to
	QueueName param.Field[string] `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uQueueBindingType] `json:"type,required"`
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uQueueBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uQueueBinding) implementsAccountWorkerDispatchNamespaceScriptSettingUpdateParamsBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uQueueBindingType string

const (
	AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uQueueBindingTypeQueue AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uQueueBindingType = "queue"
)

type AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uD1Binding struct {
	// ID of the D1 database to bind to
	ID param.Field[string] `json:"id,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name param.Field[string] `json:"name,required"`
	// The class of resource that the binding provides.
	Type param.Field[AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uD1BindingType] `json:"type,required"`
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uD1Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uD1Binding) implementsAccountWorkerDispatchNamespaceScriptSettingUpdateParamsBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uD1BindingType string

const (
	AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uD1BindingTypeD1 AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uD1BindingType = "d1"
)

type AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uDispatchNamespaceBinding struct {
	// Namespace to bind to
	Namespace param.Field[string] `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type param.Field[AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uDispatchNamespaceBindingType] `json:"type,required"`
	// Outbound worker
	Outbound param.Field[AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uDispatchNamespaceBindingOutbound] `json:"outbound"`
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uDispatchNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uDispatchNamespaceBinding) implementsAccountWorkerDispatchNamespaceScriptSettingUpdateParamsBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uDispatchNamespaceBindingType string

const (
	AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uDispatchNamespaceBindingTypeDispatchNamespace AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params param.Field[[]string] `json:"params"`
	// Outbound worker
	Worker param.Field[AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker] `json:"worker"`
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uDispatchNamespaceBindingOutbound) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker
type AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker
	Service param.Field[string] `json:"service"`
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uMtlsCertBinding struct {
	// The class of resource that the binding provides.
	Type param.Field[AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uMtlsCertBindingType] `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID param.Field[string] `json:"certificate_id"`
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uMtlsCertBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uMtlsCertBinding) implementsAccountWorkerDispatchNamespaceScriptSettingUpdateParamsBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uMtlsCertBindingType string

const (
	AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uMtlsCertBindingTypeMtlsCertificate AccountWorkerDispatchNamespaceScriptSettingUpdateParamsBindingsAvYbsl2uMtlsCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by
// [AccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrationsAvYbsl2uSingleStepMigrations],
// [AccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrationsAvYbsl2uSteppedMigrations].
type AccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrations interface {
	implementsAccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrations()
}

// A single set of migrations to apply.
type AccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrationsAvYbsl2uSingleStepMigrations struct {
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
	RenamedClasses param.Field[[]AccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrationsAvYbsl2uSingleStepMigrationsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]AccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrationsAvYbsl2uSingleStepMigrationsTransferredClass] `json:"transferred_classes"`
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrationsAvYbsl2uSingleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrationsAvYbsl2uSingleStepMigrations) implementsAccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrations() {
}

type AccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrationsAvYbsl2uSingleStepMigrationsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrationsAvYbsl2uSingleStepMigrationsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrationsAvYbsl2uSingleStepMigrationsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrationsAvYbsl2uSingleStepMigrationsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrationsAvYbsl2uSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]AccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrationsAvYbsl2uSteppedMigrationsStep] `json:"steps"`
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrationsAvYbsl2uSteppedMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrationsAvYbsl2uSteppedMigrations) implementsAccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrations() {
}

type AccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrationsAvYbsl2uSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]AccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]AccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass] `json:"transferred_classes"`
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrationsAvYbsl2uSteppedMigrationsStep) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateParamsMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerDispatchNamespaceScriptSettingUpdateParamsPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode param.Field[AccountWorkerDispatchNamespaceScriptSettingUpdateParamsPlacementMode] `json:"mode"`
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateParamsPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type AccountWorkerDispatchNamespaceScriptSettingUpdateParamsPlacementMode string

const (
	AccountWorkerDispatchNamespaceScriptSettingUpdateParamsPlacementModeSmart AccountWorkerDispatchNamespaceScriptSettingUpdateParamsPlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type AccountWorkerDispatchNamespaceScriptSettingUpdateParamsTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service param.Field[string] `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace param.Field[string] `json:"namespace"`
}

func (r AccountWorkerDispatchNamespaceScriptSettingUpdateParamsTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
