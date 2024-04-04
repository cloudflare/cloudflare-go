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
func (r *RecordService) New(ctx context.Context, params RecordNewParams, opts ...option.RequestOption) (res *DNSRecord, err error) {
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
func (r *RecordService) Update(ctx context.Context, dnsRecordID string, params RecordUpdateParams, opts ...option.RequestOption) (res *DNSRecord, err error) {
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
func (r *RecordService) List(ctx context.Context, params RecordListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[DNSRecord], err error) {
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
func (r *RecordService) ListAutoPaging(ctx context.Context, params RecordListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[DNSRecord] {
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
func (r *RecordService) Edit(ctx context.Context, dnsRecordID string, params RecordEditParams, opts ...option.RequestOption) (res *DNSRecord, err error) {
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
func (r *RecordService) Get(ctx context.Context, dnsRecordID string, query RecordGetParams, opts ...option.RequestOption) (res *DNSRecord, err error) {
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

type DNSRecord struct {
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Record type.
	Type DNSRecordType `json:"type,required"`
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
	Meta shared.UnnamedSchemaRef162 `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool        `json:"proxiable"`
	Tags      interface{} `json:"tags,required"`
	TTL       interface{} `json:"ttl,required"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string      `json:"zone_name" format:"hostname"`
	Data     interface{} `json:"data,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64       `json:"priority"`
	JSON     dnsRecordJSON `json:"-"`
	union    DNSRecordUnion
}

// dnsRecordJSON contains the JSON metadata for the struct [DNSRecord]
type dnsRecordJSON struct {
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

func (r dnsRecordJSON) RawJSON() string {
	return r.raw
}

func (r *DNSRecord) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r DNSRecord) AsUnion() DNSRecordUnion {
	return r.union
}

// Union satisfied by [dns.DNSRecordDNSRecordsARecord],
// [dns.DNSRecordDNSRecordsAAAARecord], [dns.DNSRecordDNSRecordsCAARecord],
// [dns.DNSRecordDNSRecordsCERTRecord], [dns.DNSRecordDNSRecordsCNAMERecord],
// [dns.DNSRecordDNSRecordsDNSKEYRecord], [dns.DNSRecordDNSRecordsDSRecord],
// [dns.DNSRecordDNSRecordsHTTPSRecord], [dns.DNSRecordDNSRecordsLOCRecord],
// [dns.DNSRecordDNSRecordsMXRecord], [dns.DNSRecordDNSRecordsNAPTRRecord],
// [dns.DNSRecordDNSRecordsNSRecord], [dns.DNSRecordDNSRecordsPTRRecord],
// [dns.DNSRecordDNSRecordsSMIMEARecord], [dns.DNSRecordDNSRecordsSRVRecord],
// [dns.DNSRecordDNSRecordsSSHFPRecord], [dns.DNSRecordDNSRecordsSVCBRecord],
// [dns.DNSRecordDNSRecordsTLSARecord], [dns.DNSRecordDNSRecordsTXTRecord] or
// [dns.DNSRecordDNSRecordsURIRecord].
type DNSRecordUnion interface {
	implementsDNSDNSRecord()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordDNSRecordsARecord{}),
			DiscriminatorValue: "A",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordDNSRecordsAAAARecord{}),
			DiscriminatorValue: "AAAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordDNSRecordsCAARecord{}),
			DiscriminatorValue: "CAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordDNSRecordsCERTRecord{}),
			DiscriminatorValue: "CERT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordDNSRecordsCNAMERecord{}),
			DiscriminatorValue: "CNAME",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordDNSRecordsDNSKEYRecord{}),
			DiscriminatorValue: "DNSKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordDNSRecordsDSRecord{}),
			DiscriminatorValue: "DS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordDNSRecordsHTTPSRecord{}),
			DiscriminatorValue: "HTTPS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordDNSRecordsLOCRecord{}),
			DiscriminatorValue: "LOC",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordDNSRecordsMXRecord{}),
			DiscriminatorValue: "MX",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordDNSRecordsNAPTRRecord{}),
			DiscriminatorValue: "NAPTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordDNSRecordsNSRecord{}),
			DiscriminatorValue: "NS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordDNSRecordsPTRRecord{}),
			DiscriminatorValue: "PTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordDNSRecordsSMIMEARecord{}),
			DiscriminatorValue: "SMIMEA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordDNSRecordsSRVRecord{}),
			DiscriminatorValue: "SRV",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordDNSRecordsSSHFPRecord{}),
			DiscriminatorValue: "SSHFP",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordDNSRecordsSVCBRecord{}),
			DiscriminatorValue: "SVCB",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordDNSRecordsTLSARecord{}),
			DiscriminatorValue: "TLSA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordDNSRecordsTXTRecord{}),
			DiscriminatorValue: "TXT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordDNSRecordsURIRecord{}),
			DiscriminatorValue: "URI",
		},
	)
}

type DNSRecordDNSRecordsARecord struct {
	// A valid IPv4 address.
	Content string `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsARecordType `json:"type,required"`
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
	Meta shared.UnnamedSchemaRef162 `json:"meta"`
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
	TTL DNSRecordDNSRecordsARecordTTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsARecordJSON `json:"-"`
}

// dnsRecordDNSRecordsARecordJSON contains the JSON metadata for the struct
// [DNSRecordDNSRecordsARecord]
type dnsRecordDNSRecordsARecordJSON struct {
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

func (r *DNSRecordDNSRecordsARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsARecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordDNSRecordsARecord) implementsDNSDNSRecord() {}

// Record type.
type DNSRecordDNSRecordsARecordType string

const (
	DNSRecordDNSRecordsARecordTypeA DNSRecordDNSRecordsARecordType = "A"
)

func (r DNSRecordDNSRecordsARecordType) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsARecordTypeA:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.DNSRecordDNSRecordsARecordTTLNumber].
type DNSRecordDNSRecordsARecordTTLUnion interface {
	ImplementsDNSDNSRecordDNSRecordsARecordTTLUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsARecordTTLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordDNSRecordsARecordTTLNumber(0)),
		},
	)
}

type DNSRecordDNSRecordsARecordTTLNumber float64

const (
	DNSRecordDNSRecordsARecordTTLNumber1 DNSRecordDNSRecordsARecordTTLNumber = 1
)

func (r DNSRecordDNSRecordsARecordTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsARecordTTLNumber1:
		return true
	}
	return false
}

type DNSRecordDNSRecordsAAAARecord struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsAAAARecordType `json:"type,required"`
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
	Meta shared.UnnamedSchemaRef162 `json:"meta"`
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
	TTL DNSRecordDNSRecordsAAAARecordTTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                            `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsAAAARecordJSON `json:"-"`
}

// dnsRecordDNSRecordsAAAARecordJSON contains the JSON metadata for the struct
// [DNSRecordDNSRecordsAAAARecord]
type dnsRecordDNSRecordsAAAARecordJSON struct {
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

func (r *DNSRecordDNSRecordsAAAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsAAAARecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordDNSRecordsAAAARecord) implementsDNSDNSRecord() {}

// Record type.
type DNSRecordDNSRecordsAAAARecordType string

const (
	DNSRecordDNSRecordsAAAARecordTypeAAAA DNSRecordDNSRecordsAAAARecordType = "AAAA"
)

func (r DNSRecordDNSRecordsAAAARecordType) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsAAAARecordTypeAAAA:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.DNSRecordDNSRecordsAAAARecordTTLNumber].
type DNSRecordDNSRecordsAAAARecordTTLUnion interface {
	ImplementsDNSDNSRecordDNSRecordsAAAARecordTTLUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsAAAARecordTTLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordDNSRecordsAAAARecordTTLNumber(0)),
		},
	)
}

type DNSRecordDNSRecordsAAAARecordTTLNumber float64

const (
	DNSRecordDNSRecordsAAAARecordTTLNumber1 DNSRecordDNSRecordsAAAARecordTTLNumber = 1
)

func (r DNSRecordDNSRecordsAAAARecordTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsAAAARecordTTLNumber1:
		return true
	}
	return false
}

type DNSRecordDNSRecordsCAARecord struct {
	// Components of a CAA record.
	Data DNSRecordDNSRecordsCAARecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsCAARecordType `json:"type,required"`
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
	Meta shared.UnnamedSchemaRef162 `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsCAARecordTTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                           `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsCAARecordJSON `json:"-"`
}

// dnsRecordDNSRecordsCAARecordJSON contains the JSON metadata for the struct
// [DNSRecordDNSRecordsCAARecord]
type dnsRecordDNSRecordsCAARecordJSON struct {
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

func (r *DNSRecordDNSRecordsCAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsCAARecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordDNSRecordsCAARecord) implementsDNSDNSRecord() {}

// Components of a CAA record.
type DNSRecordDNSRecordsCAARecordData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                               `json:"value"`
	JSON  dnsRecordDNSRecordsCAARecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsCAARecordDataJSON contains the JSON metadata for the struct
// [DNSRecordDNSRecordsCAARecordData]
type dnsRecordDNSRecordsCAARecordDataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsCAARecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsCAARecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordDNSRecordsCAARecordType string

const (
	DNSRecordDNSRecordsCAARecordTypeCAA DNSRecordDNSRecordsCAARecordType = "CAA"
)

func (r DNSRecordDNSRecordsCAARecordType) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsCAARecordTypeCAA:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.DNSRecordDNSRecordsCAARecordTTLNumber].
type DNSRecordDNSRecordsCAARecordTTLUnion interface {
	ImplementsDNSDNSRecordDNSRecordsCAARecordTTLUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsCAARecordTTLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordDNSRecordsCAARecordTTLNumber(0)),
		},
	)
}

type DNSRecordDNSRecordsCAARecordTTLNumber float64

const (
	DNSRecordDNSRecordsCAARecordTTLNumber1 DNSRecordDNSRecordsCAARecordTTLNumber = 1
)

func (r DNSRecordDNSRecordsCAARecordTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsCAARecordTTLNumber1:
		return true
	}
	return false
}

type DNSRecordDNSRecordsCERTRecord struct {
	// Components of a CERT record.
	Data DNSRecordDNSRecordsCERTRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsCERTRecordType `json:"type,required"`
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
	Meta shared.UnnamedSchemaRef162 `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsCERTRecordTTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                            `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsCERTRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsCERTRecordJSON contains the JSON metadata for the struct
// [DNSRecordDNSRecordsCERTRecord]
type dnsRecordDNSRecordsCERTRecordJSON struct {
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

func (r *DNSRecordDNSRecordsCERTRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsCERTRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordDNSRecordsCERTRecord) implementsDNSDNSRecord() {}

// Components of a CERT record.
type DNSRecordDNSRecordsCERTRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64                               `json:"type"`
	JSON dnsRecordDNSRecordsCERTRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsCERTRecordDataJSON contains the JSON metadata for the struct
// [DNSRecordDNSRecordsCERTRecordData]
type dnsRecordDNSRecordsCERTRecordDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsCERTRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsCERTRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordDNSRecordsCERTRecordType string

const (
	DNSRecordDNSRecordsCERTRecordTypeCERT DNSRecordDNSRecordsCERTRecordType = "CERT"
)

func (r DNSRecordDNSRecordsCERTRecordType) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsCERTRecordTypeCERT:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.DNSRecordDNSRecordsCERTRecordTTLNumber].
type DNSRecordDNSRecordsCERTRecordTTLUnion interface {
	ImplementsDNSDNSRecordDNSRecordsCERTRecordTTLUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsCERTRecordTTLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordDNSRecordsCERTRecordTTLNumber(0)),
		},
	)
}

type DNSRecordDNSRecordsCERTRecordTTLNumber float64

const (
	DNSRecordDNSRecordsCERTRecordTTLNumber1 DNSRecordDNSRecordsCERTRecordTTLNumber = 1
)

func (r DNSRecordDNSRecordsCERTRecordTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsCERTRecordTTLNumber1:
		return true
	}
	return false
}

type DNSRecordDNSRecordsCNAMERecord struct {
	// A valid hostname. Must not match the record's name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsCNAMERecordType `json:"type,required"`
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
	Meta shared.UnnamedSchemaRef162 `json:"meta"`
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
	TTL DNSRecordDNSRecordsCNAMERecordTTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                             `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsCNAMERecordJSON `json:"-"`
}

// dnsRecordDNSRecordsCNAMERecordJSON contains the JSON metadata for the struct
// [DNSRecordDNSRecordsCNAMERecord]
type dnsRecordDNSRecordsCNAMERecordJSON struct {
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

func (r *DNSRecordDNSRecordsCNAMERecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsCNAMERecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordDNSRecordsCNAMERecord) implementsDNSDNSRecord() {}

// Record type.
type DNSRecordDNSRecordsCNAMERecordType string

const (
	DNSRecordDNSRecordsCNAMERecordTypeCNAME DNSRecordDNSRecordsCNAMERecordType = "CNAME"
)

func (r DNSRecordDNSRecordsCNAMERecordType) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsCNAMERecordTypeCNAME:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.DNSRecordDNSRecordsCNAMERecordTTLNumber].
type DNSRecordDNSRecordsCNAMERecordTTLUnion interface {
	ImplementsDNSDNSRecordDNSRecordsCNAMERecordTTLUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsCNAMERecordTTLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordDNSRecordsCNAMERecordTTLNumber(0)),
		},
	)
}

type DNSRecordDNSRecordsCNAMERecordTTLNumber float64

const (
	DNSRecordDNSRecordsCNAMERecordTTLNumber1 DNSRecordDNSRecordsCNAMERecordTTLNumber = 1
)

func (r DNSRecordDNSRecordsCNAMERecordTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsCNAMERecordTTLNumber1:
		return true
	}
	return false
}

type DNSRecordDNSRecordsDNSKEYRecord struct {
	// Components of a DNSKEY record.
	Data DNSRecordDNSRecordsDNSKEYRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsDNSKEYRecordType `json:"type,required"`
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
	Meta shared.UnnamedSchemaRef162 `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsDNSKEYRecordTTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                              `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsDNSKEYRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsDNSKEYRecordJSON contains the JSON metadata for the struct
// [DNSRecordDNSRecordsDNSKEYRecord]
type dnsRecordDNSRecordsDNSKEYRecordJSON struct {
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

func (r *DNSRecordDNSRecordsDNSKEYRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsDNSKEYRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordDNSRecordsDNSKEYRecord) implementsDNSDNSRecord() {}

// Components of a DNSKEY record.
type DNSRecordDNSRecordsDNSKEYRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                                  `json:"public_key"`
	JSON      dnsRecordDNSRecordsDNSKEYRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsDNSKEYRecordDataJSON contains the JSON metadata for the
// struct [DNSRecordDNSRecordsDNSKEYRecordData]
type dnsRecordDNSRecordsDNSKEYRecordDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsDNSKEYRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsDNSKEYRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordDNSRecordsDNSKEYRecordType string

const (
	DNSRecordDNSRecordsDNSKEYRecordTypeDNSKEY DNSRecordDNSRecordsDNSKEYRecordType = "DNSKEY"
)

func (r DNSRecordDNSRecordsDNSKEYRecordType) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsDNSKEYRecordTypeDNSKEY:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.DNSRecordDNSRecordsDNSKEYRecordTTLNumber].
type DNSRecordDNSRecordsDNSKEYRecordTTLUnion interface {
	ImplementsDNSDNSRecordDNSRecordsDNSKEYRecordTTLUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsDNSKEYRecordTTLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordDNSRecordsDNSKEYRecordTTLNumber(0)),
		},
	)
}

type DNSRecordDNSRecordsDNSKEYRecordTTLNumber float64

const (
	DNSRecordDNSRecordsDNSKEYRecordTTLNumber1 DNSRecordDNSRecordsDNSKEYRecordTTLNumber = 1
)

func (r DNSRecordDNSRecordsDNSKEYRecordTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsDNSKEYRecordTTLNumber1:
		return true
	}
	return false
}

