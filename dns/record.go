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
	path := fmt.Sprintf("zones/%s/dns_records", params.ZoneID)
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
	path := fmt.Sprintf("zones/%s/dns_records/%s", params.ZoneID, dnsRecordID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List, search, sort, and filter a zones' DNS records.
func (r *RecordService) List(ctx context.Context, params RecordListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[DNSRecord], err error) {
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
func (r *RecordService) ListAutoPaging(ctx context.Context, params RecordListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[DNSRecord] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete DNS Record
func (r *RecordService) Delete(ctx context.Context, dnsRecordID string, body RecordDeleteParams, opts ...option.RequestOption) (res *RecordDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RecordDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_records/%s", body.ZoneID, dnsRecordID)
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
	path := fmt.Sprintf("zones/%s/dns_records/%s", params.ZoneID, dnsRecordID)
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
func (r *RecordService) Scan(ctx context.Context, body RecordScanParams, opts ...option.RequestOption) (res *RecordScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RecordScanResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_records/scan", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [dns.DNSRecordA], [dns.DNSRecordAAAA], [dns.DNSRecordCAA],
// [dns.DNSRecordCert], [dns.DNSRecordCNAME], [dns.DNSRecordDNSKEY],
// [dns.DNSRecordDS], [dns.DNSRecordHTTPS], [dns.DNSRecordLOC], [dns.DNSRecordMX],
// [dns.DNSRecordNAPTR], [dns.DNSRecordNS], [dns.DNSRecordPTR],
// [dns.DNSRecordSmimea], [dns.DNSRecordSRV], [dns.DNSRecordSSHFP],
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
			Type:               reflect.TypeOf(DNSRecordCert{}),
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
			Type:               reflect.TypeOf(DNSRecordSmimea{}),
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

type DNSRecordCert struct {
	// Components of a CERT record.
	Data DNSRecordCertData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordCertType `json:"type,required"`
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
	Meta DNSRecordCertMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordCertTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string            `json:"zone_name" format:"hostname"`
	JSON     dnsRecordCertJSON `json:"-"`
}

// dnsRecordCertJSON contains the JSON metadata for the struct [DNSRecordCert]
type dnsRecordCertJSON struct {
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

func (r *DNSRecordCert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordCertJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordCert) implementsDNSDNSRecord() {}

// Components of a CERT record.
type DNSRecordCertData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64               `json:"type"`
	JSON dnsRecordCertDataJSON `json:"-"`
}

// dnsRecordCertDataJSON contains the JSON metadata for the struct
// [DNSRecordCertData]
type dnsRecordCertDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordCertData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordCertDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordCertType string

const (
	DNSRecordCertTypeCert DNSRecordCertType = "CERT"
)

func (r DNSRecordCertType) IsKnown() bool {
	switch r {
	case DNSRecordCertTypeCert:
		return true
	}
	return false
}

// Extra Cloudflare-specific information about the record.
type DNSRecordCertMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                `json:"source"`
	JSON   dnsRecordCertMetaJSON `json:"-"`
}

// dnsRecordCertMetaJSON contains the JSON metadata for the struct
// [DNSRecordCertMeta]
type dnsRecordCertMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordCertMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordCertMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.DNSRecordCertTTLNumber].
type DNSRecordCertTTL interface {
	ImplementsDNSDNSRecordCertTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordCertTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordCertTTLNumber(0)),
		},
	)
}

type DNSRecordCertTTLNumber float64

const (
	DNSRecordCertTTLNumber1 DNSRecordCertTTLNumber = 1
)

func (r DNSRecordCertTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordCertTTLNumber1:
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
	Content interface{} `json:"content,required"`
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

type DNSRecordSmimea struct {
	// Components of a SMIMEA record.
	Data DNSRecordSmimeaData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordSmimeaType `json:"type,required"`
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
	Meta DNSRecordSmimeaMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordSmimeaTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string              `json:"zone_name" format:"hostname"`
	JSON     dnsRecordSmimeaJSON `json:"-"`
}

// dnsRecordSmimeaJSON contains the JSON metadata for the struct [DNSRecordSmimea]
type dnsRecordSmimeaJSON struct {
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

func (r *DNSRecordSmimea) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordSmimeaJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordSmimea) implementsDNSDNSRecord() {}

// Components of a SMIMEA record.
type DNSRecordSmimeaData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                 `json:"usage"`
	JSON  dnsRecordSmimeaDataJSON `json:"-"`
}

