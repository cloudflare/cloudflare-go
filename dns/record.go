// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// RecordService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewRecordService] method instead.
type RecordService struct {
	Options []option.RequestOption
}

// NewRecordService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRecordService(opts ...option.RequestOption) (r *RecordService) {
	r = &RecordService{}
	r.Options = opts
	return
}

// Create a new DNS record for a zone.
//
// Notes:
//
//   - A/AAAA records cannot exist on the same name as CNAME records.
//   - NS records cannot exist on the same name as any other record type.
//   - Domain names are always represented in Punycode, even if Unicode characters
//     were used when creating the record.
func (r *RecordService) New(ctx context.Context, params RecordNewParams, opts ...option.RequestOption) (res *Record, err error) {
	opts = append(r.Options[:], opts...)
	var env RecordNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_records", params.getZoneID())
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Overwrite an existing DNS record. Notes:
//
//   - A/AAAA records cannot exist on the same name as CNAME records.
//   - NS records cannot exist on the same name as any other record type.
//   - Domain names are always represented in Punycode, even if Unicode characters
//     were used when creating the record.
func (r *RecordService) Update(ctx context.Context, dnsRecordID string, params RecordUpdateParams, opts ...option.RequestOption) (res *Record, err error) {
	opts = append(r.Options[:], opts...)
	var env RecordUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_records/%s", params.getZoneID(), dnsRecordID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List, search, sort, and filter a zones' DNS records.
func (r *RecordService) List(ctx context.Context, params RecordListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[Record], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/dns_records", params.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List, search, sort, and filter a zones' DNS records.
func (r *RecordService) ListAutoPaging(ctx context.Context, params RecordListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[Record] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete DNS Record
func (r *RecordService) Delete(ctx context.Context, dnsRecordID string, params RecordDeleteParams, opts ...option.RequestOption) (res *RecordDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RecordDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_records/%s", params.ZoneID, dnsRecordID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update an existing DNS record. Notes:
//
//   - A/AAAA records cannot exist on the same name as CNAME records.
//   - NS records cannot exist on the same name as any other record type.
//   - Domain names are always represented in Punycode, even if Unicode characters
//     were used when creating the record.
func (r *RecordService) Edit(ctx context.Context, dnsRecordID string, params RecordEditParams, opts ...option.RequestOption) (res *Record, err error) {
	opts = append(r.Options[:], opts...)
	var env RecordEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_records/%s", params.getZoneID(), dnsRecordID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// You can export your
// [BIND config](https://en.wikipedia.org/wiki/Zone_file "Zone file") through this
// endpoint.
//
// See
// [the documentation](https://developers.cloudflare.com/dns/manage-dns-records/how-to/import-and-export/ "Import and export records")
// for more information.
func (r *RecordService) Export(ctx context.Context, query RecordExportParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := fmt.Sprintf("zones/%s/dns_records/export", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// DNS Record Details
func (r *RecordService) Get(ctx context.Context, dnsRecordID string, query RecordGetParams, opts ...option.RequestOption) (res *Record, err error) {
	opts = append(r.Options[:], opts...)
	var env RecordGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_records/%s", query.ZoneID, dnsRecordID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// You can upload your
// [BIND config](https://en.wikipedia.org/wiki/Zone_file "Zone file") through this
// endpoint. It assumes that cURL is called from a location with bind_config.txt
// (valid BIND config) present.
//
// See
// [the documentation](https://developers.cloudflare.com/dns/manage-dns-records/how-to/import-and-export/ "Import and export records")
// for more information.
func (r *RecordService) Import(ctx context.Context, params RecordImportParams, opts ...option.RequestOption) (res *RecordImportResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RecordImportResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_records/import", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Scan for common DNS records on your domain and automatically add them to your
// zone. Useful if you haven't updated your nameservers yet.
func (r *RecordService) Scan(ctx context.Context, params RecordScanParams, opts ...option.RequestOption) (res *RecordScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RecordScanResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_records/scan", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ARecord struct {
	// A valid IPv4 address.
	Content string `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type ARecordType `json:"type,required"`
	// Identifier
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Whether this record can be modified/deleted (true means it's managed by
	// Cloudflare).
	Locked bool `json:"locked"`
	// Extra Cloudflare-specific information about the record.
	Meta UnnamedSchemaRef1391721bdb8938d5420a8bc738a44deb `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string      `json:"zone_name" format:"hostname"`
	JSON     aRecordJSON `json:"-"`
}

// aRecordJSON contains the JSON metadata for the struct [ARecord]
type aRecordJSON struct {
	Content     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Proxied     apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aRecordJSON) RawJSON() string {
	return r.raw
}

func (r ARecord) implementsDNSRecord() {}

// Record type.
type ARecordType string

const (
	ARecordTypeA ARecordType = "A"
)

func (r ARecordType) IsKnown() bool {
	switch r {
	case ARecordTypeA:
		return true
	}
	return false
}

type AAAARecord struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type AAAARecordType `json:"type,required"`
	// Identifier
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Whether this record can be modified/deleted (true means it's managed by
	// Cloudflare).
	Locked bool `json:"locked"`
	// Extra Cloudflare-specific information about the record.
	Meta UnnamedSchemaRef1391721bdb8938d5420a8bc738a44deb `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string         `json:"zone_name" format:"hostname"`
	JSON     aaaaRecordJSON `json:"-"`
}

// aaaaRecordJSON contains the JSON metadata for the struct [AAAARecord]
type aaaaRecordJSON struct {
	Content     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Proxied     apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AAAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aaaaRecordJSON) RawJSON() string {
	return r.raw
}

func (r AAAARecord) implementsDNSRecord() {}

// Record type.
type AAAARecordType string

const (
	AAAARecordTypeAAAA AAAARecordType = "AAAA"
)

func (r AAAARecordType) IsKnown() bool {
	switch r {
	case AAAARecordTypeAAAA:
		return true
	}
	return false
}

type CAARecord struct {
	// Components of a CAA record.
	Data CAARecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type CAARecordType `json:"type,required"`
	// Identifier
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// Formatted CAA content. See 'data' to set CAA properties.
	Content string `json:"content"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Whether this record can be modified/deleted (true means it's managed by
	// Cloudflare).
	Locked bool `json:"locked"`
	// Extra Cloudflare-specific information about the record.
	Meta UnnamedSchemaRef1391721bdb8938d5420a8bc738a44deb `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string        `json:"zone_name" format:"hostname"`
	JSON     caaRecordJSON `json:"-"`
}

// caaRecordJSON contains the JSON metadata for the struct [CAARecord]
type caaRecordJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r caaRecordJSON) RawJSON() string {
	return r.raw
}

func (r CAARecord) implementsDNSRecord() {}

// Components of a CAA record.
type CAARecordData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string            `json:"value"`
	JSON  caaRecordDataJSON `json:"-"`
}

// caaRecordDataJSON contains the JSON metadata for the struct [CAARecordData]
type caaRecordDataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CAARecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r caaRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type CAARecordType string

const (
	CAARecordTypeCAA CAARecordType = "CAA"
)

func (r CAARecordType) IsKnown() bool {
	switch r {
	case CAARecordTypeCAA:
		return true
	}
	return false
}

type CERTRecord struct {
	// Components of a CERT record.
	Data CERTRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type CERTRecordType `json:"type,required"`
	// Identifier
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// Formatted CERT content. See 'data' to set CERT properties.
	Content string `json:"content"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Whether this record can be modified/deleted (true means it's managed by
	// Cloudflare).
	Locked bool `json:"locked"`
	// Extra Cloudflare-specific information about the record.
	Meta UnnamedSchemaRef1391721bdb8938d5420a8bc738a44deb `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string         `json:"zone_name" format:"hostname"`
	JSON     certRecordJSON `json:"-"`
}

// certRecordJSON contains the JSON metadata for the struct [CERTRecord]
type certRecordJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CERTRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certRecordJSON) RawJSON() string {
	return r.raw
}

func (r CERTRecord) implementsDNSRecord() {}

// Components of a CERT record.
type CERTRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64            `json:"type"`
	JSON certRecordDataJSON `json:"-"`
}

// certRecordDataJSON contains the JSON metadata for the struct [CERTRecordData]
type certRecordDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CERTRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type CERTRecordType string

const (
	CERTRecordTypeCERT CERTRecordType = "CERT"
)

func (r CERTRecordType) IsKnown() bool {
	switch r {
	case CERTRecordTypeCERT:
		return true
	}
	return false
}

type CNAMERecord struct {
	// A valid hostname. Must not match the record's name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type CNAMERecordType `json:"type,required"`
	// Identifier
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Whether this record can be modified/deleted (true means it's managed by
	// Cloudflare).
	Locked bool `json:"locked"`
	// Extra Cloudflare-specific information about the record.
	Meta UnnamedSchemaRef1391721bdb8938d5420a8bc738a44deb `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string          `json:"zone_name" format:"hostname"`
	JSON     cnameRecordJSON `json:"-"`
}

// cnameRecordJSON contains the JSON metadata for the struct [CNAMERecord]
type cnameRecordJSON struct {
	Content     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Proxied     apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CNAMERecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cnameRecordJSON) RawJSON() string {
	return r.raw
}

func (r CNAMERecord) implementsDNSRecord() {}

// Record type.
type CNAMERecordType string

const (
	CNAMERecordTypeCNAME CNAMERecordType = "CNAME"
)

func (r CNAMERecordType) IsKnown() bool {
	switch r {
	case CNAMERecordTypeCNAME:
		return true
	}
	return false
}

type DNSKEYRecord struct {
	// Components of a DNSKEY record.
	Data DNSKEYRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSKEYRecordType `json:"type,required"`
	// Identifier
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// Formatted DNSKEY content. See 'data' to set DNSKEY properties.
	Content string `json:"content"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Whether this record can be modified/deleted (true means it's managed by
	// Cloudflare).
	Locked bool `json:"locked"`
	// Extra Cloudflare-specific information about the record.
	Meta UnnamedSchemaRef1391721bdb8938d5420a8bc738a44deb `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string           `json:"zone_name" format:"hostname"`
	JSON     dnskeyRecordJSON `json:"-"`
}

// dnskeyRecordJSON contains the JSON metadata for the struct [DNSKEYRecord]
type dnskeyRecordJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSKEYRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnskeyRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSKEYRecord) implementsDNSRecord() {}

// Components of a DNSKEY record.
type DNSKEYRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string               `json:"public_key"`
	JSON      dnskeyRecordDataJSON `json:"-"`
}

// dnskeyRecordDataJSON contains the JSON metadata for the struct
// [DNSKEYRecordData]
type dnskeyRecordDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSKEYRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnskeyRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSKEYRecordType string

const (
	DNSKEYRecordTypeDNSKEY DNSKEYRecordType = "DNSKEY"
)

func (r DNSKEYRecordType) IsKnown() bool {
	switch r {
	case DNSKEYRecordTypeDNSKEY:
		return true
	}
	return false
}

type DSRecord struct {
	// Components of a DS record.
	Data DSRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DSRecordType `json:"type,required"`
	// Identifier
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// Formatted DS content. See 'data' to set DS properties.
	Content string `json:"content"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Whether this record can be modified/deleted (true means it's managed by
	// Cloudflare).
	Locked bool `json:"locked"`
	// Extra Cloudflare-specific information about the record.
	Meta UnnamedSchemaRef1391721bdb8938d5420a8bc738a44deb `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string       `json:"zone_name" format:"hostname"`
	JSON     dsRecordJSON `json:"-"`
}

// dsRecordJSON contains the JSON metadata for the struct [DSRecord]
type dsRecordJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dsRecordJSON) RawJSON() string {
	return r.raw
}

func (r DSRecord) implementsDNSRecord() {}

// Components of a DS record.
type DSRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64          `json:"key_tag"`
	JSON   dsRecordDataJSON `json:"-"`
}

// dsRecordDataJSON contains the JSON metadata for the struct [DSRecordData]
type dsRecordDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DSRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dsRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DSRecordType string

const (
	DSRecordTypeDS DSRecordType = "DS"
)

func (r DSRecordType) IsKnown() bool {
	switch r {
	case DSRecordTypeDS:
		return true
	}
	return false
}

type HTTPSRecord struct {
	// Components of a HTTPS record.
	Data HTTPSRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type HTTPSRecordType `json:"type,required"`
	// Identifier
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// Formatted HTTPS content. See 'data' to set HTTPS properties.
	Content string `json:"content"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Whether this record can be modified/deleted (true means it's managed by
	// Cloudflare).
	Locked bool `json:"locked"`
	// Extra Cloudflare-specific information about the record.
	Meta UnnamedSchemaRef1391721bdb8938d5420a8bc738a44deb `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string          `json:"zone_name" format:"hostname"`
	JSON     httpsRecordJSON `json:"-"`
}

// httpsRecordJSON contains the JSON metadata for the struct [HTTPSRecord]
type httpsRecordJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpsRecordJSON) RawJSON() string {
	return r.raw
}

func (r HTTPSRecord) implementsDNSRecord() {}

// Components of a HTTPS record.
type HTTPSRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string              `json:"value"`
	JSON  httpsRecordDataJSON `json:"-"`
}

// httpsRecordDataJSON contains the JSON metadata for the struct [HTTPSRecordData]
type httpsRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPSRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpsRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type HTTPSRecordType string

const (
	HTTPSRecordTypeHTTPS HTTPSRecordType = "HTTPS"
)

func (r HTTPSRecordType) IsKnown() bool {
	switch r {
	case HTTPSRecordTypeHTTPS:
		return true
	}
	return false
}

type LOCRecord struct {
	// Components of a LOC record.
	Data LOCRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type LOCRecordType `json:"type,required"`
	// Identifier
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// Formatted LOC content. See 'data' to set LOC properties.
	Content string `json:"content"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Whether this record can be modified/deleted (true means it's managed by
	// Cloudflare).
	Locked bool `json:"locked"`
	// Extra Cloudflare-specific information about the record.
	Meta UnnamedSchemaRef1391721bdb8938d5420a8bc738a44deb `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string        `json:"zone_name" format:"hostname"`
	JSON     locRecordJSON `json:"-"`
}

// locRecordJSON contains the JSON metadata for the struct [LOCRecord]
type locRecordJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LOCRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r locRecordJSON) RawJSON() string {
	return r.raw
}

func (r LOCRecord) implementsDNSRecord() {}

// Components of a LOC record.
type LOCRecordData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection LOCRecordDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection LOCRecordDataLongDirection `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes float64 `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds float64 `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz float64 `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert float64 `json:"precision_vert"`
	// Size of location in meters.
	Size float64           `json:"size"`
	JSON locRecordDataJSON `json:"-"`
}

// locRecordDataJSON contains the JSON metadata for the struct [LOCRecordData]
type locRecordDataJSON struct {
	Altitude      apijson.Field
	LatDegrees    apijson.Field
	LatDirection  apijson.Field
	LatMinutes    apijson.Field
	LatSeconds    apijson.Field
	LongDegrees   apijson.Field
	LongDirection apijson.Field
	LongMinutes   apijson.Field
	LongSeconds   apijson.Field
	PrecisionHorz apijson.Field
	PrecisionVert apijson.Field
	Size          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LOCRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r locRecordDataJSON) RawJSON() string {
	return r.raw
}

// Latitude direction.
type LOCRecordDataLatDirection string

const (
	LOCRecordDataLatDirectionN LOCRecordDataLatDirection = "N"
	LOCRecordDataLatDirectionS LOCRecordDataLatDirection = "S"
)

func (r LOCRecordDataLatDirection) IsKnown() bool {
	switch r {
	case LOCRecordDataLatDirectionN, LOCRecordDataLatDirectionS:
		return true
	}
	return false
}

// Longitude direction.
type LOCRecordDataLongDirection string

const (
	LOCRecordDataLongDirectionE LOCRecordDataLongDirection = "E"
	LOCRecordDataLongDirectionW LOCRecordDataLongDirection = "W"
)

func (r LOCRecordDataLongDirection) IsKnown() bool {
	switch r {
	case LOCRecordDataLongDirectionE, LOCRecordDataLongDirectionW:
		return true
	}
	return false
}

// Record type.
type LOCRecordType string

const (
	LOCRecordTypeLOC LOCRecordType = "LOC"
)

func (r LOCRecordType) IsKnown() bool {
	switch r {
	case LOCRecordTypeLOC:
		return true
	}
	return false
}

type MXRecord struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type MXRecordType `json:"type,required"`
	// Identifier
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Whether this record can be modified/deleted (true means it's managed by
	// Cloudflare).
	Locked bool `json:"locked"`
	// Extra Cloudflare-specific information about the record.
	Meta UnnamedSchemaRef1391721bdb8938d5420a8bc738a44deb `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string       `json:"zone_name" format:"hostname"`
	JSON     mxRecordJSON `json:"-"`
}

// mxRecordJSON contains the JSON metadata for the struct [MXRecord]
type mxRecordJSON struct {
	Content     apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MXRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mxRecordJSON) RawJSON() string {
	return r.raw
}

func (r MXRecord) implementsDNSRecord() {}

// Record type.
type MXRecordType string

const (
	MXRecordTypeMX MXRecordType = "MX"
)

func (r MXRecordType) IsKnown() bool {
	switch r {
	case MXRecordTypeMX:
		return true
	}
	return false
}

type NAPTRRecord struct {
	// Components of a NAPTR record.
	Data NAPTRRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type NAPTRRecordType `json:"type,required"`
	// Identifier
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// Formatted NAPTR content. See 'data' to set NAPTR properties.
	Content string `json:"content"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Whether this record can be modified/deleted (true means it's managed by
	// Cloudflare).
	Locked bool `json:"locked"`
	// Extra Cloudflare-specific information about the record.
	Meta UnnamedSchemaRef1391721bdb8938d5420a8bc738a44deb `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string          `json:"zone_name" format:"hostname"`
	JSON     naptrRecordJSON `json:"-"`
}

// naptrRecordJSON contains the JSON metadata for the struct [NAPTRRecord]
type naptrRecordJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NAPTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r naptrRecordJSON) RawJSON() string {
	return r.raw
}

func (r NAPTRRecord) implementsDNSRecord() {}

// Components of a NAPTR record.
type NAPTRRecordData struct {
	// Flags.
	Flags string `json:"flags"`
	// Order.
	Order float64 `json:"order"`
	// Preference.
	Preference float64 `json:"preference"`
	// Regex.
	Regex string `json:"regex"`
	// Replacement.
	Replacement string `json:"replacement"`
	// Service.
	Service string              `json:"service"`
	JSON    naptrRecordDataJSON `json:"-"`
}

// naptrRecordDataJSON contains the JSON metadata for the struct [NAPTRRecordData]
type naptrRecordDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NAPTRRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r naptrRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type NAPTRRecordType string

const (
	NAPTRRecordTypeNAPTR NAPTRRecordType = "NAPTR"
)

func (r NAPTRRecordType) IsKnown() bool {
	switch r {
	case NAPTRRecordTypeNAPTR:
		return true
	}
	return false
}

type NSRecord struct {
	// A valid name server host name.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type NSRecordType `json:"type,required"`
	// Identifier
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Whether this record can be modified/deleted (true means it's managed by
	// Cloudflare).
	Locked bool `json:"locked"`
	// Extra Cloudflare-specific information about the record.
	Meta UnnamedSchemaRef1391721bdb8938d5420a8bc738a44deb `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string       `json:"zone_name" format:"hostname"`
	JSON     nsRecordJSON `json:"-"`
}

// nsRecordJSON contains the JSON metadata for the struct [NSRecord]
type nsRecordJSON struct {
	Content     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nsRecordJSON) RawJSON() string {
	return r.raw
}

func (r NSRecord) implementsDNSRecord() {}

// Record type.
type NSRecordType string

const (
	NSRecordTypeNS NSRecordType = "NS"
)

func (r NSRecordType) IsKnown() bool {
	switch r {
	case NSRecordTypeNS:
		return true
	}
	return false
}

type PTRRecord struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type PTRRecordType `json:"type,required"`
	// Identifier
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Whether this record can be modified/deleted (true means it's managed by
	// Cloudflare).
	Locked bool `json:"locked"`
	// Extra Cloudflare-specific information about the record.
	Meta UnnamedSchemaRef1391721bdb8938d5420a8bc738a44deb `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string        `json:"zone_name" format:"hostname"`
	JSON     ptrRecordJSON `json:"-"`
}

// ptrRecordJSON contains the JSON metadata for the struct [PTRRecord]
type ptrRecordJSON struct {
	Content     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ptrRecordJSON) RawJSON() string {
	return r.raw
}

func (r PTRRecord) implementsDNSRecord() {}

// Record type.
type PTRRecordType string

const (
	PTRRecordTypePTR PTRRecordType = "PTR"
)

func (r PTRRecordType) IsKnown() bool {
	switch r {
	case PTRRecordTypePTR:
		return true
	}
	return false
}

type Record struct {
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Record type.
	Type RecordType `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Identifier
	ID string `json:"id"`
	// Whether this record can be modified/deleted (true means it's managed by
	// Cloudflare).
	Locked bool `json:"locked"`
	// Extra Cloudflare-specific information about the record.
	Meta UnnamedSchemaRef1391721bdb8938d5420a8bc738a44deb `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool        `json:"proxiable"`
	Tags      interface{} `json:"tags,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string      `json:"zone_name" format:"hostname"`
	Data     interface{} `json:"data,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64    `json:"priority"`
	JSON     recordJSON `json:"-"`
	union    RecordUnion
}

// recordJSON contains the JSON metadata for the struct [Record]
type recordJSON struct {
	Content     apijson.Field
	Name        apijson.Field
	Proxied     apijson.Field
	Type        apijson.Field
	Comment     apijson.Field
	CreatedOn   apijson.Field
	ID          apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	Data        apijson.Field
	Priority    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r recordJSON) RawJSON() string {
	return r.raw
}

func (r *Record) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r Record) AsUnion() RecordUnion {
	return r.union
}

// Union satisfied by [dns.ARecord], [dns.AAAARecord], [dns.CAARecord],
// [dns.CERTRecord], [dns.CNAMERecord], [dns.DNSKEYRecord], [dns.DSRecord],
// [dns.HTTPSRecord], [dns.LOCRecord], [dns.MXRecord], [dns.NAPTRRecord],
// [dns.NSRecord], [dns.PTRRecord], [dns.SMIMEARecord], [dns.SRVRecord],
// [dns.SSHFPRecord], [dns.SVCBRecord], [dns.TLSARecord], [dns.TXTRecord] or
// [dns.URIRecord].
type RecordUnion interface {
	implementsDNSRecord()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ARecord{}),
			DiscriminatorValue: "A",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AAAARecord{}),
			DiscriminatorValue: "AAAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CAARecord{}),
			DiscriminatorValue: "CAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CERTRecord{}),
			DiscriminatorValue: "CERT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CNAMERecord{}),
			DiscriminatorValue: "CNAME",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSKEYRecord{}),
			DiscriminatorValue: "DNSKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DSRecord{}),
			DiscriminatorValue: "DS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(HTTPSRecord{}),
			DiscriminatorValue: "HTTPS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LOCRecord{}),
			DiscriminatorValue: "LOC",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MXRecord{}),
			DiscriminatorValue: "MX",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(NAPTRRecord{}),
			DiscriminatorValue: "NAPTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(NSRecord{}),
			DiscriminatorValue: "NS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PTRRecord{}),
			DiscriminatorValue: "PTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SMIMEARecord{}),
			DiscriminatorValue: "SMIMEA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SRVRecord{}),
			DiscriminatorValue: "SRV",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SSHFPRecord{}),
			DiscriminatorValue: "SSHFP",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SVCBRecord{}),
			DiscriminatorValue: "SVCB",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TLSARecord{}),
			DiscriminatorValue: "TLSA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TXTRecord{}),
			DiscriminatorValue: "TXT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(URIRecord{}),
			DiscriminatorValue: "URI",
		},
	)
}

// Record type.
type RecordType string

const (
	RecordTypeA      RecordType = "A"
	RecordTypeAAAA   RecordType = "AAAA"
	RecordTypeCAA    RecordType = "CAA"
	RecordTypeCERT   RecordType = "CERT"
	RecordTypeCNAME  RecordType = "CNAME"
	RecordTypeDNSKEY RecordType = "DNSKEY"
	RecordTypeDS     RecordType = "DS"
	RecordTypeHTTPS  RecordType = "HTTPS"
	RecordTypeLOC    RecordType = "LOC"
	RecordTypeMX     RecordType = "MX"
	RecordTypeNAPTR  RecordType = "NAPTR"
	RecordTypeNS     RecordType = "NS"
	RecordTypePTR    RecordType = "PTR"
	RecordTypeSMIMEA RecordType = "SMIMEA"
	RecordTypeSRV    RecordType = "SRV"
	RecordTypeSSHFP  RecordType = "SSHFP"
	RecordTypeSVCB   RecordType = "SVCB"
	RecordTypeTLSA   RecordType = "TLSA"
	RecordTypeTXT    RecordType = "TXT"
	RecordTypeURI    RecordType = "URI"
)

func (r RecordType) IsKnown() bool {
	switch r {
	case RecordTypeA, RecordTypeAAAA, RecordTypeCAA, RecordTypeCERT, RecordTypeCNAME, RecordTypeDNSKEY, RecordTypeDS, RecordTypeHTTPS, RecordTypeLOC, RecordTypeMX, RecordTypeNAPTR, RecordTypeNS, RecordTypePTR, RecordTypeSMIMEA, RecordTypeSRV, RecordTypeSSHFP, RecordTypeSVCB, RecordTypeTLSA, RecordTypeTXT, RecordTypeURI:
		return true
	}
	return false
}

type RecordTagsItem = string

type RecordTagsItemParam = string

type SMIMEARecord struct {
	// Components of a SMIMEA record.
	Data SMIMEARecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type SMIMEARecordType `json:"type,required"`
	// Identifier
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// Formatted SMIMEA content. See 'data' to set SMIMEA properties.
	Content string `json:"content"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Whether this record can be modified/deleted (true means it's managed by
	// Cloudflare).
	Locked bool `json:"locked"`
	// Extra Cloudflare-specific information about the record.
	Meta UnnamedSchemaRef1391721bdb8938d5420a8bc738a44deb `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string           `json:"zone_name" format:"hostname"`
	JSON     smimeaRecordJSON `json:"-"`
}

// smimeaRecordJSON contains the JSON metadata for the struct [SMIMEARecord]
type smimeaRecordJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SMIMEARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smimeaRecordJSON) RawJSON() string {
	return r.raw
}

func (r SMIMEARecord) implementsDNSRecord() {}

// Components of a SMIMEA record.
type SMIMEARecordData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64              `json:"usage"`
	JSON  smimeaRecordDataJSON `json:"-"`
}

// smimeaRecordDataJSON contains the JSON metadata for the struct
// [SMIMEARecordData]
type smimeaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SMIMEARecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smimeaRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type SMIMEARecordType string

const (
	SMIMEARecordTypeSMIMEA SMIMEARecordType = "SMIMEA"
)

func (r SMIMEARecordType) IsKnown() bool {
	switch r {
	case SMIMEARecordTypeSMIMEA:
		return true
	}
	return false
}

type SRVRecord struct {
	// Components of a SRV record.
	Data SRVRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name string `json:"name,required"`
	// Record type.
	Type SRVRecordType `json:"type,required"`
	// Identifier
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// Priority, weight, port, and SRV target. See 'data' for setting the individual
	// component values.
	Content string `json:"content"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Whether this record can be modified/deleted (true means it's managed by
	// Cloudflare).
	Locked bool `json:"locked"`
	// Extra Cloudflare-specific information about the record.
	Meta UnnamedSchemaRef1391721bdb8938d5420a8bc738a44deb `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string        `json:"zone_name" format:"hostname"`
	JSON     srvRecordJSON `json:"-"`
}

// srvRecordJSON contains the JSON metadata for the struct [SRVRecord]
type srvRecordJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SRVRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r srvRecordJSON) RawJSON() string {
	return r.raw
}

func (r SRVRecord) implementsDNSRecord() {}

// Components of a SRV record.
type SRVRecordData struct {
	// A valid hostname. Deprecated in favor of the regular 'name' outside the data
	// map. This data map field represents the remainder of the full 'name' after the
	// service and protocol.
	Name string `json:"name" format:"hostname"`
	// The port of the service.
	Port float64 `json:"port"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// A valid protocol, prefixed with an underscore. Deprecated in favor of the
	// regular 'name' outside the data map. This data map field normally represents the
	// second label of that 'name'.
	Proto string `json:"proto"`
	// A service type, prefixed with an underscore. Deprecated in favor of the regular
	// 'name' outside the data map. This data map field normally represents the first
	// label of that 'name'.
	Service string `json:"service"`
	// A valid hostname.
	Target string `json:"target" format:"hostname"`
	// The record weight.
	Weight float64           `json:"weight"`
	JSON   srvRecordDataJSON `json:"-"`
}

// srvRecordDataJSON contains the JSON metadata for the struct [SRVRecordData]
type srvRecordDataJSON struct {
	Name        apijson.Field
	Port        apijson.Field
	Priority    apijson.Field
	Proto       apijson.Field
	Service     apijson.Field
	Target      apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SRVRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r srvRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type SRVRecordType string

const (
	SRVRecordTypeSRV SRVRecordType = "SRV"
)

func (r SRVRecordType) IsKnown() bool {
	switch r {
	case SRVRecordTypeSRV:
		return true
	}
	return false
}

type SSHFPRecord struct {
	// Components of a SSHFP record.
	Data SSHFPRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type SSHFPRecordType `json:"type,required"`
	// Identifier
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// Formatted SSHFP content. See 'data' to set SSHFP properties.
	Content string `json:"content"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Whether this record can be modified/deleted (true means it's managed by
	// Cloudflare).
	Locked bool `json:"locked"`
	// Extra Cloudflare-specific information about the record.
	Meta UnnamedSchemaRef1391721bdb8938d5420a8bc738a44deb `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string          `json:"zone_name" format:"hostname"`
	JSON     sshfpRecordJSON `json:"-"`
}

// sshfpRecordJSON contains the JSON metadata for the struct [SSHFPRecord]
type sshfpRecordJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSHFPRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sshfpRecordJSON) RawJSON() string {
	return r.raw
}

func (r SSHFPRecord) implementsDNSRecord() {}

// Components of a SSHFP record.
type SSHFPRecordData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64             `json:"type"`
	JSON sshfpRecordDataJSON `json:"-"`
}

// sshfpRecordDataJSON contains the JSON metadata for the struct [SSHFPRecordData]
type sshfpRecordDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSHFPRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sshfpRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type SSHFPRecordType string

const (
	SSHFPRecordTypeSSHFP SSHFPRecordType = "SSHFP"
)

func (r SSHFPRecordType) IsKnown() bool {
	switch r {
	case SSHFPRecordTypeSSHFP:
		return true
	}
	return false
}

type SVCBRecord struct {
	// Components of a SVCB record.
	Data SVCBRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type SVCBRecordType `json:"type,required"`
	// Identifier
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// Formatted SVCB content. See 'data' to set SVCB properties.
	Content string `json:"content"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Whether this record can be modified/deleted (true means it's managed by
	// Cloudflare).
	Locked bool `json:"locked"`
	// Extra Cloudflare-specific information about the record.
	Meta UnnamedSchemaRef1391721bdb8938d5420a8bc738a44deb `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string         `json:"zone_name" format:"hostname"`
	JSON     svcbRecordJSON `json:"-"`
}

// svcbRecordJSON contains the JSON metadata for the struct [SVCBRecord]
type svcbRecordJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SVCBRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r svcbRecordJSON) RawJSON() string {
	return r.raw
}

func (r SVCBRecord) implementsDNSRecord() {}

// Components of a SVCB record.
type SVCBRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string             `json:"value"`
	JSON  svcbRecordDataJSON `json:"-"`
}

// svcbRecordDataJSON contains the JSON metadata for the struct [SVCBRecordData]
type svcbRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SVCBRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r svcbRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type SVCBRecordType string

const (
	SVCBRecordTypeSVCB SVCBRecordType = "SVCB"
)

func (r SVCBRecordType) IsKnown() bool {
	switch r {
	case SVCBRecordTypeSVCB:
		return true
	}
	return false
}

type TLSARecord struct {
	// Components of a TLSA record.
	Data TLSARecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type TLSARecordType `json:"type,required"`
	// Identifier
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// Formatted TLSA content. See 'data' to set TLSA properties.
	Content string `json:"content"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Whether this record can be modified/deleted (true means it's managed by
	// Cloudflare).
	Locked bool `json:"locked"`
	// Extra Cloudflare-specific information about the record.
	Meta UnnamedSchemaRef1391721bdb8938d5420a8bc738a44deb `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string         `json:"zone_name" format:"hostname"`
	JSON     tlsaRecordJSON `json:"-"`
}

// tlsaRecordJSON contains the JSON metadata for the struct [TLSARecord]
type tlsaRecordJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TLSARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tlsaRecordJSON) RawJSON() string {
	return r.raw
}

func (r TLSARecord) implementsDNSRecord() {}

// Components of a TLSA record.
type TLSARecordData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64            `json:"usage"`
	JSON  tlsaRecordDataJSON `json:"-"`
}

// tlsaRecordDataJSON contains the JSON metadata for the struct [TLSARecordData]
type tlsaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TLSARecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tlsaRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type TLSARecordType string

const (
	TLSARecordTypeTLSA TLSARecordType = "TLSA"
)

func (r TLSARecordType) IsKnown() bool {
	switch r {
	case TLSARecordTypeTLSA:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.TTLNumber].
type TTLUnion interface {
	ImplementsDNSTTLUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TTLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(TTLNumber(0)),
		},
	)
}

type TTLNumber float64

const (
	TTLNumber1 TTLNumber = 1
)

func (r TTLNumber) IsKnown() bool {
	switch r {
	case TTLNumber1:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat], [dns.TTLNumber].
type TTLUnionParam interface {
	ImplementsDNSTTLUnionParam()
}

type TXTRecord struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type TXTRecordType `json:"type,required"`
	// Identifier
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Whether this record can be modified/deleted (true means it's managed by
	// Cloudflare).
	Locked bool `json:"locked"`
	// Extra Cloudflare-specific information about the record.
	Meta UnnamedSchemaRef1391721bdb8938d5420a8bc738a44deb `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string        `json:"zone_name" format:"hostname"`
	JSON     txtRecordJSON `json:"-"`
}

// txtRecordJSON contains the JSON metadata for the struct [TXTRecord]
type txtRecordJSON struct {
	Content     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TXTRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r txtRecordJSON) RawJSON() string {
	return r.raw
}

func (r TXTRecord) implementsDNSRecord() {}

// Record type.
type TXTRecordType string

const (
	TXTRecordTypeTXT TXTRecordType = "TXT"
)

func (r TXTRecordType) IsKnown() bool {
	switch r {
	case TXTRecordTypeTXT:
		return true
	}
	return false
}

// Extra Cloudflare-specific information about the record.
type UnnamedSchemaRef1391721bdb8938d5420a8bc738a44deb struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                               `json:"source"`
	JSON   unnamedSchemaRef1391721bdb8938d5420a8bc738a44debJSON `json:"-"`
}

// unnamedSchemaRef1391721bdb8938d5420a8bc738a44debJSON contains the JSON metadata
// for the struct [UnnamedSchemaRef1391721bdb8938d5420a8bc738a44deb]
type unnamedSchemaRef1391721bdb8938d5420a8bc738a44debJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef1391721bdb8938d5420a8bc738a44deb) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef1391721bdb8938d5420a8bc738a44debJSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef2f895e023ae55b55d2f5925449b819cd struct {
	// When the file parsing ended.
	EndTime time.Time `json:"end_time" format:"date-time"`
	// Processing time of the file in seconds.
	ProcessTime float64 `json:"process_time"`
	// When the file parsing started.
	StartTime time.Time                                            `json:"start_time" format:"date-time"`
	JSON      unnamedSchemaRef2f895e023ae55b55d2f5925449b819cdJSON `json:"-"`
}

// unnamedSchemaRef2f895e023ae55b55d2f5925449b819cdJSON contains the JSON metadata
// for the struct [UnnamedSchemaRef2f895e023ae55b55d2f5925449b819cd]
type unnamedSchemaRef2f895e023ae55b55d2f5925449b819cdJSON struct {
	EndTime     apijson.Field
	ProcessTime apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef2f895e023ae55b55d2f5925449b819cd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef2f895e023ae55b55d2f5925449b819cdJSON) RawJSON() string {
	return r.raw
}

type URIRecord struct {
	// Components of a URI record.
	Data URIRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type URIRecordType `json:"type,required"`
	// Identifier
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// Formatted URI content. See 'data' to set URI properties.
	Content string `json:"content"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Whether this record can be modified/deleted (true means it's managed by
	// Cloudflare).
	Locked bool `json:"locked"`
	// Extra Cloudflare-specific information about the record.
	Meta UnnamedSchemaRef1391721bdb8938d5420a8bc738a44deb `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string        `json:"zone_name" format:"hostname"`
	JSON     uriRecordJSON `json:"-"`
}

// uriRecordJSON contains the JSON metadata for the struct [URIRecord]
type uriRecordJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URIRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uriRecordJSON) RawJSON() string {
	return r.raw
}

func (r URIRecord) implementsDNSRecord() {}

// Components of a URI record.
type URIRecordData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64           `json:"weight"`
	JSON   uriRecordDataJSON `json:"-"`
}

// uriRecordDataJSON contains the JSON metadata for the struct [URIRecordData]
type uriRecordDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URIRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uriRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type URIRecordType string

const (
	URIRecordTypeURI URIRecordType = "URI"
)

func (r URIRecordType) IsKnown() bool {
	switch r {
	case URIRecordTypeURI:
		return true
	}
	return false
}

type RecordDeleteResponse struct {
	// Identifier
	ID   string                   `json:"id"`
	JSON recordDeleteResponseJSON `json:"-"`
}

// recordDeleteResponseJSON contains the JSON metadata for the struct
// [RecordDeleteResponse]
type recordDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type RecordImportResponse struct {
	// Number of DNS records added.
	RecsAdded float64 `json:"recs_added"`
	// Total number of DNS records parsed.
	TotalRecordsParsed float64                  `json:"total_records_parsed"`
	JSON               recordImportResponseJSON `json:"-"`
}

// recordImportResponseJSON contains the JSON metadata for the struct
// [RecordImportResponse]
type recordImportResponseJSON struct {
	RecsAdded          apijson.Field
	TotalRecordsParsed apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RecordImportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordImportResponseJSON) RawJSON() string {
	return r.raw
}

type RecordScanResponse struct {
	// Number of DNS records added.
	RecsAdded float64 `json:"recs_added"`
	// Total number of DNS records parsed.
	TotalRecordsParsed float64                `json:"total_records_parsed"`
	JSON               recordScanResponseJSON `json:"-"`
}

// recordScanResponseJSON contains the JSON metadata for the struct
// [RecordScanResponse]
type recordScanResponseJSON struct {
	RecsAdded          apijson.Field
	TotalRecordsParsed apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RecordScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordScanResponseJSON) RawJSON() string {
	return r.raw
}

// This interface is a union satisfied by one of the following:
// [RecordNewParamsARecord], [RecordNewParamsAAAARecord],
// [RecordNewParamsCAARecord], [RecordNewParamsCERTRecord],
// [RecordNewParamsCNAMERecord], [RecordNewParamsDNSKEYRecord],
// [RecordNewParamsDSRecord], [RecordNewParamsHTTPSRecord],
// [RecordNewParamsLOCRecord], [RecordNewParamsMXRecord],
// [RecordNewParamsNAPTRRecord], [RecordNewParamsNSRecord],
// [RecordNewParamsPTRRecord], [RecordNewParamsSMIMEARecord],
// [RecordNewParamsSRVRecord], [RecordNewParamsSSHFPRecord],
// [RecordNewParamsSVCBRecord], [RecordNewParamsTLSARecord],
// [RecordNewParamsTXTRecord], [RecordNewParamsURIRecord].
type RecordNewParams interface {
	ImplementsRecordNewParams()

	getZoneID() param.Field[string]
}

type RecordNewParamsARecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A valid IPv4 address.
	Content param.Field[string] `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordNewParamsARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsARecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsARecord) ImplementsRecordNewParams() {

}

// Record type.
type RecordNewParamsARecordType string

const (
	RecordNewParamsARecordTypeA RecordNewParamsARecordType = "A"
)

func (r RecordNewParamsARecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsARecordTypeA:
		return true
	}
	return false
}

type RecordNewParamsAAAARecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A valid IPv6 address.
	Content param.Field[string] `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsAAAARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordNewParamsAAAARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsAAAARecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsAAAARecord) ImplementsRecordNewParams() {

}

// Record type.
type RecordNewParamsAAAARecordType string

const (
	RecordNewParamsAAAARecordTypeAAAA RecordNewParamsAAAARecordType = "AAAA"
)

func (r RecordNewParamsAAAARecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsAAAARecordTypeAAAA:
		return true
	}
	return false
}

type RecordNewParamsCAARecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a CAA record.
	Data param.Field[RecordNewParamsCAARecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsCAARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordNewParamsCAARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsCAARecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsCAARecord) ImplementsRecordNewParams() {

}

// Components of a CAA record.
type RecordNewParamsCAARecordData struct {
	// Flags for the CAA record.
	Flags param.Field[float64] `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag param.Field[string] `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value param.Field[string] `json:"value"`
}

func (r RecordNewParamsCAARecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordNewParamsCAARecordType string

const (
	RecordNewParamsCAARecordTypeCAA RecordNewParamsCAARecordType = "CAA"
)

func (r RecordNewParamsCAARecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsCAARecordTypeCAA:
		return true
	}
	return false
}

type RecordNewParamsCERTRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a CERT record.
	Data param.Field[RecordNewParamsCERTRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsCERTRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordNewParamsCERTRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsCERTRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsCERTRecord) ImplementsRecordNewParams() {

}

// Components of a CERT record.
type RecordNewParamsCERTRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
	// Type.
	Type param.Field[float64] `json:"type"`
}

func (r RecordNewParamsCERTRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordNewParamsCERTRecordType string

const (
	RecordNewParamsCERTRecordTypeCERT RecordNewParamsCERTRecordType = "CERT"
)

func (r RecordNewParamsCERTRecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsCERTRecordTypeCERT:
		return true
	}
	return false
}

type RecordNewParamsCNAMERecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A valid hostname. Must not match the record's name.
	Content param.Field[interface{}] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsCNAMERecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordNewParamsCNAMERecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsCNAMERecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsCNAMERecord) ImplementsRecordNewParams() {

}

// Record type.
type RecordNewParamsCNAMERecordType string

const (
	RecordNewParamsCNAMERecordTypeCNAME RecordNewParamsCNAMERecordType = "CNAME"
)

func (r RecordNewParamsCNAMERecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsCNAMERecordTypeCNAME:
		return true
	}
	return false
}

type RecordNewParamsDNSKEYRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a DNSKEY record.
	Data param.Field[RecordNewParamsDNSKEYRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsDNSKEYRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordNewParamsDNSKEYRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsDNSKEYRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsDNSKEYRecord) ImplementsRecordNewParams() {

}

// Components of a DNSKEY record.
type RecordNewParamsDNSKEYRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Flags.
	Flags param.Field[float64] `json:"flags"`
	// Protocol.
	Protocol param.Field[float64] `json:"protocol"`
	// Public Key.
	PublicKey param.Field[string] `json:"public_key"`
}

func (r RecordNewParamsDNSKEYRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordNewParamsDNSKEYRecordType string

const (
	RecordNewParamsDNSKEYRecordTypeDNSKEY RecordNewParamsDNSKEYRecordType = "DNSKEY"
)

func (r RecordNewParamsDNSKEYRecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSKEYRecordTypeDNSKEY:
		return true
	}
	return false
}

type RecordNewParamsDSRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a DS record.
	Data param.Field[RecordNewParamsDSRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsDSRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordNewParamsDSRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsDSRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsDSRecord) ImplementsRecordNewParams() {

}

// Components of a DS record.
type RecordNewParamsDSRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Digest.
	Digest param.Field[string] `json:"digest"`
	// Digest Type.
	DigestType param.Field[float64] `json:"digest_type"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
}

func (r RecordNewParamsDSRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordNewParamsDSRecordType string

const (
	RecordNewParamsDSRecordTypeDS RecordNewParamsDSRecordType = "DS"
)

func (r RecordNewParamsDSRecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsDSRecordTypeDS:
		return true
	}
	return false
}

type RecordNewParamsHTTPSRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a HTTPS record.
	Data param.Field[RecordNewParamsHTTPSRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsHTTPSRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordNewParamsHTTPSRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsHTTPSRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsHTTPSRecord) ImplementsRecordNewParams() {

}

// Components of a HTTPS record.
type RecordNewParamsHTTPSRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r RecordNewParamsHTTPSRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordNewParamsHTTPSRecordType string

const (
	RecordNewParamsHTTPSRecordTypeHTTPS RecordNewParamsHTTPSRecordType = "HTTPS"
)

func (r RecordNewParamsHTTPSRecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsHTTPSRecordTypeHTTPS:
		return true
	}
	return false
}

type RecordNewParamsLOCRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a LOC record.
	Data param.Field[RecordNewParamsLOCRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsLOCRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordNewParamsLOCRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsLOCRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsLOCRecord) ImplementsRecordNewParams() {

}

// Components of a LOC record.
type RecordNewParamsLOCRecordData struct {
	// Altitude of location in meters.
	Altitude param.Field[float64] `json:"altitude"`
	// Degrees of latitude.
	LatDegrees param.Field[float64] `json:"lat_degrees"`
	// Latitude direction.
	LatDirection param.Field[RecordNewParamsLOCRecordDataLatDirection] `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes param.Field[float64] `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds param.Field[float64] `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees param.Field[float64] `json:"long_degrees"`
	// Longitude direction.
	LongDirection param.Field[RecordNewParamsLOCRecordDataLongDirection] `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes param.Field[float64] `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds param.Field[float64] `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz param.Field[float64] `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert param.Field[float64] `json:"precision_vert"`
	// Size of location in meters.
	Size param.Field[float64] `json:"size"`
}

func (r RecordNewParamsLOCRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Latitude direction.
type RecordNewParamsLOCRecordDataLatDirection string

const (
	RecordNewParamsLOCRecordDataLatDirectionN RecordNewParamsLOCRecordDataLatDirection = "N"
	RecordNewParamsLOCRecordDataLatDirectionS RecordNewParamsLOCRecordDataLatDirection = "S"
)

func (r RecordNewParamsLOCRecordDataLatDirection) IsKnown() bool {
	switch r {
	case RecordNewParamsLOCRecordDataLatDirectionN, RecordNewParamsLOCRecordDataLatDirectionS:
		return true
	}
	return false
}

// Longitude direction.
type RecordNewParamsLOCRecordDataLongDirection string

const (
	RecordNewParamsLOCRecordDataLongDirectionE RecordNewParamsLOCRecordDataLongDirection = "E"
	RecordNewParamsLOCRecordDataLongDirectionW RecordNewParamsLOCRecordDataLongDirection = "W"
)

func (r RecordNewParamsLOCRecordDataLongDirection) IsKnown() bool {
	switch r {
	case RecordNewParamsLOCRecordDataLongDirectionE, RecordNewParamsLOCRecordDataLongDirectionW:
		return true
	}
	return false
}

// Record type.
type RecordNewParamsLOCRecordType string

const (
	RecordNewParamsLOCRecordTypeLOC RecordNewParamsLOCRecordType = "LOC"
)

func (r RecordNewParamsLOCRecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsLOCRecordTypeLOC:
		return true
	}
	return false
}

type RecordNewParamsMXRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A valid mail server hostname.
	Content param.Field[string] `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority,required"`
	// Record type.
	Type param.Field[RecordNewParamsMXRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordNewParamsMXRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsMXRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsMXRecord) ImplementsRecordNewParams() {

}

// Record type.
type RecordNewParamsMXRecordType string

const (
	RecordNewParamsMXRecordTypeMX RecordNewParamsMXRecordType = "MX"
)

func (r RecordNewParamsMXRecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsMXRecordTypeMX:
		return true
	}
	return false
}

type RecordNewParamsNAPTRRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a NAPTR record.
	Data param.Field[RecordNewParamsNAPTRRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsNAPTRRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordNewParamsNAPTRRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsNAPTRRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsNAPTRRecord) ImplementsRecordNewParams() {

}

// Components of a NAPTR record.
type RecordNewParamsNAPTRRecordData struct {
	// Flags.
	Flags param.Field[string] `json:"flags"`
	// Order.
	Order param.Field[float64] `json:"order"`
	// Preference.
	Preference param.Field[float64] `json:"preference"`
	// Regex.
	Regex param.Field[string] `json:"regex"`
	// Replacement.
	Replacement param.Field[string] `json:"replacement"`
	// Service.
	Service param.Field[string] `json:"service"`
}

func (r RecordNewParamsNAPTRRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordNewParamsNAPTRRecordType string

const (
	RecordNewParamsNAPTRRecordTypeNAPTR RecordNewParamsNAPTRRecordType = "NAPTR"
)

func (r RecordNewParamsNAPTRRecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsNAPTRRecordTypeNAPTR:
		return true
	}
	return false
}

type RecordNewParamsNSRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A valid name server host name.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsNSRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordNewParamsNSRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsNSRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsNSRecord) ImplementsRecordNewParams() {

}

