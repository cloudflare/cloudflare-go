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
func (r *ProjectService) New(ctx context.Context, params ProjectNewParams, opts ...option.RequestOption) (res *ProjectNewResponse, err error) {
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
func (r *ProjectService) List(ctx context.Context, params ProjectListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[ProjectListResponse], err error) {
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
func (r *ProjectService) ListAutoPaging(ctx context.Context, params ProjectListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[ProjectListResponse] {
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
func (r *ProjectService) Edit(ctx context.Context, projectName string, params ProjectEditParams, opts ...option.RequestOption) (res *ProjectEditResponse, err error) {
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
func (r *ProjectService) Get(ctx context.Context, projectName string, query ProjectGetParams, opts ...option.RequestOption) (res *ProjectGetResponse, err error) {
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

type ProjectNewResponse struct {
	// ID of the project.
	ID string `json:"id,required"`
	// Most recent production deployment of the project.
	CanonicalDeployment ProjectNewResponseCanonicalDeployment `json:"canonical_deployment,required,nullable"`
	// When the project was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Configs for deployments in a project.
	DeploymentConfigs ProjectNewResponseDeploymentConfigs `json:"deployment_configs,required"`
	// Framework the project is using.
	Framework string `json:"framework,required"`
	// Version of the framework the project is using.
	FrameworkVersion string `json:"framework_version,required"`
	// Most recent deployment of the project.
	LatestDeployment ProjectNewResponseLatestDeployment `json:"latest_deployment,required,nullable"`
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
	BuildConfig ProjectNewResponseBuildConfig `json:"build_config"`
	// A list of associated custom domains for the project.
	Domains []string `json:"domains"`
	// Configs for the project source control.
	Source ProjectNewResponseSource `json:"source"`
	// The Cloudflare subdomain associated with the project.
	Subdomain string                 `json:"subdomain"`
	JSON      projectNewResponseJSON `json:"-"`
}

// projectNewResponseJSON contains the JSON metadata for the struct
// [ProjectNewResponse]
type projectNewResponseJSON struct {
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

func (r *ProjectNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseJSON) RawJSON() string {
	return r.raw
}

// Most recent production deployment of the project.
type ProjectNewResponseCanonicalDeployment struct {
	// Id of the deployment.
	ID string `json:"id,required"`
	// A list of alias URLs pointing to this deployment.
	Aliases []string `json:"aliases,required,nullable"`
	// Configs for the project build process.
	BuildConfig ProjectNewResponseCanonicalDeploymentBuildConfig `json:"build_config,required"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger ProjectNewResponseCanonicalDeploymentDeploymentTrigger `json:"deployment_trigger,required"`
	// Environment variables used for builds and Pages Functions.
	EnvVars map[string]ProjectNewResponseCanonicalDeploymentEnvVar `json:"env_vars,required,nullable"`
	// Type of deploy.
	Environment ProjectNewResponseCanonicalDeploymentEnvironment `json:"environment,required"`
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
	Source ProjectNewResponseCanonicalDeploymentSource `json:"source,required"`
	// List of past stages.
	Stages []Stage `json:"stages,required"`
	// The live URL to view this deployment.
	URL string `json:"url,required"`
	// Whether the deployment uses functions.
	UsesFunctions bool                                      `json:"uses_functions,nullable"`
	JSON          projectNewResponseCanonicalDeploymentJSON `json:"-"`
}

// projectNewResponseCanonicalDeploymentJSON contains the JSON metadata for the
// struct [ProjectNewResponseCanonicalDeployment]
type projectNewResponseCanonicalDeploymentJSON struct {
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

func (r *ProjectNewResponseCanonicalDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseCanonicalDeploymentJSON) RawJSON() string {
	return r.raw
}

// Configs for the project build process.
type ProjectNewResponseCanonicalDeploymentBuildConfig struct {
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
	RootDir string                                               `json:"root_dir,nullable"`
	JSON    projectNewResponseCanonicalDeploymentBuildConfigJSON `json:"-"`
}

// projectNewResponseCanonicalDeploymentBuildConfigJSON contains the JSON metadata
// for the struct [ProjectNewResponseCanonicalDeploymentBuildConfig]
type projectNewResponseCanonicalDeploymentBuildConfigJSON struct {
	WebAnalyticsTag   apijson.Field
	WebAnalyticsToken apijson.Field
	BuildCaching      apijson.Field
	BuildCommand      apijson.Field
	DestinationDir    apijson.Field
	RootDir           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectNewResponseCanonicalDeploymentBuildConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseCanonicalDeploymentBuildConfigJSON) RawJSON() string {
	return r.raw
}

// Info about what caused the deployment.
type ProjectNewResponseCanonicalDeploymentDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata ProjectNewResponseCanonicalDeploymentDeploymentTriggerMetadata `json:"metadata,required"`
	// What caused the deployment.
	Type ProjectNewResponseCanonicalDeploymentDeploymentTriggerType `json:"type,required"`
	JSON projectNewResponseCanonicalDeploymentDeploymentTriggerJSON `json:"-"`
}

// projectNewResponseCanonicalDeploymentDeploymentTriggerJSON contains the JSON
// metadata for the struct [ProjectNewResponseCanonicalDeploymentDeploymentTrigger]
type projectNewResponseCanonicalDeploymentDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseCanonicalDeploymentDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseCanonicalDeploymentDeploymentTriggerJSON) RawJSON() string {
	return r.raw
}

// Additional info about the trigger.
type ProjectNewResponseCanonicalDeploymentDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch,required"`
	// Whether the deployment trigger commit was dirty.
	CommitDirty bool `json:"commit_dirty,required"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash,required"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                             `json:"commit_message,required"`
	JSON          projectNewResponseCanonicalDeploymentDeploymentTriggerMetadataJSON `json:"-"`
}

// projectNewResponseCanonicalDeploymentDeploymentTriggerMetadataJSON contains the
// JSON metadata for the struct
// [ProjectNewResponseCanonicalDeploymentDeploymentTriggerMetadata]
type projectNewResponseCanonicalDeploymentDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitDirty   apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectNewResponseCanonicalDeploymentDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseCanonicalDeploymentDeploymentTriggerMetadataJSON) RawJSON() string {
	return r.raw
}

// What caused the deployment.
type ProjectNewResponseCanonicalDeploymentDeploymentTriggerType string

const (
	ProjectNewResponseCanonicalDeploymentDeploymentTriggerTypeGitHubPush ProjectNewResponseCanonicalDeploymentDeploymentTriggerType = "github:push"
	ProjectNewResponseCanonicalDeploymentDeploymentTriggerTypeADHoc      ProjectNewResponseCanonicalDeploymentDeploymentTriggerType = "ad_hoc"
	ProjectNewResponseCanonicalDeploymentDeploymentTriggerTypeDeployHook ProjectNewResponseCanonicalDeploymentDeploymentTriggerType = "deploy_hook"
)

func (r ProjectNewResponseCanonicalDeploymentDeploymentTriggerType) IsKnown() bool {
	switch r {
	case ProjectNewResponseCanonicalDeploymentDeploymentTriggerTypeGitHubPush, ProjectNewResponseCanonicalDeploymentDeploymentTriggerTypeADHoc, ProjectNewResponseCanonicalDeploymentDeploymentTriggerTypeDeployHook:
		return true
	}
	return false
}

// A plaintext environment variable.
type ProjectNewResponseCanonicalDeploymentEnvVar struct {
	Type ProjectNewResponseCanonicalDeploymentEnvVarsType `json:"type,required"`
	// Environment variable value.
	Value string                                          `json:"value,required"`
	JSON  projectNewResponseCanonicalDeploymentEnvVarJSON `json:"-"`
	union ProjectNewResponseCanonicalDeploymentEnvVarsUnion
}

// projectNewResponseCanonicalDeploymentEnvVarJSON contains the JSON metadata for
// the struct [ProjectNewResponseCanonicalDeploymentEnvVar]
type projectNewResponseCanonicalDeploymentEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectNewResponseCanonicalDeploymentEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectNewResponseCanonicalDeploymentEnvVar) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectNewResponseCanonicalDeploymentEnvVar{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectNewResponseCanonicalDeploymentEnvVarsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectNewResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVar],
// [ProjectNewResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVar].
func (r ProjectNewResponseCanonicalDeploymentEnvVar) AsUnion() ProjectNewResponseCanonicalDeploymentEnvVarsUnion {
	return r.union
}

// A plaintext environment variable.
//
// Union satisfied by
// [ProjectNewResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVar] or
// [ProjectNewResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVar].
type ProjectNewResponseCanonicalDeploymentEnvVarsUnion interface {
	implementsProjectNewResponseCanonicalDeploymentEnvVar()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectNewResponseCanonicalDeploymentEnvVarsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectNewResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVar{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectNewResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVar{}),
			DiscriminatorValue: "secret_text",
		},
	)
}

// A plaintext environment variable.
type ProjectNewResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVar struct {
	Type ProjectNewResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarType `json:"type,required"`
	// Environment variable value.
	Value string                                                               `json:"value,required"`
	JSON  projectNewResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarJSON `json:"-"`
}

// projectNewResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarJSON contains
// the JSON metadata for the struct
// [ProjectNewResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVar]
type projectNewResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectNewResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVar) implementsProjectNewResponseCanonicalDeploymentEnvVar() {
}

type ProjectNewResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarType string

const (
	ProjectNewResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarTypePlainText ProjectNewResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r ProjectNewResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectNewResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type ProjectNewResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVar struct {
	Type ProjectNewResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarType `json:"type,required"`
	// Secret value.
	Value string                                                                `json:"value,required"`
	JSON  projectNewResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarJSON `json:"-"`
}

// projectNewResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarJSON contains
// the JSON metadata for the struct
// [ProjectNewResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVar]
type projectNewResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectNewResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVar) implementsProjectNewResponseCanonicalDeploymentEnvVar() {
}

type ProjectNewResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarType string

const (
	ProjectNewResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarTypeSecretText ProjectNewResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r ProjectNewResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectNewResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type ProjectNewResponseCanonicalDeploymentEnvVarsType string

const (
	ProjectNewResponseCanonicalDeploymentEnvVarsTypePlainText  ProjectNewResponseCanonicalDeploymentEnvVarsType = "plain_text"
	ProjectNewResponseCanonicalDeploymentEnvVarsTypeSecretText ProjectNewResponseCanonicalDeploymentEnvVarsType = "secret_text"
)

func (r ProjectNewResponseCanonicalDeploymentEnvVarsType) IsKnown() bool {
	switch r {
	case ProjectNewResponseCanonicalDeploymentEnvVarsTypePlainText, ProjectNewResponseCanonicalDeploymentEnvVarsTypeSecretText:
		return true
	}
	return false
}

// Type of deploy.
type ProjectNewResponseCanonicalDeploymentEnvironment string

const (
	ProjectNewResponseCanonicalDeploymentEnvironmentPreview    ProjectNewResponseCanonicalDeploymentEnvironment = "preview"
	ProjectNewResponseCanonicalDeploymentEnvironmentProduction ProjectNewResponseCanonicalDeploymentEnvironment = "production"
)

func (r ProjectNewResponseCanonicalDeploymentEnvironment) IsKnown() bool {
	switch r {
	case ProjectNewResponseCanonicalDeploymentEnvironmentPreview, ProjectNewResponseCanonicalDeploymentEnvironmentProduction:
		return true
	}
	return false
}

// Configs for the project source control.
type ProjectNewResponseCanonicalDeploymentSource struct {
	Config ProjectNewResponseCanonicalDeploymentSourceConfig `json:"config,required"`
	// The source control management provider.
	Type ProjectNewResponseCanonicalDeploymentSourceType `json:"type,required"`
	JSON projectNewResponseCanonicalDeploymentSourceJSON `json:"-"`
}

// projectNewResponseCanonicalDeploymentSourceJSON contains the JSON metadata for
// the struct [ProjectNewResponseCanonicalDeploymentSource]
type projectNewResponseCanonicalDeploymentSourceJSON struct {
	Config      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseCanonicalDeploymentSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseCanonicalDeploymentSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectNewResponseCanonicalDeploymentSourceConfig struct {
	// The owner of the repository.
	Owner string `json:"owner,required"`
	// Whether to enable PR comments.
	PrCommentsEnabled bool `json:"pr_comments_enabled,required"`
	// The production branch of the repository.
	ProductionBranch string `json:"production_branch,required"`
	// The name of the repository.
	RepoName string `json:"repo_name,required"`
	// Whether to enable automatic deployments when pushing to the source repository.
	// When disabled, no deployments (production or preview) will be triggered
	// automatically.
	//
	// Deprecated: Use `production_deployments_enabled` and
	// `preview_deployment_setting` for more granular control.
	DeploymentsEnabled bool `json:"deployments_enabled"`
	// The owner ID of the repository.
	OwnerID string `json:"owner_id"`
	// A list of paths that should be excluded from triggering a preview deployment.
	// Wildcard syntax (`*`) is supported.
	PathExcludes []string `json:"path_excludes"`
	// A list of paths that should be watched to trigger a preview deployment. Wildcard
	// syntax (`*`) is supported.
	PathIncludes []string `json:"path_includes"`
	// A list of branches that should not trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchExcludes []string `json:"preview_branch_excludes"`
	// A list of branches that should trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchIncludes []string `json:"preview_branch_includes"`
	// Controls whether commits to preview branches trigger a preview deployment.
	PreviewDeploymentSetting ProjectNewResponseCanonicalDeploymentSourceConfigPreviewDeploymentSetting `json:"preview_deployment_setting"`
	// Whether to trigger a production deployment on commits to the production branch.
	ProductionDeploymentsEnabled bool `json:"production_deployments_enabled"`
	// The ID of the repository.
	RepoID string                                                `json:"repo_id"`
	JSON   projectNewResponseCanonicalDeploymentSourceConfigJSON `json:"-"`
}

// projectNewResponseCanonicalDeploymentSourceConfigJSON contains the JSON metadata
// for the struct [ProjectNewResponseCanonicalDeploymentSourceConfig]
type projectNewResponseCanonicalDeploymentSourceConfigJSON struct {
	Owner                        apijson.Field
	PrCommentsEnabled            apijson.Field
	ProductionBranch             apijson.Field
	RepoName                     apijson.Field
	DeploymentsEnabled           apijson.Field
	OwnerID                      apijson.Field
	PathExcludes                 apijson.Field
	PathIncludes                 apijson.Field
	PreviewBranchExcludes        apijson.Field
	PreviewBranchIncludes        apijson.Field
	PreviewDeploymentSetting     apijson.Field
	ProductionDeploymentsEnabled apijson.Field
	RepoID                       apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ProjectNewResponseCanonicalDeploymentSourceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseCanonicalDeploymentSourceConfigJSON) RawJSON() string {
	return r.raw
}

// Controls whether commits to preview branches trigger a preview deployment.
type ProjectNewResponseCanonicalDeploymentSourceConfigPreviewDeploymentSetting string

const (
	ProjectNewResponseCanonicalDeploymentSourceConfigPreviewDeploymentSettingAll    ProjectNewResponseCanonicalDeploymentSourceConfigPreviewDeploymentSetting = "all"
	ProjectNewResponseCanonicalDeploymentSourceConfigPreviewDeploymentSettingNone   ProjectNewResponseCanonicalDeploymentSourceConfigPreviewDeploymentSetting = "none"
	ProjectNewResponseCanonicalDeploymentSourceConfigPreviewDeploymentSettingCustom ProjectNewResponseCanonicalDeploymentSourceConfigPreviewDeploymentSetting = "custom"
)

func (r ProjectNewResponseCanonicalDeploymentSourceConfigPreviewDeploymentSetting) IsKnown() bool {
	switch r {
	case ProjectNewResponseCanonicalDeploymentSourceConfigPreviewDeploymentSettingAll, ProjectNewResponseCanonicalDeploymentSourceConfigPreviewDeploymentSettingNone, ProjectNewResponseCanonicalDeploymentSourceConfigPreviewDeploymentSettingCustom:
		return true
	}
	return false
}

// The source control management provider.
type ProjectNewResponseCanonicalDeploymentSourceType string

const (
	ProjectNewResponseCanonicalDeploymentSourceTypeGitHub ProjectNewResponseCanonicalDeploymentSourceType = "github"
	ProjectNewResponseCanonicalDeploymentSourceTypeGitlab ProjectNewResponseCanonicalDeploymentSourceType = "gitlab"
)

func (r ProjectNewResponseCanonicalDeploymentSourceType) IsKnown() bool {
	switch r {
	case ProjectNewResponseCanonicalDeploymentSourceTypeGitHub, ProjectNewResponseCanonicalDeploymentSourceTypeGitlab:
		return true
	}
	return false
}

// Configs for deployments in a project.
type ProjectNewResponseDeploymentConfigs struct {
	// Configs for preview deploys.
	Preview ProjectNewResponseDeploymentConfigsPreview `json:"preview,required"`
	// Configs for production deploys.
	Production ProjectNewResponseDeploymentConfigsProduction `json:"production,required"`
	JSON       projectNewResponseDeploymentConfigsJSON       `json:"-"`
}

// projectNewResponseDeploymentConfigsJSON contains the JSON metadata for the
// struct [ProjectNewResponseDeploymentConfigs]
type projectNewResponseDeploymentConfigsJSON struct {
	Preview     apijson.Field
	Production  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsJSON) RawJSON() string {
	return r.raw
}

// Configs for preview deploys.
type ProjectNewResponseDeploymentConfigsPreview struct {
	// Whether to always use the latest compatibility date for Pages Functions.
	AlwaysUseLatestCompatibilityDate bool `json:"always_use_latest_compatibility_date,required"`
	// The major version of the build image to use for Pages Functions.
	BuildImageMajorVersion int64 `json:"build_image_major_version,required"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate string `json:"compatibility_date,required"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags []string `json:"compatibility_flags,required"`
	// Environment variables used for builds and Pages Functions.
	EnvVars map[string]ProjectNewResponseDeploymentConfigsPreviewEnvVar `json:"env_vars,required,nullable"`
	// Whether to fail open when the deployment config cannot be applied.
	FailOpen bool `json:"fail_open,required"`
	// The usage model for Pages Functions.
	//
	// Deprecated: All new projects now use the Standard usage model.
	UsageModel ProjectNewResponseDeploymentConfigsPreviewUsageModel `json:"usage_model,required"`
	// Constellation bindings used for Pages Functions.
	AIBindings map[string]ProjectNewResponseDeploymentConfigsPreviewAIBinding `json:"ai_bindings"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets map[string]ProjectNewResponseDeploymentConfigsPreviewAnalyticsEngineDataset `json:"analytics_engine_datasets"`
	// Browser bindings used for Pages Functions.
	Browsers map[string]ProjectNewResponseDeploymentConfigsPreviewBrowser `json:"browsers"`
	// D1 databases used for Pages Functions.
	D1Databases map[string]ProjectNewResponseDeploymentConfigsPreviewD1Database `json:"d1_databases"`
	// Durable Object namespaces used for Pages Functions.
	DurableObjectNamespaces map[string]ProjectNewResponseDeploymentConfigsPreviewDurableObjectNamespace `json:"durable_object_namespaces"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings map[string]ProjectNewResponseDeploymentConfigsPreviewHyperdriveBinding `json:"hyperdrive_bindings"`
	// KV namespaces used for Pages Functions.
	KVNamespaces map[string]ProjectNewResponseDeploymentConfigsPreviewKVNamespace `json:"kv_namespaces"`
	// Limits for Pages Functions.
	Limits ProjectNewResponseDeploymentConfigsPreviewLimits `json:"limits"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates map[string]ProjectNewResponseDeploymentConfigsPreviewMTLSCertificate `json:"mtls_certificates"`
	// Placement setting used for Pages Functions.
	Placement ProjectNewResponseDeploymentConfigsPreviewPlacement `json:"placement"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers map[string]ProjectNewResponseDeploymentConfigsPreviewQueueProducer `json:"queue_producers"`
	// R2 buckets used for Pages Functions.
	R2Buckets map[string]ProjectNewResponseDeploymentConfigsPreviewR2Bucket `json:"r2_buckets"`
	// Services used for Pages Functions.
	Services map[string]ProjectNewResponseDeploymentConfigsPreviewService `json:"services"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings map[string]ProjectNewResponseDeploymentConfigsPreviewVectorizeBinding `json:"vectorize_bindings"`
	// Hash of the Wrangler configuration used for the deployment.
	WranglerConfigHash string                                         `json:"wrangler_config_hash"`
	JSON               projectNewResponseDeploymentConfigsPreviewJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewJSON contains the JSON metadata for
// the struct [ProjectNewResponseDeploymentConfigsPreview]
type projectNewResponseDeploymentConfigsPreviewJSON struct {
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

func (r *ProjectNewResponseDeploymentConfigsPreview) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewJSON) RawJSON() string {
	return r.raw
}

// A plaintext environment variable.
type ProjectNewResponseDeploymentConfigsPreviewEnvVar struct {
	Type ProjectNewResponseDeploymentConfigsPreviewEnvVarsType `json:"type,required"`
	// Environment variable value.
	Value string                                               `json:"value,required"`
	JSON  projectNewResponseDeploymentConfigsPreviewEnvVarJSON `json:"-"`
	union ProjectNewResponseDeploymentConfigsPreviewEnvVarsUnion
}

// projectNewResponseDeploymentConfigsPreviewEnvVarJSON contains the JSON metadata
// for the struct [ProjectNewResponseDeploymentConfigsPreviewEnvVar]
type projectNewResponseDeploymentConfigsPreviewEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectNewResponseDeploymentConfigsPreviewEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectNewResponseDeploymentConfigsPreviewEnvVar) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectNewResponseDeploymentConfigsPreviewEnvVar{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectNewResponseDeploymentConfigsPreviewEnvVarsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectNewResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar],
// [ProjectNewResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar].
func (r ProjectNewResponseDeploymentConfigsPreviewEnvVar) AsUnion() ProjectNewResponseDeploymentConfigsPreviewEnvVarsUnion {
	return r.union
}

// A plaintext environment variable.
//
// Union satisfied by
// [ProjectNewResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar] or
// [ProjectNewResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar].
type ProjectNewResponseDeploymentConfigsPreviewEnvVarsUnion interface {
	implementsProjectNewResponseDeploymentConfigsPreviewEnvVar()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectNewResponseDeploymentConfigsPreviewEnvVarsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectNewResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectNewResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar{}),
			DiscriminatorValue: "secret_text",
		},
	)
}

// A plaintext environment variable.
type ProjectNewResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar struct {
	Type ProjectNewResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType `json:"type,required"`
	// Environment variable value.
	Value string                                                                    `json:"value,required"`
	JSON  projectNewResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarJSON
// contains the JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar]
type projectNewResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectNewResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar) implementsProjectNewResponseDeploymentConfigsPreviewEnvVar() {
}

type ProjectNewResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType string

const (
	ProjectNewResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarTypePlainText ProjectNewResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r ProjectNewResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectNewResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type ProjectNewResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar struct {
	Type ProjectNewResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType `json:"type,required"`
	// Secret value.
	Value string                                                                     `json:"value,required"`
	JSON  projectNewResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarJSON
// contains the JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar]
type projectNewResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectNewResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar) implementsProjectNewResponseDeploymentConfigsPreviewEnvVar() {
}

type ProjectNewResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType string

const (
	ProjectNewResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarTypeSecretText ProjectNewResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r ProjectNewResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectNewResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type ProjectNewResponseDeploymentConfigsPreviewEnvVarsType string

const (
	ProjectNewResponseDeploymentConfigsPreviewEnvVarsTypePlainText  ProjectNewResponseDeploymentConfigsPreviewEnvVarsType = "plain_text"
	ProjectNewResponseDeploymentConfigsPreviewEnvVarsTypeSecretText ProjectNewResponseDeploymentConfigsPreviewEnvVarsType = "secret_text"
)

func (r ProjectNewResponseDeploymentConfigsPreviewEnvVarsType) IsKnown() bool {
	switch r {
	case ProjectNewResponseDeploymentConfigsPreviewEnvVarsTypePlainText, ProjectNewResponseDeploymentConfigsPreviewEnvVarsTypeSecretText:
		return true
	}
	return false
}

// The usage model for Pages Functions.
type ProjectNewResponseDeploymentConfigsPreviewUsageModel string

const (
	ProjectNewResponseDeploymentConfigsPreviewUsageModelStandard ProjectNewResponseDeploymentConfigsPreviewUsageModel = "standard"
	ProjectNewResponseDeploymentConfigsPreviewUsageModelBundled  ProjectNewResponseDeploymentConfigsPreviewUsageModel = "bundled"
	ProjectNewResponseDeploymentConfigsPreviewUsageModelUnbound  ProjectNewResponseDeploymentConfigsPreviewUsageModel = "unbound"
)

func (r ProjectNewResponseDeploymentConfigsPreviewUsageModel) IsKnown() bool {
	switch r {
	case ProjectNewResponseDeploymentConfigsPreviewUsageModelStandard, ProjectNewResponseDeploymentConfigsPreviewUsageModelBundled, ProjectNewResponseDeploymentConfigsPreviewUsageModelUnbound:
		return true
	}
	return false
}

// AI binding.
type ProjectNewResponseDeploymentConfigsPreviewAIBinding struct {
	ProjectID string                                                  `json:"project_id,required"`
	JSON      projectNewResponseDeploymentConfigsPreviewAIBindingJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewAIBindingJSON contains the JSON
// metadata for the struct [ProjectNewResponseDeploymentConfigsPreviewAIBinding]
type projectNewResponseDeploymentConfigsPreviewAIBindingJSON struct {
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewAIBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewAIBindingJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine binding.
type ProjectNewResponseDeploymentConfigsPreviewAnalyticsEngineDataset struct {
	// Name of the dataset.
	Dataset string                                                               `json:"dataset,required"`
	JSON    projectNewResponseDeploymentConfigsPreviewAnalyticsEngineDatasetJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewAnalyticsEngineDatasetJSON contains
// the JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsPreviewAnalyticsEngineDataset]
type projectNewResponseDeploymentConfigsPreviewAnalyticsEngineDatasetJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewAnalyticsEngineDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewAnalyticsEngineDatasetJSON) RawJSON() string {
	return r.raw
}

// Browser binding.
type ProjectNewResponseDeploymentConfigsPreviewBrowser struct {
	JSON projectNewResponseDeploymentConfigsPreviewBrowserJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewBrowserJSON contains the JSON metadata
// for the struct [ProjectNewResponseDeploymentConfigsPreviewBrowser]
type projectNewResponseDeploymentConfigsPreviewBrowserJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewBrowser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewBrowserJSON) RawJSON() string {
	return r.raw
}

// D1 binding.
type ProjectNewResponseDeploymentConfigsPreviewD1Database struct {
	// UUID of the D1 database.
	ID   string                                                   `json:"id,required"`
	JSON projectNewResponseDeploymentConfigsPreviewD1DatabaseJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewD1DatabaseJSON contains the JSON
// metadata for the struct [ProjectNewResponseDeploymentConfigsPreviewD1Database]
type projectNewResponseDeploymentConfigsPreviewD1DatabaseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewD1Database) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewD1DatabaseJSON) RawJSON() string {
	return r.raw
}

// Durable Object binding.
type ProjectNewResponseDeploymentConfigsPreviewDurableObjectNamespace struct {
	// ID of the Durable Object namespace.
	NamespaceID string                                                               `json:"namespace_id,required"`
	JSON        projectNewResponseDeploymentConfigsPreviewDurableObjectNamespaceJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewDurableObjectNamespaceJSON contains
// the JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsPreviewDurableObjectNamespace]
type projectNewResponseDeploymentConfigsPreviewDurableObjectNamespaceJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

// Hyperdrive binding.
type ProjectNewResponseDeploymentConfigsPreviewHyperdriveBinding struct {
	ID   string                                                          `json:"id,required"`
	JSON projectNewResponseDeploymentConfigsPreviewHyperdriveBindingJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewHyperdriveBindingJSON contains the
// JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsPreviewHyperdriveBinding]
type projectNewResponseDeploymentConfigsPreviewHyperdriveBindingJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewHyperdriveBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewHyperdriveBindingJSON) RawJSON() string {
	return r.raw
}

// KV namespace binding.
type ProjectNewResponseDeploymentConfigsPreviewKVNamespace struct {
	// ID of the KV namespace.
	NamespaceID string                                                    `json:"namespace_id,required"`
	JSON        projectNewResponseDeploymentConfigsPreviewKVNamespaceJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewKVNamespaceJSON contains the JSON
// metadata for the struct [ProjectNewResponseDeploymentConfigsPreviewKVNamespace]
type projectNewResponseDeploymentConfigsPreviewKVNamespaceJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewKVNamespaceJSON) RawJSON() string {
	return r.raw
}

// Limits for Pages Functions.
type ProjectNewResponseDeploymentConfigsPreviewLimits struct {
	// CPU time limit in milliseconds.
	CPUMs int64                                                `json:"cpu_ms,required"`
	JSON  projectNewResponseDeploymentConfigsPreviewLimitsJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewLimitsJSON contains the JSON metadata
// for the struct [ProjectNewResponseDeploymentConfigsPreviewLimits]
type projectNewResponseDeploymentConfigsPreviewLimitsJSON struct {
	CPUMs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewLimitsJSON) RawJSON() string {
	return r.raw
}

// mTLS binding.
type ProjectNewResponseDeploymentConfigsPreviewMTLSCertificate struct {
	CertificateID string                                                        `json:"certificate_id,required"`
	JSON          projectNewResponseDeploymentConfigsPreviewMTLSCertificateJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewMTLSCertificateJSON contains the JSON
// metadata for the struct
// [ProjectNewResponseDeploymentConfigsPreviewMTLSCertificate]
type projectNewResponseDeploymentConfigsPreviewMTLSCertificateJSON struct {
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewMTLSCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewMTLSCertificateJSON) RawJSON() string {
	return r.raw
}

// Placement setting used for Pages Functions.
type ProjectNewResponseDeploymentConfigsPreviewPlacement struct {
	// Placement mode.
	Mode string                                                  `json:"mode,required"`
	JSON projectNewResponseDeploymentConfigsPreviewPlacementJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewPlacementJSON contains the JSON
// metadata for the struct [ProjectNewResponseDeploymentConfigsPreviewPlacement]
type projectNewResponseDeploymentConfigsPreviewPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewPlacementJSON) RawJSON() string {
	return r.raw
}

// Queue Producer binding.
type ProjectNewResponseDeploymentConfigsPreviewQueueProducer struct {
	// Name of the Queue.
	Name string                                                      `json:"name,required"`
	JSON projectNewResponseDeploymentConfigsPreviewQueueProducerJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewQueueProducerJSON contains the JSON
// metadata for the struct
// [ProjectNewResponseDeploymentConfigsPreviewQueueProducer]
type projectNewResponseDeploymentConfigsPreviewQueueProducerJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewQueueProducer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewQueueProducerJSON) RawJSON() string {
	return r.raw
}

// R2 binding.
type ProjectNewResponseDeploymentConfigsPreviewR2Bucket struct {
	// Name of the R2 bucket.
	Name string `json:"name,required"`
	// Jurisdiction of the R2 bucket.
	Jurisdiction string                                                 `json:"jurisdiction,nullable"`
	JSON         projectNewResponseDeploymentConfigsPreviewR2BucketJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewR2BucketJSON contains the JSON
// metadata for the struct [ProjectNewResponseDeploymentConfigsPreviewR2Bucket]
type projectNewResponseDeploymentConfigsPreviewR2BucketJSON struct {
	Name         apijson.Field
	Jurisdiction apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewR2BucketJSON) RawJSON() string {
	return r.raw
}

// Service binding.
type ProjectNewResponseDeploymentConfigsPreviewService struct {
	// The Service environment.
	Environment string `json:"environment,required"`
	// The Service name.
	Service string `json:"service,required"`
	// The entrypoint to bind to.
	Entrypoint string                                                `json:"entrypoint,nullable"`
	JSON       projectNewResponseDeploymentConfigsPreviewServiceJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewServiceJSON contains the JSON metadata
// for the struct [ProjectNewResponseDeploymentConfigsPreviewService]
type projectNewResponseDeploymentConfigsPreviewServiceJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	Entrypoint  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewServiceJSON) RawJSON() string {
	return r.raw
}

// Vectorize binding.
type ProjectNewResponseDeploymentConfigsPreviewVectorizeBinding struct {
	IndexName string                                                         `json:"index_name,required"`
	JSON      projectNewResponseDeploymentConfigsPreviewVectorizeBindingJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsPreviewVectorizeBindingJSON contains the JSON
// metadata for the struct
// [ProjectNewResponseDeploymentConfigsPreviewVectorizeBinding]
type projectNewResponseDeploymentConfigsPreviewVectorizeBindingJSON struct {
	IndexName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsPreviewVectorizeBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsPreviewVectorizeBindingJSON) RawJSON() string {
	return r.raw
}

// Configs for production deploys.
type ProjectNewResponseDeploymentConfigsProduction struct {
	// Whether to always use the latest compatibility date for Pages Functions.
	AlwaysUseLatestCompatibilityDate bool `json:"always_use_latest_compatibility_date,required"`
	// The major version of the build image to use for Pages Functions.
	BuildImageMajorVersion int64 `json:"build_image_major_version,required"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate string `json:"compatibility_date,required"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags []string `json:"compatibility_flags,required"`
	// Environment variables used for builds and Pages Functions.
	EnvVars map[string]ProjectNewResponseDeploymentConfigsProductionEnvVar `json:"env_vars,required,nullable"`
	// Whether to fail open when the deployment config cannot be applied.
	FailOpen bool `json:"fail_open,required"`
	// The usage model for Pages Functions.
	//
	// Deprecated: All new projects now use the Standard usage model.
	UsageModel ProjectNewResponseDeploymentConfigsProductionUsageModel `json:"usage_model,required"`
	// Constellation bindings used for Pages Functions.
	AIBindings map[string]ProjectNewResponseDeploymentConfigsProductionAIBinding `json:"ai_bindings"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets map[string]ProjectNewResponseDeploymentConfigsProductionAnalyticsEngineDataset `json:"analytics_engine_datasets"`
	// Browser bindings used for Pages Functions.
	Browsers map[string]ProjectNewResponseDeploymentConfigsProductionBrowser `json:"browsers"`
	// D1 databases used for Pages Functions.
	D1Databases map[string]ProjectNewResponseDeploymentConfigsProductionD1Database `json:"d1_databases"`
	// Durable Object namespaces used for Pages Functions.
	DurableObjectNamespaces map[string]ProjectNewResponseDeploymentConfigsProductionDurableObjectNamespace `json:"durable_object_namespaces"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings map[string]ProjectNewResponseDeploymentConfigsProductionHyperdriveBinding `json:"hyperdrive_bindings"`
	// KV namespaces used for Pages Functions.
	KVNamespaces map[string]ProjectNewResponseDeploymentConfigsProductionKVNamespace `json:"kv_namespaces"`
	// Limits for Pages Functions.
	Limits ProjectNewResponseDeploymentConfigsProductionLimits `json:"limits"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates map[string]ProjectNewResponseDeploymentConfigsProductionMTLSCertificate `json:"mtls_certificates"`
	// Placement setting used for Pages Functions.
	Placement ProjectNewResponseDeploymentConfigsProductionPlacement `json:"placement"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers map[string]ProjectNewResponseDeploymentConfigsProductionQueueProducer `json:"queue_producers"`
	// R2 buckets used for Pages Functions.
	R2Buckets map[string]ProjectNewResponseDeploymentConfigsProductionR2Bucket `json:"r2_buckets"`
	// Services used for Pages Functions.
	Services map[string]ProjectNewResponseDeploymentConfigsProductionService `json:"services"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings map[string]ProjectNewResponseDeploymentConfigsProductionVectorizeBinding `json:"vectorize_bindings"`
	// Hash of the Wrangler configuration used for the deployment.
	WranglerConfigHash string                                            `json:"wrangler_config_hash"`
	JSON               projectNewResponseDeploymentConfigsProductionJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionJSON contains the JSON metadata for
// the struct [ProjectNewResponseDeploymentConfigsProduction]
type projectNewResponseDeploymentConfigsProductionJSON struct {
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

func (r *ProjectNewResponseDeploymentConfigsProduction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionJSON) RawJSON() string {
	return r.raw
}

// A plaintext environment variable.
type ProjectNewResponseDeploymentConfigsProductionEnvVar struct {
	Type ProjectNewResponseDeploymentConfigsProductionEnvVarsType `json:"type,required"`
	// Environment variable value.
	Value string                                                  `json:"value,required"`
	JSON  projectNewResponseDeploymentConfigsProductionEnvVarJSON `json:"-"`
	union ProjectNewResponseDeploymentConfigsProductionEnvVarsUnion
}

// projectNewResponseDeploymentConfigsProductionEnvVarJSON contains the JSON
// metadata for the struct [ProjectNewResponseDeploymentConfigsProductionEnvVar]
type projectNewResponseDeploymentConfigsProductionEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectNewResponseDeploymentConfigsProductionEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectNewResponseDeploymentConfigsProductionEnvVar) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectNewResponseDeploymentConfigsProductionEnvVar{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectNewResponseDeploymentConfigsProductionEnvVarsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectNewResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar],
// [ProjectNewResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar].
func (r ProjectNewResponseDeploymentConfigsProductionEnvVar) AsUnion() ProjectNewResponseDeploymentConfigsProductionEnvVarsUnion {
	return r.union
}

// A plaintext environment variable.
//
// Union satisfied by
// [ProjectNewResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar] or
// [ProjectNewResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar].
type ProjectNewResponseDeploymentConfigsProductionEnvVarsUnion interface {
	implementsProjectNewResponseDeploymentConfigsProductionEnvVar()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectNewResponseDeploymentConfigsProductionEnvVarsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectNewResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectNewResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar{}),
			DiscriminatorValue: "secret_text",
		},
	)
}

// A plaintext environment variable.
type ProjectNewResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar struct {
	Type ProjectNewResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType `json:"type,required"`
	// Environment variable value.
	Value string                                                                       `json:"value,required"`
	JSON  projectNewResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarJSON
// contains the JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar]
type projectNewResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectNewResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar) implementsProjectNewResponseDeploymentConfigsProductionEnvVar() {
}

type ProjectNewResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType string

const (
	ProjectNewResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarTypePlainText ProjectNewResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r ProjectNewResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectNewResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type ProjectNewResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar struct {
	Type ProjectNewResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType `json:"type,required"`
	// Secret value.
	Value string                                                                        `json:"value,required"`
	JSON  projectNewResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarJSON
// contains the JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar]
type projectNewResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectNewResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar) implementsProjectNewResponseDeploymentConfigsProductionEnvVar() {
}

type ProjectNewResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType string

const (
	ProjectNewResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarTypeSecretText ProjectNewResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r ProjectNewResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectNewResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type ProjectNewResponseDeploymentConfigsProductionEnvVarsType string

const (
	ProjectNewResponseDeploymentConfigsProductionEnvVarsTypePlainText  ProjectNewResponseDeploymentConfigsProductionEnvVarsType = "plain_text"
	ProjectNewResponseDeploymentConfigsProductionEnvVarsTypeSecretText ProjectNewResponseDeploymentConfigsProductionEnvVarsType = "secret_text"
)

func (r ProjectNewResponseDeploymentConfigsProductionEnvVarsType) IsKnown() bool {
	switch r {
	case ProjectNewResponseDeploymentConfigsProductionEnvVarsTypePlainText, ProjectNewResponseDeploymentConfigsProductionEnvVarsTypeSecretText:
		return true
	}
	return false
}

// The usage model for Pages Functions.
type ProjectNewResponseDeploymentConfigsProductionUsageModel string

const (
	ProjectNewResponseDeploymentConfigsProductionUsageModelStandard ProjectNewResponseDeploymentConfigsProductionUsageModel = "standard"
	ProjectNewResponseDeploymentConfigsProductionUsageModelBundled  ProjectNewResponseDeploymentConfigsProductionUsageModel = "bundled"
	ProjectNewResponseDeploymentConfigsProductionUsageModelUnbound  ProjectNewResponseDeploymentConfigsProductionUsageModel = "unbound"
)

func (r ProjectNewResponseDeploymentConfigsProductionUsageModel) IsKnown() bool {
	switch r {
	case ProjectNewResponseDeploymentConfigsProductionUsageModelStandard, ProjectNewResponseDeploymentConfigsProductionUsageModelBundled, ProjectNewResponseDeploymentConfigsProductionUsageModelUnbound:
		return true
	}
	return false
}

// AI binding.
type ProjectNewResponseDeploymentConfigsProductionAIBinding struct {
	ProjectID string                                                     `json:"project_id,required"`
	JSON      projectNewResponseDeploymentConfigsProductionAIBindingJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionAIBindingJSON contains the JSON
// metadata for the struct [ProjectNewResponseDeploymentConfigsProductionAIBinding]
type projectNewResponseDeploymentConfigsProductionAIBindingJSON struct {
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionAIBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionAIBindingJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine binding.
type ProjectNewResponseDeploymentConfigsProductionAnalyticsEngineDataset struct {
	// Name of the dataset.
	Dataset string                                                                  `json:"dataset,required"`
	JSON    projectNewResponseDeploymentConfigsProductionAnalyticsEngineDatasetJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionAnalyticsEngineDatasetJSON contains
// the JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsProductionAnalyticsEngineDataset]
type projectNewResponseDeploymentConfigsProductionAnalyticsEngineDatasetJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionAnalyticsEngineDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionAnalyticsEngineDatasetJSON) RawJSON() string {
	return r.raw
}

// Browser binding.
type ProjectNewResponseDeploymentConfigsProductionBrowser struct {
	JSON projectNewResponseDeploymentConfigsProductionBrowserJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionBrowserJSON contains the JSON
// metadata for the struct [ProjectNewResponseDeploymentConfigsProductionBrowser]
type projectNewResponseDeploymentConfigsProductionBrowserJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionBrowser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionBrowserJSON) RawJSON() string {
	return r.raw
}

// D1 binding.
type ProjectNewResponseDeploymentConfigsProductionD1Database struct {
	// UUID of the D1 database.
	ID   string                                                      `json:"id,required"`
	JSON projectNewResponseDeploymentConfigsProductionD1DatabaseJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionD1DatabaseJSON contains the JSON
// metadata for the struct
// [ProjectNewResponseDeploymentConfigsProductionD1Database]
type projectNewResponseDeploymentConfigsProductionD1DatabaseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionD1Database) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionD1DatabaseJSON) RawJSON() string {
	return r.raw
}

// Durable Object binding.
type ProjectNewResponseDeploymentConfigsProductionDurableObjectNamespace struct {
	// ID of the Durable Object namespace.
	NamespaceID string                                                                  `json:"namespace_id,required"`
	JSON        projectNewResponseDeploymentConfigsProductionDurableObjectNamespaceJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionDurableObjectNamespaceJSON contains
// the JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsProductionDurableObjectNamespace]
type projectNewResponseDeploymentConfigsProductionDurableObjectNamespaceJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

// Hyperdrive binding.
type ProjectNewResponseDeploymentConfigsProductionHyperdriveBinding struct {
	ID   string                                                             `json:"id,required"`
	JSON projectNewResponseDeploymentConfigsProductionHyperdriveBindingJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionHyperdriveBindingJSON contains the
// JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsProductionHyperdriveBinding]
type projectNewResponseDeploymentConfigsProductionHyperdriveBindingJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionHyperdriveBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionHyperdriveBindingJSON) RawJSON() string {
	return r.raw
}

// KV namespace binding.
type ProjectNewResponseDeploymentConfigsProductionKVNamespace struct {
	// ID of the KV namespace.
	NamespaceID string                                                       `json:"namespace_id,required"`
	JSON        projectNewResponseDeploymentConfigsProductionKVNamespaceJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionKVNamespaceJSON contains the JSON
// metadata for the struct
// [ProjectNewResponseDeploymentConfigsProductionKVNamespace]
type projectNewResponseDeploymentConfigsProductionKVNamespaceJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionKVNamespaceJSON) RawJSON() string {
	return r.raw
}

// Limits for Pages Functions.
type ProjectNewResponseDeploymentConfigsProductionLimits struct {
	// CPU time limit in milliseconds.
	CPUMs int64                                                   `json:"cpu_ms,required"`
	JSON  projectNewResponseDeploymentConfigsProductionLimitsJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionLimitsJSON contains the JSON
// metadata for the struct [ProjectNewResponseDeploymentConfigsProductionLimits]
type projectNewResponseDeploymentConfigsProductionLimitsJSON struct {
	CPUMs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionLimitsJSON) RawJSON() string {
	return r.raw
}

// mTLS binding.
type ProjectNewResponseDeploymentConfigsProductionMTLSCertificate struct {
	CertificateID string                                                           `json:"certificate_id,required"`
	JSON          projectNewResponseDeploymentConfigsProductionMTLSCertificateJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionMTLSCertificateJSON contains the
// JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsProductionMTLSCertificate]
type projectNewResponseDeploymentConfigsProductionMTLSCertificateJSON struct {
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionMTLSCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionMTLSCertificateJSON) RawJSON() string {
	return r.raw
}

// Placement setting used for Pages Functions.
type ProjectNewResponseDeploymentConfigsProductionPlacement struct {
	// Placement mode.
	Mode string                                                     `json:"mode,required"`
	JSON projectNewResponseDeploymentConfigsProductionPlacementJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionPlacementJSON contains the JSON
// metadata for the struct [ProjectNewResponseDeploymentConfigsProductionPlacement]
type projectNewResponseDeploymentConfigsProductionPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionPlacementJSON) RawJSON() string {
	return r.raw
}

// Queue Producer binding.
type ProjectNewResponseDeploymentConfigsProductionQueueProducer struct {
	// Name of the Queue.
	Name string                                                         `json:"name,required"`
	JSON projectNewResponseDeploymentConfigsProductionQueueProducerJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionQueueProducerJSON contains the JSON
// metadata for the struct
// [ProjectNewResponseDeploymentConfigsProductionQueueProducer]
type projectNewResponseDeploymentConfigsProductionQueueProducerJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionQueueProducer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionQueueProducerJSON) RawJSON() string {
	return r.raw
}

// R2 binding.
type ProjectNewResponseDeploymentConfigsProductionR2Bucket struct {
	// Name of the R2 bucket.
	Name string `json:"name,required"`
	// Jurisdiction of the R2 bucket.
	Jurisdiction string                                                    `json:"jurisdiction,nullable"`
	JSON         projectNewResponseDeploymentConfigsProductionR2BucketJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionR2BucketJSON contains the JSON
// metadata for the struct [ProjectNewResponseDeploymentConfigsProductionR2Bucket]
type projectNewResponseDeploymentConfigsProductionR2BucketJSON struct {
	Name         apijson.Field
	Jurisdiction apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionR2BucketJSON) RawJSON() string {
	return r.raw
}

// Service binding.
type ProjectNewResponseDeploymentConfigsProductionService struct {
	// The Service environment.
	Environment string `json:"environment,required"`
	// The Service name.
	Service string `json:"service,required"`
	// The entrypoint to bind to.
	Entrypoint string                                                   `json:"entrypoint,nullable"`
	JSON       projectNewResponseDeploymentConfigsProductionServiceJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionServiceJSON contains the JSON
// metadata for the struct [ProjectNewResponseDeploymentConfigsProductionService]
type projectNewResponseDeploymentConfigsProductionServiceJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	Entrypoint  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionServiceJSON) RawJSON() string {
	return r.raw
}

// Vectorize binding.
type ProjectNewResponseDeploymentConfigsProductionVectorizeBinding struct {
	IndexName string                                                            `json:"index_name,required"`
	JSON      projectNewResponseDeploymentConfigsProductionVectorizeBindingJSON `json:"-"`
}

// projectNewResponseDeploymentConfigsProductionVectorizeBindingJSON contains the
// JSON metadata for the struct
// [ProjectNewResponseDeploymentConfigsProductionVectorizeBinding]
type projectNewResponseDeploymentConfigsProductionVectorizeBindingJSON struct {
	IndexName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseDeploymentConfigsProductionVectorizeBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseDeploymentConfigsProductionVectorizeBindingJSON) RawJSON() string {
	return r.raw
}

// Most recent deployment of the project.
type ProjectNewResponseLatestDeployment struct {
	// Id of the deployment.
	ID string `json:"id,required"`
	// A list of alias URLs pointing to this deployment.
	Aliases []string `json:"aliases,required,nullable"`
	// Configs for the project build process.
	BuildConfig ProjectNewResponseLatestDeploymentBuildConfig `json:"build_config,required"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger ProjectNewResponseLatestDeploymentDeploymentTrigger `json:"deployment_trigger,required"`
	// Environment variables used for builds and Pages Functions.
	EnvVars map[string]ProjectNewResponseLatestDeploymentEnvVar `json:"env_vars,required,nullable"`
	// Type of deploy.
	Environment ProjectNewResponseLatestDeploymentEnvironment `json:"environment,required"`
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
	Source ProjectNewResponseLatestDeploymentSource `json:"source,required"`
	// List of past stages.
	Stages []Stage `json:"stages,required"`
	// The live URL to view this deployment.
	URL string `json:"url,required"`
	// Whether the deployment uses functions.
	UsesFunctions bool                                   `json:"uses_functions,nullable"`
	JSON          projectNewResponseLatestDeploymentJSON `json:"-"`
}

// projectNewResponseLatestDeploymentJSON contains the JSON metadata for the struct
// [ProjectNewResponseLatestDeployment]
type projectNewResponseLatestDeploymentJSON struct {
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

func (r *ProjectNewResponseLatestDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseLatestDeploymentJSON) RawJSON() string {
	return r.raw
}

// Configs for the project build process.
type ProjectNewResponseLatestDeploymentBuildConfig struct {
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
	RootDir string                                            `json:"root_dir,nullable"`
	JSON    projectNewResponseLatestDeploymentBuildConfigJSON `json:"-"`
}

// projectNewResponseLatestDeploymentBuildConfigJSON contains the JSON metadata for
// the struct [ProjectNewResponseLatestDeploymentBuildConfig]
type projectNewResponseLatestDeploymentBuildConfigJSON struct {
	WebAnalyticsTag   apijson.Field
	WebAnalyticsToken apijson.Field
	BuildCaching      apijson.Field
	BuildCommand      apijson.Field
	DestinationDir    apijson.Field
	RootDir           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectNewResponseLatestDeploymentBuildConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseLatestDeploymentBuildConfigJSON) RawJSON() string {
	return r.raw
}

// Info about what caused the deployment.
type ProjectNewResponseLatestDeploymentDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata ProjectNewResponseLatestDeploymentDeploymentTriggerMetadata `json:"metadata,required"`
	// What caused the deployment.
	Type ProjectNewResponseLatestDeploymentDeploymentTriggerType `json:"type,required"`
	JSON projectNewResponseLatestDeploymentDeploymentTriggerJSON `json:"-"`
}

// projectNewResponseLatestDeploymentDeploymentTriggerJSON contains the JSON
// metadata for the struct [ProjectNewResponseLatestDeploymentDeploymentTrigger]
type projectNewResponseLatestDeploymentDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseLatestDeploymentDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseLatestDeploymentDeploymentTriggerJSON) RawJSON() string {
	return r.raw
}

// Additional info about the trigger.
type ProjectNewResponseLatestDeploymentDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch,required"`
	// Whether the deployment trigger commit was dirty.
	CommitDirty bool `json:"commit_dirty,required"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash,required"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                          `json:"commit_message,required"`
	JSON          projectNewResponseLatestDeploymentDeploymentTriggerMetadataJSON `json:"-"`
}

// projectNewResponseLatestDeploymentDeploymentTriggerMetadataJSON contains the
// JSON metadata for the struct
// [ProjectNewResponseLatestDeploymentDeploymentTriggerMetadata]
type projectNewResponseLatestDeploymentDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitDirty   apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectNewResponseLatestDeploymentDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseLatestDeploymentDeploymentTriggerMetadataJSON) RawJSON() string {
	return r.raw
}

// What caused the deployment.
type ProjectNewResponseLatestDeploymentDeploymentTriggerType string

const (
	ProjectNewResponseLatestDeploymentDeploymentTriggerTypeGitHubPush ProjectNewResponseLatestDeploymentDeploymentTriggerType = "github:push"
	ProjectNewResponseLatestDeploymentDeploymentTriggerTypeADHoc      ProjectNewResponseLatestDeploymentDeploymentTriggerType = "ad_hoc"
	ProjectNewResponseLatestDeploymentDeploymentTriggerTypeDeployHook ProjectNewResponseLatestDeploymentDeploymentTriggerType = "deploy_hook"
)

func (r ProjectNewResponseLatestDeploymentDeploymentTriggerType) IsKnown() bool {
	switch r {
	case ProjectNewResponseLatestDeploymentDeploymentTriggerTypeGitHubPush, ProjectNewResponseLatestDeploymentDeploymentTriggerTypeADHoc, ProjectNewResponseLatestDeploymentDeploymentTriggerTypeDeployHook:
		return true
	}
	return false
}

// A plaintext environment variable.
type ProjectNewResponseLatestDeploymentEnvVar struct {
	Type ProjectNewResponseLatestDeploymentEnvVarsType `json:"type,required"`
	// Environment variable value.
	Value string                                       `json:"value,required"`
	JSON  projectNewResponseLatestDeploymentEnvVarJSON `json:"-"`
	union ProjectNewResponseLatestDeploymentEnvVarsUnion
}

// projectNewResponseLatestDeploymentEnvVarJSON contains the JSON metadata for the
// struct [ProjectNewResponseLatestDeploymentEnvVar]
type projectNewResponseLatestDeploymentEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectNewResponseLatestDeploymentEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectNewResponseLatestDeploymentEnvVar) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectNewResponseLatestDeploymentEnvVar{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectNewResponseLatestDeploymentEnvVarsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectNewResponseLatestDeploymentEnvVarsPagesPlainTextEnvVar],
// [ProjectNewResponseLatestDeploymentEnvVarsPagesSecretTextEnvVar].
func (r ProjectNewResponseLatestDeploymentEnvVar) AsUnion() ProjectNewResponseLatestDeploymentEnvVarsUnion {
	return r.union
}

// A plaintext environment variable.
//
// Union satisfied by
// [ProjectNewResponseLatestDeploymentEnvVarsPagesPlainTextEnvVar] or
// [ProjectNewResponseLatestDeploymentEnvVarsPagesSecretTextEnvVar].
type ProjectNewResponseLatestDeploymentEnvVarsUnion interface {
	implementsProjectNewResponseLatestDeploymentEnvVar()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectNewResponseLatestDeploymentEnvVarsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectNewResponseLatestDeploymentEnvVarsPagesPlainTextEnvVar{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectNewResponseLatestDeploymentEnvVarsPagesSecretTextEnvVar{}),
			DiscriminatorValue: "secret_text",
		},
	)
}

// A plaintext environment variable.
type ProjectNewResponseLatestDeploymentEnvVarsPagesPlainTextEnvVar struct {
	Type ProjectNewResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarType `json:"type,required"`
	// Environment variable value.
	Value string                                                            `json:"value,required"`
	JSON  projectNewResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarJSON `json:"-"`
}

// projectNewResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarJSON contains the
// JSON metadata for the struct
// [ProjectNewResponseLatestDeploymentEnvVarsPagesPlainTextEnvVar]
type projectNewResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseLatestDeploymentEnvVarsPagesPlainTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectNewResponseLatestDeploymentEnvVarsPagesPlainTextEnvVar) implementsProjectNewResponseLatestDeploymentEnvVar() {
}

type ProjectNewResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarType string

const (
	ProjectNewResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarTypePlainText ProjectNewResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r ProjectNewResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectNewResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type ProjectNewResponseLatestDeploymentEnvVarsPagesSecretTextEnvVar struct {
	Type ProjectNewResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarType `json:"type,required"`
	// Secret value.
	Value string                                                             `json:"value,required"`
	JSON  projectNewResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarJSON `json:"-"`
}

// projectNewResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarJSON contains the
// JSON metadata for the struct
// [ProjectNewResponseLatestDeploymentEnvVarsPagesSecretTextEnvVar]
type projectNewResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseLatestDeploymentEnvVarsPagesSecretTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectNewResponseLatestDeploymentEnvVarsPagesSecretTextEnvVar) implementsProjectNewResponseLatestDeploymentEnvVar() {
}

type ProjectNewResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarType string

const (
	ProjectNewResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarTypeSecretText ProjectNewResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r ProjectNewResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectNewResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type ProjectNewResponseLatestDeploymentEnvVarsType string

const (
	ProjectNewResponseLatestDeploymentEnvVarsTypePlainText  ProjectNewResponseLatestDeploymentEnvVarsType = "plain_text"
	ProjectNewResponseLatestDeploymentEnvVarsTypeSecretText ProjectNewResponseLatestDeploymentEnvVarsType = "secret_text"
)

func (r ProjectNewResponseLatestDeploymentEnvVarsType) IsKnown() bool {
	switch r {
	case ProjectNewResponseLatestDeploymentEnvVarsTypePlainText, ProjectNewResponseLatestDeploymentEnvVarsTypeSecretText:
		return true
	}
	return false
}

// Type of deploy.
type ProjectNewResponseLatestDeploymentEnvironment string

const (
	ProjectNewResponseLatestDeploymentEnvironmentPreview    ProjectNewResponseLatestDeploymentEnvironment = "preview"
	ProjectNewResponseLatestDeploymentEnvironmentProduction ProjectNewResponseLatestDeploymentEnvironment = "production"
)

func (r ProjectNewResponseLatestDeploymentEnvironment) IsKnown() bool {
	switch r {
	case ProjectNewResponseLatestDeploymentEnvironmentPreview, ProjectNewResponseLatestDeploymentEnvironmentProduction:
		return true
	}
	return false
}

// Configs for the project source control.
type ProjectNewResponseLatestDeploymentSource struct {
	Config ProjectNewResponseLatestDeploymentSourceConfig `json:"config,required"`
	// The source control management provider.
	Type ProjectNewResponseLatestDeploymentSourceType `json:"type,required"`
	JSON projectNewResponseLatestDeploymentSourceJSON `json:"-"`
}

// projectNewResponseLatestDeploymentSourceJSON contains the JSON metadata for the
// struct [ProjectNewResponseLatestDeploymentSource]
type projectNewResponseLatestDeploymentSourceJSON struct {
	Config      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseLatestDeploymentSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseLatestDeploymentSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectNewResponseLatestDeploymentSourceConfig struct {
	// The owner of the repository.
	Owner string `json:"owner,required"`
	// Whether to enable PR comments.
	PrCommentsEnabled bool `json:"pr_comments_enabled,required"`
	// The production branch of the repository.
	ProductionBranch string `json:"production_branch,required"`
	// The name of the repository.
	RepoName string `json:"repo_name,required"`
	// Whether to enable automatic deployments when pushing to the source repository.
	// When disabled, no deployments (production or preview) will be triggered
	// automatically.
	//
	// Deprecated: Use `production_deployments_enabled` and
	// `preview_deployment_setting` for more granular control.
	DeploymentsEnabled bool `json:"deployments_enabled"`
	// The owner ID of the repository.
	OwnerID string `json:"owner_id"`
	// A list of paths that should be excluded from triggering a preview deployment.
	// Wildcard syntax (`*`) is supported.
	PathExcludes []string `json:"path_excludes"`
	// A list of paths that should be watched to trigger a preview deployment. Wildcard
	// syntax (`*`) is supported.
	PathIncludes []string `json:"path_includes"`
	// A list of branches that should not trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchExcludes []string `json:"preview_branch_excludes"`
	// A list of branches that should trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchIncludes []string `json:"preview_branch_includes"`
	// Controls whether commits to preview branches trigger a preview deployment.
	PreviewDeploymentSetting ProjectNewResponseLatestDeploymentSourceConfigPreviewDeploymentSetting `json:"preview_deployment_setting"`
	// Whether to trigger a production deployment on commits to the production branch.
	ProductionDeploymentsEnabled bool `json:"production_deployments_enabled"`
	// The ID of the repository.
	RepoID string                                             `json:"repo_id"`
	JSON   projectNewResponseLatestDeploymentSourceConfigJSON `json:"-"`
}

// projectNewResponseLatestDeploymentSourceConfigJSON contains the JSON metadata
// for the struct [ProjectNewResponseLatestDeploymentSourceConfig]
type projectNewResponseLatestDeploymentSourceConfigJSON struct {
	Owner                        apijson.Field
	PrCommentsEnabled            apijson.Field
	ProductionBranch             apijson.Field
	RepoName                     apijson.Field
	DeploymentsEnabled           apijson.Field
	OwnerID                      apijson.Field
	PathExcludes                 apijson.Field
	PathIncludes                 apijson.Field
	PreviewBranchExcludes        apijson.Field
	PreviewBranchIncludes        apijson.Field
	PreviewDeploymentSetting     apijson.Field
	ProductionDeploymentsEnabled apijson.Field
	RepoID                       apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ProjectNewResponseLatestDeploymentSourceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseLatestDeploymentSourceConfigJSON) RawJSON() string {
	return r.raw
}

// Controls whether commits to preview branches trigger a preview deployment.
type ProjectNewResponseLatestDeploymentSourceConfigPreviewDeploymentSetting string

const (
	ProjectNewResponseLatestDeploymentSourceConfigPreviewDeploymentSettingAll    ProjectNewResponseLatestDeploymentSourceConfigPreviewDeploymentSetting = "all"
	ProjectNewResponseLatestDeploymentSourceConfigPreviewDeploymentSettingNone   ProjectNewResponseLatestDeploymentSourceConfigPreviewDeploymentSetting = "none"
	ProjectNewResponseLatestDeploymentSourceConfigPreviewDeploymentSettingCustom ProjectNewResponseLatestDeploymentSourceConfigPreviewDeploymentSetting = "custom"
)

func (r ProjectNewResponseLatestDeploymentSourceConfigPreviewDeploymentSetting) IsKnown() bool {
	switch r {
	case ProjectNewResponseLatestDeploymentSourceConfigPreviewDeploymentSettingAll, ProjectNewResponseLatestDeploymentSourceConfigPreviewDeploymentSettingNone, ProjectNewResponseLatestDeploymentSourceConfigPreviewDeploymentSettingCustom:
		return true
	}
	return false
}

// The source control management provider.
type ProjectNewResponseLatestDeploymentSourceType string

const (
	ProjectNewResponseLatestDeploymentSourceTypeGitHub ProjectNewResponseLatestDeploymentSourceType = "github"
	ProjectNewResponseLatestDeploymentSourceTypeGitlab ProjectNewResponseLatestDeploymentSourceType = "gitlab"
)

func (r ProjectNewResponseLatestDeploymentSourceType) IsKnown() bool {
	switch r {
	case ProjectNewResponseLatestDeploymentSourceTypeGitHub, ProjectNewResponseLatestDeploymentSourceTypeGitlab:
		return true
	}
	return false
}

// Configs for the project build process.
type ProjectNewResponseBuildConfig struct {
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
	RootDir string                            `json:"root_dir,nullable"`
	JSON    projectNewResponseBuildConfigJSON `json:"-"`
}

// projectNewResponseBuildConfigJSON contains the JSON metadata for the struct
// [ProjectNewResponseBuildConfig]
type projectNewResponseBuildConfigJSON struct {
	WebAnalyticsTag   apijson.Field
	WebAnalyticsToken apijson.Field
	BuildCaching      apijson.Field
	BuildCommand      apijson.Field
	DestinationDir    apijson.Field
	RootDir           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectNewResponseBuildConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseBuildConfigJSON) RawJSON() string {
	return r.raw
}

// Configs for the project source control.
type ProjectNewResponseSource struct {
	Config ProjectNewResponseSourceConfig `json:"config,required"`
	// The source control management provider.
	Type ProjectNewResponseSourceType `json:"type,required"`
	JSON projectNewResponseSourceJSON `json:"-"`
}

// projectNewResponseSourceJSON contains the JSON metadata for the struct
// [ProjectNewResponseSource]
type projectNewResponseSourceJSON struct {
	Config      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectNewResponseSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectNewResponseSourceConfig struct {
	// The owner of the repository.
	Owner string `json:"owner,required"`
	// Whether to enable PR comments.
	PrCommentsEnabled bool `json:"pr_comments_enabled,required"`
	// The production branch of the repository.
	ProductionBranch string `json:"production_branch,required"`
	// The name of the repository.
	RepoName string `json:"repo_name,required"`
	// Whether to enable automatic deployments when pushing to the source repository.
	// When disabled, no deployments (production or preview) will be triggered
	// automatically.
	//
	// Deprecated: Use `production_deployments_enabled` and
	// `preview_deployment_setting` for more granular control.
	DeploymentsEnabled bool `json:"deployments_enabled"`
	// The owner ID of the repository.
	OwnerID string `json:"owner_id"`
	// A list of paths that should be excluded from triggering a preview deployment.
	// Wildcard syntax (`*`) is supported.
	PathExcludes []string `json:"path_excludes"`
	// A list of paths that should be watched to trigger a preview deployment. Wildcard
	// syntax (`*`) is supported.
	PathIncludes []string `json:"path_includes"`
	// A list of branches that should not trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchExcludes []string `json:"preview_branch_excludes"`
	// A list of branches that should trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchIncludes []string `json:"preview_branch_includes"`
	// Controls whether commits to preview branches trigger a preview deployment.
	PreviewDeploymentSetting ProjectNewResponseSourceConfigPreviewDeploymentSetting `json:"preview_deployment_setting"`
	// Whether to trigger a production deployment on commits to the production branch.
	ProductionDeploymentsEnabled bool `json:"production_deployments_enabled"`
	// The ID of the repository.
	RepoID string                             `json:"repo_id"`
	JSON   projectNewResponseSourceConfigJSON `json:"-"`
}

// projectNewResponseSourceConfigJSON contains the JSON metadata for the struct
// [ProjectNewResponseSourceConfig]
type projectNewResponseSourceConfigJSON struct {
	Owner                        apijson.Field
	PrCommentsEnabled            apijson.Field
	ProductionBranch             apijson.Field
	RepoName                     apijson.Field
	DeploymentsEnabled           apijson.Field
	OwnerID                      apijson.Field
	PathExcludes                 apijson.Field
	PathIncludes                 apijson.Field
	PreviewBranchExcludes        apijson.Field
	PreviewBranchIncludes        apijson.Field
	PreviewDeploymentSetting     apijson.Field
	ProductionDeploymentsEnabled apijson.Field
	RepoID                       apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ProjectNewResponseSourceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectNewResponseSourceConfigJSON) RawJSON() string {
	return r.raw
}

// Controls whether commits to preview branches trigger a preview deployment.
type ProjectNewResponseSourceConfigPreviewDeploymentSetting string

const (
	ProjectNewResponseSourceConfigPreviewDeploymentSettingAll    ProjectNewResponseSourceConfigPreviewDeploymentSetting = "all"
	ProjectNewResponseSourceConfigPreviewDeploymentSettingNone   ProjectNewResponseSourceConfigPreviewDeploymentSetting = "none"
	ProjectNewResponseSourceConfigPreviewDeploymentSettingCustom ProjectNewResponseSourceConfigPreviewDeploymentSetting = "custom"
)

func (r ProjectNewResponseSourceConfigPreviewDeploymentSetting) IsKnown() bool {
	switch r {
	case ProjectNewResponseSourceConfigPreviewDeploymentSettingAll, ProjectNewResponseSourceConfigPreviewDeploymentSettingNone, ProjectNewResponseSourceConfigPreviewDeploymentSettingCustom:
		return true
	}
	return false
}

// The source control management provider.
type ProjectNewResponseSourceType string

const (
	ProjectNewResponseSourceTypeGitHub ProjectNewResponseSourceType = "github"
	ProjectNewResponseSourceTypeGitlab ProjectNewResponseSourceType = "gitlab"
)

func (r ProjectNewResponseSourceType) IsKnown() bool {
	switch r {
	case ProjectNewResponseSourceTypeGitHub, ProjectNewResponseSourceTypeGitlab:
		return true
	}
	return false
}

type ProjectListResponse struct {
	// ID of the project.
	ID string `json:"id,required"`
	// Most recent production deployment of the project.
	CanonicalDeployment ProjectListResponseCanonicalDeployment `json:"canonical_deployment,required,nullable"`
	// When the project was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Configs for deployments in a project.
	DeploymentConfigs ProjectListResponseDeploymentConfigs `json:"deployment_configs,required"`
	// Framework the project is using.
	Framework string `json:"framework,required"`
	// Version of the framework the project is using.
	FrameworkVersion string `json:"framework_version,required"`
	// Most recent deployment of the project.
	LatestDeployment ProjectListResponseLatestDeployment `json:"latest_deployment,required,nullable"`
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
	BuildConfig ProjectListResponseBuildConfig `json:"build_config"`
	// A list of associated custom domains for the project.
	Domains []string `json:"domains"`
	// Configs for the project source control.
	Source ProjectListResponseSource `json:"source"`
	// The Cloudflare subdomain associated with the project.
	Subdomain string                  `json:"subdomain"`
	JSON      projectListResponseJSON `json:"-"`
}

// projectListResponseJSON contains the JSON metadata for the struct
// [ProjectListResponse]
type projectListResponseJSON struct {
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

func (r *ProjectListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseJSON) RawJSON() string {
	return r.raw
}

// Most recent production deployment of the project.
type ProjectListResponseCanonicalDeployment struct {
	// Id of the deployment.
	ID string `json:"id,required"`
	// A list of alias URLs pointing to this deployment.
	Aliases []string `json:"aliases,required,nullable"`
	// Configs for the project build process.
	BuildConfig ProjectListResponseCanonicalDeploymentBuildConfig `json:"build_config,required"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger ProjectListResponseCanonicalDeploymentDeploymentTrigger `json:"deployment_trigger,required"`
	// Environment variables used for builds and Pages Functions.
	EnvVars map[string]ProjectListResponseCanonicalDeploymentEnvVar `json:"env_vars,required,nullable"`
	// Type of deploy.
	Environment ProjectListResponseCanonicalDeploymentEnvironment `json:"environment,required"`
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
	Source ProjectListResponseCanonicalDeploymentSource `json:"source,required"`
	// List of past stages.
	Stages []Stage `json:"stages,required"`
	// The live URL to view this deployment.
	URL string `json:"url,required"`
	// Whether the deployment uses functions.
	UsesFunctions bool                                       `json:"uses_functions,nullable"`
	JSON          projectListResponseCanonicalDeploymentJSON `json:"-"`
}

// projectListResponseCanonicalDeploymentJSON contains the JSON metadata for the
// struct [ProjectListResponseCanonicalDeployment]
type projectListResponseCanonicalDeploymentJSON struct {
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

func (r *ProjectListResponseCanonicalDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseCanonicalDeploymentJSON) RawJSON() string {
	return r.raw
}

// Configs for the project build process.
type ProjectListResponseCanonicalDeploymentBuildConfig struct {
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
	RootDir string                                                `json:"root_dir,nullable"`
	JSON    projectListResponseCanonicalDeploymentBuildConfigJSON `json:"-"`
}

// projectListResponseCanonicalDeploymentBuildConfigJSON contains the JSON metadata
// for the struct [ProjectListResponseCanonicalDeploymentBuildConfig]
type projectListResponseCanonicalDeploymentBuildConfigJSON struct {
	WebAnalyticsTag   apijson.Field
	WebAnalyticsToken apijson.Field
	BuildCaching      apijson.Field
	BuildCommand      apijson.Field
	DestinationDir    apijson.Field
	RootDir           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectListResponseCanonicalDeploymentBuildConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseCanonicalDeploymentBuildConfigJSON) RawJSON() string {
	return r.raw
}

// Info about what caused the deployment.
type ProjectListResponseCanonicalDeploymentDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata ProjectListResponseCanonicalDeploymentDeploymentTriggerMetadata `json:"metadata,required"`
	// What caused the deployment.
	Type ProjectListResponseCanonicalDeploymentDeploymentTriggerType `json:"type,required"`
	JSON projectListResponseCanonicalDeploymentDeploymentTriggerJSON `json:"-"`
}

// projectListResponseCanonicalDeploymentDeploymentTriggerJSON contains the JSON
// metadata for the struct
// [ProjectListResponseCanonicalDeploymentDeploymentTrigger]
type projectListResponseCanonicalDeploymentDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseCanonicalDeploymentDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseCanonicalDeploymentDeploymentTriggerJSON) RawJSON() string {
	return r.raw
}

// Additional info about the trigger.
type ProjectListResponseCanonicalDeploymentDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch,required"`
	// Whether the deployment trigger commit was dirty.
	CommitDirty bool `json:"commit_dirty,required"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash,required"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                              `json:"commit_message,required"`
	JSON          projectListResponseCanonicalDeploymentDeploymentTriggerMetadataJSON `json:"-"`
}

// projectListResponseCanonicalDeploymentDeploymentTriggerMetadataJSON contains the
// JSON metadata for the struct
// [ProjectListResponseCanonicalDeploymentDeploymentTriggerMetadata]
type projectListResponseCanonicalDeploymentDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitDirty   apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectListResponseCanonicalDeploymentDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseCanonicalDeploymentDeploymentTriggerMetadataJSON) RawJSON() string {
	return r.raw
}

// What caused the deployment.
type ProjectListResponseCanonicalDeploymentDeploymentTriggerType string

const (
	ProjectListResponseCanonicalDeploymentDeploymentTriggerTypeGitHubPush ProjectListResponseCanonicalDeploymentDeploymentTriggerType = "github:push"
	ProjectListResponseCanonicalDeploymentDeploymentTriggerTypeADHoc      ProjectListResponseCanonicalDeploymentDeploymentTriggerType = "ad_hoc"
	ProjectListResponseCanonicalDeploymentDeploymentTriggerTypeDeployHook ProjectListResponseCanonicalDeploymentDeploymentTriggerType = "deploy_hook"
)

func (r ProjectListResponseCanonicalDeploymentDeploymentTriggerType) IsKnown() bool {
	switch r {
	case ProjectListResponseCanonicalDeploymentDeploymentTriggerTypeGitHubPush, ProjectListResponseCanonicalDeploymentDeploymentTriggerTypeADHoc, ProjectListResponseCanonicalDeploymentDeploymentTriggerTypeDeployHook:
		return true
	}
	return false
}

// A plaintext environment variable.
type ProjectListResponseCanonicalDeploymentEnvVar struct {
	Type ProjectListResponseCanonicalDeploymentEnvVarsType `json:"type,required"`
	// Environment variable value.
	Value string                                           `json:"value,required"`
	JSON  projectListResponseCanonicalDeploymentEnvVarJSON `json:"-"`
	union ProjectListResponseCanonicalDeploymentEnvVarsUnion
}

// projectListResponseCanonicalDeploymentEnvVarJSON contains the JSON metadata for
// the struct [ProjectListResponseCanonicalDeploymentEnvVar]
type projectListResponseCanonicalDeploymentEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectListResponseCanonicalDeploymentEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectListResponseCanonicalDeploymentEnvVar) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectListResponseCanonicalDeploymentEnvVar{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectListResponseCanonicalDeploymentEnvVarsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectListResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVar],
// [ProjectListResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVar].
func (r ProjectListResponseCanonicalDeploymentEnvVar) AsUnion() ProjectListResponseCanonicalDeploymentEnvVarsUnion {
	return r.union
}

// A plaintext environment variable.
//
// Union satisfied by
// [ProjectListResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVar] or
// [ProjectListResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVar].
type ProjectListResponseCanonicalDeploymentEnvVarsUnion interface {
	implementsProjectListResponseCanonicalDeploymentEnvVar()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectListResponseCanonicalDeploymentEnvVarsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectListResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVar{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectListResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVar{}),
			DiscriminatorValue: "secret_text",
		},
	)
}

// A plaintext environment variable.
type ProjectListResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVar struct {
	Type ProjectListResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarType `json:"type,required"`
	// Environment variable value.
	Value string                                                                `json:"value,required"`
	JSON  projectListResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarJSON `json:"-"`
}

// projectListResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarJSON contains
// the JSON metadata for the struct
// [ProjectListResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVar]
type projectListResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectListResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVar) implementsProjectListResponseCanonicalDeploymentEnvVar() {
}

type ProjectListResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarType string

const (
	ProjectListResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarTypePlainText ProjectListResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r ProjectListResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectListResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type ProjectListResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVar struct {
	Type ProjectListResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarType `json:"type,required"`
	// Secret value.
	Value string                                                                 `json:"value,required"`
	JSON  projectListResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarJSON `json:"-"`
}

// projectListResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarJSON contains
// the JSON metadata for the struct
// [ProjectListResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVar]
type projectListResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectListResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVar) implementsProjectListResponseCanonicalDeploymentEnvVar() {
}

type ProjectListResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarType string

const (
	ProjectListResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarTypeSecretText ProjectListResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r ProjectListResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectListResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type ProjectListResponseCanonicalDeploymentEnvVarsType string

const (
	ProjectListResponseCanonicalDeploymentEnvVarsTypePlainText  ProjectListResponseCanonicalDeploymentEnvVarsType = "plain_text"
	ProjectListResponseCanonicalDeploymentEnvVarsTypeSecretText ProjectListResponseCanonicalDeploymentEnvVarsType = "secret_text"
)

func (r ProjectListResponseCanonicalDeploymentEnvVarsType) IsKnown() bool {
	switch r {
	case ProjectListResponseCanonicalDeploymentEnvVarsTypePlainText, ProjectListResponseCanonicalDeploymentEnvVarsTypeSecretText:
		return true
	}
	return false
}

// Type of deploy.
type ProjectListResponseCanonicalDeploymentEnvironment string

const (
	ProjectListResponseCanonicalDeploymentEnvironmentPreview    ProjectListResponseCanonicalDeploymentEnvironment = "preview"
	ProjectListResponseCanonicalDeploymentEnvironmentProduction ProjectListResponseCanonicalDeploymentEnvironment = "production"
)

func (r ProjectListResponseCanonicalDeploymentEnvironment) IsKnown() bool {
	switch r {
	case ProjectListResponseCanonicalDeploymentEnvironmentPreview, ProjectListResponseCanonicalDeploymentEnvironmentProduction:
		return true
	}
	return false
}

// Configs for the project source control.
type ProjectListResponseCanonicalDeploymentSource struct {
	Config ProjectListResponseCanonicalDeploymentSourceConfig `json:"config,required"`
	// The source control management provider.
	Type ProjectListResponseCanonicalDeploymentSourceType `json:"type,required"`
	JSON projectListResponseCanonicalDeploymentSourceJSON `json:"-"`
}

// projectListResponseCanonicalDeploymentSourceJSON contains the JSON metadata for
// the struct [ProjectListResponseCanonicalDeploymentSource]
type projectListResponseCanonicalDeploymentSourceJSON struct {
	Config      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseCanonicalDeploymentSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseCanonicalDeploymentSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectListResponseCanonicalDeploymentSourceConfig struct {
	// The owner of the repository.
	Owner string `json:"owner,required"`
	// Whether to enable PR comments.
	PrCommentsEnabled bool `json:"pr_comments_enabled,required"`
	// The production branch of the repository.
	ProductionBranch string `json:"production_branch,required"`
	// The name of the repository.
	RepoName string `json:"repo_name,required"`
	// Whether to enable automatic deployments when pushing to the source repository.
	// When disabled, no deployments (production or preview) will be triggered
	// automatically.
	//
	// Deprecated: Use `production_deployments_enabled` and
	// `preview_deployment_setting` for more granular control.
	DeploymentsEnabled bool `json:"deployments_enabled"`
	// The owner ID of the repository.
	OwnerID string `json:"owner_id"`
	// A list of paths that should be excluded from triggering a preview deployment.
	// Wildcard syntax (`*`) is supported.
	PathExcludes []string `json:"path_excludes"`
	// A list of paths that should be watched to trigger a preview deployment. Wildcard
	// syntax (`*`) is supported.
	PathIncludes []string `json:"path_includes"`
	// A list of branches that should not trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchExcludes []string `json:"preview_branch_excludes"`
	// A list of branches that should trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchIncludes []string `json:"preview_branch_includes"`
	// Controls whether commits to preview branches trigger a preview deployment.
	PreviewDeploymentSetting ProjectListResponseCanonicalDeploymentSourceConfigPreviewDeploymentSetting `json:"preview_deployment_setting"`
	// Whether to trigger a production deployment on commits to the production branch.
	ProductionDeploymentsEnabled bool `json:"production_deployments_enabled"`
	// The ID of the repository.
	RepoID string                                                 `json:"repo_id"`
	JSON   projectListResponseCanonicalDeploymentSourceConfigJSON `json:"-"`
}

// projectListResponseCanonicalDeploymentSourceConfigJSON contains the JSON
// metadata for the struct [ProjectListResponseCanonicalDeploymentSourceConfig]
type projectListResponseCanonicalDeploymentSourceConfigJSON struct {
	Owner                        apijson.Field
	PrCommentsEnabled            apijson.Field
	ProductionBranch             apijson.Field
	RepoName                     apijson.Field
	DeploymentsEnabled           apijson.Field
	OwnerID                      apijson.Field
	PathExcludes                 apijson.Field
	PathIncludes                 apijson.Field
	PreviewBranchExcludes        apijson.Field
	PreviewBranchIncludes        apijson.Field
	PreviewDeploymentSetting     apijson.Field
	ProductionDeploymentsEnabled apijson.Field
	RepoID                       apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ProjectListResponseCanonicalDeploymentSourceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseCanonicalDeploymentSourceConfigJSON) RawJSON() string {
	return r.raw
}

// Controls whether commits to preview branches trigger a preview deployment.
type ProjectListResponseCanonicalDeploymentSourceConfigPreviewDeploymentSetting string

const (
	ProjectListResponseCanonicalDeploymentSourceConfigPreviewDeploymentSettingAll    ProjectListResponseCanonicalDeploymentSourceConfigPreviewDeploymentSetting = "all"
	ProjectListResponseCanonicalDeploymentSourceConfigPreviewDeploymentSettingNone   ProjectListResponseCanonicalDeploymentSourceConfigPreviewDeploymentSetting = "none"
	ProjectListResponseCanonicalDeploymentSourceConfigPreviewDeploymentSettingCustom ProjectListResponseCanonicalDeploymentSourceConfigPreviewDeploymentSetting = "custom"
)

func (r ProjectListResponseCanonicalDeploymentSourceConfigPreviewDeploymentSetting) IsKnown() bool {
	switch r {
	case ProjectListResponseCanonicalDeploymentSourceConfigPreviewDeploymentSettingAll, ProjectListResponseCanonicalDeploymentSourceConfigPreviewDeploymentSettingNone, ProjectListResponseCanonicalDeploymentSourceConfigPreviewDeploymentSettingCustom:
		return true
	}
	return false
}

// The source control management provider.
type ProjectListResponseCanonicalDeploymentSourceType string

const (
	ProjectListResponseCanonicalDeploymentSourceTypeGitHub ProjectListResponseCanonicalDeploymentSourceType = "github"
	ProjectListResponseCanonicalDeploymentSourceTypeGitlab ProjectListResponseCanonicalDeploymentSourceType = "gitlab"
)

func (r ProjectListResponseCanonicalDeploymentSourceType) IsKnown() bool {
	switch r {
	case ProjectListResponseCanonicalDeploymentSourceTypeGitHub, ProjectListResponseCanonicalDeploymentSourceTypeGitlab:
		return true
	}
	return false
}

// Configs for deployments in a project.
type ProjectListResponseDeploymentConfigs struct {
	// Configs for preview deploys.
	Preview ProjectListResponseDeploymentConfigsPreview `json:"preview,required"`
	// Configs for production deploys.
	Production ProjectListResponseDeploymentConfigsProduction `json:"production,required"`
	JSON       projectListResponseDeploymentConfigsJSON       `json:"-"`
}

// projectListResponseDeploymentConfigsJSON contains the JSON metadata for the
// struct [ProjectListResponseDeploymentConfigs]
type projectListResponseDeploymentConfigsJSON struct {
	Preview     apijson.Field
	Production  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsJSON) RawJSON() string {
	return r.raw
}

// Configs for preview deploys.
type ProjectListResponseDeploymentConfigsPreview struct {
	// Whether to always use the latest compatibility date for Pages Functions.
	AlwaysUseLatestCompatibilityDate bool `json:"always_use_latest_compatibility_date,required"`
	// The major version of the build image to use for Pages Functions.
	BuildImageMajorVersion int64 `json:"build_image_major_version,required"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate string `json:"compatibility_date,required"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags []string `json:"compatibility_flags,required"`
	// Environment variables used for builds and Pages Functions.
	EnvVars map[string]ProjectListResponseDeploymentConfigsPreviewEnvVar `json:"env_vars,required,nullable"`
	// Whether to fail open when the deployment config cannot be applied.
	FailOpen bool `json:"fail_open,required"`
	// The usage model for Pages Functions.
	//
	// Deprecated: All new projects now use the Standard usage model.
	UsageModel ProjectListResponseDeploymentConfigsPreviewUsageModel `json:"usage_model,required"`
	// Constellation bindings used for Pages Functions.
	AIBindings map[string]ProjectListResponseDeploymentConfigsPreviewAIBinding `json:"ai_bindings"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets map[string]ProjectListResponseDeploymentConfigsPreviewAnalyticsEngineDataset `json:"analytics_engine_datasets"`
	// Browser bindings used for Pages Functions.
	Browsers map[string]ProjectListResponseDeploymentConfigsPreviewBrowser `json:"browsers"`
	// D1 databases used for Pages Functions.
	D1Databases map[string]ProjectListResponseDeploymentConfigsPreviewD1Database `json:"d1_databases"`
	// Durable Object namespaces used for Pages Functions.
	DurableObjectNamespaces map[string]ProjectListResponseDeploymentConfigsPreviewDurableObjectNamespace `json:"durable_object_namespaces"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings map[string]ProjectListResponseDeploymentConfigsPreviewHyperdriveBinding `json:"hyperdrive_bindings"`
	// KV namespaces used for Pages Functions.
	KVNamespaces map[string]ProjectListResponseDeploymentConfigsPreviewKVNamespace `json:"kv_namespaces"`
	// Limits for Pages Functions.
	Limits ProjectListResponseDeploymentConfigsPreviewLimits `json:"limits"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates map[string]ProjectListResponseDeploymentConfigsPreviewMTLSCertificate `json:"mtls_certificates"`
	// Placement setting used for Pages Functions.
	Placement ProjectListResponseDeploymentConfigsPreviewPlacement `json:"placement"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers map[string]ProjectListResponseDeploymentConfigsPreviewQueueProducer `json:"queue_producers"`
	// R2 buckets used for Pages Functions.
	R2Buckets map[string]ProjectListResponseDeploymentConfigsPreviewR2Bucket `json:"r2_buckets"`
	// Services used for Pages Functions.
	Services map[string]ProjectListResponseDeploymentConfigsPreviewService `json:"services"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings map[string]ProjectListResponseDeploymentConfigsPreviewVectorizeBinding `json:"vectorize_bindings"`
	// Hash of the Wrangler configuration used for the deployment.
	WranglerConfigHash string                                          `json:"wrangler_config_hash"`
	JSON               projectListResponseDeploymentConfigsPreviewJSON `json:"-"`
}

// projectListResponseDeploymentConfigsPreviewJSON contains the JSON metadata for
// the struct [ProjectListResponseDeploymentConfigsPreview]
type projectListResponseDeploymentConfigsPreviewJSON struct {
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

func (r *ProjectListResponseDeploymentConfigsPreview) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsPreviewJSON) RawJSON() string {
	return r.raw
}

// A plaintext environment variable.
type ProjectListResponseDeploymentConfigsPreviewEnvVar struct {
	Type ProjectListResponseDeploymentConfigsPreviewEnvVarsType `json:"type,required"`
	// Environment variable value.
	Value string                                                `json:"value,required"`
	JSON  projectListResponseDeploymentConfigsPreviewEnvVarJSON `json:"-"`
	union ProjectListResponseDeploymentConfigsPreviewEnvVarsUnion
}

// projectListResponseDeploymentConfigsPreviewEnvVarJSON contains the JSON metadata
// for the struct [ProjectListResponseDeploymentConfigsPreviewEnvVar]
type projectListResponseDeploymentConfigsPreviewEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectListResponseDeploymentConfigsPreviewEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectListResponseDeploymentConfigsPreviewEnvVar) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectListResponseDeploymentConfigsPreviewEnvVar{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectListResponseDeploymentConfigsPreviewEnvVarsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectListResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar],
// [ProjectListResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar].
func (r ProjectListResponseDeploymentConfigsPreviewEnvVar) AsUnion() ProjectListResponseDeploymentConfigsPreviewEnvVarsUnion {
	return r.union
}

// A plaintext environment variable.
//
// Union satisfied by
// [ProjectListResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar] or
// [ProjectListResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar].
type ProjectListResponseDeploymentConfigsPreviewEnvVarsUnion interface {
	implementsProjectListResponseDeploymentConfigsPreviewEnvVar()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectListResponseDeploymentConfigsPreviewEnvVarsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectListResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectListResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar{}),
			DiscriminatorValue: "secret_text",
		},
	)
}

// A plaintext environment variable.
type ProjectListResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar struct {
	Type ProjectListResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType `json:"type,required"`
	// Environment variable value.
	Value string                                                                     `json:"value,required"`
	JSON  projectListResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarJSON `json:"-"`
}

// projectListResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarJSON
// contains the JSON metadata for the struct
// [ProjectListResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar]
type projectListResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectListResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar) implementsProjectListResponseDeploymentConfigsPreviewEnvVar() {
}

type ProjectListResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType string

const (
	ProjectListResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarTypePlainText ProjectListResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r ProjectListResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectListResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type ProjectListResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar struct {
	Type ProjectListResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType `json:"type,required"`
	// Secret value.
	Value string                                                                      `json:"value,required"`
	JSON  projectListResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarJSON `json:"-"`
}

// projectListResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarJSON
// contains the JSON metadata for the struct
// [ProjectListResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar]
type projectListResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectListResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar) implementsProjectListResponseDeploymentConfigsPreviewEnvVar() {
}

type ProjectListResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType string

const (
	ProjectListResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarTypeSecretText ProjectListResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r ProjectListResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectListResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type ProjectListResponseDeploymentConfigsPreviewEnvVarsType string

const (
	ProjectListResponseDeploymentConfigsPreviewEnvVarsTypePlainText  ProjectListResponseDeploymentConfigsPreviewEnvVarsType = "plain_text"
	ProjectListResponseDeploymentConfigsPreviewEnvVarsTypeSecretText ProjectListResponseDeploymentConfigsPreviewEnvVarsType = "secret_text"
)

func (r ProjectListResponseDeploymentConfigsPreviewEnvVarsType) IsKnown() bool {
	switch r {
	case ProjectListResponseDeploymentConfigsPreviewEnvVarsTypePlainText, ProjectListResponseDeploymentConfigsPreviewEnvVarsTypeSecretText:
		return true
	}
	return false
}

// The usage model for Pages Functions.
type ProjectListResponseDeploymentConfigsPreviewUsageModel string

const (
	ProjectListResponseDeploymentConfigsPreviewUsageModelStandard ProjectListResponseDeploymentConfigsPreviewUsageModel = "standard"
	ProjectListResponseDeploymentConfigsPreviewUsageModelBundled  ProjectListResponseDeploymentConfigsPreviewUsageModel = "bundled"
	ProjectListResponseDeploymentConfigsPreviewUsageModelUnbound  ProjectListResponseDeploymentConfigsPreviewUsageModel = "unbound"
)

func (r ProjectListResponseDeploymentConfigsPreviewUsageModel) IsKnown() bool {
	switch r {
	case ProjectListResponseDeploymentConfigsPreviewUsageModelStandard, ProjectListResponseDeploymentConfigsPreviewUsageModelBundled, ProjectListResponseDeploymentConfigsPreviewUsageModelUnbound:
		return true
	}
	return false
}

// AI binding.
type ProjectListResponseDeploymentConfigsPreviewAIBinding struct {
	ProjectID string                                                   `json:"project_id,required"`
	JSON      projectListResponseDeploymentConfigsPreviewAIBindingJSON `json:"-"`
}

// projectListResponseDeploymentConfigsPreviewAIBindingJSON contains the JSON
// metadata for the struct [ProjectListResponseDeploymentConfigsPreviewAIBinding]
type projectListResponseDeploymentConfigsPreviewAIBindingJSON struct {
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsPreviewAIBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsPreviewAIBindingJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine binding.
type ProjectListResponseDeploymentConfigsPreviewAnalyticsEngineDataset struct {
	// Name of the dataset.
	Dataset string                                                                `json:"dataset,required"`
	JSON    projectListResponseDeploymentConfigsPreviewAnalyticsEngineDatasetJSON `json:"-"`
}

// projectListResponseDeploymentConfigsPreviewAnalyticsEngineDatasetJSON contains
// the JSON metadata for the struct
// [ProjectListResponseDeploymentConfigsPreviewAnalyticsEngineDataset]
type projectListResponseDeploymentConfigsPreviewAnalyticsEngineDatasetJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsPreviewAnalyticsEngineDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsPreviewAnalyticsEngineDatasetJSON) RawJSON() string {
	return r.raw
}

// Browser binding.
type ProjectListResponseDeploymentConfigsPreviewBrowser struct {
	JSON projectListResponseDeploymentConfigsPreviewBrowserJSON `json:"-"`
}

// projectListResponseDeploymentConfigsPreviewBrowserJSON contains the JSON
// metadata for the struct [ProjectListResponseDeploymentConfigsPreviewBrowser]
type projectListResponseDeploymentConfigsPreviewBrowserJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsPreviewBrowser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsPreviewBrowserJSON) RawJSON() string {
	return r.raw
}

// D1 binding.
type ProjectListResponseDeploymentConfigsPreviewD1Database struct {
	// UUID of the D1 database.
	ID   string                                                    `json:"id,required"`
	JSON projectListResponseDeploymentConfigsPreviewD1DatabaseJSON `json:"-"`
}

// projectListResponseDeploymentConfigsPreviewD1DatabaseJSON contains the JSON
// metadata for the struct [ProjectListResponseDeploymentConfigsPreviewD1Database]
type projectListResponseDeploymentConfigsPreviewD1DatabaseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsPreviewD1Database) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsPreviewD1DatabaseJSON) RawJSON() string {
	return r.raw
}

// Durable Object binding.
type ProjectListResponseDeploymentConfigsPreviewDurableObjectNamespace struct {
	// ID of the Durable Object namespace.
	NamespaceID string                                                                `json:"namespace_id,required"`
	JSON        projectListResponseDeploymentConfigsPreviewDurableObjectNamespaceJSON `json:"-"`
}

// projectListResponseDeploymentConfigsPreviewDurableObjectNamespaceJSON contains
// the JSON metadata for the struct
// [ProjectListResponseDeploymentConfigsPreviewDurableObjectNamespace]
type projectListResponseDeploymentConfigsPreviewDurableObjectNamespaceJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsPreviewDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsPreviewDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

// Hyperdrive binding.
type ProjectListResponseDeploymentConfigsPreviewHyperdriveBinding struct {
	ID   string                                                           `json:"id,required"`
	JSON projectListResponseDeploymentConfigsPreviewHyperdriveBindingJSON `json:"-"`
}

// projectListResponseDeploymentConfigsPreviewHyperdriveBindingJSON contains the
// JSON metadata for the struct
// [ProjectListResponseDeploymentConfigsPreviewHyperdriveBinding]
type projectListResponseDeploymentConfigsPreviewHyperdriveBindingJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsPreviewHyperdriveBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsPreviewHyperdriveBindingJSON) RawJSON() string {
	return r.raw
}

