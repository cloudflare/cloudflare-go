package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// APIShieldCreateOperation should be used when creating an operation in API Shield Endpoint Management.
type APIShieldCreateOperation struct {
	Method   string `json:"method"`
	Host     string `json:"host"`
	Endpoint string `json:"endpoint"`
}

// APIShieldOperation represents an operation stored in API Shield Endpoint Management.
type APIShieldOperation struct {
	APIShieldCreateOperation
	ID          string         `json:"operation_id"`
	LastUpdated time.Time      `json:"last_updated"`
	Features    map[string]any `json:"features,omitempty"`
}

// APIShieldGetOperationParams represents the query parameters to set when retrieving an operation.
//
// See API documentation https://developers.cloudflare.com/api/operations/api-shield-endpoint-management-retrieve-information-about-an-operation
type APIShieldGetOperationParams struct {
	// Features represents a set of features to return in `features` object when
	// performing making read requests against an Operation or listing operations.
	Features []string
}

// APIShieldGetOperationsParams represents the query parameters to set when retrieving operations
//
// See API documentation https://developers.cloudflare.com/api/operations/api-shield-endpoint-management-retrieve-information-about-all-operations-on-a-zone
type APIShieldGetOperationsParams struct {
	// Features represents a set of features to return in `features` object when
	// performing making read requests against an Operation or listing operations.
	Features []string
	// Direction to order results.
	Direction string
	// OrderBy when requesting a feature, the feature keys are available for ordering as well, e.g., thresholds.suggested_threshold.
	OrderBy string
	// Filters to only return operations that match filtering criteria, see APIShieldGetOperationsFilters
	Filters *APIShieldGetOperationsFilters
	// Pagination options to apply to the request.
	Pagination *PaginationOptions
}

// APIShieldGetOperationsFilters represents the filtering query parameters to set when retrieving operations
//
// See API documentation https://developers.cloudflare.com/api/operations/api-shield-endpoint-management-retrieve-information-about-all-operations-on-a-zone
type APIShieldGetOperationsFilters struct {
	// Host filters results to only include the specified hosts.
	Hosts []string
	// Host filters results to only include the specified methods.
	Methods []string
	// Endpoint filter results to only include endpoints containing this pattern.
	Endpoint string
}

// APIShieldGetOperationResponse represents the response from the api_gateway/operations/{id} endpoint.
type APIShieldGetOperationResponse struct {
	Result APIShieldOperation `json:"result"`
	Response
}

// APIShieldGetOperationsResponse represents the response from the api_gateway/operations endpoint.
type APIShieldGetOperationsResponse struct {
	Result     []APIShieldOperation `json:"result"`
	ResultInfo `json:"result_info"`
	Response
}

// APIShieldDeleteOperationResponse represents the response from the api_gateway/operations/{id} endpoint (DELETE).
type APIShieldDeleteOperationResponse struct {
	Result interface{} `json:"result"`
	Response
}

// GetAPIShieldOperation returns information about an operation
//
// API documentation https://developers.cloudflare.com/api/operations/api-shield-endpoint-management-retrieve-information-about-an-operation
func (api *API) GetAPIShieldOperation(ctx context.Context, rc *ResourceContainer, operationID string, params *APIShieldGetOperationParams) (*APIShieldOperation, error) {
	uri := fmt.Sprintf("/zones/%s/api_gateway/operations/%s", rc.Identifier, operationID)

	if params != nil {
		uri = strings.Join([]string{uri, params.Encode()}, "?")
	}

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	var asResponse APIShieldGetOperationResponse
	err = json.Unmarshal(res, &asResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return &asResponse.Result, nil
}

// GetAPIShieldOperations retrieve information about all operations on a zone
//
// API documentation https://developers.cloudflare.com/api/operations/api-shield-endpoint-management-retrieve-information-about-all-operations-on-a-zone
func (api *API) GetAPIShieldOperations(ctx context.Context, rc *ResourceContainer, params *APIShieldGetOperationsParams) ([]APIShieldOperation, ResultInfo, error) {
	uri := fmt.Sprintf("/zones/%s/api_gateway/operations", rc.Identifier)
	if params != nil {
		uri = strings.Join([]string{uri, params.Encode()}, "?")
	}

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, ResultInfo{}, err
	}

	var asResponse APIShieldGetOperationsResponse
	err = json.Unmarshal(res, &asResponse)
	if err != nil {
		return nil, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return asResponse.Result, asResponse.ResultInfo, nil
}

// PostAPIShieldOperations add one or more operations to a zone.
//
// API documentation https://developers.cloudflare.com/api/operations/api-shield-endpoint-management-add-operations-to-a-zone
func (api *API) PostAPIShieldOperations(ctx context.Context, rc *ResourceContainer, operations []APIShieldCreateOperation) ([]APIShieldOperation, error) {
	uri := fmt.Sprintf("/zones/%s/api_gateway/operations", rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, operations)
	if err != nil {
		return nil, err
	}

	// Result should be all the operations added to the zone, similar to doing GetAPIShieldOperations
	var asResponse APIShieldGetOperationsResponse
	err = json.Unmarshal(res, &asResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return asResponse.Result, nil
}

// DeleteAPIShieldOperation deletes a single operation
//
// API documentation https://developers.cloudflare.com/api/operations/api-shield-endpoint-management-delete-an-operation
func (api *API) DeleteAPIShieldOperation(ctx context.Context, rc *ResourceContainer, operationID string) error {
	uri := fmt.Sprintf("/zones/%s/api_gateway/operations/%s", rc.Identifier, operationID)

	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	var asResponse APIShieldDeleteOperationResponse
	err = json.Unmarshal(res, &asResponse)
	if err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return nil
}

// Encode is a custom method for encoding APIShieldGetOperationParams into a usable HTTP
// query parameter string.
func (a APIShieldGetOperationParams) Encode() string {
	v := url.Values{}
	for _, f := range a.Features {
		v.Add("feature", f)
	}

	return v.Encode()
}

// Encode is a custom method for encoding APIShieldGetOperationsParams into a usable HTTP
// query parameter string.
func (a APIShieldGetOperationsParams) Encode() string {
	v := url.Values{}
	for _, f := range a.Features {
		v.Add("feature", f)
	}

	if a.Direction != "" {
		v.Set("direction", a.Direction)
	}

	if a.OrderBy != "" {
		v.Set("order", a.OrderBy)
	}

	if a.Pagination != nil {
		v.Set("page", strconv.Itoa(a.Pagination.Page))
		v.Set("per_page", strconv.Itoa(a.Pagination.PerPage))
	}

	if a.Filters != nil {
		if a.Filters.Endpoint != "" {
			v.Set("endpoint", a.Filters.Endpoint)
		}

		for _, h := range a.Filters.Hosts {
			v.Add("host", h)
		}

		for _, m := range a.Filters.Methods {
			v.Add("method", m)
		}
	}

	return v.Encode()
}
