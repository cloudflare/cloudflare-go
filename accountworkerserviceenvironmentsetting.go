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

// AccountWorkerServiceEnvironmentSettingService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkerServiceEnvironmentSettingService] method instead.
type AccountWorkerServiceEnvironmentSettingService struct {
	Options []option.RequestOption
}

// NewAccountWorkerServiceEnvironmentSettingService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountWorkerServiceEnvironmentSettingService(opts ...option.RequestOption) (r *AccountWorkerServiceEnvironmentSettingService) {
	r = &AccountWorkerServiceEnvironmentSettingService{}
	r.Options = opts
	return
}

// Get script settings from a worker with an environment
func (r *AccountWorkerServiceEnvironmentSettingService) Get(ctx context.Context, accountIdentifier string, serviceName string, environmentName string, opts ...option.RequestOption) (res *AccountWorkerServiceEnvironmentSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/services/%s/environments/%s/settings", accountIdentifier, serviceName, environmentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Patch script metadata, such as bindings
func (r *AccountWorkerServiceEnvironmentSettingService) Update(ctx context.Context, accountIdentifier string, serviceName string, environmentName string, body AccountWorkerServiceEnvironmentSettingUpdateParams, opts ...option.RequestOption) (res *AccountWorkerServiceEnvironmentSettingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/services/%s/environments/%s/settings", accountIdentifier, serviceName, environmentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type AccountWorkerServiceEnvironmentSettingGetResponse struct {
	// List of bindings attached to this Worker
	Bindings []AccountWorkerServiceEnvironmentSettingGetResponseBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations AccountWorkerServiceEnvironmentSettingGetResponseMigrations `json:"migrations"`
	Placement  AccountWorkerServiceEnvironmentSettingGetResponsePlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []AccountWorkerServiceEnvironmentSettingGetResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                                `json:"usage_model"`
	JSON       accountWorkerServiceEnvironmentSettingGetResponseJSON `json:"-"`
}

// accountWorkerServiceEnvironmentSettingGetResponseJSON contains the JSON metadata
// for the struct [AccountWorkerServiceEnvironmentSettingGetResponse]
type accountWorkerServiceEnvironmentSettingGetResponseJSON struct {
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

func (r *AccountWorkerServiceEnvironmentSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uKvNamespaceBinding],
// [AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uServiceBinding],
// [AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDoBinding],
// [AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uR2Binding],
// [AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uQueueBinding],
// [AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uD1Binding],
// [AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBinding]
// or
// [AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uMtlsCertBinding].
type AccountWorkerServiceEnvironmentSettingGetResponseBinding interface {
	implementsAccountWorkerServiceEnvironmentSettingGetResponseBinding()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountWorkerServiceEnvironmentSettingGetResponseBinding)(nil)).Elem(), "")
}

type AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uKvNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uKvNamespaceBindingType `json:"type,required"`
	JSON accountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uKvNamespaceBindingJSON `json:"-"`
}

// accountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uKvNamespaceBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uKvNamespaceBinding]
type accountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uKvNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uKvNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uKvNamespaceBinding) implementsAccountWorkerServiceEnvironmentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uKvNamespaceBindingType string

const (
	AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uKvNamespaceBindingTypeKvNamespace AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uKvNamespaceBindingType = "kv_namespace"
)

type AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uServiceBindingType `json:"type,required"`
	JSON accountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uServiceBindingJSON `json:"-"`
}

// accountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uServiceBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uServiceBinding]
type accountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uServiceBinding) implementsAccountWorkerServiceEnvironmentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uServiceBindingType string

const (
	AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uServiceBindingTypeService AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uServiceBindingType = "service"
)

type AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                                         `json:"script_name"`
	JSON       accountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDoBindingJSON `json:"-"`
}

// accountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDoBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDoBinding]
type accountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDoBinding) implementsAccountWorkerServiceEnvironmentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDoBindingType string

