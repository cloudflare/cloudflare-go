// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/option"
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

func (r SingleStepMigrationParam) implementsWorkersScriptUpdateParamsMetadataMigrationsUnion() {}

func (r SingleStepMigrationParam) ImplementsWorkersForPlatformsDispatchNamespaceScriptUpdateParamsMetadataMigrationsUnion() {
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
