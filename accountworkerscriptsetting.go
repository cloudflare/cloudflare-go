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

// AccountWorkerScriptSettingService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountWorkerScriptSettingService] method instead.
type AccountWorkerScriptSettingService struct {
	Options []option.RequestOption
}

// NewAccountWorkerScriptSettingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountWorkerScriptSettingService(opts ...option.RequestOption) (r *AccountWorkerScriptSettingService) {
	r = &AccountWorkerScriptSettingService{}
	r.Options = opts
	return
}

// Get script metadata and config, such as bindings or usage model
func (r *AccountWorkerScriptSettingService) Get(ctx context.Context, accountIdentifier string, scriptName string, opts ...option.RequestOption) (res *AccountWorkerScriptSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/settings", accountIdentifier, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Patch script metadata or config, such as bindings or usage model
func (r *AccountWorkerScriptSettingService) Update(ctx context.Context, accountIdentifier string, scriptName string, body AccountWorkerScriptSettingUpdateParams, opts ...option.RequestOption) (res *AccountWorkerScriptSettingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/settings", accountIdentifier, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type AccountWorkerScriptSettingGetResponse struct {
	// List of bindings attached to this Worker
	Bindings []AccountWorkerScriptSettingGetResponseBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations AccountWorkerScriptSettingGetResponseMigrations `json:"migrations"`
	Placement  AccountWorkerScriptSettingGetResponsePlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []AccountWorkerScriptSettingGetResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                    `json:"usage_model"`
	JSON       accountWorkerScriptSettingGetResponseJSON `json:"-"`
}

// accountWorkerScriptSettingGetResponseJSON contains the JSON metadata for the
// struct [AccountWorkerScriptSettingGetResponse]
type accountWorkerScriptSettingGetResponseJSON struct {
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

func (r *AccountWorkerScriptSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uKvNamespaceBinding],
// [AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uServiceBinding],
// [AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uDoBinding],
// [AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uR2Binding],
// [AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uQueueBinding],
// [AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uD1Binding],
// [AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBinding]
// or [AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uMtlsCertBinding].
type AccountWorkerScriptSettingGetResponseBinding interface {
	implementsAccountWorkerScriptSettingGetResponseBinding()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountWorkerScriptSettingGetResponseBinding)(nil)).Elem(), "")
}

type AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uKvNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uKvNamespaceBindingType `json:"type,required"`
	JSON accountWorkerScriptSettingGetResponseBindingsAvYbsl2uKvNamespaceBindingJSON `json:"-"`
}

// accountWorkerScriptSettingGetResponseBindingsAvYbsl2uKvNamespaceBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uKvNamespaceBinding]
type accountWorkerScriptSettingGetResponseBindingsAvYbsl2uKvNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uKvNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uKvNamespaceBinding) implementsAccountWorkerScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uKvNamespaceBindingType string

const (
	AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uKvNamespaceBindingTypeKvNamespace AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uKvNamespaceBindingType = "kv_namespace"
)

type AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uServiceBindingType `json:"type,required"`
	JSON accountWorkerScriptSettingGetResponseBindingsAvYbsl2uServiceBindingJSON `json:"-"`
}

// accountWorkerScriptSettingGetResponseBindingsAvYbsl2uServiceBindingJSON contains
// the JSON metadata for the struct
// [AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uServiceBinding]
type accountWorkerScriptSettingGetResponseBindingsAvYbsl2uServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uServiceBinding) implementsAccountWorkerScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uServiceBindingType string

const (
	AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uServiceBindingTypeService AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uServiceBindingType = "service"
)

type AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                             `json:"script_name"`
	JSON       accountWorkerScriptSettingGetResponseBindingsAvYbsl2uDoBindingJSON `json:"-"`
}

// accountWorkerScriptSettingGetResponseBindingsAvYbsl2uDoBindingJSON contains the
// JSON metadata for the struct
// [AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uDoBinding]
type accountWorkerScriptSettingGetResponseBindingsAvYbsl2uDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uDoBinding) implementsAccountWorkerScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uDoBindingType string

const (
	AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uDoBindingTypeDurableObjectNamespace AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uDoBindingType = "durable_object_namespace"
)

type AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uR2BindingType `json:"type,required"`
	JSON accountWorkerScriptSettingGetResponseBindingsAvYbsl2uR2BindingJSON `json:"-"`
}

