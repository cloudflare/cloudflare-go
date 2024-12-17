package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

// GatewayCategory represents a single gateway category.
type GatewayCategory struct {
	Beta          *bool             `json:"beta,omitempty"`
	Class         string            `json:"class"`
	Description   string            `json:"description"`
	ID            int               `json:"id"`
	Name          string            `json:"name"`
	Subcategories []GatewayCategory `json:"subcategories"`
}

// GatewayCategoriesResponse represents the response from the list
// gateway categories endpoint.
type GatewayCategoriesResponse struct {
	Success    bool              `json:"success"`
	Result     []GatewayCategory `json:"result"`
	Errors     []string          `json:"errors"`
	Messages   []string          `json:"messages"`
	ResultInfo ResultInfo        `json:"result_info"`
}

// ListGatewayCategoriesParams represents the parameters for listing gateway categories.
type ListGatewayCategoriesParams struct {
	ResultInfo
}

// ListGatewayCategories returns all gateway categories within an account.
//
// API reference: https://developers.cloudflare.com/api/resources/zero_trust/subresources/gateway/subresources/categories/methods/list/
func (api *API) ListGatewayCategories(ctx context.Context, rc *ResourceContainer, params ListGatewayCategoriesParams) ([]GatewayCategory, ResultInfo, error) {
	uri := fmt.Sprintf("/accounts/%s/gateway/categories", rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []GatewayCategory{}, ResultInfo{}, err
	}

	var gResponse GatewayCategoriesResponse
	err = json.Unmarshal(res, &gResponse)
	if err != nil {
		return []GatewayCategory{}, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return gResponse.Result, gResponse.ResultInfo, nil
}
