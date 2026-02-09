// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
	"github.com/tidwall/gjson"
)

// ProjectService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectService] method instead.
type ProjectService struct {
	Options     []option.RequestOption
	Deployments *ProjectDeploymentService
	Domains     *ProjectDomainService
}

// NewProjectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProjectService(opts ...option.RequestOption) (r *ProjectService) {
	r = &ProjectService{}
	r.Options = opts
	r.Deployments = NewProjectDeploymentService(opts...)
	r.Domains = NewProjectDomainService(opts...)
	return
}

// Create a new project.
func (r *ProjectService) New(ctx context.Context, params ProjectNewParams, opts ...option.RequestOption) (res *Project, err error) {
	var env ProjectNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a list of all user projects.
func (r *ProjectService) List(ctx context.Context, params ProjectListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[Project], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Fetch a list of all user projects.
func (r *ProjectService) ListAutoPaging(ctx context.Context, params ProjectListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[Project] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete a project by name.
func (r *ProjectService) Delete(ctx context.Context, projectName string, body ProjectDeleteParams, opts ...option.RequestOption) (res *ProjectDeleteResponse, err error) {
	var env ProjectDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s", body.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Set new attributes for an existing project. Modify environment variables. To
// delete an environment variable, set the key to null.
func (r *ProjectService) Edit(ctx context.Context, projectName string, params ProjectEditParams, opts ...option.RequestOption) (res *Project, err error) {
	var env ProjectEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s", params.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a project by name.
func (r *ProjectService) Get(ctx context.Context, projectName string, query ProjectGetParams, opts ...option.RequestOption) (res *Project, err error) {
	var env ProjectGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s", query.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Purge all cached build artifacts for a Pages project
func (r *ProjectService) PurgeBuildCache(ctx context.Context, projectName string, body ProjectPurgeBuildCacheParams, opts ...option.RequestOption) (res *ProjectPurgeBuildCacheResponse, err error) {
	var env ProjectPurgeBuildCacheResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/purge_build_cache", body.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Deployment struct {
	// Id of the deployment.
	ID string `json:"id,required"`
	// A list of alias URLs pointing to this deployment.
	Aliases []string `json:"aliases,required,nullable"`
	// Configs for the project build process.
	BuildConfig DeploymentBuildConfig `json:"build_config,required"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger DeploymentDeploymentTrigger `json:"deployment_trigger,required"`
	// Environment variables used for builds and Pages Functions.
	EnvVars map[string]DeploymentEnvVar `json:"env_vars,required,nullable"`
	// Type of deploy.
	Environment DeploymentEnvironment `json:"environment,required"`
	// If the deployment has been skipped.
	IsSkipped bool `json:"is_skipped,required"`
	// The status of the deployment.
	LatestStage Stage `json:"latest_stage,required"`
	// When the deployment was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Id of the project.
	ProjectID string `json:"project_id,required"`
	// Name of the project.
	ProjectName string `json:"project_name,required"`
	// Short Id (8 character) of the deployment.
	ShortID string `json:"short_id,required"`
	// Configs for the project source control.
	Source DeploymentSource `json:"source,required"`
	// List of past stages.
	Stages []Stage `json:"stages,required"`
	// The live URL to view this deployment.
	URL string `json:"url,required"`
	// Whether the deployment uses functions.
	UsesFunctions bool           `json:"uses_functions,nullable"`
	JSON          deploymentJSON `json:"-"`
}

// deploymentJSON contains the JSON metadata for the struct [Deployment]
type deploymentJSON struct {
	ID                apijson.Field
	Aliases           apijson.Field
	BuildConfig       apijson.Field
	CreatedOn         apijson.Field
	DeploymentTrigger apijson.Field
	EnvVars           apijson.Field
	Environment       apijson.Field
	IsSkipped         apijson.Field
	LatestStage       apijson.Field
	ModifiedOn        apijson.Field
	ProjectID         apijson.Field
	ProjectName       apijson.Field
	ShortID           apijson.Field
	Source            apijson.Field
	Stages            apijson.Field
	URL               apijson.Field
	UsesFunctions     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Deployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentJSON) RawJSON() string {
	return r.raw
}

// Configs for the project build process.
type DeploymentBuildConfig struct {
	// The classifying tag for analytics.
	WebAnalyticsTag string `json:"web_analytics_tag,required,nullable"`
	// The auth token for analytics.
	WebAnalyticsToken string `json:"web_analytics_token,required,nullable"`
	// Enable build caching for the project.
	BuildCaching bool `json:"build_caching,nullable"`
	// Command used to build project.
	BuildCommand string `json:"build_command,nullable"`
	// Assets output directory of the build.
	DestinationDir string `json:"destination_dir,nullable"`
	// Directory to run the command.
	RootDir string                    `json:"root_dir,nullable"`
	JSON    deploymentBuildConfigJSON `json:"-"`
}

// deploymentBuildConfigJSON contains the JSON metadata for the struct
// [DeploymentBuildConfig]
type deploymentBuildConfigJSON struct {
	WebAnalyticsTag   apijson.Field
	WebAnalyticsToken apijson.Field
	BuildCaching      apijson.Field
	BuildCommand      apijson.Field
	DestinationDir    apijson.Field
	RootDir           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DeploymentBuildConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentBuildConfigJSON) RawJSON() string {
	return r.raw
}

// Info about what caused the deployment.
type DeploymentDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata DeploymentDeploymentTriggerMetadata `json:"metadata,required"`
	// What caused the deployment.
	Type DeploymentDeploymentTriggerType `json:"type,required"`
	JSON deploymentDeploymentTriggerJSON `json:"-"`
}

// deploymentDeploymentTriggerJSON contains the JSON metadata for the struct
// [DeploymentDeploymentTrigger]
type deploymentDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentDeploymentTriggerJSON) RawJSON() string {
	return r.raw
}

// Additional info about the trigger.
type DeploymentDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch,required"`
	// Whether the deployment trigger commit was dirty.
	CommitDirty bool `json:"commit_dirty,required"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash,required"`
	// Message of the deployment trigger commit.
	CommitMessage string                                  `json:"commit_message,required"`
	JSON          deploymentDeploymentTriggerMetadataJSON `json:"-"`
}

// deploymentDeploymentTriggerMetadataJSON contains the JSON metadata for the
// struct [DeploymentDeploymentTriggerMetadata]
type deploymentDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitDirty   apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DeploymentDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentDeploymentTriggerMetadataJSON) RawJSON() string {
	return r.raw
}

// What caused the deployment.
type DeploymentDeploymentTriggerType string

const (
	DeploymentDeploymentTriggerTypeGitHubPush DeploymentDeploymentTriggerType = "github:push"
	DeploymentDeploymentTriggerTypeADHoc      DeploymentDeploymentTriggerType = "ad_hoc"
	DeploymentDeploymentTriggerTypeDeployHook DeploymentDeploymentTriggerType = "deploy_hook"
)

func (r DeploymentDeploymentTriggerType) IsKnown() bool {
	switch r {
	case DeploymentDeploymentTriggerTypeGitHubPush, DeploymentDeploymentTriggerTypeADHoc, DeploymentDeploymentTriggerTypeDeployHook:
		return true
	}
	return false
}

// A plaintext environment variable.
type DeploymentEnvVar struct {
	Type DeploymentEnvVarsType `json:"type,required"`
	// Environment variable value.
	Value string               `json:"value,required"`
	JSON  deploymentEnvVarJSON `json:"-"`
	union DeploymentEnvVarsUnion
}

// deploymentEnvVarJSON contains the JSON metadata for the struct
// [DeploymentEnvVar]
type deploymentEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r deploymentEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r *DeploymentEnvVar) UnmarshalJSON(data []byte) (err error) {
	*r = DeploymentEnvVar{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DeploymentEnvVarsUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [DeploymentEnvVarsPagesPlainTextEnvVar],
// [DeploymentEnvVarsPagesSecretTextEnvVar].
func (r DeploymentEnvVar) AsUnion() DeploymentEnvVarsUnion {
	return r.union
}

// A plaintext environment variable.
//
// Union satisfied by [DeploymentEnvVarsPagesPlainTextEnvVar] or
// [DeploymentEnvVarsPagesSecretTextEnvVar].
type DeploymentEnvVarsUnion interface {
	implementsDeploymentEnvVar()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DeploymentEnvVarsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DeploymentEnvVarsPagesPlainTextEnvVar{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DeploymentEnvVarsPagesSecretTextEnvVar{}),
			DiscriminatorValue: "secret_text",
		},
	)
}

// A plaintext environment variable.
type DeploymentEnvVarsPagesPlainTextEnvVar struct {
	Type DeploymentEnvVarsPagesPlainTextEnvVarType `json:"type,required"`
	// Environment variable value.
	Value string                                    `json:"value,required"`
	JSON  deploymentEnvVarsPagesPlainTextEnvVarJSON `json:"-"`
}

// deploymentEnvVarsPagesPlainTextEnvVarJSON contains the JSON metadata for the
// struct [DeploymentEnvVarsPagesPlainTextEnvVar]
type deploymentEnvVarsPagesPlainTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentEnvVarsPagesPlainTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentEnvVarsPagesPlainTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r DeploymentEnvVarsPagesPlainTextEnvVar) implementsDeploymentEnvVar() {}

type DeploymentEnvVarsPagesPlainTextEnvVarType string

const (
	DeploymentEnvVarsPagesPlainTextEnvVarTypePlainText DeploymentEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r DeploymentEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case DeploymentEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type DeploymentEnvVarsPagesSecretTextEnvVar struct {
	Type DeploymentEnvVarsPagesSecretTextEnvVarType `json:"type,required"`
	// Secret value.
	Value string                                     `json:"value,required"`
	JSON  deploymentEnvVarsPagesSecretTextEnvVarJSON `json:"-"`
}

// deploymentEnvVarsPagesSecretTextEnvVarJSON contains the JSON metadata for the
// struct [DeploymentEnvVarsPagesSecretTextEnvVar]
type deploymentEnvVarsPagesSecretTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentEnvVarsPagesSecretTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentEnvVarsPagesSecretTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r DeploymentEnvVarsPagesSecretTextEnvVar) implementsDeploymentEnvVar() {}

type DeploymentEnvVarsPagesSecretTextEnvVarType string

const (
	DeploymentEnvVarsPagesSecretTextEnvVarTypeSecretText DeploymentEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r DeploymentEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case DeploymentEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type DeploymentEnvVarsType string

const (
	DeploymentEnvVarsTypePlainText  DeploymentEnvVarsType = "plain_text"
	DeploymentEnvVarsTypeSecretText DeploymentEnvVarsType = "secret_text"
)

func (r DeploymentEnvVarsType) IsKnown() bool {
	switch r {
	case DeploymentEnvVarsTypePlainText, DeploymentEnvVarsTypeSecretText:
		return true
	}
	return false
}

// Type of deploy.
type DeploymentEnvironment string

const (
	DeploymentEnvironmentPreview    DeploymentEnvironment = "preview"
	DeploymentEnvironmentProduction DeploymentEnvironment = "production"
)

func (r DeploymentEnvironment) IsKnown() bool {
	switch r {
	case DeploymentEnvironmentPreview, DeploymentEnvironmentProduction:
		return true
	}
	return false
}

// Configs for the project source control.
type DeploymentSource struct {
	Config DeploymentSourceConfig `json:"config,required"`
	// The source control management provider.
	Type DeploymentSourceType `json:"type,required"`
	JSON deploymentSourceJSON `json:"-"`
}

// deploymentSourceJSON contains the JSON metadata for the struct
// [DeploymentSource]
type deploymentSourceJSON struct {
	Config      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeploymentSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentSourceJSON) RawJSON() string {
	return r.raw
}

type DeploymentSourceConfig struct {
	// Whether to enable automatic deployments when pushing to the source repository.
	// When disabled, no deployments (production or preview) will be triggered
	// automatically.
	//
	// Deprecated: Use `production_deployments_enabled` and
	// `preview_deployment_setting` for more granular control.
	DeploymentsEnabled bool `json:"deployments_enabled,required"`
	// The owner of the repository.
	Owner string `json:"owner,required"`
	// The owner ID of the repository.
	OwnerID string `json:"owner_id,required"`
	// A list of paths that should be excluded from triggering a preview deployment.
	// Wildcard syntax (`*`) is supported.
	PathExcludes []string `json:"path_excludes,required"`
	// A list of paths that should be watched to trigger a preview deployment. Wildcard
	// syntax (`*`) is supported.
	PathIncludes []string `json:"path_includes,required"`
	// Whether to enable PR comments.
	PrCommentsEnabled bool `json:"pr_comments_enabled,required"`
	// A list of branches that should not trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchExcludes []string `json:"preview_branch_excludes,required"`
	// A list of branches that should trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchIncludes []string `json:"preview_branch_includes,required"`
	// Controls whether commits to preview branches trigger a preview deployment.
	PreviewDeploymentSetting DeploymentSourceConfigPreviewDeploymentSetting `json:"preview_deployment_setting,required"`
	// The production branch of the repository.
	ProductionBranch string `json:"production_branch,required"`
	// Whether to trigger a production deployment on commits to the production branch.
	ProductionDeploymentsEnabled bool `json:"production_deployments_enabled,required"`
	// The ID of the repository.
	RepoID string `json:"repo_id,required"`
	// The name of the repository.
	RepoName string                     `json:"repo_name,required"`
	JSON     deploymentSourceConfigJSON `json:"-"`
}

// deploymentSourceConfigJSON contains the JSON metadata for the struct
// [DeploymentSourceConfig]
type deploymentSourceConfigJSON struct {
	DeploymentsEnabled           apijson.Field
	Owner                        apijson.Field
	OwnerID                      apijson.Field
	PathExcludes                 apijson.Field
	PathIncludes                 apijson.Field
	PrCommentsEnabled            apijson.Field
	PreviewBranchExcludes        apijson.Field
	PreviewBranchIncludes        apijson.Field
	PreviewDeploymentSetting     apijson.Field
	ProductionBranch             apijson.Field
	ProductionDeploymentsEnabled apijson.Field
	RepoID                       apijson.Field
	RepoName                     apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *DeploymentSourceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deploymentSourceConfigJSON) RawJSON() string {
	return r.raw
}

// Controls whether commits to preview branches trigger a preview deployment.
type DeploymentSourceConfigPreviewDeploymentSetting string

const (
	DeploymentSourceConfigPreviewDeploymentSettingAll    DeploymentSourceConfigPreviewDeploymentSetting = "all"
	DeploymentSourceConfigPreviewDeploymentSettingNone   DeploymentSourceConfigPreviewDeploymentSetting = "none"
	DeploymentSourceConfigPreviewDeploymentSettingCustom DeploymentSourceConfigPreviewDeploymentSetting = "custom"
)

func (r DeploymentSourceConfigPreviewDeploymentSetting) IsKnown() bool {
	switch r {
	case DeploymentSourceConfigPreviewDeploymentSettingAll, DeploymentSourceConfigPreviewDeploymentSettingNone, DeploymentSourceConfigPreviewDeploymentSettingCustom:
		return true
	}
	return false
}

// The source control management provider.
type DeploymentSourceType string

const (
	DeploymentSourceTypeGitHub DeploymentSourceType = "github"
	DeploymentSourceTypeGitlab DeploymentSourceType = "gitlab"
)

func (r DeploymentSourceType) IsKnown() bool {
	switch r {
	case DeploymentSourceTypeGitHub, DeploymentSourceTypeGitlab:
		return true
	}
	return false
}

type Project struct {
	// ID of the project.
	ID string `json:"id,required"`
	// Most recent production deployment of the project.
	CanonicalDeployment Deployment `json:"canonical_deployment,required,nullable"`
	// When the project was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Configs for deployments in a project.
	DeploymentConfigs ProjectDeploymentConfigs `json:"deployment_configs,required"`
	// Framework the project is using.
	Framework string `json:"framework,required"`
	// Version of the framework the project is using.
	FrameworkVersion string `json:"framework_version,required"`
	// Most recent deployment of the project.
	LatestDeployment Deployment `json:"latest_deployment,required,nullable"`
	// Name of the project.
	Name string `json:"name,required"`
	// Name of the preview script.
	PreviewScriptName string `json:"preview_script_name,required"`
	// Production branch of the project. Used to identify production deployments.
	ProductionBranch string `json:"production_branch,required"`
	// Name of the production script.
	ProductionScriptName string `json:"production_script_name,required"`
	// Whether the project uses functions.
	UsesFunctions bool `json:"uses_functions,required,nullable"`
	// Configs for the project build process.
	BuildConfig ProjectBuildConfig `json:"build_config"`
	// A list of associated custom domains for the project.
	Domains []string `json:"domains"`
	// Configs for the project source control.
	Source ProjectSource `json:"source"`
	// The Cloudflare subdomain associated with the project.
	Subdomain string      `json:"subdomain"`
	JSON      projectJSON `json:"-"`
}

// projectJSON contains the JSON metadata for the struct [Project]
type projectJSON struct {
	ID                   apijson.Field
	CanonicalDeployment  apijson.Field
	CreatedOn            apijson.Field
	DeploymentConfigs    apijson.Field
	Framework            apijson.Field
	FrameworkVersion     apijson.Field
	LatestDeployment     apijson.Field
	Name                 apijson.Field
	PreviewScriptName    apijson.Field
	ProductionBranch     apijson.Field
	ProductionScriptName apijson.Field
	UsesFunctions        apijson.Field
	BuildConfig          apijson.Field
	Domains              apijson.Field
	Source               apijson.Field
	Subdomain            apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *Project) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectJSON) RawJSON() string {
	return r.raw
}

// Configs for deployments in a project.
type ProjectDeploymentConfigs struct {
	// Configs for preview deploys.
	Preview ProjectDeploymentConfigsPreview `json:"preview,required"`
	// Configs for production deploys.
	Production ProjectDeploymentConfigsProduction `json:"production,required"`
	JSON       projectDeploymentConfigsJSON       `json:"-"`
}

// projectDeploymentConfigsJSON contains the JSON metadata for the struct
// [ProjectDeploymentConfigs]
type projectDeploymentConfigsJSON struct {
	Preview     apijson.Field
	Production  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsJSON) RawJSON() string {
	return r.raw
}

// Configs for preview deploys.
type ProjectDeploymentConfigsPreview struct {
	// Whether to always use the latest compatibility date for Pages Functions.
	AlwaysUseLatestCompatibilityDate bool `json:"always_use_latest_compatibility_date,required"`
	// The major version of the build image to use for Pages Functions.
	BuildImageMajorVersion int64 `json:"build_image_major_version,required"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate string `json:"compatibility_date,required"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags []string `json:"compatibility_flags,required"`
	// Environment variables used for builds and Pages Functions.
	EnvVars map[string]ProjectDeploymentConfigsPreviewEnvVar `json:"env_vars,required,nullable"`
	// Whether to fail open when the deployment config cannot be applied.
	FailOpen bool `json:"fail_open,required"`
	// The usage model for Pages Functions.
	//
	// Deprecated: All new projects now use the Standard usage model.
	UsageModel ProjectDeploymentConfigsPreviewUsageModel `json:"usage_model,required"`
	// Constellation bindings used for Pages Functions.
	AIBindings map[string]ProjectDeploymentConfigsPreviewAIBinding `json:"ai_bindings"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets map[string]ProjectDeploymentConfigsPreviewAnalyticsEngineDataset `json:"analytics_engine_datasets"`
	// Browser bindings used for Pages Functions.
	Browsers map[string]ProjectDeploymentConfigsPreviewBrowser `json:"browsers"`
	// D1 databases used for Pages Functions.
	D1Databases map[string]ProjectDeploymentConfigsPreviewD1Database `json:"d1_databases"`
	// Durable Object namespaces used for Pages Functions.
	DurableObjectNamespaces map[string]ProjectDeploymentConfigsPreviewDurableObjectNamespace `json:"durable_object_namespaces"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings map[string]ProjectDeploymentConfigsPreviewHyperdriveBinding `json:"hyperdrive_bindings"`
	// KV namespaces used for Pages Functions.
	KVNamespaces map[string]ProjectDeploymentConfigsPreviewKVNamespace `json:"kv_namespaces"`
	// Limits for Pages Functions.
	Limits ProjectDeploymentConfigsPreviewLimits `json:"limits"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates map[string]ProjectDeploymentConfigsPreviewMTLSCertificate `json:"mtls_certificates"`
	// Placement setting used for Pages Functions.
	Placement ProjectDeploymentConfigsPreviewPlacement `json:"placement"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers map[string]ProjectDeploymentConfigsPreviewQueueProducer `json:"queue_producers"`
	// R2 buckets used for Pages Functions.
	R2Buckets map[string]ProjectDeploymentConfigsPreviewR2Bucket `json:"r2_buckets"`
	// Services used for Pages Functions.
	Services map[string]ProjectDeploymentConfigsPreviewService `json:"services"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings map[string]ProjectDeploymentConfigsPreviewVectorizeBinding `json:"vectorize_bindings"`
	// Hash of the Wrangler configuration used for the deployment.
	WranglerConfigHash string                              `json:"wrangler_config_hash"`
	JSON               projectDeploymentConfigsPreviewJSON `json:"-"`
}

// projectDeploymentConfigsPreviewJSON contains the JSON metadata for the struct
// [ProjectDeploymentConfigsPreview]
type projectDeploymentConfigsPreviewJSON struct {
	AlwaysUseLatestCompatibilityDate apijson.Field
	BuildImageMajorVersion           apijson.Field
	CompatibilityDate                apijson.Field
	CompatibilityFlags               apijson.Field
	EnvVars                          apijson.Field
	FailOpen                         apijson.Field
	UsageModel                       apijson.Field
	AIBindings                       apijson.Field
	AnalyticsEngineDatasets          apijson.Field
	Browsers                         apijson.Field
	D1Databases                      apijson.Field
	DurableObjectNamespaces          apijson.Field
	HyperdriveBindings               apijson.Field
	KVNamespaces                     apijson.Field
	Limits                           apijson.Field
	MTLSCertificates                 apijson.Field
	Placement                        apijson.Field
	QueueProducers                   apijson.Field
	R2Buckets                        apijson.Field
	Services                         apijson.Field
	VectorizeBindings                apijson.Field
	WranglerConfigHash               apijson.Field
	raw                              string
	ExtraFields                      map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreview) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewJSON) RawJSON() string {
	return r.raw
}

// A plaintext environment variable.
type ProjectDeploymentConfigsPreviewEnvVar struct {
	Type ProjectDeploymentConfigsPreviewEnvVarsType `json:"type,required"`
	// Environment variable value.
	Value string                                    `json:"value,required"`
	JSON  projectDeploymentConfigsPreviewEnvVarJSON `json:"-"`
	union ProjectDeploymentConfigsPreviewEnvVarsUnion
}

// projectDeploymentConfigsPreviewEnvVarJSON contains the JSON metadata for the
// struct [ProjectDeploymentConfigsPreviewEnvVar]
type projectDeploymentConfigsPreviewEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectDeploymentConfigsPreviewEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectDeploymentConfigsPreviewEnvVar) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectDeploymentConfigsPreviewEnvVar{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectDeploymentConfigsPreviewEnvVarsUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar],
// [ProjectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar].
func (r ProjectDeploymentConfigsPreviewEnvVar) AsUnion() ProjectDeploymentConfigsPreviewEnvVarsUnion {
	return r.union
}

// A plaintext environment variable.
//
// Union satisfied by [ProjectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar]
// or [ProjectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar].
type ProjectDeploymentConfigsPreviewEnvVarsUnion interface {
	implementsProjectDeploymentConfigsPreviewEnvVar()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectDeploymentConfigsPreviewEnvVarsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar{}),
			DiscriminatorValue: "secret_text",
		},
	)
}

// A plaintext environment variable.
type ProjectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar struct {
	Type ProjectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType `json:"type,required"`
	// Environment variable value.
	Value string                                                         `json:"value,required"`
	JSON  projectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarJSON `json:"-"`
}

// projectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarJSON contains the JSON
// metadata for the struct
// [ProjectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar]
type projectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar) implementsProjectDeploymentConfigsPreviewEnvVar() {
}

type ProjectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType string

const (
	ProjectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarTypePlainText ProjectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r ProjectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type ProjectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar struct {
	Type ProjectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType `json:"type,required"`
	// Secret value.
	Value string                                                          `json:"value,required"`
	JSON  projectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarJSON `json:"-"`
}

// projectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarJSON contains the
// JSON metadata for the struct
// [ProjectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar]
type projectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar) implementsProjectDeploymentConfigsPreviewEnvVar() {
}

type ProjectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType string

const (
	ProjectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarTypeSecretText ProjectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r ProjectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type ProjectDeploymentConfigsPreviewEnvVarsType string

const (
	ProjectDeploymentConfigsPreviewEnvVarsTypePlainText  ProjectDeploymentConfigsPreviewEnvVarsType = "plain_text"
	ProjectDeploymentConfigsPreviewEnvVarsTypeSecretText ProjectDeploymentConfigsPreviewEnvVarsType = "secret_text"
)

func (r ProjectDeploymentConfigsPreviewEnvVarsType) IsKnown() bool {
	switch r {
	case ProjectDeploymentConfigsPreviewEnvVarsTypePlainText, ProjectDeploymentConfigsPreviewEnvVarsTypeSecretText:
		return true
	}
	return false
}

// The usage model for Pages Functions.
type ProjectDeploymentConfigsPreviewUsageModel string

const (
	ProjectDeploymentConfigsPreviewUsageModelStandard ProjectDeploymentConfigsPreviewUsageModel = "standard"
	ProjectDeploymentConfigsPreviewUsageModelBundled  ProjectDeploymentConfigsPreviewUsageModel = "bundled"
	ProjectDeploymentConfigsPreviewUsageModelUnbound  ProjectDeploymentConfigsPreviewUsageModel = "unbound"
)

func (r ProjectDeploymentConfigsPreviewUsageModel) IsKnown() bool {
	switch r {
	case ProjectDeploymentConfigsPreviewUsageModelStandard, ProjectDeploymentConfigsPreviewUsageModelBundled, ProjectDeploymentConfigsPreviewUsageModelUnbound:
		return true
	}
	return false
}

// AI binding.
type ProjectDeploymentConfigsPreviewAIBinding struct {
	ProjectID string                                       `json:"project_id,required"`
	JSON      projectDeploymentConfigsPreviewAIBindingJSON `json:"-"`
}

// projectDeploymentConfigsPreviewAIBindingJSON contains the JSON metadata for the
// struct [ProjectDeploymentConfigsPreviewAIBinding]
type projectDeploymentConfigsPreviewAIBindingJSON struct {
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewAIBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewAIBindingJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine binding.
type ProjectDeploymentConfigsPreviewAnalyticsEngineDataset struct {
	// Name of the dataset.
	Dataset string                                                    `json:"dataset,required"`
	JSON    projectDeploymentConfigsPreviewAnalyticsEngineDatasetJSON `json:"-"`
}

// projectDeploymentConfigsPreviewAnalyticsEngineDatasetJSON contains the JSON
// metadata for the struct [ProjectDeploymentConfigsPreviewAnalyticsEngineDataset]
type projectDeploymentConfigsPreviewAnalyticsEngineDatasetJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewAnalyticsEngineDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewAnalyticsEngineDatasetJSON) RawJSON() string {
	return r.raw
}

// Browser binding.
type ProjectDeploymentConfigsPreviewBrowser struct {
	JSON projectDeploymentConfigsPreviewBrowserJSON `json:"-"`
}

// projectDeploymentConfigsPreviewBrowserJSON contains the JSON metadata for the
// struct [ProjectDeploymentConfigsPreviewBrowser]
type projectDeploymentConfigsPreviewBrowserJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewBrowser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewBrowserJSON) RawJSON() string {
	return r.raw
}

// D1 binding.
type ProjectDeploymentConfigsPreviewD1Database struct {
	// UUID of the D1 database.
	ID   string                                        `json:"id,required"`
	JSON projectDeploymentConfigsPreviewD1DatabaseJSON `json:"-"`
}

// projectDeploymentConfigsPreviewD1DatabaseJSON contains the JSON metadata for the
// struct [ProjectDeploymentConfigsPreviewD1Database]
type projectDeploymentConfigsPreviewD1DatabaseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewD1Database) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewD1DatabaseJSON) RawJSON() string {
	return r.raw
}

// Durable Object binding.
type ProjectDeploymentConfigsPreviewDurableObjectNamespace struct {
	// ID of the Durable Object namespace.
	NamespaceID string                                                    `json:"namespace_id,required"`
	JSON        projectDeploymentConfigsPreviewDurableObjectNamespaceJSON `json:"-"`
}

// projectDeploymentConfigsPreviewDurableObjectNamespaceJSON contains the JSON
// metadata for the struct [ProjectDeploymentConfigsPreviewDurableObjectNamespace]
type projectDeploymentConfigsPreviewDurableObjectNamespaceJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

// Hyperdrive binding.
type ProjectDeploymentConfigsPreviewHyperdriveBinding struct {
	ID   string                                               `json:"id,required"`
	JSON projectDeploymentConfigsPreviewHyperdriveBindingJSON `json:"-"`
}

// projectDeploymentConfigsPreviewHyperdriveBindingJSON contains the JSON metadata
// for the struct [ProjectDeploymentConfigsPreviewHyperdriveBinding]
type projectDeploymentConfigsPreviewHyperdriveBindingJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewHyperdriveBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewHyperdriveBindingJSON) RawJSON() string {
	return r.raw
}

// KV namespace binding.
type ProjectDeploymentConfigsPreviewKVNamespace struct {
	// ID of the KV namespace.
	NamespaceID string                                         `json:"namespace_id,required"`
	JSON        projectDeploymentConfigsPreviewKVNamespaceJSON `json:"-"`
}

// projectDeploymentConfigsPreviewKVNamespaceJSON contains the JSON metadata for
// the struct [ProjectDeploymentConfigsPreviewKVNamespace]
type projectDeploymentConfigsPreviewKVNamespaceJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewKVNamespaceJSON) RawJSON() string {
	return r.raw
}

// Limits for Pages Functions.
type ProjectDeploymentConfigsPreviewLimits struct {
	// CPU time limit in milliseconds.
	CPUMs int64                                     `json:"cpu_ms,required"`
	JSON  projectDeploymentConfigsPreviewLimitsJSON `json:"-"`
}

// projectDeploymentConfigsPreviewLimitsJSON contains the JSON metadata for the
// struct [ProjectDeploymentConfigsPreviewLimits]
type projectDeploymentConfigsPreviewLimitsJSON struct {
	CPUMs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewLimitsJSON) RawJSON() string {
	return r.raw
}

// mTLS binding.
type ProjectDeploymentConfigsPreviewMTLSCertificate struct {
	CertificateID string                                             `json:"certificate_id,required"`
	JSON          projectDeploymentConfigsPreviewMTLSCertificateJSON `json:"-"`
}

// projectDeploymentConfigsPreviewMTLSCertificateJSON contains the JSON metadata
// for the struct [ProjectDeploymentConfigsPreviewMTLSCertificate]
type projectDeploymentConfigsPreviewMTLSCertificateJSON struct {
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewMTLSCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewMTLSCertificateJSON) RawJSON() string {
	return r.raw
}

// Placement setting used for Pages Functions.
type ProjectDeploymentConfigsPreviewPlacement struct {
	// Placement mode.
	Mode string                                       `json:"mode,required"`
	JSON projectDeploymentConfigsPreviewPlacementJSON `json:"-"`
}

// projectDeploymentConfigsPreviewPlacementJSON contains the JSON metadata for the
// struct [ProjectDeploymentConfigsPreviewPlacement]
type projectDeploymentConfigsPreviewPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewPlacementJSON) RawJSON() string {
	return r.raw
}

// Queue Producer binding.
type ProjectDeploymentConfigsPreviewQueueProducer struct {
	// Name of the Queue.
	Name string                                           `json:"name,required"`
	JSON projectDeploymentConfigsPreviewQueueProducerJSON `json:"-"`
}

// projectDeploymentConfigsPreviewQueueProducerJSON contains the JSON metadata for
// the struct [ProjectDeploymentConfigsPreviewQueueProducer]
type projectDeploymentConfigsPreviewQueueProducerJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewQueueProducer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewQueueProducerJSON) RawJSON() string {
	return r.raw
}

// R2 binding.
type ProjectDeploymentConfigsPreviewR2Bucket struct {
	// Name of the R2 bucket.
	Name string `json:"name,required"`
	// Jurisdiction of the R2 bucket.
	Jurisdiction string                                      `json:"jurisdiction,nullable"`
	JSON         projectDeploymentConfigsPreviewR2BucketJSON `json:"-"`
}

// projectDeploymentConfigsPreviewR2BucketJSON contains the JSON metadata for the
// struct [ProjectDeploymentConfigsPreviewR2Bucket]
type projectDeploymentConfigsPreviewR2BucketJSON struct {
	Name         apijson.Field
	Jurisdiction apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewR2BucketJSON) RawJSON() string {
	return r.raw
}

// Service binding.
type ProjectDeploymentConfigsPreviewService struct {
	// The Service environment.
	Environment string `json:"environment,required"`
	// The Service name.
	Service string `json:"service,required"`
	// The entrypoint to bind to.
	Entrypoint string                                     `json:"entrypoint,nullable"`
	JSON       projectDeploymentConfigsPreviewServiceJSON `json:"-"`
}

// projectDeploymentConfigsPreviewServiceJSON contains the JSON metadata for the
// struct [ProjectDeploymentConfigsPreviewService]
type projectDeploymentConfigsPreviewServiceJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	Entrypoint  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewServiceJSON) RawJSON() string {
	return r.raw
}

// Vectorize binding.
type ProjectDeploymentConfigsPreviewVectorizeBinding struct {
	IndexName string                                              `json:"index_name,required"`
	JSON      projectDeploymentConfigsPreviewVectorizeBindingJSON `json:"-"`
}

// projectDeploymentConfigsPreviewVectorizeBindingJSON contains the JSON metadata
// for the struct [ProjectDeploymentConfigsPreviewVectorizeBinding]
type projectDeploymentConfigsPreviewVectorizeBindingJSON struct {
	IndexName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsPreviewVectorizeBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsPreviewVectorizeBindingJSON) RawJSON() string {
	return r.raw
}

// Configs for production deploys.
type ProjectDeploymentConfigsProduction struct {
	// Whether to always use the latest compatibility date for Pages Functions.
	AlwaysUseLatestCompatibilityDate bool `json:"always_use_latest_compatibility_date,required"`
	// The major version of the build image to use for Pages Functions.
	BuildImageMajorVersion int64 `json:"build_image_major_version,required"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate string `json:"compatibility_date,required"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags []string `json:"compatibility_flags,required"`
	// Environment variables used for builds and Pages Functions.
	EnvVars map[string]ProjectDeploymentConfigsProductionEnvVar `json:"env_vars,required,nullable"`
	// Whether to fail open when the deployment config cannot be applied.
	FailOpen bool `json:"fail_open,required"`
	// The usage model for Pages Functions.
	//
	// Deprecated: All new projects now use the Standard usage model.
	UsageModel ProjectDeploymentConfigsProductionUsageModel `json:"usage_model,required"`
	// Constellation bindings used for Pages Functions.
	AIBindings map[string]ProjectDeploymentConfigsProductionAIBinding `json:"ai_bindings"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets map[string]ProjectDeploymentConfigsProductionAnalyticsEngineDataset `json:"analytics_engine_datasets"`
	// Browser bindings used for Pages Functions.
	Browsers map[string]ProjectDeploymentConfigsProductionBrowser `json:"browsers"`
	// D1 databases used for Pages Functions.
	D1Databases map[string]ProjectDeploymentConfigsProductionD1Database `json:"d1_databases"`
	// Durable Object namespaces used for Pages Functions.
	DurableObjectNamespaces map[string]ProjectDeploymentConfigsProductionDurableObjectNamespace `json:"durable_object_namespaces"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings map[string]ProjectDeploymentConfigsProductionHyperdriveBinding `json:"hyperdrive_bindings"`
	// KV namespaces used for Pages Functions.
	KVNamespaces map[string]ProjectDeploymentConfigsProductionKVNamespace `json:"kv_namespaces"`
	// Limits for Pages Functions.
	Limits ProjectDeploymentConfigsProductionLimits `json:"limits"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates map[string]ProjectDeploymentConfigsProductionMTLSCertificate `json:"mtls_certificates"`
	// Placement setting used for Pages Functions.
	Placement ProjectDeploymentConfigsProductionPlacement `json:"placement"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers map[string]ProjectDeploymentConfigsProductionQueueProducer `json:"queue_producers"`
	// R2 buckets used for Pages Functions.
	R2Buckets map[string]ProjectDeploymentConfigsProductionR2Bucket `json:"r2_buckets"`
	// Services used for Pages Functions.
	Services map[string]ProjectDeploymentConfigsProductionService `json:"services"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings map[string]ProjectDeploymentConfigsProductionVectorizeBinding `json:"vectorize_bindings"`
	// Hash of the Wrangler configuration used for the deployment.
	WranglerConfigHash string                                 `json:"wrangler_config_hash"`
	JSON               projectDeploymentConfigsProductionJSON `json:"-"`
}

// projectDeploymentConfigsProductionJSON contains the JSON metadata for the struct
// [ProjectDeploymentConfigsProduction]
type projectDeploymentConfigsProductionJSON struct {
	AlwaysUseLatestCompatibilityDate apijson.Field
	BuildImageMajorVersion           apijson.Field
	CompatibilityDate                apijson.Field
	CompatibilityFlags               apijson.Field
	EnvVars                          apijson.Field
	FailOpen                         apijson.Field
	UsageModel                       apijson.Field
	AIBindings                       apijson.Field
	AnalyticsEngineDatasets          apijson.Field
	Browsers                         apijson.Field
	D1Databases                      apijson.Field
	DurableObjectNamespaces          apijson.Field
	HyperdriveBindings               apijson.Field
	KVNamespaces                     apijson.Field
	Limits                           apijson.Field
	MTLSCertificates                 apijson.Field
	Placement                        apijson.Field
	QueueProducers                   apijson.Field
	R2Buckets                        apijson.Field
	Services                         apijson.Field
	VectorizeBindings                apijson.Field
	WranglerConfigHash               apijson.Field
	raw                              string
	ExtraFields                      map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProduction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionJSON) RawJSON() string {
	return r.raw
}

// A plaintext environment variable.
type ProjectDeploymentConfigsProductionEnvVar struct {
	Type ProjectDeploymentConfigsProductionEnvVarsType `json:"type,required"`
	// Environment variable value.
	Value string                                       `json:"value,required"`
	JSON  projectDeploymentConfigsProductionEnvVarJSON `json:"-"`
	union ProjectDeploymentConfigsProductionEnvVarsUnion
}

// projectDeploymentConfigsProductionEnvVarJSON contains the JSON metadata for the
// struct [ProjectDeploymentConfigsProductionEnvVar]
type projectDeploymentConfigsProductionEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectDeploymentConfigsProductionEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectDeploymentConfigsProductionEnvVar) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectDeploymentConfigsProductionEnvVar{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectDeploymentConfigsProductionEnvVarsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar],
// [ProjectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar].
func (r ProjectDeploymentConfigsProductionEnvVar) AsUnion() ProjectDeploymentConfigsProductionEnvVarsUnion {
	return r.union
}

// A plaintext environment variable.
//
// Union satisfied by
// [ProjectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar] or
// [ProjectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar].
type ProjectDeploymentConfigsProductionEnvVarsUnion interface {
	implementsProjectDeploymentConfigsProductionEnvVar()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectDeploymentConfigsProductionEnvVarsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar{}),
			DiscriminatorValue: "secret_text",
		},
	)
}

