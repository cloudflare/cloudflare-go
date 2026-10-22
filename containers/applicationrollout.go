// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package containers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// ApplicationRolloutService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewApplicationRolloutService] method instead.
type ApplicationRolloutService struct {
	Options []option.RequestOption
}

// NewApplicationRolloutService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewApplicationRolloutService(opts ...option.RequestOption) (r *ApplicationRolloutService) {
	r = &ApplicationRolloutService{}
	r.Options = opts
	return
}

// Creates a rollout to update the application's configuration across instances
// with minimal downtime. Rollouts apply only to scheduler-backed applications with
// `scheduling_policy: "default"`. Versions and rollouts do not apply to
// applications with `scheduling_policy: "durable_object"`.
func (r *ApplicationRolloutService) New(ctx context.Context, applicationID string, params ApplicationRolloutNewParams, opts ...option.RequestOption) (res *ApplicationRolloutNewResponse, err error) {
	var env ApplicationRolloutNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if applicationID == "" {
		err = errors.New("missing required application_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/containers/applications/%s/rollouts", params.AccountID, applicationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Represents the status and metadata of a rollout process for an application. For
// "rolling" strategy: includes steps and progress with instance counts. For
// "new_instances" strategy: the response omits steps and progress. Use percentage,
// version_distribution, and health.summary for status.
type ApplicationRolloutNewResponse struct {
	// An identifier for a specific rollout within an application.
	ID string `json:"id" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	CreatedAt string `json:"created_at" api:"required"`
	// User-specified container configuration changes.
	CurrentConfiguration ApplicationRolloutNewResponseCurrentConfiguration `json:"current_configuration" api:"required"`
	// Current application version before the rollout.
	CurrentVersion int64                               `json:"current_version" api:"required"`
	Description    string                              `json:"description" api:"required"`
	Health         ApplicationRolloutNewResponseHealth `json:"health" api:"required"`
	// Kind of the rollout process.
	//
	// - "full_auto": For rolling rollouts, starts progressing steps upon rollout
	//   creation. For new_instances rollouts, advances percentage targets
	//   automatically after target-version health is observed.
	// - "full_manual": Requires manually progressing each step in the rollout using
	//   the UpdateRollout's action paramater.
	// - "durable_objects_auto": Default when the application is a DO application.
	Kind ApplicationRolloutNewResponseKind `json:"kind" api:"required"`
	// Timestamp of the most recent update to status, health, or progress.
	LastUpdatedAt string `json:"last_updated_at" api:"required"`
	// Current status of the rollout.
	Status ApplicationRolloutNewResponseStatus `json:"status" api:"required"`
	// The rollout strategy.
	//
	// - "rolling": Step-based rollout with health gates. Actively replaces instances
	//   to reach each step's target percentage. Response includes steps and progress.
	// - "new_instances": Percentage control over version distribution. Version sync
	//   actively replaces instances to match the configured percentage. "full_auto"
	//   ramps through fixed percentage targets after target-version health is
	//   observed. Response includes percentage, version_distribution, and
	//   health.summary.
	Strategy ApplicationRolloutNewResponseStrategy `json:"strategy" api:"required"`
	// User-specified container configuration changes.
	TargetConfiguration ApplicationRolloutNewResponseTargetConfiguration `json:"target_configuration" api:"required"`
	// Target application version after the rollout is complete and applied to all
	// current instances.
	TargetVersion int64 `json:"target_version" api:"required"`
	// Current target version percentage (0-100). Only present for "new_instances"
	// strategy.
	Percentage int64 `json:"percentage"`
	// Progress details of an application rollout.
	Progress ApplicationRolloutNewResponseProgress `json:"progress"`
	// Timestamp when the rollout started.
	StartedAt time.Time                           `json:"started_at" format:"date-time"`
	Steps     []ApplicationRolloutNewResponseStep `json:"steps"`
	// Version percentage distribution. Only present for "new_instances" strategy. For
	// "rolling" strategy, see progress.version_distribution instead.
	VersionDistribution ApplicationRolloutNewResponseVersionDistribution `json:"version_distribution"`
	JSON                applicationRolloutNewResponseJSON                `json:"-"`
}

// applicationRolloutNewResponseJSON contains the JSON metadata for the struct
// [ApplicationRolloutNewResponse]
type applicationRolloutNewResponseJSON struct {
	ID                   apijson.Field
	CreatedAt            apijson.Field
	CurrentConfiguration apijson.Field
	CurrentVersion       apijson.Field
	Description          apijson.Field
	Health               apijson.Field
	Kind                 apijson.Field
	LastUpdatedAt        apijson.Field
	Status               apijson.Field
	Strategy             apijson.Field
	TargetConfiguration  apijson.Field
	TargetVersion        apijson.Field
	Percentage           apijson.Field
	Progress             apijson.Field
	StartedAt            apijson.Field
	Steps                apijson.Field
	VersionDistribution  apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ApplicationRolloutNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationRolloutNewResponseJSON) RawJSON() string {
	return r.raw
}

// User-specified container configuration changes.
type ApplicationRolloutNewResponseCurrentConfiguration struct {
	AuthorizedKeys []ApplicationRolloutNewResponseCurrentConfigurationAuthorizedKey `json:"authorized_keys"`
	// The command that runs when the container starts, passed to the entrypoint. You
	// can override this at run-time. If you override only the command, it gets passed
	// to the default entrypoint specified in the image.
	Command []string `json:"command"`
	// The entry point for the container, specifying the executable to run when the
	// container starts. You can override this at run-time. If you do, the default
	// command from the image is ignored. Specify both entrypoint and command at
	// run-time to completely replace the image defaults.
	Entrypoint []string `json:"entrypoint"`
	// Container environment variables.
	EnvironmentVariables []ApplicationRolloutNewResponseCurrentConfigurationEnvironmentVariable `json:"environment_variables"`
	// Image url.
	Image string `json:"image"`
	// The instance type configures vCPU, memory, and disk.
	//
	// - "lite": 1/16 vCPU, 256 MiB memory, 2 GB disk
	// - "basic": 1/4 vCPU, 1 GiB memory, 4 GB disk
	// - "standard-1": 1/2 vCPU, 4 GiB memory, 8 GB disk
	// - "standard-2": 1 vCPU, 6 GiB memory, 12 GB disk
	// - "standard-3": 2 vCPU, 8 GiB memory, 16 GB disk
	// - "standard-4": 4 vCPU, 12 GiB memory, 20 GB disk
	InstanceType ApplicationRolloutNewResponseCurrentConfigurationInstanceType `json:"instance_type"`
	// Settings for deployment observability such as logging.
	Observability ApplicationRolloutNewResponseCurrentConfigurationObservability `json:"observability"`
	JSON          applicationRolloutNewResponseCurrentConfigurationJSON          `json:"-"`
}

// applicationRolloutNewResponseCurrentConfigurationJSON contains the JSON metadata
// for the struct [ApplicationRolloutNewResponseCurrentConfiguration]
type applicationRolloutNewResponseCurrentConfigurationJSON struct {
	AuthorizedKeys       apijson.Field
	Command              apijson.Field
	Entrypoint           apijson.Field
	EnvironmentVariables apijson.Field
	Image                apijson.Field
	InstanceType         apijson.Field
	Observability        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ApplicationRolloutNewResponseCurrentConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationRolloutNewResponseCurrentConfigurationJSON) RawJSON() string {
	return r.raw
}

// User-provided SSH public key.
type ApplicationRolloutNewResponseCurrentConfigurationAuthorizedKey struct {
	// An SSH public key.
	PublicKey string `json:"public_key" api:"required"`
	// Optional human readable name for this key.
	Name string                                                             `json:"name"`
	JSON applicationRolloutNewResponseCurrentConfigurationAuthorizedKeyJSON `json:"-"`
}

// applicationRolloutNewResponseCurrentConfigurationAuthorizedKeyJSON contains the
// JSON metadata for the struct
// [ApplicationRolloutNewResponseCurrentConfigurationAuthorizedKey]
type applicationRolloutNewResponseCurrentConfigurationAuthorizedKeyJSON struct {
	PublicKey   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationRolloutNewResponseCurrentConfigurationAuthorizedKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationRolloutNewResponseCurrentConfigurationAuthorizedKeyJSON) RawJSON() string {
	return r.raw
}

// An environment variable with a value set.
type ApplicationRolloutNewResponseCurrentConfigurationEnvironmentVariable struct {
	// An environment variable name.
	Name string `json:"name" api:"required"`
	// An environment variable value.
	Value string                                                                   `json:"value" api:"required"`
	JSON  applicationRolloutNewResponseCurrentConfigurationEnvironmentVariableJSON `json:"-"`
}

// applicationRolloutNewResponseCurrentConfigurationEnvironmentVariableJSON
// contains the JSON metadata for the struct
// [ApplicationRolloutNewResponseCurrentConfigurationEnvironmentVariable]
type applicationRolloutNewResponseCurrentConfigurationEnvironmentVariableJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationRolloutNewResponseCurrentConfigurationEnvironmentVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationRolloutNewResponseCurrentConfigurationEnvironmentVariableJSON) RawJSON() string {
	return r.raw
}

// The instance type configures vCPU, memory, and disk.
//
// - "lite": 1/16 vCPU, 256 MiB memory, 2 GB disk
// - "basic": 1/4 vCPU, 1 GiB memory, 4 GB disk
// - "standard-1": 1/2 vCPU, 4 GiB memory, 8 GB disk
// - "standard-2": 1 vCPU, 6 GiB memory, 12 GB disk
// - "standard-3": 2 vCPU, 8 GiB memory, 16 GB disk
// - "standard-4": 4 vCPU, 12 GiB memory, 20 GB disk
type ApplicationRolloutNewResponseCurrentConfigurationInstanceType string

const (
	ApplicationRolloutNewResponseCurrentConfigurationInstanceTypeLite      ApplicationRolloutNewResponseCurrentConfigurationInstanceType = "lite"
	ApplicationRolloutNewResponseCurrentConfigurationInstanceTypeBasic     ApplicationRolloutNewResponseCurrentConfigurationInstanceType = "basic"
	ApplicationRolloutNewResponseCurrentConfigurationInstanceTypeStandard1 ApplicationRolloutNewResponseCurrentConfigurationInstanceType = "standard-1"
	ApplicationRolloutNewResponseCurrentConfigurationInstanceTypeStandard2 ApplicationRolloutNewResponseCurrentConfigurationInstanceType = "standard-2"
	ApplicationRolloutNewResponseCurrentConfigurationInstanceTypeStandard3 ApplicationRolloutNewResponseCurrentConfigurationInstanceType = "standard-3"
	ApplicationRolloutNewResponseCurrentConfigurationInstanceTypeStandard4 ApplicationRolloutNewResponseCurrentConfigurationInstanceType = "standard-4"
)

func (r ApplicationRolloutNewResponseCurrentConfigurationInstanceType) IsKnown() bool {
	switch r {
	case ApplicationRolloutNewResponseCurrentConfigurationInstanceTypeLite, ApplicationRolloutNewResponseCurrentConfigurationInstanceTypeBasic, ApplicationRolloutNewResponseCurrentConfigurationInstanceTypeStandard1, ApplicationRolloutNewResponseCurrentConfigurationInstanceTypeStandard2, ApplicationRolloutNewResponseCurrentConfigurationInstanceTypeStandard3, ApplicationRolloutNewResponseCurrentConfigurationInstanceTypeStandard4:
		return true
	}
	return false
}

// Settings for deployment observability such as logging.
type ApplicationRolloutNewResponseCurrentConfigurationObservability struct {
	// Observability logging settings.
	Logs ApplicationRolloutNewResponseCurrentConfigurationObservabilityLogs `json:"logs"`
	JSON applicationRolloutNewResponseCurrentConfigurationObservabilityJSON `json:"-"`
}

// applicationRolloutNewResponseCurrentConfigurationObservabilityJSON contains the
// JSON metadata for the struct
// [ApplicationRolloutNewResponseCurrentConfigurationObservability]
type applicationRolloutNewResponseCurrentConfigurationObservabilityJSON struct {
	Logs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationRolloutNewResponseCurrentConfigurationObservability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationRolloutNewResponseCurrentConfigurationObservabilityJSON) RawJSON() string {
	return r.raw
}

// Observability logging settings.
type ApplicationRolloutNewResponseCurrentConfigurationObservabilityLogs struct {
	Enabled bool                                                                   `json:"enabled"`
	JSON    applicationRolloutNewResponseCurrentConfigurationObservabilityLogsJSON `json:"-"`
}

// applicationRolloutNewResponseCurrentConfigurationObservabilityLogsJSON contains
// the JSON metadata for the struct
// [ApplicationRolloutNewResponseCurrentConfigurationObservabilityLogs]
type applicationRolloutNewResponseCurrentConfigurationObservabilityLogsJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationRolloutNewResponseCurrentConfigurationObservabilityLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationRolloutNewResponseCurrentConfigurationObservabilityLogsJSON) RawJSON() string {
	return r.raw
}

type ApplicationRolloutNewResponseHealth struct {
	Errors []ApplicationRolloutNewResponseHealthError `json:"errors" api:"required"`
	// Shows a count of application instance states.
	Instances ApplicationRolloutNewResponseHealthInstances `json:"instances" api:"required"`
	// High-level health assessment. Only populated for "new_instances" strategy. Based
	// on a sample of target-version instances rather than a full count.
	//
	// - "pending": Zero target-version instances exist yet.
	// - "healthy": Every sampled target-version instance reports running or active.
	// - "degraded": Some sampled instances remain starting or scheduling.
	// - "unhealthy": One or more sampled instances have failed.
	Summary ApplicationRolloutNewResponseHealthSummary `json:"summary"`
	JSON    applicationRolloutNewResponseHealthJSON    `json:"-"`
}

// applicationRolloutNewResponseHealthJSON contains the JSON metadata for the
// struct [ApplicationRolloutNewResponseHealth]
type applicationRolloutNewResponseHealthJSON struct {
	Errors      apijson.Field
	Instances   apijson.Field
	Summary     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationRolloutNewResponseHealth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationRolloutNewResponseHealthJSON) RawJSON() string {
	return r.raw
}

type ApplicationRolloutNewResponseHealthError struct {
	// An event within a Placement or a Job.
	Event ApplicationRolloutNewResponseHealthErrorsEvent `json:"event" api:"required"`
	// An instance ID represents an identifier of an instance configuration that
	// maintains an underlying placement.
	InstanceID string                                       `json:"instance_id" api:"required"`
	JSON       applicationRolloutNewResponseHealthErrorJSON `json:"-"`
}

// applicationRolloutNewResponseHealthErrorJSON contains the JSON metadata for the
// struct [ApplicationRolloutNewResponseHealthError]
type applicationRolloutNewResponseHealthErrorJSON struct {
	Event       apijson.Field
	InstanceID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationRolloutNewResponseHealthError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationRolloutNewResponseHealthErrorJSON) RawJSON() string {
	return r.raw
}

// An event within a Placement or a Job.
type ApplicationRolloutNewResponseHealthErrorsEvent struct {
	ID      string                 `json:"id" api:"required"`
	Details map[string]interface{} `json:"details" api:"required"`
	Message string                 `json:"message" api:"required"`
	// Name of the event that describes the kind event that happened.
	//
	// - SchedulerPlaced: It's the first event that creates a container placement. It
	//   happens when the Containers runtime was able to retrieve deployment resources
	//   and start verifying everything is correct.
	// - NetworkingIPAssigned: It's sent when the Containers runtime maps the IP to the
	//   container.
	// - VMStarted: It's sent when the Containers runtime starts the VM. The container
	//   might remain unhealthy at this point.
	// - ImagePulled: It's sent when the Containers runtime pulls the image
	//   successfully.
	// - ImagePullError: It's sent when the Containers runtime is having issues pulling
	//   the image. The message and details have more information on what happened for
	//   debugging.
	// - VMFailedToStart: It's sent when the Containers runtime was unable to boot the
	//   VM.
	// - VMStopping: It's sent when the scheduler is stopping the VM.
	// - VMStopped: It's sent when the VM finally exits.
	// - VMFailed: It's sent when the scheduling of the VM failed in the current
	//   location.
	// - RuntimeStartFailed: It's sent when the runtime hits an internal error.
	// - SSHStarted: It's sent when the container gains network connectivity and opens
	//   the SSH port. Containers only send this event when SSH keys exist.
	// - CheckUpdate: Sent when the status of a health or readiness check changes. This
	//   may also affect the health status of the placement.
	// - DurableObjectConnected: Sent when a durable object instance connects and gains
	//   control of the deployment. This event is only sent for durable object
	//   deployments. It is sent after VMStarted.
	// - ContainerStarted: It's sent when the container starts running.
	Name         ApplicationRolloutNewResponseHealthErrorsEventName `json:"name" api:"required"`
	StatusChange map[string]interface{}                             `json:"statusChange" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	Time string                                             `json:"time" api:"required"`
	Type ApplicationRolloutNewResponseHealthErrorsEventType `json:"type" api:"required"`
	JSON applicationRolloutNewResponseHealthErrorsEventJSON `json:"-"`
}

// applicationRolloutNewResponseHealthErrorsEventJSON contains the JSON metadata
// for the struct [ApplicationRolloutNewResponseHealthErrorsEvent]
type applicationRolloutNewResponseHealthErrorsEventJSON struct {
	ID           apijson.Field
	Details      apijson.Field
	Message      apijson.Field
	Name         apijson.Field
	StatusChange apijson.Field
	Time         apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ApplicationRolloutNewResponseHealthErrorsEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationRolloutNewResponseHealthErrorsEventJSON) RawJSON() string {
	return r.raw
}

// Name of the event that describes the kind event that happened.
//
//   - SchedulerPlaced: It's the first event that creates a container placement. It
//     happens when the Containers runtime was able to retrieve deployment resources
//     and start verifying everything is correct.
//   - NetworkingIPAssigned: It's sent when the Containers runtime maps the IP to the
//     container.
//   - VMStarted: It's sent when the Containers runtime starts the VM. The container
//     might remain unhealthy at this point.
//   - ImagePulled: It's sent when the Containers runtime pulls the image
//     successfully.
//   - ImagePullError: It's sent when the Containers runtime is having issues pulling
//     the image. The message and details have more information on what happened for
//     debugging.
//   - VMFailedToStart: It's sent when the Containers runtime was unable to boot the
//     VM.
//   - VMStopping: It's sent when the scheduler is stopping the VM.
//   - VMStopped: It's sent when the VM finally exits.
//   - VMFailed: It's sent when the scheduling of the VM failed in the current
//     location.
//   - RuntimeStartFailed: It's sent when the runtime hits an internal error.
//   - SSHStarted: It's sent when the container gains network connectivity and opens
//     the SSH port. Containers only send this event when SSH keys exist.
//   - CheckUpdate: Sent when the status of a health or readiness check changes. This
//     may also affect the health status of the placement.
//   - DurableObjectConnected: Sent when a durable object instance connects and gains
//     control of the deployment. This event is only sent for durable object
//     deployments. It is sent after VMStarted.
//   - ContainerStarted: It's sent when the container starts running.
type ApplicationRolloutNewResponseHealthErrorsEventName string

const (
	ApplicationRolloutNewResponseHealthErrorsEventNameSchedulerPlaced              ApplicationRolloutNewResponseHealthErrorsEventName = "SchedulerPlaced"
	ApplicationRolloutNewResponseHealthErrorsEventNameNetworkingIPAssigned         ApplicationRolloutNewResponseHealthErrorsEventName = "NetworkingIPAssigned"
	ApplicationRolloutNewResponseHealthErrorsEventNameVmStarted                    ApplicationRolloutNewResponseHealthErrorsEventName = "VMStarted"
	ApplicationRolloutNewResponseHealthErrorsEventNameImagePulled                  ApplicationRolloutNewResponseHealthErrorsEventName = "ImagePulled"
	ApplicationRolloutNewResponseHealthErrorsEventNameImagePullError               ApplicationRolloutNewResponseHealthErrorsEventName = "ImagePullError"
	ApplicationRolloutNewResponseHealthErrorsEventNameVmFailedToStart              ApplicationRolloutNewResponseHealthErrorsEventName = "VMFailedToStart"
	ApplicationRolloutNewResponseHealthErrorsEventNameNetworkingIPAssignmentFailed ApplicationRolloutNewResponseHealthErrorsEventName = "NetworkingIPAssignmentFailed"
	ApplicationRolloutNewResponseHealthErrorsEventNameVmRunning                    ApplicationRolloutNewResponseHealthErrorsEventName = "VMRunning"
	ApplicationRolloutNewResponseHealthErrorsEventNameVmStopping                   ApplicationRolloutNewResponseHealthErrorsEventName = "VMStopping"
	ApplicationRolloutNewResponseHealthErrorsEventNameVmStopped                    ApplicationRolloutNewResponseHealthErrorsEventName = "VMStopped"
	ApplicationRolloutNewResponseHealthErrorsEventNameVmFailed                     ApplicationRolloutNewResponseHealthErrorsEventName = "VMFailed"
	ApplicationRolloutNewResponseHealthErrorsEventNameRuntimeStartFailed           ApplicationRolloutNewResponseHealthErrorsEventName = "RuntimeStartFailed"
	ApplicationRolloutNewResponseHealthErrorsEventNameSSHStarted                   ApplicationRolloutNewResponseHealthErrorsEventName = "SSHStarted"
	ApplicationRolloutNewResponseHealthErrorsEventNameServiceHealthUpdates         ApplicationRolloutNewResponseHealthErrorsEventName = "ServiceHealthUpdates"
	ApplicationRolloutNewResponseHealthErrorsEventNameCheckUpdate                  ApplicationRolloutNewResponseHealthErrorsEventName = "CheckUpdate"
	ApplicationRolloutNewResponseHealthErrorsEventNameDurableObjectConnected       ApplicationRolloutNewResponseHealthErrorsEventName = "DurableObjectConnected"
	ApplicationRolloutNewResponseHealthErrorsEventNameContainerStarted             ApplicationRolloutNewResponseHealthErrorsEventName = "ContainerStarted"
)

func (r ApplicationRolloutNewResponseHealthErrorsEventName) IsKnown() bool {
	switch r {
	case ApplicationRolloutNewResponseHealthErrorsEventNameSchedulerPlaced, ApplicationRolloutNewResponseHealthErrorsEventNameNetworkingIPAssigned, ApplicationRolloutNewResponseHealthErrorsEventNameVmStarted, ApplicationRolloutNewResponseHealthErrorsEventNameImagePulled, ApplicationRolloutNewResponseHealthErrorsEventNameImagePullError, ApplicationRolloutNewResponseHealthErrorsEventNameVmFailedToStart, ApplicationRolloutNewResponseHealthErrorsEventNameNetworkingIPAssignmentFailed, ApplicationRolloutNewResponseHealthErrorsEventNameVmRunning, ApplicationRolloutNewResponseHealthErrorsEventNameVmStopping, ApplicationRolloutNewResponseHealthErrorsEventNameVmStopped, ApplicationRolloutNewResponseHealthErrorsEventNameVmFailed, ApplicationRolloutNewResponseHealthErrorsEventNameRuntimeStartFailed, ApplicationRolloutNewResponseHealthErrorsEventNameSSHStarted, ApplicationRolloutNewResponseHealthErrorsEventNameServiceHealthUpdates, ApplicationRolloutNewResponseHealthErrorsEventNameCheckUpdate, ApplicationRolloutNewResponseHealthErrorsEventNameDurableObjectConnected, ApplicationRolloutNewResponseHealthErrorsEventNameContainerStarted:
		return true
	}
	return false
}

type ApplicationRolloutNewResponseHealthErrorsEventType string

const (
	ApplicationRolloutNewResponseHealthErrorsEventTypeInfo        ApplicationRolloutNewResponseHealthErrorsEventType = "Info"
	ApplicationRolloutNewResponseHealthErrorsEventTypeError       ApplicationRolloutNewResponseHealthErrorsEventType = "Error"
	ApplicationRolloutNewResponseHealthErrorsEventTypeWarn        ApplicationRolloutNewResponseHealthErrorsEventType = "Warn"
	ApplicationRolloutNewResponseHealthErrorsEventTypeUserError   ApplicationRolloutNewResponseHealthErrorsEventType = "UserError"
	ApplicationRolloutNewResponseHealthErrorsEventTypeSystemError ApplicationRolloutNewResponseHealthErrorsEventType = "SystemError"
)

func (r ApplicationRolloutNewResponseHealthErrorsEventType) IsKnown() bool {
	switch r {
	case ApplicationRolloutNewResponseHealthErrorsEventTypeInfo, ApplicationRolloutNewResponseHealthErrorsEventTypeError, ApplicationRolloutNewResponseHealthErrorsEventTypeWarn, ApplicationRolloutNewResponseHealthErrorsEventTypeUserError, ApplicationRolloutNewResponseHealthErrorsEventTypeSystemError:
		return true
	}
	return false
}

// Shows a count of application instance states.
type ApplicationRolloutNewResponseHealthInstances struct {
	// Number of instances whose runtime reports the container as running
	// (container_status = "running"). This is a subset of the placements that remain
	// up: an instance that is already bound to a Durable Object and serving traffic is
	// counted under "assigned" until its container_status catches up to "running", so
	// container_status can briefly lag Durable Object attachment under churn. To
	// estimate running, Durable-Object-bound instances, sum "active" + "assigned"
	// rather than reading "active" alone.
	Active int64 `json:"active" api:"required"`
	// Number of instances bound to a Durable Object with a running placement whose
	// container_status remains behind "running". These count as live, serving
	// instances; "active" + "assigned" approximates the running, Durable-Object-bound
	// count.
	Assigned int64                                            `json:"assigned" api:"required"`
	JSON     applicationRolloutNewResponseHealthInstancesJSON `json:"-"`
}

// applicationRolloutNewResponseHealthInstancesJSON contains the JSON metadata for
// the struct [ApplicationRolloutNewResponseHealthInstances]
type applicationRolloutNewResponseHealthInstancesJSON struct {
	Active      apijson.Field
	Assigned    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationRolloutNewResponseHealthInstances) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationRolloutNewResponseHealthInstancesJSON) RawJSON() string {
	return r.raw
}

// High-level health assessment. Only populated for "new_instances" strategy. Based
// on a sample of target-version instances rather than a full count.
//
// - "pending": Zero target-version instances exist yet.
// - "healthy": Every sampled target-version instance reports running or active.
// - "degraded": Some sampled instances remain starting or scheduling.
// - "unhealthy": One or more sampled instances have failed.
type ApplicationRolloutNewResponseHealthSummary string

const (
	ApplicationRolloutNewResponseHealthSummaryHealthy   ApplicationRolloutNewResponseHealthSummary = "healthy"
	ApplicationRolloutNewResponseHealthSummaryDegraded  ApplicationRolloutNewResponseHealthSummary = "degraded"
	ApplicationRolloutNewResponseHealthSummaryUnhealthy ApplicationRolloutNewResponseHealthSummary = "unhealthy"
	ApplicationRolloutNewResponseHealthSummaryPending   ApplicationRolloutNewResponseHealthSummary = "pending"
)

func (r ApplicationRolloutNewResponseHealthSummary) IsKnown() bool {
	switch r {
	case ApplicationRolloutNewResponseHealthSummaryHealthy, ApplicationRolloutNewResponseHealthSummaryDegraded, ApplicationRolloutNewResponseHealthSummaryUnhealthy, ApplicationRolloutNewResponseHealthSummaryPending:
		return true
	}
	return false
}

// Kind of the rollout process.
//
//   - "full_auto": For rolling rollouts, starts progressing steps upon rollout
//     creation. For new_instances rollouts, advances percentage targets
//     automatically after target-version health is observed.
//   - "full_manual": Requires manually progressing each step in the rollout using
//     the UpdateRollout's action paramater.
//   - "durable_objects_auto": Default when the application is a DO application.
type ApplicationRolloutNewResponseKind string

const (
	ApplicationRolloutNewResponseKindFullAuto           ApplicationRolloutNewResponseKind = "full_auto"
	ApplicationRolloutNewResponseKindFullManual         ApplicationRolloutNewResponseKind = "full_manual"
	ApplicationRolloutNewResponseKindDurableObjectsAuto ApplicationRolloutNewResponseKind = "durable_objects_auto"
)

func (r ApplicationRolloutNewResponseKind) IsKnown() bool {
	switch r {
	case ApplicationRolloutNewResponseKindFullAuto, ApplicationRolloutNewResponseKindFullManual, ApplicationRolloutNewResponseKindDurableObjectsAuto:
		return true
	}
	return false
}

// Current status of the rollout.
type ApplicationRolloutNewResponseStatus string

const (
	ApplicationRolloutNewResponseStatusPending     ApplicationRolloutNewResponseStatus = "pending"
	ApplicationRolloutNewResponseStatusProgressing ApplicationRolloutNewResponseStatus = "progressing"
	ApplicationRolloutNewResponseStatusCompleted   ApplicationRolloutNewResponseStatus = "completed"
	ApplicationRolloutNewResponseStatusReverted    ApplicationRolloutNewResponseStatus = "reverted"
	ApplicationRolloutNewResponseStatusReplaced    ApplicationRolloutNewResponseStatus = "replaced"
)

func (r ApplicationRolloutNewResponseStatus) IsKnown() bool {
	switch r {
	case ApplicationRolloutNewResponseStatusPending, ApplicationRolloutNewResponseStatusProgressing, ApplicationRolloutNewResponseStatusCompleted, ApplicationRolloutNewResponseStatusReverted, ApplicationRolloutNewResponseStatusReplaced:
		return true
	}
	return false
}

// The rollout strategy.
//
//   - "rolling": Step-based rollout with health gates. Actively replaces instances
//     to reach each step's target percentage. Response includes steps and progress.
//   - "new_instances": Percentage control over version distribution. Version sync
//     actively replaces instances to match the configured percentage. "full_auto"
//     ramps through fixed percentage targets after target-version health is
//     observed. Response includes percentage, version_distribution, and
//     health.summary.
type ApplicationRolloutNewResponseStrategy string

const (
	ApplicationRolloutNewResponseStrategyRolling      ApplicationRolloutNewResponseStrategy = "rolling"
	ApplicationRolloutNewResponseStrategyNewInstances ApplicationRolloutNewResponseStrategy = "new_instances"
)

func (r ApplicationRolloutNewResponseStrategy) IsKnown() bool {
	switch r {
	case ApplicationRolloutNewResponseStrategyRolling, ApplicationRolloutNewResponseStrategyNewInstances:
		return true
	}
	return false
}

// User-specified container configuration changes.
type ApplicationRolloutNewResponseTargetConfiguration struct {
	AuthorizedKeys []ApplicationRolloutNewResponseTargetConfigurationAuthorizedKey `json:"authorized_keys"`
	// The command that runs when the container starts, passed to the entrypoint. You
	// can override this at run-time. If you override only the command, it gets passed
	// to the default entrypoint specified in the image.
	Command []string `json:"command"`
	// The entry point for the container, specifying the executable to run when the
	// container starts. You can override this at run-time. If you do, the default
	// command from the image is ignored. Specify both entrypoint and command at
	// run-time to completely replace the image defaults.
	Entrypoint []string `json:"entrypoint"`
	// Container environment variables.
	EnvironmentVariables []ApplicationRolloutNewResponseTargetConfigurationEnvironmentVariable `json:"environment_variables"`
	// Image url.
	Image string `json:"image"`
	// The instance type configures vCPU, memory, and disk.
	//
	// - "lite": 1/16 vCPU, 256 MiB memory, 2 GB disk
	// - "basic": 1/4 vCPU, 1 GiB memory, 4 GB disk
	// - "standard-1": 1/2 vCPU, 4 GiB memory, 8 GB disk
	// - "standard-2": 1 vCPU, 6 GiB memory, 12 GB disk
	// - "standard-3": 2 vCPU, 8 GiB memory, 16 GB disk
	// - "standard-4": 4 vCPU, 12 GiB memory, 20 GB disk
	InstanceType ApplicationRolloutNewResponseTargetConfigurationInstanceType `json:"instance_type"`
	// Settings for deployment observability such as logging.
	Observability ApplicationRolloutNewResponseTargetConfigurationObservability `json:"observability"`
	JSON          applicationRolloutNewResponseTargetConfigurationJSON          `json:"-"`
}

// applicationRolloutNewResponseTargetConfigurationJSON contains the JSON metadata
// for the struct [ApplicationRolloutNewResponseTargetConfiguration]
type applicationRolloutNewResponseTargetConfigurationJSON struct {
	AuthorizedKeys       apijson.Field
	Command              apijson.Field
	Entrypoint           apijson.Field
	EnvironmentVariables apijson.Field
	Image                apijson.Field
	InstanceType         apijson.Field
	Observability        apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ApplicationRolloutNewResponseTargetConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationRolloutNewResponseTargetConfigurationJSON) RawJSON() string {
	return r.raw
}

// User-provided SSH public key.
type ApplicationRolloutNewResponseTargetConfigurationAuthorizedKey struct {
	// An SSH public key.
	PublicKey string `json:"public_key" api:"required"`
	// Optional human readable name for this key.
	Name string                                                            `json:"name"`
	JSON applicationRolloutNewResponseTargetConfigurationAuthorizedKeyJSON `json:"-"`
}

// applicationRolloutNewResponseTargetConfigurationAuthorizedKeyJSON contains the
// JSON metadata for the struct
// [ApplicationRolloutNewResponseTargetConfigurationAuthorizedKey]
type applicationRolloutNewResponseTargetConfigurationAuthorizedKeyJSON struct {
	PublicKey   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationRolloutNewResponseTargetConfigurationAuthorizedKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationRolloutNewResponseTargetConfigurationAuthorizedKeyJSON) RawJSON() string {
	return r.raw
}

// An environment variable with a value set.
type ApplicationRolloutNewResponseTargetConfigurationEnvironmentVariable struct {
	// An environment variable name.
	Name string `json:"name" api:"required"`
	// An environment variable value.
	Value string                                                                  `json:"value" api:"required"`
	JSON  applicationRolloutNewResponseTargetConfigurationEnvironmentVariableJSON `json:"-"`
}

// applicationRolloutNewResponseTargetConfigurationEnvironmentVariableJSON contains
// the JSON metadata for the struct
// [ApplicationRolloutNewResponseTargetConfigurationEnvironmentVariable]
type applicationRolloutNewResponseTargetConfigurationEnvironmentVariableJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationRolloutNewResponseTargetConfigurationEnvironmentVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationRolloutNewResponseTargetConfigurationEnvironmentVariableJSON) RawJSON() string {
	return r.raw
}

// The instance type configures vCPU, memory, and disk.
//
// - "lite": 1/16 vCPU, 256 MiB memory, 2 GB disk
// - "basic": 1/4 vCPU, 1 GiB memory, 4 GB disk
// - "standard-1": 1/2 vCPU, 4 GiB memory, 8 GB disk
// - "standard-2": 1 vCPU, 6 GiB memory, 12 GB disk
// - "standard-3": 2 vCPU, 8 GiB memory, 16 GB disk
// - "standard-4": 4 vCPU, 12 GiB memory, 20 GB disk
type ApplicationRolloutNewResponseTargetConfigurationInstanceType string

const (
	ApplicationRolloutNewResponseTargetConfigurationInstanceTypeLite      ApplicationRolloutNewResponseTargetConfigurationInstanceType = "lite"
	ApplicationRolloutNewResponseTargetConfigurationInstanceTypeBasic     ApplicationRolloutNewResponseTargetConfigurationInstanceType = "basic"
	ApplicationRolloutNewResponseTargetConfigurationInstanceTypeStandard1 ApplicationRolloutNewResponseTargetConfigurationInstanceType = "standard-1"
	ApplicationRolloutNewResponseTargetConfigurationInstanceTypeStandard2 ApplicationRolloutNewResponseTargetConfigurationInstanceType = "standard-2"
	ApplicationRolloutNewResponseTargetConfigurationInstanceTypeStandard3 ApplicationRolloutNewResponseTargetConfigurationInstanceType = "standard-3"
	ApplicationRolloutNewResponseTargetConfigurationInstanceTypeStandard4 ApplicationRolloutNewResponseTargetConfigurationInstanceType = "standard-4"
)

func (r ApplicationRolloutNewResponseTargetConfigurationInstanceType) IsKnown() bool {
	switch r {
	case ApplicationRolloutNewResponseTargetConfigurationInstanceTypeLite, ApplicationRolloutNewResponseTargetConfigurationInstanceTypeBasic, ApplicationRolloutNewResponseTargetConfigurationInstanceTypeStandard1, ApplicationRolloutNewResponseTargetConfigurationInstanceTypeStandard2, ApplicationRolloutNewResponseTargetConfigurationInstanceTypeStandard3, ApplicationRolloutNewResponseTargetConfigurationInstanceTypeStandard4:
		return true
	}
	return false
}

// Settings for deployment observability such as logging.
type ApplicationRolloutNewResponseTargetConfigurationObservability struct {
	// Observability logging settings.
	Logs ApplicationRolloutNewResponseTargetConfigurationObservabilityLogs `json:"logs"`
	JSON applicationRolloutNewResponseTargetConfigurationObservabilityJSON `json:"-"`
}

// applicationRolloutNewResponseTargetConfigurationObservabilityJSON contains the
// JSON metadata for the struct
// [ApplicationRolloutNewResponseTargetConfigurationObservability]
type applicationRolloutNewResponseTargetConfigurationObservabilityJSON struct {
	Logs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationRolloutNewResponseTargetConfigurationObservability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationRolloutNewResponseTargetConfigurationObservabilityJSON) RawJSON() string {
	return r.raw
}

// Observability logging settings.
type ApplicationRolloutNewResponseTargetConfigurationObservabilityLogs struct {
	Enabled bool                                                                  `json:"enabled"`
	JSON    applicationRolloutNewResponseTargetConfigurationObservabilityLogsJSON `json:"-"`
}

// applicationRolloutNewResponseTargetConfigurationObservabilityLogsJSON contains
// the JSON metadata for the struct
// [ApplicationRolloutNewResponseTargetConfigurationObservabilityLogs]
type applicationRolloutNewResponseTargetConfigurationObservabilityLogsJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationRolloutNewResponseTargetConfigurationObservabilityLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationRolloutNewResponseTargetConfigurationObservabilityLogsJSON) RawJSON() string {
	return r.raw
}

