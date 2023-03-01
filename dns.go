package cloudflare

import (
	"context"
	"encoding/json"
	"errors"
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
	Type      string        `url:"type,omitempty"`
	Name      string        `url:"name,omitempty"`
	Content   string        `url:"content,omitempty"`
	Proxied   *bool         `url:"proxied,omitempty"`
	Comment   string        `url:"comment,omitempty"`
	Tags      []string      `url:"tag,omitempty"` // potentially multiple `tag=`
	TagMatch  string        `url:"tag-match,omitempty"`
	Order     string        `url:"order,omitempty"`
	Direction ListDirection `url:"direction,omitempty"`
	Match     string        `url:"match,omitempty"`
	Priority  *uint16       `url:"-"`

	ResultInfo
}

type UpdateDNSRecordParams struct {
	Type     string      `json:"type,omitempty"`
	Name     string      `json:"name,omitempty"`
	Content  string      `json:"content,omitempty"`
	Data     interface{} `json:"data,omitempty"` // data for: SRV, LOC
	ID       string      `json:"-"`
	Priority *uint16     `json:"-"` // internal use only
	TTL      int         `json:"ttl,omitempty"`
	Proxied  *bool       `json:"proxied,omitempty"`
	Comment  string      `json:"comment"`
	Tags     []string    `json:"tags"`
}

// DNSListResponse represents the response from the list DNS records endpoint.
type DNSListResponse struct {
	Result []DNSRecord `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// listDNSRecordsDefaultPageSize represents the default per_page size of the API.
var listDNSRecordsDefaultPageSize int = 100

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

	params.Name = toUTS46ASCII(params.Name)

	autoPaginate := true
	if params.PerPage >= 1 || params.Page >= 1 {
		autoPaginate = false
	}

	if params.PerPage < 1 {
		params.PerPage = listDNSRecordsDefaultPageSize
	}

	if params.Page < 1 {
		params.Page = 1
	}

	var records []DNSRecord
	var lastResultInfo ResultInfo

	for {
		uri := buildURI(fmt.Sprintf("/zones/%s/dns_records", rc.Identifier), params)
		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []DNSRecord{}, &ResultInfo{}, err
		}
		var listResponse DNSListResponse
		err = json.Unmarshal(res, &listResponse)
		if err != nil {
			return []DNSRecord{}, &ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
		}
		records = append(records, listResponse.Result...)
		lastResultInfo = listResponse.ResultInfo
		params.ResultInfo = listResponse.ResultInfo.Next()
		if params.ResultInfo.Done() || !autoPaginate {
			break
		}
	}
	return records, &lastResultInfo, nil
}

// ErrMissingDNSRecordID is for when DNS record ID is needed but not given.
var ErrMissingDNSRecordID = errors.New("required DNS record ID missing")

// GetDNSRecord returns a single DNS record for the given zone & record
// identifiers.
//
// API reference: https://api.cloudflare.com/#dns-records-for-a-zone-dns-record-details
func (api *API) GetDNSRecord(ctx context.Context, rc *ResourceContainer, recordID string) (DNSRecord, error) {
	if rc.Identifier == "" {
		return DNSRecord{}, ErrMissingZoneID
	}
	if recordID == "" {
		return DNSRecord{}, ErrMissingDNSRecordID
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

	if params.ID == "" {
		return ErrMissingDNSRecordID
	}

	params.Name = toUTS46ASCII(params.Name)

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
	if recordID == "" {
		return ErrMissingDNSRecordID
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
