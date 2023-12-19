package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

type ZarazConfig = map[string]interface{}

type UpdateZarazConfigParams = map[string]interface{}

func (api *API) GetZarazConfig(ctx context.Context, rc *ResourceContainer) (ZarazConfig, error) {
	if rc.Identifier == "" {
		return ZarazConfig{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/settings/zaraz/config", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ZarazConfig{}, err
	}

	var recordResp ZarazConfig
	err = json.Unmarshal(res, &recordResp)
	if err != nil {
		return ZarazConfig{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return recordResp, nil
}

func (api *API) UpdateZarazConfig(ctx context.Context, rc *ResourceContainer, params UpdateZarazConfigParams) error {
	if rc.Identifier == "" {
		return ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/settings/zaraz/config", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return err
	}

	var recordResp *ZarazConfig
	err = json.Unmarshal(res, &recordResp)
	if err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return nil
}