// Progress details of an application rollout.
type ApplicationRolloutNewResponseProgress struct {
	// Current step being executed in the rollout process. Initialized to 0.
	CurrentStep int64 `json:"current_step" api:"required"`
	// Total number of instances the rollout affects.
	TotalInstances int64 `json:"total_instances" api:"required"`
	// Total number of steps in the rollout.
	TotalSteps int64 `json:"total_steps" api:"required"`
	// Number of instances updated in the rollout process.
	UpdatedInstances int64 `json:"updated_instances" api:"required"`
	// Expected distribution of instances per version, based on the current percentage
	// split. Populated during active rollouts. Values derive from the version
	// percentage weights rather than actual running instance counts.
	VersionDistribution ApplicationRolloutNewResponseProgressVersionDistribution `json:"version_distribution"`
	JSON                applicationRolloutNewResponseProgressJSON                `json:"-"`
}

// applicationRolloutNewResponseProgressJSON contains the JSON metadata for the
// struct [ApplicationRolloutNewResponseProgress]
type applicationRolloutNewResponseProgressJSON struct {
	CurrentStep         apijson.Field
	TotalInstances      apijson.Field
	TotalSteps          apijson.Field
	UpdatedInstances    apijson.Field
	VersionDistribution apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ApplicationRolloutNewResponseProgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationRolloutNewResponseProgressJSON) RawJSON() string {
	return r.raw
}

