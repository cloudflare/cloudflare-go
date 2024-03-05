// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// DNSRecordService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewDNSRecordService] method instead.
type DNSRecordService struct {
	Options []option.RequestOption
}

// NewDNSRecordService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDNSRecordService(opts ...option.RequestOption) (r *DNSRecordService) {
	r = &DNSRecordService{}
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
func (r *DNSRecordService) New(ctx context.Context, params DNSRecordNewParams, opts ...option.RequestOption) (res *DNSRecordsDNSRecord, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSRecordNewResponseEnvelope
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
func (r *DNSRecordService) Update(ctx context.Context, dnsRecordID string, params DNSRecordUpdateParams, opts ...option.RequestOption) (res *DNSRecordsDNSRecord, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSRecordUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_records/%s", params.ZoneID, dnsRecordID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List, search, sort, and filter a zones' DNS records.
func (r *DNSRecordService) List(ctx context.Context, params DNSRecordListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[DNSRecordsDNSRecord], err error) {
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
func (r *DNSRecordService) ListAutoPaging(ctx context.Context, params DNSRecordListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[DNSRecordsDNSRecord] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete DNS Record
func (r *DNSRecordService) Delete(ctx context.Context, dnsRecordID string, body DNSRecordDeleteParams, opts ...option.RequestOption) (res *DNSRecordDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSRecordDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_records/%s", body.ZoneID, dnsRecordID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
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
func (r *DNSRecordService) Edit(ctx context.Context, dnsRecordID string, params DNSRecordEditParams, opts ...option.RequestOption) (res *DNSRecordsDNSRecord, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSRecordEditResponseEnvelope
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
func (r *DNSRecordService) Export(ctx context.Context, query DNSRecordExportParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := fmt.Sprintf("zones/%s/dns_records/export", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// DNS Record Details
func (r *DNSRecordService) Get(ctx context.Context, dnsRecordID string, query DNSRecordGetParams, opts ...option.RequestOption) (res *DNSRecordsDNSRecord, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSRecordGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_records/%s", query.ZoneID, dnsRecordID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
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
func (r *DNSRecordService) Import(ctx context.Context, params DNSRecordImportParams, opts ...option.RequestOption) (res *DNSRecordImportResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSRecordImportResponseEnvelope
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
func (r *DNSRecordService) Scan(ctx context.Context, body DNSRecordScanParams, opts ...option.RequestOption) (res *DNSRecordScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSRecordScanResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_records/scan", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [DNSRecordsDNSRecordDNSRecordsARecord],
// [DNSRecordsDNSRecordDNSRecordsAaaaRecord],
// [DNSRecordsDNSRecordDNSRecordsCaaRecord],
// [DNSRecordsDNSRecordDNSRecordsCertRecord],
// [DNSRecordsDNSRecordDNSRecordsCnameRecord],
// [DNSRecordsDNSRecordDNSRecordsDnskeyRecord],
// [DNSRecordsDNSRecordDNSRecordsDsRecord],
// [DNSRecordsDNSRecordDNSRecordsHTTPSRecord],
// [DNSRecordsDNSRecordDNSRecordsLocRecord],
// [DNSRecordsDNSRecordDNSRecordsMxRecord],
// [DNSRecordsDNSRecordDNSRecordsNaptrRecord],
// [DNSRecordsDNSRecordDNSRecordsNsRecord],
// [DNSRecordsDNSRecordDNSRecordsPtrRecord],
// [DNSRecordsDNSRecordDNSRecordsSmimeaRecord],
// [DNSRecordsDNSRecordDNSRecordsSrvRecord],
// [DNSRecordsDNSRecordDNSRecordsSshfpRecord],
// [DNSRecordsDNSRecordDNSRecordsSvcbRecord],
// [DNSRecordsDNSRecordDNSRecordsTlsaRecord],
// [DNSRecordsDNSRecordDNSRecordsTxtRecord] or
// [DNSRecordsDNSRecordDNSRecordsUriRecord].
type DNSRecordsDNSRecord interface {
	implementsDNSRecordsDNSRecord()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*DNSRecordsDNSRecord)(nil)).Elem(), "")
}

type DNSRecordsDNSRecordDNSRecordsARecord struct {
	// A valid IPv4 address.
	Content string `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordsDNSRecordDNSRecordsARecordType `json:"type,required"`
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
	Meta DNSRecordsDNSRecordDNSRecordsARecordMeta `json:"meta"`
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
	TTL DNSRecordsDNSRecordDNSRecordsARecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                   `json:"zone_name" format:"hostname"`
	JSON     dnsRecordsDNSRecordDNSRecordsARecordJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsARecordJSON contains the JSON metadata for the
// struct [DNSRecordsDNSRecordDNSRecordsARecord]
type dnsRecordsDNSRecordDNSRecordsARecordJSON struct {
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

func (r *DNSRecordsDNSRecordDNSRecordsARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordsDNSRecordDNSRecordsARecord) implementsDNSRecordsDNSRecord() {}

// Record type.
type DNSRecordsDNSRecordDNSRecordsARecordType string

const (
	DNSRecordsDNSRecordDNSRecordsARecordTypeA DNSRecordsDNSRecordDNSRecordsARecordType = "A"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordsDNSRecordDNSRecordsARecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                       `json:"source"`
	JSON   dnsRecordsDNSRecordDNSRecordsARecordMetaJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsARecordMetaJSON contains the JSON metadata for the
// struct [DNSRecordsDNSRecordDNSRecordsARecordMeta]
type dnsRecordsDNSRecordDNSRecordsARecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordsDNSRecordDNSRecordsARecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordsDNSRecordDNSRecordsARecordTTLNumber].
type DNSRecordsDNSRecordDNSRecordsARecordTTL interface {
	ImplementsDNSRecordsDNSRecordDNSRecordsARecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordsDNSRecordDNSRecordsARecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordsDNSRecordDNSRecordsARecordTTLNumber float64

const (
	DNSRecordsDNSRecordDNSRecordsARecordTTLNumber1 DNSRecordsDNSRecordDNSRecordsARecordTTLNumber = 1
)

type DNSRecordsDNSRecordDNSRecordsAaaaRecord struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordsDNSRecordDNSRecordsAaaaRecordType `json:"type,required"`
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
	Meta DNSRecordsDNSRecordDNSRecordsAaaaRecordMeta `json:"meta"`
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
	TTL DNSRecordsDNSRecordDNSRecordsAaaaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordsDNSRecordDNSRecordsAaaaRecordJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsAaaaRecordJSON contains the JSON metadata for the
// struct [DNSRecordsDNSRecordDNSRecordsAaaaRecord]
type dnsRecordsDNSRecordDNSRecordsAaaaRecordJSON struct {
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

func (r *DNSRecordsDNSRecordDNSRecordsAaaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordsDNSRecordDNSRecordsAaaaRecord) implementsDNSRecordsDNSRecord() {}

// Record type.
type DNSRecordsDNSRecordDNSRecordsAaaaRecordType string

const (
	DNSRecordsDNSRecordDNSRecordsAaaaRecordTypeAaaa DNSRecordsDNSRecordDNSRecordsAaaaRecordType = "AAAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordsDNSRecordDNSRecordsAaaaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordsDNSRecordDNSRecordsAaaaRecordMetaJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsAaaaRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordsDNSRecordDNSRecordsAaaaRecordMeta]
type dnsRecordsDNSRecordDNSRecordsAaaaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordsDNSRecordDNSRecordsAaaaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordsDNSRecordDNSRecordsAaaaRecordTTLNumber].
type DNSRecordsDNSRecordDNSRecordsAaaaRecordTTL interface {
	ImplementsDNSRecordsDNSRecordDNSRecordsAaaaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordsDNSRecordDNSRecordsAaaaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordsDNSRecordDNSRecordsAaaaRecordTTLNumber float64

const (
	DNSRecordsDNSRecordDNSRecordsAaaaRecordTTLNumber1 DNSRecordsDNSRecordDNSRecordsAaaaRecordTTLNumber = 1
)

type DNSRecordsDNSRecordDNSRecordsCaaRecord struct {
	// Components of a CAA record.
	Data DNSRecordsDNSRecordDNSRecordsCaaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordsDNSRecordDNSRecordsCaaRecordType `json:"type,required"`
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
	Meta DNSRecordsDNSRecordDNSRecordsCaaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordsDNSRecordDNSRecordsCaaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordsDNSRecordDNSRecordsCaaRecordJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsCaaRecordJSON contains the JSON metadata for the
// struct [DNSRecordsDNSRecordDNSRecordsCaaRecord]
type dnsRecordsDNSRecordDNSRecordsCaaRecordJSON struct {
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

func (r *DNSRecordsDNSRecordDNSRecordsCaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordsDNSRecordDNSRecordsCaaRecord) implementsDNSRecordsDNSRecord() {}

// Components of a CAA record.
type DNSRecordsDNSRecordDNSRecordsCaaRecordData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                                         `json:"value"`
	JSON  dnsRecordsDNSRecordDNSRecordsCaaRecordDataJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsCaaRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordsDNSRecordDNSRecordsCaaRecordData]
type dnsRecordsDNSRecordDNSRecordsCaaRecordDataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordsDNSRecordDNSRecordsCaaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordsDNSRecordDNSRecordsCaaRecordType string

const (
	DNSRecordsDNSRecordDNSRecordsCaaRecordTypeCaa DNSRecordsDNSRecordDNSRecordsCaaRecordType = "CAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordsDNSRecordDNSRecordsCaaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                         `json:"source"`
	JSON   dnsRecordsDNSRecordDNSRecordsCaaRecordMetaJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsCaaRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordsDNSRecordDNSRecordsCaaRecordMeta]
type dnsRecordsDNSRecordDNSRecordsCaaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordsDNSRecordDNSRecordsCaaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordsDNSRecordDNSRecordsCaaRecordTTLNumber].
type DNSRecordsDNSRecordDNSRecordsCaaRecordTTL interface {
	ImplementsDNSRecordsDNSRecordDNSRecordsCaaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordsDNSRecordDNSRecordsCaaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordsDNSRecordDNSRecordsCaaRecordTTLNumber float64

const (
	DNSRecordsDNSRecordDNSRecordsCaaRecordTTLNumber1 DNSRecordsDNSRecordDNSRecordsCaaRecordTTLNumber = 1
)

type DNSRecordsDNSRecordDNSRecordsCertRecord struct {
	// Components of a CERT record.
	Data DNSRecordsDNSRecordDNSRecordsCertRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordsDNSRecordDNSRecordsCertRecordType `json:"type,required"`
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
	Meta DNSRecordsDNSRecordDNSRecordsCertRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordsDNSRecordDNSRecordsCertRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordsDNSRecordDNSRecordsCertRecordJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsCertRecordJSON contains the JSON metadata for the
// struct [DNSRecordsDNSRecordDNSRecordsCertRecord]
type dnsRecordsDNSRecordDNSRecordsCertRecordJSON struct {
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

func (r *DNSRecordsDNSRecordDNSRecordsCertRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordsDNSRecordDNSRecordsCertRecord) implementsDNSRecordsDNSRecord() {}

// Components of a CERT record.
type DNSRecordsDNSRecordDNSRecordsCertRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64                                         `json:"type"`
	JSON dnsRecordsDNSRecordDNSRecordsCertRecordDataJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsCertRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordsDNSRecordDNSRecordsCertRecordData]
type dnsRecordsDNSRecordDNSRecordsCertRecordDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordsDNSRecordDNSRecordsCertRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordsDNSRecordDNSRecordsCertRecordType string

const (
	DNSRecordsDNSRecordDNSRecordsCertRecordTypeCert DNSRecordsDNSRecordDNSRecordsCertRecordType = "CERT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordsDNSRecordDNSRecordsCertRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordsDNSRecordDNSRecordsCertRecordMetaJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsCertRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordsDNSRecordDNSRecordsCertRecordMeta]
type dnsRecordsDNSRecordDNSRecordsCertRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordsDNSRecordDNSRecordsCertRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordsDNSRecordDNSRecordsCertRecordTTLNumber].
type DNSRecordsDNSRecordDNSRecordsCertRecordTTL interface {
	ImplementsDNSRecordsDNSRecordDNSRecordsCertRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordsDNSRecordDNSRecordsCertRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordsDNSRecordDNSRecordsCertRecordTTLNumber float64

const (
	DNSRecordsDNSRecordDNSRecordsCertRecordTTLNumber1 DNSRecordsDNSRecordDNSRecordsCertRecordTTLNumber = 1
)

type DNSRecordsDNSRecordDNSRecordsCnameRecord struct {
	// A valid hostname. Must not match the record's name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordsDNSRecordDNSRecordsCnameRecordType `json:"type,required"`
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
	Meta DNSRecordsDNSRecordDNSRecordsCnameRecordMeta `json:"meta"`
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
	TTL DNSRecordsDNSRecordDNSRecordsCnameRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordsDNSRecordDNSRecordsCnameRecordJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsCnameRecordJSON contains the JSON metadata for the
// struct [DNSRecordsDNSRecordDNSRecordsCnameRecord]
type dnsRecordsDNSRecordDNSRecordsCnameRecordJSON struct {
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

func (r *DNSRecordsDNSRecordDNSRecordsCnameRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordsDNSRecordDNSRecordsCnameRecord) implementsDNSRecordsDNSRecord() {}

// Record type.
type DNSRecordsDNSRecordDNSRecordsCnameRecordType string

const (
	DNSRecordsDNSRecordDNSRecordsCnameRecordTypeCname DNSRecordsDNSRecordDNSRecordsCnameRecordType = "CNAME"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordsDNSRecordDNSRecordsCnameRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordsDNSRecordDNSRecordsCnameRecordMetaJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsCnameRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordsDNSRecordDNSRecordsCnameRecordMeta]
type dnsRecordsDNSRecordDNSRecordsCnameRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordsDNSRecordDNSRecordsCnameRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordsDNSRecordDNSRecordsCnameRecordTTLNumber].
type DNSRecordsDNSRecordDNSRecordsCnameRecordTTL interface {
	ImplementsDNSRecordsDNSRecordDNSRecordsCnameRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordsDNSRecordDNSRecordsCnameRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordsDNSRecordDNSRecordsCnameRecordTTLNumber float64

const (
	DNSRecordsDNSRecordDNSRecordsCnameRecordTTLNumber1 DNSRecordsDNSRecordDNSRecordsCnameRecordTTLNumber = 1
)

type DNSRecordsDNSRecordDNSRecordsDnskeyRecord struct {
	// Components of a DNSKEY record.
	Data DNSRecordsDNSRecordDNSRecordsDnskeyRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordsDNSRecordDNSRecordsDnskeyRecordType `json:"type,required"`
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
	Meta DNSRecordsDNSRecordDNSRecordsDnskeyRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordsDNSRecordDNSRecordsDnskeyRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordsDNSRecordDNSRecordsDnskeyRecordJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsDnskeyRecordJSON contains the JSON metadata for the
// struct [DNSRecordsDNSRecordDNSRecordsDnskeyRecord]
type dnsRecordsDNSRecordDNSRecordsDnskeyRecordJSON struct {
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

func (r *DNSRecordsDNSRecordDNSRecordsDnskeyRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordsDNSRecordDNSRecordsDnskeyRecord) implementsDNSRecordsDNSRecord() {}

// Components of a DNSKEY record.
type DNSRecordsDNSRecordDNSRecordsDnskeyRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                                            `json:"public_key"`
	JSON      dnsRecordsDNSRecordDNSRecordsDnskeyRecordDataJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsDnskeyRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordsDNSRecordDNSRecordsDnskeyRecordData]
type dnsRecordsDNSRecordDNSRecordsDnskeyRecordDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordsDNSRecordDNSRecordsDnskeyRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordsDNSRecordDNSRecordsDnskeyRecordType string

const (
	DNSRecordsDNSRecordDNSRecordsDnskeyRecordTypeDnskey DNSRecordsDNSRecordDNSRecordsDnskeyRecordType = "DNSKEY"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordsDNSRecordDNSRecordsDnskeyRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordsDNSRecordDNSRecordsDnskeyRecordMetaJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsDnskeyRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordsDNSRecordDNSRecordsDnskeyRecordMeta]
type dnsRecordsDNSRecordDNSRecordsDnskeyRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordsDNSRecordDNSRecordsDnskeyRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordsDNSRecordDNSRecordsDnskeyRecordTTLNumber].
type DNSRecordsDNSRecordDNSRecordsDnskeyRecordTTL interface {
	ImplementsDNSRecordsDNSRecordDNSRecordsDnskeyRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordsDNSRecordDNSRecordsDnskeyRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordsDNSRecordDNSRecordsDnskeyRecordTTLNumber float64

const (
	DNSRecordsDNSRecordDNSRecordsDnskeyRecordTTLNumber1 DNSRecordsDNSRecordDNSRecordsDnskeyRecordTTLNumber = 1
)

type DNSRecordsDNSRecordDNSRecordsDsRecord struct {
	// Components of a DS record.
	Data DNSRecordsDNSRecordDNSRecordsDsRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordsDNSRecordDNSRecordsDsRecordType `json:"type,required"`
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
	Meta DNSRecordsDNSRecordDNSRecordsDsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordsDNSRecordDNSRecordsDsRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                    `json:"zone_name" format:"hostname"`
	JSON     dnsRecordsDNSRecordDNSRecordsDsRecordJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsDsRecordJSON contains the JSON metadata for the
// struct [DNSRecordsDNSRecordDNSRecordsDsRecord]
type dnsRecordsDNSRecordDNSRecordsDsRecordJSON struct {
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

func (r *DNSRecordsDNSRecordDNSRecordsDsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordsDNSRecordDNSRecordsDsRecord) implementsDNSRecordsDNSRecord() {}

// Components of a DS record.
type DNSRecordsDNSRecordDNSRecordsDsRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                                       `json:"key_tag"`
	JSON   dnsRecordsDNSRecordDNSRecordsDsRecordDataJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsDsRecordDataJSON contains the JSON metadata for the
// struct [DNSRecordsDNSRecordDNSRecordsDsRecordData]
type dnsRecordsDNSRecordDNSRecordsDsRecordDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordsDNSRecordDNSRecordsDsRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordsDNSRecordDNSRecordsDsRecordType string

const (
	DNSRecordsDNSRecordDNSRecordsDsRecordTypeDs DNSRecordsDNSRecordDNSRecordsDsRecordType = "DS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordsDNSRecordDNSRecordsDsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                        `json:"source"`
	JSON   dnsRecordsDNSRecordDNSRecordsDsRecordMetaJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsDsRecordMetaJSON contains the JSON metadata for the
// struct [DNSRecordsDNSRecordDNSRecordsDsRecordMeta]
type dnsRecordsDNSRecordDNSRecordsDsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordsDNSRecordDNSRecordsDsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordsDNSRecordDNSRecordsDsRecordTTLNumber].
type DNSRecordsDNSRecordDNSRecordsDsRecordTTL interface {
	ImplementsDNSRecordsDNSRecordDNSRecordsDsRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordsDNSRecordDNSRecordsDsRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordsDNSRecordDNSRecordsDsRecordTTLNumber float64

const (
	DNSRecordsDNSRecordDNSRecordsDsRecordTTLNumber1 DNSRecordsDNSRecordDNSRecordsDsRecordTTLNumber = 1
)

type DNSRecordsDNSRecordDNSRecordsHTTPSRecord struct {
	// Components of a HTTPS record.
	Data DNSRecordsDNSRecordDNSRecordsHTTPSRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordsDNSRecordDNSRecordsHTTPSRecordType `json:"type,required"`
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
	Meta DNSRecordsDNSRecordDNSRecordsHTTPSRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordsDNSRecordDNSRecordsHTTPSRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordsDNSRecordDNSRecordsHTTPSRecordJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsHTTPSRecordJSON contains the JSON metadata for the
// struct [DNSRecordsDNSRecordDNSRecordsHTTPSRecord]
type dnsRecordsDNSRecordDNSRecordsHTTPSRecordJSON struct {
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

func (r *DNSRecordsDNSRecordDNSRecordsHTTPSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordsDNSRecordDNSRecordsHTTPSRecord) implementsDNSRecordsDNSRecord() {}

// Components of a HTTPS record.
type DNSRecordsDNSRecordDNSRecordsHTTPSRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                           `json:"value"`
	JSON  dnsRecordsDNSRecordDNSRecordsHTTPSRecordDataJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsHTTPSRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordsDNSRecordDNSRecordsHTTPSRecordData]
type dnsRecordsDNSRecordDNSRecordsHTTPSRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordsDNSRecordDNSRecordsHTTPSRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordsDNSRecordDNSRecordsHTTPSRecordType string

const (
	DNSRecordsDNSRecordDNSRecordsHTTPSRecordTypeHTTPS DNSRecordsDNSRecordDNSRecordsHTTPSRecordType = "HTTPS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordsDNSRecordDNSRecordsHTTPSRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordsDNSRecordDNSRecordsHTTPSRecordMetaJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsHTTPSRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordsDNSRecordDNSRecordsHTTPSRecordMeta]
type dnsRecordsDNSRecordDNSRecordsHTTPSRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordsDNSRecordDNSRecordsHTTPSRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordsDNSRecordDNSRecordsHTTPSRecordTTLNumber].
type DNSRecordsDNSRecordDNSRecordsHTTPSRecordTTL interface {
	ImplementsDNSRecordsDNSRecordDNSRecordsHTTPSRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordsDNSRecordDNSRecordsHTTPSRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordsDNSRecordDNSRecordsHTTPSRecordTTLNumber float64

const (
	DNSRecordsDNSRecordDNSRecordsHTTPSRecordTTLNumber1 DNSRecordsDNSRecordDNSRecordsHTTPSRecordTTLNumber = 1
)

type DNSRecordsDNSRecordDNSRecordsLocRecord struct {
	// Components of a LOC record.
	Data DNSRecordsDNSRecordDNSRecordsLocRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordsDNSRecordDNSRecordsLocRecordType `json:"type,required"`
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
	Meta DNSRecordsDNSRecordDNSRecordsLocRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordsDNSRecordDNSRecordsLocRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordsDNSRecordDNSRecordsLocRecordJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsLocRecordJSON contains the JSON metadata for the
// struct [DNSRecordsDNSRecordDNSRecordsLocRecord]
type dnsRecordsDNSRecordDNSRecordsLocRecordJSON struct {
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

func (r *DNSRecordsDNSRecordDNSRecordsLocRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordsDNSRecordDNSRecordsLocRecord) implementsDNSRecordsDNSRecord() {}

// Components of a LOC record.
type DNSRecordsDNSRecordDNSRecordsLocRecordData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection DNSRecordsDNSRecordDNSRecordsLocRecordDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection DNSRecordsDNSRecordDNSRecordsLocRecordDataLongDirection `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes float64 `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds float64 `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz float64 `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert float64 `json:"precision_vert"`
	// Size of location in meters.
	Size float64                                        `json:"size"`
	JSON dnsRecordsDNSRecordDNSRecordsLocRecordDataJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsLocRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordsDNSRecordDNSRecordsLocRecordData]
type dnsRecordsDNSRecordDNSRecordsLocRecordDataJSON struct {
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

func (r *DNSRecordsDNSRecordDNSRecordsLocRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Latitude direction.
type DNSRecordsDNSRecordDNSRecordsLocRecordDataLatDirection string

const (
	DNSRecordsDNSRecordDNSRecordsLocRecordDataLatDirectionN DNSRecordsDNSRecordDNSRecordsLocRecordDataLatDirection = "N"
	DNSRecordsDNSRecordDNSRecordsLocRecordDataLatDirectionS DNSRecordsDNSRecordDNSRecordsLocRecordDataLatDirection = "S"
)

// Longitude direction.
type DNSRecordsDNSRecordDNSRecordsLocRecordDataLongDirection string

const (
	DNSRecordsDNSRecordDNSRecordsLocRecordDataLongDirectionE DNSRecordsDNSRecordDNSRecordsLocRecordDataLongDirection = "E"
	DNSRecordsDNSRecordDNSRecordsLocRecordDataLongDirectionW DNSRecordsDNSRecordDNSRecordsLocRecordDataLongDirection = "W"
)

// Record type.
type DNSRecordsDNSRecordDNSRecordsLocRecordType string

const (
	DNSRecordsDNSRecordDNSRecordsLocRecordTypeLoc DNSRecordsDNSRecordDNSRecordsLocRecordType = "LOC"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordsDNSRecordDNSRecordsLocRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                         `json:"source"`
	JSON   dnsRecordsDNSRecordDNSRecordsLocRecordMetaJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsLocRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordsDNSRecordDNSRecordsLocRecordMeta]
type dnsRecordsDNSRecordDNSRecordsLocRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordsDNSRecordDNSRecordsLocRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordsDNSRecordDNSRecordsLocRecordTTLNumber].
type DNSRecordsDNSRecordDNSRecordsLocRecordTTL interface {
	ImplementsDNSRecordsDNSRecordDNSRecordsLocRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordsDNSRecordDNSRecordsLocRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordsDNSRecordDNSRecordsLocRecordTTLNumber float64

const (
	DNSRecordsDNSRecordDNSRecordsLocRecordTTLNumber1 DNSRecordsDNSRecordDNSRecordsLocRecordTTLNumber = 1
)

type DNSRecordsDNSRecordDNSRecordsMxRecord struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordsDNSRecordDNSRecordsMxRecordType `json:"type,required"`
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
	Meta DNSRecordsDNSRecordDNSRecordsMxRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordsDNSRecordDNSRecordsMxRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                    `json:"zone_name" format:"hostname"`
	JSON     dnsRecordsDNSRecordDNSRecordsMxRecordJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsMxRecordJSON contains the JSON metadata for the
// struct [DNSRecordsDNSRecordDNSRecordsMxRecord]
type dnsRecordsDNSRecordDNSRecordsMxRecordJSON struct {
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

func (r *DNSRecordsDNSRecordDNSRecordsMxRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordsDNSRecordDNSRecordsMxRecord) implementsDNSRecordsDNSRecord() {}

// Record type.
type DNSRecordsDNSRecordDNSRecordsMxRecordType string

const (
	DNSRecordsDNSRecordDNSRecordsMxRecordTypeMx DNSRecordsDNSRecordDNSRecordsMxRecordType = "MX"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordsDNSRecordDNSRecordsMxRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                        `json:"source"`
	JSON   dnsRecordsDNSRecordDNSRecordsMxRecordMetaJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsMxRecordMetaJSON contains the JSON metadata for the
// struct [DNSRecordsDNSRecordDNSRecordsMxRecordMeta]
type dnsRecordsDNSRecordDNSRecordsMxRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordsDNSRecordDNSRecordsMxRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordsDNSRecordDNSRecordsMxRecordTTLNumber].
type DNSRecordsDNSRecordDNSRecordsMxRecordTTL interface {
	ImplementsDNSRecordsDNSRecordDNSRecordsMxRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordsDNSRecordDNSRecordsMxRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordsDNSRecordDNSRecordsMxRecordTTLNumber float64

const (
	DNSRecordsDNSRecordDNSRecordsMxRecordTTLNumber1 DNSRecordsDNSRecordDNSRecordsMxRecordTTLNumber = 1
)

type DNSRecordsDNSRecordDNSRecordsNaptrRecord struct {
	// Components of a NAPTR record.
	Data DNSRecordsDNSRecordDNSRecordsNaptrRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordsDNSRecordDNSRecordsNaptrRecordType `json:"type,required"`
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
	Meta DNSRecordsDNSRecordDNSRecordsNaptrRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordsDNSRecordDNSRecordsNaptrRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordsDNSRecordDNSRecordsNaptrRecordJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsNaptrRecordJSON contains the JSON metadata for the
// struct [DNSRecordsDNSRecordDNSRecordsNaptrRecord]
type dnsRecordsDNSRecordDNSRecordsNaptrRecordJSON struct {
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

func (r *DNSRecordsDNSRecordDNSRecordsNaptrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordsDNSRecordDNSRecordsNaptrRecord) implementsDNSRecordsDNSRecord() {}

// Components of a NAPTR record.
type DNSRecordsDNSRecordDNSRecordsNaptrRecordData struct {
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
	Service string                                           `json:"service"`
	JSON    dnsRecordsDNSRecordDNSRecordsNaptrRecordDataJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsNaptrRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordsDNSRecordDNSRecordsNaptrRecordData]
type dnsRecordsDNSRecordDNSRecordsNaptrRecordDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordsDNSRecordDNSRecordsNaptrRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordsDNSRecordDNSRecordsNaptrRecordType string

const (
	DNSRecordsDNSRecordDNSRecordsNaptrRecordTypeNaptr DNSRecordsDNSRecordDNSRecordsNaptrRecordType = "NAPTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordsDNSRecordDNSRecordsNaptrRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordsDNSRecordDNSRecordsNaptrRecordMetaJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsNaptrRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordsDNSRecordDNSRecordsNaptrRecordMeta]
type dnsRecordsDNSRecordDNSRecordsNaptrRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordsDNSRecordDNSRecordsNaptrRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordsDNSRecordDNSRecordsNaptrRecordTTLNumber].
type DNSRecordsDNSRecordDNSRecordsNaptrRecordTTL interface {
	ImplementsDNSRecordsDNSRecordDNSRecordsNaptrRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordsDNSRecordDNSRecordsNaptrRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordsDNSRecordDNSRecordsNaptrRecordTTLNumber float64

const (
	DNSRecordsDNSRecordDNSRecordsNaptrRecordTTLNumber1 DNSRecordsDNSRecordDNSRecordsNaptrRecordTTLNumber = 1
)

type DNSRecordsDNSRecordDNSRecordsNsRecord struct {
	// A valid name server host name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordsDNSRecordDNSRecordsNsRecordType `json:"type,required"`
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
	Meta DNSRecordsDNSRecordDNSRecordsNsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordsDNSRecordDNSRecordsNsRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                    `json:"zone_name" format:"hostname"`
	JSON     dnsRecordsDNSRecordDNSRecordsNsRecordJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsNsRecordJSON contains the JSON metadata for the
// struct [DNSRecordsDNSRecordDNSRecordsNsRecord]
type dnsRecordsDNSRecordDNSRecordsNsRecordJSON struct {
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

func (r *DNSRecordsDNSRecordDNSRecordsNsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordsDNSRecordDNSRecordsNsRecord) implementsDNSRecordsDNSRecord() {}

// Record type.
type DNSRecordsDNSRecordDNSRecordsNsRecordType string

const (
	DNSRecordsDNSRecordDNSRecordsNsRecordTypeNs DNSRecordsDNSRecordDNSRecordsNsRecordType = "NS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordsDNSRecordDNSRecordsNsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                        `json:"source"`
	JSON   dnsRecordsDNSRecordDNSRecordsNsRecordMetaJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsNsRecordMetaJSON contains the JSON metadata for the
// struct [DNSRecordsDNSRecordDNSRecordsNsRecordMeta]
type dnsRecordsDNSRecordDNSRecordsNsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordsDNSRecordDNSRecordsNsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordsDNSRecordDNSRecordsNsRecordTTLNumber].
type DNSRecordsDNSRecordDNSRecordsNsRecordTTL interface {
	ImplementsDNSRecordsDNSRecordDNSRecordsNsRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordsDNSRecordDNSRecordsNsRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordsDNSRecordDNSRecordsNsRecordTTLNumber float64

const (
	DNSRecordsDNSRecordDNSRecordsNsRecordTTLNumber1 DNSRecordsDNSRecordDNSRecordsNsRecordTTLNumber = 1
)

type DNSRecordsDNSRecordDNSRecordsPtrRecord struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordsDNSRecordDNSRecordsPtrRecordType `json:"type,required"`
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
	Meta DNSRecordsDNSRecordDNSRecordsPtrRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordsDNSRecordDNSRecordsPtrRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordsDNSRecordDNSRecordsPtrRecordJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsPtrRecordJSON contains the JSON metadata for the
// struct [DNSRecordsDNSRecordDNSRecordsPtrRecord]
type dnsRecordsDNSRecordDNSRecordsPtrRecordJSON struct {
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

func (r *DNSRecordsDNSRecordDNSRecordsPtrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordsDNSRecordDNSRecordsPtrRecord) implementsDNSRecordsDNSRecord() {}

// Record type.
type DNSRecordsDNSRecordDNSRecordsPtrRecordType string

const (
	DNSRecordsDNSRecordDNSRecordsPtrRecordTypePtr DNSRecordsDNSRecordDNSRecordsPtrRecordType = "PTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordsDNSRecordDNSRecordsPtrRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                         `json:"source"`
	JSON   dnsRecordsDNSRecordDNSRecordsPtrRecordMetaJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsPtrRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordsDNSRecordDNSRecordsPtrRecordMeta]
type dnsRecordsDNSRecordDNSRecordsPtrRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordsDNSRecordDNSRecordsPtrRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordsDNSRecordDNSRecordsPtrRecordTTLNumber].
type DNSRecordsDNSRecordDNSRecordsPtrRecordTTL interface {
	ImplementsDNSRecordsDNSRecordDNSRecordsPtrRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordsDNSRecordDNSRecordsPtrRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordsDNSRecordDNSRecordsPtrRecordTTLNumber float64

const (
	DNSRecordsDNSRecordDNSRecordsPtrRecordTTLNumber1 DNSRecordsDNSRecordDNSRecordsPtrRecordTTLNumber = 1
)

type DNSRecordsDNSRecordDNSRecordsSmimeaRecord struct {
	// Components of a SMIMEA record.
	Data DNSRecordsDNSRecordDNSRecordsSmimeaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordsDNSRecordDNSRecordsSmimeaRecordType `json:"type,required"`
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
	Meta DNSRecordsDNSRecordDNSRecordsSmimeaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordsDNSRecordDNSRecordsSmimeaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordsDNSRecordDNSRecordsSmimeaRecordJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsSmimeaRecordJSON contains the JSON metadata for the
// struct [DNSRecordsDNSRecordDNSRecordsSmimeaRecord]
type dnsRecordsDNSRecordDNSRecordsSmimeaRecordJSON struct {
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

func (r *DNSRecordsDNSRecordDNSRecordsSmimeaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordsDNSRecordDNSRecordsSmimeaRecord) implementsDNSRecordsDNSRecord() {}

// Components of a SMIMEA record.
type DNSRecordsDNSRecordDNSRecordsSmimeaRecordData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                           `json:"usage"`
	JSON  dnsRecordsDNSRecordDNSRecordsSmimeaRecordDataJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsSmimeaRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordsDNSRecordDNSRecordsSmimeaRecordData]
type dnsRecordsDNSRecordDNSRecordsSmimeaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordsDNSRecordDNSRecordsSmimeaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordsDNSRecordDNSRecordsSmimeaRecordType string

const (
	DNSRecordsDNSRecordDNSRecordsSmimeaRecordTypeSmimea DNSRecordsDNSRecordDNSRecordsSmimeaRecordType = "SMIMEA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordsDNSRecordDNSRecordsSmimeaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordsDNSRecordDNSRecordsSmimeaRecordMetaJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsSmimeaRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordsDNSRecordDNSRecordsSmimeaRecordMeta]
type dnsRecordsDNSRecordDNSRecordsSmimeaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordsDNSRecordDNSRecordsSmimeaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordsDNSRecordDNSRecordsSmimeaRecordTTLNumber].
type DNSRecordsDNSRecordDNSRecordsSmimeaRecordTTL interface {
	ImplementsDNSRecordsDNSRecordDNSRecordsSmimeaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordsDNSRecordDNSRecordsSmimeaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordsDNSRecordDNSRecordsSmimeaRecordTTLNumber float64

const (
	DNSRecordsDNSRecordDNSRecordsSmimeaRecordTTLNumber1 DNSRecordsDNSRecordDNSRecordsSmimeaRecordTTLNumber = 1
)

type DNSRecordsDNSRecordDNSRecordsSrvRecord struct {
	// Components of a SRV record.
	Data DNSRecordsDNSRecordDNSRecordsSrvRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordsDNSRecordDNSRecordsSrvRecordType `json:"type,required"`
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
	Meta DNSRecordsDNSRecordDNSRecordsSrvRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordsDNSRecordDNSRecordsSrvRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordsDNSRecordDNSRecordsSrvRecordJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsSrvRecordJSON contains the JSON metadata for the
// struct [DNSRecordsDNSRecordDNSRecordsSrvRecord]
type dnsRecordsDNSRecordDNSRecordsSrvRecordJSON struct {
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

func (r *DNSRecordsDNSRecordDNSRecordsSrvRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordsDNSRecordDNSRecordsSrvRecord) implementsDNSRecordsDNSRecord() {}

// Components of a SRV record.
type DNSRecordsDNSRecordDNSRecordsSrvRecordData struct {
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
	Weight float64                                        `json:"weight"`
	JSON   dnsRecordsDNSRecordDNSRecordsSrvRecordDataJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsSrvRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordsDNSRecordDNSRecordsSrvRecordData]
type dnsRecordsDNSRecordDNSRecordsSrvRecordDataJSON struct {
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

func (r *DNSRecordsDNSRecordDNSRecordsSrvRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordsDNSRecordDNSRecordsSrvRecordType string

const (
	DNSRecordsDNSRecordDNSRecordsSrvRecordTypeSrv DNSRecordsDNSRecordDNSRecordsSrvRecordType = "SRV"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordsDNSRecordDNSRecordsSrvRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                         `json:"source"`
	JSON   dnsRecordsDNSRecordDNSRecordsSrvRecordMetaJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsSrvRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordsDNSRecordDNSRecordsSrvRecordMeta]
type dnsRecordsDNSRecordDNSRecordsSrvRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordsDNSRecordDNSRecordsSrvRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordsDNSRecordDNSRecordsSrvRecordTTLNumber].
type DNSRecordsDNSRecordDNSRecordsSrvRecordTTL interface {
	ImplementsDNSRecordsDNSRecordDNSRecordsSrvRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordsDNSRecordDNSRecordsSrvRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordsDNSRecordDNSRecordsSrvRecordTTLNumber float64

const (
	DNSRecordsDNSRecordDNSRecordsSrvRecordTTLNumber1 DNSRecordsDNSRecordDNSRecordsSrvRecordTTLNumber = 1
)

type DNSRecordsDNSRecordDNSRecordsSshfpRecord struct {
	// Components of a SSHFP record.
	Data DNSRecordsDNSRecordDNSRecordsSshfpRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordsDNSRecordDNSRecordsSshfpRecordType `json:"type,required"`
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
	Meta DNSRecordsDNSRecordDNSRecordsSshfpRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordsDNSRecordDNSRecordsSshfpRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordsDNSRecordDNSRecordsSshfpRecordJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsSshfpRecordJSON contains the JSON metadata for the
// struct [DNSRecordsDNSRecordDNSRecordsSshfpRecord]
type dnsRecordsDNSRecordDNSRecordsSshfpRecordJSON struct {
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

func (r *DNSRecordsDNSRecordDNSRecordsSshfpRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordsDNSRecordDNSRecordsSshfpRecord) implementsDNSRecordsDNSRecord() {}

// Components of a SSHFP record.
type DNSRecordsDNSRecordDNSRecordsSshfpRecordData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                                          `json:"type"`
	JSON dnsRecordsDNSRecordDNSRecordsSshfpRecordDataJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsSshfpRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordsDNSRecordDNSRecordsSshfpRecordData]
type dnsRecordsDNSRecordDNSRecordsSshfpRecordDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordsDNSRecordDNSRecordsSshfpRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordsDNSRecordDNSRecordsSshfpRecordType string

const (
	DNSRecordsDNSRecordDNSRecordsSshfpRecordTypeSshfp DNSRecordsDNSRecordDNSRecordsSshfpRecordType = "SSHFP"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordsDNSRecordDNSRecordsSshfpRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordsDNSRecordDNSRecordsSshfpRecordMetaJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsSshfpRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordsDNSRecordDNSRecordsSshfpRecordMeta]
type dnsRecordsDNSRecordDNSRecordsSshfpRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordsDNSRecordDNSRecordsSshfpRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordsDNSRecordDNSRecordsSshfpRecordTTLNumber].
type DNSRecordsDNSRecordDNSRecordsSshfpRecordTTL interface {
	ImplementsDNSRecordsDNSRecordDNSRecordsSshfpRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordsDNSRecordDNSRecordsSshfpRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordsDNSRecordDNSRecordsSshfpRecordTTLNumber float64

const (
	DNSRecordsDNSRecordDNSRecordsSshfpRecordTTLNumber1 DNSRecordsDNSRecordDNSRecordsSshfpRecordTTLNumber = 1
)

type DNSRecordsDNSRecordDNSRecordsSvcbRecord struct {
	// Components of a SVCB record.
	Data DNSRecordsDNSRecordDNSRecordsSvcbRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordsDNSRecordDNSRecordsSvcbRecordType `json:"type,required"`
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
	Meta DNSRecordsDNSRecordDNSRecordsSvcbRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordsDNSRecordDNSRecordsSvcbRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordsDNSRecordDNSRecordsSvcbRecordJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsSvcbRecordJSON contains the JSON metadata for the
// struct [DNSRecordsDNSRecordDNSRecordsSvcbRecord]
type dnsRecordsDNSRecordDNSRecordsSvcbRecordJSON struct {
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

func (r *DNSRecordsDNSRecordDNSRecordsSvcbRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordsDNSRecordDNSRecordsSvcbRecord) implementsDNSRecordsDNSRecord() {}

// Components of a SVCB record.
type DNSRecordsDNSRecordDNSRecordsSvcbRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                          `json:"value"`
	JSON  dnsRecordsDNSRecordDNSRecordsSvcbRecordDataJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsSvcbRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordsDNSRecordDNSRecordsSvcbRecordData]
type dnsRecordsDNSRecordDNSRecordsSvcbRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordsDNSRecordDNSRecordsSvcbRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordsDNSRecordDNSRecordsSvcbRecordType string

const (
	DNSRecordsDNSRecordDNSRecordsSvcbRecordTypeSvcb DNSRecordsDNSRecordDNSRecordsSvcbRecordType = "SVCB"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordsDNSRecordDNSRecordsSvcbRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordsDNSRecordDNSRecordsSvcbRecordMetaJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsSvcbRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordsDNSRecordDNSRecordsSvcbRecordMeta]
type dnsRecordsDNSRecordDNSRecordsSvcbRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordsDNSRecordDNSRecordsSvcbRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordsDNSRecordDNSRecordsSvcbRecordTTLNumber].
type DNSRecordsDNSRecordDNSRecordsSvcbRecordTTL interface {
	ImplementsDNSRecordsDNSRecordDNSRecordsSvcbRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordsDNSRecordDNSRecordsSvcbRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordsDNSRecordDNSRecordsSvcbRecordTTLNumber float64

const (
	DNSRecordsDNSRecordDNSRecordsSvcbRecordTTLNumber1 DNSRecordsDNSRecordDNSRecordsSvcbRecordTTLNumber = 1
)

type DNSRecordsDNSRecordDNSRecordsTlsaRecord struct {
	// Components of a TLSA record.
	Data DNSRecordsDNSRecordDNSRecordsTlsaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordsDNSRecordDNSRecordsTlsaRecordType `json:"type,required"`
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
	Meta DNSRecordsDNSRecordDNSRecordsTlsaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordsDNSRecordDNSRecordsTlsaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordsDNSRecordDNSRecordsTlsaRecordJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsTlsaRecordJSON contains the JSON metadata for the
// struct [DNSRecordsDNSRecordDNSRecordsTlsaRecord]
type dnsRecordsDNSRecordDNSRecordsTlsaRecordJSON struct {
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

func (r *DNSRecordsDNSRecordDNSRecordsTlsaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordsDNSRecordDNSRecordsTlsaRecord) implementsDNSRecordsDNSRecord() {}

// Components of a TLSA record.
type DNSRecordsDNSRecordDNSRecordsTlsaRecordData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                         `json:"usage"`
	JSON  dnsRecordsDNSRecordDNSRecordsTlsaRecordDataJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsTlsaRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordsDNSRecordDNSRecordsTlsaRecordData]
type dnsRecordsDNSRecordDNSRecordsTlsaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordsDNSRecordDNSRecordsTlsaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordsDNSRecordDNSRecordsTlsaRecordType string

const (
	DNSRecordsDNSRecordDNSRecordsTlsaRecordTypeTlsa DNSRecordsDNSRecordDNSRecordsTlsaRecordType = "TLSA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordsDNSRecordDNSRecordsTlsaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordsDNSRecordDNSRecordsTlsaRecordMetaJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsTlsaRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordsDNSRecordDNSRecordsTlsaRecordMeta]
type dnsRecordsDNSRecordDNSRecordsTlsaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordsDNSRecordDNSRecordsTlsaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordsDNSRecordDNSRecordsTlsaRecordTTLNumber].
type DNSRecordsDNSRecordDNSRecordsTlsaRecordTTL interface {
	ImplementsDNSRecordsDNSRecordDNSRecordsTlsaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordsDNSRecordDNSRecordsTlsaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordsDNSRecordDNSRecordsTlsaRecordTTLNumber float64

const (
	DNSRecordsDNSRecordDNSRecordsTlsaRecordTTLNumber1 DNSRecordsDNSRecordDNSRecordsTlsaRecordTTLNumber = 1
)

type DNSRecordsDNSRecordDNSRecordsTxtRecord struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordsDNSRecordDNSRecordsTxtRecordType `json:"type,required"`
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
	Meta DNSRecordsDNSRecordDNSRecordsTxtRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordsDNSRecordDNSRecordsTxtRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordsDNSRecordDNSRecordsTxtRecordJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsTxtRecordJSON contains the JSON metadata for the
// struct [DNSRecordsDNSRecordDNSRecordsTxtRecord]
type dnsRecordsDNSRecordDNSRecordsTxtRecordJSON struct {
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

func (r *DNSRecordsDNSRecordDNSRecordsTxtRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordsDNSRecordDNSRecordsTxtRecord) implementsDNSRecordsDNSRecord() {}

// Record type.
type DNSRecordsDNSRecordDNSRecordsTxtRecordType string

const (
	DNSRecordsDNSRecordDNSRecordsTxtRecordTypeTxt DNSRecordsDNSRecordDNSRecordsTxtRecordType = "TXT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordsDNSRecordDNSRecordsTxtRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                         `json:"source"`
	JSON   dnsRecordsDNSRecordDNSRecordsTxtRecordMetaJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsTxtRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordsDNSRecordDNSRecordsTxtRecordMeta]
type dnsRecordsDNSRecordDNSRecordsTxtRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordsDNSRecordDNSRecordsTxtRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordsDNSRecordDNSRecordsTxtRecordTTLNumber].
type DNSRecordsDNSRecordDNSRecordsTxtRecordTTL interface {
	ImplementsDNSRecordsDNSRecordDNSRecordsTxtRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordsDNSRecordDNSRecordsTxtRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordsDNSRecordDNSRecordsTxtRecordTTLNumber float64

const (
	DNSRecordsDNSRecordDNSRecordsTxtRecordTTLNumber1 DNSRecordsDNSRecordDNSRecordsTxtRecordTTLNumber = 1
)

type DNSRecordsDNSRecordDNSRecordsUriRecord struct {
	// Components of a URI record.
	Data DNSRecordsDNSRecordDNSRecordsUriRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordsDNSRecordDNSRecordsUriRecordType `json:"type,required"`
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
	Meta DNSRecordsDNSRecordDNSRecordsUriRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordsDNSRecordDNSRecordsUriRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordsDNSRecordDNSRecordsUriRecordJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsUriRecordJSON contains the JSON metadata for the
// struct [DNSRecordsDNSRecordDNSRecordsUriRecord]
type dnsRecordsDNSRecordDNSRecordsUriRecordJSON struct {
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

func (r *DNSRecordsDNSRecordDNSRecordsUriRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordsDNSRecordDNSRecordsUriRecord) implementsDNSRecordsDNSRecord() {}

// Components of a URI record.
type DNSRecordsDNSRecordDNSRecordsUriRecordData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64                                        `json:"weight"`
	JSON   dnsRecordsDNSRecordDNSRecordsUriRecordDataJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsUriRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordsDNSRecordDNSRecordsUriRecordData]
type dnsRecordsDNSRecordDNSRecordsUriRecordDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordsDNSRecordDNSRecordsUriRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordsDNSRecordDNSRecordsUriRecordType string

const (
	DNSRecordsDNSRecordDNSRecordsUriRecordTypeUri DNSRecordsDNSRecordDNSRecordsUriRecordType = "URI"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordsDNSRecordDNSRecordsUriRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                         `json:"source"`
	JSON   dnsRecordsDNSRecordDNSRecordsUriRecordMetaJSON `json:"-"`
}

// dnsRecordsDNSRecordDNSRecordsUriRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordsDNSRecordDNSRecordsUriRecordMeta]
type dnsRecordsDNSRecordDNSRecordsUriRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordsDNSRecordDNSRecordsUriRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordsDNSRecordDNSRecordsUriRecordTTLNumber].
type DNSRecordsDNSRecordDNSRecordsUriRecordTTL interface {
	ImplementsDNSRecordsDNSRecordDNSRecordsUriRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordsDNSRecordDNSRecordsUriRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordsDNSRecordDNSRecordsUriRecordTTLNumber float64

const (
	DNSRecordsDNSRecordDNSRecordsUriRecordTTLNumber1 DNSRecordsDNSRecordDNSRecordsUriRecordTTLNumber = 1
)

type DNSRecordDeleteResponse struct {
	// Identifier
	ID   string                      `json:"id"`
	JSON dnsRecordDeleteResponseJSON `json:"-"`
}

// dnsRecordDeleteResponseJSON contains the JSON metadata for the struct
// [DNSRecordDeleteResponse]
type dnsRecordDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordImportResponse struct {
	// Number of DNS records added.
	RecsAdded float64 `json:"recs_added"`
	// Total number of DNS records parsed.
	TotalRecordsParsed float64                     `json:"total_records_parsed"`
	JSON               dnsRecordImportResponseJSON `json:"-"`
}

// dnsRecordImportResponseJSON contains the JSON metadata for the struct
// [DNSRecordImportResponse]
type dnsRecordImportResponseJSON struct {
	RecsAdded          apijson.Field
	TotalRecordsParsed apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DNSRecordImportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordScanResponse struct {
	// Number of DNS records added.
	RecsAdded float64 `json:"recs_added"`
	// Total number of DNS records parsed.
	TotalRecordsParsed float64                   `json:"total_records_parsed"`
	JSON               dnsRecordScanResponseJSON `json:"-"`
}

// dnsRecordScanResponseJSON contains the JSON metadata for the struct
// [DNSRecordScanResponse]
type dnsRecordScanResponseJSON struct {
	RecsAdded          apijson.Field
	TotalRecordsParsed apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DNSRecordScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordNewParamsType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string]                 `json:"comment"`
	Data    param.Field[DNSRecordNewParamsData] `json:"data"`
	Meta    param.Field[DNSRecordNewParamsMeta] `json:"meta"`
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
	TTL param.Field[DNSRecordNewParamsTTL] `json:"ttl"`
}

func (r DNSRecordNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordNewParamsType string

const (
	DNSRecordNewParamsTypeUri DNSRecordNewParamsType = "URI"
)

type DNSRecordNewParamsData struct {
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
	Flags param.Field[string] `json:"flags"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
	// Degrees of latitude.
	LatDegrees param.Field[float64] `json:"lat_degrees"`
	// Latitude direction.
	LatDirection param.Field[DNSRecordNewParamsDataLatDirection] `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes param.Field[float64] `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds param.Field[float64] `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees param.Field[float64] `json:"long_degrees"`
	// Longitude direction.
	LongDirection param.Field[DNSRecordNewParamsDataLongDirection] `json:"long_direction"`
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

func (r DNSRecordNewParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Latitude direction.
type DNSRecordNewParamsDataLatDirection string

const (
	DNSRecordNewParamsDataLatDirectionN DNSRecordNewParamsDataLatDirection = "N"
	DNSRecordNewParamsDataLatDirectionS DNSRecordNewParamsDataLatDirection = "S"
)

// Longitude direction.
type DNSRecordNewParamsDataLongDirection string

const (
	DNSRecordNewParamsDataLongDirectionE DNSRecordNewParamsDataLongDirection = "E"
	DNSRecordNewParamsDataLongDirectionW DNSRecordNewParamsDataLongDirection = "W"
)

type DNSRecordNewParamsMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded param.Field[bool] `json:"auto_added"`
	// Where the record originated from.
	Source param.Field[string] `json:"source"`
}

func (r DNSRecordNewParamsMeta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat], [DNSRecordNewParamsTTLNumber].
type DNSRecordNewParamsTTL interface {
	ImplementsDNSRecordNewParamsTTL()
}

type DNSRecordNewParamsTTLNumber float64

const (
	DNSRecordNewParamsTTLNumber1 DNSRecordNewParamsTTLNumber = 1
)

type DNSRecordNewResponseEnvelope struct {
	Errors   []DNSRecordNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSRecordNewResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSRecordsDNSRecord                    `json:"result,required"`
	// Whether the API call was successful
	Success DNSRecordNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    dnsRecordNewResponseEnvelopeJSON    `json:"-"`
}

// dnsRecordNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseEnvelope]
type dnsRecordNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordNewResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    dnsRecordNewResponseEnvelopeErrorsJSON `json:"-"`
}

// dnsRecordNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseEnvelopeErrors]
type dnsRecordNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordNewResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    dnsRecordNewResponseEnvelopeMessagesJSON `json:"-"`
}

// dnsRecordNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseEnvelopeMessages]
type dnsRecordNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DNSRecordNewResponseEnvelopeSuccess bool

const (
	DNSRecordNewResponseEnvelopeSuccessTrue DNSRecordNewResponseEnvelopeSuccess = true
)

type DNSRecordUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordUpdateParamsType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string]                    `json:"comment"`
	Data    param.Field[DNSRecordUpdateParamsData] `json:"data"`
	Meta    param.Field[DNSRecordUpdateParamsMeta] `json:"meta"`
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
	TTL param.Field[DNSRecordUpdateParamsTTL] `json:"ttl"`
}

func (r DNSRecordUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordUpdateParamsType string

const (
	DNSRecordUpdateParamsTypeUri DNSRecordUpdateParamsType = "URI"
)

type DNSRecordUpdateParamsData struct {
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
	Flags param.Field[string] `json:"flags"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
	// Degrees of latitude.
	LatDegrees param.Field[float64] `json:"lat_degrees"`
	// Latitude direction.
	LatDirection param.Field[DNSRecordUpdateParamsDataLatDirection] `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes param.Field[float64] `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds param.Field[float64] `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees param.Field[float64] `json:"long_degrees"`
	// Longitude direction.
	LongDirection param.Field[DNSRecordUpdateParamsDataLongDirection] `json:"long_direction"`
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

func (r DNSRecordUpdateParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Latitude direction.
type DNSRecordUpdateParamsDataLatDirection string

const (
	DNSRecordUpdateParamsDataLatDirectionN DNSRecordUpdateParamsDataLatDirection = "N"
	DNSRecordUpdateParamsDataLatDirectionS DNSRecordUpdateParamsDataLatDirection = "S"
)

// Longitude direction.
type DNSRecordUpdateParamsDataLongDirection string

const (
	DNSRecordUpdateParamsDataLongDirectionE DNSRecordUpdateParamsDataLongDirection = "E"
	DNSRecordUpdateParamsDataLongDirectionW DNSRecordUpdateParamsDataLongDirection = "W"
)

type DNSRecordUpdateParamsMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded param.Field[bool] `json:"auto_added"`
	// Where the record originated from.
	Source param.Field[string] `json:"source"`
}

func (r DNSRecordUpdateParamsMeta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat], [DNSRecordUpdateParamsTTLNumber].
type DNSRecordUpdateParamsTTL interface {
	ImplementsDNSRecordUpdateParamsTTL()
}

type DNSRecordUpdateParamsTTLNumber float64

const (
	DNSRecordUpdateParamsTTLNumber1 DNSRecordUpdateParamsTTLNumber = 1
)

type DNSRecordUpdateResponseEnvelope struct {
	Errors   []DNSRecordUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSRecordUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSRecordsDNSRecord                       `json:"result,required"`
	// Whether the API call was successful
	Success DNSRecordUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    dnsRecordUpdateResponseEnvelopeJSON    `json:"-"`
}

// dnsRecordUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseEnvelope]
type dnsRecordUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordUpdateResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    dnsRecordUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// dnsRecordUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DNSRecordUpdateResponseEnvelopeErrors]
type dnsRecordUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordUpdateResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    dnsRecordUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// dnsRecordUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DNSRecordUpdateResponseEnvelopeMessages]
type dnsRecordUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DNSRecordUpdateResponseEnvelopeSuccess bool

const (
	DNSRecordUpdateResponseEnvelopeSuccessTrue DNSRecordUpdateResponseEnvelopeSuccess = true
)

type DNSRecordListParams struct {
	// Identifier
	ZoneID  param.Field[string]                     `path:"zone_id,required"`
	Comment param.Field[DNSRecordListParamsComment] `query:"comment"`
	// DNS record content.
	Content param.Field[string] `query:"content"`
	// Direction to order DNS records in.
	Direction param.Field[DNSRecordListParamsDirection] `query:"direction"`
	// Whether to match all search requirements or at least one (any). If set to `all`,
	// acts like a logical AND between filters. If set to `any`, acts like a logical OR
	// instead. Note that the interaction between tag filters is controlled by the
	// `tag-match` parameter instead.
	Match param.Field[DNSRecordListParamsMatch] `query:"match"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `query:"name"`
	// Field to order DNS records by.
	Order param.Field[DNSRecordListParamsOrder] `query:"order"`
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
	Search param.Field[string]                 `query:"search"`
	Tag    param.Field[DNSRecordListParamsTag] `query:"tag"`
	// Whether to match all tag search requirements or at least one (any). If set to
	// `all`, acts like a logical AND between tag filters. If set to `any`, acts like a
	// logical OR instead. Note that the regular `match` parameter is still used to
	// combine the resulting condition with other filters that aren't related to tags.
	TagMatch param.Field[DNSRecordListParamsTagMatch] `query:"tag_match"`
	// Record type.
	Type param.Field[DNSRecordListParamsType] `query:"type"`
}

// URLQuery serializes [DNSRecordListParams]'s query parameters as `url.Values`.
func (r DNSRecordListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DNSRecordListParamsComment struct {
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

// URLQuery serializes [DNSRecordListParamsComment]'s query parameters as
// `url.Values`.
func (r DNSRecordListParamsComment) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order DNS records in.
type DNSRecordListParamsDirection string

const (
	DNSRecordListParamsDirectionAsc  DNSRecordListParamsDirection = "asc"
	DNSRecordListParamsDirectionDesc DNSRecordListParamsDirection = "desc"
)

// Whether to match all search requirements or at least one (any). If set to `all`,
// acts like a logical AND between filters. If set to `any`, acts like a logical OR
// instead. Note that the interaction between tag filters is controlled by the
// `tag-match` parameter instead.
type DNSRecordListParamsMatch string

const (
	DNSRecordListParamsMatchAny DNSRecordListParamsMatch = "any"
	DNSRecordListParamsMatchAll DNSRecordListParamsMatch = "all"
)

// Field to order DNS records by.
type DNSRecordListParamsOrder string

const (
	DNSRecordListParamsOrderType    DNSRecordListParamsOrder = "type"
	DNSRecordListParamsOrderName    DNSRecordListParamsOrder = "name"
	DNSRecordListParamsOrderContent DNSRecordListParamsOrder = "content"
	DNSRecordListParamsOrderTTL     DNSRecordListParamsOrder = "ttl"
	DNSRecordListParamsOrderProxied DNSRecordListParamsOrder = "proxied"
)

type DNSRecordListParamsTag struct {
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

// URLQuery serializes [DNSRecordListParamsTag]'s query parameters as `url.Values`.
func (r DNSRecordListParamsTag) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Whether to match all tag search requirements or at least one (any). If set to
// `all`, acts like a logical AND between tag filters. If set to `any`, acts like a
// logical OR instead. Note that the regular `match` parameter is still used to
// combine the resulting condition with other filters that aren't related to tags.
type DNSRecordListParamsTagMatch string

const (
	DNSRecordListParamsTagMatchAny DNSRecordListParamsTagMatch = "any"
	DNSRecordListParamsTagMatchAll DNSRecordListParamsTagMatch = "all"
)

// Record type.
type DNSRecordListParamsType string

const (
	DNSRecordListParamsTypeA      DNSRecordListParamsType = "A"
	DNSRecordListParamsTypeAaaa   DNSRecordListParamsType = "AAAA"
	DNSRecordListParamsTypeCaa    DNSRecordListParamsType = "CAA"
	DNSRecordListParamsTypeCert   DNSRecordListParamsType = "CERT"
	DNSRecordListParamsTypeCname  DNSRecordListParamsType = "CNAME"
	DNSRecordListParamsTypeDnskey DNSRecordListParamsType = "DNSKEY"
	DNSRecordListParamsTypeDs     DNSRecordListParamsType = "DS"
	DNSRecordListParamsTypeHTTPS  DNSRecordListParamsType = "HTTPS"
	DNSRecordListParamsTypeLoc    DNSRecordListParamsType = "LOC"
	DNSRecordListParamsTypeMx     DNSRecordListParamsType = "MX"
	DNSRecordListParamsTypeNaptr  DNSRecordListParamsType = "NAPTR"
	DNSRecordListParamsTypeNs     DNSRecordListParamsType = "NS"
	DNSRecordListParamsTypePtr    DNSRecordListParamsType = "PTR"
	DNSRecordListParamsTypeSmimea DNSRecordListParamsType = "SMIMEA"
	DNSRecordListParamsTypeSrv    DNSRecordListParamsType = "SRV"
	DNSRecordListParamsTypeSshfp  DNSRecordListParamsType = "SSHFP"
	DNSRecordListParamsTypeSvcb   DNSRecordListParamsType = "SVCB"
	DNSRecordListParamsTypeTlsa   DNSRecordListParamsType = "TLSA"
	DNSRecordListParamsTypeTxt    DNSRecordListParamsType = "TXT"
	DNSRecordListParamsTypeUri    DNSRecordListParamsType = "URI"
)

type DNSRecordDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type DNSRecordDeleteResponseEnvelope struct {
	Result DNSRecordDeleteResponse             `json:"result"`
	JSON   dnsRecordDeleteResponseEnvelopeJSON `json:"-"`
}

// dnsRecordDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSRecordDeleteResponseEnvelope]
type dnsRecordDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordEditParamsType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string]                  `json:"comment"`
	Data    param.Field[DNSRecordEditParamsData] `json:"data"`
	Meta    param.Field[DNSRecordEditParamsMeta] `json:"meta"`
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
	TTL param.Field[DNSRecordEditParamsTTL] `json:"ttl"`
}

func (r DNSRecordEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordEditParamsType string

const (
	DNSRecordEditParamsTypeUri DNSRecordEditParamsType = "URI"
)

type DNSRecordEditParamsData struct {
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
	Flags param.Field[string] `json:"flags"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
	// Degrees of latitude.
	LatDegrees param.Field[float64] `json:"lat_degrees"`
	// Latitude direction.
	LatDirection param.Field[DNSRecordEditParamsDataLatDirection] `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes param.Field[float64] `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds param.Field[float64] `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees param.Field[float64] `json:"long_degrees"`
	// Longitude direction.
	LongDirection param.Field[DNSRecordEditParamsDataLongDirection] `json:"long_direction"`
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

func (r DNSRecordEditParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Latitude direction.
type DNSRecordEditParamsDataLatDirection string

const (
	DNSRecordEditParamsDataLatDirectionN DNSRecordEditParamsDataLatDirection = "N"
	DNSRecordEditParamsDataLatDirectionS DNSRecordEditParamsDataLatDirection = "S"
)

// Longitude direction.
type DNSRecordEditParamsDataLongDirection string

const (
	DNSRecordEditParamsDataLongDirectionE DNSRecordEditParamsDataLongDirection = "E"
	DNSRecordEditParamsDataLongDirectionW DNSRecordEditParamsDataLongDirection = "W"
)

type DNSRecordEditParamsMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded param.Field[bool] `json:"auto_added"`
	// Where the record originated from.
	Source param.Field[string] `json:"source"`
}

func (r DNSRecordEditParamsMeta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat], [DNSRecordEditParamsTTLNumber].
type DNSRecordEditParamsTTL interface {
	ImplementsDNSRecordEditParamsTTL()
}

type DNSRecordEditParamsTTLNumber float64

const (
	DNSRecordEditParamsTTLNumber1 DNSRecordEditParamsTTLNumber = 1
)

type DNSRecordEditResponseEnvelope struct {
	Errors   []DNSRecordEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSRecordEditResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSRecordsDNSRecord                     `json:"result,required"`
	// Whether the API call was successful
	Success DNSRecordEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    dnsRecordEditResponseEnvelopeJSON    `json:"-"`
}

// dnsRecordEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseEnvelope]
type dnsRecordEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordEditResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    dnsRecordEditResponseEnvelopeErrorsJSON `json:"-"`
}

// dnsRecordEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DNSRecordEditResponseEnvelopeErrors]
type dnsRecordEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordEditResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    dnsRecordEditResponseEnvelopeMessagesJSON `json:"-"`
}

// dnsRecordEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DNSRecordEditResponseEnvelopeMessages]
type dnsRecordEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DNSRecordEditResponseEnvelopeSuccess bool

const (
	DNSRecordEditResponseEnvelopeSuccessTrue DNSRecordEditResponseEnvelopeSuccess = true
)

type DNSRecordExportParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type DNSRecordGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type DNSRecordGetResponseEnvelope struct {
	Errors   []DNSRecordGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSRecordGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSRecordsDNSRecord                    `json:"result,required"`
	// Whether the API call was successful
	Success DNSRecordGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dnsRecordGetResponseEnvelopeJSON    `json:"-"`
}

// dnsRecordGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseEnvelope]
type dnsRecordGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordGetResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    dnsRecordGetResponseEnvelopeErrorsJSON `json:"-"`
}

// dnsRecordGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseEnvelopeErrors]
type dnsRecordGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordGetResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    dnsRecordGetResponseEnvelopeMessagesJSON `json:"-"`
}

// dnsRecordGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseEnvelopeMessages]
type dnsRecordGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DNSRecordGetResponseEnvelopeSuccess bool

const (
	DNSRecordGetResponseEnvelopeSuccessTrue DNSRecordGetResponseEnvelopeSuccess = true
)

type DNSRecordImportParams struct {
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

func (r DNSRecordImportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordImportResponseEnvelope struct {
	Errors   []DNSRecordImportResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSRecordImportResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSRecordImportResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DNSRecordImportResponseEnvelopeSuccess `json:"success,required"`
	Timing  DNSRecordImportResponseEnvelopeTiming  `json:"timing"`
	JSON    dnsRecordImportResponseEnvelopeJSON    `json:"-"`
}

// dnsRecordImportResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSRecordImportResponseEnvelope]
type dnsRecordImportResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	Timing      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordImportResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordImportResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    dnsRecordImportResponseEnvelopeErrorsJSON `json:"-"`
}

// dnsRecordImportResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DNSRecordImportResponseEnvelopeErrors]
type dnsRecordImportResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordImportResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordImportResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    dnsRecordImportResponseEnvelopeMessagesJSON `json:"-"`
}

// dnsRecordImportResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DNSRecordImportResponseEnvelopeMessages]
type dnsRecordImportResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordImportResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DNSRecordImportResponseEnvelopeSuccess bool

const (
	DNSRecordImportResponseEnvelopeSuccessTrue DNSRecordImportResponseEnvelopeSuccess = true
)

type DNSRecordImportResponseEnvelopeTiming struct {
	// When the file parsing ended.
	EndTime time.Time `json:"end_time" format:"date-time"`
	// Processing time of the file in seconds.
	ProcessTime float64 `json:"process_time"`
	// When the file parsing started.
	StartTime time.Time                                 `json:"start_time" format:"date-time"`
	JSON      dnsRecordImportResponseEnvelopeTimingJSON `json:"-"`
}

// dnsRecordImportResponseEnvelopeTimingJSON contains the JSON metadata for the
// struct [DNSRecordImportResponseEnvelopeTiming]
type dnsRecordImportResponseEnvelopeTimingJSON struct {
	EndTime     apijson.Field
	ProcessTime apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordImportResponseEnvelopeTiming) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordScanParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type DNSRecordScanResponseEnvelope struct {
	Errors   []DNSRecordScanResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSRecordScanResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSRecordScanResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DNSRecordScanResponseEnvelopeSuccess `json:"success,required"`
	Timing  DNSRecordScanResponseEnvelopeTiming  `json:"timing"`
	JSON    dnsRecordScanResponseEnvelopeJSON    `json:"-"`
}

// dnsRecordScanResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSRecordScanResponseEnvelope]
type dnsRecordScanResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	Timing      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordScanResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordScanResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    dnsRecordScanResponseEnvelopeErrorsJSON `json:"-"`
}

// dnsRecordScanResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DNSRecordScanResponseEnvelopeErrors]
type dnsRecordScanResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordScanResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordScanResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    dnsRecordScanResponseEnvelopeMessagesJSON `json:"-"`
}

// dnsRecordScanResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DNSRecordScanResponseEnvelopeMessages]
type dnsRecordScanResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordScanResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DNSRecordScanResponseEnvelopeSuccess bool

const (
	DNSRecordScanResponseEnvelopeSuccessTrue DNSRecordScanResponseEnvelopeSuccess = true
)

type DNSRecordScanResponseEnvelopeTiming struct {
	// When the file parsing ended.
	EndTime time.Time `json:"end_time" format:"date-time"`
	// Processing time of the file in seconds.
	ProcessTime float64 `json:"process_time"`
	// When the file parsing started.
	StartTime time.Time                               `json:"start_time" format:"date-time"`
	JSON      dnsRecordScanResponseEnvelopeTimingJSON `json:"-"`
}

// dnsRecordScanResponseEnvelopeTimingJSON contains the JSON metadata for the
// struct [DNSRecordScanResponseEnvelopeTiming]
type dnsRecordScanResponseEnvelopeTimingJSON struct {
	EndTime     apijson.Field
	ProcessTime apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordScanResponseEnvelopeTiming) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