type DNSRecordDNSRecordsDSRecord struct {
	// Components of a DS record.
	Data DNSRecordDNSRecordsDSRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsDSRecordType `json:"type,required"`
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
	Meta shared.UnnamedSchemaRef162 `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsDSRecordTTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsDSRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsDSRecordJSON contains the JSON metadata for the struct
// [DNSRecordDNSRecordsDSRecord]
type dnsRecordDNSRecordsDSRecordJSON struct {
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

func (r *DNSRecordDNSRecordsDSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsDSRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordDNSRecordsDSRecord) implementsDNSDNSRecord() {}

// Components of a DS record.
type DNSRecordDNSRecordsDSRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                             `json:"key_tag"`
	JSON   dnsRecordDNSRecordsDSRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsDSRecordDataJSON contains the JSON metadata for the struct
// [DNSRecordDNSRecordsDSRecordData]
type dnsRecordDNSRecordsDSRecordDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsDSRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsDSRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordDNSRecordsDSRecordType string

const (
	DNSRecordDNSRecordsDSRecordTypeDS DNSRecordDNSRecordsDSRecordType = "DS"
)

func (r DNSRecordDNSRecordsDSRecordType) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsDSRecordTypeDS:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.DNSRecordDNSRecordsDSRecordTTLNumber].
type DNSRecordDNSRecordsDSRecordTTLUnion interface {
	ImplementsDNSDNSRecordDNSRecordsDSRecordTTLUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsDSRecordTTLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordDNSRecordsDSRecordTTLNumber(0)),
		},
	)
}

type DNSRecordDNSRecordsDSRecordTTLNumber float64

const (
	DNSRecordDNSRecordsDSRecordTTLNumber1 DNSRecordDNSRecordsDSRecordTTLNumber = 1
)

func (r DNSRecordDNSRecordsDSRecordTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsDSRecordTTLNumber1:
		return true
	}
	return false
}

type DNSRecordDNSRecordsHTTPSRecord struct {
	// Components of a HTTPS record.
	Data DNSRecordDNSRecordsHTTPSRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsHTTPSRecordType `json:"type,required"`
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
	Meta shared.UnnamedSchemaRef162 `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsHTTPSRecordTTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                             `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsHTTPSRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsHTTPSRecordJSON contains the JSON metadata for the struct
// [DNSRecordDNSRecordsHTTPSRecord]
type dnsRecordDNSRecordsHTTPSRecordJSON struct {
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

func (r *DNSRecordDNSRecordsHTTPSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsHTTPSRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordDNSRecordsHTTPSRecord) implementsDNSDNSRecord() {}

// Components of a HTTPS record.
type DNSRecordDNSRecordsHTTPSRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                 `json:"value"`
	JSON  dnsRecordDNSRecordsHTTPSRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsHTTPSRecordDataJSON contains the JSON metadata for the struct
// [DNSRecordDNSRecordsHTTPSRecordData]
type dnsRecordDNSRecordsHTTPSRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsHTTPSRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsHTTPSRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordDNSRecordsHTTPSRecordType string

const (
	DNSRecordDNSRecordsHTTPSRecordTypeHTTPS DNSRecordDNSRecordsHTTPSRecordType = "HTTPS"
)

func (r DNSRecordDNSRecordsHTTPSRecordType) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsHTTPSRecordTypeHTTPS:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.DNSRecordDNSRecordsHTTPSRecordTTLNumber].
type DNSRecordDNSRecordsHTTPSRecordTTLUnion interface {
	ImplementsDNSDNSRecordDNSRecordsHTTPSRecordTTLUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsHTTPSRecordTTLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordDNSRecordsHTTPSRecordTTLNumber(0)),
		},
	)
}

type DNSRecordDNSRecordsHTTPSRecordTTLNumber float64

const (
	DNSRecordDNSRecordsHTTPSRecordTTLNumber1 DNSRecordDNSRecordsHTTPSRecordTTLNumber = 1
)

func (r DNSRecordDNSRecordsHTTPSRecordTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsHTTPSRecordTTLNumber1:
		return true
	}
	return false
}

type DNSRecordDNSRecordsLOCRecord struct {
	// Components of a LOC record.
	Data DNSRecordDNSRecordsLOCRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsLOCRecordType `json:"type,required"`
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
	Meta shared.UnnamedSchemaRef162 `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsLOCRecordTTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                           `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsLOCRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsLOCRecordJSON contains the JSON metadata for the struct
// [DNSRecordDNSRecordsLOCRecord]
type dnsRecordDNSRecordsLOCRecordJSON struct {
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

func (r *DNSRecordDNSRecordsLOCRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsLOCRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordDNSRecordsLOCRecord) implementsDNSDNSRecord() {}

// Components of a LOC record.
type DNSRecordDNSRecordsLOCRecordData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection DNSRecordDNSRecordsLOCRecordDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection DNSRecordDNSRecordsLOCRecordDataLongDirection `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes float64 `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds float64 `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz float64 `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert float64 `json:"precision_vert"`
	// Size of location in meters.
	Size float64                              `json:"size"`
	JSON dnsRecordDNSRecordsLOCRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsLOCRecordDataJSON contains the JSON metadata for the struct
// [DNSRecordDNSRecordsLOCRecordData]
type dnsRecordDNSRecordsLOCRecordDataJSON struct {
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

func (r *DNSRecordDNSRecordsLOCRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsLOCRecordDataJSON) RawJSON() string {
	return r.raw
}

// Latitude direction.
type DNSRecordDNSRecordsLOCRecordDataLatDirection string

const (
	DNSRecordDNSRecordsLOCRecordDataLatDirectionN DNSRecordDNSRecordsLOCRecordDataLatDirection = "N"
	DNSRecordDNSRecordsLOCRecordDataLatDirectionS DNSRecordDNSRecordsLOCRecordDataLatDirection = "S"
)

func (r DNSRecordDNSRecordsLOCRecordDataLatDirection) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsLOCRecordDataLatDirectionN, DNSRecordDNSRecordsLOCRecordDataLatDirectionS:
		return true
	}
	return false
}

// Longitude direction.
type DNSRecordDNSRecordsLOCRecordDataLongDirection string

const (
	DNSRecordDNSRecordsLOCRecordDataLongDirectionE DNSRecordDNSRecordsLOCRecordDataLongDirection = "E"
	DNSRecordDNSRecordsLOCRecordDataLongDirectionW DNSRecordDNSRecordsLOCRecordDataLongDirection = "W"
)

func (r DNSRecordDNSRecordsLOCRecordDataLongDirection) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsLOCRecordDataLongDirectionE, DNSRecordDNSRecordsLOCRecordDataLongDirectionW:
		return true
	}
	return false
}

// Record type.
type DNSRecordDNSRecordsLOCRecordType string

const (
	DNSRecordDNSRecordsLOCRecordTypeLOC DNSRecordDNSRecordsLOCRecordType = "LOC"
)

func (r DNSRecordDNSRecordsLOCRecordType) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsLOCRecordTypeLOC:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.DNSRecordDNSRecordsLOCRecordTTLNumber].
type DNSRecordDNSRecordsLOCRecordTTLUnion interface {
	ImplementsDNSDNSRecordDNSRecordsLOCRecordTTLUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsLOCRecordTTLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordDNSRecordsLOCRecordTTLNumber(0)),
		},
	)
}

type DNSRecordDNSRecordsLOCRecordTTLNumber float64

const (
	DNSRecordDNSRecordsLOCRecordTTLNumber1 DNSRecordDNSRecordsLOCRecordTTLNumber = 1
)

func (r DNSRecordDNSRecordsLOCRecordTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsLOCRecordTTLNumber1:
		return true
	}
	return false
}

type DNSRecordDNSRecordsMXRecord struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordDNSRecordsMXRecordType `json:"type,required"`
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
	Meta shared.UnnamedSchemaRef162 `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsMXRecordTTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsMXRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsMXRecordJSON contains the JSON metadata for the struct
// [DNSRecordDNSRecordsMXRecord]
type dnsRecordDNSRecordsMXRecordJSON struct {
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

func (r *DNSRecordDNSRecordsMXRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsMXRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordDNSRecordsMXRecord) implementsDNSDNSRecord() {}

// Record type.
type DNSRecordDNSRecordsMXRecordType string

const (
	DNSRecordDNSRecordsMXRecordTypeMX DNSRecordDNSRecordsMXRecordType = "MX"
)

func (r DNSRecordDNSRecordsMXRecordType) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsMXRecordTypeMX:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.DNSRecordDNSRecordsMXRecordTTLNumber].
type DNSRecordDNSRecordsMXRecordTTLUnion interface {
	ImplementsDNSDNSRecordDNSRecordsMXRecordTTLUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsMXRecordTTLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordDNSRecordsMXRecordTTLNumber(0)),
		},
	)
}

type DNSRecordDNSRecordsMXRecordTTLNumber float64

const (
	DNSRecordDNSRecordsMXRecordTTLNumber1 DNSRecordDNSRecordsMXRecordTTLNumber = 1
)

func (r DNSRecordDNSRecordsMXRecordTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsMXRecordTTLNumber1:
		return true
	}
	return false
}

type DNSRecordDNSRecordsNAPTRRecord struct {
	// Components of a NAPTR record.
	Data DNSRecordDNSRecordsNAPTRRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsNAPTRRecordType `json:"type,required"`
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
	Meta shared.UnnamedSchemaRef162 `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsNAPTRRecordTTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                             `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsNAPTRRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsNAPTRRecordJSON contains the JSON metadata for the struct
// [DNSRecordDNSRecordsNAPTRRecord]
type dnsRecordDNSRecordsNAPTRRecordJSON struct {
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

func (r *DNSRecordDNSRecordsNAPTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsNAPTRRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordDNSRecordsNAPTRRecord) implementsDNSDNSRecord() {}

// Components of a NAPTR record.
type DNSRecordDNSRecordsNAPTRRecordData struct {
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
	Service string                                 `json:"service"`
	JSON    dnsRecordDNSRecordsNAPTRRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsNAPTRRecordDataJSON contains the JSON metadata for the struct
// [DNSRecordDNSRecordsNAPTRRecordData]
type dnsRecordDNSRecordsNAPTRRecordDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsNAPTRRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsNAPTRRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordDNSRecordsNAPTRRecordType string

const (
	DNSRecordDNSRecordsNAPTRRecordTypeNAPTR DNSRecordDNSRecordsNAPTRRecordType = "NAPTR"
)

func (r DNSRecordDNSRecordsNAPTRRecordType) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsNAPTRRecordTypeNAPTR:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.DNSRecordDNSRecordsNAPTRRecordTTLNumber].
type DNSRecordDNSRecordsNAPTRRecordTTLUnion interface {
	ImplementsDNSDNSRecordDNSRecordsNAPTRRecordTTLUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsNAPTRRecordTTLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordDNSRecordsNAPTRRecordTTLNumber(0)),
		},
	)
}

type DNSRecordDNSRecordsNAPTRRecordTTLNumber float64

const (
	DNSRecordDNSRecordsNAPTRRecordTTLNumber1 DNSRecordDNSRecordsNAPTRRecordTTLNumber = 1
)

func (r DNSRecordDNSRecordsNAPTRRecordTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsNAPTRRecordTTLNumber1:
		return true
	}
	return false
}

type DNSRecordDNSRecordsNSRecord struct {
	// A valid name server host name.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsNSRecordType `json:"type,required"`
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
	Meta shared.UnnamedSchemaRef162 `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsNSRecordTTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsNSRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsNSRecordJSON contains the JSON metadata for the struct
// [DNSRecordDNSRecordsNSRecord]
type dnsRecordDNSRecordsNSRecordJSON struct {
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

func (r *DNSRecordDNSRecordsNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsNSRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordDNSRecordsNSRecord) implementsDNSDNSRecord() {}

// Record type.
type DNSRecordDNSRecordsNSRecordType string

const (
	DNSRecordDNSRecordsNSRecordTypeNS DNSRecordDNSRecordsNSRecordType = "NS"
)

func (r DNSRecordDNSRecordsNSRecordType) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsNSRecordTypeNS:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.DNSRecordDNSRecordsNSRecordTTLNumber].
type DNSRecordDNSRecordsNSRecordTTLUnion interface {
	ImplementsDNSDNSRecordDNSRecordsNSRecordTTLUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsNSRecordTTLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordDNSRecordsNSRecordTTLNumber(0)),
		},
	)
}

type DNSRecordDNSRecordsNSRecordTTLNumber float64

const (
	DNSRecordDNSRecordsNSRecordTTLNumber1 DNSRecordDNSRecordsNSRecordTTLNumber = 1
)

func (r DNSRecordDNSRecordsNSRecordTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsNSRecordTTLNumber1:
		return true
	}
	return false
}

type DNSRecordDNSRecordsPTRRecord struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsPTRRecordType `json:"type,required"`
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
	Meta shared.UnnamedSchemaRef162 `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsPTRRecordTTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                           `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsPTRRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsPTRRecordJSON contains the JSON metadata for the struct
// [DNSRecordDNSRecordsPTRRecord]
type dnsRecordDNSRecordsPTRRecordJSON struct {
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

func (r *DNSRecordDNSRecordsPTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsPTRRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordDNSRecordsPTRRecord) implementsDNSDNSRecord() {}

// Record type.
type DNSRecordDNSRecordsPTRRecordType string

const (
	DNSRecordDNSRecordsPTRRecordTypePTR DNSRecordDNSRecordsPTRRecordType = "PTR"
)

func (r DNSRecordDNSRecordsPTRRecordType) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsPTRRecordTypePTR:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.DNSRecordDNSRecordsPTRRecordTTLNumber].
type DNSRecordDNSRecordsPTRRecordTTLUnion interface {
	ImplementsDNSDNSRecordDNSRecordsPTRRecordTTLUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsPTRRecordTTLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordDNSRecordsPTRRecordTTLNumber(0)),
		},
	)
}

type DNSRecordDNSRecordsPTRRecordTTLNumber float64

const (
	DNSRecordDNSRecordsPTRRecordTTLNumber1 DNSRecordDNSRecordsPTRRecordTTLNumber = 1
)

func (r DNSRecordDNSRecordsPTRRecordTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsPTRRecordTTLNumber1:
		return true
	}
	return false
}

type DNSRecordDNSRecordsSMIMEARecord struct {
	// Components of a SMIMEA record.
	Data DNSRecordDNSRecordsSMIMEARecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsSMIMEARecordType `json:"type,required"`
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
	Meta shared.UnnamedSchemaRef162 `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsSMIMEARecordTTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                              `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsSMIMEARecordJSON `json:"-"`
}

// dnsRecordDNSRecordsSMIMEARecordJSON contains the JSON metadata for the struct
// [DNSRecordDNSRecordsSMIMEARecord]
type dnsRecordDNSRecordsSMIMEARecordJSON struct {
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

func (r *DNSRecordDNSRecordsSMIMEARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsSMIMEARecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordDNSRecordsSMIMEARecord) implementsDNSDNSRecord() {}

// Components of a SMIMEA record.
type DNSRecordDNSRecordsSMIMEARecordData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                 `json:"usage"`
	JSON  dnsRecordDNSRecordsSMIMEARecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsSMIMEARecordDataJSON contains the JSON metadata for the
// struct [DNSRecordDNSRecordsSMIMEARecordData]
type dnsRecordDNSRecordsSMIMEARecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsSMIMEARecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsSMIMEARecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordDNSRecordsSMIMEARecordType string

const (
	DNSRecordDNSRecordsSMIMEARecordTypeSMIMEA DNSRecordDNSRecordsSMIMEARecordType = "SMIMEA"
)

func (r DNSRecordDNSRecordsSMIMEARecordType) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsSMIMEARecordTypeSMIMEA:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.DNSRecordDNSRecordsSMIMEARecordTTLNumber].
type DNSRecordDNSRecordsSMIMEARecordTTLUnion interface {
	ImplementsDNSDNSRecordDNSRecordsSMIMEARecordTTLUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsSMIMEARecordTTLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordDNSRecordsSMIMEARecordTTLNumber(0)),
		},
	)
}

type DNSRecordDNSRecordsSMIMEARecordTTLNumber float64

const (
	DNSRecordDNSRecordsSMIMEARecordTTLNumber1 DNSRecordDNSRecordsSMIMEARecordTTLNumber = 1
)

func (r DNSRecordDNSRecordsSMIMEARecordTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsSMIMEARecordTTLNumber1:
		return true
	}
	return false
}

type DNSRecordDNSRecordsSRVRecord struct {
	// Components of a SRV record.
	Data DNSRecordDNSRecordsSRVRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsSRVRecordType `json:"type,required"`
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
	Meta shared.UnnamedSchemaRef162 `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsSRVRecordTTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                           `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsSRVRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsSRVRecordJSON contains the JSON metadata for the struct
// [DNSRecordDNSRecordsSRVRecord]
type dnsRecordDNSRecordsSRVRecordJSON struct {
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

func (r *DNSRecordDNSRecordsSRVRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsSRVRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordDNSRecordsSRVRecord) implementsDNSDNSRecord() {}

// Components of a SRV record.
type DNSRecordDNSRecordsSRVRecordData struct {
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
	Weight float64                              `json:"weight"`
	JSON   dnsRecordDNSRecordsSRVRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsSRVRecordDataJSON contains the JSON metadata for the struct
// [DNSRecordDNSRecordsSRVRecordData]
type dnsRecordDNSRecordsSRVRecordDataJSON struct {
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

func (r *DNSRecordDNSRecordsSRVRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsSRVRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordDNSRecordsSRVRecordType string

const (
	DNSRecordDNSRecordsSRVRecordTypeSRV DNSRecordDNSRecordsSRVRecordType = "SRV"
)

func (r DNSRecordDNSRecordsSRVRecordType) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsSRVRecordTypeSRV:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.DNSRecordDNSRecordsSRVRecordTTLNumber].
type DNSRecordDNSRecordsSRVRecordTTLUnion interface {
	ImplementsDNSDNSRecordDNSRecordsSRVRecordTTLUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsSRVRecordTTLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordDNSRecordsSRVRecordTTLNumber(0)),
		},
	)
}

type DNSRecordDNSRecordsSRVRecordTTLNumber float64

const (
	DNSRecordDNSRecordsSRVRecordTTLNumber1 DNSRecordDNSRecordsSRVRecordTTLNumber = 1
)

func (r DNSRecordDNSRecordsSRVRecordTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsSRVRecordTTLNumber1:
		return true
	}
	return false
}

type DNSRecordDNSRecordsSSHFPRecord struct {
	// Components of a SSHFP record.
	Data DNSRecordDNSRecordsSSHFPRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsSSHFPRecordType `json:"type,required"`
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
	Meta shared.UnnamedSchemaRef162 `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsSSHFPRecordTTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                             `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsSSHFPRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsSSHFPRecordJSON contains the JSON metadata for the struct
// [DNSRecordDNSRecordsSSHFPRecord]
type dnsRecordDNSRecordsSSHFPRecordJSON struct {
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

func (r *DNSRecordDNSRecordsSSHFPRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsSSHFPRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordDNSRecordsSSHFPRecord) implementsDNSDNSRecord() {}

// Components of a SSHFP record.
type DNSRecordDNSRecordsSSHFPRecordData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                                `json:"type"`
	JSON dnsRecordDNSRecordsSSHFPRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsSSHFPRecordDataJSON contains the JSON metadata for the struct
// [DNSRecordDNSRecordsSSHFPRecordData]
type dnsRecordDNSRecordsSSHFPRecordDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsSSHFPRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsSSHFPRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordDNSRecordsSSHFPRecordType string

const (
	DNSRecordDNSRecordsSSHFPRecordTypeSSHFP DNSRecordDNSRecordsSSHFPRecordType = "SSHFP"
)

func (r DNSRecordDNSRecordsSSHFPRecordType) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsSSHFPRecordTypeSSHFP:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.DNSRecordDNSRecordsSSHFPRecordTTLNumber].
type DNSRecordDNSRecordsSSHFPRecordTTLUnion interface {
	ImplementsDNSDNSRecordDNSRecordsSSHFPRecordTTLUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsSSHFPRecordTTLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordDNSRecordsSSHFPRecordTTLNumber(0)),
		},
	)
}

type DNSRecordDNSRecordsSSHFPRecordTTLNumber float64

const (
	DNSRecordDNSRecordsSSHFPRecordTTLNumber1 DNSRecordDNSRecordsSSHFPRecordTTLNumber = 1
)

func (r DNSRecordDNSRecordsSSHFPRecordTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsSSHFPRecordTTLNumber1:
		return true
	}
	return false
}

type DNSRecordDNSRecordsSVCBRecord struct {
	// Components of a SVCB record.
	Data DNSRecordDNSRecordsSVCBRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsSVCBRecordType `json:"type,required"`
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
	Meta shared.UnnamedSchemaRef162 `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsSVCBRecordTTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                            `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsSVCBRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsSVCBRecordJSON contains the JSON metadata for the struct
// [DNSRecordDNSRecordsSVCBRecord]
type dnsRecordDNSRecordsSVCBRecordJSON struct {
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

func (r *DNSRecordDNSRecordsSVCBRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsSVCBRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordDNSRecordsSVCBRecord) implementsDNSDNSRecord() {}

// Components of a SVCB record.
type DNSRecordDNSRecordsSVCBRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                `json:"value"`
	JSON  dnsRecordDNSRecordsSVCBRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsSVCBRecordDataJSON contains the JSON metadata for the struct
// [DNSRecordDNSRecordsSVCBRecordData]
type dnsRecordDNSRecordsSVCBRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsSVCBRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsSVCBRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordDNSRecordsSVCBRecordType string

const (
	DNSRecordDNSRecordsSVCBRecordTypeSVCB DNSRecordDNSRecordsSVCBRecordType = "SVCB"
)

func (r DNSRecordDNSRecordsSVCBRecordType) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsSVCBRecordTypeSVCB:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.DNSRecordDNSRecordsSVCBRecordTTLNumber].
type DNSRecordDNSRecordsSVCBRecordTTLUnion interface {
	ImplementsDNSDNSRecordDNSRecordsSVCBRecordTTLUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsSVCBRecordTTLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordDNSRecordsSVCBRecordTTLNumber(0)),
		},
	)
}

type DNSRecordDNSRecordsSVCBRecordTTLNumber float64

const (
	DNSRecordDNSRecordsSVCBRecordTTLNumber1 DNSRecordDNSRecordsSVCBRecordTTLNumber = 1
)

func (r DNSRecordDNSRecordsSVCBRecordTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsSVCBRecordTTLNumber1:
		return true
	}
	return false
}

type DNSRecordDNSRecordsTLSARecord struct {
	// Components of a TLSA record.
	Data DNSRecordDNSRecordsTLSARecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsTLSARecordType `json:"type,required"`
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
	Meta shared.UnnamedSchemaRef162 `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsTLSARecordTTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                            `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsTLSARecordJSON `json:"-"`
}

// dnsRecordDNSRecordsTLSARecordJSON contains the JSON metadata for the struct
// [DNSRecordDNSRecordsTLSARecord]
type dnsRecordDNSRecordsTLSARecordJSON struct {
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

func (r *DNSRecordDNSRecordsTLSARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsTLSARecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordDNSRecordsTLSARecord) implementsDNSDNSRecord() {}

// Components of a TLSA record.
type DNSRecordDNSRecordsTLSARecordData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                               `json:"usage"`
	JSON  dnsRecordDNSRecordsTLSARecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsTLSARecordDataJSON contains the JSON metadata for the struct
// [DNSRecordDNSRecordsTLSARecordData]
type dnsRecordDNSRecordsTLSARecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsTLSARecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsTLSARecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordDNSRecordsTLSARecordType string

const (
	DNSRecordDNSRecordsTLSARecordTypeTLSA DNSRecordDNSRecordsTLSARecordType = "TLSA"
)

func (r DNSRecordDNSRecordsTLSARecordType) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsTLSARecordTypeTLSA:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.DNSRecordDNSRecordsTLSARecordTTLNumber].
type DNSRecordDNSRecordsTLSARecordTTLUnion interface {
	ImplementsDNSDNSRecordDNSRecordsTLSARecordTTLUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsTLSARecordTTLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordDNSRecordsTLSARecordTTLNumber(0)),
		},
	)
}

type DNSRecordDNSRecordsTLSARecordTTLNumber float64

const (
	DNSRecordDNSRecordsTLSARecordTTLNumber1 DNSRecordDNSRecordsTLSARecordTTLNumber = 1
)

func (r DNSRecordDNSRecordsTLSARecordTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsTLSARecordTTLNumber1:
		return true
	}
	return false
}

type DNSRecordDNSRecordsTXTRecord struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsTXTRecordType `json:"type,required"`
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
	Meta shared.UnnamedSchemaRef162 `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsTXTRecordTTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                           `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsTXTRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsTXTRecordJSON contains the JSON metadata for the struct
// [DNSRecordDNSRecordsTXTRecord]
type dnsRecordDNSRecordsTXTRecordJSON struct {
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

func (r *DNSRecordDNSRecordsTXTRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsTXTRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordDNSRecordsTXTRecord) implementsDNSDNSRecord() {}

// Record type.
type DNSRecordDNSRecordsTXTRecordType string

const (
	DNSRecordDNSRecordsTXTRecordTypeTXT DNSRecordDNSRecordsTXTRecordType = "TXT"
)

func (r DNSRecordDNSRecordsTXTRecordType) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsTXTRecordTypeTXT:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.DNSRecordDNSRecordsTXTRecordTTLNumber].
type DNSRecordDNSRecordsTXTRecordTTLUnion interface {
	ImplementsDNSDNSRecordDNSRecordsTXTRecordTTLUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsTXTRecordTTLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordDNSRecordsTXTRecordTTLNumber(0)),
		},
	)
}

type DNSRecordDNSRecordsTXTRecordTTLNumber float64

const (
	DNSRecordDNSRecordsTXTRecordTTLNumber1 DNSRecordDNSRecordsTXTRecordTTLNumber = 1
)

func (r DNSRecordDNSRecordsTXTRecordTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsTXTRecordTTLNumber1:
		return true
	}
	return false
}

type DNSRecordDNSRecordsURIRecord struct {
	// Components of a URI record.
	Data DNSRecordDNSRecordsURIRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordDNSRecordsURIRecordType `json:"type,required"`
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
	Meta shared.UnnamedSchemaRef162 `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsURIRecordTTLUnion `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                           `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsURIRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsURIRecordJSON contains the JSON metadata for the struct
// [DNSRecordDNSRecordsURIRecord]
type dnsRecordDNSRecordsURIRecordJSON struct {
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

func (r *DNSRecordDNSRecordsURIRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsURIRecordJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordDNSRecordsURIRecord) implementsDNSDNSRecord() {}

// Components of a URI record.
type DNSRecordDNSRecordsURIRecordData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64                              `json:"weight"`
	JSON   dnsRecordDNSRecordsURIRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsURIRecordDataJSON contains the JSON metadata for the struct
// [DNSRecordDNSRecordsURIRecordData]
type dnsRecordDNSRecordsURIRecordDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsURIRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSRecordsURIRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordDNSRecordsURIRecordType string

const (
	DNSRecordDNSRecordsURIRecordTypeURI DNSRecordDNSRecordsURIRecordType = "URI"
)

func (r DNSRecordDNSRecordsURIRecordType) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsURIRecordTypeURI:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.DNSRecordDNSRecordsURIRecordTTLNumber].
type DNSRecordDNSRecordsURIRecordTTLUnion interface {
	ImplementsDNSDNSRecordDNSRecordsURIRecordTTLUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsURIRecordTTLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordDNSRecordsURIRecordTTLNumber(0)),
		},
	)
}

