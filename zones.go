package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

type ZonesService service

type ZoneParams struct {
	Match       string `url:"match,omitempty"`
	Name        string `url:"name,omitempty"`
	AccountName string `url:"account.name,omitempty"`
	Status      string `url:"status,omitempty"`
	AccountID   string `url:"account.id,omitempty"`
	Direction   string `url:"direction,omitempty"`

	PaginationOptions
}

// Get fetches a single zone.
//
// API reference: https://api.cloudflare.com/#zone-zone-details
func (s *ZonesService) Get(ctx context.Context, zoneID string) (Zone, error) {
	if !isValidZoneIdentifier(zoneID) {
		return Zone{}, fmt.Errorf(errInvalidZoneIdentifer, zoneID)
	}

	res, _ := s.client.Call(context.Background(), http.MethodGet, "/zones/"+zoneID, nil)

	var r ZoneResponse
	err := json.Unmarshal(res, &r)
	if err != nil {
		return Zone{}, fmt.Errorf("failed to unmarshal zone JSON data: %w", err)
	}

	return r.Result, nil
}

// List returns all zones that match the provided `ZoneParams` struct.
//
// API reference: https://api.cloudflare.com/#zone-list-zones
func (s *ZonesService) List(ctx context.Context, params ZoneParams) ([]Zone, error) {
	v, _ := query.Values(params)
	queryParams := v.Encode()
	if queryParams != "" {
		queryParams = "?" + queryParams
	}

	res, _ := s.client.Call(context.Background(), http.MethodGet, "/zones"+queryParams, nil)

	var r ZonesResponse
	err := json.Unmarshal(res, &r)
	if err != nil {
		return []Zone{}, fmt.Errorf("failed to unmarshal zone JSON data: %w", err)
	}

	return r.Result, nil
}

// Delete deletes a zone based on ID.
//
// API reference: https://api.cloudflare.com/#zone-delete-zone
func (s *ZonesService) Delete(ctx context.Context, zoneID string) error {
	if !isValidZoneIdentifier(zoneID) {
		return fmt.Errorf(errInvalidZoneIdentifer, zoneID)
	}

	res, _ := s.client.Call(context.Background(), http.MethodDelete, "/zones/"+zoneID, nil)

	var r ZoneResponse
	err := json.Unmarshal(res, &r)
	if err != nil {
		return fmt.Errorf("failed to unmarshal zone JSON data: %w", err)
	}

	return nil
}
