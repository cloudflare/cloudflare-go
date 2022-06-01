package cloudflare

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type TieredCacheSmartTopologyValue string

const (
	TieredCacheSmartTopologyOn  TieredCacheSmartTopologyValue = "on"
	TieredCacheSmartTopologyOff TieredCacheSmartTopologyValue = "off"
)

type TieredCacheSmartTopologyUpdateOptions struct {
	Value TieredCacheSmartTopologyValue `json:"value"`
}

type TieredCacheSmartTopology struct {
	ID         string                        `json:"id"`
	ModifiedOn string                        `json:"modified_on"`
	Value      TieredCacheSmartTopologyValue `json:"value"`
}

type TieredCacheSmartTopologyResponse struct {
	Response
	Result TieredCacheSmartTopology `json:"result"`
}

// TieredCacheSmartTopology reads topology for cache.
//
// API reference: TBD.
func (api *API) TieredCacheSmartTopology(ctx context.Context, zoneID string) (TieredCacheSmartTopology, error) {
	res, err := api.makeRequestContext(ctx, http.MethodGet, "/zones/"+zoneID+"/cache/tiered_cache_smart_topology_enable", nil)
	if err != nil {
		return TieredCacheSmartTopology{}, err
	}
	response := TieredCacheSmartTopologyResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return TieredCacheSmartTopology{}, errors.Wrap(err, errUnmarshalError)
	}
	return response.Result, nil
}

// UpdateTieredCacheSmartTopology updates topology for cache.
//
// API reference: TBD.
func (api *API) UpdateTieredCacheSmartTopology(ctx context.Context, zoneID string, options TieredCacheSmartTopologyUpdateOptions) (TieredCacheSmartTopology, error) {
	res, err := api.makeRequestContext(ctx, http.MethodPatch, "/zones/"+zoneID+"/cache/tiered_cache_smart_topology_enable", options)
	if err != nil {
		return TieredCacheSmartTopology{}, err
	}
	response := TieredCacheSmartTopologyResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return TieredCacheSmartTopology{}, errors.Wrap(err, errUnmarshalError)
	}
	return response.Result, nil
}