type DNSRecordDNSRecordsURIRecordTTLNumber float64

const (
	DNSRecordDNSRecordsURIRecordTTLNumber1 DNSRecordDNSRecordsURIRecordTTLNumber = 1
)

func (r DNSRecordDNSRecordsURIRecordTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordDNSRecordsURIRecordTTLNumber1:
		return true
	}
	return false
}

// Record type.
type DNSRecordType string

const (
	DNSRecordTypeA      DNSRecordType = "A"
	DNSRecordTypeAAAA   DNSRecordType = "AAAA"
	DNSRecordTypeCAA    DNSRecordType = "CAA"
	DNSRecordTypeCERT   DNSRecordType = "CERT"
	DNSRecordTypeCNAME  DNSRecordType = "CNAME"
	DNSRecordTypeDNSKEY DNSRecordType = "DNSKEY"
	DNSRecordTypeDS     DNSRecordType = "DS"
	DNSRecordTypeHTTPS  DNSRecordType = "HTTPS"
	DNSRecordTypeLOC    DNSRecordType = "LOC"
	DNSRecordTypeMX     DNSRecordType = "MX"
	DNSRecordTypeNAPTR  DNSRecordType = "NAPTR"
	DNSRecordTypeNS     DNSRecordType = "NS"
	DNSRecordTypePTR    DNSRecordType = "PTR"
	DNSRecordTypeSMIMEA DNSRecordType = "SMIMEA"
	DNSRecordTypeSRV    DNSRecordType = "SRV"
	DNSRecordTypeSSHFP  DNSRecordType = "SSHFP"
	DNSRecordTypeSVCB   DNSRecordType = "SVCB"
	DNSRecordTypeTLSA   DNSRecordType = "TLSA"
	DNSRecordTypeTXT    DNSRecordType = "TXT"
	DNSRecordTypeURI    DNSRecordType = "URI"
)

func (r DNSRecordType) IsKnown() bool {
	switch r {
	case DNSRecordTypeA, DNSRecordTypeAAAA, DNSRecordTypeCAA, DNSRecordTypeCERT, DNSRecordTypeCNAME, DNSRecordTypeDNSKEY, DNSRecordTypeDS, DNSRecordTypeHTTPS, DNSRecordTypeLOC, DNSRecordTypeMX, DNSRecordTypeNAPTR, DNSRecordTypeNS, DNSRecordTypePTR, DNSRecordTypeSMIMEA, DNSRecordTypeSRV, DNSRecordTypeSSHFP, DNSRecordTypeSVCB, DNSRecordTypeTLSA, DNSRecordTypeTXT, DNSRecordTypeURI:
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
// [RecordNewParamsDNSRecordsARecord], [RecordNewParamsDNSRecordsAAAARecord],
// [RecordNewParamsDNSRecordsCAARecord], [RecordNewParamsDNSRecordsCERTRecord],
// [RecordNewParamsDNSRecordsCNAMERecord], [RecordNewParamsDNSRecordsDNSKEYRecord],
// [RecordNewParamsDNSRecordsDSRecord], [RecordNewParamsDNSRecordsHTTPSRecord],
// [RecordNewParamsDNSRecordsLOCRecord], [RecordNewParamsDNSRecordsMXRecord],
// [RecordNewParamsDNSRecordsNAPTRRecord], [RecordNewParamsDNSRecordsNSRecord],
// [RecordNewParamsDNSRecordsPTRRecord], [RecordNewParamsDNSRecordsSMIMEARecord],
// [RecordNewParamsDNSRecordsSRVRecord], [RecordNewParamsDNSRecordsSSHFPRecord],
// [RecordNewParamsDNSRecordsSVCBRecord], [RecordNewParamsDNSRecordsTLSARecord],
// [RecordNewParamsDNSRecordsTXTRecord], [RecordNewParamsDNSRecordsURIRecord].
type RecordNewParams interface {
	ImplementsRecordNewParams()

	getZoneID() param.Field[string]
}

type RecordNewParamsDNSRecordsARecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A valid IPv4 address.
	Content param.Field[string] `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsDNSRecordsARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordNewParamsDNSRecordsARecordTTLUnion] `json:"ttl"`
}

func (r RecordNewParamsDNSRecordsARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsDNSRecordsARecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsDNSRecordsARecord) ImplementsRecordNewParams() {

}

// Record type.
type RecordNewParamsDNSRecordsARecordType string

const (
	RecordNewParamsDNSRecordsARecordTypeA RecordNewParamsDNSRecordsARecordType = "A"
)

func (r RecordNewParamsDNSRecordsARecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsARecordTypeA:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordNewParamsDNSRecordsARecordTTLNumber].
type RecordNewParamsDNSRecordsARecordTTLUnion interface {
	ImplementsDNSRecordNewParamsDNSRecordsARecordTTLUnion()
}

type RecordNewParamsDNSRecordsARecordTTLNumber float64

const (
	RecordNewParamsDNSRecordsARecordTTLNumber1 RecordNewParamsDNSRecordsARecordTTLNumber = 1
)

func (r RecordNewParamsDNSRecordsARecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsARecordTTLNumber1:
		return true
	}
	return false
}

type RecordNewParamsDNSRecordsAAAARecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A valid IPv6 address.
	Content param.Field[string] `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsDNSRecordsAAAARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordNewParamsDNSRecordsAAAARecordTTLUnion] `json:"ttl"`
}

func (r RecordNewParamsDNSRecordsAAAARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsDNSRecordsAAAARecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsDNSRecordsAAAARecord) ImplementsRecordNewParams() {

}

// Record type.
type RecordNewParamsDNSRecordsAAAARecordType string

const (
	RecordNewParamsDNSRecordsAAAARecordTypeAAAA RecordNewParamsDNSRecordsAAAARecordType = "AAAA"
)

func (r RecordNewParamsDNSRecordsAAAARecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsAAAARecordTypeAAAA:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordNewParamsDNSRecordsAAAARecordTTLNumber].
type RecordNewParamsDNSRecordsAAAARecordTTLUnion interface {
	ImplementsDNSRecordNewParamsDNSRecordsAAAARecordTTLUnion()
}

type RecordNewParamsDNSRecordsAAAARecordTTLNumber float64

const (
	RecordNewParamsDNSRecordsAAAARecordTTLNumber1 RecordNewParamsDNSRecordsAAAARecordTTLNumber = 1
)

func (r RecordNewParamsDNSRecordsAAAARecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsAAAARecordTTLNumber1:
		return true
	}
	return false
}

type RecordNewParamsDNSRecordsCAARecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a CAA record.
	Data param.Field[RecordNewParamsDNSRecordsCAARecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsDNSRecordsCAARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordNewParamsDNSRecordsCAARecordTTLUnion] `json:"ttl"`
}

func (r RecordNewParamsDNSRecordsCAARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsDNSRecordsCAARecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsDNSRecordsCAARecord) ImplementsRecordNewParams() {

}

// Components of a CAA record.
type RecordNewParamsDNSRecordsCAARecordData struct {
	// Flags for the CAA record.
	Flags param.Field[float64] `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag param.Field[string] `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value param.Field[string] `json:"value"`
}

func (r RecordNewParamsDNSRecordsCAARecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordNewParamsDNSRecordsCAARecordType string

const (
	RecordNewParamsDNSRecordsCAARecordTypeCAA RecordNewParamsDNSRecordsCAARecordType = "CAA"
)

func (r RecordNewParamsDNSRecordsCAARecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsCAARecordTypeCAA:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordNewParamsDNSRecordsCAARecordTTLNumber].
type RecordNewParamsDNSRecordsCAARecordTTLUnion interface {
	ImplementsDNSRecordNewParamsDNSRecordsCAARecordTTLUnion()
}

type RecordNewParamsDNSRecordsCAARecordTTLNumber float64

const (
	RecordNewParamsDNSRecordsCAARecordTTLNumber1 RecordNewParamsDNSRecordsCAARecordTTLNumber = 1
)

func (r RecordNewParamsDNSRecordsCAARecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsCAARecordTTLNumber1:
		return true
	}
	return false
}

type RecordNewParamsDNSRecordsCERTRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a CERT record.
	Data param.Field[RecordNewParamsDNSRecordsCERTRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsDNSRecordsCERTRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordNewParamsDNSRecordsCERTRecordTTLUnion] `json:"ttl"`
}

func (r RecordNewParamsDNSRecordsCERTRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsDNSRecordsCERTRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsDNSRecordsCERTRecord) ImplementsRecordNewParams() {

}

// Components of a CERT record.
type RecordNewParamsDNSRecordsCERTRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
	// Type.
	Type param.Field[float64] `json:"type"`
}

func (r RecordNewParamsDNSRecordsCERTRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordNewParamsDNSRecordsCERTRecordType string

const (
	RecordNewParamsDNSRecordsCERTRecordTypeCERT RecordNewParamsDNSRecordsCERTRecordType = "CERT"
)

func (r RecordNewParamsDNSRecordsCERTRecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsCERTRecordTypeCERT:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordNewParamsDNSRecordsCERTRecordTTLNumber].
type RecordNewParamsDNSRecordsCERTRecordTTLUnion interface {
	ImplementsDNSRecordNewParamsDNSRecordsCERTRecordTTLUnion()
}

type RecordNewParamsDNSRecordsCERTRecordTTLNumber float64

const (
	RecordNewParamsDNSRecordsCERTRecordTTLNumber1 RecordNewParamsDNSRecordsCERTRecordTTLNumber = 1
)

func (r RecordNewParamsDNSRecordsCERTRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsCERTRecordTTLNumber1:
		return true
	}
	return false
}

type RecordNewParamsDNSRecordsCNAMERecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A valid hostname. Must not match the record's name.
	Content param.Field[interface{}] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsDNSRecordsCNAMERecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordNewParamsDNSRecordsCNAMERecordTTLUnion] `json:"ttl"`
}

func (r RecordNewParamsDNSRecordsCNAMERecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsDNSRecordsCNAMERecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsDNSRecordsCNAMERecord) ImplementsRecordNewParams() {

}

// Record type.
type RecordNewParamsDNSRecordsCNAMERecordType string

const (
	RecordNewParamsDNSRecordsCNAMERecordTypeCNAME RecordNewParamsDNSRecordsCNAMERecordType = "CNAME"
)

func (r RecordNewParamsDNSRecordsCNAMERecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsCNAMERecordTypeCNAME:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordNewParamsDNSRecordsCNAMERecordTTLNumber].
type RecordNewParamsDNSRecordsCNAMERecordTTLUnion interface {
	ImplementsDNSRecordNewParamsDNSRecordsCNAMERecordTTLUnion()
}

type RecordNewParamsDNSRecordsCNAMERecordTTLNumber float64

const (
	RecordNewParamsDNSRecordsCNAMERecordTTLNumber1 RecordNewParamsDNSRecordsCNAMERecordTTLNumber = 1
)

func (r RecordNewParamsDNSRecordsCNAMERecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsCNAMERecordTTLNumber1:
		return true
	}
	return false
}

type RecordNewParamsDNSRecordsDNSKEYRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a DNSKEY record.
	Data param.Field[RecordNewParamsDNSRecordsDNSKEYRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsDNSRecordsDNSKEYRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordNewParamsDNSRecordsDNSKEYRecordTTLUnion] `json:"ttl"`
}

func (r RecordNewParamsDNSRecordsDNSKEYRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsDNSRecordsDNSKEYRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsDNSRecordsDNSKEYRecord) ImplementsRecordNewParams() {

}

