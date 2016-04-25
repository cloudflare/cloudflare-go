package cloudflare

import (
	"encoding/json"
	"fmt"
	"net/url"
)

/*
Create a DNS record.

API reference:
  https://api.cloudflare.com/#dns-records-for-a-zone-create-dns-record
  POST /zones/:zone_identifier/dns_records
*/
func (api *API) CreateDNSRecord(zid string, rr DNSRecord) error {
	uri := "/zones/" + zid + "/dns_records"
	res, err := api.makeRequest("POST", uri, rr)
	if err != nil {
		fmt.Println("Error with makeRequest")
		return err
	}
	var r DNSRecordResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		fmt.Println("Error with unmarshal")
		return err
	}
	return nil
}

/*
Fetches DNS records for a zone.

API reference:
  https://api.cloudflare.com/#dns-records-for-a-zone-list-dns-records
  GET /zones/:zone_identifier/dns_records
*/
func (api *API) DNSRecords(zid string, rr DNSRecord) ([]DNSRecord, error) {
	// Construct a query string
	v := url.Values{}
	if rr.Name != "" {
		v.Set("name", rr.Name)
	}
	if rr.Type != "" {
		v.Set("type", rr.Type)
	}
	if rr.Content != "" {
		v.Set("content", rr.Content)
	}
	var query string
	if len(v) > 0 {
		query = "?" + v.Encode()
	}
	uri := "/zones/" + zid + "/dns_records" + query
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return []DNSRecord{}, err
	}
	var r DNSListResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return []DNSRecord{}, err
	}
	return r.Result, nil
}

/*
Fetches a single DNS record.

API reference:
  https://api.cloudflare.com/#dns-records-for-a-zone-dns-record-details
  GET /zones/:zone_identifier/dns_records/:identifier
*/
func (api *API) DNSRecord(zid, rid string) (DNSRecord, error) {
	uri := "/zones/" + zid + "/dns_records/" + rid
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return DNSRecord{}, err
	}
	var r DNSRecordResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return DNSRecord{}, err
	}
	return r.Result, nil
}

/*
Change a DNS record.

API reference:
  https://api.cloudflare.com/#dns-records-for-a-zone-update-dns-record
  PUT /zones/:zone_identifier/dns_records/:identifier
*/
func (api *API) UpdateDNSRecord(zid, rid string, rr DNSRecord) error {
	rec, err := api.DNSRecord(zid, rid)
	if err != nil {
		return err
	}
	rr.Name = rec.Name
	rr.Type = rec.Type
	uri := "/zones/" + zid + "/dns_records/" + rid
	res, err := api.makeRequest("PUT", uri, rr)
	if err != nil {
		fmt.Println("Error with makeRequest")
		return err
	}
	var r DNSRecordResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		fmt.Println("Error with unmarshal")
		return err
	}
	return err
}

/*
Delete a DNS record.

API reference:
  https://api.cloudflare.com/#dns-records-for-a-zone-delete-dns-record
  DELETE /zones/:zone_identifier/dns_records/:identifier
*/
func (api *API) DeleteDNSRecord(zid, rid string) error {
	uri := "/zones/" + zid + "/dns_records/" + rid
	res, err := api.makeRequest("DELETE", uri, nil)
	if err != nil {
		fmt.Println("Error with makeRequest")
		return err
	}
	var r DNSRecordResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		fmt.Println("Error with unmarshal")
		return err
	}
	return nil
}