// A plaintext environment variable.
type ProjectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar struct {
	Type ProjectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType `json:"type,required"`
	// Environment variable value.
	Value string                                                            `json:"value,required"`
	JSON  projectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarJSON `json:"-"`
}

// projectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarJSON contains the
// JSON metadata for the struct
// [ProjectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar]
type projectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar) implementsProjectDeploymentConfigsProductionEnvVar() {
}

type ProjectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType string

const (
	ProjectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarTypePlainText ProjectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r ProjectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type ProjectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar struct {
	Type ProjectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType `json:"type,required"`
	// Secret value.
	Value string                                                             `json:"value,required"`
	JSON  projectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarJSON `json:"-"`
}

// projectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarJSON contains the
// JSON metadata for the struct
// [ProjectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar]
type projectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar) implementsProjectDeploymentConfigsProductionEnvVar() {
}

type ProjectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType string

const (
	ProjectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarTypeSecretText ProjectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r ProjectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type ProjectDeploymentConfigsProductionEnvVarsType string

const (
	ProjectDeploymentConfigsProductionEnvVarsTypePlainText  ProjectDeploymentConfigsProductionEnvVarsType = "plain_text"
	ProjectDeploymentConfigsProductionEnvVarsTypeSecretText ProjectDeploymentConfigsProductionEnvVarsType = "secret_text"
)

func (r ProjectDeploymentConfigsProductionEnvVarsType) IsKnown() bool {
	switch r {
	case ProjectDeploymentConfigsProductionEnvVarsTypePlainText, ProjectDeploymentConfigsProductionEnvVarsTypeSecretText:
		return true
	}
	return false
}

// The usage model for Pages Functions.
type ProjectDeploymentConfigsProductionUsageModel string

const (
	ProjectDeploymentConfigsProductionUsageModelStandard ProjectDeploymentConfigsProductionUsageModel = "standard"
	ProjectDeploymentConfigsProductionUsageModelBundled  ProjectDeploymentConfigsProductionUsageModel = "bundled"
	ProjectDeploymentConfigsProductionUsageModelUnbound  ProjectDeploymentConfigsProductionUsageModel = "unbound"
)

func (r ProjectDeploymentConfigsProductionUsageModel) IsKnown() bool {
	switch r {
	case ProjectDeploymentConfigsProductionUsageModelStandard, ProjectDeploymentConfigsProductionUsageModelBundled, ProjectDeploymentConfigsProductionUsageModelUnbound:
		return true
	}
	return false
}

// AI binding.
type ProjectDeploymentConfigsProductionAIBinding struct {
	ProjectID string                                          `json:"project_id,required"`
	JSON      projectDeploymentConfigsProductionAIBindingJSON `json:"-"`
}

// projectDeploymentConfigsProductionAIBindingJSON contains the JSON metadata for
// the struct [ProjectDeploymentConfigsProductionAIBinding]
type projectDeploymentConfigsProductionAIBindingJSON struct {
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionAIBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionAIBindingJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine binding.
type ProjectDeploymentConfigsProductionAnalyticsEngineDataset struct {
	// Name of the dataset.
	Dataset string                                                       `json:"dataset,required"`
	JSON    projectDeploymentConfigsProductionAnalyticsEngineDatasetJSON `json:"-"`
}

// projectDeploymentConfigsProductionAnalyticsEngineDatasetJSON contains the JSON
// metadata for the struct
// [ProjectDeploymentConfigsProductionAnalyticsEngineDataset]
type projectDeploymentConfigsProductionAnalyticsEngineDatasetJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionAnalyticsEngineDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionAnalyticsEngineDatasetJSON) RawJSON() string {
	return r.raw
}

// Browser binding.
type ProjectDeploymentConfigsProductionBrowser struct {
	JSON projectDeploymentConfigsProductionBrowserJSON `json:"-"`
}

// projectDeploymentConfigsProductionBrowserJSON contains the JSON metadata for the
// struct [ProjectDeploymentConfigsProductionBrowser]
type projectDeploymentConfigsProductionBrowserJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionBrowser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionBrowserJSON) RawJSON() string {
	return r.raw
}

