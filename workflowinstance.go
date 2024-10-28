// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v3/shared"
	"github.com/tidwall/gjson"
)

// WorkflowInstanceService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowInstanceService] method instead.
type WorkflowInstanceService struct {
	Options []option.RequestOption
	Status  *WorkflowInstanceStatusService
}

// NewWorkflowInstanceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkflowInstanceService(opts ...option.RequestOption) (r *WorkflowInstanceService) {
	r = &WorkflowInstanceService{}
	r.Options = opts
	r.Status = NewWorkflowInstanceStatusService(opts...)
	return
}

// Create a new workflow instance
func (r *WorkflowInstanceService) New(ctx context.Context, workflowName string, params WorkflowInstanceNewParams, opts ...option.RequestOption) (res *WorkflowInstanceNewResponse, err error) {
	var env WorkflowInstanceNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if workflowName == "" {
		err = errors.New("missing required workflow_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workflows/%s/instances", params.AccountID, workflowName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List of workflow instances
func (r *WorkflowInstanceService) List(ctx context.Context, workflowName string, params WorkflowInstanceListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[WorkflowInstanceListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if workflowName == "" {
		err = errors.New("missing required workflow_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workflows/%s/instances", params.AccountID, workflowName)
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

// List of workflow instances
func (r *WorkflowInstanceService) ListAutoPaging(ctx context.Context, workflowName string, params WorkflowInstanceListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[WorkflowInstanceListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, workflowName, params, opts...))
}

// Get logs and status from instance
func (r *WorkflowInstanceService) Get(ctx context.Context, workflowName string, instanceID string, query WorkflowInstanceGetParams, opts ...option.RequestOption) (res *WorkflowInstanceGetResponse, err error) {
	var env WorkflowInstanceGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if workflowName == "" {
		err = errors.New("missing required workflow_name parameter")
		return
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workflows/%s/instances/%s", query.AccountID, workflowName, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkflowInstanceNewResponse struct {
	ID         string                            `json:"id,required"`
	Status     WorkflowInstanceNewResponseStatus `json:"status,required"`
	VersionID  string                            `json:"version_id,required" format:"uuid"`
	WorkflowID string                            `json:"workflow_id,required" format:"uuid"`
	JSON       workflowInstanceNewResponseJSON   `json:"-"`
}

// workflowInstanceNewResponseJSON contains the JSON metadata for the struct
// [WorkflowInstanceNewResponse]
type workflowInstanceNewResponseJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	VersionID   apijson.Field
	WorkflowID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowInstanceNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowInstanceNewResponseJSON) RawJSON() string {
	return r.raw
}

type WorkflowInstanceNewResponseStatus string

const (
	WorkflowInstanceNewResponseStatusQueued          WorkflowInstanceNewResponseStatus = "queued"
	WorkflowInstanceNewResponseStatusRunning         WorkflowInstanceNewResponseStatus = "running"
	WorkflowInstanceNewResponseStatusPaused          WorkflowInstanceNewResponseStatus = "paused"
	WorkflowInstanceNewResponseStatusErrored         WorkflowInstanceNewResponseStatus = "errored"
	WorkflowInstanceNewResponseStatusTerminated      WorkflowInstanceNewResponseStatus = "terminated"
	WorkflowInstanceNewResponseStatusComplete        WorkflowInstanceNewResponseStatus = "complete"
	WorkflowInstanceNewResponseStatusWaitingForPause WorkflowInstanceNewResponseStatus = "waitingForPause"
	WorkflowInstanceNewResponseStatusWaiting         WorkflowInstanceNewResponseStatus = "waiting"
	WorkflowInstanceNewResponseStatusUnknown         WorkflowInstanceNewResponseStatus = "unknown"
)

func (r WorkflowInstanceNewResponseStatus) IsKnown() bool {
	switch r {
	case WorkflowInstanceNewResponseStatusQueued, WorkflowInstanceNewResponseStatusRunning, WorkflowInstanceNewResponseStatusPaused, WorkflowInstanceNewResponseStatusErrored, WorkflowInstanceNewResponseStatusTerminated, WorkflowInstanceNewResponseStatusComplete, WorkflowInstanceNewResponseStatusWaitingForPause, WorkflowInstanceNewResponseStatusWaiting, WorkflowInstanceNewResponseStatusUnknown:
		return true
	}
	return false
}

type WorkflowInstanceListResponse struct {
	ID         string                             `json:"id,required"`
	CreatedOn  time.Time                          `json:"created_on,required" format:"date-time"`
	ModifiedOn time.Time                          `json:"modified_on,required" format:"date-time"`
	Status     WorkflowInstanceListResponseStatus `json:"status,required"`
	VersionID  string                             `json:"version_id,required" format:"uuid"`
	WorkflowID string                             `json:"workflow_id,required" format:"uuid"`
	JSON       workflowInstanceListResponseJSON   `json:"-"`
}

// workflowInstanceListResponseJSON contains the JSON metadata for the struct
// [WorkflowInstanceListResponse]
type workflowInstanceListResponseJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	Status      apijson.Field
	VersionID   apijson.Field
	WorkflowID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowInstanceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowInstanceListResponseJSON) RawJSON() string {
	return r.raw
}

type WorkflowInstanceListResponseStatus string

const (
	WorkflowInstanceListResponseStatusQueued          WorkflowInstanceListResponseStatus = "queued"
	WorkflowInstanceListResponseStatusRunning         WorkflowInstanceListResponseStatus = "running"
	WorkflowInstanceListResponseStatusPaused          WorkflowInstanceListResponseStatus = "paused"
	WorkflowInstanceListResponseStatusErrored         WorkflowInstanceListResponseStatus = "errored"
	WorkflowInstanceListResponseStatusTerminated      WorkflowInstanceListResponseStatus = "terminated"
	WorkflowInstanceListResponseStatusComplete        WorkflowInstanceListResponseStatus = "complete"
	WorkflowInstanceListResponseStatusWaitingForPause WorkflowInstanceListResponseStatus = "waitingForPause"
	WorkflowInstanceListResponseStatusWaiting         WorkflowInstanceListResponseStatus = "waiting"
	WorkflowInstanceListResponseStatusUnknown         WorkflowInstanceListResponseStatus = "unknown"
)

func (r WorkflowInstanceListResponseStatus) IsKnown() bool {
	switch r {
	case WorkflowInstanceListResponseStatusQueued, WorkflowInstanceListResponseStatusRunning, WorkflowInstanceListResponseStatusPaused, WorkflowInstanceListResponseStatusErrored, WorkflowInstanceListResponseStatusTerminated, WorkflowInstanceListResponseStatusComplete, WorkflowInstanceListResponseStatusWaitingForPause, WorkflowInstanceListResponseStatusWaiting, WorkflowInstanceListResponseStatusUnknown:
		return true
	}
	return false
}

type WorkflowInstanceGetResponse struct {
	End       time.Time                              `json:"end,required,nullable" format:"date-time"`
	Error     WorkflowInstanceGetResponseError       `json:"error,required,nullable"`
	Output    WorkflowInstanceGetResponseOutputUnion `json:"output,required"`
	Params    interface{}                            `json:"params,required"`
	Queued    time.Time                              `json:"queued,required" format:"date-time"`
	Start     time.Time                              `json:"start,required,nullable" format:"date-time"`
	Status    WorkflowInstanceGetResponseStatus      `json:"status,required"`
	Steps     []WorkflowInstanceGetResponseStep      `json:"steps,required"`
	Success   bool                                   `json:"success,required,nullable"`
	Trigger   WorkflowInstanceGetResponseTrigger     `json:"trigger,required"`
	VersionID string                                 `json:"versionId,required" format:"uuid"`
	JSON      workflowInstanceGetResponseJSON        `json:"-"`
}

// workflowInstanceGetResponseJSON contains the JSON metadata for the struct
// [WorkflowInstanceGetResponse]
type workflowInstanceGetResponseJSON struct {
	End         apijson.Field
	Error       apijson.Field
	Output      apijson.Field
	Params      apijson.Field
	Queued      apijson.Field
	Start       apijson.Field
	Status      apijson.Field
	Steps       apijson.Field
	Success     apijson.Field
	Trigger     apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowInstanceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowInstanceGetResponseJSON) RawJSON() string {
	return r.raw
}

type WorkflowInstanceGetResponseError struct {
	Message string                               `json:"message,required"`
	Name    string                               `json:"name,required"`
	JSON    workflowInstanceGetResponseErrorJSON `json:"-"`
}

// workflowInstanceGetResponseErrorJSON contains the JSON metadata for the struct
// [WorkflowInstanceGetResponseError]
type workflowInstanceGetResponseErrorJSON struct {
	Message     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowInstanceGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowInstanceGetResponseErrorJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type WorkflowInstanceGetResponseOutputUnion interface {
	ImplementsWorkflowInstanceGetResponseOutputUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WorkflowInstanceGetResponseOutputUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type WorkflowInstanceGetResponseStatus string

const (
	WorkflowInstanceGetResponseStatusQueued          WorkflowInstanceGetResponseStatus = "queued"
	WorkflowInstanceGetResponseStatusRunning         WorkflowInstanceGetResponseStatus = "running"
	WorkflowInstanceGetResponseStatusPaused          WorkflowInstanceGetResponseStatus = "paused"
	WorkflowInstanceGetResponseStatusErrored         WorkflowInstanceGetResponseStatus = "errored"
	WorkflowInstanceGetResponseStatusTerminated      WorkflowInstanceGetResponseStatus = "terminated"
	WorkflowInstanceGetResponseStatusComplete        WorkflowInstanceGetResponseStatus = "complete"
	WorkflowInstanceGetResponseStatusWaitingForPause WorkflowInstanceGetResponseStatus = "waitingForPause"
	WorkflowInstanceGetResponseStatusWaiting         WorkflowInstanceGetResponseStatus = "waiting"
	WorkflowInstanceGetResponseStatusUnknown         WorkflowInstanceGetResponseStatus = "unknown"
)

func (r WorkflowInstanceGetResponseStatus) IsKnown() bool {
	switch r {
	case WorkflowInstanceGetResponseStatusQueued, WorkflowInstanceGetResponseStatusRunning, WorkflowInstanceGetResponseStatusPaused, WorkflowInstanceGetResponseStatusErrored, WorkflowInstanceGetResponseStatusTerminated, WorkflowInstanceGetResponseStatusComplete, WorkflowInstanceGetResponseStatusWaitingForPause, WorkflowInstanceGetResponseStatusWaiting, WorkflowInstanceGetResponseStatusUnknown:
		return true
	}
	return false
}

type WorkflowInstanceGetResponseStep struct {
	// This field can have the runtime type of
	// [[]WorkflowInstanceGetResponseStepsObjectAttempt].
	Attempts interface{} `json:"attempts,required"`
	// This field can have the runtime type of
	// [WorkflowInstanceGetResponseStepsObjectConfig].
	Config interface{} `json:"config,required"`
	End    time.Time   `json:"end,nullable" format:"date-time"`
	Name   string      `json:"name"`
	// This field can have the runtime type of [interface{}].
	Output  interface{}                          `json:"output,required"`
	Start   time.Time                            `json:"start" format:"date-time"`
	Success bool                                 `json:"success,nullable"`
	Type    WorkflowInstanceGetResponseStepsType `json:"type,required"`
	// This field can have the runtime type of
	// [WorkflowInstanceGetResponseStepsObjectError].
	Error    interface{} `json:"error,required"`
	Finished bool        `json:"finished"`
	// This field can have the runtime type of
	// [WorkflowInstanceGetResponseStepsObjectTrigger].
	Trigger interface{}                         `json:"trigger,required"`
	JSON    workflowInstanceGetResponseStepJSON `json:"-"`
	union   WorkflowInstanceGetResponseStepsUnion
}

// workflowInstanceGetResponseStepJSON contains the JSON metadata for the struct
// [WorkflowInstanceGetResponseStep]
type workflowInstanceGetResponseStepJSON struct {
	Attempts    apijson.Field
	Config      apijson.Field
	End         apijson.Field
	Name        apijson.Field
	Output      apijson.Field
	Start       apijson.Field
	Success     apijson.Field
	Type        apijson.Field
	Error       apijson.Field
	Finished    apijson.Field
	Trigger     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r workflowInstanceGetResponseStepJSON) RawJSON() string {
	return r.raw
}

func (r *WorkflowInstanceGetResponseStep) UnmarshalJSON(data []byte) (err error) {
	*r = WorkflowInstanceGetResponseStep{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [WorkflowInstanceGetResponseStepsUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [WorkflowInstanceGetResponseStepsObject],
// [WorkflowInstanceGetResponseStepsObject],
// [WorkflowInstanceGetResponseStepsObject].
func (r WorkflowInstanceGetResponseStep) AsUnion() WorkflowInstanceGetResponseStepsUnion {
	return r.union
}

// Union satisfied by [WorkflowInstanceGetResponseStepsObject],
// [WorkflowInstanceGetResponseStepsObject] or
// [WorkflowInstanceGetResponseStepsObject].
type WorkflowInstanceGetResponseStepsUnion interface {
	implementsWorkflowInstanceGetResponseStep()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WorkflowInstanceGetResponseStepsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WorkflowInstanceGetResponseStepsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WorkflowInstanceGetResponseStepsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WorkflowInstanceGetResponseStepsObject{}),
		},
	)
}

type WorkflowInstanceGetResponseStepsObject struct {
	Attempts []WorkflowInstanceGetResponseStepsObjectAttempt `json:"attempts,required"`
	Config   WorkflowInstanceGetResponseStepsObjectConfig    `json:"config,required"`
	End      time.Time                                       `json:"end,required,nullable" format:"date-time"`
	Name     string                                          `json:"name,required"`
	Output   interface{}                                     `json:"output,required"`
	Start    time.Time                                       `json:"start,required" format:"date-time"`
	Success  bool                                            `json:"success,required,nullable"`
	Type     WorkflowInstanceGetResponseStepsObjectType      `json:"type,required"`
	JSON     workflowInstanceGetResponseStepsObjectJSON      `json:"-"`
}

// workflowInstanceGetResponseStepsObjectJSON contains the JSON metadata for the
// struct [WorkflowInstanceGetResponseStepsObject]
type workflowInstanceGetResponseStepsObjectJSON struct {
	Attempts    apijson.Field
	Config      apijson.Field
	End         apijson.Field
	Name        apijson.Field
	Output      apijson.Field
	Start       apijson.Field
	Success     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowInstanceGetResponseStepsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowInstanceGetResponseStepsObjectJSON) RawJSON() string {
	return r.raw
}

func (r WorkflowInstanceGetResponseStepsObject) implementsWorkflowInstanceGetResponseStep() {}

type WorkflowInstanceGetResponseStepsObjectAttempt struct {
	End     time.Time                                           `json:"end,required,nullable" format:"date-time"`
	Error   WorkflowInstanceGetResponseStepsObjectAttemptsError `json:"error,required,nullable"`
	Start   time.Time                                           `json:"start,required" format:"date-time"`
	Success bool                                                `json:"success,required,nullable"`
	JSON    workflowInstanceGetResponseStepsObjectAttemptJSON   `json:"-"`
}

// workflowInstanceGetResponseStepsObjectAttemptJSON contains the JSON metadata for
// the struct [WorkflowInstanceGetResponseStepsObjectAttempt]
type workflowInstanceGetResponseStepsObjectAttemptJSON struct {
	End         apijson.Field
	Error       apijson.Field
	Start       apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowInstanceGetResponseStepsObjectAttempt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowInstanceGetResponseStepsObjectAttemptJSON) RawJSON() string {
	return r.raw
}

type WorkflowInstanceGetResponseStepsObjectAttemptsError struct {
	Message string                                                  `json:"message,required"`
	Name    string                                                  `json:"name,required"`
	JSON    workflowInstanceGetResponseStepsObjectAttemptsErrorJSON `json:"-"`
}

// workflowInstanceGetResponseStepsObjectAttemptsErrorJSON contains the JSON
// metadata for the struct [WorkflowInstanceGetResponseStepsObjectAttemptsError]
type workflowInstanceGetResponseStepsObjectAttemptsErrorJSON struct {
	Message     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowInstanceGetResponseStepsObjectAttemptsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowInstanceGetResponseStepsObjectAttemptsErrorJSON) RawJSON() string {
	return r.raw
}

type WorkflowInstanceGetResponseStepsObjectConfig struct {
	Retries WorkflowInstanceGetResponseStepsObjectConfigRetries      `json:"retries,required"`
	Timeout WorkflowInstanceGetResponseStepsObjectConfigTimeoutUnion `json:"timeout,required"`
	JSON    workflowInstanceGetResponseStepsObjectConfigJSON         `json:"-"`
}

// workflowInstanceGetResponseStepsObjectConfigJSON contains the JSON metadata for
// the struct [WorkflowInstanceGetResponseStepsObjectConfig]
type workflowInstanceGetResponseStepsObjectConfigJSON struct {
	Retries     apijson.Field
	Timeout     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowInstanceGetResponseStepsObjectConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowInstanceGetResponseStepsObjectConfigJSON) RawJSON() string {
	return r.raw
}

type WorkflowInstanceGetResponseStepsObjectConfigRetries struct {
	Delay   WorkflowInstanceGetResponseStepsObjectConfigRetriesDelayUnion `json:"delay,required"`
	Limit   float64                                                       `json:"limit,required"`
	Backoff WorkflowInstanceGetResponseStepsObjectConfigRetriesBackoff    `json:"backoff"`
	JSON    workflowInstanceGetResponseStepsObjectConfigRetriesJSON       `json:"-"`
}

// workflowInstanceGetResponseStepsObjectConfigRetriesJSON contains the JSON
// metadata for the struct [WorkflowInstanceGetResponseStepsObjectConfigRetries]
type workflowInstanceGetResponseStepsObjectConfigRetriesJSON struct {
	Delay       apijson.Field
	Limit       apijson.Field
	Backoff     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowInstanceGetResponseStepsObjectConfigRetries) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowInstanceGetResponseStepsObjectConfigRetriesJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type WorkflowInstanceGetResponseStepsObjectConfigRetriesDelayUnion interface {
	ImplementsWorkflowInstanceGetResponseStepsObjectConfigRetriesDelayUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WorkflowInstanceGetResponseStepsObjectConfigRetriesDelayUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type WorkflowInstanceGetResponseStepsObjectConfigRetriesBackoff string

const (
	WorkflowInstanceGetResponseStepsObjectConfigRetriesBackoffConstant    WorkflowInstanceGetResponseStepsObjectConfigRetriesBackoff = "constant"
	WorkflowInstanceGetResponseStepsObjectConfigRetriesBackoffLinear      WorkflowInstanceGetResponseStepsObjectConfigRetriesBackoff = "linear"
	WorkflowInstanceGetResponseStepsObjectConfigRetriesBackoffExponential WorkflowInstanceGetResponseStepsObjectConfigRetriesBackoff = "exponential"
)

func (r WorkflowInstanceGetResponseStepsObjectConfigRetriesBackoff) IsKnown() bool {
	switch r {
	case WorkflowInstanceGetResponseStepsObjectConfigRetriesBackoffConstant, WorkflowInstanceGetResponseStepsObjectConfigRetriesBackoffLinear, WorkflowInstanceGetResponseStepsObjectConfigRetriesBackoffExponential:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type WorkflowInstanceGetResponseStepsObjectConfigTimeoutUnion interface {
	ImplementsWorkflowInstanceGetResponseStepsObjectConfigTimeoutUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WorkflowInstanceGetResponseStepsObjectConfigTimeoutUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type WorkflowInstanceGetResponseStepsObjectType string

const (
	WorkflowInstanceGetResponseStepsObjectTypeStep WorkflowInstanceGetResponseStepsObjectType = "step"
)

func (r WorkflowInstanceGetResponseStepsObjectType) IsKnown() bool {
	switch r {
	case WorkflowInstanceGetResponseStepsObjectTypeStep:
		return true
	}
	return false
}

type WorkflowInstanceGetResponseStepsType string

const (
	WorkflowInstanceGetResponseStepsTypeStep        WorkflowInstanceGetResponseStepsType = "step"
	WorkflowInstanceGetResponseStepsTypeSleep       WorkflowInstanceGetResponseStepsType = "sleep"
	WorkflowInstanceGetResponseStepsTypeTermination WorkflowInstanceGetResponseStepsType = "termination"
)

func (r WorkflowInstanceGetResponseStepsType) IsKnown() bool {
	switch r {
	case WorkflowInstanceGetResponseStepsTypeStep, WorkflowInstanceGetResponseStepsTypeSleep, WorkflowInstanceGetResponseStepsTypeTermination:
		return true
	}
	return false
}

type WorkflowInstanceGetResponseTrigger struct {
	Source WorkflowInstanceGetResponseTriggerSource `json:"source,required"`
	JSON   workflowInstanceGetResponseTriggerJSON   `json:"-"`
}

// workflowInstanceGetResponseTriggerJSON contains the JSON metadata for the struct
// [WorkflowInstanceGetResponseTrigger]
type workflowInstanceGetResponseTriggerJSON struct {
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowInstanceGetResponseTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowInstanceGetResponseTriggerJSON) RawJSON() string {
	return r.raw
}

type WorkflowInstanceGetResponseTriggerSource string

const (
	WorkflowInstanceGetResponseTriggerSourceUnknown WorkflowInstanceGetResponseTriggerSource = "unknown"
	WorkflowInstanceGetResponseTriggerSourceAPI     WorkflowInstanceGetResponseTriggerSource = "api"
	WorkflowInstanceGetResponseTriggerSourceBinding WorkflowInstanceGetResponseTriggerSource = "binding"
	WorkflowInstanceGetResponseTriggerSourceEvent   WorkflowInstanceGetResponseTriggerSource = "event"
	WorkflowInstanceGetResponseTriggerSourceCron    WorkflowInstanceGetResponseTriggerSource = "cron"
)

func (r WorkflowInstanceGetResponseTriggerSource) IsKnown() bool {
	switch r {
	case WorkflowInstanceGetResponseTriggerSourceUnknown, WorkflowInstanceGetResponseTriggerSourceAPI, WorkflowInstanceGetResponseTriggerSourceBinding, WorkflowInstanceGetResponseTriggerSourceEvent, WorkflowInstanceGetResponseTriggerSourceCron:
		return true
	}
	return false
}

type WorkflowInstanceNewParams struct {
	AccountID  param.Field[string]      `path:"account_id,required"`
	InstanceID param.Field[string]      `json:"instance_id"`
	Params     param.Field[interface{}] `json:"params"`
}

func (r WorkflowInstanceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkflowInstanceNewResponseEnvelope struct {
	Errors     []WorkflowInstanceNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []WorkflowInstanceNewResponseEnvelopeMessages `json:"messages,required"`
	Result     WorkflowInstanceNewResponse                   `json:"result,required"`
	Success    WorkflowInstanceNewResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo WorkflowInstanceNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       workflowInstanceNewResponseEnvelopeJSON       `json:"-"`
}

// workflowInstanceNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [WorkflowInstanceNewResponseEnvelope]
type workflowInstanceNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowInstanceNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowInstanceNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WorkflowInstanceNewResponseEnvelopeErrors struct {
	Code    float64                                       `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    workflowInstanceNewResponseEnvelopeErrorsJSON `json:"-"`
}

// workflowInstanceNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WorkflowInstanceNewResponseEnvelopeErrors]
type workflowInstanceNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowInstanceNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowInstanceNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WorkflowInstanceNewResponseEnvelopeMessages struct {
	Code    float64                                         `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    workflowInstanceNewResponseEnvelopeMessagesJSON `json:"-"`
}

// workflowInstanceNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [WorkflowInstanceNewResponseEnvelopeMessages]
type workflowInstanceNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowInstanceNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowInstanceNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type WorkflowInstanceNewResponseEnvelopeSuccess bool

const (
	WorkflowInstanceNewResponseEnvelopeSuccessTrue WorkflowInstanceNewResponseEnvelopeSuccess = true
)

func (r WorkflowInstanceNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WorkflowInstanceNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type WorkflowInstanceNewResponseEnvelopeResultInfo struct {
	Count      float64                                           `json:"count,required"`
	Page       float64                                           `json:"page,required"`
	PerPage    float64                                           `json:"per_page,required"`
	TotalCount float64                                           `json:"total_count,required"`
	JSON       workflowInstanceNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// workflowInstanceNewResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [WorkflowInstanceNewResponseEnvelopeResultInfo]
type workflowInstanceNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowInstanceNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowInstanceNewResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type WorkflowInstanceListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// In ISO 8601 with no timezone offsets and in UTC.
	DateEnd param.Field[time.Time] `query:"date_end" format:"date-time"`
	// In ISO 8601 with no timezone offsets and in UTC.
	DateStart param.Field[time.Time]                        `query:"date_start" format:"date-time"`
	Page      param.Field[float64]                          `query:"page"`
	PerPage   param.Field[float64]                          `query:"per_page"`
	Status    param.Field[WorkflowInstanceListParamsStatus] `query:"status"`
}

// URLQuery serializes [WorkflowInstanceListParams]'s query parameters as
// `url.Values`.
func (r WorkflowInstanceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type WorkflowInstanceListParamsStatus string

const (
	WorkflowInstanceListParamsStatusQueued          WorkflowInstanceListParamsStatus = "queued"
	WorkflowInstanceListParamsStatusRunning         WorkflowInstanceListParamsStatus = "running"
	WorkflowInstanceListParamsStatusPaused          WorkflowInstanceListParamsStatus = "paused"
	WorkflowInstanceListParamsStatusErrored         WorkflowInstanceListParamsStatus = "errored"
	WorkflowInstanceListParamsStatusTerminated      WorkflowInstanceListParamsStatus = "terminated"
	WorkflowInstanceListParamsStatusComplete        WorkflowInstanceListParamsStatus = "complete"
	WorkflowInstanceListParamsStatusWaitingForPause WorkflowInstanceListParamsStatus = "waitingForPause"
	WorkflowInstanceListParamsStatusWaiting         WorkflowInstanceListParamsStatus = "waiting"
	WorkflowInstanceListParamsStatusUnknown         WorkflowInstanceListParamsStatus = "unknown"
)

func (r WorkflowInstanceListParamsStatus) IsKnown() bool {
	switch r {
	case WorkflowInstanceListParamsStatusQueued, WorkflowInstanceListParamsStatusRunning, WorkflowInstanceListParamsStatusPaused, WorkflowInstanceListParamsStatusErrored, WorkflowInstanceListParamsStatusTerminated, WorkflowInstanceListParamsStatusComplete, WorkflowInstanceListParamsStatusWaitingForPause, WorkflowInstanceListParamsStatusWaiting, WorkflowInstanceListParamsStatusUnknown:
		return true
	}
	return false
}

type WorkflowInstanceGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type WorkflowInstanceGetResponseEnvelope struct {
	Errors     []WorkflowInstanceGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []WorkflowInstanceGetResponseEnvelopeMessages `json:"messages,required"`
	Result     WorkflowInstanceGetResponse                   `json:"result,required"`
	Success    WorkflowInstanceGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo WorkflowInstanceGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       workflowInstanceGetResponseEnvelopeJSON       `json:"-"`
}

// workflowInstanceGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [WorkflowInstanceGetResponseEnvelope]
type workflowInstanceGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowInstanceGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowInstanceGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WorkflowInstanceGetResponseEnvelopeErrors struct {
	Code    float64                                       `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    workflowInstanceGetResponseEnvelopeErrorsJSON `json:"-"`
}

// workflowInstanceGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WorkflowInstanceGetResponseEnvelopeErrors]
type workflowInstanceGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowInstanceGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowInstanceGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WorkflowInstanceGetResponseEnvelopeMessages struct {
	Code    float64                                         `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    workflowInstanceGetResponseEnvelopeMessagesJSON `json:"-"`
}

// workflowInstanceGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [WorkflowInstanceGetResponseEnvelopeMessages]
type workflowInstanceGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowInstanceGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowInstanceGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type WorkflowInstanceGetResponseEnvelopeSuccess bool

const (
	WorkflowInstanceGetResponseEnvelopeSuccessTrue WorkflowInstanceGetResponseEnvelopeSuccess = true
)

func (r WorkflowInstanceGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WorkflowInstanceGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type WorkflowInstanceGetResponseEnvelopeResultInfo struct {
	Count      float64                                           `json:"count,required"`
	Page       float64                                           `json:"page,required"`
	PerPage    float64                                           `json:"per_page,required"`
	TotalCount float64                                           `json:"total_count,required"`
	JSON       workflowInstanceGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// workflowInstanceGetResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [WorkflowInstanceGetResponseEnvelopeResultInfo]
type workflowInstanceGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowInstanceGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowInstanceGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