// KV namespace binding.
type ProjectListResponseDeploymentConfigsPreviewKVNamespace struct {
	// ID of the KV namespace.
	NamespaceID string                                                     `json:"namespace_id,required"`
	JSON        projectListResponseDeploymentConfigsPreviewKVNamespaceJSON `json:"-"`
}

// projectListResponseDeploymentConfigsPreviewKVNamespaceJSON contains the JSON
// metadata for the struct [ProjectListResponseDeploymentConfigsPreviewKVNamespace]
type projectListResponseDeploymentConfigsPreviewKVNamespaceJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsPreviewKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsPreviewKVNamespaceJSON) RawJSON() string {
	return r.raw
}

// Limits for Pages Functions.
type ProjectListResponseDeploymentConfigsPreviewLimits struct {
	// CPU time limit in milliseconds.
	CPUMs int64                                                 `json:"cpu_ms,required"`
	JSON  projectListResponseDeploymentConfigsPreviewLimitsJSON `json:"-"`
}

// projectListResponseDeploymentConfigsPreviewLimitsJSON contains the JSON metadata
// for the struct [ProjectListResponseDeploymentConfigsPreviewLimits]
type projectListResponseDeploymentConfigsPreviewLimitsJSON struct {
	CPUMs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsPreviewLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsPreviewLimitsJSON) RawJSON() string {
	return r.raw
}

