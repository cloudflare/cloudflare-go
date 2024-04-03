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

// Union satisfied by [dns.DNSRecordA], [dns.DNSRecordAAAA], [dns.DNSRecordCAA],
// [dns.DNSRecordCERT], [dns.DNSRecordCNAME], [dns.DNSRecordDNSKEY],
// [dns.DNSRecordDS], [dns.DNSRecordHTTPS], [dns.DNSRecordLOC], [dns.DNSRecordMX],
// [dns.DNSRecordNAPTR], [dns.DNSRecordNS], [dns.DNSRecordPTR],
// [dns.DNSRecordSMIMEA], [dns.DNSRecordSRV], [dns.DNSRecordSSHFP],
// [dns.DNSRecordSVCB], [dns.DNSRecordTLSA], [dns.DNSRecordTXT] or
// [dns.DNSRecordURI].
type DNSRecord interface {
	implementsDNSDNSRecord()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecord)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordA{}),
			DiscriminatorValue: "A",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordAAAA{}),
			DiscriminatorValue: "AAAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordCAA{}),
			DiscriminatorValue: "CAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordCERT{}),
			DiscriminatorValue: "CERT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordCNAME{}),
			DiscriminatorValue: "CNAME",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordDNSKEY{}),
			DiscriminatorValue: "DNSKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordDS{}),
			DiscriminatorValue: "DS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordHTTPS{}),
			DiscriminatorValue: "HTTPS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordLOC{}),
			DiscriminatorValue: "LOC",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordMX{}),
			DiscriminatorValue: "MX",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordNAPTR{}),
			DiscriminatorValue: "NAPTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordNS{}),
			DiscriminatorValue: "NS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordPTR{}),
			DiscriminatorValue: "PTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordSMIMEA{}),
			DiscriminatorValue: "SMIMEA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordSRV{}),
			DiscriminatorValue: "SRV",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordSSHFP{}),
			DiscriminatorValue: "SSHFP",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordSVCB{}),
			DiscriminatorValue: "SVCB",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordTLSA{}),
			DiscriminatorValue: "TLSA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordTXT{}),
			DiscriminatorValue: "TXT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordURI{}),
			DiscriminatorValue: "URI",
		},
	)
}

type DNSRecordA struct {
	// A valid IPv4 address.
	Content string `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordAType `json:"type,required"`
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
	Meta DNSRecordAMeta `json:"meta"`
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
	TTL DNSRecordATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordAJSON `json:"-"`
}

// dnsRecordAJSON contains the JSON metadata for the struct [DNSRecordA]
type dnsRecordAJSON struct {
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

func (r *DNSRecordA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordAJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordA) implementsDNSDNSRecord() {}

// Record type.
type DNSRecordAType string

const (
	DNSRecordATypeA DNSRecordAType = "A"
)

func (r DNSRecordAType) IsKnown() bool {
	switch r {
	case DNSRecordATypeA:
		return true
	}
	return false
}

// Extra Cloudflare-specific information about the record.
type DNSRecordAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string             `json:"source"`
	JSON   dnsRecordAMetaJSON `json:"-"`
}

// dnsRecordAMetaJSON contains the JSON metadata for the struct [DNSRecordAMeta]
type dnsRecordAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.DNSRecordATTLNumber].
type DNSRecordATTL interface {
	ImplementsDNSDNSRecordAttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordATTLNumber(0)),
		},
	)
}

type DNSRecordATTLNumber float64

const (
	DNSRecordATTLNumber1 DNSRecordATTLNumber = 1
)

func (r DNSRecordATTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordATTLNumber1:
		return true
	}
	return false
}

type DNSRecordAAAA struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordAAAAType `json:"type,required"`
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
	Meta DNSRecordAAAAMeta `json:"meta"`
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
	TTL DNSRecordAAAATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string            `json:"zone_name" format:"hostname"`
	JSON     dnsRecordAAAAJSON `json:"-"`
}

// dnsRecordAAAAJSON contains the JSON metadata for the struct [DNSRecordAAAA]
type dnsRecordAAAAJSON struct {
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

func (r *DNSRecordAAAA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordAAAAJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordAAAA) implementsDNSDNSRecord() {}

// Record type.
type DNSRecordAAAAType string

const (
	DNSRecordAAAATypeAAAA DNSRecordAAAAType = "AAAA"
)

func (r DNSRecordAAAAType) IsKnown() bool {
	switch r {
	case DNSRecordAAAATypeAAAA:
		return true
	}
	return false
}

// Extra Cloudflare-specific information about the record.
type DNSRecordAAAAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                `json:"source"`
	JSON   dnsRecordAAAAMetaJSON `json:"-"`
}

// dnsRecordAAAAMetaJSON contains the JSON metadata for the struct
// [DNSRecordAAAAMeta]
type dnsRecordAAAAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordAAAAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordAAAAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.DNSRecordAAAATTLNumber].
type DNSRecordAAAATTL interface {
	ImplementsDNSDNSRecordAaaattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordAAAATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordAAAATTLNumber(0)),
		},
	)
}

type DNSRecordAAAATTLNumber float64

const (
	DNSRecordAAAATTLNumber1 DNSRecordAAAATTLNumber = 1
)

func (r DNSRecordAAAATTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordAAAATTLNumber1:
		return true
	}
	return false
}

type DNSRecordCAA struct {
	// Components of a CAA record.
	Data DNSRecordCAAData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordCAAType `json:"type,required"`
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
	Meta DNSRecordCAAMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordCAATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string           `json:"zone_name" format:"hostname"`
	JSON     dnsRecordCAAJSON `json:"-"`
}

// dnsRecordCAAJSON contains the JSON metadata for the struct [DNSRecordCAA]
type dnsRecordCAAJSON struct {
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

func (r *DNSRecordCAA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordCAAJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordCAA) implementsDNSDNSRecord() {}

// Components of a CAA record.
type DNSRecordCAAData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string               `json:"value"`
	JSON  dnsRecordCAADataJSON `json:"-"`
}

// dnsRecordCAADataJSON contains the JSON metadata for the struct
// [DNSRecordCAAData]
type dnsRecordCAADataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordCAAData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordCAADataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordCAAType string

const (
	DNSRecordCAATypeCAA DNSRecordCAAType = "CAA"
)

func (r DNSRecordCAAType) IsKnown() bool {
	switch r {
	case DNSRecordCAATypeCAA:
		return true
	}
	return false
}

// Extra Cloudflare-specific information about the record.
type DNSRecordCAAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string               `json:"source"`
	JSON   dnsRecordCAAMetaJSON `json:"-"`
}

// dnsRecordCAAMetaJSON contains the JSON metadata for the struct
// [DNSRecordCAAMeta]
type dnsRecordCAAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordCAAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordCAAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.DNSRecordCAATTLNumber].
type DNSRecordCAATTL interface {
	ImplementsDNSDNSRecordCaattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordCAATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordCAATTLNumber(0)),
		},
	)
}

type DNSRecordCAATTLNumber float64

const (
	DNSRecordCAATTLNumber1 DNSRecordCAATTLNumber = 1
)

func (r DNSRecordCAATTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordCAATTLNumber1:
		return true
	}
	return false
}

type DNSRecordCERT struct {
	// Components of a CERT record.
	Data DNSRecordCERTData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordCERTType `json:"type,required"`
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
	Meta DNSRecordCERTMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordCERTTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string            `json:"zone_name" format:"hostname"`
	JSON     dnsRecordCERTJSON `json:"-"`
}

// dnsRecordCERTJSON contains the JSON metadata for the struct [DNSRecordCERT]
type dnsRecordCERTJSON struct {
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

func (r *DNSRecordCERT) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordCERTJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordCERT) implementsDNSDNSRecord() {}

// Components of a CERT record.
type DNSRecordCERTData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64               `json:"type"`
	JSON dnsRecordCERTDataJSON `json:"-"`
}

// dnsRecordCERTDataJSON contains the JSON metadata for the struct
// [DNSRecordCERTData]
type dnsRecordCERTDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordCERTData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordCERTDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordCERTType string

const (
	DNSRecordCERTTypeCERT DNSRecordCERTType = "CERT"
)

func (r DNSRecordCERTType) IsKnown() bool {
	switch r {
	case DNSRecordCERTTypeCERT:
		return true
	}
	return false
}

// Extra Cloudflare-specific information about the record.
type DNSRecordCERTMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                `json:"source"`
	JSON   dnsRecordCERTMetaJSON `json:"-"`
}

// dnsRecordCERTMetaJSON contains the JSON metadata for the struct
// [DNSRecordCERTMeta]
type dnsRecordCERTMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordCERTMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordCERTMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.DNSRecordCERTTTLNumber].
type DNSRecordCERTTTL interface {
	ImplementsDNSDNSRecordCertttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordCERTTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordCERTTTLNumber(0)),
		},
	)
}

type DNSRecordCERTTTLNumber float64

const (
	DNSRecordCERTTTLNumber1 DNSRecordCERTTTLNumber = 1
)

func (r DNSRecordCERTTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordCERTTTLNumber1:
		return true
	}
	return false
}

type DNSRecordCNAME struct {
	// A valid hostname. Must not match the record's name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordCNAMEType `json:"type,required"`
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
	Meta DNSRecordCNAMEMeta `json:"meta"`
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
	TTL DNSRecordCNAMETTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string             `json:"zone_name" format:"hostname"`
	JSON     dnsRecordCNAMEJSON `json:"-"`
}

// dnsRecordCNAMEJSON contains the JSON metadata for the struct [DNSRecordCNAME]
type dnsRecordCNAMEJSON struct {
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

func (r *DNSRecordCNAME) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordCNAMEJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordCNAME) implementsDNSDNSRecord() {}

// Record type.
type DNSRecordCNAMEType string

const (
	DNSRecordCNAMETypeCNAME DNSRecordCNAMEType = "CNAME"
)

func (r DNSRecordCNAMEType) IsKnown() bool {
	switch r {
	case DNSRecordCNAMETypeCNAME:
		return true
	}
	return false
}

// Extra Cloudflare-specific information about the record.
type DNSRecordCNAMEMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                 `json:"source"`
	JSON   dnsRecordCNAMEMetaJSON `json:"-"`
}

// dnsRecordCNAMEMetaJSON contains the JSON metadata for the struct
// [DNSRecordCNAMEMeta]
type dnsRecordCNAMEMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordCNAMEMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordCNAMEMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.DNSRecordCNAMETTLNumber].
type DNSRecordCNAMETTL interface {
	ImplementsDNSDNSRecordCnamettl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordCNAMETTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordCNAMETTLNumber(0)),
		},
	)
}

type DNSRecordCNAMETTLNumber float64

const (
	DNSRecordCNAMETTLNumber1 DNSRecordCNAMETTLNumber = 1
)

func (r DNSRecordCNAMETTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordCNAMETTLNumber1:
		return true
	}
	return false
}

type DNSRecordDNSKEY struct {
	// Components of a DNSKEY record.
	Data DNSRecordDNSKEYData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSKEYType `json:"type,required"`
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
	Meta DNSRecordDNSKEYMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSKEYTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string              `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSKEYJSON `json:"-"`
}

