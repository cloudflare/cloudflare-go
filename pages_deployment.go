package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"errors"
)

// SizeOptions can be passed to a list request to configure size and cursor location.
// These values will be defaulted if omitted.
//
// This should be swapped to ResultInfoCursors once the types are corrected.
type SizeOptions struct {
	Size   int  `json:"size,omitempty" url:"size,omitempty"`
	Before *int `json:"before,omitempty" url:"before,omitempty"`
	After  *int `json:"after,omitempty" url:"after,omitempty"`
}

// PagesDeploymentStageLogs represents the logs for a Pages deployment stage.
type PagesDeploymentStageLogs struct {
	Name      string                         `json:"name"`
	StartedOn *time.Time                     `json:"started_on"`
	EndedOn   *time.Time                     `json:"ended_on"`
	Status    string                         `json:"status"`
	Start     int                            `json:"start"`
	End       int                            `json:"end"`
	Total     int                            `json:"total"`
	Data      []PagesDeploymentStageLogEntry `json:"data"`
}

// PagesDeploymentStageLogEntry represents a single log entry in a Pages deployment stage.
type PagesDeploymentStageLogEntry struct {
	ID        int        `json:"id"`
	Timestamp *time.Time `json:"timestamp"`
	Message   string     `json:"message"`
}

// PagesDeploymentLogs represents the logs for a Pages deployment.
type PagesDeploymentLogs struct {
	Total                 int                       `json:"total"`
	IncludesContainerLogs bool                      `json:"includes_container_logs"`
	Data                  []PagesDeploymentLogEntry `json:"data"`
}

// PagesDeploymentLogEntry represents a single log entry in a Pages deployment.
type PagesDeploymentLogEntry struct {
	Timestamp *time.Time `json:"ts"`
	Line      string     `json:"line"`
}

type pagesDeploymentListResponse struct {
	Response
	Result     []PagesProjectDeployment `json:"result"`
	ResultInfo `json:"result_info"`
}

type pagesDeploymentResponse struct {
	Response
	Result PagesProjectDeployment `json:"result"`
}

type pagesDeploymentStageLogsResponse struct {
	Response
	Result     PagesDeploymentStageLogs `json:"result"`
	ResultInfo `json:"result_info"`
}

type pagesDeploymentLogsResponse struct {
	Response
	Result     PagesDeploymentLogs `json:"result"`
	ResultInfo `json:"result_info"`
}

type ListPagesDeploymentsParams struct {
	AccountID   string
	ProjectName string

	PaginationOptions
}

type GetPagesDeploymentInfoParams struct {
	AccountID    string
	ProjectName  string
	DeploymentID string
}

type GetPagesDeploymentStageLogsParams struct {
	AccountID    string
	ProjectName  string
	DeploymentID string
	StageName    string

	SizeOptions
}

type GetPagesDeploymentLogsParams struct {
	AccountID    string
	ProjectName  string
	DeploymentID string

	SizeOptions
}

type DeletePagesDeploymentParams struct {
	AccountID    string
	ProjectName  string
	DeploymentID string
}

type CreatePagesDeploymentParams struct {
	AccountID   string
	ProjectName string
}

type RetryPagesDeploymentParams struct {
	AccountID    string
	ProjectName  string
	DeploymentID string
}

type RollbackPagesDeploymentParams struct {
	AccountID    string
	ProjectName  string
	DeploymentID string
}

var (
	ErrMissingProjectName  = errors.New("required missing project name")
	ErrMissingDeploymentID = errors.New("required missing deployment ID")
	ErrMissingStageName    = errors.New("required missing stage name")
)

// ListPagesDeployments returns all deployments for a Pages project.
//
// API reference: https://api.cloudflare.com/#pages-deployment-get-deployments
func (api *API) ListPagesDeployments(ctx context.Context, params ListPagesDeploymentsParams) ([]PagesProjectDeployment, ResultInfo, error) {
	if params.AccountID == "" {
		return []PagesProjectDeployment{}, ResultInfo{}, ErrMissingAccountID
	}

	if params.ProjectName == "" {
		return []PagesProjectDeployment{}, ResultInfo{}, ErrMissingProjectName
	}

	uri := buildURI(fmt.Sprintf("/accounts/%s/pages/projects/%s/deployments", params.AccountID, params.ProjectName), params.PaginationOptions)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []PagesProjectDeployment{}, ResultInfo{}, err
	}
	var r pagesDeploymentListResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return []PagesProjectDeployment{}, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, r.ResultInfo, nil
}

