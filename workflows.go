package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

// WorkflowInstance represents a single instance of a workflow.
type WorkflowInstance struct {
	ID         string     `json:"id"`
	CreatedOn  time.Time  `json:"created_on"`
	ModifiedOn time.Time  `json:"modified_on"`
	StartedOn  *time.Time `json:"started_on"`
	EndedOn    *time.Time `json:"ended_on"`
	WorkflowID string     `json:"workflow_id"`
	VersionID  string     `json:"version_id"`
	Status     string     `json:"status"`
}

// Workflow represents a workflow and its configuration.
type Workflow struct {
	Name        string     `json:"name"`
	ID          string     `json:"id"`
	CreatedOn   time.Time  `json:"created_on"`
	ModifiedOn  time.Time  `json:"modified_on"`
	ScriptName  string     `json:"script_name"`
	ClassName   string     `json:"class_name"`
	TriggeredOn *time.Time `json:"triggered_on"`
	Instances   struct {
		Complete        int `json:"complete"`
		Errored        int `json:"errored"`
		Paused         int `json:"paused"`
		Queued         int `json:"queued"`
		Running        int `json:"running"`
		Terminated     int `json:"terminated"`
		Unknown        int `json:"unknown"`
		Waiting        int `json:"waiting"`
		WaitingForPause int `json:"waitingForPause"`
	} `json:"instances"`
}

// WorkflowVersion represents a specific version of a workflow.
type WorkflowVersion struct {
	CreatedOn  time.Time `json:"created_on"`
	ModifiedOn time.Time `json:"modified_on"`
	ID         string    `json:"id"`
	WorkflowID string    `json:"workflow_id"`
	ClassName  string    `json:"class_name"`
}

// ListWorkflowsParams provides parameters for listing workflows.
type ListWorkflowsParams struct {
	PerPage int `url:"per_page,omitempty"`
	Page    int `url:"page,omitempty"`
}

// ListWorkflowInstancesParams provides parameters for listing workflow instances.
type ListWorkflowInstancesParams struct {
	PerPage    int       `url:"per_page,omitempty"`
	Page       int       `url:"page,omitempty"`
	Status     string    `url:"status,omitempty"`
	DateStart  time.Time `url:"date_start,omitempty"`
	DateEnd    time.Time `url:"date_end,omitempty"`
}

// CreateWorkflowParams provides parameters for creating or modifying a workflow.
type CreateWorkflowParams struct {
	ScriptName string `json:"script_name"`
	ClassName  string `json:"class_name"`
}

// ListWorkflowVersionsParams provides parameters for listing workflow versions.
type ListWorkflowVersionsParams struct {
	PerPage int `url:"per_page,omitempty"`
	Page    int `url:"page,omitempty"`
}

// WorkflowResponse is the response for a single workflow operation.
type WorkflowResponse struct {
	Response
	Result Workflow `json:"result"`
}

// WorkflowListResponse is the response for listing workflows.
type WorkflowListResponse struct {
	Response
	Result     []Workflow `json:"result"`
	ResultInfo ResultInfo `json:"result_info"`
}

// WorkflowInstanceResponse is the response for a single workflow instance operation.
type WorkflowInstanceResponse struct {
	Response
	Result WorkflowInstance `json:"result"`
}

// WorkflowInstanceListResponse is the response for listing workflow instances.
type WorkflowInstanceListResponse struct {
	Response
	Result     []WorkflowInstance `json:"result"`
	ResultInfo ResultInfo         `json:"result_info"`
}

// WorkflowVersionResponse is the response for a single workflow version operation.
type WorkflowVersionResponse struct {
	Response
	Result WorkflowVersion `json:"result"`
}

// WorkflowVersionListResponse is the response for listing workflow versions.
type WorkflowVersionListResponse struct {
	Response
	Result     []WorkflowVersion `json:"result"`
	ResultInfo ResultInfo        `json:"result_info"`
}