// dnsRecordSmimeaDataJSON contains the JSON metadata for the struct
// [DNSRecordSmimeaData]
type dnsRecordSmimeaDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordSmimeaData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordSmimeaDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordSmimeaType string

const (
	DNSRecordSmimeaTypeSmimea DNSRecordSmimeaType = "SMIMEA"
)

func (r DNSRecordSmimeaType) IsKnown() bool {
	switch r {
	case DNSRecordSmimeaTypeSmimea:
		return true
	}
	return false
}

// Extra Cloudflare-specific information about the record.
type DNSRecordSmimeaMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                  `json:"source"`
	JSON   dnsRecordSmimeaMetaJSON `json:"-"`
}

// dnsRecordSmimeaMetaJSON contains the JSON metadata for the struct
// [DNSRecordSmimeaMeta]
type dnsRecordSmimeaMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordSmimeaMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordSmimeaMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.DNSRecordSmimeaTTLNumber].
type DNSRecordSmimeaTTL interface {
	ImplementsDNSDNSRecordSmimeaTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordSmimeaTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordSmimeaTTLNumber(0)),
		},
	)
}

type DNSRecordSmimeaTTLNumber float64

const (
	DNSRecordSmimeaTTLNumber1 DNSRecordSmimeaTTLNumber = 1
)