// D1 binding.
type ProjectDeploymentConfigsProductionD1Database struct {
	// UUID of the D1 database.
	ID   string                                           `json:"id,required"`
	JSON projectDeploymentConfigsProductionD1DatabaseJSON `json:"-"`
}

// projectDeploymentConfigsProductionD1DatabaseJSON contains the JSON metadata for
// the struct [ProjectDeploymentConfigsProductionD1Database]
type projectDeploymentConfigsProductionD1DatabaseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionD1Database) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionD1DatabaseJSON) RawJSON() string {
	return r.raw
}

// Durable Object binding.
type ProjectDeploymentConfigsProductionDurableObjectNamespace struct {
	// ID of the Durable Object namespace.
	NamespaceID string                                                       `json:"namespace_id,required"`
	JSON        projectDeploymentConfigsProductionDurableObjectNamespaceJSON `json:"-"`
}

// projectDeploymentConfigsProductionDurableObjectNamespaceJSON contains the JSON
// metadata for the struct
// [ProjectDeploymentConfigsProductionDurableObjectNamespace]
type projectDeploymentConfigsProductionDurableObjectNamespaceJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

// Hyperdrive binding.
type ProjectDeploymentConfigsProductionHyperdriveBinding struct {
	ID   string                                                  `json:"id,required"`
	JSON projectDeploymentConfigsProductionHyperdriveBindingJSON `json:"-"`
}