// ListWorkflows returns a list of workflows for an account.
func (api *API) ListWorkflows(ctx context.Context, rc *ResourceContainer, params ListWorkflowsParams) ([]Workflow, ResultInfo, error) {
	if rc.Level != AccountRouteLevel {
		return []Workflow{}, ResultInfo{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return []Workflow{}, ResultInfo{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/workflows", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []Workflow{}, ResultInfo{}, err
	}

	var workflowListResponse WorkflowListResponse
	err = json.Unmarshal(res, &workflowListResponse)
	if err != nil {
		return []Workflow{}, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return workflowListResponse.Result, workflowListResponse.ResultInfo, nil
}

// CreateWorkflow creates or modifies a workflow.
func (api *API) CreateWorkflow(ctx context.Context, rc *ResourceContainer, workflowName string, params CreateWorkflowParams) (Workflow, error) {
	if rc.Level != AccountRouteLevel {
		return Workflow{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return Workflow{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/workflows/%s", rc.Identifier, workflowName)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return Workflow{}, err
	}

	var workflowResponse WorkflowResponse
	err = json.Unmarshal(res, &workflowResponse)
	if err != nil {
		return Workflow{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return workflowResponse.Result, nil
}

// GetWorkflow returns a specific workflow's details.
func (api *API) GetWorkflow(ctx context.Context, rc *ResourceContainer, workflowName string) (Workflow, error) {
	if rc.Level != AccountRouteLevel {
		return Workflow{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return Workflow{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/workflows/%s", rc.Identifier, workflowName)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return Workflow{}, err
	}

	var workflowResponse WorkflowResponse
	err = json.Unmarshal(res, &workflowResponse)
	if err != nil {
		return Workflow{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return workflowResponse.Result, nil
}

// ListWorkflowInstances lists all instances of a specific workflow.
func (api *API) ListWorkflowInstances(ctx context.Context, rc *ResourceContainer, workflowName string, params ListWorkflowInstancesParams) ([]WorkflowInstance, ResultInfo, error) {
	if rc.Level != AccountRouteLevel {
		return []WorkflowInstance{}, ResultInfo{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return []WorkflowInstance{}, ResultInfo{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/workflows/%s/instances", rc.Identifier, workflowName)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []WorkflowInstance{}, ResultInfo{}, err
	}

	var instanceListResponse WorkflowInstanceListResponse
	err = json.Unmarshal(res, &instanceListResponse)
	if err != nil {
		return []WorkflowInstance{}, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return instanceListResponse.Result, instanceListResponse.ResultInfo, nil
}

// GetWorkflowInstance returns details of a specific workflow instance.
func (api *API) GetWorkflowInstance(ctx context.Context, rc *ResourceContainer, workflowName, instanceID string) (WorkflowInstance, error) {
	if rc.Level != AccountRouteLevel {
		return WorkflowInstance{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return WorkflowInstance{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/workflows/%s/instances/%s", rc.Identifier, workflowName, instanceID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return WorkflowInstance{}, err
	}

	var instanceResponse WorkflowInstanceResponse
	err = json.Unmarshal(res, &instanceResponse)
	if err != nil {
		return WorkflowInstance{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return instanceResponse.Result, nil
}

// CreateWorkflowInstance creates a new workflow instance.
func (api *API) CreateWorkflowInstance(ctx context.Context, rc *ResourceContainer, workflowName string, params map[string]interface{}) (WorkflowInstance, error) {
	if rc.Level != AccountRouteLevel {
		return WorkflowInstance{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return WorkflowInstance{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/workflows/%s/instances", rc.Identifier, workflowName)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return WorkflowInstance{}, err
	}

	var instanceResponse WorkflowInstanceResponse
	err = json.Unmarshal(res, &instanceResponse)
	if err != nil {
		return WorkflowInstance{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return instanceResponse.Result, nil
}

// ListWorkflowVersions lists all versions of a specific workflow.
func (api *API) ListWorkflowVersions(ctx context.Context, rc *ResourceContainer, workflowName string, params ListWorkflowVersionsParams) ([]WorkflowVersion, ResultInfo, error) {
	if rc.Level != AccountRouteLevel {
		return []WorkflowVersion{}, ResultInfo{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return []WorkflowVersion{}, ResultInfo{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/workflows/%s/versions", rc.Identifier, workflowName)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []WorkflowVersion{}, ResultInfo{}, err
	}

	var versionListResponse WorkflowVersionListResponse
	err = json.Unmarshal(res, &versionListResponse)
	if err != nil {
		return []WorkflowVersion{}, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return versionListResponse.Result, versionListResponse.ResultInfo, nil
}

// GetWorkflowVersion returns details of a specific workflow version.
func (api *API) GetWorkflowVersion(ctx context.Context, rc *ResourceContainer, workflowName, versionID string) (WorkflowVersion, error) {
	if rc.Level != AccountRouteLevel {
		return WorkflowVersion{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return WorkflowVersion{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/workflows/%s/versions/%s", rc.Identifier, workflowName, versionID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return WorkflowVersion{}, err
	}

	var versionResponse WorkflowVersionResponse
	err = json.Unmarshal(res, &versionResponse)
	if err != nil {
		return WorkflowVersion{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return versionResponse.Result, nil
}