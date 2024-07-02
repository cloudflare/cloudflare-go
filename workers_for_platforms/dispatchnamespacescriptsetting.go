// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/cloudflare/cloudflare-go/v2/workers"
	"github.com/tidwall/gjson"
)

// DispatchNamespaceScriptSettingService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDispatchNamespaceScriptSettingService] method instead.
type DispatchNamespaceScriptSettingService struct {
	Options []option.RequestOption
}

// NewDispatchNamespaceScriptSettingService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDispatchNamespaceScriptSettingService(opts ...option.RequestOption) (r *DispatchNamespaceScriptSettingService) {
	r = &DispatchNamespaceScriptSettingService{}
	r.Options = opts
	return
}

// Patch script metadata, such as bindings
func (r *DispatchNamespaceScriptSettingService) Edit(ctx context.Context, dispatchNamespace string, scriptName string, params DispatchNamespaceScriptSettingEditParams, opts ...option.RequestOption) (res *DispatchNamespaceScriptSettingEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DispatchNamespaceScriptSettingEditResponseEnvelope
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dispatchNamespace == "" {
		err = errors.New("missing required dispatch_namespace parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/settings", params.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get script settings from a script uploaded to a Workers for Platforms namespace.
func (r *DispatchNamespaceScriptSettingService) Get(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptSettingGetParams, opts ...option.RequestOption) (res *DispatchNamespaceScriptSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DispatchNamespaceScriptSettingGetResponseEnvelope
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dispatchNamespace == "" {
		err = errors.New("missing required dispatch_namespace parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/settings", query.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DispatchNamespaceScriptSettingEditResponse struct {
	// List of bindings attached to this Worker
	Bindings []workers.Binding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Limits to apply for this Worker.
	Limits DispatchNamespaceScriptSettingEditResponseLimits `json:"limits"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations DispatchNamespaceScriptSettingEditResponseMigrations `json:"migrations"`
	Placement  workers.PlacementConfiguration                       `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []workers.ConsumerScript `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                         `json:"usage_model"`
	JSON       dispatchNamespaceScriptSettingEditResponseJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseJSON contains the JSON metadata for
// the struct [DispatchNamespaceScriptSettingEditResponse]
type dispatchNamespaceScriptSettingEditResponseJSON struct {
	Bindings           apijson.Field
	CompatibilityDate  apijson.Field
	CompatibilityFlags apijson.Field
	Limits             apijson.Field
	Logpush            apijson.Field
	Migrations         apijson.Field
	Placement          apijson.Field
	Tags               apijson.Field
	TailConsumers      apijson.Field
	UsageModel         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseJSON) RawJSON() string {
	return r.raw
}

// Limits to apply for this Worker.
type DispatchNamespaceScriptSettingEditResponseLimits struct {
	// The amount of CPU time this Worker can use in milliseconds.
	CPUMs int64                                                `json:"cpu_ms"`
	JSON  dispatchNamespaceScriptSettingEditResponseLimitsJSON `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseLimitsJSON contains the JSON metadata
// for the struct [DispatchNamespaceScriptSettingEditResponseLimits]
type dispatchNamespaceScriptSettingEditResponseLimitsJSON struct {
	CPUMs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseLimitsJSON) RawJSON() string {
	return r.raw
}

// Migrations to apply for Durable Objects associated with this Worker.
type DispatchNamespaceScriptSettingEditResponseMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// This field can have the runtime type of [[]string].
	DeletedClasses interface{} `json:"deleted_classes,required"`
	// This field can have the runtime type of [[]string].
	NewClasses interface{} `json:"new_classes,required"`
	// This field can have the runtime type of
	// [[]workers.SingleStepMigrationRenamedClass].
	RenamedClasses interface{} `json:"renamed_classes,required"`
	// This field can have the runtime type of
	// [[]workers.SingleStepMigrationTransferredClass].
	TransferredClasses interface{} `json:"transferred_classes,required"`
	// This field can have the runtime type of [[]workers.MigrationStep].
	Steps interface{}                                              `json:"steps,required"`
	JSON  dispatchNamespaceScriptSettingEditResponseMigrationsJSON `json:"-"`
	union DispatchNamespaceScriptSettingEditResponseMigrationsUnion
}

// dispatchNamespaceScriptSettingEditResponseMigrationsJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptSettingEditResponseMigrations]
type dispatchNamespaceScriptSettingEditResponseMigrationsJSON struct {
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

func (r dispatchNamespaceScriptSettingEditResponseMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r *DispatchNamespaceScriptSettingEditResponseMigrations) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DispatchNamespaceScriptSettingEditResponseMigrationsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [workers.SingleStepMigration],
// [workers.SteppedMigration].
func (r DispatchNamespaceScriptSettingEditResponseMigrations) AsUnion() DispatchNamespaceScriptSettingEditResponseMigrationsUnion {
	return r.union
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by [workers.SingleStepMigration] or [workers.SteppedMigration].
type DispatchNamespaceScriptSettingEditResponseMigrationsUnion interface {
	ImplementsWorkersForPlatformsDispatchNamespaceScriptSettingEditResponseMigrations()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DispatchNamespaceScriptSettingEditResponseMigrationsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(workers.SingleStepMigration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(workers.SteppedMigration{}),
		},
	)
}

type DispatchNamespaceScriptSettingGetResponse struct {
	// List of bindings attached to this Worker
	Bindings []workers.Binding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []string `json:"compatibility_flags"`
	// Limits to apply for this Worker.
	Limits DispatchNamespaceScriptSettingGetResponseLimits `json:"limits"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations DispatchNamespaceScriptSettingGetResponseMigrations `json:"migrations"`
	Placement  workers.PlacementConfiguration                      `json:"placement"`
	// Tags to help you manage your Workers
	Tags []string `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []workers.ConsumerScript `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string                                        `json:"usage_model"`
	JSON       dispatchNamespaceScriptSettingGetResponseJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseJSON contains the JSON metadata for the
// struct [DispatchNamespaceScriptSettingGetResponse]
type dispatchNamespaceScriptSettingGetResponseJSON struct {
	Bindings           apijson.Field
	CompatibilityDate  apijson.Field
	CompatibilityFlags apijson.Field
	Limits             apijson.Field
	Logpush            apijson.Field
	Migrations         apijson.Field
	Placement          apijson.Field
	Tags               apijson.Field
	TailConsumers      apijson.Field
	UsageModel         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseJSON) RawJSON() string {
	return r.raw
}

// Limits to apply for this Worker.
type DispatchNamespaceScriptSettingGetResponseLimits struct {
	// The amount of CPU time this Worker can use in milliseconds.
	CPUMs int64                                               `json:"cpu_ms"`
	JSON  dispatchNamespaceScriptSettingGetResponseLimitsJSON `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseLimitsJSON contains the JSON metadata
// for the struct [DispatchNamespaceScriptSettingGetResponseLimits]
type dispatchNamespaceScriptSettingGetResponseLimitsJSON struct {
	CPUMs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseLimitsJSON) RawJSON() string {
	return r.raw
}

// Migrations to apply for Durable Objects associated with this Worker.
type DispatchNamespaceScriptSettingGetResponseMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag string `json:"old_tag"`
	// This field can have the runtime type of [[]string].
	DeletedClasses interface{} `json:"deleted_classes,required"`
	// This field can have the runtime type of [[]string].
	NewClasses interface{} `json:"new_classes,required"`
	// This field can have the runtime type of
	// [[]workers.SingleStepMigrationRenamedClass].
	RenamedClasses interface{} `json:"renamed_classes,required"`
	// This field can have the runtime type of
	// [[]workers.SingleStepMigrationTransferredClass].
	TransferredClasses interface{} `json:"transferred_classes,required"`
	// This field can have the runtime type of [[]workers.MigrationStep].
	Steps interface{}                                             `json:"steps,required"`
	JSON  dispatchNamespaceScriptSettingGetResponseMigrationsJSON `json:"-"`
	union DispatchNamespaceScriptSettingGetResponseMigrationsUnion
}

// dispatchNamespaceScriptSettingGetResponseMigrationsJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptSettingGetResponseMigrations]
type dispatchNamespaceScriptSettingGetResponseMigrationsJSON struct {
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

func (r dispatchNamespaceScriptSettingGetResponseMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r *DispatchNamespaceScriptSettingGetResponseMigrations) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DispatchNamespaceScriptSettingGetResponseMigrationsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [workers.SingleStepMigration],
// [workers.SteppedMigration].
func (r DispatchNamespaceScriptSettingGetResponseMigrations) AsUnion() DispatchNamespaceScriptSettingGetResponseMigrationsUnion {
	return r.union
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by [workers.SingleStepMigration] or [workers.SteppedMigration].
type DispatchNamespaceScriptSettingGetResponseMigrationsUnion interface {
	ImplementsWorkersForPlatformsDispatchNamespaceScriptSettingGetResponseMigrations()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DispatchNamespaceScriptSettingGetResponseMigrationsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(workers.SingleStepMigration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(workers.SteppedMigration{}),
		},
	)
}

type DispatchNamespaceScriptSettingEditParams struct {
	// Identifier
	AccountID param.Field[string]                                           `path:"account_id,required"`
	Settings  param.Field[DispatchNamespaceScriptSettingEditParamsSettings] `json:"settings"`
}

func (r DispatchNamespaceScriptSettingEditParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type DispatchNamespaceScriptSettingEditParamsSettings struct {
	// List of bindings attached to this Worker
	Bindings param.Field[[]workers.BindingUnionParam] `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// Limits to apply for this Worker.
	Limits param.Field[DispatchNamespaceScriptSettingEditParamsSettingsLimits] `json:"limits"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations param.Field[DispatchNamespaceScriptSettingEditParamsSettingsMigrationsUnion] `json:"migrations"`
	Placement  param.Field[workers.PlacementConfigurationParam]                             `json:"placement"`
	// Tags to help you manage your Workers
	Tags param.Field[[]string] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]workers.ConsumerScriptParam] `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel param.Field[string] `json:"usage_model"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Limits to apply for this Worker.
type DispatchNamespaceScriptSettingEditParamsSettingsLimits struct {
	// The amount of CPU time this Worker can use in milliseconds.
	CPUMs param.Field[int64] `json:"cpu_ms"`
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsLimits) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Migrations to apply for Durable Objects associated with this Worker.
type DispatchNamespaceScriptSettingEditParamsSettingsMigrations struct {
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

func (r DispatchNamespaceScriptSettingEditParamsSettingsMigrations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DispatchNamespaceScriptSettingEditParamsSettingsMigrations) ImplementsWorkersForPlatformsDispatchNamespaceScriptSettingEditParamsSettingsMigrationsUnion() {
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by [workers.SingleStepMigrationParam],
// [workers.SteppedMigrationParam],
// [DispatchNamespaceScriptSettingEditParamsSettingsMigrations].
type DispatchNamespaceScriptSettingEditParamsSettingsMigrationsUnion interface {
	ImplementsWorkersForPlatformsDispatchNamespaceScriptSettingEditParamsSettingsMigrationsUnion()
}

type DispatchNamespaceScriptSettingEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DispatchNamespaceScriptSettingEditResponseEnvelopeSuccess `json:"success,required"`
	Result  DispatchNamespaceScriptSettingEditResponse                `json:"result"`
	JSON    dispatchNamespaceScriptSettingEditResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptSettingEditResponseEnvelopeJSON contains the JSON
// metadata for the struct [DispatchNamespaceScriptSettingEditResponseEnvelope]
type dispatchNamespaceScriptSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DispatchNamespaceScriptSettingEditResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptSettingEditResponseEnvelopeSuccessTrue DispatchNamespaceScriptSettingEditResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptSettingEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DispatchNamespaceScriptSettingGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DispatchNamespaceScriptSettingGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success DispatchNamespaceScriptSettingGetResponseEnvelopeSuccess `json:"success,required"`
	Result  DispatchNamespaceScriptSettingGetResponse                `json:"result"`
	JSON    dispatchNamespaceScriptSettingGetResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceScriptSettingGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [DispatchNamespaceScriptSettingGetResponseEnvelope]
type dispatchNamespaceScriptSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptSettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DispatchNamespaceScriptSettingGetResponseEnvelopeSuccess bool

const (
	DispatchNamespaceScriptSettingGetResponseEnvelopeSuccessTrue DispatchNamespaceScriptSettingGetResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceScriptSettingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptSettingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
