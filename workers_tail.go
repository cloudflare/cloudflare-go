package cloudflare

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var (
	ErrMissingScriptName = errors.New("required script name missing")
	ErrMissingTailID     = errors.New("required tail id missing")
)

type WorkersTail struct {
	ID        string     `json:"id"`
	URL       string     `json:"url"`
	ExpiresAt *time.Time `json:"expires_at"`
}

type StartWorkersTailParameters struct {
	AccountID  string
	ScriptName string
}

type StartWorkersTailResponse struct {
	Response
	Result WorkersTail
}

type ListWorkersTailParameters struct {
	AccountID  string
	ScriptName string
}

type ListWorkersTailResponse struct {
	Response
	Result WorkersTail
}
type DeleteWorkersTailParameters struct {
	AccountID  string
	ScriptName string
	TailID     string
}

// StartWorkersTail Starts a tail that receives logs and exception from a Worker
//
// API reference: https://api.cloudflare.com/#worker-tail-logs-start-tail
func (api *API) StartWorkersTail(ctx context.Context, params StartWorkersTailParameters) (WorkersTail, error) {
	if params.AccountID == "" {
		return WorkersTail{}, ErrMissingAccountID
	}
	if params.ScriptName == "" {
		return WorkersTail{}, ErrMissingScriptName
	}

	uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s/tails", params.AccountID, params.ScriptName)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return WorkersTail{}, err
	}
	var workerstailResponse StartWorkersTailResponse
	if err := json.Unmarshal(res, &workerstailResponse); err != nil {
		return WorkersTail{}, err
	}
	return workerstailResponse.Result, nil
}

// ListWorkersTail Get list of tails currently deployed on a worker
//
// API reference: https://api.cloudflare.com/#worker-tail-logs-list-tails
func (api *API) ListWorkersTail(ctx context.Context, params ListWorkersTailParameters) (WorkersTail, error) {
	if params.AccountID == "" {
		return WorkersTail{}, ErrMissingAccountID
	}
	if params.ScriptName == "" {
		return WorkersTail{}, ErrMissingScriptName
	}

	uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s/tails", params.AccountID, params.ScriptName)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return WorkersTail{}, err
	}
	var workerstailResponse ListWorkersTailResponse
	if err := json.Unmarshal(res, &workerstailResponse); err != nil {
		return WorkersTail{}, err
	}
	return workerstailResponse.Result, nil
}

// DeleteWorkersTail Deletes a tail from a Worker
//
// API reference: https://api.cloudflare.com/#worker-tail-logs-delete-tail
func (api *API) DeleteWorkersTail(ctx context.Context, params DeleteWorkersTailParameters) error {
	if params.AccountID == "" {
		return ErrMissingAccountID
	}
	if params.ScriptName == "" {
		return ErrMissingScriptName
	}
	if params.TailID == "" {
		return ErrMissingTailID
	}

	uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s/tails/%s", params.AccountID, params.ScriptName, params.TailID)
	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	return nil
}
