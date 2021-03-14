package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)
// SecondaryDNSZone contains the high level structure of a secondary DNS zone.
type SecondaryDNSZone struct {
	ID                 string    `json:"id,omitempty"`
	Name               string    `json:"name,omitempty"`
	Primaries          []string  `json:"primaries,omitempty"`
	AutoRefreshSeconds int       `json:"auto_refresh_seconds,omitempty"`
	SoaSerial          int       `json:"soa_serial,omitempty"`
	CreatedTime        time.Time `json:"created_time,omitempty"`
	CheckedTime        time.Time `json:"checked_time,omitempty"`
	ModifiedTime       time.Time `json:"modified_time,omitempty"`
}

// SecondaryDNSZoneDetailResponse is the API response for a single secondary
// DNS zone.
type SecondaryDNSZoneDetailResponse struct {
	Response
	Result SecondaryDNSZone `json:"result"`
}
// GetSecondaryDNSZone returns the secondary DNS zone configuration for a
// single zone.
//
// API reference: https://api.cloudflare.com/#secondary-dns-secondary-zone-configuration-details
func (api *API) GetSecondaryDNSZone(ctx context.Context, zoneID string) (SecondaryDNSZone, error) {
	uri := fmt.Sprintf("/zones/%s/secondary_dns", zoneID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return SecondaryDNSZone{}, err
	}

	var r SecondaryDNSZoneDetailResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return SecondaryDNSZone{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// CreateSecondaryDNSZone creates a secondary DNS zone.
//
// API reference: https://api.cloudflare.com/#secondary-dns-create-secondary-zone-configuration
func (api *API) CreateSecondaryDNSZone(ctx context.Context, zoneID string, zone SecondaryDNSZone) (SecondaryDNSZone, error) {
	if err := validateRequiredSecondaryDNSZoneValues(zone); err != nil {
		return SecondaryDNSZone{}, err
	}

	uri := fmt.Sprintf("/zones/%s/secondary_dns", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri,
		SecondaryDNSZone{
			Name:               zone.Name,
			AutoRefreshSeconds: zone.AutoRefreshSeconds,
			Primaries:          zone.Primaries,
		},
	)

	if err != nil {
		return SecondaryDNSZone{}, err
	}

	result := SecondaryDNSZoneDetailResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return SecondaryDNSZone{}, errors.Wrap(err, errUnmarshalError)
	}

	return result.Result, nil
}