const (
	AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDoBindingTypeDurableObjectNamespace AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDoBindingType = "durable_object_namespace"
)

type AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uR2BindingType `json:"type,required"`
	JSON accountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uR2BindingJSON `json:"-"`
}

// accountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uR2BindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uR2Binding]
type accountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uR2Binding) implementsAccountWorkerServiceEnvironmentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uR2BindingType string

const (
	AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uR2BindingTypeR2Bucket AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uR2BindingType = "r2_bucket"
)

type AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uQueueBindingType `json:"type,required"`
	JSON accountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uQueueBindingJSON `json:"-"`
}

// accountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uQueueBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uQueueBinding]
type accountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uQueueBinding) implementsAccountWorkerServiceEnvironmentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uQueueBindingType string

const (
	AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uQueueBindingTypeQueue AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uQueueBindingType = "queue"
)

type AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uD1BindingType `json:"type,required"`
	JSON accountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uD1BindingJSON `json:"-"`
}

// accountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uD1BindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uD1Binding]
type accountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uD1Binding) implementsAccountWorkerServiceEnvironmentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uD1BindingType string

const (
	AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uD1BindingTypeD1 AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uD1BindingType = "d1"
)

type AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     accountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingJSON     `json:"-"`
}

// accountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBinding]
type accountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBinding) implementsAccountWorkerServiceEnvironmentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingType string

const (
	AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingTypeDispatchNamespace AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   accountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// accountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutbound]
type accountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Outbound worker
type AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                                      `json:"service"`
	JSON    accountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// accountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker]
type accountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uMtlsCertBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uMtlsCertBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                                               `json:"certificate_id"`
	JSON          accountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uMtlsCertBindingJSON `json:"-"`
}

// accountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uMtlsCertBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uMtlsCertBinding]
type accountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uMtlsCertBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uMtlsCertBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uMtlsCertBinding) implementsAccountWorkerServiceEnvironmentSettingGetResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uMtlsCertBindingType string

const (
	AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uMtlsCertBindingTypeMtlsCertificate AccountWorkerServiceEnvironmentSettingGetResponseBindingsAvYbsl2uMtlsCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [AccountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSingleStepMigrations]
// or
// [AccountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSteppedMigrations].
type AccountWorkerServiceEnvironmentSettingGetResponseMigrations interface {
	implementsAccountWorkerServiceEnvironmentSettingGetResponseMigrations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountWorkerServiceEnvironmentSettingGetResponseMigrations)(nil)).Elem(), "")
}

// A single set of migrations to apply.
type AccountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSingleStepMigrations struct {
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
	RenamedClasses []AccountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []AccountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               accountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsJSON               `json:"-"`
}

// accountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSingleStepMigrations]
type accountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSingleStepMigrations) implementsAccountWorkerServiceEnvironmentSettingGetResponseMigrations() {
}

type AccountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClass struct {
	From string                                                                                                  `json:"from"`
	To   string                                                                                                  `json:"to"`
	JSON accountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// accountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClass]
type accountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClass struct {
	From       string                                                                                                      `json:"from"`
	FromScript string                                                                                                      `json:"from_script"`
	To         string                                                                                                      `json:"to"`
	JSON       accountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// accountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClass]
type accountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []AccountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStep `json:"steps"`
	JSON  accountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsJSON   `json:"-"`
}

// accountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSteppedMigrations]
type accountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSteppedMigrations) implementsAccountWorkerServiceEnvironmentSettingGetResponseMigrations() {
}

type AccountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []AccountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []AccountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               accountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepJSON                `json:"-"`
}

// accountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStep]
type accountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass struct {
	From string                                                                                                    `json:"from"`
	To   string                                                                                                    `json:"to"`
	JSON accountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// accountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass]
type accountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                                                        `json:"from"`
	FromScript string                                                                                                        `json:"from_script"`
	To         string                                                                                                        `json:"to"`
	JSON       accountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// accountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass]
type accountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingGetResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerServiceEnvironmentSettingGetResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode AccountWorkerServiceEnvironmentSettingGetResponsePlacementMode `json:"mode"`
	JSON accountWorkerServiceEnvironmentSettingGetResponsePlacementJSON `json:"-"`
}