// Components of a DNSKEY record.
type RecordNewParamsDNSRecordsDNSKEYRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Flags.
	Flags param.Field[float64] `json:"flags"`
	// Protocol.
	Protocol param.Field[float64] `json:"protocol"`
	// Public Key.
	PublicKey param.Field[string] `json:"public_key"`
}

func (r RecordNewParamsDNSRecordsDNSKEYRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordNewParamsDNSRecordsDNSKEYRecordType string

const (
	RecordNewParamsDNSRecordsDNSKEYRecordTypeDNSKEY RecordNewParamsDNSRecordsDNSKEYRecordType = "DNSKEY"
)

func (r RecordNewParamsDNSRecordsDNSKEYRecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsDNSKEYRecordTypeDNSKEY:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordNewParamsDNSRecordsDNSKEYRecordTTLNumber].
type RecordNewParamsDNSRecordsDNSKEYRecordTTLUnion interface {
	ImplementsDNSRecordNewParamsDNSRecordsDNSKEYRecordTTLUnion()
}

type RecordNewParamsDNSRecordsDNSKEYRecordTTLNumber float64

const (
	RecordNewParamsDNSRecordsDNSKEYRecordTTLNumber1 RecordNewParamsDNSRecordsDNSKEYRecordTTLNumber = 1
)

func (r RecordNewParamsDNSRecordsDNSKEYRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsDNSKEYRecordTTLNumber1:
		return true
	}
	return false
}

type RecordNewParamsDNSRecordsDSRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a DS record.
	Data param.Field[RecordNewParamsDNSRecordsDSRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsDNSRecordsDSRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordNewParamsDNSRecordsDSRecordTTLUnion] `json:"ttl"`
}

func (r RecordNewParamsDNSRecordsDSRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsDNSRecordsDSRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsDNSRecordsDSRecord) ImplementsRecordNewParams() {

}

// Components of a DS record.
type RecordNewParamsDNSRecordsDSRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Digest.
	Digest param.Field[string] `json:"digest"`
	// Digest Type.
	DigestType param.Field[float64] `json:"digest_type"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
}

func (r RecordNewParamsDNSRecordsDSRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordNewParamsDNSRecordsDSRecordType string

const (
	RecordNewParamsDNSRecordsDSRecordTypeDS RecordNewParamsDNSRecordsDSRecordType = "DS"
)

func (r RecordNewParamsDNSRecordsDSRecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsDSRecordTypeDS:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordNewParamsDNSRecordsDSRecordTTLNumber].
type RecordNewParamsDNSRecordsDSRecordTTLUnion interface {
	ImplementsDNSRecordNewParamsDNSRecordsDSRecordTTLUnion()
}

type RecordNewParamsDNSRecordsDSRecordTTLNumber float64

const (
	RecordNewParamsDNSRecordsDSRecordTTLNumber1 RecordNewParamsDNSRecordsDSRecordTTLNumber = 1
)

func (r RecordNewParamsDNSRecordsDSRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsDSRecordTTLNumber1:
		return true
	}
	return false
}

type RecordNewParamsDNSRecordsHTTPSRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a HTTPS record.
	Data param.Field[RecordNewParamsDNSRecordsHTTPSRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsDNSRecordsHTTPSRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordNewParamsDNSRecordsHTTPSRecordTTLUnion] `json:"ttl"`
}

func (r RecordNewParamsDNSRecordsHTTPSRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsDNSRecordsHTTPSRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsDNSRecordsHTTPSRecord) ImplementsRecordNewParams() {

}

// Components of a HTTPS record.
type RecordNewParamsDNSRecordsHTTPSRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r RecordNewParamsDNSRecordsHTTPSRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordNewParamsDNSRecordsHTTPSRecordType string

const (
	RecordNewParamsDNSRecordsHTTPSRecordTypeHTTPS RecordNewParamsDNSRecordsHTTPSRecordType = "HTTPS"
)

func (r RecordNewParamsDNSRecordsHTTPSRecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsHTTPSRecordTypeHTTPS:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordNewParamsDNSRecordsHTTPSRecordTTLNumber].
type RecordNewParamsDNSRecordsHTTPSRecordTTLUnion interface {
	ImplementsDNSRecordNewParamsDNSRecordsHTTPSRecordTTLUnion()
}

type RecordNewParamsDNSRecordsHTTPSRecordTTLNumber float64

const (
	RecordNewParamsDNSRecordsHTTPSRecordTTLNumber1 RecordNewParamsDNSRecordsHTTPSRecordTTLNumber = 1
)

func (r RecordNewParamsDNSRecordsHTTPSRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsHTTPSRecordTTLNumber1:
		return true
	}
	return false
}

type RecordNewParamsDNSRecordsLOCRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a LOC record.
	Data param.Field[RecordNewParamsDNSRecordsLOCRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsDNSRecordsLOCRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordNewParamsDNSRecordsLOCRecordTTLUnion] `json:"ttl"`
}

func (r RecordNewParamsDNSRecordsLOCRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsDNSRecordsLOCRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsDNSRecordsLOCRecord) ImplementsRecordNewParams() {

}

// Components of a LOC record.
type RecordNewParamsDNSRecordsLOCRecordData struct {
	// Altitude of location in meters.
	Altitude param.Field[float64] `json:"altitude"`
	// Degrees of latitude.
	LatDegrees param.Field[float64] `json:"lat_degrees"`
	// Latitude direction.
	LatDirection param.Field[RecordNewParamsDNSRecordsLOCRecordDataLatDirection] `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes param.Field[float64] `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds param.Field[float64] `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees param.Field[float64] `json:"long_degrees"`
	// Longitude direction.
	LongDirection param.Field[RecordNewParamsDNSRecordsLOCRecordDataLongDirection] `json:"long_direction"`
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

func (r RecordNewParamsDNSRecordsLOCRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Latitude direction.
type RecordNewParamsDNSRecordsLOCRecordDataLatDirection string

const (
	RecordNewParamsDNSRecordsLOCRecordDataLatDirectionN RecordNewParamsDNSRecordsLOCRecordDataLatDirection = "N"
	RecordNewParamsDNSRecordsLOCRecordDataLatDirectionS RecordNewParamsDNSRecordsLOCRecordDataLatDirection = "S"
)

func (r RecordNewParamsDNSRecordsLOCRecordDataLatDirection) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsLOCRecordDataLatDirectionN, RecordNewParamsDNSRecordsLOCRecordDataLatDirectionS:
		return true
	}
	return false
}

// Longitude direction.
type RecordNewParamsDNSRecordsLOCRecordDataLongDirection string

const (
	RecordNewParamsDNSRecordsLOCRecordDataLongDirectionE RecordNewParamsDNSRecordsLOCRecordDataLongDirection = "E"
	RecordNewParamsDNSRecordsLOCRecordDataLongDirectionW RecordNewParamsDNSRecordsLOCRecordDataLongDirection = "W"
)

func (r RecordNewParamsDNSRecordsLOCRecordDataLongDirection) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsLOCRecordDataLongDirectionE, RecordNewParamsDNSRecordsLOCRecordDataLongDirectionW:
		return true
	}
	return false
}

// Record type.
type RecordNewParamsDNSRecordsLOCRecordType string

const (
	RecordNewParamsDNSRecordsLOCRecordTypeLOC RecordNewParamsDNSRecordsLOCRecordType = "LOC"
)

func (r RecordNewParamsDNSRecordsLOCRecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsLOCRecordTypeLOC:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordNewParamsDNSRecordsLOCRecordTTLNumber].
type RecordNewParamsDNSRecordsLOCRecordTTLUnion interface {
	ImplementsDNSRecordNewParamsDNSRecordsLOCRecordTTLUnion()
}

type RecordNewParamsDNSRecordsLOCRecordTTLNumber float64

const (
	RecordNewParamsDNSRecordsLOCRecordTTLNumber1 RecordNewParamsDNSRecordsLOCRecordTTLNumber = 1
)

func (r RecordNewParamsDNSRecordsLOCRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsLOCRecordTTLNumber1:
		return true
	}
	return false
}

type RecordNewParamsDNSRecordsMXRecord struct {
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
	Type param.Field[RecordNewParamsDNSRecordsMXRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordNewParamsDNSRecordsMXRecordTTLUnion] `json:"ttl"`
}

func (r RecordNewParamsDNSRecordsMXRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsDNSRecordsMXRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsDNSRecordsMXRecord) ImplementsRecordNewParams() {

}

// Record type.
type RecordNewParamsDNSRecordsMXRecordType string

const (
	RecordNewParamsDNSRecordsMXRecordTypeMX RecordNewParamsDNSRecordsMXRecordType = "MX"
)

func (r RecordNewParamsDNSRecordsMXRecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsMXRecordTypeMX:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordNewParamsDNSRecordsMXRecordTTLNumber].
type RecordNewParamsDNSRecordsMXRecordTTLUnion interface {
	ImplementsDNSRecordNewParamsDNSRecordsMXRecordTTLUnion()
}

type RecordNewParamsDNSRecordsMXRecordTTLNumber float64

const (
	RecordNewParamsDNSRecordsMXRecordTTLNumber1 RecordNewParamsDNSRecordsMXRecordTTLNumber = 1
)

func (r RecordNewParamsDNSRecordsMXRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsMXRecordTTLNumber1:
		return true
	}
	return false
}

type RecordNewParamsDNSRecordsNAPTRRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a NAPTR record.
	Data param.Field[RecordNewParamsDNSRecordsNAPTRRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsDNSRecordsNAPTRRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordNewParamsDNSRecordsNAPTRRecordTTLUnion] `json:"ttl"`
}

func (r RecordNewParamsDNSRecordsNAPTRRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsDNSRecordsNAPTRRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsDNSRecordsNAPTRRecord) ImplementsRecordNewParams() {

}

// Components of a NAPTR record.
type RecordNewParamsDNSRecordsNAPTRRecordData struct {
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

func (r RecordNewParamsDNSRecordsNAPTRRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordNewParamsDNSRecordsNAPTRRecordType string

const (
	RecordNewParamsDNSRecordsNAPTRRecordTypeNAPTR RecordNewParamsDNSRecordsNAPTRRecordType = "NAPTR"
)

func (r RecordNewParamsDNSRecordsNAPTRRecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsNAPTRRecordTypeNAPTR:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordNewParamsDNSRecordsNAPTRRecordTTLNumber].
type RecordNewParamsDNSRecordsNAPTRRecordTTLUnion interface {
	ImplementsDNSRecordNewParamsDNSRecordsNAPTRRecordTTLUnion()
}

type RecordNewParamsDNSRecordsNAPTRRecordTTLNumber float64

const (
	RecordNewParamsDNSRecordsNAPTRRecordTTLNumber1 RecordNewParamsDNSRecordsNAPTRRecordTTLNumber = 1
)

func (r RecordNewParamsDNSRecordsNAPTRRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsNAPTRRecordTTLNumber1:
		return true
	}
	return false
}

type RecordNewParamsDNSRecordsNSRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A valid name server host name.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsDNSRecordsNSRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordNewParamsDNSRecordsNSRecordTTLUnion] `json:"ttl"`
}

func (r RecordNewParamsDNSRecordsNSRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsDNSRecordsNSRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsDNSRecordsNSRecord) ImplementsRecordNewParams() {

}

// Record type.
type RecordNewParamsDNSRecordsNSRecordType string

const (
	RecordNewParamsDNSRecordsNSRecordTypeNS RecordNewParamsDNSRecordsNSRecordType = "NS"
)

func (r RecordNewParamsDNSRecordsNSRecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsNSRecordTypeNS:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordNewParamsDNSRecordsNSRecordTTLNumber].
type RecordNewParamsDNSRecordsNSRecordTTLUnion interface {
	ImplementsDNSRecordNewParamsDNSRecordsNSRecordTTLUnion()
}

type RecordNewParamsDNSRecordsNSRecordTTLNumber float64

const (
	RecordNewParamsDNSRecordsNSRecordTTLNumber1 RecordNewParamsDNSRecordsNSRecordTTLNumber = 1
)

func (r RecordNewParamsDNSRecordsNSRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsNSRecordTTLNumber1:
		return true
	}
	return false
}

type RecordNewParamsDNSRecordsPTRRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Domain name pointing to the address.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsDNSRecordsPTRRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordNewParamsDNSRecordsPTRRecordTTLUnion] `json:"ttl"`
}

func (r RecordNewParamsDNSRecordsPTRRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsDNSRecordsPTRRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsDNSRecordsPTRRecord) ImplementsRecordNewParams() {

}

// Record type.
type RecordNewParamsDNSRecordsPTRRecordType string

const (
	RecordNewParamsDNSRecordsPTRRecordTypePTR RecordNewParamsDNSRecordsPTRRecordType = "PTR"
)

func (r RecordNewParamsDNSRecordsPTRRecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsPTRRecordTypePTR:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordNewParamsDNSRecordsPTRRecordTTLNumber].
type RecordNewParamsDNSRecordsPTRRecordTTLUnion interface {
	ImplementsDNSRecordNewParamsDNSRecordsPTRRecordTTLUnion()
}

type RecordNewParamsDNSRecordsPTRRecordTTLNumber float64

const (
	RecordNewParamsDNSRecordsPTRRecordTTLNumber1 RecordNewParamsDNSRecordsPTRRecordTTLNumber = 1
)

func (r RecordNewParamsDNSRecordsPTRRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsPTRRecordTTLNumber1:
		return true
	}
	return false
}

type RecordNewParamsDNSRecordsSMIMEARecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a SMIMEA record.
	Data param.Field[RecordNewParamsDNSRecordsSMIMEARecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsDNSRecordsSMIMEARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordNewParamsDNSRecordsSMIMEARecordTTLUnion] `json:"ttl"`
}

func (r RecordNewParamsDNSRecordsSMIMEARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsDNSRecordsSMIMEARecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsDNSRecordsSMIMEARecord) ImplementsRecordNewParams() {

}

// Components of a SMIMEA record.
type RecordNewParamsDNSRecordsSMIMEARecordData struct {
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r RecordNewParamsDNSRecordsSMIMEARecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordNewParamsDNSRecordsSMIMEARecordType string

const (
	RecordNewParamsDNSRecordsSMIMEARecordTypeSMIMEA RecordNewParamsDNSRecordsSMIMEARecordType = "SMIMEA"
)

func (r RecordNewParamsDNSRecordsSMIMEARecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsSMIMEARecordTypeSMIMEA:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordNewParamsDNSRecordsSMIMEARecordTTLNumber].
type RecordNewParamsDNSRecordsSMIMEARecordTTLUnion interface {
	ImplementsDNSRecordNewParamsDNSRecordsSMIMEARecordTTLUnion()
}

type RecordNewParamsDNSRecordsSMIMEARecordTTLNumber float64

const (
	RecordNewParamsDNSRecordsSMIMEARecordTTLNumber1 RecordNewParamsDNSRecordsSMIMEARecordTTLNumber = 1
)

func (r RecordNewParamsDNSRecordsSMIMEARecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsSMIMEARecordTTLNumber1:
		return true
	}
	return false
}

type RecordNewParamsDNSRecordsSRVRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a SRV record.
	Data param.Field[RecordNewParamsDNSRecordsSRVRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsDNSRecordsSRVRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordNewParamsDNSRecordsSRVRecordTTLUnion] `json:"ttl"`
}

func (r RecordNewParamsDNSRecordsSRVRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsDNSRecordsSRVRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsDNSRecordsSRVRecord) ImplementsRecordNewParams() {

}

// Components of a SRV record.
type RecordNewParamsDNSRecordsSRVRecordData struct {
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

func (r RecordNewParamsDNSRecordsSRVRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordNewParamsDNSRecordsSRVRecordType string

const (
	RecordNewParamsDNSRecordsSRVRecordTypeSRV RecordNewParamsDNSRecordsSRVRecordType = "SRV"
)

func (r RecordNewParamsDNSRecordsSRVRecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsSRVRecordTypeSRV:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordNewParamsDNSRecordsSRVRecordTTLNumber].
type RecordNewParamsDNSRecordsSRVRecordTTLUnion interface {
	ImplementsDNSRecordNewParamsDNSRecordsSRVRecordTTLUnion()
}

type RecordNewParamsDNSRecordsSRVRecordTTLNumber float64

const (
	RecordNewParamsDNSRecordsSRVRecordTTLNumber1 RecordNewParamsDNSRecordsSRVRecordTTLNumber = 1
)

func (r RecordNewParamsDNSRecordsSRVRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsSRVRecordTTLNumber1:
		return true
	}
	return false
}

type RecordNewParamsDNSRecordsSSHFPRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a SSHFP record.
	Data param.Field[RecordNewParamsDNSRecordsSSHFPRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsDNSRecordsSSHFPRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordNewParamsDNSRecordsSSHFPRecordTTLUnion] `json:"ttl"`
}

func (r RecordNewParamsDNSRecordsSSHFPRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsDNSRecordsSSHFPRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsDNSRecordsSSHFPRecord) ImplementsRecordNewParams() {

}

// Components of a SSHFP record.
type RecordNewParamsDNSRecordsSSHFPRecordData struct {
	// algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// fingerprint.
	Fingerprint param.Field[string] `json:"fingerprint"`
	// type.
	Type param.Field[float64] `json:"type"`
}

func (r RecordNewParamsDNSRecordsSSHFPRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordNewParamsDNSRecordsSSHFPRecordType string

const (
	RecordNewParamsDNSRecordsSSHFPRecordTypeSSHFP RecordNewParamsDNSRecordsSSHFPRecordType = "SSHFP"
)

func (r RecordNewParamsDNSRecordsSSHFPRecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsSSHFPRecordTypeSSHFP:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordNewParamsDNSRecordsSSHFPRecordTTLNumber].
type RecordNewParamsDNSRecordsSSHFPRecordTTLUnion interface {
	ImplementsDNSRecordNewParamsDNSRecordsSSHFPRecordTTLUnion()
}

type RecordNewParamsDNSRecordsSSHFPRecordTTLNumber float64

const (
	RecordNewParamsDNSRecordsSSHFPRecordTTLNumber1 RecordNewParamsDNSRecordsSSHFPRecordTTLNumber = 1
)

func (r RecordNewParamsDNSRecordsSSHFPRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsSSHFPRecordTTLNumber1:
		return true
	}
	return false
}

type RecordNewParamsDNSRecordsSVCBRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a SVCB record.
	Data param.Field[RecordNewParamsDNSRecordsSVCBRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsDNSRecordsSVCBRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordNewParamsDNSRecordsSVCBRecordTTLUnion] `json:"ttl"`
}