// Expected distribution of instances per version, based on the current percentage
// split. Populated during active rollouts. Values derive from the version
// percentage weights rather than actual running instance counts.
type ApplicationRolloutNewResponseProgressVersionDistribution struct {
	// Expected number of instances remaining on the current (old) version based on the
	// current percentage split. Only populated for "rolling" strategy.
	CurrentVersionInstances int64 `json:"current_version_instances"`
	// The percentage of new instances being scheduled on the current version (100 -
	// target_version_percentage). Only populated for "new_instances" strategy.
	CurrentVersionPercentage int64 `json:"current_version_percentage"`
	// Expected number of instances scheduled for the target (new) version based on the
	// current percentage split. Only populated for "rolling" strategy.
	TargetVersionInstances int64 `json:"target_version_instances"`
	// The active percentage of new instances being scheduled on the target version.
	// For "rolling", this reflects the step_size.percentage of the current active
	// step. For "new_instances", this reflects the user-set percentage.
	TargetVersionPercentage int64                                                        `json:"target_version_percentage"`
	JSON                    applicationRolloutNewResponseProgressVersionDistributionJSON `json:"-"`
}

// applicationRolloutNewResponseProgressVersionDistributionJSON contains the JSON
// metadata for the struct
// [ApplicationRolloutNewResponseProgressVersionDistribution]
type applicationRolloutNewResponseProgressVersionDistributionJSON struct {
	CurrentVersionInstances  apijson.Field
	CurrentVersionPercentage apijson.Field
	TargetVersionInstances   apijson.Field
	TargetVersionPercentage  apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *ApplicationRolloutNewResponseProgressVersionDistribution) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationRolloutNewResponseProgressVersionDistributionJSON) RawJSON() string {
	return r.raw
}

