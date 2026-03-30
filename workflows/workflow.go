// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflows

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// WorkflowService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowService] method instead.
type WorkflowService struct {
	Options   []option.RequestOption
	Instances *InstanceService
	Versions  *VersionService
}

// NewWorkflowService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkflowService(opts ...option.RequestOption) (r *WorkflowService) {
	r = &WorkflowService{}
	r.Options = opts
	r.Instances = NewInstanceService(opts...)
	r.Versions = NewVersionService(opts...)
	return
}

// Creates a new workflow or updates an existing workflow definition.
func (r *WorkflowService) Update(ctx context.Context, workflowName string, params WorkflowUpdateParams, opts ...option.RequestOption) (res *WorkflowUpdateResponse, err error) {
	var env WorkflowUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if workflowName == "" {
		err = errors.New("missing required workflow_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workflows/%s", params.AccountID, workflowName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists all workflows configured for the account.
func (r *WorkflowService) List(ctx context.Context, params WorkflowListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[WorkflowListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workflows", params.AccountID)
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

// Lists all workflows configured for the account.
func (r *WorkflowService) ListAutoPaging(ctx context.Context, params WorkflowListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[WorkflowListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes a Workflow. This only deletes the Workflow and does not delete or modify
// any Worker associated to this Workflow or bounded to it.
func (r *WorkflowService) Delete(ctx context.Context, workflowName string, body WorkflowDeleteParams, opts ...option.RequestOption) (res *WorkflowDeleteResponse, err error) {
	var env WorkflowDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if workflowName == "" {
		err = errors.New("missing required workflow_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workflows/%s", body.AccountID, workflowName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieves configuration and metadata for a specific workflow.
func (r *WorkflowService) Get(ctx context.Context, workflowName string, query WorkflowGetParams, opts ...option.RequestOption) (res *WorkflowGetResponse, err error) {
	var env WorkflowGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if workflowName == "" {
		err = errors.New("missing required workflow_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workflows/%s", query.AccountID, workflowName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type WorkflowUpdateResponse struct {
	ID                string                     `json:"id" api:"required" format:"uuid"`
	ClassName         string                     `json:"class_name" api:"required"`
	CreatedOn         time.Time                  `json:"created_on" api:"required" format:"date-time"`
	IsDeleted         float64                    `json:"is_deleted" api:"required"`
	ModifiedOn        time.Time                  `json:"modified_on" api:"required" format:"date-time"`
	Name              string                     `json:"name" api:"required"`
	ScriptName        string                     `json:"script_name" api:"required"`
	TerminatorRunning float64                    `json:"terminator_running" api:"required"`
	TriggeredOn       time.Time                  `json:"triggered_on" api:"required,nullable" format:"date-time"`
	VersionID         string                     `json:"version_id" api:"required" format:"uuid"`
	JSON              workflowUpdateResponseJSON `json:"-"`
}

// workflowUpdateResponseJSON contains the JSON metadata for the struct
// [WorkflowUpdateResponse]
type workflowUpdateResponseJSON struct {
	ID                apijson.Field
	ClassName         apijson.Field
	CreatedOn         apijson.Field
	IsDeleted         apijson.Field
	ModifiedOn        apijson.Field
	Name              apijson.Field
	ScriptName        apijson.Field
	TerminatorRunning apijson.Field
	TriggeredOn       apijson.Field
	VersionID         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *WorkflowUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type WorkflowListResponse struct {
	ID          string                        `json:"id" api:"required" format:"uuid"`
	ClassName   string                        `json:"class_name" api:"required"`
	CreatedOn   time.Time                     `json:"created_on" api:"required" format:"date-time"`
	Instances   WorkflowListResponseInstances `json:"instances" api:"required"`
	ModifiedOn  time.Time                     `json:"modified_on" api:"required" format:"date-time"`
	Name        string                        `json:"name" api:"required"`
	ScriptName  string                        `json:"script_name" api:"required"`
	TriggeredOn time.Time                     `json:"triggered_on" api:"required,nullable" format:"date-time"`
	JSON        workflowListResponseJSON      `json:"-"`
}

// workflowListResponseJSON contains the JSON metadata for the struct
// [WorkflowListResponse]
type workflowListResponseJSON struct {
	ID          apijson.Field
	ClassName   apijson.Field
	CreatedOn   apijson.Field
	Instances   apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	ScriptName  apijson.Field
	TriggeredOn apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowListResponseJSON) RawJSON() string {
	return r.raw
}

type WorkflowListResponseInstances struct {
	Complete        float64                           `json:"complete"`
	Errored         float64                           `json:"errored"`
	Paused          float64                           `json:"paused"`
	Queued          float64                           `json:"queued"`
	Running         float64                           `json:"running"`
	Terminated      float64                           `json:"terminated"`
	Waiting         float64                           `json:"waiting"`
	WaitingForPause float64                           `json:"waitingForPause"`
	JSON            workflowListResponseInstancesJSON `json:"-"`
}

// workflowListResponseInstancesJSON contains the JSON metadata for the struct
// [WorkflowListResponseInstances]
type workflowListResponseInstancesJSON struct {
	Complete        apijson.Field
	Errored         apijson.Field
	Paused          apijson.Field
	Queued          apijson.Field
	Running         apijson.Field
	Terminated      apijson.Field
	Waiting         apijson.Field
	WaitingForPause apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WorkflowListResponseInstances) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowListResponseInstancesJSON) RawJSON() string {
	return r.raw
}

type WorkflowDeleteResponse struct {
	Status  WorkflowDeleteResponseStatus `json:"status" api:"required"`
	Success bool                         `json:"success" api:"required,nullable"`
	JSON    workflowDeleteResponseJSON   `json:"-"`
}

// workflowDeleteResponseJSON contains the JSON metadata for the struct
// [WorkflowDeleteResponse]
type workflowDeleteResponseJSON struct {
	Status      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type WorkflowDeleteResponseStatus string

const (
	WorkflowDeleteResponseStatusOk WorkflowDeleteResponseStatus = "ok"
)

func (r WorkflowDeleteResponseStatus) IsKnown() bool {
	switch r {
	case WorkflowDeleteResponseStatusOk:
		return true
	}
	return false
}

type WorkflowGetResponse struct {
	ID          string                       `json:"id" api:"required" format:"uuid"`
	ClassName   string                       `json:"class_name" api:"required"`
	CreatedOn   time.Time                    `json:"created_on" api:"required" format:"date-time"`
	Instances   WorkflowGetResponseInstances `json:"instances" api:"required"`
	ModifiedOn  time.Time                    `json:"modified_on" api:"required" format:"date-time"`
	Name        string                       `json:"name" api:"required"`
	ScriptName  string                       `json:"script_name" api:"required"`
	TriggeredOn time.Time                    `json:"triggered_on" api:"required,nullable" format:"date-time"`
	JSON        workflowGetResponseJSON      `json:"-"`
}

// workflowGetResponseJSON contains the JSON metadata for the struct
// [WorkflowGetResponse]
type workflowGetResponseJSON struct {
	ID          apijson.Field
	ClassName   apijson.Field
	CreatedOn   apijson.Field
	Instances   apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	ScriptName  apijson.Field
	TriggeredOn apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowGetResponseJSON) RawJSON() string {
	return r.raw
}

type WorkflowGetResponseInstances struct {
	Complete        float64                          `json:"complete"`
	Errored         float64                          `json:"errored"`
	Paused          float64                          `json:"paused"`
	Queued          float64                          `json:"queued"`
	Running         float64                          `json:"running"`
	Terminated      float64                          `json:"terminated"`
	Waiting         float64                          `json:"waiting"`
	WaitingForPause float64                          `json:"waitingForPause"`
	JSON            workflowGetResponseInstancesJSON `json:"-"`
}

// workflowGetResponseInstancesJSON contains the JSON metadata for the struct
// [WorkflowGetResponseInstances]
type workflowGetResponseInstancesJSON struct {
	Complete        apijson.Field
	Errored         apijson.Field
	Paused          apijson.Field
	Queued          apijson.Field
	Running         apijson.Field
	Terminated      apijson.Field
	Waiting         apijson.Field
	WaitingForPause apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WorkflowGetResponseInstances) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowGetResponseInstancesJSON) RawJSON() string {
	return r.raw
}

type WorkflowUpdateParams struct {
	AccountID  param.Field[string]                     `path:"account_id" api:"required"`
	ClassName  param.Field[string]                     `json:"class_name" api:"required"`
	ScriptName param.Field[string]                     `json:"script_name" api:"required"`
	Limits     param.Field[WorkflowUpdateParamsLimits] `json:"limits"`
}

func (r WorkflowUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkflowUpdateParamsLimits struct {
	Steps param.Field[int64] `json:"steps"`
}

func (r WorkflowUpdateParamsLimits) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkflowUpdateResponseEnvelope struct {
	Errors     []WorkflowUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages   []WorkflowUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	Result     WorkflowUpdateResponse                   `json:"result" api:"required"`
	Success    WorkflowUpdateResponseEnvelopeSuccess    `json:"success" api:"required"`
	ResultInfo WorkflowUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       workflowUpdateResponseEnvelopeJSON       `json:"-"`
}

// workflowUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [WorkflowUpdateResponseEnvelope]
type workflowUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WorkflowUpdateResponseEnvelopeErrors struct {
	Code    float64                                  `json:"code" api:"required"`
	Message string                                   `json:"message" api:"required"`
	JSON    workflowUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// workflowUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WorkflowUpdateResponseEnvelopeErrors]
type workflowUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WorkflowUpdateResponseEnvelopeMessages struct {
	Code    float64                                    `json:"code" api:"required"`
	Message string                                     `json:"message" api:"required"`
	JSON    workflowUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// workflowUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WorkflowUpdateResponseEnvelopeMessages]
type workflowUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type WorkflowUpdateResponseEnvelopeSuccess bool

const (
	WorkflowUpdateResponseEnvelopeSuccessTrue WorkflowUpdateResponseEnvelopeSuccess = true
)

func (r WorkflowUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WorkflowUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type WorkflowUpdateResponseEnvelopeResultInfo struct {
	Count      float64                                      `json:"count" api:"required"`
	PerPage    float64                                      `json:"per_page" api:"required"`
	TotalCount float64                                      `json:"total_count" api:"required"`
	Cursor     string                                       `json:"cursor"`
	Page       float64                                      `json:"page"`
	JSON       workflowUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// workflowUpdateResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [WorkflowUpdateResponseEnvelopeResultInfo]
type workflowUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	Cursor      apijson.Field
	Page        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type WorkflowListParams struct {
	AccountID param.Field[string]  `path:"account_id" api:"required"`
	Page      param.Field[float64] `query:"page"`
	PerPage   param.Field[float64] `query:"per_page"`
	// Allows filtering workflows` name.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [WorkflowListParams]'s query parameters as `url.Values`.
func (r WorkflowListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type WorkflowDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type WorkflowDeleteResponseEnvelope struct {
	Errors     []WorkflowDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages   []WorkflowDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	Result     WorkflowDeleteResponse                   `json:"result" api:"required"`
	Success    WorkflowDeleteResponseEnvelopeSuccess    `json:"success" api:"required"`
	ResultInfo WorkflowDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       workflowDeleteResponseEnvelopeJSON       `json:"-"`
}

// workflowDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [WorkflowDeleteResponseEnvelope]
type workflowDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WorkflowDeleteResponseEnvelopeErrors struct {
	Code    float64                                  `json:"code" api:"required"`
	Message string                                   `json:"message" api:"required"`
	JSON    workflowDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// workflowDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WorkflowDeleteResponseEnvelopeErrors]
type workflowDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WorkflowDeleteResponseEnvelopeMessages struct {
	Code    float64                                    `json:"code" api:"required"`
	Message string                                     `json:"message" api:"required"`
	JSON    workflowDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// workflowDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WorkflowDeleteResponseEnvelopeMessages]
type workflowDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type WorkflowDeleteResponseEnvelopeSuccess bool

const (
	WorkflowDeleteResponseEnvelopeSuccessTrue WorkflowDeleteResponseEnvelopeSuccess = true
)

func (r WorkflowDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WorkflowDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type WorkflowDeleteResponseEnvelopeResultInfo struct {
	Count      float64                                      `json:"count" api:"required"`
	PerPage    float64                                      `json:"per_page" api:"required"`
	TotalCount float64                                      `json:"total_count" api:"required"`
	Cursor     string                                       `json:"cursor"`
	Page       float64                                      `json:"page"`
	JSON       workflowDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// workflowDeleteResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [WorkflowDeleteResponseEnvelopeResultInfo]
type workflowDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	Cursor      apijson.Field
	Page        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type WorkflowGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type WorkflowGetResponseEnvelope struct {
	Errors     []WorkflowGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages   []WorkflowGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result     WorkflowGetResponse                   `json:"result" api:"required"`
	Success    WorkflowGetResponseEnvelopeSuccess    `json:"success" api:"required"`
	ResultInfo WorkflowGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       workflowGetResponseEnvelopeJSON       `json:"-"`
}

// workflowGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [WorkflowGetResponseEnvelope]
type workflowGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WorkflowGetResponseEnvelopeErrors struct {
	Code    float64                               `json:"code" api:"required"`
	Message string                                `json:"message" api:"required"`
	JSON    workflowGetResponseEnvelopeErrorsJSON `json:"-"`
}

// workflowGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [WorkflowGetResponseEnvelopeErrors]
type workflowGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WorkflowGetResponseEnvelopeMessages struct {
	Code    float64                                 `json:"code" api:"required"`
	Message string                                  `json:"message" api:"required"`
	JSON    workflowGetResponseEnvelopeMessagesJSON `json:"-"`
}

// workflowGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WorkflowGetResponseEnvelopeMessages]
type workflowGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type WorkflowGetResponseEnvelopeSuccess bool

const (
	WorkflowGetResponseEnvelopeSuccessTrue WorkflowGetResponseEnvelopeSuccess = true
)

func (r WorkflowGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WorkflowGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type WorkflowGetResponseEnvelopeResultInfo struct {
	Count      float64                                   `json:"count" api:"required"`
	PerPage    float64                                   `json:"per_page" api:"required"`
	TotalCount float64                                   `json:"total_count" api:"required"`
	Cursor     string                                    `json:"cursor"`
	Page       float64                                   `json:"page"`
	JSON       workflowGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// workflowGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [WorkflowGetResponseEnvelopeResultInfo]
type workflowGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	Cursor      apijson.Field
	Page        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