func (r RecordNewParamsDNSRecordsSVCBRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsDNSRecordsSVCBRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsDNSRecordsSVCBRecord) ImplementsRecordNewParams() {

}

// Components of a SVCB record.
type RecordNewParamsDNSRecordsSVCBRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r RecordNewParamsDNSRecordsSVCBRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordNewParamsDNSRecordsSVCBRecordType string

const (
	RecordNewParamsDNSRecordsSVCBRecordTypeSVCB RecordNewParamsDNSRecordsSVCBRecordType = "SVCB"
)

func (r RecordNewParamsDNSRecordsSVCBRecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsSVCBRecordTypeSVCB:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordNewParamsDNSRecordsSVCBRecordTTLNumber].
type RecordNewParamsDNSRecordsSVCBRecordTTLUnion interface {
	ImplementsDNSRecordNewParamsDNSRecordsSVCBRecordTTLUnion()
}

type RecordNewParamsDNSRecordsSVCBRecordTTLNumber float64

const (
	RecordNewParamsDNSRecordsSVCBRecordTTLNumber1 RecordNewParamsDNSRecordsSVCBRecordTTLNumber = 1
)

func (r RecordNewParamsDNSRecordsSVCBRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsSVCBRecordTTLNumber1:
		return true
	}
	return false
}

type RecordNewParamsDNSRecordsTLSARecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a TLSA record.
	Data param.Field[RecordNewParamsDNSRecordsTLSARecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsDNSRecordsTLSARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordNewParamsDNSRecordsTLSARecordTTLUnion] `json:"ttl"`
}

func (r RecordNewParamsDNSRecordsTLSARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsDNSRecordsTLSARecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsDNSRecordsTLSARecord) ImplementsRecordNewParams() {

}

// Components of a TLSA record.
type RecordNewParamsDNSRecordsTLSARecordData struct {
	// certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r RecordNewParamsDNSRecordsTLSARecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordNewParamsDNSRecordsTLSARecordType string

const (
	RecordNewParamsDNSRecordsTLSARecordTypeTLSA RecordNewParamsDNSRecordsTLSARecordType = "TLSA"
)

func (r RecordNewParamsDNSRecordsTLSARecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsTLSARecordTypeTLSA:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordNewParamsDNSRecordsTLSARecordTTLNumber].
type RecordNewParamsDNSRecordsTLSARecordTTLUnion interface {
	ImplementsDNSRecordNewParamsDNSRecordsTLSARecordTTLUnion()
}

type RecordNewParamsDNSRecordsTLSARecordTTLNumber float64

const (
	RecordNewParamsDNSRecordsTLSARecordTTLNumber1 RecordNewParamsDNSRecordsTLSARecordTTLNumber = 1
)

func (r RecordNewParamsDNSRecordsTLSARecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsTLSARecordTTLNumber1:
		return true
	}
	return false
}

type RecordNewParamsDNSRecordsTXTRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Text content for the record.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsDNSRecordsTXTRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordNewParamsDNSRecordsTXTRecordTTLUnion] `json:"ttl"`
}

func (r RecordNewParamsDNSRecordsTXTRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsDNSRecordsTXTRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsDNSRecordsTXTRecord) ImplementsRecordNewParams() {

}

// Record type.
type RecordNewParamsDNSRecordsTXTRecordType string

const (
	RecordNewParamsDNSRecordsTXTRecordTypeTXT RecordNewParamsDNSRecordsTXTRecordType = "TXT"
)

func (r RecordNewParamsDNSRecordsTXTRecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsTXTRecordTypeTXT:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordNewParamsDNSRecordsTXTRecordTTLNumber].
type RecordNewParamsDNSRecordsTXTRecordTTLUnion interface {
	ImplementsDNSRecordNewParamsDNSRecordsTXTRecordTTLUnion()
}

type RecordNewParamsDNSRecordsTXTRecordTTLNumber float64

const (
	RecordNewParamsDNSRecordsTXTRecordTTLNumber1 RecordNewParamsDNSRecordsTXTRecordTTLNumber = 1
)

func (r RecordNewParamsDNSRecordsTXTRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsTXTRecordTTLNumber1:
		return true
	}
	return false
}

type RecordNewParamsDNSRecordsURIRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a URI record.
	Data param.Field[RecordNewParamsDNSRecordsURIRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority,required"`
	// Record type.
	Type param.Field[RecordNewParamsDNSRecordsURIRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordNewParamsDNSRecordsURIRecordTTLUnion] `json:"ttl"`
}

func (r RecordNewParamsDNSRecordsURIRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordNewParamsDNSRecordsURIRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordNewParamsDNSRecordsURIRecord) ImplementsRecordNewParams() {

}

// Components of a URI record.
type RecordNewParamsDNSRecordsURIRecordData struct {
	// The record content.
	Content param.Field[string] `json:"content"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r RecordNewParamsDNSRecordsURIRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordNewParamsDNSRecordsURIRecordType string

const (
	RecordNewParamsDNSRecordsURIRecordTypeURI RecordNewParamsDNSRecordsURIRecordType = "URI"
)

func (r RecordNewParamsDNSRecordsURIRecordType) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsURIRecordTypeURI:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordNewParamsDNSRecordsURIRecordTTLNumber].
type RecordNewParamsDNSRecordsURIRecordTTLUnion interface {
	ImplementsDNSRecordNewParamsDNSRecordsURIRecordTTLUnion()
}

type RecordNewParamsDNSRecordsURIRecordTTLNumber float64

const (
	RecordNewParamsDNSRecordsURIRecordTTLNumber1 RecordNewParamsDNSRecordsURIRecordTTLNumber = 1
)

func (r RecordNewParamsDNSRecordsURIRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordNewParamsDNSRecordsURIRecordTTLNumber1:
		return true
	}
	return false
}

type RecordNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   DNSRecord             `json:"result,required"`
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
// [RecordUpdateParamsDNSRecordsARecord], [RecordUpdateParamsDNSRecordsAAAARecord],
// [RecordUpdateParamsDNSRecordsCAARecord],
// [RecordUpdateParamsDNSRecordsCERTRecord],
// [RecordUpdateParamsDNSRecordsCNAMERecord],
// [RecordUpdateParamsDNSRecordsDNSKEYRecord],
// [RecordUpdateParamsDNSRecordsDSRecord],
// [RecordUpdateParamsDNSRecordsHTTPSRecord],
// [RecordUpdateParamsDNSRecordsLOCRecord], [RecordUpdateParamsDNSRecordsMXRecord],
// [RecordUpdateParamsDNSRecordsNAPTRRecord],
// [RecordUpdateParamsDNSRecordsNSRecord], [RecordUpdateParamsDNSRecordsPTRRecord],
// [RecordUpdateParamsDNSRecordsSMIMEARecord],
// [RecordUpdateParamsDNSRecordsSRVRecord],
// [RecordUpdateParamsDNSRecordsSSHFPRecord],
// [RecordUpdateParamsDNSRecordsSVCBRecord],
// [RecordUpdateParamsDNSRecordsTLSARecord],
// [RecordUpdateParamsDNSRecordsTXTRecord],
// [RecordUpdateParamsDNSRecordsURIRecord].
type RecordUpdateParams interface {
	ImplementsRecordUpdateParams()

	getZoneID() param.Field[string]
}

type RecordUpdateParamsDNSRecordsARecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A valid IPv4 address.
	Content param.Field[string] `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsDNSRecordsARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordUpdateParamsDNSRecordsARecordTTLUnion] `json:"ttl"`
}

func (r RecordUpdateParamsDNSRecordsARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsDNSRecordsARecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsDNSRecordsARecord) ImplementsRecordUpdateParams() {

}

// Record type.
type RecordUpdateParamsDNSRecordsARecordType string

const (
	RecordUpdateParamsDNSRecordsARecordTypeA RecordUpdateParamsDNSRecordsARecordType = "A"
)

func (r RecordUpdateParamsDNSRecordsARecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsARecordTypeA:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordUpdateParamsDNSRecordsARecordTTLNumber].
type RecordUpdateParamsDNSRecordsARecordTTLUnion interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsARecordTTLUnion()
}

type RecordUpdateParamsDNSRecordsARecordTTLNumber float64

const (
	RecordUpdateParamsDNSRecordsARecordTTLNumber1 RecordUpdateParamsDNSRecordsARecordTTLNumber = 1
)

func (r RecordUpdateParamsDNSRecordsARecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsARecordTTLNumber1:
		return true
	}
	return false
}

type RecordUpdateParamsDNSRecordsAAAARecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A valid IPv6 address.
	Content param.Field[string] `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsDNSRecordsAAAARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordUpdateParamsDNSRecordsAAAARecordTTLUnion] `json:"ttl"`
}

func (r RecordUpdateParamsDNSRecordsAAAARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsDNSRecordsAAAARecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsDNSRecordsAAAARecord) ImplementsRecordUpdateParams() {

}

// Record type.
type RecordUpdateParamsDNSRecordsAAAARecordType string

const (
	RecordUpdateParamsDNSRecordsAAAARecordTypeAAAA RecordUpdateParamsDNSRecordsAAAARecordType = "AAAA"
)

func (r RecordUpdateParamsDNSRecordsAAAARecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsAAAARecordTypeAAAA:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordUpdateParamsDNSRecordsAAAARecordTTLNumber].
type RecordUpdateParamsDNSRecordsAAAARecordTTLUnion interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsAAAARecordTTLUnion()
}

type RecordUpdateParamsDNSRecordsAAAARecordTTLNumber float64

const (
	RecordUpdateParamsDNSRecordsAAAARecordTTLNumber1 RecordUpdateParamsDNSRecordsAAAARecordTTLNumber = 1
)

func (r RecordUpdateParamsDNSRecordsAAAARecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsAAAARecordTTLNumber1:
		return true
	}
	return false
}

type RecordUpdateParamsDNSRecordsCAARecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a CAA record.
	Data param.Field[RecordUpdateParamsDNSRecordsCAARecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsDNSRecordsCAARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordUpdateParamsDNSRecordsCAARecordTTLUnion] `json:"ttl"`
}

func (r RecordUpdateParamsDNSRecordsCAARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsDNSRecordsCAARecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsDNSRecordsCAARecord) ImplementsRecordUpdateParams() {

}

// Components of a CAA record.
type RecordUpdateParamsDNSRecordsCAARecordData struct {
	// Flags for the CAA record.
	Flags param.Field[float64] `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag param.Field[string] `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value param.Field[string] `json:"value"`
}

func (r RecordUpdateParamsDNSRecordsCAARecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordUpdateParamsDNSRecordsCAARecordType string

const (
	RecordUpdateParamsDNSRecordsCAARecordTypeCAA RecordUpdateParamsDNSRecordsCAARecordType = "CAA"
)

func (r RecordUpdateParamsDNSRecordsCAARecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsCAARecordTypeCAA:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordUpdateParamsDNSRecordsCAARecordTTLNumber].
type RecordUpdateParamsDNSRecordsCAARecordTTLUnion interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsCAARecordTTLUnion()
}

type RecordUpdateParamsDNSRecordsCAARecordTTLNumber float64

const (
	RecordUpdateParamsDNSRecordsCAARecordTTLNumber1 RecordUpdateParamsDNSRecordsCAARecordTTLNumber = 1
)

func (r RecordUpdateParamsDNSRecordsCAARecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsCAARecordTTLNumber1:
		return true
	}
	return false
}

type RecordUpdateParamsDNSRecordsCERTRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a CERT record.
	Data param.Field[RecordUpdateParamsDNSRecordsCERTRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsDNSRecordsCERTRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordUpdateParamsDNSRecordsCERTRecordTTLUnion] `json:"ttl"`
}

func (r RecordUpdateParamsDNSRecordsCERTRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsDNSRecordsCERTRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsDNSRecordsCERTRecord) ImplementsRecordUpdateParams() {

}

// Components of a CERT record.
type RecordUpdateParamsDNSRecordsCERTRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
	// Type.
	Type param.Field[float64] `json:"type"`
}

func (r RecordUpdateParamsDNSRecordsCERTRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordUpdateParamsDNSRecordsCERTRecordType string

const (
	RecordUpdateParamsDNSRecordsCERTRecordTypeCERT RecordUpdateParamsDNSRecordsCERTRecordType = "CERT"
)

func (r RecordUpdateParamsDNSRecordsCERTRecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsCERTRecordTypeCERT:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordUpdateParamsDNSRecordsCERTRecordTTLNumber].
type RecordUpdateParamsDNSRecordsCERTRecordTTLUnion interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsCERTRecordTTLUnion()
}

type RecordUpdateParamsDNSRecordsCERTRecordTTLNumber float64

const (
	RecordUpdateParamsDNSRecordsCERTRecordTTLNumber1 RecordUpdateParamsDNSRecordsCERTRecordTTLNumber = 1
)

func (r RecordUpdateParamsDNSRecordsCERTRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsCERTRecordTTLNumber1:
		return true
	}
	return false
}

type RecordUpdateParamsDNSRecordsCNAMERecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A valid hostname. Must not match the record's name.
	Content param.Field[interface{}] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsDNSRecordsCNAMERecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordUpdateParamsDNSRecordsCNAMERecordTTLUnion] `json:"ttl"`
}

func (r RecordUpdateParamsDNSRecordsCNAMERecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsDNSRecordsCNAMERecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsDNSRecordsCNAMERecord) ImplementsRecordUpdateParams() {

}

// Record type.
type RecordUpdateParamsDNSRecordsCNAMERecordType string

const (
	RecordUpdateParamsDNSRecordsCNAMERecordTypeCNAME RecordUpdateParamsDNSRecordsCNAMERecordType = "CNAME"
)

func (r RecordUpdateParamsDNSRecordsCNAMERecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsCNAMERecordTypeCNAME:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordUpdateParamsDNSRecordsCNAMERecordTTLNumber].
type RecordUpdateParamsDNSRecordsCNAMERecordTTLUnion interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsCNAMERecordTTLUnion()
}

type RecordUpdateParamsDNSRecordsCNAMERecordTTLNumber float64

const (
	RecordUpdateParamsDNSRecordsCNAMERecordTTLNumber1 RecordUpdateParamsDNSRecordsCNAMERecordTTLNumber = 1
)

func (r RecordUpdateParamsDNSRecordsCNAMERecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsCNAMERecordTTLNumber1:
		return true
	}
	return false
}

type RecordUpdateParamsDNSRecordsDNSKEYRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a DNSKEY record.
	Data param.Field[RecordUpdateParamsDNSRecordsDNSKEYRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsDNSRecordsDNSKEYRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordUpdateParamsDNSRecordsDNSKEYRecordTTLUnion] `json:"ttl"`
}

func (r RecordUpdateParamsDNSRecordsDNSKEYRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsDNSRecordsDNSKEYRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsDNSRecordsDNSKEYRecord) ImplementsRecordUpdateParams() {

}

// Components of a DNSKEY record.
type RecordUpdateParamsDNSRecordsDNSKEYRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Flags.
	Flags param.Field[float64] `json:"flags"`
	// Protocol.
	Protocol param.Field[float64] `json:"protocol"`
	// Public Key.
	PublicKey param.Field[string] `json:"public_key"`
}

func (r RecordUpdateParamsDNSRecordsDNSKEYRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordUpdateParamsDNSRecordsDNSKEYRecordType string

const (
	RecordUpdateParamsDNSRecordsDNSKEYRecordTypeDNSKEY RecordUpdateParamsDNSRecordsDNSKEYRecordType = "DNSKEY"
)

func (r RecordUpdateParamsDNSRecordsDNSKEYRecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsDNSKEYRecordTypeDNSKEY:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordUpdateParamsDNSRecordsDNSKEYRecordTTLNumber].
type RecordUpdateParamsDNSRecordsDNSKEYRecordTTLUnion interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsDNSKEYRecordTTLUnion()
}

type RecordUpdateParamsDNSRecordsDNSKEYRecordTTLNumber float64

const (
	RecordUpdateParamsDNSRecordsDNSKEYRecordTTLNumber1 RecordUpdateParamsDNSRecordsDNSKEYRecordTTLNumber = 1
)

func (r RecordUpdateParamsDNSRecordsDNSKEYRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsDNSKEYRecordTTLNumber1:
		return true
	}
	return false
}

type RecordUpdateParamsDNSRecordsDSRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a DS record.
	Data param.Field[RecordUpdateParamsDNSRecordsDSRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsDNSRecordsDSRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordUpdateParamsDNSRecordsDSRecordTTLUnion] `json:"ttl"`
}

func (r RecordUpdateParamsDNSRecordsDSRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsDNSRecordsDSRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsDNSRecordsDSRecord) ImplementsRecordUpdateParams() {

}

// Components of a DS record.
type RecordUpdateParamsDNSRecordsDSRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Digest.
	Digest param.Field[string] `json:"digest"`
	// Digest Type.
	DigestType param.Field[float64] `json:"digest_type"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
}

func (r RecordUpdateParamsDNSRecordsDSRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordUpdateParamsDNSRecordsDSRecordType string

const (
	RecordUpdateParamsDNSRecordsDSRecordTypeDS RecordUpdateParamsDNSRecordsDSRecordType = "DS"
)

func (r RecordUpdateParamsDNSRecordsDSRecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsDSRecordTypeDS:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordUpdateParamsDNSRecordsDSRecordTTLNumber].
type RecordUpdateParamsDNSRecordsDSRecordTTLUnion interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsDSRecordTTLUnion()
}

type RecordUpdateParamsDNSRecordsDSRecordTTLNumber float64

const (
	RecordUpdateParamsDNSRecordsDSRecordTTLNumber1 RecordUpdateParamsDNSRecordsDSRecordTTLNumber = 1
)

func (r RecordUpdateParamsDNSRecordsDSRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsDSRecordTTLNumber1:
		return true
	}
	return false
}

type RecordUpdateParamsDNSRecordsHTTPSRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a HTTPS record.
	Data param.Field[RecordUpdateParamsDNSRecordsHTTPSRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsDNSRecordsHTTPSRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordUpdateParamsDNSRecordsHTTPSRecordTTLUnion] `json:"ttl"`
}

func (r RecordUpdateParamsDNSRecordsHTTPSRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsDNSRecordsHTTPSRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsDNSRecordsHTTPSRecord) ImplementsRecordUpdateParams() {

}