// projectDeploymentConfigsProductionHyperdriveBindingJSON contains the JSON
// metadata for the struct [ProjectDeploymentConfigsProductionHyperdriveBinding]
type projectDeploymentConfigsProductionHyperdriveBindingJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionHyperdriveBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionHyperdriveBindingJSON) RawJSON() string {
	return r.raw
}

// KV namespace binding.
type ProjectDeploymentConfigsProductionKVNamespace struct {
	// ID of the KV namespace.
	NamespaceID string                                            `json:"namespace_id,required"`
	JSON        projectDeploymentConfigsProductionKVNamespaceJSON `json:"-"`
}

// projectDeploymentConfigsProductionKVNamespaceJSON contains the JSON metadata for
// the struct [ProjectDeploymentConfigsProductionKVNamespace]
type projectDeploymentConfigsProductionKVNamespaceJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionKVNamespaceJSON) RawJSON() string {
	return r.raw
}

// Limits for Pages Functions.
type ProjectDeploymentConfigsProductionLimits struct {
	// CPU time limit in milliseconds.
	CPUMs int64                                        `json:"cpu_ms,required"`
	JSON  projectDeploymentConfigsProductionLimitsJSON `json:"-"`
}

// projectDeploymentConfigsProductionLimitsJSON contains the JSON metadata for the
// struct [ProjectDeploymentConfigsProductionLimits]
type projectDeploymentConfigsProductionLimitsJSON struct {
	CPUMs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionLimitsJSON) RawJSON() string {
	return r.raw
}

// mTLS binding.
type ProjectDeploymentConfigsProductionMTLSCertificate struct {
	CertificateID string                                                `json:"certificate_id,required"`
	JSON          projectDeploymentConfigsProductionMTLSCertificateJSON `json:"-"`
}

// projectDeploymentConfigsProductionMTLSCertificateJSON contains the JSON metadata
// for the struct [ProjectDeploymentConfigsProductionMTLSCertificate]
type projectDeploymentConfigsProductionMTLSCertificateJSON struct {
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionMTLSCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionMTLSCertificateJSON) RawJSON() string {
	return r.raw
}

// Placement setting used for Pages Functions.
type ProjectDeploymentConfigsProductionPlacement struct {
	// Placement mode.
	Mode string                                          `json:"mode,required"`
	JSON projectDeploymentConfigsProductionPlacementJSON `json:"-"`
}

// projectDeploymentConfigsProductionPlacementJSON contains the JSON metadata for
// the struct [ProjectDeploymentConfigsProductionPlacement]
type projectDeploymentConfigsProductionPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionPlacementJSON) RawJSON() string {
	return r.raw
}

// Queue Producer binding.
type ProjectDeploymentConfigsProductionQueueProducer struct {
	// Name of the Queue.
	Name string                                              `json:"name,required"`
	JSON projectDeploymentConfigsProductionQueueProducerJSON `json:"-"`
}

// projectDeploymentConfigsProductionQueueProducerJSON contains the JSON metadata
// for the struct [ProjectDeploymentConfigsProductionQueueProducer]
type projectDeploymentConfigsProductionQueueProducerJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionQueueProducer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionQueueProducerJSON) RawJSON() string {
	return r.raw
}

// R2 binding.
type ProjectDeploymentConfigsProductionR2Bucket struct {
	// Name of the R2 bucket.
	Name string `json:"name,required"`
	// Jurisdiction of the R2 bucket.
	Jurisdiction string                                         `json:"jurisdiction,nullable"`
	JSON         projectDeploymentConfigsProductionR2BucketJSON `json:"-"`
}

// projectDeploymentConfigsProductionR2BucketJSON contains the JSON metadata for
// the struct [ProjectDeploymentConfigsProductionR2Bucket]
type projectDeploymentConfigsProductionR2BucketJSON struct {
	Name         apijson.Field
	Jurisdiction apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionR2BucketJSON) RawJSON() string {
	return r.raw
}

// Service binding.
type ProjectDeploymentConfigsProductionService struct {
	// The Service environment.
	Environment string `json:"environment,required"`
	// The Service name.
	Service string `json:"service,required"`
	// The entrypoint to bind to.
	Entrypoint string                                        `json:"entrypoint,nullable"`
	JSON       projectDeploymentConfigsProductionServiceJSON `json:"-"`
}

// projectDeploymentConfigsProductionServiceJSON contains the JSON metadata for the
// struct [ProjectDeploymentConfigsProductionService]
type projectDeploymentConfigsProductionServiceJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	Entrypoint  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionServiceJSON) RawJSON() string {
	return r.raw
}

// Vectorize binding.
type ProjectDeploymentConfigsProductionVectorizeBinding struct {
	IndexName string                                                 `json:"index_name,required"`
	JSON      projectDeploymentConfigsProductionVectorizeBindingJSON `json:"-"`
}

// projectDeploymentConfigsProductionVectorizeBindingJSON contains the JSON
// metadata for the struct [ProjectDeploymentConfigsProductionVectorizeBinding]
type projectDeploymentConfigsProductionVectorizeBindingJSON struct {
	IndexName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeploymentConfigsProductionVectorizeBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeploymentConfigsProductionVectorizeBindingJSON) RawJSON() string {
	return r.raw
}

// Configs for the project build process.
type ProjectBuildConfig struct {
	// The classifying tag for analytics.
	WebAnalyticsTag string `json:"web_analytics_tag,required,nullable"`
	// The auth token for analytics.
	WebAnalyticsToken string `json:"web_analytics_token,required,nullable"`
	// Enable build caching for the project.
	BuildCaching bool `json:"build_caching,nullable"`
	// Command used to build project.
	BuildCommand string `json:"build_command,nullable"`
	// Assets output directory of the build.
	DestinationDir string `json:"destination_dir,nullable"`
	// Directory to run the command.
	RootDir string                 `json:"root_dir,nullable"`
	JSON    projectBuildConfigJSON `json:"-"`
}

// projectBuildConfigJSON contains the JSON metadata for the struct
// [ProjectBuildConfig]
type projectBuildConfigJSON struct {
	WebAnalyticsTag   apijson.Field
	WebAnalyticsToken apijson.Field
	BuildCaching      apijson.Field
	BuildCommand      apijson.Field
	DestinationDir    apijson.Field
	RootDir           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectBuildConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectBuildConfigJSON) RawJSON() string {
	return r.raw
}

// Configs for the project source control.
type ProjectSource struct {
	Config ProjectSourceConfig `json:"config,required"`
	// The source control management provider.
	Type ProjectSourceType `json:"type,required"`
	JSON projectSourceJSON `json:"-"`
}

// projectSourceJSON contains the JSON metadata for the struct [ProjectSource]
type projectSourceJSON struct {
	Config      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectSourceConfig struct {
	// Whether to enable automatic deployments when pushing to the source repository.
	// When disabled, no deployments (production or preview) will be triggered
	// automatically.
	//
	// Deprecated: Use `production_deployments_enabled` and
	// `preview_deployment_setting` for more granular control.
	DeploymentsEnabled bool `json:"deployments_enabled,required"`
	// The owner of the repository.
	Owner string `json:"owner,required"`
	// The owner ID of the repository.
	OwnerID string `json:"owner_id,required"`
	// A list of paths that should be excluded from triggering a preview deployment.
	// Wildcard syntax (`*`) is supported.
	PathExcludes []string `json:"path_excludes,required"`
	// A list of paths that should be watched to trigger a preview deployment. Wildcard
	// syntax (`*`) is supported.
	PathIncludes []string `json:"path_includes,required"`
	// Whether to enable PR comments.
	PrCommentsEnabled bool `json:"pr_comments_enabled,required"`
	// A list of branches that should not trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchExcludes []string `json:"preview_branch_excludes,required"`
	// A list of branches that should trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchIncludes []string `json:"preview_branch_includes,required"`
	// Controls whether commits to preview branches trigger a preview deployment.
	PreviewDeploymentSetting ProjectSourceConfigPreviewDeploymentSetting `json:"preview_deployment_setting,required"`
	// The production branch of the repository.
	ProductionBranch string `json:"production_branch,required"`
	// Whether to trigger a production deployment on commits to the production branch.
	ProductionDeploymentsEnabled bool `json:"production_deployments_enabled,required"`
	// The ID of the repository.
	RepoID string `json:"repo_id,required"`
	// The name of the repository.
	RepoName string                  `json:"repo_name,required"`
	JSON     projectSourceConfigJSON `json:"-"`
}

// projectSourceConfigJSON contains the JSON metadata for the struct
// [ProjectSourceConfig]
type projectSourceConfigJSON struct {
	DeploymentsEnabled           apijson.Field
	Owner                        apijson.Field
	OwnerID                      apijson.Field
	PathExcludes                 apijson.Field
	PathIncludes                 apijson.Field
	PrCommentsEnabled            apijson.Field
	PreviewBranchExcludes        apijson.Field
	PreviewBranchIncludes        apijson.Field
	PreviewDeploymentSetting     apijson.Field
	ProductionBranch             apijson.Field
	ProductionDeploymentsEnabled apijson.Field
	RepoID                       apijson.Field
	RepoName                     apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ProjectSourceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectSourceConfigJSON) RawJSON() string {
	return r.raw
}

// Controls whether commits to preview branches trigger a preview deployment.
type ProjectSourceConfigPreviewDeploymentSetting string

const (
	ProjectSourceConfigPreviewDeploymentSettingAll    ProjectSourceConfigPreviewDeploymentSetting = "all"
	ProjectSourceConfigPreviewDeploymentSettingNone   ProjectSourceConfigPreviewDeploymentSetting = "none"
	ProjectSourceConfigPreviewDeploymentSettingCustom ProjectSourceConfigPreviewDeploymentSetting = "custom"
)

func (r ProjectSourceConfigPreviewDeploymentSetting) IsKnown() bool {
	switch r {
	case ProjectSourceConfigPreviewDeploymentSettingAll, ProjectSourceConfigPreviewDeploymentSettingNone, ProjectSourceConfigPreviewDeploymentSettingCustom:
		return true
	}
	return false
}

// The source control management provider.
type ProjectSourceType string

const (
	ProjectSourceTypeGitHub ProjectSourceType = "github"
	ProjectSourceTypeGitlab ProjectSourceType = "gitlab"
)

func (r ProjectSourceType) IsKnown() bool {
	switch r {
	case ProjectSourceTypeGitHub, ProjectSourceTypeGitlab:
		return true
	}
	return false
}

// The status of the deployment.
type Stage struct {
	// When the stage ended.
	EndedOn time.Time `json:"ended_on,required,nullable" format:"date-time"`
	// The current build stage.
	Name StageName `json:"name,required"`
	// When the stage started.
	StartedOn time.Time `json:"started_on,required,nullable" format:"date-time"`
	// State of the current stage.
	Status StageStatus `json:"status,required"`
	JSON   stageJSON   `json:"-"`
}

// stageJSON contains the JSON metadata for the struct [Stage]
type stageJSON struct {
	EndedOn     apijson.Field
	Name        apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Stage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stageJSON) RawJSON() string {
	return r.raw
}

// The current build stage.
type StageName string

const (
	StageNameQueued     StageName = "queued"
	StageNameInitialize StageName = "initialize"
	StageNameCloneRepo  StageName = "clone_repo"
	StageNameBuild      StageName = "build"
	StageNameDeploy     StageName = "deploy"
)

func (r StageName) IsKnown() bool {
	switch r {
	case StageNameQueued, StageNameInitialize, StageNameCloneRepo, StageNameBuild, StageNameDeploy:
		return true
	}
	return false
}

// State of the current stage.
type StageStatus string

const (
	StageStatusSuccess  StageStatus = "success"
	StageStatusIdle     StageStatus = "idle"
	StageStatusActive   StageStatus = "active"
	StageStatusFailure  StageStatus = "failure"
	StageStatusCanceled StageStatus = "canceled"
)

func (r StageStatus) IsKnown() bool {
	switch r {
	case StageStatusSuccess, StageStatusIdle, StageStatusActive, StageStatusFailure, StageStatusCanceled:
		return true
	}
	return false
}

type ProjectDeleteResponse = interface{}

type ProjectPurgeBuildCacheResponse = interface{}

type ProjectNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Name of the project.
	Name param.Field[string] `json:"name,required"`
	// Production branch of the project. Used to identify production deployments.
	ProductionBranch param.Field[string] `json:"production_branch,required"`
	// Configs for the project build process.
	BuildConfig param.Field[ProjectNewParamsBuildConfig] `json:"build_config"`
	// Configs for deployments in a project.
	DeploymentConfigs param.Field[ProjectNewParamsDeploymentConfigs] `json:"deployment_configs"`
	// Configs for the project source control.
	Source param.Field[ProjectNewParamsSource] `json:"source"`
}