// Steps within the rollout process.
type ApplicationRolloutNewResponseStep struct {
	// The sequential order of the rollout step, automatically assigned starting from
	// 1, based on the total number of steps in the rollout process.
	ID int64 `json:"id" api:"required"`
	// Description of the rollout step.
	Description string `json:"description" api:"required"`
	// Status of the rollout step.
	Status   ApplicationRolloutNewResponseStepsStatus   `json:"status" api:"required"`
	StepSize ApplicationRolloutNewResponseStepsStepSize `json:"step_size" api:"required"`
	// UTC timestamp string in ISO 8601 format.
	CompletedAt string `json:"completed_at"`
	// Reason for the step's current status.
	Reason string `json:"reason"`
	// UTC timestamp string in ISO 8601 format.
	StartedAt string                                `json:"started_at"`
	JSON      applicationRolloutNewResponseStepJSON `json:"-"`
}

// applicationRolloutNewResponseStepJSON contains the JSON metadata for the struct
// [ApplicationRolloutNewResponseStep]
type applicationRolloutNewResponseStepJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Status      apijson.Field
	StepSize    apijson.Field
	CompletedAt apijson.Field
	Reason      apijson.Field
	StartedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationRolloutNewResponseStep) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationRolloutNewResponseStepJSON) RawJSON() string {
	return r.raw
}