// mTLS binding.
type ProjectListResponseDeploymentConfigsPreviewMTLSCertificate struct {
	CertificateID string                                                         `json:"certificate_id,required"`
	JSON          projectListResponseDeploymentConfigsPreviewMTLSCertificateJSON `json:"-"`
}

// projectListResponseDeploymentConfigsPreviewMTLSCertificateJSON contains the JSON
// metadata for the struct
// [ProjectListResponseDeploymentConfigsPreviewMTLSCertificate]
type projectListResponseDeploymentConfigsPreviewMTLSCertificateJSON struct {
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsPreviewMTLSCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsPreviewMTLSCertificateJSON) RawJSON() string {
	return r.raw
}

// Placement setting used for Pages Functions.
type ProjectListResponseDeploymentConfigsPreviewPlacement struct {
	// Placement mode.
	Mode string                                                   `json:"mode,required"`
	JSON projectListResponseDeploymentConfigsPreviewPlacementJSON `json:"-"`
}

// projectListResponseDeploymentConfigsPreviewPlacementJSON contains the JSON
// metadata for the struct [ProjectListResponseDeploymentConfigsPreviewPlacement]
type projectListResponseDeploymentConfigsPreviewPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsPreviewPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsPreviewPlacementJSON) RawJSON() string {
	return r.raw
}

// Queue Producer binding.
type ProjectListResponseDeploymentConfigsPreviewQueueProducer struct {
	// Name of the Queue.
	Name string                                                       `json:"name,required"`
	JSON projectListResponseDeploymentConfigsPreviewQueueProducerJSON `json:"-"`
}

// projectListResponseDeploymentConfigsPreviewQueueProducerJSON contains the JSON
// metadata for the struct
// [ProjectListResponseDeploymentConfigsPreviewQueueProducer]
type projectListResponseDeploymentConfigsPreviewQueueProducerJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsPreviewQueueProducer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsPreviewQueueProducerJSON) RawJSON() string {
	return r.raw
}

// R2 binding.
type ProjectListResponseDeploymentConfigsPreviewR2Bucket struct {
	// Name of the R2 bucket.
	Name string `json:"name,required"`
	// Jurisdiction of the R2 bucket.
	Jurisdiction string                                                  `json:"jurisdiction,nullable"`
	JSON         projectListResponseDeploymentConfigsPreviewR2BucketJSON `json:"-"`
}

// projectListResponseDeploymentConfigsPreviewR2BucketJSON contains the JSON
// metadata for the struct [ProjectListResponseDeploymentConfigsPreviewR2Bucket]
type projectListResponseDeploymentConfigsPreviewR2BucketJSON struct {
	Name         apijson.Field
	Jurisdiction apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsPreviewR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsPreviewR2BucketJSON) RawJSON() string {
	return r.raw
}

// Service binding.
type ProjectListResponseDeploymentConfigsPreviewService struct {
	// The Service environment.
	Environment string `json:"environment,required"`
	// The Service name.
	Service string `json:"service,required"`
	// The entrypoint to bind to.
	Entrypoint string                                                 `json:"entrypoint,nullable"`
	JSON       projectListResponseDeploymentConfigsPreviewServiceJSON `json:"-"`
}

// projectListResponseDeploymentConfigsPreviewServiceJSON contains the JSON
// metadata for the struct [ProjectListResponseDeploymentConfigsPreviewService]
type projectListResponseDeploymentConfigsPreviewServiceJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	Entrypoint  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsPreviewService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsPreviewServiceJSON) RawJSON() string {
	return r.raw
}

// Vectorize binding.
type ProjectListResponseDeploymentConfigsPreviewVectorizeBinding struct {
	IndexName string                                                          `json:"index_name,required"`
	JSON      projectListResponseDeploymentConfigsPreviewVectorizeBindingJSON `json:"-"`
}

// projectListResponseDeploymentConfigsPreviewVectorizeBindingJSON contains the
// JSON metadata for the struct
// [ProjectListResponseDeploymentConfigsPreviewVectorizeBinding]
type projectListResponseDeploymentConfigsPreviewVectorizeBindingJSON struct {
	IndexName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsPreviewVectorizeBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsPreviewVectorizeBindingJSON) RawJSON() string {
	return r.raw
}