// Record type.
type RecordNewParamsNSRecordType string

const (
	RecordNewParamsNSRecordTypeNS RecordNewParamsNSRecordType = "NS"
)

func (r RecordNewParamsNSRecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsNSRecordTypeNS:
		return true
	}
	return false
}

type RecordNewParamsPTRRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Domain name pointing to the address.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsPTRRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordNewParamsPTRRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsPTRRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsPTRRecord) ImplementsRecordNewParams() {

}

// Record type.
type RecordNewParamsPTRRecordType string

const (
	RecordNewParamsPTRRecordTypePTR RecordNewParamsPTRRecordType = "PTR"
)

func (r RecordNewParamsPTRRecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsPTRRecordTypePTR:
		return true
	}
	return false
}

type RecordNewParamsSMIMEARecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a SMIMEA record.
	Data param.Field[RecordNewParamsSMIMEARecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsSMIMEARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordNewParamsSMIMEARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsSMIMEARecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsSMIMEARecord) ImplementsRecordNewParams() {

}

// Components of a SMIMEA record.
type RecordNewParamsSMIMEARecordData struct {
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r RecordNewParamsSMIMEARecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordNewParamsSMIMEARecordType string

const (
	RecordNewParamsSMIMEARecordTypeSMIMEA RecordNewParamsSMIMEARecordType = "SMIMEA"
)

func (r RecordNewParamsSMIMEARecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsSMIMEARecordTypeSMIMEA:
		return true
	}
	return false
}

type RecordNewParamsSRVRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a SRV record.
	Data param.Field[RecordNewParamsSRVRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsSRVRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordNewParamsSRVRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsSRVRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsSRVRecord) ImplementsRecordNewParams() {

}

