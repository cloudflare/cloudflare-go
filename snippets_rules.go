package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

type SnippetsRulesResponse struct {
	Response
	Result []SnippetRule `json:"result"`
}

type SnippetRule struct {
	ID          string `json:"id"`
	Enabled     *bool  `json:"enabled,omitempty"`
	Expression  string `json:"expression"`
	SnippetName string `json:"snippet_name"`
	Description string `json:"description"`
}

func (api *API) ListZoneSnippetsRules(ctx context.Context, rc *ResourceContainer) ([]SnippetRule, error) {
	if rc.Identifier == "" {
		return nil, ErrMissingZoneID
	}

	uri := buildURI(fmt.Sprintf("/zones/%s/snippets/snippet_rules", rc.Identifier), nil)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	result := SnippetsRulesResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}

func (api *API) UpdateZoneSnippetsRules(ctx context.Context, rc *ResourceContainer, params []SnippetRule) ([]SnippetRule, error) {
	if rc.Identifier == "" {
		return nil, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/snippets/snippet_rules", rc.Identifier)

	payload, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, payload)
	if err != nil {
		return nil, err
	}

	result := SnippetsRulesResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}
