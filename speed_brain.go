package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

// SpeedBrainSetting defines the structure for speed brain settings.
type SpeedBrainSetting struct {
	ID         string     `json:"id"`
	Editable   *bool      `json:"editable,omitempty"`
	Value      string     `json:"value"`
	ModifiedOn *time.Time `json:"modified_on"`
}

// SpeedBrainSettingResponse wraps the response from the API containing the result.
type SpeedBrainSettingResponse struct {
	Response
	Result SpeedBrainSetting `json:"result"`
}

// GetSpeedBrainSettings retrieves the speed brain settings for a specific zone.
func (api *API) GetSpeedBrainSettings(ctx context.Context, zoneID string) (SpeedBrainSetting, error) {
	uri := fmt.Sprintf("/zones/%s/settings/speed_brain", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return SpeedBrainSetting{}, err
	}
	var r SpeedBrainSettingResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return SpeedBrainSetting{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// PatchSpeedBrainSettings updates the speed brain settings for a specific zone.
func (api *API) PatchSpeedBrainSettings(ctx context.Context, zoneID string, value string) (SpeedBrainSetting, error) {
	uri := fmt.Sprintf("/zones/%s/settings/speed_brain", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, SpeedBrainSetting{Value: value})
	if err != nil {
		return SpeedBrainSetting{}, err
	}
	var r SpeedBrainSettingResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return SpeedBrainSetting{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}