// Configs for production deploys.
type ProjectListResponseDeploymentConfigsProduction struct {
	// Whether to always use the latest compatibility date for Pages Functions.
	AlwaysUseLatestCompatibilityDate bool `json:"always_use_latest_compatibility_date,required"`
	// The major version of the build image to use for Pages Functions.
	BuildImageMajorVersion int64 `json:"build_image_major_version,required"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate string `json:"compatibility_date,required"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags []string `json:"compatibility_flags,required"`
	// Environment variables used for builds and Pages Functions.
	EnvVars map[string]ProjectListResponseDeploymentConfigsProductionEnvVar `json:"env_vars,required,nullable"`
	// Whether to fail open when the deployment config cannot be applied.
	FailOpen bool `json:"fail_open,required"`
	// The usage model for Pages Functions.
	//
	// Deprecated: All new projects now use the Standard usage model.
	UsageModel ProjectListResponseDeploymentConfigsProductionUsageModel `json:"usage_model,required"`
	// Constellation bindings used for Pages Functions.
	AIBindings map[string]ProjectListResponseDeploymentConfigsProductionAIBinding `json:"ai_bindings"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets map[string]ProjectListResponseDeploymentConfigsProductionAnalyticsEngineDataset `json:"analytics_engine_datasets"`
	// Browser bindings used for Pages Functions.
	Browsers map[string]ProjectListResponseDeploymentConfigsProductionBrowser `json:"browsers"`
	// D1 databases used for Pages Functions.
	D1Databases map[string]ProjectListResponseDeploymentConfigsProductionD1Database `json:"d1_databases"`
	// Durable Object namespaces used for Pages Functions.
	DurableObjectNamespaces map[string]ProjectListResponseDeploymentConfigsProductionDurableObjectNamespace `json:"durable_object_namespaces"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings map[string]ProjectListResponseDeploymentConfigsProductionHyperdriveBinding `json:"hyperdrive_bindings"`
	// KV namespaces used for Pages Functions.
	KVNamespaces map[string]ProjectListResponseDeploymentConfigsProductionKVNamespace `json:"kv_namespaces"`
	// Limits for Pages Functions.
	Limits ProjectListResponseDeploymentConfigsProductionLimits `json:"limits"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates map[string]ProjectListResponseDeploymentConfigsProductionMTLSCertificate `json:"mtls_certificates"`
	// Placement setting used for Pages Functions.
	Placement ProjectListResponseDeploymentConfigsProductionPlacement `json:"placement"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers map[string]ProjectListResponseDeploymentConfigsProductionQueueProducer `json:"queue_producers"`
	// R2 buckets used for Pages Functions.
	R2Buckets map[string]ProjectListResponseDeploymentConfigsProductionR2Bucket `json:"r2_buckets"`
	// Services used for Pages Functions.
	Services map[string]ProjectListResponseDeploymentConfigsProductionService `json:"services"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings map[string]ProjectListResponseDeploymentConfigsProductionVectorizeBinding `json:"vectorize_bindings"`
	// Hash of the Wrangler configuration used for the deployment.
	WranglerConfigHash string                                             `json:"wrangler_config_hash"`
	JSON               projectListResponseDeploymentConfigsProductionJSON `json:"-"`
}

// projectListResponseDeploymentConfigsProductionJSON contains the JSON metadata
// for the struct [ProjectListResponseDeploymentConfigsProduction]
type projectListResponseDeploymentConfigsProductionJSON struct {
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

func (r *ProjectListResponseDeploymentConfigsProduction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsProductionJSON) RawJSON() string {
	return r.raw
}

// A plaintext environment variable.
type ProjectListResponseDeploymentConfigsProductionEnvVar struct {
	Type ProjectListResponseDeploymentConfigsProductionEnvVarsType `json:"type,required"`
	// Environment variable value.
	Value string                                                   `json:"value,required"`
	JSON  projectListResponseDeploymentConfigsProductionEnvVarJSON `json:"-"`
	union ProjectListResponseDeploymentConfigsProductionEnvVarsUnion
}

// projectListResponseDeploymentConfigsProductionEnvVarJSON contains the JSON
// metadata for the struct [ProjectListResponseDeploymentConfigsProductionEnvVar]
type projectListResponseDeploymentConfigsProductionEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectListResponseDeploymentConfigsProductionEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectListResponseDeploymentConfigsProductionEnvVar) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectListResponseDeploymentConfigsProductionEnvVar{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectListResponseDeploymentConfigsProductionEnvVarsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectListResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar],
// [ProjectListResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar].
func (r ProjectListResponseDeploymentConfigsProductionEnvVar) AsUnion() ProjectListResponseDeploymentConfigsProductionEnvVarsUnion {
	return r.union
}

// A plaintext environment variable.
//
// Union satisfied by
// [ProjectListResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar] or
// [ProjectListResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar].
type ProjectListResponseDeploymentConfigsProductionEnvVarsUnion interface {
	implementsProjectListResponseDeploymentConfigsProductionEnvVar()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectListResponseDeploymentConfigsProductionEnvVarsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectListResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectListResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar{}),
			DiscriminatorValue: "secret_text",
		},
	)
}

// A plaintext environment variable.
type ProjectListResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar struct {
	Type ProjectListResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType `json:"type,required"`
	// Environment variable value.
	Value string                                                                        `json:"value,required"`
	JSON  projectListResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarJSON `json:"-"`
}

// projectListResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarJSON
// contains the JSON metadata for the struct
// [ProjectListResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar]
type projectListResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectListResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar) implementsProjectListResponseDeploymentConfigsProductionEnvVar() {
}

type ProjectListResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType string

const (
	ProjectListResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarTypePlainText ProjectListResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r ProjectListResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectListResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type ProjectListResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar struct {
	Type ProjectListResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType `json:"type,required"`
	// Secret value.
	Value string                                                                         `json:"value,required"`
	JSON  projectListResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarJSON `json:"-"`
}

// projectListResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarJSON
// contains the JSON metadata for the struct
// [ProjectListResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar]
type projectListResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectListResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar) implementsProjectListResponseDeploymentConfigsProductionEnvVar() {
}

type ProjectListResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType string

const (
	ProjectListResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarTypeSecretText ProjectListResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r ProjectListResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectListResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type ProjectListResponseDeploymentConfigsProductionEnvVarsType string

const (
	ProjectListResponseDeploymentConfigsProductionEnvVarsTypePlainText  ProjectListResponseDeploymentConfigsProductionEnvVarsType = "plain_text"
	ProjectListResponseDeploymentConfigsProductionEnvVarsTypeSecretText ProjectListResponseDeploymentConfigsProductionEnvVarsType = "secret_text"
)

func (r ProjectListResponseDeploymentConfigsProductionEnvVarsType) IsKnown() bool {
	switch r {
	case ProjectListResponseDeploymentConfigsProductionEnvVarsTypePlainText, ProjectListResponseDeploymentConfigsProductionEnvVarsTypeSecretText:
		return true
	}
	return false
}

// The usage model for Pages Functions.
type ProjectListResponseDeploymentConfigsProductionUsageModel string

const (
	ProjectListResponseDeploymentConfigsProductionUsageModelStandard ProjectListResponseDeploymentConfigsProductionUsageModel = "standard"
	ProjectListResponseDeploymentConfigsProductionUsageModelBundled  ProjectListResponseDeploymentConfigsProductionUsageModel = "bundled"
	ProjectListResponseDeploymentConfigsProductionUsageModelUnbound  ProjectListResponseDeploymentConfigsProductionUsageModel = "unbound"
)

func (r ProjectListResponseDeploymentConfigsProductionUsageModel) IsKnown() bool {
	switch r {
	case ProjectListResponseDeploymentConfigsProductionUsageModelStandard, ProjectListResponseDeploymentConfigsProductionUsageModelBundled, ProjectListResponseDeploymentConfigsProductionUsageModelUnbound:
		return true
	}
	return false
}

// AI binding.
type ProjectListResponseDeploymentConfigsProductionAIBinding struct {
	ProjectID string                                                      `json:"project_id,required"`
	JSON      projectListResponseDeploymentConfigsProductionAIBindingJSON `json:"-"`
}

// projectListResponseDeploymentConfigsProductionAIBindingJSON contains the JSON
// metadata for the struct
// [ProjectListResponseDeploymentConfigsProductionAIBinding]
type projectListResponseDeploymentConfigsProductionAIBindingJSON struct {
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsProductionAIBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsProductionAIBindingJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine binding.
type ProjectListResponseDeploymentConfigsProductionAnalyticsEngineDataset struct {
	// Name of the dataset.
	Dataset string                                                                   `json:"dataset,required"`
	JSON    projectListResponseDeploymentConfigsProductionAnalyticsEngineDatasetJSON `json:"-"`
}

// projectListResponseDeploymentConfigsProductionAnalyticsEngineDatasetJSON
// contains the JSON metadata for the struct
// [ProjectListResponseDeploymentConfigsProductionAnalyticsEngineDataset]
type projectListResponseDeploymentConfigsProductionAnalyticsEngineDatasetJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsProductionAnalyticsEngineDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsProductionAnalyticsEngineDatasetJSON) RawJSON() string {
	return r.raw
}

// Browser binding.
type ProjectListResponseDeploymentConfigsProductionBrowser struct {
	JSON projectListResponseDeploymentConfigsProductionBrowserJSON `json:"-"`
}

// projectListResponseDeploymentConfigsProductionBrowserJSON contains the JSON
// metadata for the struct [ProjectListResponseDeploymentConfigsProductionBrowser]
type projectListResponseDeploymentConfigsProductionBrowserJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsProductionBrowser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsProductionBrowserJSON) RawJSON() string {
	return r.raw
}

// D1 binding.
type ProjectListResponseDeploymentConfigsProductionD1Database struct {
	// UUID of the D1 database.
	ID   string                                                       `json:"id,required"`
	JSON projectListResponseDeploymentConfigsProductionD1DatabaseJSON `json:"-"`
}

// projectListResponseDeploymentConfigsProductionD1DatabaseJSON contains the JSON
// metadata for the struct
// [ProjectListResponseDeploymentConfigsProductionD1Database]
type projectListResponseDeploymentConfigsProductionD1DatabaseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsProductionD1Database) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsProductionD1DatabaseJSON) RawJSON() string {
	return r.raw
}

// Durable Object binding.
type ProjectListResponseDeploymentConfigsProductionDurableObjectNamespace struct {
	// ID of the Durable Object namespace.
	NamespaceID string                                                                   `json:"namespace_id,required"`
	JSON        projectListResponseDeploymentConfigsProductionDurableObjectNamespaceJSON `json:"-"`
}

// projectListResponseDeploymentConfigsProductionDurableObjectNamespaceJSON
// contains the JSON metadata for the struct
// [ProjectListResponseDeploymentConfigsProductionDurableObjectNamespace]
type projectListResponseDeploymentConfigsProductionDurableObjectNamespaceJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsProductionDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsProductionDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

// Hyperdrive binding.
type ProjectListResponseDeploymentConfigsProductionHyperdriveBinding struct {
	ID   string                                                              `json:"id,required"`
	JSON projectListResponseDeploymentConfigsProductionHyperdriveBindingJSON `json:"-"`
}

// projectListResponseDeploymentConfigsProductionHyperdriveBindingJSON contains the
// JSON metadata for the struct
// [ProjectListResponseDeploymentConfigsProductionHyperdriveBinding]
type projectListResponseDeploymentConfigsProductionHyperdriveBindingJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsProductionHyperdriveBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsProductionHyperdriveBindingJSON) RawJSON() string {
	return r.raw
}

// KV namespace binding.
type ProjectListResponseDeploymentConfigsProductionKVNamespace struct {
	// ID of the KV namespace.
	NamespaceID string                                                        `json:"namespace_id,required"`
	JSON        projectListResponseDeploymentConfigsProductionKVNamespaceJSON `json:"-"`
}

// projectListResponseDeploymentConfigsProductionKVNamespaceJSON contains the JSON
// metadata for the struct
// [ProjectListResponseDeploymentConfigsProductionKVNamespace]
type projectListResponseDeploymentConfigsProductionKVNamespaceJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsProductionKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsProductionKVNamespaceJSON) RawJSON() string {
	return r.raw
}

// Limits for Pages Functions.
type ProjectListResponseDeploymentConfigsProductionLimits struct {
	// CPU time limit in milliseconds.
	CPUMs int64                                                    `json:"cpu_ms,required"`
	JSON  projectListResponseDeploymentConfigsProductionLimitsJSON `json:"-"`
}

// projectListResponseDeploymentConfigsProductionLimitsJSON contains the JSON
// metadata for the struct [ProjectListResponseDeploymentConfigsProductionLimits]
type projectListResponseDeploymentConfigsProductionLimitsJSON struct {
	CPUMs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsProductionLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsProductionLimitsJSON) RawJSON() string {
	return r.raw
}

// mTLS binding.
type ProjectListResponseDeploymentConfigsProductionMTLSCertificate struct {
	CertificateID string                                                            `json:"certificate_id,required"`
	JSON          projectListResponseDeploymentConfigsProductionMTLSCertificateJSON `json:"-"`
}

// projectListResponseDeploymentConfigsProductionMTLSCertificateJSON contains the
// JSON metadata for the struct
// [ProjectListResponseDeploymentConfigsProductionMTLSCertificate]
type projectListResponseDeploymentConfigsProductionMTLSCertificateJSON struct {
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsProductionMTLSCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsProductionMTLSCertificateJSON) RawJSON() string {
	return r.raw
}

// Placement setting used for Pages Functions.
type ProjectListResponseDeploymentConfigsProductionPlacement struct {
	// Placement mode.
	Mode string                                                      `json:"mode,required"`
	JSON projectListResponseDeploymentConfigsProductionPlacementJSON `json:"-"`
}

// projectListResponseDeploymentConfigsProductionPlacementJSON contains the JSON
// metadata for the struct
// [ProjectListResponseDeploymentConfigsProductionPlacement]
type projectListResponseDeploymentConfigsProductionPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsProductionPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsProductionPlacementJSON) RawJSON() string {
	return r.raw
}

// Queue Producer binding.
type ProjectListResponseDeploymentConfigsProductionQueueProducer struct {
	// Name of the Queue.
	Name string                                                          `json:"name,required"`
	JSON projectListResponseDeploymentConfigsProductionQueueProducerJSON `json:"-"`
}

// projectListResponseDeploymentConfigsProductionQueueProducerJSON contains the
// JSON metadata for the struct
// [ProjectListResponseDeploymentConfigsProductionQueueProducer]
type projectListResponseDeploymentConfigsProductionQueueProducerJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsProductionQueueProducer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsProductionQueueProducerJSON) RawJSON() string {
	return r.raw
}

// R2 binding.
type ProjectListResponseDeploymentConfigsProductionR2Bucket struct {
	// Name of the R2 bucket.
	Name string `json:"name,required"`
	// Jurisdiction of the R2 bucket.
	Jurisdiction string                                                     `json:"jurisdiction,nullable"`
	JSON         projectListResponseDeploymentConfigsProductionR2BucketJSON `json:"-"`
}

// projectListResponseDeploymentConfigsProductionR2BucketJSON contains the JSON
// metadata for the struct [ProjectListResponseDeploymentConfigsProductionR2Bucket]
type projectListResponseDeploymentConfigsProductionR2BucketJSON struct {
	Name         apijson.Field
	Jurisdiction apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsProductionR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsProductionR2BucketJSON) RawJSON() string {
	return r.raw
}

// Service binding.
type ProjectListResponseDeploymentConfigsProductionService struct {
	// The Service environment.
	Environment string `json:"environment,required"`
	// The Service name.
	Service string `json:"service,required"`
	// The entrypoint to bind to.
	Entrypoint string                                                    `json:"entrypoint,nullable"`
	JSON       projectListResponseDeploymentConfigsProductionServiceJSON `json:"-"`
}

// projectListResponseDeploymentConfigsProductionServiceJSON contains the JSON
// metadata for the struct [ProjectListResponseDeploymentConfigsProductionService]
type projectListResponseDeploymentConfigsProductionServiceJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	Entrypoint  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsProductionService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsProductionServiceJSON) RawJSON() string {
	return r.raw
}

// Vectorize binding.
type ProjectListResponseDeploymentConfigsProductionVectorizeBinding struct {
	IndexName string                                                             `json:"index_name,required"`
	JSON      projectListResponseDeploymentConfigsProductionVectorizeBindingJSON `json:"-"`
}

// projectListResponseDeploymentConfigsProductionVectorizeBindingJSON contains the
// JSON metadata for the struct
// [ProjectListResponseDeploymentConfigsProductionVectorizeBinding]
type projectListResponseDeploymentConfigsProductionVectorizeBindingJSON struct {
	IndexName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseDeploymentConfigsProductionVectorizeBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseDeploymentConfigsProductionVectorizeBindingJSON) RawJSON() string {
	return r.raw
}

// Most recent deployment of the project.
type ProjectListResponseLatestDeployment struct {
	// Id of the deployment.
	ID string `json:"id,required"`
	// A list of alias URLs pointing to this deployment.
	Aliases []string `json:"aliases,required,nullable"`
	// Configs for the project build process.
	BuildConfig ProjectListResponseLatestDeploymentBuildConfig `json:"build_config,required"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger ProjectListResponseLatestDeploymentDeploymentTrigger `json:"deployment_trigger,required"`
	// Environment variables used for builds and Pages Functions.
	EnvVars map[string]ProjectListResponseLatestDeploymentEnvVar `json:"env_vars,required,nullable"`
	// Type of deploy.
	Environment ProjectListResponseLatestDeploymentEnvironment `json:"environment,required"`
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
	Source ProjectListResponseLatestDeploymentSource `json:"source,required"`
	// List of past stages.
	Stages []Stage `json:"stages,required"`
	// The live URL to view this deployment.
	URL string `json:"url,required"`
	// Whether the deployment uses functions.
	UsesFunctions bool                                    `json:"uses_functions,nullable"`
	JSON          projectListResponseLatestDeploymentJSON `json:"-"`
}

// projectListResponseLatestDeploymentJSON contains the JSON metadata for the
// struct [ProjectListResponseLatestDeployment]
type projectListResponseLatestDeploymentJSON struct {
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

func (r *ProjectListResponseLatestDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseLatestDeploymentJSON) RawJSON() string {
	return r.raw
}

// Configs for the project build process.
type ProjectListResponseLatestDeploymentBuildConfig struct {
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
	RootDir string                                             `json:"root_dir,nullable"`
	JSON    projectListResponseLatestDeploymentBuildConfigJSON `json:"-"`
}

// projectListResponseLatestDeploymentBuildConfigJSON contains the JSON metadata
// for the struct [ProjectListResponseLatestDeploymentBuildConfig]
type projectListResponseLatestDeploymentBuildConfigJSON struct {
	WebAnalyticsTag   apijson.Field
	WebAnalyticsToken apijson.Field
	BuildCaching      apijson.Field
	BuildCommand      apijson.Field
	DestinationDir    apijson.Field
	RootDir           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectListResponseLatestDeploymentBuildConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseLatestDeploymentBuildConfigJSON) RawJSON() string {
	return r.raw
}

// Info about what caused the deployment.
type ProjectListResponseLatestDeploymentDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata ProjectListResponseLatestDeploymentDeploymentTriggerMetadata `json:"metadata,required"`
	// What caused the deployment.
	Type ProjectListResponseLatestDeploymentDeploymentTriggerType `json:"type,required"`
	JSON projectListResponseLatestDeploymentDeploymentTriggerJSON `json:"-"`
}

// projectListResponseLatestDeploymentDeploymentTriggerJSON contains the JSON
// metadata for the struct [ProjectListResponseLatestDeploymentDeploymentTrigger]
type projectListResponseLatestDeploymentDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseLatestDeploymentDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseLatestDeploymentDeploymentTriggerJSON) RawJSON() string {
	return r.raw
}

// Additional info about the trigger.
type ProjectListResponseLatestDeploymentDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch,required"`
	// Whether the deployment trigger commit was dirty.
	CommitDirty bool `json:"commit_dirty,required"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash,required"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                           `json:"commit_message,required"`
	JSON          projectListResponseLatestDeploymentDeploymentTriggerMetadataJSON `json:"-"`
}

// projectListResponseLatestDeploymentDeploymentTriggerMetadataJSON contains the
// JSON metadata for the struct
// [ProjectListResponseLatestDeploymentDeploymentTriggerMetadata]
type projectListResponseLatestDeploymentDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitDirty   apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectListResponseLatestDeploymentDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseLatestDeploymentDeploymentTriggerMetadataJSON) RawJSON() string {
	return r.raw
}

// What caused the deployment.
type ProjectListResponseLatestDeploymentDeploymentTriggerType string

const (
	ProjectListResponseLatestDeploymentDeploymentTriggerTypeGitHubPush ProjectListResponseLatestDeploymentDeploymentTriggerType = "github:push"
	ProjectListResponseLatestDeploymentDeploymentTriggerTypeADHoc      ProjectListResponseLatestDeploymentDeploymentTriggerType = "ad_hoc"
	ProjectListResponseLatestDeploymentDeploymentTriggerTypeDeployHook ProjectListResponseLatestDeploymentDeploymentTriggerType = "deploy_hook"
)

func (r ProjectListResponseLatestDeploymentDeploymentTriggerType) IsKnown() bool {
	switch r {
	case ProjectListResponseLatestDeploymentDeploymentTriggerTypeGitHubPush, ProjectListResponseLatestDeploymentDeploymentTriggerTypeADHoc, ProjectListResponseLatestDeploymentDeploymentTriggerTypeDeployHook:
		return true
	}
	return false
}

// A plaintext environment variable.
type ProjectListResponseLatestDeploymentEnvVar struct {
	Type ProjectListResponseLatestDeploymentEnvVarsType `json:"type,required"`
	// Environment variable value.
	Value string                                        `json:"value,required"`
	JSON  projectListResponseLatestDeploymentEnvVarJSON `json:"-"`
	union ProjectListResponseLatestDeploymentEnvVarsUnion
}

// projectListResponseLatestDeploymentEnvVarJSON contains the JSON metadata for the
// struct [ProjectListResponseLatestDeploymentEnvVar]
type projectListResponseLatestDeploymentEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectListResponseLatestDeploymentEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectListResponseLatestDeploymentEnvVar) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectListResponseLatestDeploymentEnvVar{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectListResponseLatestDeploymentEnvVarsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectListResponseLatestDeploymentEnvVarsPagesPlainTextEnvVar],
// [ProjectListResponseLatestDeploymentEnvVarsPagesSecretTextEnvVar].
func (r ProjectListResponseLatestDeploymentEnvVar) AsUnion() ProjectListResponseLatestDeploymentEnvVarsUnion {
	return r.union
}

// A plaintext environment variable.
//
// Union satisfied by
// [ProjectListResponseLatestDeploymentEnvVarsPagesPlainTextEnvVar] or
// [ProjectListResponseLatestDeploymentEnvVarsPagesSecretTextEnvVar].
type ProjectListResponseLatestDeploymentEnvVarsUnion interface {
	implementsProjectListResponseLatestDeploymentEnvVar()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectListResponseLatestDeploymentEnvVarsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectListResponseLatestDeploymentEnvVarsPagesPlainTextEnvVar{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectListResponseLatestDeploymentEnvVarsPagesSecretTextEnvVar{}),
			DiscriminatorValue: "secret_text",
		},
	)
}

// A plaintext environment variable.
type ProjectListResponseLatestDeploymentEnvVarsPagesPlainTextEnvVar struct {
	Type ProjectListResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarType `json:"type,required"`
	// Environment variable value.
	Value string                                                             `json:"value,required"`
	JSON  projectListResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarJSON `json:"-"`
}

// projectListResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarJSON contains the
// JSON metadata for the struct
// [ProjectListResponseLatestDeploymentEnvVarsPagesPlainTextEnvVar]
type projectListResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseLatestDeploymentEnvVarsPagesPlainTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectListResponseLatestDeploymentEnvVarsPagesPlainTextEnvVar) implementsProjectListResponseLatestDeploymentEnvVar() {
}

type ProjectListResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarType string

const (
	ProjectListResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarTypePlainText ProjectListResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r ProjectListResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectListResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type ProjectListResponseLatestDeploymentEnvVarsPagesSecretTextEnvVar struct {
	Type ProjectListResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarType `json:"type,required"`
	// Secret value.
	Value string                                                              `json:"value,required"`
	JSON  projectListResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarJSON `json:"-"`
}

// projectListResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarJSON contains the
// JSON metadata for the struct
// [ProjectListResponseLatestDeploymentEnvVarsPagesSecretTextEnvVar]
type projectListResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseLatestDeploymentEnvVarsPagesSecretTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectListResponseLatestDeploymentEnvVarsPagesSecretTextEnvVar) implementsProjectListResponseLatestDeploymentEnvVar() {
}

type ProjectListResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarType string

const (
	ProjectListResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarTypeSecretText ProjectListResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r ProjectListResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectListResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type ProjectListResponseLatestDeploymentEnvVarsType string

const (
	ProjectListResponseLatestDeploymentEnvVarsTypePlainText  ProjectListResponseLatestDeploymentEnvVarsType = "plain_text"
	ProjectListResponseLatestDeploymentEnvVarsTypeSecretText ProjectListResponseLatestDeploymentEnvVarsType = "secret_text"
)

func (r ProjectListResponseLatestDeploymentEnvVarsType) IsKnown() bool {
	switch r {
	case ProjectListResponseLatestDeploymentEnvVarsTypePlainText, ProjectListResponseLatestDeploymentEnvVarsTypeSecretText:
		return true
	}
	return false
}

// Type of deploy.
type ProjectListResponseLatestDeploymentEnvironment string

const (
	ProjectListResponseLatestDeploymentEnvironmentPreview    ProjectListResponseLatestDeploymentEnvironment = "preview"
	ProjectListResponseLatestDeploymentEnvironmentProduction ProjectListResponseLatestDeploymentEnvironment = "production"
)

func (r ProjectListResponseLatestDeploymentEnvironment) IsKnown() bool {
	switch r {
	case ProjectListResponseLatestDeploymentEnvironmentPreview, ProjectListResponseLatestDeploymentEnvironmentProduction:
		return true
	}
	return false
}

// Configs for the project source control.
type ProjectListResponseLatestDeploymentSource struct {
	Config ProjectListResponseLatestDeploymentSourceConfig `json:"config,required"`
	// The source control management provider.
	Type ProjectListResponseLatestDeploymentSourceType `json:"type,required"`
	JSON projectListResponseLatestDeploymentSourceJSON `json:"-"`
}

// projectListResponseLatestDeploymentSourceJSON contains the JSON metadata for the
// struct [ProjectListResponseLatestDeploymentSource]
type projectListResponseLatestDeploymentSourceJSON struct {
	Config      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseLatestDeploymentSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseLatestDeploymentSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectListResponseLatestDeploymentSourceConfig struct {
	// The owner of the repository.
	Owner string `json:"owner,required"`
	// Whether to enable PR comments.
	PrCommentsEnabled bool `json:"pr_comments_enabled,required"`
	// The production branch of the repository.
	ProductionBranch string `json:"production_branch,required"`
	// The name of the repository.
	RepoName string `json:"repo_name,required"`
	// Whether to enable automatic deployments when pushing to the source repository.
	// When disabled, no deployments (production or preview) will be triggered
	// automatically.
	//
	// Deprecated: Use `production_deployments_enabled` and
	// `preview_deployment_setting` for more granular control.
	DeploymentsEnabled bool `json:"deployments_enabled"`
	// The owner ID of the repository.
	OwnerID string `json:"owner_id"`
	// A list of paths that should be excluded from triggering a preview deployment.
	// Wildcard syntax (`*`) is supported.
	PathExcludes []string `json:"path_excludes"`
	// A list of paths that should be watched to trigger a preview deployment. Wildcard
	// syntax (`*`) is supported.
	PathIncludes []string `json:"path_includes"`
	// A list of branches that should not trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchExcludes []string `json:"preview_branch_excludes"`
	// A list of branches that should trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchIncludes []string `json:"preview_branch_includes"`
	// Controls whether commits to preview branches trigger a preview deployment.
	PreviewDeploymentSetting ProjectListResponseLatestDeploymentSourceConfigPreviewDeploymentSetting `json:"preview_deployment_setting"`
	// Whether to trigger a production deployment on commits to the production branch.
	ProductionDeploymentsEnabled bool `json:"production_deployments_enabled"`
	// The ID of the repository.
	RepoID string                                              `json:"repo_id"`
	JSON   projectListResponseLatestDeploymentSourceConfigJSON `json:"-"`
}

// projectListResponseLatestDeploymentSourceConfigJSON contains the JSON metadata
// for the struct [ProjectListResponseLatestDeploymentSourceConfig]
type projectListResponseLatestDeploymentSourceConfigJSON struct {
	Owner                        apijson.Field
	PrCommentsEnabled            apijson.Field
	ProductionBranch             apijson.Field
	RepoName                     apijson.Field
	DeploymentsEnabled           apijson.Field
	OwnerID                      apijson.Field
	PathExcludes                 apijson.Field
	PathIncludes                 apijson.Field
	PreviewBranchExcludes        apijson.Field
	PreviewBranchIncludes        apijson.Field
	PreviewDeploymentSetting     apijson.Field
	ProductionDeploymentsEnabled apijson.Field
	RepoID                       apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ProjectListResponseLatestDeploymentSourceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseLatestDeploymentSourceConfigJSON) RawJSON() string {
	return r.raw
}

// Controls whether commits to preview branches trigger a preview deployment.
type ProjectListResponseLatestDeploymentSourceConfigPreviewDeploymentSetting string

const (
	ProjectListResponseLatestDeploymentSourceConfigPreviewDeploymentSettingAll    ProjectListResponseLatestDeploymentSourceConfigPreviewDeploymentSetting = "all"
	ProjectListResponseLatestDeploymentSourceConfigPreviewDeploymentSettingNone   ProjectListResponseLatestDeploymentSourceConfigPreviewDeploymentSetting = "none"
	ProjectListResponseLatestDeploymentSourceConfigPreviewDeploymentSettingCustom ProjectListResponseLatestDeploymentSourceConfigPreviewDeploymentSetting = "custom"
)

func (r ProjectListResponseLatestDeploymentSourceConfigPreviewDeploymentSetting) IsKnown() bool {
	switch r {
	case ProjectListResponseLatestDeploymentSourceConfigPreviewDeploymentSettingAll, ProjectListResponseLatestDeploymentSourceConfigPreviewDeploymentSettingNone, ProjectListResponseLatestDeploymentSourceConfigPreviewDeploymentSettingCustom:
		return true
	}
	return false
}

// The source control management provider.
type ProjectListResponseLatestDeploymentSourceType string

const (
	ProjectListResponseLatestDeploymentSourceTypeGitHub ProjectListResponseLatestDeploymentSourceType = "github"
	ProjectListResponseLatestDeploymentSourceTypeGitlab ProjectListResponseLatestDeploymentSourceType = "gitlab"
)

func (r ProjectListResponseLatestDeploymentSourceType) IsKnown() bool {
	switch r {
	case ProjectListResponseLatestDeploymentSourceTypeGitHub, ProjectListResponseLatestDeploymentSourceTypeGitlab:
		return true
	}
	return false
}

// Configs for the project build process.
type ProjectListResponseBuildConfig struct {
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
	RootDir string                             `json:"root_dir,nullable"`
	JSON    projectListResponseBuildConfigJSON `json:"-"`
}

// projectListResponseBuildConfigJSON contains the JSON metadata for the struct
// [ProjectListResponseBuildConfig]
type projectListResponseBuildConfigJSON struct {
	WebAnalyticsTag   apijson.Field
	WebAnalyticsToken apijson.Field
	BuildCaching      apijson.Field
	BuildCommand      apijson.Field
	DestinationDir    apijson.Field
	RootDir           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectListResponseBuildConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseBuildConfigJSON) RawJSON() string {
	return r.raw
}

// Configs for the project source control.
type ProjectListResponseSource struct {
	Config ProjectListResponseSourceConfig `json:"config,required"`
	// The source control management provider.
	Type ProjectListResponseSourceType `json:"type,required"`
	JSON projectListResponseSourceJSON `json:"-"`
}

// projectListResponseSourceJSON contains the JSON metadata for the struct
// [ProjectListResponseSource]
type projectListResponseSourceJSON struct {
	Config      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectListResponseSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectListResponseSourceConfig struct {
	// The owner of the repository.
	Owner string `json:"owner,required"`
	// Whether to enable PR comments.
	PrCommentsEnabled bool `json:"pr_comments_enabled,required"`
	// The production branch of the repository.
	ProductionBranch string `json:"production_branch,required"`
	// The name of the repository.
	RepoName string `json:"repo_name,required"`
	// Whether to enable automatic deployments when pushing to the source repository.
	// When disabled, no deployments (production or preview) will be triggered
	// automatically.
	//
	// Deprecated: Use `production_deployments_enabled` and
	// `preview_deployment_setting` for more granular control.
	DeploymentsEnabled bool `json:"deployments_enabled"`
	// The owner ID of the repository.
	OwnerID string `json:"owner_id"`
	// A list of paths that should be excluded from triggering a preview deployment.
	// Wildcard syntax (`*`) is supported.
	PathExcludes []string `json:"path_excludes"`
	// A list of paths that should be watched to trigger a preview deployment. Wildcard
	// syntax (`*`) is supported.
	PathIncludes []string `json:"path_includes"`
	// A list of branches that should not trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchExcludes []string `json:"preview_branch_excludes"`
	// A list of branches that should trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchIncludes []string `json:"preview_branch_includes"`
	// Controls whether commits to preview branches trigger a preview deployment.
	PreviewDeploymentSetting ProjectListResponseSourceConfigPreviewDeploymentSetting `json:"preview_deployment_setting"`
	// Whether to trigger a production deployment on commits to the production branch.
	ProductionDeploymentsEnabled bool `json:"production_deployments_enabled"`
	// The ID of the repository.
	RepoID string                              `json:"repo_id"`
	JSON   projectListResponseSourceConfigJSON `json:"-"`
}

// projectListResponseSourceConfigJSON contains the JSON metadata for the struct
// [ProjectListResponseSourceConfig]
type projectListResponseSourceConfigJSON struct {
	Owner                        apijson.Field
	PrCommentsEnabled            apijson.Field
	ProductionBranch             apijson.Field
	RepoName                     apijson.Field
	DeploymentsEnabled           apijson.Field
	OwnerID                      apijson.Field
	PathExcludes                 apijson.Field
	PathIncludes                 apijson.Field
	PreviewBranchExcludes        apijson.Field
	PreviewBranchIncludes        apijson.Field
	PreviewDeploymentSetting     apijson.Field
	ProductionDeploymentsEnabled apijson.Field
	RepoID                       apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ProjectListResponseSourceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectListResponseSourceConfigJSON) RawJSON() string {
	return r.raw
}

// Controls whether commits to preview branches trigger a preview deployment.
type ProjectListResponseSourceConfigPreviewDeploymentSetting string

const (
	ProjectListResponseSourceConfigPreviewDeploymentSettingAll    ProjectListResponseSourceConfigPreviewDeploymentSetting = "all"
	ProjectListResponseSourceConfigPreviewDeploymentSettingNone   ProjectListResponseSourceConfigPreviewDeploymentSetting = "none"
	ProjectListResponseSourceConfigPreviewDeploymentSettingCustom ProjectListResponseSourceConfigPreviewDeploymentSetting = "custom"
)

func (r ProjectListResponseSourceConfigPreviewDeploymentSetting) IsKnown() bool {
	switch r {
	case ProjectListResponseSourceConfigPreviewDeploymentSettingAll, ProjectListResponseSourceConfigPreviewDeploymentSettingNone, ProjectListResponseSourceConfigPreviewDeploymentSettingCustom:
		return true
	}
	return false
}

// The source control management provider.
type ProjectListResponseSourceType string

const (
	ProjectListResponseSourceTypeGitHub ProjectListResponseSourceType = "github"
	ProjectListResponseSourceTypeGitlab ProjectListResponseSourceType = "gitlab"
)

func (r ProjectListResponseSourceType) IsKnown() bool {
	switch r {
	case ProjectListResponseSourceTypeGitHub, ProjectListResponseSourceTypeGitlab:
		return true
	}
	return false
}

type ProjectDeleteResponse = interface{}

type ProjectEditResponse struct {
	// ID of the project.
	ID string `json:"id,required"`
	// Most recent production deployment of the project.
	CanonicalDeployment ProjectEditResponseCanonicalDeployment `json:"canonical_deployment,required,nullable"`
	// When the project was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Configs for deployments in a project.
	DeploymentConfigs ProjectEditResponseDeploymentConfigs `json:"deployment_configs,required"`
	// Framework the project is using.
	Framework string `json:"framework,required"`
	// Version of the framework the project is using.
	FrameworkVersion string `json:"framework_version,required"`
	// Most recent deployment of the project.
	LatestDeployment ProjectEditResponseLatestDeployment `json:"latest_deployment,required,nullable"`
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
	BuildConfig ProjectEditResponseBuildConfig `json:"build_config"`
	// A list of associated custom domains for the project.
	Domains []string `json:"domains"`
	// Configs for the project source control.
	Source ProjectEditResponseSource `json:"source"`
	// The Cloudflare subdomain associated with the project.
	Subdomain string                  `json:"subdomain"`
	JSON      projectEditResponseJSON `json:"-"`
}

// projectEditResponseJSON contains the JSON metadata for the struct
// [ProjectEditResponse]
type projectEditResponseJSON struct {
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

func (r *ProjectEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseJSON) RawJSON() string {
	return r.raw
}

// Most recent production deployment of the project.
type ProjectEditResponseCanonicalDeployment struct {
	// Id of the deployment.
	ID string `json:"id,required"`
	// A list of alias URLs pointing to this deployment.
	Aliases []string `json:"aliases,required,nullable"`
	// Configs for the project build process.
	BuildConfig ProjectEditResponseCanonicalDeploymentBuildConfig `json:"build_config,required"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger ProjectEditResponseCanonicalDeploymentDeploymentTrigger `json:"deployment_trigger,required"`
	// Environment variables used for builds and Pages Functions.
	EnvVars map[string]ProjectEditResponseCanonicalDeploymentEnvVar `json:"env_vars,required,nullable"`
	// Type of deploy.
	Environment ProjectEditResponseCanonicalDeploymentEnvironment `json:"environment,required"`
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
	Source ProjectEditResponseCanonicalDeploymentSource `json:"source,required"`
	// List of past stages.
	Stages []Stage `json:"stages,required"`
	// The live URL to view this deployment.
	URL string `json:"url,required"`
	// Whether the deployment uses functions.
	UsesFunctions bool                                       `json:"uses_functions,nullable"`
	JSON          projectEditResponseCanonicalDeploymentJSON `json:"-"`
}

