package cloudflare

import (
	"encoding/json"
	"fmt"
)

/*
Create a DNS record.

API reference:
  https://api.cloudflare.com/#dns-records-for-a-zone-create-dns-record
  POST /zones/:zone_identifier/dns_records
*/
func (api *API) CreateDNSRecord(zone string, rr NewDNSRecord) error {
	z, err := api.ListZones(zone)
	if err != nil {
		return err
	}
	// TODO(jamesog): This is brittle, fix it
	zid := z[0].ID
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
func (api *API) DNSRecords(zone string) ([]DNSRecord, error) {
	z, err := api.ListZones(zone)
	if err != nil {
		return []DNSRecord{}, err
	}
	// TODO(jamesog): This is brittle, fix it
	zid := z[0].ID
	uri := "/zones/" + zid + "/dns_records"
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return []DNSRecord{}, err
	}
	var r DNSRecordResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return []DNSRecord{}, err
	}
	return r.Result, nil
}

// https://api.cloudflare.com/#dns-records-for-a-zone-dns-record-details
// GET /zones/:zone_identifier/dns_records/:identifier
func (api *API) DNSRecord() {
}

/*
Change a DNS record.

API reference:
  https://api.cloudflare.com/#dns-records-for-a-zone-update-dns-record
  PUT /zones/:zone_identifier/dns_records/:identifier
*/
func (api *API) UpdateDNSRecord() {
}

/*
Delete a DNS record.

API reference:
  https://api.cloudflare.com/#dns-records-for-a-zone-delete-dns-record
  DELETE /zones/:zone_identifier/dns_records/:identifier
*/
func (api *API) DeleteDNSRecord() {
}