// Status of the rollout step.
type ApplicationRolloutNewResponseStepsStatus string

const (
	ApplicationRolloutNewResponseStepsStatusPending     ApplicationRolloutNewResponseStepsStatus = "pending"
	ApplicationRolloutNewResponseStepsStatusProgressing ApplicationRolloutNewResponseStepsStatus = "progressing"
	ApplicationRolloutNewResponseStepsStatusReverting   ApplicationRolloutNewResponseStepsStatus = "reverting"
	ApplicationRolloutNewResponseStepsStatusCompleted   ApplicationRolloutNewResponseStepsStatus = "completed"
	ApplicationRolloutNewResponseStepsStatusReverted    ApplicationRolloutNewResponseStepsStatus = "reverted"
)

func (r ApplicationRolloutNewResponseStepsStatus) IsKnown() bool {
	switch r {
	case ApplicationRolloutNewResponseStepsStatusPending, ApplicationRolloutNewResponseStepsStatusProgressing, ApplicationRolloutNewResponseStepsStatusReverting, ApplicationRolloutNewResponseStepsStatusCompleted, ApplicationRolloutNewResponseStepsStatusReverted:
		return true
	}
	return false
}

type ApplicationRolloutNewResponseStepsStepSize struct {
	// Percentage of instances affected in this step. Min 10% and Max 100%.
	Percentage int64                                          `json:"percentage" api:"required"`
	JSON       applicationRolloutNewResponseStepsStepSizeJSON `json:"-"`
}