// dnsRecordDNSKEYJSON contains the JSON metadata for the struct [DNSRecordDNSKEY]
type dnsRecordDNSKEYJSON struct {
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

func (r *DNSRecordDNSKEY) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSKEYJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordDNSKEY) implementsDNSDNSRecord() {}

// Components of a DNSKEY record.
type DNSRecordDNSKEYData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                  `json:"public_key"`
	JSON      dnsRecordDNSKEYDataJSON `json:"-"`
}

// dnsRecordDNSKEYDataJSON contains the JSON metadata for the struct
// [DNSRecordDNSKEYData]
type dnsRecordDNSKEYDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSKEYData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSKEYDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordDNSKEYType string

const (
	DNSRecordDNSKEYTypeDNSKEY DNSRecordDNSKEYType = "DNSKEY"
)

func (r DNSRecordDNSKEYType) IsKnown() bool {
	switch r {
	case DNSRecordDNSKEYTypeDNSKEY:
		return true
	}
	return false
}

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSKEYMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                  `json:"source"`
	JSON   dnsRecordDNSKEYMetaJSON `json:"-"`
}

// dnsRecordDNSKEYMetaJSON contains the JSON metadata for the struct
// [DNSRecordDNSKEYMeta]
type dnsRecordDNSKEYMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSKEYMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDNSKEYMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.DNSRecordDNSKEYTTLNumber].
type DNSRecordDNSKEYTTL interface {
	ImplementsDNSDNSRecordDnskeyttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSKEYTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordDNSKEYTTLNumber(0)),
		},
	)
}

type DNSRecordDNSKEYTTLNumber float64

const (
	DNSRecordDNSKEYTTLNumber1 DNSRecordDNSKEYTTLNumber = 1
)

func (r DNSRecordDNSKEYTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordDNSKEYTTLNumber1:
		return true
	}
	return false
}

type DNSRecordDS struct {
	// Components of a DS record.
	Data DNSRecordDSData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDSType `json:"type,required"`
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
	Meta DNSRecordDSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDSJSON `json:"-"`
}

// dnsRecordDSJSON contains the JSON metadata for the struct [DNSRecordDS]
type dnsRecordDSJSON struct {
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

func (r *DNSRecordDS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDSJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordDS) implementsDNSDNSRecord() {}

// Components of a DS record.
type DNSRecordDSData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64             `json:"key_tag"`
	JSON   dnsRecordDSDataJSON `json:"-"`
}

// dnsRecordDSDataJSON contains the JSON metadata for the struct [DNSRecordDSData]
type dnsRecordDSDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDSData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDSDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordDSType string

const (
	DNSRecordDSTypeDS DNSRecordDSType = "DS"
)

func (r DNSRecordDSType) IsKnown() bool {
	switch r {
	case DNSRecordDSTypeDS:
		return true
	}
	return false
}

// Extra Cloudflare-specific information about the record.
type DNSRecordDSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string              `json:"source"`
	JSON   dnsRecordDSMetaJSON `json:"-"`
}

// dnsRecordDSMetaJSON contains the JSON metadata for the struct [DNSRecordDSMeta]
type dnsRecordDSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordDSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.DNSRecordDSTTLNumber].
type DNSRecordDSTTL interface {
	ImplementsDNSDNSRecordDsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordDSTTLNumber(0)),
		},
	)
}

type DNSRecordDSTTLNumber float64

const (
	DNSRecordDSTTLNumber1 DNSRecordDSTTLNumber = 1
)

func (r DNSRecordDSTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordDSTTLNumber1:
		return true
	}
	return false
}