// Components of a SRV record.
type RecordNewParamsSRVRecordData struct {
	// A valid hostname. Deprecated in favor of the regular 'name' outside the data
	// map. This data map field represents the remainder of the full 'name' after the
	// service and protocol.
	Name param.Field[string] `json:"name" format:"hostname"`
	// The port of the service.
	Port param.Field[float64] `json:"port"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// A valid protocol, prefixed with an underscore. Deprecated in favor of the
	// regular 'name' outside the data map. This data map field normally represents the
	// second label of that 'name'.
	Proto param.Field[string] `json:"proto"`
	// A service type, prefixed with an underscore. Deprecated in favor of the regular
	// 'name' outside the data map. This data map field normally represents the first
	// label of that 'name'.
	Service param.Field[string] `json:"service"`
	// A valid hostname.
	Target param.Field[string] `json:"target" format:"hostname"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r RecordNewParamsSRVRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordNewParamsSRVRecordType string

const (
	RecordNewParamsSRVRecordTypeSRV RecordNewParamsSRVRecordType = "SRV"
)

func (r RecordNewParamsSRVRecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsSRVRecordTypeSRV:
		return true
	}
	return false
}

type RecordNewParamsSSHFPRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a SSHFP record.
	Data param.Field[RecordNewParamsSSHFPRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsSSHFPRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordNewParamsSSHFPRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsSSHFPRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsSSHFPRecord) ImplementsRecordNewParams() {

}

// Components of a SSHFP record.
type RecordNewParamsSSHFPRecordData struct {
	// algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// fingerprint.
	Fingerprint param.Field[string] `json:"fingerprint"`
	// type.
	Type param.Field[float64] `json:"type"`
}

func (r RecordNewParamsSSHFPRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordNewParamsSSHFPRecordType string

const (
	RecordNewParamsSSHFPRecordTypeSSHFP RecordNewParamsSSHFPRecordType = "SSHFP"
)

func (r RecordNewParamsSSHFPRecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsSSHFPRecordTypeSSHFP:
		return true
	}
	return false
}

type RecordNewParamsSVCBRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a SVCB record.
	Data param.Field[RecordNewParamsSVCBRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsSVCBRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordNewParamsSVCBRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsSVCBRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsSVCBRecord) ImplementsRecordNewParams() {

}

// Components of a SVCB record.
type RecordNewParamsSVCBRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r RecordNewParamsSVCBRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordNewParamsSVCBRecordType string

const (
	RecordNewParamsSVCBRecordTypeSVCB RecordNewParamsSVCBRecordType = "SVCB"
)

func (r RecordNewParamsSVCBRecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsSVCBRecordTypeSVCB:
		return true
	}
	return false
}

type RecordNewParamsTLSARecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a TLSA record.
	Data param.Field[RecordNewParamsTLSARecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsTLSARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordNewParamsTLSARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsTLSARecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsTLSARecord) ImplementsRecordNewParams() {

}

// Components of a TLSA record.
type RecordNewParamsTLSARecordData struct {
	// certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r RecordNewParamsTLSARecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordNewParamsTLSARecordType string

const (
	RecordNewParamsTLSARecordTypeTLSA RecordNewParamsTLSARecordType = "TLSA"
)

func (r RecordNewParamsTLSARecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsTLSARecordTypeTLSA:
		return true
	}
	return false
}

type RecordNewParamsTXTRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Text content for the record.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsTXTRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordNewParamsTXTRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsTXTRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsTXTRecord) ImplementsRecordNewParams() {

}

// Record type.
type RecordNewParamsTXTRecordType string

const (
	RecordNewParamsTXTRecordTypeTXT RecordNewParamsTXTRecordType = "TXT"
)

func (r RecordNewParamsTXTRecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsTXTRecordTypeTXT:
		return true
	}
	return false
}

type RecordNewParamsURIRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a URI record.
	Data param.Field[RecordNewParamsURIRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority,required"`
	// Record type.
	Type param.Field[RecordNewParamsURIRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordNewParamsURIRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsURIRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsURIRecord) ImplementsRecordNewParams() {

}

// Components of a URI record.
type RecordNewParamsURIRecordData struct {
	// The record content.
	Content param.Field[string] `json:"content"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r RecordNewParamsURIRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordNewParamsURIRecordType string

const (
	RecordNewParamsURIRecordTypeURI RecordNewParamsURIRecordType = "URI"
)

func (r RecordNewParamsURIRecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsURIRecordTypeURI:
		return true
	}
	return false
}

type RecordNewResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   Record                                                    `json:"result,required"`
	// Whether the API call was successful
	Success RecordNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    recordNewResponseEnvelopeJSON    `json:"-"`
}

// recordNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RecordNewResponseEnvelope]
type recordNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RecordNewResponseEnvelopeSuccess bool

const (
	RecordNewResponseEnvelopeSuccessTrue RecordNewResponseEnvelopeSuccess = true
)

func (r RecordNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RecordNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

// This interface is a union satisfied by one of the following:
// [RecordUpdateParamsARecord], [RecordUpdateParamsAAAARecord],
// [RecordUpdateParamsCAARecord], [RecordUpdateParamsCERTRecord],
// [RecordUpdateParamsCNAMERecord], [RecordUpdateParamsDNSKEYRecord],
// [RecordUpdateParamsDSRecord], [RecordUpdateParamsHTTPSRecord],
// [RecordUpdateParamsLOCRecord], [RecordUpdateParamsMXRecord],
// [RecordUpdateParamsNAPTRRecord], [RecordUpdateParamsNSRecord],
// [RecordUpdateParamsPTRRecord], [RecordUpdateParamsSMIMEARecord],
// [RecordUpdateParamsSRVRecord], [RecordUpdateParamsSSHFPRecord],
// [RecordUpdateParamsSVCBRecord], [RecordUpdateParamsTLSARecord],
// [RecordUpdateParamsTXTRecord], [RecordUpdateParamsURIRecord].
type RecordUpdateParams interface {
	ImplementsRecordUpdateParams()

	getZoneID() param.Field[string]
}

type RecordUpdateParamsARecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A valid IPv4 address.
	Content param.Field[string] `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordUpdateParamsARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsARecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsARecord) ImplementsRecordUpdateParams() {

}

// Record type.
type RecordUpdateParamsARecordType string

const (
	RecordUpdateParamsARecordTypeA RecordUpdateParamsARecordType = "A"
)

func (r RecordUpdateParamsARecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsARecordTypeA:
		return true
	}
	return false
}

type RecordUpdateParamsAAAARecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A valid IPv6 address.
	Content param.Field[string] `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsAAAARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordUpdateParamsAAAARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsAAAARecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsAAAARecord) ImplementsRecordUpdateParams() {

}

// Record type.
type RecordUpdateParamsAAAARecordType string

const (
	RecordUpdateParamsAAAARecordTypeAAAA RecordUpdateParamsAAAARecordType = "AAAA"
)

func (r RecordUpdateParamsAAAARecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsAAAARecordTypeAAAA:
		return true
	}
	return false
}

type RecordUpdateParamsCAARecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a CAA record.
	Data param.Field[RecordUpdateParamsCAARecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsCAARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordUpdateParamsCAARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsCAARecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsCAARecord) ImplementsRecordUpdateParams() {

}

// Components of a CAA record.
type RecordUpdateParamsCAARecordData struct {
	// Flags for the CAA record.
	Flags param.Field[float64] `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag param.Field[string] `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value param.Field[string] `json:"value"`
}

func (r RecordUpdateParamsCAARecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordUpdateParamsCAARecordType string

const (
	RecordUpdateParamsCAARecordTypeCAA RecordUpdateParamsCAARecordType = "CAA"
)

func (r RecordUpdateParamsCAARecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsCAARecordTypeCAA:
		return true
	}
	return false
}

type RecordUpdateParamsCERTRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a CERT record.
	Data param.Field[RecordUpdateParamsCERTRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsCERTRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordUpdateParamsCERTRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsCERTRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsCERTRecord) ImplementsRecordUpdateParams() {

}

// Components of a CERT record.
type RecordUpdateParamsCERTRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
	// Type.
	Type param.Field[float64] `json:"type"`
}

func (r RecordUpdateParamsCERTRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordUpdateParamsCERTRecordType string

const (
	RecordUpdateParamsCERTRecordTypeCERT RecordUpdateParamsCERTRecordType = "CERT"
)

func (r RecordUpdateParamsCERTRecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsCERTRecordTypeCERT:
		return true
	}
	return false
}

type RecordUpdateParamsCNAMERecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A valid hostname. Must not match the record's name.
	Content param.Field[interface{}] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsCNAMERecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordUpdateParamsCNAMERecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsCNAMERecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsCNAMERecord) ImplementsRecordUpdateParams() {

}

// Record type.
type RecordUpdateParamsCNAMERecordType string

const (
	RecordUpdateParamsCNAMERecordTypeCNAME RecordUpdateParamsCNAMERecordType = "CNAME"
)

func (r RecordUpdateParamsCNAMERecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsCNAMERecordTypeCNAME:
		return true
	}
	return false
}

type RecordUpdateParamsDNSKEYRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a DNSKEY record.
	Data param.Field[RecordUpdateParamsDNSKEYRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsDNSKEYRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordUpdateParamsDNSKEYRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsDNSKEYRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsDNSKEYRecord) ImplementsRecordUpdateParams() {

}

// Components of a DNSKEY record.
type RecordUpdateParamsDNSKEYRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Flags.
	Flags param.Field[float64] `json:"flags"`
	// Protocol.
	Protocol param.Field[float64] `json:"protocol"`
	// Public Key.
	PublicKey param.Field[string] `json:"public_key"`
}

func (r RecordUpdateParamsDNSKEYRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordUpdateParamsDNSKEYRecordType string

const (
	RecordUpdateParamsDNSKEYRecordTypeDNSKEY RecordUpdateParamsDNSKEYRecordType = "DNSKEY"
)

func (r RecordUpdateParamsDNSKEYRecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSKEYRecordTypeDNSKEY:
		return true
	}
	return false
}

type RecordUpdateParamsDSRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a DS record.
	Data param.Field[RecordUpdateParamsDSRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsDSRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordUpdateParamsDSRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsDSRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsDSRecord) ImplementsRecordUpdateParams() {

}

// Components of a DS record.
type RecordUpdateParamsDSRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Digest.
	Digest param.Field[string] `json:"digest"`
	// Digest Type.
	DigestType param.Field[float64] `json:"digest_type"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
}

func (r RecordUpdateParamsDSRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordUpdateParamsDSRecordType string

const (
	RecordUpdateParamsDSRecordTypeDS RecordUpdateParamsDSRecordType = "DS"
)

func (r RecordUpdateParamsDSRecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDSRecordTypeDS:
		return true
	}
	return false
}

type RecordUpdateParamsHTTPSRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a HTTPS record.
	Data param.Field[RecordUpdateParamsHTTPSRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsHTTPSRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordUpdateParamsHTTPSRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsHTTPSRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsHTTPSRecord) ImplementsRecordUpdateParams() {

}

// Components of a HTTPS record.
type RecordUpdateParamsHTTPSRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r RecordUpdateParamsHTTPSRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordUpdateParamsHTTPSRecordType string

const (
	RecordUpdateParamsHTTPSRecordTypeHTTPS RecordUpdateParamsHTTPSRecordType = "HTTPS"
)

func (r RecordUpdateParamsHTTPSRecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsHTTPSRecordTypeHTTPS:
		return true
	}
	return false
}

type RecordUpdateParamsLOCRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a LOC record.
	Data param.Field[RecordUpdateParamsLOCRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsLOCRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordUpdateParamsLOCRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsLOCRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsLOCRecord) ImplementsRecordUpdateParams() {

}