// GetPagesDeploymentInfo returns a deployment for a Pages project.
//
// API reference: https://api.cloudflare.com/#pages-deployment-get-deployment-info
func (api *API) GetPagesDeploymentInfo(ctx context.Context, params GetPagesDeploymentInfoParams) (PagesProjectDeployment, error) {
	if params.AccountID == "" {
		return PagesProjectDeployment{}, ErrMissingAccountID
	}

	if params.ProjectName == "" {
		return PagesProjectDeployment{}, ErrMissingProjectName
	}

	if params.DeploymentID == "" {
		return PagesProjectDeployment{}, ErrMissingDeploymentID
	}

	uri := fmt.Sprintf("/accounts/%s/pages/projects/%s/deployments/%s", params.AccountID, params.ProjectName, params.DeploymentID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return PagesProjectDeployment{}, err
	}
	var r pagesDeploymentResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return PagesProjectDeployment{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// GetPagesDeploymentStageLogs returns the logs for a Pages deployment stage.
//
// API reference: https://api.cloudflare.com/#pages-deployment-get-deployment-stage-logs
//
// Deprecated: Use GetPagesDeploymentLogs instead.
func (api *API) GetPagesDeploymentStageLogs(ctx context.Context, params GetPagesDeploymentStageLogsParams) (PagesDeploymentStageLogs, error) {
	if params.AccountID == "" {
		return PagesDeploymentStageLogs{}, ErrMissingAccountID
	}

	if params.ProjectName == "" {
		return PagesDeploymentStageLogs{}, ErrMissingProjectName
	}

	if params.DeploymentID == "" {
		return PagesDeploymentStageLogs{}, ErrMissingDeploymentID
	}

	if params.StageName == "" {
		return PagesDeploymentStageLogs{}, ErrMissingStageName
	}

	uri := buildURI(
		fmt.Sprintf("/accounts/%s/pages/projects/%s/deployments/%s/history/%s/logs", params.AccountID, params.ProjectName, params.DeploymentID, params.StageName),
		params.SizeOptions,
	)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return PagesDeploymentStageLogs{}, err
	}
	var r pagesDeploymentStageLogsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return PagesDeploymentStageLogs{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// GetPagesDeploymentStageLogs returns the logs for a Pages deployment stage.
//
// API reference: https://api.cloudflare.com/#pages-deployment-get-deployment-logs
func (api *API) GetPagesDeploymentLogs(ctx context.Context, params GetPagesDeploymentLogsParams) (PagesDeploymentLogs, error) {
	if params.AccountID == "" {
		return PagesDeploymentLogs{}, ErrMissingAccountID
	}

	if params.ProjectName == "" {
		return PagesDeploymentLogs{}, ErrMissingProjectName
	}

	if params.DeploymentID == "" {
		return PagesDeploymentLogs{}, ErrMissingDeploymentID
	}

	uri := buildURI(
		fmt.Sprintf("/accounts/%s/pages/projects/%s/deployments/%s/history/logs", params.AccountID, params.ProjectName, params.DeploymentID),
		params.SizeOptions,
	)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return PagesDeploymentLogs{}, err
	}
	var r pagesDeploymentLogsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return PagesDeploymentLogs{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// DeletePagesDeployment deletes a Pages deployment.
//
// API reference: https://api.cloudflare.com/#pages-deployment-delete-deployment
func (api *API) DeletePagesDeployment(ctx context.Context, params DeletePagesDeploymentParams) error {
	if params.AccountID == "" {
		return ErrMissingAccountID
	}

	if params.ProjectName == "" {
		return ErrMissingProjectName
	}

	if params.DeploymentID == "" {
		return ErrMissingDeploymentID
	}

	uri := fmt.Sprintf("/accounts/%s/pages/projects/%s/deployments/%s", params.AccountID, params.ProjectName, params.DeploymentID)

	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	return nil
}

// CreatePagesDeployment creates a Pages production deployment.
//
// API reference: https://api.cloudflare.com/#pages-deployment-create-deployment
func (api *API) CreatePagesDeployment(ctx context.Context, params CreatePagesDeploymentParams) (PagesProjectDeployment, error) {
	if params.AccountID == "" {
		return PagesProjectDeployment{}, ErrMissingAccountID
	}

	if params.ProjectName == "" {
		return PagesProjectDeployment{}, ErrMissingProjectName
	}

	uri := fmt.Sprintf("/accounts/%s/pages/projects/%s/deployments", params.AccountID, params.ProjectName)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return PagesProjectDeployment{}, err
	}
	var r pagesDeploymentResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return PagesProjectDeployment{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// RetryPagesDeployment retries a specific Pages deployment.
//
// API reference: https://api.cloudflare.com/#pages-deployment-retry-deployment
func (api *API) RetryPagesDeployment(ctx context.Context, params RetryPagesDeploymentParams) (PagesProjectDeployment, error) {
	if params.AccountID == "" {
		return PagesProjectDeployment{}, ErrMissingAccountID
	}

	if params.ProjectName == "" {
		return PagesProjectDeployment{}, ErrMissingProjectName
	}

	if params.DeploymentID == "" {
		return PagesProjectDeployment{}, ErrMissingDeploymentID
	}

	uri := fmt.Sprintf("/accounts/%s/pages/projects/%s/deployments/%s/retry", params.AccountID, params.ProjectName, params.DeploymentID)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return PagesProjectDeployment{}, err
	}
	var r pagesDeploymentResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return PagesProjectDeployment{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// RollbackPagesDeployment rollbacks the Pages production deployment to a previous production deployment.
//
// API reference: https://api.cloudflare.com/#pages-deployment-rollback-deployment
func (api *API) RollbackPagesDeployment(ctx context.Context, params RollbackPagesDeploymentParams) (PagesProjectDeployment, error) {
	if params.AccountID == "" {
		return PagesProjectDeployment{}, ErrMissingAccountID
	}

	if params.ProjectName == "" {
		return PagesProjectDeployment{}, ErrMissingProjectName
	}

	if params.DeploymentID == "" {
		return PagesProjectDeployment{}, ErrMissingDeploymentID
	}

	uri := fmt.Sprintf("/accounts/%s/pages/projects/%s/deployments/%s/rollback", params.AccountID, params.ProjectName, params.DeploymentID)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return PagesProjectDeployment{}, err
	}
	var r pagesDeploymentResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return PagesProjectDeployment{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}