// projectEditResponseCanonicalDeploymentJSON contains the JSON metadata for the
// struct [ProjectEditResponseCanonicalDeployment]
type projectEditResponseCanonicalDeploymentJSON struct {
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

func (r *ProjectEditResponseCanonicalDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseCanonicalDeploymentJSON) RawJSON() string {
	return r.raw
}

// Configs for the project build process.
type ProjectEditResponseCanonicalDeploymentBuildConfig struct {
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
	RootDir string                                                `json:"root_dir,nullable"`
	JSON    projectEditResponseCanonicalDeploymentBuildConfigJSON `json:"-"`
}

// projectEditResponseCanonicalDeploymentBuildConfigJSON contains the JSON metadata
// for the struct [ProjectEditResponseCanonicalDeploymentBuildConfig]
type projectEditResponseCanonicalDeploymentBuildConfigJSON struct {
	WebAnalyticsTag   apijson.Field
	WebAnalyticsToken apijson.Field
	BuildCaching      apijson.Field
	BuildCommand      apijson.Field
	DestinationDir    apijson.Field
	RootDir           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectEditResponseCanonicalDeploymentBuildConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseCanonicalDeploymentBuildConfigJSON) RawJSON() string {
	return r.raw
}

// Info about what caused the deployment.
type ProjectEditResponseCanonicalDeploymentDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata ProjectEditResponseCanonicalDeploymentDeploymentTriggerMetadata `json:"metadata,required"`
	// What caused the deployment.
	Type ProjectEditResponseCanonicalDeploymentDeploymentTriggerType `json:"type,required"`
	JSON projectEditResponseCanonicalDeploymentDeploymentTriggerJSON `json:"-"`
}

// projectEditResponseCanonicalDeploymentDeploymentTriggerJSON contains the JSON
// metadata for the struct
// [ProjectEditResponseCanonicalDeploymentDeploymentTrigger]
type projectEditResponseCanonicalDeploymentDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseCanonicalDeploymentDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseCanonicalDeploymentDeploymentTriggerJSON) RawJSON() string {
	return r.raw
}

// Additional info about the trigger.
type ProjectEditResponseCanonicalDeploymentDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch,required"`
	// Whether the deployment trigger commit was dirty.
	CommitDirty bool `json:"commit_dirty,required"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash,required"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                              `json:"commit_message,required"`
	JSON          projectEditResponseCanonicalDeploymentDeploymentTriggerMetadataJSON `json:"-"`
}

// projectEditResponseCanonicalDeploymentDeploymentTriggerMetadataJSON contains the
// JSON metadata for the struct
// [ProjectEditResponseCanonicalDeploymentDeploymentTriggerMetadata]
type projectEditResponseCanonicalDeploymentDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitDirty   apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectEditResponseCanonicalDeploymentDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseCanonicalDeploymentDeploymentTriggerMetadataJSON) RawJSON() string {
	return r.raw
}

// What caused the deployment.
type ProjectEditResponseCanonicalDeploymentDeploymentTriggerType string

const (
	ProjectEditResponseCanonicalDeploymentDeploymentTriggerTypeGitHubPush ProjectEditResponseCanonicalDeploymentDeploymentTriggerType = "github:push"
	ProjectEditResponseCanonicalDeploymentDeploymentTriggerTypeADHoc      ProjectEditResponseCanonicalDeploymentDeploymentTriggerType = "ad_hoc"
	ProjectEditResponseCanonicalDeploymentDeploymentTriggerTypeDeployHook ProjectEditResponseCanonicalDeploymentDeploymentTriggerType = "deploy_hook"
)

func (r ProjectEditResponseCanonicalDeploymentDeploymentTriggerType) IsKnown() bool {
	switch r {
	case ProjectEditResponseCanonicalDeploymentDeploymentTriggerTypeGitHubPush, ProjectEditResponseCanonicalDeploymentDeploymentTriggerTypeADHoc, ProjectEditResponseCanonicalDeploymentDeploymentTriggerTypeDeployHook:
		return true
	}
	return false
}

// A plaintext environment variable.
type ProjectEditResponseCanonicalDeploymentEnvVar struct {
	Type ProjectEditResponseCanonicalDeploymentEnvVarsType `json:"type,required"`
	// Environment variable value.
	Value string                                           `json:"value,required"`
	JSON  projectEditResponseCanonicalDeploymentEnvVarJSON `json:"-"`
	union ProjectEditResponseCanonicalDeploymentEnvVarsUnion
}

// projectEditResponseCanonicalDeploymentEnvVarJSON contains the JSON metadata for
// the struct [ProjectEditResponseCanonicalDeploymentEnvVar]
type projectEditResponseCanonicalDeploymentEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectEditResponseCanonicalDeploymentEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectEditResponseCanonicalDeploymentEnvVar) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectEditResponseCanonicalDeploymentEnvVar{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectEditResponseCanonicalDeploymentEnvVarsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectEditResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVar],
// [ProjectEditResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVar].
func (r ProjectEditResponseCanonicalDeploymentEnvVar) AsUnion() ProjectEditResponseCanonicalDeploymentEnvVarsUnion {
	return r.union
}

// A plaintext environment variable.
//
// Union satisfied by
// [ProjectEditResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVar] or
// [ProjectEditResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVar].
type ProjectEditResponseCanonicalDeploymentEnvVarsUnion interface {
	implementsProjectEditResponseCanonicalDeploymentEnvVar()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectEditResponseCanonicalDeploymentEnvVarsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectEditResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVar{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectEditResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVar{}),
			DiscriminatorValue: "secret_text",
		},
	)
}

// A plaintext environment variable.
type ProjectEditResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVar struct {
	Type ProjectEditResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarType `json:"type,required"`
	// Environment variable value.
	Value string                                                                `json:"value,required"`
	JSON  projectEditResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarJSON `json:"-"`
}

// projectEditResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarJSON contains
// the JSON metadata for the struct
// [ProjectEditResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVar]
type projectEditResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectEditResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVar) implementsProjectEditResponseCanonicalDeploymentEnvVar() {
}

type ProjectEditResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarType string

const (
	ProjectEditResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarTypePlainText ProjectEditResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r ProjectEditResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectEditResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type ProjectEditResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVar struct {
	Type ProjectEditResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarType `json:"type,required"`
	// Secret value.
	Value string                                                                 `json:"value,required"`
	JSON  projectEditResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarJSON `json:"-"`
}

// projectEditResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarJSON contains
// the JSON metadata for the struct
// [ProjectEditResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVar]
type projectEditResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectEditResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVar) implementsProjectEditResponseCanonicalDeploymentEnvVar() {
}

type ProjectEditResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarType string

const (
	ProjectEditResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarTypeSecretText ProjectEditResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r ProjectEditResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectEditResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type ProjectEditResponseCanonicalDeploymentEnvVarsType string

const (
	ProjectEditResponseCanonicalDeploymentEnvVarsTypePlainText  ProjectEditResponseCanonicalDeploymentEnvVarsType = "plain_text"
	ProjectEditResponseCanonicalDeploymentEnvVarsTypeSecretText ProjectEditResponseCanonicalDeploymentEnvVarsType = "secret_text"
)

func (r ProjectEditResponseCanonicalDeploymentEnvVarsType) IsKnown() bool {
	switch r {
	case ProjectEditResponseCanonicalDeploymentEnvVarsTypePlainText, ProjectEditResponseCanonicalDeploymentEnvVarsTypeSecretText:
		return true
	}
	return false
}

// Type of deploy.
type ProjectEditResponseCanonicalDeploymentEnvironment string

const (
	ProjectEditResponseCanonicalDeploymentEnvironmentPreview    ProjectEditResponseCanonicalDeploymentEnvironment = "preview"
	ProjectEditResponseCanonicalDeploymentEnvironmentProduction ProjectEditResponseCanonicalDeploymentEnvironment = "production"
)

func (r ProjectEditResponseCanonicalDeploymentEnvironment) IsKnown() bool {
	switch r {
	case ProjectEditResponseCanonicalDeploymentEnvironmentPreview, ProjectEditResponseCanonicalDeploymentEnvironmentProduction:
		return true
	}
	return false
}

// Configs for the project source control.
type ProjectEditResponseCanonicalDeploymentSource struct {
	Config ProjectEditResponseCanonicalDeploymentSourceConfig `json:"config,required"`
	// The source control management provider.
	Type ProjectEditResponseCanonicalDeploymentSourceType `json:"type,required"`
	JSON projectEditResponseCanonicalDeploymentSourceJSON `json:"-"`
}

// projectEditResponseCanonicalDeploymentSourceJSON contains the JSON metadata for
// the struct [ProjectEditResponseCanonicalDeploymentSource]
type projectEditResponseCanonicalDeploymentSourceJSON struct {
	Config      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseCanonicalDeploymentSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseCanonicalDeploymentSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectEditResponseCanonicalDeploymentSourceConfig struct {
	// The owner of the repository.
	Owner string `json:"owner,required"`
	// Whether to enable PR comments.
	PrCommentsEnabled bool `json:"pr_comments_enabled,required"`
	// The production branch of the repository.
	ProductionBranch string `json:"production_branch,required"`
	// The name of the repository.
	RepoName string `json:"repo_name,required"`
	// Whether to enable automatic deployments when pushing to the source repository.
	// When disabled, no deployments (production or preview) will be triggered
	// automatically.
	//
	// Deprecated: Use `production_deployments_enabled` and
	// `preview_deployment_setting` for more granular control.
	DeploymentsEnabled bool `json:"deployments_enabled"`
	// The owner ID of the repository.
	OwnerID string `json:"owner_id"`
	// A list of paths that should be excluded from triggering a preview deployment.
	// Wildcard syntax (`*`) is supported.
	PathExcludes []string `json:"path_excludes"`
	// A list of paths that should be watched to trigger a preview deployment. Wildcard
	// syntax (`*`) is supported.
	PathIncludes []string `json:"path_includes"`
	// A list of branches that should not trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchExcludes []string `json:"preview_branch_excludes"`
	// A list of branches that should trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchIncludes []string `json:"preview_branch_includes"`
	// Controls whether commits to preview branches trigger a preview deployment.
	PreviewDeploymentSetting ProjectEditResponseCanonicalDeploymentSourceConfigPreviewDeploymentSetting `json:"preview_deployment_setting"`
	// Whether to trigger a production deployment on commits to the production branch.
	ProductionDeploymentsEnabled bool `json:"production_deployments_enabled"`
	// The ID of the repository.
	RepoID string                                                 `json:"repo_id"`
	JSON   projectEditResponseCanonicalDeploymentSourceConfigJSON `json:"-"`
}

// projectEditResponseCanonicalDeploymentSourceConfigJSON contains the JSON
// metadata for the struct [ProjectEditResponseCanonicalDeploymentSourceConfig]
type projectEditResponseCanonicalDeploymentSourceConfigJSON struct {
	Owner                        apijson.Field
	PrCommentsEnabled            apijson.Field
	ProductionBranch             apijson.Field
	RepoName                     apijson.Field
	DeploymentsEnabled           apijson.Field
	OwnerID                      apijson.Field
	PathExcludes                 apijson.Field
	PathIncludes                 apijson.Field
	PreviewBranchExcludes        apijson.Field
	PreviewBranchIncludes        apijson.Field
	PreviewDeploymentSetting     apijson.Field
	ProductionDeploymentsEnabled apijson.Field
	RepoID                       apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ProjectEditResponseCanonicalDeploymentSourceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseCanonicalDeploymentSourceConfigJSON) RawJSON() string {
	return r.raw
}

// Controls whether commits to preview branches trigger a preview deployment.
type ProjectEditResponseCanonicalDeploymentSourceConfigPreviewDeploymentSetting string

const (
	ProjectEditResponseCanonicalDeploymentSourceConfigPreviewDeploymentSettingAll    ProjectEditResponseCanonicalDeploymentSourceConfigPreviewDeploymentSetting = "all"
	ProjectEditResponseCanonicalDeploymentSourceConfigPreviewDeploymentSettingNone   ProjectEditResponseCanonicalDeploymentSourceConfigPreviewDeploymentSetting = "none"
	ProjectEditResponseCanonicalDeploymentSourceConfigPreviewDeploymentSettingCustom ProjectEditResponseCanonicalDeploymentSourceConfigPreviewDeploymentSetting = "custom"
)

func (r ProjectEditResponseCanonicalDeploymentSourceConfigPreviewDeploymentSetting) IsKnown() bool {
	switch r {
	case ProjectEditResponseCanonicalDeploymentSourceConfigPreviewDeploymentSettingAll, ProjectEditResponseCanonicalDeploymentSourceConfigPreviewDeploymentSettingNone, ProjectEditResponseCanonicalDeploymentSourceConfigPreviewDeploymentSettingCustom:
		return true
	}
	return false
}

// The source control management provider.
type ProjectEditResponseCanonicalDeploymentSourceType string

const (
	ProjectEditResponseCanonicalDeploymentSourceTypeGitHub ProjectEditResponseCanonicalDeploymentSourceType = "github"
	ProjectEditResponseCanonicalDeploymentSourceTypeGitlab ProjectEditResponseCanonicalDeploymentSourceType = "gitlab"
)

func (r ProjectEditResponseCanonicalDeploymentSourceType) IsKnown() bool {
	switch r {
	case ProjectEditResponseCanonicalDeploymentSourceTypeGitHub, ProjectEditResponseCanonicalDeploymentSourceTypeGitlab:
		return true
	}
	return false
}

// Configs for deployments in a project.
type ProjectEditResponseDeploymentConfigs struct {
	// Configs for preview deploys.
	Preview ProjectEditResponseDeploymentConfigsPreview `json:"preview,required"`
	// Configs for production deploys.
	Production ProjectEditResponseDeploymentConfigsProduction `json:"production,required"`
	JSON       projectEditResponseDeploymentConfigsJSON       `json:"-"`
}

// projectEditResponseDeploymentConfigsJSON contains the JSON metadata for the
// struct [ProjectEditResponseDeploymentConfigs]
type projectEditResponseDeploymentConfigsJSON struct {
	Preview     apijson.Field
	Production  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsJSON) RawJSON() string {
	return r.raw
}

// Configs for preview deploys.
type ProjectEditResponseDeploymentConfigsPreview struct {
	// Whether to always use the latest compatibility date for Pages Functions.
	AlwaysUseLatestCompatibilityDate bool `json:"always_use_latest_compatibility_date,required"`
	// The major version of the build image to use for Pages Functions.
	BuildImageMajorVersion int64 `json:"build_image_major_version,required"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate string `json:"compatibility_date,required"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags []string `json:"compatibility_flags,required"`
	// Environment variables used for builds and Pages Functions.
	EnvVars map[string]ProjectEditResponseDeploymentConfigsPreviewEnvVar `json:"env_vars,required,nullable"`
	// Whether to fail open when the deployment config cannot be applied.
	FailOpen bool `json:"fail_open,required"`
	// The usage model for Pages Functions.
	//
	// Deprecated: All new projects now use the Standard usage model.
	UsageModel ProjectEditResponseDeploymentConfigsPreviewUsageModel `json:"usage_model,required"`
	// Constellation bindings used for Pages Functions.
	AIBindings map[string]ProjectEditResponseDeploymentConfigsPreviewAIBinding `json:"ai_bindings"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets map[string]ProjectEditResponseDeploymentConfigsPreviewAnalyticsEngineDataset `json:"analytics_engine_datasets"`
	// Browser bindings used for Pages Functions.
	Browsers map[string]ProjectEditResponseDeploymentConfigsPreviewBrowser `json:"browsers"`
	// D1 databases used for Pages Functions.
	D1Databases map[string]ProjectEditResponseDeploymentConfigsPreviewD1Database `json:"d1_databases"`
	// Durable Object namespaces used for Pages Functions.
	DurableObjectNamespaces map[string]ProjectEditResponseDeploymentConfigsPreviewDurableObjectNamespace `json:"durable_object_namespaces"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings map[string]ProjectEditResponseDeploymentConfigsPreviewHyperdriveBinding `json:"hyperdrive_bindings"`
	// KV namespaces used for Pages Functions.
	KVNamespaces map[string]ProjectEditResponseDeploymentConfigsPreviewKVNamespace `json:"kv_namespaces"`
	// Limits for Pages Functions.
	Limits ProjectEditResponseDeploymentConfigsPreviewLimits `json:"limits"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates map[string]ProjectEditResponseDeploymentConfigsPreviewMTLSCertificate `json:"mtls_certificates"`
	// Placement setting used for Pages Functions.
	Placement ProjectEditResponseDeploymentConfigsPreviewPlacement `json:"placement"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers map[string]ProjectEditResponseDeploymentConfigsPreviewQueueProducer `json:"queue_producers"`
	// R2 buckets used for Pages Functions.
	R2Buckets map[string]ProjectEditResponseDeploymentConfigsPreviewR2Bucket `json:"r2_buckets"`
	// Services used for Pages Functions.
	Services map[string]ProjectEditResponseDeploymentConfigsPreviewService `json:"services"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings map[string]ProjectEditResponseDeploymentConfigsPreviewVectorizeBinding `json:"vectorize_bindings"`
	// Hash of the Wrangler configuration used for the deployment.
	WranglerConfigHash string                                          `json:"wrangler_config_hash"`
	JSON               projectEditResponseDeploymentConfigsPreviewJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewJSON contains the JSON metadata for
// the struct [ProjectEditResponseDeploymentConfigsPreview]
type projectEditResponseDeploymentConfigsPreviewJSON struct {
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

func (r *ProjectEditResponseDeploymentConfigsPreview) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewJSON) RawJSON() string {
	return r.raw
}

// A plaintext environment variable.
type ProjectEditResponseDeploymentConfigsPreviewEnvVar struct {
	Type ProjectEditResponseDeploymentConfigsPreviewEnvVarsType `json:"type,required"`
	// Environment variable value.
	Value string                                                `json:"value,required"`
	JSON  projectEditResponseDeploymentConfigsPreviewEnvVarJSON `json:"-"`
	union ProjectEditResponseDeploymentConfigsPreviewEnvVarsUnion
}

// projectEditResponseDeploymentConfigsPreviewEnvVarJSON contains the JSON metadata
// for the struct [ProjectEditResponseDeploymentConfigsPreviewEnvVar]
type projectEditResponseDeploymentConfigsPreviewEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectEditResponseDeploymentConfigsPreviewEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectEditResponseDeploymentConfigsPreviewEnvVar) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectEditResponseDeploymentConfigsPreviewEnvVar{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectEditResponseDeploymentConfigsPreviewEnvVarsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectEditResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar],
// [ProjectEditResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar].
func (r ProjectEditResponseDeploymentConfigsPreviewEnvVar) AsUnion() ProjectEditResponseDeploymentConfigsPreviewEnvVarsUnion {
	return r.union
}

// A plaintext environment variable.
//
// Union satisfied by
// [ProjectEditResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar] or
// [ProjectEditResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar].
type ProjectEditResponseDeploymentConfigsPreviewEnvVarsUnion interface {
	implementsProjectEditResponseDeploymentConfigsPreviewEnvVar()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectEditResponseDeploymentConfigsPreviewEnvVarsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectEditResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectEditResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar{}),
			DiscriminatorValue: "secret_text",
		},
	)
}

// A plaintext environment variable.
type ProjectEditResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar struct {
	Type ProjectEditResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType `json:"type,required"`
	// Environment variable value.
	Value string                                                                     `json:"value,required"`
	JSON  projectEditResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarJSON
// contains the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar]
type projectEditResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectEditResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar) implementsProjectEditResponseDeploymentConfigsPreviewEnvVar() {
}

type ProjectEditResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType string