// accountWorkerScriptSettingGetResponseBindingsAvYbsl2uR2BindingJSON contains the
// JSON metadata for the struct
// [AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uR2Binding]
type accountWorkerScriptSettingGetResponseBindingsAvYbsl2uR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uR2Binding) implementsAccountWorkerScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uR2BindingType string

const (
	AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uR2BindingTypeR2Bucket AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uR2BindingType = "r2_bucket"
)

type AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uQueueBindingType `json:"type,required"`
	JSON accountWorkerScriptSettingGetResponseBindingsAvYbsl2uQueueBindingJSON `json:"-"`
}

// accountWorkerScriptSettingGetResponseBindingsAvYbsl2uQueueBindingJSON contains
// the JSON metadata for the struct
// [AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uQueueBinding]
type accountWorkerScriptSettingGetResponseBindingsAvYbsl2uQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uQueueBinding) implementsAccountWorkerScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uQueueBindingType string

const (
	AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uQueueBindingTypeQueue AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uQueueBindingType = "queue"
)

type AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uD1BindingType `json:"type,required"`
	JSON accountWorkerScriptSettingGetResponseBindingsAvYbsl2uD1BindingJSON `json:"-"`
}

// accountWorkerScriptSettingGetResponseBindingsAvYbsl2uD1BindingJSON contains the
// JSON metadata for the struct
// [AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uD1Binding]
type accountWorkerScriptSettingGetResponseBindingsAvYbsl2uD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uD1Binding) implementsAccountWorkerScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uD1BindingType string

const (
	AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uD1BindingTypeD1 AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uD1BindingType = "d1"
)

type AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     accountWorkerScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingJSON     `json:"-"`
}

// accountWorkerScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBinding]
type accountWorkerScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBinding) implementsAccountWorkerScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingType string

const (
	AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingTypeDispatchNamespace AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   accountWorkerScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// accountWorkerScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutbound]
type accountWorkerScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Outbound worker
type AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                          `json:"service"`
	JSON    accountWorkerScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// accountWorkerScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker]
type accountWorkerScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uMtlsCertBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uMtlsCertBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                                   `json:"certificate_id"`
	JSON          accountWorkerScriptSettingGetResponseBindingsAvYbsl2uMtlsCertBindingJSON `json:"-"`
}

// accountWorkerScriptSettingGetResponseBindingsAvYbsl2uMtlsCertBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uMtlsCertBinding]
type accountWorkerScriptSettingGetResponseBindingsAvYbsl2uMtlsCertBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uMtlsCertBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uMtlsCertBinding) implementsAccountWorkerScriptSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uMtlsCertBindingType string

const (
	AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uMtlsCertBindingTypeMtlsCertificate AccountWorkerScriptSettingGetResponseBindingsAvYbsl2uMtlsCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [AccountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrations] or
// [AccountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrations].
type AccountWorkerScriptSettingGetResponseMigrations interface {
	implementsAccountWorkerScriptSettingGetResponseMigrations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountWorkerScriptSettingGetResponseMigrations)(nil)).Elem(), "")
}

// A single set of migrations to apply.
type AccountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrations struct {
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
	RenamedClasses []AccountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []AccountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               accountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsJSON               `json:"-"`
}

// accountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrations]
type accountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrations) implementsAccountWorkerScriptSettingGetResponseMigrations() {
}

type AccountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClass struct {
	From string                                                                                      `json:"from"`
	To   string                                                                                      `json:"to"`
	JSON accountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// accountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClass]
type accountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClass struct {
	From       string                                                                                          `json:"from"`
	FromScript string                                                                                          `json:"from_script"`
	To         string                                                                                          `json:"to"`
	JSON       accountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// accountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClass]
type accountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []AccountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStep `json:"steps"`
	JSON  accountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsJSON   `json:"-"`
}

// accountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrations]
type accountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrations) implementsAccountWorkerScriptSettingGetResponseMigrations() {
}

type AccountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []AccountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []AccountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               accountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepJSON                `json:"-"`
}

// accountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStep]
type accountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass struct {
	From string                                                                                        `json:"from"`
	To   string                                                                                        `json:"to"`
	JSON accountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// accountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass]
type accountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                                            `json:"from"`
	FromScript string                                                                                            `json:"from_script"`
	To         string                                                                                            `json:"to"`
	JSON       accountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// accountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass]
type accountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptSettingGetResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode AccountWorkerScriptSettingGetResponsePlacementMode `json:"mode"`
	JSON accountWorkerScriptSettingGetResponsePlacementJSON `json:"-"`
}

// accountWorkerScriptSettingGetResponsePlacementJSON contains the JSON metadata
// for the struct [AccountWorkerScriptSettingGetResponsePlacement]
type accountWorkerScriptSettingGetResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingGetResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type AccountWorkerScriptSettingGetResponsePlacementMode string

const (
	AccountWorkerScriptSettingGetResponsePlacementModeSmart AccountWorkerScriptSettingGetResponsePlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type AccountWorkerScriptSettingGetResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                `json:"namespace"`
	JSON      accountWorkerScriptSettingGetResponseTailConsumerJSON `json:"-"`
}

// accountWorkerScriptSettingGetResponseTailConsumerJSON contains the JSON metadata
// for the struct [AccountWorkerScriptSettingGetResponseTailConsumer]
type accountWorkerScriptSettingGetResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingGetResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptSettingUpdateResponse struct {
	// List of bindings attached to this Worker
	Bindings []AccountWorkerScriptSettingUpdateResponseBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations AccountWorkerScriptSettingUpdateResponseMigrations `json:"migrations"`
	Placement  AccountWorkerScriptSettingUpdateResponsePlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []AccountWorkerScriptSettingUpdateResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                       `json:"usage_model"`
	JSON       accountWorkerScriptSettingUpdateResponseJSON `json:"-"`
}

// accountWorkerScriptSettingUpdateResponseJSON contains the JSON metadata for the
// struct [AccountWorkerScriptSettingUpdateResponse]
type accountWorkerScriptSettingUpdateResponseJSON struct {
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

func (r *AccountWorkerScriptSettingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBinding],
// [AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uServiceBinding],
// [AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDoBinding],
// [AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uR2Binding],
// [AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uQueueBinding],
// [AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uD1Binding],
// [AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBinding]
// or [AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uMtlsCertBinding].
type AccountWorkerScriptSettingUpdateResponseBinding interface {
	implementsAccountWorkerScriptSettingUpdateResponseBinding()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountWorkerScriptSettingUpdateResponseBinding)(nil)).Elem(), "")
}

type AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBindingType `json:"type,required"`
	JSON accountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBindingJSON `json:"-"`
}

// accountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBinding]
type accountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBinding) implementsAccountWorkerScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBindingType string

const (
	AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBindingTypeKvNamespace AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBindingType = "kv_namespace"
)

type AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uServiceBindingType `json:"type,required"`
	JSON accountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uServiceBindingJSON `json:"-"`
}

// accountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uServiceBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uServiceBinding]
type accountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uServiceBinding) implementsAccountWorkerScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uServiceBindingType string

const (
	AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uServiceBindingTypeService AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uServiceBindingType = "service"
)

type AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                                `json:"script_name"`
	JSON       accountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDoBindingJSON `json:"-"`
}

// accountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDoBindingJSON contains
// the JSON metadata for the struct
// [AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDoBinding]
type accountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDoBinding) implementsAccountWorkerScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDoBindingType string

const (
	AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDoBindingTypeDurableObjectNamespace AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDoBindingType = "durable_object_namespace"
)

type AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uR2BindingType `json:"type,required"`
	JSON accountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uR2BindingJSON `json:"-"`
}

// accountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uR2BindingJSON contains
// the JSON metadata for the struct
// [AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uR2Binding]
type accountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uR2Binding) implementsAccountWorkerScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uR2BindingType string

const (
	AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uR2BindingTypeR2Bucket AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uR2BindingType = "r2_bucket"
)

type AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uQueueBindingType `json:"type,required"`
	JSON accountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uQueueBindingJSON `json:"-"`
}

// accountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uQueueBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uQueueBinding]
type accountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uQueueBinding) implementsAccountWorkerScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uQueueBindingType string

const (
	AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uQueueBindingTypeQueue AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uQueueBindingType = "queue"
)

type AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uD1BindingType `json:"type,required"`
	JSON accountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uD1BindingJSON `json:"-"`
}

// accountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uD1BindingJSON contains
// the JSON metadata for the struct
// [AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uD1Binding]
type accountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uD1Binding) implementsAccountWorkerScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uD1BindingType string

const (
	AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uD1BindingTypeD1 AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uD1BindingType = "d1"
)

type AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     accountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingJSON     `json:"-"`
}

// accountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBinding]
type accountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBinding) implementsAccountWorkerScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingType string

const (
	AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingTypeDispatchNamespace AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   accountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// accountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutbound]
type accountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Outbound worker
type AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                             `json:"service"`
	JSON    accountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// accountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker]
type accountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uMtlsCertBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uMtlsCertBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                                      `json:"certificate_id"`
	JSON          accountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uMtlsCertBindingJSON `json:"-"`
}

// accountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uMtlsCertBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uMtlsCertBinding]
type accountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uMtlsCertBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uMtlsCertBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uMtlsCertBinding) implementsAccountWorkerScriptSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uMtlsCertBindingType string

const (
	AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uMtlsCertBindingTypeMtlsCertificate AccountWorkerScriptSettingUpdateResponseBindingsAvYbsl2uMtlsCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [AccountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrations]
// or
// [AccountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrations].
type AccountWorkerScriptSettingUpdateResponseMigrations interface {
	implementsAccountWorkerScriptSettingUpdateResponseMigrations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountWorkerScriptSettingUpdateResponseMigrations)(nil)).Elem(), "")
}

// A single set of migrations to apply.
type AccountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrations struct {
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
	RenamedClasses []AccountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []AccountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               accountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsJSON               `json:"-"`
}

// accountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrations]
type accountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrations) implementsAccountWorkerScriptSettingUpdateResponseMigrations() {
}

type AccountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClass struct {
	From string                                                                                         `json:"from"`
	To   string                                                                                         `json:"to"`
	JSON accountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// accountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClass]
type accountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClass struct {
	From       string                                                                                             `json:"from"`
	FromScript string                                                                                             `json:"from_script"`
	To         string                                                                                             `json:"to"`
	JSON       accountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// accountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClass]
type accountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []AccountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStep `json:"steps"`
	JSON  accountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsJSON   `json:"-"`
}

// accountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrations]
type accountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrations) implementsAccountWorkerScriptSettingUpdateResponseMigrations() {
}

type AccountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []AccountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []AccountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               accountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepJSON                `json:"-"`
}

// accountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStep]
type accountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass struct {
	From string                                                                                           `json:"from"`
	To   string                                                                                           `json:"to"`
	JSON accountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// accountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass]
type accountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                                               `json:"from"`
	FromScript string                                                                                               `json:"from_script"`
	To         string                                                                                               `json:"to"`
	JSON       accountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// accountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass]
type accountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptSettingUpdateResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode AccountWorkerScriptSettingUpdateResponsePlacementMode `json:"mode"`
	JSON accountWorkerScriptSettingUpdateResponsePlacementJSON `json:"-"`
}

// accountWorkerScriptSettingUpdateResponsePlacementJSON contains the JSON metadata
// for the struct [AccountWorkerScriptSettingUpdateResponsePlacement]
type accountWorkerScriptSettingUpdateResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingUpdateResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type AccountWorkerScriptSettingUpdateResponsePlacementMode string

const (
	AccountWorkerScriptSettingUpdateResponsePlacementModeSmart AccountWorkerScriptSettingUpdateResponsePlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type AccountWorkerScriptSettingUpdateResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                   `json:"namespace"`
	JSON      accountWorkerScriptSettingUpdateResponseTailConsumerJSON `json:"-"`
}

// accountWorkerScriptSettingUpdateResponseTailConsumerJSON contains the JSON
// metadata for the struct [AccountWorkerScriptSettingUpdateResponseTailConsumer]
type accountWorkerScriptSettingUpdateResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptSettingUpdateResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptSettingUpdateParams struct {
	Settings param.Field[AccountWorkerScriptSettingUpdateParamsSettings] `json:"settings"`
}

