// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
)

// WorkflowInstanceStatusService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowInstanceStatusService] method instead.
type WorkflowInstanceStatusService struct {
	Options []option.RequestOption
}

// NewWorkflowInstanceStatusService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkflowInstanceStatusService(opts ...option.RequestOption) (r *WorkflowInstanceStatusService) {
	r = &WorkflowInstanceStatusService{}
	r.Options = opts
	return
}

// Change status of instance
func (r *WorkflowInstanceStatusService) Edit(ctx context.Context, workflowName string, instanceID string, params WorkflowInstanceStatusEditParams, opts ...option.RequestOption) (res *WorkflowInstanceStatusEditResponse, err error) {
	var env WorkflowInstanceStatusEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
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
	path := fmt.Sprintf("accounts/%s/workflows/%s/instances/%s/status", params.AccountID, workflowName, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkflowInstanceStatusEditResponse struct {
	Status WorkflowInstanceStatusEditResponseStatus `json:"status,required"`
	// In ISO 8601 with no timezone offsets and in UTC.
	Timestamp time.Time                              `json:"timestamp,required" format:"date-time"`
	JSON      workflowInstanceStatusEditResponseJSON `json:"-"`
}

// workflowInstanceStatusEditResponseJSON contains the JSON metadata for the struct
// [WorkflowInstanceStatusEditResponse]
type workflowInstanceStatusEditResponseJSON struct {
	Status      apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowInstanceStatusEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowInstanceStatusEditResponseJSON) RawJSON() string {
	return r.raw
}

type WorkflowInstanceStatusEditResponseStatus string

const (
	WorkflowInstanceStatusEditResponseStatusQueued          WorkflowInstanceStatusEditResponseStatus = "queued"
	WorkflowInstanceStatusEditResponseStatusRunning         WorkflowInstanceStatusEditResponseStatus = "running"
	WorkflowInstanceStatusEditResponseStatusPaused          WorkflowInstanceStatusEditResponseStatus = "paused"
	WorkflowInstanceStatusEditResponseStatusErrored         WorkflowInstanceStatusEditResponseStatus = "errored"
	WorkflowInstanceStatusEditResponseStatusTerminated      WorkflowInstanceStatusEditResponseStatus = "terminated"
	WorkflowInstanceStatusEditResponseStatusComplete        WorkflowInstanceStatusEditResponseStatus = "complete"
	WorkflowInstanceStatusEditResponseStatusWaitingForPause WorkflowInstanceStatusEditResponseStatus = "waitingForPause"
	WorkflowInstanceStatusEditResponseStatusWaiting         WorkflowInstanceStatusEditResponseStatus = "waiting"
	WorkflowInstanceStatusEditResponseStatusUnknown         WorkflowInstanceStatusEditResponseStatus = "unknown"
)

func (r WorkflowInstanceStatusEditResponseStatus) IsKnown() bool {
	switch r {
	case WorkflowInstanceStatusEditResponseStatusQueued, WorkflowInstanceStatusEditResponseStatusRunning, WorkflowInstanceStatusEditResponseStatusPaused, WorkflowInstanceStatusEditResponseStatusErrored, WorkflowInstanceStatusEditResponseStatusTerminated, WorkflowInstanceStatusEditResponseStatusComplete, WorkflowInstanceStatusEditResponseStatusWaitingForPause, WorkflowInstanceStatusEditResponseStatusWaiting, WorkflowInstanceStatusEditResponseStatusUnknown:
		return true
	}
	return false
}

type WorkflowInstanceStatusEditParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Possible actions to apply to instance
	Status param.Field[WorkflowInstanceStatusEditParamsStatus] `json:"status,required"`
}

func (r WorkflowInstanceStatusEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Possible actions to apply to instance
type WorkflowInstanceStatusEditParamsStatus string

const (
	WorkflowInstanceStatusEditParamsStatusResume    WorkflowInstanceStatusEditParamsStatus = "resume"
	WorkflowInstanceStatusEditParamsStatusPause     WorkflowInstanceStatusEditParamsStatus = "pause"
	WorkflowInstanceStatusEditParamsStatusTerminate WorkflowInstanceStatusEditParamsStatus = "terminate"
)

func (r WorkflowInstanceStatusEditParamsStatus) IsKnown() bool {
	switch r {
	case WorkflowInstanceStatusEditParamsStatusResume, WorkflowInstanceStatusEditParamsStatusPause, WorkflowInstanceStatusEditParamsStatusTerminate:
		return true
	}
	return false
}

type WorkflowInstanceStatusEditResponseEnvelope struct {
	Errors     []WorkflowInstanceStatusEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []WorkflowInstanceStatusEditResponseEnvelopeMessages `json:"messages,required"`
	Result     WorkflowInstanceStatusEditResponse                   `json:"result,required"`
	Success    WorkflowInstanceStatusEditResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo WorkflowInstanceStatusEditResponseEnvelopeResultInfo `json:"result_info"`
	JSON       workflowInstanceStatusEditResponseEnvelopeJSON       `json:"-"`
}

// workflowInstanceStatusEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [WorkflowInstanceStatusEditResponseEnvelope]
type workflowInstanceStatusEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowInstanceStatusEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowInstanceStatusEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WorkflowInstanceStatusEditResponseEnvelopeErrors struct {
	Code    float64                                              `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    workflowInstanceStatusEditResponseEnvelopeErrorsJSON `json:"-"`
}

// workflowInstanceStatusEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [WorkflowInstanceStatusEditResponseEnvelopeErrors]
type workflowInstanceStatusEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowInstanceStatusEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowInstanceStatusEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WorkflowInstanceStatusEditResponseEnvelopeMessages struct {
	Code    float64                                                `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    workflowInstanceStatusEditResponseEnvelopeMessagesJSON `json:"-"`
}

// workflowInstanceStatusEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [WorkflowInstanceStatusEditResponseEnvelopeMessages]
type workflowInstanceStatusEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowInstanceStatusEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowInstanceStatusEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type WorkflowInstanceStatusEditResponseEnvelopeSuccess bool

const (
	WorkflowInstanceStatusEditResponseEnvelopeSuccessTrue WorkflowInstanceStatusEditResponseEnvelopeSuccess = true
)

func (r WorkflowInstanceStatusEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WorkflowInstanceStatusEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type WorkflowInstanceStatusEditResponseEnvelopeResultInfo struct {
	Count      float64                                                  `json:"count,required"`
	Page       float64                                                  `json:"page,required"`
	PerPage    float64                                                  `json:"per_page,required"`
	TotalCount float64                                                  `json:"total_count,required"`
	JSON       workflowInstanceStatusEditResponseEnvelopeResultInfoJSON `json:"-"`
}

// workflowInstanceStatusEditResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [WorkflowInstanceStatusEditResponseEnvelopeResultInfo]
type workflowInstanceStatusEditResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowInstanceStatusEditResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowInstanceStatusEditResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