// accountWorkerServiceEnvironmentSettingGetResponsePlacementJSON contains the JSON
// metadata for the struct
// [AccountWorkerServiceEnvironmentSettingGetResponsePlacement]
type accountWorkerServiceEnvironmentSettingGetResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingGetResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type AccountWorkerServiceEnvironmentSettingGetResponsePlacementMode string

const (
	AccountWorkerServiceEnvironmentSettingGetResponsePlacementModeSmart AccountWorkerServiceEnvironmentSettingGetResponsePlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type AccountWorkerServiceEnvironmentSettingGetResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                            `json:"namespace"`
	JSON      accountWorkerServiceEnvironmentSettingGetResponseTailConsumerJSON `json:"-"`
}

// accountWorkerServiceEnvironmentSettingGetResponseTailConsumerJSON contains the
// JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingGetResponseTailConsumer]
type accountWorkerServiceEnvironmentSettingGetResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingGetResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerServiceEnvironmentSettingUpdateResponse struct {
	// List of bindings attached to this Worker
	Bindings []AccountWorkerServiceEnvironmentSettingUpdateResponseBinding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations AccountWorkerServiceEnvironmentSettingUpdateResponseMigrations `json:"migrations"`
	Placement  AccountWorkerServiceEnvironmentSettingUpdateResponsePlacement  `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []AccountWorkerServiceEnvironmentSettingUpdateResponseTailConsumer `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                                   `json:"usage_model"`
	JSON       accountWorkerServiceEnvironmentSettingUpdateResponseJSON `json:"-"`
}

// accountWorkerServiceEnvironmentSettingUpdateResponseJSON contains the JSON
// metadata for the struct [AccountWorkerServiceEnvironmentSettingUpdateResponse]
type accountWorkerServiceEnvironmentSettingUpdateResponseJSON struct {
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

func (r *AccountWorkerServiceEnvironmentSettingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A binding to allow the Worker to communicate with resources
//
// Union satisfied by
// [AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBinding],
// [AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uServiceBinding],
// [AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDoBinding],
// [AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uR2Binding],
// [AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uQueueBinding],
// [AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uD1Binding],
// [AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBinding]
// or
// [AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uMtlsCertBinding].
type AccountWorkerServiceEnvironmentSettingUpdateResponseBinding interface {
	implementsAccountWorkerServiceEnvironmentSettingUpdateResponseBinding()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountWorkerServiceEnvironmentSettingUpdateResponseBinding)(nil)).Elem(), "")
}

type AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBindingType `json:"type,required"`
	JSON accountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBindingJSON `json:"-"`
}

// accountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBinding]
type accountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBindingJSON struct {
	Name        apijson.Field
	NamespaceID apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBinding) implementsAccountWorkerServiceEnvironmentSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBindingType string

const (
	AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBindingTypeKvNamespace AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uKvNamespaceBindingType = "kv_namespace"
)

type AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of Worker to bind to
	Service string `json:"service,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uServiceBindingType `json:"type,required"`
	JSON accountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uServiceBindingJSON `json:"-"`
}

// accountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uServiceBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uServiceBinding]
type accountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uServiceBindingJSON struct {
	Environment apijson.Field
	Name        apijson.Field
	Service     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uServiceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uServiceBinding) implementsAccountWorkerServiceEnvironmentSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uServiceBindingType string

const (
	AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uServiceBindingTypeService AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uServiceBindingType = "service"
)

type AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDoBinding struct {
	// The exported class name of the Durable Object
	ClassName string `json:"class_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDoBindingType `json:"type,required"`
	// The environment of the script_name to bind to
	Environment string `json:"environment"`
	// Namespace identifier tag.
	NamespaceID string `json:"namespace_id"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName string                                                                            `json:"script_name"`
	JSON       accountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDoBindingJSON `json:"-"`
}

// accountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDoBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDoBinding]
type accountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDoBindingJSON struct {
	ClassName   apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	Environment apijson.Field
	NamespaceID apijson.Field
	ScriptName  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDoBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDoBinding) implementsAccountWorkerServiceEnvironmentSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDoBindingType string

const (
	AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDoBindingTypeDurableObjectNamespace AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDoBindingType = "durable_object_namespace"
)

type AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uR2Binding struct {
	// R2 bucket to bind to
	BucketName string `json:"bucket_name,required"`
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uR2BindingType `json:"type,required"`
	JSON accountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uR2BindingJSON `json:"-"`
}

// accountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uR2BindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uR2Binding]
type accountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uR2BindingJSON struct {
	BucketName  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uR2Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uR2Binding) implementsAccountWorkerServiceEnvironmentSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uR2BindingType string

const (
	AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uR2BindingTypeR2Bucket AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uR2BindingType = "r2_bucket"
)

type AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uQueueBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Name of the Queue to bind to
	QueueName string `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uQueueBindingType `json:"type,required"`
	JSON accountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uQueueBindingJSON `json:"-"`
}

// accountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uQueueBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uQueueBinding]
type accountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uQueueBindingJSON struct {
	Name        apijson.Field
	QueueName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uQueueBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uQueueBinding) implementsAccountWorkerServiceEnvironmentSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uQueueBindingType string

const (
	AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uQueueBindingTypeQueue AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uQueueBindingType = "queue"
)

type AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uD1Binding struct {
	// ID of the D1 database to bind to
	ID string `json:"id,required"`
	// A JavaScript variable name for the binding.
	Binding string `json:"binding,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uD1BindingType `json:"type,required"`
	JSON accountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uD1BindingJSON `json:"-"`
}

// accountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uD1BindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uD1Binding]
type accountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uD1BindingJSON struct {
	ID          apijson.Field
	Binding     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uD1Binding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uD1Binding) implementsAccountWorkerServiceEnvironmentSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uD1BindingType string

const (
	AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uD1BindingTypeD1 AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uD1BindingType = "d1"
)

type AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// Namespace to bind to
	Namespace string `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingType `json:"type,required"`
	// Outbound worker
	Outbound AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutbound `json:"outbound"`
	JSON     accountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingJSON     `json:"-"`
}

// accountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBinding]
type accountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingJSON struct {
	Name        apijson.Field
	Namespace   apijson.Field
	Type        apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBinding) implementsAccountWorkerServiceEnvironmentSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingType string

const (
	AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingTypeDispatchNamespace AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params []string `json:"params"`
	// Outbound worker
	Worker AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker `json:"worker"`
	JSON   accountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundJSON   `json:"-"`
}

// accountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutbound]
type accountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundJSON struct {
	Params      apijson.Field
	Worker      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Outbound worker
type AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment string `json:"environment"`
	// Name of the outbound worker
	Service string                                                                                                         `json:"service"`
	JSON    accountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorkerJSON `json:"-"`
}

// accountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorkerJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker]
type accountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorkerJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uMtlsCertBinding struct {
	// A JavaScript variable name for the binding.
	Name string `json:"name,required"`
	// The class of resource that the binding provides.
	Type AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uMtlsCertBindingType `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID string                                                                                  `json:"certificate_id"`
	JSON          accountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uMtlsCertBindingJSON `json:"-"`
}

// accountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uMtlsCertBindingJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uMtlsCertBinding]
type accountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uMtlsCertBindingJSON struct {
	Name          apijson.Field
	Type          apijson.Field
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uMtlsCertBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uMtlsCertBinding) implementsAccountWorkerServiceEnvironmentSettingUpdateResponseBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uMtlsCertBindingType string

const (
	AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uMtlsCertBindingTypeMtlsCertificate AccountWorkerServiceEnvironmentSettingUpdateResponseBindingsAvYbsl2uMtlsCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by
// [AccountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrations]
// or
// [AccountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrations].
type AccountWorkerServiceEnvironmentSettingUpdateResponseMigrations interface {
	implementsAccountWorkerServiceEnvironmentSettingUpdateResponseMigrations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AccountWorkerServiceEnvironmentSettingUpdateResponseMigrations)(nil)).Elem(), "")
}

// A single set of migrations to apply.
type AccountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrations struct {
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
	RenamedClasses []AccountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []AccountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClass `json:"transferred_classes"`
	JSON               accountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsJSON               `json:"-"`
}

// accountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrations]
type accountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	NewTag             apijson.Field
	OldTag             apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrations) implementsAccountWorkerServiceEnvironmentSettingUpdateResponseMigrations() {
}

type AccountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClass struct {
	From string                                                                                                     `json:"from"`
	To   string                                                                                                     `json:"to"`
	JSON accountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClassJSON `json:"-"`
}

// accountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClassJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClass]
type accountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClass struct {
	From       string                                                                                                         `json:"from"`
	FromScript string                                                                                                         `json:"from_script"`
	To         string                                                                                                         `json:"to"`
	JSON       accountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClassJSON `json:"-"`
}

// accountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClassJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClass]
type accountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSingleStepMigrationsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// Migrations to apply in order.
	Steps []AccountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStep `json:"steps"`
	JSON  accountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsJSON   `json:"-"`
}

// accountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrations]
type accountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsJSON struct {
	NewTag      apijson.Field
	OldTag      apijson.Field
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AccountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrations) implementsAccountWorkerServiceEnvironmentSettingUpdateResponseMigrations() {
}

type AccountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses []string `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses []string `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses []AccountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses []AccountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass `json:"transferred_classes"`
	JSON               accountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepJSON                `json:"-"`
}

// accountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStep]
type accountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepJSON struct {
	DeletedClasses     apijson.Field
	NewClasses         apijson.Field
	RenamedClasses     apijson.Field
	TransferredClasses apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass struct {
	From string                                                                                                       `json:"from"`
	To   string                                                                                                       `json:"to"`
	JSON accountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClassJSON `json:"-"`
}

// accountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClassJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass]
type accountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClassJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass struct {
	From       string                                                                                                           `json:"from"`
	FromScript string                                                                                                           `json:"from_script"`
	To         string                                                                                                           `json:"to"`
	JSON       accountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClassJSON `json:"-"`
}

// accountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClassJSON
// contains the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass]
type accountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClassJSON struct {
	From        apijson.Field
	FromScript  apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingUpdateResponseMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerServiceEnvironmentSettingUpdateResponsePlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode AccountWorkerServiceEnvironmentSettingUpdateResponsePlacementMode `json:"mode"`
	JSON accountWorkerServiceEnvironmentSettingUpdateResponsePlacementJSON `json:"-"`
}

// accountWorkerServiceEnvironmentSettingUpdateResponsePlacementJSON contains the
// JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingUpdateResponsePlacement]
type accountWorkerServiceEnvironmentSettingUpdateResponsePlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingUpdateResponsePlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type AccountWorkerServiceEnvironmentSettingUpdateResponsePlacementMode string

const (
	AccountWorkerServiceEnvironmentSettingUpdateResponsePlacementModeSmart AccountWorkerServiceEnvironmentSettingUpdateResponsePlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type AccountWorkerServiceEnvironmentSettingUpdateResponseTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                                                               `json:"namespace"`
	JSON      accountWorkerServiceEnvironmentSettingUpdateResponseTailConsumerJSON `json:"-"`
}

// accountWorkerServiceEnvironmentSettingUpdateResponseTailConsumerJSON contains
// the JSON metadata for the struct
// [AccountWorkerServiceEnvironmentSettingUpdateResponseTailConsumer]
type accountWorkerServiceEnvironmentSettingUpdateResponseTailConsumerJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerServiceEnvironmentSettingUpdateResponseTailConsumer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerServiceEnvironmentSettingUpdateParams struct {
	// List of bindings attached to this Worker
	Bindings param.Field[[]AccountWorkerServiceEnvironmentSettingUpdateParamsBinding] `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations param.Field[AccountWorkerServiceEnvironmentSettingUpdateParamsMigrations] `json:"migrations"`
	Placement  param.Field[AccountWorkerServiceEnvironmentSettingUpdateParamsPlacement]  `json:"placement"`
	// Tags to help you manage your Workers
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]AccountWorkerServiceEnvironmentSettingUpdateParamsTailConsumer] `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel param.Field[string] `json:"usage_model"`
}

func (r AccountWorkerServiceEnvironmentSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A binding to allow the Worker to communicate with resources
//
// Satisfied by
// [AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBinding],
// [AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uServiceBinding],
// [AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uDoBinding],
// [AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uR2Binding],
// [AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uQueueBinding],
// [AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uD1Binding],
// [AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uDispatchNamespaceBinding],
// [AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uMtlsCertBinding].
type AccountWorkerServiceEnvironmentSettingUpdateParamsBinding interface {
	implementsAccountWorkerServiceEnvironmentSettingUpdateParamsBinding()
}

type AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBinding struct {
	// The class of resource that the binding provides.
	Type param.Field[AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBindingType] `json:"type,required"`
}

func (r AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBinding) implementsAccountWorkerServiceEnvironmentSettingUpdateParamsBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBindingType string

const (
	AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBindingTypeKvNamespace AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uKvNamespaceBindingType = "kv_namespace"
)

type AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uServiceBinding struct {
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment,required"`
	// Name of Worker to bind to
	Service param.Field[string] `json:"service,required"`
	// The class of resource that the binding provides.
	Type param.Field[AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uServiceBindingType] `json:"type,required"`
}