// Components of a LOC record.
type RecordUpdateParamsLOCRecordData struct {
	// Altitude of location in meters.
	Altitude param.Field[float64] `json:"altitude"`
	// Degrees of latitude.
	LatDegrees param.Field[float64] `json:"lat_degrees"`
	// Latitude direction.
	LatDirection param.Field[RecordUpdateParamsLOCRecordDataLatDirection] `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes param.Field[float64] `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds param.Field[float64] `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees param.Field[float64] `json:"long_degrees"`
	// Longitude direction.
	LongDirection param.Field[RecordUpdateParamsLOCRecordDataLongDirection] `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes param.Field[float64] `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds param.Field[float64] `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz param.Field[float64] `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert param.Field[float64] `json:"precision_vert"`
	// Size of location in meters.
	Size param.Field[float64] `json:"size"`
}

func (r RecordUpdateParamsLOCRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Latitude direction.
type RecordUpdateParamsLOCRecordDataLatDirection string

const (
	RecordUpdateParamsLOCRecordDataLatDirectionN RecordUpdateParamsLOCRecordDataLatDirection = "N"
	RecordUpdateParamsLOCRecordDataLatDirectionS RecordUpdateParamsLOCRecordDataLatDirection = "S"
)

func (r RecordUpdateParamsLOCRecordDataLatDirection) IsKnown() bool {
	switch r {
	case RecordUpdateParamsLOCRecordDataLatDirectionN, RecordUpdateParamsLOCRecordDataLatDirectionS:
		return true
	}
	return false
}

// Longitude direction.
type RecordUpdateParamsLOCRecordDataLongDirection string

const (
	RecordUpdateParamsLOCRecordDataLongDirectionE RecordUpdateParamsLOCRecordDataLongDirection = "E"
	RecordUpdateParamsLOCRecordDataLongDirectionW RecordUpdateParamsLOCRecordDataLongDirection = "W"
)

func (r RecordUpdateParamsLOCRecordDataLongDirection) IsKnown() bool {
	switch r {
	case RecordUpdateParamsLOCRecordDataLongDirectionE, RecordUpdateParamsLOCRecordDataLongDirectionW:
		return true
	}
	return false
}

// Record type.
type RecordUpdateParamsLOCRecordType string

const (
	RecordUpdateParamsLOCRecordTypeLOC RecordUpdateParamsLOCRecordType = "LOC"
)

func (r RecordUpdateParamsLOCRecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsLOCRecordTypeLOC:
		return true
	}
	return false
}

type RecordUpdateParamsMXRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A valid mail server hostname.
	Content param.Field[string] `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsMXRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordUpdateParamsMXRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsMXRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsMXRecord) ImplementsRecordUpdateParams() {

}

// Record type.
type RecordUpdateParamsMXRecordType string

const (
	RecordUpdateParamsMXRecordTypeMX RecordUpdateParamsMXRecordType = "MX"
)

func (r RecordUpdateParamsMXRecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsMXRecordTypeMX:
		return true
	}
	return false
}

type RecordUpdateParamsNAPTRRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a NAPTR record.
	Data param.Field[RecordUpdateParamsNAPTRRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsNAPTRRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordUpdateParamsNAPTRRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsNAPTRRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsNAPTRRecord) ImplementsRecordUpdateParams() {

}

// Components of a NAPTR record.
type RecordUpdateParamsNAPTRRecordData struct {
	// Flags.
	Flags param.Field[string] `json:"flags"`
	// Order.
	Order param.Field[float64] `json:"order"`
	// Preference.
	Preference param.Field[float64] `json:"preference"`
	// Regex.
	Regex param.Field[string] `json:"regex"`
	// Replacement.
	Replacement param.Field[string] `json:"replacement"`
	// Service.
	Service param.Field[string] `json:"service"`
}

func (r RecordUpdateParamsNAPTRRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordUpdateParamsNAPTRRecordType string

const (
	RecordUpdateParamsNAPTRRecordTypeNAPTR RecordUpdateParamsNAPTRRecordType = "NAPTR"
)

func (r RecordUpdateParamsNAPTRRecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsNAPTRRecordTypeNAPTR:
		return true
	}
	return false
}

type RecordUpdateParamsNSRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A valid name server host name.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsNSRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordUpdateParamsNSRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsNSRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsNSRecord) ImplementsRecordUpdateParams() {

}

// Record type.
type RecordUpdateParamsNSRecordType string

const (
	RecordUpdateParamsNSRecordTypeNS RecordUpdateParamsNSRecordType = "NS"
)

func (r RecordUpdateParamsNSRecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsNSRecordTypeNS:
		return true
	}
	return false
}

type RecordUpdateParamsPTRRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Domain name pointing to the address.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsPTRRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordUpdateParamsPTRRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsPTRRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsPTRRecord) ImplementsRecordUpdateParams() {

}

// Record type.
type RecordUpdateParamsPTRRecordType string

const (
	RecordUpdateParamsPTRRecordTypePTR RecordUpdateParamsPTRRecordType = "PTR"
)

func (r RecordUpdateParamsPTRRecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsPTRRecordTypePTR:
		return true
	}
	return false
}

type RecordUpdateParamsSMIMEARecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a SMIMEA record.
	Data param.Field[RecordUpdateParamsSMIMEARecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsSMIMEARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordUpdateParamsSMIMEARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsSMIMEARecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsSMIMEARecord) ImplementsRecordUpdateParams() {

}

// Components of a SMIMEA record.
type RecordUpdateParamsSMIMEARecordData struct {
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r RecordUpdateParamsSMIMEARecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordUpdateParamsSMIMEARecordType string

const (
	RecordUpdateParamsSMIMEARecordTypeSMIMEA RecordUpdateParamsSMIMEARecordType = "SMIMEA"
)

func (r RecordUpdateParamsSMIMEARecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsSMIMEARecordTypeSMIMEA:
		return true
	}
	return false
}

type RecordUpdateParamsSRVRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a SRV record.
	Data param.Field[RecordUpdateParamsSRVRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsSRVRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordUpdateParamsSRVRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsSRVRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsSRVRecord) ImplementsRecordUpdateParams() {

}

// Components of a SRV record.
type RecordUpdateParamsSRVRecordData struct {
	// A valid hostname. Deprecated in favor of the regular 'name' outside the data
	// map. This data map field represents the remainder of the full 'name' after the
	// service and protocol.
	Name param.Field[string] `json:"name" format:"hostname"`
	// The port of the service.
	Port param.Field[float64] `json:"port"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// A valid protocol, prefixed with an underscore. Deprecated in favor of the
	// regular 'name' outside the data map. This data map field normally represents the
	// second label of that 'name'.
	Proto param.Field[string] `json:"proto"`
	// A service type, prefixed with an underscore. Deprecated in favor of the regular
	// 'name' outside the data map. This data map field normally represents the first
	// label of that 'name'.
	Service param.Field[string] `json:"service"`
	// A valid hostname.
	Target param.Field[string] `json:"target" format:"hostname"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r RecordUpdateParamsSRVRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordUpdateParamsSRVRecordType string

const (
	RecordUpdateParamsSRVRecordTypeSRV RecordUpdateParamsSRVRecordType = "SRV"
)

func (r RecordUpdateParamsSRVRecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsSRVRecordTypeSRV:
		return true
	}
	return false
}

type RecordUpdateParamsSSHFPRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a SSHFP record.
	Data param.Field[RecordUpdateParamsSSHFPRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsSSHFPRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordUpdateParamsSSHFPRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsSSHFPRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsSSHFPRecord) ImplementsRecordUpdateParams() {

}

// Components of a SSHFP record.
type RecordUpdateParamsSSHFPRecordData struct {
	// algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// fingerprint.
	Fingerprint param.Field[string] `json:"fingerprint"`
	// type.
	Type param.Field[float64] `json:"type"`
}

func (r RecordUpdateParamsSSHFPRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordUpdateParamsSSHFPRecordType string

const (
	RecordUpdateParamsSSHFPRecordTypeSSHFP RecordUpdateParamsSSHFPRecordType = "SSHFP"
)

func (r RecordUpdateParamsSSHFPRecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsSSHFPRecordTypeSSHFP:
		return true
	}
	return false
}

type RecordUpdateParamsSVCBRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a SVCB record.
	Data param.Field[RecordUpdateParamsSVCBRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsSVCBRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordUpdateParamsSVCBRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsSVCBRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsSVCBRecord) ImplementsRecordUpdateParams() {

}

// Components of a SVCB record.
type RecordUpdateParamsSVCBRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r RecordUpdateParamsSVCBRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordUpdateParamsSVCBRecordType string

const (
	RecordUpdateParamsSVCBRecordTypeSVCB RecordUpdateParamsSVCBRecordType = "SVCB"
)

func (r RecordUpdateParamsSVCBRecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsSVCBRecordTypeSVCB:
		return true
	}
	return false
}

type RecordUpdateParamsTLSARecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a TLSA record.
	Data param.Field[RecordUpdateParamsTLSARecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsTLSARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordUpdateParamsTLSARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsTLSARecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsTLSARecord) ImplementsRecordUpdateParams() {

}

// Components of a TLSA record.
type RecordUpdateParamsTLSARecordData struct {
	// certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r RecordUpdateParamsTLSARecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordUpdateParamsTLSARecordType string

const (
	RecordUpdateParamsTLSARecordTypeTLSA RecordUpdateParamsTLSARecordType = "TLSA"
)

func (r RecordUpdateParamsTLSARecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsTLSARecordTypeTLSA:
		return true
	}
	return false
}

type RecordUpdateParamsTXTRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Text content for the record.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsTXTRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordUpdateParamsTXTRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsTXTRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsTXTRecord) ImplementsRecordUpdateParams() {

}

// Record type.
type RecordUpdateParamsTXTRecordType string

const (
	RecordUpdateParamsTXTRecordTypeTXT RecordUpdateParamsTXTRecordType = "TXT"
)

func (r RecordUpdateParamsTXTRecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsTXTRecordTypeTXT:
		return true
	}
	return false
}

type RecordUpdateParamsURIRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a URI record.
	Data param.Field[RecordUpdateParamsURIRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsURIRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordUpdateParamsURIRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsURIRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsURIRecord) ImplementsRecordUpdateParams() {

}

