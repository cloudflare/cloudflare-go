package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"golang.org/x/net/idna"
)

// DNSRecord represents a DNS record in a zone.
type DNSRecord struct {
	CreatedOn  time.Time   `json:"created_on,omitempty"`
	ModifiedOn time.Time   `json:"modified_on,omitempty"`
	Type       string      `json:"type,omitempty"`
	Name       string      `json:"name,omitempty"`
	Content    string      `json:"content,omitempty"`
	Meta       interface{} `json:"meta,omitempty"`
	Data       interface{} `json:"data,omitempty"` // data returned by: SRV, LOC
	ID         string      `json:"id,omitempty"`
	ZoneID     string      `json:"zone_id,omitempty"`
	ZoneName   string      `json:"zone_name,omitempty"`
	Priority   *uint16     `json:"priority,omitempty"`
	TTL        int         `json:"ttl,omitempty"`
	Proxied    *bool       `json:"proxied,omitempty"`
	Proxiable  bool        `json:"proxiable,omitempty"`
	Locked     bool        `json:"locked,omitempty"`
}

// DNSRecordResponse represents the response from the DNS endpoint.
type DNSRecordResponse struct {
	Result DNSRecord `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// DNSListResponse represents the response from the list DNS records endpoint.
type DNSListResponse struct {
	Result []DNSRecord `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// nontransitionalLookup implements the nontransitional processing as specified in
// Unicode Technical Standard 46 with almost all checkings off to maximize user freedom.
var nontransitionalLookup = idna.New(
	idna.MapForLookup(),
	idna.StrictDomainName(false),
	idna.ValidateLabels(false),
)

// toUTS46ASCII tries to convert IDNs (international domain names)
// from Unicode form to Punycode, using non-transitional process specified
// in UTS 46.
//
// Note: conversion errors are silently discarded and partial conversion
// results are used.
func toUTS46ASCII(name string) string {
	name, _ = nontransitionalLookup.ToASCII(name)
	return name
}

// CreateDNSRecord creates a DNS record for the zone identifier.
//
// API reference: https://api.cloudflare.com/#dns-records-for-a-zone-create-dns-record
func (api *API) CreateDNSRecord(ctx context.Context, zoneID string, rr DNSRecord) (*DNSRecordResponse, error) {
	rr.Name = toUTS46ASCII(rr.Name)

	uri := fmt.Sprintf("/zones/%s/dns_records", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, rr)
	if err != nil {
		return nil, err
	}

	var recordResp *DNSRecordResponse
	err = json.Unmarshal(res, &recordResp)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return recordResp, nil
}

// DNSRecords returns a slice of DNS records for the given zone identifier.
//
// This takes a DNSRecord to allow filtering of the results returned.
//
// API reference: https://api.cloudflare.com/#dns-records-for-a-zone-list-dns-records
func (api *API) DNSRecords(ctx context.Context, zoneID string, rr DNSRecord) ([]DNSRecord, error) {
	// Construct a query string
	v := url.Values{}
	// Using default per_page value as specified by the API
	if rr.Name != "" {
		v.Set("name", toUTS46ASCII(rr.Name))
	}
	if rr.Type != "" {
		v.Set("type", rr.Type)
	}
	if rr.Content != "" {
		v.Set("content", rr.Content)
	}

	var records []DNSRecord
	page := 1

	// Loop over makeRequest until what we've fetched all records
	for {
		v.Set("page", strconv.Itoa(page))
		uri := fmt.Sprintf("/zones/%s/dns_records?%s", zoneID, v.Encode())
		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []DNSRecord{}, err
		}
		var r DNSListResponse
		err = json.Unmarshal(res, &r)
		if err != nil {
			return []DNSRecord{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
		}
		records = append(records, r.Result...)
		if r.ResultInfo.Page >= r.ResultInfo.TotalPages {
			break
		}
		// Loop around and fetch the next page
		page++
	}
	return records, nil
}

// DNSRecord returns a single DNS record for the given zone & record
// identifiers.
//
// API reference: https://api.cloudflare.com/#dns-records-for-a-zone-dns-record-details
func (api *API) DNSRecord(ctx context.Context, zoneID, recordID string) (DNSRecord, error) {
	uri := fmt.Sprintf("/zones/%s/dns_records/%s", zoneID, recordID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return DNSRecord{}, err
	}
	var r DNSRecordResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return DNSRecord{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// UpdateDNSRecord updates a single DNS record for the given zone & record
// identifiers.
//
// API reference: https://api.cloudflare.com/#dns-records-for-a-zone-update-dns-record
func (api *API) UpdateDNSRecord(ctx context.Context, zoneID, recordID string, rr DNSRecord) error {
	rr.Name = toUTS46ASCII(rr.Name)

	// Populate the record name from the existing one if the update didn't
	// specify it.
	if rr.Name == "" || rr.Type == "" {
		rec, err := api.DNSRecord(ctx, zoneID, recordID)
		if err != nil {
			return err
		}

		if rr.Name == "" {
			rr.Name = rec.Name
		}
		if rr.Type == "" {
			rr.Type = rec.Type
		}
	}
	uri := fmt.Sprintf("/zones/%s/dns_records/%s", zoneID, recordID)
	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, rr)
	if err != nil {
		return err
	}
	var r DNSRecordResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return nil
}

// DeleteDNSRecord deletes a single DNS record for the given zone & record
// identifiers.
//
// API reference: https://api.cloudflare.com/#dns-records-for-a-zone-delete-dns-record
func (api *API) DeleteDNSRecord(ctx context.Context, zoneID, recordID string) error {
	uri := fmt.Sprintf("/zones/%s/dns_records/%s", zoneID, recordID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	var r DNSRecordResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return nil
}