// applicationRolloutNewResponseStepsStepSizeJSON contains the JSON metadata for
// the struct [ApplicationRolloutNewResponseStepsStepSize]
type applicationRolloutNewResponseStepsStepSizeJSON struct {
	Percentage  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationRolloutNewResponseStepsStepSize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationRolloutNewResponseStepsStepSizeJSON) RawJSON() string {
	return r.raw
}

// Version percentage distribution. Only present for "new_instances" strategy. For
// "rolling" strategy, see progress.version_distribution instead.
type ApplicationRolloutNewResponseVersionDistribution struct {
	// Percentage of instances on the current (old) version.
	CurrentVersionPercentage int64 `json:"current_version_percentage" api:"required"`
	// Percentage of instances on the target (new) version.
	TargetVersionPercentage int64                                                `json:"target_version_percentage" api:"required"`
	JSON                    applicationRolloutNewResponseVersionDistributionJSON `json:"-"`
}

// applicationRolloutNewResponseVersionDistributionJSON contains the JSON metadata
// for the struct [ApplicationRolloutNewResponseVersionDistribution]
type applicationRolloutNewResponseVersionDistributionJSON struct {
	CurrentVersionPercentage apijson.Field
	TargetVersionPercentage  apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *ApplicationRolloutNewResponseVersionDistribution) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationRolloutNewResponseVersionDistributionJSON) RawJSON() string {
	return r.raw
}

type ApplicationRolloutNewParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Description of the rollout process.
	Description param.Field[string] `json:"description" api:"required"`
	// Strategy used for the rollout.
	//
	// - "rolling": Step-based rollout with health gates. Actively replaces instances
	//   to reach each step's target percentage.
	// - "new_instances": Percentage control over version distribution. Version sync
	//   actively replaces instances to match the configured percentage. The
	//   "full_auto" kind advances through fixed percentage targets after
	//   target-version health is observed.
	Strategy param.Field[ApplicationRolloutNewParamsStrategy] `json:"strategy" api:"required"`
	// User-specified container configuration changes.
	TargetConfiguration param.Field[ApplicationRolloutNewParamsTargetConfiguration] `json:"target_configuration" api:"required"`
	// Kind of the rollout process. Defaults to "full_auto".
	//
	// - "full_auto": For rolling rollouts, starts progressing steps upon rollout
	//   creation. For new_instances rollouts, advances percentage targets
	//   automatically after target-version health is observed.
	// - "full_manual": Requires manually progressing each step in the rollout using
	//   the UpdateRollout's action parameter.
	Kind param.Field[ApplicationRolloutNewParamsKind] `json:"kind"`
	// Initial target version percentage (0-100). Version sync actively replaces
	// instances to match. Required when strategy is "new_instances" and kind is
	// "full_manual". When strategy is "new_instances" and kind is "full_auto", omitted
	// percentage starts at 10% or the smallest percentage that targets at least one
	// instance. Unused for "rolling".
	Percentage param.Field[int64] `json:"percentage"`
	// Percentage of rollout to increase in each step when "steps" is absent.
	// Applicable values: 5, 10, 20, 25, 50, 100. These create rollouts with 20, 10, 5,
	// 4, 2, 1 steps respectively. Only valid for "rolling" strategy.
	StepPercentage param.Field[ApplicationRolloutNewParamsStepPercentage] `json:"step_percentage"`
	// Steps defining the rollout process, used when "step_percentage" is absent.
	// Specify only one of "step_percentage" or "steps" when creating a rollout.
	// "steps" allow granular control over each step. Only valid for "rolling"
	// strategy.
	Steps param.Field[[]ApplicationRolloutNewParamsStep] `json:"steps"`
}