func (r AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uServiceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uServiceBinding) implementsAccountWorkerServiceEnvironmentSettingUpdateParamsBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uServiceBindingType string

const (
	AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uServiceBindingTypeService AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uServiceBindingType = "service"
)

type AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uDoBinding struct {
	// The exported class name of the Durable Object
	ClassName param.Field[string] `json:"class_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uDoBindingType] `json:"type,required"`
	// The environment of the script_name to bind to
	Environment param.Field[string] `json:"environment"`
	// The script where the Durable Object is defined, if it is external to this Worker
	ScriptName param.Field[string] `json:"script_name"`
}

func (r AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uDoBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uDoBinding) implementsAccountWorkerServiceEnvironmentSettingUpdateParamsBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uDoBindingType string

const (
	AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uDoBindingTypeDurableObjectNamespace AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uDoBindingType = "durable_object_namespace"
)

type AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uR2Binding struct {
	// R2 bucket to bind to
	BucketName param.Field[string] `json:"bucket_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uR2BindingType] `json:"type,required"`
}

func (r AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uR2Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uR2Binding) implementsAccountWorkerServiceEnvironmentSettingUpdateParamsBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uR2BindingType string

const (
	AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uR2BindingTypeR2Bucket AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uR2BindingType = "r2_bucket"
)

type AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uQueueBinding struct {
	// Name of the Queue to bind to
	QueueName param.Field[string] `json:"queue_name,required"`
	// The class of resource that the binding provides.
	Type param.Field[AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uQueueBindingType] `json:"type,required"`
}

func (r AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uQueueBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uQueueBinding) implementsAccountWorkerServiceEnvironmentSettingUpdateParamsBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uQueueBindingType string

const (
	AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uQueueBindingTypeQueue AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uQueueBindingType = "queue"
)

type AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uD1Binding struct {
	// ID of the D1 database to bind to
	ID param.Field[string] `json:"id,required"`
	// The name of the D1 database associated with the 'id' provided.
	Name param.Field[string] `json:"name,required"`
	// The class of resource that the binding provides.
	Type param.Field[AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uD1BindingType] `json:"type,required"`
}

func (r AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uD1Binding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uD1Binding) implementsAccountWorkerServiceEnvironmentSettingUpdateParamsBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uD1BindingType string

const (
	AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uD1BindingTypeD1 AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uD1BindingType = "d1"
)

type AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uDispatchNamespaceBinding struct {
	// Namespace to bind to
	Namespace param.Field[string] `json:"namespace,required"`
	// The class of resource that the binding provides.
	Type param.Field[AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uDispatchNamespaceBindingType] `json:"type,required"`
	// Outbound worker
	Outbound param.Field[AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uDispatchNamespaceBindingOutbound] `json:"outbound"`
}

func (r AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uDispatchNamespaceBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uDispatchNamespaceBinding) implementsAccountWorkerServiceEnvironmentSettingUpdateParamsBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uDispatchNamespaceBindingType string

const (
	AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uDispatchNamespaceBindingTypeDispatchNamespace AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uDispatchNamespaceBindingType = "dispatch_namespace"
)

// Outbound worker
type AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uDispatchNamespaceBindingOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the
	// parameters
	Params param.Field[[]string] `json:"params"`
	// Outbound worker
	Worker param.Field[AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker] `json:"worker"`
}

func (r AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uDispatchNamespaceBindingOutbound) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Outbound worker
type AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker struct {
	// Environment of the outbound worker
	Environment param.Field[string] `json:"environment"`
	// Name of the outbound worker
	Service param.Field[string] `json:"service"`
}

func (r AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uDispatchNamespaceBindingOutboundWorker) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uMtlsCertBinding struct {
	// The class of resource that the binding provides.
	Type param.Field[AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uMtlsCertBindingType] `json:"type,required"`
	// ID of the certificate to bind to
	CertificateID param.Field[string] `json:"certificate_id"`
}

func (r AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uMtlsCertBinding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uMtlsCertBinding) implementsAccountWorkerServiceEnvironmentSettingUpdateParamsBinding() {
}

// The class of resource that the binding provides.
type AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uMtlsCertBindingType string

const (
	AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uMtlsCertBindingTypeMtlsCertificate AccountWorkerServiceEnvironmentSettingUpdateParamsBindingsAvYbsl2uMtlsCertBindingType = "mtls_certificate"
)

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by
// [AccountWorkerServiceEnvironmentSettingUpdateParamsMigrationsAvYbsl2uSingleStepMigrations],
// [AccountWorkerServiceEnvironmentSettingUpdateParamsMigrationsAvYbsl2uSteppedMigrations].
type AccountWorkerServiceEnvironmentSettingUpdateParamsMigrations interface {
	implementsAccountWorkerServiceEnvironmentSettingUpdateParamsMigrations()
}

// A single set of migrations to apply.
type AccountWorkerServiceEnvironmentSettingUpdateParamsMigrationsAvYbsl2uSingleStepMigrations struct {
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
	RenamedClasses param.Field[[]AccountWorkerServiceEnvironmentSettingUpdateParamsMigrationsAvYbsl2uSingleStepMigrationsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]AccountWorkerServiceEnvironmentSettingUpdateParamsMigrationsAvYbsl2uSingleStepMigrationsTransferredClass] `json:"transferred_classes"`
}

func (r AccountWorkerServiceEnvironmentSettingUpdateParamsMigrationsAvYbsl2uSingleStepMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerServiceEnvironmentSettingUpdateParamsMigrationsAvYbsl2uSingleStepMigrations) implementsAccountWorkerServiceEnvironmentSettingUpdateParamsMigrations() {
}

