package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

type ZoneIdentifier string

type ZonesService service

type ZoneCreateParams struct {
	Name      string   `json:"name"`
	JumpStart bool     `json:"jump_start"`
	Type      string   `json:"type"`
	Account   *Account `json:"organization,omitempty"`
}

type ZoneParams struct {
	Match       string `url:"match,omitempty"`
	Name        string `url:"name,omitempty"`
	AccountName string `url:"account.name,omitempty"`
	Status      string `url:"status,omitempty"`
	AccountID   string `url:"account.id,omitempty"`
	Direction   string `url:"direction,omitempty"`
	Page        int    `json:"page,omitempty" url:"page,omitempty"`
	PerPage     int    `json:"per_page,omitempty" url:"per_page,omitempty"`
}

// ZoneIdentifierValue accepts a string and returns a ZoneIdentifier for use in
// methods that require the stricter type.
func ZoneIdentifierValue(z string) ZoneIdentifier {
	return ZoneIdentifier(z)
}

func (z ZoneIdentifier) String() string {
	return string(z)
}

func (z ZoneIdentifier) Validate() error {
	matches, _ := regexp.MatchString(`^[0-9a-fA-F]{32}$`, z.String())
	if !matches {
		return fmt.Errorf(errInvalidZoneIdentifer, z)
	}
	return nil
}

// New creates a new zone.
//
// API reference: https://api.cloudflare.com/#zone-zone-details
func (s *ZonesService) New(ctx context.Context, zone *ZoneCreateParams) (Zone, error) {
	res, err := s.client.Call(ctx, http.MethodPost, "/zones", zone)
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
func (s *ZonesService) Get(ctx context.Context, zoneID ZoneIdentifier) (Zone, error) {
	if err := zoneID.Validate(); err != nil {
		return Zone{}, err
	}

	res, _ := s.client.Call(ctx, http.MethodGet, "/zones/"+zoneID.String(), nil)

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
func (s *ZonesService) List(ctx context.Context, params *ZoneParams) ([]Zone, *ResultInfo, error) {
	res, _ := s.client.Call(ctx, http.MethodGet, buildURI("/zones", params), nil)

	var r ZonesResponse
	err := json.Unmarshal(res, &r)
	if err != nil {
		return []Zone{}, &ResultInfo{}, fmt.Errorf("failed to unmarshal zone JSON data: %w", err)
	}

	return r.Result, &r.ResultInfo, nil
}

// Delete deletes a zone based on ID.
//
// API reference: https://api.cloudflare.com/#zone-delete-zone
func (s *ZonesService) Delete(ctx context.Context, zoneID ZoneIdentifier) error {
	if err := zoneID.Validate(); err != nil {
		return err
	}

	res, _ := s.client.Call(ctx, http.MethodDelete, "/zones/"+zoneID.String(), nil)

	var r ZoneResponse
	err := json.Unmarshal(res, &r)
	if err != nil {
		return fmt.Errorf("failed to unmarshal zone JSON data: %w", err)
	}

	return nil
}