const (
	ProjectEditResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarTypePlainText ProjectEditResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r ProjectEditResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectEditResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type ProjectEditResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar struct {
	Type ProjectEditResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType `json:"type,required"`
	// Secret value.
	Value string                                                                      `json:"value,required"`
	JSON  projectEditResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarJSON
// contains the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar]
type projectEditResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectEditResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar) implementsProjectEditResponseDeploymentConfigsPreviewEnvVar() {
}

type ProjectEditResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType string

const (
	ProjectEditResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarTypeSecretText ProjectEditResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r ProjectEditResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectEditResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type ProjectEditResponseDeploymentConfigsPreviewEnvVarsType string

const (
	ProjectEditResponseDeploymentConfigsPreviewEnvVarsTypePlainText  ProjectEditResponseDeploymentConfigsPreviewEnvVarsType = "plain_text"
	ProjectEditResponseDeploymentConfigsPreviewEnvVarsTypeSecretText ProjectEditResponseDeploymentConfigsPreviewEnvVarsType = "secret_text"
)

func (r ProjectEditResponseDeploymentConfigsPreviewEnvVarsType) IsKnown() bool {
	switch r {
	case ProjectEditResponseDeploymentConfigsPreviewEnvVarsTypePlainText, ProjectEditResponseDeploymentConfigsPreviewEnvVarsTypeSecretText:
		return true
	}
	return false
}

// The usage model for Pages Functions.
type ProjectEditResponseDeploymentConfigsPreviewUsageModel string

const (
	ProjectEditResponseDeploymentConfigsPreviewUsageModelStandard ProjectEditResponseDeploymentConfigsPreviewUsageModel = "standard"
	ProjectEditResponseDeploymentConfigsPreviewUsageModelBundled  ProjectEditResponseDeploymentConfigsPreviewUsageModel = "bundled"
	ProjectEditResponseDeploymentConfigsPreviewUsageModelUnbound  ProjectEditResponseDeploymentConfigsPreviewUsageModel = "unbound"
)

func (r ProjectEditResponseDeploymentConfigsPreviewUsageModel) IsKnown() bool {
	switch r {
	case ProjectEditResponseDeploymentConfigsPreviewUsageModelStandard, ProjectEditResponseDeploymentConfigsPreviewUsageModelBundled, ProjectEditResponseDeploymentConfigsPreviewUsageModelUnbound:
		return true
	}
	return false
}

// AI binding.
type ProjectEditResponseDeploymentConfigsPreviewAIBinding struct {
	ProjectID string                                                   `json:"project_id,required"`
	JSON      projectEditResponseDeploymentConfigsPreviewAIBindingJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewAIBindingJSON contains the JSON
// metadata for the struct [ProjectEditResponseDeploymentConfigsPreviewAIBinding]
type projectEditResponseDeploymentConfigsPreviewAIBindingJSON struct {
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewAIBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewAIBindingJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine binding.
type ProjectEditResponseDeploymentConfigsPreviewAnalyticsEngineDataset struct {
	// Name of the dataset.
	Dataset string                                                                `json:"dataset,required"`
	JSON    projectEditResponseDeploymentConfigsPreviewAnalyticsEngineDatasetJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewAnalyticsEngineDatasetJSON contains
// the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsPreviewAnalyticsEngineDataset]
type projectEditResponseDeploymentConfigsPreviewAnalyticsEngineDatasetJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewAnalyticsEngineDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewAnalyticsEngineDatasetJSON) RawJSON() string {
	return r.raw
}

// Browser binding.
type ProjectEditResponseDeploymentConfigsPreviewBrowser struct {
	JSON projectEditResponseDeploymentConfigsPreviewBrowserJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewBrowserJSON contains the JSON
// metadata for the struct [ProjectEditResponseDeploymentConfigsPreviewBrowser]
type projectEditResponseDeploymentConfigsPreviewBrowserJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewBrowser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewBrowserJSON) RawJSON() string {
	return r.raw
}

// D1 binding.
type ProjectEditResponseDeploymentConfigsPreviewD1Database struct {
	// UUID of the D1 database.
	ID   string                                                    `json:"id,required"`
	JSON projectEditResponseDeploymentConfigsPreviewD1DatabaseJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewD1DatabaseJSON contains the JSON
// metadata for the struct [ProjectEditResponseDeploymentConfigsPreviewD1Database]
type projectEditResponseDeploymentConfigsPreviewD1DatabaseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewD1Database) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewD1DatabaseJSON) RawJSON() string {
	return r.raw
}

// Durable Object binding.
type ProjectEditResponseDeploymentConfigsPreviewDurableObjectNamespace struct {
	// ID of the Durable Object namespace.
	NamespaceID string                                                                `json:"namespace_id,required"`
	JSON        projectEditResponseDeploymentConfigsPreviewDurableObjectNamespaceJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewDurableObjectNamespaceJSON contains
// the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsPreviewDurableObjectNamespace]
type projectEditResponseDeploymentConfigsPreviewDurableObjectNamespaceJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

// Hyperdrive binding.
type ProjectEditResponseDeploymentConfigsPreviewHyperdriveBinding struct {
	ID   string                                                           `json:"id,required"`
	JSON projectEditResponseDeploymentConfigsPreviewHyperdriveBindingJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewHyperdriveBindingJSON contains the
// JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsPreviewHyperdriveBinding]
type projectEditResponseDeploymentConfigsPreviewHyperdriveBindingJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewHyperdriveBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewHyperdriveBindingJSON) RawJSON() string {
	return r.raw
}

// KV namespace binding.
type ProjectEditResponseDeploymentConfigsPreviewKVNamespace struct {
	// ID of the KV namespace.
	NamespaceID string                                                     `json:"namespace_id,required"`
	JSON        projectEditResponseDeploymentConfigsPreviewKVNamespaceJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewKVNamespaceJSON contains the JSON
// metadata for the struct [ProjectEditResponseDeploymentConfigsPreviewKVNamespace]
type projectEditResponseDeploymentConfigsPreviewKVNamespaceJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewKVNamespaceJSON) RawJSON() string {
	return r.raw
}

// Limits for Pages Functions.
type ProjectEditResponseDeploymentConfigsPreviewLimits struct {
	// CPU time limit in milliseconds.
	CPUMs int64                                                 `json:"cpu_ms,required"`
	JSON  projectEditResponseDeploymentConfigsPreviewLimitsJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewLimitsJSON contains the JSON metadata
// for the struct [ProjectEditResponseDeploymentConfigsPreviewLimits]
type projectEditResponseDeploymentConfigsPreviewLimitsJSON struct {
	CPUMs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewLimitsJSON) RawJSON() string {
	return r.raw
}

// mTLS binding.
type ProjectEditResponseDeploymentConfigsPreviewMTLSCertificate struct {
	CertificateID string                                                         `json:"certificate_id,required"`
	JSON          projectEditResponseDeploymentConfigsPreviewMTLSCertificateJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewMTLSCertificateJSON contains the JSON
// metadata for the struct
// [ProjectEditResponseDeploymentConfigsPreviewMTLSCertificate]
type projectEditResponseDeploymentConfigsPreviewMTLSCertificateJSON struct {
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewMTLSCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewMTLSCertificateJSON) RawJSON() string {
	return r.raw
}

// Placement setting used for Pages Functions.
type ProjectEditResponseDeploymentConfigsPreviewPlacement struct {
	// Placement mode.
	Mode string                                                   `json:"mode,required"`
	JSON projectEditResponseDeploymentConfigsPreviewPlacementJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewPlacementJSON contains the JSON
// metadata for the struct [ProjectEditResponseDeploymentConfigsPreviewPlacement]
type projectEditResponseDeploymentConfigsPreviewPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewPlacementJSON) RawJSON() string {
	return r.raw
}

// Queue Producer binding.
type ProjectEditResponseDeploymentConfigsPreviewQueueProducer struct {
	// Name of the Queue.
	Name string                                                       `json:"name,required"`
	JSON projectEditResponseDeploymentConfigsPreviewQueueProducerJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewQueueProducerJSON contains the JSON
// metadata for the struct
// [ProjectEditResponseDeploymentConfigsPreviewQueueProducer]
type projectEditResponseDeploymentConfigsPreviewQueueProducerJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewQueueProducer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewQueueProducerJSON) RawJSON() string {
	return r.raw
}

// R2 binding.
type ProjectEditResponseDeploymentConfigsPreviewR2Bucket struct {
	// Name of the R2 bucket.
	Name string `json:"name,required"`
	// Jurisdiction of the R2 bucket.
	Jurisdiction string                                                  `json:"jurisdiction,nullable"`
	JSON         projectEditResponseDeploymentConfigsPreviewR2BucketJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewR2BucketJSON contains the JSON
// metadata for the struct [ProjectEditResponseDeploymentConfigsPreviewR2Bucket]
type projectEditResponseDeploymentConfigsPreviewR2BucketJSON struct {
	Name         apijson.Field
	Jurisdiction apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewR2BucketJSON) RawJSON() string {
	return r.raw
}

// Service binding.
type ProjectEditResponseDeploymentConfigsPreviewService struct {
	// The Service environment.
	Environment string `json:"environment,required"`
	// The Service name.
	Service string `json:"service,required"`
	// The entrypoint to bind to.
	Entrypoint string                                                 `json:"entrypoint,nullable"`
	JSON       projectEditResponseDeploymentConfigsPreviewServiceJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewServiceJSON contains the JSON
// metadata for the struct [ProjectEditResponseDeploymentConfigsPreviewService]
type projectEditResponseDeploymentConfigsPreviewServiceJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	Entrypoint  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewServiceJSON) RawJSON() string {
	return r.raw
}

// Vectorize binding.
type ProjectEditResponseDeploymentConfigsPreviewVectorizeBinding struct {
	IndexName string                                                          `json:"index_name,required"`
	JSON      projectEditResponseDeploymentConfigsPreviewVectorizeBindingJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsPreviewVectorizeBindingJSON contains the
// JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsPreviewVectorizeBinding]
type projectEditResponseDeploymentConfigsPreviewVectorizeBindingJSON struct {
	IndexName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsPreviewVectorizeBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsPreviewVectorizeBindingJSON) RawJSON() string {
	return r.raw
}

// Configs for production deploys.
type ProjectEditResponseDeploymentConfigsProduction struct {
	// Whether to always use the latest compatibility date for Pages Functions.
	AlwaysUseLatestCompatibilityDate bool `json:"always_use_latest_compatibility_date,required"`
	// The major version of the build image to use for Pages Functions.
	BuildImageMajorVersion int64 `json:"build_image_major_version,required"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate string `json:"compatibility_date,required"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags []string `json:"compatibility_flags,required"`
	// Environment variables used for builds and Pages Functions.
	EnvVars map[string]ProjectEditResponseDeploymentConfigsProductionEnvVar `json:"env_vars,required,nullable"`
	// Whether to fail open when the deployment config cannot be applied.
	FailOpen bool `json:"fail_open,required"`
	// The usage model for Pages Functions.
	//
	// Deprecated: All new projects now use the Standard usage model.
	UsageModel ProjectEditResponseDeploymentConfigsProductionUsageModel `json:"usage_model,required"`
	// Constellation bindings used for Pages Functions.
	AIBindings map[string]ProjectEditResponseDeploymentConfigsProductionAIBinding `json:"ai_bindings"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets map[string]ProjectEditResponseDeploymentConfigsProductionAnalyticsEngineDataset `json:"analytics_engine_datasets"`
	// Browser bindings used for Pages Functions.
	Browsers map[string]ProjectEditResponseDeploymentConfigsProductionBrowser `json:"browsers"`
	// D1 databases used for Pages Functions.
	D1Databases map[string]ProjectEditResponseDeploymentConfigsProductionD1Database `json:"d1_databases"`
	// Durable Object namespaces used for Pages Functions.
	DurableObjectNamespaces map[string]ProjectEditResponseDeploymentConfigsProductionDurableObjectNamespace `json:"durable_object_namespaces"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings map[string]ProjectEditResponseDeploymentConfigsProductionHyperdriveBinding `json:"hyperdrive_bindings"`
	// KV namespaces used for Pages Functions.
	KVNamespaces map[string]ProjectEditResponseDeploymentConfigsProductionKVNamespace `json:"kv_namespaces"`
	// Limits for Pages Functions.
	Limits ProjectEditResponseDeploymentConfigsProductionLimits `json:"limits"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates map[string]ProjectEditResponseDeploymentConfigsProductionMTLSCertificate `json:"mtls_certificates"`
	// Placement setting used for Pages Functions.
	Placement ProjectEditResponseDeploymentConfigsProductionPlacement `json:"placement"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers map[string]ProjectEditResponseDeploymentConfigsProductionQueueProducer `json:"queue_producers"`
	// R2 buckets used for Pages Functions.
	R2Buckets map[string]ProjectEditResponseDeploymentConfigsProductionR2Bucket `json:"r2_buckets"`
	// Services used for Pages Functions.
	Services map[string]ProjectEditResponseDeploymentConfigsProductionService `json:"services"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings map[string]ProjectEditResponseDeploymentConfigsProductionVectorizeBinding `json:"vectorize_bindings"`
	// Hash of the Wrangler configuration used for the deployment.
	WranglerConfigHash string                                             `json:"wrangler_config_hash"`
	JSON               projectEditResponseDeploymentConfigsProductionJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionJSON contains the JSON metadata
// for the struct [ProjectEditResponseDeploymentConfigsProduction]
type projectEditResponseDeploymentConfigsProductionJSON struct {
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

func (r *ProjectEditResponseDeploymentConfigsProduction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionJSON) RawJSON() string {
	return r.raw
}

// A plaintext environment variable.
type ProjectEditResponseDeploymentConfigsProductionEnvVar struct {
	Type ProjectEditResponseDeploymentConfigsProductionEnvVarsType `json:"type,required"`
	// Environment variable value.
	Value string                                                   `json:"value,required"`
	JSON  projectEditResponseDeploymentConfigsProductionEnvVarJSON `json:"-"`
	union ProjectEditResponseDeploymentConfigsProductionEnvVarsUnion
}

// projectEditResponseDeploymentConfigsProductionEnvVarJSON contains the JSON
// metadata for the struct [ProjectEditResponseDeploymentConfigsProductionEnvVar]
type projectEditResponseDeploymentConfigsProductionEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectEditResponseDeploymentConfigsProductionEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectEditResponseDeploymentConfigsProductionEnvVar) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectEditResponseDeploymentConfigsProductionEnvVar{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectEditResponseDeploymentConfigsProductionEnvVarsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectEditResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar],
// [ProjectEditResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar].
func (r ProjectEditResponseDeploymentConfigsProductionEnvVar) AsUnion() ProjectEditResponseDeploymentConfigsProductionEnvVarsUnion {
	return r.union
}

// A plaintext environment variable.
//
// Union satisfied by
// [ProjectEditResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar] or
// [ProjectEditResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar].
type ProjectEditResponseDeploymentConfigsProductionEnvVarsUnion interface {
	implementsProjectEditResponseDeploymentConfigsProductionEnvVar()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectEditResponseDeploymentConfigsProductionEnvVarsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectEditResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectEditResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar{}),
			DiscriminatorValue: "secret_text",
		},
	)
}

// A plaintext environment variable.
type ProjectEditResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar struct {
	Type ProjectEditResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType `json:"type,required"`
	// Environment variable value.
	Value string                                                                        `json:"value,required"`
	JSON  projectEditResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarJSON
// contains the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar]
type projectEditResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectEditResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar) implementsProjectEditResponseDeploymentConfigsProductionEnvVar() {
}

type ProjectEditResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType string

const (
	ProjectEditResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarTypePlainText ProjectEditResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r ProjectEditResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectEditResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type ProjectEditResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar struct {
	Type ProjectEditResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType `json:"type,required"`
	// Secret value.
	Value string                                                                         `json:"value,required"`
	JSON  projectEditResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarJSON
// contains the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar]
type projectEditResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectEditResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar) implementsProjectEditResponseDeploymentConfigsProductionEnvVar() {
}

type ProjectEditResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType string

const (
	ProjectEditResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarTypeSecretText ProjectEditResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r ProjectEditResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectEditResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type ProjectEditResponseDeploymentConfigsProductionEnvVarsType string

const (
	ProjectEditResponseDeploymentConfigsProductionEnvVarsTypePlainText  ProjectEditResponseDeploymentConfigsProductionEnvVarsType = "plain_text"
	ProjectEditResponseDeploymentConfigsProductionEnvVarsTypeSecretText ProjectEditResponseDeploymentConfigsProductionEnvVarsType = "secret_text"
)

func (r ProjectEditResponseDeploymentConfigsProductionEnvVarsType) IsKnown() bool {
	switch r {
	case ProjectEditResponseDeploymentConfigsProductionEnvVarsTypePlainText, ProjectEditResponseDeploymentConfigsProductionEnvVarsTypeSecretText:
		return true
	}
	return false
}

// The usage model for Pages Functions.
type ProjectEditResponseDeploymentConfigsProductionUsageModel string

const (
	ProjectEditResponseDeploymentConfigsProductionUsageModelStandard ProjectEditResponseDeploymentConfigsProductionUsageModel = "standard"
	ProjectEditResponseDeploymentConfigsProductionUsageModelBundled  ProjectEditResponseDeploymentConfigsProductionUsageModel = "bundled"
	ProjectEditResponseDeploymentConfigsProductionUsageModelUnbound  ProjectEditResponseDeploymentConfigsProductionUsageModel = "unbound"
)

func (r ProjectEditResponseDeploymentConfigsProductionUsageModel) IsKnown() bool {
	switch r {
	case ProjectEditResponseDeploymentConfigsProductionUsageModelStandard, ProjectEditResponseDeploymentConfigsProductionUsageModelBundled, ProjectEditResponseDeploymentConfigsProductionUsageModelUnbound:
		return true
	}
	return false
}

// AI binding.
type ProjectEditResponseDeploymentConfigsProductionAIBinding struct {
	ProjectID string                                                      `json:"project_id,required"`
	JSON      projectEditResponseDeploymentConfigsProductionAIBindingJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionAIBindingJSON contains the JSON
// metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionAIBinding]
type projectEditResponseDeploymentConfigsProductionAIBindingJSON struct {
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionAIBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionAIBindingJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine binding.
type ProjectEditResponseDeploymentConfigsProductionAnalyticsEngineDataset struct {
	// Name of the dataset.
	Dataset string                                                                   `json:"dataset,required"`
	JSON    projectEditResponseDeploymentConfigsProductionAnalyticsEngineDatasetJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionAnalyticsEngineDatasetJSON
// contains the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionAnalyticsEngineDataset]
type projectEditResponseDeploymentConfigsProductionAnalyticsEngineDatasetJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionAnalyticsEngineDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionAnalyticsEngineDatasetJSON) RawJSON() string {
	return r.raw
}

// Browser binding.
type ProjectEditResponseDeploymentConfigsProductionBrowser struct {
	JSON projectEditResponseDeploymentConfigsProductionBrowserJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionBrowserJSON contains the JSON
// metadata for the struct [ProjectEditResponseDeploymentConfigsProductionBrowser]
type projectEditResponseDeploymentConfigsProductionBrowserJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionBrowser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionBrowserJSON) RawJSON() string {
	return r.raw
}

// D1 binding.
type ProjectEditResponseDeploymentConfigsProductionD1Database struct {
	// UUID of the D1 database.
	ID   string                                                       `json:"id,required"`
	JSON projectEditResponseDeploymentConfigsProductionD1DatabaseJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionD1DatabaseJSON contains the JSON
// metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionD1Database]
type projectEditResponseDeploymentConfigsProductionD1DatabaseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionD1Database) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionD1DatabaseJSON) RawJSON() string {
	return r.raw
}

// Durable Object binding.
type ProjectEditResponseDeploymentConfigsProductionDurableObjectNamespace struct {
	// ID of the Durable Object namespace.
	NamespaceID string                                                                   `json:"namespace_id,required"`
	JSON        projectEditResponseDeploymentConfigsProductionDurableObjectNamespaceJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionDurableObjectNamespaceJSON
// contains the JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionDurableObjectNamespace]
type projectEditResponseDeploymentConfigsProductionDurableObjectNamespaceJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

// Hyperdrive binding.
type ProjectEditResponseDeploymentConfigsProductionHyperdriveBinding struct {
	ID   string                                                              `json:"id,required"`
	JSON projectEditResponseDeploymentConfigsProductionHyperdriveBindingJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionHyperdriveBindingJSON contains the
// JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionHyperdriveBinding]
type projectEditResponseDeploymentConfigsProductionHyperdriveBindingJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionHyperdriveBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionHyperdriveBindingJSON) RawJSON() string {
	return r.raw
}

// KV namespace binding.
type ProjectEditResponseDeploymentConfigsProductionKVNamespace struct {
	// ID of the KV namespace.
	NamespaceID string                                                        `json:"namespace_id,required"`
	JSON        projectEditResponseDeploymentConfigsProductionKVNamespaceJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionKVNamespaceJSON contains the JSON
// metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionKVNamespace]
type projectEditResponseDeploymentConfigsProductionKVNamespaceJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionKVNamespaceJSON) RawJSON() string {
	return r.raw
}

// Limits for Pages Functions.
type ProjectEditResponseDeploymentConfigsProductionLimits struct {
	// CPU time limit in milliseconds.
	CPUMs int64                                                    `json:"cpu_ms,required"`
	JSON  projectEditResponseDeploymentConfigsProductionLimitsJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionLimitsJSON contains the JSON
// metadata for the struct [ProjectEditResponseDeploymentConfigsProductionLimits]
type projectEditResponseDeploymentConfigsProductionLimitsJSON struct {
	CPUMs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionLimitsJSON) RawJSON() string {
	return r.raw
}

// mTLS binding.
type ProjectEditResponseDeploymentConfigsProductionMTLSCertificate struct {
	CertificateID string                                                            `json:"certificate_id,required"`
	JSON          projectEditResponseDeploymentConfigsProductionMTLSCertificateJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionMTLSCertificateJSON contains the
// JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionMTLSCertificate]
type projectEditResponseDeploymentConfigsProductionMTLSCertificateJSON struct {
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionMTLSCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionMTLSCertificateJSON) RawJSON() string {
	return r.raw
}

// Placement setting used for Pages Functions.
type ProjectEditResponseDeploymentConfigsProductionPlacement struct {
	// Placement mode.
	Mode string                                                      `json:"mode,required"`
	JSON projectEditResponseDeploymentConfigsProductionPlacementJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionPlacementJSON contains the JSON
// metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionPlacement]
type projectEditResponseDeploymentConfigsProductionPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionPlacementJSON) RawJSON() string {
	return r.raw
}

// Queue Producer binding.
type ProjectEditResponseDeploymentConfigsProductionQueueProducer struct {
	// Name of the Queue.
	Name string                                                          `json:"name,required"`
	JSON projectEditResponseDeploymentConfigsProductionQueueProducerJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionQueueProducerJSON contains the
// JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionQueueProducer]
type projectEditResponseDeploymentConfigsProductionQueueProducerJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionQueueProducer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionQueueProducerJSON) RawJSON() string {
	return r.raw
}

// R2 binding.
type ProjectEditResponseDeploymentConfigsProductionR2Bucket struct {
	// Name of the R2 bucket.
	Name string `json:"name,required"`
	// Jurisdiction of the R2 bucket.
	Jurisdiction string                                                     `json:"jurisdiction,nullable"`
	JSON         projectEditResponseDeploymentConfigsProductionR2BucketJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionR2BucketJSON contains the JSON
// metadata for the struct [ProjectEditResponseDeploymentConfigsProductionR2Bucket]
type projectEditResponseDeploymentConfigsProductionR2BucketJSON struct {
	Name         apijson.Field
	Jurisdiction apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionR2BucketJSON) RawJSON() string {
	return r.raw
}

// Service binding.
type ProjectEditResponseDeploymentConfigsProductionService struct {
	// The Service environment.
	Environment string `json:"environment,required"`
	// The Service name.
	Service string `json:"service,required"`
	// The entrypoint to bind to.
	Entrypoint string                                                    `json:"entrypoint,nullable"`
	JSON       projectEditResponseDeploymentConfigsProductionServiceJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionServiceJSON contains the JSON
// metadata for the struct [ProjectEditResponseDeploymentConfigsProductionService]
type projectEditResponseDeploymentConfigsProductionServiceJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	Entrypoint  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionServiceJSON) RawJSON() string {
	return r.raw
}

// Vectorize binding.
type ProjectEditResponseDeploymentConfigsProductionVectorizeBinding struct {
	IndexName string                                                             `json:"index_name,required"`
	JSON      projectEditResponseDeploymentConfigsProductionVectorizeBindingJSON `json:"-"`
}

// projectEditResponseDeploymentConfigsProductionVectorizeBindingJSON contains the
// JSON metadata for the struct
// [ProjectEditResponseDeploymentConfigsProductionVectorizeBinding]
type projectEditResponseDeploymentConfigsProductionVectorizeBindingJSON struct {
	IndexName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseDeploymentConfigsProductionVectorizeBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseDeploymentConfigsProductionVectorizeBindingJSON) RawJSON() string {
	return r.raw
}

// Most recent deployment of the project.
type ProjectEditResponseLatestDeployment struct {
	// Id of the deployment.
	ID string `json:"id,required"`
	// A list of alias URLs pointing to this deployment.
	Aliases []string `json:"aliases,required,nullable"`
	// Configs for the project build process.
	BuildConfig ProjectEditResponseLatestDeploymentBuildConfig `json:"build_config,required"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger ProjectEditResponseLatestDeploymentDeploymentTrigger `json:"deployment_trigger,required"`
	// Environment variables used for builds and Pages Functions.
	EnvVars map[string]ProjectEditResponseLatestDeploymentEnvVar `json:"env_vars,required,nullable"`
	// Type of deploy.
	Environment ProjectEditResponseLatestDeploymentEnvironment `json:"environment,required"`
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
	Source ProjectEditResponseLatestDeploymentSource `json:"source,required"`
	// List of past stages.
	Stages []Stage `json:"stages,required"`
	// The live URL to view this deployment.
	URL string `json:"url,required"`
	// Whether the deployment uses functions.
	UsesFunctions bool                                    `json:"uses_functions,nullable"`
	JSON          projectEditResponseLatestDeploymentJSON `json:"-"`
}

// projectEditResponseLatestDeploymentJSON contains the JSON metadata for the
// struct [ProjectEditResponseLatestDeployment]
type projectEditResponseLatestDeploymentJSON struct {
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

func (r *ProjectEditResponseLatestDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseLatestDeploymentJSON) RawJSON() string {
	return r.raw
}

// Configs for the project build process.
type ProjectEditResponseLatestDeploymentBuildConfig struct {
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
	RootDir string                                             `json:"root_dir,nullable"`
	JSON    projectEditResponseLatestDeploymentBuildConfigJSON `json:"-"`
}

// projectEditResponseLatestDeploymentBuildConfigJSON contains the JSON metadata
// for the struct [ProjectEditResponseLatestDeploymentBuildConfig]
type projectEditResponseLatestDeploymentBuildConfigJSON struct {
	WebAnalyticsTag   apijson.Field
	WebAnalyticsToken apijson.Field
	BuildCaching      apijson.Field
	BuildCommand      apijson.Field
	DestinationDir    apijson.Field
	RootDir           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectEditResponseLatestDeploymentBuildConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseLatestDeploymentBuildConfigJSON) RawJSON() string {
	return r.raw
}

// Info about what caused the deployment.
type ProjectEditResponseLatestDeploymentDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata ProjectEditResponseLatestDeploymentDeploymentTriggerMetadata `json:"metadata,required"`
	// What caused the deployment.
	Type ProjectEditResponseLatestDeploymentDeploymentTriggerType `json:"type,required"`
	JSON projectEditResponseLatestDeploymentDeploymentTriggerJSON `json:"-"`
}

// projectEditResponseLatestDeploymentDeploymentTriggerJSON contains the JSON
// metadata for the struct [ProjectEditResponseLatestDeploymentDeploymentTrigger]
type projectEditResponseLatestDeploymentDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseLatestDeploymentDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseLatestDeploymentDeploymentTriggerJSON) RawJSON() string {
	return r.raw
}

// Additional info about the trigger.
type ProjectEditResponseLatestDeploymentDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch,required"`
	// Whether the deployment trigger commit was dirty.
	CommitDirty bool `json:"commit_dirty,required"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash,required"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                           `json:"commit_message,required"`
	JSON          projectEditResponseLatestDeploymentDeploymentTriggerMetadataJSON `json:"-"`
}

// projectEditResponseLatestDeploymentDeploymentTriggerMetadataJSON contains the
// JSON metadata for the struct
// [ProjectEditResponseLatestDeploymentDeploymentTriggerMetadata]
type projectEditResponseLatestDeploymentDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitDirty   apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectEditResponseLatestDeploymentDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseLatestDeploymentDeploymentTriggerMetadataJSON) RawJSON() string {
	return r.raw
}

// What caused the deployment.
type ProjectEditResponseLatestDeploymentDeploymentTriggerType string

const (
	ProjectEditResponseLatestDeploymentDeploymentTriggerTypeGitHubPush ProjectEditResponseLatestDeploymentDeploymentTriggerType = "github:push"
	ProjectEditResponseLatestDeploymentDeploymentTriggerTypeADHoc      ProjectEditResponseLatestDeploymentDeploymentTriggerType = "ad_hoc"
	ProjectEditResponseLatestDeploymentDeploymentTriggerTypeDeployHook ProjectEditResponseLatestDeploymentDeploymentTriggerType = "deploy_hook"
)

func (r ProjectEditResponseLatestDeploymentDeploymentTriggerType) IsKnown() bool {
	switch r {
	case ProjectEditResponseLatestDeploymentDeploymentTriggerTypeGitHubPush, ProjectEditResponseLatestDeploymentDeploymentTriggerTypeADHoc, ProjectEditResponseLatestDeploymentDeploymentTriggerTypeDeployHook:
		return true
	}
	return false
}

// A plaintext environment variable.
type ProjectEditResponseLatestDeploymentEnvVar struct {
	Type ProjectEditResponseLatestDeploymentEnvVarsType `json:"type,required"`
	// Environment variable value.
	Value string                                        `json:"value,required"`
	JSON  projectEditResponseLatestDeploymentEnvVarJSON `json:"-"`
	union ProjectEditResponseLatestDeploymentEnvVarsUnion
}

// projectEditResponseLatestDeploymentEnvVarJSON contains the JSON metadata for the
// struct [ProjectEditResponseLatestDeploymentEnvVar]
type projectEditResponseLatestDeploymentEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectEditResponseLatestDeploymentEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectEditResponseLatestDeploymentEnvVar) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectEditResponseLatestDeploymentEnvVar{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectEditResponseLatestDeploymentEnvVarsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectEditResponseLatestDeploymentEnvVarsPagesPlainTextEnvVar],
// [ProjectEditResponseLatestDeploymentEnvVarsPagesSecretTextEnvVar].
func (r ProjectEditResponseLatestDeploymentEnvVar) AsUnion() ProjectEditResponseLatestDeploymentEnvVarsUnion {
	return r.union
}

// A plaintext environment variable.
//
// Union satisfied by
// [ProjectEditResponseLatestDeploymentEnvVarsPagesPlainTextEnvVar] or
// [ProjectEditResponseLatestDeploymentEnvVarsPagesSecretTextEnvVar].
type ProjectEditResponseLatestDeploymentEnvVarsUnion interface {
	implementsProjectEditResponseLatestDeploymentEnvVar()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectEditResponseLatestDeploymentEnvVarsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectEditResponseLatestDeploymentEnvVarsPagesPlainTextEnvVar{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectEditResponseLatestDeploymentEnvVarsPagesSecretTextEnvVar{}),
			DiscriminatorValue: "secret_text",
		},
	)
}

// A plaintext environment variable.
type ProjectEditResponseLatestDeploymentEnvVarsPagesPlainTextEnvVar struct {
	Type ProjectEditResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarType `json:"type,required"`
	// Environment variable value.
	Value string                                                             `json:"value,required"`
	JSON  projectEditResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarJSON `json:"-"`
}

// projectEditResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarJSON contains the
// JSON metadata for the struct
// [ProjectEditResponseLatestDeploymentEnvVarsPagesPlainTextEnvVar]
type projectEditResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseLatestDeploymentEnvVarsPagesPlainTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectEditResponseLatestDeploymentEnvVarsPagesPlainTextEnvVar) implementsProjectEditResponseLatestDeploymentEnvVar() {
}

type ProjectEditResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarType string

const (
	ProjectEditResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarTypePlainText ProjectEditResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r ProjectEditResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectEditResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type ProjectEditResponseLatestDeploymentEnvVarsPagesSecretTextEnvVar struct {
	Type ProjectEditResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarType `json:"type,required"`
	// Secret value.
	Value string                                                              `json:"value,required"`
	JSON  projectEditResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarJSON `json:"-"`
}

// projectEditResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarJSON contains the
// JSON metadata for the struct
// [ProjectEditResponseLatestDeploymentEnvVarsPagesSecretTextEnvVar]
type projectEditResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseLatestDeploymentEnvVarsPagesSecretTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectEditResponseLatestDeploymentEnvVarsPagesSecretTextEnvVar) implementsProjectEditResponseLatestDeploymentEnvVar() {
}

type ProjectEditResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarType string

const (
	ProjectEditResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarTypeSecretText ProjectEditResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r ProjectEditResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectEditResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type ProjectEditResponseLatestDeploymentEnvVarsType string

const (
	ProjectEditResponseLatestDeploymentEnvVarsTypePlainText  ProjectEditResponseLatestDeploymentEnvVarsType = "plain_text"
	ProjectEditResponseLatestDeploymentEnvVarsTypeSecretText ProjectEditResponseLatestDeploymentEnvVarsType = "secret_text"
)

func (r ProjectEditResponseLatestDeploymentEnvVarsType) IsKnown() bool {
	switch r {
	case ProjectEditResponseLatestDeploymentEnvVarsTypePlainText, ProjectEditResponseLatestDeploymentEnvVarsTypeSecretText:
		return true
	}
	return false
}

// Type of deploy.
type ProjectEditResponseLatestDeploymentEnvironment string

const (
	ProjectEditResponseLatestDeploymentEnvironmentPreview    ProjectEditResponseLatestDeploymentEnvironment = "preview"
	ProjectEditResponseLatestDeploymentEnvironmentProduction ProjectEditResponseLatestDeploymentEnvironment = "production"
)

func (r ProjectEditResponseLatestDeploymentEnvironment) IsKnown() bool {
	switch r {
	case ProjectEditResponseLatestDeploymentEnvironmentPreview, ProjectEditResponseLatestDeploymentEnvironmentProduction:
		return true
	}
	return false
}

// Configs for the project source control.
type ProjectEditResponseLatestDeploymentSource struct {
	Config ProjectEditResponseLatestDeploymentSourceConfig `json:"config,required"`
	// The source control management provider.
	Type ProjectEditResponseLatestDeploymentSourceType `json:"type,required"`
	JSON projectEditResponseLatestDeploymentSourceJSON `json:"-"`
}

// projectEditResponseLatestDeploymentSourceJSON contains the JSON metadata for the
// struct [ProjectEditResponseLatestDeploymentSource]
type projectEditResponseLatestDeploymentSourceJSON struct {
	Config      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseLatestDeploymentSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseLatestDeploymentSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectEditResponseLatestDeploymentSourceConfig struct {
	// The owner of the repository.
	Owner string `json:"owner,required"`
	// Whether to enable PR comments.
	PrCommentsEnabled bool `json:"pr_comments_enabled,required"`
	// The production branch of the repository.
	ProductionBranch string `json:"production_branch,required"`
	// The name of the repository.
	RepoName string `json:"repo_name,required"`
	// Whether to enable automatic deployments when pushing to the source repository.
	// When disabled, no deployments (production or preview) will be triggered
	// automatically.
	//
	// Deprecated: Use `production_deployments_enabled` and
	// `preview_deployment_setting` for more granular control.
	DeploymentsEnabled bool `json:"deployments_enabled"`
	// The owner ID of the repository.
	OwnerID string `json:"owner_id"`
	// A list of paths that should be excluded from triggering a preview deployment.
	// Wildcard syntax (`*`) is supported.
	PathExcludes []string `json:"path_excludes"`
	// A list of paths that should be watched to trigger a preview deployment. Wildcard
	// syntax (`*`) is supported.
	PathIncludes []string `json:"path_includes"`
	// A list of branches that should not trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchExcludes []string `json:"preview_branch_excludes"`
	// A list of branches that should trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchIncludes []string `json:"preview_branch_includes"`
	// Controls whether commits to preview branches trigger a preview deployment.
	PreviewDeploymentSetting ProjectEditResponseLatestDeploymentSourceConfigPreviewDeploymentSetting `json:"preview_deployment_setting"`
	// Whether to trigger a production deployment on commits to the production branch.
	ProductionDeploymentsEnabled bool `json:"production_deployments_enabled"`
	// The ID of the repository.
	RepoID string                                              `json:"repo_id"`
	JSON   projectEditResponseLatestDeploymentSourceConfigJSON `json:"-"`
}

// projectEditResponseLatestDeploymentSourceConfigJSON contains the JSON metadata
// for the struct [ProjectEditResponseLatestDeploymentSourceConfig]
type projectEditResponseLatestDeploymentSourceConfigJSON struct {
	Owner                        apijson.Field
	PrCommentsEnabled            apijson.Field
	ProductionBranch             apijson.Field
	RepoName                     apijson.Field
	DeploymentsEnabled           apijson.Field
	OwnerID                      apijson.Field
	PathExcludes                 apijson.Field
	PathIncludes                 apijson.Field
	PreviewBranchExcludes        apijson.Field
	PreviewBranchIncludes        apijson.Field
	PreviewDeploymentSetting     apijson.Field
	ProductionDeploymentsEnabled apijson.Field
	RepoID                       apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ProjectEditResponseLatestDeploymentSourceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseLatestDeploymentSourceConfigJSON) RawJSON() string {
	return r.raw
}

// Controls whether commits to preview branches trigger a preview deployment.
type ProjectEditResponseLatestDeploymentSourceConfigPreviewDeploymentSetting string

const (
	ProjectEditResponseLatestDeploymentSourceConfigPreviewDeploymentSettingAll    ProjectEditResponseLatestDeploymentSourceConfigPreviewDeploymentSetting = "all"
	ProjectEditResponseLatestDeploymentSourceConfigPreviewDeploymentSettingNone   ProjectEditResponseLatestDeploymentSourceConfigPreviewDeploymentSetting = "none"
	ProjectEditResponseLatestDeploymentSourceConfigPreviewDeploymentSettingCustom ProjectEditResponseLatestDeploymentSourceConfigPreviewDeploymentSetting = "custom"
)

func (r ProjectEditResponseLatestDeploymentSourceConfigPreviewDeploymentSetting) IsKnown() bool {
	switch r {
	case ProjectEditResponseLatestDeploymentSourceConfigPreviewDeploymentSettingAll, ProjectEditResponseLatestDeploymentSourceConfigPreviewDeploymentSettingNone, ProjectEditResponseLatestDeploymentSourceConfigPreviewDeploymentSettingCustom:
		return true
	}
	return false
}

// The source control management provider.
type ProjectEditResponseLatestDeploymentSourceType string

const (
	ProjectEditResponseLatestDeploymentSourceTypeGitHub ProjectEditResponseLatestDeploymentSourceType = "github"
	ProjectEditResponseLatestDeploymentSourceTypeGitlab ProjectEditResponseLatestDeploymentSourceType = "gitlab"
)

func (r ProjectEditResponseLatestDeploymentSourceType) IsKnown() bool {
	switch r {
	case ProjectEditResponseLatestDeploymentSourceTypeGitHub, ProjectEditResponseLatestDeploymentSourceTypeGitlab:
		return true
	}
	return false
}

// Configs for the project build process.
type ProjectEditResponseBuildConfig struct {
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
	RootDir string                             `json:"root_dir,nullable"`
	JSON    projectEditResponseBuildConfigJSON `json:"-"`
}

// projectEditResponseBuildConfigJSON contains the JSON metadata for the struct
// [ProjectEditResponseBuildConfig]
type projectEditResponseBuildConfigJSON struct {
	WebAnalyticsTag   apijson.Field
	WebAnalyticsToken apijson.Field
	BuildCaching      apijson.Field
	BuildCommand      apijson.Field
	DestinationDir    apijson.Field
	RootDir           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectEditResponseBuildConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseBuildConfigJSON) RawJSON() string {
	return r.raw
}

// Configs for the project source control.
type ProjectEditResponseSource struct {
	Config ProjectEditResponseSourceConfig `json:"config,required"`
	// The source control management provider.
	Type ProjectEditResponseSourceType `json:"type,required"`
	JSON projectEditResponseSourceJSON `json:"-"`
}

// projectEditResponseSourceJSON contains the JSON metadata for the struct
// [ProjectEditResponseSource]
type projectEditResponseSourceJSON struct {
	Config      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectEditResponseSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectEditResponseSourceConfig struct {
	// The owner of the repository.
	Owner string `json:"owner,required"`
	// Whether to enable PR comments.
	PrCommentsEnabled bool `json:"pr_comments_enabled,required"`
	// The production branch of the repository.
	ProductionBranch string `json:"production_branch,required"`
	// The name of the repository.
	RepoName string `json:"repo_name,required"`
	// Whether to enable automatic deployments when pushing to the source repository.
	// When disabled, no deployments (production or preview) will be triggered
	// automatically.
	//
	// Deprecated: Use `production_deployments_enabled` and
	// `preview_deployment_setting` for more granular control.
	DeploymentsEnabled bool `json:"deployments_enabled"`
	// The owner ID of the repository.
	OwnerID string `json:"owner_id"`
	// A list of paths that should be excluded from triggering a preview deployment.
	// Wildcard syntax (`*`) is supported.
	PathExcludes []string `json:"path_excludes"`
	// A list of paths that should be watched to trigger a preview deployment. Wildcard
	// syntax (`*`) is supported.
	PathIncludes []string `json:"path_includes"`
	// A list of branches that should not trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchExcludes []string `json:"preview_branch_excludes"`
	// A list of branches that should trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchIncludes []string `json:"preview_branch_includes"`
	// Controls whether commits to preview branches trigger a preview deployment.
	PreviewDeploymentSetting ProjectEditResponseSourceConfigPreviewDeploymentSetting `json:"preview_deployment_setting"`
	// Whether to trigger a production deployment on commits to the production branch.
	ProductionDeploymentsEnabled bool `json:"production_deployments_enabled"`
	// The ID of the repository.
	RepoID string                              `json:"repo_id"`
	JSON   projectEditResponseSourceConfigJSON `json:"-"`
}

// projectEditResponseSourceConfigJSON contains the JSON metadata for the struct
// [ProjectEditResponseSourceConfig]
type projectEditResponseSourceConfigJSON struct {
	Owner                        apijson.Field
	PrCommentsEnabled            apijson.Field
	ProductionBranch             apijson.Field
	RepoName                     apijson.Field
	DeploymentsEnabled           apijson.Field
	OwnerID                      apijson.Field
	PathExcludes                 apijson.Field
	PathIncludes                 apijson.Field
	PreviewBranchExcludes        apijson.Field
	PreviewBranchIncludes        apijson.Field
	PreviewDeploymentSetting     apijson.Field
	ProductionDeploymentsEnabled apijson.Field
	RepoID                       apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ProjectEditResponseSourceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectEditResponseSourceConfigJSON) RawJSON() string {
	return r.raw
}

// Controls whether commits to preview branches trigger a preview deployment.
type ProjectEditResponseSourceConfigPreviewDeploymentSetting string

const (
	ProjectEditResponseSourceConfigPreviewDeploymentSettingAll    ProjectEditResponseSourceConfigPreviewDeploymentSetting = "all"
	ProjectEditResponseSourceConfigPreviewDeploymentSettingNone   ProjectEditResponseSourceConfigPreviewDeploymentSetting = "none"
	ProjectEditResponseSourceConfigPreviewDeploymentSettingCustom ProjectEditResponseSourceConfigPreviewDeploymentSetting = "custom"
)

func (r ProjectEditResponseSourceConfigPreviewDeploymentSetting) IsKnown() bool {
	switch r {
	case ProjectEditResponseSourceConfigPreviewDeploymentSettingAll, ProjectEditResponseSourceConfigPreviewDeploymentSettingNone, ProjectEditResponseSourceConfigPreviewDeploymentSettingCustom:
		return true
	}
	return false
}

// The source control management provider.
type ProjectEditResponseSourceType string

const (
	ProjectEditResponseSourceTypeGitHub ProjectEditResponseSourceType = "github"
	ProjectEditResponseSourceTypeGitlab ProjectEditResponseSourceType = "gitlab"
)

func (r ProjectEditResponseSourceType) IsKnown() bool {
	switch r {
	case ProjectEditResponseSourceTypeGitHub, ProjectEditResponseSourceTypeGitlab:
		return true
	}
	return false
}

type ProjectGetResponse struct {
	// ID of the project.
	ID string `json:"id,required"`
	// Most recent production deployment of the project.
	CanonicalDeployment ProjectGetResponseCanonicalDeployment `json:"canonical_deployment,required,nullable"`
	// When the project was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Configs for deployments in a project.
	DeploymentConfigs ProjectGetResponseDeploymentConfigs `json:"deployment_configs,required"`
	// Framework the project is using.
	Framework string `json:"framework,required"`
	// Version of the framework the project is using.
	FrameworkVersion string `json:"framework_version,required"`
	// Most recent deployment of the project.
	LatestDeployment ProjectGetResponseLatestDeployment `json:"latest_deployment,required,nullable"`
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
	BuildConfig ProjectGetResponseBuildConfig `json:"build_config"`
	// A list of associated custom domains for the project.
	Domains []string `json:"domains"`
	// Configs for the project source control.
	Source ProjectGetResponseSource `json:"source"`
	// The Cloudflare subdomain associated with the project.
	Subdomain string                 `json:"subdomain"`
	JSON      projectGetResponseJSON `json:"-"`
}

// projectGetResponseJSON contains the JSON metadata for the struct
// [ProjectGetResponse]
type projectGetResponseJSON struct {
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

func (r *ProjectGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseJSON) RawJSON() string {
	return r.raw
}

// Most recent production deployment of the project.
type ProjectGetResponseCanonicalDeployment struct {
	// Id of the deployment.
	ID string `json:"id,required"`
	// A list of alias URLs pointing to this deployment.
	Aliases []string `json:"aliases,required,nullable"`
	// Configs for the project build process.
	BuildConfig ProjectGetResponseCanonicalDeploymentBuildConfig `json:"build_config,required"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger ProjectGetResponseCanonicalDeploymentDeploymentTrigger `json:"deployment_trigger,required"`
	// Environment variables used for builds and Pages Functions.
	EnvVars map[string]ProjectGetResponseCanonicalDeploymentEnvVar `json:"env_vars,required,nullable"`
	// Type of deploy.
	Environment ProjectGetResponseCanonicalDeploymentEnvironment `json:"environment,required"`
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
	Source ProjectGetResponseCanonicalDeploymentSource `json:"source,required"`
	// List of past stages.
	Stages []Stage `json:"stages,required"`
	// The live URL to view this deployment.
	URL string `json:"url,required"`
	// Whether the deployment uses functions.
	UsesFunctions bool                                      `json:"uses_functions,nullable"`
	JSON          projectGetResponseCanonicalDeploymentJSON `json:"-"`
}

// projectGetResponseCanonicalDeploymentJSON contains the JSON metadata for the
// struct [ProjectGetResponseCanonicalDeployment]
type projectGetResponseCanonicalDeploymentJSON struct {
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

func (r *ProjectGetResponseCanonicalDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseCanonicalDeploymentJSON) RawJSON() string {
	return r.raw
}

// Configs for the project build process.
type ProjectGetResponseCanonicalDeploymentBuildConfig struct {
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
	RootDir string                                               `json:"root_dir,nullable"`
	JSON    projectGetResponseCanonicalDeploymentBuildConfigJSON `json:"-"`
}

// projectGetResponseCanonicalDeploymentBuildConfigJSON contains the JSON metadata
// for the struct [ProjectGetResponseCanonicalDeploymentBuildConfig]
type projectGetResponseCanonicalDeploymentBuildConfigJSON struct {
	WebAnalyticsTag   apijson.Field
	WebAnalyticsToken apijson.Field
	BuildCaching      apijson.Field
	BuildCommand      apijson.Field
	DestinationDir    apijson.Field
	RootDir           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectGetResponseCanonicalDeploymentBuildConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseCanonicalDeploymentBuildConfigJSON) RawJSON() string {
	return r.raw
}

// Info about what caused the deployment.
type ProjectGetResponseCanonicalDeploymentDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata ProjectGetResponseCanonicalDeploymentDeploymentTriggerMetadata `json:"metadata,required"`
	// What caused the deployment.
	Type ProjectGetResponseCanonicalDeploymentDeploymentTriggerType `json:"type,required"`
	JSON projectGetResponseCanonicalDeploymentDeploymentTriggerJSON `json:"-"`
}

// projectGetResponseCanonicalDeploymentDeploymentTriggerJSON contains the JSON
// metadata for the struct [ProjectGetResponseCanonicalDeploymentDeploymentTrigger]
type projectGetResponseCanonicalDeploymentDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseCanonicalDeploymentDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseCanonicalDeploymentDeploymentTriggerJSON) RawJSON() string {
	return r.raw
}

// Additional info about the trigger.
type ProjectGetResponseCanonicalDeploymentDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch,required"`
	// Whether the deployment trigger commit was dirty.
	CommitDirty bool `json:"commit_dirty,required"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash,required"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                             `json:"commit_message,required"`
	JSON          projectGetResponseCanonicalDeploymentDeploymentTriggerMetadataJSON `json:"-"`
}

// projectGetResponseCanonicalDeploymentDeploymentTriggerMetadataJSON contains the
// JSON metadata for the struct
// [ProjectGetResponseCanonicalDeploymentDeploymentTriggerMetadata]
type projectGetResponseCanonicalDeploymentDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitDirty   apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectGetResponseCanonicalDeploymentDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseCanonicalDeploymentDeploymentTriggerMetadataJSON) RawJSON() string {
	return r.raw
}

// What caused the deployment.
type ProjectGetResponseCanonicalDeploymentDeploymentTriggerType string

const (
	ProjectGetResponseCanonicalDeploymentDeploymentTriggerTypeGitHubPush ProjectGetResponseCanonicalDeploymentDeploymentTriggerType = "github:push"
	ProjectGetResponseCanonicalDeploymentDeploymentTriggerTypeADHoc      ProjectGetResponseCanonicalDeploymentDeploymentTriggerType = "ad_hoc"
	ProjectGetResponseCanonicalDeploymentDeploymentTriggerTypeDeployHook ProjectGetResponseCanonicalDeploymentDeploymentTriggerType = "deploy_hook"
)

func (r ProjectGetResponseCanonicalDeploymentDeploymentTriggerType) IsKnown() bool {
	switch r {
	case ProjectGetResponseCanonicalDeploymentDeploymentTriggerTypeGitHubPush, ProjectGetResponseCanonicalDeploymentDeploymentTriggerTypeADHoc, ProjectGetResponseCanonicalDeploymentDeploymentTriggerTypeDeployHook:
		return true
	}
	return false
}

// A plaintext environment variable.
type ProjectGetResponseCanonicalDeploymentEnvVar struct {
	Type ProjectGetResponseCanonicalDeploymentEnvVarsType `json:"type,required"`
	// Environment variable value.
	Value string                                          `json:"value,required"`
	JSON  projectGetResponseCanonicalDeploymentEnvVarJSON `json:"-"`
	union ProjectGetResponseCanonicalDeploymentEnvVarsUnion
}

// projectGetResponseCanonicalDeploymentEnvVarJSON contains the JSON metadata for
// the struct [ProjectGetResponseCanonicalDeploymentEnvVar]
type projectGetResponseCanonicalDeploymentEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectGetResponseCanonicalDeploymentEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectGetResponseCanonicalDeploymentEnvVar) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectGetResponseCanonicalDeploymentEnvVar{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectGetResponseCanonicalDeploymentEnvVarsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectGetResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVar],
// [ProjectGetResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVar].
func (r ProjectGetResponseCanonicalDeploymentEnvVar) AsUnion() ProjectGetResponseCanonicalDeploymentEnvVarsUnion {
	return r.union
}

// A plaintext environment variable.
//
// Union satisfied by
// [ProjectGetResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVar] or
// [ProjectGetResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVar].
type ProjectGetResponseCanonicalDeploymentEnvVarsUnion interface {
	implementsProjectGetResponseCanonicalDeploymentEnvVar()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectGetResponseCanonicalDeploymentEnvVarsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectGetResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVar{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectGetResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVar{}),
			DiscriminatorValue: "secret_text",
		},
	)
}

// A plaintext environment variable.
type ProjectGetResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVar struct {
	Type ProjectGetResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarType `json:"type,required"`
	// Environment variable value.
	Value string                                                               `json:"value,required"`
	JSON  projectGetResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarJSON `json:"-"`
}

// projectGetResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarJSON contains
// the JSON metadata for the struct
// [ProjectGetResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVar]
type projectGetResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectGetResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVar) implementsProjectGetResponseCanonicalDeploymentEnvVar() {
}

type ProjectGetResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarType string

const (
	ProjectGetResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarTypePlainText ProjectGetResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r ProjectGetResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectGetResponseCanonicalDeploymentEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type ProjectGetResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVar struct {
	Type ProjectGetResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarType `json:"type,required"`
	// Secret value.
	Value string                                                                `json:"value,required"`
	JSON  projectGetResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarJSON `json:"-"`
}

// projectGetResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarJSON contains
// the JSON metadata for the struct
// [ProjectGetResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVar]
type projectGetResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectGetResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVar) implementsProjectGetResponseCanonicalDeploymentEnvVar() {
}

type ProjectGetResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarType string

const (
	ProjectGetResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarTypeSecretText ProjectGetResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r ProjectGetResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectGetResponseCanonicalDeploymentEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type ProjectGetResponseCanonicalDeploymentEnvVarsType string

const (
	ProjectGetResponseCanonicalDeploymentEnvVarsTypePlainText  ProjectGetResponseCanonicalDeploymentEnvVarsType = "plain_text"
	ProjectGetResponseCanonicalDeploymentEnvVarsTypeSecretText ProjectGetResponseCanonicalDeploymentEnvVarsType = "secret_text"
)

func (r ProjectGetResponseCanonicalDeploymentEnvVarsType) IsKnown() bool {
	switch r {
	case ProjectGetResponseCanonicalDeploymentEnvVarsTypePlainText, ProjectGetResponseCanonicalDeploymentEnvVarsTypeSecretText:
		return true
	}
	return false
}

// Type of deploy.
type ProjectGetResponseCanonicalDeploymentEnvironment string

const (
	ProjectGetResponseCanonicalDeploymentEnvironmentPreview    ProjectGetResponseCanonicalDeploymentEnvironment = "preview"
	ProjectGetResponseCanonicalDeploymentEnvironmentProduction ProjectGetResponseCanonicalDeploymentEnvironment = "production"
)

func (r ProjectGetResponseCanonicalDeploymentEnvironment) IsKnown() bool {
	switch r {
	case ProjectGetResponseCanonicalDeploymentEnvironmentPreview, ProjectGetResponseCanonicalDeploymentEnvironmentProduction:
		return true
	}
	return false
}

// Configs for the project source control.
type ProjectGetResponseCanonicalDeploymentSource struct {
	Config ProjectGetResponseCanonicalDeploymentSourceConfig `json:"config,required"`
	// The source control management provider.
	Type ProjectGetResponseCanonicalDeploymentSourceType `json:"type,required"`
	JSON projectGetResponseCanonicalDeploymentSourceJSON `json:"-"`
}

// projectGetResponseCanonicalDeploymentSourceJSON contains the JSON metadata for
// the struct [ProjectGetResponseCanonicalDeploymentSource]
type projectGetResponseCanonicalDeploymentSourceJSON struct {
	Config      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseCanonicalDeploymentSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseCanonicalDeploymentSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectGetResponseCanonicalDeploymentSourceConfig struct {
	// The owner of the repository.
	Owner string `json:"owner,required"`
	// Whether to enable PR comments.
	PrCommentsEnabled bool `json:"pr_comments_enabled,required"`
	// The production branch of the repository.
	ProductionBranch string `json:"production_branch,required"`
	// The name of the repository.
	RepoName string `json:"repo_name,required"`
	// Whether to enable automatic deployments when pushing to the source repository.
	// When disabled, no deployments (production or preview) will be triggered
	// automatically.
	//
	// Deprecated: Use `production_deployments_enabled` and
	// `preview_deployment_setting` for more granular control.
	DeploymentsEnabled bool `json:"deployments_enabled"`
	// The owner ID of the repository.
	OwnerID string `json:"owner_id"`
	// A list of paths that should be excluded from triggering a preview deployment.
	// Wildcard syntax (`*`) is supported.
	PathExcludes []string `json:"path_excludes"`
	// A list of paths that should be watched to trigger a preview deployment. Wildcard
	// syntax (`*`) is supported.
	PathIncludes []string `json:"path_includes"`
	// A list of branches that should not trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchExcludes []string `json:"preview_branch_excludes"`
	// A list of branches that should trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchIncludes []string `json:"preview_branch_includes"`
	// Controls whether commits to preview branches trigger a preview deployment.
	PreviewDeploymentSetting ProjectGetResponseCanonicalDeploymentSourceConfigPreviewDeploymentSetting `json:"preview_deployment_setting"`
	// Whether to trigger a production deployment on commits to the production branch.
	ProductionDeploymentsEnabled bool `json:"production_deployments_enabled"`
	// The ID of the repository.
	RepoID string                                                `json:"repo_id"`
	JSON   projectGetResponseCanonicalDeploymentSourceConfigJSON `json:"-"`
}

// projectGetResponseCanonicalDeploymentSourceConfigJSON contains the JSON metadata
// for the struct [ProjectGetResponseCanonicalDeploymentSourceConfig]
type projectGetResponseCanonicalDeploymentSourceConfigJSON struct {
	Owner                        apijson.Field
	PrCommentsEnabled            apijson.Field
	ProductionBranch             apijson.Field
	RepoName                     apijson.Field
	DeploymentsEnabled           apijson.Field
	OwnerID                      apijson.Field
	PathExcludes                 apijson.Field
	PathIncludes                 apijson.Field
	PreviewBranchExcludes        apijson.Field
	PreviewBranchIncludes        apijson.Field
	PreviewDeploymentSetting     apijson.Field
	ProductionDeploymentsEnabled apijson.Field
	RepoID                       apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ProjectGetResponseCanonicalDeploymentSourceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseCanonicalDeploymentSourceConfigJSON) RawJSON() string {
	return r.raw
}

// Controls whether commits to preview branches trigger a preview deployment.
type ProjectGetResponseCanonicalDeploymentSourceConfigPreviewDeploymentSetting string

const (
	ProjectGetResponseCanonicalDeploymentSourceConfigPreviewDeploymentSettingAll    ProjectGetResponseCanonicalDeploymentSourceConfigPreviewDeploymentSetting = "all"
	ProjectGetResponseCanonicalDeploymentSourceConfigPreviewDeploymentSettingNone   ProjectGetResponseCanonicalDeploymentSourceConfigPreviewDeploymentSetting = "none"
	ProjectGetResponseCanonicalDeploymentSourceConfigPreviewDeploymentSettingCustom ProjectGetResponseCanonicalDeploymentSourceConfigPreviewDeploymentSetting = "custom"
)

func (r ProjectGetResponseCanonicalDeploymentSourceConfigPreviewDeploymentSetting) IsKnown() bool {
	switch r {
	case ProjectGetResponseCanonicalDeploymentSourceConfigPreviewDeploymentSettingAll, ProjectGetResponseCanonicalDeploymentSourceConfigPreviewDeploymentSettingNone, ProjectGetResponseCanonicalDeploymentSourceConfigPreviewDeploymentSettingCustom:
		return true
	}
	return false
}

// The source control management provider.
type ProjectGetResponseCanonicalDeploymentSourceType string

const (
	ProjectGetResponseCanonicalDeploymentSourceTypeGitHub ProjectGetResponseCanonicalDeploymentSourceType = "github"
	ProjectGetResponseCanonicalDeploymentSourceTypeGitlab ProjectGetResponseCanonicalDeploymentSourceType = "gitlab"
)

func (r ProjectGetResponseCanonicalDeploymentSourceType) IsKnown() bool {
	switch r {
	case ProjectGetResponseCanonicalDeploymentSourceTypeGitHub, ProjectGetResponseCanonicalDeploymentSourceTypeGitlab:
		return true
	}
	return false
}

// Configs for deployments in a project.
type ProjectGetResponseDeploymentConfigs struct {
	// Configs for preview deploys.
	Preview ProjectGetResponseDeploymentConfigsPreview `json:"preview,required"`
	// Configs for production deploys.
	Production ProjectGetResponseDeploymentConfigsProduction `json:"production,required"`
	JSON       projectGetResponseDeploymentConfigsJSON       `json:"-"`
}

// projectGetResponseDeploymentConfigsJSON contains the JSON metadata for the
// struct [ProjectGetResponseDeploymentConfigs]
type projectGetResponseDeploymentConfigsJSON struct {
	Preview     apijson.Field
	Production  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsJSON) RawJSON() string {
	return r.raw
}

// Configs for preview deploys.
type ProjectGetResponseDeploymentConfigsPreview struct {
	// Whether to always use the latest compatibility date for Pages Functions.
	AlwaysUseLatestCompatibilityDate bool `json:"always_use_latest_compatibility_date,required"`
	// The major version of the build image to use for Pages Functions.
	BuildImageMajorVersion int64 `json:"build_image_major_version,required"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate string `json:"compatibility_date,required"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags []string `json:"compatibility_flags,required"`
	// Environment variables used for builds and Pages Functions.
	EnvVars map[string]ProjectGetResponseDeploymentConfigsPreviewEnvVar `json:"env_vars,required,nullable"`
	// Whether to fail open when the deployment config cannot be applied.
	FailOpen bool `json:"fail_open,required"`
	// The usage model for Pages Functions.
	//
	// Deprecated: All new projects now use the Standard usage model.
	UsageModel ProjectGetResponseDeploymentConfigsPreviewUsageModel `json:"usage_model,required"`
	// Constellation bindings used for Pages Functions.
	AIBindings map[string]ProjectGetResponseDeploymentConfigsPreviewAIBinding `json:"ai_bindings"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets map[string]ProjectGetResponseDeploymentConfigsPreviewAnalyticsEngineDataset `json:"analytics_engine_datasets"`
	// Browser bindings used for Pages Functions.
	Browsers map[string]ProjectGetResponseDeploymentConfigsPreviewBrowser `json:"browsers"`
	// D1 databases used for Pages Functions.
	D1Databases map[string]ProjectGetResponseDeploymentConfigsPreviewD1Database `json:"d1_databases"`
	// Durable Object namespaces used for Pages Functions.
	DurableObjectNamespaces map[string]ProjectGetResponseDeploymentConfigsPreviewDurableObjectNamespace `json:"durable_object_namespaces"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings map[string]ProjectGetResponseDeploymentConfigsPreviewHyperdriveBinding `json:"hyperdrive_bindings"`
	// KV namespaces used for Pages Functions.
	KVNamespaces map[string]ProjectGetResponseDeploymentConfigsPreviewKVNamespace `json:"kv_namespaces"`
	// Limits for Pages Functions.
	Limits ProjectGetResponseDeploymentConfigsPreviewLimits `json:"limits"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates map[string]ProjectGetResponseDeploymentConfigsPreviewMTLSCertificate `json:"mtls_certificates"`
	// Placement setting used for Pages Functions.
	Placement ProjectGetResponseDeploymentConfigsPreviewPlacement `json:"placement"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers map[string]ProjectGetResponseDeploymentConfigsPreviewQueueProducer `json:"queue_producers"`
	// R2 buckets used for Pages Functions.
	R2Buckets map[string]ProjectGetResponseDeploymentConfigsPreviewR2Bucket `json:"r2_buckets"`
	// Services used for Pages Functions.
	Services map[string]ProjectGetResponseDeploymentConfigsPreviewService `json:"services"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings map[string]ProjectGetResponseDeploymentConfigsPreviewVectorizeBinding `json:"vectorize_bindings"`
	// Hash of the Wrangler configuration used for the deployment.
	WranglerConfigHash string                                         `json:"wrangler_config_hash"`
	JSON               projectGetResponseDeploymentConfigsPreviewJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewJSON contains the JSON metadata for
// the struct [ProjectGetResponseDeploymentConfigsPreview]
type projectGetResponseDeploymentConfigsPreviewJSON struct {
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

func (r *ProjectGetResponseDeploymentConfigsPreview) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewJSON) RawJSON() string {
	return r.raw
}

// A plaintext environment variable.
type ProjectGetResponseDeploymentConfigsPreviewEnvVar struct {
	Type ProjectGetResponseDeploymentConfigsPreviewEnvVarsType `json:"type,required"`
	// Environment variable value.
	Value string                                               `json:"value,required"`
	JSON  projectGetResponseDeploymentConfigsPreviewEnvVarJSON `json:"-"`
	union ProjectGetResponseDeploymentConfigsPreviewEnvVarsUnion
}

// projectGetResponseDeploymentConfigsPreviewEnvVarJSON contains the JSON metadata
// for the struct [ProjectGetResponseDeploymentConfigsPreviewEnvVar]
type projectGetResponseDeploymentConfigsPreviewEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectGetResponseDeploymentConfigsPreviewEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectGetResponseDeploymentConfigsPreviewEnvVar) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectGetResponseDeploymentConfigsPreviewEnvVar{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectGetResponseDeploymentConfigsPreviewEnvVarsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectGetResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar],
// [ProjectGetResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar].
func (r ProjectGetResponseDeploymentConfigsPreviewEnvVar) AsUnion() ProjectGetResponseDeploymentConfigsPreviewEnvVarsUnion {
	return r.union
}

// A plaintext environment variable.
//
// Union satisfied by
// [ProjectGetResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar] or
// [ProjectGetResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar].
type ProjectGetResponseDeploymentConfigsPreviewEnvVarsUnion interface {
	implementsProjectGetResponseDeploymentConfigsPreviewEnvVar()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectGetResponseDeploymentConfigsPreviewEnvVarsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectGetResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectGetResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar{}),
			DiscriminatorValue: "secret_text",
		},
	)
}

// A plaintext environment variable.
type ProjectGetResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar struct {
	Type ProjectGetResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType `json:"type,required"`
	// Environment variable value.
	Value string                                                                    `json:"value,required"`
	JSON  projectGetResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarJSON
// contains the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar]
type projectGetResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectGetResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVar) implementsProjectGetResponseDeploymentConfigsPreviewEnvVar() {
}

type ProjectGetResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType string

const (
	ProjectGetResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarTypePlainText ProjectGetResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r ProjectGetResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectGetResponseDeploymentConfigsPreviewEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type ProjectGetResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar struct {
	Type ProjectGetResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType `json:"type,required"`
	// Secret value.
	Value string                                                                     `json:"value,required"`
	JSON  projectGetResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarJSON
// contains the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar]
type projectGetResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectGetResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVar) implementsProjectGetResponseDeploymentConfigsPreviewEnvVar() {
}

type ProjectGetResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType string

const (
	ProjectGetResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarTypeSecretText ProjectGetResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r ProjectGetResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectGetResponseDeploymentConfigsPreviewEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type ProjectGetResponseDeploymentConfigsPreviewEnvVarsType string

const (
	ProjectGetResponseDeploymentConfigsPreviewEnvVarsTypePlainText  ProjectGetResponseDeploymentConfigsPreviewEnvVarsType = "plain_text"
	ProjectGetResponseDeploymentConfigsPreviewEnvVarsTypeSecretText ProjectGetResponseDeploymentConfigsPreviewEnvVarsType = "secret_text"
)

func (r ProjectGetResponseDeploymentConfigsPreviewEnvVarsType) IsKnown() bool {
	switch r {
	case ProjectGetResponseDeploymentConfigsPreviewEnvVarsTypePlainText, ProjectGetResponseDeploymentConfigsPreviewEnvVarsTypeSecretText:
		return true
	}
	return false
}

// The usage model for Pages Functions.
type ProjectGetResponseDeploymentConfigsPreviewUsageModel string

const (
	ProjectGetResponseDeploymentConfigsPreviewUsageModelStandard ProjectGetResponseDeploymentConfigsPreviewUsageModel = "standard"
	ProjectGetResponseDeploymentConfigsPreviewUsageModelBundled  ProjectGetResponseDeploymentConfigsPreviewUsageModel = "bundled"
	ProjectGetResponseDeploymentConfigsPreviewUsageModelUnbound  ProjectGetResponseDeploymentConfigsPreviewUsageModel = "unbound"
)

func (r ProjectGetResponseDeploymentConfigsPreviewUsageModel) IsKnown() bool {
	switch r {
	case ProjectGetResponseDeploymentConfigsPreviewUsageModelStandard, ProjectGetResponseDeploymentConfigsPreviewUsageModelBundled, ProjectGetResponseDeploymentConfigsPreviewUsageModelUnbound:
		return true
	}
	return false
}

// AI binding.
type ProjectGetResponseDeploymentConfigsPreviewAIBinding struct {
	ProjectID string                                                  `json:"project_id,required"`
	JSON      projectGetResponseDeploymentConfigsPreviewAIBindingJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewAIBindingJSON contains the JSON
// metadata for the struct [ProjectGetResponseDeploymentConfigsPreviewAIBinding]
type projectGetResponseDeploymentConfigsPreviewAIBindingJSON struct {
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewAIBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewAIBindingJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine binding.
type ProjectGetResponseDeploymentConfigsPreviewAnalyticsEngineDataset struct {
	// Name of the dataset.
	Dataset string                                                               `json:"dataset,required"`
	JSON    projectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasetJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasetJSON contains
// the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsPreviewAnalyticsEngineDataset]
type projectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasetJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewAnalyticsEngineDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewAnalyticsEngineDatasetJSON) RawJSON() string {
	return r.raw
}

// Browser binding.
type ProjectGetResponseDeploymentConfigsPreviewBrowser struct {
	JSON projectGetResponseDeploymentConfigsPreviewBrowserJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewBrowserJSON contains the JSON metadata
// for the struct [ProjectGetResponseDeploymentConfigsPreviewBrowser]
type projectGetResponseDeploymentConfigsPreviewBrowserJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewBrowser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewBrowserJSON) RawJSON() string {
	return r.raw
}

// D1 binding.
type ProjectGetResponseDeploymentConfigsPreviewD1Database struct {
	// UUID of the D1 database.
	ID   string                                                   `json:"id,required"`
	JSON projectGetResponseDeploymentConfigsPreviewD1DatabaseJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewD1DatabaseJSON contains the JSON
// metadata for the struct [ProjectGetResponseDeploymentConfigsPreviewD1Database]
type projectGetResponseDeploymentConfigsPreviewD1DatabaseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewD1Database) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewD1DatabaseJSON) RawJSON() string {
	return r.raw
}

// Durable Object binding.
type ProjectGetResponseDeploymentConfigsPreviewDurableObjectNamespace struct {
	// ID of the Durable Object namespace.
	NamespaceID string                                                               `json:"namespace_id,required"`
	JSON        projectGetResponseDeploymentConfigsPreviewDurableObjectNamespaceJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewDurableObjectNamespaceJSON contains
// the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsPreviewDurableObjectNamespace]
type projectGetResponseDeploymentConfigsPreviewDurableObjectNamespaceJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

// Hyperdrive binding.
type ProjectGetResponseDeploymentConfigsPreviewHyperdriveBinding struct {
	ID   string                                                          `json:"id,required"`
	JSON projectGetResponseDeploymentConfigsPreviewHyperdriveBindingJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewHyperdriveBindingJSON contains the
// JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsPreviewHyperdriveBinding]
type projectGetResponseDeploymentConfigsPreviewHyperdriveBindingJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewHyperdriveBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewHyperdriveBindingJSON) RawJSON() string {
	return r.raw
}

// KV namespace binding.
type ProjectGetResponseDeploymentConfigsPreviewKVNamespace struct {
	// ID of the KV namespace.
	NamespaceID string                                                    `json:"namespace_id,required"`
	JSON        projectGetResponseDeploymentConfigsPreviewKVNamespaceJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewKVNamespaceJSON contains the JSON
// metadata for the struct [ProjectGetResponseDeploymentConfigsPreviewKVNamespace]
type projectGetResponseDeploymentConfigsPreviewKVNamespaceJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewKVNamespaceJSON) RawJSON() string {
	return r.raw
}

// Limits for Pages Functions.
type ProjectGetResponseDeploymentConfigsPreviewLimits struct {
	// CPU time limit in milliseconds.
	CPUMs int64                                                `json:"cpu_ms,required"`
	JSON  projectGetResponseDeploymentConfigsPreviewLimitsJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewLimitsJSON contains the JSON metadata
// for the struct [ProjectGetResponseDeploymentConfigsPreviewLimits]
type projectGetResponseDeploymentConfigsPreviewLimitsJSON struct {
	CPUMs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewLimitsJSON) RawJSON() string {
	return r.raw
}

// mTLS binding.
type ProjectGetResponseDeploymentConfigsPreviewMTLSCertificate struct {
	CertificateID string                                                        `json:"certificate_id,required"`
	JSON          projectGetResponseDeploymentConfigsPreviewMTLSCertificateJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewMTLSCertificateJSON contains the JSON
// metadata for the struct
// [ProjectGetResponseDeploymentConfigsPreviewMTLSCertificate]
type projectGetResponseDeploymentConfigsPreviewMTLSCertificateJSON struct {
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewMTLSCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewMTLSCertificateJSON) RawJSON() string {
	return r.raw
}

// Placement setting used for Pages Functions.
type ProjectGetResponseDeploymentConfigsPreviewPlacement struct {
	// Placement mode.
	Mode string                                                  `json:"mode,required"`
	JSON projectGetResponseDeploymentConfigsPreviewPlacementJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewPlacementJSON contains the JSON
// metadata for the struct [ProjectGetResponseDeploymentConfigsPreviewPlacement]
type projectGetResponseDeploymentConfigsPreviewPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewPlacementJSON) RawJSON() string {
	return r.raw
}

// Queue Producer binding.
type ProjectGetResponseDeploymentConfigsPreviewQueueProducer struct {
	// Name of the Queue.
	Name string                                                      `json:"name,required"`
	JSON projectGetResponseDeploymentConfigsPreviewQueueProducerJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewQueueProducerJSON contains the JSON
// metadata for the struct
// [ProjectGetResponseDeploymentConfigsPreviewQueueProducer]
type projectGetResponseDeploymentConfigsPreviewQueueProducerJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewQueueProducer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewQueueProducerJSON) RawJSON() string {
	return r.raw
}

// R2 binding.
type ProjectGetResponseDeploymentConfigsPreviewR2Bucket struct {
	// Name of the R2 bucket.
	Name string `json:"name,required"`
	// Jurisdiction of the R2 bucket.
	Jurisdiction string                                                 `json:"jurisdiction,nullable"`
	JSON         projectGetResponseDeploymentConfigsPreviewR2BucketJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewR2BucketJSON contains the JSON
// metadata for the struct [ProjectGetResponseDeploymentConfigsPreviewR2Bucket]
type projectGetResponseDeploymentConfigsPreviewR2BucketJSON struct {
	Name         apijson.Field
	Jurisdiction apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewR2BucketJSON) RawJSON() string {
	return r.raw
}

// Service binding.
type ProjectGetResponseDeploymentConfigsPreviewService struct {
	// The Service environment.
	Environment string `json:"environment,required"`
	// The Service name.
	Service string `json:"service,required"`
	// The entrypoint to bind to.
	Entrypoint string                                                `json:"entrypoint,nullable"`
	JSON       projectGetResponseDeploymentConfigsPreviewServiceJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewServiceJSON contains the JSON metadata
// for the struct [ProjectGetResponseDeploymentConfigsPreviewService]
type projectGetResponseDeploymentConfigsPreviewServiceJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	Entrypoint  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewServiceJSON) RawJSON() string {
	return r.raw
}

// Vectorize binding.
type ProjectGetResponseDeploymentConfigsPreviewVectorizeBinding struct {
	IndexName string                                                         `json:"index_name,required"`
	JSON      projectGetResponseDeploymentConfigsPreviewVectorizeBindingJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsPreviewVectorizeBindingJSON contains the JSON
// metadata for the struct
// [ProjectGetResponseDeploymentConfigsPreviewVectorizeBinding]
type projectGetResponseDeploymentConfigsPreviewVectorizeBindingJSON struct {
	IndexName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsPreviewVectorizeBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsPreviewVectorizeBindingJSON) RawJSON() string {
	return r.raw
}

// Configs for production deploys.
type ProjectGetResponseDeploymentConfigsProduction struct {
	// Whether to always use the latest compatibility date for Pages Functions.
	AlwaysUseLatestCompatibilityDate bool `json:"always_use_latest_compatibility_date,required"`
	// The major version of the build image to use for Pages Functions.
	BuildImageMajorVersion int64 `json:"build_image_major_version,required"`
	// Compatibility date used for Pages Functions.
	CompatibilityDate string `json:"compatibility_date,required"`
	// Compatibility flags used for Pages Functions.
	CompatibilityFlags []string `json:"compatibility_flags,required"`
	// Environment variables used for builds and Pages Functions.
	EnvVars map[string]ProjectGetResponseDeploymentConfigsProductionEnvVar `json:"env_vars,required,nullable"`
	// Whether to fail open when the deployment config cannot be applied.
	FailOpen bool `json:"fail_open,required"`
	// The usage model for Pages Functions.
	//
	// Deprecated: All new projects now use the Standard usage model.
	UsageModel ProjectGetResponseDeploymentConfigsProductionUsageModel `json:"usage_model,required"`
	// Constellation bindings used for Pages Functions.
	AIBindings map[string]ProjectGetResponseDeploymentConfigsProductionAIBinding `json:"ai_bindings"`
	// Analytics Engine bindings used for Pages Functions.
	AnalyticsEngineDatasets map[string]ProjectGetResponseDeploymentConfigsProductionAnalyticsEngineDataset `json:"analytics_engine_datasets"`
	// Browser bindings used for Pages Functions.
	Browsers map[string]ProjectGetResponseDeploymentConfigsProductionBrowser `json:"browsers"`
	// D1 databases used for Pages Functions.
	D1Databases map[string]ProjectGetResponseDeploymentConfigsProductionD1Database `json:"d1_databases"`
	// Durable Object namespaces used for Pages Functions.
	DurableObjectNamespaces map[string]ProjectGetResponseDeploymentConfigsProductionDurableObjectNamespace `json:"durable_object_namespaces"`
	// Hyperdrive bindings used for Pages Functions.
	HyperdriveBindings map[string]ProjectGetResponseDeploymentConfigsProductionHyperdriveBinding `json:"hyperdrive_bindings"`
	// KV namespaces used for Pages Functions.
	KVNamespaces map[string]ProjectGetResponseDeploymentConfigsProductionKVNamespace `json:"kv_namespaces"`
	// Limits for Pages Functions.
	Limits ProjectGetResponseDeploymentConfigsProductionLimits `json:"limits"`
	// mTLS bindings used for Pages Functions.
	MTLSCertificates map[string]ProjectGetResponseDeploymentConfigsProductionMTLSCertificate `json:"mtls_certificates"`
	// Placement setting used for Pages Functions.
	Placement ProjectGetResponseDeploymentConfigsProductionPlacement `json:"placement"`
	// Queue Producer bindings used for Pages Functions.
	QueueProducers map[string]ProjectGetResponseDeploymentConfigsProductionQueueProducer `json:"queue_producers"`
	// R2 buckets used for Pages Functions.
	R2Buckets map[string]ProjectGetResponseDeploymentConfigsProductionR2Bucket `json:"r2_buckets"`
	// Services used for Pages Functions.
	Services map[string]ProjectGetResponseDeploymentConfigsProductionService `json:"services"`
	// Vectorize bindings used for Pages Functions.
	VectorizeBindings map[string]ProjectGetResponseDeploymentConfigsProductionVectorizeBinding `json:"vectorize_bindings"`
	// Hash of the Wrangler configuration used for the deployment.
	WranglerConfigHash string                                            `json:"wrangler_config_hash"`
	JSON               projectGetResponseDeploymentConfigsProductionJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionJSON contains the JSON metadata for
// the struct [ProjectGetResponseDeploymentConfigsProduction]
type projectGetResponseDeploymentConfigsProductionJSON struct {
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

func (r *ProjectGetResponseDeploymentConfigsProduction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionJSON) RawJSON() string {
	return r.raw
}

// A plaintext environment variable.
type ProjectGetResponseDeploymentConfigsProductionEnvVar struct {
	Type ProjectGetResponseDeploymentConfigsProductionEnvVarsType `json:"type,required"`
	// Environment variable value.
	Value string                                                  `json:"value,required"`
	JSON  projectGetResponseDeploymentConfigsProductionEnvVarJSON `json:"-"`
	union ProjectGetResponseDeploymentConfigsProductionEnvVarsUnion
}

// projectGetResponseDeploymentConfigsProductionEnvVarJSON contains the JSON
// metadata for the struct [ProjectGetResponseDeploymentConfigsProductionEnvVar]
type projectGetResponseDeploymentConfigsProductionEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectGetResponseDeploymentConfigsProductionEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectGetResponseDeploymentConfigsProductionEnvVar) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectGetResponseDeploymentConfigsProductionEnvVar{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectGetResponseDeploymentConfigsProductionEnvVarsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectGetResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar],
// [ProjectGetResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar].
func (r ProjectGetResponseDeploymentConfigsProductionEnvVar) AsUnion() ProjectGetResponseDeploymentConfigsProductionEnvVarsUnion {
	return r.union
}

// A plaintext environment variable.
//
// Union satisfied by
// [ProjectGetResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar] or
// [ProjectGetResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar].
type ProjectGetResponseDeploymentConfigsProductionEnvVarsUnion interface {
	implementsProjectGetResponseDeploymentConfigsProductionEnvVar()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectGetResponseDeploymentConfigsProductionEnvVarsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectGetResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectGetResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar{}),
			DiscriminatorValue: "secret_text",
		},
	)
}

// A plaintext environment variable.
type ProjectGetResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar struct {
	Type ProjectGetResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType `json:"type,required"`
	// Environment variable value.
	Value string                                                                       `json:"value,required"`
	JSON  projectGetResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarJSON
// contains the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar]
type projectGetResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectGetResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVar) implementsProjectGetResponseDeploymentConfigsProductionEnvVar() {
}

type ProjectGetResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType string

const (
	ProjectGetResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarTypePlainText ProjectGetResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r ProjectGetResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectGetResponseDeploymentConfigsProductionEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type ProjectGetResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar struct {
	Type ProjectGetResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType `json:"type,required"`
	// Secret value.
	Value string                                                                        `json:"value,required"`
	JSON  projectGetResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarJSON
// contains the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar]
type projectGetResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectGetResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVar) implementsProjectGetResponseDeploymentConfigsProductionEnvVar() {
}

type ProjectGetResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType string

const (
	ProjectGetResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarTypeSecretText ProjectGetResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r ProjectGetResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectGetResponseDeploymentConfigsProductionEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type ProjectGetResponseDeploymentConfigsProductionEnvVarsType string

const (
	ProjectGetResponseDeploymentConfigsProductionEnvVarsTypePlainText  ProjectGetResponseDeploymentConfigsProductionEnvVarsType = "plain_text"
	ProjectGetResponseDeploymentConfigsProductionEnvVarsTypeSecretText ProjectGetResponseDeploymentConfigsProductionEnvVarsType = "secret_text"
)

func (r ProjectGetResponseDeploymentConfigsProductionEnvVarsType) IsKnown() bool {
	switch r {
	case ProjectGetResponseDeploymentConfigsProductionEnvVarsTypePlainText, ProjectGetResponseDeploymentConfigsProductionEnvVarsTypeSecretText:
		return true
	}
	return false
}

// The usage model for Pages Functions.
type ProjectGetResponseDeploymentConfigsProductionUsageModel string

const (
	ProjectGetResponseDeploymentConfigsProductionUsageModelStandard ProjectGetResponseDeploymentConfigsProductionUsageModel = "standard"
	ProjectGetResponseDeploymentConfigsProductionUsageModelBundled  ProjectGetResponseDeploymentConfigsProductionUsageModel = "bundled"
	ProjectGetResponseDeploymentConfigsProductionUsageModelUnbound  ProjectGetResponseDeploymentConfigsProductionUsageModel = "unbound"
)

func (r ProjectGetResponseDeploymentConfigsProductionUsageModel) IsKnown() bool {
	switch r {
	case ProjectGetResponseDeploymentConfigsProductionUsageModelStandard, ProjectGetResponseDeploymentConfigsProductionUsageModelBundled, ProjectGetResponseDeploymentConfigsProductionUsageModelUnbound:
		return true
	}
	return false
}

// AI binding.
type ProjectGetResponseDeploymentConfigsProductionAIBinding struct {
	ProjectID string                                                     `json:"project_id,required"`
	JSON      projectGetResponseDeploymentConfigsProductionAIBindingJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionAIBindingJSON contains the JSON
// metadata for the struct [ProjectGetResponseDeploymentConfigsProductionAIBinding]
type projectGetResponseDeploymentConfigsProductionAIBindingJSON struct {
	ProjectID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionAIBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionAIBindingJSON) RawJSON() string {
	return r.raw
}

// Analytics Engine binding.
type ProjectGetResponseDeploymentConfigsProductionAnalyticsEngineDataset struct {
	// Name of the dataset.
	Dataset string                                                                  `json:"dataset,required"`
	JSON    projectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasetJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasetJSON contains
// the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionAnalyticsEngineDataset]
type projectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasetJSON struct {
	Dataset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionAnalyticsEngineDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionAnalyticsEngineDatasetJSON) RawJSON() string {
	return r.raw
}

// Browser binding.
type ProjectGetResponseDeploymentConfigsProductionBrowser struct {
	JSON projectGetResponseDeploymentConfigsProductionBrowserJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionBrowserJSON contains the JSON
// metadata for the struct [ProjectGetResponseDeploymentConfigsProductionBrowser]
type projectGetResponseDeploymentConfigsProductionBrowserJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionBrowser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionBrowserJSON) RawJSON() string {
	return r.raw
}

// D1 binding.
type ProjectGetResponseDeploymentConfigsProductionD1Database struct {
	// UUID of the D1 database.
	ID   string                                                      `json:"id,required"`
	JSON projectGetResponseDeploymentConfigsProductionD1DatabaseJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionD1DatabaseJSON contains the JSON
// metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionD1Database]
type projectGetResponseDeploymentConfigsProductionD1DatabaseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionD1Database) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionD1DatabaseJSON) RawJSON() string {
	return r.raw
}

// Durable Object binding.
type ProjectGetResponseDeploymentConfigsProductionDurableObjectNamespace struct {
	// ID of the Durable Object namespace.
	NamespaceID string                                                                  `json:"namespace_id,required"`
	JSON        projectGetResponseDeploymentConfigsProductionDurableObjectNamespaceJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionDurableObjectNamespaceJSON contains
// the JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionDurableObjectNamespace]
type projectGetResponseDeploymentConfigsProductionDurableObjectNamespaceJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionDurableObjectNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionDurableObjectNamespaceJSON) RawJSON() string {
	return r.raw
}

// Hyperdrive binding.
type ProjectGetResponseDeploymentConfigsProductionHyperdriveBinding struct {
	ID   string                                                             `json:"id,required"`
	JSON projectGetResponseDeploymentConfigsProductionHyperdriveBindingJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionHyperdriveBindingJSON contains the
// JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionHyperdriveBinding]
type projectGetResponseDeploymentConfigsProductionHyperdriveBindingJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionHyperdriveBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionHyperdriveBindingJSON) RawJSON() string {
	return r.raw
}

// KV namespace binding.
type ProjectGetResponseDeploymentConfigsProductionKVNamespace struct {
	// ID of the KV namespace.
	NamespaceID string                                                       `json:"namespace_id,required"`
	JSON        projectGetResponseDeploymentConfigsProductionKVNamespaceJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionKVNamespaceJSON contains the JSON
// metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionKVNamespace]
type projectGetResponseDeploymentConfigsProductionKVNamespaceJSON struct {
	NamespaceID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionKVNamespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionKVNamespaceJSON) RawJSON() string {
	return r.raw
}

// Limits for Pages Functions.
type ProjectGetResponseDeploymentConfigsProductionLimits struct {
	// CPU time limit in milliseconds.
	CPUMs int64                                                   `json:"cpu_ms,required"`
	JSON  projectGetResponseDeploymentConfigsProductionLimitsJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionLimitsJSON contains the JSON
// metadata for the struct [ProjectGetResponseDeploymentConfigsProductionLimits]
type projectGetResponseDeploymentConfigsProductionLimitsJSON struct {
	CPUMs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionLimitsJSON) RawJSON() string {
	return r.raw
}

// mTLS binding.
type ProjectGetResponseDeploymentConfigsProductionMTLSCertificate struct {
	CertificateID string                                                           `json:"certificate_id,required"`
	JSON          projectGetResponseDeploymentConfigsProductionMTLSCertificateJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionMTLSCertificateJSON contains the
// JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionMTLSCertificate]
type projectGetResponseDeploymentConfigsProductionMTLSCertificateJSON struct {
	CertificateID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionMTLSCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionMTLSCertificateJSON) RawJSON() string {
	return r.raw
}

// Placement setting used for Pages Functions.
type ProjectGetResponseDeploymentConfigsProductionPlacement struct {
	// Placement mode.
	Mode string                                                     `json:"mode,required"`
	JSON projectGetResponseDeploymentConfigsProductionPlacementJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionPlacementJSON contains the JSON
// metadata for the struct [ProjectGetResponseDeploymentConfigsProductionPlacement]
type projectGetResponseDeploymentConfigsProductionPlacementJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionPlacement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionPlacementJSON) RawJSON() string {
	return r.raw
}

// Queue Producer binding.
type ProjectGetResponseDeploymentConfigsProductionQueueProducer struct {
	// Name of the Queue.
	Name string                                                         `json:"name,required"`
	JSON projectGetResponseDeploymentConfigsProductionQueueProducerJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionQueueProducerJSON contains the JSON
// metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionQueueProducer]
type projectGetResponseDeploymentConfigsProductionQueueProducerJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionQueueProducer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionQueueProducerJSON) RawJSON() string {
	return r.raw
}

// R2 binding.
type ProjectGetResponseDeploymentConfigsProductionR2Bucket struct {
	// Name of the R2 bucket.
	Name string `json:"name,required"`
	// Jurisdiction of the R2 bucket.
	Jurisdiction string                                                    `json:"jurisdiction,nullable"`
	JSON         projectGetResponseDeploymentConfigsProductionR2BucketJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionR2BucketJSON contains the JSON
// metadata for the struct [ProjectGetResponseDeploymentConfigsProductionR2Bucket]
type projectGetResponseDeploymentConfigsProductionR2BucketJSON struct {
	Name         apijson.Field
	Jurisdiction apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionR2Bucket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionR2BucketJSON) RawJSON() string {
	return r.raw
}

// Service binding.
type ProjectGetResponseDeploymentConfigsProductionService struct {
	// The Service environment.
	Environment string `json:"environment,required"`
	// The Service name.
	Service string `json:"service,required"`
	// The entrypoint to bind to.
	Entrypoint string                                                   `json:"entrypoint,nullable"`
	JSON       projectGetResponseDeploymentConfigsProductionServiceJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionServiceJSON contains the JSON
// metadata for the struct [ProjectGetResponseDeploymentConfigsProductionService]
type projectGetResponseDeploymentConfigsProductionServiceJSON struct {
	Environment apijson.Field
	Service     apijson.Field
	Entrypoint  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionServiceJSON) RawJSON() string {
	return r.raw
}

// Vectorize binding.
type ProjectGetResponseDeploymentConfigsProductionVectorizeBinding struct {
	IndexName string                                                            `json:"index_name,required"`
	JSON      projectGetResponseDeploymentConfigsProductionVectorizeBindingJSON `json:"-"`
}

// projectGetResponseDeploymentConfigsProductionVectorizeBindingJSON contains the
// JSON metadata for the struct
// [ProjectGetResponseDeploymentConfigsProductionVectorizeBinding]
type projectGetResponseDeploymentConfigsProductionVectorizeBindingJSON struct {
	IndexName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseDeploymentConfigsProductionVectorizeBinding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseDeploymentConfigsProductionVectorizeBindingJSON) RawJSON() string {
	return r.raw
}

// Most recent deployment of the project.
type ProjectGetResponseLatestDeployment struct {
	// Id of the deployment.
	ID string `json:"id,required"`
	// A list of alias URLs pointing to this deployment.
	Aliases []string `json:"aliases,required,nullable"`
	// Configs for the project build process.
	BuildConfig ProjectGetResponseLatestDeploymentBuildConfig `json:"build_config,required"`
	// When the deployment was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Info about what caused the deployment.
	DeploymentTrigger ProjectGetResponseLatestDeploymentDeploymentTrigger `json:"deployment_trigger,required"`
	// Environment variables used for builds and Pages Functions.
	EnvVars map[string]ProjectGetResponseLatestDeploymentEnvVar `json:"env_vars,required,nullable"`
	// Type of deploy.
	Environment ProjectGetResponseLatestDeploymentEnvironment `json:"environment,required"`
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
	Source ProjectGetResponseLatestDeploymentSource `json:"source,required"`
	// List of past stages.
	Stages []Stage `json:"stages,required"`
	// The live URL to view this deployment.
	URL string `json:"url,required"`
	// Whether the deployment uses functions.
	UsesFunctions bool                                   `json:"uses_functions,nullable"`
	JSON          projectGetResponseLatestDeploymentJSON `json:"-"`
}

// projectGetResponseLatestDeploymentJSON contains the JSON metadata for the struct
// [ProjectGetResponseLatestDeployment]
type projectGetResponseLatestDeploymentJSON struct {
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

func (r *ProjectGetResponseLatestDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseLatestDeploymentJSON) RawJSON() string {
	return r.raw
}

// Configs for the project build process.
type ProjectGetResponseLatestDeploymentBuildConfig struct {
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
	RootDir string                                            `json:"root_dir,nullable"`
	JSON    projectGetResponseLatestDeploymentBuildConfigJSON `json:"-"`
}

// projectGetResponseLatestDeploymentBuildConfigJSON contains the JSON metadata for
// the struct [ProjectGetResponseLatestDeploymentBuildConfig]
type projectGetResponseLatestDeploymentBuildConfigJSON struct {
	WebAnalyticsTag   apijson.Field
	WebAnalyticsToken apijson.Field
	BuildCaching      apijson.Field
	BuildCommand      apijson.Field
	DestinationDir    apijson.Field
	RootDir           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectGetResponseLatestDeploymentBuildConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseLatestDeploymentBuildConfigJSON) RawJSON() string {
	return r.raw
}

// Info about what caused the deployment.
type ProjectGetResponseLatestDeploymentDeploymentTrigger struct {
	// Additional info about the trigger.
	Metadata ProjectGetResponseLatestDeploymentDeploymentTriggerMetadata `json:"metadata,required"`
	// What caused the deployment.
	Type ProjectGetResponseLatestDeploymentDeploymentTriggerType `json:"type,required"`
	JSON projectGetResponseLatestDeploymentDeploymentTriggerJSON `json:"-"`
}

// projectGetResponseLatestDeploymentDeploymentTriggerJSON contains the JSON
// metadata for the struct [ProjectGetResponseLatestDeploymentDeploymentTrigger]
type projectGetResponseLatestDeploymentDeploymentTriggerJSON struct {
	Metadata    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseLatestDeploymentDeploymentTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseLatestDeploymentDeploymentTriggerJSON) RawJSON() string {
	return r.raw
}

// Additional info about the trigger.
type ProjectGetResponseLatestDeploymentDeploymentTriggerMetadata struct {
	// Where the trigger happened.
	Branch string `json:"branch,required"`
	// Whether the deployment trigger commit was dirty.
	CommitDirty bool `json:"commit_dirty,required"`
	// Hash of the deployment trigger commit.
	CommitHash string `json:"commit_hash,required"`
	// Message of the deployment trigger commit.
	CommitMessage string                                                          `json:"commit_message,required"`
	JSON          projectGetResponseLatestDeploymentDeploymentTriggerMetadataJSON `json:"-"`
}

// projectGetResponseLatestDeploymentDeploymentTriggerMetadataJSON contains the
// JSON metadata for the struct
// [ProjectGetResponseLatestDeploymentDeploymentTriggerMetadata]
type projectGetResponseLatestDeploymentDeploymentTriggerMetadataJSON struct {
	Branch        apijson.Field
	CommitDirty   apijson.Field
	CommitHash    apijson.Field
	CommitMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProjectGetResponseLatestDeploymentDeploymentTriggerMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseLatestDeploymentDeploymentTriggerMetadataJSON) RawJSON() string {
	return r.raw
}

// What caused the deployment.
type ProjectGetResponseLatestDeploymentDeploymentTriggerType string

const (
	ProjectGetResponseLatestDeploymentDeploymentTriggerTypeGitHubPush ProjectGetResponseLatestDeploymentDeploymentTriggerType = "github:push"
	ProjectGetResponseLatestDeploymentDeploymentTriggerTypeADHoc      ProjectGetResponseLatestDeploymentDeploymentTriggerType = "ad_hoc"
	ProjectGetResponseLatestDeploymentDeploymentTriggerTypeDeployHook ProjectGetResponseLatestDeploymentDeploymentTriggerType = "deploy_hook"
)

func (r ProjectGetResponseLatestDeploymentDeploymentTriggerType) IsKnown() bool {
	switch r {
	case ProjectGetResponseLatestDeploymentDeploymentTriggerTypeGitHubPush, ProjectGetResponseLatestDeploymentDeploymentTriggerTypeADHoc, ProjectGetResponseLatestDeploymentDeploymentTriggerTypeDeployHook:
		return true
	}
	return false
}

// A plaintext environment variable.
type ProjectGetResponseLatestDeploymentEnvVar struct {
	Type ProjectGetResponseLatestDeploymentEnvVarsType `json:"type,required"`
	// Environment variable value.
	Value string                                       `json:"value,required"`
	JSON  projectGetResponseLatestDeploymentEnvVarJSON `json:"-"`
	union ProjectGetResponseLatestDeploymentEnvVarsUnion
}

// projectGetResponseLatestDeploymentEnvVarJSON contains the JSON metadata for the
// struct [ProjectGetResponseLatestDeploymentEnvVar]
type projectGetResponseLatestDeploymentEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r projectGetResponseLatestDeploymentEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r *ProjectGetResponseLatestDeploymentEnvVar) UnmarshalJSON(data []byte) (err error) {
	*r = ProjectGetResponseLatestDeploymentEnvVar{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ProjectGetResponseLatestDeploymentEnvVarsUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ProjectGetResponseLatestDeploymentEnvVarsPagesPlainTextEnvVar],
// [ProjectGetResponseLatestDeploymentEnvVarsPagesSecretTextEnvVar].
func (r ProjectGetResponseLatestDeploymentEnvVar) AsUnion() ProjectGetResponseLatestDeploymentEnvVarsUnion {
	return r.union
}

// A plaintext environment variable.
//
// Union satisfied by
// [ProjectGetResponseLatestDeploymentEnvVarsPagesPlainTextEnvVar] or
// [ProjectGetResponseLatestDeploymentEnvVarsPagesSecretTextEnvVar].
type ProjectGetResponseLatestDeploymentEnvVarsUnion interface {
	implementsProjectGetResponseLatestDeploymentEnvVar()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ProjectGetResponseLatestDeploymentEnvVarsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectGetResponseLatestDeploymentEnvVarsPagesPlainTextEnvVar{}),
			DiscriminatorValue: "plain_text",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ProjectGetResponseLatestDeploymentEnvVarsPagesSecretTextEnvVar{}),
			DiscriminatorValue: "secret_text",
		},
	)
}

// A plaintext environment variable.
type ProjectGetResponseLatestDeploymentEnvVarsPagesPlainTextEnvVar struct {
	Type ProjectGetResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarType `json:"type,required"`
	// Environment variable value.
	Value string                                                            `json:"value,required"`
	JSON  projectGetResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarJSON `json:"-"`
}

// projectGetResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarJSON contains the
// JSON metadata for the struct
// [ProjectGetResponseLatestDeploymentEnvVarsPagesPlainTextEnvVar]
type projectGetResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseLatestDeploymentEnvVarsPagesPlainTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectGetResponseLatestDeploymentEnvVarsPagesPlainTextEnvVar) implementsProjectGetResponseLatestDeploymentEnvVar() {
}

type ProjectGetResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarType string

const (
	ProjectGetResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarTypePlainText ProjectGetResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarType = "plain_text"
)

func (r ProjectGetResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectGetResponseLatestDeploymentEnvVarsPagesPlainTextEnvVarTypePlainText:
		return true
	}
	return false
}

// An encrypted environment variable.
type ProjectGetResponseLatestDeploymentEnvVarsPagesSecretTextEnvVar struct {
	Type ProjectGetResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarType `json:"type,required"`
	// Secret value.
	Value string                                                             `json:"value,required"`
	JSON  projectGetResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarJSON `json:"-"`
}

// projectGetResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarJSON contains the
// JSON metadata for the struct
// [ProjectGetResponseLatestDeploymentEnvVarsPagesSecretTextEnvVar]
type projectGetResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseLatestDeploymentEnvVarsPagesSecretTextEnvVar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarJSON) RawJSON() string {
	return r.raw
}

func (r ProjectGetResponseLatestDeploymentEnvVarsPagesSecretTextEnvVar) implementsProjectGetResponseLatestDeploymentEnvVar() {
}

type ProjectGetResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarType string

const (
	ProjectGetResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarTypeSecretText ProjectGetResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarType = "secret_text"
)

func (r ProjectGetResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarType) IsKnown() bool {
	switch r {
	case ProjectGetResponseLatestDeploymentEnvVarsPagesSecretTextEnvVarTypeSecretText:
		return true
	}
	return false
}

type ProjectGetResponseLatestDeploymentEnvVarsType string

const (
	ProjectGetResponseLatestDeploymentEnvVarsTypePlainText  ProjectGetResponseLatestDeploymentEnvVarsType = "plain_text"
	ProjectGetResponseLatestDeploymentEnvVarsTypeSecretText ProjectGetResponseLatestDeploymentEnvVarsType = "secret_text"
)

func (r ProjectGetResponseLatestDeploymentEnvVarsType) IsKnown() bool {
	switch r {
	case ProjectGetResponseLatestDeploymentEnvVarsTypePlainText, ProjectGetResponseLatestDeploymentEnvVarsTypeSecretText:
		return true
	}
	return false
}

// Type of deploy.
type ProjectGetResponseLatestDeploymentEnvironment string

const (
	ProjectGetResponseLatestDeploymentEnvironmentPreview    ProjectGetResponseLatestDeploymentEnvironment = "preview"
	ProjectGetResponseLatestDeploymentEnvironmentProduction ProjectGetResponseLatestDeploymentEnvironment = "production"
)

func (r ProjectGetResponseLatestDeploymentEnvironment) IsKnown() bool {
	switch r {
	case ProjectGetResponseLatestDeploymentEnvironmentPreview, ProjectGetResponseLatestDeploymentEnvironmentProduction:
		return true
	}
	return false
}

// Configs for the project source control.
type ProjectGetResponseLatestDeploymentSource struct {
	Config ProjectGetResponseLatestDeploymentSourceConfig `json:"config,required"`
	// The source control management provider.
	Type ProjectGetResponseLatestDeploymentSourceType `json:"type,required"`
	JSON projectGetResponseLatestDeploymentSourceJSON `json:"-"`
}

// projectGetResponseLatestDeploymentSourceJSON contains the JSON metadata for the
// struct [ProjectGetResponseLatestDeploymentSource]
type projectGetResponseLatestDeploymentSourceJSON struct {
	Config      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseLatestDeploymentSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseLatestDeploymentSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectGetResponseLatestDeploymentSourceConfig struct {
	// The owner of the repository.
	Owner string `json:"owner,required"`
	// Whether to enable PR comments.
	PrCommentsEnabled bool `json:"pr_comments_enabled,required"`
	// The production branch of the repository.
	ProductionBranch string `json:"production_branch,required"`
	// The name of the repository.
	RepoName string `json:"repo_name,required"`
	// Whether to enable automatic deployments when pushing to the source repository.
	// When disabled, no deployments (production or preview) will be triggered
	// automatically.
	//
	// Deprecated: Use `production_deployments_enabled` and
	// `preview_deployment_setting` for more granular control.
	DeploymentsEnabled bool `json:"deployments_enabled"`
	// The owner ID of the repository.
	OwnerID string `json:"owner_id"`
	// A list of paths that should be excluded from triggering a preview deployment.
	// Wildcard syntax (`*`) is supported.
	PathExcludes []string `json:"path_excludes"`
	// A list of paths that should be watched to trigger a preview deployment. Wildcard
	// syntax (`*`) is supported.
	PathIncludes []string `json:"path_includes"`
	// A list of branches that should not trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchExcludes []string `json:"preview_branch_excludes"`
	// A list of branches that should trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchIncludes []string `json:"preview_branch_includes"`
	// Controls whether commits to preview branches trigger a preview deployment.
	PreviewDeploymentSetting ProjectGetResponseLatestDeploymentSourceConfigPreviewDeploymentSetting `json:"preview_deployment_setting"`
	// Whether to trigger a production deployment on commits to the production branch.
	ProductionDeploymentsEnabled bool `json:"production_deployments_enabled"`
	// The ID of the repository.
	RepoID string                                             `json:"repo_id"`
	JSON   projectGetResponseLatestDeploymentSourceConfigJSON `json:"-"`
}

// projectGetResponseLatestDeploymentSourceConfigJSON contains the JSON metadata
// for the struct [ProjectGetResponseLatestDeploymentSourceConfig]
type projectGetResponseLatestDeploymentSourceConfigJSON struct {
	Owner                        apijson.Field
	PrCommentsEnabled            apijson.Field
	ProductionBranch             apijson.Field
	RepoName                     apijson.Field
	DeploymentsEnabled           apijson.Field
	OwnerID                      apijson.Field
	PathExcludes                 apijson.Field
	PathIncludes                 apijson.Field
	PreviewBranchExcludes        apijson.Field
	PreviewBranchIncludes        apijson.Field
	PreviewDeploymentSetting     apijson.Field
	ProductionDeploymentsEnabled apijson.Field
	RepoID                       apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ProjectGetResponseLatestDeploymentSourceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseLatestDeploymentSourceConfigJSON) RawJSON() string {
	return r.raw
}

// Controls whether commits to preview branches trigger a preview deployment.
type ProjectGetResponseLatestDeploymentSourceConfigPreviewDeploymentSetting string

const (
	ProjectGetResponseLatestDeploymentSourceConfigPreviewDeploymentSettingAll    ProjectGetResponseLatestDeploymentSourceConfigPreviewDeploymentSetting = "all"
	ProjectGetResponseLatestDeploymentSourceConfigPreviewDeploymentSettingNone   ProjectGetResponseLatestDeploymentSourceConfigPreviewDeploymentSetting = "none"
	ProjectGetResponseLatestDeploymentSourceConfigPreviewDeploymentSettingCustom ProjectGetResponseLatestDeploymentSourceConfigPreviewDeploymentSetting = "custom"
)

func (r ProjectGetResponseLatestDeploymentSourceConfigPreviewDeploymentSetting) IsKnown() bool {
	switch r {
	case ProjectGetResponseLatestDeploymentSourceConfigPreviewDeploymentSettingAll, ProjectGetResponseLatestDeploymentSourceConfigPreviewDeploymentSettingNone, ProjectGetResponseLatestDeploymentSourceConfigPreviewDeploymentSettingCustom:
		return true
	}
	return false
}

// The source control management provider.
type ProjectGetResponseLatestDeploymentSourceType string

const (
	ProjectGetResponseLatestDeploymentSourceTypeGitHub ProjectGetResponseLatestDeploymentSourceType = "github"
	ProjectGetResponseLatestDeploymentSourceTypeGitlab ProjectGetResponseLatestDeploymentSourceType = "gitlab"
)

func (r ProjectGetResponseLatestDeploymentSourceType) IsKnown() bool {
	switch r {
	case ProjectGetResponseLatestDeploymentSourceTypeGitHub, ProjectGetResponseLatestDeploymentSourceTypeGitlab:
		return true
	}
	return false
}

// Configs for the project build process.
type ProjectGetResponseBuildConfig struct {
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
	RootDir string                            `json:"root_dir,nullable"`
	JSON    projectGetResponseBuildConfigJSON `json:"-"`
}

// projectGetResponseBuildConfigJSON contains the JSON metadata for the struct
// [ProjectGetResponseBuildConfig]
type projectGetResponseBuildConfigJSON struct {
	WebAnalyticsTag   apijson.Field
	WebAnalyticsToken apijson.Field
	BuildCaching      apijson.Field
	BuildCommand      apijson.Field
	DestinationDir    apijson.Field
	RootDir           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ProjectGetResponseBuildConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseBuildConfigJSON) RawJSON() string {
	return r.raw
}

// Configs for the project source control.
type ProjectGetResponseSource struct {
	Config ProjectGetResponseSourceConfig `json:"config,required"`
	// The source control management provider.
	Type ProjectGetResponseSourceType `json:"type,required"`
	JSON projectGetResponseSourceJSON `json:"-"`
}

// projectGetResponseSourceJSON contains the JSON metadata for the struct
// [ProjectGetResponseSource]
type projectGetResponseSourceJSON struct {
	Config      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectGetResponseSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectGetResponseSourceConfig struct {
	// The owner of the repository.
	Owner string `json:"owner,required"`
	// Whether to enable PR comments.
	PrCommentsEnabled bool `json:"pr_comments_enabled,required"`
	// The production branch of the repository.
	ProductionBranch string `json:"production_branch,required"`
	// The name of the repository.
	RepoName string `json:"repo_name,required"`
	// Whether to enable automatic deployments when pushing to the source repository.
	// When disabled, no deployments (production or preview) will be triggered
	// automatically.
	//
	// Deprecated: Use `production_deployments_enabled` and
	// `preview_deployment_setting` for more granular control.
	DeploymentsEnabled bool `json:"deployments_enabled"`
	// The owner ID of the repository.
	OwnerID string `json:"owner_id"`
	// A list of paths that should be excluded from triggering a preview deployment.
	// Wildcard syntax (`*`) is supported.
	PathExcludes []string `json:"path_excludes"`
	// A list of paths that should be watched to trigger a preview deployment. Wildcard
	// syntax (`*`) is supported.
	PathIncludes []string `json:"path_includes"`
	// A list of branches that should not trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchExcludes []string `json:"preview_branch_excludes"`
	// A list of branches that should trigger a preview deployment. Wildcard syntax
	// (`*`) is supported. Must be used with `preview_deployment_setting` set to
	// `custom`.
	PreviewBranchIncludes []string `json:"preview_branch_includes"`
	// Controls whether commits to preview branches trigger a preview deployment.
	PreviewDeploymentSetting ProjectGetResponseSourceConfigPreviewDeploymentSetting `json:"preview_deployment_setting"`
	// Whether to trigger a production deployment on commits to the production branch.
	ProductionDeploymentsEnabled bool `json:"production_deployments_enabled"`
	// The ID of the repository.
	RepoID string                             `json:"repo_id"`
	JSON   projectGetResponseSourceConfigJSON `json:"-"`
}

// projectGetResponseSourceConfigJSON contains the JSON metadata for the struct
// [ProjectGetResponseSourceConfig]
type projectGetResponseSourceConfigJSON struct {
	Owner                        apijson.Field
	PrCommentsEnabled            apijson.Field
	ProductionBranch             apijson.Field
	RepoName                     apijson.Field
	DeploymentsEnabled           apijson.Field
	OwnerID                      apijson.Field
	PathExcludes                 apijson.Field
	PathIncludes                 apijson.Field
	PreviewBranchExcludes        apijson.Field
	PreviewBranchIncludes        apijson.Field
	PreviewDeploymentSetting     apijson.Field
	ProductionDeploymentsEnabled apijson.Field
	RepoID                       apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ProjectGetResponseSourceConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectGetResponseSourceConfigJSON) RawJSON() string {
	return r.raw
}

// Controls whether commits to preview branches trigger a preview deployment.
type ProjectGetResponseSourceConfigPreviewDeploymentSetting string

const (
	ProjectGetResponseSourceConfigPreviewDeploymentSettingAll    ProjectGetResponseSourceConfigPreviewDeploymentSetting = "all"
	ProjectGetResponseSourceConfigPreviewDeploymentSettingNone   ProjectGetResponseSourceConfigPreviewDeploymentSetting = "none"
	ProjectGetResponseSourceConfigPreviewDeploymentSettingCustom ProjectGetResponseSourceConfigPreviewDeploymentSetting = "custom"
)

func (r ProjectGetResponseSourceConfigPreviewDeploymentSetting) IsKnown() bool {
	switch r {
	case ProjectGetResponseSourceConfigPreviewDeploymentSettingAll, ProjectGetResponseSourceConfigPreviewDeploymentSettingNone, ProjectGetResponseSourceConfigPreviewDeploymentSettingCustom:
		return true
	}
	return false
}

// The source control management provider.
type ProjectGetResponseSourceType string

const (
	ProjectGetResponseSourceTypeGitHub ProjectGetResponseSourceType = "github"
	ProjectGetResponseSourceTypeGitlab ProjectGetResponseSourceType = "gitlab"
)

func (r ProjectGetResponseSourceType) IsKnown() bool {
	switch r {
	case ProjectGetResponseSourceTypeGitHub, ProjectGetResponseSourceTypeGitlab:
		return true
	}
	return false
}

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
	Result   ProjectNewResponse                   `json:"result,required"`
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
	Result   ProjectEditResponse                   `json:"result,required"`
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
	Result   ProjectGetResponse                   `json:"result,required"`
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
