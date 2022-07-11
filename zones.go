package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
)

type ZonesService service

type ZoneCreateParams struct {
	Name      string   `json:"name"`
	JumpStart bool     `json:"jump_start"`
	Type      string   `json:"type"`
	Account   *Account `json:"organization,omitempty"`
}

type ZoneListParams struct {
	Match       string `url:"match,omitempty"`
	Name        string `url:"name,omitempty"`
	AccountName string `url:"account.name,omitempty"`
	Status      string `url:"status,omitempty"`
	AccountID   string `url:"account.id,omitempty"`
	Direction   string `url:"direction,omitempty"`
	Page        int    `json:"page,omitempty" url:"page,omitempty"`
	PerPage     int    `json:"per_page,omitempty" url:"per_page,omitempty"`
}

type ZoneUpdateParams struct {
	ID                string
	Paused            *bool    `json:"paused"`
	VanityNameServers []string `json:"vanity_name_servers,omitempty"`
	Plan              ZonePlan `json:"plan,omitempty"`
	Type              string   `json:"type,omitempty"`
}

// New creates a new zone.
//
// API reference: https://api.cloudflare.com/#zone-zone-details
func (s *ZonesService) New(ctx context.Context, zone *ZoneCreateParams) (Zone, error) {
	res, err := s.client.post(ctx, "/zones", zone)
	if err != nil {
		return Zone{}, err
	}

	var r ZoneResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return Zone{}, fmt.Errorf("failed to unmarshal zone JSON data: %w", err)
	}

	return r.Result, nil
}

// Get fetches a single zone.
//
// API reference: https://api.cloudflare.com/#zone-zone-details
func (s *ZonesService) Get(ctx context.Context, rc *ResourceContainer) (Zone, error) {
	uri := fmt.Sprintf("/zones/%s", rc.Identifier)
	res, err := s.client.get(ctx, uri, nil)
	if err != nil {
		return Zone{}, fmt.Errorf("failed to fetch zones: %w", err)
	}

	var r ZoneResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return Zone{}, fmt.Errorf("failed to unmarshal zone JSON data: %w", err)
	}

	return r.Result, nil
}

// List returns all zones that match the provided `ZoneParams` struct.
//
// API reference: https://api.cloudflare.com/#zone-list-zones
func (s *ZonesService) List(ctx context.Context, params *ZoneListParams) ([]Zone, *ResultInfo, error) {
	res, _ := s.client.get(ctx, buildURI("/zones", params), nil)

	var r ZonesResponse
	err := json.Unmarshal(res, &r)
	if err != nil {
		return []Zone{}, &ResultInfo{}, fmt.Errorf("failed to unmarshal zone JSON data: %w", err)
	}

	return r.Result, &r.ResultInfo, nil
}

// Update modifies an existing zone.
//
// API reference: https://api.cloudflare.com/#zone-edit-zone
func (s *ZonesService) Update(ctx context.Context, params *ZoneUpdateParams) ([]Zone, error) {
	uri := fmt.Sprintf("/zones/%s", params.ID)
	res, _ := s.client.patch(ctx, uri, params)

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
func (s *ZonesService) Delete(ctx context.Context, rc *ResourceContainer) error {
	uri := fmt.Sprintf("/zones/%s", rc.Identifier)
	res, _ := s.client.delete(ctx, uri, nil)

	var r ZoneResponse
	err := json.Unmarshal(res, &r)
	if err != nil {
		return fmt.Errorf("failed to unmarshal zone JSON data: %w", err)
	}

	return nil
}