func (r ProjectNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for the project build process.
type ProjectNewParamsBuildConfig struct {
	// Enable build caching for the project.
	BuildCaching param.Field[bool] `json:"build_caching"`
	// Command used to build project.
	BuildCommand param.Field[string] `json:"build_command"`
	// Output directory of the build.
	DestinationDir param.Field[string] `json:"destination_dir"`
	// Directory to run the command.
	RootDir param.Field[string] `json:"root_dir"`
	// The classifying tag for analytics.
	WebAnalyticsTag param.Field[string] `json:"web_analytics_tag"`
	// The auth token for analytics.
	WebAnalyticsToken param.Field[string] `json:"web_analytics_token"`
}

func (r ProjectNewParamsBuildConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for deployments in a project.
type ProjectNewParamsDeploymentConfigs struct {
	// Configs for preview deploys.
	Preview param.Field[ProjectNewParamsDeploymentConfigsPreview] `json:"preview"`
	// Configs for production deploys.
	Production param.Field[ProjectNewParamsDeploymentConfigsProduction] `json:"production"`
}

func (r ProjectNewParamsDeploymentConfigs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for preview deploys.
type ProjectNewParamsDeploymentConfigsPreview struct {
	// Constellation bindings used for Pages Functions.
	AIBindings param.Field[map[string]ProjectNewParamsDeploymentConfigsPreviewAIBindings] `json:"ai_bindings"`
	// Whether to always use the latest compatibility date for Pages Functions.
	AlwaysUseLatestCompatibilityDate param.Field[bool] `json:"always_use_latest_compatibility_date"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets param.Field[map[string]ProjectNewParamsDeploymentConfigsPreviewAnalyticsEngineDatasets] `json:"analytics_engine_datasets"`
	// Browser bindings used for Pages Functions.
	Browsers param.Field[map[string]ProjectNewParamsDeploymentConfigsPreviewBrowsers] `json:"browsers"`
	// The major version of the build image to use for Pages Functions.
	BuildImageMajorVersion param.Field[int64] `json:"build_image_major_version"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases param.Field[map[string]ProjectNewParamsDeploymentConfigsPreviewD1Databases] `json:"d1_databases"`
	// Durable Object namespaces used for Pages Functions.
	DurableObjectNamespaces param.Field[map[string]ProjectNewParamsDeploymentConfigsPreviewDurableObjectNamespaces] `json:"durable_object_namespaces"`
	// Environment variables used for builds and Pages Functions.
	EnvVars param.Field[map[string]ProjectNewParamsDeploymentConfigsPreviewEnvVarsUnion] `json:"env_vars"`
	// Whether to fail open when the deployment config cannot be applied.
	FailOpen param.Field[bool] `json:"fail_open"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings param.Field[map[string]ProjectNewParamsDeploymentConfigsPreviewHyperdriveBindings] `json:"hyperdrive_bindings"`
	// KV namespaces used for Pages Functions.
	KVNamespaces param.Field[map[string]ProjectNewParamsDeploymentConfigsPreviewKVNamespaces] `json:"kv_namespaces"`
	// Limits for Pages Functions.
	Limits param.Field[ProjectNewParamsDeploymentConfigsPreviewLimits] `json:"limits"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates param.Field[map[string]ProjectNewParamsDeploymentConfigsPreviewMTLSCertificates] `json:"mtls_certificates"`
	// Placement setting used for Pages Functions.
	Placement param.Field[ProjectNewParamsDeploymentConfigsPreviewPlacement] `json:"placement"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers param.Field[map[string]ProjectNewParamsDeploymentConfigsPreviewQueueProducers] `json:"queue_producers"`
	// R2 buckets used for Pages Functions.
	R2Buckets param.Field[map[string]ProjectNewParamsDeploymentConfigsPreviewR2Buckets] `json:"r2_buckets"`
	// Services used for Pages Functions.
	Services param.Field[map[string]ProjectNewParamsDeploymentConfigsPreviewServices] `json:"services"`
	// The usage model for Pages Functions.
	//
	// Deprecated: All new projects now use the Standard usage model.
	UsageModel param.Field[ProjectNewParamsDeploymentConfigsPreviewUsageModel] `json:"usage_model"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings param.Field[map[string]ProjectNewParamsDeploymentConfigsPreviewVectorizeBindings] `json:"vectorize_bindings"`
	// Hash of the Wrangler configuration used for the deployment.
	WranglerConfigHash param.Field[string] `json:"wrangler_config_hash"`
}

func (r ProjectNewParamsDeploymentConfigsPreview) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AI binding.
type ProjectNewParamsDeploymentConfigsPreviewAIBindings struct {
	ProjectID param.Field[string] `json:"project_id,required"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewAIBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Analytics Engine binding.
type ProjectNewParamsDeploymentConfigsPreviewAnalyticsEngineDatasets struct {
	// Name of the dataset.
	Dataset param.Field[string] `json:"dataset,required"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewAnalyticsEngineDatasets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Browser binding.
type ProjectNewParamsDeploymentConfigsPreviewBrowsers struct {
}

func (r ProjectNewParamsDeploymentConfigsPreviewBrowsers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// D1 binding.
type ProjectNewParamsDeploymentConfigsPreviewD1Databases struct {
	// UUID of the D1 database.
	ID param.Field[string] `json:"id,required"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewD1Databases) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Durable Object binding.
type ProjectNewParamsDeploymentConfigsPreviewDurableObjectNamespaces struct {
	// ID of the Durable Object namespace.
	NamespaceID param.Field[string] `json:"namespace_id,required"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewDurableObjectNamespaces) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A plaintext environment variable.
type ProjectNewParamsDeploymentConfigsPreviewEnvVars struct {
	Type param.Field[ProjectNewParamsDeploymentConfigsPreviewEnvVarsType] `json:"type,required"`
	// Environment variable value.
	Value param.Field[string] `json:"value,required"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewEnvVars) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectNewParamsDeploymentConfigsPreviewEnvVars) implementsProjectNewParamsDeploymentConfigsPreviewEnvVarsUnion() {
}

// A plaintext environment variable.
//
// Satisfied by
// [pages.ProjectNewParamsDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar],
// [pages.ProjectNewParamsDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar],
// [ProjectNewParamsDeploymentConfigsPreviewEnvVars].
type ProjectNewParamsDeploymentConfigsPreviewEnvVarsUnion interface {
	implementsProjectNewParamsDeploymentConfigsPreviewEnvVarsUnion()
}

// A plaintext environment variable.
type ProjectNewParamsDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar struct {
	Type param.Field[ProjectNewParamsDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType] `json:"type,required"`
	// Environment variable value.
	Value param.Field[string] `json:"value,required"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectNewParamsDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar) implementsProjectNewParamsDeploymentConfigsPreviewEnvVarsUnion() {
}

type ProjectNewParamsDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType string

const (
	ProjectNewParamsDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarTypePlainText ProjectNewParamsDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r ProjectNewParamsDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectNewParamsDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type ProjectNewParamsDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar struct {
	Type param.Field[ProjectNewParamsDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType] `json:"type,required"`
	// Secret value.
	Value param.Field[string] `json:"value,required"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectNewParamsDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar) implementsProjectNewParamsDeploymentConfigsPreviewEnvVarsUnion() {
}

type ProjectNewParamsDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType string

const (
	ProjectNewParamsDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarTypeSecretText ProjectNewParamsDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r ProjectNewParamsDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectNewParamsDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type ProjectNewParamsDeploymentConfigsPreviewEnvVarsType string

const (
	ProjectNewParamsDeploymentConfigsPreviewEnvVarsTypePlainText  ProjectNewParamsDeploymentConfigsPreviewEnvVarsType = "plain_text"
	ProjectNewParamsDeploymentConfigsPreviewEnvVarsTypeSecretText ProjectNewParamsDeploymentConfigsPreviewEnvVarsType = "secret_text"
)

func (r ProjectNewParamsDeploymentConfigsPreviewEnvVarsType) IsKnown() bool {
	switch r {
	case ProjectNewParamsDeploymentConfigsPreviewEnvVarsTypePlainText, ProjectNewParamsDeploymentConfigsPreviewEnvVarsTypeSecretText:
		return true
	}
	return false
}

// Hyperdrive binding.
type ProjectNewParamsDeploymentConfigsPreviewHyperdriveBindings struct {
	ID param.Field[string] `json:"id,required"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewHyperdriveBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// KV namespace binding.
type ProjectNewParamsDeploymentConfigsPreviewKVNamespaces struct {
	// ID of the KV namespace.
	NamespaceID param.Field[string] `json:"namespace_id,required"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewKVNamespaces) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Limits for Pages Functions.
type ProjectNewParamsDeploymentConfigsPreviewLimits struct {
	// CPU time limit in milliseconds.
	CPUMs param.Field[int64] `json:"cpu_ms,required"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewLimits) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// mTLS binding.
type ProjectNewParamsDeploymentConfigsPreviewMTLSCertificates struct {
	CertificateID param.Field[string] `json:"certificate_id,required"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewMTLSCertificates) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Placement setting used for Pages Functions.
type ProjectNewParamsDeploymentConfigsPreviewPlacement struct {
	// Placement mode.
	Mode param.Field[string] `json:"mode,required"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Queue Producer binding.
type ProjectNewParamsDeploymentConfigsPreviewQueueProducers struct {
	// Name of the Queue.
	Name param.Field[string] `json:"name,required"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewQueueProducers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// R2 binding.
type ProjectNewParamsDeploymentConfigsPreviewR2Buckets struct {
	// Name of the R2 bucket.
	Name param.Field[string] `json:"name,required"`
	// Jurisdiction of the R2 bucket.
	Jurisdiction param.Field[string] `json:"jurisdiction"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewR2Buckets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Service binding.
type ProjectNewParamsDeploymentConfigsPreviewServices struct {
	// The Service name.
	Service param.Field[string] `json:"service,required"`
	// The entrypoint to bind to.
	Entrypoint param.Field[string] `json:"entrypoint"`
	// The Service environment.
	Environment param.Field[string] `json:"environment"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewServices) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The usage model for Pages Functions.
type ProjectNewParamsDeploymentConfigsPreviewUsageModel string

const (
	ProjectNewParamsDeploymentConfigsPreviewUsageModelStandard ProjectNewParamsDeploymentConfigsPreviewUsageModel = "standard"
	ProjectNewParamsDeploymentConfigsPreviewUsageModelBundled  ProjectNewParamsDeploymentConfigsPreviewUsageModel = "bundled"
	ProjectNewParamsDeploymentConfigsPreviewUsageModelUnbound  ProjectNewParamsDeploymentConfigsPreviewUsageModel = "unbound"
)

func (r ProjectNewParamsDeploymentConfigsPreviewUsageModel) IsKnown() bool {
	switch r {
	case ProjectNewParamsDeploymentConfigsPreviewUsageModelStandard, ProjectNewParamsDeploymentConfigsPreviewUsageModelBundled, ProjectNewParamsDeploymentConfigsPreviewUsageModelUnbound:
		return true
	}
	return false
}

// Vectorize binding.
type ProjectNewParamsDeploymentConfigsPreviewVectorizeBindings struct {
	IndexName param.Field[string] `json:"index_name,required"`
}

func (r ProjectNewParamsDeploymentConfigsPreviewVectorizeBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for production deploys.
type ProjectNewParamsDeploymentConfigsProduction struct {
	// Constellation bindings used for Pages Functions.
	AIBindings param.Field[map[string]ProjectNewParamsDeploymentConfigsProductionAIBindings] `json:"ai_bindings"`
	// Whether to always use the latest compatibility date for Pages Functions.
	AlwaysUseLatestCompatibilityDate param.Field[bool] `json:"always_use_latest_compatibility_date"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets param.Field[map[string]ProjectNewParamsDeploymentConfigsProductionAnalyticsEngineDatasets] `json:"analytics_engine_datasets"`
	// Browser bindings used for Pages Functions.
	Browsers param.Field[map[string]ProjectNewParamsDeploymentConfigsProductionBrowsers] `json:"browsers"`
	// The major version of the build image to use for Pages Functions.
	BuildImageMajorVersion param.Field[int64] `json:"build_image_major_version"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases param.Field[map[string]ProjectNewParamsDeploymentConfigsProductionD1Databases] `json:"d1_databases"`
	// Durable Object namespaces used for Pages Functions.
	DurableObjectNamespaces param.Field[map[string]ProjectNewParamsDeploymentConfigsProductionDurableObjectNamespaces] `json:"durable_object_namespaces"`
	// Environment variables used for builds and Pages Functions.
	EnvVars param.Field[map[string]ProjectNewParamsDeploymentConfigsProductionEnvVarsUnion] `json:"env_vars"`
	// Whether to fail open when the deployment config cannot be applied.
	FailOpen param.Field[bool] `json:"fail_open"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings param.Field[map[string]ProjectNewParamsDeploymentConfigsProductionHyperdriveBindings] `json:"hyperdrive_bindings"`
	// KV namespaces used for Pages Functions.
	KVNamespaces param.Field[map[string]ProjectNewParamsDeploymentConfigsProductionKVNamespaces] `json:"kv_namespaces"`
	// Limits for Pages Functions.
	Limits param.Field[ProjectNewParamsDeploymentConfigsProductionLimits] `json:"limits"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates param.Field[map[string]ProjectNewParamsDeploymentConfigsProductionMTLSCertificates] `json:"mtls_certificates"`
	// Placement setting used for Pages Functions.
	Placement param.Field[ProjectNewParamsDeploymentConfigsProductionPlacement] `json:"placement"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers param.Field[map[string]ProjectNewParamsDeploymentConfigsProductionQueueProducers] `json:"queue_producers"`
	// R2 buckets used for Pages Functions.
	R2Buckets param.Field[map[string]ProjectNewParamsDeploymentConfigsProductionR2Buckets] `json:"r2_buckets"`
	// Services used for Pages Functions.
	Services param.Field[map[string]ProjectNewParamsDeploymentConfigsProductionServices] `json:"services"`
	// The usage model for Pages Functions.
	//
	// Deprecated: All new projects now use the Standard usage model.
	UsageModel param.Field[ProjectNewParamsDeploymentConfigsProductionUsageModel] `json:"usage_model"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings param.Field[map[string]ProjectNewParamsDeploymentConfigsProductionVectorizeBindings] `json:"vectorize_bindings"`
	// Hash of the Wrangler configuration used for the deployment.
	WranglerConfigHash param.Field[string] `json:"wrangler_config_hash"`
}

func (r ProjectNewParamsDeploymentConfigsProduction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AI binding.
type ProjectNewParamsDeploymentConfigsProductionAIBindings struct {
	ProjectID param.Field[string] `json:"project_id,required"`
}

func (r ProjectNewParamsDeploymentConfigsProductionAIBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Analytics Engine binding.
type ProjectNewParamsDeploymentConfigsProductionAnalyticsEngineDatasets struct {
	// Name of the dataset.
	Dataset param.Field[string] `json:"dataset,required"`
}

func (r ProjectNewParamsDeploymentConfigsProductionAnalyticsEngineDatasets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Browser binding.
type ProjectNewParamsDeploymentConfigsProductionBrowsers struct {
}

func (r ProjectNewParamsDeploymentConfigsProductionBrowsers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// D1 binding.
type ProjectNewParamsDeploymentConfigsProductionD1Databases struct {
	// UUID of the D1 database.
	ID param.Field[string] `json:"id,required"`
}

func (r ProjectNewParamsDeploymentConfigsProductionD1Databases) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Durable Object binding.
type ProjectNewParamsDeploymentConfigsProductionDurableObjectNamespaces struct {
	// ID of the Durable Object namespace.
	NamespaceID param.Field[string] `json:"namespace_id,required"`
}

func (r ProjectNewParamsDeploymentConfigsProductionDurableObjectNamespaces) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A plaintext environment variable.
type ProjectNewParamsDeploymentConfigsProductionEnvVars struct {
	Type param.Field[ProjectNewParamsDeploymentConfigsProductionEnvVarsType] `json:"type,required"`
	// Environment variable value.
	Value param.Field[string] `json:"value,required"`
}

func (r ProjectNewParamsDeploymentConfigsProductionEnvVars) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectNewParamsDeploymentConfigsProductionEnvVars) implementsProjectNewParamsDeploymentConfigsProductionEnvVarsUnion() {
}

// A plaintext environment variable.
//
// Satisfied by
// [pages.ProjectNewParamsDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar],
// [pages.ProjectNewParamsDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar],
// [ProjectNewParamsDeploymentConfigsProductionEnvVars].
type ProjectNewParamsDeploymentConfigsProductionEnvVarsUnion interface {
	implementsProjectNewParamsDeploymentConfigsProductionEnvVarsUnion()
}

// A plaintext environment variable.
type ProjectNewParamsDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar struct {
	Type param.Field[ProjectNewParamsDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType] `json:"type,required"`
	// Environment variable value.
	Value param.Field[string] `json:"value,required"`
}

func (r ProjectNewParamsDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectNewParamsDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar) implementsProjectNewParamsDeploymentConfigsProductionEnvVarsUnion() {
}

type ProjectNewParamsDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType string

const (
	ProjectNewParamsDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarTypePlainText ProjectNewParamsDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r ProjectNewParamsDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectNewParamsDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type ProjectNewParamsDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar struct {
	Type param.Field[ProjectNewParamsDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType] `json:"type,required"`
	// Secret value.
	Value param.Field[string] `json:"value,required"`
}

func (r ProjectNewParamsDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectNewParamsDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar) implementsProjectNewParamsDeploymentConfigsProductionEnvVarsUnion() {
}

type ProjectNewParamsDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType string

const (
	ProjectNewParamsDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarTypeSecretText ProjectNewParamsDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r ProjectNewParamsDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectNewParamsDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type ProjectNewParamsDeploymentConfigsProductionEnvVarsType string

const (
	ProjectNewParamsDeploymentConfigsProductionEnvVarsTypePlainText  ProjectNewParamsDeploymentConfigsProductionEnvVarsType = "plain_text"
	ProjectNewParamsDeploymentConfigsProductionEnvVarsTypeSecretText ProjectNewParamsDeploymentConfigsProductionEnvVarsType = "secret_text"
)

func (r ProjectNewParamsDeploymentConfigsProductionEnvVarsType) IsKnown() bool {
	switch r {
	case ProjectNewParamsDeploymentConfigsProductionEnvVarsTypePlainText, ProjectNewParamsDeploymentConfigsProductionEnvVarsTypeSecretText:
		return true
	}
	return false
}

// Hyperdrive binding.
type ProjectNewParamsDeploymentConfigsProductionHyperdriveBindings struct {
	ID param.Field[string] `json:"id,required"`
}

func (r ProjectNewParamsDeploymentConfigsProductionHyperdriveBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// KV namespace binding.
type ProjectNewParamsDeploymentConfigsProductionKVNamespaces struct {
	// ID of the KV namespace.
	NamespaceID param.Field[string] `json:"namespace_id,required"`
}

func (r ProjectNewParamsDeploymentConfigsProductionKVNamespaces) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Limits for Pages Functions.
type ProjectNewParamsDeploymentConfigsProductionLimits struct {
	// CPU time limit in milliseconds.
	CPUMs param.Field[int64] `json:"cpu_ms,required"`
}

func (r ProjectNewParamsDeploymentConfigsProductionLimits) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// mTLS binding.
type ProjectNewParamsDeploymentConfigsProductionMTLSCertificates struct {
	CertificateID param.Field[string] `json:"certificate_id,required"`
}

func (r ProjectNewParamsDeploymentConfigsProductionMTLSCertificates) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Placement setting used for Pages Functions.
type ProjectNewParamsDeploymentConfigsProductionPlacement struct {
	// Placement mode.
	Mode param.Field[string] `json:"mode,required"`
}

func (r ProjectNewParamsDeploymentConfigsProductionPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Queue Producer binding.
type ProjectNewParamsDeploymentConfigsProductionQueueProducers struct {
	// Name of the Queue.
	Name param.Field[string] `json:"name,required"`
}

func (r ProjectNewParamsDeploymentConfigsProductionQueueProducers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// R2 binding.
type ProjectNewParamsDeploymentConfigsProductionR2Buckets struct {
	// Name of the R2 bucket.
	Name param.Field[string] `json:"name,required"`
	// Jurisdiction of the R2 bucket.
	Jurisdiction param.Field[string] `json:"jurisdiction"`
}

func (r ProjectNewParamsDeploymentConfigsProductionR2Buckets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Service binding.
type ProjectNewParamsDeploymentConfigsProductionServices struct {
	// The Service name.
	Service param.Field[string] `json:"service,required"`
	// The entrypoint to bind to.
	Entrypoint param.Field[string] `json:"entrypoint"`
	// The Service environment.
	Environment param.Field[string] `json:"environment"`
}

func (r ProjectNewParamsDeploymentConfigsProductionServices) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The usage model for Pages Functions.
type ProjectNewParamsDeploymentConfigsProductionUsageModel string

const (
	ProjectNewParamsDeploymentConfigsProductionUsageModelStandard ProjectNewParamsDeploymentConfigsProductionUsageModel = "standard"
	ProjectNewParamsDeploymentConfigsProductionUsageModelBundled  ProjectNewParamsDeploymentConfigsProductionUsageModel = "bundled"
	ProjectNewParamsDeploymentConfigsProductionUsageModelUnbound  ProjectNewParamsDeploymentConfigsProductionUsageModel = "unbound"
)

func (r ProjectNewParamsDeploymentConfigsProductionUsageModel) IsKnown() bool {
	switch r {
	case ProjectNewParamsDeploymentConfigsProductionUsageModelStandard, ProjectNewParamsDeploymentConfigsProductionUsageModelBundled, ProjectNewParamsDeploymentConfigsProductionUsageModelUnbound:
		return true
	}
	return false
}

// Vectorize binding.
type ProjectNewParamsDeploymentConfigsProductionVectorizeBindings struct {
	IndexName param.Field[string] `json:"index_name,required"`
}

func (r ProjectNewParamsDeploymentConfigsProductionVectorizeBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for the project source control.
type ProjectNewParamsSource struct {
	Config param.Field[ProjectNewParamsSourceConfig] `json:"config,required"`
	// The source control management provider.
	Type param.Field[ProjectNewParamsSourceType] `json:"type,required"`
}

func (r ProjectNewParamsSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectNewParamsSourceConfig struct {
	// Whether to enable automatic deployments when pushing to the source repository.
	// When disabled, no deployments (production or preview) will be triggered
	// automatically.
	//
	// Deprecated: Use `production_deployments_enabled` and
	// `preview_deployment_setting` for more granular control.
	DeploymentsEnabled param.Field[bool] `json:"deployments_enabled"`
	// The owner of the repository.
	Owner param.Field[string] `json:"owner"`
	// The owner ID of the repository.
	OwnerID param.Field[string] `json:"owner_id"`
	// A list of paths that should be excluded from triggering a preview deployment.
	// Wildcard syntax (`*`) is supported.
	PathExcludes param.Field[[]string] `json:"path_excludes"`
	// A list of paths that should be watched to trigger a preview deployment. Wildcard
	// syntax (`*`) is supported.
	PathIncludes param.Field[[]string] `json:"path_includes"`
	// Whether to enable PR comments.
	PrCommentsEnabled param.Field[bool] `json:"pr_comments_enabled"`
	// A list of branches that should not trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchExcludes param.Field[[]string] `json:"preview_branch_excludes"`
	// A list of branches that should trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchIncludes param.Field[[]string] `json:"preview_branch_includes"`
	// Controls whether commits to preview branches trigger a preview deployment.
	PreviewDeploymentSetting param.Field[ProjectNewParamsSourceConfigPreviewDeploymentSetting] `json:"preview_deployment_setting"`
	// The production branch of the repository.
	ProductionBranch param.Field[string] `json:"production_branch"`
	// Whether to trigger a production deployment on commits to the production branch.
	ProductionDeploymentsEnabled param.Field[bool] `json:"production_deployments_enabled"`
	// The ID of the repository.
	RepoID param.Field[string] `json:"repo_id"`
	// The name of the repository.
	RepoName param.Field[string] `json:"repo_name"`
}

func (r ProjectNewParamsSourceConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls whether commits to preview branches trigger a preview deployment.
type ProjectNewParamsSourceConfigPreviewDeploymentSetting string

const (
	ProjectNewParamsSourceConfigPreviewDeploymentSettingAll    ProjectNewParamsSourceConfigPreviewDeploymentSetting = "all"
	ProjectNewParamsSourceConfigPreviewDeploymentSettingNone   ProjectNewParamsSourceConfigPreviewDeploymentSetting = "none"
	ProjectNewParamsSourceConfigPreviewDeploymentSettingCustom ProjectNewParamsSourceConfigPreviewDeploymentSetting = "custom"
)

func (r ProjectNewParamsSourceConfigPreviewDeploymentSetting) IsKnown() bool {
	switch r {
	case ProjectNewParamsSourceConfigPreviewDeploymentSettingAll, ProjectNewParamsSourceConfigPreviewDeploymentSettingNone, ProjectNewParamsSourceConfigPreviewDeploymentSettingCustom:
		return true
	}
	return false
}

// The source control management provider.
type ProjectNewParamsSourceType string

const (
	ProjectNewParamsSourceTypeGitHub ProjectNewParamsSourceType = "github"
	ProjectNewParamsSourceTypeGitlab ProjectNewParamsSourceType = "gitlab"
)

func (r ProjectNewParamsSourceType) IsKnown() bool {
	switch r {
	case ProjectNewParamsSourceTypeGitHub, ProjectNewParamsSourceTypeGitlab:
		return true
	}
	return false
}

type ProjectNewResponseEnvelope struct {
	Errors   []ProjectNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectNewResponseEnvelopeMessages `json:"messages,required"`
	Result   Project                              `json:"result,required"`
	// Whether the API call was successful.
	Success ProjectNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectNewResponseEnvelopeJSON    `json:"-"`
}

// projectNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ProjectNewResponseEnvelope]
type projectNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ProjectNewResponseEnvelopeErrors struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           ProjectNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             projectNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// projectNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ProjectNewResponseEnvelopeErrors]
type projectNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectNewResponseEnvelopeErrorsSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    projectNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// projectNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [ProjectNewResponseEnvelopeErrorsSource]
type projectNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectNewResponseEnvelopeMessages struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           ProjectNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             projectNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// projectNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [ProjectNewResponseEnvelopeMessages]
type projectNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ProjectNewResponseEnvelopeMessagesSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    projectNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// projectNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [ProjectNewResponseEnvelopeMessagesSource]
type projectNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ProjectNewResponseEnvelopeSuccess bool

const (
	ProjectNewResponseEnvelopeSuccessTrue ProjectNewResponseEnvelopeSuccess = true
)

func (r ProjectNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Which page of projects to fetch.
	Page param.Field[int64] `query:"page"`
	// How many projects to return per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [ProjectListParams]'s query parameters as `url.Values`.
func (r ProjectListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ProjectDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDeleteResponseEnvelope struct {
	Errors   []ProjectDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ProjectDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success ProjectDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDeleteResponseEnvelopeJSON    `json:"-"`
}

// projectDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [ProjectDeleteResponseEnvelope]
type projectDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ProjectDeleteResponseEnvelopeErrors struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           ProjectDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             projectDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// projectDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ProjectDeleteResponseEnvelopeErrors]
type projectDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    projectDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// projectDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [ProjectDeleteResponseEnvelopeErrorsSource]
type projectDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectDeleteResponseEnvelopeMessages struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           ProjectDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             projectDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// projectDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ProjectDeleteResponseEnvelopeMessages]
type projectDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ProjectDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    projectDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// projectDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [ProjectDeleteResponseEnvelopeMessagesSource]
type projectDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ProjectDeleteResponseEnvelopeSuccess bool

const (
	ProjectDeleteResponseEnvelopeSuccessTrue ProjectDeleteResponseEnvelopeSuccess = true
)

func (r ProjectDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectEditParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Configs for the project build process.
	BuildConfig param.Field[ProjectEditParamsBuildConfig] `json:"build_config"`
	// Configs for deployments in a project.
	DeploymentConfigs param.Field[ProjectEditParamsDeploymentConfigs] `json:"deployment_configs"`
	// Name of the project.
	Name param.Field[string] `json:"name"`
	// Production branch of the project. Used to identify production deployments.
	ProductionBranch param.Field[string] `json:"production_branch"`
	// Configs for the project source control.
	Source param.Field[ProjectEditParamsSource] `json:"source"`
}

func (r ProjectEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for the project build process.
type ProjectEditParamsBuildConfig struct {
	// Enable build caching for the project.
	BuildCaching param.Field[bool] `json:"build_caching"`
	// Command used to build project.
	BuildCommand param.Field[string] `json:"build_command"`
	// Output directory of the build.
	DestinationDir param.Field[string] `json:"destination_dir"`
	// Directory to run the command.
	RootDir param.Field[string] `json:"root_dir"`
	// The classifying tag for analytics.
	WebAnalyticsTag param.Field[string] `json:"web_analytics_tag"`
	// The auth token for analytics.
	WebAnalyticsToken param.Field[string] `json:"web_analytics_token"`
}

func (r ProjectEditParamsBuildConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for deployments in a project.
type ProjectEditParamsDeploymentConfigs struct {
	// Configs for preview deploys.
	Preview param.Field[ProjectEditParamsDeploymentConfigsPreview] `json:"preview"`
	// Configs for production deploys.
	Production param.Field[ProjectEditParamsDeploymentConfigsProduction] `json:"production"`
}

func (r ProjectEditParamsDeploymentConfigs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for preview deploys.
type ProjectEditParamsDeploymentConfigsPreview struct {
	// Constellation bindings used for Pages Functions.
	AIBindings param.Field[map[string]ProjectEditParamsDeploymentConfigsPreviewAIBindings] `json:"ai_bindings"`
	// Whether to always use the latest compatibility date for Pages Functions.
	AlwaysUseLatestCompatibilityDate param.Field[bool] `json:"always_use_latest_compatibility_date"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets param.Field[map[string]ProjectEditParamsDeploymentConfigsPreviewAnalyticsEngineDatasets] `json:"analytics_engine_datasets"`
	// Browser bindings used for Pages Functions.
	Browsers param.Field[map[string]ProjectEditParamsDeploymentConfigsPreviewBrowsers] `json:"browsers"`
	// The major version of the build image to use for Pages Functions.
	BuildImageMajorVersion param.Field[int64] `json:"build_image_major_version"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases param.Field[map[string]ProjectEditParamsDeploymentConfigsPreviewD1Databases] `json:"d1_databases"`
	// Durable Object namespaces used for Pages Functions.
	DurableObjectNamespaces param.Field[map[string]ProjectEditParamsDeploymentConfigsPreviewDurableObjectNamespaces] `json:"durable_object_namespaces"`
	// Environment variables used for builds and Pages Functions.
	EnvVars param.Field[map[string]ProjectEditParamsDeploymentConfigsPreviewEnvVarsUnion] `json:"env_vars"`
	// Whether to fail open when the deployment config cannot be applied.
	FailOpen param.Field[bool] `json:"fail_open"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings param.Field[map[string]ProjectEditParamsDeploymentConfigsPreviewHyperdriveBindings] `json:"hyperdrive_bindings"`
	// KV namespaces used for Pages Functions.
	KVNamespaces param.Field[map[string]ProjectEditParamsDeploymentConfigsPreviewKVNamespaces] `json:"kv_namespaces"`
	// Limits for Pages Functions.
	Limits param.Field[ProjectEditParamsDeploymentConfigsPreviewLimits] `json:"limits"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates param.Field[map[string]ProjectEditParamsDeploymentConfigsPreviewMTLSCertificates] `json:"mtls_certificates"`
	// Placement setting used for Pages Functions.
	Placement param.Field[ProjectEditParamsDeploymentConfigsPreviewPlacement] `json:"placement"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers param.Field[map[string]ProjectEditParamsDeploymentConfigsPreviewQueueProducers] `json:"queue_producers"`
	// R2 buckets used for Pages Functions.
	R2Buckets param.Field[map[string]ProjectEditParamsDeploymentConfigsPreviewR2Buckets] `json:"r2_buckets"`
	// Services used for Pages Functions.
	Services param.Field[map[string]ProjectEditParamsDeploymentConfigsPreviewServices] `json:"services"`
	// The usage model for Pages Functions.
	//
	// Deprecated: All new projects now use the Standard usage model.
	UsageModel param.Field[ProjectEditParamsDeploymentConfigsPreviewUsageModel] `json:"usage_model"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings param.Field[map[string]ProjectEditParamsDeploymentConfigsPreviewVectorizeBindings] `json:"vectorize_bindings"`
	// Hash of the Wrangler configuration used for the deployment.
	WranglerConfigHash param.Field[string] `json:"wrangler_config_hash"`
}

func (r ProjectEditParamsDeploymentConfigsPreview) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AI binding.
type ProjectEditParamsDeploymentConfigsPreviewAIBindings struct {
	ProjectID param.Field[string] `json:"project_id,required"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewAIBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Analytics Engine binding.
type ProjectEditParamsDeploymentConfigsPreviewAnalyticsEngineDatasets struct {
	// Name of the dataset.
	Dataset param.Field[string] `json:"dataset,required"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewAnalyticsEngineDatasets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Browser binding.
type ProjectEditParamsDeploymentConfigsPreviewBrowsers struct {
}

func (r ProjectEditParamsDeploymentConfigsPreviewBrowsers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// D1 binding.
type ProjectEditParamsDeploymentConfigsPreviewD1Databases struct {
	// UUID of the D1 database.
	ID param.Field[string] `json:"id,required"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewD1Databases) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Durable Object binding.
type ProjectEditParamsDeploymentConfigsPreviewDurableObjectNamespaces struct {
	// ID of the Durable Object namespace.
	NamespaceID param.Field[string] `json:"namespace_id,required"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewDurableObjectNamespaces) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A plaintext environment variable.
type ProjectEditParamsDeploymentConfigsPreviewEnvVars struct {
	Type param.Field[ProjectEditParamsDeploymentConfigsPreviewEnvVarsType] `json:"type,required"`
	// Environment variable value.
	Value param.Field[string] `json:"value,required"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewEnvVars) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectEditParamsDeploymentConfigsPreviewEnvVars) implementsProjectEditParamsDeploymentConfigsPreviewEnvVarsUnion() {
}

// A plaintext environment variable.
//
// Satisfied by
// [pages.ProjectEditParamsDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar],
// [pages.ProjectEditParamsDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar],
// [ProjectEditParamsDeploymentConfigsPreviewEnvVars].
type ProjectEditParamsDeploymentConfigsPreviewEnvVarsUnion interface {
	implementsProjectEditParamsDeploymentConfigsPreviewEnvVarsUnion()
}

// A plaintext environment variable.
type ProjectEditParamsDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar struct {
	Type param.Field[ProjectEditParamsDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType] `json:"type,required"`
	// Environment variable value.
	Value param.Field[string] `json:"value,required"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectEditParamsDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar) implementsProjectEditParamsDeploymentConfigsPreviewEnvVarsUnion() {
}

type ProjectEditParamsDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType string

const (
	ProjectEditParamsDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarTypePlainText ProjectEditParamsDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r ProjectEditParamsDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectEditParamsDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type ProjectEditParamsDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar struct {
	Type param.Field[ProjectEditParamsDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType] `json:"type,required"`
	// Secret value.
	Value param.Field[string] `json:"value,required"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectEditParamsDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar) implementsProjectEditParamsDeploymentConfigsPreviewEnvVarsUnion() {
}

type ProjectEditParamsDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType string

const (
	ProjectEditParamsDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarTypeSecretText ProjectEditParamsDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r ProjectEditParamsDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectEditParamsDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type ProjectEditParamsDeploymentConfigsPreviewEnvVarsType string

const (
	ProjectEditParamsDeploymentConfigsPreviewEnvVarsTypePlainText  ProjectEditParamsDeploymentConfigsPreviewEnvVarsType = "plain_text"
	ProjectEditParamsDeploymentConfigsPreviewEnvVarsTypeSecretText ProjectEditParamsDeploymentConfigsPreviewEnvVarsType = "secret_text"
)

func (r ProjectEditParamsDeploymentConfigsPreviewEnvVarsType) IsKnown() bool {
	switch r {
	case ProjectEditParamsDeploymentConfigsPreviewEnvVarsTypePlainText, ProjectEditParamsDeploymentConfigsPreviewEnvVarsTypeSecretText:
		return true
	}
	return false
}

// Hyperdrive binding.
type ProjectEditParamsDeploymentConfigsPreviewHyperdriveBindings struct {
	ID param.Field[string] `json:"id,required"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewHyperdriveBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// KV namespace binding.
type ProjectEditParamsDeploymentConfigsPreviewKVNamespaces struct {
	// ID of the KV namespace.
	NamespaceID param.Field[string] `json:"namespace_id,required"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewKVNamespaces) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Limits for Pages Functions.
type ProjectEditParamsDeploymentConfigsPreviewLimits struct {
	// CPU time limit in milliseconds.
	CPUMs param.Field[int64] `json:"cpu_ms,required"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewLimits) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// mTLS binding.
type ProjectEditParamsDeploymentConfigsPreviewMTLSCertificates struct {
	CertificateID param.Field[string] `json:"certificate_id,required"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewMTLSCertificates) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Placement setting used for Pages Functions.
type ProjectEditParamsDeploymentConfigsPreviewPlacement struct {
	// Placement mode.
	Mode param.Field[string] `json:"mode,required"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Queue Producer binding.
type ProjectEditParamsDeploymentConfigsPreviewQueueProducers struct {
	// Name of the Queue.
	Name param.Field[string] `json:"name,required"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewQueueProducers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// R2 binding.
type ProjectEditParamsDeploymentConfigsPreviewR2Buckets struct {
	// Name of the R2 bucket.
	Name param.Field[string] `json:"name,required"`
	// Jurisdiction of the R2 bucket.
	Jurisdiction param.Field[string] `json:"jurisdiction"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewR2Buckets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Service binding.
type ProjectEditParamsDeploymentConfigsPreviewServices struct {
	// The Service name.
	Service param.Field[string] `json:"service,required"`
	// The entrypoint to bind to.
	Entrypoint param.Field[string] `json:"entrypoint"`
	// The Service environment.
	Environment param.Field[string] `json:"environment"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewServices) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The usage model for Pages Functions.
type ProjectEditParamsDeploymentConfigsPreviewUsageModel string

const (
	ProjectEditParamsDeploymentConfigsPreviewUsageModelStandard ProjectEditParamsDeploymentConfigsPreviewUsageModel = "standard"
	ProjectEditParamsDeploymentConfigsPreviewUsageModelBundled  ProjectEditParamsDeploymentConfigsPreviewUsageModel = "bundled"
	ProjectEditParamsDeploymentConfigsPreviewUsageModelUnbound  ProjectEditParamsDeploymentConfigsPreviewUsageModel = "unbound"
)

func (r ProjectEditParamsDeploymentConfigsPreviewUsageModel) IsKnown() bool {
	switch r {
	case ProjectEditParamsDeploymentConfigsPreviewUsageModelStandard, ProjectEditParamsDeploymentConfigsPreviewUsageModelBundled, ProjectEditParamsDeploymentConfigsPreviewUsageModelUnbound:
		return true
	}
	return false
}

// Vectorize binding.
type ProjectEditParamsDeploymentConfigsPreviewVectorizeBindings struct {
	IndexName param.Field[string] `json:"index_name,required"`
}

func (r ProjectEditParamsDeploymentConfigsPreviewVectorizeBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for production deploys.
type ProjectEditParamsDeploymentConfigsProduction struct {
	// Constellation bindings used for Pages Functions.
	AIBindings param.Field[map[string]ProjectEditParamsDeploymentConfigsProductionAIBindings] `json:"ai_bindings"`
	// Whether to always use the latest compatibility date for Pages Functions.
	AlwaysUseLatestCompatibilityDate param.Field[bool] `json:"always_use_latest_compatibility_date"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets param.Field[map[string]ProjectEditParamsDeploymentConfigsProductionAnalyticsEngineDatasets] `json:"analytics_engine_datasets"`
	// Browser bindings used for Pages Functions.
	Browsers param.Field[map[string]ProjectEditParamsDeploymentConfigsProductionBrowsers] `json:"browsers"`
	// The major version of the build image to use for Pages Functions.
	BuildImageMajorVersion param.Field[int64] `json:"build_image_major_version"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate param.Field[string] `json:"compatibility_date"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags param.Field[[]string] `json:"compatibility_flags"`
	// D1 databases used for Pages Functions.
	D1Databases param.Field[map[string]ProjectEditParamsDeploymentConfigsProductionD1Databases] `json:"d1_databases"`
	// Durable Object namespaces used for Pages Functions.
	DurableObjectNamespaces param.Field[map[string]ProjectEditParamsDeploymentConfigsProductionDurableObjectNamespaces] `json:"durable_object_namespaces"`
	// Environment variables used for builds and Pages Functions.
	EnvVars param.Field[map[string]ProjectEditParamsDeploymentConfigsProductionEnvVarsUnion] `json:"env_vars"`
	// Whether to fail open when the deployment config cannot be applied.
	FailOpen param.Field[bool] `json:"fail_open"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings param.Field[map[string]ProjectEditParamsDeploymentConfigsProductionHyperdriveBindings] `json:"hyperdrive_bindings"`
	// KV namespaces used for Pages Functions.
	KVNamespaces param.Field[map[string]ProjectEditParamsDeploymentConfigsProductionKVNamespaces] `json:"kv_namespaces"`
	// Limits for Pages Functions.
	Limits param.Field[ProjectEditParamsDeploymentConfigsProductionLimits] `json:"limits"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates param.Field[map[string]ProjectEditParamsDeploymentConfigsProductionMTLSCertificates] `json:"mtls_certificates"`
	// Placement setting used for Pages Functions.
	Placement param.Field[ProjectEditParamsDeploymentConfigsProductionPlacement] `json:"placement"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers param.Field[map[string]ProjectEditParamsDeploymentConfigsProductionQueueProducers] `json:"queue_producers"`
	// R2 buckets used for Pages Functions.
	R2Buckets param.Field[map[string]ProjectEditParamsDeploymentConfigsProductionR2Buckets] `json:"r2_buckets"`
	// Services used for Pages Functions.
	Services param.Field[map[string]ProjectEditParamsDeploymentConfigsProductionServices] `json:"services"`
	// The usage model for Pages Functions.
	//
	// Deprecated: All new projects now use the Standard usage model.
	UsageModel param.Field[ProjectEditParamsDeploymentConfigsProductionUsageModel] `json:"usage_model"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings param.Field[map[string]ProjectEditParamsDeploymentConfigsProductionVectorizeBindings] `json:"vectorize_bindings"`
	// Hash of the Wrangler configuration used for the deployment.
	WranglerConfigHash param.Field[string] `json:"wrangler_config_hash"`
}

func (r ProjectEditParamsDeploymentConfigsProduction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AI binding.
type ProjectEditParamsDeploymentConfigsProductionAIBindings struct {
	ProjectID param.Field[string] `json:"project_id,required"`
}

func (r ProjectEditParamsDeploymentConfigsProductionAIBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Analytics Engine binding.
type ProjectEditParamsDeploymentConfigsProductionAnalyticsEngineDatasets struct {
	// Name of the dataset.
	Dataset param.Field[string] `json:"dataset,required"`
}

func (r ProjectEditParamsDeploymentConfigsProductionAnalyticsEngineDatasets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Browser binding.
type ProjectEditParamsDeploymentConfigsProductionBrowsers struct {
}

func (r ProjectEditParamsDeploymentConfigsProductionBrowsers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// D1 binding.
type ProjectEditParamsDeploymentConfigsProductionD1Databases struct {
	// UUID of the D1 database.
	ID param.Field[string] `json:"id,required"`
}

func (r ProjectEditParamsDeploymentConfigsProductionD1Databases) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Durable Object binding.
type ProjectEditParamsDeploymentConfigsProductionDurableObjectNamespaces struct {
	// ID of the Durable Object namespace.
	NamespaceID param.Field[string] `json:"namespace_id,required"`
}

func (r ProjectEditParamsDeploymentConfigsProductionDurableObjectNamespaces) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A plaintext environment variable.
type ProjectEditParamsDeploymentConfigsProductionEnvVars struct {
	Type param.Field[ProjectEditParamsDeploymentConfigsProductionEnvVarsType] `json:"type,required"`
	// Environment variable value.
	Value param.Field[string] `json:"value,required"`
}

func (r ProjectEditParamsDeploymentConfigsProductionEnvVars) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectEditParamsDeploymentConfigsProductionEnvVars) implementsProjectEditParamsDeploymentConfigsProductionEnvVarsUnion() {
}

// A plaintext environment variable.
//
// Satisfied by
// [pages.ProjectEditParamsDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar],
// [pages.ProjectEditParamsDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar],
// [ProjectEditParamsDeploymentConfigsProductionEnvVars].
type ProjectEditParamsDeploymentConfigsProductionEnvVarsUnion interface {
	implementsProjectEditParamsDeploymentConfigsProductionEnvVarsUnion()
}

// A plaintext environment variable.
type ProjectEditParamsDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar struct {
	Type param.Field[ProjectEditParamsDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType] `json:"type,required"`
	// Environment variable value.
	Value param.Field[string] `json:"value,required"`
}

func (r ProjectEditParamsDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectEditParamsDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar) implementsProjectEditParamsDeploymentConfigsProductionEnvVarsUnion() {
}

type ProjectEditParamsDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType string

const (
	ProjectEditParamsDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarTypePlainText ProjectEditParamsDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r ProjectEditParamsDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectEditParamsDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type ProjectEditParamsDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar struct {
	Type param.Field[ProjectEditParamsDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType] `json:"type,required"`
	// Secret value.
	Value param.Field[string] `json:"value,required"`
}

func (r ProjectEditParamsDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectEditParamsDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar) implementsProjectEditParamsDeploymentConfigsProductionEnvVarsUnion() {
}

type ProjectEditParamsDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType string

const (
	ProjectEditParamsDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarTypeSecretText ProjectEditParamsDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r ProjectEditParamsDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectEditParamsDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type ProjectEditParamsDeploymentConfigsProductionEnvVarsType string

const (
	ProjectEditParamsDeploymentConfigsProductionEnvVarsTypePlainText  ProjectEditParamsDeploymentConfigsProductionEnvVarsType = "plain_text"
	ProjectEditParamsDeploymentConfigsProductionEnvVarsTypeSecretText ProjectEditParamsDeploymentConfigsProductionEnvVarsType = "secret_text"
)

func (r ProjectEditParamsDeploymentConfigsProductionEnvVarsType) IsKnown() bool {
	switch r {
	case ProjectEditParamsDeploymentConfigsProductionEnvVarsTypePlainText, ProjectEditParamsDeploymentConfigsProductionEnvVarsTypeSecretText:
		return true
	}
	return false
}

// Hyperdrive binding.
type ProjectEditParamsDeploymentConfigsProductionHyperdriveBindings struct {
	ID param.Field[string] `json:"id,required"`
}

func (r ProjectEditParamsDeploymentConfigsProductionHyperdriveBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// KV namespace binding.
type ProjectEditParamsDeploymentConfigsProductionKVNamespaces struct {
	// ID of the KV namespace.
	NamespaceID param.Field[string] `json:"namespace_id,required"`
}

func (r ProjectEditParamsDeploymentConfigsProductionKVNamespaces) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Limits for Pages Functions.
type ProjectEditParamsDeploymentConfigsProductionLimits struct {
	// CPU time limit in milliseconds.
	CPUMs param.Field[int64] `json:"cpu_ms,required"`
}

func (r ProjectEditParamsDeploymentConfigsProductionLimits) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// mTLS binding.
type ProjectEditParamsDeploymentConfigsProductionMTLSCertificates struct {
	CertificateID param.Field[string] `json:"certificate_id,required"`
}

func (r ProjectEditParamsDeploymentConfigsProductionMTLSCertificates) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Placement setting used for Pages Functions.
type ProjectEditParamsDeploymentConfigsProductionPlacement struct {
	// Placement mode.
	Mode param.Field[string] `json:"mode,required"`
}

func (r ProjectEditParamsDeploymentConfigsProductionPlacement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Queue Producer binding.
type ProjectEditParamsDeploymentConfigsProductionQueueProducers struct {
	// Name of the Queue.
	Name param.Field[string] `json:"name,required"`
}

func (r ProjectEditParamsDeploymentConfigsProductionQueueProducers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// R2 binding.
type ProjectEditParamsDeploymentConfigsProductionR2Buckets struct {
	// Name of the R2 bucket.
	Name param.Field[string] `json:"name,required"`
	// Jurisdiction of the R2 bucket.
	Jurisdiction param.Field[string] `json:"jurisdiction"`
}

func (r ProjectEditParamsDeploymentConfigsProductionR2Buckets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Service binding.
type ProjectEditParamsDeploymentConfigsProductionServices struct {
	// The Service name.
	Service param.Field[string] `json:"service,required"`
	// The entrypoint to bind to.
	Entrypoint param.Field[string] `json:"entrypoint"`
	// The Service environment.
	Environment param.Field[string] `json:"environment"`
}

func (r ProjectEditParamsDeploymentConfigsProductionServices) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The usage model for Pages Functions.
type ProjectEditParamsDeploymentConfigsProductionUsageModel string

const (
	ProjectEditParamsDeploymentConfigsProductionUsageModelStandard ProjectEditParamsDeploymentConfigsProductionUsageModel = "standard"
	ProjectEditParamsDeploymentConfigsProductionUsageModelBundled  ProjectEditParamsDeploymentConfigsProductionUsageModel = "bundled"
	ProjectEditParamsDeploymentConfigsProductionUsageModelUnbound  ProjectEditParamsDeploymentConfigsProductionUsageModel = "unbound"
)

func (r ProjectEditParamsDeploymentConfigsProductionUsageModel) IsKnown() bool {
	switch r {
	case ProjectEditParamsDeploymentConfigsProductionUsageModelStandard, ProjectEditParamsDeploymentConfigsProductionUsageModelBundled, ProjectEditParamsDeploymentConfigsProductionUsageModelUnbound:
		return true
	}
	return false
}

// Vectorize binding.
type ProjectEditParamsDeploymentConfigsProductionVectorizeBindings struct {
	IndexName param.Field[string] `json:"index_name,required"`
}

func (r ProjectEditParamsDeploymentConfigsProductionVectorizeBindings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configs for the project source control.
type ProjectEditParamsSource struct {
	Config param.Field[ProjectEditParamsSourceConfig] `json:"config,required"`
	// The source control management provider.
	Type param.Field[ProjectEditParamsSourceType] `json:"type,required"`
}

func (r ProjectEditParamsSource) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectEditParamsSourceConfig struct {
	// Whether to enable automatic deployments when pushing to the source repository.
	// When disabled, no deployments (production or preview) will be triggered
	// automatically.
	//
	// Deprecated: Use `production_deployments_enabled` and
	// `preview_deployment_setting` for more granular control.
	DeploymentsEnabled param.Field[bool] `json:"deployments_enabled"`
	// The owner of the repository.
	Owner param.Field[string] `json:"owner"`
	// The owner ID of the repository.
	OwnerID param.Field[string] `json:"owner_id"`
	// A list of paths that should be excluded from triggering a preview deployment.
	// Wildcard syntax (`*`) is supported.
	PathExcludes param.Field[[]string] `json:"path_excludes"`
	// A list of paths that should be watched to trigger a preview deployment. Wildcard
	// syntax (`*`) is supported.
	PathIncludes param.Field[[]string] `json:"path_includes"`
	// Whether to enable PR comments.
	PrCommentsEnabled param.Field[bool] `json:"pr_comments_enabled"`
	// A list of branches that should not trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchExcludes param.Field[[]string] `json:"preview_branch_excludes"`
	// A list of branches that should trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchIncludes param.Field[[]string] `json:"preview_branch_includes"`
	// Controls whether commits to preview branches trigger a preview deployment.
	PreviewDeploymentSetting param.Field[ProjectEditParamsSourceConfigPreviewDeploymentSetting] `json:"preview_deployment_setting"`
	// The production branch of the repository.
	ProductionBranch param.Field[string] `json:"production_branch"`
	// Whether to trigger a production deployment on commits to the production branch.
	ProductionDeploymentsEnabled param.Field[bool] `json:"production_deployments_enabled"`
	// The ID of the repository.
	RepoID param.Field[string] `json:"repo_id"`
	// The name of the repository.
	RepoName param.Field[string] `json:"repo_name"`
}

func (r ProjectEditParamsSourceConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls whether commits to preview branches trigger a preview deployment.
type ProjectEditParamsSourceConfigPreviewDeploymentSetting string

const (
	ProjectEditParamsSourceConfigPreviewDeploymentSettingAll    ProjectEditParamsSourceConfigPreviewDeploymentSetting = "all"
	ProjectEditParamsSourceConfigPreviewDeploymentSettingNone   ProjectEditParamsSourceConfigPreviewDeploymentSetting = "none"
	ProjectEditParamsSourceConfigPreviewDeploymentSettingCustom ProjectEditParamsSourceConfigPreviewDeploymentSetting = "custom"
)

func (r ProjectEditParamsSourceConfigPreviewDeploymentSetting) IsKnown() bool {
	switch r {
	case ProjectEditParamsSourceConfigPreviewDeploymentSettingAll, ProjectEditParamsSourceConfigPreviewDeploymentSettingNone, ProjectEditParamsSourceConfigPreviewDeploymentSettingCustom:
		return true
	}
	return false
}

// The source control management provider.
type ProjectEditParamsSourceType string

const (
	ProjectEditParamsSourceTypeGitHub ProjectEditParamsSourceType = "github"
	ProjectEditParamsSourceTypeGitlab ProjectEditParamsSourceType = "gitlab"
)

func (r ProjectEditParamsSourceType) IsKnown() bool {
	switch r {
	case ProjectEditParamsSourceTypeGitHub, ProjectEditParamsSourceTypeGitlab:
		return true
	}
	return false
}

type ProjectEditResponseEnvelope struct {
	Errors   []ProjectEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectEditResponseEnvelopeMessages `json:"messages,required"`
	Result   Project                               `json:"result,required"`
	// Whether the API call was successful.
	Success ProjectEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectEditResponseEnvelopeJSON    `json:"-"`
}

// projectEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [ProjectEditResponseEnvelope]
type projectEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ProjectEditResponseEnvelopeErrors struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           ProjectEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             projectEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// projectEditResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ProjectEditResponseEnvelopeErrors]
type projectEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectEditResponseEnvelopeErrorsSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    projectEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// projectEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [ProjectEditResponseEnvelopeErrorsSource]
type projectEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectEditResponseEnvelopeMessages struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           ProjectEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             projectEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// projectEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ProjectEditResponseEnvelopeMessages]
type projectEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ProjectEditResponseEnvelopeMessagesSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    projectEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// projectEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [ProjectEditResponseEnvelopeMessagesSource]
type projectEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ProjectEditResponseEnvelopeSuccess bool

const (
	ProjectEditResponseEnvelopeSuccessTrue ProjectEditResponseEnvelopeSuccess = true
)

func (r ProjectEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectGetResponseEnvelope struct {
	Errors   []ProjectGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectGetResponseEnvelopeMessages `json:"messages,required"`
	Result   Project                              `json:"result,required"`
	// Whether the API call was successful.
	Success ProjectGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectGetResponseEnvelopeJSON    `json:"-"`
}

// projectGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ProjectGetResponseEnvelope]
type projectGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ProjectGetResponseEnvelopeErrors struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           ProjectGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             projectGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// projectGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ProjectGetResponseEnvelopeErrors]
type projectGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectGetResponseEnvelopeErrorsSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    projectGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// projectGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [ProjectGetResponseEnvelopeErrorsSource]
type projectGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectGetResponseEnvelopeMessages struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           ProjectGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             projectGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// projectGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [ProjectGetResponseEnvelopeMessages]
type projectGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ProjectGetResponseEnvelopeMessagesSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    projectGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// projectGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [ProjectGetResponseEnvelopeMessagesSource]
type projectGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ProjectGetResponseEnvelopeSuccess bool

const (
	ProjectGetResponseEnvelopeSuccessTrue ProjectGetResponseEnvelopeSuccess = true
)

func (r ProjectGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectPurgeBuildCacheParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectPurgeBuildCacheResponseEnvelope struct {
	Errors   []ProjectPurgeBuildCacheResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectPurgeBuildCacheResponseEnvelopeMessages `json:"messages,required"`
	Result   ProjectPurgeBuildCacheResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success ProjectPurgeBuildCacheResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectPurgeBuildCacheResponseEnvelopeJSON    `json:"-"`
}

// projectPurgeBuildCacheResponseEnvelopeJSON contains the JSON metadata for the
// struct [ProjectPurgeBuildCacheResponseEnvelope]
type projectPurgeBuildCacheResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectPurgeBuildCacheResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectPurgeBuildCacheResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ProjectPurgeBuildCacheResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           ProjectPurgeBuildCacheResponseEnvelopeErrorsSource `json:"source"`
	JSON             projectPurgeBuildCacheResponseEnvelopeErrorsJSON   `json:"-"`
}

// projectPurgeBuildCacheResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ProjectPurgeBuildCacheResponseEnvelopeErrors]
type projectPurgeBuildCacheResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectPurgeBuildCacheResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectPurgeBuildCacheResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectPurgeBuildCacheResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    projectPurgeBuildCacheResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// projectPurgeBuildCacheResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [ProjectPurgeBuildCacheResponseEnvelopeErrorsSource]
type projectPurgeBuildCacheResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectPurgeBuildCacheResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectPurgeBuildCacheResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectPurgeBuildCacheResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           ProjectPurgeBuildCacheResponseEnvelopeMessagesSource `json:"source"`
	JSON             projectPurgeBuildCacheResponseEnvelopeMessagesJSON   `json:"-"`
}

// projectPurgeBuildCacheResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ProjectPurgeBuildCacheResponseEnvelopeMessages]
type projectPurgeBuildCacheResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectPurgeBuildCacheResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectPurgeBuildCacheResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ProjectPurgeBuildCacheResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    projectPurgeBuildCacheResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// projectPurgeBuildCacheResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [ProjectPurgeBuildCacheResponseEnvelopeMessagesSource]
type projectPurgeBuildCacheResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectPurgeBuildCacheResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectPurgeBuildCacheResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ProjectPurgeBuildCacheResponseEnvelopeSuccess bool

const (
	ProjectPurgeBuildCacheResponseEnvelopeSuccessTrue ProjectPurgeBuildCacheResponseEnvelopeSuccess = true
)

func (r ProjectPurgeBuildCacheResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectPurgeBuildCacheResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