func (r DNSRecordSmimeaTTLNumber) IsKnown() bool {
	switch r {
	case DNSRecordSmimeaTTLNumber1:
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

type RecordNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordNewParamsType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Formatted URI content. See 'data' to set URI properties.
	Content param.Field[interface{}]         `json:"content"`
	Data    param.Field[RecordNewParamsData] `json:"data"`
	Meta    param.Field[RecordNewParamsMeta] `json:"meta"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordNewParamsTTL] `json:"ttl"`
}

func (r RecordNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordNewParamsType string

const (
	RecordNewParamsTypeURI    RecordNewParamsType = "URI"
	RecordNewParamsTypeTXT    RecordNewParamsType = "TXT"
	RecordNewParamsTypeTLSA   RecordNewParamsType = "TLSA"
	RecordNewParamsTypeSVCB   RecordNewParamsType = "SVCB"
	RecordNewParamsTypeSSHFP  RecordNewParamsType = "SSHFP"
	RecordNewParamsTypeSRV    RecordNewParamsType = "SRV"
	RecordNewParamsTypeSmimea RecordNewParamsType = "SMIMEA"
	RecordNewParamsTypePTR    RecordNewParamsType = "PTR"
	RecordNewParamsTypeNS     RecordNewParamsType = "NS"
	RecordNewParamsTypeNAPTR  RecordNewParamsType = "NAPTR"
	RecordNewParamsTypeMX     RecordNewParamsType = "MX"
	RecordNewParamsTypeLOC    RecordNewParamsType = "LOC"
	RecordNewParamsTypeHTTPS  RecordNewParamsType = "HTTPS"
	RecordNewParamsTypeDS     RecordNewParamsType = "DS"
	RecordNewParamsTypeDNSKEY RecordNewParamsType = "DNSKEY"
	RecordNewParamsTypeCNAME  RecordNewParamsType = "CNAME"
	RecordNewParamsTypeCert   RecordNewParamsType = "CERT"
	RecordNewParamsTypeCAA    RecordNewParamsType = "CAA"
	RecordNewParamsTypeAAAA   RecordNewParamsType = "AAAA"
	RecordNewParamsTypeA      RecordNewParamsType = "A"
)

func (r RecordNewParamsType) IsKnown() bool {
	switch r {
	case RecordNewParamsTypeURI, RecordNewParamsTypeTXT, RecordNewParamsTypeTLSA, RecordNewParamsTypeSVCB, RecordNewParamsTypeSSHFP, RecordNewParamsTypeSRV, RecordNewParamsTypeSmimea, RecordNewParamsTypePTR, RecordNewParamsTypeNS, RecordNewParamsTypeNAPTR, RecordNewParamsTypeMX, RecordNewParamsTypeLOC, RecordNewParamsTypeHTTPS, RecordNewParamsTypeDS, RecordNewParamsTypeDNSKEY, RecordNewParamsTypeCNAME, RecordNewParamsTypeCert, RecordNewParamsTypeCAA, RecordNewParamsTypeAAAA, RecordNewParamsTypeA:
		return true
	}
	return false
}

type RecordNewParamsData struct {
	// algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Altitude of location in meters.
	Altitude param.Field[float64] `json:"altitude"`
	// certificate.
	Certificate param.Field[string] `json:"certificate"`
	// The record content.
	Content param.Field[string] `json:"content"`
	// Digest.
	Digest param.Field[string] `json:"digest"`
	// Digest Type.
	DigestType param.Field[float64] `json:"digest_type"`
	// fingerprint.
	Fingerprint param.Field[string] `json:"fingerprint"`
	// Flags.
	Flags param.Field[interface{}] `json:"flags"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
	// Degrees of latitude.
	LatDegrees param.Field[float64] `json:"lat_degrees"`
	// Latitude direction.
	LatDirection param.Field[RecordNewParamsDataLatDirection] `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes param.Field[float64] `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds param.Field[float64] `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees param.Field[float64] `json:"long_degrees"`
	// Longitude direction.
	LongDirection param.Field[RecordNewParamsDataLongDirection] `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes param.Field[float64] `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds param.Field[float64] `json:"long_seconds"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// A valid hostname. Deprecated in favor of the regular 'name' outside the data
	// map. This data map field represents the remainder of the full 'name' after the
	// service and protocol.
	Name param.Field[string] `json:"name" format:"hostname"`
	// Order.
	Order param.Field[float64] `json:"order"`
	// The port of the service.
	Port param.Field[float64] `json:"port"`
	// Horizontal precision of location.
	PrecisionHorz param.Field[float64] `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert param.Field[float64] `json:"precision_vert"`
	// Preference.
	Preference param.Field[float64] `json:"preference"`
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// A valid protocol, prefixed with an underscore. Deprecated in favor of the
	// regular 'name' outside the data map. This data map field normally represents the
	// second label of that 'name'.
	Proto param.Field[string] `json:"proto"`
	// Protocol.
	Protocol param.Field[float64] `json:"protocol"`
	// Public Key.
	PublicKey param.Field[string] `json:"public_key"`
	// Regex.
	Regex param.Field[string] `json:"regex"`
	// Replacement.
	Replacement param.Field[string] `json:"replacement"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// A service type, prefixed with an underscore. Deprecated in favor of the regular
	// 'name' outside the data map. This data map field normally represents the first
	// label of that 'name'.
	Service param.Field[string] `json:"service"`
	// Size of location in meters.
	Size param.Field[float64] `json:"size"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag param.Field[string] `json:"tag"`
	// target.
	Target param.Field[string] `json:"target"`
	// type.
	Type param.Field[float64] `json:"type"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
	// value.
	Value param.Field[string] `json:"value"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r RecordNewParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Latitude direction.
type RecordNewParamsDataLatDirection string

const (
	RecordNewParamsDataLatDirectionN RecordNewParamsDataLatDirection = "N"
	RecordNewParamsDataLatDirectionS RecordNewParamsDataLatDirection = "S"
)

func (r RecordNewParamsDataLatDirection) IsKnown() bool {
	switch r {
	case RecordNewParamsDataLatDirectionN, RecordNewParamsDataLatDirectionS:
		return true
	}
	return false
}

// Longitude direction.
type RecordNewParamsDataLongDirection string

const (
	RecordNewParamsDataLongDirectionE RecordNewParamsDataLongDirection = "E"
	RecordNewParamsDataLongDirectionW RecordNewParamsDataLongDirection = "W"
)

func (r RecordNewParamsDataLongDirection) IsKnown() bool {
	switch r {
	case RecordNewParamsDataLongDirectionE, RecordNewParamsDataLongDirectionW:
		return true
	}
	return false
}

type RecordNewParamsMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded param.Field[bool] `json:"auto_added"`
	// Where the record originated from.
	Source param.Field[string] `json:"source"`
}

func (r RecordNewParamsMeta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat], [dns.RecordNewParamsTTLNumber].
type RecordNewParamsTTL interface {
	ImplementsDNSRecordNewParamsTTL()
}

type RecordNewParamsTTLNumber float64

const (
	RecordNewParamsTTLNumber1 RecordNewParamsTTLNumber = 1
)

func (r RecordNewParamsTTLNumber) IsKnown() bool {
	switch r {
	case RecordNewParamsTTLNumber1:
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

type RecordUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Formatted URI content. See 'data' to set URI properties.
	Content param.Field[interface{}]            `json:"content"`
	Data    param.Field[RecordUpdateParamsData] `json:"data"`
	Meta    param.Field[RecordUpdateParamsMeta] `json:"meta"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordUpdateParamsTTL] `json:"ttl"`
}

