package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

type ZarazConfig = map[string]interface{}

type ZarazConfigResponse struct {
	Result ZarazConfig `json:"result"`
	Response
}

type ZarazWorkflowResponse struct {
	Result string `json:"result"`
	Response
}

type ZarazPublishResponse struct {
	Result string `json:"result"`
	Response
}

type UpdateZarazConfigParams = map[string]interface{}

type ZarazConfigRow struct {
	ID          int64      `json:"id,omitempty"`
	UserId      string     `json:"usedId,omitempty"`
	Description string     `json:"description,omitempty"`
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
}

type GetZarazConfigHistoryResponse struct {
	Result []ZarazConfigRow `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

type GetZarazConfigHistoryParams struct {
	ResultInfo
}

type GetZarazConfigsByIdResponse = map[string]interface{}

// listZarazConfigHistoryDefaultPageSize represents the default per_page size of the API.
var listZarazConfigHistoryDefaultPageSize int = 100

func (api *API) GetZarazConfig(ctx context.Context, rc *ResourceContainer) (ZarazConfigResponse, error) {
	if rc.Identifier == "" {
		return ZarazConfigResponse{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/settings/zaraz/v2/config", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ZarazConfigResponse{}, err
	}

	var recordResp ZarazConfigResponse
	err = json.Unmarshal(res, &recordResp)
	if err != nil {
		return ZarazConfigResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return recordResp, nil
}

func (api *API) UpdateZarazConfig(ctx context.Context, rc *ResourceContainer, params UpdateZarazConfigParams) (ZarazConfigResponse, error) {
	if rc.Identifier == "" {
		return ZarazConfigResponse{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/settings/zaraz/v2/config", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return ZarazConfigResponse{}, err
	}

	var updateResp ZarazConfigResponse
	err = json.Unmarshal(res, &updateResp)
	if err != nil {
		return ZarazConfigResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return updateResp, nil
}

func (api *API) GetZarazWorkflow(ctx context.Context, rc *ResourceContainer) (ZarazWorkflowResponse, error) {
	if rc.Identifier == "" {
		return ZarazWorkflowResponse{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/settings/zaraz/v2/workflow", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ZarazWorkflowResponse{}, err
	}

	var response ZarazWorkflowResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return ZarazWorkflowResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response, nil
}

func (api *API) UpdateZarazWorkflow(ctx context.Context, rc *ResourceContainer, workflowToUpdate string) (ZarazWorkflowResponse, error) {
	if rc.Identifier == "" {
		return ZarazWorkflowResponse{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/settings/zaraz/v2/workflow", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, workflowToUpdate)
	if err != nil {
		return ZarazWorkflowResponse{}, err
	}

	var response ZarazWorkflowResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return ZarazWorkflowResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response, nil
}

func (api *API) PublishZarazConfig(ctx context.Context, rc *ResourceContainer, description string) (ZarazPublishResponse, error) {
	if rc.Identifier == "" {
		return ZarazPublishResponse{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/settings/zaraz/v2/publish", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, description)
	if err != nil {
		return ZarazPublishResponse{}, err
	}

	var response ZarazPublishResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return ZarazPublishResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response, nil
}

func (api *API) GetZarazConfigHistory(ctx context.Context, rc *ResourceContainer, params GetZarazConfigHistoryParams) ([]ZarazConfigRow, *ResultInfo, error) {
	if rc.Identifier == "" {
		return nil, nil, ErrMissingZoneID
	}

	autoPaginate := true
	if params.PerPage >= 1 || params.Page >= 1 {
		autoPaginate = false
	}

	if params.PerPage < 1 {
		params.PerPage = listZarazConfigHistoryDefaultPageSize
	}

	if params.Page < 1 {
		params.Page = 1
	}

	var records []ZarazConfigRow
	var lastResultInfo ResultInfo

	for {
		uri := buildURI(fmt.Sprintf("/zones/%s/settings/zaraz/v2/history", rc.Identifier), params)
		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []ZarazConfigRow{}, &ResultInfo{}, err
		}
		var listResponse GetZarazConfigHistoryResponse
		err = json.Unmarshal(res, &listResponse)
		if err != nil {
			return []ZarazConfigRow{}, &ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
		}
		records = append(records, listResponse.Result...)
		lastResultInfo = listResponse.ResultInfo
		params.ResultInfo = listResponse.ResultInfo.Next()
		if params.ResultInfo.Done() || !autoPaginate {
			break
		}
	}
	return records, &lastResultInfo, nil
}

func (api *API) GetDefaultZarazConfig(ctx context.Context, rc *ResourceContainer) (ZarazConfigResponse, error) {
	if rc.Identifier == "" {
		return ZarazConfigResponse{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/settings/zaraz/v2/default", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ZarazConfigResponse{}, err
	}

	var recordResp ZarazConfigResponse
	err = json.Unmarshal(res, &recordResp)
	if err != nil {
		return ZarazConfigResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return recordResp, nil
}

func (api *API) ExportZarazConfig(ctx context.Context, rc *ResourceContainer) error {
	if rc.Identifier == "" {
		return ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/settings/zaraz/v2/export", rc.Identifier)
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
