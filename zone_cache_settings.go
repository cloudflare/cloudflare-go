package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// ZoneSetting contains settings for a zone.
type ZoneCacheSetting struct {
	ID         string      `json:"id"`
	ModifiedOn string      `json:"modified_on,omitempty"`
	Value      interface{} `json:"value"`
}

type ZoneCacheSettingSingleResponse struct {
	Response
	Result ZoneCacheSetting `json:"result"`
}

// ZoneSingleCacheSetting returns information about specified cache setting to the specified zone.
//
// API reference: https://api.cloudflare.com/#zone-cache-settings-properties
func (api *API) ZoneSingleCacheSetting(ctx context.Context, zoneID, settingName string) (ZoneCacheSetting, error) {
	uri := fmt.Sprintf("/zones/%s/cache/%s", zoneID, settingName)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ZoneCacheSetting{}, err
	}
	var r ZoneCacheSettingSingleResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return ZoneCacheSetting{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// UpdateZoneSingleCacheSetting updates the specified cache setting for a given zone.
//
// API reference: https://api.cloudflare.com/#zone-cache-settings-properties
func (api *API) UpdateZoneSingleCacheSetting(ctx context.Context, zoneID, settingName string, setting ZoneCacheSetting) (ZoneCacheSetting, error) {
	uri := fmt.Sprintf("/zones/%s/cache/%s", zoneID, settingName)
	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, setting)
	if err != nil {
		return ZoneCacheSetting{}, err
	}

	response := &ZoneCacheSettingSingleResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return ZoneCacheSetting{}, errors.Wrap(err, errUnmarshalError)
	}

	return response.Result, nil
}

// DeleteZoneSingleCacheSetting deletes the specified cache setting for a given zone.
//
// API reference: https://api.cloudflare.com/#zone-cache-settings-properties
func (api *API) DeleteZoneSingleCacheSetting(ctx context.Context, zoneID, settingName string) error {
	uri := fmt.Sprintf("/zones/%s/cache/%s", zoneID, settingName)
	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return nil
}
