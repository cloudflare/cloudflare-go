package cloudflare

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Filter holds the structure of the filter type.
type Filter struct {
	ID          string `json:"id,omitempty"`
	Expression  string `json:"expression"`
	Paused      bool   `json:"paused"`
	Description string `json:"description"`

	// Property is mentioned in documentation however isn't populated in
	// any of the API requests. For now, let's just omit it unless it's
	// provided.
	Ref string `json:"ref,omitempty"`
}

// FiltersDetailResponse is the API response that is returned
// for requesting all filters on a zone.
type FiltersDetailResponse struct {
	Result     []Filter `json:"result"`
	ResultInfo `json:"result_info"`
	Response
}

// FilterDetailResponse is the API response that is returned
// for requesting a single filter on a zone.
type FilterDetailResponse struct {
	Result     Filter `json:"result"`
	ResultInfo `json:"result_info"`
	Response
}

// FilterValidateExpression represents the JSON payload for checking
// an expression.
type FilterValidateExpression struct {
	Expression string `json:"expression"`
}

// FilterValidateExpressionResponse represents the API response for
// checking the expression. It conforms to the JSON API approach however
// we don't need all of the fields exposed.
type FilterValidateExpressionResponse struct {
	Success bool                                `json:"success"`
	Errors  []FilterValidationExpressionMessage `json:"errors"`
}

type FilterValidationExpressionMessage struct {
	Message string `json:"message"`
}

// Filter returns a single filter in a zone based on the filter ID.
//
// API reference: TBC
func (api *API) Filter(zoneID, filterID string) (Filter, error) {
	uri := fmt.Sprintf("/zones/%s/filters/%s", zoneID, filterID)

	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return Filter{}, errors.Wrap(err, errMakeRequestError)
	}

	var filterResponse FilterDetailResponse
	err = json.Unmarshal(res, &filterResponse)
	if err != nil {
		return Filter{}, errors.Wrap(err, errUnmarshalError)
	}

	return filterResponse.Result, nil
}
