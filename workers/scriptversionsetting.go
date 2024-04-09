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
func (r *ScriptVersionSettingService) Edit(ctx context.Context, scriptName string, params ScriptVersionSettingEditParams, opts ...option.RequestOption) (res *Settings, err error) {
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
func (r *ScriptVersionSettingService) Get(ctx context.Context, scriptName string, query ScriptVersionSettingGetParams, opts ...option.RequestOption) (res *Settings, err error) {
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

type CompatibilityFlags = string

type CompatibilityFlagsParam = string

type Settings struct {
	// List of bindings attached to this Worker
	Bindings []Binding `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate string `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags []CompatibilityFlags `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations SettingsMigrations     `json:"migrations"`
	Placement  PlacementConfiguration `json:"placement"`
	// Tags to help you manage your Workers
	Tags []Tags `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []ConsumerScriptItem `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel string       `json:"usage_model"`
	JSON       settingsJSON `json:"-"`
}

// settingsJSON contains the JSON metadata for the struct [Settings]
type settingsJSON struct {
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

func (r *Settings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingsJSON) RawJSON() string {
	return r.raw
}

// Migrations to apply for Durable Objects associated with this Worker.
type SettingsMigrations struct {
	// Tag to set as the latest migration tag.
	NewTag string `json:"new_tag"`
	// Tag used to verify against the latest migration tag for this Worker. If they
	// don't match, the upload is rejected.
	OldTag             string                 `json:"old_tag"`
	DeletedClasses     interface{}            `json:"deleted_classes,required"`
	NewClasses         interface{}            `json:"new_classes,required"`
	RenamedClasses     interface{}            `json:"renamed_classes,required"`
	TransferredClasses interface{}            `json:"transferred_classes,required"`
	Steps              interface{}            `json:"steps,required"`
	JSON               settingsMigrationsJSON `json:"-"`
	union              SettingsMigrationsUnion
}

// settingsMigrationsJSON contains the JSON metadata for the struct
// [SettingsMigrations]
type settingsMigrationsJSON struct {
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

func (r settingsMigrationsJSON) RawJSON() string {
	return r.raw
}

func (r *SettingsMigrations) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r SettingsMigrations) AsUnion() SettingsMigrationsUnion {
	return r.union
}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Union satisfied by [workers.SingleStepMigration] or [workers.SteppedMigration].
type SettingsMigrationsUnion interface {
	implementsWorkersSettingsMigrations()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SettingsMigrationsUnion)(nil)).Elem(),
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

type SettingsParam struct {
	// List of bindings attached to this Worker
	Bindings param.Field[[]BindingUnionParam] `json:"bindings"`
	// Opt your Worker into changes after this date
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Opt your Worker into specific changes
	CompatibilityFlags param.Field[[]CompatibilityFlagsParam] `json:"compatibility_flags"`
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Migrations to apply for Durable Objects associated with this Worker.
	Migrations param.Field[SettingsMigrationsUnionParam] `json:"migrations"`
	Placement  param.Field[PlacementConfigurationParam]  `json:"placement"`
	// Tags to help you manage your Workers
	Tags param.Field[[]TagsParam] `json:"tags"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]ConsumerScriptItemParam] `json:"tail_consumers"`
	// Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').
	UsageModel param.Field[string] `json:"usage_model"`
}

func (r SettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Migrations to apply for Durable Objects associated with this Worker.
type SettingsMigrationsParam struct {
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

func (r SettingsMigrationsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingsMigrationsParam) implementsWorkersSettingsMigrationsUnionParam() {}

// Migrations to apply for Durable Objects associated with this Worker.
//
// Satisfied by [workers.SingleStepMigrationParam],
// [workers.SteppedMigrationParam], [SettingsMigrationsParam].
type SettingsMigrationsUnionParam interface {
	implementsWorkersSettingsMigrationsUnionParam()
}

type Tags = string

type TagsParam = string

type ScriptVersionSettingEditParams struct {
	// Identifier
	AccountID param.Field[string]        `path:"account_id,required"`
	Settings  param.Field[SettingsParam] `json:"settings"`
}

func (r ScriptVersionSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScriptVersionSettingEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   Settings                                                  `json:"result,required"`
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
	Result   Settings                                                  `json:"result,required"`
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
