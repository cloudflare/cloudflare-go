package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// ZoneCacheReserve is the structure of the API object for the
// cache reserve setting.
type ZoneCacheReserve struct {
	ID         string    `json:"id,omitempty"`
	ModifiedOn time.Time `json:"modified_on,omitempty"`
	Value      string    `json:"value"`
}

// CacheReserveDetailsResponse is the API response for the
// cache reserve setting.
type CacheReserveDetailsResponse struct {
	Result ZoneCacheReserve `json:"result"`
	Response
}

type updateZoneCacheReserveRequest struct {
	Value string `json:"value"`
}

type zoneCacheReserveSingleResponse struct {
	Response
	Result ZoneCacheReserve `json:"result"`
}

// GetZoneCacheReserve returns information about the current cache reserve settings
//
// API reference: https://developers.cloudflare.com/api/operations/zone-cache-settings-get-cache-reserve-setting
func (api *API) GetZoneCacheReserve(ctx context.Context, zoneID string) (ZoneCacheReserve, error) {
	uri := fmt.Sprintf("/zones/%s/cache/cache_reserve", zoneID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ZoneCacheReserve{}, err
	}

	var cacheReserveDetailsResponse CacheReserveDetailsResponse
	err = json.Unmarshal(res, &cacheReserveDetailsResponse)
	if err != nil {
		return ZoneCacheReserve{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return cacheReserveDetailsResponse.Result, nil
}

// UpdateZoneCacheReserve updates the cache reserve setting for a zone
//
// API reference: https://developers.cloudflare.com/api/operations/zone-cache-settings-change-cache-reserve-setting
func (api *API) UpdateZoneCacheReserve(ctx context.Context, zoneID string, value string) (ZoneCacheReserve, error) {
	uri := fmt.Sprintf("/zones/%s/cache/cache_reserve", zoneID)

	updateReq := &updateZoneCacheReserveRequest{Value: value}
	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, updateReq)
	if err != nil {
		return ZoneCacheReserve{}, err
	}

	response := &zoneCacheReserveSingleResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return ZoneCacheReserve{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, nil
}