func (r AccountWorkerScriptSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerScriptSettingUpdateParamsSettings struct {
	// List of bindings attached to this Worker
	Bindings param.Field[[]AccountWorkerScriptSettingUpdateParamsSettingsBinding] `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations param.Field[AccountWorkerScriptSettingUpdateParamsSettingsMigrations] `json:"migrations"`
	Placement  param.Field[AccountWorkerScriptSettingUpdateParamsSettingsPlacement]  `json:"placement"`
	// Tags to help you manage your Workers
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]AccountWorkerScriptSettingUpdateParamsSettingsTailConsumer] `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel param.Field[string] `json:"usage_model"`
}

func (r AccountWorkerScriptSettingUpdateParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A binding to allow the Worker to communicate with resources
//
// Satisfied by
// [AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uKvNamespaceBinding],
// [AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uServiceBinding],
// [AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uDoBinding],
// [AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uR2Binding],
// [AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uQueueBinding],
// [AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uD1Binding],
// [AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uDispatchNamespaceBinding],
// [AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uMtlsCertBinding].
type AccountWorkerScriptSettingUpdateParamsSettingsBinding interface {
	implementsAccountWorkerScriptSettingUpdateParamsSettingsBinding()
}

type AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uKvNamespaceBinding struct {
	// The class of resource that the binding provides.
	Type param.Field[AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uKvNamespaceBindingType] `json:"type,required"`
}

func (r AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uKvNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uKvNamespaceBinding) implementsAccountWorkerScriptSettingUpdateParamsSettingsBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uKvNamespaceBindingType string

const (
	AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uKvNamespaceBindingTypeKvNamespace AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uKvNamespaceBindingType = "kv_namespace"
)

type AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// Name of Worker to bind to
	Service param.Field[string] `json:"service,required"`
	// The class of resource that the binding provides.
	Type param.Field[AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uServiceBindingType] `json:"type,required"`
}

func (r AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uServiceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uServiceBinding) implementsAccountWorkerScriptSettingUpdateParamsSettingsBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uServiceBindingType string

const (
	AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uServiceBindingTypeService AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uServiceBindingType = "service"
)

type AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uDoBinding struct {
	// The exported class name of the Durable Object
	ClassName param.Field[string] `json:"class_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uDoBindingType] `json:"type,required"`
	// The environment of the script_name to bind to
	Environment param.Field[string] `json:"environment"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName param.Field[string] `json:"script_name"`
}

func (r AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uDoBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uDoBinding) implementsAccountWorkerScriptSettingUpdateParamsSettingsBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uDoBindingType string

const (
	AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uDoBindingTypeDurableObjectNamespace AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uDoBindingType = "durable_object_namespace"
)

type AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uR2Binding struct {
	// R2 bucket to bind to
	BucketName param.Field[string] `json:"bucket_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uR2BindingType] `json:"type,required"`
}

func (r AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uR2Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uR2Binding) implementsAccountWorkerScriptSettingUpdateParamsSettingsBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uR2BindingType string

const (
	AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uR2BindingTypeR2Bucket AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uR2BindingType = "r2_bucket"
)

type AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uQueueBinding struct {
	// Name of the Queue to bind to
	QueueName param.Field[string] `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uQueueBindingType] `json:"type,required"`
}

func (r AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uQueueBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uQueueBinding) implementsAccountWorkerScriptSettingUpdateParamsSettingsBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uQueueBindingType string

const (
	AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uQueueBindingTypeQueue AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uQueueBindingType = "queue"
)

type AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uD1Binding struct {
	// ID of the D1 database to bind to
	ID param.Field[string] `json:"id,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name param.Field[string] `json:"name,required"`
	// The class of resource that the binding provides.
	Type param.Field[AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uD1BindingType] `json:"type,required"`
}

func (r AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uD1Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uD1Binding) implementsAccountWorkerScriptSettingUpdateParamsSettingsBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uD1BindingType string

const (
	AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uD1BindingTypeD1 AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uD1BindingType = "d1"
)

type AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uDispatchNamespaceBinding struct {
	// Namespace to bind to
	Namespace param.Field[string] `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type param.Field[AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uDispatchNamespaceBindingType] `json:"type,required"`
	// Outbound worker
	Outbound param.Field[AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uDispatchNamespaceBindingOutbound] `json:"outbound"`
}

func (r AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uDispatchNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uDispatchNamespaceBinding) implementsAccountWorkerScriptSettingUpdateParamsSettingsBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uDispatchNamespaceBindingType string

const (
	AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uDispatchNamespaceBindingTypeDispatchNamespace AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params param.Field[[]string] `json:"params"`
	// Outbound worker
	Worker param.Field[AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker] `json:"worker"`
}

