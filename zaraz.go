package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

type ZarazConfig = map[string]interface{}

type UpdateZarazConfigParams = map[string]interface{}

type ZarazConfigRow struct {
	ID          int64     `json:"id,omitempty"`
	UserId      string    `json:"usedId,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
}

type GetZarazConfigHistoryResponse struct {
	Data  []ZarazConfigRow `json:"data"`
	Count int              `json:"count"`
}

type GetZarazConfigsByIdResponse = map[string]interface{}

func (api *API) GetZarazConfig(ctx context.Context, rc *ResourceContainer) (ZarazConfig, error) {
	if rc.Identifier == "" {
		return ZarazConfig{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/settings/zaraz/config", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ZarazConfig{}, err
	}

	var recordResp ZarazConfig
	err = json.Unmarshal(res, &recordResp)
	if err != nil {
		return ZarazConfig{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return recordResp, nil
}

func (api *API) UpdateZarazConfig(ctx context.Context, rc *ResourceContainer, params UpdateZarazConfigParams) error {
	if rc.Identifier == "" {
		return ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/settings/zaraz/config", rc.Identifier)
	_, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return err
	}

	return nil
}

func (api *API) GetZarazWorkflow(ctx context.Context, rc *ResourceContainer) (string, error) {
	if rc.Identifier == "" {
		return "", ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/settings/zaraz/workflow", rc.Identifier)
	workflow, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return "", err
	}

	return string(workflow), nil
}

func (api *API) UpdateZarazWorkflow(ctx context.Context, rc *ResourceContainer, workflowToUpdate string) error {
	if rc.Identifier == "" {
		return ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/settings/zaraz/workflow", rc.Identifier)
	_, err := api.makeRequestContext(ctx, http.MethodPut, uri, workflowToUpdate)
	if err != nil {
		return err
	}

	return nil
}

func (api *API) PublishZarazConfig(ctx context.Context, rc *ResourceContainer, description string) error {
	if rc.Identifier == "" {
		return ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/settings/zaraz/publish", rc.Identifier)
	_, err := api.makeRequestContext(ctx, http.MethodPost, uri, description)
	if err != nil {
		return err
	}

	return nil
}

func (api *API) GetZarazConfigHistory(ctx context.Context, rc *ResourceContainer, limit int64, offset int64) (GetZarazConfigHistoryResponse, error) {
	if rc.Identifier == "" {
		return GetZarazConfigHistoryResponse{
			Data:  []ZarazConfigRow{},
			Count: 0,
		}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/settings/zaraz/history?limit=%d&offset=%d", rc.Identifier, limit, offset)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return GetZarazConfigHistoryResponse{
			Data:  []ZarazConfigRow{},
			Count: 0,
		}, err
	}

	var recordResp GetZarazConfigHistoryResponse
	err = json.Unmarshal(res, &recordResp)
	if err != nil {
		return GetZarazConfigHistoryResponse{
			Data:  []ZarazConfigRow{},
			Count: 0,
		}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return recordResp, nil
}

func (api *API) GetDefaultZarazConfig(ctx context.Context, rc *ResourceContainer) (ZarazConfig, error) {
	if rc.Identifier == "" {
		return ZarazConfig{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/settings/zaraz/default", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ZarazConfig{}, err
	}

	var recordResp ZarazConfig
	err = json.Unmarshal(res, &recordResp)
	if err != nil {
		return ZarazConfig{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return recordResp, nil
}

func (api *API) ExportZarazConfig(ctx context.Context, rc *ResourceContainer) error {
	if rc.Identifier == "" {
		return ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/settings/zaraz/export", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return err
	}

	var recordResp ZarazConfig
	err = json.Unmarshal(res, &recordResp)
	if err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return nil
}
