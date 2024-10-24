// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/packages/pagination"
)

// WorkflowService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowService] method instead.
type WorkflowService struct {
	Options   []option.RequestOption
	Instances *WorkflowInstanceService
	Versions  *WorkflowVersionService
}

// NewWorkflowService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkflowService(opts ...option.RequestOption) (r *WorkflowService) {
	r = &WorkflowService{}
	r.Options = opts
	r.Instances = NewWorkflowInstanceService(opts...)
	r.Versions = NewWorkflowVersionService(opts...)
	return
}

// Create/modify Workflow
func (r *WorkflowService) Update(ctx context.Context, workflowName string, params WorkflowUpdateParams, opts ...option.RequestOption) (res *WorkflowUpdateResponse, err error) {
	var env WorkflowUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if workflowName == "" {
		err = errors.New("missing required workflow_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workflows/%s", params.AccountID, workflowName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all Workflows
func (r *WorkflowService) List(ctx context.Context, params WorkflowListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[WorkflowListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
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

// List all Workflows
func (r *WorkflowService) ListAutoPaging(ctx context.Context, params WorkflowListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[WorkflowListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Get Workflow details
func (r *WorkflowService) Get(ctx context.Context, workflowName string, query WorkflowGetParams, opts ...option.RequestOption) (res *WorkflowGetResponse, err error) {
	var env WorkflowGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if workflowName == "" {
		err = errors.New("missing required workflow_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workflows/%s", query.AccountID, workflowName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkflowUpdateResponse struct {
	ID          string                     `json:"id,required" format:"uuid"`
	ClassName   string                     `json:"class_name,required"`
	CreatedOn   time.Time                  `json:"created_on,required" format:"date-time"`
	ModifiedOn  time.Time                  `json:"modified_on,required" format:"date-time"`
	Name        string                     `json:"name,required"`
	ScriptName  string                     `json:"script_name,required"`
	TriggeredOn time.Time                  `json:"triggered_on,required,nullable" format:"date-time"`
	VersionID   string                     `json:"version_id,required" format:"uuid"`
	JSON        workflowUpdateResponseJSON `json:"-"`
}

// workflowUpdateResponseJSON contains the JSON metadata for the struct
// [WorkflowUpdateResponse]
type workflowUpdateResponseJSON struct {
	ID          apijson.Field
	ClassName   apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	ScriptName  apijson.Field
	TriggeredOn apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type WorkflowListResponse struct {
	ID          string                        `json:"id,required" format:"uuid"`
	ClassName   string                        `json:"class_name,required"`
	CreatedOn   time.Time                     `json:"created_on,required" format:"date-time"`
	Instances   WorkflowListResponseInstances `json:"instances,required"`
	ModifiedOn  time.Time                     `json:"modified_on,required" format:"date-time"`
	Name        string                        `json:"name,required"`
	ScriptName  string                        `json:"script_name,required"`
	TriggeredOn time.Time                     `json:"triggered_on,required,nullable" format:"date-time"`
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
	Unknown         float64                           `json:"unknown"`
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
	Unknown         apijson.Field
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

type WorkflowGetResponse struct {
	ID          string                       `json:"id,required" format:"uuid"`
	ClassName   string                       `json:"class_name,required"`
	CreatedOn   time.Time                    `json:"created_on,required" format:"date-time"`
	Instances   WorkflowGetResponseInstances `json:"instances,required"`
	ModifiedOn  time.Time                    `json:"modified_on,required" format:"date-time"`
	Name        string                       `json:"name,required"`
	ScriptName  string                       `json:"script_name,required"`
	TriggeredOn time.Time                    `json:"triggered_on,required,nullable" format:"date-time"`
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
	Unknown         float64                          `json:"unknown"`
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
	Unknown         apijson.Field
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
	AccountID  param.Field[string] `path:"account_id,required"`
	ClassName  param.Field[string] `json:"class_name,required"`
	ScriptName param.Field[string] `json:"script_name,required"`
}

func (r WorkflowUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkflowUpdateResponseEnvelope struct {
	Errors     []WorkflowUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []WorkflowUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result     WorkflowUpdateResponse                   `json:"result,required"`
	Success    WorkflowUpdateResponseEnvelopeSuccess    `json:"success,required"`
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
	Code    float64                                  `json:"code,required"`
	Message string                                   `json:"message,required"`
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
	Code    float64                                    `json:"code,required"`
	Message string                                     `json:"message,required"`
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
	Count      float64                                      `json:"count,required"`
	Page       float64                                      `json:"page,required"`
	PerPage    float64                                      `json:"per_page,required"`
	TotalCount float64                                      `json:"total_count,required"`
	JSON       workflowUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// workflowUpdateResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [WorkflowUpdateResponseEnvelopeResultInfo]
type workflowUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
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
	AccountID param.Field[string]  `path:"account_id,required"`
	Page      param.Field[float64] `query:"page"`
	PerPage   param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [WorkflowListParams]'s query parameters as `url.Values`.
func (r WorkflowListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type WorkflowGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type WorkflowGetResponseEnvelope struct {
	Errors     []WorkflowGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []WorkflowGetResponseEnvelopeMessages `json:"messages,required"`
	Result     WorkflowGetResponse                   `json:"result,required"`
	Success    WorkflowGetResponseEnvelopeSuccess    `json:"success,required"`
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
	Code    float64                               `json:"code,required"`
	Message string                                `json:"message,required"`
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
	Code    float64                                 `json:"code,required"`
	Message string                                  `json:"message,required"`
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
	Count      float64                                   `json:"count,required"`
	Page       float64                                   `json:"page,required"`
	PerPage    float64                                   `json:"per_page,required"`
	TotalCount float64                                   `json:"total_count,required"`
	JSON       workflowGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// workflowGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [WorkflowGetResponseEnvelopeResultInfo]
type workflowGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkflowGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workflowGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