// Components of a HTTPS record.
type RecordUpdateParamsDNSRecordsHTTPSRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r RecordUpdateParamsDNSRecordsHTTPSRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordUpdateParamsDNSRecordsHTTPSRecordType string

const (
	RecordUpdateParamsDNSRecordsHTTPSRecordTypeHTTPS RecordUpdateParamsDNSRecordsHTTPSRecordType = "HTTPS"
)

func (r RecordUpdateParamsDNSRecordsHTTPSRecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsHTTPSRecordTypeHTTPS:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordUpdateParamsDNSRecordsHTTPSRecordTTLNumber].
type RecordUpdateParamsDNSRecordsHTTPSRecordTTLUnion interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsHTTPSRecordTTLUnion()
}

type RecordUpdateParamsDNSRecordsHTTPSRecordTTLNumber float64

const (
	RecordUpdateParamsDNSRecordsHTTPSRecordTTLNumber1 RecordUpdateParamsDNSRecordsHTTPSRecordTTLNumber = 1
)

func (r RecordUpdateParamsDNSRecordsHTTPSRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsHTTPSRecordTTLNumber1:
		return true
	}
	return false
}

type RecordUpdateParamsDNSRecordsLOCRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a LOC record.
	Data param.Field[RecordUpdateParamsDNSRecordsLOCRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsDNSRecordsLOCRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordUpdateParamsDNSRecordsLOCRecordTTLUnion] `json:"ttl"`
}

func (r RecordUpdateParamsDNSRecordsLOCRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsDNSRecordsLOCRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsDNSRecordsLOCRecord) ImplementsRecordUpdateParams() {

}

// Components of a LOC record.
type RecordUpdateParamsDNSRecordsLOCRecordData struct {
	// Altitude of location in meters.
	Altitude param.Field[float64] `json:"altitude"`
	// Degrees of latitude.
	LatDegrees param.Field[float64] `json:"lat_degrees"`
	// Latitude direction.
	LatDirection param.Field[RecordUpdateParamsDNSRecordsLOCRecordDataLatDirection] `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes param.Field[float64] `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds param.Field[float64] `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees param.Field[float64] `json:"long_degrees"`
	// Longitude direction.
	LongDirection param.Field[RecordUpdateParamsDNSRecordsLOCRecordDataLongDirection] `json:"long_direction"`
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

func (r RecordUpdateParamsDNSRecordsLOCRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Latitude direction.
type RecordUpdateParamsDNSRecordsLOCRecordDataLatDirection string

const (
	RecordUpdateParamsDNSRecordsLOCRecordDataLatDirectionN RecordUpdateParamsDNSRecordsLOCRecordDataLatDirection = "N"
	RecordUpdateParamsDNSRecordsLOCRecordDataLatDirectionS RecordUpdateParamsDNSRecordsLOCRecordDataLatDirection = "S"
)

func (r RecordUpdateParamsDNSRecordsLOCRecordDataLatDirection) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsLOCRecordDataLatDirectionN, RecordUpdateParamsDNSRecordsLOCRecordDataLatDirectionS:
		return true
	}
	return false
}

// Longitude direction.
type RecordUpdateParamsDNSRecordsLOCRecordDataLongDirection string

const (
	RecordUpdateParamsDNSRecordsLOCRecordDataLongDirectionE RecordUpdateParamsDNSRecordsLOCRecordDataLongDirection = "E"
	RecordUpdateParamsDNSRecordsLOCRecordDataLongDirectionW RecordUpdateParamsDNSRecordsLOCRecordDataLongDirection = "W"
)

func (r RecordUpdateParamsDNSRecordsLOCRecordDataLongDirection) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsLOCRecordDataLongDirectionE, RecordUpdateParamsDNSRecordsLOCRecordDataLongDirectionW:
		return true
	}
	return false
}

// Record type.
type RecordUpdateParamsDNSRecordsLOCRecordType string

const (
	RecordUpdateParamsDNSRecordsLOCRecordTypeLOC RecordUpdateParamsDNSRecordsLOCRecordType = "LOC"
)

func (r RecordUpdateParamsDNSRecordsLOCRecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsLOCRecordTypeLOC:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordUpdateParamsDNSRecordsLOCRecordTTLNumber].
type RecordUpdateParamsDNSRecordsLOCRecordTTLUnion interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsLOCRecordTTLUnion()
}

type RecordUpdateParamsDNSRecordsLOCRecordTTLNumber float64

const (
	RecordUpdateParamsDNSRecordsLOCRecordTTLNumber1 RecordUpdateParamsDNSRecordsLOCRecordTTLNumber = 1
)

func (r RecordUpdateParamsDNSRecordsLOCRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsLOCRecordTTLNumber1:
		return true
	}
	return false
}

type RecordUpdateParamsDNSRecordsMXRecord struct {
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
	Type param.Field[RecordUpdateParamsDNSRecordsMXRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordUpdateParamsDNSRecordsMXRecordTTLUnion] `json:"ttl"`
}

func (r RecordUpdateParamsDNSRecordsMXRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsDNSRecordsMXRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsDNSRecordsMXRecord) ImplementsRecordUpdateParams() {

}

// Record type.
type RecordUpdateParamsDNSRecordsMXRecordType string

const (
	RecordUpdateParamsDNSRecordsMXRecordTypeMX RecordUpdateParamsDNSRecordsMXRecordType = "MX"
)

func (r RecordUpdateParamsDNSRecordsMXRecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsMXRecordTypeMX:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordUpdateParamsDNSRecordsMXRecordTTLNumber].
type RecordUpdateParamsDNSRecordsMXRecordTTLUnion interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsMXRecordTTLUnion()
}

type RecordUpdateParamsDNSRecordsMXRecordTTLNumber float64

const (
	RecordUpdateParamsDNSRecordsMXRecordTTLNumber1 RecordUpdateParamsDNSRecordsMXRecordTTLNumber = 1
)

func (r RecordUpdateParamsDNSRecordsMXRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsMXRecordTTLNumber1:
		return true
	}
	return false
}

type RecordUpdateParamsDNSRecordsNAPTRRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a NAPTR record.
	Data param.Field[RecordUpdateParamsDNSRecordsNAPTRRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsDNSRecordsNAPTRRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordUpdateParamsDNSRecordsNAPTRRecordTTLUnion] `json:"ttl"`
}

func (r RecordUpdateParamsDNSRecordsNAPTRRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsDNSRecordsNAPTRRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsDNSRecordsNAPTRRecord) ImplementsRecordUpdateParams() {

}

// Components of a NAPTR record.
type RecordUpdateParamsDNSRecordsNAPTRRecordData struct {
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

func (r RecordUpdateParamsDNSRecordsNAPTRRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordUpdateParamsDNSRecordsNAPTRRecordType string

const (
	RecordUpdateParamsDNSRecordsNAPTRRecordTypeNAPTR RecordUpdateParamsDNSRecordsNAPTRRecordType = "NAPTR"
)

func (r RecordUpdateParamsDNSRecordsNAPTRRecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsNAPTRRecordTypeNAPTR:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordUpdateParamsDNSRecordsNAPTRRecordTTLNumber].
type RecordUpdateParamsDNSRecordsNAPTRRecordTTLUnion interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsNAPTRRecordTTLUnion()
}

type RecordUpdateParamsDNSRecordsNAPTRRecordTTLNumber float64

const (
	RecordUpdateParamsDNSRecordsNAPTRRecordTTLNumber1 RecordUpdateParamsDNSRecordsNAPTRRecordTTLNumber = 1
)

func (r RecordUpdateParamsDNSRecordsNAPTRRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsNAPTRRecordTTLNumber1:
		return true
	}
	return false
}

type RecordUpdateParamsDNSRecordsNSRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A valid name server host name.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsDNSRecordsNSRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordUpdateParamsDNSRecordsNSRecordTTLUnion] `json:"ttl"`
}

func (r RecordUpdateParamsDNSRecordsNSRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsDNSRecordsNSRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsDNSRecordsNSRecord) ImplementsRecordUpdateParams() {

}

// Record type.
type RecordUpdateParamsDNSRecordsNSRecordType string

const (
	RecordUpdateParamsDNSRecordsNSRecordTypeNS RecordUpdateParamsDNSRecordsNSRecordType = "NS"
)

func (r RecordUpdateParamsDNSRecordsNSRecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsNSRecordTypeNS:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordUpdateParamsDNSRecordsNSRecordTTLNumber].
type RecordUpdateParamsDNSRecordsNSRecordTTLUnion interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsNSRecordTTLUnion()
}

type RecordUpdateParamsDNSRecordsNSRecordTTLNumber float64

const (
	RecordUpdateParamsDNSRecordsNSRecordTTLNumber1 RecordUpdateParamsDNSRecordsNSRecordTTLNumber = 1
)

func (r RecordUpdateParamsDNSRecordsNSRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsNSRecordTTLNumber1:
		return true
	}
	return false
}

type RecordUpdateParamsDNSRecordsPTRRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Domain name pointing to the address.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsDNSRecordsPTRRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordUpdateParamsDNSRecordsPTRRecordTTLUnion] `json:"ttl"`
}

func (r RecordUpdateParamsDNSRecordsPTRRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsDNSRecordsPTRRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsDNSRecordsPTRRecord) ImplementsRecordUpdateParams() {

}

// Record type.
type RecordUpdateParamsDNSRecordsPTRRecordType string

const (
	RecordUpdateParamsDNSRecordsPTRRecordTypePTR RecordUpdateParamsDNSRecordsPTRRecordType = "PTR"
)

func (r RecordUpdateParamsDNSRecordsPTRRecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsPTRRecordTypePTR:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordUpdateParamsDNSRecordsPTRRecordTTLNumber].
type RecordUpdateParamsDNSRecordsPTRRecordTTLUnion interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsPTRRecordTTLUnion()
}

type RecordUpdateParamsDNSRecordsPTRRecordTTLNumber float64

const (
	RecordUpdateParamsDNSRecordsPTRRecordTTLNumber1 RecordUpdateParamsDNSRecordsPTRRecordTTLNumber = 1
)

func (r RecordUpdateParamsDNSRecordsPTRRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsPTRRecordTTLNumber1:
		return true
	}
	return false
}

type RecordUpdateParamsDNSRecordsSMIMEARecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a SMIMEA record.
	Data param.Field[RecordUpdateParamsDNSRecordsSMIMEARecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsDNSRecordsSMIMEARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordUpdateParamsDNSRecordsSMIMEARecordTTLUnion] `json:"ttl"`
}

func (r RecordUpdateParamsDNSRecordsSMIMEARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsDNSRecordsSMIMEARecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsDNSRecordsSMIMEARecord) ImplementsRecordUpdateParams() {

}

// Components of a SMIMEA record.
type RecordUpdateParamsDNSRecordsSMIMEARecordData struct {
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r RecordUpdateParamsDNSRecordsSMIMEARecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordUpdateParamsDNSRecordsSMIMEARecordType string

const (
	RecordUpdateParamsDNSRecordsSMIMEARecordTypeSMIMEA RecordUpdateParamsDNSRecordsSMIMEARecordType = "SMIMEA"
)

func (r RecordUpdateParamsDNSRecordsSMIMEARecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsSMIMEARecordTypeSMIMEA:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordUpdateParamsDNSRecordsSMIMEARecordTTLNumber].
type RecordUpdateParamsDNSRecordsSMIMEARecordTTLUnion interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsSMIMEARecordTTLUnion()
}

type RecordUpdateParamsDNSRecordsSMIMEARecordTTLNumber float64

const (
	RecordUpdateParamsDNSRecordsSMIMEARecordTTLNumber1 RecordUpdateParamsDNSRecordsSMIMEARecordTTLNumber = 1
)

func (r RecordUpdateParamsDNSRecordsSMIMEARecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsSMIMEARecordTTLNumber1:
		return true
	}
	return false
}

type RecordUpdateParamsDNSRecordsSRVRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a SRV record.
	Data param.Field[RecordUpdateParamsDNSRecordsSRVRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsDNSRecordsSRVRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordUpdateParamsDNSRecordsSRVRecordTTLUnion] `json:"ttl"`
}

func (r RecordUpdateParamsDNSRecordsSRVRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsDNSRecordsSRVRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsDNSRecordsSRVRecord) ImplementsRecordUpdateParams() {

}

// Components of a SRV record.
type RecordUpdateParamsDNSRecordsSRVRecordData struct {
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

func (r RecordUpdateParamsDNSRecordsSRVRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordUpdateParamsDNSRecordsSRVRecordType string

const (
	RecordUpdateParamsDNSRecordsSRVRecordTypeSRV RecordUpdateParamsDNSRecordsSRVRecordType = "SRV"
)

func (r RecordUpdateParamsDNSRecordsSRVRecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsSRVRecordTypeSRV:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordUpdateParamsDNSRecordsSRVRecordTTLNumber].
type RecordUpdateParamsDNSRecordsSRVRecordTTLUnion interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsSRVRecordTTLUnion()
}

type RecordUpdateParamsDNSRecordsSRVRecordTTLNumber float64

const (
	RecordUpdateParamsDNSRecordsSRVRecordTTLNumber1 RecordUpdateParamsDNSRecordsSRVRecordTTLNumber = 1
)

func (r RecordUpdateParamsDNSRecordsSRVRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsSRVRecordTTLNumber1:
		return true
	}
	return false
}

type RecordUpdateParamsDNSRecordsSSHFPRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a SSHFP record.
	Data param.Field[RecordUpdateParamsDNSRecordsSSHFPRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsDNSRecordsSSHFPRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordUpdateParamsDNSRecordsSSHFPRecordTTLUnion] `json:"ttl"`
}

func (r RecordUpdateParamsDNSRecordsSSHFPRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsDNSRecordsSSHFPRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsDNSRecordsSSHFPRecord) ImplementsRecordUpdateParams() {

}

// Components of a SSHFP record.
type RecordUpdateParamsDNSRecordsSSHFPRecordData struct {
	// algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// fingerprint.
	Fingerprint param.Field[string] `json:"fingerprint"`
	// type.
	Type param.Field[float64] `json:"type"`
}

func (r RecordUpdateParamsDNSRecordsSSHFPRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordUpdateParamsDNSRecordsSSHFPRecordType string

const (
	RecordUpdateParamsDNSRecordsSSHFPRecordTypeSSHFP RecordUpdateParamsDNSRecordsSSHFPRecordType = "SSHFP"
)

func (r RecordUpdateParamsDNSRecordsSSHFPRecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsSSHFPRecordTypeSSHFP:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordUpdateParamsDNSRecordsSSHFPRecordTTLNumber].
type RecordUpdateParamsDNSRecordsSSHFPRecordTTLUnion interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsSSHFPRecordTTLUnion()
}

type RecordUpdateParamsDNSRecordsSSHFPRecordTTLNumber float64

const (
	RecordUpdateParamsDNSRecordsSSHFPRecordTTLNumber1 RecordUpdateParamsDNSRecordsSSHFPRecordTTLNumber = 1
)

func (r RecordUpdateParamsDNSRecordsSSHFPRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsSSHFPRecordTTLNumber1:
		return true
	}
	return false
}

type RecordUpdateParamsDNSRecordsSVCBRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a SVCB record.
	Data param.Field[RecordUpdateParamsDNSRecordsSVCBRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsDNSRecordsSVCBRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordUpdateParamsDNSRecordsSVCBRecordTTLUnion] `json:"ttl"`
}

func (r RecordUpdateParamsDNSRecordsSVCBRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsDNSRecordsSVCBRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsDNSRecordsSVCBRecord) ImplementsRecordUpdateParams() {

}

// Components of a SVCB record.
type RecordUpdateParamsDNSRecordsSVCBRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r RecordUpdateParamsDNSRecordsSVCBRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordUpdateParamsDNSRecordsSVCBRecordType string

const (
	RecordUpdateParamsDNSRecordsSVCBRecordTypeSVCB RecordUpdateParamsDNSRecordsSVCBRecordType = "SVCB"
)

func (r RecordUpdateParamsDNSRecordsSVCBRecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsSVCBRecordTypeSVCB:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordUpdateParamsDNSRecordsSVCBRecordTTLNumber].
type RecordUpdateParamsDNSRecordsSVCBRecordTTLUnion interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsSVCBRecordTTLUnion()
}

type RecordUpdateParamsDNSRecordsSVCBRecordTTLNumber float64

const (
	RecordUpdateParamsDNSRecordsSVCBRecordTTLNumber1 RecordUpdateParamsDNSRecordsSVCBRecordTTLNumber = 1
)

func (r RecordUpdateParamsDNSRecordsSVCBRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsSVCBRecordTTLNumber1:
		return true
	}
	return false
}

type RecordUpdateParamsDNSRecordsTLSARecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a TLSA record.
	Data param.Field[RecordUpdateParamsDNSRecordsTLSARecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsDNSRecordsTLSARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordUpdateParamsDNSRecordsTLSARecordTTLUnion] `json:"ttl"`
}

func (r RecordUpdateParamsDNSRecordsTLSARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsDNSRecordsTLSARecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsDNSRecordsTLSARecord) ImplementsRecordUpdateParams() {

}

// Components of a TLSA record.
type RecordUpdateParamsDNSRecordsTLSARecordData struct {
	// certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r RecordUpdateParamsDNSRecordsTLSARecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordUpdateParamsDNSRecordsTLSARecordType string

const (
	RecordUpdateParamsDNSRecordsTLSARecordTypeTLSA RecordUpdateParamsDNSRecordsTLSARecordType = "TLSA"
)

func (r RecordUpdateParamsDNSRecordsTLSARecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsTLSARecordTypeTLSA:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordUpdateParamsDNSRecordsTLSARecordTTLNumber].
type RecordUpdateParamsDNSRecordsTLSARecordTTLUnion interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsTLSARecordTTLUnion()
}

type RecordUpdateParamsDNSRecordsTLSARecordTTLNumber float64

const (
	RecordUpdateParamsDNSRecordsTLSARecordTTLNumber1 RecordUpdateParamsDNSRecordsTLSARecordTTLNumber = 1
)

func (r RecordUpdateParamsDNSRecordsTLSARecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsTLSARecordTTLNumber1:
		return true
	}
	return false
}

type RecordUpdateParamsDNSRecordsTXTRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Text content for the record.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsDNSRecordsTXTRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordUpdateParamsDNSRecordsTXTRecordTTLUnion] `json:"ttl"`
}

func (r RecordUpdateParamsDNSRecordsTXTRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsDNSRecordsTXTRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsDNSRecordsTXTRecord) ImplementsRecordUpdateParams() {

}

// Record type.
type RecordUpdateParamsDNSRecordsTXTRecordType string

const (
	RecordUpdateParamsDNSRecordsTXTRecordTypeTXT RecordUpdateParamsDNSRecordsTXTRecordType = "TXT"
)

func (r RecordUpdateParamsDNSRecordsTXTRecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsTXTRecordTypeTXT:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordUpdateParamsDNSRecordsTXTRecordTTLNumber].
type RecordUpdateParamsDNSRecordsTXTRecordTTLUnion interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsTXTRecordTTLUnion()
}

type RecordUpdateParamsDNSRecordsTXTRecordTTLNumber float64

const (
	RecordUpdateParamsDNSRecordsTXTRecordTTLNumber1 RecordUpdateParamsDNSRecordsTXTRecordTTLNumber = 1
)

func (r RecordUpdateParamsDNSRecordsTXTRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsTXTRecordTTLNumber1:
		return true
	}
	return false
}

type RecordUpdateParamsDNSRecordsURIRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a URI record.
	Data param.Field[RecordUpdateParamsDNSRecordsURIRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsDNSRecordsURIRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordUpdateParamsDNSRecordsURIRecordTTLUnion] `json:"ttl"`
}

func (r RecordUpdateParamsDNSRecordsURIRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordUpdateParamsDNSRecordsURIRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordUpdateParamsDNSRecordsURIRecord) ImplementsRecordUpdateParams() {

}

// Components of a URI record.
type RecordUpdateParamsDNSRecordsURIRecordData struct {
	// The record content.
	Content param.Field[string] `json:"content"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r RecordUpdateParamsDNSRecordsURIRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordUpdateParamsDNSRecordsURIRecordType string

const (
	RecordUpdateParamsDNSRecordsURIRecordTypeURI RecordUpdateParamsDNSRecordsURIRecordType = "URI"
)

func (r RecordUpdateParamsDNSRecordsURIRecordType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsURIRecordTypeURI:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordUpdateParamsDNSRecordsURIRecordTTLNumber].
type RecordUpdateParamsDNSRecordsURIRecordTTLUnion interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsURIRecordTTLUnion()
}