type DNSRecordHTTPS struct {
	// Components of a HTTPS record.
	Data DNSRecordHTTPSData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordHTTPSType `json:"type,required"`
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
	Meta DNSRecordHTTPSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordHTTPSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string             `json:"zone_name" format:"hostname"`
	JSON     dnsRecordHTTPSJSON `json:"-"`
}

// dnsRecordHTTPSJSON contains the JSON metadata for the struct [DNSRecordHTTPS]
type dnsRecordHTTPSJSON struct {
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

func (r *DNSRecordHTTPS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordHTTPSJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordHTTPS) implementsDNSDNSRecord() {}

// Components of a HTTPS record.
type DNSRecordHTTPSData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                 `json:"value"`
	JSON  dnsRecordHTTPSDataJSON `json:"-"`
}

// dnsRecordHTTPSDataJSON contains the JSON metadata for the struct
// [DNSRecordHTTPSData]
type dnsRecordHTTPSDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordHTTPSData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordHTTPSDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordHTTPSType string

const (
	DNSRecordHTTPSTypeHTTPS DNSRecordHTTPSType = "HTTPS"
)

func (r DNSRecordHTTPSType) IsKnown() bool {
	switch r {
	case DNSRecordHTTPSTypeHTTPS:
		return true
	}
	return false
}

// Extra Cloudflare-specific information about the record.
type DNSRecordHTTPSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                 `json:"source"`
	JSON   dnsRecordHTTPSMetaJSON `json:"-"`
}

// dnsRecordHTTPSMetaJSON contains the JSON metadata for the struct
// [DNSRecordHTTPSMeta]
type dnsRecordHTTPSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordHTTPSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordHTTPSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.DNSRecordHTTPSTTLNumber].
type DNSRecordHTTPSTTL interface {
	ImplementsDNSDNSRecordHttpsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordHTTPSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordHTTPSTTLNumber(0)),
		},
	)
}

type DNSRecordHTTPSTTLNumber float64

const (
	DNSRecordHTTPSTTLNumber1 DNSRecordHTTPSTTLNumber = 1
)

func (r DNSRecordHTTPSTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordHTTPSTTLNumber1:
		return true
	}
	return false
}

type DNSRecordLOC struct {
	// Components of a LOC record.
	Data DNSRecordLOCData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordLOCType `json:"type,required"`
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
	Meta DNSRecordLOCMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordLOCTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string           `json:"zone_name" format:"hostname"`
	JSON     dnsRecordLOCJSON `json:"-"`
}

// dnsRecordLOCJSON contains the JSON metadata for the struct [DNSRecordLOC]
type dnsRecordLOCJSON struct {
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

func (r *DNSRecordLOC) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordLOCJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordLOC) implementsDNSDNSRecord() {}

// Components of a LOC record.
type DNSRecordLOCData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection DNSRecordLOCDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection DNSRecordLOCDataLongDirection `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes float64 `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds float64 `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz float64 `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert float64 `json:"precision_vert"`
	// Size of location in meters.
	Size float64              `json:"size"`
	JSON dnsRecordLOCDataJSON `json:"-"`
}

// dnsRecordLOCDataJSON contains the JSON metadata for the struct
// [DNSRecordLOCData]
type dnsRecordLOCDataJSON struct {
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

func (r *DNSRecordLOCData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordLOCDataJSON) RawJSON() string {
	return r.raw
}

// Latitude direction.
type DNSRecordLOCDataLatDirection string

const (
	DNSRecordLOCDataLatDirectionN DNSRecordLOCDataLatDirection = "N"
	DNSRecordLOCDataLatDirectionS DNSRecordLOCDataLatDirection = "S"
)

func (r DNSRecordLOCDataLatDirection) IsKnown() bool {
	switch r {
	case DNSRecordLOCDataLatDirectionN, DNSRecordLOCDataLatDirectionS:
		return true
	}
	return false
}

// Longitude direction.
type DNSRecordLOCDataLongDirection string

const (
	DNSRecordLOCDataLongDirectionE DNSRecordLOCDataLongDirection = "E"
	DNSRecordLOCDataLongDirectionW DNSRecordLOCDataLongDirection = "W"
)

func (r DNSRecordLOCDataLongDirection) IsKnown() bool {
	switch r {
	case DNSRecordLOCDataLongDirectionE, DNSRecordLOCDataLongDirectionW:
		return true
	}
	return false
}

// Record type.
type DNSRecordLOCType string

const (
	DNSRecordLOCTypeLOC DNSRecordLOCType = "LOC"
)

func (r DNSRecordLOCType) IsKnown() bool {
	switch r {
	case DNSRecordLOCTypeLOC:
		return true
	}
	return false
}

// Extra Cloudflare-specific information about the record.
type DNSRecordLOCMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string               `json:"source"`
	JSON   dnsRecordLOCMetaJSON `json:"-"`
}

// dnsRecordLOCMetaJSON contains the JSON metadata for the struct
// [DNSRecordLOCMeta]
type dnsRecordLOCMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordLOCMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordLOCMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.DNSRecordLOCTTLNumber].
type DNSRecordLOCTTL interface {
	ImplementsDNSDNSRecordLocttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordLOCTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordLOCTTLNumber(0)),
		},
	)
}

type DNSRecordLOCTTLNumber float64

const (
	DNSRecordLOCTTLNumber1 DNSRecordLOCTTLNumber = 1
)

func (r DNSRecordLOCTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordLOCTTLNumber1:
		return true
	}
	return false
}

type DNSRecordMX struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordMXType `json:"type,required"`
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
	Meta DNSRecordMXMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordMXTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordMXJSON `json:"-"`
}

// dnsRecordMXJSON contains the JSON metadata for the struct [DNSRecordMX]
type dnsRecordMXJSON struct {
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

func (r *DNSRecordMX) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordMXJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordMX) implementsDNSDNSRecord() {}

// Record type.
type DNSRecordMXType string

const (
	DNSRecordMXTypeMX DNSRecordMXType = "MX"
)

func (r DNSRecordMXType) IsKnown() bool {
	switch r {
	case DNSRecordMXTypeMX:
		return true
	}
	return false
}

// Extra Cloudflare-specific information about the record.
type DNSRecordMXMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string              `json:"source"`
	JSON   dnsRecordMXMetaJSON `json:"-"`
}

// dnsRecordMXMetaJSON contains the JSON metadata for the struct [DNSRecordMXMeta]
type dnsRecordMXMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordMXMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordMXMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.DNSRecordMXTTLNumber].
type DNSRecordMXTTL interface {
	ImplementsDNSDNSRecordMxttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordMXTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordMXTTLNumber(0)),
		},
	)
}

type DNSRecordMXTTLNumber float64

const (
	DNSRecordMXTTLNumber1 DNSRecordMXTTLNumber = 1
)

func (r DNSRecordMXTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordMXTTLNumber1:
		return true
	}
	return false
}