func (r AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uDispatchNamespaceBindingOutbound) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker
type AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker
	Service param.Field[string] `json:"service"`
}

func (r AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uMtlsCertBinding struct {
	// The class of resource that the binding provides.
	Type param.Field[AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uMtlsCertBindingType] `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID param.Field[string] `json:"certificate_id"`
}

func (r AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uMtlsCertBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uMtlsCertBinding) implementsAccountWorkerScriptSettingUpdateParamsSettingsBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uMtlsCertBindingType string

const (
	AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uMtlsCertBindingTypeMtlsCertificate AccountWorkerScriptSettingUpdateParamsSettingsBindingsAvYbsl2uMtlsCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by
// [AccountWorkerScriptSettingUpdateParamsSettingsMigrationsAvYbsl2uSingleStepMigrations],
// [AccountWorkerScriptSettingUpdateParamsSettingsMigrationsAvYbsl2uSteppedMigrations].
type AccountWorkerScriptSettingUpdateParamsSettingsMigrations interface {
	implementsAccountWorkerScriptSettingUpdateParamsSettingsMigrations()
}

// A single set of migrations to apply.
type AccountWorkerScriptSettingUpdateParamsSettingsMigrationsAvYbsl2uSingleStepMigrations struct {
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
	RenamedClasses param.Field[[]AccountWorkerScriptSettingUpdateParamsSettingsMigrationsAvYbsl2uSingleStepMigrationsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]AccountWorkerScriptSettingUpdateParamsSettingsMigrationsAvYbsl2uSingleStepMigrationsTransferredClass] `json:"transferred_classes"`
}

func (r AccountWorkerScriptSettingUpdateParamsSettingsMigrationsAvYbsl2uSingleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerScriptSettingUpdateParamsSettingsMigrationsAvYbsl2uSingleStepMigrations) implementsAccountWorkerScriptSettingUpdateParamsSettingsMigrations() {
}

type AccountWorkerScriptSettingUpdateParamsSettingsMigrationsAvYbsl2uSingleStepMigrationsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r AccountWorkerScriptSettingUpdateParamsSettingsMigrationsAvYbsl2uSingleStepMigrationsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerScriptSettingUpdateParamsSettingsMigrationsAvYbsl2uSingleStepMigrationsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r AccountWorkerScriptSettingUpdateParamsSettingsMigrationsAvYbsl2uSingleStepMigrationsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerScriptSettingUpdateParamsSettingsMigrationsAvYbsl2uSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]AccountWorkerScriptSettingUpdateParamsSettingsMigrationsAvYbsl2uSteppedMigrationsStep] `json:"steps"`
}

func (r AccountWorkerScriptSettingUpdateParamsSettingsMigrationsAvYbsl2uSteppedMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerScriptSettingUpdateParamsSettingsMigrationsAvYbsl2uSteppedMigrations) implementsAccountWorkerScriptSettingUpdateParamsSettingsMigrations() {
}

type AccountWorkerScriptSettingUpdateParamsSettingsMigrationsAvYbsl2uSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]AccountWorkerScriptSettingUpdateParamsSettingsMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]AccountWorkerScriptSettingUpdateParamsSettingsMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass] `json:"transferred_classes"`
}

func (r AccountWorkerScriptSettingUpdateParamsSettingsMigrationsAvYbsl2uSteppedMigrationsStep) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerScriptSettingUpdateParamsSettingsMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r AccountWorkerScriptSettingUpdateParamsSettingsMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerScriptSettingUpdateParamsSettingsMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r AccountWorkerScriptSettingUpdateParamsSettingsMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerScriptSettingUpdateParamsSettingsPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode param.Field[AccountWorkerScriptSettingUpdateParamsSettingsPlacementMode] `json:"mode"`
}

func (r AccountWorkerScriptSettingUpdateParamsSettingsPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type AccountWorkerScriptSettingUpdateParamsSettingsPlacementMode string

const (
	AccountWorkerScriptSettingUpdateParamsSettingsPlacementModeSmart AccountWorkerScriptSettingUpdateParamsSettingsPlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type AccountWorkerScriptSettingUpdateParamsSettingsTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service param.Field[string] `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace param.Field[string] `json:"namespace"`
}

func (r AccountWorkerScriptSettingUpdateParamsSettingsTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