type RecordUpdateParamsDNSRecordsURIRecordTTLNumber float64

const (
	RecordUpdateParamsDNSRecordsURIRecordTTLNumber1 RecordUpdateParamsDNSRecordsURIRecordTTLNumber = 1
)

func (r RecordUpdateParamsDNSRecordsURIRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDNSRecordsURIRecordTTLNumber1:
		return true
	}
	return false
}

type RecordUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   DNSRecord             `json:"result,required"`
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
// [RecordEditParamsDNSRecordsARecord], [RecordEditParamsDNSRecordsAAAARecord],
// [RecordEditParamsDNSRecordsCAARecord], [RecordEditParamsDNSRecordsCERTRecord],
// [RecordEditParamsDNSRecordsCNAMERecord],
// [RecordEditParamsDNSRecordsDNSKEYRecord], [RecordEditParamsDNSRecordsDSRecord],
// [RecordEditParamsDNSRecordsHTTPSRecord], [RecordEditParamsDNSRecordsLOCRecord],
// [RecordEditParamsDNSRecordsMXRecord], [RecordEditParamsDNSRecordsNAPTRRecord],
// [RecordEditParamsDNSRecordsNSRecord], [RecordEditParamsDNSRecordsPTRRecord],
// [RecordEditParamsDNSRecordsSMIMEARecord], [RecordEditParamsDNSRecordsSRVRecord],
// [RecordEditParamsDNSRecordsSSHFPRecord], [RecordEditParamsDNSRecordsSVCBRecord],
// [RecordEditParamsDNSRecordsTLSARecord], [RecordEditParamsDNSRecordsTXTRecord],
// [RecordEditParamsDNSRecordsURIRecord].
type RecordEditParams interface {
	ImplementsRecordEditParams()

	getZoneID() param.Field[string]
}

type RecordEditParamsDNSRecordsARecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A valid IPv4 address.
	Content param.Field[string] `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsDNSRecordsARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordEditParamsDNSRecordsARecordTTLUnion] `json:"ttl"`
}

func (r RecordEditParamsDNSRecordsARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsDNSRecordsARecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsDNSRecordsARecord) ImplementsRecordEditParams() {

}

// Record type.
type RecordEditParamsDNSRecordsARecordType string

const (
	RecordEditParamsDNSRecordsARecordTypeA RecordEditParamsDNSRecordsARecordType = "A"
)

func (r RecordEditParamsDNSRecordsARecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsARecordTypeA:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordEditParamsDNSRecordsARecordTTLNumber].
type RecordEditParamsDNSRecordsARecordTTLUnion interface {
	ImplementsDNSRecordEditParamsDNSRecordsARecordTTLUnion()
}

type RecordEditParamsDNSRecordsARecordTTLNumber float64

const (
	RecordEditParamsDNSRecordsARecordTTLNumber1 RecordEditParamsDNSRecordsARecordTTLNumber = 1
)

func (r RecordEditParamsDNSRecordsARecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsARecordTTLNumber1:
		return true
	}
	return false
}

type RecordEditParamsDNSRecordsAAAARecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A valid IPv6 address.
	Content param.Field[string] `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsDNSRecordsAAAARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordEditParamsDNSRecordsAAAARecordTTLUnion] `json:"ttl"`
}

func (r RecordEditParamsDNSRecordsAAAARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsDNSRecordsAAAARecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsDNSRecordsAAAARecord) ImplementsRecordEditParams() {

}

// Record type.
type RecordEditParamsDNSRecordsAAAARecordType string

const (
	RecordEditParamsDNSRecordsAAAARecordTypeAAAA RecordEditParamsDNSRecordsAAAARecordType = "AAAA"
)

func (r RecordEditParamsDNSRecordsAAAARecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsAAAARecordTypeAAAA:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordEditParamsDNSRecordsAAAARecordTTLNumber].
type RecordEditParamsDNSRecordsAAAARecordTTLUnion interface {
	ImplementsDNSRecordEditParamsDNSRecordsAAAARecordTTLUnion()
}

type RecordEditParamsDNSRecordsAAAARecordTTLNumber float64

const (
	RecordEditParamsDNSRecordsAAAARecordTTLNumber1 RecordEditParamsDNSRecordsAAAARecordTTLNumber = 1
)

func (r RecordEditParamsDNSRecordsAAAARecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsAAAARecordTTLNumber1:
		return true
	}
	return false
}

type RecordEditParamsDNSRecordsCAARecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a CAA record.
	Data param.Field[RecordEditParamsDNSRecordsCAARecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsDNSRecordsCAARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordEditParamsDNSRecordsCAARecordTTLUnion] `json:"ttl"`
}

func (r RecordEditParamsDNSRecordsCAARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsDNSRecordsCAARecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsDNSRecordsCAARecord) ImplementsRecordEditParams() {

}

// Components of a CAA record.
type RecordEditParamsDNSRecordsCAARecordData struct {
	// Flags for the CAA record.
	Flags param.Field[float64] `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag param.Field[string] `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value param.Field[string] `json:"value"`
}

func (r RecordEditParamsDNSRecordsCAARecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordEditParamsDNSRecordsCAARecordType string

const (
	RecordEditParamsDNSRecordsCAARecordTypeCAA RecordEditParamsDNSRecordsCAARecordType = "CAA"
)

func (r RecordEditParamsDNSRecordsCAARecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsCAARecordTypeCAA:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordEditParamsDNSRecordsCAARecordTTLNumber].
type RecordEditParamsDNSRecordsCAARecordTTLUnion interface {
	ImplementsDNSRecordEditParamsDNSRecordsCAARecordTTLUnion()
}

type RecordEditParamsDNSRecordsCAARecordTTLNumber float64

const (
	RecordEditParamsDNSRecordsCAARecordTTLNumber1 RecordEditParamsDNSRecordsCAARecordTTLNumber = 1
)

func (r RecordEditParamsDNSRecordsCAARecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsCAARecordTTLNumber1:
		return true
	}
	return false
}

type RecordEditParamsDNSRecordsCERTRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a CERT record.
	Data param.Field[RecordEditParamsDNSRecordsCERTRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsDNSRecordsCERTRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordEditParamsDNSRecordsCERTRecordTTLUnion] `json:"ttl"`
}

func (r RecordEditParamsDNSRecordsCERTRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsDNSRecordsCERTRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsDNSRecordsCERTRecord) ImplementsRecordEditParams() {

}

// Components of a CERT record.
type RecordEditParamsDNSRecordsCERTRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
	// Type.
	Type param.Field[float64] `json:"type"`
}

func (r RecordEditParamsDNSRecordsCERTRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordEditParamsDNSRecordsCERTRecordType string

const (
	RecordEditParamsDNSRecordsCERTRecordTypeCERT RecordEditParamsDNSRecordsCERTRecordType = "CERT"
)

func (r RecordEditParamsDNSRecordsCERTRecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsCERTRecordTypeCERT:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordEditParamsDNSRecordsCERTRecordTTLNumber].
type RecordEditParamsDNSRecordsCERTRecordTTLUnion interface {
	ImplementsDNSRecordEditParamsDNSRecordsCERTRecordTTLUnion()
}

type RecordEditParamsDNSRecordsCERTRecordTTLNumber float64

const (
	RecordEditParamsDNSRecordsCERTRecordTTLNumber1 RecordEditParamsDNSRecordsCERTRecordTTLNumber = 1
)

func (r RecordEditParamsDNSRecordsCERTRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsCERTRecordTTLNumber1:
		return true
	}
	return false
}

type RecordEditParamsDNSRecordsCNAMERecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A valid hostname. Must not match the record's name.
	Content param.Field[interface{}] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsDNSRecordsCNAMERecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordEditParamsDNSRecordsCNAMERecordTTLUnion] `json:"ttl"`
}

func (r RecordEditParamsDNSRecordsCNAMERecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsDNSRecordsCNAMERecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsDNSRecordsCNAMERecord) ImplementsRecordEditParams() {

}

// Record type.
type RecordEditParamsDNSRecordsCNAMERecordType string

const (
	RecordEditParamsDNSRecordsCNAMERecordTypeCNAME RecordEditParamsDNSRecordsCNAMERecordType = "CNAME"
)

func (r RecordEditParamsDNSRecordsCNAMERecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsCNAMERecordTypeCNAME:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordEditParamsDNSRecordsCNAMERecordTTLNumber].
type RecordEditParamsDNSRecordsCNAMERecordTTLUnion interface {
	ImplementsDNSRecordEditParamsDNSRecordsCNAMERecordTTLUnion()
}

type RecordEditParamsDNSRecordsCNAMERecordTTLNumber float64

const (
	RecordEditParamsDNSRecordsCNAMERecordTTLNumber1 RecordEditParamsDNSRecordsCNAMERecordTTLNumber = 1
)

func (r RecordEditParamsDNSRecordsCNAMERecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsCNAMERecordTTLNumber1:
		return true
	}
	return false
}

type RecordEditParamsDNSRecordsDNSKEYRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a DNSKEY record.
	Data param.Field[RecordEditParamsDNSRecordsDNSKEYRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsDNSRecordsDNSKEYRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordEditParamsDNSRecordsDNSKEYRecordTTLUnion] `json:"ttl"`
}

func (r RecordEditParamsDNSRecordsDNSKEYRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsDNSRecordsDNSKEYRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsDNSRecordsDNSKEYRecord) ImplementsRecordEditParams() {

}

// Components of a DNSKEY record.
type RecordEditParamsDNSRecordsDNSKEYRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Flags.
	Flags param.Field[float64] `json:"flags"`
	// Protocol.
	Protocol param.Field[float64] `json:"protocol"`
	// Public Key.
	PublicKey param.Field[string] `json:"public_key"`
}

func (r RecordEditParamsDNSRecordsDNSKEYRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordEditParamsDNSRecordsDNSKEYRecordType string

const (
	RecordEditParamsDNSRecordsDNSKEYRecordTypeDNSKEY RecordEditParamsDNSRecordsDNSKEYRecordType = "DNSKEY"
)

func (r RecordEditParamsDNSRecordsDNSKEYRecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsDNSKEYRecordTypeDNSKEY:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordEditParamsDNSRecordsDNSKEYRecordTTLNumber].
type RecordEditParamsDNSRecordsDNSKEYRecordTTLUnion interface {
	ImplementsDNSRecordEditParamsDNSRecordsDNSKEYRecordTTLUnion()
}

type RecordEditParamsDNSRecordsDNSKEYRecordTTLNumber float64

const (
	RecordEditParamsDNSRecordsDNSKEYRecordTTLNumber1 RecordEditParamsDNSRecordsDNSKEYRecordTTLNumber = 1
)

func (r RecordEditParamsDNSRecordsDNSKEYRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsDNSKEYRecordTTLNumber1:
		return true
	}
	return false
}

type RecordEditParamsDNSRecordsDSRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a DS record.
	Data param.Field[RecordEditParamsDNSRecordsDSRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsDNSRecordsDSRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordEditParamsDNSRecordsDSRecordTTLUnion] `json:"ttl"`
}

func (r RecordEditParamsDNSRecordsDSRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsDNSRecordsDSRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsDNSRecordsDSRecord) ImplementsRecordEditParams() {

}

// Components of a DS record.
type RecordEditParamsDNSRecordsDSRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Digest.
	Digest param.Field[string] `json:"digest"`
	// Digest Type.
	DigestType param.Field[float64] `json:"digest_type"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
}

func (r RecordEditParamsDNSRecordsDSRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordEditParamsDNSRecordsDSRecordType string

const (
	RecordEditParamsDNSRecordsDSRecordTypeDS RecordEditParamsDNSRecordsDSRecordType = "DS"
)

func (r RecordEditParamsDNSRecordsDSRecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsDSRecordTypeDS:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordEditParamsDNSRecordsDSRecordTTLNumber].
type RecordEditParamsDNSRecordsDSRecordTTLUnion interface {
	ImplementsDNSRecordEditParamsDNSRecordsDSRecordTTLUnion()
}

type RecordEditParamsDNSRecordsDSRecordTTLNumber float64

const (
	RecordEditParamsDNSRecordsDSRecordTTLNumber1 RecordEditParamsDNSRecordsDSRecordTTLNumber = 1
)

func (r RecordEditParamsDNSRecordsDSRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsDSRecordTTLNumber1:
		return true
	}
	return false
}

type RecordEditParamsDNSRecordsHTTPSRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a HTTPS record.
	Data param.Field[RecordEditParamsDNSRecordsHTTPSRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsDNSRecordsHTTPSRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordEditParamsDNSRecordsHTTPSRecordTTLUnion] `json:"ttl"`
}

func (r RecordEditParamsDNSRecordsHTTPSRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsDNSRecordsHTTPSRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsDNSRecordsHTTPSRecord) ImplementsRecordEditParams() {

}

// Components of a HTTPS record.
type RecordEditParamsDNSRecordsHTTPSRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r RecordEditParamsDNSRecordsHTTPSRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordEditParamsDNSRecordsHTTPSRecordType string

const (
	RecordEditParamsDNSRecordsHTTPSRecordTypeHTTPS RecordEditParamsDNSRecordsHTTPSRecordType = "HTTPS"
)

func (r RecordEditParamsDNSRecordsHTTPSRecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsHTTPSRecordTypeHTTPS:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordEditParamsDNSRecordsHTTPSRecordTTLNumber].
type RecordEditParamsDNSRecordsHTTPSRecordTTLUnion interface {
	ImplementsDNSRecordEditParamsDNSRecordsHTTPSRecordTTLUnion()
}

type RecordEditParamsDNSRecordsHTTPSRecordTTLNumber float64

const (
	RecordEditParamsDNSRecordsHTTPSRecordTTLNumber1 RecordEditParamsDNSRecordsHTTPSRecordTTLNumber = 1
)

func (r RecordEditParamsDNSRecordsHTTPSRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsHTTPSRecordTTLNumber1:
		return true
	}
	return false
}

type RecordEditParamsDNSRecordsLOCRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a LOC record.
	Data param.Field[RecordEditParamsDNSRecordsLOCRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsDNSRecordsLOCRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordEditParamsDNSRecordsLOCRecordTTLUnion] `json:"ttl"`
}

func (r RecordEditParamsDNSRecordsLOCRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsDNSRecordsLOCRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsDNSRecordsLOCRecord) ImplementsRecordEditParams() {

}

// Components of a LOC record.
type RecordEditParamsDNSRecordsLOCRecordData struct {
	// Altitude of location in meters.
	Altitude param.Field[float64] `json:"altitude"`
	// Degrees of latitude.
	LatDegrees param.Field[float64] `json:"lat_degrees"`
	// Latitude direction.
	LatDirection param.Field[RecordEditParamsDNSRecordsLOCRecordDataLatDirection] `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes param.Field[float64] `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds param.Field[float64] `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees param.Field[float64] `json:"long_degrees"`
	// Longitude direction.
	LongDirection param.Field[RecordEditParamsDNSRecordsLOCRecordDataLongDirection] `json:"long_direction"`
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

func (r RecordEditParamsDNSRecordsLOCRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Latitude direction.
type RecordEditParamsDNSRecordsLOCRecordDataLatDirection string

const (
	RecordEditParamsDNSRecordsLOCRecordDataLatDirectionN RecordEditParamsDNSRecordsLOCRecordDataLatDirection = "N"
	RecordEditParamsDNSRecordsLOCRecordDataLatDirectionS RecordEditParamsDNSRecordsLOCRecordDataLatDirection = "S"
)

func (r RecordEditParamsDNSRecordsLOCRecordDataLatDirection) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsLOCRecordDataLatDirectionN, RecordEditParamsDNSRecordsLOCRecordDataLatDirectionS:
		return true
	}
	return false
}

// Longitude direction.
type RecordEditParamsDNSRecordsLOCRecordDataLongDirection string

const (
	RecordEditParamsDNSRecordsLOCRecordDataLongDirectionE RecordEditParamsDNSRecordsLOCRecordDataLongDirection = "E"
	RecordEditParamsDNSRecordsLOCRecordDataLongDirectionW RecordEditParamsDNSRecordsLOCRecordDataLongDirection = "W"
)