type DNSRecordNAPTR struct {
	// Components of a NAPTR record.
	Data DNSRecordNAPTRData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNAPTRType `json:"type,required"`
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
	Meta DNSRecordNAPTRMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNAPTRTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string             `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNAPTRJSON `json:"-"`
}

// dnsRecordNAPTRJSON contains the JSON metadata for the struct [DNSRecordNAPTR]
type dnsRecordNAPTRJSON struct {
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

func (r *DNSRecordNAPTR) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNAPTRJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordNAPTR) implementsDNSDNSRecord() {}

// Components of a NAPTR record.
type DNSRecordNAPTRData struct {
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
	Service string                 `json:"service"`
	JSON    dnsRecordNAPTRDataJSON `json:"-"`
}

// dnsRecordNAPTRDataJSON contains the JSON metadata for the struct
// [DNSRecordNAPTRData]
type dnsRecordNAPTRDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNAPTRData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNAPTRDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordNAPTRType string

const (
	DNSRecordNAPTRTypeNAPTR DNSRecordNAPTRType = "NAPTR"
)

func (r DNSRecordNAPTRType) IsKnown() bool {
	switch r {
	case DNSRecordNAPTRTypeNAPTR:
		return true
	}
	return false
}

// Extra Cloudflare-specific information about the record.
type DNSRecordNAPTRMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                 `json:"source"`
	JSON   dnsRecordNAPTRMetaJSON `json:"-"`
}

// dnsRecordNAPTRMetaJSON contains the JSON metadata for the struct
// [DNSRecordNAPTRMeta]
type dnsRecordNAPTRMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNAPTRMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNAPTRMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.DNSRecordNAPTRTTLNumber].
type DNSRecordNAPTRTTL interface {
	ImplementsDNSDNSRecordNaptrttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNAPTRTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordNAPTRTTLNumber(0)),
		},
	)
}

type DNSRecordNAPTRTTLNumber float64

const (
	DNSRecordNAPTRTTLNumber1 DNSRecordNAPTRTTLNumber = 1
)

func (r DNSRecordNAPTRTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordNAPTRTTLNumber1:
		return true
	}
	return false
}

type DNSRecordNS struct {
	// A valid name server host name.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNSType `json:"type,required"`
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
	Meta DNSRecordNSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNSJSON `json:"-"`
}

// dnsRecordNSJSON contains the JSON metadata for the struct [DNSRecordNS]
type dnsRecordNSJSON struct {
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

func (r *DNSRecordNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNSJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordNS) implementsDNSDNSRecord() {}

// Record type.
type DNSRecordNSType string

const (
	DNSRecordNSTypeNS DNSRecordNSType = "NS"
)

func (r DNSRecordNSType) IsKnown() bool {
	switch r {
	case DNSRecordNSTypeNS:
		return true
	}
	return false
}

// Extra Cloudflare-specific information about the record.
type DNSRecordNSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string              `json:"source"`
	JSON   dnsRecordNSMetaJSON `json:"-"`
}

// dnsRecordNSMetaJSON contains the JSON metadata for the struct [DNSRecordNSMeta]
type dnsRecordNSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.DNSRecordNSTTLNumber].
type DNSRecordNSTTL interface {
	ImplementsDNSDNSRecordNsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordNSTTLNumber(0)),
		},
	)
}

type DNSRecordNSTTLNumber float64

const (
	DNSRecordNSTTLNumber1 DNSRecordNSTTLNumber = 1
)

func (r DNSRecordNSTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordNSTTLNumber1:
		return true
	}
	return false
}

type DNSRecordPTR struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordPTRType `json:"type,required"`
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
	Meta DNSRecordPTRMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordPTRTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string           `json:"zone_name" format:"hostname"`
	JSON     dnsRecordPTRJSON `json:"-"`
}

// dnsRecordPTRJSON contains the JSON metadata for the struct [DNSRecordPTR]
type dnsRecordPTRJSON struct {
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

func (r *DNSRecordPTR) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordPTRJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordPTR) implementsDNSDNSRecord() {}

// Record type.
type DNSRecordPTRType string

const (
	DNSRecordPTRTypePTR DNSRecordPTRType = "PTR"
)

func (r DNSRecordPTRType) IsKnown() bool {
	switch r {
	case DNSRecordPTRTypePTR:
		return true
	}
	return false
}

// Extra Cloudflare-specific information about the record.
type DNSRecordPTRMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string               `json:"source"`
	JSON   dnsRecordPTRMetaJSON `json:"-"`
}

// dnsRecordPTRMetaJSON contains the JSON metadata for the struct
// [DNSRecordPTRMeta]
type dnsRecordPTRMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordPTRMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordPTRMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.DNSRecordPTRTTLNumber].
type DNSRecordPTRTTL interface {
	ImplementsDNSDNSRecordPtrttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordPTRTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordPTRTTLNumber(0)),
		},
	)
}

type DNSRecordPTRTTLNumber float64

const (
	DNSRecordPTRTTLNumber1 DNSRecordPTRTTLNumber = 1
)

func (r DNSRecordPTRTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordPTRTTLNumber1:
		return true
	}
	return false
}

type DNSRecordSMIMEA struct {
	// Components of a SMIMEA record.
	Data DNSRecordSMIMEAData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordSMIMEAType `json:"type,required"`
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
	Meta DNSRecordSMIMEAMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordSMIMEATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string              `json:"zone_name" format:"hostname"`
	JSON     dnsRecordSMIMEAJSON `json:"-"`
}

// dnsRecordSMIMEAJSON contains the JSON metadata for the struct [DNSRecordSMIMEA]
type dnsRecordSMIMEAJSON struct {
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

func (r *DNSRecordSMIMEA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordSMIMEAJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordSMIMEA) implementsDNSDNSRecord() {}

// Components of a SMIMEA record.
type DNSRecordSMIMEAData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                 `json:"usage"`
	JSON  dnsRecordSMIMEADataJSON `json:"-"`
}

// dnsRecordSMIMEADataJSON contains the JSON metadata for the struct
// [DNSRecordSMIMEAData]
type dnsRecordSMIMEADataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordSMIMEAData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordSMIMEADataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordSMIMEAType string

const (
	DNSRecordSMIMEATypeSMIMEA DNSRecordSMIMEAType = "SMIMEA"
)

func (r DNSRecordSMIMEAType) IsKnown() bool {
	switch r {
	case DNSRecordSMIMEATypeSMIMEA:
		return true
	}
	return false
}

// Extra Cloudflare-specific information about the record.
type DNSRecordSMIMEAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                  `json:"source"`
	JSON   dnsRecordSMIMEAMetaJSON `json:"-"`
}

// dnsRecordSMIMEAMetaJSON contains the JSON metadata for the struct
// [DNSRecordSMIMEAMeta]
type dnsRecordSMIMEAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordSMIMEAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordSMIMEAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.DNSRecordSMIMEATTLNumber].
type DNSRecordSMIMEATTL interface {
	ImplementsDNSDNSRecordSmimeattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordSMIMEATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordSMIMEATTLNumber(0)),
		},
	)
}

type DNSRecordSMIMEATTLNumber float64

const (
	DNSRecordSMIMEATTLNumber1 DNSRecordSMIMEATTLNumber = 1
)

func (r DNSRecordSMIMEATTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordSMIMEATTLNumber1:
		return true
	}
	return false
}

type DNSRecordSRV struct {
	// Components of a SRV record.
	Data DNSRecordSRVData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordSRVType `json:"type,required"`
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
	Meta DNSRecordSRVMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordSRVTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string           `json:"zone_name" format:"hostname"`
	JSON     dnsRecordSRVJSON `json:"-"`
}

// dnsRecordSRVJSON contains the JSON metadata for the struct [DNSRecordSRV]
type dnsRecordSRVJSON struct {
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

func (r *DNSRecordSRV) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordSRVJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordSRV) implementsDNSDNSRecord() {}

// Components of a SRV record.
type DNSRecordSRVData struct {
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
	Weight float64              `json:"weight"`
	JSON   dnsRecordSRVDataJSON `json:"-"`
}

// dnsRecordSRVDataJSON contains the JSON metadata for the struct
// [DNSRecordSRVData]
type dnsRecordSRVDataJSON struct {
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

func (r *DNSRecordSRVData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordSRVDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordSRVType string

const (
	DNSRecordSRVTypeSRV DNSRecordSRVType = "SRV"
)

func (r DNSRecordSRVType) IsKnown() bool {
	switch r {
	case DNSRecordSRVTypeSRV:
		return true
	}
	return false
}

// Extra Cloudflare-specific information about the record.
type DNSRecordSRVMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string               `json:"source"`
	JSON   dnsRecordSRVMetaJSON `json:"-"`
}

// dnsRecordSRVMetaJSON contains the JSON metadata for the struct
// [DNSRecordSRVMeta]
type dnsRecordSRVMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordSRVMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordSRVMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.DNSRecordSRVTTLNumber].
type DNSRecordSRVTTL interface {
	ImplementsDNSDNSRecordSrvttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordSRVTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordSRVTTLNumber(0)),
		},
	)
}

type DNSRecordSRVTTLNumber float64

const (
	DNSRecordSRVTTLNumber1 DNSRecordSRVTTLNumber = 1
)

func (r DNSRecordSRVTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordSRVTTLNumber1:
		return true
	}
	return false
}

type DNSRecordSSHFP struct {
	// Components of a SSHFP record.
	Data DNSRecordSSHFPData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordSSHFPType `json:"type,required"`
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
	Meta DNSRecordSSHFPMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordSSHFPTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string             `json:"zone_name" format:"hostname"`
	JSON     dnsRecordSSHFPJSON `json:"-"`
}

// dnsRecordSSHFPJSON contains the JSON metadata for the struct [DNSRecordSSHFP]
type dnsRecordSSHFPJSON struct {
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

func (r *DNSRecordSSHFP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordSSHFPJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordSSHFP) implementsDNSDNSRecord() {}

// Components of a SSHFP record.
type DNSRecordSSHFPData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                `json:"type"`
	JSON dnsRecordSSHFPDataJSON `json:"-"`
}

// dnsRecordSSHFPDataJSON contains the JSON metadata for the struct
// [DNSRecordSSHFPData]
type dnsRecordSSHFPDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordSSHFPData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordSSHFPDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordSSHFPType string

const (
	DNSRecordSSHFPTypeSSHFP DNSRecordSSHFPType = "SSHFP"
)

func (r DNSRecordSSHFPType) IsKnown() bool {
	switch r {
	case DNSRecordSSHFPTypeSSHFP:
		return true
	}
	return false
}

// Extra Cloudflare-specific information about the record.
type DNSRecordSSHFPMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                 `json:"source"`
	JSON   dnsRecordSSHFPMetaJSON `json:"-"`
}

// dnsRecordSSHFPMetaJSON contains the JSON metadata for the struct
// [DNSRecordSSHFPMeta]
type dnsRecordSSHFPMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordSSHFPMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordSSHFPMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.DNSRecordSSHFPTTLNumber].
type DNSRecordSSHFPTTL interface {
	ImplementsDNSDNSRecordSshfpttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordSSHFPTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordSSHFPTTLNumber(0)),
		},
	)
}

type DNSRecordSSHFPTTLNumber float64

const (
	DNSRecordSSHFPTTLNumber1 DNSRecordSSHFPTTLNumber = 1
)

func (r DNSRecordSSHFPTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordSSHFPTTLNumber1:
		return true
	}
	return false
}

type DNSRecordSVCB struct {
	// Components of a SVCB record.
	Data DNSRecordSVCBData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordSVCBType `json:"type,required"`
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
	Meta DNSRecordSVCBMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordSVCBTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string            `json:"zone_name" format:"hostname"`
	JSON     dnsRecordSVCBJSON `json:"-"`
}

// dnsRecordSVCBJSON contains the JSON metadata for the struct [DNSRecordSVCB]
type dnsRecordSVCBJSON struct {
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

func (r *DNSRecordSVCB) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordSVCBJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordSVCB) implementsDNSDNSRecord() {}

// Components of a SVCB record.
type DNSRecordSVCBData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                `json:"value"`
	JSON  dnsRecordSVCBDataJSON `json:"-"`
}

// dnsRecordSVCBDataJSON contains the JSON metadata for the struct
// [DNSRecordSVCBData]
type dnsRecordSVCBDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordSVCBData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordSVCBDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordSVCBType string

const (
	DNSRecordSVCBTypeSVCB DNSRecordSVCBType = "SVCB"
)

func (r DNSRecordSVCBType) IsKnown() bool {
	switch r {
	case DNSRecordSVCBTypeSVCB:
		return true
	}
	return false
}

// Extra Cloudflare-specific information about the record.
type DNSRecordSVCBMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                `json:"source"`
	JSON   dnsRecordSVCBMetaJSON `json:"-"`
}

// dnsRecordSVCBMetaJSON contains the JSON metadata for the struct
// [DNSRecordSVCBMeta]
type dnsRecordSVCBMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordSVCBMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordSVCBMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.DNSRecordSVCBTTLNumber].
type DNSRecordSVCBTTL interface {
	ImplementsDNSDNSRecordSvcbttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordSVCBTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordSVCBTTLNumber(0)),
		},
	)
}

type DNSRecordSVCBTTLNumber float64

const (
	DNSRecordSVCBTTLNumber1 DNSRecordSVCBTTLNumber = 1
)

func (r DNSRecordSVCBTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordSVCBTTLNumber1:
		return true
	}
	return false
}

type DNSRecordTLSA struct {
	// Components of a TLSA record.
	Data DNSRecordTLSAData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordTLSAType `json:"type,required"`
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
	Meta DNSRecordTLSAMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordTLSATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string            `json:"zone_name" format:"hostname"`
	JSON     dnsRecordTLSAJSON `json:"-"`
}

// dnsRecordTLSAJSON contains the JSON metadata for the struct [DNSRecordTLSA]
type dnsRecordTLSAJSON struct {
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

func (r *DNSRecordTLSA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordTLSAJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordTLSA) implementsDNSDNSRecord() {}

// Components of a TLSA record.
type DNSRecordTLSAData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64               `json:"usage"`
	JSON  dnsRecordTLSADataJSON `json:"-"`
}

// dnsRecordTLSADataJSON contains the JSON metadata for the struct
// [DNSRecordTLSAData]
type dnsRecordTLSADataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordTLSAData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordTLSADataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordTLSAType string

const (
	DNSRecordTLSATypeTLSA DNSRecordTLSAType = "TLSA"
)

func (r DNSRecordTLSAType) IsKnown() bool {
	switch r {
	case DNSRecordTLSATypeTLSA:
		return true
	}
	return false
}

// Extra Cloudflare-specific information about the record.
type DNSRecordTLSAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                `json:"source"`
	JSON   dnsRecordTLSAMetaJSON `json:"-"`
}

// dnsRecordTLSAMetaJSON contains the JSON metadata for the struct
// [DNSRecordTLSAMeta]
type dnsRecordTLSAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordTLSAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordTLSAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.DNSRecordTLSATTLNumber].
type DNSRecordTLSATTL interface {
	ImplementsDNSDNSRecordTlsattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordTLSATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordTLSATTLNumber(0)),
		},
	)
}

type DNSRecordTLSATTLNumber float64

const (
	DNSRecordTLSATTLNumber1 DNSRecordTLSATTLNumber = 1
)

func (r DNSRecordTLSATTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordTLSATTLNumber1:
		return true
	}
	return false
}

type DNSRecordTXT struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordTXTType `json:"type,required"`
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
	Meta DNSRecordTXTMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordTXTTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string           `json:"zone_name" format:"hostname"`
	JSON     dnsRecordTXTJSON `json:"-"`
}

// dnsRecordTXTJSON contains the JSON metadata for the struct [DNSRecordTXT]
type dnsRecordTXTJSON struct {
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

func (r *DNSRecordTXT) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordTXTJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordTXT) implementsDNSDNSRecord() {}

// Record type.
type DNSRecordTXTType string

const (
	DNSRecordTXTTypeTXT DNSRecordTXTType = "TXT"
)

func (r DNSRecordTXTType) IsKnown() bool {
	switch r {
	case DNSRecordTXTTypeTXT:
		return true
	}
	return false
}

// Extra Cloudflare-specific information about the record.
type DNSRecordTXTMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string               `json:"source"`
	JSON   dnsRecordTXTMetaJSON `json:"-"`
}

// dnsRecordTXTMetaJSON contains the JSON metadata for the struct
// [DNSRecordTXTMeta]
type dnsRecordTXTMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordTXTMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordTXTMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.DNSRecordTXTTTLNumber].
type DNSRecordTXTTTL interface {
	ImplementsDNSDNSRecordTxtttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordTXTTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordTXTTTLNumber(0)),
		},
	)
}

type DNSRecordTXTTTLNumber float64

const (
	DNSRecordTXTTTLNumber1 DNSRecordTXTTTLNumber = 1
)

func (r DNSRecordTXTTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordTXTTTLNumber1:
		return true
	}
	return false
}

type DNSRecordURI struct {
	// Components of a URI record.
	Data DNSRecordURIData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordURIType `json:"type,required"`
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
	Meta DNSRecordURIMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordURITTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string           `json:"zone_name" format:"hostname"`
	JSON     dnsRecordURIJSON `json:"-"`
}

// dnsRecordURIJSON contains the JSON metadata for the struct [DNSRecordURI]
type dnsRecordURIJSON struct {
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

func (r *DNSRecordURI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordURIJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordURI) implementsDNSDNSRecord() {}

// Components of a URI record.
type DNSRecordURIData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64              `json:"weight"`
	JSON   dnsRecordURIDataJSON `json:"-"`
}

// dnsRecordURIDataJSON contains the JSON metadata for the struct
// [DNSRecordURIData]
type dnsRecordURIDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordURIData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordURIDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordURIType string

const (
	DNSRecordURITypeURI DNSRecordURIType = "URI"
)

func (r DNSRecordURIType) IsKnown() bool {
	switch r {
	case DNSRecordURITypeURI:
		return true
	}
	return false
}

// Extra Cloudflare-specific information about the record.
type DNSRecordURIMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string               `json:"source"`
	JSON   dnsRecordURIMetaJSON `json:"-"`
}

// dnsRecordURIMetaJSON contains the JSON metadata for the struct
// [DNSRecordURIMeta]
type dnsRecordURIMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordURIMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordURIMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.DNSRecordURITTLNumber].
type DNSRecordURITTL interface {
	ImplementsDNSDNSRecordUrittl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordURITTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordURITTLNumber(0)),
		},
	)
}

type DNSRecordURITTLNumber float64

const (
	DNSRecordURITTLNumber1 DNSRecordURITTLNumber = 1
)

func (r DNSRecordURITTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordURITTLNumber1:
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
	TTL param.Field[RecordNewParamsDNSRecordsARecordTTL] `json:"ttl"`
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
type RecordNewParamsDNSRecordsARecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsARecordTTL()
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
	TTL param.Field[RecordNewParamsDNSRecordsAAAARecordTTL] `json:"ttl"`
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
type RecordNewParamsDNSRecordsAAAARecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsAAAARecordTTL()
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
	TTL param.Field[RecordNewParamsDNSRecordsCAARecordTTL] `json:"ttl"`
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
type RecordNewParamsDNSRecordsCAARecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsCAARecordTTL()
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
	TTL param.Field[RecordNewParamsDNSRecordsCERTRecordTTL] `json:"ttl"`
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
type RecordNewParamsDNSRecordsCERTRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsCERTRecordTTL()
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
	TTL param.Field[RecordNewParamsDNSRecordsCNAMERecordTTL] `json:"ttl"`
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
type RecordNewParamsDNSRecordsCNAMERecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsCNAMERecordTTL()
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
	TTL param.Field[RecordNewParamsDNSRecordsDNSKEYRecordTTL] `json:"ttl"`
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
type RecordNewParamsDNSRecordsDNSKEYRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsDNSKEYRecordTTL()
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
	TTL param.Field[RecordNewParamsDNSRecordsDSRecordTTL] `json:"ttl"`
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
type RecordNewParamsDNSRecordsDSRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsDSRecordTTL()
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
	TTL param.Field[RecordNewParamsDNSRecordsHTTPSRecordTTL] `json:"ttl"`
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
type RecordNewParamsDNSRecordsHTTPSRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsHTTPSRecordTTL()
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
	TTL param.Field[RecordNewParamsDNSRecordsLOCRecordTTL] `json:"ttl"`
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
type RecordNewParamsDNSRecordsLOCRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsLOCRecordTTL()
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
	TTL param.Field[RecordNewParamsDNSRecordsMXRecordTTL] `json:"ttl"`
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
type RecordNewParamsDNSRecordsMXRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsMXRecordTTL()
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
	TTL param.Field[RecordNewParamsDNSRecordsNAPTRRecordTTL] `json:"ttl"`
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
type RecordNewParamsDNSRecordsNAPTRRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsNAPTRRecordTTL()
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
	TTL param.Field[RecordNewParamsDNSRecordsNSRecordTTL] `json:"ttl"`
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
type RecordNewParamsDNSRecordsNSRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsNSRecordTTL()
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
	TTL param.Field[RecordNewParamsDNSRecordsPTRRecordTTL] `json:"ttl"`
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
type RecordNewParamsDNSRecordsPTRRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsPTRRecordTTL()
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
	TTL param.Field[RecordNewParamsDNSRecordsSMIMEARecordTTL] `json:"ttl"`
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
type RecordNewParamsDNSRecordsSMIMEARecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsSMIMEARecordTTL()
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
	TTL param.Field[RecordNewParamsDNSRecordsSRVRecordTTL] `json:"ttl"`
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
type RecordNewParamsDNSRecordsSRVRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsSRVRecordTTL()
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
	TTL param.Field[RecordNewParamsDNSRecordsSSHFPRecordTTL] `json:"ttl"`
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
type RecordNewParamsDNSRecordsSSHFPRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsSSHFPRecordTTL()
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
	TTL param.Field[RecordNewParamsDNSRecordsSVCBRecordTTL] `json:"ttl"`
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
type RecordNewParamsDNSRecordsSVCBRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsSVCBRecordTTL()
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
	TTL param.Field[RecordNewParamsDNSRecordsTLSARecordTTL] `json:"ttl"`
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
type RecordNewParamsDNSRecordsTLSARecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsTLSARecordTTL()
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
	TTL param.Field[RecordNewParamsDNSRecordsTXTRecordTTL] `json:"ttl"`
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
type RecordNewParamsDNSRecordsTXTRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsTXTRecordTTL()
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
	TTL param.Field[RecordNewParamsDNSRecordsURIRecordTTL] `json:"ttl"`
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
type RecordNewParamsDNSRecordsURIRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsURIRecordTTL()
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
	Errors   []RecordNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RecordNewResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSRecord                           `json:"result,required"`
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

type RecordNewResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    recordNewResponseEnvelopeErrorsJSON `json:"-"`
}

// recordNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RecordNewResponseEnvelopeErrors]
type recordNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RecordNewResponseEnvelopeMessages struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    recordNewResponseEnvelopeMessagesJSON `json:"-"`
}

// recordNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RecordNewResponseEnvelopeMessages]
type recordNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseEnvelopeMessagesJSON) RawJSON() string {
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
	TTL param.Field[RecordUpdateParamsDNSRecordsARecordTTL] `json:"ttl"`
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
type RecordUpdateParamsDNSRecordsARecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsARecordTTL()
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
	TTL param.Field[RecordUpdateParamsDNSRecordsAAAARecordTTL] `json:"ttl"`
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
type RecordUpdateParamsDNSRecordsAAAARecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsAAAARecordTTL()
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
	TTL param.Field[RecordUpdateParamsDNSRecordsCAARecordTTL] `json:"ttl"`
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
type RecordUpdateParamsDNSRecordsCAARecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsCAARecordTTL()
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
	TTL param.Field[RecordUpdateParamsDNSRecordsCERTRecordTTL] `json:"ttl"`
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
type RecordUpdateParamsDNSRecordsCERTRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsCERTRecordTTL()
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
	TTL param.Field[RecordUpdateParamsDNSRecordsCNAMERecordTTL] `json:"ttl"`
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
type RecordUpdateParamsDNSRecordsCNAMERecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsCNAMERecordTTL()
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
	TTL param.Field[RecordUpdateParamsDNSRecordsDNSKEYRecordTTL] `json:"ttl"`
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
type RecordUpdateParamsDNSRecordsDNSKEYRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsDNSKEYRecordTTL()
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
	TTL param.Field[RecordUpdateParamsDNSRecordsDSRecordTTL] `json:"ttl"`
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
type RecordUpdateParamsDNSRecordsDSRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsDSRecordTTL()
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
	TTL param.Field[RecordUpdateParamsDNSRecordsHTTPSRecordTTL] `json:"ttl"`
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
type RecordUpdateParamsDNSRecordsHTTPSRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsHTTPSRecordTTL()
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
	TTL param.Field[RecordUpdateParamsDNSRecordsLOCRecordTTL] `json:"ttl"`
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
type RecordUpdateParamsDNSRecordsLOCRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsLOCRecordTTL()
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
	TTL param.Field[RecordUpdateParamsDNSRecordsMXRecordTTL] `json:"ttl"`
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
type RecordUpdateParamsDNSRecordsMXRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsMXRecordTTL()
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
	TTL param.Field[RecordUpdateParamsDNSRecordsNAPTRRecordTTL] `json:"ttl"`
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
type RecordUpdateParamsDNSRecordsNAPTRRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsNAPTRRecordTTL()
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
	TTL param.Field[RecordUpdateParamsDNSRecordsNSRecordTTL] `json:"ttl"`
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
type RecordUpdateParamsDNSRecordsNSRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsNSRecordTTL()
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
	TTL param.Field[RecordUpdateParamsDNSRecordsPTRRecordTTL] `json:"ttl"`
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
type RecordUpdateParamsDNSRecordsPTRRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsPTRRecordTTL()
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
	TTL param.Field[RecordUpdateParamsDNSRecordsSMIMEARecordTTL] `json:"ttl"`
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
type RecordUpdateParamsDNSRecordsSMIMEARecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsSMIMEARecordTTL()
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
	TTL param.Field[RecordUpdateParamsDNSRecordsSRVRecordTTL] `json:"ttl"`
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
type RecordUpdateParamsDNSRecordsSRVRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsSRVRecordTTL()
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
	TTL param.Field[RecordUpdateParamsDNSRecordsSSHFPRecordTTL] `json:"ttl"`
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
type RecordUpdateParamsDNSRecordsSSHFPRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsSSHFPRecordTTL()
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
	TTL param.Field[RecordUpdateParamsDNSRecordsSVCBRecordTTL] `json:"ttl"`
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
type RecordUpdateParamsDNSRecordsSVCBRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsSVCBRecordTTL()
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
	TTL param.Field[RecordUpdateParamsDNSRecordsTLSARecordTTL] `json:"ttl"`
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
type RecordUpdateParamsDNSRecordsTLSARecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsTLSARecordTTL()
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
	TTL param.Field[RecordUpdateParamsDNSRecordsTXTRecordTTL] `json:"ttl"`
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
type RecordUpdateParamsDNSRecordsTXTRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsTXTRecordTTL()
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
	TTL param.Field[RecordUpdateParamsDNSRecordsURIRecordTTL] `json:"ttl"`
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
type RecordUpdateParamsDNSRecordsURIRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsURIRecordTTL()
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
	Errors   []RecordUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RecordUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSRecord                              `json:"result,required"`
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

type RecordUpdateResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    recordUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// recordUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RecordUpdateResponseEnvelopeErrors]
type recordUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RecordUpdateResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    recordUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// recordUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RecordUpdateResponseEnvelopeMessages]
type recordUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
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
	TTL param.Field[RecordEditParamsDNSRecordsARecordTTL] `json:"ttl"`
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
type RecordEditParamsDNSRecordsARecordTTL interface {
	ImplementsDNSRecordEditParamsDNSRecordsARecordTTL()
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
	TTL param.Field[RecordEditParamsDNSRecordsAAAARecordTTL] `json:"ttl"`
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
type RecordEditParamsDNSRecordsAAAARecordTTL interface {
	ImplementsDNSRecordEditParamsDNSRecordsAAAARecordTTL()
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
	TTL param.Field[RecordEditParamsDNSRecordsCAARecordTTL] `json:"ttl"`
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
type RecordEditParamsDNSRecordsCAARecordTTL interface {
	ImplementsDNSRecordEditParamsDNSRecordsCAARecordTTL()
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
	TTL param.Field[RecordEditParamsDNSRecordsCERTRecordTTL] `json:"ttl"`
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
type RecordEditParamsDNSRecordsCERTRecordTTL interface {
	ImplementsDNSRecordEditParamsDNSRecordsCERTRecordTTL()
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
	TTL param.Field[RecordEditParamsDNSRecordsCNAMERecordTTL] `json:"ttl"`
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
type RecordEditParamsDNSRecordsCNAMERecordTTL interface {
	ImplementsDNSRecordEditParamsDNSRecordsCNAMERecordTTL()
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
	TTL param.Field[RecordEditParamsDNSRecordsDNSKEYRecordTTL] `json:"ttl"`
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
type RecordEditParamsDNSRecordsDNSKEYRecordTTL interface {
	ImplementsDNSRecordEditParamsDNSRecordsDNSKEYRecordTTL()
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
	TTL param.Field[RecordEditParamsDNSRecordsDSRecordTTL] `json:"ttl"`
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
type RecordEditParamsDNSRecordsDSRecordTTL interface {
	ImplementsDNSRecordEditParamsDNSRecordsDSRecordTTL()
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
	TTL param.Field[RecordEditParamsDNSRecordsHTTPSRecordTTL] `json:"ttl"`
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
type RecordEditParamsDNSRecordsHTTPSRecordTTL interface {
	ImplementsDNSRecordEditParamsDNSRecordsHTTPSRecordTTL()
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
	TTL param.Field[RecordEditParamsDNSRecordsLOCRecordTTL] `json:"ttl"`
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
type RecordEditParamsDNSRecordsLOCRecordTTL interface {
	ImplementsDNSRecordEditParamsDNSRecordsLOCRecordTTL()
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
	TTL param.Field[RecordEditParamsDNSRecordsMXRecordTTL] `json:"ttl"`
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
type RecordEditParamsDNSRecordsMXRecordTTL interface {
	ImplementsDNSRecordEditParamsDNSRecordsMXRecordTTL()
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
	TTL param.Field[RecordEditParamsDNSRecordsNAPTRRecordTTL] `json:"ttl"`
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
type RecordEditParamsDNSRecordsNAPTRRecordTTL interface {
	ImplementsDNSRecordEditParamsDNSRecordsNAPTRRecordTTL()
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
	TTL param.Field[RecordEditParamsDNSRecordsNSRecordTTL] `json:"ttl"`
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
type RecordEditParamsDNSRecordsNSRecordTTL interface {
	ImplementsDNSRecordEditParamsDNSRecordsNSRecordTTL()
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
	TTL param.Field[RecordEditParamsDNSRecordsPTRRecordTTL] `json:"ttl"`
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
type RecordEditParamsDNSRecordsPTRRecordTTL interface {
	ImplementsDNSRecordEditParamsDNSRecordsPTRRecordTTL()
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
	TTL param.Field[RecordEditParamsDNSRecordsSMIMEARecordTTL] `json:"ttl"`
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
type RecordEditParamsDNSRecordsSMIMEARecordTTL interface {
	ImplementsDNSRecordEditParamsDNSRecordsSMIMEARecordTTL()
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
	TTL param.Field[RecordEditParamsDNSRecordsSRVRecordTTL] `json:"ttl"`
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
type RecordEditParamsDNSRecordsSRVRecordTTL interface {
	ImplementsDNSRecordEditParamsDNSRecordsSRVRecordTTL()
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
	TTL param.Field[RecordEditParamsDNSRecordsSSHFPRecordTTL] `json:"ttl"`
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
type RecordEditParamsDNSRecordsSSHFPRecordTTL interface {
	ImplementsDNSRecordEditParamsDNSRecordsSSHFPRecordTTL()
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
	TTL param.Field[RecordEditParamsDNSRecordsSVCBRecordTTL] `json:"ttl"`
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
type RecordEditParamsDNSRecordsSVCBRecordTTL interface {
	ImplementsDNSRecordEditParamsDNSRecordsSVCBRecordTTL()
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
	TTL param.Field[RecordEditParamsDNSRecordsTLSARecordTTL] `json:"ttl"`
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
type RecordEditParamsDNSRecordsTLSARecordTTL interface {
	ImplementsDNSRecordEditParamsDNSRecordsTLSARecordTTL()
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
	TTL param.Field[RecordEditParamsDNSRecordsTXTRecordTTL] `json:"ttl"`
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
type RecordEditParamsDNSRecordsTXTRecordTTL interface {
	ImplementsDNSRecordEditParamsDNSRecordsTXTRecordTTL()
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
	TTL param.Field[RecordEditParamsDNSRecordsURIRecordTTL] `json:"ttl"`
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
type RecordEditParamsDNSRecordsURIRecordTTL interface {
	ImplementsDNSRecordEditParamsDNSRecordsURIRecordTTL()
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
	Errors   []RecordEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RecordEditResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSRecord                            `json:"result,required"`
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

type RecordEditResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    recordEditResponseEnvelopeErrorsJSON `json:"-"`
}

// recordEditResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RecordEditResponseEnvelopeErrors]
type recordEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RecordEditResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    recordEditResponseEnvelopeMessagesJSON `json:"-"`
}

// recordEditResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RecordEditResponseEnvelopeMessages]
type recordEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseEnvelopeMessagesJSON) RawJSON() string {
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
	Errors   []RecordGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RecordGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSRecord                           `json:"result,required"`
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

type RecordGetResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    recordGetResponseEnvelopeErrorsJSON `json:"-"`
}

// recordGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RecordGetResponseEnvelopeErrors]
type recordGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RecordGetResponseEnvelopeMessages struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    recordGetResponseEnvelopeMessagesJSON `json:"-"`
}

// recordGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RecordGetResponseEnvelopeMessages]
type recordGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseEnvelopeMessagesJSON) RawJSON() string {
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
	Errors   []RecordImportResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RecordImportResponseEnvelopeMessages `json:"messages,required"`
	Result   RecordImportResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success RecordImportResponseEnvelopeSuccess `json:"success,required"`
	Timing  RecordImportResponseEnvelopeTiming  `json:"timing"`
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

type RecordImportResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    recordImportResponseEnvelopeErrorsJSON `json:"-"`
}

// recordImportResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RecordImportResponseEnvelopeErrors]
type recordImportResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordImportResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordImportResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RecordImportResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    recordImportResponseEnvelopeMessagesJSON `json:"-"`
}

// recordImportResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RecordImportResponseEnvelopeMessages]
type recordImportResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordImportResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordImportResponseEnvelopeMessagesJSON) RawJSON() string {
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

type RecordImportResponseEnvelopeTiming struct {
	// When the file parsing ended.
	EndTime time.Time `json:"end_time" format:"date-time"`
	// Processing time of the file in seconds.
	ProcessTime float64 `json:"process_time"`
	// When the file parsing started.
	StartTime time.Time                              `json:"start_time" format:"date-time"`
	JSON      recordImportResponseEnvelopeTimingJSON `json:"-"`
}

// recordImportResponseEnvelopeTimingJSON contains the JSON metadata for the struct
// [RecordImportResponseEnvelopeTiming]
type recordImportResponseEnvelopeTimingJSON struct {
	EndTime     apijson.Field
	ProcessTime apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordImportResponseEnvelopeTiming) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordImportResponseEnvelopeTimingJSON) RawJSON() string {
	return r.raw
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
	Errors   []RecordScanResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RecordScanResponseEnvelopeMessages `json:"messages,required"`
	Result   RecordScanResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success RecordScanResponseEnvelopeSuccess `json:"success,required"`
	Timing  RecordScanResponseEnvelopeTiming  `json:"timing"`
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

type RecordScanResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    recordScanResponseEnvelopeErrorsJSON `json:"-"`
}

// recordScanResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RecordScanResponseEnvelopeErrors]
type recordScanResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordScanResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordScanResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RecordScanResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    recordScanResponseEnvelopeMessagesJSON `json:"-"`
}

// recordScanResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RecordScanResponseEnvelopeMessages]
type recordScanResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordScanResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordScanResponseEnvelopeMessagesJSON) RawJSON() string {
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

type RecordScanResponseEnvelopeTiming struct {
	// When the file parsing ended.
	EndTime time.Time `json:"end_time" format:"date-time"`
	// Processing time of the file in seconds.
	ProcessTime float64 `json:"process_time"`
	// When the file parsing started.
	StartTime time.Time                            `json:"start_time" format:"date-time"`
	JSON      recordScanResponseEnvelopeTimingJSON `json:"-"`
}

// recordScanResponseEnvelopeTimingJSON contains the JSON metadata for the struct
// [RecordScanResponseEnvelopeTiming]
type recordScanResponseEnvelopeTimingJSON struct {
	EndTime     apijson.Field
	ProcessTime apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordScanResponseEnvelopeTiming) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordScanResponseEnvelopeTimingJSON) RawJSON() string {
	return r.raw
}