type AccountWorkerServiceEnvironmentSettingUpdateParamsMigrationsAvYbsl2uSingleStepMigrationsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r AccountWorkerServiceEnvironmentSettingUpdateParamsMigrationsAvYbsl2uSingleStepMigrationsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerServiceEnvironmentSettingUpdateParamsMigrationsAvYbsl2uSingleStepMigrationsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r AccountWorkerServiceEnvironmentSettingUpdateParamsMigrationsAvYbsl2uSingleStepMigrationsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerServiceEnvironmentSettingUpdateParamsMigrationsAvYbsl2uSteppedMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag param.Field[string] `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag param.Field[string] `json:"old_tag"`
	// Migrations to apply in order.
	Steps param.Field[[]AccountWorkerServiceEnvironmentSettingUpdateParamsMigrationsAvYbsl2uSteppedMigrationsStep] `json:"steps"`
}

func (r AccountWorkerServiceEnvironmentSettingUpdateParamsMigrationsAvYbsl2uSteppedMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountWorkerServiceEnvironmentSettingUpdateParamsMigrationsAvYbsl2uSteppedMigrations) implementsAccountWorkerServiceEnvironmentSettingUpdateParamsMigrations() {
}

type AccountWorkerServiceEnvironmentSettingUpdateParamsMigrationsAvYbsl2uSteppedMigrationsStep struct {
	// A list of classes to delete Durable Object namespaces from.
	DeletedClasses param.Field[[]string] `json:"deleted_classes"`
	// A list of classes to create Durable Object namespaces from.
	NewClasses param.Field[[]string] `json:"new_classes"`
	// A list of classes with Durable Object namespaces that were renamed.
	RenamedClasses param.Field[[]AccountWorkerServiceEnvironmentSettingUpdateParamsMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass] `json:"renamed_classes"`
	// A list of transfers for Durable Object namespaces from a different Worker and
	// class to a class defined in this Worker.
	TransferredClasses param.Field[[]AccountWorkerServiceEnvironmentSettingUpdateParamsMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass] `json:"transferred_classes"`
}

func (r AccountWorkerServiceEnvironmentSettingUpdateParamsMigrationsAvYbsl2uSteppedMigrationsStep) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerServiceEnvironmentSettingUpdateParamsMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass struct {
	From param.Field[string] `json:"from"`
	To   param.Field[string] `json:"to"`
}

func (r AccountWorkerServiceEnvironmentSettingUpdateParamsMigrationsAvYbsl2uSteppedMigrationsStepsRenamedClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerServiceEnvironmentSettingUpdateParamsMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass struct {
	From       param.Field[string] `json:"from"`
	FromScript param.Field[string] `json:"from_script"`
	To         param.Field[string] `json:"to"`
}

func (r AccountWorkerServiceEnvironmentSettingUpdateParamsMigrationsAvYbsl2uSteppedMigrationsStepsTransferredClass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerServiceEnvironmentSettingUpdateParamsPlacement struct {
	// Enables
	// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
	// Only `"smart"` is currently supported
	Mode param.Field[AccountWorkerServiceEnvironmentSettingUpdateParamsPlacementMode] `json:"mode"`
}

func (r AccountWorkerServiceEnvironmentSettingUpdateParamsPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables
// [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).
// Only `"smart"` is currently supported
type AccountWorkerServiceEnvironmentSettingUpdateParamsPlacementMode string

const (
	AccountWorkerServiceEnvironmentSettingUpdateParamsPlacementModeSmart AccountWorkerServiceEnvironmentSettingUpdateParamsPlacementMode = "smart"
)

// A reference to a script that will consume logs from the attached Worker.
type AccountWorkerServiceEnvironmentSettingUpdateParamsTailConsumer struct {
	// Name of Worker that is to be the consumer.
	Service param.Field[string] `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace param.Field[string] `json:"namespace"`
}

func (r AccountWorkerServiceEnvironmentSettingUpdateParamsTailConsumer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
