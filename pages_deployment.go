package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// SizeOptions can be passed to a list request to configure size and cursor location
// These values will be defaulted if omitted
type SizeOptions struct {
	Size   int `json:"size,omitempty"`
	Before int `json:"before,omitempty"`
	After  int `json:"after,omitempty"`
}

// PagesDeploymentStageLogEntry represents the logs for a Pages deployment stage.
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

// ListPagesDeployments returns all deployments for a Pages project.
//
// API reference: https://api.cloudflare.com/#pages-deployment-get-deployments
func (api *API) ListPagesDeployments(ctx context.Context, accountID string, projectName string, pageOpts PaginationOptions) ([]PagesProjectDeployment, ResultInfo, error) {
	v := url.Values{}
	if pageOpts.PerPage > 0 {
		v.Set("per_page", strconv.Itoa(pageOpts.PerPage))
	}
	if pageOpts.Page > 0 {
		v.Set("page", strconv.Itoa(pageOpts.Page))
	}

	uri := fmt.Sprintf("/accounts/%s/pages/projects/%s/deployments", accountID, projectName)
	if len(v) > 0 {
		uri = fmt.Sprintf("%s?%s", uri, v.Encode())
	}

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []PagesProjectDeployment{}, ResultInfo{}, err
	}
	var r pagesDeploymentListResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return []PagesProjectDeployment{}, ResultInfo{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, r.ResultInfo, nil
}

// GetPagesDeploymentInfo returns a deployment for a Pages project.
//
// API reference: https://api.cloudflare.com/#pages-deployment-get-deployment-info
func (api *API) GetPagesDeploymentInfo(ctx context.Context, accountID string, projectName string, deploymentID string) (PagesProjectDeployment, error) {
	uri := fmt.Sprintf("/accounts/%s/pages/projects/%s/deployments/%s", accountID, projectName, deploymentID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return PagesProjectDeployment{}, err
	}
	var r pagesDeploymentResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return PagesProjectDeployment{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// GetPagesDeploymentStageLogs returns the logs for a Pages deployment stage.
//
// API reference: https://api.cloudflare.com/#pages-deployment-get-deployment-stage-logs
func (api *API) GetPagesDeploymentStageLogs(ctx context.Context, accountID string, projectName string, deploymentID string, stageName string, sizeOpts SizeOptions) (PagesDeploymentStageLogs, error) {
	v := url.Values{}
	if sizeOpts.Size > 0 {
		v.Set("size", strconv.Itoa(sizeOpts.Size))
	}
	if sizeOpts.Before > 0 {
		v.Set("before", strconv.Itoa(sizeOpts.Before))
	} else if sizeOpts.After > 0 {
		v.Set("after", strconv.Itoa(sizeOpts.After))
	}

	uri := fmt.Sprintf("/accounts/%s/pages/projects/%s/deployments/%s/history/%s/logs", accountID, projectName, deploymentID, stageName)
	if len(v) > 0 {
		uri = fmt.Sprintf("%s?%s", uri, v.Encode())
	}

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return PagesDeploymentStageLogs{}, err
	}
	var r pagesDeploymentStageLogsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return PagesDeploymentStageLogs{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// DeletePagesDeployment deletes a Pages deployment.
//
// API reference: https://api.cloudflare.com/#pages-deployment-delete-deployment
func (api *API) DeletePagesDeployment(ctx context.Context, accountID string, projectName string, deploymentID string) error {
	uri := fmt.Sprintf("/accounts/%s/pages/projects/%s/deployments/%s", accountID, projectName, deploymentID)

	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	return nil
}

// CreatePagesDeployment creates a Pages production deployment.
//
// API reference: https://api.cloudflare.com/#pages-deployment-create-deployment
func (api *API) CreatePagesDeployment(ctx context.Context, accountID string, projectName string) (PagesProjectDeployment, error) {
	uri := fmt.Sprintf("/accounts/%s/pages/projects/%s/deployments", accountID, projectName)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return PagesProjectDeployment{}, err
	}
	var r pagesDeploymentResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return PagesProjectDeployment{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// RetryPagesDeployment retries a specific Pages deployment.
//
// API reference: https://api.cloudflare.com/#pages-deployment-retry-deployment
func (api *API) RetryPagesDeployment(ctx context.Context, accountID string, projectName string, deploymentID string) (PagesProjectDeployment, error) {
	uri := fmt.Sprintf("/accounts/%s/pages/projects/%s/deployments/%s/retry", accountID, projectName, deploymentID)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return PagesProjectDeployment{}, err
	}
	var r pagesDeploymentResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return PagesProjectDeployment{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// RollbackPagesDeployment rollbacks the Pages production deployment to a previous production deployment.
//
// API reference: https://api.cloudflare.com/#pages-deployment-rollback-deployment
func (api *API) RollbackPagesDeployment(ctx context.Context, accountID string, projectName string, deploymentID string) (PagesProjectDeployment, error) {
	uri := fmt.Sprintf("/accounts/%s/pages/projects/%s/deployments/%s/rollback", accountID, projectName, deploymentID)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return PagesProjectDeployment{}, err
	}
	var r pagesDeploymentResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return PagesProjectDeployment{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}