func (r ApplicationRolloutNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Strategy used for the rollout.
//
//   - "rolling": Step-based rollout with health gates. Actively replaces instances
//     to reach each step's target percentage.
//   - "new_instances": Percentage control over version distribution. Version sync
//     actively replaces instances to match the configured percentage. The
//     "full_auto" kind advances through fixed percentage targets after
//     target-version health is observed.
type ApplicationRolloutNewParamsStrategy string

const (
	ApplicationRolloutNewParamsStrategyRolling      ApplicationRolloutNewParamsStrategy = "rolling"
	ApplicationRolloutNewParamsStrategyNewInstances ApplicationRolloutNewParamsStrategy = "new_instances"
)

func (r ApplicationRolloutNewParamsStrategy) IsKnown() bool {
	switch r {
	case ApplicationRolloutNewParamsStrategyRolling, ApplicationRolloutNewParamsStrategyNewInstances:
		return true
	}
	return false
}

// User-specified container configuration changes.
type ApplicationRolloutNewParamsTargetConfiguration struct {
	AuthorizedKeys param.Field[[]ApplicationRolloutNewParamsTargetConfigurationAuthorizedKey] `json:"authorized_keys"`
	// The command that runs when the container starts, passed to the entrypoint. You
	// can override this at run-time. If you override only the command, it gets passed
	// to the default entrypoint specified in the image.
	Command param.Field[[]string] `json:"command"`
	// The entry point for the container, specifying the executable to run when the
	// container starts. You can override this at run-time. If you do, the default
	// command from the image is ignored. Specify both entrypoint and command at
	// run-time to completely replace the image defaults.
	Entrypoint param.Field[[]string] `json:"entrypoint"`
	// Container environment variables.
	EnvironmentVariables param.Field[[]ApplicationRolloutNewParamsTargetConfigurationEnvironmentVariable] `json:"environment_variables"`
	// Image url.
	Image param.Field[string] `json:"image"`
	// The instance type configures vCPU, memory, and disk.
	//
	// - "lite": 1/16 vCPU, 256 MiB memory, 2 GB disk
	// - "basic": 1/4 vCPU, 1 GiB memory, 4 GB disk
	// - "standard-1": 1/2 vCPU, 4 GiB memory, 8 GB disk
	// - "standard-2": 1 vCPU, 6 GiB memory, 12 GB disk
	// - "standard-3": 2 vCPU, 8 GiB memory, 16 GB disk
	// - "standard-4": 4 vCPU, 12 GiB memory, 20 GB disk
	InstanceType param.Field[ApplicationRolloutNewParamsTargetConfigurationInstanceType] `json:"instance_type"`
	// Settings for deployment observability such as logging.
	Observability param.Field[ApplicationRolloutNewParamsTargetConfigurationObservability] `json:"observability"`
}

func (r ApplicationRolloutNewParamsTargetConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// User-provided SSH public key.
type ApplicationRolloutNewParamsTargetConfigurationAuthorizedKey struct {
	// An SSH public key.
	PublicKey param.Field[string] `json:"public_key" api:"required"`
	// Optional human readable name for this key.
	Name param.Field[string] `json:"name"`
}

func (r ApplicationRolloutNewParamsTargetConfigurationAuthorizedKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An environment variable with a value set.
type ApplicationRolloutNewParamsTargetConfigurationEnvironmentVariable struct {
	// An environment variable name.
	Name param.Field[string] `json:"name" api:"required"`
	// An environment variable value.
	Value param.Field[string] `json:"value" api:"required"`
}

func (r ApplicationRolloutNewParamsTargetConfigurationEnvironmentVariable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The instance type configures vCPU, memory, and disk.
//
// - "lite": 1/16 vCPU, 256 MiB memory, 2 GB disk
// - "basic": 1/4 vCPU, 1 GiB memory, 4 GB disk
// - "standard-1": 1/2 vCPU, 4 GiB memory, 8 GB disk
// - "standard-2": 1 vCPU, 6 GiB memory, 12 GB disk
// - "standard-3": 2 vCPU, 8 GiB memory, 16 GB disk
// - "standard-4": 4 vCPU, 12 GiB memory, 20 GB disk
type ApplicationRolloutNewParamsTargetConfigurationInstanceType string

const (
	ApplicationRolloutNewParamsTargetConfigurationInstanceTypeLite      ApplicationRolloutNewParamsTargetConfigurationInstanceType = "lite"
	ApplicationRolloutNewParamsTargetConfigurationInstanceTypeBasic     ApplicationRolloutNewParamsTargetConfigurationInstanceType = "basic"
	ApplicationRolloutNewParamsTargetConfigurationInstanceTypeStandard1 ApplicationRolloutNewParamsTargetConfigurationInstanceType = "standard-1"
	ApplicationRolloutNewParamsTargetConfigurationInstanceTypeStandard2 ApplicationRolloutNewParamsTargetConfigurationInstanceType = "standard-2"
	ApplicationRolloutNewParamsTargetConfigurationInstanceTypeStandard3 ApplicationRolloutNewParamsTargetConfigurationInstanceType = "standard-3"
	ApplicationRolloutNewParamsTargetConfigurationInstanceTypeStandard4 ApplicationRolloutNewParamsTargetConfigurationInstanceType = "standard-4"
)

func (r ApplicationRolloutNewParamsTargetConfigurationInstanceType) IsKnown() bool {
	switch r {
	case ApplicationRolloutNewParamsTargetConfigurationInstanceTypeLite, ApplicationRolloutNewParamsTargetConfigurationInstanceTypeBasic, ApplicationRolloutNewParamsTargetConfigurationInstanceTypeStandard1, ApplicationRolloutNewParamsTargetConfigurationInstanceTypeStandard2, ApplicationRolloutNewParamsTargetConfigurationInstanceTypeStandard3, ApplicationRolloutNewParamsTargetConfigurationInstanceTypeStandard4:
		return true
	}
	return false
}

// Settings for deployment observability such as logging.
type ApplicationRolloutNewParamsTargetConfigurationObservability struct {
	// Observability logging settings.
	Logs param.Field[ApplicationRolloutNewParamsTargetConfigurationObservabilityLogs] `json:"logs"`
}

func (r ApplicationRolloutNewParamsTargetConfigurationObservability) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Observability logging settings.
type ApplicationRolloutNewParamsTargetConfigurationObservabilityLogs struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ApplicationRolloutNewParamsTargetConfigurationObservabilityLogs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Kind of the rollout process. Defaults to "full_auto".
//
//   - "full_auto": For rolling rollouts, starts progressing steps upon rollout
//     creation. For new_instances rollouts, advances percentage targets
//     automatically after target-version health is observed.
//   - "full_manual": Requires manually progressing each step in the rollout using
//     the UpdateRollout's action parameter.
type ApplicationRolloutNewParamsKind string

const (
	ApplicationRolloutNewParamsKindFullAuto   ApplicationRolloutNewParamsKind = "full_auto"
	ApplicationRolloutNewParamsKindFullManual ApplicationRolloutNewParamsKind = "full_manual"
)

func (r ApplicationRolloutNewParamsKind) IsKnown() bool {
	switch r {
	case ApplicationRolloutNewParamsKindFullAuto, ApplicationRolloutNewParamsKindFullManual:
		return true
	}
	return false
}

// Percentage of rollout to increase in each step when "steps" is absent.
// Applicable values: 5, 10, 20, 25, 50, 100. These create rollouts with 20, 10, 5,
// 4, 2, 1 steps respectively. Only valid for "rolling" strategy.
type ApplicationRolloutNewParamsStepPercentage int64

const (
	ApplicationRolloutNewParamsStepPercentage5   ApplicationRolloutNewParamsStepPercentage = 5
	ApplicationRolloutNewParamsStepPercentage10  ApplicationRolloutNewParamsStepPercentage = 10
	ApplicationRolloutNewParamsStepPercentage20  ApplicationRolloutNewParamsStepPercentage = 20
	ApplicationRolloutNewParamsStepPercentage25  ApplicationRolloutNewParamsStepPercentage = 25
	ApplicationRolloutNewParamsStepPercentage50  ApplicationRolloutNewParamsStepPercentage = 50
	ApplicationRolloutNewParamsStepPercentage100 ApplicationRolloutNewParamsStepPercentage = 100
)

func (r ApplicationRolloutNewParamsStepPercentage) IsKnown() bool {
	switch r {
	case ApplicationRolloutNewParamsStepPercentage5, ApplicationRolloutNewParamsStepPercentage10, ApplicationRolloutNewParamsStepPercentage20, ApplicationRolloutNewParamsStepPercentage25, ApplicationRolloutNewParamsStepPercentage50, ApplicationRolloutNewParamsStepPercentage100:
		return true
	}
	return false
}

// Steps defining the rollout process.
type ApplicationRolloutNewParamsStep struct {
	// Description of the rollout step.
	Description param.Field[string]                                   `json:"description" api:"required"`
	StepSize    param.Field[ApplicationRolloutNewParamsStepsStepSize] `json:"step_size" api:"required"`
}

func (r ApplicationRolloutNewParamsStep) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ApplicationRolloutNewParamsStepsStepSize struct {
	// Percentage of instances affected in this step. Min 10% and Max 100%.
	Percentage param.Field[int64] `json:"percentage" api:"required"`
}

func (r ApplicationRolloutNewParamsStepsStepSize) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ApplicationRolloutNewResponseEnvelope struct {
	Errors   []ApplicationRolloutNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ApplicationRolloutNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Represents the status and metadata of a rollout process for an application. For
	// "rolling" strategy: includes steps and progress with instance counts. For
	// "new_instances" strategy: the response omits steps and progress. Use percentage,
	// version_distribution, and health.summary for status.
	Result ApplicationRolloutNewResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success bool                                      `json:"success" api:"required"`
	JSON    applicationRolloutNewResponseEnvelopeJSON `json:"-"`
}

// applicationRolloutNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ApplicationRolloutNewResponseEnvelope]
type applicationRolloutNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationRolloutNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationRolloutNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ApplicationRolloutNewResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code" api:"required"`
	Message          string                                            `json:"message" api:"required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           ApplicationRolloutNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             applicationRolloutNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// applicationRolloutNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ApplicationRolloutNewResponseEnvelopeErrors]
type applicationRolloutNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ApplicationRolloutNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationRolloutNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ApplicationRolloutNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    applicationRolloutNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// applicationRolloutNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ApplicationRolloutNewResponseEnvelopeErrorsSource]
type applicationRolloutNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationRolloutNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationRolloutNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ApplicationRolloutNewResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code" api:"required"`
	Message          string                                              `json:"message" api:"required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           ApplicationRolloutNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             applicationRolloutNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// applicationRolloutNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ApplicationRolloutNewResponseEnvelopeMessages]
type applicationRolloutNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ApplicationRolloutNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationRolloutNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ApplicationRolloutNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    applicationRolloutNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// applicationRolloutNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [ApplicationRolloutNewResponseEnvelopeMessagesSource]
type applicationRolloutNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApplicationRolloutNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r applicationRolloutNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
