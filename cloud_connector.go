package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

type CloudConnectorRulesResponse struct {
	Response
	Result []CloudConnectorRule `json:"result"`
}

type CloudConnectorRuleParameters struct {
	Host string `json:"host"`
}

type CloudConnectorRule struct {
	ID          string                       `json:"id"`
	Enabled     *bool                        `json:"enabled,omitempty"`
	Expression  string                       `json:"expression"`
	Provider    string                       `json:"provider"`
	Parameters  CloudConnectorRuleParameters `json:"parameters"`
	Description string                       `json:"description"`
}

func (api *API) ListZoneCloudConnectorRules(ctx context.Context, rc *ResourceContainer) ([]CloudConnectorRule, error) {
	if rc.Identifier == "" {
		return nil, ErrMissingZoneID
	}

	uri := buildURI(fmt.Sprintf("/zones/%s/cloud_connector/rules", rc.Identifier), nil)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	result := CloudConnectorRulesResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}

func (api *API) UpdateZoneCloudConnectorRules(ctx context.Context, rc *ResourceContainer, params []CloudConnectorRule) ([]CloudConnectorRule, error) {
	if rc.Identifier == "" {
		return nil, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/cloud_connector/rules", rc.Identifier)

	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, payload)
	if err != nil {
		return nil, err
	}

	result := CloudConnectorRulesResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}