func (r RecordUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordUpdateParamsType string

const (
	RecordUpdateParamsTypeURI    RecordUpdateParamsType = "URI"
	RecordUpdateParamsTypeTXT    RecordUpdateParamsType = "TXT"
	RecordUpdateParamsTypeTLSA   RecordUpdateParamsType = "TLSA"
	RecordUpdateParamsTypeSVCB   RecordUpdateParamsType = "SVCB"
	RecordUpdateParamsTypeSSHFP  RecordUpdateParamsType = "SSHFP"
	RecordUpdateParamsTypeSRV    RecordUpdateParamsType = "SRV"
	RecordUpdateParamsTypeSmimea RecordUpdateParamsType = "SMIMEA"
	RecordUpdateParamsTypePTR    RecordUpdateParamsType = "PTR"
	RecordUpdateParamsTypeNS     RecordUpdateParamsType = "NS"
	RecordUpdateParamsTypeNAPTR  RecordUpdateParamsType = "NAPTR"
	RecordUpdateParamsTypeMX     RecordUpdateParamsType = "MX"
	RecordUpdateParamsTypeLOC    RecordUpdateParamsType = "LOC"
	RecordUpdateParamsTypeHTTPS  RecordUpdateParamsType = "HTTPS"
	RecordUpdateParamsTypeDS     RecordUpdateParamsType = "DS"
	RecordUpdateParamsTypeDNSKEY RecordUpdateParamsType = "DNSKEY"
	RecordUpdateParamsTypeCNAME  RecordUpdateParamsType = "CNAME"
	RecordUpdateParamsTypeCert   RecordUpdateParamsType = "CERT"
	RecordUpdateParamsTypeCAA    RecordUpdateParamsType = "CAA"
	RecordUpdateParamsTypeAAAA   RecordUpdateParamsType = "AAAA"
	RecordUpdateParamsTypeA      RecordUpdateParamsType = "A"
)

func (r RecordUpdateParamsType) IsKnown() bool {
	switch r {
	case RecordUpdateParamsTypeURI, RecordUpdateParamsTypeTXT, RecordUpdateParamsTypeTLSA, RecordUpdateParamsTypeSVCB, RecordUpdateParamsTypeSSHFP, RecordUpdateParamsTypeSRV, RecordUpdateParamsTypeSmimea, RecordUpdateParamsTypePTR, RecordUpdateParamsTypeNS, RecordUpdateParamsTypeNAPTR, RecordUpdateParamsTypeMX, RecordUpdateParamsTypeLOC, RecordUpdateParamsTypeHTTPS, RecordUpdateParamsTypeDS, RecordUpdateParamsTypeDNSKEY, RecordUpdateParamsTypeCNAME, RecordUpdateParamsTypeCert, RecordUpdateParamsTypeCAA, RecordUpdateParamsTypeAAAA, RecordUpdateParamsTypeA:
		return true
	}
	return false
}

type RecordUpdateParamsData struct {
	// algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Altitude of location in meters.
	Altitude param.Field[float64] `json:"altitude"`
	// certificate.
	Certificate param.Field[string] `json:"certificate"`
	// The record content.
	Content param.Field[string] `json:"content"`
	// Digest.
	Digest param.Field[string] `json:"digest"`
	// Digest Type.
	DigestType param.Field[float64] `json:"digest_type"`
	// fingerprint.
	Fingerprint param.Field[string] `json:"fingerprint"`
	// Flags.
	Flags param.Field[interface{}] `json:"flags"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
	// Degrees of latitude.
	LatDegrees param.Field[float64] `json:"lat_degrees"`
	// Latitude direction.
	LatDirection param.Field[RecordUpdateParamsDataLatDirection] `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes param.Field[float64] `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds param.Field[float64] `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees param.Field[float64] `json:"long_degrees"`
	// Longitude direction.
	LongDirection param.Field[RecordUpdateParamsDataLongDirection] `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes param.Field[float64] `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds param.Field[float64] `json:"long_seconds"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// A valid hostname. Deprecated in favor of the regular 'name' outside the data
	// map. This data map field represents the remainder of the full 'name' after the
	// service and protocol.
	Name param.Field[string] `json:"name" format:"hostname"`
	// Order.
	Order param.Field[float64] `json:"order"`
	// The port of the service.
	Port param.Field[float64] `json:"port"`
	// Horizontal precision of location.
	PrecisionHorz param.Field[float64] `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert param.Field[float64] `json:"precision_vert"`
	// Preference.
	Preference param.Field[float64] `json:"preference"`
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// A valid protocol, prefixed with an underscore. Deprecated in favor of the
	// regular 'name' outside the data map. This data map field normally represents the
	// second label of that 'name'.
	Proto param.Field[string] `json:"proto"`
	// Protocol.
	Protocol param.Field[float64] `json:"protocol"`
	// Public Key.
	PublicKey param.Field[string] `json:"public_key"`
	// Regex.
	Regex param.Field[string] `json:"regex"`
	// Replacement.
	Replacement param.Field[string] `json:"replacement"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// A service type, prefixed with an underscore. Deprecated in favor of the regular
	// 'name' outside the data map. This data map field normally represents the first
	// label of that 'name'.
	Service param.Field[string] `json:"service"`
	// Size of location in meters.
	Size param.Field[float64] `json:"size"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag param.Field[string] `json:"tag"`
	// target.
	Target param.Field[string] `json:"target"`
	// type.
	Type param.Field[float64] `json:"type"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
	// value.
	Value param.Field[string] `json:"value"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r RecordUpdateParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Latitude direction.
type RecordUpdateParamsDataLatDirection string

const (
	RecordUpdateParamsDataLatDirectionN RecordUpdateParamsDataLatDirection = "N"
	RecordUpdateParamsDataLatDirectionS RecordUpdateParamsDataLatDirection = "S"
)

func (r RecordUpdateParamsDataLatDirection) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDataLatDirectionN, RecordUpdateParamsDataLatDirectionS:
		return true
	}
	return false
}