// Components of a URI record.
type RecordUpdateParamsURIRecordData struct {
	// The record content.
	Content param.Field[string] `json:"content"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r RecordUpdateParamsURIRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordUpdateParamsURIRecordType string

const (
	RecordUpdateParamsURIRecordTypeURI RecordUpdateParamsURIRecordType = "URI"
)

func (r RecordUpdateParamsURIRecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsURIRecordTypeURI:
		return true
	}
	return false
}

type RecordUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   Record                                                    `json:"result,required"`
	// Whether the API call was successful
	Success RecordUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    recordUpdateResponseEnvelopeJSON    `json:"-"`
}

// recordUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RecordUpdateResponseEnvelope]
type recordUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RecordUpdateResponseEnvelopeSuccess bool

const (
	RecordUpdateResponseEnvelopeSuccessTrue RecordUpdateResponseEnvelopeSuccess = true
)

func (r RecordUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RecordUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RecordListParams struct {
	// Identifier
	ZoneID  param.Field[string]                  `path:"zone_id,required"`
	Comment param.Field[RecordListParamsComment] `query:"comment"`
	// DNS record content.
	Content param.Field[string] `query:"content"`
	// Direction to order DNS records in.
	Direction param.Field[RecordListParamsDirection] `query:"direction"`
	// Whether to match all search requirements or at least one (any). If set to `all`,
	// acts like a logical AND between filters. If set to `any`, acts like a logical OR
	// instead. Note that the interaction between tag filters is controlled by the
	// `tag-match` parameter instead.
	Match param.Field[RecordListParamsMatch] `query:"match"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `query:"name"`
	// Field to order DNS records by.
	Order param.Field[RecordListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of DNS records per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `query:"proxied"`
	// Allows searching in multiple properties of a DNS record simultaneously. This
	// parameter is intended for human users, not automation. Its exact behavior is
	// intentionally left unspecified and is subject to change in the future. This
	// parameter works independently of the `match` setting. For automated searches,
	// please use the other available parameters.
	Search param.Field[string]              `query:"search"`
	Tag    param.Field[RecordListParamsTag] `query:"tag"`
	// Whether to match all tag search requirements or at least one (any). If set to
	// `all`, acts like a logical AND between tag filters. If set to `any`, acts like a
	// logical OR instead. Note that the regular `match` parameter is still used to
	// combine the resulting condition with other filters that aren't related to tags.
	TagMatch param.Field[RecordListParamsTagMatch] `query:"tag_match"`
	// Record type.
	Type param.Field[RecordListParamsType] `query:"type"`
}

// URLQuery serializes [RecordListParams]'s query parameters as `url.Values`.
func (r RecordListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RecordListParamsComment struct {
	// If this parameter is present, only records _without_ a comment are returned.
	Absent param.Field[string] `query:"absent"`
	// Substring of the DNS record comment. Comment filters are case-insensitive.
	Contains param.Field[string] `query:"contains"`
	// Suffix of the DNS record comment. Comment filters are case-insensitive.
	Endswith param.Field[string] `query:"endswith"`
	// Exact value of the DNS record comment. Comment filters are case-insensitive.
	Exact param.Field[string] `query:"exact"`
	// If this parameter is present, only records _with_ a comment are returned.
	Present param.Field[string] `query:"present"`
	// Prefix of the DNS record comment. Comment filters are case-insensitive.
	Startswith param.Field[string] `query:"startswith"`
}

// URLQuery serializes [RecordListParamsComment]'s query parameters as
// `url.Values`.
func (r RecordListParamsComment) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order DNS records in.
type RecordListParamsDirection string

const (
	RecordListParamsDirectionAsc  RecordListParamsDirection = "asc"
	RecordListParamsDirectionDesc RecordListParamsDirection = "desc"
)

func (r RecordListParamsDirection) IsKnown() bool {
	switch r {
	case RecordListParamsDirectionAsc, RecordListParamsDirectionDesc:
		return true
	}
	return false
}

// Whether to match all search requirements or at least one (any). If set to `all`,
// acts like a logical AND between filters. If set to `any`, acts like a logical OR
// instead. Note that the interaction between tag filters is controlled by the
// `tag-match` parameter instead.
type RecordListParamsMatch string

const (
	RecordListParamsMatchAny RecordListParamsMatch = "any"
	RecordListParamsMatchAll RecordListParamsMatch = "all"
)

func (r RecordListParamsMatch) IsKnown() bool {
	switch r {
	case RecordListParamsMatchAny, RecordListParamsMatchAll:
		return true
	}
	return false
}

// Field to order DNS records by.
type RecordListParamsOrder string

const (
	RecordListParamsOrderType    RecordListParamsOrder = "type"
	RecordListParamsOrderName    RecordListParamsOrder = "name"
	RecordListParamsOrderContent RecordListParamsOrder = "content"
	RecordListParamsOrderTTL     RecordListParamsOrder = "ttl"
	RecordListParamsOrderProxied RecordListParamsOrder = "proxied"
)

func (r RecordListParamsOrder) IsKnown() bool {
	switch r {
	case RecordListParamsOrderType, RecordListParamsOrderName, RecordListParamsOrderContent, RecordListParamsOrderTTL, RecordListParamsOrderProxied:
		return true
	}
	return false
}

type RecordListParamsTag struct {
	// Name of a tag which must _not_ be present on the DNS record. Tag filters are
	// case-insensitive.
	Absent param.Field[string] `query:"absent"`
	// A tag and value, of the form `<tag-name>:<tag-value>`. The API will only return
	// DNS records that have a tag named `<tag-name>` whose value contains
	// `<tag-value>`. Tag filters are case-insensitive.
	Contains param.Field[string] `query:"contains"`
	// A tag and value, of the form `<tag-name>:<tag-value>`. The API will only return
	// DNS records that have a tag named `<tag-name>` whose value ends with
	// `<tag-value>`. Tag filters are case-insensitive.
	Endswith param.Field[string] `query:"endswith"`
	// A tag and value, of the form `<tag-name>:<tag-value>`. The API will only return
	// DNS records that have a tag named `<tag-name>` whose value is `<tag-value>`. Tag
	// filters are case-insensitive.
	Exact param.Field[string] `query:"exact"`
	// Name of a tag which must be present on the DNS record. Tag filters are
	// case-insensitive.
	Present param.Field[string] `query:"present"`
	// A tag and value, of the form `<tag-name>:<tag-value>`. The API will only return
	// DNS records that have a tag named `<tag-name>` whose value starts with
	// `<tag-value>`. Tag filters are case-insensitive.
	Startswith param.Field[string] `query:"startswith"`
}

// URLQuery serializes [RecordListParamsTag]'s query parameters as `url.Values`.
func (r RecordListParamsTag) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Whether to match all tag search requirements or at least one (any). If set to
// `all`, acts like a logical AND between tag filters. If set to `any`, acts like a
// logical OR instead. Note that the regular `match` parameter is still used to
// combine the resulting condition with other filters that aren't related to tags.
type RecordListParamsTagMatch string

const (
	RecordListParamsTagMatchAny RecordListParamsTagMatch = "any"
	RecordListParamsTagMatchAll RecordListParamsTagMatch = "all"
)

func (r RecordListParamsTagMatch) IsKnown() bool {
	switch r {
	case RecordListParamsTagMatchAny, RecordListParamsTagMatchAll:
		return true
	}
	return false
}

// Record type.
type RecordListParamsType string

const (
	RecordListParamsTypeA      RecordListParamsType = "A"
	RecordListParamsTypeAAAA   RecordListParamsType = "AAAA"
	RecordListParamsTypeCAA    RecordListParamsType = "CAA"
	RecordListParamsTypeCERT   RecordListParamsType = "CERT"
	RecordListParamsTypeCNAME  RecordListParamsType = "CNAME"
	RecordListParamsTypeDNSKEY RecordListParamsType = "DNSKEY"
	RecordListParamsTypeDS     RecordListParamsType = "DS"
	RecordListParamsTypeHTTPS  RecordListParamsType = "HTTPS"
	RecordListParamsTypeLOC    RecordListParamsType = "LOC"
	RecordListParamsTypeMX     RecordListParamsType = "MX"
	RecordListParamsTypeNAPTR  RecordListParamsType = "NAPTR"
	RecordListParamsTypeNS     RecordListParamsType = "NS"
	RecordListParamsTypePTR    RecordListParamsType = "PTR"
	RecordListParamsTypeSMIMEA RecordListParamsType = "SMIMEA"
	RecordListParamsTypeSRV    RecordListParamsType = "SRV"
	RecordListParamsTypeSSHFP  RecordListParamsType = "SSHFP"
	RecordListParamsTypeSVCB   RecordListParamsType = "SVCB"
	RecordListParamsTypeTLSA   RecordListParamsType = "TLSA"
	RecordListParamsTypeTXT    RecordListParamsType = "TXT"
	RecordListParamsTypeURI    RecordListParamsType = "URI"
)

func (r RecordListParamsType) IsKnown() bool {
	switch r {
	case RecordListParamsTypeA, RecordListParamsTypeAAAA, RecordListParamsTypeCAA, RecordListParamsTypeCERT, RecordListParamsTypeCNAME, RecordListParamsTypeDNSKEY, RecordListParamsTypeDS, RecordListParamsTypeHTTPS, RecordListParamsTypeLOC, RecordListParamsTypeMX, RecordListParamsTypeNAPTR, RecordListParamsTypeNS, RecordListParamsTypePTR, RecordListParamsTypeSMIMEA, RecordListParamsTypeSRV, RecordListParamsTypeSSHFP, RecordListParamsTypeSVCB, RecordListParamsTypeTLSA, RecordListParamsTypeTXT, RecordListParamsTypeURI:
		return true
	}
	return false
}

type RecordDeleteParams struct {
	// Identifier
	ZoneID param.Field[string]      `path:"zone_id,required"`
	Body   param.Field[interface{}] `json:"body,required"`
}

func (r RecordDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RecordDeleteResponseEnvelope struct {
	Result RecordDeleteResponse             `json:"result"`
	JSON   recordDeleteResponseEnvelopeJSON `json:"-"`
}

// recordDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RecordDeleteResponseEnvelope]
type recordDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// This interface is a union satisfied by one of the following:
// [RecordEditParamsARecord], [RecordEditParamsAAAARecord],
// [RecordEditParamsCAARecord], [RecordEditParamsCERTRecord],
// [RecordEditParamsCNAMERecord], [RecordEditParamsDNSKEYRecord],
// [RecordEditParamsDSRecord], [RecordEditParamsHTTPSRecord],
// [RecordEditParamsLOCRecord], [RecordEditParamsMXRecord],
// [RecordEditParamsNAPTRRecord], [RecordEditParamsNSRecord],
// [RecordEditParamsPTRRecord], [RecordEditParamsSMIMEARecord],
// [RecordEditParamsSRVRecord], [RecordEditParamsSSHFPRecord],
// [RecordEditParamsSVCBRecord], [RecordEditParamsTLSARecord],
// [RecordEditParamsTXTRecord], [RecordEditParamsURIRecord].
type RecordEditParams interface {
	ImplementsRecordEditParams()

	getZoneID() param.Field[string]
}

type RecordEditParamsARecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A valid IPv4 address.
	Content param.Field[string] `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordEditParamsARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsARecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsARecord) ImplementsRecordEditParams() {

}

// Record type.
type RecordEditParamsARecordType string

const (
	RecordEditParamsARecordTypeA RecordEditParamsARecordType = "A"
)

func (r RecordEditParamsARecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsARecordTypeA:
		return true
	}
	return false
}

type RecordEditParamsAAAARecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A valid IPv6 address.
	Content param.Field[string] `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsAAAARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordEditParamsAAAARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsAAAARecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsAAAARecord) ImplementsRecordEditParams() {

}

// Record type.
type RecordEditParamsAAAARecordType string

const (
	RecordEditParamsAAAARecordTypeAAAA RecordEditParamsAAAARecordType = "AAAA"
)

func (r RecordEditParamsAAAARecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsAAAARecordTypeAAAA:
		return true
	}
	return false
}

type RecordEditParamsCAARecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a CAA record.
	Data param.Field[RecordEditParamsCAARecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsCAARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordEditParamsCAARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsCAARecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsCAARecord) ImplementsRecordEditParams() {

}

// Components of a CAA record.
type RecordEditParamsCAARecordData struct {
	// Flags for the CAA record.
	Flags param.Field[float64] `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag param.Field[string] `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value param.Field[string] `json:"value"`
}

func (r RecordEditParamsCAARecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordEditParamsCAARecordType string

const (
	RecordEditParamsCAARecordTypeCAA RecordEditParamsCAARecordType = "CAA"
)

func (r RecordEditParamsCAARecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsCAARecordTypeCAA:
		return true
	}
	return false
}

type RecordEditParamsCERTRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a CERT record.
	Data param.Field[RecordEditParamsCERTRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsCERTRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordEditParamsCERTRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsCERTRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsCERTRecord) ImplementsRecordEditParams() {

}

// Components of a CERT record.
type RecordEditParamsCERTRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
	// Type.
	Type param.Field[float64] `json:"type"`
}

func (r RecordEditParamsCERTRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordEditParamsCERTRecordType string

const (
	RecordEditParamsCERTRecordTypeCERT RecordEditParamsCERTRecordType = "CERT"
)

func (r RecordEditParamsCERTRecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsCERTRecordTypeCERT:
		return true
	}
	return false
}

type RecordEditParamsCNAMERecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A valid hostname. Must not match the record's name.
	Content param.Field[interface{}] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsCNAMERecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordEditParamsCNAMERecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsCNAMERecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsCNAMERecord) ImplementsRecordEditParams() {

}

// Record type.
type RecordEditParamsCNAMERecordType string

const (
	RecordEditParamsCNAMERecordTypeCNAME RecordEditParamsCNAMERecordType = "CNAME"
)

func (r RecordEditParamsCNAMERecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsCNAMERecordTypeCNAME:
		return true
	}
	return false
}

type RecordEditParamsDNSKEYRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a DNSKEY record.
	Data param.Field[RecordEditParamsDNSKEYRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsDNSKEYRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordEditParamsDNSKEYRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsDNSKEYRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsDNSKEYRecord) ImplementsRecordEditParams() {

}

// Components of a DNSKEY record.
type RecordEditParamsDNSKEYRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Flags.
	Flags param.Field[float64] `json:"flags"`
	// Protocol.
	Protocol param.Field[float64] `json:"protocol"`
	// Public Key.
	PublicKey param.Field[string] `json:"public_key"`
}

func (r RecordEditParamsDNSKEYRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordEditParamsDNSKEYRecordType string

const (
	RecordEditParamsDNSKEYRecordTypeDNSKEY RecordEditParamsDNSKEYRecordType = "DNSKEY"
)

func (r RecordEditParamsDNSKEYRecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSKEYRecordTypeDNSKEY:
		return true
	}
	return false
}

type RecordEditParamsDSRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a DS record.
	Data param.Field[RecordEditParamsDSRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsDSRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordEditParamsDSRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsDSRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsDSRecord) ImplementsRecordEditParams() {

}

// Components of a DS record.
type RecordEditParamsDSRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Digest.
	Digest param.Field[string] `json:"digest"`
	// Digest Type.
	DigestType param.Field[float64] `json:"digest_type"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
}

func (r RecordEditParamsDSRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordEditParamsDSRecordType string

const (
	RecordEditParamsDSRecordTypeDS RecordEditParamsDSRecordType = "DS"
)

func (r RecordEditParamsDSRecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsDSRecordTypeDS:
		return true
	}
	return false
}

type RecordEditParamsHTTPSRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a HTTPS record.
	Data param.Field[RecordEditParamsHTTPSRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsHTTPSRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordEditParamsHTTPSRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsHTTPSRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsHTTPSRecord) ImplementsRecordEditParams() {

}

// Components of a HTTPS record.
type RecordEditParamsHTTPSRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r RecordEditParamsHTTPSRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordEditParamsHTTPSRecordType string

const (
	RecordEditParamsHTTPSRecordTypeHTTPS RecordEditParamsHTTPSRecordType = "HTTPS"
)

func (r RecordEditParamsHTTPSRecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsHTTPSRecordTypeHTTPS:
		return true
	}
	return false
}

type RecordEditParamsLOCRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a LOC record.
	Data param.Field[RecordEditParamsLOCRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsLOCRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordEditParamsLOCRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsLOCRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsLOCRecord) ImplementsRecordEditParams() {

}

// Components of a LOC record.
type RecordEditParamsLOCRecordData struct {
	// Altitude of location in meters.
	Altitude param.Field[float64] `json:"altitude"`
	// Degrees of latitude.
	LatDegrees param.Field[float64] `json:"lat_degrees"`
	// Latitude direction.
	LatDirection param.Field[RecordEditParamsLOCRecordDataLatDirection] `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes param.Field[float64] `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds param.Field[float64] `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees param.Field[float64] `json:"long_degrees"`
	// Longitude direction.
	LongDirection param.Field[RecordEditParamsLOCRecordDataLongDirection] `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes param.Field[float64] `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds param.Field[float64] `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz param.Field[float64] `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert param.Field[float64] `json:"precision_vert"`
	// Size of location in meters.
	Size param.Field[float64] `json:"size"`
}

