package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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
	Comment    string      `json:"comment,omitempty"`
	Tags       []string    `json:"tags,omitempty"`
}

// DNSRecordResponse represents the response from the DNS endpoint.
type DNSRecordResponse struct {
	Result DNSRecord `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

type ListDirection string

const (
	ListDirectionAsc  ListDirection = "asc"
	ListDirectionDesc ListDirection = "desc"
)

type ListDNSRecordsParams struct {
	CreatedOn  time.Time     `json:"created_on,omitempty" url:"created_on,omitempty"`
	ModifiedOn time.Time     `json:"modified_on,omitempty" url:"modified_on,omitempty"`
	Type       string        `json:"type,omitempty" url:"type,omitempty"`
	Name       string        `json:"name,omitempty" url:"name,omitempty"`
	Content    string        `json:"content,omitempty" url:"content,omitempty"`
	Proxied    *bool         `json:"proxied,omitempty" url:"proxied,omitempty"`
	Comment    string        `json:"comment,omitempty" url:"comment,omitempty"`
	Tags       []string      `json:"tags,omitempty"`
	Order      string        `url:"order,omitempty"`
	Direction  ListDirection `url:"direction,omitempty"`
	Match      string        `url:"match,omitempty"`
	Priority   *uint16       `json:"priority,omitempty"`

	ResultInfo
}

type UpdateDNSRecordParams struct {
	CreatedOn  time.Time   `json:"created_on,omitempty" url:"created_on,omitempty"`
	ModifiedOn time.Time   `json:"modified_on,omitempty" url:"modified_on,omitempty"`
	Type       string      `json:"type,omitempty" url:"type,omitempty"`
	Name       string      `json:"name,omitempty" url:"name,omitempty"`
	Content    string      `json:"content,omitempty" url:"content,omitempty"`
	Meta       interface{} `json:"meta,omitempty"`
	Data       interface{} `json:"data,omitempty"` // data returned by: SRV, LOC
	ID         string      `json:"id,omitempty"`
	ZoneID     string      `json:"zone_id,omitempty"`
	ZoneName   string      `json:"zone_name,omitempty"`
	Priority   *uint16     `json:"priority,omitempty"`
	TTL        int         `json:"ttl,omitempty"`
	Proxied    *bool       `json:"proxied,omitempty" url:"proxied,omitempty"`
	Proxiable  bool        `json:"proxiable,omitempty"`
	Locked     bool        `json:"locked,omitempty"`
	Comment    string      `json:"comment,omitempty" url:"comment,omitempty"`
	Tags       []string    `json:"tags,omitempty"`
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

type CreateDNSRecordParams struct {
	CreatedOn  time.Time   `json:"created_on,omitempty" url:"created_on,omitempty"`
	ModifiedOn time.Time   `json:"modified_on,omitempty" url:"modified_on,omitempty"`
	Type       string      `json:"type,omitempty" url:"type,omitempty"`
	Name       string      `json:"name,omitempty" url:"name,omitempty"`
	Content    string      `json:"content,omitempty" url:"content,omitempty"`
	Meta       interface{} `json:"meta,omitempty"`
	Data       interface{} `json:"data,omitempty"` // data returned by: SRV, LOC
	ID         string      `json:"id,omitempty"`
	ZoneID     string      `json:"zone_id,omitempty"`
	ZoneName   string      `json:"zone_name,omitempty"`
	Priority   *uint16     `json:"priority,omitempty"`
	TTL        int         `json:"ttl,omitempty"`
	Proxied    *bool       `json:"proxied,omitempty" url:"proxied,omitempty"`
	Proxiable  bool        `json:"proxiable,omitempty"`
	Locked     bool        `json:"locked,omitempty"`
	Comment    string      `json:"comment,omitempty" url:"comment,omitempty"`
	Tags       []string    `json:"tags,omitempty"`
}

// CreateDNSRecord creates a DNS record for the zone identifier.
//
// API reference: https://api.cloudflare.com/#dns-records-for-a-zone-create-dns-record
func (api *API) CreateDNSRecord(ctx context.Context, rc *ResourceContainer, params CreateDNSRecordParams) (*DNSRecordResponse, error) {
	if rc.Identifier == "" {
		return nil, ErrMissingZoneID
	}
	params.Name = toUTS46ASCII(params.Name)

	uri := fmt.Sprintf("/zones/%s/dns_records", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
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

// ListDNSRecords returns a slice of DNS records for the given zone identifier.
//
// API reference: https://api.cloudflare.com/#dns-records-for-a-zone-list-dns-records
func (api *API) ListDNSRecords(ctx context.Context, rc *ResourceContainer, params ListDNSRecordsParams) ([]DNSRecord, *ResultInfo, error) {
	if rc.Identifier == "" {
		return nil, nil, ErrMissingZoneID
	}

	if params.Name != "" {
		params.Name = toUTS46ASCII(params.Name)
	}

	autoPaginate := true
	if params.PerPage >= 1 || params.Page >= 1 {
		autoPaginate = false
	}

	if params.PerPage < 1 {
		params.PerPage = 50
	}

	if params.Page < 1 {
		params.Page = 1
	}

	var records []DNSRecord
	var listResponse DNSListResponse

	for {
		uri := buildURI(fmt.Sprintf("/zones/%s/dns_records", rc.Identifier), params)
		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []DNSRecord{}, &ResultInfo{}, err
		}
		err = json.Unmarshal(res, &listResponse)
		if err != nil {
			return []DNSRecord{}, &ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
		}
		records = append(records, listResponse.Result...)
		params.ResultInfo = listResponse.ResultInfo.Next()
		if params.ResultInfo.Done() || !autoPaginate {
			break
		}
	}
	return records, &listResponse.ResultInfo, nil
}

// GetDNSRecord returns a single DNS record for the given zone & record
// identifiers.
//
// API reference: https://api.cloudflare.com/#dns-records-for-a-zone-dns-record-details
func (api *API) GetDNSRecord(ctx context.Context, rc *ResourceContainer, recordID string) (DNSRecord, error) {
	if rc.Identifier == "" {
		return DNSRecord{}, ErrMissingZoneID
	}
	uri := fmt.Sprintf("/zones/%s/dns_records/%s", rc.Identifier, recordID)
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
func (api *API) UpdateDNSRecord(ctx context.Context, rc *ResourceContainer, params UpdateDNSRecordParams) error {
	if rc.Identifier == "" {
		return ErrMissingZoneID
	}

	params.Name = toUTS46ASCII(params.Name)

	// Populate the record name from the existing one if the update didn't
	// specify it.
	if params.Name == "" || params.Type == "" {
		rec, err := api.GetDNSRecord(ctx, rc, params.ID)
		if err != nil {
			return err
		}

		if params.Name == "" {
			params.Name = rec.Name
		}
		if params.Type == "" {
			params.Type = rec.Type
		}
	}

	uri := fmt.Sprintf("/zones/%s/dns_records/%s", rc.Identifier, params.ID)
	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, params)
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
func (api *API) DeleteDNSRecord(ctx context.Context, rc *ResourceContainer, recordID string) error {
	if rc.Identifier == "" {
		return ErrMissingZoneID
	}
	uri := fmt.Sprintf("/zones/%s/dns_records/%s", rc.Identifier, recordID)
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