// Longitude direction.
type RecordUpdateParamsDataLongDirection string

const (
	RecordUpdateParamsDataLongDirectionE RecordUpdateParamsDataLongDirection = "E"
	RecordUpdateParamsDataLongDirectionW RecordUpdateParamsDataLongDirection = "W"
)

func (r RecordUpdateParamsDataLongDirection) IsKnown() bool {
	switch r {
	case RecordUpdateParamsDataLongDirectionE, RecordUpdateParamsDataLongDirectionW:
		return true
	}
	return false
}

type RecordUpdateParamsMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded param.Field[bool] `json:"auto_added"`
	// Where the record originated from.
	Source param.Field[string] `json:"source"`
}

func (r RecordUpdateParamsMeta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat], [dns.RecordUpdateParamsTTLNumber].
type RecordUpdateParamsTTL interface {
	ImplementsDNSRecordUpdateParamsTTL()
}

type RecordUpdateParamsTTLNumber float64

const (
	RecordUpdateParamsTTLNumber1 RecordUpdateParamsTTLNumber = 1
)

func (r RecordUpdateParamsTTLNumber) IsKnown() bool {
	switch r {
	case RecordUpdateParamsTTLNumber1:
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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
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
	RecordListParamsTypeCert   RecordListParamsType = "CERT"
	RecordListParamsTypeCNAME  RecordListParamsType = "CNAME"
	RecordListParamsTypeDNSKEY RecordListParamsType = "DNSKEY"
	RecordListParamsTypeDS     RecordListParamsType = "DS"
	RecordListParamsTypeHTTPS  RecordListParamsType = "HTTPS"
	RecordListParamsTypeLOC    RecordListParamsType = "LOC"
	RecordListParamsTypeMX     RecordListParamsType = "MX"
	RecordListParamsTypeNAPTR  RecordListParamsType = "NAPTR"
	RecordListParamsTypeNS     RecordListParamsType = "NS"
	RecordListParamsTypePTR    RecordListParamsType = "PTR"
	RecordListParamsTypeSmimea RecordListParamsType = "SMIMEA"
	RecordListParamsTypeSRV    RecordListParamsType = "SRV"
	RecordListParamsTypeSSHFP  RecordListParamsType = "SSHFP"
	RecordListParamsTypeSVCB   RecordListParamsType = "SVCB"
	RecordListParamsTypeTLSA   RecordListParamsType = "TLSA"
	RecordListParamsTypeTXT    RecordListParamsType = "TXT"
	RecordListParamsTypeURI    RecordListParamsType = "URI"
)

func (r RecordListParamsType) IsKnown() bool {
	switch r {
	case RecordListParamsTypeA, RecordListParamsTypeAAAA, RecordListParamsTypeCAA, RecordListParamsTypeCert, RecordListParamsTypeCNAME, RecordListParamsTypeDNSKEY, RecordListParamsTypeDS, RecordListParamsTypeHTTPS, RecordListParamsTypeLOC, RecordListParamsTypeMX, RecordListParamsTypeNAPTR, RecordListParamsTypeNS, RecordListParamsTypePTR, RecordListParamsTypeSmimea, RecordListParamsTypeSRV, RecordListParamsTypeSSHFP, RecordListParamsTypeSVCB, RecordListParamsTypeTLSA, RecordListParamsTypeTXT, RecordListParamsTypeURI:
		return true
	}
	return false
}

type RecordDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
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

type RecordEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordEditParamsType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Formatted URI content. See 'data' to set URI properties.
	Content param.Field[interface{}]          `json:"content"`
	Data    param.Field[RecordEditParamsData] `json:"data"`
	Meta    param.Field[RecordEditParamsMeta] `json:"meta"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[RecordEditParamsTTL] `json:"ttl"`
}

func (r RecordEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type RecordEditParamsType string

const (
	RecordEditParamsTypeURI    RecordEditParamsType = "URI"
	RecordEditParamsTypeTXT    RecordEditParamsType = "TXT"
	RecordEditParamsTypeTLSA   RecordEditParamsType = "TLSA"
	RecordEditParamsTypeSVCB   RecordEditParamsType = "SVCB"
	RecordEditParamsTypeSSHFP  RecordEditParamsType = "SSHFP"
	RecordEditParamsTypeSRV    RecordEditParamsType = "SRV"
	RecordEditParamsTypeSmimea RecordEditParamsType = "SMIMEA"
	RecordEditParamsTypePTR    RecordEditParamsType = "PTR"
	RecordEditParamsTypeNS     RecordEditParamsType = "NS"
	RecordEditParamsTypeNAPTR  RecordEditParamsType = "NAPTR"
	RecordEditParamsTypeMX     RecordEditParamsType = "MX"
	RecordEditParamsTypeLOC    RecordEditParamsType = "LOC"
	RecordEditParamsTypeHTTPS  RecordEditParamsType = "HTTPS"
	RecordEditParamsTypeDS     RecordEditParamsType = "DS"
	RecordEditParamsTypeDNSKEY RecordEditParamsType = "DNSKEY"
	RecordEditParamsTypeCNAME  RecordEditParamsType = "CNAME"
	RecordEditParamsTypeCert   RecordEditParamsType = "CERT"
	RecordEditParamsTypeCAA    RecordEditParamsType = "CAA"
	RecordEditParamsTypeAAAA   RecordEditParamsType = "AAAA"
	RecordEditParamsTypeA      RecordEditParamsType = "A"
)

func (r RecordEditParamsType) IsKnown() bool {
	switch r {
	case RecordEditParamsTypeURI, RecordEditParamsTypeTXT, RecordEditParamsTypeTLSA, RecordEditParamsTypeSVCB, RecordEditParamsTypeSSHFP, RecordEditParamsTypeSRV, RecordEditParamsTypeSmimea, RecordEditParamsTypePTR, RecordEditParamsTypeNS, RecordEditParamsTypeNAPTR, RecordEditParamsTypeMX, RecordEditParamsTypeLOC, RecordEditParamsTypeHTTPS, RecordEditParamsTypeDS, RecordEditParamsTypeDNSKEY, RecordEditParamsTypeCNAME, RecordEditParamsTypeCert, RecordEditParamsTypeCAA, RecordEditParamsTypeAAAA, RecordEditParamsTypeA:
		return true
	}
	return false
}

type RecordEditParamsData struct {
	// algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Altitude of location in meters.
	Altitude param.Field[float64] `json:"altitude"`
	// certificate.
	Certificate param.Field[string] `json:"certificate"`
	// The record content.
	Content param.Field[string] `json:"content"`
	// Digest.
	Digest param.Field[string] `json:"digest"`
	// Digest Type.
	DigestType param.Field[float64] `json:"digest_type"`
	// fingerprint.
	Fingerprint param.Field[string] `json:"fingerprint"`
	// Flags.
	Flags param.Field[interface{}] `json:"flags"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
	// Degrees of latitude.
	LatDegrees param.Field[float64] `json:"lat_degrees"`
	// Latitude direction.
	LatDirection param.Field[RecordEditParamsDataLatDirection] `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes param.Field[float64] `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds param.Field[float64] `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees param.Field[float64] `json:"long_degrees"`
	// Longitude direction.
	LongDirection param.Field[RecordEditParamsDataLongDirection] `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes param.Field[float64] `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds param.Field[float64] `json:"long_seconds"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// A valid hostname. Deprecated in favor of the regular 'name' outside the data
	// map. This data map field represents the remainder of the full 'name' after the
	// service and protocol.
	Name param.Field[string] `json:"name" format:"hostname"`
	// Order.
	Order param.Field[float64] `json:"order"`
	// The port of the service.
	Port param.Field[float64] `json:"port"`
	// Horizontal precision of location.
	PrecisionHorz param.Field[float64] `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert param.Field[float64] `json:"precision_vert"`
	// Preference.
	Preference param.Field[float64] `json:"preference"`
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// A valid protocol, prefixed with an underscore. Deprecated in favor of the
	// regular 'name' outside the data map. This data map field normally represents the
	// second label of that 'name'.
	Proto param.Field[string] `json:"proto"`
	// Protocol.
	Protocol param.Field[float64] `json:"protocol"`
	// Public Key.
	PublicKey param.Field[string] `json:"public_key"`
	// Regex.
	Regex param.Field[string] `json:"regex"`
	// Replacement.
	Replacement param.Field[string] `json:"replacement"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// A service type, prefixed with an underscore. Deprecated in favor of the regular
	// 'name' outside the data map. This data map field normally represents the first
	// label of that 'name'.
	Service param.Field[string] `json:"service"`
	// Size of location in meters.
	Size param.Field[float64] `json:"size"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag param.Field[string] `json:"tag"`
	// target.
	Target param.Field[string] `json:"target"`
	// type.
	Type param.Field[float64] `json:"type"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
	// value.
	Value param.Field[string] `json:"value"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r RecordEditParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Latitude direction.
type RecordEditParamsDataLatDirection string

const (
	RecordEditParamsDataLatDirectionN RecordEditParamsDataLatDirection = "N"
	RecordEditParamsDataLatDirectionS RecordEditParamsDataLatDirection = "S"
)

func (r RecordEditParamsDataLatDirection) IsKnown() bool {
	switch r {
	case RecordEditParamsDataLatDirectionN, RecordEditParamsDataLatDirectionS:
		return true
	}
	return false
}

// Longitude direction.
type RecordEditParamsDataLongDirection string

const (
	RecordEditParamsDataLongDirectionE RecordEditParamsDataLongDirection = "E"
	RecordEditParamsDataLongDirectionW RecordEditParamsDataLongDirection = "W"
)

func (r RecordEditParamsDataLongDirection) IsKnown() bool {
	switch r {
	case RecordEditParamsDataLongDirectionE, RecordEditParamsDataLongDirectionW:
		return true
	}
	return false
}

type RecordEditParamsMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded param.Field[bool] `json:"auto_added"`
	// Where the record originated from.
	Source param.Field[string] `json:"source"`
}

func (r RecordEditParamsMeta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat], [dns.RecordEditParamsTTLNumber].
type RecordEditParamsTTL interface {
	ImplementsDNSRecordEditParamsTTL()
}

type RecordEditParamsTTLNumber float64

const (
	RecordEditParamsTTLNumber1 RecordEditParamsTTLNumber = 1
)

func (r RecordEditParamsTTLNumber) IsKnown() bool {
	switch r {
	case RecordEditParamsTTLNumber1:
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
	ZoneID param.Field[string] `path:"zone_id,required"`
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