func (r RecordEditParamsLOCRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Latitude direction.
type RecordEditParamsLOCRecordDataLatDirection string

const (
	RecordEditParamsLOCRecordDataLatDirectionN RecordEditParamsLOCRecordDataLatDirection = "N"
	RecordEditParamsLOCRecordDataLatDirectionS RecordEditParamsLOCRecordDataLatDirection = "S"
)

func (r RecordEditParamsLOCRecordDataLatDirection) IsKnown() bool {
	switch r {
	case RecordEditParamsLOCRecordDataLatDirectionN, RecordEditParamsLOCRecordDataLatDirectionS:
		return true
	}
	return false
}

// Longitude direction.
type RecordEditParamsLOCRecordDataLongDirection string

const (
	RecordEditParamsLOCRecordDataLongDirectionE RecordEditParamsLOCRecordDataLongDirection = "E"
	RecordEditParamsLOCRecordDataLongDirectionW RecordEditParamsLOCRecordDataLongDirection = "W"
)

func (r RecordEditParamsLOCRecordDataLongDirection) IsKnown() bool {
	switch r {
	case RecordEditParamsLOCRecordDataLongDirectionE, RecordEditParamsLOCRecordDataLongDirectionW:
		return true
	}
	return false
}

// Record type.
type RecordEditParamsLOCRecordType string

const (
	RecordEditParamsLOCRecordTypeLOC RecordEditParamsLOCRecordType = "LOC"
)

func (r RecordEditParamsLOCRecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsLOCRecordTypeLOC:
		return true
	}
	return false
}

type RecordEditParamsMXRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A valid mail server hostname.
	Content param.Field[string] `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority,required"`
	// Record type.
	Type param.Field[RecordEditParamsMXRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordEditParamsMXRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsMXRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsMXRecord) ImplementsRecordEditParams() {

}

// Record type.
type RecordEditParamsMXRecordType string

const (
	RecordEditParamsMXRecordTypeMX RecordEditParamsMXRecordType = "MX"
)

func (r RecordEditParamsMXRecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsMXRecordTypeMX:
		return true
	}
	return false
}

type RecordEditParamsNAPTRRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a NAPTR record.
	Data param.Field[RecordEditParamsNAPTRRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsNAPTRRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordEditParamsNAPTRRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsNAPTRRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsNAPTRRecord) ImplementsRecordEditParams() {

}

// Components of a NAPTR record.
type RecordEditParamsNAPTRRecordData struct {
	// Flags.
	Flags param.Field[string] `json:"flags"`
	// Order.
	Order param.Field[float64] `json:"order"`
	// Preference.
	Preference param.Field[float64] `json:"preference"`
	// Regex.
	Regex param.Field[string] `json:"regex"`
	// Replacement.
	Replacement param.Field[string] `json:"replacement"`
	// Service.
	Service param.Field[string] `json:"service"`
}

func (r RecordEditParamsNAPTRRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordEditParamsNAPTRRecordType string

const (
	RecordEditParamsNAPTRRecordTypeNAPTR RecordEditParamsNAPTRRecordType = "NAPTR"
)

func (r RecordEditParamsNAPTRRecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsNAPTRRecordTypeNAPTR:
		return true
	}
	return false
}

type RecordEditParamsNSRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A valid name server host name.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsNSRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordEditParamsNSRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsNSRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsNSRecord) ImplementsRecordEditParams() {

}

// Record type.
type RecordEditParamsNSRecordType string

const (
	RecordEditParamsNSRecordTypeNS RecordEditParamsNSRecordType = "NS"
)

func (r RecordEditParamsNSRecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsNSRecordTypeNS:
		return true
	}
	return false
}

type RecordEditParamsPTRRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Domain name pointing to the address.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsPTRRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordEditParamsPTRRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsPTRRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsPTRRecord) ImplementsRecordEditParams() {

}

// Record type.
type RecordEditParamsPTRRecordType string

const (
	RecordEditParamsPTRRecordTypePTR RecordEditParamsPTRRecordType = "PTR"
)

func (r RecordEditParamsPTRRecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsPTRRecordTypePTR:
		return true
	}
	return false
}

type RecordEditParamsSMIMEARecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a SMIMEA record.
	Data param.Field[RecordEditParamsSMIMEARecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsSMIMEARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordEditParamsSMIMEARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsSMIMEARecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsSMIMEARecord) ImplementsRecordEditParams() {

}

// Components of a SMIMEA record.
type RecordEditParamsSMIMEARecordData struct {
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r RecordEditParamsSMIMEARecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordEditParamsSMIMEARecordType string

const (
	RecordEditParamsSMIMEARecordTypeSMIMEA RecordEditParamsSMIMEARecordType = "SMIMEA"
)

func (r RecordEditParamsSMIMEARecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsSMIMEARecordTypeSMIMEA:
		return true
	}
	return false
}

type RecordEditParamsSRVRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a SRV record.
	Data param.Field[RecordEditParamsSRVRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsSRVRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordEditParamsSRVRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsSRVRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsSRVRecord) ImplementsRecordEditParams() {

}

// Components of a SRV record.
type RecordEditParamsSRVRecordData struct {
	// A valid hostname. Deprecated in favor of the regular 'name' outside the data
	// map. This data map field represents the remainder of the full 'name' after the
	// service and protocol.
	Name param.Field[string] `json:"name" format:"hostname"`
	// The port of the service.
	Port param.Field[float64] `json:"port"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// A valid protocol, prefixed with an underscore. Deprecated in favor of the
	// regular 'name' outside the data map. This data map field normally represents the
	// second label of that 'name'.
	Proto param.Field[string] `json:"proto"`
	// A service type, prefixed with an underscore. Deprecated in favor of the regular
	// 'name' outside the data map. This data map field normally represents the first
	// label of that 'name'.
	Service param.Field[string] `json:"service"`
	// A valid hostname.
	Target param.Field[string] `json:"target" format:"hostname"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r RecordEditParamsSRVRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordEditParamsSRVRecordType string

const (
	RecordEditParamsSRVRecordTypeSRV RecordEditParamsSRVRecordType = "SRV"
)

func (r RecordEditParamsSRVRecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsSRVRecordTypeSRV:
		return true
	}
	return false
}

type RecordEditParamsSSHFPRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a SSHFP record.
	Data param.Field[RecordEditParamsSSHFPRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsSSHFPRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordEditParamsSSHFPRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsSSHFPRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsSSHFPRecord) ImplementsRecordEditParams() {

}

// Components of a SSHFP record.
type RecordEditParamsSSHFPRecordData struct {
	// algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// fingerprint.
	Fingerprint param.Field[string] `json:"fingerprint"`
	// type.
	Type param.Field[float64] `json:"type"`
}

func (r RecordEditParamsSSHFPRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordEditParamsSSHFPRecordType string

const (
	RecordEditParamsSSHFPRecordTypeSSHFP RecordEditParamsSSHFPRecordType = "SSHFP"
)

func (r RecordEditParamsSSHFPRecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsSSHFPRecordTypeSSHFP:
		return true
	}
	return false
}

type RecordEditParamsSVCBRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a SVCB record.
	Data param.Field[RecordEditParamsSVCBRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsSVCBRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordEditParamsSVCBRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsSVCBRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsSVCBRecord) ImplementsRecordEditParams() {

}

// Components of a SVCB record.
type RecordEditParamsSVCBRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r RecordEditParamsSVCBRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordEditParamsSVCBRecordType string

const (
	RecordEditParamsSVCBRecordTypeSVCB RecordEditParamsSVCBRecordType = "SVCB"
)

func (r RecordEditParamsSVCBRecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsSVCBRecordTypeSVCB:
		return true
	}
	return false
}

type RecordEditParamsTLSARecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a TLSA record.
	Data param.Field[RecordEditParamsTLSARecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsTLSARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordEditParamsTLSARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsTLSARecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsTLSARecord) ImplementsRecordEditParams() {

}

// Components of a TLSA record.
type RecordEditParamsTLSARecordData struct {
	// certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r RecordEditParamsTLSARecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordEditParamsTLSARecordType string

const (
	RecordEditParamsTLSARecordTypeTLSA RecordEditParamsTLSARecordType = "TLSA"
)

func (r RecordEditParamsTLSARecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsTLSARecordTypeTLSA:
		return true
	}
	return false
}

type RecordEditParamsTXTRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Text content for the record.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsTXTRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordEditParamsTXTRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsTXTRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsTXTRecord) ImplementsRecordEditParams() {

}

// Record type.
type RecordEditParamsTXTRecordType string

const (
	RecordEditParamsTXTRecordTypeTXT RecordEditParamsTXTRecordType = "TXT"
)

func (r RecordEditParamsTXTRecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsTXTRecordTypeTXT:
		return true
	}
	return false
}

type RecordEditParamsURIRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a URI record.
	Data param.Field[RecordEditParamsURIRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority,required"`
	// Record type.
	Type param.Field[RecordEditParamsURIRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsItemParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTLUnionParam] `json:"ttl"`
}

func (r RecordEditParamsURIRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsURIRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsURIRecord) ImplementsRecordEditParams() {

}

// Components of a URI record.
type RecordEditParamsURIRecordData struct {
	// The record content.
	Content param.Field[string] `json:"content"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r RecordEditParamsURIRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordEditParamsURIRecordType string

const (
	RecordEditParamsURIRecordTypeURI RecordEditParamsURIRecordType = "URI"
)

func (r RecordEditParamsURIRecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsURIRecordTypeURI:
		return true
	}
	return false
}

type RecordEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   Record                                                    `json:"result,required"`
	// Whether the API call was successful
	Success RecordEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    recordEditResponseEnvelopeJSON    `json:"-"`
}

// recordEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [RecordEditResponseEnvelope]
type recordEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RecordEditResponseEnvelopeSuccess bool

const (
	RecordEditResponseEnvelopeSuccessTrue RecordEditResponseEnvelopeSuccess = true
)

func (r RecordEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RecordEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RecordExportParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RecordGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RecordGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   Record                                                    `json:"result,required"`
	// Whether the API call was successful
	Success RecordGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    recordGetResponseEnvelopeJSON    `json:"-"`
}

// recordGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RecordGetResponseEnvelope]
type recordGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RecordGetResponseEnvelopeSuccess bool

const (
	RecordGetResponseEnvelopeSuccessTrue RecordGetResponseEnvelopeSuccess = true
)

func (r RecordGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RecordGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RecordImportParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// BIND config to import.
	//
	// **Tip:** When using cURL, a file can be uploaded using
	// `--form 'file=@bind_config.txt'`.
	File param.Field[string] `json:"file,required"`
	// Whether or not proxiable records should receive the performance and security
	// benefits of Cloudflare.
	//
	// The value should be either `true` or `false`.
	Proxied param.Field[string] `json:"proxied"`
}

func (r RecordImportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RecordImportResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   RecordImportResponse                                      `json:"result,required"`
	// Whether the API call was successful
	Success RecordImportResponseEnvelopeSuccess              `json:"success,required"`
	Timing  UnnamedSchemaRef2f895e023ae55b55d2f5925449b819cd `json:"timing"`
	JSON    recordImportResponseEnvelopeJSON                 `json:"-"`
}

// recordImportResponseEnvelopeJSON contains the JSON metadata for the struct
// [RecordImportResponseEnvelope]
type recordImportResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	Timing      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordImportResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordImportResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RecordImportResponseEnvelopeSuccess bool

const (
	RecordImportResponseEnvelopeSuccessTrue RecordImportResponseEnvelopeSuccess = true
)

func (r RecordImportResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RecordImportResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RecordScanParams struct {
	// Identifier
	ZoneID param.Field[string]      `path:"zone_id,required"`
	Body   param.Field[interface{}] `json:"body,required"`
}

func (r RecordScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RecordScanResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   RecordScanResponse                                        `json:"result,required"`
	// Whether the API call was successful
	Success RecordScanResponseEnvelopeSuccess                `json:"success,required"`
	Timing  UnnamedSchemaRef2f895e023ae55b55d2f5925449b819cd `json:"timing"`
	JSON    recordScanResponseEnvelopeJSON                   `json:"-"`
}

// recordScanResponseEnvelopeJSON contains the JSON metadata for the struct
// [RecordScanResponseEnvelope]
type recordScanResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	Timing      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordScanResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordScanResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RecordScanResponseEnvelopeSuccess bool

const (
	RecordScanResponseEnvelopeSuccessTrue RecordScanResponseEnvelopeSuccess = true
)

func (r RecordScanResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RecordScanResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