func (r RecordEditParamsDNSRecordsLOCRecordDataLongDirection) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsLOCRecordDataLongDirectionE, RecordEditParamsDNSRecordsLOCRecordDataLongDirectionW:
		return true
	}
	return false
}

// Record type.
type RecordEditParamsDNSRecordsLOCRecordType string

const (
	RecordEditParamsDNSRecordsLOCRecordTypeLOC RecordEditParamsDNSRecordsLOCRecordType = "LOC"
)

func (r RecordEditParamsDNSRecordsLOCRecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsLOCRecordTypeLOC:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordEditParamsDNSRecordsLOCRecordTTLNumber].
type RecordEditParamsDNSRecordsLOCRecordTTLUnion interface {
	ImplementsDNSRecordEditParamsDNSRecordsLOCRecordTTLUnion()
}

type RecordEditParamsDNSRecordsLOCRecordTTLNumber float64

const (
	RecordEditParamsDNSRecordsLOCRecordTTLNumber1 RecordEditParamsDNSRecordsLOCRecordTTLNumber = 1
)

func (r RecordEditParamsDNSRecordsLOCRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsLOCRecordTTLNumber1:
		return true
	}
	return false
}

type RecordEditParamsDNSRecordsMXRecord struct {
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
	Type param.Field[RecordEditParamsDNSRecordsMXRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordEditParamsDNSRecordsMXRecordTTLUnion] `json:"ttl"`
}

func (r RecordEditParamsDNSRecordsMXRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsDNSRecordsMXRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsDNSRecordsMXRecord) ImplementsRecordEditParams() {

}

// Record type.
type RecordEditParamsDNSRecordsMXRecordType string

const (
	RecordEditParamsDNSRecordsMXRecordTypeMX RecordEditParamsDNSRecordsMXRecordType = "MX"
)

func (r RecordEditParamsDNSRecordsMXRecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsMXRecordTypeMX:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordEditParamsDNSRecordsMXRecordTTLNumber].
type RecordEditParamsDNSRecordsMXRecordTTLUnion interface {
	ImplementsDNSRecordEditParamsDNSRecordsMXRecordTTLUnion()
}

type RecordEditParamsDNSRecordsMXRecordTTLNumber float64

const (
	RecordEditParamsDNSRecordsMXRecordTTLNumber1 RecordEditParamsDNSRecordsMXRecordTTLNumber = 1
)

func (r RecordEditParamsDNSRecordsMXRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsMXRecordTTLNumber1:
		return true
	}
	return false
}

type RecordEditParamsDNSRecordsNAPTRRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a NAPTR record.
	Data param.Field[RecordEditParamsDNSRecordsNAPTRRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsDNSRecordsNAPTRRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordEditParamsDNSRecordsNAPTRRecordTTLUnion] `json:"ttl"`
}

func (r RecordEditParamsDNSRecordsNAPTRRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsDNSRecordsNAPTRRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsDNSRecordsNAPTRRecord) ImplementsRecordEditParams() {

}

// Components of a NAPTR record.
type RecordEditParamsDNSRecordsNAPTRRecordData struct {
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

func (r RecordEditParamsDNSRecordsNAPTRRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordEditParamsDNSRecordsNAPTRRecordType string

const (
	RecordEditParamsDNSRecordsNAPTRRecordTypeNAPTR RecordEditParamsDNSRecordsNAPTRRecordType = "NAPTR"
)

func (r RecordEditParamsDNSRecordsNAPTRRecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsNAPTRRecordTypeNAPTR:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordEditParamsDNSRecordsNAPTRRecordTTLNumber].
type RecordEditParamsDNSRecordsNAPTRRecordTTLUnion interface {
	ImplementsDNSRecordEditParamsDNSRecordsNAPTRRecordTTLUnion()
}

type RecordEditParamsDNSRecordsNAPTRRecordTTLNumber float64

const (
	RecordEditParamsDNSRecordsNAPTRRecordTTLNumber1 RecordEditParamsDNSRecordsNAPTRRecordTTLNumber = 1
)

func (r RecordEditParamsDNSRecordsNAPTRRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsNAPTRRecordTTLNumber1:
		return true
	}
	return false
}

type RecordEditParamsDNSRecordsNSRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A valid name server host name.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsDNSRecordsNSRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordEditParamsDNSRecordsNSRecordTTLUnion] `json:"ttl"`
}

func (r RecordEditParamsDNSRecordsNSRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsDNSRecordsNSRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsDNSRecordsNSRecord) ImplementsRecordEditParams() {

}

// Record type.
type RecordEditParamsDNSRecordsNSRecordType string

const (
	RecordEditParamsDNSRecordsNSRecordTypeNS RecordEditParamsDNSRecordsNSRecordType = "NS"
)

func (r RecordEditParamsDNSRecordsNSRecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsNSRecordTypeNS:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordEditParamsDNSRecordsNSRecordTTLNumber].
type RecordEditParamsDNSRecordsNSRecordTTLUnion interface {
	ImplementsDNSRecordEditParamsDNSRecordsNSRecordTTLUnion()
}

type RecordEditParamsDNSRecordsNSRecordTTLNumber float64

const (
	RecordEditParamsDNSRecordsNSRecordTTLNumber1 RecordEditParamsDNSRecordsNSRecordTTLNumber = 1
)

func (r RecordEditParamsDNSRecordsNSRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsNSRecordTTLNumber1:
		return true
	}
	return false
}

type RecordEditParamsDNSRecordsPTRRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Domain name pointing to the address.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsDNSRecordsPTRRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordEditParamsDNSRecordsPTRRecordTTLUnion] `json:"ttl"`
}

func (r RecordEditParamsDNSRecordsPTRRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsDNSRecordsPTRRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsDNSRecordsPTRRecord) ImplementsRecordEditParams() {

}

// Record type.
type RecordEditParamsDNSRecordsPTRRecordType string

const (
	RecordEditParamsDNSRecordsPTRRecordTypePTR RecordEditParamsDNSRecordsPTRRecordType = "PTR"
)

func (r RecordEditParamsDNSRecordsPTRRecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsPTRRecordTypePTR:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordEditParamsDNSRecordsPTRRecordTTLNumber].
type RecordEditParamsDNSRecordsPTRRecordTTLUnion interface {
	ImplementsDNSRecordEditParamsDNSRecordsPTRRecordTTLUnion()
}

type RecordEditParamsDNSRecordsPTRRecordTTLNumber float64

const (
	RecordEditParamsDNSRecordsPTRRecordTTLNumber1 RecordEditParamsDNSRecordsPTRRecordTTLNumber = 1
)

func (r RecordEditParamsDNSRecordsPTRRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsPTRRecordTTLNumber1:
		return true
	}
	return false
}

type RecordEditParamsDNSRecordsSMIMEARecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a SMIMEA record.
	Data param.Field[RecordEditParamsDNSRecordsSMIMEARecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsDNSRecordsSMIMEARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordEditParamsDNSRecordsSMIMEARecordTTLUnion] `json:"ttl"`
}

func (r RecordEditParamsDNSRecordsSMIMEARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsDNSRecordsSMIMEARecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsDNSRecordsSMIMEARecord) ImplementsRecordEditParams() {

}

// Components of a SMIMEA record.
type RecordEditParamsDNSRecordsSMIMEARecordData struct {
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r RecordEditParamsDNSRecordsSMIMEARecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordEditParamsDNSRecordsSMIMEARecordType string

const (
	RecordEditParamsDNSRecordsSMIMEARecordTypeSMIMEA RecordEditParamsDNSRecordsSMIMEARecordType = "SMIMEA"
)

func (r RecordEditParamsDNSRecordsSMIMEARecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsSMIMEARecordTypeSMIMEA:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordEditParamsDNSRecordsSMIMEARecordTTLNumber].
type RecordEditParamsDNSRecordsSMIMEARecordTTLUnion interface {
	ImplementsDNSRecordEditParamsDNSRecordsSMIMEARecordTTLUnion()
}

type RecordEditParamsDNSRecordsSMIMEARecordTTLNumber float64

const (
	RecordEditParamsDNSRecordsSMIMEARecordTTLNumber1 RecordEditParamsDNSRecordsSMIMEARecordTTLNumber = 1
)

func (r RecordEditParamsDNSRecordsSMIMEARecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsSMIMEARecordTTLNumber1:
		return true
	}
	return false
}

type RecordEditParamsDNSRecordsSRVRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a SRV record.
	Data param.Field[RecordEditParamsDNSRecordsSRVRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsDNSRecordsSRVRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordEditParamsDNSRecordsSRVRecordTTLUnion] `json:"ttl"`
}

func (r RecordEditParamsDNSRecordsSRVRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsDNSRecordsSRVRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsDNSRecordsSRVRecord) ImplementsRecordEditParams() {

}

// Components of a SRV record.
type RecordEditParamsDNSRecordsSRVRecordData struct {
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

func (r RecordEditParamsDNSRecordsSRVRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordEditParamsDNSRecordsSRVRecordType string

const (
	RecordEditParamsDNSRecordsSRVRecordTypeSRV RecordEditParamsDNSRecordsSRVRecordType = "SRV"
)

func (r RecordEditParamsDNSRecordsSRVRecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsSRVRecordTypeSRV:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordEditParamsDNSRecordsSRVRecordTTLNumber].
type RecordEditParamsDNSRecordsSRVRecordTTLUnion interface {
	ImplementsDNSRecordEditParamsDNSRecordsSRVRecordTTLUnion()
}

type RecordEditParamsDNSRecordsSRVRecordTTLNumber float64

const (
	RecordEditParamsDNSRecordsSRVRecordTTLNumber1 RecordEditParamsDNSRecordsSRVRecordTTLNumber = 1
)

func (r RecordEditParamsDNSRecordsSRVRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsSRVRecordTTLNumber1:
		return true
	}
	return false
}

type RecordEditParamsDNSRecordsSSHFPRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a SSHFP record.
	Data param.Field[RecordEditParamsDNSRecordsSSHFPRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsDNSRecordsSSHFPRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordEditParamsDNSRecordsSSHFPRecordTTLUnion] `json:"ttl"`
}

func (r RecordEditParamsDNSRecordsSSHFPRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsDNSRecordsSSHFPRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsDNSRecordsSSHFPRecord) ImplementsRecordEditParams() {

}

// Components of a SSHFP record.
type RecordEditParamsDNSRecordsSSHFPRecordData struct {
	// algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// fingerprint.
	Fingerprint param.Field[string] `json:"fingerprint"`
	// type.
	Type param.Field[float64] `json:"type"`
}

func (r RecordEditParamsDNSRecordsSSHFPRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordEditParamsDNSRecordsSSHFPRecordType string

const (
	RecordEditParamsDNSRecordsSSHFPRecordTypeSSHFP RecordEditParamsDNSRecordsSSHFPRecordType = "SSHFP"
)

func (r RecordEditParamsDNSRecordsSSHFPRecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsSSHFPRecordTypeSSHFP:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordEditParamsDNSRecordsSSHFPRecordTTLNumber].
type RecordEditParamsDNSRecordsSSHFPRecordTTLUnion interface {
	ImplementsDNSRecordEditParamsDNSRecordsSSHFPRecordTTLUnion()
}

type RecordEditParamsDNSRecordsSSHFPRecordTTLNumber float64

const (
	RecordEditParamsDNSRecordsSSHFPRecordTTLNumber1 RecordEditParamsDNSRecordsSSHFPRecordTTLNumber = 1
)

func (r RecordEditParamsDNSRecordsSSHFPRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsSSHFPRecordTTLNumber1:
		return true
	}
	return false
}

type RecordEditParamsDNSRecordsSVCBRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a SVCB record.
	Data param.Field[RecordEditParamsDNSRecordsSVCBRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsDNSRecordsSVCBRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordEditParamsDNSRecordsSVCBRecordTTLUnion] `json:"ttl"`
}

func (r RecordEditParamsDNSRecordsSVCBRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsDNSRecordsSVCBRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsDNSRecordsSVCBRecord) ImplementsRecordEditParams() {

}

// Components of a SVCB record.
type RecordEditParamsDNSRecordsSVCBRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r RecordEditParamsDNSRecordsSVCBRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordEditParamsDNSRecordsSVCBRecordType string

const (
	RecordEditParamsDNSRecordsSVCBRecordTypeSVCB RecordEditParamsDNSRecordsSVCBRecordType = "SVCB"
)

func (r RecordEditParamsDNSRecordsSVCBRecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsSVCBRecordTypeSVCB:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordEditParamsDNSRecordsSVCBRecordTTLNumber].
type RecordEditParamsDNSRecordsSVCBRecordTTLUnion interface {
	ImplementsDNSRecordEditParamsDNSRecordsSVCBRecordTTLUnion()
}

type RecordEditParamsDNSRecordsSVCBRecordTTLNumber float64

const (
	RecordEditParamsDNSRecordsSVCBRecordTTLNumber1 RecordEditParamsDNSRecordsSVCBRecordTTLNumber = 1
)

func (r RecordEditParamsDNSRecordsSVCBRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsSVCBRecordTTLNumber1:
		return true
	}
	return false
}

type RecordEditParamsDNSRecordsTLSARecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a TLSA record.
	Data param.Field[RecordEditParamsDNSRecordsTLSARecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsDNSRecordsTLSARecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordEditParamsDNSRecordsTLSARecordTTLUnion] `json:"ttl"`
}

func (r RecordEditParamsDNSRecordsTLSARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsDNSRecordsTLSARecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsDNSRecordsTLSARecord) ImplementsRecordEditParams() {

}

// Components of a TLSA record.
type RecordEditParamsDNSRecordsTLSARecordData struct {
	// certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r RecordEditParamsDNSRecordsTLSARecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordEditParamsDNSRecordsTLSARecordType string

const (
	RecordEditParamsDNSRecordsTLSARecordTypeTLSA RecordEditParamsDNSRecordsTLSARecordType = "TLSA"
)

func (r RecordEditParamsDNSRecordsTLSARecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsTLSARecordTypeTLSA:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordEditParamsDNSRecordsTLSARecordTTLNumber].
type RecordEditParamsDNSRecordsTLSARecordTTLUnion interface {
	ImplementsDNSRecordEditParamsDNSRecordsTLSARecordTTLUnion()
}

type RecordEditParamsDNSRecordsTLSARecordTTLNumber float64

const (
	RecordEditParamsDNSRecordsTLSARecordTTLNumber1 RecordEditParamsDNSRecordsTLSARecordTTLNumber = 1
)

func (r RecordEditParamsDNSRecordsTLSARecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsTLSARecordTTLNumber1:
		return true
	}
	return false
}

type RecordEditParamsDNSRecordsTXTRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Text content for the record.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsDNSRecordsTXTRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordEditParamsDNSRecordsTXTRecordTTLUnion] `json:"ttl"`
}

func (r RecordEditParamsDNSRecordsTXTRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsDNSRecordsTXTRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsDNSRecordsTXTRecord) ImplementsRecordEditParams() {

}

// Record type.
type RecordEditParamsDNSRecordsTXTRecordType string

const (
	RecordEditParamsDNSRecordsTXTRecordTypeTXT RecordEditParamsDNSRecordsTXTRecordType = "TXT"
)

func (r RecordEditParamsDNSRecordsTXTRecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsTXTRecordTypeTXT:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordEditParamsDNSRecordsTXTRecordTTLNumber].
type RecordEditParamsDNSRecordsTXTRecordTTLUnion interface {
	ImplementsDNSRecordEditParamsDNSRecordsTXTRecordTTLUnion()
}

type RecordEditParamsDNSRecordsTXTRecordTTLNumber float64

const (
	RecordEditParamsDNSRecordsTXTRecordTTLNumber1 RecordEditParamsDNSRecordsTXTRecordTTLNumber = 1
)

func (r RecordEditParamsDNSRecordsTXTRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsTXTRecordTTLNumber1:
		return true
	}
	return false
}

type RecordEditParamsDNSRecordsURIRecord struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Components of a URI record.
	Data param.Field[RecordEditParamsDNSRecordsURIRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority,required"`
	// Record type.
	Type param.Field[RecordEditParamsDNSRecordsURIRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordEditParamsDNSRecordsURIRecordTTLUnion] `json:"ttl"`
}

func (r RecordEditParamsDNSRecordsURIRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordEditParamsDNSRecordsURIRecord) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RecordEditParamsDNSRecordsURIRecord) ImplementsRecordEditParams() {

}

// Components of a URI record.
type RecordEditParamsDNSRecordsURIRecordData struct {
	// The record content.
	Content param.Field[string] `json:"content"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r RecordEditParamsDNSRecordsURIRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordEditParamsDNSRecordsURIRecordType string

const (
	RecordEditParamsDNSRecordsURIRecordTypeURI RecordEditParamsDNSRecordsURIRecordType = "URI"
)

func (r RecordEditParamsDNSRecordsURIRecordType) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsURIRecordTypeURI:
		return true
	}
	return false
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [dns.RecordEditParamsDNSRecordsURIRecordTTLNumber].
type RecordEditParamsDNSRecordsURIRecordTTLUnion interface {
	ImplementsDNSRecordEditParamsDNSRecordsURIRecordTTLUnion()
}

type RecordEditParamsDNSRecordsURIRecordTTLNumber float64

const (
	RecordEditParamsDNSRecordsURIRecordTTLNumber1 RecordEditParamsDNSRecordsURIRecordTTLNumber = 1
)

func (r RecordEditParamsDNSRecordsURIRecordTTLNumber) IsKnown() bool {
	switch r {
	case RecordEditParamsDNSRecordsURIRecordTTLNumber1:
		return true
	}
	return false
}

type RecordEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   DNSRecord             `json:"result,required"`
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   DNSRecord             `json:"result,required"`
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   RecordImportResponse  `json:"result,required"`
	// Whether the API call was successful
	Success RecordImportResponseEnvelopeSuccess `json:"success,required"`
	Timing  shared.UnnamedSchemaRef62           `json:"timing"`
	JSON    recordImportResponseEnvelopeJSON    `json:"-"`
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   RecordScanResponse    `json:"result,required"`
	// Whether the API call was successful
	Success RecordScanResponseEnvelopeSuccess `json:"success,required"`
	Timing  shared.UnnamedSchemaRef62         `json:"timing"`
	JSON    recordScanResponseEnvelopeJSON    `json:"-"`
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
