package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type CustomNameserverRecord struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type CustomNameserver struct {
	NSName string `json:"ns_name"`
	NSSet  int    `json:"ns_set"`
}

type CustomNameserverResult struct {
	DNSRecords []CustomNameserverRecord `json:"dns_records"`
	NSName     string                   `json:"ns_name"`
	NSSet      int                      `json:"ns_set"`
	Status     string                   `json:"status"`
	ZoneTag    string                   `json:"zone_tag"`
}

type CustomNameserverZoneMetadata struct {
	Type    string `json:"type"`
	NSSet   string `json:"ns_set"`
	Enabled bool   `json:"enabled"`
}

type customNameserverListResponse struct {
	Response
	Result []CustomNameserverResult `json:"result"`
}

type customNameserverCreateResponse struct {
	Response
	Result CustomNameserverResult `json:"result"`
}

type customNameserverZoneMetadata struct {
	Response
	Result CustomNameserverZoneMetadata
}

// GetAccountCustomNameservers lists an account's custom nameservers.
//
// API documentation: https://developers.cloudflare.com/api/operations/account-level-custom-nameservers-list-account-custom-nameservers
func (api *API) GetAccountCustomNameservers(ctx context.Context, rc *ResourceContainer) ([]CustomNameserverResult, error) {
	uri := fmt.Sprintf("/accounts/%s/custom_ns", rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	response := &customNameserverListResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, nil
}

// CreateAccountCustomNameserver adds an account custom nameserver.
//
// API documentation: https://developers.cloudflare.com/api/operations/account-level-custom-nameservers-add-account-custom-nameserver
func (api *API) CreateAccountCustomNameserver(ctx context.Context, rc *ResourceContainer, param CustomNameserver) (CustomNameserverResult, error) {
	uri := fmt.Sprintf("/accounts/%s/custom_ns", rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, param)
	if err != nil {
		return CustomNameserverResult{}, err
	}

	response := &customNameserverCreateResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return CustomNameserverResult{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, nil
}

// DeleteAccountCustomNameserver removes an account custom nameserver.
//
// API documentation: https://developers.cloudflare.com/api/operations/account-level-custom-nameservers-delete-account-custom-nameserver
func (api *API) DeleteAccountCustomNameserver(ctx context.Context, rc *ResourceContainer, nsName string) error {
	uri := fmt.Sprintf("/accounts/%s/custom_ns/%s", rc.Identifier, nsName)

	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetEligibleZonesAccountCustomNameservers lists zones eligible for account custom nameservers.
//
// API documentation: https://developers.cloudflare.com/api/operations/account-level-custom-nameservers-get-eligible-zones-for-account-custom-nameservers
func (api *API) GetEligibleZonesAccountCustomNameservers(ctx context.Context, rc *ResourceContainer) ([]string, error) {
	uri := fmt.Sprintf("/accounts/%s/custom_ns/availability", rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	response := struct {
		Result []string `json:"result"`
	}{}

	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, nil
}

// GetAccountCustomNameserverZoneMetadata get metadata for account-level custom nameservers on a zone.
//
// API documentation: https://developers.cloudflare.com/api/operations/account-level-custom-nameservers-usage-for-a-zone-get-account-custom-nameserver-related-zone-metadata
func (api *API) GetAccountCustomNameserverZoneMetadata(ctx context.Context, rc *ResourceContainer) (CustomNameserverZoneMetadata, error) {
	uri := fmt.Sprintf("/zones/%s/custom_ns", rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return CustomNameserverZoneMetadata{}, err
	}

	response := &customNameserverZoneMetadata{}

	err = json.Unmarshal(res, &response)
	if err != nil {
		return CustomNameserverZoneMetadata{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, nil
}

// UpdateAccountCustomNameserverZoneMetadata set metadata for account-level custom nameservers on a zone.
//
// API documentation: https://developers.cloudflare.com/api/operations/account-level-custom-nameservers-usage-for-a-zone-set-account-custom-nameserver-related-zone-metadata
func (api *API) UpdateAccountCustomNameserverZoneMetadata(ctx context.Context, rc *ResourceContainer, param CustomNameserverZoneMetadata) error {
	uri := fmt.Sprintf("/zones/%s/custom_ns", rc.Identifier)

	param.Type = ""
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, param)
	if err != nil {
		return nil
	}

	fmt.Println(string(res))

	return nil
}
