package cloudflare

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

// LogpullRetentionConfiguration describes a the structure of a Logpull Retention
// payload.
type LogpullRetentionConfiguration struct {
	Flag bool `json:"flag"`
}

// LogpullRetentionConfigurationResponse is the API response, containing the
// Logpull retention result.
type LogpullRetentionConfigurationResponse struct {
	Response
	Result LogpullRetentionConfiguration `json:"result"`
}

// GetLogpullRetentionFlag gets the current setting flag.
//
// API reference: https://developers.cloudflare.com/logs/logpull-api/enabling-log-retention/
func (api *API) GetLogpullRetentionFlag(ctx context.Context, zoneID string) (*LogpullRetentionConfiguration, error) {
	uri := "/zones/" + zoneID + "/logs/control/retention/flag"
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return &LogpullRetentionConfiguration{}, err
	}
	var r LogpullRetentionConfigurationResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}
	return &r.Result, nil
}

// SetLogpullRetentionFlag updates the retention flag to the defined boolean.
//
// API reference: https://developers.cloudflare.com/logs/logpull-api/enabling-log-retention/
func (api *API) SetLogpullRetentionFlag(ctx context.Context, zoneID string, enabled bool) (*LogpullRetentionConfiguration, error) {
	uri := "/zones/" + zoneID + "/logs/control/retention/flag"
	flagPayload := LogpullRetentionConfiguration{Flag: enabled}

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, flagPayload)
	if err != nil {
		return &LogpullRetentionConfiguration{}, err
	}
	var r LogpullRetentionConfigurationResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return &LogpullRetentionConfiguration{}, err
	}
	return &r.Result, nil
}
